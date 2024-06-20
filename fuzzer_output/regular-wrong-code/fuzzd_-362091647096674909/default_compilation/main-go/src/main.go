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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-739), (func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(121)
		}
		return (_dafny.MultiSetOf(false)).Cardinality()
	})())
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cww"), _dafny.UnicodeSeqOfUtf8Bytes("g")), (func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("iekbdyjkt")
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-653))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))
	})())
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	return !((((_dafny.SetOf(true)).Cardinality()).Cmp(_dafny.IntOfInt64(-923)) >= 0) || (true))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Contains(false), !(false), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvsmrmtk")).Cardinality())).Cmp(_dafny.IntOfInt64(-139)) == 0)
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_dafny.IntOfInt64(391)).Plus(_dafny.IntOfInt64(820))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 D0, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D0_.Create_DC1_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(553), (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("rejvchf"))).Cardinality())).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), Companion_D0_.Create_DC1_((_dafny.MultiSetOf(true, false)).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-482))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kd")).Cardinality()))).Minus(_dafny.IntOfInt64(74)))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.Zero).Minus(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-377), _dafny.IntOfInt64(741))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-377)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(741)) < 0) {
				_coll0.Add((_1_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("levb")).Cardinality())), (_dafny.Zero).Minus((_dafny.SetOf(true, false, false)).Cardinality()))
			}
		}
		return _coll0.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(393), _dafny.IntOfInt64(-101))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(393)).Cmp(_2_v1) <= 0) && ((_2_v1).Cmp(_dafny.IntOfInt64(-101)) < 0) {
				_coll1.Add((_2_v1).Plus((_dafny.SetOf(true, true)).Cardinality()), _dafny.IntOfInt64(-736))
			}
		}
		return _coll1.ToMap()
	}())).Cardinality()), (_dafny.IntOfInt64(-318)).Times(_dafny.IntOfInt64(841)))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(true))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 D0, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Set, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	})), _dafny.UnicodeSeqOfUtf8Bytes("vtjpii")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nnq")))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(Companion_D5_.Create_DC13_((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("epbtsmm")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("upj")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_4_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-734)
		}))).Cardinality()))).Cardinality(), false, false)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ju")).Cardinality()), false, true)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfInt64(-703), true, false), Companion_D5_.Create_DC13_((_dafny.MultiSetOf(_dafny.CodePoint('m'))).Cardinality(), true, false)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(194))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_5_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()), true, true)))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v0 _dafny.Sequence
			_6_v0 = interface{}(_compr_2).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(Companion_D5_.Create_DC13_((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("epbtsmm")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("upj")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_4_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-734)
			}))).Cardinality()))).Cardinality(), false, false)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ju")).Cardinality()), false, true)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfInt64(-703), true, false), Companion_D5_.Create_DC13_((_dafny.MultiSetOf(_dafny.CodePoint('m'))).Cardinality(), true, false)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(194))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_5_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))).Cardinality()), true, true)))).Contains(_6_v0) {
				_coll2.Add(_6_v0)
			}
		}
		return _coll2.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.SeqOf(Companion_D5_.Create_DC13_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('n'))).Cardinality()), true, false)), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfInt64(-505), true, !(false))), _dafny.SeqOf(Companion_D5_.Create_DC13_(_dafny.IntOfInt64(953), true, false))))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 bool, p3 _dafny.Set, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(!((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("teghdwfgg")).Cardinality()))).Cardinality()))).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dhgmuljeq")).Cardinality())))), Companion_Default___.SafeDivisionInt((_dafny.SetOf(false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Cardinality()), (func() bool {
		if true {
			return true
		}
		return !(true)
	})())
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D0 = Companion_D0_.Create_DC0_(_dafny.MultiSetOf(false, false, false))
	_ = _source0
	if _source0.Is_DC1() {
		var _7___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _7___mcc_h0
		var _8_cf1 _dafny.Int = _7___mcc_h0
		_ = _8_cf1
		return _dafny.CodePoint('u')
	} else if _source0.Is_DC2() {
		var _9___mcc_h1 bool = _source0.Get_().(D0_DC2).Cf2
		_ = _9___mcc_h1
		var _10___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC2).Cf3
		_ = _10___mcc_h2
		var _11___mcc_h3 bool = _source0.Get_().(D0_DC2).Cf4
		_ = _11___mcc_h3
		var _12_cf4 bool = _11___mcc_h3
		_ = _12_cf4
		var _13_cf3 _dafny.Int = _10___mcc_h2
		_ = _13_cf3
		var _14_cf2 bool = _9___mcc_h1
		_ = _14_cf2
		return _dafny.CodePoint('h')
	} else if _source0.Is_DC0() {
		var _15___mcc_h4 _dafny.MultiSet = _source0.Get_().(D0_DC0).Cf0
		_ = _15___mcc_h4
		var _16_cf0 _dafny.MultiSet = _15___mcc_h4
		_ = _16_cf0
		return _dafny.CodePoint('k')
	} else {
		var _17___mcc_h5 D0 = _source0.Get_().(D0_DC3).Cf5
		_ = _17___mcc_h5
		var _18_cf5 D0 = _17___mcc_h5
		_ = _18_cf5
		return _dafny.CodePoint('w')
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 D2, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_19_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-681)
	}))).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(-454)), _dafny.MultiSetOf(_dafny.IntOfInt64(-972)), _dafny.MultiSetOf(_dafny.IntOfInt64(756), _dafny.IntOfInt64(399), (func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_20_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("umygtlqs")).Cardinality()), _dafny.IntOfInt64(503))).Cardinality())
		}))).Cardinality()))).Cardinality()), false))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _21_v0 _dafny.Map
			_21_v0 = interface{}(_compr_3).(_dafny.Map)
			if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_20_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("umygtlqs")).Cardinality()), _dafny.IntOfInt64(503))).Cardinality())
			}))).Cardinality()))).Cardinality()), false))).Contains(_21_v0) {
				_coll3.Add(_21_v0)
			}
		}
		return _coll3.ToSet()
	}()).Cardinality()), _dafny.MultiSetOf(_dafny.IntOfInt64(817), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true)).Cardinality()))).Intersection((_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(529), _dafny.IntOfInt64(926)))).Union(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-43))), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _22_v1 _dafny.MultiSet
			_22_v1 = interface{}(_compr_4).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-43))), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))), _22_v1) {
				_coll4.Add(_22_v1)
			}
		}
		return _coll4.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 bool, p2 bool, p3 D0, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(554), _dafny.IntOfInt64(634))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(554)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(634)) < 0) {
				_coll5.Add((_23_v0).Plus(_dafny.IntOfInt64(707)), !(false))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality()))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true), true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_24_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('q'), _dafny.CodePoint('v'))).Cardinality()))
	})))).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iytbot")).Cardinality()))).Intersection(_dafny.MultiSetOf((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(593), _dafny.IntOfInt64(-150))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _25_v1 _dafny.Int
			_25_v1 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(593)).Cmp(_25_v1) <= 0) && ((_25_v1).Cmp(_dafny.IntOfInt64(-150)) < 0) {
				_coll6.Add((_25_v1).Times(_dafny.IntOfInt64(680)), _dafny.IntOfInt64(551))
			}
		}
		return _coll6.ToMap()
	}()).Cardinality())), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-732))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_26_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(false))).Cardinality()
	}))))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('g')
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(745))).Difference((_dafny.SetOf(_dafny.IntOfInt64(69))).Union(_dafny.SetOf(_dafny.IntOfInt64(184), _dafny.IntOfInt64(786), _dafny.IntOfInt64(16))))
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((false) == (true), !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wqt")).Cardinality())).Cmp(_dafny.IntOfInt64(695)) <= 0), true, true)
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	var _source1 D21 = Companion_D21_.Create_DC50_(_dafny.MultiSetOf(Companion_D9_.Create_DC20_(_dafny.UnicodeSeqOfUtf8Bytes("ere"), true, _dafny.IntOfInt64(-948)), Companion_D9_.Create_DC20_(_dafny.UnicodeSeqOfUtf8Bytes("nmnxpj"), true, _dafny.IntOfInt64(189))))
	_ = _source1
	if _source1.Is_DC51() {
		var _27___mcc_h0 bool = _source1.Get_().(D21_DC51).Cf81
		_ = _27___mcc_h0
		var _28_cf81 bool = _27___mcc_h0
		_ = _28_cf81
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rlotprgxp")).Cardinality()), _dafny.SeqOf(true))).Merge(func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-654), _dafny.IntOfInt64(-939))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _29_v0 _dafny.Int
				_29_v0 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(-654)).Cmp(_29_v0) <= 0) && ((_29_v0).Cmp(_dafny.IntOfInt64(-939)) < 0) {
					_coll7.Add((_29_v0).Times(_dafny.IntOfInt64(-920)), _dafny.SeqOf(_28_cf81))
				}
			}
			return _coll7.ToMap()
		}())
	} else if _source1.Is_DC50() {
		var _30___mcc_h1 _dafny.MultiSet = _source1.Get_().(D21_DC50).Cf80
		_ = _30___mcc_h1
		var _31_cf80 _dafny.MultiSet = _30___mcc_h1
		_ = _31_cf80
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-855), _dafny.SeqOf(!(false)))
	} else {
		var _32___mcc_h2 D21 = _source1.Get_().(D21_DC52).Cf82
		_ = _32___mcc_h2
		var _33_cf82 D21 = _32___mcc_h2
		_ = _33_cf82
		return func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(201), _dafny.IntOfInt64(952))); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _34_v1 _dafny.Int
				_34_v1 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.IntOfInt64(201)).Cmp(_34_v1) <= 0) && ((_34_v1).Cmp(_dafny.IntOfInt64(952)) < 0) {
					_coll8.Add((_34_v1).Plus(_dafny.IntOfInt64(637)), _dafny.SeqOf(true, false, true))
				}
			}
			return _coll8.ToMap()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC7_(!(false), !((_dafny.SetOf(_dafny.IntOfInt64(664))).IsSubsetOf(_dafny.SetOf((_dafny.SetOf(_dafny.SetOf(false, false), _dafny.SetOf(false))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, globalState *GlobalState) D6 {
	var _source2 D5 = Companion_D5_.Create_DC14_(!(false))
	_ = _source2
	if _source2.Is_DC13() {
		var _35___mcc_h0 _dafny.Int = _source2.Get_().(D5_DC13).Cf20
		_ = _35___mcc_h0
		var _36___mcc_h1 bool = _source2.Get_().(D5_DC13).Cf21
		_ = _36___mcc_h1
		var _37___mcc_h2 bool = _source2.Get_().(D5_DC13).Cf22
		_ = _37___mcc_h2
		var _38_cf22 bool = _37___mcc_h2
		_ = _38_cf22
		var _39_cf21 bool = _36___mcc_h1
		_ = _39_cf21
		var _40_cf20 _dafny.Int = _35___mcc_h0
		_ = _40_cf20
		return Companion_D6_.Create_DC15_(func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(153), _dafny.IntOfInt64(-772))); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _41_v0 _dafny.Int
				_41_v0 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(153)).Cmp(_41_v0) <= 0) && ((_41_v0).Cmp(_dafny.IntOfInt64(-772)) < 0) {
					_coll9.Add((_41_v0).Plus(_40_cf20), _dafny.SeqOf(_39_cf21))
				}
			}
			return _coll9.ToMap()
		}())
	} else if _source2.Is_DC14() {
		var _42___mcc_h3 bool = _source2.Get_().(D5_DC14).Cf23
		_ = _42___mcc_h3
		var _43_cf23 bool = _42___mcc_h3
		_ = _43_cf23
		return Companion_D6_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(490), _dafny.SeqOf(_43_cf23)))
	} else {
		var _44___mcc_h4 _dafny.Set = _source2.Get_().(D5_DC12).Cf19
		_ = _44___mcc_h4
		var _45_cf19 _dafny.Set = _44___mcc_h4
		_ = _45_cf19
		return Companion_D6_.Create_DC15_(func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(555), _dafny.IntOfInt64(490))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _46_v1 _dafny.Int
				_46_v1 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(555)).Cmp(_46_v1) <= 0) && ((_46_v1).Cmp(_dafny.IntOfInt64(490)) < 0) {
					_coll10.Add((_46_v1).Plus(_dafny.IntOfInt64(355)), _dafny.SeqOf(true, false, true, false))
				}
			}
			return _coll10.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC13_(((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vsqstwekh")).Cardinality())))).Plus(_dafny.IntOfInt64(-868)), false, true)
}
func (_static *CompanionStruct_Default___) Fm41(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false)).Union((_dafny.SetOf(true, !(true), false, !(false), false)).Intersection(_dafny.SetOf(!(true), true, true)))
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.IntOfInt64(898), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-776))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_47_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}))).Cardinality()))
		}
		return _dafny.MultiSetOf(_dafny.IntOfInt64(827))
	})()).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(-34)))
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.MultiSetOf(false, false, false, false))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, globalState *GlobalState) D4 {
	if false {
		return Companion_D4_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})))
	} else {
		return Companion_D4_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_49_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 D2, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_50_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('r'))).Cardinality())
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(89))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_51_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kqu")).Cardinality())
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ehybm")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(485))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_52_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(588)
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-665), _dafny.IntOfInt64(731))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(413), _dafny.IntOfInt64(65))
	})()
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(542), false))).Cardinality()), _dafny.IntOfInt64(396)), _dafny.SeqOf(_dafny.IntOfInt64(-939), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm48(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("uarmdrl"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("lpsc")))
}
func (_static *CompanionStruct_Default___) Fm49(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if !(!(!(true))) {
			return true
		}
		return true
	})(), _dafny.IntOfInt64(657))
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('u'))))).Union(_dafny.MultiSetOf(Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D13_.Create_DC29_(false, false)).Dtor_cf49(), _dafny.CodePoint('i'))), Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('y'))), Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(true))), _dafny.CodePoint('c')))))
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll11 = _dafny.NewBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(((func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(297), _dafny.IntOfInt64(936))); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _53_v1 _dafny.Int
					_53_v1 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(297)).Cmp(_53_v1) <= 0) && ((_53_v1).Cmp(_dafny.IntOfInt64(936)) < 0) {
						_coll13.Add(Companion_Default___.SafeModuloInt(_53_v1, _dafny.IntOfInt64(-866)), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(292))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality())
					}
				}
				return _coll13.ToMap()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (Companion_D11_.Create_DC23_(!(true), _dafny.IntOfInt64(-234), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), Companion_D5_.Create_DC14_(false))).Dtor_cf37()))).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _54_v0 _dafny.Map
				_54_v0 = interface{}(_compr_12).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(297), _dafny.IntOfInt64(936))); ; {
						_compr_14, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _53_v1 _dafny.Int
						_53_v1 = interface{}(_compr_14).(_dafny.Int)
						if ((_dafny.IntOfInt64(297)).Cmp(_53_v1) <= 0) && ((_53_v1).Cmp(_dafny.IntOfInt64(936)) < 0) {
							_coll14.Add(Companion_Default___.SafeModuloInt(_53_v1, _dafny.IntOfInt64(-866)), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(292))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality())
						}
					}
					return _coll14.ToMap()
				}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (Companion_D11_.Create_DC23_(!(true), _dafny.IntOfInt64(-234), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), Companion_D5_.Create_DC14_(false))).Dtor_cf37())), _54_v0) {
					_coll12.Add(_54_v0, _dafny.UnicodeSeqOfUtf8Bytes("ccyc"))
				}
			}
			return _coll12.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(973)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-461))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_55_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))))).Keys().Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _56_v2 _dafny.Map
			_56_v2 = interface{}(_compr_11).(_dafny.Map)
			if ((func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(297), _dafny.IntOfInt64(936))); ; {
						_compr_16, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _53_v1 _dafny.Int
						_53_v1 = interface{}(_compr_16).(_dafny.Int)
						if ((_dafny.IntOfInt64(297)).Cmp(_53_v1) <= 0) && ((_53_v1).Cmp(_dafny.IntOfInt64(936)) < 0) {
							_coll16.Add(Companion_Default___.SafeModuloInt(_53_v1, _dafny.IntOfInt64(-866)), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(292))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality())
						}
					}
					return _coll16.ToMap()
				}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (Companion_D11_.Create_DC23_(!(true), _dafny.IntOfInt64(-234), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), Companion_D5_.Create_DC14_(false))).Dtor_cf37()))).Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _54_v0 _dafny.Map
					_54_v0 = interface{}(_compr_15).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
						var _coll17 = _dafny.NewMapBuilder()
						_ = _coll17
						for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(297), _dafny.IntOfInt64(936))); ; {
							_compr_17, _ok17 := _iter17()
							if !_ok17 {
								break
							}
							var _53_v1 _dafny.Int
							_53_v1 = interface{}(_compr_17).(_dafny.Int)
							if ((_dafny.IntOfInt64(297)).Cmp(_53_v1) <= 0) && ((_53_v1).Cmp(_dafny.IntOfInt64(936)) < 0) {
								_coll17.Add(Companion_Default___.SafeModuloInt(_53_v1, _dafny.IntOfInt64(-866)), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(292))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality())
							}
						}
						return _coll17.ToMap()
					}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (Companion_D11_.Create_DC23_(!(true), _dafny.IntOfInt64(-234), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), Companion_D5_.Create_DC14_(false))).Dtor_cf37())), _54_v0) {
						_coll15.Add(_54_v0, _dafny.UnicodeSeqOfUtf8Bytes("ccyc"))
					}
				}
				return _coll15.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(973)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-461))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_55_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))))).Contains(_56_v2) {
				_coll11.Add(_56_v2)
			}
		}
		return _coll11.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm52(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('b'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.CodePoint('f')))
}
func (_static *CompanionStruct_Default___) Fm53(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ttv"), _dafny.UnicodeSeqOfUtf8Bytes("ctxpl"), _dafny.UnicodeSeqOfUtf8Bytes("ljm"))).Union(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(505))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_57_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Set, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll18 = _dafny.NewBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((func() _dafny.Map {
			var _coll19 = _dafny.NewMapBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SetOf((func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Elements()); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _59_v1 _dafny.Int
					_59_v1 = interface{}(_compr_20).(_dafny.Int)
					if (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg21 _dafny.Int) interface{} {
							return coer21(arg21)
						}
					}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('p')
					}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Contains(_59_v1) {
						_coll20.Add((_59_v1).Times(_dafny.IntOfInt64(666)), !(true))
					}
				}
				return _coll20.ToMap()
			}()).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(296)), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(true))).Cardinality())))).Elements()); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _60_v0 _dafny.Set
				_60_v0 = interface{}(_compr_19).(_dafny.Set)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SetOf((func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter21 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('p')
					}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Elements()); ; {
						_compr_21, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _59_v1 _dafny.Int
						_59_v1 = interface{}(_compr_21).(_dafny.Int)
						if (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg23 _dafny.Int) interface{} {
								return coer23(arg23)
							}
						}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Contains(_59_v1) {
							_coll21.Add((_59_v1).Times(_dafny.IntOfInt64(666)), !(true))
						}
					}
					return _coll21.ToMap()
				}()).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(296)), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(true))).Cardinality()))), _60_v0) {
					_coll19.Add(_60_v0, _dafny.UnicodeSeqOfUtf8Bytes("bgfnoiq"))
				}
			}
			return _coll19.ToMap()
		}()).Keys().Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _61_v2 _dafny.Set
			_61_v2 = interface{}(_compr_18).(_dafny.Set)
			if (func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SetOf((func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter23 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg24 _dafny.Int) interface{} {
							return coer24(arg24)
						}
					}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('p')
					}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Elements()); ; {
						_compr_23, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _59_v1 _dafny.Int
						_59_v1 = interface{}(_compr_23).(_dafny.Int)
						if (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg25 _dafny.Int) interface{} {
								return coer25(arg25)
							}
						}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Contains(_59_v1) {
							_coll23.Add((_59_v1).Times(_dafny.IntOfInt64(666)), !(true))
						}
					}
					return _coll23.ToMap()
				}()).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(296)), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(true))).Cardinality())))).Elements()); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _60_v0 _dafny.Set
					_60_v0 = interface{}(_compr_22).(_dafny.Set)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SetOf((func() _dafny.Map {
						var _coll24 = _dafny.NewMapBuilder()
						_ = _coll24
						for _iter24 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg26 _dafny.Int) interface{} {
								return coer26(arg26)
							}
						}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Elements()); ; {
							_compr_24, _ok24 := _iter24()
							if !_ok24 {
								break
							}
							var _59_v1 _dafny.Int
							_59_v1 = interface{}(_compr_24).(_dafny.Int)
							if (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg27 _dafny.Int) interface{} {
									return coer27(arg27)
								}
							}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('p')
							}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(47))).Cardinality())).Contains(_59_v1) {
								_coll24.Add((_59_v1).Times(_dafny.IntOfInt64(666)), !(true))
							}
						}
						return _coll24.ToMap()
					}()).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(296)), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(true))).Cardinality()))), _60_v0) {
						_coll22.Add(_60_v0, _dafny.UnicodeSeqOfUtf8Bytes("bgfnoiq"))
					}
				}
				return _coll22.ToMap()
			}()).Contains(_61_v2) {
				_coll18.Add(_61_v2)
			}
		}
		return _coll18.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(8))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}(func(_62_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('f'))
	}))
}
func (_static *CompanionStruct_Default___) Fm56(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll25 = _dafny.NewMapBuilder()
		_ = _coll25
		for _iter25 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(784), (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(-19)), _dafny.SeqOf(_dafny.IntOfInt64(428)), _dafny.SeqOf((_dafny.SetOf(true, true, false)).Cardinality(), (_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(437)))).Cardinality())).Keys().Elements()); ; {
			_compr_25, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _63_v0 _dafny.Int
			_63_v0 = interface{}(_compr_25).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(784), (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(-19)), _dafny.SeqOf(_dafny.IntOfInt64(428)), _dafny.SeqOf((_dafny.SetOf(true, true, false)).Cardinality(), (_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(437)))).Cardinality())).Contains(_63_v0) {
				_coll25.Add(Companion_Default___.SafeDivisionInt(_63_v0, _dafny.IntOfInt64(581)), true)
			}
		}
		return _coll25.ToMap()
	}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, true)).Cardinality(), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll26 = _dafny.NewMapBuilder()
		_ = _coll26
		for _iter26 := _dafny.Iterate((_dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality(), (func() _dafny.Set {
			var _coll27 = _dafny.NewBuilder()
			_ = _coll27
			for _iter27 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(88), true)).Keys().Elements()); ; {
				_compr_27, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _64_v2 _dafny.Int
				_64_v2 = interface{}(_compr_27).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(88), true)).Contains(_64_v2) {
					_coll27.Add((_64_v2).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ss")).Cardinality()))).Cardinality())))))
				}
			}
			return _coll27.ToSet()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wqn")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dltidiom")).Cardinality()), _dafny.IntOfInt64(399))).Elements()); ; {
			_compr_26, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _65_v1 _dafny.Int
			_65_v1 = interface{}(_compr_26).(_dafny.Int)
			if (_dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality(), (func() _dafny.Set {
				var _coll28 = _dafny.NewBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(88), true)).Keys().Elements()); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _66_v2 _dafny.Int
					_66_v2 = interface{}(_compr_28).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(88), true)).Contains(_66_v2) {
						_coll28.Add((_66_v2).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ss")).Cardinality()))).Cardinality())))))
					}
				}
				return _coll28.ToSet()
			}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wqn")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dltidiom")).Cardinality()), _dafny.IntOfInt64(399))).Contains(_65_v1) {
				_coll26.Add((_65_v1).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(185))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}(func(_67_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(148), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfInt64(10), _dafny.IntOfInt64(813), _dafny.IntOfInt64(-879)))
			}
		}
		return _coll26.ToMap()
	}()).Cardinality(), true)))
}
func (_static *CompanionStruct_Default___) Fm57(globalState *GlobalState) D19 {
	return Companion_D19_.Create_DC47_((_dafny.SetOf(!(false), false)).IsSubsetOf(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm58(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(true), false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(true), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(false), false)))
}
func (_static *CompanionStruct_Default___) Fm59(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(655), _dafny.IntOfInt64(-349)))).Merge((func() _dafny.Map {
		var _coll29 = _dafny.NewMapBuilder()
		_ = _coll29
		for _iter29 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(false)), true)).Keys().Elements()); ; {
			_compr_29, _ok29 := _iter29()
			if !_ok29 {
				break
			}
			var _68_v0 _dafny.Set
			_68_v0 = interface{}(_compr_29).(_dafny.Set)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(false)), true)).Contains(_68_v0) {
				_coll29.Add(_68_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-697), _dafny.IntOfInt64(-69)))
			}
		}
		return _coll29.ToMap()
	}()).Merge((Companion_D31_.Create_DC68_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), func() _dafny.Map {
		var _coll30 = _dafny.NewMapBuilder()
		_ = _coll30
		for _iter30 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(109), !(false))).Cardinality())).Elements()); ; {
			_compr_30, _ok30 := _iter30()
			if !_ok30 {
				break
			}
			var _69_v1 _dafny.Int
			_69_v1 = interface{}(_compr_30).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(109), !(false))).Cardinality()), _69_v1) {
				_coll30.Add((_69_v1).Times(_dafny.IntOfInt64(-893)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qubdos")).Cardinality()))
			}
		}
		return _coll30.ToMap()
	}()))).Dtor_cf112()))
}
func (_static *CompanionStruct_Default___) Fm60(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(219))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg30 _dafny.Int) interface{} {
			return coer30(arg30)
		}
	}(func(_70_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(486)
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(449))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg31 _dafny.Int) interface{} {
			return coer31(arg31)
		}
	}(func(_71_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}(func(_72_i2 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-574)
		}))
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(78))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer33(arg33)
		}
	}(func(_73_i3 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(337)
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(974))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg34 _dafny.Int) interface{} {
			return coer34(arg34)
		}
	}(func(_74_i4 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfInt64(-185), (_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfInt64(734), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(913))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_75_i5 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(781)
		}))).Cardinality()), _dafny.IntOfInt64(-302))
	}))))
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Int, p1 _dafny.Sequence, p2 D0, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll31 = _dafny.NewBuilder()
		_ = _coll31
		for _iter31 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-252))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_76_i0 _dafny.Int) _dafny.Int {
			return (func() _dafny.Set {
				var _coll32 = _dafny.NewBuilder()
				_ = _coll32
				for _iter32 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(-448), _dafny.IntOfInt64(124))).Elements()); ; {
					_compr_32, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _77_v0 _dafny.Int
					_77_v0 = interface{}(_compr_32).(_dafny.Int)
					if (_dafny.SetOf(_dafny.IntOfInt64(-448), _dafny.IntOfInt64(124))).Contains(_77_v0) {
						_coll32.Add((_77_v0).Plus(_dafny.IntOfInt64(212)))
					}
				}
				return _coll32.ToSet()
			}()).Cardinality()
		}))).Cardinality())))).Elements()); ; {
			_compr_31, _ok31 := _iter31()
			if !_ok31 {
				break
			}
			var _78_v1 _dafny.Sequence
			_78_v1 = interface{}(_compr_31).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-252))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}(func(_76_i0 _dafny.Int) _dafny.Int {
				return (func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter33 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(-448), _dafny.IntOfInt64(124))).Elements()); ; {
						_compr_33, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _79_v0 _dafny.Int
						_79_v0 = interface{}(_compr_33).(_dafny.Int)
						if (_dafny.SetOf(_dafny.IntOfInt64(-448), _dafny.IntOfInt64(124))).Contains(_79_v0) {
							_coll33.Add((_79_v0).Plus(_dafny.IntOfInt64(212)))
						}
					}
					return _coll33.ToSet()
				}()).Cardinality()
			}))).Cardinality())))).Contains(_78_v1) {
				_coll31.Add(_78_v1)
			}
		}
		return _coll31.ToSet()
	}()).Union(_dafny.SetOf(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("faysnx")).Cardinality()), _dafny.IntOfInt64(-59), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-141))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg38 _dafny.Int) interface{} {
			return coer38(arg38)
		}
	}(func(_80_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), _dafny.IntOfInt64(494))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-772), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-478))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg39 _dafny.Int) interface{} {
			return coer39(arg39)
		}
	}(func(_81_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qs")).Cardinality()), _dafny.IntOfInt64(995))).Cardinality()))).Cardinality()))))).Difference((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(88)), _dafny.SeqOf((_dafny.SetOf(false, true, false)).Cardinality()), _dafny.SeqOf((_dafny.MultiSetOf(false, true, false, false)).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(-711), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eavsgwtj")).Cardinality())), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.IntOfInt64(-41))).Cardinality()), _dafny.IntOfInt64(753)))).Union(func() _dafny.Set {
		var _coll34 = _dafny.NewBuilder()
		_ = _coll34
		for _iter34 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(837)), Companion_D5_.Create_DC14_(false))).Keys().Elements()); ; {
			_compr_34, _ok34 := _iter34()
			if !_ok34 {
				break
			}
			var _82_v2 _dafny.Sequence
			_82_v2 = interface{}(_compr_34).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(837)), Companion_D5_.Create_DC14_(false))).Contains(_82_v2) {
				_coll34.Add(_82_v2)
			}
		}
		return _coll34.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm62(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source3 D4 = Companion_D4_.Create_DC11_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())))
	_ = _source3
	if _source3.Is_DC11() {
		var _83___mcc_h0 _dafny.Int = _source3.Get_().(D4_DC11).Cf18
		_ = _83___mcc_h0
		var _84_cf18 _dafny.Int = _83___mcc_h0
		_ = _84_cf18
		return func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter35 := _dafny.Iterate((func() _dafny.Set {
				var _coll36 = _dafny.NewBuilder()
				_ = _coll36
				for _iter36 := _dafny.Iterate((func() _dafny.Map {
					var _coll37 = _dafny.NewMapBuilder()
					_ = _coll37
					for _iter37 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Keys().Elements()); ; {
						_compr_37, _ok37 := _iter37()
						if !_ok37 {
							break
						}
						var _85_v1 _dafny.Sequence
						_85_v1 = interface{}(_compr_37).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Contains(_85_v1) {
							_coll37.Add(_85_v1, _84_cf18)
						}
					}
					return _coll37.ToMap()
				}()).Keys().Elements()); ; {
					_compr_36, _ok36 := _iter36()
					if !_ok36 {
						break
					}
					var _86_v2 _dafny.Sequence
					_86_v2 = interface{}(_compr_36).(_dafny.Sequence)
					if (func() _dafny.Map {
						var _coll38 = _dafny.NewMapBuilder()
						_ = _coll38
						for _iter38 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Keys().Elements()); ; {
							_compr_38, _ok38 := _iter38()
							if !_ok38 {
								break
							}
							var _85_v1 _dafny.Sequence
							_85_v1 = interface{}(_compr_38).(_dafny.Sequence)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Contains(_85_v1) {
								_coll38.Add(_85_v1, _84_cf18)
							}
						}
						return _coll38.ToMap()
					}()).Contains(_86_v2) {
						_coll36.Add(_86_v2)
					}
				}
				return _coll36.ToSet()
			}()).Elements()); ; {
				_compr_35, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _87_v0 _dafny.Sequence
				_87_v0 = interface{}(_compr_35).(_dafny.Sequence)
				if (func() _dafny.Set {
					var _coll39 = _dafny.NewBuilder()
					_ = _coll39
					for _iter39 := _dafny.Iterate((func() _dafny.Map {
						var _coll40 = _dafny.NewMapBuilder()
						_ = _coll40
						for _iter40 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Keys().Elements()); ; {
							_compr_40, _ok40 := _iter40()
							if !_ok40 {
								break
							}
							var _85_v1 _dafny.Sequence
							_85_v1 = interface{}(_compr_40).(_dafny.Sequence)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Contains(_85_v1) {
								_coll40.Add(_85_v1, _84_cf18)
							}
						}
						return _coll40.ToMap()
					}()).Keys().Elements()); ; {
						_compr_39, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _88_v2 _dafny.Sequence
						_88_v2 = interface{}(_compr_39).(_dafny.Sequence)
						if (func() _dafny.Map {
							var _coll41 = _dafny.NewMapBuilder()
							_ = _coll41
							for _iter41 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Keys().Elements()); ; {
								_compr_41, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _85_v1 _dafny.Sequence
								_85_v1 = interface{}(_compr_41).(_dafny.Sequence)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Contains(_85_v1) {
									_coll41.Add(_85_v1, _84_cf18)
								}
							}
							return _coll41.ToMap()
						}()).Contains(_88_v2) {
							_coll39.Add(_88_v2)
						}
					}
					return _coll39.ToSet()
				}()).Contains(_87_v0) {
					_coll35.Add(_87_v0, true)
				}
			}
			return _coll35.ToMap()
		}()
	} else {
		var _89___mcc_h1 _dafny.Sequence = _source3.Get_().(D4_DC10).Cf17
		_ = _89___mcc_h1
		var _90_cf17 _dafny.Sequence = _89___mcc_h1
		_ = _90_cf17
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true))
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var _91_v0 _dafny.Array
	_ = _91_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
	_ = _nw0
	_91_v0 = _nw0
	var _92_v1 _dafny.Map
	_ = _92_v1
	_92_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v0, (p0).Cmp(p3) != 0)
	if (func() bool {
		if (_92_v1).Contains(_91_v0) {
			return (_92_v1).Get(_91_v0).(bool)
		}
		return !(p2) || (!(p2))
	})() {
		var _93_v2 _dafny.Set
		_ = _93_v2
		_93_v2 = _dafny.SetOf(false, p2)
		var _94_v3 _dafny.Map
		_ = _94_v3
		_94_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v2, p0)
		var _95_v4 _dafny.Sequence
		_ = _95_v4
		_95_v4 = _dafny.SeqOf(p3)
		var _96_v5 _dafny.Sequence
		_ = _96_v5
		_96_v5 = _dafny.SeqOf(p2, p2)
		var _97_v6 _dafny.Map
		_ = _97_v6
		_97_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
		var _98_v7 _dafny.Map
		_ = _98_v7
		_98_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_95_v4).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_96_v5, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_97_v6).Cardinality()), _dafny.IntOfUint32((_96_v5).Cardinality()))).Uint32(), false)).Cardinality()))
		_94_v3 = (_94_v3).Update(_dafny.SetOf(p2), (func() _dafny.Int {
			if (_97_v6).Contains(_dafny.IntOfInt64(-908)) {
				return (_97_v6).Get(_dafny.IntOfInt64(-908)).(_dafny.Int)
			}
			return p0
		})())
		var _99_v8 _dafny.Array
		_ = _99_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_100_p2 bool) func(_dafny.Int) bool {
				return func(_101_i0 _dafny.Int) bool {
					return _100_p2
				}
			})(p2)
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
		_99_v8 = _nw1
		var _102_v9 _dafny.Sequence
		_ = _102_v9
		_102_v9 = _dafny.SeqOf(_99_v8)
		var _103_v10 D27
		_ = _103_v10
		_103_v10 = Companion_D27_.Create_DC63_((p1).Cardinality(), p0, p0, p2)
		var _104_v11 _dafny.Sequence
		_ = _104_v11
		_104_v11 = _dafny.SeqOf(_95_v4, _95_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(985))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_105_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_106_i2 _dafny.Int) _dafny.Int {
				return _105_p0
			}
		})(p0))), _95_v4)
		var _107_v12 _dafny.Map
		_ = _107_v12
		_107_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_103_v10).Dtor_cf107(), _104_v11)
		var _108_v13 D1
		_ = _108_v13
		_108_v13 = Companion_D1_.Create_DC5_(p3, p0, (func() _dafny.Sequence {
			if (_107_v12).Contains(true) {
				return (_107_v12).Get(true).(_dafny.Sequence)
			}
			return _104_v11
		})(), _99_v8, p0)
		var _109_v14 _dafny.Set
		_ = _109_v14
		_109_v14 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)).Cardinality())
		var _110_v15 _dafny.Map
		_ = _110_v15
		_110_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _99_v8)
		var _pat_let_tv0 = _104_v11
		_ = _pat_let_tv0
		var _pat_let_tv1 = p3
		_ = _pat_let_tv1
		var _111_v16 _dafny.Array
		_ = _111_v16
		var _nwElement0_0 D1 = Companion_D1_.Create_DC5_(p3, p3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(844))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}((func(_112_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_113_i1 _dafny.Int) _dafny.Sequence {
				return _112_v4
			}
		})(_95_v4))), _99_v8, _dafny.IntOfInt64(685))
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(17))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_0, 0)
		(_nw2).ArraySet1(_108_v13, 1)
		(_nw2).ArraySet1(_108_v13, 2)
		(_nw2).ArraySet1(_108_v13, 3)
		(_nw2).ArraySet1(func(_pat_let0_0 D1) D1 {
			return func(_114_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 _dafny.Sequence) D1 {
					return func(_115_dt__update_hcf9_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC5_((_114_dt__update__tmp_h0).Dtor_cf7(), (_114_dt__update__tmp_h0).Dtor_cf8(), _115_dt__update_hcf9_h0, (_114_dt__update__tmp_h0).Dtor_cf10(), (_114_dt__update__tmp_h0).Dtor_cf11())
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_108_v13), 4)
		(_nw2).ArraySet1(_108_v13, 5)
		(_nw2).ArraySet1((func() D1 {
			if p2 {
				return _108_v13
			}
			return _108_v13
		})(), 6)
		(_nw2).ArraySet1((func() D1 {
			if p2 {
				return _108_v13
			}
			return Companion_D1_.Create_DC5_(p0, p3, _104_v11, _99_v8, (_dafny.Zero).Minus(p3))
		})(), 7)
		(_nw2).ArraySet1(Companion_D1_.Create_DC5_(p3, _dafny.IntOfInt64(806), _104_v11, _99_v8, p0), 8)
		(_nw2).ArraySet1(Companion_D1_.Create_DC5_(p0, p3, _dafny.SeqOf(_dafny.SeqOf((p1).Cardinality(), p3, p0, p3, (_109_v14).Cardinality())), (func() _dafny.Array {
			if (_110_v15).Contains(p2) {
				return (_110_v15).Get(p2).(_dafny.Array)
			}
			return _99_v8
		})(), p3), 9)
		(_nw2).ArraySet1(Companion_D1_.Create_DC5_(p3, p3, _104_v11, _99_v8, p0), 10)
		(_nw2).ArraySet1(_108_v13, 11)
		(_nw2).ArraySet1(_108_v13, 12)
		(_nw2).ArraySet1(_108_v13, 13)
		(_nw2).ArraySet1(_108_v13, 14)
		(_nw2).ArraySet1(_108_v13, 15)
		(_nw2).ArraySet1(func(_pat_let2_0 D1) D1 {
			return func(_116_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let3_0 _dafny.Int) D1 {
					return func(_117_dt__update_hcf11_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC5_((_116_dt__update__tmp_h1).Dtor_cf7(), (_116_dt__update__tmp_h1).Dtor_cf8(), (_116_dt__update__tmp_h1).Dtor_cf9(), (_116_dt__update__tmp_h1).Dtor_cf10(), _117_dt__update_hcf11_h0)
					}(_pat_let3_0)
				}(_pat_let_tv1)
			}(_pat_let2_0)
		}(_108_v13), 16)
		_111_v16 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_111_v16), 0))
		_ = _index0
		(_111_v16).ArraySet1(_108_v13, (_index0).Int())
		var _118_v17 _dafny.Sequence
		_ = _118_v17
		_118_v17 = _dafny.SeqOf(_dafny.SeqOf(_99_v8), _102_v9, _102_v9, _102_v9, _102_v9)
		var _119_v18 _dafny.Map
		_ = _119_v18
		_119_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_118_v17).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_118_v17).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _120_v19 _dafny.Sequence
		_ = _120_v19
		_120_v19 = _dafny.UnicodeSeqOfUtf8Bytes("tlsyfds")
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_111_v16), 0))
		_ = _index1
		var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_102_v9, (func() _dafny.Sequence {
			if (_119_v18).Contains(_dafny.IntOfUint32((_120_v19).Cardinality())) {
				return (_119_v18).Get(_dafny.IntOfUint32((_120_v19).Cardinality())).(_dafny.Sequence)
			}
			return _102_v9
		})())
		_ = _rhs0
		var _rhs1 D1 = _108_v13
		_ = _rhs1
		var _lhs0 _dafny.Array = _111_v16
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_111_v16), 0))
		_ = _lhs1
		_102_v9 = _rhs0
		(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
		var _121_v20 D0
		_ = _121_v20
		_121_v20 = Companion_D0_.Create_DC0_(p1)
		var _122_v21 D0
		_ = _122_v21
		_122_v21 = Companion_D0_.Create_DC0_((p1).Union((_121_v20).Dtor_cf0()))
		var _source4 D0 = _122_v21
		_ = _source4
		if _source4.Is_DC1() {
			var _123___mcc_h0 _dafny.Int = _source4.Get_().(D0_DC1).Cf1
			_ = _123___mcc_h0
			var _124_cf1 _dafny.Int = _123___mcc_h0
			_ = _124_cf1
			var _125_v22 D9
			_ = _125_v22
			_125_v22 = Companion_D9_.Create_DC20_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(86))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}(func(_126_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), p2, p3)
			var _127_v23 _dafny.MultiSet
			_ = _127_v23
			_127_v23 = _dafny.MultiSetOf(_125_v22)
			var _128_v24 D21
			_ = _128_v24
			_128_v24 = Companion_D21_.Create_DC50_(_127_v23)
			var _129_v25 _dafny.Map
			_ = _129_v25
			_129_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_96_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_96_v5).Cardinality()))).Uint32()).(bool), p2)
			var _130_v26 D5
			_ = _130_v26
			_130_v26 = Companion_D5_.Create_DC14_(p2)
			var _131_v27 D11
			_ = _131_v27
			_131_v27 = Companion_D11_.Create_DC23_(p2, _124_cf1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), _130_v26)
			var _132_v28 _dafny.Array
			_ = _132_v28
			var _nwElement0_1 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			_ = _nwElement0_1
			var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(15))
			_ = _nw3
			(_nw3).ArraySet1(_nwElement0_1, 0)
			(_nw3).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), 1)
			(_nw3).ArraySet1(_129_v25, 2)
			(_nw3).ArraySet1(_129_v25, 3)
			(_nw3).ArraySet1(_129_v25, 4)
			(_nw3).ArraySet1(_129_v25, 5)
			(_nw3).ArraySet1(_129_v25, 6)
			(_nw3).ArraySet1(_129_v25, 7)
			(_nw3).ArraySet1(_129_v25, 8)
			(_nw3).ArraySet1((_131_v27).Dtor_cf38(), 9)
			(_nw3).ArraySet1(_129_v25, 10)
			(_nw3).ArraySet1((_131_v27).Dtor_cf38(), 11)
			(_nw3).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_96_v5).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_96_v5).Cardinality()))).Uint32()).(bool), p2), 12)
			(_nw3).ArraySet1(_129_v25, 13)
			(_nw3).ArraySet1(_129_v25, 14)
			_132_v28 = _nw3
			var _133_v29 _dafny.MultiSet
			_ = _133_v29
			_133_v29 = _dafny.MultiSetOf(_120_v19)
			var _134_v30 _dafny.MultiSet
			_ = _134_v30
			_134_v30 = _133_v29
			r0 = (_124_cf1).Minus((_dafny.Zero).Minus((Companion_D23_.Create_DC55_(_128_v24, _132_v28, _134_v30, p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Cardinality())).Dtor_cf89()))
			var _pat_let_tv2 = _93_v2
			_ = _pat_let_tv2
			var _pat_let_tv3 = _93_v2
			_ = _pat_let_tv3
			var _rhs2 bool = p2
			_ = _rhs2
			var _rhs3 D5 = func(_pat_let4_0 D5) D5 {
				return func(_135_dt__update__tmp_h2 D5) D5 {
					return func(_pat_let5_0 bool) D5 {
						return func(_136_dt__update_hcf23_h0 bool) D5 {
							return Companion_D5_.Create_DC14_(_136_dt__update_hcf23_h0)
						}(_pat_let5_0)
					}((_pat_let_tv2).IsSubsetOf(_pat_let_tv3))
				}(_pat_let4_0)
			}(_130_v26)
			_ = _rhs3
			var _rhs4 _dafny.Sequence = _120_v19
			_ = _rhs4
			var _rhs5 _dafny.Int = Companion_Default___.SafeDivisionInt((p3).Minus(_dafny.IntOfUint32((_96_v5).Cardinality())), p0)
			_ = _rhs5
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			_lhs2.F13 = _rhs2
			_130_v26 = _rhs3
			_120_v19 = _rhs4
			r0 = _rhs5
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_99_v8), 0))
			_ = _index2
			(_99_v8).ArraySet1((!(p2)) || (p2), (_index2).Int())
			var _137_v31 _dafny.Array
			_ = _137_v31
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
			_ = _nw4
			_137_v31 = _nw4
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_99_v8), 0))
			_ = _index3
			var _rhs6 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-374))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_138_v19 _dafny.Sequence, _139_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_140_i4 _dafny.Int) _dafny.CodePoint {
					return (_138_v19).Select((Companion_Default___.SafeIndex(_139_p0, _dafny.IntOfUint32((_138_v19).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_120_v19, p0))), _120_v19), _dafny.Companion_Sequence_.Concatenate(_120_v19, _dafny.UnicodeSeqOfUtf8Bytes("cnqxtfmaj"))), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(p3, _124_cf1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-374))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_141_v19 _dafny.Sequence, _142_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_143_i4 _dafny.Int) _dafny.CodePoint {
					return (_141_v19).Select((Companion_Default___.SafeIndex(_142_p0, _dafny.IntOfUint32((_141_v19).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_120_v19, p0))), _120_v19), _dafny.Companion_Sequence_.Concatenate(_120_v19, _dafny.UnicodeSeqOfUtf8Bytes("cnqxtfmaj")))).Cardinality()))).Uint32(), _dafny.CodePoint('q'))
			_ = _rhs6
			var _rhs7 _dafny.Int = (_124_cf1).Plus(p0)
			_ = _rhs7
			var _rhs8 bool = (func() bool {
				if p2 {
					return p2
				}
				return p2
			})()
			_ = _rhs8
			var _rhs9 bool = p2
			_ = _rhs9
			var _rhs10 _dafny.Array = _137_v31
			_ = _rhs10
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			var _lhs4 _dafny.Array = _99_v8
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_99_v8), 0))
			_ = _lhs5
			_120_v19 = _rhs6
			_lhs3.F7 = _rhs7
			r1 = _rhs8
			(_lhs4).ArraySet1(_rhs9, (_lhs5).Int())
			_137_v31 = _rhs10
			var _144_v32 _dafny.CodePoint
			_ = _144_v32
			_144_v32 = _dafny.CodePoint('q')
			var _145_v33 T1
			_ = _145_v33
			var _nw5 *C6 = New_C6_()
			_ = _nw5
			_nw5.Ctor__(_144_v32, (_96_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_96_v5).Cardinality()))).Uint32()).(bool))
			_145_v33 = _nw5
			var _146_v34 _dafny.Array
			_ = _146_v34
			var _nwElement0_2 _dafny.CodePoint = _145_v33.F23()
			_ = _nwElement0_2
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
			_ = _nw6
			(_nw6).ArraySet1CodePoint(_nwElement0_2, 0)
			(_nw6).ArraySet1CodePoint(_145_v33.F23(), 1)
			(_nw6).ArraySet1CodePoint(_144_v32, 2)
			(_nw6).ArraySet1CodePoint(_145_v33.F23(), 3)
			_146_v34 = _nw6
			var _nw7 *C10 = New_C10_()
			_ = _nw7
			_nw7.Ctor__(p2, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(872), p0), _144_v32, p2, _146_v34, Companion_Default___.Fm2(globalState))
			_145_v33 = _nw7
		} else if _source4.Is_DC2() {
			var _147___mcc_h1 bool = _source4.Get_().(D0_DC2).Cf2
			_ = _147___mcc_h1
			var _148___mcc_h2 _dafny.Int = _source4.Get_().(D0_DC2).Cf3
			_ = _148___mcc_h2
			var _149___mcc_h3 bool = _source4.Get_().(D0_DC2).Cf4
			_ = _149___mcc_h3
			var _150_cf4 bool = _149___mcc_h3
			_ = _150_cf4
			var _151_cf3 _dafny.Int = _148___mcc_h2
			_ = _151_cf3
			var _152_cf2 bool = _147___mcc_h1
			_ = _152_cf2
			var _153_v35 _dafny.Array
			_ = _153_v35
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw8
			_153_v35 = _nw8
			var _154_v36 _dafny.Sequence
			_ = _154_v36
			_154_v36 = _dafny.SeqOf(_153_v35, _153_v35)
			var _155_v37 D18
			_ = _155_v37
			_155_v37 = Companion_D18_.Create_DC42_((p3).Minus(p3), _dafny.Companion_Sequence_.Contains(_154_v36, _153_v35))
			_155_v37 = _155_v37
			var _156_v38 _dafny.Map
			_ = _156_v38
			_156_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_120_v19, _120_v19), p0)
			_156_v38 = (_156_v38).Update((_152_cf2) == (false), p0)
			var _157_v39 _dafny.Array
			_ = _157_v39
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw9
			_157_v39 = _nw9
			var _158_v40 _dafny.Map
			_ = _158_v40
			_158_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_cf4, _152_cf2)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_157_v39), 0))
			_ = _index4
			(_157_v39).ArraySet1(_158_v40, (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_157_v39), 0))
			_ = _index5
			(_157_v39).ArraySet1(_158_v40, (_index5).Int())
			var _159_v41 D19
			_ = _159_v41
			_159_v41 = Companion_D19_.Create_DC47_(true)
			var _160_v42 D19
			_ = _160_v42
			_160_v42 = Companion_D19_.Create_DC48_(_159_v41)
			var _161_v43 D19
			_ = _161_v43
			_161_v43 = Companion_D19_.Create_DC48_(_160_v42)
			var _pat_let_tv4 = _160_v42
			_ = _pat_let_tv4
			_161_v43 = func(_pat_let6_0 D19) D19 {
				return func(_162_dt__update__tmp_h3 D19) D19 {
					return func(_pat_let7_0 D19) D19 {
						return func(_163_dt__update_hcf78_h0 D19) D19 {
							return Companion_D19_.Create_DC48_(_163_dt__update_hcf78_h0)
						}(_pat_let7_0)
					}(_pat_let_tv4)
				}(_pat_let6_0)
			}(_161_v43)
		} else if _source4.Is_DC0() {
			var _164___mcc_h4 _dafny.MultiSet = _source4.Get_().(D0_DC0).Cf0
			_ = _164___mcc_h4
			var _165_cf0 _dafny.MultiSet = _164___mcc_h4
			_ = _165_cf0
			var _166_v44 _dafny.CodePoint
			_ = _166_v44
			_166_v44 = _dafny.CodePoint('f')
			var _167_v45 _dafny.MultiSet
			_ = _167_v45
			_167_v45 = _dafny.MultiSetOf(_166_v44)
			_167_v45 = _167_v45
			var _168_v46 _dafny.Array
			_ = _168_v46
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
			_ = _nw10
			_168_v46 = _nw10
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_168_v46), 0))
			_ = _index6
			(_168_v46).ArraySet1CodePoint(_166_v44, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_168_v46), 0))
			_ = _index7
			(_168_v46).ArraySet1CodePoint(_166_v44, (_index7).Int())
			var _169_v47 _dafny.MultiSet
			_ = _169_v47
			_169_v47 = _dafny.MultiSetOf(_120_v19, _120_v19)
			var _170_v48 _dafny.Map
			_ = _170_v48
			_170_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _169_v47)
			_169_v47 = (func() _dafny.MultiSet {
				if (_170_v48).Contains((p3).Plus(p0)) {
					return (_170_v48).Get((p3).Plus(p0)).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_120_v19)
			})()
			var _171_v49 _dafny.Sequence
			_ = _171_v49
			_171_v49 = _dafny.SeqOf(_120_v19, _120_v19)
			_120_v19 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_120_v19, _120_v19), (_171_v49).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_171_v49).Cardinality()))).Uint32()).(_dafny.Sequence))
		} else {
			var _172___mcc_h5 D0 = _source4.Get_().(D0_DC3).Cf5
			_ = _172___mcc_h5
			var _173_cf5 D0 = _172___mcc_h5
			_ = _173_cf5
			var _174_v50 _dafny.CodePoint
			_ = _174_v50
			_174_v50 = _dafny.CodePoint('v')
			var _175_v51 T0
			_ = _175_v51
			var _nw11 *C0 = New_C0_()
			_ = _nw11
			_nw11.Ctor__(Companion_Default___.Fm48(globalState), p2, _174_v50, p2)
			_175_v51 = _nw11
			var _176_v52 _dafny.Array
			_ = _176_v52
			var _nwElement0_3 T0 = _175_v51
			_ = _nwElement0_3
			var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(16))
			_ = _nw12
			(_nw12).ArraySet1(_nwElement0_3, 0)
			(_nw12).ArraySet1(_175_v51, 1)
			(_nw12).ArraySet1(_175_v51, 2)
			(_nw12).ArraySet1(_175_v51, 3)
			(_nw12).ArraySet1(_175_v51, 4)
			(_nw12).ArraySet1(_175_v51, 5)
			(_nw12).ArraySet1(_175_v51, 6)
			(_nw12).ArraySet1(_175_v51, 7)
			(_nw12).ArraySet1(_175_v51, 8)
			(_nw12).ArraySet1(_175_v51, 9)
			(_nw12).ArraySet1(_175_v51, 10)
			(_nw12).ArraySet1(_175_v51, 11)
			(_nw12).ArraySet1((_175_v51), 12)
			(_nw12).ArraySet1(_175_v51, 13)
			(_nw12).ArraySet1(_175_v51, 14)
			(_nw12).ArraySet1(_175_v51, 15)
			_176_v52 = _nw12
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_176_v52), 0))
			_ = _index8
			(_176_v52).ArraySet1(_175_v51, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_176_v52), 0))
			_ = _index9
			(_176_v52).ArraySet1(_175_v51, (_index9).Int())
			var _177_v53 _dafny.Array
			_ = _177_v53
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
			_ = _nw13
			_177_v53 = _nw13
			var _178_v55 _dafny.Set
			_ = _178_v55
			_178_v55 = _dafny.SetOf(_175_v51.F23())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_177_v53), 0))
			_ = _index10
			(_177_v53).ArraySet1((func() _dafny.Set {
				var _coll42 = _dafny.NewBuilder()
				_ = _coll42
				for _iter42 := _dafny.Iterate((func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter43 := _dafny.Iterate((_178_v55).Elements()); ; {
						_compr_43, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _179_v54 _dafny.CodePoint
						_179_v54 = interface{}(_compr_43).(_dafny.CodePoint)
						if (_178_v55).Contains(_179_v54) {
							_coll43.Add(_179_v54, !((_175_v51).F24()))
						}
					}
					return _coll43.ToMap()
				}()).Keys().Elements()); ; {
					_compr_42, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _180_v56 _dafny.CodePoint
					_180_v56 = interface{}(_compr_42).(_dafny.CodePoint)
					if (func() _dafny.Map {
						var _coll44 = _dafny.NewMapBuilder()
						_ = _coll44
						for _iter44 := _dafny.Iterate((_178_v55).Elements()); ; {
							_compr_44, _ok44 := _iter44()
							if !_ok44 {
								break
							}
							var _179_v54 _dafny.CodePoint
							_179_v54 = interface{}(_compr_44).(_dafny.CodePoint)
							if (_178_v55).Contains(_179_v54) {
								_coll44.Add(_179_v54, !((_175_v51).F24()))
							}
						}
						return _coll44.ToMap()
					}()).Contains(_180_v56) {
						_coll42.Add(_180_v56)
					}
				}
				return _coll42.ToSet()
			}()).Intersection(_178_v55), (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_177_v53), 0))
			_ = _index11
			(_177_v53).ArraySet1(_178_v55, (_index11).Int())
			var _181_v57 _dafny.Array
			_ = _181_v57
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_1
			var _nw14 _dafny.Array
			_ = _nw14
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw14 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.CodePoint = (func(_182_v51 T0) func(_dafny.Int) _dafny.CodePoint {
					return func(_183_i5 _dafny.Int) _dafny.CodePoint {
						return _182_v51.F23()
					}
				})(_175_v51)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw14 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw14).ArraySet1CodePoint(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw14).ArraySet1CodePoint(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_181_v57 = _nw14
			var _184_v58 *C7
			_ = _184_v58
			var _nw15 *C7 = New_C7_()
			_ = _nw15
			_nw15.Ctor__((_175_v51).F24(), true, _181_v57, _175_v51.F23(), false, !((_dafny.MultiSetOf(p0, _dafny.IntOfInt64(448), p0)).IsProperSubsetOf(_dafny.MultiSetFromSeq(_95_v4))))
			_184_v58 = _nw15
			_184_v58 = _184_v58
			(globalState).F13 = (_184_v58).F31()
		}
		var _185_v59 _dafny.Map
		_ = _185_v59
		_185_v59 = _97_v6
		var _186_v60 _dafny.Array
		_ = _186_v60
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_2
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.CodePoint = func(_187_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw16 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw16).ArraySet1CodePoint(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw16).ArraySet1CodePoint(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_186_v60 = _nw16
		var _188_v61 _dafny.Map
		_ = _188_v61
		_188_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _186_v60)
		var _189_v62 _dafny.CodePoint
		_ = _189_v62
		_189_v62 = _dafny.CodePoint('s')
		var _190_v63 *C4
		_ = _190_v63
		var _nw17 *C4 = New_C4_()
		_ = _nw17
		_nw17.Ctor__(_dafny.IntOfUint32((_120_v19).Cardinality()), _185_v59, (func() _dafny.Array {
			if (_188_v61).Contains(false) {
				return (_188_v61).Get(false).(_dafny.Array)
			}
			return _186_v60
		})(), _189_v62, p2)
		_190_v63 = _nw17
		(globalState).F13 = (func() bool {
			if (_190_v63).Fm22(p0, p2, globalState) {
				return p2
			}
			return p2
		})()
	} else {
		r2 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-599), (Companion_Default___.Fm62(p0, globalState)).Cardinality())
		r0 = p0
		var _191_v64 _dafny.Array
		_ = _191_v64
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw18
		_191_v64 = _nw18
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))
		_ = _index12
		(_191_v64).ArraySet1(false, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))
		_ = _index13
		(_191_v64).ArraySet1(p2, (_index13).Int())
		var _192_v65 _dafny.Sequence
		_ = _192_v65
		_192_v65 = _dafny.SeqOf(!((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool)))
		var _193_v66 _dafny.Map
		_ = _193_v66
		_193_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _192_v65)
		var _194_v67 _dafny.Set
		_ = _194_v67
		_194_v67 = _dafny.SetOf(_dafny.IntOfInt64(425))
		var _195_v68 _dafny.Sequence
		_ = _195_v68
		_195_v68 = _dafny.UnicodeSeqOfUtf8Bytes("k")
		var _196_v69 _dafny.Map
		_ = _196_v69
		_196_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pfiap")).Cardinality()))
		var _197_v70 _dafny.Sequence
		_ = _197_v70
		_197_v70 = _dafny.SeqOf(p3, p3, p0, _dafny.IntOfInt64(141), p3)
		var _198_v71 _dafny.Set
		_ = _198_v71
		_198_v71 = _dafny.SetOf((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool), (_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool))
		var _199_v72 _dafny.MultiSet
		_ = _199_v72
		_199_v72 = _dafny.MultiSetOf(_198_v71)
		var _200_v74 _dafny.Array
		_ = _200_v74
		var _nwElement0_4 _dafny.Int = (p3).Times((_193_v66).Cardinality())
		_ = _nwElement0_4
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(29))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_4, 0)
		(_nw19).ArraySet1((_194_v67).Cardinality(), 1)
		(_nw19).ArraySet1(p3, 2)
		(_nw19).ArraySet1(_dafny.IntOfInt64(-354), 3)
		(_nw19).ArraySet1(_dafny.IntOfInt64(-37), 4)
		(_nw19).ArraySet1(Companion_Default___.SafeModuloInt(p3, p3), 5)
		(_nw19).ArraySet1(p3, 6)
		(_nw19).ArraySet1(_dafny.IntOfUint32((_195_v68).Cardinality()), 7)
		(_nw19).ArraySet1(_dafny.IntOfInt64(641), 8)
		(_nw19).ArraySet1(((_196_v69).Cardinality()).Minus((func() _dafny.Int {
			if (_196_v69).Contains(!((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool))) {
				return (_196_v69).Get(!((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool))).(_dafny.Int)
			}
			return p3
		})()), 9)
		(_nw19).ArraySet1(p3, 10)
		(_nw19).ArraySet1(_dafny.IntOfUint32((_197_v70).Cardinality()), 11)
		(_nw19).ArraySet1(Companion_Default___.SafeDivisionInt(p0, p0), 12)
		(_nw19).ArraySet1((func() _dafny.Int {
			if true {
				return p3
			}
			return (_dafny.Zero).Minus(p3)
		})(), 13)
		(_nw19).ArraySet1(p3, 14)
		(_nw19).ArraySet1((func() _dafny.Int {
			if (_199_v72).Contains(_198_v71) {
				return (_199_v72).Multiplicity(_198_v71)
			}
			return _dafny.IntOfInt64(184)
		})(), 15)
		(_nw19).ArraySet1(p3, 16)
		(_nw19).ArraySet1(p0, 17)
		(_nw19).ArraySet1(p0, 18)
		(_nw19).ArraySet1(p3, 19)
		(_nw19).ArraySet1(p0, 20)
		(_nw19).ArraySet1(p3, 21)
		(_nw19).ArraySet1(p0, 22)
		(_nw19).ArraySet1(p3, 23)
		(_nw19).ArraySet1(p3, 24)
		(_nw19).ArraySet1(_dafny.IntOfInt64(-424), 25)
		(_nw19).ArraySet1(p3, 26)
		(_nw19).ArraySet1(((func() _dafny.Set {
			var _coll45 = _dafny.NewBuilder()
			_ = _coll45
			for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(649), _dafny.IntOfInt64(62))); ; {
				_compr_45, _ok45 := _iter45()
				if !_ok45 {
					break
				}
				var _201_v73 _dafny.Int
				_201_v73 = interface{}(_compr_45).(_dafny.Int)
				if ((_dafny.IntOfInt64(649)).Cmp(_201_v73) <= 0) && ((_201_v73).Cmp(_dafny.IntOfInt64(62)) < 0) {
					_coll45.Add(Companion_Default___.SafeDivisionInt(_201_v73, _dafny.IntOfInt64(234)))
				}
			}
			return _coll45.ToSet()
		}()).Intersection(_194_v67)).Cardinality(), 27)
		(_nw19).ArraySet1(p0, 28)
		_200_v74 = _nw19
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_200_v74), 0))
		_ = _index14
		(_200_v74).ArraySet1(p0, (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_200_v74), 0))
		_ = _index15
		var _rhs11 _dafny.Int = p0
		_ = _rhs11
		var _rhs12 _dafny.Int = (p0).Times(p3)
		_ = _rhs12
		var _lhs6 _dafny.Array = _200_v74
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_200_v74), 0))
		_ = _lhs7
		var _lhs8 *GlobalState = globalState
		_ = _lhs8
		(_lhs6).ArraySet1(_rhs11, (_lhs7).Int())
		_lhs8.F7 = _rhs12
		var _202_v75 _dafny.Array
		_ = _202_v75
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
		_ = _nw20
		_202_v75 = _nw20
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_202_v75), 0))
		_ = _index16
		(_202_v75).ArraySet1(_198_v71, (_index16).Int())
		var _203_v76 _dafny.Sequence
		_ = _203_v76
		_203_v76 = _dafny.SeqOf(_198_v71)
		var _204_v77 _dafny.CodePoint
		_ = _204_v77
		_204_v77 = _dafny.CodePoint('h')
		var _205_v78 _dafny.Map
		_ = _205_v78
		_205_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool), p2)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_202_v75), 0))
		_ = _index17
		var _rhs13 bool = Companion_Default___.Fm2(globalState)
		_ = _rhs13
		var _rhs14 _dafny.Set = ((_198_v71).Union(_198_v71)).Intersection(((_203_v76).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(p2, (_dafny.SetOf(_204_v77)).Cardinality(), globalState), _dafny.IntOfUint32((_203_v76).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_dafny.SetOf((func() bool {
			if (_205_v78).Contains((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool)) {
				return (_205_v78).Get((_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool)).(bool)
			}
			return (_191_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_191_v64), 0))).Int()).(bool)
		})(), p2)))
		_ = _rhs14
		var _lhs9 _dafny.Array = _202_v75
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_202_v75), 0))
		_ = _lhs10
		r1 = _rhs13
		(_lhs9).ArraySet1(_rhs14, (_lhs10).Int())
	}
	var _206_v79 _dafny.CodePoint
	_ = _206_v79
	_206_v79 = _dafny.CodePoint('a')
	var _207_v80 *C6
	_ = _207_v80
	var _nw21 *C6 = New_C6_()
	_ = _nw21
	_nw21.Ctor__(_206_v79, true)
	_207_v80 = _nw21
	var _208_v81 D23
	_ = _208_v81
	_208_v81 = Companion_D23_.Create_DC54_(_207_v80)
	_207_v80 = (_208_v81).Dtor_cf84()
	var _209_v82 _dafny.Sequence
	_ = _209_v82
	_209_v82 = _dafny.UnicodeSeqOfUtf8Bytes("yj")
	var _210_v83 D9
	_ = _210_v83
	_210_v83 = Companion_D9_.Create_DC20_(_209_v82, p2, p0)
	var _211_v84 _dafny.MultiSet
	_ = _211_v84
	_211_v84 = _dafny.MultiSetOf(_210_v83, _210_v83, _210_v83)
	var _212_v85 D21
	_ = _212_v85
	_212_v85 = Companion_D21_.Create_DC50_(_211_v84)
	var _213_v86 _dafny.Sequence
	_ = _213_v86
	_213_v86 = _dafny.SeqOf(_212_v85, _212_v85, Companion_D21_.Create_DC50_(_211_v84))
	_213_v86 = _213_v86
	var _214_v87 _dafny.Array
	_ = _214_v87
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_3
	var _nw22 _dafny.Array
	_ = _nw22
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw22 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) bool = (func(_215_p2 bool) func(_dafny.Int) bool {
			return func(_216_i7 _dafny.Int) bool {
				return _215_p2
			}
		})(p2)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw22 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw22).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw22).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_214_v87 = _nw22
	_214_v87 = _214_v87
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_214_v87), 0))
	_ = _index18
	(_214_v87).ArraySet1(p2, (_index18).Int())
	var _217_v88 _dafny.Map
	_ = _217_v88
	_217_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_209_v82).Cardinality())), p2)
	var _218_v89 _dafny.Array
	_ = _218_v89
	var _nwElement0_5 _dafny.Int = (_217_v88).Cardinality()
	_ = _nwElement0_5
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(7))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_5, 0)
	(_nw23).ArraySet1(p0, 1)
	(_nw23).ArraySet1(p0, 2)
	(_nw23).ArraySet1(p3, 3)
	(_nw23).ArraySet1(_dafny.IntOfUint32((_209_v82).Cardinality()), 4)
	(_nw23).ArraySet1(p3, 5)
	(_nw23).ArraySet1(p3, 6)
	_218_v89 = _nw23
	var _219_v90 _dafny.Map
	_ = _219_v90
	_219_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_218_v89, p1)
	var _220_v91 _dafny.Sequence
	_ = _220_v91
	_220_v91 = _dafny.SeqOf(p0, p0)
	var _221_v92 _dafny.Sequence
	_ = _221_v92
	_221_v92 = _dafny.SeqOf(_220_v91, _220_v91)
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_214_v87), 0))
	_ = _index19
	(_214_v87).ArraySet1(((func() _dafny.MultiSet {
		if (_219_v90).Contains(_218_v89) {
			return (_219_v90).Get(_218_v89).(_dafny.MultiSet)
		}
		return p1
	})()).IsDisjointFrom((p1).Update(p2, Companion_Default___.Abs(_dafny.IntOfUint32((_221_v92).Cardinality())))), (_index19).Int())
	var _222_v93 _dafny.Map
	_ = _222_v93
	_222_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("nrtq"), _218_v89)
	r2 = ((_222_v93).Merge(_222_v93)).Cardinality()
	r0 = p0
	r1 = (_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_214_v87), 0))).Int()).(bool)
	r2 = p0
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _223_v0 bool
	_ = _223_v0
	_223_v0 = true
	var _224_v1 _dafny.Array
	_ = _224_v1
	var _nwElement0_6 bool = _223_v0
	_ = _nwElement0_6
	var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
	_ = _nw24
	(_nw24).ArraySet1(_nwElement0_6, 0)
	_224_v1 = _nw24
	var _225_v2 _dafny.Int
	_ = _225_v2
	_225_v2 = _dafny.IntOfInt64(230)
	var _226_v3 _dafny.Sequence
	_ = _226_v3
	_226_v3 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v1, (_dafny.Zero).Minus(_225_v2)))
	var _227_v4 _dafny.Map
	_ = _227_v4
	_227_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(298), _223_v0)
	var _228_v5 _dafny.Map
	_ = _228_v5
	_228_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_225_v2, _225_v2)
	var _229_v6 _dafny.Map
	_ = _229_v6
	_229_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))
	var _230_v7 _dafny.Sequence
	_ = _230_v7
	_230_v7 = _dafny.SeqOf(_223_v0)
	var _231_v8 _dafny.Map
	_ = _231_v8
	_231_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v0, _223_v0)
	var _232_v9 _dafny.Map
	_ = _232_v9
	_232_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v7, true)
	var _233_v10 _dafny.Sequence
	_ = _233_v10
	_233_v10 = _dafny.UnicodeSeqOfUtf8Bytes("wjoqku")
	var _234_v11 _dafny.Array
	_ = _234_v11
	var _nwElement0_7 _dafny.Int = _225_v2
	_ = _nwElement0_7
	var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(26))
	_ = _nw25
	(_nw25).ArraySet1(_nwElement0_7, 0)
	(_nw25).ArraySet1((_227_v4).Cardinality(), 1)
	(_nw25).ArraySet1(_225_v2, 2)
	(_nw25).ArraySet1((func() _dafny.Int {
		if (_228_v5).Contains(_225_v2) {
			return (_228_v5).Get(_225_v2).(_dafny.Int)
		}
		return _225_v2
	})(), 3)
	(_nw25).ArraySet1(_225_v2, 4)
	(_nw25).ArraySet1((_229_v6).Cardinality(), 5)
	(_nw25).ArraySet1(_225_v2, 6)
	(_nw25).ArraySet1(_225_v2, 7)
	(_nw25).ArraySet1(_dafny.IntOfInt64(-846), 8)
	(_nw25).ArraySet1(_225_v2, 9)
	(_nw25).ArraySet1(_dafny.IntOfInt64(379), 10)
	(_nw25).ArraySet1(_225_v2, 11)
	(_nw25).ArraySet1(_225_v2, 12)
	(_nw25).ArraySet1(_225_v2, 13)
	(_nw25).ArraySet1(_225_v2, 14)
	(_nw25).ArraySet1((_dafny.Zero).Minus(_225_v2), 15)
	(_nw25).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_230_v7, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_230_v7).Cardinality()), _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32(), _223_v0)).Cardinality()), 16)
	(_nw25).ArraySet1(_225_v2, 17)
	(_nw25).ArraySet1((_dafny.Zero).Minus(_225_v2), 18)
	(_nw25).ArraySet1(_225_v2, 19)
	(_nw25).ArraySet1((_dafny.Zero).Minus((_231_v8).Cardinality()), 20)
	(_nw25).ArraySet1((_232_v9).Cardinality(), 21)
	(_nw25).ArraySet1(_dafny.IntOfUint32((_233_v10).Cardinality()), 22)
	(_nw25).ArraySet1(_225_v2, 23)
	(_nw25).ArraySet1((_227_v4).Cardinality(), 24)
	(_nw25).ArraySet1(_225_v2, 25)
	_234_v11 = _nw25
	var _235_v12 _dafny.Set
	_ = _235_v12
	_235_v12 = _dafny.SetOf(_224_v1)
	var _236_v13 _dafny.CodePoint
	_ = _236_v13
	_236_v13 = _dafny.CodePoint('a')
	var _237_v14 _dafny.Map
	_ = _237_v14
	_237_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_225_v2), _231_v8)
	var _238_globalState *GlobalState
	_ = _238_globalState
	var _nw26 *GlobalState = New_GlobalState_()
	_ = _nw26
	_nw26.Ctor__(_224_v1, _dafny.IntOfInt64(986), (_226_v3).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_226_v3).Cardinality()))).Uint32()).(_dafny.Map), true, _234_v11, false, _224_v1, _dafny.IntOfInt64(986), _233_v10, true, _235_v12, _dafny.IntOfInt64(737), _dafny.Companion_Sequence_.Update(_233_v10, (Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_233_v10).Cardinality()))).Uint32(), _236_v13), false, _dafny.CodePoint('a'), (_231_v8).Merge((func() _dafny.Map {
		if (_237_v14).Contains(_225_v2) {
			return (_237_v14).Get(_225_v2).(_dafny.Map)
		}
		return _231_v8
	})()), _dafny.IntOfInt64(108), _dafny.Companion_Sequence_.Concatenate(_233_v10, _233_v10), _dafny.UnicodeSeqOfUtf8Bytes("n"), false, _dafny.IntOfInt64(-137), _234_v11, _234_v11)
	_238_globalState = _nw26
	var _hi0 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
		if true {
			return (_dafny.Zero).Minus(_225_v2)
		}
		return Companion_Default___.Fm0(_223_v0, _225_v2, _238_globalState)
	})())
	_ = _hi0
	for _239_i0 := _dafny.IntOfInt64(425); _239_i0.Cmp(_hi0) < 0; _239_i0 = _239_i0.Plus(_dafny.One) {
		if true {
			(_238_globalState).F7 = _225_v2
			(_238_globalState).F7 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sulow")).Cardinality())
			_233_v10 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(723))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}(func(_240_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))
			(_238_globalState).F7 = _225_v2
			var _241_v15 _dafny.MultiSet
			_ = _241_v15
			_241_v15 = _dafny.MultiSetOf(_223_v0)
			var _242_v16 _dafny.Int
			_ = _242_v16
			var _243_v17 bool
			_ = _243_v17
			var _244_v18 _dafny.Int
			_ = _244_v18
			var _out0 _dafny.Int
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			_out0, _out1, _out2 = Companion_Default___.M0((_dafny.Zero).Minus(_239_i0), (_241_v15).Intersection(_241_v15), _223_v0, _dafny.IntOfUint32((_230_v7).Cardinality()), _238_globalState)
			_242_v16 = _out0
			_243_v17 = _out1
			_244_v18 = _out2
		} else {
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_224_v1), 0))
			_ = _index20
			(_224_v1).ArraySet1(_223_v0, (_index20).Int())
			var _245_v19 _dafny.Sequence
			_ = _245_v19
			_245_v19 = _dafny.SeqOf(_dafny.IntOfInt64(313))
			var _246_v20 _dafny.Sequence
			_ = _246_v20
			_246_v20 = _dafny.SeqOf(_245_v19, _245_v19, _245_v19)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_224_v1), 0))
			_ = _index21
			(_224_v1).ArraySet1(_dafny.Companion_Sequence_.Equal((_246_v20).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_225_v2), _dafny.IntOfUint32((_246_v20).Cardinality()))).Uint32()).(_dafny.Sequence), _245_v19), (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_224_v1), 0))
			_ = _index22
			(_224_v1).ArraySet1(((_dafny.IntOfUint32((Companion_Default___.Fm1(Companion_Default___.Fm2(_238_globalState), _238_globalState)).Cardinality())).Plus(_225_v2)).Cmp(_dafny.IntOfInt64(591)) <= 0, (_index22).Int())
			var _247_v21 _dafny.MultiSet
			_ = _247_v21
			_247_v21 = _dafny.MultiSetOf(_223_v0, !(_223_v0), false, _223_v0, _223_v0)
			var _248_v22 _dafny.Int
			_ = _248_v22
			var _249_v23 bool
			_ = _249_v23
			var _250_v24 _dafny.Int
			_ = _250_v24
			var _out3 _dafny.Int
			_ = _out3
			var _out4 bool
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out3, _out4, _out5 = Companion_Default___.M0(_225_v2, _247_v21, (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool), _239_i0, _238_globalState)
			_248_v22 = _out3
			_249_v23 = _out4
			_250_v24 = _out5
			var _251_v25 D0
			_ = _251_v25
			_251_v25 = Companion_D0_.Create_DC0_(Companion_Default___.Fm3(_223_v0, _238_globalState))
			var _252_v26 _dafny.Set
			_ = _252_v26
			_252_v26 = _dafny.SetOf(_250_v24)
			var _253_v27 _dafny.Int
			_ = _253_v27
			var _254_v28 bool
			_ = _254_v28
			var _255_v29 _dafny.Int
			_ = _255_v29
			var _out6 _dafny.Int
			_ = _out6
			var _out7 bool
			_ = _out7
			var _out8 _dafny.Int
			_ = _out8
			_out6, _out7, _out8 = Companion_Default___.M0(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lrbqh")).Cardinality()), ((_251_v25).Dtor_cf0()).Difference(_247_v21), _223_v0, (_250_v24).Times((_252_v26).Cardinality()), _238_globalState)
			_253_v27 = _out6
			_254_v28 = _out7
			_255_v29 = _out8
			(_238_globalState).F7 = _225_v2
		}
		var _256_v30 _dafny.Array
		_ = _256_v30
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_4
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Sequence = func(_257_i2 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("feg")
			}
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw27 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw27).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw27).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_256_v30 = _nw27
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_256_v30), 0))
		_ = _index23
		(_256_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_233_v10, _233_v10), (_index23).Int())
		var _258_v31 _dafny.MultiSet
		_ = _258_v31
		_258_v31 = _dafny.MultiSetOf(_223_v0, _223_v0, _223_v0)
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_256_v30), 0))
		_ = _index24
		(_256_v30).ArraySet1(Companion_Default___.Fm1((_225_v2).Cmp((func() _dafny.Int {
			if (_258_v31).Contains(!(_223_v0)) {
				return (_258_v31).Multiplicity(!(_223_v0))
			}
			return _239_i0
		})()) <= 0, _238_globalState), (_index24).Int())
		var _259_v32 _dafny.Array
		_ = _259_v32
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(5))
		_ = _nw28
		_259_v32 = _nw28
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_259_v32), 0))
		_ = _index25
		(_259_v32).ArraySet1(Companion_D0_.Create_DC2_(_223_v0, _239_i0, _223_v0), (_index25).Int())
		var _260_v34 _dafny.Set
		_ = _260_v34
		_260_v34 = _dafny.SetOf(_225_v2)
		var _261_v35 D0
		_ = _261_v35
		_261_v35 = Companion_D0_.Create_DC2_(!(_223_v0), (func() _dafny.Map {
			var _coll46 = _dafny.NewMapBuilder()
			_ = _coll46
			for _iter46 := _dafny.Iterate((_260_v34).Elements()); ; {
				_compr_46, _ok46 := _iter46()
				if !_ok46 {
					break
				}
				var _262_v33 _dafny.Int
				_262_v33 = interface{}(_compr_46).(_dafny.Int)
				if (_260_v34).Contains(_262_v33) {
					_coll46.Add(Companion_Default___.SafeDivisionInt(_262_v33, _239_i0), _236_v13)
				}
			}
			return _coll46.ToMap()
		}()).Cardinality(), _223_v0)
		var _pat_let_tv5 = _223_v0
		_ = _pat_let_tv5
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_259_v32), 0))
		_ = _index26
		(_259_v32).ArraySet1(func(_pat_let8_0 D0) D0 {
			return func(_263_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let9_0 bool) D0 {
					return func(_264_dt__update_hcf2_h0 bool) D0 {
						return Companion_D0_.Create_DC2_(_264_dt__update_hcf2_h0, (_263_dt__update__tmp_h0).Dtor_cf3(), (_263_dt__update__tmp_h0).Dtor_cf4())
					}(_pat_let9_0)
				}(!(_pat_let_tv5))
			}(_pat_let8_0)
		}(_261_v35), (_index26).Int())
		var _rhs15 _dafny.Int = _239_i0
		_ = _rhs15
		var _rhs16 _dafny.Int = Companion_Default___.Fm0(false, _239_i0, _238_globalState)
		_ = _rhs16
		var _rhs17 _dafny.MultiSet = ((_dafny.MultiSetOf(_223_v0, _223_v0)).Union(_dafny.MultiSetOf(_223_v0, _223_v0))).Intersection(_dafny.MultiSetOf(_223_v0, !(_223_v0), true))
		_ = _rhs17
		var _rhs18 bool = _223_v0
		_ = _rhs18
		var _lhs11 *GlobalState = _238_globalState
		_ = _lhs11
		var _lhs12 *GlobalState = _238_globalState
		_ = _lhs12
		_lhs11.F7 = _rhs15
		_225_v2 = _rhs16
		_258_v31 = _rhs17
		_lhs12.F13 = _rhs18
	}
	(_238_globalState).F13 = !((_225_v2).Cmp(_225_v2) >= 0) || (!(_dafny.Companion_Sequence_.IsPrefixOf(_233_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(215))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg46 _dafny.Int) interface{} {
			return coer46(arg46)
		}
	}((func(_265_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_266_i3 _dafny.Int) _dafny.CodePoint {
			return _265_v13
		}
	})(_236_v13))))))
	var _267_v36 _dafny.Set
	_ = _267_v36
	_267_v36 = _dafny.SetOf(_223_v0)
	var _268_v37 _dafny.Set
	_ = _268_v37
	_268_v37 = _dafny.SetOf(Companion_Default___.Fm0(_223_v0, _225_v2, _238_globalState), _225_v2)
	var _269_v38 _dafny.Set
	_ = _269_v38
	_269_v38 = _dafny.SetOf(_dafny.IntOfInt64(462), (_268_v37).Cardinality(), _225_v2, _225_v2, _225_v2)
	_225_v2 = (_225_v2).Minus(Companion_Default___.SafeModuloInt(((_228_v5).Update((_267_v36).Cardinality(), (_268_v37).Cardinality())).Cardinality(), Companion_Default___.Fm0(_223_v0, _dafny.IntOfUint32((_dafny.SeqOf(_223_v0)).Cardinality()), _238_globalState)))
	var _270_v39 _dafny.MultiSet
	_ = _270_v39
	_270_v39 = _dafny.MultiSetOf((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_225_v2, (_dafny.Zero).Minus(_225_v2))))
	_270_v39 = (_270_v39).Intersection(_270_v39)
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
	_ = _index27
	(_234_v11).ArraySet1(_225_v2, (_index27).Int())
	var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
	_ = _index28
	(_234_v11).ArraySet1(_225_v2, (_index28).Int())
	var _271_v40 _dafny.Array
	_ = _271_v40
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(22))
	_ = _nw29
	_271_v40 = _nw29
	var _272_v41 _dafny.MultiSet
	_ = _272_v41
	_272_v41 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("lpbndwsy"), Companion_Default___.Fm1(_223_v0, _238_globalState), _233_v10, _233_v10)
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_271_v40), 0))
	_ = _index29
	(_271_v40).ArraySet1(_272_v41, (_index29).Int())
	var _273_v42 D0
	_ = _273_v42
	_273_v42 = Companion_D0_.Create_DC1_(_225_v2)
	var _pat_let_tv6 = _272_v41
	_ = _pat_let_tv6
	var _pat_let_tv7 = _233_v10
	_ = _pat_let_tv7
	var _pat_let_tv8 = _230_v7
	_ = _pat_let_tv8
	var _pat_let_tv9 = _272_v41
	_ = _pat_let_tv9
	var _pat_let_tv10 = _272_v41
	_ = _pat_let_tv10
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_271_v40), 0))
	_ = _index30
	var _rhs19 _dafny.MultiSet = func(_source5 D0) _dafny.MultiSet {
		if _source5.Is_DC1() {
			var _274___mcc_h0 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
			_ = _274___mcc_h0
			var _275_cf1 _dafny.Int = _274___mcc_h0
			_ = _275_cf1
			return (_pat_let_tv6).Update(_pat_let_tv7, Companion_Default___.Abs(_dafny.IntOfUint32((_pat_let_tv8).Cardinality())))
		} else if _source5.Is_DC2() {
			var _276___mcc_h1 bool = _source5.Get_().(D0_DC2).Cf2
			_ = _276___mcc_h1
			var _277___mcc_h2 _dafny.Int = _source5.Get_().(D0_DC2).Cf3
			_ = _277___mcc_h2
			var _278___mcc_h3 bool = _source5.Get_().(D0_DC2).Cf4
			_ = _278___mcc_h3
			var _279_cf4 bool = _278___mcc_h3
			_ = _279_cf4
			var _280_cf3 _dafny.Int = _277___mcc_h2
			_ = _280_cf3
			var _281_cf2 bool = _276___mcc_h1
			_ = _281_cf2
			return _pat_let_tv9
		} else if _source5.Is_DC0() {
			var _282___mcc_h4 _dafny.MultiSet = _source5.Get_().(D0_DC0).Cf0
			_ = _282___mcc_h4
			var _283_cf0 _dafny.MultiSet = _282___mcc_h4
			_ = _283_cf0
			return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("nt"))
		} else {
			var _284___mcc_h5 D0 = _source5.Get_().(D0_DC3).Cf5
			_ = _284___mcc_h5
			var _285_cf5 D0 = _284___mcc_h5
			_ = _285_cf5
			return _pat_let_tv10
		}
	}(_273_v42)
	_ = _rhs19
	var _rhs20 bool = ((_dafny.Zero).Minus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))).Cmp(Companion_Default___.SafeModuloInt((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _225_v2)) == 0
	_ = _rhs20
	var _lhs13 _dafny.Array = _271_v40
	_ = _lhs13
	var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_271_v40), 0))
	_ = _lhs14
	var _lhs15 *GlobalState = _238_globalState
	_ = _lhs15
	(_lhs13).ArraySet1(_rhs19, (_lhs14).Int())
	_lhs15.F13 = _rhs20
	var _286_v43 _dafny.Map
	_ = _286_v43
	_286_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v1, (_233_v10).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_233_v10).Cardinality()))).Uint32()).(_dafny.CodePoint))
	var _287_v44 _dafny.Array
	_ = _287_v44
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_5
	var _nw30 _dafny.Array
	_ = _nw30
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw30 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_288_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_289_i4 _dafny.Int) _dafny.CodePoint {
				return _288_v13
			}
		})(_236_v13)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw30 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw30).ArraySet1CodePoint(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw30).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_287_v44 = _nw30
	var _290_v45 *C3
	_ = _290_v45
	var _nw31 *C3 = New_C3_()
	_ = _nw31
	_nw31.Ctor__(_286_v43, _287_v44, Companion_Default___.Fm29((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _223_v0, _223_v0, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _238_globalState), ((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).Cmp((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)) != 0)
	_290_v45 = _nw31
	var _291_v46 D5
	_ = _291_v46
	_291_v46 = Companion_D5_.Create_DC14_(_223_v0)
	var _292_v47 _dafny.Set
	_ = _292_v47
	_292_v47 = _dafny.SetOf(_291_v46, _291_v46, Companion_D5_.Create_DC14_(_223_v0), func(_pat_let10_0 D5) D5 {
		return func(_293_dt__update__tmp_h1 D5) D5 {
			return func(_pat_let11_0 bool) D5 {
				return func(_294_dt__update_hcf23_h0 bool) D5 {
					return Companion_D5_.Create_DC14_(_294_dt__update_hcf23_h0)
				}(_pat_let11_0)
			}(true)
		}(_pat_let10_0)
	}(Companion_D5_.Create_DC14_(_223_v0)), _291_v46)
	var _295_v49 _dafny.Map
	_ = _295_v49
	_295_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_292_v47, func() _dafny.Map {
		var _coll47 = _dafny.NewMapBuilder()
		_ = _coll47
		for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(568), _dafny.IntOfInt64(81))); ; {
			_compr_47, _ok47 := _iter47()
			if !_ok47 {
				break
			}
			var _296_v48 _dafny.Int
			_296_v48 = interface{}(_compr_47).(_dafny.Int)
			if ((_dafny.IntOfInt64(568)).Cmp(_296_v48) <= 0) && ((_296_v48).Cmp(_dafny.IntOfInt64(81)) < 0) {
				_coll47.Add(Companion_Default___.SafeModuloInt(_296_v48, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vlhwkla")).Cardinality())), (func() _dafny.Int {
					if (_270_v39).Contains(_dafny.IntOfInt64(395)) {
						return (_270_v39).Multiplicity(_dafny.IntOfInt64(395))
					}
					return (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)
				})())
			}
		}
		return _coll47.ToMap()
	}())
	var _297_v52 _dafny.Sequence
	_ = _297_v52
	_297_v52 = _dafny.SeqOf(_295_v49, func() _dafny.Map {
		var _coll48 = _dafny.NewMapBuilder()
		_ = _coll48
		for _iter48 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(_291_v46), _292_v47, func() _dafny.Set {
			var _coll49 = _dafny.NewBuilder()
			_ = _coll49
			for _iter49 := _dafny.Iterate((Companion_Default___.Fm58(_236_v13, _dafny.IntOfUint32((_233_v10).Cardinality()), _223_v0, _238_globalState)).Keys().Elements()); ; {
				_compr_49, _ok49 := _iter49()
				if !_ok49 {
					break
				}
				var _298_v51 D5
				_298_v51 = interface{}(_compr_49).(D5)
				if (Companion_Default___.Fm58(_236_v13, _dafny.IntOfUint32((_233_v10).Cardinality()), _223_v0, _238_globalState)).Contains(_298_v51) {
					_coll49.Add(_298_v51)
				}
			}
			return _coll49.ToSet()
		}())).Elements()); ; {
			_compr_48, _ok48 := _iter48()
			if !_ok48 {
				break
			}
			var _299_v50 _dafny.Set
			_299_v50 = interface{}(_compr_48).(_dafny.Set)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(_291_v46), _292_v47, func() _dafny.Set {
				var _coll50 = _dafny.NewBuilder()
				_ = _coll50
				for _iter50 := _dafny.Iterate((Companion_Default___.Fm58(_236_v13, _dafny.IntOfUint32((_233_v10).Cardinality()), _223_v0, _238_globalState)).Keys().Elements()); ; {
					_compr_50, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _300_v51 D5
					_300_v51 = interface{}(_compr_50).(D5)
					if (Companion_Default___.Fm58(_236_v13, _dafny.IntOfUint32((_233_v10).Cardinality()), _223_v0, _238_globalState)).Contains(_300_v51) {
						_coll50.Add(_300_v51)
					}
				}
				return _coll50.ToSet()
			}()), _299_v50) {
				_coll48.Add(_299_v50, _228_v5)
			}
		}
		return _coll48.ToMap()
	}(), (_295_v49).Merge(_295_v49), Companion_Default___.Fm59(_233_v10, _238_globalState), _295_v49)
	_295_v49 = (_297_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-874), _dafny.IntOfUint32((_297_v52).Cardinality()))).Uint32()).(_dafny.Map)
	var _301_v53 _dafny.Array
	_ = _301_v53
	var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
	_ = _nw32
	_301_v53 = _nw32
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_301_v53), 0))
	_ = _index31
	(_301_v53).ArraySet1(_dafny.SeqOf(_224_v1), (_index31).Int())
	var _302_v54 _dafny.Sequence
	_ = _302_v54
	_302_v54 = _dafny.SeqOf(_224_v1, _224_v1, _224_v1, _224_v1, _224_v1)
	var _303_v55 _dafny.Sequence
	_ = _303_v55
	_303_v55 = _dafny.SeqOf(_302_v54, _dafny.SeqOf(_224_v1))
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_301_v53), 0))
	_ = _index32
	(_301_v53).ArraySet1((_303_v55).Select((Companion_Default___.SafeIndex((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_303_v55).Cardinality()))).Uint32()).(_dafny.Sequence), (_index32).Int())
	var _hi1 _dafny.Int = _225_v2
	_ = _hi1
	for _304_i5 := Companion_Default___.SafeDivisionInt(_225_v2, _225_v2); _304_i5.Cmp(_hi1) < 0; _304_i5 = _304_i5.Plus(_dafny.One) {
		var _305_v56 _dafny.Array
		_ = _305_v56
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_6
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Map = (func(_306_v0 bool, _307_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_308_i6 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v0, _307_v2)
				}
			})(_223_v0, _225_v2)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw33 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw33).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw33).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_305_v56 = _nw33
		var _309_v57 *C1
		_ = _309_v57
		var _nw34 *C1 = New_C1_()
		_ = _nw34
		_nw34.Ctor__(true, false)
		_309_v57 = _nw34
		var _310_v58 _dafny.Map
		_ = _310_v58
		_310_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v56, _309_v57)
		var _311_v59 _dafny.Sequence
		_ = _311_v59
		_311_v59 = _dafny.SeqOf((func() *C1 {
			if (_310_v58).Contains(_305_v56) {
				return (_310_v58).Get(_305_v56).(*C1)
			}
			return _309_v57
		})(), (func() *C1 {
			if (_309_v57).F39() {
				return _309_v57
			}
			return _309_v57
		})())
		var _rhs21 _dafny.Array = _234_v11
		_ = _rhs21
		var _rhs22 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_311_v59, _311_v59)
		_ = _rhs22
		var _rhs23 _dafny.Array = _224_v1
		_ = _rhs23
		var _lhs16 *GlobalState = _238_globalState
		_ = _lhs16
		_lhs16.F22 = _rhs21
		_311_v59 = _rhs22
		_224_v1 = _rhs23
		(_238_globalState).F13 = (_230_v7).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0((_309_v57).F39(), (_228_v5).Cardinality(), _238_globalState), _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32()).(bool)
		if (_dafny.IntOfInt64(62)).Cmp((_304_i5).Minus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))) == 0 {
			var _312_v60 *C10
			_ = _312_v60
			var _nw35 *C10 = New_C10_()
			_ = _nw35
			_nw35.Ctor__((true) == ((_309_v57).F39()), (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.CodePoint('j'), (_225_v2).Cmp(_225_v2) <= 0, _287_v44, _223_v0)
			_312_v60 = _nw35
			var _313_v61 _dafny.Sequence
			_ = _313_v61
			_313_v61 = _dafny.SeqOf((func() _dafny.Int {
				if (_309_v57).F39() {
					return _312_v60.F29
				}
				return _312_v60.F29
			})())
			var _rhs24 *C10 = _312_v60
			_ = _rhs24
			var _rhs25 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _223_v0 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(238))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}((func(_314_v60 *C10) func(_dafny.Int) _dafny.Int {
						return func(_315_i7 _dafny.Int) _dafny.Int {
							return _314_v60.F29
						}
					})(_312_v60)))
				}
				return _313_v61
			})(), _313_v61), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_233_v10).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _223_v0 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(238))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}((func(_316_v60 *C10) func(_dafny.Int) _dafny.Int {
						return func(_317_i7 _dafny.Int) _dafny.Int {
							return _316_v60.F29
						}
					})(_312_v60)))
				}
				return _313_v61
			})(), _313_v61)).Cardinality()))).Uint32(), _304_i5)
			_ = _rhs25
			var _rhs26 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_313_v61, _313_v61)
			_ = _rhs26
			var _rhs27 _dafny.Int = _225_v2
			_ = _rhs27
			var _rhs28 _dafny.Int = _304_i5
			_ = _rhs28
			var _lhs17 *GlobalState = _238_globalState
			_ = _lhs17
			var _lhs18 *GlobalState = _238_globalState
			_ = _lhs18
			_312_v60 = _rhs24
			_313_v61 = _rhs25
			_313_v61 = _rhs26
			_lhs17.F7 = _rhs27
			_lhs18.F7 = _rhs28
			var _318_v62 _dafny.MultiSet
			_ = _318_v62
			_318_v62 = _dafny.MultiSetOf(_223_v0, (_309_v57).F39(), (_309_v57).F39())
			var _319_v63 *C2
			_ = _319_v63
			var _nw36 *C2 = New_C2_()
			_ = _nw36
			_nw36.Ctor__(_318_v62, _287_v44, _236_v13, (_312_v60).F28())
			_319_v63 = _nw36
			(_238_globalState).F13 = _223_v0
			var _320_v64 *C0
			_ = _320_v64
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_312_v60).F28(), _dafny.UnicodeSeqOfUtf8Bytes("w")), ((_309_v57).F39()) && (_223_v0), _dafny.CodePoint('l'), (_312_v60).F28())
			_320_v64 = _nw37
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _index33
			var _rhs29 *C0 = _320_v64
			_ = _rhs29
			var _rhs30 _dafny.MultiSet = _318_v62
			_ = _rhs30
			var _rhs31 _dafny.Int = Companion_Default___.SafeModuloInt(_225_v2, _312_v60.F29)
			_ = _rhs31
			var _rhs32 *C3 = _290_v45
			_ = _rhs32
			var _rhs33 _dafny.Int = (func() _dafny.Int {
				if _223_v0 {
					return (Companion_Default___.Fm3(!((_309_v57).F39()), _238_globalState)).Cardinality()
				}
				return (_312_v60).Fm9(_313_v61, _312_v60.F29, (_309_v57).F39(), _238_globalState)
			})()
			_ = _rhs33
			var _lhs19 *GlobalState = _238_globalState
			_ = _lhs19
			var _lhs20 _dafny.Array = _234_v11
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _lhs21
			_320_v64 = _rhs29
			_318_v62 = _rhs30
			_lhs19.F7 = _rhs31
			_290_v45 = _rhs32
			(_lhs20).ArraySet1(_rhs33, (_lhs21).Int())
			_223_v0 = (_320_v64).F38()
		} else {
			var _321_v65 _dafny.Map
			_ = _321_v65
			_321_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_309_v57).F39(), Companion_Default___.Fm1((_309_v57).F39(), _238_globalState))
			var _322_v66 T2
			_ = _322_v66
			var _nw38 *C2 = New_C2_()
			_ = _nw38
			_nw38.Ctor__(_dafny.MultiSetFromSeq(_230_v7), _287_v44, _236_v13, (_309_v57).F39())
			_322_v66 = _nw38
			var _323_v67 _dafny.Map
			_ = _323_v67
			_323_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_321_v65).Contains((_309_v57).F39()) {
					return (_321_v65).Get((_309_v57).F39()).(_dafny.Sequence)
				}
				return _233_v10
			})(), (func() T2 {
				if _223_v0 {
					return _322_v66
				}
				return _322_v66
			})())
			_323_v67 = (_323_v67).Update(_233_v10, _322_v66)
			var _324_v68 _dafny.Sequence
			_ = _324_v68
			_324_v68 = _dafny.SeqOf(_304_i5, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))
			var _325_v69 _dafny.Map
			_ = _325_v69
			_325_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_322_v66).F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_228_v5).Cardinality(), (_322_v66).F24()))
			var _326_v71 _dafny.Sequence
			_ = _326_v71
			_326_v71 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_327_i5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_328_i8 _dafny.Int) _dafny.Int {
					return _327_i5
				}
			})(_304_i5))), (Companion_Default___.SafeIndex(((func() _dafny.Map {
				if (_325_v69).Contains((_309_v57).F39()) {
					return (_325_v69).Get((_309_v57).F39()).(_dafny.Map)
				}
				return func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-632), _dafny.IntOfInt64(635))); ; {
						_compr_51, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _329_v70 _dafny.Int
						_329_v70 = interface{}(_compr_51).(_dafny.Int)
						if ((_dafny.IntOfInt64(-632)).Cmp(_329_v70) <= 0) && ((_329_v70).Cmp(_dafny.IntOfInt64(635)) < 0) {
							_coll51.Add((_329_v70).Plus((_227_v4).Cardinality()), (Companion_D6_.Create_DC16_(_231_v8, (_309_v57).F39(), false)).Dtor_cf27())
						}
					}
					return _coll51.ToMap()
				}()
			})()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_330_i5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_331_i8 _dafny.Int) _dafny.Int {
					return _330_i5
				}
			})(_304_i5)))).Cardinality()))).Uint32(), _dafny.IntOfUint32((_233_v10).Cardinality())), _324_v68), _324_v68, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(930))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_332_v11 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_333_i9 _dafny.Int) _dafny.Int {
					return (_332_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_332_v11), 0))).Int()).(_dafny.Int)
				}
			})(_234_v11))), _324_v68, _324_v68)
			_324_v68 = (_326_v71).Select((Companion_Default___.SafeIndex(((_309_v57).Fm33(_238_globalState)).Minus(_304_i5), _dafny.IntOfUint32((_326_v71).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _334_v72 bool
			_ = _334_v72
			var _335_v73 _dafny.CodePoint
			_ = _335_v73
			var _336_v74 bool
			_ = _336_v74
			var _out9 bool
			_ = _out9
			var _out10 _dafny.CodePoint
			_ = _out10
			var _out11 bool
			_ = _out11
			_out9, _out10, _out11 = (_290_v45).M11(_223_v0, _238_globalState)
			_334_v72 = _out9
			_335_v73 = _out10
			_336_v74 = _out11
			var _337_v75 T1
			_ = _337_v75
			var _nw39 *C9 = New_C9_()
			_ = _nw39
			_nw39.Ctor__(_236_v13, _336_v74)
			_337_v75 = _nw39
			_337_v75 = _337_v75
			_326_v71 = (func() _dafny.Sequence {
				if _334_v72 {
					return Companion_Default___.Fm60((_309_v57).F39(), _238_globalState)
				}
				return _326_v71
			})()
		}
		_236_v13 = _236_v13
	}
	var _338_v76 D21
	_ = _338_v76
	_338_v76 = Companion_D21_.Create_DC51_((_223_v0) || (_223_v0))
	var _source6 D21 = _338_v76
	_ = _source6
	if _source6.Is_DC51() {
		var _339___mcc_h6 bool = _source6.Get_().(D21_DC51).Cf81
		_ = _339___mcc_h6
		var _340_cf81 bool = _339___mcc_h6
		_ = _340_cf81
		(_238_globalState).F13 = !(_223_v0)
		if ((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(-409)) >= 0 {
			var _341_v77 *C1
			_ = _341_v77
			var _nw40 *C1 = New_C1_()
			_ = _nw40
			_nw40.Ctor__(_223_v0, _340_cf81)
			_341_v77 = _nw40
			var _342_v78 _dafny.MultiSet
			_ = _342_v78
			_342_v78 = _dafny.MultiSetOf((_341_v77).Fm32(_340_cf81, _225_v2, _238_globalState), (_230_v7).Select((Companion_Default___.SafeIndex((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32()).(bool))
			var _343_v79 _dafny.Int
			_ = _343_v79
			var _344_v80 bool
			_ = _344_v80
			var _345_v81 _dafny.Int
			_ = _345_v81
			var _out12 _dafny.Int
			_ = _out12
			var _out13 bool
			_ = _out13
			var _out14 _dafny.Int
			_ = _out14
			_out12, _out13, _out14 = Companion_Default___.M0(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v77, _dafny.CodePoint('p'))).Cardinality()).Plus((_270_v39).Cardinality()), _342_v78, _223_v0, _225_v2, _238_globalState)
			_343_v79 = _out12
			_344_v80 = _out13
			_345_v81 = _out14
			var _346_v82 _dafny.Array
			_ = _346_v82
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw41
			_346_v82 = _nw41
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_346_v82), 0))
			_ = _index34
			(_346_v82).ArraySet1(_234_v11, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_346_v82), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _index36
			var _rhs34 _dafny.Map = (_227_v4).Merge(_227_v4)
			_ = _rhs34
			var _rhs35 bool = _344_v80
			_ = _rhs35
			var _rhs36 _dafny.Array = _234_v11
			_ = _rhs36
			var _rhs37 _dafny.Int = Companion_Default___.SafeDivisionInt(_225_v2, _225_v2)
			_ = _rhs37
			var _rhs38 bool = _344_v80
			_ = _rhs38
			var _lhs22 _dafny.Array = _346_v82
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_346_v82), 0))
			_ = _lhs23
			var _lhs24 _dafny.Array = _234_v11
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _lhs25
			_227_v4 = _rhs34
			_344_v80 = _rhs35
			(_lhs22).ArraySet1(_rhs36, (_lhs23).Int())
			(_lhs24).ArraySet1(_rhs37, (_lhs25).Int())
			_223_v0 = _rhs38
			var _347_v83 *C2
			_ = _347_v83
			var _nw42 *C2 = New_C2_()
			_ = _nw42
			_nw42.Ctor__(_342_v78, _287_v44, (func() _dafny.CodePoint {
				if (_341_v77).F39() {
					return _236_v13
				}
				return _236_v13
			})(), _344_v80)
			_347_v83 = _nw42
			_340_cf81 = _344_v80
			var _348_v84 *C4
			_ = _348_v84
			var _nw43 *C4 = New_C4_()
			_ = _nw43
			_nw43.Ctor__((_347_v83).Fm24((_347_v83).Fm24(_345_v81, _340_cf81, _238_globalState), (_341_v77).F39(), _238_globalState), _228_v5, _287_v44, _236_v13, ((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).Cmp(_343_v79) >= 0)
			_348_v84 = _nw43
			var _nw44 *C4 = New_C4_()
			_ = _nw44
			_nw44.Ctor__(_343_v79, _348_v84.F34, _287_v44, _236_v13, _223_v0)
			_348_v84 = _nw44
		} else {
			var _349_v85 bool
			_ = _349_v85
			var _350_v86 _dafny.CodePoint
			_ = _350_v86
			var _351_v87 bool
			_ = _351_v87
			var _out15 bool
			_ = _out15
			var _out16 _dafny.CodePoint
			_ = _out16
			var _out17 bool
			_ = _out17
			_out15, _out16, _out17 = (_290_v45).M11(false, _238_globalState)
			_349_v85 = _out15
			_350_v86 = _out16
			_351_v87 = _out17
			var _352_v88 *C5
			_ = _352_v88
			var _nw45 *C5 = New_C5_()
			_ = _nw45
			_nw45.Ctor__(_224_v1)
			_352_v88 = _nw45
			var _353_v89 _dafny.Array
			_ = _353_v89
			var _nwElement0_8 *C5 = _352_v88
			_ = _nwElement0_8
			var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(12))
			_ = _nw46
			(_nw46).ArraySet1(_nwElement0_8, 0)
			(_nw46).ArraySet1(_352_v88, 1)
			(_nw46).ArraySet1(_352_v88, 2)
			(_nw46).ArraySet1(_352_v88, 3)
			(_nw46).ArraySet1(_352_v88, 4)
			(_nw46).ArraySet1(_352_v88, 5)
			(_nw46).ArraySet1(_352_v88, 6)
			(_nw46).ArraySet1(_352_v88, 7)
			(_nw46).ArraySet1(_352_v88, 8)
			(_nw46).ArraySet1(_352_v88, 9)
			(_nw46).ArraySet1(_352_v88, 10)
			(_nw46).ArraySet1(_352_v88, 11)
			_353_v89 = _nw46
			var _354_v90 D16
			_ = _354_v90
			_354_v90 = Companion_D16_.Create_DC37_(_352_v88)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_353_v89), 0))
			_ = _index37
			(_353_v89).ArraySet1((_354_v90).Dtor_cf61(), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_353_v89), 0))
			_ = _index38
			var _nw47 *C5 = New_C5_()
			_ = _nw47
			_nw47.Ctor__(_224_v1)
			(_353_v89).ArraySet1(_nw47, (_index38).Int())
			_233_v10 = _233_v10
			var _355_v91 _dafny.MultiSet
			_ = _355_v91
			_355_v91 = _dafny.MultiSetOf((_dafny.IntOfInt64(-238)).Cmp(_225_v2) == 0, _340_cf81)
			var _356_v94 *C4
			_ = _356_v94
			var _nw48 *C4 = New_C4_()
			_ = _nw48
			_nw48.Ctor__(_225_v2, _228_v5, _287_v44, _350_v86, _340_cf81)
			_356_v94 = _nw48
			var _357_v95 D24
			_ = _357_v95
			_357_v95 = Companion_D24_.Create_DC56_(_356_v94)
			var _358_v96 _dafny.Map
			_ = _358_v96
			_358_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll52 = _dafny.NewBuilder()
				_ = _coll52
				for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-204), _dafny.IntOfInt64(551))); ; {
					_compr_52, _ok52 := _iter52()
					if !_ok52 {
						break
					}
					var _359_v93 _dafny.Int
					_359_v93 = interface{}(_compr_52).(_dafny.Int)
					if ((_dafny.IntOfInt64(-204)).Cmp(_359_v93) <= 0) && ((_359_v93).Cmp(_dafny.IntOfInt64(551)) < 0) {
						_coll52.Add((_359_v93).Times((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll52.ToSet()
			}(), (_357_v95).Dtor_cf90())
			var _rhs39 _dafny.Int = Companion_Default___.SafeModuloInt(_225_v2, ((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).Minus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)))
			_ = _rhs39
			var _rhs40 _dafny.MultiSet = Companion_Default___.Fm3(_351_v87, _238_globalState)
			_ = _rhs40
			var _rhs41 bool = (_358_v96).Contains(func() _dafny.Set {
				var _coll53 = _dafny.NewBuilder()
				_ = _coll53
				for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(609), _dafny.IntOfInt64(-929))); ; {
					_compr_53, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _360_v92 _dafny.Int
					_360_v92 = interface{}(_compr_53).(_dafny.Int)
					if ((_dafny.IntOfInt64(609)).Cmp(_360_v92) <= 0) && ((_360_v92).Cmp(_dafny.IntOfInt64(-929)) < 0) {
						_coll53.Add((_360_v92).Minus((_dafny.Zero).Minus(_225_v2)))
					}
				}
				return _coll53.ToSet()
			}())
			_ = _rhs41
			var _lhs26 *GlobalState = _238_globalState
			_ = _lhs26
			_lhs26.F7 = _rhs39
			_355_v91 = _rhs40
			_340_cf81 = _rhs41
			var _361_v97 _dafny.Map
			_ = _361_v97
			_361_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("b"), _234_v11)
			var _362_v98 _dafny.Array
			_ = _362_v98
			var _nwElement0_9 _dafny.Array = _234_v11
			_ = _nwElement0_9
			var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(4))
			_ = _nw49
			(_nw49).ArraySet1(_nwElement0_9, 0)
			(_nw49).ArraySet1((func() _dafny.Array {
				if _351_v87 {
					return _234_v11
				}
				return _234_v11
			})(), 1)
			(_nw49).ArraySet1((func() _dafny.Array {
				if (_361_v97).Contains(_dafny.Companion_Sequence_.Update(_233_v10, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_351_v87, (_356_v94).F33(), _238_globalState), _dafny.IntOfUint32((_233_v10).Cardinality()))).Uint32(), _236_v13)) {
					return (_361_v97).Get(_dafny.Companion_Sequence_.Update(_233_v10, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_351_v87, (_356_v94).F33(), _238_globalState), _dafny.IntOfUint32((_233_v10).Cardinality()))).Uint32(), _236_v13)).(_dafny.Array)
				}
				return _234_v11
			})(), 2)
			(_nw49).ArraySet1(_234_v11, 3)
			_362_v98 = _nw49
			var _363_v99 D25
			_ = _363_v99
			_363_v99 = Companion_D25_.Create_DC58_(_362_v98)
			_362_v98 = (_363_v99).Dtor_cf96()
		}
		var _364_v100 _dafny.Set
		_ = _364_v100
		_364_v100 = _dafny.SetOf(_287_v44, _287_v44)
		var _365_v101 _dafny.Map
		_ = _365_v101
		_365_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v0, _364_v100)
		if (((func() _dafny.Set {
			if (_365_v101).Contains(_340_cf81) {
				return (_365_v101).Get(_340_cf81).(_dafny.Set)
			}
			return _364_v100
		})()).Cardinality()).Cmp(Companion_Default___.SafeDivisionInt((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))) >= 0 {
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_224_v1), 0))
			_ = _index39
			(_224_v1).ArraySet1(_223_v0, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_224_v1), 0))
			_ = _index40
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _index41
			var _rhs42 _dafny.CodePoint = _236_v13
			_ = _rhs42
			var _rhs43 bool = _223_v0
			_ = _rhs43
			var _rhs44 bool = !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_366_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_367_i10 _dafny.Int) _dafny.CodePoint {
					return _366_v13
				}
			})(_236_v13))), _233_v10)
			_ = _rhs44
			var _rhs45 bool = !(_340_cf81)
			_ = _rhs45
			var _rhs46 _dafny.Int = (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)
			_ = _rhs46
			var _lhs27 *GlobalState = _238_globalState
			_ = _lhs27
			var _lhs28 _dafny.Array = _224_v1
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_224_v1), 0))
			_ = _lhs29
			var _lhs30 *GlobalState = _238_globalState
			_ = _lhs30
			var _lhs31 _dafny.Array = _234_v11
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _lhs32
			_236_v13 = _rhs42
			_lhs27.F13 = _rhs43
			(_lhs28).ArraySet1(_rhs44, (_lhs29).Int())
			_lhs30.F13 = _rhs45
			(_lhs31).ArraySet1(_rhs46, (_lhs32).Int())
			var _368_v102 _dafny.MultiSet
			_ = _368_v102
			_368_v102 = _270_v39
			var _369_v103 _dafny.Sequence
			_ = _369_v103
			_369_v103 = _dafny.SeqOf((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))
			var _nwElement0_10 _dafny.Int = _225_v2
			_ = _nwElement0_10
			var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(27))
			_ = _nw50
			(_nw50).ArraySet1(_nwElement0_10, 0)
			(_nw50).ArraySet1(_225_v2, 1)
			(_nw50).ArraySet1(Companion_Default___.SafeModuloInt(_225_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fghnvwo")).Cardinality())), 2)
			(_nw50).ArraySet1((_273_v42).Dtor_cf1(), 3)
			(_nw50).ArraySet1(((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).Times((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)), 4)
			(_nw50).ArraySet1(_dafny.IntOfInt64(946), 5)
			(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_225_v2, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)), 6)
			(_nw50).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_230_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32(), (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool))).Cardinality())), 7)
			(_nw50).ArraySet1((_225_v2).Plus(_225_v2), 8)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 9)
			(_nw50).ArraySet1(_225_v2, 10)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 11)
			(_nw50).ArraySet1(Companion_Default___.Fm0((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool), _225_v2, _238_globalState), 12)
			(_nw50).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm21(_267_v36, (_270_v39), _238_globalState)).Cardinality()), 13)
			(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_225_v2, (func() _dafny.Int {
				if (_228_v5).Contains((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)) {
					return (_228_v5).Get((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return _225_v2
			})()), 14)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 15)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 16)
			(_nw50).ArraySet1(_dafny.IntOfInt64(86), 17)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 18)
			(_nw50).ArraySet1(_225_v2, 19)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 20)
			(_nw50).ArraySet1(_225_v2, 21)
			(_nw50).ArraySet1(_225_v2, 22)
			(_nw50).ArraySet1((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), 23)
			(_nw50).ArraySet1(Companion_Default___.Fm0(_223_v0, (_227_v4).Cardinality(), _238_globalState), 24)
			(_nw50).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_370_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_371_i11 _dafny.Int) _dafny.CodePoint {
					return _370_v13
				}
			})(_236_v13)))).Cardinality()), (_369_v103).Select((Companion_Default___.SafeIndex((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_369_v103).Cardinality()))).Uint32()).(_dafny.Int)), 25)
			(_nw50).ArraySet1(_225_v2, 26)
			_234_v11 = _nw50
			(_238_globalState).F7 = Companion_Default___.SafeModuloInt(((_369_v103).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_369_v103).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_dafny.Zero).Minus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))), _dafny.IntOfInt64(-101))
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _index42
			(_234_v11).ArraySet1(Companion_Default___.SafeDivisionInt((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)).Minus(_225_v2))), (_index42).Int())
			(_238_globalState).F13 = _340_cf81
		} else {
			var _372_v104 _dafny.Sequence
			_ = _372_v104
			_372_v104 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _225_v2))
			var _373_v105 _dafny.Set
			_ = _373_v105
			_373_v105 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(865))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}((func(_374_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_375_i12 _dafny.Int) _dafny.Int {
					return _374_v2
				}
			})(_225_v2))))
			var _376_v107 _dafny.Map
			_ = _376_v107
			_376_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}(func(_377_i13 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(672)
			})), _340_cf81)
			var _378_v109 _dafny.Sequence
			_ = _378_v109
			_378_v109 = _dafny.SeqOf(_373_v105)
			var _379_v110 _dafny.Map
			_ = _379_v110
			_379_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_372_v104, _236_v13)
			var _380_v112 D13
			_ = _380_v112
			_380_v112 = Companion_D13_.Create_DC28_(_340_cf81, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _225_v2)
			var _381_v114 _dafny.Map
			_ = _381_v114
			_381_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v0, _233_v10)
			var _382_v115 D2
			_ = _382_v115
			_382_v115 = Companion_D2_.Create_DC6_(_381_v114)
			var _383_v116 _dafny.Array
			_ = _383_v116
			var _nwElement0_11 _dafny.Set = func() _dafny.Set {
				var _coll54 = _dafny.NewBuilder()
				_ = _coll54
				for _iter54 := _dafny.Iterate((_373_v105).Elements()); ; {
					_compr_54, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _384_v106 _dafny.Sequence
					_384_v106 = interface{}(_compr_54).(_dafny.Sequence)
					if (_373_v105).Contains(_384_v106) {
						_coll54.Add(_384_v106)
					}
				}
				return _coll54.ToSet()
			}()
			_ = _nwElement0_11
			var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(22))
			_ = _nw51
			(_nw51).ArraySet1(_nwElement0_11, 0)
			(_nw51).ArraySet1((_373_v105).Intersection(_373_v105), 1)
			(_nw51).ArraySet1(_373_v105, 2)
			(_nw51).ArraySet1(_dafny.SetOf(_372_v104), 3)
			(_nw51).ArraySet1((func() _dafny.Set {
				if _340_cf81 {
					return _373_v105
				}
				return _373_v105
			})(), 4)
			(_nw51).ArraySet1(_373_v105, 5)
			(_nw51).ArraySet1((func() _dafny.Set {
				var _coll55 = _dafny.NewBuilder()
				_ = _coll55
				for _iter55 := _dafny.Iterate((_376_v107).Keys().Elements()); ; {
					_compr_55, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _385_v108 _dafny.Sequence
					_385_v108 = interface{}(_compr_55).(_dafny.Sequence)
					if (_376_v107).Contains(_385_v108) {
						_coll55.Add(_385_v108)
					}
				}
				return _coll55.ToSet()
			}()).Intersection(_373_v105), 6)
			(_nw51).ArraySet1((_378_v109).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_378_v109).Cardinality()))).Uint32()).(_dafny.Set), 7)
			(_nw51).ArraySet1((_378_v109).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_378_v109).Cardinality()))).Uint32()).(_dafny.Set), 8)
			(_nw51).ArraySet1(_373_v105, 9)
			(_nw51).ArraySet1(((_378_v109).Select((Companion_Default___.SafeIndex((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_378_v109).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_373_v105), 10)
			(_nw51).ArraySet1(func() _dafny.Set {
				var _coll56 = _dafny.NewBuilder()
				_ = _coll56
				for _iter56 := _dafny.Iterate((_379_v110).Keys().Elements()); ; {
					_compr_56, _ok56 := _iter56()
					if !_ok56 {
						break
					}
					var _386_v111 _dafny.Sequence
					_386_v111 = interface{}(_compr_56).(_dafny.Sequence)
					if (_379_v110).Contains(_386_v111) {
						_coll56.Add(_386_v111)
					}
				}
				return _coll56.ToSet()
			}(), 11)
			(_nw51).ArraySet1(_373_v105, 12)
			(_nw51).ArraySet1((_373_v105).Difference(_373_v105), 13)
			(_nw51).ArraySet1(_373_v105, 14)
			(_nw51).ArraySet1((func() _dafny.Set {
				if (_380_v112).Dtor_cf45() {
					return func() _dafny.Set {
						var _coll57 = _dafny.NewBuilder()
						_ = _coll57
						for _iter57 := _dafny.Iterate((_dafny.SeqOf((_290_v45).Fm7(_238_globalState))).Elements()); ; {
							_compr_57, _ok57 := _iter57()
							if !_ok57 {
								break
							}
							var _387_v113 _dafny.Sequence
							_387_v113 = interface{}(_compr_57).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_290_v45).Fm7(_238_globalState)), _387_v113) {
								_coll57.Add(_387_v113)
							}
						}
						return _coll57.ToSet()
					}()
				}
				return _373_v105
			})(), 15)
			(_nw51).ArraySet1(_dafny.SetOf(_372_v104, Companion_Default___.Fm45(_223_v0, _382_v115, (func() _dafny.Int {
				if (_228_v5).Contains(_dafny.IntOfInt64(591)) {
					return (_228_v5).Get(_dafny.IntOfInt64(591)).(_dafny.Int)
				}
				return _225_v2
			})(), _238_globalState)), 16)
			(_nw51).ArraySet1(_dafny.SetOf(_372_v104), 17)
			(_nw51).ArraySet1((_378_v109).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_378_v109).Cardinality()))).Uint32()).(_dafny.Set), 18)
			(_nw51).ArraySet1(_373_v105, 19)
			(_nw51).ArraySet1(_373_v105, 20)
			(_nw51).ArraySet1(_373_v105, 21)
			_383_v116 = _nw51
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_383_v116), 0))
			_ = _index43
			(_383_v116).ArraySet1(_dafny.SetOf(_372_v104), (_index43).Int())
			var _388_v117 T1
			_ = _388_v117
			var _nw52 *C10 = New_C10_()
			_ = _nw52
			_nw52.Ctor__(_223_v0, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _236_v13, _223_v0, _287_v44, _340_cf81)
			_388_v117 = _nw52
			var _389_v118 _dafny.Map
			_ = _389_v118
			_389_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_340_cf81, _388_v117)
			var _390_v119 _dafny.Array
			_ = _390_v119
			var _nw53 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
			_ = _nw53
			_390_v119 = _nw53
			var _391_v120 D18
			_ = _391_v120
			_391_v120 = Companion_D18_.Create_DC44_(Companion_D18_.Create_DC41_(_390_v119))
			var _392_v121 D11
			_ = _392_v121
			_392_v121 = Companion_D11_.Create_DC23_((_388_v117).F24(), _225_v2, _231_v8, _291_v46)
			var _393_v122 D0
			_ = _393_v122
			_393_v122 = Companion_D0_.Create_DC2_(_340_cf81, _225_v2, _340_cf81)
			var _394_v123 D27
			_ = _394_v123
			_394_v123 = Companion_D27_.Create_DC61_(_389_v118)
			var _395_v124 _dafny.Sequence
			_ = _395_v124
			_395_v124 = _dafny.SeqOf((func() _dafny.Map {
				if _340_cf81 {
					return (_394_v123).Dtor_cf99()
				}
				return _389_v118
			})(), _389_v118)
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_383_v116), 0))
			_ = _index44
			var _rhs47 bool = true
			_ = _rhs47
			var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_372_v104, _dafny.Companion_Sequence_.Concatenate(_372_v104, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.MultiSetOf(_391_v120)).Cardinality()), (Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_391_v120)).Cardinality())).Cardinality()))).Uint32(), _225_v2)))
			_ = _rhs48
			var _rhs49 _dafny.Sequence = _230_v7
			_ = _rhs49
			var _rhs50 _dafny.Set = Companion_Default___.Fm61((_392_v121).Dtor_cf37(), Companion_Default___.Fm1(_223_v0, _238_globalState), (func() D0 {
				if !(false) {
					return _393_v122
				}
				return _393_v122
			})(), _225_v2, _238_globalState)
			_ = _rhs50
			var _rhs51 _dafny.Map = (_395_v124).Select((Companion_Default___.SafeIndex((_225_v2).Plus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_395_v124).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs51
			var _lhs33 *GlobalState = _238_globalState
			_ = _lhs33
			var _lhs34 _dafny.Array = _383_v116
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_383_v116), 0))
			_ = _lhs35
			_lhs33.F13 = _rhs47
			_372_v104 = _rhs48
			_230_v7 = _rhs49
			(_lhs34).ArraySet1(_rhs50, (_lhs35).Int())
			_389_v118 = _rhs51
			(_238_globalState).F13 = (_388_v117).F24()
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _index45
			(_234_v11).ArraySet1(_225_v2, (_index45).Int())
			(_238_globalState).F13 = _223_v0
			var _396_v125 T2
			_ = _396_v125
			var _nw54 *C4 = New_C4_()
			_ = _nw54
			_nw54.Ctor__(_225_v2, _228_v5, _287_v44, _388_v117.F23(), _340_cf81)
			_396_v125 = _nw54
			var _397_v126 _dafny.Map
			_ = _397_v126
			_397_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_396_v125, _340_cf81)
			var _398_v127 bool
			_ = _398_v127
			var _399_v128 _dafny.Int
			_ = _399_v128
			var _400_v129 bool
			_ = _400_v129
			var _out18 bool
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			var _out20 bool
			_ = _out20
			_out18, _out19, _out20 = (_290_v45).M3((func() bool {
				if (_397_v126).Contains(_396_v125) {
					return (_397_v126).Get(_396_v125).(bool)
				}
				return (_396_v125).F24()
			})(), _236_v13, _238_globalState)
			_398_v127 = _out18
			_399_v128 = _out19
			_400_v129 = _out20
		}
		_223_v0 = _223_v0
	} else if _source6.Is_DC50() {
		var _401___mcc_h7 _dafny.MultiSet = _source6.Get_().(D21_DC50).Cf80
		_ = _401___mcc_h7
		var _402_cf80 _dafny.MultiSet = _401___mcc_h7
		_ = _402_cf80
		var _403_v130 *C10
		_ = _403_v130
		var _nw55 *C10 = New_C10_()
		_ = _nw55
		_nw55.Ctor__((_dafny.IntOfInt64(41)).Cmp((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)) >= 0, _225_v2, _236_v13, false, _287_v44, _223_v0)
		_403_v130 = _nw55
		var _404_v131 _dafny.Sequence
		_ = _404_v131
		_404_v131 = _dafny.SeqOf(_dafny.IntOfInt64(685), _dafny.IntOfInt64(402), _225_v2)
		_404_v131 = _404_v131
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
		_ = _index46
		(_234_v11).ArraySet1(_403_v130.F29, (_index46).Int())
		_233_v10 = _dafny.UnicodeSeqOfUtf8Bytes("ueed")
	} else {
		var _405___mcc_h8 D21 = _source6.Get_().(D21_DC52).Cf82
		_ = _405___mcc_h8
		var _406_cf82 D21 = _405___mcc_h8
		_ = _406_cf82
		var _407_v132 _dafny.Array
		_ = _407_v132
		var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw56
		_407_v132 = _nw56
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_407_v132), 0))
		_ = _index47
		(_407_v132).ArraySet1(_234_v11, (_index47).Int())
		var _408_v133 _dafny.Map
		_ = _408_v133
		_408_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _234_v11)
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_407_v132), 0))
		_ = _index48
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
		_ = _index49
		var _rhs52 _dafny.Array = _234_v11
		_ = _rhs52
		var _rhs53 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xy")).Cardinality())
		_ = _rhs53
		var _rhs54 _dafny.Int = _225_v2
		_ = _rhs54
		var _rhs55 _dafny.Map = _408_v133
		_ = _rhs55
		var _lhs36 _dafny.Array = _407_v132
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_407_v132), 0))
		_ = _lhs37
		var _lhs38 _dafny.Array = _234_v11
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
		_ = _lhs39
		(_lhs36).ArraySet1(_rhs52, (_lhs37).Int())
		_225_v2 = _rhs53
		(_lhs38).ArraySet1(_rhs54, (_lhs39).Int())
		_408_v133 = _rhs55
		var _409_v134 *C7
		_ = _409_v134
		var _nw57 *C7 = New_C7_()
		_ = _nw57
		_nw57.Ctor__(!((func() bool {
			if _223_v0 {
				return _223_v0
			}
			return _223_v0
		})()), false, _287_v44, _236_v13, _dafny.Companion_Sequence_.IsProperPrefixOf(_230_v7, _230_v7), (_290_v45).Fm4(_225_v2, _230_v7, _238_globalState))
		_409_v134 = _nw57
		_233_v10 = _dafny.UnicodeSeqOfUtf8Bytes("f")
		_233_v10 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_233_v10, (Companion_Default___.SafeIndex((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_233_v10).Cardinality()))).Uint32(), _236_v13), _233_v10)
	}
	var _rhs56 _dafny.Int = (_dafny.Zero).Minus(_225_v2)
	_ = _rhs56
	var _rhs57 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kbtyeu"), _233_v10), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(_223_v0, _238_globalState), (Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((Companion_Default___.Fm1(_223_v0, _238_globalState)).Cardinality()))).Uint32(), _236_v13))
	_ = _rhs57
	var _lhs40 *GlobalState = _238_globalState
	_ = _lhs40
	_lhs40.F7 = _rhs56
	_223_v0 = _rhs57
	var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
	_ = _index50
	(_224_v1).ArraySet1((func() bool {
		if _223_v0 {
			return _223_v0
		}
		return _223_v0
	})(), (_index50).Int())
	var _410_v135 D18
	_ = _410_v135
	_410_v135 = Companion_D18_.Create_DC42_((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _223_v0)
	var _411_v136 _dafny.Set
	_ = _411_v136
	_411_v136 = _dafny.SetOf(_410_v135)
	var _412_v137 _dafny.Map
	_ = _412_v137
	_412_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_410_v135, (_267_v36).Cardinality())
	var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
	_ = _index51
	(_224_v1).ArraySet1(!(_411_v136).Equals(func() _dafny.Set {
		var _coll58 = _dafny.NewBuilder()
		_ = _coll58
		for _iter58 := _dafny.Iterate((_412_v137).Keys().Elements()); ; {
			_compr_58, _ok58 := _iter58()
			if !_ok58 {
				break
			}
			var _413_v138 D18
			_413_v138 = interface{}(_compr_58).(D18)
			if (_412_v137).Contains(_413_v138) {
				_coll58.Add(_413_v138)
			}
		}
		return _coll58.ToSet()
	}()), (_index51).Int())
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_224_v1), 0))); ; {
		_guard_loop_0, _ok59 := _iter59()
		if !_ok59 {
			break
		}
		var _414_i14 _dafny.Int
		_414_i14 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_414_i14).Sign() != -1) && ((_414_i14).Cmp(_dafny.ArrayLen((_224_v1), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_224_v1, (_414_i14).Int(), (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)))
		}
	}
	for _iter60 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok60 := _iter60()
		if !_ok60 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	var _415_v139 D9
	_ = _415_v139
	_415_v139 = Companion_D9_.Create_DC20_(_233_v10, _223_v0, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))
	var _416_v140 D4
	_ = _416_v140
	_416_v140 = Companion_D4_.Create_DC10_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jnm"), (Companion_Default___.SafeIndex((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jnm")).Cardinality()))).Uint32(), _236_v13), (_415_v139).Dtor_cf31()))
	var _source7 D4 = _416_v140
	_ = _source7
	if _source7.Is_DC11() {
		var _417___mcc_h9 _dafny.Int = _source7.Get_().(D4_DC11).Cf18
		_ = _417___mcc_h9
		var _418_cf18 _dafny.Int = _417___mcc_h9
		_ = _418_cf18
		var _419_v141 _dafny.Sequence
		_ = _419_v141
		_419_v141 = _dafny.SeqOf((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _418_cf18, (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))
		var _420_v142 _dafny.Map
		_ = _420_v142
		_420_v142 = _228_v5
		var _421_v143 *C4
		_ = _421_v143
		var _nw58 *C4 = New_C4_()
		_ = _nw58
		_nw58.Ctor__(_dafny.IntOfUint32((_419_v141).Cardinality()), _420_v142, _287_v44, _236_v13, (_290_v45).Fm4((func() _dafny.Int {
			if (_228_v5).Contains(_225_v2) {
				return (_228_v5).Get(_225_v2).(_dafny.Int)
			}
			return (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)
		})(), _230_v7, _238_globalState))
		_421_v143 = _nw58
		var _422_v144 D24
		_ = _422_v144
		_422_v144 = Companion_D24_.Create_DC56_(_421_v143)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
		_ = _index52
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
		_ = _index53
		var _rhs58 bool = (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)
		_ = _rhs58
		var _rhs59 D24 = _422_v144
		_ = _rhs59
		var _rhs60 bool = _223_v0
		_ = _rhs60
		var _rhs61 bool = (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)
		_ = _rhs61
		var _lhs41 _dafny.Array = _224_v1
		_ = _lhs41
		var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
		_ = _lhs42
		var _lhs43 _dafny.Array = _224_v1
		_ = _lhs43
		var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
		_ = _lhs44
		var _lhs45 *GlobalState = _238_globalState
		_ = _lhs45
		(_lhs41).ArraySet1(_rhs58, (_lhs42).Int())
		_422_v144 = _rhs59
		(_lhs43).ArraySet1(_rhs60, (_lhs44).Int())
		_lhs45.F13 = _rhs61
		(_238_globalState).F7 = (_dafny.Zero).Minus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))
		var _423_v145 bool
		_ = _423_v145
		var _424_v146 _dafny.CodePoint
		_ = _424_v146
		var _425_v147 bool
		_ = _425_v147
		var _out21 bool
		_ = _out21
		var _out22 _dafny.CodePoint
		_ = _out22
		var _out23 bool
		_ = _out23
		_out21, _out22, _out23 = (_290_v45).M11((func() bool {
			if (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool) {
				return _223_v0
			}
			return _223_v0
		})(), _238_globalState)
		_423_v145 = _out21
		_424_v146 = _out22
		_425_v147 = _out23
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
		_ = _index54
		(_224_v1).ArraySet1(!(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_230_v7, _dafny.SeqOf((_230_v7).Select((Companion_Default___.SafeIndex(_418_cf18, _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32()).(bool))), _dafny.Companion_Sequence_.Concatenate(_230_v7, _230_v7))), (_index54).Int())
	} else {
		var _426___mcc_h10 _dafny.Sequence = _source7.Get_().(D4_DC10).Cf17
		_ = _426___mcc_h10
		var _427_cf17 _dafny.Sequence = _426___mcc_h10
		_ = _427_cf17
		var _428_v148 _dafny.Array
		_ = _428_v148
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_7
		var _nw59 _dafny.Array
		_ = _nw59
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw59 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Map = (func(_429_v0 bool, _430_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
				return func(_431_i15 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_v0, _430_v13)
				}
			})(_223_v0, _236_v13)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw59 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw59).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw59).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_428_v148 = _nw59
		var _432_v149 D14
		_ = _432_v149
		_432_v149 = Companion_D14_.Create_DC30_(_428_v148)
		var _433_v150 D14
		_ = _433_v150
		_433_v150 = Companion_D14_.Create_DC32_(_432_v149)
		var _434_v151 D14
		_ = _434_v151
		_434_v151 = Companion_D14_.Create_DC32_((func() D14 {
			if _223_v0 {
				return _433_v150
			}
			return Companion_D14_.Create_DC31_(_223_v0)
		})())
		var _435_v152 _dafny.Map
		_ = _435_v152
		_435_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _225_v2)
		var _436_v153 *C4
		_ = _436_v153
		var _nw60 *C4 = New_C4_()
		_ = _nw60
		_nw60.Ctor__(_225_v2, _435_v152, _287_v44, _236_v13, true)
		_436_v153 = _nw60
		var _437_v155 _dafny.Map
		_ = _437_v155
		_437_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_225_v2, _410_v135)
		var _438_v156 _dafny.Sequence
		_ = _438_v156
		_438_v156 = _dafny.SeqOf(Companion_Default___.Fm0((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool), (_436_v153).F33(), _238_globalState))
		var _rhs62 bool = (_290_v45).Fm4((_225_v2).Plus((_dafny.MultiSetOf(_436_v153, _436_v153)).Cardinality()), _230_v7, _238_globalState)
		_ = _rhs62
		var _rhs63 D14 = _434_v151
		_ = _rhs63
		var _rhs64 bool = (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)
		_ = _rhs64
		var _rhs65 _dafny.Int = (func() _dafny.Map {
			var _coll59 = _dafny.NewMapBuilder()
			_ = _coll59
			for _iter61 := _dafny.Iterate(((_437_v155).Update((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _410_v135)).Keys().Elements()); ; {
				_compr_59, _ok61 := _iter61()
				if !_ok61 {
					break
				}
				var _439_v154 _dafny.Int
				_439_v154 = interface{}(_compr_59).(_dafny.Int)
				if ((_437_v155).Update((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _410_v135)).Contains(_439_v154) {
					_coll59.Add(Companion_Default___.SafeModuloInt(_439_v154, (_436_v153).F33()), _236_v13)
				}
			}
			return _coll59.ToMap()
		}()).Cardinality()
		_ = _rhs65
		var _rhs66 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
			if !((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)) {
				return _438_v156
			}
			return _438_v156
		})(), _dafny.Companion_Sequence_.Update(_438_v156, (Companion_Default___.SafeIndex((_436_v153).F33(), _dafny.IntOfUint32((_438_v156).Cardinality()))).Uint32(), (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)))
		_ = _rhs66
		var _lhs46 *GlobalState = _238_globalState
		_ = _lhs46
		var _lhs47 *GlobalState = _238_globalState
		_ = _lhs47
		var _lhs48 *GlobalState = _238_globalState
		_ = _lhs48
		var _lhs49 *GlobalState = _238_globalState
		_ = _lhs49
		_lhs46.F13 = _rhs62
		_434_v151 = _rhs63
		_lhs47.F13 = _rhs64
		_lhs48.F7 = _rhs65
		_lhs49.F13 = _rhs66
		if !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_427_cf17, _233_v10), _dafny.CodePoint('a')) {
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))
			_ = _index55
			(_224_v1).ArraySet1((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool), (_index55).Int())
			(_238_globalState).F7 = Companion_Default___.SafeModuloInt((_225_v2).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_230_v7, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_427_cf17).Cardinality()), _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32(), _223_v0)).Cardinality())), (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))
			var _440_v157 *C8
			_ = _440_v157
			var _nw61 *C8 = New_C8_()
			_ = _nw61
			_nw61.Ctor__(_223_v0)
			_440_v157 = _nw61
			var _441_v158 *C8
			_ = _441_v158
			_441_v158 = _440_v157
			var _442_v159 _dafny.Map
			_ = _442_v159
			_442_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_436_v153).F33(), (_441_v158))
			_442_v159 = _442_v159
			_225_v2 = _225_v2
			var _443_v160 *C3
			_ = _443_v160
			_443_v160 = _290_v45
			_290_v45 = (_443_v160)
		} else {
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))
			_ = _index56
			(_234_v11).ArraySet1(((_dafny.Zero).Minus((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))).Times(_225_v2), (_index56).Int())
			(_238_globalState).F7 = (_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)
			var _444_v161 D14
			_ = _444_v161
			_444_v161 = Companion_D14_.Create_DC30_(_428_v148)
			_444_v161 = _444_v161
			var _445_v162 bool
			_ = _445_v162
			var _446_v163 _dafny.Int
			_ = _446_v163
			var _447_v164 bool
			_ = _447_v164
			var _out24 bool
			_ = _out24
			var _out25 _dafny.Int
			_ = _out25
			var _out26 bool
			_ = _out26
			_out24, _out25, _out26 = (_436_v153).M3((_225_v2).Cmp((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int)) > 0, _236_v13, _238_globalState)
			_445_v162 = _out24
			_446_v163 = _out25
			_447_v164 = _out26
			(_238_globalState).F13 = (_230_v7).Select((Companion_Default___.SafeIndex(_225_v2, _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32()).(bool)
		}
		var _448_v165 _dafny.MultiSet
		_ = _448_v165
		_448_v165 = _dafny.MultiSetOf(_223_v0)
		var _449_v166 *C2
		_ = _449_v166
		var _nw62 *C2 = New_C2_()
		_ = _nw62
		_nw62.Ctor__(_448_v165, (func() _dafny.Array {
			if (_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool) {
				return _287_v44
			}
			return _287_v44
		})(), _236_v13, _223_v0)
		_449_v166 = _nw62
		(_238_globalState).F7 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _225_v2), _dafny.SeqOf((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int))), !(_223_v0))).Cardinality()
	}
	var _450_v167 _dafny.Map
	_ = _450_v167
	_450_v167 = _228_v5
	var _source8 _dafny.Map = _450_v167
	_ = _source8
	var _451___mcc_h11 _dafny.Map = _source8
	_ = _451___mcc_h11
	var _452_cf16 _dafny.Map = _451___mcc_h11
	_ = _452_cf16
	var _453_v168 _dafny.Int
	_ = _453_v168
	var _454_v169 bool
	_ = _454_v169
	var _455_v170 _dafny.Int
	_ = _455_v170
	var _out27 _dafny.Int
	_ = _out27
	var _out28 bool
	_ = _out28
	var _out29 _dafny.Int
	_ = _out29
	_out27, _out28, _out29 = Companion_Default___.M0((_234_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_234_v11), 0))).Int()).(_dafny.Int), _dafny.MultiSetOf((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)), ((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool)) == (true), _225_v2, _238_globalState)
	_453_v168 = _out27
	_454_v169 = _out28
	_455_v170 = _out29
	var _456_v171 _dafny.MultiSet
	_ = _456_v171
	_456_v171 = _dafny.MultiSetOf(_234_v11, _234_v11)
	_456_v171 = _456_v171
	(_238_globalState).F7 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_238_globalState), _223_v0)).Cardinality()
	var _457_v172 _dafny.MultiSet
	_ = _457_v172
	_457_v172 = _dafny.MultiSetOf((_224_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_224_v1), 0))).Int()).(bool), _454_v169, _223_v0)
	var _458_v173 *C2
	_ = _458_v173
	var _nw63 *C2 = New_C2_()
	_ = _nw63
	_nw63.Ctor__(_457_v172, _287_v44, (func() _dafny.CodePoint {
		if (_290_v45).Fm4(_455_v170, _230_v7, _238_globalState) {
			return _236_v13
		}
		return _dafny.CodePoint('a')
	})(), (false) == (!(_223_v0)))
	_458_v173 = _nw63
	_dafny.Print(_223_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_225_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_226_v3).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(298), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(230), _dafny.IntOfInt64(230))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_230_v7, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_231_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_233_v10.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v11).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v12).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_236_v13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-230), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F0()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F2()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F6()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_238_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F12().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_238_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F15()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F18().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_globalState.F22).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v36).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v37).Equals(_dafny.SetOf(_dafny.Zero, _dafny.IntOfInt64(230))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v38).Equals(_dafny.SetOf(_dafny.IntOfInt64(462), _dafny.IntOfInt64(2), _dafny.IntOfInt64(230))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v39).Equals(_dafny.MultiSetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_271_v40).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("lpbndwsy"), _dafny.UnicodeSeqOfUtf8Bytes("cwwgiekbdyjkt"), _dafny.UnicodeSeqOfUtf8Bytes("wjoqku"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v41).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("lpbndwsy"), _dafny.UnicodeSeqOfUtf8Bytes("cwwgiekbdyjkt"), _dafny.UnicodeSeqOfUtf8Bytes("wjoqku"), _dafny.UnicodeSeqOfUtf8Bytes("wjoqku"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v42).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v43).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v44).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v44).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v44).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v44).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v44).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_290_v45).F35()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_291_v46).Dtor_cf23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_292_v47).Equals(_dafny.SetOf(Companion_D5_.Create_DC14_(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v49).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), _dafny.NewMapBuilder().ToMap())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_297_v52, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), _dafny.NewMapBuilder().ToMap()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(230), _dafny.IntOfInt64(230))).UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(230), _dafny.IntOfInt64(230))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), _dafny.NewMapBuilder().ToMap()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(655), _dafny.IntOfInt64(-349))).UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-697), _dafny.IntOfInt64(-69))).UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-893), _dafny.IntOfInt64(6))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_D5_.Create_DC14_(true)), _dafny.NewMapBuilder().ToMap()))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_301_v53).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence)).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_302_v54).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_303_v55).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_338_v76).Dtor_cf81())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_410_v135).Dtor_cf68())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_410_v135).Dtor_cf69())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_411_v136).Equals(_dafny.SetOf(Companion_D18_.Create_DC42_(_dafny.Zero, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_412_v137).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D18_.Create_DC42_(_dafny.Zero, true), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_415_v139).Dtor_cf31().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_415_v139).Dtor_cf32())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_415_v139).Dtor_cf33())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_416_v140).Dtor_cf17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_450_v167).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(230), _dafny.IntOfInt64(230))))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
	Cf3 _dafny.Int
	Cf4 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool, Cf3 _dafny.Int, Cf4 bool) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.MultiSet
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.MultiSet) D0 {
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
	return Companion_D0_.Create_DC1_(_dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf0() _dafny.MultiSet {
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
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
	Cf7  _dafny.Int
	Cf8  _dafny.Int
	Cf9  _dafny.Sequence
	Cf10 _dafny.Array
	Cf11 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 _dafny.Int, Cf8 _dafny.Int, Cf9 _dafny.Sequence, Cf10 _dafny.Array, Cf11 _dafny.Int) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf6 _dafny.CodePoint
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.CodePoint) D1 {
	return D1{D1_DC4{Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(_dafny.Zero, _dafny.Zero, _dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Equals(data2.Cf9) && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6 == data2.Cf6
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
	Cf13 bool
	Cf14 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 bool, Cf14 bool) D2 {
	return D2{D2_DC7{Cf13, Cf14}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf12 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf12 _dafny.Map) D2 {
	return D2{D2_DC6{Cf12}}
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
	return Companion_D2_.Create_DC7_(false, false)
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() bool {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf12
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
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf16 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf16 _dafny.Map) D3 {
	return D3{D3_DC9{Cf16}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf16() _dafny.Map {
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
	Cf18 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf18 _dafny.Int) D4 {
	return D4{D4_DC11{Cf18}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf17 _dafny.Sequence
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf17 _dafny.Sequence) D4 {
	return D4{D4_DC10{Cf17}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.Zero)
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + data.Cf17.VerbatimString(true) + ")"
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

type D5_DC13 struct {
	Cf20 _dafny.Int
	Cf21 bool
	Cf22 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf20 _dafny.Int, Cf21 bool, Cf22 bool) D5 {
	return D5{D5_DC13{Cf20, Cf21, Cf22}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf23 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf23 bool) D5 {
	return D5{D5_DC14{Cf23}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf19 _dafny.Set
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf19 _dafny.Set) D5 {
	return D5{D5_DC12{Cf19}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(_dafny.Zero, false, false)
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf22() bool {
	return _this.Get_().(D5_DC13).Cf22
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC14).Cf23
}

func (_this D5) Dtor_cf19() _dafny.Set {
	return _this.Get_().(D5_DC12).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
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
	Cf25 _dafny.Map
	Cf26 bool
	Cf27 bool
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf25 _dafny.Map, Cf26 bool, Cf27 bool) D6 {
	return D6{D6_DC16{Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf24 _dafny.Map
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf24 _dafny.Map) D6 {
	return D6{D6_DC15{Cf24}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.EmptyMap, false, false)
}

func (_this D6) Dtor_cf25() _dafny.Map {
	return _this.Get_().(D6_DC16).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC16).Cf26
}

func (_this D6) Dtor_cf27() bool {
	return _this.Get_().(D6_DC16).Cf27
}

func (_this D6) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D6_DC15).Cf24
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf28 *C1
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf28 *C1) D7 {
	return D7{D7_DC17{Cf28}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D7) Dtor_cf28() *C1 {
	return _this.Get_().(D7_DC17).Cf28
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf28 == data2.Cf28
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

type D8_DC18 struct {
	Cf29 _dafny.Sequence
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf29 _dafny.Sequence) D8 {
	return D8{D8_DC18{Cf29}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D8_DC18).Cf29
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf29.Equals(data2.Cf29)
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

type D9_DC20 struct {
	Cf31 _dafny.Sequence
	Cf32 bool
	Cf33 _dafny.Int
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf31 _dafny.Sequence, Cf32 bool, Cf33 _dafny.Int) D9 {
	return D9{D9_DC20{Cf31, Cf32, Cf33}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC19 struct {
	Cf30 _dafny.Map
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf30 _dafny.Map) D9 {
	return D9{D9_DC19{Cf30}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC20_(_dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D9) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D9_DC20).Cf31
}

func (_this D9) Dtor_cf32() bool {
	return _this.Get_().(D9_DC20).Cf32
}

func (_this D9) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D9_DC20).Cf33
}

func (_this D9) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D9_DC19).Cf30
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC20:
		{
			return "D9.DC20" + "(" + data.Cf31.VerbatimString(true) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf31.Equals(data2.Cf31) && data1.Cf32 == data2.Cf32 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf34 _dafny.Sequence
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf34 _dafny.Sequence) D10 {
	return D10{D10_DC21{Cf34}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D10) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D10_DC21).Cf34
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf34) + ")"
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

type D11_DC23 struct {
	Cf36 bool
	Cf37 _dafny.Int
	Cf38 _dafny.Map
	Cf39 D5
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf36 bool, Cf37 _dafny.Int, Cf38 _dafny.Map, Cf39 D5) D11 {
	return D11{D11_DC23{Cf36, Cf37, Cf38, Cf39}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

type D11_DC22 struct {
	Cf35 _dafny.Array
}

func (D11_DC22) isD11() {}

func (CompanionStruct_D11_) Create_DC22_(Cf35 _dafny.Array) D11 {
	return D11{D11_DC22{Cf35}}
}

func (_this D11) Is_DC22() bool {
	_, ok := _this.Get_().(D11_DC22)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC23_(false, _dafny.Zero, _dafny.EmptyMap, Companion_D5_.Default())
}

func (_this D11) Dtor_cf36() bool {
	return _this.Get_().(D11_DC23).Cf36
}

func (_this D11) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D11_DC23).Cf37
}

func (_this D11) Dtor_cf38() _dafny.Map {
	return _this.Get_().(D11_DC23).Cf38
}

func (_this D11) Dtor_cf39() D5 {
	return _this.Get_().(D11_DC23).Cf39
}

func (_this D11) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D11_DC22).Cf35
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D11_DC22:
		{
			return "D11.DC22" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Equals(data2.Cf38) && data1.Cf39.Equals(data2.Cf39)
		}
	case D11_DC22:
		{
			data2, ok := other.Get_().(D11_DC22)
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

type D12_DC25 struct {
	Cf41 bool
	Cf42 _dafny.CodePoint
}

func (D12_DC25) isD12() {}

func (CompanionStruct_D12_) Create_DC25_(Cf41 bool, Cf42 _dafny.CodePoint) D12 {
	return D12{D12_DC25{Cf41, Cf42}}
}

func (_this D12) Is_DC25() bool {
	_, ok := _this.Get_().(D12_DC25)
	return ok
}

type D12_DC24 struct {
	Cf40 _dafny.Array
}

func (D12_DC24) isD12() {}

func (CompanionStruct_D12_) Create_DC24_(Cf40 _dafny.Array) D12 {
	return D12{D12_DC24{Cf40}}
}

func (_this D12) Is_DC24() bool {
	_, ok := _this.Get_().(D12_DC24)
	return ok
}

type D12_DC26 struct {
	Cf43 D12
}

func (D12_DC26) isD12() {}

func (CompanionStruct_D12_) Create_DC26_(Cf43 D12) D12 {
	return D12{D12_DC26{Cf43}}
}

func (_this D12) Is_DC26() bool {
	_, ok := _this.Get_().(D12_DC26)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC25_(false, _dafny.CodePoint('D'))
}

func (_this D12) Dtor_cf41() bool {
	return _this.Get_().(D12_DC25).Cf41
}

func (_this D12) Dtor_cf42() _dafny.CodePoint {
	return _this.Get_().(D12_DC25).Cf42
}

func (_this D12) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D12_DC24).Cf40
}

func (_this D12) Dtor_cf43() D12 {
	return _this.Get_().(D12_DC26).Cf43
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC25:
		{
			return "D12.DC25" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D12_DC24:
		{
			return "D12.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D12_DC26:
		{
			return "D12.DC26" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC25:
		{
			data2, ok := other.Get_().(D12_DC25)
			return ok && data1.Cf41 == data2.Cf41 && data1.Cf42 == data2.Cf42
		}
	case D12_DC24:
		{
			data2, ok := other.Get_().(D12_DC24)
			return ok && data1.Cf40 == data2.Cf40
		}
	case D12_DC26:
		{
			data2, ok := other.Get_().(D12_DC26)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D13_DC28 struct {
	Cf45 bool
	Cf46 _dafny.Int
	Cf47 _dafny.Int
	Cf48 _dafny.Int
}

func (D13_DC28) isD13() {}

func (CompanionStruct_D13_) Create_DC28_(Cf45 bool, Cf46 _dafny.Int, Cf47 _dafny.Int, Cf48 _dafny.Int) D13 {
	return D13{D13_DC28{Cf45, Cf46, Cf47, Cf48}}
}

func (_this D13) Is_DC28() bool {
	_, ok := _this.Get_().(D13_DC28)
	return ok
}

type D13_DC29 struct {
	Cf49 bool
	Cf50 bool
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf49 bool, Cf50 bool) D13 {
	return D13{D13_DC29{Cf49, Cf50}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

type D13_DC27 struct {
	Cf44 _dafny.Map
}

func (D13_DC27) isD13() {}

func (CompanionStruct_D13_) Create_DC27_(Cf44 _dafny.Map) D13 {
	return D13{D13_DC27{Cf44}}
}

func (_this D13) Is_DC27() bool {
	_, ok := _this.Get_().(D13_DC27)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC28_(false, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D13) Dtor_cf45() bool {
	return _this.Get_().(D13_DC28).Cf45
}

func (_this D13) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D13_DC28).Cf46
}

func (_this D13) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D13_DC28).Cf47
}

func (_this D13) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D13_DC28).Cf48
}

func (_this D13) Dtor_cf49() bool {
	return _this.Get_().(D13_DC29).Cf49
}

func (_this D13) Dtor_cf50() bool {
	return _this.Get_().(D13_DC29).Cf50
}

func (_this D13) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D13_DC27).Cf44
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC28:
		{
			return "D13.DC28" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D13_DC27:
		{
			return "D13.DC27" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC28:
		{
			data2, ok := other.Get_().(D13_DC28)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50
		}
	case D13_DC27:
		{
			data2, ok := other.Get_().(D13_DC27)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D14_DC31 struct {
	Cf52 bool
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf52 bool) D14 {
	return D14{D14_DC31{Cf52}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

type D14_DC30 struct {
	Cf51 _dafny.Array
}

func (D14_DC30) isD14() {}

func (CompanionStruct_D14_) Create_DC30_(Cf51 _dafny.Array) D14 {
	return D14{D14_DC30{Cf51}}
}

func (_this D14) Is_DC30() bool {
	_, ok := _this.Get_().(D14_DC30)
	return ok
}

type D14_DC32 struct {
	Cf53 D14
}

func (D14_DC32) isD14() {}

func (CompanionStruct_D14_) Create_DC32_(Cf53 D14) D14 {
	return D14{D14_DC32{Cf53}}
}

func (_this D14) Is_DC32() bool {
	_, ok := _this.Get_().(D14_DC32)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC31_(false)
}

func (_this D14) Dtor_cf52() bool {
	return _this.Get_().(D14_DC31).Cf52
}

func (_this D14) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D14_DC30).Cf51
}

func (_this D14) Dtor_cf53() D14 {
	return _this.Get_().(D14_DC32).Cf53
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D14_DC30:
		{
			return "D14.DC30" + "(" + _dafny.String(data.Cf51) + ")"
		}
	case D14_DC32:
		{
			return "D14.DC32" + "(" + _dafny.String(data.Cf53) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf52 == data2.Cf52
		}
	case D14_DC30:
		{
			data2, ok := other.Get_().(D14_DC30)
			return ok && data1.Cf51 == data2.Cf51
		}
	case D14_DC32:
		{
			data2, ok := other.Get_().(D14_DC32)
			return ok && data1.Cf53.Equals(data2.Cf53)
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
	Cf55 bool
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_(Cf55 bool) D15 {
	return D15{D15_DC34{Cf55}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

type D15_DC35 struct {
	Cf56 _dafny.Int
	Cf57 _dafny.Int
	Cf58 _dafny.Int
}

func (D15_DC35) isD15() {}

func (CompanionStruct_D15_) Create_DC35_(Cf56 _dafny.Int, Cf57 _dafny.Int, Cf58 _dafny.Int) D15 {
	return D15{D15_DC35{Cf56, Cf57, Cf58}}
}

func (_this D15) Is_DC35() bool {
	_, ok := _this.Get_().(D15_DC35)
	return ok
}

type D15_DC36 struct {
	Cf59 bool
	Cf60 _dafny.Int
}

func (D15_DC36) isD15() {}

func (CompanionStruct_D15_) Create_DC36_(Cf59 bool, Cf60 _dafny.Int) D15 {
	return D15{D15_DC36{Cf59, Cf60}}
}

func (_this D15) Is_DC36() bool {
	_, ok := _this.Get_().(D15_DC36)
	return ok
}

type D15_DC33 struct {
	Cf54 _dafny.Map
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf54 _dafny.Map) D15 {
	return D15{D15_DC33{Cf54}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC34_(false)
}

func (_this D15) Dtor_cf55() bool {
	return _this.Get_().(D15_DC34).Cf55
}

func (_this D15) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D15_DC35).Cf56
}

func (_this D15) Dtor_cf57() _dafny.Int {
	return _this.Get_().(D15_DC35).Cf57
}

func (_this D15) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D15_DC35).Cf58
}

func (_this D15) Dtor_cf59() bool {
	return _this.Get_().(D15_DC36).Cf59
}

func (_this D15) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D15_DC36).Cf60
}

func (_this D15) Dtor_cf54() _dafny.Map {
	return _this.Get_().(D15_DC33).Cf54
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC34:
		{
			return "D15.DC34" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D15_DC35:
		{
			return "D15.DC35" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ")"
		}
	case D15_DC36:
		{
			return "D15.DC36" + "(" + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf54) + ")"
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
			return ok && data1.Cf55 == data2.Cf55
		}
	case D15_DC35:
		{
			data2, ok := other.Get_().(D15_DC35)
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0 && data1.Cf57.Cmp(data2.Cf57) == 0 && data1.Cf58.Cmp(data2.Cf58) == 0
		}
	case D15_DC36:
		{
			data2, ok := other.Get_().(D15_DC36)
			return ok && data1.Cf59 == data2.Cf59 && data1.Cf60.Cmp(data2.Cf60) == 0
		}
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D16_DC38 struct {
	Cf62 _dafny.Sequence
	Cf63 _dafny.Array
	Cf64 _dafny.Set
}

func (D16_DC38) isD16() {}

func (CompanionStruct_D16_) Create_DC38_(Cf62 _dafny.Sequence, Cf63 _dafny.Array, Cf64 _dafny.Set) D16 {
	return D16{D16_DC38{Cf62, Cf63, Cf64}}
}

func (_this D16) Is_DC38() bool {
	_, ok := _this.Get_().(D16_DC38)
	return ok
}

type D16_DC37 struct {
	Cf61 *C5
}

func (D16_DC37) isD16() {}

func (CompanionStruct_D16_) Create_DC37_(Cf61 *C5) D16 {
	return D16{D16_DC37{Cf61}}
}

func (_this D16) Is_DC37() bool {
	_, ok := _this.Get_().(D16_DC37)
	return ok
}

type D16_DC39 struct {
	Cf65 D16
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf65 D16) D16 {
	return D16{D16_DC39{Cf65}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC38_(_dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySet)
}

func (_this D16) Dtor_cf62() _dafny.Sequence {
	return _this.Get_().(D16_DC38).Cf62
}

func (_this D16) Dtor_cf63() _dafny.Array {
	return _this.Get_().(D16_DC38).Cf63
}

func (_this D16) Dtor_cf64() _dafny.Set {
	return _this.Get_().(D16_DC38).Cf64
}

func (_this D16) Dtor_cf61() *C5 {
	return _this.Get_().(D16_DC37).Cf61
}

func (_this D16) Dtor_cf65() D16 {
	return _this.Get_().(D16_DC39).Cf65
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC38:
		{
			return "D16.DC38" + "(" + data.Cf62.VerbatimString(true) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D16_DC37:
		{
			return "D16.DC37" + "(" + _dafny.String(data.Cf61) + ")"
		}
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC38:
		{
			data2, ok := other.Get_().(D16_DC38)
			return ok && data1.Cf62.Equals(data2.Cf62) && data1.Cf63 == data2.Cf63 && data1.Cf64.Equals(data2.Cf64)
		}
	case D16_DC37:
		{
			data2, ok := other.Get_().(D16_DC37)
			return ok && data1.Cf61 == data2.Cf61
		}
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D17_DC40 struct {
	Cf66 T0
}

func (D17_DC40) isD17() {}

func (CompanionStruct_D17_) Create_DC40_(Cf66 T0) D17 {
	return D17{D17_DC40{Cf66}}
}

func (_this D17) Is_DC40() bool {
	_, ok := _this.Get_().(D17_DC40)
	return ok
}

func (CompanionStruct_D17_) Default() T0 {
	return (T0)(nil)
}

func (_this D17) Dtor_cf66() T0 {
	return _this.Get_().(D17_DC40).Cf66
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC40:
		{
			return "D17.DC40" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC40:
		{
			data2, ok := other.Get_().(D17_DC40)
			return ok && _dafny.AreEqual(data1.Cf66, data2.Cf66)
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

type D18_DC42 struct {
	Cf68 _dafny.Int
	Cf69 bool
}

func (D18_DC42) isD18() {}

func (CompanionStruct_D18_) Create_DC42_(Cf68 _dafny.Int, Cf69 bool) D18 {
	return D18{D18_DC42{Cf68, Cf69}}
}

func (_this D18) Is_DC42() bool {
	_, ok := _this.Get_().(D18_DC42)
	return ok
}

type D18_DC43 struct {
	Cf70 _dafny.Array
	Cf71 _dafny.Int
	Cf72 _dafny.Sequence
}

func (D18_DC43) isD18() {}

func (CompanionStruct_D18_) Create_DC43_(Cf70 _dafny.Array, Cf71 _dafny.Int, Cf72 _dafny.Sequence) D18 {
	return D18{D18_DC43{Cf70, Cf71, Cf72}}
}

func (_this D18) Is_DC43() bool {
	_, ok := _this.Get_().(D18_DC43)
	return ok
}

type D18_DC41 struct {
	Cf67 _dafny.Array
}

func (D18_DC41) isD18() {}

func (CompanionStruct_D18_) Create_DC41_(Cf67 _dafny.Array) D18 {
	return D18{D18_DC41{Cf67}}
}

func (_this D18) Is_DC41() bool {
	_, ok := _this.Get_().(D18_DC41)
	return ok
}

type D18_DC44 struct {
	Cf73 D18
}

func (D18_DC44) isD18() {}

func (CompanionStruct_D18_) Create_DC44_(Cf73 D18) D18 {
	return D18{D18_DC44{Cf73}}
}

func (_this D18) Is_DC44() bool {
	_, ok := _this.Get_().(D18_DC44)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC42_(_dafny.Zero, false)
}

func (_this D18) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D18_DC42).Cf68
}

func (_this D18) Dtor_cf69() bool {
	return _this.Get_().(D18_DC42).Cf69
}

func (_this D18) Dtor_cf70() _dafny.Array {
	return _this.Get_().(D18_DC43).Cf70
}

func (_this D18) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D18_DC43).Cf71
}

func (_this D18) Dtor_cf72() _dafny.Sequence {
	return _this.Get_().(D18_DC43).Cf72
}

func (_this D18) Dtor_cf67() _dafny.Array {
	return _this.Get_().(D18_DC41).Cf67
}

func (_this D18) Dtor_cf73() D18 {
	return _this.Get_().(D18_DC44).Cf73
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC42:
		{
			return "D18.DC42" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D18_DC43:
		{
			return "D18.DC43" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + data.Cf72.VerbatimString(true) + ")"
		}
	case D18_DC41:
		{
			return "D18.DC41" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D18_DC44:
		{
			return "D18.DC44" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC42:
		{
			data2, ok := other.Get_().(D18_DC42)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69 == data2.Cf69
		}
	case D18_DC43:
		{
			data2, ok := other.Get_().(D18_DC43)
			return ok && data1.Cf70 == data2.Cf70 && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72.Equals(data2.Cf72)
		}
	case D18_DC41:
		{
			data2, ok := other.Get_().(D18_DC41)
			return ok && data1.Cf67 == data2.Cf67
		}
	case D18_DC44:
		{
			data2, ok := other.Get_().(D18_DC44)
			return ok && data1.Cf73.Equals(data2.Cf73)
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

type D19_DC46 struct {
	Cf75 _dafny.Set
	Cf76 _dafny.Sequence
}

func (D19_DC46) isD19() {}

func (CompanionStruct_D19_) Create_DC46_(Cf75 _dafny.Set, Cf76 _dafny.Sequence) D19 {
	return D19{D19_DC46{Cf75, Cf76}}
}

func (_this D19) Is_DC46() bool {
	_, ok := _this.Get_().(D19_DC46)
	return ok
}

type D19_DC47 struct {
	Cf77 bool
}

func (D19_DC47) isD19() {}

func (CompanionStruct_D19_) Create_DC47_(Cf77 bool) D19 {
	return D19{D19_DC47{Cf77}}
}

func (_this D19) Is_DC47() bool {
	_, ok := _this.Get_().(D19_DC47)
	return ok
}

type D19_DC45 struct {
	Cf74 _dafny.Map
}

func (D19_DC45) isD19() {}

func (CompanionStruct_D19_) Create_DC45_(Cf74 _dafny.Map) D19 {
	return D19{D19_DC45{Cf74}}
}

func (_this D19) Is_DC45() bool {
	_, ok := _this.Get_().(D19_DC45)
	return ok
}

type D19_DC48 struct {
	Cf78 D19
}

func (D19_DC48) isD19() {}

func (CompanionStruct_D19_) Create_DC48_(Cf78 D19) D19 {
	return D19{D19_DC48{Cf78}}
}

func (_this D19) Is_DC48() bool {
	_, ok := _this.Get_().(D19_DC48)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC46_(_dafny.EmptySet, _dafny.EmptySeq)
}

func (_this D19) Dtor_cf75() _dafny.Set {
	return _this.Get_().(D19_DC46).Cf75
}

func (_this D19) Dtor_cf76() _dafny.Sequence {
	return _this.Get_().(D19_DC46).Cf76
}

func (_this D19) Dtor_cf77() bool {
	return _this.Get_().(D19_DC47).Cf77
}

func (_this D19) Dtor_cf74() _dafny.Map {
	return _this.Get_().(D19_DC45).Cf74
}

func (_this D19) Dtor_cf78() D19 {
	return _this.Get_().(D19_DC48).Cf78
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC46:
		{
			return "D19.DC46" + "(" + _dafny.String(data.Cf75) + ", " + data.Cf76.VerbatimString(true) + ")"
		}
	case D19_DC47:
		{
			return "D19.DC47" + "(" + _dafny.String(data.Cf77) + ")"
		}
	case D19_DC45:
		{
			return "D19.DC45" + "(" + _dafny.String(data.Cf74) + ")"
		}
	case D19_DC48:
		{
			return "D19.DC48" + "(" + _dafny.String(data.Cf78) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC46:
		{
			data2, ok := other.Get_().(D19_DC46)
			return ok && data1.Cf75.Equals(data2.Cf75) && data1.Cf76.Equals(data2.Cf76)
		}
	case D19_DC47:
		{
			data2, ok := other.Get_().(D19_DC47)
			return ok && data1.Cf77 == data2.Cf77
		}
	case D19_DC45:
		{
			data2, ok := other.Get_().(D19_DC45)
			return ok && data1.Cf74.Equals(data2.Cf74)
		}
	case D19_DC48:
		{
			data2, ok := other.Get_().(D19_DC48)
			return ok && data1.Cf78.Equals(data2.Cf78)
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

type D20_DC49 struct {
	Cf79 _dafny.MultiSet
}

func (D20_DC49) isD20() {}

func (CompanionStruct_D20_) Create_DC49_(Cf79 _dafny.MultiSet) D20 {
	return D20{D20_DC49{Cf79}}
}

func (_this D20) Is_DC49() bool {
	_, ok := _this.Get_().(D20_DC49)
	return ok
}

func (CompanionStruct_D20_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D20) Dtor_cf79() _dafny.MultiSet {
	return _this.Get_().(D20_DC49).Cf79
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC49:
		{
			return "D20.DC49" + "(" + _dafny.String(data.Cf79) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC49:
		{
			data2, ok := other.Get_().(D20_DC49)
			return ok && data1.Cf79.Equals(data2.Cf79)
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

type D21_DC51 struct {
	Cf81 bool
}

func (D21_DC51) isD21() {}

func (CompanionStruct_D21_) Create_DC51_(Cf81 bool) D21 {
	return D21{D21_DC51{Cf81}}
}

func (_this D21) Is_DC51() bool {
	_, ok := _this.Get_().(D21_DC51)
	return ok
}

type D21_DC50 struct {
	Cf80 _dafny.MultiSet
}

func (D21_DC50) isD21() {}

func (CompanionStruct_D21_) Create_DC50_(Cf80 _dafny.MultiSet) D21 {
	return D21{D21_DC50{Cf80}}
}

func (_this D21) Is_DC50() bool {
	_, ok := _this.Get_().(D21_DC50)
	return ok
}

type D21_DC52 struct {
	Cf82 D21
}

func (D21_DC52) isD21() {}

func (CompanionStruct_D21_) Create_DC52_(Cf82 D21) D21 {
	return D21{D21_DC52{Cf82}}
}

func (_this D21) Is_DC52() bool {
	_, ok := _this.Get_().(D21_DC52)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC51_(false)
}

func (_this D21) Dtor_cf81() bool {
	return _this.Get_().(D21_DC51).Cf81
}

func (_this D21) Dtor_cf80() _dafny.MultiSet {
	return _this.Get_().(D21_DC50).Cf80
}

func (_this D21) Dtor_cf82() D21 {
	return _this.Get_().(D21_DC52).Cf82
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC51:
		{
			return "D21.DC51" + "(" + _dafny.String(data.Cf81) + ")"
		}
	case D21_DC50:
		{
			return "D21.DC50" + "(" + _dafny.String(data.Cf80) + ")"
		}
	case D21_DC52:
		{
			return "D21.DC52" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC51:
		{
			data2, ok := other.Get_().(D21_DC51)
			return ok && data1.Cf81 == data2.Cf81
		}
	case D21_DC50:
		{
			data2, ok := other.Get_().(D21_DC50)
			return ok && data1.Cf80.Equals(data2.Cf80)
		}
	case D21_DC52:
		{
			data2, ok := other.Get_().(D21_DC52)
			return ok && data1.Cf82.Equals(data2.Cf82)
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

type D22_DC53 struct {
	Cf83 T3
}

func (D22_DC53) isD22() {}

func (CompanionStruct_D22_) Create_DC53_(Cf83 T3) D22 {
	return D22{D22_DC53{Cf83}}
}

func (_this D22) Is_DC53() bool {
	_, ok := _this.Get_().(D22_DC53)
	return ok
}

func (CompanionStruct_D22_) Default() T3 {
	return (T3)(nil)
}

func (_this D22) Dtor_cf83() T3 {
	return _this.Get_().(D22_DC53).Cf83
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC53:
		{
			return "D22.DC53" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC53:
		{
			data2, ok := other.Get_().(D22_DC53)
			return ok && _dafny.AreEqual(data1.Cf83, data2.Cf83)
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

type D23_DC55 struct {
	Cf85 D21
	Cf86 _dafny.Array
	Cf87 _dafny.MultiSet
	Cf88 bool
	Cf89 _dafny.Int
}

func (D23_DC55) isD23() {}

func (CompanionStruct_D23_) Create_DC55_(Cf85 D21, Cf86 _dafny.Array, Cf87 _dafny.MultiSet, Cf88 bool, Cf89 _dafny.Int) D23 {
	return D23{D23_DC55{Cf85, Cf86, Cf87, Cf88, Cf89}}
}

func (_this D23) Is_DC55() bool {
	_, ok := _this.Get_().(D23_DC55)
	return ok
}

type D23_DC54 struct {
	Cf84 *C6
}

func (D23_DC54) isD23() {}

func (CompanionStruct_D23_) Create_DC54_(Cf84 *C6) D23 {
	return D23{D23_DC54{Cf84}}
}

func (_this D23) Is_DC54() bool {
	_, ok := _this.Get_().(D23_DC54)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC55_(Companion_D21_.Default(), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMultiSet, false, _dafny.Zero)
}

func (_this D23) Dtor_cf85() D21 {
	return _this.Get_().(D23_DC55).Cf85
}

func (_this D23) Dtor_cf86() _dafny.Array {
	return _this.Get_().(D23_DC55).Cf86
}

func (_this D23) Dtor_cf87() _dafny.MultiSet {
	return _this.Get_().(D23_DC55).Cf87
}

func (_this D23) Dtor_cf88() bool {
	return _this.Get_().(D23_DC55).Cf88
}

func (_this D23) Dtor_cf89() _dafny.Int {
	return _this.Get_().(D23_DC55).Cf89
}

func (_this D23) Dtor_cf84() *C6 {
	return _this.Get_().(D23_DC54).Cf84
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC55:
		{
			return "D23.DC55" + "(" + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ")"
		}
	case D23_DC54:
		{
			return "D23.DC54" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC55:
		{
			data2, ok := other.Get_().(D23_DC55)
			return ok && data1.Cf85.Equals(data2.Cf85) && data1.Cf86 == data2.Cf86 && data1.Cf87.Equals(data2.Cf87) && data1.Cf88 == data2.Cf88 && data1.Cf89.Cmp(data2.Cf89) == 0
		}
	case D23_DC54:
		{
			data2, ok := other.Get_().(D23_DC54)
			return ok && data1.Cf84 == data2.Cf84
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

type D24_DC57 struct {
	Cf91 T3
	Cf92 bool
	Cf93 _dafny.CodePoint
	Cf94 _dafny.Int
	Cf95 bool
}

func (D24_DC57) isD24() {}

func (CompanionStruct_D24_) Create_DC57_(Cf91 T3, Cf92 bool, Cf93 _dafny.CodePoint, Cf94 _dafny.Int, Cf95 bool) D24 {
	return D24{D24_DC57{Cf91, Cf92, Cf93, Cf94, Cf95}}
}

func (_this D24) Is_DC57() bool {
	_, ok := _this.Get_().(D24_DC57)
	return ok
}

type D24_DC56 struct {
	Cf90 *C4
}

func (D24_DC56) isD24() {}

func (CompanionStruct_D24_) Create_DC56_(Cf90 *C4) D24 {
	return D24{D24_DC56{Cf90}}
}

func (_this D24) Is_DC56() bool {
	_, ok := _this.Get_().(D24_DC56)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC57_((T3)(nil), false, _dafny.CodePoint('D'), _dafny.Zero, false)
}

func (_this D24) Dtor_cf91() T3 {
	return _this.Get_().(D24_DC57).Cf91
}

func (_this D24) Dtor_cf92() bool {
	return _this.Get_().(D24_DC57).Cf92
}

func (_this D24) Dtor_cf93() _dafny.CodePoint {
	return _this.Get_().(D24_DC57).Cf93
}

func (_this D24) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D24_DC57).Cf94
}

func (_this D24) Dtor_cf95() bool {
	return _this.Get_().(D24_DC57).Cf95
}

func (_this D24) Dtor_cf90() *C4 {
	return _this.Get_().(D24_DC56).Cf90
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC57:
		{
			return "D24.DC57" + "(" + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ", " + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ")"
		}
	case D24_DC56:
		{
			return "D24.DC56" + "(" + _dafny.String(data.Cf90) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC57:
		{
			data2, ok := other.Get_().(D24_DC57)
			return ok && _dafny.AreEqual(data1.Cf91, data2.Cf91) && data1.Cf92 == data2.Cf92 && data1.Cf93 == data2.Cf93 && data1.Cf94.Cmp(data2.Cf94) == 0 && data1.Cf95 == data2.Cf95
		}
	case D24_DC56:
		{
			data2, ok := other.Get_().(D24_DC56)
			return ok && data1.Cf90 == data2.Cf90
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

type D25_DC59 struct {
	Cf97 bool
}

func (D25_DC59) isD25() {}

func (CompanionStruct_D25_) Create_DC59_(Cf97 bool) D25 {
	return D25{D25_DC59{Cf97}}
}

func (_this D25) Is_DC59() bool {
	_, ok := _this.Get_().(D25_DC59)
	return ok
}

type D25_DC58 struct {
	Cf96 _dafny.Array
}

func (D25_DC58) isD25() {}

func (CompanionStruct_D25_) Create_DC58_(Cf96 _dafny.Array) D25 {
	return D25{D25_DC58{Cf96}}
}

func (_this D25) Is_DC58() bool {
	_, ok := _this.Get_().(D25_DC58)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC59_(false)
}

func (_this D25) Dtor_cf97() bool {
	return _this.Get_().(D25_DC59).Cf97
}

func (_this D25) Dtor_cf96() _dafny.Array {
	return _this.Get_().(D25_DC58).Cf96
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC59:
		{
			return "D25.DC59" + "(" + _dafny.String(data.Cf97) + ")"
		}
	case D25_DC58:
		{
			return "D25.DC58" + "(" + _dafny.String(data.Cf96) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC59:
		{
			data2, ok := other.Get_().(D25_DC59)
			return ok && data1.Cf97 == data2.Cf97
		}
	case D25_DC58:
		{
			data2, ok := other.Get_().(D25_DC58)
			return ok && data1.Cf96 == data2.Cf96
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

type D26_DC60 struct {
	Cf98 _dafny.MultiSet
}

func (D26_DC60) isD26() {}

func (CompanionStruct_D26_) Create_DC60_(Cf98 _dafny.MultiSet) D26 {
	return D26{D26_DC60{Cf98}}
}

func (_this D26) Is_DC60() bool {
	_, ok := _this.Get_().(D26_DC60)
	return ok
}

func (CompanionStruct_D26_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D26) Dtor_cf98() _dafny.MultiSet {
	return _this.Get_().(D26_DC60).Cf98
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC60:
		{
			return "D26.DC60" + "(" + _dafny.String(data.Cf98) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC60:
		{
			data2, ok := other.Get_().(D26_DC60)
			return ok && data1.Cf98.Equals(data2.Cf98)
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

type D27_DC62 struct {
	Cf100 _dafny.Sequence
	Cf101 _dafny.CodePoint
	Cf102 bool
	Cf103 _dafny.Sequence
}

func (D27_DC62) isD27() {}

func (CompanionStruct_D27_) Create_DC62_(Cf100 _dafny.Sequence, Cf101 _dafny.CodePoint, Cf102 bool, Cf103 _dafny.Sequence) D27 {
	return D27{D27_DC62{Cf100, Cf101, Cf102, Cf103}}
}

func (_this D27) Is_DC62() bool {
	_, ok := _this.Get_().(D27_DC62)
	return ok
}

type D27_DC63 struct {
	Cf104 _dafny.Int
	Cf105 _dafny.Int
	Cf106 _dafny.Int
	Cf107 bool
}

func (D27_DC63) isD27() {}

func (CompanionStruct_D27_) Create_DC63_(Cf104 _dafny.Int, Cf105 _dafny.Int, Cf106 _dafny.Int, Cf107 bool) D27 {
	return D27{D27_DC63{Cf104, Cf105, Cf106, Cf107}}
}

func (_this D27) Is_DC63() bool {
	_, ok := _this.Get_().(D27_DC63)
	return ok
}

type D27_DC61 struct {
	Cf99 _dafny.Map
}

func (D27_DC61) isD27() {}

func (CompanionStruct_D27_) Create_DC61_(Cf99 _dafny.Map) D27 {
	return D27{D27_DC61{Cf99}}
}

func (_this D27) Is_DC61() bool {
	_, ok := _this.Get_().(D27_DC61)
	return ok
}

type D27_DC64 struct {
	Cf108 D27
}

func (D27_DC64) isD27() {}

func (CompanionStruct_D27_) Create_DC64_(Cf108 D27) D27 {
	return D27{D27_DC64{Cf108}}
}

func (_this D27) Is_DC64() bool {
	_, ok := _this.Get_().(D27_DC64)
	return ok
}

func (CompanionStruct_D27_) Default() D27 {
	return Companion_D27_.Create_DC62_(_dafny.EmptySeq, _dafny.CodePoint('D'), false, _dafny.EmptySeq)
}

func (_this D27) Dtor_cf100() _dafny.Sequence {
	return _this.Get_().(D27_DC62).Cf100
}

func (_this D27) Dtor_cf101() _dafny.CodePoint {
	return _this.Get_().(D27_DC62).Cf101
}

func (_this D27) Dtor_cf102() bool {
	return _this.Get_().(D27_DC62).Cf102
}

func (_this D27) Dtor_cf103() _dafny.Sequence {
	return _this.Get_().(D27_DC62).Cf103
}

func (_this D27) Dtor_cf104() _dafny.Int {
	return _this.Get_().(D27_DC63).Cf104
}

func (_this D27) Dtor_cf105() _dafny.Int {
	return _this.Get_().(D27_DC63).Cf105
}

func (_this D27) Dtor_cf106() _dafny.Int {
	return _this.Get_().(D27_DC63).Cf106
}

func (_this D27) Dtor_cf107() bool {
	return _this.Get_().(D27_DC63).Cf107
}

func (_this D27) Dtor_cf99() _dafny.Map {
	return _this.Get_().(D27_DC61).Cf99
}

func (_this D27) Dtor_cf108() D27 {
	return _this.Get_().(D27_DC64).Cf108
}

func (_this D27) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D27_DC62:
		{
			return "D27.DC62" + "(" + data.Cf100.VerbatimString(true) + ", " + _dafny.String(data.Cf101) + ", " + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D27_DC63:
		{
			return "D27.DC63" + "(" + _dafny.String(data.Cf104) + ", " + _dafny.String(data.Cf105) + ", " + _dafny.String(data.Cf106) + ", " + _dafny.String(data.Cf107) + ")"
		}
	case D27_DC61:
		{
			return "D27.DC61" + "(" + _dafny.String(data.Cf99) + ")"
		}
	case D27_DC64:
		{
			return "D27.DC64" + "(" + _dafny.String(data.Cf108) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D27) Equals(other D27) bool {
	switch data1 := _this.Get_().(type) {
	case D27_DC62:
		{
			data2, ok := other.Get_().(D27_DC62)
			return ok && data1.Cf100.Equals(data2.Cf100) && data1.Cf101 == data2.Cf101 && data1.Cf102 == data2.Cf102 && data1.Cf103.Equals(data2.Cf103)
		}
	case D27_DC63:
		{
			data2, ok := other.Get_().(D27_DC63)
			return ok && data1.Cf104.Cmp(data2.Cf104) == 0 && data1.Cf105.Cmp(data2.Cf105) == 0 && data1.Cf106.Cmp(data2.Cf106) == 0 && data1.Cf107 == data2.Cf107
		}
	case D27_DC61:
		{
			data2, ok := other.Get_().(D27_DC61)
			return ok && data1.Cf99.Equals(data2.Cf99)
		}
	case D27_DC64:
		{
			data2, ok := other.Get_().(D27_DC64)
			return ok && data1.Cf108.Equals(data2.Cf108)
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

type D28_DC65 struct {
	Cf109 *C8
}

func (D28_DC65) isD28() {}

func (CompanionStruct_D28_) Create_DC65_(Cf109 *C8) D28 {
	return D28{D28_DC65{Cf109}}
}

func (_this D28) Is_DC65() bool {
	_, ok := _this.Get_().(D28_DC65)
	return ok
}

func (CompanionStruct_D28_) Default() *C8 {
	return (*C8)(nil)
}

func (_this D28) Dtor_cf109() *C8 {
	return _this.Get_().(D28_DC65).Cf109
}

func (_this D28) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D28_DC65:
		{
			return "D28.DC65" + "(" + _dafny.String(data.Cf109) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D28) Equals(other D28) bool {
	switch data1 := _this.Get_().(type) {
	case D28_DC65:
		{
			data2, ok := other.Get_().(D28_DC65)
			return ok && data1.Cf109 == data2.Cf109
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

// Definition of datatype D29
type D29 struct {
	Data_D29_
}

func (_this D29) Get_() Data_D29_ {
	return _this.Data_D29_
}

type Data_D29_ interface {
	isD29()
}

type CompanionStruct_D29_ struct {
}

var Companion_D29_ = CompanionStruct_D29_{}

type D29_DC66 struct {
	Cf110 *C3
}

func (D29_DC66) isD29() {}

func (CompanionStruct_D29_) Create_DC66_(Cf110 *C3) D29 {
	return D29{D29_DC66{Cf110}}
}

func (_this D29) Is_DC66() bool {
	_, ok := _this.Get_().(D29_DC66)
	return ok
}

func (CompanionStruct_D29_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D29) Dtor_cf110() *C3 {
	return _this.Get_().(D29_DC66).Cf110
}

func (_this D29) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D29_DC66:
		{
			return "D29.DC66" + "(" + _dafny.String(data.Cf110) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D29) Equals(other D29) bool {
	switch data1 := _this.Get_().(type) {
	case D29_DC66:
		{
			data2, ok := other.Get_().(D29_DC66)
			return ok && data1.Cf110 == data2.Cf110
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D29) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D29)
	return ok && _this.Equals(typed)
}

func Type_D29_() _dafny.TypeDescriptor {
	return type_D29_{}
}

type type_D29_ struct {
}

func (_this type_D29_) Default() interface{} {
	return Companion_D29_.Default()
}

func (_this type_D29_) String() string {
	return "main.D29"
}
func (_this D29) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D29{}

// End of datatype D29

// Definition of datatype D30
type D30 struct {
	Data_D30_
}

func (_this D30) Get_() Data_D30_ {
	return _this.Data_D30_
}

type Data_D30_ interface {
	isD30()
}

type CompanionStruct_D30_ struct {
}

var Companion_D30_ = CompanionStruct_D30_{}

type D30_DC67 struct {
	Cf111 _dafny.Map
}

func (D30_DC67) isD30() {}

func (CompanionStruct_D30_) Create_DC67_(Cf111 _dafny.Map) D30 {
	return D30{D30_DC67{Cf111}}
}

func (_this D30) Is_DC67() bool {
	_, ok := _this.Get_().(D30_DC67)
	return ok
}

func (CompanionStruct_D30_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D30) Dtor_cf111() _dafny.Map {
	return _this.Get_().(D30_DC67).Cf111
}

func (_this D30) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D30_DC67:
		{
			return "D30.DC67" + "(" + _dafny.String(data.Cf111) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D30) Equals(other D30) bool {
	switch data1 := _this.Get_().(type) {
	case D30_DC67:
		{
			data2, ok := other.Get_().(D30_DC67)
			return ok && data1.Cf111.Equals(data2.Cf111)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D30) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D30)
	return ok && _this.Equals(typed)
}

func Type_D30_() _dafny.TypeDescriptor {
	return type_D30_{}
}

type type_D30_ struct {
}

func (_this type_D30_) Default() interface{} {
	return Companion_D30_.Default()
}

func (_this type_D30_) String() string {
	return "main.D30"
}
func (_this D30) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D30{}

// End of datatype D30

// Definition of datatype D31
type D31 struct {
	Data_D31_
}

func (_this D31) Get_() Data_D31_ {
	return _this.Data_D31_
}

type Data_D31_ interface {
	isD31()
}

type CompanionStruct_D31_ struct {
}

var Companion_D31_ = CompanionStruct_D31_{}

type D31_DC69 struct {
	Cf113 _dafny.Int
}

func (D31_DC69) isD31() {}

func (CompanionStruct_D31_) Create_DC69_(Cf113 _dafny.Int) D31 {
	return D31{D31_DC69{Cf113}}
}

func (_this D31) Is_DC69() bool {
	_, ok := _this.Get_().(D31_DC69)
	return ok
}

type D31_DC70 struct {
	Cf114 _dafny.Int
}

func (D31_DC70) isD31() {}

func (CompanionStruct_D31_) Create_DC70_(Cf114 _dafny.Int) D31 {
	return D31{D31_DC70{Cf114}}
}

func (_this D31) Is_DC70() bool {
	_, ok := _this.Get_().(D31_DC70)
	return ok
}

type D31_DC68 struct {
	Cf112 _dafny.Map
}

func (D31_DC68) isD31() {}

func (CompanionStruct_D31_) Create_DC68_(Cf112 _dafny.Map) D31 {
	return D31{D31_DC68{Cf112}}
}

func (_this D31) Is_DC68() bool {
	_, ok := _this.Get_().(D31_DC68)
	return ok
}

func (CompanionStruct_D31_) Default() D31 {
	return Companion_D31_.Create_DC69_(_dafny.Zero)
}

func (_this D31) Dtor_cf113() _dafny.Int {
	return _this.Get_().(D31_DC69).Cf113
}

func (_this D31) Dtor_cf114() _dafny.Int {
	return _this.Get_().(D31_DC70).Cf114
}

func (_this D31) Dtor_cf112() _dafny.Map {
	return _this.Get_().(D31_DC68).Cf112
}

func (_this D31) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D31_DC69:
		{
			return "D31.DC69" + "(" + _dafny.String(data.Cf113) + ")"
		}
	case D31_DC70:
		{
			return "D31.DC70" + "(" + _dafny.String(data.Cf114) + ")"
		}
	case D31_DC68:
		{
			return "D31.DC68" + "(" + _dafny.String(data.Cf112) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D31) Equals(other D31) bool {
	switch data1 := _this.Get_().(type) {
	case D31_DC69:
		{
			data2, ok := other.Get_().(D31_DC69)
			return ok && data1.Cf113.Cmp(data2.Cf113) == 0
		}
	case D31_DC70:
		{
			data2, ok := other.Get_().(D31_DC70)
			return ok && data1.Cf114.Cmp(data2.Cf114) == 0
		}
	case D31_DC68:
		{
			data2, ok := other.Get_().(D31_DC68)
			return ok && data1.Cf112.Equals(data2.Cf112)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D31) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D31)
	return ok && _this.Equals(typed)
}

func Type_D31_() _dafny.TypeDescriptor {
	return type_D31_{}
}

type type_D31_ struct {
}

func (_this type_D31_) Default() interface{} {
	return Companion_D31_.Default()
}

func (_this type_D31_) String() string {
	return "main.D31"
}
func (_this D31) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D31{}

// End of datatype D31

// Definition of trait T0
type T0 interface {
	String() string
	F23() _dafny.CodePoint
	F23_set_(value _dafny.CodePoint)
	Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool
	F24() bool
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
	F23() _dafny.CodePoint
	F23_set_(value _dafny.CodePoint)
	Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool
	F24() bool
	Fm5(globalState *GlobalState) D0
	Fm6(globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) (_dafny.Sequence, D0, bool)
	M2(globalState *GlobalState) (bool, _dafny.Int, bool)
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
	F23() _dafny.CodePoint
	F23_set_(value _dafny.CodePoint)
	Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool
	F24() bool
	Fm7(globalState *GlobalState) _dafny.Sequence
	Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map
	M3(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int, bool)
	F25() _dafny.Array
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

// Definition of trait T3
type T3 interface {
	String() string
	Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int
	M4(p0 bool, globalState *GlobalState) bool
	F27() bool
}
type CompanionStruct_T3_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T3_ = CompanionStruct_T3_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T3_) CastTo_(x interface{}) T3 {
	var t T3
	t, _ = x.(T3)
	return t
}

// End of trait T3

// Definition of class GlobalState
type GlobalState struct {
	F7   _dafny.Int
	F10  _dafny.Set
	F13  bool
	F22  _dafny.Array
	_f0  _dafny.Array
	_f1  _dafny.Int
	_f2  _dafny.Map
	_f3  bool
	_f4  _dafny.Array
	_f5  bool
	_f6  _dafny.Array
	_f8  _dafny.Sequence
	_f9  bool
	_f11 _dafny.Int
	_f12 _dafny.Sequence
	_f14 _dafny.CodePoint
	_f15 _dafny.Map
	_f16 _dafny.Int
	_f17 _dafny.Sequence
	_f18 _dafny.Sequence
	_f19 bool
	_f20 _dafny.Int
	_f21 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F7 = _dafny.Zero
	_this.F10 = _dafny.EmptySet
	_this.F13 = false
	_this.F22 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.EmptyMap
	_this._f3 = false
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f5 = false
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptySeq
	_this._f9 = false
	_this._f11 = _dafny.Zero
	_this._f12 = _dafny.EmptySeq
	_this._f14 = _dafny.CodePoint('D')
	_this._f15 = _dafny.EmptyMap
	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.EmptySeq
	_this._f18 = _dafny.EmptySeq
	_this._f19 = false
	_this._f20 = _dafny.Zero
	_this._f21 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Int, f2 _dafny.Map, f3 bool, f4 _dafny.Array, f5 bool, f6 _dafny.Array, f7 _dafny.Int, f8 _dafny.Sequence, f9 bool, f10 _dafny.Set, f11 _dafny.Int, f12 _dafny.Sequence, f13 bool, f14 _dafny.CodePoint, f15 _dafny.Map, f16 _dafny.Int, f17 _dafny.Sequence, f18 _dafny.Sequence, f19 bool, f20 _dafny.Int, f21 _dafny.Array, f22 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Map {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Array {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Sequence {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.CodePoint {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Map {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Sequence {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() _dafny.Int {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() _dafny.Array {
	{
		return _this._f21
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f23 _dafny.CodePoint
	_f24 bool
	F37  _dafny.Map
	_f38 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
	_this.F37 = _dafny.EmptyMap
	_this._f38 = false
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

func (_this *C0) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C0) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C0) F24() bool {
	return _this._f24
}
func (_this *C0) Ctor__(f37 _dafny.Map, f38 bool, f23 _dafny.CodePoint, f24 bool) {
	{
		(_this).F37 = f37
		(_this)._f38 = f38
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C0) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !(!(((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F24()))).Cardinality())).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), (Companion_D0_.Create_DC2_((_this).F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(624), true)).Cardinality(), (_this).F24())).Dtor_cf4())).Cardinality()).Minus((_dafny.MultiSetOf((_this).F24())).Cardinality())))) >= 0))
	}
}
func (_this *C0) Fm26(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Set, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24(), (_this).F24())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf((_this).F38())).Cardinality())).Cardinality())), (_dafny.IntOfInt64(403)).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bdnuosq")).Cardinality()))))
	}
}
func (_this *C0) F38() bool {
	{
		return _this._f38
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f27 bool
	_f39 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f27 = false
	_this._f39 = false
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_}
}

var _ T3 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F27() bool {
	return _this._f27
}
func (_this *C1) Ctor__(f39 bool, f27 bool) {
	{
		(_this)._f39 = f39
		(_this)._f27 = f27
	}
}
func (_this *C1) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("snfxoe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(957))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}(func(_459_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-692))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}(func(_460_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})))).Cardinality()).Times((_dafny.MultiSetOf((_this).F39(), (_this).F27(), (_this).F27())).Cardinality())), _dafny.IntOfInt64(934))
	}
}
func (_this *C1) Fm32(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("afadxteqa"), _dafny.IntOfInt64(915))).Cardinality(), (_this).F39())).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(155), _dafny.IntOfUint32((_dafny.SeqOf((_this).F39())).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(590), _dafny.IntOfUint32((_dafny.SeqOf((_this).F39())).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(671), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yahpsabt")).Cardinality())))))
	}
}
func (_this *C1) Fm33(globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(506), ((func() _dafny.Set {
			var _coll60 = _dafny.NewBuilder()
			_ = _coll60
			for _iter62 := _dafny.Iterate((_dafny.SeqOf(Companion_D5_.Create_DC14_((_this).F39()))).Elements()); ; {
				_compr_60, _ok62 := _iter62()
				if !_ok62 {
					break
				}
				var _461_v0 D5
				_461_v0 = interface{}(_compr_60).(D5)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D5_.Create_DC14_((_this).F39())), _461_v0) {
					_coll60.Add(_461_v0)
				}
			}
			return _coll60.ToSet()
		}()).Union(_dafny.SetOf(Companion_D5_.Create_DC14_((_this).F39()), Companion_D5_.Create_DC14_((_this).F27())))).Cardinality())
	}
}
func (_this *C1) M4(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _462_i0 _dafny.Int
		_ = _462_i0
		_462_i0 = _dafny.Zero
		{
			for true {
				{
					if (_462_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_462_i0 = (_462_i0).Plus(_dafny.One)
					var _463_v0 _dafny.CodePoint
					_ = _463_v0
					_463_v0 = _dafny.CodePoint('h')
					_463_v0 = _463_v0
					if (_this).F39() {
						var _464_v1 _dafny.Int
						_ = _464_v1
						_464_v1 = _dafny.IntOfInt64(852)
						var _465_v2 _dafny.Sequence
						_ = _465_v2
						_465_v2 = _dafny.SeqOf(_464_v1, _464_v1, _464_v1)
						var _466_v3 _dafny.MultiSet
						_ = _466_v3
						_466_v3 = _dafny.MultiSetOf(_464_v1, _464_v1, _dafny.IntOfUint32((_465_v2).Cardinality()))
						var _467_v4 _dafny.Map
						_ = _467_v4
						_467_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F39(), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("dr"), (Companion_Default___.SafeIndex(_464_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dr")).Cardinality()))).Uint32(), _dafny.CodePoint('h')))
						var _468_v5 D2
						_ = _468_v5
						_468_v5 = Companion_D2_.Create_DC6_(_467_v4)
						var _469_v6 _dafny.Map
						_ = _469_v6
						_469_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v5, p0)
						var _470_v7 _dafny.Map
						_ = _470_v7
						_470_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v1, true)
						var _471_v8 _dafny.Sequence
						_ = _471_v8
						_471_v8 = _dafny.UnicodeSeqOfUtf8Bytes("nubq")
						var _472_v9 _dafny.Map
						_ = _472_v9
						_472_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F39(), p0)
						var _473_v10 _dafny.Array
						_ = _473_v10
						var _nwElement0_12 _dafny.Int = _464_v1
						_ = _nwElement0_12
						var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(18))
						_ = _nw64
						(_nw64).ArraySet1(_nwElement0_12, 0)
						(_nw64).ArraySet1(_464_v1, 1)
						(_nw64).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_465_v2).Cardinality()), (_this).Fm33(globalState)), 2)
						(_nw64).ArraySet1((_466_v3).Cardinality(), 3)
						(_nw64).ArraySet1((_469_v6).Cardinality(), 4)
						(_nw64).ArraySet1((_dafny.Zero).Minus((_dafny.MultiSetOf((func() bool {
							if (_470_v7).Contains(_dafny.IntOfInt64(483)) {
								return (_470_v7).Get(_dafny.IntOfInt64(483)).(bool)
							}
							return p0
						})())).Cardinality()), 5)
						(_nw64).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jovube")).Cardinality()), 6)
						(_nw64).ArraySet1((_this).Fm9(_465_v2, _464_v1, (_this).F27(), globalState), 7)
						(_nw64).ArraySet1((_this).Fm33(globalState), 8)
						(_nw64).ArraySet1(_464_v1, 9)
						(_nw64).ArraySet1(_464_v1, 10)
						(_nw64).ArraySet1(_464_v1, 11)
						(_nw64).ArraySet1(_464_v1, 12)
						(_nw64).ArraySet1(_dafny.IntOfInt64(-64), 13)
						(_nw64).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_471_v8, _471_v8)).Cardinality()), 14)
						(_nw64).ArraySet1((func() _dafny.Int {
							if (_this).F39() {
								return _464_v1
							}
							return _464_v1
						})(), 15)
						(_nw64).ArraySet1((_466_v3).Cardinality(), 16)
						(_nw64).ArraySet1((_472_v9).Cardinality(), 17)
						_473_v10 = _nw64
						(globalState).F22 = _473_v10
						_463_v0 = Companion_Default___.Fm34((_464_v1).Minus(_464_v1), p0, globalState)
						(globalState).F7 = _464_v1
						var _474_v11 _dafny.Set
						_ = _474_v11
						_474_v11 = _dafny.SetOf(_464_v1)
						var _475_v12 _dafny.Sequence
						_ = _475_v12
						_475_v12 = _dafny.SeqOf(_474_v11)
						var _476_v13 _dafny.MultiSet
						_ = _476_v13
						_476_v13 = _dafny.MultiSetOf(_dafny.SetOf(_464_v1), (_475_v12).Select((Companion_Default___.SafeIndex(_464_v1, _dafny.IntOfUint32((_475_v12).Cardinality()))).Uint32()).(_dafny.Set), _dafny.SetOf(_464_v1), _474_v11, Companion_Default___.Fm35((_this).F27(), globalState))
						_465_v2 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hrwmpm")).Cardinality()), _464_v1), _465_v2), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).Fm9(_465_v2, _464_v1, (_this).F27(), globalState), _464_v1, (_470_v7).Cardinality()), _dafny.SeqOf((_476_v13).Cardinality(), _464_v1))), (Companion_Default___.SafeIndex((_464_v1).Minus((_this).Fm9(_465_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ulr"), (Companion_Default___.SafeIndex(_464_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ulr")).Cardinality()))).Uint32(), _463_v0)).Cardinality()), p0, globalState)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hrwmpm")).Cardinality()), _464_v1), _465_v2), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).Fm9(_465_v2, _464_v1, (_this).F27(), globalState), _464_v1, (_470_v7).Cardinality()), _dafny.SeqOf((_476_v13).Cardinality(), _464_v1)))).Cardinality()))).Uint32(), _dafny.IntOfUint32((_465_v2).Cardinality()))
						_464_v1 = (_dafny.Zero).Minus(_464_v1)
					} else {
						var _477_v14 _dafny.Int
						_ = _477_v14
						_477_v14 = _dafny.IntOfInt64(699)
						var _478_v15 _dafny.Sequence
						_ = _478_v15
						_478_v15 = _dafny.SeqOf(_477_v14)
						var _479_v16 _dafny.Sequence
						_ = _479_v16
						_479_v16 = _dafny.SeqOf(_478_v15, _478_v15)
						var _480_v17 _dafny.Array
						_ = _480_v17
						var _len0_8 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_8
						var _nw65 _dafny.Array
						_ = _nw65
						if _len0_8.Cmp(_dafny.Zero) == 0 {
							_nw65 = _dafny.NewArray(_len0_8)
						} else {
							var _init8 func(_dafny.Int) bool = func(_481_i1 _dafny.Int) bool {
								return (_this).F39()
							}
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
						_480_v17 = _nw65
						var _482_v18 D1
						_ = _482_v18
						_482_v18 = Companion_D1_.Create_DC5_(_477_v14, _477_v14, _479_v16, _480_v17, _477_v14)
						var _rhs67 _dafny.Int = (_482_v18).Dtor_cf8()
						_ = _rhs67
						var _rhs68 bool = !((_this).F27()) || (p0)
						_ = _rhs68
						var _rhs69 bool = (_this).F39()
						_ = _rhs69
						var _lhs50 *GlobalState = globalState
						_ = _lhs50
						_lhs50.F7 = _rhs67
						r0 = _rhs68
						r0 = _rhs69
						(globalState).F13 = p0
						var _483_v19 _dafny.Sequence
						_ = _483_v19
						_483_v19 = _dafny.SeqOf(p0)
						_477_v14 = _dafny.IntOfUint32((_483_v19).Cardinality())
						(globalState).F13 = (_this).F39()
						var _484_v20 _dafny.Array
						_ = _484_v20
						var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
						_ = _nw66
						_484_v20 = _nw66
						var _485_v21 _dafny.Array
						_ = _485_v21
						var _nwElement0_13 _dafny.Array = _484_v20
						_ = _nwElement0_13
						var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(9))
						_ = _nw67
						(_nw67).ArraySet1(_nwElement0_13, 0)
						(_nw67).ArraySet1(_484_v20, 1)
						(_nw67).ArraySet1(_484_v20, 2)
						(_nw67).ArraySet1(_484_v20, 3)
						(_nw67).ArraySet1(_484_v20, 4)
						(_nw67).ArraySet1(_484_v20, 5)
						(_nw67).ArraySet1(_484_v20, 6)
						(_nw67).ArraySet1(_484_v20, 7)
						(_nw67).ArraySet1(_484_v20, 8)
						_485_v21 = _nw67
						var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_485_v21), 0))
						_ = _index57
						(_485_v21).ArraySet1(_484_v20, (_index57).Int())
						var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_485_v21), 0))
						_ = _index58
						(_485_v21).ArraySet1(_484_v20, (_index58).Int())
					}
					var _486_v22 _dafny.Int
					_ = _486_v22
					_486_v22 = _dafny.IntOfInt64(683)
					(globalState).F7 = (_dafny.Zero).Minus(_486_v22)
					var _487_v23 _dafny.Map
					_ = _487_v23
					_487_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_486_v22, Companion_Default___.Fm36(globalState))
					var _488_v24 _dafny.Array
					_ = _488_v24
					var _len0_9 _dafny.Int = _dafny.IntOfInt64(27)
					_ = _len0_9
					var _nw68 _dafny.Array
					_ = _nw68
					if _len0_9.Cmp(_dafny.Zero) == 0 {
						_nw68 = _dafny.NewArray(_len0_9)
					} else {
						var _init9 func(_dafny.Int) _dafny.Map = (func(_489_p0 bool) func(_dafny.Int) _dafny.Map {
							return func(_490_i2 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_489_p0, _489_p0)
							}
						})(p0)
						_ = _init9
						var _element0_9 = _init9(_dafny.Zero)
						_ = _element0_9
						_nw68 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
						(_nw68).ArraySet1(_element0_9, 0)
						var _nativeLen0_9 = (_len0_9).Int()
						_ = _nativeLen0_9
						for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
							(_nw68).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
						}
					}
					_488_v24 = _nw68
					var _491_v25 _dafny.Set
					_ = _491_v25
					_491_v25 = _dafny.SetOf(_486_v22)
					var _492_v26 _dafny.Sequence
					_ = _492_v26
					_492_v26 = _dafny.SeqOf(false)
					var _493_v27 _dafny.Map
					_ = _493_v27
					_493_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_488_v24, (_487_v23).Merge(((Companion_D6_.Create_DC15_(_487_v23)).Dtor_cf24()).Update((_491_v25).Cardinality(), _492_v26)))
					_487_v23 = (func() _dafny.Map {
						if (_493_v27).Contains(_488_v24) {
							return (_493_v27).Get(_488_v24).(_dafny.Map)
						}
						return Companion_Default___.Fm37(_486_v22, p0, globalState)
					})()
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _494_v28 _dafny.MultiSet
		_ = _494_v28
		_494_v28 = _dafny.MultiSetOf((_this).F27())
		var _495_v29 _dafny.Int
		_ = _495_v29
		_495_v29 = _dafny.IntOfInt64(-13)
		var _496_v30 _dafny.Set
		_ = _496_v30
		_496_v30 = _dafny.SetOf(_dafny.IntOfInt64(-220), _495_v29)
		var _497_v31 _dafny.Map
		_ = _497_v31
		_497_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_v28, (_496_v30).IsProperSubsetOf(_dafny.SetOf((_496_v30).Cardinality(), (func() _dafny.Int {
			if (_494_v28).Contains((_this).F27()) {
				return (_494_v28).Multiplicity((_this).F27())
			}
			return _495_v29
		})())))
		_497_v31 = _497_v31
		var _498_v32 _dafny.Sequence
		_ = _498_v32
		_498_v32 = _dafny.UnicodeSeqOfUtf8Bytes("pkn")
		var _499_v33 *C0
		_ = _499_v33
		var _nw69 *C0 = New_C0_()
		_ = _nw69
		_nw69.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F39(), _498_v32)).Update(p0, _498_v32), p0, _dafny.CodePoint('m'), p0)
		_499_v33 = _nw69
		(globalState).F13 = p0
		var _500_v34 D5
		_ = _500_v34
		_500_v34 = Companion_D5_.Create_DC14_((_this).F27())
		_500_v34 = Companion_D5_.Create_DC14_((_this).F39())
		var _501_v35 _dafny.CodePoint
		_ = _501_v35
		_501_v35 = _dafny.CodePoint('u')
		var _502_v36 _dafny.Map
		_ = _502_v36
		_502_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v35, (_500_v34).Dtor_cf23())
		_502_v36 = ((_502_v36).Merge(_502_v36)).Merge(_502_v36)
		r0 = (_this).F27()
		return r0
	}
}
func (_this *C1) F39() bool {
	{
		return _this._f39
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f23 _dafny.CodePoint
	_f24 bool
	_f25 _dafny.Array
	_f36 _dafny.MultiSet
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
	_this._f25 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f36 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C2) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C2) F24() bool {
	return _this._f24
}
func (_this *C2) F25() _dafny.Array {
	return _this._f25
}
func (_this *C2) Ctor__(f36 _dafny.MultiSet, f25 _dafny.Array, f23 _dafny.CodePoint, f24 bool) {
	{
		(_this)._f36 = f36
		(_this)._f25 = f25
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C2) Fm7(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality())), ((_dafny.SetOf(!((_this).F24()), (_this).F24())).Union(_dafny.SetOf((_this).F24()))).Cardinality(), (_dafny.IntOfInt64(998)).Plus(_dafny.IntOfInt64(870)))
	}
}
func (_this *C2) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("ptgbrq"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("sxchetkux")))
	}
}
func (_this *C2) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C2) Fm24(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())))).Union(func() _dafny.Set {
			var _coll61 = _dafny.NewBuilder()
			_ = _coll61
			for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(584), _dafny.IntOfInt64(-980))); ; {
				_compr_61, _ok63 := _iter63()
				if !_ok63 {
					break
				}
				var _503_v0 _dafny.Int
				_503_v0 = interface{}(_compr_61).(_dafny.Int)
				if ((_dafny.IntOfInt64(584)).Cmp(_503_v0) <= 0) && ((_503_v0).Cmp(_dafny.IntOfInt64(-980)) < 0) {
					_coll61.Add(Companion_Default___.SafeModuloInt(_503_v0, _dafny.IntOfInt64(982)))
				}
			}
			return _coll61.ToSet()
		}())).Union((Companion_D5_.Create_DC12_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jayht")).Cardinality()))).Cardinality()))).Dtor_cf19())).Cardinality())
	}
}
func (_this *C2) Fm25(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf((_this).F24())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F24()), _dafny.SeqOf(false)))
	}
}
func (_this *C2) M3(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _504_v0 _dafny.Int
		_ = _504_v0
		_504_v0 = _dafny.IntOfInt64(268)
		var _505_v1 _dafny.Set
		_ = _505_v1
		_505_v1 = _dafny.SetOf(_504_v0, _504_v0)
		var _506_v2 D5
		_ = _506_v2
		_506_v2 = Companion_D5_.Create_DC12_(_505_v1)
		var _source9 D5 = _506_v2
		_ = _source9
		if _source9.Is_DC13() {
			var _507___mcc_h0 _dafny.Int = _source9.Get_().(D5_DC13).Cf20
			_ = _507___mcc_h0
			var _508___mcc_h1 bool = _source9.Get_().(D5_DC13).Cf21
			_ = _508___mcc_h1
			var _509___mcc_h2 bool = _source9.Get_().(D5_DC13).Cf22
			_ = _509___mcc_h2
			var _510_cf22 bool = _509___mcc_h2
			_ = _510_cf22
			var _511_cf21 bool = _508___mcc_h1
			_ = _511_cf21
			var _512_cf20 _dafny.Int = _507___mcc_h0
			_ = _512_cf20
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen(((_this).F25()), 0))
			_ = _index59
			((_this).F25()).ArraySet1CodePoint(p1, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen(((_this).F25()), 0))
			_ = _index60
			((_this).F25()).ArraySet1CodePoint(_dafny.CodePoint('d'), (_index60).Int())
			var _513_v3 D0
			_ = _513_v3
			_513_v3 = Companion_D0_.Create_DC1_(_504_v0)
			var _514_v4 _dafny.Map
			_ = _514_v4
			_514_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_513_v3), _513_v3), _504_v0)
			var _515_v5 _dafny.Sequence
			_ = _515_v5
			_515_v5 = _dafny.SeqOf(p0, (_this).F24())
			_514_v4 = (_514_v4).Update((func() bool {
				if (_515_v5).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0((_this).F24(), _504_v0, globalState), _dafny.IntOfUint32((_515_v5).Cardinality()))).Uint32()).(bool) {
					return _511_cf21
				}
				return _510_cf22
			})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_515_v5).Cardinality())))
			var _516_v6 _dafny.Map
			_ = _516_v6
			_516_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_512_cf20).Times(_504_v0), p0)
			_516_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_504_v0, (_510_cf22) == ((_this).F24()))
			var _517_v7 _dafny.Array
			_ = _517_v7
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw70
			_517_v7 = _nw70
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_517_v7), 0))
			_ = _index61
			(_517_v7).ArraySet1(false, (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_517_v7), 0))
			_ = _index62
			(_517_v7).ArraySet1(p0, (_index62).Int())
		} else if _source9.Is_DC14() {
			var _518___mcc_h3 bool = _source9.Get_().(D5_DC14).Cf23
			_ = _518___mcc_h3
			var _519_cf23 bool = _518___mcc_h3
			_ = _519_cf23
			(globalState).F7 = (_dafny.Zero).Minus((func() _dafny.Int {
				if _519_cf23 {
					return _504_v0
				}
				return _504_v0
			})())
			var _520_v8 _dafny.Sequence
			_ = _520_v8
			_520_v8 = _dafny.SeqOf((_this).F24())
			var _521_v9 _dafny.Array
			_ = _521_v9
			var _nwElement0_14 bool = p0
			_ = _nwElement0_14
			var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(12))
			_ = _nw71
			(_nw71).ArraySet1(_nwElement0_14, 0)
			(_nw71).ArraySet1(p0, 1)
			(_nw71).ArraySet1(_519_cf23, 2)
			(_nw71).ArraySet1(p0, 3)
			(_nw71).ArraySet1(_519_cf23, 4)
			(_nw71).ArraySet1(p0, 5)
			(_nw71).ArraySet1((_504_v0).Cmp((_dafny.Zero).Minus(_504_v0)) != 0, 6)
			(_nw71).ArraySet1((_this).F24(), 7)
			(_nw71).ArraySet1(Companion_Default___.Fm2(globalState), 8)
			(_nw71).ArraySet1((func() bool {
				if true {
					return (_this).Fm4(_504_v0, _dafny.Companion_Sequence_.Update(_520_v8, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-415), _dafny.IntOfUint32((_520_v8).Cardinality()))).Uint32(), p0), globalState)
				}
				return p0
			})(), 9)
			(_nw71).ArraySet1(p0, 10)
			(_nw71).ArraySet1((func() bool {
				if _519_cf23 {
					return (_this).F24()
				}
				return p0
			})(), 11)
			_521_v9 = _nw71
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_521_v9), 0))
			_ = _index63
			(_521_v9).ArraySet1(true, (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_521_v9), 0))
			_ = _index64
			(_521_v9).ArraySet1((_this).F24(), (_index64).Int())
			var _522_v10 _dafny.Sequence
			_ = _522_v10
			_522_v10 = _dafny.UnicodeSeqOfUtf8Bytes("yf")
			var _523_v11 _dafny.Sequence
			_ = _523_v11
			_523_v11 = _dafny.SeqOf(_504_v0)
			var _524_v12 _dafny.Sequence
			_ = _524_v12
			_524_v12 = _dafny.SeqOf(_523_v11)
			var _525_v13 D1
			_ = _525_v13
			_525_v13 = Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_522_v10).Cardinality()), _504_v0, _524_v12, _521_v9, _504_v0)
			var _pat_let_tv11 = _504_v0
			_ = _pat_let_tv11
			var _pat_let_tv12 = _521_v9
			_ = _pat_let_tv12
			var _pat_let_tv13 = _524_v12
			_ = _pat_let_tv13
			var _pat_let_tv14 = _504_v0
			_ = _pat_let_tv14
			var _pat_let_tv15 = _523_v11
			_ = _pat_let_tv15
			var _pat_let_tv16 = _504_v0
			_ = _pat_let_tv16
			var _pat_let_tv17 = _523_v11
			_ = _pat_let_tv17
			_521_v9 = (func(_pat_let12_0 D1) D1 {
				return func(_530_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let17_0 _dafny.Int) D1 {
						return func(_531_dt__update_hcf7_h1 _dafny.Int) D1 {
							return func(_pat_let18_0 _dafny.Int) D1 {
								return func(_534_dt__update_hcf8_h0 _dafny.Int) D1 {
									return Companion_D1_.Create_DC5_(_531_dt__update_hcf7_h1, _534_dt__update_hcf8_h0, (_530_dt__update__tmp_h1).Dtor_cf9(), (_530_dt__update__tmp_h1).Dtor_cf10(), (_530_dt__update__tmp_h1).Dtor_cf11())
								}(_pat_let18_0)
							}((_pat_let_tv15).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-263))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg58 _dafny.Int) interface{} {
									return coer58(arg58)
								}
							}((func(_532_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_533_i0 _dafny.Int) _dafny.Int {
									return _532_v0
								}
							})(_pat_let_tv16)))).Cardinality()), _dafny.IntOfUint32((_pat_let_tv17).Cardinality()))).Uint32()).(_dafny.Int))
						}(_pat_let17_0)
					}(_pat_let_tv14)
				}(_pat_let12_0)
			}(func(_pat_let13_0 D1) D1 {
				return func(_526_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let14_0 _dafny.Int) D1 {
						return func(_527_dt__update_hcf7_h0 _dafny.Int) D1 {
							return func(_pat_let15_0 _dafny.Array) D1 {
								return func(_528_dt__update_hcf10_h0 _dafny.Array) D1 {
									return func(_pat_let16_0 _dafny.Sequence) D1 {
										return func(_529_dt__update_hcf9_h0 _dafny.Sequence) D1 {
											return Companion_D1_.Create_DC5_(_527_dt__update_hcf7_h0, (_526_dt__update__tmp_h0).Dtor_cf8(), _529_dt__update_hcf9_h0, _528_dt__update_hcf10_h0, (_526_dt__update__tmp_h0).Dtor_cf11())
										}(_pat_let16_0)
									}(_pat_let_tv13)
								}(_pat_let15_0)
							}(_pat_let_tv12)
						}(_pat_let14_0)
					}(_pat_let_tv11)
				}(_pat_let13_0)
			}(_525_v13))).Dtor_cf10()
			_522_v10 = _dafny.Companion_Sequence_.Update(_522_v10, (Companion_Default___.SafeIndex(_504_v0, _dafny.IntOfUint32((_522_v10).Cardinality()))).Uint32(), _this.F23())
		} else {
			var _535___mcc_h4 _dafny.Set = _source9.Get_().(D5_DC12).Cf19
			_ = _535___mcc_h4
			var _536_cf19 _dafny.Set = _535___mcc_h4
			_ = _536_cf19
			(globalState).F13 = (_this).F24()
			var _537_v14 _dafny.Sequence
			_ = _537_v14
			_537_v14 = _dafny.SeqOf((_this).F24(), !(p0))
			var _538_v15 _dafny.Array
			_ = _538_v15
			var _nwElement0_15 bool = p0
			_ = _nwElement0_15
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(21))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_15, 0)
			(_nw72).ArraySet1(p0, 1)
			(_nw72).ArraySet1(Companion_Default___.Fm2(globalState), 2)
			(_nw72).ArraySet1((_this).F24(), 3)
			(_nw72).ArraySet1(p0, 4)
			(_nw72).ArraySet1(p0, 5)
			(_nw72).ArraySet1(p0, 6)
			(_nw72).ArraySet1((_537_v14).Select((Companion_Default___.SafeIndex(_504_v0, _dafny.IntOfUint32((_537_v14).Cardinality()))).Uint32()).(bool), 7)
			(_nw72).ArraySet1((_this).F24(), 8)
			(_nw72).ArraySet1(p0, 9)
			(_nw72).ArraySet1(!((_this).F24()), 10)
			(_nw72).ArraySet1(false, 11)
			(_nw72).ArraySet1((_this).F24(), 12)
			(_nw72).ArraySet1(false, 13)
			(_nw72).ArraySet1((_this).F24(), 14)
			(_nw72).ArraySet1(p0, 15)
			(_nw72).ArraySet1((_this).F24(), 16)
			(_nw72).ArraySet1((_this).F24(), 17)
			(_nw72).ArraySet1(!(p0), 18)
			(_nw72).ArraySet1((_this).F24(), 19)
			(_nw72).ArraySet1(!(false), 20)
			_538_v15 = _nw72
			var _539_v16 _dafny.Array
			_ = _539_v16
			var _nwElement0_16 _dafny.Array = _538_v15
			_ = _nwElement0_16
			var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.One)
			_ = _nw73
			(_nw73).ArraySet1(_nwElement0_16, 0)
			_539_v16 = _nw73
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))
			_ = _index65
			(_539_v16).ArraySet1(_538_v15, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))
			_ = _index66
			(_539_v16).ArraySet1(_538_v15, (_index66).Int())
			var _540_v17 *C0
			_ = _540_v17
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm1(p0, globalState)), false, p1, false)
			_540_v17 = _nw74
			var _541_v18 _dafny.Sequence
			_ = _541_v18
			_541_v18 = _dafny.UnicodeSeqOfUtf8Bytes("t")
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_539_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))).Int()))
			_ = _arr0
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_dafny.ArrayCastTo((_539_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))).Int()))), 0))
			_ = _index67
			(_arr0).ArraySet1((p0) && (p0), (_index67).Int())
			var _542_v19 _dafny.Set
			_ = _542_v19
			_542_v19 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_537_v14, (Companion_Default___.SafeIndex(_504_v0, _dafny.IntOfUint32((_537_v14).Cardinality()))).Uint32(), (_this).F24()))
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_539_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))).Int()))
			_ = _arr1
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_dafny.ArrayCastTo((_539_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))).Int()))), 0))
			_ = _index68
			var _rhs70 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("eeheeq")
			_ = _rhs70
			var _rhs71 bool = !(_dafny.SetOf(_537_v14)).Equals(_542_v19)
			_ = _rhs71
			var _lhs51 _dafny.Array = _dafny.ArrayCastTo((_539_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))).Int()))
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_dafny.ArrayCastTo((_539_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_539_v16), 0))).Int()))), 0))
			_ = _lhs52
			_541_v18 = _rhs70
			(_lhs51).ArraySet1(_rhs71, (_lhs52).Int())
		}
		var _543_v20 _dafny.Sequence
		_ = _543_v20
		_543_v20 = _dafny.UnicodeSeqOfUtf8Bytes("mcs")
		var _544_v21 _dafny.Map
		_ = _544_v21
		_544_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_504_v0, false)
		var _545_v22 _dafny.Sequence
		_ = _545_v22
		_545_v22 = _dafny.SeqOf(_dafny.IntOfUint32((_543_v20).Cardinality()), (_544_v21).Cardinality())
		var _546_v23 _dafny.Sequence
		_ = _546_v23
		_546_v23 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-728))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}((func(_547_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_548_i2 _dafny.Int) _dafny.Int {
				return (Companion_D4_.Create_DC11_(_547_v0)).Dtor_cf18()
			}
		})(_504_v0))), _545_v22)
		var _549_v24 _dafny.Sequence
		_ = _549_v24
		_549_v24 = _dafny.SeqOf((_this).F24())
		var _550_v25 _dafny.Array
		_ = _550_v25
		var _nwElement0_17 bool = p0
		_ = _nwElement0_17
		var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
		_ = _nw75
		(_nw75).ArraySet1(_nwElement0_17, 0)
		(_nw75).ArraySet1(!(false), 1)
		(_nw75).ArraySet1((_this).F24(), 2)
		(_nw75).ArraySet1(p0, 3)
		(_nw75).ArraySet1(true, 4)
		(_nw75).ArraySet1((_this).F24(), 5)
		(_nw75).ArraySet1((_this).F24(), 6)
		(_nw75).ArraySet1(Companion_Default___.Fm2(globalState), 7)
		(_nw75).ArraySet1((_this).F24(), 8)
		(_nw75).ArraySet1((_this).F24(), 9)
		(_nw75).ArraySet1(p0, 10)
		(_nw75).ArraySet1(p0, 11)
		(_nw75).ArraySet1(!((_this).F24()), 12)
		(_nw75).ArraySet1((_this).F24(), 13)
		(_nw75).ArraySet1(p0, 14)
		(_nw75).ArraySet1(p0, 15)
		(_nw75).ArraySet1(true, 16)
		(_nw75).ArraySet1(p0, 17)
		(_nw75).ArraySet1(p0, 18)
		(_nw75).ArraySet1((_this).Fm4(_dafny.IntOfInt64(-538), _549_v24, globalState), 19)
		(_nw75).ArraySet1(p0, 20)
		(_nw75).ArraySet1(p0, 21)
		(_nw75).ArraySet1(p0, 22)
		(_nw75).ArraySet1(p0, 23)
		(_nw75).ArraySet1(p0, 24)
		(_nw75).ArraySet1((_this).F24(), 25)
		(_nw75).ArraySet1((_this).F24(), 26)
		(_nw75).ArraySet1((_this).F24(), 27)
		(_nw75).ArraySet1(true, 28)
		_550_v25 = _nw75
		var _551_v26 D1
		_ = _551_v26
		_551_v26 = Companion_D1_.Create_DC5_(_504_v0, _504_v0, _546_v23, _550_v25, (_dafny.Zero).Minus((_505_v1).Cardinality()))
		var _552_v27 _dafny.MultiSet
		_ = _552_v27
		_552_v27 = _dafny.MultiSetOf((_551_v26).Dtor_cf10(), _550_v25, _550_v25, _550_v25)
		var _553_i1 _dafny.Int
		_ = _553_i1
		_553_i1 = _dafny.Zero
		{
			for !(((_552_v27).Union(_552_v27)).IsSubsetOf(_552_v27)) {
				{
					if (_553_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_553_i1 = (_553_i1).Plus(_dafny.One)
					var _554_v28 _dafny.Sequence
					_ = _554_v28
					_554_v28 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_543_v20, (Companion_Default___.SafeIndex(_504_v0, _dafny.IntOfUint32((_543_v20).Cardinality()))).Uint32(), _this.F23()), _543_v20)
					if _dafny.Companion_Sequence_.Equal((_554_v28).Select((Companion_Default___.SafeIndex(_504_v0, _dafny.IntOfUint32((_554_v28).Cardinality()))).Uint32()).(_dafny.Sequence), _543_v20) {
						(globalState).F13 = (_this).F24()
						_545_v22 = _545_v22
						var _555_v29 _dafny.Array
						_ = _555_v29
						var _len0_10 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_10
						var _nw76 _dafny.Array
						_ = _nw76
						if _len0_10.Cmp(_dafny.Zero) == 0 {
							_nw76 = _dafny.NewArray(_len0_10)
						} else {
							var _init10 func(_dafny.Int) _dafny.Int = (func(_556_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_557_i3 _dafny.Int) _dafny.Int {
									return (_557_i3).Times(_556_v0)
								}
							})(_504_v0)
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
						_555_v29 = _nw76
						var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_555_v29), 0))
						_ = _index69
						(_555_v29).ArraySet1(_504_v0, (_index69).Int())
						var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_555_v29), 0))
						_ = _index70
						(_555_v29).ArraySet1(_dafny.IntOfInt64(553), (_index70).Int())
						var _558_v30 _dafny.Map
						_ = _558_v30
						_558_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_555_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_555_v29), 0))).Int()).(_dafny.Int))
						var _559_v31 _dafny.Map
						_ = _559_v31
						_559_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_558_v30, p0)
						_559_v31 = (_559_v31).Update(_558_v30, (_504_v0).Cmp(_dafny.IntOfUint32((_545_v22).Cardinality())) != 0)
						r0 = ((_504_v0).Plus(_504_v0)).Cmp(((_555_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_555_v29), 0))).Int()).(_dafny.Int)).Minus((_this).Fm24(_dafny.IntOfInt64(975), (_this).F24(), globalState))) > 0
					} else {
						(globalState).F13 = (_504_v0).Cmp(_504_v0) < 0
						r1 = Companion_Default___.Fm0(_dafny.Companion_Sequence_.IsProperPrefixOf(_543_v20, _543_v20), _504_v0, globalState)
						var _560_v32 D0
						_ = _560_v32
						_560_v32 = Companion_D0_.Create_DC1_(_504_v0)
						_560_v32 = Companion_D0_.Create_DC1_((_504_v0).Times(_504_v0))
						var _561_v33 _dafny.Array
						_ = _561_v33
						var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
						_ = _nw77
						_561_v33 = _nw77
						var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_561_v33), 0))
						_ = _index71
						(_561_v33).ArraySet1((_504_v0).Plus(_dafny.IntOfInt64(604)), (_index71).Int())
						var _562_v34 _dafny.Map
						_ = _562_v34
						_562_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_561_v33, _dafny.Companion_Sequence_.Concatenate(_545_v22, _545_v22))
						var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_561_v33), 0))
						_ = _index72
						(_561_v33).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm1(p0, globalState)).Cardinality()), (_index72).Int())
						var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_561_v33), 0))
						_ = _index73
						var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_561_v33), 0))
						_ = _index74
						var _rhs72 _dafny.Int = (_504_v0).Plus((_504_v0).Minus(_504_v0))
						_ = _rhs72
						var _rhs73 bool = (_this).F24()
						_ = _rhs73
						var _rhs74 _dafny.Map = _562_v34
						_ = _rhs74
						var _rhs75 _dafny.Int = _dafny.IntOfUint32((_543_v20).Cardinality())
						_ = _rhs75
						var _rhs76 _dafny.Int = _504_v0
						_ = _rhs76
						var _lhs53 _dafny.Array = _561_v33
						_ = _lhs53
						var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_561_v33), 0))
						_ = _lhs54
						var _lhs55 *GlobalState = globalState
						_ = _lhs55
						var _lhs56 _dafny.Array = _561_v33
						_ = _lhs56
						var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_561_v33), 0))
						_ = _lhs57
						(_lhs53).ArraySet1(_rhs72, (_lhs54).Int())
						r0 = _rhs73
						_562_v34 = _rhs74
						_lhs55.F7 = _rhs75
						(_lhs56).ArraySet1(_rhs76, (_lhs57).Int())
						var _563_v35 _dafny.Map
						_ = _563_v35
						_563_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(389))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg60 _dafny.Int) interface{} {
								return coer60(arg60)
							}
						}(func(_564_i4 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						})), _dafny.IntOfInt64(472))
						_563_v35 = (_563_v35).Update(_dafny.UnicodeSeqOfUtf8Bytes("cjl"), (_561_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_561_v33), 0))).Int()).(_dafny.Int))
					}
					if !(!(p0)) {
						var _565_v36 _dafny.Map
						_ = _565_v36
						_565_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_550_v25, _543_v20)
						var _566_v37 D4
						_ = _566_v37
						_566_v37 = Companion_D4_.Create_DC10_((func() _dafny.Sequence {
							if (_565_v36).Contains(_550_v25) {
								return (_565_v36).Get(_550_v25).(_dafny.Sequence)
							}
							return _543_v20
						})())
						var _rhs77 _dafny.Int = (_504_v0).Minus(_504_v0)
						_ = _rhs77
						var _rhs78 D4 = _566_v37
						_ = _rhs78
						var _rhs79 bool = ((_dafny.Zero).Minus(_504_v0)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_543_v20, _543_v20)).Cardinality())) >= 0
						_ = _rhs79
						var _rhs80 _dafny.Int = (_dafny.Zero).Minus(_504_v0)
						_ = _rhs80
						var _rhs81 bool = p0
						_ = _rhs81
						var _lhs58 *GlobalState = globalState
						_ = _lhs58
						var _lhs59 *GlobalState = globalState
						_ = _lhs59
						var _lhs60 *GlobalState = globalState
						_ = _lhs60
						var _lhs61 *GlobalState = globalState
						_ = _lhs61
						_lhs58.F7 = _rhs77
						_566_v37 = _rhs78
						_lhs59.F13 = _rhs79
						_lhs60.F7 = _rhs80
						_lhs61.F13 = _rhs81
						(globalState).F7 = (_504_v0).Minus(_504_v0)
						(_this).F23_set_(_this.F23())
						r0 = (Companion_D2_.Create_DC7_(!((_this).F24()), p0)).Dtor_cf14()
						_543_v20 = _543_v20
					} else {
						(globalState).F13 = (_this).F24()
						var _567_v38 _dafny.Set
						_ = _567_v38
						_567_v38 = _dafny.SetOf(!((_this).F24()))
						var _568_v39 _dafny.Map
						_ = _568_v39
						_568_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
						var _569_v40 _dafny.Map
						_ = _569_v40
						_569_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), (_this).F24())).Merge(_568_v39)).Cardinality())
						var _570_v41 D0
						_ = _570_v41
						_570_v41 = Companion_D0_.Create_DC1_(_504_v0)
						var _571_v42 D0
						_ = _571_v42
						_571_v42 = Companion_D0_.Create_DC3_(_570_v41)
						var _572_v43 D0
						_ = _572_v43
						_572_v43 = Companion_D0_.Create_DC3_(_571_v42)
						var _573_v44 _dafny.Array
						_ = _573_v44
						var _nw78 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(9))
						_ = _nw78
						_573_v44 = _nw78
						var _574_v45 D4
						_ = _574_v45
						_574_v45 = Companion_D4_.Create_DC10_(_dafny.UnicodeSeqOfUtf8Bytes("jgil"))
						var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_573_v44), 0))
						_ = _index75
						(_573_v44).ArraySet1(_574_v45, (_index75).Int())
						var _575_v46 _dafny.Map
						_ = _575_v46
						_575_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
						var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_573_v44), 0))
						_ = _index76
						var _rhs82 _dafny.Set = _567_v38
						_ = _rhs82
						var _rhs83 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_504_v0).Cmp((_575_v46).Cardinality()) < 0, (_504_v0).Minus(_504_v0))
						_ = _rhs83
						var _rhs84 D0 = _572_v43
						_ = _rhs84
						var _rhs85 D4 = _574_v45
						_ = _rhs85
						var _lhs62 _dafny.Array = _573_v44
						_ = _lhs62
						var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_573_v44), 0))
						_ = _lhs63
						_567_v38 = _rhs82
						_569_v40 = _rhs83
						_572_v43 = _rhs84
						(_lhs62).ArraySet1(_rhs85, (_lhs63).Int())
						_504_v0 = _504_v0
						(globalState).F7 = _504_v0
						var _576_v47 _dafny.Array
						_ = _576_v47
						var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
						_ = _nw79
						_576_v47 = _nw79
						var _577_v48 _dafny.Array
						_ = _577_v48
						var _len0_11 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_11
						var _nw80 _dafny.Array
						_ = _nw80
						if _len0_11.Cmp(_dafny.Zero) == 0 {
							_nw80 = _dafny.NewArray(_len0_11)
						} else {
							var _init11 func(_dafny.Int) _dafny.Int = func(_578_i5 _dafny.Int) _dafny.Int {
								return (_578_i5).Plus(_dafny.IntOfInt64(291))
							}
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
						_577_v48 = _nw80
						var _579_v49 _dafny.Map
						_ = _579_v49
						_579_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_577_v48, _504_v0)
						var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_576_v47), 0))
						_ = _index77
						(_576_v47).ArraySet1(_579_v49, (_index77).Int())
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_576_v47), 0))
						_ = _index78
						var _rhs86 bool = p0
						_ = _rhs86
						var _rhs87 _dafny.Map = ((func() _dafny.Map {
							if Companion_Default___.Fm2(globalState) {
								return ((_579_v49).Update(_577_v48, (_575_v46).Cardinality())).Update(_577_v48, _504_v0)
							}
							return _579_v49
						})()).Merge(_579_v49)
						_ = _rhs87
						var _lhs64 _dafny.Array = _576_v47
						_ = _lhs64
						var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_576_v47), 0))
						_ = _lhs65
						r2 = _rhs86
						(_lhs64).ArraySet1(_rhs87, (_lhs65).Int())
					}
					var _580_v50 _dafny.Map
					_ = _580_v50
					_580_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
					var _581_v51 D0
					_ = _581_v51
					_581_v51 = Companion_D0_.Create_DC2_(p0, _504_v0, (func() bool {
						if (_580_v50).Contains(p0) {
							return (_580_v50).Get(p0).(bool)
						}
						return !(p0)
					})())
					var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_550_v25), 0))
					_ = _index79
					(_550_v25).ArraySet1((_581_v51).Dtor_cf4(), (_index79).Int())
					var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_550_v25), 0))
					_ = _index80
					(_550_v25).ArraySet1((_this).F24(), (_index80).Int())
					r2 = p0
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _582_v52 _dafny.Array
		_ = _582_v52
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_12
		var _nw81 _dafny.Array
		_ = _nw81
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw81 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.MultiSet = (func(_583_v0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
				return func(_584_i6 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_583_v0)
				}
			})(_504_v0)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw81 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw81).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw81).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_582_v52 = _nw81
		var _585_v53 _dafny.Map
		_ = _585_v53
		_585_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
		var _586_v54 _dafny.MultiSet
		_ = _586_v54
		_586_v54 = _dafny.MultiSetOf(_504_v0, (_585_v53).Cardinality())
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_582_v52), 0))
		_ = _index81
		(_582_v52).ArraySet1((_dafny.MultiSetOf(_504_v0)).Intersection(_586_v54), (_index81).Int())
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_582_v52), 0))
		_ = _index82
		(_582_v52).ArraySet1((_dafny.MultiSetFromSeq(_545_v22)).Intersection(_586_v54), (_index82).Int())
		var _hi2 _dafny.Int = (func() _dafny.Int {
			if (_this).F24() {
				return (_dafny.Zero).Minus(_504_v0)
			}
			return (_dafny.Zero).Minus(_504_v0)
		})()
		_ = _hi2
		for _587_i7 := _504_v0; _587_i7.Cmp(_hi2) < 0; _587_i7 = _587_i7.Plus(_dafny.One) {
			(globalState).F13 = (_505_v1).IsDisjointFrom(_505_v1)
			r0 = (_this).F24()
			var _588_v55 _dafny.Array
			_ = _588_v55
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
			_ = _nw82
			_588_v55 = _nw82
			var _589_v56 D5
			_ = _589_v56
			_589_v56 = Companion_D5_.Create_DC13_(_504_v0, (_this).F24(), p0)
			var _590_v57 _dafny.Set
			_ = _590_v57
			_590_v57 = _dafny.SetOf(_dafny.SeqOf(_589_v56, _589_v56))
			var _591_v58 _dafny.Map
			_ = _591_v58
			_591_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _590_v57)
			var _rhs88 bool = (Companion_Default___.Fm27((_this).F24(), p0, globalState)).Equals((func() _dafny.Set {
				if (_591_v58).Contains(p0) {
					return (_591_v58).Get(p0).(_dafny.Set)
				}
				return _590_v57
			})())
			_ = _rhs88
			var _rhs89 _dafny.CodePoint = _this.F23()
			_ = _rhs89
			var _rhs90 _dafny.Array = (func() _dafny.Array {
				if (_this).F24() {
					return _588_v55
				}
				return _588_v55
			})()
			_ = _rhs90
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			var _lhs67 *C2 = _this
			_ = _lhs67
			_lhs66.F13 = _rhs88
			_lhs67.F23_set_(_rhs89)
			_588_v55 = _rhs90
			var _592_v59 *C0
			_ = _592_v59
			var _nw83 *C0 = New_C0_()
			_ = _nw83
			_nw83.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.UnicodeSeqOfUtf8Bytes("owqe"))).Update((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("wdpkd")), (_this).F24(), _this.F23(), true)
			_592_v59 = _nw83
		}
		var _593_v60 _dafny.Array
		_ = _593_v60
		var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw84
		_593_v60 = _nw84
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_593_v60), 0))
		_ = _index83
		(_593_v60).ArraySet1(_504_v0, (_index83).Int())
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_593_v60), 0))
		_ = _index84
		(_593_v60).ArraySet1(Companion_Default___.SafeDivisionInt((_586_v54).Cardinality(), _504_v0), (_index84).Int())
		r1 = _dafny.IntOfUint32((_543_v20).Cardinality())
		r0 = (_dafny.MultiSetOf(false)).IsProperSubsetOf((_this).F36())
		r1 = _504_v0
		r2 = (_this).F24()
		return r0, r1, r2
	}
}
func (_this *C2) M12(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _594_v0 _dafny.Sequence
		_ = _594_v0
		_594_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ulpqjjbk")
		r1 = _dafny.Companion_Sequence_.Concatenate(_594_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}(func(_595_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		})))
		var _596_v1 _dafny.Sequence
		_ = _596_v1
		_596_v1 = _dafny.SeqOf((_this).F24())
		var _rhs91 _dafny.Int = (p1).Plus(p1)
		_ = _rhs91
		var _rhs92 _dafny.Int = _dafny.IntOfUint32((_596_v1).Cardinality())
		_ = _rhs92
		var _lhs68 *GlobalState = globalState
		_ = _lhs68
		var _lhs69 *GlobalState = globalState
		_ = _lhs69
		_lhs68.F7 = _rhs91
		_lhs69.F7 = _rhs92
		var _597_v2 D1
		_ = _597_v2
		_597_v2 = Companion_D1_.Create_DC4_((_594_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_594_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _source10 D1 = _597_v2
		_ = _source10
		if _source10.Is_DC5() {
			var _598___mcc_h0 _dafny.Int = _source10.Get_().(D1_DC5).Cf7
			_ = _598___mcc_h0
			var _599___mcc_h1 _dafny.Int = _source10.Get_().(D1_DC5).Cf8
			_ = _599___mcc_h1
			var _600___mcc_h2 _dafny.Sequence = _source10.Get_().(D1_DC5).Cf9
			_ = _600___mcc_h2
			var _601___mcc_h3 _dafny.Array = _source10.Get_().(D1_DC5).Cf10
			_ = _601___mcc_h3
			var _602___mcc_h4 _dafny.Int = _source10.Get_().(D1_DC5).Cf11
			_ = _602___mcc_h4
			var _603_cf11 _dafny.Int = _602___mcc_h4
			_ = _603_cf11
			var _604_cf10 _dafny.Array = _601___mcc_h3
			_ = _604_cf10
			var _605_cf9 _dafny.Sequence = _600___mcc_h2
			_ = _605_cf9
			var _606_cf8 _dafny.Int = _599___mcc_h1
			_ = _606_cf8
			var _607_cf7 _dafny.Int = _598___mcc_h0
			_ = _607_cf7
			var _608_v3 _dafny.Sequence
			_ = _608_v3
			_608_v3 = _dafny.SeqOf(_603_cf11)
			var _609_v4 _dafny.Set
			_ = _609_v4
			_609_v4 = _dafny.SetOf(_607_cf7, (_dafny.Zero).Minus(p1), (_608_v3).Select((Companion_Default___.SafeIndex(_603_cf11, _dafny.IntOfUint32((_608_v3).Cardinality()))).Uint32()).(_dafny.Int))
			var _rhs93 _dafny.Int = Companion_Default___.Fm0((_this).F24(), _607_cf7, globalState)
			_ = _rhs93
			var _rhs94 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_609_v4).Cardinality()), _dafny.IntOfUint32((_596_v1).Cardinality()))
			_ = _rhs94
			var _rhs95 bool = (p1).Cmp((_dafny.Zero).Minus(_606_cf8)) < 0
			_ = _rhs95
			var _lhs70 *GlobalState = globalState
			_ = _lhs70
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			_603_cf11 = _rhs93
			_lhs70.F7 = _rhs94
			_lhs71.F13 = _rhs95
			var _610_v5 D0
			_ = _610_v5
			_610_v5 = Companion_D0_.Create_DC2_((_this).F24(), _607_cf7, (_this).F24())
			var _611_v6 _dafny.Map
			_ = _611_v6
			_611_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _607_cf7)
			var _pat_let_tv18 = _611_v6
			_ = _pat_let_tv18
			_610_v5 = func(_pat_let19_0 D0) D0 {
				return func(_612_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let20_0 _dafny.Int) D0 {
						return func(_613_dt__update_hcf3_h0 _dafny.Int) D0 {
							return func(_pat_let21_0 bool) D0 {
								return func(_614_dt__update_hcf4_h0 bool) D0 {
									return Companion_D0_.Create_DC2_((_612_dt__update__tmp_h0).Dtor_cf2(), _613_dt__update_hcf3_h0, _614_dt__update_hcf4_h0)
								}(_pat_let21_0)
							}(true)
						}(_pat_let20_0)
					}((_pat_let_tv18).Cardinality())
				}(_pat_let19_0)
			}(_610_v5)
			var _615_v7 _dafny.Map
			_ = _615_v7
			_615_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _594_v0)
			var _616_v8 D2
			_ = _616_v8
			_616_v8 = Companion_D2_.Create_DC6_(_615_v7)
			_616_v8 = (func() D2 {
				if (_this).F24() {
					return _616_v8
				}
				return Companion_D2_.Create_DC6_(_615_v7)
			})()
			var _617_v9 _dafny.Map
			_ = _617_v9
			_617_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(-372))
			var _618_v10 _dafny.Set
			_ = _618_v10
			_618_v10 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _606_cf8), _617_v9, _617_v9, _617_v9)
			var _source11 D0 = Companion_Default___.Fm28(_594_v0, _608_v3, (_596_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_594_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_594_v0).Cardinality()), _dafny.IntOfUint32((_594_v0).Cardinality()))).Uint32(), _this.F23())).Cardinality()), _dafny.IntOfUint32((_596_v1).Cardinality()))).Uint32()).(bool), _618_v10, globalState)
			_ = _source11
			if _source11.Is_DC1() {
				var _619___mcc_h6 _dafny.Int = _source11.Get_().(D0_DC1).Cf1
				_ = _619___mcc_h6
				var _620_cf1 _dafny.Int = _619___mcc_h6
				_ = _620_cf1
				var _621_v11 *C0
				_ = _621_v11
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__(_615_v7, (Companion_D0_.Create_DC2_(true, _607_cf7, (_this).F24())).Dtor_cf4(), _this.F23(), (_this).F24())
				_621_v11 = _nw85
				var _622_v12 _dafny.Sequence
				_ = _622_v12
				_622_v12 = _dafny.SeqOf(_604_cf10)
				var _623_v13 _dafny.Map
				_ = _623_v13
				_623_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_621_v11).F38(), _606_cf8, globalState), (_622_v12).Select((Companion_Default___.SafeIndex((_609_v4).Cardinality(), _dafny.IntOfUint32((_622_v12).Cardinality()))).Uint32()).(_dafny.Array))
				_623_v13 = (_623_v13).Update((_608_v3).Select((Companion_Default___.SafeIndex(_603_cf11, _dafny.IntOfUint32((_608_v3).Cardinality()))).Uint32()).(_dafny.Int), _604_cf10)
				var _624_v14 _dafny.Map
				_ = _624_v14
				_624_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_594_v0).Cardinality()), _603_cf11)
				_624_v14 = (_624_v14).Update((_dafny.IntOfInt64(843)).Minus(p1), _dafny.IntOfInt64(474))
				var _625_v15 _dafny.Map
				_ = _625_v15
				_625_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_617_v9).Cardinality(), _596_v1)
				_625_v15 = (_625_v15).Update((_dafny.Zero).Minus(_606_cf8), _dafny.Companion_Sequence_.Concatenate(_596_v1, _596_v1))
			} else if _source11.Is_DC2() {
				var _626___mcc_h7 bool = _source11.Get_().(D0_DC2).Cf2
				_ = _626___mcc_h7
				var _627___mcc_h8 _dafny.Int = _source11.Get_().(D0_DC2).Cf3
				_ = _627___mcc_h8
				var _628___mcc_h9 bool = _source11.Get_().(D0_DC2).Cf4
				_ = _628___mcc_h9
				var _629_cf4 bool = _628___mcc_h9
				_ = _629_cf4
				var _630_cf3 _dafny.Int = _627___mcc_h8
				_ = _630_cf3
				var _631_cf2 bool = _626___mcc_h7
				_ = _631_cf2
				(globalState).F7 = Companion_Default___.SafeModuloInt((_607_cf7).Plus(_603_cf11), _607_cf7)
				(globalState).F7 = _630_cf3
				var _632_v16 _dafny.Map
				_ = _632_v16
				_632_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _603_cf11)
				var _pat_let_tv19 = p1
				_ = _pat_let_tv19
				var _pat_let_tv20 = _631_cf2
				_ = _pat_let_tv20
				var _rhs96 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (func(_pat_let22_0 D0) D0 {
						return func(_633_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let23_0 _dafny.Int) D0 {
								return func(_634_dt__update_hcf3_h1 _dafny.Int) D0 {
									return func(_pat_let24_0 bool) D0 {
										return func(_635_dt__update_hcf4_h1 bool) D0 {
											return Companion_D0_.Create_DC2_((_633_dt__update__tmp_h1).Dtor_cf2(), _634_dt__update_hcf3_h1, _635_dt__update_hcf4_h1)
										}(_pat_let24_0)
									}(_pat_let_tv20)
								}(_pat_let23_0)
							}(_pat_let_tv19)
						}(_pat_let22_0)
					}(_610_v5)).Dtor_cf4() {
						return _this.F23()
					}
					return Companion_Default___.Fm29(_630_cf3, _629_cf4, _631_cf2, _603_cf11, globalState)
				})()
				_ = _rhs96
				var _rhs97 _dafny.Map = (func() _dafny.Map {
					var _coll62 = _dafny.NewMapBuilder()
					_ = _coll62
					for _iter64 := _dafny.Iterate((_608_v3).Elements()); ; {
						_compr_62, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _636_v17 _dafny.Int
						_636_v17 = interface{}(_compr_62).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_608_v3, _636_v17) {
							_coll62.Add(Companion_Default___.SafeModuloInt(_636_v17, _630_cf3), _603_cf11)
						}
					}
					return _coll62.ToMap()
				}()).Merge((_632_v16).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_cf8, _603_cf11)))
				_ = _rhs97
				var _lhs72 *C2 = _this
				_ = _lhs72
				_lhs72.F23_set_(_rhs96)
				_632_v16 = _rhs97
				var _rhs98 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_594_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_594_v0, _594_v0), (Companion_Default___.SafeIndex(_606_cf8, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_594_v0, _594_v0)).Cardinality()))).Uint32(), _dafny.CodePoint('v')))
				_ = _rhs98
				var _rhs99 _dafny.CodePoint = _this.F23()
				_ = _rhs99
				var _lhs73 *C2 = _this
				_ = _lhs73
				r1 = _rhs98
				_lhs73.F23_set_(_rhs99)
			} else if _source11.Is_DC0() {
				var _637___mcc_h10 _dafny.MultiSet = _source11.Get_().(D0_DC0).Cf0
				_ = _637___mcc_h10
				var _638_cf0 _dafny.MultiSet = _637___mcc_h10
				_ = _638_cf0
				_594_v0 = _dafny.Companion_Sequence_.Update(_594_v0, (Companion_Default___.SafeIndex(_606_cf8, _dafny.IntOfUint32((_594_v0).Cardinality()))).Uint32(), _dafny.CodePoint('a'))
				var _639_v18 _dafny.Array
				_ = _639_v18
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_13
				var _nw86 _dafny.Array
				_ = _nw86
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw86 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Int = (func(_640_v2 D1) func(_dafny.Int) _dafny.Int {
						return func(_641_i1 _dafny.Int) _dafny.Int {
							return (_641_i1).Times((_dafny.SetOf((_640_v2).Dtor_cf6(), _this.F23())).Cardinality())
						}
					})(_597_v2)
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
				_639_v18 = _nw86
				var _rhs100 _dafny.Int = ((_dafny.Zero).Minus(_606_cf8)).Times((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_603_cf11, _603_cf11)))
				_ = _rhs100
				var _rhs101 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_596_v1, _596_v1), false)
				_ = _rhs101
				var _rhs102 _dafny.Int = (_607_cf7).Minus((_dafny.Zero).Minus(_607_cf7))
				_ = _rhs102
				var _rhs103 _dafny.Array = _639_v18
				_ = _rhs103
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				var _lhs76 *GlobalState = globalState
				_ = _lhs76
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				_lhs74.F7 = _rhs100
				_lhs75.F13 = _rhs101
				_lhs76.F7 = _rhs102
				_lhs77.F22 = _rhs103
				var _642_v19 D2
				_ = _642_v19
				_642_v19 = Companion_D2_.Create_DC7_((_this).F24(), (_this).F24())
				var _643_v20 *C0
				_ = _643_v20
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__(_615_v7, (_642_v19).Dtor_cf14(), _this.F23(), (_this).F24())
				_643_v20 = _nw87
				var _644_v21 _dafny.MultiSet
				_ = _644_v21
				_644_v21 = _dafny.MultiSetOf(_dafny.IntOfUint32((_594_v0).Cardinality()))
				var _645_v22 _dafny.Set
				_ = _645_v22
				_645_v22 = _dafny.SetOf((_644_v21).Difference(_644_v21))
				var _646_v23 _dafny.Sequence
				_ = _646_v23
				_646_v23 = _dafny.SeqOf(_594_v0, _594_v0)
				var _rhs104 _dafny.Set = Companion_Default___.Fm30((func() D2 {
					if (_this).F24() {
						return _616_v8
					}
					return _616_v8
				})(), _646_v23, globalState)
				_ = _rhs104
				var _rhs105 _dafny.Int = (_dafny.Zero).Minus(_607_cf7)
				_ = _rhs105
				_645_v22 = _rhs104
				r0 = _rhs105
			} else {
				var _647___mcc_h11 D0 = _source11.Get_().(D0_DC3).Cf5
				_ = _647___mcc_h11
				var _648_cf5 D0 = _647___mcc_h11
				_ = _648_cf5
				_603_cf11 = _607_cf7
				var _649_v24 _dafny.MultiSet
				_ = _649_v24
				_649_v24 = _dafny.MultiSetOf(_dafny.IntOfUint32((_596_v1).Cardinality()), _607_cf7, p1)
				var _650_v25 _dafny.Sequence
				_ = _650_v25
				_650_v25 = _dafny.SeqOf(_dafny.MultiSetOf(_607_cf7), _649_v24, _649_v24)
				var _651_v26 D5
				_ = _651_v26
				_651_v26 = Companion_D5_.Create_DC12_(_dafny.SetOf(p1))
				var _652_v27 _dafny.Array
				_ = _652_v27
				var _nwElement0_18 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ogv"), _594_v0)).Cardinality())
				_ = _nwElement0_18
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(15))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_18, 0)
				(_nw88).ArraySet1(_606_cf8, 1)
				(_nw88).ArraySet1((_609_v4).Cardinality(), 2)
				(_nw88).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mcybk")).Cardinality()), 3)
				(_nw88).ArraySet1((((_651_v26).Dtor_cf19()).Cardinality()).Plus((func() _dafny.Int {
					if (_649_v24).Contains(_607_cf7) {
						return (_649_v24).Multiplicity(_607_cf7)
					}
					return _dafny.IntOfUint32((_594_v0).Cardinality())
				})()), 4)
				(_nw88).ArraySet1(_dafny.IntOfInt64(276), 5)
				(_nw88).ArraySet1(_606_cf8, 6)
				(_nw88).ArraySet1((_dafny.IntOfUint32((_608_v3).Cardinality())).Plus((func() _dafny.Int {
					if (_617_v9).Contains((_596_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_608_v3).Cardinality()), _dafny.IntOfUint32((_596_v1).Cardinality()))).Uint32()).(bool)) {
						return (_617_v9).Get((_596_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_608_v3).Cardinality()), _dafny.IntOfUint32((_596_v1).Cardinality()))).Uint32()).(bool)).(_dafny.Int)
					}
					return p1
				})()), 7)
				(_nw88).ArraySet1(p1, 8)
				(_nw88).ArraySet1(_607_cf7, 9)
				(_nw88).ArraySet1(p1, 10)
				(_nw88).ArraySet1((_dafny.Zero).Minus((p1).Plus(p1)), 11)
				(_nw88).ArraySet1(Companion_Default___.Fm0(false, _606_cf8, globalState), 12)
				(_nw88).ArraySet1(_603_cf11, 13)
				(_nw88).ArraySet1(_dafny.IntOfUint32(((Companion_D4_.Create_DC10_(_594_v0)).Dtor_cf17()).Cardinality()), 14)
				_652_v27 = _nw88
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_652_v27), 0))
				_ = _index85
				(_652_v27).ArraySet1(p1, (_index85).Int())
				var _653_v28 D5
				_ = _653_v28
				_653_v28 = Companion_D5_.Create_DC13_(_603_cf11, true, (_this).F24())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_604_cf10), 0))
				_ = _index86
				(_604_cf10).ArraySet1((func(_pat_let25_0 D5) D5 {
					return func(_654_dt__update__tmp_h2 D5) D5 {
						return func(_pat_let26_0 bool) D5 {
							return func(_655_dt__update_hcf21_h0 bool) D5 {
								return Companion_D5_.Create_DC13_((_654_dt__update__tmp_h2).Dtor_cf20(), _655_dt__update_hcf21_h0, (_654_dt__update__tmp_h2).Dtor_cf22())
							}(_pat_let26_0)
						}((_this).F24())
					}(_pat_let25_0)
				}(_653_v28)).Dtor_cf21(), (_index86).Int())
				var _656_v29 _dafny.Map
				_ = _656_v29
				_656_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _604_cf10)
				var _657_v30 _dafny.Set
				_ = _657_v30
				_657_v30 = _dafny.SetOf((_this).F24())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_652_v27), 0))
				_ = _index87
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_604_cf10), 0))
				_ = _index88
				var _rhs106 _dafny.Sequence = Companion_Default___.Fm31((_this).F24(), !((_this).F24()), (_this).F24(), _610_v5, globalState)
				_ = _rhs106
				var _rhs107 _dafny.Int = (Companion_Default___.SafeDivisionInt(_607_cf7, _607_cf7)).Times(((_656_v29).Cardinality()).Minus(_603_cf11))
				_ = _rhs107
				var _rhs108 bool = (_657_v30).IsDisjointFrom((_657_v30).Intersection(_dafny.SetOf((_596_v1).Select((Companion_Default___.SafeIndex(_606_cf8, _dafny.IntOfUint32((_596_v1).Cardinality()))).Uint32()).(bool))))
				_ = _rhs108
				var _lhs78 _dafny.Array = _652_v27
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_652_v27), 0))
				_ = _lhs79
				var _lhs80 _dafny.Array = _604_cf10
				_ = _lhs80
				var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_604_cf10), 0))
				_ = _lhs81
				_650_v25 = _rhs106
				(_lhs78).ArraySet1(_rhs107, (_lhs79).Int())
				(_lhs80).ArraySet1(_rhs108, (_lhs81).Int())
				var _658_v31 _dafny.Array
				_ = _658_v31
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw89
				_658_v31 = _nw89
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_658_v31), 0))
				_ = _index89
				(_658_v31).ArraySet1(_594_v0, (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_658_v31), 0))
				_ = _index90
				(_658_v31).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qxxltvc"), (_index90).Int())
				var _659_v32 _dafny.Map
				_ = _659_v32
				_659_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_649_v24, _603_cf11)
				var _660_v33 _dafny.Map
				_ = _660_v33
				_660_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_603_cf11, _603_cf11)
				_659_v32 = (_659_v32).Update(_649_v24, Companion_Default___.Fm0(!((_604_cf10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_604_cf10), 0))).Int()).(bool)), (_660_v33).Cardinality(), globalState))
			}
		} else {
			var _661___mcc_h5 _dafny.CodePoint = _source10.Get_().(D1_DC4).Cf6
			_ = _661___mcc_h5
			var _662_cf6 _dafny.CodePoint = _661___mcc_h5
			_ = _662_cf6
			var _source12 D2 = Companion_D2_.Create_DC7_((_this).F24(), (_this).F24())
			_ = _source12
			if _source12.Is_DC7() {
				var _663___mcc_h12 bool = _source12.Get_().(D2_DC7).Cf13
				_ = _663___mcc_h12
				var _664___mcc_h13 bool = _source12.Get_().(D2_DC7).Cf14
				_ = _664___mcc_h13
				var _665_cf14 bool = _664___mcc_h13
				_ = _665_cf14
				var _666_cf13 bool = _663___mcc_h12
				_ = _666_cf13
				var _667_v34 D4
				_ = _667_v34
				_667_v34 = Companion_D4_.Create_DC11_(p1)
				var _pat_let_tv21 = p1
				_ = _pat_let_tv21
				var _pat_let_tv22 = p1
				_ = _pat_let_tv22
				_667_v34 = func(_pat_let27_0 D4) D4 {
					return func(_668_dt__update__tmp_h3 D4) D4 {
						return func(_pat_let28_0 _dafny.Int) D4 {
							return func(_669_dt__update_hcf18_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC11_(_669_dt__update_hcf18_h0)
							}(_pat_let28_0)
						}((_pat_let_tv21).Plus(_pat_let_tv22))
					}(_pat_let27_0)
				}(Companion_D4_.Create_DC11_(p1))
				var _670_v35 _dafny.Map
				_ = _670_v35
				_670_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p1)
				_670_v35 = (_670_v35).Update(!(!(true)) || (_665_cf14), p1)
				var _671_v36 _dafny.Sequence
				_ = _671_v36
				_671_v36 = _dafny.SeqOf(p1)
				var _672_v37 T3
				_ = _672_v37
				var _nw90 *C1 = New_C1_()
				_ = _nw90
				_nw90.Ctor__(true, _666_cf13)
				_672_v37 = _nw90
				var _673_v38 _dafny.Array
				_ = _673_v38
				var _nwElement0_19 T3 = _672_v37
				_ = _nwElement0_19
				var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(13))
				_ = _nw91
				(_nw91).ArraySet1(_nwElement0_19, 0)
				(_nw91).ArraySet1(_672_v37, 1)
				(_nw91).ArraySet1(_672_v37, 2)
				(_nw91).ArraySet1(_672_v37, 3)
				(_nw91).ArraySet1(_672_v37, 4)
				(_nw91).ArraySet1(_672_v37, 5)
				(_nw91).ArraySet1(_672_v37, 6)
				(_nw91).ArraySet1(_672_v37, 7)
				(_nw91).ArraySet1(_672_v37, 8)
				(_nw91).ArraySet1(_672_v37, 9)
				(_nw91).ArraySet1(_672_v37, 10)
				(_nw91).ArraySet1(_672_v37, 11)
				(_nw91).ArraySet1(_672_v37, 12)
				_673_v38 = _nw91
				var _674_v39 _dafny.Map
				_ = _674_v39
				_674_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_673_v38, _this.F23())
				var _rhs109 _dafny.Sequence = (func() _dafny.Sequence {
					if _665_cf14 {
						return _dafny.SeqOf(_dafny.IntOfInt64(-18))
					}
					return _dafny.Companion_Sequence_.Concatenate(_671_v36, _671_v36)
				})()
				_ = _rhs109
				var _rhs110 _dafny.Map = _674_v39
				_ = _rhs110
				_671_v36 = _rhs109
				_674_v39 = _rhs110
				var _675_v40 _dafny.Array
				_ = _675_v40
				var _nwElement0_20 _dafny.Sequence = _596_v1
				_ = _nwElement0_20
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(18))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_20, 0)
				(_nw92).ArraySet1((func() _dafny.Sequence {
					if _666_cf13 {
						return _596_v1
					}
					return _dafny.SeqOf((_672_v37).F27())
				})(), 1)
				(_nw92).ArraySet1(_596_v1, 2)
				(_nw92).ArraySet1(_596_v1, 3)
				(_nw92).ArraySet1(_596_v1, 4)
				(_nw92).ArraySet1(_596_v1, 5)
				(_nw92).ArraySet1(_596_v1, 6)
				(_nw92).ArraySet1(_596_v1, 7)
				(_nw92).ArraySet1(_596_v1, 8)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_672_v37).F27(), (_this).F24()), _dafny.SeqOf(_665_cf14)), 9)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F24()), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality()))).Uint32(), (_672_v37).F27()), 10)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm2(globalState)), _596_v1), 11)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, _666_cf13), _596_v1), 12)
				(_nw92).ArraySet1(_596_v1, 13)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_596_v1, _596_v1), 14)
				(_nw92).ArraySet1(Companion_Default___.Fm36(globalState), 15)
				(_nw92).ArraySet1(_596_v1, 16)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_596_v1, _dafny.SeqOf((_672_v37).F27())), 17)
				_675_v40 = _nw92
				var _676_v41 _dafny.Map
				_ = _676_v41
				_676_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(440), _596_v1)
				var _677_v42 _dafny.Map
				_ = _677_v42
				_677_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_672_v37).F27())
				var _678_v43 D6
				_ = _678_v43
				_678_v43 = Companion_D6_.Create_DC16_(_677_v42, (_672_v37).F27(), (_672_v37).F27())
				var _679_v44 _dafny.Set
				_ = _679_v44
				_679_v44 = _dafny.SetOf((_this).F25())
				var _680_v45 _dafny.Array
				_ = _680_v45
				var _nwElement0_21 bool = _665_cf14
				_ = _nwElement0_21
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(27))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_21, 0)
				(_nw93).ArraySet1(_dafny.Companion_Sequence_.Contains(_594_v0, _662_cf6), 1)
				(_nw93).ArraySet1(_665_cf14, 2)
				(_nw93).ArraySet1((_this).F24(), 3)
				(_nw93).ArraySet1(_665_cf14, 4)
				(_nw93).ArraySet1(((_this).F36()).IsProperSubsetOf(_dafny.MultiSetOf(false, (_this).F24())), 5)
				(_nw93).ArraySet1(!(_676_v41).Equals(_676_v41), 6)
				(_nw93).ArraySet1((_665_cf14) == (_666_cf13), 7)
				(_nw93).ArraySet1((_678_v43).Dtor_cf27(), 8)
				(_nw93).ArraySet1(false, 9)
				(_nw93).ArraySet1(!(_665_cf14), 10)
				(_nw93).ArraySet1((Companion_D2_.Create_DC7_((_672_v37).F27(), (_672_v37).F27())).Dtor_cf13(), 11)
				(_nw93).ArraySet1(_666_cf13, 12)
				(_nw93).ArraySet1(!((_679_v44).IsProperSubsetOf(_679_v44)), 13)
				(_nw93).ArraySet1((_this).F24(), 14)
				(_nw93).ArraySet1((_672_v37).F27(), 15)
				(_nw93).ArraySet1((_this).F24(), 16)
				(_nw93).ArraySet1((_672_v37).F27(), 17)
				(_nw93).ArraySet1((_this).F24(), 18)
				(_nw93).ArraySet1(true, 19)
				(_nw93).ArraySet1((_672_v37).F27(), 20)
				(_nw93).ArraySet1(!(false), 21)
				(_nw93).ArraySet1((_this).F24(), 22)
				(_nw93).ArraySet1((_this).F24(), 23)
				(_nw93).ArraySet1(!((p1).Cmp(p1) > 0), 24)
				(_nw93).ArraySet1(true, 25)
				(_nw93).ArraySet1((_this).F24(), 26)
				_680_v45 = _nw93
				var _681_v46 _dafny.Map
				_ = _681_v46
				_681_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _675_v40)
				var _rhs111 _dafny.Array = (func() _dafny.Array {
					if (_681_v46).Contains(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(675), p1)) {
						return (_681_v46).Get(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(675), p1)).(_dafny.Array)
					}
					return _675_v40
				})()
				_ = _rhs111
				var _rhs112 bool = true
				_ = _rhs112
				var _rhs113 _dafny.Array = _680_v45
				_ = _rhs113
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				_675_v40 = _rhs111
				_lhs82.F13 = _rhs112
				_680_v45 = _rhs113
			} else if _source12.Is_DC6() {
				var _682___mcc_h14 _dafny.Map = _source12.Get_().(D2_DC6).Cf12
				_ = _682___mcc_h14
				var _683_cf12 _dafny.Map = _682___mcc_h14
				_ = _683_cf12
				var _684_v47 _dafny.Map
				_ = _684_v47
				_684_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_594_v0, _dafny.IntOfUint32((_594_v0).Cardinality()))
				r0 = (func() _dafny.Int {
					if (_684_v47).Contains(_594_v0) {
						return (_684_v47).Get(_594_v0).(_dafny.Int)
					}
					return p1
				})()
				(globalState).F13 = !((_this).F24())
				var _685_v48 _dafny.Array
				_ = _685_v48
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw94
				_685_v48 = _nw94
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_685_v48), 0))
				_ = _index91
				(_685_v48).ArraySet1(p1, (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_685_v48), 0))
				_ = _index92
				(_685_v48).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_685_v48), 0))
				_ = _index93
				(_685_v48).ArraySet1((_685_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_685_v48), 0))).Int()).(_dafny.Int), (_index93).Int())
			} else {
				var _686___mcc_h15 D2 = _source12.Get_().(D2_DC8).Cf15
				_ = _686___mcc_h15
				var _687_cf15 D2 = _686___mcc_h15
				_ = _687_cf15
				var _688_v49 _dafny.MultiSet
				_ = _688_v49
				_688_v49 = _dafny.MultiSetOf(_dafny.IntOfInt64(535))
				var _689_v50 _dafny.Map
				_ = _689_v50
				_689_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), true)
				var _690_v51 _dafny.Map
				_ = _690_v51
				_690_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_689_v50, _dafny.IntOfInt64(-793))
				var _691_v52 _dafny.Array
				_ = _691_v52
				var _nwElement0_22 _dafny.Int = p1
				_ = _nwElement0_22
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_22, 0)
				(_nw95).ArraySet1((_dafny.Zero).Minus(p1), 1)
				(_nw95).ArraySet1(p1, 2)
				(_nw95).ArraySet1(p1, 3)
				(_nw95).ArraySet1((func() _dafny.Int {
					if ((_this).F36()).Contains((_this).F24()) {
						return ((_this).F36()).Multiplicity((_this).F24())
					}
					return _dafny.IntOfInt64(-489)
				})(), 4)
				(_nw95).ArraySet1((_688_v49).Cardinality(), 5)
				(_nw95).ArraySet1(p1, 6)
				(_nw95).ArraySet1(p1, 7)
				(_nw95).ArraySet1(Companion_Default___.Fm0(true, _dafny.IntOfUint32((_594_v0).Cardinality()), globalState), 8)
				(_nw95).ArraySet1(_dafny.IntOfInt64(529), 9)
				(_nw95).ArraySet1((func() _dafny.Int {
					if (_690_v51).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), true)) {
						return (_690_v51).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), true)).(_dafny.Int)
					}
					return _dafny.IntOfInt64(444)
				})(), 10)
				(_nw95).ArraySet1(p1, 11)
				(_nw95).ArraySet1((_dafny.Zero).Minus((_dafny.MultiSetOf((_596_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_596_v1).Cardinality()))).Uint32()).(bool))).Cardinality()), 12)
				(_nw95).ArraySet1(p1, 13)
				(_nw95).ArraySet1(p1, 14)
				(_nw95).ArraySet1(p1, 15)
				_691_v52 = _nw95
				var _692_v53 _dafny.Map
				_ = _692_v53
				_692_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v52, (_this).F24())
				_692_v53 = _692_v53
				var _693_v54 D6
				_ = _693_v54
				_693_v54 = Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F24()), (_this).F24(), (_this).F24())
				var _694_v55 _dafny.Map
				_ = _694_v55
				_694_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
				var _695_v56 _dafny.Set
				_ = _695_v56
				_695_v56 = _dafny.SetOf(p1, p1, p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _691_v52)).Cardinality(), p1)
				var _pat_let_tv23 = _689_v50
				_ = _pat_let_tv23
				var _696_v57 _dafny.Array
				_ = _696_v57
				var _nwElement0_23 bool = (_this).F24()
				_ = _nwElement0_23
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(20))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_23, 0)
				(_nw96).ArraySet1((_this).F24(), 1)
				(_nw96).ArraySet1((_this).F24(), 2)
				(_nw96).ArraySet1((_this).F24(), 3)
				(_nw96).ArraySet1((_this).F24(), 4)
				(_nw96).ArraySet1((_this).F24(), 5)
				(_nw96).ArraySet1((_this).F24(), 6)
				(_nw96).ArraySet1((_this).F24(), 7)
				(_nw96).ArraySet1((_this).F24(), 8)
				(_nw96).ArraySet1(!(_688_v49).Contains((_dafny.Zero).Minus(p1)), 9)
				(_nw96).ArraySet1((_this).F24(), 10)
				(_nw96).ArraySet1((_this).F24(), 11)
				(_nw96).ArraySet1((_this).F24(), 12)
				(_nw96).ArraySet1(!((func(_pat_let29_0 D6) D6 {
					return func(_697_dt__update__tmp_h4 D6) D6 {
						return func(_pat_let30_0 _dafny.Map) D6 {
							return func(_698_dt__update_hcf25_h0 _dafny.Map) D6 {
								return Companion_D6_.Create_DC16_(_698_dt__update_hcf25_h0, (_697_dt__update__tmp_h4).Dtor_cf26(), (_697_dt__update__tmp_h4).Dtor_cf27())
							}(_pat_let30_0)
						}(_pat_let_tv23)
					}(_pat_let29_0)
				}(_693_v54)).Dtor_cf27()) || ((_this).F24()), 13)
				(_nw96).ArraySet1((p1).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
					if (_694_v55).Contains((_this).F24()) {
						return (_694_v55).Get((_this).F24()).(_dafny.Int)
					}
					return p1
				})())) > 0, 14)
				(_nw96).ArraySet1((_this).F24(), 15)
				(_nw96).ArraySet1((p1).Cmp((_695_v56).Cardinality()) != 0, 16)
				(_nw96).ArraySet1(!((_this).F24()), 17)
				(_nw96).ArraySet1(((_this).F36()).IsProperSubsetOf((_this).F36()), 18)
				(_nw96).ArraySet1((_this).F24(), 19)
				_696_v57 = _nw96
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_696_v57), 0))
				_ = _index94
				(_696_v57).ArraySet1((_this).F24(), (_index94).Int())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_696_v57), 0))
				_ = _index95
				(_696_v57).ArraySet1((false) == (false), (_index95).Int())
				var _699_v58 _dafny.Sequence
				_ = _699_v58
				_699_v58 = _dafny.SeqOf(_691_v52, _691_v52, _691_v52)
				var _700_v59 _dafny.Map
				_ = _700_v59
				_700_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F24())
				_699_v58 = (func() _dafny.Sequence {
					if (func() _dafny.Set {
						var _coll63 = _dafny.NewBuilder()
						_ = _coll63
						for _iter65 := _dafny.Iterate((_700_v59).Keys().Elements()); ; {
							_compr_63, _ok65 := _iter65()
							if !_ok65 {
								break
							}
							var _701_v60 _dafny.Int
							_701_v60 = interface{}(_compr_63).(_dafny.Int)
							if (_700_v59).Contains(_701_v60) {
								_coll63.Add((_701_v60).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-807), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mo")).Cardinality())))).Cardinality()))
							}
						}
						return _coll63.ToSet()
					}()).IsSubsetOf(_695_v56) {
						return _699_v58
					}
					return _dafny.Companion_Sequence_.Concatenate(_699_v58, _699_v58)
				})()
				(globalState).F13 = (_696_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_696_v57), 0))).Int()).(bool)
			}
			var _702_v61 *C1
			_ = _702_v61
			var _nw97 *C1 = New_C1_()
			_ = _nw97
			_nw97.Ctor__((_this).F24(), (_this).F24())
			_702_v61 = _nw97
			var _703_v62 _dafny.Array
			_ = _703_v62
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_14
			var _nw98 _dafny.Array
			_ = _nw98
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw98 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Sequence = (func(_704_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_705_i2 _dafny.Int) _dafny.Sequence {
						return _704_v0
					}
				})(_594_v0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw98 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw98).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw98).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_703_v62 = _nw98
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))
			_ = _index96
			(_703_v62).ArraySet1(_594_v0, (_index96).Int())
			var _706_v63 D4
			_ = _706_v63
			_706_v63 = Companion_D4_.Create_DC10_(_594_v0)
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))
			_ = _index97
			(_703_v62).ArraySet1(((func() D4 {
				if (_this).F24() {
					return _706_v63
				}
				return _706_v63
			})()).Dtor_cf17(), (_index97).Int())
			if Companion_Default___.Fm2(globalState) {
				var _707_v64 _dafny.Array
				_ = _707_v64
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_15
				var _nw99 _dafny.Array
				_ = _nw99
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw99 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Int = (func(_708_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_709_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_709_i3, _708_p1)
						}
					})(p1)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw99 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw99).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw99).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_707_v64 = _nw99
				var _710_v65 _dafny.Map
				_ = _710_v65
				_710_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _707_v64)
				_710_v65 = (_710_v65).Update(p1, _707_v64)
				(globalState).F13 = !((_this).F36()).Contains((_702_v61).F39())
				var _711_v66 _dafny.Array
				_ = _711_v66
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw100
				_711_v66 = _nw100
				var _712_v67 _dafny.Set
				_ = _712_v67
				_712_v67 = _dafny.SetOf(p1)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_711_v66), 0))
				_ = _index98
				(_711_v66).ArraySet1(!((_702_v61).Fm32((_702_v61).F39(), (_712_v67).Cardinality(), globalState)), (_index98).Int())
				var _713_v68 _dafny.Sequence
				_ = _713_v68
				_713_v68 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("sn"), (_703_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))).Int()).(_dafny.Sequence))
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_711_v66), 0))
				_ = _index99
				var _rhs114 bool = (_702_v61).F39()
				_ = _rhs114
				var _rhs115 _dafny.Int = Companion_Default___.SafeModuloInt(p1, p1)
				_ = _rhs115
				var _rhs116 _dafny.Sequence = (_713_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_594_v0).Cardinality()), _dafny.IntOfUint32((_713_v68).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs116
				var _rhs117 bool = (p1).Cmp(p1) <= 0
				_ = _rhs117
				var _rhs118 _dafny.Int = p1
				_ = _rhs118
				var _lhs83 *GlobalState = globalState
				_ = _lhs83
				var _lhs84 _dafny.Array = _711_v66
				_ = _lhs84
				var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_711_v66), 0))
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				_lhs83.F13 = _rhs114
				r2 = _rhs115
				r1 = _rhs116
				(_lhs84).ArraySet1(_rhs117, (_lhs85).Int())
				_lhs86.F7 = _rhs118
				var _714_v69 _dafny.Array
				_ = _714_v69
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_16
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Map = (func(_715_p1 _dafny.Int, _716_v61 *C1) func(_dafny.Int) _dafny.Map {
						return func(_717_i4 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-632), _715_p1)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_716_v61).F39(), _715_p1))
						}
					})(p1, _702_v61)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw101 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw101).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw101).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_714_v69 = _nw101
				var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
				_ = _nw102
				_714_v69 = _nw102
				var _718_v70 _dafny.Set
				_ = _718_v70
				_718_v70 = _dafny.SetOf((_this).F24())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_711_v66), 0))
				_ = _index100
				(_711_v66).ArraySet1((func() bool {
					if (_711_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_711_v66), 0))).Int()).(bool) {
						return (_this).F24()
					}
					return (_718_v70).IsSubsetOf(_718_v70)
				})(), (_index100).Int())
			} else {
				(globalState).F13 = (func() bool {
					if _dafny.Companion_Sequence_.Equal((_703_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))).Int()).(_dafny.Sequence), (_703_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))).Int()).(_dafny.Sequence)) {
						return (_596_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_596_v1).Cardinality()))).Uint32()).(bool)
					}
					return !((_this).F24())
				})()
				(globalState).F13 = (_702_v61).F39()
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))
				_ = _index101
				(_703_v62).ArraySet1((_703_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_703_v62), 0))).Int()).(_dafny.Sequence), (_index101).Int())
				(globalState).F13 = (_702_v61).F39()
				var _719_v71 D2
				_ = _719_v71
				_719_v71 = Companion_D2_.Create_DC7_(true, (_702_v61).F39())
				var _720_v72 _dafny.Array
				_ = _720_v72
				var _nwElement0_24 D2 = _719_v71
				_ = _nwElement0_24
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(23))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_24, 0)
				(_nw103).ArraySet1(_719_v71, 1)
				(_nw103).ArraySet1(_719_v71, 2)
				(_nw103).ArraySet1(Companion_D2_.Create_DC7_((_this).Fm4(p1, _596_v1, globalState), (_702_v61).F39()), 3)
				(_nw103).ArraySet1(_719_v71, 4)
				(_nw103).ArraySet1(_719_v71, 5)
				(_nw103).ArraySet1(_719_v71, 6)
				(_nw103).ArraySet1(_719_v71, 7)
				(_nw103).ArraySet1(_719_v71, 8)
				(_nw103).ArraySet1(Companion_D2_.Create_DC7_((_702_v61).F39(), (_702_v61).F39()), 9)
				(_nw103).ArraySet1(_719_v71, 10)
				(_nw103).ArraySet1(_719_v71, 11)
				(_nw103).ArraySet1(_719_v71, 12)
				(_nw103).ArraySet1(_719_v71, 13)
				(_nw103).ArraySet1(func(_pat_let31_0 D2) D2 {
					return func(_721_dt__update__tmp_h5 D2) D2 {
						return func(_pat_let32_0 bool) D2 {
							return func(_722_dt__update_hcf13_h0 bool) D2 {
								return Companion_D2_.Create_DC7_(_722_dt__update_hcf13_h0, (_721_dt__update__tmp_h5).Dtor_cf14())
							}(_pat_let32_0)
						}((_this).F24())
					}(_pat_let31_0)
				}(_719_v71), 14)
				(_nw103).ArraySet1(_719_v71, 15)
				(_nw103).ArraySet1(_719_v71, 16)
				(_nw103).ArraySet1(Companion_Default___.Fm38(true, (_702_v61).F39(), _dafny.IntOfUint32((_596_v1).Cardinality()), globalState), 17)
				(_nw103).ArraySet1(_719_v71, 18)
				(_nw103).ArraySet1(_719_v71, 19)
				(_nw103).ArraySet1(_719_v71, 20)
				(_nw103).ArraySet1(_719_v71, 21)
				(_nw103).ArraySet1(_719_v71, 22)
				_720_v72 = _nw103
				var _723_v73 _dafny.Sequence
				_ = _723_v73
				_723_v73 = _dafny.SeqOf(_720_v72, _720_v72, _720_v72, _720_v72, _720_v72)
				var _724_v74 _dafny.Map
				_ = _724_v74
				_724_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_702_v61).F39())
				var _725_v75 _dafny.Sequence
				_ = _725_v75
				_725_v75 = _dafny.SeqOf(p1)
				var _726_v76 _dafny.Sequence
				_ = _726_v76
				_726_v76 = _dafny.SeqOf((_dafny.Zero).Minus((_725_v75).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_725_v75).Cardinality()))).Uint32()).(_dafny.Int)))
				var _rhs119 _dafny.Sequence = _723_v73
				_ = _rhs119
				var _rhs120 bool = (_702_v61).F39()
				_ = _rhs120
				var _rhs121 bool = (_this).F24()
				_ = _rhs121
				var _rhs122 _dafny.Int = (_dafny.Zero).Minus((p1).Times(Companion_Default___.SafeModuloInt((_724_v74).Cardinality(), (_726_v76).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_726_v76).Cardinality()))).Uint32()).(_dafny.Int))))
				_ = _rhs122
				var _rhs123 bool = true
				_ = _rhs123
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				_723_v73 = _rhs119
				_lhs87.F13 = _rhs120
				_lhs88.F13 = _rhs121
				r2 = _rhs122
				_lhs89.F13 = _rhs123
			}
		}
		var _727_v77 _dafny.Array
		_ = _727_v77
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_17
		var _nw104 _dafny.Array
		_ = _nw104
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw104 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.CodePoint = func(_728_i6 _dafny.Int) _dafny.CodePoint {
				return _this.F23()
			}
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw104 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw104).ArraySet1CodePoint(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw104).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_727_v77 = _nw104
		for _iter66 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_727_v77), 0))); ; {
			_guard_loop_1, _ok66 := _iter66()
			if !_ok66 {
				break
			}
			var _729_i5 _dafny.Int
			_729_i5 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_729_i5).Sign() != -1) && ((_729_i5).Cmp(_dafny.ArrayLen((_727_v77), 0)) < 0)) {
				(_727_v77).ArraySet1CodePoint(_this.F23(), (_729_i5).Int())
			}
		}
		var _730_v78 T3
		_ = _730_v78
		var _nw105 *C1 = New_C1_()
		_ = _nw105
		_nw105.Ctor__((_this).F24(), (_this).F24())
		_730_v78 = _nw105
		var _731_v79 _dafny.Array
		_ = _731_v79
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_18
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) bool = func(_732_i7 _dafny.Int) bool {
				return (_this).F24()
			}
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw106 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw106).ArraySet1(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw106).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		_731_v79 = _nw106
		var _733_v80 _dafny.Map
		_ = _733_v80
		_733_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_730_v78, (func() _dafny.Array {
			if (_730_v78).F27() {
				return _731_v79
			}
			return _731_v79
		})())
		_733_v80 = ((_733_v80).Merge(_733_v80)).Merge((_733_v80).Update(_730_v78, _731_v79))
		if _dafny.Companion_Sequence_.Contains(_594_v0, _this.F23()) {
			var _734_v81 D2
			_ = _734_v81
			_734_v81 = Companion_D2_.Create_DC7_((_730_v78).F27(), (_730_v78).F27())
			(globalState).F13 = (_734_v81).Dtor_cf14()
			var _735_v82 _dafny.Map
			_ = _735_v82
			_735_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_730_v78).F27(), _594_v0)
			var _736_v83 *C0
			_ = _736_v83
			var _nw107 *C0 = New_C0_()
			_ = _nw107
			_nw107.Ctor__(_735_v82, !((_this).F24()), (_594_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-439), _dafny.IntOfUint32((_594_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), (p1).Cmp(_dafny.IntOfUint32((_594_v0).Cardinality())) >= 0)
			_736_v83 = _nw107
			var _737_v84 _dafny.Map
			_ = _737_v84
			_737_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _596_v1)
			var _738_v85 D6
			_ = _738_v85
			_738_v85 = Companion_D6_.Create_DC15_(_737_v84)
			var _pat_let_tv24 = _596_v1
			_ = _pat_let_tv24
			var _739_v87 _dafny.Array
			_ = _739_v87
			var _nwElement0_25 D6 = _738_v85
			_ = _nwElement0_25
			var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(8))
			_ = _nw108
			(_nw108).ArraySet1(_nwElement0_25, 0)
			(_nw108).ArraySet1(Companion_Default___.Fm39(p1, globalState), 1)
			(_nw108).ArraySet1((func() D6 {
				if (_736_v83).F38() {
					return _738_v85
				}
				return _738_v85
			})(), 2)
			(_nw108).ArraySet1(func(_pat_let33_0 D6) D6 {
				return func(_740_dt__update__tmp_h6 D6) D6 {
					return func(_pat_let34_0 _dafny.Map) D6 {
						return func(_742_dt__update_hcf24_h0 _dafny.Map) D6 {
							return Companion_D6_.Create_DC15_(_742_dt__update_hcf24_h0)
						}(_pat_let34_0)
					}(func() _dafny.Map {
						var _coll64 = _dafny.NewMapBuilder()
						_ = _coll64
						for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(36), _dafny.IntOfInt64(-788))); ; {
							_compr_64, _ok67 := _iter67()
							if !_ok67 {
								break
							}
							var _741_v86 _dafny.Int
							_741_v86 = interface{}(_compr_64).(_dafny.Int)
							if ((_dafny.IntOfInt64(36)).Cmp(_741_v86) <= 0) && ((_741_v86).Cmp(_dafny.IntOfInt64(-788)) < 0) {
								_coll64.Add((_741_v86).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("exgcvpsfo")).Cardinality()))).Cardinality())), _pat_let_tv24)
							}
						}
						return _coll64.ToMap()
					}())
				}(_pat_let33_0)
			}(Companion_Default___.Fm39(_dafny.IntOfInt64(278), globalState)), 3)
			(_nw108).ArraySet1(_738_v85, 4)
			(_nw108).ArraySet1(_738_v85, 5)
			(_nw108).ArraySet1(Companion_D6_.Create_DC15_(_737_v84), 6)
			(_nw108).ArraySet1(_738_v85, 7)
			_739_v87 = _nw108
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_739_v87), 0))
			_ = _index102
			(_739_v87).ArraySet1(_738_v85, (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_739_v87), 0))
			_ = _index103
			(_739_v87).ArraySet1(_738_v85, (_index103).Int())
			var _743_v88 D2
			_ = _743_v88
			_743_v88 = Companion_D2_.Create_DC6_((_735_v82).Merge(_735_v82))
			_743_v88 = _743_v88
			var _744_v89 _dafny.Sequence
			_ = _744_v89
			_744_v89 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-678))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}(func(_745_i8 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(556)
			}))).Cardinality())))
			(globalState).F13 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(54), p1, p1), _dafny.SeqOf(p1)), _dafny.Companion_Sequence_.Concatenate(_744_v89, _744_v89))
		} else {
			var _746_v90 _dafny.Map
			_ = _746_v90
			_746_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_730_v78).F27(), (_730_v78).F27())
			var _747_v91 _dafny.Sequence
			_ = _747_v91
			_747_v91 = _dafny.SeqOf(_746_v90, _746_v90, _746_v90, _746_v90)
			var _748_v92 _dafny.Map
			_ = _748_v92
			_748_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_597_v2, _747_v91)
			r0 = (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_748_v92).Contains(_597_v2) {
					return (_748_v92).Get(_597_v2).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_747_v91, _dafny.SeqOf(_746_v90))
			})()).Cardinality()))
			var _749_v93 D0
			_ = _749_v93
			_749_v93 = Companion_D0_.Create_DC2_((_this).Fm4(p1, _596_v1, globalState), p1, (_730_v78).F27())
			var _750_v94 _dafny.Map
			_ = _750_v94
			_750_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v93, (_this).F24())
			if (func() bool {
				if (_750_v94).Contains(_749_v93) {
					return (_750_v94).Get(_749_v93).(bool)
				}
				return (_730_v78).F27()
			})() {
				(_this).F23_set_((func() _dafny.CodePoint {
					if (_this).F24() {
						return Companion_Default___.Fm29(p1, (_730_v78).F27(), (_730_v78).F27(), p1, globalState)
					}
					return _this.F23()
				})())
				(globalState).F7 = (_dafny.IntOfUint32((_594_v0).Cardinality())).Plus((_dafny.Zero).Minus(p1))
				var _751_v95 _dafny.Map
				_ = _751_v95
				_751_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm40((_730_v78).F27(), globalState)).Dtor_cf20(), !((_730_v78).F27()))
				var _752_v96 _dafny.Sequence
				_ = _752_v96
				_752_v96 = _dafny.SeqOf(p1)
				_751_v95 = (_751_v95).Update(((_730_v78).Fm9(_752_v96, p1, (_730_v78).F27(), globalState)).Plus(p1), (_730_v78).F27())
				(globalState).F7 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((p1).Times(_dafny.IntOfUint32((_752_v96).Cardinality()))), p1)
				var _753_v97 _dafny.Set
				_ = _753_v97
				_753_v97 = _dafny.SetOf((_730_v78).F27(), (_730_v78).F27())
				_751_v95 = (_751_v95).Update((_753_v97).Cardinality(), _dafny.Companion_Sequence_.Equal(_752_v96, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vbav")).Cardinality()), p1, (_753_v97).Cardinality())))
			} else {
				var _754_v98 _dafny.Sequence
				_ = _754_v98
				_754_v98 = _dafny.SeqOf(_594_v0, _594_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_594_v0, _594_v0), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_594_v0, _594_v0)).Cardinality()))).Uint32(), _this.F23()), _dafny.Companion_Sequence_.Concatenate(_594_v0, _594_v0))
				r1 = (_754_v98).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_754_v98).Cardinality()))).Uint32()).(_dafny.Sequence)
				(globalState).F13 = (_730_v78).F27()
				(globalState).F7 = p1
				(globalState).F7 = p1
				var _755_v99 _dafny.Set
				_ = _755_v99
				_755_v99 = _dafny.SetOf(_dafny.CodePoint('r'), _this.F23(), _this.F23())
				_755_v99 = _dafny.SetOf((func(_pat_let35_0 D1) D1 {
					return func(_756_dt__update__tmp_h7 D1) D1 {
						return func(_pat_let36_0 _dafny.CodePoint) D1 {
							return func(_757_dt__update_hcf6_h0 _dafny.CodePoint) D1 {
								return Companion_D1_.Create_DC4_(_757_dt__update_hcf6_h0)
							}(_pat_let36_0)
						}(_this.F23())
					}(_pat_let35_0)
				}(_597_v2)).Dtor_cf6(), _this.F23())
			}
			if (_730_v78).F27() {
				var _rhs124 _dafny.CodePoint = _this.F23()
				_ = _rhs124
				var _rhs125 bool = (_730_v78).F27()
				_ = _rhs125
				var _rhs126 bool = (_this).F24()
				_ = _rhs126
				var _lhs90 *C2 = _this
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				_lhs90.F23_set_(_rhs124)
				_lhs91.F13 = _rhs125
				_lhs92.F13 = _rhs126
				(globalState).F13 = (_730_v78).F27()
				var _758_v100 D2
				_ = _758_v100
				_758_v100 = Companion_D2_.Create_DC7_((_730_v78).F27(), (_730_v78).F27())
				var _759_v101 D5
				_ = _759_v101
				_759_v101 = Companion_D5_.Create_DC13_(_dafny.IntOfInt64(377), !_dafny.Companion_Sequence_.Equal(_596_v1, _596_v1), !((_758_v100).Dtor_cf13()) || (Companion_Default___.Fm2(globalState)))
				_759_v101 = _759_v101
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw109
				_731_v79 = _nw109
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw110
				_731_v79 = _nw110
			} else {
				var _760_v102 _dafny.Map
				_ = _760_v102
				_760_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _594_v0)
				var _761_v103 _dafny.MultiSet
				_ = _761_v103
				_761_v103 = _dafny.MultiSetOf(_dafny.IntOfInt64(975))
				var _762_v104 *C0
				_ = _762_v104
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__(_760_v102, (_this).Fm4((_761_v103).Cardinality(), _596_v1, globalState), _this.F23(), (_730_v78).F27())
				_762_v104 = _nw111
				var _763_v105 _dafny.Array
				_ = _763_v105
				var _nw112 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
				_ = _nw112
				_763_v105 = _nw112
				var _764_v106 *C1
				_ = _764_v106
				var _nw113 *C1 = New_C1_()
				_ = _nw113
				_nw113.Ctor__((_730_v78).F27(), (_730_v78).F27())
				_764_v106 = _nw113
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_763_v105), 0))
				_ = _index104
				(_763_v105).ArraySet1(_764_v106, (_index104).Int())
				var _765_v107 *C1
				_ = _765_v107
				_765_v107 = _764_v106
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_763_v105), 0))
				_ = _index105
				(_763_v105).ArraySet1((_765_v107), (_index105).Int())
				var _766_v108 bool
				_ = _766_v108
				var _out30 bool
				_ = _out30
				_out30 = (_764_v106).M4((_this).F24(), globalState)
				_766_v108 = _out30
				_766_v108 = !((_this).F24())
				(globalState).F13 = ((_this).F24()) || (((_this).F24()) || ((_730_v78).F27()))
			}
			r0 = (_dafny.IntOfUint32((_594_v0).Cardinality())).Times((_this).Fm24(p1, (_730_v78).F27(), globalState))
			var _767_v109 _dafny.Sequence
			_ = _767_v109
			_767_v109 = _dafny.SeqOf(p1, p1, Companion_Default___.Fm0((_730_v78).F27(), _dafny.IntOfInt64(395), globalState))
			var _768_v110 _dafny.Map
			_ = _768_v110
			_768_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), (_this).F24())
			(globalState).F7 = (_dafny.IntOfInt64(571)).Minus((_dafny.Zero).Minus((_767_v109).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_768_v110)).Cardinality()), _dafny.IntOfUint32((_767_v109).Cardinality()))).Uint32()).(_dafny.Int)))
		}
		r0 = (func() _dafny.Int {
			if (_this).F24() {
				return (func() _dafny.Int {
					if (_730_v78).F27() {
						return (_730_v78).Fm9(_dafny.SeqOf(p1, p1), p1, (_this).F24(), globalState)
					}
					return p1
				})()
			}
			return (_dafny.Zero).Minus((p1).Minus(p1))
		})()
		r1 = _594_v0
		r2 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(406), (func() _dafny.Int {
			if (_this).F24() {
				return p1
			}
			return p1
		})())
		return r0, r1, r2
	}
}
func (_this *C2) F36() _dafny.MultiSet {
	{
		return _this._f36
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f23 _dafny.CodePoint
	_f24 bool
	_f25 _dafny.Array
	_f35 _dafny.Map
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
	_this._f25 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f35 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C3) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C3) F24() bool {
	return _this._f24
}
func (_this *C3) F25() _dafny.Array {
	return _this._f25
}
func (_this *C3) Ctor__(f35 _dafny.Map, f25 _dafny.Array, f23 _dafny.CodePoint, f24 bool) {
	{
		(_this)._f35 = f35
		(_this)._f25 = f25
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C3) Fm7(globalState *GlobalState) _dafny.Sequence {
	{
		if ((_this).F24()) || (false) {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(926))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}(func(_769_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("er")).Cardinality())
			}))
		} else {
			return _dafny.SeqOf(_dafny.IntOfInt64(278))
		}
	}
}
func (_this *C3) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("n"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("f")))
	}
}
func (_this *C3) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return ((_dafny.IntOfInt64(538)).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(947))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}(func(_770_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		}))).Cardinality()))).Cmp(((func() _dafny.Set {
			if (_this).F24() {
				return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("qkrmexiw"))
			}
			return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("aaucpfakq"), _dafny.UnicodeSeqOfUtf8Bytes("hrrodm"), _dafny.UnicodeSeqOfUtf8Bytes("qgsjiaqub"))
		})()).Cardinality()) <= 0
	}
}
func (_this *C3) M3(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _771_v0 _dafny.Int
		_ = _771_v0
		_771_v0 = _dafny.IntOfInt64(331)
		(globalState).F7 = (_771_v0).Minus(_771_v0)
		var _772_v1 _dafny.Sequence
		_ = _772_v1
		_772_v1 = _dafny.UnicodeSeqOfUtf8Bytes("gbplvd")
		var _773_v2 D4
		_ = _773_v2
		_773_v2 = Companion_D4_.Create_DC10_(_772_v1)
		_772_v1 = _dafny.Companion_Sequence_.Concatenate(_772_v1, (_773_v2).Dtor_cf17())
		var _774_v3 _dafny.MultiSet
		_ = _774_v3
		_774_v3 = _dafny.MultiSetOf(_771_v0, _771_v0, _771_v0, _771_v0, _771_v0)
		var _775_i0 _dafny.Int
		_ = _775_i0
		_775_i0 = _dafny.Zero
		{
			for !((_774_v3).IsDisjointFrom((func() _dafny.MultiSet {
				if (_this).F24() {
					return _774_v3
				}
				return _774_v3
			})())) {
				{
					if (_775_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_775_i0 = (_775_i0).Plus(_dafny.One)
					var _776_v4 bool
					_ = _776_v4
					var _777_v5 _dafny.CodePoint
					_ = _777_v5
					var _778_v6 bool
					_ = _778_v6
					var _out31 bool
					_ = _out31
					var _out32 _dafny.CodePoint
					_ = _out32
					var _out33 bool
					_ = _out33
					_out31, _out32, _out33 = (_this).M11(p0, globalState)
					_776_v4 = _out31
					_777_v5 = _out32
					_778_v6 = _out33
					var _779_v7 _dafny.Map
					_ = _779_v7
					_779_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _772_v1)
					var _780_v8 *C0
					_ = _780_v8
					var _nw114 *C0 = New_C0_()
					_ = _nw114
					_nw114.Ctor__(_779_v7, _778_v6, p1, true)
					_780_v8 = _nw114
					r1 = _771_v0
					(globalState).F13 = (_this).F24()
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _781_v9 _dafny.Map
		_ = _781_v9
		_781_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_772_v1, _771_v0)
		(globalState).F7 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_772_v1, _771_v0)).Merge(_781_v9)).Cardinality()
		r0 = p0
		(globalState).F13 = (_this).F24()
		r0 = p0
		r1 = _771_v0
		var _782_v10 _dafny.MultiSet
		_ = _782_v10
		_782_v10 = _dafny.MultiSetOf((_this).F24(), p0)
		var _783_v11 _dafny.Map
		_ = _783_v11
		_783_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_772_v1, _782_v10)
		r2 = !((func() _dafny.MultiSet {
			if (_783_v11).Contains(_772_v1) {
				return (_783_v11).Get(_772_v1).(_dafny.MultiSet)
			}
			return _782_v10
		})()).Contains(p0)
		return r0, r1, r2
	}
}
func (_this *C3) M11(p0 bool, globalState *GlobalState) (bool, _dafny.CodePoint, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var r2 bool = false
		_ = r2
		var _784_v0 D5
		_ = _784_v0
		_784_v0 = Companion_D5_.Create_DC14_(p0)
		r0 = (_784_v0).Dtor_cf23()
		var _785_i0 _dafny.Int
		_ = _785_i0
		_785_i0 = _dafny.Zero
		{
			for (_this).F24() {
				{
					if (_785_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_785_i0 = (_785_i0).Plus(_dafny.One)
					var _786_v1 _dafny.Int
					_ = _786_v1
					_786_v1 = _dafny.IntOfInt64(49)
					(globalState).F7 = _786_v1
					var _source13 D5 = _784_v0
					_ = _source13
					if _source13.Is_DC13() {
						var _787___mcc_h0 _dafny.Int = _source13.Get_().(D5_DC13).Cf20
						_ = _787___mcc_h0
						var _788___mcc_h1 bool = _source13.Get_().(D5_DC13).Cf21
						_ = _788___mcc_h1
						var _789___mcc_h2 bool = _source13.Get_().(D5_DC13).Cf22
						_ = _789___mcc_h2
						var _790_cf22 bool = _789___mcc_h2
						_ = _790_cf22
						var _791_cf21 bool = _788___mcc_h1
						_ = _791_cf21
						var _792_cf20 _dafny.Int = _787___mcc_h0
						_ = _792_cf20
						var _793_v2 _dafny.Sequence
						_ = _793_v2
						_793_v2 = _dafny.SeqOf(_786_v1, (_dafny.Zero).Minus(_786_v1))
						var _794_v3 _dafny.Sequence
						_ = _794_v3
						_794_v3 = _dafny.SeqOf(_dafny.SeqOf(_786_v1, _dafny.IntOfInt64(-258), _786_v1, _786_v1, _792_cf20))
						(globalState).F7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_793_v2, (_794_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(690))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}(func(_795_i1 _dafny.Int) _dafny.CodePoint {
							return _this.F23()
						}))).Cardinality()), _dafny.IntOfUint32((_794_v3).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality())
						(_this).F23_set_(_this.F23())
						var _796_v4 *C1
						_ = _796_v4
						var _nw115 *C1 = New_C1_()
						_ = _nw115
						_nw115.Ctor__(_790_cf22, p0)
						_796_v4 = _nw115
						(globalState).F7 = _786_v1
					} else if _source13.Is_DC14() {
						var _797___mcc_h3 bool = _source13.Get_().(D5_DC14).Cf23
						_ = _797___mcc_h3
						var _798_cf23 bool = _797___mcc_h3
						_ = _798_cf23
						var _799_v5 _dafny.Array
						_ = _799_v5
						var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
						_ = _nw116
						_799_v5 = _nw116
						(globalState).F22 = _799_v5
						var _800_v6 _dafny.Set
						_ = _800_v6
						_800_v6 = _dafny.SetOf(p0)
						_800_v6 = Companion_Default___.Fm41(globalState)
						var _801_v7 _dafny.Sequence
						_ = _801_v7
						_801_v7 = _dafny.SeqOf((_this).F24(), p0, false, !(_798_cf23))
						var _802_v8 _dafny.Map
						_ = _802_v8
						_802_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_801_v7, _786_v1)
						var _803_v9 _dafny.Map
						_ = _803_v9
						_803_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm36(globalState))
						var _804_v10 _dafny.Array
						_ = _804_v10
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(4)
						_ = _len0_19
						var _nw117 _dafny.Array
						_ = _nw117
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw117 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) bool = func(_805_i2 _dafny.Int) bool {
								return false
							}
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw117 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw117).ArraySet1(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw117).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_804_v10 = _nw117
						var _806_v11 _dafny.Map
						_ = _806_v11
						_806_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_804_v10, _786_v1)
						_802_v8 = (_802_v8).Update(_dafny.Companion_Sequence_.Concatenate(_801_v7, (func() _dafny.Sequence {
							if (_803_v9).Contains(p0) {
								return (_803_v9).Get(p0).(_dafny.Sequence)
							}
							return _801_v7
						})()), (func() _dafny.Int {
							if (_806_v11).Contains(_804_v10) {
								return (_806_v11).Get(_804_v10).(_dafny.Int)
							}
							return _786_v1
						})())
						_798_cf23 = p0
					} else {
						var _807___mcc_h4 _dafny.Set = _source13.Get_().(D5_DC12).Cf19
						_ = _807___mcc_h4
						var _808_cf19 _dafny.Set = _807___mcc_h4
						_ = _808_cf19
						var _809_v12 _dafny.Set
						_ = _809_v12
						_809_v12 = _dafny.SetOf(_this.F23(), _this.F23(), _this.F23())
						var _810_v13 _dafny.Sequence
						_ = _810_v13
						_810_v13 = _dafny.SeqOf(_786_v1)
						var _811_v14 _dafny.Sequence
						_ = _811_v14
						_811_v14 = _dafny.UnicodeSeqOfUtf8Bytes("cemw")
						var _812_v15 _dafny.Map
						_ = _812_v15
						_812_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_786_v1, (_dafny.Zero).Minus((_810_v13).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(p0, _dafny.IntOfUint32((_811_v14).Cardinality()), globalState), _dafny.IntOfUint32((_810_v13).Cardinality()))).Uint32()).(_dafny.Int)))
						var _813_v16 _dafny.Map
						_ = _813_v16
						_813_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v12).IsProperSubsetOf(_809_v12), (_812_v15).Cardinality())
						_813_v16 = (_813_v16).Update(p0, _786_v1)
						var _814_v18 _dafny.Map
						_ = _814_v18
						_814_v18 = (_812_v15).Merge(func() _dafny.Map {
							var _coll65 = _dafny.NewMapBuilder()
							_ = _coll65
							for _iter68 := _dafny.Iterate((_810_v13).Elements()); ; {
								_compr_65, _ok68 := _iter68()
								if !_ok68 {
									break
								}
								var _815_v17 _dafny.Int
								_815_v17 = interface{}(_compr_65).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_810_v13, _815_v17) {
									_coll65.Add(Companion_Default___.SafeDivisionInt(_815_v17, _786_v1), _786_v1)
								}
							}
							return _coll65.ToMap()
						}())
						_814_v18 = _812_v15
						var _816_v19 _dafny.Sequence
						_ = _816_v19
						_816_v19 = _dafny.SeqOf(p0)
						var _817_v20 D0
						_ = _817_v20
						_817_v20 = Companion_D0_.Create_DC0_(_dafny.MultiSetFromSeq(_816_v19))
						var _818_v21 _dafny.Map
						_ = _818_v21
						_818_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
						var _819_v22 _dafny.Array
						_ = _819_v22
						var _nwElement0_26 bool = p0
						_ = _nwElement0_26
						var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(26))
						_ = _nw118
						(_nw118).ArraySet1(_nwElement0_26, 0)
						(_nw118).ArraySet1((_786_v1).Cmp(Companion_Default___.Fm0(!((_this).F24()), _786_v1, globalState)) < 0, 1)
						(_nw118).ArraySet1((_this).F24(), 2)
						(_nw118).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_816_v19, _816_v19), 3)
						(_nw118).ArraySet1((_this).F24(), 4)
						(_nw118).ArraySet1(p0, 5)
						(_nw118).ArraySet1((_this).F24(), 6)
						(_nw118).ArraySet1(!(false), 7)
						(_nw118).ArraySet1((p0) || (p0), 8)
						(_nw118).ArraySet1(p0, 9)
						(_nw118).ArraySet1((_this).F24(), 10)
						(_nw118).ArraySet1(p0, 11)
						(_nw118).ArraySet1(((_817_v20).Dtor_cf0()).IsDisjointFrom(_dafny.MultiSetOf(p0, (_this).F24())), 12)
						(_nw118).ArraySet1(((_818_v21).Cardinality()).Cmp((_818_v21).Cardinality()) <= 0, 13)
						(_nw118).ArraySet1(p0, 14)
						(_nw118).ArraySet1(Companion_Default___.Fm2(globalState), 15)
						(_nw118).ArraySet1(p0, 16)
						(_nw118).ArraySet1(p0, 17)
						(_nw118).ArraySet1((_816_v19).Select((Companion_Default___.SafeIndex(_786_v1, _dafny.IntOfUint32((_816_v19).Cardinality()))).Uint32()).(bool), 18)
						(_nw118).ArraySet1(p0, 19)
						(_nw118).ArraySet1((_this).F24(), 20)
						(_nw118).ArraySet1((_this).F24(), 21)
						(_nw118).ArraySet1((_dafny.IntOfUint32((_811_v14).Cardinality())).Cmp(_786_v1) == 0, 22)
						(_nw118).ArraySet1(p0, 23)
						(_nw118).ArraySet1((_this).F24(), 24)
						(_nw118).ArraySet1((_this).F24(), 25)
						_819_v22 = _nw118
						var _820_v23 _dafny.Map
						_ = _820_v23
						_820_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-580), (_this).F24())
						var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_819_v22), 0))
						_ = _index106
						(_819_v22).ArraySet1((func() bool {
							if (_820_v23).Contains(_dafny.IntOfUint32((_810_v13).Cardinality())) {
								return (_820_v23).Get(_dafny.IntOfUint32((_810_v13).Cardinality())).(bool)
							}
							return !((_this).F24())
						})(), (_index106).Int())
						var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_819_v22), 0))
						_ = _index107
						(_819_v22).ArraySet1((_this).F24(), (_index107).Int())
						var _821_v24 _dafny.Array
						_ = _821_v24
						var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
						_ = _nw119
						_821_v24 = _nw119
						var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_821_v24), 0))
						_ = _index108
						(_821_v24).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_786_v1, _786_v1)), (_index108).Int())
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_821_v24), 0))
						_ = _index109
						(_821_v24).ArraySet1(_786_v1, (_index109).Int())
					}
					var _822_v25 _dafny.Array
					_ = _822_v25
					var _len0_20 _dafny.Int = _dafny.One
					_ = _len0_20
					var _nw120 _dafny.Array
					_ = _nw120
					if _len0_20.Cmp(_dafny.Zero) == 0 {
						_nw120 = _dafny.NewArray(_len0_20)
					} else {
						var _init20 func(_dafny.Int) D4 = (func(_823_v1 _dafny.Int) func(_dafny.Int) D4 {
							return func(_824_i3 _dafny.Int) D4 {
								return Companion_D4_.Create_DC11_(_823_v1)
							}
						})(_786_v1)
						_ = _init20
						var _element0_20 = _init20(_dafny.Zero)
						_ = _element0_20
						_nw120 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
						(_nw120).ArraySet1(_element0_20, 0)
						var _nativeLen0_20 = (_len0_20).Int()
						_ = _nativeLen0_20
						for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
							(_nw120).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
						}
					}
					_822_v25 = _nw120
					var _825_v26 D4
					_ = _825_v26
					_825_v26 = Companion_D4_.Create_DC11_(_dafny.IntOfInt64(911))
					var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_822_v25), 0))
					_ = _index110
					(_822_v25).ArraySet1(_825_v26, (_index110).Int())
					var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_822_v25), 0))
					_ = _index111
					(_822_v25).ArraySet1(_825_v26, (_index111).Int())
					var _826_v27 _dafny.Array
					_ = _826_v27
					var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
					_ = _nw121
					_826_v27 = _nw121
					var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_826_v27), 0))
					_ = _index112
					(_826_v27).ArraySet1((_dafny.Zero).Minus(_786_v1), (_index112).Int())
					var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_826_v27), 0))
					_ = _index113
					(_826_v27).ArraySet1((_786_v1).Times(_786_v1), (_index113).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		(globalState).F13 = Companion_Default___.Fm2(globalState)
		var _827_v28 _dafny.Sequence
		_ = _827_v28
		_827_v28 = _dafny.UnicodeSeqOfUtf8Bytes("plwsf")
		var _828_v29 _dafny.Int
		_ = _828_v29
		_828_v29 = _dafny.IntOfInt64(669)
		var _829_v30 _dafny.Array
		_ = _829_v30
		var _nwElement0_27 _dafny.Sequence = _827_v28
		_ = _nwElement0_27
		var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(24))
		_ = _nw122
		(_nw122).ArraySet1(_nwElement0_27, 0)
		(_nw122).ArraySet1(_827_v28, 1)
		(_nw122).ArraySet1(_827_v28, 2)
		(_nw122).ArraySet1(_827_v28, 3)
		(_nw122).ArraySet1(_827_v28, 4)
		(_nw122).ArraySet1(_827_v28, 5)
		(_nw122).ArraySet1(_827_v28, 6)
		(_nw122).ArraySet1(_827_v28, 7)
		(_nw122).ArraySet1(_827_v28, 8)
		(_nw122).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(774))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}(func(_830_i4 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		})), 9)
		(_nw122).ArraySet1(_827_v28, 10)
		(_nw122).ArraySet1(_827_v28, 11)
		(_nw122).ArraySet1(_827_v28, 12)
		(_nw122).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_827_v28, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-90))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg67 _dafny.Int) interface{} {
				return coer67(arg67)
			}
		}(func(_831_i5 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		})), (Companion_Default___.SafeIndex(_828_v29, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-90))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}(func(_832_i5 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		}))).Cardinality()))).Uint32(), _this.F23())), 13)
		(_nw122).ArraySet1(_827_v28, 14)
		(_nw122).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_827_v28, _827_v28), 15)
		(_nw122).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(439))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_833_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), 16)
		(_nw122).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(769))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}(func(_834_i7 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		})), 17)
		(_nw122).ArraySet1((func() _dafny.Sequence {
			if (_this).F24() {
				return _dafny.UnicodeSeqOfUtf8Bytes("lti")
			}
			return _827_v28
		})(), 18)
		(_nw122).ArraySet1(_827_v28, 19)
		(_nw122).ArraySet1(_827_v28, 20)
		(_nw122).ArraySet1(_827_v28, 21)
		(_nw122).ArraySet1(_827_v28, 22)
		(_nw122).ArraySet1(_827_v28, 23)
		_829_v30 = _nw122
		var _835_v31 _dafny.Sequence
		_ = _835_v31
		_835_v31 = _dafny.SeqOf(p0)
		var _836_v32 _dafny.Map
		_ = _836_v32
		_836_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_835_v31).Cardinality()), p0)
		var _837_v33 _dafny.Set
		_ = _837_v33
		_837_v33 = _dafny.SetOf((func() bool {
			if (_836_v32).Contains(Companion_Default___.Fm0((_this).F24(), _828_v29, globalState)) {
				return (_836_v32).Get(Companion_Default___.Fm0((_this).F24(), _828_v29, globalState)).(bool)
			}
			return (_this).F24()
		})(), true, (_this).F24())
		var _pat_let_tv25 = _828_v29
		_ = _pat_let_tv25
		var _pat_let_tv26 = _828_v29
		_ = _pat_let_tv26
		var _rhs127 _dafny.Array = _829_v30
		_ = _rhs127
		var _rhs128 bool = !(p0)
		_ = _rhs128
		var _rhs129 _dafny.Int = func(_source14 D5) _dafny.Int {
			if _source14.Is_DC13() {
				var _838___mcc_h5 _dafny.Int = _source14.Get_().(D5_DC13).Cf20
				_ = _838___mcc_h5
				var _839___mcc_h6 bool = _source14.Get_().(D5_DC13).Cf21
				_ = _839___mcc_h6
				var _840___mcc_h7 bool = _source14.Get_().(D5_DC13).Cf22
				_ = _840___mcc_h7
				var _841_cf22 bool = _840___mcc_h7
				_ = _841_cf22
				var _842_cf21 bool = _839___mcc_h6
				_ = _842_cf21
				var _843_cf20 _dafny.Int = _838___mcc_h5
				_ = _843_cf20
				return _843_cf20
			} else if _source14.Is_DC14() {
				var _844___mcc_h8 bool = _source14.Get_().(D5_DC14).Cf23
				_ = _844___mcc_h8
				var _845_cf23 bool = _844___mcc_h8
				_ = _845_cf23
				return _pat_let_tv25
			} else {
				var _846___mcc_h9 _dafny.Set = _source14.Get_().(D5_DC12).Cf19
				_ = _846___mcc_h9
				var _847_cf19 _dafny.Set = _846___mcc_h9
				_ = _847_cf19
				return _pat_let_tv26
			}
		}(Companion_D5_.Create_DC13_((_dafny.Zero).Minus((_837_v33).Cardinality()), p0, false))
		_ = _rhs129
		var _rhs130 bool = p0
		_ = _rhs130
		var _rhs131 bool = (_this).F24()
		_ = _rhs131
		var _lhs93 *GlobalState = globalState
		_ = _lhs93
		var _lhs94 *GlobalState = globalState
		_ = _lhs94
		var _lhs95 *GlobalState = globalState
		_ = _lhs95
		_829_v30 = _rhs127
		_lhs93.F13 = _rhs128
		_lhs94.F7 = _rhs129
		_lhs95.F13 = _rhs130
		r0 = _rhs131
		var _848_v34 _dafny.Array
		_ = _848_v34
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_21
		var _nw123 _dafny.Array
		_ = _nw123
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw123 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) bool = (func(_849_p0 bool) func(_dafny.Int) bool {
				return func(_850_i9 _dafny.Int) bool {
					return _849_p0
				}
			})(p0)
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw123 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw123).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw123).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_848_v34 = _nw123
		for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_848_v34), 0))); ; {
			_guard_loop_2, _ok69 := _iter69()
			if !_ok69 {
				break
			}
			var _851_i8 _dafny.Int
			_851_i8 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_851_i8).Sign() != -1) && ((_851_i8).Cmp(_dafny.ArrayLen((_848_v34), 0)) < 0)) {
				(_848_v34).ArraySet1((_dafny.MultiSetOf((_this).F24(), p0)).IsSubsetOf(_dafny.MultiSetOf(false, (_this).F24())), (_851_i8).Int())
			}
		}
		(globalState).F13 = (_this).F24()
		r0 = p0
		r1 = Companion_Default___.Fm34((func() _dafny.Int {
			if true {
				return _828_v29
			}
			return _828_v29
		})(), !(false), globalState)
		r2 = (_this).F24()
		return r0, r1, r2
	}
}
func (_this *C3) F35() _dafny.Map {
	{
		return _this._f35
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f23 _dafny.CodePoint
	_f24 bool
	_f25 _dafny.Array
	F34  _dafny.Map
	_f33 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
	_this._f25 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F34 = _dafny.EmptyMap
	_this._f33 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C4) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C4) F24() bool {
	return _this._f24
}
func (_this *C4) F25() _dafny.Array {
	return _this._f25
}
func (_this *C4) Ctor__(f33 _dafny.Int, f34 _dafny.Map, f25 _dafny.Array, f23 _dafny.CodePoint, f24 bool) {
	{
		(_this)._f33 = f33
		(_this).F34 = f34
		(_this)._f25 = f25
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C4) Fm7(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(_dafny.IntOfInt64(207), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(-927))).Cardinality()).Times((func() _dafny.Map {
			var _coll66 = _dafny.NewMapBuilder()
			_ = _coll66
			for _iter70 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()))).Elements()); ; {
				_compr_66, _ok70 := _iter70()
				if !_ok70 {
					break
				}
				var _852_v0 _dafny.Map
				_852_v0 = interface{}(_compr_66).(_dafny.Map)
				if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()))).Contains(_852_v0) {
					_coll66.Add(_852_v0, (_this).F33())
				}
			}
			return _coll66.ToMap()
		}()).Cardinality()))
	}
}
func (_this *C4) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((func() D2 {
			if (_this).F24() {
				return Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("h")))
			}
			return Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("d")))
		})()).Dtor_cf12()
	}
}
func (_this *C4) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C4) Fm22(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C4) Fm23(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C4) M3(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		if p0 {
			var _853_v0 _dafny.Sequence
			_ = _853_v0
			_853_v0 = _dafny.SeqOf(!(p0), p0)
			var _854_v1 _dafny.MultiSet
			_ = _854_v1
			_854_v1 = _dafny.MultiSetOf(!((_this).F24()), p0, (_853_v0).Select((Companion_Default___.SafeIndex((_this).F33(), _dafny.IntOfUint32((_853_v0).Cardinality()))).Uint32()).(bool))
			var _855_v2 *C2
			_ = _855_v2
			var _nw124 *C2 = New_C2_()
			_ = _nw124
			_nw124.Ctor__((_854_v1).Intersection(_dafny.MultiSetOf((_this).F24(), (_this).F24())), (_this).F25(), _dafny.CodePoint('c'), Companion_Default___.Fm2(globalState))
			_855_v2 = _nw124
			var _856_v3 _dafny.Set
			_ = _856_v3
			_856_v3 = _dafny.SetOf(_dafny.IntOfInt64(315))
			(globalState).F13 = (_856_v3).IsSubsetOf(_856_v3)
			var _857_v4 *C1
			_ = _857_v4
			var _nw125 *C1 = New_C1_()
			_ = _nw125
			_nw125.Ctor__(true, (_this).F24())
			_857_v4 = _nw125
			var _858_v5 _dafny.Sequence
			_ = _858_v5
			_858_v5 = _dafny.UnicodeSeqOfUtf8Bytes("c")
			var _859_v6 _dafny.Map
			_ = _859_v6
			_859_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_857_v4).F39(), _858_v5)
			var _source15 D2 = Companion_D2_.Create_DC6_(_859_v6)
			_ = _source15
			if _source15.Is_DC7() {
				var _860___mcc_h0 bool = _source15.Get_().(D2_DC7).Cf13
				_ = _860___mcc_h0
				var _861___mcc_h1 bool = _source15.Get_().(D2_DC7).Cf14
				_ = _861___mcc_h1
				var _862_cf14 bool = _861___mcc_h1
				_ = _862_cf14
				var _863_cf13 bool = _860___mcc_h0
				_ = _863_cf13
				var _864_v7 _dafny.Sequence
				_ = _864_v7
				_864_v7 = _dafny.SeqOf((_this).F33())
				var _865_v8 _dafny.Sequence
				_ = _865_v8
				_865_v8 = _dafny.SeqOf(_864_v7)
				var _866_v9 _dafny.Map
				_ = _866_v9
				_866_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(43), _dafny.IntOfUint32((_865_v8).Cardinality()))
				var _867_v11 _dafny.Map
				_ = _867_v11
				_867_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_857_v4).F39(), p0)
				var _868_v12 _dafny.Set
				_ = _868_v12
				_868_v12 = _dafny.SetOf(Companion_D6_.Create_DC16_(_867_v11, (_857_v4).F39(), (_857_v4).F39()))
				var _rhs132 bool = _862_cf14
				_ = _rhs132
				var _rhs133 _dafny.Sequence = _dafny.SeqOf(_862_cf14, (_857_v4).F39())
				_ = _rhs133
				var _rhs134 _dafny.Map = func() _dafny.Map {
					var _coll67 = _dafny.NewMapBuilder()
					_ = _coll67
					for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(452), _dafny.IntOfInt64(718))); ; {
						_compr_67, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _869_v10 _dafny.Int
						_869_v10 = interface{}(_compr_67).(_dafny.Int)
						if ((_dafny.IntOfInt64(452)).Cmp(_869_v10) <= 0) && ((_869_v10).Cmp(_dafny.IntOfInt64(718)) < 0) {
							_coll67.Add(Companion_Default___.SafeDivisionInt(_869_v10, (_this).F33()), ((_dafny.SetOf(_863_cf13, (_857_v4).F39())).Cardinality()).Plus((_856_v3).Cardinality()))
						}
					}
					return _coll67.ToMap()
				}()
				_ = _rhs134
				var _rhs135 _dafny.Int = ((_this).F33()).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _868_v12)).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(662))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}(func(_870_i0 _dafny.Int) _dafny.Int {
					return (_this).F33()
				}))).Cardinality())))
				_ = _rhs135
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				_863_cf13 = _rhs132
				_853_v0 = _rhs133
				_866_v9 = _rhs134
				_lhs96.F7 = _rhs135
				var _871_v13 D5
				_ = _871_v13
				_871_v13 = Companion_D5_.Create_DC13_((_this).F33(), (_857_v4).F39(), (_857_v4).F39())
				var _872_v14 _dafny.MultiSet
				_ = _872_v14
				_872_v14 = _dafny.MultiSetOf(_871_v13, _871_v13)
				var _873_v15 _dafny.MultiSet
				_ = _873_v15
				_873_v15 = _dafny.MultiSetOf(_858_v5)
				var _874_v16 _dafny.Array
				_ = _874_v16
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_22
				var _nw126 _dafny.Array
				_ = _nw126
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw126 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_875_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_876_i1 _dafny.Int) _dafny.Sequence {
							return _875_v0
						}
					})(_853_v0)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw126 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw126).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw126).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_874_v16 = _nw126
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_874_v16), 0))
				_ = _index114
				(_874_v16).ArraySet1((_855_v2).Fm25(_863_cf13, globalState), (_index114).Int())
				var _877_v17 _dafny.MultiSet
				_ = _877_v17
				_877_v17 = _dafny.MultiSetOf((_this).F33())
				var _878_v18 _dafny.Set
				_ = _878_v18
				_878_v18 = _dafny.SetOf(_877_v17, _877_v17, Companion_Default___.Fm42((_this).F24(), _dafny.IntOfUint32((_858_v5).Cardinality()), globalState))
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_874_v16), 0))
				_ = _index115
				var _rhs136 _dafny.MultiSet = (_872_v14).Union(_872_v14)
				_ = _rhs136
				var _rhs137 _dafny.MultiSet = _873_v15
				_ = _rhs137
				var _rhs138 bool = (_857_v4).Fm32(_863_cf13, (_this).F33(), globalState)
				_ = _rhs138
				var _rhs139 bool = !((_878_v18).IsSubsetOf(func() _dafny.Set {
					var _coll68 = _dafny.NewBuilder()
					_ = _coll68
					for _iter72 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_877_v17, (_this).F33())).Keys().Elements()); ; {
						_compr_68, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _879_v19 _dafny.MultiSet
						_879_v19 = interface{}(_compr_68).(_dafny.MultiSet)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_877_v17, (_this).F33())).Contains(_879_v19) {
							_coll68.Add(_879_v19)
						}
					}
					return _coll68.ToSet()
				}()))
				_ = _rhs139
				var _rhs140 _dafny.Sequence = _853_v0
				_ = _rhs140
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				var _lhs99 _dafny.Array = _874_v16
				_ = _lhs99
				var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_874_v16), 0))
				_ = _lhs100
				_872_v14 = _rhs136
				_873_v15 = _rhs137
				_lhs97.F13 = _rhs138
				_lhs98.F13 = _rhs139
				(_lhs99).ArraySet1(_rhs140, (_lhs100).Int())
				var _880_v20 _dafny.Array
				_ = _880_v20
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw127
				_880_v20 = _nw127
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_880_v20), 0))
				_ = _index116
				(_880_v20).ArraySet1((_857_v4).F39(), (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_880_v20), 0))
				_ = _index117
				var _rhs141 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F33(), (_dafny.Zero).Minus(((_dafny.Zero).Minus((_this).F33())).Minus((_this).F33())))
				_ = _rhs141
				var _rhs142 bool = (_857_v4).F39()
				_ = _rhs142
				var _lhs101 *GlobalState = globalState
				_ = _lhs101
				var _lhs102 _dafny.Array = _880_v20
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_880_v20), 0))
				_ = _lhs103
				_lhs101.F7 = _rhs141
				(_lhs102).ArraySet1(_rhs142, (_lhs103).Int())
				r0 = true
			} else if _source15.Is_DC6() {
				var _881___mcc_h2 _dafny.Map = _source15.Get_().(D2_DC6).Cf12
				_ = _881___mcc_h2
				var _882_cf12 _dafny.Map = _881___mcc_h2
				_ = _882_cf12
				r2 = p0
				var _883_v21 _dafny.Array
				_ = _883_v21
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw128
				_883_v21 = _nw128
				var _884_v23 _dafny.Map
				_ = _884_v23
				_884_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() _dafny.Set {
					var _coll69 = _dafny.NewBuilder()
					_ = _coll69
					for _iter73 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(643), _dafny.IntOfInt64(65))); ; {
						_compr_69, _ok73 := _iter73()
						if !_ok73 {
							break
						}
						var _885_v22 _dafny.Int
						_885_v22 = interface{}(_compr_69).(_dafny.Int)
						if ((_dafny.IntOfInt64(643)).Cmp(_885_v22) <= 0) && ((_885_v22).Cmp(_dafny.IntOfInt64(65)) < 0) {
							_coll69.Add((_885_v22).Minus((_this).F33()))
						}
					}
					return _coll69.ToSet()
				}()).Cardinality())
				var _886_v24 _dafny.Map
				_ = _886_v24
				_886_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_884_v23).Cardinality(), _854_v1)
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_883_v21), 0))
				_ = _index118
				(_883_v21).ArraySet1(_886_v24, (_index118).Int())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_883_v21), 0))
				_ = _index119
				(_883_v21).ArraySet1(_886_v24, (_index119).Int())
				var _887_v25 _dafny.Map
				_ = _887_v25
				_887_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(148), p0)
				(globalState).F13 = (_853_v0).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F33(), (_887_v25).Cardinality()), _dafny.IntOfUint32((_853_v0).Cardinality()))).Uint32()).(bool)
				r2 = !(Companion_Default___.Fm2(globalState))
			} else {
				var _888___mcc_h3 D2 = _source15.Get_().(D2_DC8).Cf15
				_ = _888___mcc_h3
				var _889_cf15 D2 = _888___mcc_h3
				_ = _889_cf15
				r2 = (_857_v4).F39()
				var _890_v27 D4
				_ = _890_v27
				_890_v27 = Companion_D4_.Create_DC11_((func() _dafny.Map {
					var _coll70 = _dafny.NewMapBuilder()
					_ = _coll70
					for _iter74 := _dafny.Iterate((_dafny.SeqOf(_this.F23(), p1)).Elements()); ; {
						_compr_70, _ok74 := _iter74()
						if !_ok74 {
							break
						}
						var _891_v26 _dafny.CodePoint
						_891_v26 = interface{}(_compr_70).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_this.F23(), p1), _891_v26) {
							_coll70.Add(_891_v26, (_857_v4).F39())
						}
					}
					return _coll70.ToMap()
				}()).Cardinality())
				var _892_v28 D0
				_ = _892_v28
				_892_v28 = Companion_D0_.Create_DC2_((_857_v4).F39(), (_890_v27).Dtor_cf18(), (_857_v4).F39())
				var _pat_let_tv27 = p0
				_ = _pat_let_tv27
				var _pat_let_tv28 = _857_v4
				_ = _pat_let_tv28
				(globalState).F7 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_858_v5, Companion_Default___.Fm1(true, globalState)), func(_pat_let37_0 D0) D0 {
					return func(_896_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let41_0 bool) D0 {
							return func(_897_dt__update_hcf4_h1 bool) D0 {
								return func(_pat_let42_0 bool) D0 {
									return func(_898_dt__update_hcf2_h0 bool) D0 {
										return Companion_D0_.Create_DC2_(_898_dt__update_hcf2_h0, (_896_dt__update__tmp_h1).Dtor_cf3(), _897_dt__update_hcf4_h1)
									}(_pat_let42_0)
								}((_pat_let_tv28).F39())
							}(_pat_let41_0)
						}((_this).F24())
					}(_pat_let37_0)
				}(func(_pat_let38_0 D0) D0 {
					return func(_893_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let39_0 bool) D0 {
							return func(_894_dt__update_hcf4_h0 bool) D0 {
								return func(_pat_let40_0 _dafny.Int) D0 {
									return func(_895_dt__update_hcf3_h0 _dafny.Int) D0 {
										return Companion_D0_.Create_DC2_((_893_dt__update__tmp_h0).Dtor_cf2(), _895_dt__update_hcf3_h0, _894_dt__update_hcf4_h0)
									}(_pat_let40_0)
								}(_dafny.IntOfInt64(795))
							}(_pat_let39_0)
						}(_pat_let_tv27)
					}(_pat_let38_0)
				}(_892_v28)))).Cardinality()
				r0 = (_857_v4).F39()
				var _899_v29 *C2
				_ = _899_v29
				var _nw129 *C2 = New_C2_()
				_ = _nw129
				_nw129.Ctor__((_854_v1).Update((_857_v4).F39(), Companion_Default___.Abs((_this).F33())), (_this).F25(), _dafny.CodePoint('f'), _dafny.Companion_Sequence_.IsProperPrefixOf(_858_v5, _dafny.UnicodeSeqOfUtf8Bytes("i")))
				_899_v29 = _nw129
			}
			r0 = (_857_v4).F39()
		} else {
			var _900_v30 _dafny.Array
			_ = _900_v30
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw130
			_900_v30 = _nw130
			var _901_v31 _dafny.Array
			_ = _901_v31
			var _nwElement0_28 _dafny.Array = _900_v30
			_ = _nwElement0_28
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(9))
			_ = _nw131
			(_nw131).ArraySet1(_nwElement0_28, 0)
			(_nw131).ArraySet1(_900_v30, 1)
			(_nw131).ArraySet1(_900_v30, 2)
			(_nw131).ArraySet1(_900_v30, 3)
			(_nw131).ArraySet1(_900_v30, 4)
			(_nw131).ArraySet1(_900_v30, 5)
			(_nw131).ArraySet1(_900_v30, 6)
			(_nw131).ArraySet1(_900_v30, 7)
			(_nw131).ArraySet1(_900_v30, 8)
			_901_v31 = _nw131
			_901_v31 = _901_v31
			var _902_v32 _dafny.Array
			_ = _902_v32
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw132
			_902_v32 = _nw132
			var _903_v33 _dafny.Sequence
			_ = _903_v33
			_903_v33 = _dafny.UnicodeSeqOfUtf8Bytes("hhp")
			var _904_v34 _dafny.MultiSet
			_ = _904_v34
			_904_v34 = _dafny.MultiSetOf((_this).F33(), _dafny.IntOfUint32((_903_v33).Cardinality()))
			var _905_v35 _dafny.MultiSet
			_ = _905_v35
			_905_v35 = _dafny.MultiSetOf(!(p0), (_this).F24(), (_this).F24(), (_this).F24())
			var _rhs143 _dafny.Array = _902_v32
			_ = _rhs143
			var _rhs144 _dafny.Int = (_this).F33()
			_ = _rhs144
			var _rhs145 _dafny.Array = _900_v30
			_ = _rhs145
			var _rhs146 _dafny.Int = Companion_Default___.SafeModuloInt(((func() _dafny.Int {
				if (_904_v34).Contains((_this).F33()) {
					return (_904_v34).Multiplicity((_this).F33())
				}
				return _dafny.IntOfInt64(-963)
			})()).Times((_905_v35).Cardinality()), (_this).F33())
			_ = _rhs146
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			_lhs104.F22 = _rhs143
			r1 = _rhs144
			_900_v30 = _rhs145
			r1 = _rhs146
			var _906_v36 _dafny.Map
			_ = _906_v36
			_906_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F33())
			_906_v36 = (_906_v36).Update(p1, _dafny.IntOfInt64(-924))
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_902_v32), 0))
			_ = _index120
			(_902_v32).ArraySet1((_this).F33(), (_index120).Int())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_902_v32), 0))
			_ = _index121
			(_902_v32).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_903_v33, _dafny.UnicodeSeqOfUtf8Bytes("fb"))).Cardinality()), (_index121).Int())
			var _907_v37 _dafny.Map
			_ = _907_v37
			_907_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if !(false) {
					return p0
				}
				return (_this).F24()
			})(), (_this).F33())
			var _rhs147 _dafny.CodePoint = p1
			_ = _rhs147
			var _rhs148 _dafny.Map = _907_v37
			_ = _rhs148
			var _rhs149 _dafny.Int = ((_this).F33()).Plus((_this).F33())
			_ = _rhs149
			var _rhs150 bool = p0
			_ = _rhs150
			var _rhs151 bool = (_this).F24()
			_ = _rhs151
			var _lhs105 *C4 = _this
			_ = _lhs105
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			_lhs105.F23_set_(_rhs147)
			_907_v37 = _rhs148
			_lhs106.F7 = _rhs149
			r2 = _rhs150
			_lhs107.F13 = _rhs151
		}
		if (p0) || (p0) {
			var _908_v38 _dafny.Map
			_ = _908_v38
			_908_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), (_this).F24())
			var _909_v39 _dafny.Sequence
			_ = _909_v39
			_909_v39 = _dafny.SeqOf(_908_v38, _908_v38)
			r0 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_909_v39, _909_v39), _909_v39)
			var _910_v40 _dafny.Sequence
			_ = _910_v40
			_910_v40 = _dafny.SeqOf((_this).F33())
			var _911_v41 _dafny.MultiSet
			_ = _911_v41
			_911_v41 = _dafny.MultiSetOf(_910_v40)
			var _912_v42 _dafny.Array
			_ = _912_v42
			var _nwElement0_29 _dafny.Int = (_this).F33()
			_ = _nwElement0_29
			var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(2))
			_ = _nw133
			(_nw133).ArraySet1(_nwElement0_29, 0)
			(_nw133).ArraySet1((_this).F33(), 1)
			_912_v42 = _nw133
			var _913_v43 _dafny.Map
			_ = _913_v43
			_913_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_911_v41, _912_v42)
			_913_v43 = (_913_v43).Update(_911_v41, (func() _dafny.Array {
				if (_this).F24() {
					return _912_v42
				}
				return _912_v42
			})())
			var _914_v44 _dafny.Sequence
			_ = _914_v44
			_914_v44 = _dafny.UnicodeSeqOfUtf8Bytes("jp")
			var _915_v45 _dafny.MultiSet
			_ = _915_v45
			_915_v45 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}(func(_916_i2 _dafny.Int) _dafny.CodePoint {
				return _this.F23()
			})))
			var _917_v46 _dafny.Map
			_ = _917_v46
			_917_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v44, _915_v45)
			_917_v46 = (_917_v46).Update(_914_v44, _915_v45)
			var _918_v47 _dafny.Sequence
			_ = _918_v47
			_918_v47 = _dafny.SeqOf(false, p0, !((_this).F24()), (_this).F24())
			var _919_v48 _dafny.Map
			_ = _919_v48
			_919_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
			var _920_v49 _dafny.Array
			_ = _920_v49
			var _nwElement0_30 bool = (_this).F24()
			_ = _nwElement0_30
			var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(28))
			_ = _nw134
			(_nw134).ArraySet1(_nwElement0_30, 0)
			(_nw134).ArraySet1(p0, 1)
			(_nw134).ArraySet1(p0, 2)
			(_nw134).ArraySet1(p0, 3)
			(_nw134).ArraySet1(p0, 4)
			(_nw134).ArraySet1((_this).F24(), 5)
			(_nw134).ArraySet1(p0, 6)
			(_nw134).ArraySet1(p0, 7)
			(_nw134).ArraySet1(p0, 8)
			(_nw134).ArraySet1((_this).F24(), 9)
			(_nw134).ArraySet1(p0, 10)
			(_nw134).ArraySet1(!((_918_v47).Select((Companion_Default___.SafeIndex((_this).F33(), _dafny.IntOfUint32((_918_v47).Cardinality()))).Uint32()).(bool)), 11)
			(_nw134).ArraySet1(p0, 12)
			(_nw134).ArraySet1((_this).F24(), 13)
			(_nw134).ArraySet1((_this).F24(), 14)
			(_nw134).ArraySet1(false, 15)
			(_nw134).ArraySet1((_this).F24(), 16)
			(_nw134).ArraySet1(p0, 17)
			(_nw134).ArraySet1(true, 18)
			(_nw134).ArraySet1((func() bool {
				if (_919_v48).Contains(true) {
					return (_919_v48).Get(true).(bool)
				}
				return true
			})(), 19)
			(_nw134).ArraySet1((_this).F24(), 20)
			(_nw134).ArraySet1((_918_v47).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F33()), _dafny.IntOfUint32((_918_v47).Cardinality()))).Uint32()).(bool), 21)
			(_nw134).ArraySet1((_this).Fm4((_this).F33(), _918_v47, globalState), 22)
			(_nw134).ArraySet1(p0, 23)
			(_nw134).ArraySet1(false, 24)
			(_nw134).ArraySet1((_this).F24(), 25)
			(_nw134).ArraySet1((_this).F24(), 26)
			(_nw134).ArraySet1(true, 27)
			_920_v49 = _nw134
			var _921_v50 _dafny.Map
			_ = _921_v50
			_921_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_920_v49, _this.F23())
			var _922_v51 _dafny.Set
			_ = _922_v51
			_922_v51 = _dafny.SetOf((_914_v44).Select((Companion_Default___.SafeIndex((_this).F33(), _dafny.IntOfUint32((_914_v44).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F23())
			var _923_v52 _dafny.MultiSet
			_ = _923_v52
			_923_v52 = _dafny.MultiSetOf((_this).F24(), p0)
			var _924_v53 D0
			_ = _924_v53
			_924_v53 = Companion_D0_.Create_DC0_(_923_v52)
			var _925_v54 D0
			_ = _925_v54
			_925_v54 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(_924_v53)))
			var _926_v55 _dafny.Array
			_ = _926_v55
			var _nwElement0_31 D0 = Companion_D0_.Create_DC3_(_924_v53)
			_ = _nwElement0_31
			var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(8))
			_ = _nw135
			(_nw135).ArraySet1(_nwElement0_31, 0)
			(_nw135).ArraySet1(_925_v54, 1)
			(_nw135).ArraySet1(_925_v54, 2)
			(_nw135).ArraySet1(_925_v54, 3)
			(_nw135).ArraySet1(_925_v54, 4)
			(_nw135).ArraySet1(Companion_D0_.Create_DC3_(_924_v53), 5)
			(_nw135).ArraySet1(_925_v54, 6)
			(_nw135).ArraySet1(_925_v54, 7)
			_926_v55 = _nw135
			var _927_v56 _dafny.Map
			_ = _927_v56
			_927_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_926_v55, _dafny.SetOf(_this.F23(), p1))
			var _928_v57 *C3
			_ = _928_v57
			var _nw136 *C3 = New_C3_()
			_ = _nw136
			_nw136.Ctor__(_921_v50, (_this).F25(), _this.F23(), !(((func() _dafny.Set {
				if (_927_v56).Contains(_926_v55) {
					return (_927_v56).Get(_926_v55).(_dafny.Set)
				}
				return _922_v51
			})()).IsProperSubsetOf(_922_v51)))
			_928_v57 = _nw136
			var _nw137 *C3 = New_C3_()
			_ = _nw137
			_nw137.Ctor__((_928_v57).F35(), (_this).F25(), _this.F23(), ((_dafny.SetOf(_912_v42, _912_v42)).Cardinality()).Cmp((_dafny.Zero).Minus((_this).F33())) == 0)
			_928_v57 = _nw137
			var _929_v58 _dafny.Set
			_ = _929_v58
			_929_v58 = _dafny.SetOf((_this).F33(), _dafny.IntOfUint32(((_928_v57).Fm7(globalState)).Cardinality()))
			var _930_v59 D5
			_ = _930_v59
			_930_v59 = Companion_D5_.Create_DC12_(_929_v58)
			var _931_v60 _dafny.Map
			_ = _931_v60
			_931_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), _this.F23())
			r0 = !(((_929_v58).Difference(_dafny.SetOf((_dafny.Zero).Minus((_919_v48).Cardinality()), (_this).F33(), (_this).F33(), (_this).F33(), (_931_v60).Cardinality()))).IsSubsetOf((_930_v59).Dtor_cf19()))
		} else {
			(globalState).F7 = (_this).F33()
			var _932_v61 _dafny.MultiSet
			_ = _932_v61
			_932_v61 = _dafny.MultiSetOf(!((_this).F24()))
			var _933_v62 *C2
			_ = _933_v62
			var _nw138 *C2 = New_C2_()
			_ = _nw138
			_nw138.Ctor__((_dafny.MultiSetOf((_this).F24(), p0, (_this).F24())).Intersection(_932_v61), (_this).F25(), (func() _dafny.CodePoint {
				if p0 {
					return _this.F23()
				}
				return _this.F23()
			})(), (_932_v61).IsDisjointFrom(_932_v61))
			_933_v62 = _nw138
			var _934_v63 _dafny.Array
			_ = _934_v63
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw139
			_934_v63 = _nw139
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_934_v63), 0))
			_ = _index122
			(_934_v63).ArraySet1(p0, (_index122).Int())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_934_v63), 0))
			_ = _index123
			(_934_v63).ArraySet1(((_this).F33()).Cmp((_dafny.Zero).Minus((_this).F33())) == 0, (_index123).Int())
			var _935_v64 _dafny.Sequence
			_ = _935_v64
			_935_v64 = _dafny.SeqOf(p0)
			(globalState).F13 = (_935_v64).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F33(), (_this).F33()), _dafny.IntOfUint32((_935_v64).Cardinality()))).Uint32()).(bool)
			var _936_v65 _dafny.Int
			_ = _936_v65
			var _out34 _dafny.Int
			_ = _out34
			_out34 = (_this).M10((_934_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_934_v63), 0))).Int()).(bool), globalState)
			_936_v65 = _out34
		}
		(globalState).F13 = Companion_Default___.Fm2(globalState)
		var _937_v66 _dafny.Sequence
		_ = _937_v66
		_937_v66 = _dafny.SeqOf((_this).F24(), (_this).F24(), p0, !(p0), p0)
		var _938_v67 _dafny.Map
		_ = _938_v67
		_938_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _937_v66)
		(globalState).F13 = !((_938_v67).Contains((_this).F24()))
		var _hi3 _dafny.Int = Companion_Default___.Fm0((_this).F24(), (_this).F33(), globalState)
		_ = _hi3
		for _939_i3 := (_this).F33(); _939_i3.Cmp(_hi3) < 0; _939_i3 = _939_i3.Plus(_dafny.One) {
			if (_939_i3).Cmp(_939_i3) > 0 {
				(globalState).F7 = _939_i3
				var _940_v68 _dafny.Set
				_ = _940_v68
				_940_v68 = _dafny.SetOf((_this).F24(), (_this).F24())
				var _941_v69 _dafny.Sequence
				_ = _941_v69
				_941_v69 = _dafny.SeqOf((_940_v68).Difference(_940_v68))
				var _942_v70 _dafny.Map
				_ = _942_v70
				_942_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_940_v68, _940_v68), _941_v69))
				var _943_v71 _dafny.Set
				_ = _943_v71
				_943_v71 = _dafny.SetOf(_dafny.IntOfInt64(999), _939_i3)
				var _944_v72 _dafny.MultiSet
				_ = _944_v72
				_944_v72 = _dafny.MultiSetOf(_943_v71)
				var _rhs152 bool = !(!((_this).F24()))
				_ = _rhs152
				var _rhs153 _dafny.Sequence = (func() _dafny.Sequence {
					if (_942_v70).Contains(false) {
						return (_942_v70).Get(false).(_dafny.Sequence)
					}
					return _941_v69
				})()
				_ = _rhs153
				var _rhs154 _dafny.Int = (_dafny.Zero).Minus(_939_i3)
				_ = _rhs154
				var _rhs155 bool = (_this).F24()
				_ = _rhs155
				var _rhs156 _dafny.Int = ((_944_v72).Cardinality()).Times((_this).F33())
				_ = _rhs156
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				_lhs108.F13 = _rhs152
				_941_v69 = _rhs153
				_lhs109.F7 = _rhs154
				_lhs110.F13 = _rhs155
				_lhs111.F7 = _rhs156
				(_this).F34 = _this.F34
				var _945_v73 _dafny.Array
				_ = _945_v73
				var _nw140 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw140
				_945_v73 = _nw140
				var _946_v74 _dafny.Map
				_ = _946_v74
				_946_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F24())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_945_v73), 0))
				_ = _index124
				(_945_v73).ArraySet1((_this).Fm23((_this).F33(), _946_v74, _937_v66, globalState), (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_945_v73), 0))
				_ = _index125
				(_945_v73).ArraySet1((_this).Fm23(_939_i3, _946_v74, _dafny.Companion_Sequence_.Concatenate(_937_v66, _937_v66), globalState), (_index125).Int())
				var _947_v75 _dafny.Sequence
				_ = _947_v75
				_947_v75 = _dafny.SeqOf((_this).F33(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality())))
				var _948_v76 _dafny.MultiSet
				_ = _948_v76
				_948_v76 = _dafny.MultiSetOf((_this).F24(), (_this).F24(), false)
				var _949_v77 _dafny.Map
				_ = _949_v77
				_949_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_947_v75, ((func() _dafny.Int {
					if (_948_v76).Contains(false) {
						return (_948_v76).Multiplicity(false)
					}
					return (_this).F33()
				})()).Minus(_939_i3))
				_949_v77 = (_949_v77).Update(_947_v75, _939_i3)
			} else {
				var _950_v78 _dafny.Sequence
				_ = _950_v78
				_950_v78 = _dafny.UnicodeSeqOfUtf8Bytes("jctwmpljq")
				_950_v78 = _dafny.Companion_Sequence_.Concatenate(_950_v78, _dafny.Companion_Sequence_.Update(_950_v78, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0, (_this).F24(), true, p0)).Cardinality())), _dafny.IntOfUint32((_950_v78).Cardinality()))).Uint32(), _this.F23()))
				var _951_v79 _dafny.Array
				_ = _951_v79
				var _nw141 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw141
				_951_v79 = _nw141
				_951_v79 = (func() _dafny.Array {
					if p0 {
						return _951_v79
					}
					return _951_v79
				})()
				var _952_v80 _dafny.Array
				_ = _952_v80
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_23
				var _nw142 _dafny.Array
				_ = _nw142
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw142 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) D0 = (func(_953_p0 bool) func(_dafny.Int) D0 {
						return func(_954_i4 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_dafny.MultiSetOf(!(_953_p0)))
						}
					})(p0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw142 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw142).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw142).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_952_v80 = _nw142
				var _955_v81 _dafny.MultiSet
				_ = _955_v81
				_955_v81 = _dafny.MultiSetOf((_this).F24())
				var _956_v82 D0
				_ = _956_v82
				_956_v82 = Companion_D0_.Create_DC0_(_955_v81)
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_952_v80), 0))
				_ = _index126
				(_952_v80).ArraySet1(_956_v82, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_952_v80), 0))
				_ = _index127
				(_952_v80).ArraySet1(Companion_Default___.Fm43(p0, (_this).F24(), p0, globalState), (_index127).Int())
				var _957_v83 _dafny.Sequence
				_ = _957_v83
				_957_v83 = _dafny.SeqOf(_950_v78, _950_v78)
				var _rhs157 bool = (_this).F24()
				_ = _rhs157
				var _rhs158 bool = (_939_i3).Cmp(_939_i3) > 0
				_ = _rhs158
				var _rhs159 _dafny.Sequence = _957_v83
				_ = _rhs159
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				_lhs112.F13 = _rhs157
				r0 = _rhs158
				_957_v83 = _rhs159
				var _958_v84 _dafny.Map
				_ = _958_v84
				_958_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(43))
				var _959_v85 _dafny.Map
				_ = _959_v85
				_959_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _958_v84)
				var _960_v86 _dafny.Map
				_ = _960_v86
				_960_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), ((func() _dafny.Map {
					if (_959_v85).Contains(_this.F23()) {
						return (_959_v85).Get(_this.F23()).(_dafny.Map)
					}
					return _958_v84
				})()).Contains(!(p0)))
				r1 = (_960_v86).Cardinality()
			}
			var _961_v87 D5
			_ = _961_v87
			_961_v87 = Companion_D5_.Create_DC14_(p0)
			var _962_v88 _dafny.Sequence
			_ = _962_v88
			_962_v88 = _dafny.SeqOf(_961_v87, _961_v87)
			var _963_v89 _dafny.Map
			_ = _963_v89
			_963_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _962_v88)
			_963_v89 = (_963_v89).Update((_this).F24(), _962_v88)
			var _964_v90 *C1
			_ = _964_v90
			var _nw143 *C1 = New_C1_()
			_ = _nw143
			_nw143.Ctor__((_this).F24(), true)
			_964_v90 = _nw143
			var _965_v91 *C1
			_ = _965_v91
			_965_v91 = _964_v90
			var _966_v92 _dafny.Map
			_ = _966_v92
			_966_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_this).F24() {
					return _939_i3
				}
				return (_this).F33()
			})(), _965_v91)
			_966_v92 = (_966_v92).Update(Companion_Default___.SafeDivisionInt((_this).F33(), (_this).F33()), _965_v91)
			var _967_v93 _dafny.Sequence
			_ = _967_v93
			_967_v93 = _dafny.UnicodeSeqOfUtf8Bytes("nvxmib")
			(globalState).F7 = Companion_Default___.SafeModuloInt(_939_i3, _dafny.IntOfUint32((_967_v93).Cardinality()))
		}
		var _968_v94 _dafny.Array
		_ = _968_v94
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_24
		var _nw144 _dafny.Array
		_ = _nw144
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw144 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) bool = func(_969_i5 _dafny.Int) bool {
				return false
			}
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw144 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw144).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw144).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_968_v94 = _nw144
		var _970_v95 _dafny.Map
		_ = _970_v95
		_970_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_968_v94, _this.F23())
		var _971_v96 *C3
		_ = _971_v96
		var _nw145 *C3 = New_C3_()
		_ = _nw145
		_nw145.Ctor__(_970_v95, (_this).F25(), p1, p0)
		_971_v96 = _nw145
		r0 = p0
		r1 = (_this).F33()
		var _972_v98 _dafny.MultiSet
		_ = _972_v98
		_972_v98 = _dafny.MultiSetOf((_this).F33())
		var _973_v99 _dafny.Set
		_ = _973_v99
		_973_v99 = _dafny.SetOf(_968_v94)
		var _974_v100 _dafny.Map
		_ = _974_v100
		_974_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll71 = _dafny.NewMapBuilder()
			_ = _coll71
			for _iter75 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_972_v98).Contains((_this).F33()) {
					return (_972_v98).Multiplicity((_this).F33())
				}
				return (_this).F33()
			})(), _dafny.IntOfInt64(240))).Keys().Elements()); ; {
				_compr_71, _ok75 := _iter75()
				if !_ok75 {
					break
				}
				var _975_v97 _dafny.Int
				_975_v97 = interface{}(_compr_71).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_972_v98).Contains((_this).F33()) {
						return (_972_v98).Multiplicity((_this).F33())
					}
					return (_this).F33()
				})(), _dafny.IntOfInt64(240))).Contains(_975_v97) {
					_coll71.Add((_975_v97).Times((_this).F33()), (_this).F33())
				}
			}
			return _coll71.ToMap()
		}(), _973_v99)
		var _976_v101 _dafny.Sequence
		_ = _976_v101
		_976_v101 = _dafny.UnicodeSeqOfUtf8Bytes("ktaasd")
		var _977_v102 _dafny.Map
		_ = _977_v102
		_977_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_976_v101).Cardinality()), (_this).F33())
		r2 = (_973_v99).IsSubsetOf((func() _dafny.Set {
			if (_974_v100).Contains(_977_v102) {
				return (_974_v100).Get(_977_v102).(_dafny.Set)
			}
			return _973_v99
		})())
		return r0, r1, r2
	}
}
func (_this *C4) M10(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _978_v0 _dafny.Map
		_ = _978_v0
		_978_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F33(), ((_this).F33()).Cmp(_dafny.IntOfInt64(237)) >= 0)
		_978_v0 = (_978_v0).Update((_this).F33(), (func() bool {
			if p0 {
				return (_this).F24()
			}
			return p0
		})())
		var _979_v1 _dafny.Map
		_ = _979_v1
		_979_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F24())
		_979_v1 = (_979_v1).Update(!((_this).F24()), (_this).F24())
		(globalState).F13 = ((func() bool {
			if false {
				return p0
			}
			return p0
		})()) && (p0)
		var _hi4 _dafny.Int = (_this).F33()
		_ = _hi4
		for _980_i0 := (_this).F33(); _980_i0.Cmp(_hi4) < 0; _980_i0 = _980_i0.Plus(_dafny.One) {
			_979_v1 = _979_v1
			var _981_v2 _dafny.Array
			_ = _981_v2
			var _nw146 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw146
			_981_v2 = _nw146
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_981_v2), 0))
			_ = _index128
			(_981_v2).ArraySet1(true, (_index128).Int())
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_981_v2), 0))
			_ = _index129
			(_981_v2).ArraySet1((_this).F24(), (_index129).Int())
			(globalState).F13 = !(p0)
			var _982_v3 _dafny.Sequence
			_ = _982_v3
			_982_v3 = _dafny.SeqOf((_this).F24())
			var _983_v4 _dafny.MultiSet
			_ = _983_v4
			_983_v4 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_981_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_981_v2), 0))).Int()).(bool)), _982_v3))
			_983_v4 = _983_v4
		}
		var _984_i1 _dafny.Int
		_ = _984_i1
		_984_i1 = _dafny.Zero
		{
			for ((_this).F33()).Cmp((_this).F33()) == 0 {
				{
					if (_984_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_984_i1 = (_984_i1).Plus(_dafny.One)
					(globalState).F13 = Companion_Default___.Fm2(globalState)
					(globalState).F7 = (_this).F33()
					var _985_v5 _dafny.Map
					_ = _985_v5
					_985_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (_this).F33())
					var _986_v6 _dafny.Sequence
					_ = _986_v6
					_986_v6 = _dafny.SeqOf((_this).F33())
					var _987_v7 _dafny.Sequence
					_ = _987_v7
					_987_v7 = _986_v6
					var _988_v8 D0
					_ = _988_v8
					_988_v8 = Companion_D0_.Create_DC2_(!(!((_this).F24())), (func() _dafny.Int {
						if (_985_v5).Contains((_this).F33()) {
							return (_985_v5).Get((_this).F33()).(_dafny.Int)
						}
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((_987_v7).Cardinality()))
					})(), p0)
					var _989_v9 _dafny.Array
					_ = _989_v9
					var _len0_25 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_25
					var _nw147 _dafny.Array
					_ = _nw147
					if _len0_25.Cmp(_dafny.Zero) == 0 {
						_nw147 = _dafny.NewArray(_len0_25)
					} else {
						var _init25 func(_dafny.Int) _dafny.Int = func(_990_i2 _dafny.Int) _dafny.Int {
							return (_990_i2).Times((_this).F33())
						}
						_ = _init25
						var _element0_25 = _init25(_dafny.Zero)
						_ = _element0_25
						_nw147 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
						(_nw147).ArraySet1(_element0_25, 0)
						var _nativeLen0_25 = (_len0_25).Int()
						_ = _nativeLen0_25
						for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
							(_nw147).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
						}
					}
					_989_v9 = _nw147
					var _991_v10 _dafny.Map
					_ = _991_v10
					_991_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_988_v8, _989_v9)
					var _992_v11 _dafny.Sequence
					_ = _992_v11
					_992_v11 = _dafny.SeqOf(p0)
					_991_v10 = (_991_v10).Update(Companion_D0_.Create_DC2_(!(p0), _dafny.IntOfUint32((_992_v11).Cardinality()), (_this).F24()), _989_v9)
					var _993_v12 _dafny.Array
					_ = _993_v12
					var _len0_26 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_26
					var _nw148 _dafny.Array
					_ = _nw148
					if _len0_26.Cmp(_dafny.Zero) == 0 {
						_nw148 = _dafny.NewArray(_len0_26)
					} else {
						var _init26 func(_dafny.Int) _dafny.Sequence = (func(_994_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_995_i3 _dafny.Int) _dafny.Sequence {
								return _994_v6
							}
						})(_986_v6)
						_ = _init26
						var _element0_26 = _init26(_dafny.Zero)
						_ = _element0_26
						_nw148 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
						(_nw148).ArraySet1(_element0_26, 0)
						var _nativeLen0_26 = (_len0_26).Int()
						_ = _nativeLen0_26
						for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
							(_nw148).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
						}
					}
					_993_v12 = _nw148
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_993_v12), 0))
					_ = _index130
					(_993_v12).ArraySet1((_this).Fm7(globalState), (_index130).Int())
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_989_v9), 0))
					_ = _index131
					(_989_v9).ArraySet1(((_this).F33()).Plus((_this).F33()), (_index131).Int())
					var _996_v13 _dafny.Sequence
					_ = _996_v13
					_996_v13 = _dafny.UnicodeSeqOfUtf8Bytes("xfjecqwws")
					var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_993_v12), 0))
					_ = _index132
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_989_v9), 0))
					_ = _index133
					var _rhs160 _dafny.Sequence = _986_v6
					_ = _rhs160
					var _rhs161 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(498)).Minus((_dafny.Zero).Minus((_dafny.SetOf((_this).F33(), _dafny.IntOfUint32((_996_v13).Cardinality()))).Cardinality())), (_dafny.IntOfInt64(-59)).Plus((_this).F33()))
					_ = _rhs161
					var _rhs162 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F33(), (_this).F33())
					_ = _rhs162
					var _lhs113 _dafny.Array = _993_v12
					_ = _lhs113
					var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_993_v12), 0))
					_ = _lhs114
					var _lhs115 *GlobalState = globalState
					_ = _lhs115
					var _lhs116 _dafny.Array = _989_v9
					_ = _lhs116
					var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_989_v9), 0))
					_ = _lhs117
					(_lhs113).ArraySet1(_rhs160, (_lhs114).Int())
					_lhs115.F7 = _rhs161
					(_lhs116).ArraySet1(_rhs162, (_lhs117).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r0 = (_this).F33()
		var _997_v14 _dafny.MultiSet
		_ = _997_v14
		_997_v14 = _dafny.MultiSetOf(p0)
		r0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_997_v14).Contains((_this).F24()) {
				return (_997_v14).Multiplicity((_this).F24())
			}
			return (_this).F33()
		})(), (_this).F33())))
		return r0
	}
}
func (_this *C4) F33() _dafny.Int {
	{
		return _this._f33
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f32 _dafny.Array
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f32 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C5) Ctor__(f32 _dafny.Array) {
	{
		(_this)._f32 = f32
	}
}
func (_this *C5) Fm17(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfInt64(744)).Times(_dafny.IntOfInt64(-870))).Plus(((_dafny.MultiSetOf(_dafny.CodePoint('j'))).Intersection(_dafny.MultiSetOf(_dafny.CodePoint('v')))).Cardinality())
	}
}
func (_this *C5) M8(globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _998_v0 bool
		_ = _998_v0
		_998_v0 = true
		var _999_v1 _dafny.Sequence
		_ = _999_v1
		_999_v1 = _dafny.SeqOf(_998_v0, _998_v0, _998_v0)
		var _1000_v2 _dafny.Int
		_ = _1000_v2
		_1000_v2 = _dafny.IntOfInt64(987)
		var _1001_v3 _dafny.MultiSet
		_ = _1001_v3
		_1001_v3 = _dafny.MultiSetOf(_dafny.IntOfUint32((_999_v1).Cardinality()), _1000_v2)
		var _1002_v4 _dafny.Map
		_ = _1002_v4
		_1002_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_998_v0, _998_v0, _998_v0)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-624))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg73 _dafny.Int) interface{} {
				return coer73(arg73)
			}
		}(func(_1003_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})))
		(globalState).F13 = ((Companion_Default___.Fm18(globalState)).Difference(_1001_v3)).IsProperSubsetOf((_1001_v3).Update(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1002_v4).Contains(_1000_v2) {
				return (_1002_v4).Get(_1000_v2).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("rnfuhyqq")
		})()).Cardinality()), Companion_Default___.Abs(_1000_v2)))
		var _1004_v5 _dafny.Set
		_ = _1004_v5
		_1004_v5 = _dafny.SetOf(_dafny.IntOfInt64(164))
		_1004_v5 = (_1004_v5).Union(func() _dafny.Set {
			var _coll72 = _dafny.NewBuilder()
			_ = _coll72
			for _iter76 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(413), _dafny.IntOfInt64(860))); ; {
				_compr_72, _ok76 := _iter76()
				if !_ok76 {
					break
				}
				var _1005_v6 _dafny.Int
				_1005_v6 = interface{}(_compr_72).(_dafny.Int)
				if ((_dafny.IntOfInt64(413)).Cmp(_1005_v6) <= 0) && ((_1005_v6).Cmp(_dafny.IntOfInt64(860)) < 0) {
					_coll72.Add((_1005_v6).Minus(_1000_v2))
				}
			}
			return _coll72.ToSet()
		}())
		var _hi5 _dafny.Int = _1000_v2
		_ = _hi5
		for _1006_i1 := _1000_v2; _1006_i1.Cmp(_hi5) < 0; _1006_i1 = _1006_i1.Plus(_dafny.One) {
			if _998_v0 {
				var _1007_v7 _dafny.Sequence
				_ = _1007_v7
				_1007_v7 = _dafny.UnicodeSeqOfUtf8Bytes("qwdbajgpg")
				var _1008_v8 _dafny.Map
				_ = _1008_v8
				_1008_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v0, (func() _dafny.Int {
					if (_1001_v3).Contains((_1004_v5).Cardinality()) {
						return (_1001_v3).Multiplicity((_1004_v5).Cardinality())
					}
					return _1000_v2
				})())
				var _1009_v9 _dafny.CodePoint
				_ = _1009_v9
				_1009_v9 = _dafny.CodePoint('k')
				var _1010_v10 _dafny.Array
				_ = _1010_v10
				var _nwElement0_32 _dafny.Int = _1000_v2
				_ = _nwElement0_32
				var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(6))
				_ = _nw149
				(_nw149).ArraySet1(_nwElement0_32, 0)
				(_nw149).ArraySet1(Companion_Default___.Fm0(_998_v0, _1000_v2, globalState), 1)
				(_nw149).ArraySet1((_this).Fm17(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1007_v7, (Companion_Default___.SafeIndex((_1008_v8).Cardinality(), _dafny.IntOfUint32((_1007_v7).Cardinality()))).Uint32(), _1009_v9)).Cardinality()), _dafny.CodePoint('g'), Companion_Default___.Fm2(globalState), globalState), 2)
				(_nw149).ArraySet1(_1006_i1, 3)
				(_nw149).ArraySet1((_1006_i1).Minus(_1006_i1), 4)
				(_nw149).ArraySet1(Companion_Default___.Fm0(_998_v0, _1006_i1, globalState), 5)
				_1010_v10 = _nw149
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1010_v10), 0))
				_ = _index134
				(_1010_v10).ArraySet1(_dafny.IntOfInt64(454), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1010_v10), 0))
				_ = _index135
				(_1010_v10).ArraySet1(Companion_Default___.SafeDivisionInt((_this).Fm17(_1006_i1, _1009_v9, _998_v0, globalState), _1006_i1), (_index135).Int())
				(globalState).F7 = _1006_i1
				var _1011_v11 _dafny.Map
				_ = _1011_v11
				_1011_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v0, _998_v0)
				var _1012_v12 _dafny.Map
				_ = _1012_v12
				_1012_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v0, _dafny.Companion_Sequence_.Update(_999_v1, (Companion_Default___.SafeIndex(_1006_i1, _dafny.IntOfUint32((_999_v1).Cardinality()))).Uint32(), _998_v0)), (_1011_v11).Merge(Companion_Default___.Fm19(globalState)))
				var _1013_v13 _dafny.Map
				_ = _1013_v13
				_1013_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v0, _999_v1)
				_1012_v12 = (_1012_v12).Update((func() _dafny.Map {
					if _998_v0 {
						return _1013_v13
					}
					return _1013_v13
				})(), _1011_v11)
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen(((_this).F32()), 0))
				_ = _index136
				((_this).F32()).ArraySet1(_998_v0, (_index136).Int())
				var _1014_v14 _dafny.Sequence
				_ = _1014_v14
				_1014_v14 = _dafny.SeqOf((_1010_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1010_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-537))
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen(((_this).F32()), 0))
				_ = _index137
				((_this).F32()).ArraySet1((_dafny.MultiSetOf(_1014_v14, _1014_v14)).IsProperSubsetOf(_dafny.MultiSetOf(_1014_v14)), (_index137).Int())
				(globalState).F7 = (_1010_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1010_v10), 0))).Int()).(_dafny.Int)
			} else {
				var _1015_v15 _dafny.Array
				_ = _1015_v15
				var _out35 _dafny.Array
				_ = _out35
				_out35 = (_this).M9((_dafny.IntOfInt64(-814)).Cmp(_1000_v2) <= 0, _998_v0, globalState)
				_1015_v15 = _out35
				(globalState).F13 = _998_v0
				(globalState).F7 = _1000_v2
				var _1016_v16 _dafny.Sequence
				_ = _1016_v16
				_1016_v16 = _dafny.SeqOf(_1000_v2, _1000_v2)
				(globalState).F7 = Companion_Default___.Fm0(false, Companion_Default___.SafeDivisionInt(_1006_i1, _dafny.IntOfUint32((_1016_v16).Cardinality())), globalState)
				var _1017_v17 _dafny.Array
				_ = _1017_v17
				var _nwElement0_33 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-368))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}((func(_1018_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1019_i2 _dafny.Int) _dafny.Int {
						return _1018_i1
					}
				})(_1006_i1)))
				_ = _nwElement0_33
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(29))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_33, 0)
				(_nw150).ArraySet1(_dafny.SeqOf(_1006_i1, _1006_i1, _1000_v2), 1)
				(_nw150).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(182))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1020_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1021_i3 _dafny.Int) _dafny.Int {
						return _1020_v2
					}
				})(_1000_v2))), 2)
				(_nw150).ArraySet1(_1016_v16, 3)
				(_nw150).ArraySet1(_1016_v16, 4)
				(_nw150).ArraySet1(_1016_v16, 5)
				(_nw150).ArraySet1(_1016_v16, 6)
				(_nw150).ArraySet1(_1016_v16, 7)
				(_nw150).ArraySet1(_dafny.SeqOf(_1000_v2, _1006_i1), 8)
				(_nw150).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-645))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1022_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1023_i4 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_1022_v2)
					}
				})(_1000_v2))), 9)
				(_nw150).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1024_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1025_i5 _dafny.Int) _dafny.Int {
						return _1024_i1
					}
				})(_1006_i1))), 10)
				(_nw150).ArraySet1(_1016_v16, 11)
				(_nw150).ArraySet1(_1016_v16, 12)
				(_nw150).ArraySet1(_1016_v16, 13)
				(_nw150).ArraySet1(_1016_v16, 14)
				(_nw150).ArraySet1(_1016_v16, 15)
				(_nw150).ArraySet1(_1016_v16, 16)
				(_nw150).ArraySet1(_1016_v16, 17)
				(_nw150).ArraySet1(_1016_v16, 18)
				(_nw150).ArraySet1(_dafny.SeqOf(_1006_i1, _1000_v2), 19)
				(_nw150).ArraySet1(_1016_v16, 20)
				(_nw150).ArraySet1(_1016_v16, 21)
				(_nw150).ArraySet1(_1016_v16, 22)
				(_nw150).ArraySet1(_1016_v16, 23)
				(_nw150).ArraySet1(_dafny.SeqOf(_1006_i1), 24)
				(_nw150).ArraySet1(_1016_v16, 25)
				(_nw150).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_999_v1).Cardinality()), _1000_v2), 26)
				(_nw150).ArraySet1(_1016_v16, 27)
				(_nw150).ArraySet1(_1016_v16, 28)
				_1017_v17 = _nw150
				var _1026_v18 _dafny.Map
				_ = _1026_v18
				_1026_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1017_v17, _998_v0)
				var _1027_v19 _dafny.MultiSet
				_ = _1027_v19
				_1027_v19 = _dafny.MultiSetOf((func() bool {
					if (_1026_v18).Contains(_1017_v17) {
						return (_1026_v18).Get(_1017_v17).(bool)
					}
					return _998_v0
				})())
				var _1028_v20 D0
				_ = _1028_v20
				_1028_v20 = Companion_D0_.Create_DC0_(_1027_v19)
				var _1029_v21 _dafny.Array
				_ = _1029_v21
				var _nwElement0_34 D0 = _1028_v20
				_ = _nwElement0_34
				var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(19))
				_ = _nw151
				(_nw151).ArraySet1(_nwElement0_34, 0)
				(_nw151).ArraySet1(_1028_v20, 1)
				(_nw151).ArraySet1(_1028_v20, 2)
				(_nw151).ArraySet1((func() D0 {
					if !(_998_v0) {
						return _1028_v20
					}
					return _1028_v20
				})(), 3)
				(_nw151).ArraySet1(_1028_v20, 4)
				(_nw151).ArraySet1(Companion_D0_.Create_DC0_(_dafny.MultiSetOf(!(true), _998_v0)), 5)
				(_nw151).ArraySet1(_1028_v20, 6)
				(_nw151).ArraySet1(_1028_v20, 7)
				(_nw151).ArraySet1(_1028_v20, 8)
				(_nw151).ArraySet1(_1028_v20, 9)
				(_nw151).ArraySet1(_1028_v20, 10)
				(_nw151).ArraySet1(_1028_v20, 11)
				(_nw151).ArraySet1(Companion_D0_.Create_DC0_(_1027_v19), 12)
				(_nw151).ArraySet1(Companion_Default___.Fm20((_1027_v19).Cardinality(), Companion_Default___.Fm20(_1000_v2, _1028_v20, globalState), globalState), 13)
				(_nw151).ArraySet1(_1028_v20, 14)
				(_nw151).ArraySet1(_1028_v20, 15)
				(_nw151).ArraySet1(_1028_v20, 16)
				(_nw151).ArraySet1(_1028_v20, 17)
				(_nw151).ArraySet1(_1028_v20, 18)
				_1029_v21 = _nw151
				var _1030_v22 _dafny.Sequence
				_ = _1030_v22
				_1030_v22 = _dafny.UnicodeSeqOfUtf8Bytes("dgjysbxnh")
				var _1031_v23 _dafny.Array
				_ = _1031_v23
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_27
				var _nw152 _dafny.Array
				_ = _nw152
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw152 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Int = (func(_1032_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1033_i6 _dafny.Int) _dafny.Int {
							return (_1033_i6).Times(_1032_v2)
						}
					})(_1000_v2)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw152 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw152).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw152).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1031_v23 = _nw152
				var _rhs163 _dafny.Array = _1029_v21
				_ = _rhs163
				var _rhs164 _dafny.Array = _1031_v23
				_ = _rhs164
				var _rhs165 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_1000_v2, _1006_i1), _1006_i1))
				_ = _rhs165
				var _rhs166 bool = _998_v0
				_ = _rhs166
				var _rhs167 _dafny.Sequence = _1030_v22
				_ = _rhs167
				var _lhs118 *GlobalState = globalState
				_ = _lhs118
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				_1029_v21 = _rhs163
				_lhs118.F22 = _rhs164
				_lhs119.F7 = _rhs165
				_lhs120.F13 = _rhs166
				_1030_v22 = _rhs167
			}
			_1000_v2 = _1006_i1
			(globalState).F13 = _998_v0
			var _1034_v24 _dafny.CodePoint
			_ = _1034_v24
			_1034_v24 = _dafny.CodePoint('i')
			_1034_v24 = _1034_v24
		}
		var _1035_v25 _dafny.Map
		_ = _1035_v25
		_1035_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v2, _998_v0)).Cardinality()).Cmp(_1000_v2) <= 0, _998_v0)
		_1035_v25 = (_1035_v25).Update(_998_v0, _998_v0)
		var _1036_v26 _dafny.Array
		_ = _1036_v26
		var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw153
		_1036_v26 = _nw153
		var _1037_v27 _dafny.Map
		_ = _1037_v27
		_1037_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v2, _1000_v2)
		var _1038_v28 _dafny.Map
		_ = _1038_v28
		_1038_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v27, _998_v0)
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1036_v26), 0))
		_ = _index138
		(_1036_v26).ArraySet1(((_1038_v28).Cardinality()).Plus(_1000_v2), (_index138).Int())
		var _1039_v29 _dafny.Sequence
		_ = _1039_v29
		_1039_v29 = _dafny.UnicodeSeqOfUtf8Bytes("ejirnglkd")
		var _1040_v30 D2
		_ = _1040_v30
		_1040_v30 = Companion_D2_.Create_DC7_(_998_v0, _998_v0)
		var _1041_v31 D2
		_ = _1041_v31
		_1041_v31 = Companion_D2_.Create_DC7_(_dafny.Companion_Sequence_.IsProperPrefixOf(_1039_v29, _1039_v29), (_1040_v30).Dtor_cf13())
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1036_v26), 0))
		_ = _index139
		var _rhs168 _dafny.Int = _1000_v2
		_ = _rhs168
		var _rhs169 D2 = _1040_v30
		_ = _rhs169
		var _rhs170 _dafny.Int = (_1000_v2).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if _998_v0 {
				return _999_v1
			}
			return _dafny.SeqOf(_998_v0)
		})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1000_v2), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _998_v0 {
				return _999_v1
			}
			return _dafny.SeqOf(_998_v0)
		})()).Cardinality()))).Uint32(), _998_v0)).Cardinality()))
		_ = _rhs170
		var _rhs171 bool = _998_v0
		_ = _rhs171
		var _rhs172 _dafny.Int = Companion_Default___.SafeDivisionInt(_1000_v2, (_1000_v2).Times(_dafny.IntOfInt64(304)))
		_ = _rhs172
		var _lhs121 _dafny.Array = _1036_v26
		_ = _lhs121
		var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1036_v26), 0))
		_ = _lhs122
		var _lhs123 *GlobalState = globalState
		_ = _lhs123
		var _lhs124 *GlobalState = globalState
		_ = _lhs124
		var _lhs125 *GlobalState = globalState
		_ = _lhs125
		(_lhs121).ArraySet1(_rhs168, (_lhs122).Int())
		_1041_v31 = _rhs169
		_lhs123.F7 = _rhs170
		_lhs124.F13 = _rhs171
		_lhs125.F7 = _rhs172
		(globalState).F13 = _998_v0
		r0 = (_this).F32()
		return r0
	}
}
func (_this *C5) M9(p0 bool, p1 bool, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _1042_v0 _dafny.Sequence
		_ = _1042_v0
		_1042_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ri")
		var _1043_v1 _dafny.Sequence
		_ = _1043_v1
		_1043_v1 = _dafny.SeqOf(_1042_v0, _dafny.UnicodeSeqOfUtf8Bytes("gwjxmn"))
		var _1044_v2 _dafny.Set
		_ = _1044_v2
		_1044_v2 = _dafny.SetOf(p1, p0)
		var _1045_v3 _dafny.Int
		_ = _1045_v3
		_1045_v3 = _dafny.IntOfInt64(858)
		var _1046_v4 _dafny.MultiSet
		_ = _1046_v4
		_1046_v4 = _dafny.MultiSetOf(_1045_v3, _1045_v3)
		_1043_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1042_v0, _1042_v0, _1042_v0), Companion_Default___.Fm21(_1044_v2, _1046_v4, globalState))
		if p0 {
			var _1047_v5 _dafny.CodePoint
			_ = _1047_v5
			_1047_v5 = _dafny.CodePoint('v')
			var _1048_v6 _dafny.Map
			_ = _1048_v6
			_1048_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("av"), (Companion_Default___.SafeIndex(_1045_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("av")).Cardinality()))).Uint32(), _1047_v5))
			var _1049_v7 *C0
			_ = _1049_v7
			var _nw154 *C0 = New_C0_()
			_ = _nw154
			_nw154.Ctor__((func() _dafny.Map {
				if p0 {
					return _1048_v6
				}
				return _1048_v6
			})(), (!(p0)) && (p1), _1047_v5, p0)
			_1049_v7 = _nw154
			(globalState).F7 = _1045_v3
			var _1050_v8 _dafny.Sequence
			_ = _1050_v8
			_1050_v8 = _dafny.SeqOf((_1049_v7).F38(), !((_1049_v7).F38()))
			var _1051_v9 _dafny.Array
			_ = _1051_v9
			var _nwElement0_35 _dafny.Sequence = _1050_v8
			_ = _nwElement0_35
			var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(13))
			_ = _nw155
			(_nw155).ArraySet1(_nwElement0_35, 0)
			(_nw155).ArraySet1(_1050_v8, 1)
			(_nw155).ArraySet1(_dafny.SeqOf(p1, p1, (_1049_v7).F38()), 2)
			(_nw155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1050_v8, _1050_v8), 3)
			(_nw155).ArraySet1(_1050_v8, 4)
			(_nw155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1049_v7).F38(), p0), _1050_v8), 5)
			(_nw155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1050_v8, _1050_v8), 6)
			(_nw155).ArraySet1(_1050_v8, 7)
			(_nw155).ArraySet1(_dafny.SeqOf(!((_1049_v7).F38()), p1, !(p1)), 8)
			(_nw155).ArraySet1(_dafny.SeqOf(p0), 9)
			(_nw155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1050_v8, _1050_v8), 10)
			(_nw155).ArraySet1(_1050_v8, 11)
			(_nw155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _1050_v8), 12)
			_1051_v9 = _nw155
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1051_v9), 0))
			_ = _index140
			(_1051_v9).ArraySet1(_1050_v8, (_index140).Int())
			var _1052_v10 _dafny.MultiSet
			_ = _1052_v10
			_1052_v10 = _dafny.MultiSetOf(p0)
			var _1053_v11 _dafny.Set
			_ = _1053_v11
			_1053_v11 = _dafny.SetOf(_dafny.IntOfInt64(829), _1045_v3, _1045_v3, _1045_v3)
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1051_v9), 0))
			_ = _index141
			var _rhs173 _dafny.CodePoint = _1047_v5
			_ = _rhs173
			var _rhs174 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _1050_v8), _dafny.Companion_Sequence_.Update(_1050_v8, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1052_v10).Contains(true) {
					return (_1052_v10).Multiplicity(true)
				}
				return _1045_v3
			})(), _dafny.IntOfUint32((_1050_v8).Cardinality()))).Uint32(), p1))
			_ = _rhs174
			var _rhs175 bool = false
			_ = _rhs175
			var _rhs176 bool = (_1053_v11).IsDisjointFrom(_1053_v11)
			_ = _rhs176
			var _lhs126 _dafny.Array = _1051_v9
			_ = _lhs126
			var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1051_v9), 0))
			_ = _lhs127
			var _lhs128 *GlobalState = globalState
			_ = _lhs128
			var _lhs129 *GlobalState = globalState
			_ = _lhs129
			_1047_v5 = _rhs173
			(_lhs126).ArraySet1(_rhs174, (_lhs127).Int())
			_lhs128.F13 = _rhs175
			_lhs129.F13 = _rhs176
			(globalState).F7 = _1045_v3
			if (_1049_v7).F38() {
				var _1054_v12 _dafny.Sequence
				_ = _1054_v12
				_1054_v12 = _dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1045_v3, _1045_v3)))
				_1054_v12 = (_1049_v7).Fm26(_dafny.IntOfUint32((_1042_v0).Cardinality()), (func() _dafny.Sequence {
					if p0 {
						return (_1051_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1051_v9), 0))).Int()).(_dafny.Sequence)
					}
					return (_1051_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1051_v9), 0))).Int()).(_dafny.Sequence)
				})(), func() _dafny.Set {
					var _coll73 = _dafny.NewBuilder()
					_ = _coll73
					for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(885), _dafny.IntOfInt64(965))); ; {
						_compr_73, _ok77 := _iter77()
						if !_ok77 {
							break
						}
						var _1055_v13 _dafny.Int
						_1055_v13 = interface{}(_compr_73).(_dafny.Int)
						if ((_dafny.IntOfInt64(885)).Cmp(_1055_v13) <= 0) && ((_1055_v13).Cmp(_dafny.IntOfInt64(965)) < 0) {
							_coll73.Add(Companion_Default___.SafeModuloInt(_1055_v13, _1045_v3))
						}
					}
					return _coll73.ToSet()
				}(), (_1049_v7).F38(), globalState)
				(globalState).F13 = p1
				(globalState).F13 = (_1049_v7).F38()
				(globalState).F7 = _1045_v3
				var _1056_v14 _dafny.Array
				_ = _1056_v14
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
				_ = _nw156
				_1056_v14 = _nw156
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1056_v14), 0))
				_ = _index142
				(_1056_v14).ArraySet1CodePoint(_1047_v5, (_index142).Int())
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1056_v14), 0))
				_ = _index143
				var _rhs177 _dafny.Sequence = (_1043_v1).Select((Companion_Default___.SafeIndex(_1045_v3, _dafny.IntOfUint32((_1043_v1).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs177
				var _rhs178 _dafny.CodePoint = _dafny.CodePoint('l')
				_ = _rhs178
				var _lhs130 _dafny.Array = _1056_v14
				_ = _lhs130
				var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1056_v14), 0))
				_ = _lhs131
				_1042_v0 = _rhs177
				(_lhs130).ArraySet1CodePoint(_rhs178, (_lhs131).Int())
			} else {
				var _1057_v15 _dafny.Array
				_ = _1057_v15
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
				_ = _nw157
				_1057_v15 = _nw157
				var _1058_v16 *C2
				_ = _1058_v16
				var _nw158 *C2 = New_C2_()
				_ = _nw158
				_nw158.Ctor__(_1052_v10, _1057_v15, _1047_v5, (_1049_v7).F38())
				_1058_v16 = _nw158
				var _1059_v17 _dafny.Sequence
				_ = _1059_v17
				_1059_v17 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_this).Fm17(_1045_v3, _1047_v5, (_1049_v7).F38(), globalState), _1045_v3))
				_1059_v17 = _dafny.Companion_Sequence_.Concatenate((_1058_v16).Fm7(globalState), _1059_v17)
				var _1060_v18 _dafny.MultiSet
				_ = _1060_v18
				_1060_v18 = _dafny.MultiSetOf(_1053_v11, (_1053_v11).Difference(_1053_v11), _1053_v11, _1053_v11, (_1053_v11).Union(_1053_v11))
				_1060_v18 = ((_1060_v18).Difference(_1060_v18)).Intersection((_1060_v18).Union(_dafny.MultiSetOf(_1053_v11, _1053_v11)))
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen(((_this).F32()), 0))
				_ = _index144
				((_this).F32()).ArraySet1((_1049_v7).F38(), (_index144).Int())
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen(((_this).F32()), 0))
				_ = _index145
				((_this).F32()).ArraySet1((_1049_v7).F38(), (_index145).Int())
				(globalState).F7 = _1045_v3
			}
		} else {
			var _1061_v19 *C1
			_ = _1061_v19
			var _nw159 *C1 = New_C1_()
			_ = _nw159
			_nw159.Ctor__(p1, p1)
			_1061_v19 = _nw159
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))
			_ = _index146
			((_this).F32()).ArraySet1(p1, (_index146).Int())
			var _1062_v20 _dafny.Sequence
			_ = _1062_v20
			_1062_v20 = _dafny.SeqOf((_1061_v19).F39())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))
			_ = _index147
			((_this).F32()).ArraySet1(!((func() bool {
				if false {
					return (_1045_v3).Cmp(_dafny.IntOfUint32((_1062_v20).Cardinality())) < 0
				}
				return true
			})()), (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))
			_ = _index148
			((_this).F32()).ArraySet1(((_this).F32()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))).Int()).(bool), (_index148).Int())
			var _1063_v21 _dafny.Array
			_ = _1063_v21
			var _len0_28 _dafny.Int = _dafny.One
			_ = _len0_28
			var _nw160 _dafny.Array
			_ = _nw160
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw160 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Map = (func(_1064_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_1065_i0 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1064_v20)
					}
				})(_1062_v20)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw160 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw160).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw160).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1063_v21 = _nw160
			var _1066_v22 _dafny.Map
			_ = _1066_v22
			_1066_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1061_v19).F39(), Companion_Default___.Fm36(globalState))
			var _1067_v23 D9
			_ = _1067_v23
			_1067_v23 = Companion_D9_.Create_DC19_(_1066_v22)
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1063_v21), 0))
			_ = _index149
			(_1063_v21).ArraySet1((_1067_v23).Dtor_cf30(), (_index149).Int())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1063_v21), 0))
			_ = _index150
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))
			_ = _index151
			var _rhs179 _dafny.Map = _1066_v22
			_ = _rhs179
			var _rhs180 bool = p0
			_ = _rhs180
			var _rhs181 bool = (_1061_v19).F39()
			_ = _rhs181
			var _rhs182 bool = p0
			_ = _rhs182
			var _lhs132 _dafny.Array = _1063_v21
			_ = _lhs132
			var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1063_v21), 0))
			_ = _lhs133
			var _lhs134 *GlobalState = globalState
			_ = _lhs134
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			var _lhs136 _dafny.Array = (_this).F32()
			_ = _lhs136
			var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))
			_ = _lhs137
			(_lhs132).ArraySet1(_rhs179, (_lhs133).Int())
			_lhs134.F13 = _rhs180
			_lhs135.F13 = _rhs181
			(_lhs136).ArraySet1(_rhs182, (_lhs137).Int())
			if _dafny.Companion_Sequence_.IsPrefixOf(_1062_v20, _1062_v20) {
				var _1068_v24 *C1
				_ = _1068_v24
				var _nw161 *C1 = New_C1_()
				_ = _nw161
				_nw161.Ctor__((func() bool {
					if p1 {
						return ((_this).F32()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))).Int()).(bool)
					}
					return p1
				})(), p1)
				_1068_v24 = _nw161
				(globalState).F13 = (((_this).F32()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))).Int()).(bool)) == ((_1068_v24).F39())
				(globalState).F7 = _1045_v3
				(globalState).F7 = (_1045_v3).Plus((_dafny.Zero).Minus(_1045_v3))
				var _1069_v25 _dafny.Array
				_ = _1069_v25
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_29
				var _nw162 _dafny.Array
				_ = _nw162
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw162 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.MultiSet = (func(_1070_v19 *C1, _1071_p1 bool, _1072_p0 bool, _1073_v24 *C1) func(_dafny.Int) _dafny.MultiSet {
						return func(_1074_i1 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F32()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))).Int()).(bool), ((_this).F32()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))).Int()).(bool)), (_1070_v19).F39(), _1071_p1), Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1072_p0), _1072_p0), (_1070_v19).F39(), (_1073_v24).F39()))).Intersection(_dafny.MultiSetOf(Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F32()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F32()), 0))).Int()).(bool), (_1070_v19).F39()), (_1073_v24).F39(), !((_1070_v19).F39()))))
						}
					})(_1061_v19, p1, p0, _1068_v24)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw162 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw162).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw162).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1069_v25 = _nw162
				var _1075_v26 _dafny.Map
				_ = _1075_v26
				_1075_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _1076_v27 D6
				_ = _1076_v27
				_1076_v27 = Companion_D6_.Create_DC16_(_1075_v26, true, p1)
				var _1077_v28 _dafny.MultiSet
				_ = _1077_v28
				_1077_v28 = _dafny.MultiSetOf(_1076_v27, _1076_v27)
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1069_v25), 0))
				_ = _index152
				(_1069_v25).ArraySet1((_1077_v28).Union((_1077_v28).Update(_1076_v27, Companion_Default___.Abs(_1045_v3))), (_index152).Int())
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_1069_v25), 0))
				_ = _index153
				(_1069_v25).ArraySet1(_1077_v28, (_index153).Int())
			} else {
				var _1078_v29 _dafny.Map
				_ = _1078_v29
				_1078_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1061_v19).F39(), _1042_v0)
				var _1079_v30 D2
				_ = _1079_v30
				_1079_v30 = Companion_D2_.Create_DC6_(_1078_v29)
				_1079_v30 = _1079_v30
				var _1080_v31 _dafny.CodePoint
				_ = _1080_v31
				_1080_v31 = _dafny.CodePoint('c')
				_1080_v31 = _dafny.CodePoint('d')
				var _1081_v32 _dafny.MultiSet
				_ = _1081_v32
				_1081_v32 = _dafny.MultiSetOf(!((_1061_v19).F39()))
				var _1082_v33 D1
				_ = _1082_v33
				_1082_v33 = Companion_D1_.Create_DC4_(_1080_v31)
				var _1083_v34 _dafny.Array
				_ = _1083_v34
				var _nwElement0_36 _dafny.CodePoint = _1080_v31
				_ = _nwElement0_36
				var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(28))
				_ = _nw163
				(_nw163).ArraySet1CodePoint(_nwElement0_36, 0)
				(_nw163).ArraySet1CodePoint(_1080_v31, 1)
				(_nw163).ArraySet1CodePoint(_dafny.CodePoint('j'), 2)
				(_nw163).ArraySet1CodePoint(_1080_v31, 3)
				(_nw163).ArraySet1CodePoint(_1080_v31, 4)
				(_nw163).ArraySet1CodePoint(_1080_v31, 5)
				(_nw163).ArraySet1CodePoint(_dafny.CodePoint('r'), 6)
				(_nw163).ArraySet1CodePoint(_1080_v31, 7)
				(_nw163).ArraySet1CodePoint(_1080_v31, 8)
				(_nw163).ArraySet1CodePoint(_1080_v31, 9)
				(_nw163).ArraySet1CodePoint(_1080_v31, 10)
				(_nw163).ArraySet1CodePoint(_1080_v31, 11)
				(_nw163).ArraySet1CodePoint(_1080_v31, 12)
				(_nw163).ArraySet1CodePoint(_1080_v31, 13)
				(_nw163).ArraySet1CodePoint(_1080_v31, 14)
				(_nw163).ArraySet1CodePoint((_1082_v33).Dtor_cf6(), 15)
				(_nw163).ArraySet1CodePoint(_1080_v31, 16)
				(_nw163).ArraySet1CodePoint(_1080_v31, 17)
				(_nw163).ArraySet1CodePoint(_dafny.CodePoint('o'), 18)
				(_nw163).ArraySet1CodePoint(_1080_v31, 19)
				(_nw163).ArraySet1CodePoint(_1080_v31, 20)
				(_nw163).ArraySet1CodePoint(_1080_v31, 21)
				(_nw163).ArraySet1CodePoint(_1080_v31, 22)
				(_nw163).ArraySet1CodePoint((_1082_v33).Dtor_cf6(), 23)
				(_nw163).ArraySet1CodePoint(_1080_v31, 24)
				(_nw163).ArraySet1CodePoint(_dafny.CodePoint('s'), 25)
				(_nw163).ArraySet1CodePoint(_dafny.CodePoint('a'), 26)
				(_nw163).ArraySet1CodePoint(_1080_v31, 27)
				_1083_v34 = _nw163
				var _1084_v35 *C2
				_ = _1084_v35
				var _nw164 *C2 = New_C2_()
				_ = _nw164
				_nw164.Ctor__(_1081_v32, _1083_v34, _1080_v31, (_1061_v19).F39())
				_1084_v35 = _nw164
				var _1085_v36 D4
				_ = _1085_v36
				_1085_v36 = Companion_D4_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(198))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1086_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1087_i2 _dafny.Int) _dafny.CodePoint {
						return _1086_v31
					}
				})(_1080_v31))))
				var _1088_v37 _dafny.Set
				_ = _1088_v37
				_1088_v37 = _dafny.SetOf(_dafny.IntOfInt64(-793))
				_1085_v36 = func(_pat_let43_0 D4) D4 {
					return func(_1089_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let44_0 _dafny.Sequence) D4 {
							return func(_1090_dt__update_hcf17_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC10_(_1090_dt__update_hcf17_h0)
							}(_pat_let44_0)
						}(_dafny.UnicodeSeqOfUtf8Bytes("qrk"))
					}(_pat_let43_0)
				}(Companion_Default___.Fm44((_1088_v37).Cardinality(), globalState))
				var _1091_v38 _dafny.Map
				_ = _1091_v38
				_1091_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1061_v19).F39(), _dafny.IntOfInt64(899))
				var _1092_v39 _dafny.Sequence
				_ = _1092_v39
				_1092_v39 = _dafny.SeqOf((_1091_v38).Cardinality())
				var _1093_v40 _dafny.Array
				_ = _1093_v40
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_30
				var _nw165 _dafny.Array
				_ = _nw165
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw165 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Int = (func(_1094_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1095_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1095_i3, _1094_v3)
						}
					})(_1045_v3)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw165 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw165).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw165).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1093_v40 = _nw165
				var _1096_v41 _dafny.Map
				_ = _1096_v41
				_1096_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1093_v40, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1042_v0, (Companion_Default___.SafeIndex(_1045_v3, _dafny.IntOfUint32((_1042_v0).Cardinality()))).Uint32(), _1080_v31), (Companion_Default___.SafeIndex(_1045_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1042_v0, (Companion_Default___.SafeIndex(_1045_v3, _dafny.IntOfUint32((_1042_v0).Cardinality()))).Uint32(), _1080_v31)).Cardinality()))).Uint32(), _1080_v31))
				var _1097_v42 _dafny.Map
				_ = _1097_v42
				_1097_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1061_v19).Fm9(_1092_v39, _1045_v3, (_1061_v19).F39(), globalState), _1096_v41)
				_1097_v42 = (_1097_v42).Update((_1061_v19).Fm9(_1092_v39, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-662))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}((func(_1098_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1099_i4 _dafny.Int) _dafny.Int {
						return _1098_v3
					}
				})(_1045_v3)))).Cardinality()), (_1061_v19).F39(), globalState), _1096_v41)
			}
		}
		(globalState).F7 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1045_v3), _1045_v3)
		var _1100_v43 _dafny.Map
		_ = _1100_v43
		_1100_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1045_v3, (_dafny.Zero).Minus(_1045_v3))
		var _hi6 _dafny.Int = (func() _dafny.Int {
			if (_1100_v43).Contains(_1045_v3) {
				return (_1100_v43).Get(_1045_v3).(_dafny.Int)
			}
			return _1045_v3
		})()
		_ = _hi6
		for _1101_i5 := (_dafny.IntOfInt64(56)).Plus(_dafny.IntOfInt64(431)); _1101_i5.Cmp(_hi6) < 0; _1101_i5 = _1101_i5.Plus(_dafny.One) {
			var _1102_v44 _dafny.Array
			_ = _1102_v44
			var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw166
			_1102_v44 = _nw166
			_1102_v44 = _1102_v44
			(globalState).F13 = !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(575))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg80 _dafny.Int) interface{} {
					return coer80(arg80)
				}
			}(func(_1103_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), _1042_v0), _1042_v0))
			(globalState).F13 = p0
			_1046_v4 = (_1046_v4).Intersection(_1046_v4)
		}
		var _1104_v45 _dafny.Array
		_ = _1104_v45
		var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(9))
		_ = _nw167
		_1104_v45 = _nw167
		var _1105_v46 _dafny.Array
		_ = _1105_v46
		var _len0_31 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_31
		var _nw168 _dafny.Array
		_ = _nw168
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw168 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) _dafny.Int = func(_1106_i7 _dafny.Int) _dafny.Int {
				return (_1106_i7).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
			}
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw168 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw168).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw168).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_1105_v46 = _nw168
		var _1107_v47 _dafny.Set
		_ = _1107_v47
		_1107_v47 = _dafny.SetOf(_1105_v46)
		var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1104_v45), 0))
		_ = _index154
		(_1104_v45).ArraySet1(_1107_v47, (_index154).Int())
		var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1104_v45), 0))
		_ = _index155
		(_1104_v45).ArraySet1((func() _dafny.Set {
			if !(p1) {
				return _1107_v47
			}
			return _1107_v47
		})(), (_index155).Int())
		var _1108_v48 _dafny.Array
		_ = _1108_v48
		var _len0_32 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_32
		var _nw169 _dafny.Array
		_ = _nw169
		if _len0_32.Cmp(_dafny.Zero) == 0 {
			_nw169 = _dafny.NewArray(_len0_32)
		} else {
			var _init32 func(_dafny.Int) _dafny.CodePoint = func(_1109_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}
			_ = _init32
			var _element0_32 = _init32(_dafny.Zero)
			_ = _element0_32
			_nw169 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
			(_nw169).ArraySet1CodePoint(_element0_32, 0)
			var _nativeLen0_32 = (_len0_32).Int()
			_ = _nativeLen0_32
			for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
				(_nw169).ArraySet1CodePoint(_init32(_dafny.IntOf(_i0_32)), _i0_32)
			}
		}
		_1108_v48 = _nw169
		var _1110_v49 _dafny.CodePoint
		_ = _1110_v49
		_1110_v49 = _dafny.CodePoint('g')
		var _1111_v50 *C4
		_ = _1111_v50
		var _nw170 *C4 = New_C4_()
		_ = _nw170
		_nw170.Ctor__(_dafny.IntOfInt64(-768), _1100_v43, _1108_v48, _1110_v49, p0)
		_1111_v50 = _nw170
		var _1112_v51 _dafny.Array
		_ = _1112_v51
		var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw171
		_1112_v51 = _nw171
		var _1113_v52 _dafny.Map
		_ = _1113_v52
		_1113_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1112_v51)
		r0 = (func() _dafny.Array {
			if (_1113_v52).Contains((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1043_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-248), _dafny.IntOfUint32((_1043_v1).Cardinality()))).Uint32(), _1042_v0)).Cardinality())).Cmp(_1045_v3) >= 0) {
				return (_1113_v52).Get((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1043_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-248), _dafny.IntOfUint32((_1043_v1).Cardinality()))).Uint32(), _1042_v0)).Cardinality())).Cmp(_1045_v3) >= 0).(_dafny.Array)
			}
			return _1112_v51
		})()
		return r0
	}
}
func (_this *C5) F32() _dafny.Array {
	{
		return _this._f32
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f23 _dafny.CodePoint
	_f24 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
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

func (_this *C6) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C6) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C6) F24() bool {
	return _this._f24
}
func (_this *C6) Ctor__(f23 _dafny.CodePoint, f24 bool) {
	{
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C6) Fm5(globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_this.F23())).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("lwt"))).Cardinality()))
	}
}
func (_this *C6) Fm6(globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1114_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("gt")))).Cardinality()
	}
}
func (_this *C6) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf(_dafny.MultiSetOf((_this).F24()), (Companion_D0_.Create_DC0_(_dafny.MultiSetOf((_this).F24(), (_this).F24(), (_this).F24(), (_this).F24(), (_this).F24()))).Dtor_cf0())).IsDisjointFrom((_dafny.MultiSetOf(_dafny.MultiSetOf((_this).F24()))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.MultiSetOf(false, (_this).F24())))))
	}
}
func (_this *C6) Fm16(globalState *GlobalState) bool {
	{
		return !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(298))).Contains((_this).F24())
	}
}
func (_this *C6) M1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) (_dafny.Sequence, D0, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var r2 bool = false
		_ = r2
		var _1115_v0 _dafny.Int
		_ = _1115_v0
		_1115_v0 = _dafny.IntOfInt64(914)
		var _hi7 _dafny.Int = _1115_v0
		_ = _hi7
		for _1116_i0 := _1115_v0; _1116_i0.Cmp(_hi7) < 0; _1116_i0 = _1116_i0.Plus(_dafny.One) {
			var _1117_v1 _dafny.Map
			_ = _1117_v1
			_1117_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("hblxvj"))
			var _1118_v2 *C0
			_ = _1118_v2
			var _nw172 *C0 = New_C0_()
			_ = _nw172
			_nw172.Ctor__(_1117_v1, (_this).F24(), _this.F23(), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xu")).Cardinality())).Cmp(_1116_i0) == 0)
			_1118_v2 = _nw172
			var _1119_v3 _dafny.Sequence
			_ = _1119_v3
			_1119_v3 = _dafny.UnicodeSeqOfUtf8Bytes("c")
			(globalState).F7 = _dafny.IntOfUint32((_1119_v3).Cardinality())
			var _rhs183 _dafny.Int = Companion_Default___.SafeDivisionInt(_1116_i0, _1115_v0)
			_ = _rhs183
			var _rhs184 _dafny.Int = _1115_v0
			_ = _rhs184
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			_lhs138.F7 = _rhs183
			_1115_v0 = _rhs184
			var _1120_v4 _dafny.Sequence
			_ = _1120_v4
			_1120_v4 = _dafny.SeqOf((_this).F24(), (_this).F24())
			if (func() bool {
				if (_1118_v2).F38() {
					return (func() bool {
						if (_this).F24() {
							return !(p1)
						}
						return true
					})()
				}
				return (_1120_v4).Select((Companion_Default___.SafeIndex(_1115_v0, _dafny.IntOfUint32((_1120_v4).Cardinality()))).Uint32()).(bool)
			})() {
				(globalState).F7 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1115_v0), (_1115_v0).Minus(_dafny.IntOfUint32((_1120_v4).Cardinality())))
				var _1121_v5 _dafny.Array
				_ = _1121_v5
				var _nw173 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw173
				_1121_v5 = _nw173
				var _1122_v6 _dafny.Map
				_ = _1122_v6
				_1122_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1121_v5, _1115_v0)
				var _1123_v7 _dafny.Sequence
				_ = _1123_v7
				_1123_v7 = _dafny.SeqOf(_1122_v6, _1122_v6)
				var _1124_v8 _dafny.Map
				_ = _1124_v8
				_1124_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), ((_1123_v7).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1116_i0), _dafny.IntOfUint32((_1123_v7).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
				var _1125_v9 _dafny.Map
				_ = _1125_v9
				_1125_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1116_i0, (_1124_v8).Merge(_1124_v8))
				_1125_v9 = (_1125_v9).Update((_dafny.Zero).Minus(_1116_i0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.SeqOf(_1115_v0)).Cardinality())))
				r0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_1118_v2).F38(), globalState), _1119_v3)
				(_this).F23_set_(_this.F23())
				_1115_v0 = _1115_v0
			} else {
				var _1126_v10 _dafny.Set
				_ = _1126_v10
				_1126_v10 = _dafny.SetOf(_1115_v0, _1115_v0)
				var _1127_v11 _dafny.Sequence
				_ = _1127_v11
				_1127_v11 = _dafny.SeqOf(_1126_v10, _1126_v10)
				(globalState).F7 = _dafny.IntOfUint32((_dafny.SeqOf((_dafny.IntOfUint32((_1127_v11).Cardinality())).Plus(_1116_i0), _1115_v0, _dafny.IntOfInt64(-974))).Cardinality())
				var _1128_v12 _dafny.Array
				_ = _1128_v12
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_33
				var _nw174 _dafny.Array
				_ = _nw174
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw174 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Int = (func(_1129_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1130_i1 _dafny.Int) _dafny.Int {
							return (_1130_i1).Minus(_1129_v0)
						}
					})(_1115_v0)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw174 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw174).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw174).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1128_v12 = _nw174
				(globalState).F7 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1115_v0, _1116_i0)), Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(_1128_v12)).Cardinality(), _1115_v0))
				(globalState).F7 = _1116_i0
				(globalState).F13 = (_this).F24()
				var _1131_v13 _dafny.Map
				_ = _1131_v13
				_1131_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p1)
				var _rhs185 _dafny.Int = _1116_i0
				_ = _rhs185
				var _rhs186 _dafny.Map = Companion_Default___.Fm19(globalState)
				_ = _rhs186
				_1115_v0 = _rhs185
				_1131_v13 = _rhs186
			}
		}
		var _1132_v14 _dafny.Array
		_ = _1132_v14
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_34
		var _nw175 _dafny.Array
		_ = _nw175
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw175 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) _dafny.Int = (func(_1133_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1134_i2 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1134_i2, _1133_v0)
				}
			})(_1115_v0)
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw175 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw175).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw175).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1132_v14 = _nw175
		var _1135_v15 _dafny.Sequence
		_ = _1135_v15
		_1135_v15 = _dafny.SeqOf(_1132_v14)
		var _1136_v16 _dafny.Sequence
		_ = _1136_v16
		_1136_v16 = _dafny.SeqOf((_1135_v15).Select((Companion_Default___.SafeIndex(_1115_v0, _dafny.IntOfUint32((_1135_v15).Cardinality()))).Uint32()).(_dafny.Array), _1132_v14)
		var _1137_v17 _dafny.Sequence
		_ = _1137_v17
		_1137_v17 = _dafny.UnicodeSeqOfUtf8Bytes("bniwatogj")
		var _1138_v18 _dafny.Map
		_ = _1138_v18
		_1138_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1137_v17)
		var _1139_v19 D2
		_ = _1139_v19
		_1139_v19 = Companion_D2_.Create_DC6_(_1138_v18)
		_1115_v0 = (_dafny.IntOfUint32(((func() _dafny.Sequence {
			if p1 {
				return _1136_v16
			}
			return _1136_v16
		})()).Cardinality())).Minus(_dafny.IntOfUint32((Companion_Default___.Fm45((_this).F24(), _1139_v19, _1115_v0, globalState)).Cardinality()))
		r2 = (_this).F24()
		var _1140_v20 *C1
		_ = _1140_v20
		var _nw176 *C1 = New_C1_()
		_ = _nw176
		_nw176.Ctor__(p1, p1)
		_1140_v20 = _nw176
		var _1141_v21 *C1
		_ = _1141_v21
		_1141_v21 = _1140_v20
		var _1142_v22 _dafny.Sequence
		_ = _1142_v22
		_1142_v22 = _dafny.SeqOf(_1141_v21, _1141_v21, _1141_v21, _1141_v21, _1141_v21)
		var _1143_v23 _dafny.Sequence
		_ = _1143_v23
		_1143_v23 = _1142_v22
		var _1144_v24 _dafny.MultiSet
		_ = _1144_v24
		_1144_v24 = _dafny.MultiSetOf(_1142_v22, _dafny.SeqOf(_1141_v21, _1141_v21, _1141_v21), _1142_v22, (_1143_v23), _1142_v22)
		_1144_v24 = ((_1144_v24).Update(_1142_v22, Companion_Default___.Abs(_1115_v0))).Difference(_1144_v24)
		var _1145_i3 _dafny.Int
		_ = _1145_i3
		_1145_i3 = _dafny.Zero
		{
			for p1 {
				{
					if (_1145_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1145_i3 = (_1145_i3).Plus(_dafny.One)
					(globalState).F7 = _1115_v0
					(globalState).F22 = _1132_v14
					var _1146_v25 D5
					_ = _1146_v25
					_1146_v25 = Companion_D5_.Create_DC13_((_dafny.Zero).Minus(_1115_v0), (_1140_v20).F39(), (_this).F24())
					var _1147_v26 _dafny.Map
					_ = _1147_v26
					_1147_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(402), (_1146_v25).Dtor_cf21())
					_1147_v26 = (_1147_v26).Update(_dafny.IntOfInt64(431), p1)
					var _1148_v27 _dafny.Sequence
					_ = _1148_v27
					_1148_v27 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(153))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg82 _dafny.Int) interface{} {
							return coer82(arg82)
						}
					}((func(_1149_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1150_i4 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_1149_v0)
						}
					})(_1115_v0))))
					var _1151_v28 _dafny.Array
					_ = _1151_v28
					var _nwElement0_37 bool = p1
					_ = _nwElement0_37
					var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(10))
					_ = _nw177
					(_nw177).ArraySet1(_nwElement0_37, 0)
					(_nw177).ArraySet1(p1, 1)
					(_nw177).ArraySet1((_this).F24(), 2)
					(_nw177).ArraySet1(p1, 3)
					(_nw177).ArraySet1((_this).F24(), 4)
					(_nw177).ArraySet1((_this).F24(), 5)
					(_nw177).ArraySet1((_this).F24(), 6)
					(_nw177).ArraySet1(p1, 7)
					(_nw177).ArraySet1((_this).F24(), 8)
					(_nw177).ArraySet1((_1140_v20).F39(), 9)
					_1151_v28 = _nw177
					var _1152_v29 D1
					_ = _1152_v29
					_1152_v29 = Companion_D1_.Create_DC5_((_1115_v0).Plus(_1115_v0), _dafny.IntOfUint32((_1137_v17).Cardinality()), _1148_v27, (func() _dafny.Array {
						if true {
							return _1151_v28
						}
						return _1151_v28
					})(), _1115_v0)
					var _source16 D1 = _1152_v29
					_ = _source16
					if _source16.Is_DC5() {
						var _1153___mcc_h0 _dafny.Int = _source16.Get_().(D1_DC5).Cf7
						_ = _1153___mcc_h0
						var _1154___mcc_h1 _dafny.Int = _source16.Get_().(D1_DC5).Cf8
						_ = _1154___mcc_h1
						var _1155___mcc_h2 _dafny.Sequence = _source16.Get_().(D1_DC5).Cf9
						_ = _1155___mcc_h2
						var _1156___mcc_h3 _dafny.Array = _source16.Get_().(D1_DC5).Cf10
						_ = _1156___mcc_h3
						var _1157___mcc_h4 _dafny.Int = _source16.Get_().(D1_DC5).Cf11
						_ = _1157___mcc_h4
						var _1158_cf11 _dafny.Int = _1157___mcc_h4
						_ = _1158_cf11
						var _1159_cf10 _dafny.Array = _1156___mcc_h3
						_ = _1159_cf10
						var _1160_cf9 _dafny.Sequence = _1155___mcc_h2
						_ = _1160_cf9
						var _1161_cf8 _dafny.Int = _1154___mcc_h1
						_ = _1161_cf8
						var _1162_cf7 _dafny.Int = _1153___mcc_h0
						_ = _1162_cf7
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1159_cf10), 0))
						_ = _index156
						(_1159_cf10).ArraySet1((_1140_v20).F39(), (_index156).Int())
						var _1163_v30 _dafny.Map
						_ = _1163_v30
						_1163_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1140_v20).F39(), (_1140_v20).F39())
						var _1164_v31 _dafny.MultiSet
						_ = _1164_v31
						_1164_v31 = _dafny.MultiSetOf((_1140_v20).F39())
						var _1165_v32 _dafny.Sequence
						_ = _1165_v32
						_1165_v32 = _dafny.SeqOf((_1140_v20).F39(), p1)
						var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1159_cf10), 0))
						_ = _index157
						(_1159_cf10).ArraySet1((func() bool {
							if (_1163_v30).Contains((_1140_v20).F39()) {
								return (_1163_v30).Get((_1140_v20).F39()).(bool)
							}
							return (_dafny.MultiSetFromSeq(_1165_v32)).IsProperSubsetOf(_1164_v31)
						})(), (_index157).Int())
						var _1166_v33 D11
						_ = _1166_v33
						_1166_v33 = Companion_D11_.Create_DC22_(_1132_v14)
						var _1167_v34 _dafny.Map
						_ = _1167_v34
						_1167_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if p1 {
								return true
							}
							return (_this).F24()
						})(), (_1166_v33).Dtor_cf35())
						var _1168_v35 D5
						_ = _1168_v35
						_1168_v35 = Companion_D5_.Create_DC14_((_1140_v20).F39())
						(globalState).F22 = (func() _dafny.Array {
							if (_1167_v34).Contains(((_1140_v20).Fm32((_this).F24(), _1161_cf8, globalState)) == ((_1168_v35).Dtor_cf23())) {
								return (_1167_v34).Get(((_1140_v20).Fm32((_this).F24(), _1161_cf8, globalState)) == ((_1168_v35).Dtor_cf23())).(_dafny.Array)
							}
							return _1132_v14
						})()
						var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1159_cf10), 0))
						_ = _index158
						(_1159_cf10).ArraySet1(!((_this).F24()), (_index158).Int())
						var _1169_v36 _dafny.Map
						_ = _1169_v36
						_1169_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1137_v17)
						_1161_cf8 = (_1169_v36).Cardinality()
					} else {
						var _1170___mcc_h5 _dafny.CodePoint = _source16.Get_().(D1_DC4).Cf6
						_ = _1170___mcc_h5
						var _1171_cf6 _dafny.CodePoint = _1170___mcc_h5
						_ = _1171_cf6
						var _1172_v37 bool
						_ = _1172_v37
						var _out36 bool
						_ = _out36
						_out36 = (_1140_v20).M4(false, globalState)
						_1172_v37 = _out36
						_1115_v0 = _1115_v0
						var _1173_v38 _dafny.Sequence
						_ = _1173_v38
						_1173_v38 = _dafny.SeqOf(_1115_v0, _1115_v0, _1115_v0, _dafny.IntOfUint32((_1137_v17).Cardinality()))
						var _1174_v39 _dafny.Array
						_ = _1174_v39
						var _len0_35 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_35
						var _nw178 _dafny.Array
						_ = _nw178
						if _len0_35.Cmp(_dafny.Zero) == 0 {
							_nw178 = _dafny.NewArray(_len0_35)
						} else {
							var _init35 func(_dafny.Int) _dafny.CodePoint = func(_1175_i5 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('s')
							}
							_ = _init35
							var _element0_35 = _init35(_dafny.Zero)
							_ = _element0_35
							_nw178 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
							(_nw178).ArraySet1CodePoint(_element0_35, 0)
							var _nativeLen0_35 = (_len0_35).Int()
							_ = _nativeLen0_35
							for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
								(_nw178).ArraySet1CodePoint(_init35(_dafny.IntOf(_i0_35)), _i0_35)
							}
						}
						_1174_v39 = _nw178
						var _1176_v40 D4
						_ = _1176_v40
						_1176_v40 = Companion_D4_.Create_DC10_(_1137_v17)
						var _1177_v41 *C4
						_ = _1177_v41
						var _nw179 *C4 = New_C4_()
						_ = _nw179
						_nw179.Ctor__(_1115_v0, Companion_Default___.Fm46(_1115_v0, _1173_v38, globalState), _1174_v39, Companion_Default___.Fm34(_dafny.IntOfUint32(((_1176_v40).Dtor_cf17()).Cardinality()), (_1140_v20).F39(), globalState), p1)
						_1177_v41 = _nw179
						var _rhs187 _dafny.Int = _dafny.IntOfInt64(239)
						_ = _rhs187
						var _rhs188 _dafny.Int = (_dafny.Zero).Minus((_1177_v41).F33())
						_ = _rhs188
						var _lhs139 *GlobalState = globalState
						_ = _lhs139
						_lhs139.F7 = _rhs187
						_1115_v0 = _rhs188
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1178_v42 _dafny.Sequence
		_ = _1178_v42
		_1178_v42 = _dafny.SeqOf((_this).F24(), false, (func() bool {
			if !((_this).F24()) {
				return p1
			}
			return (_this).F24()
		})())
		if (_1178_v42).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.IntOfUint32((_1178_v42).Cardinality()))).Uint32()).(bool) {
			var _1179_v43 _dafny.Set
			_ = _1179_v43
			_1179_v43 = _dafny.SetOf((_1140_v20).F39())
			(globalState).F7 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-258))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg83 _dafny.Int) interface{} {
					return coer83(arg83)
				}
			}((func(_1180_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1181_i6 _dafny.Int) _dafny.CodePoint {
					return _1180_p0
				}
			})(p0))), _1137_v17)).Cardinality()), (_dafny.Zero).Minus((_1179_v43).Cardinality()))
			_1143_v23 = _1143_v23
			_1135_v15 = _1136_v16
			_1115_v0 = _1115_v0
			var _1182_v44 _dafny.Array
			_ = _1182_v44
			var _nw180 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw180
			_1182_v44 = _nw180
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1182_v44), 0))
			_ = _index159
			(_1182_v44).ArraySet1((_1115_v0).Cmp(_dafny.IntOfInt64(805)) > 0, (_index159).Int())
			var _1183_v45 _dafny.Sequence
			_ = _1183_v45
			_1183_v45 = _dafny.SeqOf(_1115_v0)
			var _1184_v46 _dafny.Map
			_ = _1184_v46
			_1184_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1178_v42).Cardinality()), (_1183_v45))
			var _1185_v47 _dafny.Map
			_ = _1185_v47
			_1185_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1184_v46, (func() bool {
				if true {
					return (_this).F24()
				}
				return (_this).F24()
			})())
			var _1186_v48 _dafny.Set
			_ = _1186_v48
			_1186_v48 = _dafny.SetOf(_1115_v0, _1115_v0, _dafny.IntOfInt64(50))
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1182_v44), 0))
			_ = _index160
			var _rhs189 bool = (func() bool {
				if (_1185_v47).Contains(Companion_Default___.Fm47((Companion_Default___.Fm18(globalState)).Cardinality(), globalState)) {
					return (_1185_v47).Get(Companion_Default___.Fm47((Companion_Default___.Fm18(globalState)).Cardinality(), globalState)).(bool)
				}
				return (_dafny.SetOf(Companion_Default___.Fm0(p1, _1115_v0, globalState))).IsProperSubsetOf(_1186_v48)
			})()
			_ = _rhs189
			var _rhs190 _dafny.CodePoint = _dafny.CodePoint('i')
			_ = _rhs190
			var _rhs191 _dafny.Int = _1115_v0
			_ = _rhs191
			var _lhs140 _dafny.Array = _1182_v44
			_ = _lhs140
			var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1182_v44), 0))
			_ = _lhs141
			var _lhs142 *C6 = _this
			_ = _lhs142
			var _lhs143 *GlobalState = globalState
			_ = _lhs143
			(_lhs140).ArraySet1(_rhs189, (_lhs141).Int())
			_lhs142.F23_set_(_rhs190)
			_lhs143.F7 = _rhs191
		} else {
			_1115_v0 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg84 _dafny.Int) interface{} {
					return coer84(arg84)
				}
			}((func(_1187_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_1188_i7 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_1187_v0)
				}
			})(_1115_v0)))).Cardinality())
			var _rhs192 bool = (_1140_v20).F39()
			_ = _rhs192
			var _rhs193 _dafny.Int = (_1115_v0).Times((func() _dafny.Int {
				if (_this).F24() {
					return _1115_v0
				}
				return _dafny.IntOfInt64(-78)
			})())
			_ = _rhs193
			var _rhs194 _dafny.Sequence = _1178_v42
			_ = _rhs194
			var _lhs144 *GlobalState = globalState
			_ = _lhs144
			r2 = _rhs192
			_lhs144.F7 = _rhs193
			_1178_v42 = _rhs194
			var _1189_v49 _dafny.Array
			_ = _1189_v49
			var _nw181 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(18))
			_ = _nw181
			_1189_v49 = _nw181
			var _1190_v51 D6
			_ = _1190_v51
			_1190_v51 = Companion_D6_.Create_DC15_(func() _dafny.Map {
				var _coll74 = _dafny.NewMapBuilder()
				_ = _coll74
				for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(327), _dafny.IntOfInt64(336))); ; {
					_compr_74, _ok78 := _iter78()
					if !_ok78 {
						break
					}
					var _1191_v50 _dafny.Int
					_1191_v50 = interface{}(_compr_74).(_dafny.Int)
					if ((_dafny.IntOfInt64(327)).Cmp(_1191_v50) <= 0) && ((_1191_v50).Cmp(_dafny.IntOfInt64(336)) < 0) {
						_coll74.Add(Companion_Default___.SafeDivisionInt(_1191_v50, _1115_v0), _1178_v42)
					}
				}
				return _coll74.ToMap()
			}())
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1189_v49), 0))
			_ = _index161
			(_1189_v49).ArraySet1(_1190_v51, (_index161).Int())
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1189_v49), 0))
			_ = _index162
			(_1189_v49).ArraySet1(Companion_Default___.Fm39(_1115_v0, globalState), (_index162).Int())
			var _1192_v52 _dafny.MultiSet
			_ = _1192_v52
			_1192_v52 = _dafny.MultiSetOf(_1115_v0)
			var _1193_v53 _dafny.Map
			_ = _1193_v53
			_1193_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1137_v17, _1192_v52)
			var _1194_v54 D5
			_ = _1194_v54
			_1194_v54 = Companion_D5_.Create_DC14_((_1140_v20).F39())
			var _1195_v55 _dafny.Map
			_ = _1195_v55
			_1195_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(410), !(p1))
			var _1196_v56 _dafny.Map
			_ = _1196_v56
			_1196_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), (_1140_v20).F39())
			var _1197_v57 _dafny.Map
			_ = _1197_v57
			_1197_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() bool {
				if (_1196_v56).Contains((_1140_v20).F39()) {
					return (_1196_v56).Get((_1140_v20).F39()).(bool)
				}
				return (_1140_v20).F39()
			})())
			var _1198_v58 _dafny.Set
			_ = _1198_v58
			_1198_v58 = _dafny.SetOf(_1192_v52)
			var _1199_v59 _dafny.Map
			_ = _1199_v59
			_1199_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1140_v20).F39(), _1198_v58)
			var _1200_v60 _dafny.Array
			_ = _1200_v60
			var _nwElement0_38 bool = p1
			_ = _nwElement0_38
			var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(27))
			_ = _nw182
			(_nw182).ArraySet1(_nwElement0_38, 0)
			(_nw182).ArraySet1(_dafny.Companion_Sequence_.Contains(_1137_v17, p0), 1)
			(_nw182).ArraySet1((func() bool {
				if (_this).F24() {
					return (_1140_v20).F39()
				}
				return (_this).F24()
			})(), 2)
			(_nw182).ArraySet1((_this).F24(), 3)
			(_nw182).ArraySet1(p1, 4)
			(_nw182).ArraySet1((_1192_v52).IsDisjointFrom(_1192_v52), 5)
			(_nw182).ArraySet1((_1193_v53).Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg85 _dafny.Int) interface{} {
					return coer85(arg85)
				}
			}((func(_1201_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1202_i8 _dafny.Int) _dafny.CodePoint {
					return _1201_p0
				}
			})(p0))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1115_v0), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}((func(_1203_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1204_i8 _dafny.Int) _dafny.CodePoint {
					return _1203_p0
				}
			})(p0)))).Cardinality()))).Uint32(), p0)), 6)
			(_nw182).ArraySet1(p1, 7)
			(_nw182).ArraySet1((_1115_v0).Cmp(_1115_v0) != 0, 8)
			(_nw182).ArraySet1((_1194_v54).Dtor_cf23(), 9)
			(_nw182).ArraySet1((false) == ((_this).F24()), 10)
			(_nw182).ArraySet1((_1140_v20).F39(), 11)
			(_nw182).ArraySet1((_this).F24(), 12)
			(_nw182).ArraySet1((_this).F24(), 13)
			(_nw182).ArraySet1((_this).F24(), 14)
			(_nw182).ArraySet1(p1, 15)
			(_nw182).ArraySet1((_1140_v20).F39(), 16)
			(_nw182).ArraySet1((_this).F24(), 17)
			(_nw182).ArraySet1(!((_this).F24()), 18)
			(_nw182).ArraySet1(p1, 19)
			(_nw182).ArraySet1(p1, 20)
			(_nw182).ArraySet1((func() bool {
				if (_this).F24() {
					return (_1140_v20).F39()
				}
				return (func() bool {
					if (_1195_v55).Contains(_1115_v0) {
						return (_1195_v55).Get(_1115_v0).(bool)
					}
					return (_1140_v20).F39()
				})()
			})(), 21)
			(_nw182).ArraySet1((func() bool {
				if (_1197_v57).Contains(p1) {
					return (_1197_v57).Get(p1).(bool)
				}
				return !(false)
			})(), 22)
			(_nw182).ArraySet1(p1, 23)
			(_nw182).ArraySet1((_1140_v20).F39(), 24)
			(_nw182).ArraySet1((((func() _dafny.Set {
				if (_1199_v59).Contains((_1140_v20).F39()) {
					return (_1199_v59).Get((_1140_v20).F39()).(_dafny.Set)
				}
				return _1198_v58
			})()).Cardinality()).Cmp(_1115_v0) < 0, 25)
			(_nw182).ArraySet1((_1140_v20).F39(), 26)
			_1200_v60 = _nw182
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1200_v60), 0))
			_ = _index163
			(_1200_v60).ArraySet1((_1140_v20).F39(), (_index163).Int())
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1200_v60), 0))
			_ = _index164
			(_1200_v60).ArraySet1(p1, (_index164).Int())
			var _1205_v61 _dafny.Array
			_ = _1205_v61
			var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw183
			_1205_v61 = _nw183
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1205_v61), 0))
			_ = _index165
			(_1205_v61).ArraySet1(_1200_v60, (_index165).Int())
			var _1206_v62 _dafny.Map
			_ = _1206_v62
			_1206_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1200_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1200_v60), 0))).Int()).(bool)), _1200_v60)
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1205_v61), 0))
			_ = _index166
			(_1205_v61).ArraySet1((func() _dafny.Array {
				if (_1206_v62).Contains(p1) {
					return (_1206_v62).Get(p1).(_dafny.Array)
				}
				return _1200_v60
			})(), (_index166).Int())
		}
		r0 = _dafny.Companion_Sequence_.Concatenate(_1137_v17, _1137_v17)
		var _1207_v63 D0
		_ = _1207_v63
		_1207_v63 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1138_v18).Contains(p1) {
				return (_1138_v18).Get(p1).(_dafny.Sequence)
			}
			return _1137_v17
		})()).Cardinality()))
		r1 = (func() D0 {
			if (_this).F24() {
				return _1207_v63
			}
			return _1207_v63
		})()
		r2 = !((_this).F24())
		return r0, r1, r2
	}
}
func (_this *C6) M2(globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _source17 D2 = Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("hin")))
		_ = _source17
		if _source17.Is_DC7() {
			var _1208___mcc_h0 bool = _source17.Get_().(D2_DC7).Cf13
			_ = _1208___mcc_h0
			var _1209___mcc_h1 bool = _source17.Get_().(D2_DC7).Cf14
			_ = _1209___mcc_h1
			var _1210_cf14 bool = _1209___mcc_h1
			_ = _1210_cf14
			var _1211_cf13 bool = _1208___mcc_h0
			_ = _1211_cf13
			var _1212_v0 _dafny.Array
			_ = _1212_v0
			var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
			_ = _nw184
			_1212_v0 = _nw184
			_1212_v0 = _1212_v0
			var _1213_v1 _dafny.Array
			_ = _1213_v1
			var _nw185 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw185
			_1213_v1 = _nw185
			_1213_v1 = _1213_v1
			var _1214_v2 _dafny.Int
			_ = _1214_v2
			_1214_v2 = _dafny.IntOfInt64(486)
			r2 = (_1214_v2).Cmp((_dafny.Zero).Minus(_1214_v2)) <= 0
			(globalState).F7 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(711))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg87 _dafny.Int) interface{} {
					return coer87(arg87)
				}
			}(func(_1215_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-356), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(711))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}(func(_1216_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))).Cardinality()))).Uint32(), _this.F23())).Cardinality())).Times((_dafny.Zero).Minus(_1214_v2))
		} else if _source17.Is_DC6() {
			var _1217___mcc_h2 _dafny.Map = _source17.Get_().(D2_DC6).Cf12
			_ = _1217___mcc_h2
			var _1218_cf12 _dafny.Map = _1217___mcc_h2
			_ = _1218_cf12
			var _1219_v3 _dafny.Sequence
			_ = _1219_v3
			_1219_v3 = _dafny.SeqOf(false)
			var _1220_v4 _dafny.Int
			_ = _1220_v4
			_1220_v4 = _dafny.IntOfInt64(158)
			var _1221_v5 _dafny.Sequence
			_ = _1221_v5
			_1221_v5 = _dafny.SeqOf(_dafny.IntOfUint32((_1219_v3).Cardinality()), _1220_v4, _dafny.IntOfInt64(430))
			var _1222_v6 _dafny.Map
			_ = _1222_v6
			_1222_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1221_v5)
			_1222_v6 = (_1222_v6).Merge(_1222_v6)
			var _1223_v7 _dafny.Array
			_ = _1223_v7
			var _nw186 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw186
			_1223_v7 = _nw186
			var _1224_v8 _dafny.Map
			_ = _1224_v8
			_1224_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1220_v4)
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1223_v7), 0))
			_ = _index167
			(_1223_v7).ArraySet1((_1224_v8).Merge(_1224_v8), (_index167).Int())
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1223_v7), 0))
			_ = _index168
			(_1223_v7).ArraySet1(_1224_v8, (_index168).Int())
			var _1225_v9 _dafny.Sequence
			_ = _1225_v9
			_1225_v9 = _dafny.UnicodeSeqOfUtf8Bytes("urkqv")
			var _1226_v10 _dafny.Map
			_ = _1226_v10
			_1226_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1220_v4, _1220_v4)
			var _1227_v12 *C0
			_ = _1227_v12
			var _nw187 *C0 = New_C0_()
			_ = _nw187
			_nw187.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1225_v9)).Merge(_1218_cf12), ((func() _dafny.Set {
				var _coll75 = _dafny.NewBuilder()
				_ = _coll75
				for _iter79 := _dafny.Iterate((_1226_v10).Keys().Elements()); ; {
					_compr_75, _ok79 := _iter79()
					if !_ok79 {
						break
					}
					var _1228_v11 _dafny.Int
					_1228_v11 = interface{}(_compr_75).(_dafny.Int)
					if (_1226_v10).Contains(_1228_v11) {
						_coll75.Add((_1228_v11).Plus(_dafny.IntOfInt64(476)))
					}
				}
				return _coll75.ToSet()
			}()).Cardinality()).Cmp(_1220_v4) < 0, _this.F23(), true)
			_1227_v12 = _nw187
			_1219_v3 = _1219_v3
		} else {
			var _1229___mcc_h3 D2 = _source17.Get_().(D2_DC8).Cf15
			_ = _1229___mcc_h3
			var _1230_cf15 D2 = _1229___mcc_h3
			_ = _1230_cf15
			var _1231_v13 *C1
			_ = _1231_v13
			var _nw188 *C1 = New_C1_()
			_ = _nw188
			_nw188.Ctor__((_this).F24(), (_this).F24())
			_1231_v13 = _nw188
			var _1232_v14 *C1
			_ = _1232_v14
			_1232_v14 = _1231_v13
			var _1233_v15 _dafny.Sequence
			_ = _1233_v15
			_1233_v15 = _dafny.SeqOf(_1231_v13, _1231_v13, _1231_v13)
			var _1234_v16 _dafny.Int
			_ = _1234_v16
			_1234_v16 = _dafny.IntOfInt64(-984)
			var _source18 *C1 = (_1233_v15).Select((Companion_Default___.SafeIndex(_1234_v16, _dafny.IntOfUint32((_1233_v15).Cardinality()))).Uint32()).(*C1)
			_ = _source18
			var _1235___mcc_h4 *C1 = _source18
			_ = _1235___mcc_h4
			var _1236_cf28 *C1 = _1235___mcc_h4
			_ = _1236_cf28
			var _1237_v17 _dafny.Array
			_ = _1237_v17
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw189
			_1237_v17 = _nw189
			var _1238_v18 _dafny.Array
			_ = _1238_v18
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_36
			var _nw190 _dafny.Array
			_ = _nw190
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw190 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Int = (func(_1239_v16 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1240_i1 _dafny.Int) _dafny.Int {
						return (_1240_i1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1239_v16)).Cardinality()))
					}
				})(_1234_v16)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw190 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw190).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw190).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1238_v18 = _nw190
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1237_v17), 0))
			_ = _index169
			(_1237_v17).ArraySet1(_1238_v18, (_index169).Int())
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1237_v17), 0))
			_ = _index170
			(_1237_v17).ArraySet1(_1238_v18, (_index170).Int())
			var _1241_v19 _dafny.Sequence
			_ = _1241_v19
			_1241_v19 = _dafny.SeqOf(!((_1236_cf28).F39()), (_1231_v13).F39())
			var _1242_v20 _dafny.Sequence
			_ = _1242_v20
			_1242_v20 = _dafny.SeqOf(_1241_v19)
			var _1243_v21 _dafny.Int
			_ = _1243_v21
			var _1244_v22 bool
			_ = _1244_v22
			var _1245_v23 _dafny.Int
			_ = _1245_v23
			var _out37 _dafny.Int
			_ = _out37
			var _out38 bool
			_ = _out38
			var _out39 _dafny.Int
			_ = _out39
			_out37, _out38, _out39 = Companion_Default___.M0(_1234_v16, _dafny.MultiSetOf((_this).F24()), ((_this).F24()) == ((_1236_cf28).F39()), _dafny.IntOfUint32((_1242_v20).Cardinality()), globalState)
			_1243_v21 = _out37
			_1244_v22 = _out38
			_1245_v23 = _out39
			var _1246_v24 _dafny.Array
			_ = _1246_v24
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw191
			_1246_v24 = _nw191
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_1246_v24), 0))
			_ = _index171
			(_1246_v24).ArraySet1(_1244_v22, (_index171).Int())
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1246_v24), 0))
			_ = _index172
			(_1246_v24).ArraySet1(_1244_v22, (_index172).Int())
			var _1247_v25 _dafny.Sequence
			_ = _1247_v25
			_1247_v25 = _dafny.SeqOf(_1243_v21)
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_1246_v24), 0))
			_ = _index173
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1246_v24), 0))
			_ = _index174
			var _rhs195 bool = (_this).F24()
			_ = _rhs195
			var _rhs196 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1242_v20, _dafny.SeqOf(_1241_v19))
			_ = _rhs196
			var _rhs197 bool = (func() bool {
				if _1244_v22 {
					return !_dafny.Companion_Sequence_.Equal(_1247_v25, _1247_v25)
				}
				return true
			})()
			_ = _rhs197
			var _lhs145 _dafny.Array = _1246_v24
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_1246_v24), 0))
			_ = _lhs146
			var _lhs147 _dafny.Array = _1246_v24
			_ = _lhs147
			var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1246_v24), 0))
			_ = _lhs148
			(_lhs145).ArraySet1(_rhs195, (_lhs146).Int())
			_1242_v20 = _rhs196
			(_lhs147).ArraySet1(_rhs197, (_lhs148).Int())
			var _1248_v26 _dafny.Map
			_ = _1248_v26
			_1248_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1244_v22, _1244_v22)
			var _1249_v27 _dafny.Map
			_ = _1249_v27
			_1249_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v19, (_1236_cf28).F39())
			var _1250_v28 _dafny.Map
			_ = _1250_v28
			_1250_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_1248_v26).Contains((_1231_v13).F39()) {
					return (_1248_v26).Get((_1231_v13).F39()).(bool)
				}
				return (func() bool {
					if (_1249_v27).Contains(_1241_v19) {
						return (_1249_v27).Get(_1241_v19).(bool)
					}
					return !((_1231_v13).F39())
				})()
			})(), _1243_v21)
			var _1251_v29 _dafny.Sequence
			_ = _1251_v29
			_1251_v29 = _dafny.SeqOf(_1250_v28, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(globalState), (_1231_v13).Fm33(globalState)))
			_1250_v28 = (_1251_v29).Select((Companion_Default___.SafeIndex(_1243_v21, _dafny.IntOfUint32((_1251_v29).Cardinality()))).Uint32()).(_dafny.Map)
			var _1252_v30 _dafny.Sequence
			_ = _1252_v30
			_1252_v30 = _dafny.UnicodeSeqOfUtf8Bytes("ycfwea")
			var _1253_v31 _dafny.Map
			_ = _1253_v31
			_1253_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F24()), _1252_v30)
			var _1254_v32 D2
			_ = _1254_v32
			_1254_v32 = Companion_D2_.Create_DC6_(_1253_v31)
			var _source19 D2 = _1254_v32
			_ = _source19
			if _source19.Is_DC7() {
				var _1255___mcc_h5 bool = _source19.Get_().(D2_DC7).Cf13
				_ = _1255___mcc_h5
				var _1256___mcc_h6 bool = _source19.Get_().(D2_DC7).Cf14
				_ = _1256___mcc_h6
				var _1257_cf14 bool = _1256___mcc_h6
				_ = _1257_cf14
				var _1258_cf13 bool = _1255___mcc_h5
				_ = _1258_cf13
				var _1259_v33 _dafny.Sequence
				_ = _1259_v33
				_1259_v33 = _dafny.SeqOf(_1258_cf13)
				var _1260_v34 _dafny.Map
				_ = _1260_v34
				_1260_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(globalState), _1259_v33)
				_1260_v34 = (_1260_v34).Update(_1258_cf13, _1259_v33)
				var _1261_v35 _dafny.Map
				_ = _1261_v35
				_1261_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1231_v13).F39(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1252_v30).Cardinality()), (_1231_v13).F39()))
				var _1262_v36 _dafny.Map
				_ = _1262_v36
				_1262_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v16, Companion_Default___.Fm2(globalState))
				_1261_v35 = (_1261_v35).Update((_this).F24(), _1262_v36)
				var _1263_v37 _dafny.Map
				_ = _1263_v37
				_1263_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1260_v34).Cardinality(), _1234_v16)
				var _1264_v38 _dafny.Sequence
				_ = _1264_v38
				_1264_v38 = _dafny.SeqOf(_1234_v16)
				var _1265_v39 _dafny.Set
				_ = _1265_v39
				_1265_v39 = _dafny.SetOf((_this).F24(), (Companion_D5_.Create_DC13_(_1234_v16, (_1231_v13).F39(), (_1259_v33).Select((Companion_Default___.SafeIndex(_1234_v16, _dafny.IntOfUint32((_1259_v33).Cardinality()))).Uint32()).(bool))).Dtor_cf21())
				var _1266_v40 _dafny.Sequence
				_ = _1266_v40
				_1266_v40 = _dafny.SeqOf(_dafny.SetOf((_1231_v13).F39(), (_1231_v13).F39()), _1265_v39)
				var _1267_v41 _dafny.Array
				_ = _1267_v41
				var _nwElement0_39 _dafny.Int = _1234_v16
				_ = _nwElement0_39
				var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(23))
				_ = _nw192
				(_nw192).ArraySet1(_nwElement0_39, 0)
				(_nw192).ArraySet1(_1234_v16, 1)
				(_nw192).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1234_v16), _1234_v16), 2)
				(_nw192).ArraySet1(_1234_v16, 3)
				(_nw192).ArraySet1(_1234_v16, 4)
				(_nw192).ArraySet1(_1234_v16, 5)
				(_nw192).ArraySet1(_1234_v16, 6)
				(_nw192).ArraySet1(_1234_v16, 7)
				(_nw192).ArraySet1(Companion_Default___.SafeModuloInt(_1234_v16, (func() _dafny.Int {
					if (_1263_v37).Contains(_dafny.IntOfUint32((_1264_v38).Cardinality())) {
						return (_1263_v37).Get(_dafny.IntOfUint32((_1264_v38).Cardinality())).(_dafny.Int)
					}
					return _1234_v16
				})()), 8)
				(_nw192).ArraySet1(_1234_v16, 9)
				(_nw192).ArraySet1((_this).Fm6(globalState), 10)
				(_nw192).ArraySet1((_1234_v16).Times(((_1266_v40).Select((Companion_Default___.SafeIndex(_1234_v16, _dafny.IntOfUint32((_1266_v40).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()), 11)
				(_nw192).ArraySet1(_1234_v16, 12)
				(_nw192).ArraySet1((_dafny.Zero).Minus(_1234_v16), 13)
				(_nw192).ArraySet1(_1234_v16, 14)
				(_nw192).ArraySet1(_1234_v16, 15)
				(_nw192).ArraySet1((_1264_v38).Select((Companion_Default___.SafeIndex(_1234_v16, _dafny.IntOfUint32((_1264_v38).Cardinality()))).Uint32()).(_dafny.Int), 16)
				(_nw192).ArraySet1((_1262_v36).Cardinality(), 17)
				(_nw192).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1234_v16)), 18)
				(_nw192).ArraySet1(_1234_v16, 19)
				(_nw192).ArraySet1(_1234_v16, 20)
				(_nw192).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1258_cf13, _1234_v16)).Cardinality(), _1234_v16)), 21)
				(_nw192).ArraySet1((_1234_v16).Plus(_dafny.IntOfUint32((_1252_v30).Cardinality())), 22)
				_1267_v41 = _nw192
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1267_v41), 0))
				_ = _index175
				(_1267_v41).ArraySet1((_1263_v37).Cardinality(), (_index175).Int())
				var _1268_v42 _dafny.MultiSet
				_ = _1268_v42
				_1268_v42 = _dafny.MultiSetOf(_this.F23())
				var _1269_v43 _dafny.Map
				_ = _1269_v43
				_1269_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1258_cf13, false)
				var _1270_v44 _dafny.MultiSet
				_ = _1270_v44
				_1270_v44 = _dafny.MultiSetOf(_1269_v43, _1269_v43, _1269_v43)
				var _1271_v45 _dafny.MultiSet
				_ = _1271_v45
				_1271_v45 = _dafny.MultiSetOf(_1270_v44, _1270_v44, (_dafny.MultiSetOf(_1269_v43)).Update(_1269_v43, Companion_Default___.Abs(_1234_v16)), _1270_v44, _1270_v44)
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1267_v41), 0))
				_ = _index176
				var _rhs198 _dafny.Int = _1234_v16
				_ = _rhs198
				var _rhs199 _dafny.Int = (func() _dafny.Int {
					if (_1268_v42).Contains(_this.F23()) {
						return (_1268_v42).Multiplicity(_this.F23())
					}
					return (func() _dafny.Int {
						if (_1271_v45).Contains(_1270_v44) {
							return (_1271_v45).Multiplicity(_1270_v44)
						}
						return _1234_v16
					})()
				})()
				_ = _rhs199
				var _rhs200 _dafny.Int = ((_dafny.IntOfInt64(821)).Times(Companion_Default___.Fm0((_1231_v13).F39(), _1234_v16, globalState))).Times(_1234_v16)
				_ = _rhs200
				var _lhs149 _dafny.Array = _1267_v41
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1267_v41), 0))
				_ = _lhs150
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				(_lhs149).ArraySet1(_rhs198, (_lhs150).Int())
				_lhs151.F7 = _rhs199
				_lhs152.F7 = _rhs200
				(globalState).F7 = Companion_Default___.Fm0((_1231_v13).F39(), _1234_v16, globalState)
			} else if _source19.Is_DC6() {
				var _1272___mcc_h7 _dafny.Map = _source19.Get_().(D2_DC6).Cf12
				_ = _1272___mcc_h7
				var _1273_cf12 _dafny.Map = _1272___mcc_h7
				_ = _1273_cf12
				var _1274_v46 _dafny.MultiSet
				_ = _1274_v46
				_1274_v46 = _dafny.MultiSetOf((_1231_v13).F39(), (_this).F24())
				_1274_v46 = _1274_v46
				var _1275_v48 _dafny.Map
				_ = _1275_v48
				_1275_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll76 = _dafny.NewMapBuilder()
					_ = _coll76
					for _iter80 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(704), _dafny.IntOfInt64(468))); ; {
						_compr_76, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _1276_v47 _dafny.Int
						_1276_v47 = interface{}(_compr_76).(_dafny.Int)
						if ((_dafny.IntOfInt64(704)).Cmp(_1276_v47) <= 0) && ((_1276_v47).Cmp(_dafny.IntOfInt64(468)) < 0) {
							_coll76.Add((_1276_v47).Times(_1234_v16), _1234_v16)
						}
					}
					return _coll76.ToMap()
				}(), _dafny.IntOfInt64(39))
				var _1277_v49 _dafny.MultiSet
				_ = _1277_v49
				_1277_v49 = _dafny.MultiSetOf(_1234_v16, _1234_v16, _1234_v16)
				var _1278_v50 _dafny.Set
				_ = _1278_v50
				_1278_v50 = _dafny.SetOf((_1231_v13).F39())
				var _nwElement0_40 _dafny.Int = _1234_v16
				_ = _nwElement0_40
				var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(19))
				_ = _nw193
				(_nw193).ArraySet1(_nwElement0_40, 0)
				(_nw193).ArraySet1(_dafny.IntOfInt64(31), 1)
				(_nw193).ArraySet1((_1234_v16).Plus(_1234_v16), 2)
				(_nw193).ArraySet1((_1275_v48).Cardinality(), 3)
				(_nw193).ArraySet1(_1234_v16, 4)
				(_nw193).ArraySet1((func() _dafny.Int {
					if (_1277_v49).Contains(_1234_v16) {
						return (_1277_v49).Multiplicity(_1234_v16)
					}
					return _1234_v16
				})(), 5)
				(_nw193).ArraySet1(_1234_v16, 6)
				(_nw193).ArraySet1(_1234_v16, 7)
				(_nw193).ArraySet1((_1234_v16).Minus(_1234_v16), 8)
				(_nw193).ArraySet1(_1234_v16, 9)
				(_nw193).ArraySet1(Companion_Default___.SafeDivisionInt(_1234_v16, _1234_v16), 10)
				(_nw193).ArraySet1((_1234_v16).Plus(Companion_Default___.Fm0((_1231_v13).F39(), _1234_v16, globalState)), 11)
				(_nw193).ArraySet1(_1234_v16, 12)
				(_nw193).ArraySet1(_1234_v16, 13)
				(_nw193).ArraySet1((_1234_v16).Plus((_dafny.Zero).Minus(_1234_v16)), 14)
				(_nw193).ArraySet1(_1234_v16, 15)
				(_nw193).ArraySet1(_1234_v16, 16)
				(_nw193).ArraySet1(Companion_Default___.Fm0((_this).F24(), (_1278_v50).Cardinality(), globalState), 17)
				(_nw193).ArraySet1(_1234_v16, 18)
				(globalState).F22 = _nw193
				(globalState).F7 = Companion_Default___.SafeModuloInt(_1234_v16, _1234_v16)
				var _1279_v51 _dafny.Map
				_ = _1279_v51
				_1279_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(105), _1234_v16)
				var _1280_v52 _dafny.Array
				_ = _1280_v52
				var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
				_ = _nw194
				_1280_v52 = _nw194
				var _1281_v53 D12
				_ = _1281_v53
				_1281_v53 = Companion_D12_.Create_DC24_(_1280_v52)
				var _1282_v54 D0
				_ = _1282_v54
				_1282_v54 = Companion_D0_.Create_DC1_((_1274_v46).Cardinality())
				var _1283_v55 *C4
				_ = _1283_v55
				var _nw195 *C4 = New_C4_()
				_ = _nw195
				_nw195.Ctor__((_1234_v16).Minus(_1234_v16), _1279_v51, (_1281_v53).Dtor_cf40(), Companion_Default___.Fm34((_1282_v54).Dtor_cf1(), false, globalState), (_1231_v13).F39())
				_1283_v55 = _nw195
			} else {
				var _1284___mcc_h8 D2 = _source19.Get_().(D2_DC8).Cf15
				_ = _1284___mcc_h8
				var _1285_cf15 D2 = _1284___mcc_h8
				_ = _1285_cf15
				var _1286_v56 _dafny.Map
				_ = _1286_v56
				_1286_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1231_v13).F39(), (_1231_v13).F39())
				_1286_v56 = (_1286_v56).Update((_1231_v13).F39(), (_1231_v13).F39())
				var _1287_v59 _dafny.Map
				_ = _1287_v59
				_1287_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll77 = _dafny.NewMapBuilder()
					_ = _coll77
					for _iter81 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(183), _dafny.IntOfInt64(311))); ; {
						_compr_77, _ok81 := _iter81()
						if !_ok81 {
							break
						}
						var _1288_v57 _dafny.Int
						_1288_v57 = interface{}(_compr_77).(_dafny.Int)
						if ((_dafny.IntOfInt64(183)).Cmp(_1288_v57) <= 0) && ((_1288_v57).Cmp(_dafny.IntOfInt64(311)) < 0) {
							_coll77.Add((_1288_v57).Times((func() _dafny.Map {
								var _coll78 = _dafny.NewMapBuilder()
								_ = _coll78
								for _iter82 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-155), _dafny.IntOfInt64(293))); ; {
									_compr_78, _ok82 := _iter82()
									if !_ok82 {
										break
									}
									var _1289_v58 _dafny.Int
									_1289_v58 = interface{}(_compr_78).(_dafny.Int)
									if ((_dafny.IntOfInt64(-155)).Cmp(_1289_v58) <= 0) && ((_1289_v58).Cmp(_dafny.IntOfInt64(293)) < 0) {
										_coll78.Add(Companion_Default___.SafeDivisionInt(_1289_v58, _1234_v16), _1286_v56)
									}
								}
								return _coll78.ToMap()
							}()).Cardinality()), _1234_v16)
						}
					}
					return _coll77.ToMap()
				}()).Cardinality(), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1((_this).F24(), globalState), (Companion_Default___.SafeIndex(_1234_v16, _dafny.IntOfUint32((Companion_Default___.Fm1((_this).F24(), globalState)).Cardinality()))).Uint32(), _this.F23()))
				_1252_v30 = (func() _dafny.Sequence {
					if (_1287_v59).Contains(_1234_v16) {
						return (_1287_v59).Get(_1234_v16).(_dafny.Sequence)
					}
					return _1252_v30
				})()
				_1252_v30 = _1252_v30
				var _1290_v60 _dafny.Array
				_ = _1290_v60
				var _nw196 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw196
				_1290_v60 = _nw196
				var _1291_v61 *C5
				_ = _1291_v61
				var _nw197 *C5 = New_C5_()
				_ = _nw197
				_nw197.Ctor__(_1290_v60)
				_1291_v61 = _nw197
			}
			(globalState).F13 = (_this).F24()
			var _1292_v62 _dafny.MultiSet
			_ = _1292_v62
			_1292_v62 = _dafny.MultiSetOf(_1234_v16, _1234_v16)
			var _1293_v63 _dafny.Map
			_ = _1293_v63
			_1293_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1234_v16).Cmp((_dafny.Zero).Minus((_1292_v62).Cardinality())) < 0, _1232_v14)
			_1293_v63 = (_1293_v63).Update(((_1231_v13).F39()) == ((_this).F24()), _1232_v14)
		}
		var _1294_v64 _dafny.Map
		_ = _1294_v64
		_1294_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
		var _1295_v65 _dafny.Array
		_ = _1295_v65
		var _nwElement0_41 _dafny.Map = (_1294_v64).Merge(_1294_v64)
		_ = _nwElement0_41
		var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(24))
		_ = _nw198
		(_nw198).ArraySet1(_nwElement0_41, 0)
		(_nw198).ArraySet1((_1294_v64).Merge(_1294_v64), 1)
		(_nw198).ArraySet1((Companion_Default___.Fm19(globalState)).Merge(_1294_v64), 2)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F24()), (_this).F24()), 3)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm16(globalState), (_this).F24()), 4)
		(_nw198).ArraySet1(_1294_v64, 5)
		(_nw198).ArraySet1(_1294_v64, 6)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()), 7)
		(_nw198).ArraySet1(_1294_v64, 8)
		(_nw198).ArraySet1(_1294_v64, 9)
		(_nw198).ArraySet1(_1294_v64, 10)
		(_nw198).ArraySet1((_1294_v64).Merge(_1294_v64), 11)
		(_nw198).ArraySet1(_1294_v64, 12)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()), 13)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), false), 14)
		(_nw198).ArraySet1((_1294_v64).Merge(_1294_v64), 15)
		(_nw198).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), !((_this).F24()))).Update((_this).F24(), (_this).F24()), 16)
		(_nw198).ArraySet1(_1294_v64, 17)
		(_nw198).ArraySet1(_1294_v64, 18)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()), 19)
		(_nw198).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()), 20)
		(_nw198).ArraySet1((_1294_v64).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())), 21)
		(_nw198).ArraySet1(_1294_v64, 22)
		(_nw198).ArraySet1((_1294_v64).Merge(_1294_v64), 23)
		_1295_v65 = _nw198
		for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1295_v65), 0))); ; {
			_guard_loop_3, _ok83 := _iter83()
			if !_ok83 {
				break
			}
			var _1296_i2 _dafny.Int
			_1296_i2 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1296_i2).Sign() != -1) && ((_1296_i2).Cmp(_dafny.ArrayLen((_1295_v65), 0)) < 0)) {
				(_1295_v65).ArraySet1((_1294_v64).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())), (_1296_i2).Int())
			}
		}
		var _1297_v66 _dafny.Int
		_ = _1297_v66
		_1297_v66 = _dafny.IntOfInt64(590)
		var _1298_v67 _dafny.Map
		_ = _1298_v67
		_1298_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1297_v66, _1297_v66)
		var _1299_v68 _dafny.Map
		_ = _1299_v68
		_1299_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1297_v66, _1298_v67)
		var _1300_v69 _dafny.Sequence
		_ = _1300_v69
		_1300_v69 = _dafny.SeqOf(((func() _dafny.Map {
			if (_1299_v68).Contains(_1297_v66) {
				return (_1299_v68).Get(_1297_v66).(_dafny.Map)
			}
			return _1298_v67
		})()).Cardinality(), _1297_v66, _1297_v66)
		_1300_v69 = _dafny.Companion_Sequence_.Concatenate(_1300_v69, _1300_v69)
		r1 = _1297_v66
		var _1301_v70 _dafny.Sequence
		_ = _1301_v70
		_1301_v70 = _dafny.SeqOf(!(true))
		var _1302_v71 _dafny.Array
		_ = _1302_v71
		var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
		_ = _nw199
		_1302_v71 = _nw199
		var _1303_v72 _dafny.MultiSet
		_ = _1303_v72
		_1303_v72 = _dafny.MultiSetOf(_this.F23())
		var _1304_v73 _dafny.Set
		_ = _1304_v73
		_1304_v73 = _dafny.SetOf((func() _dafny.Int {
			if (_1303_v72).Contains(_this.F23()) {
				return (_1303_v72).Multiplicity(_this.F23())
			}
			return (_this).Fm6(globalState)
		})(), _dafny.IntOfInt64(-668))
		var _1305_v74 *C2
		_ = _1305_v74
		var _nw200 *C2 = New_C2_()
		_ = _nw200
		_nw200.Ctor__(_dafny.MultiSetFromSeq(_1301_v70), (func() _dafny.Array {
			if (_this).F24() {
				return _1302_v71
			}
			return _1302_v71
		})(), _this.F23(), ((_1304_v73).Cardinality()).Cmp(_1297_v66) < 0)
		_1305_v74 = _nw200
		r0 = true
		r0 = (_this).F24()
		r1 = _1297_v66
		r2 = (_this).F24()
		return r0, r1, r2
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f23 _dafny.CodePoint
	_f24 bool
	_f25 _dafny.Array
	_f27 bool
	_f30 bool
	_f31 bool
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
	_this._f25 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f27 = false
	_this._f30 = false
	_this._f31 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T3_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C7{}
var _ T3 = &C7{}
var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C7) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C7) F24() bool {
	return _this._f24
}
func (_this *C7) F25() _dafny.Array {
	return _this._f25
}
func (_this *C7) F27() bool {
	return _this._f27
}
func (_this *C7) Ctor__(f30 bool, f31 bool, f25 _dafny.Array, f23 _dafny.CodePoint, f24 bool, f27 bool) {
	{
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f25 = f25
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f27 = f27
	}
}
func (_this *C7) Fm7(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(564)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}(func(_1306_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		}))).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F24()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!((_this).F30()))).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg90 _dafny.Int) interface{} {
				return coer90(arg90)
			}
		}(func(_1307_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(802)
		}))))
	}
}
func (_this *C7) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg91 _dafny.Int) interface{} {
				return coer91(arg91)
			}
		}(func(_1308_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		}))))).Dtor_cf12()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("y"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("sufkeaph"))))
	}
}
func (_this *C7) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(252)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(!(!((_this).F30())))).Cardinality())) > 0
	}
}
func (_this *C7) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if (_this).F30() {
				return _dafny.IntOfInt64(-200)
			}
			return _dafny.IntOfInt64(-98)
		})())), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(891))).Cardinality())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ych")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-903), _dafny.IntOfInt64(-298))).Cardinality())).Cardinality()))
	}
}
func (_this *C7) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf((func() _dafny.Sequence {
			if (_this).F31() {
				return _dafny.SeqOf((_this).F27())
			}
			return _dafny.SeqOf(true, (_this).F24(), (_this).F24())
		})())
	}
}
func (_this *C7) M3(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1309_v0 _dafny.Array
		_ = _1309_v0
		var _len0_37 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_37
		var _nw201 _dafny.Array
		_ = _nw201
		if _len0_37.Cmp(_dafny.Zero) == 0 {
			_nw201 = _dafny.NewArray(_len0_37)
		} else {
			var _init37 func(_dafny.Int) bool = func(_1310_i0 _dafny.Int) bool {
				return (_this).F30()
			}
			_ = _init37
			var _element0_37 = _init37(_dafny.Zero)
			_ = _element0_37
			_nw201 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
			(_nw201).ArraySet1(_element0_37, 0)
			var _nativeLen0_37 = (_len0_37).Int()
			_ = _nativeLen0_37
			for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
				(_nw201).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
			}
		}
		_1309_v0 = _nw201
		var _1311_v1 _dafny.Int
		_ = _1311_v1
		_1311_v1 = _dafny.IntOfInt64(235)
		var _1312_v2 _dafny.Map
		_ = _1312_v2
		_1312_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1313_v3 _dafny.Array
		_ = _1313_v3
		var _nwElement0_42 _dafny.Int = _1311_v1
		_ = _nwElement0_42
		var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(10))
		_ = _nw202
		(_nw202).ArraySet1(_nwElement0_42, 0)
		(_nw202).ArraySet1(_dafny.IntOfInt64(441), 1)
		(_nw202).ArraySet1(_dafny.IntOfInt64(-229), 2)
		(_nw202).ArraySet1(_1311_v1, 3)
		(_nw202).ArraySet1(_1311_v1, 4)
		(_nw202).ArraySet1(_1311_v1, 5)
		(_nw202).ArraySet1(_1311_v1, 6)
		(_nw202).ArraySet1(_dafny.IntOfInt64(459), 7)
		(_nw202).ArraySet1(_dafny.IntOfInt64(855), 8)
		(_nw202).ArraySet1((_1312_v2).Cardinality(), 9)
		_1313_v3 = _nw202
		var _1314_v4 _dafny.Sequence
		_ = _1314_v4
		_1314_v4 = _dafny.SeqOf(_1313_v3)
		var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))
		_ = _index177
		(_1309_v0).ArraySet1(_dafny.Companion_Sequence_.Equal(_1314_v4, _1314_v4), (_index177).Int())
		var _1315_v5 _dafny.Map
		_ = _1315_v5
		_1315_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _1316_v6 _dafny.Map
		_ = _1316_v6
		_1316_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1315_v5).Cardinality(), _1311_v1)
		var _1317_v7 _dafny.Map
		_ = _1317_v7
		_1317_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1316_v6)
		var _1318_v8 _dafny.Map
		_ = _1318_v8
		_1318_v8 = (func() _dafny.Map {
			if (_1317_v7).Contains((_this).F31()) {
				return (_1317_v7).Get((_this).F31()).(_dafny.Map)
			}
			return _1316_v6
		})()
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))
		_ = _index178
		(_1309_v0).ArraySet1(func(_source20 _dafny.Map) bool {
			var _1319___mcc_h0 _dafny.Map = _source20
			_ = _1319___mcc_h0
			var _1320_cf16 _dafny.Map = _1319___mcc_h0
			_ = _1320_cf16
			return !((_this).F24()) || ((_this).F30())
		}(_1318_v8), (_index178).Int())
		var _1321_v9 _dafny.Array
		_ = _1321_v9
		var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
		_ = _nw203
		_1321_v9 = _nw203
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1321_v9), 0))
		_ = _index179
		(_1321_v9).ArraySet1(_1309_v0, (_index179).Int())
		var _1322_v10 _dafny.Map
		_ = _1322_v10
		_1322_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v1, false)
		var _1323_v11 _dafny.Map
		_ = _1323_v11
		_1323_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1322_v10, _1309_v0)
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1321_v9), 0))
		_ = _index180
		(_1321_v9).ArraySet1((func() _dafny.Array {
			if (_1323_v11).Contains(_1322_v10) {
				return (_1323_v11).Get(_1322_v10).(_dafny.Array)
			}
			return _1309_v0
		})(), (_index180).Int())
		var _1324_v12 _dafny.Array
		_ = _1324_v12
		var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw204
		_1324_v12 = _nw204
		var _1325_v13 _dafny.Array
		_ = _1325_v13
		var _nwElement0_43 _dafny.Array = _1313_v3
		_ = _nwElement0_43
		var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(14))
		_ = _nw205
		(_nw205).ArraySet1(_nwElement0_43, 0)
		(_nw205).ArraySet1(_1313_v3, 1)
		(_nw205).ArraySet1(_1313_v3, 2)
		(_nw205).ArraySet1(_1313_v3, 3)
		(_nw205).ArraySet1(_1313_v3, 4)
		(_nw205).ArraySet1(_1313_v3, 5)
		(_nw205).ArraySet1(_1313_v3, 6)
		(_nw205).ArraySet1(_1313_v3, 7)
		(_nw205).ArraySet1(_1313_v3, 8)
		(_nw205).ArraySet1(_1313_v3, 9)
		(_nw205).ArraySet1(_1313_v3, 10)
		(_nw205).ArraySet1(_1313_v3, 11)
		(_nw205).ArraySet1(_1313_v3, 12)
		(_nw205).ArraySet1(_1313_v3, 13)
		_1325_v13 = _nw205
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1324_v12), 0))
		_ = _index181
		(_1324_v12).ArraySet1(_1325_v13, (_index181).Int())
		var _1326_v14 T1
		_ = _1326_v14
		var _nw206 *C6 = New_C6_()
		_ = _nw206
		_nw206.Ctor__(_this.F23(), true)
		_1326_v14 = _nw206
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1324_v12), 0))
		_ = _index182
		var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))
		_ = _index183
		var _rhs201 _dafny.Array = (func() _dafny.Array {
			if (_this).F30() {
				return (func() _dafny.Array {
					if (_this).F30() {
						return _1325_v13
					}
					return _1325_v13
				})()
			}
			return _1325_v13
		})()
		_ = _rhs201
		var _rhs202 bool = (_1311_v1).Cmp(_dafny.IntOfInt64(72)) < 0
		_ = _rhs202
		var _rhs203 T1 = _1326_v14
		_ = _rhs203
		var _lhs153 _dafny.Array = _1324_v12
		_ = _lhs153
		var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1324_v12), 0))
		_ = _lhs154
		var _lhs155 _dafny.Array = _1309_v0
		_ = _lhs155
		var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))
		_ = _lhs156
		(_lhs153).ArraySet1(_rhs201, (_lhs154).Int())
		(_lhs155).ArraySet1(_rhs202, (_lhs156).Int())
		_1326_v14 = _rhs203
		var _1327_v15 _dafny.Set
		_ = _1327_v15
		_1327_v15 = _dafny.SetOf(_1311_v1)
		var _1328_v16 D5
		_ = _1328_v16
		_1328_v16 = Companion_D5_.Create_DC12_((_1327_v15).Union(_1327_v15))
		var _source21 D5 = _1328_v16
		_ = _source21
		if _source21.Is_DC13() {
			var _1329___mcc_h1 _dafny.Int = _source21.Get_().(D5_DC13).Cf20
			_ = _1329___mcc_h1
			var _1330___mcc_h2 bool = _source21.Get_().(D5_DC13).Cf21
			_ = _1330___mcc_h2
			var _1331___mcc_h3 bool = _source21.Get_().(D5_DC13).Cf22
			_ = _1331___mcc_h3
			var _1332_cf22 bool = _1331___mcc_h3
			_ = _1332_cf22
			var _1333_cf21 bool = _1330___mcc_h2
			_ = _1333_cf21
			var _1334_cf20 _dafny.Int = _1329___mcc_h1
			_ = _1334_cf20
			var _1335_v17 _dafny.MultiSet
			_ = _1335_v17
			_1335_v17 = _dafny.MultiSetOf(p0)
			r1 = Companion_Default___.SafeModuloInt(_1311_v1, (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(650))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}(func(_1336_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(48)
			})), (Companion_Default___.SafeIndex(_1334_cf20, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(650))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}(func(_1337_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(48)
			}))).Cardinality()))).Uint32(), (_dafny.SetOf(_1334_cf20, (_1335_v17).Cardinality(), _1334_cf20)).Cardinality())).Cardinality())).Times((_1326_v14).Fm6(globalState)))
			if !((_this).F24()) {
				_1313_v3 = _1313_v3
				var _1338_v18 *C1
				_ = _1338_v18
				var _nw207 *C1 = New_C1_()
				_ = _nw207
				_nw207.Ctor__(!((func() bool {
					if (_1326_v14).F24() {
						return _1333_cf21
					}
					return _1333_cf21
				})()), p0)
				_1338_v18 = _nw207
				var _1339_v19 D12
				_ = _1339_v19
				_1339_v19 = Companion_D12_.Create_DC25_(((_this).F24()) && ((_1326_v14).F24()), _1326_v14.F23())
				_1339_v19 = _1339_v19
				var _1340_v20 _dafny.Sequence
				_ = _1340_v20
				_1340_v20 = _dafny.UnicodeSeqOfUtf8Bytes("rwgqydm")
				var _1341_v21 _dafny.Set
				_ = _1341_v21
				_1341_v21 = _dafny.SetOf((_1326_v14).F24())
				var _1342_v22 _dafny.Map
				_ = _1342_v22
				_1342_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1340_v20, _1341_v21)
				_1342_v22 = (_1342_v22).Update(_1340_v20, _1341_v21)
				_1332_cf22 = (_this).F31()
			} else {
				(globalState).F22 = _1313_v3
				var _1343_v23 _dafny.Sequence
				_ = _1343_v23
				_1343_v23 = _dafny.UnicodeSeqOfUtf8Bytes("jlmrjcjpx")
				var _1344_v24 D5
				_ = _1344_v24
				_1344_v24 = Companion_D5_.Create_DC14_(Companion_Default___.Fm2(globalState))
				var _1345_v25 _dafny.Map
				_ = _1345_v25
				_1345_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1344_v24, Companion_Default___.Fm0(_1333_cf21, _1334_cf20, globalState))
				var _1346_v26 _dafny.Map
				_ = _1346_v26
				_1346_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_1343_v23, _1326_v14.F23()), _1345_v25)
				var _1347_v28 _dafny.Set
				_ = _1347_v28
				_1347_v28 = _dafny.SetOf(_1344_v24)
				_1346_v26 = (_1346_v26).Update(_1333_cf21, func() _dafny.Map {
					var _coll79 = _dafny.NewMapBuilder()
					_ = _coll79
					for _iter84 := _dafny.Iterate((_1347_v28).Elements()); ; {
						_compr_79, _ok84 := _iter84()
						if !_ok84 {
							break
						}
						var _1348_v27 D5
						_1348_v27 = interface{}(_compr_79).(D5)
						if (_1347_v28).Contains(_1348_v27) {
							_coll79.Add(_1348_v27, _1334_cf20)
						}
					}
					return _coll79.ToMap()
				}())
				_1334_cf20 = _1334_cf20
				var _1349_v29 _dafny.Sequence
				_ = _1349_v29
				_1349_v29 = _dafny.SeqOf(_1311_v1, _1334_cf20, _1334_cf20, _1311_v1)
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1313_v3), 0))
				_ = _index184
				(_1313_v3).ArraySet1(_dafny.IntOfUint32((_1349_v29).Cardinality()), (_index184).Int())
				var _1350_v30 _dafny.Map
				_ = _1350_v30
				_1350_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v1, _1313_v3)
				var _1351_v31 _dafny.Map
				_ = _1351_v31
				_1351_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1312_v2, _1333_cf21)
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1313_v3), 0))
				_ = _index185
				var _rhs204 bool = (_1311_v1).Cmp((_dafny.Zero).Minus(_1334_cf20)) == 0
				_ = _rhs204
				var _rhs205 _dafny.Int = (_1334_cf20).Plus(Companion_Default___.Fm0(_1333_cf21, _1311_v1, globalState))
				_ = _rhs205
				var _rhs206 _dafny.Array = (func() _dafny.Array {
					if (_1350_v30).Contains((_1351_v31).Cardinality()) {
						return (_1350_v30).Get((_1351_v31).Cardinality()).(_dafny.Array)
					}
					return (func() _dafny.Array {
						if false {
							return _1313_v3
						}
						return _1313_v3
					})()
				})()
				_ = _rhs206
				var _rhs207 _dafny.Int = (_dafny.Zero).Minus(_1311_v1)
				_ = _rhs207
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 _dafny.Array = _1313_v3
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1313_v3), 0))
				_ = _lhs161
				_lhs157.F13 = _rhs204
				_lhs158.F7 = _rhs205
				_lhs159.F22 = _rhs206
				(_lhs160).ArraySet1(_rhs207, (_lhs161).Int())
				_1333_cf21 = (_1334_cf20).Cmp((_1313_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1313_v3), 0))).Int()).(_dafny.Int)) != 0
			}
			_1333_cf21 = (_this).F30()
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))
			_ = _index186
			(_1309_v0).ArraySet1((_1309_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))).Int()).(bool), (_index186).Int())
		} else if _source21.Is_DC14() {
			var _1352___mcc_h4 bool = _source21.Get_().(D5_DC14).Cf23
			_ = _1352___mcc_h4
			var _1353_cf23 bool = _1352___mcc_h4
			_ = _1353_cf23
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1313_v3), 0))
			_ = _index187
			(_1313_v3).ArraySet1(Companion_Default___.SafeModuloInt(_1311_v1, Companion_Default___.Fm0((_this).F24(), _dafny.IntOfInt64(262), globalState)), (_index187).Int())
			var _1354_v32 _dafny.MultiSet
			_ = _1354_v32
			_1354_v32 = _dafny.MultiSetOf((_this).F24())
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1313_v3), 0))
			_ = _index188
			(_1313_v3).ArraySet1((func() _dafny.Int {
				if (_1354_v32).Contains(false) {
					return (_1354_v32).Multiplicity(false)
				}
				return _1311_v1
			})(), (_index188).Int())
			var _1355_v33 _dafny.Sequence
			_ = _1355_v33
			_1355_v33 = _dafny.UnicodeSeqOfUtf8Bytes("dqbli")
			_1355_v33 = _dafny.Companion_Sequence_.Concatenate(_1355_v33, _1355_v33)
			var _1356_v34 _dafny.Array
			_ = _1356_v34
			var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
			_ = _nw208
			_1356_v34 = _nw208
			var _1357_v35 _dafny.Map
			_ = _1357_v35
			_1357_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1313_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1313_v3), 0))).Int()).(_dafny.Int), (_this).F25())
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1356_v34), 0))
			_ = _index189
			(_1356_v34).ArraySet1(_1357_v35, (_index189).Int())
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1356_v34), 0))
			_ = _index190
			(_1356_v34).ArraySet1(_1357_v35, (_index190).Int())
			_1355_v33 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1((_1309_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))).Int()).(bool), globalState), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_1313_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1313_v3), 0))).Int()).(_dafny.Int), _1311_v1), _dafny.IntOfUint32((Companion_Default___.Fm1((_1309_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))).Int()).(bool), globalState)).Cardinality()))).Uint32(), _1326_v14.F23())
		} else {
			var _1358___mcc_h5 _dafny.Set = _source21.Get_().(D5_DC12).Cf19
			_ = _1358___mcc_h5
			var _1359_cf19 _dafny.Set = _1358___mcc_h5
			_ = _1359_cf19
			var _1360_v36 _dafny.Sequence
			_ = _1360_v36
			_1360_v36 = _dafny.SeqOf(_dafny.ArrayCastTo((_1321_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1321_v9), 0))).Int())), _1309_v0, _dafny.ArrayCastTo((_1321_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1321_v9), 0))).Int())), _dafny.ArrayCastTo((_1321_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1321_v9), 0))).Int())))
			_1360_v36 = _dafny.SeqOf(_1309_v0, (func() _dafny.Array {
				if (_1326_v14).F24() {
					return (_1360_v36).Select((Companion_Default___.SafeIndex(_1311_v1, _dafny.IntOfUint32((_1360_v36).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _1309_v0
			})(), _dafny.ArrayCastTo((_1321_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1321_v9), 0))).Int())))
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))
			_ = _index191
			(_1309_v0).ArraySet1((_1326_v14).F24(), (_index191).Int())
			var _1361_v37 _dafny.Sequence
			_ = _1361_v37
			_1361_v37 = _dafny.SeqOf(_1311_v1)
			_1361_v37 = _1361_v37
			var _1362_v38 _dafny.Sequence
			_ = _1362_v38
			_1362_v38 = _dafny.SeqOf((_this).F24(), !((_this).F31()), (_1326_v14).F24(), (_this).F24())
			var _1363_v39 D4
			_ = _1363_v39
			_1363_v39 = Companion_D4_.Create_DC11_(_1311_v1)
			var _rhs208 bool = (_this).Fm4((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dupgk")).Cardinality())).Times(_1311_v1), _1362_v38, globalState)
			_ = _rhs208
			var _rhs209 _dafny.Int = (_1363_v39).Dtor_cf18()
			_ = _rhs209
			var _rhs210 bool = (_1362_v38).Select((Companion_Default___.SafeIndex(_1311_v1, _dafny.IntOfUint32((_1362_v38).Cardinality()))).Uint32()).(bool)
			_ = _rhs210
			var _rhs211 bool = p0
			_ = _rhs211
			var _lhs162 *GlobalState = globalState
			_ = _lhs162
			var _lhs163 *GlobalState = globalState
			_ = _lhs163
			_lhs162.F13 = _rhs208
			_1311_v1 = _rhs209
			_lhs163.F13 = _rhs210
			r2 = _rhs211
		}
		_1316_v6 = (_1316_v6).Update(_1311_v1, _1311_v1)
		if (func() bool {
			if (_1309_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))).Int()).(bool) {
				return (func() bool {
					if (_1322_v10).Contains(_dafny.IntOfInt64(-678)) {
						return (_1322_v10).Get(_dafny.IntOfInt64(-678)).(bool)
					}
					return (_1326_v14).F24()
				})()
			}
			return p0
		})() {
			(globalState).F7 = _1311_v1
			(_this).F23_set_((func() _dafny.CodePoint {
				if (_this).F31() {
					return _1326_v14.F23()
				}
				return _1326_v14.F23()
			})())
			var _1364_v40 _dafny.Map
			_ = _1364_v40
			_1364_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1326_v14).F24(), _1311_v1)
			var _1365_v41 _dafny.Sequence
			_ = _1365_v41
			_1365_v41 = _dafny.SeqOf((_1364_v40).Cardinality())
			r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1365_v41, _1365_v41)).Cardinality())
			var _1366_v42 _dafny.MultiSet
			_ = _1366_v42
			_1366_v42 = _dafny.MultiSetOf(Companion_Default___.Fm2(globalState))
			r1 = (((_1366_v42).Union((_1366_v42).Update(p0, Companion_Default___.Abs(_1311_v1)))).Cardinality()).Plus((_dafny.Zero).Minus(_1311_v1))
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1313_v3), 0))
			_ = _index192
			(_1313_v3).ArraySet1(_dafny.IntOfInt64(342), (_index192).Int())
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1313_v3), 0))
			_ = _index193
			(_1313_v3).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_1311_v1, _1311_v1), _1311_v1), (_index193).Int())
		} else {
			var _1367_v43 D5
			_ = _1367_v43
			_1367_v43 = Companion_D5_.Create_DC14_((_1326_v14).F24())
			_1367_v43 = _1367_v43
			var _1368_v44 _dafny.Sequence
			_ = _1368_v44
			_1368_v44 = _dafny.UnicodeSeqOfUtf8Bytes("qy")
			var _1369_v45 _dafny.Map
			_ = _1369_v45
			_1369_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1368_v44)
			var _1370_v46 _dafny.Sequence
			_ = _1370_v46
			_1370_v46 = _dafny.SeqOf((_1309_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1309_v0), 0))).Int()).(bool), (_this).F30())
			var _1371_v47 _dafny.Set
			_ = _1371_v47
			_1371_v47 = _dafny.SetOf(_1368_v44, (func() _dafny.Sequence {
				if (_1369_v45).Contains((_1326_v14).Fm4(_1311_v1, _1370_v46, globalState)) {
					return (_1369_v45).Get((_1326_v14).Fm4(_1311_v1, _1370_v46, globalState)).(_dafny.Sequence)
				}
				return _1368_v44
			})())
			var _1372_v48 _dafny.Map
			_ = _1372_v48
			_1372_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1368_v44, (_this).F24())
			_1371_v47 = (func() _dafny.Set {
				var _coll80 = _dafny.NewBuilder()
				_ = _coll80
				for _iter85 := _dafny.Iterate((_1372_v48).Keys().Elements()); ; {
					_compr_80, _ok85 := _iter85()
					if !_ok85 {
						break
					}
					var _1373_v49 _dafny.Sequence
					_1373_v49 = interface{}(_compr_80).(_dafny.Sequence)
					if (_1372_v48).Contains(_1373_v49) {
						_coll80.Add(_1373_v49)
					}
				}
				return _coll80.ToSet()
			}()).Difference(_1371_v47)
			_1368_v44 = _1368_v44
			var _1374_v50 _dafny.Set
			_ = _1374_v50
			_1374_v50 = _dafny.SetOf((_1370_v46).Select((Companion_Default___.SafeIndex(_1311_v1, _dafny.IntOfUint32((_1370_v46).Cardinality()))).Uint32()).(bool), (_this).F30())
			var _1375_v51 _dafny.Map
			_ = _1375_v51
			_1375_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _1374_v50)
			_1375_v51 = (_1375_v51).Update(false, _1374_v50)
			(globalState).F7 = (_1311_v1).Plus((_dafny.Zero).Minus(_1311_v1))
		}
		var _1376_v52 _dafny.MultiSet
		_ = _1376_v52
		_1376_v52 = _dafny.MultiSetOf(false)
		var _1377_v53 _dafny.Sequence
		_ = _1377_v53
		_1377_v53 = _dafny.SeqOf(_1376_v52)
		var _1378_v54 _dafny.Sequence
		_ = _1378_v54
		_1378_v54 = _dafny.SeqOf(_1311_v1, _1311_v1)
		r0 = (_dafny.SetOf((_1377_v53).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1378_v54).Cardinality()), _dafny.IntOfUint32((_1377_v53).Cardinality()))).Uint32()).(_dafny.MultiSet), _dafny.MultiSetOf((_1326_v14).F24()))).IsDisjointFrom(func() _dafny.Set {
			var _coll81 = _dafny.NewBuilder()
			_ = _coll81
			for _iter86 := _dafny.Iterate((_dafny.MultiSetFromSeq(_1377_v53)).Elements()); ; {
				_compr_81, _ok86 := _iter86()
				if !_ok86 {
					break
				}
				var _1379_v55 _dafny.MultiSet
				_1379_v55 = interface{}(_compr_81).(_dafny.MultiSet)
				if (_dafny.MultiSetFromSeq(_1377_v53)).Contains(_1379_v55) {
					_coll81.Add(_1379_v55)
				}
			}
			return _coll81.ToSet()
		}())
		var _1380_v57 _dafny.Set
		_ = _1380_v57
		_1380_v57 = _dafny.SetOf(Companion_D5_.Create_DC12_(func() _dafny.Set {
			var _coll82 = _dafny.NewBuilder()
			_ = _coll82
			for _iter87 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-668), _dafny.IntOfInt64(-506))); ; {
				_compr_82, _ok87 := _iter87()
				if !_ok87 {
					break
				}
				var _1381_v56 _dafny.Int
				_1381_v56 = interface{}(_compr_82).(_dafny.Int)
				if ((_dafny.IntOfInt64(-668)).Cmp(_1381_v56) <= 0) && ((_1381_v56).Cmp(_dafny.IntOfInt64(-506)) < 0) {
					_coll82.Add((_1381_v56).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg94 _dafny.Int) interface{} {
							return coer94(arg94)
						}
					}(func(_1382_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					}))).Cardinality())))
				}
			}
			return _coll82.ToSet()
		}()), _1328_v16)
		var _1383_v58 _dafny.Set
		_ = _1383_v58
		_1383_v58 = _dafny.SetOf(_1380_v57)
		r1 = (_1383_v58).Cardinality()
		r2 = p0
		return r0, r1, r2
	}
}
func (_this *C7) M4(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(globalState).F13 = (_this).F24()
		var _1384_v0 _dafny.MultiSet
		_ = _1384_v0
		_1384_v0 = _dafny.MultiSetOf((_this).F30(), (_this).F24())
		var _1385_v1 D0
		_ = _1385_v1
		_1385_v1 = Companion_D0_.Create_DC0_(_1384_v0)
		var _1386_v2 *C2
		_ = _1386_v2
		var _nw209 *C2 = New_C2_()
		_ = _nw209
		_nw209.Ctor__((func() _dafny.MultiSet {
			if (_this).F31() {
				return _1384_v0
			}
			return (_1385_v1).Dtor_cf0()
		})(), (_this).F25(), _this.F23(), (_this).F24())
		_1386_v2 = _nw209
		var _1387_i0 _dafny.Int
		_ = _1387_i0
		_1387_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_1387_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1387_i0 = (_1387_i0).Plus(_dafny.One)
					var _1388_v3 _dafny.Array
					_ = _1388_v3
					var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
					_ = _nw210
					_1388_v3 = _nw210
					var _1389_v4 _dafny.Sequence
					_ = _1389_v4
					_1389_v4 = _dafny.UnicodeSeqOfUtf8Bytes("gtxxn")
					var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_1388_v3), 0))
					_ = _index194
					(_1388_v3).ArraySet1(_1389_v4, (_index194).Int())
					var _1390_v5 _dafny.Int
					_ = _1390_v5
					_1390_v5 = _dafny.IntOfInt64(924)
					var _1391_v6 _dafny.Sequence
					_ = _1391_v6
					_1391_v6 = _dafny.SeqOf((_this).F24())
					var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_1388_v3), 0))
					_ = _index195
					(_1388_v3).ArraySet1(_dafny.Companion_Sequence_.Update(_1389_v4, (Companion_Default___.SafeIndex((Companion_Default___.Fm0((_this).F27(), _1390_v5, globalState)).Times(_dafny.IntOfUint32((_1391_v6).Cardinality())), _dafny.IntOfUint32((_1389_v4).Cardinality()))).Uint32(), _this.F23()), (_index195).Int())
					var _1392_v7 _dafny.Sequence
					_ = _1392_v7
					_1392_v7 = _dafny.SeqOf(Companion_Default___.Fm0(true, _1390_v5, globalState))
					var _1393_v8 _dafny.Sequence
					_ = _1393_v8
					_1393_v8 = _dafny.SeqOf((_1392_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1392_v7).Cardinality()), _dafny.IntOfUint32((_1392_v7).Cardinality()))).Uint32()).(_dafny.Int))
					_1393_v8 = (func() _dafny.Sequence {
						if p0 {
							return _1393_v8
						}
						return _1393_v8
					})()
					var _1394_v9 D5
					_ = _1394_v9
					_1394_v9 = Companion_D5_.Create_DC13_(_1390_v5, true, (_this).F30())
					var _rhs212 _dafny.Int = _1390_v5
					_ = _rhs212
					var _rhs213 D5 = Companion_D5_.Create_DC13_(_1390_v5, (_this).F24(), !_dafny.Companion_Sequence_.Contains(_1391_v6, (_this).F31()))
					_ = _rhs213
					var _lhs164 *GlobalState = globalState
					_ = _lhs164
					_lhs164.F7 = _rhs212
					_1394_v9 = _rhs213
					var _1395_v10 _dafny.Array
					_ = _1395_v10
					var _nw211 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
					_ = _nw211
					_1395_v10 = _nw211
					var _1396_v11 _dafny.Map
					_ = _1396_v11
					_1396_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1395_v10, _this.F23())
					var _1397_v12 *C3
					_ = _1397_v12
					var _nw212 *C3 = New_C3_()
					_ = _nw212
					_nw212.Ctor__((_1396_v11).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1395_v10, _this.F23())), (_this).F25(), _this.F23(), true)
					_1397_v12 = _nw212
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		if (_this).F31() {
			(globalState).F13 = !((_this).F31())
			var _1398_v13 _dafny.Array
			_ = _1398_v13
			var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw213
			_1398_v13 = _nw213
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))
			_ = _index196
			(_1398_v13).ArraySet1(_dafny.IntOfInt64(-212), (_index196).Int())
			var _1399_v14 _dafny.Int
			_ = _1399_v14
			_1399_v14 = _dafny.IntOfInt64(17)
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))
			_ = _index197
			(_1398_v13).ArraySet1(_1399_v14, (_index197).Int())
			var _1400_v15 *C1
			_ = _1400_v15
			var _nw214 *C1 = New_C1_()
			_ = _nw214
			_nw214.Ctor__((_this).F27(), (_this).F30())
			_1400_v15 = _nw214
			var _1401_v16 *C1
			_ = _1401_v16
			_1401_v16 = _1400_v15
			var _1402_v17 _dafny.Sequence
			_ = _1402_v17
			_1402_v17 = _dafny.SeqOf(_1401_v16)
			var _1403_v18 _dafny.Sequence
			_ = _1403_v18
			_1403_v18 = (func() _dafny.Sequence {
				if (_this).F27() {
					return _dafny.SeqOf(_1401_v16)
				}
				return _1402_v17
			})()
			var _source22 _dafny.Sequence = _1403_v18
			_ = _source22
			var _1404___mcc_h0 _dafny.Sequence = _source22
			_ = _1404___mcc_h0
			var _1405_cf34 _dafny.Sequence = _1404___mcc_h0
			_ = _1405_cf34
			(globalState).F13 = (_this).F31()
			var _1406_v19 _dafny.Sequence
			_ = _1406_v19
			_1406_v19 = _dafny.UnicodeSeqOfUtf8Bytes("eqrnaakjo")
			var _1407_v20 _dafny.Set
			_ = _1407_v20
			_1407_v20 = _dafny.SetOf(_dafny.IntOfInt64(883), _1399_v14)
			var _1408_v21 _dafny.Sequence
			_ = _1408_v21
			_1408_v21 = _dafny.SeqOf(_1399_v14, _dafny.IntOfUint32((_1406_v19).Cardinality()), (_1407_v20).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_1399_v14)))
			var _1409_v22 _dafny.Sequence
			_ = _1409_v22
			_1409_v22 = _dafny.SeqOf(_1400_v15)
			var _1410_v23 _dafny.MultiSet
			_ = _1410_v23
			_1410_v23 = _dafny.MultiSetOf(_1400_v15, _1400_v15, _1400_v15)
			var _1411_v24 _dafny.Sequence
			_ = _1411_v24
			_1411_v24 = _dafny.SeqOf(_dafny.MultiSetOf(_1400_v15, _1400_v15), _dafny.MultiSetFromSeq(_1409_v22), _1410_v23, _1410_v23)
			var _1412_v25 _dafny.Map
			_ = _1412_v25
			_1412_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1399_v14, (_1398_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))).Int()).(_dafny.Int))
			var _rhs214 _dafny.Int = (_this).Fm9(_1408_v21, (_1398_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))).Int()).(_dafny.Int), (_1400_v15).F39(), globalState)
			_ = _rhs214
			var _rhs215 _dafny.Int = (_dafny.IntOfUint32((_1411_v24).Cardinality())).Minus(_1399_v14)
			_ = _rhs215
			var _rhs216 bool = ((_this).F31()) || (!(_1412_v25).Contains((_dafny.Zero).Minus((_1398_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))).Int()).(_dafny.Int))))
			_ = _rhs216
			var _lhs165 *GlobalState = globalState
			_ = _lhs165
			_1399_v14 = _rhs214
			_1399_v14 = _rhs215
			_lhs165.F13 = _rhs216
			var _1413_v26 _dafny.Map
			_ = _1413_v26
			_1413_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), (_1398_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))).Int()).(_dafny.Int))
			(globalState).F7 = ((_1413_v26).Cardinality()).Times(Companion_Default___.SafeDivisionInt(_1399_v14, _1399_v14))
			(globalState).F7 = (_1398_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1398_v13), 0))).Int()).(_dafny.Int)
			(globalState).F7 = _1399_v14
			var _1414_v27 _dafny.Map
			_ = _1414_v27
			_1414_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(_1399_v14)).Times(_1399_v14), p0)
			_1414_v27 = (_1414_v27).Update(_1399_v14, (_1400_v15).F39())
		} else {
			var _1415_v28 _dafny.Int
			_ = _1415_v28
			_1415_v28 = _dafny.IntOfInt64(782)
			var _1416_v29 _dafny.Sequence
			_ = _1416_v29
			_1416_v29 = _dafny.UnicodeSeqOfUtf8Bytes("byo")
			if (_1415_v28).Cmp(_dafny.IntOfUint32((_1416_v29).Cardinality())) > 0 {
				(globalState).F13 = !(_dafny.Companion_Sequence_.IsPrefixOf(_1416_v29, _1416_v29))
				var _1417_v30 _dafny.Sequence
				_ = _1417_v30
				_1417_v30 = _dafny.SeqOf(Companion_Default___.Fm2(globalState), (func() bool {
					if (_this).F30() {
						return p0
					}
					return p0
				})(), (_this).F24())
				var _1418_v31 _dafny.Map
				_ = _1418_v31
				_1418_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1415_v28, _1415_v28)
				var _rhs217 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.Int {
					if (_this).F30() {
						return _dafny.IntOfInt64(386)
					}
					return _1415_v28
				})()).Times(Companion_Default___.Fm0(!((_this).F27()), (func() _dafny.Int {
					if (_1384_v0).Contains(true) {
						return (_1384_v0).Multiplicity(true)
					}
					return _1415_v28
				})(), globalState)))
				_ = _rhs217
				var _rhs218 _dafny.Sequence = _1417_v30
				_ = _rhs218
				var _rhs219 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_1418_v31).Contains(_1415_v28) {
						return (_1418_v31).Get(_1415_v28).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(_1415_v28)
				})())
				_ = _rhs219
				var _lhs166 *GlobalState = globalState
				_ = _lhs166
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				_lhs166.F7 = _rhs217
				_1417_v30 = _rhs218
				_lhs167.F7 = _rhs219
				var _1419_v32 _dafny.Set
				_ = _1419_v32
				_1419_v32 = _dafny.SetOf((_this).F24())
				var _1420_v33 _dafny.Map
				_ = _1420_v33
				_1420_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1419_v32).Cardinality())
				var _1421_v34 _dafny.Set
				_ = _1421_v34
				_1421_v34 = _dafny.SetOf(_1420_v33, _1420_v33)
				var _1422_v35 D0
				_ = _1422_v35
				_1422_v35 = Companion_D0_.Create_DC3_(Companion_Default___.Fm28(_1416_v29, _dafny.SeqOf(_1415_v28), Companion_Default___.Fm2(globalState), _1421_v34, globalState))
				var _1423_v36 _dafny.Map
				_ = _1423_v36
				_1423_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1419_v32).IsDisjointFrom(_1419_v32), _1422_v35)
				_1423_v36 = (_1423_v36).Update((_this).F31(), _1422_v35)
				var _1424_v37 _dafny.Array
				_ = _1424_v37
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_38
				var _nw215 _dafny.Array
				_ = _nw215
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw215 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Int = (func(_1425_v29 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1426_i1 _dafny.Int) _dafny.Int {
							return (_1426_i1).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1425_v29).Cardinality())))
						}
					})(_1416_v29)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw215 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw215).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw215).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1424_v37 = _nw215
				var _1427_v38 _dafny.Set
				_ = _1427_v38
				_1427_v38 = _dafny.SetOf(_1415_v28, _1415_v28, _1415_v28, _1415_v28)
				var _1428_v39 _dafny.MultiSet
				_ = _1428_v39
				_1428_v39 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1417_v30).Cardinality()))
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1424_v37), 0))
				_ = _index198
				(_1424_v37).ArraySet1(((_1427_v38).Union(func() _dafny.Set {
					var _coll83 = _dafny.NewBuilder()
					_ = _coll83
					for _iter88 := _dafny.Iterate((_1428_v39).Elements()); ; {
						_compr_83, _ok88 := _iter88()
						if !_ok88 {
							break
						}
						var _1429_v40 _dafny.Int
						_1429_v40 = interface{}(_compr_83).(_dafny.Int)
						if (_1428_v39).Contains(_1429_v40) {
							_coll83.Add(Companion_Default___.SafeModuloInt(_1429_v40, _dafny.IntOfInt64(321)))
						}
					}
					return _coll83.ToSet()
				}())).Cardinality(), (_index198).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1424_v37), 0))
				_ = _index199
				(_1424_v37).ArraySet1(_1415_v28, (_index199).Int())
				var _1430_v41 *C1
				_ = _1430_v41
				var _nw216 *C1 = New_C1_()
				_ = _nw216
				_nw216.Ctor__(true, (_this).F30())
				_1430_v41 = _nw216
				var _1431_v42 *C1
				_ = _1431_v42
				_1431_v42 = _1430_v41
				_1431_v42 = _1431_v42
			} else {
				var _1432_v43 *C2
				_ = _1432_v43
				var _nw217 *C2 = New_C2_()
				_ = _nw217
				_nw217.Ctor__(_1384_v0, (_this).F25(), _this.F23(), false)
				_1432_v43 = _nw217
				var _1433_v44 _dafny.MultiSet
				_ = _1433_v44
				_1433_v44 = _dafny.MultiSetOf(_1415_v28)
				var _1434_v45 _dafny.Map
				_ = _1434_v45
				_1434_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1416_v29).Select((Companion_Default___.SafeIndex(_1415_v28, _dafny.IntOfUint32((_1416_v29).Cardinality()))).Uint32()).(_dafny.CodePoint), (_1433_v44).IsProperSubsetOf(_1433_v44))
				var _1435_v46 _dafny.Array
				_ = _1435_v46
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_39
				var _nw218 _dafny.Array
				_ = _nw218
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw218 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) bool = func(_1436_i2 _dafny.Int) bool {
						return (_this).F27()
					}
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw218 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw218).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw218).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1435_v46 = _nw218
				var _1437_v47 _dafny.Map
				_ = _1437_v47
				_1437_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if (_this).F24() {
						return _1435_v46
					}
					return _1435_v46
				})(), _1434_v45)
				_1434_v45 = (func() _dafny.Map {
					if (_1437_v47).Contains(_1435_v46) {
						return (_1437_v47).Get(_1435_v46).(_dafny.Map)
					}
					return ((_1434_v45).Update(_this.F23(), p0)).Merge(_1434_v45)
				})()
				(globalState).F13 = (_this).F30()
				_1415_v28 = _1415_v28
				var _1438_v48 _dafny.Sequence
				_ = _1438_v48
				_1438_v48 = _dafny.SeqOf((_this).F31(), (_this).F24(), (_this).F31(), (_this).F30(), (_this).F31())
				var _1439_v49 _dafny.Map
				_ = _1439_v49
				_1439_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1415_v28, _1438_v48)
				var _1440_v50 D6
				_ = _1440_v50
				_1440_v50 = Companion_D6_.Create_DC15_(_1439_v49)
				var _1441_v51 _dafny.Sequence
				_ = _1441_v51
				_1441_v51 = _dafny.SeqOf(_1440_v50)
				var _1442_v52 _dafny.Map
				_ = _1442_v52
				_1442_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_1441_v51, _1441_v51), _1415_v28)
				(globalState).F7 = (_1442_v52).Cardinality()
			}
			var _1443_v53 _dafny.Array
			_ = _1443_v53
			var _nw219 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw219
			_1443_v53 = _nw219
			var _1444_v54 _dafny.Map
			_ = _1444_v54
			_1444_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v53, Companion_Default___.Fm34(_dafny.IntOfInt64(123), (_this).F24(), globalState))
			var _1445_v55 *C3
			_ = _1445_v55
			var _nw220 *C3 = New_C3_()
			_ = _nw220
			_nw220.Ctor__(_1444_v54, (_this).F25(), _this.F23(), (_this).F31())
			_1445_v55 = _nw220
			var _1446_v56 _dafny.MultiSet
			_ = _1446_v56
			_1446_v56 = _dafny.MultiSetOf(_this.F23())
			var _1447_v57 _dafny.Map
			_ = _1447_v57
			_1447_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1446_v56).Cardinality(), _1445_v55)
			var _1448_v58 _dafny.Sequence
			_ = _1448_v58
			_1448_v58 = _dafny.SeqOf(_1445_v55, _1445_v55)
			var _1449_v59 _dafny.Array
			_ = _1449_v59
			var _nwElement0_44 *C3 = _1445_v55
			_ = _nwElement0_44
			var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(27))
			_ = _nw221
			(_nw221).ArraySet1(_nwElement0_44, 0)
			(_nw221).ArraySet1(_1445_v55, 1)
			(_nw221).ArraySet1(_1445_v55, 2)
			(_nw221).ArraySet1(_1445_v55, 3)
			(_nw221).ArraySet1((func() *C3 {
				if (_1447_v57).Contains(_1415_v28) {
					return (_1447_v57).Get(_1415_v28).(*C3)
				}
				return _1445_v55
			})(), 4)
			(_nw221).ArraySet1(_1445_v55, 5)
			(_nw221).ArraySet1(_1445_v55, 6)
			(_nw221).ArraySet1(_1445_v55, 7)
			(_nw221).ArraySet1(_1445_v55, 8)
			(_nw221).ArraySet1(_1445_v55, 9)
			(_nw221).ArraySet1(_1445_v55, 10)
			(_nw221).ArraySet1(_1445_v55, 11)
			(_nw221).ArraySet1(_1445_v55, 12)
			(_nw221).ArraySet1(_1445_v55, 13)
			(_nw221).ArraySet1((_1448_v58).Select((Companion_Default___.SafeIndex(_1415_v28, _dafny.IntOfUint32((_1448_v58).Cardinality()))).Uint32()).(*C3), 14)
			(_nw221).ArraySet1(_1445_v55, 15)
			(_nw221).ArraySet1(_1445_v55, 16)
			(_nw221).ArraySet1(_1445_v55, 17)
			(_nw221).ArraySet1(_1445_v55, 18)
			(_nw221).ArraySet1(_1445_v55, 19)
			(_nw221).ArraySet1(_1445_v55, 20)
			(_nw221).ArraySet1(_1445_v55, 21)
			(_nw221).ArraySet1(_1445_v55, 22)
			(_nw221).ArraySet1(_1445_v55, 23)
			(_nw221).ArraySet1((_1448_v58).Select((Companion_Default___.SafeIndex(_1415_v28, _dafny.IntOfUint32((_1448_v58).Cardinality()))).Uint32()).(*C3), 24)
			(_nw221).ArraySet1(_1445_v55, 25)
			(_nw221).ArraySet1(_1445_v55, 26)
			_1449_v59 = _nw221
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1449_v59), 0))
			_ = _index200
			(_1449_v59).ArraySet1(_1445_v55, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1449_v59), 0))
			_ = _index201
			(_1449_v59).ArraySet1(_1445_v55, (_index201).Int())
			var _1450_v60 D2
			_ = _1450_v60
			_1450_v60 = Companion_D2_.Create_DC7_(p0, p0)
			var _1451_v61 _dafny.Set
			_ = _1451_v61
			_1451_v61 = _dafny.SetOf(_1450_v60, _1450_v60, _1450_v60, _1450_v60)
			var _pat_let_tv29 = p0
			_ = _pat_let_tv29
			_1451_v61 = _dafny.SetOf(func(_pat_let45_0 D2) D2 {
				return func(_1452_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let46_0 bool) D2 {
						return func(_1453_dt__update_hcf14_h0 bool) D2 {
							return Companion_D2_.Create_DC7_((_1452_dt__update__tmp_h0).Dtor_cf13(), _1453_dt__update_hcf14_h0)
						}(_pat_let46_0)
					}(_pat_let_tv29)
				}(_pat_let45_0)
			}(_1450_v60))
			var _1454_v63 _dafny.Set
			_ = _1454_v63
			_1454_v63 = _dafny.SetOf(_1416_v29)
			if !((func() _dafny.Map {
				var _coll84 = _dafny.NewMapBuilder()
				_ = _coll84
				for _iter89 := _dafny.Iterate((_1454_v63).Elements()); ; {
					_compr_84, _ok89 := _iter89()
					if !_ok89 {
						break
					}
					var _1455_v62 _dafny.Sequence
					_1455_v62 = interface{}(_compr_84).(_dafny.Sequence)
					if (_1454_v63).Contains(_1455_v62) {
						_coll84.Add(_1455_v62, (Companion_D9_.Create_DC20_(_1416_v29, p0, _1415_v28)).Dtor_cf31())
					}
				}
				return _coll84.ToMap()
			}()).Update(_1416_v29, _1416_v29)).Contains((Companion_D4_.Create_DC10_(Companion_Default___.Fm1(p0, globalState))).Dtor_cf17()) {
				var _1456_v64 *C3
				_ = _1456_v64
				var _nw222 *C3 = New_C3_()
				_ = _nw222
				_nw222.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v53, _this.F23()), (func() _dafny.Array {
					if (_this).F30() {
						return (_this).F25()
					}
					return (_this).F25()
				})(), _dafny.CodePoint('e'), (_this).F31())
				_1456_v64 = _nw222
				var _1457_v65 _dafny.Map
				_ = _1457_v65
				_1457_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v53, Companion_Default___.Fm0((_this).F30(), (_1386_v2).Fm24((_dafny.Zero).Minus(_1415_v28), true, globalState), globalState))
				var _rhs220 _dafny.CodePoint = (func() _dafny.CodePoint {
					if p0 {
						return _this.F23()
					}
					return _this.F23()
				})()
				_ = _rhs220
				var _rhs221 _dafny.Int = (func() _dafny.Int {
					if (_1415_v28).Cmp(_1415_v28) < 0 {
						return _1415_v28
					}
					return (_1457_v65).Cardinality()
				})()
				_ = _rhs221
				var _lhs168 *C7 = _this
				_ = _lhs168
				var _lhs169 *GlobalState = globalState
				_ = _lhs169
				_lhs168.F23_set_(_rhs220)
				_lhs169.F7 = _rhs221
				_1415_v28 = _dafny.IntOfUint32((_1416_v29).Cardinality())
				var _1458_v66 D12
				_ = _1458_v66
				_1458_v66 = Companion_D12_.Create_DC25_((_this).F31(), _this.F23())
				_1458_v66 = _1458_v66
				var _1459_v67 bool
				_ = _1459_v67
				var _1460_v68 _dafny.Int
				_ = _1460_v68
				var _1461_v69 bool
				_ = _1461_v69
				var _out40 bool
				_ = _out40
				var _out41 _dafny.Int
				_ = _out41
				var _out42 bool
				_ = _out42
				_out40, _out41, _out42 = (_1445_v55).M3((_this).F27(), _this.F23(), globalState)
				_1459_v67 = _out40
				_1460_v68 = _out41
				_1461_v69 = _out42
			} else {
				(globalState).F7 = _1415_v28
				var _1462_v70 _dafny.Map
				_ = _1462_v70
				_1462_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.CodePoint('q'))
				var _1463_v71 D13
				_ = _1463_v71
				_1463_v71 = Companion_D13_.Create_DC27_(_1462_v70)
				var _1464_v72 D5
				_ = _1464_v72
				_1464_v72 = Companion_D5_.Create_DC14_((_this).F27())
				var _1465_v73 D11
				_ = _1465_v73
				_1465_v73 = Companion_D11_.Create_DC23_(p0, (_dafny.Zero).Minus(_1415_v28), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p0), _1464_v72)
				var _1466_v74 _dafny.Sequence
				_ = _1466_v74
				_1466_v74 = _dafny.SeqOf((_1465_v73).Dtor_cf36())
				var _1467_v75 _dafny.Array
				_ = _1467_v75
				var _nwElement0_45 _dafny.Map = (_1462_v70).Update((_this).F24(), _this.F23())
				_ = _nwElement0_45
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(8))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_45, 0)
				(_nw223).ArraySet1((_1463_v71).Dtor_cf44(), 1)
				(_nw223).ArraySet1((_1462_v70).Update((_this).F27(), _this.F23()), 2)
				(_nw223).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1466_v74).Select((Companion_Default___.SafeIndex(_1415_v28, _dafny.IntOfUint32((_1466_v74).Cardinality()))).Uint32()).(bool), _this.F23()), 3)
				(_nw223).ArraySet1(_1462_v70, 4)
				(_nw223).ArraySet1(_1462_v70, 5)
				(_nw223).ArraySet1(_1462_v70, 6)
				(_nw223).ArraySet1(_1462_v70, 7)
				_1467_v75 = _nw223
				var _1468_v76 _dafny.Array
				_ = _1468_v76
				var _nw224 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw224
				_1468_v76 = _nw224
				var _1469_v77 _dafny.Sequence
				_ = _1469_v77
				_1469_v77 = _dafny.SeqOf(_1467_v75, _1467_v75, _1467_v75, _1467_v75, _1467_v75)
				var _1470_v78 _dafny.Map
				_ = _1470_v78
				_1470_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1468_v76, (_1469_v77).Select((Companion_Default___.SafeIndex(_1415_v28, _dafny.IntOfUint32((_1469_v77).Cardinality()))).Uint32()).(_dafny.Array))
				var _1471_v79 D14
				_ = _1471_v79
				_1471_v79 = Companion_D14_.Create_DC30_(_1467_v75)
				var _1472_v80 _dafny.Map
				_ = _1472_v80
				_1472_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1467_v75)
				var _1473_v81 _dafny.Array
				_ = _1473_v81
				var _nwElement0_46 _dafny.Array = _1467_v75
				_ = _nwElement0_46
				var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(25))
				_ = _nw225
				(_nw225).ArraySet1(_nwElement0_46, 0)
				(_nw225).ArraySet1(_1467_v75, 1)
				(_nw225).ArraySet1(_1467_v75, 2)
				(_nw225).ArraySet1((func() _dafny.Array {
					if (_1470_v78).Contains(_1468_v76) {
						return (_1470_v78).Get(_1468_v76).(_dafny.Array)
					}
					return _1467_v75
				})(), 3)
				(_nw225).ArraySet1(_1467_v75, 4)
				(_nw225).ArraySet1(_1467_v75, 5)
				(_nw225).ArraySet1(_1467_v75, 6)
				(_nw225).ArraySet1(_1467_v75, 7)
				(_nw225).ArraySet1(_1467_v75, 8)
				(_nw225).ArraySet1(_1467_v75, 9)
				(_nw225).ArraySet1(_1467_v75, 10)
				(_nw225).ArraySet1(_1467_v75, 11)
				(_nw225).ArraySet1(_1467_v75, 12)
				(_nw225).ArraySet1(_1467_v75, 13)
				(_nw225).ArraySet1(_1467_v75, 14)
				(_nw225).ArraySet1(_1467_v75, 15)
				(_nw225).ArraySet1((_1471_v79).Dtor_cf51(), 16)
				(_nw225).ArraySet1(_1467_v75, 17)
				(_nw225).ArraySet1(_1467_v75, 18)
				(_nw225).ArraySet1((func() _dafny.Array {
					if (_1472_v80).Contains((_this).F24()) {
						return (_1472_v80).Get((_this).F24()).(_dafny.Array)
					}
					return _1467_v75
				})(), 19)
				(_nw225).ArraySet1(_1467_v75, 20)
				(_nw225).ArraySet1(_1467_v75, 21)
				(_nw225).ArraySet1(_1467_v75, 22)
				(_nw225).ArraySet1(_1467_v75, 23)
				(_nw225).ArraySet1(_1467_v75, 24)
				_1473_v81 = _nw225
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_1473_v81), 0))
				_ = _index202
				(_1473_v81).ArraySet1(_1467_v75, (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_1473_v81), 0))
				_ = _index203
				(_1473_v81).ArraySet1(_1467_v75, (_index203).Int())
				(_this).F23_set_(_this.F23())
				var _1474_v82 _dafny.Sequence
				_ = _1474_v82
				_1474_v82 = _dafny.SeqOf(_1415_v28)
				var _1475_v83 _dafny.MultiSet
				_ = _1475_v83
				_1475_v83 = _dafny.MultiSetOf(_1415_v28, _1415_v28, _1415_v28)
				(globalState).F13 = (_1475_v83).IsProperSubsetOf(_dafny.MultiSetFromSeq(_1474_v82))
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_1443_v53), 0))
				_ = _index204
				(_1443_v53).ArraySet1(((_this).F24()) && ((_this).F27()), (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_1443_v53), 0))
				_ = _index205
				(_1443_v53).ArraySet1(p0, (_index205).Int())
			}
			var _1476_v84 *C5
			_ = _1476_v84
			var _nw226 *C5 = New_C5_()
			_ = _nw226
			_nw226.Ctor__(_1443_v53)
			_1476_v84 = _nw226
		}
		var _1477_v85 _dafny.Array
		_ = _1477_v85
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_40
		var _nw227 _dafny.Array
		_ = _nw227
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw227 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) bool = (func(_1478_p0 bool) func(_dafny.Int) bool {
				return func(_1479_i3 _dafny.Int) bool {
					return _1478_p0
				}
			})(p0)
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw227 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw227).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw227).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1477_v85 = _nw227
		var _1480_v86 _dafny.Map
		_ = _1480_v86
		_1480_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1477_v85, _dafny.IntOfInt64(550))
		var _1481_v87 _dafny.Int
		_ = _1481_v87
		_1481_v87 = _dafny.IntOfInt64(629)
		var _1482_v88 _dafny.MultiSet
		_ = _1482_v88
		_1482_v88 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1481_v87), _1481_v87, _1481_v87, _1481_v87)
		var _1483_v89 _dafny.Map
		_ = _1483_v89
		_1483_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F31()), (_1386_v2).Fm4(_1481_v87, _dafny.SeqOf((_this).F24(), (_this).F31(), (_this).F30()), globalState))
		var _rhs222 _dafny.Map = _1480_v86
		_ = _rhs222
		var _rhs223 _dafny.MultiSet = ((_1482_v88).Difference(Companion_Default___.Fm18(globalState))).Difference(_1482_v88)
		_ = _rhs223
		var _rhs224 _dafny.Int = (_1481_v87).Plus((_1483_v89).Cardinality())
		_ = _rhs224
		var _lhs170 *GlobalState = globalState
		_ = _lhs170
		_1480_v86 = _rhs222
		_1482_v88 = _rhs223
		_lhs170.F7 = _rhs224
		var _1484_v90 _dafny.Array
		_ = _1484_v90
		var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw228
		_1484_v90 = _nw228
		var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1484_v90), 0))
		_ = _index206
		(_1484_v90).ArraySet1(_1481_v87, (_index206).Int())
		var _1485_v91 _dafny.Sequence
		_ = _1485_v91
		_1485_v91 = _dafny.UnicodeSeqOfUtf8Bytes("oaniuldr")
		var _1486_v92 _dafny.Map
		_ = _1486_v92
		_1486_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.Companion_Sequence_.Update(_1485_v91, (Companion_Default___.SafeIndex(_1481_v87, _dafny.IntOfUint32((_1485_v91).Cardinality()))).Uint32(), _this.F23()))
		var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1484_v90), 0))
		_ = _index207
		var _rhs225 bool = (_1486_v92).Equals((_1486_v92).Merge(_1486_v92))
		_ = _rhs225
		var _rhs226 _dafny.Int = _1481_v87
		_ = _rhs226
		var _lhs171 *GlobalState = globalState
		_ = _lhs171
		var _lhs172 _dafny.Array = _1484_v90
		_ = _lhs172
		var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1484_v90), 0))
		_ = _lhs173
		_lhs171.F13 = _rhs225
		(_lhs172).ArraySet1(_rhs226, (_lhs173).Int())
		r0 = ((_this).F27()) && ((_this).F24())
		return r0
	}
}
func (_this *C7) M7(p0 _dafny.Int, p1 bool, globalState *GlobalState) {
	{
		var _1487_v0 _dafny.Int
		_ = _1487_v0
		var _1488_v1 bool
		_ = _1488_v1
		var _1489_v2 _dafny.Int
		_ = _1489_v2
		var _out43 _dafny.Int
		_ = _out43
		var _out44 bool
		_ = _out44
		var _out45 _dafny.Int
		_ = _out45
		_out43, _out44, _out45 = Companion_Default___.M0((p0).Plus(p0), Companion_Default___.Fm3((_this).F30(), globalState), true, (p0).Plus(p0), globalState)
		_1487_v0 = _out43
		_1488_v1 = _out44
		_1489_v2 = _out45
		(globalState).F7 = (p0).Minus(p0)
		var _1490_v3 _dafny.Map
		_ = _1490_v3
		_1490_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p0)
		var _1491_v4 _dafny.Sequence
		_ = _1491_v4
		_1491_v4 = _dafny.UnicodeSeqOfUtf8Bytes("dfrewfj")
		var _1492_v5 _dafny.Set
		_ = _1492_v5
		_1492_v5 = _dafny.SetOf(_1489_v2)
		var _1493_v6 D5
		_ = _1493_v6
		_1493_v6 = Companion_D5_.Create_DC12_(_1492_v5)
		var _pat_let_tv30 = p1
		_ = _pat_let_tv30
		var _pat_let_tv31 = _1487_v0
		_ = _pat_let_tv31
		var _rhs227 _dafny.Map = _1490_v3
		_ = _rhs227
		var _rhs228 _dafny.Int = p0
		_ = _rhs228
		var _rhs229 bool = (_this).F24()
		_ = _rhs229
		var _rhs230 bool = !(func(_source23 D5) bool {
			if _source23.Is_DC13() {
				var _1494___mcc_h0 _dafny.Int = _source23.Get_().(D5_DC13).Cf20
				_ = _1494___mcc_h0
				var _1495___mcc_h1 bool = _source23.Get_().(D5_DC13).Cf21
				_ = _1495___mcc_h1
				var _1496___mcc_h2 bool = _source23.Get_().(D5_DC13).Cf22
				_ = _1496___mcc_h2
				var _1497_cf22 bool = _1496___mcc_h2
				_ = _1497_cf22
				var _1498_cf21 bool = _1495___mcc_h1
				_ = _1498_cf21
				var _1499_cf20 _dafny.Int = _1494___mcc_h0
				_ = _1499_cf20
				return !(_pat_let_tv30) || (_1498_cf21)
			} else if _source23.Is_DC14() {
				var _1500___mcc_h3 bool = _source23.Get_().(D5_DC14).Cf23
				_ = _1500___mcc_h3
				var _1501_cf23 bool = _1500___mcc_h3
				_ = _1501_cf23
				return _1501_cf23
			} else {
				var _1502___mcc_h4 _dafny.Set = _source23.Get_().(D5_DC12).Cf19
				_ = _1502___mcc_h4
				var _1503_cf19 _dafny.Set = _1502___mcc_h4
				_ = _1503_cf19
				return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wpj")).Cardinality())).Cmp(_pat_let_tv31) < 0
			}
		}(_1493_v6))
		_ = _rhs230
		var _rhs231 _dafny.Sequence = (func() _dafny.Sequence {
			if _1488_v1 {
				return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}(func(_1504_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), (Companion_Default___.SafeIndex(_1489_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}(func(_1505_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				}))).Cardinality()))).Uint32(), _this.F23())
			}
			return _dafny.Companion_Sequence_.Concatenate(_1491_v4, _1491_v4)
		})()
		_ = _rhs231
		_1490_v3 = _rhs227
		_1489_v2 = _rhs228
		_1488_v1 = _rhs229
		_1488_v1 = _rhs230
		_1491_v4 = _rhs231
		var _hi8 _dafny.Int = (p0).Plus(_1489_v2)
		_ = _hi8
		for _1506_i1 := p0; _1506_i1.Cmp(_hi8) < 0; _1506_i1 = _1506_i1.Plus(_dafny.One) {
			var _1507_v7 _dafny.Map
			_ = _1507_v7
			_1507_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if p1 {
					return _1487_v0
				}
				return p0
			})(), !((_this).F27()))
			var _1508_v8 _dafny.Sequence
			_ = _1508_v8
			_1508_v8 = _dafny.SeqOf(p1, (_this).F27())
			var _1509_v9 _dafny.Sequence
			_ = _1509_v9
			_1509_v9 = _dafny.SeqOf(false, (_this).Fm4(p0, _dafny.SeqOf((_this).F30()), globalState), (_1508_v8).Select((Companion_Default___.SafeIndex(_1489_v2, _dafny.IntOfUint32((_1508_v8).Cardinality()))).Uint32()).(bool), (_this).F30(), (_this).F27())
			_1507_v7 = (_1507_v7).Update(Companion_Default___.SafeDivisionInt(_1489_v2, _1489_v2), (_dafny.IntOfUint32((_1509_v9).Cardinality())).Cmp(_1487_v0) >= 0)
			_1488_v1 = p1
			(globalState).F13 = p1
			(globalState).F13 = (_this).Fm4(_1487_v0, _1508_v8, globalState)
		}
		var _1510_v10 _dafny.Map
		_ = _1510_v10
		_1510_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(p0, _1487_v0))
		var _1511_v11 D4
		_ = _1511_v11
		_1511_v11 = Companion_D4_.Create_DC10_(_1491_v4)
		var _1512_v12 _dafny.Map
		_ = _1512_v12
		_1512_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1510_v10, _1511_v11)
		var _1513_v14 T0
		_ = _1513_v14
		var _nw229 *C0 = New_C0_()
		_ = _nw229
		_nw229.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _1491_v4), (_this).F31(), _this.F23(), (_this).F31())
		_1513_v14 = _nw229
		var _1514_v15 _dafny.Map
		_ = _1514_v15
		_1514_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1488_v1, _1513_v14)
		var _1515_v16 _dafny.Map
		_ = _1515_v16
		_1515_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1489_v2, _1514_v15)
		var _1516_v17 _dafny.Sequence
		_ = _1516_v17
		_1516_v17 = _dafny.SeqOf(_1487_v0, (_1515_v16).Cardinality(), p0, _1487_v0)
		_1512_v12 = (_1512_v12).Update((func() _dafny.Map {
			if !(!((_this).F31())) {
				return func() _dafny.Map {
					var _coll85 = _dafny.NewMapBuilder()
					_ = _coll85
					for _iter90 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-590), _dafny.IntOfInt64(391))); ; {
						_compr_85, _ok90 := _iter90()
						if !_ok90 {
							break
						}
						var _1517_v13 _dafny.Int
						_1517_v13 = interface{}(_compr_85).(_dafny.Int)
						if ((_dafny.IntOfInt64(-590)).Cmp(_1517_v13) <= 0) && ((_1517_v13).Cmp(_dafny.IntOfInt64(391)) < 0) {
							_coll85.Add((_1517_v13).Plus(_dafny.IntOfInt64(202)), _dafny.SeqOf(p0))
						}
					}
					return _coll85.ToMap()
				}()
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1491_v4).Cardinality()), _1516_v17)
		})(), Companion_D4_.Create_DC10_(_1491_v4))
		var _hi9 _dafny.Int = Companion_Default___.SafeDivisionInt((Companion_D0_.Create_DC2_(!(_1488_v1), _1487_v0, _1488_v1)).Dtor_cf3(), _1489_v2)
		_ = _hi9
		for _1518_i2 := _dafny.IntOfInt64(743); _1518_i2.Cmp(_hi9) < 0; _1518_i2 = _1518_i2.Plus(_dafny.One) {
			if !((_1492_v5).IsProperSubsetOf(_1492_v5)) {
				var _rhs232 bool = (func() bool {
					if (_this).F27() {
						return (func() bool {
							if false {
								return (_1513_v14).F24()
							}
							return (_1513_v14).F24()
						})()
					}
					return (_this).F31()
				})()
				_ = _rhs232
				var _rhs233 _dafny.Sequence = _1491_v4
				_ = _rhs233
				var _lhs174 *GlobalState = globalState
				_ = _lhs174
				_lhs174.F13 = _rhs232
				_1491_v4 = _rhs233
				var _1519_v18 _dafny.Sequence
				_ = _1519_v18
				_1519_v18 = _dafny.SeqOf((_this).F24())
				var _1520_v19 D14
				_ = _1520_v19
				_1520_v19 = Companion_D14_.Create_DC31_((_1513_v14).F24())
				_1519_v18 = _dafny.Companion_Sequence_.Concatenate(_1519_v18, _dafny.SeqOf((_1520_v19).Dtor_cf52()))
				var _1521_v20 _dafny.Sequence
				_ = _1521_v20
				_1521_v20 = _dafny.SeqOf(Companion_Default___.Fm36(globalState))
				var _1522_v21 _dafny.MultiSet
				_ = _1522_v21
				_1522_v21 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1521_v20).Cardinality()), _1489_v2, _1518_i2)
				var _1523_v22 _dafny.Map
				_ = _1523_v22
				_1523_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1522_v21).IsDisjointFrom(_1522_v21), _1492_v5)
				_1489_v2 = (_1523_v22).Cardinality()
				var _1524_v23 _dafny.Map
				_ = _1524_v23
				_1524_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1518_i2).Cmp(_1489_v2) >= 0, !((_this).F27()) || (!(_1488_v1)))
				_1524_v23 = (_1524_v23).Update(true, ((_dafny.Zero).Minus(_1487_v0)).Cmp(p0) != 0)
				_1519_v18 = _1519_v18
			} else {
				var _1525_v24 _dafny.MultiSet
				_ = _1525_v24
				_1525_v24 = _dafny.MultiSetOf((_1513_v14).F24())
				var _1526_v25 _dafny.Map
				_ = _1526_v25
				_1526_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1518_i2, _1489_v2)
				var _1527_v26 _dafny.Array
				_ = _1527_v26
				var _nwElement0_47 _dafny.Int = _dafny.IntOfUint32((_1491_v4).Cardinality())
				_ = _nwElement0_47
				var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(17))
				_ = _nw230
				(_nw230).ArraySet1(_nwElement0_47, 0)
				(_nw230).ArraySet1(p0, 1)
				(_nw230).ArraySet1(((_dafny.Zero).Minus(_1518_i2)).Minus(_dafny.IntOfInt64(862)), 2)
				(_nw230).ArraySet1(_1489_v2, 3)
				(_nw230).ArraySet1((_dafny.IntOfInt64(788)).Times((_1525_v24).Cardinality()), 4)
				(_nw230).ArraySet1((func() _dafny.Int {
					if (_1525_v24).Contains((_this).F24()) {
						return (_1525_v24).Multiplicity((_this).F24())
					}
					return p0
				})(), 5)
				(_nw230).ArraySet1(_dafny.IntOfInt64(-929), 6)
				(_nw230).ArraySet1(p0, 7)
				(_nw230).ArraySet1(_1518_i2, 8)
				(_nw230).ArraySet1(_dafny.IntOfInt64(-59), 9)
				(_nw230).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}(func(_1528_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F23()
				}))).Cardinality())).Plus(_1487_v0), 10)
				(_nw230).ArraySet1(_1489_v2, 11)
				(_nw230).ArraySet1((p0).Minus((_dafny.Zero).Minus(_1518_i2)), 12)
				(_nw230).ArraySet1(p0, 13)
				(_nw230).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1489_v2)), 14)
				(_nw230).ArraySet1(p0, 15)
				(_nw230).ArraySet1((func() _dafny.Int {
					if (_1526_v25).Contains(_1489_v2) {
						return (_1526_v25).Get(_1489_v2).(_dafny.Int)
					}
					return _1487_v0
				})(), 16)
				_1527_v26 = _nw230
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_1527_v26), 0))
				_ = _index208
				(_1527_v26).ArraySet1(_1487_v0, (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_1527_v26), 0))
				_ = _index209
				(_1527_v26).ArraySet1((func() _dafny.Int {
					if (_1525_v24).Contains(!((_1513_v14).F24())) {
						return (_1525_v24).Multiplicity(!((_1513_v14).F24()))
					}
					return (_dafny.Zero).Minus((_1516_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.IntOfUint32((_1516_v17).Cardinality()))).Uint32()).(_dafny.Int))
				})(), (_index209).Int())
				var _1529_v27 _dafny.Sequence
				_ = _1529_v27
				_1529_v27 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1491_v4, _1491_v4), _dafny.UnicodeSeqOfUtf8Bytes("ifjf"), _1491_v4)
				_1529_v27 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1529_v27, _1529_v27), _dafny.Companion_Sequence_.Concatenate(_1529_v27, _1529_v27))
				var _1530_v28 D11
				_ = _1530_v28
				_1530_v28 = Companion_D11_.Create_DC22_(_1527_v26)
				(globalState).F22 = (_1530_v28).Dtor_cf35()
				_1488_v1 = (_this).F24()
				var _1531_v29 _dafny.Map
				_ = _1531_v29
				_1531_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1491_v4)
				var _1532_v30 _dafny.Sequence
				_ = _1532_v30
				_1532_v30 = _dafny.SeqOf(_1488_v1, (_this).F27(), (_this).F27())
				var _1533_v31 *C0
				_ = _1533_v31
				var _nw231 *C0 = New_C0_()
				_ = _nw231
				_nw231.Ctor__(((_1531_v29).Update((_this).F30(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_1534_v14 T0) func(_dafny.Int) _dafny.CodePoint {
					return func(_1535_i4 _dafny.Int) _dafny.CodePoint {
						return _1534_v14.F23()
					}
				})(_1513_v14))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1491_v4)), (_1487_v0).Cmp(_dafny.IntOfUint32((_1532_v30).Cardinality())) >= 0, _1513_v14.F23(), (_this).F24())
				_1533_v31 = _nw231
			}
			_1488_v1 = (_1513_v14).F24()
			var _1536_v32 D12
			_ = _1536_v32
			_1536_v32 = Companion_D12_.Create_DC25_(!(p1), _1513_v14.F23())
			var _1537_v33 D12
			_ = _1537_v33
			_1537_v33 = Companion_D12_.Create_DC26_(_1536_v32)
			_1537_v33 = _1537_v33
			(globalState).F13 = (_this).F31()
		}
	}
}
func (_this *C7) F30() bool {
	{
		return _this._f30
	}
}
func (_this *C7) F31() bool {
	{
		return _this._f31
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f27 bool
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f27 = false
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_}
}

var _ T3 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F27() bool {
	return _this._f27
}
func (_this *C8) Ctor__(f27 bool) {
	{
		(_this)._f27 = f27
	}
}
func (_this *C8) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-598)
	}
}
func (_this *C8) M4(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		if (_this).F27() {
			var _1538_v0 _dafny.MultiSet
			_ = _1538_v0
			_1538_v0 = _dafny.MultiSetOf(false)
			var _1539_v1 _dafny.Set
			_ = _1539_v1
			_1539_v1 = _dafny.SetOf(Companion_D0_.Create_DC0_(_1538_v0))
			(globalState).F13 = (_1539_v1).IsDisjointFrom(_1539_v1)
			var _1540_v2 _dafny.Sequence
			_ = _1540_v2
			_1540_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ktgrmbaw")
			var _1541_v3 _dafny.Set
			_ = _1541_v3
			_1541_v3 = _dafny.SetOf(_dafny.IntOfUint32((_1540_v2).Cardinality()))
			var _1542_v4 _dafny.Sequence
			_ = _1542_v4
			_1542_v4 = _dafny.SeqOf(_1541_v3)
			var _1543_v5 _dafny.Int
			_ = _1543_v5
			_1543_v5 = _dafny.IntOfInt64(57)
			(globalState).F13 = !(_1541_v3).Equals(((_1542_v4).Select((Companion_Default___.SafeIndex(_1543_v5, _dafny.IntOfUint32((_1542_v4).Cardinality()))).Uint32()).(_dafny.Set)).Union(_dafny.SetOf(_dafny.IntOfInt64(639), _1543_v5)))
			var _1544_v6 _dafny.CodePoint
			_ = _1544_v6
			_1544_v6 = _dafny.CodePoint('y')
			_1544_v6 = _1544_v6
			var _1545_v7 _dafny.Sequence
			_ = _1545_v7
			_1545_v7 = _dafny.SeqOf(_1543_v5, _1543_v5, _1543_v5)
			var _1546_v8 _dafny.Sequence
			_ = _1546_v8
			_1546_v8 = _dafny.SeqOf(_1545_v7)
			var _1547_v9 _dafny.Array
			_ = _1547_v9
			var _nw232 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw232
			_1547_v9 = _nw232
			var _1548_v10 _dafny.Map
			_ = _1548_v10
			_1548_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1543_v5, _1547_v9)
			var _1549_v11 _dafny.Set
			_ = _1549_v11
			_1549_v11 = _dafny.SetOf((func() _dafny.Array {
				if (_1548_v10).Contains(_1543_v5) {
					return (_1548_v10).Get(_1543_v5).(_dafny.Array)
				}
				return _1547_v9
			})(), _1547_v9, _1547_v9, _1547_v9)
			var _1550_v12 _dafny.Sequence
			_ = _1550_v12
			_1550_v12 = _dafny.SeqOf((_this).F27(), (_this).F27())
			var _1551_v13 _dafny.Sequence
			_ = _1551_v13
			_1551_v13 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1550_v12, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1540_v2).Cardinality()), _dafny.IntOfUint32((_1550_v12).Cardinality()))).Uint32(), (_this).F27()))
			var _rhs234 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1543_v5, (_1549_v11).Cardinality(), _1543_v5), (Companion_Default___.SafeIndex(_1543_v5, _dafny.IntOfUint32((_dafny.SeqOf(_1543_v5, (_1549_v11).Cardinality(), _1543_v5)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1540_v2).Cardinality())), _1545_v7), _1546_v8)
			_ = _rhs234
			var _rhs235 _dafny.Int = _1543_v5
			_ = _rhs235
			var _rhs236 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
				if p0 {
					return _1551_v13
				}
				return _1551_v13
			})(), _dafny.SeqOf(_dafny.SeqOf(p0, !(!((_this).F27())))))
			_ = _rhs236
			var _rhs237 _dafny.Sequence = _1545_v7
			_ = _rhs237
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			var _lhs176 *GlobalState = globalState
			_ = _lhs176
			_1546_v8 = _rhs234
			_lhs175.F7 = _rhs235
			_lhs176.F13 = _rhs236
			_1545_v7 = _rhs237
			var _1552_v14 _dafny.Array
			_ = _1552_v14
			var _nw233 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
			_ = _nw233
			_1552_v14 = _nw233
			var _1553_v15 D1
			_ = _1553_v15
			_1553_v15 = Companion_D1_.Create_DC4_(_1544_v6)
			var _rhs238 _dafny.CodePoint = (_1553_v15).Dtor_cf6()
			_ = _rhs238
			var _rhs239 _dafny.Array = _1552_v14
			_ = _rhs239
			var _rhs240 _dafny.Int = (_1543_v5).Plus((_1543_v5).Plus(Companion_Default___.Fm0(p0, _1543_v5, globalState)))
			_ = _rhs240
			var _lhs177 *GlobalState = globalState
			_ = _lhs177
			_1544_v6 = _rhs238
			_1552_v14 = _rhs239
			_lhs177.F7 = _rhs240
		} else {
			var _1554_v16 _dafny.Int
			_ = _1554_v16
			_1554_v16 = _dafny.IntOfInt64(253)
			var _1555_v17 _dafny.Map
			_ = _1555_v17
			_1555_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F27(), _dafny.IntOfInt64(570), globalState), !((Companion_D0_.Create_DC2_((_this).F27(), _1554_v16, (_this).F27())).Dtor_cf4()))
			if (func() bool {
				if (_1555_v17).Contains(_dafny.IntOfInt64(166)) {
					return (_1555_v17).Get(_dafny.IntOfInt64(166)).(bool)
				}
				return p0
			})() {
				var _1556_v18 _dafny.Array
				_ = _1556_v18
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_41
				var _nw234 _dafny.Array
				_ = _nw234
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw234 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) _dafny.Int = (func(_1557_v16 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1558_i0 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1558_i0, _1557_v16)
						}
					})(_1554_v16)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw234 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw234).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw234).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1556_v18 = _nw234
				var _1559_v19 _dafny.Set
				_ = _1559_v19
				_1559_v19 = _dafny.SetOf(_dafny.IntOfInt64(409), _1554_v16, _1554_v16)
				var _rhs241 _dafny.Array = _1556_v18
				_ = _rhs241
				var _rhs242 _dafny.Int = _1554_v16
				_ = _rhs242
				var _rhs243 bool = (_1559_v19).IsDisjointFrom(_1559_v19)
				_ = _rhs243
				var _lhs178 *GlobalState = globalState
				_ = _lhs178
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				_lhs178.F22 = _rhs241
				_lhs179.F7 = _rhs242
				_lhs180.F13 = _rhs243
				var _1560_v20 _dafny.Array
				_ = _1560_v20
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_42
				var _nw235 _dafny.Array
				_ = _nw235
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw235 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) bool = (func(_1561_p0 bool) func(_dafny.Int) bool {
						return func(_1562_i1 _dafny.Int) bool {
							return _1561_p0
						}
					})(p0)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw235 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw235).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw235).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1560_v20 = _nw235
				var _1563_v21 _dafny.Array
				_ = _1563_v21
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_43
				var _nw236 _dafny.Array
				_ = _nw236
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw236 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) D0 = (func(_1564_v16 _dafny.Int) func(_dafny.Int) D0 {
						return func(_1565_i2 _dafny.Int) D0 {
							return Companion_D0_.Create_DC1_(_1564_v16)
						}
					})(_1554_v16)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw236 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw236).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw236).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1563_v21 = _nw236
				var _1566_v22 _dafny.Set
				_ = _1566_v22
				_1566_v22 = _dafny.SetOf(_1563_v21)
				var _rhs244 _dafny.Int = (_dafny.Zero).Minus((((_1566_v22).Union(_1566_v22)).Union(_1566_v22)).Cardinality())
				_ = _rhs244
				var _rhs245 _dafny.Array = _1560_v20
				_ = _rhs245
				var _rhs246 _dafny.Int = (_1554_v16).Plus(_1554_v16)
				_ = _rhs246
				var _lhs181 *GlobalState = globalState
				_ = _lhs181
				_1554_v16 = _rhs244
				_1560_v20 = _rhs245
				_lhs181.F7 = _rhs246
				var _1567_v23 _dafny.CodePoint
				_ = _1567_v23
				_1567_v23 = _dafny.CodePoint('o')
				var _1568_v24 _dafny.Array
				_ = _1568_v24
				var _nwElement0_48 _dafny.CodePoint = _1567_v23
				_ = _nwElement0_48
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(7))
				_ = _nw237
				(_nw237).ArraySet1CodePoint(_nwElement0_48, 0)
				(_nw237).ArraySet1CodePoint(_1567_v23, 1)
				(_nw237).ArraySet1CodePoint(_1567_v23, 2)
				(_nw237).ArraySet1CodePoint(_1567_v23, 3)
				(_nw237).ArraySet1CodePoint(_1567_v23, 4)
				(_nw237).ArraySet1CodePoint(_1567_v23, 5)
				(_nw237).ArraySet1CodePoint(_1567_v23, 6)
				_1568_v24 = _nw237
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_1568_v24), 0))
				_ = _index210
				(_1568_v24).ArraySet1CodePoint(_1567_v23, (_index210).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_1568_v24), 0))
				_ = _index211
				(_1568_v24).ArraySet1CodePoint(_1567_v23, (_index211).Int())
				(globalState).F7 = Companion_Default___.Fm0((_this).F27(), _1554_v16, globalState)
				var _1569_v25 D0
				_ = _1569_v25
				_1569_v25 = Companion_D0_.Create_DC2_(p0, _1554_v16, p0)
				var _1570_v26 _dafny.Set
				_ = _1570_v26
				_1570_v26 = _dafny.SetOf((func() bool {
					if (_this).F27() {
						return p0
					}
					return !(p0)
				})(), (_1569_v25).Dtor_cf2())
				_1570_v26 = (_1570_v26).Difference((_1570_v26).Union(_1570_v26))
			} else {
				var _1571_v27 _dafny.Sequence
				_ = _1571_v27
				_1571_v27 = _dafny.SeqOf(_1554_v16)
				var _1572_v28 _dafny.Sequence
				_ = _1572_v28
				_1572_v28 = _dafny.UnicodeSeqOfUtf8Bytes("lwpgw")
				var _1573_v29 _dafny.Array
				_ = _1573_v29
				var _nwElement0_49 _dafny.Int = _1554_v16
				_ = _nwElement0_49
				var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(10))
				_ = _nw238
				(_nw238).ArraySet1(_nwElement0_49, 0)
				(_nw238).ArraySet1(_1554_v16, 1)
				(_nw238).ArraySet1((func() _dafny.Int {
					if (_this).F27() {
						return _dafny.IntOfInt64(137)
					}
					return _1554_v16
				})(), 2)
				(_nw238).ArraySet1(_1554_v16, 3)
				(_nw238).ArraySet1(_1554_v16, 4)
				(_nw238).ArraySet1(_1554_v16, 5)
				(_nw238).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_1571_v27).Cardinality())), _1554_v16), 6)
				(_nw238).ArraySet1(_1554_v16, 7)
				(_nw238).ArraySet1(Companion_Default___.Fm0(p0, _1554_v16, globalState), 8)
				(_nw238).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1572_v28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(2))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg99 _dafny.Int) interface{} {
						return coer99(arg99)
					}
				}(func(_1574_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})))).Cardinality()), 9)
				_1573_v29 = _nw238
				(globalState).F22 = _1573_v29
				var _1575_v30 _dafny.Sequence
				_ = _1575_v30
				_1575_v30 = _dafny.SeqOf((_this).F27())
				var _1576_v31 D0
				_ = _1576_v31
				_1576_v31 = Companion_D0_.Create_DC1_(_1554_v16)
				var _1577_v32 _dafny.Map
				_ = _1577_v32
				_1577_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_1575_v30, p0), _1576_v31)
				var _1578_v33 _dafny.Sequence
				_ = _1578_v33
				_1578_v33 = _dafny.SeqOf(_1571_v27, _dafny.SeqOf(_1554_v16))
				var _1579_v34 _dafny.CodePoint
				_ = _1579_v34
				_1579_v34 = _dafny.CodePoint('c')
				var _1580_v35 _dafny.Array
				_ = _1580_v35
				var _nwElement0_50 _dafny.CodePoint = _1579_v34
				_ = _nwElement0_50
				var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(19))
				_ = _nw239
				(_nw239).ArraySet1CodePoint(_nwElement0_50, 0)
				(_nw239).ArraySet1CodePoint(_1579_v34, 1)
				(_nw239).ArraySet1CodePoint(_1579_v34, 2)
				(_nw239).ArraySet1CodePoint(_1579_v34, 3)
				(_nw239).ArraySet1CodePoint(_1579_v34, 4)
				(_nw239).ArraySet1CodePoint(_1579_v34, 5)
				(_nw239).ArraySet1CodePoint(_1579_v34, 6)
				(_nw239).ArraySet1CodePoint(_1579_v34, 7)
				(_nw239).ArraySet1CodePoint(_1579_v34, 8)
				(_nw239).ArraySet1CodePoint(_1579_v34, 9)
				(_nw239).ArraySet1CodePoint(_1579_v34, 10)
				(_nw239).ArraySet1CodePoint(_1579_v34, 11)
				(_nw239).ArraySet1CodePoint(_1579_v34, 12)
				(_nw239).ArraySet1CodePoint(_1579_v34, 13)
				(_nw239).ArraySet1CodePoint(_1579_v34, 14)
				(_nw239).ArraySet1CodePoint(_1579_v34, 15)
				(_nw239).ArraySet1CodePoint(_1579_v34, 16)
				(_nw239).ArraySet1CodePoint(_1579_v34, 17)
				(_nw239).ArraySet1CodePoint(_1579_v34, 18)
				_1580_v35 = _nw239
				var _1581_v36 D12
				_ = _1581_v36
				_1581_v36 = Companion_D12_.Create_DC24_(_1580_v35)
				var _1582_v37 T3
				_ = _1582_v37
				var _nw240 *C7 = New_C7_()
				_ = _nw240
				_nw240.Ctor__((_this).F27(), !(true), (_1581_v36).Dtor_cf40(), _1579_v34, (_this).F27(), p0)
				_1582_v37 = _nw240
				var _1583_v38 _dafny.Map
				_ = _1583_v38
				_1583_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1582_v37, (_this).F27())
				var _rhs247 _dafny.Map = Companion_Default___.Fm13(Companion_Default___.Fm14((_this).F27(), globalState), !(_dafny.Companion_Sequence_.IsProperPrefixOf(_1578_v33, _dafny.Companion_Sequence_.Update(_1578_v33, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.IntOfUint32((_1578_v33).Cardinality()))).Uint32(), _dafny.SeqOf(_1554_v16)))), _1554_v16, (func() bool {
					if (_1583_v38).Contains(_1582_v37) {
						return (_1583_v38).Get(_1582_v37).(bool)
					}
					return (_1582_v37).F27()
				})(), globalState)
				_ = _rhs247
				var _rhs248 _dafny.Int = Companion_Default___.SafeDivisionInt((_1571_v27).Select((Companion_Default___.SafeIndex(_1554_v16, _dafny.IntOfUint32((_1571_v27).Cardinality()))).Uint32()).(_dafny.Int), _1554_v16)
				_ = _rhs248
				var _lhs182 *GlobalState = globalState
				_ = _lhs182
				_1577_v32 = _rhs247
				_lhs182.F7 = _rhs248
				_1572_v28 = _dafny.Companion_Sequence_.Concatenate(_1572_v28, _1572_v28)
				var _nw241 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw241
				_1573_v29 = _nw241
				var _1584_v39 _dafny.Array
				_ = _1584_v39
				var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
				_ = _nw242
				_1584_v39 = _nw242
				var _1585_v40 _dafny.Map
				_ = _1585_v40
				_1585_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1575_v30).Cardinality()), Companion_D14_.Create_DC30_(_1584_v39))
				var _1586_v41 _dafny.Sequence
				_ = _1586_v41
				_1586_v41 = _dafny.SeqOf(_1585_v40, _1585_v40, _1585_v40)
				(globalState).F7 = (_this).Fm9(_1571_v27, _dafny.IntOfUint32((_1586_v41).Cardinality()), !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("g"), _dafny.CodePoint('f')), globalState)
			}
			var _1587_v42 _dafny.Set
			_ = _1587_v42
			_1587_v42 = _dafny.SetOf(_dafny.IntOfInt64(-926))
			var _1588_v43 _dafny.Map
			_ = _1588_v43
			_1588_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1587_v42, _1587_v42)
			var _1589_v44 _dafny.Map
			_ = _1589_v44
			_1589_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				if (_1588_v43).Contains(_1587_v42) {
					return (_1588_v43).Get(_1587_v42).(_dafny.Set)
				}
				return _1587_v42
			})(), _1554_v16)
			var _1590_v45 _dafny.MultiSet
			_ = _1590_v45
			_1590_v45 = _dafny.MultiSetOf(_1554_v16)
			_1554_v16 = Companion_Default___.SafeModuloInt(((_1589_v44).Update(_1587_v42, _dafny.IntOfInt64(101))).Cardinality(), (func() _dafny.Int {
				if p0 {
					return (_1590_v45).Cardinality()
				}
				return _1554_v16
			})())
			(globalState).F7 = (_dafny.IntOfInt64(363)).Minus(Companion_Default___.SafeDivisionInt(_1554_v16, _1554_v16))
			var _1591_v46 _dafny.Map
			_ = _1591_v46
			_1591_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F27())
			var _source24 D11 = Companion_D11_.Create_DC23_((_this).F27(), (_1554_v16).Times(_1554_v16), _1591_v46, Companion_D5_.Create_DC14_(p0))
			_ = _source24
			if _source24.Is_DC23() {
				var _1592___mcc_h0 bool = _source24.Get_().(D11_DC23).Cf36
				_ = _1592___mcc_h0
				var _1593___mcc_h1 _dafny.Int = _source24.Get_().(D11_DC23).Cf37
				_ = _1593___mcc_h1
				var _1594___mcc_h2 _dafny.Map = _source24.Get_().(D11_DC23).Cf38
				_ = _1594___mcc_h2
				var _1595___mcc_h3 D5 = _source24.Get_().(D11_DC23).Cf39
				_ = _1595___mcc_h3
				var _1596_cf39 D5 = _1595___mcc_h3
				_ = _1596_cf39
				var _1597_cf38 _dafny.Map = _1594___mcc_h2
				_ = _1597_cf38
				var _1598_cf37 _dafny.Int = _1593___mcc_h1
				_ = _1598_cf37
				var _1599_cf36 bool = _1592___mcc_h0
				_ = _1599_cf36
				_1598_cf37 = _1554_v16
				var _1600_v47 _dafny.Array
				_ = _1600_v47
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_44
				var _nw243 _dafny.Array
				_ = _nw243
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw243 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) _dafny.Int = (func(_1601_cf37 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1602_i4 _dafny.Int) _dafny.Int {
							return (_1602_i4).Plus(_1601_cf37)
						}
					})(_1598_cf37)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw243 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw243).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw243).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1600_v47 = _nw243
				var _1603_v48 _dafny.MultiSet
				_ = _1603_v48
				_1603_v48 = _dafny.MultiSetOf(_1600_v47)
				var _1604_v49 _dafny.CodePoint
				_ = _1604_v49
				_1604_v49 = _dafny.CodePoint('e')
				var _1605_v50 _dafny.Sequence
				_ = _1605_v50
				_1605_v50 = _dafny.UnicodeSeqOfUtf8Bytes("o")
				var _1606_v51 _dafny.Sequence
				_ = _1606_v51
				_1606_v51 = _dafny.SeqOf(true, (_this).F27())
				var _1607_v52 _dafny.Map
				_ = _1607_v52
				_1607_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1598_cf37, _1606_v51)
				var _1608_v53 _dafny.Map
				_ = _1608_v53
				_1608_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_1605_v50, _1604_v49), (_1607_v52).Cardinality())
				var _rhs249 _dafny.MultiSet = _1603_v48
				_ = _rhs249
				var _rhs250 _dafny.Map = _1608_v53
				_ = _rhs250
				_1603_v48 = _rhs249
				_1608_v53 = _rhs250
				var _1609_v54 _dafny.Array
				_ = _1609_v54
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_45
				var _nw244 _dafny.Array
				_ = _nw244
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw244 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) bool = (func(_1610_p0 bool) func(_dafny.Int) bool {
						return func(_1611_i5 _dafny.Int) bool {
							return _1610_p0
						}
					})(p0)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw244 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw244).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw244).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1609_v54 = _nw244
				var _1612_v55 _dafny.Set
				_ = _1612_v55
				_1612_v55 = _dafny.SetOf(_1609_v54)
				(globalState).F10 = _1612_v55
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_1600_v47), 0))
				_ = _index212
				(_1600_v47).ArraySet1((_dafny.IntOfInt64(59)).Times(Companion_Default___.Fm0(p0, (_dafny.Zero).Minus(_1598_cf37), globalState)), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_1600_v47), 0))
				_ = _index213
				(_1600_v47).ArraySet1(_1554_v16, (_index213).Int())
				var _1613_v56 _dafny.Sequence
				_ = _1613_v56
				_1613_v56 = _dafny.SeqOf(_1554_v16, _1554_v16, _1554_v16)
				var _1614_v57 _dafny.Map
				_ = _1614_v57
				_1614_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1598_cf37, _1598_cf37)
				var _1615_v58 _dafny.Map
				_ = _1615_v58
				_1615_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1613_v56).Cardinality()), _dafny.IntOfUint32((_1606_v51).Cardinality())), _1614_v57)
				var _1616_v59 D9
				_ = _1616_v59
				_1616_v59 = Companion_D9_.Create_DC20_(_1605_v50, (_this).F27(), _1598_cf37)
				var _1617_v60 _dafny.Map
				_ = _1617_v60
				_1617_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1616_v59, Companion_Default___.Fm0(_1599_cf36, _1598_cf37, globalState))
				var _1618_v61 _dafny.Map
				_ = _1618_v61
				_1618_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1617_v60, _1605_v50)
				var _1619_v63 _dafny.MultiSet
				_ = _1619_v63
				_1619_v63 = _dafny.MultiSetOf(Companion_D9_.Create_DC20_(_1605_v50, p0, _1554_v16))
				var _1620_v65 D15
				_ = _1620_v65
				_1620_v65 = Companion_D15_.Create_DC33_(func() _dafny.Map {
					var _coll86 = _dafny.NewMapBuilder()
					_ = _coll86
					for _iter91 := _dafny.Iterate((func() _dafny.Set {
						var _coll87 = _dafny.NewBuilder()
						_ = _coll87
						for _iter92 := _dafny.Iterate((_1619_v63).Elements()); ; {
							_compr_87, _ok92 := _iter92()
							if !_ok92 {
								break
							}
							var _1621_v64 D9
							_1621_v64 = interface{}(_compr_87).(D9)
							if (_1619_v63).Contains(_1621_v64) {
								_coll87.Add(_1621_v64)
							}
						}
						return _coll87.ToSet()
					}()).Elements()); ; {
						_compr_86, _ok91 := _iter91()
						if !_ok91 {
							break
						}
						var _1622_v62 D9
						_1622_v62 = interface{}(_compr_86).(D9)
						if (func() _dafny.Set {
							var _coll88 = _dafny.NewBuilder()
							_ = _coll88
							for _iter93 := _dafny.Iterate((_1619_v63).Elements()); ; {
								_compr_88, _ok93 := _iter93()
								if !_ok93 {
									break
								}
								var _1623_v64 D9
								_1623_v64 = interface{}(_compr_88).(D9)
								if (_1619_v63).Contains(_1623_v64) {
									_coll88.Add(_1623_v64)
								}
							}
							return _coll88.ToSet()
						}()).Contains(_1622_v62) {
							_coll86.Add(_1622_v62, _1554_v16)
						}
					}
					return _coll86.ToMap()
				}())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_1600_v47), 0))
				_ = _index214
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_1600_v47), 0))
				_ = _index215
				var _rhs251 _dafny.Int = (_1615_v58).Cardinality()
				_ = _rhs251
				var _rhs252 _dafny.Int = (_1598_cf37).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0((_this).F27(), _1554_v16, globalState)))
				_ = _rhs252
				var _rhs253 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1618_v61).Contains((_1617_v60).Merge((_1620_v65).Dtor_cf54())) {
						return (_1618_v61).Get((_1617_v60).Merge((_1620_v65).Dtor_cf54())).(_dafny.Sequence)
					}
					return _1605_v50
				})()
				_ = _rhs253
				var _rhs254 _dafny.Sequence = (func() _dafny.Sequence {
					if !(_1599_cf36) {
						return _dafny.Companion_Sequence_.Update(Companion_Default___.Fm36(globalState), (Companion_Default___.SafeIndex(_1554_v16, _dafny.IntOfUint32((Companion_Default___.Fm36(globalState)).Cardinality()))).Uint32(), (_this).F27())
					}
					return _dafny.Companion_Sequence_.Concatenate(_1606_v51, _1606_v51)
				})()
				_ = _rhs254
				var _lhs183 _dafny.Array = _1600_v47
				_ = _lhs183
				var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_1600_v47), 0))
				_ = _lhs184
				var _lhs185 _dafny.Array = _1600_v47
				_ = _lhs185
				var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_1600_v47), 0))
				_ = _lhs186
				(_lhs183).ArraySet1(_rhs251, (_lhs184).Int())
				(_lhs185).ArraySet1(_rhs252, (_lhs186).Int())
				_1605_v50 = _rhs253
				_1606_v51 = _rhs254
			} else {
				var _1624___mcc_h4 _dafny.Array = _source24.Get_().(D11_DC22).Cf35
				_ = _1624___mcc_h4
				var _1625_cf35 _dafny.Array = _1624___mcc_h4
				_ = _1625_cf35
				var _1626_v66 _dafny.Array
				_ = _1626_v66
				var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
				_ = _nw245
				_1626_v66 = _nw245
				var _1627_v67 _dafny.Array
				_ = _1627_v67
				var _nw246 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw246
				_1627_v67 = _nw246
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_1626_v66), 0))
				_ = _index216
				(_1626_v66).ArraySet1(_1627_v67, (_index216).Int())
				var _1628_v68 _dafny.Array
				_ = _1628_v68
				var _nw247 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
				_ = _nw247
				_1628_v68 = _nw247
				var _1629_v69 _dafny.Map
				_ = _1629_v69
				_1629_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1628_v68, _1627_v67)
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_1626_v66), 0))
				_ = _index217
				(_1626_v66).ArraySet1((func() _dafny.Array {
					if (_1629_v69).Contains(_1628_v68) {
						return (_1629_v69).Get(_1628_v68).(_dafny.Array)
					}
					return _1627_v67
				})(), (_index217).Int())
				var _1630_v70 _dafny.CodePoint
				_ = _1630_v70
				_1630_v70 = _dafny.CodePoint('q')
				var _1631_v71 _dafny.Map
				_ = _1631_v71
				_1631_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1554_v16).Minus(_1554_v16), _1630_v70)
				_1631_v71 = (_1631_v71).Update(_1554_v16, _1630_v70)
				var _1632_v72 _dafny.Array
				_ = _1632_v72
				var _nw248 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(25))
				_ = _nw248
				_1632_v72 = _nw248
				var _1633_v73 _dafny.Map
				_ = _1633_v73
				_1633_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v16, _1632_v72)
				var _1634_v74 _dafny.Sequence
				_ = _1634_v74
				_1634_v74 = _dafny.SeqOf(_1554_v16, _1554_v16, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(949))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1635_v70 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1636_i6 _dafny.Int) _dafny.CodePoint {
						return _1635_v70
					}
				})(_1630_v70)))).Cardinality()))
				_1633_v73 = (_1633_v73).Update(_dafny.IntOfUint32((_1634_v74).Cardinality()), _1632_v72)
				var _1637_v75 _dafny.Map
				_ = _1637_v75
				_1637_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v16, _1590_v45)
				_1637_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v16, _1590_v45)
			}
			(globalState).F13 = !(Companion_Default___.Fm2(globalState)) || (!(p0))
		}
		var _1638_v76 _dafny.Int
		_ = _1638_v76
		_1638_v76 = _dafny.IntOfInt64(549)
		var _hi10 _dafny.Int = _1638_v76
		_ = _hi10
		for _1639_i7 := _1638_v76; _1639_i7.Cmp(_hi10) < 0; _1639_i7 = _1639_i7.Plus(_dafny.One) {
			(globalState).F13 = (p0) == (false)
			var _1640_v77 _dafny.Array
			_ = _1640_v77
			var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
			_ = _nw249
			_1640_v77 = _nw249
			var _1641_v78 _dafny.CodePoint
			_ = _1641_v78
			_1641_v78 = _dafny.CodePoint('d')
			var _1642_v79 _dafny.Sequence
			_ = _1642_v79
			_1642_v79 = _dafny.UnicodeSeqOfUtf8Bytes("bjjnna")
			var _1643_v80 D9
			_ = _1643_v80
			_1643_v80 = Companion_D9_.Create_DC20_(_1642_v79, p0, (_dafny.SetOf(_1639_i7)).Cardinality())
			var _1644_v81 _dafny.Sequence
			_ = _1644_v81
			_1644_v81 = _dafny.SeqOf(_1639_i7, (_1643_v80).Dtor_cf33())
			var _1645_v82 _dafny.Sequence
			_ = _1645_v82
			_1645_v82 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_1642_v79).Cardinality())), (_this).F27())).Cardinality(), _dafny.IntOfInt64(720), _dafny.IntOfUint32((_1644_v81).Cardinality()))
			var _1646_v83 *C7
			_ = _1646_v83
			var _nw250 *C7 = New_C7_()
			_ = _nw250
			_nw250.Ctor__(true, !(p0), _1640_v77, _1641_v78, !(!(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg101 _dafny.Int) interface{} {
					return coer101(arg101)
				}
			}((func(_1647_i7 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1648_i8 _dafny.Int) _dafny.Int {
					return _1647_i7
				}
			})(_1639_i7))), _1644_v81))), !(Companion_Default___.Fm2(globalState)) || (p0))
			_1646_v83 = _nw250
			var _1649_v84 _dafny.Map
			_ = _1649_v84
			_1649_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1644_v81)
			var _1650_v85 _dafny.MultiSet
			_ = _1650_v85
			_1650_v85 = _dafny.MultiSetOf(true)
			var _1651_v86 _dafny.Int
			_ = _1651_v86
			var _1652_v87 bool
			_ = _1652_v87
			var _1653_v88 _dafny.Int
			_ = _1653_v88
			var _out46 _dafny.Int
			_ = _out46
			var _out47 bool
			_ = _out47
			var _out48 _dafny.Int
			_ = _out48
			_out46, _out47, _out48 = Companion_Default___.M0(((func() _dafny.Map {
				if (_1646_v83).F31() {
					return _1649_v84
				}
				return _1649_v84
			})()).Cardinality(), _1650_v85, (_1646_v83).F30(), _1639_i7, globalState)
			_1651_v86 = _out46
			_1652_v87 = _out47
			_1653_v88 = _out48
			_1642_v79 = _dafny.UnicodeSeqOfUtf8Bytes("k")
		}
		var _1654_v89 _dafny.Sequence
		_ = _1654_v89
		_1654_v89 = _dafny.UnicodeSeqOfUtf8Bytes("vgc")
		_1654_v89 = _1654_v89
		_1638_v76 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1654_v89, _1654_v89), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg102 _dafny.Int) interface{} {
				return coer102(arg102)
			}
		}(func(_1655_i9 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})))).Cardinality())
		r0 = Companion_Default___.Fm2(globalState)
		_1638_v76 = Companion_Default___.SafeDivisionInt(_1638_v76, (_1638_v76).Plus(_1638_v76))
		r0 = (Companion_D5_.Create_DC13_((_dafny.Zero).Minus(_dafny.IntOfUint32((_1654_v89).Cardinality())), p0, p0)).Dtor_cf22()
		return r0
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f23 _dafny.CodePoint
	_f24 bool
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C9{}
var _ T0 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C9) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C9) F24() bool {
	return _this._f24
}
func (_this *C9) Ctor__(f23 _dafny.CodePoint, f24 bool) {
	{
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C9) Fm5(globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_((_this).F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24(), !((_this).F24()))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F24())))), (_this).F24())).Cardinality())).Cardinality(), (_this).F24())))
	}
}
func (_this *C9) Fm6(globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.SetOf(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(219)), Companion_D0_.Create_DC1_(_dafny.IntOfInt64(367)))).Difference(_dafny.SetOf(Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf(_this.F23(), _dafny.CodePoint('u'))).Cardinality())), Companion_D0_.Create_DC1_(_dafny.IntOfInt64(715)), Companion_D0_.Create_DC1_((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(380)), (_this).F24())).Cardinality())).Cardinality())))).Intersection((_dafny.SetOf(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(890)))).Union(_dafny.SetOf(Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("obaunck")).Cardinality())), Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-720)), Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality())), Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(83), _dafny.IntOfInt64(946))).Cardinality())))))).Cardinality()
	}
}
func (_this *C9) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus(((func() _dafny.Map {
			var _coll89 = _dafny.NewMapBuilder()
			_ = _coll89
			for _iter94 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(456), _dafny.IntOfInt64(762))); ; {
				_compr_89, _ok94 := _iter94()
				if !_ok94 {
					break
				}
				var _1656_v0 _dafny.Int
				_1656_v0 = interface{}(_compr_89).(_dafny.Int)
				if ((_dafny.IntOfInt64(456)).Cmp(_1656_v0) <= 0) && ((_1656_v0).Cmp(_dafny.IntOfInt64(762)) < 0) {
					_coll89.Add((_1656_v0).Plus(_dafny.IntOfInt64(36)), false)
				}
			}
			return _coll89.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(613), (_this).F24()))).Cardinality())).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(760))).Uint32(), func(coer103 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg103 _dafny.Int) interface{} {
				return coer103(arg103)
			}
		}(func(_1657_i0 _dafny.Int) D0 {
			return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(987)))
		}))).Cardinality()))) > 0
	}
}
func (_this *C9) Fm11(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C9) M1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) (_dafny.Sequence, D0, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var r2 bool = false
		_ = r2
		var _1658_v0 _dafny.Int
		_ = _1658_v0
		_1658_v0 = _dafny.IntOfInt64(135)
		var _1659_v1 D0
		_ = _1659_v1
		_1659_v1 = Companion_D0_.Create_DC2_(p1, _1658_v0, false)
		if (_1659_v1).Dtor_cf2() {
			var _1660_v2 _dafny.Sequence
			_ = _1660_v2
			_1660_v2 = _dafny.SeqOf(_dafny.IntOfInt64(357))
			var _1661_v3 _dafny.Map
			_ = _1661_v3
			_1661_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(p1, globalState), _dafny.IntOfUint32((_1660_v2).Cardinality()))
			var _1662_v4 _dafny.Sequence
			_ = _1662_v4
			_1662_v4 = _dafny.UnicodeSeqOfUtf8Bytes("rptgus")
			_1661_v3 = (_1661_v3).Update(_1662_v4, _1658_v0)
			var _1663_v5 _dafny.Array
			_ = _1663_v5
			var _nw251 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw251
			_1663_v5 = _nw251
			var _1664_v6 _dafny.Array
			_ = _1664_v6
			var _nw252 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw252
			_1664_v6 = _nw252
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_1664_v6), 0))
			_ = _index218
			(_1664_v6).ArraySet1(_1658_v0, (_index218).Int())
			var _1665_v7 _dafny.Array
			_ = _1665_v7
			var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
			_ = _nw253
			_1665_v7 = _nw253
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_1664_v6), 0))
			_ = _index219
			var _rhs255 _dafny.Array = (func() _dafny.Array {
				if p1 {
					return _1663_v5
				}
				return _1663_v5
			})()
			_ = _rhs255
			var _rhs256 _dafny.Int = _1658_v0
			_ = _rhs256
			var _rhs257 _dafny.Array = _1665_v7
			_ = _rhs257
			var _lhs187 _dafny.Array = _1664_v6
			_ = _lhs187
			var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_1664_v6), 0))
			_ = _lhs188
			_1663_v5 = _rhs255
			(_lhs187).ArraySet1(_rhs256, (_lhs188).Int())
			_1665_v7 = _rhs257
			_1658_v0 = (Companion_Default___.Fm12(_this.F23(), p1, _1658_v0, globalState)).Dtor_cf1()
			_1662_v4 = _1662_v4
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1663_v5), 0))
			_ = _index220
			(_1663_v5).ArraySet1((_this).F24(), (_index220).Int())
			var _1666_v8 _dafny.Set
			_ = _1666_v8
			_1666_v8 = _dafny.SetOf((_1664_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_1664_v6), 0))).Int()).(_dafny.Int))
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1663_v5), 0))
			_ = _index221
			(_1663_v5).ArraySet1((_1666_v8).IsDisjointFrom(_1666_v8), (_index221).Int())
		} else {
			var _1667_v9 _dafny.MultiSet
			_ = _1667_v9
			_1667_v9 = _dafny.MultiSetOf(p1)
			var _1668_v10 _dafny.Map
			_ = _1668_v10
			_1668_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1) || ((_this).F24()), (_1667_v9).Contains((_this).F24()))
			var _1669_v11 _dafny.Array
			_ = _1669_v11
			var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
			_ = _nw254
			_1669_v11 = _nw254
			var _1670_v12 T3
			_ = _1670_v12
			var _nw255 *C7 = New_C7_()
			_ = _nw255
			_nw255.Ctor__(p1, p1, _1669_v11, _dafny.CodePoint('j'), p1, true)
			_1670_v12 = _nw255
			var _1671_v13 _dafny.Map
			_ = _1671_v13
			_1671_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1670_v12, _this.F23())
			var _1672_v14 _dafny.Map
			_ = _1672_v14
			_1672_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v0, _1671_v13)
			_1668_v10 = (_1668_v10).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v0, _1671_v13)).Equals(_1672_v14), (_this).F24())
			(globalState).F7 = _1658_v0
			if p1 {
				var _1673_v15 _dafny.Sequence
				_ = _1673_v15
				_1673_v15 = _dafny.UnicodeSeqOfUtf8Bytes("hapquyumu")
				var _1674_v16 _dafny.Map
				_ = _1674_v16
				_1674_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1673_v15)
				var _1675_v17 T0
				_ = _1675_v17
				var _nw256 *C0 = New_C0_()
				_ = _nw256
				_nw256.Ctor__((_1674_v16).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1673_v15)), p1, _this.F23(), p1)
				_1675_v17 = _nw256
				_1675_v17 = _1675_v17
				var _pat_let_tv32 = _1670_v12
				_ = _pat_let_tv32
				r2 = (func(_pat_let47_0 D0) D0 {
					return func(_1676_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let48_0 bool) D0 {
							return func(_1677_dt__update_hcf2_h0 bool) D0 {
								return Companion_D0_.Create_DC2_(_1677_dt__update_hcf2_h0, (_1676_dt__update__tmp_h0).Dtor_cf3(), (_1676_dt__update__tmp_h0).Dtor_cf4())
							}(_pat_let48_0)
						}((_pat_let_tv32).F27())
					}(_pat_let47_0)
				}(_1659_v1)).Dtor_cf4()
				var _1678_v18 _dafny.Array
				_ = _1678_v18
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_46
				var _nw257 _dafny.Array
				_ = _nw257
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw257 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Int = (func(_1679_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1680_i0 _dafny.Int) _dafny.Int {
							return (_1680_i0).Plus(_1679_v0)
						}
					})(_1658_v0)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw257 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw257).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw257).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1678_v18 = _nw257
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1678_v18), 0))
				_ = _index222
				(_1678_v18).ArraySet1(_1658_v0, (_index222).Int())
				var _1681_v19 _dafny.Sequence
				_ = _1681_v19
				_1681_v19 = _dafny.SeqOf((_dafny.Zero).Minus((_1658_v0).Plus(_1658_v0)))
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1678_v18), 0))
				_ = _index223
				var _rhs258 _dafny.Int = _dafny.IntOfUint32((_1681_v19).Cardinality())
				_ = _rhs258
				var _rhs259 _dafny.Int = _1658_v0
				_ = _rhs259
				var _lhs189 _dafny.Array = _1678_v18
				_ = _lhs189
				var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1678_v18), 0))
				_ = _lhs190
				(_lhs189).ArraySet1(_rhs258, (_lhs190).Int())
				_1658_v0 = _rhs259
				var _1682_v20 *C8
				_ = _1682_v20
				var _nw258 *C8 = New_C8_()
				_ = _nw258
				_nw258.Ctor__((_this).F24())
				_1682_v20 = _nw258
				_1658_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1678_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1678_v18), 0))).Int()).(_dafny.Int), _1658_v0))
			} else {
				var _1683_v21 *C6
				_ = _1683_v21
				var _nw259 *C6 = New_C6_()
				_ = _nw259
				_nw259.Ctor__(p0, !((_1670_v12).F27()))
				_1683_v21 = _nw259
				var _1684_v22 _dafny.Sequence
				_ = _1684_v22
				_1684_v22 = _dafny.SeqOf(_1669_v11)
				var _1685_v23 D14
				_ = _1685_v23
				_1685_v23 = Companion_D14_.Create_DC31_((_1670_v12).F27())
				var _1686_v24 *C7
				_ = _1686_v24
				var _nw260 *C7 = New_C7_()
				_ = _nw260
				_nw260.Ctor__(false, (_1670_v12).F27(), (_1684_v22).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1658_v0), _dafny.IntOfUint32((_1684_v22).Cardinality()))).Uint32()).(_dafny.Array), _dafny.CodePoint('w'), (_this).F24(), (_1685_v23).Dtor_cf52())
				_1686_v24 = _nw260
				var _1687_v25 _dafny.Array
				_ = _1687_v25
				var _nw261 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw261
				_1687_v25 = _nw261
				var _1688_v26 *C5
				_ = _1688_v26
				var _nw262 *C5 = New_C5_()
				_ = _nw262
				_nw262.Ctor__(_1687_v25)
				_1688_v26 = _nw262
				var _1689_v27 _dafny.Array
				_ = _1689_v27
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_47
				var _nw263 _dafny.Array
				_ = _nw263
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw263 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) D15 = func(_1690_i1 _dafny.Int) D15 {
						return Companion_D15_.Create_DC34_(false)
					}
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw263 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw263).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw263).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1689_v27 = _nw263
				var _1691_v28 _dafny.Sequence
				_ = _1691_v28
				_1691_v28 = _dafny.SeqOf(_1689_v27)
				var _1692_v29 D16
				_ = _1692_v29
				_1692_v29 = Companion_D16_.Create_DC37_(_1688_v26)
				var _1693_v30 D16
				_ = _1693_v30
				_1693_v30 = Companion_D16_.Create_DC37_((_1692_v29).Dtor_cf61())
				var _rhs260 _dafny.Int = _dafny.IntOfUint32((_1691_v28).Cardinality())
				_ = _rhs260
				var _rhs261 _dafny.Int = _1658_v0
				_ = _rhs261
				var _rhs262 *C5 = (_1693_v30).Dtor_cf61()
				_ = _rhs262
				var _rhs263 _dafny.Int = (_this).Fm6(globalState)
				_ = _rhs263
				var _rhs264 bool = true
				_ = _rhs264
				var _lhs191 *GlobalState = globalState
				_ = _lhs191
				var _lhs192 *GlobalState = globalState
				_ = _lhs192
				var _lhs193 *GlobalState = globalState
				_ = _lhs193
				_1658_v0 = _rhs260
				_lhs191.F7 = _rhs261
				_1688_v26 = _rhs262
				_lhs192.F7 = _rhs263
				_lhs193.F13 = _rhs264
				var _1694_v31 _dafny.Sequence
				_ = _1694_v31
				_1694_v31 = _dafny.UnicodeSeqOfUtf8Bytes("dtuyeu")
				r0 = _1694_v31
				r2 = ((_1670_v12).F27()) || (false)
			}
			var _1695_v32 _dafny.Array
			_ = _1695_v32
			var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw264
			_1695_v32 = _nw264
			var _1696_v33 _dafny.Sequence
			_ = _1696_v33
			_1696_v33 = _dafny.SeqOf(_1695_v32)
			var _1697_v34 _dafny.Array
			_ = _1697_v34
			var _nwElement0_51 _dafny.Array = _1695_v32
			_ = _nwElement0_51
			var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(17))
			_ = _nw265
			(_nw265).ArraySet1(_nwElement0_51, 0)
			(_nw265).ArraySet1(_1695_v32, 1)
			(_nw265).ArraySet1(_1695_v32, 2)
			(_nw265).ArraySet1(_1695_v32, 3)
			(_nw265).ArraySet1(_1695_v32, 4)
			(_nw265).ArraySet1(_1695_v32, 5)
			(_nw265).ArraySet1(_1695_v32, 6)
			(_nw265).ArraySet1(_1695_v32, 7)
			(_nw265).ArraySet1(_1695_v32, 8)
			(_nw265).ArraySet1(_1695_v32, 9)
			(_nw265).ArraySet1(_1695_v32, 10)
			(_nw265).ArraySet1(_1695_v32, 11)
			(_nw265).ArraySet1(_1695_v32, 12)
			(_nw265).ArraySet1(_1695_v32, 13)
			(_nw265).ArraySet1(_1695_v32, 14)
			(_nw265).ArraySet1(_1695_v32, 15)
			(_nw265).ArraySet1((_1696_v33).Select((Companion_Default___.SafeIndex(_1658_v0, _dafny.IntOfUint32((_1696_v33).Cardinality()))).Uint32()).(_dafny.Array), 16)
			_1697_v34 = _nw265
			var _1698_v35 _dafny.Map
			_ = _1698_v35
			_1698_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1697_v34)
			_1698_v35 = (_1698_v35).Update(_dafny.CodePoint('y'), _1697_v34)
			var _1699_v36 _dafny.MultiSet
			_ = _1699_v36
			_1699_v36 = _dafny.MultiSetOf(p0)
			var _1700_v37 _dafny.MultiSet
			_ = _1700_v37
			_1700_v37 = _dafny.MultiSetOf(_1699_v36)
			var _1701_v38 *C0
			_ = _1701_v38
			var _nw266 *C0 = New_C0_()
			_ = _nw266
			_nw266.Ctor__((Companion_Default___.Fm48(globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1670_v12).F27(), _dafny.UnicodeSeqOfUtf8Bytes("ufkfma"))), (_this).F24(), _this.F23(), (_1700_v37).IsProperSubsetOf(_1700_v37))
			_1701_v38 = _nw266
		}
		var _1702_v39 D5
		_ = _1702_v39
		_1702_v39 = Companion_D5_.Create_DC13_(_1658_v0, !((_this).F24()), (_this).F24())
		var _source25 D5 = _1702_v39
		_ = _source25
		if _source25.Is_DC13() {
			var _1703___mcc_h0 _dafny.Int = _source25.Get_().(D5_DC13).Cf20
			_ = _1703___mcc_h0
			var _1704___mcc_h1 bool = _source25.Get_().(D5_DC13).Cf21
			_ = _1704___mcc_h1
			var _1705___mcc_h2 bool = _source25.Get_().(D5_DC13).Cf22
			_ = _1705___mcc_h2
			var _1706_cf22 bool = _1705___mcc_h2
			_ = _1706_cf22
			var _1707_cf21 bool = _1704___mcc_h1
			_ = _1707_cf21
			var _1708_cf20 _dafny.Int = _1703___mcc_h0
			_ = _1708_cf20
			var _1709_v40 _dafny.Sequence
			_ = _1709_v40
			_1709_v40 = _dafny.UnicodeSeqOfUtf8Bytes("ntvyfg")
			(globalState).F13 = _dafny.Companion_Sequence_.Equal(_1709_v40, _dafny.Companion_Sequence_.Concatenate(_1709_v40, _1709_v40))
			var _1710_v41 _dafny.Array
			_ = _1710_v41
			var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw267
			_1710_v41 = _nw267
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1710_v41), 0))
			_ = _index224
			(_1710_v41).ArraySet1((_this).Fm6(globalState), (_index224).Int())
			var _1711_v42 _dafny.Map
			_ = _1711_v42
			_1711_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(p1)), _1658_v0)
			var _1712_v43 _dafny.Sequence
			_ = _1712_v43
			_1712_v43 = _dafny.SeqOf(_1711_v42)
			var _1713_v44 _dafny.Set
			_ = _1713_v44
			_1713_v44 = _dafny.SetOf(p0, _this.F23(), _this.F23(), _this.F23(), _this.F23())
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1710_v41), 0))
			_ = _index225
			(_1710_v41).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1712_v43, _dafny.Companion_Sequence_.Update(_1712_v43, (Companion_Default___.SafeIndex((_1713_v44).Cardinality(), _dafny.IntOfUint32((_1712_v43).Cardinality()))).Uint32(), _1711_v42))).Cardinality()), (_index225).Int())
			var _1714_v45 _dafny.MultiSet
			_ = _1714_v45
			_1714_v45 = _dafny.MultiSetOf(p1)
			var _1715_v46 _dafny.Sequence
			_ = _1715_v46
			_1715_v46 = _dafny.SeqOf(_1714_v45)
			var _1716_v47 _dafny.Array
			_ = _1716_v47
			var _nw268 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
			_ = _nw268
			_1716_v47 = _nw268
			var _1717_v48 *C2
			_ = _1717_v48
			var _nw269 *C2 = New_C2_()
			_ = _nw269
			_nw269.Ctor__((_1715_v46).Select((Companion_Default___.SafeIndex((_1710_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1710_v41), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1715_v46).Cardinality()))).Uint32()).(_dafny.MultiSet), _1716_v47, _dafny.CodePoint('s'), p1)
			_1717_v48 = _nw269
			_1717_v48 = _1717_v48
			var _1718_v49 _dafny.Map
			_ = _1718_v49
			_1718_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1710_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1710_v41), 0))).Int()).(_dafny.Int)), _1707_cf21)
			_1718_v49 = (_1718_v49).Update((_1710_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1710_v41), 0))).Int()).(_dafny.Int), _1706_cf22)
		} else if _source25.Is_DC14() {
			var _1719___mcc_h3 bool = _source25.Get_().(D5_DC14).Cf23
			_ = _1719___mcc_h3
			var _1720_cf23 bool = _1719___mcc_h3
			_ = _1720_cf23
			(globalState).F13 = _1720_cf23
			_1659_v1 = _1659_v1
			if ((func() bool {
				if p1 {
					return (_this).Fm11(_1658_v0, (_this).F24(), globalState)
				}
				return p1
			})()) && ((_this).F24()) {
				var _1721_v50 _dafny.Sequence
				_ = _1721_v50
				_1721_v50 = _dafny.SeqOf(p1, false, (_this).F24(), _1720_cf23, _1720_cf23)
				var _1722_v51 _dafny.Map
				_ = _1722_v51
				_1722_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v0, _this.F23())
				var _1723_v52 _dafny.Map
				_ = _1723_v52
				_1723_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1721_v50).Select((Companion_Default___.SafeIndex((_1722_v51).Cardinality(), _dafny.IntOfUint32((_1721_v50).Cardinality()))).Uint32()).(bool), _1720_cf23)
				r2 = (func() bool {
					if false {
						return (_this).F24()
					}
					return (_this).Fm4((_1723_v52).Cardinality(), Companion_Default___.Fm36(globalState), globalState)
				})()
				var _1724_v53 _dafny.Sequence
				_ = _1724_v53
				_1724_v53 = _dafny.UnicodeSeqOfUtf8Bytes("hj")
				r0 = _1724_v53
				(globalState).F13 = _1720_cf23
				(globalState).F7 = _1658_v0
				var _1725_v54 _dafny.MultiSet
				_ = _1725_v54
				_1725_v54 = _dafny.MultiSetOf(!(false))
				var _1726_v55 _dafny.Int
				_ = _1726_v55
				var _1727_v56 bool
				_ = _1727_v56
				var _1728_v57 _dafny.Int
				_ = _1728_v57
				var _out49 _dafny.Int
				_ = _out49
				var _out50 bool
				_ = _out50
				var _out51 _dafny.Int
				_ = _out51
				_out49, _out50, _out51 = Companion_Default___.M0(_1658_v0, (_dafny.MultiSetOf(p1, (_this).F24())).Intersection(_1725_v54), p1, Companion_Default___.SafeDivisionInt(_1658_v0, _dafny.IntOfInt64(878)), globalState)
				_1726_v55 = _out49
				_1727_v56 = _out50
				_1728_v57 = _out51
			} else {
				var _1729_v58 _dafny.Array
				_ = _1729_v58
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_48
				var _nw270 _dafny.Array
				_ = _nw270
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw270 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) bool = func(_1730_i2 _dafny.Int) bool {
						return false
					}
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw270 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw270).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw270).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1729_v58 = _nw270
				var _1731_v59 _dafny.Map
				_ = _1731_v59
				_1731_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1729_v58, p0)
				var _1732_v60 _dafny.Map
				_ = _1732_v60
				_1732_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_1658_v0, _1658_v0), _1731_v59)
				_1732_v60 = (_1732_v60).Update(_1658_v0, _1731_v59)
				var _1733_v61 _dafny.Array
				_ = _1733_v61
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_49
				var _nw271 _dafny.Array
				_ = _nw271
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw271 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Int = (func(_1734_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1735_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1735_i3, _1734_v0)
						}
					})(_1658_v0)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw271 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw271).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw271).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1733_v61 = _nw271
				var _1736_v62 _dafny.Sequence
				_ = _1736_v62
				_1736_v62 = _dafny.SeqOf(p1)
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1733_v61), 0))
				_ = _index226
				(_1733_v61).ArraySet1((_1658_v0).Times(Companion_Default___.Fm0(_1720_cf23, _dafny.IntOfUint32((_1736_v62).Cardinality()), globalState)), (_index226).Int())
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1733_v61), 0))
				_ = _index227
				(_1733_v61).ArraySet1(_1658_v0, (_index227).Int())
				(globalState).F7 = _1658_v0
				var _1737_v63 _dafny.Sequence
				_ = _1737_v63
				_1737_v63 = _dafny.SeqOf(p0, p0)
				var _1738_v64 _dafny.Array
				_ = _1738_v64
				var _nwElement0_52 _dafny.CodePoint = p0
				_ = _nwElement0_52
				var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(27))
				_ = _nw272
				(_nw272).ArraySet1CodePoint(_nwElement0_52, 0)
				(_nw272).ArraySet1CodePoint(_this.F23(), 1)
				(_nw272).ArraySet1CodePoint((_1737_v63).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1737_v63).Cardinality()), _dafny.IntOfUint32((_1737_v63).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
				(_nw272).ArraySet1CodePoint(_this.F23(), 3)
				(_nw272).ArraySet1CodePoint(p0, 4)
				(_nw272).ArraySet1CodePoint(p0, 5)
				(_nw272).ArraySet1CodePoint((_1737_v63).Select((Companion_Default___.SafeIndex(_1658_v0, _dafny.IntOfUint32((_1737_v63).Cardinality()))).Uint32()).(_dafny.CodePoint), 6)
				(_nw272).ArraySet1CodePoint(_this.F23(), 7)
				(_nw272).ArraySet1CodePoint(_this.F23(), 8)
				(_nw272).ArraySet1CodePoint(_this.F23(), 9)
				(_nw272).ArraySet1CodePoint(p0, 10)
				(_nw272).ArraySet1CodePoint(p0, 11)
				(_nw272).ArraySet1CodePoint(_this.F23(), 12)
				(_nw272).ArraySet1CodePoint(_dafny.CodePoint('u'), 13)
				(_nw272).ArraySet1CodePoint(p0, 14)
				(_nw272).ArraySet1CodePoint(_this.F23(), 15)
				(_nw272).ArraySet1CodePoint(_this.F23(), 16)
				(_nw272).ArraySet1CodePoint(p0, 17)
				(_nw272).ArraySet1CodePoint(_this.F23(), 18)
				(_nw272).ArraySet1CodePoint(p0, 19)
				(_nw272).ArraySet1CodePoint(p0, 20)
				(_nw272).ArraySet1CodePoint((_1737_v63).Select((Companion_Default___.SafeIndex((_1733_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1733_v61), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1737_v63).Cardinality()))).Uint32()).(_dafny.CodePoint), 21)
				(_nw272).ArraySet1CodePoint(p0, 22)
				(_nw272).ArraySet1CodePoint(p0, 23)
				(_nw272).ArraySet1CodePoint(Companion_Default___.Fm29((_1733_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1733_v61), 0))).Int()).(_dafny.Int), p1, _1720_cf23, _1658_v0, globalState), 24)
				(_nw272).ArraySet1CodePoint(p0, 25)
				(_nw272).ArraySet1CodePoint(p0, 26)
				_1738_v64 = _nw272
				var _1739_v65 D12
				_ = _1739_v65
				_1739_v65 = Companion_D12_.Create_DC24_(_1738_v64)
				var _1740_v66 *C7
				_ = _1740_v66
				var _nw273 *C7 = New_C7_()
				_ = _nw273
				_nw273.Ctor__((func() bool {
					if true {
						return (_this).F24()
					}
					return true
				})(), (_this).F24(), (_1739_v65).Dtor_cf40(), _this.F23(), !(((_1733_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1733_v61), 0))).Int()).(_dafny.Int)).Cmp(_1658_v0) > 0), p1)
				_1740_v66 = _nw273
				var _1741_v67 _dafny.Array
				_ = _1741_v67
				var _nw274 _dafny.Array = _dafny.NewArrayWithValue(Companion_D13_.Default(), _dafny.IntOfInt64(6))
				_ = _nw274
				_1741_v67 = _nw274
				var _1742_v68 D13
				_ = _1742_v68
				_1742_v68 = Companion_D13_.Create_DC29_((_1740_v66).F30(), (_1740_v66).F31())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1741_v67), 0))
				_ = _index228
				(_1741_v67).ArraySet1(_1742_v68, (_index228).Int())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1741_v67), 0))
				_ = _index229
				(_1741_v67).ArraySet1(_1742_v68, (_index229).Int())
			}
			_1720_cf23 = (_this).F24()
		} else {
			var _1743___mcc_h4 _dafny.Set = _source25.Get_().(D5_DC12).Cf19
			_ = _1743___mcc_h4
			var _1744_cf19 _dafny.Set = _1743___mcc_h4
			_ = _1744_cf19
			var _1745_v69 _dafny.Array
			_ = _1745_v69
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_50
			var _nw275 _dafny.Array
			_ = _nw275
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw275 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Int = (func(_1746_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1747_i4 _dafny.Int) _dafny.Int {
						return (_1747_i4).Minus(_1746_v0)
					}
				})(_1658_v0)
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw275 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw275).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw275).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_1745_v69 = _nw275
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1745_v69), 0))
			_ = _index230
			(_1745_v69).ArraySet1(_1658_v0, (_index230).Int())
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1745_v69), 0))
			_ = _index231
			(_1745_v69).ArraySet1(_1658_v0, (_index231).Int())
			var _1748_v70 _dafny.Sequence
			_ = _1748_v70
			_1748_v70 = _dafny.SeqOf(_1658_v0)
			var _rhs265 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1748_v70, _1748_v70)
			_ = _rhs265
			var _rhs266 bool = !(p1)
			_ = _rhs266
			var _lhs194 *GlobalState = globalState
			_ = _lhs194
			_1748_v70 = _rhs265
			_lhs194.F13 = _rhs266
			var _1749_v71 _dafny.Array
			_ = _1749_v71
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_51
			var _nw276 _dafny.Array
			_ = _nw276
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw276 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) bool = (func(_1750_p1 bool) func(_dafny.Int) bool {
					return func(_1751_i5 _dafny.Int) bool {
						return (Companion_D2_.Create_DC7_(_1750_p1, (_this).F24())).Dtor_cf14()
					}
				})(p1)
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw276 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw276).ArraySet1(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw276).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_1749_v71 = _nw276
			var _1752_v72 _dafny.Sequence
			_ = _1752_v72
			_1752_v72 = _dafny.UnicodeSeqOfUtf8Bytes("cbxbjje")
			var _1753_v73 _dafny.Map
			_ = _1753_v73
			_1753_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1752_v72)
			var _1754_v74 _dafny.Sequence
			_ = _1754_v74
			_1754_v74 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_1753_v73).Contains(false) {
					return (_1753_v73).Get(false).(_dafny.Sequence)
				}
				return _1752_v72
			})())
			var _1755_v75 _dafny.Map
			_ = _1755_v75
			_1755_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p1)
			var _1756_v76 _dafny.Map
			_ = _1756_v76
			_1756_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1754_v74).Cardinality()), (_1755_v75).Cardinality()), _1749_v71)
			_1749_v71 = (func() _dafny.Array {
				if (_1756_v76).Contains(_1658_v0) {
					return (_1756_v76).Get(_1658_v0).(_dafny.Array)
				}
				return _1749_v71
			})()
			var _1757_v77 D0
			_ = _1757_v77
			_1757_v77 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-87))
			r1 = _1757_v77
		}
		if p1 {
			var _1758_v78 _dafny.Array
			_ = _1758_v78
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_52
			var _nw277 _dafny.Array
			_ = _nw277
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw277 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Int = (func(_1759_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1760_i6 _dafny.Int) _dafny.Int {
						return (_1760_i6).Minus(_1759_v0)
					}
				})(_1658_v0)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw277 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw277).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw277).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1758_v78 = _nw277
			(globalState).F22 = _1758_v78
			var _1761_v79 _dafny.Map
			_ = _1761_v79
			_1761_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
			(globalState).F7 = Companion_Default___.Fm0(p1, (_1761_v79).Cardinality(), globalState)
			(globalState).F13 = (_this).F24()
			var _1762_v80 T1
			_ = _1762_v80
			var _nw278 *C6 = New_C6_()
			_ = _nw278
			_nw278.Ctor__(_this.F23(), (_this).F24())
			_1762_v80 = _nw278
			_1762_v80 = _1762_v80
			(globalState).F13 = (func() bool {
				if false {
					return (_1762_v80).F24()
				}
				return (_this).F24()
			})()
		} else {
			var _rhs267 _dafny.Int = _dafny.IntOfInt64(86)
			_ = _rhs267
			var _rhs268 bool = (func() bool {
				if p1 {
					return (_this).Fm11(_1658_v0, p1, globalState)
				}
				return !(p1) || (p1)
			})()
			_ = _rhs268
			var _rhs269 bool = (_this).Fm11(_dafny.IntOfInt64(-862), (_this).F24(), globalState)
			_ = _rhs269
			var _rhs270 bool = (_this).Fm11(_1658_v0, p1, globalState)
			_ = _rhs270
			var _rhs271 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _rhs271
			var _lhs195 *GlobalState = globalState
			_ = _lhs195
			var _lhs196 *GlobalState = globalState
			_ = _lhs196
			var _lhs197 *GlobalState = globalState
			_ = _lhs197
			var _lhs198 *GlobalState = globalState
			_ = _lhs198
			_lhs195.F7 = _rhs267
			r2 = _rhs268
			_lhs196.F13 = _rhs269
			_lhs197.F13 = _rhs270
			_lhs198.F7 = _rhs271
			var _1763_v81 _dafny.Sequence
			_ = _1763_v81
			_1763_v81 = _dafny.UnicodeSeqOfUtf8Bytes("c")
			var _1764_v82 _dafny.Array
			_ = _1764_v82
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_53
			var _nw279 _dafny.Array
			_ = _nw279
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw279 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) bool = func(_1765_i7 _dafny.Int) bool {
					return (_this).F24()
				}
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw279 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw279).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw279).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_1764_v82 = _nw279
			var _1766_v83 _dafny.Set
			_ = _1766_v83
			_1766_v83 = _dafny.SetOf((_this).F24(), (_this).F24())
			var _1767_v84 D16
			_ = _1767_v84
			_1767_v84 = Companion_D16_.Create_DC38_(_dafny.Companion_Sequence_.Concatenate(_1763_v81, _1763_v81), _1764_v82, _1766_v83)
			var _source26 D16 = _1767_v84
			_ = _source26
			if _source26.Is_DC38() {
				var _1768___mcc_h5 _dafny.Sequence = _source26.Get_().(D16_DC38).Cf62
				_ = _1768___mcc_h5
				var _1769___mcc_h6 _dafny.Array = _source26.Get_().(D16_DC38).Cf63
				_ = _1769___mcc_h6
				var _1770___mcc_h7 _dafny.Set = _source26.Get_().(D16_DC38).Cf64
				_ = _1770___mcc_h7
				var _1771_cf64 _dafny.Set = _1770___mcc_h7
				_ = _1771_cf64
				var _1772_cf63 _dafny.Array = _1769___mcc_h6
				_ = _1772_cf63
				var _1773_cf62 _dafny.Sequence = _1768___mcc_h5
				_ = _1773_cf62
				var _1774_v85 _dafny.Map
				_ = _1774_v85
				_1774_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(524), _1658_v0)
				var _1775_v86 _dafny.Map
				_ = _1775_v86
				_1775_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v0, (_1774_v85).Cardinality())
				var _1776_v87 _dafny.Map
				_ = _1776_v87
				_1776_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1658_v0).Minus((_1775_v86).Cardinality()), _1658_v0)
				_1774_v85 = _1774_v85
				var _1777_v88 *C5
				_ = _1777_v88
				var _nw280 *C5 = New_C5_()
				_ = _nw280
				_nw280.Ctor__(_1772_cf63)
				_1777_v88 = _nw280
				var _1778_v89 _dafny.Array
				_ = _1778_v89
				var _nw281 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw281
				_1778_v89 = _nw281
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1778_v89), 0))
				_ = _index232
				(_1778_v89).ArraySet1((_dafny.Zero).Minus(_1658_v0), (_index232).Int())
				var _1779_v90 _dafny.Map
				_ = _1779_v90
				_1779_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p1)
				var _1780_v91 D5
				_ = _1780_v91
				_1780_v91 = Companion_D5_.Create_DC14_(p1)
				var _1781_v92 D11
				_ = _1781_v92
				_1781_v92 = Companion_D11_.Create_DC23_(p1, _dafny.IntOfUint32((_dafny.SeqOf(p1, !(p1))).Cardinality()), (_1779_v90).Update((_this).F24(), !((_this).F24())), _1780_v91)
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1778_v89), 0))
				_ = _index233
				var _rhs272 _dafny.Int = _1658_v0
				_ = _rhs272
				var _rhs273 _dafny.Int = (_1781_v92).Dtor_cf37()
				_ = _rhs273
				var _rhs274 bool = (_this).F24()
				_ = _rhs274
				var _lhs199 _dafny.Array = _1778_v89
				_ = _lhs199
				var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1778_v89), 0))
				_ = _lhs200
				var _lhs201 *GlobalState = globalState
				_ = _lhs201
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				(_lhs199).ArraySet1(_rhs272, (_lhs200).Int())
				_lhs201.F7 = _rhs273
				_lhs202.F13 = _rhs274
				_1779_v90 = (_1779_v90).Update((_this).F24(), (_this).F24())
			} else if _source26.Is_DC37() {
				var _1782___mcc_h8 *C5 = _source26.Get_().(D16_DC37).Cf61
				_ = _1782___mcc_h8
				var _1783_cf61 *C5 = _1782___mcc_h8
				_ = _1783_cf61
				var _1784_v93 _dafny.Array
				_ = _1784_v93
				var _nw282 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw282
				_1784_v93 = _nw282
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1784_v93), 0))
				_ = _index234
				(_1784_v93).ArraySet1(_1658_v0, (_index234).Int())
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1784_v93), 0))
				_ = _index235
				(_1784_v93).ArraySet1((_1658_v0).Times(_1658_v0), (_index235).Int())
				(_this).F23_set_(p0)
				_1658_v0 = Companion_Default___.Fm0(!((_1658_v0).Cmp(Companion_Default___.Fm0(true, _dafny.IntOfInt64(364), globalState)) >= 0), (_1784_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_1784_v93), 0))).Int()).(_dafny.Int), globalState)
				var _1785_v94 _dafny.Array
				_ = _1785_v94
				var _nw283 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
				_ = _nw283
				_1785_v94 = _nw283
				var _1786_v95 _dafny.MultiSet
				_ = _1786_v95
				_1786_v95 = _dafny.MultiSetOf(_1658_v0)
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1785_v94), 0))
				_ = _index236
				(_1785_v94).ArraySet1((_1786_v95).Intersection(_1786_v95), (_index236).Int())
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1785_v94), 0))
				_ = _index237
				(_1785_v94).ArraySet1(_1786_v95, (_index237).Int())
			} else {
				var _1787___mcc_h9 D16 = _source26.Get_().(D16_DC39).Cf65
				_ = _1787___mcc_h9
				var _1788_cf65 D16 = _1787___mcc_h9
				_ = _1788_cf65
				(globalState).F13 = !(p1) || (false)
				var _1789_v96 _dafny.Sequence
				_ = _1789_v96
				_1789_v96 = _dafny.SeqOf(_1658_v0)
				var _1790_v97 _dafny.Array
				_ = _1790_v97
				var _nwElement0_53 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1763_v81).Cardinality()))
				_ = _nwElement0_53
				var _nw284 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(25))
				_ = _nw284
				(_nw284).ArraySet1(_nwElement0_53, 0)
				(_nw284).ArraySet1((_1789_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.IntOfUint32((_1789_v96).Cardinality()))).Uint32()).(_dafny.Int), 1)
				(_nw284).ArraySet1(_1658_v0, 2)
				(_nw284).ArraySet1(_dafny.IntOfInt64(521), 3)
				(_nw284).ArraySet1((_dafny.Zero).Minus(_1658_v0), 4)
				(_nw284).ArraySet1((_dafny.MultiSetFromSeq(_1789_v96)).Cardinality(), 5)
				(_nw284).ArraySet1(_1658_v0, 6)
				(_nw284).ArraySet1((_dafny.Zero).Minus(_1658_v0), 7)
				(_nw284).ArraySet1((_1659_v1).Dtor_cf3(), 8)
				(_nw284).ArraySet1(_1658_v0, 9)
				(_nw284).ArraySet1(_1658_v0, 10)
				(_nw284).ArraySet1(_1658_v0, 11)
				(_nw284).ArraySet1(_1658_v0, 12)
				(_nw284).ArraySet1(_1658_v0, 13)
				(_nw284).ArraySet1(Companion_Default___.Fm0(!(p1), _1658_v0, globalState), 14)
				(_nw284).ArraySet1(_1658_v0, 15)
				(_nw284).ArraySet1(_1658_v0, 16)
				(_nw284).ArraySet1(_1658_v0, 17)
				(_nw284).ArraySet1(_1658_v0, 18)
				(_nw284).ArraySet1(_1658_v0, 19)
				(_nw284).ArraySet1(_1658_v0, 20)
				(_nw284).ArraySet1((_1658_v0).Times(_1658_v0), 21)
				(_nw284).ArraySet1(_1658_v0, 22)
				(_nw284).ArraySet1(_1658_v0, 23)
				(_nw284).ArraySet1((_1658_v0).Plus(_dafny.IntOfUint32((_1763_v81).Cardinality())), 24)
				_1790_v97 = _nw284
				var _1791_v98 D5
				_ = _1791_v98
				_1791_v98 = Companion_D5_.Create_DC14_(p1)
				var _1792_v99 D11
				_ = _1792_v99
				_1792_v99 = Companion_D11_.Create_DC23_(p1, _1658_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), _1791_v98)
				var _1793_v100 _dafny.Sequence
				_ = _1793_v100
				_1793_v100 = _dafny.SeqOf(_1792_v99, _1792_v99, _1792_v99, _1792_v99, _1792_v99)
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1790_v97), 0))
				_ = _index238
				(_1790_v97).ArraySet1((((_dafny.MultiSetFromSeq(_1793_v100)).Update(_1792_v99, Companion_Default___.Abs(_1658_v0))).Intersection(_dafny.MultiSetOf(_1792_v99))).Cardinality(), (_index238).Int())
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1790_v97), 0))
				_ = _index239
				(_1790_v97).ArraySet1(Companion_Default___.SafeDivisionInt(_1658_v0, ((_dafny.Zero).Minus(_1658_v0)).Minus(_1658_v0)), (_index239).Int())
				_1658_v0 = _dafny.IntOfInt64(375)
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1790_v97), 0))
				_ = _index240
				(_1790_v97).ArraySet1(_1658_v0, (_index240).Int())
			}
			var _1794_v101 _dafny.Array
			_ = _1794_v101
			var _nw285 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw285
			_1794_v101 = _nw285
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1794_v101), 0))
			_ = _index241
			(_1794_v101).ArraySet1(_1658_v0, (_index241).Int())
			var _1795_v102 _dafny.Map
			_ = _1795_v102
			_1795_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1658_v0)
			var _1796_v103 _dafny.Sequence
			_ = _1796_v103
			_1796_v103 = _dafny.SeqOf(p1)
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1794_v101), 0))
			_ = _index242
			var _rhs275 _dafny.Int = (func() _dafny.Int {
				if (_1795_v102).Contains((_this).F24()) {
					return (_1795_v102).Get((_this).F24()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_1796_v103).Cardinality())
			})()
			_ = _rhs275
			var _rhs276 _dafny.Int = _1658_v0
			_ = _rhs276
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 _dafny.Array = _1794_v101
			_ = _lhs204
			var _lhs205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1794_v101), 0))
			_ = _lhs205
			_lhs203.F7 = _rhs275
			(_lhs204).ArraySet1(_rhs276, (_lhs205).Int())
			var _1797_v104 *C8
			_ = _1797_v104
			var _nw286 *C8 = New_C8_()
			_ = _nw286
			_nw286.Ctor__((_this).F24())
			_1797_v104 = _nw286
			var _1798_v105 _dafny.Sequence
			_ = _1798_v105
			_1798_v105 = _dafny.SeqOf(_1658_v0, (_1794_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1794_v101), 0))).Int()).(_dafny.Int))
			var _1799_v106 _dafny.Sequence
			_ = _1799_v106
			_1799_v106 = _dafny.SeqOf(_1658_v0, (_1798_v105).Select((Companion_Default___.SafeIndex((_1794_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1794_v101), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1798_v105).Cardinality()))).Uint32()).(_dafny.Int))
			var _1800_v107 _dafny.MultiSet
			_ = _1800_v107
			_1800_v107 = _dafny.MultiSetOf(_1658_v0, _dafny.IntOfUint32((_1799_v106).Cardinality()), _1658_v0, _1658_v0)
			r2 = (func() bool {
				if !((_this).F24()) {
					return (_dafny.MultiSetOf((_1794_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1794_v101), 0))).Int()).(_dafny.Int))).IsDisjointFrom(_1800_v107)
				}
				return (_this).F24()
			})()
		}
		var _1801_v108 _dafny.Map
		_ = _1801_v108
		_1801_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1702_v39, (_this).F24())
		_1801_v108 = (_1801_v108).Update(_1702_v39, (_this).F24())
		if (_this).F24() {
			var _1802_v109 _dafny.Array
			_ = _1802_v109
			var _nw287 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw287
			_1802_v109 = _nw287
			var _1803_v110 _dafny.Map
			_ = _1803_v110
			_1803_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1802_v109, _dafny.CodePoint('y'))
			var _1804_v111 _dafny.Sequence
			_ = _1804_v111
			_1804_v111 = _dafny.UnicodeSeqOfUtf8Bytes("hu")
			var _1805_v112 _dafny.Array
			_ = _1805_v112
			var _nwElement0_54 _dafny.CodePoint = _this.F23()
			_ = _nwElement0_54
			var _nw288 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(11))
			_ = _nw288
			(_nw288).ArraySet1CodePoint(_nwElement0_54, 0)
			(_nw288).ArraySet1CodePoint(_dafny.CodePoint('p'), 1)
			(_nw288).ArraySet1CodePoint(p0, 2)
			(_nw288).ArraySet1CodePoint(p0, 3)
			(_nw288).ArraySet1CodePoint(_this.F23(), 4)
			(_nw288).ArraySet1CodePoint(p0, 5)
			(_nw288).ArraySet1CodePoint(_this.F23(), 6)
			(_nw288).ArraySet1CodePoint(_dafny.CodePoint('h'), 7)
			(_nw288).ArraySet1CodePoint(p0, 8)
			(_nw288).ArraySet1CodePoint(p0, 9)
			(_nw288).ArraySet1CodePoint((_1804_v111).Select((Companion_Default___.SafeIndex(_1658_v0, _dafny.IntOfUint32((_1804_v111).Cardinality()))).Uint32()).(_dafny.CodePoint), 10)
			_1805_v112 = _nw288
			var _1806_v113 *C3
			_ = _1806_v113
			var _nw289 *C3 = New_C3_()
			_ = _nw289
			_nw289.Ctor__(_1803_v110, _1805_v112, p0, (_this).F24())
			_1806_v113 = _nw289
			var _1807_v114 _dafny.Sequence
			_ = _1807_v114
			_1807_v114 = _dafny.SeqOf(_1658_v0, _1658_v0, _dafny.IntOfInt64(129), _1658_v0)
			var _1808_v115 _dafny.Map
			_ = _1808_v115
			_1808_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1806_v113, Companion_Default___.SafeDivisionInt((_1807_v114).Select((Companion_Default___.SafeIndex(_1658_v0, _dafny.IntOfUint32((_1807_v114).Cardinality()))).Uint32()).(_dafny.Int), _1658_v0))
			_1808_v115 = _1808_v115
			var _1809_v116 T3
			_ = _1809_v116
			var _nw290 *C7 = New_C7_()
			_ = _nw290
			_nw290.Ctor__(false, (_this).F24(), _1805_v112, p0, p1, !(true))
			_1809_v116 = _nw290
			var _1810_v117 _dafny.Array
			_ = _1810_v117
			var _nwElement0_55 T3 = _1809_v116
			_ = _nwElement0_55
			var _nw291 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(7))
			_ = _nw291
			(_nw291).ArraySet1(_nwElement0_55, 0)
			(_nw291).ArraySet1(_1809_v116, 1)
			(_nw291).ArraySet1(_1809_v116, 2)
			(_nw291).ArraySet1(_1809_v116, 3)
			(_nw291).ArraySet1(_1809_v116, 4)
			(_nw291).ArraySet1(_1809_v116, 5)
			(_nw291).ArraySet1(_1809_v116, 6)
			_1810_v117 = _nw291
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1810_v117), 0))
			_ = _index243
			(_1810_v117).ArraySet1(_1809_v116, (_index243).Int())
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1810_v117), 0))
			_ = _index244
			(_1810_v117).ArraySet1((func() T3 {
				if (_1809_v116).F27() {
					return _1809_v116
				}
				return _1809_v116
			})(), (_index244).Int())
			var _1811_v118 _dafny.Map
			_ = _1811_v118
			_1811_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1804_v111, (func() _dafny.Array {
				if !((_1809_v116).F27()) {
					return _1802_v109
				}
				return _1802_v109
			})())
			var _1812_v119 _dafny.Sequence
			_ = _1812_v119
			_1812_v119 = _dafny.SeqOf((_this).F24(), (_this).F24())
			var _1813_v120 D4
			_ = _1813_v120
			_1813_v120 = Companion_D4_.Create_DC10_(_1804_v111)
			var _1814_v121 _dafny.Map
			_ = _1814_v121
			_1814_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1804_v111)
			var _1815_v122 _dafny.Sequence
			_ = _1815_v122
			_1815_v122 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_1814_v121).Contains(p1) {
					return (_1814_v121).Get(p1).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("fpa")
			})(), _1804_v111, _dafny.UnicodeSeqOfUtf8Bytes("bfygf"))
			var _1816_v123 _dafny.MultiSet
			_ = _1816_v123
			_1816_v123 = _dafny.MultiSetOf((_1809_v116).F27())
			var _1817_v124 *C2
			_ = _1817_v124
			var _nw292 *C2 = New_C2_()
			_ = _nw292
			_nw292.Ctor__(_1816_v123, _1805_v112, _this.F23(), (_this).F24())
			_1817_v124 = _nw292
			var _1818_v125 _dafny.Sequence
			_ = _1818_v125
			_1818_v125 = _dafny.SeqOf(_1817_v124, _1817_v124)
			var _rhs277 bool = ((_this).Fm4((_this).Fm6(globalState), _1812_v119, globalState)) || (p1)
			_ = _rhs277
			var _rhs278 _dafny.Map = _1811_v118
			_ = _rhs278
			var _rhs279 _dafny.Int = (_1658_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1813_v120)).Cardinality()))
			_ = _rhs279
			var _rhs280 bool = (_this).F24()
			_ = _rhs280
			var _rhs281 _dafny.Sequence = (_1815_v122).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_1658_v0, _dafny.IntOfUint32((_1818_v125).Cardinality())), _dafny.IntOfUint32((_1815_v122).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs281
			var _lhs206 *GlobalState = globalState
			_ = _lhs206
			var _lhs207 *GlobalState = globalState
			_ = _lhs207
			r2 = _rhs277
			_1811_v118 = _rhs278
			_lhs206.F7 = _rhs279
			_lhs207.F13 = _rhs280
			_1804_v111 = _rhs281
			var _1819_v126 *C1
			_ = _1819_v126
			var _nw293 *C1 = New_C1_()
			_ = _nw293
			_nw293.Ctor__(p1, _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(Companion_Default___.Fm2(globalState), (_1809_v116).F27()), _1812_v119))
			_1819_v126 = _nw293
			r2 = (_this).F24()
		} else {
			var _1820_v127 _dafny.Map
			_ = _1820_v127
			_1820_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1658_v0).Times(_1658_v0), _1658_v0)
			_1820_v127 = _1820_v127
			(globalState).F7 = (_dafny.Zero).Minus(_1658_v0)
			_1659_v1 = _1659_v1
			var _1821_v128 _dafny.Array
			_ = _1821_v128
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_54
			var _nw294 _dafny.Array
			_ = _nw294
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw294 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.Int = (func(_1822_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1823_i8 _dafny.Int) _dafny.Int {
						return (_1823_i8).Times(_1822_v0)
					}
				})(_1658_v0)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw294 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw294).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw294).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_1821_v128 = _nw294
			var _1824_v129 _dafny.Set
			_ = _1824_v129
			_1824_v129 = _dafny.SetOf(_1658_v0)
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))
			_ = _index245
			(_1821_v128).ArraySet1((_1824_v129).Cardinality(), (_index245).Int())
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))
			_ = _index246
			(_1821_v128).ArraySet1(((func() _dafny.Int {
				if (_this).F24() {
					return _1658_v0
				}
				return _1658_v0
			})()).Plus(_1658_v0), (_index246).Int())
			var _1825_v130 _dafny.MultiSet
			_ = _1825_v130
			_1825_v130 = _dafny.MultiSetOf(_1658_v0)
			if ((_1825_v130).Difference(_dafny.MultiSetOf((_this).Fm6(globalState), _1658_v0))).IsSubsetOf(_1825_v130) {
				_1820_v127 = (_1820_v127).Update((_dafny.Zero).Minus(((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int)).Plus((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int))), (_dafny.Zero).Minus(_1658_v0))
				var _1826_v131 _dafny.Array
				_ = _1826_v131
				var _nw295 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw295
				_1826_v131 = _nw295
				var _1827_v132 *C5
				_ = _1827_v132
				var _nw296 *C5 = New_C5_()
				_ = _nw296
				_nw296.Ctor__(_1826_v131)
				_1827_v132 = _nw296
				r2 = (_this).F24()
				(globalState).F13 = (Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_1658_v0))).Cardinality()), !((_this).F24()), p1)).Dtor_cf22()
				var _1828_v133 _dafny.MultiSet
				_ = _1828_v133
				_1828_v133 = _dafny.MultiSetOf((_this).F24())
				var _1829_v134 D0
				_ = _1829_v134
				_1829_v134 = Companion_D0_.Create_DC0_(_1828_v133)
				_1829_v134 = _1829_v134
			} else {
				(globalState).F7 = _1658_v0
				var _1830_v135 _dafny.Map
				_ = _1830_v135
				_1830_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F24(), (_this).F24(), p1), p1)
				var _1831_v136 _dafny.Sequence
				_ = _1831_v136
				_1831_v136 = _dafny.SeqOf(p1)
				var _1832_v137 _dafny.Sequence
				_ = _1832_v137
				_1832_v137 = _dafny.UnicodeSeqOfUtf8Bytes("blrabrw")
				var _1833_v138 _dafny.Map
				_ = _1833_v138
				_1833_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v0, _1832_v137)
				var _1834_v139 _dafny.Map
				_ = _1834_v139
				_1834_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4((_1830_v135).Cardinality(), _1831_v136, globalState), (func() _dafny.Sequence {
					if (_1833_v138).Contains(_1658_v0) {
						return (_1833_v138).Get(_1658_v0).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-330))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg104 _dafny.Int) interface{} {
							return coer104(arg104)
						}
					}(func(_1835_i9 _dafny.Int) _dafny.CodePoint {
						return _this.F23()
					}))
				})())
				var _1836_v140 *C1
				_ = _1836_v140
				var _nw297 *C1 = New_C1_()
				_ = _nw297
				_nw297.Ctor__((_this).F24(), false)
				_1836_v140 = _nw297
				var _1837_v141 _dafny.Map
				_ = _1837_v141
				_1837_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1836_v140, _1658_v0)
				var _1838_v142 _dafny.Map
				_ = _1838_v142
				_1838_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1837_v141, (_this).F24())
				var _1839_v143 _dafny.Map
				_ = _1839_v143
				_1839_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F24()), (_dafny.Zero).Minus((_1838_v142).Cardinality()))
				var _1840_v144 _dafny.Map
				_ = _1840_v144
				_1840_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1839_v143, p1)
				var _1841_v145 *C0
				_ = _1841_v145
				var _nw298 *C0 = New_C0_()
				_ = _nw298
				_nw298.Ctor__(_1834_v139, (_this).F24(), Companion_Default___.Fm29((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int), p1, p1, (_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int), globalState), (_1840_v144).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_1831_v136).Cardinality())), (_1836_v140).F39())))
				_1841_v145 = _nw298
				(_this).F23_set_((_1832_v137).Select((Companion_Default___.SafeIndex(_1658_v0, _dafny.IntOfUint32((_1832_v137).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _1842_v146 _dafny.Sequence
				_ = _1842_v146
				_1842_v146 = _dafny.SeqOf((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int), ((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int)).Minus((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int)), _1658_v0, (_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int))
				(globalState).F7 = (_1842_v146).Select((Companion_Default___.SafeIndex((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1842_v146).Cardinality()))).Uint32()).(_dafny.Int)
				var _1843_v147 _dafny.MultiSet
				_ = _1843_v147
				_1843_v147 = _dafny.MultiSetOf((_1841_v145).F38())
				(globalState).F7 = (func() _dafny.Int {
					if !(_1843_v147).Contains(!(!((_1836_v140).F39()))) {
						return (_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int)
					}
					return Companion_Default___.SafeModuloInt((_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int), (_1821_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1821_v128), 0))).Int()).(_dafny.Int))
				})()
			}
		}
		var _1844_v148 _dafny.Array
		_ = _1844_v148
		var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw299
		_1844_v148 = _nw299
		var _1845_v149 _dafny.Sequence
		_ = _1845_v149
		_1845_v149 = _dafny.SeqOf(_1844_v148, _1844_v148)
		var _hi11 _dafny.Int = (func() _dafny.Int {
			if (_this).F24() {
				return _dafny.IntOfUint32((_1845_v149).Cardinality())
			}
			return _1658_v0
		})()
		_ = _hi11
		for _1846_i10 := _1658_v0; _1846_i10.Cmp(_hi11) < 0; _1846_i10 = _1846_i10.Plus(_dafny.One) {
			var _1847_v150 _dafny.Sequence
			_ = _1847_v150
			_1847_v150 = _dafny.UnicodeSeqOfUtf8Bytes("enqsgx")
			(globalState).F13 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_this).F24(), globalState), _1847_v150), _1847_v150)
			var _1848_v151 _dafny.Array
			_ = _1848_v151
			var _nw300 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw300
			_1848_v151 = _nw300
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1848_v151), 0))
			_ = _index247
			(_1848_v151).ArraySet1((_this).F24(), (_index247).Int())
			var _1849_v152 _dafny.Map
			_ = _1849_v152
			_1849_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1847_v150)
			var _1850_v153 _dafny.Sequence
			_ = _1850_v153
			_1850_v153 = _dafny.SeqOf(p1)
			var _1851_v154 T0
			_ = _1851_v154
			var _nw301 *C0 = New_C0_()
			_ = _nw301
			_nw301.Ctor__(_1849_v152, (_1850_v153).Select((Companion_Default___.SafeIndex(_1846_i10, _dafny.IntOfUint32((_1850_v153).Cardinality()))).Uint32()).(bool), _this.F23(), !((_this).F24()))
			_1851_v154 = _nw301
			var _1852_v155 _dafny.MultiSet
			_ = _1852_v155
			_1852_v155 = _dafny.MultiSetOf(_1846_i10)
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1848_v151), 0))
			_ = _index248
			var _rhs282 bool = false
			_ = _rhs282
			var _rhs283 _dafny.Int = (_dafny.Zero).Minus(_1846_i10)
			_ = _rhs283
			var _rhs284 T0 = _1851_v154
			_ = _rhs284
			var _rhs285 _dafny.MultiSet = (_1852_v155).Intersection(_dafny.MultiSetOf(_1846_i10, _1846_i10, _1658_v0))
			_ = _rhs285
			var _rhs286 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1850_v153, _dafny.SeqOf(true)), _dafny.Companion_Sequence_.Concatenate(_1850_v153, _1850_v153))
			_ = _rhs286
			var _lhs208 _dafny.Array = _1848_v151
			_ = _lhs208
			var _lhs209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1848_v151), 0))
			_ = _lhs209
			(_lhs208).ArraySet1(_rhs282, (_lhs209).Int())
			_1658_v0 = _rhs283
			_1851_v154 = _rhs284
			_1852_v155 = _rhs285
			_1850_v153 = _rhs286
			var _1853_v156 *C8
			_ = _1853_v156
			var _nw302 *C8 = New_C8_()
			_ = _nw302
			_nw302.Ctor__(p1)
			_1853_v156 = _nw302
			var _1854_v157 _dafny.MultiSet
			_ = _1854_v157
			_1854_v157 = _dafny.MultiSetOf(_1853_v156)
			var _1855_v158 _dafny.Set
			_ = _1855_v158
			_1855_v158 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lufwkuo")).Cardinality()), _1846_i10, ((_1854_v157).Update(_1853_v156, Companion_Default___.Abs(_1846_i10))).Cardinality())
			var _1856_v159 _dafny.Set
			_ = _1856_v159
			_1856_v159 = _dafny.SetOf((_1855_v158).Cardinality(), _1846_i10, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(114))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1857_v154 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_1858_i11 _dafny.Int) _dafny.CodePoint {
					return _1857_v154.F23()
				}
			})(_1851_v154)))).Cardinality()))
			(globalState).F7 = (func() _dafny.Int {
				if p1 {
					return Companion_Default___.SafeDivisionInt(_1658_v0, (_1856_v159).Cardinality())
				}
				return _1846_i10
			})()
			var _1859_v160 _dafny.Map
			_ = _1859_v160
			_1859_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("y"), _1851_v154.F23())
			(_1851_v154).F23_set_((func() _dafny.CodePoint {
				if (_1859_v160).Contains(_dafny.Companion_Sequence_.Concatenate(_1847_v150, _1847_v150)) {
					return (_1859_v160).Get(_dafny.Companion_Sequence_.Concatenate(_1847_v150, _1847_v150)).(_dafny.CodePoint)
				}
				return _1851_v154.F23()
			})())
		}
		r0 = _dafny.UnicodeSeqOfUtf8Bytes("oyguroyys")
		var _1860_v161 D0
		_ = _1860_v161
		_1860_v161 = Companion_D0_.Create_DC1_(_1658_v0)
		r1 = _1860_v161
		var _1861_v162 *C1
		_ = _1861_v162
		var _nw303 *C1 = New_C1_()
		_ = _nw303
		_nw303.Ctor__(p1, p1)
		_1861_v162 = _nw303
		var _1862_v163 _dafny.Set
		_ = _1862_v163
		_1862_v163 = _dafny.SetOf(_1861_v162, _1861_v162)
		r2 = (_1862_v163).IsProperSubsetOf(_1862_v163)
		return r0, r1, r2
	}
}
func (_this *C9) M2(globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _source27 D13 = Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _this.F23()))
		_ = _source27
		if _source27.Is_DC28() {
			var _1863___mcc_h0 bool = _source27.Get_().(D13_DC28).Cf45
			_ = _1863___mcc_h0
			var _1864___mcc_h1 _dafny.Int = _source27.Get_().(D13_DC28).Cf46
			_ = _1864___mcc_h1
			var _1865___mcc_h2 _dafny.Int = _source27.Get_().(D13_DC28).Cf47
			_ = _1865___mcc_h2
			var _1866___mcc_h3 _dafny.Int = _source27.Get_().(D13_DC28).Cf48
			_ = _1866___mcc_h3
			var _1867_cf48 _dafny.Int = _1866___mcc_h3
			_ = _1867_cf48
			var _1868_cf47 _dafny.Int = _1865___mcc_h2
			_ = _1868_cf47
			var _1869_cf46 _dafny.Int = _1864___mcc_h1
			_ = _1869_cf46
			var _1870_cf45 bool = _1863___mcc_h0
			_ = _1870_cf45
			var _1871_v0 _dafny.Array
			_ = _1871_v0
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_55
			var _nw304 _dafny.Array
			_ = _nw304
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw304 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) bool = (func(_1872_cf45 bool) func(_dafny.Int) bool {
					return func(_1873_i0 _dafny.Int) bool {
						return _1872_cf45
					}
				})(_1870_cf45)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw304 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw304).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw304).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_1871_v0 = _nw304
			var _1874_v1 _dafny.Sequence
			_ = _1874_v1
			_1874_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ig")
			var _1875_v2 _dafny.Map
			_ = _1875_v2
			_1875_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1871_v0, (_1874_v1).Select((Companion_Default___.SafeIndex(_1868_cf47, _dafny.IntOfUint32((_1874_v1).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _1876_v3 _dafny.Array
			_ = _1876_v3
			var _nw305 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
			_ = _nw305
			_1876_v3 = _nw305
			var _1877_v4 *C3
			_ = _1877_v4
			var _nw306 *C3 = New_C3_()
			_ = _nw306
			_nw306.Ctor__(_1875_v2, _1876_v3, _this.F23(), _1870_cf45)
			_1877_v4 = _nw306
			var _1878_v5 D2
			_ = _1878_v5
			_1878_v5 = Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1874_v1))
			var _1879_v6 _dafny.Sequence
			_ = _1879_v6
			_1879_v6 = Companion_Default___.Fm45((_this).F24(), _1878_v5, _dafny.IntOfUint32((_1874_v1).Cardinality()), globalState)
			var _1880_v7 _dafny.Map
			_ = _1880_v7
			_1880_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F24()) || (true), (_1879_v6))
			var _1881_v8 _dafny.Sequence
			_ = _1881_v8
			_1881_v8 = _dafny.SeqOf(_dafny.IntOfInt64(-557), _1868_cf47)
			_1880_v7 = (_1880_v7).Update((_this).F24(), _1881_v8)
			var _1882_v9 D12
			_ = _1882_v9
			_1882_v9 = Companion_D12_.Create_DC25_((_this).F24(), (Companion_D1_.Create_DC4_(_this.F23())).Dtor_cf6())
			if (func() bool {
				if _1870_cf45 {
					return (_this).F24()
				}
				return (_1882_v9).Dtor_cf41()
			})() {
				_1882_v9 = _1882_v9
				var _1883_v10 _dafny.Map
				_ = _1883_v10
				_1883_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1870_cf45, _dafny.UnicodeSeqOfUtf8Bytes("sworc"))
				var _1884_v11 T0
				_ = _1884_v11
				var _nw307 *C0 = New_C0_()
				_ = _nw307
				_nw307.Ctor__(_1883_v10, _1870_cf45, _this.F23(), _1870_cf45)
				_1884_v11 = _nw307
				var _1885_v12 T0
				_ = _1885_v12
				_1885_v12 = _1884_v11
				var _1886_v13 _dafny.Array
				_ = _1886_v13
				var _nwElement0_56 T0 = (_1885_v12)
				_ = _nwElement0_56
				var _nw308 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(14))
				_ = _nw308
				(_nw308).ArraySet1(_nwElement0_56, 0)
				(_nw308).ArraySet1(_1884_v11, 1)
				(_nw308).ArraySet1(_1884_v11, 2)
				(_nw308).ArraySet1(_1884_v11, 3)
				(_nw308).ArraySet1(_1884_v11, 4)
				(_nw308).ArraySet1((func() T0 {
					if (_1884_v11).F24() {
						return _1884_v11
					}
					return _1884_v11
				})(), 5)
				(_nw308).ArraySet1(_1884_v11, 6)
				(_nw308).ArraySet1(_1884_v11, 7)
				(_nw308).ArraySet1(_1884_v11, 8)
				(_nw308).ArraySet1(_1884_v11, 9)
				(_nw308).ArraySet1((func() T0 {
					if (_1884_v11).F24() {
						return _1884_v11
					}
					return _1884_v11
				})(), 10)
				(_nw308).ArraySet1(_1884_v11, 11)
				(_nw308).ArraySet1((func() T0 {
					if (_1884_v11).F24() {
						return _1884_v11
					}
					return _1884_v11
				})(), 12)
				(_nw308).ArraySet1(_1884_v11, 13)
				_1886_v13 = _nw308
				var _1887_v14 _dafny.MultiSet
				_ = _1887_v14
				_1887_v14 = _dafny.MultiSetOf(_1867_cf48, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1874_v1).Cardinality())))
				var _1888_v15 _dafny.Map
				_ = _1888_v15
				_1888_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1870_cf45, (Companion_D5_.Create_DC13_((_dafny.Zero).Minus(_1868_cf47), !((_this).F24()), (_1884_v11).F24())).Dtor_cf20())
				var _1889_v16 _dafny.Map
				_ = _1889_v16
				_1889_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1888_v15).Cardinality(), (_this).F24())
				var _1890_v17 _dafny.Array
				_ = _1890_v17
				var _nwElement0_57 _dafny.Int = _1869_cf46
				_ = _nwElement0_57
				var _nw309 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(19))
				_ = _nw309
				(_nw309).ArraySet1(_nwElement0_57, 0)
				(_nw309).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1874_v1).Cardinality()), _1867_cf48), 1)
				(_nw309).ArraySet1(_1867_cf48, 2)
				(_nw309).ArraySet1((_this).Fm6(globalState), 3)
				(_nw309).ArraySet1(_1867_cf48, 4)
				(_nw309).ArraySet1(_1868_cf47, 5)
				(_nw309).ArraySet1(Companion_Default___.SafeModuloInt(_1867_cf48, _dafny.IntOfUint32((_1874_v1).Cardinality())), 6)
				(_nw309).ArraySet1(_dafny.IntOfInt64(310), 7)
				(_nw309).ArraySet1(_dafny.IntOfInt64(592), 8)
				(_nw309).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_1887_v14).Contains(_1869_cf46) {
						return (_1887_v14).Multiplicity(_1869_cf46)
					}
					return (_1889_v16).Cardinality()
				})(), _1868_cf47), 9)
				(_nw309).ArraySet1(Companion_Default___.SafeDivisionInt(_1869_cf46, _1868_cf47), 10)
				(_nw309).ArraySet1((_dafny.SetOf(_1867_cf48, _dafny.IntOfUint32((_1874_v1).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_1867_cf48)), _1869_cf46, _1868_cf47)).Cardinality(), 11)
				(_nw309).ArraySet1(_1869_cf46, 12)
				(_nw309).ArraySet1(_1868_cf47, 13)
				(_nw309).ArraySet1(_1867_cf48, 14)
				(_nw309).ArraySet1(_dafny.IntOfInt64(132), 15)
				(_nw309).ArraySet1(_1868_cf47, 16)
				(_nw309).ArraySet1(_1869_cf46, 17)
				(_nw309).ArraySet1(Companion_Default___.SafeModuloInt(_1867_cf48, _1868_cf47), 18)
				_1890_v17 = _nw309
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1890_v17), 0))
				_ = _index249
				(_1890_v17).ArraySet1(_dafny.IntOfInt64(-225), (_index249).Int())
				var _1891_v18 _dafny.Array
				_ = _1891_v18
				var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
				_ = _nw310
				_1891_v18 = _nw310
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1891_v18), 0))
				_ = _index250
				(_1891_v18).ArraySet1(Companion_Default___.Fm49(globalState), (_index250).Int())
				var _1892_v19 D18
				_ = _1892_v19
				_1892_v19 = Companion_D18_.Create_DC41_(_1886_v13)
				var _1893_v20 _dafny.Sequence
				_ = _1893_v20
				_1893_v20 = _dafny.SeqOf((_1884_v11).F24())
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1890_v17), 0))
				_ = _index251
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1891_v18), 0))
				_ = _index252
				var _rhs287 _dafny.Int = (_dafny.Zero).Minus(_1868_cf47)
				_ = _rhs287
				var _rhs288 _dafny.Array = (_1892_v19).Dtor_cf67()
				_ = _rhs288
				var _rhs289 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("adnjjpgt"), _1874_v1)).Cardinality()), _1867_cf48)
				_ = _rhs289
				var _rhs290 bool = ((_this).Fm4(_1868_cf47, _1893_v20, globalState)) && ((_this).F24())
				_ = _rhs290
				var _rhs291 _dafny.Map = _1888_v15
				_ = _rhs291
				var _lhs210 *GlobalState = globalState
				_ = _lhs210
				var _lhs211 _dafny.Array = _1890_v17
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1890_v17), 0))
				_ = _lhs212
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 _dafny.Array = _1891_v18
				_ = _lhs214
				var _lhs215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1891_v18), 0))
				_ = _lhs215
				_lhs210.F7 = _rhs287
				_1886_v13 = _rhs288
				(_lhs211).ArraySet1(_rhs289, (_lhs212).Int())
				_lhs213.F13 = _rhs290
				(_lhs214).ArraySet1(_rhs291, (_lhs215).Int())
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1890_v17), 0))
				_ = _index253
				(_1890_v17).ArraySet1((_dafny.Zero).Minus(_1869_cf46), (_index253).Int())
				var _1894_v21 *C1
				_ = _1894_v21
				var _nw311 *C1 = New_C1_()
				_ = _nw311
				_nw311.Ctor__((_this).F24(), (_this).F24())
				_1894_v21 = _nw311
				var _1895_v22 *C1
				_ = _1895_v22
				_1895_v22 = _1894_v21
				_1895_v22 = _1895_v22
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_1871_v0), 0))
				_ = _index254
				(_1871_v0).ArraySet1(((_this).F24()) && ((_this).F24()), (_index254).Int())
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_1871_v0), 0))
				_ = _index255
				(_1871_v0).ArraySet1((_this).F24(), (_index255).Int())
			} else {
				var _1896_v23 _dafny.Map
				_ = _1896_v23
				_1896_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wjgyr")).Cardinality()))
				var _1897_v24 _dafny.Set
				_ = _1897_v24
				_1897_v24 = _dafny.SetOf((_this).F24())
				var _1898_v25 _dafny.Map
				_ = _1898_v25
				_1898_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1897_v24).Cardinality(), _dafny.IntOfInt64(151))
				_1896_v23 = (_1896_v23).Update((_this).F24(), (func() _dafny.Int {
					if (_1898_v25).Contains(_1868_cf47) {
						return (_1898_v25).Get(_1868_cf47).(_dafny.Int)
					}
					return _1868_cf47
				})())
				var _1899_v26 _dafny.Set
				_ = _1899_v26
				_1899_v26 = _dafny.SetOf((_dafny.Zero).Minus(_1867_cf48))
				var _1900_v27 _dafny.Set
				_ = _1900_v27
				_1900_v27 = _dafny.SetOf(_1868_cf47, (_1899_v26).Cardinality())
				var _1901_v28 _dafny.Sequence
				_ = _1901_v28
				_1901_v28 = _dafny.SeqOf(_1870_cf45)
				var _1902_v29 _dafny.Set
				_ = _1902_v29
				_1902_v29 = _dafny.SetOf(_1900_v27)
				var _1903_v30 _dafny.Map
				_ = _1903_v30
				_1903_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1901_v28, (_1902_v29).Cardinality())
				var _1904_v31 _dafny.Map
				_ = _1904_v31
				_1904_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1874_v1, (Companion_Default___.SafeIndex((_1899_v26).Cardinality(), _dafny.IntOfUint32((_1874_v1).Cardinality()))).Uint32(), _this.F23()), _1903_v30)
				_1904_v31 = (_1904_v31).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uso"), _1874_v1), (func() _dafny.Map {
					var _coll90 = _dafny.NewMapBuilder()
					_ = _coll90
					for _iter95 := _dafny.Iterate((_dafny.SeqOf(_1901_v28, _1901_v28, _1901_v28)).Elements()); ; {
						_compr_90, _ok95 := _iter95()
						if !_ok95 {
							break
						}
						var _1905_v32 _dafny.Sequence
						_1905_v32 = interface{}(_compr_90).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1901_v28, _1901_v28, _1901_v28), _1905_v32) {
							_coll90.Add(_1905_v32, _dafny.IntOfInt64(66))
						}
					}
					return _coll90.ToMap()
				}()).Merge(_1903_v30))
				var _1906_v33 _dafny.MultiSet
				_ = _1906_v33
				_1906_v33 = _dafny.MultiSetOf(_1870_cf45, false)
				var _1907_v34 _dafny.Sequence
				_ = _1907_v34
				_1907_v34 = _dafny.SeqOf(_1876_v3)
				var _1908_v35 *C2
				_ = _1908_v35
				var _nw312 *C2 = New_C2_()
				_ = _nw312
				_nw312.Ctor__(_1906_v33, (_1907_v34).Select((Companion_Default___.SafeIndex(_1867_cf48, _dafny.IntOfUint32((_1907_v34).Cardinality()))).Uint32()).(_dafny.Array), _this.F23(), _1870_cf45)
				_1908_v35 = _nw312
				_1870_cf45 = (_this).F24()
				(globalState).F13 = (_this).F24()
			}
			var _1909_v36 _dafny.Map
			_ = _1909_v36
			_1909_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1871_v0)
			(globalState).F7 = (_1909_v36).Cardinality()
		} else if _source27.Is_DC29() {
			var _1910___mcc_h4 bool = _source27.Get_().(D13_DC29).Cf49
			_ = _1910___mcc_h4
			var _1911___mcc_h5 bool = _source27.Get_().(D13_DC29).Cf50
			_ = _1911___mcc_h5
			var _1912_cf50 bool = _1911___mcc_h5
			_ = _1912_cf50
			var _1913_cf49 bool = _1910___mcc_h4
			_ = _1913_cf49
			var _1914_v37 _dafny.Int
			_ = _1914_v37
			_1914_v37 = _dafny.IntOfInt64(-698)
			var _1915_v38 _dafny.Array
			_ = _1915_v38
			var _nwElement0_58 _dafny.Int = _1914_v37
			_ = _nwElement0_58
			var _nw313 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(7))
			_ = _nw313
			(_nw313).ArraySet1(_nwElement0_58, 0)
			(_nw313).ArraySet1(((_dafny.Zero).Minus(_1914_v37)).Times(_dafny.IntOfInt64(944)), 1)
			(_nw313).ArraySet1(_1914_v37, 2)
			(_nw313).ArraySet1(_1914_v37, 3)
			(_nw313).ArraySet1(_1914_v37, 4)
			(_nw313).ArraySet1(_1914_v37, 5)
			(_nw313).ArraySet1(_1914_v37, 6)
			_1915_v38 = _nw313
			(globalState).F22 = _1915_v38
			var _1916_v39 _dafny.Array
			_ = _1916_v39
			var _len0_56 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_56
			var _nw314 _dafny.Array
			_ = _nw314
			if _len0_56.Cmp(_dafny.Zero) == 0 {
				_nw314 = _dafny.NewArray(_len0_56)
			} else {
				var _init56 func(_dafny.Int) bool = func(_1917_i1 _dafny.Int) bool {
					return false
				}
				_ = _init56
				var _element0_56 = _init56(_dafny.Zero)
				_ = _element0_56
				_nw314 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
				(_nw314).ArraySet1(_element0_56, 0)
				var _nativeLen0_56 = (_len0_56).Int()
				_ = _nativeLen0_56
				for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
					(_nw314).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
				}
			}
			_1916_v39 = _nw314
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1916_v39), 0))
			_ = _index256
			(_1916_v39).ArraySet1((_this).F24(), (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1916_v39), 0))
			_ = _index257
			(_1916_v39).ArraySet1(_1913_cf49, (_index257).Int())
			var _1918_v40 _dafny.Array
			_ = _1918_v40
			var _len0_57 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_57
			var _nw315 _dafny.Array
			_ = _nw315
			if _len0_57.Cmp(_dafny.Zero) == 0 {
				_nw315 = _dafny.NewArray(_len0_57)
			} else {
				var _init57 func(_dafny.Int) _dafny.Map = (func(_1919_cf50 bool) func(_dafny.Int) _dafny.Map {
					return func(_1920_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1919_cf50, _dafny.IntOfInt64(736))
					}
				})(_1912_cf50)
				_ = _init57
				var _element0_57 = _init57(_dafny.Zero)
				_ = _element0_57
				_nw315 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
				(_nw315).ArraySet1(_element0_57, 0)
				var _nativeLen0_57 = (_len0_57).Int()
				_ = _nativeLen0_57
				for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
					(_nw315).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
				}
			}
			_1918_v40 = _nw315
			_1918_v40 = _1918_v40
			(globalState).F13 = (_this).F24()
		} else {
			var _1921___mcc_h6 _dafny.Map = _source27.Get_().(D13_DC27).Cf44
			_ = _1921___mcc_h6
			var _1922_cf44 _dafny.Map = _1921___mcc_h6
			_ = _1922_cf44
			var _1923_v41 _dafny.Sequence
			_ = _1923_v41
			_1923_v41 = _dafny.SeqOf(!((_this).F24()))
			var _1924_v42 _dafny.Set
			_ = _1924_v42
			_1924_v42 = _dafny.SetOf((_this).F24())
			var _1925_v43 _dafny.Map
			_ = _1925_v43
			_1925_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1923_v41, _1924_v42)
			var _1926_v44 _dafny.Int
			_ = _1926_v44
			_1926_v44 = _dafny.IntOfInt64(35)
			var _1927_v45 _dafny.Sequence
			_ = _1927_v45
			_1927_v45 = _dafny.UnicodeSeqOfUtf8Bytes("tghfl")
			var _rhs292 _dafny.Map = _1925_v43
			_ = _rhs292
			var _rhs293 _dafny.Int = _1926_v44
			_ = _rhs293
			var _rhs294 bool = _dafny.Companion_Sequence_.Contains(_1927_v45, _this.F23())
			_ = _rhs294
			var _rhs295 bool = !((_this).F24())
			_ = _rhs295
			var _lhs216 *GlobalState = globalState
			_ = _lhs216
			var _lhs217 *GlobalState = globalState
			_ = _lhs217
			_1925_v43 = _rhs292
			r1 = _rhs293
			_lhs216.F13 = _rhs294
			_lhs217.F13 = _rhs295
			var _1928_v46 T1
			_ = _1928_v46
			var _nw316 *C6 = New_C6_()
			_ = _nw316
			_nw316.Ctor__(_this.F23(), (_this).F24())
			_1928_v46 = _nw316
			var _1929_v47 _dafny.Sequence
			_ = _1929_v47
			_1929_v47 = _dafny.SeqOf(_1928_v46, _1928_v46)
			_1929_v47 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1929_v47, _dafny.SeqOf(_1928_v46)), _1929_v47)
			if ((_1928_v46).F24()) || ((_this).F24()) {
				(globalState).F7 = _1926_v44
				var _1930_v48 _dafny.Array
				_ = _1930_v48
				var _nw317 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw317
				_1930_v48 = _nw317
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1930_v48), 0))
				_ = _index258
				(_1930_v48).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1923_v41).Cardinality())), (_index258).Int())
				var _1931_v49 _dafny.Map
				_ = _1931_v49
				_1931_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1926_v44, (_this).F24())
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1930_v48), 0))
				_ = _index259
				(_1930_v48).ArraySet1(((_1931_v49).Merge((func() _dafny.Map {
					if (_this).F24() {
						return _1931_v49
					}
					return _1931_v49
				})())).Cardinality(), (_index259).Int())
				var _1932_v50 _dafny.MultiSet
				_ = _1932_v50
				_1932_v50 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1928_v46).F24(), (_1928_v46).F24()))
				var _1933_v51 _dafny.MultiSet
				_ = _1933_v51
				_1933_v51 = _dafny.MultiSetOf((_this).F24())
				var _1934_v52 _dafny.Set
				_ = _1934_v52
				_1934_v52 = _dafny.SetOf(_1926_v44)
				var _1935_v53 _dafny.Int
				_ = _1935_v53
				var _1936_v54 bool
				_ = _1936_v54
				var _1937_v55 _dafny.Int
				_ = _1937_v55
				var _out52 _dafny.Int
				_ = _out52
				var _out53 bool
				_ = _out53
				var _out54 _dafny.Int
				_ = _out54
				_out52, _out53, _out54 = Companion_Default___.M0((_1932_v50).Cardinality(), (_1933_v51).Update(true, Companion_Default___.Abs((_1934_v52).Cardinality())), (_1928_v46).F24(), _1926_v44, globalState)
				_1935_v53 = _out52
				_1936_v54 = _out53
				_1937_v55 = _out54
				var _1938_v56 _dafny.Array
				_ = _1938_v56
				var _nw318 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw318
				_1938_v56 = _nw318
				var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
				_ = _nw319
				_1938_v56 = _nw319
				_1923_v41 = _1923_v41
			} else {
				_1926_v44 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(726), _1926_v44)
				var _1939_v57 _dafny.Map
				_ = _1939_v57
				_1939_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1928_v46).F24(), _1923_v41)
				_1939_v57 = (_1939_v57).Update((_1928_v46).F24(), _1923_v41)
				(globalState).F13 = (_1928_v46).F24()
				var _1940_v58 _dafny.Map
				_ = _1940_v58
				_1940_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1926_v44, (_1928_v46).F24())
				var _1941_v59 _dafny.Sequence
				_ = _1941_v59
				_1941_v59 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_1926_v44, _1926_v44), (_1940_v58).Cardinality())
				var _rhs296 bool = (_1926_v44).Cmp((_dafny.Zero).Minus(_1926_v44)) <= 0
				_ = _rhs296
				var _rhs297 _dafny.Int = (_1926_v44).Times((_dafny.IntOfUint32((_1923_v41).Cardinality())).Minus(_1926_v44))
				_ = _rhs297
				var _rhs298 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1926_v44, _1926_v44), _1941_v59)
				_ = _rhs298
				var _rhs299 bool = (_this).F24()
				_ = _rhs299
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				r2 = _rhs296
				_lhs218.F7 = _rhs297
				_1941_v59 = _rhs298
				_lhs219.F13 = _rhs299
				var _1942_v60 D14
				_ = _1942_v60
				_1942_v60 = Companion_D14_.Create_DC31_((_1928_v46).F24())
				r2 = (_1942_v60).Dtor_cf52()
			}
			var _1943_v61 _dafny.Array
			_ = _1943_v61
			var _nw320 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw320
			_1943_v61 = _nw320
			var _1944_v62 _dafny.Map
			_ = _1944_v62
			_1944_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1943_v61, _1928_v46.F23())
			var _1945_v63 _dafny.Array
			_ = _1945_v63
			var _len0_58 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_58
			var _nw321 _dafny.Array
			_ = _nw321
			if _len0_58.Cmp(_dafny.Zero) == 0 {
				_nw321 = _dafny.NewArray(_len0_58)
			} else {
				var _init58 func(_dafny.Int) _dafny.CodePoint = func(_1946_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}
				_ = _init58
				var _element0_58 = _init58(_dafny.Zero)
				_ = _element0_58
				_nw321 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
				(_nw321).ArraySet1CodePoint(_element0_58, 0)
				var _nativeLen0_58 = (_len0_58).Int()
				_ = _nativeLen0_58
				for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
					(_nw321).ArraySet1CodePoint(_init58(_dafny.IntOf(_i0_58)), _i0_58)
				}
			}
			_1945_v63 = _nw321
			var _1947_v65 *C3
			_ = _1947_v65
			var _nw322 *C3 = New_C3_()
			_ = _nw322
			_nw322.Ctor__(_1944_v62, _1945_v63, (_1927_v45).Select((Companion_Default___.SafeIndex(_1926_v44, _dafny.IntOfUint32((_1927_v45).Cardinality()))).Uint32()).(_dafny.CodePoint), (_dafny.SetOf(_1926_v44)).Equals(func() _dafny.Set {
				var _coll91 = _dafny.NewBuilder()
				_ = _coll91
				for _iter96 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(164), _dafny.IntOfInt64(617))); ; {
					_compr_91, _ok96 := _iter96()
					if !_ok96 {
						break
					}
					var _1948_v64 _dafny.Int
					_1948_v64 = interface{}(_compr_91).(_dafny.Int)
					if ((_dafny.IntOfInt64(164)).Cmp(_1948_v64) <= 0) && ((_1948_v64).Cmp(_dafny.IntOfInt64(617)) < 0) {
						_coll91.Add((_1948_v64).Minus(_1926_v44))
					}
				}
				return _coll91.ToSet()
			}()))
			_1947_v65 = _nw322
		}
		var _1949_v66 _dafny.Int
		_ = _1949_v66
		_1949_v66 = _dafny.IntOfInt64(-734)
		var _1950_v67 _dafny.Sequence
		_ = _1950_v67
		_1950_v67 = _dafny.UnicodeSeqOfUtf8Bytes("yfjo")
		var _1951_v68 _dafny.Set
		_ = _1951_v68
		_1951_v68 = _dafny.SetOf((_this).Fm6(globalState), _1949_v66, _1949_v66)
		var _1952_v69 _dafny.Map
		_ = _1952_v69
		_1952_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1949_v66, _this.F23())
		var _1953_v70 _dafny.Array
		_ = _1953_v70
		var _nw323 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw323
		_1953_v70 = _nw323
		var _1954_v71 _dafny.Sequence
		_ = _1954_v71
		_1954_v71 = _dafny.SeqOf(_1953_v70)
		var _1955_v72 _dafny.Sequence
		_ = _1955_v72
		_1955_v72 = _dafny.SeqOf(_1952_v69, _1952_v69, _1952_v69, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1954_v71).Cardinality()), _this.F23()))
		var _1956_v73 _dafny.MultiSet
		_ = _1956_v73
		_1956_v73 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1949_v66), _1949_v66)
		var _1957_v74 _dafny.Array
		_ = _1957_v74
		var _nwElement0_59 _dafny.Int = _1949_v66
		_ = _nwElement0_59
		var _nw324 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(21))
		_ = _nw324
		(_nw324).ArraySet1(_nwElement0_59, 0)
		(_nw324).ArraySet1(_1949_v66, 1)
		(_nw324).ArraySet1((_1949_v66).Times(_dafny.IntOfUint32((_1950_v67).Cardinality())), 2)
		(_nw324).ArraySet1(_1949_v66, 3)
		(_nw324).ArraySet1((_1949_v66).Minus(_dafny.IntOfInt64(434)), 4)
		(_nw324).ArraySet1(_1949_v66, 5)
		(_nw324).ArraySet1(_1949_v66, 6)
		(_nw324).ArraySet1(_dafny.IntOfUint32((_1950_v67).Cardinality()), 7)
		(_nw324).ArraySet1((_1951_v68).Cardinality(), 8)
		(_nw324).ArraySet1(_1949_v66, 9)
		(_nw324).ArraySet1((_1949_v66).Minus(_1949_v66), 10)
		(_nw324).ArraySet1(_1949_v66, 11)
		(_nw324).ArraySet1((_1949_v66).Minus(_dafny.IntOfUint32((_1955_v72).Cardinality())), 12)
		(_nw324).ArraySet1(_1949_v66, 13)
		(_nw324).ArraySet1(_1949_v66, 14)
		(_nw324).ArraySet1(_1949_v66, 15)
		(_nw324).ArraySet1((_1949_v66).Plus(_1949_v66), 16)
		(_nw324).ArraySet1(_1949_v66, 17)
		(_nw324).ArraySet1(Companion_Default___.SafeModuloInt((_1956_v73).Cardinality(), _1949_v66), 18)
		(_nw324).ArraySet1(_1949_v66, 19)
		(_nw324).ArraySet1(_dafny.IntOfInt64(139), 20)
		_1957_v74 = _nw324
		var _1958_v75 _dafny.Map
		_ = _1958_v75
		_1958_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1949_v66, _dafny.IntOfInt64(904))
		var _1959_v76 _dafny.Map
		_ = _1959_v76
		_1959_v76 = _1958_v75
		var _1960_v77 _dafny.Array
		_ = _1960_v77
		var _len0_59 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_59
		var _nw325 _dafny.Array
		_ = _nw325
		if _len0_59.Cmp(_dafny.Zero) == 0 {
			_nw325 = _dafny.NewArray(_len0_59)
		} else {
			var _init59 func(_dafny.Int) _dafny.CodePoint = func(_1961_i4 _dafny.Int) _dafny.CodePoint {
				return _this.F23()
			}
			_ = _init59
			var _element0_59 = _init59(_dafny.Zero)
			_ = _element0_59
			_nw325 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
			(_nw325).ArraySet1CodePoint(_element0_59, 0)
			var _nativeLen0_59 = (_len0_59).Int()
			_ = _nativeLen0_59
			for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
				(_nw325).ArraySet1CodePoint(_init59(_dafny.IntOf(_i0_59)), _i0_59)
			}
		}
		_1960_v77 = _nw325
		var _1962_v78 T2
		_ = _1962_v78
		var _nw326 *C4 = New_C4_()
		_ = _nw326
		_nw326.Ctor__(_dafny.IntOfInt64(872), _1959_v76, _1960_v77, _this.F23(), !((_this).F24()))
		_1962_v78 = _nw326
		var _1963_v79 _dafny.Sequence
		_ = _1963_v79
		_1963_v79 = _dafny.SeqOf(_1962_v78, _1962_v78, _1962_v78, _1962_v78, _1962_v78)
		var _1964_v80 _dafny.Sequence
		_ = _1964_v80
		_1964_v80 = _dafny.SeqOf((_1963_v79).Select((Companion_Default___.SafeIndex(_1949_v66, _dafny.IntOfUint32((_1963_v79).Cardinality()))).Uint32()).(T2), _1962_v78)
		var _1965_v81 _dafny.Map
		_ = _1965_v81
		_1965_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1964_v80)
		var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1957_v74), 0))
		_ = _index260
		(_1957_v74).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1965_v81).Contains((_1962_v78).F24()) {
				return (_1965_v81).Get((_1962_v78).F24()).(_dafny.Sequence)
			}
			return _1964_v80
		})()).Cardinality()), (_index260).Int())
		var _1966_v82 _dafny.Map
		_ = _1966_v82
		_1966_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1962_v78).F24(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F24(), (_this).F24())).Cardinality()))
		var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1957_v74), 0))
		_ = _index261
		(_1957_v74).ArraySet1(Companion_Default___.SafeDivisionInt(_1949_v66, (func() _dafny.Int {
			if (_1966_v82).Contains((_1962_v78).F24()) {
				return (_1966_v82).Get((_1962_v78).F24()).(_dafny.Int)
			}
			return _dafny.IntOfInt64(-751)
		})()), (_index261).Int())
		r1 = (_1957_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1957_v74), 0))).Int()).(_dafny.Int)
		var _1967_v83 _dafny.Map
		_ = _1967_v83
		_1967_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1962_v78).F25(), ((_1957_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1957_v74), 0))).Int()).(_dafny.Int)).Minus((_1957_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1957_v74), 0))).Int()).(_dafny.Int)))
		_1967_v83 = (_1967_v83).Update(_1960_v77, _1949_v66)
		var _rhs300 _dafny.CodePoint = _1962_v78.F23()
		_ = _rhs300
		var _rhs301 _dafny.CodePoint = Companion_Default___.Fm34(Companion_Default___.SafeModuloInt((_1951_v68).Cardinality(), _1949_v66), (_1962_v78).F24(), globalState)
		_ = _rhs301
		var _lhs220 T2 = _1962_v78
		_ = _lhs220
		var _lhs221 *C9 = _this
		_ = _lhs221
		_lhs220.F23_set_(_rhs300)
		_lhs221.F23_set_(_rhs301)
		var _1968_v84 _dafny.Array
		_ = _1968_v84
		var _len0_60 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_60
		var _nw327 _dafny.Array
		_ = _nw327
		if _len0_60.Cmp(_dafny.Zero) == 0 {
			_nw327 = _dafny.NewArray(_len0_60)
		} else {
			var _init60 func(_dafny.Int) _dafny.Set = (func(_1969_v68 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_1970_i5 _dafny.Int) _dafny.Set {
					return _1969_v68
				}
			})(_1951_v68)
			_ = _init60
			var _element0_60 = _init60(_dafny.Zero)
			_ = _element0_60
			_nw327 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
			(_nw327).ArraySet1(_element0_60, 0)
			var _nativeLen0_60 = (_len0_60).Int()
			_ = _nativeLen0_60
			for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
				(_nw327).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
			}
		}
		_1968_v84 = _nw327
		var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_1968_v84), 0))
		_ = _index262
		(_1968_v84).ArraySet1((Companion_D5_.Create_DC12_(_1951_v68)).Dtor_cf19(), (_index262).Int())
		var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_1968_v84), 0))
		_ = _index263
		(_1968_v84).ArraySet1(_1951_v68, (_index263).Int())
		r0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1950_v67, _dafny.UnicodeSeqOfUtf8Bytes("ebhil")), _1950_v67)
		r1 = _1949_v66
		r2 = (Companion_D15_.Create_DC34_(true)).Dtor_cf55()
		return r0, r1, r2
	}
}
func (_this *C9) M5(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1971_v0 _dafny.Sequence
		_ = _1971_v0
		_1971_v0 = _dafny.UnicodeSeqOfUtf8Bytes("bnuau")
		var _hi12 _dafny.Int = Companion_Default___.SafeDivisionInt(p3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1971_v0).Cardinality())))
		_ = _hi12
		for _1972_i0 := p3; _1972_i0.Cmp(_hi12) < 0; _1972_i0 = _1972_i0.Plus(_dafny.One) {
			if false {
				_1971_v0 = _1971_v0
				(globalState).F13 = p2
				(globalState).F7 = (func() _dafny.Int {
					if p2 {
						return p3
					}
					return (func() _dafny.Int {
						if (_this).F24() {
							return _1972_i0
						}
						return _1972_i0
					})()
				})()
				var _1973_v1 _dafny.Array
				_ = _1973_v1
				var _len0_61 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_61
				var _nw328 _dafny.Array
				_ = _nw328
				if _len0_61.Cmp(_dafny.Zero) == 0 {
					_nw328 = _dafny.NewArray(_len0_61)
				} else {
					var _init61 func(_dafny.Int) _dafny.Int = func(_1974_i1 _dafny.Int) _dafny.Int {
						return (_1974_i1).Plus(_dafny.IntOfInt64(930))
					}
					_ = _init61
					var _element0_61 = _init61(_dafny.Zero)
					_ = _element0_61
					_nw328 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
					(_nw328).ArraySet1(_element0_61, 0)
					var _nativeLen0_61 = (_len0_61).Int()
					_ = _nativeLen0_61
					for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
						(_nw328).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
					}
				}
				_1973_v1 = _nw328
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1973_v1), 0))
				_ = _index264
				(_1973_v1).ArraySet1(_1972_i0, (_index264).Int())
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1973_v1), 0))
				_ = _index265
				(_1973_v1).ArraySet1(Companion_Default___.SafeDivisionInt(p3, p3), (_index265).Int())
				var _1975_v2 _dafny.Sequence
				_ = _1975_v2
				_1975_v2 = _dafny.SeqOf(p1)
				var _1976_v3 *C1
				_ = _1976_v3
				var _nw329 *C1 = New_C1_()
				_ = _nw329
				_nw329.Ctor__((_this).Fm4(_dafny.IntOfInt64(475), _1975_v2, globalState), p2)
				_1976_v3 = _nw329
			} else {
				var _rhs302 bool = (_this).F24()
				_ = _rhs302
				var _rhs303 bool = !(p2) || ((func() bool {
					if p2 {
						return p1
					}
					return true
				})())
				_ = _rhs303
				var _lhs222 *GlobalState = globalState
				_ = _lhs222
				var _lhs223 *GlobalState = globalState
				_ = _lhs223
				_lhs222.F13 = _rhs302
				_lhs223.F13 = _rhs303
				var _1977_v4 _dafny.Set
				_ = _1977_v4
				_1977_v4 = _dafny.SetOf((_dafny.Zero).Minus(_1972_i0), p3)
				_1977_v4 = _1977_v4
				var _1978_v5 _dafny.MultiSet
				_ = _1978_v5
				_1978_v5 = _dafny.MultiSetOf(_1972_i0, Companion_Default___.Fm0((_this).F24(), p3, globalState))
				var _1979_v6 _dafny.Sequence
				_ = _1979_v6
				_1979_v6 = _dafny.SeqOf(p2, (_this).F24())
				var _1980_v7 _dafny.Map
				_ = _1980_v7
				_1980_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1971_v0, (_this).F24())
				var _1981_v8 _dafny.Map
				_ = _1981_v8
				_1981_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1978_v5).IsDisjointFrom((_1978_v5).Update(_dafny.IntOfUint32((_1979_v6).Cardinality()), Companion_Default___.Abs(p3))), (_1980_v7).Cardinality())
				_1981_v8 = _1981_v8
				(globalState).F13 = false
				var _1982_v9 _dafny.Array
				_ = _1982_v9
				var _len0_62 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_62
				var _nw330 _dafny.Array
				_ = _nw330
				if _len0_62.Cmp(_dafny.Zero) == 0 {
					_nw330 = _dafny.NewArray(_len0_62)
				} else {
					var _init62 func(_dafny.Int) _dafny.Int = (func(_1983_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1984_i2 _dafny.Int) _dafny.Int {
							return (_1984_i2).Times(_dafny.IntOfUint32((_1983_v6).Cardinality()))
						}
					})(_1979_v6)
					_ = _init62
					var _element0_62 = _init62(_dafny.Zero)
					_ = _element0_62
					_nw330 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
					(_nw330).ArraySet1(_element0_62, 0)
					var _nativeLen0_62 = (_len0_62).Int()
					_ = _nativeLen0_62
					for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
						(_nw330).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
					}
				}
				_1982_v9 = _nw330
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1982_v9), 0))
				_ = _index266
				(_1982_v9).ArraySet1(p3, (_index266).Int())
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1982_v9), 0))
				_ = _index267
				(_1982_v9).ArraySet1(Companion_Default___.SafeModuloInt(_1972_i0, p3), (_index267).Int())
			}
			var _1985_v10 _dafny.Map
			_ = _1985_v10
			_1985_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(445))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}(func(_1986_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(635)
			}))).Cardinality()))
			var _1987_v11 _dafny.MultiSet
			_ = _1987_v11
			_1987_v11 = _dafny.MultiSetOf(_1972_i0, (_dafny.Zero).Minus((func() _dafny.Int {
				if (_1985_v10).Contains(p2) {
					return (_1985_v10).Get(p2).(_dafny.Int)
				}
				return _1972_i0
			})()), _1972_i0)
			var _1988_v12 _dafny.MultiSet
			_ = _1988_v12
			_1988_v12 = _dafny.MultiSetOf(_1971_v0)
			(globalState).F13 = !((_1987_v11).IsSubsetOf(_1987_v11)) || ((_dafny.MultiSetOf(_1971_v0)).IsSubsetOf(_1988_v12))
			var _1989_v13 _dafny.Array
			_ = _1989_v13
			var _nw331 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw331
			_1989_v13 = _nw331
			(globalState).F22 = _1989_v13
			var _1990_v14 _dafny.Array
			_ = _1990_v14
			var _nw332 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw332
			_1990_v14 = _nw332
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_1990_v14), 0))
			_ = _index268
			(_1990_v14).ArraySet1(_1989_v13, (_index268).Int())
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_1990_v14), 0))
			_ = _index269
			(_1990_v14).ArraySet1((func() _dafny.Array {
				if p1 {
					return _1989_v13
				}
				return _1989_v13
			})(), (_index269).Int())
		}
		var _1991_v15 _dafny.Array
		_ = _1991_v15
		var _len0_63 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_63
		var _nw333 _dafny.Array
		_ = _nw333
		if _len0_63.Cmp(_dafny.Zero) == 0 {
			_nw333 = _dafny.NewArray(_len0_63)
		} else {
			var _init63 func(_dafny.Int) bool = (func(_1992_p1 bool) func(_dafny.Int) bool {
				return func(_1993_i4 _dafny.Int) bool {
					return _1992_p1
				}
			})(p1)
			_ = _init63
			var _element0_63 = _init63(_dafny.Zero)
			_ = _element0_63
			_nw333 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
			(_nw333).ArraySet1(_element0_63, 0)
			var _nativeLen0_63 = (_len0_63).Int()
			_ = _nativeLen0_63
			for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
				(_nw333).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
			}
		}
		_1991_v15 = _nw333
		var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
		_ = _index270
		(_1991_v15).ArraySet1(p2, (_index270).Int())
		var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
		_ = _index271
		(_1991_v15).ArraySet1((_this).F24(), (_index271).Int())
		var _1994_v16 _dafny.Array
		_ = _1994_v16
		var _nwElement0_60 _dafny.CodePoint = _this.F23()
		_ = _nwElement0_60
		var _nw334 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(5))
		_ = _nw334
		(_nw334).ArraySet1CodePoint(_nwElement0_60, 0)
		(_nw334).ArraySet1CodePoint((_1971_v0).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1971_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
		(_nw334).ArraySet1CodePoint(_this.F23(), 2)
		(_nw334).ArraySet1CodePoint(Companion_Default___.Fm34(p3, (_this).F24(), globalState), 3)
		(_nw334).ArraySet1CodePoint(_this.F23(), 4)
		_1994_v16 = _nw334
		var _source28 D12 = Companion_D12_.Create_DC24_(_1994_v16)
		_ = _source28
		if _source28.Is_DC25() {
			var _1995___mcc_h0 bool = _source28.Get_().(D12_DC25).Cf41
			_ = _1995___mcc_h0
			var _1996___mcc_h1 _dafny.CodePoint = _source28.Get_().(D12_DC25).Cf42
			_ = _1996___mcc_h1
			var _1997_cf42 _dafny.CodePoint = _1996___mcc_h1
			_ = _1997_cf42
			var _1998_cf41 bool = _1995___mcc_h0
			_ = _1998_cf41
			(globalState).F13 = p1
			r0 = _1998_cf41
			var _1999_v17 _dafny.Set
			_ = _1999_v17
			_1999_v17 = _dafny.SetOf(p1)
			if (_dafny.SetOf(_1998_cf41)).Equals(_1999_v17) {
				var _2000_v18 _dafny.Array
				_ = _2000_v18
				var _nw335 _dafny.Array = _dafny.NewArrayWithValue(Companion_D18_.Default(), _dafny.IntOfInt64(5))
				_ = _nw335
				_2000_v18 = _nw335
				var _2001_v19 T0
				_ = _2001_v19
				var _nw336 *C0 = New_C0_()
				_ = _nw336
				_nw336.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1971_v0), _1998_cf41, _this.F23(), _1998_cf41)
				_2001_v19 = _nw336
				var _2002_v20 _dafny.Map
				_ = _2002_v20
				_2002_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2001_v19)
				var _2003_v21 _dafny.Map
				_ = _2003_v21
				_2003_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2001_v19)
				var _2004_v22 _dafny.Map
				_ = _2004_v22
				_2004_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2001_v19.F23(), _2001_v19)
				var _2005_v23 _dafny.Array
				_ = _2005_v23
				var _nwElement0_61 T0 = _2001_v19
				_ = _nwElement0_61
				var _nw337 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(13))
				_ = _nw337
				(_nw337).ArraySet1(_nwElement0_61, 0)
				(_nw337).ArraySet1((func() T0 {
					if (_2002_v20).Contains((_this).F24()) {
						return (_2002_v20).Get((_this).F24()).(T0)
					}
					return _2001_v19
				})(), 1)
				(_nw337).ArraySet1((func() T0 {
					if (_2003_v21).Contains(p3) {
						return (_2003_v21).Get(p3).(T0)
					}
					return _2001_v19
				})(), 2)
				(_nw337).ArraySet1(_2001_v19, 3)
				(_nw337).ArraySet1(_2001_v19, 4)
				(_nw337).ArraySet1(_2001_v19, 5)
				(_nw337).ArraySet1(_2001_v19, 6)
				(_nw337).ArraySet1(_2001_v19, 7)
				(_nw337).ArraySet1((func() T0 {
					if (_2004_v22).Contains(_2001_v19.F23()) {
						return (_2004_v22).Get(_2001_v19.F23()).(T0)
					}
					return _2001_v19
				})(), 8)
				(_nw337).ArraySet1(_2001_v19, 9)
				(_nw337).ArraySet1(_2001_v19, 10)
				(_nw337).ArraySet1(_2001_v19, 11)
				(_nw337).ArraySet1(_2001_v19, 12)
				_2005_v23 = _nw337
				var _pat_let_tv33 = _2005_v23
				_ = _pat_let_tv33
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_2000_v18), 0))
				_ = _index272
				(_2000_v18).ArraySet1(func(_pat_let49_0 D18) D18 {
					return func(_2006_dt__update__tmp_h0 D18) D18 {
						return func(_pat_let50_0 _dafny.Array) D18 {
							return func(_2007_dt__update_hcf67_h0 _dafny.Array) D18 {
								return Companion_D18_.Create_DC41_(_2007_dt__update_hcf67_h0)
							}(_pat_let50_0)
						}(_pat_let_tv33)
					}(_pat_let49_0)
				}(Companion_D18_.Create_DC41_(_2005_v23)), (_index272).Int())
				var _2008_v24 D18
				_ = _2008_v24
				_2008_v24 = Companion_D18_.Create_DC41_(_2005_v23)
				var _pat_let_tv34 = _2005_v23
				_ = _pat_let_tv34
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_2000_v18), 0))
				_ = _index273
				(_2000_v18).ArraySet1(func(_pat_let51_0 D18) D18 {
					return func(_2009_dt__update__tmp_h1 D18) D18 {
						return func(_pat_let52_0 _dafny.Array) D18 {
							return func(_2010_dt__update_hcf67_h1 _dafny.Array) D18 {
								return Companion_D18_.Create_DC41_(_2010_dt__update_hcf67_h1)
							}(_pat_let52_0)
						}(_pat_let_tv34)
					}(_pat_let51_0)
				}(_2008_v24), (_index273).Int())
				var _2011_v25 _dafny.Map
				_ = _2011_v25
				_2011_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1998_cf41)
				var _2012_v26 _dafny.Map
				_ = _2012_v26
				_2012_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2011_v25).Cardinality(), p1)
				var _2013_v27 *C8
				_ = _2013_v27
				var _nw338 *C8 = New_C8_()
				_ = _nw338
				_nw338.Ctor__((_2011_v25).Equals(_2011_v25))
				_2013_v27 = _nw338
				_1971_v0 = (func() _dafny.Sequence {
					if !(_1998_cf41) {
						return _1971_v0
					}
					return _1971_v0
				})()
				(globalState).F13 = (_this).F24()
				var _2014_v28 _dafny.MultiSet
				_ = _2014_v28
				_2014_v28 = _dafny.MultiSetOf((_2001_v19).F24(), false, (p3).Cmp(p3) < 0, (_2001_v19).F24())
				_2014_v28 = _dafny.MultiSetOf((func() _dafny.Set {
					var _coll92 = _dafny.NewBuilder()
					_ = _coll92
					for _iter97 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(767), _dafny.IntOfInt64(828))); ; {
						_compr_92, _ok97 := _iter97()
						if !_ok97 {
							break
						}
						var _2015_v29 _dafny.Int
						_2015_v29 = interface{}(_compr_92).(_dafny.Int)
						if ((_dafny.IntOfInt64(767)).Cmp(_2015_v29) <= 0) && ((_2015_v29).Cmp(_dafny.IntOfInt64(828)) < 0) {
							_coll92.Add(Companion_Default___.SafeDivisionInt(_2015_v29, p3))
						}
					}
					return _coll92.ToSet()
				}()).IsDisjointFrom(func() _dafny.Set {
					var _coll93 = _dafny.NewBuilder()
					_ = _coll93
					for _iter98 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(529), _dafny.IntOfInt64(499))); ; {
						_compr_93, _ok98 := _iter98()
						if !_ok98 {
							break
						}
						var _2016_v30 _dafny.Int
						_2016_v30 = interface{}(_compr_93).(_dafny.Int)
						if ((_dafny.IntOfInt64(529)).Cmp(_2016_v30) <= 0) && ((_2016_v30).Cmp(_dafny.IntOfInt64(499)) < 0) {
							_coll93.Add((_2016_v30).Minus(p3))
						}
					}
					return _coll93.ToSet()
				}()))
			} else {
				var _2017_v31 _dafny.Map
				_ = _2017_v31
				_2017_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
				var _2018_v32 *C4
				_ = _2018_v32
				var _nw339 *C4 = New_C4_()
				_ = _nw339
				_nw339.Ctor__(p3, _2017_v31, _1994_v16, _this.F23(), !(p1))
				_2018_v32 = _nw339
				var _2019_v33 _dafny.Array
				_ = _2019_v33
				var _nw340 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(10))
				_ = _nw340
				_2019_v33 = _nw340
				var _2020_v34 *C7
				_ = _2020_v34
				var _nw341 *C7 = New_C7_()
				_ = _nw341
				_nw341.Ctor__((_this).F24(), (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), _1994_v16, _1997_cf42, p1, p1)
				_2020_v34 = _nw341
				var _2021_v35 _dafny.MultiSet
				_ = _2021_v35
				_2021_v35 = _dafny.MultiSetOf(_2020_v34)
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_2019_v33), 0))
				_ = _index274
				(_2019_v33).ArraySet1(_2021_v35, (_index274).Int())
				var _2022_v36 _dafny.Sequence
				_ = _2022_v36
				_2022_v36 = _dafny.SeqOf(p3)
				var _2023_v37 _dafny.MultiSet
				_ = _2023_v37
				_2023_v37 = _dafny.MultiSetOf(_dafny.IntOfInt64(905), _dafny.IntOfUint32((_dafny.SeqOf(_2022_v36)).Cardinality()), _dafny.IntOfUint32((_1971_v0).Cardinality()))
				var _2024_v38 D15
				_ = _2024_v38
				_2024_v38 = Companion_D15_.Create_DC34_(p1)
				var _2025_v39 _dafny.Sequence
				_ = _2025_v39
				_2025_v39 = _dafny.SeqOf((_2020_v34).F31(), p2)
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_2019_v33), 0))
				_ = _index275
				var _rhs304 _dafny.Set = _1999_v17
				_ = _rhs304
				var _rhs305 _dafny.MultiSet = ((_dafny.MultiSetOf(_2020_v34)).Intersection(_2021_v35)).Difference(_2021_v35)
				_ = _rhs305
				var _rhs306 _dafny.MultiSet = _dafny.MultiSetOf((_2018_v32).F33(), ((_dafny.Zero).Minus(p3)).Plus((_2018_v32).F33()), ((_dafny.MultiSetOf(_2024_v38)).Cardinality()).Minus(_dafny.IntOfUint32((_2025_v39).Cardinality())))
				_ = _rhs306
				var _lhs224 _dafny.Array = _2019_v33
				_ = _lhs224
				var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_2019_v33), 0))
				_ = _lhs225
				_1999_v17 = _rhs304
				(_lhs224).ArraySet1(_rhs305, (_lhs225).Int())
				_2023_v37 = _rhs306
				var _2026_v40 _dafny.Map
				_ = _2026_v40
				_2026_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1998_cf41, (_2020_v34).F31())
				(globalState).F13 = (func() bool {
					if (_2026_v40).Contains(_1998_cf41) {
						return (_2026_v40).Get(_1998_cf41).(bool)
					}
					return _1998_cf41
				})()
				var _2027_v41 _dafny.Set
				_ = _2027_v41
				_2027_v41 = _dafny.SetOf((_2018_v32).F33())
				(globalState).F13 = ((_dafny.SetOf((_2018_v32).F33())).Union(_2027_v41)).IsSubsetOf(_2027_v41)
				var _2028_v42 _dafny.Array
				_ = _2028_v42
				var _len0_64 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_64
				var _nw342 _dafny.Array
				_ = _nw342
				if _len0_64.Cmp(_dafny.Zero) == 0 {
					_nw342 = _dafny.NewArray(_len0_64)
				} else {
					var _init64 func(_dafny.Int) _dafny.Map = (func(_2029_p2 bool) func(_dafny.Int) _dafny.Map {
						return func(_2030_i5 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2029_p2, _this.F23())
						}
					})(p2)
					_ = _init64
					var _element0_64 = _init64(_dafny.Zero)
					_ = _element0_64
					_nw342 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
					(_nw342).ArraySet1(_element0_64, 0)
					var _nativeLen0_64 = (_len0_64).Int()
					_ = _nativeLen0_64
					for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
						(_nw342).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
					}
				}
				_2028_v42 = _nw342
				var _2031_v43 D14
				_ = _2031_v43
				_2031_v43 = Companion_D14_.Create_DC30_(_2028_v42)
				var _2032_v44 _dafny.Map
				_ = _2032_v44
				_2032_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2018_v32).F33(), _2031_v43)
				_2032_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p3).Plus(p3), Companion_D14_.Create_DC30_(_2028_v42))
			}
			var _2033_v45 D12
			_ = _2033_v45
			_2033_v45 = Companion_D12_.Create_DC25_((p3).Cmp(p3) >= 0, _1997_cf42)
			_2033_v45 = _2033_v45
		} else if _source28.Is_DC24() {
			var _2034___mcc_h2 _dafny.Array = _source28.Get_().(D12_DC24).Cf40
			_ = _2034___mcc_h2
			var _2035_cf40 _dafny.Array = _2034___mcc_h2
			_ = _2035_cf40
			var _2036_v46 _dafny.Map
			_ = _2036_v46
			_2036_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm2(globalState))
			var _2037_v47 _dafny.Set
			_ = _2037_v47
			_2037_v47 = _dafny.SetOf((func() bool {
				if (_2036_v46).Contains(p3) {
					return (_2036_v46).Get(p3).(bool)
				}
				return (_this).F24()
			})(), (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
			var _2038_v48 _dafny.Map
			_ = _2038_v48
			_2038_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_2037_v47).Union(_2037_v47))
			_2038_v48 = (_2038_v48).Update(p3, _dafny.SetOf(p1, false))
			if p1 {
				var _2039_v49 D12
				_ = _2039_v49
				_2039_v49 = Companion_D12_.Create_DC25_((_this).F24(), _this.F23())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_1994_v16), 0))
				_ = _index276
				(_1994_v16).ArraySet1CodePoint((_2039_v49).Dtor_cf42(), (_index276).Int())
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_1994_v16), 0))
				_ = _index277
				var _rhs307 _dafny.CodePoint = _this.F23()
				_ = _rhs307
				var _rhs308 bool = !((_this).F24()) || (!((_this).F24()))
				_ = _rhs308
				var _rhs309 bool = p2
				_ = _rhs309
				var _lhs226 _dafny.Array = _1994_v16
				_ = _lhs226
				var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_1994_v16), 0))
				_ = _lhs227
				var _lhs228 *GlobalState = globalState
				_ = _lhs228
				var _lhs229 *GlobalState = globalState
				_ = _lhs229
				(_lhs226).ArraySet1CodePoint(_rhs307, (_lhs227).Int())
				_lhs228.F13 = _rhs308
				_lhs229.F13 = _rhs309
				(globalState).F7 = Companion_Default___.SafeModuloInt((p3).Plus((_dafny.Zero).Minus(p3)), (func() _dafny.Int {
					if (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool) {
						return p3
					}
					return p3
				})())
				(globalState).F13 = p2
				var _2040_v50 _dafny.Array
				_ = _2040_v50
				var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw343
				_2040_v50 = _nw343
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_2040_v50), 0))
				_ = _index278
				(_2040_v50).ArraySet1(p3, (_index278).Int())
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_2040_v50), 0))
				_ = _index279
				(_2040_v50).ArraySet1((_dafny.Zero).Minus(p3), (_index279).Int())
				(globalState).F7 = p3
			} else {
				var _2041_v51 _dafny.Map
				_ = _2041_v51
				_2041_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
				(globalState).F7 = Companion_Default___.Fm0((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), (func() _dafny.Int {
					if (_2041_v51).Contains(p2) {
						return (_2041_v51).Get(p2).(_dafny.Int)
					}
					return (_2036_v46).Cardinality()
				})(), globalState)
				var _2042_v52 _dafny.Array
				_ = _2042_v52
				var _nw344 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
				_ = _nw344
				_2042_v52 = _nw344
				var _2043_v53 D15
				_ = _2043_v53
				_2043_v53 = Companion_D15_.Create_DC34_((_this).F24())
				var _2044_v54 _dafny.Map
				_ = _2044_v54
				_2044_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _2043_v53)
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_2042_v52), 0))
				_ = _index280
				(_2042_v52).ArraySet1(_2044_v54, (_index280).Int())
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_2042_v52), 0))
				_ = _index281
				(_2042_v52).ArraySet1(_2044_v54, (_index281).Int())
				(globalState).F7 = _dafny.IntOfInt64(263)
				_1971_v0 = _dafny.Companion_Sequence_.Concatenate(_1971_v0, _dafny.Companion_Sequence_.Update(_1971_v0, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1971_v0).Cardinality()))).Uint32(), _dafny.CodePoint('o')))
				var _2045_v55 _dafny.Sequence
				_ = _2045_v55
				_2045_v55 = _dafny.SeqOf(_dafny.SetOf(!((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool)), p2, p1, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool)))
				(globalState).F13 = !(((p3).Plus(p3)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2045_v55, _dafny.SeqOf(Companion_Default___.Fm41(globalState)))).Cardinality())) < 0)
			}
			(_this).F23_set_(_this.F23())
			(_this).F23_set_(_this.F23())
		} else {
			var _2046___mcc_h3 D12 = _source28.Get_().(D12_DC26).Cf43
			_ = _2046___mcc_h3
			var _2047_cf43 D12 = _2046___mcc_h3
			_ = _2047_cf43
			_1971_v0 = _1971_v0
			var _2048_v56 _dafny.Sequence
			_ = _2048_v56
			_2048_v56 = _dafny.SeqOf((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
			var _2049_v57 _dafny.MultiSet
			_ = _2049_v57
			_2049_v57 = _dafny.MultiSetOf((_this).Fm11(_dafny.IntOfInt64(-498), p1, globalState), (_this).F24())
			var _2050_v58 _dafny.Sequence
			_ = _2050_v58
			_2050_v58 = _dafny.SeqOf(p3)
			var _2051_v59 _dafny.Set
			_ = _2051_v59
			_2051_v59 = _dafny.SetOf(_2050_v58)
			var _2052_v60 _dafny.Int
			_ = _2052_v60
			var _2053_v61 bool
			_ = _2053_v61
			var _2054_v62 _dafny.Int
			_ = _2054_v62
			var _out55 _dafny.Int
			_ = _out55
			var _out56 bool
			_ = _out56
			var _out57 _dafny.Int
			_ = _out57
			_out55, _out56, _out57 = Companion_Default___.M0(_dafny.IntOfUint32((_2048_v56).Cardinality()), _2049_v57, !(!((_2051_v59).IsProperSubsetOf(_2051_v59))), (p3).Times((_dafny.Zero).Minus(p3)), globalState)
			_2052_v60 = _out55
			_2053_v61 = _out56
			_2054_v62 = _out57
			var _2055_v63 _dafny.Set
			_ = _2055_v63
			_2055_v63 = _dafny.SetOf(p1, p2, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
			var _2056_v64 *C8
			_ = _2056_v64
			var _nw345 *C8 = New_C8_()
			_ = _nw345
			_nw345.Ctor__(!(_2055_v63).Equals(_2055_v63))
			_2056_v64 = _nw345
			_2056_v64 = _2056_v64
			var _2057_v65 _dafny.Set
			_ = _2057_v65
			_2057_v65 = _dafny.SetOf((_2052_v60).Times((_2050_v58).Select((Companion_Default___.SafeIndex((_2055_v63).Cardinality(), _dafny.IntOfUint32((_2050_v58).Cardinality()))).Uint32()).(_dafny.Int)), _2054_v62, p3, _2052_v60, _2052_v60)
			_2057_v65 = _dafny.SetOf(p3)
		}
		var _hi13 _dafny.Int = p3
		_ = _hi13
		for _2058_i6 := (_dafny.Zero).Minus(p3); _2058_i6.Cmp(_hi13) < 0; _2058_i6 = _2058_i6.Plus(_dafny.One) {
			if (func() bool {
				if (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool) {
					return p1
				}
				return p1
			})() {
				var _2059_v66 _dafny.Sequence
				_ = _2059_v66
				_2059_v66 = _dafny.SeqOf(p1, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), p1)
				(globalState).F7 = (_dafny.IntOfUint32((_2059_v66).Cardinality())).Times(p3)
				var _2060_v67 _dafny.Map
				_ = _2060_v67
				_2060_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1991_v15, p3)
				var _2061_v68 _dafny.MultiSet
				_ = _2061_v68
				_2061_v68 = _dafny.MultiSetOf(p2)
				_2060_v67 = (_2060_v67).Update(_1991_v15, Companion_Default___.SafeModuloInt(_2058_i6, (_2061_v68).Cardinality()))
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
				_ = _index282
				(_1991_v15).ArraySet1((_2061_v68).Equals((_2061_v68).Union(_2061_v68)), (_index282).Int())
				(globalState).F13 = !(p1) || ((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
				var _2062_v69 _dafny.Map
				_ = _2062_v69
				_2062_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1991_v15, _this.F23())
				var _2063_v70 *C3
				_ = _2063_v70
				var _nw346 *C3 = New_C3_()
				_ = _nw346
				_nw346.Ctor__(_2062_v69, _1994_v16, _this.F23(), !((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool)) || (p2))
				_2063_v70 = _nw346
			} else {
				var _2064_v71 _dafny.Array
				_ = _2064_v71
				var _nw347 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw347
				_2064_v71 = _nw347
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_2064_v71), 0))
				_ = _index283
				(_2064_v71).ArraySet1(p3, (_index283).Int())
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_2064_v71), 0))
				_ = _index284
				(_2064_v71).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1971_v0, _1971_v0)).Cardinality()), (_index284).Int())
				var _2065_v72 _dafny.Map
				_ = _2065_v72
				_2065_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1991_v15, _this.F23())
				var _2066_v73 *C3
				_ = _2066_v73
				var _nw348 *C3 = New_C3_()
				_ = _nw348
				_nw348.Ctor__(_2065_v72, _1994_v16, _this.F23(), p1)
				_2066_v73 = _nw348
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_2064_v71), 0))
				_ = _index285
				(_2064_v71).ArraySet1((_2064_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_2064_v71), 0))).Int()).(_dafny.Int), (_index285).Int())
				var _2067_v74 _dafny.MultiSet
				_ = _2067_v74
				_2067_v74 = _dafny.MultiSetOf((_2064_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_2064_v71), 0))).Int()).(_dafny.Int))
				var _2068_v75 *C1
				_ = _2068_v75
				var _nw349 *C1 = New_C1_()
				_ = _nw349
				_nw349.Ctor__(true, (_dafny.MultiSetOf(_dafny.IntOfInt64(-913))).IsProperSubsetOf(_2067_v74))
				_2068_v75 = _nw349
				var _2069_v76 D5
				_ = _2069_v76
				_2069_v76 = Companion_D5_.Create_DC12_(_dafny.SetOf(p3))
				var _2070_v77 _dafny.MultiSet
				_ = _2070_v77
				_2070_v77 = _dafny.MultiSetOf(_2069_v76)
				_2070_v77 = _2070_v77
			}
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1994_v16), 0))
			_ = _index286
			(_1994_v16).ArraySet1CodePoint(_this.F23(), (_index286).Int())
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1994_v16), 0))
			_ = _index287
			var _rhs310 _dafny.Int = Companion_Default___.SafeDivisionInt(_2058_i6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(341))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}((func(_2071_v0 _dafny.Sequence, _2072_p3 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_2073_i7 _dafny.Int) _dafny.CodePoint {
					return (_2071_v0).Select((Companion_Default___.SafeIndex(_2072_p3, _dafny.IntOfUint32((_2071_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_1971_v0, p3))), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(341))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}((func(_2074_v0 _dafny.Sequence, _2075_p3 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_2076_i7 _dafny.Int) _dafny.CodePoint {
					return (_2074_v0).Select((Companion_Default___.SafeIndex(_2075_p3, _dafny.IntOfUint32((_2074_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_1971_v0, p3)))).Cardinality()))).Uint32(), _this.F23()), _1971_v0)).Cardinality()))
			_ = _rhs310
			var _rhs311 _dafny.CodePoint = _dafny.CodePoint('n')
			_ = _rhs311
			var _rhs312 _dafny.Array = _1991_v15
			_ = _rhs312
			var _rhs313 _dafny.CodePoint = _this.F23()
			_ = _rhs313
			var _lhs230 *GlobalState = globalState
			_ = _lhs230
			var _lhs231 _dafny.Array = _1994_v16
			_ = _lhs231
			var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1994_v16), 0))
			_ = _lhs232
			var _lhs233 *C9 = _this
			_ = _lhs233
			_lhs230.F7 = _rhs310
			(_lhs231).ArraySet1CodePoint(_rhs311, (_lhs232).Int())
			_1991_v15 = _rhs312
			_lhs233.F23_set_(_rhs313)
			r0 = !(true)
			var _2077_v78 _dafny.Sequence
			_ = _2077_v78
			_2077_v78 = _dafny.SeqOf(p1)
			var _2078_v79 _dafny.Set
			_ = _2078_v79
			_2078_v79 = _dafny.SetOf((_this).F24(), !((_this).F24()), (_this).Fm4(_dafny.IntOfInt64(-465), _2077_v78, globalState))
			(globalState).F7 = (_2078_v79).Cardinality()
		}
		if p2 {
			var _2079_v80 _dafny.MultiSet
			_ = _2079_v80
			_2079_v80 = _dafny.MultiSetOf(!(p1))
			var _2080_v81 _dafny.Int
			_ = _2080_v81
			var _2081_v82 bool
			_ = _2081_v82
			var _2082_v83 _dafny.Int
			_ = _2082_v83
			var _out58 _dafny.Int
			_ = _out58
			var _out59 bool
			_ = _out59
			var _out60 _dafny.Int
			_ = _out60
			_out58, _out59, _out60 = Companion_Default___.M0((_this).Fm6(globalState), (Companion_Default___.Fm3(p2, globalState)).Union(_2079_v80), p2, p3, globalState)
			_2080_v81 = _out58
			_2081_v82 = _out59
			_2082_v83 = _out60
			var _2083_v84 D13
			_ = _2083_v84
			_2083_v84 = Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F23()))
			var _pat_let_tv35 = _2081_v82
			_ = _pat_let_tv35
			var _2084_v85 _dafny.MultiSet
			_ = _2084_v85
			_2084_v85 = _dafny.MultiSetOf(_2083_v84, (func() D13 {
				if (_this).F24() {
					return func(_pat_let53_0 D13) D13 {
						return func(_2085_dt__update__tmp_h2 D13) D13 {
							return func(_pat_let54_0 _dafny.Map) D13 {
								return func(_2086_dt__update_hcf44_h0 _dafny.Map) D13 {
									return Companion_D13_.Create_DC27_(_2086_dt__update_hcf44_h0)
								}(_pat_let54_0)
							}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv35, _this.F23()))
						}(_pat_let53_0)
					}(_2083_v84)
				}
				return _2083_v84
			})(), _2083_v84)
			_2084_v85 = Companion_Default___.Fm50((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), _2080_v81, globalState)
			var _2087_v86 _dafny.Sequence
			_ = _2087_v86
			_2087_v86 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(599))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}((func(_2088_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_2089_i8 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_2088_v0).Cardinality())
				}
			})(_1971_v0)))).Cardinality()))
			var _2090_v87 _dafny.Map
			_ = _2090_v87
			_2090_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2087_v86, (Companion_Default___.SafeIndex(_2080_v81, _dafny.IntOfUint32((_2087_v86).Cardinality()))).Uint32(), _2082_v83)).Cardinality()), _this.F23())
			(_this).F23_set_((func() _dafny.CodePoint {
				if _2081_v82 {
					return _this.F23()
				}
				return (func() _dafny.CodePoint {
					if (_2090_v87).Contains(_dafny.IntOfInt64(-794)) {
						return (_2090_v87).Get(_dafny.IntOfInt64(-794)).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('b')
				})()
			})())
			_2080_v81 = p3
			var _2091_v88 _dafny.Set
			_ = _2091_v88
			_2091_v88 = _dafny.SetOf(_2081_v82)
			(globalState).F7 = (_dafny.Zero).Minus(((_2091_v88).Union(_2091_v88)).Cardinality())
		} else {
			(globalState).F7 = _dafny.IntOfInt64(650)
			var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
			_ = _index288
			(_1991_v15).ArraySet1((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), (_index288).Int())
			var _2092_v89 D2
			_ = _2092_v89
			_2092_v89 = Companion_D2_.Create_DC7_(!(false), p2)
			var _2093_v90 _dafny.Sequence
			_ = _2093_v90
			_2093_v90 = _dafny.SeqOf(_2092_v89)
			_2093_v90 = _dafny.Companion_Sequence_.Concatenate(_2093_v90, _dafny.Companion_Sequence_.Concatenate(_2093_v90, _dafny.SeqOf(_2092_v89)))
			var _2094_v91 _dafny.MultiSet
			_ = _2094_v91
			_2094_v91 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("aynidalob"), _1971_v0, _1971_v0, _1971_v0, _1971_v0)
			var _2095_v92 _dafny.Set
			_ = _2095_v92
			_2095_v92 = _dafny.SetOf(_dafny.IntOfInt64(271), (func() _dafny.Int {
				if (_2094_v91).Contains(_1971_v0) {
					return (_2094_v91).Multiplicity(_1971_v0)
				}
				return p3
			})())
			_2095_v92 = _2095_v92
			var _2096_v93 _dafny.Set
			_ = _2096_v93
			_2096_v93 = _dafny.SetOf((_this).F24())
			var _2097_v94 _dafny.Map
			_ = _2097_v94
			_2097_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_2096_v93).Cardinality())
			_2097_v94 = _2097_v94
		}
		if (func() bool {
			if true {
				return (_this).F24()
			}
			return (_this).F24()
		})() {
			var _2098_v95 _dafny.Map
			_ = _2098_v95
			_2098_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), _dafny.IntOfUint32((_1971_v0).Cardinality()), globalState), p3)
			var _2099_v96 _dafny.Set
			_ = _2099_v96
			_2099_v96 = _dafny.SetOf(_2098_v95)
			var _2100_v97 _dafny.Map
			_ = _2100_v97
			_2100_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
			var _2101_v98 _dafny.Sequence
			_ = _2101_v98
			_2101_v98 = _dafny.SeqOf((_this).F24(), (_this).F24())
			var _2102_v99 *C8
			_ = _2102_v99
			var _nw350 *C8 = New_C8_()
			_ = _nw350
			_nw350.Ctor__(p2)
			_2102_v99 = _nw350
			var _2103_v100 _dafny.Map
			_ = _2103_v100
			_2103_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), _2102_v99)
			_2099_v96 = Companion_Default___.Fm51(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2100_v97, _2101_v98), p3, p3, (func() bool {
				if p1 {
					return (_2101_v98).Select((Companion_Default___.SafeIndex((_2103_v100).Cardinality(), _dafny.IntOfUint32((_2101_v98).Cardinality()))).Uint32()).(bool)
				}
				return p1
			})(), globalState)
			var _2104_v101 _dafny.MultiSet
			_ = _2104_v101
			_2104_v101 = _dafny.MultiSetOf(p2, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool), (_2101_v98).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_2101_v98).Cardinality()))).Uint32()).(bool))
			if !((_dafny.MultiSetOf(p3)).Update((_2104_v101).Cardinality(), Companion_Default___.Abs(p3))).Contains(Companion_Default___.Fm0(!((_this).F24()), p3, globalState)) {
				var _2105_v102 _dafny.Int
				_ = _2105_v102
				var _2106_v103 bool
				_ = _2106_v103
				var _2107_v104 _dafny.Int
				_ = _2107_v104
				var _out61 _dafny.Int
				_ = _out61
				var _out62 bool
				_ = _out62
				var _out63 _dafny.Int
				_ = _out63
				_out61, _out62, _out63 = Companion_Default___.M0(p3, (_2104_v101).Union(_2104_v101), p1, p3, globalState)
				_2105_v102 = _out61
				_2106_v103 = _out62
				_2107_v104 = _out63
				(globalState).F7 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1971_v0).Cardinality()))
				var _2108_v105 _dafny.Array
				_ = _2108_v105
				var _nw351 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
				_ = _nw351
				_2108_v105 = _nw351
				var _2109_v106 *C6
				_ = _2109_v106
				var _nw352 *C6 = New_C6_()
				_ = _nw352
				_nw352.Ctor__(_this.F23(), !(Companion_Default___.Fm2(globalState)))
				_2109_v106 = _nw352
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_2108_v105), 0))
				_ = _index289
				(_2108_v105).ArraySet1(_2109_v106, (_index289).Int())
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_2108_v105), 0))
				_ = _index290
				(_2108_v105).ArraySet1(_2109_v106, (_index290).Int())
				var _2110_v107 D13
				_ = _2110_v107
				_2110_v107 = Companion_D13_.Create_DC28_(_dafny.Companion_Sequence_.IsProperPrefixOf(_1971_v0, _1971_v0), (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1971_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2101_v98).Cardinality()), _dafny.IntOfUint32((_1971_v0).Cardinality()))).Uint32(), _this.F23()), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_2105_v102), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1971_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2101_v98).Cardinality()), _dafny.IntOfUint32((_1971_v0).Cardinality()))).Uint32(), _this.F23())).Cardinality()))).Uint32(), _this.F23())).Cardinality())).Times(p3)), (_2107_v104).Plus(_dafny.IntOfInt64(920)), (_2107_v104).Times(_dafny.IntOfInt64(592)))
				var _rhs314 D13 = _2110_v107
				_ = _rhs314
				var _rhs315 bool = !(_2106_v103) || ((_this).F24())
				_ = _rhs315
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				_2110_v107 = _rhs314
				_lhs234.F13 = _rhs315
				_2107_v104 = (Companion_D15_.Create_DC36_(Companion_Default___.Fm2(globalState), (_2100_v97).Cardinality())).Dtor_cf60()
			} else {
				var _2111_v108 _dafny.Map
				_ = _2111_v108
				_2111_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.CodePoint('t'))
				var _2112_v109 *C5
				_ = _2112_v109
				var _nw353 *C5 = New_C5_()
				_ = _nw353
				_nw353.Ctor__(_1991_v15)
				_2112_v109 = _nw353
				var _2113_v110 _dafny.Map
				_ = _2113_v110
				_2113_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2112_v109, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
				var _2114_v111 _dafny.Map
				_ = _2114_v111
				_2114_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2113_v110).Cardinality(), p1)
				_2111_v108 = Companion_Default___.Fm52((_2114_v111).Cardinality(), !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_this.F23()), _1971_v0), p3, p3, globalState)
				var _2115_v112 T1
				_ = _2115_v112
				var _nw354 *C6 = New_C6_()
				_ = _nw354
				_nw354.Ctor__(_this.F23(), (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
				_2115_v112 = _nw354
				var _2116_v113 D5
				_ = _2116_v113
				_2116_v113 = Companion_D5_.Create_DC13_((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p3, (func() _dafny.Int {
					if (_2104_v101).Contains((_this).F24()) {
						return (_2104_v101).Multiplicity((_this).F24())
					}
					return _dafny.IntOfUint32((Companion_Default___.Fm1((_2115_v112).F24(), globalState)).Cardinality())
				})())), (_2115_v112).F24(), (_2115_v112).F24())
				var _2117_v114 _dafny.Map
				_ = _2117_v114
				_2117_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2115_v112).F24(), _dafny.IntOfInt64(522))
				var _2118_v115 _dafny.Map
				_ = _2118_v115
				_2118_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1991_v15, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
				var _2119_v116 _dafny.Map
				_ = _2119_v116
				_2119_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), (_2118_v115).Cardinality())
				var _2120_v117 *C6
				_ = _2120_v117
				var _nw355 *C6 = New_C6_()
				_ = _nw355
				_nw355.Ctor__(_2115_v112.F23(), p2)
				_2120_v117 = _nw355
				var _2121_v118 _dafny.Map
				_ = _2121_v118
				_2121_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2120_v117, p3)
				var _2122_v119 _dafny.Array
				_ = _2122_v119
				var _nwElement0_62 _dafny.Int = p3
				_ = _nwElement0_62
				var _nw356 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(29))
				_ = _nw356
				(_nw356).ArraySet1(_nwElement0_62, 0)
				(_nw356).ArraySet1(_dafny.IntOfInt64(565), 1)
				(_nw356).ArraySet1(p3, 2)
				(_nw356).ArraySet1(p3, 3)
				(_nw356).ArraySet1(p3, 4)
				(_nw356).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F23())).Merge(_2111_v108)).Cardinality(), 5)
				(_nw356).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(216), (_2117_v114).Cardinality()), 6)
				(_nw356).ArraySet1(p3, 7)
				(_nw356).ArraySet1(p3, 8)
				(_nw356).ArraySet1((p3).Times(_dafny.IntOfUint32((_1971_v0).Cardinality())), 9)
				(_nw356).ArraySet1(p3, 10)
				(_nw356).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), (_this).Fm6(globalState))).Merge(_2119_v116)).Cardinality(), 11)
				(_nw356).ArraySet1(p3, 12)
				(_nw356).ArraySet1(p3, 13)
				(_nw356).ArraySet1((func() _dafny.Int {
					if (_2104_v101).Contains((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool)) {
						return (_2104_v101).Multiplicity((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
					}
					return _dafny.IntOfInt64(218)
				})(), 14)
				(_nw356).ArraySet1((p3).Minus(p3), 15)
				(_nw356).ArraySet1(Companion_Default___.SafeDivisionInt((_2121_v118).Cardinality(), p3), 16)
				(_nw356).ArraySet1(p3, 17)
				(_nw356).ArraySet1(p3, 18)
				(_nw356).ArraySet1(_dafny.IntOfInt64(463), 19)
				(_nw356).ArraySet1(Companion_Default___.SafeDivisionInt(p3, p3), 20)
				(_nw356).ArraySet1(_dafny.IntOfUint32((_1971_v0).Cardinality()), 21)
				(_nw356).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(287), p3), 22)
				(_nw356).ArraySet1(_dafny.IntOfInt64(-90), 23)
				(_nw356).ArraySet1(_dafny.IntOfUint32((_2101_v98).Cardinality()), 24)
				(_nw356).ArraySet1(p3, 25)
				(_nw356).ArraySet1(p3, 26)
				(_nw356).ArraySet1(_dafny.IntOfInt64(420), 27)
				(_nw356).ArraySet1(p3, 28)
				_2122_v119 = _nw356
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2122_v119), 0))
				_ = _index291
				(_2122_v119).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p3, p3, p3)).Cardinality()), (_index291).Int())
				var _2123_v120 _dafny.MultiSet
				_ = _2123_v120
				_2123_v120 = _dafny.MultiSetOf(p3)
				var _2124_v121 D11
				_ = _2124_v121
				_2124_v121 = Companion_D11_.Create_DC23_((_this).F24(), (_2123_v120).Cardinality(), _2100_v97, Companion_D5_.Create_DC14_((_2115_v112).F24()))
				var _pat_let_tv36 = p3
				_ = _pat_let_tv36
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2122_v119), 0))
				_ = _index292
				var _rhs316 T1 = _2115_v112
				_ = _rhs316
				var _rhs317 bool = !((func() bool {
					if (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool) {
						return (_this).F24()
					}
					return p1
				})())
				_ = _rhs317
				var _rhs318 D5 = func(_pat_let55_0 D5) D5 {
					return func(_2125_dt__update__tmp_h3 D5) D5 {
						return func(_pat_let56_0 _dafny.Int) D5 {
							return func(_2126_dt__update_hcf20_h0 _dafny.Int) D5 {
								return Companion_D5_.Create_DC13_(_2126_dt__update_hcf20_h0, (_2125_dt__update__tmp_h3).Dtor_cf21(), (_2125_dt__update__tmp_h3).Dtor_cf22())
							}(_pat_let56_0)
						}(_pat_let_tv36)
					}(_pat_let55_0)
				}(_2116_v113)
				_ = _rhs318
				var _rhs319 _dafny.Int = Companion_Default___.Fm0(p2, (_2104_v101).Cardinality(), globalState)
				_ = _rhs319
				var _rhs320 _dafny.Int = (func() _dafny.Int {
					if (_2124_v121).Dtor_cf36() {
						return _dafny.IntOfInt64(390)
					}
					return (func() _dafny.Int {
						if (_2123_v120).Contains(p3) {
							return (_2123_v120).Multiplicity(p3)
						}
						return (_2117_v114).Cardinality()
					})()
				})()
				_ = _rhs320
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				var _lhs236 _dafny.Array = _2122_v119
				_ = _lhs236
				var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2122_v119), 0))
				_ = _lhs237
				var _lhs238 *GlobalState = globalState
				_ = _lhs238
				_2115_v112 = _rhs316
				_lhs235.F13 = _rhs317
				_2116_v113 = _rhs318
				(_lhs236).ArraySet1(_rhs319, (_lhs237).Int())
				_lhs238.F7 = _rhs320
				(globalState).F13 = p1
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_2122_v119), 0))
				_ = _index293
				(_2122_v119).ArraySet1((_2115_v112).Fm6(globalState), (_index293).Int())
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
				_ = _index294
				(_1991_v15).ArraySet1(!(!(true) || (p1)) || (!((_2115_v112).F24()) || ((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))), (_index294).Int())
			}
			if (_this).F24() {
				var _2127_v122 *C6
				_ = _2127_v122
				var _nw357 *C6 = New_C6_()
				_ = _nw357
				_nw357.Ctor__(_this.F23(), p1)
				_2127_v122 = _nw357
				var _2128_v123 D5
				_ = _2128_v123
				_2128_v123 = Companion_D5_.Create_DC14_(p1)
				var _2129_v124 D11
				_ = _2129_v124
				_2129_v124 = Companion_D11_.Create_DC23_((_this).F24(), _dafny.IntOfUint32((_1971_v0).Cardinality()), _2100_v97, _2128_v123)
				_1991_v15 = (func() _dafny.Array {
					if (_2129_v124).Dtor_cf36() {
						return _1991_v15
					}
					return _1991_v15
				})()
				var _2130_v125 _dafny.Array
				_ = _2130_v125
				var _len0_65 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_65
				var _nw358 _dafny.Array
				_ = _nw358
				if _len0_65.Cmp(_dafny.Zero) == 0 {
					_nw358 = _dafny.NewArray(_len0_65)
				} else {
					var _init65 func(_dafny.Int) _dafny.Sequence = (func(_2131_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2132_i9 _dafny.Int) _dafny.Sequence {
							return _2131_v0
						}
					})(_1971_v0)
					_ = _init65
					var _element0_65 = _init65(_dafny.Zero)
					_ = _element0_65
					_nw358 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
					(_nw358).ArraySet1(_element0_65, 0)
					var _nativeLen0_65 = (_len0_65).Int()
					_ = _nativeLen0_65
					for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
						(_nw358).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
					}
				}
				_2130_v125 = _nw358
				_2130_v125 = _2130_v125
				_2098_v95 = (_2098_v95).Update(p3, (func() _dafny.Int {
					if (_2104_v101).Contains(false) {
						return (_2104_v101).Multiplicity(false)
					}
					return p3
				})())
				var _2133_v126 _dafny.Set
				_ = _2133_v126
				_2133_v126 = _dafny.SetOf(((_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool)) == ((_this).F24()))
				_2133_v126 = (_2133_v126).Union(_2133_v126)
			} else {
				r0 = p1
				(globalState).F13 = p1
				(globalState).F13 = !_dafny.Companion_Sequence_.Contains(_1971_v0, _this.F23())
				r0 = p2
				var _2134_v127 _dafny.Set
				_ = _2134_v127
				_2134_v127 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_1971_v0, _dafny.UnicodeSeqOfUtf8Bytes("ykpfixig")))
				_2134_v127 = Companion_Default___.Fm53(globalState)
			}
			_2100_v97 = (_2100_v97).Update(p2, (_1991_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))).Int()).(bool))
			(globalState).F7 = p3
		} else {
			var _2135_v128 _dafny.Map
			_ = _2135_v128
			_2135_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			var _2136_v129 T2
			_ = _2136_v129
			var _nw359 *C4 = New_C4_()
			_ = _nw359
			_nw359.Ctor__(p3, _2135_v128, _1994_v16, _this.F23(), (_this).F24())
			_2136_v129 = _nw359
			var _2137_v130 _dafny.Map
			_ = _2137_v130
			_2137_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2136_v129)
			var _2138_v131 _dafny.Sequence
			_ = _2138_v131
			_2138_v131 = _dafny.SeqOf(p2)
			_2137_v130 = (_2137_v130).Update((_dafny.IntOfUint32((_2138_v131).Cardinality())).Plus(p3), _2136_v129)
			var _2139_v132 _dafny.MultiSet
			_ = _2139_v132
			_2139_v132 = _dafny.MultiSetOf(Companion_Default___.Fm2(globalState))
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
			_ = _index295
			(_1991_v15).ArraySet1((!(p1) || (true)) || ((_2139_v132).IsSubsetOf(_2139_v132)), (_index295).Int())
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1991_v15), 0))
			_ = _index296
			(_1991_v15).ArraySet1(!((_this).F24()), (_index296).Int())
			(globalState).F7 = _dafny.IntOfInt64(57)
			(globalState).F13 = p1
		}
		var _pat_let_tv37 = p1
		_ = _pat_let_tv37
		var _pat_let_tv38 = p1
		_ = _pat_let_tv38
		var _pat_let_tv39 = _1991_v15
		_ = _pat_let_tv39
		var _pat_let_tv40 = _1991_v15
		_ = _pat_let_tv40
		var _pat_let_tv41 = p3
		_ = _pat_let_tv41
		var _pat_let_tv42 = p3
		_ = _pat_let_tv42
		var _pat_let_tv43 = p2
		_ = _pat_let_tv43
		var _pat_let_tv44 = _1971_v0
		_ = _pat_let_tv44
		r0 = func(_source29 D5) bool {
			if _source29.Is_DC13() {
				var _2140___mcc_h4 _dafny.Int = _source29.Get_().(D5_DC13).Cf20
				_ = _2140___mcc_h4
				var _2141___mcc_h5 bool = _source29.Get_().(D5_DC13).Cf21
				_ = _2141___mcc_h5
				var _2142___mcc_h6 bool = _source29.Get_().(D5_DC13).Cf22
				_ = _2142___mcc_h6
				var _2143_cf22 bool = _2142___mcc_h6
				_ = _2143_cf22
				var _2144_cf21 bool = _2141___mcc_h5
				_ = _2144_cf21
				var _2145_cf20 _dafny.Int = _2140___mcc_h4
				_ = _2145_cf20
				return (_pat_let_tv37) == (_pat_let_tv38)
			} else if _source29.Is_DC14() {
				var _2146___mcc_h7 bool = _source29.Get_().(D5_DC14).Cf23
				_ = _2146___mcc_h7
				var _2147_cf23 bool = _2146___mcc_h7
				_ = _2147_cf23
				return (_pat_let_tv40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_pat_let_tv39), 0))).Int()).(bool)
			} else {
				var _2148___mcc_h8 _dafny.Set = _source29.Get_().(D5_DC12).Cf19
				_ = _2148___mcc_h8
				var _2149_cf19 _dafny.Set = _2148___mcc_h8
				_ = _2149_cf19
				return (_dafny.MultiSetOf(_pat_let_tv41, _pat_let_tv42)).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv43)).Cardinality()), _dafny.IntOfUint32((_pat_let_tv44).Cardinality()), _dafny.IntOfInt64(622)))
			}
		}(Companion_D5_.Create_DC13_(_dafny.IntOfUint32((_1971_v0).Cardinality()), true, p1))
		return r0
	}
}
func (_this *C9) M6(p0 bool, p1 D0, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		if p3 {
			var _2150_v0 _dafny.MultiSet
			_ = _2150_v0
			_2150_v0 = _dafny.MultiSetOf(true)
			var _2151_v1 _dafny.Map
			_ = _2151_v1
			_2151_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2150_v0, _this.F23())
			_2151_v1 = _2151_v1
			var _2152_v2 _dafny.Int
			_ = _2152_v2
			_2152_v2 = _dafny.IntOfInt64(641)
			var _2153_v3 _dafny.Sequence
			_ = _2153_v3
			_2153_v3 = _dafny.SeqOf(true)
			var _2154_v4 _dafny.Map
			_ = _2154_v4
			_2154_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2152_v2, _2152_v2)
			var _2155_v5 _dafny.Map
			_ = _2155_v5
			_2155_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2152_v2)
			var _2156_v6 _dafny.Array
			_ = _2156_v6
			var _nwElement0_63 _dafny.Map = _2154_v4
			_ = _nwElement0_63
			var _nw360 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(10))
			_ = _nw360
			(_nw360).ArraySet1(_nwElement0_63, 0)
			(_nw360).ArraySet1(_2154_v4, 1)
			(_nw360).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), _2152_v2), 2)
			(_nw360).ArraySet1(_2154_v4, 3)
			(_nw360).ArraySet1(_2154_v4, 4)
			(_nw360).ArraySet1(_2154_v4, 5)
			(_nw360).ArraySet1(_2154_v4, 6)
			(_nw360).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2150_v0).Cardinality(), (_dafny.Zero).Minus((_2155_v5).Cardinality())), 7)
			(_nw360).ArraySet1(_2154_v4, 8)
			(_nw360).ArraySet1(_2154_v4, 9)
			_2156_v6 = _nw360
			var _2157_v7 _dafny.Map
			_ = _2157_v7
			_2157_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2156_v6, (_this).Fm6(globalState))
			var _2158_v8 D9
			_ = _2158_v8
			_2158_v8 = Companion_D9_.Create_DC20_(p2, p0, _2152_v2)
			var _2159_v9 _dafny.Sequence
			_ = _2159_v9
			_2159_v9 = _dafny.SeqOf(_dafny.IntOfInt64(-691))
			var _2160_v10 _dafny.MultiSet
			_ = _2160_v10
			_2160_v10 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2159_v9).Cardinality()))
			var _2161_v11 _dafny.Array
			_ = _2161_v11
			var _nwElement0_64 _dafny.Int = _2152_v2
			_ = _nwElement0_64
			var _nw361 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(22))
			_ = _nw361
			(_nw361).ArraySet1(_nwElement0_64, 0)
			(_nw361).ArraySet1(_2152_v2, 1)
			(_nw361).ArraySet1(_2152_v2, 2)
			(_nw361).ArraySet1(_2152_v2, 3)
			(_nw361).ArraySet1(_2152_v2, 4)
			(_nw361).ArraySet1((_2152_v2).Plus(_2152_v2), 5)
			(_nw361).ArraySet1(_2152_v2, 6)
			(_nw361).ArraySet1(_2152_v2, 7)
			(_nw361).ArraySet1((func() _dafny.Int {
				if (_2150_v0).Contains((_this).F24()) {
					return (_2150_v0).Multiplicity((_this).F24())
				}
				return _2152_v2
			})(), 8)
			(_nw361).ArraySet1(_dafny.IntOfInt64(648), 9)
			(_nw361).ArraySet1(_2152_v2, 10)
			(_nw361).ArraySet1(_2152_v2, 11)
			(_nw361).ArraySet1((func() _dafny.Int {
				if (_2153_v3).Select((Companion_Default___.SafeIndex(_2152_v2, _dafny.IntOfUint32((_2153_v3).Cardinality()))).Uint32()).(bool) {
					return _2152_v2
				}
				return (func() _dafny.Int {
					if (_2157_v7).Contains(_2156_v6) {
						return (_2157_v7).Get(_2156_v6).(_dafny.Int)
					}
					return _2152_v2
				})()
			})(), 12)
			(_nw361).ArraySet1((_2152_v2).Plus(((_dafny.MultiSetOf(_2158_v8)).Update(Companion_D9_.Create_DC20_(p2, (_this).F24(), (_2160_v10).Cardinality()), Companion_Default___.Abs(_2152_v2))).Cardinality()), 13)
			(_nw361).ArraySet1(Companion_Default___.SafeDivisionInt(_2152_v2, _dafny.IntOfUint32((p2).Cardinality())), 14)
			(_nw361).ArraySet1((_dafny.Zero).Minus(_2152_v2), 15)
			(_nw361).ArraySet1(_2152_v2, 16)
			(_nw361).ArraySet1((_2152_v2).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xhna")).Cardinality())), 17)
			(_nw361).ArraySet1(_2152_v2, 18)
			(_nw361).ArraySet1(_2152_v2, 19)
			(_nw361).ArraySet1(_2152_v2, 20)
			(_nw361).ArraySet1((func() _dafny.Int {
				if (_this).F24() {
					return _2152_v2
				}
				return _2152_v2
			})(), 21)
			_2161_v11 = _nw361
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_2161_v11), 0))
			_ = _index297
			(_2161_v11).ArraySet1(_2152_v2, (_index297).Int())
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_2161_v11), 0))
			_ = _index298
			(_2161_v11).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((p1).Dtor_cf3()), _2152_v2), (_index298).Int())
			var _2162_v12 _dafny.Set
			_ = _2162_v12
			_2162_v12 = _dafny.SetOf(_2150_v0, _2150_v0, _2150_v0, (_dafny.MultiSetOf((_this).F24())).Intersection(_dafny.MultiSetFromSeq(_2153_v3)))
			_2162_v12 = _dafny.SetOf(Companion_Default___.Fm3(p3, globalState), (_2150_v0).Update((_this).F24(), Companion_Default___.Abs((_2161_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_2161_v11), 0))).Int()).(_dafny.Int))))
			var _2163_v13 _dafny.Sequence
			_ = _2163_v13
			_2163_v13 = _dafny.UnicodeSeqOfUtf8Bytes("d")
			var _2164_v14 _dafny.Map
			_ = _2164_v14
			_2164_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2152_v2, _this.F23())
			_2163_v13 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1((func() bool {
				if (_this).F24() {
					return p3
				}
				return (_this).F24()
			})(), globalState), (Companion_Default___.SafeIndex((_2161_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_2161_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm1((func() bool {
				if (_this).F24() {
					return p3
				}
				return (_this).F24()
			})(), globalState)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
				if (_2164_v14).Contains(_dafny.IntOfInt64(702)) {
					return (_2164_v14).Get(_dafny.IntOfInt64(702)).(_dafny.CodePoint)
				}
				return _this.F23()
			})())
			var _2165_v15 _dafny.Array
			_ = _2165_v15
			var _nw362 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw362
			_2165_v15 = _nw362
			var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_2165_v15), 0))
			_ = _index299
			(_2165_v15).ArraySet1(!((_this).F24()) || (p3), (_index299).Int())
			var _2166_v16 _dafny.Map
			_ = _2166_v16
			_2166_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), false)
			var _2167_v17 _dafny.Map
			_ = _2167_v17
			_2167_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2163_v13)
			var _2168_v18 T0
			_ = _2168_v18
			var _nw363 *C0 = New_C0_()
			_ = _nw363
			_nw363.Ctor__(_2167_v17, true, _this.F23(), (_this).F24())
			_2168_v18 = _nw363
			var _2169_v19 _dafny.Map
			_ = _2169_v19
			_2169_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2168_v18, (_this).F24())
			var _2170_v20 T3
			_ = _2170_v20
			var _nw364 *C1 = New_C1_()
			_ = _nw364
			_nw364.Ctor__(p3, (func() bool {
				if (_2166_v16).Contains((func() bool {
					if (_2169_v19).Contains(_2168_v18) {
						return (_2169_v19).Get(_2168_v18).(bool)
					}
					return true
				})()) {
					return (_2166_v16).Get((func() bool {
						if (_2169_v19).Contains(_2168_v18) {
							return (_2169_v19).Get(_2168_v18).(bool)
						}
						return true
					})()).(bool)
				}
				return true
			})())
			_2170_v20 = _nw364
			var _2171_v21 _dafny.Sequence
			_ = _2171_v21
			_2171_v21 = _dafny.SeqOf(_2170_v20)
			var _2172_v22 _dafny.Map
			_ = _2172_v22
			_2172_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2163_v13, _2171_v21)
			var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_2165_v15), 0))
			_ = _index300
			var _rhs321 _dafny.Int = ((_2172_v22).Update(_2163_v13, _2171_v21)).Cardinality()
			_ = _rhs321
			var _rhs322 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfInt64(-745))).Minus(_dafny.IntOfUint32((_2163_v13).Cardinality()))
			_ = _rhs322
			var _rhs323 bool = (_2168_v18).F24()
			_ = _rhs323
			var _rhs324 bool = (_2170_v20).F27()
			_ = _rhs324
			var _lhs239 *GlobalState = globalState
			_ = _lhs239
			var _lhs240 *GlobalState = globalState
			_ = _lhs240
			var _lhs241 _dafny.Array = _2165_v15
			_ = _lhs241
			var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_2165_v15), 0))
			_ = _lhs242
			_lhs239.F7 = _rhs321
			_lhs240.F7 = _rhs322
			r0 = _rhs323
			(_lhs241).ArraySet1(_rhs324, (_lhs242).Int())
		} else {
			var _2173_v23 _dafny.Int
			_ = _2173_v23
			_2173_v23 = _dafny.IntOfInt64(861)
			var _2174_v24 _dafny.Map
			_ = _2174_v24
			_2174_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p2)
			var _2175_v25 _dafny.Map
			_ = _2175_v25
			_2175_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2173_v23, _2174_v24)
			var _2176_v26 T0
			_ = _2176_v26
			var _nw365 *C0 = New_C0_()
			_ = _nw365
			_nw365.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC7_(!((_this).F24()), (_this).F24())).Dtor_cf14(), p2)).Merge((func() _dafny.Map {
				if (_2175_v25).Contains(_2173_v23) {
					return (_2175_v25).Get(_2173_v23).(_dafny.Map)
				}
				return _2174_v24
			})()), p3, _this.F23(), p0)
			_2176_v26 = _nw365
			_2176_v26 = _2176_v26
			var _2177_v27 _dafny.Set
			_ = _2177_v27
			_2177_v27 = _dafny.SetOf((_this).F24())
			var _2178_v28 _dafny.Set
			_ = _2178_v28
			_2178_v28 = _dafny.SetOf((_dafny.Zero).Minus(_2173_v23), _2173_v23)
			var _2179_v29 _dafny.Set
			_ = _2179_v29
			_2179_v29 = _dafny.SetOf(_2178_v28, _2178_v28)
			(globalState).F7 = Companion_Default___.Fm0((_2179_v29).IsSubsetOf(Companion_Default___.Fm54(_2177_v27, globalState)), (_2173_v23).Minus(_2173_v23), globalState)
			var _2180_v30 _dafny.Array
			_ = _2180_v30
			var _nwElement0_65 bool = p0
			_ = _nwElement0_65
			var _nw366 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(23))
			_ = _nw366
			(_nw366).ArraySet1(_nwElement0_65, 0)
			(_nw366).ArraySet1((_this).Fm11(_2173_v23, (_2176_v26).F24(), globalState), 1)
			(_nw366).ArraySet1(p3, 2)
			(_nw366).ArraySet1((_2176_v26).F24(), 3)
			(_nw366).ArraySet1(p3, 4)
			(_nw366).ArraySet1(p0, 5)
			(_nw366).ArraySet1(p3, 6)
			(_nw366).ArraySet1((_this).F24(), 7)
			(_nw366).ArraySet1((_this).F24(), 8)
			(_nw366).ArraySet1(p3, 9)
			(_nw366).ArraySet1((_2176_v26).F24(), 10)
			(_nw366).ArraySet1((_this).F24(), 11)
			(_nw366).ArraySet1(p0, 12)
			(_nw366).ArraySet1(p3, 13)
			(_nw366).ArraySet1((_this).F24(), 14)
			(_nw366).ArraySet1(p0, 15)
			(_nw366).ArraySet1(p3, 16)
			(_nw366).ArraySet1(false, 17)
			(_nw366).ArraySet1((_this).F24(), 18)
			(_nw366).ArraySet1((_this).F24(), 19)
			(_nw366).ArraySet1((_2176_v26).F24(), 20)
			(_nw366).ArraySet1(p3, 21)
			(_nw366).ArraySet1(p0, 22)
			_2180_v30 = _nw366
			var _2181_v31 _dafny.Sequence
			_ = _2181_v31
			_2181_v31 = _dafny.SeqOf(_2180_v30, _2180_v30)
			var _2182_v32 _dafny.Array
			_ = _2182_v32
			var _nwElement0_66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2181_v31, _2181_v31)
			_ = _nwElement0_66
			var _nw367 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.One)
			_ = _nw367
			(_nw367).ArraySet1(_nwElement0_66, 0)
			_2182_v32 = _nw367
			var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_2182_v32), 0))
			_ = _index301
			(_2182_v32).ArraySet1(_2181_v31, (_index301).Int())
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_2182_v32), 0))
			_ = _index302
			(_2182_v32).ArraySet1(_dafny.SeqOf(_2180_v30), (_index302).Int())
			(globalState).F7 = (_2173_v23).Times(_2173_v23)
			var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_2180_v30), 0))
			_ = _index303
			(_2180_v30).ArraySet1((_this).F24(), (_index303).Int())
			var _2183_v33 _dafny.Sequence
			_ = _2183_v33
			_2183_v33 = _dafny.SeqOf(p3, (_this).F24(), p0, p3, (_2176_v26).F24())
			var _2184_v34 _dafny.Sequence
			_ = _2184_v34
			_2184_v34 = _dafny.SeqOf(_2183_v33)
			var _2185_v35 _dafny.Map
			_ = _2185_v35
			_2185_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2173_v23, _2173_v23)
			var _2186_v36 _dafny.Sequence
			_ = _2186_v36
			_2186_v36 = _dafny.SeqOf(_2173_v23, (_this).Fm6(globalState))
			var _2187_v37 *C1
			_ = _2187_v37
			var _nw368 *C1 = New_C1_()
			_ = _nw368
			_nw368.Ctor__(p3, p0)
			_2187_v37 = _nw368
			var _2188_v38 _dafny.Map
			_ = _2188_v38
			_2188_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2187_v37, _2173_v23)
			var _2189_v39 _dafny.Map
			_ = _2189_v39
			_2189_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2173_v23, _2188_v38)
			var _2190_v40 _dafny.Map
			_ = _2190_v40
			_2190_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2173_v23, false)
			var _2191_v41 D19
			_ = _2191_v41
			_2191_v41 = Companion_D19_.Create_DC45_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2187_v37, _2173_v23))
			var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_2180_v30), 0))
			_ = _index304
			var _rhs325 bool = (_2185_v35).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F24()), _2183_v33), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F24()), _2183_v33)).Cardinality()))).Uint32(), (_this).F24()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2184_v34).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F24()), _2183_v33), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F24()), _2183_v33)).Cardinality()))).Uint32(), (_this).F24())).Cardinality()))).Uint32(), p0)).Cardinality())))
			_ = _rhs325
			var _rhs326 _dafny.Int = (func() _dafny.Int {
				if !(p0) || ((_this).F24()) {
					return _dafny.IntOfUint32((_2186_v36).Cardinality())
				}
				return _2173_v23
			})()
			_ = _rhs326
			var _rhs327 bool = (_dafny.IntOfInt64(-491)).Cmp(((_2189_v39).Update((_2190_v40).Cardinality(), (_2191_v41).Dtor_cf74())).Cardinality()) > 0
			_ = _rhs327
			var _rhs328 bool = (func() bool {
				if (_this).F24() {
					return (_2187_v37).F39()
				}
				return (_this).Fm4(_2173_v23, _2183_v33, globalState)
			})()
			_ = _rhs328
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			var _lhs244 _dafny.Array = _2180_v30
			_ = _lhs244
			var _lhs245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_2180_v30), 0))
			_ = _lhs245
			r0 = _rhs325
			_2173_v23 = _rhs326
			_lhs243.F13 = _rhs327
			(_lhs244).ArraySet1(_rhs328, (_lhs245).Int())
		}
		var _2192_v42 _dafny.Int
		_ = _2192_v42
		_2192_v42 = _dafny.IntOfInt64(458)
		var _2193_v43 _dafny.Map
		_ = _2193_v43
		_2193_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2192_v42, p0)
		if !((func() bool {
			if (_2193_v43).Contains(_2192_v42) {
				return (_2193_v43).Get(_2192_v42).(bool)
			}
			return p3
		})()) {
			var _2194_v44 _dafny.Sequence
			_ = _2194_v44
			_2194_v44 = _dafny.SeqOf((_2192_v42).Times(_2192_v42))
			var _rhs329 _dafny.CodePoint = _this.F23()
			_ = _rhs329
			var _rhs330 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2194_v44, _2194_v44), _2194_v44)
			_ = _rhs330
			var _lhs246 *C9 = _this
			_ = _lhs246
			_lhs246.F23_set_(_rhs329)
			_2194_v44 = _rhs330
			var _2195_v45 *C8
			_ = _2195_v45
			var _nw369 *C8 = New_C8_()
			_ = _nw369
			_nw369.Ctor__(p3)
			_2195_v45 = _nw369
			var _2196_v46 _dafny.Sequence
			_ = _2196_v46
			_2196_v46 = _dafny.SeqOf(p0, (_this).F24())
			var _2197_v47 _dafny.Set
			_ = _2197_v47
			_2197_v47 = _dafny.SetOf((_this).F24(), !((_2196_v46).Select((Companion_Default___.SafeIndex(_2192_v42, _dafny.IntOfUint32((_2196_v46).Cardinality()))).Uint32()).(bool)), (_this).F24(), p3, (_this).F24())
			var _2198_v48 _dafny.Map
			_ = _2198_v48
			_2198_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2197_v47, _dafny.MultiSetFromSeq(_2194_v44))
			var _2199_v49 D15
			_ = _2199_v49
			_2199_v49 = Companion_D15_.Create_DC35_(_2192_v42, _2192_v42, (_2197_v47).Cardinality())
			var _2200_v50 _dafny.Sequence
			_ = _2200_v50
			_2200_v50 = _dafny.SeqOf(_2199_v49, _2199_v49)
			_2198_v48 = (_2198_v48).Merge((func() _dafny.Map {
				if p3 {
					return _2198_v48
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2197_v47, Companion_Default___.Fm42(p3, _dafny.IntOfUint32((_2200_v50).Cardinality()), globalState))
			})())
			var _2201_v51 _dafny.Map
			_ = _2201_v51
			_2201_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
			var _2202_v52 _dafny.Sequence
			_ = _2202_v52
			_2202_v52 = _dafny.SeqOf(_2194_v44)
			var _2203_v53 _dafny.Map
			_ = _2203_v53
			_2203_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(399))
			var _2204_v54 _dafny.Set
			_ = _2204_v54
			_2204_v54 = _dafny.SetOf(_2192_v42, _dafny.IntOfUint32((_2202_v52).Cardinality()), (_2203_v53).Cardinality(), _2192_v42, _dafny.IntOfUint32((p2).Cardinality()))
			var _2205_v55 *C0
			_ = _2205_v55
			var _nw370 *C0 = New_C0_()
			_ = _nw370
			_nw370.Ctor__(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p2)).Update(Companion_Default___.Fm2(globalState), _dafny.UnicodeSeqOfUtf8Bytes("b"))).Merge(_2201_v51), (_2204_v54).IsSubsetOf(_dafny.SetOf((_this).Fm6(globalState))), _this.F23(), true)
			_2205_v55 = _nw370
			var _2206_v56 D15
			_ = _2206_v56
			_2206_v56 = Companion_D15_.Create_DC34_((_this).F24())
			var _source30 D15 = _2206_v56
			_ = _source30
			if _source30.Is_DC34() {
				var _2207___mcc_h0 bool = _source30.Get_().(D15_DC34).Cf55
				_ = _2207___mcc_h0
				var _2208_cf55 bool = _2207___mcc_h0
				_ = _2208_cf55
				(globalState).F7 = _2192_v42
				(globalState).F7 = _2192_v42
				var _2209_v57 _dafny.Array
				_ = _2209_v57
				var _nw371 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw371
				_2209_v57 = _nw371
				var _2210_v58 _dafny.Map
				_ = _2210_v58
				_2210_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2192_v42, _dafny.IntOfInt64(-824))
				var _2211_v59 _dafny.Map
				_ = _2211_v59
				_2211_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F23())
				var _2212_v60 _dafny.Sequence
				_ = _2212_v60
				_2212_v60 = _dafny.SeqOf(_2211_v59)
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_2209_v57), 0))
				_ = _index305
				(_2209_v57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm55(_2208_cf55, _2210_v58, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(404))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}((func(_2213_v42 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2214_i0 _dafny.Int) _dafny.Int {
						return _2213_v42
					}
				})(_2192_v42))), globalState), _2212_v60), (_index305).Int())
				var _2215_v61 _dafny.Array
				_ = _2215_v61
				var _nw372 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
				_ = _nw372
				_2215_v61 = _nw372
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_2215_v61), 0))
				_ = _index306
				(_2215_v61).ArraySet1(p2, (_index306).Int())
				var _2216_v62 _dafny.Array
				_ = _2216_v62
				var _len0_66 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_66
				var _nw373 _dafny.Array
				_ = _nw373
				if _len0_66.Cmp(_dafny.Zero) == 0 {
					_nw373 = _dafny.NewArray(_len0_66)
				} else {
					var _init66 func(_dafny.Int) _dafny.Int = (func(_2217_v42 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2218_i1 _dafny.Int) _dafny.Int {
							return (_2218_i1).Times(_2217_v42)
						}
					})(_2192_v42)
					_ = _init66
					var _element0_66 = _init66(_dafny.Zero)
					_ = _element0_66
					_nw373 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
					(_nw373).ArraySet1(_element0_66, 0)
					var _nativeLen0_66 = (_len0_66).Int()
					_ = _nativeLen0_66
					for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
						(_nw373).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
					}
				}
				_2216_v62 = _nw373
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2216_v62), 0))
				_ = _index307
				(_2216_v62).ArraySet1(((_this).Fm6(globalState)).Plus(_dafny.IntOfUint32((_2194_v44).Cardinality())), (_index307).Int())
				var _2219_v63 _dafny.MultiSet
				_ = _2219_v63
				_2219_v63 = _dafny.MultiSetOf(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(887))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}(func(_2220_i2 _dafny.Int) _dafny.CodePoint {
					return _this.F23()
				})), p2, p2, p2)
				var _2221_v64 _dafny.MultiSet
				_ = _2221_v64
				_2221_v64 = _2219_v63
				var _2222_v65 _dafny.Map
				_ = _2222_v65
				_2222_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _dafny.MultiSetOf(p2))
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_2209_v57), 0))
				_ = _index308
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_2215_v61), 0))
				_ = _index309
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2216_v62), 0))
				_ = _index310
				var _rhs331 _dafny.Sequence = _2212_v60
				_ = _rhs331
				var _rhs332 bool = !(p3) || ((_2221_v64).IsProperSubsetOf((func() _dafny.MultiSet {
					if (_2222_v65).Contains(_dafny.CodePoint('j')) {
						return (_2222_v65).Get(_dafny.CodePoint('j')).(_dafny.MultiSet)
					}
					return _2219_v63
				})()))
				_ = _rhs332
				var _rhs333 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p2, p2), _dafny.Companion_Sequence_.Concatenate(p2, p2))
				_ = _rhs333
				var _rhs334 _dafny.Int = _2192_v42
				_ = _rhs334
				var _lhs247 _dafny.Array = _2209_v57
				_ = _lhs247
				var _lhs248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_2209_v57), 0))
				_ = _lhs248
				var _lhs249 _dafny.Array = _2215_v61
				_ = _lhs249
				var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_2215_v61), 0))
				_ = _lhs250
				var _lhs251 _dafny.Array = _2216_v62
				_ = _lhs251
				var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2216_v62), 0))
				_ = _lhs252
				(_lhs247).ArraySet1(_rhs331, (_lhs248).Int())
				_2208_cf55 = _rhs332
				(_lhs249).ArraySet1(_rhs333, (_lhs250).Int())
				(_lhs251).ArraySet1(_rhs334, (_lhs252).Int())
				var _2223_v66 _dafny.Map
				_ = _2223_v66
				_2223_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2205_v55).Fm26((_2216_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2216_v62), 0))).Int()).(_dafny.Int), _2196_v46, _2204_v54, (_2205_v55).F38(), globalState), (_2216_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_2216_v62), 0))).Int()).(_dafny.Int))
				var _2224_v67 _dafny.MultiSet
				_ = _2224_v67
				_2224_v67 = _dafny.MultiSetOf((_2223_v66).Cardinality(), _2192_v42)
				(globalState).F13 = (_2224_v67).IsSubsetOf(_2224_v67)
			} else if _source30.Is_DC35() {
				var _2225___mcc_h1 _dafny.Int = _source30.Get_().(D15_DC35).Cf56
				_ = _2225___mcc_h1
				var _2226___mcc_h2 _dafny.Int = _source30.Get_().(D15_DC35).Cf57
				_ = _2226___mcc_h2
				var _2227___mcc_h3 _dafny.Int = _source30.Get_().(D15_DC35).Cf58
				_ = _2227___mcc_h3
				var _2228_cf58 _dafny.Int = _2227___mcc_h3
				_ = _2228_cf58
				var _2229_cf57 _dafny.Int = _2226___mcc_h2
				_ = _2229_cf57
				var _2230_cf56 _dafny.Int = _2225___mcc_h1
				_ = _2230_cf56
				var _2231_v68 _dafny.Sequence
				_ = _2231_v68
				_2231_v68 = _dafny.UnicodeSeqOfUtf8Bytes("uiky")
				var _2232_v69 D19
				_ = _2232_v69
				_2232_v69 = Companion_D19_.Create_DC46_(_2197_v47, p2)
				var _2233_v70 _dafny.Sequence
				_ = _2233_v70
				_2233_v70 = _dafny.SeqOf(_2231_v68, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_this).F24(), globalState), p2), p2, (_2232_v69).Dtor_cf76())
				_2231_v68 = (_2233_v70).Select((Companion_Default___.SafeIndex(_2229_cf57, _dafny.IntOfUint32((_2233_v70).Cardinality()))).Uint32()).(_dafny.Sequence)
				var _2234_v71 _dafny.MultiSet
				_ = _2234_v71
				_2234_v71 = _dafny.MultiSetOf((_2205_v55).F38(), (_this).F24())
				var _2235_v72 _dafny.Map
				_ = _2235_v72
				_2235_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2229_cf57, _2234_v71)
				_2193_v43 = (_2193_v43).Update(Companion_Default___.SafeDivisionInt(_2192_v42, ((func() _dafny.MultiSet {
					if (_2235_v72).Contains(_2228_cf58) {
						return (_2235_v72).Get(_2228_cf58).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf((_2205_v55).F38(), p3)
				})()).Cardinality()), !((_2205_v55).F38()) || (p3))
				var _2236_v73 _dafny.Map
				_ = _2236_v73
				_2236_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2192_v42, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(972))).Uint32(), func(coer112 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}((func(_2237_cf57 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2238_i3 _dafny.Int) _dafny.Int {
						return _2237_cf57
					}
				})(_2229_cf57))))
				_2236_v73 = (_2236_v73).Update((_dafny.IntOfUint32((_2231_v68).Cardinality())).Minus(_2192_v42), _2194_v44)
				var _2239_v74 _dafny.Map
				_ = _2239_v74
				_2239_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2194_v44, _this.F23())
				var _2240_v75 _dafny.Map
				_ = _2240_v75
				_2240_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2231_v68, _2239_v74)
				_2240_v75 = (_2240_v75).Update(p2, _2239_v74)
			} else if _source30.Is_DC36() {
				var _2241___mcc_h4 bool = _source30.Get_().(D15_DC36).Cf59
				_ = _2241___mcc_h4
				var _2242___mcc_h5 _dafny.Int = _source30.Get_().(D15_DC36).Cf60
				_ = _2242___mcc_h5
				var _2243_cf60 _dafny.Int = _2242___mcc_h5
				_ = _2243_cf60
				var _2244_cf59 bool = _2241___mcc_h4
				_ = _2244_cf59
				var _2245_v76 _dafny.Array
				_ = _2245_v76
				var _len0_67 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_67
				var _nw374 _dafny.Array
				_ = _nw374
				if _len0_67.Cmp(_dafny.Zero) == 0 {
					_nw374 = _dafny.NewArray(_len0_67)
				} else {
					var _init67 func(_dafny.Int) _dafny.CodePoint = func(_2246_i4 _dafny.Int) _dafny.CodePoint {
						return _this.F23()
					}
					_ = _init67
					var _element0_67 = _init67(_dafny.Zero)
					_ = _element0_67
					_nw374 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
					(_nw374).ArraySet1CodePoint(_element0_67, 0)
					var _nativeLen0_67 = (_len0_67).Int()
					_ = _nativeLen0_67
					for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
						(_nw374).ArraySet1CodePoint(_init67(_dafny.IntOf(_i0_67)), _i0_67)
					}
				}
				_2245_v76 = _nw374
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2245_v76), 0))
				_ = _index311
				(_2245_v76).ArraySet1CodePoint(_this.F23(), (_index311).Int())
				var _2247_v77 _dafny.Array
				_ = _2247_v77
				var _nw375 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw375
				_2247_v77 = _nw375
				var _2248_v78 _dafny.MultiSet
				_ = _2248_v78
				_2248_v78 = _dafny.MultiSetOf(_2247_v77, _2247_v77)
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2245_v76), 0))
				_ = _index312
				var _rhs335 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_2244_cf59, globalState), _dafny.Companion_Sequence_.Concatenate(p2, _dafny.UnicodeSeqOfUtf8Bytes("qyif")))).Cardinality())
				_ = _rhs335
				var _rhs336 _dafny.CodePoint = _this.F23()
				_ = _rhs336
				var _rhs337 _dafny.Int = (func() _dafny.Int {
					if (_2248_v78).Contains(_2247_v77) {
						return (_2248_v78).Multiplicity(_2247_v77)
					}
					return _2243_cf60
				})()
				_ = _rhs337
				var _rhs338 bool = p3
				_ = _rhs338
				var _lhs253 *GlobalState = globalState
				_ = _lhs253
				var _lhs254 _dafny.Array = _2245_v76
				_ = _lhs254
				var _lhs255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2245_v76), 0))
				_ = _lhs255
				var _lhs256 *GlobalState = globalState
				_ = _lhs256
				var _lhs257 *GlobalState = globalState
				_ = _lhs257
				_lhs253.F7 = _rhs335
				(_lhs254).ArraySet1CodePoint(_rhs336, (_lhs255).Int())
				_lhs256.F7 = _rhs337
				_lhs257.F13 = _rhs338
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2245_v76), 0))
				_ = _index313
				(_2245_v76).ArraySet1CodePoint(_this.F23(), (_index313).Int())
				_2243_cf60 = _2192_v42
				var _2249_v79 _dafny.Array
				_ = _2249_v79
				var _nw376 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw376
				_2249_v79 = _nw376
				var _2250_v80 _dafny.Set
				_ = _2250_v80
				_2250_v80 = _dafny.SetOf(_2249_v79, _2249_v79)
				(globalState).F7 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_2250_v80).Cardinality(), _2192_v42))))
			} else {
				var _2251___mcc_h6 _dafny.Map = _source30.Get_().(D15_DC33).Cf54
				_ = _2251___mcc_h6
				var _2252_cf54 _dafny.Map = _2251___mcc_h6
				_ = _2252_cf54
				var _2253_v81 _dafny.Array
				_ = _2253_v81
				var _nw377 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw377
				_2253_v81 = _nw377
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_2253_v81), 0))
				_ = _index314
				(_2253_v81).ArraySet1(_dafny.IntOfInt64(-378), (_index314).Int())
				var _2254_v82 _dafny.MultiSet
				_ = _2254_v82
				_2254_v82 = _dafny.MultiSetOf(_2203_v53)
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_2253_v81), 0))
				_ = _index315
				(_2253_v81).ArraySet1(((_2254_v82).Intersection((_2254_v82).Difference((_2254_v82).Update(_2203_v53, Companion_Default___.Abs(_2192_v42))))).Cardinality(), (_index315).Int())
				var _2255_v83 T0
				_ = _2255_v83
				var _nw378 *C0 = New_C0_()
				_ = _nw378
				_nw378.Ctor__(_2201_v51, (_2205_v55).F38(), _this.F23(), (_2205_v55).F38())
				_2255_v83 = _nw378
				var _2256_v84 T0
				_ = _2256_v84
				_2256_v84 = _2255_v83
				_2256_v84 = _2256_v84
				var _2257_v85 _dafny.Array
				_ = _2257_v85
				var _nw379 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
				_ = _nw379
				_2257_v85 = _nw379
				var _2258_v86 _dafny.Map
				_ = _2258_v86
				_2258_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2257_v85, p2)
				_2258_v86 = (_2258_v86).Update(_2257_v85, p2)
				(globalState).F13 = p0
			}
		} else {
			(globalState).F7 = (_2192_v42).Times((_2192_v42).Plus(_2192_v42))
			var _2259_v87 _dafny.Map
			_ = _2259_v87
			_2259_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F23())
			(globalState).F13 = !(!(!(_2259_v87).Equals(_2259_v87)))
			var _2260_v88 _dafny.Map
			_ = _2260_v88
			_2260_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
			var _2261_v89 *C0
			_ = _2261_v89
			var _nw380 *C0 = New_C0_()
			_ = _nw380
			_nw380.Ctor__(_2260_v88, p0, _this.F23(), p0)
			_2261_v89 = _nw380
			_2192_v42 = ((_this).Fm6(globalState)).Minus(_2192_v42)
			var _2262_v90 _dafny.Map
			_ = _2262_v90
			_2262_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), (_this).F24())).Cardinality()).Minus(_2192_v42)), Companion_Default___.Fm56(p3, _2192_v42, globalState))
			_2262_v90 = (func() _dafny.Map {
				var _coll94 = _dafny.NewMapBuilder()
				_ = _coll94
				for _iter99 := _dafny.Iterate((_dafny.MultiSetOf(_2192_v42)).Elements()); ; {
					_compr_94, _ok99 := _iter99()
					if !_ok99 {
						break
					}
					var _2263_v91 _dafny.Int
					_2263_v91 = interface{}(_compr_94).(_dafny.Int)
					if (_dafny.MultiSetOf(_2192_v42)).Contains(_2263_v91) {
						_coll94.Add(Companion_Default___.SafeModuloInt(_2263_v91, _2192_v42), _2193_v43)
					}
				}
				return _coll94.ToMap()
			}()).Merge(_2262_v90)
		}
		var _2264_v92 _dafny.Sequence
		_ = _2264_v92
		_2264_v92 = _dafny.SeqOf(_dafny.IntOfUint32((p2).Cardinality()))
		var _2265_v93 _dafny.Set
		_ = _2265_v93
		_2265_v93 = _dafny.SetOf((_this).Fm6(globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2192_v42), _2264_v92)).Cardinality()))
		_2265_v93 = _2265_v93
		var _2266_v94 _dafny.Set
		_ = _2266_v94
		_2266_v94 = _dafny.SetOf(true)
		var _2267_v95 _dafny.MultiSet
		_ = _2267_v95
		_2267_v95 = _dafny.MultiSetOf(_2192_v42, (_dafny.Zero).Minus(_2192_v42))
		var _2268_v96 _dafny.Map
		_ = _2268_v96
		_2268_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_2266_v94, _dafny.SetOf(p0, p3, true))).Cardinality(), (_2267_v95).Cardinality())
		_2268_v96 = (_2268_v96).Merge(_2268_v96)
		var _2269_v97 D9
		_ = _2269_v97
		_2269_v97 = Companion_D9_.Create_DC20_(_dafny.UnicodeSeqOfUtf8Bytes("of"), (_this).F24(), _2192_v42)
		var _2270_v98 _dafny.MultiSet
		_ = _2270_v98
		_2270_v98 = _dafny.MultiSetOf(_2269_v97)
		var _2271_v99 D21
		_ = _2271_v99
		_2271_v99 = Companion_D21_.Create_DC50_(_2270_v98)
		_2270_v98 = (_2271_v99).Dtor_cf80()
		var _2272_i5 _dafny.Int
		_ = _2272_i5
		_2272_i5 = _dafny.Zero
		{
			for (_this).F24() {
				{
					if (_2272_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_2272_i5 = (_2272_i5).Plus(_dafny.One)
					_2193_v43 = (_2193_v43).Update((_dafny.Zero).Minus((_2192_v42).Plus(_2192_v42)), (_2192_v42).Cmp(_2192_v42) == 0)
					var _2273_v100 _dafny.MultiSet
					_ = _2273_v100
					_2273_v100 = _dafny.MultiSetOf(_2265_v93)
					if (_2273_v100).IsProperSubsetOf((func() _dafny.MultiSet {
						if (_this).F24() {
							return _2273_v100
						}
						return _2273_v100
					})()) {
						var _2274_v101 _dafny.Array
						_ = _2274_v101
						var _nw381 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
						_ = _nw381
						_2274_v101 = _nw381
						var _2275_v102 _dafny.Array
						_ = _2275_v102
						var _nwElement0_67 bool = p0
						_ = _nwElement0_67
						var _nw382 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(7))
						_ = _nw382
						(_nw382).ArraySet1(_nwElement0_67, 0)
						(_nw382).ArraySet1(Companion_Default___.Fm2(globalState), 1)
						(_nw382).ArraySet1(p0, 2)
						(_nw382).ArraySet1((_this).F24(), 3)
						(_nw382).ArraySet1(true, 4)
						(_nw382).ArraySet1(p0, 5)
						(_nw382).ArraySet1((_this).F24(), 6)
						_2275_v102 = _nw382
						var _2276_v103 _dafny.Array
						_ = _2276_v103
						var _nwElement0_68 _dafny.Array = _2275_v102
						_ = _nwElement0_68
						var _nw383 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(3))
						_ = _nw383
						(_nw383).ArraySet1(_nwElement0_68, 0)
						(_nw383).ArraySet1(_2275_v102, 1)
						(_nw383).ArraySet1(_2275_v102, 2)
						_2276_v103 = _nw383
						var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_2276_v103), 0))
						_ = _index316
						(_2276_v103).ArraySet1(_2275_v102, (_index316).Int())
						var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_2276_v103), 0))
						_ = _index317
						var _rhs339 _dafny.Array = _2274_v101
						_ = _rhs339
						var _rhs340 bool = (_this).F24()
						_ = _rhs340
						var _rhs341 bool = _dafny.Companion_Sequence_.IsPrefixOf(p2, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(729))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg113 _dafny.Int) interface{} {
								return coer113(arg113)
							}
						}((func(_2277_p2 _dafny.Sequence, _2278_v42 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
							return func(_2279_i6 _dafny.Int) _dafny.CodePoint {
								return (_2277_p2).Select((Companion_Default___.SafeIndex(_2278_v42, _dafny.IntOfUint32((_2277_p2).Cardinality()))).Uint32()).(_dafny.CodePoint)
							}
						})(p2, _2192_v42))), p2))
						_ = _rhs341
						var _rhs342 _dafny.Int = _2192_v42
						_ = _rhs342
						var _rhs343 _dafny.Array = _2275_v102
						_ = _rhs343
						var _lhs258 *GlobalState = globalState
						_ = _lhs258
						var _lhs259 *GlobalState = globalState
						_ = _lhs259
						var _lhs260 _dafny.Array = _2276_v103
						_ = _lhs260
						var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_2276_v103), 0))
						_ = _lhs261
						_2274_v101 = _rhs339
						_lhs258.F13 = _rhs340
						r0 = _rhs341
						_lhs259.F7 = _rhs342
						(_lhs260).ArraySet1(_rhs343, (_lhs261).Int())
						var _2280_v104 _dafny.Array
						_ = _2280_v104
						var _len0_68 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_68
						var _nw384 _dafny.Array
						_ = _nw384
						if _len0_68.Cmp(_dafny.Zero) == 0 {
							_nw384 = _dafny.NewArray(_len0_68)
						} else {
							var _init68 func(_dafny.Int) _dafny.CodePoint = func(_2281_i7 _dafny.Int) _dafny.CodePoint {
								return _this.F23()
							}
							_ = _init68
							var _element0_68 = _init68(_dafny.Zero)
							_ = _element0_68
							_nw384 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
							(_nw384).ArraySet1CodePoint(_element0_68, 0)
							var _nativeLen0_68 = (_len0_68).Int()
							_ = _nativeLen0_68
							for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
								(_nw384).ArraySet1CodePoint(_init68(_dafny.IntOf(_i0_68)), _i0_68)
							}
						}
						_2280_v104 = _nw384
						_2280_v104 = _2280_v104
						var _2282_v105 _dafny.MultiSet
						_ = _2282_v105
						_2282_v105 = _dafny.MultiSetOf(p0)
						var _2283_v106 *C2
						_ = _2283_v106
						var _nw385 *C2 = New_C2_()
						_ = _nw385
						_nw385.Ctor__(_2282_v105, _2280_v104, _this.F23(), p3)
						_2283_v106 = _nw385
						var _2284_v107 D5
						_ = _2284_v107
						_2284_v107 = Companion_D5_.Create_DC13_(_2192_v42, p3, p3)
						var _rhs344 _dafny.Int = (_2284_v107).Dtor_cf20()
						_ = _rhs344
						var _rhs345 _dafny.Int = _2192_v42
						_ = _rhs345
						var _rhs346 bool = !_dafny.Companion_Sequence_.Contains(p2, _this.F23())
						_ = _rhs346
						var _rhs347 _dafny.Int = (_dafny.Zero).Minus(_2192_v42)
						_ = _rhs347
						var _lhs262 *GlobalState = globalState
						_ = _lhs262
						var _lhs263 *GlobalState = globalState
						_ = _lhs263
						var _lhs264 *GlobalState = globalState
						_ = _lhs264
						_lhs262.F7 = _rhs344
						_2192_v42 = _rhs345
						_lhs263.F13 = _rhs346
						_lhs264.F7 = _rhs347
						r0 = !((_this).F24())
					} else {
						var _2285_v108 _dafny.Array
						_ = _2285_v108
						var _len0_69 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_69
						var _nw386 _dafny.Array
						_ = _nw386
						if _len0_69.Cmp(_dafny.Zero) == 0 {
							_nw386 = _dafny.NewArray(_len0_69)
						} else {
							var _init69 func(_dafny.Int) _dafny.MultiSet = (func(_2286_v95 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
								return func(_2287_i8 _dafny.Int) _dafny.MultiSet {
									return _2286_v95
								}
							})(_2267_v95)
							_ = _init69
							var _element0_69 = _init69(_dafny.Zero)
							_ = _element0_69
							_nw386 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
							(_nw386).ArraySet1(_element0_69, 0)
							var _nativeLen0_69 = (_len0_69).Int()
							_ = _nativeLen0_69
							for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
								(_nw386).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
							}
						}
						_2285_v108 = _nw386
						_2285_v108 = _2285_v108
						(globalState).F13 = (_2192_v42).Cmp(_2192_v42) != 0
						var _2288_v109 _dafny.Sequence
						_ = _2288_v109
						_2288_v109 = _dafny.UnicodeSeqOfUtf8Bytes("cjiuamnmc")
						_2288_v109 = _2288_v109
						var _2289_v110 D0
						_ = _2289_v110
						_2289_v110 = Companion_D0_.Create_DC1_((_2268_v96).Cardinality())
						var _2290_v111 _dafny.Sequence
						_ = _2290_v111
						_2290_v111 = _dafny.SeqOf(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(38)), _2289_v110, _2289_v110)
						var _2291_v112 _dafny.Map
						_ = _2291_v112
						_2291_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(339), _dafny.SetOf(_2192_v42))
						var _2292_v113 _dafny.Sequence
						_ = _2292_v113
						_2292_v113 = _dafny.SeqOf(true, p3)
						var _2293_v114 _dafny.Array
						_ = _2293_v114
						var _nwElement0_69 bool = p3
						_ = _nwElement0_69
						var _nw387 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(28))
						_ = _nw387
						(_nw387).ArraySet1(_nwElement0_69, 0)
						(_nw387).ArraySet1(p3, 1)
						(_nw387).ArraySet1(p3, 2)
						(_nw387).ArraySet1((_this).F24(), 3)
						(_nw387).ArraySet1(false, 4)
						(_nw387).ArraySet1(p0, 5)
						(_nw387).ArraySet1(p0, 6)
						(_nw387).ArraySet1(p3, 7)
						(_nw387).ArraySet1((_this).Fm4(_2192_v42, _2292_v113, globalState), 8)
						(_nw387).ArraySet1(Companion_Default___.Fm2(globalState), 9)
						(_nw387).ArraySet1(p0, 10)
						(_nw387).ArraySet1(false, 11)
						(_nw387).ArraySet1(p0, 12)
						(_nw387).ArraySet1(p3, 13)
						(_nw387).ArraySet1(p0, 14)
						(_nw387).ArraySet1(p3, 15)
						(_nw387).ArraySet1(false, 16)
						(_nw387).ArraySet1(p0, 17)
						(_nw387).ArraySet1(p0, 18)
						(_nw387).ArraySet1(!((_this).F24()), 19)
						(_nw387).ArraySet1(p0, 20)
						(_nw387).ArraySet1(p0, 21)
						(_nw387).ArraySet1(p3, 22)
						(_nw387).ArraySet1(p0, 23)
						(_nw387).ArraySet1(true, 24)
						(_nw387).ArraySet1(p3, 25)
						(_nw387).ArraySet1(p3, 26)
						(_nw387).ArraySet1(p0, 27)
						_2293_v114 = _nw387
						var _2294_v115 _dafny.Array
						_ = _2294_v115
						var _nw388 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
						_ = _nw388
						_2294_v115 = _nw388
						var _2295_v116 *C3
						_ = _2295_v116
						var _nw389 *C3 = New_C3_()
						_ = _nw389
						_nw389.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2293_v114, _this.F23()), _2294_v115, _dafny.CodePoint('v'), false)
						_2295_v116 = _nw389
						var _2296_v117 _dafny.Sequence
						_ = _2296_v117
						_2296_v117 = _dafny.SeqOf(_2295_v116, _2295_v116)
						var _2297_v118 bool
						_ = _2297_v118
						var _out64 bool
						_ = _out64
						_out64 = (_this).M5(_2290_v111, (_2265_v93).IsDisjointFrom((func() _dafny.Set {
							if (_2291_v112).Contains(_2192_v42) {
								return (_2291_v112).Get(_2192_v42).(_dafny.Set)
							}
							return _2265_v93
						})()), !(false), _dafny.IntOfUint32((_2296_v117).Cardinality()), globalState)
						_2297_v118 = _out64
						(globalState).F7 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_2288_v109).Cardinality()))
					}
					var _2298_v119 _dafny.Map
					_ = _2298_v119
					_2298_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _2299_v120 _dafny.Sequence
					_ = _2299_v120
					_2299_v120 = _dafny.SeqOf(p0)
					var _2300_v121 _dafny.Map
					_ = _2300_v121
					_2300_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2299_v120, _dafny.UnicodeSeqOfUtf8Bytes("ri"))
					var _2301_v122 _dafny.Map
					_ = _2301_v122
					_2301_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC7_(p3, p0)).Dtor_cf13(), (func() _dafny.Map {
						if (_this).Fm11((_2298_v119).Cardinality(), (_this).F24(), globalState) {
							return _2300_v121
						}
						return _2300_v121
					})())
					_2301_v122 = (_2301_v122).Update(!_dafny.Companion_Sequence_.Contains(p2, _this.F23()), _2300_v121)
					var _2302_v123 _dafny.Array
					_ = _2302_v123
					var _nw390 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(16))
					_ = _nw390
					_2302_v123 = _nw390
					var _2303_v124 D4
					_ = _2303_v124
					_2303_v124 = Companion_D4_.Create_DC11_(_2192_v42)
					var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_2302_v123), 0))
					_ = _index318
					(_2302_v123).ArraySet1(_2303_v124, (_index318).Int())
					var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_2302_v123), 0))
					_ = _index319
					(_2302_v123).ArraySet1(Companion_D4_.Create_DC11_(_2192_v42), (_index319).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _2304_v125 _dafny.Sequence
		_ = _2304_v125
		_2304_v125 = _dafny.SeqOf((_2265_v93).Contains(_2192_v42))
		r0 = (_2304_v125).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
			if (_this).F24() {
				return _2192_v42
			}
			return _2192_v42
		})()), _dafny.IntOfUint32((_2304_v125).Cardinality()))).Uint32()).(bool)
		return r0
	}
}

// End of class C9

// Definition of class C10
type C10 struct {
	_f23 _dafny.CodePoint
	_f24 bool
	_f25 _dafny.Array
	_f27 bool
	F29  _dafny.Int
	_f28 bool
}

func New_C10_() *C10 {
	_this := C10{}

	_this._f23 = _dafny.CodePoint('D')
	_this._f24 = false
	_this._f25 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f27 = false
	_this.F29 = _dafny.Zero
	_this._f28 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T3_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C10{}
var _ T2 = &C10{}
var _ T3 = &C10{}
var _ T0 = &C10{}
var _ _dafny.TraitOffspring = &C10{}

func (_this *C10) F23() _dafny.CodePoint {
	return _this._f23
}
func (_this *C10) F23_set_(value _dafny.CodePoint) {
	_this._f23 = value
}
func (_this *C10) F24() bool {
	return _this._f24
}
func (_this *C10) F25() _dafny.Array {
	return _this._f25
}
func (_this *C10) F27() bool {
	return _this._f27
}
func (_this *C10) Ctor__(f28 bool, f29 _dafny.Int, f23 _dafny.CodePoint, f24 bool, f25 _dafny.Array, f27 bool) {
	{
		(_this)._f28 = f28
		(_this).F29 = f29
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this)._f27 = f27
	}
}
func (_this *C10) Fm5(globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC3_((func() D0 {
			if (_this).F27() {
				return Companion_D0_.Create_DC0_(_dafny.MultiSetOf((_this).F28()))
			}
			return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(997))).Cardinality()), (_this).F28()))
		})())
	}
}
func (_this *C10) Fm6(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_this.F29).Plus(_this.F29))
	}
}
func (_this *C10) Fm4(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !((_this).F28())
	}
}
func (_this *C10) Fm7(globalState *GlobalState) _dafny.Sequence {
	{
		if (_this.F29).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_this).F27())).Cardinality())) >= 0 {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F29))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(874))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg114 _dafny.Int) interface{} {
					return coer114(arg114)
				}
			}(func(_2305_i0 _dafny.Int) _dafny.Int {
				return _this.F29
			})))
		} else {
			return _dafny.SeqOf(_this.F29)
		}
	}
}
func (_this *C10) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.UnicodeSeqOfUtf8Bytes("th"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.UnicodeSeqOfUtf8Bytes("jedxluog")))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _dafny.UnicodeSeqOfUtf8Bytes("afbr")))
	}
}
func (_this *C10) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _this.F29)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bfg")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _this.F23())).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _this.F29)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _this.F29)))).Cardinality()
	}
}
func (_this *C10) Fm10(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		return _this.F23()
	}
}
func (_this *C10) M1(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) (_dafny.Sequence, D0, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var r2 bool = false
		_ = r2
		var _2306_v0 _dafny.Array
		_ = _2306_v0
		var _len0_70 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_70
		var _nw391 _dafny.Array
		_ = _nw391
		if _len0_70.Cmp(_dafny.Zero) == 0 {
			_nw391 = _dafny.NewArray(_len0_70)
		} else {
			var _init70 func(_dafny.Int) _dafny.Sequence = (func(_2307_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_2308_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qccj"), (Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qccj")).Cardinality()))).Uint32(), _2307_p0)
				}
			})(p0)
			_ = _init70
			var _element0_70 = _init70(_dafny.Zero)
			_ = _element0_70
			_nw391 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
			(_nw391).ArraySet1(_element0_70, 0)
			var _nativeLen0_70 = (_len0_70).Int()
			_ = _nativeLen0_70
			for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
				(_nw391).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
			}
		}
		_2306_v0 = _nw391
		var _2309_v1 _dafny.Sequence
		_ = _2309_v1
		_2309_v1 = _dafny.UnicodeSeqOfUtf8Bytes("nbcmfn")
		var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))
		_ = _index320
		(_2306_v0).ArraySet1(_2309_v1, (_index320).Int())
		var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))
		_ = _index321
		(_2306_v0).ArraySet1(_dafny.Companion_Sequence_.Update(_2309_v1, (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_this.F29, (_dafny.Zero).Minus(_this.F29)), _dafny.IntOfUint32((_2309_v1).Cardinality()))).Uint32(), p0), (_index321).Int())
		var _2310_v2 *C1
		_ = _2310_v2
		var _nw392 *C1 = New_C1_()
		_ = _nw392
		_nw392.Ctor__((!((_this).F24())) && (p1), (_this).F28())
		_2310_v2 = _nw392
		var _2311_v3 _dafny.Sequence
		_ = _2311_v3
		_2311_v3 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_this).F28(), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(679))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg115 _dafny.Int) interface{} {
				return coer115(arg115)
			}
		}(func(_2312_i1 _dafny.Int) _dafny.CodePoint {
			return _this.F23()
		}))), _2309_v1, _2309_v1, _dafny.Companion_Sequence_.Concatenate(_2309_v1, _2309_v1), (_2306_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))).Int()).(_dafny.Sequence))
		_2311_v3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm21(_dafny.SetOf((_this).F27()), _dafny.MultiSetFromSeq((_this).Fm7(globalState)), globalState), _2311_v3), _2311_v3)
		var _2313_v4 _dafny.Array
		_ = _2313_v4
		var _nw393 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw393
		_2313_v4 = _nw393
		var _2314_v5 *C7
		_ = _2314_v5
		var _nw394 *C7 = New_C7_()
		_ = _nw394
		_nw394.Ctor__((_2310_v2).F39(), (_2310_v2).F39(), (_this).F25(), _this.F23(), !((_this).F27()), (_this).F24())
		_2314_v5 = _nw394
		var _2315_v6 _dafny.Array
		_ = _2315_v6
		var _nwElement0_70 *C7 = _2314_v5
		_ = _nwElement0_70
		var _nw395 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(14))
		_ = _nw395
		(_nw395).ArraySet1(_nwElement0_70, 0)
		(_nw395).ArraySet1(_2314_v5, 1)
		(_nw395).ArraySet1(_2314_v5, 2)
		(_nw395).ArraySet1(_2314_v5, 3)
		(_nw395).ArraySet1(_2314_v5, 4)
		(_nw395).ArraySet1(_2314_v5, 5)
		(_nw395).ArraySet1(_2314_v5, 6)
		(_nw395).ArraySet1(_2314_v5, 7)
		(_nw395).ArraySet1(_2314_v5, 8)
		(_nw395).ArraySet1(_2314_v5, 9)
		(_nw395).ArraySet1(_2314_v5, 10)
		(_nw395).ArraySet1(_2314_v5, 11)
		(_nw395).ArraySet1(_2314_v5, 12)
		(_nw395).ArraySet1(_2314_v5, 13)
		_2315_v6 = _nw395
		var _2316_v7 _dafny.Map
		_ = _2316_v7
		_2316_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _2315_v6)
		var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_2313_v4), 0))
		_ = _index322
		(_2313_v4).ArraySet1((func() _dafny.Array {
			if (_2316_v7).Contains(_this.F23()) {
				return (_2316_v7).Get(_this.F23()).(_dafny.Array)
			}
			return _2315_v6
		})(), (_index322).Int())
		var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_2313_v4), 0))
		_ = _index323
		var _nwElement0_71 *C7 = _2314_v5
		_ = _nwElement0_71
		var _nw396 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(4))
		_ = _nw396
		(_nw396).ArraySet1(_nwElement0_71, 0)
		(_nw396).ArraySet1(_2314_v5, 1)
		(_nw396).ArraySet1((func() *C7 {
			if (_2314_v5).F30() {
				return _2314_v5
			}
			return _2314_v5
		})(), 2)
		(_nw396).ArraySet1(_2314_v5, 3)
		(_2313_v4).ArraySet1(_nw396, (_index323).Int())
		var _ingredients1 = _dafny.NewBuilder()
		_ = _ingredients1
		for _iter100 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2306_v0), 0))); ; {
			_guard_loop_4, _ok100 := _iter100()
			if !_ok100 {
				break
			}
			var _2317_i2 _dafny.Int
			_2317_i2 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_2317_i2).Sign() != -1) && ((_2317_i2).Cmp(_dafny.ArrayLen((_2306_v0), 0)) < 0)) {
				_ingredients1.Add(_dafny.TupleOf(_2306_v0, (_2317_i2).Int(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_2306_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("rmtvniulv")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(310))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg116 _dafny.Int) interface{} {
						return coer116(arg116)
					}
				}(func(_2318_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				})))))
			}
		}
		for _iter101 := _dafny.Iterate(_ingredients1); ; {
			_tup1, _ok101 := _iter101()
			if !_ok101 {
				break
			}
			(_dafny.ArrayCastTo((*(_tup1.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup1.(_dafny.Tuple)).IndexInt(2)), (*(_tup1.(_dafny.Tuple)).IndexInt(1)).(int))
		}
		var _2319_v8 D9
		_ = _2319_v8
		_2319_v8 = Companion_D9_.Create_DC20_((_2306_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm2(globalState), _dafny.IntOfUint32((_2309_v1).Cardinality()))
		var _2320_v9 _dafny.MultiSet
		_ = _2320_v9
		_2320_v9 = _dafny.MultiSetOf(_2319_v8)
		var _2321_v10 D21
		_ = _2321_v10
		_2321_v10 = Companion_D21_.Create_DC50_((_2320_v9).Difference(_2320_v9))
		var _source31 D21 = _2321_v10
		_ = _source31
		if _source31.Is_DC51() {
			var _2322___mcc_h0 bool = _source31.Get_().(D21_DC51).Cf81
			_ = _2322___mcc_h0
			var _2323_cf81 bool = _2322___mcc_h0
			_ = _2323_cf81
			_2309_v1 = (func() _dafny.Sequence {
				if !(_2323_cf81) || ((_2314_v5).F30()) {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(43))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg117 _dafny.Int) interface{} {
							return coer117(arg117)
						}
					}((func(_2324_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2325_i4 _dafny.Int) _dafny.CodePoint {
							return _2324_p0
						}
					})(p0)))
				}
				return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_this).F24(), globalState), _2309_v1)
			})()
			var _2326_v11 bool
			_ = _2326_v11
			var _out65 bool
			_ = _out65
			_out65 = (_2314_v5).M4((_2314_v5).F30(), globalState)
			_2326_v11 = _out65
			var _2327_v12 _dafny.Sequence
			_ = _2327_v12
			_2327_v12 = _dafny.SeqOf((_2314_v5).F30(), !((_2314_v5).F30()), (_2319_v8).Dtor_cf32())
			_2326_v11 = (_dafny.IntOfUint32((_2327_v12).Cardinality())).Cmp(_this.F29) < 0
			var _2328_v13 _dafny.Sequence
			_ = _2328_v13
			_2328_v13 = _dafny.SeqOf(_this.F29)
			var _2329_v14 _dafny.Map
			_ = _2329_v14
			_2329_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(843), (_2328_v13).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2328_v13).Cardinality()))).Uint32()).(_dafny.Int))
			_2327_v12 = _dafny.Companion_Sequence_.Update(_2327_v12, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_2329_v14).Contains(_dafny.IntOfInt64(157)) {
					return (_2329_v14).Get(_dafny.IntOfInt64(157)).(_dafny.Int)
				}
				return _this.F29
			})(), _dafny.IntOfUint32((_2327_v12).Cardinality()))).Uint32(), _2326_v11)
		} else if _source31.Is_DC50() {
			var _2330___mcc_h1 _dafny.MultiSet = _source31.Get_().(D21_DC50).Cf80
			_ = _2330___mcc_h1
			var _2331_cf80 _dafny.MultiSet = _2330___mcc_h1
			_ = _2331_cf80
			var _2332_v15 *C9
			_ = _2332_v15
			var _nw397 *C9 = New_C9_()
			_ = _nw397
			_nw397.Ctor__(_this.F23(), (_this).F28())
			_2332_v15 = _nw397
			var _2333_v16 _dafny.Sequence
			_ = _2333_v16
			_2333_v16 = _dafny.SeqOf((_2314_v5).F30(), false)
			var _2334_v17 _dafny.Array
			_ = _2334_v17
			var _nwElement0_72 bool = (_this).F27()
			_ = _nwElement0_72
			var _nw398 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(26))
			_ = _nw398
			(_nw398).ArraySet1(_nwElement0_72, 0)
			(_nw398).ArraySet1((_2314_v5).F30(), 1)
			(_nw398).ArraySet1((_this).Fm4(_this.F29, _2333_v16, globalState), 2)
			(_nw398).ArraySet1(true, 3)
			(_nw398).ArraySet1((_this).F28(), 4)
			(_nw398).ArraySet1(p1, 5)
			(_nw398).ArraySet1((_2314_v5).F30(), 6)
			(_nw398).ArraySet1(true, 7)
			(_nw398).ArraySet1((_this).F28(), 8)
			(_nw398).ArraySet1((_this).F24(), 9)
			(_nw398).ArraySet1(false, 10)
			(_nw398).ArraySet1((_this).F28(), 11)
			(_nw398).ArraySet1((_2310_v2).F39(), 12)
			(_nw398).ArraySet1((_2310_v2).F39(), 13)
			(_nw398).ArraySet1(!(!((_2310_v2).F39())), 14)
			(_nw398).ArraySet1((_2314_v5).F30(), 15)
			(_nw398).ArraySet1(p1, 16)
			(_nw398).ArraySet1(!((_2310_v2).F39()), 17)
			(_nw398).ArraySet1(true, 18)
			(_nw398).ArraySet1((_2314_v5).F31(), 19)
			(_nw398).ArraySet1((_this).F27(), 20)
			(_nw398).ArraySet1(p1, 21)
			(_nw398).ArraySet1((_2314_v5).F31(), 22)
			(_nw398).ArraySet1((_2314_v5).Fm4(_this.F29, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(false, (_this).F24(), (_2310_v2).F39()), (Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_dafny.SeqOf(false, (_this).F24(), (_2310_v2).F39())).Cardinality()))).Uint32(), (_2310_v2).F39()), globalState), 23)
			(_nw398).ArraySet1((_2310_v2).F39(), 24)
			(_nw398).ArraySet1((_2314_v5).F31(), 25)
			_2334_v17 = _nw398
			var _2335_v18 *C5
			_ = _2335_v18
			var _nw399 *C5 = New_C5_()
			_ = _nw399
			_nw399.Ctor__(_2334_v17)
			_2335_v18 = _nw399
			if (_this.F29).Cmp(_this.F29) >= 0 {
				(globalState).F7 = _this.F29
				(globalState).F7 = _this.F29
				var _2336_v19 _dafny.MultiSet
				_ = _2336_v19
				_2336_v19 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_2306_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()), _this.F29, _this.F29, (_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F29)), _this.F29)
				var _2337_v20 D5
				_ = _2337_v20
				_2337_v20 = Companion_D5_.Create_DC13_((func() _dafny.Int {
					if (_2336_v19).Contains(_this.F29) {
						return (_2336_v19).Multiplicity(_this.F29)
					}
					return _this.F29
				})(), true, false)
				_2337_v20 = _2337_v20
				var _2338_v21 _dafny.Array
				_ = _2338_v21
				var _nw400 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw400
				_2338_v21 = _nw400
				(globalState).F22 = _2338_v21
				var _2339_v22 _dafny.Set
				_ = _2339_v22
				_2339_v22 = _dafny.SetOf(true, (_2314_v5).F31())
				var _2340_v23 _dafny.Map
				_ = _2340_v23
				_2340_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2339_v22).Intersection(_2339_v22), _2311_v3)
				_2340_v23 = (_2340_v23).Update((_2339_v22).Difference(_2339_v22), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("e")))
			} else {
				(globalState).F7 = (Companion_D5_.Create_DC13_(_this.F29, (_2310_v2).F39(), (_this).F24())).Dtor_cf20()
				(_this).F29 = _this.F29
				var _2341_v24 _dafny.Array
				_ = _2341_v24
				var _len0_71 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_71
				var _nw401 _dafny.Array
				_ = _nw401
				if _len0_71.Cmp(_dafny.Zero) == 0 {
					_nw401 = _dafny.NewArray(_len0_71)
				} else {
					var _init71 func(_dafny.Int) _dafny.Int = func(_2342_i5 _dafny.Int) _dafny.Int {
						return (_2342_i5).Times(_dafny.IntOfUint32((_dafny.SeqOf(_this.F29, _this.F29)).Cardinality()))
					}
					_ = _init71
					var _element0_71 = _init71(_dafny.Zero)
					_ = _element0_71
					_nw401 = _dafny.NewArrayFromExample(_element0_71, nil, _len0_71)
					(_nw401).ArraySet1(_element0_71, 0)
					var _nativeLen0_71 = (_len0_71).Int()
					_ = _nativeLen0_71
					for _i0_71 := 1; _i0_71 < _nativeLen0_71; _i0_71++ {
						(_nw401).ArraySet1(_init71(_dafny.IntOf(_i0_71)), _i0_71)
					}
				}
				_2341_v24 = _nw401
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_2341_v24), 0))
				_ = _index324
				(_2341_v24).ArraySet1(_this.F29, (_index324).Int())
				var _2343_v25 _dafny.Sequence
				_ = _2343_v25
				_2343_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2333_v16, _2333_v16)).Cardinality()), _this.F29)
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_2341_v24), 0))
				_ = _index325
				var _rhs348 _dafny.Int = (_dafny.Zero).Minus(_this.F29)
				_ = _rhs348
				var _rhs349 _dafny.CodePoint = _this.F23()
				_ = _rhs349
				var _rhs350 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F29), (Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_dafny.SeqOf(_this.F29)).Cardinality()))).Uint32(), _this.F29), _2343_v25), _dafny.Companion_Sequence_.Concatenate(_2343_v25, _2343_v25))
				_ = _rhs350
				var _rhs351 _dafny.Int = (_dafny.Zero).Minus(_this.F29)
				_ = _rhs351
				var _lhs265 _dafny.Array = _2341_v24
				_ = _lhs265
				var _lhs266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_2341_v24), 0))
				_ = _lhs266
				var _lhs267 *C10 = _this
				_ = _lhs267
				var _lhs268 *GlobalState = globalState
				_ = _lhs268
				(_lhs265).ArraySet1(_rhs348, (_lhs266).Int())
				_lhs267.F23_set_(_rhs349)
				_2343_v25 = _rhs350
				_lhs268.F7 = _rhs351
				var _2344_v26 D21
				_ = _2344_v26
				_2344_v26 = Companion_D21_.Create_DC51_(true)
				var _2345_v27 _dafny.Map
				_ = _2345_v27
				_2345_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2344_v26).Dtor_cf81(), (_2314_v5).F31())
				_2345_v27 = (_2345_v27).Update((_2314_v5).F30(), (_2333_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2309_v1).Cardinality()), _dafny.IntOfUint32((_2333_v16).Cardinality()))).Uint32()).(bool))
				(globalState).F13 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2309_v1, (_2306_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))).Int()).(_dafny.Sequence)), _2309_v1)
			}
			var _2346_v28 *C1
			_ = _2346_v28
			var _nw402 *C1 = New_C1_()
			_ = _nw402
			_nw402.Ctor__((_2310_v2).F39(), (_2310_v2).F39())
			_2346_v28 = _nw402
		} else {
			var _2347___mcc_h2 D21 = _source31.Get_().(D21_DC52).Cf82
			_ = _2347___mcc_h2
			var _2348_cf82 D21 = _2347___mcc_h2
			_ = _2348_cf82
			(_this).F29 = _this.F29
			r2 = (_2314_v5).F31()
			var _2349_v29 _dafny.Array
			_ = _2349_v29
			var _nw403 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw403
			_2349_v29 = _nw403
			var _2350_v30 *C5
			_ = _2350_v30
			var _nw404 *C5 = New_C5_()
			_ = _nw404
			_nw404.Ctor__(_2349_v29)
			_2350_v30 = _nw404
			var _2351_v31 D16
			_ = _2351_v31
			_2351_v31 = Companion_D16_.Create_DC37_(_2350_v30)
			var _2352_v32 D16
			_ = _2352_v32
			_2352_v32 = Companion_D16_.Create_DC39_(_2351_v31)
			var _pat_let_tv45 = _2351_v31
			_ = _pat_let_tv45
			var _2353_v33 _dafny.Array
			_ = _2353_v33
			var _nwElement0_73 D16 = _2352_v32
			_ = _nwElement0_73
			var _nw405 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(17))
			_ = _nw405
			(_nw405).ArraySet1(_nwElement0_73, 0)
			(_nw405).ArraySet1(_2352_v32, 1)
			(_nw405).ArraySet1(_2352_v32, 2)
			(_nw405).ArraySet1(_2352_v32, 3)
			(_nw405).ArraySet1(_2352_v32, 4)
			(_nw405).ArraySet1(_2352_v32, 5)
			(_nw405).ArraySet1(Companion_D16_.Create_DC39_(_2351_v31), 6)
			(_nw405).ArraySet1((func() D16 {
				if (_2314_v5).F31() {
					return _2352_v32
				}
				return _2352_v32
			})(), 7)
			(_nw405).ArraySet1(func(_pat_let57_0 D16) D16 {
				return func(_2354_dt__update__tmp_h0 D16) D16 {
					return func(_pat_let58_0 D16) D16 {
						return func(_2355_dt__update_hcf65_h0 D16) D16 {
							return Companion_D16_.Create_DC39_(_2355_dt__update_hcf65_h0)
						}(_pat_let58_0)
					}(_pat_let_tv45)
				}(_pat_let57_0)
			}(_2352_v32), 8)
			(_nw405).ArraySet1(Companion_D16_.Create_DC39_(_2351_v31), 9)
			(_nw405).ArraySet1(_2352_v32, 10)
			(_nw405).ArraySet1(_2352_v32, 11)
			(_nw405).ArraySet1(_2352_v32, 12)
			(_nw405).ArraySet1(Companion_D16_.Create_DC39_(_2351_v31), 13)
			(_nw405).ArraySet1(_2352_v32, 14)
			(_nw405).ArraySet1(_2352_v32, 15)
			(_nw405).ArraySet1(_2352_v32, 16)
			_2353_v33 = _nw405
			var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_2353_v33), 0))
			_ = _index326
			(_2353_v33).ArraySet1(Companion_D16_.Create_DC39_(_2351_v31), (_index326).Int())
			var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_2353_v33), 0))
			_ = _index327
			(_2353_v33).ArraySet1(_2352_v32, (_index327).Int())
			var _2356_v34 D21
			_ = _2356_v34
			_2356_v34 = Companion_D21_.Create_DC51_((_2314_v5).F30())
			_2356_v34 = _2356_v34
		}
		r0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-100))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg118 _dafny.Int) interface{} {
				return coer118(arg118)
			}
		}((func(_2357_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2358_i6 _dafny.Int) _dafny.CodePoint {
				return _2357_p0
			}
		})(p0)))
		var _2359_v35 D0
		_ = _2359_v35
		_2359_v35 = Companion_D0_.Create_DC1_(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_2306_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_2306_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()), _this.F29))
		r1 = _2359_v35
		var _2360_v36 _dafny.MultiSet
		_ = _2360_v36
		_2360_v36 = _dafny.MultiSetOf(_this.F29)
		r2 = (((func() _dafny.MultiSet {
			if (_this).F24() {
				return _2360_v36
			}
			return _2360_v36
		})()).Cardinality()).Cmp(_this.F29) > 0
		return r0, r1, r2
	}
}
func (_this *C10) M2(globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		(globalState).F13 = (_this).F24()
		(globalState).F7 = Companion_Default___.SafeModuloInt(_this.F29, _this.F29)
		var _2361_v0 _dafny.Array
		_ = _2361_v0
		var _nw406 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(22))
		_ = _nw406
		_2361_v0 = _nw406
		var _2362_v1 _dafny.Map
		_ = _2362_v1
		_2362_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _dafny.IntOfInt64(-956))
		var _2363_v2 D0
		_ = _2363_v2
		_2363_v2 = Companion_D0_.Create_DC1_((_2362_v1).Cardinality())
		var _2364_v3 D0
		_ = _2364_v3
		_2364_v3 = Companion_D0_.Create_DC3_(_2363_v2)
		var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_2361_v0), 0))
		_ = _index328
		(_2361_v0).ArraySet1(_2364_v3, (_index328).Int())
		var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_2361_v0), 0))
		_ = _index329
		(_2361_v0).ArraySet1(_2364_v3, (_index329).Int())
		var _2365_v4 _dafny.Array
		_ = _2365_v4
		var _nw407 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw407
		_2365_v4 = _nw407
		var _2366_v5 _dafny.MultiSet
		_ = _2366_v5
		_2366_v5 = _dafny.MultiSetOf(_2365_v4)
		(globalState).F7 = (_this.F29).Minus((_2366_v5).Cardinality())
		var _2367_v6 _dafny.Sequence
		_ = _2367_v6
		_2367_v6 = _dafny.UnicodeSeqOfUtf8Bytes("i")
		var _2368_v7 _dafny.MultiSet
		_ = _2368_v7
		_2368_v7 = _dafny.MultiSetOf(_2367_v6, _2367_v6)
		var _2369_v8 _dafny.MultiSet
		_ = _2369_v8
		_2369_v8 = (_2368_v7).Difference(_dafny.MultiSetOf(_2367_v6, _2367_v6))
		var _source32 _dafny.MultiSet = _2369_v8
		_ = _source32
		var _2370___mcc_h0 _dafny.MultiSet = _source32
		_ = _2370___mcc_h0
		var _2371_cf79 _dafny.MultiSet = _2370___mcc_h0
		_ = _2371_cf79
		var _2372_v9 _dafny.Map
		_ = _2372_v9
		_2372_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(865))
		var _2373_v10 _dafny.MultiSet
		_ = _2373_v10
		_2373_v10 = _dafny.MultiSetOf((_2372_v9).Cardinality(), _this.F29)
		var _2374_v11 _dafny.Map
		_ = _2374_v11
		_2374_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_2373_v10).IsProperSubsetOf(_2373_v10))
		var _2375_v12 _dafny.Array
		_ = _2375_v12
		var _nw408 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw408
		_2375_v12 = _nw408
		var _2376_v13 *C5
		_ = _2376_v13
		var _nw409 *C5 = New_C5_()
		_ = _nw409
		_nw409.Ctor__(_2375_v12)
		_2376_v13 = _nw409
		var _2377_v14 _dafny.Set
		_ = _2377_v14
		_2377_v14 = _dafny.SetOf(_2376_v13, _2376_v13, _2376_v13, _2376_v13, _2376_v13)
		_2374_v11 = (_2374_v11).Update(Companion_Default___.Fm2(globalState), (_2377_v14).IsDisjointFrom(_dafny.SetOf(_2376_v13)))
		var _2378_v15 _dafny.Set
		_ = _2378_v15
		_2378_v15 = _dafny.SetOf(_2375_v12, (_2376_v13).F32(), _2375_v12)
		(globalState).F13 = (_2378_v15).IsProperSubsetOf(_2378_v15)
		(globalState).F13 = (_this).F28()
		var _2379_v16 _dafny.MultiSet
		_ = _2379_v16
		_2379_v16 = _dafny.MultiSetOf((_this).F27())
		var _2380_v17 _dafny.Map
		_ = _2380_v17
		_2380_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F25())
		var _2381_v18 *C2
		_ = _2381_v18
		var _nw410 *C2 = New_C2_()
		_ = _nw410
		_nw410.Ctor__(_2379_v16, (func() _dafny.Array {
			if (_2380_v17).Contains((_this).F27()) {
				return (_2380_v17).Get((_this).F27()).(_dafny.Array)
			}
			return (_this).F25()
		})(), _this.F23(), (_this).F24())
		_2381_v18 = _nw410
		var _hi14 _dafny.Int = _this.F29
		_ = _hi14
		for _2382_i0 := (_dafny.Zero).Minus(_this.F29); _2382_i0.Cmp(_hi14) < 0; _2382_i0 = _2382_i0.Plus(_dafny.One) {
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))
			_ = _index330
			(_2365_v4).ArraySet1(_2382_i0, (_index330).Int())
			var _2383_v19 _dafny.Sequence
			_ = _2383_v19
			_2383_v19 = _dafny.SeqOf(_this.F29)
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))
			_ = _index331
			(_2365_v4).ArraySet1((Companion_Default___.SafeDivisionInt(_this.F29, (_2383_v19).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2383_v19).Cardinality()))).Uint32()).(_dafny.Int))).Times((_2382_i0).Minus(_2382_i0)), (_index331).Int())
			var _2384_v20 D12
			_ = _2384_v20
			_2384_v20 = Companion_D12_.Create_DC25_((_this).F24(), (_2367_v6).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2367_v6).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _2385_v21 D12
			_ = _2385_v21
			_2385_v21 = Companion_D12_.Create_DC26_(_2384_v20)
			var _2386_v22 D12
			_ = _2386_v22
			_2386_v22 = Companion_D12_.Create_DC26_(_2384_v20)
			var _source33 D12 = _2386_v22
			_ = _source33
			if _source33.Is_DC25() {
				var _2387___mcc_h1 bool = _source33.Get_().(D12_DC25).Cf41
				_ = _2387___mcc_h1
				var _2388___mcc_h2 _dafny.CodePoint = _source33.Get_().(D12_DC25).Cf42
				_ = _2388___mcc_h2
				var _2389_cf42 _dafny.CodePoint = _2388___mcc_h2
				_ = _2389_cf42
				var _2390_cf41 bool = _2387___mcc_h1
				_ = _2390_cf41
				var _2391_v23 _dafny.MultiSet
				_ = _2391_v23
				_2391_v23 = _dafny.MultiSetOf((_this).F28(), (_this).F24(), (_this).F27())
				var _2392_v24 _dafny.Int
				_ = _2392_v24
				var _2393_v25 bool
				_ = _2393_v25
				var _2394_v26 _dafny.Int
				_ = _2394_v26
				var _out66 _dafny.Int
				_ = _out66
				var _out67 bool
				_ = _out67
				var _out68 _dafny.Int
				_ = _out68
				_out66, _out67, _out68 = Companion_Default___.M0(_dafny.IntOfUint32((_2367_v6).Cardinality()), _2391_v23, _2390_cf41, ((_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int)).Plus((_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int)), globalState)
				_2392_v24 = _out66
				_2393_v25 = _out67
				_2394_v26 = _out68
				r2 = (_this).F28()
				var _2395_v27 D0
				_ = _2395_v27
				_2395_v27 = Companion_D0_.Create_DC2_((_this).F24(), (_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int), (_this).F24())
				_2390_cf41 = (_2395_v27).Dtor_cf4()
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen(((_this).F25()), 0))
				_ = _index332
				((_this).F25()).ArraySet1CodePoint(_2389_cf42, (_index332).Int())
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen(((_this).F25()), 0))
				_ = _index333
				((_this).F25()).ArraySet1CodePoint(_this.F23(), (_index333).Int())
			} else if _source33.Is_DC24() {
				var _2396___mcc_h3 _dafny.Array = _source33.Get_().(D12_DC24).Cf40
				_ = _2396___mcc_h3
				var _2397_cf40 _dafny.Array = _2396___mcc_h3
				_ = _2397_cf40
				var _2398_v28 _dafny.Array
				_ = _2398_v28
				var _nw411 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
				_ = _nw411
				_2398_v28 = _nw411
				var _2399_v29 _dafny.Array
				_ = _2399_v29
				var _nw412 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw412
				_2399_v29 = _nw412
				var _2400_v30 D12
				_ = _2400_v30
				_2400_v30 = Companion_D12_.Create_DC25_((_this).F24(), _this.F23())
				var _2401_v31 *C3
				_ = _2401_v31
				var _nw413 *C3 = New_C3_()
				_ = _nw413
				_nw413.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2399_v29, (_2400_v30).Dtor_cf42()), _2397_cf40, _this.F23(), (_this).F24())
				_2401_v31 = _nw413
				var _2402_v32 _dafny.Map
				_ = _2402_v32
				_2402_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int), _2401_v31)
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_2398_v28), 0))
				_ = _index334
				(_2398_v28).ArraySet1(_2402_v32, (_index334).Int())
				var _2403_v34 _dafny.Set
				_ = _2403_v34
				_2403_v34 = _dafny.SetOf((_this).Fm9(_2383_v19, _2382_i0, (_this).F24(), globalState), _2382_i0)
				var _2404_v35 _dafny.Set
				_ = _2404_v35
				_2404_v35 = _dafny.SetOf(!((_this).F24()) || (Companion_Default___.Fm2(globalState)), (_2403_v34).IsSubsetOf(func() _dafny.Set {
					var _coll95 = _dafny.NewBuilder()
					_ = _coll95
					for _iter102 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-657), _dafny.IntOfInt64(218))); ; {
						_compr_95, _ok102 := _iter102()
						if !_ok102 {
							break
						}
						var _2405_v33 _dafny.Int
						_2405_v33 = interface{}(_compr_95).(_dafny.Int)
						if ((_dafny.IntOfInt64(-657)).Cmp(_2405_v33) <= 0) && ((_2405_v33).Cmp(_dafny.IntOfInt64(218)) < 0) {
							_coll95.Add((_2405_v33).Minus(_this.F29))
						}
					}
					return _coll95.ToSet()
				}()))
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))
				_ = _index335
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_2398_v28), 0))
				_ = _index336
				var _rhs352 _dafny.Int = (_2404_v35).Cardinality()
				_ = _rhs352
				var _rhs353 _dafny.Map = _2402_v32
				_ = _rhs353
				var _lhs269 _dafny.Array = _2365_v4
				_ = _lhs269
				var _lhs270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))
				_ = _lhs270
				var _lhs271 _dafny.Array = _2398_v28
				_ = _lhs271
				var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_2398_v28), 0))
				_ = _lhs272
				(_lhs269).ArraySet1(_rhs352, (_lhs270).Int())
				(_lhs271).ArraySet1(_rhs353, (_lhs272).Int())
				(globalState).F7 = _dafny.IntOfUint32((_2367_v6).Cardinality())
				var _2406_v36 D15
				_ = _2406_v36
				_2406_v36 = Companion_D15_.Create_DC34_(Companion_Default___.Fm2(globalState))
				var _2407_v37 _dafny.Map
				_ = _2407_v37
				_2407_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, (_this).F24())
				var _2408_v38 _dafny.Map
				_ = _2408_v38
				_2408_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2406_v36, (_2407_v37).Cardinality())
				var _2409_v39 *C6
				_ = _2409_v39
				var _nw414 *C6 = New_C6_()
				_ = _nw414
				_nw414.Ctor__(_this.F23(), (_this).F24())
				_2409_v39 = _nw414
				var _2410_v40 _dafny.Map
				_ = _2410_v40
				_2410_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _2409_v39)
				var _2411_v41 _dafny.MultiSet
				_ = _2411_v41
				_2411_v41 = _dafny.MultiSetOf(_2382_i0)
				var _2412_v42 _dafny.Map
				_ = _2412_v42
				_2412_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2408_v38).Update(_2406_v36, (_dafny.MultiSetOf(_2410_v40, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _2409_v39), _2410_v40, _2410_v40, _2410_v40)).Cardinality()), _2411_v41)
				var _2413_v43 _dafny.Set
				_ = _2413_v43
				_2413_v43 = _dafny.SetOf(_2367_v6, _dafny.Companion_Sequence_.Update(_2367_v6, (Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2367_v6).Cardinality()))).Uint32(), _dafny.CodePoint('u')), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aeorvcm"), _2367_v6), _2367_v6)
				var _2414_v44 _dafny.Map
				_ = _2414_v44
				_2414_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _2399_v29)
				var _2415_v45 _dafny.Array
				_ = _2415_v45
				var _nw415 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw415
				_2415_v45 = _nw415
				var _2416_v46 _dafny.Sequence
				_ = _2416_v46
				_2416_v46 = _dafny.SeqOf(_2367_v6)
				var _rhs354 _dafny.Map = _2412_v42
				_ = _rhs354
				var _rhs355 bool = (_2414_v44).Contains(true)
				_ = _rhs355
				var _rhs356 _dafny.Array = _2415_v45
				_ = _rhs356
				var _rhs357 _dafny.Set = func() _dafny.Set {
					var _coll96 = _dafny.NewBuilder()
					_ = _coll96
					for _iter103 := _dafny.Iterate((_2416_v46).Elements()); ; {
						_compr_96, _ok103 := _iter103()
						if !_ok103 {
							break
						}
						var _2417_v47 _dafny.Sequence
						_2417_v47 = interface{}(_compr_96).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_2416_v46, _2417_v47) {
							_coll96.Add(_2417_v47)
						}
					}
					return _coll96.ToSet()
				}()
				_ = _rhs357
				var _lhs273 *GlobalState = globalState
				_ = _lhs273
				var _lhs274 *GlobalState = globalState
				_ = _lhs274
				_2412_v42 = _rhs354
				_lhs273.F13 = _rhs355
				_lhs274.F22 = _rhs356
				_2413_v43 = _rhs357
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_2399_v29), 0))
				_ = _index337
				(_2399_v29).ArraySet1((Companion_Default___.Fm0((_this).F28(), _this.F29, globalState)).Cmp(_this.F29) >= 0, (_index337).Int())
				var _2418_v48 _dafny.Sequence
				_ = _2418_v48
				_2418_v48 = _dafny.SeqOf(!((_dafny.IntOfInt64(-418)).Cmp(_dafny.IntOfUint32((_2367_v6).Cardinality())) != 0), (_this).F28(), true)
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_2399_v29), 0))
				_ = _index338
				(_2399_v29).ArraySet1(!((_2418_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-95), _dafny.IntOfUint32((_2418_v48).Cardinality()))).Uint32()).(bool)), (_index338).Int())
			} else {
				var _2419___mcc_h4 D12 = _source33.Get_().(D12_DC26).Cf43
				_ = _2419___mcc_h4
				var _2420_cf43 D12 = _2419___mcc_h4
				_ = _2420_cf43
				var _2421_v49 _dafny.Array
				_ = _2421_v49
				var _nw416 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw416
				_2421_v49 = _nw416
				_2421_v49 = _2421_v49
				r2 = (_this).F24()
				var _2422_v50 _dafny.MultiSet
				_ = _2422_v50
				_2422_v50 = _dafny.MultiSetOf(false, (_this).F24())
				var _2423_v51 _dafny.Int
				_ = _2423_v51
				var _2424_v52 bool
				_ = _2424_v52
				var _2425_v53 _dafny.Int
				_ = _2425_v53
				var _out69 _dafny.Int
				_ = _out69
				var _out70 bool
				_ = _out70
				var _out71 _dafny.Int
				_ = _out71
				_out69, _out70, _out71 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_2382_i0), _2382_i0), _2422_v50, (func() bool {
					if (_this).F28() {
						return (_this).F28()
					}
					return (_this).F28()
				})(), _this.F29, globalState)
				_2423_v51 = _out69
				_2424_v52 = _out70
				_2425_v53 = _out71
				var _2426_v54 T3
				_ = _2426_v54
				var _nw417 *C8 = New_C8_()
				_ = _nw417
				_nw417.Ctor__(_2424_v52)
				_2426_v54 = _nw417
				var _2427_v55 T3
				_ = _2427_v55
				_2427_v55 = _2426_v54
				var _2428_v56 _dafny.Map
				_ = _2428_v56
				_2428_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(298), (_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int))).Cardinality(), _2426_v54)
				var _2429_v57 _dafny.Map
				_ = _2429_v57
				_2429_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2367_v6)
				var _2430_v58 D2
				_ = _2430_v58
				_2430_v58 = Companion_D2_.Create_DC6_(_2429_v57)
				var _2431_v59 _dafny.MultiSet
				_ = _2431_v59
				_2431_v59 = _dafny.MultiSetOf(_2423_v51, _dafny.IntOfUint32((Companion_Default___.Fm45((_this).F27(), _2430_v58, (_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))
				var _2432_v60 _dafny.Array
				_ = _2432_v60
				var _nwElement0_74 T3 = _2426_v54
				_ = _nwElement0_74
				var _nw418 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(19))
				_ = _nw418
				(_nw418).ArraySet1(_nwElement0_74, 0)
				(_nw418).ArraySet1(_2426_v54, 1)
				(_nw418).ArraySet1(_2426_v54, 2)
				(_nw418).ArraySet1(_2426_v54, 3)
				(_nw418).ArraySet1(_2426_v54, 4)
				(_nw418).ArraySet1((_2427_v55), 5)
				(_nw418).ArraySet1(_2426_v54, 6)
				(_nw418).ArraySet1((_2427_v55), 7)
				(_nw418).ArraySet1(_2426_v54, 8)
				(_nw418).ArraySet1(_2426_v54, 9)
				(_nw418).ArraySet1(_2426_v54, 10)
				(_nw418).ArraySet1(_2426_v54, 11)
				(_nw418).ArraySet1(_2426_v54, 12)
				(_nw418).ArraySet1((func() T3 {
					if (_2428_v56).Contains((func() _dafny.Int {
						if (_2431_v59).Contains(_2425_v53) {
							return (_2431_v59).Multiplicity(_2425_v53)
						}
						return (_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int)
					})()) {
						return (_2428_v56).Get((func() _dafny.Int {
							if (_2431_v59).Contains(_2425_v53) {
								return (_2431_v59).Multiplicity(_2425_v53)
							}
							return (_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int)
						})()).(T3)
					}
					return _2426_v54
				})(), 13)
				(_nw418).ArraySet1((_2427_v55), 14)
				(_nw418).ArraySet1(_2426_v54, 15)
				(_nw418).ArraySet1(_2426_v54, 16)
				(_nw418).ArraySet1(_2426_v54, 17)
				(_nw418).ArraySet1(_2426_v54, 18)
				_2432_v60 = _nw418
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_2432_v60), 0))
				_ = _index339
				(_2432_v60).ArraySet1(_2426_v54, (_index339).Int())
				var _2433_v61 _dafny.Map
				_ = _2433_v61
				_2433_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2426_v54).F27(), false)
				var _2434_v62 _dafny.Map
				_ = _2434_v62
				_2434_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2433_v61, _2367_v6)
				var _2435_v63 _dafny.Sequence
				_ = _2435_v63
				_2435_v63 = _dafny.SeqOf(_2367_v6, (func() _dafny.Sequence {
					if (_2434_v62).Contains(_2433_v61) {
						return (_2434_v62).Get(_2433_v61).(_dafny.Sequence)
					}
					return _2367_v6
				})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}(func(_2436_i1 _dafny.Int) _dafny.CodePoint {
					return _this.F23()
				})))
				var _2437_v64 _dafny.MultiSet
				_ = _2437_v64
				_2437_v64 = _dafny.MultiSetOf(_2383_v19)
				var _2438_v65 _dafny.Sequence
				_ = _2438_v65
				_2438_v65 = _dafny.SeqOf(_2424_v52, (_this).F24(), (_this).F27(), false)
				var _2439_v67 _dafny.Map
				_ = _2439_v67
				_2439_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _this.F23())
				var _2440_v68 _dafny.Map
				_ = _2440_v68
				_2440_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2365_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_2365_v4), 0))).Int()).(_dafny.Int), (_2439_v67).Cardinality())
				var _2441_v69 _dafny.Map
				_ = _2441_v69
				_2441_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2440_v68).Cardinality(), (_this).F27())
				var _2442_v70 _dafny.Map
				_ = _2442_v70
				_2442_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2438_v65, _2440_v68)
				var _2443_v71 _dafny.Map
				_ = _2443_v71
				_2443_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2438_v65, func() _dafny.Map {
					var _coll97 = _dafny.NewMapBuilder()
					_ = _coll97
					for _iter104 := _dafny.Iterate((_2441_v69).Keys().Elements()); ; {
						_compr_97, _ok104 := _iter104()
						if !_ok104 {
							break
						}
						var _2444_v66 _dafny.Int
						_2444_v66 = interface{}(_compr_97).(_dafny.Int)
						if (_2441_v69).Contains(_2444_v66) {
							_coll97.Add(Companion_Default___.SafeModuloInt(_2444_v66, _this.F29), _dafny.IntOfUint32((_2438_v65).Cardinality()))
						}
					}
					return _coll97.ToMap()
				}())).Merge(_2442_v70), _2426_v54)
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_2432_v60), 0))
				_ = _index340
				var _rhs358 _dafny.Int = (_dafny.IntOfUint32((_2435_v63).Cardinality())).Times((func() _dafny.Int {
					if (_2437_v64).Contains(_2383_v19) {
						return (_2437_v64).Multiplicity(_2383_v19)
					}
					return _2423_v51
				})())
				_ = _rhs358
				var _rhs359 T3 = (func() T3 {
					if (_2443_v71).Contains(_2442_v70) {
						return (_2443_v71).Get(_2442_v70).(T3)
					}
					return _2426_v54
				})()
				_ = _rhs359
				var _lhs275 *GlobalState = globalState
				_ = _lhs275
				var _lhs276 _dafny.Array = _2432_v60
				_ = _lhs276
				var _lhs277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_2432_v60), 0))
				_ = _lhs277
				_lhs275.F7 = _rhs358
				(_lhs276).ArraySet1(_rhs359, (_lhs277).Int())
			}
			(_this).F29 = _this.F29
			var _2445_v72 *C1
			_ = _2445_v72
			var _nw419 *C1 = New_C1_()
			_ = _nw419
			_nw419.Ctor__((_this).F28(), (func() bool {
				if (_this).F24() {
					return (_this).F28()
				}
				return (_this).F28()
			})())
			_2445_v72 = _nw419
		}
		var _2446_v73 _dafny.Sequence
		_ = _2446_v73
		_2446_v73 = _dafny.SeqOf(_dafny.IntOfInt64(718), _this.F29)
		r0 = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_2446_v73).Cardinality()))).Cmp(_this.F29) < 0
		r1 = _this.F29
		var _2447_v74 _dafny.Set
		_ = _2447_v74
		_2447_v74 = _dafny.SetOf((_this).F24())
		var _2448_v75 _dafny.Sequence
		_ = _2448_v75
		_2448_v75 = _dafny.SeqOf((_this).F24())
		r2 = !((_this).F28()) || ((_2447_v74).IsDisjointFrom(_dafny.SetOf((_2448_v75).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2448_v75).Cardinality()))).Uint32()).(bool))))
		return r0, r1, r2
	}
}
func (_this *C10) M3(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _hi15 _dafny.Int = _this.F29
		_ = _hi15
		for _2449_i0 := _this.F29; _2449_i0.Cmp(_hi15) < 0; _2449_i0 = _2449_i0.Plus(_dafny.One) {
			(globalState).F13 = (_this).F28()
			var _2450_v0 _dafny.MultiSet
			_ = _2450_v0
			_2450_v0 = _dafny.MultiSetOf((_this).F24())
			var _2451_v1 _dafny.Sequence
			_ = _2451_v1
			_2451_v1 = _dafny.SeqOf(_2449_i0, _this.F29, _dafny.IntOfInt64(-260))
			var _2452_v2 _dafny.Int
			_ = _2452_v2
			var _2453_v3 bool
			_ = _2453_v3
			var _2454_v4 _dafny.Int
			_ = _2454_v4
			var _out72 _dafny.Int
			_ = _out72
			var _out73 bool
			_ = _out73
			var _out74 _dafny.Int
			_ = _out74
			_out72, _out73, _out74 = Companion_Default___.M0(_this.F29, (_2450_v0).Difference((_2450_v0).Update(p0, Companion_Default___.Abs(_2449_i0))), (_this).F28(), Companion_Default___.Fm0((_this).F27(), (_2451_v1).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2451_v1).Cardinality()))).Uint32()).(_dafny.Int), globalState), globalState)
			_2452_v2 = _out72
			_2453_v3 = _out73
			_2454_v4 = _out74
			var _2455_v5 *C6
			_ = _2455_v5
			var _nw420 *C6 = New_C6_()
			_ = _nw420
			_nw420.Ctor__((_this).Fm10((_this).F27(), _2452_v2, globalState), Companion_Default___.Fm2(globalState))
			_2455_v5 = _nw420
			var _2456_v6 D23
			_ = _2456_v6
			_2456_v6 = Companion_D23_.Create_DC54_(_2455_v5)
			_2455_v5 = (_2456_v6).Dtor_cf84()
			var _2457_v7 _dafny.Set
			_ = _2457_v7
			_2457_v7 = _dafny.SetOf(_2452_v2)
			var _2458_v8 _dafny.Map
			_ = _2458_v8
			_2458_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Set {
				if !(p0) {
					return _2457_v7
				}
				return _2457_v7
			})())
			var _2459_v10 _dafny.Map
			_ = _2459_v10
			_2459_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2449_i0, _2452_v2)
			var _2460_v11 *C4
			_ = _2460_v11
			var _nw421 *C4 = New_C4_()
			_ = _nw421
			_nw421.Ctor__((func() _dafny.Set {
				var _coll98 = _dafny.NewBuilder()
				_ = _coll98
				for _iter105 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(272), _dafny.IntOfInt64(866))); ; {
					_compr_98, _ok105 := _iter105()
					if !_ok105 {
						break
					}
					var _2461_v9 _dafny.Int
					_2461_v9 = interface{}(_compr_98).(_dafny.Int)
					if ((_dafny.IntOfInt64(272)).Cmp(_2461_v9) <= 0) && ((_2461_v9).Cmp(_dafny.IntOfInt64(866)) < 0) {
						_coll98.Add((_2461_v9).Minus(_2454_v4))
					}
				}
				return _coll98.ToSet()
			}()).Cardinality(), _2459_v10, (_this).F25(), _this.F23(), _2453_v3)
			_2460_v11 = _nw421
			var _2462_v12 _dafny.Sequence
			_ = _2462_v12
			_2462_v12 = _dafny.SeqOf(_2460_v11)
			var _2463_v13 _dafny.Sequence
			_ = _2463_v13
			_2463_v13 = _dafny.SeqOf(p0)
			var _rhs360 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23(), _dafny.SetOf(_dafny.IntOfInt64(687)))).Merge((_2458_v8).Merge(_2458_v8))
			_ = _rhs360
			var _rhs361 _dafny.Int = (_dafny.MultiSetFromSeq(_2451_v1)).Cardinality()
			_ = _rhs361
			var _rhs362 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2462_v12).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2463_v13, _2463_v13)).Cardinality()))
			_ = _rhs362
			var _rhs363 bool = _2453_v3
			_ = _rhs363
			var _lhs278 *C10 = _this
			_ = _lhs278
			var _lhs279 *GlobalState = globalState
			_ = _lhs279
			_2458_v8 = _rhs360
			_lhs278.F29 = _rhs361
			r1 = _rhs362
			_lhs279.F13 = _rhs363
		}
		var _hi16 _dafny.Int = _this.F29
		_ = _hi16
		for _2464_i1 := _dafny.IntOfInt64(186); _2464_i1.Cmp(_hi16) < 0; _2464_i1 = _2464_i1.Plus(_dafny.One) {
			var _2465_v14 _dafny.Set
			_ = _2465_v14
			_2465_v14 = _dafny.SetOf(!((_this).F27()), p0, p0)
			var _2466_v15 _dafny.Sequence
			_ = _2466_v15
			_2466_v15 = _dafny.UnicodeSeqOfUtf8Bytes("eeajmfk")
			var _2467_v16 D19
			_ = _2467_v16
			_2467_v16 = Companion_D19_.Create_DC46_((_2465_v14).Union(_2465_v14), _2466_v15)
			_2467_v16 = _2467_v16
			(_this).F29 = _this.F29
			(globalState).F7 = (_2464_i1).Plus(_this.F29)
			(_this).F23_set_(_this.F23())
		}
		var _2468_v17 _dafny.Array
		_ = _2468_v17
		var _nw422 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
		_ = _nw422
		_2468_v17 = _nw422
		var _2469_v18 _dafny.Array
		_ = _2469_v18
		var _len0_72 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_72
		var _nw423 _dafny.Array
		_ = _nw423
		if _len0_72.Cmp(_dafny.Zero) == 0 {
			_nw423 = _dafny.NewArray(_len0_72)
		} else {
			var _init72 func(_dafny.Int) bool = func(_2470_i2 _dafny.Int) bool {
				return (_this).F28()
			}
			_ = _init72
			var _element0_72 = _init72(_dafny.Zero)
			_ = _element0_72
			_nw423 = _dafny.NewArrayFromExample(_element0_72, nil, _len0_72)
			(_nw423).ArraySet1(_element0_72, 0)
			var _nativeLen0_72 = (_len0_72).Int()
			_ = _nativeLen0_72
			for _i0_72 := 1; _i0_72 < _nativeLen0_72; _i0_72++ {
				(_nw423).ArraySet1(_init72(_dafny.IntOf(_i0_72)), _i0_72)
			}
		}
		_2469_v18 = _nw423
		var _2471_v19 _dafny.Set
		_ = _2471_v19
		_2471_v19 = _dafny.SetOf(_2469_v18, _2469_v18, _2469_v18, _2469_v18, _2469_v18)
		var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2468_v17), 0))
		_ = _index341
		(_2468_v17).ArraySet1(_2471_v19, (_index341).Int())
		var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2468_v17), 0))
		_ = _index342
		(_2468_v17).ArraySet1(((func() _dafny.Set {
			if (_this).F27() {
				return _2471_v19
			}
			return _dafny.SetOf(_2469_v18, _2469_v18)
		})()).Intersection(_2471_v19), (_index342).Int())
		r2 = (_this).F24()
		var _2472_v20 _dafny.Array
		_ = _2472_v20
		var _nw424 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
		_ = _nw424
		_2472_v20 = _nw424
		var _2473_v21 _dafny.Array
		_ = _2473_v21
		var _nw425 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw425
		_2473_v21 = _nw425
		var _2474_v22 _dafny.Map
		_ = _2474_v22
		_2474_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2473_v21)
		var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_2472_v20), 0))
		_ = _index343
		(_2472_v20).ArraySet1((_2474_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2473_v21)), (_index343).Int())
		var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_2472_v20), 0))
		_ = _index344
		(_2472_v20).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _2473_v21)).Merge(_2474_v22)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2473_v21)), (_index344).Int())
		var _hi17 _dafny.Int = _this.F29
		_ = _hi17
		for _2475_i3 := _this.F29; _2475_i3.Cmp(_hi17) < 0; _2475_i3 = _2475_i3.Plus(_dafny.One) {
			var _2476_v23 _dafny.Set
			_ = _2476_v23
			_2476_v23 = _dafny.SetOf((_this).F24())
			var _2477_v24 *C5
			_ = _2477_v24
			var _nw426 *C5 = New_C5_()
			_ = _nw426
			_nw426.Ctor__(_2469_v18)
			_2477_v24 = _nw426
			var _2478_v25 _dafny.Set
			_ = _2478_v25
			_2478_v25 = _dafny.SetOf(_2477_v24, _2477_v24)
			var _rhs364 _dafny.Array = _2473_v21
			_ = _rhs364
			var _rhs365 bool = true
			_ = _rhs365
			var _rhs366 _dafny.Int = (_dafny.Zero).Minus((((func() _dafny.Set {
				if (_this).F24() {
					return _2476_v23
				}
				return _2476_v23
			})()).Union(_2476_v23)).Cardinality())
			_ = _rhs366
			var _rhs367 bool = (func() bool {
				if (_this).F27() {
					return (func() bool {
						if false {
							return Companion_Default___.Fm2(globalState)
						}
						return p0
					})()
				}
				return (_2478_v25).IsProperSubsetOf(_2478_v25)
			})()
			_ = _rhs367
			var _lhs280 *GlobalState = globalState
			_ = _lhs280
			var _lhs281 *C10 = _this
			_ = _lhs281
			_lhs280.F22 = _rhs364
			r0 = _rhs365
			_lhs281.F29 = _rhs366
			r2 = _rhs367
			var _2479_v26 _dafny.Sequence
			_ = _2479_v26
			_2479_v26 = _dafny.SeqOf(_this.F29)
			var _2480_v27 T1
			_ = _2480_v27
			var _nw427 *C6 = New_C6_()
			_ = _nw427
			_nw427.Ctor__(p1, p0)
			_2480_v27 = _nw427
			var _2481_v28 _dafny.Array
			_ = _2481_v28
			var _nwElement0_75 T1 = _2480_v27
			_ = _nwElement0_75
			var _nw428 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(17))
			_ = _nw428
			(_nw428).ArraySet1(_nwElement0_75, 0)
			(_nw428).ArraySet1(_2480_v27, 1)
			(_nw428).ArraySet1(_2480_v27, 2)
			(_nw428).ArraySet1(_2480_v27, 3)
			(_nw428).ArraySet1(_2480_v27, 4)
			(_nw428).ArraySet1(_2480_v27, 5)
			(_nw428).ArraySet1(_2480_v27, 6)
			(_nw428).ArraySet1(_2480_v27, 7)
			(_nw428).ArraySet1((func() T1 {
				if p0 {
					return _2480_v27
				}
				return _2480_v27
			})(), 8)
			(_nw428).ArraySet1(_2480_v27, 9)
			(_nw428).ArraySet1(_2480_v27, 10)
			(_nw428).ArraySet1(_2480_v27, 11)
			(_nw428).ArraySet1(_2480_v27, 12)
			(_nw428).ArraySet1(_2480_v27, 13)
			(_nw428).ArraySet1(_2480_v27, 14)
			(_nw428).ArraySet1(_2480_v27, 15)
			(_nw428).ArraySet1(_2480_v27, 16)
			_2481_v28 = _nw428
			var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_2481_v28), 0))
			_ = _index345
			(_2481_v28).ArraySet1(_2480_v27, (_index345).Int())
			var _2482_v29 _dafny.Sequence
			_ = _2482_v29
			_2482_v29 = _dafny.UnicodeSeqOfUtf8Bytes("vfbyqdoi")
			var _2483_v30 _dafny.Sequence
			_ = _2483_v30
			_2483_v30 = _dafny.SeqOf(p0)
			var _2484_v32 D5
			_ = _2484_v32
			_2484_v32 = Companion_D5_.Create_DC12_(func() _dafny.Set {
				var _coll99 = _dafny.NewBuilder()
				_ = _coll99
				for _iter106 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_2482_v29).Cardinality()), _dafny.IntOfUint32((_2483_v30).Cardinality()))).Elements()); ; {
					_compr_99, _ok106 := _iter106()
					if !_ok106 {
						break
					}
					var _2485_v31 _dafny.Int
					_2485_v31 = interface{}(_compr_99).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_2482_v29).Cardinality()), _dafny.IntOfUint32((_2483_v30).Cardinality())), _2485_v31) {
						_coll99.Add((_2485_v31).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ru")).Cardinality())))
					}
				}
				return _coll99.ToSet()
			}())
			var _2486_v33 _dafny.Map
			_ = _2486_v33
			_2486_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2482_v29)
			var _2487_v34 D2
			_ = _2487_v34
			_2487_v34 = Companion_D2_.Create_DC6_(_2486_v33)
			var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_2481_v28), 0))
			_ = _index346
			var _rhs368 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2475_i3), Companion_Default___.Fm45((_this).F28(), _2487_v34, _dafny.IntOfInt64(633), globalState))
			_ = _rhs368
			var _rhs369 _dafny.CodePoint = _2480_v27.F23()
			_ = _rhs369
			var _rhs370 T1 = _2480_v27
			_ = _rhs370
			var _rhs371 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2482_v29, _2482_v29)
			_ = _rhs371
			var _rhs372 D5 = (func() D5 {
				if Companion_Default___.Fm2(globalState) {
					return _2484_v32
				}
				return _2484_v32
			})()
			_ = _rhs372
			var _lhs282 *C10 = _this
			_ = _lhs282
			var _lhs283 _dafny.Array = _2481_v28
			_ = _lhs283
			var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_2481_v28), 0))
			_ = _lhs284
			_2479_v26 = _rhs368
			_lhs282.F23_set_(_rhs369)
			(_lhs283).ArraySet1(_rhs370, (_lhs284).Int())
			_2482_v29 = _rhs371
			_2484_v32 = _rhs372
			if !(false) || ((_this).F27()) {
				(_this).F23_set_(_2480_v27.F23())
				(_this).F29 = (_dafny.Zero).Minus(_2475_i3)
				var _2488_v35 D1
				_ = _2488_v35
				_2488_v35 = Companion_D1_.Create_DC4_(_2480_v27.F23())
				(_this).F23_set_((_2488_v35).Dtor_cf6())
				var _2489_v36 _dafny.Array
				_ = _2489_v36
				var _len0_73 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_73
				var _nw429 _dafny.Array
				_ = _nw429
				if _len0_73.Cmp(_dafny.Zero) == 0 {
					_nw429 = _dafny.NewArray(_len0_73)
				} else {
					var _init73 func(_dafny.Int) _dafny.Sequence = (func(_2490_v29 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2491_i4 _dafny.Int) _dafny.Sequence {
							return _2490_v29
						}
					})(_2482_v29)
					_ = _init73
					var _element0_73 = _init73(_dafny.Zero)
					_ = _element0_73
					_nw429 = _dafny.NewArrayFromExample(_element0_73, nil, _len0_73)
					(_nw429).ArraySet1(_element0_73, 0)
					var _nativeLen0_73 = (_len0_73).Int()
					_ = _nativeLen0_73
					for _i0_73 := 1; _i0_73 < _nativeLen0_73; _i0_73++ {
						(_nw429).ArraySet1(_init73(_dafny.IntOf(_i0_73)), _i0_73)
					}
				}
				_2489_v36 = _nw429
				var _2492_v37 _dafny.Map
				_ = _2492_v37
				_2492_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2482_v29, _2482_v29)
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_2489_v36), 0))
				_ = _index347
				(_2489_v36).ArraySet1((func() _dafny.Sequence {
					if (_2492_v37).Contains(_2482_v29) {
						return (_2492_v37).Get(_2482_v29).(_dafny.Sequence)
					}
					return _2482_v29
				})(), (_index347).Int())
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_2489_v36), 0))
				_ = _index348
				(_2489_v36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2482_v29, _dafny.Companion_Sequence_.Concatenate(_2482_v29, _dafny.UnicodeSeqOfUtf8Bytes("lydhu"))), (_index348).Int())
				(globalState).F7 = _this.F29
			} else {
				var _2493_v38 _dafny.MultiSet
				_ = _2493_v38
				_2493_v38 = _dafny.MultiSetOf((_this).F24(), (_2480_v27).F24(), true)
				var _2494_v39 _dafny.Map
				_ = _2494_v39
				_2494_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_2493_v38), _dafny.Companion_Sequence_.Concatenate(_2482_v29, _2482_v29))
				var _2495_v40 D0
				_ = _2495_v40
				_2495_v40 = Companion_D0_.Create_DC0_(_2493_v38)
				_2494_v39 = (_2494_v39).Update(_2495_v40, _dafny.Companion_Sequence_.Concatenate(_2482_v29, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg120 _dafny.Int) interface{} {
						return coer120(arg120)
					}
				}((func(_2496_v27 T1) func(_dafny.Int) _dafny.CodePoint {
					return func(_2497_i5 _dafny.Int) _dafny.CodePoint {
						return _2496_v27.F23()
					}
				})(_2480_v27)))))
				r1 = (_this.F29).Times(_dafny.IntOfInt64(968))
				var _2498_v41 _dafny.Array
				_ = _2498_v41
				var _nwElement0_76 _dafny.CodePoint = _2480_v27.F23()
				_ = _nwElement0_76
				var _nw430 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(7))
				_ = _nw430
				(_nw430).ArraySet1CodePoint(_nwElement0_76, 0)
				(_nw430).ArraySet1CodePoint(_this.F23(), 1)
				(_nw430).ArraySet1CodePoint(_2480_v27.F23(), 2)
				(_nw430).ArraySet1CodePoint(Companion_Default___.Fm29(_2475_i3, (_this).F28(), (_this).F27(), (_dafny.Zero).Minus(_2475_i3), globalState), 3)
				(_nw430).ArraySet1CodePoint(_dafny.CodePoint('w'), 4)
				(_nw430).ArraySet1CodePoint(_2480_v27.F23(), 5)
				(_nw430).ArraySet1CodePoint(_this.F23(), 6)
				_2498_v41 = _nw430
				_2498_v41 = _2498_v41
				(globalState).F13 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_2479_v26).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2479_v26).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfInt64(471))).Cmp(_dafny.IntOfInt64(552)) != 0
				var _2499_v42 _dafny.Map
				_ = _2499_v42
				_2499_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_i3, _2482_v29)
				_2499_v42 = (_2499_v42).Update(Companion_Default___.SafeModuloInt(_this.F29, (_2476_v23).Cardinality()), _2482_v29)
			}
			var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_2469_v18), 0))
			_ = _index349
			(_2469_v18).ArraySet1(((_this).F28()) && (p0), (_index349).Int())
			var _2500_v43 _dafny.Map
			_ = _2500_v43
			_2500_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F27())
			var _2501_v44 _dafny.Sequence
			_ = _2501_v44
			_2501_v44 = _dafny.SeqOf(_2500_v43, _2500_v43, _2500_v43, _2500_v43)
			var _2502_v45 _dafny.Map
			_ = _2502_v45
			_2502_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2473_v21, (_2501_v44).Select((Companion_Default___.SafeIndex((_2476_v23).Cardinality(), _dafny.IntOfUint32((_2501_v44).Cardinality()))).Uint32()).(_dafny.Map))
			var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen(((_2477_v24).F32()), 0))
			_ = _index350
			((_2477_v24).F32()).ArraySet1((_this).F28(), (_index350).Int())
			var _2503_v46 _dafny.Map
			_ = _2503_v46
			_2503_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2477_v24).F32(), p1)
			var _2504_v47 *C3
			_ = _2504_v47
			var _nw431 *C3 = New_C3_()
			_ = _nw431
			_nw431.Ctor__(_2503_v46, (_this).F25(), _this.F23(), (_this).F27())
			_2504_v47 = _nw431
			var _2505_v48 _dafny.Set
			_ = _2505_v48
			_2505_v48 = _dafny.SetOf(_2504_v47, _2504_v47, _2504_v47, _2504_v47, _2504_v47)
			var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_2469_v18), 0))
			_ = _index351
			var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen(((_2477_v24).F32()), 0))
			_ = _index352
			var _rhs373 bool = !((_dafny.IntOfInt64(916)).Cmp((_this.F29).Plus(_2475_i3)) != 0)
			_ = _rhs373
			var _rhs374 _dafny.Map = _2502_v45
			_ = _rhs374
			var _rhs375 bool = (_2480_v27).Fm4((_2505_v48).Cardinality(), Companion_Default___.Fm36(globalState), globalState)
			_ = _rhs375
			var _lhs285 _dafny.Array = _2469_v18
			_ = _lhs285
			var _lhs286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_2469_v18), 0))
			_ = _lhs286
			var _lhs287 _dafny.Array = (_2477_v24).F32()
			_ = _lhs287
			var _lhs288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen(((_2477_v24).F32()), 0))
			_ = _lhs288
			(_lhs285).ArraySet1(_rhs373, (_lhs286).Int())
			_2502_v45 = _rhs374
			(_lhs287).ArraySet1(_rhs375, (_lhs288).Int())
		}
		r0 = (_this).F28()
		r1 = _this.F29
		var _2506_v49 D6
		_ = _2506_v49
		_2506_v49 = Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0), (_this).F27(), p0)
		r2 = (_2506_v49).Dtor_cf26()
		return r0, r1, r2
	}
}
func (_this *C10) M4(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _2507_v0 _dafny.Sequence
		_ = _2507_v0
		_2507_v0 = _dafny.SeqOf(_this.F29, _this.F29)
		var _2508_v1 _dafny.Map
		_ = _2508_v1
		_2508_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2507_v0, _this.F23())
		var _2509_v2 _dafny.MultiSet
		_ = _2509_v2
		_2509_v2 = _dafny.MultiSetOf((func() _dafny.CodePoint {
			if (_2508_v1).Contains(_2507_v0) {
				return (_2508_v1).Get(_2507_v0).(_dafny.CodePoint)
			}
			return _this.F23()
		})(), _this.F23())
		_2509_v2 = _2509_v2
		var _2510_v3 _dafny.Map
		_ = _2510_v3
		_2510_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p0)
		var _2511_v4 _dafny.Map
		_ = _2511_v4
		_2511_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_2510_v3).Contains((_this).F24()) {
				return (_2510_v3).Get((_this).F24()).(bool)
			}
			return (_this).F28()
		})(), (_this.F29).Minus(_this.F29))
		var _2512_v5 _dafny.MultiSet
		_ = _2512_v5
		_2512_v5 = _dafny.MultiSetOf((_this).F28(), (_this).F28(), !((_this).F27()), (_this).F28())
		var _2513_v6 _dafny.Map
		_ = _2513_v6
		_2513_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2512_v5).Cardinality(), (_this).F24())
		_2511_v4 = (_2511_v4).Update(((_2513_v6).Cardinality()).Cmp((Companion_Default___.Fm3((_this).F24(), globalState)).Cardinality()) >= 0, _this.F29)
		(globalState).F7 = _this.F29
		(_this).F29 = (_this).Fm6(globalState)
		var _2514_v7 _dafny.Array
		_ = _2514_v7
		var _nw432 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
		_ = _nw432
		_2514_v7 = _nw432
		var _2515_v8 _dafny.Sequence
		_ = _2515_v8
		_2515_v8 = _dafny.SeqOf(_2514_v7)
		var _2516_v9 _dafny.Array
		_ = _2516_v9
		var _nwElement0_77 _dafny.Array = _2514_v7
		_ = _nwElement0_77
		var _nw433 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(13))
		_ = _nw433
		(_nw433).ArraySet1(_nwElement0_77, 0)
		(_nw433).ArraySet1((_2515_v8).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2515_v8).Cardinality()))).Uint32()).(_dafny.Array), 1)
		(_nw433).ArraySet1(_2514_v7, 2)
		(_nw433).ArraySet1(_2514_v7, 3)
		(_nw433).ArraySet1(_2514_v7, 4)
		(_nw433).ArraySet1(_2514_v7, 5)
		(_nw433).ArraySet1(_2514_v7, 6)
		(_nw433).ArraySet1(_2514_v7, 7)
		(_nw433).ArraySet1(_2514_v7, 8)
		(_nw433).ArraySet1(_2514_v7, 9)
		(_nw433).ArraySet1(_2514_v7, 10)
		(_nw433).ArraySet1(_2514_v7, 11)
		(_nw433).ArraySet1(_2514_v7, 12)
		_2516_v9 = _nw433
		var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_2516_v9), 0))
		_ = _index353
		(_2516_v9).ArraySet1(_2514_v7, (_index353).Int())
		var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_2516_v9), 0))
		_ = _index354
		var _rhs376 bool = (_this).F24()
		_ = _rhs376
		var _rhs377 _dafny.Array = _2514_v7
		_ = _rhs377
		var _lhs289 _dafny.Array = _2516_v9
		_ = _lhs289
		var _lhs290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_2516_v9), 0))
		_ = _lhs290
		r0 = _rhs376
		(_lhs289).ArraySet1(_rhs377, (_lhs290).Int())
		var _2517_v10 _dafny.MultiSet
		_ = _2517_v10
		_2517_v10 = _dafny.MultiSetOf(_this.F29, (_2512_v5).Cardinality())
		var _2518_v11 _dafny.Map
		_ = _2518_v11
		_2518_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, ((_2517_v10).Difference(_2517_v10)).Cardinality())
		var _2519_v12 D19
		_ = _2519_v12
		_2519_v12 = Companion_D19_.Create_DC47_((_this).F24())
		var _2520_v13 _dafny.Array
		_ = _2520_v13
		var _nw434 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw434
		_2520_v13 = _nw434
		var _2521_v14 _dafny.Sequence
		_ = _2521_v14
		_2521_v14 = _dafny.UnicodeSeqOfUtf8Bytes("myrcehyl")
		var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_2520_v13), 0))
		_ = _index355
		(_2520_v13).ArraySet1(_2521_v14, (_index355).Int())
		var _2522_v15 _dafny.Sequence
		_ = _2522_v15
		_2522_v15 = _dafny.SeqOf((_this).F27(), (_this).F28())
		var _2523_v16 _dafny.Set
		_ = _2523_v16
		_2523_v16 = _dafny.SetOf((_2522_v15).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_2522_v15).Cardinality()))).Uint32()).(bool))
		var _2524_v17 D0
		_ = _2524_v17
		_2524_v17 = Companion_D0_.Create_DC2_((_this).F28(), (_2510_v3).Cardinality(), !(Companion_Default___.Fm2(globalState)))
		var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_2520_v13), 0))
		_ = _index356
		var _rhs378 _dafny.Map = ((_2518_v11).Update((_dafny.Zero).Minus(_this.F29), _this.F29)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _this.F29))
		_ = _rhs378
		var _rhs379 D19 = Companion_Default___.Fm57(globalState)
		_ = _rhs379
		var _rhs380 _dafny.Int = (_2524_v17).Dtor_cf3()
		_ = _rhs380
		var _rhs381 _dafny.Sequence = _2521_v14
		_ = _rhs381
		var _rhs382 _dafny.Set = (_2523_v16).Intersection(_2523_v16)
		_ = _rhs382
		var _lhs291 *GlobalState = globalState
		_ = _lhs291
		var _lhs292 _dafny.Array = _2520_v13
		_ = _lhs292
		var _lhs293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_2520_v13), 0))
		_ = _lhs293
		_2518_v11 = _rhs378
		_2519_v12 = _rhs379
		_lhs291.F7 = _rhs380
		(_lhs292).ArraySet1(_rhs381, (_lhs293).Int())
		_2523_v16 = _rhs382
		r0 = (_this).F24()
		return r0
	}
}
func (_this *C10) F28() bool {
	{
		return _this._f28
	}
}

// End of class C10
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
