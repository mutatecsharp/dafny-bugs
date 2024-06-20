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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	return ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(555), _dafny.IntOfInt64(992))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(555)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(992)) < 0) {
				_coll0.Add((_0_v0).Times(_dafny.IntOfInt64(-542)), Companion_D2_.Create_DC6_(_dafny.SeqOf(_dafny.IntOfInt64(123), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(623))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg0 _dafny.Int) interface{} {
						return coer0(arg0)
					}
				}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(161))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg1 _dafny.Int) interface{} {
						return coer1(arg1)
					}
				}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})), true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Cardinality()) >= 0
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 D0, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(55)), true)).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf((Companion_D2_.Create_DC6_(_dafny.SeqOf(_dafny.IntOfInt64(523), (_dafny.MultiSetOf(_dafny.CodePoint('p'), _dafny.CodePoint('l'))).Cardinality()))).Dtor_cf11(), _dafny.SeqOf(_dafny.IntOfInt64(156), _dafny.IntOfInt64(-520)))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Sequence
			_3_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((Companion_D2_.Create_DC6_(_dafny.SeqOf(_dafny.IntOfInt64(523), (_dafny.MultiSetOf(_dafny.CodePoint('p'), _dafny.CodePoint('l'))).Cardinality()))).Dtor_cf11(), _dafny.SeqOf(_dafny.IntOfInt64(156), _dafny.IntOfInt64(-520))), _3_v0) {
				_coll1.Add(_3_v0, !(true))
			}
		}
		return _coll1.ToMap()
	}())).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(615))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_4_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(742)
		})), _dafny.IntOfInt64(225))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _5_v1 _dafny.Sequence
			_5_v1 = interface{}(_compr_2).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(615))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_4_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(742)
			})), _dafny.IntOfInt64(225))).Contains(_5_v1) {
				_coll2.Add(_5_v1, true)
			}
		}
		return _coll2.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if (!(!(false))) || (true) {
		return (_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('t'), _dafny.CodePoint('c'), _dafny.CodePoint('h')), func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), true)).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _6_v0 _dafny.CodePoint
				_6_v0 = interface{}(_compr_3).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), true)).Contains(_6_v0) {
					_coll3.Add(_6_v0)
				}
			}
			return _coll3.ToSet()
		}())).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('e'), _dafny.CodePoint('n')))))
	} else {
		return _dafny.MultiSetOf(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _8_v1 _dafny.CodePoint
				_8_v1 = interface{}(_compr_4).(_dafny.CodePoint)
				if (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})))).Contains(_8_v1) {
					_coll4.Add(_8_v1)
				}
			}
			return _coll4.ToSet()
		}(), _dafny.SetOf(_dafny.CodePoint('y'), _dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('n')), func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.UnicodeSeqOfUtf8Bytes("iasqoksj"))).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _9_v2 _dafny.CodePoint
				_9_v2 = interface{}(_compr_5).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.UnicodeSeqOfUtf8Bytes("iasqoksj"))).Contains(_9_v2) {
					_coll5.Add(_9_v2)
				}
			}
			return _coll5.ToSet()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(183))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_10_i0 _dafny.Int) _dafny.Set {
		return func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.CodePoint('q'))).Keys().Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _11_v0 _dafny.CodePoint
				_11_v0 = interface{}(_compr_6).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.CodePoint('q'))).Contains(_11_v0) {
					_coll6.Add(_11_v0)
				}
			}
			return _coll6.ToSet()
		}()
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(872))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_12_i1 _dafny.Int) _dafny.Set {
		return _dafny.SetOf(_dafny.CodePoint('o'))
	})))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('o')
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("anbp"), _dafny.UnicodeSeqOfUtf8Bytes("qmbavh")), _dafny.UnicodeSeqOfUtf8Bytes("e"))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqOf((_dafny.SetOf(false, false, true)).Cardinality())
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(148))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_13_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-169)))).Cardinality())
		}))
	})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, !(!(true)))).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-274))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(525), _dafny.IntOfInt64(504))).Cardinality(), _dafny.IntOfInt64(-521)), _dafny.SeqOf(_dafny.IntOfInt64(174), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(false, false)).Cardinality()), _dafny.CodePoint('p'))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("jogx")
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.MultiSetOf(true, true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sxblvklw")).Cardinality()), _dafny.IntOfInt64(135)))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hvymmqds")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(true))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ygsxci")).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(false)))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(105))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_14_i0 _dafny.Int) _dafny.Set {
		return (func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-371))).Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _15_v0 _dafny.Int
				_15_v0 = interface{}(_compr_7).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-371)), _15_v0) {
					_coll7.Add((_15_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(792))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(504))).Cardinality())).Cardinality()))
				}
			}
			return _coll7.ToSet()
		}()).Union(_dafny.SetOf(_dafny.IntOfInt64(101), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkuhorvqx")).Cardinality()))).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _16_v1 _dafny.Int
				_16_v1 = interface{}(_compr_8).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkuhorvqx")).Cardinality()))).Contains(_16_v1) {
					_coll8.Add((_16_v1).Minus(_dafny.IntOfInt64(-619)), Companion_D9_.Create_DC23_(_dafny.MultiSetOf(true)))
				}
			}
			return _coll8.ToMap()
		}()).Cardinality())).Cardinality())))
	}))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-459))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_17_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
	})), _dafny.SeqOf((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(786), _dafny.IntOfInt64(626))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(786)).Cmp(_18_v0) <= 0) && ((_18_v0).Cmp(_dafny.IntOfInt64(626)) < 0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_18_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("olteeivdc")).Cardinality())))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-830))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_19_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(386)
	}))).Cardinality())))).IsProperSubsetOf(_dafny.SetOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(885), (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("mxe"))).Cardinality(), _dafny.IntOfInt64(-13)), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(12))).Cardinality(), _dafny.IntOfInt64(170), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tjay")).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(890)))) {
		if true {
			return _dafny.CodePoint('d')
		} else {
			return _dafny.CodePoint('q')
		}
	} else if true {
		return _dafny.CodePoint('d')
	} else {
		return _dafny.CodePoint('o')
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 bool, globalState *GlobalState) D2 {
	if true {
		return Companion_D2_.Create_DC7_(_dafny.SeqOf(_dafny.CodePoint('h')), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eurek")).Cardinality()), (_dafny.SetOf(_dafny.CodePoint('u'), _dafny.CodePoint('d'))).Cardinality(), true)
	} else {
		return Companion_D2_.Create_DC7_(_dafny.SeqOf(_dafny.CodePoint('v')), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()), _dafny.IntOfInt64(798), !(false))
	}
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(102))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('w'))).Cardinality())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(47))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(686))))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Set, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false, (Companion_D9_.Create_DC25_(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality()), true)).Dtor_cf41(), false, true)).Difference(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(782), _dafny.IntOfInt64(-958))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _20_v0 _dafny.Int
			_20_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(782)).Cmp(_20_v0) <= 0) && ((_20_v0).Cmp(_dafny.IntOfInt64(-958)) < 0) {
				_coll10.Add((_20_v0).Minus(_dafny.IntOfInt64(86)))
			}
		}
		return _coll10.ToSet()
	}()).Intersection(_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(411), _dafny.IntOfInt64(918), _dafny.IntOfInt64(498), _dafny.IntOfInt64(372), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(656))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_21_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(75)
	})))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(478)))).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dslxbs")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.CodePoint('q'))
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())).Times(_dafny.IntOfInt64(742)), _dafny.IntOfInt64(652))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jjjqasxw"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-236))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), _dafny.UnicodeSeqOfUtf8Bytes("idsrwnse")))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, p2 _dafny.Set, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source0 D2 = Companion_D2_.Create_DC7_(_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('a'), _dafny.CodePoint('y')), (_dafny.Zero).Minus(_dafny.IntOfInt64(-221)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(625))).Cardinality()), false)
	_ = _source0
	if _source0.Is_DC7() {
		var _23___mcc_h0 _dafny.Sequence = _source0.Get_().(D2_DC7).Cf12
		_ = _23___mcc_h0
		var _24___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC7).Cf13
		_ = _24___mcc_h1
		var _25___mcc_h2 _dafny.Int = _source0.Get_().(D2_DC7).Cf14
		_ = _25___mcc_h2
		var _26___mcc_h3 bool = _source0.Get_().(D2_DC7).Cf15
		_ = _26___mcc_h3
		var _27_cf15 bool = _26___mcc_h3
		_ = _27_cf15
		var _28_cf14 _dafny.Int = _25___mcc_h2
		_ = _28_cf14
		var _29_cf13 _dafny.Int = _24___mcc_h1
		_ = _29_cf13
		var _30_cf12 _dafny.Sequence = _23___mcc_h0
		_ = _30_cf12
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_29_cf13, _27_cf15)).Merge(func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(109), _dafny.IntOfInt64(892))); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _31_v0 _dafny.Int
				_31_v0 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(109)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(892)) < 0) {
					_coll11.Add(Companion_Default___.SafeDivisionInt(_31_v0, _28_cf14), _27_cf15)
				}
			}
			return _coll11.ToMap()
		}())
	} else if _source0.Is_DC8() {
		var _32___mcc_h4 bool = _source0.Get_().(D2_DC8).Cf16
		_ = _32___mcc_h4
		var _33___mcc_h5 bool = _source0.Get_().(D2_DC8).Cf17
		_ = _33___mcc_h5
		var _34___mcc_h6 bool = _source0.Get_().(D2_DC8).Cf18
		_ = _34___mcc_h6
		var _35_cf18 bool = _34___mcc_h6
		_ = _35_cf18
		var _36_cf17 bool = _33___mcc_h5
		_ = _36_cf17
		var _37_cf16 bool = _32___mcc_h4
		_ = _37_cf16
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_cf17, _dafny.IntOfInt64(741))).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(142), _36_cf17))
	} else {
		var _38___mcc_h7 _dafny.Sequence = _source0.Get_().(D2_DC6).Cf11
		_ = _38___mcc_h7
		var _39_cf11 _dafny.Sequence = _38___mcc_h7
		_ = _39_cf11
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(138), true)
	}
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.CodePoint {
	var _source1 D11 = (func() D11 {
		if true {
			return Companion_D11_.Create_DC30_(true)
		}
		return Companion_D11_.Create_DC30_(true)
	})()
	_ = _source1
	if _source1.Is_DC28() {
		var _40___mcc_h0 bool = _source1.Get_().(D11_DC28).Cf44
		_ = _40___mcc_h0
		var _41_cf44 bool = _40___mcc_h0
		_ = _41_cf44
		return _dafny.CodePoint('y')
	} else if _source1.Is_DC29() {
		var _42___mcc_h1 _dafny.Int = _source1.Get_().(D11_DC29).Cf45
		_ = _42___mcc_h1
		var _43___mcc_h2 _dafny.Int = _source1.Get_().(D11_DC29).Cf46
		_ = _43___mcc_h2
		var _44___mcc_h3 _dafny.Int = _source1.Get_().(D11_DC29).Cf47
		_ = _44___mcc_h3
		var _45___mcc_h4 _dafny.Int = _source1.Get_().(D11_DC29).Cf48
		_ = _45___mcc_h4
		var _46_cf48 _dafny.Int = _45___mcc_h4
		_ = _46_cf48
		var _47_cf47 _dafny.Int = _44___mcc_h3
		_ = _47_cf47
		var _48_cf46 _dafny.Int = _43___mcc_h2
		_ = _48_cf46
		var _49_cf45 _dafny.Int = _42___mcc_h1
		_ = _49_cf45
		return _dafny.CodePoint('a')
	} else if _source1.Is_DC30() {
		var _50___mcc_h5 bool = _source1.Get_().(D11_DC30).Cf49
		_ = _50___mcc_h5
		var _51_cf49 bool = _50___mcc_h5
		_ = _51_cf49
		return _dafny.CodePoint('u')
	} else if _source1.Is_DC27() {
		var _52___mcc_h6 _dafny.Set = _source1.Get_().(D11_DC27).Cf43
		_ = _52___mcc_h6
		var _53_cf43 _dafny.Set = _52___mcc_h6
		_ = _53_cf43
		return _dafny.CodePoint('p')
	} else {
		var _54___mcc_h7 D11 = _source1.Get_().(D11_DC31).Cf50
		_ = _54___mcc_h7
		var _55_cf50 D11 = _54___mcc_h7
		_ = _55_cf50
		if false {
			return _dafny.CodePoint('a')
		} else {
			return _dafny.CodePoint('v')
		}
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(568))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_56_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-181))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-152))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_57_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-182)
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("eolvsmea"))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(673), _dafny.IntOfInt64(562))))
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(true))).Intersection((_dafny.MultiSetOf(!(true))).Intersection(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(890), _dafny.IntOfInt64(345))); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _58_v0 _dafny.Int
			_58_v0 = interface{}(_compr_12).(_dafny.Int)
			if ((_dafny.IntOfInt64(890)).Cmp(_58_v0) <= 0) && ((_58_v0).Cmp(_dafny.IntOfInt64(345)) < 0) {
				_coll12.Add(Companion_Default___.SafeModuloInt(_58_v0, _dafny.IntOfInt64(137)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gmgrm")).Cardinality())))
			}
		}
		return _coll12.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, globalState *GlobalState) _dafny.Map {
	var _source2 D5 = Companion_D5_.Create_DC15_(Companion_D5_.Create_DC12_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-443))).Cardinality()))))
	_ = _source2
	if _source2.Is_DC13() {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D11_.Create_DC28_(false)).Dtor_cf44(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((Companion_D9_.Create_DC25_(true, (_dafny.SetOf(false)).Cardinality(), false)).Dtor_cf40())).Cardinality()), _dafny.IntOfInt64(-792))).Cardinality()))).Cardinality()))).Merge(func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _59_v1 _dafny.Sequence
					_59_v1 = interface{}(_compr_14).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Contains(_59_v1) {
						_coll14.Add(_59_v1, _dafny.IntOfInt64(-606))
					}
				}
				return _coll14.ToMap()
			}()).Keys().Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _60_v0 _dafny.Sequence
				_60_v0 = interface{}(_compr_13).(_dafny.Sequence)
				if (func() _dafny.Map {
					var _coll15 = _dafny.NewMapBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Elements()); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _59_v1 _dafny.Sequence
						_59_v1 = interface{}(_compr_15).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Contains(_59_v1) {
							_coll15.Add(_59_v1, _dafny.IntOfInt64(-606))
						}
					}
					return _coll15.ToMap()
				}()).Contains(_60_v0) {
					_coll13.Add(_60_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(478), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(150))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}(func(_61_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					}))).Cardinality()))).Cardinality()))
				}
			}
			return _coll13.ToMap()
		}())
	} else if _source2.Is_DC14() {
		var _62___mcc_h0 bool = _source2.Get_().(D5_DC14).Cf23
		_ = _62___mcc_h0
		var _63_cf23 bool = _62___mcc_h0
		_ = _63_cf23
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_63_cf23, _63_cf23, _63_cf23), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_cf23, _dafny.IntOfInt64(683)))
	} else if _source2.Is_DC12() {
		var _64___mcc_h1 _dafny.Map = _source2.Get_().(D5_DC12).Cf22
		_ = _64___mcc_h1
		var _65_cf22 _dafny.Map = _64___mcc_h1
		_ = _65_cf22
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, false), _65_cf22)
	} else {
		var _66___mcc_h2 D5 = _source2.Get_().(D5_DC15).Cf24
		_ = _66___mcc_h2
		var _67_cf24 D5 = _66___mcc_h2
		_ = _67_cf24
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(164)))).Merge(func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, false), _dafny.UnicodeSeqOfUtf8Bytes("uevfjb"))).Keys().Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _68_v2 _dafny.Sequence
				_68_v2 = interface{}(_compr_16).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, false), _dafny.UnicodeSeqOfUtf8Bytes("uevfjb"))).Contains(_68_v2) {
					_coll16.Add(_68_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfInt64(-902))))
				}
			}
			return _coll16.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source3 D11 = Companion_D11_.Create_DC28_(true)
	_ = _source3
	if _source3.Is_DC28() {
		var _69___mcc_h0 bool = _source3.Get_().(D11_DC28).Cf44
		_ = _69___mcc_h0
		var _70_cf44 bool = _69___mcc_h0
		_ = _70_cf44
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(320), (_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(189))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_71_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_72_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality())
		})))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(402), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vavsi")).Cardinality()))).Cardinality()))
	} else if _source3.Is_DC29() {
		var _73___mcc_h1 _dafny.Int = _source3.Get_().(D11_DC29).Cf45
		_ = _73___mcc_h1
		var _74___mcc_h2 _dafny.Int = _source3.Get_().(D11_DC29).Cf46
		_ = _74___mcc_h2
		var _75___mcc_h3 _dafny.Int = _source3.Get_().(D11_DC29).Cf47
		_ = _75___mcc_h3
		var _76___mcc_h4 _dafny.Int = _source3.Get_().(D11_DC29).Cf48
		_ = _76___mcc_h4
		var _77_cf48 _dafny.Int = _76___mcc_h4
		_ = _77_cf48
		var _78_cf47 _dafny.Int = _75___mcc_h3
		_ = _78_cf47
		var _79_cf46 _dafny.Int = _74___mcc_h2
		_ = _79_cf46
		var _80_cf45 _dafny.Int = _73___mcc_h1
		_ = _80_cf45
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_cf47, _78_cf47)
	} else if _source3.Is_DC30() {
		var _81___mcc_h5 bool = _source3.Get_().(D11_DC30).Cf49
		_ = _81___mcc_h5
		var _82_cf49 bool = _81___mcc_h5
		_ = _82_cf49
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(109), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_cf49, false)).Cardinality())
	} else if _source3.Is_DC27() {
		var _83___mcc_h6 _dafny.Set = _source3.Get_().(D11_DC27).Cf43
		_ = _83___mcc_h6
		var _84_cf43 _dafny.Set = _83___mcc_h6
		_ = _84_cf43
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(80), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(408))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false, false)).Cardinality(), _dafny.IntOfInt64(172)))
	} else {
		var _85___mcc_h7 D11 = _source3.Get_().(D11_DC31).Cf50
		_ = _85___mcc_h7
		var _86_cf50 D11 = _85___mcc_h7
		_ = _86_cf50
		return func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(698))).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _87_v0 _dafny.Int
				_87_v0 = interface{}(_compr_17).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(698))).Contains(_87_v0) {
					_coll17.Add(Companion_Default___.SafeDivisionInt(_87_v0, _dafny.IntOfInt64(-374)), _dafny.IntOfInt64(531))
				}
			}
			return _coll17.ToMap()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aynax")).Cardinality())))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_88_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-406))
	}))))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, globalState *GlobalState) D4 {
	var _source4 D11 = Companion_D11_.Create_DC29_(_dafny.IntOfInt64(-157), _dafny.IntOfInt64(677), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(48))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_89_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-188)
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_90_i1 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality())).Cardinality()))
	})))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tsx")).Cardinality()))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(681))
	_ = _source4
	if _source4.Is_DC28() {
		var _91___mcc_h0 bool = _source4.Get_().(D11_DC28).Cf44
		_ = _91___mcc_h0
		var _92_cf44 bool = _91___mcc_h0
		_ = _92_cf44
		return Companion_D4_.Create_DC11_(_92_cf44)
	} else if _source4.Is_DC29() {
		var _93___mcc_h1 _dafny.Int = _source4.Get_().(D11_DC29).Cf45
		_ = _93___mcc_h1
		var _94___mcc_h2 _dafny.Int = _source4.Get_().(D11_DC29).Cf46
		_ = _94___mcc_h2
		var _95___mcc_h3 _dafny.Int = _source4.Get_().(D11_DC29).Cf47
		_ = _95___mcc_h3
		var _96___mcc_h4 _dafny.Int = _source4.Get_().(D11_DC29).Cf48
		_ = _96___mcc_h4
		var _97_cf48 _dafny.Int = _96___mcc_h4
		_ = _97_cf48
		var _98_cf47 _dafny.Int = _95___mcc_h3
		_ = _98_cf47
		var _99_cf46 _dafny.Int = _94___mcc_h2
		_ = _99_cf46
		var _100_cf45 _dafny.Int = _93___mcc_h1
		_ = _100_cf45
		return Companion_D4_.Create_DC11_(false)
	} else if _source4.Is_DC30() {
		var _101___mcc_h5 bool = _source4.Get_().(D11_DC30).Cf49
		_ = _101___mcc_h5
		var _102_cf49 bool = _101___mcc_h5
		_ = _102_cf49
		return Companion_D4_.Create_DC11_(_102_cf49)
	} else if _source4.Is_DC27() {
		var _103___mcc_h6 _dafny.Set = _source4.Get_().(D11_DC27).Cf43
		_ = _103___mcc_h6
		var _104_cf43 _dafny.Set = _103___mcc_h6
		_ = _104_cf43
		return Companion_D4_.Create_DC11_(false)
	} else {
		var _105___mcc_h7 D11 = _source4.Get_().(D11_DC31).Cf50
		_ = _105___mcc_h7
		var _106_cf50 D11 = _105___mcc_h7
		_ = _106_cf50
		return Companion_D4_.Create_DC11_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm42(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fgdv")).Cardinality()))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfInt64(-820), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), _dafny.IntOfInt64(103), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aheg")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qehktwkga")).Cardinality()))).Cardinality())).Cardinality())), (_dafny.IntOfInt64(-193)).Plus(_dafny.IntOfUint32((_dafny.SeqOf(false, !(!(!(true))))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm43(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("yaynwj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(188))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_107_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.UnicodeSeqOfUtf8Bytes("yrbinxrmv"))).Difference((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.UnicodeSeqOfUtf8Bytes("fmj"))).Difference((Companion_D12_.Create_DC32_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("mtvdgkj"), _dafny.UnicodeSeqOfUtf8Bytes("xxyl"))))).Dtor_cf51()))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(855))).Uint32(), func(coer23 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_108_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC2_((Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(27))).Uint32(), func(coer24 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_109_i1 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(true)
		}))).Cardinality()), _dafny.IntOfInt64(894)))).Dtor_cf3())
	}))
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(false, false))
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(821), _dafny.IntOfInt64(144))); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _110_v0 _dafny.Int
			_110_v0 = interface{}(_compr_18).(_dafny.Int)
			if ((_dafny.IntOfInt64(821)).Cmp(_110_v0) <= 0) && ((_110_v0).Cmp(_dafny.IntOfInt64(144)) < 0) {
				_coll18.Add((_110_v0).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dav")).Cardinality()))), _dafny.IntOfInt64(-163))
			}
		}
		return _coll18.ToMap()
	}(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll19 = _dafny.NewMapBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(767), _dafny.IntOfInt64(440))); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _111_v1 _dafny.Int
			_111_v1 = interface{}(_compr_19).(_dafny.Int)
			if ((_dafny.IntOfInt64(767)).Cmp(_111_v1) <= 0) && ((_111_v1).Cmp(_dafny.IntOfInt64(440)) < 0) {
				_coll19.Add(Companion_Default___.SafeDivisionInt(_111_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((Companion_D2_.Create_DC6_(_dafny.SeqOf(_dafny.IntOfInt64(185), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kvwjkcv")).Cardinality()), _dafny.IntOfInt64(685)))).Dtor_cf11()).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(152), (_dafny.SetOf(!(true))).Cardinality())).Cardinality())).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(70), _dafny.IntOfInt64(800))).Cardinality())
			}
		}
		return _coll19.ToMap()
	}()).Cardinality()), _dafny.IntOfInt64(489))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(955), _dafny.IntOfInt64(299))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(383))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yi")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(203))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_112_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality())), _dafny.IntOfInt64(120))))
}
func (_static *CompanionStruct_Default___) Fm47(p0 D0, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('f'), (func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('i')
		}
		return _dafny.CodePoint('l')
	})(), _dafny.CodePoint('o'))
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.CodePoint, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(_dafny.SeqOf(_dafny.IntOfInt64(794), _dafny.IntOfInt64(-102)))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _113_v0 _dafny.Int
	_ = _113_v0
	_113_v0 = _dafny.IntOfInt64(-3)
	var _114_v1 _dafny.Set
	_ = _114_v1
	_114_v1 = _dafny.SetOf(_113_v0)
	var _115_v2 bool
	_ = _115_v2
	_115_v2 = false
	var _116_v3 _dafny.Sequence
	_ = _116_v3
	_116_v3 = _dafny.SeqOf(true, _115_v2)
	var _117_v4 *C2
	_ = _117_v4
	var _nw0 *C2 = New_C2_()
	_ = _nw0
	_nw0.Ctor__(_115_v2)
	_117_v4 = _nw0
	var _118_v5 _dafny.MultiSet
	_ = _118_v5
	_118_v5 = _dafny.MultiSetOf(_117_v4, _117_v4)
	var _119_v6 *C4
	_ = _119_v6
	var _nw1 *C4 = New_C4_()
	_ = _nw1
	_nw1.Ctor__(_115_v2, false)
	_119_v6 = _nw1
	var _120_v7 _dafny.Set
	_ = _120_v7
	_120_v7 = _dafny.SetOf(_119_v6)
	var _121_v8 _dafny.Array
	_ = _121_v8
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
	_ = _nw2
	_121_v8 = _nw2
	var _122_v9 _dafny.Map
	_ = _122_v9
	_122_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v8, (_119_v6).F16())
	var _123_v10 _dafny.Map
	_ = _123_v10
	_123_v10 = _122_v9
	var _124_v11 _dafny.Map
	_ = _124_v11
	_124_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_119_v6).F16(), _115_v2)
	var _125_v12 D4
	_ = _125_v12
	_125_v12 = Companion_D4_.Create_DC11_((_119_v6).F16())
	var _126_v13 _dafny.Map
	_ = _126_v13
	_126_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v0, _113_v0)
	var _127_v14 _dafny.Sequence
	_ = _127_v14
	_127_v14 = _dafny.UnicodeSeqOfUtf8Bytes("wm")
	var _128_v15 _dafny.CodePoint
	_ = _128_v15
	_128_v15 = _dafny.CodePoint('v')
	var _129_v16 _dafny.MultiSet
	_ = _129_v16
	_129_v16 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("yyl"), _dafny.Companion_Sequence_.Update(_127_v14, (Companion_Default___.SafeIndex(_113_v0, _dafny.IntOfUint32((_127_v14).Cardinality()))).Uint32(), _128_v15))
	var _130_v17 _dafny.MultiSet
	_ = _130_v17
	_130_v17 = _dafny.MultiSetOf(_116_v3)
	var _131_v18 _dafny.Array
	_ = _131_v18
	var _nwElement0_0 bool = (_113_v0).Cmp((_114_v1).Cardinality()) < 0
	_ = _nwElement0_0
	var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(29))
	_ = _nw3
	(_nw3).ArraySet1(_nwElement0_0, 0)
	(_nw3).ArraySet1((_116_v3).Select((Companion_Default___.SafeIndex((_118_v5).Cardinality(), _dafny.IntOfUint32((_116_v3).Cardinality()))).Uint32()).(bool), 1)
	(_nw3).ArraySet1((_120_v7).IsProperSubsetOf(_120_v7), 2)
	(_nw3).ArraySet1(_115_v2, 3)
	(_nw3).ArraySet1(!((_119_v6).F16()), 4)
	(_nw3).ArraySet1((_119_v6).F16(), 5)
	(_nw3).ArraySet1((_119_v6).F16(), 6)
	(_nw3).ArraySet1(_115_v2, 7)
	(_nw3).ArraySet1(true, 8)
	(_nw3).ArraySet1(_115_v2, 9)
	(_nw3).ArraySet1(_115_v2, 10)
	(_nw3).ArraySet1(_115_v2, 11)
	(_nw3).ArraySet1((_119_v6).F16(), 12)
	(_nw3).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v8, _115_v2)).Equals((_122_v9).Update(_121_v8, (_119_v6).F16())), 13)
	(_nw3).ArraySet1((func() bool {
		if (_122_v9).Contains(_121_v8) {
			return (_122_v9).Get(_121_v8).(bool)
		}
		return _115_v2
	})(), 14)
	(_nw3).ArraySet1((_115_v2) == (_115_v2), 15)
	(_nw3).ArraySet1((Companion_Default___.Fm25(_114_v1, _124_v11, _113_v0, _115_v2, globalState)).IsDisjointFrom(Companion_Default___.Fm25(_114_v1, Companion_Default___.Fm20(_115_v2, _113_v0, (_119_v6).F16(), globalState), _113_v0, (_119_v6).F16(), globalState)), 16)
	(_nw3).ArraySet1(_115_v2, 17)
	(_nw3).ArraySet1((_119_v6).F16(), 18)
	(_nw3).ArraySet1((_119_v6).F16(), 19)
	(_nw3).ArraySet1(false, 20)
	(_nw3).ArraySet1((_dafny.IntOfInt64(-421)).Cmp(_113_v0) >= 0, 21)
	(_nw3).ArraySet1((func() bool {
		if (_119_v6).F16() {
			return _115_v2
		}
		return (_119_v6).F16()
	})(), 22)
	(_nw3).ArraySet1(!((_125_v12).Dtor_cf21()), 23)
	(_nw3).ArraySet1(Companion_Default___.Fm0(_126_v13, _113_v0, globalState), 24)
	(_nw3).ArraySet1((_129_v16).IsSubsetOf(_dafny.MultiSetOf(_127_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_132_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	})), _127_v14)), 25)
	(_nw3).ArraySet1((!(false)) || (_115_v2), 26)
	(_nw3).ArraySet1(!((func() bool {
		if (_119_v6).F16() {
			return _115_v2
		}
		return (_119_v6).F16()
	})()), 27)
	(_nw3).ArraySet1((_113_v0).Cmp((func() _dafny.Int {
		if (_130_v17).Contains(_116_v3) {
			return (_130_v17).Multiplicity(_116_v3)
		}
		return _113_v0
	})()) <= 0, 28)
	_131_v18 = _nw3
	_131_v18 = _131_v18
	var _133_v19 _dafny.Sequence
	_ = _133_v19
	_133_v19 = _dafny.SeqOf(_116_v3)
	var _source5 _dafny.Sequence = (_133_v19).Select((Companion_Default___.SafeIndex(_113_v0, _dafny.IntOfUint32((_133_v19).Cardinality()))).Uint32()).(_dafny.Sequence)
	_ = _source5
	var _134___mcc_h0 _dafny.Sequence = _source5
	_ = _134___mcc_h0
	var _135_cf34 _dafny.Sequence = _134___mcc_h0
	_ = _135_cf34
	r0 = _113_v0
	_115_v2 = (_119_v6).F16()
	var _136_v20 _dafny.Array
	_ = _136_v20
	var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
	_ = _nw4
	_136_v20 = _nw4
	var _137_v21 _dafny.MultiSet
	_ = _137_v21
	_137_v21 = _dafny.MultiSetOf(_128_v15)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_136_v20), 0))
	_ = _index0
	(_136_v20).ArraySet1((_137_v21).Intersection(_137_v21), (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_136_v20), 0))
	_ = _index1
	var _rhs0 _dafny.MultiSet = (_137_v21).Union(_dafny.MultiSetOf(_128_v15, _dafny.CodePoint('s')))
	_ = _rhs0
	var _rhs1 bool = !(_115_v2)
	_ = _rhs1
	var _lhs0 _dafny.Array = _136_v20
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_136_v20), 0))
	_ = _lhs1
	(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
	_115_v2 = _rhs1
	_115_v2 = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rokcgkova")).Cardinality()))).Cmp(_113_v0) <= 0
	var _138_v22 D1
	_ = _138_v22
	_138_v22 = Companion_D1_.Create_DC3_(_dafny.IntOfInt64(104))
	var _139_i1 _dafny.Int
	_ = _139_i1
	_139_i1 = _dafny.Zero
	{
		var _pat_let_tv0 = _119_v6
		_ = _pat_let_tv0
		var _pat_let_tv1 = _119_v6
		_ = _pat_let_tv1
		var _pat_let_tv2 = _119_v6
		_ = _pat_let_tv2
		var _pat_let_tv3 = _115_v2
		_ = _pat_let_tv3
		var _pat_let_tv4 = _119_v6
		_ = _pat_let_tv4
		var _pat_let_tv5 = _113_v0
		_ = _pat_let_tv5
		for func(_source6 D1) bool {
			if _source6.Is_DC3() {
				var _144___mcc_h1 _dafny.Int = _source6.Get_().(D1_DC3).Cf4
				_ = _144___mcc_h1
				var _145_cf4 _dafny.Int = _144___mcc_h1
				_ = _145_cf4
				return true
			} else if _source6.Is_DC4() {
				var _146___mcc_h2 _dafny.Array = _source6.Get_().(D1_DC4).Cf5
				_ = _146___mcc_h2
				var _147___mcc_h3 _dafny.Int = _source6.Get_().(D1_DC4).Cf6
				_ = _147___mcc_h3
				var _148___mcc_h4 _dafny.CodePoint = _source6.Get_().(D1_DC4).Cf7
				_ = _148___mcc_h4
				var _149_cf7 _dafny.CodePoint = _148___mcc_h4
				_ = _149_cf7
				var _150_cf6 _dafny.Int = _147___mcc_h3
				_ = _150_cf6
				var _151_cf5 _dafny.Array = _146___mcc_h2
				_ = _151_cf5
				return (_dafny.SetOf((_pat_let_tv0).F16())).IsSubsetOf(_dafny.SetOf((_pat_let_tv1).F16()))
			} else if _source6.Is_DC5() {
				var _152___mcc_h5 _dafny.Sequence = _source6.Get_().(D1_DC5).Cf8
				_ = _152___mcc_h5
				var _153___mcc_h6 _dafny.Array = _source6.Get_().(D1_DC5).Cf9
				_ = _153___mcc_h6
				var _154___mcc_h7 _dafny.Int = _source6.Get_().(D1_DC5).Cf10
				_ = _154___mcc_h7
				var _155_cf10 _dafny.Int = _154___mcc_h7
				_ = _155_cf10
				var _156_cf9 _dafny.Array = _153___mcc_h6
				_ = _156_cf9
				var _157_cf8 _dafny.Sequence = _152___mcc_h5
				_ = _157_cf8
				return (_pat_let_tv2).F16()
			} else {
				var _158___mcc_h8 _dafny.Map = _source6.Get_().(D1_DC2).Cf3
				_ = _158___mcc_h8
				var _159_cf3 _dafny.Map = _158___mcc_h8
				_ = _159_cf3
				return ((_dafny.SetOf(_pat_let_tv3, (_pat_let_tv4).F16())).Cardinality()).Cmp(_pat_let_tv5) == 0
			}
		}((func() D1 {
			if _115_v2 {
				return _138_v22
			}
			return _138_v22
		})()) {
			{
				if (_139_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_139_i1 = (_139_i1).Plus(_dafny.One)
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_121_v8), 0))
				_ = _index2
				(_121_v8).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nntti"), _127_v14), (Companion_Default___.SafeIndex(_113_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nntti"), _127_v14)).Cardinality()))).Uint32(), _128_v15)).Cardinality()), (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_121_v8), 0))
				_ = _index3
				(_121_v8).ArraySet1((_113_v0).Plus(_113_v0), (_index3).Int())
				var _140_v23 *C4
				_ = _140_v23
				var _nw5 *C4 = New_C4_()
				_ = _nw5
				_nw5.Ctor__(((_121_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_121_v8), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_121_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_121_v8), 0))).Int()).(_dafny.Int))) != 0, (_114_v1).IsSubsetOf((Companion_D11_.Create_DC27_(_114_v1)).Dtor_cf43()))
				_140_v23 = _nw5
				var _141_v24 _dafny.Map
				_ = _141_v24
				_141_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_113_v0), _131_v18)
				var _142_v25 D1
				_ = _142_v25
				_142_v25 = Companion_D1_.Create_DC4_(_131_v18, _dafny.IntOfUint32((_127_v14).Cardinality()), _128_v15)
				_131_v18 = (func() _dafny.Array {
					if (_141_v24).Contains(_113_v0) {
						return (_141_v24).Get(_113_v0).(_dafny.Array)
					}
					return (_142_v25).Dtor_cf5()
				})()
				var _143_v26 _dafny.Array
				_ = _143_v26
				var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw6
				_143_v26 = _nw6
				var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_143_v26), 0))
				_ = _index4
				(_143_v26).ArraySet1(_131_v18, (_index4).Int())
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_143_v26), 0))
				_ = _index5
				(_143_v26).ArraySet1(_131_v18, (_index5).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _160_v27 _dafny.Set
	_ = _160_v27
	_160_v27 = _dafny.SetOf(_124_v11, _124_v11, _124_v11)
	_115_v2 = (_160_v27).Equals(_160_v27)
	var _161_v28 _dafny.Array
	_ = _161_v28
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(22))
	_ = _nw7
	_161_v28 = _nw7
	var _162_v29 D7
	_ = _162_v29
	_162_v29 = Companion_D7_.Create_DC20_(_dafny.IntOfInt64(-722), _128_v15)
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_161_v28), 0))
	_ = _index6
	(_161_v28).ArraySet1(_162_v29, (_index6).Int())
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_161_v28), 0))
	_ = _index7
	(_161_v28).ArraySet1(_162_v29, (_index7).Int())
	_113_v0 = _113_v0
	r0 = _113_v0
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _163_v0 bool
	_ = _163_v0
	_163_v0 = false
	var _164_v1 _dafny.Sequence
	_ = _164_v1
	_164_v1 = _dafny.SeqOf(_163_v0, !(_163_v0), _163_v0, !(_163_v0))
	var _165_v2 _dafny.Int
	_ = _165_v2
	_165_v2 = _dafny.IntOfInt64(-199)
	var _166_v3 _dafny.Array
	_ = _166_v3
	var _nwElement0_1 bool = _163_v0
	_ = _nwElement0_1
	var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(27))
	_ = _nw8
	(_nw8).ArraySet1(_nwElement0_1, 0)
	(_nw8).ArraySet1(_163_v0, 1)
	(_nw8).ArraySet1(_163_v0, 2)
	(_nw8).ArraySet1(_163_v0, 3)
	(_nw8).ArraySet1(_163_v0, 4)
	(_nw8).ArraySet1(_163_v0, 5)
	(_nw8).ArraySet1(_163_v0, 6)
	(_nw8).ArraySet1(_163_v0, 7)
	(_nw8).ArraySet1(true, 8)
	(_nw8).ArraySet1(_163_v0, 9)
	(_nw8).ArraySet1(_163_v0, 10)
	(_nw8).ArraySet1(_163_v0, 11)
	(_nw8).ArraySet1(!(_163_v0), 12)
	(_nw8).ArraySet1(true, 13)
	(_nw8).ArraySet1(_163_v0, 14)
	(_nw8).ArraySet1(_163_v0, 15)
	(_nw8).ArraySet1(_163_v0, 16)
	(_nw8).ArraySet1(_163_v0, 17)
	(_nw8).ArraySet1(_163_v0, 18)
	(_nw8).ArraySet1(_163_v0, 19)
	(_nw8).ArraySet1(_163_v0, 20)
	(_nw8).ArraySet1(_163_v0, 21)
	(_nw8).ArraySet1(_163_v0, 22)
	(_nw8).ArraySet1((_164_v1).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_164_v1).Cardinality()))).Uint32()).(bool), 23)
	(_nw8).ArraySet1(_163_v0, 24)
	(_nw8).ArraySet1(_163_v0, 25)
	(_nw8).ArraySet1(_163_v0, 26)
	_166_v3 = _nw8
	var _167_v4 D0
	_ = _167_v4
	_167_v4 = Companion_D0_.Create_DC0_(false)
	var _168_v5 _dafny.Set
	_ = _168_v5
	_168_v5 = _dafny.SetOf(_164_v1, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_163_v0, (_167_v4).Dtor_cf0(), !(false), _163_v0, _163_v0), (Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_dafny.SeqOf(_163_v0, (_167_v4).Dtor_cf0(), !(false), _163_v0, _163_v0)).Cardinality()))).Uint32(), false), _164_v1)
	var _169_v6 _dafny.Sequence
	_ = _169_v6
	_169_v6 = _dafny.UnicodeSeqOfUtf8Bytes("nftgk")
	var _170_globalState *GlobalState
	_ = _170_globalState
	var _nw9 *GlobalState = New_GlobalState_()
	_ = _nw9
	_nw9.Ctor__(_dafny.IntOfInt64(-586), _166_v3, true, false, _164_v1, (_168_v5).Union(_168_v5), false, _dafny.MultiSetOf(_165_v2), true, _dafny.IntOfInt64(721), _166_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(777))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}(func(_171_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), _169_v6), _dafny.IntOfInt64(604))
	_170_globalState = _nw9
	var _172_v7 D0
	_ = _172_v7
	_172_v7 = Companion_D0_.Create_DC1_(_163_v0, Companion_Default___.SafeDivisionInt(_165_v2, _165_v2))
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
	_ = _index8
	(_166_v3).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_169_v6, _169_v6), (_index8).Int())
	var _173_v8 _dafny.Map
	_ = _173_v8
	_173_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_165_v2), _dafny.IntOfInt64(833))
	var _174_v9 _dafny.MultiSet
	_ = _174_v9
	_174_v9 = _dafny.MultiSetOf(_165_v2)
	var _175_v10 _dafny.CodePoint
	_ = _175_v10
	_175_v10 = _dafny.CodePoint('d')
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
	_ = _index9
	var _rhs2 D0 = _172_v7
	_ = _rhs2
	var _rhs3 bool = Companion_Default___.Fm0(_173_v8, (_dafny.Zero).Minus((_174_v9).Cardinality()), _170_globalState)
	_ = _rhs3
	var _rhs4 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_169_v6, _dafny.UnicodeSeqOfUtf8Bytes("qtgkj")), _175_v10)
	_ = _rhs4
	var _rhs5 bool = _163_v0
	_ = _rhs5
	var _rhs6 _dafny.Array = _166_v3
	_ = _rhs6
	var _lhs2 _dafny.Array = _166_v3
	_ = _lhs2
	var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
	_ = _lhs3
	_172_v7 = _rhs2
	(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
	_163_v0 = _rhs4
	_163_v0 = _rhs5
	_166_v3 = _rhs6
	var _176_i1 _dafny.Int
	_ = _176_i1
	_176_i1 = _dafny.Zero
	{
		for false {
			{
				if (_176_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_176_i1 = (_176_i1).Plus(_dafny.One)
				var _177_v11 _dafny.Array
				_ = _177_v11
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_0
				var _nw10 _dafny.Array
				_ = _nw10
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw10 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.Sequence = (func(_178_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_179_i2 _dafny.Int) _dafny.Sequence {
							return _178_v1
						}
					})(_164_v1)
					_ = _init0
					var _element0_0 = _init0(_dafny.Zero)
					_ = _element0_0
					_nw10 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
					(_nw10).ArraySet1(_element0_0, 0)
					var _nativeLen0_0 = (_len0_0).Int()
					_ = _nativeLen0_0
					for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
						(_nw10).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
					}
				}
				_177_v11 = _nw10
				var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_177_v11), 0))
				_ = _index10
				(_177_v11).ArraySet1(_164_v1, (_index10).Int())
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_177_v11), 0))
				_ = _index11
				(_177_v11).ArraySet1(_164_v1, (_index11).Int())
				var _180_v12 _dafny.Array
				_ = _180_v12
				var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
				_ = _nw11
				_180_v12 = _nw11
				var _181_v13 _dafny.Map
				_ = _181_v13
				_181_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v3, Companion_Default___.Fm0(_173_v8, _165_v2, _170_globalState))
				var _182_v14 _dafny.Map
				_ = _182_v14
				_182_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v0, (func() bool {
					if (_181_v13).Contains(_166_v3) {
						return (_181_v13).Get(_166_v3).(bool)
					}
					return true
				})())
				var _183_v15 _dafny.Array
				_ = _183_v15
				var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw12
				_183_v15 = _nw12
				var _184_v16 _dafny.Set
				_ = _184_v16
				_184_v16 = _dafny.SetOf(_183_v15, _183_v15)
				var _185_v17 _dafny.Map
				_ = _185_v17
				_185_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_182_v14).Cardinality(), _184_v16)
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_180_v12), 0))
				_ = _index12
				(_180_v12).ArraySet1(((func() _dafny.Set {
					if (_185_v17).Contains(_dafny.IntOfInt64(-75)) {
						return (_185_v17).Get(_dafny.IntOfInt64(-75)).(_dafny.Set)
					}
					return _184_v16
				})()).Union(_184_v16), (_index12).Int())
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_180_v12), 0))
				_ = _index13
				(_180_v12).ArraySet1(_dafny.SetOf(_183_v15), (_index13).Int())
				_175_v10 = _175_v10
				var _186_v18 _dafny.Map
				_ = _186_v18
				_186_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wm"), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
				var _187_v19 _dafny.Int
				_ = _187_v19
				var _out0 _dafny.Int
				_ = _out0
				_out0 = Companion_Default___.M0(_186_v18, _170_globalState)
				_187_v19 = _out0
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _188_v20 _dafny.Array
	_ = _188_v20
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_1
	var _nw13 _dafny.Array
	_ = _nw13
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw13 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = (func(_189_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_190_i3 _dafny.Int) _dafny.Int {
				return (_190_i3).Plus(_189_v2)
			}
		})(_165_v2)
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
	_188_v20 = _nw13
	var _191_v21 _dafny.Map
	_ = _191_v21
	_191_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v2, _188_v20)
	_191_v21 = _191_v21
	var _192_v22 D1
	_ = _192_v22
	_192_v22 = Companion_D1_.Create_DC2_(_173_v8)
	var _193_i4 _dafny.Int
	_ = _193_i4
	_193_i4 = _dafny.Zero
	{
		for Companion_Default___.Fm0((_192_v22).Dtor_cf3(), _165_v2, _170_globalState) {
			{
				if (_193_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_193_i4 = (_193_i4).Plus(_dafny.One)
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
				_ = _index14
				(_166_v3).ArraySet1(Companion_Default___.Fm0(_173_v8, _165_v2, _170_globalState), (_index14).Int())
				var _194_v23 _dafny.Map
				_ = _194_v23
				_194_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v2, !(_163_v0))
				var _195_v24 _dafny.Sequence
				_ = _195_v24
				_195_v24 = _dafny.SeqOf(_165_v2)
				var _196_v25 _dafny.Map
				_ = _196_v25
				_196_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v23, _195_v24)
				var _197_v26 _dafny.Set
				_ = _197_v26
				_197_v26 = _dafny.SetOf((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool), _163_v0, _163_v0, !(false))
				var _198_v27 _dafny.Map
				_ = _198_v27
				_198_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_197_v26).Cardinality(), _195_v24)
				var _199_v28 _dafny.Map
				_ = _199_v28
				_199_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_196_v25).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v23, (func() _dafny.Sequence {
					if (_198_v27).Contains(_dafny.IntOfUint32((_169_v6).Cardinality())) {
						return (_198_v27).Get(_dafny.IntOfUint32((_169_v6).Cardinality())).(_dafny.Sequence)
					}
					return _195_v24
				})())), Companion_Default___.Fm0(_173_v8, _165_v2, _170_globalState))
				_199_v28 = (_199_v28).Update((_196_v25).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v23, _195_v24)), _163_v0)
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_188_v20), 0))
				_ = _index15
				(_188_v20).ArraySet1((func() _dafny.Set {
					var _coll20 = _dafny.NewBuilder()
					_ = _coll20
					for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(448), _dafny.IntOfInt64(86))); ; {
						_compr_20, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _200_v29 _dafny.Int
						_200_v29 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(448)).Cmp(_200_v29) <= 0) && ((_200_v29).Cmp(_dafny.IntOfInt64(86)) < 0) {
							_coll20.Add((_200_v29).Times(_165_v2))
						}
					}
					return _coll20.ToSet()
				}()).Cardinality(), (_index15).Int())
				var _201_v30 _dafny.Map
				_ = _201_v30
				_201_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v0, _dafny.IntOfUint32((_195_v24).Cardinality()))
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_188_v20), 0))
				_ = _index16
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
				_ = _index17
				var _rhs7 bool = false
				_ = _rhs7
				var _rhs8 _dafny.Int = _165_v2
				_ = _rhs8
				var _rhs9 _dafny.Array = _188_v20
				_ = _rhs9
				var _rhs10 _dafny.Int = (_dafny.IntOfInt64(782)).Plus(_165_v2)
				_ = _rhs10
				var _rhs11 bool = !((_201_v30).Merge(_201_v30)).Contains((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
				_ = _rhs11
				var _lhs4 _dafny.Array = _188_v20
				_ = _lhs4
				var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_188_v20), 0))
				_ = _lhs5
				var _lhs6 _dafny.Array = _166_v3
				_ = _lhs6
				var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
				_ = _lhs7
				_163_v0 = _rhs7
				(_lhs4).ArraySet1(_rhs8, (_lhs5).Int())
				_188_v20 = _rhs9
				_165_v2 = _rhs10
				(_lhs6).ArraySet1(_rhs11, (_lhs7).Int())
				var _202_v31 _dafny.Sequence
				_ = _202_v31
				_202_v31 = _dafny.SeqOf(_166_v3, _166_v3)
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
				_ = _index18
				(_166_v3).ArraySet1(!_dafny.Companion_Sequence_.Contains(_202_v31, _166_v3), (_index18).Int())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _203_v32 _dafny.Map
	_ = _203_v32
	_203_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)), _165_v2)
	var _hi0 _dafny.Int = _165_v2
	_ = _hi0
	for _204_i5 := ((_203_v32).Merge(_203_v32)).Cardinality(); _204_i5.Cmp(_hi0) < 0; _204_i5 = _204_i5.Plus(_dafny.One) {
		var _205_v33 _dafny.Array
		_ = _205_v33
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
		_ = _nw14
		_205_v33 = _nw14
		var _206_v34 _dafny.Sequence
		_ = _206_v34
		_206_v34 = _dafny.SeqOf(_165_v2)
		var _207_v35 _dafny.Map
		_ = _207_v35
		_207_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_206_v34, _206_v34), (func() bool {
			if _163_v0 {
				return true
			}
			return _163_v0
		})())
		var _208_v36 _dafny.Map
		_ = _208_v36
		_208_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v20, _165_v2)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index19
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index20
		var _rhs12 _dafny.Int = _165_v2
		_ = _rhs12
		var _rhs13 _dafny.Array = _205_v33
		_ = _rhs13
		var _rhs14 _dafny.Map = Companion_Default___.Fm1(false, _172_v7, _170_globalState)
		_ = _rhs14
		var _rhs15 bool = (func() bool {
			if (func() bool {
				if (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool) {
					return (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)
				}
				return _163_v0
			})() {
				return !(_208_v36).Equals(_208_v36)
			}
			return true
		})()
		_ = _rhs15
		var _rhs16 bool = (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)
		_ = _rhs16
		var _lhs8 _dafny.Array = _166_v3
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _lhs9
		var _lhs10 _dafny.Array = _166_v3
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _lhs11
		_165_v2 = _rhs12
		_205_v33 = _rhs13
		_207_v35 = _rhs14
		(_lhs8).ArraySet1(_rhs15, (_lhs9).Int())
		(_lhs10).ArraySet1(_rhs16, (_lhs11).Int())
		_165_v2 = (_dafny.Zero).Minus((_204_i5).Minus(_204_i5))
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index21
		(_166_v3).ArraySet1((_165_v2).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mjidffht")).Cardinality())) != 0, (_index21).Int())
		_163_v0 = ((_dafny.Zero).Minus(_204_i5)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_169_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-249))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_209_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_210_i6 _dafny.Int) _dafny.CodePoint {
				return _209_v10
			}
		})(_175_v10))))).Cardinality())) == 0
	}
	var _211_v37 _dafny.Set
	_ = _211_v37
	_211_v37 = _dafny.SetOf(_175_v10)
	var _212_v38 _dafny.MultiSet
	_ = _212_v38
	_212_v38 = _dafny.MultiSetOf(_211_v37)
	var _213_v39 _dafny.MultiSet
	_ = _213_v39
	_213_v39 = _dafny.MultiSetOf(false, (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool), (_164_v1).Select((Companion_Default___.SafeIndex((_174_v9).Cardinality(), _dafny.IntOfUint32((_164_v1).Cardinality()))).Uint32()).(bool))
	var _214_v41 _dafny.Sequence
	_ = _214_v41
	_214_v41 = _dafny.SeqOf(_211_v37, func() _dafny.Set {
		var _coll21 = _dafny.NewBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))).Keys().Elements()); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _215_v40 _dafny.CodePoint
			_215_v40 = interface{}(_compr_21).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))).Contains(_215_v40) {
				_coll21.Add(_215_v40)
			}
		}
		return _coll21.ToSet()
	}())
	var _216_v42 _dafny.Array
	_ = _216_v42
	var _nwElement0_2 _dafny.MultiSet = _212_v38
	_ = _nwElement0_2
	var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(9))
	_ = _nw15
	(_nw15).ArraySet1(_nwElement0_2, 0)
	(_nw15).ArraySet1(_212_v38, 1)
	(_nw15).ArraySet1(Companion_Default___.Fm2(!(false), _165_v2, _170_globalState), 2)
	(_nw15).ArraySet1(_dafny.MultiSetFromSeq(Companion_Default___.Fm3((_213_v39).Cardinality(), _165_v2, _dafny.SeqOf(_174_v9, _174_v9), _175_v10, _170_globalState)), 3)
	(_nw15).ArraySet1(_212_v38, 4)
	(_nw15).ArraySet1(_212_v38, 5)
	(_nw15).ArraySet1(_dafny.MultiSetFromSeq(_214_v41), 6)
	(_nw15).ArraySet1((_dafny.MultiSetOf(_211_v37, _211_v37)).Update(_211_v37, Companion_Default___.Abs(_165_v2)), 7)
	(_nw15).ArraySet1(_212_v38, 8)
	_216_v42 = _nw15
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_216_v42), 0))
	_ = _index22
	(_216_v42).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(389))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}((func(_217_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.Set {
		return func(_218_i7 _dafny.Int) _dafny.Set {
			return func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v10, _dafny.IntOfInt64(613))).Keys().Elements()); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _219_v43 _dafny.CodePoint
					_219_v43 = interface{}(_compr_22).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v10, _dafny.IntOfInt64(613))).Contains(_219_v43) {
						_coll22.Add(_219_v43)
					}
				}
				return _coll22.ToSet()
			}()
		}
	})(_175_v10)))), (_index22).Int())
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_216_v42), 0))
	_ = _index23
	(_216_v42).ArraySet1((_212_v38).Update(_211_v37, Companion_Default___.Abs(_165_v2)), (_index23).Int())
	var _hi1 _dafny.Int = (_165_v2).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v2, _165_v2)).Cardinality())
	_ = _hi1
	for _220_i8 := _165_v2; _220_i8.Cmp(_hi1) < 0; _220_i8 = _220_i8.Plus(_dafny.One) {
		_188_v20 = _188_v20
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index24
		(_166_v3).ArraySet1(!(!((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))), (_index24).Int())
		_165_v2 = _220_i8
		if (_165_v2).Cmp(_dafny.IntOfInt64(452)) <= 0 {
			_165_v2 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_164_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_169_v6).Cardinality()), _dafny.IntOfUint32((_164_v1).Cardinality()))).Uint32(), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)), _dafny.SeqOf(!(_163_v0), _163_v0, _163_v0, _163_v0))).Cardinality())).Plus(_220_i8)
			var _221_v44 _dafny.Map
			_ = _221_v44
			_221_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_165_v2), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
			var _222_v46 *C2
			_ = _222_v46
			var _nw16 *C2 = New_C2_()
			_ = _nw16
			_nw16.Ctor__(!(_221_v44).Equals(func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(603), _dafny.IntOfInt64(37))); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _223_v45 _dafny.Int
					_223_v45 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(603)).Cmp(_223_v45) <= 0) && ((_223_v45).Cmp(_dafny.IntOfInt64(37)) < 0) {
						_coll23.Add((_223_v45).Plus(_165_v2), _163_v0)
					}
				}
				return _coll23.ToMap()
			}()))
			_222_v46 = _nw16
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
			_ = _index25
			(_166_v3).ArraySet1(!(_163_v0), (_index25).Int())
			_163_v0 = (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)
			var _224_v47 _dafny.Array
			_ = _224_v47
			var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw17
			_224_v47 = _nw17
			var _225_v48 _dafny.Array
			_ = _225_v48
			var _nw18 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
			_ = _nw18
			_225_v48 = _nw18
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_224_v47), 0))
			_ = _index26
			(_224_v47).ArraySet1(_225_v48, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_224_v47), 0))
			_ = _index27
			var _rhs17 bool = ((_165_v2).Minus(_dafny.IntOfInt64(957))).Cmp(_dafny.IntOfInt64(855)) <= 0
			_ = _rhs17
			var _rhs18 _dafny.Int = _220_i8
			_ = _rhs18
			var _rhs19 _dafny.Int = (func() _dafny.Int {
				if (_173_v8).Contains(((_221_v44).Merge(_221_v44)).Cardinality()) {
					return (_173_v8).Get(((_221_v44).Merge(_221_v44)).Cardinality()).(_dafny.Int)
				}
				return _220_i8
			})()
			_ = _rhs19
			var _rhs20 _dafny.Array = _225_v48
			_ = _rhs20
			var _lhs12 _dafny.Array = _224_v47
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_224_v47), 0))
			_ = _lhs13
			_163_v0 = _rhs17
			_165_v2 = _rhs18
			_165_v2 = _rhs19
			(_lhs12).ArraySet1(_rhs20, (_lhs13).Int())
		} else {
			_169_v6 = _dafny.UnicodeSeqOfUtf8Bytes("tkuqi")
			var _226_v49 D2
			_ = _226_v49
			_226_v49 = Companion_D2_.Create_DC8_(_163_v0, _163_v0, _163_v0)
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
			_ = _index28
			var _rhs21 bool = _163_v0
			_ = _rhs21
			var _rhs22 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_220_i8).Times(_dafny.IntOfInt64(479)), _165_v2)))
			_ = _rhs22
			var _rhs23 bool = (_226_v49).Dtor_cf17()
			_ = _rhs23
			var _lhs14 _dafny.Array = _166_v3
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
			_ = _lhs15
			(_lhs14).ArraySet1(_rhs21, (_lhs15).Int())
			_165_v2 = _rhs22
			_163_v0 = _rhs23
			_175_v10 = _dafny.CodePoint('l')
			var _227_v50 _dafny.Array
			_ = _227_v50
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
			_ = _nw19
			_227_v50 = _nw19
			var _228_v51 _dafny.Set
			_ = _228_v51
			_228_v51 = _dafny.SetOf(_165_v2, _165_v2)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_227_v50), 0))
			_ = _index29
			(_227_v50).ArraySet1(_228_v51, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_227_v50), 0))
			_ = _index30
			(_227_v50).ArraySet1(_228_v51, (_index30).Int())
			var _229_v52 T1
			_ = _229_v52
			var _nw20 *C5 = New_C5_()
			_ = _nw20
			_nw20.Ctor__(!(_163_v0), _188_v20)
			_229_v52 = _nw20
			_229_v52 = _229_v52
		}
	}
	var _230_i9 _dafny.Int
	_ = _230_i9
	_230_i9 = _dafny.Zero
	{
		for ((func() _dafny.Int {
			if (_213_v39).Contains(_163_v0) {
				return (_213_v39).Multiplicity(_163_v0)
			}
			return _dafny.IntOfInt64(-671)
		})()).Cmp(_165_v2) != 0 {
			{
				if (_230_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_230_i9 = (_230_i9).Plus(_dafny.One)
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_188_v20), 0))
				_ = _index31
				(_188_v20).ArraySet1(_165_v2, (_index31).Int())
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_188_v20), 0))
				_ = _index32
				(_188_v20).ArraySet1((func() _dafny.Int {
					if (_203_v32).Contains((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)) {
						return (_203_v32).Get((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)).(_dafny.Int)
					}
					return _165_v2
				})(), (_index32).Int())
				var _231_v53 _dafny.Map
				_ = _231_v53
				_231_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v6, (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
				var _232_v54 _dafny.Int
				_ = _232_v54
				var _out1 _dafny.Int
				_ = _out1
				_out1 = Companion_Default___.M0(_231_v53, _170_globalState)
				_232_v54 = _out1
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
				_ = _index33
				(_166_v3).ArraySet1(_163_v0, (_index33).Int())
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_188_v20), 0))
				_ = _index34
				var _rhs24 _dafny.Int = _dafny.IntOfInt64(711)
				_ = _rhs24
				var _rhs25 _dafny.Int = (_232_v54).Plus(_dafny.IntOfInt64(-301))
				_ = _rhs25
				var _lhs16 _dafny.Array = _188_v20
				_ = _lhs16
				var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_188_v20), 0))
				_ = _lhs17
				_165_v2 = _rhs24
				(_lhs16).ArraySet1(_rhs25, (_lhs17).Int())
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _233_v55 *C3
	_ = _233_v55
	var _nw21 *C3 = New_C3_()
	_ = _nw21
	_nw21.Ctor__(_163_v0, !((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)))
	_233_v55 = _nw21
	_233_v55 = (func() *C3 {
		if !_dafny.Companion_Sequence_.Contains(_164_v1, (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)) {
			return _233_v55
		}
		return _233_v55
	})()
	var _234_v56 _dafny.Sequence
	_ = _234_v56
	_234_v56 = _dafny.SeqOf(_188_v20, _188_v20, (func() _dafny.Array {
		if (_233_v55).F19() {
			return _188_v20
		}
		return _188_v20
	})())
	_234_v56 = _234_v56
	for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_188_v20), 0))); ; {
		_guard_loop_0, _ok24 := _iter24()
		if !_ok24 {
			break
		}
		var _235_i10 _dafny.Int
		_235_i10 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_235_i10).Sign() != -1) && ((_235_i10).Cmp(_dafny.ArrayLen((_188_v20), 0)) < 0)) {
			(_188_v20).ArraySet1(Companion_Default___.SafeDivisionInt(_235_i10, _165_v2), (_235_i10).Int())
		}
	}
	if _163_v0 {
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index35
		(_166_v3).ArraySet1(true, (_index35).Int())
		var _236_v57 _dafny.Array
		_ = _236_v57
		var _out2 _dafny.Array
		_ = _out2
		_out2 = (_233_v55).M1(Companion_Default___.Fm31(_170_globalState), _170_globalState)
		_236_v57 = _out2
		var _237_v58 _dafny.Array
		_ = _237_v58
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_2
		var _nw22 _dafny.Array
		_ = _nw22
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw22 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Set = (func(_238_v55 *C3) func(_dafny.Int) _dafny.Set {
				return func(_239_i11 _dafny.Int) _dafny.Set {
					return _dafny.SetOf((_238_v55).F19())
				}
			})(_233_v55)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw22 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw22).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw22).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_237_v58 = _nw22
		var _240_v60 _dafny.Set
		_ = _240_v60
		_240_v60 = _dafny.SetOf((_233_v55).F19(), (_233_v55).F19(), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool), !(Companion_Default___.Fm0(func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter25 := _dafny.Iterate((_174_v9).Elements()); ; {
				_compr_24, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _241_v59 _dafny.Int
				_241_v59 = interface{}(_compr_24).(_dafny.Int)
				if (_174_v9).Contains(_241_v59) {
					_coll24.Add((_241_v59).Plus(_165_v2), _165_v2)
				}
			}
			return _coll24.ToMap()
		}(), _165_v2, _170_globalState)))
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_237_v58), 0))
		_ = _index36
		(_237_v58).ArraySet1((func() _dafny.Set {
			if (_164_v1).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_164_v1).Cardinality()))).Uint32()).(bool) {
				return _240_v60
			}
			return _dafny.SetOf((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
		})(), (_index36).Int())
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_237_v58), 0))
		_ = _index37
		(_237_v58).ArraySet1((_240_v60).Union(_240_v60), (_index37).Int())
		var _242_v61 *C3
		_ = _242_v61
		var _nw23 *C3 = New_C3_()
		_ = _nw23
		_nw23.Ctor__((_233_v55).F19(), ((_dafny.SetOf(_165_v2)).Cardinality()).Cmp(_165_v2) > 0)
		_242_v61 = _nw23
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index38
		(_166_v3).ArraySet1((Companion_Default___.SafeModuloInt(_165_v2, _165_v2)).Cmp(((_240_v60).Difference((_237_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_237_v58), 0))).Int()).(_dafny.Set))).Cardinality()) < 0, (_index38).Int())
	} else {
		var _243_v63 _dafny.Sequence
		_ = _243_v63
		_243_v63 = _dafny.SeqOf(_165_v2, _165_v2)
		var _pat_let_tv6 = _165_v2
		_ = _pat_let_tv6
		var _244_v64 _dafny.Map
		_ = _244_v64
		_244_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_233_v55).Fm4(_163_v0, func(_pat_let0_0 D0) D0 {
			return func(_245_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 _dafny.Int) D0 {
					return func(_246_dt__update_hcf2_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_((_245_dt__update__tmp_h0).Dtor_cf1(), _246_dt__update_hcf2_h0)
					}(_pat_let1_0)
				}(_pat_let_tv6)
			}(_pat_let0_0)
		}(_172_v7), (_203_v32).Cardinality(), func() _dafny.Map {
			var _coll25 = _dafny.NewMapBuilder()
			_ = _coll25
			for _iter26 := _dafny.Iterate((_243_v63).Elements()); ; {
				_compr_25, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _247_v62 _dafny.Int
				_247_v62 = interface{}(_compr_25).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_243_v63, _247_v62) {
					_coll25.Add((_247_v62).Minus(_165_v2), _165_v2)
				}
			}
			return _coll25.ToMap()
		}(), _170_globalState), _203_v32)
		_165_v2 = (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_233_v55).F19(), Companion_Default___.Fm31(_170_globalState))).Merge((func() _dafny.Map {
			if (_244_v64).Contains((_233_v55).F19()) {
				return (_244_v64).Get((_233_v55).F19()).(_dafny.Map)
			}
			return _203_v32
		})())).Cardinality())
		var _248_v65 _dafny.Array
		_ = _248_v65
		var _nw24 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw24
		_248_v65 = _nw24
		_248_v65 = _248_v65
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index39
		(_166_v3).ArraySet1((_233_v55).F19(), (_index39).Int())
		_165_v2 = _165_v2
		_174_v9 = _174_v9
	}
	var _249_v66 _dafny.Array
	_ = _249_v66
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
	_ = _nw25
	_249_v66 = _nw25
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))
	_ = _index40
	(_249_v66).ArraySet1(_188_v20, (_index40).Int())
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))
	_ = _index41
	(_249_v66).ArraySet1(_188_v20, (_index41).Int())
	for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_166_v3), 0))); ; {
		_guard_loop_1, _ok27 := _iter27()
		if !_ok27 {
			break
		}
		var _250_i12 _dafny.Int
		_250_i12 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_250_i12).Sign() != -1) && ((_250_i12).Cmp(_dafny.ArrayLen((_166_v3), 0)) < 0)) {
			(_166_v3).ArraySet1(!(_163_v0) || ((func() bool {
				if (_233_v55).F19() {
					return (_233_v55).F19()
				}
				return (_233_v55).F19()
			})()), (_250_i12).Int())
		}
	}
	var _251_v67 D5
	_ = _251_v67
	_251_v67 = Companion_D5_.Create_DC13_()
	var _source7 D5 = _251_v67
	_ = _source7
	if _source7.Is_DC13() {
		var _252_v68 _dafny.Set
		_ = _252_v68
		_252_v68 = _dafny.SetOf(_165_v2)
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index42
		var _rhs26 _dafny.Int = (_165_v2).Minus(_165_v2)
		_ = _rhs26
		var _rhs27 bool = (_252_v68).Equals(_dafny.SetOf(_165_v2, _dafny.IntOfUint32((_169_v6).Cardinality())))
		_ = _rhs27
		var _lhs18 _dafny.Array = _166_v3
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _lhs19
		_165_v2 = _rhs26
		(_lhs18).ArraySet1(_rhs27, (_lhs19).Int())
		var _253_v69 _dafny.Array
		_ = _253_v69
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(8))
		_ = _nw26
		_253_v69 = _nw26
		var _254_v70 _dafny.Map
		_ = _254_v70
		_254_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_233_v55, _173_v8)
		var _255_v71 _dafny.Sequence
		_ = _255_v71
		_255_v71 = _dafny.SeqOf(_173_v8)
		var _pat_let_tv7 = _165_v2
		_ = _pat_let_tv7
		var _pat_let_tv8 = _165_v2
		_ = _pat_let_tv8
		var _256_v72 _dafny.Array
		_ = _256_v72
		var _nwElement0_3 _dafny.Map = (func() _dafny.Map {
			if (_254_v70).Contains(_233_v55) {
				return (_254_v70).Get(_233_v55).(_dafny.Map)
			}
			return _173_v8
		})()
		_ = _nwElement0_3
		var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(23))
		_ = _nw27
		(_nw27).ArraySet1(_nwElement0_3, 0)
		(_nw27).ArraySet1(_173_v8, 1)
		(_nw27).ArraySet1(_173_v8, 2)
		(_nw27).ArraySet1(_173_v8, 3)
		(_nw27).ArraySet1((_173_v8).Update(_165_v2, _165_v2), 4)
		(_nw27).ArraySet1((func() _dafny.Map {
			if (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool) {
				return (_255_v71).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_255_v71).Cardinality()))).Uint32()).(_dafny.Map)
			}
			return _173_v8
		})(), 5)
		(_nw27).ArraySet1((func() _dafny.Map {
			if _163_v0 {
				return _173_v8
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(303), _dafny.IntOfInt64(904))
		})(), 6)
		(_nw27).ArraySet1(((func(_pat_let2_0 D1) D1 {
			return func(_257_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let3_0 _dafny.Map) D1 {
					return func(_258_dt__update_hcf3_h0 _dafny.Map) D1 {
						return Companion_D1_.Create_DC2_(_258_dt__update_hcf3_h0)
					}(_pat_let3_0)
				}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv7, _pat_let_tv8))
			}(_pat_let2_0)
		}(Companion_D1_.Create_DC2_(_173_v8))).Dtor_cf3()).Merge(_173_v8), 7)
		(_nw27).ArraySet1(_173_v8, 8)
		(_nw27).ArraySet1(Companion_Default___.Fm39(_165_v2, _170_globalState), 9)
		(_nw27).ArraySet1(((_173_v8).Update(_dafny.IntOfUint32((Companion_Default___.Fm28(_dafny.IntOfInt64(203), _165_v2, !(true), _170_globalState)).Cardinality()), _165_v2)).Merge(_173_v8), 10)
		(_nw27).ArraySet1(_173_v8, 11)
		(_nw27).ArraySet1(_173_v8, 12)
		(_nw27).ArraySet1(_173_v8, 13)
		(_nw27).ArraySet1(_173_v8, 14)
		(_nw27).ArraySet1(_173_v8, 15)
		(_nw27).ArraySet1(_173_v8, 16)
		(_nw27).ArraySet1(_173_v8, 17)
		(_nw27).ArraySet1((func() _dafny.Map {
			if _163_v0 {
				return _173_v8
			}
			return _173_v8
		})(), 18)
		(_nw27).ArraySet1(Companion_Default___.Fm39(_165_v2, _170_globalState), 19)
		(_nw27).ArraySet1(_173_v8, 20)
		(_nw27).ArraySet1(_173_v8, 21)
		(_nw27).ArraySet1(_173_v8, 22)
		_256_v72 = _nw27
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_256_v72), 0))
		_ = _index43
		(_256_v72).ArraySet1(_173_v8, (_index43).Int())
		var _259_v73 _dafny.Sequence
		_ = _259_v73
		_259_v73 = _dafny.SeqOf(_165_v2, _165_v2)
		var _260_v74 _dafny.Set
		_ = _260_v74
		_260_v74 = _dafny.SetOf(_259_v73, _259_v73, _259_v73, _259_v73)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))
		_ = _index44
		(_188_v20).ArraySet1((_260_v74).Cardinality(), (_index44).Int())
		var _261_v75 _dafny.Map
		_ = _261_v75
		_261_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
		var _262_v76 D4
		_ = _262_v76
		_262_v76 = Companion_D4_.Create_DC11_((_233_v55).F19())
		var _263_v77 _dafny.Set
		_ = _263_v77
		_263_v77 = _dafny.SetOf((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool), !((_262_v76).Dtor_cf21()))
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_256_v72), 0))
		_ = _index45
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))
		_ = _index46
		var _rhs28 _dafny.Array = (func() _dafny.Array {
			if !(_261_v75).Contains((_164_v1).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_164_v1).Cardinality()))).Uint32()).(bool)) {
				return _253_v69
			}
			return _253_v69
		})()
		_ = _rhs28
		var _rhs29 _dafny.Map = (_173_v8).Update((_174_v9).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_233_v55).F19() {
				return _259_v73
			}
			return _259_v73
		})()).Cardinality()))
		_ = _rhs29
		var _rhs30 _dafny.Int = (_dafny.IntOfInt64(-208)).Minus(((_263_v77).Union(_dafny.SetOf((_233_v55).F19()))).Cardinality())
		_ = _rhs30
		var _lhs20 _dafny.Array = _256_v72
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_256_v72), 0))
		_ = _lhs21
		var _lhs22 _dafny.Array = _188_v20
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))
		_ = _lhs23
		_253_v69 = _rhs28
		(_lhs20).ArraySet1(_rhs29, (_lhs21).Int())
		(_lhs22).ArraySet1(_rhs30, (_lhs23).Int())
		if (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool) {
			var _264_v78 _dafny.Sequence
			_ = _264_v78
			_264_v78 = _dafny.SeqOf(_259_v73, (func() _dafny.Sequence {
				if (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool) {
					return _259_v73
				}
				return _259_v73
			})(), _259_v73, _259_v73, _259_v73)
			_264_v78 = _264_v78
			var _265_v79 *C3
			_ = _265_v79
			var _nw28 *C3 = New_C3_()
			_ = _nw28
			_nw28.Ctor__(!((_233_v55).F19()) || ((_233_v55).F19()), (func() bool {
				if _163_v0 {
					return (_233_v55).F19()
				}
				return (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)
			})())
			_265_v79 = _nw28
			var _266_v80 _dafny.Array
			_ = _266_v80
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
			_ = _nw29
			_266_v80 = _nw29
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
			_ = _nw30
			_266_v80 = _nw30
			_166_v3 = _166_v3
			var _267_v81 *C3
			_ = _267_v81
			var _nw31 *C3 = New_C3_()
			_ = _nw31
			_nw31.Ctor__((_233_v55).F19(), (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool))
			_267_v81 = _nw31
		} else {
			var _268_v82 _dafny.Array
			_ = _268_v82
			var _out3 _dafny.Array
			_ = _out3
			_out3 = (_233_v55).M1((_188_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))).Int()).(_dafny.Int), _170_globalState)
			_268_v82 = _out3
			var _269_v83 D1
			_ = _269_v83
			_269_v83 = Companion_D1_.Create_DC5_(_169_v6, _166_v3, _165_v2)
			var _270_v84 _dafny.Map
			_ = _270_v84
			var _271_v85 bool
			_ = _271_v85
			var _272_v86 bool
			_ = _272_v86
			var _273_v87 _dafny.Sequence
			_ = _273_v87
			var _out4 _dafny.Map
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Sequence
			_ = _out7
			_out4, _out5, _out6, _out7 = (_233_v55).M2((_188_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Update(_169_v6, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-928), _dafny.IntOfUint32((_169_v6).Cardinality()))).Uint32(), _175_v10), (_269_v83).Dtor_cf10(), _170_globalState)
			_270_v84 = _out4
			_271_v85 = _out5
			_272_v86 = _out6
			_273_v87 = _out7
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))
			_ = _index47
			(_188_v20).ArraySet1(Companion_Default___.SafeDivisionInt(((_259_v73).Select((Companion_Default___.SafeIndex((_188_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_259_v73).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_165_v2), (_dafny.Zero).Minus((_188_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))).Int()).(_dafny.Int))), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))
			_ = _index48
			(_188_v20).ArraySet1((Companion_Default___.Fm31(_170_globalState)).Times((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(((_174_v9).Update(_165_v2, Companion_Default___.Abs((_188_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_188_v20), 0))).Int()).(_dafny.Int)))).Cardinality(), _165_v2))), (_index48).Int())
			_166_v3 = _166_v3
		}
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index49
		(_166_v3).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("lv"), _169_v6), (_index49).Int())
	} else if _source7.Is_DC14() {
		var _274___mcc_h0 bool = _source7.Get_().(D5_DC14).Cf23
		_ = _274___mcc_h0
		var _275_cf23 bool = _274___mcc_h0
		_ = _275_cf23
		var _276_v88 _dafny.Array
		_ = _276_v88
		var _nwElement0_4 _dafny.Array = _166_v3
		_ = _nwElement0_4
		var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(3))
		_ = _nw32
		(_nw32).ArraySet1(_nwElement0_4, 0)
		(_nw32).ArraySet1(_166_v3, 1)
		(_nw32).ArraySet1(_166_v3, 2)
		_276_v88 = _nw32
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_276_v88), 0))
		_ = _index50
		(_276_v88).ArraySet1((func() _dafny.Array {
			if _275_cf23 {
				return _166_v3
			}
			return _166_v3
		})(), (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_276_v88), 0))
		_ = _index51
		(_276_v88).ArraySet1(_166_v3, (_index51).Int())
		var _277_v89 _dafny.Sequence
		_ = _277_v89
		_277_v89 = _dafny.SeqOf(_165_v2)
		var _278_v90 _dafny.Sequence
		_ = _278_v90
		_278_v90 = _dafny.SeqOf(_166_v3, _166_v3)
		var _279_v91 _dafny.Map
		_ = _279_v91
		_279_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_277_v89, Companion_Default___.Fm16(_dafny.IntOfUint32((_278_v90).Cardinality()), _165_v2, (_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool), _165_v2, _170_globalState)), _dafny.IntOfUint32((_169_v6).Cardinality()))
		_279_v91 = (_279_v91).Update(_277_v89, Companion_Default___.SafeModuloInt(_165_v2, _165_v2))
		var _280_v92 _dafny.MultiSet
		_ = _280_v92
		_280_v92 = _dafny.MultiSetOf(_192_v22, _192_v22, _192_v22)
		var _pat_let_tv9 = _173_v8
		_ = _pat_let_tv9
		_169_v6 = Companion_Default___.Fm17((_dafny.MultiSetOf(_192_v22, _192_v22, func(_pat_let4_0 D1) D1 {
			return func(_281_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let5_0 _dafny.Map) D1 {
					return func(_282_dt__update_hcf3_h1 _dafny.Map) D1 {
						return Companion_D1_.Create_DC2_(_282_dt__update_hcf3_h1)
					}(_pat_let5_0)
				}(_pat_let_tv9)
			}(_pat_let4_0)
		}(_192_v22))).Intersection(_280_v92), _170_globalState)
		_165_v2 = (_dafny.IntOfInt64(297)).Minus(_165_v2)
	} else if _source7.Is_DC12() {
		var _283___mcc_h1 _dafny.Map = _source7.Get_().(D5_DC12).Cf22
		_ = _283___mcc_h1
		var _284_cf22 _dafny.Map = _283___mcc_h1
		_ = _284_cf22
		var _285_v93 *C2
		_ = _285_v93
		var _nw33 *C2 = New_C2_()
		_ = _nw33
		_nw33.Ctor__(!(!(_163_v0)))
		_285_v93 = _nw33
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))
		_ = _index52
		(_249_v66).ArraySet1(_188_v20, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))
		_ = _index53
		(_249_v66).ArraySet1(_188_v20, (_index53).Int())
		var _286_v94 *C5
		_ = _286_v94
		var _nw34 *C5 = New_C5_()
		_ = _nw34
		_nw34.Ctor__(!(true) || (_163_v0), _188_v20)
		_286_v94 = _nw34
	} else {
		var _287___mcc_h2 D5 = _source7.Get_().(D5_DC15).Cf24
		_ = _287___mcc_h2
		var _288_cf24 D5 = _287___mcc_h2
		_ = _288_cf24
		if (_233_v55).F19() {
			_163_v0 = (_233_v55).F19()
			_163_v0 = (_233_v55).F19()
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))
			_ = _index54
			(_249_v66).ArraySet1(_188_v20, (_index54).Int())
			var _289_v95 _dafny.Array
			_ = _289_v95
			var _out8 _dafny.Array
			_ = _out8
			_out8 = (_233_v55).M1(_165_v2, _170_globalState)
			_289_v95 = _out8
			_163_v0 = (_233_v55).F19()
		} else {
			var _290_v96 _dafny.Map
			_ = _290_v96
			var _291_v97 bool
			_ = _291_v97
			var _292_v98 bool
			_ = _292_v98
			var _293_v99 _dafny.Sequence
			_ = _293_v99
			var _out9 _dafny.Map
			_ = _out9
			var _out10 bool
			_ = _out10
			var _out11 bool
			_ = _out11
			var _out12 _dafny.Sequence
			_ = _out12
			_out9, _out10, _out11, _out12 = (_233_v55).M2(_165_v2, _169_v6, (_dafny.Zero).Minus(_165_v2), _170_globalState)
			_290_v96 = _out9
			_291_v97 = _out10
			_292_v98 = _out11
			_293_v99 = _out12
			var _294_v100 T0
			_ = _294_v100
			var _nw35 *C5 = New_C5_()
			_ = _nw35
			_nw35.Ctor__(_291_v97, _188_v20)
			_294_v100 = _nw35
			var _295_v101 _dafny.Map
			_ = _295_v101
			_295_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v2, _294_v100)
			_292_v98 = !(!((((_dafny.Zero).Minus(_165_v2)).Times((_295_v101).Cardinality())).Cmp(_165_v2) >= 0))
			var _296_v102 _dafny.Sequence
			_ = _296_v102
			_296_v102 = _dafny.SeqOf(_166_v3)
			var _297_v103 D1
			_ = _297_v103
			_297_v103 = Companion_D1_.Create_DC4_((_296_v102).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_296_v102).Cardinality()))).Uint32()).(_dafny.Array), _165_v2, _175_v10)
			_292_v98 = ((_297_v103).Dtor_cf6()).Cmp(_165_v2) < 0
			_165_v2 = Companion_Default___.SafeDivisionInt((_165_v2).Plus((_dafny.Zero).Minus(_165_v2)), _165_v2)
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_188_v20), 0))
			_ = _index55
			(_188_v20).ArraySet1((_165_v2).Minus(_165_v2), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_188_v20), 0))
			_ = _index56
			(_188_v20).ArraySet1(_165_v2, (_index56).Int())
		}
		if _163_v0 {
			var _298_v104 _dafny.Sequence
			_ = _298_v104
			_298_v104 = _dafny.SeqOf(_165_v2, _165_v2, (_174_v9).Cardinality())
			var _299_v105 _dafny.Sequence
			_ = _299_v105
			_299_v105 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("epfsk")).Cardinality()), _165_v2), _298_v104, _298_v104)
			var _300_v106 _dafny.Map
			_ = _300_v106
			var _301_v107 bool
			_ = _301_v107
			var _302_v108 bool
			_ = _302_v108
			var _303_v109 _dafny.Sequence
			_ = _303_v109
			var _out13 _dafny.Map
			_ = _out13
			var _out14 bool
			_ = _out14
			var _out15 bool
			_ = _out15
			var _out16 _dafny.Sequence
			_ = _out16
			_out13, _out14, _out15, _out16 = (_233_v55).M2(((_dafny.SetOf((_233_v55).F19(), (_233_v55).F19())).Cardinality()).Plus(_dafny.IntOfUint32(((_299_v105).Select((Companion_Default___.SafeIndex(_165_v2, _dafny.IntOfUint32((_299_v105).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), _169_v6, _165_v2, _170_globalState)
			_300_v106 = _out13
			_301_v107 = _out14
			_302_v108 = _out15
			_303_v109 = _out16
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))
			_ = _arr0
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))
			_ = _index57
			(_arr0).ArraySet1(Companion_Default___.SafeModuloInt(_165_v2, _dafny.IntOfUint32((_303_v109).Cardinality())), (_index57).Int())
			var _304_v110 _dafny.Array
			_ = _304_v110
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_3
			var _nw36 _dafny.Array
			_ = _nw36
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw36 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Set = (func(_305_v108 bool, _306_v107 bool, _307_v3 _dafny.Array, _308_v0 bool, _309_v55 *C3) func(_dafny.Int) _dafny.Set {
					return func(_310_i13 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v108, _306_v107), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_307_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_307_v3), 0))).Int()).(bool), _308_v0))).Union(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _305_v108), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v107, (_309_v55).F19())))
					}
				})(_302_v108, _301_v107, _166_v3, _163_v0, _233_v55)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw36 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw36).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw36).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_304_v110 = _nw36
			var _311_v111 _dafny.Map
			_ = _311_v111
			_311_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v0, !(_301_v107))
			var _312_v112 _dafny.Set
			_ = _312_v112
			_312_v112 = _dafny.SetOf(_311_v111, _311_v111)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_304_v110), 0))
			_ = _index58
			(_304_v110).ArraySet1(_312_v112, (_index58).Int())
			var _313_v113 _dafny.Set
			_ = _313_v113
			_313_v113 = _dafny.SetOf(_165_v2)
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))
			_ = _arr1
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))
			_ = _index59
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_304_v110), 0))
			_ = _index60
			var _rhs31 _dafny.Int = (_165_v2).Plus((_313_v113).Cardinality())
			_ = _rhs31
			var _rhs32 _dafny.Int = _165_v2
			_ = _rhs32
			var _rhs33 _dafny.MultiSet = _213_v39
			_ = _rhs33
			var _rhs34 _dafny.Int = _165_v2
			_ = _rhs34
			var _rhs35 _dafny.Set = _312_v112
			_ = _rhs35
			var _lhs24 _dafny.Array = _dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _304_v110
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_304_v110), 0))
			_ = _lhs27
			_165_v2 = _rhs31
			_165_v2 = _rhs32
			_213_v39 = _rhs33
			(_lhs24).ArraySet1(_rhs34, (_lhs25).Int())
			(_lhs26).ArraySet1(_rhs35, (_lhs27).Int())
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))
			_ = _arr2
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))
			_ = _index61
			var _rhs36 _dafny.Int = (_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))).Int()).(_dafny.Int)
			_ = _rhs36
			var _rhs37 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))).Int()).(_dafny.Int), _165_v2)
			_ = _rhs37
			var _lhs28 _dafny.Array = _dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))
			_ = _lhs29
			(_lhs28).ArraySet1(_rhs36, (_lhs29).Int())
			_173_v8 = _rhs37
			var _314_v114 _dafny.Map
			_ = _314_v114
			_314_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _298_v104)
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))
			_ = _arr3
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_dafny.ArrayCastTo((_249_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))).Int()))), 0))
			_ = _index62
			(_arr3).ArraySet1(Companion_Default___.SafeModuloInt((_314_v114).Cardinality(), Companion_Default___.Fm13(_302_v108, (_233_v55).F19(), _175_v10, _170_globalState)), (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
			_ = _index63
			(_166_v3).ArraySet1(((_233_v55).F19()) == ((_233_v55).F19()), (_index63).Int())
		} else {
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
			_ = _index64
			(_166_v3).ArraySet1(false, (_index64).Int())
			var _315_v116 _dafny.Sequence
			_ = _315_v116
			_315_v116 = _dafny.SeqOf(_dafny.IntOfUint32((_169_v6).Cardinality()), _165_v2)
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
			_ = _index65
			(_166_v3).ArraySet1(Companion_Default___.Fm0(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter28 := _dafny.Iterate((_315_v116).Elements()); ; {
					_compr_26, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _316_v115 _dafny.Int
					_316_v115 = interface{}(_compr_26).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_315_v116, _316_v115) {
						_coll26.Add((_316_v115).Plus(_165_v2), ((_174_v9).Update(_dafny.IntOfInt64(272), Companion_Default___.Abs(_165_v2))).Cardinality())
					}
				}
				return _coll26.ToMap()
			}(), _165_v2, _170_globalState), (_index65).Int())
			var _317_v117 D2
			_ = _317_v117
			_317_v117 = Companion_D2_.Create_DC6_(_315_v116)
			var _318_v118 D2
			_ = _318_v118
			_318_v118 = Companion_D2_.Create_DC6_((_317_v117).Dtor_cf11())
			var _319_v119 _dafny.Map
			_ = _319_v119
			_319_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_165_v2), _169_v6)
			var _320_v120 _dafny.Map
			_ = _320_v120
			_320_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_233_v55).F19(), _319_v119)
			var _rhs38 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("od")).Cardinality())
			_ = _rhs38
			var _rhs39 D2 = Companion_Default___.Fm48(_175_v10, ((func() _dafny.Map {
				if (_320_v120).Contains((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)) {
					return (_320_v120).Get((_166_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))).Int()).(bool)).(_dafny.Map)
				}
				return _319_v119
			})()).Update(_165_v2, _169_v6), _165_v2, _dafny.Companion_Sequence_.IsProperPrefixOf(_169_v6, _169_v6), _170_globalState)
			_ = _rhs39
			var _rhs40 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(706))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_321_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_322_i14 _dafny.Int) _dafny.Int {
					return _321_v2
				}
			})(_165_v2)))
			_ = _rhs40
			var _rhs41 _dafny.Map = (_203_v32).Update((_233_v55).F19(), _dafny.IntOfInt64(488))
			_ = _rhs41
			_165_v2 = _rhs38
			_318_v118 = _rhs39
			_315_v116 = _rhs40
			_203_v32 = _rhs41
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_249_v66), 0))
			_ = _index66
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw37
			(_249_v66).ArraySet1(_nw37, (_index66).Int())
			var _323_v121 T1
			_ = _323_v121
			var _nw38 *C5 = New_C5_()
			_ = _nw38
			_nw38.Ctor__(true, _188_v20)
			_323_v121 = _nw38
			_323_v121 = _323_v121
		}
		var _324_v122 _dafny.Array
		_ = _324_v122
		var _out17 _dafny.Array
		_ = _out17
		_out17 = (_233_v55).M1((_dafny.IntOfInt64(85)).Times(_165_v2), _170_globalState)
		_324_v122 = _out17
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _index67
		var _rhs42 bool = _163_v0
		_ = _rhs42
		var _rhs43 _dafny.CodePoint = (_233_v55).Fm27(_170_globalState)
		_ = _rhs43
		var _lhs30 _dafny.Array = _166_v3
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_166_v3), 0))
		_ = _lhs31
		(_lhs30).ArraySet1(_rhs42, (_lhs31).Int())
		_175_v10 = _rhs43
	}
	_165_v2 = ((_dafny.Zero).Minus((_dafny.Zero).Minus(_165_v2))).Times(_165_v2)
	_dafny.Print(_163_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_164_v1, _dafny.SeqOf(false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_165_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v3).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v4).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v5).Equals(_dafny.SetOf(_dafny.SeqOf(false, true, false, true), _dafny.SeqOf(false, false, true, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_169_v6.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_170_globalState).F4(), _dafny.SeqOf(false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F5()).Equals(_dafny.SetOf(_dafny.SeqOf(false, true, false, true), _dafny.SeqOf(false, false, true, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F7()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-199))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_170_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F11().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v7).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v7).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(199), _dafny.IntOfInt64(833))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v9).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-199))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_v10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_176_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v20).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_191_v21).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_192_v22).Dtor_cf3()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(199), _dafny.IntOfInt64(833))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_193_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v32).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(78001))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_211_v37).Equals(_dafny.SetOf(_dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v38).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_213_v39).Equals(_dafny.MultiSetOf(false, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_214_v41, _dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('d')), _dafny.SetOf(_dafny.CodePoint('p')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('t'), _dafny.CodePoint('c'), _dafny.CodePoint('h')), _dafny.SetOf(_dafny.CodePoint('v')), _dafny.SetOf(_dafny.CodePoint('e'), _dafny.CodePoint('n')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('i')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')), _dafny.SetOf(_dafny.CodePoint('o')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')), _dafny.SetOf(_dafny.CodePoint('p')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_216_v42).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_230_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v55).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_234_v56).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_249_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
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

type D1_DC3 struct {
	Cf4 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Int) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf5 _dafny.Array
	Cf6 _dafny.Int
	Cf7 _dafny.CodePoint
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 _dafny.Array, Cf6 _dafny.Int, Cf7 _dafny.CodePoint) D1 {
	return D1{D1_DC4{Cf5, Cf6, Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf8  _dafny.Sequence
	Cf9  _dafny.Array
	Cf10 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 _dafny.Sequence, Cf9 _dafny.Array, Cf10 _dafny.Int) D1 {
	return D1{D1_DC5{Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC2 struct {
	Cf3 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf3 _dafny.Map) D1 {
	return D1{D1_DC2{Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.Zero)
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + data.Cf8.VerbatimString(true) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf3) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf8.Equals(data2.Cf8) && data1.Cf9 == data2.Cf9 && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf12 _dafny.Sequence
	Cf13 _dafny.Int
	Cf14 _dafny.Int
	Cf15 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.Sequence, Cf13 _dafny.Int, Cf14 _dafny.Int, Cf15 bool) D2 {
	return D2{D2_DC7{Cf12, Cf13, Cf14, Cf15}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf16 bool
	Cf17 bool
	Cf18 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf16 bool, Cf17 bool, Cf18 bool) D2 {
	return D2{D2_DC8{Cf16, Cf17, Cf18}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC6 struct {
	Cf11 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf11}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero, false)
}

func (_this D2) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf15() bool {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf16() bool {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf17() bool {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) Dtor_cf18() bool {
	return _this.Get_().(D2_DC8).Cf18
}

func (_this D2) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + data.Cf12.VerbatimString(true) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf12.Equals(data2.Cf12) && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf19 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf19 _dafny.Map) D3 {
	return D3{D3_DC9{Cf19}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf19
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf21 bool
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf21 bool) D4 {
	return D4{D4_DC11{Cf21}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf20 _dafny.Array
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf20 _dafny.Array) D4 {
	return D4{D4_DC10{Cf20}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(false)
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC11).Cf21
}

func (_this D4) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D4_DC10).Cf20
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && data1.Cf21 == data2.Cf21
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf20 == data2.Cf20
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
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_() D5 {
	return D5{D5_DC13{}}
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
	Cf22 _dafny.Map
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf22 _dafny.Map) D5 {
	return D5{D5_DC12{Cf22}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC15 struct {
	Cf24 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf24 D5) D5 {
	return D5{D5_DC15{Cf24}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_()
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC14).Cf23
}

func (_this D5) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D5_DC12).Cf22
}

func (_this D5) Dtor_cf24() D5 {
	return _this.Get_().(D5_DC15).Cf24
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf24) + ")"
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
			_, ok := other.Get_().(D5_DC13)
			return ok
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf26 _dafny.Int
	Cf27 _dafny.Int
	Cf28 _dafny.Array
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf26 _dafny.Int, Cf27 _dafny.Int, Cf28 _dafny.Array) D6 {
	return D6{D6_DC17{Cf26, Cf27, Cf28}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf25 _dafny.Map
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf25 _dafny.Map) D6 {
	return D6{D6_DC16{Cf25}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC18 struct {
	Cf29 D6
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf29 D6) D6 {
	return D6{D6_DC18{Cf29}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(_dafny.Zero, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D6) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf27
}

func (_this D6) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D6_DC17).Cf28
}

func (_this D6) Dtor_cf25() _dafny.Map {
	return _this.Get_().(D6_DC16).Cf25
}

func (_this D6) Dtor_cf29() D6 {
	return _this.Get_().(D6_DC18).Cf29
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
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
	Cf31 _dafny.Int
	Cf32 _dafny.CodePoint
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf31 _dafny.Int, Cf32 _dafny.CodePoint) D7 {
	return D7{D7_DC20{Cf31, Cf32}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC19 struct {
	Cf30 T1
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf30 T1) D7 {
	return D7{D7_DC19{Cf30}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC21 struct {
	Cf33 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf33 D7) D7 {
	return D7{D7_DC21{Cf33}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D7) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf31
}

func (_this D7) Dtor_cf32() _dafny.CodePoint {
	return _this.Get_().(D7_DC20).Cf32
}

func (_this D7) Dtor_cf30() T1 {
	return _this.Get_().(D7_DC19).Cf30
}

func (_this D7) Dtor_cf33() D7 {
	return _this.Get_().(D7_DC21).Cf33
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf33) + ")"
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
			return ok && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && _dafny.AreEqual(data1.Cf30, data2.Cf30)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf33.Equals(data2.Cf33)
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

type D8_DC22 struct {
	Cf34 _dafny.Sequence
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf34 _dafny.Sequence) D8 {
	return D8{D8_DC22{Cf34}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D8_DC22).Cf34
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D9_DC24 struct {
	Cf36 bool
	Cf37 bool
	Cf38 _dafny.CodePoint
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf36 bool, Cf37 bool, Cf38 _dafny.CodePoint) D9 {
	return D9{D9_DC24{Cf36, Cf37, Cf38}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC25 struct {
	Cf39 bool
	Cf40 _dafny.Int
	Cf41 bool
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf39 bool, Cf40 _dafny.Int, Cf41 bool) D9 {
	return D9{D9_DC25{Cf39, Cf40, Cf41}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC23 struct {
	Cf35 _dafny.MultiSet
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf35 _dafny.MultiSet) D9 {
	return D9{D9_DC23{Cf35}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(false, false, _dafny.CodePoint('D'))
}

func (_this D9) Dtor_cf36() bool {
	return _this.Get_().(D9_DC24).Cf36
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC24).Cf37
}

func (_this D9) Dtor_cf38() _dafny.CodePoint {
	return _this.Get_().(D9_DC24).Cf38
}

func (_this D9) Dtor_cf39() bool {
	return _this.Get_().(D9_DC25).Cf39
}

func (_this D9) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf40
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC25).Cf41
}

func (_this D9) Dtor_cf35() _dafny.MultiSet {
	return _this.Get_().(D9_DC23).Cf35
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41 == data2.Cf41
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf35.Equals(data2.Cf35)
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
	Cf42 _dafny.Map
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf42 _dafny.Map) D10 {
	return D10{D10_DC26{Cf42}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D10) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D10_DC26).Cf42
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf42.Equals(data2.Cf42)
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

type D11_DC28 struct {
	Cf44 bool
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf44 bool) D11 {
	return D11{D11_DC28{Cf44}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC29 struct {
	Cf45 _dafny.Int
	Cf46 _dafny.Int
	Cf47 _dafny.Int
	Cf48 _dafny.Int
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf45 _dafny.Int, Cf46 _dafny.Int, Cf47 _dafny.Int, Cf48 _dafny.Int) D11 {
	return D11{D11_DC29{Cf45, Cf46, Cf47, Cf48}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC30 struct {
	Cf49 bool
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf49 bool) D11 {
	return D11{D11_DC30{Cf49}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC27 struct {
	Cf43 _dafny.Set
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf43 _dafny.Set) D11 {
	return D11{D11_DC27{Cf43}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC31 struct {
	Cf50 D11
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf50 D11) D11 {
	return D11{D11_DC31{Cf50}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC28_(false)
}

func (_this D11) Dtor_cf44() bool {
	return _this.Get_().(D11_DC28).Cf44
}

func (_this D11) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D11_DC29).Cf45
}

func (_this D11) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D11_DC29).Cf46
}

func (_this D11) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D11_DC29).Cf47
}

func (_this D11) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D11_DC29).Cf48
}

func (_this D11) Dtor_cf49() bool {
	return _this.Get_().(D11_DC30).Cf49
}

func (_this D11) Dtor_cf43() _dafny.Set {
	return _this.Get_().(D11_DC27).Cf43
}

func (_this D11) Dtor_cf50() D11 {
	return _this.Get_().(D11_DC31).Cf50
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf44 == data2.Cf44
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf49 == data2.Cf49
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf50.Equals(data2.Cf50)
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
	Cf52 _dafny.Sequence
	Cf53 bool
	Cf54 _dafny.Int
	Cf55 bool
	Cf56 bool
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf52 _dafny.Sequence, Cf53 bool, Cf54 _dafny.Int, Cf55 bool, Cf56 bool) D12 {
	return D12{D12_DC33{Cf52, Cf53, Cf54, Cf55, Cf56}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC32 struct {
	Cf51 _dafny.MultiSet
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf51 _dafny.MultiSet) D12 {
	return D12{D12_DC32{Cf51}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC34 struct {
	Cf57 D12
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf57 D12) D12 {
	return D12{D12_DC34{Cf57}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC33_(_dafny.EmptySeq, false, _dafny.Zero, false, false)
}

func (_this D12) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D12_DC33).Cf52
}

func (_this D12) Dtor_cf53() bool {
	return _this.Get_().(D12_DC33).Cf53
}

func (_this D12) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D12_DC33).Cf54
}

func (_this D12) Dtor_cf55() bool {
	return _this.Get_().(D12_DC33).Cf55
}

func (_this D12) Dtor_cf56() bool {
	return _this.Get_().(D12_DC33).Cf56
}

func (_this D12) Dtor_cf51() _dafny.MultiSet {
	return _this.Get_().(D12_DC32).Cf51
}

func (_this D12) Dtor_cf57() D12 {
	return _this.Get_().(D12_DC34).Cf57
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33" + "(" + data.Cf52.VerbatimString(true) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf51) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf57) + ")"
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
			return ok && data1.Cf52.Equals(data2.Cf52) && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0 && data1.Cf55 == data2.Cf55 && data1.Cf56 == data2.Cf56
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf51.Equals(data2.Cf51)
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf57.Equals(data2.Cf57)
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

// Definition of trait T0
type T0 interface {
	String() string
	F13() bool
	F13_set_(value bool)
	Fm4(p0 bool, p1 D0, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool
	Fm5(globalState *GlobalState) D1
	M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Array
	M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Sequence)
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
	Fm6(globalState *GlobalState) bool
	Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int
	F14() _dafny.Array
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
	_f0  _dafny.Int
	_f1  _dafny.Array
	_f2  bool
	_f3  bool
	_f4  _dafny.Sequence
	_f5  _dafny.Set
	_f6  bool
	_f7  _dafny.MultiSet
	_f8  bool
	_f9  _dafny.Int
	_f10 _dafny.Array
	_f11 _dafny.Sequence
	_f12 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = false
	_this._f3 = false
	_this._f4 = _dafny.EmptySeq
	_this._f5 = _dafny.EmptySet
	_this._f6 = false
	_this._f7 = _dafny.EmptyMultiSet
	_this._f8 = false
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = _dafny.EmptySeq
	_this._f12 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Array, f2 bool, f3 bool, f4 _dafny.Sequence, f5 _dafny.Set, f6 bool, f7 _dafny.MultiSet, f8 bool, f9 _dafny.Int, f10 _dafny.Array, f11 _dafny.Sequence, f12 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
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
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Set {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.MultiSet {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Array {
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f15 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f15 = _dafny.Zero
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

func (_this *C0) Ctor__(f15 _dafny.Int) {
	{
		(_this)._f15 = f15
	}
}
func (_this *C0) Fm10(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Contains(!(false) || (false))
	}
}
func (_this *C0) F15() _dafny.Int {
	{
		return _this._f15
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f14 _dafny.Array
	F18  _dafny.Sequence
	_f17 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F18 = _dafny.EmptySeq
	_this._f17 = false
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

func (_this *C1) F14() _dafny.Array {
	return _this._f14
}
func (_this *C1) Ctor__(f17 bool, f18 _dafny.Sequence, f14 _dafny.Array) {
	{
		(_this)._f17 = f17
		(_this).F18 = f18
		(_this)._f14 = f14
	}
}
func (_this *C1) Fm6(globalState *GlobalState) bool {
	{
		return (_this).F17()
	}
}
func (_this *C1) Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-859)
	}
}
func (_this *C1) Fm14(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((_dafny.IntOfUint32((_this.F18).Cardinality())).Cmp(Companion_Default___.SafeDivisionInt((_dafny.SetOf((_this).F17())).Cardinality(), _dafny.IntOfInt64(41))) >= 0)
	}
}
func (_this *C1) Fm15(globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.MultiSetOf((_this).F17())).Difference(_dafny.MultiSetOf((_this).F17()))).Union(_dafny.MultiSetOf(false, (_this).F17()))).Cardinality()
	}
}
func (_this *C1) M5(p0 _dafny.Array, globalState *GlobalState) {
	{
		var _325_i0 _dafny.Int
		_ = _325_i0
		_325_i0 = _dafny.Zero
		{
			for (_this).F17() {
				{
					if (_325_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_325_i0 = (_325_i0).Plus(_dafny.One)
					var _326_v0 _dafny.Int
					_ = _326_v0
					_326_v0 = _dafny.IntOfInt64(669)
					var _327_v1 bool
					_ = _327_v1
					_327_v1 = true
					var _rhs44 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(703), (_326_v0).Times(_326_v0))
					_ = _rhs44
					var _rhs45 bool = (_this).F17()
					_ = _rhs45
					var _rhs46 bool = (_326_v0).Cmp(_326_v0) <= 0
					_ = _rhs46
					_326_v0 = _rhs44
					_327_v1 = _rhs45
					_327_v1 = _rhs46
					var _328_v2 _dafny.Map
					_ = _328_v2
					_328_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm14(_326_v0, (_this).F17(), _326_v0, _326_v0, globalState), (_this).F14())
					var _329_v3 _dafny.Map
					_ = _329_v3
					_329_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_326_v0, (func() _dafny.Array {
						if (_328_v2).Contains((_this).F17()) {
							return (_328_v2).Get((_this).F17()).(_dafny.Array)
						}
						return (_this).F14()
					})())
					var _rhs47 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_326_v0, (_this).F14())).Merge(_329_v3)).Cardinality()
					_ = _rhs47
					var _rhs48 bool = false
					_ = _rhs48
					_326_v0 = _rhs47
					_327_v1 = _rhs48
					var _330_v4 _dafny.CodePoint
					_ = _330_v4
					_330_v4 = _dafny.CodePoint('t')
					_330_v4 = (_this.F18).Select((Companion_Default___.SafeIndex(_326_v0, _dafny.IntOfUint32((_this.F18).Cardinality()))).Uint32()).(_dafny.CodePoint)
					var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index68
					((_this).F14()).ArraySet1((_326_v0).Minus(_326_v0), (_index68).Int())
					var _331_v5 _dafny.Set
					_ = _331_v5
					_331_v5 = _dafny.SetOf(true, _327_v1, (_this).F17())
					var _332_v6 _dafny.Map
					_ = _332_v6
					_332_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_331_v5).Difference(_dafny.SetOf((_this).F17(), true)), ((_dafny.Zero).Minus(_326_v0)).Minus(_326_v0))
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index69
					((_this).F14()).ArraySet1((_332_v6).Cardinality(), (_index69).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _333_v7 _dafny.Int
		_ = _333_v7
		_333_v7 = _dafny.IntOfInt64(-432)
		var _hi2 _dafny.Int = _dafny.IntOfInt64(98)
		_ = _hi2
		for _334_i1 := _333_v7; _334_i1.Cmp(_hi2) < 0; _334_i1 = _334_i1.Plus(_dafny.One) {
			var _335_v8 _dafny.Map
			_ = _335_v8
			_335_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _334_i1)
			var _336_v9 _dafny.Sequence
			_ = _336_v9
			_336_v9 = _dafny.SeqOf((_this).F17())
			var _337_v10 _dafny.Map
			_ = _337_v10
			_337_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v9, _333_v7)
			var _338_v11 _dafny.Sequence
			_ = _338_v11
			_338_v11 = _dafny.SeqOf(_334_i1, ((_337_v10).Update(_dafny.SeqOf((_this).F17(), (_this).F17()), _334_i1)).Cardinality(), _dafny.IntOfInt64(740))
			var _339_v12 _dafny.Map
			_ = _339_v12
			_339_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (func() _dafny.Sequence {
				if (_this).F17() {
					return Companion_Default___.Fm16(_333_v7, (_335_v8).Cardinality(), (_this).F17(), _334_i1, globalState)
				}
				return _338_v11
			})())
			_339_v12 = (_339_v12).Update(!((_this).F17()) || ((_this).F17()), _338_v11)
			_333_v7 = (Companion_Default___.SafeDivisionInt(_334_i1, _333_v7)).Times(_333_v7)
			_333_v7 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(160), _dafny.IntOfInt64(5))
			var _340_v13 _dafny.MultiSet
			_ = _340_v13
			_340_v13 = _dafny.MultiSetOf(_334_i1)
			var _341_v14 _dafny.Set
			_ = _341_v14
			_341_v14 = _dafny.SetOf((_this).F17())
			var _342_v15 _dafny.Set
			_ = _342_v15
			_342_v15 = _dafny.SetOf(_341_v14)
			_333_v7 = (func() _dafny.Int {
				if (_340_v13).Contains((_333_v7).Plus(_dafny.IntOfUint32((_this.F18).Cardinality()))) {
					return (_340_v13).Multiplicity((_333_v7).Plus(_dafny.IntOfUint32((_this.F18).Cardinality())))
				}
				return (_342_v15).Cardinality()
			})()
		}
		var _343_i2 _dafny.Int
		_ = _343_i2
		_343_i2 = _dafny.Zero
		{
			for !(!((_this).F17())) {
				{
					if (_343_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_343_i2 = (_343_i2).Plus(_dafny.One)
					var _344_v16 *C0
					_ = _344_v16
					var _nw39 *C0 = New_C0_()
					_ = _nw39
					_nw39.Ctor__(_333_v7)
					_344_v16 = _nw39
					var _345_v17 _dafny.Map
					_ = _345_v17
					_345_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v7, (_344_v16).F15())
					var _346_v18 D1
					_ = _346_v18
					_346_v18 = Companion_D1_.Create_DC2_(_345_v17)
					var _source8 D1 = _346_v18
					_ = _source8
					if _source8.Is_DC3() {
						var _347___mcc_h0 _dafny.Int = _source8.Get_().(D1_DC3).Cf4
						_ = _347___mcc_h0
						var _348_cf4 _dafny.Int = _347___mcc_h0
						_ = _348_cf4
						var _349_v19 D1
						_ = _349_v19
						_349_v19 = Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(335))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}(func(_350_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('m')
						}))).Cardinality()))
						var _351_v20 _dafny.Array
						_ = _351_v20
						var _nwElement0_5 D1 = Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_this.F18).Cardinality()))
						_ = _nwElement0_5
						var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(26))
						_ = _nw40
						(_nw40).ArraySet1(_nwElement0_5, 0)
						(_nw40).ArraySet1(_349_v19, 1)
						(_nw40).ArraySet1(_349_v19, 2)
						(_nw40).ArraySet1(Companion_D1_.Create_DC3_((_344_v16).F15()), 3)
						(_nw40).ArraySet1(_349_v19, 4)
						(_nw40).ArraySet1(_349_v19, 5)
						(_nw40).ArraySet1(_349_v19, 6)
						(_nw40).ArraySet1(_349_v19, 7)
						(_nw40).ArraySet1(_349_v19, 8)
						(_nw40).ArraySet1(_349_v19, 9)
						(_nw40).ArraySet1(Companion_D1_.Create_DC3_(_348_cf4), 10)
						(_nw40).ArraySet1(_349_v19, 11)
						(_nw40).ArraySet1(_349_v19, 12)
						(_nw40).ArraySet1(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(115)), 13)
						(_nw40).ArraySet1(_349_v19, 14)
						(_nw40).ArraySet1(_349_v19, 15)
						(_nw40).ArraySet1(_349_v19, 16)
						(_nw40).ArraySet1(_349_v19, 17)
						(_nw40).ArraySet1(_349_v19, 18)
						(_nw40).ArraySet1(_349_v19, 19)
						(_nw40).ArraySet1(_349_v19, 20)
						(_nw40).ArraySet1(_349_v19, 21)
						(_nw40).ArraySet1(_349_v19, 22)
						(_nw40).ArraySet1(_349_v19, 23)
						(_nw40).ArraySet1(_349_v19, 24)
						(_nw40).ArraySet1(_349_v19, 25)
						_351_v20 = _nw40
						var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_351_v20), 0))
						_ = _index70
						(_351_v20).ArraySet1(_349_v19, (_index70).Int())
						var _352_v21 bool
						_ = _352_v21
						_352_v21 = true
						var _353_v22 _dafny.Sequence
						_ = _353_v22
						_353_v22 = _dafny.SeqOf(_dafny.IntOfUint32((_this.F18).Cardinality()), _dafny.IntOfInt64(31), _348_cf4)
						var _pat_let_tv10 = _352_v21
						_ = _pat_let_tv10
						var _pat_let_tv11 = _353_v22
						_ = _pat_let_tv11
						var _pat_let_tv12 = _348_cf4
						_ = _pat_let_tv12
						var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_351_v20), 0))
						_ = _index71
						var _rhs49 D1 = func(_pat_let6_0 D1) D1 {
							return func(_354_dt__update__tmp_h0 D1) D1 {
								return func(_pat_let7_0 _dafny.Int) D1 {
									return func(_355_dt__update_hcf4_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC3_(_355_dt__update_hcf4_h0)
									}(_pat_let7_0)
								}((func() _dafny.Int {
									if _pat_let_tv10 {
										return _dafny.IntOfUint32((_pat_let_tv11).Cardinality())
									}
									return _pat_let_tv12
								})())
							}(_pat_let6_0)
						}(Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("x")).Cardinality())))
						_ = _rhs49
						var _rhs50 _dafny.Int = Companion_Default___.SafeDivisionInt(_333_v7, _348_cf4)
						_ = _rhs50
						var _rhs51 bool = !((_this).F17())
						_ = _rhs51
						var _rhs52 _dafny.Int = (_344_v16).F15()
						_ = _rhs52
						var _rhs53 bool = (_this).F17()
						_ = _rhs53
						var _lhs32 _dafny.Array = _351_v20
						_ = _lhs32
						var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_351_v20), 0))
						_ = _lhs33
						(_lhs32).ArraySet1(_rhs49, (_lhs33).Int())
						_333_v7 = _rhs50
						_352_v21 = _rhs51
						_348_cf4 = _rhs52
						_352_v21 = _rhs53
						var _356_v23 D2
						_ = _356_v23
						_356_v23 = Companion_D2_.Create_DC7_(_this.F18, (_344_v16).F15(), (_344_v16).F15(), (_this).F17())
						var _357_v24 _dafny.Sequence
						_ = _357_v24
						_357_v24 = _dafny.SeqOf(_352_v21, _352_v21)
						var _358_v25 _dafny.CodePoint
						_ = _358_v25
						_358_v25 = _dafny.CodePoint('l')
						var _pat_let_tv13 = _345_v17
						_ = _pat_let_tv13
						var _359_v26 _dafny.MultiSet
						_ = _359_v26
						_359_v26 = _dafny.MultiSetOf(_346_v18, func(_pat_let8_0 D1) D1 {
							return func(_360_dt__update__tmp_h1 D1) D1 {
								return func(_pat_let9_0 _dafny.Map) D1 {
									return func(_361_dt__update_hcf3_h0 _dafny.Map) D1 {
										return Companion_D1_.Create_DC2_(_361_dt__update_hcf3_h0)
									}(_pat_let9_0)
								}(_pat_let_tv13)
							}(_pat_let8_0)
						}(_346_v18))
						var _362_v27 _dafny.Array
						_ = _362_v27
						var _nwElement0_6 bool = _352_v21
						_ = _nwElement0_6
						var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(25))
						_ = _nw41
						(_nw41).ArraySet1(_nwElement0_6, 0)
						(_nw41).ArraySet1((_356_v23).Dtor_cf15(), 1)
						(_nw41).ArraySet1(_352_v21, 2)
						(_nw41).ArraySet1((_this).F17(), 3)
						(_nw41).ArraySet1(_352_v21, 4)
						(_nw41).ArraySet1((_333_v7).Cmp(_dafny.IntOfUint32((_357_v24).Cardinality())) < 0, 5)
						(_nw41).ArraySet1((_this).F17(), 6)
						(_nw41).ArraySet1(_352_v21, 7)
						(_nw41).ArraySet1(_352_v21, 8)
						(_nw41).ArraySet1((_this).F17(), 9)
						(_nw41).ArraySet1((_this).F17(), 10)
						(_nw41).ArraySet1(!(!(_352_v21)), 11)
						(_nw41).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("eyusi"), _358_v25), 12)
						(_nw41).ArraySet1(_352_v21, 13)
						(_nw41).ArraySet1((_357_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_353_v22).Cardinality()), _dafny.IntOfUint32((_357_v24).Cardinality()))).Uint32()).(bool), 14)
						(_nw41).ArraySet1(!(_352_v21), 15)
						(_nw41).ArraySet1(false, 16)
						(_nw41).ArraySet1((_this).F17(), 17)
						(_nw41).ArraySet1(false, 18)
						(_nw41).ArraySet1(_352_v21, 19)
						(_nw41).ArraySet1((_352_v21) == ((_this).F17()), 20)
						(_nw41).ArraySet1((_this).F17(), 21)
						(_nw41).ArraySet1(_352_v21, 22)
						(_nw41).ArraySet1((_this).F17(), 23)
						(_nw41).ArraySet1(!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm17(_359_v26, globalState), _this.F18), 24)
						_362_v27 = _nw41
						var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_362_v27), 0))
						_ = _index72
						(_362_v27).ArraySet1(!(!((_this).F17()) || ((_this).F17())), (_index72).Int())
						var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_362_v27), 0))
						_ = _index73
						(_362_v27).ArraySet1(((_344_v16).F15()).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(250), _dafny.IntOfInt64(848))) == 0, (_index73).Int())
						var _363_v28 _dafny.Sequence
						_ = _363_v28
						_363_v28 = _dafny.SeqOf(_dafny.SeqOf((_344_v16).F15()))
						var _364_v29 _dafny.MultiSet
						_ = _364_v29
						_364_v29 = _dafny.MultiSetOf(_dafny.SeqOf((_344_v16).F15(), (_344_v16).F15()))
						var _rhs54 _dafny.Sequence = (_363_v28).Select((Companion_Default___.SafeIndex((_344_v16).F15(), _dafny.IntOfUint32((_363_v28).Cardinality()))).Uint32()).(_dafny.Sequence)
						_ = _rhs54
						var _rhs55 bool = (_dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_353_v22, (Companion_Default___.SafeIndex(_333_v7, _dafny.IntOfUint32((_353_v22).Cardinality()))).Uint32(), (_344_v16).F15()), Companion_Default___.Fm16(_dafny.IntOfInt64(263), (_344_v16).F15(), !(!((_this).F17())), _333_v7, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(523))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg32 _dafny.Int) interface{} {
								return coer32(arg32)
							}
						}((func(_365_v27 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_366_i4 _dafny.Int) _dafny.Int {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_365_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_365_v27), 0))).Int()).(bool), (_this).F17())).Cardinality()
							}
						})(_362_v27))))).IsDisjointFrom(_364_v29)
						_ = _rhs55
						var _rhs56 _dafny.Int = _dafny.IntOfInt64(172)
						_ = _rhs56
						_353_v22 = _rhs54
						_352_v21 = _rhs55
						_333_v7 = _rhs56
						var _367_v30 _dafny.Int
						_ = _367_v30
						var _368_v31 _dafny.Int
						_ = _368_v31
						var _369_v32 bool
						_ = _369_v32
						var _out18 _dafny.Int
						_ = _out18
						var _out19 _dafny.Int
						_ = _out19
						var _out20 bool
						_ = _out20
						_out18, _out19, _out20 = (_this).M6((_362_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_362_v27), 0))).Int()).(bool), globalState)
						_367_v30 = _out18
						_368_v31 = _out19
						_369_v32 = _out20
					} else if _source8.Is_DC4() {
						var _370___mcc_h1 _dafny.Array = _source8.Get_().(D1_DC4).Cf5
						_ = _370___mcc_h1
						var _371___mcc_h2 _dafny.Int = _source8.Get_().(D1_DC4).Cf6
						_ = _371___mcc_h2
						var _372___mcc_h3 _dafny.CodePoint = _source8.Get_().(D1_DC4).Cf7
						_ = _372___mcc_h3
						var _373_cf7 _dafny.CodePoint = _372___mcc_h3
						_ = _373_cf7
						var _374_cf6 _dafny.Int = _371___mcc_h2
						_ = _374_cf6
						var _375_cf5 _dafny.Array = _370___mcc_h1
						_ = _375_cf5
						var _376_v33 _dafny.Set
						_ = _376_v33
						_376_v33 = _dafny.SetOf((_this).F17(), (_this).F17())
						var _377_v34 _dafny.Map
						_ = _377_v34
						_377_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
							if (_this).F17() {
								return _this.F18
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("mkrgtsv")
						})(), Companion_Default___.SafeDivisionInt(_333_v7, (_376_v33).Cardinality()))
						_377_v34 = (_377_v34).Update(_this.F18, (_344_v16).F15())
						var _378_v35 _dafny.Map
						_ = _378_v35
						_378_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17())
						_374_cf6 = (_dafny.Zero).Minus((_333_v7).Times((_378_v35).Cardinality()))
						(_this).F18 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_this.F18, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-952))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}((func(_379_cf7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_380_i5 _dafny.Int) _dafny.CodePoint {
								return _379_cf7
							}
						})(_373_cf7)))), (Companion_Default___.SafeIndex((_344_v16).F15(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F18, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-952))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}((func(_381_cf7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_382_i5 _dafny.Int) _dafny.CodePoint {
								return _381_cf7
							}
						})(_373_cf7))))).Cardinality()))).Uint32(), _373_cf7)
						var _383_v36 bool
						_ = _383_v36
						_383_v36 = true
						_383_v36 = ((_dafny.IntOfInt64(-718)).Plus(_374_cf6)).Cmp(_dafny.IntOfInt64(447)) <= 0
					} else if _source8.Is_DC5() {
						var _384___mcc_h4 _dafny.Sequence = _source8.Get_().(D1_DC5).Cf8
						_ = _384___mcc_h4
						var _385___mcc_h5 _dafny.Array = _source8.Get_().(D1_DC5).Cf9
						_ = _385___mcc_h5
						var _386___mcc_h6 _dafny.Int = _source8.Get_().(D1_DC5).Cf10
						_ = _386___mcc_h6
						var _387_cf10 _dafny.Int = _386___mcc_h6
						_ = _387_cf10
						var _388_cf9 _dafny.Array = _385___mcc_h5
						_ = _388_cf9
						var _389_cf8 _dafny.Sequence = _384___mcc_h4
						_ = _389_cf8
						var _390_v37 _dafny.Map
						_ = _390_v37
						_390_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v7, (_this).F17())
						var _391_v38 _dafny.Array
						_ = _391_v38
						var _nwElement0_7 D1 = _346_v18
						_ = _nwElement0_7
						var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(21))
						_ = _nw42
						(_nw42).ArraySet1(_nwElement0_7, 0)
						(_nw42).ArraySet1(Companion_D1_.Create_DC2_(_345_v17), 1)
						(_nw42).ArraySet1(_346_v18, 2)
						(_nw42).ArraySet1(_346_v18, 3)
						(_nw42).ArraySet1(_346_v18, 4)
						(_nw42).ArraySet1(_346_v18, 5)
						(_nw42).ArraySet1(_346_v18, 6)
						(_nw42).ArraySet1(_346_v18, 7)
						(_nw42).ArraySet1(_346_v18, 8)
						(_nw42).ArraySet1(Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm18((_this).F17(), (_this).F17(), (_344_v16).F15(), _dafny.IntOfInt64(348), globalState)).Cardinality(), (_390_v37).Cardinality())), 9)
						(_nw42).ArraySet1((func() D1 {
							if !(true) {
								return _346_v18
							}
							return _346_v18
						})(), 10)
						(_nw42).ArraySet1(Companion_D1_.Create_DC2_(_345_v17), 11)
						(_nw42).ArraySet1(_346_v18, 12)
						(_nw42).ArraySet1(_346_v18, 13)
						(_nw42).ArraySet1(_346_v18, 14)
						(_nw42).ArraySet1(_346_v18, 15)
						(_nw42).ArraySet1(_346_v18, 16)
						(_nw42).ArraySet1(_346_v18, 17)
						(_nw42).ArraySet1(_346_v18, 18)
						(_nw42).ArraySet1(Companion_D1_.Create_DC2_(_345_v17), 19)
						(_nw42).ArraySet1(Companion_Default___.Fm19(globalState), 20)
						_391_v38 = _nw42
						var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_391_v38), 0))
						_ = _index74
						(_391_v38).ArraySet1(_346_v18, (_index74).Int())
						var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_391_v38), 0))
						_ = _index75
						(_391_v38).ArraySet1(_346_v18, (_index75).Int())
						var _392_v39 *C0
						_ = _392_v39
						var _nw43 *C0 = New_C0_()
						_ = _nw43
						_nw43.Ctor__(_333_v7)
						_392_v39 = _nw43
						var _393_v40 bool
						_ = _393_v40
						_393_v40 = false
						var _394_v41 _dafny.Sequence
						_ = _394_v41
						_394_v41 = _dafny.SeqOf(_393_v40)
						_393_v40 = (false) || ((_394_v41).Select((Companion_Default___.SafeIndex(_387_cf10, _dafny.IntOfUint32((_394_v41).Cardinality()))).Uint32()).(bool))
						var _395_v42 _dafny.Sequence
						_ = _395_v42
						_395_v42 = _dafny.SeqOf((_dafny.Zero).Minus((_344_v16).F15()), (_344_v16).F15())
						_333_v7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_395_v42, _dafny.Companion_Sequence_.Concatenate(_395_v42, _395_v42))).Cardinality())
					} else {
						var _396___mcc_h7 _dafny.Map = _source8.Get_().(D1_DC2).Cf3
						_ = _396___mcc_h7
						var _397_cf3 _dafny.Map = _396___mcc_h7
						_ = _397_cf3
						var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index76
						((_this).F14()).ArraySet1((_333_v7).Plus((_344_v16).F15()), (_index76).Int())
						var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index77
						((_this).F14()).ArraySet1((_344_v16).F15(), (_index77).Int())
						var _398_v43 _dafny.Array
						_ = _398_v43
						var _nw44 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
						_ = _nw44
						_398_v43 = _nw44
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_398_v43), 0))
						_ = _index78
						(_398_v43).ArraySet1((_this).F17(), (_index78).Int())
						var _399_v44 bool
						_ = _399_v44
						_399_v44 = false
						var _400_v45 _dafny.MultiSet
						_ = _400_v45
						_400_v45 = _dafny.MultiSetOf(_399_v44, !((_this).F17()))
						var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_398_v43), 0))
						_ = _index79
						var _rhs57 bool = !((((_400_v45).Cardinality()).Plus((_344_v16).F15())).Cmp((_344_v16).F15()) != 0)
						_ = _rhs57
						var _rhs58 bool = ((_344_v16).F15()).Cmp(_dafny.IntOfUint32((_this.F18).Cardinality())) < 0
						_ = _rhs58
						var _lhs34 _dafny.Array = _398_v43
						_ = _lhs34
						var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_398_v43), 0))
						_ = _lhs35
						(_lhs34).ArraySet1(_rhs57, (_lhs35).Int())
						_399_v44 = _rhs58
						var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_398_v43), 0))
						_ = _index80
						(_398_v43).ArraySet1(!(true), (_index80).Int())
						_399_v44 = (_398_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_398_v43), 0))).Int()).(bool)
					}
					var _401_v46 _dafny.Sequence
					_ = _401_v46
					_401_v46 = _dafny.SeqOf((_this).F17(), !((_this).F17()))
					var _402_v47 _dafny.Map
					_ = _402_v47
					_402_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_344_v16).F15(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F17(), (_this).F17()), _401_v46))
					var _403_v48 _dafny.Set
					_ = _403_v48
					_403_v48 = _dafny.SetOf(_401_v46, _dafny.SeqOf((_this).F17(), (_this).F17()))
					var _404_v49 _dafny.MultiSet
					_ = _404_v49
					_404_v49 = _dafny.MultiSetOf(true)
					var _405_v50 _dafny.Sequence
					_ = _405_v50
					_405_v50 = _dafny.SeqOf((_403_v48).Cardinality(), (_404_v49).Cardinality())
					_402_v47 = (_402_v47).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_this).F17(), (_this).F17()), (_405_v50).Select((Companion_Default___.SafeIndex(_333_v7, _dafny.IntOfUint32((_405_v50).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality(), _401_v46)
					var _406_v51 D0
					_ = _406_v51
					_406_v51 = Companion_D0_.Create_DC0_((_this).F17())
					var _source9 D0 = _406_v51
					_ = _source9
					if _source9.Is_DC1() {
						var _407___mcc_h8 bool = _source9.Get_().(D0_DC1).Cf1
						_ = _407___mcc_h8
						var _408___mcc_h9 _dafny.Int = _source9.Get_().(D0_DC1).Cf2
						_ = _408___mcc_h9
						var _409_cf2 _dafny.Int = _408___mcc_h9
						_ = _409_cf2
						var _410_cf1 bool = _407___mcc_h8
						_ = _410_cf1
						var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index81
						((_this).F14()).ArraySet1(Companion_Default___.SafeDivisionInt((_344_v16).F15(), _409_cf2), (_index81).Int())
						var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index82
						((_this).F14()).ArraySet1(_409_cf2, (_index82).Int())
						var _411_v52 _dafny.Array
						_ = _411_v52
						var _len0_4 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_4
						var _nw45 _dafny.Array
						_ = _nw45
						if _len0_4.Cmp(_dafny.Zero) == 0 {
							_nw45 = _dafny.NewArray(_len0_4)
						} else {
							var _init4 func(_dafny.Int) _dafny.Int = (func(_412_v16 *C0) func(_dafny.Int) _dafny.Int {
								return func(_413_i6 _dafny.Int) _dafny.Int {
									return (_413_i6).Times((_412_v16).F15())
								}
							})(_344_v16)
							_ = _init4
							var _element0_4 = _init4(_dafny.Zero)
							_ = _element0_4
							_nw45 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
							(_nw45).ArraySet1(_element0_4, 0)
							var _nativeLen0_4 = (_len0_4).Int()
							_ = _nativeLen0_4
							for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
								(_nw45).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
							}
						}
						_411_v52 = _nw45
						var _len0_5 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_5
						var _nw46 _dafny.Array
						_ = _nw46
						if _len0_5.Cmp(_dafny.Zero) == 0 {
							_nw46 = _dafny.NewArray(_len0_5)
						} else {
							var _init5 func(_dafny.Int) _dafny.Int = (func(_414_v16 *C0) func(_dafny.Int) _dafny.Int {
								return func(_415_i7 _dafny.Int) _dafny.Int {
									return (_415_i7).Times((_414_v16).F15())
								}
							})(_344_v16)
							_ = _init5
							var _element0_5 = _init5(_dafny.Zero)
							_ = _element0_5
							_nw46 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
							(_nw46).ArraySet1(_element0_5, 0)
							var _nativeLen0_5 = (_len0_5).Int()
							_ = _nativeLen0_5
							for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
								(_nw46).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
							}
						}
						_411_v52 = _nw46
						var _416_v53 _dafny.Map
						_ = _416_v53
						_416_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _410_cf1)
						var _417_v54 _dafny.Map
						_ = _417_v54
						_417_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(994))
						var _418_v55 _dafny.Set
						_ = _418_v55
						_418_v55 = _dafny.SetOf((_344_v16).F15(), (_344_v16).F15())
						var _419_v56 _dafny.Array
						_ = _419_v56
						var _nwElement0_8 bool = _410_cf1
						_ = _nwElement0_8
						var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
						_ = _nw47
						(_nw47).ArraySet1(_nwElement0_8, 0)
						(_nw47).ArraySet1((_this).F17(), 1)
						(_nw47).ArraySet1(_410_cf1, 2)
						(_nw47).ArraySet1(_410_cf1, 3)
						(_nw47).ArraySet1(!(_410_cf1), 4)
						(_nw47).ArraySet1((_this).F17(), 5)
						(_nw47).ArraySet1(_410_cf1, 6)
						(_nw47).ArraySet1(false, 7)
						(_nw47).ArraySet1(_410_cf1, 8)
						(_nw47).ArraySet1((_this).F17(), 9)
						(_nw47).ArraySet1((_this).F17(), 10)
						(_nw47).ArraySet1(false, 11)
						(_nw47).ArraySet1((_this).F17(), 12)
						(_nw47).ArraySet1(false, 13)
						(_nw47).ArraySet1(_410_cf1, 14)
						(_nw47).ArraySet1(true, 15)
						(_nw47).ArraySet1((_this).F17(), 16)
						(_nw47).ArraySet1(_410_cf1, 17)
						(_nw47).ArraySet1(_410_cf1, 18)
						(_nw47).ArraySet1(true, 19)
						(_nw47).ArraySet1((_401_v46).Select((Companion_Default___.SafeIndex((_417_v54).Cardinality(), _dafny.IntOfUint32((_401_v46).Cardinality()))).Uint32()).(bool), 20)
						(_nw47).ArraySet1((_this).Fm14((_344_v16).F15(), _410_cf1, (_344_v16).F15(), (_418_v55).Cardinality(), globalState), 21)
						(_nw47).ArraySet1((_this).F17(), 22)
						(_nw47).ArraySet1(false, 23)
						(_nw47).ArraySet1(_410_cf1, 24)
						_419_v56 = _nw47
						var _420_v57 _dafny.Sequence
						_ = _420_v57
						_420_v57 = _dafny.SeqOf(_416_v53)
						var _421_v58 _dafny.Map
						_ = _421_v58
						_421_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v56, (_420_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-257))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}(func(_422_i8 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('j')
						}))).Cardinality()), _dafny.IntOfUint32((_420_v57).Cardinality()))).Uint32()).(_dafny.Map))
						var _423_v59 _dafny.Map
						_ = _423_v59
						_423_v59 = _416_v53
						var _424_v60 _dafny.Array
						_ = _424_v60
						var _nwElement0_9 _dafny.Map = _416_v53
						_ = _nwElement0_9
						var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(10))
						_ = _nw48
						(_nw48).ArraySet1(_nwElement0_9, 0)
						(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_410_cf1), !((_this).F17())), 1)
						(_nw48).ArraySet1((_416_v53).Merge(_416_v53), 2)
						(_nw48).ArraySet1((func() _dafny.Map {
							if (_421_v58).Contains(_419_v56) {
								return (_421_v58).Get(_419_v56).(_dafny.Map)
							}
							return (_416_v53).Update(_410_cf1, (func() bool {
								if (_416_v53).Contains((_this).F17()) {
									return (_416_v53).Get((_this).F17()).(bool)
								}
								return _410_cf1
							})())
						})(), 3)
						(_nw48).ArraySet1(_416_v53, 4)
						(_nw48).ArraySet1((_423_v59), 5)
						(_nw48).ArraySet1(_416_v53, 6)
						(_nw48).ArraySet1(_416_v53, 7)
						(_nw48).ArraySet1(_416_v53, 8)
						(_nw48).ArraySet1((Companion_Default___.Fm20(_410_cf1, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int), _410_cf1, globalState)).Merge(_416_v53), 9)
						_424_v60 = _nw48
						var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_424_v60), 0))
						_ = _index83
						(_424_v60).ArraySet1(_416_v53, (_index83).Int())
						var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_424_v60), 0))
						_ = _index84
						(_424_v60).ArraySet1(_416_v53, (_index84).Int())
						_411_v52 = _411_v52
					} else {
						var _425___mcc_h10 bool = _source9.Get_().(D0_DC0).Cf0
						_ = _425___mcc_h10
						var _426_cf0 bool = _425___mcc_h10
						_ = _426_cf0
						var _427_v61 _dafny.Int
						_ = _427_v61
						var _428_v62 _dafny.Int
						_ = _428_v62
						var _429_v63 bool
						_ = _429_v63
						var _out21 _dafny.Int
						_ = _out21
						var _out22 _dafny.Int
						_ = _out22
						var _out23 bool
						_ = _out23
						_out21, _out22, _out23 = (_this).M6((_401_v46).Select((Companion_Default___.SafeIndex((_344_v16).F15(), _dafny.IntOfUint32((_401_v46).Cardinality()))).Uint32()).(bool), globalState)
						_427_v61 = _out21
						_428_v62 = _out22
						_429_v63 = _out23
						_405_v50 = _dafny.Companion_Sequence_.Concatenate(_405_v50, _405_v50)
						var _430_v64 _dafny.CodePoint
						_ = _430_v64
						_430_v64 = _dafny.CodePoint('h')
						var _431_v65 _dafny.Map
						_ = _431_v65
						_431_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_427_v61, _430_v64)
						var _432_v66 _dafny.Map
						_ = _432_v66
						_432_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _431_v65)
						_432_v66 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _431_v65)).Merge(_432_v66)
						var _433_v67 _dafny.Array
						_ = _433_v67
						var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
						_ = _len0_6
						var _nw49 _dafny.Array
						_ = _nw49
						if _len0_6.Cmp(_dafny.Zero) == 0 {
							_nw49 = _dafny.NewArray(_len0_6)
						} else {
							var _init6 func(_dafny.Int) bool = (func(_434_v63 bool) func(_dafny.Int) bool {
								return func(_435_i9 _dafny.Int) bool {
									return !(_434_v63)
								}
							})(_429_v63)
							_ = _init6
							var _element0_6 = _init6(_dafny.Zero)
							_ = _element0_6
							_nw49 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
							(_nw49).ArraySet1(_element0_6, 0)
							var _nativeLen0_6 = (_len0_6).Int()
							_ = _nativeLen0_6
							for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
								(_nw49).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
							}
						}
						_433_v67 = _nw49
						var _436_v68 _dafny.Map
						_ = _436_v68
						_436_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_433_v67, _dafny.SetOf((_this).Fm6(globalState)))
						var _437_v69 _dafny.Set
						_ = _437_v69
						_437_v69 = _dafny.SetOf(_429_v63, true, _429_v63)
						_436_v68 = (_436_v68).Update(_433_v67, _437_v69)
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _438_v70 bool
		_ = _438_v70
		_438_v70 = false
		_438_v70 = _438_v70
		var _439_v71 _dafny.Int
		_ = _439_v71
		var _440_v72 _dafny.Int
		_ = _440_v72
		var _441_v73 bool
		_ = _441_v73
		var _out24 _dafny.Int
		_ = _out24
		var _out25 _dafny.Int
		_ = _out25
		var _out26 bool
		_ = _out26
		_out24, _out25, _out26 = (_this).M6((_333_v7).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-78))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_442_i10 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))).Cardinality())) < 0, globalState)
		_439_v71 = _out24
		_440_v72 = _out25
		_441_v73 = _out26
		var _443_v74 *C0
		_ = _443_v74
		var _nw50 *C0 = New_C0_()
		_ = _nw50
		_nw50.Ctor__((_this).Fm15(globalState))
		_443_v74 = _nw50
	}
}
func (_this *C1) M6(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _444_v0 _dafny.Int
		_ = _444_v0
		_444_v0 = _dafny.IntOfInt64(848)
		var _445_v1 _dafny.Map
		_ = _445_v1
		_445_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v0, p0)
		var _446_v2 _dafny.Map
		_ = _446_v2
		_446_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _445_v1)
		var _447_v4 _dafny.Map
		_ = _447_v4
		_447_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v0, ((func() _dafny.Map {
			if (_446_v2).Contains((_this).F17()) {
				return (_446_v2).Get((_this).F17()).(_dafny.Map)
			}
			return func() _dafny.Map {
				var _coll27 = _dafny.NewMapBuilder()
				_ = _coll27
				for _iter29 := _dafny.Iterate((_445_v1).Keys().Elements()); ; {
					_compr_27, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _448_v3 _dafny.Int
					_448_v3 = interface{}(_compr_27).(_dafny.Int)
					if (_445_v1).Contains(_448_v3) {
						_coll27.Add((_448_v3).Plus(_dafny.IntOfInt64(-153)), (_this).F17())
					}
				}
				return _coll27.ToMap()
			}()
		})()).Cardinality())
		var _449_v5 D0
		_ = _449_v5
		_449_v5 = Companion_D0_.Create_DC1_(p0, _444_v0)
		if Companion_Default___.Fm0(_447_v4, (_dafny.MultiSetOf((_449_v5).Dtor_cf1())).Cardinality(), globalState) {
			_444_v0 = _444_v0
			var _450_v6 _dafny.Map
			_ = _450_v6
			_450_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, p0)
			var _451_v7 _dafny.Map
			_ = _451_v7
			_451_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, _450_v6)
			var _452_v8 _dafny.Int
			_ = _452_v8
			var _out27 _dafny.Int
			_ = _out27
			_out27 = Companion_Default___.M0((func() _dafny.Map {
				if (_451_v7).Contains(_this.F18) {
					return (_451_v7).Get(_this.F18).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(498))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}(func(_453_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				})), (_this).F17())
			})(), globalState)
			_452_v8 = _out27
			var _rhs59 _dafny.Int = _452_v8
			_ = _rhs59
			var _rhs60 bool = p0
			_ = _rhs60
			r1 = _rhs59
			r2 = _rhs60
			var _454_v9 _dafny.Array
			_ = _454_v9
			var _nw51 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw51
			_454_v9 = _nw51
			var _455_v10 _dafny.MultiSet
			_ = _455_v10
			_455_v10 = _dafny.MultiSetOf(_454_v9)
			r2 = !((func() bool {
				if p0 {
					return Companion_Default___.Fm0(_447_v4, (_this).Fm7(_dafny.IntOfInt64(107), (_455_v10).Cardinality(), (_this).F17(), globalState), globalState)
				}
				return (_this).F17()
			})())
			var _456_v11 *C0
			_ = _456_v11
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__(_444_v0)
			_456_v11 = _nw52
		} else {
			var _457_v12 _dafny.Array
			_ = _457_v12
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw53
			_457_v12 = _nw53
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
			_ = _index85
			(_457_v12).ArraySet1(p0, (_index85).Int())
			var _458_v13 D4
			_ = _458_v13
			_458_v13 = Companion_D4_.Create_DC10_((_this).F14())
			var _459_v14 _dafny.Array
			_ = _459_v14
			var _nwElement0_10 _dafny.Array = (_this).F14()
			_ = _nwElement0_10
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(17))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_10, 0)
			(_nw54).ArraySet1((_this).F14(), 1)
			(_nw54).ArraySet1((_458_v13).Dtor_cf20(), 2)
			(_nw54).ArraySet1((_this).F14(), 3)
			(_nw54).ArraySet1((_this).F14(), 4)
			(_nw54).ArraySet1((_this).F14(), 5)
			(_nw54).ArraySet1((_this).F14(), 6)
			(_nw54).ArraySet1((_this).F14(), 7)
			(_nw54).ArraySet1((_this).F14(), 8)
			(_nw54).ArraySet1((_this).F14(), 9)
			(_nw54).ArraySet1((_this).F14(), 10)
			(_nw54).ArraySet1((_this).F14(), 11)
			(_nw54).ArraySet1((_this).F14(), 12)
			(_nw54).ArraySet1((_this).F14(), 13)
			(_nw54).ArraySet1((_this).F14(), 14)
			(_nw54).ArraySet1((_this).F14(), 15)
			(_nw54).ArraySet1((_this).F14(), 16)
			_459_v14 = _nw54
			var _460_v15 _dafny.Array
			_ = _460_v15
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw55
			_460_v15 = _nw55
			var _461_v16 _dafny.Sequence
			_ = _461_v16
			_461_v16 = _dafny.SeqOf((_dafny.Zero).Minus(_444_v0), _444_v0, _444_v0)
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_460_v15), 0))
			_ = _index86
			(_460_v15).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_461_v16, _461_v16), (_index86).Int())
			var _462_v17 _dafny.Array
			_ = _462_v17
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
			_ = _nw56
			_462_v17 = _nw56
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))
			_ = _index87
			(_462_v17).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("okriy"), (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
			_ = _index88
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_460_v15), 0))
			_ = _index89
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))
			_ = _index90
			var _rhs61 bool = false
			_ = _rhs61
			var _rhs62 _dafny.Array = _459_v14
			_ = _rhs62
			var _rhs63 _dafny.Int = (_444_v0).Times(_444_v0)
			_ = _rhs63
			var _rhs64 _dafny.Sequence = _461_v16
			_ = _rhs64
			var _rhs65 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}(func(_463_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), _this.F18)
			_ = _rhs65
			var _lhs36 _dafny.Array = _457_v12
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _460_v15
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_460_v15), 0))
			_ = _lhs39
			var _lhs40 _dafny.Array = _462_v17
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))
			_ = _lhs41
			(_lhs36).ArraySet1(_rhs61, (_lhs37).Int())
			_459_v14 = _rhs62
			r1 = _rhs63
			(_lhs38).ArraySet1(_rhs64, (_lhs39).Int())
			(_lhs40).ArraySet1(_rhs65, (_lhs41).Int())
			var _464_v18 D1
			_ = _464_v18
			_464_v18 = Companion_D1_.Create_DC3_(_444_v0)
			var _pat_let_tv14 = _444_v0
			_ = _pat_let_tv14
			var _source10 D1 = func(_pat_let10_0 D1) D1 {
				return func(_465_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let11_0 _dafny.Int) D1 {
						return func(_466_dt__update_hcf4_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC3_(_466_dt__update_hcf4_h0)
						}(_pat_let11_0)
					}(_pat_let_tv14)
				}(_pat_let10_0)
			}(_464_v18)
			_ = _source10
			if _source10.Is_DC3() {
				var _467___mcc_h0 _dafny.Int = _source10.Get_().(D1_DC3).Cf4
				_ = _467___mcc_h0
				var _468_cf4 _dafny.Int = _467___mcc_h0
				_ = _468_cf4
				var _469_v19 *C0
				_ = _469_v19
				var _nw57 *C0 = New_C0_()
				_ = _nw57
				_nw57.Ctor__(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_444_v0), _468_cf4))
				_469_v19 = _nw57
				var _470_v20 D1
				_ = _470_v20
				_470_v20 = Companion_D1_.Create_DC2_(_447_v4)
				_470_v20 = _470_v20
				_468_cf4 = (_469_v19).F15()
				_457_v12 = _457_v12
			} else if _source10.Is_DC4() {
				var _471___mcc_h1 _dafny.Array = _source10.Get_().(D1_DC4).Cf5
				_ = _471___mcc_h1
				var _472___mcc_h2 _dafny.Int = _source10.Get_().(D1_DC4).Cf6
				_ = _472___mcc_h2
				var _473___mcc_h3 _dafny.CodePoint = _source10.Get_().(D1_DC4).Cf7
				_ = _473___mcc_h3
				var _474_cf7 _dafny.CodePoint = _473___mcc_h3
				_ = _474_cf7
				var _475_cf6 _dafny.Int = _472___mcc_h2
				_ = _475_cf6
				var _476_cf5 _dafny.Array = _471___mcc_h1
				_ = _476_cf5
				var _477_v21 _dafny.Set
				_ = _477_v21
				_477_v21 = _dafny.SetOf((_dafny.Zero).Minus((_this).Fm7(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qqslcms"), (Companion_Default___.SafeIndex(_475_cf6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqslcms")).Cardinality()))).Uint32(), _474_cf7)).Cardinality()), _444_v0, (_457_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))).Int()).(bool), globalState)), _475_cf6, _475_cf6)
				var _478_v22 _dafny.Sequence
				_ = _478_v22
				_478_v22 = _dafny.SeqOf(_477_v21, _477_v21, _477_v21, _477_v21, _477_v21)
				var _479_v23 _dafny.Map
				_ = _479_v23
				_479_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
				var _480_v24 _dafny.Map
				_ = _480_v24
				_480_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm14(_475_cf6, (_457_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))).Int()).(bool), (_479_v23).Cardinality(), _475_cf6, globalState), Companion_Default___.Fm21(!(p0), globalState))
				_478_v22 = (func() _dafny.Sequence {
					if (_480_v24).Contains((_457_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))).Int()).(bool)) {
						return (_480_v24).Get((_457_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))).Int()).(bool)).(_dafny.Sequence)
					}
					return _478_v22
				})()
				var _481_v25 _dafny.Map
				_ = _481_v25
				_481_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _444_v0)
				var _482_v26 D1
				_ = _482_v26
				_482_v26 = Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_474_cf7, _474_cf7)).Cardinality(), (_481_v25).Cardinality())).Update(_dafny.IntOfUint32((_dafny.SeqOf(_444_v0, _dafny.IntOfInt64(113))).Cardinality()), _dafny.IntOfUint32(((_462_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))).Int()).(_dafny.Sequence)).Cardinality())))
				var _483_v27 _dafny.MultiSet
				_ = _483_v27
				_483_v27 = _dafny.MultiSetOf(_482_v26)
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))
				_ = _index91
				(_462_v17).ArraySet1(Companion_Default___.Fm17((_483_v27).Union(_dafny.MultiSetOf(_482_v26, Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_475_cf6, _444_v0)), _482_v26, _482_v26, _482_v26)), globalState), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
				_ = _index92
				(_457_v12).ArraySet1(p0, (_index92).Int())
				var _484_v28 _dafny.Map
				_ = _484_v28
				_484_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_445_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(365))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_485_v1 _dafny.Map, _486_p0 bool, _487_v12 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_488_i2 _dafny.Int) _dafny.Map {
						return (_485_v1).Update(_dafny.IntOfUint32((_dafny.SeqOf((_this).F17(), _486_p0, (_this).F17())).Cardinality()), (_487_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_487_v12), 0))).Int()).(bool))
					}
				})(_445_v1, p0, _457_v12))))
				var _489_v29 _dafny.Sequence
				_ = _489_v29
				_489_v29 = _dafny.SeqOf(_445_v1, _445_v1)
				var _490_v30 _dafny.Map
				_ = _490_v30
				_490_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
					if (_484_v28).Contains(_445_v1) {
						return (_484_v28).Get(_445_v1).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}((func(_491_v1 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_492_i3 _dafny.Int) _dafny.Map {
							return _491_v1
						}
					})(_445_v1)))
				})(), _489_v29), _449_v5)
				_490_v30 = _490_v30
			} else if _source10.Is_DC5() {
				var _493___mcc_h4 _dafny.Sequence = _source10.Get_().(D1_DC5).Cf8
				_ = _493___mcc_h4
				var _494___mcc_h5 _dafny.Array = _source10.Get_().(D1_DC5).Cf9
				_ = _494___mcc_h5
				var _495___mcc_h6 _dafny.Int = _source10.Get_().(D1_DC5).Cf10
				_ = _495___mcc_h6
				var _496_cf10 _dafny.Int = _495___mcc_h6
				_ = _496_cf10
				var _497_cf9 _dafny.Array = _494___mcc_h5
				_ = _497_cf9
				var _498_cf8 _dafny.Sequence = _493___mcc_h4
				_ = _498_cf8
				var _499_v31 *C0
				_ = _499_v31
				var _nw58 *C0 = New_C0_()
				_ = _nw58
				_nw58.Ctor__(_444_v0)
				_499_v31 = _nw58
				_499_v31 = _499_v31
				var _500_v32 *C0
				_ = _500_v32
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(_dafny.IntOfInt64(523))
				_500_v32 = _nw59
				var _501_v33 *C0
				_ = _501_v33
				var _nw60 *C0 = New_C0_()
				_ = _nw60
				_nw60.Ctor__((_499_v31).F15())
				_501_v33 = _nw60
				r0 = _496_cf10
			} else {
				var _502___mcc_h7 _dafny.Map = _source10.Get_().(D1_DC2).Cf3
				_ = _502___mcc_h7
				var _503_cf3 _dafny.Map = _502___mcc_h7
				_ = _503_cf3
				_444_v0 = _444_v0
				_444_v0 = ((func() _dafny.Int {
					if p0 {
						return _444_v0
					}
					return _444_v0
				})()).Minus(_444_v0)
				var _504_v34 _dafny.CodePoint
				_ = _504_v34
				_504_v34 = _dafny.CodePoint('b')
				var _505_v35 _dafny.Map
				_ = _505_v35
				_505_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v0, _504_v34)
				var _506_v36 _dafny.Map
				_ = _506_v36
				_506_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _505_v35)
				_506_v36 = (_506_v36).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _505_v35))
				var _507_v37 _dafny.Array
				_ = _507_v37
				var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw61
				_507_v37 = _nw61
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_507_v37), 0))
				_ = _index93
				(_507_v37).ArraySet1(_462_v17, (_index93).Int())
				var _508_v38 _dafny.MultiSet
				_ = _508_v38
				_508_v38 = _dafny.MultiSetOf(p0)
				var _509_v39 _dafny.MultiSet
				_ = _509_v39
				_509_v39 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(221))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_510_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_511_i4 _dafny.Int) _dafny.CodePoint {
						return _510_v34
					}
				})(_504_v34))))
				var _512_v40 D2
				_ = _512_v40
				_512_v40 = Companion_D2_.Create_DC7_((_462_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))).Int()).(_dafny.Sequence), _444_v0, (func() _dafny.Int {
					if (_509_v39).Contains((_462_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))).Int()).(_dafny.Sequence)) {
						return (_509_v39).Multiplicity((_462_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_462_v17), 0))).Int()).(_dafny.Sequence))
					}
					return _dafny.IntOfInt64(204)
				})(), (_this).F17())
				var _513_v41 _dafny.Map
				_ = _513_v41
				_513_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_504_v34, _447_v4)
				var _514_v42 D1
				_ = _514_v42
				_514_v42 = Companion_D1_.Create_DC2_((func() _dafny.Map {
					if (_513_v41).Contains(_504_v34) {
						return (_513_v41).Get(_504_v34).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v0, _dafny.IntOfUint32((_461_v16).Cardinality()))
				})())
				var _515_v43 _dafny.MultiSet
				_ = _515_v43
				_515_v43 = _dafny.MultiSetOf(_514_v42, _514_v42, Companion_Default___.Fm19(globalState))
				var _516_v44 D1
				_ = _516_v44
				_516_v44 = Companion_D1_.Create_DC5_(Companion_Default___.Fm17(_515_v43, globalState), _457_v12, _444_v0)
				var _517_v45 _dafny.Set
				_ = _517_v45
				_517_v45 = _dafny.SetOf(false)
				var _518_v46 _dafny.Sequence
				_ = _518_v46
				_518_v46 = _dafny.SeqOf(_517_v45, _517_v45)
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_507_v37), 0))
				_ = _index94
				var _rhs66 _dafny.Array = _462_v17
				_ = _rhs66
				var _rhs67 _dafny.Int = (_512_v40).Dtor_cf14()
				_ = _rhs67
				var _rhs68 _dafny.Int = (_516_v44).Dtor_cf10()
				_ = _rhs68
				var _rhs69 _dafny.MultiSet = (_508_v38).Update((_444_v0).Cmp(((_518_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.IntOfUint32((_518_v46).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()) < 0, Companion_Default___.Abs(_444_v0))
				_ = _rhs69
				var _rhs70 _dafny.Int = _444_v0
				_ = _rhs70
				var _lhs42 _dafny.Array = _507_v37
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_507_v37), 0))
				_ = _lhs43
				(_lhs42).ArraySet1(_rhs66, (_lhs43).Int())
				r1 = _rhs67
				r1 = _rhs68
				_508_v38 = _rhs69
				_444_v0 = _rhs70
			}
			_457_v12 = _457_v12
			var _519_v47 _dafny.Sequence
			_ = _519_v47
			_519_v47 = _dafny.SeqOf(_457_v12, _457_v12)
			var _520_v48 _dafny.CodePoint
			_ = _520_v48
			_520_v48 = _dafny.CodePoint('m')
			var _521_v49 _dafny.Array
			_ = _521_v49
			var _nwElement0_11 _dafny.Array = (_519_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_520_v48)).Cardinality()), _dafny.IntOfUint32((_519_v47).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _nwElement0_11
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(19))
			_ = _nw62
			(_nw62).ArraySet1(_nwElement0_11, 0)
			(_nw62).ArraySet1(_457_v12, 1)
			(_nw62).ArraySet1((func() _dafny.Array {
				if p0 {
					return _457_v12
				}
				return _457_v12
			})(), 2)
			(_nw62).ArraySet1(_457_v12, 3)
			(_nw62).ArraySet1(_457_v12, 4)
			(_nw62).ArraySet1(_457_v12, 5)
			(_nw62).ArraySet1(_457_v12, 6)
			(_nw62).ArraySet1(_457_v12, 7)
			(_nw62).ArraySet1(_457_v12, 8)
			(_nw62).ArraySet1(_457_v12, 9)
			(_nw62).ArraySet1((_519_v47).Select((Companion_Default___.SafeIndex(_444_v0, _dafny.IntOfUint32((_519_v47).Cardinality()))).Uint32()).(_dafny.Array), 10)
			(_nw62).ArraySet1(_457_v12, 11)
			(_nw62).ArraySet1(_457_v12, 12)
			(_nw62).ArraySet1(_457_v12, 13)
			(_nw62).ArraySet1(_457_v12, 14)
			(_nw62).ArraySet1(_457_v12, 15)
			(_nw62).ArraySet1(_457_v12, 16)
			(_nw62).ArraySet1(_457_v12, 17)
			(_nw62).ArraySet1(_457_v12, 18)
			_521_v49 = _nw62
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_521_v49), 0))
			_ = _index95
			(_521_v49).ArraySet1(_457_v12, (_index95).Int())
			var _522_v50 _dafny.Map
			_ = _522_v50
			_522_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if (_this).F17() {
					return p0
				}
				return (_this).F17()
			})()), (_this).F17())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_521_v49), 0))
			_ = _index96
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
			_ = _index97
			var _rhs71 _dafny.Array = (_519_v47).Select((Companion_Default___.SafeIndex(_444_v0, _dafny.IntOfUint32((_519_v47).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs71
			var _rhs72 bool = (func() bool {
				if (_522_v50).Contains((p0) && ((_this).F17())) {
					return (_522_v50).Get((p0) && ((_this).F17())).(bool)
				}
				return p0
			})()
			_ = _rhs72
			var _lhs44 _dafny.Array = _521_v49
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_521_v49), 0))
			_ = _lhs45
			var _lhs46 _dafny.Array = _457_v12
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
			_ = _lhs47
			(_lhs44).ArraySet1(_rhs71, (_lhs45).Int())
			(_lhs46).ArraySet1(_rhs72, (_lhs47).Int())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_457_v12), 0))
			_ = _index98
			(_457_v12).ArraySet1((_444_v0).Cmp(_444_v0) > 0, (_index98).Int())
		}
		var _523_i5 _dafny.Int
		_ = _523_i5
		_523_i5 = _dafny.Zero
		{
			for p0 {
				{
					if (_523_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_523_i5 = (_523_i5).Plus(_dafny.One)
					r2 = (_this).F17()
					r2 = (_this).F17()
					(_this).F18 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if false {
							return _this.F18
						}
						return _this.F18
					})(), _this.F18)
					var _524_v51 _dafny.Sequence
					_ = _524_v51
					_524_v51 = _dafny.SeqOf(_444_v0, _444_v0, (func() _dafny.Int {
						if false {
							return _444_v0
						}
						return (_dafny.Zero).Minus(_444_v0)
					})())
					var _rhs73 _dafny.Int = (_444_v0).Minus(_444_v0)
					_ = _rhs73
					var _rhs74 _dafny.Int = _444_v0
					_ = _rhs74
					var _rhs75 bool = (_this).F17()
					_ = _rhs75
					var _rhs76 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_524_v51, _524_v51)
					_ = _rhs76
					_444_v0 = _rhs73
					r0 = _rhs74
					r2 = _rhs75
					_524_v51 = _rhs76
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _525_v52 _dafny.Array
		_ = _525_v52
		var _nwElement0_12 bool = p0
		_ = _nwElement0_12
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_12, 0)
		(_nw63).ArraySet1((_this).F17(), 1)
		(_nw63).ArraySet1((_this).F17(), 2)
		(_nw63).ArraySet1((_this).F17(), 3)
		_525_v52 = _nw63
		var _526_v53 _dafny.CodePoint
		_ = _526_v53
		_526_v53 = _dafny.CodePoint('o')
		var _527_v54 D1
		_ = _527_v54
		_527_v54 = Companion_D1_.Create_DC4_(_525_v52, _444_v0, _526_v53)
		var _528_v55 _dafny.MultiSet
		_ = _528_v55
		_528_v55 = _dafny.MultiSetOf(_525_v52, _525_v52)
		var _529_v56 _dafny.Sequence
		_ = _529_v56
		_529_v56 = _dafny.SeqOf(_528_v55, _528_v55, _528_v55)
		r0 = ((_527_v54).Dtor_cf6()).Minus(((_529_v56).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.IntOfUint32((_529_v56).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality())
		var _530_v57 _dafny.Array
		_ = _530_v57
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_7
		var _nw64 _dafny.Array
		_ = _nw64
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw64 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) D0 = func(_531_i6 _dafny.Int) D0 {
				return Companion_D0_.Create_DC0_(true)
			}
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
		_530_v57 = _nw64
		var _532_v58 D0
		_ = _532_v58
		_532_v58 = Companion_D0_.Create_DC0_((_this).F17())
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_530_v57), 0))
		_ = _index99
		(_530_v57).ArraySet1(_532_v58, (_index99).Int())
		var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_530_v57), 0))
		_ = _index100
		(_530_v57).ArraySet1(_532_v58, (_index100).Int())
		r2 = !_dafny.Companion_Sequence_.Contains(_this.F18, _526_v53)
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index101
		((_this).F14()).ArraySet1(_444_v0, (_index101).Int())
		var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index102
		((_this).F14()).ArraySet1(_dafny.IntOfInt64(-219), (_index102).Int())
		var _533_v59 _dafny.Set
		_ = _533_v59
		_533_v59 = _dafny.SetOf((_this).F17())
		var _534_v60 _dafny.MultiSet
		_ = _534_v60
		_534_v60 = _dafny.MultiSetOf(p0, (_533_v59).IsDisjointFrom(_533_v59), (true) || (Companion_Default___.Fm0(_447_v4, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int), globalState)))
		r0 = (func() _dafny.Int {
			if (_534_v60).Contains(true) {
				return (_534_v60).Multiplicity(true)
			}
			return _dafny.IntOfInt64(870)
		})()
		r1 = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int)
		r2 = p0
		return r0, r1, r2
	}
}
func (_this *C1) F17() bool {
	{
		return _this._f17
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f13 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f13 = false
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

func (_this *C2) F13() bool {
	return _this._f13
}
func (_this *C2) F13_set_(value bool) {
	_this._f13 = value
}
func (_this *C2) Ctor__(f13 bool) {
	{
		(_this)._f13 = f13
	}
}
func (_this *C2) Fm4(p0 bool, p1 D0, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return _this.F13()
	}
}
func (_this *C2) Fm5(globalState *GlobalState) D1 {
	{
		var _source11 D5 = Companion_D5_.Create_DC13_()
		_ = _source11
		if _source11.Is_DC13() {
			return Companion_D1_.Create_DC3_(_dafny.IntOfInt64(614))
		} else if _source11.Is_DC14() {
			var _535___mcc_h0 bool = _source11.Get_().(D5_DC14).Cf23
			_ = _535___mcc_h0
			var _536_cf23 bool = _535___mcc_h0
			_ = _536_cf23
			return Companion_D1_.Create_DC3_(_dafny.IntOfInt64(504))
		} else if _source11.Is_DC12() {
			var _537___mcc_h1 _dafny.Map = _source11.Get_().(D5_DC12).Cf22
			_ = _537___mcc_h1
			var _538_cf22 _dafny.Map = _537___mcc_h1
			_ = _538_cf22
			return Companion_D1_.Create_DC3_((func() _dafny.Int {
				if (_538_cf22).Contains(_this.F13()) {
					return (_538_cf22).Get(_this.F13()).(_dafny.Int)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())).Cardinality()
			})())
		} else {
			var _539___mcc_h2 D5 = _source11.Get_().(D5_DC15).Cf24
			_ = _539___mcc_h2
			var _540_cf24 D5 = _539___mcc_h2
			_ = _540_cf24
			return Companion_D1_.Create_DC3_((_dafny.SetOf(_this.F13())).Cardinality())
		}
	}
}
func (_this *C2) Fm29(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D5 {
	{
		if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(248), _this.F13())).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eaogdxqm")).Cardinality()))).Equals(func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter30 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(214), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bqw")).Cardinality()))).Elements()); ; {
				_compr_28, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _541_v0 _dafny.Int
				_541_v0 = interface{}(_compr_28).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(214), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bqw")).Cardinality())), _541_v0) {
					_coll28.Add((_541_v0).Plus(_dafny.IntOfInt64(24)), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_this.F13())).Cardinality()), _dafny.IntOfInt64(-428))).Cardinality()))
				}
			}
			return _coll28.ToMap()
		}()) {
			return Companion_D5_.Create_DC13_()
		} else {
			return Companion_D5_.Create_DC13_()
		}
	}
}
func (_this *C2) Fm30(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		return _this.F13()
	}
}
func (_this *C2) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _542_v0 _dafny.Map
		_ = _542_v0
		_542_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())
		var _hi3 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(p0))
		_ = _hi3
		for _543_i0 := ((_542_v0).Cardinality()).Plus(p0); _543_i0.Cmp(_hi3) < 0; _543_i0 = _543_i0.Plus(_dafny.One) {
			if true {
				var _544_v1 _dafny.Int
				_ = _544_v1
				_544_v1 = _dafny.IntOfInt64(661)
				var _545_v2 _dafny.CodePoint
				_ = _545_v2
				_545_v2 = _dafny.CodePoint('v')
				var _546_v3 _dafny.Map
				_ = _546_v3
				_546_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_v2, _543_i0)
				var _547_v4 _dafny.Sequence
				_ = _547_v4
				_547_v4 = _dafny.SeqOf((_546_v3).Cardinality())
				_544_v1 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_547_v4).Cardinality())), (Companion_D2_.Create_DC7_(_dafny.SeqOf(_545_v2, _545_v2), p0, _544_v1, false)).Dtor_cf14())
				var _548_v5 _dafny.Map
				_ = _548_v5
				_548_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _549_v6 _dafny.Sequence
				_ = _549_v6
				_549_v6 = _dafny.SeqOf(_this.F13())
				var _550_v7 _dafny.Map
				_ = _550_v7
				_550_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm31(globalState), (_549_v6).Select((Companion_Default___.SafeIndex(_543_i0, _dafny.IntOfUint32((_549_v6).Cardinality()))).Uint32()).(bool))
				var _551_v8 _dafny.Map
				_ = _551_v8
				_551_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_548_v5).Contains(_543_i0) {
						return (_548_v5).Get(_543_i0).(_dafny.Int)
					}
					return Companion_Default___.Fm31(globalState)
				})(), _550_v7)
				_551_v8 = (_551_v8).Update(_543_i0, _550_v7)
				var _552_v9 _dafny.Sequence
				_ = _552_v9
				_552_v9 = _dafny.UnicodeSeqOfUtf8Bytes("ioqxh")
				var _553_v10 _dafny.Map
				_ = _553_v10
				_553_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_552_v9, !(_this.F13()))
				var _554_v11 _dafny.Int
				_ = _554_v11
				var _out28 _dafny.Int
				_ = _out28
				_out28 = Companion_Default___.M0(_553_v10, globalState)
				_554_v11 = _out28
				var _555_v12 _dafny.Array
				_ = _555_v12
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw65
				_555_v12 = _nw65
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_555_v12), 0))
				_ = _index103
				(_555_v12).ArraySet1(_554_v11, (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_555_v12), 0))
				_ = _index104
				(_555_v12).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_543_i0)), (_index104).Int())
				(_this).F13_set_(_this.F13())
			} else {
				(_this).F13_set_(_this.F13())
				var _556_v13 _dafny.Sequence
				_ = _556_v13
				_556_v13 = _dafny.UnicodeSeqOfUtf8Bytes("fnbb")
				var _557_v14 _dafny.Array
				_ = _557_v14
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw66
				_557_v14 = _nw66
				var _558_v15 _dafny.CodePoint
				_ = _558_v15
				_558_v15 = _dafny.CodePoint('j')
				var _559_v16 D1
				_ = _559_v16
				_559_v16 = Companion_D1_.Create_DC4_(_557_v14, p0, _558_v15)
				var _560_v17 _dafny.Set
				_ = _560_v17
				_560_v17 = _dafny.SetOf(_this.F13())
				var _561_v18 _dafny.Map
				_ = _561_v18
				_561_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _560_v17)
				var _562_v19 _dafny.Map
				_ = _562_v19
				_562_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _560_v17)
				var _563_v20 D1
				_ = _563_v20
				_563_v20 = Companion_D1_.Create_DC5_(_556_v13, (_559_v16).Dtor_cf5(), ((func() _dafny.Set {
					if (_561_v18).Contains(p0) {
						return (_561_v18).Get(p0).(_dafny.Set)
					}
					return (func() _dafny.Set {
						if (_562_v19).Contains(_this.F13()) {
							return (_562_v19).Get(_this.F13()).(_dafny.Set)
						}
						return _560_v17
					})()
				})()).Cardinality())
				var _564_v21 _dafny.Map
				_ = _564_v21
				_564_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), (_542_v0).Cardinality())
				var _565_v22 _dafny.Sequence
				_ = _565_v22
				_565_v22 = _dafny.SeqOf(_556_v13)
				var _pat_let_tv15 = _556_v13
				_ = _pat_let_tv15
				var _566_v23 _dafny.Array
				_ = _566_v23
				var _nwElement0_13 D1 = _563_v20
				_ = _nwElement0_13
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(23))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_13, 0)
				(_nw67).ArraySet1(_563_v20, 1)
				(_nw67).ArraySet1(func(_pat_let12_0 D1) D1 {
					return func(_567_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let13_0 _dafny.Sequence) D1 {
							return func(_568_dt__update_hcf8_h0 _dafny.Sequence) D1 {
								return func(_pat_let14_0 _dafny.Int) D1 {
									return func(_569_dt__update_hcf10_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC5_(_568_dt__update_hcf8_h0, (_567_dt__update__tmp_h0).Dtor_cf9(), _569_dt__update_hcf10_h0)
									}(_pat_let14_0)
								}(_543_i0)
							}(_pat_let13_0)
						}(_pat_let_tv15)
					}(_pat_let12_0)
				}(_563_v20), 2)
				(_nw67).ArraySet1(_563_v20, 3)
				(_nw67).ArraySet1(Companion_D1_.Create_DC5_(_556_v13, _557_v14, _543_i0), 4)
				(_nw67).ArraySet1(_563_v20, 5)
				(_nw67).ArraySet1(Companion_D1_.Create_DC5_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_570_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_571_i1 _dafny.Int) _dafny.CodePoint {
						return _570_v15
					}
				})(_558_v15))), _557_v14, _543_i0), 6)
				(_nw67).ArraySet1(_563_v20, 7)
				(_nw67).ArraySet1(Companion_D1_.Create_DC5_(_556_v13, _557_v14, _543_i0), 8)
				(_nw67).ArraySet1(_563_v20, 9)
				(_nw67).ArraySet1(_563_v20, 10)
				(_nw67).ArraySet1(_563_v20, 11)
				(_nw67).ArraySet1(_563_v20, 12)
				(_nw67).ArraySet1(_563_v20, 13)
				(_nw67).ArraySet1(_563_v20, 14)
				(_nw67).ArraySet1(_563_v20, 15)
				(_nw67).ArraySet1(Companion_D1_.Create_DC5_(Companion_Default___.Fm32(_dafny.UnicodeSeqOfUtf8Bytes("xehdabxv"), _this.F13(), (func() _dafny.Int {
					if (_564_v21).Contains(_this.F13()) {
						return (_564_v21).Get(_this.F13()).(_dafny.Int)
					}
					return p0
				})(), _543_i0, globalState), _557_v14, _543_i0), 16)
				(_nw67).ArraySet1(Companion_D1_.Create_DC5_(_556_v13, _557_v14, _dafny.IntOfUint32((_565_v22).Cardinality())), 17)
				(_nw67).ArraySet1(_563_v20, 18)
				(_nw67).ArraySet1(_563_v20, 19)
				(_nw67).ArraySet1(_563_v20, 20)
				(_nw67).ArraySet1(_563_v20, 21)
				(_nw67).ArraySet1(_563_v20, 22)
				_566_v23 = _nw67
				var _rhs77 _dafny.Array = (func() _dafny.Array {
					if _this.F13() {
						return _566_v23
					}
					return _566_v23
				})()
				_ = _rhs77
				var _rhs78 bool = _this.F13()
				_ = _rhs78
				var _lhs48 *C2 = _this
				_ = _lhs48
				_566_v23 = _rhs77
				_lhs48.F13_set_(_rhs78)
				var _572_v24 _dafny.Array
				_ = _572_v24
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_8
				var _nw68 _dafny.Array
				_ = _nw68
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw68 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = (func(_573_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_574_i2 _dafny.Int) _dafny.Int {
							return (_574_i2).Plus((_dafny.Zero).Minus(_573_i0))
						}
					})(_543_i0)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw68 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw68).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw68).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_572_v24 = _nw68
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_572_v24), 0))
				_ = _index105
				(_572_v24).ArraySet1(Companion_Default___.Fm31(globalState), (_index105).Int())
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_572_v24), 0))
				_ = _index106
				(_572_v24).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(593))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_575_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_576_i3 _dafny.Int) _dafny.CodePoint {
						return _575_v15
					}
				})(_558_v15))), _556_v13), _dafny.Companion_Sequence_.Concatenate(_556_v13, _556_v13))).Cardinality()), (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_572_v24), 0))
				_ = _index107
				(_572_v24).ArraySet1((_572_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_572_v24), 0))).Int()).(_dafny.Int), (_index107).Int())
				(_this).F13_set_(_this.F13())
			}
			var _577_v25 _dafny.Sequence
			_ = _577_v25
			_577_v25 = _dafny.UnicodeSeqOfUtf8Bytes("hujnpvv")
			var _578_v26 _dafny.CodePoint
			_ = _578_v26
			_578_v26 = _dafny.CodePoint('w')
			(_this).F13_set_(!(_dafny.Companion_Sequence_.IsPrefixOf(_577_v25, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_577_v25, _577_v25), (Companion_Default___.SafeIndex(_543_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_577_v25, _577_v25)).Cardinality()))).Uint32(), _578_v26), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_577_v25, _577_v25), (Companion_Default___.SafeIndex(_543_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_577_v25, _577_v25)).Cardinality()))).Uint32(), _578_v26)).Cardinality()))).Uint32(), _dafny.CodePoint('f')))))
			var _579_v27 _dafny.Int
			_ = _579_v27
			_579_v27 = _dafny.IntOfInt64(918)
			_579_v27 = _543_i0
			_579_v27 = p0
		}
		var _580_v28 _dafny.Array
		_ = _580_v28
		var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw69
		_580_v28 = _nw69
		var _581_v29 _dafny.Map
		_ = _581_v29
		_581_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_580_v28, _580_v28)
		_581_v29 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_580_v28, _580_v28)).Merge((_581_v29).Merge(_581_v29))
		var _582_v30 *C0
		_ = _582_v30
		var _nw70 *C0 = New_C0_()
		_ = _nw70
		_nw70.Ctor__(p0)
		_582_v30 = _nw70
		var _583_v31 _dafny.Array
		_ = _583_v31
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_9
		var _nw71 _dafny.Array
		_ = _nw71
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw71 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) bool = func(_584_i4 _dafny.Int) bool {
				return _this.F13()
			}
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw71 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw71).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw71).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_583_v31 = _nw71
		var _585_v32 D1
		_ = _585_v32
		_585_v32 = Companion_D1_.Create_DC4_(_583_v31, (_582_v30).F15(), _dafny.CodePoint('m'))
		var _586_v33 _dafny.Sequence
		_ = _586_v33
		_586_v33 = _dafny.UnicodeSeqOfUtf8Bytes("nqogre")
		var _587_v34 _dafny.Map
		_ = _587_v34
		_587_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
			if false {
				return _dafny.CodePoint('t')
			}
			return (_585_v32).Dtor_cf7()
		})(), _586_v33)
		var _588_v35 _dafny.CodePoint
		_ = _588_v35
		_588_v35 = _dafny.CodePoint('f')
		_587_v34 = (_587_v34).Update(_588_v35, _586_v33)
		var _589_v36 _dafny.Sequence
		_ = _589_v36
		_589_v36 = _dafny.SeqOf(_this.F13())
		var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_580_v28), 0))
		_ = _index108
		(_580_v28).ArraySet1(_dafny.IntOfUint32((_589_v36).Cardinality()), (_index108).Int())
		var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_580_v28), 0))
		_ = _index109
		(_580_v28).ArraySet1((_582_v30).F15(), (_index109).Int())
		var _590_v37 _dafny.MultiSet
		_ = _590_v37
		_590_v37 = _dafny.MultiSetOf((_580_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_580_v28), 0))).Int()).(_dafny.Int))
		_589_v36 = _dafny.SeqOf((_590_v37).IsProperSubsetOf(_590_v37), !((_dafny.IntOfInt64(523)).Cmp(p0) < 0))
		r0 = _580_v28
		return r0
	}
}
func (_this *C2) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Sequence) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _591_v0 _dafny.Set
		_ = _591_v0
		_591_v0 = _dafny.SetOf(p0)
		var _hi4 _dafny.Int = p2
		_ = _hi4
		for _592_i0 := (Companion_Default___.Fm33(_this.F13(), false, _591_v0, _this.F13(), globalState)).Cardinality(); _592_i0.Cmp(_hi4) < 0; _592_i0 = _592_i0.Plus(_dafny.One) {
			var _593_v1 _dafny.Array
			_ = _593_v1
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw72
			_593_v1 = _nw72
			var _594_v2 _dafny.Array
			_ = _594_v2
			var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw73
			_594_v2 = _nw73
			var _595_v3 _dafny.CodePoint
			_ = _595_v3
			_595_v3 = _dafny.CodePoint('p')
			var _596_v4 _dafny.MultiSet
			_ = _596_v4
			_596_v4 = _dafny.MultiSetOf(_595_v3)
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_594_v2), 0))
			_ = _index110
			(_594_v2).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wawnlrb")).Cardinality()), (func() _dafny.Int {
				if (_596_v4).Contains(_595_v3) {
					return (_596_v4).Multiplicity(_595_v3)
				}
				return _592_i0
			})()), (_index110).Int())
			var _597_v5 _dafny.Set
			_ = _597_v5
			_597_v5 = _dafny.SetOf(_this.F13(), _this.F13(), _this.F13())
			var _598_v6 _dafny.Sequence
			_ = _598_v6
			_598_v6 = _dafny.SeqOf(p0, Companion_Default___.Fm31(globalState), (_597_v5).Cardinality(), p0)
			var _599_v7 _dafny.MultiSet
			_ = _599_v7
			_599_v7 = _dafny.MultiSetOf(_598_v6, _598_v6)
			var _600_v8 _dafny.Map
			_ = _600_v8
			_600_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _592_i0)
			var _601_v9 D5
			_ = _601_v9
			_601_v9 = Companion_D5_.Create_DC12_(_600_v8)
			var _602_v10 _dafny.Map
			_ = _602_v10
			_602_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _601_v9)
			var _603_v11 _dafny.Map
			_ = _603_v11
			_603_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm30(_599_v7, p0, _602_v10, globalState), _dafny.CodePoint('h'))
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_594_v2), 0))
			_ = _index111
			var _rhs79 bool = _this.F13()
			_ = _rhs79
			var _rhs80 _dafny.Array = _593_v1
			_ = _rhs80
			var _rhs81 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tasrqgp")).Cardinality())).Minus(((_603_v11).Merge(_603_v11)).Cardinality())
			_ = _rhs81
			var _lhs49 _dafny.Array = _594_v2
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_594_v2), 0))
			_ = _lhs50
			r2 = _rhs79
			_593_v1 = _rhs80
			(_lhs49).ArraySet1(_rhs81, (_lhs50).Int())
			r1 = (((_594_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_594_v2), 0))).Int()).(_dafny.Int)).Cmp(_592_i0) <= 0) || (_this.F13())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_594_v2), 0))
			_ = _index112
			(_594_v2).ArraySet1(_dafny.IntOfInt64(62), (_index112).Int())
			var _604_v12 _dafny.Array
			_ = _604_v12
			var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(18))
			_ = _nw74
			_604_v12 = _nw74
			var _605_v13 _dafny.Array
			_ = _605_v13
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw75
			_605_v13 = _nw75
			var _606_v14 _dafny.MultiSet
			_ = _606_v14
			_606_v14 = _dafny.MultiSetOf(_605_v13)
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_604_v12), 0))
			_ = _index113
			(_604_v12).ArraySet1(_606_v14, (_index113).Int())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_604_v12), 0))
			_ = _index114
			(_604_v12).ArraySet1((_606_v14).Update(_605_v13, Companion_Default___.Abs(_592_i0)), (_index114).Int())
		}
		var _607_v16 _dafny.Map
		_ = _607_v16
		_607_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), p0)
		var _608_v17 _dafny.CodePoint
		_ = _608_v17
		_608_v17 = _dafny.CodePoint('w')
		var _609_v18 _dafny.Map
		_ = _609_v18
		_609_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_607_v16).Cardinality(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _608_v17), p2)
		var _610_v19 _dafny.Int
		_ = _610_v19
		var _out29 _dafny.Int
		_ = _out29
		_out29 = Companion_Default___.M0(func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter31 := _dafny.Iterate((_609_v18).Keys().Elements()); ; {
				_compr_29, _ok31 := _iter31()
				if !_ok31 {
					break
				}
				var _611_v15 _dafny.Sequence
				_611_v15 = interface{}(_compr_29).(_dafny.Sequence)
				if (_609_v18).Contains(_611_v15) {
					_coll29.Add(_611_v15, _this.F13())
				}
			}
			return _coll29.ToMap()
		}(), globalState)
		_610_v19 = _out29
		var _612_i1 _dafny.Int
		_ = _612_i1
		_612_i1 = _dafny.Zero
		{
			for _this.F13() {
				{
					if (_612_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_612_i1 = (_612_i1).Plus(_dafny.One)
					var _613_v20 _dafny.MultiSet
					_ = _613_v20
					_613_v20 = _dafny.MultiSetOf(_dafny.IntOfInt64(251))
					var _614_v21 D2
					_ = _614_v21
					_614_v21 = Companion_D2_.Create_DC7_(p1, p2, _610_v19, !(_this.F13()))
					var _615_v22 _dafny.Set
					_ = _615_v22
					_615_v22 = _dafny.SetOf(_608_v17)
					var _616_v23 _dafny.Map
					_ = _616_v23
					_616_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_615_v22).Cardinality(), p2)
					var _617_v24 _dafny.Map
					_ = _617_v24
					_617_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_613_v20, (_dafny.Zero).Minus(_610_v19))
					var _618_v25 D0
					_ = _618_v25
					_618_v25 = Companion_D0_.Create_DC1_(_this.F13(), p2)
					var _619_v26 _dafny.Array
					_ = _619_v26
					var _nwElement0_14 bool = _this.F13()
					_ = _nwElement0_14
					var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(5))
					_ = _nw76
					(_nw76).ArraySet1(_nwElement0_14, 0)
					(_nw76).ArraySet1(!_dafny.Companion_Sequence_.Contains(p1, Companion_Default___.Fm34(globalState)), 1)
					(_nw76).ArraySet1(!(_617_v24).Contains((_613_v20).Update(_dafny.IntOfUint32(((_614_v21).Dtor_cf12()).Cardinality()), Companion_Default___.Abs((_dafny.Zero).Minus((_616_v23).Cardinality())))), 2)
					(_nw76).ArraySet1((_this).Fm4(_this.F13(), _618_v25, _610_v19, _616_v23, globalState), 3)
					(_nw76).ArraySet1(_this.F13(), 4)
					_619_v26 = _nw76
					var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))
					_ = _index115
					(_619_v26).ArraySet1(true, (_index115).Int())
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))
					_ = _index116
					(_619_v26).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}((func(_620_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_621_i2 _dafny.Int) _dafny.CodePoint {
							return _620_v17
						}
					})(_608_v17))), _dafny.UnicodeSeqOfUtf8Bytes("k")), (_index116).Int())
					var _622_v27 _dafny.Sequence
					_ = _622_v27
					_622_v27 = _dafny.UnicodeSeqOfUtf8Bytes("y")
					_622_v27 = _622_v27
					if (_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool) {
						var _623_v28 _dafny.Array
						_ = _623_v28
						var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
						_ = _nw77
						_623_v28 = _nw77
						var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
						_ = _nw78
						_623_v28 = _nw78
						_610_v19 = (Companion_Default___.Fm31(globalState)).Times((_dafny.IntOfUint32((_622_v27).Cardinality())).Times(Companion_Default___.Fm31(globalState)))
						var _624_v29 _dafny.Sequence
						_ = _624_v29
						_624_v29 = _dafny.SeqOf(!((_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool)), (_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool))
						var _625_v30 _dafny.Set
						_ = _625_v30
						_625_v30 = _dafny.SetOf((_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool))
						var _626_v31 _dafny.Sequence
						_ = _626_v31
						_626_v31 = _dafny.SeqOf((_625_v30).Cardinality(), _610_v19)
						var _627_v32 D1
						_ = _627_v32
						_627_v32 = Companion_D1_.Create_DC4_(_619_v26, p2, _dafny.CodePoint('y'))
						var _628_v33 _dafny.MultiSet
						_ = _628_v33
						_628_v33 = _dafny.MultiSetOf((_627_v32).Dtor_cf7())
						var _629_v34 _dafny.Array
						_ = _629_v34
						var _nwElement0_15 _dafny.Int = (_dafny.MultiSetFromSeq(_624_v29)).Cardinality()
						_ = _nwElement0_15
						var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(15))
						_ = _nw79
						(_nw79).ArraySet1(_nwElement0_15, 0)
						(_nw79).ArraySet1(_610_v19, 1)
						(_nw79).ArraySet1(_dafny.IntOfInt64(300), 2)
						(_nw79).ArraySet1(p0, 3)
						(_nw79).ArraySet1(p0, 4)
						(_nw79).ArraySet1(_dafny.IntOfUint32((_624_v29).Cardinality()), 5)
						(_nw79).ArraySet1((_607_v16).Cardinality(), 6)
						(_nw79).ArraySet1(p2, 7)
						(_nw79).ArraySet1(p2, 8)
						(_nw79).ArraySet1((_626_v31).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_628_v33).Contains(_608_v17) {
								return (_628_v33).Multiplicity(_608_v17)
							}
							return _610_v19
						})(), _dafny.IntOfUint32((_626_v31).Cardinality()))).Uint32()).(_dafny.Int), 9)
						(_nw79).ArraySet1(_610_v19, 10)
						(_nw79).ArraySet1(p0, 11)
						(_nw79).ArraySet1((func() _dafny.Int {
							if (_613_v20).Contains(_610_v19) {
								return (_613_v20).Multiplicity(_610_v19)
							}
							return p0
						})(), 12)
						(_nw79).ArraySet1(_dafny.IntOfInt64(304), 13)
						(_nw79).ArraySet1(p2, 14)
						_629_v34 = _nw79
						var _630_v35 *C1
						_ = _630_v35
						var _nw80 *C1 = New_C1_()
						_ = _nw80
						_nw80.Ctor__(_this.F13(), _622_v27, _629_v34)
						_630_v35 = _nw80
						_610_v19 = p0
						var _631_v36 _dafny.Map
						_ = _631_v36
						_631_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _608_v17)
						var _632_v37 _dafny.Array
						_ = _632_v37
						var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
						_ = _nw81
						_632_v37 = _nw81
						var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_632_v37), 0))
						_ = _index117
						(_632_v37).ArraySet1(_624_v29, (_index117).Int())
						var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_632_v37), 0))
						_ = _index118
						var _rhs82 _dafny.Int = _dafny.IntOfInt64(-25)
						_ = _rhs82
						var _rhs83 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(p1, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("krd"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("krd")).Cardinality()))).Uint32(), _dafny.CodePoint('o'))), _608_v17)
						_ = _rhs83
						var _rhs84 _dafny.Sequence = _626_v31
						_ = _rhs84
						var _rhs85 _dafny.Sequence = _dafny.SeqOf((_this.F13()) == (_this.F13()), (_630_v35).F17())
						_ = _rhs85
						var _rhs86 _dafny.Sequence = _622_v27
						_ = _rhs86
						var _lhs51 _dafny.Array = _632_v37
						_ = _lhs51
						var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_632_v37), 0))
						_ = _lhs52
						_610_v19 = _rhs82
						_631_v36 = _rhs83
						_626_v31 = _rhs84
						(_lhs51).ArraySet1(_rhs85, (_lhs52).Int())
						_622_v27 = _rhs86
					} else {
						var _633_v38 _dafny.Array
						_ = _633_v38
						var _nwElement0_16 _dafny.Int = p2
						_ = _nwElement0_16
						var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(3))
						_ = _nw82
						(_nw82).ArraySet1(_nwElement0_16, 0)
						(_nw82).ArraySet1(_610_v19, 1)
						(_nw82).ArraySet1(_dafny.IntOfInt64(149), 2)
						_633_v38 = _nw82
						var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_633_v38), 0))
						_ = _index119
						(_633_v38).ArraySet1(_610_v19, (_index119).Int())
						var _634_v39 _dafny.Sequence
						_ = _634_v39
						_634_v39 = _dafny.SeqOf(_610_v19, p0, p0)
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_633_v38), 0))
						_ = _index120
						(_633_v38).ArraySet1((func() _dafny.Int {
							if (_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool) {
								return (p2).Times(p2)
							}
							return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_634_v39, _634_v39)).Cardinality())
						})(), (_index120).Int())
						var _635_v40 *C0
						_ = _635_v40
						var _nw83 *C0 = New_C0_()
						_ = _nw83
						_nw83.Ctor__(_610_v19)
						_635_v40 = _nw83
						var _636_v41 _dafny.Array
						_ = _636_v41
						var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
						_ = _nw84
						_636_v41 = _nw84
						_636_v41 = _636_v41
						var _rhs87 _dafny.CodePoint = _608_v17
						_ = _rhs87
						var _rhs88 bool = _this.F13()
						_ = _rhs88
						var _lhs53 *C2 = _this
						_ = _lhs53
						_608_v17 = _rhs87
						_lhs53.F13_set_(_rhs88)
						_622_v27 = (func() _dafny.Sequence {
							if (func() bool {
								if (_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool) {
									return (_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool)
								}
								return (_619_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_619_v26), 0))).Int()).(bool)
							})() {
								return p1
							}
							return p1
						})()
					}
					var _637_v42 _dafny.Map
					_ = _637_v42
					_637_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_622_v27, _619_v26)
					_637_v42 = (_637_v42).Update(_622_v27, _619_v26)
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _hi5 _dafny.Int = _dafny.IntOfInt64(151)
		_ = _hi5
		for _638_i3 := (_dafny.Zero).Minus(((_591_v0).Difference(_591_v0)).Cardinality()); _638_i3.Cmp(_hi5) < 0; _638_i3 = _638_i3.Plus(_dafny.One) {
			var _639_v43 _dafny.Sequence
			_ = _639_v43
			_639_v43 = _dafny.SeqOf(_this.F13())
			var _640_v44 _dafny.Sequence
			_ = _640_v44
			_640_v44 = _dafny.SeqOf((_639_v43).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_639_v43).Cardinality()))).Uint32()).(bool))
			var _641_v45 _dafny.Map
			_ = _641_v45
			_641_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), _this.F13())
			(_this).F13_set_(((_641_v45).Contains(_639_v43)) && (_this.F13()))
			var _642_v46 _dafny.Array
			_ = _642_v46
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw85
			_642_v46 = _nw85
			var _643_v47 *C1
			_ = _643_v47
			var _nw86 *C1 = New_C1_()
			_ = _nw86
			_nw86.Ctor__(_this.F13(), _dafny.UnicodeSeqOfUtf8Bytes("gwxskgni"), _642_v46)
			_643_v47 = _nw86
			var _644_v48 D0
			_ = _644_v48
			_644_v48 = Companion_D0_.Create_DC0_(_this.F13())
			var _pat_let_tv16 = _643_v47
			_ = _pat_let_tv16
			var _pat_let_tv17 = _643_v47
			_ = _pat_let_tv17
			var _645_v49 _dafny.Array
			_ = _645_v49
			var _nwElement0_17 D0 = func(_pat_let15_0 D0) D0 {
				return func(_646_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let16_0 bool) D0 {
						return func(_647_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_647_dt__update_hcf0_h0)
						}(_pat_let16_0)
					}(false)
				}(_pat_let15_0)
			}(_644_v48)
			_ = _nwElement0_17
			var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(6))
			_ = _nw87
			(_nw87).ArraySet1(_nwElement0_17, 0)
			(_nw87).ArraySet1(func(_pat_let17_0 D0) D0 {
				return func(_648_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let18_0 bool) D0 {
						return func(_649_dt__update_hcf0_h1 bool) D0 {
							return Companion_D0_.Create_DC0_(_649_dt__update_hcf0_h1)
						}(_pat_let18_0)
					}((_pat_let_tv16).F17())
				}(_pat_let17_0)
			}(_644_v48), 1)
			(_nw87).ArraySet1(_644_v48, 2)
			(_nw87).ArraySet1(func(_pat_let19_0 D0) D0 {
				return func(_650_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let20_0 bool) D0 {
						return func(_651_dt__update_hcf0_h2 bool) D0 {
							return Companion_D0_.Create_DC0_(_651_dt__update_hcf0_h2)
						}(_pat_let20_0)
					}((_pat_let_tv17).F17())
				}(_pat_let19_0)
			}(_644_v48), 3)
			(_nw87).ArraySet1(_644_v48, 4)
			(_nw87).ArraySet1(Companion_D0_.Create_DC0_(_this.F13()), 5)
			_645_v49 = _nw87
			(_643_v47).M5(_645_v49, globalState)
			var _652_v50 _dafny.Map
			_ = _652_v50
			_652_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_643_v47).F17(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_638_i3, _638_i3))
			var _653_v51 _dafny.Map
			_ = _653_v51
			_653_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_639_v43), !(_652_v50).Contains((_643_v47).F17()))
			var _654_v52 _dafny.MultiSet
			_ = _654_v52
			_654_v52 = _dafny.MultiSetOf(true)
			var _655_v53 _dafny.Array
			_ = _655_v53
			var _nwElement0_18 bool = (_643_v47).F17()
			_ = _nwElement0_18
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(20))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_18, 0)
			(_nw88).ArraySet1(_this.F13(), 1)
			(_nw88).ArraySet1(_this.F13(), 2)
			(_nw88).ArraySet1((_643_v47).F17(), 3)
			(_nw88).ArraySet1(_this.F13(), 4)
			(_nw88).ArraySet1(!(_this.F13()), 5)
			(_nw88).ArraySet1(!(_this.F13()), 6)
			(_nw88).ArraySet1(!((_654_v52).IsDisjointFrom(_654_v52)), 7)
			(_nw88).ArraySet1(!((_643_v47).F17()), 8)
			(_nw88).ArraySet1((_643_v47).F17(), 9)
			(_nw88).ArraySet1(_this.F13(), 10)
			(_nw88).ArraySet1((_643_v47).F17(), 11)
			(_nw88).ArraySet1((p0).Cmp(p0) == 0, 12)
			(_nw88).ArraySet1(true, 13)
			(_nw88).ArraySet1(_this.F13(), 14)
			(_nw88).ArraySet1(!(_this.F13()) || ((_643_v47).F17()), 15)
			(_nw88).ArraySet1((p2).Cmp(p2) > 0, 16)
			(_nw88).ArraySet1(_this.F13(), 17)
			(_nw88).ArraySet1(_this.F13(), 18)
			(_nw88).ArraySet1(!(_this.F13()), 19)
			_655_v53 = _nw88
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_655_v53), 0))
			_ = _index121
			(_655_v53).ArraySet1((_643_v47).Fm6(globalState), (_index121).Int())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_655_v53), 0))
			_ = _index122
			(_655_v53).ArraySet1(_this.F13(), (_index122).Int())
			var _656_v54 _dafny.Sequence
			_ = _656_v54
			_656_v54 = _dafny.SeqOf(_638_i3, p2)
			var _657_v55 _dafny.Map
			_ = _657_v55
			_657_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_643_v47).Fm7(_638_i3, (_dafny.Zero).Minus((_656_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_643_v47.F18).Cardinality()), _dafny.IntOfUint32((_656_v54).Cardinality()))).Uint32()).(_dafny.Int)), _this.F13(), globalState), _dafny.IntOfUint32((p1).Cardinality()))
			var _658_v56 _dafny.Set
			_ = _658_v56
			_658_v56 = _dafny.SetOf((_643_v47).Fm14(_638_i3, (_643_v47).F17(), (_657_v55).Cardinality(), p2, globalState))
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_655_v53), 0))
			_ = _index123
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_655_v53), 0))
			_ = _index124
			var _rhs89 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_this.F13(), (_644_v48).Dtor_cf0(), (_643_v47).F17()), (_643_v47).F17())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_640_v44), (_643_v47).F17()))
			_ = _rhs89
			var _rhs90 bool = true
			_ = _rhs90
			var _rhs91 bool = (_658_v56).IsSubsetOf(_658_v56)
			_ = _rhs91
			var _rhs92 bool = !((_643_v47).F17())
			_ = _rhs92
			var _rhs93 bool = !_dafny.Companion_Sequence_.Equal(p1, _643_v47.F18)
			_ = _rhs93
			var _lhs54 _dafny.Array = _655_v53
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_655_v53), 0))
			_ = _lhs55
			var _lhs56 _dafny.Array = _655_v53
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_655_v53), 0))
			_ = _lhs57
			_653_v51 = _rhs89
			r1 = _rhs90
			(_lhs54).ArraySet1(_rhs91, (_lhs55).Int())
			(_lhs56).ArraySet1(_rhs92, (_lhs57).Int())
			r1 = _rhs93
		}
		var _659_v57 _dafny.Set
		_ = _659_v57
		_659_v57 = _dafny.SetOf(_this.F13(), !(_this.F13()))
		var _660_v58 _dafny.Map
		_ = _660_v58
		_660_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())
		var _661_v59 _dafny.Array
		_ = _661_v59
		var _nwElement0_19 _dafny.Int = p0
		_ = _nwElement0_19
		var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(16))
		_ = _nw89
		(_nw89).ArraySet1(_nwElement0_19, 0)
		(_nw89).ArraySet1(p2, 1)
		(_nw89).ArraySet1((_dafny.Zero).Minus(p2), 2)
		(_nw89).ArraySet1(p0, 3)
		(_nw89).ArraySet1(p0, 4)
		(_nw89).ArraySet1((_659_v57).Cardinality(), 5)
		(_nw89).ArraySet1(p2, 6)
		(_nw89).ArraySet1(p0, 7)
		(_nw89).ArraySet1(p2, 8)
		(_nw89).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 9)
		(_nw89).ArraySet1(_610_v19, 10)
		(_nw89).ArraySet1(p0, 11)
		(_nw89).ArraySet1(_610_v19, 12)
		(_nw89).ArraySet1((_660_v58).Cardinality(), 13)
		(_nw89).ArraySet1(_610_v19, 14)
		(_nw89).ArraySet1(p2, 15)
		_661_v59 = _nw89
		var _662_v60 D4
		_ = _662_v60
		_662_v60 = Companion_D4_.Create_DC10_((func() _dafny.Array {
			if _this.F13() {
				return _661_v59
			}
			return _661_v59
		})())
		var _source12 D4 = _662_v60
		_ = _source12
		if _source12.Is_DC11() {
			var _663___mcc_h0 bool = _source12.Get_().(D4_DC11).Cf21
			_ = _663___mcc_h0
			var _664_cf21 bool = _663___mcc_h0
			_ = _664_cf21
			var _665_v61 *C0
			_ = _665_v61
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(822), Companion_Default___.Fm31(globalState)))
			_665_v61 = _nw90
			var _666_v62 _dafny.Map
			_ = _666_v62
			_666_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm35(p0, globalState), false)
			var _667_v63 _dafny.Array
			_ = _667_v63
			var _nwElement0_20 _dafny.Array = _661_v59
			_ = _nwElement0_20
			var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(14))
			_ = _nw91
			(_nw91).ArraySet1(_nwElement0_20, 0)
			(_nw91).ArraySet1(_661_v59, 1)
			(_nw91).ArraySet1(_661_v59, 2)
			(_nw91).ArraySet1(_661_v59, 3)
			(_nw91).ArraySet1(_661_v59, 4)
			(_nw91).ArraySet1(_661_v59, 5)
			(_nw91).ArraySet1(_661_v59, 6)
			(_nw91).ArraySet1(_661_v59, 7)
			(_nw91).ArraySet1(_661_v59, 8)
			(_nw91).ArraySet1(_661_v59, 9)
			(_nw91).ArraySet1(_661_v59, 10)
			(_nw91).ArraySet1(_661_v59, 11)
			(_nw91).ArraySet1(_661_v59, 12)
			(_nw91).ArraySet1(_661_v59, 13)
			_667_v63 = _nw91
			var _668_v64 *C1
			_ = _668_v64
			var _nw92 *C1 = New_C1_()
			_ = _nw92
			_nw92.Ctor__(_this.F13(), p1, _661_v59)
			_668_v64 = _nw92
			var _669_v65 _dafny.Map
			_ = _669_v65
			_669_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_667_v63, _dafny.SetOf(_668_v64))
			var _670_v66 _dafny.Set
			_ = _670_v66
			_670_v66 = _dafny.SetOf(_668_v64)
			var _671_v67 _dafny.Map
			_ = _671_v67
			_671_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Minus((_666_v62).Cardinality()), (func() _dafny.Int {
				if !(_664_cf21) {
					return (_dafny.Zero).Minus(((func() _dafny.Set {
						if (_669_v65).Contains(_667_v63) {
							return (_669_v65).Get(_667_v63).(_dafny.Set)
						}
						return _670_v66
					})()).Cardinality())
				}
				return _610_v19
			})())
			_671_v67 = _671_v67
			var _672_v68 _dafny.Array
			_ = _672_v68
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw93
			_672_v68 = _nw93
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_672_v68), 0))
			_ = _index125
			(_672_v68).ArraySet1((func() bool {
				if _this.F13() {
					return false
				}
				return _664_cf21
			})(), (_index125).Int())
			var _673_v69 _dafny.MultiSet
			_ = _673_v69
			_673_v69 = _dafny.MultiSetOf(_672_v68, _672_v68, _672_v68)
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_672_v68), 0))
			_ = _index126
			(_672_v68).ArraySet1(!(((_673_v69).IsDisjointFrom(_673_v69)) || (_664_cf21)), (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_672_v68), 0))
			_ = _index127
			(_672_v68).ArraySet1((_668_v64).F17(), (_index127).Int())
		} else {
			var _674___mcc_h1 _dafny.Array = _source12.Get_().(D4_DC10).Cf20
			_ = _674___mcc_h1
			var _675_cf20 _dafny.Array = _674___mcc_h1
			_ = _675_cf20
			var _676_v70 _dafny.Sequence
			_ = _676_v70
			_676_v70 = _dafny.SeqOf(_this.F13())
			var _677_v71 *C1
			_ = _677_v71
			var _nw94 *C1 = New_C1_()
			_ = _nw94
			_nw94.Ctor__((!((_676_v70).Select((Companion_Default___.SafeIndex(_610_v19, _dafny.IntOfUint32((_676_v70).Cardinality()))).Uint32()).(bool))) == (_this.F13()), p1, _675_cf20)
			_677_v71 = _nw94
			if _this.F13() {
				var _678_v72 _dafny.Sequence
				_ = _678_v72
				_678_v72 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("cqgr"), (Companion_Default___.SafeIndex(_610_v19, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cqgr")).Cardinality()))).Uint32(), _608_v17), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-754))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_679_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_680_i4 _dafny.Int) _dafny.CodePoint {
						return _679_v17
					}
				})(_608_v17))), _dafny.UnicodeSeqOfUtf8Bytes("uthxbo"))
				var _681_v73 _dafny.Map
				_ = _681_v73
				_681_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_677_v71).F17(), _678_v72)
				var _682_v74 _dafny.Sequence
				_ = _682_v74
				_682_v74 = _dafny.SeqOf((_659_v57).Cardinality())
				r1 = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
					if _this.F13() {
						return (func() _dafny.Sequence {
							if (_681_v73).Contains((_677_v71).F17()) {
								return (_681_v73).Get((_677_v71).F17()).(_dafny.Sequence)
							}
							return _678_v72
						})()
					}
					return _678_v72
				})(), _dafny.SeqOf(Companion_Default___.Fm32(_677_v71.F18, (_677_v71).F17(), p2, (_682_v74).Select((Companion_Default___.SafeIndex((_607_v16).Cardinality(), _dafny.IntOfUint32((_682_v74).Cardinality()))).Uint32()).(_dafny.Int), globalState)))
				_610_v19 = (_dafny.Zero).Minus((_610_v19).Times(_610_v19))
				_591_v0 = (func() _dafny.Set {
					if (_677_v71).F17() {
						return _591_v0
					}
					return _591_v0
				})()
				_610_v19 = (_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality())).Times((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_609_v18).Cardinality(), _dafny.IntOfInt64(416))))
				var _683_v75 *C1
				_ = _683_v75
				var _nw95 *C1 = New_C1_()
				_ = _nw95
				_nw95.Ctor__((_677_v71).F17(), p1, _675_cf20)
				_683_v75 = _nw95
			} else {
				var _684_v76 _dafny.Array
				_ = _684_v76
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_10
				var _nw96 _dafny.Array
				_ = _nw96
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw96 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Sequence = (func(_685_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_686_i5 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg46 _dafny.Int) interface{} {
									return coer46(arg46)
								}
							}((func(_687_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_688_i6 _dafny.Int) _dafny.CodePoint {
									return _687_v17
								}
							})(_685_v17)))
						}
					})(_608_v17)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw96 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw96).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw96).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_684_v76 = _nw96
				var _689_v77 _dafny.Sequence
				_ = _689_v77
				_689_v77 = _dafny.SeqOf(_684_v76, _684_v76, _684_v76, _684_v76, _684_v76)
				_684_v76 = (_689_v77).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_v19, true)).Cardinality(), _dafny.IntOfUint32((_689_v77).Cardinality()))).Uint32()).(_dafny.Array)
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_675_cf20), 0))
				_ = _index128
				(_675_cf20).ArraySet1(p0, (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_675_cf20), 0))
				_ = _index129
				(_675_cf20).ArraySet1(Companion_Default___.SafeModuloInt(p2, p0), (_index129).Int())
				r2 = _this.F13()
				_610_v19 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.IntOfInt64(-336)).Minus((_675_cf20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_675_cf20), 0))).Int()).(_dafny.Int))))
				var _690_v78 _dafny.Map
				_ = _690_v78
				_690_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F13())
				_690_v78 = (_690_v78).Merge(_690_v78)
			}
			r2 = (_677_v71).F17()
			var _source13 D4 = (func() D4 {
				if (p2).Cmp(_610_v19) > 0 {
					return _662_v60
				}
				return Companion_D4_.Create_DC10_(_675_cf20)
			})()
			_ = _source13
			if _source13.Is_DC11() {
				var _691___mcc_h2 bool = _source13.Get_().(D4_DC11).Cf21
				_ = _691___mcc_h2
				var _692_cf21 bool = _691___mcc_h2
				_ = _692_cf21
				var _693_v79 _dafny.Array
				_ = _693_v79
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw97
				_693_v79 = _nw97
				var _694_v80 _dafny.Sequence
				_ = _694_v80
				_694_v80 = _dafny.SeqOf(p0)
				var _695_v81 _dafny.Sequence
				_ = _695_v81
				_695_v81 = _dafny.SeqOf((_694_v80).Select((Companion_Default___.SafeIndex(_610_v19, _dafny.IntOfUint32((_694_v80).Cardinality()))).Uint32()).(_dafny.Int))
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_693_v79), 0))
				_ = _index130
				(_693_v79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_695_v81, _694_v80), (_index130).Int())
				var _696_v82 _dafny.Array
				_ = _696_v82
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw98
				_696_v82 = _nw98
				var _697_v83 D0
				_ = _697_v83
				_697_v83 = Companion_D0_.Create_DC1_((_677_v71).F17(), p0)
				var _698_v84 D5
				_ = _698_v84
				_698_v84 = Companion_D5_.Create_DC14_(_692_cf21)
				var _699_v85 D4
				_ = _699_v85
				_699_v85 = Companion_D4_.Create_DC11_(_this.F13())
				var _700_v86 _dafny.Array
				_ = _700_v86
				var _nwElement0_21 bool = (_677_v71).F17()
				_ = _nwElement0_21
				var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(26))
				_ = _nw99
				(_nw99).ArraySet1(_nwElement0_21, 0)
				(_nw99).ArraySet1(_this.F13(), 1)
				(_nw99).ArraySet1(_692_cf21, 2)
				(_nw99).ArraySet1((_677_v71).F17(), 3)
				(_nw99).ArraySet1(false, 4)
				(_nw99).ArraySet1(_692_cf21, 5)
				(_nw99).ArraySet1(false, 6)
				(_nw99).ArraySet1(!(!((_677_v71).F17())), 7)
				(_nw99).ArraySet1(_692_cf21, 8)
				(_nw99).ArraySet1(false, 9)
				(_nw99).ArraySet1((_677_v71).F17(), 10)
				(_nw99).ArraySet1((_this).Fm4((_677_v71).F17(), _697_v83, (Companion_Default___.Fm36(globalState)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(106), _610_v19), globalState), 11)
				(_nw99).ArraySet1(_692_cf21, 12)
				(_nw99).ArraySet1(!((_677_v71).Fm14(p0, (_677_v71).F17(), _610_v19, (_677_v71).Fm7(p2, _dafny.IntOfUint32((_677_v71.F18).Cardinality()), _this.F13(), globalState), globalState)), 13)
				(_nw99).ArraySet1((_698_v84).Dtor_cf23(), 14)
				(_nw99).ArraySet1(_this.F13(), 15)
				(_nw99).ArraySet1(_this.F13(), 16)
				(_nw99).ArraySet1((_676_v70).Select((Companion_Default___.SafeIndex(_610_v19, _dafny.IntOfUint32((_676_v70).Cardinality()))).Uint32()).(bool), 17)
				(_nw99).ArraySet1((_677_v71).F17(), 18)
				(_nw99).ArraySet1((_677_v71).F17(), 19)
				(_nw99).ArraySet1((_677_v71).F17(), 20)
				(_nw99).ArraySet1(!((_677_v71).F17()), 21)
				(_nw99).ArraySet1((_677_v71).Fm6(globalState), 22)
				(_nw99).ArraySet1((_677_v71).F17(), 23)
				(_nw99).ArraySet1(true, 24)
				(_nw99).ArraySet1((_699_v85).Dtor_cf21(), 25)
				_700_v86 = _nw99
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_696_v82), 0))
				_ = _index131
				(_696_v82).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _700_v86), (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_700_v86), 0))
				_ = _index132
				(_700_v86).ArraySet1((_677_v71).F17(), (_index132).Int())
				var _701_v87 _dafny.MultiSet
				_ = _701_v87
				_701_v87 = _dafny.MultiSetOf((_dafny.Zero).Minus(p2))
				var _702_v88 _dafny.Map
				_ = _702_v88
				_702_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_701_v87).Cardinality()).Times(p2), _700_v86)
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_693_v79), 0))
				_ = _index133
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_696_v82), 0))
				_ = _index134
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_700_v86), 0))
				_ = _index135
				var _rhs94 _dafny.Sequence = Companion_Default___.Fm35((_677_v71).Fm15(globalState), globalState)
				_ = _rhs94
				var _rhs95 _dafny.Array = _661_v59
				_ = _rhs95
				var _rhs96 _dafny.Map = _702_v88
				_ = _rhs96
				var _rhs97 bool = (_this.F13()) == (_this.F13())
				_ = _rhs97
				var _rhs98 _dafny.Array = _675_cf20
				_ = _rhs98
				var _lhs58 _dafny.Array = _693_v79
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_693_v79), 0))
				_ = _lhs59
				var _lhs60 _dafny.Array = _696_v82
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_696_v82), 0))
				_ = _lhs61
				var _lhs62 _dafny.Array = _700_v86
				_ = _lhs62
				var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_700_v86), 0))
				_ = _lhs63
				(_lhs58).ArraySet1(_rhs94, (_lhs59).Int())
				_661_v59 = _rhs95
				(_lhs60).ArraySet1(_rhs96, (_lhs61).Int())
				(_lhs62).ArraySet1(_rhs97, (_lhs63).Int())
				_675_cf20 = _rhs98
				_610_v19 = (func() _dafny.Int {
					if (_607_v16).Contains(_this.F13()) {
						return (_607_v16).Get(_this.F13()).(_dafny.Int)
					}
					return p0
				})()
				var _703_v89 _dafny.Map
				_ = _703_v89
				_703_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), _675_cf20)
				var _704_v90 _dafny.Map
				_ = _704_v90
				_704_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (p0).Cmp((_703_v89).Cardinality()) < 0)
				var _rhs99 bool = (_677_v71).Fm6(globalState)
				_ = _rhs99
				var _rhs100 _dafny.Set = _659_v57
				_ = _rhs100
				var _rhs101 _dafny.Int = (p2).Plus((func() _dafny.Int {
					if false {
						return p0
					}
					return (_677_v71).Fm7(_610_v19, p2, (_677_v71).F17(), globalState)
				})())
				_ = _rhs101
				var _rhs102 bool = (func() bool {
					if (_704_v90).Contains(Companion_Default___.SafeDivisionInt(p0, _610_v19)) {
						return (_704_v90).Get(Companion_Default___.SafeDivisionInt(p0, _610_v19)).(bool)
					}
					return (_701_v87).IsSubsetOf(_701_v87)
				})()
				_ = _rhs102
				var _rhs103 bool = _this.F13()
				_ = _rhs103
				r2 = _rhs99
				_659_v57 = _rhs100
				_610_v19 = _rhs101
				_692_cf21 = _rhs102
				r1 = _rhs103
				_610_v19 = _dafny.IntOfInt64(976)
			} else {
				var _705___mcc_h3 _dafny.Array = _source13.Get_().(D4_DC10).Cf20
				_ = _705___mcc_h3
				var _706_cf20 _dafny.Array = _705___mcc_h3
				_ = _706_cf20
				_610_v19 = (((_660_v58).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_677_v71).F17(), (_677_v71).F17()))).Merge((_660_v58).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_677_v71).F17(), (_677_v71).F17())))).Cardinality()
				_610_v19 = _610_v19
				var _707_v91 _dafny.MultiSet
				_ = _707_v91
				_707_v91 = _dafny.MultiSetOf(_this.F13())
				_610_v19 = Companion_Default___.SafeDivisionInt((p2).Times(((_707_v91).Update((_677_v71).F17(), Companion_Default___.Abs(p0))).Cardinality()), _610_v19)
				var _708_v92 _dafny.Sequence
				_ = _708_v92
				_708_v92 = _dafny.SeqOf(_677_v71)
				var _709_v93 _dafny.Map
				_ = _709_v93
				_709_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(581))
				var _710_v94 _dafny.Map
				_ = _710_v94
				_710_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_709_v93).Contains(p0) {
						return (_709_v93).Get(p0).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_709_v93).Contains(_dafny.IntOfInt64(26)) {
							return (_709_v93).Get(_dafny.IntOfInt64(26)).(_dafny.Int)
						}
						return p2
					})()
				})(), _dafny.UnicodeSeqOfUtf8Bytes("t"))
				var _711_v95 T1
				_ = _711_v95
				var _nw100 *C1 = New_C1_()
				_ = _nw100
				_nw100.Ctor__(((_677_v71).Fm15(globalState)).Cmp(_dafny.IntOfUint32((_708_v92).Cardinality())) <= 0, (func() _dafny.Sequence {
					if (_710_v94).Contains(_610_v19) {
						return (_710_v94).Get(_610_v19).(_dafny.Sequence)
					}
					return p1
				})(), _706_cf20)
				_711_v95 = _nw100
				var _712_v96 _dafny.Map
				_ = _712_v96
				_712_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_v19, _607_v16)
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen(((_711_v95).F14()), 0))
				_ = _index136
				((_711_v95).F14()).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_677_v71.F18, _677_v71.F18, _dafny.UnicodeSeqOfUtf8Bytes("sp"), _677_v71.F18, _677_v71.F18), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf(_677_v71.F18, _677_v71.F18, _dafny.UnicodeSeqOfUtf8Bytes("sp"), _677_v71.F18, _677_v71.F18)).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_677_v71.F18, _677_v71.F18, _dafny.UnicodeSeqOfUtf8Bytes("sp"), _677_v71.F18, _677_v71.F18), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf(_677_v71.F18, _677_v71.F18, _dafny.UnicodeSeqOfUtf8Bytes("sp"), _677_v71.F18, _677_v71.F18)).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), _677_v71.F18)).Cardinality()), (_index136).Int())
				var _713_v97 D7
				_ = _713_v97
				_713_v97 = Companion_D7_.Create_DC19_(_711_v95)
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen(((_711_v95).F14()), 0))
				_ = _index137
				var _rhs104 T1 = (_713_v97).Dtor_cf30()
				_ = _rhs104
				var _rhs105 _dafny.Map = Companion_Default___.Fm37((_677_v71).F17(), globalState)
				_ = _rhs105
				var _rhs106 _dafny.Int = Companion_Default___.SafeDivisionInt(p2, (_dafny.IntOfUint32((p1).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-849))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_714_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_715_i7 _dafny.Int) _dafny.CodePoint {
						return _714_v17
					}
				})(_608_v17)))).Cardinality())))
				_ = _rhs106
				var _lhs64 _dafny.Array = (_711_v95).F14()
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen(((_711_v95).F14()), 0))
				_ = _lhs65
				_711_v95 = _rhs104
				_712_v96 = _rhs105
				(_lhs64).ArraySet1(_rhs106, (_lhs65).Int())
			}
		}
		var _716_v98 _dafny.Array
		_ = _716_v98
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_11
		var _nw101 _dafny.Array
		_ = _nw101
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw101 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = func(_717_i8 _dafny.Int) bool {
				return _this.F13()
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw101 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw101).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw101).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_716_v98 = _nw101
		var _718_v99 _dafny.MultiSet
		_ = _718_v99
		_718_v99 = _dafny.MultiSetOf(_716_v98, _716_v98)
		(_this).F13_set_(!((_718_v99).Union(_718_v99)).Equals((_dafny.MultiSetOf(_716_v98)).Intersection(_718_v99)))
		var _719_v100 _dafny.Map
		_ = _719_v100
		_719_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _716_v98)
		r0 = (_719_v100).Merge(_719_v100)
		r1 = _this.F13()
		var _720_v101 D2
		_ = _720_v101
		_720_v101 = Companion_D2_.Create_DC7_(_dafny.SeqOf(_608_v17), (_659_v57).Cardinality(), _dafny.IntOfInt64(749), _this.F13())
		r2 = ((func() D2 {
			if _this.F13() {
				return _720_v101
			}
			return _720_v101
		})()).Dtor_cf15()
		var _721_v102 _dafny.Sequence
		_ = _721_v102
		_721_v102 = _dafny.SeqOf(_this.F13())
		r3 = (func() _dafny.Sequence {
			if _dafny.Companion_Sequence_.Contains(_721_v102, _this.F13()) {
				return _721_v102
			}
			return _dafny.Companion_Sequence_.Concatenate(_721_v102, _721_v102)
		})()
		return r0, r1, r2, r3
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f13 bool
	_f19 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f13 = false
	_this._f19 = false
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

func (_this *C3) F13() bool {
	return _this._f13
}
func (_this *C3) F13_set_(value bool) {
	_this._f13 = value
}
func (_this *C3) Ctor__(f19 bool, f13 bool) {
	{
		(_this)._f19 = f19
		(_this)._f13 = f13
	}
}
func (_this *C3) Fm4(p0 bool, p1 D0, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		var _source14 D1 = Companion_D1_.Create_DC2_(func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(439), _dafny.IntOfInt64(-49))); ; {
				_compr_30, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _722_v0 _dafny.Int
				_722_v0 = interface{}(_compr_30).(_dafny.Int)
				if ((_dafny.IntOfInt64(439)).Cmp(_722_v0) <= 0) && ((_722_v0).Cmp(_dafny.IntOfInt64(-49)) < 0) {
					_coll30.Add((_722_v0).Plus(_dafny.IntOfInt64(-610)), _dafny.IntOfInt64(836))
				}
			}
			return _coll30.ToMap()
		}())
		_ = _source14
		if _source14.Is_DC3() {
			var _723___mcc_h0 _dafny.Int = _source14.Get_().(D1_DC3).Cf4
			_ = _723___mcc_h0
			var _724_cf4 _dafny.Int = _723___mcc_h0
			_ = _724_cf4
			return true
		} else if _source14.Is_DC4() {
			var _725___mcc_h1 _dafny.Array = _source14.Get_().(D1_DC4).Cf5
			_ = _725___mcc_h1
			var _726___mcc_h2 _dafny.Int = _source14.Get_().(D1_DC4).Cf6
			_ = _726___mcc_h2
			var _727___mcc_h3 _dafny.CodePoint = _source14.Get_().(D1_DC4).Cf7
			_ = _727___mcc_h3
			var _728_cf7 _dafny.CodePoint = _727___mcc_h3
			_ = _728_cf7
			var _729_cf6 _dafny.Int = _726___mcc_h2
			_ = _729_cf6
			var _730_cf5 _dafny.Array = _725___mcc_h1
			_ = _730_cf5
			return _this.F13()
		} else if _source14.Is_DC5() {
			var _731___mcc_h4 _dafny.Sequence = _source14.Get_().(D1_DC5).Cf8
			_ = _731___mcc_h4
			var _732___mcc_h5 _dafny.Array = _source14.Get_().(D1_DC5).Cf9
			_ = _732___mcc_h5
			var _733___mcc_h6 _dafny.Int = _source14.Get_().(D1_DC5).Cf10
			_ = _733___mcc_h6
			var _734_cf10 _dafny.Int = _733___mcc_h6
			_ = _734_cf10
			var _735_cf9 _dafny.Array = _732___mcc_h5
			_ = _735_cf9
			var _736_cf8 _dafny.Sequence = _731___mcc_h4
			_ = _736_cf8
			return false
		} else {
			var _737___mcc_h7 _dafny.Map = _source14.Get_().(D1_DC2).Cf3
			_ = _737___mcc_h7
			var _738_cf3 _dafny.Map = _737___mcc_h7
			_ = _738_cf3
			return _this.F13()
		}
	}
}
func (_this *C3) Fm5(globalState *GlobalState) D1 {
	{
		if _this.F13() {
			return Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F19(), _this.F13())).Cardinality()))
		} else {
			return Companion_D1_.Create_DC3_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F13()), (_this).F19())).Cardinality())
		}
	}
}
func (_this *C3) Fm27(globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('d')
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _hi6 _dafny.Int = p0
		_ = _hi6
		for _739_i0 := p0; _739_i0.Cmp(_hi6) < 0; _739_i0 = _739_i0.Plus(_dafny.One) {
			var _740_v0 _dafny.Sequence
			_ = _740_v0
			_740_v0 = _dafny.SeqOf(_this.F13(), (_this).F19())
			(_this).F13_set_((_740_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_741_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			}))).Cardinality()), _dafny.IntOfUint32((_740_v0).Cardinality()))).Uint32()).(bool))
			var _742_v1 _dafny.Array
			_ = _742_v1
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw102
			_742_v1 = _nw102
			var _743_v2 _dafny.Array
			_ = _743_v2
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw103
			_743_v2 = _nw103
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_742_v1), 0))
			_ = _index138
			(_742_v1).ArraySet1(_743_v2, (_index138).Int())
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_742_v1), 0))
			_ = _index139
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_12
			var _nw104 _dafny.Array
			_ = _nw104
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw104 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) bool = func(_744_i2 _dafny.Int) bool {
					return false
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw104 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw104).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw104).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			(_742_v1).ArraySet1(_nw104, (_index139).Int())
			var _745_v3 _dafny.Sequence
			_ = _745_v3
			_745_v3 = _dafny.UnicodeSeqOfUtf8Bytes("rn")
			_745_v3 = _745_v3
			if !((_this).F19()) {
				var _746_v4 _dafny.Int
				_ = _746_v4
				_746_v4 = _dafny.IntOfInt64(687)
				var _747_v5 _dafny.Set
				_ = _747_v5
				_747_v5 = _dafny.SetOf(_dafny.ArrayCastTo((_742_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_742_v1), 0))).Int())), _743_v2, _dafny.ArrayCastTo((_742_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_742_v1), 0))).Int())), _dafny.ArrayCastTo((_742_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_742_v1), 0))).Int())))
				_746_v4 = ((_747_v5).Difference(_747_v5)).Cardinality()
				var _748_v6 _dafny.Array
				_ = _748_v6
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_13
				var _nw105 _dafny.Array
				_ = _nw105
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw105 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_749_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_750_i3 _dafny.Int) _dafny.Sequence {
							return _749_v3
						}
					})(_745_v3)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw105 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw105).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw105).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_748_v6 = _nw105
				var _751_v7 _dafny.CodePoint
				_ = _751_v7
				_751_v7 = _dafny.CodePoint('n')
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_748_v6), 0))
				_ = _index140
				(_748_v6).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(138))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_752_v3 _dafny.Sequence, _753_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_754_i4 _dafny.Int) _dafny.CodePoint {
						return (_752_v3).Select((Companion_Default___.SafeIndex(_753_p0, _dafny.IntOfUint32((_752_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_745_v3, p0))), (Companion_Default___.SafeIndex(_746_v4, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(138))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_755_v3 _dafny.Sequence, _756_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_757_i4 _dafny.Int) _dafny.CodePoint {
						return (_755_v3).Select((Companion_Default___.SafeIndex(_756_p0, _dafny.IntOfUint32((_755_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_745_v3, p0)))).Cardinality()))).Uint32(), _751_v7), (_index140).Int())
				var _758_v8 _dafny.Sequence
				_ = _758_v8
				_758_v8 = _dafny.SeqOf(_745_v3)
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_748_v6), 0))
				_ = _index141
				(_748_v6).ArraySet1((_758_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_758_v8).Cardinality()))).Uint32()).(_dafny.Sequence), (_index141).Int())
				var _759_v9 _dafny.Map
				_ = _759_v9
				_759_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_739_i0), _dafny.IntOfInt64(-33))
				_746_v4 = ((_759_v9).Cardinality()).Times(_746_v4)
				var _760_v10 _dafny.Map
				_ = _760_v10
				_760_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), p0)
				_760_v10 = (_760_v10).Update(_this.F13(), _dafny.IntOfInt64(-151))
				var _761_v11 _dafny.Array
				_ = _761_v11
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw106
				_761_v11 = _nw106
				var _762_v12 *C1
				_ = _762_v12
				var _nw107 *C1 = New_C1_()
				_ = _nw107
				_nw107.Ctor__((_this).F19(), _745_v3, _761_v11)
				_762_v12 = _nw107
			} else {
				var _763_v13 _dafny.Int
				_ = _763_v13
				_763_v13 = _dafny.IntOfInt64(-760)
				_763_v13 = (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _763_v13)).Cardinality()).Minus(_dafny.IntOfInt64(625)))
				var _764_v14 _dafny.CodePoint
				_ = _764_v14
				_764_v14 = _dafny.CodePoint('j')
				var _765_v15 _dafny.Map
				_ = _765_v15
				_765_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_764_v14, (_this).F19())
				(_this).F13_set_((func() bool {
					if (_765_v15).Contains(_764_v14) {
						return (_765_v15).Get(_764_v14).(bool)
					}
					return (_this).F19()
				})())
				var _766_v16 _dafny.Array
				_ = _766_v16
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_14
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.Int = (func(_767_v13 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_768_i5 _dafny.Int) _dafny.Int {
							return (_768_i5).Plus((_dafny.Zero).Minus(_767_v13))
						}
					})(_763_v13)
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw108 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw108).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw108).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_766_v16 = _nw108
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_766_v16), 0))
				_ = _index142
				(_766_v16).ArraySet1(_dafny.IntOfInt64(-945), (_index142).Int())
				var _769_v17 D0
				_ = _769_v17
				_769_v17 = Companion_D0_.Create_DC1_(_this.F13(), p0)
				var _770_v18 T0
				_ = _770_v18
				var _nw109 *C2 = New_C2_()
				_ = _nw109
				_nw109.Ctor__(_this.F13())
				_770_v18 = _nw109
				var _771_v19 _dafny.Sequence
				_ = _771_v19
				_771_v19 = _dafny.SeqOf(_770_v18)
				var _772_v20 _dafny.Map
				_ = _772_v20
				_772_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v13, Companion_Default___.Fm28(_dafny.IntOfUint32((_771_v19).Cardinality()), _763_v13, _this.F13(), globalState))
				var _773_v21 *C0
				_ = _773_v21
				var _nw110 *C0 = New_C0_()
				_ = _nw110
				_nw110.Ctor__(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_772_v20).Contains(_763_v13) {
						return (_772_v20).Get(_763_v13).(_dafny.Sequence)
					}
					return _745_v3
				})()).Cardinality()))
				_773_v21 = _nw110
				var _774_v22 _dafny.Map
				_ = _774_v22
				_774_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_773_v21).F15(), _dafny.IntOfUint32((_745_v3).Cardinality()))
				var _775_v23 _dafny.Set
				_ = _775_v23
				_775_v23 = _dafny.SetOf((_this).Fm4((_this).F19(), _769_v17, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_v21, _770_v18.F13())).Cardinality(), _774_v22, globalState))
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_766_v16), 0))
				_ = _index143
				(_766_v16).ArraySet1((_775_v23).Cardinality(), (_index143).Int())
				var _776_v24 D5
				_ = _776_v24
				_776_v24 = Companion_D5_.Create_DC14_(!(false))
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_766_v16), 0))
				_ = _index144
				var _rhs107 D5 = _776_v24
				_ = _rhs107
				var _rhs108 _dafny.Int = (_dafny.Zero).Minus(_739_i0)
				_ = _rhs108
				var _rhs109 bool = _this.F13()
				_ = _rhs109
				var _lhs66 _dafny.Array = _766_v16
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_766_v16), 0))
				_ = _lhs67
				var _lhs68 *C3 = _this
				_ = _lhs68
				_776_v24 = _rhs107
				(_lhs66).ArraySet1(_rhs108, (_lhs67).Int())
				_lhs68.F13_set_(_rhs109)
				var _777_v25 _dafny.Map
				_ = _777_v25
				_777_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_776_v24).Dtor_cf23(), (_766_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_766_v16), 0))).Int()).(_dafny.Int))
				var _778_v26 _dafny.MultiSet
				_ = _778_v26
				_778_v26 = _dafny.MultiSetOf(_739_i0, _739_i0)
				_777_v25 = (_777_v25).Update((_778_v26).IsSubsetOf(_778_v26), _763_v13)
			}
		}
		var _hi7 _dafny.Int = p0
		_ = _hi7
		for _779_i6 := p0; _779_i6.Cmp(_hi7) < 0; _779_i6 = _779_i6.Plus(_dafny.One) {
			(_this).F13_set_(_this.F13())
			var _780_v27 D2
			_ = _780_v27
			_780_v27 = Companion_D2_.Create_DC8_(_this.F13(), _this.F13(), _this.F13())
			var _781_v28 _dafny.Sequence
			_ = _781_v28
			_781_v28 = _dafny.SeqOf(_780_v27, _780_v27)
			var _782_v29 _dafny.Map
			_ = _782_v29
			_782_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), false)
			_781_v28 = _dafny.SeqOf(_780_v27, Companion_D2_.Create_DC8_((_this).F19(), _this.F13(), (func() bool {
				if (_782_v29).Contains((_this).F19()) {
					return (_782_v29).Get((_this).F19()).(bool)
				}
				return _this.F13()
			})()), _780_v27)
			var _783_v30 _dafny.Int
			_ = _783_v30
			_783_v30 = _dafny.IntOfInt64(328)
			_783_v30 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_783_v30, _783_v30))
			var _784_v31 _dafny.Sequence
			_ = _784_v31
			_784_v31 = _dafny.UnicodeSeqOfUtf8Bytes("orhdrkah")
			var _785_v32 _dafny.Array
			_ = _785_v32
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw111
			_785_v32 = _nw111
			var _786_v33 D6
			_ = _786_v33
			_786_v33 = Companion_D6_.Create_DC17_((_dafny.Zero).Minus(_779_i6), p0, _785_v32)
			var _787_v34 D6
			_ = _787_v34
			_787_v34 = Companion_D6_.Create_DC18_(_786_v33)
			var _788_v35 _dafny.Map
			_ = _788_v35
			_788_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_787_v34, _this.F13())
			var _789_v36 D1
			_ = _789_v36
			_789_v36 = Companion_D1_.Create_DC5_(_784_v31, _785_v32, (_788_v35).Cardinality())
			var _pat_let_tv18 = _783_v30
			_ = _pat_let_tv18
			_789_v36 = func(_pat_let21_0 D1) D1 {
				return func(_790_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let22_0 _dafny.Sequence) D1 {
						return func(_791_dt__update_hcf8_h0 _dafny.Sequence) D1 {
							return func(_pat_let23_0 _dafny.Int) D1 {
								return func(_792_dt__update_hcf10_h0 _dafny.Int) D1 {
									return Companion_D1_.Create_DC5_(_791_dt__update_hcf8_h0, (_790_dt__update__tmp_h0).Dtor_cf9(), _792_dt__update_hcf10_h0)
								}(_pat_let23_0)
							}(_pat_let_tv18)
						}(_pat_let22_0)
					}(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("f"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("f")).Cardinality()))).Uint32(), _dafny.CodePoint('s')))
				}(_pat_let21_0)
			}(_789_v36)
		}
		var _793_v37 D6
		_ = _793_v37
		_793_v37 = Companion_D6_.Create_DC16_(Companion_Default___.Fm38(_this.F13(), globalState))
		var _794_v38 D6
		_ = _794_v38
		_794_v38 = Companion_D6_.Create_DC18_(_793_v37)
		var _795_v39 D6
		_ = _795_v39
		_795_v39 = Companion_D6_.Create_DC18_(_793_v37)
		_795_v39 = _795_v39
		var _796_v40 _dafny.CodePoint
		_ = _796_v40
		_796_v40 = _dafny.CodePoint('q')
		var _797_v41 D7
		_ = _797_v41
		_797_v41 = Companion_D7_.Create_DC20_(p0, _796_v40)
		var _798_v42 D0
		_ = _798_v42
		_798_v42 = Companion_D0_.Create_DC1_(!(_this.F13()), p0)
		var _799_v43 _dafny.Sequence
		_ = _799_v43
		_799_v43 = _dafny.SeqOf(p0)
		var _800_v44 _dafny.Map
		_ = _800_v44
		_800_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _dafny.CodePoint('x'))
		var _801_v45 _dafny.Array
		_ = _801_v45
		var _nwElement0_22 _dafny.CodePoint = _796_v40
		_ = _nwElement0_22
		var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
		_ = _nw112
		(_nw112).ArraySet1CodePoint(_nwElement0_22, 0)
		(_nw112).ArraySet1CodePoint(_796_v40, 1)
		(_nw112).ArraySet1CodePoint(_796_v40, 2)
		(_nw112).ArraySet1CodePoint((_797_v41).Dtor_cf32(), 3)
		(_nw112).ArraySet1CodePoint(_dafny.CodePoint('a'), 4)
		(_nw112).ArraySet1CodePoint(_796_v40, 5)
		(_nw112).ArraySet1CodePoint(_796_v40, 6)
		(_nw112).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _this.F13() {
				return _796_v40
			}
			return _796_v40
		})(), 7)
		(_nw112).ArraySet1CodePoint(_dafny.CodePoint('h'), 8)
		(_nw112).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_this).Fm4((_this).F19(), _798_v42, p0, Companion_Default___.Fm39(_dafny.IntOfUint32((_799_v43).Cardinality()), globalState), globalState) {
				return _796_v40
			}
			return _796_v40
		})(), 9)
		(_nw112).ArraySet1CodePoint(_796_v40, 10)
		(_nw112).ArraySet1CodePoint(_796_v40, 11)
		(_nw112).ArraySet1CodePoint(_796_v40, 12)
		(_nw112).ArraySet1CodePoint((func() _dafny.CodePoint {
			if false {
				return _796_v40
			}
			return _796_v40
		})(), 13)
		(_nw112).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _this.F13() {
				return _dafny.CodePoint('e')
			}
			return _796_v40
		})(), 14)
		(_nw112).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_800_v44).Contains(true) {
				return (_800_v44).Get(true).(_dafny.CodePoint)
			}
			return _796_v40
		})(), 15)
		_801_v45 = _nw112
		_801_v45 = _801_v45
		var _802_v46 _dafny.Sequence
		_ = _802_v46
		_802_v46 = _dafny.UnicodeSeqOfUtf8Bytes("hvovm")
		var _803_v47 _dafny.Map
		_ = _803_v47
		_803_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fls")).Cardinality()), p0)
		var _804_v48 _dafny.Sequence
		_ = _804_v48
		_804_v48 = _dafny.SeqOf(!((_this).F19()))
		var _805_v49 _dafny.Set
		_ = _805_v49
		_805_v49 = _dafny.SetOf((_this).F19())
		var _806_v50 _dafny.Array
		_ = _806_v50
		var _nwElement0_23 bool = Companion_Default___.Fm0(_803_v47, p0, globalState)
		_ = _nwElement0_23
		var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(6))
		_ = _nw113
		(_nw113).ArraySet1(_nwElement0_23, 0)
		(_nw113).ArraySet1((_804_v48).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm31(globalState), _dafny.IntOfUint32((_804_v48).Cardinality()))).Uint32()).(bool), 1)
		(_nw113).ArraySet1((_804_v48).Select((Companion_Default___.SafeIndex((_805_v49).Cardinality(), _dafny.IntOfUint32((_804_v48).Cardinality()))).Uint32()).(bool), 2)
		(_nw113).ArraySet1(_this.F13(), 3)
		(_nw113).ArraySet1((_this).F19(), 4)
		(_nw113).ArraySet1(_this.F13(), 5)
		_806_v50 = _nw113
		var _807_v51 D1
		_ = _807_v51
		_807_v51 = Companion_D1_.Create_DC5_(_dafny.Companion_Sequence_.Concatenate(_802_v46, _802_v46), _806_v50, p0)
		var _source15 D1 = _807_v51
		_ = _source15
		if _source15.Is_DC3() {
			var _808___mcc_h0 _dafny.Int = _source15.Get_().(D1_DC3).Cf4
			_ = _808___mcc_h0
			var _809_cf4 _dafny.Int = _808___mcc_h0
			_ = _809_cf4
			var _810_v52 D1
			_ = _810_v52
			_810_v52 = Companion_D1_.Create_DC4_(_806_v50, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(955))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_811_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_812_i7 _dafny.Int) _dafny.Int {
					return _811_cf4
				}
			})(_809_cf4)))).Cardinality())), _796_v40)
			var _813_v53 _dafny.Set
			_ = _813_v53
			_813_v53 = _dafny.SetOf(_810_v52, _810_v52, _810_v52, _810_v52)
			var _814_v54 _dafny.Sequence
			_ = _814_v54
			_814_v54 = _dafny.SeqOf(_813_v53)
			if ((_814_v54).Select((Companion_Default___.SafeIndex(_809_cf4, _dafny.IntOfUint32((_814_v54).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf((_813_v53).Difference(_813_v53)) {
				var _815_v55 _dafny.Map
				_ = _815_v55
				_815_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), p0)
				_815_v55 = (_815_v55).Update(!(!(_this.F13()) || (_this.F13())), _dafny.IntOfUint32((_dafny.SeqOf(_805_v49, _805_v49)).Cardinality()))
				var _816_v56 _dafny.Array
				_ = _816_v56
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_15
				var _nw114 _dafny.Array
				_ = _nw114
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw114 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Int = (func(_817_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_818_i8 _dafny.Int) _dafny.Int {
							return (_818_i8).Minus(_817_cf4)
						}
					})(_809_cf4)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw114 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw114).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw114).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_816_v56 = _nw114
				var _819_v57 *C1
				_ = _819_v57
				var _nw115 *C1 = New_C1_()
				_ = _nw115
				_nw115.Ctor__(_this.F13(), _802_v46, _816_v56)
				_819_v57 = _nw115
				var _820_v58 _dafny.MultiSet
				_ = _820_v58
				_820_v58 = _dafny.MultiSetOf((_797_v41).Dtor_cf31())
				_809_cf4 = (func() _dafny.Int {
					if (_820_v58).Contains(_809_cf4) {
						return (_820_v58).Multiplicity(_809_cf4)
					}
					return (_dafny.IntOfInt64(-547)).Times(_809_cf4)
				})()
				var _821_v59 _dafny.Array
				_ = _821_v59
				var _nwElement0_24 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(239))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_822_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_823_i9 _dafny.Int) _dafny.CodePoint {
						return _822_v40
					}
				})(_796_v40)))
				_ = _nwElement0_24
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(3))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_24, 0)
				(_nw116).ArraySet1(_802_v46, 1)
				(_nw116).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("f"), 2)
				_821_v59 = _nw116
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_821_v59), 0))
				_ = _index145
				(_821_v59).ArraySet1(_802_v46, (_index145).Int())
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_821_v59), 0))
				_ = _index146
				(_821_v59).ArraySet1(_dafny.Companion_Sequence_.Update(_819_v57.F18, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_819_v57.F18).Cardinality()))).Uint32(), _796_v40), (_index146).Int())
				var _rhs110 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("bplas")
				_ = _rhs110
				var _rhs111 *C1 = _819_v57
				_ = _rhs111
				var _rhs112 _dafny.Sequence = _799_v43
				_ = _rhs112
				var _lhs69 *C1 = _819_v57
				_ = _lhs69
				_lhs69.F18 = _rhs110
				_819_v57 = _rhs111
				_799_v43 = _rhs112
			} else {
				var _824_v60 _dafny.Array
				_ = _824_v60
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_16
				var _nw117 _dafny.Array
				_ = _nw117
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw117 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = (func(_825_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_826_i10 _dafny.Int) _dafny.Int {
							return (_826_i10).Times(_825_cf4)
						}
					})(_809_cf4)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw117 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw117).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw117).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_824_v60 = _nw117
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_824_v60), 0))
				_ = _index147
				(_824_v60).ArraySet1(Companion_Default___.Fm31(globalState), (_index147).Int())
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_824_v60), 0))
				_ = _index148
				(_824_v60).ArraySet1(p0, (_index148).Int())
				var _827_v61 _dafny.Map
				_ = _827_v61
				_827_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_824_v60, _824_v60)
				_827_v61 = (_827_v61).Update(_824_v60, (func() _dafny.Array {
					if _this.F13() {
						return _824_v60
					}
					return _824_v60
				})())
				var _828_v62 _dafny.Set
				_ = _828_v62
				_828_v62 = _dafny.SetOf((_824_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_824_v60), 0))).Int()).(_dafny.Int))
				var _829_v63 D2
				_ = _829_v63
				_829_v63 = Companion_D2_.Create_DC7_(_802_v46, _809_cf4, _dafny.IntOfInt64(-989), _this.F13())
				var _830_v64 _dafny.MultiSet
				_ = _830_v64
				_830_v64 = _dafny.MultiSetOf((_828_v62).Difference(_dafny.SetOf((_829_v63).Dtor_cf13(), (_824_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_824_v60), 0))).Int()).(_dafny.Int), (_824_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_824_v60), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_831_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_832_i11 _dafny.Int) _dafny.CodePoint {
						return _831_v40
					}
				})(_796_v40)))).Cardinality()))), _828_v62)
				_830_v64 = _830_v64
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_824_v60), 0))
				_ = _index149
				(_824_v60).ArraySet1(_809_cf4, (_index149).Int())
				_796_v40 = _dafny.CodePoint('v')
			}
			(_this).F13_set_((Companion_Default___.Fm31(globalState)).Cmp(p0) >= 0)
			var _833_v65 _dafny.Array
			_ = _833_v65
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw118
			_833_v65 = _nw118
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_833_v65), 0))
			_ = _index150
			(_833_v65).ArraySet1((func() _dafny.Sequence {
				if _this.F13() {
					return _802_v46
				}
				return _802_v46
			})(), (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_833_v65), 0))
			_ = _index151
			(_833_v65).ArraySet1((Companion_D2_.Create_DC7_(_802_v46, _dafny.IntOfInt64(-656), _dafny.IntOfUint32((_799_v43).Cardinality()), _this.F13())).Dtor_cf12(), (_index151).Int())
			var _834_v66 _dafny.Array
			_ = _834_v66
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_17
			var _nw119 _dafny.Array
			_ = _nw119
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw119 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Map = (func(_835_p0 _dafny.Int, _836_cf4 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_837_i12 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_835_p0, _836_cf4)
					}
				})(p0, _809_cf4)
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
			_834_v66 = _nw119
			_834_v66 = _834_v66
		} else if _source15.Is_DC4() {
			var _838___mcc_h1 _dafny.Array = _source15.Get_().(D1_DC4).Cf5
			_ = _838___mcc_h1
			var _839___mcc_h2 _dafny.Int = _source15.Get_().(D1_DC4).Cf6
			_ = _839___mcc_h2
			var _840___mcc_h3 _dafny.CodePoint = _source15.Get_().(D1_DC4).Cf7
			_ = _840___mcc_h3
			var _841_cf7 _dafny.CodePoint = _840___mcc_h3
			_ = _841_cf7
			var _842_cf6 _dafny.Int = _839___mcc_h2
			_ = _842_cf6
			var _843_cf5 _dafny.Array = _838___mcc_h1
			_ = _843_cf5
			_842_cf6 = Companion_Default___.SafeModuloInt(p0, (func() _dafny.Int {
				if (_this).F19() {
					return Companion_Default___.Fm31(globalState)
				}
				return _842_cf6
			})())
			var _844_v67 _dafny.Array
			_ = _844_v67
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_18
			var _nw120 _dafny.Array
			_ = _nw120
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw120 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) _dafny.Int = func(_845_i13 _dafny.Int) _dafny.Int {
					return (_845_i13).Times(_dafny.IntOfInt64(862))
				}
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
			_844_v67 = _nw120
			var _846_v68 _dafny.Sequence
			_ = _846_v68
			_846_v68 = _dafny.SeqOf(_dafny.SeqOf(Companion_Default___.Fm31(globalState), (_805_v49).Cardinality()), _dafny.SeqOf(_842_cf6, p0), _799_v43, _799_v43)
			var _847_v69 _dafny.Map
			_ = _847_v69
			_847_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F13())
			var _848_v70 _dafny.MultiSet
			_ = _848_v70
			_848_v70 = _dafny.MultiSetOf(_this.F13(), (_this).F19())
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
			_ = _index152
			(_844_v67).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_803_v47).Contains(p0) {
					return (_803_v47).Get(p0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_846_v68).Cardinality())
			})())), ((_847_v69).Update((_848_v70).Cardinality(), (_this).F19())).Cardinality()), (_index152).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
			_ = _index153
			(_844_v67).ArraySet1(_842_cf6, (_index153).Int())
			var _849_v71 _dafny.Map
			_ = _849_v71
			_849_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_841_cf7, _806_v50)
			_849_v71 = (_849_v71).Update(_841_cf7, _806_v50)
			var _850_v72 _dafny.MultiSet
			_ = _850_v72
			_850_v72 = _dafny.MultiSetOf(_842_cf6)
			var _851_v73 _dafny.Map
			_ = _851_v73
			_851_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), (_this).F19())
			if (_dafny.MultiSetOf((_dafny.Zero).Minus(((_851_v73).Update(_this.F13(), _this.F13())).Cardinality()), (_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int), _842_cf6)).IsSubsetOf(_850_v72) {
				var _852_v74 D1
				_ = _852_v74
				_852_v74 = Companion_D1_.Create_DC4_(_806_v50, _842_cf6, _796_v40)
				var _853_v75 *C2
				_ = _853_v75
				var _nw121 *C2 = New_C2_()
				_ = _nw121
				_nw121.Ctor__(_this.F13())
				_853_v75 = _nw121
				var _854_v76 _dafny.Sequence
				_ = _854_v76
				_854_v76 = _dafny.SeqOf(_853_v75)
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
				_ = _index154
				var _rhs113 _dafny.CodePoint = _841_cf7
				_ = _rhs113
				var _rhs114 _dafny.Int = (_852_v74).Dtor_cf6()
				_ = _rhs114
				var _rhs115 bool = ((_dafny.IntOfInt64(890)).Times((_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfInt64(948)) < 0
				_ = _rhs115
				var _rhs116 _dafny.MultiSet = ((_850_v72).Difference(Companion_Default___.Fm40(_dafny.IntOfUint32((_854_v76).Cardinality()), _this.F13(), _this.F13(), (_this).F19(), globalState))).Difference(_850_v72)
				_ = _rhs116
				var _rhs117 _dafny.Int = (_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int)
				_ = _rhs117
				var _lhs70 *C3 = _this
				_ = _lhs70
				var _lhs71 _dafny.Array = _844_v67
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
				_ = _lhs72
				_841_cf7 = _rhs113
				_842_cf6 = _rhs114
				_lhs70.F13_set_(_rhs115)
				_850_v72 = _rhs116
				(_lhs71).ArraySet1(_rhs117, (_lhs72).Int())
				var _855_v77 _dafny.Array
				_ = _855_v77
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw122
				_855_v77 = _nw122
				var _856_v78 _dafny.Map
				_ = _856_v78
				_856_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_855_v77, Companion_Default___.Fm41(p0, globalState))
				_856_v78 = (_856_v78).Update(_855_v77, Companion_D4_.Create_DC11_(!((_this).F19())))
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
				_ = _index155
				(_844_v67).ArraySet1((_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int), (_index155).Int())
				var _857_v79 _dafny.Map
				_ = _857_v79
				_857_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_850_v72, (_dafny.IntOfInt64(888)).Minus(_842_cf6))
				var _858_v80 _dafny.Set
				_ = _858_v80
				_858_v80 = _dafny.SetOf((_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int))
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
				_ = _index156
				(_844_v67).ArraySet1((func() _dafny.Int {
					if (_857_v79).Contains(Companion_Default___.Fm40((_dafny.Zero).Minus((_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int)), false, _this.F13(), _this.F13(), globalState)) {
						return (_857_v79).Get(Companion_Default___.Fm40((_dafny.Zero).Minus((_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int)), false, _this.F13(), _this.F13(), globalState)).(_dafny.Int)
					}
					return (_858_v80).Cardinality()
				})(), (_index156).Int())
				var _859_v81 *C2
				_ = _859_v81
				var _nw123 *C2 = New_C2_()
				_ = _nw123
				_nw123.Ctor__(_this.F13())
				_859_v81 = _nw123
			} else {
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))
				_ = _index157
				(_844_v67).ArraySet1((_844_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_844_v67), 0))).Int()).(_dafny.Int), (_index157).Int())
				(_this).F13_set_((_804_v48).Select((Companion_Default___.SafeIndex((_842_cf6).Plus(_dafny.IntOfUint32((_802_v46).Cardinality())), _dafny.IntOfUint32((_804_v48).Cardinality()))).Uint32()).(bool))
				(_this).F13_set_((p0).Cmp(_842_cf6) < 0)
				var _860_v83 _dafny.Map
				_ = _860_v83
				_860_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v46, _796_v40)
				var _861_v84 _dafny.Map
				_ = _861_v84
				_861_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(150))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_862_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_863_i14 _dafny.Int) _dafny.CodePoint {
						return _862_v40
					}
				})(_796_v40))), (_804_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.IntOfUint32((_804_v48).Cardinality()))).Uint32()).(bool))
				var _864_v85 _dafny.Int
				_ = _864_v85
				var _out30 _dafny.Int
				_ = _out30
				_out30 = Companion_Default___.M0((func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter33 := _dafny.Iterate((_860_v83).Keys().Elements()); ; {
						_compr_31, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _865_v82 _dafny.Sequence
						_865_v82 = interface{}(_compr_31).(_dafny.Sequence)
						if (_860_v83).Contains(_865_v82) {
							_coll31.Add(_865_v82, true)
						}
					}
					return _coll31.ToMap()
				}()).Merge(_861_v84), globalState)
				_864_v85 = _out30
				_842_cf6 = (_dafny.Zero).Minus(p0)
			}
		} else if _source15.Is_DC5() {
			var _866___mcc_h4 _dafny.Sequence = _source15.Get_().(D1_DC5).Cf8
			_ = _866___mcc_h4
			var _867___mcc_h5 _dafny.Array = _source15.Get_().(D1_DC5).Cf9
			_ = _867___mcc_h5
			var _868___mcc_h6 _dafny.Int = _source15.Get_().(D1_DC5).Cf10
			_ = _868___mcc_h6
			var _869_cf10 _dafny.Int = _868___mcc_h6
			_ = _869_cf10
			var _870_cf9 _dafny.Array = _867___mcc_h5
			_ = _870_cf9
			var _871_cf8 _dafny.Sequence = _866___mcc_h4
			_ = _871_cf8
			(_this).F13_set_(!((_this).F19()))
			var _872_v86 _dafny.Map
			_ = _872_v86
			_872_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(808), _806_v50)
			var _873_v87 _dafny.Set
			_ = _873_v87
			_873_v87 = _dafny.SetOf(_870_cf9, _870_cf9, (func() _dafny.Array {
				if (_872_v86).Contains(_869_cf10) {
					return (_872_v86).Get(_869_cf10).(_dafny.Array)
				}
				return _870_cf9
			})())
			_873_v87 = (_dafny.SetOf(_806_v50, _806_v50)).Difference(_873_v87)
			_869_cf10 = _869_cf10
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_806_v50), 0))
			_ = _index158
			(_806_v50).ArraySet1(((_this).F19()) || (_this.F13()), (_index158).Int())
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_806_v50), 0))
			_ = _index159
			var _rhs118 _dafny.Array = (func() _dafny.Array {
				if !(_this.F13()) || (false) {
					return _870_cf9
				}
				return _806_v50
			})()
			_ = _rhs118
			var _rhs119 _dafny.Int = Companion_Default___.SafeModuloInt(_869_cf10, _869_cf10)
			_ = _rhs119
			var _rhs120 bool = _this.F13()
			_ = _rhs120
			var _lhs73 _dafny.Array = _806_v50
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_806_v50), 0))
			_ = _lhs74
			_870_cf9 = _rhs118
			_869_cf10 = _rhs119
			(_lhs73).ArraySet1(_rhs120, (_lhs74).Int())
		} else {
			var _874___mcc_h7 _dafny.Map = _source15.Get_().(D1_DC2).Cf3
			_ = _874___mcc_h7
			var _875_cf3 _dafny.Map = _874___mcc_h7
			_ = _875_cf3
			var _876_v88 _dafny.Int
			_ = _876_v88
			_876_v88 = _dafny.IntOfInt64(120)
			var _877_v89 _dafny.Sequence
			_ = _877_v89
			_877_v89 = _dafny.SeqOf(_804_v48)
			var _878_v90 _dafny.Sequence
			_ = _878_v90
			_878_v90 = (_877_v89).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_877_v89).Cardinality()))).Uint32()).(_dafny.Sequence)
			_876_v88 = _dafny.IntOfUint32((_878_v90).Cardinality())
			var _879_v91 *C0
			_ = _879_v91
			var _nw124 *C0 = New_C0_()
			_ = _nw124
			_nw124.Ctor__(_876_v88)
			_879_v91 = _nw124
			var _880_v92 _dafny.Array
			_ = _880_v92
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_19
			var _nw125 _dafny.Array
			_ = _nw125
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw125 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = (func(_881_v91 *C0) func(_dafny.Int) _dafny.Int {
					return func(_882_i15 _dafny.Int) _dafny.Int {
						return (_882_i15).Times((_881_v91).F15())
					}
				})(_879_v91)
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
			_880_v92 = _nw125
			var _883_v93 *C1
			_ = _883_v93
			var _nw126 *C1 = New_C1_()
			_ = _nw126
			_nw126.Ctor__(_this.F13(), _dafny.Companion_Sequence_.Concatenate(_802_v46, _dafny.UnicodeSeqOfUtf8Bytes("uyflgki")), _880_v92)
			_883_v93 = _nw126
			_876_v88 = (_879_v91).F15()
		}
		var _884_v94 D2
		_ = _884_v94
		_884_v94 = Companion_D2_.Create_DC8_(true, _this.F13(), _this.F13())
		(_this).F13_set_(!(func(_source16 D2) bool {
			if _source16.Is_DC7() {
				var _885___mcc_h8 _dafny.Sequence = _source16.Get_().(D2_DC7).Cf12
				_ = _885___mcc_h8
				var _886___mcc_h9 _dafny.Int = _source16.Get_().(D2_DC7).Cf13
				_ = _886___mcc_h9
				var _887___mcc_h10 _dafny.Int = _source16.Get_().(D2_DC7).Cf14
				_ = _887___mcc_h10
				var _888___mcc_h11 bool = _source16.Get_().(D2_DC7).Cf15
				_ = _888___mcc_h11
				var _889_cf15 bool = _888___mcc_h11
				_ = _889_cf15
				var _890_cf14 _dafny.Int = _887___mcc_h10
				_ = _890_cf14
				var _891_cf13 _dafny.Int = _886___mcc_h9
				_ = _891_cf13
				var _892_cf12 _dafny.Sequence = _885___mcc_h8
				_ = _892_cf12
				return _this.F13()
			} else if _source16.Is_DC8() {
				var _893___mcc_h12 bool = _source16.Get_().(D2_DC8).Cf16
				_ = _893___mcc_h12
				var _894___mcc_h13 bool = _source16.Get_().(D2_DC8).Cf17
				_ = _894___mcc_h13
				var _895___mcc_h14 bool = _source16.Get_().(D2_DC8).Cf18
				_ = _895___mcc_h14
				var _896_cf18 bool = _895___mcc_h14
				_ = _896_cf18
				var _897_cf17 bool = _894___mcc_h13
				_ = _897_cf17
				var _898_cf16 bool = _893___mcc_h12
				_ = _898_cf16
				return _898_cf16
			} else {
				var _899___mcc_h15 _dafny.Sequence = _source16.Get_().(D2_DC6).Cf11
				_ = _899___mcc_h15
				var _900_cf11 _dafny.Sequence = _899___mcc_h15
				_ = _900_cf11
				return (_this).F19()
			}
		}(_884_v94)))
		var _901_v95 _dafny.Array
		_ = _901_v95
		var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw127
		_901_v95 = _nw127
		var _902_v96 D4
		_ = _902_v96
		_902_v96 = Companion_D4_.Create_DC10_(_901_v95)
		r0 = (_902_v96).Dtor_cf20()
		return r0
	}
}
func (_this *C3) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Sequence) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _903_i0 _dafny.Int
		_ = _903_i0
		_903_i0 = _dafny.Zero
		{
			for false {
				{
					if (_903_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_903_i0 = (_903_i0).Plus(_dafny.One)
					var _904_v0 _dafny.Array
					_ = _904_v0
					var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
					_ = _nw128
					_904_v0 = _nw128
					var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_904_v0), 0))
					_ = _index160
					(_904_v0).ArraySet1(p2, (_index160).Int())
					var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_904_v0), 0))
					_ = _index161
					(_904_v0).ArraySet1(p2, (_index161).Int())
					var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_904_v0), 0))
					_ = _index162
					(_904_v0).ArraySet1((func() _dafny.Int {
						if _this.F13() {
							return p0
						}
						return p2
					})(), (_index162).Int())
					var _905_v1 _dafny.Array
					_ = _905_v1
					var _len0_20 _dafny.Int = _dafny.IntOfInt64(26)
					_ = _len0_20
					var _nw129 _dafny.Array
					_ = _nw129
					if _len0_20.Cmp(_dafny.Zero) == 0 {
						_nw129 = _dafny.NewArray(_len0_20)
					} else {
						var _init20 func(_dafny.Int) _dafny.MultiSet = func(_906_i1 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F13(), (_this).F19()))).Union(_dafny.MultiSetOf(_this.F13()))
						}
						_ = _init20
						var _element0_20 = _init20(_dafny.Zero)
						_ = _element0_20
						_nw129 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
						(_nw129).ArraySet1(_element0_20, 0)
						var _nativeLen0_20 = (_len0_20).Int()
						_ = _nativeLen0_20
						for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
							(_nw129).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
						}
					}
					_905_v1 = _nw129
					var _907_v2 _dafny.MultiSet
					_ = _907_v2
					_907_v2 = _dafny.MultiSetOf((_this).F19(), (_this).F19())
					var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_905_v1), 0))
					_ = _index163
					(_905_v1).ArraySet1(_907_v2, (_index163).Int())
					var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_905_v1), 0))
					_ = _index164
					(_905_v1).ArraySet1((_dafny.MultiSetOf(true)).Union(_dafny.MultiSetOf(_this.F13())), (_index164).Int())
					var _908_v3 _dafny.Sequence
					_ = _908_v3
					_908_v3 = _dafny.SeqOf(p2)
					var _909_v4 _dafny.Map
					_ = _909_v4
					_909_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_908_v3, _908_v3)
					(_this).F13_set_(!(_909_v4).Contains(_908_v3))
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _910_v6 _dafny.Sequence
		_ = _910_v6
		_910_v6 = _dafny.SeqOf(_dafny.IntOfInt64(785))
		var _911_v7 _dafny.Map
		_ = _911_v7
		_911_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _dafny.IntOfUint32((_910_v6).Cardinality()))
		var _912_v8 _dafny.MultiSet
		_ = _912_v8
		_912_v8 = _dafny.MultiSetOf((_911_v7).Cardinality(), (_dafny.MultiSetOf((_this).F19())).Cardinality())
		var _913_v9 _dafny.Map
		_ = _913_v9
		_913_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		r2 = (Companion_Default___.Fm0(func() _dafny.Map {
			var _coll32 = _dafny.NewMapBuilder()
			_ = _coll32
			for _iter34 := _dafny.Iterate((_912_v8).Elements()); ; {
				_compr_32, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _914_v5 _dafny.Int
				_914_v5 = interface{}(_compr_32).(_dafny.Int)
				if (_912_v8).Contains(_914_v5) {
					_coll32.Add(Companion_Default___.SafeDivisionInt(_914_v5, p2), p0)
				}
			}
			return _coll32.ToMap()
		}(), (_913_v9).Cardinality(), globalState)) && ((_this).F19())
		var _915_i2 _dafny.Int
		_ = _915_i2
		_915_i2 = _dafny.Zero
		{
			for (p0).Cmp(p0) < 0 {
				{
					if (_915_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_915_i2 = (_915_i2).Plus(_dafny.One)
					var _916_v10 D0
					_ = _916_v10
					_916_v10 = Companion_D0_.Create_DC0_(_this.F13())
					var _917_v11 _dafny.Map
					_ = _917_v11
					_917_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_916_v10).Dtor_cf0())
					var _918_v13 _dafny.Sequence
					_ = _918_v13
					_918_v13 = _dafny.SeqOf(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-25))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg55 _dafny.Int) interface{} {
							return coer55(arg55)
						}
					}(func(_919_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('c')
					})), p1)
					var _920_v15 _dafny.Sequence
					_ = _920_v15
					_920_v15 = _dafny.SeqOf(_this.F13(), _this.F13())
					var _921_v16 _dafny.Sequence
					_ = _921_v16
					_921_v16 = _dafny.SeqOf(_920_v15)
					var _922_v17 _dafny.Array
					_ = _922_v17
					var _nwElement0_25 _dafny.Int = p2
					_ = _nwElement0_25
					var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(18))
					_ = _nw130
					(_nw130).ArraySet1(_nwElement0_25, 0)
					(_nw130).ArraySet1(p0, 1)
					(_nw130).ArraySet1(p2, 2)
					(_nw130).ArraySet1((func() _dafny.Set {
						var _coll33 = _dafny.NewBuilder()
						_ = _coll33
						for _iter35 := _dafny.Iterate((func() _dafny.Map {
							var _coll34 = _dafny.NewMapBuilder()
							_ = _coll34
							for _iter36 := _dafny.Iterate((_918_v13).Elements()); ; {
								_compr_34, _ok36 := _iter36()
								if !_ok36 {
									break
								}
								var _923_v12 _dafny.Sequence
								_923_v12 = interface{}(_compr_34).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_918_v13, _923_v12) {
									_coll34.Add(_923_v12, (_this).F19())
								}
							}
							return _coll34.ToMap()
						}()).Keys().Elements()); ; {
							_compr_33, _ok35 := _iter35()
							if !_ok35 {
								break
							}
							var _924_v14 _dafny.Sequence
							_924_v14 = interface{}(_compr_33).(_dafny.Sequence)
							if (func() _dafny.Map {
								var _coll35 = _dafny.NewMapBuilder()
								_ = _coll35
								for _iter37 := _dafny.Iterate((_918_v13).Elements()); ; {
									_compr_35, _ok37 := _iter37()
									if !_ok37 {
										break
									}
									var _923_v12 _dafny.Sequence
									_923_v12 = interface{}(_compr_35).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_918_v13, _923_v12) {
										_coll35.Add(_923_v12, (_this).F19())
									}
								}
								return _coll35.ToMap()
							}()).Contains(_924_v14) {
								_coll33.Add(_924_v14)
							}
						}
						return _coll33.ToSet()
					}()).Cardinality(), 3)
					(_nw130).ArraySet1(p0, 4)
					(_nw130).ArraySet1(p0, 5)
					(_nw130).ArraySet1(p0, 6)
					(_nw130).ArraySet1((_dafny.Zero).Minus(p2), 7)
					(_nw130).ArraySet1(p2, 8)
					(_nw130).ArraySet1(p2, 9)
					(_nw130).ArraySet1(p0, 10)
					(_nw130).ArraySet1(_dafny.IntOfUint32((_921_v16).Cardinality()), 11)
					(_nw130).ArraySet1(p2, 12)
					(_nw130).ArraySet1(p2, 13)
					(_nw130).ArraySet1(p0, 14)
					(_nw130).ArraySet1(p2, 15)
					(_nw130).ArraySet1(p2, 16)
					(_nw130).ArraySet1(p0, 17)
					_922_v17 = _nw130
					var _925_v18 *C1
					_ = _925_v18
					var _nw131 *C1 = New_C1_()
					_ = _nw131
					_nw131.Ctor__((func() bool {
						if (_917_v11).Contains(p2) {
							return (_917_v11).Get(p2).(bool)
						}
						return Companion_Default___.Fm0(_913_v9, p0, globalState)
					})(), _dafny.UnicodeSeqOfUtf8Bytes("wceigad"), _922_v17)
					_925_v18 = _nw131
					var _926_v19 _dafny.CodePoint
					_ = _926_v19
					_926_v19 = _dafny.CodePoint('p')
					_926_v19 = _926_v19
					(_this).F13_set_((_925_v18).F17())
					var _927_v20 _dafny.Map
					_ = _927_v20
					_927_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_920_v15, (_this).F19())
					_927_v20 = (_927_v20).Update(_920_v15, (_this).F19())
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _928_v21 D2
		_ = _928_v21
		_928_v21 = Companion_D2_.Create_DC6_(_910_v6)
		var _pat_let_tv19 = _910_v6
		_ = _pat_let_tv19
		var _pat_let_tv20 = p2
		_ = _pat_let_tv20
		var _pat_let_tv21 = _910_v6
		_ = _pat_let_tv21
		var _pat_let_tv22 = _912_v8
		_ = _pat_let_tv22
		if func(_source17 D2) bool {
			if _source17.Is_DC7() {
				var _929___mcc_h0 _dafny.Sequence = _source17.Get_().(D2_DC7).Cf12
				_ = _929___mcc_h0
				var _930___mcc_h1 _dafny.Int = _source17.Get_().(D2_DC7).Cf13
				_ = _930___mcc_h1
				var _931___mcc_h2 _dafny.Int = _source17.Get_().(D2_DC7).Cf14
				_ = _931___mcc_h2
				var _932___mcc_h3 bool = _source17.Get_().(D2_DC7).Cf15
				_ = _932___mcc_h3
				var _933_cf15 bool = _932___mcc_h3
				_ = _933_cf15
				var _934_cf14 _dafny.Int = _931___mcc_h2
				_ = _934_cf14
				var _935_cf13 _dafny.Int = _930___mcc_h1
				_ = _935_cf13
				var _936_cf12 _dafny.Sequence = _929___mcc_h0
				_ = _936_cf12
				return _this.F13()
			} else if _source17.Is_DC8() {
				var _937___mcc_h4 bool = _source17.Get_().(D2_DC8).Cf16
				_ = _937___mcc_h4
				var _938___mcc_h5 bool = _source17.Get_().(D2_DC8).Cf17
				_ = _938___mcc_h5
				var _939___mcc_h6 bool = _source17.Get_().(D2_DC8).Cf18
				_ = _939___mcc_h6
				var _940_cf18 bool = _939___mcc_h6
				_ = _940_cf18
				var _941_cf17 bool = _938___mcc_h5
				_ = _941_cf17
				var _942_cf16 bool = _937___mcc_h4
				_ = _942_cf16
				return (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('g'))).Cardinality())).Cmp((_pat_let_tv19).Select((Companion_Default___.SafeIndex(_pat_let_tv20, _dafny.IntOfUint32((_pat_let_tv21).Cardinality()))).Uint32()).(_dafny.Int)) < 0
			} else {
				var _943___mcc_h7 _dafny.Sequence = _source17.Get_().(D2_DC6).Cf11
				_ = _943___mcc_h7
				var _944_cf11 _dafny.Sequence = _943___mcc_h7
				_ = _944_cf11
				return ((_pat_let_tv22).Cardinality()).Cmp(_dafny.IntOfInt64(75)) < 0
			}
		}(_928_v21) {
			_913_v9 = (_913_v9).Update(p0, _dafny.IntOfInt64(733))
			var _945_v22 _dafny.Set
			_ = _945_v22
			_945_v22 = _dafny.SetOf(p0, _dafny.IntOfInt64(150))
			var _946_v23 _dafny.Array
			_ = _946_v23
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_21
			var _nw132 _dafny.Array
			_ = _nw132
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw132 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_947_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_948_i4 _dafny.Int) _dafny.Int {
						return (_948_i4).Plus(_947_p0)
					}
				})(p0)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw132 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw132).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw132).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_946_v23 = _nw132
			var _949_v24 _dafny.Map
			_ = _949_v24
			_949_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v23, !((_this).F19()))
			var _950_v25 _dafny.Map
			_ = _950_v25
			_950_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_945_v22, (_949_v24).Merge(_949_v24))
			_950_v25 = (_950_v25).Update((_dafny.SetOf(p0)).Difference(_945_v22), _949_v24)
			var _951_v27 *C1
			_ = _951_v27
			var _nw133 *C1 = New_C1_()
			_ = _nw133
			_nw133.Ctor__(Companion_Default___.Fm0(func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-765), _dafny.IntOfInt64(322))); ; {
					_compr_36, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _952_v26 _dafny.Int
					_952_v26 = interface{}(_compr_36).(_dafny.Int)
					if ((_dafny.IntOfInt64(-765)).Cmp(_952_v26) <= 0) && ((_952_v26).Cmp(_dafny.IntOfInt64(322)) < 0) {
						_coll36.Add((_952_v26).Plus((_912_v8).Cardinality()), p0)
					}
				}
				return _coll36.ToMap()
			}(), (_910_v6).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_910_v6).Cardinality()))).Uint32()).(_dafny.Int), globalState), p1, _946_v23)
			_951_v27 = _nw133
			_951_v27 = _951_v27
			var _953_v28 D4
			_ = _953_v28
			_953_v28 = Companion_D4_.Create_DC11_((_951_v27).F17())
			var _954_v29 _dafny.MultiSet
			_ = _954_v29
			_954_v29 = _dafny.MultiSetOf(_953_v28)
			var _955_v30 _dafny.Sequence
			_ = _955_v30
			_955_v30 = _dafny.SeqOf(_954_v29, _954_v29, _dafny.MultiSetOf(_953_v28), _954_v29, _954_v29)
			var _956_v31 _dafny.Array
			_ = _956_v31
			var _nwElement0_26 bool = _this.F13()
			_ = _nwElement0_26
			var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(12))
			_ = _nw134
			(_nw134).ArraySet1(_nwElement0_26, 0)
			(_nw134).ArraySet1(false, 1)
			(_nw134).ArraySet1((_951_v27).F17(), 2)
			(_nw134).ArraySet1(((_951_v27).F17()) || ((_this).F19()), 3)
			(_nw134).ArraySet1((_this).F19(), 4)
			(_nw134).ArraySet1(_this.F13(), 5)
			(_nw134).ArraySet1(_this.F13(), 6)
			(_nw134).ArraySet1(_this.F13(), 7)
			(_nw134).ArraySet1((_951_v27).F17(), 8)
			(_nw134).ArraySet1(_this.F13(), 9)
			(_nw134).ArraySet1((_954_v29).IsDisjointFrom((_955_v30).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_955_v30).Cardinality()))).Uint32()).(_dafny.MultiSet)), 10)
			(_nw134).ArraySet1((_951_v27).F17(), 11)
			_956_v31 = _nw134
			var _957_v32 D5
			_ = _957_v32
			_957_v32 = Companion_D5_.Create_DC14_(!((_this).F19()))
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_956_v31), 0))
			_ = _index165
			(_956_v31).ArraySet1((func(_pat_let24_0 D5) D5 {
				return func(_958_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let25_0 bool) D5 {
						return func(_959_dt__update_hcf23_h0 bool) D5 {
							return Companion_D5_.Create_DC14_(_959_dt__update_hcf23_h0)
						}(_pat_let25_0)
					}(_this.F13())
				}(_pat_let24_0)
			}(_957_v32)).Dtor_cf23(), (_index165).Int())
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_956_v31), 0))
			_ = _index166
			(_956_v31).ArraySet1(_this.F13(), (_index166).Int())
			var _960_v33 _dafny.Int
			_ = _960_v33
			_960_v33 = _dafny.IntOfInt64(470)
			_960_v33 = _dafny.IntOfInt64(203)
		} else {
			var _961_v34 _dafny.Map
			_ = _961_v34
			_961_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())
			var _rhs121 bool = !(!(!(!(true))) || (!((_this).F19())))
			_ = _rhs121
			var _rhs122 bool = !(!(((((_961_v34).Update(_this.F13(), (_this).F19())).Cardinality()).Cmp(p2) >= 0) || (_this.F13())))
			_ = _rhs122
			r2 = _rhs121
			r1 = _rhs122
			var _962_v35 _dafny.Map
			_ = _962_v35
			_962_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.SafeModuloInt(p0, p0))
			var _963_v36 _dafny.CodePoint
			_ = _963_v36
			_963_v36 = _dafny.CodePoint('d')
			_962_v35 = (_962_v35).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _963_v36), p1), (_dafny.Zero).Minus(p0))
			_963_v36 = Companion_Default___.Fm34(globalState)
			var _964_v38 _dafny.Array
			_ = _964_v38
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_22
			var _nw135 _dafny.Array
			_ = _nw135
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw135 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = (func(_965_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
					return func(_966_i5 _dafny.Int) _dafny.Int {
						return (_966_i5).Plus((Companion_D7_.Create_DC20_((func() _dafny.Map {
							var _coll37 = _dafny.NewMapBuilder()
							_ = _coll37
							for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-622), _dafny.IntOfInt64(300))); ; {
								_compr_37, _ok39 := _iter39()
								if !_ok39 {
									break
								}
								var _967_v37 _dafny.Int
								_967_v37 = interface{}(_compr_37).(_dafny.Int)
								if ((_dafny.IntOfInt64(-622)).Cmp(_967_v37) <= 0) && ((_967_v37).Cmp(_dafny.IntOfInt64(300)) < 0) {
									_coll37.Add(Companion_Default___.SafeDivisionInt(_967_v37, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.IntOfInt64(969))
								}
							}
							return _coll37.ToMap()
						}()).Cardinality(), (Companion_D7_.Create_DC20_(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _965_v36)).Dtor_cf32())).Dtor_cf31())
					}
				})(_963_v36)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw135 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw135).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw135).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_964_v38 = _nw135
			var _968_v39 _dafny.Set
			_ = _968_v39
			_968_v39 = _dafny.SetOf(_964_v38, _964_v38)
			var _rhs123 bool = _this.F13()
			_ = _rhs123
			var _rhs124 bool = !((_968_v39).IsDisjointFrom((_968_v39).Difference(_968_v39)))
			_ = _rhs124
			r1 = _rhs123
			r2 = _rhs124
			var _969_v40 _dafny.Array
			_ = _969_v40
			var _nw136 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(18))
			_ = _nw136
			_969_v40 = _nw136
			var _970_v41 T1
			_ = _970_v41
			var _nw137 *C1 = New_C1_()
			_ = _nw137
			_nw137.Ctor__((_this).F19(), p1, _964_v38)
			_970_v41 = _nw137
			var _971_v42 D7
			_ = _971_v42
			_971_v42 = Companion_D7_.Create_DC19_(_970_v41)
			var _972_v43 D7
			_ = _972_v43
			_972_v43 = Companion_D7_.Create_DC21_(_971_v42)
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_969_v40), 0))
			_ = _index167
			(_969_v40).ArraySet1(_972_v43, (_index167).Int())
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_969_v40), 0))
			_ = _index168
			(_969_v40).ArraySet1(_972_v43, (_index168).Int())
		}
		var _973_v44 _dafny.Sequence
		_ = _973_v44
		_973_v44 = _dafny.SeqOf(_this.F13())
		var _974_v45 _dafny.Sequence
		_ = _974_v45
		_974_v45 = _dafny.SeqOf((_973_v44).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_973_v44).Cardinality()))).Uint32()).(bool))
		var _975_v46 _dafny.Array
		_ = _975_v46
		var _nwElement0_27 bool = (_this).F19()
		_ = _nwElement0_27
		var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(26))
		_ = _nw138
		(_nw138).ArraySet1(_nwElement0_27, 0)
		(_nw138).ArraySet1(_this.F13(), 1)
		(_nw138).ArraySet1(!((_this).F19()), 2)
		(_nw138).ArraySet1((_974_v45).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_974_v45).Cardinality()))).Uint32()).(bool), 3)
		(_nw138).ArraySet1(_this.F13(), 4)
		(_nw138).ArraySet1((_this).Fm4((_this).F19(), Companion_Default___.Fm42(globalState), p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm31(globalState)), globalState), 5)
		(_nw138).ArraySet1(!((_this).F19()), 6)
		(_nw138).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(379))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}(func(_976_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))), 7)
		(_nw138).ArraySet1((_this).F19(), 8)
		(_nw138).ArraySet1((_this).F19(), 9)
		(_nw138).ArraySet1(!(_this.F13()), 10)
		(_nw138).ArraySet1(_this.F13(), 11)
		(_nw138).ArraySet1(_this.F13(), 12)
		(_nw138).ArraySet1(_this.F13(), 13)
		(_nw138).ArraySet1(_this.F13(), 14)
		(_nw138).ArraySet1((_this).F19(), 15)
		(_nw138).ArraySet1(_this.F13(), 16)
		(_nw138).ArraySet1(_dafny.Companion_Sequence_.Equal(_910_v6, _910_v6), 17)
		(_nw138).ArraySet1(_this.F13(), 18)
		(_nw138).ArraySet1((_this).F19(), 19)
		(_nw138).ArraySet1((_this).F19(), 20)
		(_nw138).ArraySet1(_this.F13(), 21)
		(_nw138).ArraySet1(false, 22)
		(_nw138).ArraySet1(_this.F13(), 23)
		(_nw138).ArraySet1(!((_this).F19()), 24)
		(_nw138).ArraySet1(!(_this.F13()), 25)
		_975_v46 = _nw138
		_975_v46 = _975_v46
		var _977_v47 _dafny.Array
		_ = _977_v47
		var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw139
		_977_v47 = _nw139
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_977_v47), 0))
		_ = _index169
		(_977_v47).ArraySet1(p0, (_index169).Int())
		var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_977_v47), 0))
		_ = _index170
		(_977_v47).ArraySet1(Companion_Default___.Fm31(globalState), (_index170).Int())
		var _978_v48 _dafny.Map
		_ = _978_v48
		_978_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_977_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_977_v47), 0))).Int()).(_dafny.Int), _975_v46)
		r0 = (_978_v48).Merge(_978_v48)
		r1 = (_dafny.MultiSetOf(_this.F13(), (_this).F19())).Contains((true) || ((_this).F19()))
		r2 = (_this).F19()
		r3 = _dafny.SeqOf(true, ((_977_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_977_v47), 0))).Int()).(_dafny.Int)).Cmp(p2) != 0, (_this).F19(), (func() bool {
			if (_this).F19() {
				return _this.F13()
			}
			return true
		})())
		return r0, r1, r2, r3
	}
}
func (_this *C3) F19() bool {
	{
		return _this._f19
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f13 bool
	_f16 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f13 = false
	_this._f16 = false
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

func (_this *C4) F13() bool {
	return _this._f13
}
func (_this *C4) F13_set_(value bool) {
	_this._f13 = value
}
func (_this *C4) Ctor__(f16 bool, f13 bool) {
	{
		(_this)._f16 = f16
		(_this)._f13 = f13
	}
}
func (_this *C4) Fm4(p0 bool, p1 D0, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return (_this).F16()
	}
}
func (_this *C4) Fm5(globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC3_(_dafny.IntOfInt64(621))
	}
}
func (_this *C4) Fm12(p0 _dafny.Int, p1 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C4) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _hi8 _dafny.Int = p0
		_ = _hi8
		for _979_i0 := p0; _979_i0.Cmp(_hi8) < 0; _979_i0 = _979_i0.Plus(_dafny.One) {
			var _980_v0 *C0
			_ = _980_v0
			var _nw140 *C0 = New_C0_()
			_ = _nw140
			_nw140.Ctor__(_979_i0)
			_980_v0 = _nw140
			var _981_v1 _dafny.Array
			_ = _981_v1
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_23
			var _nw141 _dafny.Array
			_ = _nw141
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw141 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Sequence = func(_982_i1 _dafny.Int) _dafny.Sequence {
					return (func() _dafny.Sequence {
						if _this.F13() {
							return _dafny.UnicodeSeqOfUtf8Bytes("hwoxu")
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("pxu")
					})()
				}
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
			_981_v1 = _nw141
			var _983_v2 _dafny.Sequence
			_ = _983_v2
			_983_v2 = _dafny.UnicodeSeqOfUtf8Bytes("fylr")
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_981_v1), 0))
			_ = _index171
			(_981_v1).ArraySet1(_983_v2, (_index171).Int())
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_981_v1), 0))
			_ = _index172
			(_981_v1).ArraySet1(_983_v2, (_index172).Int())
			_981_v1 = _981_v1
			var _984_v3 _dafny.Sequence
			_ = _984_v3
			_984_v3 = _dafny.SeqOf((_980_v0).F15())
			var _985_v4 _dafny.Sequence
			_ = _985_v4
			_985_v4 = _dafny.SeqOf(_this.F13())
			var _986_v5 D0
			_ = _986_v5
			_986_v5 = Companion_D0_.Create_DC1_(_this.F13(), p0)
			var _987_v6 _dafny.CodePoint
			_ = _987_v6
			_987_v6 = _dafny.CodePoint('k')
			var _988_v7 _dafny.Map
			_ = _988_v7
			_988_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F16())
			var _989_v8 _dafny.Array
			_ = _989_v8
			var _nwElement0_28 _dafny.Map = _988_v7
			_ = _nwElement0_28
			var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(2))
			_ = _nw142
			(_nw142).ArraySet1(_nwElement0_28, 0)
			(_nw142).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13((_this).F16(), (_this).F16(), _987_v6, globalState), _this.F13()), 1)
			_989_v8 = _nw142
			var _990_v9 _dafny.Map
			_ = _990_v9
			_990_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_989_v8, (_980_v0).F15())
			var _991_v10 _dafny.Array
			_ = _991_v10
			var _nwElement0_29 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _nwElement0_29
			var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(9))
			_ = _nw143
			(_nw143).ArraySet1(_nwElement0_29, 0)
			(_nw143).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_this).F16() {
					return _985_v4
				}
				return _985_v4
			})()).Cardinality()), 1)
			(_nw143).ArraySet1((_980_v0).F15(), 2)
			(_nw143).ArraySet1(Companion_Default___.SafeDivisionInt(p0, Companion_Default___.Fm13((_986_v5).Dtor_cf1(), (_this).F16(), _987_v6, globalState)), 3)
			(_nw143).ArraySet1((func() _dafny.Int {
				if (_990_v9).Contains(_989_v8) {
					return (_990_v9).Get(_989_v8).(_dafny.Int)
				}
				return p0
			})(), 4)
			(_nw143).ArraySet1(_dafny.IntOfInt64(355), 5)
			(_nw143).ArraySet1((p0).Plus((_980_v0).F15()), 6)
			(_nw143).ArraySet1(p0, 7)
			(_nw143).ArraySet1(p0, 8)
			_991_v10 = _nw143
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_991_v10), 0))
			_ = _index173
			(_991_v10).ArraySet1((_979_i0).Minus(p0), (_index173).Int())
			var _992_v11 _dafny.Int
			_ = _992_v11
			_992_v11 = _dafny.IntOfInt64(151)
			var _993_v12 T1
			_ = _993_v12
			var _nw144 *C1 = New_C1_()
			_ = _nw144
			_nw144.Ctor__((_this).F16(), _dafny.UnicodeSeqOfUtf8Bytes("luf"), _991_v10)
			_993_v12 = _nw144
			var _994_v13 _dafny.MultiSet
			_ = _994_v13
			_994_v13 = _dafny.MultiSetOf((_this).F16(), (_this).F16())
			var _995_v14 _dafny.Sequence
			_ = _995_v14
			_995_v14 = _dafny.SeqOf(_993_v12)
			var _996_v15 D2
			_ = _996_v15
			_996_v15 = Companion_D2_.Create_DC7_(_983_v2, _992_v11, _dafny.IntOfInt64(518), (_this).F16())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_991_v10), 0))
			_ = _index174
			var _rhs125 _dafny.Sequence = _984_v3
			_ = _rhs125
			var _rhs126 bool = (_dafny.IntOfInt64(560)).Cmp(_dafny.IntOfInt64(947)) > 0
			_ = _rhs126
			var _rhs127 _dafny.Int = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_992_v11), _992_v11)).Minus((_980_v0).F15())
			_ = _rhs127
			var _rhs128 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (_979_i0).Plus((_993_v12).Fm7((_980_v0).F15(), (_994_v13).Cardinality(), true, globalState)))
			_ = _rhs128
			var _rhs129 T1 = (_995_v14).Select((Companion_Default___.SafeIndex(((_980_v0).F15()).Minus((_996_v15).Dtor_cf13()), _dafny.IntOfUint32((_995_v14).Cardinality()))).Uint32()).(T1)
			_ = _rhs129
			var _lhs75 *C4 = _this
			_ = _lhs75
			var _lhs76 _dafny.Array = _991_v10
			_ = _lhs76
			var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_991_v10), 0))
			_ = _lhs77
			_984_v3 = _rhs125
			_lhs75.F13_set_(_rhs126)
			(_lhs76).ArraySet1(_rhs127, (_lhs77).Int())
			_992_v11 = _rhs128
			_993_v12 = _rhs129
		}
		var _997_v16 _dafny.CodePoint
		_ = _997_v16
		_997_v16 = _dafny.CodePoint('n')
		var _998_v17 _dafny.Sequence
		_ = _998_v17
		_998_v17 = _dafny.SeqOf(_997_v16, Companion_Default___.Fm22(_this.F13(), _this.F13(), globalState))
		var _999_v18 _dafny.Array
		_ = _999_v18
		var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw145
		_999_v18 = _nw145
		var _1000_v19 *C1
		_ = _1000_v19
		var _nw146 *C1 = New_C1_()
		_ = _nw146
		_nw146.Ctor__(true, _998_v17, _999_v18)
		_1000_v19 = _nw146
		var _1001_v20 _dafny.Sequence
		_ = _1001_v20
		_1001_v20 = _dafny.SeqOf(_1000_v19)
		var _1002_v21 D0
		_ = _1002_v21
		_1002_v21 = Companion_D0_.Create_DC1_(_this.F13(), _dafny.IntOfUint32((_1001_v20).Cardinality()))
		var _1003_v22 D2
		_ = _1003_v22
		_1003_v22 = Companion_D2_.Create_DC7_(_998_v17, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F13())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(18))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}((func(_1004_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1005_i2 _dafny.Int) _dafny.Int {
				return _1004_p0
			}
		})(p0)))).Cardinality()), !((_1002_v21).Dtor_cf1()))
		var _1006_v25 _dafny.Map
		_ = _1006_v25
		_1006_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(393), _1003_v22)
		var _1007_v27 _dafny.Map
		_ = _1007_v27
		_1007_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality(), p0)
		var _1008_v28 _dafny.Array
		_ = _1008_v28
		var _nwElement0_30 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22)
		_ = _nwElement0_30
		var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(29))
		_ = _nw147
		(_nw147).ArraySet1(_nwElement0_30, 0)
		(_nw147).ArraySet1(func() _dafny.Map {
			var _coll38 = _dafny.NewMapBuilder()
			_ = _coll38
			for _iter40 := _dafny.Iterate((func() _dafny.Map {
				var _coll39 = _dafny.NewMapBuilder()
				_ = _coll39
				for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(457), _dafny.IntOfInt64(35))); ; {
					_compr_39, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _1009_v24 _dafny.Int
					_1009_v24 = interface{}(_compr_39).(_dafny.Int)
					if ((_dafny.IntOfInt64(457)).Cmp(_1009_v24) <= 0) && ((_1009_v24).Cmp(_dafny.IntOfInt64(35)) < 0) {
						_coll39.Add((_1009_v24).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(p0)).Cardinality(), (_1000_v19).F17())).Cardinality()), p0)
					}
				}
				return _coll39.ToMap()
			}()).Keys().Elements()); ; {
				_compr_38, _ok40 := _iter40()
				if !_ok40 {
					break
				}
				var _1010_v23 _dafny.Int
				_1010_v23 = interface{}(_compr_38).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(457), _dafny.IntOfInt64(35))); ; {
						_compr_40, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _1009_v24 _dafny.Int
						_1009_v24 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(457)).Cmp(_1009_v24) <= 0) && ((_1009_v24).Cmp(_dafny.IntOfInt64(35)) < 0) {
							_coll40.Add((_1009_v24).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(p0)).Cardinality(), (_1000_v19).F17())).Cardinality()), p0)
						}
					}
					return _coll40.ToMap()
				}()).Contains(_1010_v23) {
					_coll38.Add(Companion_Default___.SafeModuloInt(_1010_v23, p0), _1003_v22)
				}
			}
			return _coll38.ToMap()
		}(), 1)
		(_nw147).ArraySet1(_1006_v25, 2)
		(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22), 3)
		(_nw147).ArraySet1(_1006_v25, 4)
		(_nw147).ArraySet1(func() _dafny.Map {
			var _coll41 = _dafny.NewMapBuilder()
			_ = _coll41
			for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(500), _dafny.IntOfInt64(-712))); ; {
				_compr_41, _ok43 := _iter43()
				if !_ok43 {
					break
				}
				var _1011_v26 _dafny.Int
				_1011_v26 = interface{}(_compr_41).(_dafny.Int)
				if ((_dafny.IntOfInt64(500)).Cmp(_1011_v26) <= 0) && ((_1011_v26).Cmp(_dafny.IntOfInt64(-712)) < 0) {
					_coll41.Add(Companion_Default___.SafeModuloInt(_1011_v26, p0), _1003_v22)
				}
			}
			return _coll41.ToMap()
		}(), 5)
		(_nw147).ArraySet1(_1006_v25, 6)
		(_nw147).ArraySet1(_1006_v25, 7)
		(_nw147).ArraySet1(_1006_v25, 8)
		(_nw147).ArraySet1(_1006_v25, 9)
		(_nw147).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22)).Update(p0, _1003_v22), 10)
		(_nw147).ArraySet1(_1006_v25, 11)
		(_nw147).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm23(_dafny.IntOfUint32((_1000_v19.F18).Cardinality()), (_this).F16(), globalState))).Merge(_1006_v25), 12)
		(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22), 13)
		(_nw147).ArraySet1(_1006_v25, 14)
		(_nw147).ArraySet1((_1006_v25).Merge(_1006_v25), 15)
		(_nw147).ArraySet1((_1006_v25).Merge(_1006_v25), 16)
		(_nw147).ArraySet1(_1006_v25, 17)
		(_nw147).ArraySet1(_1006_v25, 18)
		(_nw147).ArraySet1(_1006_v25, 19)
		(_nw147).ArraySet1(_1006_v25, 20)
		(_nw147).ArraySet1(_1006_v25, 21)
		(_nw147).ArraySet1((_1006_v25).Merge(_1006_v25), 22)
		(_nw147).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22)).Update(p0, _1003_v22), 23)
		(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22), 24)
		(_nw147).ArraySet1((_1006_v25).Update(p0, Companion_D2_.Create_DC7_(_1000_v19.F18, p0, (_dafny.SetOf((_this).F16(), (_this).Fm4((_this).F16(), _1002_v21, p0, _1007_v27, globalState))).Cardinality(), (_this).F16())), 25)
		(_nw147).ArraySet1(_1006_v25, 26)
		(_nw147).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22)).Merge(_1006_v25), 27)
		(_nw147).ArraySet1(_1006_v25, 28)
		_1008_v28 = _nw147
		var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1008_v28), 0))
		_ = _index175
		(_1008_v28).ArraySet1(_1006_v25, (_index175).Int())
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1008_v28), 0))
		_ = _index176
		(_1008_v28).ArraySet1(((_1006_v25).Merge(_1006_v25)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1003_v22)), (_index176).Int())
		var _1012_v29 _dafny.Map
		_ = _1012_v29
		_1012_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1003_v22).Dtor_cf15(), p0)
		(_this).F13_set_((p0).Cmp(((_1012_v29).Merge(_1012_v29)).Cardinality()) <= 0)
		var _1013_v30 _dafny.Int
		_ = _1013_v30
		_1013_v30 = _dafny.IntOfInt64(888)
		_1013_v30 = _1013_v30
		_1013_v30 = (Companion_Default___.SafeModuloInt(p0, p0)).Plus(_dafny.IntOfInt64(-301))
		var _hi9 _dafny.Int = p0
		_ = _hi9
		for _1014_i3 := _1013_v30; _1014_i3.Cmp(_hi9) < 0; _1014_i3 = _1014_i3.Plus(_dafny.One) {
			var _1015_v31 _dafny.Sequence
			_ = _1015_v31
			_1015_v31 = _dafny.SeqOf(p0)
			(_this).F13_set_(((_1012_v29).Cardinality()).Cmp((_1015_v31).Select((Companion_Default___.SafeIndex(_1013_v30, _dafny.IntOfUint32((_1015_v31).Cardinality()))).Uint32()).(_dafny.Int)) > 0)
			var _1016_v32 _dafny.Sequence
			_ = _1016_v32
			_1016_v32 = _dafny.SeqOf((_1000_v19).F17())
			var _rhs130 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1016_v32, _1016_v32)
			_ = _rhs130
			var _rhs131 _dafny.Int = (_dafny.MultiSetOf((func() _dafny.Int {
				if (_1007_v27).Contains((_1007_v27).Cardinality()) {
					return (_1007_v27).Get((_1007_v27).Cardinality()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_1001_v20).Cardinality())
			})(), _1014_i3, p0)).Cardinality()
			_ = _rhs131
			_1016_v32 = _rhs130
			_1013_v30 = _rhs131
			(_this).F13_set_(_this.F13())
			var _1017_v33 *C1
			_ = _1017_v33
			var _nw148 *C1 = New_C1_()
			_ = _nw148
			_nw148.Ctor__((_1000_v19).Fm14(p0, (_1000_v19).F17(), _1014_i3, _1013_v30, globalState), _1000_v19.F18, _999_v18)
			_1017_v33 = _nw148
		}
		r0 = _999_v18
		return r0
	}
}
func (_this *C4) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Sequence) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _1018_v0 _dafny.Map
		_ = _1018_v0
		_1018_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.UnicodeSeqOfUtf8Bytes("mvnyliga"))
		var _hi10 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1018_v0).Contains(p2) {
				return (_1018_v0).Get(p2).(_dafny.Sequence)
			}
			return p1
		})()).Cardinality())
		_ = _hi10
		for _1019_i0 := p2; _1019_i0.Cmp(_hi10) < 0; _1019_i0 = _1019_i0.Plus(_dafny.One) {
			var _1020_v1 _dafny.Int
			_ = _1020_v1
			_1020_v1 = _dafny.IntOfInt64(-188)
			var _1021_v2 _dafny.MultiSet
			_ = _1021_v2
			_1021_v2 = _dafny.MultiSetOf(_this.F13())
			var _rhs132 _dafny.Int = (_1021_v2).Cardinality()
			_ = _rhs132
			var _rhs133 _dafny.Int = ((Companion_Default___.Fm24(_1020_v1, _this.F13(), globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _1019_i0))).Cardinality()
			_ = _rhs133
			var _rhs134 bool = (_this.F13()) == ((_this).F16())
			_ = _rhs134
			var _rhs135 bool = _this.F13()
			_ = _rhs135
			var _lhs78 *C4 = _this
			_ = _lhs78
			_1020_v1 = _rhs132
			_1020_v1 = _rhs133
			r1 = _rhs134
			_lhs78.F13_set_(_rhs135)
			var _1022_v3 _dafny.Array
			_ = _1022_v3
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw149
			_1022_v3 = _nw149
			_1022_v3 = _1022_v3
			var _1023_v4 _dafny.Array
			_ = _1023_v4
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
			_ = _nw150
			_1023_v4 = _nw150
			var _1024_v5 _dafny.Array
			_ = _1024_v5
			var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
			_ = _nw151
			_1024_v5 = _nw151
			var _1025_v6 _dafny.Map
			_ = _1025_v6
			_1025_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus(_1019_i0))
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1024_v5), 0))
			_ = _index177
			(_1024_v5).ArraySet1((_1025_v6).Update(p0, p2), (_index177).Int())
			var _1026_v7 _dafny.CodePoint
			_ = _1026_v7
			_1026_v7 = _dafny.CodePoint('s')
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1024_v5), 0))
			_ = _index178
			var _rhs136 _dafny.Array = _1023_v4
			_ = _rhs136
			var _rhs137 _dafny.Map = (_1025_v6).Update(Companion_Default___.Fm13((_this).F16(), true, _1026_v7, globalState), p2)
			_ = _rhs137
			var _rhs138 bool = !(_this.F13())
			_ = _rhs138
			var _rhs139 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.Companion_Sequence_.Concatenate(p1, p1))).Cardinality())
			_ = _rhs139
			var _rhs140 _dafny.Int = _1020_v1
			_ = _rhs140
			var _lhs79 _dafny.Array = _1024_v5
			_ = _lhs79
			var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1024_v5), 0))
			_ = _lhs80
			var _lhs81 *C4 = _this
			_ = _lhs81
			_1023_v4 = _rhs136
			(_lhs79).ArraySet1(_rhs137, (_lhs80).Int())
			_lhs81.F13_set_(_rhs138)
			_1020_v1 = _rhs139
			_1020_v1 = _rhs140
			var _1027_v8 _dafny.Map
			_ = _1027_v8
			_1027_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F16()) == (!((_this).F16())), _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1026_v7))
			var _1028_v9 _dafny.Map
			_ = _1028_v9
			_1028_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_1019_i0)).Cardinality(), _dafny.SetOf(_1020_v1))
			var _1029_v10 _dafny.Array
			_ = _1029_v10
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_24
			var _nw152 _dafny.Array
			_ = _nw152
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw152 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Int = (func(_1030_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1031_i1 _dafny.Int) _dafny.Int {
						return (_1031_i1).Times(_1030_p2)
					}
				})(p2)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw152 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw152).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw152).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1029_v10 = _nw152
			var _1032_v11 _dafny.Map
			_ = _1032_v11
			_1032_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1028_v9).Cardinality(), _1029_v10)
			_1027_v8 = (_1027_v8).Update(((_1032_v11).Cardinality()).Cmp(_dafny.IntOfInt64(-440)) < 0, _dafny.UnicodeSeqOfUtf8Bytes("dct"))
		}
		var _1033_v12 _dafny.Array
		_ = _1033_v12
		var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw153
		_1033_v12 = _nw153
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))
		_ = _index179
		(_1033_v12).ArraySet1(false, (_index179).Int())
		var _1034_v13 _dafny.Int
		_ = _1034_v13
		_1034_v13 = _dafny.IntOfInt64(539)
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_1033_v12), 0))
		_ = _index180
		(_1033_v12).ArraySet1((_this).F16(), (_index180).Int())
		var _1035_v14 _dafny.Set
		_ = _1035_v14
		_1035_v14 = _dafny.SetOf(true, _this.F13())
		var _1036_v15 _dafny.Set
		_ = _1036_v15
		_1036_v15 = _dafny.SetOf(_1034_v13)
		var _1037_v16 _dafny.Map
		_ = _1037_v16
		_1037_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _this.F13())
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))
		_ = _index181
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_1033_v12), 0))
		_ = _index182
		var _rhs141 bool = _this.F13()
		_ = _rhs141
		var _rhs142 _dafny.Int = Companion_Default___.SafeDivisionInt(_1034_v13, _dafny.IntOfInt64(-663))
		_ = _rhs142
		var _rhs143 bool = !((_1035_v14).Union(Companion_Default___.Fm25(_1036_v15, _1037_v16, p0, _this.F13(), globalState))).Equals((_1035_v14).Difference(_1035_v14))
		_ = _rhs143
		var _rhs144 bool = _this.F13()
		_ = _rhs144
		var _lhs82 _dafny.Array = _1033_v12
		_ = _lhs82
		var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))
		_ = _lhs83
		var _lhs84 _dafny.Array = _1033_v12
		_ = _lhs84
		var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_1033_v12), 0))
		_ = _lhs85
		var _lhs86 *C4 = _this
		_ = _lhs86
		(_lhs82).ArraySet1(_rhs141, (_lhs83).Int())
		_1034_v13 = _rhs142
		(_lhs84).ArraySet1(_rhs143, (_lhs85).Int())
		_lhs86.F13_set_(_rhs144)
		_1034_v13 = _dafny.IntOfInt64(326)
		var _1038_v17 D1
		_ = _1038_v17
		_1038_v17 = Companion_D1_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("arq"), _1033_v12, _dafny.IntOfInt64(764))
		var _1039_v18 D4
		_ = _1039_v18
		_1039_v18 = Companion_D4_.Create_DC11_((_this).F16())
		var _1040_v19 _dafny.CodePoint
		_ = _1040_v19
		_1040_v19 = _dafny.CodePoint('m')
		var _1041_v20 _dafny.MultiSet
		_ = _1041_v20
		_1041_v20 = _dafny.MultiSetOf(_1034_v13, (_1035_v14).Cardinality(), p0)
		var _1042_v21 _dafny.Map
		_ = _1042_v21
		_1042_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(_1034_v13))
		var _1043_v22 _dafny.Array
		_ = _1043_v22
		var _len0_25 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_25
		var _nw154 _dafny.Array
		_ = _nw154
		if _len0_25.Cmp(_dafny.Zero) == 0 {
			_nw154 = _dafny.NewArray(_len0_25)
		} else {
			var _init25 func(_dafny.Int) _dafny.Int = func(_1044_i2 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_1044_i2, _dafny.IntOfInt64(836))
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
		_1043_v22 = _nw154
		var _1045_v23 _dafny.Map
		_ = _1045_v23
		_1045_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _1034_v13)
		var _1046_v24 D5
		_ = _1046_v24
		_1046_v24 = Companion_D5_.Create_DC12_(_1045_v23)
		var _1047_v25 _dafny.Map
		_ = _1047_v25
		_1047_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1043_v22, ((_1046_v24).Dtor_cf22()).Cardinality())
		var _1048_v26 _dafny.Sequence
		_ = _1048_v26
		_1048_v26 = _dafny.SeqOf((_this).F16())
		var _1049_v27 _dafny.Sequence
		_ = _1049_v27
		_1049_v27 = _dafny.SeqOf(_1034_v13, p0)
		var _1050_v28 _dafny.Array
		_ = _1050_v28
		var _nwElement0_31 _dafny.Int = (_1038_v17).Dtor_cf10()
		_ = _nwElement0_31
		var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(26))
		_ = _nw155
		(_nw155).ArraySet1(_nwElement0_31, 0)
		(_nw155).ArraySet1(_dafny.IntOfInt64(168), 1)
		(_nw155).ArraySet1(p0, 2)
		(_nw155).ArraySet1((_dafny.IntOfUint32((p1).Cardinality())).Plus(_dafny.IntOfInt64(590)), 3)
		(_nw155).ArraySet1(p0, 4)
		(_nw155).ArraySet1(Companion_Default___.Fm13(!((_1039_v18).Dtor_cf21()), _this.F13(), _1040_v19, globalState), 5)
		(_nw155).ArraySet1(p0, 6)
		(_nw155).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_1041_v20).Cardinality()), (_1042_v21).Cardinality())), 7)
		(_nw155).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm13((_this).F16(), _this.F13(), _dafny.CodePoint('d'), globalState), (_dafny.Zero).Minus(p0)), 8)
		(_nw155).ArraySet1(p0, 9)
		(_nw155).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p2, (Companion_Default___.Fm26(_1040_v19, p0, _1034_v13, globalState)).Cardinality())), 10)
		(_nw155).ArraySet1(p0, 11)
		(_nw155).ArraySet1((func() _dafny.Int {
			if (_1047_v25).Contains(_1043_v22) {
				return (_1047_v25).Get(_1043_v22).(_dafny.Int)
			}
			return _dafny.IntOfUint32((p1).Cardinality())
		})(), 12)
		(_nw155).ArraySet1(Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfUint32((_1048_v26).Cardinality())), 13)
		(_nw155).ArraySet1(p2, 14)
		(_nw155).ArraySet1(p0, 15)
		(_nw155).ArraySet1(_dafny.IntOfInt64(532), 16)
		(_nw155).ArraySet1(p2, 17)
		(_nw155).ArraySet1(p0, 18)
		(_nw155).ArraySet1((_1034_v13).Plus(_1034_v13), 19)
		(_nw155).ArraySet1((_dafny.IntOfInt64(707)).Times(p2), 20)
		(_nw155).ArraySet1((_1034_v13).Plus(_1034_v13), 21)
		(_nw155).ArraySet1(p2, 22)
		(_nw155).ArraySet1((_dafny.Zero).Minus(((_1049_v27).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1049_v27).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p2)), 23)
		(_nw155).ArraySet1(_dafny.IntOfInt64(333), 24)
		(_nw155).ArraySet1(((_dafny.MultiSetOf(_1040_v19, _1040_v19)).Update(_1040_v19, Companion_Default___.Abs(_1034_v13))).Cardinality(), 25)
		_1050_v28 = _nw155
		var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1043_v22), 0))
		_ = _index183
		(_1043_v22).ArraySet1((p2).Times(_dafny.IntOfUint32((_1049_v27).Cardinality())), (_index183).Int())
		var _1051_v29 _dafny.Map
		_ = _1051_v29
		_1051_v29 = _1037_v16
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1043_v22), 0))
		_ = _index184
		var _rhs145 _dafny.Int = p0
		_ = _rhs145
		var _rhs146 _dafny.Map = _1051_v29
		_ = _rhs146
		var _lhs87 _dafny.Array = _1043_v22
		_ = _lhs87
		var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1043_v22), 0))
		_ = _lhs88
		(_lhs87).ArraySet1(_rhs145, (_lhs88).Int())
		_1051_v29 = _rhs146
		var _1052_v30 _dafny.Array
		_ = _1052_v30
		var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
		_ = _nw156
		_1052_v30 = _nw156
		var _1053_v31 _dafny.Map
		_ = _1053_v31
		_1053_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v26, (_1045_v23).Update((_1033_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))).Int()).(bool), _1034_v13))
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1052_v30), 0))
		_ = _index185
		(_1052_v30).ArraySet1(_1053_v31, (_index185).Int())
		var _1054_v32 D6
		_ = _1054_v32
		_1054_v32 = Companion_D6_.Create_DC16_(_1053_v31)
		var _1055_v33 _dafny.Map
		_ = _1055_v33
		_1055_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1033_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))).Int()).(bool), p1)
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))
		_ = _index186
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1052_v30), 0))
		_ = _index187
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1043_v22), 0))
		_ = _index188
		var _rhs147 bool = !(_this.F13()) || ((_1048_v26).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1048_v26).Cardinality()))).Uint32()).(bool))
		_ = _rhs147
		var _rhs148 _dafny.Map = (_1054_v32).Dtor_cf25()
		_ = _rhs148
		var _rhs149 _dafny.Int = (_dafny.Zero).Minus((p0).Minus(((_1043_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1043_v22), 0))).Int()).(_dafny.Int)).Plus((_1055_v33).Cardinality())))
		_ = _rhs149
		var _rhs150 _dafny.Int = (p2).Plus(_dafny.IntOfInt64(534))
		_ = _rhs150
		var _rhs151 _dafny.Array = _1050_v28
		_ = _rhs151
		var _lhs89 _dafny.Array = _1033_v12
		_ = _lhs89
		var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))
		_ = _lhs90
		var _lhs91 _dafny.Array = _1052_v30
		_ = _lhs91
		var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_1052_v30), 0))
		_ = _lhs92
		var _lhs93 _dafny.Array = _1043_v22
		_ = _lhs93
		var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1043_v22), 0))
		_ = _lhs94
		(_lhs89).ArraySet1(_rhs147, (_lhs90).Int())
		(_lhs91).ArraySet1(_rhs148, (_lhs92).Int())
		(_lhs93).ArraySet1(_rhs149, (_lhs94).Int())
		_1034_v13 = _rhs150
		_1043_v22 = _rhs151
		_1035_v14 = _1035_v14
		var _1056_v34 _dafny.Sequence
		_ = _1056_v34
		_1056_v34 = _dafny.SeqOf(_1033_v12)
		var _1057_v35 _dafny.Map
		_ = _1057_v35
		_1057_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1037_v16).Cardinality(), (_1056_v34).Select((Companion_Default___.SafeIndex(_1034_v13, _dafny.IntOfUint32((_1056_v34).Cardinality()))).Uint32()).(_dafny.Array))
		r0 = (func() _dafny.Map {
			if (_1033_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))).Int()).(bool) {
				return _1057_v35
			}
			return _1057_v35
		})()
		r1 = (func() bool {
			if _this.F13() {
				return (_1048_v26).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_1048_v26)).Cardinality(), _dafny.IntOfUint32((_1048_v26).Cardinality()))).Uint32()).(bool)
			}
			return (_1033_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1033_v12), 0))).Int()).(bool)
		})()
		var _1058_v36 _dafny.MultiSet
		_ = _1058_v36
		_1058_v36 = _dafny.MultiSetOf((_this).F16(), _this.F13(), !((_this).F16()))
		r2 = ((_1058_v36).Difference(_1058_v36)).IsSubsetOf(_1058_v36)
		r3 = _1048_v26
		return r0, r1, r2, r3
	}
}
func (_this *C4) M4(p0 bool, p1 bool, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1059_v0 D4
		_ = _1059_v0
		_1059_v0 = Companion_D4_.Create_DC11_(p1)
		var _1060_v1 _dafny.CodePoint
		_ = _1060_v1
		_1060_v1 = _dafny.CodePoint('o')
		var _1061_v2 _dafny.Sequence
		_ = _1061_v2
		_1061_v2 = _dafny.SeqOf(_1060_v1)
		var _1062_v3 _dafny.Int
		_ = _1062_v3
		_1062_v3 = _dafny.IntOfInt64(956)
		var _1063_v4 D2
		_ = _1063_v4
		_1063_v4 = Companion_D2_.Create_DC7_(_1061_v2, _1062_v3, _dafny.IntOfInt64(593), p0)
		var _pat_let_tv23 = _1063_v4
		_ = _pat_let_tv23
		var _source18 D4 = func(_pat_let26_0 D4) D4 {
			return func(_1064_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let27_0 bool) D4 {
					return func(_1065_dt__update_hcf21_h0 bool) D4 {
						return Companion_D4_.Create_DC11_(_1065_dt__update_hcf21_h0)
					}(_pat_let27_0)
				}((_pat_let_tv23).Dtor_cf15())
			}(_pat_let26_0)
		}(_1059_v0)
		_ = _source18
		if _source18.Is_DC11() {
			var _1066___mcc_h0 bool = _source18.Get_().(D4_DC11).Cf21
			_ = _1066___mcc_h0
			var _1067_cf21 bool = _1066___mcc_h0
			_ = _1067_cf21
			var _1068_v5 _dafny.Array
			_ = _1068_v5
			var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw157
			_1068_v5 = _nw157
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1068_v5), 0))
			_ = _index189
			(_1068_v5).ArraySet1(_1062_v3, (_index189).Int())
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1068_v5), 0))
			_ = _index190
			(_1068_v5).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm13(_1067_cf21, _1067_cf21, _1060_v1, globalState), (_dafny.Zero).Minus((_dafny.Zero).Minus((_1062_v3).Times(_1062_v3)))), (_index190).Int())
			var _1069_v6 _dafny.Set
			_ = _1069_v6
			_1069_v6 = _dafny.SetOf(p1, _this.F13(), _1067_cf21)
			_1069_v6 = _1069_v6
			var _1070_v7 _dafny.Map
			_ = _1070_v7
			_1070_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), !(p1))
			_1070_v7 = (_1070_v7).Update(p1, _1067_cf21)
			var _1071_v8 *C1
			_ = _1071_v8
			var _nw158 *C1 = New_C1_()
			_ = _nw158
			_nw158.Ctor__(_1067_cf21, _dafny.Companion_Sequence_.Update(_1061_v2, (Companion_Default___.SafeIndex((_1070_v7).Cardinality(), _dafny.IntOfUint32((_1061_v2).Cardinality()))).Uint32(), _dafny.CodePoint('c')), _1068_v5)
			_1071_v8 = _nw158
		} else {
			var _1072___mcc_h1 _dafny.Array = _source18.Get_().(D4_DC10).Cf20
			_ = _1072___mcc_h1
			var _1073_cf20 _dafny.Array = _1072___mcc_h1
			_ = _1073_cf20
			var _1074_v9 _dafny.Array
			_ = _1074_v9
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_26
			var _nw159 _dafny.Array
			_ = _nw159
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw159 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) bool = (func(_1075_p1 bool) func(_dafny.Int) bool {
					return func(_1076_i0 _dafny.Int) bool {
						return _1075_p1
					}
				})(p1)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw159 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw159).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw159).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1074_v9 = _nw159
			var _1077_v10 D1
			_ = _1077_v10
			_1077_v10 = Companion_D1_.Create_DC4_(_1074_v9, _1062_v3, Companion_Default___.Fm22((_this).F16(), p0, globalState))
			var _1078_v11 _dafny.Array
			_ = _1078_v11
			var _nwElement0_32 _dafny.CodePoint = _1060_v1
			_ = _nwElement0_32
			var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(28))
			_ = _nw160
			(_nw160).ArraySet1CodePoint(_nwElement0_32, 0)
			(_nw160).ArraySet1CodePoint(_1060_v1, 1)
			(_nw160).ArraySet1CodePoint(_1060_v1, 2)
			(_nw160).ArraySet1CodePoint((_1077_v10).Dtor_cf7(), 3)
			(_nw160).ArraySet1CodePoint(_1060_v1, 4)
			(_nw160).ArraySet1CodePoint(_1060_v1, 5)
			(_nw160).ArraySet1CodePoint(_1060_v1, 6)
			(_nw160).ArraySet1CodePoint(_dafny.CodePoint('u'), 7)
			(_nw160).ArraySet1CodePoint(_dafny.CodePoint('k'), 8)
			(_nw160).ArraySet1CodePoint(_1060_v1, 9)
			(_nw160).ArraySet1CodePoint(_1060_v1, 10)
			(_nw160).ArraySet1CodePoint(_dafny.CodePoint('o'), 11)
			(_nw160).ArraySet1CodePoint(_dafny.CodePoint('p'), 12)
			(_nw160).ArraySet1CodePoint(_1060_v1, 13)
			(_nw160).ArraySet1CodePoint(_1060_v1, 14)
			(_nw160).ArraySet1CodePoint(_dafny.CodePoint('s'), 15)
			(_nw160).ArraySet1CodePoint(_1060_v1, 16)
			(_nw160).ArraySet1CodePoint(_1060_v1, 17)
			(_nw160).ArraySet1CodePoint((_1061_v2).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1061_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), 18)
			(_nw160).ArraySet1CodePoint(_1060_v1, 19)
			(_nw160).ArraySet1CodePoint((_1061_v2).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1061_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), 20)
			(_nw160).ArraySet1CodePoint(_1060_v1, 21)
			(_nw160).ArraySet1CodePoint(_1060_v1, 22)
			(_nw160).ArraySet1CodePoint(_1060_v1, 23)
			(_nw160).ArraySet1CodePoint(_1060_v1, 24)
			(_nw160).ArraySet1CodePoint(_1060_v1, 25)
			(_nw160).ArraySet1CodePoint((_1077_v10).Dtor_cf7(), 26)
			(_nw160).ArraySet1CodePoint(_1060_v1, 27)
			_1078_v11 = _nw160
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1078_v11), 0))
			_ = _index191
			(_1078_v11).ArraySet1CodePoint(Companion_Default___.Fm22(_this.F13(), p0, globalState), (_index191).Int())
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1078_v11), 0))
			_ = _index192
			var _rhs152 bool = (!((func() bool {
				if _this.F13() {
					return (_this).F16()
				}
				return p1
			})())) || ((p1) && ((_this).F16()))
			_ = _rhs152
			var _rhs153 _dafny.CodePoint = _1060_v1
			_ = _rhs153
			var _lhs95 *C4 = _this
			_ = _lhs95
			var _lhs96 _dafny.Array = _1078_v11
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1078_v11), 0))
			_ = _lhs97
			_lhs95.F13_set_(_rhs152)
			(_lhs96).ArraySet1CodePoint(_rhs153, (_lhs97).Int())
			var _1079_v12 _dafny.Map
			_ = _1079_v12
			_1079_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v3, _1062_v3)
			var _1080_v13 T0
			_ = _1080_v13
			var _nw161 *C2 = New_C2_()
			_ = _nw161
			_nw161.Ctor__(true)
			_1080_v13 = _nw161
			r1 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1079_v12, _1080_v13)).Cardinality()).Plus(_dafny.IntOfUint32((_1061_v2).Cardinality()))).Cmp(_dafny.IntOfInt64(459)) > 0
			var _1081_v14 _dafny.Map
			_ = _1081_v14
			_1081_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v3, (_this).F16())
			var _1082_v15 _dafny.Map
			_ = _1082_v15
			_1082_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_1081_v14).Contains((_dafny.Zero).Minus(_1062_v3)) {
					return (_1081_v14).Get((_dafny.Zero).Minus(_1062_v3)).(bool)
				}
				return p0
			})(), Companion_Default___.SafeDivisionInt(_1062_v3, _1062_v3))
			_1082_v15 = (_1082_v15).Update(_1080_v13.F13(), Companion_Default___.SafeDivisionInt(_1062_v3, _1062_v3))
			_1081_v14 = _1081_v14
		}
		(_this).F13_set_(_this.F13())
		var _1083_v16 _dafny.Sequence
		_ = _1083_v16
		_1083_v16 = _dafny.SeqOf(!(_this.F13()), p1, _this.F13(), (_this).F16())
		var _1084_v17 _dafny.Array
		_ = _1084_v17
		var _len0_27 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_27
		var _nw162 _dafny.Array
		_ = _nw162
		if _len0_27.Cmp(_dafny.Zero) == 0 {
			_nw162 = _dafny.NewArray(_len0_27)
		} else {
			var _init27 func(_dafny.Int) bool = func(_1085_i1 _dafny.Int) bool {
				return _this.F13()
			}
			_ = _init27
			var _element0_27 = _init27(_dafny.Zero)
			_ = _element0_27
			_nw162 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
			(_nw162).ArraySet1(_element0_27, 0)
			var _nativeLen0_27 = (_len0_27).Int()
			_ = _nativeLen0_27
			for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
				(_nw162).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
			}
		}
		_1084_v17 = _nw162
		var _1086_v18 _dafny.Map
		_ = _1086_v18
		_1086_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_v16, _1084_v17)
		_1086_v18 = (_1086_v18).Update(_dafny.SeqOf((_1083_v16).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32()).(bool), (_this).F16(), !((_this).F16())), _1084_v17)
		var _1087_i2 _dafny.Int
		_ = _1087_i2
		_1087_i2 = _dafny.Zero
		{
			for p1 {
				{
					if (_1087_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1087_i2 = (_1087_i2).Plus(_dafny.One)
					if (_this).F16() {
						var _1088_v19 _dafny.Map
						_ = _1088_v19
						_1088_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _1062_v3)
						var _1089_v20 _dafny.Map
						_ = _1089_v20
						_1089_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_1088_v19).Cardinality())
						var _1090_v21 _dafny.Set
						_ = _1090_v21
						_1090_v21 = _dafny.SetOf(true, !(p0))
						var _1091_v22 _dafny.Set
						_ = _1091_v22
						_1091_v22 = _dafny.SetOf(_1062_v3, (_1090_v21).Cardinality())
						var _1092_v23 _dafny.Map
						_ = _1092_v23
						_1092_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), p0)
						var _rhs154 bool = (_1083_v16).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_1062_v3, _1062_v3), _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32()).(bool)
						_ = _rhs154
						var _rhs155 _dafny.Int = ((func() _dafny.Int {
							if p0 {
								return (_1088_v19).Cardinality()
							}
							return _1062_v3
						})()).Times(Companion_Default___.SafeDivisionInt(_1062_v3, _1062_v3))
						_ = _rhs155
						var _rhs156 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1062_v3, _dafny.IntOfInt64(-641), _dafny.IntOfUint32((_1061_v2).Cardinality()), _1062_v3), Companion_Default___.Fm25(_1091_v22, _1092_v23, _dafny.IntOfUint32((_dafny.SeqOf(_this.F13(), false, (_this).F16())).Cardinality()), p0, globalState))).Cardinality()
						_ = _rhs156
						r0 = _rhs154
						_1062_v3 = _rhs155
						r2 = _rhs156
						var _1093_v24 D0
						_ = _1093_v24
						_1093_v24 = Companion_D0_.Create_DC1_((_this).F16(), _dafny.IntOfInt64(-634))
						var _1094_v25 _dafny.Map
						_ = _1094_v25
						_1094_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1062_v3), _1062_v3)
						var _1095_v26 _dafny.Map
						_ = _1095_v26
						_1095_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1061_v2).Cardinality()), (_1094_v25).Cardinality())
						var _1096_v27 _dafny.Sequence
						_ = _1096_v27
						_1096_v27 = _dafny.SeqOf(_1062_v3)
						_1060_v1 = Companion_Default___.Fm22((_this).Fm4((_1083_v16).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32()).(bool), _1093_v24, _1062_v3, (_1094_v25).Update(_1062_v3, _1062_v3), globalState), (_1062_v3).Cmp((_1096_v27).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1096_v27).Cardinality()))).Uint32()).(_dafny.Int)) >= 0, globalState)
						var _1097_v28 _dafny.Array
						_ = _1097_v28
						var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
						_ = _nw163
						_1097_v28 = _nw163
						var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1097_v28), 0))
						_ = _index193
						(_1097_v28).ArraySet1(_1062_v3, (_index193).Int())
						var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1097_v28), 0))
						_ = _index194
						var _rhs157 bool = ((_dafny.IntOfUint32((_1061_v2).Cardinality())).Minus(_1062_v3)).Cmp((_1089_v20).Cardinality()) < 0
						_ = _rhs157
						var _rhs158 bool = _this.F13()
						_ = _rhs158
						var _rhs159 _dafny.Int = _1062_v3
						_ = _rhs159
						var _lhs98 *C4 = _this
						_ = _lhs98
						var _lhs99 _dafny.Array = _1097_v28
						_ = _lhs99
						var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1097_v28), 0))
						_ = _lhs100
						_lhs98.F13_set_(_rhs157)
						r0 = _rhs158
						(_lhs99).ArraySet1(_rhs159, (_lhs100).Int())
						_1096_v27 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(303))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg58 _dafny.Int) interface{} {
								return coer58(arg58)
							}
						}((func(_1098_v28 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_1099_i3 _dafny.Int) _dafny.Int {
								return (_1098_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1098_v28), 0))).Int()).(_dafny.Int)
							}
						})(_1097_v28)))
						var _1100_v29 _dafny.MultiSet
						_ = _1100_v29
						_1100_v29 = _dafny.MultiSetOf(((_1097_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1097_v28), 0))).Int()).(_dafny.Int)).Minus(_1062_v3), _dafny.IntOfInt64(741), (_1097_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1097_v28), 0))).Int()).(_dafny.Int), _1062_v3)
						var _1101_v30 T1
						_ = _1101_v30
						var _nw164 *C1 = New_C1_()
						_ = _nw164
						_nw164.Ctor__(p0, _dafny.UnicodeSeqOfUtf8Bytes("nldyh"), _1097_v28)
						_1101_v30 = _nw164
						var _1102_v31 _dafny.Sequence
						_ = _1102_v31
						_1102_v31 = _dafny.SeqOf(_1101_v30, _1101_v30)
						var _1103_v32 D7
						_ = _1103_v32
						_1103_v32 = Companion_D7_.Create_DC19_(_1101_v30)
						var _1104_v33 _dafny.Array
						_ = _1104_v33
						var _nwElement0_33 T1 = _1101_v30
						_ = _nwElement0_33
						var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(14))
						_ = _nw165
						(_nw165).ArraySet1(_nwElement0_33, 0)
						(_nw165).ArraySet1((_1102_v31).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1102_v31).Cardinality()))).Uint32()).(T1), 1)
						(_nw165).ArraySet1(_1101_v30, 2)
						(_nw165).ArraySet1((_1103_v32).Dtor_cf30(), 3)
						(_nw165).ArraySet1(_1101_v30, 4)
						(_nw165).ArraySet1(_1101_v30, 5)
						(_nw165).ArraySet1(_1101_v30, 6)
						(_nw165).ArraySet1(_1101_v30, 7)
						(_nw165).ArraySet1(_1101_v30, 8)
						(_nw165).ArraySet1(_1101_v30, 9)
						(_nw165).ArraySet1(_1101_v30, 10)
						(_nw165).ArraySet1(_1101_v30, 11)
						(_nw165).ArraySet1(_1101_v30, 12)
						(_nw165).ArraySet1(_1101_v30, 13)
						_1104_v33 = _nw165
						var _1105_v34 _dafny.Sequence
						_ = _1105_v34
						_1105_v34 = _dafny.SeqOf(_1104_v33)
						var _rhs160 _dafny.CodePoint = _1060_v1
						_ = _rhs160
						var _rhs161 _dafny.Sequence = (func() _dafny.Sequence {
							if p1 {
								return _1083_v16
							}
							return _1083_v16
						})()
						_ = _rhs161
						var _rhs162 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1061_v2, _1061_v2)
						_ = _rhs162
						var _rhs163 _dafny.MultiSet = (_1100_v29).Update(_1062_v3, Companion_Default___.Abs((_1090_v21).Cardinality()))
						_ = _rhs163
						var _rhs164 _dafny.Sequence = _1105_v34
						_ = _rhs164
						_1060_v1 = _rhs160
						_1083_v16 = _rhs161
						r1 = _rhs162
						_1100_v29 = _rhs163
						_1105_v34 = _rhs164
					} else {
						var _1106_v35 _dafny.Array
						_ = _1106_v35
						var _nwElement0_34 _dafny.Int = _1062_v3
						_ = _nwElement0_34
						var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(4))
						_ = _nw166
						(_nw166).ArraySet1(_nwElement0_34, 0)
						(_nw166).ArraySet1(_dafny.IntOfInt64(330), 1)
						(_nw166).ArraySet1(_1062_v3, 2)
						(_nw166).ArraySet1(Companion_Default___.SafeDivisionInt(_1062_v3, _1062_v3), 3)
						_1106_v35 = _nw166
						var _1107_v36 _dafny.Map
						_ = _1107_v36
						_1107_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1106_v35, p0)
						var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1106_v35), 0))
						_ = _index195
						(_1106_v35).ArraySet1((_1107_v36).Cardinality(), (_index195).Int())
						var _1108_v37 _dafny.Map
						_ = _1108_v37
						_1108_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), (_1083_v16).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32()).(bool))
						var _1109_v38 _dafny.Map
						_ = _1109_v38
						_1109_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1108_v37, _1062_v3)
						var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1106_v35), 0))
						_ = _index196
						(_1106_v35).ArraySet1(((_1109_v38).Update(_1108_v37, _1062_v3)).Cardinality(), (_index196).Int())
						var _1110_v39 *C1
						_ = _1110_v39
						var _nw167 *C1 = New_C1_()
						_ = _nw167
						_nw167.Ctor__(p0, _1061_v2, _1106_v35)
						_1110_v39 = _nw167
						var _1111_v40 _dafny.Map
						_ = _1111_v40
						_1111_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F13()), _dafny.IntOfUint32((_1061_v2).Cardinality()))
						var _1112_v41 _dafny.Sequence
						_ = _1112_v41
						_1112_v41 = _dafny.SeqOf(_1062_v3, _dafny.IntOfInt64(431))
						var _1113_v42 _dafny.MultiSet
						_ = _1113_v42
						_1113_v42 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1110_v39.F18).Cardinality()))
						var _1114_v43 _dafny.Sequence
						_ = _1114_v43
						_1114_v43 = _dafny.SeqOf((_1112_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.IntOfUint32((_1112_v41).Cardinality()))).Uint32()).(_dafny.Int), (func() _dafny.Int {
							if (_1113_v42).Contains(_dafny.IntOfUint32((_1061_v2).Cardinality())) {
								return (_1113_v42).Multiplicity(_dafny.IntOfUint32((_1061_v2).Cardinality()))
							}
							return (_1106_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1106_v35), 0))).Int()).(_dafny.Int)
						})(), (_1106_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1106_v35), 0))).Int()).(_dafny.Int), _1062_v3)
						var _rhs165 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
							if !((_this).F16()) {
								return _1061_v2
							}
							return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("aiwrcowth"), (Companion_Default___.SafeIndex((_1111_v40).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aiwrcowth")).Cardinality()))).Uint32(), _1060_v1)
						})(), _1110_v39.F18)
						_ = _rhs165
						var _rhs166 bool = !(!(p1))
						_ = _rhs166
						var _rhs167 _dafny.Int = (_1114_v43).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1114_v43).Cardinality()))).Uint32()).(_dafny.Int)
						_ = _rhs167
						_1061_v2 = _rhs165
						r0 = _rhs166
						_1062_v3 = _rhs167
						r2 = (_1106_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1106_v35), 0))).Int()).(_dafny.Int)
						_1113_v42 = (_dafny.MultiSetFromSeq(_1112_v41)).Update(_1062_v3, Companion_Default___.Abs(((_1106_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1106_v35), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(31))))
					}
					var _1115_v44 D1
					_ = _1115_v44
					_1115_v44 = Companion_D1_.Create_DC4_(_1084_v17, _1062_v3, _1060_v1)
					var _1116_v45 _dafny.Sequence
					_ = _1116_v45
					_1116_v45 = _dafny.SeqOf(_1061_v2, _dafny.Companion_Sequence_.Update(_1061_v2, (Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1061_v2).Cardinality()))).Uint32(), _dafny.CodePoint('l')), _1061_v2)
					var _1117_v46 _dafny.Array
					_ = _1117_v46
					var _nwElement0_35 _dafny.Int = (_1062_v3).Minus((_1115_v44).Dtor_cf6())
					_ = _nwElement0_35
					var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(11))
					_ = _nw168
					(_nw168).ArraySet1(_nwElement0_35, 0)
					(_nw168).ArraySet1(_dafny.IntOfInt64(609), 1)
					(_nw168).ArraySet1(Companion_Default___.SafeModuloInt(_1062_v3, _dafny.IntOfUint32((_dafny.SeqOf((_1063_v4).Dtor_cf13())).Cardinality())), 2)
					(_nw168).ArraySet1(_1062_v3, 3)
					(_nw168).ArraySet1(_1062_v3, 4)
					(_nw168).ArraySet1((_dafny.Zero).Minus(_1062_v3), 5)
					(_nw168).ArraySet1(Companion_Default___.SafeDivisionInt(_1062_v3, _dafny.IntOfUint32((_1083_v16).Cardinality())), 6)
					(_nw168).ArraySet1((_1062_v3).Times(_1062_v3), 7)
					(_nw168).ArraySet1((_1062_v3).Times(_1062_v3), 8)
					(_nw168).ArraySet1((_dafny.Zero).Minus((_1062_v3).Times(_dafny.IntOfUint32((_1116_v45).Cardinality()))), 9)
					(_nw168).ArraySet1(_1062_v3, 10)
					_1117_v46 = _nw168
					var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
					_ = _index197
					(_1117_v46).ArraySet1(_dafny.IntOfUint32((_1061_v2).Cardinality()), (_index197).Int())
					var _1118_v47 D6
					_ = _1118_v47
					_1118_v47 = Companion_D6_.Create_DC17_(_1062_v3, _1062_v3, _1084_v17)
					var _1119_v48 _dafny.Set
					_ = _1119_v48
					_1119_v48 = _dafny.SetOf(_1118_v47)
					var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
					_ = _index198
					(_1117_v46).ArraySet1(((_1119_v48).Union(_1119_v48)).Cardinality(), (_index198).Int())
					var _1120_v49 _dafny.MultiSet
					_ = _1120_v49
					_1120_v49 = _dafny.MultiSetOf(_1060_v1)
					var _1121_v50 _dafny.Sequence
					_ = _1121_v50
					_1121_v50 = _dafny.SeqOf(Companion_Default___.Fm31(globalState))
					var _rhs168 _dafny.Int = _dafny.IntOfInt64(120)
					_ = _rhs168
					var _rhs169 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dcdkyipa")).Cardinality())
					_ = _rhs169
					var _rhs170 _dafny.MultiSet = _dafny.MultiSetFromSeq(Companion_Default___.Fm28((func() _dafny.Int {
						if _this.F13() {
							return _dafny.IntOfUint32((_1061_v2).Cardinality())
						}
						return (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)
					})(), Companion_Default___.Fm31(globalState), !_dafny.Companion_Sequence_.Equal(_1121_v50, _1121_v50), globalState))
					_ = _rhs170
					var _rhs171 _dafny.Array = _1117_v46
					_ = _rhs171
					_1062_v3 = _rhs168
					_1062_v3 = _rhs169
					_1120_v49 = _rhs170
					_1117_v46 = _rhs171
					var _1122_v51 _dafny.Sequence
					_ = _1122_v51
					_1122_v51 = _dafny.SeqOf(_1121_v50)
					var _1123_v52 D2
					_ = _1123_v52
					_1123_v52 = Companion_D2_.Create_DC6_((_1122_v51).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm31(globalState), _dafny.IntOfUint32((_1122_v51).Cardinality()))).Uint32()).(_dafny.Sequence))
					var _source19 D2 = _1123_v52
					_ = _source19
					if _source19.Is_DC7() {
						var _1124___mcc_h2 _dafny.Sequence = _source19.Get_().(D2_DC7).Cf12
						_ = _1124___mcc_h2
						var _1125___mcc_h3 _dafny.Int = _source19.Get_().(D2_DC7).Cf13
						_ = _1125___mcc_h3
						var _1126___mcc_h4 _dafny.Int = _source19.Get_().(D2_DC7).Cf14
						_ = _1126___mcc_h4
						var _1127___mcc_h5 bool = _source19.Get_().(D2_DC7).Cf15
						_ = _1127___mcc_h5
						var _1128_cf15 bool = _1127___mcc_h5
						_ = _1128_cf15
						var _1129_cf14 _dafny.Int = _1126___mcc_h4
						_ = _1129_cf14
						var _1130_cf13 _dafny.Int = _1125___mcc_h3
						_ = _1130_cf13
						var _1131_cf12 _dafny.Sequence = _1124___mcc_h2
						_ = _1131_cf12
						var _1132_v53 _dafny.Map
						_ = _1132_v53
						_1132_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ajufne"), _1061_v2), (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int))
						_1132_v53 = _1132_v53
						var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1084_v17), 0))
						_ = _index199
						(_1084_v17).ArraySet1(Companion_Default___.Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_1062_v3)), (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int), globalState), (_index199).Int())
						var _1133_v54 _dafny.Map
						_ = _1133_v54
						_1133_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1120_v49).Cardinality()), _this.F13())
						var _1134_v55 _dafny.Set
						_ = _1134_v55
						_1134_v55 = _dafny.SetOf(_1083_v16)
						var _1135_v56 *C2
						_ = _1135_v56
						var _nw169 *C2 = New_C2_()
						_ = _nw169
						_nw169.Ctor__(true)
						_1135_v56 = _nw169
						var _1136_v57 _dafny.Map
						_ = _1136_v57
						_1136_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1135_v56, (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int))
						var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _index200
						var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1084_v17), 0))
						_ = _index201
						var _rhs172 bool = (_1128_cf15) && (p1)
						_ = _rhs172
						var _rhs173 _dafny.Int = Companion_Default___.SafeModuloInt(_1130_cf13, (_1130_cf13).Minus((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)))
						_ = _rhs173
						var _rhs174 bool = !_dafny.Companion_Sequence_.Contains(_1121_v50, (_1062_v3).Minus(_dafny.IntOfInt64(635)))
						_ = _rhs174
						var _rhs175 bool = ((_1134_v55).Intersection(_1134_v55)).IsProperSubsetOf(_dafny.SetOf(_1083_v16, _1083_v16))
						_ = _rhs175
						var _rhs176 _dafny.Map = (_1133_v54).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1122_v51).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_1136_v57).Contains(_1135_v56) {
								return (_1136_v57).Get(_1135_v56).(_dafny.Int)
							}
							return (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)
						})(), _dafny.IntOfUint32((_1122_v51).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), p0)).Merge(_1133_v54))
						_ = _rhs176
						var _lhs101 _dafny.Array = _1117_v46
						_ = _lhs101
						var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _lhs102
						var _lhs103 _dafny.Array = _1084_v17
						_ = _lhs103
						var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1084_v17), 0))
						_ = _lhs104
						r1 = _rhs172
						(_lhs101).ArraySet1(_rhs173, (_lhs102).Int())
						(_lhs103).ArraySet1(_rhs174, (_lhs104).Int())
						r0 = _rhs175
						_1133_v54 = _rhs176
						var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _index202
						(_1117_v46).ArraySet1((Companion_Default___.Fm13(p1, (_this).F16(), _1060_v1, globalState)).Minus(_dafny.IntOfInt64(-352)), (_index202).Int())
						var _1137_v58 _dafny.Array
						_ = _1137_v58
						var _len0_28 _dafny.Int = _dafny.IntOfInt64(15)
						_ = _len0_28
						var _nw170 _dafny.Array
						_ = _nw170
						if _len0_28.Cmp(_dafny.Zero) == 0 {
							_nw170 = _dafny.NewArray(_len0_28)
						} else {
							var _init28 func(_dafny.Int) _dafny.Sequence = (func(_1138_v16 _dafny.Sequence, _1139_v17 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
								return func(_1140_i4 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_1138_v16, _dafny.SeqOf((_1139_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1139_v17), 0))).Int()).(bool)))
								}
							})(_1083_v16, _1084_v17)
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
						_1137_v58 = _nw170
						_1137_v58 = _1137_v58
					} else if _source19.Is_DC8() {
						var _1141___mcc_h6 bool = _source19.Get_().(D2_DC8).Cf16
						_ = _1141___mcc_h6
						var _1142___mcc_h7 bool = _source19.Get_().(D2_DC8).Cf17
						_ = _1142___mcc_h7
						var _1143___mcc_h8 bool = _source19.Get_().(D2_DC8).Cf18
						_ = _1143___mcc_h8
						var _1144_cf18 bool = _1143___mcc_h8
						_ = _1144_cf18
						var _1145_cf17 bool = _1142___mcc_h7
						_ = _1145_cf17
						var _1146_cf16 bool = _1141___mcc_h6
						_ = _1146_cf16
						var _1147_v59 _dafny.Map
						_ = _1147_v59
						_1147_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1144_cf18, (_1120_v49).Cardinality())
						var _1148_v60 _dafny.MultiSet
						_ = _1148_v60
						_1148_v60 = _dafny.MultiSetOf(_1062_v3, (func() _dafny.Int {
							if (_1147_v59).Contains(_1144_cf18) {
								return (_1147_v59).Get(_1144_cf18).(_dafny.Int)
							}
							return (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)
						})(), (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int))
						var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _index203
						(_1117_v46).ArraySet1((_1148_v60).Cardinality(), (_index203).Int())
						var _1149_v61 _dafny.MultiSet
						_ = _1149_v61
						_1149_v61 = _dafny.MultiSetOf(_1061_v2, _dafny.Companion_Sequence_.Concatenate(_1061_v2, _dafny.UnicodeSeqOfUtf8Bytes("hpqdcyt")), (_1116_v45).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-673)), _dafny.IntOfUint32((_1116_v45).Cardinality()))).Uint32()).(_dafny.Sequence))
						_1149_v61 = ((_dafny.MultiSetFromSeq(_1116_v45)).Intersection(Companion_Default___.Fm43(globalState))).Intersection(_1149_v61)
						var _1150_v62 _dafny.Map
						_ = _1150_v62
						_1150_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1062_v3, _dafny.IntOfInt64(767))
						var _1151_v63 D1
						_ = _1151_v63
						_1151_v63 = Companion_D1_.Create_DC2_(_1150_v62)
						var _1152_v64 _dafny.Sequence
						_ = _1152_v64
						_1152_v64 = _dafny.SeqOf(_1151_v63, _1151_v63)
						_1152_v64 = Companion_Default___.Fm44(_1062_v3, false, _dafny.IntOfInt64(704), !_dafny.Companion_Sequence_.Equal(_1061_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(330))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}((func(_1153_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1154_i5 _dafny.Int) _dafny.CodePoint {
								return _1153_v1
							}
						})(_1060_v1)))), globalState)
						var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _index204
						(_1117_v46).ArraySet1((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int), (_index204).Int())
					} else {
						var _1155___mcc_h9 _dafny.Sequence = _source19.Get_().(D2_DC6).Cf11
						_ = _1155___mcc_h9
						var _1156_cf11 _dafny.Sequence = _1155___mcc_h9
						_ = _1156_cf11
						var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _index205
						(_1117_v46).ArraySet1((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int), (_index205).Int())
						var _1157_v65 _dafny.Sequence
						_ = _1157_v65
						_1157_v65 = _dafny.SeqOf(_1083_v16, _1083_v16)
						var _1158_v66 D9
						_ = _1158_v66
						_1158_v66 = Companion_D9_.Create_DC23_(_dafny.MultiSetFromSeq((_1157_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-343), _dafny.IntOfUint32((_1157_v65).Cardinality()))).Uint32()).(_dafny.Sequence)))
						var _1159_v67 _dafny.MultiSet
						_ = _1159_v67
						_1159_v67 = _dafny.MultiSetOf((_this).F16())
						var _1160_v68 *C3
						_ = _1160_v68
						var _nw171 *C3 = New_C3_()
						_ = _nw171
						_nw171.Ctor__(!((_1159_v67).IsProperSubsetOf((_1158_v66).Dtor_cf35())), _this.F13())
						_1160_v68 = _nw171
						var _1161_v69 D1
						_ = _1161_v69
						_1161_v69 = Companion_D1_.Create_DC5_(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("caoyl"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("caoyl")).Cardinality()))).Uint32(), _1060_v1), _1084_v17, _1062_v3)
						var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))
						_ = _index206
						(_1117_v46).ArraySet1(((_1161_v69).Dtor_cf10()).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_1156_cf11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1061_v2).Cardinality()), _dafny.IntOfUint32((_1156_cf11).Cardinality()))).Uint32()).(_dafny.Int)))), (_index206).Int())
						var _1162_v70 D5
						_ = _1162_v70
						_1162_v70 = Companion_D5_.Create_DC14_(p0)
						var _1163_v71 D5
						_ = _1163_v71
						_1163_v71 = Companion_D5_.Create_DC15_(_1162_v70)
						var _1164_v72 _dafny.MultiSet
						_ = _1164_v72
						_1164_v72 = _dafny.MultiSetOf(_1062_v3, (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int))
						var _1165_v73 _dafny.MultiSet
						_ = _1165_v73
						_1165_v73 = _dafny.MultiSetOf(_1164_v72, _1164_v72)
						var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_1084_v17), 0))
						_ = _index207
						(_1084_v17).ArraySet1((_1165_v73).IsProperSubsetOf(_1165_v73), (_index207).Int())
						var _1166_v74 D5
						_ = _1166_v74
						_1166_v74 = Companion_D5_.Create_DC12_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _1062_v3))
						var _1167_v75 _dafny.Map
						_ = _1167_v75
						_1167_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1061_v2).Cardinality()), _1084_v17)
						var _1168_v76 _dafny.Map
						_ = _1168_v76
						_1168_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(476), (func() _dafny.Array {
							if (_1167_v75).Contains((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)) {
								return (_1167_v75).Get((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)).(_dafny.Array)
							}
							return _1084_v17
						})())
						var _1169_v77 _dafny.Map
						_ = _1169_v77
						_1169_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1084_v17, _1084_v17)
						var _1170_v78 _dafny.Array
						_ = _1170_v78
						var _nwElement0_36 _dafny.Array = _1084_v17
						_ = _nwElement0_36
						var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(21))
						_ = _nw172
						(_nw172).ArraySet1(_nwElement0_36, 0)
						(_nw172).ArraySet1(_1084_v17, 1)
						(_nw172).ArraySet1(_1084_v17, 2)
						(_nw172).ArraySet1(_1084_v17, 3)
						(_nw172).ArraySet1(_1084_v17, 4)
						(_nw172).ArraySet1(_1084_v17, 5)
						(_nw172).ArraySet1((_1161_v69).Dtor_cf9(), 6)
						(_nw172).ArraySet1((func() _dafny.Array {
							if _this.F13() {
								return _1084_v17
							}
							return _1084_v17
						})(), 7)
						(_nw172).ArraySet1((func() _dafny.Array {
							if (_1168_v76).Contains((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)) {
								return (_1168_v76).Get((_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int)).(_dafny.Array)
							}
							return _1084_v17
						})(), 8)
						(_nw172).ArraySet1(_1084_v17, 9)
						(_nw172).ArraySet1(_1084_v17, 10)
						(_nw172).ArraySet1(_1084_v17, 11)
						(_nw172).ArraySet1(_1084_v17, 12)
						(_nw172).ArraySet1(_1084_v17, 13)
						(_nw172).ArraySet1(_1084_v17, 14)
						(_nw172).ArraySet1(_1084_v17, 15)
						(_nw172).ArraySet1((func() _dafny.Array {
							if (_1169_v77).Contains(_1084_v17) {
								return (_1169_v77).Get(_1084_v17).(_dafny.Array)
							}
							return _1084_v17
						})(), 16)
						(_nw172).ArraySet1(_1084_v17, 17)
						(_nw172).ArraySet1(_1084_v17, 18)
						(_nw172).ArraySet1(_1084_v17, 19)
						(_nw172).ArraySet1(_1084_v17, 20)
						_1170_v78 = _nw172
						var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1170_v78), 0))
						_ = _index208
						(_1170_v78).ArraySet1(_1084_v17, (_index208).Int())
						var _1171_v79 _dafny.Array
						_ = _1171_v79
						var _nwElement0_37 _dafny.Sequence = Companion_Default___.Fm32(_1061_v2, (_this).F16(), _1062_v3, (_1117_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1117_v46), 0))).Int()).(_dafny.Int), globalState)
						_ = _nwElement0_37
						var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(2))
						_ = _nw173
						(_nw173).ArraySet1(_nwElement0_37, 0)
						(_nw173).ArraySet1(_1061_v2, 1)
						_1171_v79 = _nw173
						var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_1171_v79), 0))
						_ = _index209
						(_1171_v79).ArraySet1(_1061_v2, (_index209).Int())
						var _1172_v80 _dafny.Map
						_ = _1172_v80
						_1172_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1061_v2, _1061_v2), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1061_v2, _1061_v2)).Cardinality()))).Uint32(), _1060_v1), !(_dafny.Companion_Sequence_.IsProperPrefixOf(_1061_v2, _1061_v2)))
						var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_1084_v17), 0))
						_ = _index210
						var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1170_v78), 0))
						_ = _index211
						var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_1171_v79), 0))
						_ = _index212
						var _rhs177 D5 = _1163_v71
						_ = _rhs177
						var _rhs178 bool = (func() bool {
							if (_1172_v80).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(70))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg60 _dafny.Int) interface{} {
									return coer60(arg60)
								}
							}((func(_1173_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1174_i6 _dafny.Int) _dafny.CodePoint {
									return _1173_v1
								}
							})(_1060_v1)))) {
								return (_1172_v80).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(70))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg61 _dafny.Int) interface{} {
										return coer61(arg61)
									}
								}((func(_1175_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_1176_i6 _dafny.Int) _dafny.CodePoint {
										return _1175_v1
									}
								})(_1060_v1)))).(bool)
							}
							return (_this).F16()
						})()
						_ = _rhs178
						var _rhs179 D5 = _1166_v74
						_ = _rhs179
						var _rhs180 _dafny.Array = _1084_v17
						_ = _rhs180
						var _rhs181 _dafny.Sequence = _1061_v2
						_ = _rhs181
						var _lhs105 _dafny.Array = _1084_v17
						_ = _lhs105
						var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_1084_v17), 0))
						_ = _lhs106
						var _lhs107 _dafny.Array = _1170_v78
						_ = _lhs107
						var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1170_v78), 0))
						_ = _lhs108
						var _lhs109 _dafny.Array = _1171_v79
						_ = _lhs109
						var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_1171_v79), 0))
						_ = _lhs110
						_1163_v71 = _rhs177
						(_lhs105).ArraySet1(_rhs178, (_lhs106).Int())
						_1166_v74 = _rhs179
						(_lhs107).ArraySet1(_rhs180, (_lhs108).Int())
						(_lhs109).ArraySet1(_rhs181, (_lhs110).Int())
					}
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1177_v81 _dafny.Map
		_ = _1177_v81
		_1177_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
		var _1178_v82 _dafny.Sequence
		_ = _1178_v82
		_1178_v82 = _dafny.Companion_Sequence_.Update(_1083_v16, (Companion_Default___.SafeIndex((_1177_v81).Cardinality(), _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32(), _this.F13())
		var _1179_v83 _dafny.Array
		_ = _1179_v83
		var _nwElement0_38 _dafny.Sequence = _1178_v82
		_ = _nwElement0_38
		var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(15))
		_ = _nw174
		(_nw174).ArraySet1(_nwElement0_38, 0)
		(_nw174).ArraySet1(_1178_v82, 1)
		(_nw174).ArraySet1(_1178_v82, 2)
		(_nw174).ArraySet1(_1178_v82, 3)
		(_nw174).ArraySet1(_1178_v82, 4)
		(_nw174).ArraySet1(_1178_v82, 5)
		(_nw174).ArraySet1(_1178_v82, 6)
		(_nw174).ArraySet1(Companion_Default___.Fm45(p0, _dafny.IntOfInt64(67), globalState), 7)
		(_nw174).ArraySet1(_1178_v82, 8)
		(_nw174).ArraySet1(_1178_v82, 9)
		(_nw174).ArraySet1(_1178_v82, 10)
		(_nw174).ArraySet1(Companion_Default___.Fm45(_this.F13(), _1062_v3, globalState), 11)
		(_nw174).ArraySet1(_1178_v82, 12)
		(_nw174).ArraySet1(_1178_v82, 13)
		(_nw174).ArraySet1(_1178_v82, 14)
		_1179_v83 = _nw174
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_1179_v83), 0))
		_ = _index213
		(_1179_v83).ArraySet1(_1178_v82, (_index213).Int())
		var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_1179_v83), 0))
		_ = _index214
		(_1179_v83).ArraySet1(_1178_v82, (_index214).Int())
		var _1180_v84 _dafny.Set
		_ = _1180_v84
		_1180_v84 = _dafny.SetOf(_1062_v3, _1062_v3)
		var _1181_v85 _dafny.Set
		_ = _1181_v85
		_1181_v85 = _dafny.SetOf((_this).F16())
		var _1182_v86 _dafny.Map
		_ = _1182_v86
		_1182_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _1181_v85)
		var _1183_v87 _dafny.Sequence
		_ = _1183_v87
		_1183_v87 = _dafny.SeqOf(_1181_v85, _1181_v85)
		var _1184_v88 _dafny.MultiSet
		_ = _1184_v88
		_1184_v88 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1061_v2).Cardinality()))
		var _1185_v89 _dafny.Map
		_ = _1185_v89
		_1185_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1184_v88, p1)
		var _1186_v90 _dafny.Array
		_ = _1186_v90
		var _nwElement0_39 _dafny.Set = _dafny.SetOf(p1, (_1083_v16).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32()).(bool))
		_ = _nwElement0_39
		var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(23))
		_ = _nw175
		(_nw175).ArraySet1(_nwElement0_39, 0)
		(_nw175).ArraySet1((func() _dafny.Set {
			if (func() bool {
				if (_1177_v81).Contains(p0) {
					return (_1177_v81).Get(p0).(bool)
				}
				return _this.F13()
			})() {
				return Companion_Default___.Fm25(_1180_v84, _1177_v81, _1062_v3, _this.F13(), globalState)
			}
			return _dafny.SetOf(p1)
		})(), 1)
		(_nw175).ArraySet1((func() _dafny.Set {
			if (_1182_v86).Contains(!(_this.F13())) {
				return (_1182_v86).Get(!(_this.F13())).(_dafny.Set)
			}
			return _dafny.SetOf(p0)
		})(), 2)
		(_nw175).ArraySet1(_1181_v85, 3)
		(_nw175).ArraySet1((func() _dafny.Set {
			if true {
				return _dafny.SetOf((_this).F16(), p0)
			}
			return _1181_v85
		})(), 4)
		(_nw175).ArraySet1(_1181_v85, 5)
		(_nw175).ArraySet1(_1181_v85, 6)
		(_nw175).ArraySet1(_1181_v85, 7)
		(_nw175).ArraySet1((_1183_v87).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1183_v87).Cardinality()))).Uint32()).(_dafny.Set), 8)
		(_nw175).ArraySet1(_1181_v85, 9)
		(_nw175).ArraySet1(_1181_v85, 10)
		(_nw175).ArraySet1((func() _dafny.Set {
			if (func() bool {
				if (_1185_v89).Contains(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}(func(_1187_i7 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(915)
				})))) {
					return (_1185_v89).Get(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg63 _dafny.Int) interface{} {
							return coer63(arg63)
						}
					}(func(_1188_i7 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(915)
					})))).(bool)
				}
				return p0
			})() {
				return _1181_v85
			}
			return _1181_v85
		})(), 11)
		(_nw175).ArraySet1((_1181_v85).Intersection(_1181_v85), 12)
		(_nw175).ArraySet1(_1181_v85, 13)
		(_nw175).ArraySet1((func() _dafny.Set {
			if p1 {
				return _dafny.SetOf(p1, (_this).F16())
			}
			return _1181_v85
		})(), 14)
		(_nw175).ArraySet1(_1181_v85, 15)
		(_nw175).ArraySet1(_1181_v85, 16)
		(_nw175).ArraySet1(_1181_v85, 17)
		(_nw175).ArraySet1((_1181_v85).Difference(_1181_v85), 18)
		(_nw175).ArraySet1((func() _dafny.Set {
			if _this.F13() {
				return _1181_v85
			}
			return _dafny.SetOf((_this).F16())
		})(), 19)
		(_nw175).ArraySet1(_1181_v85, 20)
		(_nw175).ArraySet1(_dafny.SetOf(true), 21)
		(_nw175).ArraySet1(_1181_v85, 22)
		_1186_v90 = _nw175
		var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1186_v90), 0))
		_ = _index215
		(_1186_v90).ArraySet1((_1181_v85).Intersection(_1181_v85), (_index215).Int())
		var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1186_v90), 0))
		_ = _index216
		var _rhs182 _dafny.Array = _1084_v17
		_ = _rhs182
		var _rhs183 _dafny.Set = ((_1181_v85).Union(_dafny.SetOf((_1083_v16).Select((Companion_Default___.SafeIndex(_1062_v3, _dafny.IntOfUint32((_1083_v16).Cardinality()))).Uint32()).(bool), _this.F13(), p0, (_this).F16()))).Union(_1181_v85)
		_ = _rhs183
		var _lhs111 _dafny.Array = _1186_v90
		_ = _lhs111
		var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1186_v90), 0))
		_ = _lhs112
		_1084_v17 = _rhs182
		(_lhs111).ArraySet1(_rhs183, (_lhs112).Int())
		r0 = _dafny.Companion_Sequence_.IsPrefixOf(_1061_v2, _1061_v2)
		r1 = (_1181_v85).IsProperSubsetOf(((_1186_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1186_v90), 0))).Int()).(_dafny.Set)).Intersection(_1181_v85))
		r2 = ((_1177_v81).Cardinality()).Plus(_1062_v3)
		return r0, r1, r2
	}
}
func (_this *C4) F16() bool {
	{
		return _this._f16
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f13 bool
	_f14 _dafny.Array
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f13 = false
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C5{}
var _ T1 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F13() bool {
	return _this._f13
}
func (_this *C5) F13_set_(value bool) {
	_this._f13 = value
}
func (_this *C5) F14() _dafny.Array {
	return _this._f14
}
func (_this *C5) Ctor__(f13 bool, f14 _dafny.Array) {
	{
		(_this)._f13 = f13
		(_this)._f14 = f14
	}
}
func (_this *C5) Fm4(p0 bool, p1 D0, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Int {
			if _this.F13() {
				return _dafny.IntOfInt64(790)
			}
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F13())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-516), _dafny.IntOfInt64(198), (_dafny.SetOf(_dafny.CodePoint('m'), _dafny.CodePoint('q'))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("flisrw")).Cardinality())), (_dafny.SetOf(_this.F13())).Cardinality())).Cardinality()))).Cardinality())
		})()).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(757), _dafny.IntOfInt64(-211))).Cardinality()) < 0
	}
}
func (_this *C5) Fm5(globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC3_(Companion_Default___.SafeModuloInt((_dafny.SetOf(_this.F13())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(346))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}(func(_1189_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality())))
	}
}
func (_this *C5) Fm6(globalState *GlobalState) bool {
	{
		return !((Companion_D0_.Create_DC1_(_this.F13(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("box")).Cardinality()))).Dtor_cf1())
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(457), _dafny.IntOfInt64(-669))).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nuqc")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(114)))).Cardinality()))
	}
}
func (_this *C5) Fm8(p0 D1, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("rq"), _dafny.CodePoint('a'))
	}
}
func (_this *C5) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((func() _dafny.Int {
			if _this.F13() {
				return _dafny.IntOfInt64(-694)
			}
			return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F13(), !(_this.F13()), _this.F13()))).Cardinality()
		})()), (_dafny.IntOfInt64(-159)).Times((_dafny.MultiSetOf(!(_this.F13()), _this.F13())).Cardinality()))
	}
}
func (_this *C5) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _1190_v0 _dafny.Array
		_ = _1190_v0
		var _nwElement0_40 bool = !(!(_this.F13()))
		_ = _nwElement0_40
		var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.One)
		_ = _nw176
		(_nw176).ArraySet1(_nwElement0_40, 0)
		_1190_v0 = _nw176
		var _1191_v1 D1
		_ = _1191_v1
		_1191_v1 = Companion_D1_.Create_DC4_(_1190_v0, p0, _dafny.CodePoint('i'))
		var _source20 D1 = _1191_v1
		_ = _source20
		if _source20.Is_DC3() {
			var _1192___mcc_h0 _dafny.Int = _source20.Get_().(D1_DC3).Cf4
			_ = _1192___mcc_h0
			var _1193_cf4 _dafny.Int = _1192___mcc_h0
			_ = _1193_cf4
			var _1194_v2 _dafny.MultiSet
			_ = _1194_v2
			_1194_v2 = _dafny.MultiSetOf(_1193_cf4)
			_1193_cf4 = (_this).Fm7((_this).Fm9(p0, globalState), Companion_Default___.SafeDivisionInt(p0, ((_1194_v2).Update(p0, Companion_Default___.Abs(p0))).Cardinality()), _this.F13(), globalState)
			if !(true) {
				var _1195_v3 _dafny.Sequence
				_ = _1195_v3
				_1195_v3 = _dafny.UnicodeSeqOfUtf8Bytes("gljnndblh")
				var _1196_v4 _dafny.Int
				_ = _1196_v4
				var _out31 _dafny.Int
				_ = _out31
				_out31 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1195_v3, _this.F13()), globalState)
				_1196_v4 = _out31
				var _1197_v5 _dafny.Array
				_ = _1197_v5
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw177
				_1197_v5 = _nw177
				_1197_v5 = _1197_v5
				var _1198_v6 *C0
				_ = _1198_v6
				var _nw178 *C0 = New_C0_()
				_ = _nw178
				_nw178.Ctor__((_this).Fm7(_1193_cf4, (_dafny.Zero).Minus(_1196_v4), _this.F13(), globalState))
				_1198_v6 = _nw178
				var _1199_v7 bool
				_ = _1199_v7
				var _1200_v8 _dafny.Int
				_ = _1200_v8
				var _out32 bool
				_ = _out32
				var _out33 _dafny.Int
				_ = _out33
				_out32, _out33 = (_this).M3(_1193_cf4, (_1198_v6).F15(), (_1198_v6).F15(), globalState)
				_1199_v7 = _out32
				_1200_v8 = _out33
				_1195_v3 = _1195_v3
			} else {
				var _rhs184 bool = _this.F13()
				_ = _rhs184
				var _rhs185 _dafny.Int = _1193_cf4
				_ = _rhs185
				var _lhs113 *C5 = _this
				_ = _lhs113
				_lhs113.F13_set_(_rhs184)
				_1193_cf4 = _rhs185
				var _1201_v9 _dafny.Map
				_ = _1201_v9
				_1201_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), (_1194_v2).Update(p0, Companion_Default___.Abs(p0)))
				_1201_v9 = (_1201_v9).Update(_this.F13(), (_1194_v2).Difference(_1194_v2))
				_1193_cf4 = p0
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index217
				((_this).F14()).ArraySet1(p0, (_index217).Int())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index218
				((_this).F14()).ArraySet1(_1193_cf4, (_index218).Int())
				var _1202_v10 _dafny.Map
				_ = _1202_v10
				_1202_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int), _this.F13())
				_1193_cf4 = (_1202_v10).Cardinality()
			}
			var _1203_v11 _dafny.Sequence
			_ = _1203_v11
			_1203_v11 = _dafny.SeqOf((_this).F14())
			var _1204_v12 *C0
			_ = _1204_v12
			var _nw179 *C0 = New_C0_()
			_ = _nw179
			_nw179.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32((_1203_v11).Cardinality())))
			_1204_v12 = _nw179
			var _1205_v13 _dafny.Sequence
			_ = _1205_v13
			_1205_v13 = _dafny.UnicodeSeqOfUtf8Bytes("dqmf")
			var _1206_v14 _dafny.CodePoint
			_ = _1206_v14
			_1206_v14 = _dafny.CodePoint('w')
			_1205_v13 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1205_v13, _dafny.UnicodeSeqOfUtf8Bytes("abddeo")), (Companion_Default___.SafeIndex((p0).Times(_dafny.IntOfUint32((_1205_v13).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1205_v13, _dafny.UnicodeSeqOfUtf8Bytes("abddeo"))).Cardinality()))).Uint32(), _1206_v14)
		} else if _source20.Is_DC4() {
			var _1207___mcc_h1 _dafny.Array = _source20.Get_().(D1_DC4).Cf5
			_ = _1207___mcc_h1
			var _1208___mcc_h2 _dafny.Int = _source20.Get_().(D1_DC4).Cf6
			_ = _1208___mcc_h2
			var _1209___mcc_h3 _dafny.CodePoint = _source20.Get_().(D1_DC4).Cf7
			_ = _1209___mcc_h3
			var _1210_cf7 _dafny.CodePoint = _1209___mcc_h3
			_ = _1210_cf7
			var _1211_cf6 _dafny.Int = _1208___mcc_h2
			_ = _1211_cf6
			var _1212_cf5 _dafny.Array = _1207___mcc_h1
			_ = _1212_cf5
			var _1213_v15 _dafny.MultiSet
			_ = _1213_v15
			_1213_v15 = _dafny.MultiSetOf(_this.F13(), _this.F13())
			_1211_cf6 = (_1213_v15).Cardinality()
			(_this).F13_set_(_this.F13())
			var _1214_v16 _dafny.Sequence
			_ = _1214_v16
			_1214_v16 = _dafny.UnicodeSeqOfUtf8Bytes("buj")
			_1214_v16 = _1214_v16
			var _1215_v17 _dafny.Map
			_ = _1215_v17
			_1215_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1211_cf6, (_this).Fm9((_this).Fm9(p0, globalState), globalState))
			var _1216_v18 _dafny.Map
			_ = _1216_v18
			_1216_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(685), _1215_v17)
			var _1217_v19 _dafny.Map
			_ = _1217_v19
			_1217_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())
			_1216_v18 = (_1216_v18).Update((_dafny.Zero).Minus((_this).Fm9((_1217_v19).Cardinality(), globalState)), (_1215_v17).Update(p0, (_1213_v15).Cardinality()))
		} else if _source20.Is_DC5() {
			var _1218___mcc_h4 _dafny.Sequence = _source20.Get_().(D1_DC5).Cf8
			_ = _1218___mcc_h4
			var _1219___mcc_h5 _dafny.Array = _source20.Get_().(D1_DC5).Cf9
			_ = _1219___mcc_h5
			var _1220___mcc_h6 _dafny.Int = _source20.Get_().(D1_DC5).Cf10
			_ = _1220___mcc_h6
			var _1221_cf10 _dafny.Int = _1220___mcc_h6
			_ = _1221_cf10
			var _1222_cf9 _dafny.Array = _1219___mcc_h5
			_ = _1222_cf9
			var _1223_cf8 _dafny.Sequence = _1218___mcc_h4
			_ = _1223_cf8
			var _1224_v20 _dafny.Array
			_ = _1224_v20
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_29
			var _nw180 _dafny.Array
			_ = _nw180
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw180 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.MultiSet = (func(_1225_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_1226_i0 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetFromSeq((Companion_D2_.Create_DC6_(_dafny.SeqOf(_1225_p0))).Dtor_cf11())
					}
				})(p0)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw180 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw180).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw180).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1224_v20 = _nw180
			_1224_v20 = _1224_v20
			var _1227_v21 D0
			_ = _1227_v21
			_1227_v21 = Companion_D0_.Create_DC0_(_this.F13())
			_1227_v21 = _1227_v21
			var _1228_v22 _dafny.Map
			_ = _1228_v22
			_1228_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1191_v1)
			var _1229_v23 _dafny.Map
			_ = _1229_v23
			_1229_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _dafny.CodePoint('x'))
			_1228_v22 = (_1228_v22).Update((_1229_v23).Cardinality(), _1191_v1)
			var _1230_v24 _dafny.Map
			_ = _1230_v24
			_1230_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1223_cf8, Companion_Default___.Fm11(_this.F13(), p0, _this.F13(), globalState))
			var _1231_v25 _dafny.CodePoint
			_ = _1231_v25
			_1231_v25 = _dafny.CodePoint('b')
			_1230_v24 = (_1230_v24).Update(_1223_cf8, _1231_v25)
		} else {
			var _1232___mcc_h7 _dafny.Map = _source20.Get_().(D1_DC2).Cf3
			_ = _1232___mcc_h7
			var _1233_cf3 _dafny.Map = _1232___mcc_h7
			_ = _1233_cf3
			var _1234_v26 _dafny.Int
			_ = _1234_v26
			_1234_v26 = _dafny.IntOfInt64(283)
			_1234_v26 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1234_v26))
			_1234_v26 = p0
			_1234_v26 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v26, (func() _dafny.Int {
				if !(_this.F13()) {
					return _1234_v26
				}
				return _1234_v26
			})())).Cardinality()
			(_this).F13_set_(_this.F13())
		}
		var _1235_v27 _dafny.Int
		_ = _1235_v27
		_1235_v27 = _dafny.IntOfInt64(-407)
		var _1236_v28 _dafny.Sequence
		_ = _1236_v28
		_1236_v28 = _dafny.UnicodeSeqOfUtf8Bytes("sffwtcir")
		_1235_v27 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1236_v28, _1236_v28)).Cardinality())
		var _1237_v29 _dafny.Array
		_ = _1237_v29
		var _nw181 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
		_ = _nw181
		_1237_v29 = _nw181
		var _1238_v30 _dafny.Set
		_ = _1238_v30
		_1238_v30 = _dafny.SetOf(_this.F13())
		var _1239_v31 _dafny.Set
		_ = _1239_v31
		_1239_v31 = _dafny.SetOf(_1238_v30)
		var _1240_v32 *C0
		_ = _1240_v32
		var _nw182 *C0 = New_C0_()
		_ = _nw182
		_nw182.Ctor__((_1239_v31).Cardinality())
		_1240_v32 = _nw182
		var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1237_v29), 0))
		_ = _index219
		(_1237_v29).ArraySet1(_1240_v32, (_index219).Int())
		var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1237_v29), 0))
		_ = _index220
		(_1237_v29).ArraySet1(_1240_v32, (_index220).Int())
		(_this).F13_set_(true)
		var _1241_v33 *C0
		_ = _1241_v33
		var _nw183 *C0 = New_C0_()
		_ = _nw183
		_nw183.Ctor__(p0)
		_1241_v33 = _nw183
		var _1242_v34 _dafny.Set
		_ = _1242_v34
		_1242_v34 = _dafny.SetOf(_dafny.IntOfInt64(549))
		var _1243_v35 _dafny.Map
		_ = _1243_v35
		_1243_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-689)), _1242_v34)
		var _1244_v36 _dafny.MultiSet
		_ = _1244_v36
		_1244_v36 = _dafny.MultiSetOf(false, _this.F13())
		var _1245_v37 _dafny.Map
		_ = _1245_v37
		_1245_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1244_v36, (_1241_v33).F15())
		var _1246_v38 _dafny.Sequence
		_ = _1246_v38
		_1246_v38 = _dafny.SeqOf(p0)
		var _1247_v39 _dafny.Map
		_ = _1247_v39
		_1247_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1246_v38).Select((Companion_Default___.SafeIndex((_this).Fm9(_1235_v27, globalState), _dafny.IntOfUint32((_1246_v38).Cardinality()))).Uint32()).(_dafny.Int), (_this).F14())
		var _rhs186 bool = ((((func() _dafny.Set {
			if (_1243_v35).Contains((_1240_v32).F15()) {
				return (_1243_v35).Get((_1240_v32).F15()).(_dafny.Set)
			}
			return _1242_v34
		})()).Difference(_1242_v34)).Cardinality()).Cmp((_1245_v37).Cardinality()) <= 0
		_ = _rhs186
		var _rhs187 *C0 = _1241_v33
		_ = _rhs187
		var _rhs188 _dafny.Int = (p0).Times((_1247_v39).Cardinality())
		_ = _rhs188
		var _lhs114 *C5 = _this
		_ = _lhs114
		_lhs114.F13_set_(_rhs186)
		_1240_v32 = _rhs187
		_1235_v27 = _rhs188
		r0 = (_this).F14()
		return r0
	}
}
func (_this *C5) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Sequence) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _1248_i0 _dafny.Int
		_ = _1248_i0
		_1248_i0 = _dafny.Zero
		{
			for _this.F13() {
				{
					if (_1248_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1248_i0 = (_1248_i0).Plus(_dafny.One)
					var _1249_v0 _dafny.Int
					_ = _1249_v0
					_1249_v0 = _dafny.IntOfInt64(21)
					_1249_v0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(p2, p2), (func() _dafny.Int {
						if _this.F13() {
							return p0
						}
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(545))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}(func(_1250_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('v')
						}))).Cardinality())
					})())
					var _1251_v1 _dafny.Array
					_ = _1251_v1
					var _len0_30 _dafny.Int = _dafny.IntOfInt64(20)
					_ = _len0_30
					var _nw184 _dafny.Array
					_ = _nw184
					if _len0_30.Cmp(_dafny.Zero) == 0 {
						_nw184 = _dafny.NewArray(_len0_30)
					} else {
						var _init30 func(_dafny.Int) _dafny.Sequence = (func(_1252_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_1253_i2 _dafny.Int) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Concatenate(_1252_p1, _1252_p1)
							}
						})(p1)
						_ = _init30
						var _element0_30 = _init30(_dafny.Zero)
						_ = _element0_30
						_nw184 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
						(_nw184).ArraySet1(_element0_30, 0)
						var _nativeLen0_30 = (_len0_30).Int()
						_ = _nativeLen0_30
						for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
							(_nw184).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
						}
					}
					_1251_v1 = _nw184
					var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1251_v1), 0))
					_ = _index221
					(_1251_v1).ArraySet1(p1, (_index221).Int())
					var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1251_v1), 0))
					_ = _index222
					(_1251_v1).ArraySet1(p1, (_index222).Int())
					var _1254_v2 _dafny.MultiSet
					_ = _1254_v2
					_1254_v2 = _dafny.MultiSetOf(p2)
					var _1255_v3 _dafny.Set
					_ = _1255_v3
					_1255_v3 = _dafny.SetOf(_1254_v2)
					var _1256_v4 _dafny.CodePoint
					_ = _1256_v4
					_1256_v4 = _dafny.CodePoint('t')
					var _1257_v5 _dafny.Map
					_ = _1257_v5
					_1257_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1255_v3, _dafny.Companion_Sequence_.Concatenate(p1, _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_1249_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1256_v4)))
					_1257_v5 = (_1257_v5).Update(_1255_v3, _dafny.Companion_Sequence_.Update((_1251_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1251_v1), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_1251_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1251_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1256_v4))
					var _1258_v6 _dafny.Array
					_ = _1258_v6
					var _nw185 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
					_ = _nw185
					_1258_v6 = _nw185
					var _1259_v7 T0
					_ = _1259_v7
					var _nw186 *C3 = New_C3_()
					_ = _nw186
					_nw186.Ctor__(_this.F13(), _this.F13())
					_1259_v7 = _nw186
					var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1258_v6), 0))
					_ = _index223
					(_1258_v6).ArraySet1(_1259_v7, (_index223).Int())
					var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1258_v6), 0))
					_ = _index224
					var _nw187 *C3 = New_C3_()
					_ = _nw187
					_nw187.Ctor__(_this.F13(), !(_this.F13()))
					(_1258_v6).ArraySet1(_nw187, (_index224).Int())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1260_v8 *C3
		_ = _1260_v8
		var _nw188 *C3 = New_C3_()
		_ = _nw188
		_nw188.Ctor__((func() bool {
			if _this.F13() {
				return false
			}
			return _this.F13()
		})(), _this.F13())
		_1260_v8 = _nw188
		if _this.F13() {
			var _1261_v9 _dafny.Sequence
			_ = _1261_v9
			_1261_v9 = _dafny.SeqOf(!(_this.F13()))
			var _1262_v10 _dafny.Map
			_ = _1262_v10
			_1262_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_1261_v9).Cardinality()))
			var _1263_v11 D0
			_ = _1263_v11
			_1263_v11 = Companion_D0_.Create_DC1_(_this.F13(), (func() _dafny.Int {
				if (_1262_v10).Contains(p2) {
					return (_1262_v10).Get(p2).(_dafny.Int)
				}
				return p0
			})())
			var _source21 D0 = _1263_v11
			_ = _source21
			if _source21.Is_DC1() {
				var _1264___mcc_h0 bool = _source21.Get_().(D0_DC1).Cf1
				_ = _1264___mcc_h0
				var _1265___mcc_h1 _dafny.Int = _source21.Get_().(D0_DC1).Cf2
				_ = _1265___mcc_h1
				var _1266_cf2 _dafny.Int = _1265___mcc_h1
				_ = _1266_cf2
				var _1267_cf1 bool = _1264___mcc_h0
				_ = _1267_cf1
				(_this).F13_set_(Companion_Default___.Fm0((_1262_v10).Merge(_1262_v10), (_dafny.IntOfUint32((p1).Cardinality())).Times(p0), globalState))
				_1266_cf2 = _dafny.IntOfInt64(207)
				_1266_cf2 = p0
				var _1268_v12 _dafny.Set
				_ = _1268_v12
				_1268_v12 = _dafny.SetOf(p2, p0, p0)
				var _1269_v13 _dafny.Map
				_ = _1269_v13
				_1269_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1261_v9).Cardinality()), (_1260_v8).F19())
				var _1270_v15 _dafny.Array
				_ = _1270_v15
				var _nwElement0_41 _dafny.Set = _1268_v12
				_ = _nwElement0_41
				var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(5))
				_ = _nw189
				(_nw189).ArraySet1(_nwElement0_41, 0)
				(_nw189).ArraySet1(_1268_v12, 1)
				(_nw189).ArraySet1((_1268_v12).Intersection(_1268_v12), 2)
				(_nw189).ArraySet1(_1268_v12, 3)
				(_nw189).ArraySet1(func() _dafny.Set {
					var _coll42 = _dafny.NewBuilder()
					_ = _coll42
					for _iter44 := _dafny.Iterate((_1269_v13).Keys().Elements()); ; {
						_compr_42, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _1271_v14 _dafny.Int
						_1271_v14 = interface{}(_compr_42).(_dafny.Int)
						if (_1269_v13).Contains(_1271_v14) {
							_coll42.Add((_1271_v14).Times(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, false))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(442))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg66 _dafny.Int) interface{} {
									return coer66(arg66)
								}
							}(func(_1272_i3 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('o')
							}))).Cardinality()))).Cardinality())))
						}
					}
					return _coll42.ToSet()
				}(), 4)
				_1270_v15 = _nw189
				var _1273_v16 _dafny.Map
				_ = _1273_v16
				_1273_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1260_v8)
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1270_v15), 0))
				_ = _index225
				(_1270_v15).ArraySet1(_dafny.SetOf((_1273_v16).Cardinality()), (_index225).Int())
				var _1274_v17 _dafny.Sequence
				_ = _1274_v17
				_1274_v17 = _dafny.UnicodeSeqOfUtf8Bytes("ucptmuvvm")
				var _1275_v18 _dafny.Sequence
				_ = _1275_v18
				_1275_v18 = _dafny.SeqOf(p2)
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1270_v15), 0))
				_ = _index226
				var _rhs189 _dafny.Set = ((_1268_v12).Union(_1268_v12)).Difference((_dafny.SetOf(_1266_cf2, _1266_cf2, _dafny.IntOfUint32((_1275_v18).Cardinality()))).Intersection(_1268_v12))
				_ = _rhs189
				var _rhs190 bool = _this.F13()
				_ = _rhs190
				var _rhs191 bool = (_1260_v8).F19()
				_ = _rhs191
				var _rhs192 _dafny.Int = (_dafny.SetOf((p0).Plus(_1266_cf2), (_dafny.Zero).Minus((_dafny.Zero).Minus(p2)), p0)).Cardinality()
				_ = _rhs192
				var _rhs193 _dafny.Sequence = p1
				_ = _rhs193
				var _lhs115 _dafny.Array = _1270_v15
				_ = _lhs115
				var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1270_v15), 0))
				_ = _lhs116
				(_lhs115).ArraySet1(_rhs189, (_lhs116).Int())
				_1267_cf1 = _rhs190
				_1267_cf1 = _rhs191
				_1266_cf2 = _rhs192
				_1274_v17 = _rhs193
			} else {
				var _1276___mcc_h2 bool = _source21.Get_().(D0_DC0).Cf0
				_ = _1276___mcc_h2
				var _1277_cf0 bool = _1276___mcc_h2
				_ = _1277_cf0
				var _1278_v19 _dafny.CodePoint
				_ = _1278_v19
				_1278_v19 = _dafny.CodePoint('m')
				var _1279_v20 _dafny.Map
				_ = _1279_v20
				_1279_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1260_v8).F19(), (_1260_v8).F19())
				var _1280_v21 _dafny.MultiSet
				_ = _1280_v21
				_1280_v21 = _dafny.MultiSetOf((func() bool {
					if (_1279_v20).Contains(_1277_cf0) {
						return (_1279_v20).Get(_1277_cf0).(bool)
					}
					return (_1260_v8).F19()
				})(), _this.F13())
				var _1281_v22 D2
				_ = _1281_v22
				_1281_v22 = Companion_D2_.Create_DC7_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.CodePoint('b'), _1278_v19, _1278_v19), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('b'), _1278_v19, _1278_v19)).Cardinality()))).Uint32(), _1278_v19), (_1280_v21).Cardinality(), p2, _1277_cf0)
				var _1282_v23 _dafny.Map
				_ = _1282_v23
				_1282_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1281_v22).Dtor_cf15()), p2)
				_1282_v23 = (_1282_v23).Update(!((_1260_v8).F19()), p2)
				var _1283_v24 _dafny.Map
				_ = _1283_v24
				_1283_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1262_v10)
				var _1284_v25 *C4
				_ = _1284_v25
				var _nw190 *C4 = New_C4_()
				_ = _nw190
				_nw190.Ctor__((_this).Fm8(Companion_D1_.Create_DC2_((func() _dafny.Map {
					if (_1283_v24).Contains(p2) {
						return (_1283_v24).Get(p2).(_dafny.Map)
					}
					return Companion_Default___.Fm18((_1260_v8).F19(), (_1260_v8).F19(), (_1282_v23).Cardinality(), (_this).Fm7(p0, p2, _1277_cf0, globalState), globalState)
				})()), p2, globalState), !((_1260_v8).F19()))
				_1284_v25 = _nw190
				var _1285_v26 D7
				_ = _1285_v26
				_1285_v26 = Companion_D7_.Create_DC20_(_dafny.IntOfInt64(606), _1278_v19)
				var _1286_v27 D7
				_ = _1286_v27
				_1286_v27 = Companion_D7_.Create_DC21_(_1285_v26)
				var _1287_v28 D7
				_ = _1287_v28
				_1287_v28 = Companion_D7_.Create_DC21_(_1285_v26)
				var _1288_v29 D7
				_ = _1288_v29
				_1288_v29 = Companion_D7_.Create_DC21_(_1286_v27)
				_1288_v29 = _1288_v29
				_1279_v20 = (_1279_v20).Merge((_1279_v20).Merge(_1279_v20))
			}
			var _1289_v30 _dafny.CodePoint
			_ = _1289_v30
			_1289_v30 = _dafny.CodePoint('h')
			var _1290_v31 _dafny.Map
			_ = _1290_v31
			_1290_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D4_.Create_DC11_((_1260_v8).F19())).Dtor_cf21(), _this.F13())
			var _1291_v32 D2
			_ = _1291_v32
			_1291_v32 = Companion_D2_.Create_DC8_(_this.F13(), false, _this.F13())
			_1289_v30 = Companion_Default___.Fm11((func() bool {
				if (_1290_v31).Contains((_1260_v8).F19()) {
					return (_1290_v31).Get((_1260_v8).F19()).(bool)
				}
				return (_1291_v32).Dtor_cf17()
			})(), p2, Companion_Default___.Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-513)), p2, globalState), globalState)
			var _1292_v33 _dafny.Array
			_ = _1292_v33
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw191
			_1292_v33 = _nw191
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))
			_ = _index227
			(_1292_v33).ArraySet1((p0).Cmp(p2) <= 0, (_index227).Int())
			var _1293_v34 *C4
			_ = _1293_v34
			var _nw192 *C4 = New_C4_()
			_ = _nw192
			_nw192.Ctor__(_this.F13(), (_1260_v8).F19())
			_1293_v34 = _nw192
			var _1294_v35 _dafny.Set
			_ = _1294_v35
			_1294_v35 = _dafny.SetOf((_1260_v8).F19())
			var _1295_v36 _dafny.Map
			_ = _1295_v36
			_1295_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1293_v34, _1294_v35)
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))
			_ = _index228
			var _rhs194 bool = (_1260_v8).F19()
			_ = _rhs194
			var _rhs195 bool = ((_this).Fm9(((func() _dafny.Set {
				if (_1295_v36).Contains(_1293_v34) {
					return (_1295_v36).Get(_1293_v34).(_dafny.Set)
				}
				return _dafny.SetOf(_this.F13())
			})()).Cardinality(), globalState)).Cmp(p0) >= 0
			_ = _rhs195
			var _lhs117 *C5 = _this
			_ = _lhs117
			var _lhs118 _dafny.Array = _1292_v33
			_ = _lhs118
			var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))
			_ = _lhs119
			_lhs117.F13_set_(_rhs194)
			(_lhs118).ArraySet1(_rhs195, (_lhs119).Int())
			var _1296_v37 _dafny.Sequence
			_ = _1296_v37
			_1296_v37 = _1261_v9
			var _source22 _dafny.Sequence = _1296_v37
			_ = _source22
			var _1297___mcc_h3 _dafny.Sequence = _source22
			_ = _1297___mcc_h3
			var _1298_cf34 _dafny.Sequence = _1297___mcc_h3
			_ = _1298_cf34
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))
			_ = _index229
			(_1292_v33).ArraySet1((_1260_v8).F19(), (_index229).Int())
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index230
			((_this).F14()).ArraySet1(p2, (_index230).Int())
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index231
			((_this).F14()).ArraySet1(p2, (_index231).Int())
			var _1299_v38 D5
			_ = _1299_v38
			_1299_v38 = Companion_D5_.Create_DC14_((_1292_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))).Int()).(bool))
			var _1300_v39 _dafny.Set
			_ = _1300_v39
			_1300_v39 = _dafny.SetOf(Companion_D5_.Create_DC14_((_1293_v34).F16()), _1299_v38, _1299_v38, _1299_v38)
			var _1301_v40 _dafny.Array
			_ = _1301_v40
			var _nwElement0_42 _dafny.Array = _1292_v33
			_ = _nwElement0_42
			var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(9))
			_ = _nw193
			(_nw193).ArraySet1(_nwElement0_42, 0)
			(_nw193).ArraySet1(_1292_v33, 1)
			(_nw193).ArraySet1(_1292_v33, 2)
			(_nw193).ArraySet1(_1292_v33, 3)
			(_nw193).ArraySet1(_1292_v33, 4)
			(_nw193).ArraySet1(_1292_v33, 5)
			(_nw193).ArraySet1(_1292_v33, 6)
			(_nw193).ArraySet1(_1292_v33, 7)
			(_nw193).ArraySet1(_1292_v33, 8)
			_1301_v40 = _nw193
			var _1302_v41 _dafny.Map
			_ = _1302_v41
			_1302_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1300_v39, _1301_v40)
			_1302_v41 = (_1302_v41).Update(_1300_v39, _1301_v40)
			var _1303_v42 _dafny.Array
			_ = _1303_v42
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_31
			var _nw194 _dafny.Array
			_ = _nw194
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw194 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) D9 = (func(_1304_v34 *C4) func(_dafny.Int) D9 {
					return func(_1305_i4 _dafny.Int) D9 {
						return Companion_D9_.Create_DC23_(_dafny.MultiSetOf((_1304_v34).F16()))
					}
				})(_1293_v34)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw194 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw194).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw194).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1303_v42 = _nw194
			var _1306_v43 D1
			_ = _1306_v43
			_1306_v43 = Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm7(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int), (_dafny.MultiSetOf((_dafny.Zero).Minus(p0))).Cardinality(), (_1293_v34).F16(), globalState)))
			var _1307_v44 _dafny.MultiSet
			_ = _1307_v44
			_1307_v44 = _dafny.MultiSetOf((_this).Fm8(_1306_v43, p0, globalState), (_1293_v34).F16())
			var _1308_v45 D9
			_ = _1308_v45
			_1308_v45 = Companion_D9_.Create_DC23_(_1307_v44)
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1303_v42), 0))
			_ = _index232
			(_1303_v42).ArraySet1(_1308_v45, (_index232).Int())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))
			_ = _index233
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index234
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1303_v42), 0))
			_ = _index235
			var _rhs196 bool = !(false)
			_ = _rhs196
			var _rhs197 _dafny.Int = (func() _dafny.Int {
				if (_1307_v44).Contains(!(true)) {
					return (_1307_v44).Multiplicity(!(true))
				}
				return p2
			})()
			_ = _rhs197
			var _rhs198 D9 = _1308_v45
			_ = _rhs198
			var _rhs199 _dafny.CodePoint = _1289_v30
			_ = _rhs199
			var _lhs120 _dafny.Array = _1292_v33
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v33), 0))
			_ = _lhs121
			var _lhs122 _dafny.Array = (_this).F14()
			_ = _lhs122
			var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs123
			var _lhs124 _dafny.Array = _1303_v42
			_ = _lhs124
			var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1303_v42), 0))
			_ = _lhs125
			(_lhs120).ArraySet1(_rhs196, (_lhs121).Int())
			(_lhs122).ArraySet1(_rhs197, (_lhs123).Int())
			(_lhs124).ArraySet1(_rhs198, (_lhs125).Int())
			_1289_v30 = _rhs199
			_1261_v9 = _1261_v9
		} else {
			var _1309_v46 _dafny.Sequence
			_ = _1309_v46
			_1309_v46 = _dafny.SeqOf(p0, _dafny.IntOfInt64(627), p0, _dafny.IntOfUint32((p1).Cardinality()))
			var _1310_v47 *C3
			_ = _1310_v47
			var _nw195 *C3 = New_C3_()
			_ = _nw195
			_nw195.Ctor__((func() bool {
				if (_1260_v8).F19() {
					return true
				}
				return true
			})(), !_dafny.Companion_Sequence_.Equal(_1309_v46, _1309_v46))
			_1310_v47 = _nw195
			var _1311_v48 _dafny.Int
			_ = _1311_v48
			_1311_v48 = _dafny.IntOfInt64(58)
			var _1312_v49 _dafny.CodePoint
			_ = _1312_v49
			_1312_v49 = _dafny.CodePoint('r')
			_1311_v48 = Companion_Default___.Fm13((_1260_v8).F19(), _this.F13(), _1312_v49, globalState)
			_1311_v48 = p0
			var _1313_v50 *C3
			_ = _1313_v50
			var _nw196 *C3 = New_C3_()
			_ = _nw196
			_nw196.Ctor__(!(_dafny.SetOf(false)).Contains(_this.F13()), (p0).Cmp((_dafny.Zero).Minus(p0)) != 0)
			_1313_v50 = _nw196
			var _1314_v51 _dafny.Array
			_ = _1314_v51
			var _nw197 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
			_ = _nw197
			_1314_v51 = _nw197
			var _1315_v52 *C0
			_ = _1315_v52
			var _nw198 *C0 = New_C0_()
			_ = _nw198
			_nw198.Ctor__(p2)
			_1315_v52 = _nw198
			var _1316_v53 _dafny.Map
			_ = _1316_v53
			_1316_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1315_v52, (_1310_v47).F19())
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_1314_v51), 0))
			_ = _index236
			(_1314_v51).ArraySet1(_1316_v53, (_index236).Int())
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_1314_v51), 0))
			_ = _index237
			(_1314_v51).ArraySet1(((_1316_v53).Update(_1315_v52, (_1260_v8).F19())).Merge(_1316_v53), (_index237).Int())
		}
		var _1317_v54 _dafny.Sequence
		_ = _1317_v54
		_1317_v54 = _dafny.UnicodeSeqOfUtf8Bytes("jvtqh")
		_1317_v54 = p1
		var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index238
		((_this).F14()).ArraySet1(p2, (_index238).Int())
		var _1318_v55 _dafny.MultiSet
		_ = _1318_v55
		_1318_v55 = _dafny.MultiSetOf(p0, p0)
		var _1319_v56 _dafny.Sequence
		_ = _1319_v56
		_1319_v56 = _dafny.SeqOf(_dafny.SeqOf(p0, (_1318_v55).Cardinality(), p0))
		var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index239
		((_this).F14()).ArraySet1(_dafny.IntOfUint32((_1319_v56).Cardinality()), (_index239).Int())
		var _1320_v57 _dafny.Map
		_ = _1320_v57
		_1320_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _1321_v58 _dafny.Map
		_ = _1321_v58
		_1321_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(_1320_v57, (_dafny.MultiSetOf(_this.F13())).Cardinality(), globalState))
		var _1322_v59 _dafny.Map
		_ = _1322_v59
		_1322_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1321_v58).Merge(_1321_v58), _1317_v54)
		_1322_v59 = (_1322_v59).Update(_1321_v58, _dafny.UnicodeSeqOfUtf8Bytes("hbsuyxpfb"))
		var _1323_v60 _dafny.Array
		_ = _1323_v60
		var _len0_32 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_32
		var _nw199 _dafny.Array
		_ = _nw199
		if _len0_32.Cmp(_dafny.Zero) == 0 {
			_nw199 = _dafny.NewArray(_len0_32)
		} else {
			var _init32 func(_dafny.Int) bool = func(_1324_i5 _dafny.Int) bool {
				return _this.F13()
			}
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
		_1323_v60 = _nw199
		var _1325_v61 _dafny.Sequence
		_ = _1325_v61
		_1325_v61 = _dafny.SeqOf(_1323_v60)
		var _1326_v62 _dafny.Map
		_ = _1326_v62
		_1326_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(73), (_1325_v61).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1325_v61).Cardinality()))).Uint32()).(_dafny.Array))
		var _1327_v63 _dafny.Sequence
		_ = _1327_v63
		_1327_v63 = _dafny.SeqOf(p2)
		var _1328_v64 D1
		_ = _1328_v64
		_1328_v64 = Companion_D1_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("yvhsai"), _1323_v60, p2)
		var _1329_v65 _dafny.Sequence
		_ = _1329_v65
		_1329_v65 = _dafny.SeqOf((_1326_v62).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1327_v63).Cardinality()), (Companion_D1_.Create_DC5_((_1328_v64).Dtor_cf8(), _1323_v60, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int))).Dtor_cf9())))
		r0 = (_1329_v65).Select((Companion_Default___.SafeIndex(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1329_v65).Cardinality()))).Uint32()).(_dafny.Map)
		var _1330_v66 _dafny.Sequence
		_ = _1330_v66
		_1330_v66 = _dafny.SeqOf((_1260_v8).F19())
		r1 = (p0).Cmp((_dafny.IntOfUint32((_1330_v66).Cardinality())).Minus((_this).Fm9(_dafny.IntOfUint32((_1327_v63).Cardinality()), globalState))) != 0
		r2 = (func() bool {
			if (_1260_v8).F19() {
				return (_1260_v8).F19()
			}
			return _this.F13()
		})()
		r3 = _1330_v66
		return r0, r1, r2, r3
	}
}
func (_this *C5) M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1331_v0 _dafny.CodePoint
		_ = _1331_v0
		_1331_v0 = _dafny.CodePoint('m')
		var _1332_v1 D7
		_ = _1332_v1
		_1332_v1 = Companion_D7_.Create_DC20_(p1, _1331_v0)
		r1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p1, (_1332_v1).Dtor_cf31()))
		if !(true) {
			r0 = _this.F13()
			var _1333_v2 _dafny.MultiSet
			_ = _1333_v2
			_1333_v2 = _dafny.MultiSetOf(_dafny.IntOfInt64(508))
			r1 = (_1333_v2).Cardinality()
			var _1334_v3 _dafny.Array
			_ = _1334_v3
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw200
			_1334_v3 = _nw200
			var _1335_v4 _dafny.Sequence
			_ = _1335_v4
			_1335_v4 = _dafny.SeqOf(_this.F13(), _this.F13())
			var _1336_v5 _dafny.Set
			_ = _1336_v5
			_1336_v5 = _dafny.SetOf(_dafny.IntOfUint32((_1335_v4).Cardinality()), _dafny.IntOfUint32((_1335_v4).Cardinality()))
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_1334_v3), 0))
			_ = _index240
			(_1334_v3).ArraySet1((_1336_v5).IsDisjointFrom(_dafny.SetOf((_this).Fm9(_dafny.IntOfInt64(165), globalState), p1)), (_index240).Int())
			var _1337_v6 _dafny.Array
			_ = _1337_v6
			var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw201
			_1337_v6 = _nw201
			var _1338_v7 _dafny.Sequence
			_ = _1338_v7
			_1338_v7 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_1337_v6), 0))
			_ = _index241
			(_1337_v6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1338_v7, _1338_v7), (_index241).Int())
			var _1339_v8 D1
			_ = _1339_v8
			_1339_v8 = Companion_D1_.Create_DC5_(_1338_v7, _1334_v3, p0)
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_1334_v3), 0))
			_ = _index242
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_1337_v6), 0))
			_ = _index243
			var _rhs200 bool = _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("ywfkh"), _1331_v0)
			_ = _rhs200
			var _rhs201 _dafny.Sequence = (_1339_v8).Dtor_cf8()
			_ = _rhs201
			var _lhs126 _dafny.Array = _1334_v3
			_ = _lhs126
			var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_1334_v3), 0))
			_ = _lhs127
			var _lhs128 _dafny.Array = _1337_v6
			_ = _lhs128
			var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_1337_v6), 0))
			_ = _lhs129
			(_lhs126).ArraySet1(_rhs200, (_lhs127).Int())
			(_lhs128).ArraySet1(_rhs201, (_lhs129).Int())
			_1338_v7 = _1338_v7
			var _1340_v9 _dafny.Array
			_ = _1340_v9
			var _nwElement0_43 _dafny.CodePoint = _dafny.CodePoint('c')
			_ = _nwElement0_43
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(3))
			_ = _nw202
			(_nw202).ArraySet1CodePoint(_nwElement0_43, 0)
			(_nw202).ArraySet1CodePoint(_1331_v0, 1)
			(_nw202).ArraySet1CodePoint(_1331_v0, 2)
			_1340_v9 = _nw202
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1340_v9), 0))
			_ = _index244
			(_1340_v9).ArraySet1CodePoint(_1331_v0, (_index244).Int())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1340_v9), 0))
			_ = _index245
			(_1340_v9).ArraySet1CodePoint(_dafny.CodePoint('y'), (_index245).Int())
		} else {
			r1 = (_dafny.IntOfInt64(-939)).Times(_dafny.IntOfUint32((_dafny.SeqOf(_this.F13(), _this.F13(), true)).Cardinality()))
			var _1341_v10 _dafny.Sequence
			_ = _1341_v10
			_1341_v10 = _dafny.UnicodeSeqOfUtf8Bytes("garuvy")
			var _1342_v11 _dafny.Map
			_ = _1342_v11
			_1342_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1341_v10, !(_this.F13()) || ((_this).Fm6(globalState)))
			_1342_v11 = (_1342_v11).Update(_dafny.UnicodeSeqOfUtf8Bytes("fg"), !(_this.F13()))
			var _1343_v12 _dafny.Map
			_ = _1343_v12
			_1343_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
			var _1344_v13 _dafny.MultiSet
			_ = _1344_v13
			_1344_v13 = _dafny.MultiSetOf(_this.F13())
			var _1345_v14 _dafny.Map
			_ = _1345_v14
			_1345_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_this.F13())).Cardinality()), _1344_v13)
			_1343_v12 = (_1343_v12).Update(_dafny.IntOfInt64(-5), !(_1345_v14).Contains(p0))
			var _1346_v15 D0
			_ = _1346_v15
			_1346_v15 = Companion_D0_.Create_DC0_(_this.F13())
			_1346_v15 = Companion_D0_.Create_DC0_(_this.F13())
			r1 = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(898))
		}
		var _1347_v16 _dafny.Array
		_ = _1347_v16
		var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw203
		_1347_v16 = _nw203
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1347_v16), 0))); ; {
			_guard_loop_2, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _1348_i0 _dafny.Int
			_1348_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1348_i0).Sign() != -1) && ((_1348_i0).Cmp(_dafny.ArrayLen((_1347_v16), 0)) < 0)) {
				(_1347_v16).ArraySet1((_1348_i0).Plus(p2), (_1348_i0).Int())
			}
		}
		var _1349_v17 _dafny.Sequence
		_ = _1349_v17
		_1349_v17 = _dafny.UnicodeSeqOfUtf8Bytes("isf")
		var _hi11 _dafny.Int = _dafny.IntOfUint32((_1349_v17).Cardinality())
		_ = _hi11
		for _1350_i1 := p0; _1350_i1.Cmp(_hi11) < 0; _1350_i1 = _1350_i1.Plus(_dafny.One) {
			var _1351_v18 _dafny.Array
			_ = _1351_v18
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_33
			var _nw204 _dafny.Array
			_ = _nw204
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw204 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) bool = func(_1352_i2 _dafny.Int) bool {
					return false
				}
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
			_1351_v18 = _nw204
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))
			_ = _index246
			(_1351_v18).ArraySet1(_this.F13(), (_index246).Int())
			var _1353_v19 _dafny.Sequence
			_ = _1353_v19
			_1353_v19 = _dafny.SeqOf((p0).Cmp(p1) != 0)
			var _1354_v20 _dafny.Map
			_ = _1354_v20
			_1354_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))
			_ = _index247
			(_1351_v18).ArraySet1((_1353_v19).Select((Companion_Default___.SafeIndex((_1354_v20).Cardinality(), _dafny.IntOfUint32((_1353_v19).Cardinality()))).Uint32()).(bool), (_index247).Int())
			var _1355_v21 _dafny.Sequence
			_ = _1355_v21
			_1355_v21 = _dafny.SeqOf(p2)
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index248
			((_this).F14()).ArraySet1(((_1355_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1349_v17).Cardinality()), _dafny.IntOfUint32((_1355_v21).Cardinality()))).Uint32()).(_dafny.Int)).Times(p1), (_index248).Int())
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index249
			((_this).F14()).ArraySet1((_dafny.Zero).Minus(p2), (_index249).Int())
			(_this).F13_set_((_1351_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))).Int()).(bool))
			var _1356_v23 *C4
			_ = _1356_v23
			var _nw205 *C4 = New_C4_()
			_ = _nw205
			_nw205.Ctor__(true, _this.F13())
			_1356_v23 = _nw205
			var _1357_v24 _dafny.MultiSet
			_ = _1357_v24
			_1357_v24 = _dafny.MultiSetOf((_1356_v23).F16())
			var _1358_v25 _dafny.Map
			_ = _1358_v25
			_1358_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1356_v23, _1357_v24)
			var _1359_v26 _dafny.Map
			_ = _1359_v26
			_1359_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), p2)).Cardinality(), (_dafny.Zero).Minus((_1358_v25).Cardinality()))
			var _1360_v27 _dafny.Set
			_ = _1360_v27
			_1360_v27 = _dafny.SetOf((func() _dafny.Map {
				if (_1351_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))).Int()).(bool) {
					return func() _dafny.Map {
						var _coll43 = _dafny.NewMapBuilder()
						_ = _coll43
						for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(983), _dafny.IntOfInt64(402))); ; {
							_compr_43, _ok46 := _iter46()
							if !_ok46 {
								break
							}
							var _1361_v22 _dafny.Int
							_1361_v22 = interface{}(_compr_43).(_dafny.Int)
							if ((_dafny.IntOfInt64(983)).Cmp(_1361_v22) <= 0) && ((_1361_v22).Cmp(_dafny.IntOfInt64(402)) < 0) {
								_coll43.Add((_1361_v22).Minus(p0), p0)
							}
						}
						return _coll43.ToMap()
					}()
				}
				return (_1359_v26).Update(p0, _dafny.IntOfInt64(361))
			})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1349_v17).Cardinality()), _dafny.IntOfInt64(726)), (_1359_v26).Update(p0, (_dafny.Zero).Minus((_dafny.SetOf(Companion_Default___.Fm13((_1351_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))).Int()).(bool), (_1351_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))).Int()).(bool), _1331_v0, globalState), p0, p2, (Companion_Default___.Fm24(_dafny.IntOfUint32((_1349_v17).Cardinality()), (_1351_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))).Int()).(bool), globalState)).Cardinality())).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2))
			var _1362_v28 D0
			_ = _1362_v28
			_1362_v28 = Companion_D0_.Create_DC1_((_1356_v23).F16(), _dafny.IntOfInt64(356))
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))
			_ = _index250
			var _rhs202 _dafny.Int = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(_dafny.Int)
			_ = _rhs202
			var _rhs203 _dafny.Int = p0
			_ = _rhs203
			var _rhs204 bool = (_1356_v23).Fm4(((_dafny.Zero).Minus(p1)).Cmp(_dafny.IntOfInt64(306)) < 0, _1362_v28, _1350_i1, _1359_v26, globalState)
			_ = _rhs204
			var _rhs205 _dafny.Set = Companion_Default___.Fm46(p0, globalState)
			_ = _rhs205
			var _lhs130 _dafny.Array = _1351_v18
			_ = _lhs130
			var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1351_v18), 0))
			_ = _lhs131
			r1 = _rhs202
			r1 = _rhs203
			(_lhs130).ArraySet1(_rhs204, (_lhs131).Int())
			_1360_v27 = _rhs205
		}
		var _1363_i3 _dafny.Int
		_ = _1363_i3
		_1363_i3 = _dafny.Zero
		{
			for _this.F13() {
				{
					if (_1363_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1363_i3 = (_1363_i3).Plus(_dafny.One)
					r1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p2, (func() _dafny.Int {
						if true {
							return _dafny.IntOfInt64(-158)
						}
						return p0
					})()))
					var _1364_v29 _dafny.Map
					_ = _1364_v29
					_1364_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (true) == (true))
					var _1365_v30 _dafny.Map
					_ = _1365_v30
					_1365_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())
					var _1366_v31 _dafny.Map
					_ = _1366_v31
					_1366_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1365_v30).Cardinality()), p2)
					_1364_v29 = (_1364_v29).Update((func() _dafny.Int {
						if Companion_Default___.Fm0(_1366_v31, p0, globalState) {
							return p1
						}
						return p0
					})(), _this.F13())
					var _1367_v32 D2
					_ = _1367_v32
					_1367_v32 = Companion_D2_.Create_DC7_(_1349_v17, p1, p0, _this.F13())
					var _1368_v33 _dafny.Sequence
					_ = _1368_v33
					_1368_v33 = _dafny.SeqOf(_this.F13(), _this.F13(), _this.F13(), _this.F13(), (_this).Fm4(_this.F13(), Companion_D0_.Create_DC1_(_this.F13(), _dafny.IntOfInt64(-46)), _dafny.IntOfInt64(675), _1366_v31, globalState))
					var _1369_v34 _dafny.Sequence
					_ = _1369_v34
					_1369_v34 = _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('k')))
					var _pat_let_tv24 = p1
					_ = _pat_let_tv24
					var _1370_v35 _dafny.Array
					_ = _1370_v35
					var _nwElement0_44 D2 = _1367_v32
					_ = _nwElement0_44
					var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(19))
					_ = _nw206
					(_nw206).ArraySet1(_nwElement0_44, 0)
					(_nw206).ArraySet1(_1367_v32, 1)
					(_nw206).ArraySet1(_1367_v32, 2)
					(_nw206).ArraySet1(_1367_v32, 3)
					(_nw206).ArraySet1(_1367_v32, 4)
					(_nw206).ArraySet1(func(_pat_let28_0 D2) D2 {
						return func(_1371_dt__update__tmp_h0 D2) D2 {
							return func(_pat_let29_0 _dafny.Int) D2 {
								return func(_1372_dt__update_hcf13_h0 _dafny.Int) D2 {
									return Companion_D2_.Create_DC7_((_1371_dt__update__tmp_h0).Dtor_cf12(), _1372_dt__update_hcf13_h0, (_1371_dt__update__tmp_h0).Dtor_cf14(), (_1371_dt__update__tmp_h0).Dtor_cf15())
								}(_pat_let29_0)
							}(_pat_let_tv24)
						}(_pat_let28_0)
					}(_1367_v32), 5)
					(_nw206).ArraySet1(Companion_D2_.Create_DC7_(_dafny.Companion_Sequence_.Update(_1349_v17, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1349_v17).Cardinality()))).Uint32(), _1331_v0), p1, (_dafny.Zero).Minus(p0), _this.F13()), 6)
					(_nw206).ArraySet1(_1367_v32, 7)
					(_nw206).ArraySet1(_1367_v32, 8)
					(_nw206).ArraySet1(_1367_v32, 9)
					(_nw206).ArraySet1(Companion_D2_.Create_DC7_(_dafny.SeqOf(_1331_v0), _dafny.IntOfUint32((_1368_v33).Cardinality()), p0, _this.F13()), 10)
					(_nw206).ArraySet1(_1367_v32, 11)
					(_nw206).ArraySet1(Companion_Default___.Fm23(p1, _this.F13(), globalState), 12)
					(_nw206).ArraySet1(_1367_v32, 13)
					(_nw206).ArraySet1(_1367_v32, 14)
					(_nw206).ArraySet1(Companion_Default___.Fm23(p1, _this.F13(), globalState), 15)
					(_nw206).ArraySet1(Companion_D2_.Create_DC7_((_1369_v34).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1369_v34).Cardinality()))).Uint32()).(_dafny.Sequence), p0, p0, _this.F13()), 16)
					(_nw206).ArraySet1(Companion_D2_.Create_DC7_(_dafny.SeqOf(_1331_v0, _1331_v0), p2, p1, _this.F13()), 17)
					(_nw206).ArraySet1(_1367_v32, 18)
					_1370_v35 = _nw206
					var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1370_v35), 0))
					_ = _index251
					(_1370_v35).ArraySet1(_1367_v32, (_index251).Int())
					var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_1370_v35), 0))
					_ = _index252
					(_1370_v35).ArraySet1(_1367_v32, (_index252).Int())
					(_this).F13_set_(_this.F13())
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		if _this.F13() {
			var _rhs206 _dafny.Int = _dafny.IntOfInt64(-866)
			_ = _rhs206
			var _rhs207 bool = _this.F13()
			_ = _rhs207
			var _rhs208 _dafny.Int = (p2).Minus((p2).Minus(p2))
			_ = _rhs208
			var _rhs209 _dafny.Int = _dafny.IntOfInt64(107)
			_ = _rhs209
			r1 = _rhs206
			r0 = _rhs207
			r1 = _rhs208
			r1 = _rhs209
			r1 = p2
			var _1373_v36 D9
			_ = _1373_v36
			_1373_v36 = Companion_D9_.Create_DC24_(_this.F13(), _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("nxkq"), _1331_v0), _1331_v0)
			var _source23 D9 = _1373_v36
			_ = _source23
			if _source23.Is_DC24() {
				var _1374___mcc_h0 bool = _source23.Get_().(D9_DC24).Cf36
				_ = _1374___mcc_h0
				var _1375___mcc_h1 bool = _source23.Get_().(D9_DC24).Cf37
				_ = _1375___mcc_h1
				var _1376___mcc_h2 _dafny.CodePoint = _source23.Get_().(D9_DC24).Cf38
				_ = _1376___mcc_h2
				var _1377_cf38 _dafny.CodePoint = _1376___mcc_h2
				_ = _1377_cf38
				var _1378_cf37 bool = _1375___mcc_h1
				_ = _1378_cf37
				var _1379_cf36 bool = _1374___mcc_h0
				_ = _1379_cf36
				(_this).F13_set_(_1379_cf36)
				var _1380_v37 D7
				_ = _1380_v37
				_1380_v37 = Companion_D7_.Create_DC20_(p1, _1331_v0)
				var _1381_v38 D7
				_ = _1381_v38
				_1381_v38 = Companion_D7_.Create_DC21_(_1380_v37)
				var _1382_v39 D7
				_ = _1382_v39
				_1382_v39 = Companion_D7_.Create_DC21_(_1380_v37)
				var _1383_v40 D7
				_ = _1383_v40
				_1383_v40 = Companion_D7_.Create_DC21_(_1382_v39)
				var _1384_v41 _dafny.Sequence
				_ = _1384_v41
				_1384_v41 = _dafny.SeqOf(_1383_v40, _1383_v40)
				var _1385_v42 *C2
				_ = _1385_v42
				var _nw207 *C2 = New_C2_()
				_ = _nw207
				_nw207.Ctor__(!_dafny.Companion_Sequence_.Equal(_1384_v41, _1384_v41))
				_1385_v42 = _nw207
				_1385_v42 = _1385_v42
				var _1386_v43 _dafny.Sequence
				_ = _1386_v43
				_1386_v43 = _dafny.SeqOf(p1, p1)
				var _1387_v44 _dafny.MultiSet
				_ = _1387_v44
				_1387_v44 = _dafny.MultiSetOf(_1379_cf36)
				r0 = (p2).Cmp((_1386_v43).Select((Companion_Default___.SafeIndex((_1387_v44).Cardinality(), _dafny.IntOfUint32((_1386_v43).Cardinality()))).Uint32()).(_dafny.Int)) > 0
				var _1388_v45 _dafny.Array
				_ = _1388_v45
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw208
				_1388_v45 = _nw208
				var _1389_v46 D4
				_ = _1389_v46
				_1389_v46 = Companion_D4_.Create_DC10_(_1347_v16)
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1388_v45), 0))
				_ = _index253
				(_1388_v45).ArraySet1((_1389_v46).Dtor_cf20(), (_index253).Int())
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1388_v45), 0))
				_ = _index254
				(_1388_v45).ArraySet1(_1347_v16, (_index254).Int())
			} else if _source23.Is_DC25() {
				var _1390___mcc_h3 bool = _source23.Get_().(D9_DC25).Cf39
				_ = _1390___mcc_h3
				var _1391___mcc_h4 _dafny.Int = _source23.Get_().(D9_DC25).Cf40
				_ = _1391___mcc_h4
				var _1392___mcc_h5 bool = _source23.Get_().(D9_DC25).Cf41
				_ = _1392___mcc_h5
				var _1393_cf41 bool = _1392___mcc_h5
				_ = _1393_cf41
				var _1394_cf40 _dafny.Int = _1391___mcc_h4
				_ = _1394_cf40
				var _1395_cf39 bool = _1390___mcc_h3
				_ = _1395_cf39
				var _1396_v47 _dafny.Map
				_ = _1396_v47
				_1396_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm7((_dafny.Zero).Minus(p2), p0, _1395_cf39, globalState), p0)
				var _1397_v48 *C2
				_ = _1397_v48
				var _nw209 *C2 = New_C2_()
				_ = _nw209
				_nw209.Ctor__(Companion_Default___.Fm0(_1396_v47, p0, globalState))
				_1397_v48 = _nw209
				var _1398_v49 _dafny.Map
				_ = _1398_v49
				_1398_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1394_cf40, _1397_v48)
				var _rhs210 bool = (_1393_cf41) || (_1395_cf39)
				_ = _rhs210
				var _rhs211 bool = ((_1398_v49).Cardinality()).Cmp(p1) >= 0
				_ = _rhs211
				r0 = _rhs210
				_1395_cf39 = _rhs211
				var _1399_v50 _dafny.Map
				_ = _1399_v50
				_1399_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), (_dafny.SetOf(_1395_cf39, _1393_cf41)).Cardinality())
				var _1400_v51 _dafny.Sequence
				_ = _1400_v51
				_1400_v51 = _dafny.SeqOf((_dafny.Zero).Minus((_1399_v50).Cardinality()), _dafny.IntOfInt64(690), (_dafny.Zero).Minus(_1394_cf40))
				var _1401_v52 _dafny.Map
				_ = _1401_v52
				_1401_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1395_cf39, _1395_cf39)
				var _1402_v53 _dafny.Map
				_ = _1402_v53
				_1402_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1400_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.IntOfUint32((_1400_v51).Cardinality()))).Uint32()).(_dafny.Int), _1401_v52)
				var _1403_v54 _dafny.Map
				_ = _1403_v54
				_1403_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1331_v0, _1394_cf40)
				r1 = (((_1402_v53).Update((_1403_v54).Cardinality(), _1401_v52)).Cardinality()).Plus(p2)
				r1 = Companion_Default___.SafeDivisionInt(_1394_cf40, _1394_cf40)
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1347_v16), 0))
				_ = _index255
				(_1347_v16).ArraySet1((func() _dafny.Int {
					if _1395_cf39 {
						return p0
					}
					return _dafny.IntOfInt64(664)
				})(), (_index255).Int())
				var _1404_v55 D0
				_ = _1404_v55
				_1404_v55 = Companion_D0_.Create_DC1_(_1395_cf39, _dafny.IntOfInt64(209))
				var _1405_v56 _dafny.Set
				_ = _1405_v56
				_1405_v56 = _dafny.SetOf(_1331_v0, _1331_v0)
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1347_v16), 0))
				_ = _index256
				var _rhs212 _dafny.Int = Companion_Default___.SafeDivisionInt(((Companion_Default___.Fm47(_1404_v55, p0, globalState)).Union(func() _dafny.Set {
					var _coll44 = _dafny.NewBuilder()
					_ = _coll44
					for _iter47 := _dafny.Iterate((_1405_v56).Elements()); ; {
						_compr_44, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1406_v57 _dafny.CodePoint
						_1406_v57 = interface{}(_compr_44).(_dafny.CodePoint)
						if (_1405_v56).Contains(_1406_v57) {
							_coll44.Add(_1406_v57)
						}
					}
					return _coll44.ToSet()
				}())).Cardinality(), p2)
				_ = _rhs212
				var _rhs213 _dafny.Int = _dafny.IntOfUint32((_1349_v17).Cardinality())
				_ = _rhs213
				var _rhs214 _dafny.Map = _1399_v50
				_ = _rhs214
				var _lhs132 _dafny.Array = _1347_v16
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1347_v16), 0))
				_ = _lhs133
				r1 = _rhs212
				(_lhs132).ArraySet1(_rhs213, (_lhs133).Int())
				_1399_v50 = _rhs214
			} else {
				var _1407___mcc_h6 _dafny.MultiSet = _source23.Get_().(D9_DC23).Cf35
				_ = _1407___mcc_h6
				var _1408_cf35 _dafny.MultiSet = _1407___mcc_h6
				_ = _1408_cf35
				var _1409_v58 _dafny.Array
				_ = _1409_v58
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw210
				_1409_v58 = _nw210
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _index257
				(_1409_v58).ArraySet1((func() bool {
					if _this.F13() {
						return _this.F13()
					}
					return _this.F13()
				})(), (_index257).Int())
				var _1410_v59 *C1
				_ = _1410_v59
				var _nw211 *C1 = New_C1_()
				_ = _nw211
				_nw211.Ctor__(_this.F13(), _1349_v17, (_this).F14())
				_1410_v59 = _nw211
				var _1411_v60 _dafny.Sequence
				_ = _1411_v60
				_1411_v60 = _dafny.SeqOf(_this.F13())
				var _1412_v61 _dafny.Map
				_ = _1412_v61
				_1412_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1410_v59, (_dafny.MultiSetFromSeq(_1411_v60)).Union(_1408_cf35))
				var _1413_v62 _dafny.MultiSet
				_ = _1413_v62
				_1413_v62 = _dafny.MultiSetOf(_1410_v59.F18, _1410_v59.F18)
				var _1414_v63 _dafny.Map
				_ = _1414_v63
				_1414_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F13())
				var _1415_v64 _dafny.Sequence
				_ = _1415_v64
				_1415_v64 = _dafny.SeqOf(_1349_v17, _1410_v59.F18, _1349_v17)
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _index258
				var _rhs215 bool = (((_1413_v62).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-486))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_1416_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1417_i4 _dafny.Int) _dafny.CodePoint {
						return _1416_v0
					}
				})(_1331_v0))), Companion_Default___.Abs((_1414_v63).Cardinality()))).Update((_1415_v64).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1415_v64).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Abs((func() _dafny.Int {
					if (_1408_cf35).Contains(_this.F13()) {
						return (_1408_cf35).Multiplicity(_this.F13())
					}
					return p0
				})()))).IsDisjointFrom(_1413_v62)
				_ = _rhs215
				var _rhs216 bool = _this.F13()
				_ = _rhs216
				var _rhs217 bool = (func() bool {
					if (_1410_v59).F17() {
						return (func() bool {
							if _this.F13() {
								return _this.F13()
							}
							return false
						})()
					}
					return !(_this.F13())
				})()
				_ = _rhs217
				var _rhs218 _dafny.Map = (func() _dafny.Map {
					if !((_1410_v59).F17()) || ((_1410_v59).F17()) {
						return _1412_v61
					}
					return (_1412_v61).Merge(_1412_v61)
				})()
				_ = _rhs218
				var _lhs134 *C5 = _this
				_ = _lhs134
				var _lhs135 *C5 = _this
				_ = _lhs135
				var _lhs136 _dafny.Array = _1409_v58
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _lhs137
				_lhs134.F13_set_(_rhs215)
				_lhs135.F13_set_(_rhs216)
				(_lhs136).ArraySet1(_rhs217, (_lhs137).Int())
				_1412_v61 = _rhs218
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1347_v16), 0))
				_ = _index259
				(_1347_v16).ArraySet1(p2, (_index259).Int())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _index260
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1347_v16), 0))
				_ = _index261
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _index262
				var _rhs219 bool = (_1409_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))).Int()).(bool)
				_ = _rhs219
				var _rhs220 _dafny.Int = p0
				_ = _rhs220
				var _rhs221 _dafny.Int = ((_dafny.Zero).Minus(p2)).Minus(p0)
				_ = _rhs221
				var _rhs222 bool = (_1410_v59).F17()
				_ = _rhs222
				var _rhs223 _dafny.MultiSet = (_1408_cf35).Difference(_1408_cf35)
				_ = _rhs223
				var _lhs138 _dafny.Array = _1409_v58
				_ = _lhs138
				var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _lhs139
				var _lhs140 _dafny.Array = _1347_v16
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1347_v16), 0))
				_ = _lhs141
				var _lhs142 _dafny.Array = _1409_v58
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1409_v58), 0))
				_ = _lhs143
				(_lhs138).ArraySet1(_rhs219, (_lhs139).Int())
				(_lhs140).ArraySet1(_rhs220, (_lhs141).Int())
				r1 = _rhs221
				(_lhs142).ArraySet1(_rhs222, (_lhs143).Int())
				_1408_cf35 = _rhs223
				var _1418_v65 _dafny.Set
				_ = _1418_v65
				_1418_v65 = _dafny.SetOf(_1349_v17)
				var _1419_v66 _dafny.Array
				_ = _1419_v66
				var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw212
				_1419_v66 = _nw212
				var _1420_v67 _dafny.Map
				_ = _1420_v67
				_1420_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1410_v59.F18, (_1410_v59).F17())
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1347_v16), 0))
				_ = _index263
				var _rhs224 bool = (Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((_1411_v60).Cardinality()))).Cmp(p0) >= 0
				_ = _rhs224
				var _rhs225 _dafny.Set = func() _dafny.Set {
					var _coll45 = _dafny.NewBuilder()
					_ = _coll45
					for _iter48 := _dafny.Iterate((_1420_v67).Keys().Elements()); ; {
						_compr_45, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1421_v68 _dafny.Sequence
						_1421_v68 = interface{}(_compr_45).(_dafny.Sequence)
						if (_1420_v67).Contains(_1421_v68) {
							_coll45.Add(_1421_v68)
						}
					}
					return _coll45.ToSet()
				}()
				_ = _rhs225
				var _rhs226 _dafny.CodePoint = _1331_v0
				_ = _rhs226
				var _rhs227 _dafny.Array = (_this).F14()
				_ = _rhs227
				var _rhs228 _dafny.Int = (_1410_v59).Fm7(_dafny.IntOfInt64(517), ((_1347_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1347_v16), 0))).Int()).(_dafny.Int)).Times(p2), (_1410_v59).F17(), globalState)
				_ = _rhs228
				var _lhs144 _dafny.Array = _1347_v16
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_1347_v16), 0))
				_ = _lhs145
				r0 = _rhs224
				_1418_v65 = _rhs225
				_1331_v0 = _rhs226
				_1419_v66 = _rhs227
				(_lhs144).ArraySet1(_rhs228, (_lhs145).Int())
				r0 = !(true)
			}
			var _1422_v69 *C0
			_ = _1422_v69
			var _nw213 *C0 = New_C0_()
			_ = _nw213
			_nw213.Ctor__(p0)
			_1422_v69 = _nw213
			_1349_v17 = _1349_v17
		} else {
			var _1423_v70 _dafny.Set
			_ = _1423_v70
			_1423_v70 = _dafny.SetOf(_this.F13())
			var _1424_v71 _dafny.Map
			_ = _1424_v71
			_1424_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), (_1423_v70).Cardinality())
			var _1425_v72 _dafny.Map
			_ = _1425_v72
			_1425_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1424_v71, !(_this.F13()))
			var _1426_v73 _dafny.Map
			_ = _1426_v73
			_1426_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-410), (_1425_v72).Cardinality())
			var _1427_v74 _dafny.Sequence
			_ = _1427_v74
			_1427_v74 = _dafny.SeqOf(_dafny.IntOfInt64(-185))
			var _1428_v75 _dafny.Set
			_ = _1428_v75
			_1428_v75 = _dafny.SetOf(p0, _dafny.IntOfUint32((_1427_v74).Cardinality()))
			var _1429_v76 _dafny.Map
			_ = _1429_v76
			_1429_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Int {
				if (_1426_v73).Contains(p0) {
					return (_1426_v73).Get(p0).(_dafny.Int)
				}
				return (_1428_v75).Cardinality()
			})())
			var _1430_v77 _dafny.Set
			_ = _1430_v77
			_1430_v77 = _dafny.SetOf(_dafny.SeqOf(_this.F13()))
			_1429_v76 = (_1429_v76).Update(p0, Companion_Default___.SafeModuloInt((_1430_v77).Cardinality(), p0))
			r1 = Companion_Default___.SafeModuloInt(p2, p2)
			var _1431_v78 D9
			_ = _1431_v78
			_1431_v78 = Companion_D9_.Create_DC25_(_this.F13(), (_dafny.Zero).Minus(p0), _this.F13())
			var _1432_v79 _dafny.Map
			_ = _1432_v79
			_1432_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F13(), _this.F13())
			(_this).F13_set_(!(!((Companion_Default___.Fm20(_this.F13(), p0, (_1431_v78).Dtor_cf41(), globalState)).Equals(_1432_v79)) || (false)))
			(_this).F13_set_(!(_this.F13()))
			var _1433_v80 _dafny.Array
			_ = _1433_v80
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_34
			var _nw214 _dafny.Array
			_ = _nw214
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw214 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.CodePoint = (func(_1434_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1435_i5 _dafny.Int) _dafny.CodePoint {
						return _1434_v0
					}
				})(_1331_v0)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw214 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw214).ArraySet1CodePoint(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw214).ArraySet1CodePoint(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1433_v80 = _nw214
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1433_v80), 0))
			_ = _index264
			(_1433_v80).ArraySet1CodePoint(_1331_v0, (_index264).Int())
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1433_v80), 0))
			_ = _index265
			var _rhs229 _dafny.Int = _dafny.IntOfInt64(92)
			_ = _rhs229
			var _rhs230 _dafny.CodePoint = _1331_v0
			_ = _rhs230
			var _lhs146 _dafny.Array = _1433_v80
			_ = _lhs146
			var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1433_v80), 0))
			_ = _lhs147
			r1 = _rhs229
			(_lhs146).ArraySet1CodePoint(_rhs230, (_lhs147).Int())
		}
		var _1436_v81 _dafny.Array
		_ = _1436_v81
		var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw215
		_1436_v81 = _nw215
		var _1437_v82 _dafny.MultiSet
		_ = _1437_v82
		_1437_v82 = _dafny.MultiSetOf(_1436_v81)
		r0 = (_1437_v82).IsSubsetOf(_1437_v82)
		var _1438_v83 D1
		_ = _1438_v83
		_1438_v83 = Companion_D1_.Create_DC5_(_1349_v17, _1436_v81, p0)
		r1 = (_1438_v83).Dtor_cf10()
		return r0, r1
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
