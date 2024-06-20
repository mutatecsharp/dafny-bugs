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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(616), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-434))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(26))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	})))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(582))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(356)
	}))).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(343)), _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-843))), _dafny.MultiSetOf((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('d'))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _3_v0 _dafny.CodePoint
				_3_v0 = interface{}(_compr_1).(_dafny.CodePoint)
				if (_dafny.SetOf(_dafny.CodePoint('d'))).Contains(_3_v0) {
					_coll1.Add(_3_v0, false)
				}
			}
			return _coll1.ToMap()
		}()).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v1 _dafny.CodePoint
			_4_v1 = interface{}(_compr_0).(_dafny.CodePoint)
			if (func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('d'))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _3_v0 _dafny.CodePoint
					_3_v0 = interface{}(_compr_2).(_dafny.CodePoint)
					if (_dafny.SetOf(_dafny.CodePoint('d'))).Contains(_3_v0) {
						_coll2.Add(_3_v0, false)
					}
				}
				return _coll2.ToMap()
			}()).Contains(_4_v1) {
				_coll0.Add(_4_v1)
			}
		}
		return _coll0.ToSet()
	}()).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer3 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_5_i1 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
	}))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) bool {
	return !(!((!(false)) || (!(false)))) || ((_dafny.SetOf(_dafny.CodePoint('t'))).IsProperSubsetOf(_dafny.SetOf(_dafny.CodePoint('k'))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Set, p1 _dafny.CodePoint, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(884)
		}
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xjmw")).Cardinality())
	})()).Times((func() _dafny.Int {
		if true {
			return _dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality())
		}
		return _dafny.IntOfInt64(-443)
	})())
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(!(true)), false)
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Map, p3 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC5_()
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("bx"), true))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Difference((_dafny.SetOf(false)).Union(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(465), _dafny.IntOfInt64(64))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-896), _dafny.IntOfInt64(863)))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true, false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) D2 {
	if !(false) {
		return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC5_())
	} else {
		return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC3_(_dafny.SeqOf(Companion_D0_.Create_DC1_(true), Companion_D0_.Create_DC1_(false), Companion_D0_.Create_DC1_(false))))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('s')
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vn"), _dafny.UnicodeSeqOfUtf8Bytes("pyh")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.UnicodeSeqOfUtf8Bytes("nt")))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("iql")
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.Map, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(false))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-443), _dafny.IntOfInt64(-121))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _7_v0 _dafny.Int
			_7_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-443)).Cmp(_7_v0) <= 0) && ((_7_v0).Cmp(_dafny.IntOfInt64(-121)) < 0) {
				_coll3.Add((_7_v0).Minus((_dafny.MultiSetOf((_dafny.SetOf(false)).Cardinality())).Cardinality()))
			}
		}
		return _coll3.ToSet()
	}(), _dafny.SetOf(_dafny.IntOfInt64(917), _dafny.IntOfInt64(-461)), _dafny.SetOf((_dafny.MultiSetOf(false, !(!(true)))).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(-841), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), _dafny.IntOfInt64(36)), _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))))).Cardinality())) > 0, _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.CodePoint('j')))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.MultiSetOf(_dafny.IntOfInt64(-222), _dafny.IntOfInt64(310))).Cardinality())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true, false, false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(false)))), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sp")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(303)))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(965)))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(610)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC13_((func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(392)
		}
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ntyxfwen")).Cardinality())
	})(), _dafny.SetOf(false), _dafny.CodePoint('g'), _dafny.IntOfInt64(-831), false)
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Set, globalState *GlobalState) _dafny.MultiSet {
	if false {
		return (_dafny.MultiSetOf(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC3_(_dafny.SeqOf(Companion_D0_.Create_DC1_(false), Companion_D0_.Create_DC1_(true)))))).Intersection(_dafny.MultiSetOf(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC5_())))
	} else {
		return _dafny.MultiSetOf(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC3_(_dafny.SeqOf(Companion_D0_.Create_DC1_(true), Companion_D0_.Create_DC1_(!(false)), Companion_D0_.Create_DC1_(true), Companion_D0_.Create_DC1_(true)))))), Companion_D2_.Create_DC6_(Companion_D2_.Create_DC3_(_dafny.SeqOf(Companion_D0_.Create_DC1_(true), Companion_D0_.Create_DC1_(true)))))
	}
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC18_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(698), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("todmnh")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 _dafny.MultiSet, p2 _dafny.Int, p3 _dafny.Set, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, !(!(true)))).Difference(_dafny.SetOf(false, !(false)))).Union(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(false))).Union(_dafny.MultiSetOf(false, !(true), false))).Intersection((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(!(!(true)))
		}
		return _dafny.MultiSetOf(false)
	})())
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(121), _dafny.IntOfInt64(39))).Intersection(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(953), false)).Cardinality())), (_dafny.IntOfInt64(532)).Cmp((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))).Cardinality())).Cardinality()) == 0)
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.SeqOf(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, false, true, true, false))))
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.CodePoint('l'))).Union((func() _dafny.Set {
		if true {
			return _dafny.SetOf(_dafny.CodePoint('i'), _dafny.CodePoint('x'))
		}
		return func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(492))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _10_v0 _dafny.CodePoint
				_10_v0 = interface{}(_compr_4).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(492))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), _10_v0) {
					_coll4.Add(_10_v0)
				}
			}
			return _coll4.ToSet()
		}()
	})())
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Map, p1 _dafny.Int, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-625))), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(454))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_11_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(580)
			}))))
		}
		return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(116), _dafny.IntOfInt64(990)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sng")).Cardinality())), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iajxrdi")).Cardinality()))), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pocfxd")).Cardinality()), false)).Cardinality())))
	})())).Difference((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(973), (func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf(false))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _12_v0 _dafny.Set
			_12_v0 = interface{}(_compr_5).(_dafny.Set)
			if (_dafny.MultiSetOf(_dafny.SetOf(false))).Contains(_12_v0) {
				_coll5.Add(_12_v0, false)
			}
		}
		return _coll5.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(59), _dafny.IntOfInt64(168)))).Intersection(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-520), _dafny.IntOfInt64(279)), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(150), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sq")).Cardinality()))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(789), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(-224), (_dafny.MultiSetOf(true)).Cardinality(), (_dafny.SetOf(false, true)).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-871), Companion_D18_.Create_DC39_(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sbpaa")).Cardinality()), Companion_D18_.Create_DC39_(!(true))))
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(573), Companion_D13_.Create_DC28_(true, _dafny.CodePoint('j')))).Merge(func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_13_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))).Cardinality())
			}))).Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _15_v0 _dafny.Int
				_15_v0 = interface{}(_compr_7).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_13_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					}))).Cardinality())
				})), _15_v0) {
					_coll7.Add((_15_v0).Minus(_dafny.IntOfInt64(91)), Companion_D13_.Create_DC28_(false, _dafny.CodePoint('o')))
				}
			}
			return _coll7.ToMap()
		}())).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _16_v1 _dafny.Int
			_16_v1 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(573), Companion_D13_.Create_DC28_(true, _dafny.CodePoint('j')))).Merge(func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}(func(_13_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg14 _dafny.Int) interface{} {
							return coer14(arg14)
						}
					}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					}))).Cardinality())
				}))).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _15_v0 _dafny.Int
					_15_v0 = interface{}(_compr_8).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg15 _dafny.Int) interface{} {
							return coer15(arg15)
						}
					}(func(_13_i0 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg16 _dafny.Int) interface{} {
								return coer16(arg16)
							}
						}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('w')
						}))).Cardinality())
					})), _15_v0) {
						_coll8.Add((_15_v0).Minus(_dafny.IntOfInt64(91)), Companion_D13_.Create_DC28_(false, _dafny.CodePoint('o')))
					}
				}
				return _coll8.ToMap()
			}())).Contains(_16_v1) {
				_coll6.Add(Companion_Default___.SafeDivisionInt(_16_v1, _dafny.IntOfInt64(316)))
			}
		}
		return _coll6.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(31), _dafny.IntOfInt64(724))).Cardinality()), func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(511), _dafny.IntOfInt64(620))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _17_v0 _dafny.Int
			_17_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(511)).Cmp(_17_v0) <= 0) && ((_17_v0).Cmp(_dafny.IntOfInt64(620)) < 0) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_17_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oyikk")).Cardinality())))
			}
		}
		return _coll9.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	var r0 bool = false
	_ = r0
	var _18_v0 _dafny.CodePoint
	_ = _18_v0
	_18_v0 = _dafny.CodePoint('q')
	var _19_v1 D13
	_ = _19_v1
	_19_v1 = Companion_D13_.Create_DC28_(p0, _18_v0)
	var _source0 D13 = _19_v1
	_ = _source0
	if _source0.Is_DC27() {
		var _20___mcc_h0 _dafny.CodePoint = _source0.Get_().(D13_DC27).Cf43
		_ = _20___mcc_h0
		var _21___mcc_h1 T0 = _source0.Get_().(D13_DC27).Cf44
		_ = _21___mcc_h1
		var _22___mcc_h2 _dafny.CodePoint = _source0.Get_().(D13_DC27).Cf45
		_ = _22___mcc_h2
		var _23___mcc_h3 _dafny.Int = _source0.Get_().(D13_DC27).Cf46
		_ = _23___mcc_h3
		var _24_cf46 _dafny.Int = _23___mcc_h3
		_ = _24_cf46
		var _25_cf45 _dafny.CodePoint = _22___mcc_h2
		_ = _25_cf45
		var _26_cf44 T0 = _21___mcc_h1
		_ = _26_cf44
		var _27_cf43 _dafny.CodePoint = _20___mcc_h0
		_ = _27_cf43
		var _28_v2 _dafny.Array
		_ = _28_v2
		var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw0
		_28_v2 = _nw0
		var _29_v3 D11
		_ = _29_v3
		_29_v3 = Companion_D11_.Create_DC24_(_28_v2, p0, p0)
		var _30_v4 _dafny.Set
		_ = _30_v4
		_30_v4 = _dafny.SetOf(_29_v3)
		var _31_v5 _dafny.Array
		_ = _31_v5
		_31_v5 = _28_v2
		var _32_v6 _dafny.Map
		_ = _32_v6
		_32_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _33_v7 _dafny.Sequence
		_ = _33_v7
		_33_v7 = _dafny.UnicodeSeqOfUtf8Bytes("gvt")
		var _34_v8 D18
		_ = _34_v8
		_34_v8 = Companion_D18_.Create_DC38_(_30_v4)
		var _35_v9 _dafny.Array
		_ = _35_v9
		var _nwElement0_0 _dafny.Set = (_30_v4).Difference(_30_v4)
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(20))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1((func() _dafny.Set {
			if p0 {
				return _30_v4
			}
			return _30_v4
		})(), 1)
		(_nw1).ArraySet1((_dafny.SetOf(_29_v3)).Intersection(_30_v4), 2)
		(_nw1).ArraySet1((_dafny.SetOf(Companion_D11_.Create_DC24_((_31_v5), p0, p0))).Difference(_30_v4), 3)
		(_nw1).ArraySet1(_30_v4, 4)
		(_nw1).ArraySet1(_30_v4, 5)
		(_nw1).ArraySet1(_30_v4, 6)
		(_nw1).ArraySet1(_30_v4, 7)
		(_nw1).ArraySet1(_dafny.SetOf(Companion_D11_.Create_DC24_(_28_v2, p0, p0)), 8)
		(_nw1).ArraySet1(_30_v4, 9)
		(_nw1).ArraySet1(_30_v4, 10)
		(_nw1).ArraySet1(_30_v4, 11)
		(_nw1).ArraySet1(_30_v4, 12)
		(_nw1).ArraySet1(_30_v4, 13)
		(_nw1).ArraySet1(_30_v4, 14)
		(_nw1).ArraySet1(_30_v4, 15)
		(_nw1).ArraySet1(_30_v4, 16)
		(_nw1).ArraySet1(_dafny.SetOf(Companion_D11_.Create_DC24_(_28_v2, (func() bool {
			if (_32_v6).Contains((_26_cf44).F0()) {
				return (_32_v6).Get((_26_cf44).F0()).(bool)
			}
			return true
		})(), p0), _29_v3, Companion_D11_.Create_DC24_(_28_v2, p0, (_26_cf44).Fm6(p0, _33_v7, _33_v7, _24_cf46, globalState))), 17)
		(_nw1).ArraySet1(_dafny.SetOf(_29_v3), 18)
		(_nw1).ArraySet1(((_34_v8).Dtor_cf58()).Union(_30_v4), 19)
		_35_v9 = _nw1
		_35_v9 = _35_v9
		var _36_v10 _dafny.Map
		_ = _36_v10
		_36_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0), p1)
		var _37_v11 _dafny.Map
		_ = _37_v11
		_37_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_26_cf44).F0(), (_26_cf44).F0())
		var _38_v12 *C5
		_ = _38_v12
		var _nw2 *C5 = New_C5_()
		_ = _nw2
		_nw2.Ctor__(_25_cf45, (_dafny.Zero).Minus((Companion_D5_.Create_DC12_(p1, (func() _dafny.Int {
			if (_36_v10).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)) {
				return (_36_v10).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).(_dafny.Int)
			}
			return p1
		})())).Dtor_cf15()), (_37_v11).Cardinality())
		_38_v12 = _nw2
		var _39_v13 *C2
		_ = _39_v13
		var _nw3 *C2 = New_C2_()
		_ = _nw3
		_nw3.Ctor__(_dafny.IntOfUint32((_33_v7).Cardinality()))
		_39_v13 = _nw3
		var _rhs0 *C5 = _38_v12
		_ = _rhs0
		var _rhs1 *C2 = _39_v13
		_ = _rhs1
		var _rhs2 _dafny.Int = (_38_v12).F2()
		_ = _rhs2
		var _rhs3 bool = p0
		_ = _rhs3
		var _rhs4 bool = ((_38_v12).F2()).Cmp(Companion_Default___.SafeModuloInt(p1, p1)) == 0
		_ = _rhs4
		_38_v12 = _rhs0
		_39_v13 = _rhs1
		_24_cf46 = _rhs2
		r0 = _rhs3
		r0 = _rhs4
		var _40_v14 _dafny.Set
		_ = _40_v14
		_40_v14 = _dafny.SetOf((_38_v12).F2())
		var _41_v15 D3
		_ = _41_v15
		_41_v15 = Companion_D3_.Create_DC8_(p0, (_38_v12).F2())
		var _42_v16 D7
		_ = _42_v16
		_42_v16 = Companion_D7_.Create_DC17_(p0, p0, p0)
		var _43_v17 _dafny.Set
		_ = _43_v17
		_43_v17 = _dafny.SetOf(_42_v16)
		var _rhs5 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_38_v12).F2()), (_40_v14).Cardinality())
		_ = _rhs5
		var _rhs6 bool = !((_38_v12).Fm6(p0, _33_v7, (Companion_D2_.Create_DC4_(_33_v7, p0)).Dtor_cf4(), (_41_v15).Dtor_cf9(), globalState))
		_ = _rhs6
		var _rhs7 _dafny.Int = (_38_v12).F2()
		_ = _rhs7
		var _rhs8 bool = !((_dafny.MultiSetOf(_43_v17, _43_v17)).IsProperSubsetOf(_dafny.MultiSetOf(_43_v17)))
		_ = _rhs8
		_24_cf46 = _rhs5
		r0 = _rhs6
		_24_cf46 = _rhs7
		r0 = _rhs8
		var _44_v18 *C4
		_ = _44_v18
		var _nw4 *C4 = New_C4_()
		_ = _nw4
		_nw4.Ctor__(Companion_Default___.Fm3(_dafny.SetOf(p0), (_38_v12).F1(), _dafny.IntOfInt64(-553), _dafny.IntOfInt64(258), globalState))
		_44_v18 = _nw4
	} else if _source0.Is_DC28() {
		var _45___mcc_h4 bool = _source0.Get_().(D13_DC28).Cf47
		_ = _45___mcc_h4
		var _46___mcc_h5 _dafny.CodePoint = _source0.Get_().(D13_DC28).Cf48
		_ = _46___mcc_h5
		var _47_cf48 _dafny.CodePoint = _46___mcc_h5
		_ = _47_cf48
		var _48_cf47 bool = _45___mcc_h4
		_ = _48_cf47
		if (func() bool {
			if true {
				return _48_cf47
			}
			return _48_cf47
		})() {
			var _49_v19 T1
			_ = _49_v19
			var _nw5 *C5 = New_C5_()
			_ = _nw5
			_nw5.Ctor__(_47_cf48, _dafny.IntOfInt64(515), _dafny.IntOfInt64(237))
			_49_v19 = _nw5
			var _50_v20 _dafny.Map
			_ = _50_v20
			_50_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _49_v19)
			var _51_v21 _dafny.Map
			_ = _51_v21
			_51_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, (_50_v20).Merge(_50_v20))
			_51_v21 = (_51_v21).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_cf48, _50_v20))
			var _52_v22 _dafny.Int
			_ = _52_v22
			_52_v22 = _dafny.IntOfInt64(118)
			_52_v22 = p1
			var _53_v23 _dafny.MultiSet
			_ = _53_v23
			_53_v23 = _dafny.MultiSetOf(_52_v22, (_49_v19).F0(), _52_v22, (_49_v19).F0())
			var _54_v24 _dafny.MultiSet
			_ = _54_v24
			_54_v24 = _dafny.MultiSetOf((_53_v23).Difference((_53_v23).Update(p1, Companion_Default___.Abs(p1))), _53_v23, (_53_v23).Intersection(_53_v23))
			var _55_v25 _dafny.Map
			_ = _55_v25
			_55_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(p1))
			var _56_v26 _dafny.Map
			_ = _56_v26
			_56_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v19, p1)
			var _57_v27 _dafny.Sequence
			_ = _57_v27
			_57_v27 = _dafny.SeqOf(p1)
			var _58_v28 _dafny.Map
			_ = _58_v28
			_58_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_57_v27).Cardinality()), _dafny.SeqOf(_18_v0, _47_cf48))
			_54_v24 = Companion_Default___.Fm32((_55_v25).Merge(_55_v25), (_56_v26).Cardinality(), false, _58_v28, globalState)
			var _59_v29 _dafny.Sequence
			_ = _59_v29
			_59_v29 = _dafny.UnicodeSeqOfUtf8Bytes("wkht")
			var _60_v30 D2
			_ = _60_v30
			_60_v30 = Companion_D2_.Create_DC4_(_59_v29, _48_cf47)
			var _61_v31 _dafny.Map
			_ = _61_v31
			_61_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_60_v30, p1)
			var _62_v32 D18
			_ = _62_v32
			_62_v32 = Companion_D18_.Create_DC39_(p0)
			var _pat_let_tv0 = _48_cf47
			_ = _pat_let_tv0
			var _63_v33 _dafny.Map
			_ = _63_v33
			_63_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qhqkkmfj")).Cardinality()), func(_pat_let0_0 D18) D18 {
				return func(_64_dt__update__tmp_h0 D18) D18 {
					return func(_pat_let1_0 bool) D18 {
						return func(_65_dt__update_hcf59_h0 bool) D18 {
							return Companion_D18_.Create_DC39_(_65_dt__update_hcf59_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_62_v32))
			var _66_v35 _dafny.Sequence
			_ = _66_v35
			_66_v35 = _dafny.SeqOf(Companion_Default___.Fm33((_49_v19).F0(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC4_(_59_v29, true), _52_v22), globalState), Companion_Default___.Fm33(_52_v22, _61_v31, globalState), (_63_v33).Update((_dafny.Zero).Minus((_49_v19).F0()), _62_v32), Companion_Default___.Fm33((_49_v19).F0(), _61_v31, globalState), func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(648), _dafny.IntOfInt64(747))); ; {
					_compr_10, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _67_v34 _dafny.Int
					_67_v34 = interface{}(_compr_10).(_dafny.Int)
					if ((_dafny.IntOfInt64(648)).Cmp(_67_v34) <= 0) && ((_67_v34).Cmp(_dafny.IntOfInt64(747)) < 0) {
						_coll10.Add((_67_v34).Times((_dafny.MultiSetOf(_59_v29)).Cardinality()), _62_v32)
					}
				}
				return _coll10.ToMap()
			}())
			_52_v22 = _dafny.IntOfUint32((_66_v35).Cardinality())
			var _68_v36 _dafny.Array
			_ = _68_v36
			var _nw6 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
			_ = _nw6
			_68_v36 = _nw6
			var _rhs9 bool = p0
			_ = _rhs9
			var _rhs10 _dafny.Array = _68_v36
			_ = _rhs10
			_48_cf47 = _rhs9
			_68_v36 = _rhs10
		} else {
			var _69_v37 _dafny.Int
			_ = _69_v37
			_69_v37 = _dafny.IntOfInt64(223)
			_69_v37 = _69_v37
			var _70_v38 _dafny.Map
			_ = _70_v38
			_70_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_cf47, Companion_Default___.Fm2(p0, p1, p0, _48_cf47, globalState))
			var _71_v39 _dafny.Map
			_ = _71_v39
			_71_v39 = _70_v38
			var _72_v40 _dafny.Map
			_ = _72_v40
			_72_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v39, _48_cf47)
			_72_v40 = (_72_v40).Update(_71_v39, p0)
			var _73_v41 _dafny.Array
			_ = _73_v41
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw7
			_73_v41 = _nw7
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_73_v41), 0))
			_ = _index0
			(_73_v41).ArraySet1(_69_v37, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_73_v41), 0))
			_ = _index1
			(_73_v41).ArraySet1((Companion_Default___.Fm22(globalState)).Cardinality(), (_index1).Int())
			_48_cf47 = p0
			var _74_v42 _dafny.Sequence
			_ = _74_v42
			_74_v42 = _dafny.SeqOf((_73_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_73_v41), 0))).Int()).(_dafny.Int))
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_73_v41), 0))
			_ = _index2
			(_73_v41).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_74_v42, _74_v42)).Cardinality()), Companion_Default___.SafeModuloInt(_69_v37, p1)), (_index2).Int())
		}
		var _75_v43 _dafny.Int
		_ = _75_v43
		_75_v43 = _dafny.IntOfInt64(-188)
		var _76_v44 _dafny.Set
		_ = _76_v44
		_76_v44 = _dafny.SetOf(!(_48_cf47))
		var _77_v45 _dafny.Array
		_ = _77_v45
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_0
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_78_cf47 bool) func(_dafny.Int) bool {
				return func(_79_i0 _dafny.Int) bool {
					return _78_cf47
				}
			})(_48_cf47)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw8 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw8).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw8).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_77_v45 = _nw8
		var _80_v46 D11
		_ = _80_v46
		_80_v46 = Companion_D11_.Create_DC24_(_77_v45, _48_cf47, p0)
		var _81_v47 _dafny.Set
		_ = _81_v47
		_81_v47 = _dafny.SetOf(_80_v46, func(_pat_let2_0 D11) D11 {
			return func(_82_dt__update__tmp_h1 D11) D11 {
				return func(_pat_let3_0 bool) D11 {
					return func(_83_dt__update_hcf39_h0 bool) D11 {
						return Companion_D11_.Create_DC24_((_82_dt__update__tmp_h1).Dtor_cf38(), _83_dt__update_hcf39_h0, (_82_dt__update__tmp_h1).Dtor_cf40())
					}(_pat_let3_0)
				}(false)
			}(_pat_let2_0)
		}(_80_v46))
		var _84_v48 D18
		_ = _84_v48
		_84_v48 = Companion_D18_.Create_DC38_(_81_v47)
		var _85_v49 _dafny.MultiSet
		_ = _85_v49
		_85_v49 = _dafny.MultiSetOf(Companion_D18_.Create_DC38_(_dafny.SetOf(_80_v46)), _84_v48)
		var _86_v50 D19
		_ = _86_v50
		_86_v50 = Companion_D19_.Create_DC41_(_85_v49)
		_75_v43 = Companion_Default___.Fm3(_76_v44, _47_cf48, (func() _dafny.Int {
			if _48_cf47 {
				return ((_86_v50).Dtor_cf62()).Cardinality()
			}
			return _75_v43
		})(), _75_v43, globalState)
		if !(!((_75_v43).Cmp(_dafny.IntOfInt64(182)) == 0)) {
			_75_v43 = (_dafny.IntOfInt64(641)).Minus(p1)
			_77_v45 = _77_v45
			var _87_v51 _dafny.Sequence
			_ = _87_v51
			_87_v51 = _dafny.SeqOf(p0)
			_48_cf47 = Companion_Default___.Fm2((_76_v44).IsDisjointFrom(_76_v44), Companion_Default___.SafeDivisionInt(_75_v43, (_dafny.Zero).Minus(_dafny.IntOfUint32((_87_v51).Cardinality()))), p0, true, globalState)
			var _pat_let_tv1 = p0
			_ = _pat_let_tv1
			_19_v1 = (func() D13 {
				if _48_cf47 {
					return func(_pat_let4_0 D13) D13 {
						return func(_88_dt__update__tmp_h2 D13) D13 {
							return func(_pat_let5_0 bool) D13 {
								return func(_89_dt__update_hcf47_h0 bool) D13 {
									return Companion_D13_.Create_DC28_(_89_dt__update_hcf47_h0, (_88_dt__update__tmp_h2).Dtor_cf48())
								}(_pat_let5_0)
							}(_pat_let_tv1)
						}(_pat_let4_0)
					}(_19_v1)
				}
				return _19_v1
			})()
			var _90_v52 _dafny.Set
			_ = _90_v52
			_90_v52 = _dafny.SetOf(p1)
			var _91_v53 *C1
			_ = _91_v53
			var _nw9 *C1 = New_C1_()
			_ = _nw9
			_nw9.Ctor__(p1)
			_91_v53 = _nw9
			var _92_v54 _dafny.MultiSet
			_ = _92_v54
			_92_v54 = _dafny.MultiSetOf(_91_v53)
			var _93_v55 _dafny.Sequence
			_ = _93_v55
			_93_v55 = _dafny.SeqOf(_90_v52, _dafny.SetOf(Companion_Default___.Fm3(_76_v44, _dafny.CodePoint('d'), _dafny.IntOfInt64(872), (_92_v54).Cardinality(), globalState)))
			var _94_v56 _dafny.Sequence
			_ = _94_v56
			_94_v56 = _dafny.SeqOf((_93_v55).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_93_v55).Cardinality()))).Uint32()).(_dafny.Set))
			_90_v52 = ((_93_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.IntOfUint32((_93_v55).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_90_v52)
		} else {
			var _95_v57 *C1
			_ = _95_v57
			var _nw10 *C1 = New_C1_()
			_ = _nw10
			_nw10.Ctor__((p1).Times(p1))
			_95_v57 = _nw10
			var _96_v58 _dafny.Sequence
			_ = _96_v58
			_96_v58 = _dafny.UnicodeSeqOfUtf8Bytes("pmdqlx")
			var _97_v59 D2
			_ = _97_v59
			_97_v59 = Companion_D2_.Create_DC4_(_96_v58, p0)
			var _98_v60 _dafny.Map
			_ = _98_v60
			_98_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_97_v59).Dtor_cf4(), p0)
			var _99_v61 _dafny.Map
			_ = _99_v61
			_99_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_cf47, (_95_v57).Fm15(_98_v60, _48_cf47, _48_cf47, globalState))
			var _rhs11 _dafny.Array = _77_v45
			_ = _rhs11
			var _rhs12 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_96_v58, _96_v58)
			_ = _rhs12
			var _rhs13 *C1 = _95_v57
			_ = _rhs13
			var _rhs14 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return ((func() _dafny.Int {
						if (_99_v61).Contains(p0) {
							return (_99_v61).Get(p0).(_dafny.Int)
						}
						return _75_v43
					})()).Times(_75_v43)
				}
				return _dafny.IntOfInt64(-427)
			})()
			_ = _rhs14
			_77_v45 = _rhs11
			r0 = _rhs12
			_95_v57 = _rhs13
			_75_v43 = _rhs14
			var _100_v62 T1
			_ = _100_v62
			var _nw11 *C5 = New_C5_()
			_ = _nw11
			_nw11.Ctor__((_96_v58).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_96_v58).Cardinality()))).Uint32()).(_dafny.CodePoint), p1, Companion_Default___.SafeModuloInt(_75_v43, _75_v43))
			_100_v62 = _nw11
			_100_v62 = _100_v62
			var _101_v63 *C1
			_ = _101_v63
			var _nw12 *C1 = New_C1_()
			_ = _nw12
			_nw12.Ctor__((_100_v62).F0())
			_101_v63 = _nw12
			var _102_v64 _dafny.Set
			_ = _102_v64
			_102_v64 = _dafny.SetOf((_dafny.Zero).Minus(_75_v43), p1, p1, _dafny.IntOfInt64(308))
			var _103_v65 _dafny.MultiSet
			_ = _103_v65
			_103_v65 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_99_v61).Contains(_48_cf47) {
					return (_99_v61).Get(_48_cf47).(_dafny.Int)
				}
				return (_100_v62).F0()
			})())
			var _104_v66 _dafny.Array
			_ = _104_v66
			var _nwElement0_1 _dafny.Int = p1
			_ = _nwElement0_1
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(15))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_1, 0)
			(_nw13).ArraySet1(_75_v43, 1)
			(_nw13).ArraySet1((_102_v64).Cardinality(), 2)
			(_nw13).ArraySet1(_dafny.IntOfInt64(132), 3)
			(_nw13).ArraySet1(((_100_v62).F0()).Plus((_100_v62).F0()), 4)
			(_nw13).ArraySet1(p1, 5)
			(_nw13).ArraySet1(_dafny.IntOfInt64(854), 6)
			(_nw13).ArraySet1(_dafny.IntOfInt64(842), 7)
			(_nw13).ArraySet1(p1, 8)
			(_nw13).ArraySet1(p1, 9)
			(_nw13).ArraySet1(_dafny.IntOfInt64(686), 10)
			(_nw13).ArraySet1((_100_v62).F0(), 11)
			(_nw13).ArraySet1(p1, 12)
			(_nw13).ArraySet1((_103_v65).Cardinality(), 13)
			(_nw13).ArraySet1(p1, 14)
			_104_v66 = _nw13
			var _105_v67 _dafny.Map
			_ = _105_v67
			_105_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(667), p0)
			var _106_v68 _dafny.Map
			_ = _106_v68
			_106_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_75_v43, (_105_v67).Cardinality()), _48_cf47)
			var _rhs15 _dafny.Int = p1
			_ = _rhs15
			var _rhs16 bool = (func() bool {
				if (_106_v68).Contains(p1) {
					return (_106_v68).Get(p1).(bool)
				}
				return _48_cf47
			})()
			_ = _rhs16
			var _rhs17 bool = (p1).Cmp(_75_v43) > 0
			_ = _rhs17
			var _rhs18 _dafny.Array = _104_v66
			_ = _rhs18
			_75_v43 = _rhs15
			r0 = _rhs16
			_48_cf47 = _rhs17
			_104_v66 = _rhs18
			var _107_v69 _dafny.Sequence
			_ = _107_v69
			_107_v69 = _dafny.SeqOf(p0, p0, p0)
			var _108_v70 _dafny.Map
			_ = _108_v70
			_108_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v69, _48_cf47)
			_108_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_107_v69, _107_v69), p0)
		}
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_77_v45), 0))
		_ = _index3
		(_77_v45).ArraySet1(p0, (_index3).Int())
		var _109_v71 _dafny.MultiSet
		_ = _109_v71
		_109_v71 = _dafny.MultiSetOf(_48_cf47, p0, _48_cf47)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_77_v45), 0))
		_ = _index4
		(_77_v45).ArraySet1(((func() _dafny.Int {
			if (_109_v71).Contains(false) {
				return (_109_v71).Multiplicity(false)
			}
			return _75_v43
		})()).Cmp(Companion_Default___.Fm3(_dafny.SetOf(_48_cf47, _48_cf47, _48_cf47), _18_v0, p1, p1, globalState)) < 0, (_index4).Int())
	} else if _source0.Is_DC29() {
		var _110___mcc_h6 _dafny.CodePoint = _source0.Get_().(D13_DC29).Cf49
		_ = _110___mcc_h6
		var _111___mcc_h7 bool = _source0.Get_().(D13_DC29).Cf50
		_ = _111___mcc_h7
		var _112_cf50 bool = _111___mcc_h7
		_ = _112_cf50
		var _113_cf49 _dafny.CodePoint = _110___mcc_h6
		_ = _113_cf49
		var _114_v72 _dafny.Array
		_ = _114_v72
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
		_ = _nw14
		_114_v72 = _nw14
		var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
		_ = _nw15
		_114_v72 = _nw15
		var _115_v73 _dafny.Set
		_ = _115_v73
		_115_v73 = _dafny.SetOf(_112_cf50, p0, p0, !(_112_cf50), _112_cf50)
		var _116_v74 T1
		_ = _116_v74
		var _nw16 *C5 = New_C5_()
		_ = _nw16
		_nw16.Ctor__(_dafny.CodePoint('c'), Companion_Default___.Fm3(_115_v73, _18_v0, p1, p1, globalState), (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, p1))))
		_116_v74 = _nw16
		_116_v74 = (func() T1 {
			if _112_cf50 {
				return _116_v74
			}
			return _116_v74
		})()
		var _117_v75 _dafny.Array
		_ = _117_v75
		var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
		_ = _nw17
		_117_v75 = _nw17
		var _118_v76 _dafny.Map
		_ = _118_v76
		_118_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_116_v74).F0())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_117_v75), 0))
		_ = _index5
		(_117_v75).ArraySet1((_118_v76).Merge(_118_v76), (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_117_v75), 0))
		_ = _index6
		(_117_v75).ArraySet1((((_118_v76).Update(p1, (_116_v74).F0())).Merge(_118_v76)).Merge((_118_v76).Merge(_118_v76)), (_index6).Int())
		var _119_v77 _dafny.Array
		_ = _119_v77
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_1
		var _nw18 _dafny.Array
		_ = _nw18
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw18 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_120_cf50 bool) func(_dafny.Int) bool {
				return func(_121_i1 _dafny.Int) bool {
					return _120_cf50
				}
			})(_112_cf50)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw18 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw18).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw18).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_119_v77 = _nw18
		_119_v77 = _119_v77
	} else if _source0.Is_DC26() {
		var _122___mcc_h8 *C0 = _source0.Get_().(D13_DC26).Cf42
		_ = _122___mcc_h8
		var _123_cf42 *C0 = _122___mcc_h8
		_ = _123_cf42
		var _124_v78 *C1
		_ = _124_v78
		var _nw19 *C1 = New_C1_()
		_ = _nw19
		_nw19.Ctor__(p1)
		_124_v78 = _nw19
		var _125_v79 _dafny.Sequence
		_ = _125_v79
		_125_v79 = _dafny.SeqOf(_124_v78)
		var _126_v80 _dafny.Sequence
		_ = _126_v80
		_126_v80 = _dafny.SeqOf(_125_v79)
		_125_v79 = (_126_v80).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_126_v80).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _127_v81 _dafny.Int
		_ = _127_v81
		_127_v81 = _dafny.IntOfInt64(70)
		var _128_v82 _dafny.Set
		_ = _128_v82
		_128_v82 = _dafny.SetOf(false, p0, p0)
		_127_v81 = (_127_v81).Minus(Companion_Default___.Fm3(_128_v82, _18_v0, _127_v81, p1, globalState))
		r0 = p0
		_127_v81 = _127_v81
	} else {
		var _129___mcc_h9 D13 = _source0.Get_().(D13_DC30).Cf51
		_ = _129___mcc_h9
		var _130_cf51 D13 = _129___mcc_h9
		_ = _130_cf51
		r0 = !(p0)
		var _131_v83 _dafny.Array
		_ = _131_v83
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_2
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_132_p0 bool) func(_dafny.Int) bool {
				return func(_133_i2 _dafny.Int) bool {
					return !(_132_p0)
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
		_131_v83 = _nw20
		var _134_v84 _dafny.MultiSet
		_ = _134_v84
		_134_v84 = _dafny.MultiSetOf(p1)
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_131_v83), 0))
		_ = _index7
		(_131_v83).ArraySet1((_134_v84).IsSubsetOf(_134_v84), (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_131_v83), 0))
		_ = _index8
		(_131_v83).ArraySet1(p0, (_index8).Int())
		var _135_v85 _dafny.Array
		_ = _135_v85
		var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw21
		_135_v85 = _nw21
		var _136_v86 _dafny.Set
		_ = _136_v86
		_136_v86 = _dafny.SetOf(p0)
		var _137_v87 *C1
		_ = _137_v87
		var _nw22 *C1 = New_C1_()
		_ = _nw22
		_nw22.Ctor__(p1)
		_137_v87 = _nw22
		var _138_v88 _dafny.Sequence
		_ = _138_v88
		_138_v88 = _dafny.SeqOf(_137_v87, _137_v87)
		var _139_v89 _dafny.Sequence
		_ = _139_v89
		_139_v89 = _dafny.SeqOf(_131_v83)
		var _140_v90 _dafny.Array
		_ = _140_v90
		var _nwElement0_2 _dafny.Int = (_dafny.Zero).Minus(p1)
		_ = _nwElement0_2
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(17))
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_2, 0)
		(_nw23).ArraySet1(p1, 1)
		(_nw23).ArraySet1(p1, 2)
		(_nw23).ArraySet1(p1, 3)
		(_nw23).ArraySet1(p1, 4)
		(_nw23).ArraySet1(p1, 5)
		(_nw23).ArraySet1(p1, 6)
		(_nw23).ArraySet1(_dafny.IntOfInt64(122), 7)
		(_nw23).ArraySet1(p1, 8)
		(_nw23).ArraySet1(p1, 9)
		(_nw23).ArraySet1(p1, 10)
		(_nw23).ArraySet1((_136_v86).Cardinality(), 11)
		(_nw23).ArraySet1(_dafny.IntOfUint32((_138_v88).Cardinality()), 12)
		(_nw23).ArraySet1(p1, 13)
		(_nw23).ArraySet1(p1, 14)
		(_nw23).ArraySet1(p1, 15)
		(_nw23).ArraySet1(_dafny.IntOfUint32((_139_v89).Cardinality()), 16)
		_140_v90 = _nw23
		var _141_v91 _dafny.Map
		_ = _141_v91
		_141_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_131_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_131_v83), 0))).Int()).(bool), _140_v90)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_135_v85), 0))
		_ = _index9
		(_135_v85).ArraySet1(_141_v91, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_135_v85), 0))
		_ = _index10
		(_135_v85).ArraySet1(_141_v91, (_index10).Int())
		var _142_v92 _dafny.Int
		_ = _142_v92
		_142_v92 = _dafny.IntOfInt64(601)
		var _143_v93 _dafny.Sequence
		_ = _143_v93
		_143_v93 = _dafny.SeqOf(p1, _142_v92)
		var _144_v94 _dafny.Sequence
		_ = _144_v94
		_144_v94 = _dafny.SeqOf(_134_v84)
		_142_v92 = ((func() _dafny.MultiSet {
			if (_131_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_131_v83), 0))).Int()).(bool) {
				return _dafny.MultiSetFromSeq(_143_v93)
			}
			return (_144_v94).Select((Companion_Default___.SafeIndex(_142_v92, _dafny.IntOfUint32((_144_v94).Cardinality()))).Uint32()).(_dafny.MultiSet)
		})()).Cardinality()
	}
	var _145_v95 _dafny.Int
	_ = _145_v95
	_145_v95 = _dafny.IntOfInt64(898)
	var _rhs19 _dafny.Int = p1
	_ = _rhs19
	var _rhs20 bool = p0
	_ = _rhs20
	_145_v95 = _rhs19
	r0 = _rhs20
	var _146_v96 _dafny.Sequence
	_ = _146_v96
	_146_v96 = _dafny.UnicodeSeqOfUtf8Bytes("dwej")
	var _147_v97 _dafny.Sequence
	_ = _147_v97
	_147_v97 = _dafny.SeqOf(_145_v95)
	var _148_v98 _dafny.Sequence
	_ = _148_v98
	_148_v98 = _dafny.SeqOf(p0, Companion_Default___.Fm2(p0, _dafny.IntOfUint32((_147_v97).Cardinality()), p0, p0, globalState))
	var _149_v99 T1
	_ = _149_v99
	var _nw24 *C5 = New_C5_()
	_ = _nw24
	_nw24.Ctor__(_dafny.CodePoint('w'), _dafny.IntOfUint32((_148_v98).Cardinality()), _145_v95)
	_149_v99 = _nw24
	var _150_v100 _dafny.Sequence
	_ = _150_v100
	_150_v100 = _dafny.SeqOf(_149_v99)
	var _151_v101 _dafny.Set
	_ = _151_v101
	_151_v101 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v96, !(p0))).Cardinality(), p1, p1, _dafny.IntOfUint32((_150_v100).Cardinality()))
	if (_151_v101).IsProperSubsetOf(Companion_Default___.Fm34(globalState)) {
		if p0 {
			var _152_v102 _dafny.Set
			_ = _152_v102
			_152_v102 = _151_v101
			var _153_v103 _dafny.Sequence
			_ = _153_v103
			_153_v103 = _dafny.SeqOf(_152_v102, _151_v101, _151_v101)
			var _154_v104 _dafny.Map
			_ = _154_v104
			_154_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			_153_v103 = _dafny.Companion_Sequence_.Concatenate(_153_v103, Companion_Default___.Fm35((func() _dafny.Int {
				if (_154_v104).Contains(p0) {
					return (_154_v104).Get(p0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_146_v96).Cardinality())
			})(), (_149_v99).F0(), p0, globalState))
			r0 = p0
			var _155_v105 _dafny.Array
			_ = _155_v105
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw25
			_155_v105 = _nw25
			var _156_v106 *C0
			_ = _156_v106
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__(_18_v0)
			_156_v106 = _nw26
			var _157_v107 _dafny.Map
			_ = _157_v107
			_157_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v106, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(289))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_158_v99 T1) func(_dafny.Int) _dafny.Int {
				return func(_159_i3 _dafny.Int) _dafny.Int {
					return (_158_v99).F0()
				}
			})(_149_v99))))
			var _160_v108 _dafny.Map
			_ = _160_v108
			_160_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v107, _145_v95)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_155_v105), 0))
			_ = _index11
			(_155_v105).ArraySet1((func() _dafny.Int {
				if (_160_v108).Contains(_157_v107) {
					return (_160_v108).Get(_157_v107).(_dafny.Int)
				}
				return _dafny.IntOfInt64(376)
			})(), (_index11).Int())
			var _161_v109 _dafny.Array
			_ = _161_v109
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_3
			var _nw27 _dafny.Array
			_ = _nw27
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw27 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_162_p0 bool) func(_dafny.Int) bool {
					return func(_163_i4 _dafny.Int) bool {
						return _162_p0
					}
				})(p0)
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
			_161_v109 = _nw27
			var _164_v110 _dafny.Set
			_ = _164_v110
			_164_v110 = _dafny.SetOf(_161_v109)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_155_v105), 0))
			_ = _index12
			(_155_v105).ArraySet1(((_149_v99).F0()).Minus((_164_v110).Cardinality()), (_index12).Int())
			var _165_v111 T0
			_ = _165_v111
			var _nw28 *C3 = New_C3_()
			_ = _nw28
			_nw28.Ctor__((_149_v99).F0())
			_165_v111 = _nw28
			var _166_v112 _dafny.Map
			_ = _166_v112
			_166_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_155_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_155_v105), 0))).Int()).(_dafny.Int), _dafny.SetOf(_165_v111))
			_145_v95 = Companion_Default___.SafeModuloInt((_166_v112).Cardinality(), _dafny.IntOfInt64(223))
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_155_v105), 0))
			_ = _index13
			(_155_v105).ArraySet1((_149_v99).F0(), (_index13).Int())
		} else {
			_149_v99 = _149_v99
			var _167_v113 _dafny.Set
			_ = _167_v113
			_167_v113 = _dafny.SetOf(_149_v99)
			var _rhs21 _dafny.Sequence = _146_v96
			_ = _rhs21
			var _rhs22 bool = (func() bool {
				if false {
					return p0
				}
				return p0
			})()
			_ = _rhs22
			var _rhs23 _dafny.Int = (((_167_v113).Union(_167_v113)).Cardinality()).Minus((_149_v99).F0())
			_ = _rhs23
			_146_v96 = _rhs21
			r0 = _rhs22
			_145_v95 = _rhs23
			var _168_v114 _dafny.Array
			_ = _168_v114
			var _nw29 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw29
			_168_v114 = _nw29
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_168_v114), 0))
			_ = _index14
			(_168_v114).ArraySet1(_149_v99, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_168_v114), 0))
			_ = _index15
			(_168_v114).ArraySet1(_149_v99, (_index15).Int())
			_145_v95 = (_dafny.Zero).Minus(((_149_v99).F0()).Times(((_149_v99).F0()).Minus((_dafny.Zero).Minus((_149_v99).F0()))))
			_145_v95 = (_dafny.IntOfInt64(-90)).Minus((_149_v99).F0())
		}
		var _169_v115 _dafny.Map
		_ = _169_v115
		_169_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _170_v116 _dafny.Sequence
		_ = _170_v116
		_170_v116 = _dafny.SeqOf(_169_v115, _169_v115, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), (_169_v115).Merge(_169_v115))
		_170_v116 = _170_v116
		var _171_v117 *C3
		_ = _171_v117
		var _nw30 *C3 = New_C3_()
		_ = _nw30
		_nw30.Ctor__(_dafny.IntOfUint32((_148_v98).Cardinality()))
		_171_v117 = _nw30
		var _172_v118 _dafny.Set
		_ = _172_v118
		_172_v118 = _dafny.SetOf(p0, p0)
		_145_v95 = (_dafny.IntOfInt64(720)).Minus((Companion_Default___.Fm3(_172_v118, _18_v0, (func() _dafny.Set {
			var _coll11 = _dafny.NewBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(766), _dafny.IntOfInt64(816))); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _173_v119 _dafny.Int
				_173_v119 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(766)).Cmp(_173_v119) <= 0) && ((_173_v119).Cmp(_dafny.IntOfInt64(816)) < 0) {
					_coll11.Add(Companion_Default___.SafeModuloInt(_173_v119, _dafny.IntOfInt64(918)))
				}
			}
			return _coll11.ToSet()
		}()).Cardinality(), (_149_v99).F0(), globalState)).Minus((_149_v99).F0()))
		var _174_v120 *C0
		_ = _174_v120
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__(_dafny.CodePoint('m'))
		_174_v120 = _nw31
	} else {
		var _175_v121 _dafny.Array
		_ = _175_v121
		var _nw32 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
		_ = _nw32
		_175_v121 = _nw32
		var _176_v122 D3
		_ = _176_v122
		_176_v122 = Companion_D3_.Create_DC9_(_dafny.IntOfInt64(745), _175_v121, p0)
		var _177_v123 _dafny.Map
		_ = _177_v123
		_177_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
		var _178_v124 _dafny.Array
		_ = _178_v124
		var _nwElement0_3 bool = p0
		_ = _nwElement0_3
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(26))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_3, 0)
		(_nw33).ArraySet1(p0, 1)
		(_nw33).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_146_v96, _146_v96), 2)
		(_nw33).ArraySet1(false, 3)
		(_nw33).ArraySet1(p0, 4)
		(_nw33).ArraySet1(p0, 5)
		(_nw33).ArraySet1(p0, 6)
		(_nw33).ArraySet1(p0, 7)
		(_nw33).ArraySet1(p0, 8)
		(_nw33).ArraySet1(p0, 9)
		(_nw33).ArraySet1(p0, 10)
		(_nw33).ArraySet1((func() bool {
			if p0 {
				return p0
			}
			return true
		})(), 11)
		(_nw33).ArraySet1(p0, 12)
		(_nw33).ArraySet1(!_dafny.Companion_Sequence_.Contains(_146_v96, _18_v0), 13)
		(_nw33).ArraySet1(!(!(p0)) || (p0), 14)
		(_nw33).ArraySet1(p0, 15)
		(_nw33).ArraySet1(p0, 16)
		(_nw33).ArraySet1(p0, 17)
		(_nw33).ArraySet1(true, 18)
		(_nw33).ArraySet1((_176_v122).Dtor_cf12(), 19)
		(_nw33).ArraySet1((func() bool {
			if p0 {
				return true
			}
			return (func() bool {
				if (_177_v123).Contains(p0) {
					return (_177_v123).Get(p0).(bool)
				}
				return p0
			})()
		})(), 20)
		(_nw33).ArraySet1(p0, 21)
		(_nw33).ArraySet1((func() bool {
			if p0 {
				return p0
			}
			return p0
		})(), 22)
		(_nw33).ArraySet1(p0, 23)
		(_nw33).ArraySet1(p0, 24)
		(_nw33).ArraySet1(!(p0), 25)
		_178_v124 = _nw33
		_178_v124 = _178_v124
		if p0 {
			var _179_v125 _dafny.Map
			_ = _179_v125
			_179_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v0, _145_v95)
			var _180_v126 _dafny.MultiSet
			_ = _180_v126
			_180_v126 = _dafny.MultiSetOf(_179_v125)
			_180_v126 = _180_v126
			var _181_v127 *C4
			_ = _181_v127
			var _nw34 *C4 = New_C4_()
			_ = _nw34
			_nw34.Ctor__(_dafny.IntOfUint32((_148_v98).Cardinality()))
			_181_v127 = _nw34
			var _182_v128 _dafny.Sequence
			_ = _182_v128
			_182_v128 = _dafny.SeqOf(_181_v127)
			var _183_v129 _dafny.Sequence
			_ = _183_v129
			_183_v129 = _dafny.SeqOf(_182_v128)
			r0 = _dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
				if p0 {
					return (_183_v129).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_183_v129).Cardinality()))).Uint32()).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_181_v127)
			})(), _dafny.Companion_Sequence_.Update(_182_v128, (Companion_Default___.SafeIndex((_149_v99).F0(), _dafny.IntOfUint32((_182_v128).Cardinality()))).Uint32(), _181_v127))
			var _184_v130 _dafny.Array
			_ = _184_v130
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw35
			_184_v130 = _nw35
			var _185_v131 T0
			_ = _185_v131
			var _nw36 *C4 = New_C4_()
			_ = _nw36
			_nw36.Ctor__((_149_v99).F0())
			_185_v131 = _nw36
			var _186_v132 D10
			_ = _186_v132
			_186_v132 = Companion_D10_.Create_DC21_(_185_v131)
			var _187_v133 _dafny.Map
			_ = _187_v133
			_187_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v127, (_186_v132).Dtor_cf33())
			var _188_v134 _dafny.Array
			_ = _188_v134
			var _nwElement0_4 T0 = (func() T0 {
				if p0 {
					return _185_v131
				}
				return _185_v131
			})()
			_ = _nwElement0_4
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(18))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_4, 0)
			(_nw37).ArraySet1(_185_v131, 1)
			(_nw37).ArraySet1(_185_v131, 2)
			(_nw37).ArraySet1(_185_v131, 3)
			(_nw37).ArraySet1(_185_v131, 4)
			(_nw37).ArraySet1(_185_v131, 5)
			(_nw37).ArraySet1(_185_v131, 6)
			(_nw37).ArraySet1(_185_v131, 7)
			(_nw37).ArraySet1((func() T0 {
				if (_187_v133).Contains(_181_v127) {
					return (_187_v133).Get(_181_v127).(T0)
				}
				return _185_v131
			})(), 8)
			(_nw37).ArraySet1((func() T0 {
				if p0 {
					return _185_v131
				}
				return _185_v131
			})(), 9)
			(_nw37).ArraySet1(_185_v131, 10)
			(_nw37).ArraySet1(_185_v131, 11)
			(_nw37).ArraySet1(_185_v131, 12)
			(_nw37).ArraySet1(_185_v131, 13)
			(_nw37).ArraySet1(_185_v131, 14)
			(_nw37).ArraySet1(_185_v131, 15)
			(_nw37).ArraySet1(_185_v131, 16)
			(_nw37).ArraySet1(_185_v131, 17)
			_188_v134 = _nw37
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_188_v134), 0))
			_ = _index16
			(_188_v134).ArraySet1(_185_v131, (_index16).Int())
			var _189_v135 _dafny.Map
			_ = _189_v135
			_189_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v121, p0)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_188_v134), 0))
			_ = _index17
			var _rhs24 _dafny.Array = _184_v130
			_ = _rhs24
			var _rhs25 _dafny.CodePoint = Companion_Default___.Fm17(((_189_v135).Merge(_189_v135)).Cardinality(), globalState)
			_ = _rhs25
			var _rhs26 _dafny.Int = (_149_v99).F0()
			_ = _rhs26
			var _rhs27 T0 = _185_v131
			_ = _rhs27
			var _lhs0 _dafny.Array = _188_v134
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_188_v134), 0))
			_ = _lhs1
			_184_v130 = _rhs24
			_18_v0 = _rhs25
			_145_v95 = _rhs26
			(_lhs0).ArraySet1(_rhs27, (_lhs1).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_178_v124), 0))
			_ = _index18
			(_178_v124).ArraySet1(!(!(p0)), (_index18).Int())
			var _190_v136 _dafny.Map
			_ = _190_v136
			_190_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v99, (_149_v99).F0())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_178_v124), 0))
			_ = _index19
			(_178_v124).ArraySet1((_190_v136).Equals(_190_v136), (_index19).Int())
			r0 = !(!((func() bool {
				if p0 {
					return p0
				}
				return p0
			})()))
		} else {
			_18_v0 = _18_v0
			var _191_v137 _dafny.Set
			_ = _191_v137
			_191_v137 = _dafny.SetOf((_149_v99).Fm6(p0, _dafny.UnicodeSeqOfUtf8Bytes("cqrcjrmk"), _146_v96, (_149_v99).F0(), globalState), p0, p0, p0, p0)
			_145_v95 = Companion_Default___.Fm3(_191_v137, _18_v0, Companion_Default___.SafeModuloInt(p1, (_149_v99).F0()), p1, globalState)
			r0 = p0
			_145_v95 = (_149_v99).F0()
			var _192_v139 _dafny.Array
			_ = _192_v139
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_4
			var _nw38 _dafny.Array
			_ = _nw38
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw38 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Set = (func(_193_v99 T1) func(_dafny.Int) _dafny.Set {
					return func(_194_i5 _dafny.Int) _dafny.Set {
						return func() _dafny.Set {
							var _coll12 = _dafny.NewBuilder()
							_ = _coll12
							for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_193_v99).F0(), true)).Keys().Elements()); ; {
								_compr_12, _ok12 := _iter12()
								if !_ok12 {
									break
								}
								var _195_v138 _dafny.Int
								_195_v138 = interface{}(_compr_12).(_dafny.Int)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_193_v99).F0(), true)).Contains(_195_v138) {
									_coll12.Add((_195_v138).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()))
								}
							}
							return _coll12.ToSet()
						}()
					}
				})(_149_v99)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw38 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw38).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw38).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_192_v139 = _nw38
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_192_v139), 0))
			_ = _index20
			(_192_v139).ArraySet1(_151_v101, (_index20).Int())
			var _196_v140 D16
			_ = _196_v140
			_196_v140 = Companion_D16_.Create_DC35_((_149_v99).F0())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_192_v139), 0))
			_ = _index21
			var _rhs28 _dafny.Set = Companion_Default___.Fm34(globalState)
			_ = _rhs28
			var _rhs29 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_197_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_198_i6 _dafny.Int) _dafny.CodePoint {
					return _197_v0
				}
			})(_18_v0))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_199_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_200_i6 _dafny.Int) _dafny.CodePoint {
					return _199_v0
				}
			})(_18_v0)))).Cardinality()))).Uint32(), (_146_v96).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_196_v140).Dtor_cf55()), _dafny.IntOfUint32((_146_v96).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_ = _rhs29
			var _lhs2 _dafny.Array = _192_v139
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_192_v139), 0))
			_ = _lhs3
			(_lhs2).ArraySet1(_rhs28, (_lhs3).Int())
			_146_v96 = _rhs29
		}
		var _201_v141 D3
		_ = _201_v141
		_201_v141 = Companion_D3_.Create_DC8_(p0, (_149_v99).F0())
		var _202_v142 T0
		_ = _202_v142
		var _nw39 *C1 = New_C1_()
		_ = _nw39
		_nw39.Ctor__((_149_v99).F0())
		_202_v142 = _nw39
		var _203_v143 _dafny.Sequence
		_ = _203_v143
		_203_v143 = _dafny.SeqOf(_202_v142)
		var _204_v144 _dafny.MultiSet
		_ = _204_v144
		_204_v144 = _dafny.MultiSetOf(_202_v142, _202_v142, (_203_v143).Select((Companion_Default___.SafeIndex(_145_v95, _dafny.IntOfUint32((_203_v143).Cardinality()))).Uint32()).(T0))
		var _205_v145 _dafny.Map
		_ = _205_v145
		_205_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v141, !((_dafny.MultiSetOf(_202_v142, _202_v142, _202_v142)).IsSubsetOf(_204_v144)))
		_205_v145 = (_205_v145).Update(_201_v141, p0)
		var _206_v146 *C4
		_ = _206_v146
		var _nw40 *C4 = New_C4_()
		_ = _nw40
		_nw40.Ctor__(((_149_v99).F0()).Times((_149_v99).F0()))
		_206_v146 = _nw40
		var _207_v147 *C2
		_ = _207_v147
		var _nw41 *C2 = New_C2_()
		_ = _nw41
		_nw41.Ctor__(_145_v95)
		_207_v147 = _nw41
	}
	var _208_v148 *C2
	_ = _208_v148
	var _nw42 *C2 = New_C2_()
	_ = _nw42
	_nw42.Ctor__(_dafny.IntOfUint32((_146_v96).Cardinality()))
	_208_v148 = _nw42
	var _209_v149 _dafny.Map
	_ = _209_v149
	_209_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_145_v95).Cmp(p1) != 0, _208_v148)
	var _210_v150 D6
	_ = _210_v150
	_210_v150 = Companion_D6_.Create_DC15_((_dafny.Zero).Minus(_145_v95), p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_149_v99).F0())).Cardinality())
	var _211_v151 _dafny.Set
	_ = _211_v151
	_211_v151 = _dafny.SetOf(p0)
	var _212_v152 D5
	_ = _212_v152
	_212_v152 = Companion_D5_.Create_DC13_((_149_v99).F0(), _211_v151, _18_v0, _145_v95, p0)
	var _213_v153 _dafny.Sequence
	_ = _213_v153
	_213_v153 = _dafny.SeqOf((func() D6 {
		if p0 {
			return _210_v150
		}
		return _210_v150
	})(), Companion_D6_.Create_DC15_((_212_v152).Dtor_cf20(), p0, _dafny.IntOfInt64(898)), _210_v150, _210_v150, _210_v150)
	var _214_v154 _dafny.Array
	_ = _214_v154
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_5
	var _nw43 _dafny.Array
	_ = _nw43
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw43 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Int = (func(_215_v95 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_216_i7 _dafny.Int) _dafny.Int {
				return (_216_i7).Times(_215_v95)
			}
		})(_145_v95)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw43 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw43).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw43).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_214_v154 = _nw43
	var _217_v155 _dafny.Sequence
	_ = _217_v155
	_217_v155 = _dafny.SeqOf(_214_v154, _214_v154)
	var _218_v156 _dafny.Map
	_ = _218_v156
	_218_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _214_v154)
	var _219_v157 _dafny.Array
	_ = _219_v157
	var _nwElement0_5 _dafny.Array = (_217_v155).Select((Companion_Default___.SafeIndex((_147_v97).Select((Companion_Default___.SafeIndex((_149_v99).F0(), _dafny.IntOfUint32((_147_v97).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_217_v155).Cardinality()))).Uint32()).(_dafny.Array)
	_ = _nwElement0_5
	var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(23))
	_ = _nw44
	(_nw44).ArraySet1(_nwElement0_5, 0)
	(_nw44).ArraySet1((func() _dafny.Array {
		if p0 {
			return _214_v154
		}
		return _214_v154
	})(), 1)
	(_nw44).ArraySet1(_214_v154, 2)
	(_nw44).ArraySet1(_214_v154, 3)
	(_nw44).ArraySet1(_214_v154, 4)
	(_nw44).ArraySet1(_214_v154, 5)
	(_nw44).ArraySet1(_214_v154, 6)
	(_nw44).ArraySet1(_214_v154, 7)
	(_nw44).ArraySet1(_214_v154, 8)
	(_nw44).ArraySet1((_217_v155).Select((Companion_Default___.SafeIndex((_211_v151).Cardinality(), _dafny.IntOfUint32((_217_v155).Cardinality()))).Uint32()).(_dafny.Array), 9)
	(_nw44).ArraySet1((func() _dafny.Array {
		if p0 {
			return _214_v154
		}
		return _214_v154
	})(), 10)
	(_nw44).ArraySet1(_214_v154, 11)
	(_nw44).ArraySet1(_214_v154, 12)
	(_nw44).ArraySet1((func() _dafny.Array {
		if (_218_v156).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_147_v97).Cardinality()))) {
			return (_218_v156).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_147_v97).Cardinality()))).(_dafny.Array)
		}
		return _214_v154
	})(), 13)
	(_nw44).ArraySet1(_214_v154, 14)
	(_nw44).ArraySet1(_214_v154, 15)
	(_nw44).ArraySet1(_214_v154, 16)
	(_nw44).ArraySet1(_214_v154, 17)
	(_nw44).ArraySet1(_214_v154, 18)
	(_nw44).ArraySet1(_214_v154, 19)
	(_nw44).ArraySet1(_214_v154, 20)
	(_nw44).ArraySet1(_214_v154, 21)
	(_nw44).ArraySet1(_214_v154, 22)
	_219_v157 = _nw44
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))
	_ = _index22
	(_219_v157).ArraySet1(_214_v154, (_index22).Int())
	var _220_v158 _dafny.Map
	_ = _220_v158
	_220_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_147_v97).Cardinality()), _209_v149)
	var _221_v159 _dafny.MultiSet
	_ = _221_v159
	_221_v159 = _dafny.MultiSetOf(p1)
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))
	_ = _index23
	var _rhs30 _dafny.Map = (_209_v149).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _208_v148)).Merge((func() _dafny.Map {
		if (_220_v158).Contains(_dafny.IntOfUint32((Companion_Default___.Fm19(_221_v159, (_211_v151).Cardinality(), p1, _145_v95, globalState)).Cardinality())) {
			return (_220_v158).Get(_dafny.IntOfUint32((Companion_Default___.Fm19(_221_v159, (_211_v151).Cardinality(), p1, _145_v95, globalState)).Cardinality())).(_dafny.Map)
		}
		return _209_v149
	})()))
	_ = _rhs30
	var _rhs31 _dafny.Sequence = _213_v153
	_ = _rhs31
	var _rhs32 _dafny.Array = _214_v154
	_ = _rhs32
	var _lhs4 _dafny.Array = _219_v157
	_ = _lhs4
	var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))
	_ = _lhs5
	_209_v149 = _rhs30
	_213_v153 = _rhs31
	(_lhs4).ArraySet1(_rhs32, (_lhs5).Int())
	var _222_i8 _dafny.Int
	_ = _222_i8
	_222_i8 = _dafny.Zero
	{
		for (p1).Cmp(((_dafny.Zero).Minus(_145_v95)).Minus(_145_v95)) <= 0 {
			{
				if (_222_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_222_i8 = (_222_i8).Plus(_dafny.One)
				var _223_v160 T0
				_ = _223_v160
				var _nw45 *C4 = New_C4_()
				_ = _nw45
				_nw45.Ctor__(_145_v95)
				_223_v160 = _nw45
				_145_v95 = (Companion_D13_.Create_DC27_(_18_v0, _223_v160, _18_v0, _dafny.IntOfUint32((_146_v96).Cardinality()))).Dtor_cf46()
				var _224_v161 _dafny.MultiSet
				_ = _224_v161
				_224_v161 = _dafny.MultiSetOf(p0, true)
				var _225_v162 _dafny.Map
				_ = _225_v162
				_225_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_224_v161).IsProperSubsetOf(_224_v161), _211_v151)
				var _226_v163 D7
				_ = _226_v163
				_226_v163 = Companion_D7_.Create_DC18_((_223_v160).F0())
				var _rhs33 _dafny.Int = (_223_v160).F0()
				_ = _rhs33
				var _rhs34 _dafny.Int = Companion_Default___.Fm3(_211_v151, _18_v0, (_149_v99).F0(), (_226_v163).Dtor_cf30(), globalState)
				_ = _rhs34
				var _rhs35 _dafny.Map = _225_v162
				_ = _rhs35
				var _rhs36 bool = _dafny.Companion_Sequence_.IsPrefixOf(_148_v98, _dafny.Companion_Sequence_.Concatenate(_148_v98, _dafny.SeqOf(p0)))
				_ = _rhs36
				_145_v95 = _rhs33
				_145_v95 = _rhs34
				_225_v162 = _rhs35
				r0 = _rhs36
				_145_v95 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_146_v96, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uvrdr"), _146_v96))).Cardinality())
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))
				_ = _arr0
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))), 0))
				_ = _index24
				(_arr0).ArraySet1((_149_v99).F0(), (_index24).Int())
				var _227_v164 _dafny.Map
				_ = _227_v164
				_227_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_211_v151, _18_v0, p1, (_149_v99).F0(), globalState), p0)
				var _228_v165 *C1
				_ = _228_v165
				var _nw46 *C1 = New_C1_()
				_ = _nw46
				_nw46.Ctor__((_227_v164).Cardinality())
				_228_v165 = _nw46
				var _229_v166 _dafny.MultiSet
				_ = _229_v166
				_229_v166 = _dafny.MultiSetOf(_228_v165)
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))
				_ = _arr1
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))), 0))
				_ = _index25
				(_arr1).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if p0 {
						return (func() _dafny.Int {
							if (_229_v166).Contains(_228_v165) {
								return (_229_v166).Multiplicity(_228_v165)
							}
							return (_149_v99).F0()
						})()
					}
					return (_149_v99).F0()
				})(), (_dafny.Zero).Minus((_dafny.IntOfUint32((_148_v98).Cardinality())).Plus(_145_v95))), (_index25).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _230_v167 _dafny.Sequence
	_ = _230_v167
	_230_v167 = _dafny.SeqOf(_221_v159)
	var _hi0 _dafny.Int = _145_v95
	_ = _hi0
	for _231_i9 := ((_230_v167).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_230_v167).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(); _231_i9.Cmp(_hi0) < 0; _231_i9 = _231_i9.Plus(_dafny.One) {
		var _232_v168 _dafny.Array
		_ = _232_v168
		var _nwElement0_6 bool = p0
		_ = _nwElement0_6
		var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(2))
		_ = _nw47
		(_nw47).ArraySet1(_nwElement0_6, 0)
		(_nw47).ArraySet1(p0, 1)
		_232_v168 = _nw47
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))
		_ = _index26
		(_232_v168).ArraySet1(p0, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))
		_ = _index27
		(_232_v168).ArraySet1(false, (_index27).Int())
		if (_232_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))).Int()).(bool) {
			var _233_v169 _dafny.Sequence
			_ = _233_v169
			_233_v169 = _dafny.SeqOf(_146_v96)
			var _234_v170 _dafny.Set
			_ = _234_v170
			_234_v170 = _dafny.SetOf(_146_v96, (func() _dafny.Sequence {
				if false {
					return _146_v96
				}
				return _146_v96
			})(), _dafny.Companion_Sequence_.Concatenate(_146_v96, _146_v96), (_233_v169).Select((Companion_Default___.SafeIndex(_231_i9, _dafny.IntOfUint32((_233_v169).Cardinality()))).Uint32()).(_dafny.Sequence), _146_v96)
			_234_v170 = _234_v170
			_145_v95 = (func() _dafny.Int {
				if (_221_v159).Contains((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_148_v98, (Companion_Default___.SafeIndex(_231_i9, _dafny.IntOfUint32((_148_v98).Cardinality()))).Uint32(), (_232_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))).Int()).(bool))).Cardinality())).Plus(_231_i9)) {
					return (_221_v159).Multiplicity((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_148_v98, (Companion_Default___.SafeIndex(_231_i9, _dafny.IntOfUint32((_148_v98).Cardinality()))).Uint32(), (_232_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))).Int()).(bool))).Cardinality())).Plus(_231_i9))
				}
				return (func() _dafny.Int {
					if false {
						return (_151_v101).Cardinality()
					}
					return Companion_Default___.Fm3(_211_v151, _18_v0, _dafny.IntOfUint32((_146_v96).Cardinality()), (_149_v99).F0(), globalState)
				})()
			})()
			var _235_v171 _dafny.Sequence
			_ = _235_v171
			_235_v171 = _dafny.SeqOf((func() _dafny.Array {
				if (_232_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))).Int()).(bool) {
					return _232_v168
				}
				return _232_v168
			})(), _232_v168)
			_235_v171 = _dafny.SeqOf(_232_v168)
			var _236_v172 _dafny.MultiSet
			_ = _236_v172
			_236_v172 = _dafny.MultiSetOf(_232_v168)
			_236_v172 = ((_236_v172).Union(_dafny.MultiSetOf(_232_v168))).Union(_236_v172)
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))
			_ = _arr2
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))), 0))
			_ = _index28
			(_arr2).ArraySet1(((_149_v99).F0()).Minus((_147_v97).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_147_v97).Cardinality()))).Uint32()).(_dafny.Int)), (_index28).Int())
			var _237_v173 T0
			_ = _237_v173
			var _nw48 *C1 = New_C1_()
			_ = _nw48
			_nw48.Ctor__((_149_v99).F0())
			_237_v173 = _nw48
			var _238_v174 D13
			_ = _238_v174
			_238_v174 = Companion_D13_.Create_DC27_(_dafny.CodePoint('w'), _237_v173, _dafny.CodePoint('o'), (_149_v99).F0())
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))
			_ = _arr3
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_dafny.ArrayCastTo((_219_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_219_v157), 0))).Int()))), 0))
			_ = _index29
			(_arr3).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(p1, (_149_v99).F0()), (_238_v174).Dtor_cf46()), (_index29).Int())
		} else {
			_145_v95 = (_149_v99).F0()
			var _239_v175 _dafny.Map
			_ = _239_v175
			_239_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_18_v0), (_232_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))).Int()).(bool))
			var _240_v176 _dafny.MultiSet
			_ = _240_v176
			_240_v176 = _dafny.MultiSetOf(_18_v0, _18_v0)
			var _241_v178 _dafny.Sequence
			_ = _241_v178
			_241_v178 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(788))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_242_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_243_i10 _dafny.Int) _dafny.CodePoint {
					return _242_v0
				}
			})(_18_v0))), _146_v96, _146_v96, _dafny.UnicodeSeqOfUtf8Bytes("hdjof"), _dafny.UnicodeSeqOfUtf8Bytes("edovto"))
			_239_v175 = (_239_v175).Update((_240_v176).Difference(_240_v176), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v96, _148_v98)).Equals(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_241_v178).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _244_v177 _dafny.Sequence
					_244_v177 = interface{}(_compr_13).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_241_v178, _244_v177) {
						_coll13.Add(_244_v177, _dafny.SeqOf((_232_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))).Int()).(bool)))
					}
				}
				return _coll13.ToMap()
			}()))
			var _245_v179 _dafny.Map
			_ = _245_v179
			_245_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(71), _146_v96)
			_145_v95 = ((_149_v99).F0()).Times(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if p0 {
					return (func() _dafny.Sequence {
						if (_245_v179).Contains((_149_v99).F0()) {
							return (_245_v179).Get((_149_v99).F0()).(_dafny.Sequence)
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("lv")
					})()
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-667))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}((func(_246_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_247_i11 _dafny.Int) _dafny.CodePoint {
						return _246_v0
					}
				})(_18_v0)))
			})()).Cardinality()))
			var _248_v180 *C5
			_ = _248_v180
			var _nw49 *C5 = New_C5_()
			_ = _nw49
			_nw49.Ctor__((func() _dafny.CodePoint {
				if false {
					return _18_v0
				}
				return _18_v0
			})(), _dafny.IntOfUint32((_146_v96).Cardinality()), _145_v95)
			_248_v180 = _nw49
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_232_v168), 0))
			_ = _index30
			(_232_v168).ArraySet1(p0, (_index30).Int())
		}
		var _249_v181 _dafny.Array
		_ = _249_v181
		var _nw50 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
		_ = _nw50
		_249_v181 = _nw50
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_249_v181), 0))
		_ = _index31
		(_249_v181).ArraySet1(_208_v148, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_249_v181), 0))
		_ = _index32
		(_249_v181).ArraySet1((func() *C2 {
			if ((_149_v99).F0()).Cmp((_dafny.Zero).Minus(_231_i9)) < 0 {
				return _208_v148
			}
			return _208_v148
		})(), (_index32).Int())
		var _250_v182 _dafny.Sequence
		_ = _250_v182
		_250_v182 = _dafny.SeqOf(_232_v168, _232_v168, _232_v168, _232_v168)
		_232_v168 = (_250_v182).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_149_v99).F0()), _dafny.IntOfUint32((_250_v182).Cardinality()))).Uint32()).(_dafny.Array)
	}
	r0 = p0
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _251_globalState *GlobalState
	_ = _251_globalState
	var _nw51 *GlobalState = New_GlobalState_()
	_ = _nw51
	_nw51.Ctor__()
	_251_globalState = _nw51
	var _252_v0 _dafny.Array
	_ = _252_v0
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_6
	var _nw52 _dafny.Array
	_ = _nw52
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw52 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = func(_253_i0 _dafny.Int) bool {
			return false
		}
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
	_252_v0 = _nw52
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
	_ = _index33
	(_252_v0).ArraySet1((_dafny.IntOfUint32((Companion_Default___.Fm0(_251_globalState)).Cardinality())).Cmp(_dafny.IntOfInt64(716)) <= 0, (_index33).Int())
	var _254_v1 bool
	_ = _254_v1
	_254_v1 = false
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_252_v0), 0))
	_ = _index34
	(_252_v0).ArraySet1(_254_v1, (_index34).Int())
	var _255_v2 _dafny.Int
	_ = _255_v2
	_255_v2 = _dafny.IntOfInt64(-806)
	var _256_v3 _dafny.Array
	_ = _256_v3
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(28)
	_ = _len0_7
	var _nw53 _dafny.Array
	_ = _nw53
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw53 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Int = (func(_257_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_258_i1 _dafny.Int) _dafny.Int {
				return (_258_i1).Minus(_257_v2)
			}
		})(_255_v2)
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
	_256_v3 = _nw53
	var _259_v4 _dafny.MultiSet
	_ = _259_v4
	_259_v4 = _dafny.MultiSetOf(_256_v3, _256_v3, _256_v3, _256_v3)
	var _260_v5 _dafny.MultiSet
	_ = _260_v5
	_260_v5 = _dafny.MultiSetOf(_255_v2)
	var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
	_ = _index35
	var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_252_v0), 0))
	_ = _index36
	var _rhs37 bool = (true) || ((_dafny.MultiSetOf(_255_v2)).IsProperSubsetOf(_260_v5))
	_ = _rhs37
	var _rhs38 _dafny.Array = (func() _dafny.Array {
		if _254_v1 {
			return _252_v0
		}
		return _252_v0
	})()
	_ = _rhs38
	var _rhs39 bool = _254_v1
	_ = _rhs39
	var _rhs40 _dafny.Int = _255_v2
	_ = _rhs40
	var _rhs41 _dafny.MultiSet = _259_v4
	_ = _rhs41
	var _lhs6 _dafny.Array = _252_v0
	_ = _lhs6
	var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
	_ = _lhs7
	var _lhs8 _dafny.Array = _252_v0
	_ = _lhs8
	var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_252_v0), 0))
	_ = _lhs9
	(_lhs6).ArraySet1(_rhs37, (_lhs7).Int())
	_252_v0 = _rhs38
	(_lhs8).ArraySet1(_rhs39, (_lhs9).Int())
	_255_v2 = _rhs40
	_259_v4 = _rhs41
	if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
		_255_v2 = _255_v2
		var _261_v6 _dafny.CodePoint
		_ = _261_v6
		_261_v6 = _dafny.CodePoint('p')
		_255_v2 = _dafny.IntOfUint32((Companion_Default___.Fm1(_261_v6, _251_globalState)).Cardinality())
		if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
			_255_v2 = _255_v2
			var _262_v7 _dafny.Map
			_ = _262_v7
			_262_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_254_v1, (_dafny.SetOf(!(_254_v1))).Cardinality())
			_262_v7 = (_262_v7).Update(_254_v1, (_260_v5).Cardinality())
			_261_v6 = _261_v6
			var _263_v8 _dafny.Sequence
			_ = _263_v8
			_263_v8 = _dafny.UnicodeSeqOfUtf8Bytes("qqyjceeg")
			var _264_v9 bool
			_ = _264_v9
			var _out0 bool
			_ = _out0
			_out0 = Companion_Default___.M0(_254_v1, _dafny.IntOfUint32((_263_v8).Cardinality()), _251_globalState)
			_264_v9 = _out0
			var _265_v10 _dafny.Sequence
			_ = _265_v10
			_265_v10 = _dafny.SeqOf(false)
			var _266_v11 D0
			_ = _266_v11
			_266_v11 = Companion_D0_.Create_DC1_((_265_v10).Select((Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_265_v10).Cardinality()))).Uint32()).(bool))
			var _267_v12 _dafny.Map
			_ = _267_v12
			_267_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_254_v1, _264_v9)
			_264_v9 = (func() bool {
				if (!((_266_v11).Dtor_cf1())) && ((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)) {
					return (func() bool {
						if _254_v1 {
							return (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
						}
						return _254_v1
					})()
				}
				return !(Companion_Default___.Fm2(_264_v9, _255_v2, !(true), Companion_Default___.Fm2((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _dafny.IntOfInt64(888), _254_v1, (func() bool {
					if (_267_v12).Contains((_265_v10).Select((Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_265_v10).Cardinality()))).Uint32()).(bool)) {
						return (_267_v12).Get((_265_v10).Select((Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_265_v10).Cardinality()))).Uint32()).(bool)).(bool)
					}
					return _264_v9
				})(), _251_globalState), _251_globalState))
			})()
		} else {
			var _268_v13 bool
			_ = _268_v13
			var _out1 bool
			_ = _out1
			_out1 = Companion_Default___.M0((func() bool {
				if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
					return (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
				}
				return (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
			})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(725))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_269_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_270_i2 _dafny.Int) _dafny.CodePoint {
					return _269_v6
				}
			})(_261_v6)))).Cardinality()), _251_globalState)
			_268_v13 = _out1
			var _271_v14 _dafny.Array
			_ = _271_v14
			var _nwElement0_7 _dafny.Array = _252_v0
			_ = _nwElement0_7
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(3))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_7, 0)
			(_nw54).ArraySet1(_252_v0, 1)
			(_nw54).ArraySet1(_252_v0, 2)
			_271_v14 = _nw54
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_271_v14), 0))
			_ = _index37
			(_271_v14).ArraySet1(_252_v0, (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_271_v14), 0))
			_ = _index38
			(_271_v14).ArraySet1(_252_v0, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
			_ = _index39
			(_252_v0).ArraySet1((_268_v13) == (_268_v13), (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
			_ = _index40
			(_252_v0).ArraySet1(_254_v1, (_index40).Int())
			_268_v13 = !((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
		}
		var _272_v15 _dafny.Set
		_ = _272_v15
		_272_v15 = _dafny.SetOf(_261_v6)
		var _273_v16 _dafny.Sequence
		_ = _273_v16
		_273_v16 = _dafny.SeqOf(_255_v2, _255_v2)
		var _274_v17 _dafny.Map
		_ = _274_v17
		_274_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_275_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_276_i3 _dafny.Int) _dafny.CodePoint {
				return _275_v6
			}
		})(_261_v6))), _dafny.IntOfUint32((_273_v16).Cardinality()))
		_254_v1 = ((_272_v15).Cardinality()).Cmp((_274_v17).Cardinality()) <= 0
		var _277_v18 _dafny.Map
		_ = _277_v18
		_277_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v2, _255_v2)
		var _278_v19 _dafny.Map
		_ = _278_v19
		_278_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_277_v18).Contains(_255_v2) {
				return (_277_v18).Get(_255_v2).(_dafny.Int)
			}
			return _255_v2
		})(), !(_254_v1))
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _index41
		(_252_v0).ArraySet1((func() bool {
			if _254_v1 {
				return !((func() bool {
					if (_278_v19).Contains(_255_v2) {
						return (_278_v19).Get(_255_v2).(bool)
					}
					return _254_v1
				})()) || (_254_v1)
			}
			return _254_v1
		})(), (_index41).Int())
	} else {
		var _279_v20 _dafny.CodePoint
		_ = _279_v20
		_279_v20 = _dafny.CodePoint('o')
		var _280_v21 _dafny.Sequence
		_ = _280_v21
		_280_v21 = _dafny.SeqOf((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
		var _281_v22 bool
		_ = _281_v22
		var _out2 bool
		_ = _out2
		_out2 = Companion_Default___.M0((_255_v2).Cmp(Companion_Default___.Fm3(_dafny.SetOf((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)), _279_v20, _255_v2, _255_v2, _251_globalState)) > 0, (_dafny.SetOf(_255_v2, _dafny.IntOfUint32((_280_v21).Cardinality()))).Cardinality(), _251_globalState)
		_281_v22 = _out2
		if (func() bool {
			if _254_v1 {
				return (_dafny.IntOfUint32((_280_v21).Cardinality())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_280_v21).Cardinality()))) <= 0
			}
			return _281_v22
		})() {
			var _282_v23 _dafny.Map
			_ = _282_v23
			_282_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)), _281_v22)
			var _283_v24 _dafny.Set
			_ = _283_v24
			_283_v24 = _dafny.SetOf(_281_v22)
			var _rhs42 _dafny.MultiSet = ((_260_v5).Union(_dafny.MultiSetOf(Companion_Default___.Fm3(_283_v24, _279_v20, _255_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-155))).Uint32(), func(coer24 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_284_v1 bool) func(_dafny.Int) D0 {
				return func(_285_i4 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_284_v1)
				}
			})(_254_v1)))).Cardinality()), _251_globalState), _255_v2))).Union((_260_v5).Update(_255_v2, Companion_Default___.Abs(_255_v2)))
			_ = _rhs42
			var _rhs43 _dafny.Map = Companion_Default___.Fm4(_255_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}(func(_286_i5 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-995)
			}))).Cardinality()), _dafny.IntOfInt64(185), _251_globalState)
			_ = _rhs43
			_260_v5 = _rhs42
			_282_v23 = _rhs43
			_255_v2 = (func() _dafny.Int {
				if _281_v22 {
					return _255_v2
				}
				return _255_v2
			})()
			_279_v20 = _279_v20
			var _287_v25 _dafny.MultiSet
			_ = _287_v25
			_287_v25 = _dafny.MultiSetOf(_283_v24)
			_255_v2 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(547), (_287_v25).Cardinality())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
			_ = _index42
			(_252_v0).ArraySet1((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (_index42).Int())
		} else {
			_255_v2 = (((_260_v5).Difference(_260_v5)).Cardinality()).Times(_255_v2)
			_281_v22 = (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
			_255_v2 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(989))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_288_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_289_i6 _dafny.Int) _dafny.Int {
					return _288_v2
				}
			})(_255_v2)))).Cardinality()), _255_v2)
			var _290_v26 _dafny.Sequence
			_ = _290_v26
			_290_v26 = _dafny.SeqOf(_255_v2, _255_v2)
			var _291_v27 _dafny.MultiSet
			_ = _291_v27
			_291_v27 = _dafny.MultiSetOf(_254_v1, (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _254_v1)
			var _292_v28 _dafny.Set
			_ = _292_v28
			_292_v28 = _dafny.SetOf((_291_v27).Cardinality())
			var _293_v29 _dafny.Sequence
			_ = _293_v29
			_293_v29 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_290_v26, (Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_290_v26).Cardinality()))).Uint32(), (_292_v28).Cardinality()), _290_v26), _dafny.Companion_Sequence_.Concatenate(_290_v26, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_294_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_295_i7 _dafny.Int) _dafny.Int {
					return _294_v2
				}
			})(_255_v2)))))
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
			_ = _index43
			var _rhs44 bool = !(_254_v1)
			_ = _rhs44
			var _rhs45 _dafny.Int = _dafny.IntOfUint32((_293_v29).Cardinality())
			_ = _rhs45
			var _lhs10 _dafny.Array = _252_v0
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
			_ = _lhs11
			(_lhs10).ArraySet1(_rhs44, (_lhs11).Int())
			_255_v2 = _rhs45
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_256_v3), 0))
			_ = _index44
			(_256_v3).ArraySet1((func() _dafny.Int {
				if (_260_v5).Contains(_255_v2) {
					return (_260_v5).Multiplicity(_255_v2)
				}
				return (_dafny.Zero).Minus(_255_v2)
			})(), (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_256_v3), 0))
			_ = _index45
			(_256_v3).ArraySet1(_255_v2, (_index45).Int())
		}
		var _296_v30 _dafny.Sequence
		_ = _296_v30
		_296_v30 = _dafny.SeqOf(_255_v2)
		var _297_v31 bool
		_ = _297_v31
		var _out3 bool
		_ = _out3
		_out3 = Companion_Default___.M0((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (_296_v30).Select((Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_296_v30).Cardinality()))).Uint32()).(_dafny.Int), _251_globalState)
		_297_v31 = _out3
		var _298_v32 _dafny.Array
		_ = _298_v32
		var _nwElement0_8 _dafny.Array = (_252_v0)
		_ = _nwElement0_8
		var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(6))
		_ = _nw55
		(_nw55).ArraySet1(_nwElement0_8, 0)
		(_nw55).ArraySet1(_252_v0, 1)
		(_nw55).ArraySet1(_252_v0, 2)
		(_nw55).ArraySet1(_252_v0, 3)
		(_nw55).ArraySet1(_252_v0, 4)
		(_nw55).ArraySet1(_252_v0, 5)
		_298_v32 = _nw55
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_298_v32), 0))
		_ = _index46
		(_298_v32).ArraySet1(_252_v0, (_index46).Int())
		var _299_v33 _dafny.Array
		_ = _299_v33
		var _nwElement0_9 _dafny.Array = _256_v3
		_ = _nwElement0_9
		var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(9))
		_ = _nw56
		(_nw56).ArraySet1(_nwElement0_9, 0)
		(_nw56).ArraySet1(_256_v3, 1)
		(_nw56).ArraySet1(_256_v3, 2)
		(_nw56).ArraySet1(_256_v3, 3)
		(_nw56).ArraySet1(_256_v3, 4)
		(_nw56).ArraySet1(_256_v3, 5)
		(_nw56).ArraySet1(_256_v3, 6)
		(_nw56).ArraySet1(_256_v3, 7)
		(_nw56).ArraySet1(_256_v3, 8)
		_299_v33 = _nw56
		var _300_v34 _dafny.Sequence
		_ = _300_v34
		_300_v34 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_256_v3), 0))
		_ = _index47
		(_256_v3).ArraySet1(_dafny.IntOfUint32((_300_v34).Cardinality()), (_index47).Int())
		var _301_v35 D0
		_ = _301_v35
		_301_v35 = Companion_D0_.Create_DC1_(!(_254_v1))
		var _302_v36 _dafny.Sequence
		_ = _302_v36
		_302_v36 = _dafny.SeqOf(_301_v35)
		var _303_v37 _dafny.Map
		_ = _303_v37
		_303_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (Companion_D2_.Create_DC3_(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer28 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_304_v35 D0) func(_dafny.Int) D0 {
			return func(_305_i8 _dafny.Int) D0 {
				return _304_v35
			}
		})(_301_v35))), (Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer29 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}((func(_306_v35 D0) func(_dafny.Int) D0 {
			return func(_307_i8 _dafny.Int) D0 {
				return _306_v35
			}
		})(_301_v35)))).Cardinality()))).Uint32(), _301_v35))).Dtor_cf3())
		var _308_v38 _dafny.Sequence
		_ = _308_v38
		_308_v38 = _dafny.SeqOf(_302_v36, (func() _dafny.Sequence {
			if (_303_v37).Contains(_281_v22) {
				return (_303_v37).Get(_281_v22).(_dafny.Sequence)
			}
			return _302_v36
		})(), _302_v36, _dafny.SeqOf(_301_v35, _301_v35))
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_298_v32), 0))
		_ = _index48
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_256_v3), 0))
		_ = _index49
		var _rhs46 _dafny.Array = _252_v0
		_ = _rhs46
		var _rhs47 _dafny.Array = _299_v33
		_ = _rhs47
		var _rhs48 bool = _dafny.Companion_Sequence_.IsPrefixOf(_302_v36, (_308_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.IntOfUint32((_308_v38).Cardinality()))).Uint32()).(_dafny.Sequence))
		_ = _rhs48
		var _rhs49 _dafny.Int = (_255_v2).Minus((_dafny.SetOf(_281_v22)).Cardinality())
		_ = _rhs49
		var _lhs12 _dafny.Array = _298_v32
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_298_v32), 0))
		_ = _lhs13
		var _lhs14 _dafny.Array = _256_v3
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_256_v3), 0))
		_ = _lhs15
		(_lhs12).ArraySet1(_rhs46, (_lhs13).Int())
		_299_v33 = _rhs47
		_281_v22 = _rhs48
		(_lhs14).ArraySet1(_rhs49, (_lhs15).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _index50
		(_252_v0).ArraySet1(!((_280_v21).Select((Companion_Default___.SafeIndex((_256_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_256_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_280_v21).Cardinality()))).Uint32()).(bool)), (_index50).Int())
	}
	if false {
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _index51
		(_252_v0).ArraySet1((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (_index51).Int())
		var _309_v39 bool
		_ = _309_v39
		var _out4 bool
		_ = _out4
		_out4 = Companion_Default___.M0(!((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)), _255_v2, _251_globalState)
		_309_v39 = _out4
		_252_v0 = _252_v0
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _index52
		var _rhs50 bool = (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
		_ = _rhs50
		var _rhs51 bool = _309_v39
		_ = _rhs51
		var _rhs52 _dafny.Int = _255_v2
		_ = _rhs52
		var _lhs16 _dafny.Array = _252_v0
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _lhs17
		_309_v39 = _rhs50
		(_lhs16).ArraySet1(_rhs51, (_lhs17).Int())
		_255_v2 = _rhs52
		var _310_v40 _dafny.Sequence
		_ = _310_v40
		_310_v40 = _dafny.UnicodeSeqOfUtf8Bytes("qn")
		var _311_v41 _dafny.Set
		_ = _311_v41
		_311_v41 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fylsrkxqe")).Cardinality()), _255_v2, _255_v2)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _index53
		(_252_v0).ArraySet1((func() bool {
			if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
				return (_dafny.IntOfUint32((_310_v40).Cardinality())).Cmp(_255_v2) == 0
			}
			return (_311_v41).IsProperSubsetOf(_311_v41)
		})(), (_index53).Int())
	} else {
		var _312_v42 _dafny.MultiSet
		_ = _312_v42
		_312_v42 = _dafny.MultiSetOf(_254_v1, !(_254_v1), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
		var _313_v43 _dafny.Sequence
		_ = _313_v43
		_313_v43 = _dafny.SeqOf(_254_v1)
		_312_v42 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if _254_v1 {
				return _313_v43
			}
			return _dafny.SeqOf(_254_v1, !(_254_v1))
		})(), _313_v43))
		_255_v2 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_255_v2))
		var _314_v44 _dafny.Map
		_ = _314_v44
		_314_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
		var _315_v45 bool
		_ = _315_v45
		var _out5 bool
		_ = _out5
		_out5 = Companion_Default___.M0(!((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _254_v1)).Equals((_314_v44).Update(_254_v1, Companion_Default___.Fm2((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _255_v2, (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _254_v1, _251_globalState)))), _255_v2, _251_globalState)
		_315_v45 = _out5
		if !(_254_v1) {
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_256_v3), 0))
			_ = _index54
			(_256_v3).ArraySet1((_dafny.SetOf(_315_v45, (_313_v43).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.IntOfUint32((_313_v43).Cardinality()))).Uint32()).(bool))).Cardinality(), (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_256_v3), 0))
			_ = _index55
			(_256_v3).ArraySet1(Companion_Default___.SafeModuloInt(_255_v2, _dafny.IntOfInt64(501)), (_index55).Int())
			_255_v2 = _255_v2
			_254_v1 = _254_v1
			_314_v44 = (_314_v44).Update(_315_v45, _315_v45)
			_315_v45 = false
		} else {
			var _316_v46 bool
			_ = _316_v46
			var _out6 bool
			_ = _out6
			_out6 = Companion_Default___.M0(!(_315_v45), _255_v2, _251_globalState)
			_316_v46 = _out6
			var _317_v47 bool
			_ = _317_v47
			var _out7 bool
			_ = _out7
			_out7 = Companion_Default___.M0(_254_v1, _255_v2, _251_globalState)
			_317_v47 = _out7
			_255_v2 = (_dafny.Zero).Minus((_255_v2).Plus(_255_v2))
			var _318_v48 *C1
			_ = _318_v48
			var _nw57 *C1 = New_C1_()
			_ = _nw57
			_nw57.Ctor__(_255_v2)
			_318_v48 = _nw57
			var _319_v49 _dafny.CodePoint
			_ = _319_v49
			_319_v49 = _dafny.CodePoint('g')
			var _320_v50 _dafny.Sequence
			_ = _320_v50
			_320_v50 = _dafny.SeqOf(_319_v49)
			var _321_v51 _dafny.Map
			_ = _321_v51
			_321_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v50, _315_v45)
			_321_v51 = (_321_v51).Update(_dafny.SeqOf(_319_v49), _dafny.Companion_Sequence_.IsProperPrefixOf(_313_v43, _313_v43))
		}
		var _322_v52 _dafny.CodePoint
		_ = _322_v52
		_322_v52 = _dafny.CodePoint('y')
		var _323_v53 *C5
		_ = _323_v53
		var _nw58 *C5 = New_C5_()
		_ = _nw58
		_nw58.Ctor__(_322_v52, _255_v2, _255_v2)
		_323_v53 = _nw58
	}
	var _324_v54 *C3
	_ = _324_v54
	var _nw59 *C3 = New_C3_()
	_ = _nw59
	_nw59.Ctor__(_255_v2)
	_324_v54 = _nw59
	var _325_v55 D16
	_ = _325_v55
	_325_v55 = Companion_D16_.Create_DC33_(_324_v54)
	var _326_v56 _dafny.Array
	_ = _326_v56
	var _nwElement0_10 *C3 = _324_v54
	_ = _nwElement0_10
	var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(5))
	_ = _nw60
	(_nw60).ArraySet1(_nwElement0_10, 0)
	(_nw60).ArraySet1(_324_v54, 1)
	(_nw60).ArraySet1(_324_v54, 2)
	(_nw60).ArraySet1(_324_v54, 3)
	(_nw60).ArraySet1((_325_v55).Dtor_cf54(), 4)
	_326_v56 = _nw60
	var _327_v57 _dafny.Sequence
	_ = _327_v57
	_327_v57 = _dafny.SeqOf(_326_v56)
	var _328_v58 _dafny.CodePoint
	_ = _328_v58
	_328_v58 = _dafny.CodePoint('x')
	var _329_v59 _dafny.Map
	_ = _329_v59
	_329_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_327_v57).Select((Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_327_v57).Cardinality()))).Uint32()).(_dafny.Array), _328_v58)
	_329_v59 = (_329_v59).Merge(_329_v59)
	_255_v2 = _255_v2
	_328_v58 = _328_v58
	var _330_v60 D2
	_ = _330_v60
	_330_v60 = Companion_D2_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("ontdavgjc"), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
	var _pat_let_tv2 = _330_v60
	_ = _pat_let_tv2
	var _pat_let_tv3 = _330_v60
	_ = _pat_let_tv3
	var _pat_let_tv4 = _330_v60
	_ = _pat_let_tv4
	var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
	_ = _index56
	var _rhs53 _dafny.Int = ((_255_v2).Plus(_255_v2)).Times(_dafny.IntOfInt64(-502))
	_ = _rhs53
	var _rhs54 bool = (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
	_ = _rhs54
	var _rhs55 D2 = func(_source1 D5) D2 {
		if _source1.Is_DC12() {
			var _331___mcc_h0 _dafny.Int = _source1.Get_().(D5_DC12).Cf15
			_ = _331___mcc_h0
			var _332___mcc_h1 _dafny.Int = _source1.Get_().(D5_DC12).Cf16
			_ = _332___mcc_h1
			var _333_cf16 _dafny.Int = _332___mcc_h1
			_ = _333_cf16
			var _334_cf15 _dafny.Int = _331___mcc_h0
			_ = _334_cf15
			return _pat_let_tv2
		} else if _source1.Is_DC13() {
			var _335___mcc_h2 _dafny.Int = _source1.Get_().(D5_DC13).Cf17
			_ = _335___mcc_h2
			var _336___mcc_h3 _dafny.Set = _source1.Get_().(D5_DC13).Cf18
			_ = _336___mcc_h3
			var _337___mcc_h4 _dafny.CodePoint = _source1.Get_().(D5_DC13).Cf19
			_ = _337___mcc_h4
			var _338___mcc_h5 _dafny.Int = _source1.Get_().(D5_DC13).Cf20
			_ = _338___mcc_h5
			var _339___mcc_h6 bool = _source1.Get_().(D5_DC13).Cf21
			_ = _339___mcc_h6
			var _340_cf21 bool = _339___mcc_h6
			_ = _340_cf21
			var _341_cf20 _dafny.Int = _338___mcc_h5
			_ = _341_cf20
			var _342_cf19 _dafny.CodePoint = _337___mcc_h4
			_ = _342_cf19
			var _343_cf18 _dafny.Set = _336___mcc_h3
			_ = _343_cf18
			var _344_cf17 _dafny.Int = _335___mcc_h2
			_ = _344_cf17
			return _pat_let_tv3
		} else {
			var _345___mcc_h7 _dafny.Set = _source1.Get_().(D5_DC11).Cf14
			_ = _345___mcc_h7
			var _346_cf14 _dafny.Set = _345___mcc_h7
			_ = _346_cf14
			return _pat_let_tv4
		}
	}(Companion_D5_.Create_DC12_(_255_v2, _255_v2))
	_ = _rhs55
	var _lhs18 _dafny.Array = _252_v0
	_ = _lhs18
	var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
	_ = _lhs19
	_255_v2 = _rhs53
	(_lhs18).ArraySet1(_rhs54, (_lhs19).Int())
	_330_v60 = _rhs55
	var _347_v61 _dafny.Set
	_ = _347_v61
	_347_v61 = _dafny.SetOf((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
	var _348_i9 _dafny.Int
	_ = _348_i9
	_348_i9 = _dafny.Zero
	{
		for (_347_v61).IsProperSubsetOf(_dafny.SetOf(_254_v1, _254_v1, (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))) {
			{
				if (_348_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_348_i9 = (_348_i9).Plus(_dafny.One)
				var _349_v62 D3
				_ = _349_v62
				_349_v62 = Companion_D3_.Create_DC8_((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _255_v2)
				var _350_v63 bool
				_ = _350_v63
				var _out8 bool
				_ = _out8
				_out8 = Companion_Default___.M0(((_dafny.Zero).Minus(Companion_Default___.Fm3(_dafny.SetOf(_254_v1, _254_v1), _328_v58, _255_v2, _255_v2, _251_globalState))).Cmp((_349_v62).Dtor_cf9()) < 0, _255_v2, _251_globalState)
				_350_v63 = _out8
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_256_v3), 0))
				_ = _index57
				(_256_v3).ArraySet1(Companion_Default___.SafeDivisionInt(_255_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(542))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_351_v0 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_352_i10 _dafny.Int) _dafny.CodePoint {
						return (Companion_D13_.Create_DC29_(_dafny.CodePoint('j'), (_351_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_351_v0), 0))).Int()).(bool))).Dtor_cf49()
					}
				})(_252_v0)))).Cardinality())), (_index57).Int())
				var _353_v64 *C5
				_ = _353_v64
				var _nw61 *C5 = New_C5_()
				_ = _nw61
				_nw61.Ctor__(_328_v58, _255_v2, _255_v2)
				_353_v64 = _nw61
				var _354_v65 _dafny.Map
				_ = _354_v65
				_354_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_353_v64, _255_v2)
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_256_v3), 0))
				_ = _index58
				(_256_v3).ArraySet1((func() _dafny.Int {
					if (_354_v65).Contains(_353_v64) {
						return (_354_v65).Get(_353_v64).(_dafny.Int)
					}
					return _255_v2
				})(), (_index58).Int())
				var _355_v66 _dafny.Array
				_ = _355_v66
				var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw62
				_355_v66 = _nw62
				_355_v66 = _355_v66
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_256_v3), 0))
				_ = _index59
				(_256_v3).ArraySet1((_353_v64).F2(), (_index59).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_255_v2 = Companion_Default___.SafeDivisionInt(_255_v2, _255_v2)
	var _356_i11 _dafny.Int
	_ = _356_i11
	_356_i11 = _dafny.Zero
	{
		for !(_254_v1) {
			{
				if (_356_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_356_i11 = (_356_i11).Plus(_dafny.One)
				_254_v1 = _254_v1
				_255_v2 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_255_v2, _dafny.IntOfInt64(770)), _255_v2)
				if (_255_v2).Cmp(_255_v2) == 0 {
					var _357_v68 T0
					_ = _357_v68
					var _nw63 *C3 = New_C3_()
					_ = _nw63
					_nw63.Ctor__((func() _dafny.Map {
						var _coll14 = _dafny.NewMapBuilder()
						_ = _coll14
						for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v2, _254_v1)).Keys().Elements()); ; {
							_compr_14, _ok14 := _iter14()
							if !_ok14 {
								break
							}
							var _358_v67 _dafny.Int
							_358_v67 = interface{}(_compr_14).(_dafny.Int)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v2, _254_v1)).Contains(_358_v67) {
								_coll14.Add(Companion_Default___.SafeDivisionInt(_358_v67, (_347_v61).Cardinality()), _255_v2)
							}
						}
						return _coll14.ToMap()
					}()).Cardinality())
					_357_v68 = _nw63
					var _359_v69 D13
					_ = _359_v69
					_359_v69 = Companion_D13_.Create_DC27_(_328_v58, _357_v68, _328_v58, _255_v2)
					_255_v2 = (_359_v69).Dtor_cf46()
					var _360_v70 _dafny.Set
					_ = _360_v70
					_360_v70 = _dafny.SetOf(_252_v0)
					_255_v2 = Companion_Default___.SafeDivisionInt((_360_v70).Cardinality(), (_dafny.Zero).Minus((_359_v69).Dtor_cf46()))
					var _361_v71 _dafny.Sequence
					_ = _361_v71
					_361_v71 = _dafny.UnicodeSeqOfUtf8Bytes("nmqirv")
					(_324_v54).M7(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("oaqldlvsi"), _361_v71), _254_v1, _256_v3, _dafny.IntOfUint32((_361_v71).Cardinality()), _251_globalState)
					_255_v2 = (_357_v68).F0()
					var _362_v72 _dafny.Sequence
					_ = _362_v72
					_362_v72 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg31 _dafny.Int) interface{} {
							return coer31(arg31)
						}
					}((func(_363_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_364_i12 _dafny.Int) _dafny.CodePoint {
							return _363_v58
						}
					})(_328_v58))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(259))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}((func(_365_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_366_i13 _dafny.Int) _dafny.CodePoint {
							return _365_v58
						}
					})(_328_v58))))
					var _367_v73 _dafny.Sequence
					_ = _367_v73
					_367_v73 = _dafny.SeqOf((_255_v2).Cmp(_dafny.IntOfInt64(-839)) <= 0, _dafny.Companion_Sequence_.IsPrefixOf(_361_v71, (_362_v72).Select((Companion_Default___.SafeIndex((_357_v68).F0(), _dafny.IntOfUint32((_362_v72).Cardinality()))).Uint32()).(_dafny.Sequence)), _254_v1)
					var _368_v74 _dafny.Map
					_ = _368_v74
					_368_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v58, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qmslbj")).Cardinality()))
					var _369_v75 _dafny.Set
					_ = _369_v75
					_369_v75 = _dafny.SetOf((func() _dafny.Int {
						if (_368_v74).Contains(_328_v58) {
							return (_368_v74).Get(_328_v58).(_dafny.Int)
						}
						return _255_v2
					})(), (_357_v68).F0())
					var _370_v76 _dafny.MultiSet
					_ = _370_v76
					_370_v76 = _dafny.MultiSetOf(_328_v58, _328_v58)
					var _371_v77 _dafny.Map
					_ = _371_v77
					_371_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_357_v68).F0(), _dafny.MultiSetOf(_328_v58))
					var _372_v78 _dafny.Sequence
					_ = _372_v78
					_372_v78 = _dafny.SeqOf(_370_v76, (func() _dafny.MultiSet {
						if (_371_v77).Contains(_255_v2) {
							return (_371_v77).Get(_255_v2).(_dafny.MultiSet)
						}
						return _370_v76
					})())
					var _rhs56 _dafny.Int = Companion_Default___.Fm3(_347_v61, (func() _dafny.CodePoint {
						if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
							return _328_v58
						}
						return _328_v58
					})(), (_255_v2).Times((_357_v68).F0()), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_361_v71).Cardinality()), (_369_v75).Cardinality()), _251_globalState)
					_ = _rhs56
					var _rhs57 bool = (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)
					_ = _rhs57
					var _rhs58 _dafny.Sequence = _367_v73
					_ = _rhs58
					var _rhs59 _dafny.Set = Companion_Default___.Fm10(((_372_v78).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.IntOfUint32((_372_v78).Cardinality()))).Uint32()).(_dafny.MultiSet)).Intersection(_370_v76), (_324_v54).Fm13(_255_v2, (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _251_globalState), _251_globalState)
					_ = _rhs59
					var _rhs60 _dafny.CodePoint = _328_v58
					_ = _rhs60
					_255_v2 = _rhs56
					_254_v1 = _rhs57
					_367_v73 = _rhs58
					_347_v61 = _rhs59
					_328_v58 = _rhs60
				} else {
					_255_v2 = _255_v2
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
					_ = _index60
					(_252_v0).ArraySet1((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (_index60).Int())
					var _373_v79 _dafny.Map
					_ = _373_v79
					_373_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v2, false)
					var _374_v80 _dafny.MultiSet
					_ = _374_v80
					_374_v80 = _dafny.MultiSetOf(_328_v58, _328_v58)
					var _375_v81 _dafny.Map
					_ = _375_v81
					_375_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v80, _256_v3)
					_373_v79 = (_373_v79).Update(_255_v2, ((_375_v81).Update(_374_v80, _256_v3)).Contains(_374_v80))
					var _376_v82 _dafny.Map
					_ = _376_v82
					_376_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_255_v2, _255_v2), (_dafny.SetOf((_324_v54).Fm13(_dafny.IntOfInt64(-924), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _251_globalState))).Union(_347_v61))
					_376_v82 = (_376_v82).Update(_255_v2, _347_v61)
					var _377_v83 D6
					_ = _377_v83
					_377_v83 = Companion_D6_.Create_DC14_(_256_v3)
					var _378_v84 _dafny.Sequence
					_ = _378_v84
					_378_v84 = _dafny.SeqOf(_255_v2)
					var _379_v85 _dafny.Sequence
					_ = _379_v85
					_379_v85 = _dafny.UnicodeSeqOfUtf8Bytes("ryvnthp")
					var _380_v86 *C4
					_ = _380_v86
					var _nw64 *C4 = New_C4_()
					_ = _nw64
					_nw64.Ctor__(_255_v2)
					_380_v86 = _nw64
					var _381_v87 _dafny.Sequence
					_ = _381_v87
					_381_v87 = _dafny.SeqOf(_380_v86)
					var _382_v88 _dafny.Sequence
					_ = _382_v88
					_382_v88 = _dafny.SeqOf(_380_v86, (_381_v87).Select((Companion_Default___.SafeIndex(_255_v2, _dafny.IntOfUint32((_381_v87).Cardinality()))).Uint32()).(*C4))
					(_324_v54).M7(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(627))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg33 _dafny.Int) interface{} {
							return coer33(arg33)
						}
					}(func(_383_i14 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('q')
					})), _254_v1, (func() _dafny.Array {
						if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
							return _256_v3
						}
						return (_377_v83).Dtor_cf22()
					})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_378_v84, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.IntOfUint32((_378_v84).Cardinality()))).Uint32(), _255_v2), _dafny.SeqOf(_dafny.IntOfUint32((_379_v85).Cardinality()), _dafny.IntOfUint32((_382_v88).Cardinality()), _255_v2))).Cardinality())), _251_globalState)
				}
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
				_ = _index61
				(_252_v0).ArraySet1(_254_v1, (_index61).Int())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _384_i15 _dafny.Int
	_ = _384_i15
	_384_i15 = _dafny.Zero
	{
		for (_254_v1) || ((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool)) {
			{
				if (_384_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_384_i15 = (_384_i15).Plus(_dafny.One)
				var _385_v89 T1
				_ = _385_v89
				var _nw65 *C5 = New_C5_()
				_ = _nw65
				_nw65.Ctor__(_328_v58, _255_v2, _255_v2)
				_385_v89 = _nw65
				var _386_v90 _dafny.Array
				_ = _386_v90
				var _nw66 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
				_ = _nw66
				_386_v90 = _nw66
				var _387_v91 _dafny.Map
				_ = _387_v91
				_387_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_385_v89, _386_v90)
				var _388_v92 D3
				_ = _388_v92
				_388_v92 = Companion_D3_.Create_DC9_(_255_v2, (func() _dafny.Array {
					if (_387_v91).Contains(_385_v89) {
						return (_387_v91).Get(_385_v89).(_dafny.Array)
					}
					return _386_v90
				})(), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
				var _389_v93 _dafny.Set
				_ = _389_v93
				_389_v93 = _dafny.SetOf((_388_v92).Dtor_cf10(), (_385_v89).F0(), _dafny.IntOfInt64(20))
				var _390_v94 D5
				_ = _390_v94
				_390_v94 = Companion_D5_.Create_DC13_((_255_v2).Times(Companion_Default___.Fm3(_dafny.SetOf(_254_v1), _328_v58, _255_v2, _255_v2, _251_globalState)), _347_v61, _328_v58, Companion_Default___.SafeModuloInt(_255_v2, _255_v2), (_389_v93).IsDisjointFrom(_389_v93))
				_390_v94 = _390_v94
				var _391_v95 _dafny.Map
				_ = _391_v95
				_391_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_385_v89).F0(), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool))
				_391_v95 = (_391_v95).Update((_385_v89).F0(), _254_v1)
				_256_v3 = _256_v3
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_256_v3), 0))
				_ = _index62
				(_256_v3).ArraySet1((_dafny.Zero).Minus((_385_v89).F0()), (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_256_v3), 0))
				_ = _index63
				(_256_v3).ArraySet1(_255_v2, (_index63).Int())
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _392_v96 *C3
	_ = _392_v96
	var _nw67 *C3 = New_C3_()
	_ = _nw67
	_nw67.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(747), _dafny.IntOfInt64(988)))
	_392_v96 = _nw67
	var _393_v97 _dafny.Sequence
	_ = _393_v97
	_393_v97 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(811))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg34 _dafny.Int) interface{} {
			return coer34(arg34)
		}
	}(func(_394_i16 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("nsgqps")
	}))
	var _395_v98 _dafny.Map
	_ = _395_v98
	_395_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _393_v97)
	_395_v98 = (_395_v98).Update((_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), _393_v97)
	var _396_v99 _dafny.Sequence
	_ = _396_v99
	_396_v99 = _dafny.SeqOf(_328_v58)
	(_324_v54).M7(_dafny.UnicodeSeqOfUtf8Bytes("nujtypxb"), (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool), (func() _dafny.Array {
		if true {
			return _256_v3
		}
		return _256_v3
	})(), _dafny.IntOfUint32((_396_v99).Cardinality()), _251_globalState)
	var _397_v100 _dafny.Array
	_ = _397_v100
	var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
	_ = _nw68
	_397_v100 = _nw68
	var _398_v101 _dafny.Int
	_ = _398_v101
	var _399_v102 _dafny.Int
	_ = _399_v102
	var _400_v103 _dafny.Int
	_ = _400_v103
	var _out9 _dafny.Int
	_ = _out9
	var _out10 _dafny.Int
	_ = _out10
	var _out11 _dafny.Int
	_ = _out11
	_out9, _out10, _out11 = (_324_v54).M1(_397_v100, (func() _dafny.Int {
		if _254_v1 {
			return (_dafny.Zero).Minus(_255_v2)
		}
		return (_dafny.Zero).Minus(_255_v2)
	})(), _251_globalState)
	_398_v101 = _out9
	_399_v102 = _out10
	_400_v103 = _out11
	var _401_v104 _dafny.Sequence
	_ = _401_v104
	_401_v104 = _dafny.SeqOf(_255_v2)
	var _hi1 _dafny.Int = (func() _dafny.Int {
		if (_252_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))).Int()).(bool) {
			return _255_v2
		}
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-266))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}((func(_402_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_403_i18 _dafny.Int) _dafny.CodePoint {
				return _402_v58
			}
		})(_328_v58)))).Cardinality())
	})()
	_ = _hi1
	for _404_i17 := (_401_v104).Select((Companion_Default___.SafeIndex(_398_v101, _dafny.IntOfUint32((_401_v104).Cardinality()))).Uint32()).(_dafny.Int); _404_i17.Cmp(_hi1) < 0; _404_i17 = _404_i17.Plus(_dafny.One) {
		var _405_v105 _dafny.Map
		_ = _405_v105
		_405_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_260_v5).IsProperSubsetOf(_dafny.MultiSetOf(_255_v2)), _252_v0)
		_405_v105 = (_405_v105).Update(!_dafny.Companion_Sequence_.Equal((_330_v60).Dtor_cf4(), _396_v99), (func() _dafny.Array {
			if _254_v1 {
				return _252_v0
			}
			return _252_v0
		})())
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_252_v0), 0))
		_ = _index64
		(_252_v0).ArraySet1(_254_v1, (_index64).Int())
		var _406_v106 _dafny.Map
		_ = _406_v106
		_406_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v2, _255_v2)
		var _rhs61 _dafny.Int = (func() _dafny.Int {
			if (_406_v106).Contains(Companion_Default___.Fm3(_347_v61, _328_v58, _404_i17, _404_i17, _251_globalState)) {
				return (_406_v106).Get(Companion_Default___.Fm3(_347_v61, _328_v58, _404_i17, _404_i17, _251_globalState)).(_dafny.Int)
			}
			return _399_v102
		})()
		_ = _rhs61
		var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_396_v99, _396_v99)
		_ = _rhs62
		_255_v2 = _rhs61
		_396_v99 = _rhs62
		var _407_v107 _dafny.Array
		_ = _407_v107
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_8
		var _nw69 _dafny.Array
		_ = _nw69
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw69 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Set = (func(_408_v104 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
				return func(_409_i19 _dafny.Int) _dafny.Set {
					return (_dafny.SetOf(_dafny.IntOfUint32((_408_v104).Cardinality())))
				}
			})(_401_v104)
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
		_407_v107 = _nw69
		_407_v107 = _407_v107
	}
	_dafny.Print((_252_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_254_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_255_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v3).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v4).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v5).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-806))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_325_v55).Dtor_cf54()).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_326_v56).ArrayGet1((_dafny.Zero).Int()).(*C3)).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_326_v56).ArrayGet1((_dafny.One).Int()).(*C3)).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_326_v56).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C3)).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_326_v56).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C3)).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_326_v56).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_327_v57).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_328_v58)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_329_v59).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v60).Dtor_cf4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v60).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_347_v61).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_348_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_356_i11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_384_i15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_393_v97), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_395_v98).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps"), _dafny.UnicodeSeqOfUtf8Bytes("nsgqps")))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_396_v99, _dafny.SeqOf(_dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_398_v101)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_399_v102)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_400_v103)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_401_v104, _dafny.SeqOf(_dafny.One)))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool) D0 {
	return D0{D0_DC1{Cf1}}
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
	return Companion_D0_.Create_DC1_(false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf1 == data2.Cf1
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
	Cf2 _dafny.Array
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Array) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D1) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf2 == data2.Cf2
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
	Cf4 _dafny.Sequence
	Cf5 bool
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Sequence, Cf5 bool) D2 {
	return D2{D2_DC4{Cf4, Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_() D2 {
	return D2{D2_DC5{}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC3 struct {
	Cf3 _dafny.Sequence
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf3 _dafny.Sequence) D2 {
	return D2{D2_DC3{Cf3}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC6 struct {
	Cf6 D2
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 D2) D2 {
	return D2{D2_DC6{Cf6}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.EmptySeq, false)
}

func (_this D2) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) Dtor_cf5() bool {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D2_DC3).Cf3
}

func (_this D2) Dtor_cf6() D2 {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + data.Cf4.VerbatimString(true) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf4.Equals(data2.Cf4) && data1.Cf5 == data2.Cf5
		}
	case D2_DC5:
		{
			_, ok := other.Get_().(D2_DC5)
			return ok
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
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

type D3_DC8 struct {
	Cf8 bool
	Cf9 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf8 bool, Cf9 _dafny.Int) D3 {
	return D3{D3_DC8{Cf8, Cf9}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf10 _dafny.Int
	Cf11 _dafny.Array
	Cf12 bool
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf10 _dafny.Int, Cf11 _dafny.Array, Cf12 bool) D3 {
	return D3{D3_DC9{Cf10, Cf11, Cf12}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC7 struct {
	Cf7 _dafny.Sequence
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf7 _dafny.Sequence) D3 {
	return D3{D3_DC7{Cf7}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(false, _dafny.Zero)
}

func (_this D3) Dtor_cf8() bool {
	return _this.Get_().(D3_DC8).Cf8
}

func (_this D3) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf9
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) Dtor_cf11() _dafny.Array {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC9).Cf12
}

func (_this D3) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D3_DC7).Cf7
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf13 _dafny.CodePoint
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf13 _dafny.CodePoint) D4 {
	return D4{D4_DC10{Cf13}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D4) Dtor_cf13() _dafny.CodePoint {
	return _this.Get_().(D4_DC10).Cf13
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf13 == data2.Cf13
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
	Cf15 _dafny.Int
	Cf16 _dafny.Int
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf15 _dafny.Int, Cf16 _dafny.Int) D5 {
	return D5{D5_DC12{Cf15, Cf16}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC13 struct {
	Cf17 _dafny.Int
	Cf18 _dafny.Set
	Cf19 _dafny.CodePoint
	Cf20 _dafny.Int
	Cf21 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf17 _dafny.Int, Cf18 _dafny.Set, Cf19 _dafny.CodePoint, Cf20 _dafny.Int, Cf21 bool) D5 {
	return D5{D5_DC13{Cf17, Cf18, Cf19, Cf20, Cf21}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC11 struct {
	Cf14 _dafny.Set
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf14 _dafny.Set) D5 {
	return D5{D5_DC11{Cf14}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC12_(_dafny.Zero, _dafny.Zero)
}

func (_this D5) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf15
}

func (_this D5) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf16
}

func (_this D5) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf17
}

func (_this D5) Dtor_cf18() _dafny.Set {
	return _this.Get_().(D5_DC13).Cf18
}

func (_this D5) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf14() _dafny.Set {
	return _this.Get_().(D5_DC11).Cf14
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Equals(data2.Cf18) && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D6_DC15 struct {
	Cf23 _dafny.Int
	Cf24 bool
	Cf25 _dafny.Int
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf23 _dafny.Int, Cf24 bool, Cf25 _dafny.Int) D6 {
	return D6{D6_DC15{Cf23, Cf24, Cf25}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC14 struct {
	Cf22 _dafny.Array
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf22 _dafny.Array) D6 {
	return D6{D6_DC14{Cf22}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC15_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC15).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf25
}

func (_this D6) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D6_DC14).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf22 == data2.Cf22
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
	Cf27 bool
	Cf28 bool
	Cf29 bool
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf27 bool, Cf28 bool, Cf29 bool) D7 {
	return D7{D7_DC17{Cf27, Cf28, Cf29}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC18 struct {
	Cf30 _dafny.Int
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf30 _dafny.Int) D7 {
	return D7{D7_DC18{Cf30}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC16 struct {
	Cf26 _dafny.Sequence
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf26 _dafny.Sequence) D7 {
	return D7{D7_DC16{Cf26}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC17_(false, false, false)
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC17).Cf27
}

func (_this D7) Dtor_cf28() bool {
	return _this.Get_().(D7_DC17).Cf28
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC17).Cf29
}

func (_this D7) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D7_DC18).Cf30
}

func (_this D7) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf27 == data2.Cf27 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
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

type D8_DC19 struct {
	Cf31 _dafny.Array
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf31 _dafny.Array) D8 {
	return D8{D8_DC19{Cf31}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D8) Dtor_cf31() _dafny.Array {
	return _this.Get_().(D8_DC19).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf31 == data2.Cf31
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
	Cf32 _dafny.Map
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf32 _dafny.Map) D9 {
	return D9{D9_DC20{Cf32}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D9) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D9_DC20).Cf32
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	Cf34 _dafny.Int
	Cf35 bool
	Cf36 bool
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf34 _dafny.Int, Cf35 bool, Cf36 bool) D10 {
	return D10{D10_DC22{Cf34, Cf35, Cf36}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC21 struct {
	Cf33 T0
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf33 T0) D10 {
	return D10{D10_DC21{Cf33}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_(_dafny.Zero, false, false)
}

func (_this D10) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D10_DC22).Cf34
}

func (_this D10) Dtor_cf35() bool {
	return _this.Get_().(D10_DC22).Cf35
}

func (_this D10) Dtor_cf36() bool {
	return _this.Get_().(D10_DC22).Cf36
}

func (_this D10) Dtor_cf33() T0 {
	return _this.Get_().(D10_DC21).Cf33
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && _dafny.AreEqual(data1.Cf33, data2.Cf33)
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
	Cf38 _dafny.Array
	Cf39 bool
	Cf40 bool
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf38 _dafny.Array, Cf39 bool, Cf40 bool) D11 {
	return D11{D11_DC24{Cf38, Cf39, Cf40}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC23 struct {
	Cf37 _dafny.Sequence
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf37 _dafny.Sequence) D11 {
	return D11{D11_DC23{Cf37}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC24_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, false)
}

func (_this D11) Dtor_cf38() _dafny.Array {
	return _this.Get_().(D11_DC24).Cf38
}

func (_this D11) Dtor_cf39() bool {
	return _this.Get_().(D11_DC24).Cf39
}

func (_this D11) Dtor_cf40() bool {
	return _this.Get_().(D11_DC24).Cf40
}

func (_this D11) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D11_DC23).Cf37
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf37) + ")"
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
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40
		}
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf37.Equals(data2.Cf37)
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
	Cf41 _dafny.Array
}

func (D12_DC25) isD12() {}

func (CompanionStruct_D12_) Create_DC25_(Cf41 _dafny.Array) D12 {
	return D12{D12_DC25{Cf41}}
}

func (_this D12) Is_DC25() bool {
	_, ok := _this.Get_().(D12_DC25)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D12) Dtor_cf41() _dafny.Array {
	return _this.Get_().(D12_DC25).Cf41
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC25:
		{
			return "D12.DC25" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf41 == data2.Cf41
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

type D13_DC27 struct {
	Cf43 _dafny.CodePoint
	Cf44 T0
	Cf45 _dafny.CodePoint
	Cf46 _dafny.Int
}

func (D13_DC27) isD13() {}

func (CompanionStruct_D13_) Create_DC27_(Cf43 _dafny.CodePoint, Cf44 T0, Cf45 _dafny.CodePoint, Cf46 _dafny.Int) D13 {
	return D13{D13_DC27{Cf43, Cf44, Cf45, Cf46}}
}

func (_this D13) Is_DC27() bool {
	_, ok := _this.Get_().(D13_DC27)
	return ok
}

type D13_DC28 struct {
	Cf47 bool
	Cf48 _dafny.CodePoint
}

func (D13_DC28) isD13() {}

func (CompanionStruct_D13_) Create_DC28_(Cf47 bool, Cf48 _dafny.CodePoint) D13 {
	return D13{D13_DC28{Cf47, Cf48}}
}

func (_this D13) Is_DC28() bool {
	_, ok := _this.Get_().(D13_DC28)
	return ok
}

type D13_DC29 struct {
	Cf49 _dafny.CodePoint
	Cf50 bool
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf49 _dafny.CodePoint, Cf50 bool) D13 {
	return D13{D13_DC29{Cf49, Cf50}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

type D13_DC26 struct {
	Cf42 *C0
}

func (D13_DC26) isD13() {}

func (CompanionStruct_D13_) Create_DC26_(Cf42 *C0) D13 {
	return D13{D13_DC26{Cf42}}
}

func (_this D13) Is_DC26() bool {
	_, ok := _this.Get_().(D13_DC26)
	return ok
}

type D13_DC30 struct {
	Cf51 D13
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf51 D13) D13 {
	return D13{D13_DC30{Cf51}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC27_(_dafny.CodePoint('D'), (T0)(nil), _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D13) Dtor_cf43() _dafny.CodePoint {
	return _this.Get_().(D13_DC27).Cf43
}

func (_this D13) Dtor_cf44() T0 {
	return _this.Get_().(D13_DC27).Cf44
}

func (_this D13) Dtor_cf45() _dafny.CodePoint {
	return _this.Get_().(D13_DC27).Cf45
}

func (_this D13) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D13_DC27).Cf46
}

func (_this D13) Dtor_cf47() bool {
	return _this.Get_().(D13_DC28).Cf47
}

func (_this D13) Dtor_cf48() _dafny.CodePoint {
	return _this.Get_().(D13_DC28).Cf48
}

func (_this D13) Dtor_cf49() _dafny.CodePoint {
	return _this.Get_().(D13_DC29).Cf49
}

func (_this D13) Dtor_cf50() bool {
	return _this.Get_().(D13_DC29).Cf50
}

func (_this D13) Dtor_cf42() *C0 {
	return _this.Get_().(D13_DC26).Cf42
}

func (_this D13) Dtor_cf51() D13 {
	return _this.Get_().(D13_DC30).Cf51
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC27:
		{
			return "D13.DC27" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D13_DC28:
		{
			return "D13.DC28" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D13_DC26:
		{
			return "D13.DC26" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC27:
		{
			data2, ok := other.Get_().(D13_DC27)
			return ok && data1.Cf43 == data2.Cf43 && _dafny.AreEqual(data1.Cf44, data2.Cf44) && data1.Cf45 == data2.Cf45 && data1.Cf46.Cmp(data2.Cf46) == 0
		}
	case D13_DC28:
		{
			data2, ok := other.Get_().(D13_DC28)
			return ok && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48
		}
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50
		}
	case D13_DC26:
		{
			data2, ok := other.Get_().(D13_DC26)
			return ok && data1.Cf42 == data2.Cf42
		}
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
			return ok && data1.Cf51.Equals(data2.Cf51)
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
	Cf52 _dafny.MultiSet
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf52 _dafny.MultiSet) D14 {
	return D14{D14_DC31{Cf52}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D14) Dtor_cf52() _dafny.MultiSet {
	return _this.Get_().(D14_DC31).Cf52
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf52) + ")"
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

type D15_DC32 struct {
	Cf53 _dafny.Sequence
}

func (D15_DC32) isD15() {}

func (CompanionStruct_D15_) Create_DC32_(Cf53 _dafny.Sequence) D15 {
	return D15{D15_DC32{Cf53}}
}

func (_this D15) Is_DC32() bool {
	_, ok := _this.Get_().(D15_DC32)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D15) Dtor_cf53() _dafny.Sequence {
	return _this.Get_().(D15_DC32).Cf53
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC32:
		{
			return "D15.DC32" + "(" + _dafny.String(data.Cf53) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC32:
		{
			data2, ok := other.Get_().(D15_DC32)
			return ok && data1.Cf53.Equals(data2.Cf53)
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

type D16_DC34 struct {
}

func (D16_DC34) isD16() {}

func (CompanionStruct_D16_) Create_DC34_() D16 {
	return D16{D16_DC34{}}
}

func (_this D16) Is_DC34() bool {
	_, ok := _this.Get_().(D16_DC34)
	return ok
}

type D16_DC35 struct {
	Cf55 _dafny.Int
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf55 _dafny.Int) D16 {
	return D16{D16_DC35{Cf55}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

type D16_DC33 struct {
	Cf54 *C3
}

func (D16_DC33) isD16() {}

func (CompanionStruct_D16_) Create_DC33_(Cf54 *C3) D16 {
	return D16{D16_DC33{Cf54}}
}

func (_this D16) Is_DC33() bool {
	_, ok := _this.Get_().(D16_DC33)
	return ok
}

type D16_DC36 struct {
	Cf56 D16
}

func (D16_DC36) isD16() {}

func (CompanionStruct_D16_) Create_DC36_(Cf56 D16) D16 {
	return D16{D16_DC36{Cf56}}
}

func (_this D16) Is_DC36() bool {
	_, ok := _this.Get_().(D16_DC36)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC34_()
}

func (_this D16) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf55
}

func (_this D16) Dtor_cf54() *C3 {
	return _this.Get_().(D16_DC33).Cf54
}

func (_this D16) Dtor_cf56() D16 {
	return _this.Get_().(D16_DC36).Cf56
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC34:
		{
			return "D16.DC34"
		}
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D16_DC33:
		{
			return "D16.DC33" + "(" + _dafny.String(data.Cf54) + ")"
		}
	case D16_DC36:
		{
			return "D16.DC36" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC34:
		{
			_, ok := other.Get_().(D16_DC34)
			return ok
		}
	case D16_DC35:
		{
			data2, ok := other.Get_().(D16_DC35)
			return ok && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D16_DC33:
		{
			data2, ok := other.Get_().(D16_DC33)
			return ok && data1.Cf54 == data2.Cf54
		}
	case D16_DC36:
		{
			data2, ok := other.Get_().(D16_DC36)
			return ok && data1.Cf56.Equals(data2.Cf56)
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

type D17_DC37 struct {
	Cf57 _dafny.Set
}

func (D17_DC37) isD17() {}

func (CompanionStruct_D17_) Create_DC37_(Cf57 _dafny.Set) D17 {
	return D17{D17_DC37{Cf57}}
}

func (_this D17) Is_DC37() bool {
	_, ok := _this.Get_().(D17_DC37)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D17) Dtor_cf57() _dafny.Set {
	return _this.Get_().(D17_DC37).Cf57
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC37:
		{
			return "D17.DC37" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC37:
		{
			data2, ok := other.Get_().(D17_DC37)
			return ok && data1.Cf57.Equals(data2.Cf57)
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

type D18_DC39 struct {
	Cf59 bool
}

func (D18_DC39) isD18() {}

func (CompanionStruct_D18_) Create_DC39_(Cf59 bool) D18 {
	return D18{D18_DC39{Cf59}}
}

func (_this D18) Is_DC39() bool {
	_, ok := _this.Get_().(D18_DC39)
	return ok
}

type D18_DC40 struct {
	Cf60 *C4
	Cf61 bool
}

func (D18_DC40) isD18() {}

func (CompanionStruct_D18_) Create_DC40_(Cf60 *C4, Cf61 bool) D18 {
	return D18{D18_DC40{Cf60, Cf61}}
}

func (_this D18) Is_DC40() bool {
	_, ok := _this.Get_().(D18_DC40)
	return ok
}

type D18_DC38 struct {
	Cf58 _dafny.Set
}

func (D18_DC38) isD18() {}

func (CompanionStruct_D18_) Create_DC38_(Cf58 _dafny.Set) D18 {
	return D18{D18_DC38{Cf58}}
}

func (_this D18) Is_DC38() bool {
	_, ok := _this.Get_().(D18_DC38)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC39_(false)
}

func (_this D18) Dtor_cf59() bool {
	return _this.Get_().(D18_DC39).Cf59
}

func (_this D18) Dtor_cf60() *C4 {
	return _this.Get_().(D18_DC40).Cf60
}

func (_this D18) Dtor_cf61() bool {
	return _this.Get_().(D18_DC40).Cf61
}

func (_this D18) Dtor_cf58() _dafny.Set {
	return _this.Get_().(D18_DC38).Cf58
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC39:
		{
			return "D18.DC39" + "(" + _dafny.String(data.Cf59) + ")"
		}
	case D18_DC40:
		{
			return "D18.DC40" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D18_DC38:
		{
			return "D18.DC38" + "(" + _dafny.String(data.Cf58) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC39:
		{
			data2, ok := other.Get_().(D18_DC39)
			return ok && data1.Cf59 == data2.Cf59
		}
	case D18_DC40:
		{
			data2, ok := other.Get_().(D18_DC40)
			return ok && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61
		}
	case D18_DC38:
		{
			data2, ok := other.Get_().(D18_DC38)
			return ok && data1.Cf58.Equals(data2.Cf58)
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
	Cf63 bool
	Cf64 _dafny.Int
}

func (D19_DC42) isD19() {}

func (CompanionStruct_D19_) Create_DC42_(Cf63 bool, Cf64 _dafny.Int) D19 {
	return D19{D19_DC42{Cf63, Cf64}}
}

func (_this D19) Is_DC42() bool {
	_, ok := _this.Get_().(D19_DC42)
	return ok
}

type D19_DC41 struct {
	Cf62 _dafny.MultiSet
}

func (D19_DC41) isD19() {}

func (CompanionStruct_D19_) Create_DC41_(Cf62 _dafny.MultiSet) D19 {
	return D19{D19_DC41{Cf62}}
}

func (_this D19) Is_DC41() bool {
	_, ok := _this.Get_().(D19_DC41)
	return ok
}

type D19_DC43 struct {
	Cf65 D19
}

func (D19_DC43) isD19() {}

func (CompanionStruct_D19_) Create_DC43_(Cf65 D19) D19 {
	return D19{D19_DC43{Cf65}}
}

func (_this D19) Is_DC43() bool {
	_, ok := _this.Get_().(D19_DC43)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC42_(false, _dafny.Zero)
}

func (_this D19) Dtor_cf63() bool {
	return _this.Get_().(D19_DC42).Cf63
}

func (_this D19) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D19_DC42).Cf64
}

func (_this D19) Dtor_cf62() _dafny.MultiSet {
	return _this.Get_().(D19_DC41).Cf62
}

func (_this D19) Dtor_cf65() D19 {
	return _this.Get_().(D19_DC43).Cf65
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC42:
		{
			return "D19.DC42" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D19_DC41:
		{
			return "D19.DC41" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D19_DC43:
		{
			return "D19.DC43" + "(" + _dafny.String(data.Cf65) + ")"
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
			return ok && data1.Cf63 == data2.Cf63 && data1.Cf64.Cmp(data2.Cf64) == 0
		}
	case D19_DC41:
		{
			data2, ok := other.Get_().(D19_DC41)
			return ok && data1.Cf62.Equals(data2.Cf62)
		}
	case D19_DC43:
		{
			data2, ok := other.Get_().(D19_DC43)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm5(globalState *GlobalState) bool
	Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool
	M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int)
	F0() _dafny.Int
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
	Fm5(globalState *GlobalState) bool
	Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool
	M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int)
	F0() _dafny.Int
	M2(globalState *GlobalState)
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
	_f3 _dafny.CodePoint
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f3 = _dafny.CodePoint('D')
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

func (_this *C0) Ctor__(f3 _dafny.CodePoint) {
	{
		(_this)._f3 = f3
	}
}
func (_this *C0) F3() _dafny.CodePoint {
	{
		return _this._f3
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f0 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f0 = _dafny.Zero
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

func (_this *C1) F0() _dafny.Int {
	return _this._f0
}
func (_this *C1) Ctor__(f0 _dafny.Int) {
	{
		(_this)._f0 = f0
	}
}
func (_this *C1) Fm5(globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(600))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_410_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})), _dafny.UnicodeSeqOfUtf8Bytes("s")))
	}
}
func (_this *C1) Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(false) || (true)
	}
}
func (_this *C1) Fm15(p0 _dafny.Map, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F0()
	}
}
func (_this *C1) Fm16(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((_dafny.MultiSetOf(false)).IsSubsetOf(_dafny.MultiSetOf(false, !(false))))
	}
}
func (_this *C1) M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _411_v0 bool
		_ = _411_v0
		_411_v0 = true
		var _412_i0 _dafny.Int
		_ = _412_i0
		_412_i0 = _dafny.Zero
		{
			for (func() bool {
				if true {
					return _411_v0
				}
				return (func() bool {
					if _411_v0 {
						return !(_411_v0)
					}
					return _411_v0
				})()
			})() {
				{
					if (_412_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_412_i0 = (_412_i0).Plus(_dafny.One)
					var _413_v1 _dafny.Array
					_ = _413_v1
					var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
					_ = _nw70
					_413_v1 = _nw70
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_413_v1), 0))
					_ = _index65
					(_413_v1).ArraySet1((_this).F0(), (_index65).Int())
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_413_v1), 0))
					_ = _index66
					(_413_v1).ArraySet1(p1, (_index66).Int())
					var _414_v2 bool
					_ = _414_v2
					var _out12 bool
					_ = _out12
					_out12 = Companion_Default___.M0(((_dafny.Zero).Minus(_dafny.IntOfInt64(-237))).Cmp(p1) > 0, (_this).F0(), globalState)
					_414_v2 = _out12
					var _415_v3 _dafny.CodePoint
					_ = _415_v3
					_415_v3 = _dafny.CodePoint('h')
					var _416_v4 _dafny.Map
					_ = _416_v4
					_416_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v3, _dafny.IntOfInt64(-500))
					_416_v4 = (_416_v4).Merge(_416_v4)
					var _417_v5 *C0
					_ = _417_v5
					var _nw71 *C0 = New_C0_()
					_ = _nw71
					_nw71.Ctor__(_415_v3)
					_417_v5 = _nw71
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _418_v6 _dafny.CodePoint
		_ = _418_v6
		_418_v6 = _dafny.CodePoint('e')
		var _419_v7 *C0
		_ = _419_v7
		var _nw72 *C0 = New_C0_()
		_ = _nw72
		_nw72.Ctor__(_418_v6)
		_419_v7 = _nw72
		var _420_v8 _dafny.Sequence
		_ = _420_v8
		_420_v8 = _dafny.UnicodeSeqOfUtf8Bytes("fce")
		var _421_v9 _dafny.Map
		_ = _421_v9
		_421_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v8, _411_v0)
		var _422_v10 _dafny.Sequence
		_ = _422_v10
		_422_v10 = _dafny.SeqOf(_dafny.IntOfInt64(584))
		r0 = (_this).Fm15(_421_v9, !(_411_v0), _dafny.Companion_Sequence_.Equal(_422_v10, _422_v10), globalState)
		var _423_v11 _dafny.Array
		_ = _423_v11
		var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw73
		_423_v11 = _nw73
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_423_v11), 0))
		_ = _index67
		(_423_v11).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F0(), (_dafny.Zero).Minus(p1))), (_index67).Int())
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_423_v11), 0))
		_ = _index68
		(_423_v11).ArraySet1((func() _dafny.Int {
			if _411_v0 {
				return (_this).F0()
			}
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("flbjqa")).Cardinality())
		})(), (_index68).Int())
		var _424_v12 _dafny.MultiSet
		_ = _424_v12
		_424_v12 = _dafny.MultiSetOf(_411_v0)
		var _425_v13 _dafny.Map
		_ = _425_v13
		_425_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC7_(_422_v10), (_424_v12).Equals(_424_v12))
		var _426_v14 D3
		_ = _426_v14
		_426_v14 = Companion_D3_.Create_DC7_(_422_v10)
		_425_v13 = (_425_v13).Update(_426_v14, (_424_v12).IsSubsetOf(_424_v12))
		var _427_v15 _dafny.Map
		_ = _427_v15
		_427_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(718))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}((func(_428_v7 *C0) func(_dafny.Int) _dafny.CodePoint {
			return func(_429_i2 _dafny.Int) _dafny.CodePoint {
				return (_428_v7).F3()
			}
		})(_419_v7))))
		var _430_v16 _dafny.Array
		_ = _430_v16
		var _nwElement0_11 _dafny.Sequence = (func() _dafny.Sequence {
			if _411_v0 {
				return _420_v8
			}
			return _dafny.Companion_Sequence_.Update(_420_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_420_v8).Cardinality()), _dafny.IntOfUint32((_420_v8).Cardinality()))).Uint32(), _dafny.CodePoint('y'))
		})()
		_ = _nwElement0_11
		var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(25))
		_ = _nw74
		(_nw74).ArraySet1(_nwElement0_11, 0)
		(_nw74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("btdr"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_431_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))), 1)
		(_nw74).ArraySet1(_420_v8, 2)
		(_nw74).ArraySet1((func() _dafny.Sequence {
			if (_427_v15).Contains((_this).F0()) {
				return (_427_v15).Get((_this).F0()).(_dafny.Sequence)
			}
			return _420_v8
		})(), 3)
		(_nw74).ArraySet1(_420_v8, 4)
		(_nw74).ArraySet1(_dafny.Companion_Sequence_.Update(_420_v8, (Companion_Default___.SafeIndex((_423_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_423_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_420_v8).Cardinality()))).Uint32(), Companion_Default___.Fm17(p1, globalState)), 5)
		(_nw74).ArraySet1(_420_v8, 6)
		(_nw74).ArraySet1(_420_v8, 7)
		(_nw74).ArraySet1(_420_v8, 8)
		(_nw74).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cx"), 9)
		(_nw74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_420_v8, _420_v8), 10)
		(_nw74).ArraySet1(_420_v8, 11)
		(_nw74).ArraySet1(_420_v8, 12)
		(_nw74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_420_v8, _dafny.UnicodeSeqOfUtf8Bytes("u")), 13)
		(_nw74).ArraySet1(_420_v8, 14)
		(_nw74).ArraySet1(_420_v8, 15)
		(_nw74).ArraySet1(_420_v8, 16)
		(_nw74).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(730))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}((func(_432_v7 *C0) func(_dafny.Int) _dafny.CodePoint {
			return func(_433_i3 _dafny.Int) _dafny.CodePoint {
				return (_432_v7).F3()
			}
		})(_419_v7))), 17)
		(_nw74).ArraySet1(_420_v8, 18)
		(_nw74).ArraySet1(_420_v8, 19)
		(_nw74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_420_v8, _420_v8), 20)
		(_nw74).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("oxyrf"), 21)
		(_nw74).ArraySet1(_420_v8, 22)
		(_nw74).ArraySet1(_420_v8, 23)
		(_nw74).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-916))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_434_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_435_i4 _dafny.Int) _dafny.CodePoint {
				return _434_v6
			}
		})(_418_v6))), 24)
		_430_v16 = _nw74
		var _436_v17 D0
		_ = _436_v17
		_436_v17 = Companion_D0_.Create_DC0_(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_420_v8, _420_v8)))
		var _437_v18 _dafny.Map
		_ = _437_v18
		_437_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v0, _411_v0)
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_423_v11), 0))
		_ = _index69
		var _rhs63 _dafny.Array = (func() _dafny.Array {
			if _411_v0 {
				return _430_v16
			}
			return _430_v16
		})()
		_ = _rhs63
		var _rhs64 _dafny.Int = p1
		_ = _rhs64
		var _rhs65 _dafny.Sequence = (func() _dafny.Sequence {
			if _411_v0 {
				return _420_v8
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("jhobkl")
		})()
		_ = _rhs65
		var _rhs66 D0 = (func() D0 {
			if ((_this).F0()).Cmp((_dafny.Zero).Minus((_423_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_423_v11), 0))).Int()).(_dafny.Int))) >= 0 {
				return Companion_D0_.Create_DC0_((func() bool {
					if (_437_v18).Contains(_411_v0) {
						return (_437_v18).Get(_411_v0).(bool)
					}
					return true
				})())
			}
			return _436_v17
		})()
		_ = _rhs66
		var _lhs20 _dafny.Array = _423_v11
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_423_v11), 0))
		_ = _lhs21
		_430_v16 = _rhs63
		(_lhs20).ArraySet1(_rhs64, (_lhs21).Int())
		_420_v8 = _rhs65
		_436_v17 = _rhs66
		r0 = (_this).F0()
		r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_420_v8, _420_v8), _420_v8)).Cardinality())
		r2 = p1
		return r0, r1, r2
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f0 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f0 = _dafny.Zero
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

func (_this *C2) F0() _dafny.Int {
	return _this._f0
}
func (_this *C2) Ctor__(f0 _dafny.Int) {
	{
		(_this)._f0 = f0
	}
}
func (_this *C2) Fm5(globalState *GlobalState) bool {
	{
		return ((func() bool {
			if true {
				return true
			}
			return true
		})()) && (!(!(true) || (false)))
	}
}
func (_this *C2) Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((false) || (true))
	}
}
func (_this *C2) M2(globalState *GlobalState) {
	{
		var _438_v0 bool
		_ = _438_v0
		_438_v0 = true
		var _439_v1 _dafny.Array
		_ = _439_v1
		var _nwElement0_12 bool = _438_v0
		_ = _nwElement0_12
		var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(2))
		_ = _nw75
		(_nw75).ArraySet1(_nwElement0_12, 0)
		(_nw75).ArraySet1(_438_v0, 1)
		_439_v1 = _nw75
		var _440_v2 _dafny.MultiSet
		_ = _440_v2
		_440_v2 = _dafny.MultiSetOf(_439_v1)
		var _441_v3 _dafny.Sequence
		_ = _441_v3
		_441_v3 = _dafny.SeqOf((_440_v2).Cardinality())
		var _pat_let_tv5 = _441_v3
		_ = _pat_let_tv5
		var _pat_let_tv6 = _441_v3
		_ = _pat_let_tv6
		_441_v3 = func(_source2 D0) _dafny.Sequence {
			if _source2.Is_DC1() {
				var _442___mcc_h0 bool = _source2.Get_().(D0_DC1).Cf1
				_ = _442___mcc_h0
				var _443_cf1 bool = _442___mcc_h0
				_ = _443_cf1
				return _pat_let_tv5
			} else {
				var _444___mcc_h1 bool = _source2.Get_().(D0_DC0).Cf0
				_ = _444___mcc_h1
				var _445_cf0 bool = _444___mcc_h1
				_ = _445_cf0
				return _pat_let_tv6
			}
		}(Companion_D0_.Create_DC0_(_438_v0))
		var _446_v4 _dafny.Sequence
		_ = _446_v4
		_446_v4 = _dafny.UnicodeSeqOfUtf8Bytes("bdfea")
		var _447_v5 D2
		_ = _447_v5
		_447_v5 = Companion_D2_.Create_DC4_(_446_v4, _438_v0)
		var _pat_let_tv7 = _438_v0
		_ = _pat_let_tv7
		var _pat_let_tv8 = _438_v0
		_ = _pat_let_tv8
		var _pat_let_tv9 = _438_v0
		_ = _pat_let_tv9
		var _pat_let_tv10 = _438_v0
		_ = _pat_let_tv10
		var _pat_let_tv11 = _438_v0
		_ = _pat_let_tv11
		_438_v0 = func(_source3 D2) bool {
			if _source3.Is_DC4() {
				var _448___mcc_h2 _dafny.Sequence = _source3.Get_().(D2_DC4).Cf4
				_ = _448___mcc_h2
				var _449___mcc_h3 bool = _source3.Get_().(D2_DC4).Cf5
				_ = _449___mcc_h3
				var _450_cf5 bool = _449___mcc_h3
				_ = _450_cf5
				var _451_cf4 _dafny.Sequence = _448___mcc_h2
				_ = _451_cf4
				return _pat_let_tv7
			} else if _source3.Is_DC5() {
				return _pat_let_tv8
			} else if _source3.Is_DC3() {
				var _452___mcc_h4 _dafny.Sequence = _source3.Get_().(D2_DC3).Cf3
				_ = _452___mcc_h4
				var _453_cf3 _dafny.Sequence = _452___mcc_h4
				_ = _453_cf3
				return !((_pat_let_tv9) || (_pat_let_tv10))
			} else {
				var _454___mcc_h5 D2 = _source3.Get_().(D2_DC6).Cf6
				_ = _454___mcc_h5
				var _455_cf6 D2 = _454___mcc_h5
				_ = _455_cf6
				return _pat_let_tv11
			}
		}(_447_v5)
		var _hi2 _dafny.Int = _dafny.IntOfInt64(221)
		_ = _hi2
		for _456_i0 := (_this).F0(); _456_i0.Cmp(_hi2) < 0; _456_i0 = _456_i0.Plus(_dafny.One) {
			var _457_v6 _dafny.Array
			_ = _457_v6
			var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
			_ = _nw76
			_457_v6 = _nw76
			var _458_v7 _dafny.Map
			_ = _458_v7
			_458_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_439_v1, (_this).F0())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_457_v6), 0))
			_ = _index70
			(_457_v6).ArraySet1(_458_v7, (_index70).Int())
			var _459_v8 _dafny.Map
			_ = _459_v8
			_459_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5(globalState), _458_v7)
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_457_v6), 0))
			_ = _index71
			(_457_v6).ArraySet1(((func() _dafny.Map {
				if (_459_v8).Contains(_438_v0) {
					return (_459_v8).Get(_438_v0).(_dafny.Map)
				}
				return _458_v7
			})()).Update(_439_v1, _456_i0), (_index71).Int())
			var _460_v9 _dafny.Int
			_ = _460_v9
			_460_v9 = _dafny.IntOfInt64(715)
			var _461_v10 _dafny.MultiSet
			_ = _461_v10
			_461_v10 = _dafny.MultiSetOf(_438_v0, !(_438_v0), _438_v0, _438_v0, true)
			_460_v9 = (func() _dafny.Int {
				if (_461_v10).Contains(_dafny.Companion_Sequence_.IsPrefixOf(_446_v4, _446_v4)) {
					return (_461_v10).Multiplicity(_dafny.Companion_Sequence_.IsPrefixOf(_446_v4, _446_v4))
				}
				return _456_i0
			})()
			var _462_v11 _dafny.Sequence
			_ = _462_v11
			_462_v11 = _dafny.SeqOf(_438_v0, _438_v0)
			var _463_v12 _dafny.CodePoint
			_ = _463_v12
			_463_v12 = _dafny.CodePoint('t')
			var _464_v13 bool
			_ = _464_v13
			var _out13 bool
			_ = _out13
			_out13 = Companion_Default___.M0(((_this).F0()).Cmp(_456_i0) < 0, Companion_Default___.Fm3(_dafny.SetOf(Companion_Default___.Fm2(_438_v0, _dafny.IntOfInt64(-904), _438_v0, _438_v0, globalState), _438_v0, (_462_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.IntOfUint32((_462_v11).Cardinality()))).Uint32()).(bool), _438_v0, _438_v0), _463_v12, _460_v9, _460_v9, globalState), globalState)
			_464_v13 = _out13
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_439_v1), 0))
			_ = _index72
			(_439_v1).ArraySet1(false, (_index72).Int())
			var _465_v14 _dafny.Map
			_ = _465_v14
			_465_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_456_i0).Cmp((_this).F0()) > 0, _460_v9)
			var _466_v15 _dafny.Map
			_ = _466_v15
			_466_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v14, _462_v11)
			var _467_v16 _dafny.Set
			_ = _467_v16
			_467_v16 = _dafny.SetOf(_438_v0, _438_v0, true, _464_v13)
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_439_v1), 0))
			_ = _index73
			var _rhs67 _dafny.Int = (func() _dafny.Int {
				if (_465_v14).Contains(true) {
					return (_465_v14).Get(true).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(_460_v9))
			})()
			_ = _rhs67
			var _rhs68 bool = (_456_i0).Cmp((_dafny.Zero).Minus(_456_i0)) >= 0
			_ = _rhs68
			var _rhs69 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_462_v11, _dafny.Companion_Sequence_.Concatenate(_462_v11, (func() _dafny.Sequence {
				if (_466_v15).Contains(_465_v14) {
					return (_466_v15).Get(_465_v14).(_dafny.Sequence)
				}
				return _462_v11
			})()))
			_ = _rhs69
			var _rhs70 bool = !((_dafny.SetOf(_464_v13)).Equals(_467_v16))
			_ = _rhs70
			var _lhs22 _dafny.Array = _439_v1
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_439_v1), 0))
			_ = _lhs23
			_460_v9 = _rhs67
			(_lhs22).ArraySet1(_rhs68, (_lhs23).Int())
			_462_v11 = _rhs69
			_438_v0 = _rhs70
		}
		var _468_v17 _dafny.Int
		_ = _468_v17
		_468_v17 = _dafny.IntOfInt64(653)
		var _469_v18 _dafny.Set
		_ = _469_v18
		_469_v18 = _dafny.SetOf(_438_v0)
		var _470_v19 _dafny.CodePoint
		_ = _470_v19
		_470_v19 = _dafny.CodePoint('h')
		var _471_v20 _dafny.Map
		_ = _471_v20
		_471_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v17, (_this).Fm5(globalState))
		var _472_v21 _dafny.MultiSet
		_ = _472_v21
		_472_v21 = _dafny.MultiSetOf(_471_v20, _471_v20)
		_468_v17 = Companion_Default___.Fm3((_469_v18).Union(_469_v18), _470_v19, (_468_v17).Plus((func() _dafny.Int {
			if (_472_v21).Contains(_471_v20) {
				return (_472_v21).Multiplicity(_471_v20)
			}
			return _468_v17
		})()), (_this).F0(), globalState)
		var _hi3 _dafny.Int = _468_v17
		_ = _hi3
		for _473_i1 := _468_v17; _473_i1.Cmp(_hi3) < 0; _473_i1 = _473_i1.Plus(_dafny.One) {
			var _474_v22 _dafny.MultiSet
			_ = _474_v22
			_474_v22 = _dafny.MultiSetOf(_473_i1, (_this).F0())
			_446_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_446_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(585))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_475_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_476_i2 _dafny.Int) _dafny.CodePoint {
					return _475_v19
				}
			})(_470_v19)))), Companion_Default___.Fm19(_474_v22, _473_i1, (_this).F0(), _468_v17, globalState))
			var _477_v23 _dafny.Array
			_ = _477_v23
			var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw77
			_477_v23 = _nw77
			var _478_v24 _dafny.Set
			_ = _478_v24
			_478_v24 = _dafny.SetOf(_477_v23)
			var _rhs71 _dafny.Sequence = _441_v3
			_ = _rhs71
			var _rhs72 bool = !(_438_v0) || (_438_v0)
			_ = _rhs72
			var _rhs73 _dafny.Int = (_478_v24).Cardinality()
			_ = _rhs73
			var _rhs74 _dafny.Int = (_dafny.Zero).Minus(_468_v17)
			_ = _rhs74
			_441_v3 = _rhs71
			_438_v0 = _rhs72
			_468_v17 = _rhs73
			_468_v17 = _rhs74
			var _479_v25 _dafny.Map
			_ = _479_v25
			_479_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_438_v0), _439_v1)
			_439_v1 = (func() _dafny.Array {
				if (_479_v25).Contains(_438_v0) {
					return (_479_v25).Get(_438_v0).(_dafny.Array)
				}
				return _439_v1
			})()
			var _480_v26 *C0
			_ = _480_v26
			var _nw78 *C0 = New_C0_()
			_ = _nw78
			_nw78.Ctor__(_dafny.CodePoint('n'))
			_480_v26 = _nw78
		}
		var _481_v27 _dafny.Array
		_ = _481_v27
		var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw79
		_481_v27 = _nw79
		var _482_v28 _dafny.MultiSet
		_ = _482_v28
		_482_v28 = _dafny.MultiSetOf(_481_v27, _481_v27, _481_v27, _481_v27, _481_v27)
		_468_v17 = Companion_Default___.SafeModuloInt((_this).F0(), Companion_Default___.SafeModuloInt((_482_v28).Cardinality(), (_this).F0()))
	}
}
func (_this *C2) M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _483_v0 bool
		_ = _483_v0
		_483_v0 = false
		_483_v0 = _483_v0
		var _484_v1 _dafny.Set
		_ = _484_v1
		_484_v1 = _dafny.SetOf((_this).F0(), (_this).F0())
		_483_v0 = !((_484_v1).Intersection(_484_v1)).Equals(_484_v1)
		var _485_v2 _dafny.Array
		_ = _485_v2
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_9
		var _nw80 _dafny.Array
		_ = _nw80
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw80 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.CodePoint = func(_486_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw80 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw80).ArraySet1CodePoint(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw80).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_485_v2 = _nw80
		for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_485_v2), 0))); ; {
			_guard_loop_0, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _487_i0 _dafny.Int
			_487_i0 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_487_i0).Sign() != -1) && ((_487_i0).Cmp(_dafny.ArrayLen((_485_v2), 0)) < 0)) {
				(_485_v2).ArraySet1CodePoint(_dafny.CodePoint('h'), (_487_i0).Int())
			}
		}
		var _488_v3 _dafny.MultiSet
		_ = _488_v3
		_488_v3 = _dafny.MultiSetOf(_483_v0, (_this).Fm5(globalState), (_this).Fm5(globalState))
		var _489_v4 _dafny.Map
		_ = _489_v4
		_489_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_488_v3).Cardinality(), p1)
		var _490_v5 _dafny.Array
		_ = _490_v5
		var _nwElement0_13 bool = false
		_ = _nwElement0_13
		var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(20))
		_ = _nw81
		(_nw81).ArraySet1(_nwElement0_13, 0)
		(_nw81).ArraySet1(_483_v0, 1)
		(_nw81).ArraySet1(Companion_Default___.Fm2(_483_v0, _dafny.IntOfInt64(847), _483_v0, _483_v0, globalState), 2)
		(_nw81).ArraySet1(_483_v0, 3)
		(_nw81).ArraySet1(_483_v0, 4)
		(_nw81).ArraySet1(_483_v0, 5)
		(_nw81).ArraySet1(_483_v0, 6)
		(_nw81).ArraySet1(_483_v0, 7)
		(_nw81).ArraySet1(_483_v0, 8)
		(_nw81).ArraySet1(_483_v0, 9)
		(_nw81).ArraySet1(_483_v0, 10)
		(_nw81).ArraySet1(_483_v0, 11)
		(_nw81).ArraySet1(true, 12)
		(_nw81).ArraySet1(_483_v0, 13)
		(_nw81).ArraySet1(_483_v0, 14)
		(_nw81).ArraySet1(true, 15)
		(_nw81).ArraySet1(_483_v0, 16)
		(_nw81).ArraySet1(_483_v0, 17)
		(_nw81).ArraySet1(_483_v0, 18)
		(_nw81).ArraySet1(_483_v0, 19)
		_490_v5 = _nw81
		var _491_v6 _dafny.Array
		_ = _491_v6
		_491_v6 = _490_v5
		var _492_v7 _dafny.Map
		_ = _492_v7
		_492_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_491_v6, _dafny.IntOfInt64(-32))
		var _493_v8 _dafny.Array
		_ = _493_v8
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw82
		_493_v8 = _nw82
		var _494_v9 _dafny.MultiSet
		_ = _494_v9
		_494_v9 = _dafny.MultiSetOf(_493_v8)
		var _495_v10 _dafny.Set
		_ = _495_v10
		_495_v10 = _dafny.SetOf(_483_v0, _483_v0)
		var _496_v11 _dafny.CodePoint
		_ = _496_v11
		_496_v11 = _dafny.CodePoint('k')
		var _497_v12 _dafny.Sequence
		_ = _497_v12
		_497_v12 = _dafny.SeqOf(_dafny.SeqOf(_483_v0))
		var _498_v13 _dafny.Sequence
		_ = _498_v13
		_498_v13 = _dafny.UnicodeSeqOfUtf8Bytes("ltx")
		var _499_v14 _dafny.Sequence
		_ = _499_v14
		_499_v14 = _dafny.SeqOf(true, _483_v0)
		var _500_v15 D7
		_ = _500_v15
		_500_v15 = Companion_D7_.Create_DC16_(_499_v14)
		var _501_v16 _dafny.Array
		_ = _501_v16
		var _nwElement0_14 _dafny.Int = _dafny.IntOfInt64(98)
		_ = _nwElement0_14
		var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(21))
		_ = _nw83
		(_nw83).ArraySet1(_nwElement0_14, 0)
		(_nw83).ArraySet1((_489_v4).Cardinality(), 1)
		(_nw83).ArraySet1(Companion_Default___.SafeModuloInt((_this).F0(), (_this).F0()), 2)
		(_nw83).ArraySet1((_dafny.IntOfInt64(124)).Times((_this).F0()), 3)
		(_nw83).ArraySet1((_492_v7).Cardinality(), 4)
		(_nw83).ArraySet1(p1, 5)
		(_nw83).ArraySet1((_494_v9).Cardinality(), 6)
		(_nw83).ArraySet1((_this).F0(), 7)
		(_nw83).ArraySet1(p1, 8)
		(_nw83).ArraySet1(_dafny.IntOfInt64(71), 9)
		(_nw83).ArraySet1((_this).F0(), 10)
		(_nw83).ArraySet1(p1, 11)
		(_nw83).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F0(), (_dafny.Zero).Minus(Companion_Default___.Fm3(_495_v10, _496_v11, _dafny.IntOfUint32((_497_v12).Cardinality()), p1, globalState))), 12)
		(_nw83).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_498_v13, _dafny.UnicodeSeqOfUtf8Bytes("hbu"))).Cardinality()), 13)
		(_nw83).ArraySet1(p1, 14)
		(_nw83).ArraySet1((func() _dafny.Int {
			if (_488_v3).Contains(_483_v0) {
				return (_488_v3).Multiplicity(_483_v0)
			}
			return (_this).F0()
		})(), 15)
		(_nw83).ArraySet1((func() _dafny.Int {
			if _483_v0 {
				return _dafny.IntOfUint32(((_500_v15).Dtor_cf26()).Cardinality())
			}
			return (_this).F0()
		})(), 16)
		(_nw83).ArraySet1(p1, 17)
		(_nw83).ArraySet1((_this).F0(), 18)
		(_nw83).ArraySet1(p1, 19)
		(_nw83).ArraySet1((_this).F0(), 20)
		_501_v16 = _nw83
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_501_v16), 0))); ; {
			_guard_loop_1, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _502_i2 _dafny.Int
			_502_i2 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_502_i2).Sign() != -1) && ((_502_i2).Cmp(_dafny.ArrayLen((_501_v16), 0)) < 0)) {
				(_501_v16).ArraySet1((_502_i2).Plus(p1), (_502_i2).Int())
			}
		}
		var _503_v17 D5
		_ = _503_v17
		_503_v17 = Companion_D5_.Create_DC13_(p1, _495_v10, _496_v11, _dafny.IntOfInt64(701), _483_v0)
		var _504_i3 _dafny.Int
		_ = _504_i3
		_504_i3 = _dafny.Zero
		{
			for (_503_v17).Dtor_cf21() {
				{
					if (_504_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_504_i3 = (_504_i3).Plus(_dafny.One)
					var _505_v18 _dafny.Sequence
					_ = _505_v18
					_505_v18 = _dafny.SeqOf(_493_v8)
					var _506_v19 _dafny.Map
					_ = _506_v19
					_506_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_505_v18, _505_v18), _483_v0)
					var _507_v20 D2
					_ = _507_v20
					_507_v20 = Companion_D2_.Create_DC4_(_498_v13, !(_483_v0))
					_506_v19 = (_506_v19).Update(_505_v18, (_507_v20).Dtor_cf5())
					var _508_v21 _dafny.Map
					_ = _508_v21
					_508_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v0, _496_v11)
					var _509_v22 _dafny.Sequence
					_ = _509_v22
					_509_v22 = _dafny.SeqOf((_this).F0(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v0, _483_v0)).Cardinality(), (Companion_Default___.Fm20(_483_v0, _508_v21, true, _496_v11, globalState)).Cardinality(), (_this).F0(), p1)
					var _rhs75 bool = ((_this).F0()).Cmp(Companion_Default___.SafeModuloInt((_this).F0(), (_this).F0())) != 0
					_ = _rhs75
					var _rhs76 _dafny.CodePoint = _496_v11
					_ = _rhs76
					var _rhs77 _dafny.Int = p1
					_ = _rhs77
					var _rhs78 bool = _dafny.Companion_Sequence_.Equal(_509_v22, _509_v22)
					_ = _rhs78
					var _rhs79 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F0(), (_this).F0())
					_ = _rhs79
					_483_v0 = _rhs75
					_496_v11 = _rhs76
					r1 = _rhs77
					_483_v0 = _rhs78
					r2 = _rhs79
					_483_v0 = Companion_Default___.Fm2((_483_v0) && (_483_v0), ((_this).F0()).Plus((_this).F0()), !(_483_v0), _483_v0, globalState)
					_483_v0 = (_dafny.IntOfInt64(357)).Cmp((_495_v10).Cardinality()) <= 0
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _hi4 _dafny.Int = p1
		_ = _hi4
		for _510_i4 := _dafny.IntOfUint32((_dafny.SeqOf((_this).F0())).Cardinality()); _510_i4.Cmp(_hi4) < 0; _510_i4 = _510_i4.Plus(_dafny.One) {
			_483_v0 = ((_dafny.Zero).Minus((_this).F0())).Cmp(p1) == 0
			_483_v0 = _483_v0
			if _483_v0 {
				var _511_v23 _dafny.Map
				_ = _511_v23
				_511_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_498_v13).Cardinality()), _483_v0)
				var _512_v24 _dafny.Map
				_ = _512_v24
				_512_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v0, _dafny.IntOfUint32((Companion_Default___.Fm19(Companion_Default___.Fm22(globalState), _510_i4, (_this).F0(), (_511_v23).Cardinality(), globalState)).Cardinality()))
				var _513_v25 _dafny.Sequence
				_ = _513_v25
				_513_v25 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v0, (_this).F0()), Companion_Default___.Fm21(globalState), _512_v24)
				_513_v25 = _513_v25
				var _rhs80 _dafny.Int = _510_i4
				_ = _rhs80
				var _rhs81 _dafny.CodePoint = _496_v11
				_ = _rhs81
				var _rhs82 bool = ((_510_i4).Minus(p1)).Cmp(p1) < 0
				_ = _rhs82
				r1 = _rhs80
				_496_v11 = _rhs81
				_483_v0 = _rhs82
				_483_v0 = _483_v0
				r2 = Companion_Default___.Fm3(_495_v10, _496_v11, p1, (p1).Plus(_dafny.IntOfInt64(-34)), globalState)
				r0 = _510_i4
			} else {
				r2 = _510_i4
				_493_v8 = _493_v8
				var _514_v26 _dafny.MultiSet
				_ = _514_v26
				_514_v26 = _dafny.MultiSetOf(_488_v3)
				var _515_v27 _dafny.MultiSet
				_ = _515_v27
				_515_v27 = _dafny.MultiSetOf(_dafny.IntOfUint32((_498_v13).Cardinality()), p1, (func() _dafny.Int {
					if (_514_v26).Contains(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, _483_v0))) {
						return (_514_v26).Multiplicity(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, _483_v0)))
					}
					return _dafny.IntOfUint32((_498_v13).Cardinality())
				})())
				var _516_v28 _dafny.Map
				_ = _516_v28
				_516_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if _483_v0 {
						return _490_v5
					}
					return _490_v5
				})(), _515_v27)
				_516_v28 = (_516_v28).Merge(_516_v28)
				_493_v8 = _501_v16
				var _517_v29 _dafny.Array
				_ = _517_v29
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
				_ = _nw84
				_517_v29 = _nw84
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_517_v29), 0))
				_ = _index74
				(_517_v29).ArraySet1(_dafny.SeqOf(_510_i4), (_index74).Int())
				var _518_v30 _dafny.Map
				_ = _518_v30
				_518_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _483_v0)
				var _519_v31 _dafny.Sequence
				_ = _519_v31
				_519_v31 = _dafny.SeqOf((_488_v3).Cardinality(), Companion_Default___.SafeDivisionInt((_518_v30).Cardinality(), (_this).F0()))
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_517_v29), 0))
				_ = _index75
				(_517_v29).ArraySet1(_519_v31, (_index75).Int())
			}
			if (_510_i4).Cmp(_510_i4) != 0 {
				var _520_v32 _dafny.MultiSet
				_ = _520_v32
				_520_v32 = _dafny.MultiSetOf(_510_i4)
				_483_v0 = (((_520_v32).Cardinality()).Times((_this).F0())).Cmp(p1) == 0
				r0 = (func() _dafny.Int {
					if false {
						return (_this).F0()
					}
					return p1
				})()
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_490_v5), 0))
				_ = _index76
				(_490_v5).ArraySet1(!((_this).Fm6(_483_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-317))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}(func(_521_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})), _498_v13, _510_i4, globalState)), (_index76).Int())
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_490_v5), 0))
				_ = _index77
				(_490_v5).ArraySet1((func() bool {
					if _483_v0 {
						return !((_dafny.IntOfInt64(-183)).Cmp(_510_i4) >= 0)
					}
					return _483_v0
				})(), (_index77).Int())
				r1 = _dafny.IntOfUint32((_498_v13).Cardinality())
				r0 = Companion_Default___.Fm3((_495_v10).Intersection(_495_v10), _496_v11, p1, ((_this).F0()).Minus(p1), globalState)
			} else {
				_496_v11 = _496_v11
				var _522_v33 _dafny.Map
				_ = _522_v33
				_522_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v0, _493_v8)
				_522_v33 = (_522_v33).Update(_483_v0, _501_v16)
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_490_v5), 0))
				_ = _index78
				(_490_v5).ArraySet1(_483_v0, (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_490_v5), 0))
				_ = _index79
				(_490_v5).ArraySet1((func() bool {
					if false {
						return _483_v0
					}
					return (_499_v14).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_499_v14).Cardinality()))).Uint32()).(bool)
				})(), (_index79).Int())
				var _523_v34 _dafny.Map
				_ = _523_v34
				_523_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC18_(p1), (_490_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_490_v5), 0))).Int()).(bool))
				var _524_v36 D7
				_ = _524_v36
				_524_v36 = Companion_D7_.Create_DC18_((_this).F0())
				var _525_v37 _dafny.Map
				_ = _525_v37
				_525_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, (_this).F0())), (func() _dafny.Map {
					if _483_v0 {
						return _523_v34
					}
					return func() _dafny.Map {
						var _coll15 = _dafny.NewMapBuilder()
						_ = _coll15
						for _iter17 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v36, _483_v0)).Keys().Elements()); ; {
							_compr_15, _ok17 := _iter17()
							if !_ok17 {
								break
							}
							var _526_v35 D7
							_526_v35 = interface{}(_compr_15).(D7)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v36, _483_v0)).Contains(_526_v35) {
								_coll15.Add(_526_v35, (_490_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_490_v5), 0))).Int()).(bool))
							}
						}
						return _coll15.ToMap()
					}()
				})())
				_525_v37 = (_525_v37).Merge((_525_v37).Merge(_525_v37))
				_496_v11 = _496_v11
			}
		}
		var _527_v38 _dafny.Sequence
		_ = _527_v38
		_527_v38 = _dafny.SeqOf(_dafny.IntOfUint32((_498_v13).Cardinality()))
		r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_527_v38, _527_v38)).Cardinality())
		r1 = p1
		r2 = p1
		return r0, r1, r2
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f0 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f0 = _dafny.Zero
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

func (_this *C3) F0() _dafny.Int {
	return _this._f0
}
func (_this *C3) Ctor__(f0 _dafny.Int) {
	{
		(_this)._f0 = f0
	}
}
func (_this *C3) Fm5(globalState *GlobalState) bool {
	{
		return !(!(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jdrpodwh"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg43 _dafny.Int) interface{} {
				return coer43(arg43)
			}
		}(func(_528_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("hehmjmhcu")))))
	}
}
func (_this *C3) Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		if true {
			return !((func() bool {
				if false {
					return false
				}
				return true
			})())
		} else {
			return (true) && (true)
		}
	}
}
func (_this *C3) Fm13(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C3) M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _529_v0 _dafny.Set
		_ = _529_v0
		_529_v0 = _dafny.SetOf(p1, (_this).F0())
		var _530_v1 _dafny.Set
		_ = _530_v1
		_530_v1 = _dafny.SetOf(_529_v0, _529_v0, _529_v0)
		var _531_v2 bool
		_ = _531_v2
		_531_v2 = true
		var _532_v3 _dafny.Sequence
		_ = _532_v3
		_532_v3 = _dafny.SeqOf(_530_v1)
		var _533_v4 _dafny.Map
		_ = _533_v4
		_533_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _531_v2)
		var _534_v5 _dafny.Set
		_ = _534_v5
		_534_v5 = _dafny.SetOf(_531_v2)
		var _535_v6 _dafny.CodePoint
		_ = _535_v6
		_535_v6 = _dafny.CodePoint('j')
		var _536_v7 _dafny.Map
		_ = _536_v7
		_536_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_529_v0, _531_v2)
		var _537_v9 _dafny.Sequence
		_ = _537_v9
		_537_v9 = _dafny.SeqOf(_529_v0, _529_v0)
		var _538_v11 _dafny.Array
		_ = _538_v11
		var _nwElement0_15 _dafny.Set = _530_v1
		_ = _nwElement0_15
		var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(21))
		_ = _nw85
		(_nw85).ArraySet1(_nwElement0_15, 0)
		(_nw85).ArraySet1(_530_v1, 1)
		(_nw85).ArraySet1(_530_v1, 2)
		(_nw85).ArraySet1((_530_v1).Intersection(_530_v1), 3)
		(_nw85).ArraySet1((_530_v1).Union(_530_v1), 4)
		(_nw85).ArraySet1(_530_v1, 5)
		(_nw85).ArraySet1((func() _dafny.Set {
			if _531_v2 {
				return (_532_v3).Select((Companion_Default___.SafeIndex((_533_v4).Cardinality(), _dafny.IntOfUint32((_532_v3).Cardinality()))).Uint32()).(_dafny.Set)
			}
			return _530_v1
		})(), 6)
		(_nw85).ArraySet1(_530_v1, 7)
		(_nw85).ArraySet1((_532_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3(_534_v5, _535_v6, (_this).F0(), p1, globalState), _dafny.IntOfUint32((_532_v3).Cardinality()))).Uint32()).(_dafny.Set), 8)
		(_nw85).ArraySet1(func() _dafny.Set {
			var _coll16 = _dafny.NewBuilder()
			_ = _coll16
			for _iter18 := _dafny.Iterate((_536_v7).Keys().Elements()); ; {
				_compr_16, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _539_v8 _dafny.Set
				_539_v8 = interface{}(_compr_16).(_dafny.Set)
				if (_536_v7).Contains(_539_v8) {
					_coll16.Add(_539_v8)
				}
			}
			return _coll16.ToSet()
		}(), 9)
		(_nw85).ArraySet1((_dafny.SetOf(_529_v0)).Intersection(_dafny.SetOf(_dafny.SetOf(p1), _dafny.SetOf(p1), _529_v0, _529_v0, _529_v0)), 10)
		(_nw85).ArraySet1(_530_v1, 11)
		(_nw85).ArraySet1(_530_v1, 12)
		(_nw85).ArraySet1(_dafny.SetOf(_529_v0, _529_v0, _529_v0, _529_v0), 13)
		(_nw85).ArraySet1(_530_v1, 14)
		(_nw85).ArraySet1(_dafny.SetOf(_529_v0), 15)
		(_nw85).ArraySet1(_530_v1, 16)
		(_nw85).ArraySet1(func() _dafny.Set {
			var _coll17 = _dafny.NewBuilder()
			_ = _coll17
			for _iter19 := _dafny.Iterate((_537_v9).Elements()); ; {
				_compr_17, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _540_v10 _dafny.Set
				_540_v10 = interface{}(_compr_17).(_dafny.Set)
				if _dafny.Companion_Sequence_.Contains(_537_v9, _540_v10) {
					_coll17.Add(_540_v10)
				}
			}
			return _coll17.ToSet()
		}(), 17)
		(_nw85).ArraySet1(_530_v1, 18)
		(_nw85).ArraySet1(_530_v1, 19)
		(_nw85).ArraySet1((_530_v1).Union(_530_v1), 20)
		_538_v11 = _nw85
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_538_v11), 0))
		_ = _index80
		(_538_v11).ArraySet1(_dafny.SetOf(_529_v0), (_index80).Int())
		var _541_v12 D5
		_ = _541_v12
		_541_v12 = Companion_D5_.Create_DC11_(_dafny.SetOf(_529_v0, _529_v0))
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_538_v11), 0))
		_ = _index81
		(_538_v11).ArraySet1((_541_v12).Dtor_cf14(), (_index81).Int())
		var _source4 D2 = Companion_Default___.Fm14((_this).F0(), _531_v2, _535_v6, globalState)
		_ = _source4
		if _source4.Is_DC4() {
			var _542___mcc_h0 _dafny.Sequence = _source4.Get_().(D2_DC4).Cf4
			_ = _542___mcc_h0
			var _543___mcc_h1 bool = _source4.Get_().(D2_DC4).Cf5
			_ = _543___mcc_h1
			var _544_cf5 bool = _543___mcc_h1
			_ = _544_cf5
			var _545_cf4 _dafny.Sequence = _542___mcc_h0
			_ = _545_cf4
			var _546_v13 _dafny.Array
			_ = _546_v13
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw86
			_546_v13 = _nw86
			var _547_v14 _dafny.Sequence
			_ = _547_v14
			_547_v14 = _dafny.SeqOf(false, true, _544_cf5, _544_cf5)
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_546_v13), 0))
			_ = _index82
			(_546_v13).ArraySet1(!(!(_531_v2) || ((_547_v14).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_547_v14).Cardinality()))).Uint32()).(bool))), (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_546_v13), 0))
			_ = _index83
			(_546_v13).ArraySet1(_531_v2, (_index83).Int())
			var _548_v15 _dafny.Map
			_ = _548_v15
			_548_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v2, _545_cf4)
			_548_v15 = (_548_v15).Update((_dafny.SetOf(p1, (_this).F0())).IsProperSubsetOf(_529_v0), _545_cf4)
			r2 = _dafny.IntOfUint32((_545_cf4).Cardinality())
			_546_v13 = _546_v13
		} else if _source4.Is_DC5() {
			var _549_v16 _dafny.Map
			_ = _549_v16
			_549_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
			var _550_v17 _dafny.Sequence
			_ = _550_v17
			_550_v17 = _dafny.SeqOf(_531_v2, _531_v2, _531_v2)
			var _551_v18 _dafny.Map
			_ = _551_v18
			_551_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_550_v17).Cardinality()))
			var _552_v19 _dafny.Sequence
			_ = _552_v19
			_552_v19 = _dafny.SeqOf(Companion_Default___.Fm3(_534_v5, _535_v6, (_551_v18).Cardinality(), (_this).F0(), globalState))
			var _553_v20 _dafny.Map
			_ = _553_v20
			_553_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _535_v6)
			var _554_v21 _dafny.Array
			_ = _554_v21
			var _nwElement0_16 _dafny.Int = (func() _dafny.Int {
				if (_549_v16).Contains(_531_v2) {
					return (_549_v16).Get(_531_v2).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-553)
			})()
			_ = _nwElement0_16
			var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(26))
			_ = _nw87
			(_nw87).ArraySet1(_nwElement0_16, 0)
			(_nw87).ArraySet1(p1, 1)
			(_nw87).ArraySet1(p1, 2)
			(_nw87).ArraySet1((func() _dafny.Int {
				if (_this).Fm5(globalState) {
					return (_this).F0()
				}
				return (_this).F0()
			})(), 3)
			(_nw87).ArraySet1((_552_v19).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_552_v19).Cardinality()))).Uint32()).(_dafny.Int), 4)
			(_nw87).ArraySet1(p1, 5)
			(_nw87).ArraySet1((func() _dafny.Int {
				if false {
					return (_this).F0()
				}
				return _dafny.IntOfInt64(345)
			})(), 6)
			(_nw87).ArraySet1((_dafny.Zero).Minus(p1), 7)
			(_nw87).ArraySet1(_dafny.IntOfInt64(610), 8)
			(_nw87).ArraySet1((_dafny.Zero).Minus(p1), 9)
			(_nw87).ArraySet1(Companion_Default___.Fm3(_534_v5, (func() _dafny.CodePoint {
				if (_553_v20).Contains(p1) {
					return (_553_v20).Get(p1).(_dafny.CodePoint)
				}
				return _535_v6
			})(), (_this).F0(), (_this).F0(), globalState), 10)
			(_nw87).ArraySet1((_this).F0(), 11)
			(_nw87).ArraySet1((_this).F0(), 12)
			(_nw87).ArraySet1((_529_v0).Cardinality(), 13)
			(_nw87).ArraySet1((_this).F0(), 14)
			(_nw87).ArraySet1(_dafny.IntOfInt64(-41), 15)
			(_nw87).ArraySet1(p1, 16)
			(_nw87).ArraySet1(_dafny.IntOfUint32((_552_v19).Cardinality()), 17)
			(_nw87).ArraySet1((_dafny.Zero).Minus((_this).F0()), 18)
			(_nw87).ArraySet1(p1, 19)
			(_nw87).ArraySet1(_dafny.IntOfInt64(808), 20)
			(_nw87).ArraySet1(p1, 21)
			(_nw87).ArraySet1((_this).F0(), 22)
			(_nw87).ArraySet1(p1, 23)
			(_nw87).ArraySet1((_this).F0(), 24)
			(_nw87).ArraySet1(p1, 25)
			_554_v21 = _nw87
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))
			_ = _index84
			(_554_v21).ArraySet1(p1, (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))
			_ = _index85
			(_554_v21).ArraySet1((p1).Plus(Companion_Default___.SafeDivisionInt(p1, p1)), (_index85).Int())
			var _555_v22 _dafny.MultiSet
			_ = _555_v22
			_555_v22 = _dafny.MultiSetOf(_531_v2, false)
			var _556_v23 _dafny.MultiSet
			_ = _556_v23
			_556_v23 = _dafny.MultiSetOf(p1, p1, (_554_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))).Int()).(_dafny.Int))
			var _557_v24 *C0
			_ = _557_v24
			var _nw88 *C0 = New_C0_()
			_ = _nw88
			_nw88.Ctor__(_dafny.CodePoint('v'))
			_557_v24 = _nw88
			var _558_v25 _dafny.Map
			_ = _558_v25
			_558_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_557_v24, _531_v2)
			var _559_v26 _dafny.Map
			_ = _559_v26
			_559_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_558_v25).Contains(_557_v24) {
					return (_558_v25).Get(_557_v24).(bool)
				}
				return Companion_Default___.Fm2(_531_v2, p1, _531_v2, _531_v2, globalState)
			})(), true)
			var _560_v27 D0
			_ = _560_v27
			_560_v27 = Companion_D0_.Create_DC0_(_531_v2)
			var _561_v28 _dafny.Array
			_ = _561_v28
			var _nwElement0_17 bool = (_560_v27).Dtor_cf0()
			_ = _nwElement0_17
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(13))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_17, 0)
			(_nw89).ArraySet1(true, 1)
			(_nw89).ArraySet1(false, 2)
			(_nw89).ArraySet1(_531_v2, 3)
			(_nw89).ArraySet1((_this).Fm5(globalState), 4)
			(_nw89).ArraySet1(_531_v2, 5)
			(_nw89).ArraySet1(_531_v2, 6)
			(_nw89).ArraySet1(_531_v2, 7)
			(_nw89).ArraySet1(_531_v2, 8)
			(_nw89).ArraySet1(_531_v2, 9)
			(_nw89).ArraySet1(_531_v2, 10)
			(_nw89).ArraySet1(_531_v2, 11)
			(_nw89).ArraySet1(_531_v2, 12)
			_561_v28 = _nw89
			var _562_v29 _dafny.MultiSet
			_ = _562_v29
			_562_v29 = _dafny.MultiSetOf(_561_v28)
			var _563_v30 _dafny.Array
			_ = _563_v30
			var _nwElement0_18 bool = (_555_v22).IsDisjointFrom(_dafny.MultiSetFromSeq(_550_v17))
			_ = _nwElement0_18
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(26))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_18, 0)
			(_nw90).ArraySet1((func() bool {
				if _531_v2 {
					return _531_v2
				}
				return _531_v2
			})(), 1)
			(_nw90).ArraySet1((false) == (_531_v2), 2)
			(_nw90).ArraySet1(_531_v2, 3)
			(_nw90).ArraySet1((_550_v17).Select((Companion_Default___.SafeIndex((_554_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_550_v17).Cardinality()))).Uint32()).(bool), 4)
			(_nw90).ArraySet1(!(_531_v2) || (_531_v2), 5)
			(_nw90).ArraySet1(_531_v2, 6)
			(_nw90).ArraySet1(_531_v2, 7)
			(_nw90).ArraySet1((_550_v17).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_550_v17).Cardinality()))).Uint32()).(bool), 8)
			(_nw90).ArraySet1(_531_v2, 9)
			(_nw90).ArraySet1(_531_v2, 10)
			(_nw90).ArraySet1((_556_v23).IsDisjointFrom(_556_v23), 11)
			(_nw90).ArraySet1((_550_v17).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_550_v17).Cardinality()))).Uint32()).(bool), 12)
			(_nw90).ArraySet1(((_550_v17).Select((Companion_Default___.SafeIndex((_554_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_550_v17).Cardinality()))).Uint32()).(bool)) || ((func() bool {
				if (_559_v26).Contains(_531_v2) {
					return (_559_v26).Get(_531_v2).(bool)
				}
				return false
			})()), 13)
			(_nw90).ArraySet1((_531_v2) == (_531_v2), 14)
			(_nw90).ArraySet1(true, 15)
			(_nw90).ArraySet1((_dafny.MultiSetOf(_561_v28)).IsProperSubsetOf(_562_v29), 16)
			(_nw90).ArraySet1(_531_v2, 17)
			(_nw90).ArraySet1((_this).Fm5(globalState), 18)
			(_nw90).ArraySet1((_529_v0).IsProperSubsetOf(_529_v0), 19)
			(_nw90).ArraySet1(_531_v2, 20)
			(_nw90).ArraySet1(_531_v2, 21)
			(_nw90).ArraySet1(_531_v2, 22)
			(_nw90).ArraySet1(_531_v2, 23)
			(_nw90).ArraySet1(!((_this).Fm5(globalState)), 24)
			(_nw90).ArraySet1(_531_v2, 25)
			_563_v30 = _nw90
			var _564_v31 _dafny.Sequence
			_ = _564_v31
			_564_v31 = _dafny.UnicodeSeqOfUtf8Bytes("ykehlcy")
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_561_v28), 0))
			_ = _index86
			(_561_v28).ArraySet1(!_dafny.Companion_Sequence_.Contains(_564_v31, _535_v6), (_index86).Int())
			var _565_v32 _dafny.Sequence
			_ = _565_v32
			_565_v32 = _dafny.SeqOf(_554_v21, _554_v21, _554_v21)
			var _566_v33 D2
			_ = _566_v33
			_566_v33 = Companion_D2_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(382))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}(func(_567_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})), _531_v2)
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_561_v28), 0))
			_ = _index87
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))
			_ = _index88
			var _rhs83 bool = !(!(_531_v2) || (!_dafny.Companion_Sequence_.Contains(_565_v32, _554_v21)))
			_ = _rhs83
			var _rhs84 _dafny.Int = Companion_Default___.Fm3(_534_v5, (func() _dafny.CodePoint {
				if _531_v2 {
					return (_557_v24).F3()
				}
				return (_557_v24).F3()
			})(), p1, p1, globalState)
			_ = _rhs84
			var _rhs85 bool = (_dafny.IntOfUint32(((_566_v33).Dtor_cf4()).Cardinality())).Cmp((_this).F0()) <= 0
			_ = _rhs85
			var _rhs86 bool = _531_v2
			_ = _rhs86
			var _rhs87 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_550_v17).Cardinality()))
			_ = _rhs87
			var _lhs24 _dafny.Array = _561_v28
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_561_v28), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _554_v21
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_554_v21), 0))
			_ = _lhs27
			_531_v2 = _rhs83
			r2 = _rhs84
			(_lhs24).ArraySet1(_rhs85, (_lhs25).Int())
			_531_v2 = _rhs86
			(_lhs26).ArraySet1(_rhs87, (_lhs27).Int())
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw91
			_554_v21 = _nw91
			var _568_v34 *C0
			_ = _568_v34
			var _nw92 *C0 = New_C0_()
			_ = _nw92
			_nw92.Ctor__((_557_v24).F3())
			_568_v34 = _nw92
		} else if _source4.Is_DC3() {
			var _569___mcc_h2 _dafny.Sequence = _source4.Get_().(D2_DC3).Cf3
			_ = _569___mcc_h2
			var _570_cf3 _dafny.Sequence = _569___mcc_h2
			_ = _570_cf3
			var _571_v35 _dafny.Array
			_ = _571_v35
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_10
			var _nw93 _dafny.Array
			_ = _nw93
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw93 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Int = (func(_572_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_573_i1 _dafny.Int) _dafny.Int {
						return (_573_i1).Minus(_572_p1)
					}
				})(p1)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw93 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw93).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw93).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_571_v35 = _nw93
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
			_ = _index89
			(_571_v35).ArraySet1(Companion_Default___.SafeModuloInt(p1, p1), (_index89).Int())
			var _574_v36 _dafny.MultiSet
			_ = _574_v36
			_574_v36 = _dafny.MultiSetOf(_571_v35)
			var _575_v37 D6
			_ = _575_v37
			_575_v37 = Companion_D6_.Create_DC14_(_571_v35)
			var _576_v38 *C0
			_ = _576_v38
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__(_dafny.CodePoint('a'))
			_576_v38 = _nw94
			var _577_v39 _dafny.Map
			_ = _577_v39
			_577_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _576_v38)
			var _578_v40 _dafny.Map
			_ = _578_v40
			_578_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v2, p1)
			var _579_v41 _dafny.Sequence
			_ = _579_v41
			_579_v41 = _dafny.SeqOf(_576_v38, _576_v38)
			var _580_v42 _dafny.Array
			_ = _580_v42
			var _nwElement0_19 *C0 = _576_v38
			_ = _nwElement0_19
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(26))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_19, 0)
			(_nw95).ArraySet1(_576_v38, 1)
			(_nw95).ArraySet1(_576_v38, 2)
			(_nw95).ArraySet1(_576_v38, 3)
			(_nw95).ArraySet1((func() *C0 {
				if (_577_v39).Contains((func() _dafny.Int {
					if (_578_v40).Contains(_531_v2) {
						return (_578_v40).Get(_531_v2).(_dafny.Int)
					}
					return p1
				})()) {
					return (_577_v39).Get((func() _dafny.Int {
						if (_578_v40).Contains(_531_v2) {
							return (_578_v40).Get(_531_v2).(_dafny.Int)
						}
						return p1
					})()).(*C0)
				}
				return _576_v38
			})(), 4)
			(_nw95).ArraySet1((_579_v41).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_579_v41).Cardinality()))).Uint32()).(*C0), 5)
			(_nw95).ArraySet1(_576_v38, 6)
			(_nw95).ArraySet1(_576_v38, 7)
			(_nw95).ArraySet1(_576_v38, 8)
			(_nw95).ArraySet1(_576_v38, 9)
			(_nw95).ArraySet1(_576_v38, 10)
			(_nw95).ArraySet1(_576_v38, 11)
			(_nw95).ArraySet1(_576_v38, 12)
			(_nw95).ArraySet1(_576_v38, 13)
			(_nw95).ArraySet1(_576_v38, 14)
			(_nw95).ArraySet1(_576_v38, 15)
			(_nw95).ArraySet1(_576_v38, 16)
			(_nw95).ArraySet1(_576_v38, 17)
			(_nw95).ArraySet1(_576_v38, 18)
			(_nw95).ArraySet1(_576_v38, 19)
			(_nw95).ArraySet1(_576_v38, 20)
			(_nw95).ArraySet1(_576_v38, 21)
			(_nw95).ArraySet1(_576_v38, 22)
			(_nw95).ArraySet1(_576_v38, 23)
			(_nw95).ArraySet1(_576_v38, 24)
			(_nw95).ArraySet1(_576_v38, 25)
			_580_v42 = _nw95
			var _581_v43 D3
			_ = _581_v43
			_581_v43 = Companion_D3_.Create_DC9_((_this).F0(), _580_v42, _531_v2)
			var _pat_let_tv12 = _571_v35
			_ = _pat_let_tv12
			var _pat_let_tv13 = _531_v2
			_ = _pat_let_tv13
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
			_ = _index90
			var _rhs88 bool = (_574_v36).IsProperSubsetOf((_574_v36).Intersection(_dafny.MultiSetOf((func(_pat_let6_0 D6) D6 {
				return func(_582_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let7_0 _dafny.Array) D6 {
						return func(_583_dt__update_hcf22_h0 _dafny.Array) D6 {
							return Companion_D6_.Create_DC14_(_583_dt__update_hcf22_h0)
						}(_pat_let7_0)
					}(_pat_let_tv12)
				}(_pat_let6_0)
			}(_575_v37)).Dtor_cf22(), _571_v35, _571_v35, _571_v35)))
			_ = _rhs88
			var _rhs89 _dafny.Int = ((_this).F0()).Times((func(_pat_let8_0 D3) D3 {
				return func(_584_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let9_0 _dafny.Int) D3 {
						return func(_585_dt__update_hcf10_h0 _dafny.Int) D3 {
							return func(_pat_let10_0 bool) D3 {
								return func(_586_dt__update_hcf12_h0 bool) D3 {
									return Companion_D3_.Create_DC9_(_585_dt__update_hcf10_h0, (_584_dt__update__tmp_h1).Dtor_cf11(), _586_dt__update_hcf12_h0)
								}(_pat_let10_0)
							}(_pat_let_tv13)
						}(_pat_let9_0)
					}((_this).F0())
				}(_pat_let8_0)
			}(_581_v43)).Dtor_cf10())
			_ = _rhs89
			var _lhs28 _dafny.Array = _571_v35
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
			_ = _lhs29
			_531_v2 = _rhs88
			(_lhs28).ArraySet1(_rhs89, (_lhs29).Int())
			var _587_v44 _dafny.Sequence
			_ = _587_v44
			_587_v44 = _dafny.UnicodeSeqOfUtf8Bytes("arknopxms")
			var _588_v45 D2
			_ = _588_v45
			_588_v45 = Companion_D2_.Create_DC4_(_587_v44, _531_v2)
			var _589_v46 _dafny.Set
			_ = _589_v46
			_589_v46 = _dafny.SetOf(_588_v45, _588_v45, _588_v45)
			if (_dafny.SetOf(_588_v45)).IsDisjointFrom(_589_v46) {
				_531_v2 = _531_v2
				_531_v2 = (func() bool {
					if ((_this).F0()).Cmp((_this).F0()) != 0 {
						return _531_v2
					}
					return _531_v2
				})()
				_531_v2 = !(Companion_Default___.Fm2(_531_v2, ((_this).F0()).Times((_this).F0()), true, false, globalState))
				var _590_v47 _dafny.Map
				_ = _590_v47
				_590_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm18((_571_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))).Int()).(_dafny.Int), globalState)).Cardinality()), (_this).F0())
				var _591_v48 T0
				_ = _591_v48
				var _nw96 *C1 = New_C1_()
				_ = _nw96
				_nw96.Ctor__((func() _dafny.Int {
					if (_590_v47).Contains(p1) {
						return (_590_v47).Get(p1).(_dafny.Int)
					}
					return Companion_Default___.Fm3(_534_v5, _535_v6, (_this).F0(), p1, globalState)
				})())
				_591_v48 = _nw96
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
				_ = _index91
				var _rhs90 T0 = _591_v48
				_ = _rhs90
				var _rhs91 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_531_v2)).Intersection(_dafny.SetOf((_this).Fm6(_531_v2, _587_v44, _587_v44, (_591_v48).F0(), globalState), _531_v2)), _534_v5, _534_v5)).Cardinality())
				_ = _rhs91
				var _lhs30 _dafny.Array = _571_v35
				_ = _lhs30
				var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
				_ = _lhs31
				_591_v48 = _rhs90
				(_lhs30).ArraySet1(_rhs91, (_lhs31).Int())
				var _592_v49 _dafny.Array
				_ = _592_v49
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw97
				_592_v49 = _nw97
				_592_v49 = _592_v49
			} else {
				var _593_v50 *C0
				_ = _593_v50
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__((_576_v38).F3())
				_593_v50 = _nw98
				var _594_v51 _dafny.Array
				_ = _594_v51
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw99
				_594_v51 = _nw99
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_594_v51), 0))
				_ = _index92
				(_594_v51).ArraySet1(_531_v2, (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_594_v51), 0))
				_ = _index93
				(_594_v51).ArraySet1(_531_v2, (_index93).Int())
				_531_v2 = (_594_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_594_v51), 0))).Int()).(bool)
				_588_v45 = Companion_D2_.Create_DC4_(_587_v44, _531_v2)
				r1 = ((_571_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))).Int()).(_dafny.Int)).Times((_571_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))).Int()).(_dafny.Int))
			}
			var _595_v52 D0
			_ = _595_v52
			_595_v52 = Companion_D0_.Create_DC1_(_531_v2)
			var _596_v53 _dafny.Sequence
			_ = _596_v53
			_596_v53 = _dafny.SeqOf((_this).F0())
			var _pat_let_tv14 = _531_v2
			_ = _pat_let_tv14
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
			_ = _index94
			var _rhs92 _dafny.CodePoint = (_587_v44).Select((Companion_Default___.SafeIndex((_571_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_587_v44).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs92
			var _rhs93 bool = (func(_pat_let11_0 D0) D0 {
				return func(_597_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let12_0 bool) D0 {
						return func(_598_dt__update_hcf1_h0 bool) D0 {
							return Companion_D0_.Create_DC1_(_598_dt__update_hcf1_h0)
						}(_pat_let12_0)
					}(_pat_let_tv14)
				}(_pat_let11_0)
			}(_595_v52)).Dtor_cf1()
			_ = _rhs93
			var _rhs94 _dafny.Int = _dafny.IntOfUint32((_596_v53).Cardinality())
			_ = _rhs94
			var _lhs32 _dafny.Array = _571_v35
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
			_ = _lhs33
			_535_v6 = _rhs92
			_531_v2 = _rhs93
			(_lhs32).ArraySet1(_rhs94, (_lhs33).Int())
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_571_v35), 0))
			_ = _index95
			(_571_v35).ArraySet1(_dafny.IntOfInt64(193), (_index95).Int())
		} else {
			var _599___mcc_h3 D2 = _source4.Get_().(D2_DC6).Cf6
			_ = _599___mcc_h3
			var _600_cf6 D2 = _599___mcc_h3
			_ = _600_cf6
			var _601_v54 *C0
			_ = _601_v54
			var _nw100 *C0 = New_C0_()
			_ = _nw100
			_nw100.Ctor__(_535_v6)
			_601_v54 = _nw100
			var _602_v55 *C1
			_ = _602_v55
			var _nw101 *C1 = New_C1_()
			_ = _nw101
			_nw101.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _601_v54)).Cardinality())
			_602_v55 = _nw101
			r0 = (_dafny.Zero).Minus((_this).F0())
			var _603_v56 *C0
			_ = _603_v56
			var _nw102 *C0 = New_C0_()
			_ = _nw102
			_nw102.Ctor__((_601_v54).F3())
			_603_v56 = _nw102
			r1 = (_this).F0()
		}
		var _604_i2 _dafny.Int
		_ = _604_i2
		_604_i2 = _dafny.Zero
		{
			for _531_v2 {
				{
					if (_604_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_604_i2 = (_604_i2).Plus(_dafny.One)
					var _605_v57 _dafny.Array
					_ = _605_v57
					var _nw103 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
					_ = _nw103
					_605_v57 = _nw103
					var _606_v58 D3
					_ = _606_v58
					_606_v58 = Companion_D3_.Create_DC9_(Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ffshwyvs"))).Cardinality(), _dafny.IntOfInt64(364)), _605_v57, _531_v2)
					var _source5 D3 = _606_v58
					_ = _source5
					if _source5.Is_DC8() {
						var _607___mcc_h4 bool = _source5.Get_().(D3_DC8).Cf8
						_ = _607___mcc_h4
						var _608___mcc_h5 _dafny.Int = _source5.Get_().(D3_DC8).Cf9
						_ = _608___mcc_h5
						var _609_cf9 _dafny.Int = _608___mcc_h5
						_ = _609_cf9
						var _610_cf8 bool = _607___mcc_h4
						_ = _610_cf8
						var _611_v59 _dafny.Sequence
						_ = _611_v59
						_611_v59 = _dafny.UnicodeSeqOfUtf8Bytes("qpgrnxwqu")
						var _612_v60 bool
						_ = _612_v60
						var _out14 bool
						_ = _out14
						_out14 = Companion_Default___.M0((_this).Fm6(_531_v2, _611_v59, _611_v59, Companion_Default___.Fm3(_534_v5, _535_v6, (_534_v5).Cardinality(), (_this).F0(), globalState), globalState), p1, globalState)
						_612_v60 = _out14
						r1 = _609_cf9
						var _613_v61 *C1
						_ = _613_v61
						var _nw104 *C1 = New_C1_()
						_ = _nw104
						_nw104.Ctor__((_this).F0())
						_613_v61 = _nw104
						var _614_v62 _dafny.Map
						_ = _614_v62
						_614_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_612_v60, _612_v60)
						_614_v62 = (_614_v62).Update(_531_v2, _531_v2)
					} else if _source5.Is_DC9() {
						var _615___mcc_h6 _dafny.Int = _source5.Get_().(D3_DC9).Cf10
						_ = _615___mcc_h6
						var _616___mcc_h7 _dafny.Array = _source5.Get_().(D3_DC9).Cf11
						_ = _616___mcc_h7
						var _617___mcc_h8 bool = _source5.Get_().(D3_DC9).Cf12
						_ = _617___mcc_h8
						var _618_cf12 bool = _617___mcc_h8
						_ = _618_cf12
						var _619_cf11 _dafny.Array = _616___mcc_h7
						_ = _619_cf11
						var _620_cf10 _dafny.Int = _615___mcc_h6
						_ = _620_cf10
						var _621_v63 T1
						_ = _621_v63
						var _nw105 *C2 = New_C2_()
						_ = _nw105
						_nw105.Ctor__(p1)
						_621_v63 = _nw105
						var _622_v64 _dafny.Map
						_ = _622_v64
						_622_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F0()).Cmp(p1) <= 0, _621_v63)
						_622_v64 = (_622_v64).Update(false, (func() T1 {
							if (_622_v64).Contains((_621_v63).Fm5(globalState)) {
								return (_622_v64).Get((_621_v63).Fm5(globalState)).(T1)
							}
							return _621_v63
						})())
						var _623_v65 _dafny.Array
						_ = _623_v65
						var _len0_11 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_11
						var _nw106 _dafny.Array
						_ = _nw106
						if _len0_11.Cmp(_dafny.Zero) == 0 {
							_nw106 = _dafny.NewArray(_len0_11)
						} else {
							var _init11 func(_dafny.Int) _dafny.Set = (func(_624_v0 _dafny.Set) func(_dafny.Int) _dafny.Set {
								return func(_625_i3 _dafny.Int) _dafny.Set {
									return _624_v0
								}
							})(_529_v0)
							_ = _init11
							var _element0_11 = _init11(_dafny.Zero)
							_ = _element0_11
							_nw106 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
							(_nw106).ArraySet1(_element0_11, 0)
							var _nativeLen0_11 = (_len0_11).Int()
							_ = _nativeLen0_11
							for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
								(_nw106).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
							}
						}
						_623_v65 = _nw106
						_623_v65 = _623_v65
						var _626_v66 _dafny.Array
						_ = _626_v66
						var _nw107 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
						_ = _nw107
						_626_v66 = _nw107
						var _627_v67 _dafny.MultiSet
						_ = _627_v67
						_627_v67 = _dafny.MultiSetOf((_621_v63).F0(), _620_cf10)
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_626_v66), 0))
						_ = _index96
						(_626_v66).ArraySet1(((func() _dafny.Int {
							if (_627_v67).Contains(_620_cf10) {
								return (_627_v67).Multiplicity(_620_cf10)
							}
							return (_this).F0()
						})()).Cmp((_this).F0()) < 0, (_index96).Int())
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_626_v66), 0))
						_ = _index97
						(_626_v66).ArraySet1(true, (_index97).Int())
						r2 = (_dafny.Zero).Minus((p1).Times(_620_cf10))
					} else {
						var _628___mcc_h9 _dafny.Sequence = _source5.Get_().(D3_DC7).Cf7
						_ = _628___mcc_h9
						var _629_cf7 _dafny.Sequence = _628___mcc_h9
						_ = _629_cf7
						var _630_v68 *C0
						_ = _630_v68
						var _nw108 *C0 = New_C0_()
						_ = _nw108
						_nw108.Ctor__(_535_v6)
						_630_v68 = _nw108
						_531_v2 = _531_v2
						r0 = (p1).Plus((_this).F0())
						var _631_v69 _dafny.Array
						_ = _631_v69
						var _len0_12 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_12
						var _nw109 _dafny.Array
						_ = _nw109
						if _len0_12.Cmp(_dafny.Zero) == 0 {
							_nw109 = _dafny.NewArray(_len0_12)
						} else {
							var _init12 func(_dafny.Int) _dafny.Int = (func(_632_cf7 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
								return func(_633_i4 _dafny.Int) _dafny.Int {
									return (_633_i4).Times(_dafny.IntOfUint32((_632_cf7).Cardinality()))
								}
							})(_629_cf7)
							_ = _init12
							var _element0_12 = _init12(_dafny.Zero)
							_ = _element0_12
							_nw109 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
							(_nw109).ArraySet1(_element0_12, 0)
							var _nativeLen0_12 = (_len0_12).Int()
							_ = _nativeLen0_12
							for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
								(_nw109).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
							}
						}
						_631_v69 = _nw109
						var _634_v70 _dafny.MultiSet
						_ = _634_v70
						_634_v70 = _dafny.MultiSetOf(_631_v69, _631_v69, _631_v69, _631_v69, _631_v69)
						_634_v70 = _dafny.MultiSetOf(_631_v69, _631_v69)
					}
					var _635_v71 _dafny.Array
					_ = _635_v71
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_13
					var _nw110 _dafny.Array
					_ = _nw110
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw110 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) _dafny.Int = (func(_636_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_637_i5 _dafny.Int) _dafny.Int {
								return (_637_i5).Times(_636_p1)
							}
						})(p1)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw110 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw110).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw110).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_635_v71 = _nw110
					var _638_v72 _dafny.Map
					_ = _638_v72
					_638_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _635_v71)
					_638_v72 = _638_v72
					var _639_v73 _dafny.Map
					_ = _639_v73
					_639_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), p1)
					var _640_v74 _dafny.Map
					_ = _640_v74
					_640_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_535_v6, p1)
					var _641_v75 _dafny.Array
					_ = _641_v75
					var _nwElement0_20 bool = _531_v2
					_ = _nwElement0_20
					var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(10))
					_ = _nw111
					(_nw111).ArraySet1(_nwElement0_20, 0)
					(_nw111).ArraySet1(_531_v2, 1)
					(_nw111).ArraySet1(!((_639_v73).Update((_640_v74).Cardinality(), (_this).F0())).Equals(_639_v73), 2)
					(_nw111).ArraySet1(true, 3)
					(_nw111).ArraySet1(_531_v2, 4)
					(_nw111).ArraySet1(true, 5)
					(_nw111).ArraySet1(true, 6)
					(_nw111).ArraySet1(false, 7)
					(_nw111).ArraySet1(_531_v2, 8)
					(_nw111).ArraySet1(_531_v2, 9)
					_641_v75 = _nw111
					var _642_v76 _dafny.Sequence
					_ = _642_v76
					_642_v76 = _dafny.SeqOf((_this).Fm13((_this).F0(), false, globalState), _531_v2)
					var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_641_v75), 0))
					_ = _index98
					(_641_v75).ArraySet1((_642_v76).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_639_v73).Cardinality()), _dafny.IntOfUint32((_642_v76).Cardinality()))).Uint32()).(bool), (_index98).Int())
					var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_641_v75), 0))
					_ = _index99
					(_641_v75).ArraySet1((_531_v2) == (_531_v2), (_index99).Int())
					var _643_v77 _dafny.MultiSet
					_ = _643_v77
					_643_v77 = _dafny.MultiSetOf(_531_v2, true)
					var _644_v78 _dafny.Sequence
					_ = _644_v78
					_644_v78 = _dafny.UnicodeSeqOfUtf8Bytes("dvejd")
					var _645_v79 _dafny.MultiSet
					_ = _645_v79
					_645_v79 = _dafny.MultiSetOf(_dafny.IntOfInt64(204), (_this).F0())
					var _646_v80 _dafny.MultiSet
					_ = _646_v80
					_646_v80 = _dafny.MultiSetOf(_643_v77, _643_v77)
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_641_v75), 0))
					_ = _index100
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_641_v75), 0))
					_ = _index101
					var _rhs95 bool = (((_643_v77).Union(_643_v77)).Cardinality()).Cmp(p1) != 0
					_ = _rhs95
					var _rhs96 _dafny.Int = p1
					_ = _rhs96
					var _rhs97 D3 = _606_v58
					_ = _rhs97
					var _rhs98 _dafny.Int = (((_534_v5).Union(_dafny.SetOf(_531_v2))).Union(_534_v5)).Cardinality()
					_ = _rhs98
					var _rhs99 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_644_v78, Companion_Default___.Fm19(_645_v79, (_this).F0(), (_this).F0(), (_this).F0(), globalState))).Cardinality())).Cmp(Companion_Default___.Fm3(_534_v5, _535_v6, (_this).F0(), (_646_v80).Cardinality(), globalState)) == 0
					_ = _rhs99
					var _lhs34 _dafny.Array = _641_v75
					_ = _lhs34
					var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_641_v75), 0))
					_ = _lhs35
					var _lhs36 _dafny.Array = _641_v75
					_ = _lhs36
					var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_641_v75), 0))
					_ = _lhs37
					(_lhs34).ArraySet1(_rhs95, (_lhs35).Int())
					r0 = _rhs96
					_606_v58 = _rhs97
					r0 = _rhs98
					(_lhs36).ArraySet1(_rhs99, (_lhs37).Int())
					var _647_v81 _dafny.Array
					_ = _647_v81
					var _nwElement0_21 _dafny.CodePoint = _535_v6
					_ = _nwElement0_21
					var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(5))
					_ = _nw112
					(_nw112).ArraySet1CodePoint(_nwElement0_21, 0)
					(_nw112).ArraySet1CodePoint((func() _dafny.CodePoint {
						if _531_v2 {
							return _535_v6
						}
						return _535_v6
					})(), 1)
					(_nw112).ArraySet1CodePoint(_535_v6, 2)
					(_nw112).ArraySet1CodePoint(_535_v6, 3)
					(_nw112).ArraySet1CodePoint(_535_v6, 4)
					_647_v81 = _nw112
					var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_647_v81), 0))
					_ = _index102
					(_647_v81).ArraySet1CodePoint(_dafny.CodePoint('b'), (_index102).Int())
					var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_647_v81), 0))
					_ = _index103
					(_647_v81).ArraySet1CodePoint(_535_v6, (_index103).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _648_i6 _dafny.Int
		_ = _648_i6
		_648_i6 = _dafny.Zero
		{
			for (_this).Fm13((_dafny.Zero).Minus((Companion_Default___.Fm3(_534_v5, _535_v6, p1, p1, globalState)).Minus(p1)), true, globalState) {
				{
					if (_648_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_648_i6 = (_648_i6).Plus(_dafny.One)
					_531_v2 = !(_531_v2)
					var _649_v82 _dafny.Map
					_ = _649_v82
					_649_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v2, _534_v5)
					_531_v2 = (Companion_Default___.SafeDivisionInt((_this).F0(), (Companion_D5_.Create_DC13_(_dafny.IntOfInt64(678), (func() _dafny.Set {
						if (_649_v82).Contains(_531_v2) {
							return (_649_v82).Get(_531_v2).(_dafny.Set)
						}
						return _534_v5
					})(), _535_v6, (_this).F0(), false)).Dtor_cf17())).Cmp((_this).F0()) <= 0
					_531_v2 = _531_v2
					var _650_v83 _dafny.Sequence
					_ = _650_v83
					_650_v83 = _dafny.SeqOf(_531_v2, _531_v2, true, !(_531_v2), true)
					var _651_v84 _dafny.Sequence
					_ = _651_v84
					_651_v84 = _dafny.UnicodeSeqOfUtf8Bytes("oiil")
					var _652_v85 _dafny.Sequence
					_ = _652_v85
					_652_v85 = _dafny.SeqOf(_dafny.IntOfUint32((_651_v84).Cardinality()), (_dafny.Zero).Minus(p1))
					var _653_v86 _dafny.MultiSet
					_ = _653_v86
					_653_v86 = _dafny.MultiSetOf(p1, p1)
					var _source6 D5 = Companion_Default___.Fm23((_650_v83).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F0()), _dafny.IntOfUint32((_650_v83).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus(((_534_v5).Cardinality()).Times(p1)), (_dafny.Zero).Minus((_652_v85).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_652_v85).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm19(_653_v86, (_dafny.Zero).Minus((_this).F0()), (_653_v86).Cardinality(), _dafny.IntOfUint32((_651_v84).Cardinality()), globalState), _651_v84), globalState)
					_ = _source6
					if _source6.Is_DC12() {
						var _654___mcc_h10 _dafny.Int = _source6.Get_().(D5_DC12).Cf15
						_ = _654___mcc_h10
						var _655___mcc_h11 _dafny.Int = _source6.Get_().(D5_DC12).Cf16
						_ = _655___mcc_h11
						var _656_cf16 _dafny.Int = _655___mcc_h11
						_ = _656_cf16
						var _657_cf15 _dafny.Int = _654___mcc_h10
						_ = _657_cf15
						var _658_v88 *C2
						_ = _658_v88
						var _nw113 *C2 = New_C2_()
						_ = _nw113
						_nw113.Ctor__(((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qvaepprp")).Cardinality()), _dafny.IntOfInt64(913))).Intersection(func() _dafny.Set {
							var _coll18 = _dafny.NewBuilder()
							_ = _coll18
							for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(908), _dafny.IntOfInt64(602))); ; {
								_compr_18, _ok20 := _iter20()
								if !_ok20 {
									break
								}
								var _659_v87 _dafny.Int
								_659_v87 = interface{}(_compr_18).(_dafny.Int)
								if ((_dafny.IntOfInt64(908)).Cmp(_659_v87) <= 0) && ((_659_v87).Cmp(_dafny.IntOfInt64(602)) < 0) {
									_coll18.Add((_659_v87).Minus((_dafny.Zero).Minus(p1)))
								}
							}
							return _coll18.ToSet()
						}())).Cardinality())
						_658_v88 = _nw113
						r0 = (_652_v85).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_652_v85).Cardinality()))).Uint32()).(_dafny.Int)
						_531_v2 = _531_v2
						_531_v2 = _531_v2
					} else if _source6.Is_DC13() {
						var _660___mcc_h12 _dafny.Int = _source6.Get_().(D5_DC13).Cf17
						_ = _660___mcc_h12
						var _661___mcc_h13 _dafny.Set = _source6.Get_().(D5_DC13).Cf18
						_ = _661___mcc_h13
						var _662___mcc_h14 _dafny.CodePoint = _source6.Get_().(D5_DC13).Cf19
						_ = _662___mcc_h14
						var _663___mcc_h15 _dafny.Int = _source6.Get_().(D5_DC13).Cf20
						_ = _663___mcc_h15
						var _664___mcc_h16 bool = _source6.Get_().(D5_DC13).Cf21
						_ = _664___mcc_h16
						var _665_cf21 bool = _664___mcc_h16
						_ = _665_cf21
						var _666_cf20 _dafny.Int = _663___mcc_h15
						_ = _666_cf20
						var _667_cf19 _dafny.CodePoint = _662___mcc_h14
						_ = _667_cf19
						var _668_cf18 _dafny.Set = _661___mcc_h13
						_ = _668_cf18
						var _669_cf17 _dafny.Int = _660___mcc_h12
						_ = _669_cf17
						r0 = (_this).F0()
						_666_cf20 = _dafny.IntOfInt64(-224)
						_665_cf21 = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(165))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}((func(_670_v86 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
							return func(_671_i7 _dafny.Int) _dafny.Int {
								return (_670_v86).Cardinality()
							}
						})(_653_v86)))).Cardinality())).Cmp((_this).F0()) > 0
						_665_cf21 = _531_v2
					} else {
						var _672___mcc_h17 _dafny.Set = _source6.Get_().(D5_DC11).Cf14
						_ = _672___mcc_h17
						var _673_cf14 _dafny.Set = _672___mcc_h17
						_ = _673_cf14
						r1 = _dafny.IntOfInt64(-301)
						var _674_v89 D3
						_ = _674_v89
						_674_v89 = Companion_D3_.Create_DC8_(_531_v2, p1)
						var _675_v90 _dafny.Map
						_ = _675_v90
						_675_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v2, (_674_v89).Dtor_cf9())
						var _676_v91 _dafny.MultiSet
						_ = _676_v91
						_676_v91 = _dafny.MultiSetOf(_675_v90, _675_v90, _675_v90)
						var _677_v92 _dafny.MultiSet
						_ = _677_v92
						_677_v92 = _dafny.MultiSetOf(_534_v5)
						var _rhs100 bool = !(_dafny.MultiSetOf(_534_v5, _534_v5)).Equals(_677_v92)
						_ = _rhs100
						var _rhs101 bool = _531_v2
						_ = _rhs101
						var _rhs102 bool = !((_dafny.IntOfInt64(252)).Cmp(p1) >= 0)
						_ = _rhs102
						var _rhs103 _dafny.MultiSet = ((_676_v91).Intersection(_dafny.MultiSetOf(_675_v90))).Difference(_676_v91)
						_ = _rhs103
						var _rhs104 _dafny.Int = (Companion_Default___.SafeModuloInt((_this).F0(), (_this).F0())).Times((_this).F0())
						_ = _rhs104
						_531_v2 = _rhs100
						_531_v2 = _rhs101
						_531_v2 = _rhs102
						_676_v91 = _rhs103
						r2 = _rhs104
						_531_v2 = _531_v2
						var _678_v93 _dafny.Array
						_ = _678_v93
						var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
						_ = _nw114
						_678_v93 = _nw114
						var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_678_v93), 0))
						_ = _index104
						(_678_v93).ArraySet1(p1, (_index104).Int())
						var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_678_v93), 0))
						_ = _index105
						(_678_v93).ArraySet1((_dafny.Zero).Minus((_this).F0()), (_index105).Int())
					}
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _679_i8 _dafny.Int
		_ = _679_i8
		_679_i8 = _dafny.Zero
		{
			for (Companion_Default___.Fm2(true, p1, _531_v2, _531_v2, globalState)) || (_531_v2) {
				{
					if (_679_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_679_i8 = (_679_i8).Plus(_dafny.One)
					var _680_v94 _dafny.Array
					_ = _680_v94
					var _len0_14 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_14
					var _nw115 _dafny.Array
					_ = _nw115
					if _len0_14.Cmp(_dafny.Zero) == 0 {
						_nw115 = _dafny.NewArray(_len0_14)
					} else {
						var _init14 func(_dafny.Int) D7 = (func(_681_v2 bool) func(_dafny.Int) D7 {
							return func(_682_i9 _dafny.Int) D7 {
								return (func() D7 {
									if _681_v2 {
										return Companion_D7_.Create_DC17_(_681_v2, _681_v2, _681_v2)
									}
									return Companion_D7_.Create_DC17_(_681_v2, _681_v2, _681_v2)
								})()
							}
						})(_531_v2)
						_ = _init14
						var _element0_14 = _init14(_dafny.Zero)
						_ = _element0_14
						_nw115 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
						(_nw115).ArraySet1(_element0_14, 0)
						var _nativeLen0_14 = (_len0_14).Int()
						_ = _nativeLen0_14
						for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
							(_nw115).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
						}
					}
					_680_v94 = _nw115
					var _nw116 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(13))
					_ = _nw116
					_680_v94 = _nw116
					var _683_v95 _dafny.Array
					_ = _683_v95
					var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
					_ = _nw117
					_683_v95 = _nw117
					var _684_v96 _dafny.Array
					_ = _684_v96
					var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
					_ = _nw118
					_684_v96 = _nw118
					var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_683_v95), 0))
					_ = _index106
					(_683_v95).ArraySet1(_684_v96, (_index106).Int())
					var _685_v97 *C1
					_ = _685_v97
					var _nw119 *C1 = New_C1_()
					_ = _nw119
					_nw119.Ctor__(p1)
					_685_v97 = _nw119
					var _686_v98 _dafny.Map
					_ = _686_v98
					_686_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_685_v97, _684_v96)
					var _687_v99 _dafny.Array
					_ = _687_v99
					_687_v99 = (func() _dafny.Array {
						if (_686_v98).Contains(_685_v97) {
							return (_686_v98).Get(_685_v97).(_dafny.Array)
						}
						return _684_v96
					})()
					var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_683_v95), 0))
					_ = _index107
					(_683_v95).ArraySet1((_687_v99), (_index107).Int())
					r0 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-892))
					var _688_v100 _dafny.Map
					_ = _688_v100
					_688_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_534_v5, _535_v6, (_this).F0(), p1, globalState), _dafny.SeqOf(_531_v2))
					var _689_v101 _dafny.Sequence
					_ = _689_v101
					_689_v101 = _dafny.SeqOf(_531_v2, _531_v2, _531_v2)
					_688_v100 = (_688_v100).Update(p1, _689_v101)
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		_531_v2 = _531_v2
		r0 = (_this).F0()
		var _690_v102 _dafny.Sequence
		_ = _690_v102
		_690_v102 = _dafny.SeqOf(p1)
		r1 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_690_v102).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_690_v102).Cardinality()))).Uint32()).(_dafny.Int)))
		r2 = (_533_v4).Cardinality()
		return r0, r1, r2
	}
}
func (_this *C3) M6(p0 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _691_v0 T0
		_ = _691_v0
		var _nw120 *C1 = New_C1_()
		_ = _nw120
		_nw120.Ctor__(p0)
		_691_v0 = _nw120
		var _692_v1 _dafny.Map
		_ = _692_v1
		_692_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v0, (_691_v0).F0())
		_692_v1 = (_692_v1).Update(_691_v0, (_this).F0())
		var _693_v2 bool
		_ = _693_v2
		_693_v2 = true
		var _694_v3 _dafny.Sequence
		_ = _694_v3
		_694_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ejn")
		var _695_v4 D2
		_ = _695_v4
		_695_v4 = Companion_D2_.Create_DC4_(_694_v3, (_691_v0).Fm5(globalState))
		var _696_v5 D3
		_ = _696_v5
		_696_v5 = Companion_D3_.Create_DC8_(_693_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fxkw"), (_695_v4).Dtor_cf4())).Cardinality()))
		var _source7 D3 = _696_v5
		_ = _source7
		if _source7.Is_DC8() {
			var _697___mcc_h0 bool = _source7.Get_().(D3_DC8).Cf8
			_ = _697___mcc_h0
			var _698___mcc_h1 _dafny.Int = _source7.Get_().(D3_DC8).Cf9
			_ = _698___mcc_h1
			var _699_cf9 _dafny.Int = _698___mcc_h1
			_ = _699_cf9
			var _700_cf8 bool = _697___mcc_h0
			_ = _700_cf8
			r1 = _dafny.IntOfInt64(-144)
			var _701_v6 *C1
			_ = _701_v6
			var _nw121 *C1 = New_C1_()
			_ = _nw121
			_nw121.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(276), _dafny.IntOfUint32((_694_v3).Cardinality())))
			_701_v6 = _nw121
			var _702_v7 _dafny.CodePoint
			_ = _702_v7
			_702_v7 = _dafny.CodePoint('s')
			var _703_v8 *C0
			_ = _703_v8
			var _nw122 *C0 = New_C0_()
			_ = _nw122
			_nw122.Ctor__(_702_v7)
			_703_v8 = _nw122
			var _704_v9 _dafny.MultiSet
			_ = _704_v9
			_704_v9 = _dafny.MultiSetOf(_693_v2)
			var _705_v10 _dafny.Sequence
			_ = _705_v10
			_705_v10 = _dafny.SeqOf(false)
			var _706_v11 _dafny.Sequence
			_ = _706_v11
			_706_v11 = _dafny.SeqOf(_704_v9, _dafny.MultiSetFromSeq(_705_v10))
			var _707_v12 *C2
			_ = _707_v12
			var _nw123 *C2 = New_C2_()
			_ = _nw123
			_nw123.Ctor__(Companion_Default___.SafeModuloInt(_699_cf9, ((_706_v11).Select((Companion_Default___.SafeIndex(_699_cf9, _dafny.IntOfUint32((_706_v11).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()))
			_707_v12 = _nw123
		} else if _source7.Is_DC9() {
			var _708___mcc_h2 _dafny.Int = _source7.Get_().(D3_DC9).Cf10
			_ = _708___mcc_h2
			var _709___mcc_h3 _dafny.Array = _source7.Get_().(D3_DC9).Cf11
			_ = _709___mcc_h3
			var _710___mcc_h4 bool = _source7.Get_().(D3_DC9).Cf12
			_ = _710___mcc_h4
			var _711_cf12 bool = _710___mcc_h4
			_ = _711_cf12
			var _712_cf11 _dafny.Array = _709___mcc_h3
			_ = _712_cf11
			var _713_cf10 _dafny.Int = _708___mcc_h2
			_ = _713_cf10
			var _714_v13 _dafny.Sequence
			_ = _714_v13
			_714_v13 = _dafny.SeqOf(_dafny.SetOf((_691_v0).F0(), (_691_v0).F0()))
			var _715_v14 _dafny.Set
			_ = _715_v14
			_715_v14 = _dafny.SetOf((_691_v0).F0())
			var _716_v15 _dafny.MultiSet
			_ = _716_v15
			_716_v15 = _dafny.MultiSetOf((_dafny.SetOf(p0)).Union((_714_v13).Select((Companion_Default___.SafeIndex((_691_v0).F0(), _dafny.IntOfUint32((_714_v13).Cardinality()))).Uint32()).(_dafny.Set)), _715_v14, _715_v14, (_715_v14).Union(_715_v14))
			_716_v15 = _716_v15
			var _717_v16 _dafny.Array
			_ = _717_v16
			var _nw124 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw124
			_717_v16 = _nw124
			_717_v16 = _717_v16
			var _718_v17 D6
			_ = _718_v17
			_718_v17 = Companion_D6_.Create_DC15_(_dafny.IntOfInt64(934), _711_cf12, (_691_v0).F0())
			_718_v17 = _718_v17
			_694_v3 = _694_v3
		} else {
			var _719___mcc_h5 _dafny.Sequence = _source7.Get_().(D3_DC7).Cf7
			_ = _719___mcc_h5
			var _720_cf7 _dafny.Sequence = _719___mcc_h5
			_ = _720_cf7
			r1 = (func() _dafny.Int {
				if true {
					return Companion_Default___.SafeDivisionInt(p0, p0)
				}
				return (_691_v0).F0()
			})()
			_693_v2 = _693_v2
			var _721_v18 _dafny.CodePoint
			_ = _721_v18
			_721_v18 = _dafny.CodePoint('d')
			_721_v18 = Companion_Default___.Fm17((_691_v0).F0(), globalState)
			var _722_v19 _dafny.Array
			_ = _722_v19
			var _nwElement0_22 bool = (_693_v2) && (_693_v2)
			_ = _nwElement0_22
			var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(11))
			_ = _nw125
			(_nw125).ArraySet1(_nwElement0_22, 0)
			(_nw125).ArraySet1(_693_v2, 1)
			(_nw125).ArraySet1(_693_v2, 2)
			(_nw125).ArraySet1(_693_v2, 3)
			(_nw125).ArraySet1(_693_v2, 4)
			(_nw125).ArraySet1(false, 5)
			(_nw125).ArraySet1(true, 6)
			(_nw125).ArraySet1(_693_v2, 7)
			(_nw125).ArraySet1(_693_v2, 8)
			(_nw125).ArraySet1(_693_v2, 9)
			(_nw125).ArraySet1(_693_v2, 10)
			_722_v19 = _nw125
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_722_v19), 0))
			_ = _index108
			(_722_v19).ArraySet1((_this).Fm5(globalState), (_index108).Int())
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_722_v19), 0))
			_ = _index109
			(_722_v19).ArraySet1(_693_v2, (_index109).Int())
			var _723_v20 _dafny.Map
			_ = _723_v20
			_723_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v2, _693_v2)
			var _724_v21 _dafny.Sequence
			_ = _724_v21
			_724_v21 = _dafny.SeqOf(_723_v20, _723_v20, _723_v20)
			var _725_v22 _dafny.Set
			_ = _725_v22
			_725_v22 = _dafny.SetOf(_693_v2, _693_v2, (_691_v0).Fm5(globalState))
			var _726_v23 D2
			_ = _726_v23
			_726_v23 = Companion_D2_.Create_DC5_()
			var _727_v24 D2
			_ = _727_v24
			_727_v24 = Companion_D2_.Create_DC6_(_726_v23)
			var _728_v25 _dafny.MultiSet
			_ = _728_v25
			_728_v25 = _dafny.MultiSetOf(_727_v24)
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_722_v19), 0))
			_ = _index110
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_722_v19), 0))
			_ = _index111
			var _rhs105 _dafny.Int = (_this).F0()
			_ = _rhs105
			var _rhs106 _dafny.CodePoint = Companion_Default___.Fm17(Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(p0)), globalState)
			_ = _rhs106
			var _rhs107 bool = (func() bool {
				if (func() bool {
					if _693_v2 {
						return _693_v2
					}
					return _693_v2
				})() {
					return _693_v2
				}
				return _693_v2
			})()
			_ = _rhs107
			var _rhs108 bool = ((_728_v25).Update(_727_v24, Companion_Default___.Abs(_dafny.IntOfUint32((_694_v3).Cardinality())))).IsProperSubsetOf(Companion_Default___.Fm24(_693_v2, _693_v2, ((_724_v21).Select((Companion_Default___.SafeIndex((_691_v0).F0(), _dafny.IntOfUint32((_724_v21).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _725_v22, globalState))
			_ = _rhs108
			var _lhs38 _dafny.Array = _722_v19
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_722_v19), 0))
			_ = _lhs39
			var _lhs40 _dafny.Array = _722_v19
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_722_v19), 0))
			_ = _lhs41
			r1 = _rhs105
			_721_v18 = _rhs106
			(_lhs38).ArraySet1(_rhs107, (_lhs39).Int())
			(_lhs40).ArraySet1(_rhs108, (_lhs41).Int())
		}
		var _729_v26 _dafny.Array
		_ = _729_v26
		var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw126
		_729_v26 = _nw126
		var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_729_v26), 0))
		_ = _index112
		(_729_v26).ArraySet1(p0, (_index112).Int())
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_729_v26), 0))
		_ = _index113
		(_729_v26).ArraySet1((p0).Plus((_this).F0()), (_index113).Int())
		var _730_v27 _dafny.Map
		_ = _730_v27
		_730_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_729_v26, _729_v26)
		_730_v27 = (_730_v27).Update(_729_v26, _729_v26)
		var _731_v28 _dafny.Array
		_ = _731_v28
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_15
		var _nw127 _dafny.Array
		_ = _nw127
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw127 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = func(_732_i0 _dafny.Int) bool {
				return true
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw127 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw127).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw127).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_731_v28 = _nw127
		_731_v28 = _731_v28
		var _733_v29 _dafny.Set
		_ = _733_v29
		_733_v29 = _dafny.SetOf(_693_v2, _693_v2)
		r0 = (_733_v29).IsSubsetOf(_733_v29)
		r0 = _693_v2
		var _734_v30 _dafny.Set
		_ = _734_v30
		_734_v30 = _dafny.SetOf(_729_v26, _729_v26)
		r1 = (_734_v30).Cardinality()
		return r0, r1
	}
}
func (_this *C3) M7(p0 _dafny.Sequence, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) {
	{
		var _735_v0 _dafny.Int
		_ = _735_v0
		_735_v0 = _dafny.IntOfInt64(-360)
		_735_v0 = Companion_Default___.SafeModuloInt(_735_v0, p3)
		var _736_v1 _dafny.Map
		_ = _736_v1
		_736_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
		if !((func() _dafny.Map {
			if true {
				return _736_v1
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
		})()).Equals(_736_v1) {
			var _737_v2 _dafny.Map
			_ = _737_v2
			_737_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			_735_v0 = (((_737_v2).Merge(_737_v2)).Cardinality()).Times((_this).F0())
			var _738_v3 bool
			_ = _738_v3
			_738_v3 = false
			var _739_v4 _dafny.Sequence
			_ = _739_v4
			_739_v4 = _dafny.SeqOf(p2, p2, p2, p2, p2)
			_738_v3 = _dafny.Companion_Sequence_.Equal(_739_v4, _dafny.Companion_Sequence_.Concatenate(_739_v4, _739_v4))
			var _740_v5 _dafny.Set
			_ = _740_v5
			_740_v5 = _dafny.SetOf(p3)
			_738_v3 = !((_738_v3) || ((_740_v5).IsSubsetOf(_740_v5)))
			var _741_v6 _dafny.CodePoint
			_ = _741_v6
			_741_v6 = _dafny.CodePoint('p')
			var _742_v7 _dafny.Sequence
			_ = _742_v7
			_742_v7 = _dafny.SeqOf(_741_v6, _741_v6)
			var _743_v8 T1
			_ = _743_v8
			var _nw128 *C2 = New_C2_()
			_ = _nw128
			_nw128.Ctor__((_740_v5).Cardinality())
			_743_v8 = _nw128
			var _744_v9 _dafny.Array
			_ = _744_v9
			var _nwElement0_23 T1 = _743_v8
			_ = _nwElement0_23
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(20))
			_ = _nw129
			(_nw129).ArraySet1(_nwElement0_23, 0)
			(_nw129).ArraySet1(_743_v8, 1)
			(_nw129).ArraySet1((func() T1 {
				if _738_v3 {
					return _743_v8
				}
				return _743_v8
			})(), 2)
			(_nw129).ArraySet1(_743_v8, 3)
			(_nw129).ArraySet1(_743_v8, 4)
			(_nw129).ArraySet1(_743_v8, 5)
			(_nw129).ArraySet1(_743_v8, 6)
			(_nw129).ArraySet1(_743_v8, 7)
			(_nw129).ArraySet1(_743_v8, 8)
			(_nw129).ArraySet1(_743_v8, 9)
			(_nw129).ArraySet1(_743_v8, 10)
			(_nw129).ArraySet1(_743_v8, 11)
			(_nw129).ArraySet1(_743_v8, 12)
			(_nw129).ArraySet1(_743_v8, 13)
			(_nw129).ArraySet1(_743_v8, 14)
			(_nw129).ArraySet1(_743_v8, 15)
			(_nw129).ArraySet1(_743_v8, 16)
			(_nw129).ArraySet1(_743_v8, 17)
			(_nw129).ArraySet1(_743_v8, 18)
			(_nw129).ArraySet1(_743_v8, 19)
			_744_v9 = _nw129
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_744_v9), 0))
			_ = _index114
			(_744_v9).ArraySet1(_743_v8, (_index114).Int())
			var _745_v10 _dafny.Map
			_ = _745_v10
			_745_v10 = _737_v2
			var _746_v11 _dafny.Sequence
			_ = _746_v11
			_746_v11 = _dafny.SeqOf((_743_v8).F0(), ((_745_v10).Update(_738_v3, p1)).Cardinality())
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_744_v9), 0))
			_ = _index115
			var _rhs109 _dafny.Int = (_746_v11).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_746_v11).Select((Companion_Default___.SafeIndex((_743_v8).F0(), _dafny.IntOfUint32((_746_v11).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(_735_v0)), _dafny.IntOfUint32((_746_v11).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs109
			var _rhs110 _dafny.Sequence = p0
			_ = _rhs110
			var _rhs111 T1 = _743_v8
			_ = _rhs111
			var _rhs112 bool = p1
			_ = _rhs112
			var _lhs42 _dafny.Array = _744_v9
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_744_v9), 0))
			_ = _lhs43
			_735_v0 = _rhs109
			_742_v7 = _rhs110
			(_lhs42).ArraySet1(_rhs111, (_lhs43).Int())
			_738_v3 = _rhs112
			var _747_v12 _dafny.Array
			_ = _747_v12
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_16
			var _nw130 _dafny.Array
			_ = _nw130
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw130 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) D2 = (func(_748_v7 _dafny.Sequence, _749_v3 bool) func(_dafny.Int) D2 {
					return func(_750_i0 _dafny.Int) D2 {
						return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC4_(_748_v7, _749_v3))
					}
				})(_742_v7, _738_v3)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw130 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw130).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw130).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_747_v12 = _nw130
			var _751_v13 D2
			_ = _751_v13
			_751_v13 = Companion_D2_.Create_DC3_(_dafny.SeqOf(Companion_D0_.Create_DC1_(_738_v3)))
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_747_v12), 0))
			_ = _index116
			(_747_v12).ArraySet1(Companion_D2_.Create_DC6_(_751_v13), (_index116).Int())
			var _752_v14 D2
			_ = _752_v14
			_752_v14 = Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(_751_v13))
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_747_v12), 0))
			_ = _index117
			(_747_v12).ArraySet1(_752_v14, (_index117).Int())
		} else {
			var _753_v15 _dafny.Map
			_ = _753_v15
			_753_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F0())
			var _754_v16 _dafny.Set
			_ = _754_v16
			_754_v16 = _dafny.SetOf(p1)
			var _755_v17 _dafny.CodePoint
			_ = _755_v17
			_755_v17 = _dafny.CodePoint('b')
			var _756_v18 _dafny.Map
			_ = _756_v18
			_756_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_754_v16, _755_v17, _735_v0, _dafny.IntOfInt64(215), globalState), p1)
			var _757_v19 _dafny.Map
			_ = _757_v19
			_757_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_756_v18).Contains(_735_v0) {
					return (_756_v18).Get(_735_v0).(bool)
				}
				return p1
			})(), _753_v15)
			var _758_v20 _dafny.MultiSet
			_ = _758_v20
			_758_v20 = _dafny.MultiSetOf((_753_v15).Merge((func() _dafny.Map {
				if (_757_v19).Contains(p1) {
					return (_757_v19).Get(p1).(_dafny.Map)
				}
				return _753_v15
			})()), _753_v15, _753_v15)
			var _759_v21 bool
			_ = _759_v21
			_759_v21 = true
			var _rhs113 _dafny.MultiSet = _758_v20
			_ = _rhs113
			var _rhs114 bool = p1
			_ = _rhs114
			_758_v20 = _rhs113
			_759_v21 = _rhs114
			var _760_v22 D5
			_ = _760_v22
			_760_v22 = Companion_D5_.Create_DC13_(_735_v0, _754_v16, _755_v17, (_this).F0(), _759_v21)
			_759_v21 = (_760_v22).Dtor_cf21()
			var _761_v23 _dafny.Map
			_ = _761_v23
			_761_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
			var _762_v24 _dafny.Map
			_ = _762_v24
			_762_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _761_v23)
			var _763_v25 _dafny.Array
			_ = _763_v25
			var _nwElement0_24 _dafny.Map = _761_v23
			_ = _nwElement0_24
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(6))
			_ = _nw131
			(_nw131).ArraySet1(_nwElement0_24, 0)
			(_nw131).ArraySet1((_761_v23).Update(p1, _759_v21), 1)
			(_nw131).ArraySet1(_761_v23, 2)
			(_nw131).ArraySet1(_761_v23, 3)
			(_nw131).ArraySet1((func() _dafny.Map {
				if (_762_v24).Contains(p1) {
					return (_762_v24).Get(p1).(_dafny.Map)
				}
				return _761_v23
			})(), 4)
			(_nw131).ArraySet1(_761_v23, 5)
			_763_v25 = _nw131
			var _764_v26 _dafny.Array
			_ = _764_v26
			_764_v26 = _763_v25
			var _source8 _dafny.Array = _764_v26
			_ = _source8
			var _765___mcc_h0 _dafny.Array = _source8
			_ = _765___mcc_h0
			var _766_cf31 _dafny.Array = _765___mcc_h0
			_ = _766_cf31
			var _767_v27 _dafny.Array
			_ = _767_v27
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw132
			_767_v27 = _nw132
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))
			_ = _index118
			(_767_v27).ArraySet1((true) || (true), (_index118).Int())
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))
			_ = _index119
			var _rhs115 bool = p1
			_ = _rhs115
			var _rhs116 bool = (_754_v16).IsSubsetOf((_dafny.SetOf(!(p1), true, _759_v21, false, _759_v21)).Intersection(_754_v16))
			_ = _rhs116
			var _rhs117 bool = _759_v21
			_ = _rhs117
			var _lhs44 _dafny.Array = _767_v27
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))
			_ = _lhs45
			_759_v21 = _rhs115
			(_lhs44).ArraySet1(_rhs116, (_lhs45).Int())
			_759_v21 = _rhs117
			_759_v21 = _759_v21
			var _768_v28 _dafny.Map
			_ = _768_v28
			_768_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus((_this).F0()))
			var _769_v29 _dafny.Map
			_ = _769_v29
			_769_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_768_v28).Cardinality(), _755_v17)
			var _770_v30 _dafny.Map
			_ = _770_v30
			_770_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_735_v0, _767_v27)
			_769_v29 = (_769_v29).Update((func() _dafny.Int {
				if _759_v21 {
					return (_770_v30).Cardinality()
				}
				return _735_v0
			})(), _dafny.CodePoint('b'))
			var _771_v31 _dafny.Sequence
			_ = _771_v31
			_771_v31 = _dafny.UnicodeSeqOfUtf8Bytes("fdbsmm")
			var _772_v32 D0
			_ = _772_v32
			_772_v32 = Companion_D0_.Create_DC1_((_767_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))).Int()).(bool))
			var _773_v33 _dafny.Sequence
			_ = _773_v33
			_773_v33 = _dafny.SeqOf(_772_v32)
			var _774_v34 D2
			_ = _774_v34
			_774_v34 = Companion_D2_.Create_DC3_(_773_v33)
			var _775_v35 _dafny.Sequence
			_ = _775_v35
			_775_v35 = _dafny.SeqOf(!(_759_v21))
			var _776_v36 _dafny.Array
			_ = _776_v36
			var _nwElement0_25 _dafny.Array = _767_v27
			_ = _nwElement0_25
			var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(16))
			_ = _nw133
			(_nw133).ArraySet1(_nwElement0_25, 0)
			(_nw133).ArraySet1(_767_v27, 1)
			(_nw133).ArraySet1(_767_v27, 2)
			(_nw133).ArraySet1(_767_v27, 3)
			(_nw133).ArraySet1(_767_v27, 4)
			(_nw133).ArraySet1(_767_v27, 5)
			(_nw133).ArraySet1(_767_v27, 6)
			(_nw133).ArraySet1(_767_v27, 7)
			(_nw133).ArraySet1(_767_v27, 8)
			(_nw133).ArraySet1(_767_v27, 9)
			(_nw133).ArraySet1(_767_v27, 10)
			(_nw133).ArraySet1(_767_v27, 11)
			(_nw133).ArraySet1(_767_v27, 12)
			(_nw133).ArraySet1(_767_v27, 13)
			(_nw133).ArraySet1(_767_v27, 14)
			(_nw133).ArraySet1(_767_v27, 15)
			_776_v36 = _nw133
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_776_v36), 0))
			_ = _index120
			(_776_v36).ArraySet1(_767_v27, (_index120).Int())
			var _777_v37 T0
			_ = _777_v37
			var _nw134 *C1 = New_C1_()
			_ = _nw134
			_nw134.Ctor__(p3)
			_777_v37 = _nw134
			var _778_v38 D10
			_ = _778_v38
			_778_v38 = Companion_D10_.Create_DC21_(_777_v37)
			var _779_v39 _dafny.Map
			_ = _779_v39
			_779_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_759_v21, (_778_v38).Dtor_cf33())
			var _780_v40 _dafny.Sequence
			_ = _780_v40
			_780_v40 = _dafny.SeqOf(_779_v39, _779_v39)
			var _781_v41 _dafny.MultiSet
			_ = _781_v41
			_781_v41 = _dafny.MultiSetOf(((_780_v40).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_780_v40).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))
			_ = _index121
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_776_v36), 0))
			_ = _index122
			var _rhs118 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("evawlaen"), (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("evawlaen")).Cardinality()))).Uint32(), _755_v17)
			_ = _rhs118
			var _rhs119 bool = !((_767_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))).Int()).(bool))
			_ = _rhs119
			var _rhs120 D2 = _774_v34
			_ = _rhs120
			var _rhs121 _dafny.Sequence = (func() _dafny.Sequence {
				if ((_781_v41).Cardinality()).Cmp((_this).F0()) != 0 {
					return _dafny.SeqOf(p1)
				}
				return _775_v35
			})()
			_ = _rhs121
			var _rhs122 _dafny.Array = _767_v27
			_ = _rhs122
			var _lhs46 _dafny.Array = _767_v27
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_767_v27), 0))
			_ = _lhs47
			var _lhs48 _dafny.Array = _776_v36
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_776_v36), 0))
			_ = _lhs49
			_771_v31 = _rhs118
			(_lhs46).ArraySet1(_rhs119, (_lhs47).Int())
			_774_v34 = _rhs120
			_775_v35 = _rhs121
			(_lhs48).ArraySet1(_rhs122, (_lhs49).Int())
			var _782_v42 _dafny.Array
			_ = _782_v42
			var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw135
			_782_v42 = _nw135
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_782_v42), 0))
			_ = _index123
			(_782_v42).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vcrj"), (_index123).Int())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_782_v42), 0))
			_ = _index124
			(_782_v42).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jrrxlikqd"), (_index124).Int())
			var _783_v43 _dafny.Sequence
			_ = _783_v43
			_783_v43 = _dafny.SeqOf(_735_v0)
			var _784_v44 _dafny.Sequence
			_ = _784_v44
			_784_v44 = _dafny.SeqOf(_759_v21)
			_759_v21 = (!_dafny.Companion_Sequence_.Equal(_783_v43, _dafny.Companion_Sequence_.Update(_783_v43, (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_783_v43).Cardinality()))).Uint32(), _dafny.IntOfUint32((_784_v44).Cardinality())))) && (p1)
		}
		var _785_v45 _dafny.Set
		_ = _785_v45
		_785_v45 = _dafny.SetOf(p1)
		_735_v0 = Companion_Default___.SafeDivisionInt(p3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_785_v45).Cardinality())).Cardinality())
		var _786_v46 _dafny.CodePoint
		_ = _786_v46
		_786_v46 = _dafny.CodePoint('g')
		var _787_v47 _dafny.Map
		_ = _787_v47
		_787_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
		var _788_v48 _dafny.Map
		_ = _788_v48
		_788_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_786_v46, _787_v47)
		var _789_v49 _dafny.Array
		_ = _789_v49
		var _nwElement0_26 _dafny.Map = (_788_v48).Merge(_788_v48)
		_ = _nwElement0_26
		var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(3))
		_ = _nw136
		(_nw136).ArraySet1(_nwElement0_26, 0)
		(_nw136).ArraySet1(_788_v48, 1)
		(_nw136).ArraySet1(_788_v48, 2)
		_789_v49 = _nw136
		var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_789_v49), 0))
		_ = _index125
		(_789_v49).ArraySet1(_788_v48, (_index125).Int())
		var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_789_v49), 0))
		_ = _index126
		(_789_v49).ArraySet1(_788_v48, (_index126).Int())
		var _hi5 _dafny.Int = _dafny.IntOfInt64(-100)
		_ = _hi5
		for _790_i1 := p3; _790_i1.Cmp(_hi5) < 0; _790_i1 = _790_i1.Plus(_dafny.One) {
			var _791_v50 _dafny.Array
			_ = _791_v50
			var _nw137 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw137
			_791_v50 = _nw137
			_791_v50 = _791_v50
			_735_v0 = _735_v0
			var _792_v51 _dafny.Set
			_ = _792_v51
			_792_v51 = _dafny.SetOf(_dafny.SetOf((_this).F0(), (_dafny.Zero).Minus(p3)))
			var _793_v52 D5
			_ = _793_v52
			_793_v52 = Companion_D5_.Create_DC11_(_792_v51)
			var _source9 D5 = _793_v52
			_ = _source9
			if _source9.Is_DC12() {
				var _794___mcc_h1 _dafny.Int = _source9.Get_().(D5_DC12).Cf15
				_ = _794___mcc_h1
				var _795___mcc_h2 _dafny.Int = _source9.Get_().(D5_DC12).Cf16
				_ = _795___mcc_h2
				var _796_cf16 _dafny.Int = _795___mcc_h2
				_ = _796_cf16
				var _797_cf15 _dafny.Int = _794___mcc_h1
				_ = _797_cf15
				var _798_v53 D2
				_ = _798_v53
				_798_v53 = Companion_D2_.Create_DC4_(p0, false)
				var _799_v54 _dafny.Map
				_ = _799_v54
				_799_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v53).Dtor_cf5(), p1)
				_799_v54 = _799_v54
				var _800_v55 _dafny.MultiSet
				_ = _800_v55
				_800_v55 = _dafny.MultiSetOf(_790_i1, p3)
				var _801_v56 bool
				_ = _801_v56
				var _802_v57 _dafny.Int
				_ = _802_v57
				var _out15 bool
				_ = _out15
				var _out16 _dafny.Int
				_ = _out16
				_out15, _out16 = (_this).M6(((_800_v55).Update(_796_cf16, Companion_Default___.Abs(_797_cf15))).Cardinality(), globalState)
				_801_v56 = _out15
				_802_v57 = _out16
				var _803_v58 *C0
				_ = _803_v58
				var _nw138 *C0 = New_C0_()
				_ = _nw138
				_nw138.Ctor__(_786_v46)
				_803_v58 = _nw138
				_801_v56 = !(false)
			} else if _source9.Is_DC13() {
				var _804___mcc_h3 _dafny.Int = _source9.Get_().(D5_DC13).Cf17
				_ = _804___mcc_h3
				var _805___mcc_h4 _dafny.Set = _source9.Get_().(D5_DC13).Cf18
				_ = _805___mcc_h4
				var _806___mcc_h5 _dafny.CodePoint = _source9.Get_().(D5_DC13).Cf19
				_ = _806___mcc_h5
				var _807___mcc_h6 _dafny.Int = _source9.Get_().(D5_DC13).Cf20
				_ = _807___mcc_h6
				var _808___mcc_h7 bool = _source9.Get_().(D5_DC13).Cf21
				_ = _808___mcc_h7
				var _809_cf21 bool = _808___mcc_h7
				_ = _809_cf21
				var _810_cf20 _dafny.Int = _807___mcc_h6
				_ = _810_cf20
				var _811_cf19 _dafny.CodePoint = _806___mcc_h5
				_ = _811_cf19
				var _812_cf18 _dafny.Set = _805___mcc_h4
				_ = _812_cf18
				var _813_cf17 _dafny.Int = _804___mcc_h3
				_ = _813_cf17
				var _814_v59 _dafny.Sequence
				_ = _814_v59
				_814_v59 = _dafny.SeqOf(p1, _809_cf21)
				var _815_v60 D7
				_ = _815_v60
				_815_v60 = Companion_D7_.Create_DC16_(_814_v59)
				_815_v60 = _815_v60
				var _816_v61 _dafny.MultiSet
				_ = _816_v61
				_816_v61 = _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)), p3, (_dafny.Zero).Minus(_810_cf20))
				var _817_v62 _dafny.MultiSet
				_ = _817_v62
				_817_v62 = _dafny.MultiSetOf((_816_v61).Cardinality(), (_this).F0(), _735_v0)
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((p2), 0))
				_ = _index127
				(p2).ArraySet1(_790_i1, (_index127).Int())
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((p2), 0))
				_ = _index128
				var _rhs123 _dafny.MultiSet = (func() _dafny.MultiSet {
					if p1 {
						return _816_v61
					}
					return _817_v62
				})()
				_ = _rhs123
				var _rhs124 _dafny.Int = _dafny.IntOfInt64(268)
				_ = _rhs124
				var _lhs50 _dafny.Array = p2
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((p2), 0))
				_ = _lhs51
				_817_v62 = _rhs123
				(_lhs50).ArraySet1(_rhs124, (_lhs51).Int())
				_735_v0 = Companion_Default___.SafeDivisionInt(_790_i1, _810_cf20)
				var _818_v63 _dafny.Array
				_ = _818_v63
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw139
				_818_v63 = _nw139
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_818_v63), 0))
				_ = _index129
				(_818_v63).ArraySet1(p1, (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_818_v63), 0))
				_ = _index130
				(_818_v63).ArraySet1(!(_809_cf21), (_index130).Int())
			} else {
				var _819___mcc_h8 _dafny.Set = _source9.Get_().(D5_DC11).Cf14
				_ = _819___mcc_h8
				var _820_cf14 _dafny.Set = _819___mcc_h8
				_ = _820_cf14
				_735_v0 = _790_i1
				var _821_v64 bool
				_ = _821_v64
				var _822_v65 _dafny.Int
				_ = _822_v65
				var _out17 bool
				_ = _out17
				var _out18 _dafny.Int
				_ = _out18
				_out17, _out18 = (_this).M6(_735_v0, globalState)
				_821_v64 = _out17
				_822_v65 = _out18
				_822_v65 = _790_i1
				var _823_v66 _dafny.Map
				_ = _823_v66
				_823_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_821_v64, _821_v64)
				var _824_v67 _dafny.Map
				_ = _824_v67
				_824_v67 = _823_v66
				_824_v67 = Companion_Default___.Fm25(globalState)
			}
			var _825_v68 T0
			_ = _825_v68
			var _nw140 *C1 = New_C1_()
			_ = _nw140
			_nw140.Ctor__((_this).F0())
			_825_v68 = _nw140
			var _826_v69 D10
			_ = _826_v69
			_826_v69 = Companion_D10_.Create_DC21_(_825_v68)
			var _827_v70 D0
			_ = _827_v70
			_827_v70 = Companion_D0_.Create_DC1_(p1)
			var _828_v71 _dafny.Sequence
			_ = _828_v71
			_828_v71 = _dafny.SeqOf(_827_v70)
			var _829_v72 _dafny.Map
			_ = _829_v72
			_829_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_826_v69, Companion_D2_.Create_DC3_(_828_v71))
			var _830_v73 _dafny.Sequence
			_ = _830_v73
			_830_v73 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_826_v69, Companion_D2_.Create_DC3_(_dafny.SeqOf(_827_v70)))).Merge(_829_v72), (_829_v72).Merge(_829_v72), (func() _dafny.Map {
				if p1 {
					return _829_v72
				}
				return _829_v72
			})(), (_829_v72).Merge(_829_v72), _829_v72)
			_830_v73 = _830_v73
		}
		var _831_v74 _dafny.MultiSet
		_ = _831_v74
		_831_v74 = _dafny.MultiSetOf(p1, p1, p1, true, p1)
		var _832_v75 _dafny.Sequence
		_ = _832_v75
		_832_v75 = _dafny.SeqOf((func() _dafny.Int {
			if (_831_v74).Contains(p1) {
				return (_831_v74).Multiplicity(p1)
			}
			return p3
		})(), p3)
		var _833_v76 D0
		_ = _833_v76
		_833_v76 = Companion_D0_.Create_DC1_(p1)
		var _834_v77 _dafny.Sequence
		_ = _834_v77
		_834_v77 = _dafny.SeqOf(_833_v76)
		var _835_v78 D2
		_ = _835_v78
		_835_v78 = Companion_D2_.Create_DC3_(_834_v77)
		var _836_v79 _dafny.Map
		_ = _836_v79
		_836_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _835_v78)
		var _837_v80 D2
		_ = _837_v80
		_837_v80 = Companion_D2_.Create_DC6_((func() D2 {
			if (_836_v79).Contains(p1) {
				return (_836_v79).Get(p1).(D2)
			}
			return _835_v78
		})())
		var _pat_let_tv15 = _832_v75
		_ = _pat_let_tv15
		var _pat_let_tv16 = _735_v0
		_ = _pat_let_tv16
		var _pat_let_tv17 = _832_v75
		_ = _pat_let_tv17
		var _pat_let_tv18 = _832_v75
		_ = _pat_let_tv18
		var _pat_let_tv19 = _832_v75
		_ = _pat_let_tv19
		var _pat_let_tv20 = _832_v75
		_ = _pat_let_tv20
		var _pat_let_tv21 = _735_v0
		_ = _pat_let_tv21
		var _pat_let_tv22 = _835_v78
		_ = _pat_let_tv22
		_832_v75 = func(_source10 D2) _dafny.Sequence {
			if _source10.Is_DC4() {
				var _838___mcc_h9 _dafny.Sequence = _source10.Get_().(D2_DC4).Cf4
				_ = _838___mcc_h9
				var _839___mcc_h10 bool = _source10.Get_().(D2_DC4).Cf5
				_ = _839___mcc_h10
				var _840_cf5 bool = _839___mcc_h10
				_ = _840_cf5
				var _841_cf4 _dafny.Sequence = _838___mcc_h9
				_ = _841_cf4
				return _pat_let_tv15
			} else if _source10.Is_DC5() {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F0(), (_this).F0(), _pat_let_tv16), _pat_let_tv17)
			} else if _source10.Is_DC3() {
				var _842___mcc_h11 _dafny.Sequence = _source10.Get_().(D2_DC3).Cf3
				_ = _842___mcc_h11
				var _843_cf3 _dafny.Sequence = _842___mcc_h11
				_ = _843_cf3
				return _pat_let_tv18
			} else {
				var _844___mcc_h12 D2 = _source10.Get_().(D2_DC6).Cf6
				_ = _844___mcc_h12
				var _845_cf6 D2 = _844___mcc_h12
				_ = _845_cf6
				return _dafny.Companion_Sequence_.Update(_pat_let_tv19, (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_pat_let_tv20).Cardinality()))).Uint32(), _pat_let_tv21)
			}
		}(func(_pat_let13_0 D2) D2 {
			return func(_846_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let14_0 D2) D2 {
					return func(_847_dt__update_hcf6_h0 D2) D2 {
						return Companion_D2_.Create_DC6_(_847_dt__update_hcf6_h0)
					}(_pat_let14_0)
				}(_pat_let_tv22)
			}(_pat_let13_0)
		}(_837_v80))
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f0 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f0 = _dafny.Zero
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

func (_this *C4) F0() _dafny.Int {
	return _this._f0
}
func (_this *C4) Ctor__(f0 _dafny.Int) {
	{
		(_this)._f0 = f0
	}
}
func (_this *C4) Fm5(globalState *GlobalState) bool {
	{
		return !(!(!(!(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vrjq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_848_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))), _dafny.CodePoint('i'))))))
	}
}
func (_this *C4) Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(true) || ((Companion_D3_.Create_DC8_(true, _dafny.IntOfInt64(-634))).Dtor_cf8())
	}
}
func (_this *C4) M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _849_v0 bool
		_ = _849_v0
		_849_v0 = false
		var _850_v1 _dafny.MultiSet
		_ = _850_v1
		_850_v1 = _dafny.MultiSetOf(_849_v0)
		r2 = Companion_Default___.SafeDivisionInt((_850_v1).Cardinality(), (p1).Times(p1))
		if _849_v0 {
			var _851_v3 _dafny.Set
			_ = _851_v3
			_851_v3 = _dafny.SetOf(p1, p1)
			var _852_v4 _dafny.Map
			_ = _852_v4
			_852_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _853_v6 _dafny.Sequence
			_ = _853_v6
			_853_v6 = _dafny.SeqOf((_this).F0())
			var _854_v7 D3
			_ = _854_v7
			_854_v7 = Companion_D3_.Create_DC7_(_853_v6)
			var _855_v8 _dafny.Array
			_ = _855_v8
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw141
			_855_v8 = _nw141
			var _856_v9 _dafny.CodePoint
			_ = _856_v9
			_856_v9 = _dafny.CodePoint('c')
			var _857_v10 *C0
			_ = _857_v10
			var _nw142 *C0 = New_C0_()
			_ = _nw142
			_nw142.Ctor__(_856_v9)
			_857_v10 = _nw142
			var _858_v11 _dafny.Sequence
			_ = _858_v11
			_858_v11 = _dafny.UnicodeSeqOfUtf8Bytes("wgt")
			var _859_v15 _dafny.Array
			_ = _859_v15
			var _nwElement0_27 _dafny.Map = func() _dafny.Map {
				var _coll19 = _dafny.NewMapBuilder()
				_ = _coll19
				for _iter21 := _dafny.Iterate((_851_v3).Elements()); ; {
					_compr_19, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _860_v2 _dafny.Int
					_860_v2 = interface{}(_compr_19).(_dafny.Int)
					if (_851_v3).Contains(_860_v2) {
						_coll19.Add((_860_v2).Times((_850_v1).Cardinality()), (_this).F0())
					}
				}
				return _coll19.ToMap()
			}()
			_ = _nwElement0_27
			var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(29))
			_ = _nw143
			(_nw143).ArraySet1(_nwElement0_27, 0)
			(_nw143).ArraySet1((_852_v4).Merge(func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(859), _dafny.IntOfInt64(772))); ; {
					_compr_20, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _861_v5 _dafny.Int
					_861_v5 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(859)).Cmp(_861_v5) <= 0) && ((_861_v5).Cmp(_dafny.IntOfInt64(772)) < 0) {
						_coll20.Add((_861_v5).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, Companion_D2_.Create_DC5_())).Cardinality()), (_852_v4).Cardinality())
					}
				}
				return _coll20.ToMap()
			}()), 1)
			(_nw143).ArraySet1(_852_v4, 2)
			(_nw143).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_854_v7).Dtor_cf7()).Cardinality()), (_this).F0()), 3)
			(_nw143).ArraySet1(_852_v4, 4)
			(_nw143).ArraySet1(_852_v4, 5)
			(_nw143).ArraySet1(_852_v4, 6)
			(_nw143).ArraySet1(_852_v4, 7)
			(_nw143).ArraySet1((_852_v4).Merge(_852_v4), 8)
			(_nw143).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _dafny.IntOfUint32((_dafny.SeqOf(_855_v8, _855_v8, _855_v8)).Cardinality())), 9)
			(_nw143).ArraySet1(_852_v4, 10)
			(_nw143).ArraySet1(_852_v4, 11)
			(_nw143).ArraySet1(Companion_Default___.Fm11((_this).F0(), globalState), 12)
			(_nw143).ArraySet1(_852_v4, 13)
			(_nw143).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_857_v10), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_851_v3).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_857_v10)).Cardinality()))).Uint32(), _857_v10)).Cardinality()), _dafny.IntOfUint32((_858_v11).Cardinality())), 14)
			(_nw143).ArraySet1(_852_v4, 15)
			(_nw143).ArraySet1(_852_v4, 16)
			(_nw143).ArraySet1((_852_v4).Merge(func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(896), _dafny.IntOfInt64(834))); ; {
					_compr_21, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _862_v12 _dafny.Int
					_862_v12 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(896)).Cmp(_862_v12) <= 0) && ((_862_v12).Cmp(_dafny.IntOfInt64(834)) < 0) {
						_coll21.Add((_862_v12).Times(_dafny.IntOfInt64(-918)), p1)
					}
				}
				return _coll21.ToMap()
			}()), 17)
			(_nw143).ArraySet1((_852_v4).Merge(_852_v4), 18)
			(_nw143).ArraySet1(_852_v4, 19)
			(_nw143).ArraySet1(_852_v4, 20)
			(_nw143).ArraySet1(func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(156), _dafny.IntOfInt64(62))); ; {
					_compr_22, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _863_v13 _dafny.Int
					_863_v13 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(156)).Cmp(_863_v13) <= 0) && ((_863_v13).Cmp(_dafny.IntOfInt64(62)) < 0) {
						_coll22.Add(Companion_Default___.SafeDivisionInt(_863_v13, p1), (_this).F0())
					}
				}
				return _coll22.ToMap()
			}(), 21)
			(_nw143).ArraySet1(_852_v4, 22)
			(_nw143).ArraySet1(_852_v4, 23)
			(_nw143).ArraySet1(Companion_Default___.Fm11(p1, globalState), 24)
			(_nw143).ArraySet1(_852_v4, 25)
			(_nw143).ArraySet1(func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(659), _dafny.IntOfInt64(736))); ; {
					_compr_23, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _864_v14 _dafny.Int
					_864_v14 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(659)).Cmp(_864_v14) <= 0) && ((_864_v14).Cmp(_dafny.IntOfInt64(736)) < 0) {
						_coll23.Add((_864_v14).Minus(p1), _dafny.IntOfInt64(225))
					}
				}
				return _coll23.ToMap()
			}(), 26)
			(_nw143).ArraySet1(_852_v4, 27)
			(_nw143).ArraySet1(_852_v4, 28)
			_859_v15 = _nw143
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_859_v15), 0))
			_ = _index131
			(_859_v15).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F0()), (_index131).Int())
			var _865_v16 _dafny.Array
			_ = _865_v16
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_17
			var _nw144 _dafny.Array
			_ = _nw144
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw144 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Sequence = (func(_866_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_867_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_866_v0, _866_v0)
					}
				})(_849_v0)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw144 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw144).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw144).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_865_v16 = _nw144
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))
			_ = _index132
			(_865_v16).ArraySet1(Companion_Default___.Fm12((_this).F0(), (_this).F0(), _849_v0, globalState), (_index132).Int())
			var _868_v17 _dafny.Array
			_ = _868_v17
			var _len0_18 _dafny.Int = _dafny.One
			_ = _len0_18
			var _nw145 _dafny.Array
			_ = _nw145
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw145 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = (func(_869_v0 bool) func(_dafny.Int) bool {
					return func(_870_i1 _dafny.Int) bool {
						return !(_869_v0) || (_869_v0)
					}
				})(_849_v0)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw145 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw145).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw145).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_868_v17 = _nw145
			var _871_v18 _dafny.Sequence
			_ = _871_v18
			_871_v18 = _dafny.SeqOf(_849_v0, _849_v0)
			var _872_v19 _dafny.Array
			_ = _872_v19
			_872_v19 = _868_v17
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_859_v15), 0))
			_ = _index133
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))
			_ = _index134
			var _rhs125 _dafny.Map = _852_v4
			_ = _rhs125
			var _rhs126 _dafny.Sequence = _858_v11
			_ = _rhs126
			var _rhs127 _dafny.Sequence = _871_v18
			_ = _rhs127
			var _rhs128 _dafny.Array = (_868_v17)
			_ = _rhs128
			var _rhs129 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_858_v11, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_873_v10 *C0) func(_dafny.Int) _dafny.CodePoint {
				return func(_874_i2 _dafny.Int) _dafny.CodePoint {
					return (_873_v10).F3()
				}
			})(_857_v10))), _858_v11)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_871_v18).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_858_v11, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}((func(_875_v10 *C0) func(_dafny.Int) _dafny.CodePoint {
				return func(_876_i2 _dafny.Int) _dafny.CodePoint {
					return (_875_v10).F3()
				}
			})(_857_v10))), _858_v11))).Cardinality()))).Uint32(), (_857_v10).F3())
			_ = _rhs129
			var _lhs52 _dafny.Array = _859_v15
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_859_v15), 0))
			_ = _lhs53
			var _lhs54 _dafny.Array = _865_v16
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))
			_ = _lhs55
			(_lhs52).ArraySet1(_rhs125, (_lhs53).Int())
			_858_v11 = _rhs126
			(_lhs54).ArraySet1(_rhs127, (_lhs55).Int())
			_868_v17 = _rhs128
			_858_v11 = _rhs129
			var _877_v20 *C0
			_ = _877_v20
			var _nw146 *C0 = New_C0_()
			_ = _nw146
			_nw146.Ctor__(_dafny.CodePoint('c'))
			_877_v20 = _nw146
			var _pat_let_tv23 = _853_v6
			_ = _pat_let_tv23
			var _pat_let_tv24 = _853_v6
			_ = _pat_let_tv24
			var _source11 D3 = func(_pat_let15_0 D3) D3 {
				return func(_878_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let16_0 _dafny.Sequence) D3 {
						return func(_879_dt__update_hcf7_h0 _dafny.Sequence) D3 {
							return Companion_D3_.Create_DC7_(_879_dt__update_hcf7_h0)
						}(_pat_let16_0)
					}(_dafny.Companion_Sequence_.Update(_pat_let_tv23, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.IntOfUint32((_pat_let_tv24).Cardinality()))).Uint32(), (_this).F0()))
				}(_pat_let15_0)
			}(_854_v7)
			_ = _source11
			if _source11.Is_DC8() {
				var _880___mcc_h0 bool = _source11.Get_().(D3_DC8).Cf8
				_ = _880___mcc_h0
				var _881___mcc_h1 _dafny.Int = _source11.Get_().(D3_DC8).Cf9
				_ = _881___mcc_h1
				var _882_cf9 _dafny.Int = _881___mcc_h1
				_ = _882_cf9
				var _883_cf8 bool = _880___mcc_h0
				_ = _883_cf8
				var _884_v21 T0
				_ = _884_v21
				var _nw147 *C3 = New_C3_()
				_ = _nw147
				_nw147.Ctor__((_this).F0())
				_884_v21 = _nw147
				var _885_v22 _dafny.Array
				_ = _885_v22
				var _nwElement0_28 T0 = _884_v21
				_ = _nwElement0_28
				var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(25))
				_ = _nw148
				(_nw148).ArraySet1(_nwElement0_28, 0)
				(_nw148).ArraySet1(_884_v21, 1)
				(_nw148).ArraySet1(_884_v21, 2)
				(_nw148).ArraySet1(_884_v21, 3)
				(_nw148).ArraySet1(_884_v21, 4)
				(_nw148).ArraySet1(_884_v21, 5)
				(_nw148).ArraySet1(_884_v21, 6)
				(_nw148).ArraySet1(_884_v21, 7)
				(_nw148).ArraySet1(_884_v21, 8)
				(_nw148).ArraySet1(_884_v21, 9)
				(_nw148).ArraySet1(_884_v21, 10)
				(_nw148).ArraySet1(_884_v21, 11)
				(_nw148).ArraySet1(_884_v21, 12)
				(_nw148).ArraySet1(_884_v21, 13)
				(_nw148).ArraySet1(_884_v21, 14)
				(_nw148).ArraySet1(_884_v21, 15)
				(_nw148).ArraySet1(_884_v21, 16)
				(_nw148).ArraySet1(_884_v21, 17)
				(_nw148).ArraySet1(_884_v21, 18)
				(_nw148).ArraySet1(_884_v21, 19)
				(_nw148).ArraySet1(_884_v21, 20)
				(_nw148).ArraySet1(_884_v21, 21)
				(_nw148).ArraySet1(_884_v21, 22)
				(_nw148).ArraySet1(_884_v21, 23)
				(_nw148).ArraySet1(_884_v21, 24)
				_885_v22 = _nw148
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_885_v22), 0))
				_ = _index135
				(_885_v22).ArraySet1(_884_v21, (_index135).Int())
				var _886_v23 D5
				_ = _886_v23
				_886_v23 = Companion_D5_.Create_DC13_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_858_v11).Cardinality()), _849_v0)).Cardinality(), _dafny.SetOf(_883_cf8), (_858_v11).Select((Companion_Default___.SafeIndex((_884_v21).F0(), _dafny.IntOfUint32((_858_v11).Cardinality()))).Uint32()).(_dafny.CodePoint), _882_cf9, _849_v0)
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_885_v22), 0))
				_ = _index136
				(_885_v22).ArraySet1((func() T0 {
					if _883_cf8 {
						return (func() T0 {
							if (_886_v23).Dtor_cf21() {
								return _884_v21
							}
							return _884_v21
						})()
					}
					return _884_v21
				})(), (_index136).Int())
				r1 = _882_cf9
				var _887_v24 _dafny.Set
				_ = _887_v24
				_887_v24 = _dafny.SetOf(_849_v0, _883_cf8, _883_cf8, _883_cf8)
				r1 = (((_this).F0()).Times(Companion_Default___.Fm3(_887_v24, _856_v9, _882_cf9, _dafny.IntOfUint32((_858_v11).Cardinality()), globalState))).Times((_884_v21).F0())
				var _888_v25 _dafny.Array
				_ = _888_v25
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw149
				_888_v25 = _nw149
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_888_v25), 0))
				_ = _index137
				(_888_v25).ArraySet1(_868_v17, (_index137).Int())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_888_v25), 0))
				_ = _index138
				(_888_v25).ArraySet1(_868_v17, (_index138).Int())
			} else if _source11.Is_DC9() {
				var _889___mcc_h2 _dafny.Int = _source11.Get_().(D3_DC9).Cf10
				_ = _889___mcc_h2
				var _890___mcc_h3 _dafny.Array = _source11.Get_().(D3_DC9).Cf11
				_ = _890___mcc_h3
				var _891___mcc_h4 bool = _source11.Get_().(D3_DC9).Cf12
				_ = _891___mcc_h4
				var _892_cf12 bool = _891___mcc_h4
				_ = _892_cf12
				var _893_cf11 _dafny.Array = _890___mcc_h3
				_ = _893_cf11
				var _894_cf10 _dafny.Int = _889___mcc_h2
				_ = _894_cf10
				_849_v0 = _849_v0
				var _895_v26 *C1
				_ = _895_v26
				var _nw150 *C1 = New_C1_()
				_ = _nw150
				_nw150.Ctor__(Companion_Default___.SafeModuloInt(_894_cf10, _dafny.IntOfUint32(((_865_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))).Int()).(_dafny.Sequence)).Cardinality())))
				_895_v26 = _nw150
				var _896_v27 *C1
				_ = _896_v27
				var _nw151 *C1 = New_C1_()
				_ = _nw151
				_nw151.Ctor__((_this).F0())
				_896_v27 = _nw151
				var _897_v28 _dafny.Map
				_ = _897_v28
				_897_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, p1)
				var _898_v29 _dafny.Sequence
				_ = _898_v29
				_898_v29 = _dafny.SeqOf(_897_v28, _897_v28, _897_v28)
				var _899_v30 D3
				_ = _899_v30
				_899_v30 = Companion_D3_.Create_DC8_(_849_v0, p1)
				var _900_v31 _dafny.Map
				_ = _900_v31
				_900_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_898_v29).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_898_v29).Cardinality()))).Uint32()).(_dafny.Map), (_this).Fm6((_899_v30).Dtor_cf8(), _858_v11, _858_v11, _dafny.IntOfInt64(432), globalState))
				_900_v31 = (_900_v31).Update(_897_v28, _892_cf12)
			} else {
				var _901___mcc_h5 _dafny.Sequence = _source11.Get_().(D3_DC7).Cf7
				_ = _901___mcc_h5
				var _902_cf7 _dafny.Sequence = _901___mcc_h5
				_ = _902_cf7
				var _903_v32 _dafny.Map
				_ = _903_v32
				_903_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _868_v17)
				var _904_v33 _dafny.Set
				_ = _904_v33
				_904_v33 = _dafny.SetOf(_849_v0)
				var _905_v34 _dafny.Map
				_ = _905_v34
				_905_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_904_v33).Cardinality(), _868_v17))
				var _906_v35 _dafny.Map
				_ = _906_v35
				_906_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _849_v0)
				var _rhs130 _dafny.Map = (func() _dafny.Map {
					if (_905_v34).Contains((func() bool {
						if (_906_v35).Contains((_871_v18).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_871_v18).Cardinality()))).Uint32()).(bool)) {
							return (_906_v35).Get((_871_v18).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_871_v18).Cardinality()))).Uint32()).(bool)).(bool)
						}
						return Companion_Default___.Fm2(_849_v0, (_this).F0(), _849_v0, _849_v0, globalState)
					})()) {
						return (_905_v34).Get((func() bool {
							if (_906_v35).Contains((_871_v18).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_871_v18).Cardinality()))).Uint32()).(bool)) {
								return (_906_v35).Get((_871_v18).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_871_v18).Cardinality()))).Uint32()).(bool)).(bool)
							}
							return Companion_Default___.Fm2(_849_v0, (_this).F0(), _849_v0, _849_v0, globalState)
						})()).(_dafny.Map)
					}
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _868_v17)).Update(p1, _868_v17)
				})()
				_ = _rhs130
				var _rhs131 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F0(), p1)))
				_ = _rhs131
				_903_v32 = _rhs130
				r2 = _rhs131
				var _907_v36 _dafny.MultiSet
				_ = _907_v36
				_907_v36 = _dafny.MultiSetOf((_this).F0(), Companion_Default___.Fm3(_904_v33, (_857_v10).F3(), p1, (_dafny.Zero).Minus((_this).F0()), globalState))
				var _908_v37 T1
				_ = _908_v37
				var _nw152 *C2 = New_C2_()
				_ = _nw152
				_nw152.Ctor__(p1)
				_908_v37 = _nw152
				var _909_v38 _dafny.Set
				_ = _909_v38
				_909_v38 = _dafny.SetOf(_908_v37)
				var _910_v39 *C2
				_ = _910_v39
				var _nw153 *C2 = New_C2_()
				_ = _nw153
				_nw153.Ctor__(Companion_Default___.Fm3(_904_v33, (_857_v10).F3(), p1, (func() _dafny.Int {
					if (_907_v36).Contains((_909_v38).Cardinality()) {
						return (_907_v36).Multiplicity((_909_v38).Cardinality())
					}
					return p1
				})(), globalState))
				_910_v39 = _nw153
				_910_v39 = _910_v39
				var _911_v40 _dafny.Map
				_ = _911_v40
				_911_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_855_v8, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, _849_v0)).Update(_849_v0, _849_v0))
				var _912_v41 _dafny.Map
				_ = _912_v41
				_912_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(467), _906_v35)
				var _913_v42 _dafny.Map
				_ = _913_v42
				_913_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_857_v10).F3(), (_908_v37).F0())
				var _914_v43 _dafny.Map
				_ = _914_v43
				_914_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_849_v0), (func() _dafny.Int {
					if (_913_v42).Contains((_857_v10).F3()) {
						return (_913_v42).Get((_857_v10).F3()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_858_v11).Cardinality())
				})())
				var _915_v44 _dafny.Map
				_ = _915_v44
				_915_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _849_v0)
				_906_v35 = ((func() _dafny.Map {
					if (_911_v40).Contains(_855_v8) {
						return (_911_v40).Get(_855_v8).(_dafny.Map)
					}
					return (func() _dafny.Map {
						if (_912_v41).Contains((_914_v43).Cardinality()) {
							return (_912_v41).Get((_914_v43).Cardinality()).(_dafny.Map)
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, (func() bool {
							if (_915_v44).Contains(p1) {
								return (_915_v44).Get(p1).(bool)
							}
							return _849_v0
						})())
					})()
				})()).Update(_849_v0, _849_v0)
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_855_v8), 0))
				_ = _index139
				(_855_v8).ArraySet1((_dafny.Zero).Minus((_908_v37).F0()), (_index139).Int())
				var _916_v45 _dafny.Map
				_ = _916_v45
				_916_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_849_v0), _dafny.SetOf(_849_v0))
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_855_v8), 0))
				_ = _index140
				(_855_v8).ArraySet1(Companion_Default___.Fm3(((func() _dafny.Set {
					if (_916_v45).Contains(_849_v0) {
						return (_916_v45).Get(_849_v0).(_dafny.Set)
					}
					return _904_v33
				})()).Difference(_904_v33), _dafny.CodePoint('i'), (_908_v37).F0(), (_this).F0(), globalState), (_index140).Int())
			}
			var _917_v46 _dafny.Set
			_ = _917_v46
			_917_v46 = _dafny.SetOf(!(_849_v0), _849_v0, _849_v0)
			var _918_v47 _dafny.MultiSet
			_ = _918_v47
			_918_v47 = _dafny.MultiSetOf(_dafny.IntOfInt64(721))
			var _919_v48 D0
			_ = _919_v48
			_919_v48 = Companion_D0_.Create_DC1_(((_865_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_865_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(bool))
			var _920_v49 D2
			_ = _920_v49
			_920_v49 = Companion_D2_.Create_DC3_(_dafny.SeqOf(_919_v48))
			var _921_v50 _dafny.Map
			_ = _921_v50
			_921_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _920_v49)
			var _922_v51 _dafny.MultiSet
			_ = _922_v51
			_922_v51 = _dafny.MultiSetOf(Companion_D2_.Create_DC6_((func() D2 {
				if (_921_v50).Contains(p1) {
					return (_921_v50).Get(p1).(D2)
				}
				return _920_v49
			})()), Companion_D2_.Create_DC6_(_920_v49))
			var _923_v52 _dafny.Map
			_ = _923_v52
			_923_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm18(Companion_Default___.Fm3(_917_v46, _856_v9, (_918_v47).Cardinality(), (_this).F0(), globalState), globalState), _922_v51)
			var _924_v53 *C3
			_ = _924_v53
			var _nw154 *C3 = New_C3_()
			_ = _nw154
			_nw154.Ctor__((_850_v1).Cardinality())
			_924_v53 = _nw154
			var _925_v54 _dafny.Map
			_ = _925_v54
			_925_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _924_v53)
			var _926_v55 _dafny.Map
			_ = _926_v55
			_926_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
				if (_923_v52).Contains(_858_v11) {
					return (_923_v52).Get(_858_v11).(_dafny.MultiSet)
				}
				return _922_v51
			})(), (func() *C3 {
				if (_925_v54).Contains(_dafny.IntOfInt64(693)) {
					return (_925_v54).Get(_dafny.IntOfInt64(693)).(*C3)
				}
				return _924_v53
			})())
			_926_v55 = (_926_v55).Update(_922_v51, _924_v53)
			if (p1).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_858_v11).Cardinality()), (_this).F0())) < 0 {
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_868_v17), 0))
				_ = _index141
				(_868_v17).ArraySet1(_849_v0, (_index141).Int())
				var _927_v56 D2
				_ = _927_v56
				_927_v56 = Companion_D2_.Create_DC4_(_858_v11, _849_v0)
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_868_v17), 0))
				_ = _index142
				var _rhs132 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("k"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(95))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_928_v20 *C0) func(_dafny.Int) _dafny.CodePoint {
					return func(_929_i3 _dafny.Int) _dafny.CodePoint {
						return (_928_v20).F3()
					}
				})(_877_v20))))
				_ = _rhs132
				var _rhs133 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_858_v11, (_927_v56).Dtor_cf4()))
				_ = _rhs133
				var _lhs56 _dafny.Array = _868_v17
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_868_v17), 0))
				_ = _lhs57
				_858_v11 = _rhs132
				(_lhs56).ArraySet1(_rhs133, (_lhs57).Int())
				(_924_v53).M7(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jljbktggr"), _858_v11), _849_v0, _855_v8, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _849_v0 {
						return _853_v6
					}
					return _853_v6
				})()).Cardinality()), globalState)
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_868_v17), 0))
				_ = _index143
				(_868_v17).ArraySet1((p1).Cmp(p1) > 0, (_index143).Int())
				var _930_v57 *C1
				_ = _930_v57
				var _nw155 *C1 = New_C1_()
				_ = _nw155
				_nw155.Ctor__((_this).F0())
				_930_v57 = _nw155
				var _931_v58 _dafny.Map
				_ = _931_v58
				_931_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _849_v0)
				var _932_v59 _dafny.Map
				_ = _932_v59
				_932_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_930_v57, _931_v58)
				var _933_v60 _dafny.Map
				_ = _933_v60
				_933_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, _932_v59)
				_933_v60 = (_933_v60).Update((_930_v57).Fm6(_849_v0, _858_v11, _858_v11, p1, globalState), (func() _dafny.Map {
					if _849_v0 {
						return _932_v59
					}
					return (_932_v59).Update(_930_v57, _931_v58)
				})())
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))
				_ = _index144
				(_865_v16).ArraySet1((_865_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(859), _dafny.ArrayLen((_865_v16), 0))).Int()).(_dafny.Sequence), (_index144).Int())
			} else {
				_856_v9 = (_857_v10).F3()
				var _934_v61 *C3
				_ = _934_v61
				var _nw156 *C3 = New_C3_()
				_ = _nw156
				_nw156.Ctor__(p1)
				_934_v61 = _nw156
				var _935_v62 _dafny.Array
				_ = _935_v62
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
				_ = _nw157
				_935_v62 = _nw157
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_935_v62), 0))
				_ = _index145
				(_935_v62).ArraySet1(_917_v46, (_index145).Int())
				var _936_v63 _dafny.Array
				_ = _936_v63
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(27))
				_ = _nw158
				_936_v63 = _nw158
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_936_v63), 0))
				_ = _index146
				(_936_v63).ArraySet1(_918_v47, (_index146).Int())
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_935_v62), 0))
				_ = _index147
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_936_v63), 0))
				_ = _index148
				var _rhs134 _dafny.Int = (func() _dafny.Int {
					if Companion_Default___.Fm2(_849_v0, (_this).F0(), !(_849_v0), false, globalState) {
						return p1
					}
					return Companion_Default___.Fm3(_917_v46, (_877_v20).F3(), (Companion_Default___.Fm11(_dafny.IntOfUint32((_858_v11).Cardinality()), globalState)).Cardinality(), (_851_v3).Cardinality(), globalState)
				})()
				_ = _rhs134
				var _rhs135 _dafny.Int = (_dafny.Zero).Minus((_this).F0())
				_ = _rhs135
				var _rhs136 _dafny.Array = _855_v8
				_ = _rhs136
				var _rhs137 _dafny.Set = _917_v46
				_ = _rhs137
				var _rhs138 _dafny.MultiSet = _dafny.MultiSetOf((_this).F0())
				_ = _rhs138
				var _lhs58 _dafny.Array = _935_v62
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_935_v62), 0))
				_ = _lhs59
				var _lhs60 _dafny.Array = _936_v63
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_936_v63), 0))
				_ = _lhs61
				r2 = _rhs134
				r2 = _rhs135
				_855_v8 = _rhs136
				(_lhs58).ArraySet1(_rhs137, (_lhs59).Int())
				(_lhs60).ArraySet1(_rhs138, (_lhs61).Int())
				_856_v9 = (func() _dafny.CodePoint {
					if _849_v0 {
						return _856_v9
					}
					return (_877_v20).F3()
				})()
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_868_v17), 0))
				_ = _index149
				(_868_v17).ArraySet1(true, (_index149).Int())
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_868_v17), 0))
				_ = _index150
				(_868_v17).ArraySet1(!(Companion_Default___.Fm2(((_935_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_935_v62), 0))).Int()).(_dafny.Set)).IsProperSubsetOf(_dafny.SetOf(_849_v0)), p1, (true) || (_849_v0), (func() bool {
					if _849_v0 {
						return _849_v0
					}
					return _849_v0
				})(), globalState)), (_index150).Int())
			}
		} else {
			var _937_v64 *C3
			_ = _937_v64
			var _nw159 *C3 = New_C3_()
			_ = _nw159
			_nw159.Ctor__(p1)
			_937_v64 = _nw159
			if _849_v0 {
				var _938_v65 _dafny.Array
				_ = _938_v65
				var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw160
				_938_v65 = _nw160
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))
				_ = _index151
				(_938_v65).ArraySet1(p1, (_index151).Int())
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))
				_ = _index152
				(_938_v65).ArraySet1(p1, (_index152).Int())
				var _939_v66 _dafny.Array
				_ = _939_v66
				var _len0_19 _dafny.Int = _dafny.One
				_ = _len0_19
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = (func(_940_p1 _dafny.Int) func(_dafny.Int) bool {
						return func(_941_i4 _dafny.Int) bool {
							return (_940_p1).Cmp(_dafny.IntOfInt64(854)) >= 0
						}
					})(p1)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw161 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw161).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw161).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_939_v66 = _nw161
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_939_v66), 0))
				_ = _index153
				(_939_v66).ArraySet1(_849_v0, (_index153).Int())
				var _942_v67 _dafny.Sequence
				_ = _942_v67
				_942_v67 = _dafny.SeqOf(!(_849_v0))
				var _943_v68 D7
				_ = _943_v68
				_943_v68 = Companion_D7_.Create_DC16_(_942_v67)
				var _944_v69 _dafny.CodePoint
				_ = _944_v69
				_944_v69 = _dafny.CodePoint('k')
				var _945_v70 _dafny.CodePoint
				_ = _945_v70
				_945_v70 = _944_v69
				var _946_v71 _dafny.Map
				_ = _946_v71
				_946_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int)), _945_v70)
				var _947_v72 _dafny.Sequence
				_ = _947_v72
				_947_v72 = _dafny.SeqOf(_938_v65)
				var _948_v73 _dafny.Sequence
				_ = _948_v73
				_948_v73 = _dafny.SeqOf((_this).F0())
				var _949_v74 _dafny.Set
				_ = _949_v74
				_949_v74 = _dafny.SetOf(_849_v0)
				var _950_v75 _dafny.MultiSet
				_ = _950_v75
				_950_v75 = _dafny.MultiSetOf(_dafny.IntOfUint32((_948_v73).Cardinality()), Companion_Default___.Fm3(_949_v74, _dafny.CodePoint('q'), _dafny.IntOfInt64(-458), _dafny.IntOfInt64(-6), globalState))
				var _951_v76 _dafny.Sequence
				_ = _951_v76
				_951_v76 = _dafny.SeqOf(p1, _dafny.IntOfUint32(((Companion_D11_.Create_DC23_(_947_v72)).Dtor_cf37()).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xurqdyud")).Cardinality()), (_this).F0(), (func() _dafny.Int {
					if (_950_v75).Contains((_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int)) {
						return (_950_v75).Multiplicity((_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int))
					}
					return p1
				})())
				var _952_v77 _dafny.Sequence
				_ = _952_v77
				_952_v77 = _dafny.SeqOf((_946_v71).Cardinality(), (_dafny.IntOfUint32((_951_v76).Cardinality())).Plus((_950_v75).Cardinality()))
				var _953_v78 _dafny.Sequence
				_ = _953_v78
				_953_v78 = _dafny.SeqOf(_952_v77, _951_v76, _dafny.SeqOf((_this).F0()), _dafny.SeqOf((_this).F0()))
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_939_v66), 0))
				_ = _index154
				var _rhs139 bool = _849_v0
				_ = _rhs139
				var _rhs140 D7 = _943_v68
				_ = _rhs140
				var _rhs141 _dafny.Sequence = (_953_v78).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_953_v78).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs141
				var _rhs142 bool = !(_849_v0) || (_849_v0)
				_ = _rhs142
				var _lhs62 _dafny.Array = _939_v66
				_ = _lhs62
				var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_939_v66), 0))
				_ = _lhs63
				(_lhs62).ArraySet1(_rhs139, (_lhs63).Int())
				_943_v68 = _rhs140
				_952_v77 = _rhs141
				_849_v0 = _rhs142
				var _954_v79 D7
				_ = _954_v79
				_954_v79 = Companion_D7_.Create_DC18_((_dafny.Zero).Minus((_this).F0()))
				var _955_v80 *C1
				_ = _955_v80
				var _nw162 *C1 = New_C1_()
				_ = _nw162
				_nw162.Ctor__((_this).F0())
				_955_v80 = _nw162
				var _pat_let_tv25 = _955_v80
				_ = _pat_let_tv25
				var _pat_let_tv26 = _849_v0
				_ = _pat_let_tv26
				var _956_v81 _dafny.Array
				_ = _956_v81
				var _nwElement0_29 D7 = Companion_D7_.Create_DC18_((_dafny.Zero).Minus((func() _dafny.Int {
					if (_850_v1).Contains(_849_v0) {
						return (_850_v1).Multiplicity(_849_v0)
					}
					return (_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int)
				})()))
				_ = _nwElement0_29
				var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(29))
				_ = _nw163
				(_nw163).ArraySet1(_nwElement0_29, 0)
				(_nw163).ArraySet1(_954_v79, 1)
				(_nw163).ArraySet1(_954_v79, 2)
				(_nw163).ArraySet1(_954_v79, 3)
				(_nw163).ArraySet1(_954_v79, 4)
				(_nw163).ArraySet1(_954_v79, 5)
				(_nw163).ArraySet1(Companion_D7_.Create_DC18_((_this).F0()), 6)
				(_nw163).ArraySet1(func(_pat_let17_0 D7) D7 {
					return func(_959_dt__update__tmp_h3 D7) D7 {
						return func(_pat_let20_0 _dafny.Int) D7 {
							return func(_960_dt__update_hcf30_h1 _dafny.Int) D7 {
								return Companion_D7_.Create_DC18_(_960_dt__update_hcf30_h1)
							}(_pat_let20_0)
						}((_this).F0())
					}(_pat_let17_0)
				}(func(_pat_let18_0 D7) D7 {
					return func(_957_dt__update__tmp_h2 D7) D7 {
						return func(_pat_let19_0 _dafny.Int) D7 {
							return func(_958_dt__update_hcf30_h0 _dafny.Int) D7 {
								return Companion_D7_.Create_DC18_(_958_dt__update_hcf30_h0)
							}(_pat_let19_0)
						}((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv25, _pat_let_tv26)).Cardinality())
					}(_pat_let18_0)
				}(_954_v79)), 7)
				(_nw163).ArraySet1(_954_v79, 8)
				(_nw163).ArraySet1(Companion_D7_.Create_DC18_(Companion_Default___.Fm3(_949_v74, _944_v69, p1, (_850_v1).Cardinality(), globalState)), 9)
				(_nw163).ArraySet1(Companion_Default___.Fm26(globalState), 10)
				(_nw163).ArraySet1(_954_v79, 11)
				(_nw163).ArraySet1(Companion_D7_.Create_DC18_(p1), 12)
				(_nw163).ArraySet1(_954_v79, 13)
				(_nw163).ArraySet1(_954_v79, 14)
				(_nw163).ArraySet1(_954_v79, 15)
				(_nw163).ArraySet1(_954_v79, 16)
				(_nw163).ArraySet1(Companion_D7_.Create_DC18_(p1), 17)
				(_nw163).ArraySet1(Companion_Default___.Fm26(globalState), 18)
				(_nw163).ArraySet1(Companion_D7_.Create_DC18_((_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int)), 19)
				(_nw163).ArraySet1(_954_v79, 20)
				(_nw163).ArraySet1(_954_v79, 21)
				(_nw163).ArraySet1((func() D7 {
					if _849_v0 {
						return _954_v79
					}
					return _954_v79
				})(), 22)
				(_nw163).ArraySet1(_954_v79, 23)
				(_nw163).ArraySet1(_954_v79, 24)
				(_nw163).ArraySet1(_954_v79, 25)
				(_nw163).ArraySet1(_954_v79, 26)
				(_nw163).ArraySet1(_954_v79, 27)
				(_nw163).ArraySet1(_954_v79, 28)
				_956_v81 = _nw163
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_956_v81), 0))
				_ = _index155
				(_956_v81).ArraySet1(_954_v79, (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_956_v81), 0))
				_ = _index156
				var _rhs143 D7 = _954_v79
				_ = _rhs143
				var _rhs144 _dafny.Set = (_949_v74).Union(_949_v74)
				_ = _rhs144
				var _rhs145 _dafny.Int = (_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int)
				_ = _rhs145
				var _lhs64 _dafny.Array = _956_v81
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_956_v81), 0))
				_ = _lhs65
				(_lhs64).ArraySet1(_rhs143, (_lhs65).Int())
				_949_v74 = _rhs144
				r1 = _rhs145
				var _961_v82 _dafny.Sequence
				_ = _961_v82
				_961_v82 = _dafny.UnicodeSeqOfUtf8Bytes("aixybf")
				_961_v82 = _961_v82
				r1 = (Companion_Default___.SafeDivisionInt((_938_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_938_v65), 0))).Int()).(_dafny.Int), (_this).F0())).Plus((_dafny.Zero).Minus((_this).F0()))
			} else {
				var _962_v83 _dafny.Array
				_ = _962_v83
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_20
				var _nw164 _dafny.Array
				_ = _nw164
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw164 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) bool = (func(_963_v0 bool) func(_dafny.Int) bool {
						return func(_964_i5 _dafny.Int) bool {
							return _963_v0
						}
					})(_849_v0)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw164 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw164).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw164).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_962_v83 = _nw164
				var _965_v84 _dafny.Sequence
				_ = _965_v84
				_965_v84 = _dafny.UnicodeSeqOfUtf8Bytes("elco")
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _index157
				(_962_v83).ArraySet1((_937_v64).Fm6(_849_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(87))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}(func(_966_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				})), _965_v84, (func() _dafny.Int {
					if (_850_v1).Contains(!(_849_v0)) {
						return (_850_v1).Multiplicity(!(_849_v0))
					}
					return p1
				})(), globalState), (_index157).Int())
				var _967_v85 T0
				_ = _967_v85
				var _nw165 *C3 = New_C3_()
				_ = _nw165
				_nw165.Ctor__(_dafny.IntOfUint32((_965_v84).Cardinality()))
				_967_v85 = _nw165
				var _968_v86 _dafny.Sequence
				_ = _968_v86
				_968_v86 = _dafny.SeqOf(_849_v0, _849_v0)
				var _969_v87 D7
				_ = _969_v87
				_969_v87 = Companion_D7_.Create_DC16_(_968_v86)
				var _970_v88 _dafny.MultiSet
				_ = _970_v88
				_970_v88 = _dafny.MultiSetOf(_969_v87)
				var _971_v89 _dafny.CodePoint
				_ = _971_v89
				_971_v89 = _dafny.CodePoint('d')
				var _972_v90 _dafny.Set
				_ = _972_v90
				_972_v90 = _dafny.SetOf(_849_v0, Companion_Default___.Fm2(_849_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, p1)).Cardinality(), _849_v0, _849_v0, globalState))
				var _973_v91 _dafny.Map
				_ = _973_v91
				_973_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_967_v85).F0())), (_this).F0())
				var _974_v92 _dafny.MultiSet
				_ = _974_v92
				_974_v92 = _dafny.MultiSetOf(p1, (_this).F0())
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _index158
				var _rhs146 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm3(Companion_Default___.Fm27(_849_v0, _970_v88, p1, _dafny.SetOf((_dafny.Zero).Minus((_this).F0()), (_967_v85).F0(), (_this).F0(), (_967_v85).F0()), globalState), _971_v89, (func() _dafny.Int {
					if !(_849_v0) {
						return Companion_Default___.Fm3(_972_v90, _971_v89, (_973_v91).Cardinality(), (_967_v85).F0(), globalState)
					}
					return (func() _dafny.Int {
						if (_974_v92).Contains((_this).F0()) {
							return (_974_v92).Multiplicity((_this).F0())
						}
						return (_this).F0()
					})()
				})(), (_dafny.Zero).Minus((_850_v1).Cardinality()), globalState))
				_ = _rhs146
				var _rhs147 bool = ((_850_v1).Cardinality()).Cmp((_dafny.IntOfInt64(959)).Minus(Companion_Default___.Fm3(_972_v90, _971_v89, _dafny.IntOfInt64(999), _dafny.IntOfUint32((_965_v84).Cardinality()), globalState))) < 0
				_ = _rhs147
				var _rhs148 bool = (_dafny.MultiSetOf(_dafny.IntOfUint32((_968_v86).Cardinality()), _dafny.IntOfInt64(191), p1)).IsDisjointFrom(((_974_v92).Update((_967_v85).F0(), Companion_Default___.Abs((_967_v85).F0()))).Intersection(Companion_Default___.Fm22(globalState)))
				_ = _rhs148
				var _rhs149 _dafny.Int = (p1).Times((_967_v85).F0())
				_ = _rhs149
				var _rhs150 T0 = _967_v85
				_ = _rhs150
				var _lhs66 _dafny.Array = _962_v83
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _lhs67
				r2 = _rhs146
				_849_v0 = _rhs147
				(_lhs66).ArraySet1(_rhs148, (_lhs67).Int())
				r1 = _rhs149
				_967_v85 = _rhs150
				var _975_v93 _dafny.Map
				_ = _975_v93
				_975_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if _849_v0 {
						return Companion_Default___.Fm17(p1, globalState)
					}
					return _971_v89
				})(), (_967_v85).F0())
				_975_v93 = (func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter26 := _dafny.Iterate(((_dafny.MultiSetOf(_dafny.CodePoint('e'))).Update(_971_v89, Companion_Default___.Abs(_dafny.IntOfUint32((_965_v84).Cardinality())))).Elements()); ; {
						_compr_24, _ok26 := _iter26()
						if !_ok26 {
							break
						}
						var _976_v94 _dafny.CodePoint
						_976_v94 = interface{}(_compr_24).(_dafny.CodePoint)
						if ((_dafny.MultiSetOf(_dafny.CodePoint('e'))).Update(_971_v89, Companion_Default___.Abs(_dafny.IntOfUint32((_965_v84).Cardinality())))).Contains(_976_v94) {
							_coll24.Add(_976_v94, (_967_v85).F0())
						}
					}
					return _coll24.ToMap()
				}()).Merge((_975_v93).Merge(_975_v93))
				var _977_v95 T1
				_ = _977_v95
				var _nw166 *C2 = New_C2_()
				_ = _nw166
				_nw166.Ctor__((_967_v85).F0())
				_977_v95 = _nw166
				var _978_v96 _dafny.Map
				_ = _978_v96
				_978_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_977_v95, _965_v84)
				_965_v84 = (func() _dafny.Sequence {
					if (_978_v96).Contains(_977_v95) {
						return (_978_v96).Get(_977_v95).(_dafny.Sequence)
					}
					return _965_v84
				})()
				var _979_v97 _dafny.Sequence
				_ = _979_v97
				_979_v97 = _dafny.SeqOf((_977_v95).F0())
				var _980_v98 D3
				_ = _980_v98
				_980_v98 = Companion_D3_.Create_DC7_(_979_v97)
				var _981_v99 _dafny.Set
				_ = _981_v99
				_981_v99 = _dafny.SetOf(_962_v83, _962_v83)
				var _982_v100 _dafny.Array
				_ = _982_v100
				var _nwElement0_30 *C3 = _937_v64
				_ = _nwElement0_30
				var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(21))
				_ = _nw167
				(_nw167).ArraySet1(_nwElement0_30, 0)
				(_nw167).ArraySet1(_937_v64, 1)
				(_nw167).ArraySet1(_937_v64, 2)
				(_nw167).ArraySet1(_937_v64, 3)
				(_nw167).ArraySet1(_937_v64, 4)
				(_nw167).ArraySet1(_937_v64, 5)
				(_nw167).ArraySet1(_937_v64, 6)
				(_nw167).ArraySet1(_937_v64, 7)
				(_nw167).ArraySet1(_937_v64, 8)
				(_nw167).ArraySet1(_937_v64, 9)
				(_nw167).ArraySet1(_937_v64, 10)
				(_nw167).ArraySet1(_937_v64, 11)
				(_nw167).ArraySet1(_937_v64, 12)
				(_nw167).ArraySet1(_937_v64, 13)
				(_nw167).ArraySet1(_937_v64, 14)
				(_nw167).ArraySet1(_937_v64, 15)
				(_nw167).ArraySet1(_937_v64, 16)
				(_nw167).ArraySet1(_937_v64, 17)
				(_nw167).ArraySet1(_937_v64, 18)
				(_nw167).ArraySet1(_937_v64, 19)
				(_nw167).ArraySet1(_937_v64, 20)
				_982_v100 = _nw167
				var _983_v101 _dafny.Map
				_ = _983_v101
				_983_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_982_v100, !(_849_v0))
				var _984_v102 _dafny.Set
				_ = _984_v102
				_984_v102 = _dafny.SetOf(_937_v64)
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _index159
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _index160
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _index161
				var _rhs151 bool = (_962_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))).Int()).(bool)
				_ = _rhs151
				var _rhs152 D3 = _980_v98
				_ = _rhs152
				var _rhs153 bool = !((_981_v99).IsProperSubsetOf(_981_v99)) || (_849_v0)
				_ = _rhs153
				var _rhs154 bool = (func() bool {
					if (_983_v101).Contains(_982_v100) {
						return (_983_v101).Get(_982_v100).(bool)
					}
					return (_984_v102).IsSubsetOf(_984_v102)
				})()
				_ = _rhs154
				var _lhs68 _dafny.Array = _962_v83
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _lhs69
				var _lhs70 _dafny.Array = _962_v83
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _lhs71
				var _lhs72 _dafny.Array = _962_v83
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))
				_ = _lhs73
				(_lhs68).ArraySet1(_rhs151, (_lhs69).Int())
				_980_v98 = _rhs152
				(_lhs70).ArraySet1(_rhs153, (_lhs71).Int())
				(_lhs72).ArraySet1(_rhs154, (_lhs73).Int())
				var _rhs155 bool = _849_v0
				_ = _rhs155
				var _rhs156 _dafny.Int = ((_dafny.MultiSetOf(_965_v84)).Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_965_v84, _965_v84), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_965_v84, _965_v84)).Cardinality()))).Uint32(), _971_v89), Companion_Default___.Abs(((_977_v95).F0()).Plus((_this).F0())))).Cardinality()
				_ = _rhs156
				var _rhs157 _dafny.Sequence = _dafny.SeqOf(_849_v0, (_962_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))).Int()).(bool), (_962_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_962_v83), 0))).Int()).(bool))
				_ = _rhs157
				_849_v0 = _rhs155
				r2 = _rhs156
				_968_v86 = _rhs157
			}
			var _985_v103 D2
			_ = _985_v103
			_985_v103 = Companion_D2_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-137))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}(func(_986_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), (p1).Cmp(p1) < 0)
			_985_v103 = _985_v103
			if (_849_v0) && (_849_v0) {
				var _987_v104 _dafny.Array
				_ = _987_v104
				var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw168
				_987_v104 = _nw168
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_987_v104), 0))
				_ = _index162
				(_987_v104).ArraySet1(_dafny.IntOfInt64(-560), (_index162).Int())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_987_v104), 0))
				_ = _index163
				(_987_v104).ArraySet1((_dafny.Zero).Minus(p1), (_index163).Int())
				var _988_v105 _dafny.CodePoint
				_ = _988_v105
				_988_v105 = _dafny.CodePoint('n')
				var _989_v106 _dafny.Map
				_ = _989_v106
				_989_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, _988_v105)
				_989_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((!(!(_849_v0))) == (_849_v0)), _988_v105)
				_849_v0 = _849_v0
				var _990_v107 _dafny.Sequence
				_ = _990_v107
				_990_v107 = _dafny.UnicodeSeqOfUtf8Bytes("bcmmxy")
				_990_v107 = _990_v107
				_849_v0 = !(_849_v0) || (_849_v0)
			} else {
				var _991_v108 _dafny.Array
				_ = _991_v108
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw169
				_991_v108 = _nw169
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_991_v108), 0))
				_ = _index164
				(_991_v108).ArraySet1(_849_v0, (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_991_v108), 0))
				_ = _index165
				(_991_v108).ArraySet1(_849_v0, (_index165).Int())
				var _992_v109 _dafny.Array
				_ = _992_v109
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_21
				var _nw170 _dafny.Array
				_ = _nw170
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw170 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = (func(_993_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_994_i8 _dafny.Int) _dafny.Int {
							return (_994_i8).Times(_993_p1)
						}
					})(p1)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw170 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw170).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw170).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_992_v109 = _nw170
				var _995_v110 _dafny.Sequence
				_ = _995_v110
				_995_v110 = _dafny.SeqOf(_992_v109)
				_992_v109 = (_995_v110).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_995_v110).Cardinality()))).Uint32()).(_dafny.Array)
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_991_v108), 0))
				_ = _index166
				(_991_v108).ArraySet1(((func() _dafny.Int {
					if !(_849_v0) {
						return _dafny.IntOfInt64(-114)
					}
					return (_this).F0()
				})()).Cmp(_dafny.IntOfInt64(454)) > 0, (_index166).Int())
				r0 = (_this).F0()
				var _996_v111 _dafny.Array
				_ = _996_v111
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw171
				_996_v111 = _nw171
				var _997_v112 _dafny.Array
				_ = _997_v112
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_22
				var _nw172 _dafny.Array
				_ = _nw172
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw172 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) D0 = (func(_998_v108 _dafny.Array) func(_dafny.Int) D0 {
						return func(_999_i9 _dafny.Int) D0 {
							return Companion_D0_.Create_DC1_((_998_v108).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_998_v108), 0))).Int()).(bool))
						}
					})(_991_v108)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw172 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw172).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw172).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_997_v112 = _nw172
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_996_v111), 0))
				_ = _index167
				(_996_v111).ArraySet1(_997_v112, (_index167).Int())
				var _1000_v113 _dafny.Array
				_ = _1000_v113
				_1000_v113 = _997_v112
				var _1001_v114 _dafny.Sequence
				_ = _1001_v114
				_1001_v114 = _dafny.SeqOf(_997_v112, _997_v112, (_1000_v113))
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_996_v111), 0))
				_ = _index168
				(_996_v111).ArraySet1((_1001_v114).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_1001_v114).Cardinality()))).Uint32()).(_dafny.Array), (_index168).Int())
			}
			var _1002_v115 _dafny.Map
			_ = _1002_v115
			_1002_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, true)
			var _1003_v116 _dafny.Sequence
			_ = _1003_v116
			_1003_v116 = _dafny.SeqOf(_849_v0)
			var _1004_v117 _dafny.Sequence
			_ = _1004_v117
			_1004_v117 = _dafny.SeqOf((_1002_v115).Cardinality(), (_this).F0())
			if (func() bool {
				if (_1002_v115).Contains(!((_1003_v116).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1004_v117).Cardinality()), _dafny.IntOfUint32((_1003_v116).Cardinality()))).Uint32()).(bool))) {
					return (_1002_v115).Get(!((_1003_v116).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1004_v117).Cardinality()), _dafny.IntOfUint32((_1003_v116).Cardinality()))).Uint32()).(bool))).(bool)
				}
				return false
			})() {
				_849_v0 = _849_v0
				_849_v0 = _849_v0
				_849_v0 = _849_v0
				var _1005_v118 _dafny.Sequence
				_ = _1005_v118
				_1005_v118 = _dafny.UnicodeSeqOfUtf8Bytes("qp")
				var _1006_v119 D0
				_ = _1006_v119
				_1006_v119 = Companion_D0_.Create_DC0_(false)
				var _1007_v120 _dafny.CodePoint
				_ = _1007_v120
				_1007_v120 = _dafny.CodePoint('r')
				_1005_v118 = _dafny.Companion_Sequence_.Update(_1005_v118, (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((Companion_D3_.Create_DC8_((_1006_v119).Dtor_cf0(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_850_v1).Contains(_849_v0) {
						return (_850_v1).Multiplicity(_849_v0)
					}
					return p1
				})(), (_this).F0())).Cardinality())).Dtor_cf9(), _dafny.IntOfUint32((_1005_v118).Cardinality())), _dafny.IntOfUint32((_1005_v118).Cardinality()))).Uint32(), _1007_v120)
				var _1008_v121 bool
				_ = _1008_v121
				var _1009_v122 _dafny.Int
				_ = _1009_v122
				var _out19 bool
				_ = _out19
				var _out20 _dafny.Int
				_ = _out20
				_out19, _out20 = (_937_v64).M6(Companion_Default___.SafeModuloInt(p1, (_this).F0()), globalState)
				_1008_v121 = _out19
				_1009_v122 = _out20
			} else {
				var _1010_v123 _dafny.Array
				_ = _1010_v123
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_23
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.CodePoint = func(_1011_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw173 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw173).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw173).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1010_v123 = _nw173
				var _1012_v124 _dafny.CodePoint
				_ = _1012_v124
				_1012_v124 = _dafny.CodePoint('l')
				var _1013_v125 _dafny.CodePoint
				_ = _1013_v125
				_1013_v125 = _1012_v124
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_1010_v123), 0))
				_ = _index169
				(_1010_v123).ArraySet1(_1013_v125, (_index169).Int())
				var _1014_v126 _dafny.Set
				_ = _1014_v126
				_1014_v126 = _dafny.SetOf((_this).F0(), (_this).F0())
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_1010_v123), 0))
				_ = _index170
				var _rhs158 _dafny.CodePoint = _1013_v125
				_ = _rhs158
				var _rhs159 _dafny.Int = ((((_850_v1).Update(_849_v0, Companion_Default___.Abs((_1004_v117).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1004_v117).Cardinality()))).Uint32()).(_dafny.Int)))).Update(_849_v0, Companion_Default___.Abs(_dafny.IntOfUint32((_1004_v117).Cardinality())))).Cardinality()).Times(p1)
				_ = _rhs159
				var _rhs160 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1003_v116, _1003_v116)).Cardinality())).Cmp((_this).F0()) > 0
				_ = _rhs160
				var _rhs161 _dafny.Sequence = _1004_v117
				_ = _rhs161
				var _rhs162 _dafny.Set = (_1014_v126).Intersection((_dafny.SetOf((_this).F0())).Difference(_1014_v126))
				_ = _rhs162
				var _lhs74 _dafny.Array = _1010_v123
				_ = _lhs74
				var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_1010_v123), 0))
				_ = _lhs75
				(_lhs74).ArraySet1(_rhs158, (_lhs75).Int())
				r0 = _rhs159
				_849_v0 = _rhs160
				_1004_v117 = _rhs161
				_1014_v126 = _rhs162
				var _1015_v127 _dafny.Array
				_ = _1015_v127
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
				_ = _nw174
				_1015_v127 = _nw174
				var _1016_v128 _dafny.Map
				_ = _1016_v128
				_1016_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-949), p1)
				var _1017_v129 D7
				_ = _1017_v129
				_1017_v129 = Companion_D7_.Create_DC17_(_849_v0, _849_v0, _849_v0)
				var _pat_let_tv27 = _849_v0
				_ = _pat_let_tv27
				var _1018_v130 _dafny.Map
				_ = _1018_v130
				_1018_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let21_0 D7) D7 {
					return func(_1019_dt__update__tmp_h4 D7) D7 {
						return func(_pat_let22_0 bool) D7 {
							return func(_1020_dt__update_hcf27_h0 bool) D7 {
								return Companion_D7_.Create_DC17_(_1020_dt__update_hcf27_h0, (_1019_dt__update__tmp_h4).Dtor_cf28(), (_1019_dt__update__tmp_h4).Dtor_cf29())
							}(_pat_let22_0)
						}(_pat_let_tv27)
					}(_pat_let21_0)
				}(_1017_v129), (p1).Plus((func() _dafny.Int {
					if (_1016_v128).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("puclwjroo")).Cardinality())) {
						return (_1016_v128).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("puclwjroo")).Cardinality())).(_dafny.Int)
					}
					return _dafny.IntOfInt64(967)
				})()))
				var _1021_v131 _dafny.Sequence
				_ = _1021_v131
				_1021_v131 = _dafny.UnicodeSeqOfUtf8Bytes("msay")
				var _1022_v132 _dafny.Map
				_ = _1022_v132
				_1022_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v0, _1021_v131)
				var _1023_v133 _dafny.Set
				_ = _1023_v133
				_1023_v133 = _dafny.SetOf(_849_v0, _849_v0)
				var _rhs163 _dafny.Array = _1015_v127
				_ = _rhs163
				var _rhs164 _dafny.Map = _1016_v128
				_ = _rhs164
				var _rhs165 _dafny.Int = ((func() _dafny.Int {
					if _849_v0 {
						return p1
					}
					return (_this).F0()
				})()).Plus(p1)
				_ = _rhs165
				var _rhs166 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1017_v129, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_1022_v132).Contains(_849_v0) {
						return (_1022_v132).Get(_849_v0).(_dafny.Sequence)
					}
					return _1021_v131
				})(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1022_v132).Contains(_849_v0) {
						return (_1022_v132).Get(_849_v0).(_dafny.Sequence)
					}
					return _1021_v131
				})()).Cardinality()))).Uint32(), _1012_v124)).Cardinality()))
				_ = _rhs166
				var _rhs167 _dafny.Int = Companion_Default___.Fm3(_1023_v133, _1012_v124, p1, p1, globalState)
				_ = _rhs167
				_1015_v127 = _rhs163
				_1016_v128 = _rhs164
				r2 = _rhs165
				_1018_v130 = _rhs166
				r1 = _rhs167
				var _1024_v134 *C2
				_ = _1024_v134
				var _nw175 *C2 = New_C2_()
				_ = _nw175
				_nw175.Ctor__((_this).F0())
				_1024_v134 = _nw175
				var _1025_v135 _dafny.Sequence
				_ = _1025_v135
				_1025_v135 = _dafny.SeqOf(_1024_v134, _1024_v134)
				var _1026_v136 _dafny.Array
				_ = _1026_v136
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw176
				_1026_v136 = _nw176
				var _rhs168 _dafny.Sequence = (func() _dafny.Sequence {
					if _849_v0 {
						return (_985_v103).Dtor_cf4()
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("hmljar")
				})()
				_ = _rhs168
				var _rhs169 bool = (_1003_v116).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1003_v116).Cardinality()))).Uint32()).(bool)
				_ = _rhs169
				var _rhs170 _dafny.Sequence = _1025_v135
				_ = _rhs170
				var _rhs171 _dafny.Array = _1026_v136
				_ = _rhs171
				var _rhs172 bool = _849_v0
				_ = _rhs172
				_1021_v131 = _rhs168
				_849_v0 = _rhs169
				_1025_v135 = _rhs170
				_1026_v136 = _rhs171
				_849_v0 = _rhs172
				_1014_v126 = _1014_v126
				_849_v0 = _849_v0
			}
		}
		var _1027_v137 _dafny.Sequence
		_ = _1027_v137
		_1027_v137 = _dafny.SeqOf(p1)
		_1027_v137 = _1027_v137
		if (func() bool {
			if _849_v0 {
				return _849_v0
			}
			return _849_v0
		})() {
			var _1028_v138 _dafny.Set
			_ = _1028_v138
			_1028_v138 = _dafny.SetOf(false, _849_v0)
			var _1029_v139 *C3
			_ = _1029_v139
			var _nw177 *C3 = New_C3_()
			_ = _nw177
			_nw177.Ctor__((_dafny.Zero).Minus(((_1028_v138).Intersection(_1028_v138)).Cardinality()))
			_1029_v139 = _nw177
			var _1030_v140 _dafny.CodePoint
			_ = _1030_v140
			_1030_v140 = _dafny.CodePoint('v')
			var _1031_v141 *C0
			_ = _1031_v141
			var _nw178 *C0 = New_C0_()
			_ = _nw178
			_nw178.Ctor__(_1030_v140)
			_1031_v141 = _nw178
			if _849_v0 {
				var _1032_v142 _dafny.MultiSet
				_ = _1032_v142
				_1032_v142 = _dafny.MultiSetOf((_this).F0())
				var _1033_v143 _dafny.Map
				_ = _1033_v143
				_1033_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(_1032_v142, p1, (_this).F0(), (_this).F0(), globalState), true)
				_1033_v143 = (_1033_v143).Update(Companion_Default___.Fm19(_1032_v142, p1, _dafny.IntOfUint32((_1027_v137).Cardinality()), (_this).F0(), globalState), !(_849_v0))
				var _1034_v144 _dafny.Sequence
				_ = _1034_v144
				_1034_v144 = _dafny.UnicodeSeqOfUtf8Bytes("jwm")
				_1034_v144 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(9))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1035_v140 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1036_i11 _dafny.Int) _dafny.CodePoint {
						return _1035_v140
					}
				})(_1030_v140)))
				var _1037_v145 *C1
				_ = _1037_v145
				var _nw179 *C1 = New_C1_()
				_ = _nw179
				_nw179.Ctor__((_this).F0())
				_1037_v145 = _nw179
				var _1038_v146 *C1
				_ = _1038_v146
				var _nw180 *C1 = New_C1_()
				_ = _nw180
				_nw180.Ctor__((_dafny.Zero).Minus((_this).F0()))
				_1038_v146 = _nw180
				_849_v0 = !(_dafny.MultiSetOf(_849_v0, _849_v0, _849_v0, _849_v0)).Equals((_850_v1).Difference(_850_v1))
			} else {
				var _1039_v147 _dafny.Sequence
				_ = _1039_v147
				_1039_v147 = _dafny.UnicodeSeqOfUtf8Bytes("iltk")
				_1039_v147 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(310))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_1040_v140 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1041_i12 _dafny.Int) _dafny.CodePoint {
						return _1040_v140
					}
				})(_1030_v140))), _1039_v147), _dafny.UnicodeSeqOfUtf8Bytes("gclpd"))
				var _1042_v148 _dafny.Array
				_ = _1042_v148
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw181
				_1042_v148 = _nw181
				var _1043_v149 _dafny.Map
				_ = _1043_v149
				_1043_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_849_v0) && (_849_v0), _1042_v148)
				_1043_v149 = (_1043_v149).Update((Companion_D3_.Create_DC8_(_849_v0, p1)).Dtor_cf8(), _1042_v148)
				r0 = (p1).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1030_v140, (_1031_v141).F3()), _1039_v147)).Cardinality()))
				var _1044_v150 *C1
				_ = _1044_v150
				var _nw182 *C1 = New_C1_()
				_ = _nw182
				_nw182.Ctor__((_this).F0())
				_1044_v150 = _nw182
				r2 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F0()), (_this).F0())
			}
			_849_v0 = _849_v0
			_849_v0 = _849_v0
		} else {
			_849_v0 = _849_v0
			_849_v0 = _849_v0
			var _1045_v151 _dafny.Sequence
			_ = _1045_v151
			_1045_v151 = _dafny.UnicodeSeqOfUtf8Bytes("nny")
			var _1046_v152 _dafny.CodePoint
			_ = _1046_v152
			_1046_v152 = _dafny.CodePoint('d')
			var _1047_v153 _dafny.Set
			_ = _1047_v153
			_1047_v153 = _dafny.SetOf(_849_v0, _849_v0)
			var _1048_v154 _dafny.MultiSet
			_ = _1048_v154
			_1048_v154 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1027_v137).Cardinality()), p1)
			var _1049_v155 D2
			_ = _1049_v155
			_1049_v155 = Companion_D2_.Create_DC4_(_dafny.Companion_Sequence_.Update(_1045_v151, (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1045_v151).Cardinality()))).Uint32(), _1046_v152), ((_this).F0()).Cmp(Companion_Default___.Fm3(_1047_v153, _1046_v152, (_this).F0(), (_1048_v154).Cardinality(), globalState)) <= 0)
			_1049_v155 = _1049_v155
			var _1050_v156 *C2
			_ = _1050_v156
			var _nw183 *C2 = New_C2_()
			_ = _nw183
			_nw183.Ctor__(p1)
			_1050_v156 = _nw183
			var _nw184 *C2 = New_C2_()
			_ = _nw184
			_nw184.Ctor__((p1).Times(p1))
			_1050_v156 = _nw184
			_1027_v137 = _1027_v137
		}
		var _1051_v157 _dafny.Set
		_ = _1051_v157
		_1051_v157 = _dafny.SetOf(true)
		var _1052_v158 _dafny.Sequence
		_ = _1052_v158
		_1052_v158 = _dafny.UnicodeSeqOfUtf8Bytes("kw")
		var _hi6 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1051_v157, _dafny.IntOfUint32((_1052_v158).Cardinality()))).Cardinality(), p1)
		_ = _hi6
		for _1053_i13 := (_1027_v137).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1027_v137).Cardinality()))).Uint32()).(_dafny.Int); _1053_i13.Cmp(_hi6) < 0; _1053_i13 = _1053_i13.Plus(_dafny.One) {
			_849_v0 = _849_v0
			r0 = Companion_Default___.SafeModuloInt((_this).F0(), (_1051_v157).Cardinality())
			_849_v0 = !((_dafny.IntOfUint32((_1027_v137).Cardinality())).Cmp((_dafny.MultiSetOf(p1, (_this).F0())).Cardinality()) > 0)
			r2 = (_this).F0()
		}
		var _1054_v159 _dafny.Sequence
		_ = _1054_v159
		_1054_v159 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1052_v158, _dafny.UnicodeSeqOfUtf8Bytes("j")))
		_1054_v159 = _1054_v159
		var _1055_v160 _dafny.CodePoint
		_ = _1055_v160
		_1055_v160 = _dafny.CodePoint('s')
		var _1056_v161 _dafny.Set
		_ = _1056_v161
		_1056_v161 = _dafny.SetOf(_1055_v160, _1055_v160, _1055_v160, _dafny.CodePoint('g'), _dafny.CodePoint('h'))
		r0 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(924)).Plus(p1), Companion_Default___.SafeModuloInt((_this).F0(), (_1056_v161).Cardinality()))
		r1 = (_this).F0()
		var _1057_v162 D0
		_ = _1057_v162
		_1057_v162 = Companion_D0_.Create_DC1_(_849_v0)
		var _1058_v163 _dafny.Sequence
		_ = _1058_v163
		_1058_v163 = _dafny.SeqOf(_1057_v162)
		var _1059_v164 D2
		_ = _1059_v164
		_1059_v164 = Companion_D2_.Create_DC3_(_1058_v163)
		var _pat_let_tv28 = _1052_v158
		_ = _pat_let_tv28
		var _pat_let_tv29 = p1
		_ = _pat_let_tv29
		var _pat_let_tv30 = p1
		_ = _pat_let_tv30
		r2 = func(_source12 D2) _dafny.Int {
			if _source12.Is_DC4() {
				var _1060___mcc_h6 _dafny.Sequence = _source12.Get_().(D2_DC4).Cf4
				_ = _1060___mcc_h6
				var _1061___mcc_h7 bool = _source12.Get_().(D2_DC4).Cf5
				_ = _1061___mcc_h7
				var _1062_cf5 bool = _1061___mcc_h7
				_ = _1062_cf5
				var _1063_cf4 _dafny.Sequence = _1060___mcc_h6
				_ = _1063_cf4
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC4_(_pat_let_tv28, false)).Dtor_cf5(), (_dafny.Zero).Minus(_pat_let_tv29))).Cardinality()
			} else if _source12.Is_DC5() {
				return _pat_let_tv30
			} else if _source12.Is_DC3() {
				var _1064___mcc_h8 _dafny.Sequence = _source12.Get_().(D2_DC3).Cf3
				_ = _1064___mcc_h8
				var _1065_cf3 _dafny.Sequence = _1064___mcc_h8
				_ = _1065_cf3
				return Companion_Default___.SafeModuloInt((_this).F0(), (_this).F0())
			} else {
				var _1066___mcc_h9 D2 = _source12.Get_().(D2_DC6).Cf6
				_ = _1066___mcc_h9
				var _1067_cf6 D2 = _1066___mcc_h9
				_ = _1067_cf6
				return (_this).F0()
			}
		}(_1059_v164)
		return r0, r1, r2
	}
}
func (_this *C4) M4(globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1068_v0 _dafny.Array
		_ = _1068_v0
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_24
		var _nw185 _dafny.Array
		_ = _nw185
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw185 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) _dafny.Sequence = func(_1069_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("kotlm")
			}
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw185 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw185).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw185).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_1068_v0 = _nw185
		var _1070_v1 _dafny.Sequence
		_ = _1070_v1
		_1070_v1 = _dafny.UnicodeSeqOfUtf8Bytes("pxch")
		var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1068_v0), 0))
		_ = _index171
		(_1068_v0).ArraySet1(_1070_v1, (_index171).Int())
		var _1071_v2 *C2
		_ = _1071_v2
		var _nw186 *C2 = New_C2_()
		_ = _nw186
		_nw186.Ctor__(_dafny.IntOfInt64(282))
		_1071_v2 = _nw186
		var _1072_v3 _dafny.MultiSet
		_ = _1072_v3
		_1072_v3 = _dafny.MultiSetOf(_1071_v2)
		var _1073_v4 bool
		_ = _1073_v4
		_1073_v4 = false
		var _1074_v5 D2
		_ = _1074_v5
		_1074_v5 = Companion_D2_.Create_DC4_(_1070_v1, _1073_v4)
		var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1068_v0), 0))
		_ = _index172
		(_1068_v0).ArraySet1((func() _dafny.Sequence {
			if (_1072_v3).Contains(_1071_v2) {
				return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm18((_this).F0(), globalState), _1070_v1)
			}
			return (_1074_v5).Dtor_cf4()
		})(), (_index172).Int())
		var _hi7 _dafny.Int = ((_this).F0()).Plus(_dafny.IntOfInt64(-45))
		_ = _hi7
		for _1075_i1 := Companion_Default___.SafeDivisionInt((_this).F0(), (_dafny.Zero).Minus((_this).F0())); _1075_i1.Cmp(_hi7) < 0; _1075_i1 = _1075_i1.Plus(_dafny.One) {
			_1070_v1 = (_1068_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1068_v0), 0))).Int()).(_dafny.Sequence)
			var _1076_v6 _dafny.Array
			_ = _1076_v6
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_25
			var _nw187 _dafny.Array
			_ = _nw187
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw187 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_1077_v4 bool) func(_dafny.Int) bool {
					return func(_1078_i2 _dafny.Int) bool {
						return _1077_v4
					}
				})(_1073_v4)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw187 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw187).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw187).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1076_v6 = _nw187
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1076_v6), 0))
			_ = _index173
			(_1076_v6).ArraySet1(_1073_v4, (_index173).Int())
			var _1079_v7 _dafny.Sequence
			_ = _1079_v7
			_1079_v7 = _dafny.SeqOf(_1073_v4)
			var _1080_v8 _dafny.MultiSet
			_ = _1080_v8
			_1080_v8 = _dafny.MultiSetOf(_1073_v4, _1073_v4)
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1076_v6), 0))
			_ = _index174
			(_1076_v6).ArraySet1((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1079_v7, _1079_v7), (Companion_Default___.SafeIndex(_1075_i1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1079_v7, _1079_v7)).Cardinality()))).Uint32(), _1073_v4))).IsDisjointFrom((func() _dafny.MultiSet {
				if _1073_v4 {
					return _1080_v8
				}
				return _1080_v8
			})()), (_index174).Int())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1076_v6), 0))
			_ = _index175
			(_1076_v6).ArraySet1((_1071_v2).Fm5(globalState), (_index175).Int())
			_1073_v4 = !(_1073_v4)
		}
		var _1081_v9 _dafny.MultiSet
		_ = _1081_v9
		_1081_v9 = _dafny.MultiSetOf((_this).F0())
		var _1082_v10 _dafny.Sequence
		_ = _1082_v10
		_1082_v10 = _dafny.SeqOf(_1073_v4)
		var _1083_v11 _dafny.Map
		_ = _1083_v11
		_1083_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1082_v10).Cardinality()), !(_1073_v4)), _1073_v4)
		var _1084_v12 _dafny.Map
		_ = _1084_v12
		_1084_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(904), false)
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1068_v0), 0))
		_ = _index176
		(_1068_v0).ArraySet1(Companion_Default___.Fm19(_1081_v9, (_this).F0(), ((_1083_v11).Cardinality()).Plus((_this).F0()), (_1084_v12).Cardinality(), globalState), (_index176).Int())
		if (_1071_v2).Fm6(_1073_v4, _1070_v1, _1070_v1, ((_this).F0()).Plus((_this).F0()), globalState) {
			var _1085_v13 _dafny.Int
			_ = _1085_v13
			_1085_v13 = _dafny.IntOfInt64(47)
			_1085_v13 = _1085_v13
			if _1073_v4 {
				var _1086_v14 _dafny.Array
				_ = _1086_v14
				var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
				_ = _nw188
				_1086_v14 = _nw188
				var _1087_v15 _dafny.Set
				_ = _1087_v15
				_1087_v15 = _dafny.SetOf(_dafny.IntOfInt64(636))
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1086_v14), 0))
				_ = _index177
				(_1086_v14).ArraySet1(_1087_v15, (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1086_v14), 0))
				_ = _index178
				(_1086_v14).ArraySet1((_1087_v15).Intersection((func() _dafny.Set {
					if _1073_v4 {
						return _1087_v15
					}
					return _1087_v15
				})()), (_index178).Int())
				var _1088_v16 *C3
				_ = _1088_v16
				var _nw189 *C3 = New_C3_()
				_ = _nw189
				_nw189.Ctor__(_1085_v13)
				_1088_v16 = _nw189
				var _1089_v17 _dafny.Set
				_ = _1089_v17
				_1089_v17 = _dafny.SetOf(_dafny.SetOf(_1073_v4, _1073_v4))
				var _1090_v18 _dafny.Map
				_ = _1090_v18
				_1090_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1088_v16, (_1089_v17).Union(_1089_v17))
				_1090_v18 = (_1090_v18).Update(_1088_v16, _1089_v17)
				_1073_v4 = (_1073_v4) || (!(_1073_v4))
				_1073_v4 = _1073_v4
				_1085_v13 = _dafny.IntOfInt64(976)
			} else {
				_1073_v4 = _1073_v4
				var _1091_v19 _dafny.Map
				_ = _1091_v19
				_1091_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v4, (_dafny.Zero).Minus(_1085_v13))
				var _1092_v20 _dafny.Map
				_ = _1092_v20
				_1092_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v4, _1091_v19)
				_1092_v20 = (_1092_v20).Update(_1073_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v4, _dafny.IntOfUint32((_1070_v1).Cardinality())))
				var _1093_v21 *C3
				_ = _1093_v21
				var _nw190 *C3 = New_C3_()
				_ = _nw190
				_nw190.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1070_v1, _1070_v1)).Cardinality()))
				_1093_v21 = _nw190
				var _1094_v22 _dafny.CodePoint
				_ = _1094_v22
				_1094_v22 = _dafny.CodePoint('x')
				var _1095_v23 _dafny.Map
				_ = _1095_v23
				_1095_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1085_v13, _1094_v22)
				var _1096_v24 _dafny.MultiSet
				_ = _1096_v24
				_1096_v24 = _dafny.MultiSetOf(_1095_v23, _1095_v23, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _1094_v22), _1095_v23, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), _1094_v22))
				var _1097_v25 _dafny.Sequence
				_ = _1097_v25
				_1097_v25 = _dafny.SeqOf(_1085_v13, (_dafny.Zero).Minus(_1085_v13), _1085_v13, _1085_v13, _1085_v13)
				var _1098_v26 _dafny.Map
				_ = _1098_v26
				_1098_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1096_v24).Union(_1096_v24), (_1097_v25).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.IntOfUint32((_1097_v25).Cardinality()))).Uint32()).(_dafny.Int))
				_1098_v26 = (_1098_v26).Update((_1096_v24).Difference(_dafny.MultiSetOf(_1095_v23)), (_this).F0())
				var _1099_v27 _dafny.Array
				_ = _1099_v27
				var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw191
				_1099_v27 = _nw191
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1099_v27), 0))
				_ = _index179
				(_1099_v27).ArraySet1(_1085_v13, (_index179).Int())
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1099_v27), 0))
				_ = _index180
				(_1099_v27).ArraySet1((_1085_v13).Times((_this).F0()), (_index180).Int())
				var _1100_v28 _dafny.MultiSet
				_ = _1100_v28
				_1100_v28 = _dafny.MultiSetOf(_1073_v4, _1073_v4)
				var _1101_v29 _dafny.Map
				_ = _1101_v29
				_1101_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1094_v22, ((_1100_v28).Update((_1093_v21).Fm5(globalState), Companion_Default___.Abs((_this).F0()))).IsProperSubsetOf(_1100_v28))
				var _1102_v30 _dafny.Sequence
				_ = _1102_v30
				_1102_v30 = _dafny.SeqOf(_dafny.SeqOf(_1073_v4), _1082_v10)
				var _1103_v31 _dafny.Array
				_ = _1103_v31
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw192
				_1103_v31 = _nw192
				var _1104_v32 _dafny.Set
				_ = _1104_v32
				_1104_v32 = _dafny.SetOf(_1103_v31)
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1099_v27), 0))
				_ = _index181
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1099_v27), 0))
				_ = _index182
				var _rhs173 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1082_v10, (_1102_v30).Select((Companion_Default___.SafeIndex((_1100_v28).Cardinality(), _dafny.IntOfUint32((_1102_v30).Cardinality()))).Uint32()).(_dafny.Sequence)), _1082_v10)
				_ = _rhs173
				var _rhs174 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1104_v32).Cardinality(), _1085_v13))
				_ = _rhs174
				var _rhs175 _dafny.Int = (_1085_v13).Plus((_this).F0())
				_ = _rhs175
				var _rhs176 _dafny.Map = _1101_v29
				_ = _rhs176
				var _rhs177 _dafny.Int = (_dafny.Zero).Minus((_1085_v13).Times((_this).F0()))
				_ = _rhs177
				var _lhs76 _dafny.Array = _1099_v27
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1099_v27), 0))
				_ = _lhs77
				var _lhs78 _dafny.Array = _1099_v27
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_1099_v27), 0))
				_ = _lhs79
				_1082_v10 = _rhs173
				(_lhs76).ArraySet1(_rhs174, (_lhs77).Int())
				(_lhs78).ArraySet1(_rhs175, (_lhs79).Int())
				_1101_v29 = _rhs176
				_1085_v13 = _rhs177
			}
			var _rhs178 bool = _1073_v4
			_ = _rhs178
			var _rhs179 bool = _1073_v4
			_ = _rhs179
			var _rhs180 bool = _1073_v4
			_ = _rhs180
			var _rhs181 _dafny.Int = _dafny.IntOfInt64(-963)
			_ = _rhs181
			_1073_v4 = _rhs178
			_1073_v4 = _rhs179
			_1073_v4 = _rhs180
			_1085_v13 = _rhs181
			var _1105_v33 _dafny.CodePoint
			_ = _1105_v33
			_1105_v33 = _dafny.CodePoint('s')
			var _1106_v34 *C0
			_ = _1106_v34
			var _nw193 *C0 = New_C0_()
			_ = _nw193
			_nw193.Ctor__(_1105_v33)
			_1106_v34 = _nw193
			_1106_v34 = _1106_v34
			_1085_v13 = Companion_Default___.SafeDivisionInt((_this).F0(), (_this).F0())
		} else {
			var _1107_v35 _dafny.Map
			_ = _1107_v35
			_1107_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F0()).Cmp((_this).F0()) < 0, (_this).F0())
			_1107_v35 = (_1107_v35).Update((false) || (_1073_v4), (_this).F0())
			var _1108_v36 _dafny.Array
			_ = _1108_v36
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw194
			_1108_v36 = _nw194
			var _1109_v37 _dafny.Int
			_ = _1109_v37
			var _1110_v38 _dafny.Int
			_ = _1110_v38
			var _1111_v39 _dafny.Int
			_ = _1111_v39
			var _out21 _dafny.Int
			_ = _out21
			var _out22 _dafny.Int
			_ = _out22
			var _out23 _dafny.Int
			_ = _out23
			_out21, _out22, _out23 = (_1071_v2).M1(_1108_v36, (_this).F0(), globalState)
			_1109_v37 = _out21
			_1110_v38 = _out22
			_1111_v39 = _out23
			var _1112_v40 _dafny.CodePoint
			_ = _1112_v40
			_1112_v40 = _dafny.CodePoint('y')
			_1112_v40 = _1112_v40
			var _1113_v41 _dafny.Array
			_ = _1113_v41
			var _nw195 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw195
			_1113_v41 = _nw195
			var _1114_v42 T0
			_ = _1114_v42
			var _nw196 *C3 = New_C3_()
			_ = _nw196
			_nw196.Ctor__((_this).F0())
			_1114_v42 = _nw196
			var _1115_v43 D10
			_ = _1115_v43
			_1115_v43 = Companion_D10_.Create_DC21_(_1114_v42)
			var _1116_v44 _dafny.Array
			_ = _1116_v44
			var _nwElement0_31 bool = _1073_v4
			_ = _nwElement0_31
			var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(7))
			_ = _nw197
			(_nw197).ArraySet1(_nwElement0_31, 0)
			(_nw197).ArraySet1(_1073_v4, 1)
			(_nw197).ArraySet1(_1073_v4, 2)
			(_nw197).ArraySet1(false, 3)
			(_nw197).ArraySet1(!(false), 4)
			(_nw197).ArraySet1(_1073_v4, 5)
			(_nw197).ArraySet1(_1073_v4, 6)
			_1116_v44 = _nw197
			var _1117_v45 _dafny.Map
			_ = _1117_v45
			_1117_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v43, _1116_v44)
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1113_v41), 0))
			_ = _index183
			(_1113_v41).ArraySet1(_1117_v45, (_index183).Int())
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1113_v41), 0))
			_ = _index184
			(_1113_v41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v43, _1116_v44), (_index184).Int())
			var _1118_v46 _dafny.Array
			_ = _1118_v46
			var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw198
			_1118_v46 = _nw198
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1118_v46), 0))
			_ = _index185
			(_1118_v46).ArraySet1((_dafny.Zero).Minus(_1111_v39), (_index185).Int())
			var _1119_v47 _dafny.Map
			_ = _1119_v47
			_1119_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v4, _1073_v4)
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1118_v46), 0))
			_ = _index186
			(_1118_v46).ArraySet1(((_dafny.Zero).Minus(((func() _dafny.Map {
				if !(_1073_v4) {
					return _1119_v47
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v4, _1073_v4)
			})()).Cardinality())).Times(_1110_v38), (_index186).Int())
		}
		var _1120_v48 _dafny.Int
		_ = _1120_v48
		_1120_v48 = _dafny.IntOfInt64(-963)
		_1120_v48 = (_dafny.Zero).Minus((_this).F0())
		var _1121_v49 _dafny.CodePoint
		_ = _1121_v49
		_1121_v49 = _dafny.CodePoint('t')
		var _1122_v50 _dafny.Map
		_ = _1122_v50
		_1122_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1121_v49, _1120_v48)
		var _hi8 _dafny.Int = ((_1122_v50).Merge(_1122_v50)).Cardinality()
		_ = _hi8
		for _1123_i3 := (_this).F0(); _1123_i3.Cmp(_hi8) < 0; _1123_i3 = _1123_i3.Plus(_dafny.One) {
			var _1124_v51 *C3
			_ = _1124_v51
			var _nw199 *C3 = New_C3_()
			_ = _nw199
			_nw199.Ctor__(_dafny.IntOfInt64(920))
			_1124_v51 = _nw199
			var _1125_v52 _dafny.Sequence
			_ = _1125_v52
			_1125_v52 = _dafny.SeqOf(_1123_i3)
			var _1126_v53 _dafny.Sequence
			_ = _1126_v53
			_1126_v53 = _dafny.SeqOf(_1125_v52, _1125_v52, _1125_v52, _1125_v52, _1125_v52)
			var _1127_v54 _dafny.Map
			_ = _1127_v54
			_1127_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1126_v53, _dafny.SeqOf(_1120_v48, _dafny.IntOfInt64(847), (_this).F0()))
			_1127_v54 = (_1127_v54).Update(_dafny.Companion_Sequence_.Update(_1126_v53, (Companion_Default___.SafeIndex(_1123_i3, _dafny.IntOfUint32((_1126_v53).Cardinality()))).Uint32(), _1125_v52), _1125_v52)
			_1120_v48 = (_1120_v48).Minus((_this).F0())
			var _1128_v55 _dafny.Array
			_ = _1128_v55
			var _nwElement0_32 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _nwElement0_32
			var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(2))
			_ = _nw200
			(_nw200).ArraySet1(_nwElement0_32, 0)
			(_nw200).ArraySet1((_1123_i3).Times(_1120_v48), 1)
			_1128_v55 = _nw200
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1128_v55), 0))
			_ = _index187
			(_1128_v55).ArraySet1(_1123_i3, (_index187).Int())
			var _1129_v56 _dafny.Set
			_ = _1129_v56
			_1129_v56 = _dafny.SetOf(_1073_v4, _1073_v4)
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1128_v55), 0))
			_ = _index188
			(_1128_v55).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(_1129_v56, _1121_v49, (_this).F0(), _dafny.IntOfUint32((_1082_v10).Cardinality()), globalState), _1120_v48), (_index188).Int())
		}
		var _1130_v57 _dafny.Map
		_ = _1130_v57
		_1130_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v4, _1073_v4)
		var _1131_v58 _dafny.Sequence
		_ = _1131_v58
		_1131_v58 = _dafny.SeqOf(_1130_v57)
		var _1132_v59 _dafny.Sequence
		_ = _1132_v59
		_1132_v59 = _dafny.SeqOf(_1081_v9)
		var _1133_v60 _dafny.Sequence
		_ = _1133_v60
		_1133_v60 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ecqjshdfv")).Cardinality()), _1120_v48)).Cardinality())
		var _1134_v61 _dafny.MultiSet
		_ = _1134_v61
		_1134_v61 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F0())), _dafny.MultiSetOf(_dafny.IntOfUint32((_1133_v60).Cardinality()), _1120_v48))
		r0 = ((_1130_v57).Merge((_1131_v58).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1131_v58).Cardinality()))).Uint32()).(_dafny.Map))).Update(_1073_v4, (_dafny.MultiSetFromSeq(_1132_v59)).IsDisjointFrom(_1134_v61))
		return r0
	}
}
func (_this *C4) M5(p0 bool, p1 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Array, _dafny.MultiSet) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r3
		var _1135_v0 _dafny.Int
		_ = _1135_v0
		_1135_v0 = _dafny.IntOfInt64(565)
		_1135_v0 = (_this).F0()
		var _1136_v1 _dafny.CodePoint
		_ = _1136_v1
		_1136_v1 = _dafny.CodePoint('g')
		_1136_v1 = _1136_v1
		_1135_v0 = _1135_v0
		if p0 {
			var _1137_v2 *C0
			_ = _1137_v2
			var _nw201 *C0 = New_C0_()
			_ = _nw201
			_nw201.Ctor__(_1136_v1)
			_1137_v2 = _nw201
			var _1138_v3 _dafny.Set
			_ = _1138_v3
			_1138_v3 = _dafny.SetOf(p1, p1)
			var _1139_v4 _dafny.Sequence
			_ = _1139_v4
			_1139_v4 = _dafny.SeqOf((_1138_v3).Cardinality())
			var _1140_v5 D0
			_ = _1140_v5
			_1140_v5 = Companion_D0_.Create_DC1_(true)
			var _1141_v6 _dafny.Sequence
			_ = _1141_v6
			_1141_v6 = _dafny.SeqOf(Companion_D0_.Create_DC1_(Companion_Default___.Fm2(p1, _dafny.IntOfUint32((_1139_v4).Cardinality()), p0, p1, globalState)), _1140_v5, _1140_v5)
			var _1142_v7 D2
			_ = _1142_v7
			_1142_v7 = Companion_D2_.Create_DC3_(_1141_v6)
			var _1143_v8 D2
			_ = _1143_v8
			_1143_v8 = Companion_D2_.Create_DC6_((func() D2 {
				if p1 {
					return _1142_v7
				}
				return Companion_D2_.Create_DC6_(_1142_v7)
			})())
			_1143_v8 = _1143_v8
			_1135_v0 = (_1135_v0).Times(_1135_v0)
			_1135_v0 = (Companion_Default___.SafeModuloInt((_this).F0(), (_dafny.Zero).Minus(Companion_Default___.Fm3(_1138_v3, (_1137_v2).F3(), (_1139_v4).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1139_v4).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(759), globalState)))).Plus(_1135_v0)
			var _1144_v9 *C0
			_ = _1144_v9
			var _nw202 *C0 = New_C0_()
			_ = _nw202
			_nw202.Ctor__((_1137_v2).F3())
			_1144_v9 = _nw202
		} else {
			var _1145_v10 _dafny.Array
			_ = _1145_v10
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_26
			var _nw203 _dafny.Array
			_ = _nw203
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw203 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.Int = func(_1146_i0 _dafny.Int) _dafny.Int {
					return (_1146_i0).Plus((_this).F0())
				}
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw203 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw203).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw203).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1145_v10 = _nw203
			var _1147_v11 _dafny.Map
			_ = _1147_v11
			_1147_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(301), (_this).F0())
			var _1148_v12 _dafny.Sequence
			_ = _1148_v12
			_1148_v12 = _dafny.SeqOf(_1136_v1, _1136_v1)
			var _1149_v13 _dafny.Sequence
			_ = _1149_v13
			_1149_v13 = _dafny.SeqOf((_1147_v11).Cardinality(), (_this).F0(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1148_v12).Cardinality())))
			var _1150_v14 _dafny.Sequence
			_ = _1150_v14
			_1150_v14 = _dafny.SeqOf(p1)
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
			_ = _index189
			(_1145_v10).ArraySet1(((_1149_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1150_v14).Cardinality()), _dafny.IntOfUint32((_1149_v13).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_1135_v0), (_index189).Int())
			var _1151_v15 D3
			_ = _1151_v15
			_1151_v15 = Companion_D3_.Create_DC7_(_1149_v13)
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
			_ = _index190
			var _rhs182 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((_1151_v15).Dtor_cf7(), _1149_v13)).Cardinality())
			_ = _rhs182
			var _rhs183 _dafny.Int = _1135_v0
			_ = _rhs183
			var _rhs184 _dafny.Sequence = _1148_v12
			_ = _rhs184
			var _lhs80 _dafny.Array = _1145_v10
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
			_ = _lhs81
			(_lhs80).ArraySet1(_rhs182, (_lhs81).Int())
			_1135_v0 = _rhs183
			_1148_v12 = _rhs184
			var _1152_v16 _dafny.Array
			_ = _1152_v16
			var _nwElement0_33 bool = p1
			_ = _nwElement0_33
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(18))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_33, 0)
			(_nw204).ArraySet1(p0, 1)
			(_nw204).ArraySet1(true, 2)
			(_nw204).ArraySet1(p1, 3)
			(_nw204).ArraySet1(p1, 4)
			(_nw204).ArraySet1(p1, 5)
			(_nw204).ArraySet1((_this).Fm6(p0, _1148_v12, _1148_v12, _1135_v0, globalState), 6)
			(_nw204).ArraySet1(p0, 7)
			(_nw204).ArraySet1(p0, 8)
			(_nw204).ArraySet1(!(p0), 9)
			(_nw204).ArraySet1(p0, 10)
			(_nw204).ArraySet1(Companion_Default___.Fm2(p1, _1135_v0, true, p1, globalState), 11)
			(_nw204).ArraySet1(p0, 12)
			(_nw204).ArraySet1(p1, 13)
			(_nw204).ArraySet1((p1) && (p1), 14)
			(_nw204).ArraySet1(p0, 15)
			(_nw204).ArraySet1(p0, 16)
			(_nw204).ArraySet1(p0, 17)
			_1152_v16 = _nw204
			r1 = _1152_v16
			r0 = !(p0) || (p1)
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
			_ = _index191
			(_1145_v10).ArraySet1(((_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int)).Plus(_1135_v0), (_index191).Int())
			var _1153_v17 D7
			_ = _1153_v17
			_1153_v17 = Companion_D7_.Create_DC17_(!(p1), !(p1), p1)
			var _source13 D7 = _1153_v17
			_ = _source13
			if _source13.Is_DC17() {
				var _1154___mcc_h0 bool = _source13.Get_().(D7_DC17).Cf27
				_ = _1154___mcc_h0
				var _1155___mcc_h1 bool = _source13.Get_().(D7_DC17).Cf28
				_ = _1155___mcc_h1
				var _1156___mcc_h2 bool = _source13.Get_().(D7_DC17).Cf29
				_ = _1156___mcc_h2
				var _1157_cf29 bool = _1156___mcc_h2
				_ = _1157_cf29
				var _1158_cf28 bool = _1155___mcc_h1
				_ = _1158_cf28
				var _1159_cf27 bool = _1154___mcc_h0
				_ = _1159_cf27
				var _1160_v18 _dafny.Map
				_ = _1160_v18
				_1160_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1148_v12, _1152_v16)
				_1160_v18 = (_1160_v18).Merge(_1160_v18)
				var _1161_v19 _dafny.Array
				_ = _1161_v19
				var _nw205 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
				_ = _nw205
				_1161_v19 = _nw205
				var _1162_v20 D3
				_ = _1162_v20
				_1162_v20 = Companion_D3_.Create_DC9_(_1135_v0, _1161_v19, p0)
				_1157_cf29 = (_1162_v20).Dtor_cf12()
				_1145_v10 = _1145_v10
				var _1163_v21 _dafny.Map
				_ = _1163_v21
				_1163_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1159_cf27), _dafny.IntOfUint32((_dafny.SeqOf((_this).F0())).Cardinality()))).Update(_1158_cf28, (_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int)), (_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int))
				var _1164_v22 D7
				_ = _1164_v22
				_1164_v22 = Companion_D7_.Create_DC18_((_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int))
				var _1165_v23 _dafny.Map
				_ = _1165_v23
				_1165_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1164_v22, _1135_v0)
				var _1166_v24 T1
				_ = _1166_v24
				var _nw206 *C2 = New_C2_()
				_ = _nw206
				_nw206.Ctor__(_1135_v0)
				_1166_v24 = _nw206
				var _1167_v25 _dafny.Map
				_ = _1167_v25
				_1167_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1166_v24, p0)
				var _1168_v26 _dafny.Map
				_ = _1168_v26
				_1168_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, ((_1165_v23).Update(_1164_v22, (_1167_v25).Cardinality())).Cardinality())
				var _1169_v27 _dafny.Map
				_ = _1169_v27
				_1169_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1163_v21).Contains(_1168_v26) {
						return (_1163_v21).Get(_1168_v26).(_dafny.Int)
					}
					return (_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int)
				})(), _1148_v12)
				_1169_v27 = (_1169_v27).Update((_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int), _1148_v12)
			} else if _source13.Is_DC18() {
				var _1170___mcc_h3 _dafny.Int = _source13.Get_().(D7_DC18).Cf30
				_ = _1170___mcc_h3
				var _1171_cf30 _dafny.Int = _1170___mcc_h3
				_ = _1171_cf30
				var _1172_v28 _dafny.Map
				_ = _1172_v28
				_1172_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				_1172_v28 = (_1172_v28).Update(!(p0) || (p0), p1)
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
				_ = _index192
				(_1145_v10).ArraySet1((_dafny.IntOfInt64(668)).Minus((_this).F0()), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1152_v16), 0))
				_ = _index193
				(_1152_v16).ArraySet1(p0, (_index193).Int())
				var _1173_v29 D2
				_ = _1173_v29
				_1173_v29 = Companion_D2_.Create_DC4_(_dafny.Companion_Sequence_.Update(_1148_v12, (Companion_Default___.SafeIndex(_1171_cf30, _dafny.IntOfUint32((_1148_v12).Cardinality()))).Uint32(), _1136_v1), p1)
				var _1174_v30 D0
				_ = _1174_v30
				_1174_v30 = Companion_D0_.Create_DC0_((_1173_v29).Dtor_cf5())
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1152_v16), 0))
				_ = _index194
				var _rhs185 bool = p0
				_ = _rhs185
				var _rhs186 D0 = _1174_v30
				_ = _rhs186
				var _lhs82 _dafny.Array = _1152_v16
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1152_v16), 0))
				_ = _lhs83
				(_lhs82).ArraySet1(_rhs185, (_lhs83).Int())
				_1174_v30 = _rhs186
				var _rhs187 _dafny.Sequence = _1148_v12
				_ = _rhs187
				var _rhs188 bool = p1
				_ = _rhs188
				_1148_v12 = _rhs187
				r0 = _rhs188
			} else {
				var _1175___mcc_h4 _dafny.Sequence = _source13.Get_().(D7_DC16).Cf26
				_ = _1175___mcc_h4
				var _1176_cf26 _dafny.Sequence = _1175___mcc_h4
				_ = _1176_cf26
				var _1177_v31 _dafny.Map
				_ = _1177_v31
				_1177_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1135_v0).Cmp((_1149_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1148_v12).Cardinality()), _dafny.IntOfUint32((_1149_v13).Cardinality()))).Uint32()).(_dafny.Int)) > 0, (func() bool {
					if !(p1) {
						return p1
					}
					return p0
				})())
				var _1178_v32 D2
				_ = _1178_v32
				_1178_v32 = Companion_D2_.Create_DC4_(_1148_v12, p0)
				var _1179_v33 _dafny.Map
				_ = _1179_v33
				_1179_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1178_v32).Dtor_cf5(), _1148_v12)
				_1177_v31 = (_1177_v31).Update(!((_1153_v17).Dtor_cf29()), (_this).Fm6(Companion_Default___.Fm2(p1, (_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int), p1, false, globalState), (func() _dafny.Sequence {
					if (_1179_v33).Contains(p0) {
						return (_1179_v33).Get(p0).(_dafny.Sequence)
					}
					return _1148_v12
				})(), _1148_v12, _1135_v0, globalState))
				var _1180_v34 _dafny.Array
				_ = _1180_v34
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
				_ = _nw207
				_1180_v34 = _nw207
				var _1181_v35 _dafny.MultiSet
				_ = _1181_v35
				_1181_v35 = _dafny.MultiSetOf(_1136_v1)
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1180_v34), 0))
				_ = _index195
				(_1180_v34).ArraySet1(_1181_v35, (_index195).Int())
				var _1182_v36 T0
				_ = _1182_v36
				var _nw208 *C3 = New_C3_()
				_ = _nw208
				_nw208.Ctor__((_this).F0())
				_1182_v36 = _nw208
				var _1183_v37 D10
				_ = _1183_v37
				_1183_v37 = Companion_D10_.Create_DC21_(_1182_v36)
				var _1184_v38 _dafny.Map
				_ = _1184_v38
				_1184_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1181_v35).Cardinality(), _1183_v37)
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
				_ = _index196
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1180_v34), 0))
				_ = _index197
				var _rhs189 bool = !(p0) || (!(((_1184_v38).Cardinality()).Cmp(_dafny.IntOfUint32((_1148_v12).Cardinality())) < 0))
				_ = _rhs189
				var _rhs190 bool = p0
				_ = _rhs190
				var _rhs191 _dafny.Int = (_1182_v36).F0()
				_ = _rhs191
				var _rhs192 _dafny.MultiSet = _1181_v35
				_ = _rhs192
				var _lhs84 _dafny.Array = _1145_v10
				_ = _lhs84
				var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))
				_ = _lhs85
				var _lhs86 _dafny.Array = _1180_v34
				_ = _lhs86
				var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1180_v34), 0))
				_ = _lhs87
				r0 = _rhs189
				r0 = _rhs190
				(_lhs84).ArraySet1(_rhs191, (_lhs85).Int())
				(_lhs86).ArraySet1(_rhs192, (_lhs87).Int())
				var _1185_v39 *C1
				_ = _1185_v39
				var _nw209 *C1 = New_C1_()
				_ = _nw209
				_nw209.Ctor__((_1145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1145_v10), 0))).Int()).(_dafny.Int))
				_1185_v39 = _nw209
				var _1186_v40 _dafny.Map
				_ = _1186_v40
				_1186_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1135_v0, _1185_v39)
				var _1187_v41 _dafny.Map
				_ = _1187_v41
				_1187_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_1149_v13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(88))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1188_v10 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_1189_i1 _dafny.Int) _dafny.Int {
						return (_1188_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1188_v10), 0))).Int()).(_dafny.Int)
					}
				})(_1145_v10)))), (_1186_v40).Cardinality())
				_1187_v41 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1182_v36).F0())).Merge(_1187_v41)
				var _1190_v42 _dafny.Map
				_ = _1190_v42
				_1190_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v1, _1149_v13)
				_1149_v13 = (func() _dafny.Sequence {
					if (_1190_v42).Contains(_1136_v1) {
						return (_1190_v42).Get(_1136_v1).(_dafny.Sequence)
					}
					return _1149_v13
				})()
			}
		}
		var _1191_v43 _dafny.Set
		_ = _1191_v43
		_1191_v43 = _dafny.SetOf(p1, p1, true)
		var _hi9 _dafny.Int = Companion_Default___.Fm3(_1191_v43, _1136_v1, _1135_v0, (_this).F0(), globalState)
		_ = _hi9
		for _1192_i2 := (_this).F0(); _1192_i2.Cmp(_hi9) < 0; _1192_i2 = _1192_i2.Plus(_dafny.One) {
			var _1193_v44 _dafny.Map
			_ = _1193_v44
			_1193_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1135_v0)
			_1193_v44 = (_1193_v44).Update(p1, _1192_i2)
			var _1194_v45 _dafny.Array
			_ = _1194_v45
			var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw210
			_1194_v45 = _nw210
			var _1195_v46 _dafny.Sequence
			_ = _1195_v46
			_1195_v46 = _dafny.UnicodeSeqOfUtf8Bytes("k")
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1194_v45), 0))
			_ = _index198
			(_1194_v45).ArraySet1((_1135_v0).Plus(_dafny.IntOfUint32((_1195_v46).Cardinality())), (_index198).Int())
			var _1196_v47 _dafny.Array
			_ = _1196_v47
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_27
			var _nw211 _dafny.Array
			_ = _nw211
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw211 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = (func(_1197_p0 bool) func(_dafny.Int) bool {
					return func(_1198_i3 _dafny.Int) bool {
						return _1197_p0
					}
				})(p0)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw211 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw211).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw211).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1196_v47 = _nw211
			var _1199_v48 D11
			_ = _1199_v48
			_1199_v48 = Companion_D11_.Create_DC24_(_1196_v47, p1, !(p0))
			var _1200_v49 D5
			_ = _1200_v49
			_1200_v49 = Companion_D5_.Create_DC13_(_1192_i2, _dafny.SetOf(true, true, p0, p1, p1), (_1195_v46).Select((Companion_Default___.SafeIndex(_1192_i2, _dafny.IntOfUint32((_1195_v46).Cardinality()))).Uint32()).(_dafny.CodePoint), _1192_i2, p0)
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1194_v45), 0))
			_ = _index199
			var _rhs193 _dafny.Int = ((_this).F0()).Times((_1191_v43).Cardinality())
			_ = _rhs193
			var _rhs194 D11 = Companion_D11_.Create_DC24_(_1196_v47, (_1200_v49).Dtor_cf21(), p1)
			_ = _rhs194
			var _rhs195 bool = (p1) || (p0)
			_ = _rhs195
			var _lhs88 _dafny.Array = _1194_v45
			_ = _lhs88
			var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1194_v45), 0))
			_ = _lhs89
			(_lhs88).ArraySet1(_rhs193, (_lhs89).Int())
			_1199_v48 = _rhs194
			r0 = _rhs195
			var _1201_v50 _dafny.Map
			_ = _1201_v50
			_1201_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1194_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1194_v45), 0))).Int()).(_dafny.Int), !(p1))
			var _rhs196 _dafny.Int = _1192_i2
			_ = _rhs196
			var _rhs197 bool = !(((_this).F0()).Cmp((_1194_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1194_v45), 0))).Int()).(_dafny.Int)) < 0)
			_ = _rhs197
			var _rhs198 _dafny.Map = (_1201_v50).Merge(_1201_v50)
			_ = _rhs198
			var _rhs199 bool = (p1) == ((_1191_v43).IsProperSubsetOf(_1191_v43))
			_ = _rhs199
			_1135_v0 = _rhs196
			r0 = _rhs197
			_1201_v50 = _rhs198
			r0 = _rhs199
			var _1202_v51 _dafny.Map
			_ = _1202_v51
			_1202_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1194_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1194_v45), 0))).Int()).(_dafny.Int), _1192_i2)
			var _1203_v52 _dafny.Map
			_ = _1203_v52
			_1203_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v51, _1194_v45)
			_1203_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v51, _1194_v45)
		}
		var _1204_v53 D7
		_ = _1204_v53
		_1204_v53 = Companion_D7_.Create_DC17_(p1, p1, (p1) == (p1))
		var _source14 D7 = _1204_v53
		_ = _source14
		if _source14.Is_DC17() {
			var _1205___mcc_h5 bool = _source14.Get_().(D7_DC17).Cf27
			_ = _1205___mcc_h5
			var _1206___mcc_h6 bool = _source14.Get_().(D7_DC17).Cf28
			_ = _1206___mcc_h6
			var _1207___mcc_h7 bool = _source14.Get_().(D7_DC17).Cf29
			_ = _1207___mcc_h7
			var _1208_cf29 bool = _1207___mcc_h7
			_ = _1208_cf29
			var _1209_cf28 bool = _1206___mcc_h6
			_ = _1209_cf28
			var _1210_cf27 bool = _1205___mcc_h5
			_ = _1210_cf27
			var _1211_v54 _dafny.Map
			_ = _1211_v54
			_1211_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1135_v0, true)
			_1211_v54 = _1211_v54
			_1135_v0 = _1135_v0
			var _1212_v55 _dafny.Array
			_ = _1212_v55
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_28
			var _nw212 _dafny.Array
			_ = _nw212
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw212 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = (func(_1213_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1214_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1214_i4, _1213_v0)
					}
				})(_1135_v0)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw212 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw212).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw212).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1212_v55 = _nw212
			var _1215_v56 _dafny.Sequence
			_ = _1215_v56
			_1215_v56 = _dafny.SeqOf(!(p0))
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_1212_v55), 0))
			_ = _index200
			(_1212_v55).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1215_v56, _1215_v56)).Cardinality()), (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_1212_v55), 0))
			_ = _index201
			(_1212_v55).ArraySet1(_1135_v0, (_index201).Int())
			var _1216_v57 *C1
			_ = _1216_v57
			var _nw213 *C1 = New_C1_()
			_ = _nw213
			_nw213.Ctor__((_this).F0())
			_1216_v57 = _nw213
		} else if _source14.Is_DC18() {
			var _1217___mcc_h8 _dafny.Int = _source14.Get_().(D7_DC18).Cf30
			_ = _1217___mcc_h8
			var _1218_cf30 _dafny.Int = _1217___mcc_h8
			_ = _1218_cf30
			var _1219_v58 _dafny.Array
			_ = _1219_v58
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_29
			var _nw214 _dafny.Array
			_ = _nw214
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw214 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Sequence = (func(_1220_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1221_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(Companion_D5_.Create_DC11_(_dafny.SetOf(_dafny.SetOf((_dafny.SetOf((_this).F0(), _dafny.IntOfInt64(810))).Cardinality()))), Companion_D5_.Create_DC11_(_dafny.SetOf(_dafny.SetOf((_this).F0(), _1220_v0, (_this).F0(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqlvvtuaq")).Cardinality()), (_dafny.MultiSetOf((_this).F0())).Cardinality()))))
					}
				})(_1135_v0)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw214 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw214).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw214).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1219_v58 = _nw214
			var _1222_v59 _dafny.Array
			_ = _1222_v59
			var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw215
			_1222_v59 = _nw215
			var _1223_v60 _dafny.MultiSet
			_ = _1223_v60
			_1223_v60 = _dafny.MultiSetOf(_dafny.IntOfInt64(-22), _1135_v0)
			var _1224_v61 _dafny.Sequence
			_ = _1224_v61
			_1224_v61 = _dafny.UnicodeSeqOfUtf8Bytes("kkw")
			var _rhs200 _dafny.Array = _1222_v59
			_ = _rhs200
			var _rhs201 _dafny.Int = (func() _dafny.Int {
				if (_1223_v60).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality()) {
					return (_1223_v60).Multiplicity((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality())
				}
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nqka")).Cardinality()), (_dafny.SetOf(p1, p1, p0)).Cardinality())
			})()
			_ = _rhs201
			var _rhs202 _dafny.Int = (Companion_Default___.SafeDivisionInt(_1135_v0, (_this).F0())).Minus(Companion_Default___.SafeDivisionInt(_1135_v0, _dafny.IntOfUint32((_1224_v61).Cardinality())))
			_ = _rhs202
			var _rhs203 _dafny.Array = _1219_v58
			_ = _rhs203
			r1 = _rhs200
			_1135_v0 = _rhs201
			_1135_v0 = _rhs202
			_1219_v58 = _rhs203
			var _1225_v62 _dafny.Array
			_ = _1225_v62
			var _nwElement0_34 _dafny.Array = _1222_v59
			_ = _nwElement0_34
			var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(18))
			_ = _nw216
			(_nw216).ArraySet1(_nwElement0_34, 0)
			(_nw216).ArraySet1((func() _dafny.Array {
				if !((Companion_D0_.Create_DC1_(p0)).Dtor_cf1()) {
					return _1222_v59
				}
				return _1222_v59
			})(), 1)
			(_nw216).ArraySet1(_1222_v59, 2)
			(_nw216).ArraySet1(_1222_v59, 3)
			(_nw216).ArraySet1(_1222_v59, 4)
			(_nw216).ArraySet1(_1222_v59, 5)
			(_nw216).ArraySet1(_1222_v59, 6)
			(_nw216).ArraySet1(_1222_v59, 7)
			(_nw216).ArraySet1(_1222_v59, 8)
			(_nw216).ArraySet1(_1222_v59, 9)
			(_nw216).ArraySet1(_1222_v59, 10)
			(_nw216).ArraySet1(_1222_v59, 11)
			(_nw216).ArraySet1(_1222_v59, 12)
			(_nw216).ArraySet1(_1222_v59, 13)
			(_nw216).ArraySet1(_1222_v59, 14)
			(_nw216).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1222_v59
				}
				return _1222_v59
			})(), 15)
			(_nw216).ArraySet1(_1222_v59, 16)
			(_nw216).ArraySet1(_1222_v59, 17)
			_1225_v62 = _nw216
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1225_v62), 0))
			_ = _index202
			(_1225_v62).ArraySet1(_1222_v59, (_index202).Int())
			var _1226_v63 _dafny.Array
			_ = _1226_v63
			var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw217
			_1226_v63 = _nw217
			var _1227_v64 _dafny.MultiSet
			_ = _1227_v64
			_1227_v64 = _dafny.MultiSetOf(p1, (_this).Fm6((_this).Fm5(globalState), _1224_v61, _1224_v61, _1218_cf30, globalState), p0)
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1225_v62), 0))
			_ = _index203
			var _rhs204 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1224_v61, _1224_v61)
			_ = _rhs204
			var _rhs205 _dafny.Array = _1226_v63
			_ = _rhs205
			var _rhs206 bool = p1
			_ = _rhs206
			var _rhs207 _dafny.Array = _1222_v59
			_ = _rhs207
			var _rhs208 bool = ((_dafny.MultiSetOf(!(p0))).Union(_1227_v64)).IsProperSubsetOf(((_1227_v64).Update(p1, Companion_Default___.Abs(_1218_cf30))).Difference(_1227_v64))
			_ = _rhs208
			var _lhs90 _dafny.Array = _1225_v62
			_ = _lhs90
			var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1225_v62), 0))
			_ = _lhs91
			_1224_v61 = _rhs204
			r2 = _rhs205
			r0 = _rhs206
			(_lhs90).ArraySet1(_rhs207, (_lhs91).Int())
			r0 = _rhs208
			var _1228_v65 *C0
			_ = _1228_v65
			var _nw218 *C0 = New_C0_()
			_ = _nw218
			_nw218.Ctor__(_1136_v1)
			_1228_v65 = _nw218
			_1228_v65 = _1228_v65
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_1226_v63), 0))
			_ = _index204
			(_1226_v63).ArraySet1(_1218_cf30, (_index204).Int())
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_1226_v63), 0))
			_ = _index205
			(_1226_v63).ArraySet1((_1135_v0).Plus((_this).F0()), (_index205).Int())
		} else {
			var _1229___mcc_h9 _dafny.Sequence = _source14.Get_().(D7_DC16).Cf26
			_ = _1229___mcc_h9
			var _1230_cf26 _dafny.Sequence = _1229___mcc_h9
			_ = _1230_cf26
			var _1231_v66 _dafny.Sequence
			_ = _1231_v66
			_1231_v66 = _dafny.UnicodeSeqOfUtf8Bytes("jqjpl")
			var _1232_v67 _dafny.Set
			_ = _1232_v67
			_1232_v67 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("vk"), _1231_v66)
			var _1233_v68 _dafny.Map
			_ = _1233_v68
			_1233_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), p1)
			var _1234_v69 _dafny.Map
			_ = _1234_v69
			_1234_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jut"), _1233_v68)
			r0 = (func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter27 := _dafny.Iterate((_1234_v69).Keys().Elements()); ; {
					_compr_25, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _1235_v70 _dafny.Sequence
					_1235_v70 = interface{}(_compr_25).(_dafny.Sequence)
					if (_1234_v69).Contains(_1235_v70) {
						_coll25.Add(_1235_v70)
					}
				}
				return _coll25.ToSet()
			}()).IsProperSubsetOf(_1232_v67)
			var _1236_v71 _dafny.Map
			_ = _1236_v71
			_1236_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _1237_v72 _dafny.Set
			_ = _1237_v72
			_1237_v72 = _dafny.SetOf(_1135_v0, (_1236_v71).Cardinality())
			r0 = !(_dafny.SetOf((_this).F0(), _dafny.IntOfUint32((Companion_Default___.Fm0(globalState)).Cardinality()), _1135_v0, _1135_v0)).Equals(_1237_v72)
			_1136_v1 = _1136_v1
			var _1238_v73 *C0
			_ = _1238_v73
			var _nw219 *C0 = New_C0_()
			_ = _nw219
			_nw219.Ctor__(_1136_v1)
			_1238_v73 = _nw219
		}
		r0 = ((_1191_v43).Union(_1191_v43)).IsSubsetOf((func() _dafny.Set {
			if p1 {
				return _1191_v43
			}
			return _1191_v43
		})())
		var _1239_v74 _dafny.Array
		_ = _1239_v74
		var _nw220 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw220
		_1239_v74 = _nw220
		r1 = _1239_v74
		var _1240_v75 _dafny.Sequence
		_ = _1240_v75
		_1240_v75 = _dafny.UnicodeSeqOfUtf8Bytes("shkq")
		var _1241_v76 _dafny.MultiSet
		_ = _1241_v76
		_1241_v76 = _dafny.MultiSetOf(_1135_v0)
		var _1242_v77 _dafny.Sequence
		_ = _1242_v77
		_1242_v77 = _dafny.SeqOf((_this).F0(), _1135_v0)
		var _1243_v78 _dafny.Array
		_ = _1243_v78
		var _nw221 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
		_ = _nw221
		_1243_v78 = _nw221
		var _pat_let_tv31 = _1135_v0
		_ = _pat_let_tv31
		var _1244_v79 _dafny.Array
		_ = _1244_v79
		var _nwElement0_35 _dafny.Int = (_dafny.IntOfUint32((_1240_v75).Cardinality())).Times(_1135_v0)
		_ = _nwElement0_35
		var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(7))
		_ = _nw222
		(_nw222).ArraySet1(_nwElement0_35, 0)
		(_nw222).ArraySet1(_1135_v0, 1)
		(_nw222).ArraySet1(_1135_v0, 2)
		(_nw222).ArraySet1(((_1241_v76).Difference(_dafny.MultiSetFromSeq(_1242_v77))).Cardinality(), 3)
		(_nw222).ArraySet1(_1135_v0, 4)
		(_nw222).ArraySet1((_dafny.Zero).Minus((func(_pat_let23_0 D3) D3 {
			return func(_1245_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let24_0 _dafny.Int) D3 {
					return func(_1246_dt__update_hcf10_h0 _dafny.Int) D3 {
						return Companion_D3_.Create_DC9_(_1246_dt__update_hcf10_h0, (_1245_dt__update__tmp_h0).Dtor_cf11(), (_1245_dt__update__tmp_h0).Dtor_cf12())
					}(_pat_let24_0)
				}(_pat_let_tv31)
			}(_pat_let23_0)
		}(Companion_D3_.Create_DC9_((_this).F0(), _1243_v78, p1))).Dtor_cf10()), 5)
		(_nw222).ArraySet1(_1135_v0, 6)
		_1244_v79 = _nw222
		r2 = _1244_v79
		var _1247_v80 _dafny.MultiSet
		_ = _1247_v80
		_1247_v80 = _dafny.MultiSetOf(false)
		r3 = ((_1247_v80).Update(p0, Companion_Default___.Abs(_1135_v0))).Intersection(_1247_v80)
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f0 _dafny.Int
	_f1 _dafny.CodePoint
	_f2 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.CodePoint('D')
	_this._f2 = _dafny.Zero
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

func (_this *C5) F0() _dafny.Int {
	return _this._f0
}
func (_this *C5) Ctor__(f1 _dafny.CodePoint, f2 _dafny.Int, f0 _dafny.Int) {
	{
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f0 = f0
	}
}
func (_this *C5) Fm5(globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C5) Fm6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(!((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(573))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg55 _dafny.Int) interface{} {
				return coer55(arg55)
			}
		}(func(_1248_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ilqiu")).Cardinality())
		})))).Contains((_this).F0()))) || ((true) == (true))
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('n')
	}
}
func (_this *C5) M2(globalState *GlobalState) {
	{
		var _1249_v0 *C0
		_ = _1249_v0
		var _nw223 *C0 = New_C0_()
		_ = _nw223
		_nw223.Ctor__((_this).F1())
		_1249_v0 = _nw223
		var _1250_v1 bool
		_ = _1250_v1
		_1250_v1 = true
		var _1251_v2 _dafny.Sequence
		_ = _1251_v2
		_1251_v2 = _dafny.SeqOf(true, _1250_v1)
		var _1252_v3 _dafny.Map
		_ = _1252_v3
		_1252_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v1, _1251_v2)
		var _1253_v4 _dafny.Array
		_ = _1253_v4
		var _nwElement0_36 bool = _1250_v1
		_ = _nwElement0_36
		var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(9))
		_ = _nw224
		(_nw224).ArraySet1(_nwElement0_36, 0)
		(_nw224).ArraySet1((_1250_v1) == (_1250_v1), 1)
		(_nw224).ArraySet1(_1250_v1, 2)
		(_nw224).ArraySet1(_1250_v1, 3)
		(_nw224).ArraySet1(_1250_v1, 4)
		(_nw224).ArraySet1(_1250_v1, 5)
		(_nw224).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1251_v2, (func() _dafny.Sequence {
			if (_1252_v3).Contains(_1250_v1) {
				return (_1252_v3).Get(_1250_v1).(_dafny.Sequence)
			}
			return _1251_v2
		})()), 6)
		(_nw224).ArraySet1(!(_1250_v1) || (_1250_v1), 7)
		(_nw224).ArraySet1(_1250_v1, 8)
		_1253_v4 = _nw224
		var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
		_ = _index206
		(_1253_v4).ArraySet1(_1250_v1, (_index206).Int())
		var _1254_v5 _dafny.Map
		_ = _1254_v5
		_1254_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).Fm5(globalState)) || (_1250_v1), _dafny.MultiSetOf(_1250_v1))
		var _1255_v6 _dafny.Sequence
		_ = _1255_v6
		_1255_v6 = _dafny.SeqOf(_1251_v2, _dafny.SeqOf(_1250_v1))
		var _1256_v7 _dafny.Map
		_ = _1256_v7
		_1256_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_1249_v0).F3())
		var _1257_v8 _dafny.Sequence
		_ = _1257_v8
		_1257_v8 = _dafny.UnicodeSeqOfUtf8Bytes("sargafyl")
		var _1258_v9 _dafny.Sequence
		_ = _1258_v9
		_1258_v9 = _dafny.SeqOf((func() _dafny.CodePoint {
			if (_1256_v7).Contains(_1250_v1) {
				return (_1256_v7).Get(_1250_v1).(_dafny.CodePoint)
			}
			return (_this).F1()
		})(), (_1249_v0).F3(), (_1257_v8).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_1257_v8).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F1())
		var _1259_v10 _dafny.Sequence
		_ = _1259_v10
		_1259_v10 = _dafny.SeqOf((_1258_v9).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1258_v9).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _1260_v11 _dafny.MultiSet
		_ = _1260_v11
		_1260_v11 = _dafny.MultiSetOf(_1250_v1)
		var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
		_ = _index207
		var _rhs209 bool = ((_this).F0()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1255_v6, (Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_1255_v6).Cardinality()))).Uint32(), _1251_v2)).Cardinality())) < 0
		_ = _rhs209
		var _rhs210 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1258_v9, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F0()), _dafny.IntOfUint32((_1258_v9).Cardinality()))).Uint32(), (_this).F1()), (_this).F1())
		_ = _rhs210
		var _rhs211 _dafny.Map = ((_1254_v5).Update(_1250_v1, _dafny.MultiSetOf(_1250_v1, !(_1250_v1), _1250_v1))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v1, _1260_v11))
		_ = _rhs211
		var _lhs92 _dafny.Array = _1253_v4
		_ = _lhs92
		var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
		_ = _lhs93
		(_lhs92).ArraySet1(_rhs209, (_lhs93).Int())
		_1250_v1 = _rhs210
		_1254_v5 = _rhs211
		var _1261_v12 _dafny.Int
		_ = _1261_v12
		_1261_v12 = _dafny.IntOfInt64(796)
		_1261_v12 = _dafny.IntOfInt64(234)
		var _1262_v13 _dafny.Sequence
		_ = _1262_v13
		_1262_v13 = _dafny.SeqOf(_1249_v0, _1249_v0, _1249_v0)
		var _1263_v14 _dafny.Map
		_ = _1263_v14
		_1263_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1260_v11, _1250_v1)
		var _1264_v15 _dafny.Sequence
		_ = _1264_v15
		_1264_v15 = _dafny.SeqOf(_dafny.IntOfInt64(182), (_dafny.Zero).Minus((_this).F0()))
		if !((_dafny.MultiSetFromSeq(_1264_v15)).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfUint32((_1262_v13).Cardinality()), _dafny.IntOfUint32((_1251_v2).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdcu")).Cardinality()), (_1263_v14).Cardinality()), (Companion_Default___.SafeIndex(_1261_v12, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1262_v13).Cardinality()), _dafny.IntOfUint32((_1251_v2).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdcu")).Cardinality()), (_1263_v14).Cardinality())).Cardinality()))).Uint32(), (_this).F2())).Cardinality())))) {
			var _1265_v16 _dafny.Set
			_ = _1265_v16
			_1265_v16 = _dafny.SetOf(_1261_v12, (_this).F2(), (_this).F0())
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _index208
			var _rhs212 bool = (_dafny.IntOfInt64(745)).Cmp((_dafny.IntOfUint32((_1264_v15).Cardinality())).Minus(_1261_v12)) != 0
			_ = _rhs212
			var _rhs213 _dafny.Int = _1261_v12
			_ = _rhs213
			var _rhs214 bool = (_1250_v1) && (false)
			_ = _rhs214
			var _rhs215 _dafny.Set = _1265_v16
			_ = _rhs215
			var _lhs94 _dafny.Array = _1253_v4
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _lhs95
			(_lhs94).ArraySet1(_rhs212, (_lhs95).Int())
			_1261_v12 = _rhs213
			_1250_v1 = _rhs214
			_1265_v16 = _rhs215
			var _1266_v17 _dafny.Array
			_ = _1266_v17
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_30
			var _nw225 _dafny.Array
			_ = _nw225
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw225 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Sequence = (func(_1267_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1268_i0 _dafny.Int) _dafny.Sequence {
						return _1267_v2
					}
				})(_1251_v2)
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw225 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw225).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw225).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1266_v17 = _nw225
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1266_v17), 0))
			_ = _index209
			(_1266_v17).ArraySet1(_1251_v2, (_index209).Int())
			var _1269_v18 _dafny.Sequence
			_ = _1269_v18
			_1269_v18 = _dafny.SeqOf(_1253_v4)
			var _1270_v19 D3
			_ = _1270_v19
			_1270_v19 = Companion_D3_.Create_DC7_(_1264_v15)
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1266_v17), 0))
			_ = _index210
			var _rhs216 _dafny.Array = (_1269_v18).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_1261_v12, _dafny.IntOfUint32(((_1270_v19).Dtor_cf7()).Cardinality())), _dafny.IntOfUint32((_1269_v18).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs216
			var _rhs217 bool = !(_1250_v1)
			_ = _rhs217
			var _rhs218 _dafny.Sequence = (func() _dafny.Sequence {
				if _1250_v1 {
					return _1251_v2
				}
				return _1251_v2
			})()
			_ = _rhs218
			var _rhs219 bool = (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool)
			_ = _rhs219
			var _lhs96 _dafny.Array = _1266_v17
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1266_v17), 0))
			_ = _lhs97
			_1253_v4 = _rhs216
			_1250_v1 = _rhs217
			(_lhs96).ArraySet1(_rhs218, (_lhs97).Int())
			_1250_v1 = _rhs219
			var _1271_v20 _dafny.Map
			_ = _1271_v20
			_1271_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), _1253_v4)
			if !(_1271_v20).Contains((_dafny.Zero).Minus(_1261_v12)) {
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
				_ = _index211
				(_1253_v4).ArraySet1((func() bool {
					if (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool) {
						return !(!((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool)) || ((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool)))
					}
					return (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool)
				})(), (_index211).Int())
				var _1272_v21 _dafny.MultiSet
				_ = _1272_v21
				_1272_v21 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_1261_v12, _dafny.IntOfInt64(573)), ((_this).F2()).Plus(_1261_v12), (_dafny.IntOfUint32((_1257_v8).Cardinality())).Plus((_this).F0()))
				_1272_v21 = _1272_v21
				var _1273_v22 *C0
				_ = _1273_v22
				var _nw226 *C0 = New_C0_()
				_ = _nw226
				_nw226.Ctor__((_1249_v0).F3())
				_1273_v22 = _nw226
				var _1274_v23 _dafny.Array
				_ = _1274_v23
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_31
				var _nw227 _dafny.Array
				_ = _nw227
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw227 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1275_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1276_i1 _dafny.Int) _dafny.Sequence {
							return _1275_v9
						}
					})(_1258_v9)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw227 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw227).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw227).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1274_v23 = _nw227
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_1274_v23), 0))
				_ = _index212
				(_1274_v23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("eblbebi"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eblbebi")).Cardinality()))).Uint32(), (_1249_v0).F3()), _1258_v9), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_1274_v23), 0))
				_ = _index213
				(_1274_v23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1258_v9, _1258_v9), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
				_ = _index214
				(_1253_v4).ArraySet1((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool), (_index214).Int())
			} else {
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
				_ = _index215
				(_1253_v4).ArraySet1(_1250_v1, (_index215).Int())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
				_ = _index216
				(_1253_v4).ArraySet1((_this).Fm5(globalState), (_index216).Int())
				_1261_v12 = _1261_v12
				var _1277_v24 *C0
				_ = _1277_v24
				var _nw228 *C0 = New_C0_()
				_ = _nw228
				_nw228.Ctor__((_1249_v0).F3())
				_1277_v24 = _nw228
				var _1278_v25 _dafny.MultiSet
				_ = _1278_v25
				_1278_v25 = _dafny.MultiSetOf((_this).F2())
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
				_ = _index217
				(_1253_v4).ArraySet1(!(_1278_v25).Equals((_1278_v25).Update((_this).F0(), Companion_Default___.Abs(_dafny.IntOfInt64(46)))), (_index217).Int())
			}
			var _1279_v26 _dafny.Array
			_ = _1279_v26
			var _nw229 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
			_ = _nw229
			_1279_v26 = _nw229
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_1279_v26), 0))
			_ = _index218
			(_1279_v26).ArraySet1(_1249_v0, (_index218).Int())
			var _1280_v27 _dafny.CodePoint
			_ = _1280_v27
			_1280_v27 = (_this).F1()
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_1279_v26), 0))
			_ = _index219
			var _nw230 *C0 = New_C0_()
			_ = _nw230
			_nw230.Ctor__((_1280_v27))
			(_1279_v26).ArraySet1(_nw230, (_index219).Int())
			_1264_v15 = _1264_v15
		} else {
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _index220
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _index221
			var _rhs220 bool = (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool)
			_ = _rhs220
			var _rhs221 bool = !(_1250_v1) || ((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool))
			_ = _rhs221
			var _lhs98 _dafny.Array = _1253_v4
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _lhs99
			var _lhs100 _dafny.Array = _1253_v4
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _lhs101
			(_lhs98).ArraySet1(_rhs220, (_lhs99).Int())
			(_lhs100).ArraySet1(_rhs221, (_lhs101).Int())
			var _1281_v28 D2
			_ = _1281_v28
			_1281_v28 = Companion_D2_.Create_DC5_()
			var _1282_v29 D3
			_ = _1282_v29
			_1282_v29 = Companion_D3_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(925))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_1283_v15 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1284_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_1283_v15).Cardinality())
				}
			})(_1264_v15))))
			var _1285_v30 _dafny.Map
			_ = _1285_v30
			_1285_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v1, _1282_v29)
			var _1286_v31 _dafny.Map
			_ = _1286_v31
			_1286_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1261_v12, _dafny.IntOfInt64(-499), (_this).F0()), (_1249_v0).F3())
			_1281_v28 = Companion_Default___.Fm8((_dafny.Zero).Minus(_1261_v12), _1285_v30, _1286_v31, (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool), globalState)
			var _1287_v32 D0
			_ = _1287_v32
			_1287_v32 = Companion_D0_.Create_DC0_(_1250_v1)
			_1250_v1 = (_1287_v32).Dtor_cf0()
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
			_ = _index222
			(_1253_v4).ArraySet1(!(!((_1261_v12).Cmp((_1261_v12).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1261_v12)).Cardinality()))) >= 0)), (_index222).Int())
			var _source15 D3 = _1282_v29
			_ = _source15
			if _source15.Is_DC8() {
				var _1288___mcc_h0 bool = _source15.Get_().(D3_DC8).Cf8
				_ = _1288___mcc_h0
				var _1289___mcc_h1 _dafny.Int = _source15.Get_().(D3_DC8).Cf9
				_ = _1289___mcc_h1
				var _1290_cf9 _dafny.Int = _1289___mcc_h1
				_ = _1290_cf9
				var _1291_cf8 bool = _1288___mcc_h0
				_ = _1291_cf8
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))
				_ = _index223
				(_1253_v4).ArraySet1(_1291_cf8, (_index223).Int())
				var _1292_v33 D2
				_ = _1292_v33
				_1292_v33 = Companion_D2_.Create_DC5_()
				var _1293_v34 D2
				_ = _1293_v34
				_1293_v34 = Companion_D2_.Create_DC6_(_1292_v33)
				_1293_v34 = Companion_Default___.Fm9(globalState)
				var _1294_v35 D3
				_ = _1294_v35
				_1294_v35 = Companion_D3_.Create_DC8_(_1250_v1, (_this).F0())
				_1261_v12 = Companion_Default___.SafeDivisionInt((_1294_v35).Dtor_cf9(), _1290_cf9)
				var _1295_v36 _dafny.Map
				_ = _1295_v36
				_1295_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1291_cf8), _1291_cf8)
				_1295_v36 = (_1295_v36).Update(_1291_cf8, _1291_cf8)
			} else if _source15.Is_DC9() {
				var _1296___mcc_h2 _dafny.Int = _source15.Get_().(D3_DC9).Cf10
				_ = _1296___mcc_h2
				var _1297___mcc_h3 _dafny.Array = _source15.Get_().(D3_DC9).Cf11
				_ = _1297___mcc_h3
				var _1298___mcc_h4 bool = _source15.Get_().(D3_DC9).Cf12
				_ = _1298___mcc_h4
				var _1299_cf12 bool = _1298___mcc_h4
				_ = _1299_cf12
				var _1300_cf11 _dafny.Array = _1297___mcc_h3
				_ = _1300_cf11
				var _1301_cf10 _dafny.Int = _1296___mcc_h2
				_ = _1301_cf10
				var _1302_v37 *C0
				_ = _1302_v37
				var _nw231 *C0 = New_C0_()
				_ = _nw231
				_nw231.Ctor__((_1249_v0).F3())
				_1302_v37 = _nw231
				var _1303_v38 _dafny.Sequence
				_ = _1303_v38
				_1303_v38 = _dafny.SeqOf(_1258_v9, _dafny.UnicodeSeqOfUtf8Bytes("lrbh"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xh"), _1257_v8))
				var _1304_v39 _dafny.Map
				_ = _1304_v39
				_1304_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-525))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1305_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1306_i4 _dafny.Int) _dafny.Sequence {
						return _1305_v8
					}
				})(_1257_v8))))
				var _1307_v40 _dafny.Sequence
				_ = _1307_v40
				_1307_v40 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1303_v38, (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1303_v38).Cardinality()))).Uint32(), _1258_v9), _1303_v38), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-71))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_1308_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1309_i3 _dafny.Int) _dafny.Sequence {
						return _1308_v10
					}
				})(_1259_v10))), (func() _dafny.Sequence {
					if (_1304_v39).Contains(!((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool))) {
						return (_1304_v39).Get(!((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool))).(_dafny.Sequence)
					}
					return _1303_v38
				})())
				var _rhs222 *C0 = (func() *C0 {
					if (_1251_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1301_cf10), _dafny.IntOfUint32((_1251_v2).Cardinality()))).Uint32()).(bool) {
						return _1302_v37
					}
					return _1302_v37
				})()
				_ = _rhs222
				var _rhs223 _dafny.Sequence = (_1307_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-693), _dafny.IntOfUint32((_1307_v40).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs223
				_1302_v37 = _rhs222
				_1303_v38 = _rhs223
				var _1310_v41 _dafny.MultiSet
				_ = _1310_v41
				_1310_v41 = _dafny.MultiSetOf(_dafny.CodePoint('p'), (_this).F1(), (_1302_v37).F3(), (_1249_v0).F3(), (_1249_v0).F3())
				var _1311_v42 _dafny.Set
				_ = _1311_v42
				_1311_v42 = _dafny.SetOf(_1299_cf12, _1299_cf12, (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool), (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool))
				var _1312_v43 _dafny.MultiSet
				_ = _1312_v43
				_1312_v43 = _dafny.MultiSetOf((_this).F2(), (_this).F2(), (_this).F0(), _1261_v12)
				_1299_cf12 = ((Companion_Default___.Fm3(Companion_Default___.Fm10(_1310_v41, _1299_cf12, globalState), (_1249_v0).F3(), (_this).F2(), Companion_Default___.Fm3(_1311_v42, (_1249_v0).F3(), _dafny.IntOfUint32((_1258_v9).Cardinality()), (_this).F2(), globalState), globalState)).Minus(_1301_cf10)).Cmp(((_1312_v43).Difference(_dafny.MultiSetOf((_this).F2(), _1261_v12, (_this).F2()))).Cardinality()) <= 0
				var _1313_v44 _dafny.Map
				_ = _1313_v44
				_1313_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1299_cf12, !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("uuhjyw"), (_1302_v37).F3()))
				_1313_v44 = (_1313_v44).Update(!(_1250_v1), _1250_v1)
				_1261_v12 = ((_dafny.Zero).Minus(_1261_v12)).Plus(_1261_v12)
			} else {
				var _1314___mcc_h5 _dafny.Sequence = _source15.Get_().(D3_DC7).Cf7
				_ = _1314___mcc_h5
				var _1315_cf7 _dafny.Sequence = _1314___mcc_h5
				_ = _1315_cf7
				_1259_v10 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1258_v9, _1259_v10), _dafny.Companion_Sequence_.Concatenate(_1258_v9, _1259_v10))
				var _1316_v45 D0
				_ = _1316_v45
				_1316_v45 = Companion_D0_.Create_DC1_(true)
				_1316_v45 = Companion_D0_.Create_DC1_((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool))
				var _1317_v46 _dafny.Map
				_ = _1317_v46
				_1317_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool) {
						return _1250_v1
					}
					return _1250_v1
				})(), false)
				_1317_v46 = (_1317_v46).Update(!(Companion_Default___.Fm2((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool), _1261_v12, (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool), (_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool), globalState)), !(!((func() bool {
					if (_1317_v46).Contains(_1250_v1) {
						return (_1317_v46).Get(_1250_v1).(bool)
					}
					return (_this).Fm5(globalState)
				})())))
				var _1318_v47 *C0
				_ = _1318_v47
				var _nw232 *C0 = New_C0_()
				_ = _nw232
				_nw232.Ctor__((_1249_v0).F3())
				_1318_v47 = _nw232
			}
		}
		var _1319_v48 *C0
		_ = _1319_v48
		var _nw233 *C0 = New_C0_()
		_ = _nw233
		_nw233.Ctor__((_1249_v0).F3())
		_1319_v48 = _nw233
		var _source16 D0 = Companion_D0_.Create_DC0_(_1250_v1)
		_ = _source16
		if _source16.Is_DC1() {
			var _1320___mcc_h6 bool = _source16.Get_().(D0_DC1).Cf1
			_ = _1320___mcc_h6
			var _1321_cf1 bool = _1320___mcc_h6
			_ = _1321_cf1
			_1261_v12 = (_this).F0()
			_1261_v12 = ((func() _dafny.Int {
				if !((_1253_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1253_v4), 0))).Int()).(bool)) {
					return _dafny.IntOfInt64(701)
				}
				return (_this).F2()
			})()).Minus(_dafny.IntOfUint32((_1264_v15).Cardinality()))
			var _1322_v50 _dafny.Set
			_ = _1322_v50
			_1322_v50 = _dafny.SetOf(_1321_cf1, _1250_v1)
			var _1323_v51 _dafny.Sequence
			_ = _1323_v51
			_1323_v51 = _dafny.SeqOf(_1322_v50)
			var _1324_v52 _dafny.Map
			_ = _1324_v52
			_1324_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), (_1323_v51).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1323_v51).Cardinality()))).Uint32()).(_dafny.Set))
			var _1325_v53 _dafny.Map
			_ = _1325_v53
			_1325_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3((func() _dafny.Set {
				if (_1324_v52).Contains(_dafny.CodePoint('o')) {
					return (_1324_v52).Get(_dafny.CodePoint('o')).(_dafny.Set)
				}
				return _1322_v50
			})(), (_1249_v0).F3(), (_this).F0(), _1261_v12, globalState), (_this).F0())
			_1261_v12 = ((_this).F2()).Minus(((func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter28 := _dafny.Iterate((_1325_v53).Keys().Elements()); ; {
					_compr_26, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _1326_v49 _dafny.Int
					_1326_v49 = interface{}(_compr_26).(_dafny.Int)
					if (_1325_v53).Contains(_1326_v49) {
						_coll26.Add((_1326_v49).Minus((_dafny.SetOf(false, _1250_v1)).Cardinality()), _1261_v12)
					}
				}
				return _coll26.ToMap()
			}()).Merge(_1325_v53)).Cardinality())
			_1250_v1 = true
		} else {
			var _1327___mcc_h7 bool = _source16.Get_().(D0_DC0).Cf0
			_ = _1327___mcc_h7
			var _1328_cf0 bool = _1327___mcc_h7
			_ = _1328_cf0
			var _1329_v54 _dafny.Array
			_ = _1329_v54
			var _nw234 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
			_ = _nw234
			_1329_v54 = _nw234
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1329_v54), 0))
			_ = _index224
			(_1329_v54).ArraySet1(_1249_v0, (_index224).Int())
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1329_v54), 0))
			_ = _index225
			(_1329_v54).ArraySet1(_1249_v0, (_index225).Int())
			var _1330_v55 D0
			_ = _1330_v55
			_1330_v55 = Companion_D0_.Create_DC0_(_1328_cf0)
			var _1331_v56 _dafny.Map
			_ = _1331_v56
			_1331_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1330_v55, _1253_v4)
			_1331_v56 = (_1331_v56).Update(_1330_v55, _1253_v4)
			var _1332_v57 _dafny.Array
			_ = _1332_v57
			var _nw235 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw235
			_1332_v57 = _nw235
			var _1333_v58 _dafny.Map
			_ = _1333_v58
			_1333_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), _1261_v12)
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1332_v57), 0))
			_ = _index226
			(_1332_v57).ArraySet1(((_1333_v58).Merge(_1333_v58)).Cardinality(), (_index226).Int())
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1332_v57), 0))
			_ = _index227
			(_1332_v57).ArraySet1(_dafny.IntOfInt64(929), (_index227).Int())
			_1250_v1 = (_this).Fm5(globalState)
		}
	}
}
func (_this *C5) M1(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1334_v0 bool
		_ = _1334_v0
		_1334_v0 = false
		var _1335_v1 _dafny.Set
		_ = _1335_v1
		_1335_v1 = _dafny.SetOf(_1334_v0, _1334_v0, _1334_v0)
		var _hi10 _dafny.Int = Companion_Default___.Fm3(_1335_v1, (_this).F1(), (_this).F2(), _dafny.IntOfInt64(2), globalState)
		_ = _hi10
		for _1336_i0 := p1; _1336_i0.Cmp(_hi10) < 0; _1336_i0 = _1336_i0.Plus(_dafny.One) {
			if _1334_v0 {
				var _1337_v2 _dafny.Set
				_ = _1337_v2
				_1337_v2 = _dafny.SetOf((_this).F1())
				r0 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_dafny.SetOf(_dafny.CodePoint('g'))).IsProperSubsetOf(_1337_v2) {
						return (_this).F0()
					}
					return p1
				})())
				var _1338_v3 D3
				_ = _1338_v3
				_1338_v3 = Companion_D3_.Create_DC8_(true, (_this).F2())
				var _1339_v4 _dafny.Sequence
				_ = _1339_v4
				_1339_v4 = _dafny.UnicodeSeqOfUtf8Bytes("gmdnh")
				var _1340_v5 _dafny.Map
				_ = _1340_v5
				_1340_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1338_v3).Dtor_cf9(), _dafny.IntOfUint32((_1339_v4).Cardinality()))), ((_1335_v1).Difference(_1335_v1)).Cardinality())
				_1340_v5 = (_1340_v5).Update(_1336_i0, (_this).F0())
				var _1341_v6 *C0
				_ = _1341_v6
				var _nw236 *C0 = New_C0_()
				_ = _nw236
				_nw236.Ctor__((_this).F1())
				_1341_v6 = _nw236
				_1334_v0 = _1334_v0
				r1 = Companion_Default___.SafeModuloInt(_1336_i0, (_this).F0())
			} else {
				var _1342_v7 _dafny.Array
				_ = _1342_v7
				var _nwElement0_37 _dafny.Int = (func() _dafny.Int {
					if _1334_v0 {
						return _1336_i0
					}
					return _1336_i0
				})()
				_ = _nwElement0_37
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.One)
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_37, 0)
				_1342_v7 = _nw237
				var _1343_v8 T0
				_ = _1343_v8
				var _nw238 *C3 = New_C3_()
				_ = _nw238
				_nw238.Ctor__((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F2())))
				_1343_v8 = _nw238
				var _1344_v9 _dafny.Sequence
				_ = _1344_v9
				_1344_v9 = _dafny.SeqOf(_1343_v8, _1343_v8, _1343_v8, _1343_v8)
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1342_v7), 0))
				_ = _index228
				(_1342_v7).ArraySet1(_dafny.IntOfUint32((_1344_v9).Cardinality()), (_index228).Int())
				var _1345_v10 _dafny.Sequence
				_ = _1345_v10
				_1345_v10 = _dafny.SeqOf((_this).F2(), p1, (_this).F2(), (_this).F2())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1342_v7), 0))
				_ = _index229
				(_1342_v7).ArraySet1((_1345_v10).Select((Companion_Default___.SafeIndex(((_this).F0()).Minus(_dafny.IntOfInt64(266)), _dafny.IntOfUint32((_1345_v10).Cardinality()))).Uint32()).(_dafny.Int), (_index229).Int())
				var _1346_v11 _dafny.Map
				_ = _1346_v11
				_1346_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v0, _1334_v0)
				var _1347_v12 _dafny.Map
				_ = _1347_v12
				_1347_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v0, _dafny.CodePoint('p'))
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1342_v7), 0))
				_ = _index230
				var _rhs224 bool = (func() bool {
					if _1334_v0 {
						return _1334_v0
					}
					return (_1335_v1).IsProperSubsetOf(_1335_v1)
				})()
				_ = _rhs224
				var _rhs225 bool = false
				_ = _rhs225
				var _rhs226 _dafny.Int = (((_1346_v11).Merge(Companion_Default___.Fm20(_1334_v0, _1347_v12, _1334_v0, (_this).F1(), globalState))).Cardinality()).Plus(((_this).F0()).Times(_1336_i0))
				_ = _rhs226
				var _rhs227 _dafny.Int = _dafny.IntOfInt64(-863)
				_ = _rhs227
				var _rhs228 _dafny.Array = _1342_v7
				_ = _rhs228
				var _lhs102 _dafny.Array = _1342_v7
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1342_v7), 0))
				_ = _lhs103
				_1334_v0 = _rhs224
				_1334_v0 = _rhs225
				r0 = _rhs226
				(_lhs102).ArraySet1(_rhs227, (_lhs103).Int())
				_1342_v7 = _rhs228
				var _1348_v13 D0
				_ = _1348_v13
				_1348_v13 = Companion_D0_.Create_DC1_(_1334_v0)
				var _1349_v14 D2
				_ = _1349_v14
				_1349_v14 = Companion_D2_.Create_DC6_(Companion_D2_.Create_DC3_(_dafny.SeqOf(_1348_v13, _1348_v13)))
				var _1350_v15 D2
				_ = _1350_v15
				_1350_v15 = Companion_D2_.Create_DC6_(_1349_v14)
				var _1351_v16 _dafny.Map
				_ = _1351_v16
				_1351_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1350_v15, p1)
				_1334_v0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1350_v15, (_1343_v8).F0())).Equals(_1351_v16)
				var _1352_v17 *C2
				_ = _1352_v17
				var _nw239 *C2 = New_C2_()
				_ = _nw239
				_nw239.Ctor__(_dafny.IntOfInt64(409))
				_1352_v17 = _nw239
				var _1353_v18 _dafny.Sequence
				_ = _1353_v18
				_1353_v18 = _dafny.UnicodeSeqOfUtf8Bytes("paqg")
				var _1354_v19 _dafny.Array
				_ = _1354_v19
				var _nwElement0_38 _dafny.Sequence = _1353_v18
				_ = _nwElement0_38
				var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(22))
				_ = _nw240
				(_nw240).ArraySet1(_nwElement0_38, 0)
				(_nw240).ArraySet1(_1353_v18, 1)
				(_nw240).ArraySet1(Companion_Default___.Fm18(p1, globalState), 2)
				(_nw240).ArraySet1(_1353_v18, 3)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1353_v18, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("hibaahddh"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hibaahddh")).Cardinality()))).Uint32(), (_this).F1())), 4)
				(_nw240).ArraySet1(_1353_v18, 5)
				(_nw240).ArraySet1(_1353_v18, 6)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("n"), 7)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gvptickwg"), 8)
				(_nw240).ArraySet1(_1353_v18, 9)
				(_nw240).ArraySet1(_1353_v18, 10)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("g"), _1353_v18), 11)
				(_nw240).ArraySet1(_1353_v18, 12)
				(_nw240).ArraySet1(_1353_v18, 13)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cidhpdr"), 14)
				(_nw240).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_1355_v18 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_1356_i1 _dafny.Int) _dafny.CodePoint {
						return (_1355_v18).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_1355_v18).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_1353_v18))), 15)
				(_nw240).ArraySet1(_1353_v18, 16)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("abkkeaax"), 17)
				(_nw240).ArraySet1(_1353_v18, 18)
				(_nw240).ArraySet1(_1353_v18, 19)
				(_nw240).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}(func(_1357_i2 _dafny.Int) _dafny.CodePoint {
					return (_this).F1()
				})), 20)
				(_nw240).ArraySet1(_1353_v18, 21)
				_1354_v19 = _nw240
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1354_v19), 0))
				_ = _index231
				(_1354_v19).ArraySet1(_1353_v18, (_index231).Int())
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1354_v19), 0))
				_ = _index232
				(_1354_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1353_v18, _1353_v18), (_index232).Int())
			}
			var _1358_v20 _dafny.Sequence
			_ = _1358_v20
			_1358_v20 = _dafny.UnicodeSeqOfUtf8Bytes("j")
			var _1359_v21 _dafny.Map
			_ = _1359_v21
			_1359_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), _1358_v20)
			var _1360_v22 _dafny.MultiSet
			_ = _1360_v22
			_1360_v22 = _dafny.MultiSetOf(true)
			var _1361_v23 _dafny.MultiSet
			_ = _1361_v23
			_1361_v23 = _dafny.MultiSetOf(_dafny.IntOfInt64(749), p1)
			var _1362_v24 _dafny.Sequence
			_ = _1362_v24
			_1362_v24 = _dafny.SeqOf(_1334_v0)
			var _1363_v25 D10
			_ = _1363_v25
			_1363_v25 = Companion_D10_.Create_DC22_(Companion_Default___.Fm3(Companion_Default___.Fm10(_dafny.MultiSetOf((_this).F1()), _1334_v0, globalState), (_this).F1(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1359_v21).Contains((_1360_v22).Cardinality()) {
					return (_1359_v21).Get((_1360_v22).Cardinality()).(_dafny.Sequence)
				}
				return Companion_Default___.Fm19(_1361_v23, _dafny.IntOfUint32((_1362_v24).Cardinality()), _1336_i0, p1, globalState)
			})()).Cardinality()), _1336_i0, globalState), (_1334_v0) && (_1334_v0), _1334_v0)
			_1363_v25 = _1363_v25
			r2 = (_dafny.Zero).Minus((_this).F2())
			_1334_v0 = (func() bool {
				if _1334_v0 {
					return !(_1334_v0)
				}
				return _1334_v0
			})()
		}
		var _hi11 _dafny.Int = ((_this).F0()).Minus(p1)
		_ = _hi11
		for _1364_i3 := (_this).F0(); _1364_i3.Cmp(_hi11) < 0; _1364_i3 = _1364_i3.Plus(_dafny.One) {
			var _1365_v26 _dafny.Array
			_ = _1365_v26
			var _nw241 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
			_ = _nw241
			_1365_v26 = _nw241
			var _1366_v27 _dafny.MultiSet
			_ = _1366_v27
			_1366_v27 = _dafny.MultiSetOf(_1334_v0)
			var _1367_v28 _dafny.Map
			_ = _1367_v28
			_1367_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1366_v27, !(_1334_v0))
			var _1368_v29 T0
			_ = _1368_v29
			var _nw242 *C1 = New_C1_()
			_ = _nw242
			_nw242.Ctor__((_1367_v28).Cardinality())
			_1368_v29 = _nw242
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1365_v26), 0))
			_ = _index233
			(_1365_v26).ArraySet1(_1368_v29, (_index233).Int())
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1365_v26), 0))
			_ = _index234
			(_1365_v26).ArraySet1(_1368_v29, (_index234).Int())
			var _1369_v30 _dafny.Sequence
			_ = _1369_v30
			_1369_v30 = _dafny.UnicodeSeqOfUtf8Bytes("yyrnjblm")
			_1369_v30 = _1369_v30
			r1 = (_this).F0()
			var _1370_v31 _dafny.Sequence
			_ = _1370_v31
			_1370_v31 = _dafny.SeqOf(_1334_v0, true)
			var _1371_v32 _dafny.Map
			_ = _1371_v32
			_1371_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), _1364_i3)
			var _1372_v33 _dafny.Sequence
			_ = _1372_v33
			_1372_v33 = _dafny.SeqOf(p1, (_this).F0())
			var _1373_v34 _dafny.Sequence
			_ = _1373_v34
			_1373_v34 = _dafny.SeqOf(_1372_v33)
			var _1374_v35 _dafny.Map
			_ = _1374_v35
			_1374_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v0, _1334_v0)
			var _1375_v36 _dafny.Map
			_ = _1375_v36
			_1375_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-707), true)
			var _1376_v37 _dafny.Array
			_ = _1376_v37
			var _nwElement0_39 _dafny.MultiSet = (_1366_v27).Difference(_1366_v27)
			_ = _nwElement0_39
			var _nw243 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(20))
			_ = _nw243
			(_nw243).ArraySet1(_nwElement0_39, 0)
			(_nw243).ArraySet1((_1366_v27).Union(_1366_v27), 1)
			(_nw243).ArraySet1((_dafny.MultiSetOf(_1334_v0, _1334_v0)).Union(_1366_v27), 2)
			(_nw243).ArraySet1((_1366_v27).Update(_1334_v0, Companion_Default___.Abs(_1364_i3)), 3)
			(_nw243).ArraySet1((_dafny.MultiSetFromSeq(_1370_v31)).Difference(_1366_v27), 4)
			(_nw243).ArraySet1(_1366_v27, 5)
			(_nw243).ArraySet1(_1366_v27, 6)
			(_nw243).ArraySet1((_dafny.MultiSetFromSeq(_1370_v31)).Update(_1334_v0, Companion_Default___.Abs((func() _dafny.Int {
				if (_1371_v32).Contains(_1364_i3) {
					return (_1371_v32).Get(_1364_i3).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_1373_v34).Cardinality())
			})())), 7)
			(_nw243).ArraySet1(_dafny.MultiSetOf(_1334_v0), 8)
			(_nw243).ArraySet1(_dafny.MultiSetOf(_1334_v0, _1334_v0, _1334_v0), 9)
			(_nw243).ArraySet1(_dafny.MultiSetOf(_1334_v0), 10)
			(_nw243).ArraySet1((_1366_v27).Update(_1334_v0, Companion_Default___.Abs(_dafny.IntOfInt64(895))), 11)
			(_nw243).ArraySet1(_1366_v27, 12)
			(_nw243).ArraySet1((_1366_v27).Difference(_1366_v27), 13)
			(_nw243).ArraySet1(_1366_v27, 14)
			(_nw243).ArraySet1(_1366_v27, 15)
			(_nw243).ArraySet1(_dafny.MultiSetOf(_1334_v0, _1334_v0, _1334_v0), 16)
			(_nw243).ArraySet1(_1366_v27, 17)
			(_nw243).ArraySet1(_1366_v27, 18)
			(_nw243).ArraySet1((func() _dafny.MultiSet {
				if (func() bool {
					if (_1374_v35).Contains(_1334_v0) {
						return (_1374_v35).Get(_1334_v0).(bool)
					}
					return _1334_v0
				})() {
					return _1366_v27
				}
				return Companion_Default___.Fm28(_1369_v30, true, (_1374_v35).Cardinality(), Companion_Default___.Fm2((func() bool {
					if (_1375_v36).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg61 _dafny.Int) interface{} {
							return coer61(arg61)
						}
					}(func(_1377_i4 _dafny.Int) _dafny.CodePoint {
						return (_this).F1()
					}))).Cardinality())) {
						return (_1375_v36).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg62 _dafny.Int) interface{} {
								return coer62(arg62)
							}
						}(func(_1378_i4 _dafny.Int) _dafny.CodePoint {
							return (_this).F1()
						}))).Cardinality())).(bool)
					}
					return _1334_v0
				})(), p1, _1334_v0, _1334_v0, globalState), globalState)
			})(), 19)
			_1376_v37 = _nw243
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1376_v37), 0))
			_ = _index235
			(_1376_v37).ArraySet1(_dafny.MultiSetOf((_1368_v29).Fm5(globalState)), (_index235).Int())
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1376_v37), 0))
			_ = _index236
			(_1376_v37).ArraySet1(_1366_v27, (_index236).Int())
		}
		var _1379_v38 _dafny.Map
		_ = _1379_v38
		_1379_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), (_this).F2())
		var _1380_v39 _dafny.Set
		_ = _1380_v39
		_1380_v39 = _dafny.SetOf((_1379_v38).Cardinality(), (_this).F2())
		var _1381_i5 _dafny.Int
		_ = _1381_i5
		_1381_i5 = _dafny.Zero
		{
			for !(_1334_v0) || ((_1380_v39).Contains((_this).F0())) {
				{
					if (_1381_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1381_i5 = (_1381_i5).Plus(_dafny.One)
					_1334_v0 = !(((_this).F0()).Cmp((_this).F2()) < 0)
					var _1382_v40 _dafny.Sequence
					_ = _1382_v40
					_1382_v40 = _dafny.UnicodeSeqOfUtf8Bytes("cjy")
					var _1383_v41 _dafny.Map
					_ = _1383_v41
					_1383_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v0, p1), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1382_v40).Cardinality())))
					var _1384_v42 _dafny.Map
					_ = _1384_v42
					_1384_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v0, (_this).F0())
					_1383_v41 = (_1383_v41).Update(_1384_v42, (_this).F2())
					_1334_v0 = _1334_v0
					if _1334_v0 {
						var _1385_v43 _dafny.Array
						_ = _1385_v43
						var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
						_ = _nw244
						_1385_v43 = _nw244
						var _1386_v44 _dafny.Array
						_ = _1386_v44
						var _nw245 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
						_ = _nw245
						_1386_v44 = _nw245
						var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_1385_v43), 0))
						_ = _index237
						(_1385_v43).ArraySet1(_1386_v44, (_index237).Int())
						var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_1385_v43), 0))
						_ = _index238
						(_1385_v43).ArraySet1(_1386_v44, (_index238).Int())
						var _1387_v45 _dafny.Map
						_ = _1387_v45
						_1387_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1382_v40, (_this).F2())
						var _1388_v46 _dafny.MultiSet
						_ = _1388_v46
						_1388_v46 = _dafny.MultiSetOf((_this).F0(), Companion_Default___.Fm3(_1335_v1, _dafny.CodePoint('p'), (func() _dafny.Int {
							if (_1387_v45).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg63 _dafny.Int) interface{} {
									return coer63(arg63)
								}
							}(func(_1389_i6 _dafny.Int) _dafny.CodePoint {
								return (_this).F1()
							}))) {
								return (_1387_v45).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg64 _dafny.Int) interface{} {
										return coer64(arg64)
									}
								}(func(_1390_i6 _dafny.Int) _dafny.CodePoint {
									return (_this).F1()
								}))).(_dafny.Int)
							}
							return p1
						})(), _dafny.IntOfInt64(328), globalState))
						var _1391_v47 _dafny.MultiSet
						_ = _1391_v47
						_1391_v47 = _dafny.MultiSetOf(_1388_v46, _1388_v46)
						r0 = ((_dafny.MultiSetOf(_dafny.MultiSetOf((_this).F0()), _1388_v46, _1388_v46, _dafny.MultiSetOf((_this).F2()), _1388_v46)).Difference(_1391_v47)).Cardinality()
						var _1392_v48 *C0
						_ = _1392_v48
						var _nw246 *C0 = New_C0_()
						_ = _nw246
						_nw246.Ctor__((_this).F1())
						_1392_v48 = _nw246
						_1382_v40 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-405))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}((func(_1393_v48 *C0) func(_dafny.Int) _dafny.CodePoint {
							return func(_1394_i7 _dafny.Int) _dafny.CodePoint {
								return (_1393_v48).F3()
							}
						})(_1392_v48)))
						_1380_v39 = func() _dafny.Set {
							var _coll27 = _dafny.NewBuilder()
							_ = _coll27
							for _iter29 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), p1)).Merge(func() _dafny.Map {
								var _coll28 = _dafny.NewMapBuilder()
								_ = _coll28
								for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(944), _dafny.IntOfInt64(-345))); ; {
									_compr_28, _ok30 := _iter30()
									if !_ok30 {
										break
									}
									var _1395_v49 _dafny.Int
									_1395_v49 = interface{}(_compr_28).(_dafny.Int)
									if ((_dafny.IntOfInt64(944)).Cmp(_1395_v49) <= 0) && ((_1395_v49).Cmp(_dafny.IntOfInt64(-345)) < 0) {
										_coll28.Add((_1395_v49).Times(p1), (_this).F2())
									}
								}
								return _coll28.ToMap()
							}())).Keys().Elements()); ; {
								_compr_27, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _1396_v50 _dafny.Int
								_1396_v50 = interface{}(_compr_27).(_dafny.Int)
								if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), p1)).Merge(func() _dafny.Map {
									var _coll29 = _dafny.NewMapBuilder()
									_ = _coll29
									for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(944), _dafny.IntOfInt64(-345))); ; {
										_compr_29, _ok31 := _iter31()
										if !_ok31 {
											break
										}
										var _1395_v49 _dafny.Int
										_1395_v49 = interface{}(_compr_29).(_dafny.Int)
										if ((_dafny.IntOfInt64(944)).Cmp(_1395_v49) <= 0) && ((_1395_v49).Cmp(_dafny.IntOfInt64(-345)) < 0) {
											_coll29.Add((_1395_v49).Times(p1), (_this).F2())
										}
									}
									return _coll29.ToMap()
								}())).Contains(_1396_v50) {
									_coll27.Add((_1396_v50).Plus((_dafny.Zero).Minus(((func() _dafny.Map {
										if !(false) {
											return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(204), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(872))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
												return func(arg66 _dafny.Int) interface{} {
													return coer66(arg66)
												}
											}(func(_1397_i8 _dafny.Int) _dafny.CodePoint {
												return _dafny.CodePoint('o')
											}))).Cardinality()))
										}
										return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(968), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pp")).Cardinality()))).Cardinality())).Cardinality()))
									})()).Cardinality())))
								}
							}
							return _coll27.ToSet()
						}()
					} else {
						var _1398_v51 *C1
						_ = _1398_v51
						var _nw247 *C1 = New_C1_()
						_ = _nw247
						_nw247.Ctor__(_dafny.IntOfInt64(580))
						_1398_v51 = _nw247
						_1398_v51 = _1398_v51
						var _1399_v52 T0
						_ = _1399_v52
						var _nw248 *C3 = New_C3_()
						_ = _nw248
						_nw248.Ctor__((_this).F0())
						_1399_v52 = _nw248
						var _1400_v53 _dafny.Map
						_ = _1400_v53
						_1400_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_1335_v1, (_this).F1(), (_this).F2(), _dafny.IntOfUint32((_1382_v40).Cardinality()), globalState), _1399_v52)
						var _1401_v54 _dafny.Sequence
						_ = _1401_v54
						_1401_v54 = _dafny.SeqOf(p1)
						var _1402_v55 _dafny.Sequence
						_ = _1402_v55
						_1402_v55 = _dafny.SeqOf(_1334_v0, _1334_v0)
						var _1403_v56 _dafny.Array
						_ = _1403_v56
						var _nwElement0_40 _dafny.Int = (_this).F2()
						_ = _nwElement0_40
						var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(11))
						_ = _nw249
						(_nw249).ArraySet1(_nwElement0_40, 0)
						(_nw249).ArraySet1((_1400_v53).Cardinality(), 1)
						(_nw249).ArraySet1((_this).F2(), 2)
						(_nw249).ArraySet1((_dafny.Zero).Minus((_1401_v54).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1401_v54).Cardinality()))).Uint32()).(_dafny.Int)), 3)
						(_nw249).ArraySet1(p1, 4)
						(_nw249).ArraySet1((_this).F2(), 5)
						(_nw249).ArraySet1(Companion_Default___.SafeDivisionInt(p1, (_1379_v38).Cardinality()), 6)
						(_nw249).ArraySet1((_this).F2(), 7)
						(_nw249).ArraySet1((_dafny.Zero).Minus((_this).F2()), 8)
						(_nw249).ArraySet1(((_this).F0()).Plus((_dafny.SetOf(_1334_v0, _1334_v0)).Cardinality()), 9)
						(_nw249).ArraySet1((p1).Minus((_1398_v51).Fm15(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg67 _dafny.Int) interface{} {
								return coer67(arg67)
							}
						}((func(_1404_v1 _dafny.Set) func(_dafny.Int) _dafny.CodePoint {
							return func(_1405_i9 _dafny.Int) _dafny.CodePoint {
								return (Companion_D5_.Create_DC13_(_dafny.IntOfInt64(209), _1404_v1, (_this).F1(), (_this).F2(), true)).Dtor_cf19()
							}
						})(_1335_v1))), (_this).Fm5(globalState)), (_1402_v55).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_1402_v55).Cardinality()))).Uint32()).(bool), true, globalState)), 10)
						_1403_v56 = _nw249
						var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1403_v56), 0))
						_ = _index239
						(_1403_v56).ArraySet1(((_this).F2()).Minus(p1), (_index239).Int())
						var _1406_v57 _dafny.Map
						_ = _1406_v57
						_1406_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1382_v40, _1334_v0)
						var _1407_v58 _dafny.MultiSet
						_ = _1407_v58
						_1407_v58 = _dafny.MultiSetOf(_1334_v0, (_1402_v55).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1402_v55).Cardinality()))).Uint32()).(bool), _1334_v0)
						var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1403_v56), 0))
						_ = _index240
						var _rhs229 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1401_v54).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1401_v54).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg68 _dafny.Int) interface{} {
								return coer68(arg68)
							}
						}(func(_1408_i10 _dafny.Int) _dafny.CodePoint {
							return (_this).F1()
						}))).Cardinality())))
						_ = _rhs229
						var _rhs230 _dafny.Int = (_1398_v51).Fm15(_1406_v57, false, ((_1379_v38).Update((_1399_v52).F0(), (_this).F0())).Contains((func() _dafny.Int {
							if (_1407_v58).Contains(true) {
								return (_1407_v58).Multiplicity(true)
							}
							return Companion_Default___.Fm3(_1335_v1, (_this).F1(), (_this).F2(), (_1399_v52).F0(), globalState)
						})()), globalState)
						_ = _rhs230
						var _rhs231 _dafny.Int = (_1399_v52).F0()
						_ = _rhs231
						var _lhs104 _dafny.Array = _1403_v56
						_ = _lhs104
						var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1403_v56), 0))
						_ = _lhs105
						r2 = _rhs229
						(_lhs104).ArraySet1(_rhs230, (_lhs105).Int())
						r0 = _rhs231
						var _1409_v59 D6
						_ = _1409_v59
						_1409_v59 = Companion_D6_.Create_DC14_(_1403_v56)
						_1409_v59 = _1409_v59
						_1403_v56 = _1403_v56
						_1334_v0 = _1334_v0
					}
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1410_v60 _dafny.Sequence
		_ = _1410_v60
		_1410_v60 = _dafny.UnicodeSeqOfUtf8Bytes("jiaaicmv")
		var _1411_v61 _dafny.Map
		_ = _1411_v61
		_1411_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1410_v60, _1334_v0)
		_1411_v61 = (_1411_v61).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_1412_i11 _dafny.Int) _dafny.CodePoint {
			return (_this).F1()
		})), _1334_v0)
		var _1413_v62 T1
		_ = _1413_v62
		var _nw250 *C2 = New_C2_()
		_ = _nw250
		_nw250.Ctor__((_this).F0())
		_1413_v62 = _nw250
		var _1414_v63 _dafny.Map
		_ = _1414_v63
		_1414_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _1413_v62)
		var _1415_v64 D6
		_ = _1415_v64
		_1415_v64 = Companion_D6_.Create_DC15_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(299))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}((func(_1416_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1417_i12 _dafny.Int) _dafny.Int {
				return _1416_p1
			}
		})(p1))), Companion_Default___.Fm0(globalState))).Cardinality()), (_1414_v63).Equals(_1414_v63), (_this).F0())
		var _source17 D6 = _1415_v64
		_ = _source17
		if _source17.Is_DC15() {
			var _1418___mcc_h0 _dafny.Int = _source17.Get_().(D6_DC15).Cf23
			_ = _1418___mcc_h0
			var _1419___mcc_h1 bool = _source17.Get_().(D6_DC15).Cf24
			_ = _1419___mcc_h1
			var _1420___mcc_h2 _dafny.Int = _source17.Get_().(D6_DC15).Cf25
			_ = _1420___mcc_h2
			var _1421_cf25 _dafny.Int = _1420___mcc_h2
			_ = _1421_cf25
			var _1422_cf24 bool = _1419___mcc_h1
			_ = _1422_cf24
			var _1423_cf23 _dafny.Int = _1418___mcc_h0
			_ = _1423_cf23
			var _1424_v65 *C1
			_ = _1424_v65
			var _nw251 *C1 = New_C1_()
			_ = _nw251
			_nw251.Ctor__((_1413_v62).F0())
			_1424_v65 = _nw251
			var _source18 _dafny.CodePoint = (_this).F1()
			_ = _source18
			var _1425___mcc_h4 _dafny.CodePoint = _source18
			_ = _1425___mcc_h4
			var _1426_cf13 _dafny.CodePoint = _1425___mcc_h4
			_ = _1426_cf13
			var _1427_v66 D2
			_ = _1427_v66
			_1427_v66 = Companion_D2_.Create_DC4_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1410_v60, (Companion_Default___.SafeIndex(Companion_Default___.Fm3(_1335_v1, (_this).F1(), (_this).F2(), _1421_cf25, globalState), _dafny.IntOfUint32((_1410_v60).Cardinality()))).Uint32(), (_1410_v60).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1410_v60).Cardinality()))).Uint32()).(_dafny.CodePoint)), _1410_v60), !((_1334_v0) == (_1334_v0)))
			_1427_v66 = _1427_v66
			var _pat_let_tv32 = _1410_v60
			_ = _pat_let_tv32
			var _pat_let_tv33 = _1410_v60
			_ = _pat_let_tv33
			var _pat_let_tv34 = _1426_cf13
			_ = _pat_let_tv34
			var _rhs232 D2 = func(_pat_let25_0 D2) D2 {
				return func(_1428_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let26_0 _dafny.Sequence) D2 {
						return func(_1429_dt__update_hcf4_h0 _dafny.Sequence) D2 {
							return Companion_D2_.Create_DC4_(_1429_dt__update_hcf4_h0, (_1428_dt__update__tmp_h0).Dtor_cf5())
						}(_pat_let26_0)
					}(_dafny.Companion_Sequence_.Update(_pat_let_tv32, (Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_pat_let_tv33).Cardinality()))).Uint32(), _pat_let_tv34))
				}(_pat_let25_0)
			}(Companion_D2_.Create_DC4_(_1410_v60, _1334_v0))
			_ = _rhs232
			var _rhs233 bool = _1334_v0
			_ = _rhs233
			_1427_v66 = _rhs232
			_1334_v0 = _rhs233
			_1334_v0 = _1422_cf24
			var _1430_v67 _dafny.Array
			_ = _1430_v67
			var _nw252 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
			_ = _nw252
			_1430_v67 = _nw252
			var _1431_v68 D3
			_ = _1431_v68
			_1431_v68 = Companion_D3_.Create_DC9_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1334_v0), (_1413_v62).F0())).Cardinality(), _1430_v67, _1334_v0)
			_1379_v38 = (_1379_v38).Update(_1421_cf25, (_1431_v68).Dtor_cf10())
			var _1432_v69 _dafny.MultiSet
			_ = _1432_v69
			_1432_v69 = _dafny.MultiSetOf(!(_1422_cf24))
			r2 = Companion_Default___.Fm3((_1335_v1).Difference(_dafny.SetOf(_1334_v0)), (_this).F1(), (_dafny.Zero).Minus((_1424_v65).Fm15(_1411_v61, _1334_v0, _1422_cf24, globalState)), (func() _dafny.Int {
				if _1334_v0 {
					return _dafny.IntOfInt64(579)
				}
				return (func() _dafny.Int {
					if (_1432_v69).Contains((_1424_v65).Fm5(globalState)) {
						return (_1432_v69).Multiplicity((_1424_v65).Fm5(globalState))
					}
					return (_this).F2()
				})()
			})(), globalState)
			_1379_v38 = _1379_v38
		} else {
			var _1433___mcc_h3 _dafny.Array = _source17.Get_().(D6_DC14).Cf22
			_ = _1433___mcc_h3
			var _1434_cf22 _dafny.Array = _1433___mcc_h3
			_ = _1434_cf22
			var _1435_v70 _dafny.Array
			_ = _1435_v70
			var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
			_ = _nw253
			_1435_v70 = _nw253
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1434_cf22), 0))
			_ = _index241
			(_1434_cf22).ArraySet1(_dafny.IntOfInt64(-340), (_index241).Int())
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1434_cf22), 0))
			_ = _index242
			var _rhs234 _dafny.Array = _1435_v70
			_ = _rhs234
			var _rhs235 _dafny.Int = ((_this).F0()).Minus(Companion_Default___.SafeModuloInt((_1413_v62).F0(), p1))
			_ = _rhs235
			var _rhs236 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("e")
			_ = _rhs236
			var _lhs106 _dafny.Array = _1434_cf22
			_ = _lhs106
			var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1434_cf22), 0))
			_ = _lhs107
			_1435_v70 = _rhs234
			(_lhs106).ArraySet1(_rhs235, (_lhs107).Int())
			_1410_v60 = _rhs236
			_1334_v0 = _1334_v0
			var _1436_v71 _dafny.MultiSet
			_ = _1436_v71
			_1436_v71 = _dafny.MultiSetOf((_1413_v62).F0(), _dafny.IntOfUint32((_1410_v60).Cardinality()))
			_1335_v1 = _dafny.SetOf(_1334_v0, _1334_v0, (_1436_v71).IsSubsetOf(_1436_v71))
			var _1437_v72 _dafny.Map
			_ = _1437_v72
			_1437_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp((_this).F2()) >= 0, _1334_v0)
			_1437_v72 = (func() _dafny.Map {
				if _1334_v0 {
					return _1437_v72
				}
				return _1437_v72
			})()
		}
		var _1438_v73 _dafny.Array
		_ = _1438_v73
		var _nw254 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
		_ = _nw254
		_1438_v73 = _nw254
		var _1439_v74 _dafny.Map
		_ = _1439_v74
		_1439_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v0, _1438_v73)
		var _1440_v75 _dafny.MultiSet
		_ = _1440_v75
		_1440_v75 = _dafny.MultiSetOf(p1)
		_1438_v73 = (func() _dafny.Array {
			if (_1439_v74).Contains(!(_1440_v75).Equals(_1440_v75)) {
				return (_1439_v74).Get(!(_1440_v75).Equals(_1440_v75)).(_dafny.Array)
			}
			return _1438_v73
		})()
		var _1441_v76 _dafny.Sequence
		_ = _1441_v76
		_1441_v76 = _dafny.SeqOf(p1)
		r0 = ((func() _dafny.Int {
			if (_1379_v38).Contains(_dafny.IntOfInt64(836)) {
				return (_1379_v38).Get(_dafny.IntOfInt64(836)).(_dafny.Int)
			}
			return (_dafny.Zero).Minus((_1441_v76).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1441_v76).Cardinality()))).Uint32()).(_dafny.Int))
		})()).Times(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_1379_v38).Cardinality()), (_this).F2()))
		r1 = (_this).F0()
		r2 = (_1413_v62).F0()
		return r0, r1, r2
	}
}
func (_this *C5) M3(p0 _dafny.Int, p1 T1, p2 _dafny.Sequence, globalState *GlobalState) (bool, D2) {
	{
		var r0 bool = false
		_ = r0
		var r1 D2 = Companion_D2_.Default()
		_ = r1
		var _1442_v0 bool
		_ = _1442_v0
		_1442_v0 = true
		if _1442_v0 {
			var _1443_v1 _dafny.Set
			_ = _1443_v1
			_1443_v1 = _dafny.SetOf(Companion_Default___.Fm2(!(_1442_v0), (p1).F0(), _1442_v0, _1442_v0, globalState), false, (func() bool {
				if _1442_v0 {
					return _1442_v0
				}
				return _1442_v0
			})(), !(_1442_v0), _1442_v0)
			var _1444_v2 _dafny.Sequence
			_ = _1444_v2
			_1444_v2 = _dafny.SeqOf(_dafny.SetOf(_1442_v0))
			var _1445_v3 _dafny.Sequence
			_ = _1445_v3
			_1445_v3 = _dafny.UnicodeSeqOfUtf8Bytes("hxmu")
			_1443_v1 = (_1444_v2).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3(_dafny.SetOf(false, _1442_v0, _1442_v0), (_this).F1(), _dafny.IntOfUint32((_1445_v3).Cardinality()), (_this).F0(), globalState), _dafny.IntOfUint32((_1444_v2).Cardinality()))).Uint32()).(_dafny.Set)
			var _1446_v4 _dafny.Set
			_ = _1446_v4
			_1446_v4 = _dafny.SetOf(p0, p0)
			var _1447_v5 _dafny.MultiSet
			_ = _1447_v5
			_1447_v5 = _dafny.MultiSetOf((_this).F0(), (_1446_v4).Cardinality())
			var _1448_v6 _dafny.Set
			_ = _1448_v6
			_1448_v6 = _dafny.SetOf(_1447_v5, _dafny.MultiSetOf(_dafny.IntOfInt64(-379)), _1447_v5)
			_1448_v6 = (_1448_v6).Difference(_1448_v6)
			var _1449_v7 _dafny.Array
			_ = _1449_v7
			var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw255
			_1449_v7 = _nw255
			var _1450_v10 _dafny.Array
			_ = _1450_v10
			var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw256
			_1450_v10 = _nw256
			var _1451_v11 _dafny.Map
			_ = _1451_v11
			_1451_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3((_1444_v2).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_1444_v2).Cardinality()))).Uint32()).(_dafny.Set), (_this).F1(), (_this).F2(), (func() _dafny.Set {
				var _coll30 = _dafny.NewBuilder()
				_ = _coll30
				for _iter32 := _dafny.Iterate(((func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(758), _dafny.IntOfInt64(-84))); ; {
						_compr_31, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _1452_v8 _dafny.Int
						_1452_v8 = interface{}(_compr_31).(_dafny.Int)
						if ((_dafny.IntOfInt64(758)).Cmp(_1452_v8) <= 0) && ((_1452_v8).Cmp(_dafny.IntOfInt64(-84)) < 0) {
							_coll31.Add(Companion_Default___.SafeModuloInt(_1452_v8, (_dafny.MultiSetOf((_this).F1())).Cardinality()), _1442_v0)
						}
					}
					return _coll31.ToMap()
				}()).Update(_dafny.IntOfInt64(648), _1442_v0)).Keys().Elements()); ; {
					_compr_30, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _1453_v9 _dafny.Int
					_1453_v9 = interface{}(_compr_30).(_dafny.Int)
					if ((func() _dafny.Map {
						var _coll32 = _dafny.NewMapBuilder()
						_ = _coll32
						for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(758), _dafny.IntOfInt64(-84))); ; {
							_compr_32, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _1452_v8 _dafny.Int
							_1452_v8 = interface{}(_compr_32).(_dafny.Int)
							if ((_dafny.IntOfInt64(758)).Cmp(_1452_v8) <= 0) && ((_1452_v8).Cmp(_dafny.IntOfInt64(-84)) < 0) {
								_coll32.Add(Companion_Default___.SafeModuloInt(_1452_v8, (_dafny.MultiSetOf((_this).F1())).Cardinality()), _1442_v0)
							}
						}
						return _coll32.ToMap()
					}()).Update(_dafny.IntOfInt64(648), _1442_v0)).Contains(_1453_v9) {
						_coll30.Add((_1453_v9).Plus((_dafny.SetOf(_dafny.IntOfInt64(989), _dafny.IntOfInt64(-246))).Cardinality()))
					}
				}
				return _coll30.ToSet()
			}()).Cardinality(), globalState), _1450_v10)
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_1449_v7), 0))
			_ = _index243
			(_1449_v7).ArraySet1((func() _dafny.Array {
				if (_1451_v11).Contains((p1).F0()) {
					return (_1451_v11).Get((p1).F0()).(_dafny.Array)
				}
				return _1450_v10
			})(), (_index243).Int())
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_1449_v7), 0))
			_ = _index244
			(_1449_v7).ArraySet1(_1450_v10, (_index244).Int())
			var _1454_v12 _dafny.Int
			_ = _1454_v12
			_1454_v12 = _dafny.IntOfInt64(399)
			_1454_v12 = ((p1).F0()).Plus(_1454_v12)
			var _1455_v13 _dafny.CodePoint
			_ = _1455_v13
			_1455_v13 = _dafny.CodePoint('g')
			_1455_v13 = _1455_v13
		} else {
			var _1456_v14 _dafny.Sequence
			_ = _1456_v14
			_1456_v14 = _dafny.SeqOf(_1442_v0, _1442_v0)
			var _1457_v15 _dafny.Sequence
			_ = _1457_v15
			_1457_v15 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1442_v0, _1442_v0), _1456_v14), _dafny.SeqOf(_1442_v0))
			_1456_v14 = (_1457_v15).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _1442_v0 {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1442_v0, (_this).F1())).Cardinality()
				}
				return (_dafny.Zero).Minus((_this).F0())
			})(), _dafny.IntOfUint32((_1457_v15).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _1458_v16 _dafny.Sequence
			_ = _1458_v16
			_1458_v16 = _dafny.UnicodeSeqOfUtf8Bytes("biksd")
			var _1459_v17 _dafny.Sequence
			_ = _1459_v17
			_1459_v17 = _dafny.SeqOf(_1458_v16, _1458_v16)
			var _1460_v18 _dafny.Map
			_ = _1460_v18
			_1460_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1442_v0, false)
			var _1461_v19 _dafny.Sequence
			_ = _1461_v19
			_1461_v19 = _dafny.SeqOf(_1459_v17, _1459_v17)
			var _rhs237 bool = !(true) || ((func() bool {
				if (_1460_v18).Contains((_1456_v14).Select((Companion_Default___.SafeIndex((p1).F0(), _dafny.IntOfUint32((_1456_v14).Cardinality()))).Uint32()).(bool)) {
					return (_1460_v18).Get((_1456_v14).Select((Companion_Default___.SafeIndex((p1).F0(), _dafny.IntOfUint32((_1456_v14).Cardinality()))).Uint32()).(bool)).(bool)
				}
				return true
			})())
			_ = _rhs237
			var _rhs238 _dafny.Sequence = (_1461_v19).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1461_v19).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs238
			r0 = _rhs237
			_1459_v17 = _rhs238
			var _1462_v20 _dafny.Array
			_ = _1462_v20
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_32
			var _nw257 _dafny.Array
			_ = _nw257
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw257 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) bool = (func(_1463_v0 bool) func(_dafny.Int) bool {
					return func(_1464_i0 _dafny.Int) bool {
						return _1463_v0
					}
				})(_1442_v0)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw257 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw257).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw257).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1462_v20 = _nw257
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))
			_ = _index245
			(_1462_v20).ArraySet1(false, (_index245).Int())
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))
			_ = _index246
			(_1462_v20).ArraySet1(true, (_index246).Int())
			var _1465_v21 _dafny.Map
			_ = _1465_v21
			_1465_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F1(), _1442_v0)
			var _1466_v22 D7
			_ = _1466_v22
			_1466_v22 = Companion_D7_.Create_DC17_((_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool), !((func() bool {
				if (_1465_v21).Contains((_this).F1()) {
					return (_1465_v21).Get((_this).F1()).(bool)
				}
				return !(!((_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool)))
			})()) || ((_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool)), !((_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool)) || (true))
			var _1467_v23 _dafny.Map
			_ = _1467_v23
			_1467_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1460_v18).Cardinality(), p1)
			_1466_v22 = Companion_D7_.Create_DC17_(false, _1442_v0, (_1467_v23).Contains((_this).F0()))
			var _1468_v24 _dafny.Set
			_ = _1468_v24
			_1468_v24 = _dafny.SetOf((_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool))
			var _1469_v25 _dafny.Set
			_ = _1469_v25
			_1469_v25 = _dafny.SetOf((p1).F0(), _dafny.IntOfUint32((_dafny.SeqOf((p1).F0(), (_1468_v24).Cardinality())).Cardinality()))
			var _1470_v26 _dafny.MultiSet
			_ = _1470_v26
			_1470_v26 = _dafny.MultiSetOf((_1469_v25).Cardinality(), _dafny.IntOfUint32((p2).Cardinality()), (p1).F0())
			var _1471_v27 D7
			_ = _1471_v27
			_1471_v27 = Companion_D7_.Create_DC16_(_1456_v14)
			var _1472_v28 _dafny.MultiSet
			_ = _1472_v28
			_1472_v28 = _dafny.MultiSetOf((_this).F1())
			var _1473_v29 _dafny.Array
			_ = _1473_v29
			var _nwElement0_41 _dafny.Sequence = _1456_v14
			_ = _nwElement0_41
			var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(23))
			_ = _nw258
			(_nw258).ArraySet1(_nwElement0_41, 0)
			(_nw258).ArraySet1(_1456_v14, 1)
			(_nw258).ArraySet1(_dafny.SeqOf((_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool)), 2)
			(_nw258).ArraySet1(_1456_v14, 3)
			(_nw258).ArraySet1(_1456_v14, 4)
			(_nw258).ArraySet1(_1456_v14, 5)
			(_nw258).ArraySet1(_1456_v14, 6)
			(_nw258).ArraySet1(_dafny.SeqOf(_1442_v0, _1442_v0), 7)
			(_nw258).ArraySet1(_1456_v14, 8)
			(_nw258).ArraySet1(_1456_v14, 9)
			(_nw258).ArraySet1(_1456_v14, 10)
			(_nw258).ArraySet1(_1456_v14, 11)
			(_nw258).ArraySet1(_1456_v14, 12)
			(_nw258).ArraySet1(_1456_v14, 13)
			(_nw258).ArraySet1(_1456_v14, 14)
			(_nw258).ArraySet1(_1456_v14, 15)
			(_nw258).ArraySet1(_1456_v14, 16)
			(_nw258).ArraySet1(_1456_v14, 17)
			(_nw258).ArraySet1(_1456_v14, 18)
			(_nw258).ArraySet1(_1456_v14, 19)
			(_nw258).ArraySet1(_1456_v14, 20)
			(_nw258).ArraySet1(_1456_v14, 21)
			(_nw258).ArraySet1(_1456_v14, 22)
			_1473_v29 = _nw258
			var _1474_v30 _dafny.Sequence
			_ = _1474_v30
			_1474_v30 = _dafny.SeqOf(_1473_v29)
			var _1475_v31 _dafny.MultiSet
			_ = _1475_v31
			_1475_v31 = _dafny.MultiSetOf((_1474_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.IntOfUint32((_1474_v30).Cardinality()))).Uint32()).(_dafny.Array), _1473_v29)
			var _1476_v32 *C0
			_ = _1476_v32
			var _nw259 *C0 = New_C0_()
			_ = _nw259
			_nw259.Ctor__(_dafny.CodePoint('a'))
			_1476_v32 = _nw259
			var _1477_v33 D13
			_ = _1477_v33
			_1477_v33 = Companion_D13_.Create_DC26_(_1476_v32)
			var _1478_v34 _dafny.Map
			_ = _1478_v34
			_1478_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).F0(), _1476_v32)
			var _1479_v35 _dafny.Array
			_ = _1479_v35
			var _nwElement0_42 *C0 = _1476_v32
			_ = _nwElement0_42
			var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(27))
			_ = _nw260
			(_nw260).ArraySet1(_nwElement0_42, 0)
			(_nw260).ArraySet1(_1476_v32, 1)
			(_nw260).ArraySet1(_1476_v32, 2)
			(_nw260).ArraySet1(_1476_v32, 3)
			(_nw260).ArraySet1(_1476_v32, 4)
			(_nw260).ArraySet1(_1476_v32, 5)
			(_nw260).ArraySet1(_1476_v32, 6)
			(_nw260).ArraySet1(_1476_v32, 7)
			(_nw260).ArraySet1(_1476_v32, 8)
			(_nw260).ArraySet1(_1476_v32, 9)
			(_nw260).ArraySet1(_1476_v32, 10)
			(_nw260).ArraySet1(_1476_v32, 11)
			(_nw260).ArraySet1(_1476_v32, 12)
			(_nw260).ArraySet1(_1476_v32, 13)
			(_nw260).ArraySet1(_1476_v32, 14)
			(_nw260).ArraySet1(_1476_v32, 15)
			(_nw260).ArraySet1(_1476_v32, 16)
			(_nw260).ArraySet1((_1477_v33).Dtor_cf42(), 17)
			(_nw260).ArraySet1(_1476_v32, 18)
			(_nw260).ArraySet1(_1476_v32, 19)
			(_nw260).ArraySet1(_1476_v32, 20)
			(_nw260).ArraySet1(_1476_v32, 21)
			(_nw260).ArraySet1((func() *C0 {
				if (_1478_v34).Contains((p1).F0()) {
					return (_1478_v34).Get((p1).F0()).(*C0)
				}
				return _1476_v32
			})(), 22)
			(_nw260).ArraySet1(_1476_v32, 23)
			(_nw260).ArraySet1(_1476_v32, 24)
			(_nw260).ArraySet1(_1476_v32, 25)
			(_nw260).ArraySet1(_1476_v32, 26)
			_1479_v35 = _nw260
			var _1480_v36 D10
			_ = _1480_v36
			_1480_v36 = Companion_D10_.Create_DC22_(_dafny.IntOfInt64(968), (_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool), !(_1442_v0))
			var _1481_v37 _dafny.Map
			_ = _1481_v37
			_1481_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1442_v0, (p1).F0())
			var _1482_v38 _dafny.Set
			_ = _1482_v38
			_1482_v38 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1442_v0, _1442_v0))
			var _1483_v39 D5
			_ = _1483_v39
			_1483_v39 = Companion_D5_.Create_DC13_((_this).F2(), _1468_v24, (_this).F1(), (_this).F2(), _1442_v0)
			var _pat_let_tv35 = p1
			_ = _pat_let_tv35
			var _pat_let_tv36 = _1476_v32
			_ = _pat_let_tv36
			var _1484_v40 _dafny.Array
			_ = _1484_v40
			var _nwElement0_43 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-289), (_this).F0())
			_ = _nwElement0_43
			var _nw261 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(28))
			_ = _nw261
			(_nw261).ArraySet1(_nwElement0_43, 0)
			(_nw261).ArraySet1(Companion_Default___.SafeDivisionInt((_1460_v18).Cardinality(), (p1).F0()), 1)
			(_nw261).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfInt64(73)).Plus((_this).F0())), 2)
			(_nw261).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, (p1).F0())), 3)
			(_nw261).ArraySet1((p1).F0(), 4)
			(_nw261).ArraySet1((p1).F0(), 5)
			(_nw261).ArraySet1((func() _dafny.Int {
				if _1442_v0 {
					return (_this).F0()
				}
				return (p1).F0()
			})(), 6)
			(_nw261).ArraySet1((_this).F0(), 7)
			(_nw261).ArraySet1(Companion_Default___.SafeModuloInt((p1).F0(), ((_1470_v26).Update((p1).F0(), Companion_Default___.Abs(_dafny.IntOfUint32(((_1471_v27).Dtor_cf26()).Cardinality())))).Cardinality()), 8)
			(_nw261).ArraySet1(_dafny.IntOfInt64(-526), 9)
			(_nw261).ArraySet1((p1).F0(), 10)
			(_nw261).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm3(Companion_Default___.Fm10(_1472_v28, (_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool), globalState), (_this).F1(), (p1).F0(), (_dafny.Zero).Minus((p1).F0()), globalState)), 11)
			(_nw261).ArraySet1((_dafny.Zero).Minus((p1).F0()), 12)
			(_nw261).ArraySet1((p1).F0(), 13)
			(_nw261).ArraySet1(Companion_Default___.SafeModuloInt((_this).F2(), (func() _dafny.Int {
				if (_1475_v31).Contains(_1473_v29) {
					return (_1475_v31).Multiplicity(_1473_v29)
				}
				return (Companion_D3_.Create_DC9_((p1).F0(), _1479_v35, (_1462_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_1462_v20), 0))).Int()).(bool))).Dtor_cf10()
			})()), 14)
			(_nw261).ArraySet1((_1480_v36).Dtor_cf34(), 15)
			(_nw261).ArraySet1((p1).F0(), 16)
			(_nw261).ArraySet1((_this).F0(), 17)
			(_nw261).ArraySet1((_this).F2(), 18)
			(_nw261).ArraySet1(_dafny.IntOfUint32((_1458_v16).Cardinality()), 19)
			(_nw261).ArraySet1(_dafny.IntOfUint32((_1458_v16).Cardinality()), 20)
			(_nw261).ArraySet1((_this).F0(), 21)
			(_nw261).ArraySet1((_this).F0(), 22)
			(_nw261).ArraySet1(((_dafny.Zero).Minus((p1).F0())).Times((_this).F2()), 23)
			(_nw261).ArraySet1((p1).F0(), 24)
			(_nw261).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_1481_v37).Cardinality()), (p1).F0()), 25)
			(_nw261).ArraySet1((_1482_v38).Cardinality(), 26)
			(_nw261).ArraySet1(((func(_pat_let27_0 D5) D5 {
				return func(_1485_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let28_0 _dafny.Int) D5 {
						return func(_1486_dt__update_hcf17_h0 _dafny.Int) D5 {
							return func(_pat_let29_0 _dafny.CodePoint) D5 {
								return func(_1487_dt__update_hcf19_h0 _dafny.CodePoint) D5 {
									return Companion_D5_.Create_DC13_(_1486_dt__update_hcf17_h0, (_1485_dt__update__tmp_h0).Dtor_cf18(), _1487_dt__update_hcf19_h0, (_1485_dt__update__tmp_h0).Dtor_cf20(), (_1485_dt__update__tmp_h0).Dtor_cf21())
								}(_pat_let29_0)
							}((_pat_let_tv36).F3())
						}(_pat_let28_0)
					}((_pat_let_tv35).F0())
				}(_pat_let27_0)
			}(_1483_v39)).Dtor_cf18()).Cardinality(), 27)
			_1484_v40 = _nw261
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1484_v40), 0))
			_ = _index247
			(_1484_v40).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), (_index247).Int())
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1484_v40), 0))
			_ = _index248
			(_1484_v40).ArraySet1(((p1).F0()).Times(_dafny.IntOfInt64(814)), (_index248).Int())
		}
		var _1488_v41 _dafny.Array
		_ = _1488_v41
		var _len0_33 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_33
		var _nw262 _dafny.Array
		_ = _nw262
		if _len0_33.Cmp(_dafny.Zero) == 0 {
			_nw262 = _dafny.NewArray(_len0_33)
		} else {
			var _init33 func(_dafny.Int) _dafny.Int = func(_1489_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_1489_i1, _dafny.IntOfInt64(-498))
			}
			_ = _init33
			var _element0_33 = _init33(_dafny.Zero)
			_ = _element0_33
			_nw262 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
			(_nw262).ArraySet1(_element0_33, 0)
			var _nativeLen0_33 = (_len0_33).Int()
			_ = _nativeLen0_33
			for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
				(_nw262).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
			}
		}
		_1488_v41 = _nw262
		var _1490_v42 _dafny.Sequence
		_ = _1490_v42
		_1490_v42 = _dafny.SeqOf(_1488_v41, _1488_v41)
		var _1491_v43 _dafny.MultiSet
		_ = _1491_v43
		_1491_v43 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_this).F0())), _dafny.SeqOf((p1).F0(), _dafny.IntOfInt64(531), _dafny.IntOfUint32((_1490_v42).Cardinality()), p0, p0)))
		var _1492_v44 _dafny.Array
		_ = _1492_v44
		var _nw263 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
		_ = _nw263
		_1492_v44 = _nw263
		var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1492_v44), 0))
		_ = _index249
		(_1492_v44).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_this).Fm5(globalState) {
				return _dafny.CodePoint('k')
			}
			return (_this).F1()
		})(), (_index249).Int())
		var _1493_v45 _dafny.Int
		_ = _1493_v45
		_1493_v45 = _dafny.IntOfInt64(113)
		var _1494_v46 _dafny.Set
		_ = _1494_v46
		_1494_v46 = _dafny.SetOf(true)
		var _1495_v47 _dafny.Sequence
		_ = _1495_v47
		_1495_v47 = _dafny.UnicodeSeqOfUtf8Bytes("davdt")
		var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1492_v44), 0))
		_ = _index250
		var _rhs239 _dafny.MultiSet = ((_1491_v43).Union(_1491_v43)).Intersection((_dafny.MultiSetOf(p2, p2)).Update(p2, Companion_Default___.Abs(Companion_Default___.Fm3(_1494_v46, (_this).F1(), (_this).F2(), (p1).F0(), globalState))))
		_ = _rhs239
		var _rhs240 bool = !(_1442_v0)
		_ = _rhs240
		var _rhs241 _dafny.CodePoint = (_this).F1()
		_ = _rhs241
		var _rhs242 _dafny.Int = Companion_Default___.SafeModuloInt((p1).F0(), (_1493_v45).Minus(_dafny.IntOfUint32((_1495_v47).Cardinality())))
		_ = _rhs242
		var _rhs243 _dafny.Int = _1493_v45
		_ = _rhs243
		var _lhs108 _dafny.Array = _1492_v44
		_ = _lhs108
		var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1492_v44), 0))
		_ = _lhs109
		_1491_v43 = _rhs239
		r0 = _rhs240
		(_lhs108).ArraySet1CodePoint(_rhs241, (_lhs109).Int())
		_1493_v45 = _rhs242
		_1493_v45 = _rhs243
		var _1496_v48 *C0
		_ = _1496_v48
		var _nw264 *C0 = New_C0_()
		_ = _nw264
		_nw264.Ctor__(_dafny.CodePoint('h'))
		_1496_v48 = _nw264
		var _1497_v49 D13
		_ = _1497_v49
		_1497_v49 = Companion_D13_.Create_DC26_(_1496_v48)
		var _source19 D13 = _1497_v49
		_ = _source19
		if _source19.Is_DC27() {
			var _1498___mcc_h0 _dafny.CodePoint = _source19.Get_().(D13_DC27).Cf43
			_ = _1498___mcc_h0
			var _1499___mcc_h1 T0 = _source19.Get_().(D13_DC27).Cf44
			_ = _1499___mcc_h1
			var _1500___mcc_h2 _dafny.CodePoint = _source19.Get_().(D13_DC27).Cf45
			_ = _1500___mcc_h2
			var _1501___mcc_h3 _dafny.Int = _source19.Get_().(D13_DC27).Cf46
			_ = _1501___mcc_h3
			var _1502_cf46 _dafny.Int = _1501___mcc_h3
			_ = _1502_cf46
			var _1503_cf45 _dafny.CodePoint = _1500___mcc_h2
			_ = _1503_cf45
			var _1504_cf44 T0 = _1499___mcc_h1
			_ = _1504_cf44
			var _1505_cf43 _dafny.CodePoint = _1498___mcc_h0
			_ = _1505_cf43
			var _1506_v50 _dafny.Array
			_ = _1506_v50
			var _nw265 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw265
			_1506_v50 = _nw265
			var _1507_v51 _dafny.Array
			_ = _1507_v51
			var _nw266 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw266
			_1507_v51 = _nw266
			var _1508_v52 _dafny.Sequence
			_ = _1508_v52
			_1508_v52 = _dafny.SeqOf(_1507_v51)
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_1506_v50), 0))
			_ = _index251
			(_1506_v50).ArraySet1(_1508_v52, (_index251).Int())
			var _1509_v53 _dafny.MultiSet
			_ = _1509_v53
			_1509_v53 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1507_v51)).Cardinality())))
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_1506_v50), 0))
			_ = _index252
			var _rhs244 _dafny.Sequence = _1508_v52
			_ = _rhs244
			var _rhs245 _dafny.MultiSet = Companion_Default___.Fm22(globalState)
			_ = _rhs245
			var _lhs110 _dafny.Array = _1506_v50
			_ = _lhs110
			var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_1506_v50), 0))
			_ = _lhs111
			(_lhs110).ArraySet1(_rhs244, (_lhs111).Int())
			_1509_v53 = _rhs245
			_1502_cf46 = ((p2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int)).Minus((func() _dafny.Int {
				if _1442_v0 {
					return (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1495_v47).Cardinality())))
				}
				return (_1504_cf44).F0()
			})())
			var _1510_v54 _dafny.Array
			_ = _1510_v54
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_34
			var _nw267 _dafny.Array
			_ = _nw267
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw267 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1511_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1512_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1511_v0), _dafny.SeqOf(_1511_v0))
					}
				})(_1442_v0)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw267 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw267).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw267).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1510_v54 = _nw267
			var _1513_v55 _dafny.Sequence
			_ = _1513_v55
			_1513_v55 = _dafny.SeqOf(_1510_v54, _1510_v54, _1510_v54)
			_1510_v54 = (_1513_v55).Select((Companion_Default___.SafeIndex((p1).F0(), _dafny.IntOfUint32((_1513_v55).Cardinality()))).Uint32()).(_dafny.Array)
			var _1514_v56 D0
			_ = _1514_v56
			_1514_v56 = Companion_D0_.Create_DC1_(_1442_v0)
			_1514_v56 = _1514_v56
		} else if _source19.Is_DC28() {
			var _1515___mcc_h4 bool = _source19.Get_().(D13_DC28).Cf47
			_ = _1515___mcc_h4
			var _1516___mcc_h5 _dafny.CodePoint = _source19.Get_().(D13_DC28).Cf48
			_ = _1516___mcc_h5
			var _1517_cf48 _dafny.CodePoint = _1516___mcc_h5
			_ = _1517_cf48
			var _1518_cf47 bool = _1515___mcc_h4
			_ = _1518_cf47
			var _1519_v57 _dafny.Array
			_ = _1519_v57
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_35
			var _nw268 _dafny.Array
			_ = _nw268
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw268 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) bool = func(_1520_i3 _dafny.Int) bool {
					return true
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw268 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw268).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw268).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1519_v57 = _nw268
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1519_v57), 0))
			_ = _index253
			(_1519_v57).ArraySet1(_1442_v0, (_index253).Int())
			var _1521_v58 _dafny.Sequence
			_ = _1521_v58
			_1521_v58 = _dafny.SeqOf(_1442_v0, _1442_v0)
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1519_v57), 0))
			_ = _index254
			(_1519_v57).ArraySet1((_1518_cf47) && (_dafny.Companion_Sequence_.IsPrefixOf(_1521_v58, _1521_v58)), (_index254).Int())
			_1493_v45 = (_this).F0()
			var _1522_v59 _dafny.MultiSet
			_ = _1522_v59
			_1522_v59 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(-389))
			var _1523_v60 _dafny.MultiSet
			_ = _1523_v60
			_1523_v60 = _1522_v59
			var _1524_v61 _dafny.Map
			_ = _1524_v61
			_1524_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1523_v60), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F2(), (_dafny.Zero).Minus(p0)))).IsProperSubsetOf(_1522_v59))
			var _rhs246 _dafny.Int = (p1).F0()
			_ = _rhs246
			var _rhs247 _dafny.Map = Companion_Default___.Fm29(_1442_v0, _1518_cf47, _1518_cf47, _1518_cf47, globalState)
			_ = _rhs247
			_1493_v45 = _rhs246
			_1524_v61 = _rhs247
			var _1525_v62 _dafny.Map
			_ = _1525_v62
			_1525_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1521_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gisbfojc")).Cardinality()), _dafny.IntOfUint32((_1521_v58).Cardinality()))).Uint32()).(bool), _1518_cf47)
			var _1526_v63 _dafny.Map
			_ = _1526_v63
			_1526_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1525_v62).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1442_v0)), (p0).Cmp((p1).F0()) > 0)
			_1526_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1525_v62, ((_this).F2()).Cmp((p1).F0()) <= 0)
		} else if _source19.Is_DC29() {
			var _1527___mcc_h6 _dafny.CodePoint = _source19.Get_().(D13_DC29).Cf49
			_ = _1527___mcc_h6
			var _1528___mcc_h7 bool = _source19.Get_().(D13_DC29).Cf50
			_ = _1528___mcc_h7
			var _1529_cf50 bool = _1528___mcc_h7
			_ = _1529_cf50
			var _1530_cf49 _dafny.CodePoint = _1527___mcc_h6
			_ = _1530_cf49
			if _1442_v0 {
				var _1531_v64 _dafny.Sequence
				_ = _1531_v64
				_1531_v64 = _dafny.SeqOf(_1442_v0)
				_1531_v64 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1532_v48 *C0) func(_dafny.Int) _dafny.CodePoint {
					return func(_1533_i4 _dafny.Int) _dafny.CodePoint {
						return (_1532_v48).F3()
					}
				})(_1496_v48))), _1495_v47), _1442_v0, _1529_cf50)
				var _1534_v65 _dafny.MultiSet
				_ = _1534_v65
				_1534_v65 = _dafny.MultiSetOf(_dafny.IntOfInt64(-707), (_this).F0())
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index255
				(_1488_v41).ArraySet1(((_dafny.MultiSetFromSeq(p2)).Difference(_1534_v65)).Cardinality(), (_index255).Int())
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index256
				(_1488_v41).ArraySet1(Companion_Default___.SafeModuloInt((_1493_v45).Plus((p1).F0()), (p1).F0()), (_index256).Int())
				r0 = _1529_cf50
				var _1535_v66 *C2
				_ = _1535_v66
				var _nw269 *C2 = New_C2_()
				_ = _nw269
				_nw269.Ctor__((_1488_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1488_v41), 0))).Int()).(_dafny.Int))
				_1535_v66 = _nw269
				var _1536_v67 _dafny.MultiSet
				_ = _1536_v67
				_1536_v67 = _dafny.MultiSetOf(_1535_v66)
				_1529_cf50 = ((_1536_v67).IsSubsetOf(_1536_v67)) && (!(!_dafny.Companion_Sequence_.Contains(_1495_v47, _dafny.CodePoint('w'))))
				_1529_cf50 = _1442_v0
			} else {
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index257
				(_1488_v41).ArraySet1((_this).F2(), (_index257).Int())
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index258
				(_1488_v41).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(_dafny.SetOf(_1442_v0), (_this).F1(), (p1).F0(), (p1).F0(), globalState), p0)), (_index258).Int())
				var _1537_v68 _dafny.Sequence
				_ = _1537_v68
				_1537_v68 = _dafny.SeqOf(false, _1529_cf50, _1442_v0)
				var _1538_v69 D7
				_ = _1538_v69
				_1538_v69 = Companion_D7_.Create_DC16_(_dafny.SeqOf(_1442_v0, _1529_cf50, true, true))
				_1537_v68 = (_1538_v69).Dtor_cf26()
				var _1539_v70 _dafny.MultiSet
				_ = _1539_v70
				_1539_v70 = _dafny.MultiSetOf(_1538_v69)
				var _1540_v71 _dafny.Map
				_ = _1540_v71
				_1540_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).F0(), _1529_cf50)
				var _1541_v72 _dafny.MultiSet
				_ = _1541_v72
				_1541_v72 = _dafny.MultiSetOf(_1493_v45, (p1).F0())
				var _1542_v73 _dafny.Set
				_ = _1542_v73
				_1542_v73 = _dafny.SetOf(p0, (_1541_v72).Cardinality())
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index259
				(_1488_v41).ArraySet1(Companion_Default___.Fm3(_1494_v46, (_1496_v48).F3(), Companion_Default___.Fm3(Companion_Default___.Fm27(_1442_v0, _1539_v70, (_1540_v71).Cardinality(), _1542_v73, globalState), (_1496_v48).F3(), (p1).F0(), (_this).F0(), globalState), Companion_Default___.SafeModuloInt((p1).F0(), (p1).F0()), globalState), (_index259).Int())
				_1493_v45 = (_dafny.SetOf((_1488_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1488_v41), 0))).Int()).(_dafny.Int))).Cardinality()
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index260
				(_1488_v41).ArraySet1(Companion_Default___.Fm3((_1494_v46).Intersection(_1494_v46), _dafny.CodePoint('f'), (_1542_v73).Cardinality(), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("to")).Cardinality())).Plus((_this).F2()), globalState), (_index260).Int())
			}
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1488_v41), 0))
			_ = _index261
			(_1488_v41).ArraySet1(_dafny.IntOfInt64(86), (_index261).Int())
			var _1543_v74 _dafny.Map
			_ = _1543_v74
			_1543_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_cf49, (_1492_v44).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1492_v44), 0))).Int()))
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1488_v41), 0))
			_ = _index262
			var _rhs248 _dafny.Int = ((p0).Plus(_dafny.IntOfUint32((_1495_v47).Cardinality()))).Times((p2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(_1493_v45, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F0())).Cardinality()), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs248
			var _rhs249 _dafny.Int = (_1543_v74).Cardinality()
			_ = _rhs249
			var _rhs250 bool = !(_1442_v0)
			_ = _rhs250
			var _lhs112 _dafny.Array = _1488_v41
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1488_v41), 0))
			_ = _lhs113
			_1493_v45 = _rhs248
			(_lhs112).ArraySet1(_rhs249, (_lhs113).Int())
			r0 = _rhs250
			var _1544_v75 _dafny.Set
			_ = _1544_v75
			_1544_v75 = _dafny.SetOf(p0, _1493_v45)
			_1530_cf49 = (_this).Fm7((_dafny.IntOfUint32((Companion_Default___.Fm18((_dafny.MultiSetOf(_1442_v0)).Cardinality(), globalState)).Cardinality())).Plus((_dafny.SetOf(!(!(_1529_cf50)), !(_1529_cf50))).Cardinality()), (_1488_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1488_v41), 0))).Int()).(_dafny.Int), _1529_cf50, (_dafny.Zero).Minus(Companion_Default___.Fm3(_1494_v46, (_this).F1(), (_1544_v75).Cardinality(), (_dafny.Zero).Minus(p0), globalState)), globalState)
			var _1545_v76 _dafny.Map
			_ = _1545_v76
			_1545_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), _1495_v47)
			var _1546_v77 D2
			_ = _1546_v77
			_1546_v77 = Companion_D2_.Create_DC4_(_dafny.Companion_Sequence_.Concatenate(_1495_v47, (func() _dafny.Sequence {
				if (_1545_v76).Contains(p0) {
					return (_1545_v76).Get(p0).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_1495_v47, (Companion_Default___.SafeIndex((_1488_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1488_v41), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1495_v47).Cardinality()))).Uint32(), (_1496_v48).F3())
			})()), ((_1488_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_1488_v41), 0))).Int()).(_dafny.Int)).Cmp((_this).F0()) > 0)
			r1 = _1546_v77
		} else if _source19.Is_DC26() {
			var _1547___mcc_h8 *C0 = _source19.Get_().(D13_DC26).Cf42
			_ = _1547___mcc_h8
			var _1548_cf42 *C0 = _1547___mcc_h8
			_ = _1548_cf42
			var _1549_v78 _dafny.MultiSet
			_ = _1549_v78
			_1549_v78 = _dafny.MultiSetOf((p1).F0())
			r0 = !((_1549_v78).IsProperSubsetOf((func() _dafny.MultiSet {
				if _1442_v0 {
					return _1549_v78
				}
				return _1549_v78
			})()))
			var _1550_v79 _dafny.Array
			_ = _1550_v79
			var _nwElement0_44 bool = _1442_v0
			_ = _nwElement0_44
			var _nw270 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(6))
			_ = _nw270
			(_nw270).ArraySet1(_nwElement0_44, 0)
			(_nw270).ArraySet1(_1442_v0, 1)
			(_nw270).ArraySet1(_1442_v0, 2)
			(_nw270).ArraySet1(_1442_v0, 3)
			(_nw270).ArraySet1(_1442_v0, 4)
			(_nw270).ArraySet1(_1442_v0, 5)
			_1550_v79 = _nw270
			var _1551_v80 _dafny.Sequence
			_ = _1551_v80
			_1551_v80 = _dafny.SeqOf(_1550_v79)
			_1493_v45 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1551_v80).Cardinality()), (p1).F0())
			var _1552_v81 _dafny.Array
			_ = _1552_v81
			var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw271
			_1552_v81 = _nw271
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_1552_v81), 0))
			_ = _index263
			(_1552_v81).ArraySet1(_1488_v41, (_index263).Int())
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_1552_v81), 0))
			_ = _index264
			var _rhs251 _dafny.Array = _1488_v41
			_ = _rhs251
			var _rhs252 _dafny.Array = _1550_v79
			_ = _rhs252
			var _rhs253 _dafny.Int = (p1).F0()
			_ = _rhs253
			var _rhs254 bool = (_1442_v0) || (!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm19(_1549_v78, (p1).F0(), (_this).F2(), (p1).F0(), globalState), _1495_v47))
			_ = _rhs254
			var _rhs255 _dafny.Int = (func() _dafny.Int {
				if (_dafny.MultiSetFromSeq(p2)).IsSubsetOf(_1549_v78) {
					return p0
				}
				return ((_this).F0()).Minus(_dafny.IntOfInt64(734))
			})()
			_ = _rhs255
			var _lhs114 _dafny.Array = _1552_v81
			_ = _lhs114
			var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_1552_v81), 0))
			_ = _lhs115
			(_lhs114).ArraySet1(_rhs251, (_lhs115).Int())
			_1550_v79 = _rhs252
			_1493_v45 = _rhs253
			r0 = _rhs254
			_1493_v45 = _rhs255
			_1495_v47 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1495_v47, _1495_v47), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-630))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}((func(_1553_cf42 *C0) func(_dafny.Int) _dafny.CodePoint {
				return func(_1554_i5 _dafny.Int) _dafny.CodePoint {
					return (_1553_cf42).F3()
				}
			})(_1548_cf42))))
		} else {
			var _1555___mcc_h9 D13 = _source19.Get_().(D13_DC30).Cf51
			_ = _1555___mcc_h9
			var _1556_cf51 D13 = _1555___mcc_h9
			_ = _1556_cf51
			var _1557_v82 D0
			_ = _1557_v82
			_1557_v82 = Companion_D0_.Create_DC0_(_1442_v0)
			var _1558_v83 _dafny.Sequence
			_ = _1558_v83
			_1558_v83 = _dafny.SeqOf(_1557_v82)
			var _1559_v84 _dafny.Set
			_ = _1559_v84
			_1559_v84 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1558_v83).Cardinality())))
			var _1560_v85 _dafny.Sequence
			_ = _1560_v85
			_1560_v85 = _dafny.SeqOf(_1559_v84)
			var _1561_v86 _dafny.MultiSet
			_ = _1561_v86
			_1561_v86 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1560_v85).Cardinality()))
			var _1562_v87 _dafny.Map
			_ = _1562_v87
			_1562_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fyltafcwf")).Cardinality()))
			var _1563_v88 _dafny.Sequence
			_ = _1563_v88
			_1563_v88 = _dafny.SeqOf(_1561_v86, _dafny.MultiSetFromSeq(_dafny.SeqOf(_1493_v45, p0, _dafny.IntOfInt64(222), (p1).F0())))
			var _1564_v90 _dafny.Map
			_ = _1564_v90
			_1564_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(798), _dafny.IntOfInt64(-379))); ; {
					_compr_33, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _1565_v89 _dafny.Int
					_1565_v89 = interface{}(_compr_33).(_dafny.Int)
					if ((_dafny.IntOfInt64(798)).Cmp(_1565_v89) <= 0) && ((_1565_v89).Cmp(_dafny.IntOfInt64(-379)) < 0) {
						_coll33.Add((_1565_v89).Times((_this).F0()), _dafny.IntOfUint32((_1495_v47).Cardinality()))
					}
				}
				return _coll33.ToMap()
			}(), _dafny.MultiSetOf(_dafny.IntOfUint32((_1495_v47).Cardinality()), Companion_Default___.Fm3(_1494_v46, (_1496_v48).F3(), (_1559_v84).Cardinality(), (_this).F0(), globalState)))
			var _1566_v91 _dafny.Array
			_ = _1566_v91
			var _nwElement0_45 _dafny.MultiSet = _1561_v86
			_ = _nwElement0_45
			var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(24))
			_ = _nw272
			(_nw272).ArraySet1(_nwElement0_45, 0)
			(_nw272).ArraySet1(_1561_v86, 1)
			(_nw272).ArraySet1(_1561_v86, 2)
			(_nw272).ArraySet1(_1561_v86, 3)
			(_nw272).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfInt64(-230))).Update((_1562_v87).Cardinality(), Companion_Default___.Abs((_this).F2())), 4)
			(_nw272).ArraySet1(Companion_Default___.Fm22(globalState), 5)
			(_nw272).ArraySet1(_1561_v86, 6)
			(_nw272).ArraySet1(_dafny.MultiSetOf((_dafny.Zero).Minus((_this).F2())), 7)
			(_nw272).ArraySet1(_1561_v86, 8)
			(_nw272).ArraySet1((_1561_v86).Update((_this).F0(), Companion_Default___.Abs(p0)), 9)
			(_nw272).ArraySet1((_dafny.MultiSetOf((p1).F0())).Intersection(_1561_v86), 10)
			(_nw272).ArraySet1((_1563_v88).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.IntOfUint32((_1563_v88).Cardinality()))).Uint32()).(_dafny.MultiSet), 11)
			(_nw272).ArraySet1(_1561_v86, 12)
			(_nw272).ArraySet1((_1563_v88).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1563_v88).Cardinality()))).Uint32()).(_dafny.MultiSet), 13)
			(_nw272).ArraySet1((_1561_v86).Intersection(_1561_v86), 14)
			(_nw272).ArraySet1(_1561_v86, 15)
			(_nw272).ArraySet1((func() _dafny.MultiSet {
				if (_1564_v90).Contains(_1562_v87) {
					return (_1564_v90).Get(_1562_v87).(_dafny.MultiSet)
				}
				return _1561_v86
			})(), 16)
			(_nw272).ArraySet1(_1561_v86, 17)
			(_nw272).ArraySet1(_1561_v86, 18)
			(_nw272).ArraySet1(_1561_v86, 19)
			(_nw272).ArraySet1((_1561_v86).Difference(Companion_Default___.Fm22(globalState)), 20)
			(_nw272).ArraySet1(_1561_v86, 21)
			(_nw272).ArraySet1(_dafny.MultiSetOf((_this).F2(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cm")).Cardinality())), 22)
			(_nw272).ArraySet1((_1561_v86).Intersection(_1561_v86), 23)
			_1566_v91 = _nw272
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1566_v91), 0))
			_ = _index265
			(_1566_v91).ArraySet1(_dafny.MultiSetFromSeq(p2), (_index265).Int())
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1566_v91), 0))
			_ = _index266
			(_1566_v91).ArraySet1(_1561_v86, (_index266).Int())
			var _1567_v92 _dafny.Array
			_ = _1567_v92
			var _nw273 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw273
			_1567_v92 = _nw273
			var _1568_v93 *C3
			_ = _1568_v93
			var _nw274 *C3 = New_C3_()
			_ = _nw274
			_nw274.Ctor__(_1493_v45)
			_1568_v93 = _nw274
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_1567_v92), 0))
			_ = _index267
			(_1567_v92).ArraySet1(_1568_v93, (_index267).Int())
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_1567_v92), 0))
			_ = _index268
			(_1567_v92).ArraySet1(_1568_v93, (_index268).Int())
			var _1569_v94 _dafny.Map
			_ = _1569_v94
			_1569_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(p2, p2), _1442_v0)
			_1569_v94 = (_1569_v94).Merge(_1569_v94)
			var _1570_v95 _dafny.Array
			_ = _1570_v95
			var _nw275 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw275
			_1570_v95 = _nw275
			var _1571_v96 _dafny.Int
			_ = _1571_v96
			var _1572_v97 _dafny.Int
			_ = _1572_v97
			var _1573_v98 _dafny.Int
			_ = _1573_v98
			var _out24 _dafny.Int
			_ = _out24
			var _out25 _dafny.Int
			_ = _out25
			var _out26 _dafny.Int
			_ = _out26
			_out24, _out25, _out26 = (p1).M1(_1570_v95, (_1493_v45).Plus((_dafny.Zero).Minus(_dafny.IntOfInt64(-94))), globalState)
			_1571_v96 = _out24
			_1572_v97 = _out25
			_1573_v98 = _out26
		}
		_1493_v45 = (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm0(globalState)).Cardinality()))
		var _hi12 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F2(), (_dafny.Zero).Minus((_this).F0()))
		_ = _hi12
		for _1574_i6 := (p1).F0(); _1574_i6.Cmp(_hi12) < 0; _1574_i6 = _1574_i6.Plus(_dafny.One) {
			if !(_1442_v0) {
				var _1575_v99 _dafny.Array
				_ = _1575_v99
				var _len0_36 _dafny.Int = _dafny.One
				_ = _len0_36
				var _nw276 _dafny.Array
				_ = _nw276
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw276 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Map = (func(_1576_v0 bool) func(_dafny.Int) _dafny.Map {
						return func(_1577_i7 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1576_v0, _dafny.SeqOf(_1576_v0, _1576_v0, _1576_v0))
						}
					})(_1442_v0)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw276 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw276).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw276).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1575_v99 = _nw276
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1575_v99), 0))
				_ = _index269
				(_1575_v99).ArraySet1(Companion_Default___.Fm30(_1442_v0, _1442_v0, _1442_v0, globalState), (_index269).Int())
				var _1578_v100 D13
				_ = _1578_v100
				_1578_v100 = Companion_D13_.Create_DC28_(_1442_v0, (_1496_v48).F3())
				var _1579_v101 _dafny.Map
				_ = _1579_v101
				_1579_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1442_v0, _dafny.SeqOf((_1578_v100).Dtor_cf47(), _1442_v0, _1442_v0, _1442_v0))
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1575_v99), 0))
				_ = _index270
				(_1575_v99).ArraySet1(_1579_v101, (_index270).Int())
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index271
				(_1488_v41).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1495_v47).Cardinality()), (p1).F0()), (_index271).Int())
				var _1580_v102 _dafny.Sequence
				_ = _1580_v102
				_1580_v102 = _dafny.SeqOf(_1442_v0)
				var _1581_v103 _dafny.MultiSet
				_ = _1581_v103
				_1581_v103 = _dafny.MultiSetOf((_1580_v102).Select((Companion_Default___.SafeIndex((p1).F0(), _dafny.IntOfUint32((_1580_v102).Cardinality()))).Uint32()).(bool), !((_1580_v102).Select((Companion_Default___.SafeIndex((p1).F0(), _dafny.IntOfUint32((_1580_v102).Cardinality()))).Uint32()).(bool)), _1442_v0, _1442_v0)
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index272
				var _rhs256 bool = (_1581_v103).Contains(_1442_v0)
				_ = _rhs256
				var _rhs257 _dafny.Int = (func() _dafny.Int {
					if _1442_v0 {
						return (p1).F0()
					}
					return (p1).F0()
				})()
				_ = _rhs257
				var _lhs116 _dafny.Array = _1488_v41
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_1488_v41), 0))
				_ = _lhs117
				r0 = _rhs256
				(_lhs116).ArraySet1(_rhs257, (_lhs117).Int())
				var _1582_v104 _dafny.Set
				_ = _1582_v104
				_1582_v104 = _dafny.SetOf(_1495_v47)
				var _1583_v105 _dafny.Sequence
				_ = _1583_v105
				_1583_v105 = _dafny.SeqOf(p1)
				var _1584_v106 *C3
				_ = _1584_v106
				var _nw277 *C3 = New_C3_()
				_ = _nw277
				_nw277.Ctor__((p1).F0())
				_1584_v106 = _nw277
				var _1585_v107 _dafny.MultiSet
				_ = _1585_v107
				_1585_v107 = _dafny.MultiSetOf(_1584_v106)
				var _1586_v108 _dafny.MultiSet
				_ = _1586_v108
				_1586_v108 = _dafny.MultiSetOf((_this).F2(), (p1).F0(), _1493_v45, _dafny.IntOfUint32((_1583_v105).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((p1).F0(), p0, (_1585_v107).Cardinality())).Cardinality()))
				var _1587_v109 D7
				_ = _1587_v109
				_1587_v109 = Companion_D7_.Create_DC18_((_1488_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_1488_v41), 0))).Int()).(_dafny.Int))
				_1442_v0 = (_this).Fm6((_1582_v104).IsSubsetOf(_dafny.SetOf(_1495_v47, _1495_v47)), Companion_Default___.Fm19(_1586_v108, (_this).F2(), (p1).F0(), (_1587_v109).Dtor_cf30(), globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vr"), _1495_v47), _dafny.IntOfInt64(-812), globalState)
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_1488_v41), 0))
				_ = _index273
				(_1488_v41).ArraySet1((p1).F0(), (_index273).Int())
				_1442_v0 = _1442_v0
			} else {
				var _1588_v110 _dafny.Map
				_ = _1588_v110
				_1588_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1495_v47, _1442_v0)
				r0 = !(_1588_v110).Equals(_1588_v110)
				var _1589_v111 _dafny.Map
				_ = _1589_v111
				_1589_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).F0(), Companion_Default___.Fm2(_1442_v0, (_this).F2(), !(_1442_v0), _1442_v0, globalState))
				_1589_v111 = (_1589_v111).Update(_1493_v45, false)
				_1493_v45 = (_1493_v45).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1495_v47, _1495_v47)).Cardinality()))
				var _1590_v112 _dafny.Array
				_ = _1590_v112
				var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw278
				_1590_v112 = _nw278
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1590_v112), 0))
				_ = _index274
				(_1590_v112).ArraySet1(_dafny.SeqOf((_1492_v44).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1492_v44), 0))).Int()), (_1496_v48).F3()), (_index274).Int())
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1590_v112), 0))
				_ = _index275
				(_1590_v112).ArraySet1(_1495_v47, (_index275).Int())
				var _1591_v113 _dafny.MultiSet
				_ = _1591_v113
				_1591_v113 = _dafny.MultiSetOf(_1574_i6)
				var _1592_v114 _dafny.Sequence
				_ = _1592_v114
				_1592_v114 = _dafny.SeqOf(Companion_Default___.Fm19(_1591_v113, (p1).F0(), (p1).F0(), _dafny.IntOfInt64(180), globalState))
				var _1593_v115 _dafny.Sequence
				_ = _1593_v115
				_1593_v115 = _1592_v114
				_1592_v114 = (_1593_v115)
			}
			var _1594_v116 _dafny.Sequence
			_ = _1594_v116
			_1594_v116 = _dafny.SeqOf(_1442_v0, (_dafny.IntOfInt64(-270)).Cmp((_this).F0()) >= 0)
			_1594_v116 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1594_v116, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.IntOfUint32((_1594_v116).Cardinality()))).Uint32(), _1442_v0), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(true, _1442_v0, !(_1442_v0)), (Companion_Default___.SafeIndex((p1).F0(), _dafny.IntOfUint32((_dafny.SeqOf(true, _1442_v0, !(_1442_v0))).Cardinality()))).Uint32(), _1442_v0))
			var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1492_v44), 0))
			_ = _index276
			(_1492_v44).ArraySet1CodePoint((_1496_v48).F3(), (_index276).Int())
			var _1595_v117 _dafny.Array
			_ = _1595_v117
			var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(24))
			_ = _nw279
			_1595_v117 = _nw279
			var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_1595_v117), 0))
			_ = _index277
			(_1595_v117).ArraySet1(Companion_Default___.Fm31(globalState), (_index277).Int())
			var _1596_v118 _dafny.Set
			_ = _1596_v118
			_1596_v118 = _dafny.SetOf(_dafny.CodePoint('g'))
			var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_1595_v117), 0))
			_ = _index278
			(_1595_v117).ArraySet1((_1596_v118).Union(_1596_v118), (_index278).Int())
		}
		var _1597_v119 _dafny.Sequence
		_ = _1597_v119
		_1597_v119 = _dafny.SeqOf(_1495_v47, _1495_v47, _1495_v47)
		_1495_v47 = _dafny.Companion_Sequence_.Concatenate(_1495_v47, _dafny.Companion_Sequence_.Concatenate((_1597_v119).Select((Companion_Default___.SafeIndex((_this).F0(), _dafny.IntOfUint32((_1597_v119).Cardinality()))).Uint32()).(_dafny.Sequence), _1495_v47))
		r0 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_1495_v47, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sdiwmui"), (Companion_Default___.SafeIndex(_1493_v45, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sdiwmui")).Cardinality()))).Uint32(), (_this).F1())), (_this).F1())
		r1 = Companion_D2_.Create_DC4_(_1495_v47, !(_1442_v0))
		return r0, r1
	}
}
func (_this *C5) F1() _dafny.CodePoint {
	{
		return _this._f1
	}
}
func (_this *C5) F2() _dafny.Int {
	{
		return _this._f2
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
