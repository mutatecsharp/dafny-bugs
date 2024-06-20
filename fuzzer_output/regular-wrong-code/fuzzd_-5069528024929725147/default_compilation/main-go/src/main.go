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
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D5_.Create_DC13_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()))).Cardinality())).Cardinality(), _dafny.SeqOf(Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_()))).Dtor_cf23()) || (true), (_dafny.IntOfInt64(319)).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("enyymaag")).Cardinality())) > 0)
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(997), Companion_Default___.SafeDivisionInt((Companion_D11_.Create_DC26_(true, _dafny.IntOfInt64(212))).Dtor_cf49(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(897))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-486)
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(942), _dafny.IntOfInt64(311))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(942)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(311)) < 0) {
				_coll0.Add((_1_v0).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pay")).Cardinality()))).Cardinality())))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(-559)), (_dafny.SetOf(true)).IsProperSubsetOf(_dafny.SetOf(true, true)))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Map, p2 D0, globalState *GlobalState) _dafny.Int {
	return ((_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())).Plus((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(857), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-21), false)).Cardinality())).Cardinality())))).Cardinality()))).Minus(_dafny.IntOfInt64(791))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	if _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true), _dafny.SeqOf(true, true)) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-786), _dafny.IntOfInt64(-348)), Companion_D0_.Create_DC0_(false))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kgxwf")).Cardinality()))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _2_v0 _dafny.Int
				_2_v0 = interface{}(_compr_1).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kgxwf")).Cardinality()))).Contains(_2_v0) {
					_coll1.Add((_2_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-941))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg1 _dafny.Int) interface{} {
							return coer1(arg1)
						}
					}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('b')
					}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())))).Cardinality()), _dafny.IntOfInt64(-24))
				}
			}
			return _coll1.ToMap()
		}(), Companion_D0_.Create_DC0_(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(864), _dafny.IntOfInt64(366))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _4_v1 _dafny.Int
				_4_v1 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(864)).Cmp(_4_v1) <= 0) && ((_4_v1).Cmp(_dafny.IntOfInt64(366)) < 0) {
					_coll2.Add((_4_v1).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(490), false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(!(true), false, true)).Cardinality()))
				}
			}
			return _coll2.ToMap()
		}(), Companion_D0_.Create_DC0_(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-995))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality()), (_dafny.SetOf(true)).IsDisjointFrom(_dafny.SetOf(false, true)))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("csqbv")
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, p2 D2, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(939), _dafny.IntOfInt64(8))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _6_v0 _dafny.Int
			_6_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(939)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(8)) < 0) {
				_coll3.Add((_6_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), true)).Cardinality()), false)
			}
		}
		return _coll3.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ymwukergo")).Cardinality())), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aax")).Cardinality()), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ahojjj")).Cardinality())), false))).Union(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(516), _dafny.IntOfInt64(408))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _7_v1 _dafny.Int
				_7_v1 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(516)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(408)) < 0) {
					_coll5.Add((_7_v1).Plus(_dafny.IntOfInt64(531)), false)
				}
			}
			return _coll5.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, false)).Cardinality(), true), func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_8_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-995)
			}))).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _9_v2 _dafny.Int
				_9_v2 = interface{}(_compr_6).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_8_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-995)
				})), _9_v2) {
					_coll6.Add(Companion_Default___.SafeModuloInt(_9_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("agsobcif")).Cardinality())), !(false))
				}
			}
			return _coll6.ToMap()
		}())).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _10_v3 _dafny.Map
			_10_v3 = interface{}(_compr_4).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(516), _dafny.IntOfInt64(408))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _7_v1 _dafny.Int
					_7_v1 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(516)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(408)) < 0) {
						_coll7.Add((_7_v1).Plus(_dafny.IntOfInt64(531)), false)
					}
				}
				return _coll7.ToMap()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, false)).Cardinality(), true), func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_8_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-995)
				}))).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _9_v2 _dafny.Int
					_9_v2 = interface{}(_compr_8).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}(func(_8_i0 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(-995)
					})), _9_v2) {
						_coll8.Add(Companion_Default___.SafeModuloInt(_9_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("agsobcif")).Cardinality())), !(false))
					}
				}
				return _coll8.ToMap()
			}()), _10_v3) {
				_coll4.Add(_10_v3)
			}
		}
		return _coll4.ToSet()
	}())).Union(func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(899), false))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _11_v4 _dafny.Map
			_11_v4 = interface{}(_compr_9).(_dafny.Map)
			if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(899), false))).Contains(_11_v4) {
				_coll9.Add(_11_v4)
			}
		}
		return _coll9.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm12(p0 D0, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('j')
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) bool {
	return !(func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(189))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})))).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _13_v0 _dafny.Sequence
			_13_v0 = interface{}(_compr_10).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(189))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})))).Contains(_13_v0) {
				_coll10.Add(_13_v0, false)
			}
		}
		return _coll10.ToMap()
	}()).Equals(func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ogphrbwfb")).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(407), (_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality()))).Cardinality()))).Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _14_v1 _dafny.Sequence
			_14_v1 = interface{}(_compr_11).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ogphrbwfb")).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(407), (_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality()))).Cardinality())), _14_v1) {
				_coll11.Add(_14_v1, !(true))
			}
		}
		return _coll11.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_())
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC5_(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.SeqOf(!(true))))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(true)), true), _dafny.SeqOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_(_dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('l')), true, _dafny.UnicodeSeqOfUtf8Bytes("mlnbv"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(616))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_15_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(771), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vojrh")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-52), !(true))).Merge((func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality())))))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _16_v0 _dafny.Int
			_16_v0 = interface{}(_compr_12).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()))))), _16_v0) {
				_coll12.Add((_16_v0).Times(_dafny.IntOfInt64(492)), false)
			}
		}
		return _coll12.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-583), !(false))))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wfd"), _dafny.UnicodeSeqOfUtf8Bytes("pkkiujvqb")), _dafny.UnicodeSeqOfUtf8Bytes("hhqapw"))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC5_((_dafny.MultiSetOf(false)).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-244)))).Difference(_dafny.MultiSetOf((func() _dafny.Set {
		var _coll13 = _dafny.NewBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(220)), _dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(148)))).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _17_v0 _dafny.Sequence
			_17_v0 = interface{}(_compr_13).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(220)), _dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(148)))).Contains(_17_v0) {
				_coll13.Add(_17_v0)
			}
		}
		return _coll13.ToSet()
	}()).Cardinality()))).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(207))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(872))).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_18_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(226)
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(!(!(true)), true, true, !(true), !(true))).Difference((_dafny.MultiSetOf(true)).Union(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(!(false), !(true), false))).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(731))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC9_()
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, !(true), true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-370), (_dafny.MultiSetOf(true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(false)), func() _dafny.Map {
		var _coll14 = _dafny.NewMapBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(215), _dafny.IntOfInt64(835))); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_14).(_dafny.Int)
			if ((_dafny.IntOfInt64(215)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(835)) < 0) {
				_coll14.Add((_19_v0).Times((func() _dafny.Set {
					var _coll15 = _dafny.NewBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-54), _dafny.IntOfInt64(615))); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _21_v1 _dafny.Int
						_21_v1 = interface{}(_compr_15).(_dafny.Int)
						if ((_dafny.IntOfInt64(-54)).Cmp(_21_v1) <= 0) && ((_21_v1).Cmp(_dafny.IntOfInt64(615)) < 0) {
							_coll15.Add((_21_v1).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vrmfs")).Cardinality()))).Cardinality())))
						}
					}
					return _coll15.ToSet()
				}()).Cardinality()), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(778))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_20_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gavh")).Cardinality())
				}))).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hdtu")).Cardinality()), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), true)).Cardinality())).Cardinality())).Cardinality())
			}
		}
		return _coll14.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(false), _dafny.SeqOf(true, true)) {
		return _dafny.CodePoint('t')
	} else {
		return _dafny.CodePoint('v')
	}
}
func (_static *CompanionStruct_Default___) Fm34(p0 D3, p1 bool, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("r"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vvnpdwre"), _dafny.UnicodeSeqOfUtf8Bytes("uqmklfrag")))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D5 = Companion_D5_.Create_DC14_(false, _dafny.IntOfInt64(49), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(990))).Cardinality()))
	_ = _source0
	if _source0.Is_DC13() {
		var _22___mcc_h0 bool = _source0.Get_().(D5_DC13).Cf23
		_ = _22___mcc_h0
		var _23___mcc_h1 _dafny.Int = _source0.Get_().(D5_DC13).Cf24
		_ = _23___mcc_h1
		var _24___mcc_h2 _dafny.Sequence = _source0.Get_().(D5_DC13).Cf25
		_ = _24___mcc_h2
		var _25_cf25 _dafny.Sequence = _24___mcc_h2
		_ = _25_cf25
		var _26_cf24 _dafny.Int = _23___mcc_h1
		_ = _26_cf24
		var _27_cf23 bool = _22___mcc_h0
		_ = _27_cf23
		return _25_cf25
	} else if _source0.Is_DC14() {
		var _28___mcc_h3 bool = _source0.Get_().(D5_DC14).Cf26
		_ = _28___mcc_h3
		var _29___mcc_h4 _dafny.Int = _source0.Get_().(D5_DC14).Cf27
		_ = _29___mcc_h4
		var _30___mcc_h5 _dafny.Int = _source0.Get_().(D5_DC14).Cf28
		_ = _30___mcc_h5
		var _31_cf28 _dafny.Int = _30___mcc_h5
		_ = _31_cf28
		var _32_cf27 _dafny.Int = _29___mcc_h4
		_ = _32_cf27
		var _33_cf26 bool = _28___mcc_h3
		_ = _33_cf26
		return _dafny.SeqOf(Companion_D3_.Create_DC9_())
	} else {
		var _34___mcc_h6 _dafny.Map = _source0.Get_().(D5_DC12).Cf22
		_ = _34___mcc_h6
		var _35_cf22 _dafny.Map = _34___mcc_h6
		_ = _35_cf22
		if false {
			return _dafny.SeqOf(Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_())
		} else {
			return _dafny.SeqOf(Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_(), Companion_D3_.Create_DC9_())
		}
	}
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Map {
	if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(673), _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cwaejs")).Cardinality()))).Cardinality())), _dafny.IntOfInt64(213), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-927))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_36_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()
	}))).Cardinality())))).Cardinality(), _dafny.IntOfInt64(732), (_dafny.SetOf(true)).Cardinality())).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('m'))).Cardinality())) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(515), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_37_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(!(false), true)).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("y"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false, false)).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("ay")))
	}
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false)).Union((_dafny.SetOf(false, false)).Difference(_dafny.SetOf((Companion_D5_.Create_DC13_(!(!(false)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(546))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))).Cardinality()), _dafny.SeqOf(Companion_D3_.Create_DC9_()))).Dtor_cf23(), false)))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.CodePoint('s'))).Cardinality())).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _39_v0 _dafny.Int
			_39_v0 = interface{}(_compr_16).(_dafny.Int)
			if (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.CodePoint('s'))).Cardinality())).Contains(_39_v0) {
				_coll16.Add((_39_v0).Times((_dafny.MultiSetOf(_dafny.CodePoint('q'))).Cardinality()), !(!(!(false))))
			}
		}
		return _coll16.ToMap()
	}()).Merge((func() _dafny.Map {
		if false {
			return func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(715), _dafny.IntOfInt64(203))); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _40_v1 _dafny.Int
					_40_v1 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(715)).Cmp(_40_v1) <= 0) && ((_40_v1).Cmp(_dafny.IntOfInt64(203)) < 0) {
						_coll17.Add(Companion_Default___.SafeDivisionInt(_40_v1, _dafny.IntOfInt64(-514)), true)
					}
				}
				return _coll17.ToMap()
			}()
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.CodePoint('p'), _dafny.CodePoint('m'))).Cardinality(), true)
	})())
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(true), _dafny.SeqOf(false, false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('e'))).Cardinality()), !(!(true))))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_41_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dce")).Cardinality()))
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oeo")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(848))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_42_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(157)))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.CodePoint, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC13_(!(!(!(!(true))) || (false)), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(556), _dafny.IntOfInt64(662)), _dafny.SeqOf(Companion_D3_.Create_DC9_()))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll18 = _dafny.NewBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(604), _dafny.IntOfInt64(271))); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _43_v0 _dafny.Int
			_43_v0 = interface{}(_compr_18).(_dafny.Int)
			if ((_dafny.IntOfInt64(604)).Cmp(_43_v0) <= 0) && ((_43_v0).Cmp(_dafny.IntOfInt64(271)) < 0) {
				_coll18.Add(Companion_Default___.SafeModuloInt(_43_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("po")).Cardinality())))
			}
		}
		return _coll18.ToSet()
	}()).Intersection((_dafny.SetOf((_dafny.SetOf(!(!(false)))).Cardinality())).Union(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_44_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())
	}))).Cardinality()), false)).Cardinality(), (_dafny.SetOf(true, true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm43(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(72))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_45_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(993)))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC14_((_dafny.MultiSetOf((_dafny.SetOf(false)).Cardinality())).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dwcbpk")).Cardinality())))), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_46_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})))).Cardinality()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("evt")).Cardinality())), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Cardinality(), !(false))).Cardinality()).Minus(_dafny.IntOfInt64(305)))
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.MultiSet, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, !(true)))).Cardinality())).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qgmkk")).Cardinality())) != 0)
}
func (_static *CompanionStruct_Default___) Fm46(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SetOf(_dafny.IntOfInt64(-257), _dafny.IntOfInt64(309), (_dafny.SetOf(!(false))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func() _dafny.Set {
		var _coll19 = _dafny.NewBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-750), _dafny.IntOfInt64(898))); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _47_v0 _dafny.Int
			_47_v0 = interface{}(_compr_19).(_dafny.Int)
			if ((_dafny.IntOfInt64(-750)).Cmp(_47_v0) <= 0) && ((_47_v0).Cmp(_dafny.IntOfInt64(898)) < 0) {
				_coll19.Add((_47_v0).Minus((_dafny.MultiSetOf(_dafny.SeqOf(Companion_D3_.Create_DC9_()))).Cardinality()))
			}
		}
		return _coll19.ToSet()
	}()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(_dafny.IntOfInt64(385)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC5_(_dafny.MultiSetOf(!(false), true, !(false), true, true)), (func() bool {
		if true {
			return !(true)
		}
		return false
	})())
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _48_v0 _dafny.Array
	_ = _48_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
	_ = _nw0
	_48_v0 = _nw0
	var _49_v1 _dafny.Array
	_ = _49_v1
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
	_ = _nw1
	_49_v1 = _nw1
	var _50_v2 _dafny.Int
	_ = _50_v2
	_50_v2 = _dafny.IntOfInt64(548)
	var _51_v3 bool
	_ = _51_v3
	_51_v3 = false
	var _52_v4 _dafny.Sequence
	_ = _52_v4
	_52_v4 = _dafny.SeqOf(_51_v3, _51_v3)
	var _53_v5 _dafny.MultiSet
	_ = _53_v5
	_53_v5 = _dafny.MultiSetOf((_dafny.Zero).Minus(_50_v2), _dafny.IntOfUint32((_52_v4).Cardinality()))
	var _54_v6 _dafny.Array
	_ = _54_v6
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.CodePoint = func(_55_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw2 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw2).ArraySet1CodePoint(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw2).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_54_v6 = _nw2
	var _56_globalState *GlobalState
	_ = _56_globalState
	var _nw3 *GlobalState = New_GlobalState_()
	_ = _nw3
	_nw3.Ctor__(true, true, false, true, _dafny.IntOfInt64(432), _dafny.IntOfInt64(-409), _dafny.IntOfInt64(541), _48_v0, _dafny.IntOfInt64(766), true, _dafny.CodePoint('n'), false, _dafny.IntOfInt64(-713), _49_v1, _dafny.IntOfInt64(804), _53_v5, _54_v6)
	_56_globalState = _nw3
	var _57_v7 _dafny.Set
	_ = _57_v7
	_57_v7 = _dafny.SetOf(_50_v2)
	var _58_v8 _dafny.Sequence
	_ = _58_v8
	_58_v8 = _dafny.SeqOf(_dafny.SetOf(_50_v2), _57_v7, _57_v7)
	var _59_v9 _dafny.Map
	_ = _59_v9
	_59_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_58_v8).Cardinality()), (_57_v7).Cardinality())
	var _60_v10 _dafny.Sequence
	_ = _60_v10
	_60_v10 = _dafny.UnicodeSeqOfUtf8Bytes("yvxdaq")
	var _61_v11 _dafny.CodePoint
	_ = _61_v11
	_61_v11 = _dafny.CodePoint('v')
	var _62_v12 _dafny.Map
	_ = _62_v12
	_62_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v9, _dafny.Companion_Sequence_.Update(_60_v10, (Companion_Default___.SafeIndex((_57_v7).Cardinality(), _dafny.IntOfUint32((_60_v10).Cardinality()))).Uint32(), _61_v11))
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))
	_ = _index0
	(_49_v1).ArraySet1((func() _dafny.Sequence {
		if (_62_v12).Contains(_59_v9) {
			return (_62_v12).Get(_59_v9).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("angrdqnc")
	})(), (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))
	_ = _index1
	var _rhs0 bool = _51_v3
	_ = _rhs0
	var _rhs1 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("wagfakwvy")
	_ = _rhs1
	var _rhs2 bool = _51_v3
	_ = _rhs2
	var _lhs0 *GlobalState = _56_globalState
	_ = _lhs0
	var _lhs1 _dafny.Array = _49_v1
	_ = _lhs1
	var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))
	_ = _lhs2
	var _lhs3 *GlobalState = _56_globalState
	_ = _lhs3
	_lhs0.F3 = _rhs0
	(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
	_lhs3.F9 = _rhs2
	if _51_v3 {
		var _63_v13 _dafny.Map
		_ = _63_v13
		_63_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _51_v3)
		var _64_v14 D1
		_ = _64_v14
		_64_v14 = Companion_D1_.Create_DC4_(false, _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, _51_v3)), _50_v2)
		var _65_v15 _dafny.Array
		_ = _65_v15
		var _nwElement0_0 bool = _51_v3
		_ = _nwElement0_0
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(27))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_0, 0)
		(_nw4).ArraySet1(_51_v3, 1)
		(_nw4).ArraySet1(true, 2)
		(_nw4).ArraySet1(false, 3)
		(_nw4).ArraySet1(_51_v3, 4)
		(_nw4).ArraySet1(_51_v3, 5)
		(_nw4).ArraySet1(true, 6)
		(_nw4).ArraySet1((_64_v14).Dtor_cf9(), 7)
		(_nw4).ArraySet1(!(_51_v3), 8)
		(_nw4).ArraySet1(!(_51_v3), 9)
		(_nw4).ArraySet1(_51_v3, 10)
		(_nw4).ArraySet1(_51_v3, 11)
		(_nw4).ArraySet1(true, 12)
		(_nw4).ArraySet1(_51_v3, 13)
		(_nw4).ArraySet1(_51_v3, 14)
		(_nw4).ArraySet1(_51_v3, 15)
		(_nw4).ArraySet1(_51_v3, 16)
		(_nw4).ArraySet1(_51_v3, 17)
		(_nw4).ArraySet1(_51_v3, 18)
		(_nw4).ArraySet1(_51_v3, 19)
		(_nw4).ArraySet1(_51_v3, 20)
		(_nw4).ArraySet1(Companion_Default___.Fm13(_56_globalState), 21)
		(_nw4).ArraySet1(_51_v3, 22)
		(_nw4).ArraySet1(_51_v3, 23)
		(_nw4).ArraySet1(!(!(_51_v3)), 24)
		(_nw4).ArraySet1(_51_v3, 25)
		(_nw4).ArraySet1(_51_v3, 26)
		_65_v15 = _nw4
		var _66_v16 _dafny.Set
		_ = _66_v16
		_66_v16 = _dafny.SetOf(_65_v15)
		var _67_v17 *C5
		_ = _67_v17
		var _nw5 *C5 = New_C5_()
		_ = _nw5
		_nw5.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(446))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_68_v1 _dafny.Array, _69_v2 _dafny.Int, _70_v3 bool) func(_dafny.Int) _dafny.CodePoint {
			return func(_71_i1 _dafny.Int) _dafny.CodePoint {
				return ((_68_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_68_v1), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v2, _70_v3)).Cardinality(), _dafny.IntOfUint32(((_68_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_68_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_49_v1, _50_v2, _51_v3))), _51_v3, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(969), (_dafny.Zero).Minus(_50_v2)), _63_v13, _65_v15, (_66_v16).Difference(_66_v16), _51_v3)
		_67_v17 = _nw5
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))
		_ = _index2
		(_48_v0).ArraySet1(_dafny.IntOfInt64(884), (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))
		_ = _index3
		(_48_v0).ArraySet1(_dafny.IntOfInt64(7), (_index3).Int())
		var _72_v18 D11
		_ = _72_v18
		_72_v18 = Companion_D11_.Create_DC26_(_51_v3, (_dafny.SetOf((_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))).Cardinality())
		var _73_v19 D11
		_ = _73_v19
		_73_v19 = Companion_D11_.Create_DC28_(_72_v18)
		var _74_v20 D11
		_ = _74_v20
		_74_v20 = Companion_D11_.Create_DC28_(_73_v19)
		var _75_v21 D11
		_ = _75_v21
		_75_v21 = Companion_D11_.Create_DC28_(_74_v20)
		var _pat_let_tv0 = _73_v19
		_ = _pat_let_tv0
		_75_v21 = func(_pat_let0_0 D11) D11 {
			return func(_76_dt__update__tmp_h0 D11) D11 {
				return func(_pat_let1_0 D11) D11 {
					return func(_77_dt__update_hcf53_h0 D11) D11 {
						return Companion_D11_.Create_DC28_(_77_dt__update_hcf53_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_75_v21)
		if _51_v3 {
			var _78_v22 _dafny.Set
			_ = _78_v22
			_78_v22 = _dafny.SetOf((_67_v17).F30())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))
			_ = _index4
			(_48_v0).ArraySet1((_dafny.Zero).Minus((((_78_v22).Difference(_78_v22)).Difference((_78_v22).Union(_78_v22))).Cardinality()), (_index4).Int())
			_60_v10 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(803))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_79_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_80_i2 _dafny.Int) _dafny.CodePoint {
					return _79_v11
				}
			})(_61_v11)))
			var _81_v23 _dafny.Map
			_ = _81_v23
			_81_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v7).Difference(_57_v7), (func() _dafny.Sequence {
				if _51_v3 {
					return _52_v4
				}
				return _52_v4
			})())
			_81_v23 = (_81_v23).Update((_57_v7).Difference(_57_v7), _dafny.SeqOf((_67_v17).F30(), (_67_v17).F30(), _51_v3, !(_51_v3), !(!((_67_v17).F30()))))
			var _82_v24 D5
			_ = _82_v24
			_82_v24 = Companion_D5_.Create_DC14_((func() bool {
				if (_67_v17).F30() {
					return _51_v3
				}
				return (_67_v17).F30()
			})(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(467), _50_v2), (_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))
			_82_v24 = _82_v24
			var _83_v25 _dafny.Sequence
			_ = _83_v25
			_83_v25 = _dafny.SeqOf((_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))
			var _84_v26 _dafny.MultiSet
			_ = _84_v26
			_84_v26 = _dafny.MultiSetOf(_61_v11)
			(_56_globalState).F14 = (_83_v25).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_60_v10, (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)), (Companion_Default___.SafeIndex((_84_v26).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_60_v10, (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence))).Cardinality()))).Uint32(), _61_v11)).Cardinality()), _dafny.IntOfUint32((_83_v25).Cardinality()))).Uint32()).(_dafny.Int)
		} else {
			var _85_v27 *C9
			_ = _85_v27
			var _nw6 *C9 = New_C9_()
			_ = _nw6
			_nw6.Ctor__((_53_v5).Cardinality(), _50_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, (_67_v17).F30()), _65_v15, _66_v16, (_67_v17).F30())
			_85_v27 = _nw6
			var _86_v28 _dafny.Map
			_ = _86_v28
			_86_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v27, (_dafny.Zero).Minus(_85_v27.F22))
			var _87_v29 bool
			_ = _87_v29
			var _88_v30 bool
			_ = _88_v30
			var _89_v31 _dafny.Int
			_ = _89_v31
			var _out0 bool
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			_out0, _out1, _out2 = (_67_v17).M0((func() _dafny.Int {
				if _51_v3 {
					return (func() _dafny.Int {
						if (_86_v28).Contains(_85_v27) {
							return (_86_v28).Get(_85_v27).(_dafny.Int)
						}
						return (_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int)
					})()
				}
				return _50_v2
			})(), _51_v3, _85_v27.F22, _56_globalState)
			_87_v29 = _out0
			_88_v30 = _out1
			_89_v31 = _out2
			var _90_v32 _dafny.Map
			_ = _90_v32
			_90_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_67_v17).F30(), Companion_Default___.Fm13(_56_globalState))
			var _91_v33 _dafny.Map
			_ = _91_v33
			_91_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v9, _89_v31)
			var _92_v34 _dafny.Set
			_ = _92_v34
			_92_v34 = _dafny.SetOf(_88_v30)
			var _93_v35 D7
			_ = _93_v35
			_93_v35 = Companion_D7_.Create_DC16_(_92_v34)
			_90_v32 = (_90_v32).Update((_85_v27).Fm1(_91_v33, _56_globalState), ((_93_v35).Dtor_cf30()).IsProperSubsetOf(_92_v34))
			var _94_v36 *C2
			_ = _94_v36
			var _nw7 *C2 = New_C2_()
			_ = _nw7
			_nw7.Ctor__((_66_v16).Union(_66_v16), true)
			_94_v36 = _nw7
			var _95_v37 _dafny.Sequence
			_ = _95_v37
			_95_v37 = _dafny.SeqOf((_92_v34).Cardinality())
			var _96_v38 _dafny.Sequence
			_ = _96_v38
			_96_v38 = _dafny.SeqOf(_95_v37, _95_v37, _95_v37, _95_v37)
			var _97_v39 _dafny.Array
			_ = _97_v39
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw8
			_97_v39 = _nw8
			var _98_v40 D0
			_ = _98_v40
			_98_v40 = Companion_D0_.Create_DC2_(_88_v30, _50_v2, _97_v39)
			var _99_v41 _dafny.Map
			_ = _99_v41
			_99_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v29, _dafny.CodePoint('x'))
			var _100_v42 _dafny.Array
			_ = _100_v42
			var _nwElement0_1 _dafny.Sequence = _95_v37
			_ = _nwElement0_1
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(25))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_1, 0)
			(_nw9).ArraySet1(_95_v37, 1)
			(_nw9).ArraySet1(_dafny.SeqOf(_85_v27.F22, _85_v27.F22), 2)
			(_nw9).ArraySet1(Companion_Default___.Fm40(!(_88_v30), _88_v30, _56_globalState), 3)
			(_nw9).ArraySet1(_95_v37, 4)
			(_nw9).ArraySet1(_95_v37, 5)
			(_nw9).ArraySet1(_95_v37, 6)
			(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_95_v37, _95_v37), 7)
			(_nw9).ArraySet1((_96_v38).Select((Companion_Default___.SafeIndex(_85_v27.F22, _dafny.IntOfUint32((_96_v38).Cardinality()))).Uint32()).(_dafny.Sequence), 8)
			(_nw9).ArraySet1((func() _dafny.Sequence {
				if _87_v29 {
					return _95_v37
				}
				return _95_v37
			})(), 9)
			(_nw9).ArraySet1(_95_v37, 10)
			(_nw9).ArraySet1(_95_v37, 11)
			(_nw9).ArraySet1(_95_v37, 12)
			(_nw9).ArraySet1(_95_v37, 13)
			(_nw9).ArraySet1(_dafny.SeqOf(_50_v2, _50_v2, _50_v2), 14)
			(_nw9).ArraySet1(_95_v37, 15)
			(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_101_v31 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_102_i3 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_101_v31)
				}
			})(_89_v31))), _95_v37), 16)
			(_nw9).ArraySet1(_95_v37, 17)
			(_nw9).ArraySet1(_95_v37, 18)
			(_nw9).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-558))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_103_v37 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_104_i4 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_103_v37).Cardinality())
				}
			})(_95_v37))), 19)
			(_nw9).ArraySet1(_95_v37, 20)
			(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_95_v37, _dafny.SeqOf((_98_v40).Dtor_cf6())), 21)
			(_nw9).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_105_v27 *C9) func(_dafny.Int) _dafny.Int {
				return func(_106_i5 _dafny.Int) _dafny.Int {
					return _105_v27.F22
				}
			})(_85_v27))), 22)
			(_nw9).ArraySet1(_95_v37, 23)
			(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_50_v2, (_99_v41).Cardinality()), _dafny.SeqOf(_85_v27.F22)), 24)
			_100_v42 = _nw9
			_100_v42 = _100_v42
			var _107_v45 _dafny.Array
			_ = _107_v45
			var _nwElement0_2 _dafny.Map = _59_v9
			_ = _nwElement0_2
			var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(26))
			_ = _nw10
			(_nw10).ArraySet1(_nwElement0_2, 0)
			(_nw10).ArraySet1(func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-832), _dafny.IntOfInt64(895))); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _108_v43 _dafny.Int
					_108_v43 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(-832)).Cmp(_108_v43) <= 0) && ((_108_v43).Cmp(_dafny.IntOfInt64(895)) < 0) {
						_coll20.Add(Companion_Default___.SafeDivisionInt(_108_v43, _89_v31), _50_v2)
					}
				}
				return _coll20.ToMap()
			}(), 1)
			(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(193), _85_v27.F22), 2)
			(_nw10).ArraySet1(_59_v9, 3)
			(_nw10).ArraySet1(_59_v9, 4)
			(_nw10).ArraySet1(_59_v9, 5)
			(_nw10).ArraySet1(_59_v9, 6)
			(_nw10).ArraySet1(_59_v9, 7)
			(_nw10).ArraySet1(_59_v9, 8)
			(_nw10).ArraySet1(_59_v9, 9)
			(_nw10).ArraySet1(_59_v9, 10)
			(_nw10).ArraySet1(_59_v9, 11)
			(_nw10).ArraySet1(_59_v9, 12)
			(_nw10).ArraySet1(_59_v9, 13)
			(_nw10).ArraySet1(_59_v9, 14)
			(_nw10).ArraySet1(_59_v9, 15)
			(_nw10).ArraySet1(_59_v9, 16)
			(_nw10).ArraySet1(_59_v9, 17)
			(_nw10).ArraySet1(func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(765), _dafny.IntOfInt64(4))); ; {
					_compr_21, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _109_v44 _dafny.Int
					_109_v44 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(765)).Cmp(_109_v44) <= 0) && ((_109_v44).Cmp(_dafny.IntOfInt64(4)) < 0) {
						_coll21.Add(Companion_Default___.SafeModuloInt(_109_v44, _85_v27.F22), _dafny.IntOfUint32((_dafny.SeqOf((_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))).Cardinality()))
					}
				}
				return _coll21.ToMap()
			}(), 18)
			(_nw10).ArraySet1(_59_v9, 19)
			(_nw10).ArraySet1(_59_v9, 20)
			(_nw10).ArraySet1(_59_v9, 21)
			(_nw10).ArraySet1(_59_v9, 22)
			(_nw10).ArraySet1(_59_v9, 23)
			(_nw10).ArraySet1(_59_v9, 24)
			(_nw10).ArraySet1(_59_v9, 25)
			_107_v45 = _nw10
			var _110_v46 D7
			_ = _110_v46
			_110_v46 = Companion_D7_.Create_DC17_(_107_v45, _87_v29, (_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))
			var _111_v47 D5
			_ = _111_v47
			_111_v47 = Companion_D5_.Create_DC14_((_92_v34).IsDisjointFrom(_92_v34), (_110_v46).Dtor_cf33(), _dafny.IntOfInt64(453))
			var _112_v48 D3
			_ = _112_v48
			_112_v48 = Companion_D3_.Create_DC9_()
			var _113_v49 _dafny.Sequence
			_ = _113_v49
			_113_v49 = _dafny.SeqOf(_112_v48, _112_v48, _112_v48)
			var _114_v50 _dafny.Sequence
			_ = _114_v50
			_114_v50 = _113_v49
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_65_v15), 0))
			_ = _index5
			(_65_v15).ArraySet1(_88_v30, (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_65_v15), 0))
			_ = _index6
			var _rhs3 bool = false
			_ = _rhs3
			var _rhs4 D5 = Companion_Default___.Fm44(Companion_Default___.SafeDivisionInt(_50_v2, (_67_v17).Fm0(_dafny.IntOfInt64(930), _61_v11, _56_globalState)), _56_globalState)
			_ = _rhs4
			var _rhs5 _dafny.Sequence = _114_v50
			_ = _rhs5
			var _rhs6 bool = !((_67_v17).F30())
			_ = _rhs6
			var _lhs4 *GlobalState = _56_globalState
			_ = _lhs4
			var _lhs5 _dafny.Array = _65_v15
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_65_v15), 0))
			_ = _lhs6
			_lhs4.F9 = _rhs3
			_111_v47 = _rhs4
			_114_v50 = _rhs5
			(_lhs5).ArraySet1(_rhs6, (_lhs6).Int())
		}
		var _115_v51 _dafny.Sequence
		_ = _115_v51
		_115_v51 = _dafny.SeqOf((_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))
		var _116_v52 *C1
		_ = _116_v52
		var _nw11 *C1 = New_C1_()
		_ = _nw11
		_nw11.Ctor__()
		_116_v52 = _nw11
		var _117_v53 _dafny.Array
		_ = _117_v53
		var _nwElement0_3 *C1 = _116_v52
		_ = _nwElement0_3
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(28))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_3, 0)
		(_nw12).ArraySet1(_116_v52, 1)
		(_nw12).ArraySet1(_116_v52, 2)
		(_nw12).ArraySet1(_116_v52, 3)
		(_nw12).ArraySet1(_116_v52, 4)
		(_nw12).ArraySet1(_116_v52, 5)
		(_nw12).ArraySet1(_116_v52, 6)
		(_nw12).ArraySet1(_116_v52, 7)
		(_nw12).ArraySet1(_116_v52, 8)
		(_nw12).ArraySet1(_116_v52, 9)
		(_nw12).ArraySet1(_116_v52, 10)
		(_nw12).ArraySet1(_116_v52, 11)
		(_nw12).ArraySet1((func() *C1 {
			if _51_v3 {
				return _116_v52
			}
			return _116_v52
		})(), 12)
		(_nw12).ArraySet1(_116_v52, 13)
		(_nw12).ArraySet1(_116_v52, 14)
		(_nw12).ArraySet1(_116_v52, 15)
		(_nw12).ArraySet1(_116_v52, 16)
		(_nw12).ArraySet1(_116_v52, 17)
		(_nw12).ArraySet1(_116_v52, 18)
		(_nw12).ArraySet1(_116_v52, 19)
		(_nw12).ArraySet1(_116_v52, 20)
		(_nw12).ArraySet1(_116_v52, 21)
		(_nw12).ArraySet1(_116_v52, 22)
		(_nw12).ArraySet1(_116_v52, 23)
		(_nw12).ArraySet1(_116_v52, 24)
		(_nw12).ArraySet1(_116_v52, 25)
		(_nw12).ArraySet1(_116_v52, 26)
		(_nw12).ArraySet1(_116_v52, 27)
		_117_v53 = _nw12
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_117_v53), 0))
		_ = _index7
		(_117_v53).ArraySet1(_116_v52, (_index7).Int())
		var _118_v54 T2
		_ = _118_v54
		var _nw13 *C8 = New_C8_()
		_ = _nw13
		_nw13.Ctor__((_53_v5).Cardinality(), _51_v3, _65_v15, (_66_v16).Union(_66_v16), _dafny.Companion_Sequence_.IsPrefixOf(_52_v4, _52_v4), ((_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int)).Plus(_50_v2), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), (_67_v17).F30()))
		_118_v54 = _nw13
		var _119_v55 _dafny.Sequence
		_ = _119_v55
		_119_v55 = _dafny.SeqOf(_60_v10, (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(589))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_120_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), (_67_v17).F29(), _dafny.UnicodeSeqOfUtf8Bytes("geplenmxg"))
		var _121_v56 _dafny.Sequence
		_ = _121_v56
		_121_v56 = _dafny.SeqOf(_53_v5)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_117_v53), 0))
		_ = _index8
		var _rhs7 bool = (_67_v17).F30()
		_ = _rhs7
		var _rhs8 _dafny.Int = (_118_v54).F20()
		_ = _rhs8
		var _rhs9 _dafny.Sequence = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xwcwqwn"), (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_119_v55).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xwcwqwn"), (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence))).Cardinality()))).Uint32(), ((_67_v17).F29()).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_59_v9).Contains(_dafny.IntOfUint32((_121_v56).Cardinality())) {
				return (_59_v9).Get(_dafny.IntOfUint32((_121_v56).Cardinality())).(_dafny.Int)
			}
			return (_118_v54).F20()
		})(), _dafny.IntOfUint32(((_67_v17).F29()).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()))
		_ = _rhs9
		var _rhs10 *C1 = _116_v52
		_ = _rhs10
		var _rhs11 T2 = _118_v54
		_ = _rhs11
		var _lhs7 *GlobalState = _56_globalState
		_ = _lhs7
		var _lhs8 *GlobalState = _56_globalState
		_ = _lhs8
		var _lhs9 _dafny.Array = _117_v53
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_117_v53), 0))
		_ = _lhs10
		_lhs7.F3 = _rhs7
		_lhs8.F14 = _rhs8
		_115_v51 = _rhs9
		(_lhs9).ArraySet1(_rhs10, (_lhs10).Int())
		_118_v54 = _rhs11
	} else {
		var _122_v57 _dafny.Array
		_ = _122_v57
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_1
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_123_v3 bool) func(_dafny.Int) bool {
				return func(_124_i7 _dafny.Int) bool {
					return _123_v3
				}
			})(_51_v3)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw14 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw14).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw14).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_122_v57 = _nw14
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_122_v57), 0))
		_ = _index9
		(_122_v57).ArraySet1((_51_v3) && (_51_v3), (_index9).Int())
		var _125_v58 _dafny.Set
		_ = _125_v58
		_125_v58 = _dafny.SetOf(_122_v57)
		var _126_v59 _dafny.Map
		_ = _126_v59
		_126_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, _51_v3)
		var _127_v60 *C8
		_ = _127_v60
		var _nw15 *C8 = New_C8_()
		_ = _nw15
		_nw15.Ctor__(_50_v2, _51_v3, _122_v57, _125_v58, _51_v3, _50_v2, _126_v59)
		_127_v60 = _nw15
		var _128_v61 _dafny.Sequence
		_ = _128_v61
		_128_v61 = _dafny.SeqOf(_127_v60)
		var _129_v62 _dafny.Set
		_ = _129_v62
		_129_v62 = _dafny.SetOf(_128_v61)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_122_v57), 0))
		_ = _index10
		(_122_v57).ArraySet1((_129_v62).IsProperSubsetOf(_129_v62), (_index10).Int())
		if (_52_v4).Select((Companion_Default___.SafeIndex(_127_v60.F23, _dafny.IntOfUint32((_52_v4).Cardinality()))).Uint32()).(bool) {
			var _130_v63 D10
			_ = _130_v63
			_130_v63 = Companion_D10_.Create_DC22_(_57_v7)
			var _131_v64 _dafny.Map
			_ = _131_v64
			_131_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_122_v57), 0))).Int()).(bool), _130_v63)
			var _pat_let_tv1 = _57_v7
			_ = _pat_let_tv1
			var _132_v65 _dafny.Int
			_ = _132_v65
			var _133_v66 bool
			_ = _133_v66
			var _134_v67 _dafny.Int
			_ = _134_v67
			var _out3 _dafny.Int
			_ = _out3
			var _out4 bool
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out3, _out4, _out5 = (_127_v60).M2(((_131_v64).Update(_51_v3, func(_pat_let2_0 D10) D10 {
				return func(_135_dt__update__tmp_h1 D10) D10 {
					return func(_pat_let3_0 _dafny.Set) D10 {
						return func(_136_dt__update_hcf42_h0 _dafny.Set) D10 {
							return Companion_D10_.Create_DC22_(_136_dt__update_hcf42_h0)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let2_0)
			}(_130_v63))).Cardinality(), _56_globalState)
			_132_v65 = _out3
			_133_v66 = _out4
			_134_v67 = _out5
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_54_v6), 0))
			_ = _index11
			(_54_v6).ArraySet1CodePoint(_dafny.CodePoint('o'), (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_54_v6), 0))
			_ = _index12
			(_54_v6).ArraySet1CodePoint(_61_v11, (_index12).Int())
			var _137_v68 _dafny.Sequence
			_ = _137_v68
			_137_v68 = _dafny.SeqOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(false, _51_v3, _133_v66, _133_v66, (_127_v60).F24()))
			var _138_v69 _dafny.Array
			_ = _138_v69
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
			_ = _nw16
			_138_v69 = _nw16
			var _139_v70 _dafny.Map
			_ = _139_v70
			_139_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v60.F23, _138_v69)
			var _140_v71 _dafny.MultiSet
			_ = _140_v71
			_140_v71 = _dafny.MultiSetOf(_138_v69, _138_v69, (func() _dafny.Array {
				if (_139_v70).Contains(_127_v60.F23) {
					return (_139_v70).Get(_127_v60.F23).(_dafny.Array)
				}
				return _138_v69
			})())
			var _rhs12 _dafny.Int = _134_v67
			_ = _rhs12
			var _rhs13 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_137_v68, _137_v68)
			_ = _rhs13
			var _rhs14 _dafny.Int = (_dafny.Zero).Minus(_134_v67)
			_ = _rhs14
			var _rhs15 _dafny.Array = _122_v57
			_ = _rhs15
			var _rhs16 _dafny.Int = (func() _dafny.Int {
				if (_140_v71).Contains(_138_v69) {
					return (_140_v71).Multiplicity(_138_v69)
				}
				return _127_v60.F23
			})()
			_ = _rhs16
			var _lhs11 *GlobalState = _56_globalState
			_ = _lhs11
			var _lhs12 *C8 = _127_v60
			_ = _lhs12
			_lhs11.F14 = _rhs12
			_137_v68 = _rhs13
			_134_v67 = _rhs14
			_122_v57 = _rhs15
			_lhs12.F23 = _rhs16
			var _141_v72 _dafny.Map
			_ = _141_v72
			_141_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v65, _53_v5)
			_141_v72 = (_141_v72).Merge((_141_v72).Merge(_141_v72))
			_50_v2 = _127_v60.F23
		} else {
			var _142_v73 T2
			_ = _142_v73
			var _nw17 *C9 = New_C9_()
			_ = _nw17
			_nw17.Ctor__(_50_v2, _50_v2, _126_v59, _122_v57, _125_v58, (_127_v60).F24())
			_142_v73 = _nw17
			var _143_v74 D10
			_ = _143_v74
			_143_v74 = Companion_D10_.Create_DC23_((_127_v60).F24(), (_127_v60).F24(), _142_v73)
			var _144_v75 D10
			_ = _144_v75
			_144_v75 = Companion_D10_.Create_DC24_(_143_v74)
			var _145_v76 _dafny.Array
			_ = _145_v76
			var _nwElement0_4 D10 = _144_v75
			_ = _nwElement0_4
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(5))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_4, 0)
			(_nw18).ArraySet1(Companion_D10_.Create_DC24_(_143_v74), 1)
			(_nw18).ArraySet1(_144_v75, 2)
			(_nw18).ArraySet1(_144_v75, 3)
			(_nw18).ArraySet1(Companion_D10_.Create_DC24_(_143_v74), 4)
			_145_v76 = _nw18
			var _146_v77 _dafny.Map
			_ = _146_v77
			_146_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, _145_v76)
			var _147_v78 _dafny.Array
			_ = _147_v78
			var _nwElement0_5 _dafny.Array = _145_v76
			_ = _nwElement0_5
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(5))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_5, 0)
			(_nw19).ArraySet1((func() _dafny.Array {
				if (_146_v77).Contains((_142_v73).F20()) {
					return (_146_v77).Get((_142_v73).F20()).(_dafny.Array)
				}
				return _145_v76
			})(), 1)
			(_nw19).ArraySet1(_145_v76, 2)
			(_nw19).ArraySet1(_145_v76, 3)
			(_nw19).ArraySet1(_145_v76, 4)
			_147_v78 = _nw19
			var _148_v79 D14
			_ = _148_v79
			_148_v79 = Companion_D14_.Create_DC34_(_54_v6)
			var _149_v80 _dafny.MultiSet
			_ = _149_v80
			_149_v80 = _dafny.MultiSetOf(_148_v79)
			var _150_v81 _dafny.Map
			_ = _150_v81
			_150_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(502)).Cmp(_50_v2) > 0, (_149_v80).Update(_148_v79, Companion_Default___.Abs(_127_v60.F23)))
			var _rhs17 _dafny.Array = _147_v78
			_ = _rhs17
			var _rhs18 bool = (_127_v60).F24()
			_ = _rhs18
			var _rhs19 _dafny.Map = _150_v81
			_ = _rhs19
			var _rhs20 _dafny.Int = _dafny.IntOfInt64(748)
			_ = _rhs20
			var _lhs13 *GlobalState = _56_globalState
			_ = _lhs13
			var _lhs14 *C8 = _127_v60
			_ = _lhs14
			_147_v78 = _rhs17
			_lhs13.F9 = _rhs18
			_150_v81 = _rhs19
			_lhs14.F23 = _rhs20
			var _151_v82 _dafny.MultiSet
			_ = _151_v82
			_151_v82 = _dafny.MultiSetOf((_122_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_122_v57), 0))).Int()).(bool))
			var _152_v83 _dafny.Array
			_ = _152_v83
			var _nwElement0_6 _dafny.Map = _59_v9
			_ = _nwElement0_6
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(7))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_6, 0)
			(_nw20).ArraySet1(_59_v9, 1)
			(_nw20).ArraySet1((_59_v9).Update(_127_v60.F23, (_dafny.SetOf(_51_v3, (_127_v60).F24())).Cardinality()), 2)
			(_nw20).ArraySet1(_59_v9, 3)
			(_nw20).ArraySet1(Companion_Default___.Fm30((_127_v60).F24(), _56_globalState), 4)
			(_nw20).ArraySet1(_59_v9, 5)
			(_nw20).ArraySet1(_59_v9, 6)
			_152_v83 = _nw20
			var _153_v84 D7
			_ = _153_v84
			_153_v84 = Companion_D7_.Create_DC17_(_152_v83, _51_v3, _127_v60.F23)
			var _154_v85 _dafny.Sequence
			_ = _154_v85
			_154_v85 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, (_127_v60).F24()))
			var _155_v86 bool
			_ = _155_v86
			var _156_v87 bool
			_ = _156_v87
			var _157_v88 _dafny.Int
			_ = _157_v88
			var _out6 bool
			_ = _out6
			var _out7 bool
			_ = _out7
			var _out8 _dafny.Int
			_ = _out8
			_out6, _out7, _out8 = (_127_v60).M0((func() _dafny.Int {
				if (_151_v82).Contains((_153_v84).Dtor_cf32()) {
					return (_151_v82).Multiplicity((_153_v84).Dtor_cf32())
				}
				return _50_v2
			})(), _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(201))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_158_v73 T2) func(_dafny.Int) _dafny.Map {
				return func(_159_i8 _dafny.Int) _dafny.Map {
					return _158_v73.F21()
				}
			})(_142_v73))), _154_v85), _127_v60.F23, _56_globalState)
			_155_v86 = _out6
			_156_v87 = _out7
			_157_v88 = _out8
			_156_v87 = (_142_v73).F18()
			var _160_v89 T0
			_ = _160_v89
			var _nw21 *C2 = New_C2_()
			_ = _nw21
			_nw21.Ctor__(_dafny.SetOf(_142_v73.F19(), _142_v73.F19(), _142_v73.F19(), _122_v57, _122_v57), false)
			_160_v89 = _nw21
			var _161_v90 *C2
			_ = _161_v90
			var _nw22 *C2 = New_C2_()
			_ = _nw22
			_nw22.Ctor__(_125_v58, (_dafny.IntOfInt64(562)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_60_v10).Cardinality()), _160_v89)).Cardinality()) == 0)
			_161_v90 = _nw22
			var _162_v91 _dafny.Sequence
			_ = _162_v91
			_162_v91 = _dafny.SeqOf(_161_v90)
			_161_v90 = (_162_v91).Select((Companion_Default___.SafeIndex((_50_v2).Plus(_dafny.IntOfUint32((_60_v10).Cardinality())), _dafny.IntOfUint32((_162_v91).Cardinality()))).Uint32()).(*C2)
			var _163_v92 bool
			_ = _163_v92
			var _164_v93 bool
			_ = _164_v93
			var _165_v94 _dafny.Int
			_ = _165_v94
			var _out9 bool
			_ = _out9
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out9, _out10, _out11 = (_127_v60).M0(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()), (_dafny.Zero).Minus((_142_v73).F20())), ((_142_v73).F18()) && ((_127_v60).F24()), _50_v2, _56_globalState)
			_163_v92 = _out9
			_164_v93 = _out10
			_165_v94 = _out11
		}
		var _166_v95 _dafny.Map
		_ = _166_v95
		_166_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, (_122_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_122_v57), 0))).Int()).(bool))
		var _167_v96 _dafny.MultiSet
		_ = _167_v96
		_167_v96 = _dafny.MultiSetOf(_60_v10)
		var _168_v98 _dafny.Map
		_ = _168_v98
		_168_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
			var _coll22 = _dafny.NewBuilder()
			_ = _coll22
			for _iter22 := _dafny.Iterate((_167_v96).Elements()); ; {
				_compr_22, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _169_v97 _dafny.Sequence
				_169_v97 = interface{}(_compr_22).(_dafny.Sequence)
				if (_167_v96).Contains(_169_v97) {
					_coll22.Add(_169_v97)
				}
			}
			return _coll22.ToSet()
		}(), (_122_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_122_v57), 0))).Int()).(bool))
		_166_v95 = (_166_v95).Update(_dafny.IntOfInt64(-78), (func() bool {
			if (_168_v98).Contains(_dafny.SetOf((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence))) {
				return (_168_v98).Get(_dafny.SetOf((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence))).(bool)
			}
			return _51_v3
		})())
		(_127_v60).F23 = Companion_Default___.SafeModuloInt(_127_v60.F23, Companion_Default___.SafeDivisionInt(_50_v2, _dafny.IntOfInt64(797)))
		var _170_v99 *C1
		_ = _170_v99
		var _nw23 *C1 = New_C1_()
		_ = _nw23
		_nw23.Ctor__()
		_170_v99 = _nw23
	}
	var _hi0 _dafny.Int = _50_v2
	_ = _hi0
	for _171_i9 := _50_v2; _171_i9.Cmp(_hi0) < 0; _171_i9 = _171_i9.Plus(_dafny.One) {
		var _172_v100 _dafny.Array
		_ = _172_v100
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_2
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_173_v3 bool) func(_dafny.Int) bool {
				return func(_174_i10 _dafny.Int) bool {
					return _173_v3
				}
			})(_51_v3)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw24 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw24).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw24).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_172_v100 = _nw24
		var _175_v101 _dafny.Set
		_ = _175_v101
		_175_v101 = _dafny.SetOf(_172_v100)
		var _176_v102 _dafny.Map
		_ = _176_v102
		_176_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, _51_v3)
		var _177_v103 *C8
		_ = _177_v103
		var _nw25 *C8 = New_C8_()
		_ = _nw25
		_nw25.Ctor__(_50_v2, _51_v3, _172_v100, _175_v101, _51_v3, _171_i9, _176_v102)
		_177_v103 = _nw25
		var _178_v104 _dafny.Map
		_ = _178_v104
		_178_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, _177_v103)
		var _179_v105 _dafny.Sequence
		_ = _179_v105
		_179_v105 = _dafny.SeqOf(_dafny.IntOfInt64(16), _177_v103.F23)
		var _180_v107 _dafny.Set
		_ = _180_v107
		_180_v107 = _dafny.SetOf(Companion_Default___.Fm33(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lx")).Cardinality()), _51_v3, _56_globalState), _61_v11)
		var _181_v108 _dafny.Array
		_ = _181_v108
		var _nwElement0_7 bool = (_178_v104).Equals(_178_v104)
		_ = _nwElement0_7
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(7))
		_ = _nw26
		(_nw26).ArraySet1(_nwElement0_7, 0)
		(_nw26).ArraySet1((func() _dafny.Set {
			var _coll23 = _dafny.NewBuilder()
			_ = _coll23
			for _iter23 := _dafny.Iterate((_179_v105).Elements()); ; {
				_compr_23, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _182_v106 _dafny.Int
				_182_v106 = interface{}(_compr_23).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_179_v105, _182_v106) {
					_coll23.Add((_182_v106).Minus(_dafny.IntOfInt64(481)))
				}
			}
			return _coll23.ToSet()
		}()).IsSubsetOf(_57_v7), 1)
		(_nw26).ArraySet1(false, 2)
		(_nw26).ArraySet1(false, 3)
		(_nw26).ArraySet1(_51_v3, 4)
		(_nw26).ArraySet1((_177_v103).F24(), 5)
		(_nw26).ArraySet1((_180_v107).IsProperSubsetOf(_180_v107), 6)
		_181_v108 = _nw26
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_181_v108), 0))
		_ = _index13
		(_181_v108).ArraySet1((_177_v103).F24(), (_index13).Int())
		var _183_v109 _dafny.Map
		_ = _183_v109
		_183_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_i9, _61_v11)
		var _184_v110 *C4
		_ = _184_v110
		var _nw27 *C4 = New_C4_()
		_ = _nw27
		_nw27.Ctor__((_183_v109).Cardinality(), _dafny.IntOfInt64(-688), _175_v101, true)
		_184_v110 = _nw27
		var _185_v111 _dafny.Sequence
		_ = _185_v111
		_185_v111 = _dafny.SeqOf(_184_v110)
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_181_v108), 0))
		_ = _index14
		(_181_v108).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_185_v111, _dafny.Companion_Sequence_.Concatenate(_185_v111, _185_v111))), (_index14).Int())
		var _186_v112 D3
		_ = _186_v112
		_186_v112 = Companion_D3_.Create_DC10_((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), _51_v3, (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_187_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_188_i11 _dafny.Int) _dafny.CodePoint {
				return _187_v11
			}
		})(_61_v11))), (_184_v110).F25())
		var _189_v113 _dafny.Int
		_ = _189_v113
		var _out12 _dafny.Int
		_ = _out12
		_out12 = (_177_v103).M1(_dafny.IntOfInt64(-149), _dafny.Companion_Sequence_.Contains((_186_v112).Dtor_cf18(), _dafny.CodePoint('k')), (_171_i9).Times(_171_i9), _51_v3, _56_globalState)
		_189_v113 = _out12
		_51_v3 = _51_v3
		(_56_globalState).F3 = (_177_v103).F24()
	}
	_52_v4 = _52_v4
	var _190_i12 _dafny.Int
	_ = _190_i12
	_190_i12 = _dafny.Zero
	{
		for (_50_v2).Cmp(_50_v2) == 0 {
			{
				if (_190_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_190_i12 = (_190_i12).Plus(_dafny.One)
				var _191_v114 _dafny.Map
				_ = _191_v114
				_191_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_52_v4, _59_v9)
				var _source1 D9 = Companion_D9_.Create_DC19_(_191_v114)
				_ = _source1
				if _source1.Is_DC20() {
					var _192___mcc_h0 _dafny.Map = _source1.Get_().(D9_DC20).Cf36
					_ = _192___mcc_h0
					var _193___mcc_h1 _dafny.Int = _source1.Get_().(D9_DC20).Cf37
					_ = _193___mcc_h1
					var _194___mcc_h2 bool = _source1.Get_().(D9_DC20).Cf38
					_ = _194___mcc_h2
					var _195___mcc_h3 _dafny.Map = _source1.Get_().(D9_DC20).Cf39
					_ = _195___mcc_h3
					var _196___mcc_h4 _dafny.CodePoint = _source1.Get_().(D9_DC20).Cf40
					_ = _196___mcc_h4
					var _197_cf40 _dafny.CodePoint = _196___mcc_h4
					_ = _197_cf40
					var _198_cf39 _dafny.Map = _195___mcc_h3
					_ = _198_cf39
					var _199_cf38 bool = _194___mcc_h2
					_ = _199_cf38
					var _200_cf37 _dafny.Int = _193___mcc_h1
					_ = _200_cf37
					var _201_cf36 _dafny.Map = _192___mcc_h0
					_ = _201_cf36
					var _202_v115 _dafny.Map
					_ = _202_v115
					_202_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_cf38, (func() _dafny.Int {
						if (_53_v5).Contains((_201_cf36).Cardinality()) {
							return (_53_v5).Multiplicity((_201_cf36).Cardinality())
						}
						return _200_cf37
					})())
					var _203_v116 D0
					_ = _203_v116
					_203_v116 = Companion_D0_.Create_DC0_(_199_cf38)
					var _204_v117 _dafny.Array
					_ = _204_v117
					var _len0_3 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_3
					var _nw28 _dafny.Array
					_ = _nw28
					if _len0_3.Cmp(_dafny.Zero) == 0 {
						_nw28 = _dafny.NewArray(_len0_3)
					} else {
						var _init3 func(_dafny.Int) bool = (func(_205_v3 bool) func(_dafny.Int) bool {
							return func(_206_i13 _dafny.Int) bool {
								return _205_v3
							}
						})(_51_v3)
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
					_204_v117 = _nw28
					var _207_v118 _dafny.Set
					_ = _207_v118
					_207_v118 = _dafny.SetOf(_204_v117)
					var _208_v119 *C4
					_ = _208_v119
					var _nw29 *C4 = New_C4_()
					_ = _nw29
					_nw29.Ctor__(_200_cf37, (_dafny.Zero).Minus(((_dafny.Zero).Minus(Companion_Default___.Fm7(_200_cf37, _202_v115, _203_v116, _56_globalState))).Minus(_200_cf37)), (_207_v118).Union(_207_v118), _51_v3)
					_208_v119 = _nw29
					var _209_v120 _dafny.Map
					_ = _209_v120
					_209_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, _204_v117)
					_209_v120 = ((_209_v120).Merge(_209_v120)).Merge(_209_v120)
					_209_v120 = (_209_v120).Update(_51_v3, _204_v117)
					var _210_v121 *C2
					_ = _210_v121
					var _nw30 *C2 = New_C2_()
					_ = _nw30
					_nw30.Ctor__(_207_v118, _51_v3)
					_210_v121 = _nw30
					var _211_v122 _dafny.Map
					_ = _211_v122
					_211_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(848), _210_v121)
					var _212_v123 _dafny.Map
					_ = _212_v123
					_212_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_cf37, _211_v122)
					_212_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(577), _211_v122)
				} else if _source1.Is_DC19() {
					var _213___mcc_h5 _dafny.Map = _source1.Get_().(D9_DC19).Cf35
					_ = _213___mcc_h5
					var _214_cf35 _dafny.Map = _213___mcc_h5
					_ = _214_cf35
					var _215_v124 _dafny.Array
					_ = _215_v124
					var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
					_ = _nw31
					_215_v124 = _nw31
					var _216_v125 _dafny.Set
					_ = _216_v125
					_216_v125 = _dafny.SetOf(_215_v124, _215_v124, _215_v124)
					var _217_v126 *C3
					_ = _217_v126
					var _nw32 *C3 = New_C3_()
					_ = _nw32
					_nw32.Ctor__(_215_v124, _216_v125, _51_v3)
					_217_v126 = _nw32
					var _218_v127 _dafny.Map
					_ = _218_v127
					_218_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v126, _49_v1)
					var _219_v128 _dafny.Map
					_ = _219_v128
					_219_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, _49_v1)
					var _220_v129 _dafny.Array
					_ = _220_v129
					var _nwElement0_8 _dafny.Array = _49_v1
					_ = _nwElement0_8
					var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(11))
					_ = _nw33
					(_nw33).ArraySet1(_nwElement0_8, 0)
					(_nw33).ArraySet1(_49_v1, 1)
					(_nw33).ArraySet1((func() _dafny.Array {
						if (_218_v127).Contains(_217_v126) {
							return (_218_v127).Get(_217_v126).(_dafny.Array)
						}
						return _49_v1
					})(), 2)
					(_nw33).ArraySet1(_49_v1, 3)
					(_nw33).ArraySet1(_49_v1, 4)
					(_nw33).ArraySet1(_49_v1, 5)
					(_nw33).ArraySet1(_49_v1, 6)
					(_nw33).ArraySet1((func() _dafny.Array {
						if (_219_v128).Contains(_51_v3) {
							return (_219_v128).Get(_51_v3).(_dafny.Array)
						}
						return _49_v1
					})(), 7)
					(_nw33).ArraySet1(_49_v1, 8)
					(_nw33).ArraySet1(_49_v1, 9)
					(_nw33).ArraySet1(_49_v1, 10)
					_220_v129 = _nw33
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_220_v129), 0))
					_ = _index15
					(_220_v129).ArraySet1(_49_v1, (_index15).Int())
					var _221_v130 _dafny.Map
					_ = _221_v130
					_221_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_52_v4, _dafny.SeqOf(_51_v3, _51_v3, _51_v3))
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_220_v129), 0))
					_ = _index16
					var _rhs21 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_52_v4, (func() _dafny.Sequence {
						if (_221_v130).Contains(_52_v4) {
							return (_221_v130).Get(_52_v4).(_dafny.Sequence)
						}
						return _52_v4
					})())
					_ = _rhs21
					var _rhs22 _dafny.Array = _49_v1
					_ = _rhs22
					var _rhs23 bool = !(!(_51_v3))
					_ = _rhs23
					var _lhs15 _dafny.Array = _220_v129
					_ = _lhs15
					var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_220_v129), 0))
					_ = _lhs16
					var _lhs17 *GlobalState = _56_globalState
					_ = _lhs17
					_52_v4 = _rhs21
					(_lhs15).ArraySet1(_rhs22, (_lhs16).Int())
					_lhs17.F3 = _rhs23
					(_56_globalState).F14 = _50_v2
					var _222_v131 D9
					_ = _222_v131
					_222_v131 = Companion_D9_.Create_DC19_(_214_cf35)
					var _223_v132 D9
					_ = _223_v132
					_223_v132 = Companion_D9_.Create_DC21_(_222_v131)
					var _224_v133 D9
					_ = _224_v133
					_224_v133 = Companion_D9_.Create_DC21_(_222_v131)
					var _pat_let_tv2 = _223_v132
					_ = _pat_let_tv2
					_224_v133 = func(_pat_let4_0 D9) D9 {
						return func(_225_dt__update__tmp_h2 D9) D9 {
							return func(_pat_let5_0 D9) D9 {
								return func(_226_dt__update_hcf41_h0 D9) D9 {
									return Companion_D9_.Create_DC21_(_226_dt__update_hcf41_h0)
								}(_pat_let5_0)
							}(_pat_let_tv2)
						}(_pat_let4_0)
					}(_224_v133)
					_215_v124 = _215_v124
				} else {
					var _227___mcc_h6 D9 = _source1.Get_().(D9_DC21).Cf41
					_ = _227___mcc_h6
					var _228_cf41 D9 = _227___mcc_h6
					_ = _228_cf41
					var _229_v134 _dafny.Array
					_ = _229_v134
					var _nw34 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
					_ = _nw34
					_229_v134 = _nw34
					var _230_v135 D3
					_ = _230_v135
					_230_v135 = Companion_D3_.Create_DC9_()
					var _231_v136 _dafny.Sequence
					_ = _231_v136
					_231_v136 = _dafny.SeqOf(_230_v135)
					var _232_v137 D5
					_ = _232_v137
					_232_v137 = Companion_D5_.Create_DC13_(_51_v3, _50_v2, _231_v136)
					var _233_v138 _dafny.Map
					_ = _233_v138
					_233_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, _51_v3)
					var _234_v139 _dafny.Array
					_ = _234_v139
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_4
					var _nw35 _dafny.Array
					_ = _nw35
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw35 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) bool = func(_235_i14 _dafny.Int) bool {
							return true
						}
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw35 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw35).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw35).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_234_v139 = _nw35
					var _236_v140 _dafny.Set
					_ = _236_v140
					_236_v140 = _dafny.SetOf(_234_v139)
					var _237_v141 *C7
					_ = _237_v141
					var _nw36 *C7 = New_C7_()
					_ = _nw36
					_nw36.Ctor__(_50_v2, _233_v138, _234_v139, _236_v140, _51_v3)
					_237_v141 = _nw36
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_229_v134), 0))
					_ = _index17
					(_229_v134).ArraySet1((func() *C7 {
						if (_232_v137).Dtor_cf23() {
							return _237_v141
						}
						return _237_v141
					})(), (_index17).Int())
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_229_v134), 0))
					_ = _index18
					(_229_v134).ArraySet1(_237_v141, (_index18).Int())
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))
					_ = _index19
					(_49_v1).ArraySet1((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), (_index19).Int())
					_59_v9 = (_59_v9).Update(_50_v2, (_50_v2).Plus(_50_v2))
					var _238_v142 _dafny.Map
					_ = _238_v142
					_238_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, _50_v2)
					var _239_v143 D0
					_ = _239_v143
					_239_v143 = Companion_D0_.Create_DC0_(!(true))
					(_56_globalState).F14 = Companion_Default___.Fm7(_50_v2, _238_v142, _239_v143, _56_globalState)
				}
				var _240_v144 _dafny.MultiSet
				_ = _240_v144
				_240_v144 = _dafny.MultiSetOf(_51_v3)
				var _241_v145 _dafny.Map
				_ = _241_v145
				_241_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_50_v2, _dafny.IntOfUint32(((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())), Companion_D2_.Create_DC5_(_240_v144))
				var _242_v146 D2
				_ = _242_v146
				_242_v146 = Companion_D2_.Create_DC5_(_240_v144)
				_241_v145 = (_241_v145).Update((_dafny.Zero).Minus(_50_v2), _242_v146)
				var _243_v147 _dafny.Array
				_ = _243_v147
				var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw37
				_243_v147 = _nw37
				var _rhs24 _dafny.Array = _243_v147
				_ = _rhs24
				var _rhs25 bool = _51_v3
				_ = _rhs25
				var _rhs26 bool = _51_v3
				_ = _rhs26
				var _rhs27 _dafny.Sequence = _60_v10
				_ = _rhs27
				var _lhs18 *GlobalState = _56_globalState
				_ = _lhs18
				var _lhs19 *GlobalState = _56_globalState
				_ = _lhs19
				_243_v147 = _rhs24
				_lhs18.F3 = _rhs25
				_lhs19.F9 = _rhs26
				_60_v10 = _rhs27
				var _244_v148 _dafny.Array
				_ = _244_v148
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_5
				var _nw38 _dafny.Array
				_ = _nw38
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw38 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Set = (func(_245_v3 bool) func(_dafny.Int) _dafny.Set {
						return func(_246_i15 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_245_v3)
						}
					})(_51_v3)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw38 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw38).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw38).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_244_v148 = _nw38
				var _247_v149 _dafny.Set
				_ = _247_v149
				_247_v149 = _dafny.SetOf(!(_51_v3))
				var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_244_v148), 0))
				_ = _index20
				(_244_v148).ArraySet1(_247_v149, (_index20).Int())
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_244_v148), 0))
				_ = _index21
				(_244_v148).ArraySet1(_247_v149, (_index21).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	if _51_v3 {
		if !(((func() _dafny.Int {
			if _51_v3 {
				return _50_v2
			}
			return _50_v2
		})()).Cmp(_50_v2) <= 0) {
			(_56_globalState).F14 = (func() _dafny.Int {
				if (_53_v5).Contains(_50_v2) {
					return (_53_v5).Multiplicity(_50_v2)
				}
				return _50_v2
			})()
			var _248_v150 _dafny.Sequence
			_ = _248_v150
			_248_v150 = _dafny.SeqOf(_60_v10)
			_248_v150 = _dafny.Companion_Sequence_.Update(_248_v150, (Companion_Default___.SafeIndex(_50_v2, _dafny.IntOfUint32((_248_v150).Cardinality()))).Uint32(), (_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence))
			var _249_v151 _dafny.Array
			_ = _249_v151
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw39
			_249_v151 = _nw39
			var _250_v152 _dafny.Sequence
			_ = _250_v152
			_250_v152 = _dafny.SeqOf(_50_v2)
			var _251_v153 _dafny.Sequence
			_ = _251_v153
			_251_v153 = _dafny.SeqOf(_250_v152, _250_v152)
			var _252_v154 _dafny.Array
			_ = _252_v154
			var _nwElement0_9 _dafny.Sequence = _250_v152
			_ = _nwElement0_9
			var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(10))
			_ = _nw40
			(_nw40).ArraySet1(_nwElement0_9, 0)
			(_nw40).ArraySet1(_250_v152, 1)
			(_nw40).ArraySet1((_251_v153).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_52_v4).Cardinality()), _dafny.IntOfUint32((_251_v153).Cardinality()))).Uint32()).(_dafny.Sequence), 2)
			(_nw40).ArraySet1(_250_v152, 3)
			(_nw40).ArraySet1(_250_v152, 4)
			(_nw40).ArraySet1(_250_v152, 5)
			(_nw40).ArraySet1(_250_v152, 6)
			(_nw40).ArraySet1(_250_v152, 7)
			(_nw40).ArraySet1(_250_v152, 8)
			(_nw40).ArraySet1((_250_v152), 9)
			_252_v154 = _nw40
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_249_v151), 0))
			_ = _index22
			(_249_v151).ArraySet1(_252_v154, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_249_v151), 0))
			_ = _index23
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw41
			(_249_v151).ArraySet1(_nw41, (_index23).Int())
			var _253_v155 _dafny.Array
			_ = _253_v155
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_6
			var _nw42 _dafny.Array
			_ = _nw42
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw42 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) bool = (func(_254_v3 bool) func(_dafny.Int) bool {
					return func(_255_i16 _dafny.Int) bool {
						return _254_v3
					}
				})(_51_v3)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw42 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw42).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw42).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_253_v155 = _nw42
			var _256_v156 _dafny.Set
			_ = _256_v156
			_256_v156 = _dafny.SetOf(_253_v155)
			var _257_v157 T0
			_ = _257_v157
			var _nw43 *C4 = New_C4_()
			_ = _nw43
			_nw43.Ctor__(_50_v2, _dafny.IntOfInt64(978), _256_v156, _51_v3)
			_257_v157 = _nw43
			_257_v157 = _257_v157
			var _258_v158 _dafny.Map
			_ = _258_v158
			_258_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_257_v157).F18(), (_257_v157).F18())
			var _259_v159 *C7
			_ = _259_v159
			var _nw44 *C7 = New_C7_()
			_ = _nw44
			_nw44.Ctor__((_250_v152).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-841), _dafny.IntOfUint32((_250_v152).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm45(_53_v5, (_258_v158).Cardinality(), _51_v3, _56_globalState), _253_v155, _257_v157.F17(), Companion_Default___.Fm13(_56_globalState))
			_259_v159 = _nw44
			_259_v159 = _259_v159
		} else {
			var _260_v160 _dafny.MultiSet
			_ = _260_v160
			_260_v160 = _dafny.MultiSetOf(Companion_Default___.Fm30(true, _56_globalState))
			_260_v160 = _260_v160
			(_56_globalState).F7 = _48_v0
			var _261_v161 _dafny.Array
			_ = _261_v161
			var _len0_7 _dafny.Int = _dafny.One
			_ = _len0_7
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = func(_262_i17 _dafny.Int) bool {
					return false
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw45 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw45).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw45).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_261_v161 = _nw45
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_261_v161), 0))
			_ = _index24
			(_261_v161).ArraySet1(_51_v3, (_index24).Int())
			var _263_v162 D0
			_ = _263_v162
			_263_v162 = Companion_D0_.Create_DC0_(Companion_Default___.Fm13(_56_globalState))
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_261_v161), 0))
			_ = _index25
			(_261_v161).ArraySet1((Companion_Default___.SafeModuloInt(_50_v2, _50_v2)).Cmp((func() _dafny.Int {
				if false {
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_60_v10, (Companion_Default___.SafeIndex(_50_v2, _dafny.IntOfUint32((_60_v10).Cardinality()))).Uint32(), Companion_Default___.Fm12(_263_v162, _56_globalState))).Cardinality())
				}
				return (_dafny.Zero).Minus(_50_v2)
			})()) < 0, (_index25).Int())
			var _264_v163 _dafny.MultiSet
			_ = _264_v163
			_264_v163 = _dafny.MultiSetOf(_261_v161)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_261_v161), 0))
			_ = _index26
			(_261_v161).ArraySet1(!(_264_v163).Equals(_264_v163), (_index26).Int())
			(_56_globalState).F14 = _50_v2
		}
		_60_v10 = (func() _dafny.Sequence {
			if true {
				return _dafny.Companion_Sequence_.Concatenate(_60_v10, Companion_Default___.Fm10(_56_globalState))
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("udophjwub")
		})()
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))
		_ = _index27
		(_49_v1).ArraySet1((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), (_index27).Int())
		var _265_v164 _dafny.Map
		_ = _265_v164
		_265_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, _51_v3)
		_51_v3 = (func() bool {
			if Companion_Default___.Fm13(_56_globalState) {
				return _51_v3
			}
			return ((_265_v164).Cardinality()).Cmp(_50_v2) > 0
		})()
		_50_v2 = _50_v2
	} else {
		var _266_v165 _dafny.Map
		_ = _266_v165
		_266_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, (_dafny.IntOfInt64(341)).Cmp(_dafny.IntOfUint32(((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())) <= 0)
		_266_v165 = (_266_v165).Update(_dafny.IntOfUint32((_dafny.SeqOf(_51_v3)).Cardinality()), !(_dafny.Companion_Sequence_.IsProperPrefixOf((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("kpts"))))
		(_56_globalState).F14 = _dafny.IntOfUint32(((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())
		(_56_globalState).F3 = (_51_v3) && (_51_v3)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_54_v6), 0))
		_ = _index28
		(_54_v6).ArraySet1CodePoint(_61_v11, (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_54_v6), 0))
		_ = _index29
		(_54_v6).ArraySet1CodePoint(_dafny.CodePoint('r'), (_index29).Int())
		var _267_v166 D2
		_ = _267_v166
		_267_v166 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32(((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		var _268_v167 D2
		_ = _268_v167
		_268_v167 = Companion_D2_.Create_DC7_(_267_v166)
		var _source2 D2 = _268_v167
		_ = _source2
		if _source2.Is_DC6() {
			var _269___mcc_h7 _dafny.Int = _source2.Get_().(D2_DC6).Cf13
			_ = _269___mcc_h7
			var _270_cf13 _dafny.Int = _269___mcc_h7
			_ = _270_cf13
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_48_v0), 0))
			_ = _index30
			(_48_v0).ArraySet1(_270_cf13, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_48_v0), 0))
			_ = _index31
			(_48_v0).ArraySet1(_270_cf13, (_index31).Int())
			var _271_v168 _dafny.Map
			_ = _271_v168
			_271_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, _51_v3)
			var _272_v169 _dafny.Array
			_ = _272_v169
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_8
			var _nw46 _dafny.Array
			_ = _nw46
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw46 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) bool = (func(_273_v3 bool) func(_dafny.Int) bool {
					return func(_274_i18 _dafny.Int) bool {
						return _273_v3
					}
				})(_51_v3)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw46 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw46).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw46).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_272_v169 = _nw46
			var _275_v170 _dafny.Set
			_ = _275_v170
			_275_v170 = _dafny.SetOf(_272_v169, _272_v169)
			var _276_v171 T0
			_ = _276_v171
			var _nw47 *C2 = New_C2_()
			_ = _nw47
			_nw47.Ctor__(_275_v170, _51_v3)
			_276_v171 = _nw47
			var _277_v172 _dafny.Map
			_ = _277_v172
			_277_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T0 {
				if (func() bool {
					if (_271_v168).Contains(_51_v3) {
						return (_271_v168).Get(_51_v3).(bool)
					}
					return _51_v3
				})() {
					return _276_v171
				}
				return _276_v171
			})(), _50_v2)
			_277_v172 = (_277_v172).Update(_276_v171, (_48_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_48_v0), 0))).Int()).(_dafny.Int))
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_48_v0), 0))
			_ = _index32
			(_48_v0).ArraySet1(_270_cf13, (_index32).Int())
			var _278_v173 _dafny.MultiSet
			_ = _278_v173
			_278_v173 = _dafny.MultiSetOf(!(!((func() bool {
				if (_266_v165).Contains((_dafny.Zero).Minus(_270_cf13)) {
					return (_266_v165).Get((_dafny.Zero).Minus(_270_cf13)).(bool)
				}
				return _51_v3
			})())), true, (_276_v171).F18(), (_276_v171).F18())
			_278_v173 = ((_278_v173).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_51_v3)))).Intersection(_278_v173)
		} else if _source2.Is_DC5() {
			var _279___mcc_h8 _dafny.MultiSet = _source2.Get_().(D2_DC5).Cf12
			_ = _279___mcc_h8
			var _280_cf12 _dafny.MultiSet = _279___mcc_h8
			_ = _280_cf12
			_51_v3 = _51_v3
			var _281_v174 D0
			_ = _281_v174
			_281_v174 = Companion_D0_.Create_DC0_(_51_v3)
			var _282_v175 _dafny.Array
			_ = _282_v175
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_9
			var _nw48 _dafny.Array
			_ = _nw48
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw48 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) bool = (func(_283_v3 bool) func(_dafny.Int) bool {
					return func(_284_i19 _dafny.Int) bool {
						return _283_v3
					}
				})(_51_v3)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw48 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw48).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw48).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_282_v175 = _nw48
			var _285_v176 _dafny.Set
			_ = _285_v176
			_285_v176 = _dafny.SetOf(_282_v175, _282_v175)
			var _286_v177 _dafny.Map
			_ = _286_v177
			_286_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, _51_v3)
			var _287_v178 *C8
			_ = _287_v178
			var _nw49 *C8 = New_C8_()
			_ = _nw49
			_nw49.Ctor__(Companion_Default___.Fm7(_50_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_51_v3), _dafny.IntOfInt64(274)), func(_pat_let6_0 D0) D0 {
				return func(_288_dt__update__tmp_h3 D0) D0 {
					return func(_pat_let7_0 bool) D0 {
						return func(_289_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_289_dt__update_hcf0_h0)
						}(_pat_let7_0)
					}(true)
				}(_pat_let6_0)
			}(_281_v174), _56_globalState), _51_v3, _282_v175, _285_v176, !_dafny.Companion_Sequence_.Contains(_60_v10, (_54_v6).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_54_v6), 0))).Int())), (func() _dafny.Int {
				if (_59_v9).Contains(_50_v2) {
					return (_59_v9).Get(_50_v2).(_dafny.Int)
				}
				return _50_v2
			})(), _286_v177)
			_287_v178 = _nw49
			var _290_v179 *C5
			_ = _290_v179
			var _nw50 *C5 = New_C5_()
			_ = _nw50
			_nw50.Ctor__(_60_v10, (_287_v178).F24(), _50_v2, _286_v177, _282_v175, _285_v176, _51_v3)
			_290_v179 = _nw50
			_290_v179 = _290_v179
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_282_v175), 0))
			_ = _index33
			(_282_v175).ArraySet1((_287_v178).F24(), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_282_v175), 0))
			_ = _index34
			(_282_v175).ArraySet1(!((_290_v179).F30()), (_index34).Int())
			var _291_v180 _dafny.Map
			_ = _291_v180
			_291_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_290_v179).F30(), (_287_v178).F24())
			var _292_v181 D0
			_ = _292_v181
			_292_v181 = Companion_D0_.Create_DC1_((_49_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_49_v1), 0))).Int()).(_dafny.Sequence), (_291_v180).Cardinality(), _287_v178.F23, _61_v11)
			var _pat_let_tv3 = _54_v6
			_ = _pat_let_tv3
			var _pat_let_tv4 = _54_v6
			_ = _pat_let_tv4
			var _pat_let_tv5 = _50_v2
			_ = _pat_let_tv5
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_282_v175), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_282_v175), 0))
			_ = _index36
			var _rhs28 bool = (_287_v178).F24()
			_ = _rhs28
			var _rhs29 _dafny.Int = (_50_v2).Minus((func(_pat_let8_0 D0) D0 {
				return func(_293_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let9_0 _dafny.CodePoint) D0 {
						return func(_294_dt__update_hcf4_h0 _dafny.CodePoint) D0 {
							return func(_pat_let10_0 _dafny.Int) D0 {
								return func(_295_dt__update_hcf3_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC1_((_293_dt__update__tmp_h4).Dtor_cf1(), (_293_dt__update__tmp_h4).Dtor_cf2(), _295_dt__update_hcf3_h0, _294_dt__update_hcf4_h0)
								}(_pat_let10_0)
							}(_pat_let_tv5)
						}(_pat_let9_0)
					}((_pat_let_tv4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_pat_let_tv3), 0))).Int()))
				}(_pat_let8_0)
			}(_292_v181)).Dtor_cf3())
			_ = _rhs29
			var _rhs30 _dafny.Int = _287_v178.F23
			_ = _rhs30
			var _rhs31 bool = !((_287_v178.F23).Cmp(_50_v2) > 0)
			_ = _rhs31
			var _lhs20 _dafny.Array = _282_v175
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_282_v175), 0))
			_ = _lhs21
			var _lhs22 *GlobalState = _56_globalState
			_ = _lhs22
			var _lhs23 _dafny.Array = _282_v175
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_282_v175), 0))
			_ = _lhs24
			(_lhs20).ArraySet1(_rhs28, (_lhs21).Int())
			_50_v2 = _rhs29
			_lhs22.F14 = _rhs30
			(_lhs23).ArraySet1(_rhs31, (_lhs24).Int())
		} else {
			var _296___mcc_h9 D2 = _source2.Get_().(D2_DC7).Cf14
			_ = _296___mcc_h9
			var _297_cf14 D2 = _296___mcc_h9
			_ = _297_cf14
			var _298_v183 _dafny.Map
			_ = _298_v183
			_298_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, _50_v2)
			var _299_v184 _dafny.Array
			_ = _299_v184
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_10
			var _nw51 _dafny.Array
			_ = _nw51
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw51 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) bool = (func(_300_v10 _dafny.Sequence, _301_v3 bool, _302_v1 _dafny.Array, _303_v2 _dafny.Int) func(_dafny.Int) bool {
					return func(_304_i20 _dafny.Int) bool {
						return (Companion_D3_.Create_DC10_(_300_v10, _301_v3, (_302_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_302_v1), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("go"), _303_v2)).Dtor_cf17()
					}
				})(_60_v10, _51_v3, _49_v1, _50_v2)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw51 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw51).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw51).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_299_v184 = _nw51
			var _305_v185 *C7
			_ = _305_v185
			var _nw52 *C7 = New_C7_()
			_ = _nw52
			_nw52.Ctor__(_dafny.IntOfInt64(900), func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter24 := _dafny.Iterate((_298_v183).Keys().Elements()); ; {
					_compr_24, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _306_v182 _dafny.CodePoint
					_306_v182 = interface{}(_compr_24).(_dafny.CodePoint)
					if (_298_v183).Contains(_306_v182) {
						_coll24.Add(_306_v182, false)
					}
				}
				return _coll24.ToMap()
			}(), _299_v184, _dafny.SetOf(_299_v184, _299_v184, _299_v184, _299_v184, _299_v184), _51_v3)
			_305_v185 = _nw52
			_305_v185 = _305_v185
			var _307_v186 _dafny.Map
			_ = _307_v186
			_307_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_52_v4).Select((Companion_Default___.SafeIndex(_50_v2, _dafny.IntOfUint32((_52_v4).Cardinality()))).Uint32()).(bool), _57_v7)
			var _308_v187 _dafny.Map
			_ = _308_v187
			_308_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_54_v6).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_54_v6), 0))).Int()), _61_v11)
			var _309_v188 _dafny.Map
			_ = _309_v188
			_309_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v187, _307_v186)
			var _310_v189 _dafny.MultiSet
			_ = _310_v189
			_310_v189 = _dafny.MultiSetOf(_51_v3, _51_v3)
			var _311_v190 _dafny.Set
			_ = _311_v190
			_311_v190 = _dafny.SetOf(_299_v184)
			var _312_v191 _dafny.Array
			_ = _312_v191
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_11
			var _nw53 _dafny.Array
			_ = _nw53
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw53 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Map = (func(_313_v9 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_314_i21 _dafny.Int) _dafny.Map {
						return _313_v9
					}
				})(_59_v9)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw53 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw53).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw53).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_312_v191 = _nw53
			var _315_v192 D7
			_ = _315_v192
			_315_v192 = Companion_D7_.Create_DC17_(_312_v191, _51_v3, _50_v2)
			var _316_v193 _dafny.Sequence
			_ = _316_v193
			_316_v193 = _dafny.SeqOf(_315_v192)
			var _317_v194 D12
			_ = _317_v194
			_317_v194 = Companion_D12_.Create_DC30_(_310_v189, _51_v3, Companion_D11_.Create_DC25_(_311_v190), _316_v193, false)
			var _318_v195 _dafny.Array
			_ = _318_v195
			var _nwElement0_10 _dafny.Map = _307_v186
			_ = _nwElement0_10
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(14))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_10, 0)
			(_nw54).ArraySet1((func() _dafny.Map {
				if (_309_v188).Contains(_308_v187) {
					return (_309_v188).Get(_308_v187).(_dafny.Map)
				}
				return _307_v186
			})(), 1)
			(_nw54).ArraySet1((func() _dafny.Map {
				if _51_v3 {
					return _307_v186
				}
				return Companion_Default___.Fm46(_56_globalState)
			})(), 2)
			(_nw54).ArraySet1((_307_v186).Update((_317_v194).Dtor_cf59(), _57_v7), 3)
			(_nw54).ArraySet1((_307_v186).Merge(_307_v186), 4)
			(_nw54).ArraySet1(_307_v186, 5)
			(_nw54).ArraySet1(_307_v186, 6)
			(_nw54).ArraySet1(_307_v186, 7)
			(_nw54).ArraySet1(_307_v186, 8)
			(_nw54).ArraySet1((_307_v186).Update(_51_v3, _57_v7), 9)
			(_nw54).ArraySet1(_307_v186, 10)
			(_nw54).ArraySet1((_307_v186).Merge(_307_v186), 11)
			(_nw54).ArraySet1((func() _dafny.Map {
				if _51_v3 {
					return _307_v186
				}
				return _307_v186
			})(), 12)
			(_nw54).ArraySet1(_307_v186, 13)
			_318_v195 = _nw54
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_318_v195), 0))
			_ = _index37
			(_318_v195).ArraySet1(_307_v186, (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_318_v195), 0))
			_ = _index38
			(_318_v195).ArraySet1((_307_v186).Merge(_307_v186), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_54_v6), 0))
			_ = _index39
			(_54_v6).ArraySet1CodePoint(_dafny.CodePoint('r'), (_index39).Int())
			var _319_v196 _dafny.Sequence
			_ = _319_v196
			_319_v196 = _dafny.SeqOf(_dafny.IntOfInt64(796))
			var _320_v197 _dafny.Sequence
			_ = _320_v197
			_320_v197 = _dafny.SeqOf(_319_v196, _dafny.SeqOf(_50_v2))
			(_56_globalState).F14 = (_dafny.Zero).Minus(Companion_Default___.Fm7(_50_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32(((_320_v197).Select((Companion_Default___.SafeIndex((_53_v5).Cardinality(), _dafny.IntOfUint32((_320_v197).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), Companion_D0_.Create_DC0_(!(Companion_Default___.Fm13(_56_globalState))), _56_globalState))
		}
	}
	var _321_v198 _dafny.Array
	_ = _321_v198
	var _len0_12 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_12
	var _nw55 _dafny.Array
	_ = _nw55
	if _len0_12.Cmp(_dafny.Zero) == 0 {
		_nw55 = _dafny.NewArray(_len0_12)
	} else {
		var _init12 func(_dafny.Int) bool = (func(_322_v3 bool) func(_dafny.Int) bool {
			return func(_323_i22 _dafny.Int) bool {
				return _322_v3
			}
		})(_51_v3)
		_ = _init12
		var _element0_12 = _init12(_dafny.Zero)
		_ = _element0_12
		_nw55 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
		(_nw55).ArraySet1(_element0_12, 0)
		var _nativeLen0_12 = (_len0_12).Int()
		_ = _nativeLen0_12
		for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
			(_nw55).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
		}
	}
	_321_v198 = _nw55
	var _324_v199 *C1
	_ = _324_v199
	var _nw56 *C1 = New_C1_()
	_ = _nw56
	_nw56.Ctor__()
	_324_v199 = _nw56
	var _325_v200 _dafny.Map
	_ = _325_v200
	_325_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v199, _324_v199)
	var _326_v201 _dafny.Sequence
	_ = _326_v201
	_326_v201 = _dafny.SeqOf((_325_v200).Cardinality())
	var _327_v202 *C2
	_ = _327_v202
	var _nw57 *C2 = New_C2_()
	_ = _nw57
	_nw57.Ctor__(_dafny.SetOf(_321_v198), !((_50_v2).Cmp(_dafny.IntOfUint32((_326_v201).Cardinality())) != 0))
	_327_v202 = _nw57
	var _328_v203 D3
	_ = _328_v203
	_328_v203 = Companion_D3_.Create_DC9_()
	var _329_v204 _dafny.Sequence
	_ = _329_v204
	_329_v204 = _dafny.SeqOf(_328_v203)
	var _330_v205 _dafny.Sequence
	_ = _330_v205
	_330_v205 = _329_v204
	var _331_i23 _dafny.Int
	_ = _331_i23
	_331_i23 = _dafny.Zero
	{
		var _pat_let_tv6 = _51_v3
		_ = _pat_let_tv6
		for func(_source3 _dafny.Sequence) bool {
			var _347___mcc_h10 _dafny.Sequence = _source3
			_ = _347___mcc_h10
			var _348_cf21 _dafny.Sequence = _347___mcc_h10
			_ = _348_cf21
			return _pat_let_tv6
		}(_330_v205) {
			{
				if (_331_i23).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_331_i23 = (_331_i23).Plus(_dafny.One)
				var _332_v206 _dafny.Set
				_ = _332_v206
				_332_v206 = _dafny.SetOf(_48_v0)
				var _333_v207 _dafny.Map
				_ = _333_v207
				_333_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v3, (_332_v206).IsProperSubsetOf(_332_v206))
				var _rhs32 _dafny.Array = _321_v198
				_ = _rhs32
				var _rhs33 _dafny.Int = _50_v2
				_ = _rhs33
				var _rhs34 _dafny.Map = ((_333_v207).Merge(_333_v207)).Update((_51_v3) && (false), _51_v3)
				_ = _rhs34
				var _lhs25 *GlobalState = _56_globalState
				_ = _lhs25
				_321_v198 = _rhs32
				_lhs25.F14 = _rhs33
				_333_v207 = _rhs34
				var _334_v208 _dafny.Sequence
				_ = _334_v208
				var _335_v209 D0
				_ = _335_v209
				var _336_v210 bool
				_ = _336_v210
				var _337_v211 bool
				_ = _337_v211
				var _out13 _dafny.Sequence
				_ = _out13
				var _out14 D0
				_ = _out14
				var _out15 bool
				_ = _out15
				var _out16 bool
				_ = _out16
				_out13, _out14, _out15, _out16 = (_324_v199).M5((_dafny.Zero).Minus(_50_v2), _56_globalState)
				_334_v208 = _out13
				_335_v209 = _out14
				_336_v210 = _out15
				_337_v211 = _out16
				var _338_v212 _dafny.MultiSet
				_ = _338_v212
				_338_v212 = _dafny.MultiSetOf(_336_v210)
				_338_v212 = _338_v212
				var _339_v215 D10
				_ = _339_v215
				_339_v215 = Companion_D10_.Create_DC22_(func() _dafny.Set {
					var _coll25 = _dafny.NewBuilder()
					_ = _coll25
					for _iter25 := _dafny.Iterate((_326_v201).Elements()); ; {
						_compr_25, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _340_v214 _dafny.Int
						_340_v214 = interface{}(_compr_25).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_326_v201, _340_v214) {
							_coll25.Add(Companion_Default___.SafeModuloInt(_340_v214, _dafny.IntOfInt64(238)))
						}
					}
					return _coll25.ToSet()
				}())
				var _341_v216 _dafny.Map
				_ = _341_v216
				_341_v216 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_59_v9).Cardinality(), _339_v215)
				var _342_v217 _dafny.Sequence
				_ = _342_v217
				var _343_v218 D0
				_ = _343_v218
				var _344_v219 bool
				_ = _344_v219
				var _345_v220 bool
				_ = _345_v220
				var _out17 _dafny.Sequence
				_ = _out17
				var _out18 D0
				_ = _out18
				var _out19 bool
				_ = _out19
				var _out20 bool
				_ = _out20
				_out17, _out18, _out19, _out20 = (_324_v199).M5(((func() _dafny.Map {
					var _coll26 = _dafny.NewMapBuilder()
					_ = _coll26
					for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-608), _dafny.IntOfInt64(634))); ; {
						_compr_26, _ok26 := _iter26()
						if !_ok26 {
							break
						}
						var _346_v213 _dafny.Int
						_346_v213 = interface{}(_compr_26).(_dafny.Int)
						if ((_dafny.IntOfInt64(-608)).Cmp(_346_v213) <= 0) && ((_346_v213).Cmp(_dafny.IntOfInt64(634)) < 0) {
							_coll26.Add(Companion_Default___.SafeModuloInt(_346_v213, _50_v2), Companion_D10_.Create_DC22_(_57_v7))
						}
					}
					return _coll26.ToMap()
				}()).Merge(_341_v216)).Cardinality(), _56_globalState)
				_342_v217 = _out17
				_343_v218 = _out18
				_344_v219 = _out19
				_345_v220 = _out20
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _349_v221 _dafny.Sequence
	_ = _349_v221
	var _350_v222 D0
	_ = _350_v222
	var _351_v223 bool
	_ = _351_v223
	var _352_v224 bool
	_ = _352_v224
	var _out21 _dafny.Sequence
	_ = _out21
	var _out22 D0
	_ = _out22
	var _out23 bool
	_ = _out23
	var _out24 bool
	_ = _out24
	_out21, _out22, _out23, _out24 = (_324_v199).M5((_50_v2).Minus(_50_v2), _56_globalState)
	_349_v221 = _out21
	_350_v222 = _out22
	_351_v223 = _out23
	_352_v224 = _out24
	var _353_v225 _dafny.Sequence
	_ = _353_v225
	var _354_v226 D0
	_ = _354_v226
	var _355_v227 bool
	_ = _355_v227
	var _356_v228 bool
	_ = _356_v228
	var _out25 _dafny.Sequence
	_ = _out25
	var _out26 D0
	_ = _out26
	var _out27 bool
	_ = _out27
	var _out28 bool
	_ = _out28
	_out25, _out26, _out27, _out28 = (_324_v199).M5(_dafny.IntOfInt64(832), _56_globalState)
	_353_v225 = _out25
	_354_v226 = _out26
	_355_v227 = _out27
	_356_v228 = _out28
	_351_v223 = (func() bool {
		if (_50_v2).Cmp((func() _dafny.Int {
			if (_59_v9).Contains(_50_v2) {
				return (_59_v9).Get(_50_v2).(_dafny.Int)
			}
			return _50_v2
		})()) <= 0 {
			return _352_v224
		}
		return !(_356_v228)
	})()
	(_56_globalState).F3 = _351_v223
	var _357_v229 _dafny.Sequence
	_ = _357_v229
	var _358_v230 D0
	_ = _358_v230
	var _359_v231 bool
	_ = _359_v231
	var _360_v232 bool
	_ = _360_v232
	var _out29 _dafny.Sequence
	_ = _out29
	var _out30 D0
	_ = _out30
	var _out31 bool
	_ = _out31
	var _out32 bool
	_ = _out32
	_out29, _out30, _out31, _out32 = (_324_v199).M5((_dafny.IntOfUint32((_60_v10).Cardinality())).Times(_50_v2), _56_globalState)
	_357_v229 = _out29
	_358_v230 = _out30
	_359_v231 = _out31
	_360_v232 = _out32
	for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_321_v198), 0))); ; {
		_guard_loop_0, _ok27 := _iter27()
		if !_ok27 {
			break
		}
		var _361_i24 _dafny.Int
		_361_i24 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_361_i24).Sign() != -1) && ((_361_i24).Cmp(_dafny.ArrayLen((_321_v198), 0)) < 0)) {
			(_321_v198).ArraySet1((func() bool {
				if _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("vws"), _61_v11) {
					return !(!(_359_v231))
				}
				return _359_v231
			})(), (_361_i24).Int())
		}
	}
	var _362_v233 _dafny.MultiSet
	_ = _362_v233
	_362_v233 = _dafny.MultiSetOf(_359_v231, _51_v3)
	var _363_i25 _dafny.Int
	_ = _363_i25
	_363_i25 = _dafny.Zero
	{
		for ((func() _dafny.Int {
			if (_362_v233).Contains(_352_v224) {
				return (_362_v233).Multiplicity(_352_v224)
			}
			return _dafny.IntOfUint32((_52_v4).Cardinality())
		})()).Cmp(_50_v2) == 0 {
			{
				if (_363_i25).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_363_i25 = (_363_i25).Plus(_dafny.One)
				var _364_v234 _dafny.Set
				_ = _364_v234
				_364_v234 = _dafny.SetOf(_321_v198)
				var _365_v235 _dafny.Map
				_ = _365_v235
				_365_v235 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v11, false)
				var _366_v236 T2
				_ = _366_v236
				var _nw58 *C8 = New_C8_()
				_ = _nw58
				_nw58.Ctor__(_dafny.IntOfUint32((_60_v10).Cardinality()), _355_v227, _321_v198, _364_v234, _351_v223, _50_v2, _365_v235)
				_366_v236 = _nw58
				var _367_v237 _dafny.Array
				_ = _367_v237
				var _nwElement0_11 T2 = _366_v236
				_ = _nwElement0_11
				var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(3))
				_ = _nw59
				(_nw59).ArraySet1(_nwElement0_11, 0)
				(_nw59).ArraySet1(_366_v236, 1)
				(_nw59).ArraySet1(_366_v236, 2)
				_367_v237 = _nw59
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_367_v237), 0))
				_ = _index40
				(_367_v237).ArraySet1(_366_v236, (_index40).Int())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_321_v198), 0))
				_ = _index41
				(_321_v198).ArraySet1(_51_v3, (_index41).Int())
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_367_v237), 0))
				_ = _index42
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_321_v198), 0))
				_ = _index43
				var _rhs35 T2 = _366_v236
				_ = _rhs35
				var _rhs36 _dafny.Int = (_366_v236).F20()
				_ = _rhs36
				var _rhs37 bool = _355_v227
				_ = _rhs37
				var _rhs38 bool = !(_352_v224)
				_ = _rhs38
				var _rhs39 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(915), _50_v2)
				_ = _rhs39
				var _lhs26 _dafny.Array = _367_v237
				_ = _lhs26
				var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_367_v237), 0))
				_ = _lhs27
				var _lhs28 *GlobalState = _56_globalState
				_ = _lhs28
				var _lhs29 _dafny.Array = _321_v198
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_321_v198), 0))
				_ = _lhs30
				var _lhs31 *GlobalState = _56_globalState
				_ = _lhs31
				(_lhs26).ArraySet1(_rhs35, (_lhs27).Int())
				_lhs28.F14 = _rhs36
				(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
				_lhs31.F3 = _rhs38
				_50_v2 = _rhs39
				var _368_v238 *C3
				_ = _368_v238
				var _nw60 *C3 = New_C3_()
				_ = _nw60
				_nw60.Ctor__(_366_v236.F19(), (_364_v234).Union(_366_v236.F17()), _351_v223)
				_368_v238 = _nw60
				var _369_v239 _dafny.Array
				_ = _369_v239
				var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
				_ = _nw61
				_369_v239 = _nw61
				var _370_v240 _dafny.Set
				_ = _370_v240
				_370_v240 = _dafny.SetOf(_48_v0)
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_369_v239), 0))
				_ = _index44
				(_369_v239).ArraySet1(_370_v240, (_index44).Int())
				var _371_v241 T1
				_ = _371_v241
				var _nw62 *C3 = New_C3_()
				_ = _nw62
				_nw62.Ctor__(_366_v236.F19(), _366_v236.F17(), (_366_v236).F18())
				_371_v241 = _nw62
				var _372_v242 _dafny.Map
				_ = _372_v242
				_372_v242 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, _371_v241)
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_369_v239), 0))
				_ = _index45
				var _rhs40 _dafny.Set = ((_370_v240).Intersection(_370_v240)).Difference(_370_v240)
				_ = _rhs40
				var _rhs41 _dafny.Int = (_372_v242).Cardinality()
				_ = _rhs41
				var _lhs32 _dafny.Array = _369_v239
				_ = _lhs32
				var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_369_v239), 0))
				_ = _lhs33
				(_lhs32).ArraySet1(_rhs40, (_lhs33).Int())
				_50_v2 = _rhs41
				var _373_v243 _dafny.Map
				_ = _373_v243
				_373_v243 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC5_(_362_v233), true)
				var _374_v244 D2
				_ = _374_v244
				_374_v244 = Companion_D2_.Create_DC5_(_362_v233)
				var _375_v246 _dafny.Sequence
				_ = _375_v246
				_375_v246 = _dafny.SeqOf(_374_v244)
				var _376_v248 _dafny.Sequence
				_ = _376_v248
				_376_v248 = _dafny.SeqOf(_373_v243)
				var _377_v252 _dafny.MultiSet
				_ = _377_v252
				_377_v252 = _dafny.MultiSetOf(_374_v244)
				var _378_v253 _dafny.Array
				_ = _378_v253
				var _nwElement0_12 _dafny.Map = _373_v243
				_ = _nwElement0_12
				var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(24))
				_ = _nw63
				(_nw63).ArraySet1(_nwElement0_12, 0)
				(_nw63).ArraySet1(Companion_Default___.Fm47(_356_v228, _56_globalState), 1)
				(_nw63).ArraySet1((Companion_Default___.Fm47(_51_v3, _56_globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v244, _355_v227)), 2)
				(_nw63).ArraySet1((func() _dafny.Map {
					if (_366_v236).F18() {
						return _373_v243
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v244, false)
				})(), 3)
				(_nw63).ArraySet1(_373_v243, 4)
				(_nw63).ArraySet1((_373_v243).Merge(_373_v243), 5)
				(_nw63).ArraySet1(func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter28 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_375_v246, (Companion_Default___.SafeIndex((_366_v236).F20(), _dafny.IntOfUint32((_375_v246).Cardinality()))).Uint32(), Companion_D2_.Create_DC5_(_dafny.MultiSetOf(!(_51_v3))))).Elements()); ; {
						_compr_27, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _379_v245 D2
						_379_v245 = interface{}(_compr_27).(D2)
						if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_375_v246, (Companion_Default___.SafeIndex((_366_v236).F20(), _dafny.IntOfUint32((_375_v246).Cardinality()))).Uint32(), Companion_D2_.Create_DC5_(_dafny.MultiSetOf(!(_51_v3)))), _379_v245) {
							_coll27.Add(_379_v245, _352_v224)
						}
					}
					return _coll27.ToMap()
				}(), 6)
				(_nw63).ArraySet1((func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter29 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(830))).Uint32(), func(coer28 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}((func(_380_v244 D2) func(_dafny.Int) D2 {
						return func(_381_i26 _dafny.Int) D2 {
							return _380_v244
						}
					})(_374_v244)))).Elements()); ; {
						_compr_28, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _382_v247 D2
						_382_v247 = interface{}(_compr_28).(D2)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(830))).Uint32(), func(coer29 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
							return func(arg29 _dafny.Int) interface{} {
								return coer29(arg29)
							}
						}((func(_383_v244 D2) func(_dafny.Int) D2 {
							return func(_381_i26 _dafny.Int) D2 {
								return _383_v244
							}
						})(_374_v244))), _382_v247) {
							_coll28.Add(_382_v247, (_366_v236).F18())
						}
					}
					return _coll28.ToMap()
				}()).Update(_374_v244, _51_v3), 7)
				(_nw63).ArraySet1(_373_v243, 8)
				(_nw63).ArraySet1((_373_v243).Merge(_373_v243), 9)
				(_nw63).ArraySet1(_373_v243, 10)
				(_nw63).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v244, _51_v3), 11)
				(_nw63).ArraySet1((_373_v243).Merge((_376_v248).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.IntOfUint32((_376_v248).Cardinality()))).Uint32()).(_dafny.Map)), 12)
				(_nw63).ArraySet1(func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate((_375_v246).Elements()); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _384_v249 D2
						_384_v249 = interface{}(_compr_29).(D2)
						if _dafny.Companion_Sequence_.Contains(_375_v246, _384_v249) {
							_coll29.Add(_384_v249, (Companion_D5_.Create_DC14_(false, (_dafny.SetOf((_366_v236).F20(), (_366_v236).F20())).Cardinality(), (_366_v236).F20())).Dtor_cf26())
						}
					}
					return _coll29.ToMap()
				}(), 13)
				(_nw63).ArraySet1(func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter31 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_375_v246, (Companion_Default___.SafeIndex(_50_v2, _dafny.IntOfUint32((_375_v246).Cardinality()))).Uint32(), _374_v244)).Elements()); ; {
						_compr_30, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _385_v250 D2
						_385_v250 = interface{}(_compr_30).(D2)
						if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_375_v246, (Companion_Default___.SafeIndex(_50_v2, _dafny.IntOfUint32((_375_v246).Cardinality()))).Uint32(), _374_v244), _385_v250) {
							_coll30.Add(_385_v250, _360_v232)
						}
					}
					return _coll30.ToMap()
				}(), 14)
				(_nw63).ArraySet1((_373_v243).Merge(_373_v243), 15)
				(_nw63).ArraySet1(_373_v243, 16)
				(_nw63).ArraySet1(Companion_Default___.Fm47(_51_v3, _56_globalState), 17)
				(_nw63).ArraySet1(_373_v243, 18)
				(_nw63).ArraySet1(func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter32 := _dafny.Iterate((_377_v252).Elements()); ; {
						_compr_31, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _386_v251 D2
						_386_v251 = interface{}(_compr_31).(D2)
						if (_377_v252).Contains(_386_v251) {
							_coll31.Add(_386_v251, (_366_v236).F18())
						}
					}
					return _coll31.ToMap()
				}(), 19)
				(_nw63).ArraySet1(_373_v243, 20)
				(_nw63).ArraySet1(Companion_Default___.Fm47(_352_v224, _56_globalState), 21)
				(_nw63).ArraySet1(_373_v243, 22)
				(_nw63).ArraySet1(_373_v243, 23)
				_378_v253 = _nw63
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_378_v253), 0))
				_ = _index46
				(_378_v253).ArraySet1(func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter33 := _dafny.Iterate((_377_v252).Elements()); ; {
						_compr_32, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _387_v254 D2
						_387_v254 = interface{}(_compr_32).(D2)
						if (_377_v252).Contains(_387_v254) {
							_coll32.Add(_387_v254, (_371_v241).F18())
						}
					}
					return _coll32.ToMap()
				}(), (_index46).Int())
				var _388_v255 _dafny.Array
				_ = _388_v255
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_13
				var _nw64 _dafny.Array
				_ = _nw64
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw64 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_389_v201 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_390_i27 _dafny.Int) _dafny.Sequence {
							return _389_v201
						}
					})(_326_v201)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw64 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw64).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw64).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_388_v255 = _nw64
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_388_v255), 0))
				_ = _index47
				(_388_v255).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_326_v201, _326_v201), (_index47).Int())
				var _391_v256 _dafny.Map
				_ = _391_v256
				_391_v256 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_366_v236).F20(), _359_v231)
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_378_v253), 0))
				_ = _index48
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_388_v255), 0))
				_ = _index49
				var _rhs42 _dafny.Sequence = _349_v221
				_ = _rhs42
				var _rhs43 _dafny.Int = (_366_v236).F20()
				_ = _rhs43
				var _rhs44 _dafny.Map = ((func() _dafny.Map {
					if false {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v244, (func() bool {
							if (_391_v256).Contains(_dafny.IntOfInt64(-278)) {
								return (_391_v256).Get(_dafny.IntOfInt64(-278)).(bool)
							}
							return _355_v227
						})())
					}
					return Companion_Default___.Fm47(false, _56_globalState)
				})()).Merge(_373_v243)
				_ = _rhs44
				var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_326_v201, (Companion_Default___.SafeIndex((_366_v236).F20(), _dafny.IntOfUint32((_326_v201).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_50_v2))
				_ = _rhs45
				var _lhs34 *GlobalState = _56_globalState
				_ = _lhs34
				var _lhs35 _dafny.Array = _378_v253
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_378_v253), 0))
				_ = _lhs36
				var _lhs37 _dafny.Array = _388_v255
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_388_v255), 0))
				_ = _lhs38
				_353_v225 = _rhs42
				_lhs34.F14 = _rhs43
				(_lhs35).ArraySet1(_rhs44, (_lhs36).Int())
				(_lhs37).ArraySet1(_rhs45, (_lhs38).Int())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	(_56_globalState).F3 = _360_v232
	_dafny.Print((_49_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_50_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_51_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_52_v4, _dafny.SeqOf(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_v5).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-548), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_56_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_56_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_56_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_56_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F15).Equals(_dafny.MultiSetOf(_dafny.Zero, _dafny.Zero, _dafny.IntOfInt64(3288), _dafny.IntOfInt64(3288))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_globalState.F16).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v7).Equals(_dafny.SetOf(_dafny.IntOfInt64(548))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_58_v8, _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(548)), _dafny.SetOf(_dafny.IntOfInt64(548)), _dafny.SetOf(_dafny.IntOfInt64(548)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_60_v10.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_61_v11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_62_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.One), _dafny.UnicodeSeqOfUtf8Bytes("yvxdaq"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_i12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_321_v198).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_325_v200).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_326_v201, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_329_v204, _dafny.SeqOf(Companion_D3_.Create_DC9_())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_330_v205), _dafny.SeqOf(Companion_D3_.Create_DC9_())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_331_i23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_349_v221.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_350_v222).Dtor_cf1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_350_v222).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_350_v222).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_350_v222).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_351_v223)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_352_v224)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_353_v225.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_354_v226).Dtor_cf1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_354_v226).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_354_v226).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_354_v226).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_355_v227)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_356_v228)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_357_v229.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_358_v230).Dtor_cf1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_358_v230).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_358_v230).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_358_v230).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_359_v231)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_360_v232)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_362_v233).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_363_i25)
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
	Cf1 _dafny.Sequence
	Cf2 _dafny.Int
	Cf3 _dafny.Int
	Cf4 _dafny.CodePoint
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Sequence, Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 _dafny.CodePoint) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 bool
	Cf6 _dafny.Int
	Cf7 _dafny.Array
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 bool, Cf6 _dafny.Int, Cf7 _dafny.Array) D0 {
	return D0{D0_DC2{Cf5, Cf6, Cf7}}
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

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() bool {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D0_DC2).Cf7
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
			return "D0.DC1" + "(" + data.Cf1.VerbatimString(true) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1.Equals(data2.Cf1) && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7
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
	Cf9  bool
	Cf10 _dafny.Set
	Cf11 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf9 bool, Cf10 _dafny.Set, Cf11 _dafny.Int) D1 {
	return D1{D1_DC4{Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf8 _dafny.MultiSet
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf8 _dafny.MultiSet) D1 {
	return D1{D1_DC3{Cf8}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, _dafny.EmptySet, _dafny.Zero)
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Set {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf8() _dafny.MultiSet {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10.Equals(data2.Cf10) && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf13 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf13 _dafny.Int) D2 {
	return D2{D2_DC6{Cf13}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf12 _dafny.MultiSet
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf12 _dafny.MultiSet) D2 {
	return D2{D2_DC5{Cf12}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC7 struct {
	Cf14 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf14 D2) D2 {
	return D2{D2_DC7{Cf14}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero)
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) Dtor_cf12() _dafny.MultiSet {
	return _this.Get_().(D2_DC5).Cf12
}

func (_this D2) Dtor_cf14() D2 {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D3_DC10 struct {
	Cf16 _dafny.Sequence
	Cf17 bool
	Cf18 _dafny.Sequence
	Cf19 _dafny.Sequence
	Cf20 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 _dafny.Sequence, Cf17 bool, Cf18 _dafny.Sequence, Cf19 _dafny.Sequence, Cf20 _dafny.Int) D3 {
	return D3{D3_DC10{Cf16, Cf17, Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf15 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf15 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf15}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_()
}

func (_this D3) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() bool {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf20
}

func (_this D3) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + data.Cf16.VerbatimString(true) + ", " + _dafny.String(data.Cf17) + ", " + data.Cf18.VerbatimString(true) + ", " + data.Cf19.VerbatimString(true) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf15) + ")"
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
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16) && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18) && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf21 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf21 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf21}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf21
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 bool
	Cf24 _dafny.Int
	Cf25 _dafny.Sequence
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf23 bool, Cf24 _dafny.Int, Cf25 _dafny.Sequence) D5 {
	return D5{D5_DC13{Cf23, Cf24, Cf25}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf26 bool
	Cf27 _dafny.Int
	Cf28 _dafny.Int
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf26 bool, Cf27 _dafny.Int, Cf28 _dafny.Int) D5 {
	return D5{D5_DC14{Cf26, Cf27, Cf28}}
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

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(false, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC13).Cf23
}

func (_this D5) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf24
}

func (_this D5) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D5_DC13).Cf25
}

func (_this D5) Dtor_cf26() bool {
	return _this.Get_().(D5_DC14).Cf26
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf28
}

func (_this D5) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D5_DC12).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Equals(data2.Cf25)
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28.Cmp(data2.Cf28) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf22.Equals(data2.Cf22)
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
	Cf29 _dafny.Sequence
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf29 _dafny.Sequence) D6 {
	return D6{D6_DC15{Cf29}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D6) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D6_DC15).Cf29
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf29) + ")"
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

type D7_DC17 struct {
	Cf31 _dafny.Array
	Cf32 bool
	Cf33 _dafny.Int
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf31 _dafny.Array, Cf32 bool, Cf33 _dafny.Int) D7 {
	return D7{D7_DC17{Cf31, Cf32, Cf33}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC16 struct {
	Cf30 _dafny.Set
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf30 _dafny.Set) D7 {
	return D7{D7_DC16{Cf30}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC17_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero)
}

func (_this D7) Dtor_cf31() _dafny.Array {
	return _this.Get_().(D7_DC17).Cf31
}

func (_this D7) Dtor_cf32() bool {
	return _this.Get_().(D7_DC17).Cf32
}

func (_this D7) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D7_DC17).Cf33
}

func (_this D7) Dtor_cf30() _dafny.Set {
	return _this.Get_().(D7_DC16).Cf30
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf30) + ")"
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
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf34 _dafny.Sequence
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf34 _dafny.Sequence) D8 {
	return D8{D8_DC18{Cf34}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D8_DC18).Cf34
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf34) + ")"
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

type D9_DC20 struct {
	Cf36 _dafny.Map
	Cf37 _dafny.Int
	Cf38 bool
	Cf39 _dafny.Map
	Cf40 _dafny.CodePoint
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf36 _dafny.Map, Cf37 _dafny.Int, Cf38 bool, Cf39 _dafny.Map, Cf40 _dafny.CodePoint) D9 {
	return D9{D9_DC20{Cf36, Cf37, Cf38, Cf39, Cf40}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC19 struct {
	Cf35 _dafny.Map
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf35 _dafny.Map) D9 {
	return D9{D9_DC19{Cf35}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC21 struct {
	Cf41 D9
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf41 D9) D9 {
	return D9{D9_DC21{Cf41}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC20_(_dafny.EmptyMap, _dafny.Zero, false, _dafny.EmptyMap, _dafny.CodePoint('D'))
}

func (_this D9) Dtor_cf36() _dafny.Map {
	return _this.Get_().(D9_DC20).Cf36
}

func (_this D9) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D9_DC20).Cf37
}

func (_this D9) Dtor_cf38() bool {
	return _this.Get_().(D9_DC20).Cf38
}

func (_this D9) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D9_DC20).Cf39
}

func (_this D9) Dtor_cf40() _dafny.CodePoint {
	return _this.Get_().(D9_DC20).Cf40
}

func (_this D9) Dtor_cf35() _dafny.Map {
	return _this.Get_().(D9_DC19).Cf35
}

func (_this D9) Dtor_cf41() D9 {
	return _this.Get_().(D9_DC21).Cf41
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf36.Equals(data2.Cf36) && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38 && data1.Cf39.Equals(data2.Cf39) && data1.Cf40 == data2.Cf40
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D10_DC23 struct {
	Cf43 bool
	Cf44 bool
	Cf45 T2
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf43 bool, Cf44 bool, Cf45 T2) D10 {
	return D10{D10_DC23{Cf43, Cf44, Cf45}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

type D10_DC22 struct {
	Cf42 _dafny.Set
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf42 _dafny.Set) D10 {
	return D10{D10_DC22{Cf42}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC24 struct {
	Cf46 D10
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf46 D10) D10 {
	return D10{D10_DC24{Cf46}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC23_(false, false, (T2)(nil))
}

func (_this D10) Dtor_cf43() bool {
	return _this.Get_().(D10_DC23).Cf43
}

func (_this D10) Dtor_cf44() bool {
	return _this.Get_().(D10_DC23).Cf44
}

func (_this D10) Dtor_cf45() T2 {
	return _this.Get_().(D10_DC23).Cf45
}

func (_this D10) Dtor_cf42() _dafny.Set {
	return _this.Get_().(D10_DC22).Cf42
}

func (_this D10) Dtor_cf46() D10 {
	return _this.Get_().(D10_DC24).Cf46
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D10_DC24:
		{
			return "D10.DC24" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf43 == data2.Cf43 && data1.Cf44 == data2.Cf44 && _dafny.AreEqual(data1.Cf45, data2.Cf45)
		}
	case D10_DC22:
		{
			data2, ok := other.Get_().(D10_DC22)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
			return ok && data1.Cf46.Equals(data2.Cf46)
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
	Cf48 bool
	Cf49 _dafny.Int
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf48 bool, Cf49 _dafny.Int) D11 {
	return D11{D11_DC26{Cf48, Cf49}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC27 struct {
	Cf50 _dafny.Sequence
	Cf51 _dafny.Int
	Cf52 _dafny.Int
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf50 _dafny.Sequence, Cf51 _dafny.Int, Cf52 _dafny.Int) D11 {
	return D11{D11_DC27{Cf50, Cf51, Cf52}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC25 struct {
	Cf47 _dafny.Set
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf47 _dafny.Set) D11 {
	return D11{D11_DC25{Cf47}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC28 struct {
	Cf53 D11
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf53 D11) D11 {
	return D11{D11_DC28{Cf53}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC26_(false, _dafny.Zero)
}

func (_this D11) Dtor_cf48() bool {
	return _this.Get_().(D11_DC26).Cf48
}

func (_this D11) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D11_DC26).Cf49
}

func (_this D11) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D11_DC27).Cf50
}

func (_this D11) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D11_DC27).Cf51
}

func (_this D11) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D11_DC27).Cf52
}

func (_this D11) Dtor_cf47() _dafny.Set {
	return _this.Get_().(D11_DC25).Cf47
}

func (_this D11) Dtor_cf53() D11 {
	return _this.Get_().(D11_DC28).Cf53
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf53) + ")"
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
			return ok && data1.Cf48 == data2.Cf48 && data1.Cf49.Cmp(data2.Cf49) == 0
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf50.Equals(data2.Cf50) && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf47.Equals(data2.Cf47)
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf53.Equals(data2.Cf53)
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

type D12_DC30 struct {
	Cf55 _dafny.MultiSet
	Cf56 bool
	Cf57 D11
	Cf58 _dafny.Sequence
	Cf59 bool
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf55 _dafny.MultiSet, Cf56 bool, Cf57 D11, Cf58 _dafny.Sequence, Cf59 bool) D12 {
	return D12{D12_DC30{Cf55, Cf56, Cf57, Cf58, Cf59}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

type D12_DC29 struct {
	Cf54 _dafny.Sequence
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf54 _dafny.Sequence) D12 {
	return D12{D12_DC29{Cf54}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC30_(_dafny.EmptyMultiSet, false, Companion_D11_.Default(), _dafny.EmptySeq, false)
}

func (_this D12) Dtor_cf55() _dafny.MultiSet {
	return _this.Get_().(D12_DC30).Cf55
}

func (_this D12) Dtor_cf56() bool {
	return _this.Get_().(D12_DC30).Cf56
}

func (_this D12) Dtor_cf57() D11 {
	return _this.Get_().(D12_DC30).Cf57
}

func (_this D12) Dtor_cf58() _dafny.Sequence {
	return _this.Get_().(D12_DC30).Cf58
}

func (_this D12) Dtor_cf59() bool {
	return _this.Get_().(D12_DC30).Cf59
}

func (_this D12) Dtor_cf54() _dafny.Sequence {
	return _this.Get_().(D12_DC29).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC30:
		{
			data2, ok := other.Get_().(D12_DC30)
			return ok && data1.Cf55.Equals(data2.Cf55) && data1.Cf56 == data2.Cf56 && data1.Cf57.Equals(data2.Cf57) && data1.Cf58.Equals(data2.Cf58) && data1.Cf59 == data2.Cf59
		}
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D13_DC32 struct {
	Cf61 bool
	Cf62 _dafny.Int
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf61 bool, Cf62 _dafny.Int) D13 {
	return D13{D13_DC32{Cf61, Cf62}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

type D13_DC31 struct {
	Cf60 _dafny.Map
}

func (D13_DC31) isD13() {}

func (CompanionStruct_D13_) Create_DC31_(Cf60 _dafny.Map) D13 {
	return D13{D13_DC31{Cf60}}
}

func (_this D13) Is_DC31() bool {
	_, ok := _this.Get_().(D13_DC31)
	return ok
}

type D13_DC33 struct {
	Cf63 D13
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf63 D13) D13 {
	return D13{D13_DC33{Cf63}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC32_(false, _dafny.Zero)
}

func (_this D13) Dtor_cf61() bool {
	return _this.Get_().(D13_DC32).Cf61
}

func (_this D13) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D13_DC32).Cf62
}

func (_this D13) Dtor_cf60() _dafny.Map {
	return _this.Get_().(D13_DC31).Cf60
}

func (_this D13) Dtor_cf63() D13 {
	return _this.Get_().(D13_DC33).Cf63
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC32:
		{
			return "D13.DC32" + "(" + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ")"
		}
	case D13_DC31:
		{
			return "D13.DC31" + "(" + _dafny.String(data.Cf60) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf63) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC32:
		{
			data2, ok := other.Get_().(D13_DC32)
			return ok && data1.Cf61 == data2.Cf61 && data1.Cf62.Cmp(data2.Cf62) == 0
		}
	case D13_DC31:
		{
			data2, ok := other.Get_().(D13_DC31)
			return ok && data1.Cf60.Equals(data2.Cf60)
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf63.Equals(data2.Cf63)
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
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_() D14 {
	return D14{D14_DC35{}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC34 struct {
	Cf64 _dafny.Array
}

func (D14_DC34) isD14() {}

func (CompanionStruct_D14_) Create_DC34_(Cf64 _dafny.Array) D14 {
	return D14{D14_DC34{Cf64}}
}

func (_this D14) Is_DC34() bool {
	_, ok := _this.Get_().(D14_DC34)
	return ok
}

type D14_DC36 struct {
	Cf65 D14
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf65 D14) D14 {
	return D14{D14_DC36{Cf65}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC35_()
}

func (_this D14) Dtor_cf64() _dafny.Array {
	return _this.Get_().(D14_DC34).Cf64
}

func (_this D14) Dtor_cf65() D14 {
	return _this.Get_().(D14_DC36).Cf65
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC35:
		{
			return "D14.DC35"
		}
	case D14_DC34:
		{
			return "D14.DC34" + "(" + _dafny.String(data.Cf64) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf65) + ")"
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
			_, ok := other.Get_().(D14_DC35)
			return ok
		}
	case D14_DC34:
		{
			data2, ok := other.Get_().(D14_DC34)
			return ok && data1.Cf64 == data2.Cf64
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf65.Equals(data2.Cf65)
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
	F17() _dafny.Set
	F17_set_(value _dafny.Set)
	F18() bool
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
	F17() _dafny.Set
	F17_set_(value _dafny.Set)
	F18() bool
	F19() _dafny.Array
	F19_set_(value _dafny.Array)
	Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int)
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
	F17() _dafny.Set
	F17_set_(value _dafny.Set)
	F19() _dafny.Array
	F19_set_(value _dafny.Array)
	Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int)
	F18() bool
	F21() _dafny.Map
	F21_set_(value _dafny.Map)
	Fm1(p0 _dafny.Map, globalState *GlobalState) bool
	F20() _dafny.Int
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
	F3   bool
	F7   _dafny.Array
	F9   bool
	F14  _dafny.Int
	F15  _dafny.MultiSet
	F16  _dafny.Array
	_f0  bool
	_f1  bool
	_f2  bool
	_f4  _dafny.Int
	_f5  _dafny.Int
	_f6  _dafny.Int
	_f8  _dafny.Int
	_f10 _dafny.CodePoint
	_f11 bool
	_f12 _dafny.Int
	_f13 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = false
	_this.F7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F9 = false
	_this.F14 = _dafny.Zero
	_this.F15 = _dafny.EmptyMultiSet
	_this.F16 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = false
	_this._f1 = false
	_this._f2 = false
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f10 = _dafny.CodePoint('D')
	_this._f11 = false
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 bool, f3 bool, f4 _dafny.Int, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.Array, f8 _dafny.Int, f9 bool, f10 _dafny.CodePoint, f11 bool, f12 _dafny.Int, f13 _dafny.Array, f14 _dafny.Int, f15 _dafny.MultiSet, f16 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this).F15 = f15
		(_this).F16 = f16
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Int {
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
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F10() _dafny.CodePoint {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Array {
	{
		return _this._f13
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm5(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(583)).Cmp(_dafny.IntOfInt64(122)) < 0
	}
}
func (_this *C1) M5(p0 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, D0, bool, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _392_v0 _dafny.CodePoint
		_ = _392_v0
		_392_v0 = _dafny.CodePoint('n')
		var _393_v2 _dafny.Map
		_ = _393_v2
		_393_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v0, p0)
		var _394_v4 _dafny.Sequence
		_ = _394_v4
		_394_v4 = _dafny.SeqOf((func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter34 := _dafny.Iterate((_393_v2).Keys().Elements()); ; {
				_compr_33, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _395_v1 _dafny.CodePoint
				_395_v1 = interface{}(_compr_33).(_dafny.CodePoint)
				if (_393_v2).Contains(_395_v1) {
					_coll33.Add(_395_v1, (func() _dafny.Map {
						var _coll34 = _dafny.NewMapBuilder()
						_ = _coll34
						for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-472), _dafny.IntOfInt64(765))); ; {
							_compr_34, _ok35 := _iter35()
							if !_ok35 {
								break
							}
							var _396_v3 _dafny.Int
							_396_v3 = interface{}(_compr_34).(_dafny.Int)
							if ((_dafny.IntOfInt64(-472)).Cmp(_396_v3) <= 0) && ((_396_v3).Cmp(_dafny.IntOfInt64(765)) < 0) {
								_coll34.Add((_396_v3).Minus(p0), true)
							}
						}
						return _coll34.ToMap()
					}()).Cardinality())
				}
			}
			return _coll33.ToMap()
		}()).Contains((Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("ecidbhho"), p0, (_dafny.Zero).Minus(p0), _392_v0)).Dtor_cf4()))
		var _397_i0 _dafny.Int
		_ = _397_i0
		_397_i0 = _dafny.Zero
		{
			for (_394_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_394_v4).Cardinality()))).Uint32()).(bool) {
				{
					if (_397_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_397_i0 = (_397_i0).Plus(_dafny.One)
					var _398_v5 _dafny.Array
					_ = _398_v5
					var _nw65 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
					_ = _nw65
					_398_v5 = _nw65
					var _399_v6 bool
					_ = _399_v6
					_399_v6 = true
					var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))
					_ = _index50
					(_398_v5).ArraySet1(_399_v6, (_index50).Int())
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))
					_ = _index51
					(_398_v5).ArraySet1(_399_v6, (_index51).Int())
					if (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool) {
						var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))
						_ = _index52
						var _rhs46 _dafny.Int = p0
						_ = _rhs46
						var _rhs47 bool = (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool)
						_ = _rhs47
						var _rhs48 bool = _399_v6
						_ = _rhs48
						var _rhs49 _dafny.Array = _398_v5
						_ = _rhs49
						var _lhs39 *GlobalState = globalState
						_ = _lhs39
						var _lhs40 _dafny.Array = _398_v5
						_ = _lhs40
						var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))
						_ = _lhs41
						_lhs39.F14 = _rhs46
						r2 = _rhs47
						(_lhs40).ArraySet1(_rhs48, (_lhs41).Int())
						_398_v5 = _rhs49
						var _400_v7 _dafny.Sequence
						_ = _400_v7
						_400_v7 = _dafny.SeqOf(p0)
						(globalState).F9 = (_this).Fm5(p0, _dafny.Companion_Sequence_.Concatenate(_400_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg30 _dafny.Int) interface{} {
								return coer30(arg30)
							}
						}(func(_401_i1 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(-71)
						}))), globalState)
						var _402_v8 _dafny.MultiSet
						_ = _402_v8
						_402_v8 = _dafny.MultiSetOf(p0)
						var _403_v9 D1
						_ = _403_v9
						_403_v9 = Companion_D1_.Create_DC3_(_402_v8)
						(globalState).F14 = ((_403_v9).Dtor_cf8()).Cardinality()
						(globalState).F14 = p0
						_399_v6 = _399_v6
					} else {
						var _404_v10 _dafny.Sequence
						_ = _404_v10
						_404_v10 = _dafny.UnicodeSeqOfUtf8Bytes("lylpu")
						var _405_v11 _dafny.Map
						_ = _405_v11
						_405_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_404_v10, _404_v10)
						var _406_v12 _dafny.MultiSet
						_ = _406_v12
						_406_v12 = _dafny.MultiSetOf(p0, p0)
						var _407_v13 _dafny.Array
						_ = _407_v13
						var _nwElement0_13 _dafny.Int = (func() _dafny.Int {
							if (_406_v12).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kx")).Cardinality())) {
								return (_406_v12).Multiplicity(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kx")).Cardinality()))
							}
							return p0
						})()
						_ = _nwElement0_13
						var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(2))
						_ = _nw66
						(_nw66).ArraySet1(_nwElement0_13, 0)
						(_nw66).ArraySet1(p0, 1)
						_407_v13 = _nw66
						var _408_v14 D0
						_ = _408_v14
						_408_v14 = Companion_D0_.Create_DC2_((_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool), (_405_v11).Cardinality(), _407_v13)
						var _pat_let_tv7 = p0
						_ = _pat_let_tv7
						var _409_v15 _dafny.MultiSet
						_ = _409_v15
						_409_v15 = _dafny.MultiSetOf(func(_pat_let11_0 D0) D0 {
							return func(_410_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let12_0 _dafny.Int) D0 {
									return func(_411_dt__update_hcf6_h0 _dafny.Int) D0 {
										return Companion_D0_.Create_DC2_((_410_dt__update__tmp_h0).Dtor_cf5(), _411_dt__update_hcf6_h0, (_410_dt__update__tmp_h0).Dtor_cf7())
									}(_pat_let12_0)
								}(_pat_let_tv7)
							}(_pat_let11_0)
						}(_408_v14))
						_409_v15 = (func() _dafny.MultiSet {
							if (_399_v6) || ((_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool)) {
								return (_dafny.MultiSetOf(_408_v14, Companion_D0_.Create_DC2_((_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool), p0, _407_v13), _408_v14)).Difference(_dafny.MultiSetOf(Companion_D0_.Create_DC2_((_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool), (_dafny.Zero).Minus(p0), _407_v13), _408_v14))
							}
							return _409_v15
						})()
						var _412_v16 _dafny.Array
						_ = _412_v16
						var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
						_ = _nw67
						_412_v16 = _nw67
						var _413_v17 _dafny.Map
						_ = _413_v17
						_413_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(140), (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool))
						var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_412_v16), 0))
						_ = _index53
						(_412_v16).ArraySet1(_413_v17, (_index53).Int())
						var _414_v18 _dafny.Sequence
						_ = _414_v18
						_414_v18 = _dafny.SeqOf(_404_v10)
						var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_412_v16), 0))
						_ = _index54
						(_412_v16).ArraySet1((_413_v17).Update((func() _dafny.Int {
							if _399_v6 {
								return _dafny.IntOfInt64(169)
							}
							return _dafny.IntOfUint32((_414_v18).Cardinality())
						})(), false), (_index54).Int())
						var _415_v19 _dafny.Sequence
						_ = _415_v19
						_415_v19 = _dafny.SeqOf(p0, p0)
						var _416_v20 _dafny.Map
						_ = _416_v20
						_416_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_415_v19).Cardinality()), p0), (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool))
						(globalState).F14 = ((Companion_Default___.Fm6((_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool), (func() bool {
							if (_413_v17).Contains(p0) {
								return (_413_v17).Get(p0).(bool)
							}
							return _399_v6
						})(), p0, globalState)).Merge(_416_v20)).Cardinality()
						var _rhs50 _dafny.Sequence = _404_v10
						_ = _rhs50
						var _rhs51 bool = _399_v6
						_ = _rhs51
						var _lhs42 *GlobalState = globalState
						_ = _lhs42
						r0 = _rhs50
						_lhs42.F3 = _rhs51
						var _417_v21 _dafny.Map
						_ = _417_v21
						_417_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (p0).Cmp(p0) <= 0)
						_417_v21 = (_417_v21).Update(_399_v6, (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool))
					}
					_398_v5 = (func() _dafny.Array {
						if (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool) {
							return _398_v5
						}
						return _398_v5
					})()
					var _418_v22 _dafny.Set
					_ = _418_v22
					_418_v22 = _dafny.SetOf(_399_v6)
					_418_v22 = (_418_v22).Union((_dafny.SetOf(_399_v6, (_398_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_398_v5), 0))).Int()).(bool))).Intersection(_418_v22))
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _419_v23 _dafny.Sequence
		_ = _419_v23
		_419_v23 = _dafny.UnicodeSeqOfUtf8Bytes("cmpkw")
		var _420_v24 _dafny.Sequence
		_ = _420_v24
		_420_v24 = _dafny.SeqOf(p0)
		var _421_v25 _dafny.MultiSet
		_ = _421_v25
		_421_v25 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_419_v23).Cardinality()), _dafny.IntOfUint32((_420_v24).Cardinality())), p0)
		var _422_v26 bool
		_ = _422_v26
		_422_v26 = false
		var _423_v27 _dafny.Map
		_ = _423_v27
		_423_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _422_v26), (_421_v25).Cardinality())
		var _424_v28 _dafny.Map
		_ = _424_v28
		_424_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, _422_v26)
		var _425_v29 _dafny.MultiSet
		_ = _425_v29
		_425_v29 = _dafny.MultiSetOf(_424_v28, _424_v28)
		(globalState).F14 = (func() _dafny.Int {
			if (_421_v25).Contains(((_423_v27).Cardinality()).Plus((func() _dafny.Int {
				if (_425_v29).Contains(_424_v28) {
					return (_425_v29).Multiplicity(_424_v28)
				}
				return p0
			})())) {
				return (_421_v25).Multiplicity(((_423_v27).Cardinality()).Plus((func() _dafny.Int {
					if (_425_v29).Contains(_424_v28) {
						return (_425_v29).Multiplicity(_424_v28)
					}
					return p0
				})()))
			}
			return _dafny.IntOfInt64(162)
		})()
		var _hi1 _dafny.Int = (_dafny.Zero).Minus(p0)
		_ = _hi1
		for _426_i2 := p0; _426_i2.Cmp(_hi1) < 0; _426_i2 = _426_i2.Plus(_dafny.One) {
			var _427_v30 _dafny.Map
			_ = _427_v30
			_427_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, _426_i2)
			var _428_v31 D0
			_ = _428_v31
			_428_v31 = Companion_D0_.Create_DC0_(_422_v26)
			var _429_v32 _dafny.Set
			_ = _429_v32
			_429_v32 = _dafny.SetOf(_422_v26, _422_v26)
			var _430_v33 D0
			_ = _430_v33
			_430_v33 = Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("fegnc"), Companion_Default___.Fm7(Companion_Default___.Fm7(_426_i2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, _426_i2), _428_v31, globalState), _427_v30, _428_v31, globalState), _426_i2, _392_v0)
			var _431_v34 _dafny.Map
			_ = _431_v34
			_431_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _392_v0)
			var _432_v35 _dafny.Array
			_ = _432_v35
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_14
			var _nw68 _dafny.Array
			_ = _nw68
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw68 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = (func(_433_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_434_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_434_i4, _433_i2)
					}
				})(_426_i2)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw68 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw68).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw68).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_432_v35 = _nw68
			var _435_v36 D0
			_ = _435_v36
			_435_v36 = Companion_D0_.Create_DC2_(_422_v26, _426_i2, _432_v35)
			var _pat_let_tv8 = _422_v26
			_ = _pat_let_tv8
			var _436_v37 _dafny.Array
			_ = _436_v37
			var _nwElement0_14 _dafny.Int = (_426_i2).Plus(_426_i2)
			_ = _nwElement0_14
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(25))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_14, 0)
			(_nw69).ArraySet1(Companion_Default___.Fm7(_426_i2, _427_v30, func(_pat_let13_0 D0) D0 {
				return func(_437_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let14_0 bool) D0 {
						return func(_438_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_438_dt__update_hcf0_h0)
						}(_pat_let14_0)
					}(_pat_let_tv8)
				}(_pat_let13_0)
			}(_428_v31), globalState), 1)
			(_nw69).ArraySet1(p0, 2)
			(_nw69).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5((_429_v32).Cardinality(), _420_v24, globalState), (_this).Fm5((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-495))).Uint32(), func(coer31 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_439_v23 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
				return func(_440_i3 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_dafny.IntOfUint32((_439_v23).Cardinality()))
				}
			})(_419_v23))))).Cardinality(), _420_v24, globalState))).Cardinality(), _426_i2), 3)
			(_nw69).ArraySet1((_424_v28).Cardinality(), 4)
			(_nw69).ArraySet1((_427_v30).Cardinality(), 5)
			(_nw69).ArraySet1((p0).Minus(p0), 6)
			(_nw69).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(669), _426_i2), 7)
			(_nw69).ArraySet1(p0, 8)
			(_nw69).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if _422_v26 {
					return p0
				}
				return p0
			})()), 9)
			(_nw69).ArraySet1((_430_v33).Dtor_cf2(), 10)
			(_nw69).ArraySet1(((Companion_Default___.Fm8(_422_v26, _422_v26, (func() _dafny.CodePoint {
				if (_431_v34).Contains(_dafny.IntOfUint32((_419_v23).Cardinality())) {
					return (_431_v34).Get(_dafny.IntOfUint32((_419_v23).Cardinality())).(_dafny.CodePoint)
				}
				return _392_v0
			})(), globalState)).Cardinality()).Times(_dafny.IntOfUint32((_394_v4).Cardinality())), 11)
			(_nw69).ArraySet1((_dafny.Zero).Minus((_435_v36).Dtor_cf6()), 12)
			(_nw69).ArraySet1(_426_i2, 13)
			(_nw69).ArraySet1(_dafny.IntOfInt64(193), 14)
			(_nw69).ArraySet1(_426_i2, 15)
			(_nw69).ArraySet1((_dafny.Zero).Minus(p0), 16)
			(_nw69).ArraySet1(_426_i2, 17)
			(_nw69).ArraySet1((_426_i2).Times(p0), 18)
			(_nw69).ArraySet1(_dafny.IntOfInt64(552), 19)
			(_nw69).ArraySet1(p0, 20)
			(_nw69).ArraySet1(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_419_v23).Cardinality())), 21)
			(_nw69).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_394_v4, _394_v4)).Cardinality()), 22)
			(_nw69).ArraySet1(_426_i2, 23)
			(_nw69).ArraySet1(p0, 24)
			_436_v37 = _nw69
			var _441_v38 _dafny.MultiSet
			_ = _441_v38
			_441_v38 = _dafny.MultiSetOf(_422_v26)
			var _442_v39 D2
			_ = _442_v39
			_442_v39 = Companion_D2_.Create_DC5_(_441_v38)
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_436_v37), 0))
			_ = _index55
			(_436_v37).ArraySet1(Companion_Default___.Fm7(Companion_Default___.Fm7((_427_v30).Cardinality(), _427_v30, _428_v31, globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, ((_442_v39).Dtor_cf12()).Cardinality()), _428_v31, globalState), (_index55).Int())
			var _443_v40 _dafny.Set
			_ = _443_v40
			_443_v40 = _dafny.SetOf(p0, p0)
			var _444_v41 _dafny.Map
			_ = _444_v41
			_444_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_443_v40).Cardinality(), (Companion_Default___.Fm9(_392_v0, globalState)).Cardinality())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_436_v37), 0))
			_ = _index56
			(_436_v37).ArraySet1((_dafny.Zero).Minus((_426_i2).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_444_v41).Contains(_dafny.IntOfInt64(813)) {
					return (_444_v41).Get(_dafny.IntOfInt64(813)).(_dafny.Int)
				}
				return _dafny.IntOfInt64(394)
			})()))), (_index56).Int())
			r3 = !((_this).Fm5((_436_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_436_v37), 0))).Int()).(_dafny.Int), _dafny.SeqOf(_426_i2), globalState))
			(globalState).F14 = (_dafny.Zero).Minus(p0)
			var _445_v42 *C0
			_ = _445_v42
			var _nw70 *C0 = New_C0_()
			_ = _nw70
			_nw70.Ctor__()
			_445_v42 = _nw70
		}
		var _446_v43 _dafny.Sequence
		_ = _446_v43
		_446_v43 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10(globalState), _419_v23))
		var _rhs52 bool = _422_v26
		_ = _rhs52
		var _rhs53 _dafny.Int = Companion_Default___.SafeDivisionInt((_420_v24).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_420_v24).Cardinality()))).Uint32()).(_dafny.Int), p0)
		_ = _rhs53
		var _rhs54 _dafny.MultiSet = ((_421_v25).Intersection((_421_v25).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ug")).Cardinality()), Companion_Default___.Abs(p0)))).Union(_421_v25)
		_ = _rhs54
		var _rhs55 _dafny.Sequence = _446_v43
		_ = _rhs55
		var _lhs43 *GlobalState = globalState
		_ = _lhs43
		var _lhs44 *GlobalState = globalState
		_ = _lhs44
		var _lhs45 *GlobalState = globalState
		_ = _lhs45
		_lhs43.F9 = _rhs52
		_lhs44.F14 = _rhs53
		_lhs45.F15 = _rhs54
		_446_v43 = _rhs55
		var _447_v44 _dafny.Array
		_ = _447_v44
		var _nw71 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw71
		_447_v44 = _nw71
		for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_447_v44), 0))); ; {
			_guard_loop_1, _ok36 := _iter36()
			if !_ok36 {
				break
			}
			var _448_i5 _dafny.Int
			_448_i5 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_448_i5).Sign() != -1) && ((_448_i5).Cmp(_dafny.ArrayLen((_447_v44), 0)) < 0)) {
				(_447_v44).ArraySet1(((p0).Minus((_dafny.Zero).Minus(p0))).Cmp((_dafny.SetOf(_422_v26, _422_v26, _422_v26, _422_v26)).Cardinality()) <= 0, (_448_i5).Int())
			}
		}
		if (_394_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.IntOfUint32((_394_v4).Cardinality()))).Uint32()).(bool) {
			_422_v26 = (_421_v25).IsSubsetOf((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs(_dafny.IntOfUint32((_394_v4).Cardinality()))))
			if !(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_394_v4, _394_v4), _422_v26)) {
				var _449_v45 _dafny.Map
				_ = _449_v45
				_449_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_420_v24).Cardinality()), _422_v26)
				var _450_v46 D0
				_ = _450_v46
				_450_v46 = Companion_D0_.Create_DC0_((func() bool {
					if (_449_v45).Contains(p0) {
						return (_449_v45).Get(p0).(bool)
					}
					return (_this).Fm5(p0, _420_v24, globalState)
				})())
				var _451_v47 _dafny.Map
				_ = _451_v47
				_451_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm7(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, p0), _450_v46, globalState))
				_451_v47 = _451_v47
				(globalState).F3 = _422_v26
				var _452_v48 _dafny.Map
				_ = _452_v48
				_452_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, p0)
				(globalState).F14 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(p0, (_452_v48).Cardinality()), p0)
				var _453_v49 _dafny.Array
				_ = _453_v49
				var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw72
				_453_v49 = _nw72
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_453_v49), 0))
				_ = _index57
				(_453_v49).ArraySet1(p0, (_index57).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_453_v49), 0))
				_ = _index58
				(_453_v49).ArraySet1(p0, (_index58).Int())
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_447_v44), 0))
				_ = _index59
				(_447_v44).ArraySet1(false, (_index59).Int())
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_447_v44), 0))
				_ = _index60
				(_447_v44).ArraySet1(_422_v26, (_index60).Int())
			} else {
				(globalState).F3 = _422_v26
				var _454_v50 _dafny.Map
				_ = _454_v50
				_454_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(_422_v26)), p0))
				var _455_v51 D0
				_ = _455_v51
				_455_v51 = Companion_D0_.Create_DC0_(_422_v26)
				var _456_v53 _dafny.Map
				_ = _456_v53
				_456_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _422_v26)
				var _457_v54 _dafny.Sequence
				_ = _457_v54
				_457_v54 = _dafny.SeqOf(_456_v53)
				_423_v27 = ((func() _dafny.Map {
					if (_454_v50).Contains(p0) {
						return (_454_v50).Get(p0).(_dafny.Map)
					}
					return _423_v27
				})()).Merge((func() _dafny.Map {
					if (_455_v51).Dtor_cf0() {
						return func() _dafny.Map {
							var _coll35 = _dafny.NewMapBuilder()
							_ = _coll35
							for _iter37 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_457_v54, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-468), _dafny.IntOfUint32((_457_v54).Cardinality()))).Uint32(), _456_v53)).Elements()); ; {
								_compr_35, _ok37 := _iter37()
								if !_ok37 {
									break
								}
								var _458_v52 _dafny.Map
								_458_v52 = interface{}(_compr_35).(_dafny.Map)
								if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_457_v54, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-468), _dafny.IntOfUint32((_457_v54).Cardinality()))).Uint32(), _456_v53), _458_v52) {
									_coll35.Add(_458_v52, p0)
								}
							}
							return _coll35.ToMap()
						}()
					}
					return _423_v27
				})())
				var _459_v55 _dafny.Map
				_ = _459_v55
				_459_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, (_420_v24).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_420_v24).Cardinality()))).Uint32()).(_dafny.Int))
				var _460_v56 _dafny.Array
				_ = _460_v56
				var _nwElement0_15 _dafny.Int = _dafny.IntOfInt64(122)
				_ = _nwElement0_15
				var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(9))
				_ = _nw73
				(_nw73).ArraySet1(_nwElement0_15, 0)
				(_nw73).ArraySet1(p0, 1)
				(_nw73).ArraySet1(_dafny.IntOfUint32((_394_v4).Cardinality()), 2)
				(_nw73).ArraySet1((_dafny.IntOfInt64(930)).Minus(p0), 3)
				(_nw73).ArraySet1(Companion_Default___.Fm7(p0, _459_v55, _455_v51, globalState), 4)
				(_nw73).ArraySet1(p0, 5)
				(_nw73).ArraySet1((p0).Times(p0), 6)
				(_nw73).ArraySet1(_dafny.IntOfInt64(693), 7)
				(_nw73).ArraySet1(_dafny.IntOfInt64(183), 8)
				_460_v56 = _nw73
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_460_v56), 0))
				_ = _index61
				(_460_v56).ArraySet1(p0, (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_460_v56), 0))
				_ = _index62
				(_460_v56).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), (_index62).Int())
				var _461_v57 _dafny.Set
				_ = _461_v57
				_461_v57 = _dafny.SetOf((func() _dafny.Int {
					if (_459_v55).Contains(_422_v26) {
						return (_459_v55).Get(_422_v26).(_dafny.Int)
					}
					return p0
				})())
				_461_v57 = _461_v57
				var _462_v58 *C0
				_ = _462_v58
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__()
				_462_v58 = _nw74
			}
			var _463_v59 _dafny.Map
			_ = _463_v59
			_463_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_447_v44, _419_v23)
			_463_v59 = ((_463_v59).Merge(_463_v59)).Merge(_463_v59)
			(globalState).F3 = (_this).Fm5(_dafny.IntOfInt64(655), _420_v24, globalState)
			(globalState).F14 = p0
		} else {
			var _464_v60 *C0
			_ = _464_v60
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__()
			_464_v60 = _nw75
			(globalState).F9 = _dafny.Companion_Sequence_.Equal(_419_v23, _dafny.Companion_Sequence_.Update(_419_v23, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.IntOfUint32((_419_v23).Cardinality()))).Uint32(), _392_v0))
			var _465_v61 _dafny.MultiSet
			_ = _465_v61
			_465_v61 = _dafny.MultiSetOf(_422_v26)
			var _466_v62 D2
			_ = _466_v62
			_466_v62 = Companion_D2_.Create_DC5_(_465_v61)
			var _467_v63 D1
			_ = _467_v63
			_467_v63 = Companion_D1_.Create_DC4_(_422_v26, Companion_Default___.Fm11(p0, _422_v26, _466_v62, globalState), p0)
			var _468_v64 _dafny.Map
			_ = _468_v64
			_468_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm5(_dafny.IntOfInt64(604), _420_v24, globalState))
			var _469_v65 _dafny.Set
			_ = _469_v65
			_469_v65 = _dafny.SetOf(_468_v64)
			var _470_v66 _dafny.Set
			_ = _470_v66
			_470_v66 = _dafny.SetOf(_467_v63, _467_v63, _467_v63, Companion_D1_.Create_DC4_(false, _469_v65, p0), _467_v63)
			var _471_v67 _dafny.Map
			_ = _471_v67
			_471_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, _470_v66)
			var _rhs56 _dafny.Map = _471_v67
			_ = _rhs56
			var _rhs57 _dafny.Int = (p0).Times(p0)
			_ = _rhs57
			var _rhs58 bool = !(false)
			_ = _rhs58
			var _rhs59 _dafny.Int = p0
			_ = _rhs59
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			var _lhs47 *GlobalState = globalState
			_ = _lhs47
			_471_v67 = _rhs56
			_lhs46.F14 = _rhs57
			_422_v26 = _rhs58
			_lhs47.F14 = _rhs59
			(globalState).F14 = p0
			var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw76
			(globalState).F7 = _nw76
		}
		r0 = _dafny.Companion_Sequence_.Concatenate(_419_v23, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_472_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_473_i6 _dafny.Int) _dafny.CodePoint {
				return _472_v0
			}
		})(_392_v0))), _419_v23))
		var _474_v69 _dafny.Map
		_ = _474_v69
		_474_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, p0)
		var _475_v70 _dafny.Sequence
		_ = _475_v70
		_475_v70 = _dafny.SeqOf(_474_v69, _474_v69)
		var _476_v71 _dafny.MultiSet
		_ = _476_v71
		_476_v71 = _dafny.MultiSetOf(_422_v26)
		var _477_v72 D0
		_ = _477_v72
		_477_v72 = Companion_D0_.Create_DC0_(_422_v26)
		var _478_v73 D0
		_ = _478_v73
		_478_v73 = Companion_D0_.Create_DC1_(_419_v23, ((func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter38 := _dafny.Iterate((_475_v70).Elements()); ; {
				_compr_36, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _479_v68 _dafny.Map
				_479_v68 = interface{}(_compr_36).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_475_v70, _479_v68) {
					_coll36.Add(_479_v68, _422_v26)
				}
			}
			return _coll36.ToMap()
		}()).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v26, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality()), !(true))).Cardinality(), (_476_v71).Cardinality(), Companion_Default___.Fm12(_477_v72, globalState))
		r1 = _478_v73
		var _480_v74 _dafny.Map
		_ = _480_v74
		_480_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_481_i8 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_dafny.IntOfInt64(-428))
		})))
		var _482_v75 _dafny.Set
		_ = _482_v75
		_482_v75 = _dafny.SetOf(_dafny.IntOfInt64(-46), p0)
		var _483_v76 _dafny.Sequence
		_ = _483_v76
		_483_v76 = _dafny.SeqOf(_482_v75, _482_v75, _482_v75)
		r2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}((func(_484_p0 _dafny.Int, _485_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
			return func(_486_i7 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_484_p0, _dafny.IntOfUint32((_485_v4).Cardinality()))
			}
		})(p0, _394_v4))), (func() _dafny.Sequence {
			if (_480_v74).Contains(p0) {
				return (_480_v74).Get(p0).(_dafny.Sequence)
			}
			return _483_v76
		})())
		r3 = _422_v26
		return r0, r1, r2, r3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f17 _dafny.Set
	_f18 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f17 = _dafny.EmptySet
	_this._f18 = false
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

func (_this *C2) F17() _dafny.Set {
	return _this._f17
}
func (_this *C2) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C2) F18() bool {
	return _this._f18
}
func (_this *C2) Ctor__(f17 _dafny.Set, f18 bool) {
	{
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C2) Fm19(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf((_this).F18()), _dafny.SetOf((_this).F18(), (_this).F18())), (func() _dafny.Sequence {
			if (_this).F18() {
				return _dafny.SeqOf(_dafny.SetOf((_this).F18(), (_this).F18(), (_this).F18()), _dafny.SetOf((_this).F18()))
			}
			return _dafny.SeqOf(_dafny.SetOf((_this).F18()), _dafny.SetOf((_this).F18()), _dafny.SetOf(!(false)))
		})())
	}
}
func (_this *C2) Fm20(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("lxunet")
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f17 _dafny.Set
	_f19 _dafny.Array
	_f18 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f18 = false
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

func (_this *C3) F17() _dafny.Set {
	return _this._f17
}
func (_this *C3) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C3) F19() _dafny.Array {
	return _this._f19
}
func (_this *C3) F19_set_(value _dafny.Array) {
	_this._f19 = value
}
func (_this *C3) F18() bool {
	return _this._f18
}
func (_this *C3) Ctor__(f19 _dafny.Array, f17 _dafny.Set, f18 bool) {
	{
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C3) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ooi")).Cardinality())
	}
}
func (_this *C3) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _487_v0 *C1
		_ = _487_v0
		var _nw77 *C1 = New_C1_()
		_ = _nw77
		_nw77.Ctor__()
		_487_v0 = _nw77
		(globalState).F9 = (func() bool {
			if !(p1) || (p1) {
				return p1
			}
			return (_this).F18()
		})()
		var _488_v1 _dafny.Map
		_ = _488_v1
		_488_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
		var _489_v2 _dafny.Set
		_ = _489_v2
		_489_v2 = _dafny.SetOf(_488_v1)
		var _490_v3 D1
		_ = _490_v3
		_490_v3 = Companion_D1_.Create_DC4_(true, _489_v2, _dafny.IntOfInt64(164))
		var _491_v4 _dafny.MultiSet
		_ = _491_v4
		_491_v4 = _dafny.MultiSetOf(_490_v3)
		var _492_v5 _dafny.Sequence
		_ = _492_v5
		_492_v5 = _dafny.SeqOf(Companion_D1_.Create_DC4_(p1, _489_v2, p2), _490_v3)
		var _493_v6 _dafny.Sequence
		_ = _493_v6
		_493_v6 = _dafny.SeqOf(p1, p1, (_this).F18())
		var _494_v7 _dafny.Map
		_ = _494_v7
		_494_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_493_v6).Cardinality())))
		var _495_v8 _dafny.MultiSet
		_ = _495_v8
		_495_v8 = _dafny.MultiSetOf(p2, p2)
		var _496_v9 _dafny.Map
		_ = _496_v9
		_496_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_492_v5)).IsSubsetOf(_491_v4), ((func() _dafny.MultiSet {
			if (_494_v7).Contains(p1) {
				return (_494_v7).Get(p1).(_dafny.MultiSet)
			}
			return _495_v8
		})()).Union(_495_v8))
		_496_v9 = _496_v9
		var _497_v10 _dafny.Sequence
		_ = _497_v10
		_497_v10 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _498_v11 _dafny.Array
		_ = _498_v11
		var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw78
		_498_v11 = _nw78
		var _499_v12 D0
		_ = _499_v12
		_499_v12 = Companion_D0_.Create_DC2_(p1, _dafny.IntOfUint32((_497_v10).Cardinality()), _498_v11)
		var _500_v13 _dafny.CodePoint
		_ = _500_v13
		_500_v13 = _dafny.CodePoint('v')
		var _501_v14 _dafny.Sequence
		_ = _501_v14
		_501_v14 = _dafny.SeqOf((_this).Fm0((_dafny.Zero).Minus(p2), _500_v13, globalState))
		var _502_v15 D3
		_ = _502_v15
		_502_v15 = Companion_D3_.Create_DC9_()
		var _503_v16 _dafny.Sequence
		_ = _503_v16
		_503_v16 = _dafny.SeqOf(_502_v15)
		var _504_v17 _dafny.Sequence
		_ = _504_v17
		_504_v17 = _503_v16
		var _505_v18 D5
		_ = _505_v18
		_505_v18 = Companion_D5_.Create_DC13_(p1, (_501_v14).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_501_v14).Cardinality()))).Uint32()).(_dafny.Int), _504_v17)
		if ((func() D0 {
			if p1 {
				return _499_v12
			}
			return Companion_D0_.Create_DC2_(p1, (_505_v18).Dtor_cf24(), _498_v11)
		})()).Dtor_cf5() {
			r0 = (_this).F18()
			var _506_v19 *C0
			_ = _506_v19
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__()
			_506_v19 = _nw79
			(globalState).F14 = p0
			var _507_v20 _dafny.Map
			_ = _507_v20
			_507_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F18()) || (false), (_dafny.Zero).Minus(p2))
			_507_v20 = _507_v20
			if ((_dafny.Zero).Minus(p0)).Cmp((p0).Plus(_dafny.IntOfUint32((_497_v10).Cardinality()))) != 0 {
				(globalState).F3 = !(p1) || ((p1) || ((_this).F18()))
				var _508_v21 _dafny.Set
				_ = _508_v21
				_508_v21 = _dafny.SetOf((_this).F18(), false, p1, p1, (_this).F18())
				_508_v21 = (Companion_D7_.Create_DC16_(_508_v21)).Dtor_cf30()
				var _509_v22 _dafny.Array
				_ = _509_v22
				var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw80
				_509_v22 = _nw80
				_509_v22 = (func() _dafny.Array {
					if (_this).F18() {
						return _509_v22
					}
					return _509_v22
				})()
				var _510_v23 _dafny.Array
				_ = _510_v23
				var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw81
				_510_v23 = _nw81
				_510_v23 = _510_v23
				var _511_v24 _dafny.Map
				_ = _511_v24
				_511_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)).Cardinality()) >= 0, _497_v10)
				_511_v24 = (_511_v24).Update(false, _497_v10)
			} else {
				var _512_v25 _dafny.Array
				_ = _512_v25
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_15
				var _nw82 _dafny.Array
				_ = _nw82
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw82 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Sequence = (func(_513_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_514_i0 _dafny.Int) _dafny.Sequence {
							return _513_v6
						}
					})(_493_v6)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw82 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw82).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw82).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_512_v25 = _nw82
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_512_v25), 0))
				_ = _index63
				(_512_v25).ArraySet1(_dafny.Companion_Sequence_.Update(_493_v6, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_493_v6).Cardinality()))).Uint32(), p1), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_512_v25), 0))
				_ = _index64
				(_512_v25).ArraySet1(_493_v6, (_index64).Int())
				var _rhs60 bool = false
				_ = _rhs60
				var _rhs61 _dafny.Int = (Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(p0))).Minus(p2)
				_ = _rhs61
				var _lhs48 *GlobalState = globalState
				_ = _lhs48
				r0 = _rhs60
				_lhs48.F14 = _rhs61
				var _515_v26 _dafny.Array
				_ = _515_v26
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
				_ = _nw83
				_515_v26 = _nw83
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_515_v26), 0))
				_ = _index65
				(_515_v26).ArraySet1CodePoint(_500_v13, (_index65).Int())
				var _516_v27 D3
				_ = _516_v27
				_516_v27 = Companion_D3_.Create_DC10_(_497_v10, p1, _497_v10, _dafny.Companion_Sequence_.Concatenate(_497_v10, _497_v10), p2)
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_515_v26), 0))
				_ = _index66
				var _rhs62 _dafny.Int = p0
				_ = _rhs62
				var _rhs63 _dafny.CodePoint = _500_v13
				_ = _rhs63
				var _rhs64 _dafny.Int = Companion_Default___.SafeDivisionInt((p2).Minus(p0), p0)
				_ = _rhs64
				var _rhs65 _dafny.MultiSet = _dafny.MultiSetOf(p2)
				_ = _rhs65
				var _rhs66 D3 = _516_v27
				_ = _rhs66
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				var _lhs50 _dafny.Array = _515_v26
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_515_v26), 0))
				_ = _lhs51
				var _lhs52 *GlobalState = globalState
				_ = _lhs52
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				_lhs49.F14 = _rhs62
				(_lhs50).ArraySet1CodePoint(_rhs63, (_lhs51).Int())
				_lhs52.F14 = _rhs64
				_lhs53.F15 = _rhs65
				_516_v27 = _rhs66
				var _517_v28 _dafny.MultiSet
				_ = _517_v28
				_517_v28 = _dafny.MultiSetOf(p1)
				var _518_v29 _dafny.Map
				_ = _518_v29
				_518_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(true, (_this).F18(), (_this).F18())).Cardinality()), _517_v28)
				_518_v29 = (_518_v29).Update(_dafny.IntOfUint32((_497_v10).Cardinality()), (_517_v28).Difference(_517_v28))
				var _519_v30 _dafny.Map
				_ = _519_v30
				_519_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(532))
				var _rhs67 bool = !((_519_v30).Contains((func() _dafny.Int {
					if p1 {
						return p0
					}
					return p2
				})()))
				_ = _rhs67
				var _rhs68 _dafny.Array = _this.F19()
				_ = _rhs68
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				var _lhs55 *C3 = _this
				_ = _lhs55
				_lhs54.F9 = _rhs67
				_lhs55.F19_set_(_rhs68)
			}
		} else {
			_497_v10 = _497_v10
			(globalState).F14 = p0
			_500_v13 = _500_v13
			var _520_v31 _dafny.Map
			_ = _520_v31
			_520_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _521_v32 _dafny.Sequence
			_ = _521_v32
			var _522_v33 D0
			_ = _522_v33
			var _523_v34 bool
			_ = _523_v34
			var _524_v35 bool
			_ = _524_v35
			var _out33 _dafny.Sequence
			_ = _out33
			var _out34 D0
			_ = _out34
			var _out35 bool
			_ = _out35
			var _out36 bool
			_ = _out36
			_out33, _out34, _out35, _out36 = (_487_v0).M5((func() _dafny.Int {
				if (_520_v31).Contains(_dafny.IntOfUint32((_497_v10).Cardinality())) {
					return (_520_v31).Get(_dafny.IntOfUint32((_497_v10).Cardinality())).(_dafny.Int)
				}
				return p0
			})(), globalState)
			_521_v32 = _out33
			_522_v33 = _out34
			_523_v34 = _out35
			_524_v35 = _out36
			(globalState).F3 = (p0).Cmp(p2) != 0
		}
		if false {
			(globalState).F3 = p1
			(globalState).F3 = (_this).F18()
			var _arr0 _dafny.Array = _this.F19()
			_ = _arr0
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index67
			(_arr0).ArraySet1(p1, (_index67).Int())
			var _525_v36 _dafny.Map
			_ = _525_v36
			_525_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.MultiSetOf(_498_v11, _498_v11, _498_v11))
			var _526_v37 _dafny.MultiSet
			_ = _526_v37
			_526_v37 = _dafny.MultiSetOf(_498_v11, _498_v11)
			var _arr1 _dafny.Array = _this.F19()
			_ = _arr1
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index68
			(_arr1).ArraySet1((func() bool {
				if (p1) || (false) {
					return (_this).F18()
				}
				return ((_dafny.Zero).Minus(((func() _dafny.MultiSet {
					if (_525_v36).Contains(p2) {
						return (_525_v36).Get(p2).(_dafny.MultiSet)
					}
					return _526_v37
				})()).Cardinality())).Cmp(_dafny.IntOfInt64(558)) != 0
			})(), (_index68).Int())
			var _527_v38 _dafny.Map
			_ = _527_v38
			_527_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_497_v10).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(p0)))
			var _528_v39 _dafny.MultiSet
			_ = _528_v39
			_528_v39 = _dafny.MultiSetOf(p1)
			var _529_v40 _dafny.Array
			_ = _529_v40
			var _nwElement0_16 bool = true
			_ = _nwElement0_16
			var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(22))
			_ = _nw84
			(_nw84).ArraySet1(_nwElement0_16, 0)
			(_nw84).ArraySet1(false, 1)
			(_nw84).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 2)
			(_nw84).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 3)
			(_nw84).ArraySet1((_this).F18(), 4)
			(_nw84).ArraySet1(p1, 5)
			(_nw84).ArraySet1((_this).F18(), 6)
			(_nw84).ArraySet1((_this).F18(), 7)
			(_nw84).ArraySet1((_this).F18(), 8)
			(_nw84).ArraySet1((_this).F18(), 9)
			(_nw84).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 10)
			(_nw84).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 11)
			(_nw84).ArraySet1(!((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)), 12)
			(_nw84).ArraySet1((_this).F18(), 13)
			(_nw84).ArraySet1((_this).F18(), 14)
			(_nw84).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 15)
			(_nw84).ArraySet1((_this).F18(), 16)
			(_nw84).ArraySet1(p1, 17)
			(_nw84).ArraySet1((_this).F18(), 18)
			(_nw84).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 19)
			(_nw84).ArraySet1(p1, 20)
			(_nw84).ArraySet1((_this).F18(), 21)
			_529_v40 = _nw84
			var _530_v41 _dafny.Map
			_ = _530_v41
			_530_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(Companion_D3_.Create_DC10_(_497_v10, (_this).F18(), _497_v10, _497_v10, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(337))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_531_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(664)
			}))).Cardinality())), Companion_Default___.Fm18(_dafny.CodePoint('d'), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), p2, p1, globalState), Companion_D3_.Create_DC10_(_497_v10, p1, _497_v10, _497_v10, p0)))
			var _532_v42 _dafny.Map
			_ = _532_v42
			_532_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_493_v6).Cardinality()))
			var _533_v44 _dafny.Set
			_ = _533_v44
			_533_v44 = _dafny.SetOf((func() _dafny.Int {
				if (_532_v42).Contains((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)) {
					return (_532_v42).Get((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)).(_dafny.Int)
				}
				return (func() _dafny.Map {
					var _coll37 = _dafny.NewMapBuilder()
					_ = _coll37
					for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(490), _dafny.IntOfInt64(49))); ; {
						_compr_37, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _534_v43 _dafny.Int
						_534_v43 = interface{}(_compr_37).(_dafny.Int)
						if ((_dafny.IntOfInt64(490)).Cmp(_534_v43) <= 0) && ((_534_v43).Cmp(_dafny.IntOfInt64(49)) < 0) {
							_coll37.Add(Companion_Default___.SafeModuloInt(_534_v43, p0), (_490_v3).Dtor_cf9())
						}
					}
					return _coll37.ToMap()
				}()).Cardinality()
			})(), p2)
			var _535_v45 T0
			_ = _535_v45
			var _nw85 *C2 = New_C2_()
			_ = _nw85
			_nw85.Ctor__(_this.F17(), (_this).F18())
			_535_v45 = _nw85
			var _536_v46 _dafny.Map
			_ = _536_v46
			_536_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_533_v44).Cardinality(), _535_v45)
			var _537_v47 _dafny.Map
			_ = _537_v47
			_537_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_487_v0).Fm5(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-349))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_538_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_539_i2 _dafny.Int) _dafny.Int {
					return _538_p2
				}
			})(p2)))).Cardinality()), _501_v14, globalState), true), _536_v46)
			var _nwElement0_17 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_527_v38).Contains(p0) {
					return (_527_v38).Get(p0).(_dafny.Int)
				}
				return p0
			})())
			_ = _nwElement0_17
			var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(27))
			_ = _nw86
			(_nw86).ArraySet1(_nwElement0_17, 0)
			(_nw86).ArraySet1(p2, 1)
			(_nw86).ArraySet1(p0, 2)
			(_nw86).ArraySet1((p0).Times(p2), 3)
			(_nw86).ArraySet1(p0, 4)
			(_nw86).ArraySet1(p2, 5)
			(_nw86).ArraySet1(p2, 6)
			(_nw86).ArraySet1(p2, 7)
			(_nw86).ArraySet1(p0, 8)
			(_nw86).ArraySet1((_this).Fm0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_497_v10, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_497_v10).Cardinality()))).Uint32(), _500_v13)).Cardinality()), _500_v13, globalState), 9)
			(_nw86).ArraySet1(p2, 10)
			(_nw86).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_528_v39).Cardinality(), _529_v40)).Cardinality(), 11)
			(_nw86).ArraySet1((_this).Fm0(p2, _500_v13, globalState), 12)
			(_nw86).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_493_v6).Cardinality()), p0), 13)
			(_nw86).ArraySet1((_dafny.Zero).Minus(((_530_v41).Cardinality()).Minus(p0)), 14)
			(_nw86).ArraySet1((_this).Fm0(p2, _500_v13, globalState), 15)
			(_nw86).ArraySet1(p0, 16)
			(_nw86).ArraySet1(p2, 17)
			(_nw86).ArraySet1((_537_v47).Cardinality(), 18)
			(_nw86).ArraySet1(p2, 19)
			(_nw86).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_497_v10).Cardinality()), (_this).Fm0(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pwulprhlu")).Cardinality()), _dafny.CodePoint('g'), globalState)), 20)
			(_nw86).ArraySet1(_dafny.IntOfInt64(738), 21)
			(_nw86).ArraySet1(p2, 22)
			(_nw86).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ytgmpt")).Cardinality()), 23)
			(_nw86).ArraySet1(p2, 24)
			(_nw86).ArraySet1(p2, 25)
			(_nw86).ArraySet1(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfUint32((_493_v6).Cardinality())), 26)
			(globalState).F7 = _nw86
			var _540_v48 _dafny.Sequence
			_ = _540_v48
			var _541_v49 D0
			_ = _541_v49
			var _542_v50 bool
			_ = _542_v50
			var _543_v51 bool
			_ = _543_v51
			var _out37 _dafny.Sequence
			_ = _out37
			var _out38 D0
			_ = _out38
			var _out39 bool
			_ = _out39
			var _out40 bool
			_ = _out40
			_out37, _out38, _out39, _out40 = (_487_v0).M5(p2, globalState)
			_540_v48 = _out37
			_541_v49 = _out38
			_542_v50 = _out39
			_543_v51 = _out40
		} else {
			if (_this).F18() {
				var _544_v52 *C2
				_ = _544_v52
				var _nw87 *C2 = New_C2_()
				_ = _nw87
				_nw87.Ctor__(_this.F17(), (func() bool {
					if (_this).F18() {
						return p1
					}
					return p1
				})())
				_544_v52 = _nw87
				var _545_v53 _dafny.Map
				_ = _545_v53
				_545_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), p2)
				var _rhs69 _dafny.Int = p0
				_ = _rhs69
				var _rhs70 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_544_v52).Fm20(globalState), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_this).F18() {
						return (func() _dafny.Int {
							if (_545_v53).Contains(p1) {
								return (_545_v53).Get(p1).(_dafny.Int)
							}
							return p0
						})()
					}
					return (_dafny.Zero).Minus(p0)
				})(), _dafny.IntOfUint32(((_544_v52).Fm20(globalState)).Cardinality()))).Uint32(), _500_v13)
				_ = _rhs70
				var _lhs56 *GlobalState = globalState
				_ = _lhs56
				_lhs56.F14 = _rhs69
				_497_v10 = _rhs70
				(globalState).F3 = p1
				var _546_v54 _dafny.Map
				_ = _546_v54
				_546_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())
				var _547_v55 _dafny.Map
				_ = _547_v55
				_547_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_487_v0).Fm5((_546_v54).Cardinality(), _501_v14, globalState), false)
				_547_v55 = (_546_v54).Merge((_546_v54).Update((_this).F18(), p1))
				(globalState).F9 = p1
			} else {
				r2 = p0
				var _548_v56 D5
				_ = _548_v56
				_548_v56 = Companion_D5_.Create_DC14_(p1, _dafny.IntOfInt64(-226), _dafny.IntOfUint32((_493_v6).Cardinality()))
				var _549_v57 _dafny.Map
				_ = _549_v57
				_549_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _550_v58 _dafny.Map
				_ = _550_v58
				_550_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _548_v56)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _548_v56)), _549_v57)
				_550_v58 = (_550_v58).Merge(_550_v58)
				var _551_v59 *C1
				_ = _551_v59
				var _nw88 *C1 = New_C1_()
				_ = _nw88
				_nw88.Ctor__()
				_551_v59 = _nw88
				var _552_v60 _dafny.Array
				_ = _552_v60
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw89
				_552_v60 = _nw89
				var _553_v61 _dafny.Set
				_ = _553_v61
				_553_v61 = _dafny.SetOf(_500_v13, _500_v13)
				var _554_v62 _dafny.Map
				_ = _554_v62
				_554_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SetOf((_497_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_497_v10).Cardinality()))).Uint32()).(_dafny.CodePoint), _500_v13, _dafny.CodePoint('h')))
				var _555_v63 _dafny.Sequence
				_ = _555_v63
				_555_v63 = _dafny.SeqOf(_553_v61, _dafny.SetOf(_500_v13, _500_v13))
				var _556_v65 _dafny.Array
				_ = _556_v65
				var _nwElement0_18 _dafny.Set = _553_v61
				_ = _nwElement0_18
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(15))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_18, 0)
				(_nw90).ArraySet1((func() _dafny.Set {
					if (_554_v62).Contains(true) {
						return (_554_v62).Get(true).(_dafny.Set)
					}
					return _553_v61
				})(), 1)
				(_nw90).ArraySet1(_553_v61, 2)
				(_nw90).ArraySet1(_553_v61, 3)
				(_nw90).ArraySet1(_553_v61, 4)
				(_nw90).ArraySet1(_553_v61, 5)
				(_nw90).ArraySet1(_553_v61, 6)
				(_nw90).ArraySet1((_555_v63).Select((Companion_Default___.SafeIndex((_this).Fm0(p0, _500_v13, globalState), _dafny.IntOfUint32((_555_v63).Cardinality()))).Uint32()).(_dafny.Set), 7)
				(_nw90).ArraySet1(func() _dafny.Set {
					var _coll38 = _dafny.NewBuilder()
					_ = _coll38
					for _iter40 := _dafny.Iterate((_497_v10).Elements()); ; {
						_compr_38, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _557_v64 _dafny.CodePoint
						_557_v64 = interface{}(_compr_38).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_497_v10, _557_v64) {
							_coll38.Add(_557_v64)
						}
					}
					return _coll38.ToSet()
				}(), 8)
				(_nw90).ArraySet1(_553_v61, 9)
				(_nw90).ArraySet1(_553_v61, 10)
				(_nw90).ArraySet1(_553_v61, 11)
				(_nw90).ArraySet1(_553_v61, 12)
				(_nw90).ArraySet1(_553_v61, 13)
				(_nw90).ArraySet1(_553_v61, 14)
				_556_v65 = _nw90
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_552_v60), 0))
				_ = _index69
				(_552_v60).ArraySet1(_556_v65, (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_552_v60), 0))
				_ = _index70
				(_552_v60).ArraySet1(_556_v65, (_index70).Int())
				var _558_v66 _dafny.Set
				_ = _558_v66
				_558_v66 = _dafny.SetOf(p2)
				(globalState).F9 = !(_558_v66).Contains(p0)
			}
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_498_v11), 0))
			_ = _index71
			(_498_v11).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-741), p0), (_index71).Int())
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_498_v11), 0))
			_ = _index72
			(_498_v11).ArraySet1(p2, (_index72).Int())
			var _559_v67 _dafny.Array
			_ = _559_v67
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_16
			var _nw91 _dafny.Array
			_ = _nw91
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw91 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Sequence = (func(_560_v10 _dafny.Sequence, _561_v11 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
					return func(_562_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_560_v10, (Companion_Default___.SafeIndex((_561_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_561_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_560_v10).Cardinality()))).Uint32(), _dafny.CodePoint('q'))
					}
				})(_497_v10, _498_v11)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw91 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw91).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw91).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_559_v67 = _nw91
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_559_v67), 0))
			_ = _index73
			(_559_v67).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_497_v10, _497_v10), (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_559_v67), 0))
			_ = _index74
			(_559_v67).ArraySet1(_497_v10, (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_559_v67), 0))
			_ = _index75
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_498_v11), 0))
			_ = _index76
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_559_v67), 0))
			_ = _index77
			var _rhs71 _dafny.Map = Companion_Default___.Fm21(_dafny.SeqOf((_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int), p0, p0, (_dafny.Zero).Minus(p0), p2), globalState)
			_ = _rhs71
			var _rhs72 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("nosuekq")
			_ = _rhs72
			var _rhs73 _dafny.Int = (_this).Fm0((_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int), _500_v13, globalState)
			_ = _rhs73
			var _rhs74 _dafny.Sequence = Companion_Default___.Fm22(p0, globalState)
			_ = _rhs74
			var _lhs57 _dafny.Array = _559_v67
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_559_v67), 0))
			_ = _lhs58
			var _lhs59 _dafny.Array = _498_v11
			_ = _lhs59
			var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_498_v11), 0))
			_ = _lhs60
			var _lhs61 _dafny.Array = _559_v67
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_559_v67), 0))
			_ = _lhs62
			_488_v1 = _rhs71
			(_lhs57).ArraySet1(_rhs72, (_lhs58).Int())
			(_lhs59).ArraySet1(_rhs73, (_lhs60).Int())
			(_lhs61).ArraySet1(_rhs74, (_lhs62).Int())
			var _arr2 _dafny.Array = _this.F19()
			_ = _arr2
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index78
			(_arr2).ArraySet1((_this).F18(), (_index78).Int())
			var _arr3 _dafny.Array = _this.F19()
			_ = _arr3
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index79
			(_arr3).ArraySet1(!((_this).F18()), (_index79).Int())
			var _563_v68 *C1
			_ = _563_v68
			var _nw92 *C1 = New_C1_()
			_ = _nw92
			_nw92.Ctor__()
			_563_v68 = _nw92
		}
		var _564_v69 D2
		_ = _564_v69
		_564_v69 = Companion_D2_.Create_DC7_(Companion_Default___.Fm23(p0, _dafny.IntOfInt64(942), (_this).F18(), globalState))
		var _565_v70 D2
		_ = _565_v70
		_565_v70 = Companion_D2_.Create_DC7_(_564_v69)
		var _source4 D2 = _565_v70
		_ = _source4
		if _source4.Is_DC6() {
			var _566___mcc_h0 _dafny.Int = _source4.Get_().(D2_DC6).Cf13
			_ = _566___mcc_h0
			var _567_cf13 _dafny.Int = _566___mcc_h0
			_ = _567_cf13
			var _568_v71 _dafny.Array
			_ = _568_v71
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw93
			_568_v71 = _nw93
			var _569_v72 _dafny.Array
			_ = _569_v72
			var _nwElement0_19 _dafny.Array = _498_v11
			_ = _nwElement0_19
			var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(3))
			_ = _nw94
			(_nw94).ArraySet1(_nwElement0_19, 0)
			(_nw94).ArraySet1(_498_v11, 1)
			(_nw94).ArraySet1(_498_v11, 2)
			_569_v72 = _nw94
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_568_v71), 0))
			_ = _index80
			(_568_v71).ArraySet1(_569_v72, (_index80).Int())
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_568_v71), 0))
			_ = _index81
			(_568_v71).ArraySet1(_569_v72, (_index81).Int())
			var _570_v73 *C2
			_ = _570_v73
			var _nw95 *C2 = New_C2_()
			_ = _nw95
			_nw95.Ctor__(_this.F17(), (_this).F18())
			_570_v73 = _nw95
			var _rhs75 _dafny.Int = _567_cf13
			_ = _rhs75
			var _rhs76 *C2 = _570_v73
			_ = _rhs76
			var _lhs63 *GlobalState = globalState
			_ = _lhs63
			_lhs63.F14 = _rhs75
			_570_v73 = _rhs76
			var _rhs77 bool = p1
			_ = _rhs77
			var _rhs78 _dafny.Int = p2
			_ = _rhs78
			var _lhs64 *GlobalState = globalState
			_ = _lhs64
			var _lhs65 *GlobalState = globalState
			_ = _lhs65
			_lhs64.F3 = _rhs77
			_lhs65.F14 = _rhs78
			(globalState).F3 = (((_this).F18()) && (p1)) || ((_this).F18())
		} else if _source4.Is_DC5() {
			var _571___mcc_h1 _dafny.MultiSet = _source4.Get_().(D2_DC5).Cf12
			_ = _571___mcc_h1
			var _572_cf12 _dafny.MultiSet = _571___mcc_h1
			_ = _572_cf12
			_497_v10 = _dafny.Companion_Sequence_.Concatenate(_497_v10, _497_v10)
			var _573_v74 _dafny.MultiSet
			_ = _573_v74
			_573_v74 = _dafny.MultiSetOf(_this.F19(), _this.F19(), _this.F19(), _this.F19())
			var _574_v75 _dafny.Map
			_ = _574_v75
			_574_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_501_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_501_v14).Cardinality()))).Uint32()).(_dafny.Int), p0)
			(globalState).F14 = ((_573_v74).Cardinality()).Times((func() _dafny.Int {
				if (_574_v75).Contains(p0) {
					return (_574_v75).Get(p0).(_dafny.Int)
				}
				return p0
			})())
			if !((_this).F18()) {
				(globalState).F9 = (_this).F18()
				var _575_v76 *C0
				_ = _575_v76
				var _nw96 *C0 = New_C0_()
				_ = _nw96
				_nw96.Ctor__()
				_575_v76 = _nw96
				(globalState).F9 = (_this).F18()
				var _576_v77 _dafny.Array
				_ = _576_v77
				var _nwElement0_20 *C1 = _487_v0
				_ = _nwElement0_20
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(5))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_20, 0)
				(_nw97).ArraySet1((func() *C1 {
					if p1 {
						return _487_v0
					}
					return _487_v0
				})(), 1)
				(_nw97).ArraySet1(_487_v0, 2)
				(_nw97).ArraySet1(_487_v0, 3)
				(_nw97).ArraySet1(_487_v0, 4)
				_576_v77 = _nw97
				var _577_v78 _dafny.Array
				_ = _577_v78
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_17
				var _nw98 _dafny.Array
				_ = _nw98
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw98 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.CodePoint = func(_578_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw98 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw98).ArraySet1CodePoint(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw98).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_577_v78 = _nw98
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_577_v78), 0))
				_ = _index82
				(_577_v78).ArraySet1CodePoint((func() _dafny.CodePoint {
					if p1 {
						return _500_v13
					}
					return _500_v13
				})(), (_index82).Int())
				var _579_v79 _dafny.Sequence
				_ = _579_v79
				_579_v79 = _dafny.SeqOf((func() _dafny.Array {
					if (_487_v0).Fm5(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(914))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}((func(_580_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_581_i5 _dafny.Int) _dafny.Int {
							return _580_p0
						}
					})(p0))), globalState) {
						return _576_v77
					}
					return _576_v77
				})())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_577_v78), 0))
				_ = _index83
				var _rhs79 _dafny.Array = (_579_v79).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_579_v79).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs79
				var _rhs80 _dafny.CodePoint = _500_v13
				_ = _rhs80
				var _lhs66 _dafny.Array = _577_v78
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_577_v78), 0))
				_ = _lhs67
				_576_v77 = _rhs79
				(_lhs66).ArraySet1CodePoint(_rhs80, (_lhs67).Int())
				(globalState).F3 = (_this).F18()
			} else {
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))
				_ = _index84
				(_498_v11).ArraySet1((p0).Plus(_dafny.IntOfInt64(205)), (_index84).Int())
				var _582_v80 _dafny.Set
				_ = _582_v80
				_582_v80 = _dafny.SetOf(_497_v10)
				var _583_v81 _dafny.Array
				_ = _583_v81
				var _nwElement0_21 _dafny.Sequence = _497_v10
				_ = _nwElement0_21
				var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(26))
				_ = _nw99
				(_nw99).ArraySet1(_nwElement0_21, 0)
				(_nw99).ArraySet1(_497_v10, 1)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_497_v10, _497_v10), 2)
				(_nw99).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(100))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_584_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_585_i6 _dafny.Int) _dafny.CodePoint {
						return _584_v13
					}
				})(_500_v13))), 3)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_497_v10, _497_v10), 4)
				(_nw99).ArraySet1(_497_v10, 5)
				(_nw99).ArraySet1(_497_v10, 6)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Update(_497_v10, (Companion_Default___.SafeIndex((_582_v80).Cardinality(), _dafny.IntOfUint32((_497_v10).Cardinality()))).Uint32(), _500_v13), 7)
				(_nw99).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gdic"), 8)
				(_nw99).ArraySet1(_497_v10, 9)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm22((_572_cf12).Cardinality(), globalState), _dafny.UnicodeSeqOfUtf8Bytes("fe")), 10)
				(_nw99).ArraySet1(_497_v10, 11)
				(_nw99).ArraySet1(_497_v10, 12)
				(_nw99).ArraySet1(_497_v10, 13)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_586_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_587_i7 _dafny.Int) _dafny.CodePoint {
						return _586_v13
					}
				})(_500_v13))), _497_v10), 14)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_497_v10, _dafny.UnicodeSeqOfUtf8Bytes("su")), 15)
				(_nw99).ArraySet1(_497_v10, 16)
				(_nw99).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_588_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_589_i8 _dafny.Int) _dafny.CodePoint {
						return _588_v13
					}
				})(_500_v13))), 17)
				(_nw99).ArraySet1(_497_v10, 18)
				(_nw99).ArraySet1(_497_v10, 19)
				(_nw99).ArraySet1(_497_v10, 20)
				(_nw99).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wfxv"), 21)
				(_nw99).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(917))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_590_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_591_i9 _dafny.Int) _dafny.CodePoint {
						return _590_v13
					}
				})(_500_v13))), 22)
				(_nw99).ArraySet1(_497_v10, 23)
				(_nw99).ArraySet1(Companion_Default___.Fm22(p2, globalState), 24)
				(_nw99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_497_v10, _497_v10), 25)
				_583_v81 = _nw99
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))
				_ = _index85
				(_583_v81).ArraySet1(_497_v10, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))
				_ = _index86
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))
				_ = _index87
				var _rhs81 bool = (p1) && ((_this).F18())
				_ = _rhs81
				var _rhs82 bool = p1
				_ = _rhs82
				var _rhs83 _dafny.Int = (p0).Minus((func() _dafny.Int {
					if (_this).F18() {
						return _dafny.IntOfInt64(394)
					}
					return p2
				})())
				_ = _rhs83
				var _rhs84 _dafny.Sequence = _497_v10
				_ = _rhs84
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				var _lhs70 _dafny.Array = _498_v11
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))
				_ = _lhs71
				var _lhs72 _dafny.Array = _583_v81
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))
				_ = _lhs73
				_lhs68.F3 = _rhs81
				_lhs69.F3 = _rhs82
				(_lhs70).ArraySet1(_rhs83, (_lhs71).Int())
				(_lhs72).ArraySet1(_rhs84, (_lhs73).Int())
				(globalState).F14 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yh"), (_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence)))).Cardinality())
				var _592_v82 *C2
				_ = _592_v82
				var _nw100 *C2 = New_C2_()
				_ = _nw100
				_nw100.Ctor__(_this.F17(), (_this).F18())
				_592_v82 = _nw100
				var _593_v83 _dafny.Map
				_ = _593_v83
				_593_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _592_v82)
				var _594_v84 _dafny.Map
				_ = _594_v84
				_594_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C2 {
					if (_593_v83).Contains(false) {
						return (_593_v83).Get(false).(*C2)
					}
					return _592_v82
				})(), (_this).F18())
				var _595_v85 _dafny.Sequence
				_ = _595_v85
				_595_v85 = _dafny.SeqOf(_594_v84)
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))
				_ = _index88
				(_498_v11).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_595_v85, _dafny.SeqOf(_594_v84, _594_v84)), _595_v85)).Cardinality()), (_index88).Int())
				var _596_v86 D0
				_ = _596_v86
				_596_v86 = Companion_D0_.Create_DC1_((_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence), (_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int), (_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int), ((_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32(((_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_596_v86 = (func() D0 {
					if (_this).F18() {
						return _596_v86
					}
					return _596_v86
				})()
				var _arr4 _dafny.Array = _this.F19()
				_ = _arr4
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index89
				(_arr4).ArraySet1(true, (_index89).Int())
				var _597_v87 _dafny.Array
				_ = _597_v87
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw101
				_597_v87 = _nw101
				var _598_v88 _dafny.Sequence
				_ = _598_v88
				_598_v88 = _dafny.SeqOf(_597_v87, _597_v87)
				var _599_v89 _dafny.Map
				_ = _599_v89
				_599_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-552), _597_v87)
				var _arr5 _dafny.Array = _this.F19()
				_ = _arr5
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index90
				var _rhs85 _dafny.Array = (func() _dafny.Array {
					if _dafny.Companion_Sequence_.IsProperPrefixOf((_583_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_583_v81), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("g")) {
						return (_598_v88).Select((Companion_Default___.SafeIndex((_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_598_v88).Cardinality()))).Uint32()).(_dafny.Array)
					}
					return (func() _dafny.Array {
						if (_599_v89).Contains((_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int)) {
							return (_599_v89).Get((_498_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_498_v11), 0))).Int()).(_dafny.Int)).(_dafny.Array)
						}
						return _597_v87
					})()
				})()
				_ = _rhs85
				var _rhs86 bool = p1
				_ = _rhs86
				var _rhs87 bool = (_this).F18()
				_ = _rhs87
				var _rhs88 bool = (func() bool {
					if (_this).F18() {
						return p1
					}
					return (_this).F18()
				})()
				_ = _rhs88
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				var _lhs76 _dafny.Array = _this.F19()
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_this.F19()), 0))
				_ = _lhs77
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				_lhs74.F7 = _rhs85
				_lhs75.F3 = _rhs86
				(_lhs76).ArraySet1(_rhs87, (_lhs77).Int())
				_lhs78.F9 = _rhs88
			}
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_498_v11), 0))
			_ = _index91
			(_498_v11).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.SetOf(_dafny.IntOfUint32((_493_v6).Cardinality()), _dafny.IntOfUint32((_501_v14).Cardinality()))).Cardinality(), p0), (_index91).Int())
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_498_v11), 0))
			_ = _index92
			(_498_v11).ArraySet1(p0, (_index92).Int())
		} else {
			var _600___mcc_h2 D2 = _source4.Get_().(D2_DC7).Cf14
			_ = _600___mcc_h2
			var _601_cf14 D2 = _600___mcc_h2
			_ = _601_cf14
			var _602_v90 _dafny.Map
			_ = _602_v90
			_602_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F19())
			_602_v90 = (_602_v90).Update(((_this).F18()) == ((_this).F18()), _this.F19())
			r1 = (_this).F18()
			r0 = !((_this).F18())
			var _603_v91 _dafny.Map
			_ = _603_v91
			_603_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_498_v11, _497_v10)
			_603_v91 = (_603_v91).Update(_498_v11, _497_v10)
		}
		var _604_v92 _dafny.Set
		_ = _604_v92
		_604_v92 = _dafny.SetOf(p0)
		var _605_v93 _dafny.Sequence
		_ = _605_v93
		_605_v93 = _dafny.SeqOf(_604_v92)
		r0 = (func() bool {
			if (func() _dafny.Set {
				var _coll39 = _dafny.NewBuilder()
				_ = _coll39
				for _iter41 := _dafny.Iterate((_605_v93).Elements()); ; {
					_compr_39, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _606_v94 _dafny.Set
					_606_v94 = interface{}(_compr_39).(_dafny.Set)
					if _dafny.Companion_Sequence_.Contains(_605_v93, _606_v94) {
						_coll39.Add(_606_v94)
					}
				}
				return _coll39.ToSet()
			}()).Contains(_604_v92) {
				return !((func() bool {
					if (_this).F18() {
						return (_this).F18()
					}
					return (_this).F18()
				})())
			}
			return p1
		})()
		r1 = (_487_v0).Fm5(Companion_Default___.SafeDivisionInt(p2, p0), _dafny.Companion_Sequence_.Concatenate(_501_v14, _dafny.SeqOf(p0)), globalState)
		r2 = p2
		return r0, r1, r2
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f17 _dafny.Set
	_f18 bool
	F26  _dafny.Int
	_f25 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f17 = _dafny.EmptySet
	_this._f18 = false
	_this.F26 = _dafny.Zero
	_this._f25 = _dafny.Zero
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

func (_this *C4) F17() _dafny.Set {
	return _this._f17
}
func (_this *C4) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C4) F18() bool {
	return _this._f18
}
func (_this *C4) Ctor__(f25 _dafny.Int, f26 _dafny.Int, f17 _dafny.Set, f18 bool) {
	{
		(_this)._f25 = f25
		(_this).F26 = f26
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C4) M3(p0 _dafny.CodePoint, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _607_v0 bool
		_ = _607_v0
		var _608_v1 _dafny.Map
		_ = _608_v1
		var _609_v2 _dafny.Int
		_ = _609_v2
		var _out41 bool
		_ = _out41
		var _out42 _dafny.Map
		_ = _out42
		var _out43 _dafny.Int
		_ = _out43
		_out41, _out42, _out43 = (_this).M4(globalState)
		_607_v0 = _out41
		_608_v1 = _out42
		_609_v2 = _out43
		var _610_v3 _dafny.Array
		_ = _610_v3
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_18
		var _nw102 _dafny.Array
		_ = _nw102
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw102 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) _dafny.CodePoint = (func(_611_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_612_i0 _dafny.Int) _dafny.CodePoint {
					return _611_p0
				}
			})(p0)
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw102 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw102).ArraySet1CodePoint(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw102).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		_610_v3 = _nw102
		var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_610_v3), 0))
		_ = _index93
		(_610_v3).ArraySet1CodePoint(p0, (_index93).Int())
		var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_610_v3), 0))
		_ = _index94
		(_610_v3).ArraySet1CodePoint(p0, (_index94).Int())
		var _613_v4 _dafny.Array
		_ = _613_v4
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_19
		var _nw103 _dafny.Array
		_ = _nw103
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw103 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) _dafny.Int = (func(_614_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_615_i2 _dafny.Int) _dafny.Int {
					return (_615_i2).Plus(_614_v2)
				}
			})(_609_v2)
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw103 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw103).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw103).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_613_v4 = _nw103
		for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_613_v4), 0))); ; {
			_guard_loop_2, _ok42 := _iter42()
			if !_ok42 {
				break
			}
			var _616_i1 _dafny.Int
			_616_i1 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_616_i1).Sign() != -1) && ((_616_i1).Cmp(_dafny.ArrayLen((_613_v4), 0)) < 0)) {
				(_613_v4).ArraySet1((_616_i1).Times(_this.F26), (_616_i1).Int())
			}
		}
		var _617_v5 _dafny.Sequence
		_ = _617_v5
		_617_v5 = _dafny.UnicodeSeqOfUtf8Bytes("n")
		var _618_v6 _dafny.Sequence
		_ = _618_v6
		_618_v6 = _dafny.SeqOf(_617_v5, _dafny.UnicodeSeqOfUtf8Bytes("vvpbsjve"))
		_618_v6 = _dafny.Companion_Sequence_.Update(_618_v6, (Companion_Default___.SafeIndex(_609_v2, _dafny.IntOfUint32((_618_v6).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_617_v5, _dafny.UnicodeSeqOfUtf8Bytes("aiiyumo")))
		(_this).F26 = ((_this).F25()).Times(_609_v2)
		(_this).F26 = _this.F26
		r0 = _dafny.IntOfInt64(636)
		return r0
	}
}
func (_this *C4) M4(globalState *GlobalState) (bool, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _619_v0 _dafny.Array
		_ = _619_v0
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_20
		var _nw104 _dafny.Array
		_ = _nw104
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw104 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = func(_620_i0 _dafny.Int) _dafny.Int {
				return (_620_i0).Times(_dafny.IntOfUint32((_dafny.SeqOf(_this.F26, (_this).F25(), (_this).F25())).Cardinality()))
			}
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw104 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw104).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw104).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_619_v0 = _nw104
		var _pat_let_tv9 = _619_v0
		_ = _pat_let_tv9
		var _source5 D0 = func(_pat_let15_0 D0) D0 {
			return func(_621_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let16_0 _dafny.Int) D0 {
					return func(_622_dt__update_hcf6_h0 _dafny.Int) D0 {
						return func(_pat_let17_0 _dafny.Array) D0 {
							return func(_623_dt__update_hcf7_h0 _dafny.Array) D0 {
								return Companion_D0_.Create_DC2_((_621_dt__update__tmp_h0).Dtor_cf5(), _622_dt__update_hcf6_h0, _623_dt__update_hcf7_h0)
							}(_pat_let17_0)
						}(_pat_let_tv9)
					}(_pat_let16_0)
				}(_this.F26)
			}(_pat_let15_0)
		}(Companion_D0_.Create_DC2_((_this).F18(), (_this).F25(), _619_v0))
		_ = _source5
		if _source5.Is_DC0() {
			var _624___mcc_h0 bool = _source5.Get_().(D0_DC0).Cf0
			_ = _624___mcc_h0
			var _625_cf0 bool = _624___mcc_h0
			_ = _625_cf0
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))
			_ = _index95
			(_619_v0).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F25(), _this.F26), (_index95).Int())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))
			_ = _index96
			(_619_v0).ArraySet1(_dafny.IntOfInt64(646), (_index96).Int())
			var _626_v1 _dafny.Map
			_ = _626_v1
			_626_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)).Cmp(_this.F26) >= 0, false)
			var _627_v2 _dafny.Map
			_ = _627_v2
			_627_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F18())
			var _628_v3 _dafny.MultiSet
			_ = _628_v3
			_628_v3 = _dafny.MultiSetOf(_627_v2)
			_626_v1 = Companion_Default___.Fm3(_this.F26, true, _628_v3, globalState)
			var _source6 D0 = Companion_D0_.Create_DC0_((_this).F18())
			_ = _source6
			if _source6.Is_DC0() {
				var _629___mcc_h8 bool = _source6.Get_().(D0_DC0).Cf0
				_ = _629___mcc_h8
				var _630_cf0 bool = _629___mcc_h8
				_ = _630_cf0
				var _631_v4 _dafny.Sequence
				_ = _631_v4
				_631_v4 = _dafny.UnicodeSeqOfUtf8Bytes("qnhyyx")
				_631_v4 = _631_v4
				var _632_v5 _dafny.CodePoint
				_ = _632_v5
				_632_v5 = _dafny.CodePoint('w')
				_632_v5 = _632_v5
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))
				_ = _index97
				(_619_v0).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm4((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), globalState), _this.F26), (_index97).Int())
				var _633_v6 _dafny.Set
				_ = _633_v6
				_633_v6 = _dafny.SetOf(_630_cf0)
				var _634_v7 _dafny.Array
				_ = _634_v7
				var _nwElement0_22 _dafny.Int = (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)
				_ = _nwElement0_22
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(5))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_22, 0)
				(_nw105).ArraySet1((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), 1)
				(_nw105).ArraySet1((_this).F25(), 2)
				(_nw105).ArraySet1(_this.F26, 3)
				(_nw105).ArraySet1((_633_v6).Cardinality(), 4)
				_634_v7 = _nw105
				var _635_v8 D0
				_ = _635_v8
				_635_v8 = Companion_D0_.Create_DC2_(_630_cf0, (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), _634_v7)
				var _636_v9 _dafny.MultiSet
				_ = _636_v9
				_636_v9 = _dafny.MultiSetOf((_635_v8).Dtor_cf5())
				var _637_v10 _dafny.Set
				_ = _637_v10
				_637_v10 = _dafny.SetOf((_636_v9).Cardinality())
				var _638_v11 _dafny.Set
				_ = _638_v11
				_638_v11 = _dafny.SetOf(_637_v10, _637_v10, _637_v10)
				var _639_v12 _dafny.Map
				_ = _639_v12
				_639_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_637_v10, _dafny.SetOf(_this.F26))
				(globalState).F9 = ((func() _dafny.Set {
					var _coll40 = _dafny.NewBuilder()
					_ = _coll40
					for _iter43 := _dafny.Iterate((_639_v12).Keys().Elements()); ; {
						_compr_40, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _640_v13 _dafny.Set
						_640_v13 = interface{}(_compr_40).(_dafny.Set)
						if (_639_v12).Contains(_640_v13) {
							_coll40.Add(_640_v13)
						}
					}
					return _coll40.ToSet()
				}()).Intersection(_638_v11)).IsProperSubsetOf((_638_v11).Union(_638_v11))
			} else if _source6.Is_DC1() {
				var _641___mcc_h9 _dafny.Sequence = _source6.Get_().(D0_DC1).Cf1
				_ = _641___mcc_h9
				var _642___mcc_h10 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
				_ = _642___mcc_h10
				var _643___mcc_h11 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
				_ = _643___mcc_h11
				var _644___mcc_h12 _dafny.CodePoint = _source6.Get_().(D0_DC1).Cf4
				_ = _644___mcc_h12
				var _645_cf4 _dafny.CodePoint = _644___mcc_h12
				_ = _645_cf4
				var _646_cf3 _dafny.Int = _643___mcc_h11
				_ = _646_cf3
				var _647_cf2 _dafny.Int = _642___mcc_h10
				_ = _647_cf2
				var _648_cf1 _dafny.Sequence = _641___mcc_h9
				_ = _648_cf1
				var _649_v14 _dafny.MultiSet
				_ = _649_v14
				_649_v14 = _dafny.MultiSetOf((_this).F18(), (_this).F18(), !(_625_cf0))
				_649_v14 = _649_v14
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))
				_ = _index98
				var _rhs89 _dafny.Int = _this.F26
				_ = _rhs89
				var _rhs90 _dafny.Int = _this.F26
				_ = _rhs90
				var _rhs91 _dafny.Int = (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)
				_ = _rhs91
				var _lhs79 _dafny.Array = _619_v0
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))
				_ = _lhs80
				_647_cf2 = _rhs89
				(_lhs79).ArraySet1(_rhs90, (_lhs80).Int())
				_646_cf3 = _rhs91
				var _650_v15 _dafny.Array
				_ = _650_v15
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw106
				_650_v15 = _nw106
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_650_v15), 0))
				_ = _index99
				(_650_v15).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_648_cf1, _648_cf1), (_index99).Int())
				var _651_v16 _dafny.Map
				_ = _651_v16
				_651_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_625_cf0), _645_cf4)
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_650_v15), 0))
				_ = _index100
				(_650_v15).ArraySet1(_dafny.Companion_Sequence_.Update(_648_cf1, (Companion_Default___.SafeIndex((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_648_cf1).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_651_v16).Contains((_this).F18()) {
						return (_651_v16).Get((_this).F18()).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('p')
				})()), (_index100).Int())
				var _652_v17 *C0
				_ = _652_v17
				var _nw107 *C0 = New_C0_()
				_ = _nw107
				_nw107.Ctor__()
				_652_v17 = _nw107
			} else {
				var _653___mcc_h13 bool = _source6.Get_().(D0_DC2).Cf5
				_ = _653___mcc_h13
				var _654___mcc_h14 _dafny.Int = _source6.Get_().(D0_DC2).Cf6
				_ = _654___mcc_h14
				var _655___mcc_h15 _dafny.Array = _source6.Get_().(D0_DC2).Cf7
				_ = _655___mcc_h15
				var _656_cf7 _dafny.Array = _655___mcc_h15
				_ = _656_cf7
				var _657_cf6 _dafny.Int = _654___mcc_h14
				_ = _657_cf6
				var _658_cf5 bool = _653___mcc_h13
				_ = _658_cf5
				var _659_v18 _dafny.Array
				_ = _659_v18
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_21
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) bool = (func(_660_cf0 bool) func(_dafny.Int) bool {
						return func(_661_i1 _dafny.Int) bool {
							return (func() bool {
								if true {
									return _660_cf0
								}
								return _660_cf0
							})()
						}
					})(_625_cf0)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw108 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw108).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw108).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_659_v18 = _nw108
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_659_v18), 0))
				_ = _index101
				(_659_v18).ArraySet1(false, (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_659_v18), 0))
				_ = _index102
				(_659_v18).ArraySet1(_658_cf5, (_index102).Int())
				var _662_v19 _dafny.Array
				_ = _662_v19
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_22
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_663_v1 _dafny.Map) func(_dafny.Int) _dafny.Sequence {
						return func(_664_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_663_v1), (Companion_D3_.Create_DC8_(_dafny.SeqOf(_663_v1))).Dtor_cf15())
						}
					})(_626_v1)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw109 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw109).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw109).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_662_v19 = _nw109
				var _665_v20 _dafny.Sequence
				_ = _665_v20
				_665_v20 = _dafny.SeqOf(_626_v1)
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_662_v19), 0))
				_ = _index103
				(_662_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v20, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-803))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_666_v1 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_667_i3 _dafny.Int) _dafny.Map {
						return _666_v1
					}
				})(_626_v1)))), (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_662_v19), 0))
				_ = _index104
				(_662_v19).ArraySet1(_665_v20, (_index104).Int())
				var _668_v21 D0
				_ = _668_v21
				_668_v21 = Companion_D0_.Create_DC0_(_658_cf5)
				var _669_v22 _dafny.Map
				_ = _669_v22
				_669_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13(globalState), _668_v21)
				_669_v22 = (_669_v22).Update(_625_cf0, _668_v21)
				(globalState).F14 = Companion_Default___.SafeModuloInt(_657_cf6, (_657_cf6).Plus((_this).F25()))
			}
			var _670_v23 _dafny.CodePoint
			_ = _670_v23
			_670_v23 = _dafny.CodePoint('c')
			var _671_v24 _dafny.Sequence
			_ = _671_v24
			_671_v24 = _dafny.SeqOf(_670_v23, _670_v23)
			var _source7 D3 = Companion_D3_.Create_DC10_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}(func(_672_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), _671_v24), _625_cf0, _dafny.Companion_Sequence_.Concatenate(_671_v24, _671_v24), _671_v24, Companion_Default___.Fm4((_626_v1).Cardinality(), globalState))
			_ = _source7
			if _source7.Is_DC9() {
				(globalState).F14 = Companion_Default___.SafeModuloInt((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), (_this).F25())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))
				_ = _index105
				(_619_v0).ArraySet1(((_this).F25()).Minus(((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)).Times((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int))), (_index105).Int())
				var _673_v25 _dafny.Array
				_ = _673_v25
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw110
				_673_v25 = _nw110
				var _674_v26 _dafny.Map
				_ = _674_v26
				_674_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_673_v25, _670_v23)
				_674_v26 = (_674_v26).Update(_673_v25, _670_v23)
				(globalState).F3 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-115))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_675_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_676_i5 _dafny.Int) _dafny.CodePoint {
						return _675_v23
					}
				})(_670_v23))), _670_v23)
			} else if _source7.Is_DC10() {
				var _677___mcc_h16 _dafny.Sequence = _source7.Get_().(D3_DC10).Cf16
				_ = _677___mcc_h16
				var _678___mcc_h17 bool = _source7.Get_().(D3_DC10).Cf17
				_ = _678___mcc_h17
				var _679___mcc_h18 _dafny.Sequence = _source7.Get_().(D3_DC10).Cf18
				_ = _679___mcc_h18
				var _680___mcc_h19 _dafny.Sequence = _source7.Get_().(D3_DC10).Cf19
				_ = _680___mcc_h19
				var _681___mcc_h20 _dafny.Int = _source7.Get_().(D3_DC10).Cf20
				_ = _681___mcc_h20
				var _682_cf20 _dafny.Int = _681___mcc_h20
				_ = _682_cf20
				var _683_cf19 _dafny.Sequence = _680___mcc_h19
				_ = _683_cf19
				var _684_cf18 _dafny.Sequence = _679___mcc_h18
				_ = _684_cf18
				var _685_cf17 bool = _678___mcc_h17
				_ = _685_cf17
				var _686_cf16 _dafny.Sequence = _677___mcc_h16
				_ = _686_cf16
				(globalState).F3 = _685_cf17
				var _687_v27 _dafny.Array
				_ = _687_v27
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_23
				var _nw111 _dafny.Array
				_ = _nw111
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw111 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_688_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_689_i6 _dafny.Int) _dafny.Int {
							return (_689_i6).Times((_688_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_688_v0), 0))).Int()).(_dafny.Int))
						}
					})(_619_v0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw111 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw111).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw111).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_687_v27 = _nw111
				var _690_v28 _dafny.Set
				_ = _690_v28
				_690_v28 = _dafny.SetOf(_687_v27, _687_v27, _687_v27)
				var _691_v29 D3
				_ = _691_v29
				_691_v29 = Companion_D3_.Create_DC9_()
				var _692_v30 _dafny.Sequence
				_ = _692_v30
				_692_v30 = _dafny.SeqOf(_691_v29)
				var _693_v31 _dafny.Sequence
				_ = _693_v31
				_693_v31 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(316))).Uint32(), func(coer45 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_694_v29 D3) func(_dafny.Int) D3 {
					return func(_695_i8 _dafny.Int) D3 {
						return _694_v29
					}
				})(_691_v29)))
				var _696_v32 _dafny.Sequence
				_ = _696_v32
				_696_v32 = _dafny.SeqOf(_692_v30, _692_v30, _692_v30)
				var _697_v33 _dafny.Array
				_ = _697_v33
				var _nwElement0_23 _dafny.Sequence = _692_v30
				_ = _nwElement0_23
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(26))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_23, 0)
				(_nw112).ArraySet1(_692_v30, 1)
				(_nw112).ArraySet1(_dafny.SeqOf(_691_v29), 2)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _692_v30), 3)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _692_v30), 4)
				(_nw112).ArraySet1(_692_v30, 5)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(54))).Uint32(), func(coer46 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_698_v29 D3) func(_dafny.Int) D3 {
					return func(_699_i7 _dafny.Int) D3 {
						return _698_v29
					}
				})(_691_v29)))), 6)
				(_nw112).ArraySet1((_693_v31), 7)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _dafny.SeqOf(_691_v29)), 8)
				(_nw112).ArraySet1((_696_v32).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_696_v32).Cardinality()))).Uint32()).(_dafny.Sequence), 9)
				(_nw112).ArraySet1(_692_v30, 10)
				(_nw112).ArraySet1(_692_v30, 11)
				(_nw112).ArraySet1(_dafny.SeqOf(_691_v29), 12)
				(_nw112).ArraySet1(_692_v30, 13)
				(_nw112).ArraySet1(_692_v30, 14)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _692_v30), 15)
				(_nw112).ArraySet1(_692_v30, 16)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer47 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_700_v29 D3) func(_dafny.Int) D3 {
					return func(_701_i9 _dafny.Int) D3 {
						return _700_v29
					}
				})(_691_v29))), _692_v30), 17)
				(_nw112).ArraySet1(Companion_Default___.Fm14(_685_cf17, (_this).F18(), _dafny.IntOfInt64(-325), globalState), 18)
				(_nw112).ArraySet1(_692_v30, 19)
				(_nw112).ArraySet1(_dafny.SeqOf(_691_v29), 20)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _692_v30), 21)
				(_nw112).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_692_v30, _dafny.SeqOf(_691_v29, Companion_D3_.Create_DC9_(), _691_v29)), 22)
				(_nw112).ArraySet1(_692_v30, 23)
				(_nw112).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(118))).Uint32(), func(coer48 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_702_v29 D3) func(_dafny.Int) D3 {
					return func(_703_i10 _dafny.Int) D3 {
						return _702_v29
					}
				})(_691_v29))), 24)
				(_nw112).ArraySet1(_692_v30, 25)
				_697_v33 = _nw112
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_697_v33), 0))
				_ = _index106
				(_697_v33).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(394))).Uint32(), func(coer49 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_704_v29 D3) func(_dafny.Int) D3 {
					return func(_705_i11 _dafny.Int) D3 {
						return _704_v29
					}
				})(_691_v29))), (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_697_v33), 0))
				_ = _index107
				var _rhs92 bool = Companion_Default___.Fm13(globalState)
				_ = _rhs92
				var _rhs93 _dafny.Set = _690_v28
				_ = _rhs93
				var _rhs94 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_692_v30, (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_692_v30).Cardinality()))).Uint32(), _691_v29)
				_ = _rhs94
				var _lhs81 _dafny.Array = _697_v33
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_697_v33), 0))
				_ = _lhs82
				_625_cf0 = _rhs92
				_690_v28 = _rhs93
				(_lhs81).ArraySet1(_rhs94, (_lhs82).Int())
				_685_cf17 = !((_685_cf17) && ((func() bool {
					if _625_cf0 {
						return !(_685_cf17)
					}
					return _685_cf17
				})()))
				var _706_v34 _dafny.Sequence
				_ = _706_v34
				_706_v34 = _dafny.SeqOf((_this).F18())
				var _707_v35 D3
				_ = _707_v35
				_707_v35 = Companion_D3_.Create_DC10_(_671_v24, _625_cf0, _686_cf16, _686_cf16, (_this).F25())
				var _708_v36 _dafny.Map
				_ = _708_v36
				_708_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_706_v34).Select((Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_706_v34).Cardinality()))).Uint32()).(bool), _707_v35)
				var _709_v37 _dafny.Map
				_ = _709_v37
				_709_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_708_v36).Cardinality(), (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int))
				r2 = ((_this).F25()).Minus((_709_v37).Cardinality())
			} else {
				var _710___mcc_h21 _dafny.Sequence = _source7.Get_().(D3_DC8).Cf15
				_ = _710___mcc_h21
				var _711_cf15 _dafny.Sequence = _710___mcc_h21
				_ = _711_cf15
				var _712_v38 _dafny.Array
				_ = _712_v38
				var _nwElement0_24 _dafny.Int = (_this).F25()
				_ = _nwElement0_24
				var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(7))
				_ = _nw113
				(_nw113).ArraySet1(_nwElement0_24, 0)
				(_nw113).ArraySet1((_this).F25(), 1)
				(_nw113).ArraySet1((_this).F25(), 2)
				(_nw113).ArraySet1((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), 3)
				(_nw113).ArraySet1(_this.F26, 4)
				(_nw113).ArraySet1(_this.F26, 5)
				(_nw113).ArraySet1(_dafny.IntOfInt64(244), 6)
				_712_v38 = _nw113
				var _713_v39 D0
				_ = _713_v39
				_713_v39 = Companion_D0_.Create_DC2_((_this).F18(), (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), _712_v38)
				(globalState).F14 = (_713_v39).Dtor_cf6()
				(globalState).F3 = _625_cf0
				var _714_v40 *C0
				_ = _714_v40
				var _nw114 *C0 = New_C0_()
				_ = _nw114
				_nw114.Ctor__()
				_714_v40 = _nw114
				var _715_v41 _dafny.Array
				_ = _715_v41
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw115
				_715_v41 = _nw115
				var _716_v42 _dafny.Map
				_ = _716_v42
				_716_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _715_v41)
				var _717_v43 _dafny.Map
				_ = _717_v43
				_717_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F26, _716_v42)
				var _718_v44 D5
				_ = _718_v44
				_718_v44 = Companion_D5_.Create_DC12_(_716_v42)
				_717_v43 = (_717_v43).Update((_this).F25(), (_718_v44).Dtor_cf22())
			}
		} else if _source5.Is_DC1() {
			var _719___mcc_h1 _dafny.Sequence = _source5.Get_().(D0_DC1).Cf1
			_ = _719___mcc_h1
			var _720___mcc_h2 _dafny.Int = _source5.Get_().(D0_DC1).Cf2
			_ = _720___mcc_h2
			var _721___mcc_h3 _dafny.Int = _source5.Get_().(D0_DC1).Cf3
			_ = _721___mcc_h3
			var _722___mcc_h4 _dafny.CodePoint = _source5.Get_().(D0_DC1).Cf4
			_ = _722___mcc_h4
			var _723_cf4 _dafny.CodePoint = _722___mcc_h4
			_ = _723_cf4
			var _724_cf3 _dafny.Int = _721___mcc_h3
			_ = _724_cf3
			var _725_cf2 _dafny.Int = _720___mcc_h2
			_ = _725_cf2
			var _726_cf1 _dafny.Sequence = _719___mcc_h1
			_ = _726_cf1
			_726_cf1 = _726_cf1
			var _727_v45 _dafny.Array
			_ = _727_v45
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_24
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) bool = func(_728_i12 _dafny.Int) bool {
					return (_this).F18()
				}
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw116 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw116).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw116).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_727_v45 = _nw116
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
			_ = _index108
			(_727_v45).ArraySet1((_this).F18(), (_index108).Int())
			var _729_v46 _dafny.MultiSet
			_ = _729_v46
			_729_v46 = _dafny.MultiSetOf(_619_v0, _619_v0)
			var _730_v47 _dafny.Map
			_ = _730_v47
			_730_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _729_v46)
			var _731_v48 _dafny.Sequence
			_ = _731_v48
			_731_v48 = _dafny.SeqOf(_729_v46, _729_v46)
			var _732_v49 _dafny.Array
			_ = _732_v49
			var _nwElement0_25 _dafny.MultiSet = _729_v46
			_ = _nwElement0_25
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(20))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_25, 0)
			(_nw117).ArraySet1(_729_v46, 1)
			(_nw117).ArraySet1((_729_v46).Update(_619_v0, Companion_Default___.Abs(_this.F26)), 2)
			(_nw117).ArraySet1(_729_v46, 3)
			(_nw117).ArraySet1((_729_v46).Union(_dafny.MultiSetOf(_619_v0, _619_v0, _619_v0)), 4)
			(_nw117).ArraySet1(_729_v46, 5)
			(_nw117).ArraySet1((func() _dafny.MultiSet {
				if (_730_v47).Contains((_this).F18()) {
					return (_730_v47).Get((_this).F18()).(_dafny.MultiSet)
				}
				return _729_v46
			})(), 6)
			(_nw117).ArraySet1((_729_v46).Difference(_729_v46), 7)
			(_nw117).ArraySet1(_729_v46, 8)
			(_nw117).ArraySet1(_729_v46, 9)
			(_nw117).ArraySet1((_729_v46).Intersection(_729_v46), 10)
			(_nw117).ArraySet1((_731_v48).Select((Companion_Default___.SafeIndex(_724_cf3, _dafny.IntOfUint32((_731_v48).Cardinality()))).Uint32()).(_dafny.MultiSet), 11)
			(_nw117).ArraySet1((_729_v46).Update(_619_v0, Companion_Default___.Abs(Companion_Default___.Fm4(_this.F26, globalState))), 12)
			(_nw117).ArraySet1(_729_v46, 13)
			(_nw117).ArraySet1(_dafny.MultiSetOf(_619_v0, _619_v0), 14)
			(_nw117).ArraySet1((_dafny.MultiSetOf(_619_v0)).Union(_729_v46), 15)
			(_nw117).ArraySet1(_729_v46, 16)
			(_nw117).ArraySet1(_729_v46, 17)
			(_nw117).ArraySet1(_729_v46, 18)
			(_nw117).ArraySet1(_729_v46, 19)
			_732_v49 = _nw117
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_732_v49), 0))
			_ = _index109
			(_732_v49).ArraySet1(_729_v46, (_index109).Int())
			var _733_v50 D0
			_ = _733_v50
			_733_v50 = Companion_D0_.Create_DC0_((_this).F18())
			var _734_v51 _dafny.Sequence
			_ = _734_v51
			_734_v51 = _dafny.SeqOf((_this).F18())
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
			_ = _index110
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_732_v49), 0))
			_ = _index111
			var _rhs95 _dafny.CodePoint = _723_cf4
			_ = _rhs95
			var _rhs96 bool = (_725_cf2).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F18()), (Companion_Default___.SafeIndex(_724_cf3, _dafny.IntOfUint32((_dafny.SeqOf((_this).F18())).Cardinality()))).Uint32(), (_733_v50).Dtor_cf0())).Cardinality())) < 0
			_ = _rhs96
			var _rhs97 _dafny.CodePoint = _723_cf4
			_ = _rhs97
			var _rhs98 bool = (_734_v51).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_729_v46).Contains(_619_v0) {
					return (_729_v46).Multiplicity(_619_v0)
				}
				return _724_cf3
			})(), _dafny.IntOfUint32((_734_v51).Cardinality()))).Uint32()).(bool)
			_ = _rhs98
			var _rhs99 _dafny.MultiSet = (_729_v46).Union(_dafny.MultiSetOf(_619_v0))
			_ = _rhs99
			var _lhs83 _dafny.Array = _727_v45
			_ = _lhs83
			var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
			_ = _lhs84
			var _lhs85 *GlobalState = globalState
			_ = _lhs85
			var _lhs86 _dafny.Array = _732_v49
			_ = _lhs86
			var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_732_v49), 0))
			_ = _lhs87
			_723_cf4 = _rhs95
			(_lhs83).ArraySet1(_rhs96, (_lhs84).Int())
			_723_cf4 = _rhs97
			_lhs85.F3 = _rhs98
			(_lhs86).ArraySet1(_rhs99, (_lhs87).Int())
			if (_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool) {
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_619_v0), 0))
				_ = _index112
				(_619_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(325))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_735_cf4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_736_i13 _dafny.Int) _dafny.CodePoint {
						return _735_cf4
					}
				})(_723_cf4))), _726_cf1)).Cardinality()), (_index112).Int())
				var _737_v52 _dafny.Set
				_ = _737_v52
				_737_v52 = _dafny.SetOf((_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool))
				var _738_v53 _dafny.Map
				_ = _738_v53
				_738_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_cf1, _724_cf3)
				var _739_v54 _dafny.Sequence
				_ = _739_v54
				_739_v54 = _dafny.SeqOf(_725_cf2, _725_cf2)
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
				_ = _index113
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_619_v0), 0))
				_ = _index115
				var _rhs100 bool = !(((_this).F25()).Cmp(((_737_v52).Cardinality()).Times((_this).F25())) != 0)
				_ = _rhs100
				var _rhs101 bool = false
				_ = _rhs101
				var _rhs102 bool = ((_738_v53).Cardinality()).Cmp((_dafny.IntOfUint32((_739_v54).Cardinality())).Minus(_dafny.IntOfInt64(536))) != 0
				_ = _rhs102
				var _rhs103 bool = ((_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool)) && (Companion_Default___.Fm13(globalState))
				_ = _rhs103
				var _rhs104 _dafny.Int = _dafny.IntOfInt64(647)
				_ = _rhs104
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				var _lhs89 _dafny.Array = _727_v45
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				var _lhs92 _dafny.Array = _727_v45
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))
				_ = _lhs93
				var _lhs94 _dafny.Array = _619_v0
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_619_v0), 0))
				_ = _lhs95
				_lhs88.F3 = _rhs100
				(_lhs89).ArraySet1(_rhs101, (_lhs90).Int())
				_lhs91.F9 = _rhs102
				(_lhs92).ArraySet1(_rhs103, (_lhs93).Int())
				(_lhs94).ArraySet1(_rhs104, (_lhs95).Int())
				var _740_v55 _dafny.MultiSet
				_ = _740_v55
				_740_v55 = _dafny.MultiSetOf(false, (_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool))
				var _741_v56 D2
				_ = _741_v56
				_741_v56 = Companion_D2_.Create_DC7_(Companion_D2_.Create_DC5_(_740_v55))
				var _742_v57 D2
				_ = _742_v57
				_742_v57 = Companion_D2_.Create_DC7_(_741_v56)
				_742_v57 = _742_v57
				var _743_v58 _dafny.Array
				_ = _743_v58
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_25
				var _nw118 _dafny.Array
				_ = _nw118
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw118 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Int = func(_744_i14 _dafny.Int) _dafny.Int {
						return (_744_i14).Minus((_this).F25())
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw118 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw118).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw118).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_743_v58 = _nw118
				var _745_v59 D0
				_ = _745_v59
				_745_v59 = Companion_D0_.Create_DC2_((_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool), _725_cf2, _743_v58)
				var _rhs105 bool = ((_745_v59).Dtor_cf6()).Cmp(_dafny.IntOfInt64(615)) >= 0
				_ = _rhs105
				var _rhs106 bool = (_this).F18()
				_ = _rhs106
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				_lhs96.F9 = _rhs105
				_lhs97.F3 = _rhs106
				(globalState).F3 = (_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool)
				var _746_v60 D2
				_ = _746_v60
				_746_v60 = Companion_D2_.Create_DC5_(_740_v55)
				var _747_v61 _dafny.Set
				_ = _747_v61
				_747_v61 = _dafny.SetOf(_746_v60, Companion_Default___.Fm15((_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool), _726_cf1, _726_cf1, globalState))
				(globalState).F3 = !((_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool)) || ((_747_v61).IsSubsetOf(_747_v61))
			} else {
				(globalState).F9 = (_725_cf2).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(824))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_748_cf4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_749_i15 _dafny.Int) _dafny.CodePoint {
						return _748_cf4
					}
				})(_723_cf4)))).Cardinality())) > 0
				var _750_v62 _dafny.Array
				_ = _750_v62
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_26
				var _nw119 _dafny.Array
				_ = _nw119
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw119 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Sequence = (func(_751_v51 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_752_i16 _dafny.Int) _dafny.Sequence {
							return (_751_v51)
						}
					})(_734_v51)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw119 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw119).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw119).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_750_v62 = _nw119
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_750_v62), 0))
				_ = _index116
				(_750_v62).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_734_v51, (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_734_v51).Cardinality()))).Uint32(), (_this).F18()), Companion_Default___.Fm16(globalState)), (_index116).Int())
				var _753_v63 _dafny.Sequence
				_ = _753_v63
				_753_v63 = _dafny.SeqOf(_725_cf2)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_750_v62), 0))
				_ = _index117
				(_750_v62).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(false), (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Uint32(), (_this).F18()))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_753_v63, _753_v63)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(false), (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Uint32(), (_this).F18())))).Cardinality()))).Uint32(), false), (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(false), (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Uint32(), (_this).F18()))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_753_v63, _753_v63)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Concatenate(_734_v51, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(false), (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Uint32(), (_this).F18())))).Cardinality()))).Uint32(), false)).Cardinality()))).Uint32(), (_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool)), (_index117).Int())
				var _754_v64 _dafny.Set
				_ = _754_v64
				_754_v64 = _dafny.SetOf(Companion_Default___.Fm17(globalState))
				var _755_v65 _dafny.Sequence
				_ = _755_v65
				_755_v65 = _dafny.SeqOf(!((_this).F18()), (_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool))
				_754_v64 = (func() _dafny.Set {
					if (_this).F18() {
						return _754_v64
					}
					return _dafny.SetOf(_755_v65, _755_v65)
				})()
				var _rhs107 bool = !(Companion_Default___.Fm13(globalState)) || ((_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool))
				_ = _rhs107
				var _rhs108 _dafny.Int = _this.F26
				_ = _rhs108
				var _rhs109 _dafny.Int = (_725_cf2).Minus(_724_cf3)
				_ = _rhs109
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				_lhs98.F9 = _rhs107
				_lhs99.F14 = _rhs108
				_lhs100.F14 = _rhs109
				(globalState).F14 = _724_cf3
			}
			(globalState).F3 = (_727_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_727_v45), 0))).Int()).(bool)
		} else {
			var _756___mcc_h5 bool = _source5.Get_().(D0_DC2).Cf5
			_ = _756___mcc_h5
			var _757___mcc_h6 _dafny.Int = _source5.Get_().(D0_DC2).Cf6
			_ = _757___mcc_h6
			var _758___mcc_h7 _dafny.Array = _source5.Get_().(D0_DC2).Cf7
			_ = _758___mcc_h7
			var _759_cf7 _dafny.Array = _758___mcc_h7
			_ = _759_cf7
			var _760_cf6 _dafny.Int = _757___mcc_h6
			_ = _760_cf6
			var _761_cf5 bool = _756___mcc_h5
			_ = _761_cf5
			var _762_v66 _dafny.Set
			_ = _762_v66
			_762_v66 = _dafny.SetOf((_this).F18())
			var _763_v67 _dafny.Sequence
			_ = _763_v67
			_763_v67 = _dafny.SeqOf((_762_v66).Cardinality(), (_this).F25(), _760_cf6, _760_cf6, _760_cf6)
			var _764_v68 _dafny.Set
			_ = _764_v68
			_764_v68 = _dafny.SetOf((_763_v67).Select((Companion_Default___.SafeIndex(_760_cf6, _dafny.IntOfUint32((_763_v67).Cardinality()))).Uint32()).(_dafny.Int))
			(globalState).F9 = (_764_v68).IsProperSubsetOf(_764_v68)
			var _765_v69 _dafny.CodePoint
			_ = _765_v69
			_765_v69 = _dafny.CodePoint('q')
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_759_cf7), 0))
			_ = _index118
			(_759_cf7).ArraySet1(((_this).F25()).Times((_dafny.Zero).Minus(_760_cf6)), (_index118).Int())
			var _766_v70 _dafny.MultiSet
			_ = _766_v70
			_766_v70 = _dafny.MultiSetOf((_this).F18())
			var _767_v71 D2
			_ = _767_v71
			_767_v71 = Companion_D2_.Create_DC5_(_766_v70)
			var _768_v72 D2
			_ = _768_v72
			_768_v72 = Companion_D2_.Create_DC5_((_767_v71).Dtor_cf12())
			var _769_v73 _dafny.Sequence
			_ = _769_v73
			_769_v73 = _dafny.SeqOf((_this).F18(), _761_cf5)
			var _770_v74 _dafny.Map
			_ = _770_v74
			_770_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_769_v73).Cardinality()), (_this).F25())
			var _771_v75 _dafny.Sequence
			_ = _771_v75
			_771_v75 = _dafny.UnicodeSeqOfUtf8Bytes("ayplhylwi")
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_759_cf7), 0))
			_ = _index119
			var _rhs110 _dafny.CodePoint = _dafny.CodePoint('c')
			_ = _rhs110
			var _rhs111 _dafny.Int = _dafny.IntOfInt64(-472)
			_ = _rhs111
			var _rhs112 D2 = _768_v72
			_ = _rhs112
			var _rhs113 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F26, _760_cf6)
			_ = _rhs113
			var _rhs114 _dafny.Int = ((_770_v74).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_771_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}(func(_772_i17 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})))).Cardinality()))
			_ = _rhs114
			var _lhs101 _dafny.Array = _759_cf7
			_ = _lhs101
			var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_759_cf7), 0))
			_ = _lhs102
			var _lhs103 *C4 = _this
			_ = _lhs103
			_765_v69 = _rhs110
			(_lhs101).ArraySet1(_rhs111, (_lhs102).Int())
			_767_v71 = _rhs112
			_lhs103.F26 = _rhs113
			r2 = _rhs114
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_759_cf7), 0))
			_ = _index120
			(_759_cf7).ArraySet1(_this.F26, (_index120).Int())
			var _773_v76 _dafny.Array
			_ = _773_v76
			var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
			_ = _nw120
			_773_v76 = _nw120
			var _774_v77 _dafny.Sequence
			_ = _774_v77
			_774_v77 = _dafny.SeqOf(_759_cf7)
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_773_v76), 0))
			_ = _index121
			(_773_v76).ArraySet1((func() _dafny.Sequence {
				if _761_cf5 {
					return _774_v77
				}
				return _774_v77
			})(), (_index121).Int())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_773_v76), 0))
			_ = _index122
			(_773_v76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_774_v77, _dafny.Companion_Sequence_.Update(_774_v77, (Companion_Default___.SafeIndex((_759_cf7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_759_cf7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_774_v77).Cardinality()))).Uint32(), _619_v0)), (_index122).Int())
		}
		var _775_v78 _dafny.Set
		_ = _775_v78
		_775_v78 = _dafny.SetOf((_this).F18())
		var _776_v79 _dafny.Sequence
		_ = _776_v79
		_776_v79 = _dafny.SeqOf(_this.F26, _this.F26, _dafny.IntOfInt64(224))
		if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_776_v79, _776_v79), ((_775_v78).Cardinality()).Times(_this.F26)) {
			(globalState).F9 = (_this).F18()
			var _777_v80 _dafny.Sequence
			_ = _777_v80
			_777_v80 = _dafny.SeqOf((_this).F18())
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_777_v80, _777_v80), _777_v80) {
				var _778_v81 _dafny.Array
				_ = _778_v81
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_27
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = func(_779_i18 _dafny.Int) bool {
						return (_this).F18()
					}
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw121 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw121).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw121).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_778_v81 = _nw121
				var _780_v82 T1
				_ = _780_v82
				var _nw122 *C3 = New_C3_()
				_ = _nw122
				_nw122.Ctor__(_778_v81, _dafny.SetOf(_778_v81, _778_v81), (_this).F18())
				_780_v82 = _nw122
				(_this).F26 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F25()), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F25(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _780_v82)).Cardinality())))
				var _781_v83 _dafny.CodePoint
				_ = _781_v83
				_781_v83 = _dafny.CodePoint('y')
				var _782_v84 _dafny.Sequence
				_ = _782_v84
				_782_v84 = _dafny.SeqOf(_781_v83, _781_v83)
				(globalState).F14 = (((_dafny.MultiSetOf(_781_v83)).Union(_dafny.MultiSetFromSeq(_782_v84))).Cardinality()).Plus((_this).F25())
				var _783_v85 _dafny.Set
				_ = _783_v85
				_783_v85 = _dafny.SetOf(_dafny.IntOfInt64(854), _this.F26)
				_783_v85 = _783_v85
				var _784_v86 _dafny.Map
				_ = _784_v86
				_784_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F26, _dafny.IntOfInt64(-870))
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_778_v81), 0))
				_ = _index123
				(_778_v81).ArraySet1(((func() _dafny.Int {
					if (_784_v86).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vcaginhv")).Cardinality())) {
						return (_784_v86).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vcaginhv")).Cardinality())).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-392)
				})()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_776_v79).Cardinality()))) != 0, (_index123).Int())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_778_v81), 0))
				_ = _index124
				(_778_v81).ArraySet1(((_this).F18()) == ((_this).F18()), (_index124).Int())
				var _785_v88 T0
				_ = _785_v88
				var _nw123 *C2 = New_C2_()
				_ = _nw123
				_nw123.Ctor__(_780_v82.F17(), !((_this).F18()))
				_785_v88 = _nw123
				var _786_v89 _dafny.Map
				_ = _786_v89
				_786_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_778_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_778_v81), 0))).Int()).(bool), _785_v88)
				var _787_v90 bool
				_ = _787_v90
				var _788_v91 bool
				_ = _788_v91
				var _789_v92 _dafny.Int
				_ = _789_v92
				var _out44 bool
				_ = _out44
				var _out45 bool
				_ = _out45
				var _out46 _dafny.Int
				_ = _out46
				_out44, _out45, _out46 = (_780_v82).M0((_this).F25(), (func() _dafny.Set {
					var _coll41 = _dafny.NewBuilder()
					_ = _coll41
					for _iter44 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-119))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}(func(_790_i19 _dafny.Int) _dafny.Int {
						return _this.F26
					}))).Elements()); ; {
						_compr_41, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _791_v87 _dafny.Int
						_791_v87 = interface{}(_compr_41).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-119))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg54 _dafny.Int) interface{} {
								return coer54(arg54)
							}
						}(func(_790_i19 _dafny.Int) _dafny.Int {
							return _this.F26
						})), _791_v87) {
							_coll41.Add((_791_v87).Plus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(33))).Uint32(), func(coer55 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
								return func(arg55 _dafny.Int) interface{} {
									return coer55(arg55)
								}
							}(func(_792_i20 _dafny.Int) D3 {
								return Companion_D3_.Create_DC9_()
							})))).Cardinality())))
						}
					}
					return _coll41.ToSet()
				}()).IsSubsetOf(_783_v85), ((func() _dafny.Map {
					if (_778_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_778_v81), 0))).Int()).(bool) {
						return _786_v89
					}
					return _786_v89
				})()).Cardinality(), globalState)
				_787_v90 = _out44
				_788_v91 = _out45
				_789_v92 = _out46
			} else {
				var _793_v93 *C2
				_ = _793_v93
				var _nw124 *C2 = New_C2_()
				_ = _nw124
				_nw124.Ctor__(_this.F17(), (_this).F18())
				_793_v93 = _nw124
				var _794_v94 D3
				_ = _794_v94
				_794_v94 = Companion_D3_.Create_DC9_()
				var _795_v95 _dafny.Map
				_ = _795_v95
				_795_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _794_v94)
				(globalState).F3 = !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_775_v78).Cardinality(), _794_v94)).Equals(_795_v95)
				var _796_v96 _dafny.Sequence
				_ = _796_v96
				_796_v96 = _dafny.UnicodeSeqOfUtf8Bytes("ntmfjd")
				var _797_v97 _dafny.Sequence
				_ = _797_v97
				_797_v97 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_796_v96, _796_v96))
				r2 = _dafny.IntOfUint32(((_797_v97).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_797_v97).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
				var _798_v98 _dafny.Array
				_ = _798_v98
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_28
				var _nw125 _dafny.Array
				_ = _nw125
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw125 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) bool = func(_799_i21 _dafny.Int) bool {
						return (_this).F18()
					}
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw125 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw125).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw125).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_798_v98 = _nw125
				var _800_v99 *C3
				_ = _800_v99
				var _nw126 *C3 = New_C3_()
				_ = _nw126
				_nw126.Ctor__(_798_v98, _this.F17(), (_this).F18())
				_800_v99 = _nw126
				var _801_v100 _dafny.Map
				_ = _801_v100
				_801_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F25()), _this.F26)
				_801_v100 = (_801_v100).Update(_this.F26, (_this).F25())
			}
			var _802_v101 _dafny.Array
			_ = _802_v101
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_29
			var _nw127 _dafny.Array
			_ = _nw127
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw127 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Sequence = func(_803_i22 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("wllppav")
				}
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw127 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw127).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw127).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_802_v101 = _nw127
			var _804_v102 _dafny.Sequence
			_ = _804_v102
			_804_v102 = _dafny.UnicodeSeqOfUtf8Bytes("gwvqrarpx")
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_802_v101), 0))
			_ = _index125
			(_802_v101).ArraySet1(_804_v102, (_index125).Int())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_802_v101), 0))
			_ = _index126
			(_802_v101).ArraySet1(_804_v102, (_index126).Int())
			_776_v79 = (func() _dafny.Sequence {
				if (_this).F18() {
					return _776_v79
				}
				return _776_v79
			})()
			var _805_v103 _dafny.CodePoint
			_ = _805_v103
			_805_v103 = _dafny.CodePoint('b')
			_805_v103 = _805_v103
		} else {
			(globalState).F3 = (_this).F18()
			var _806_v104 T0
			_ = _806_v104
			var _nw128 *C2 = New_C2_()
			_ = _nw128
			_nw128.Ctor__(_this.F17(), !((_this).F18()))
			_806_v104 = _nw128
			_806_v104 = _806_v104
			var _807_v105 _dafny.Array
			_ = _807_v105
			var _nw129 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw129
			_807_v105 = _nw129
			var _808_v106 *C3
			_ = _808_v106
			var _nw130 *C3 = New_C3_()
			_ = _nw130
			_nw130.Ctor__(_807_v105, _this.F17(), (_this).F18())
			_808_v106 = _nw130
			(globalState).F3 = (_this).F18()
			(globalState).F9 = Companion_Default___.Fm13(globalState)
		}
		var _809_v107 _dafny.Set
		_ = _809_v107
		_809_v107 = _dafny.SetOf((_this).F25())
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))
		_ = _index127
		(_619_v0).ArraySet1((_776_v79).Select((Companion_Default___.SafeIndex((_809_v107).Cardinality(), _dafny.IntOfUint32((_776_v79).Cardinality()))).Uint32()).(_dafny.Int), (_index127).Int())
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))
		_ = _index128
		(_619_v0).ArraySet1(_this.F26, (_index128).Int())
		var _810_i23 _dafny.Int
		_ = _810_i23
		_810_i23 = _dafny.Zero
		{
			for (_this).F18() {
				{
					if (_810_i23).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_810_i23 = (_810_i23).Plus(_dafny.One)
					var _811_v108 _dafny.Map
					_ = _811_v108
					_811_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_776_v79, _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("gsdlppl"), _dafny.UnicodeSeqOfUtf8Bytes("ghnl")))
					var _812_v109 _dafny.Sequence
					_ = _812_v109
					_812_v109 = _dafny.SeqOf(!((_this).F18()), (_this).F18())
					_811_v108 = (_811_v108).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-581)), _dafny.Companion_Sequence_.Update(_776_v79, (Companion_Default___.SafeIndex((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_776_v79).Cardinality()))).Uint32(), _dafny.IntOfUint32((_812_v109).Cardinality()))), (_this).F18())
					var _813_v110 _dafny.Map
					_ = _813_v110
					_813_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F25()), (_this).F18())
					var _814_v111 _dafny.Map
					_ = _814_v111
					_814_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int))
					_813_v110 = (_813_v110).Update((_814_v111).Merge(_814_v111), (_this).F18())
					var _815_v112 _dafny.Sequence
					_ = _815_v112
					_815_v112 = _dafny.UnicodeSeqOfUtf8Bytes("hn")
					(_this).F26 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_this).F25(), _this.F26), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(972))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg56 _dafny.Int) interface{} {
							return coer56(arg56)
						}
					}(func(_816_i24 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					})), _815_v112)).Cardinality()))
					var _817_v113 _dafny.Map
					_ = _817_v113
					_817_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(10), _815_v112)
					var _818_v114 _dafny.Sequence
					_ = _818_v114
					_818_v114 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_815_v112).Cardinality()), (_this).F25()))
					var _819_v115 _dafny.CodePoint
					_ = _819_v115
					_819_v115 = _dafny.CodePoint('c')
					_815_v112 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_815_v112, _dafny.Companion_Sequence_.Concatenate(_815_v112, (func() _dafny.Sequence {
						if (_817_v113).Contains((_776_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_818_v114).Cardinality()), _dafny.IntOfUint32((_776_v79).Cardinality()))).Uint32()).(_dafny.Int)) {
							return (_817_v113).Get((_776_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_818_v114).Cardinality()), _dafny.IntOfUint32((_776_v79).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Sequence)
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("ptbvpe")
					})())), (Companion_Default___.SafeIndex((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_815_v112, _dafny.Companion_Sequence_.Concatenate(_815_v112, (func() _dafny.Sequence {
						if (_817_v113).Contains((_776_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_818_v114).Cardinality()), _dafny.IntOfUint32((_776_v79).Cardinality()))).Uint32()).(_dafny.Int)) {
							return (_817_v113).Get((_776_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_818_v114).Cardinality()), _dafny.IntOfUint32((_776_v79).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Sequence)
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("ptbvpe")
					})()))).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
						if (_this).F18() {
							return _dafny.CodePoint('p')
						}
						return _819_v115
					})())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _820_i25 _dafny.Int
		_ = _820_i25
		_820_i25 = _dafny.Zero
		{
			for (_this.F26).Cmp((_this).F25()) > 0 {
				{
					if (_820_i25).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_820_i25 = (_820_i25).Plus(_dafny.One)
					r2 = (_dafny.Zero).Minus(Companion_Default___.Fm4(_this.F26, globalState))
					var _821_v116 _dafny.Array
					_ = _821_v116
					var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
					_ = _nw131
					_821_v116 = _nw131
					var _822_v117 _dafny.Array
					_ = _822_v117
					var _len0_30 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_30
					var _nw132 _dafny.Array
					_ = _nw132
					if _len0_30.Cmp(_dafny.Zero) == 0 {
						_nw132 = _dafny.NewArray(_len0_30)
					} else {
						var _init30 func(_dafny.Int) _dafny.Sequence = func(_823_i26 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("ekeegfkgk")
						}
						_ = _init30
						var _element0_30 = _init30(_dafny.Zero)
						_ = _element0_30
						_nw132 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
						(_nw132).ArraySet1(_element0_30, 0)
						var _nativeLen0_30 = (_len0_30).Int()
						_ = _nativeLen0_30
						for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
							(_nw132).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
						}
					}
					_822_v117 = _nw132
					var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_821_v116), 0))
					_ = _index129
					(_821_v116).ArraySet1(_822_v117, (_index129).Int())
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_821_v116), 0))
					_ = _index130
					(_821_v116).ArraySet1(_822_v117, (_index130).Int())
					var _824_v118 *C0
					_ = _824_v118
					var _nw133 *C0 = New_C0_()
					_ = _nw133
					_nw133.Ctor__()
					_824_v118 = _nw133
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))
					_ = _index131
					(_619_v0).ArraySet1(_this.F26, (_index131).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _hi2 _dafny.Int = (_this).F25()
		_ = _hi2
		for _825_i27 := _this.F26; _825_i27.Cmp(_hi2) < 0; _825_i27 = _825_i27.Plus(_dafny.One) {
			var _826_v119 _dafny.CodePoint
			_ = _826_v119
			_826_v119 = _dafny.CodePoint('d')
			_826_v119 = _826_v119
			var _827_v120 _dafny.Sequence
			_ = _827_v120
			_827_v120 = _dafny.SeqOf(_826_v119)
			var _828_v121 _dafny.Map
			_ = _828_v121
			_828_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), _this.F26)
			var _829_v122 D3
			_ = _829_v122
			_829_v122 = Companion_D3_.Create_DC10_(_827_v120, (_this).F18(), _dafny.UnicodeSeqOfUtf8Bytes("yivqbw"), _827_v120, (_828_v121).Cardinality())
			(globalState).F9 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_829_v122, (_this).F18())).Cardinality()).Cmp(_825_i27) != 0
			var _830_v123 _dafny.Map
			_ = _830_v123
			_830_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-912))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_831_v119 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_832_i28 _dafny.Int) _dafny.CodePoint {
					return _831_v119
				}
			})(_826_v119))))
			_830_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(103))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}((func(_833_v119 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_834_i29 _dafny.Int) _dafny.CodePoint {
					return _833_v119
				}
			})(_826_v119))))
			if (_this).F18() {
				var _835_v124 _dafny.Map
				_ = _835_v124
				_835_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F26).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F18(), (_this).F18()), (Companion_Default___.SafeIndex(_this.F26, _dafny.IntOfUint32((_dafny.SeqOf((_this).F18(), (_this).F18())).Cardinality()))).Uint32(), false)).Cardinality())) > 0, (_this).F25())
				_835_v124 = (_835_v124).Update(!((_this).F18()), (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int))
				var _836_v125 *C0
				_ = _836_v125
				var _nw134 *C0 = New_C0_()
				_ = _nw134
				_nw134.Ctor__()
				_836_v125 = _nw134
				var _837_v126 _dafny.Array
				_ = _837_v126
				var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(27))
				_ = _nw135
				_837_v126 = _nw135
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_837_v126), 0))
				_ = _index132
				(_837_v126).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}(func(_838_i30 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(262), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _this.F26, (_dafny.MultiSetOf((_this).F18())).Cardinality())).Cardinality(), _dafny.IntOfInt64(715))
				}))), (_index132).Int())
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_837_v126), 0))
				_ = _index133
				(_837_v126).ArraySet1((_dafny.MultiSetOf(_776_v79)).Difference(_dafny.MultiSetOf(_776_v79, _776_v79, _776_v79)), (_index133).Int())
				var _839_v127 *C0
				_ = _839_v127
				var _nw136 *C0 = New_C0_()
				_ = _nw136
				_nw136.Ctor__()
				_839_v127 = _nw136
				(globalState).F9 = (Companion_Default___.SafeDivisionInt(_this.F26, (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int))).Cmp((_dafny.IntOfInt64(-298)).Minus((_775_v78).Cardinality())) <= 0
			} else {
				(globalState).F9 = !((_this).F18())
				var _840_v128 *C1
				_ = _840_v128
				var _nw137 *C1 = New_C1_()
				_ = _nw137
				_nw137.Ctor__()
				_840_v128 = _nw137
				var _841_v129 _dafny.Map
				_ = _841_v129
				_841_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_827_v120).Cardinality()), (_this).F18())
				var _842_v130 _dafny.Map
				_ = _842_v130
				_842_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_825_i27, _841_v129)
				var _843_v131 _dafny.Map
				_ = _843_v131
				_843_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v130, (_this).F25())
				(globalState).F14 = (func() _dafny.Int {
					if (_843_v131).Contains(_842_v130) {
						return (_843_v131).Get(_842_v130).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-449)
				})()
				var _844_v132 _dafny.Array
				_ = _844_v132
				var _nwElement0_26 _dafny.Array = _619_v0
				_ = _nwElement0_26
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(6))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_26, 0)
				(_nw138).ArraySet1(_619_v0, 1)
				(_nw138).ArraySet1(_619_v0, 2)
				(_nw138).ArraySet1(_619_v0, 3)
				(_nw138).ArraySet1(_619_v0, 4)
				(_nw138).ArraySet1(_619_v0, 5)
				_844_v132 = _nw138
				var _845_v133 _dafny.Map
				_ = _845_v133
				_845_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_844_v132, _this.F26)
				var _846_v134 _dafny.Map
				_ = _846_v134
				_846_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _844_v132)
				_845_v133 = (_845_v133).Update((func() _dafny.Array {
					if (_846_v134).Contains((_840_v128).Fm5(_this.F26, _776_v79, globalState)) {
						return (_846_v134).Get((_840_v128).Fm5(_this.F26, _776_v79, globalState)).(_dafny.Array)
					}
					return _844_v132
				})(), (_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F18() {
						return _825_i27
					}
					return _this.F26
				})()))
				var _847_v135 D0
				_ = _847_v135
				_847_v135 = Companion_D0_.Create_DC2_((_this).F18(), (_this).F25(), _619_v0)
				var _848_v136 _dafny.MultiSet
				_ = _848_v136
				_848_v136 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(81))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_849_v119 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_850_i31 _dafny.Int) _dafny.CodePoint {
						return _849_v119
					}
				})(_826_v119))))
				var _851_v137 _dafny.MultiSet
				_ = _851_v137
				_851_v137 = _dafny.MultiSetOf(_825_i27)
				var _852_v138 _dafny.Array
				_ = _852_v138
				var _nwElement0_27 bool = (func() bool {
					if (_this).F18() {
						return (_this).F18()
					}
					return (_this).F18()
				})()
				_ = _nwElement0_27
				var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(14))
				_ = _nw139
				(_nw139).ArraySet1(_nwElement0_27, 0)
				(_nw139).ArraySet1(!((_this).F18()) || ((_this).F18()), 1)
				(_nw139).ArraySet1((_this).F18(), 2)
				(_nw139).ArraySet1(((_dafny.MultiSetOf(_825_i27)).Update((_this).F25(), Companion_Default___.Abs(_this.F26))).IsSubsetOf(Companion_Default___.Fm24((_847_v135).Dtor_cf5(), (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), globalState)), 3)
				(_nw139).ArraySet1((_this).F18(), 4)
				(_nw139).ArraySet1((_this).F18(), 5)
				(_nw139).ArraySet1(((func() _dafny.Int {
					if (_848_v136).Contains(_827_v120) {
						return (_848_v136).Multiplicity(_827_v120)
					}
					return (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)
				})()).Cmp(_dafny.IntOfUint32((_827_v120).Cardinality())) > 0, 6)
				(_nw139).ArraySet1((_this).F18(), 7)
				(_nw139).ArraySet1(((_this).F25()).Cmp((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)) > 0, 8)
				(_nw139).ArraySet1((_840_v128).Fm5(_this.F26, _776_v79, globalState), 9)
				(_nw139).ArraySet1(!((_this).F18()), 10)
				(_nw139).ArraySet1((_this).F18(), 11)
				(_nw139).ArraySet1((_this).F18(), 12)
				(_nw139).ArraySet1((_dafny.IntOfUint32((_827_v120).Cardinality())).Cmp((func() _dafny.Int {
					if (_851_v137).Contains(_this.F26) {
						return (_851_v137).Multiplicity(_this.F26)
					}
					return _825_i27
				})()) >= 0, 13)
				_852_v138 = _nw139
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_852_v138), 0))
				_ = _index134
				(_852_v138).ArraySet1((_this).F18(), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_852_v138), 0))
				_ = _index135
				var _rhs115 bool = (_this).F18()
				_ = _rhs115
				var _rhs116 _dafny.CodePoint = _826_v119
				_ = _rhs116
				var _lhs104 _dafny.Array = _852_v138
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_852_v138), 0))
				_ = _lhs105
				(_lhs104).ArraySet1(_rhs115, (_lhs105).Int())
				_826_v119 = _rhs116
			}
		}
		r0 = ((_this).F25()).Cmp(_dafny.IntOfInt64(730)) == 0
		var _853_v139 _dafny.Array
		_ = _853_v139
		var _nw140 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw140
		_853_v139 = _nw140
		var _854_v140 _dafny.Map
		_ = _854_v140
		_854_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_853_v139, _this.F26)
		r1 = _854_v140
		var _855_v141 _dafny.Map
		_ = _855_v141
		_855_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _809_v107)
		r2 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int), (_619_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_619_v0), 0))).Int()).(_dafny.Int)), ((func() _dafny.Set {
			if (_855_v141).Contains(!((_this).F18())) {
				return (_855_v141).Get(!((_this).F18())).(_dafny.Set)
			}
			return _809_v107
		})()).Cardinality())
		return r0, r1, r2
	}
}
func (_this *C4) F25() _dafny.Int {
	{
		return _this._f25
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f17 _dafny.Set
	_f19 _dafny.Array
	_f21 _dafny.Map
	_f20 _dafny.Int
	_f18 bool
	_f29 _dafny.Sequence
	_f30 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f21 = _dafny.EmptyMap
	_this._f20 = _dafny.Zero
	_this._f18 = false
	_this._f29 = _dafny.EmptySeq
	_this._f30 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C5{}
var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F17() _dafny.Set {
	return _this._f17
}
func (_this *C5) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C5) F19() _dafny.Array {
	return _this._f19
}
func (_this *C5) F19_set_(value _dafny.Array) {
	_this._f19 = value
}
func (_this *C5) F21() _dafny.Map {
	return _this._f21
}
func (_this *C5) F21_set_(value _dafny.Map) {
	_this._f21 = value
}
func (_this *C5) F20() _dafny.Int {
	return _this._f20
}
func (_this *C5) F18() bool {
	return _this._f18
}
func (_this *C5) Ctor__(f29 _dafny.Sequence, f30 bool, f20 _dafny.Int, f21 _dafny.Map, f19 _dafny.Array, f17 _dafny.Set, f18 bool) {
	{
		(_this)._f29 = f29
		(_this)._f30 = f30
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C5) Fm1(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		return ((_dafny.MultiSetOf((_dafny.SetOf(!((_this).F30()), (_this).F30())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F20(), (_this).F20())).Cardinality()), (_this).F20(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.UnicodeSeqOfUtf8Bytes("lbbbe"))).Cardinality(), (_this).F20())).Cardinality()).Cmp((func() _dafny.Int {
			if (_this).F18() {
				return (_this).F20()
			}
			return (_this).F20()
		})()) == 0
	}
}
func (_this *C5) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		var _source8 D1 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf((_this).F20()))
		_ = _source8
		if _source8.Is_DC4() {
			var _856___mcc_h0 bool = _source8.Get_().(D1_DC4).Cf9
			_ = _856___mcc_h0
			var _857___mcc_h1 _dafny.Set = _source8.Get_().(D1_DC4).Cf10
			_ = _857___mcc_h1
			var _858___mcc_h2 _dafny.Int = _source8.Get_().(D1_DC4).Cf11
			_ = _858___mcc_h2
			var _859_cf11 _dafny.Int = _858___mcc_h2
			_ = _859_cf11
			var _860_cf10 _dafny.Set = _857___mcc_h1
			_ = _860_cf10
			var _861_cf9 bool = _856___mcc_h0
			_ = _861_cf9
			return _dafny.IntOfInt64(902)
		} else {
			var _862___mcc_h3 _dafny.MultiSet = _source8.Get_().(D1_DC3).Cf8
			_ = _862___mcc_h3
			var _863_cf8 _dafny.MultiSet = _862___mcc_h3
			_ = _863_cf8
			return (_dafny.Zero).Minus((_this).F20())
		}
	}
}
func (_this *C5) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		r1 = (_this).F18()
		var _864_v0 _dafny.Array
		_ = _864_v0
		var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw141
		_864_v0 = _nw141
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_864_v0), 0))); ; {
			_guard_loop_3, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _865_i0 _dafny.Int
			_865_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_865_i0).Sign() != -1) && ((_865_i0).Cmp(_dafny.ArrayLen((_864_v0), 0)) < 0)) {
				(_864_v0).ArraySet1((_865_i0).Times(((_this).F20()).Minus(p0)), (_865_i0).Int())
			}
		}
		var _866_v1 _dafny.Set
		_ = _866_v1
		_866_v1 = _dafny.SetOf(true, p1)
		var _867_i1 _dafny.Int
		_ = _867_i1
		_867_i1 = _dafny.Zero
		{
			for ((_866_v1).Cardinality()).Cmp((_this).F20()) < 0 {
				{
					if (_867_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_867_i1 = (_867_i1).Plus(_dafny.One)
					var _arr6 _dafny.Array = _this.F19()
					_ = _arr6
					var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_this.F19()), 0))
					_ = _index136
					(_arr6).ArraySet1((_this).F18(), (_index136).Int())
					var _868_v2 _dafny.CodePoint
					_ = _868_v2
					_868_v2 = _dafny.CodePoint('v')
					var _arr7 _dafny.Array = _this.F19()
					_ = _arr7
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_this.F19()), 0))
					_ = _index137
					(_arr7).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_this).F18(), p1, !(p1)), false), (_index137).Int())
					var _869_v3 _dafny.Sequence
					_ = _869_v3
					_869_v3 = _dafny.SeqOf((_this).F18())
					var _arr8 _dafny.Array = _this.F19()
					_ = _arr8
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_this.F19()), 0))
					_ = _index138
					var _arr9 _dafny.Array = _this.F19()
					_ = _arr9
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_this.F19()), 0))
					_ = _index139
					var _rhs117 _dafny.Int = _dafny.IntOfUint32(((_this).F29()).Cardinality())
					_ = _rhs117
					var _rhs118 bool = (func() bool {
						if false {
							return (_this).F30()
						}
						return true
					})()
					_ = _rhs118
					var _rhs119 _dafny.CodePoint = _868_v2
					_ = _rhs119
					var _rhs120 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_869_v3, _869_v3))
					_ = _rhs120
					var _lhs106 *GlobalState = globalState
					_ = _lhs106
					var _lhs107 _dafny.Array = _this.F19()
					_ = _lhs107
					var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_this.F19()), 0))
					_ = _lhs108
					var _lhs109 _dafny.Array = _this.F19()
					_ = _lhs109
					var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_this.F19()), 0))
					_ = _lhs110
					_lhs106.F14 = _rhs117
					(_lhs107).ArraySet1(_rhs118, (_lhs108).Int())
					_868_v2 = _rhs119
					(_lhs109).ArraySet1(_rhs120, (_lhs110).Int())
					var _870_v4 _dafny.Array
					_ = _870_v4
					var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
					_ = _nw142
					_870_v4 = _nw142
					var _871_v5 _dafny.Map
					_ = _871_v5
					_871_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_868_v2, _870_v4)
					var _872_v6 _dafny.Sequence
					_ = _872_v6
					_872_v6 = _dafny.SeqOf(_870_v4, _870_v4)
					var _873_v7 _dafny.Map
					_ = _873_v7
					_873_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_872_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.IntOfUint32((_872_v6).Cardinality()))).Uint32()).(_dafny.Array))
					_871_v5 = (_871_v5).Update(_dafny.CodePoint('g'), (func() _dafny.Array {
						if (_873_v7).Contains((_this).F29()) {
							return (_873_v7).Get((_this).F29()).(_dafny.Array)
						}
						return _870_v4
					})())
					var _874_v8 _dafny.Sequence
					_ = _874_v8
					_874_v8 = _dafny.SeqOf(p0)
					var _875_v9 D1
					_ = _875_v9
					_875_v9 = Companion_D1_.Create_DC3_(_dafny.MultiSetFromSeq(_874_v8))
					_875_v9 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf((_this).F20()))
					var _876_v10 _dafny.Array
					_ = _876_v10
					var _len0_31 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_31
					var _nw143 _dafny.Array
					_ = _nw143
					if _len0_31.Cmp(_dafny.Zero) == 0 {
						_nw143 = _dafny.NewArray(_len0_31)
					} else {
						var _init31 func(_dafny.Int) bool = func(_877_i2 _dafny.Int) bool {
							return (_this).F18()
						}
						_ = _init31
						var _element0_31 = _init31(_dafny.Zero)
						_ = _element0_31
						_nw143 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
						(_nw143).ArraySet1(_element0_31, 0)
						var _nativeLen0_31 = (_len0_31).Int()
						_ = _nativeLen0_31
						for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
							(_nw143).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
						}
					}
					_876_v10 = _nw143
					_876_v10 = _876_v10
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _878_v11 _dafny.Array
		_ = _878_v11
		var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
		_ = _nw144
		_878_v11 = _nw144
		var _879_v12 _dafny.Map
		_ = _879_v12
		_879_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F18())
		var _880_v13 _dafny.Set
		_ = _880_v13
		_880_v13 = _dafny.SetOf(_879_v12, _879_v12)
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_878_v11), 0))
		_ = _index140
		(_878_v11).ArraySet1(_880_v13, (_index140).Int())
		var _881_v14 _dafny.Set
		_ = _881_v14
		_881_v14 = _dafny.SetOf((_this).F29())
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_878_v11), 0))
		_ = _index141
		(_878_v11).ArraySet1((func() _dafny.Set {
			if (_881_v14).IsProperSubsetOf(_881_v14) {
				return _880_v13
			}
			return _880_v13
		})(), (_index141).Int())
		var _arr10 _dafny.Array = _this.F19()
		_ = _arr10
		var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index142
		(_arr10).ArraySet1((_this).F30(), (_index142).Int())
		var _882_v15 _dafny.Sequence
		_ = _882_v15
		_882_v15 = _dafny.SeqOf((_this).F30())
		var _883_v16 _dafny.Sequence
		_ = _883_v16
		_883_v16 = _882_v15
		var _arr11 _dafny.Array = _this.F19()
		_ = _arr11
		var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index143
		(_arr11).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate((_883_v16), _882_v15), (func() _dafny.Sequence {
			if (_this).F30() {
				return _882_v15
			}
			return _882_v15
		})()), (_index143).Int())
		var _884_v17 _dafny.Map
		_ = _884_v17
		_884_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}(func(_885_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F29()).Cardinality())))
		var _886_v18 _dafny.Map
		_ = _886_v18
		_886_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F18())
		var _hi3 _dafny.Int = (_this).F20()
		_ = _hi3
		for _887_i3 := (func() _dafny.Int {
			if (_884_v17).Contains((_886_v18).Cardinality()) {
				return (_884_v17).Get((_886_v18).Cardinality()).(_dafny.Int)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p1)).Cardinality()
		})(); _887_i3.Cmp(_hi3) < 0; _887_i3 = _887_i3.Plus(_dafny.One) {
			var _888_v19 _dafny.Array
			_ = _888_v19
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(26))
			_ = _nw145
			_888_v19 = _nw145
			var _889_v20 D3
			_ = _889_v20
			_889_v20 = Companion_D3_.Create_DC10_((_this).F29(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_this).F29(), (_this).F29(), (_884_v17).Cardinality())
			_888_v19 = (func() _dafny.Array {
				if ((_889_v20).Dtor_cf20()).Cmp(p0) != 0 {
					return _888_v19
				}
				return (func() _dafny.Array {
					if p1 {
						return _888_v19
					}
					return _888_v19
				})()
			})()
			var _890_v22 _dafny.Map
			_ = _890_v22
			_890_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll42 = _dafny.NewMapBuilder()
				_ = _coll42
				for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(500), _dafny.IntOfInt64(409))); ; {
					_compr_42, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _891_v21 _dafny.Int
					_891_v21 = interface{}(_compr_42).(_dafny.Int)
					if ((_dafny.IntOfInt64(500)).Cmp(_891_v21) <= 0) && ((_891_v21).Cmp(_dafny.IntOfInt64(409)) < 0) {
						_coll42.Add((_891_v21).Minus(_887_i3), _dafny.IntOfInt64(147))
					}
				}
				return _coll42.ToMap()
			}(), p2)
			var _892_v23 _dafny.MultiSet
			_ = _892_v23
			_892_v23 = _dafny.MultiSetOf((_this).Fm1(_890_v22, globalState))
			(globalState).F3 = !((_892_v23).Contains(!_dafny.Companion_Sequence_.Contains(_882_v15, (_this).F30())))
			var _arr12 _dafny.Array = _this.F19()
			_ = _arr12
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index144
			var _rhs121 bool = (_this).F30()
			_ = _rhs121
			var _rhs122 _dafny.Int = _dafny.IntOfInt64(652)
			_ = _rhs122
			var _rhs123 _dafny.Int = p0
			_ = _rhs123
			var _lhs111 _dafny.Array = _this.F19()
			_ = _lhs111
			var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_this.F19()), 0))
			_ = _lhs112
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			(_lhs111).ArraySet1(_rhs121, (_lhs112).Int())
			_lhs113.F14 = _rhs122
			_lhs114.F14 = _rhs123
			var _893_v24 _dafny.Sequence
			_ = _893_v24
			_893_v24 = _dafny.SeqOf(_892_v23, _892_v23, _892_v23)
			_892_v23 = (_893_v24).Select((Companion_Default___.SafeIndex(((_866_v1).Cardinality()).Plus((_866_v1).Cardinality()), _dafny.IntOfUint32((_893_v24).Cardinality()))).Uint32()).(_dafny.MultiSet)
		}
		r0 = (_this).F30()
		var _894_v25 _dafny.Set
		_ = _894_v25
		_894_v25 = _dafny.SetOf((_this).F20())
		var _895_v26 _dafny.Array
		_ = _895_v26
		var _nw146 _dafny.Array = _dafny.NewArrayWithValue(Companion_D11_.Default(), _dafny.IntOfInt64(29))
		_ = _nw146
		_895_v26 = _nw146
		var _896_v27 _dafny.Map
		_ = _896_v27
		_896_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v26, _894_v25)
		var _897_v28 _dafny.MultiSet
		_ = _897_v28
		_897_v28 = _dafny.MultiSetOf((_this).F30(), (_this).F30())
		var _898_v29 _dafny.Map
		_ = _898_v29
		_898_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
		var _899_v30 _dafny.Sequence
		_ = _899_v30
		_899_v30 = _dafny.SeqOf((_this).F20())
		r1 = (func() bool {
			if !((_894_v25).IsDisjointFrom((func() _dafny.Set {
				if (_896_v27).Contains(_895_v26) {
					return (_896_v27).Get(_895_v26).(_dafny.Set)
				}
				return _dafny.SetOf(p0, p2, p2, (func() _dafny.Int {
					if (_897_v28).Contains((_this).F30()) {
						return (_897_v28).Multiplicity((_this).F30())
					}
					return p2
				})(), p0)
			})())) {
				return _dafny.Companion_Sequence_.Contains((_this).F29(), Companion_Default___.Fm33(p2, (func() bool {
					if (_898_v29).Contains((_this).F30()) {
						return (_898_v29).Get((_this).F30()).(bool)
					}
					return (_this).F18()
				})(), globalState))
			}
			return _dafny.Companion_Sequence_.Contains(_899_v30, _dafny.IntOfInt64(409))
		})()
		r2 = (_this).F20()
		return r0, r1, r2
	}
}
func (_this *C5) F29() _dafny.Sequence {
	{
		return _this._f29
	}
}
func (_this *C5) F30() bool {
	{
		return _this._f30
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f17 _dafny.Set
	_f19 _dafny.Array
	_f21 _dafny.Map
	_f20 _dafny.Int
	_f18 bool
	_f28 _dafny.Int
	_f27 _dafny.Map
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f21 = _dafny.EmptyMap
	_this._f20 = _dafny.Zero
	_this._f18 = false
	_this._f28 = _dafny.Zero
	_this._f27 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C6{}
var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F17() _dafny.Set {
	return _this._f17
}
func (_this *C6) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C6) F19() _dafny.Array {
	return _this._f19
}
func (_this *C6) F19_set_(value _dafny.Array) {
	_this._f19 = value
}
func (_this *C6) F21() _dafny.Map {
	return _this._f21
}
func (_this *C6) F21_set_(value _dafny.Map) {
	_this._f21 = value
}
func (_this *C6) F20() _dafny.Int {
	return _this._f20
}
func (_this *C6) F18() bool {
	return _this._f18
}
func (_this *C6) Ctor__(f27 _dafny.Map, f28 _dafny.Int, f20 _dafny.Int, f21 _dafny.Map, f19 _dafny.Array, f17 _dafny.Set, f18 bool) {
	{
		(_this)._f27 = f27
		(_this)._f28 = f28
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C6) Fm1(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		return !(true)
	}
}
func (_this *C6) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F28()
	}
}
func (_this *C6) Fm27(globalState *GlobalState) bool {
	{
		return (_this).F18()
	}
}
func (_this *C6) Fm28(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F20()
	}
}
func (_this *C6) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _hi4 _dafny.Int = (_this).F28()
		_ = _hi4
		for _900_i0 := (_dafny.Zero).Minus((func() _dafny.Int {
			if p1 {
				return p2
			}
			return p0
		})()); _900_i0.Cmp(_hi4) < 0; _900_i0 = _900_i0.Plus(_dafny.One) {
			var _901_v0 _dafny.Sequence
			_ = _901_v0
			_901_v0 = _dafny.SeqOf(p2, p2, p2)
			var _902_v1 _dafny.Map
			_ = _902_v1
			_902_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_this).F28()).Cmp(_900_i0) != 0), _dafny.Companion_Sequence_.Concatenate(_901_v0, _901_v0))
			_902_v1 = (_902_v1).Merge(_902_v1)
			var _903_v2 _dafny.Sequence
			_ = _903_v2
			_903_v2 = _901_v0
			var _904_v3 _dafny.Sequence
			_ = _904_v3
			_904_v3 = _dafny.SeqOf((_903_v2), _901_v0)
			var _905_v4 _dafny.Sequence
			_ = _905_v4
			_905_v4 = _dafny.SeqOf(p1)
			var _rhs124 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_900_i0, (_this).F28()), (_dafny.IntOfUint32((_904_v3).Cardinality())).Minus(_dafny.IntOfUint32((_905_v4).Cardinality())))
			_ = _rhs124
			var _rhs125 _dafny.Int = _dafny.IntOfUint32((_901_v0).Cardinality())
			_ = _rhs125
			r2 = _rhs124
			r2 = _rhs125
			var _906_v5 D0
			_ = _906_v5
			_906_v5 = Companion_D0_.Create_DC0_((_this).F18())
			(globalState).F3 = (_906_v5).Dtor_cf0()
			var _907_v6 _dafny.Map
			_ = _907_v6
			_907_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _908_v7 _dafny.Sequence
			_ = _908_v7
			_908_v7 = _dafny.SeqOf(_907_v6, _907_v6, (_907_v6).Update(p1, p1))
			var _source9 D3 = Companion_D3_.Create_DC8_(_908_v7)
			_ = _source9
			if _source9.Is_DC9() {
				var _909_v8 _dafny.Map
				_ = _909_v8
				_909_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29(globalState), ((_this).F20()).Times(_dafny.IntOfInt64(73)))
				_909_v8 = (_909_v8).Update(_dafny.MultiSetOf(true, p1), (_this).F20())
				var _910_v9 _dafny.Map
				_ = _910_v9
				_910_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_905_v4, Companion_Default___.Fm30((_this).F18(), globalState))
				(globalState).F14 = (((_this).Fm28((_this).F20(), _dafny.IntOfInt64(784), _910_v9, globalState)).Plus((_this).F20())).Plus((_this).F20())
				var _911_v10 *C2
				_ = _911_v10
				var _nw147 *C2 = New_C2_()
				_ = _nw147
				_nw147.Ctor__(_this.F17(), (_this).F18())
				_911_v10 = _nw147
				var _912_v11 _dafny.Sequence
				_ = _912_v11
				_912_v11 = _dafny.SeqOf(_911_v10)
				(globalState).F9 = (_dafny.IntOfUint32((_912_v11).Cardinality())).Cmp(((_dafny.Zero).Minus(p2)).Times((_this).Fm28((_this).Fm28(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yrjchpu")).Cardinality()), _910_v9, globalState), (_this).F28(), _910_v9, globalState))) >= 0
				var _913_v12 _dafny.CodePoint
				_ = _913_v12
				_913_v12 = _dafny.CodePoint('i')
				var _914_v13 _dafny.Sequence
				_ = _914_v13
				_914_v13 = _dafny.UnicodeSeqOfUtf8Bytes("jovhhqucn")
				var _915_v14 _dafny.Sequence
				_ = _915_v14
				_915_v14 = _dafny.SeqOf(_913_v12, _913_v12, (_914_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.IntOfUint32((_914_v13).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_914_v13 = _914_v13
			} else if _source9.Is_DC10() {
				var _916___mcc_h0 _dafny.Sequence = _source9.Get_().(D3_DC10).Cf16
				_ = _916___mcc_h0
				var _917___mcc_h1 bool = _source9.Get_().(D3_DC10).Cf17
				_ = _917___mcc_h1
				var _918___mcc_h2 _dafny.Sequence = _source9.Get_().(D3_DC10).Cf18
				_ = _918___mcc_h2
				var _919___mcc_h3 _dafny.Sequence = _source9.Get_().(D3_DC10).Cf19
				_ = _919___mcc_h3
				var _920___mcc_h4 _dafny.Int = _source9.Get_().(D3_DC10).Cf20
				_ = _920___mcc_h4
				var _921_cf20 _dafny.Int = _920___mcc_h4
				_ = _921_cf20
				var _922_cf19 _dafny.Sequence = _919___mcc_h3
				_ = _922_cf19
				var _923_cf18 _dafny.Sequence = _918___mcc_h2
				_ = _923_cf18
				var _924_cf17 bool = _917___mcc_h1
				_ = _924_cf17
				var _925_cf16 _dafny.Sequence = _916___mcc_h0
				_ = _925_cf16
				var _926_v15 *C1
				_ = _926_v15
				var _nw148 *C1 = New_C1_()
				_ = _nw148
				_nw148.Ctor__()
				_926_v15 = _nw148
				(globalState).F9 = true
				var _927_v16 _dafny.Set
				_ = _927_v16
				_927_v16 = _dafny.SetOf((_this).F20())
				var _arr13 _dafny.Array = _this.F19()
				_ = _arr13
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index145
				(_arr13).ArraySet1((_927_v16).IsDisjointFrom(func() _dafny.Set {
					var _coll43 = _dafny.NewBuilder()
					_ = _coll43
					for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-267), _dafny.IntOfInt64(467))); ; {
						_compr_43, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _928_v17 _dafny.Int
						_928_v17 = interface{}(_compr_43).(_dafny.Int)
						if ((_dafny.IntOfInt64(-267)).Cmp(_928_v17) <= 0) && ((_928_v17).Cmp(_dafny.IntOfInt64(467)) < 0) {
							_coll43.Add(Companion_Default___.SafeDivisionInt(_928_v17, (_this).F20()))
						}
					}
					return _coll43.ToSet()
				}()), (_index145).Int())
				var _arr14 _dafny.Array = _this.F19()
				_ = _arr14
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index146
				(_arr14).ArraySet1(_924_cf17, (_index146).Int())
				var _929_v18 _dafny.Map
				_ = _929_v18
				_929_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_900_i0, p0)
				var _930_v19 _dafny.Map
				_ = _930_v19
				_930_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_929_v18, p2)
				(globalState).F3 = (_this).Fm1(_930_v19, globalState)
			} else {
				var _931___mcc_h5 _dafny.Sequence = _source9.Get_().(D3_DC8).Cf15
				_ = _931___mcc_h5
				var _932_cf15 _dafny.Sequence = _931___mcc_h5
				_ = _932_cf15
				var _933_v20 _dafny.CodePoint
				_ = _933_v20
				_933_v20 = _dafny.CodePoint('s')
				var _934_v21 _dafny.Set
				_ = _934_v21
				_934_v21 = _dafny.SetOf(_933_v20)
				var _935_v22 _dafny.Sequence
				_ = _935_v22
				_935_v22 = _dafny.SeqOf(Companion_D3_.Create_DC9_(), Companion_Default___.Fm31(globalState))
				var _936_v23 _dafny.MultiSet
				_ = _936_v23
				_936_v23 = _dafny.MultiSetOf((Companion_D5_.Create_DC13_((_this).F18(), (_this).F28(), _935_v22)).Dtor_cf24())
				var _rhs126 bool = p1
				_ = _rhs126
				var _rhs127 bool = (_934_v21).IsSubsetOf(_934_v21)
				_ = _rhs127
				var _rhs128 bool = !(p1)
				_ = _rhs128
				var _rhs129 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_936_v23).Contains(_900_i0) {
						return (_936_v23).Multiplicity(_900_i0)
					}
					return (_this).F20()
				})(), (_900_i0).Times(p0))
				_ = _rhs129
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				var _lhs117 *GlobalState = globalState
				_ = _lhs117
				_lhs115.F9 = _rhs126
				r1 = _rhs127
				_lhs116.F3 = _rhs128
				_lhs117.F14 = _rhs129
				var _937_v24 T1
				_ = _937_v24
				var _nw149 *C3 = New_C3_()
				_ = _nw149
				_nw149.Ctor__(_this.F19(), _this.F17(), p1)
				_937_v24 = _nw149
				var _938_v25 _dafny.Set
				_ = _938_v25
				_938_v25 = _dafny.SetOf(_937_v24, _937_v24)
				_938_v25 = _938_v25
				var _939_v26 D9
				_ = _939_v26
				_939_v26 = Companion_D9_.Create_DC19_(Companion_Default___.Fm32(globalState))
				r2 = (_this).Fm28((_this).F20(), ((_this).F28()).Minus(_dafny.IntOfInt64(144)), (_939_v26).Dtor_cf35(), globalState)
				var _940_v27 _dafny.Set
				_ = _940_v27
				_940_v27 = _dafny.SetOf(p0, _dafny.IntOfInt64(-324), p0)
				var _941_v28 D10
				_ = _941_v28
				_941_v28 = Companion_D10_.Create_DC22_(_940_v27)
				_940_v27 = (((_941_v28).Dtor_cf42()).Union(_940_v27)).Intersection(func() _dafny.Set {
					var _coll44 = _dafny.NewBuilder()
					_ = _coll44
					for _iter48 := _dafny.Iterate((_940_v27).Elements()); ; {
						_compr_44, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _942_v29 _dafny.Int
						_942_v29 = interface{}(_compr_44).(_dafny.Int)
						if (_940_v27).Contains(_942_v29) {
							_coll44.Add((_942_v29).Plus(_dafny.IntOfUint32((_dafny.SeqOf(!(true), !(false))).Cardinality())))
						}
					}
					return _coll44.ToSet()
				}())
			}
		}
		var _943_i1 _dafny.Int
		_ = _943_i1
		_943_i1 = _dafny.Zero
		{
			for p1 {
				{
					if (_943_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_943_i1 = (_943_i1).Plus(_dafny.One)
					var _944_v30 *C1
					_ = _944_v30
					var _nw150 *C1 = New_C1_()
					_ = _nw150
					_nw150.Ctor__()
					_944_v30 = _nw150
					var _945_v31 _dafny.Sequence
					_ = _945_v31
					_945_v31 = _dafny.UnicodeSeqOfUtf8Bytes("ert")
					var _946_v32 _dafny.Map
					_ = _946_v32
					_946_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(118), (_this).F28())
					var _947_v33 _dafny.CodePoint
					_ = _947_v33
					_947_v33 = _dafny.CodePoint('c')
					var _948_v34 _dafny.Array
					_ = _948_v34
					var _nwElement0_28 _dafny.Int = p0
					_ = _nwElement0_28
					var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(26))
					_ = _nw151
					(_nw151).ArraySet1(_nwElement0_28, 0)
					(_nw151).ArraySet1((_this).F20(), 1)
					(_nw151).ArraySet1((_this).F28(), 2)
					(_nw151).ArraySet1((_this).F20(), 3)
					(_nw151).ArraySet1(p2, 4)
					(_nw151).ArraySet1(_dafny.IntOfInt64(887), 5)
					(_nw151).ArraySet1(_dafny.IntOfUint32((_945_v31).Cardinality()), 6)
					(_nw151).ArraySet1(((_this).F27()).Cardinality(), 7)
					(_nw151).ArraySet1(p2, 8)
					(_nw151).ArraySet1((_this).F28(), 9)
					(_nw151).ArraySet1(_dafny.IntOfInt64(548), 10)
					(_nw151).ArraySet1(p2, 11)
					(_nw151).ArraySet1(p2, 12)
					(_nw151).ArraySet1((func() _dafny.Int {
						if ((_this).F27()).Contains((_this).F18()) {
							return ((_this).F27()).Get((_this).F18()).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pkdr")).Cardinality())
					})(), 13)
					(_nw151).ArraySet1((_this).F28(), 14)
					(_nw151).ArraySet1((_946_v32).Cardinality(), 15)
					(_nw151).ArraySet1(p0, 16)
					(_nw151).ArraySet1((_this).F28(), 17)
					(_nw151).ArraySet1(p2, 18)
					(_nw151).ArraySet1((_this).F20(), 19)
					(_nw151).ArraySet1((_this).Fm0((_this).F20(), _947_v33, globalState), 20)
					(_nw151).ArraySet1((_this).F28(), 21)
					(_nw151).ArraySet1((_this).F20(), 22)
					(_nw151).ArraySet1((_this).F20(), 23)
					(_nw151).ArraySet1(p0, 24)
					(_nw151).ArraySet1((_this).F20(), 25)
					_948_v34 = _nw151
					var _949_v35 _dafny.Array
					_ = _949_v35
					var _nwElement0_29 _dafny.Array = _948_v34
					_ = _nwElement0_29
					var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(10))
					_ = _nw152
					(_nw152).ArraySet1(_nwElement0_29, 0)
					(_nw152).ArraySet1(_948_v34, 1)
					(_nw152).ArraySet1(_948_v34, 2)
					(_nw152).ArraySet1(_948_v34, 3)
					(_nw152).ArraySet1(_948_v34, 4)
					(_nw152).ArraySet1(_948_v34, 5)
					(_nw152).ArraySet1(_948_v34, 6)
					(_nw152).ArraySet1(_948_v34, 7)
					(_nw152).ArraySet1(_948_v34, 8)
					(_nw152).ArraySet1((func() _dafny.Array {
						if false {
							return _948_v34
						}
						return _948_v34
					})(), 9)
					_949_v35 = _nw152
					_949_v35 = _949_v35
					var _950_v36 _dafny.Map
					_ = _950_v36
					_950_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
					_950_v36 = (_950_v36).Update(_dafny.IntOfUint32((_945_v31).Cardinality()), true)
					var _951_v37 _dafny.Sequence
					_ = _951_v37
					_951_v37 = _dafny.SeqOf(true, true)
					var _952_v38 _dafny.Set
					_ = _952_v38
					_952_v38 = _dafny.SetOf(((_this).Fm0(_dafny.IntOfUint32((_951_v37).Cardinality()), _947_v33, globalState)).Minus((_950_v36).Cardinality()))
					var _953_v39 D10
					_ = _953_v39
					_953_v39 = Companion_D10_.Create_DC22_(_952_v38)
					var _954_v40 _dafny.Map
					_ = _954_v40
					_954_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_953_v39).Dtor_cf42())
					_952_v38 = (func() _dafny.Set {
						if (_954_v40).Contains((_this).F18()) {
							return (_954_v40).Get((_this).F18()).(_dafny.Set)
						}
						return _952_v38
					})()
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _955_v41 _dafny.Sequence
		_ = _955_v41
		_955_v41 = _dafny.UnicodeSeqOfUtf8Bytes("ky")
		var _956_v42 _dafny.CodePoint
		_ = _956_v42
		_956_v42 = _dafny.CodePoint('m')
		var _957_v43 D11
		_ = _957_v43
		_957_v43 = Companion_D11_.Create_DC25_(_this.F17())
		var _958_v44 *C4
		_ = _958_v44
		var _nw153 *C4 = New_C4_()
		_ = _nw153
		_nw153.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_955_v41).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_955_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-751), _dafny.IntOfUint32((_955_v41).Cardinality()))).Uint32(), _956_v42)).Cardinality())), p2, (_957_v43).Dtor_cf47(), ((_this).F28()).Cmp(_dafny.IntOfInt64(-569)) == 0)
		_958_v44 = _nw153
		if p1 {
			var _959_v45 _dafny.Array
			_ = _959_v45
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw154
			_959_v45 = _nw154
			var _960_v46 _dafny.Array
			_ = _960_v46
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
			_ = _nw155
			_960_v46 = _nw155
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_959_v45), 0))
			_ = _index147
			(_959_v45).ArraySet1(_960_v46, (_index147).Int())
			var _961_v47 _dafny.Map
			_ = _961_v47
			_961_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _960_v46)
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_959_v45), 0))
			_ = _index148
			(_959_v45).ArraySet1((func() _dafny.Array {
				if (_961_v47).Contains(p1) {
					return (_961_v47).Get(p1).(_dafny.Array)
				}
				return _960_v46
			})(), (_index148).Int())
			var _962_v48 _dafny.MultiSet
			_ = _962_v48
			_962_v48 = _dafny.MultiSetOf((_this).F20())
			(globalState).F15 = _962_v48
			(_958_v44).F26 = (_dafny.Zero).Minus(p2)
			_960_v46 = _960_v46
			var _963_v49 _dafny.Map
			_ = _963_v49
			_963_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_958_v44.F26, _958_v44.F26)
			var _964_v50 _dafny.Map
			_ = _964_v50
			_964_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_963_v49, (_this).F20())
			var _965_v51 _dafny.MultiSet
			_ = _965_v51
			_965_v51 = _dafny.MultiSetOf((_this).F18())
			var _966_v52 T2
			_ = _966_v52
			var _nw156 *C5 = New_C5_()
			_ = _nw156
			_nw156.Ctor__(_955_v41, (_this).Fm1((_964_v50).Update(_963_v49, (func() _dafny.Int {
				if (_965_v51).Contains((_this).F18()) {
					return (_965_v51).Multiplicity((_this).F18())
				}
				return _958_v44.F26
			})()), globalState), (_this).F28(), _this.F21(), (func() _dafny.Array {
				if p1 {
					return _this.F19()
				}
				return _this.F19()
			})(), (_this.F17()).Union(_this.F17()), (_this).F18())
			_966_v52 = _nw156
			var _967_v53 _dafny.Map
			_ = _967_v53
			_967_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_966_v52).F18(), p1)
			var _arr15 _dafny.Array = _966_v52.F19()
			_ = _arr15
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_966_v52.F19()), 0))
			_ = _index149
			(_arr15).ArraySet1((p0).Cmp((_967_v53).Cardinality()) < 0, (_index149).Int())
			var _968_v54 _dafny.Set
			_ = _968_v54
			_968_v54 = _dafny.SetOf((_this).F18())
			var _969_v55 D10
			_ = _969_v55
			_969_v55 = Companion_D10_.Create_DC23_(!((_966_v52).F18()), p1, _966_v52)
			var _arr16 _dafny.Array = _966_v52.F19()
			_ = _arr16
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_966_v52.F19()), 0))
			_ = _index150
			var _rhs130 _dafny.Int = _dafny.IntOfInt64(130)
			_ = _rhs130
			var _rhs131 _dafny.Int = _dafny.IntOfInt64(986)
			_ = _rhs131
			var _rhs132 T2 = (_969_v55).Dtor_cf45()
			_ = _rhs132
			var _rhs133 bool = !(((_966_v52).F18()) || (((_dafny.Zero).Minus(p0)).Cmp(p0) == 0))
			_ = _rhs133
			var _rhs134 _dafny.Set = _dafny.SetOf((_966_v52).F18())
			_ = _rhs134
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			var _lhs119 *GlobalState = globalState
			_ = _lhs119
			var _lhs120 _dafny.Array = _966_v52.F19()
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_966_v52.F19()), 0))
			_ = _lhs121
			_lhs118.F14 = _rhs130
			_lhs119.F14 = _rhs131
			_966_v52 = _rhs132
			(_lhs120).ArraySet1(_rhs133, (_lhs121).Int())
			_968_v54 = _rhs134
		} else {
			_955_v41 = _955_v41
			var _970_v56 _dafny.Array
			_ = _970_v56
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_32
			var _nw157 _dafny.Array
			_ = _nw157
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw157 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.CodePoint = func(_971_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw157 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw157).ArraySet1CodePoint(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw157).ArraySet1CodePoint(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_970_v56 = _nw157
			(globalState).F16 = _970_v56
			r2 = (_dafny.Zero).Minus(p0)
			(globalState).F9 = p1
			var _972_v57 _dafny.Sequence
			_ = _972_v57
			_972_v57 = _dafny.SeqOf((_this).F18(), p1)
			var _973_v58 D3
			_ = _973_v58
			_973_v58 = Companion_D3_.Create_DC9_()
			var _974_v59 _dafny.Map
			_ = _974_v59
			_974_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_972_v57, _972_v57), Companion_Default___.Fm34(_973_v58, p1, _956_v42, (_this).F18(), globalState))
			_955_v41 = (func() _dafny.Sequence {
				if (_974_v59).Contains(_dafny.Companion_Sequence_.Concatenate(_972_v57, _972_v57)) {
					return (_974_v59).Get(_dafny.Companion_Sequence_.Concatenate(_972_v57, _972_v57)).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(243))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_975_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_976_i3 _dafny.Int) _dafny.CodePoint {
						return _975_v42
					}
				})(_956_v42)))
			})()
		}
		var _977_v60 _dafny.Sequence
		_ = _977_v60
		_977_v60 = _dafny.SeqOf((_958_v44).F25())
		var _978_v62 _dafny.Sequence
		_ = _978_v62
		_978_v62 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_977_v60, (Companion_Default___.SafeIndex((func() _dafny.Map {
			var _coll45 = _dafny.NewMapBuilder()
			_ = _coll45
			for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-896), _dafny.IntOfInt64(31))); ; {
				_compr_45, _ok49 := _iter49()
				if !_ok49 {
					break
				}
				var _979_v61 _dafny.Int
				_979_v61 = interface{}(_compr_45).(_dafny.Int)
				if ((_dafny.IntOfInt64(-896)).Cmp(_979_v61) <= 0) && ((_979_v61).Cmp(_dafny.IntOfInt64(31)) < 0) {
					_coll45.Add(Companion_Default___.SafeDivisionInt(_979_v61, (_dafny.MultiSetOf(_dafny.IntOfInt64(810), _dafny.IntOfInt64(797))).Cardinality()), _977_v60)
				}
			}
			return _coll45.ToMap()
		}()).Cardinality(), _dafny.IntOfUint32((_977_v60).Cardinality()))).Uint32(), p2)).Cardinality()), (_this).F20(), p2)
		_978_v62 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-36))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg63 _dafny.Int) interface{} {
				return coer63(arg63)
			}
		}((func(_980_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_981_i4 _dafny.Int) _dafny.Int {
				return _980_p0
			}
		})(p0)))
		if (_this).F18() {
			r1 = (_this).F18()
			var _982_v63 _dafny.Sequence
			_ = _982_v63
			_982_v63 = _dafny.SeqOf((_this).F18(), p1, (_this).F18(), (_this).F18(), !((_this).F18()))
			var _rhs135 bool = false
			_ = _rhs135
			var _rhs136 bool = ((_this).F18()) || (!(_dafny.Companion_Sequence_.Equal(_982_v63, _982_v63)))
			_ = _rhs136
			var _rhs137 bool = !(true)
			_ = _rhs137
			var _lhs122 *GlobalState = globalState
			_ = _lhs122
			var _lhs123 *GlobalState = globalState
			_ = _lhs123
			r0 = _rhs135
			_lhs122.F9 = _rhs136
			_lhs123.F3 = _rhs137
			var _983_v64 _dafny.MultiSet
			_ = _983_v64
			_983_v64 = _dafny.MultiSetOf((_this).F20())
			var _984_v65 _dafny.Set
			_ = _984_v65
			_984_v65 = _dafny.SetOf(_958_v44.F26, (_958_v44).F25())
			var _nwElement0_30 bool = p1
			_ = _nwElement0_30
			var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(13))
			_ = _nw158
			(_nw158).ArraySet1(_nwElement0_30, 0)
			(_nw158).ArraySet1((_this).F18(), 1)
			(_nw158).ArraySet1((_983_v64).IsProperSubsetOf(_983_v64), 2)
			(_nw158).ArraySet1(p1, 3)
			(_nw158).ArraySet1(p1, 4)
			(_nw158).ArraySet1((_this).F18(), 5)
			(_nw158).ArraySet1((_this).F18(), 6)
			(_nw158).ArraySet1(p1, 7)
			(_nw158).ArraySet1(p1, 8)
			(_nw158).ArraySet1(!((_this).F18()), 9)
			(_nw158).ArraySet1((_this).F18(), 10)
			(_nw158).ArraySet1(!((_dafny.SetOf((_this).F20())).IsDisjointFrom(_984_v65)), 11)
			(_nw158).ArraySet1(!(!((_this).F18())) || ((_this).F18()), 12)
			(_this).F19_set_(_nw158)
			var _985_v66 bool
			_ = _985_v66
			var _986_v67 _dafny.Map
			_ = _986_v67
			var _987_v68 _dafny.Int
			_ = _987_v68
			var _out47 bool
			_ = _out47
			var _out48 _dafny.Map
			_ = _out48
			var _out49 _dafny.Int
			_ = _out49
			_out47, _out48, _out49 = (_958_v44).M4(globalState)
			_985_v66 = _out47
			_986_v67 = _out48
			_987_v68 = _out49
			var _988_v69 _dafny.Array
			_ = _988_v69
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_33
			var _nw159 _dafny.Array
			_ = _nw159
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw159 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Int = (func(_989_v44 *C4) func(_dafny.Int) _dafny.Int {
					return func(_990_i5 _dafny.Int) _dafny.Int {
						return (_990_i5).Times((_989_v44).F25())
					}
				})(_958_v44)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw159 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw159).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw159).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_988_v69 = _nw159
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_988_v69), 0))
			_ = _index151
			(_988_v69).ArraySet1(_958_v44.F26, (_index151).Int())
			var _991_v70 D0
			_ = _991_v70
			_991_v70 = Companion_D0_.Create_DC2_(_985_v66, (_958_v44).F25(), _988_v69)
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_988_v69), 0))
			_ = _index152
			var _rhs138 _dafny.Int = _958_v44.F26
			_ = _rhs138
			var _rhs139 _dafny.Int = _dafny.IntOfInt64(832)
			_ = _rhs139
			var _rhs140 _dafny.Int = (_977_v60).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F28(), _958_v44.F26), _dafny.IntOfUint32((_977_v60).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs140
			var _rhs141 _dafny.Array = (_991_v70).Dtor_cf7()
			_ = _rhs141
			var _lhs124 _dafny.Array = _988_v69
			_ = _lhs124
			var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_988_v69), 0))
			_ = _lhs125
			var _lhs126 *C4 = _958_v44
			_ = _lhs126
			(_lhs124).ArraySet1(_rhs138, (_lhs125).Int())
			_lhs126.F26 = _rhs139
			_987_v68 = _rhs140
			_988_v69 = _rhs141
		} else {
			var _992_v71 _dafny.Array
			_ = _992_v71
			var _nw160 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
			_ = _nw160
			_992_v71 = _nw160
			var _993_v72 _dafny.Map
			_ = _993_v72
			_993_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_992_v71, (_this).F18())
			_993_v72 = (_993_v72).Merge(_993_v72)
			var _994_v73 _dafny.Array
			_ = _994_v73
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw161
			_994_v73 = _nw161
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_994_v73), 0))
			_ = _index153
			(_994_v73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_955_v41, _955_v41), (_index153).Int())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_994_v73), 0))
			_ = _index154
			(_994_v73).ArraySet1(_955_v41, (_index154).Int())
			var _995_v74 _dafny.Array
			_ = _995_v74
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_34
			var _nw162 _dafny.Array
			_ = _nw162
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw162 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Int = (func(_996_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_997_i6 _dafny.Int) _dafny.Int {
						return (_997_i6).Minus(_996_p2)
					}
				})(p2)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw162 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw162).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw162).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_995_v74 = _nw162
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_995_v74), 0))
			_ = _index155
			(_995_v74).ArraySet1((_958_v44).F25(), (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_995_v74), 0))
			_ = _index156
			(_995_v74).ArraySet1((_this).F20(), (_index156).Int())
			(globalState).F9 = (_this).F18()
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_995_v74), 0))
			_ = _index157
			(_995_v74).ArraySet1((_958_v44.F26).Times(Companion_Default___.SafeDivisionInt((_this).F28(), _958_v44.F26)), (_index157).Int())
		}
		r0 = p1
		r1 = (_this).F18()
		var _998_v75 _dafny.Sequence
		_ = _998_v75
		_998_v75 = _dafny.SeqOf(true)
		r2 = Companion_Default___.SafeDivisionInt((p0).Plus((func() _dafny.Int {
			if ((_this).F27()).Contains((_this).F18()) {
				return ((_this).F27()).Get((_this).F18()).(_dafny.Int)
			}
			return (_this).Fm0(_dafny.IntOfUint32((_998_v75).Cardinality()), _956_v42, globalState)
		})()), _dafny.IntOfInt64(117))
		return r0, r1, r2
	}
}
func (_this *C6) M6(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _999_i0 _dafny.Int
		_ = _999_i0
		_999_i0 = _dafny.Zero
		{
			for (_dafny.IntOfInt64(757)).Cmp((_this).F20()) >= 0 {
				{
					if (_999_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_999_i0 = (_999_i0).Plus(_dafny.One)
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((p1), 0))
					_ = _index158
					(p1).ArraySet1((_this).F28(), (_index158).Int())
					var _1000_v0 _dafny.Array
					_ = _1000_v0
					var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(16))
					_ = _nw163
					_1000_v0 = _nw163
					var _1001_v1 _dafny.Set
					_ = _1001_v1
					_1001_v1 = _dafny.SetOf(_1000_v0, _1000_v0)
					var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((p1), 0))
					_ = _index159
					(p1).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(698), (_1001_v1).Cardinality()), (_index159).Int())
					var _1002_v2 _dafny.Array
					_ = _1002_v2
					var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
					_ = _nw164
					_1002_v2 = _nw164
					var _1003_v3 _dafny.Map
					_ = _1003_v3
					_1003_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1002_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)))
					var _1004_v4 _dafny.Map
					_ = _1004_v4
					_1004_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))
					_1003_v3 = (_1003_v3).Update(_1002_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F20())).Merge(_1004_v4))
					var _1005_v5 _dafny.Map
					_ = _1005_v5
					_1005_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F18())
					r2 = _dafny.SeqOf((p0).Cmp((_1005_v5).Cardinality()) <= 0, (_this).F18())
					var _1006_v6 _dafny.Sequence
					_ = _1006_v6
					_1006_v6 = _dafny.SeqOf(p0)
					var _1007_v7 *C4
					_ = _1007_v7
					var _nw165 *C4 = New_C4_()
					_ = _nw165
					_nw165.Ctor__(p3, Companion_Default___.SafeModuloInt((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (_this).F28()), _this.F17(), !((_this).F18()) || ((Companion_D9_.Create_DC20_(_1005_v5, _dafny.IntOfUint32((_1006_v6).Cardinality()), (_this).F18(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), _dafny.CodePoint('h'))).Dtor_cf38()))
					_1007_v7 = _nw165
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1008_i1 _dafny.Int
		_ = _1008_i1
		_1008_i1 = _dafny.Zero
		{
			for (p0).Cmp((_this).F28()) < 0 {
				{
					if (_1008_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1008_i1 = (_1008_i1).Plus(_dafny.One)
					var _1009_v8 _dafny.Sequence
					_ = _1009_v8
					_1009_v8 = _dafny.SeqOf((_this).F18(), (_this).F18())
					var _1010_v9 _dafny.Map
					_ = _1010_v9
					_1010_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1009_v8, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1009_v8).Cardinality()), (_this).F28()))
					var _1011_v10 _dafny.Map
					_ = _1011_v10
					_1011_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if ((_this).F27()).Contains(false) {
							return ((_this).F27()).Get(false).(_dafny.Int)
						}
						return p0
					})(), (_dafny.Zero).Minus((_this).Fm28(p3, _dafny.IntOfInt64(426), _1010_v9, globalState)))
					var _1012_v11 _dafny.Sequence
					_ = _1012_v11
					_1012_v11 = _dafny.SeqOf((_1011_v10).Cardinality(), (_this).F28(), p0)
					var _1013_v12 _dafny.Sequence
					_ = _1013_v12
					_1013_v12 = _dafny.UnicodeSeqOfUtf8Bytes("bimux")
					var _source10 _dafny.Sequence = Companion_Default___.Fm35(Companion_Default___.SafeModuloInt(p0, (_this).F20()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg64 _dafny.Int) interface{} {
							return coer64(arg64)
						}
					}((func(_1014_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1015_i2 _dafny.Int) _dafny.Int {
							return _1014_p0
						}
					})(p0))), _1012_v11), _dafny.IntOfUint32((_1013_v12).Cardinality()), globalState)
					_ = _source10
					var _1016___mcc_h0 _dafny.Sequence = _source10
					_ = _1016___mcc_h0
					var _1017_cf21 _dafny.Sequence = _1016___mcc_h0
					_ = _1017_cf21
					var _1018_v13 D2
					_ = _1018_v13
					_1018_v13 = Companion_D2_.Create_DC6_((_this).F28())
					(globalState).F14 = (p3).Plus((_1018_v13).Dtor_cf13())
					_1013_v12 = _1013_v12
					var _1019_v14 _dafny.Map
					_ = _1019_v14
					_1019_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F18()), (_this).F18())
					(globalState).F14 = Companion_Default___.SafeModuloInt(((_1019_v14).Cardinality()).Plus((_this).F20()), (((_this).F27()).Cardinality()).Minus((_this).F28()))
					var _arr17 _dafny.Array = _this.F19()
					_ = _arr17
					var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_this.F19()), 0))
					_ = _index160
					(_arr17).ArraySet1(false, (_index160).Int())
					var _1020_v15 _dafny.MultiSet
					_ = _1020_v15
					_1020_v15 = _dafny.MultiSetOf((_this).F18(), (_this).F18(), (_this).F18(), !((_this).F18()), (_this).F18())
					var _arr18 _dafny.Array = _this.F19()
					_ = _arr18
					var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_this.F19()), 0))
					_ = _index161
					(_arr18).ArraySet1((_1009_v8).Select((Companion_Default___.SafeIndex((_1020_v15).Cardinality(), _dafny.IntOfUint32((_1009_v8).Cardinality()))).Uint32()).(bool), (_index161).Int())
					var _1021_v16 T0
					_ = _1021_v16
					var _nw166 *C2 = New_C2_()
					_ = _nw166
					_nw166.Ctor__(_this.F17(), (_this).F18())
					_1021_v16 = _nw166
					var _1022_v17 _dafny.Map
					_ = _1022_v17
					_1022_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1013_v12, _1021_v16)
					(globalState).F14 = (func() _dafny.Int {
						if (_this).F18() {
							return p3
						}
						return Companion_Default___.SafeDivisionInt((_1022_v17).Cardinality(), (_this).F28())
					})()
					var _1023_v18 _dafny.CodePoint
					_ = _1023_v18
					_1023_v18 = _dafny.CodePoint('x')
					var _1024_v19 D11
					_ = _1024_v19
					_1024_v19 = Companion_D11_.Create_DC27_(_dafny.SeqOf(true, (_1021_v16).F18()), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).Fm0((_this).F20(), _1023_v18, globalState))), _dafny.IntOfInt64(334))
					var _pat_let_tv10 = _1009_v8
					_ = _pat_let_tv10
					var _pat_let_tv11 = _1009_v8
					_ = _pat_let_tv11
					var _1025_v20 _dafny.Map
					_ = _1025_v20
					_1025_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).Fm27(globalState)), func(_pat_let18_0 D11) D11 {
						return func(_1026_dt__update__tmp_h0 D11) D11 {
							return func(_pat_let19_0 _dafny.Int) D11 {
								return func(_1027_dt__update_hcf52_h0 _dafny.Int) D11 {
									return func(_pat_let20_0 _dafny.Sequence) D11 {
										return func(_1028_dt__update_hcf50_h0 _dafny.Sequence) D11 {
											return Companion_D11_.Create_DC27_(_1028_dt__update_hcf50_h0, (_1026_dt__update__tmp_h0).Dtor_cf51(), _1027_dt__update_hcf52_h0)
										}(_pat_let20_0)
									}(_pat_let_tv11)
								}(_pat_let19_0)
							}(_dafny.IntOfUint32((_pat_let_tv10).Cardinality()))
						}(_pat_let18_0)
					}(_1024_v19))
					_1025_v20 = (_1025_v20).Update(_1009_v8, _1024_v19)
					var _1029_v21 _dafny.MultiSet
					_ = _1029_v21
					_1029_v21 = _dafny.MultiSetOf((_this).F20())
					var _1030_v22 _dafny.Sequence
					_ = _1030_v22
					_1030_v22 = _dafny.SeqOf(_1029_v21)
					_1030_v22 = _1030_v22
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		if (_this).F18() {
			if (_this).F18() {
				var _1031_v23 _dafny.Sequence
				_ = _1031_v23
				_1031_v23 = _dafny.SeqOf(p1)
				(globalState).F7 = (_1031_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.IntOfUint32((_1031_v23).Cardinality()))).Uint32()).(_dafny.Array)
				var _1032_v24 _dafny.Map
				_ = _1032_v24
				_1032_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _this.F19())
				(globalState).F14 = ((_1032_v24).Merge((_1032_v24).Merge(_1032_v24))).Cardinality()
				var _1033_v25 _dafny.Sequence
				_ = _1033_v25
				_1033_v25 = _dafny.SeqOf(_dafny.IntOfInt64(106), p3, (_dafny.Zero).Minus(p0))
				r1 = _dafny.Companion_Sequence_.Concatenate(_1033_v25, _1033_v25)
				var _1034_v26 _dafny.Sequence
				_ = _1034_v26
				_1034_v26 = _dafny.UnicodeSeqOfUtf8Bytes("kfekrwy")
				_1034_v26 = _1034_v26
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((p1), 0))
				_ = _index162
				(p1).ArraySet1(p3, (_index162).Int())
				var _1035_v27 _dafny.CodePoint
				_ = _1035_v27
				_1035_v27 = _dafny.CodePoint('k')
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((p1), 0))
				_ = _index163
				(p1).ArraySet1((_this).Fm0((_this).F20(), _1035_v27, globalState), (_index163).Int())
			} else {
				var _1036_v28 _dafny.CodePoint
				_ = _1036_v28
				_1036_v28 = _dafny.CodePoint('d')
				var _1037_v29 _dafny.Sequence
				_ = _1037_v29
				_1037_v29 = _dafny.SeqOf(_1036_v28, _1036_v28, _1036_v28, _dafny.CodePoint('m'))
				var _1038_v30 _dafny.Array
				_ = _1038_v30
				var _nwElement0_31 _dafny.CodePoint = _1036_v28
				_ = _nwElement0_31
				var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(15))
				_ = _nw167
				(_nw167).ArraySet1CodePoint(_nwElement0_31, 0)
				(_nw167).ArraySet1CodePoint(_1036_v28, 1)
				(_nw167).ArraySet1CodePoint(_1036_v28, 2)
				(_nw167).ArraySet1CodePoint(_1036_v28, 3)
				(_nw167).ArraySet1CodePoint(Companion_Default___.Fm33(p3, (_this).F18(), globalState), 4)
				(_nw167).ArraySet1CodePoint(_1036_v28, 5)
				(_nw167).ArraySet1CodePoint(_1036_v28, 6)
				(_nw167).ArraySet1CodePoint((_1037_v29).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_1037_v29).Cardinality()))).Uint32()).(_dafny.CodePoint), 7)
				(_nw167).ArraySet1CodePoint(_1036_v28, 8)
				(_nw167).ArraySet1CodePoint(_1036_v28, 9)
				(_nw167).ArraySet1CodePoint(_1036_v28, 10)
				(_nw167).ArraySet1CodePoint(_1036_v28, 11)
				(_nw167).ArraySet1CodePoint(_1036_v28, 12)
				(_nw167).ArraySet1CodePoint(_1036_v28, 13)
				(_nw167).ArraySet1CodePoint(_dafny.CodePoint('l'), 14)
				_1038_v30 = _nw167
				(globalState).F16 = _1038_v30
				var _1039_v31 _dafny.Sequence
				_ = _1039_v31
				_1039_v31 = _dafny.SeqOf(p0)
				var _1040_v32 _dafny.Map
				_ = _1040_v32
				_1040_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfUint32((_1039_v31).Cardinality())), p3)
				_1036_v28 = Companion_Default___.Fm33((p3).Plus(p0), (_this).Fm1(_1040_v32, globalState), globalState)
				(globalState).F14 = (_dafny.Zero).Minus(p0)
				var _1041_v33 _dafny.MultiSet
				_ = _1041_v33
				_1041_v33 = _dafny.MultiSetOf(((_this).F27()).Cardinality())
				var _1042_v34 _dafny.Sequence
				_ = _1042_v34
				_1042_v34 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1039_v31, (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_1039_v31).Cardinality()))).Uint32(), (_this).F20())), _dafny.MultiSetOf(_dafny.IntOfUint32((_1037_v29).Cardinality())), _1041_v33, (_1041_v33).Difference(_1041_v33), _dafny.MultiSetOf(_dafny.IntOfInt64(500)))
				_1042_v34 = _1042_v34
				(globalState).F14 = (_dafny.Zero).Minus((_dafny.Zero).Minus((((_this).F20()).Minus(p3)).Minus((_this).F20())))
			}
			var _1043_v35 D3
			_ = _1043_v35
			_1043_v35 = Companion_D3_.Create_DC9_()
			var _1044_v36 _dafny.Sequence
			_ = _1044_v36
			_1044_v36 = _dafny.SeqOf(_1043_v35)
			var _1045_v37 D5
			_ = _1045_v37
			_1045_v37 = Companion_D5_.Create_DC13_((_this).F18(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhjby")).Cardinality()), _1044_v36)
			(globalState).F9 = (_1045_v37).Dtor_cf23()
			r0 = (_dafny.Zero).Minus((_this).F20())
			var _1046_v38 _dafny.Array
			_ = _1046_v38
			var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw168
			_1046_v38 = _nw168
			var _1047_v39 _dafny.Map
			_ = _1047_v39
			_1047_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ch"), p1)
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_1046_v38), 0))
			_ = _index164
			(_1046_v38).ArraySet1(_1047_v39, (_index164).Int())
			var _1048_v40 _dafny.Sequence
			_ = _1048_v40
			_1048_v40 = _dafny.UnicodeSeqOfUtf8Bytes("uqjs")
			var _1049_v41 D0
			_ = _1049_v41
			_1049_v41 = Companion_D0_.Create_DC2_((_this).F18(), _dafny.IntOfUint32((_1048_v40).Cardinality()), p1)
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_1046_v38), 0))
			_ = _index165
			(_1046_v38).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1048_v40, _dafny.UnicodeSeqOfUtf8Bytes("dsssjt")), (_1049_v41).Dtor_cf7()), (_index165).Int())
			if (_this).F18() {
				(globalState).F14 = p3
				var _rhs142 _dafny.Int = (_this).F28()
				_ = _rhs142
				var _rhs143 _dafny.Array = p1
				_ = _rhs143
				var _rhs144 bool = (_this).F18()
				_ = _rhs144
				var _rhs145 bool = (_this).F18()
				_ = _rhs145
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 *GlobalState = globalState
				_ = _lhs128
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				r0 = _rhs142
				_lhs127.F7 = _rhs143
				_lhs128.F9 = _rhs144
				_lhs129.F9 = _rhs145
				_1048_v40 = _dafny.UnicodeSeqOfUtf8Bytes("ydy")
				var _1050_v42 _dafny.Sequence
				_ = _1050_v42
				_1050_v42 = _dafny.SeqOf((_this).F18(), (_this).F18())
				var _1051_v43 _dafny.Map
				_ = _1051_v43
				_1051_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F19(), _1050_v42)
				var _rhs146 _dafny.Map = _1051_v43
				_ = _rhs146
				var _rhs147 _dafny.Int = (_this).F28()
				_ = _rhs147
				var _rhs148 bool = (_this).F18()
				_ = _rhs148
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				_1051_v43 = _rhs146
				_lhs130.F14 = _rhs147
				_lhs131.F3 = _rhs148
				var _1052_v44 _dafny.Array
				_ = _1052_v44
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw169
				_1052_v44 = _nw169
				_1052_v44 = _1052_v44
			} else {
				(globalState).F14 = p0
				(globalState).F9 = (func() bool {
					if (_this).F18() {
						return (_this).F18()
					}
					return true
				})()
				var _1053_v45 _dafny.CodePoint
				_ = _1053_v45
				_1053_v45 = _dafny.CodePoint('s')
				(_this).F21_set_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1053_v45, (_this).F18())).Merge(_this.F21()))
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((p1), 0))
				_ = _index166
				(p1).ArraySet1((_this).F20(), (_index166).Int())
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((p1), 0))
				_ = _index167
				(p1).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1048_v40, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1048_v40).Cardinality()))).Uint32(), _1053_v45)).Cardinality()), (_index167).Int())
				var _1054_v46 _dafny.Array
				_ = _1054_v46
				var _nwElement0_32 _dafny.CodePoint = _1053_v45
				_ = _nwElement0_32
				var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(21))
				_ = _nw170
				(_nw170).ArraySet1CodePoint(_nwElement0_32, 0)
				(_nw170).ArraySet1CodePoint(_1053_v45, 1)
				(_nw170).ArraySet1CodePoint(_1053_v45, 2)
				(_nw170).ArraySet1CodePoint(_1053_v45, 3)
				(_nw170).ArraySet1CodePoint(_1053_v45, 4)
				(_nw170).ArraySet1CodePoint(_1053_v45, 5)
				(_nw170).ArraySet1CodePoint(_1053_v45, 6)
				(_nw170).ArraySet1CodePoint(_1053_v45, 7)
				(_nw170).ArraySet1CodePoint(_1053_v45, 8)
				(_nw170).ArraySet1CodePoint(_dafny.CodePoint('a'), 9)
				(_nw170).ArraySet1CodePoint(_1053_v45, 10)
				(_nw170).ArraySet1CodePoint(_1053_v45, 11)
				(_nw170).ArraySet1CodePoint(_1053_v45, 12)
				(_nw170).ArraySet1CodePoint(_1053_v45, 13)
				(_nw170).ArraySet1CodePoint(_dafny.CodePoint('v'), 14)
				(_nw170).ArraySet1CodePoint((_1048_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1048_v40).Cardinality()), _dafny.IntOfUint32((_1048_v40).Cardinality()))).Uint32()).(_dafny.CodePoint), 15)
				(_nw170).ArraySet1CodePoint(_1053_v45, 16)
				(_nw170).ArraySet1CodePoint(_1053_v45, 17)
				(_nw170).ArraySet1CodePoint(_1053_v45, 18)
				(_nw170).ArraySet1CodePoint(_1053_v45, 19)
				(_nw170).ArraySet1CodePoint(_1053_v45, 20)
				_1054_v46 = _nw170
				var _1055_v47 _dafny.Map
				_ = _1055_v47
				_1055_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1054_v46)
				var _1056_v48 _dafny.Array
				_ = _1056_v48
				var _nwElement0_33 _dafny.Array = _1054_v46
				_ = _nwElement0_33
				var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(9))
				_ = _nw171
				(_nw171).ArraySet1(_nwElement0_33, 0)
				(_nw171).ArraySet1(_1054_v46, 1)
				(_nw171).ArraySet1(_1054_v46, 2)
				(_nw171).ArraySet1((func() _dafny.Array {
					if (_1055_v47).Contains((_this).F18()) {
						return (_1055_v47).Get((_this).F18()).(_dafny.Array)
					}
					return _1054_v46
				})(), 3)
				(_nw171).ArraySet1(_1054_v46, 4)
				(_nw171).ArraySet1(_1054_v46, 5)
				(_nw171).ArraySet1(_1054_v46, 6)
				(_nw171).ArraySet1(_1054_v46, 7)
				(_nw171).ArraySet1(_1054_v46, 8)
				_1056_v48 = _nw171
				var _1057_v49 _dafny.Map
				_ = _1057_v49
				_1057_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1056_v48, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1048_v40).Cardinality()), (_this).F20()))
				var _1058_v50 _dafny.Map
				_ = _1058_v50
				_1058_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
				_1057_v49 = (_1057_v49).Update(_1056_v48, _1058_v50)
			}
		} else {
			(globalState).F14 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}(func(_1059_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			}))).Cardinality()))
			var _1060_v51 _dafny.Sequence
			_ = _1060_v51
			_1060_v51 = _dafny.SeqOf(true, !((_this).F18()), !((_this).F18()), (_this).F18(), (_this).F18())
			var _1061_v52 _dafny.Map
			_ = _1061_v52
			_1061_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v51, false)
			var _1062_v53 _dafny.Set
			_ = _1062_v53
			_1062_v53 = _dafny.SetOf(_1060_v51)
			(globalState).F14 = (_dafny.Zero).Minus(((_1061_v52).Update(_1060_v51, !(!(_1062_v53).Equals(func() _dafny.Set {
				var _coll46 = _dafny.NewBuilder()
				_ = _coll46
				for _iter50 := _dafny.Iterate((_1062_v53).Elements()); ; {
					_compr_46, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _1063_v54 _dafny.Sequence
					_1063_v54 = interface{}(_compr_46).(_dafny.Sequence)
					if (_1062_v53).Contains(_1063_v54) {
						_coll46.Add(_1063_v54)
					}
				}
				return _coll46.ToSet()
			}())))).Cardinality())
			var _1064_v55 _dafny.Set
			_ = _1064_v55
			_1064_v55 = _dafny.SetOf((_this).F18(), (_this).F18(), (func() bool {
				if (_this).F18() {
					return (_this).F18()
				}
				return (_this).F18()
			})(), !(false) || ((_this).F18()))
			var _1065_v56 _dafny.MultiSet
			_ = _1065_v56
			_1065_v56 = _dafny.MultiSetOf((_this).F18(), (_this).F18(), (_this).F18())
			var _1066_v57 _dafny.Map
			_ = _1066_v57
			_1066_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1065_v56).Intersection(_1065_v56), (_this).F18())
			var _1067_v58 _dafny.Sequence
			_ = _1067_v58
			_1067_v58 = _dafny.UnicodeSeqOfUtf8Bytes("wupo")
			var _1068_v59 _dafny.Map
			_ = _1068_v59
			_1068_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1067_v58)
			var _rhs149 bool = !(!((_this).F18()))
			_ = _rhs149
			var _rhs150 _dafny.Set = p2
			_ = _rhs150
			var _rhs151 _dafny.Int = ((_1068_v59).Merge(Companion_Default___.Fm36(globalState))).Cardinality()
			_ = _rhs151
			var _rhs152 _dafny.Map = _1066_v57
			_ = _rhs152
			var _lhs132 *GlobalState = globalState
			_ = _lhs132
			_lhs132.F3 = _rhs149
			_1064_v55 = _rhs150
			r0 = _rhs151
			_1066_v57 = _rhs152
			var _1069_v60 T2
			_ = _1069_v60
			var _nw172 *C5 = New_C5_()
			_ = _nw172
			_nw172.Ctor__(_1067_v58, (_this).F18(), _dafny.IntOfUint32((_1060_v51).Cardinality()), _this.F21(), _this.F19(), _this.F17(), (_this).F18())
			_1069_v60 = _nw172
			var _1070_v61 D10
			_ = _1070_v61
			_1070_v61 = Companion_D10_.Create_DC23_((_this).F18(), (_this).F18(), _1069_v60)
			var _1071_v62 _dafny.Map
			_ = _1071_v62
			_1071_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1064_v55).Union(_1064_v55), _1070_v61)
			_1071_v62 = (_1071_v62).Update(Companion_Default___.Fm37((_1069_v60).F18(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yf")).Cardinality()), (_1069_v60).F18(), (_1069_v60).F18(), globalState), _1070_v61)
			(globalState).F14 = p0
		}
		if (_this).F18() {
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p1), 0))
			_ = _index168
			(p1).ArraySet1(p0, (_index168).Int())
			var _1072_v63 _dafny.Map
			_ = _1072_v63
			_1072_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-786)).Cmp((p2).Cardinality()) > 0, _dafny.IntOfInt64(202))
			var _1073_v64 _dafny.Sequence
			_ = _1073_v64
			_1073_v64 = _dafny.SeqOf((_this).F18())
			var _1074_v65 _dafny.Map
			_ = _1074_v65
			_1074_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
			var _1075_v66 _dafny.Map
			_ = _1075_v66
			_1075_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1074_v65)
			var _1076_v67 _dafny.Map
			_ = _1076_v67
			_1076_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1073_v64, (func() _dafny.Map {
				if (_1075_v66).Contains((_this).F18()) {
					return (_1075_v66).Get((_this).F18()).(_dafny.Map)
				}
				return _1074_v65
			})())
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p1), 0))
			_ = _index169
			var _rhs153 bool = (p3).Cmp((_this).Fm28(p3, (_this).F20(), _1076_v67, globalState)) > 0
			_ = _rhs153
			var _rhs154 bool = (_this).F18()
			_ = _rhs154
			var _rhs155 _dafny.Int = (_dafny.Zero).Minus(p3)
			_ = _rhs155
			var _rhs156 _dafny.Map = _1072_v63
			_ = _rhs156
			var _lhs133 *GlobalState = globalState
			_ = _lhs133
			var _lhs134 *GlobalState = globalState
			_ = _lhs134
			var _lhs135 _dafny.Array = p1
			_ = _lhs135
			var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p1), 0))
			_ = _lhs136
			_lhs133.F9 = _rhs153
			_lhs134.F9 = _rhs154
			(_lhs135).ArraySet1(_rhs155, (_lhs136).Int())
			_1072_v63 = _rhs156
			var _1077_v68 _dafny.Sequence
			_ = _1077_v68
			_1077_v68 = _dafny.UnicodeSeqOfUtf8Bytes("hifw")
			var _1078_v69 _dafny.MultiSet
			_ = _1078_v69
			_1078_v69 = _dafny.MultiSetOf(Companion_D11_.Create_DC27_(_1073_v64, (_dafny.Zero).Minus(p3), (_dafny.SetOf(p3)).Cardinality()))
			var _1079_v70 _dafny.Sequence
			_ = _1079_v70
			_1079_v70 = _dafny.SeqOf((_1078_v69).Intersection(_1078_v69), (_1078_v69).Union(_1078_v69))
			var _1080_v71 _dafny.Map
			_ = _1080_v71
			_1080_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1073_v64)
			var _rhs157 _dafny.Sequence = _1077_v68
			_ = _rhs157
			var _rhs158 _dafny.MultiSet = (_1079_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1080_v71).Contains(!(!((_this).F18()))) {
					return (_1080_v71).Get(!(!((_this).F18()))).(_dafny.Sequence)
				}
				return _1073_v64
			})()).Cardinality()), _dafny.IntOfUint32((_1079_v70).Cardinality()))).Uint32()).(_dafny.MultiSet)
			_ = _rhs158
			var _rhs159 bool = (_this).F18()
			_ = _rhs159
			var _lhs137 *GlobalState = globalState
			_ = _lhs137
			_1077_v68 = _rhs157
			_1078_v69 = _rhs158
			_lhs137.F3 = _rhs159
			var _arr19 _dafny.Array = _this.F19()
			_ = _arr19
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index170
			(_arr19).ArraySet1((_this).F18(), (_index170).Int())
			var _arr20 _dafny.Array = _this.F19()
			_ = _arr20
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index171
			(_arr20).ArraySet1(!((func() bool {
				if (func() bool {
					if !((_this).F18()) {
						return (_this).F18()
					}
					return (_this).F18()
				})() {
					return !((_this).F18())
				}
				return (_this).F18()
			})()), (_index171).Int())
			(globalState).F14 = (p0).Minus((_this).F28())
			var _1081_v72 _dafny.Array
			_ = _1081_v72
			var _nw173 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw173
			_1081_v72 = _nw173
			var _1082_v73 _dafny.Sequence
			_ = _1082_v73
			_1082_v73 = _dafny.SeqOf(_1081_v72, _1081_v72, _1081_v72)
			_1082_v73 = _dafny.Companion_Sequence_.Concatenate(_1082_v73, _1082_v73)
		} else {
			var _1083_v74 *C3
			_ = _1083_v74
			var _nw174 *C3 = New_C3_()
			_ = _nw174
			_nw174.Ctor__(_this.F19(), _this.F17(), (_this).F18())
			_1083_v74 = _nw174
			var _1084_v75 _dafny.Sequence
			_ = _1084_v75
			_1084_v75 = _dafny.UnicodeSeqOfUtf8Bytes("ikjvinf")
			var _1085_v76 D3
			_ = _1085_v76
			_1085_v76 = Companion_D3_.Create_DC9_()
			var _1086_v77 _dafny.CodePoint
			_ = _1086_v77
			_1086_v77 = _dafny.CodePoint('a')
			_1084_v75 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm34(_1085_v76, (_this).F18(), _1086_v77, (_this).F18(), globalState), _dafny.Companion_Sequence_.Concatenate(_1084_v75, _1084_v75))
			var _1087_v78 _dafny.Sequence
			_ = _1087_v78
			_1087_v78 = _dafny.SeqOf(true)
			var _1088_v79 _dafny.MultiSet
			_ = _1088_v79
			_1088_v79 = _dafny.MultiSetOf(_1087_v78, _1087_v78)
			var _1089_v80 _dafny.Map
			_ = _1089_v80
			_1089_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1085_v76, ((_1088_v79).Cardinality()).Cmp((_1083_v74).Fm0(p0, _1086_v77, globalState)) <= 0)
			_1089_v80 = (_1089_v80).Update(_1085_v76, (_this).F18())
			var _1090_v81 _dafny.Set
			_ = _1090_v81
			_1090_v81 = _dafny.SetOf(p3, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-98), (_this).F28())), _dafny.IntOfUint32((_dafny.SeqOf(((_this).F27()).Update((_this).F18(), (_this).F28()))).Cardinality()))
			_1090_v81 = (_1090_v81).Union(_1090_v81)
			if (_1087_v78).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1087_v78).Cardinality()))).Uint32()).(bool) {
				var _1091_v82 _dafny.MultiSet
				_ = _1091_v82
				_1091_v82 = _dafny.MultiSetOf(_dafny.MultiSetOf((_this).F18()))
				var _1092_v83 _dafny.Map
				_ = _1092_v83
				_1092_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F18())
				var _1093_v84 _dafny.Sequence
				_ = _1093_v84
				_1093_v84 = _dafny.SeqOf(Companion_Default___.Fm38((_1091_v82).Cardinality(), (_this).F18(), globalState), (_1092_v83).Update(_dafny.IntOfInt64(273), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F18()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).Cardinality(), true)).Update((_this).F20(), (_this).Fm27(globalState))).Update((_this).F28(), (_this).F18()), _1092_v83)
				var _1094_v86 _dafny.Map
				_ = _1094_v86
				_1094_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter51 := _dafny.Iterate((_dafny.SeqOf(p0, (_this).F28())).Elements()); ; {
						_compr_47, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1095_v85 _dafny.Int
						_1095_v85 = interface{}(_compr_47).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p0, (_this).F28()), _1095_v85) {
							_coll47.Add(Companion_Default___.SafeModuloInt(_1095_v85, (_this).F28()), _dafny.IntOfInt64(218))
						}
					}
					return _coll47.ToMap()
				}(), _dafny.IntOfUint32((_1084_v75).Cardinality()))
				_1093_v84 = (func() _dafny.Sequence {
					if (func() bool {
						if (_this).F18() {
							return (_this).F18()
						}
						return (_this).Fm1(_1094_v86, globalState)
					})() {
						return _dafny.Companion_Sequence_.Concatenate(_1093_v84, _1093_v84)
					}
					return (func() _dafny.Sequence {
						if (_this).F18() {
							return _1093_v84
						}
						return _1093_v84
					})()
				})()
				var _arr21 _dafny.Array = _this.F19()
				_ = _arr21
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index172
				(_arr21).ArraySet1((_this).F18(), (_index172).Int())
				var _arr22 _dafny.Array = _this.F19()
				_ = _arr22
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index173
				(_arr22).ArraySet1(!(Companion_Default___.Fm39(p3, globalState)).Contains((_this).Fm27(globalState)), (_index173).Int())
				var _1096_v87 _dafny.MultiSet
				_ = _1096_v87
				_1096_v87 = _dafny.MultiSetOf((_this).F18(), (_this).F18(), (_this).F18())
				var _arr23 _dafny.Array = _this.F19()
				_ = _arr23
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index174
				var _rhs160 bool = !(_1096_v87).Equals(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _1087_v78)))
				_ = _rhs160
				var _rhs161 bool = (_this).Fm27(globalState)
				_ = _rhs161
				var _rhs162 _dafny.Array = p1
				_ = _rhs162
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 _dafny.Array = _this.F19()
				_ = _lhs139
				var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))
				_ = _lhs140
				var _lhs141 *GlobalState = globalState
				_ = _lhs141
				_lhs138.F9 = _rhs160
				(_lhs139).ArraySet1(_rhs161, (_lhs140).Int())
				_lhs141.F7 = _rhs162
				var _1097_v88 _dafny.Map
				_ = _1097_v88
				_1097_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), true)
				var _rhs163 _dafny.Array = p1
				_ = _rhs163
				var _rhs164 bool = !((func() bool {
					if (_1097_v88).Contains((!((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))) || ((_this).F18())) {
						return (_1097_v88).Get((!((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))) || ((_this).F18())).(bool)
					}
					return (_this).F18()
				})())
				_ = _rhs164
				var _rhs165 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("mfq")
				_ = _rhs165
				var _rhs166 bool = _dafny.Companion_Sequence_.Contains(_1084_v75, _1086_v77)
				_ = _rhs166
				var _rhs167 bool = (_this).F18()
				_ = _rhs167
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				_lhs142.F7 = _rhs163
				_lhs143.F3 = _rhs164
				_1084_v75 = _rhs165
				_lhs144.F3 = _rhs166
				_lhs145.F9 = _rhs167
				var _1098_v90 _dafny.Map
				_ = _1098_v90
				_1098_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1087_v78, func() _dafny.Map {
					var _coll48 = _dafny.NewMapBuilder()
					_ = _coll48
					for _iter52 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))).Keys().Elements()); ; {
						_compr_48, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1099_v89 _dafny.Int
						_1099_v89 = interface{}(_compr_48).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))).Contains(_1099_v89) {
							_coll48.Add((_1099_v89).Minus(p0), p3)
						}
					}
					return _coll48.ToMap()
				}())
				var _rhs168 _dafny.Int = (_this).F28()
				_ = _rhs168
				var _rhs169 bool = (_this).F18()
				_ = _rhs169
				var _rhs170 _dafny.Int = (_this).Fm28((_this).F20(), (p0).Minus(p0), _1098_v90, globalState)
				_ = _rhs170
				var _rhs171 bool = (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
				_ = _rhs171
				var _rhs172 _dafny.Int = p3
				_ = _rhs172
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				var _lhs148 *GlobalState = globalState
				_ = _lhs148
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				var _lhs150 *GlobalState = globalState
				_ = _lhs150
				_lhs146.F14 = _rhs168
				_lhs147.F3 = _rhs169
				_lhs148.F14 = _rhs170
				_lhs149.F3 = _rhs171
				_lhs150.F14 = _rhs172
			} else {
				var _1100_v91 _dafny.Sequence
				_ = _1100_v91
				_1100_v91 = _dafny.SeqOf(_1084_v75)
				var _1101_v92 _dafny.Map
				_ = _1101_v92
				_1101_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.Companion_Sequence_.IsProperPrefixOf(_1100_v91, _1100_v91))
				_1101_v92 = (_1101_v92).Update((_this).F18(), (func() bool {
					if (_this).F18() {
						return (_this).F18()
					}
					return (_1087_v78).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1087_v78).Cardinality()))).Uint32()).(bool)
				})())
				_1090_v81 = _1090_v81
				var _1102_v93 _dafny.Sequence
				_ = _1102_v93
				_1102_v93 = _dafny.SeqOf(((_dafny.SetOf(false, (_this).F18(), (_this).F18())).Union(p2)).Cardinality(), (_this).F20())
				r0 = (_1102_v93).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1102_v93).Cardinality()))).Uint32()).(_dafny.Int)
				(_this).F19_set_(_this.F19())
				(globalState).F9 = ((_dafny.Zero).Minus((p3).Times((_this).F28()))).Cmp(p3) == 0
			}
		}
		var _1103_v94 *C2
		_ = _1103_v94
		var _nw175 *C2 = New_C2_()
		_ = _nw175
		_nw175.Ctor__((_this.F17()).Difference(_this.F17()), (_this).F18())
		_1103_v94 = _nw175
		var _1104_v95 _dafny.Sequence
		_ = _1104_v95
		_1104_v95 = _dafny.SeqOf((_this).F18())
		var _1105_v96 _dafny.MultiSet
		_ = _1105_v96
		_1105_v96 = _dafny.MultiSetOf((_this).F18())
		var _arr24 _dafny.Array = _this.F19()
		_ = _arr24
		var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_this.F19()), 0))
		_ = _index175
		(_arr24).ArraySet1((_1104_v95).Select((Companion_Default___.SafeIndex((_1105_v96).Cardinality(), _dafny.IntOfUint32((_1104_v95).Cardinality()))).Uint32()).(bool), (_index175).Int())
		var _arr25 _dafny.Array = _this.F19()
		_ = _arr25
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_this.F19()), 0))
		_ = _index176
		(_arr25).ArraySet1((_this).F18(), (_index176).Int())
		r0 = (_1105_v96).Cardinality()
		var _1106_v97 _dafny.Sequence
		_ = _1106_v97
		_1106_v97 = _dafny.SeqOf(p0)
		var _1107_v98 _dafny.Set
		_ = _1107_v98
		_1107_v98 = _dafny.SetOf((_this).F20())
		var _1108_v99 _dafny.Sequence
		_ = _1108_v99
		_1108_v99 = _dafny.UnicodeSeqOfUtf8Bytes("dsaa")
		var _1109_v100 _dafny.Map
		_ = _1109_v100
		_1109_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1104_v95).Cardinality()), p3)
		var _1110_v101 _dafny.Map
		_ = _1110_v101
		_1110_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1104_v95, _1109_v100)
		var _1111_v102 _dafny.Sequence
		_ = _1111_v102
		_1111_v102 = _dafny.Companion_Sequence_.Update(_1106_v97, (Companion_Default___.SafeIndex((_1107_v98).Cardinality(), _dafny.IntOfUint32((_1106_v97).Cardinality()))).Uint32(), (_this).Fm28(_dafny.IntOfUint32((_1108_v99).Cardinality()), p3, _1110_v101, globalState))
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1106_v97, _dafny.Companion_Sequence_.Concatenate(_1106_v97, (_1111_v102))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1108_v99, _dafny.UnicodeSeqOfUtf8Bytes("o"))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1106_v97, _dafny.Companion_Sequence_.Concatenate(_1106_v97, (_1111_v102)))).Cardinality()))).Uint32(), (_this).F20())
		r2 = _1104_v95
		return r0, r1, r2
	}
}
func (_this *C6) M7(p0 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _1112_v0 _dafny.Sequence
		_ = _1112_v0
		_1112_v0 = _dafny.UnicodeSeqOfUtf8Bytes("beccgqvi")
		(globalState).F14 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(519))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}(func(_1113_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), _1112_v0)).Cardinality()))
		var _1114_v1 _dafny.Sequence
		_ = _1114_v1
		_1114_v1 = _dafny.SeqOf((_this).F18(), (_this).F18(), (_this).F18(), (_this).F18(), (_this).F18())
		var _1115_v2 _dafny.MultiSet
		_ = _1115_v2
		_1115_v2 = _dafny.MultiSetOf(_1114_v1)
		(globalState).F3 = ((_1115_v2).Cardinality()).Cmp(p0) >= 0
		var _1116_v3 _dafny.CodePoint
		_ = _1116_v3
		_1116_v3 = _dafny.CodePoint('j')
		_1116_v3 = _1116_v3
		var _1117_v4 _dafny.Sequence
		_ = _1117_v4
		_1117_v4 = _dafny.SeqOf(_dafny.IntOfUint32((_1112_v0).Cardinality()), p0, p0)
		var _1118_v5 _dafny.Set
		_ = _1118_v5
		_1118_v5 = _dafny.SetOf((_this).F18())
		(globalState).F14 = (_dafny.IntOfUint32((_1117_v4).Cardinality())).Times(((_dafny.SetOf((_this).F18())).Union(_1118_v5)).Cardinality())
		(globalState).F14 = ((_this).F28()).Minus(_dafny.IntOfUint32((_1112_v0).Cardinality()))
		var _1119_i1 _dafny.Int
		_ = _1119_i1
		_1119_i1 = _dafny.Zero
		{
			for (_this).F18() {
				{
					if (_1119_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1119_i1 = (_1119_i1).Plus(_dafny.One)
					(globalState).F9 = (_this).F18()
					var _1120_v6 _dafny.Map
					_ = _1120_v6
					_1120_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_1112_v0, _1112_v0), (_this).F18())
					_1120_v6 = (_1120_v6).Update((_this).F18(), !((_this).F18()))
					_1112_v0 = (func() _dafny.Sequence {
						if (_this).F18() {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(276))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg67 _dafny.Int) interface{} {
									return coer67(arg67)
								}
							}((func(_1121_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1122_i2 _dafny.Int) _dafny.CodePoint {
									return _1121_v3
								}
							})(_1116_v3)))
						}
						return _1112_v0
					})()
					var _1123_v7 *C1
					_ = _1123_v7
					var _nw176 *C1 = New_C1_()
					_ = _nw176
					_nw176.Ctor__()
					_1123_v7 = _nw176
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1124_v8 T0
		_ = _1124_v8
		var _nw177 *C4 = New_C4_()
		_ = _nw177
		_nw177.Ctor__((_this).F28(), (_1118_v5).Cardinality(), _this.F17(), (_this).F18())
		_1124_v8 = _nw177
		var _1125_v9 _dafny.Sequence
		_ = _1125_v9
		_1125_v9 = _dafny.SeqOf(_1124_v8)
		var _1126_v10 D12
		_ = _1126_v10
		_1126_v10 = Companion_D12_.Create_DC29_(_1125_v9)
		var _1127_v11 _dafny.Sequence
		_ = _1127_v11
		_1127_v11 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1125_v9, (Companion_Default___.SafeIndex((func() _dafny.Int {
			if ((_this).F27()).Contains((_1124_v8).F18()) {
				return ((_this).F27()).Get((_1124_v8).F18()).(_dafny.Int)
			}
			return (_this).F28()
		})(), _dafny.IntOfUint32((_1125_v9).Cardinality()))).Uint32(), _1124_v8), _1125_v9), _1125_v9, (_1126_v10).Dtor_cf54())
		r0 = (_1127_v11).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_1127_v11).Cardinality()))).Uint32()).(_dafny.Sequence)
		r1 = (func() _dafny.Int {
			if (_this).F18() {
				return (func() _dafny.Int {
					if (_this).F18() {
						return (_this).F28()
					}
					return p0
				})()
			}
			return (_this).F20()
		})()
		var _1128_v12 D3
		_ = _1128_v12
		_1128_v12 = Companion_D3_.Create_DC10_(_1112_v0, (_1124_v8).F18(), _dafny.UnicodeSeqOfUtf8Bytes("hym"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(34))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}((func(_1129_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1130_i3 _dafny.Int) _dafny.CodePoint {
				return _1129_v3
			}
		})(_1116_v3))), (_dafny.Zero).Minus(p0))
		var _1131_v13 _dafny.MultiSet
		_ = _1131_v13
		_1131_v13 = _dafny.MultiSetOf((_1124_v8).F18(), (_this).F18(), (_this).F18(), (_1128_v12).Dtor_cf17())
		r2 = (_1131_v13).IsSubsetOf(_1131_v13)
		r3 = _1117_v4
		return r0, r1, r2, r3
	}
}
func (_this *C6) F28() _dafny.Int {
	{
		return _this._f28
	}
}
func (_this *C6) F27() _dafny.Map {
	{
		return _this._f27
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f17 _dafny.Set
	_f19 _dafny.Array
	_f21 _dafny.Map
	_f20 _dafny.Int
	_f18 bool
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f21 = _dafny.EmptyMap
	_this._f20 = _dafny.Zero
	_this._f18 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C7{}
var _ T1 = &C7{}
var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F17() _dafny.Set {
	return _this._f17
}
func (_this *C7) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C7) F19() _dafny.Array {
	return _this._f19
}
func (_this *C7) F19_set_(value _dafny.Array) {
	_this._f19 = value
}
func (_this *C7) F21() _dafny.Map {
	return _this._f21
}
func (_this *C7) F21_set_(value _dafny.Map) {
	_this._f21 = value
}
func (_this *C7) F20() _dafny.Int {
	return _this._f20
}
func (_this *C7) F18() bool {
	return _this._f18
}
func (_this *C7) Ctor__(f20 _dafny.Int, f21 _dafny.Map, f19 _dafny.Array, f17 _dafny.Set, f18 bool) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C7) Fm1(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		return (_this).F18()
	}
}
func (_this *C7) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F20()
	}
}
func (_this *C7) Fm25(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cd")).Cardinality())).Cmp((_dafny.Zero).Minus((_this).F20())) < 0, (_this).F18(), (_this).F18(), (_this).F18())
	}
}
func (_this *C7) Fm26(globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetOf((_this).F18(), (_this).F18()))
	}
}
func (_this *C7) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1132_v0 _dafny.Map
		_ = _1132_v0
		_1132_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), p1)
		var _1133_v1 _dafny.Set
		_ = _1133_v1
		_1133_v1 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(797), p1))
		var _1134_v2 D1
		_ = _1134_v2
		_1134_v2 = Companion_D1_.Create_DC4_(((func() bool {
			if (_1132_v0).Contains(p0) {
				return (_1132_v0).Get(p0).(bool)
			}
			return true
		})()) && ((_this).F18()), (_dafny.SetOf(_1132_v0)).Intersection(_1133_v1), p0)
		var _source11 D1 = _1134_v2
		_ = _source11
		if _source11.Is_DC4() {
			var _1135___mcc_h0 bool = _source11.Get_().(D1_DC4).Cf9
			_ = _1135___mcc_h0
			var _1136___mcc_h1 _dafny.Set = _source11.Get_().(D1_DC4).Cf10
			_ = _1136___mcc_h1
			var _1137___mcc_h2 _dafny.Int = _source11.Get_().(D1_DC4).Cf11
			_ = _1137___mcc_h2
			var _1138_cf11 _dafny.Int = _1137___mcc_h2
			_ = _1138_cf11
			var _1139_cf10 _dafny.Set = _1136___mcc_h1
			_ = _1139_cf10
			var _1140_cf9 bool = _1135___mcc_h0
			_ = _1140_cf9
			var _1141_v3 _dafny.Sequence
			_ = _1141_v3
			_1141_v3 = _dafny.UnicodeSeqOfUtf8Bytes("edctjguo")
			var _1142_v4 T2
			_ = _1142_v4
			var _nw178 *C5 = New_C5_()
			_ = _nw178
			_nw178.Ctor__(_1141_v3, p1, (_this).F20(), _this.F21(), _this.F19(), _this.F17(), _1140_cf9)
			_1142_v4 = _nw178
			_1142_v4 = _1142_v4
			_1138_cf11 = (p0).Minus(_1138_cf11)
			var _1143_v5 _dafny.MultiSet
			_ = _1143_v5
			_1143_v5 = _dafny.MultiSetOf(_1134_v2)
			(globalState).F14 = (func() _dafny.Int {
				if (_1143_v5).Contains(Companion_D1_.Create_DC4_(_1140_cf9, _1133_v1, _dafny.IntOfInt64(885))) {
					return (_1143_v5).Multiplicity(Companion_D1_.Create_DC4_(_1140_cf9, _1133_v1, _dafny.IntOfInt64(885)))
				}
				return (_this).F20()
			})()
			var _1144_v6 _dafny.Map
			_ = _1144_v6
			_1144_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1138_cf11).Cmp(p2) > 0, _this.F19())
			var _1145_v7 _dafny.Sequence
			_ = _1145_v7
			_1145_v7 = _dafny.SeqOf((func() bool {
				if (_1132_v0).Contains((_1142_v4).F20()) {
					return (_1132_v0).Get((_1142_v4).F20()).(bool)
				}
				return false
			})(), (_1142_v4).F18(), _1140_cf9, !(!(!((_1142_v4).F18()))), (_1142_v4).F18())
			_1144_v6 = (_1144_v6).Update((_1145_v7).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1145_v7).Cardinality()))).Uint32()).(bool), _this.F19())
		} else {
			var _1146___mcc_h3 _dafny.MultiSet = _source11.Get_().(D1_DC3).Cf8
			_ = _1146___mcc_h3
			var _1147_cf8 _dafny.MultiSet = _1146___mcc_h3
			_ = _1147_cf8
			var _1148_v8 _dafny.Sequence
			_ = _1148_v8
			_1148_v8 = _dafny.UnicodeSeqOfUtf8Bytes("k")
			var _1149_v9 _dafny.Map
			_ = _1149_v9
			_1149_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Concatenate(_1148_v8, _1148_v8))
			var _1150_v10 _dafny.Sequence
			_ = _1150_v10
			_1150_v10 = _dafny.SeqOf(_1149_v9)
			var _rhs173 bool = !(!(p1))
			_ = _rhs173
			var _rhs174 _dafny.Map = (_1150_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1150_v10).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs174
			var _rhs175 _dafny.Int = (Companion_D2_.Create_DC6_(_dafny.IntOfInt64(-752))).Dtor_cf13()
			_ = _rhs175
			var _rhs176 _dafny.Array = _this.F19()
			_ = _rhs176
			var _rhs177 _dafny.MultiSet = _1147_cf8
			_ = _rhs177
			var _lhs151 *GlobalState = globalState
			_ = _lhs151
			var _lhs152 *C7 = _this
			_ = _lhs152
			_lhs151.F3 = _rhs173
			_1149_v9 = _rhs174
			r2 = _rhs175
			_lhs152.F19_set_(_rhs176)
			_1147_cf8 = _rhs177
			var _1151_v11 _dafny.Array
			_ = _1151_v11
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_35
			var _nw179 _dafny.Array
			_ = _nw179
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw179 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = (func(_1152_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1153_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1153_i0, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_1152_p2, (_dafny.SetOf(false)).Cardinality(), (_this).F20())).Cardinality(), _1152_p2)).Cardinality()))
					}
				})(p2)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw179 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw179).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw179).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1151_v11 = _nw179
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_1151_v11), 0))
			_ = _index177
			(_1151_v11).ArraySet1(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(336)), (_index177).Int())
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_1151_v11), 0))
			_ = _index178
			(_1151_v11).ArraySet1((_this).F20(), (_index178).Int())
			var _1154_v12 _dafny.CodePoint
			_ = _1154_v12
			_1154_v12 = _dafny.CodePoint('v')
			var _arr26 _dafny.Array = _this.F19()
			_ = _arr26
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index179
			(_arr26).ArraySet1(p1, (_index179).Int())
			var _1155_v13 _dafny.Map
			_ = _1155_v13
			_1155_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), p2)
			var _1156_v14 _dafny.Sequence
			_ = _1156_v14
			_1156_v14 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_1148_v8).Cardinality()), p0))
			var _arr27 _dafny.Array = _this.F19()
			_ = _arr27
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_this.F19()), 0))
			_ = _index180
			var _rhs178 _dafny.CodePoint = (func() _dafny.CodePoint {
				if p1 {
					return _1154_v12
				}
				return _1154_v12
			})()
			_ = _rhs178
			var _rhs179 bool = ((func() _dafny.Int {
				if (_1155_v13).Contains(_dafny.IntOfUint32((_1156_v14).Cardinality())) {
					return (_1155_v13).Get(_dafny.IntOfUint32((_1156_v14).Cardinality())).(_dafny.Int)
				}
				return _dafny.IntOfInt64(249)
			})()).Cmp(Companion_Default___.SafeModuloInt((_1151_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_1151_v11), 0))).Int()).(_dafny.Int), (_1151_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_1151_v11), 0))).Int()).(_dafny.Int))) < 0
			_ = _rhs179
			var _lhs153 _dafny.Array = _this.F19()
			_ = _lhs153
			var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_this.F19()), 0))
			_ = _lhs154
			_1154_v12 = _rhs178
			(_lhs153).ArraySet1(_rhs179, (_lhs154).Int())
			var _1157_v15 D0
			_ = _1157_v15
			_1157_v15 = Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("aw"), p0, _dafny.IntOfInt64(155), _1154_v12)
			var _1158_v16 _dafny.Sequence
			_ = _1158_v16
			_1158_v16 = _dafny.SeqOf(_1148_v8, _dafny.Companion_Sequence_.Concatenate((_1157_v15).Dtor_cf1(), _1148_v8), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(557))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_1159_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _1148_v8)
			_1158_v16 = _dafny.Companion_Sequence_.Concatenate(_1158_v16, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(386))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}((func(_1160_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1161_i2 _dafny.Int) _dafny.Sequence {
					return _1160_v8
				}
			})(_1148_v8))))
		}
		var _1162_v17 _dafny.Map
		_ = _1162_v17
		_1162_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F19())
		var _1163_v18 _dafny.Map
		_ = _1163_v18
		_1163_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1162_v17).Cardinality())
		if (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v18, p0), globalState) {
			var _1164_v19 _dafny.Array
			_ = _1164_v19
			var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw180
			_1164_v19 = _nw180
			var _1165_v20 _dafny.Sequence
			_ = _1165_v20
			_1165_v20 = _dafny.SeqOf(p0, p2, p2)
			var _1166_v21 _dafny.CodePoint
			_ = _1166_v21
			_1166_v21 = _dafny.CodePoint('s')
			var _rhs180 _dafny.Array = _1164_v19
			_ = _rhs180
			var _rhs181 _dafny.Int = (_this).Fm0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1165_v20, _1165_v20)).Cardinality()), _1166_v21, globalState)
			_ = _rhs181
			var _rhs182 bool = false
			_ = _rhs182
			var _rhs183 bool = !(true)
			_ = _rhs183
			var _rhs184 bool = !(!((_this).F18()))
			_ = _rhs184
			var _lhs155 *GlobalState = globalState
			_ = _lhs155
			var _lhs156 *GlobalState = globalState
			_ = _lhs156
			var _lhs157 *GlobalState = globalState
			_ = _lhs157
			var _lhs158 *GlobalState = globalState
			_ = _lhs158
			var _lhs159 *GlobalState = globalState
			_ = _lhs159
			_lhs155.F7 = _rhs180
			_lhs156.F14 = _rhs181
			_lhs157.F3 = _rhs182
			_lhs158.F9 = _rhs183
			_lhs159.F9 = _rhs184
			var _1167_v22 _dafny.Sequence
			_ = _1167_v22
			_1167_v22 = _dafny.UnicodeSeqOfUtf8Bytes("fm")
			var _1168_v23 *C5
			_ = _1168_v23
			var _nw181 *C5 = New_C5_()
			_ = _nw181
			_nw181.Ctor__(_1167_v22, !(p1), _dafny.IntOfUint32((_1167_v22).Cardinality()), _this.F21(), _this.F19(), (_this.F17()).Difference(_this.F17()), (_this).F18())
			_1168_v23 = _nw181
			var _1169_v24 _dafny.Map
			_ = _1169_v24
			_1169_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p2)
			_1169_v24 = _1169_v24
			if false {
				_1164_v19 = _1164_v19
				_1166_v21 = _1166_v21
				r1 = (p2).Cmp((_this).F20()) >= 0
				var _1170_v25 *C2
				_ = _1170_v25
				var _nw182 *C2 = New_C2_()
				_ = _nw182
				_nw182.Ctor__(_this.F17(), (_1168_v23).F30())
				_1170_v25 = _nw182
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1164_v19), 0))
				_ = _index181
				(_1164_v19).ArraySet1((_this).F20(), (_index181).Int())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1164_v19), 0))
				_ = _index182
				(_1164_v19).ArraySet1((_this).F20(), (_index182).Int())
			} else {
				var _arr28 _dafny.Array = _this.F19()
				_ = _arr28
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index183
				(_arr28).ArraySet1(!(false) || ((_this).F18()), (_index183).Int())
				var _arr29 _dafny.Array = _this.F19()
				_ = _arr29
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index184
				(_arr29).ArraySet1((_1168_v23).F30(), (_index184).Int())
				var _1171_v26 _dafny.Array
				_ = _1171_v26
				var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw183
				_1171_v26 = _nw183
				var _1172_v27 _dafny.Sequence
				_ = _1172_v27
				_1172_v27 = _dafny.SeqOf(_1171_v26)
				var _1173_v28 _dafny.Map
				_ = _1173_v28
				_1173_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1163_v18).Contains(p0) {
						return (_1163_v18).Get(p0).(_dafny.Int)
					}
					return (_this).F20()
				})(), _1171_v26)
				var _1174_v29 _dafny.Map
				_ = _1174_v29
				_1174_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (func() _dafny.Array {
					if (_1173_v28).Contains(p2) {
						return (_1173_v28).Get(p2).(_dafny.Array)
					}
					return _1171_v26
				})())
				var _1175_v30 _dafny.Array
				_ = _1175_v30
				var _nwElement0_34 _dafny.Array = _1171_v26
				_ = _nwElement0_34
				var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(18))
				_ = _nw184
				(_nw184).ArraySet1(_nwElement0_34, 0)
				(_nw184).ArraySet1(_1171_v26, 1)
				(_nw184).ArraySet1(_1171_v26, 2)
				(_nw184).ArraySet1(_1171_v26, 3)
				(_nw184).ArraySet1(_1171_v26, 4)
				(_nw184).ArraySet1(_1171_v26, 5)
				(_nw184).ArraySet1(_1171_v26, 6)
				(_nw184).ArraySet1(_1171_v26, 7)
				(_nw184).ArraySet1((_1172_v27).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1172_v27).Cardinality()))).Uint32()).(_dafny.Array), 8)
				(_nw184).ArraySet1((func() _dafny.Array {
					if (_1173_v28).Contains(p0) {
						return (_1173_v28).Get(p0).(_dafny.Array)
					}
					return _1171_v26
				})(), 9)
				(_nw184).ArraySet1(_1171_v26, 10)
				(_nw184).ArraySet1(_1171_v26, 11)
				(_nw184).ArraySet1((func() _dafny.Array {
					if (_1174_v29).Contains(p1) {
						return (_1174_v29).Get(p1).(_dafny.Array)
					}
					return _1171_v26
				})(), 12)
				(_nw184).ArraySet1(_1171_v26, 13)
				(_nw184).ArraySet1(_1171_v26, 14)
				(_nw184).ArraySet1(_1171_v26, 15)
				(_nw184).ArraySet1(_1171_v26, 16)
				(_nw184).ArraySet1(_1171_v26, 17)
				_1175_v30 = _nw184
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1175_v30), 0))
				_ = _index185
				(_1175_v30).ArraySet1(_1171_v26, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1175_v30), 0))
				_ = _index186
				(_1175_v30).ArraySet1(_1171_v26, (_index186).Int())
				_1164_v19 = _1164_v19
				var _1176_v31 _dafny.Map
				_ = _1176_v31
				_1176_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v18, p0)
				var _1177_v35 _dafny.Array
				_ = _1177_v35
				var _nwElement0_35 _dafny.Map = (_1163_v18).Merge(_1163_v18)
				_ = _nwElement0_35
				var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(14))
				_ = _nw185
				(_nw185).ArraySet1(_nwElement0_35, 0)
				(_nw185).ArraySet1((_1163_v18).Merge(_1163_v18), 1)
				(_nw185).ArraySet1(_1163_v18, 2)
				(_nw185).ArraySet1((func() _dafny.Map {
					if (_1168_v23).Fm1(_1176_v31, globalState) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
					}
					return func() _dafny.Map {
						var _coll49 = _dafny.NewMapBuilder()
						_ = _coll49
						for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(982), _dafny.IntOfInt64(808))); ; {
							_compr_49, _ok53 := _iter53()
							if !_ok53 {
								break
							}
							var _1178_v32 _dafny.Int
							_1178_v32 = interface{}(_compr_49).(_dafny.Int)
							if ((_dafny.IntOfInt64(982)).Cmp(_1178_v32) <= 0) && ((_1178_v32).Cmp(_dafny.IntOfInt64(808)) < 0) {
								_coll49.Add((_1178_v32).Plus(p0), p2)
							}
						}
						return _coll49.ToMap()
					}()
				})(), 3)
				(_nw185).ArraySet1(_1163_v18, 4)
				(_nw185).ArraySet1((_1163_v18).Merge(func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter54 := _dafny.Iterate((_dafny.MultiSetOf((_this).F20(), p0)).Elements()); ; {
						_compr_50, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1179_v33 _dafny.Int
						_1179_v33 = interface{}(_compr_50).(_dafny.Int)
						if (_dafny.MultiSetOf((_this).F20(), p0)).Contains(_1179_v33) {
							_coll50.Add(Companion_Default___.SafeModuloInt(_1179_v33, p2), (_this).F20())
						}
					}
					return _coll50.ToMap()
				}()), 5)
				(_nw185).ArraySet1(_1163_v18, 6)
				(_nw185).ArraySet1(_1163_v18, 7)
				(_nw185).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_dafny.Zero).Minus(p0)), 8)
				(_nw185).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), 9)
				(_nw185).ArraySet1((_1163_v18).Merge(_1163_v18), 10)
				(_nw185).ArraySet1(_1163_v18, 11)
				(_nw185).ArraySet1((_1163_v18).Merge(func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(163), _dafny.IntOfInt64(967))); ; {
						_compr_51, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _1180_v34 _dafny.Int
						_1180_v34 = interface{}(_compr_51).(_dafny.Int)
						if ((_dafny.IntOfInt64(163)).Cmp(_1180_v34) <= 0) && ((_1180_v34).Cmp(_dafny.IntOfInt64(967)) < 0) {
							_coll51.Add((_1180_v34).Plus((_this).F20()), p2)
						}
					}
					return _coll51.ToMap()
				}()), 12)
				(_nw185).ArraySet1(_1163_v18, 13)
				_1177_v35 = _nw185
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1177_v35), 0))
				_ = _index187
				(_1177_v35).ArraySet1(_1163_v18, (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1177_v35), 0))
				_ = _index188
				(_1177_v35).ArraySet1(_1163_v18, (_index188).Int())
				_1132_v0 = (_1132_v0).Update(_dafny.IntOfInt64(289), p1)
			}
			var _1181_v36 D0
			_ = _1181_v36
			_1181_v36 = Companion_D0_.Create_DC2_((_this).F18(), p0, _1164_v19)
			var _1182_v37 *C4
			_ = _1182_v37
			var _nw186 *C4 = New_C4_()
			_ = _nw186
			_nw186.Ctor__(p0, p0, _this.F17(), (p0).Cmp((_dafny.Zero).Minus((_1181_v36).Dtor_cf6())) < 0)
			_1182_v37 = _nw186
		} else {
			var _1183_v38 _dafny.CodePoint
			_ = _1183_v38
			_1183_v38 = _dafny.CodePoint('i')
			var _1184_v39 _dafny.Sequence
			_ = _1184_v39
			_1184_v39 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			var _1185_v40 _dafny.Sequence
			_ = _1185_v40
			_1185_v40 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qlsh"), _1184_v39, _1184_v39, _1184_v39, _1184_v39)
			var _1186_v41 _dafny.Map
			_ = _1186_v41
			_1186_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1183_v38, _dafny.IntOfUint32(((_1185_v40).Select((Companion_Default___.SafeIndex((_this).Fm0((_this).F20(), Companion_Default___.Fm33(_dafny.IntOfUint32((_1184_v39).Cardinality()), (_this).F18(), globalState), globalState), _dafny.IntOfUint32((_1185_v40).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			var _1187_v42 _dafny.Map
			_ = _1187_v42
			_1187_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v18, (_1186_v41).Cardinality())
			if !((_this).Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(p1, globalState), (_this).F20())).Merge(_1187_v42), globalState)) {
				var _arr30 _dafny.Array = _this.F19()
				_ = _arr30
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index189
				(_arr30).ArraySet1((func() bool {
					if (_this).F18() {
						return true
					}
					return (_this).F18()
				})(), (_index189).Int())
				var _1188_v43 _dafny.Map
				_ = _1188_v43
				_1188_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _this.F19())
				var _1189_v44 D5
				_ = _1189_v44
				_1189_v44 = Companion_D5_.Create_DC12_(_1188_v43)
				var _1190_v45 _dafny.MultiSet
				_ = _1190_v45
				_1190_v45 = _dafny.MultiSetOf(_1189_v44)
				var _1191_v46 _dafny.Map
				_ = _1191_v46
				_1191_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1190_v45)
				var _1192_v47 _dafny.MultiSet
				_ = _1192_v47
				_1192_v47 = _dafny.MultiSetOf(p2)
				var _1193_v48 _dafny.Map
				_ = _1193_v48
				_1193_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1191_v46, (_1192_v47).Cardinality())
				var _arr31 _dafny.Array = _this.F19()
				_ = _arr31
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index190
				(_arr31).ArraySet1(((func() _dafny.Int {
					if (_1193_v48).Contains(_1191_v46) {
						return (_1193_v48).Get(_1191_v46).(_dafny.Int)
					}
					return _dafny.IntOfInt64(748)
				})()).Cmp(p0) >= 0, (_index190).Int())
				(globalState).F14 = (_this).F20()
				var _1194_v49 _dafny.Array
				_ = _1194_v49
				var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
				_ = _nw187
				_1194_v49 = _nw187
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1194_v49), 0))
				_ = _index191
				(_1194_v49).ArraySet1CodePoint(_1183_v38, (_index191).Int())
				var _1195_v51 _dafny.Set
				_ = _1195_v51
				_1195_v51 = _dafny.SetOf(_1183_v38, _1183_v38, _1183_v38, _1183_v38, _1183_v38)
				var _1196_v52 _dafny.MultiSet
				_ = _1196_v52
				_1196_v52 = _dafny.MultiSetOf(func() _dafny.Set {
					var _coll52 = _dafny.NewBuilder()
					_ = _coll52
					for _iter56 := _dafny.Iterate((_this.F21()).Keys().Elements()); ; {
						_compr_52, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1197_v50 _dafny.CodePoint
						_1197_v50 = interface{}(_compr_52).(_dafny.CodePoint)
						if (_this.F21()).Contains(_1197_v50) {
							_coll52.Add(_1197_v50)
						}
					}
					return _coll52.ToSet()
				}(), _1195_v51)
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1194_v49), 0))
				_ = _index192
				(_1194_v49).ArraySet1CodePoint(Companion_Default___.Fm33((func() _dafny.Int {
					if (_1196_v52).Contains(_1195_v51) {
						return (_1196_v52).Multiplicity(_1195_v51)
					}
					return p2
				})(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), globalState), (_index192).Int())
				var _arr32 _dafny.Array = _this.F19()
				_ = _arr32
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index193
				(_arr32).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_index193).Int())
				var _1198_v53 D11
				_ = _1198_v53
				_1198_v53 = Companion_D11_.Create_DC26_(p1, p2)
				var _1199_v54 _dafny.Array
				_ = _1199_v54
				var _nwElement0_36 bool = p1
				_ = _nwElement0_36
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(19))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_36, 0)
				(_nw188).ArraySet1(true, 1)
				(_nw188).ArraySet1(true, 2)
				(_nw188).ArraySet1((_this).F18(), 3)
				(_nw188).ArraySet1(true, 4)
				(_nw188).ArraySet1((_this).F18(), 5)
				(_nw188).ArraySet1((_this).F18(), 6)
				(_nw188).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 7)
				(_nw188).ArraySet1((_this).F18(), 8)
				(_nw188).ArraySet1(true, 9)
				(_nw188).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 10)
				(_nw188).ArraySet1(p1, 11)
				(_nw188).ArraySet1(!(p1), 12)
				(_nw188).ArraySet1((_this).F18(), 13)
				(_nw188).ArraySet1((_1198_v53).Dtor_cf48(), 14)
				(_nw188).ArraySet1((_this).F18(), 15)
				(_nw188).ArraySet1(p1, 16)
				(_nw188).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 17)
				(_nw188).ArraySet1(p1, 18)
				_1199_v54 = _nw188
				var _1200_v55 *C3
				_ = _1200_v55
				var _nw189 *C3 = New_C3_()
				_ = _nw189
				_nw189.Ctor__(_1199_v54, _this.F17(), (_1198_v53).Dtor_cf48())
				_1200_v55 = _nw189
				_1200_v55 = (func() *C3 {
					if (p0).Cmp(p0) < 0 {
						return _1200_v55
					}
					return _1200_v55
				})()
			} else {
				var _1201_v56 _dafny.MultiSet
				_ = _1201_v56
				_1201_v56 = _dafny.MultiSetOf(p2)
				var _1202_v57 _dafny.Sequence
				_ = _1202_v57
				_1202_v57 = _dafny.SeqOf((_dafny.Zero).Minus(p2), p0, (_dafny.Zero).Minus((p0).Plus((_this).Fm0(p0, _1183_v38, globalState))), ((_1201_v56).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F18())).Cardinality()), p0, p0, p0, p0))).Cardinality())
				var _1203_v58 _dafny.Map
				_ = _1203_v58
				_1203_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
				var _1204_v59 _dafny.Array
				_ = _1204_v59
				var _nwElement0_37 _dafny.Int = _dafny.IntOfInt64(708)
				_ = _nwElement0_37
				var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(5))
				_ = _nw190
				(_nw190).ArraySet1(_nwElement0_37, 0)
				(_nw190).ArraySet1(((_1203_v58).Cardinality()).Plus((_this).F20()), 1)
				(_nw190).ArraySet1((p0).Plus(p0), 2)
				(_nw190).ArraySet1(((_this).F20()).Times(p0), 3)
				(_nw190).ArraySet1((_dafny.Zero).Minus(p2), 4)
				_1204_v59 = _nw190
				var _1205_v60 _dafny.Sequence
				_ = _1205_v60
				_1205_v60 = _dafny.SeqOf(_1203_v58, _1203_v58)
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))
				_ = _index194
				(_1204_v59).ArraySet1((_dafny.Zero).Minus((((_1205_v60).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1205_v60).Cardinality()))).Uint32()).(_dafny.Map)).Update((_this).F18(), p1)).Cardinality()), (_index194).Int())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))
				_ = _index195
				var _rhs185 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1206_v57 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1207_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_1206_v57).Cardinality())
					}
				})(_1202_v57)))
				_ = _rhs185
				var _rhs186 _dafny.Int = p0
				_ = _rhs186
				var _rhs187 bool = (p2).Cmp((_this).F20()) < 0
				_ = _rhs187
				var _lhs160 _dafny.Array = _1204_v59
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))
				_ = _lhs161
				_1202_v57 = _rhs185
				(_lhs160).ArraySet1(_rhs186, (_lhs161).Int())
				r1 = _rhs187
				var _1208_v61 D11
				_ = _1208_v61
				_1208_v61 = Companion_D11_.Create_DC25_(_dafny.SetOf(_this.F19()))
				var _1209_v62 _dafny.Set
				_ = _1209_v62
				_1209_v62 = _dafny.SetOf(_1208_v61, _1208_v61, Companion_D11_.Create_DC25_(_this.F17()), _1208_v61)
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))
				_ = _index196
				var _rhs188 _dafny.Set = (_1209_v62).Difference(_1209_v62)
				_ = _rhs188
				var _rhs189 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F20(), (_1204_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))).Int()).(_dafny.Int))
				_ = _rhs189
				var _lhs162 _dafny.Array = _1204_v59
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))
				_ = _lhs163
				_1209_v62 = _rhs188
				(_lhs162).ArraySet1(_rhs189, (_lhs163).Int())
				var _arr33 _dafny.Array = _this.F19()
				_ = _arr33
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index197
				(_arr33).ArraySet1((_this).F18(), (_index197).Int())
				var _1210_v63 *C3
				_ = _1210_v63
				var _nw191 *C3 = New_C3_()
				_ = _nw191
				_nw191.Ctor__(_this.F19(), _this.F17(), (_this).F18())
				_1210_v63 = _nw191
				var _1211_v64 _dafny.Map
				_ = _1211_v64
				_1211_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1210_v63, (_this).F20())
				var _arr34 _dafny.Array = _this.F19()
				_ = _arr34
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index198
				(_arr34).ArraySet1(!((_1211_v64).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1210_v63, (_1201_v56).Cardinality()))).Contains(_1210_v63), (_index198).Int())
				var _1212_v65 _dafny.Map
				_ = _1212_v65
				_1212_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1204_v59, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F18()))
				var _rhs190 bool = (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
				_ = _rhs190
				var _rhs191 _dafny.Int = ((_1204_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_1204_v59), 0))).Int()).(_dafny.Int)).Times((_this).F20())
				_ = _rhs191
				var _rhs192 _dafny.Array = _1204_v59
				_ = _rhs192
				var _rhs193 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1184_v39, (Companion_Default___.SafeIndex((_1210_v63).Fm0(_dafny.IntOfUint32((Companion_Default___.Fm40((_this).F18(), p1, globalState)).Cardinality()), _1183_v38, globalState), _dafny.IntOfUint32((_1184_v39).Cardinality()))).Uint32(), _1183_v38), (Companion_Default___.SafeIndex((((func() _dafny.Map {
					if (_1212_v65).Contains(_1204_v59) {
						return (_1212_v65).Get(_1204_v59).(_dafny.Map)
					}
					return _1203_v58
				})()).Merge(_1203_v58)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1184_v39, (Companion_Default___.SafeIndex((_1210_v63).Fm0(_dafny.IntOfUint32((Companion_Default___.Fm40((_this).F18(), p1, globalState)).Cardinality()), _1183_v38, globalState), _dafny.IntOfUint32((_1184_v39).Cardinality()))).Uint32(), _1183_v38)).Cardinality()))).Uint32(), _1183_v38)
				_ = _rhs193
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				_lhs164.F3 = _rhs190
				r2 = _rhs191
				_lhs165.F7 = _rhs192
				_1184_v39 = _rhs193
				r2 = Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfUint32((_1202_v57).Cardinality()))
			}
			var _1213_v66 _dafny.MultiSet
			_ = _1213_v66
			_1213_v66 = _dafny.MultiSetOf((_this).F20())
			(globalState).F15 = _1213_v66
			var _1214_v67 _dafny.Array
			_ = _1214_v67
			var _nw192 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(2))
			_ = _nw192
			_1214_v67 = _nw192
			var _1215_v68 _dafny.Map
			_ = _1215_v68
			_1215_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1214_v67, _this.F17())
			var _1216_v69 D11
			_ = _1216_v69
			_1216_v69 = Companion_D11_.Create_DC25_((func() _dafny.Set {
				if (_1215_v68).Contains(_1214_v67) {
					return (_1215_v68).Get(_1214_v67).(_dafny.Set)
				}
				return _this.F17()
			})())
			var _1217_v70 _dafny.Array
			_ = _1217_v70
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_36
			var _nw193 _dafny.Array
			_ = _nw193
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw193 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) D9 = (func(_1218_v0 _dafny.Map, _1219_p0 _dafny.Int, _1220_p1 bool, _1221_v18 _dafny.Map, _1222_v38 _dafny.CodePoint) func(_dafny.Int) D9 {
					return func(_1223_i4 _dafny.Int) D9 {
						return Companion_D9_.Create_DC20_(_1218_v0, _1219_p0, _1220_p1, _1221_v18, _1222_v38)
					}
				})(_1132_v0, p0, p1, _1163_v18, _1183_v38)
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
			_1217_v70 = _nw193
			var _1224_v71 D9
			_ = _1224_v71
			_1224_v71 = Companion_D9_.Create_DC20_(_1132_v0, p0, p1, (_1163_v18).Update(p0, p0), _1183_v38)
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1217_v70), 0))
			_ = _index199
			(_1217_v70).ArraySet1(_1224_v71, (_index199).Int())
			var _1225_v72 _dafny.Set
			_ = _1225_v72
			_1225_v72 = _dafny.SetOf((_this).F18(), true, p1)
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1217_v70), 0))
			_ = _index200
			var _rhs194 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1184_v39, _1184_v39), _1184_v39)
			_ = _rhs194
			var _rhs195 D11 = _1216_v69
			_ = _rhs195
			var _rhs196 _dafny.Int = (func() _dafny.Int {
				if (_1163_v18).Contains(_dafny.IntOfInt64(551)) {
					return (_1163_v18).Get(_dafny.IntOfInt64(551)).(_dafny.Int)
				}
				return p0
			})()
			_ = _rhs196
			var _rhs197 bool = (_this).Fm1(_1187_v42, globalState)
			_ = _rhs197
			var _rhs198 D9 = Companion_D9_.Create_DC20_(_1132_v0, p2, (_1225_v72).IsSubsetOf(_dafny.SetOf(p1)), _1163_v18, _1183_v38)
			_ = _rhs198
			var _lhs166 _dafny.Array = _1217_v70
			_ = _lhs166
			var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1217_v70), 0))
			_ = _lhs167
			_1184_v39 = _rhs194
			_1216_v69 = _rhs195
			r2 = _rhs196
			r1 = _rhs197
			(_lhs166).ArraySet1(_rhs198, (_lhs167).Int())
			_1132_v0 = (_1132_v0).Update(p2, (_1213_v66).IsSubsetOf(_1213_v66))
			(globalState).F9 = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dca")).Cardinality())).Cmp(p2) >= 0
		}
		r1 = p1
		var _1226_v73 _dafny.CodePoint
		_ = _1226_v73
		_1226_v73 = _dafny.CodePoint('g')
		var _1227_v74 _dafny.Sequence
		_ = _1227_v74
		_1227_v74 = _dafny.UnicodeSeqOfUtf8Bytes("xiu")
		var _1228_v75 D3
		_ = _1228_v75
		_1228_v75 = Companion_D3_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(955))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}(func(_1229_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("xwn"), _1226_v73), _1227_v74, _1227_v74, _dafny.IntOfUint32((_1227_v74).Cardinality()))
		var _source12 D3 = _1228_v75
		_ = _source12
		if _source12.Is_DC9() {
			if true {
				var _1230_v76 _dafny.Sequence
				_ = _1230_v76
				_1230_v76 = _dafny.SeqOf(p0, p0)
				var _1231_v77 *C5
				_ = _1231_v77
				var _nw194 *C5 = New_C5_()
				_ = _nw194
				_nw194.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}(func(_1232_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				})), (_this).F18(), (_dafny.Zero).Minus((_1230_v76).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1230_v76).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1226_v73, (_this).F18()), _this.F19(), _this.F17(), p1)
				_1231_v77 = _nw194
				var _arr35 _dafny.Array = _this.F19()
				_ = _arr35
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index201
				(_arr35).ArraySet1((p2).Cmp(p0) <= 0, (_index201).Int())
				var _1233_v78 _dafny.Map
				_ = _1233_v78
				_1233_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v18, p0)
				var _arr36 _dafny.Array = _this.F19()
				_ = _arr36
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index202
				(_arr36).ArraySet1((_this).Fm1(_1233_v78, globalState), (_index202).Int())
				var _1234_v79 _dafny.Array
				_ = _1234_v79
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_37
				var _nw195 _dafny.Array
				_ = _nw195
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw195 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.Sequence = (func(_1235_v77 *C5, _1236_p0 _dafny.Int, _1237_v0 _dafny.Map, _1238_v76 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1239_i7 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1235_v77).F30(), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_1235_v77).F30(), (func() bool {
								if (_1237_v0).Contains(_1236_p0) {
									return (_1237_v0).Get(_1236_p0).(bool)
								}
								return true
							})()))).Cardinality())).Cardinality()), _1238_v76)
						}
					})(_1231_v77, p0, _1132_v0, _1230_v76)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw195 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw195).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw195).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1234_v79 = _nw195
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1234_v79), 0))
				_ = _index203
				(_1234_v79).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(607))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}(func(_1240_i8 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(940)
				})), (_index203).Int())
				var _1241_v80 _dafny.Set
				_ = _1241_v80
				_1241_v80 = _dafny.SetOf((_this).F20())
				var _1242_v81 _dafny.Array
				_ = _1242_v81
				var _nw196 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
				_ = _nw196
				_1242_v81 = _nw196
				var _1243_v82 _dafny.Sequence
				_ = _1243_v82
				_1243_v82 = _dafny.SeqOf(!(false))
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1234_v79), 0))
				_ = _index204
				var _rhs199 _dafny.Sequence = _1230_v76
				_ = _rhs199
				var _rhs200 _dafny.Set = _1241_v80
				_ = _rhs200
				var _rhs201 _dafny.Array = _1242_v81
				_ = _rhs201
				var _rhs202 _dafny.Int = (p0).Minus(Companion_Default___.SafeModuloInt(p0, p0))
				_ = _rhs202
				var _rhs203 bool = (_1243_v82).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1243_v82).Cardinality()))).Uint32()).(bool)
				_ = _rhs203
				var _lhs168 _dafny.Array = _1234_v79
				_ = _lhs168
				var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1234_v79), 0))
				_ = _lhs169
				var _lhs170 *GlobalState = globalState
				_ = _lhs170
				var _lhs171 *GlobalState = globalState
				_ = _lhs171
				(_lhs168).ArraySet1(_rhs199, (_lhs169).Int())
				_1241_v80 = _rhs200
				_1242_v81 = _rhs201
				_lhs170.F14 = _rhs202
				_lhs171.F9 = _rhs203
				var _1244_v83 _dafny.Map
				_ = _1244_v83
				_1244_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("mjn"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mjn")).Cardinality()))).Uint32(), _1226_v73), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("mjn"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mjn")).Cardinality()))).Uint32(), _1226_v73)).Cardinality()))).Uint32(), _1226_v73), _1227_v74))
				var _rhs204 _dafny.Map = ((_1244_v83).Merge(_1244_v83)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1231_v77).F29()))
				_ = _rhs204
				var _rhs205 _dafny.Int = (p0).Times(p0)
				_ = _rhs205
				var _rhs206 _dafny.Sequence = (_this).Fm25((_this).F18(), p1, p2, (p1) == ((_this).F18()), globalState)
				_ = _rhs206
				var _rhs207 bool = (func() bool {
					if (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool) {
						return (_this).F18()
					}
					return (_1231_v77).Fm1(_1233_v78, globalState)
				})()
				_ = _rhs207
				var _lhs172 *GlobalState = globalState
				_ = _lhs172
				var _lhs173 *GlobalState = globalState
				_ = _lhs173
				_1244_v83 = _rhs204
				_lhs172.F14 = _rhs205
				_1243_v82 = _rhs206
				_lhs173.F3 = _rhs207
				var _1245_v84 D3
				_ = _1245_v84
				_1245_v84 = Companion_D3_.Create_DC9_()
				var _1246_v85 _dafny.Sequence
				_ = _1246_v85
				_1246_v85 = _dafny.SeqOf(_1245_v84)
				var _1247_v86 D5
				_ = _1247_v86
				_1247_v86 = Companion_D5_.Create_DC13_(!((_this).Fm1(_1233_v78, globalState)), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F20())), _1246_v85)
				var _1248_v87 _dafny.Array
				_ = _1248_v87
				var _nwElement0_38 D5 = Companion_Default___.Fm41(p2, (_this).F18(), (_1231_v77).F29(), _1226_v73, globalState)
				_ = _nwElement0_38
				var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(8))
				_ = _nw197
				(_nw197).ArraySet1(_nwElement0_38, 0)
				(_nw197).ArraySet1(Companion_Default___.Fm41(_dafny.IntOfUint32(((_1231_v77).F29()).Cardinality()), (_this).F18(), _dafny.UnicodeSeqOfUtf8Bytes("c"), _dafny.CodePoint('g'), globalState), 1)
				(_nw197).ArraySet1(_1247_v86, 2)
				(_nw197).ArraySet1(_1247_v86, 3)
				(_nw197).ArraySet1(Companion_Default___.Fm41((_this).F20(), p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-486))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1249_v73 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1250_i9 _dafny.Int) _dafny.CodePoint {
						return _1249_v73
					}
				})(_1226_v73))), _dafny.CodePoint('x'), globalState), 4)
				(_nw197).ArraySet1(_1247_v86, 5)
				(_nw197).ArraySet1(_1247_v86, 6)
				(_nw197).ArraySet1(_1247_v86, 7)
				_1248_v87 = _nw197
				var _1251_v88 _dafny.Sequence
				_ = _1251_v88
				_1251_v88 = _1246_v85
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1248_v87), 0))
				_ = _index205
				(_1248_v87).ArraySet1(Companion_D5_.Create_DC13_(p1, p2, _1251_v88), (_index205).Int())
				var _1252_v89 _dafny.Map
				_ = _1252_v89
				_1252_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_1231_v77).F30())
				var _1253_v90 _dafny.Array
				_ = _1253_v90
				var _nwElement0_39 _dafny.Int = p2
				_ = _nwElement0_39
				var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(12))
				_ = _nw198
				(_nw198).ArraySet1(_nwElement0_39, 0)
				(_nw198).ArraySet1((_this).Fm0(p0, _dafny.CodePoint('g'), globalState), 1)
				(_nw198).ArraySet1((_dafny.Zero).Minus((_1252_v89).Cardinality()), 2)
				(_nw198).ArraySet1((_this).F20(), 3)
				(_nw198).ArraySet1((_this).F20(), 4)
				(_nw198).ArraySet1(p2, 5)
				(_nw198).ArraySet1(_dafny.IntOfInt64(193), 6)
				(_nw198).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1231_v77).F29(), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32(((_1231_v77).F29()).Cardinality()))).Uint32(), _1226_v73)).Cardinality()), 7)
				(_nw198).ArraySet1(p0, 8)
				(_nw198).ArraySet1((_this).F20(), 9)
				(_nw198).ArraySet1(p2, 10)
				(_nw198).ArraySet1(_dafny.IntOfInt64(18), 11)
				_1253_v90 = _nw198
				var _1254_v91 _dafny.MultiSet
				_ = _1254_v91
				_1254_v91 = _dafny.MultiSetOf(_1253_v90, _1253_v90)
				var _pat_let_tv12 = _1254_v91
				_ = _pat_let_tv12
				var _pat_let_tv13 = _1253_v90
				_ = _pat_let_tv13
				var _pat_let_tv14 = _1254_v91
				_ = _pat_let_tv14
				var _pat_let_tv15 = _1253_v90
				_ = _pat_let_tv15
				var _pat_let_tv16 = _1251_v88
				_ = _pat_let_tv16
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1248_v87), 0))
				_ = _index206
				(_1248_v87).ArraySet1(func(_pat_let21_0 D5) D5 {
					return func(_1255_dt__update__tmp_h0 D5) D5 {
						return func(_pat_let22_0 _dafny.Int) D5 {
							return func(_1256_dt__update_hcf24_h0 _dafny.Int) D5 {
								return func(_pat_let23_0 _dafny.Sequence) D5 {
									return func(_1257_dt__update_hcf25_h0 _dafny.Sequence) D5 {
										return Companion_D5_.Create_DC13_((_1255_dt__update__tmp_h0).Dtor_cf23(), _1256_dt__update_hcf24_h0, _1257_dt__update_hcf25_h0)
									}(_pat_let23_0)
								}(_pat_let_tv16)
							}(_pat_let22_0)
						}(Companion_Default___.SafeModuloInt((_this).F20(), (func() _dafny.Int {
							if (_pat_let_tv12).Contains(_pat_let_tv13) {
								return (_pat_let_tv14).Multiplicity(_pat_let_tv15)
							}
							return _dafny.IntOfInt64(913)
						})()))
					}(_pat_let21_0)
				}(_1247_v86), (_index206).Int())
			} else {
				_1227_v74 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(654))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1258_p2 _dafny.Int, _1259_v74 _dafny.Sequence, _1260_v73 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1261_i10 _dafny.Int) _dafny.CodePoint {
						return (Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("gf"), _1258_p2, _dafny.IntOfInt64(-503), (Companion_D0_.Create_DC1_(_1259_v74, (_this).F20(), _1258_p2, _1260_v73)).Dtor_cf4())).Dtor_cf4()
					}
				})(p2, _1227_v74, _1226_v73)))
				var _arr37 _dafny.Array = _this.F19()
				_ = _arr37
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index207
				(_arr37).ArraySet1((p0).Cmp((_this).F20()) != 0, (_index207).Int())
				var _arr38 _dafny.Array = _this.F19()
				_ = _arr38
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index208
				(_arr38).ArraySet1(p1, (_index208).Int())
				var _1262_v92 _dafny.Set
				_ = _1262_v92
				_1262_v92 = _dafny.SetOf(p1)
				var _1263_v93 _dafny.Map
				_ = _1263_v93
				_1263_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(p1))
				var _1264_v94 _dafny.Map
				_ = _1264_v94
				_1264_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1263_v93).Cardinality())
				var _1265_v95 _dafny.Sequence
				_ = _1265_v95
				_1265_v95 = _dafny.SeqOf((_this).F20())
				var _1266_v96 _dafny.Sequence
				_ = _1266_v96
				_1266_v96 = _dafny.SeqOf(p0, ((_1264_v94).Update((_this).F18(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_this).F18(), (_this).F18()))).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_1265_v95).Cardinality()))
				var _1267_v97 _dafny.Sequence
				_ = _1267_v97
				_1267_v97 = _dafny.SeqOf(_dafny.IntOfUint32((_1227_v74).Cardinality()), (_1266_v96).Select((Companion_Default___.SafeIndex((_this).Fm0((_this).F20(), _1226_v73, globalState), _dafny.IntOfUint32((_1266_v96).Cardinality()))).Uint32()).(_dafny.Int))
				var _1268_v98 *C4
				_ = _1268_v98
				var _nw199 *C4 = New_C4_()
				_ = _nw199
				_nw199.Ctor__((_dafny.MultiSetOf(_dafny.IntOfInt64(241))).Cardinality(), (_dafny.Zero).Minus((_this).Fm0((_1163_v18).Cardinality(), _1226_v73, globalState)), _this.F17(), !((_this).F18()))
				_1268_v98 = _nw199
				var _1269_v99 _dafny.Set
				_ = _1269_v99
				_1269_v99 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v98, Companion_Default___.Fm37((_this).F18(), _dafny.IntOfInt64(384), p1, false, globalState))).Cardinality())
				var _1270_v100 _dafny.Array
				_ = _1270_v100
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_38
				var _nw200 _dafny.Array
				_ = _nw200
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw200 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Int = func(_1271_i11 _dafny.Int) _dafny.Int {
						return (_1271_i11).Plus((_this).F20())
					}
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw200 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw200).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw200).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1270_v100 = _nw200
				var _1272_v101 D0
				_ = _1272_v101
				_1272_v101 = Companion_D0_.Create_DC2_(p1, (_1268_v98).F25(), _1270_v100)
				var _1273_v102 _dafny.Array
				_ = _1273_v102
				var _nwElement0_40 bool = p1
				_ = _nwElement0_40
				var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(20))
				_ = _nw201
				(_nw201).ArraySet1(_nwElement0_40, 0)
				(_nw201).ArraySet1(p1, 1)
				(_nw201).ArraySet1((_this).F18(), 2)
				(_nw201).ArraySet1(false, 3)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 4)
				(_nw201).ArraySet1(p1, 5)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 6)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 7)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 8)
				(_nw201).ArraySet1(true, 9)
				(_nw201).ArraySet1(true, 10)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 11)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 12)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 13)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 14)
				(_nw201).ArraySet1(!(false), 15)
				(_nw201).ArraySet1(p1, 16)
				(_nw201).ArraySet1((_this).F18(), 17)
				(_nw201).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 18)
				(_nw201).ArraySet1((_1272_v101).Dtor_cf5(), 19)
				_1273_v102 = _nw201
				var _1274_v103 _dafny.Sequence
				_ = _1274_v103
				_1274_v103 = _dafny.SeqOf(_1273_v102, _1273_v102)
				var _1275_v104 _dafny.Array
				_ = _1275_v104
				var _nwElement0_41 _dafny.Int = (_1269_v99).Cardinality()
				_ = _nwElement0_41
				var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(6))
				_ = _nw202
				(_nw202).ArraySet1(_nwElement0_41, 0)
				(_nw202).ArraySet1(_dafny.IntOfUint32((_1266_v96).Cardinality()), 1)
				(_nw202).ArraySet1((_1268_v98).F25(), 2)
				(_nw202).ArraySet1((_this).F20(), 3)
				(_nw202).ArraySet1(_dafny.IntOfUint32((_1274_v103).Cardinality()), 4)
				(_nw202).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uosnx")).Cardinality()), 5)
				_1275_v104 = _nw202
				var _1276_v105 D0
				_ = _1276_v105
				_1276_v105 = Companion_D0_.Create_DC2_((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_1262_v92).Cardinality(), _1275_v104)
				var _arr39 _dafny.Array = _this.F19()
				_ = _arr39
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index209
				var _rhs208 _dafny.Int = p2
				_ = _rhs208
				var _rhs209 _dafny.Array = (_1272_v101).Dtor_cf7()
				_ = _rhs209
				var _rhs210 _dafny.Set = _1262_v92
				_ = _rhs210
				var _rhs211 _dafny.Sequence = _1265_v95
				_ = _rhs211
				var _rhs212 bool = !((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
				_ = _rhs212
				var _lhs174 *GlobalState = globalState
				_ = _lhs174
				var _lhs175 _dafny.Array = _this.F19()
				_ = _lhs175
				var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))
				_ = _lhs176
				r2 = _rhs208
				_lhs174.F7 = _rhs209
				_1262_v92 = _rhs210
				_1267_v97 = _rhs211
				(_lhs175).ArraySet1(_rhs212, (_lhs176).Int())
				var _arr40 _dafny.Array = _this.F19()
				_ = _arr40
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_this.F19()), 0))
				_ = _index210
				(_arr40).ArraySet1((_this).F18(), (_index210).Int())
				var _1277_v106 _dafny.Array
				_ = _1277_v106
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
				_ = _nw203
				_1277_v106 = _nw203
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1277_v106), 0))
				_ = _index211
				(_1277_v106).ArraySet1CodePoint(_1226_v73, (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1277_v106), 0))
				_ = _index212
				(_1277_v106).ArraySet1CodePoint(_1226_v73, (_index212).Int())
			}
			(globalState).F9 = (_this).F18()
			var _1278_v107 _dafny.Array
			_ = _1278_v107
			var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw204
			_1278_v107 = _nw204
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1278_v107), 0))
			_ = _index213
			(_1278_v107).ArraySet1((p0).Minus(p0), (_index213).Int())
			var _1279_v108 _dafny.MultiSet
			_ = _1279_v108
			_1279_v108 = _dafny.MultiSetOf(p1)
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1278_v107), 0))
			_ = _index214
			(_1278_v107).ArraySet1((_1279_v108).Cardinality(), (_index214).Int())
			var _1280_v109 _dafny.Map
			_ = _1280_v109
			_1280_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _this.F19())
			var _1281_v110 D5
			_ = _1281_v110
			_1281_v110 = Companion_D5_.Create_DC12_(_1280_v109)
			var _1282_v111 _dafny.Sequence
			_ = _1282_v111
			_1282_v111 = _dafny.SeqOf(_1281_v110, Companion_D5_.Create_DC12_(_1280_v109))
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1278_v107), 0))
			_ = _index215
			var _rhs213 bool = p1
			_ = _rhs213
			var _rhs214 _dafny.Int = (p0).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1282_v111, _dafny.Companion_Sequence_.Update(_1282_v111, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1282_v111).Cardinality()))).Uint32(), _1281_v110))).Cardinality()))
			_ = _rhs214
			var _rhs215 _dafny.Int = Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus((_this).F20()))
			_ = _rhs215
			var _rhs216 _dafny.CodePoint = _1226_v73
			_ = _rhs216
			var _rhs217 _dafny.CodePoint = _1226_v73
			_ = _rhs217
			var _lhs177 *GlobalState = globalState
			_ = _lhs177
			var _lhs178 _dafny.Array = _1278_v107
			_ = _lhs178
			var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1278_v107), 0))
			_ = _lhs179
			_lhs177.F3 = _rhs213
			(_lhs178).ArraySet1(_rhs214, (_lhs179).Int())
			r2 = _rhs215
			_1226_v73 = _rhs216
			_1226_v73 = _rhs217
		} else if _source12.Is_DC10() {
			var _1283___mcc_h4 _dafny.Sequence = _source12.Get_().(D3_DC10).Cf16
			_ = _1283___mcc_h4
			var _1284___mcc_h5 bool = _source12.Get_().(D3_DC10).Cf17
			_ = _1284___mcc_h5
			var _1285___mcc_h6 _dafny.Sequence = _source12.Get_().(D3_DC10).Cf18
			_ = _1285___mcc_h6
			var _1286___mcc_h7 _dafny.Sequence = _source12.Get_().(D3_DC10).Cf19
			_ = _1286___mcc_h7
			var _1287___mcc_h8 _dafny.Int = _source12.Get_().(D3_DC10).Cf20
			_ = _1287___mcc_h8
			var _1288_cf20 _dafny.Int = _1287___mcc_h8
			_ = _1288_cf20
			var _1289_cf19 _dafny.Sequence = _1286___mcc_h7
			_ = _1289_cf19
			var _1290_cf18 _dafny.Sequence = _1285___mcc_h6
			_ = _1290_cf18
			var _1291_cf17 bool = _1284___mcc_h5
			_ = _1291_cf17
			var _1292_cf16 _dafny.Sequence = _1283___mcc_h4
			_ = _1292_cf16
			(globalState).F14 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(617), p2)
			var _1293_v112 *C0
			_ = _1293_v112
			var _nw205 *C0 = New_C0_()
			_ = _nw205
			_nw205.Ctor__()
			_1293_v112 = _nw205
			var _1294_v113 _dafny.Sequence
			_ = _1294_v113
			_1294_v113 = _dafny.SeqOf(Companion_Default___.Fm30(p1, globalState), _1163_v18, Companion_Default___.Fm30((_this).F18(), globalState))
			var _1295_v114 _dafny.Sequence
			_ = _1295_v114
			_1295_v114 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(((_1294_v113).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1294_v113).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfInt64(944)), p0, (_this).Fm0((_this).F20(), _1226_v73, globalState))
			_1288_cf20 = _dafny.IntOfUint32((_1295_v114).Cardinality())
			var _1296_v115 _dafny.Set
			_ = _1296_v115
			_1296_v115 = _dafny.SetOf(_1288_cf20)
			var _1297_v116 _dafny.Map
			_ = _1297_v116
			_1297_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v18, (_dafny.Zero).Minus((_this).F20()))
			var _1298_v117 _dafny.Sequence
			_ = _1298_v117
			_1298_v117 = _dafny.SeqOf((_this).Fm1(_1297_v116, globalState))
			_1296_v115 = (Companion_Default___.Fm42(_dafny.MultiSetOf(true), p1, _1298_v117, globalState)).Intersection(_1296_v115)
		} else {
			var _1299___mcc_h9 _dafny.Sequence = _source12.Get_().(D3_DC8).Cf15
			_ = _1299___mcc_h9
			var _1300_cf15 _dafny.Sequence = _1299___mcc_h9
			_ = _1300_cf15
			(globalState).F14 = (_this).F20()
			var _1301_v118 _dafny.Array
			_ = _1301_v118
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_39
			var _nw206 _dafny.Array
			_ = _nw206
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw206 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Sequence = func(_1302_i12 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F18()), _dafny.SeqOf((_this).F18()))
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw206 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw206).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw206).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1301_v118 = _nw206
			var _1303_v119 _dafny.Map
			_ = _1303_v119
			_1303_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _1304_v120 D9
			_ = _1304_v120
			_1304_v120 = Companion_D9_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1303_v119).Cardinality(), false), _dafny.IntOfInt64(-81), (_this).F18(), _1163_v18, _1226_v73)
			var _1305_v121 _dafny.Sequence
			_ = _1305_v121
			_1305_v121 = _dafny.SeqOf(!(p1))
			var _1306_v122 _dafny.Set
			_ = _1306_v122
			_1306_v122 = _dafny.SetOf((_this).F20(), p2, p0)
			var _1307_v123 _dafny.Map
			_ = _1307_v123
			_1307_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SeqOf(p1))
			var _nwElement0_42 _dafny.Sequence = _dafny.SeqOf(p1, p1, p1, (_1304_v120).Dtor_cf38())
			_ = _nwElement0_42
			var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(17))
			_ = _nw207
			(_nw207).ArraySet1(_nwElement0_42, 0)
			(_nw207).ArraySet1(_1305_v121, 1)
			(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F18()), _1305_v121), 2)
			(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1305_v121, _1305_v121), 3)
			(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1305_v121, _1305_v121), 4)
			(_nw207).ArraySet1(_1305_v121, 5)
			(_nw207).ArraySet1(_1305_v121, 6)
			(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _1305_v121), 7)
			(_nw207).ArraySet1((_this).Fm25(!((_this).F18()), p1, (_1306_v122).Cardinality(), (_this).F18(), globalState), 8)
			(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1305_v121, _dafny.SeqOf(p1, p1)), 9)
			(_nw207).ArraySet1(_dafny.SeqOf((_this).F18()), 10)
			(_nw207).ArraySet1(_1305_v121, 11)
			(_nw207).ArraySet1(_1305_v121, 12)
			(_nw207).ArraySet1(_1305_v121, 13)
			(_nw207).ArraySet1(_1305_v121, 14)
			(_nw207).ArraySet1(_1305_v121, 15)
			(_nw207).ArraySet1((func() _dafny.Sequence {
				if (_1307_v123).Contains(false) {
					return (_1307_v123).Get(false).(_dafny.Sequence)
				}
				return _1305_v121
			})(), 16)
			_1301_v118 = _nw207
			var _1308_v124 _dafny.Array
			_ = _1308_v124
			var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw208
			_1308_v124 = _nw208
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_1308_v124), 0))
			_ = _index216
			(_1308_v124).ArraySet1(_this.F19(), (_index216).Int())
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_1308_v124), 0))
			_ = _index217
			(_1308_v124).ArraySet1(_this.F19(), (_index217).Int())
			var _1309_v125 _dafny.Array
			_ = _1309_v125
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_40
			var _nw209 _dafny.Array
			_ = _nw209
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw209 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) _dafny.Int = func(_1310_i13 _dafny.Int) _dafny.Int {
					return (_1310_i13).Plus(_dafny.IntOfInt64(647))
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw209 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw209).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw209).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1309_v125 = _nw209
			(globalState).F7 = _1309_v125
		}
		var _1311_v126 _dafny.Set
		_ = _1311_v126
		_1311_v126 = _dafny.SetOf((p0).Cmp(_dafny.IntOfUint32(((_this).Fm25((_this).F18(), p1, p0, (_this).F18(), globalState)).Cardinality())) >= 0)
		var _1312_v127 _dafny.Map
		_ = _1312_v127
		_1312_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v18, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), p0)).Cardinality())
		var _1313_v128 _dafny.Sequence
		_ = _1313_v128
		_1313_v128 = _dafny.SeqOf(_1311_v126, _dafny.SetOf(p1), _1311_v126, Companion_Default___.Fm37((_this).F18(), (_this).F20(), p1, (_this).Fm1(_1312_v127, globalState), globalState))
		_1311_v126 = (_1313_v128).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1313_v128).Cardinality()))).Uint32()).(_dafny.Set)
		var _1314_v129 _dafny.Sequence
		_ = _1314_v129
		_1314_v129 = _dafny.SeqOf((_this).F18())
		var _1315_v130 *C1
		_ = _1315_v130
		var _nw210 *C1 = New_C1_()
		_ = _nw210
		_nw210.Ctor__()
		_1315_v130 = _nw210
		var _1316_v131 _dafny.Set
		_ = _1316_v131
		_1316_v131 = _dafny.SetOf(_1315_v130, _1315_v130)
		var _1317_v132 _dafny.Map
		_ = _1317_v132
		_1317_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1316_v131)
		var _1318_v133 _dafny.Map
		_ = _1318_v133
		_1318_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1227_v74, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1314_v129).Cardinality()), _dafny.IntOfUint32((_1227_v74).Cardinality()))).Uint32(), _1226_v73), (func() _dafny.Set {
			if (_1317_v132).Contains((_this).F18()) {
				return (_1317_v132).Get((_this).F18()).(_dafny.Set)
			}
			return _1316_v131
		})())
		var _1319_i14 _dafny.Int
		_ = _1319_i14
		_1319_i14 = _dafny.Zero
		{
			for !(_1318_v133).Contains(_dafny.Companion_Sequence_.Concatenate(_1227_v74, _1227_v74)) {
				{
					if (_1319_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1319_i14 = (_1319_i14).Plus(_dafny.One)
					(globalState).F14 = p2
					var _1320_v134 _dafny.Array
					_ = _1320_v134
					var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
					_ = _nw211
					_1320_v134 = _nw211
					var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_1320_v134), 0))
					_ = _index218
					(_1320_v134).ArraySet1CodePoint(_1226_v73, (_index218).Int())
					var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_1320_v134), 0))
					_ = _index219
					(_1320_v134).ArraySet1CodePoint(_1226_v73, (_index219).Int())
					r2 = p2
					var _1321_v135 D3
					_ = _1321_v135
					_1321_v135 = Companion_D3_.Create_DC9_()
					var _1322_v136 _dafny.Sequence
					_ = _1322_v136
					_1322_v136 = _dafny.SeqOf(Companion_D3_.Create_DC9_(), _1321_v135)
					var _1323_v137 _dafny.Sequence
					_ = _1323_v137
					_1323_v137 = _1322_v136
					var _source13 D5 = Companion_D5_.Create_DC13_(p1, ((_this).F20()).Times(p2), _1323_v137)
					_ = _source13
					if _source13.Is_DC13() {
						var _1324___mcc_h10 bool = _source13.Get_().(D5_DC13).Cf23
						_ = _1324___mcc_h10
						var _1325___mcc_h11 _dafny.Int = _source13.Get_().(D5_DC13).Cf24
						_ = _1325___mcc_h11
						var _1326___mcc_h12 _dafny.Sequence = _source13.Get_().(D5_DC13).Cf25
						_ = _1326___mcc_h12
						var _1327_cf25 _dafny.Sequence = _1326___mcc_h12
						_ = _1327_cf25
						var _1328_cf24 _dafny.Int = _1325___mcc_h11
						_ = _1328_cf24
						var _1329_cf23 bool = _1324___mcc_h10
						_ = _1329_cf23
						(globalState).F14 = _dafny.IntOfInt64(-602)
						var _1330_v138 _dafny.MultiSet
						_ = _1330_v138
						_1330_v138 = _dafny.MultiSetOf((_this).F20(), (_this).F20(), p0, p2)
						(globalState).F14 = (func() _dafny.Int {
							if (_1330_v138).Contains((_this).F20()) {
								return (_1330_v138).Multiplicity((_this).F20())
							}
							return p0
						})()
						var _1331_v139 *C0
						_ = _1331_v139
						var _nw212 *C0 = New_C0_()
						_ = _nw212
						_nw212.Ctor__()
						_1331_v139 = _nw212
						var _1332_v140 _dafny.Sequence
						_ = _1332_v140
						_1332_v140 = _dafny.SeqOf((_1132_v0).Cardinality())
						var _1333_v141 _dafny.Map
						_ = _1333_v141
						_1333_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1332_v140).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1332_v140).Cardinality()))).Uint32()).(_dafny.Int), _1331_v139)
						var _1334_v142 T0
						_ = _1334_v142
						var _nw213 *C2 = New_C2_()
						_ = _nw213
						_nw213.Ctor__(_this.F17(), _1329_cf23)
						_1334_v142 = _nw213
						var _1335_v143 _dafny.Sequence
						_ = _1335_v143
						_1335_v143 = _dafny.SeqOf(_1334_v142)
						var _1336_v144 _dafny.Sequence
						_ = _1336_v144
						_1336_v144 = _dafny.SeqOf(_1334_v142, _1334_v142, _1334_v142, (_1335_v143).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1335_v143).Cardinality()))).Uint32()).(T0))
						var _1337_v145 _dafny.Sequence
						_ = _1337_v145
						_1337_v145 = _dafny.SeqOf(_1331_v139, _1331_v139)
						var _1338_v146 _dafny.Array
						_ = _1338_v146
						var _nwElement0_43 *C0 = _1331_v139
						_ = _nwElement0_43
						var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(29))
						_ = _nw214
						(_nw214).ArraySet1(_nwElement0_43, 0)
						(_nw214).ArraySet1(_1331_v139, 1)
						(_nw214).ArraySet1(_1331_v139, 2)
						(_nw214).ArraySet1(_1331_v139, 3)
						(_nw214).ArraySet1(_1331_v139, 4)
						(_nw214).ArraySet1(_1331_v139, 5)
						(_nw214).ArraySet1(_1331_v139, 6)
						(_nw214).ArraySet1(_1331_v139, 7)
						(_nw214).ArraySet1(_1331_v139, 8)
						(_nw214).ArraySet1(_1331_v139, 9)
						(_nw214).ArraySet1(_1331_v139, 10)
						(_nw214).ArraySet1(_1331_v139, 11)
						(_nw214).ArraySet1(_1331_v139, 12)
						(_nw214).ArraySet1(_1331_v139, 13)
						(_nw214).ArraySet1(_1331_v139, 14)
						(_nw214).ArraySet1((func() *C0 {
							if (_1333_v141).Contains((_dafny.MultiSetFromSeq(_1335_v143)).Cardinality()) {
								return (_1333_v141).Get((_dafny.MultiSetFromSeq(_1335_v143)).Cardinality()).(*C0)
							}
							return _1331_v139
						})(), 15)
						(_nw214).ArraySet1(_1331_v139, 16)
						(_nw214).ArraySet1(_1331_v139, 17)
						(_nw214).ArraySet1(_1331_v139, 18)
						(_nw214).ArraySet1(_1331_v139, 19)
						(_nw214).ArraySet1(_1331_v139, 20)
						(_nw214).ArraySet1(_1331_v139, 21)
						(_nw214).ArraySet1(_1331_v139, 22)
						(_nw214).ArraySet1(_1331_v139, 23)
						(_nw214).ArraySet1(_1331_v139, 24)
						(_nw214).ArraySet1(_1331_v139, 25)
						(_nw214).ArraySet1((func() *C0 {
							if (_this).F18() {
								return _1331_v139
							}
							return _1331_v139
						})(), 26)
						(_nw214).ArraySet1(_1331_v139, 27)
						(_nw214).ArraySet1((_1337_v145).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-362), _dafny.IntOfUint32((_1337_v145).Cardinality()))).Uint32()).(*C0), 28)
						_1338_v146 = _nw214
						_1338_v146 = _1338_v146
						var _1339_v147 _dafny.Map
						_ = _1339_v147
						_1339_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1227_v74)
						var _1340_v148 _dafny.Array
						_ = _1340_v148
						var _nwElement0_44 _dafny.Int = p2
						_ = _nwElement0_44
						var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(15))
						_ = _nw215
						(_nw215).ArraySet1(_nwElement0_44, 0)
						(_nw215).ArraySet1((_this).F20(), 1)
						(_nw215).ArraySet1(p2, 2)
						(_nw215).ArraySet1((_this).F20(), 3)
						(_nw215).ArraySet1((_this).F20(), 4)
						(_nw215).ArraySet1(p2, 5)
						(_nw215).ArraySet1((_1339_v147).Cardinality(), 6)
						(_nw215).ArraySet1(((_1311_v126).Intersection(_1311_v126)).Cardinality(), 7)
						(_nw215).ArraySet1(_1328_cf24, 8)
						(_nw215).ArraySet1((_this).F20(), 9)
						(_nw215).ArraySet1(_1328_cf24, 10)
						(_nw215).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1328_cf24)), 11)
						(_nw215).ArraySet1((_1332_v140).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-78))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg77 _dafny.Int) interface{} {
								return coer77(arg77)
							}
						}((func(_1341_v74 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_1342_i15 _dafny.Int) _dafny.Sequence {
								return _1341_v74
							}
						})(_1227_v74)))).Cardinality()), _dafny.IntOfUint32((_1332_v140).Cardinality()))).Uint32()).(_dafny.Int), 12)
						(_nw215).ArraySet1((_this).F20(), 13)
						(_nw215).ArraySet1((func() _dafny.Int {
							if (_1334_v142).F18() {
								return p2
							}
							return (_dafny.Zero).Minus(p0)
						})(), 14)
						_1340_v148 = _nw215
						var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_1340_v148), 0))
						_ = _index220
						(_1340_v148).ArraySet1(_dafny.IntOfUint32((_1227_v74).Cardinality()), (_index220).Int())
						var _1343_v149 _dafny.MultiSet
						_ = _1343_v149
						_1343_v149 = _dafny.MultiSetOf(_1227_v74, _1227_v74, _1227_v74)
						var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_1340_v148), 0))
						_ = _index221
						(_1340_v148).ArraySet1(((_this).F20()).Times((_dafny.IntOfInt64(187)).Minus((_1343_v149).Cardinality())), (_index221).Int())
					} else if _source13.Is_DC14() {
						var _1344___mcc_h13 bool = _source13.Get_().(D5_DC14).Cf26
						_ = _1344___mcc_h13
						var _1345___mcc_h14 _dafny.Int = _source13.Get_().(D5_DC14).Cf27
						_ = _1345___mcc_h14
						var _1346___mcc_h15 _dafny.Int = _source13.Get_().(D5_DC14).Cf28
						_ = _1346___mcc_h15
						var _1347_cf28 _dafny.Int = _1346___mcc_h15
						_ = _1347_cf28
						var _1348_cf27 _dafny.Int = _1345___mcc_h14
						_ = _1348_cf27
						var _1349_cf26 bool = _1344___mcc_h13
						_ = _1349_cf26
						(_this).F19_set_(_this.F19())
						var _arr41 _dafny.Array = _this.F19()
						_ = _arr41
						var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_this.F19()), 0))
						_ = _index222
						(_arr41).ArraySet1((_this).F18(), (_index222).Int())
						var _arr42 _dafny.Array = _this.F19()
						_ = _arr42
						var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_this.F19()), 0))
						_ = _index223
						(_arr42).ArraySet1(true, (_index223).Int())
						var _1350_v150 _dafny.Sequence
						_ = _1350_v150
						_1350_v150 = _dafny.SeqOf(p2, (_this).F20())
						var _1351_v151 _dafny.Map
						_ = _1351_v151
						_1351_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1227_v74, _1350_v150)
						_1351_v151 = (_1351_v151).Update(_dafny.Companion_Sequence_.Concatenate(_1227_v74, _dafny.UnicodeSeqOfUtf8Bytes("kla")), _dafny.Companion_Sequence_.Concatenate(_1350_v150, _dafny.SeqOf((_this).F20(), p2)))
						var _1352_v152 _dafny.MultiSet
						_ = _1352_v152
						_1352_v152 = _dafny.MultiSetOf(_1349_cf26)
						(globalState).F14 = (_this).Fm0((_1352_v152).Cardinality(), (_1320_v134).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_1320_v134), 0))).Int()), globalState)
					} else {
						var _1353___mcc_h16 _dafny.Map = _source13.Get_().(D5_DC12).Cf22
						_ = _1353___mcc_h16
						var _1354_cf22 _dafny.Map = _1353___mcc_h16
						_ = _1354_cf22
						(globalState).F14 = (_dafny.Zero).Minus(p2)
						var _1355_v153 _dafny.Array
						_ = _1355_v153
						var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
						_ = _nw216
						_1355_v153 = _nw216
						var _1356_v154 D7
						_ = _1356_v154
						_1356_v154 = Companion_D7_.Create_DC17_(_1355_v153, (_this).F18(), p2)
						var _1357_v155 _dafny.Map
						_ = _1357_v155
						_1357_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1321_v135, _1356_v154)
						var _1358_v156 _dafny.Array
						_ = _1358_v156
						var _len0_41 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_41
						var _nw217 _dafny.Array
						_ = _nw217
						if _len0_41.Cmp(_dafny.Zero) == 0 {
							_nw217 = _dafny.NewArray(_len0_41)
						} else {
							var _init41 func(_dafny.Int) _dafny.Int = (func(_1359_v73 _dafny.CodePoint, _1360_v134 _dafny.Array) func(_dafny.Int) _dafny.Int {
								return func(_1361_i16 _dafny.Int) _dafny.Int {
									return (_1361_i16).Plus((_dafny.SetOf(_1359_v73, (_1360_v134).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_1360_v134), 0))).Int()))).Cardinality())
								}
							})(_1226_v73, _1320_v134)
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
						_1358_v156 = _nw217
						var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1358_v156), 0))
						_ = _index224
						(_1358_v156).ArraySet1(p0, (_index224).Int())
						var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1358_v156), 0))
						_ = _index225
						(_1358_v156).ArraySet1((_this).F20(), (_index225).Int())
						var _1362_v157 _dafny.Sequence
						_ = _1362_v157
						_1362_v157 = _dafny.SeqOf(_1357_v155)
						var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1358_v156), 0))
						_ = _index226
						var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1358_v156), 0))
						_ = _index227
						var _rhs218 _dafny.Map = ((_1362_v157).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1314_v129).Cardinality()), _dafny.IntOfUint32((_1362_v157).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1357_v155)
						_ = _rhs218
						var _rhs219 _dafny.Int = (_dafny.Zero).Minus(p2)
						_ = _rhs219
						var _rhs220 _dafny.Int = Companion_Default___.SafeModuloInt(p0, p0)
						_ = _rhs220
						var _rhs221 _dafny.Int = p0
						_ = _rhs221
						var _lhs180 *GlobalState = globalState
						_ = _lhs180
						var _lhs181 _dafny.Array = _1358_v156
						_ = _lhs181
						var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1358_v156), 0))
						_ = _lhs182
						var _lhs183 _dafny.Array = _1358_v156
						_ = _lhs183
						var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1358_v156), 0))
						_ = _lhs184
						_1357_v155 = _rhs218
						_lhs180.F14 = _rhs219
						(_lhs181).ArraySet1(_rhs220, (_lhs182).Int())
						(_lhs183).ArraySet1(_rhs221, (_lhs184).Int())
						var _1363_v158 _dafny.Array
						_ = _1363_v158
						var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
						_ = _nw218
						_1363_v158 = _nw218
						var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1363_v158), 0))
						_ = _index228
						(_1363_v158).ArraySet1(_this.F19(), (_index228).Int())
						var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1363_v158), 0))
						_ = _index229
						(_1363_v158).ArraySet1(_this.F19(), (_index229).Int())
						var _1364_v160 _dafny.Map
						_ = _1364_v160
						_1364_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Map {
							var _coll53 = _dafny.NewMapBuilder()
							_ = _coll53
							for _iter57 := _dafny.Iterate((_1132_v0).Keys().Elements()); ; {
								_compr_53, _ok57 := _iter57()
								if !_ok57 {
									break
								}
								var _1365_v159 _dafny.Int
								_1365_v159 = interface{}(_compr_53).(_dafny.Int)
								if (_1132_v0).Contains(_1365_v159) {
									_coll53.Add((_1365_v159).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg78 _dafny.Int) interface{} {
											return coer78(arg78)
										}
									}(func(_1366_i17 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('q')
									}))).Cardinality())), p1)
								}
							}
							return _coll53.ToMap()
						}()).Merge(_1132_v0))
						_1364_v160 = (_1364_v160).Update(((_this).F18()) || ((_this).F18()), (_1132_v0).Merge(_1132_v0))
					}
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		r0 = p1
		r1 = p1
		r2 = (_dafny.Zero).Minus(((_dafny.IntOfInt64(595)).Times(p0)).Minus(((_this).F20()).Plus(p2)))
		return r0, r1, r2
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f17 _dafny.Set
	_f19 _dafny.Array
	_f21 _dafny.Map
	_f20 _dafny.Int
	_f18 bool
	F23  _dafny.Int
	_f24 bool
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f21 = _dafny.EmptyMap
	_this._f20 = _dafny.Zero
	_this._f18 = false
	_this.F23 = _dafny.Zero
	_this._f24 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C8{}
var _ T2 = &C8{}
var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F17() _dafny.Set {
	return _this._f17
}
func (_this *C8) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C8) F19() _dafny.Array {
	return _this._f19
}
func (_this *C8) F19_set_(value _dafny.Array) {
	_this._f19 = value
}
func (_this *C8) F21() _dafny.Map {
	return _this._f21
}
func (_this *C8) F21_set_(value _dafny.Map) {
	_this._f21 = value
}
func (_this *C8) F20() _dafny.Int {
	return _this._f20
}
func (_this *C8) F18() bool {
	return _this._f18
}
func (_this *C8) Ctor__(f23 _dafny.Int, f24 bool, f19 _dafny.Array, f17 _dafny.Set, f18 bool, f20 _dafny.Int, f21 _dafny.Map) {
	{
		(_this).F23 = f23
		(_this)._f24 = f24
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C8) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fmocyx")).Cardinality()))).Cardinality()).Minus(_this.F23)).Times(_dafny.IntOfInt64(482))
	}
}
func (_this *C8) Fm1(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		return (_this).F18()
	}
}
func (_this *C8) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1367_i0 _dafny.Int
		_ = _1367_i0
		_1367_i0 = _dafny.Zero
		{
			for (_this).F24() {
				{
					if (_1367_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1367_i0 = (_1367_i0).Plus(_dafny.One)
					var _1368_v0 _dafny.Sequence
					_ = _1368_v0
					_1368_v0 = _dafny.UnicodeSeqOfUtf8Bytes("fgnisgrc")
					var _1369_v1 _dafny.Map
					_ = _1369_v1
					_1369_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Update((_this).F20(), (_this).F20()), _this.F23)
					_1368_v0 = (func() _dafny.Sequence {
						if (_this).Fm1(_1369_v1, globalState) {
							return _1368_v0
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("bkcjnysf")
					})()
					(globalState).F3 = !((_this).F18())
					var _1370_v2 _dafny.Array
					_ = _1370_v2
					var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
					_ = _nw219
					_1370_v2 = _nw219
					(globalState).F7 = _1370_v2
					var _1371_v3 _dafny.Array
					_ = _1371_v3
					var _nw220 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
					_ = _nw220
					_1371_v3 = _nw220
					_1371_v3 = (func() _dafny.Array {
						if (_this).F18() {
							return _1371_v3
						}
						return _1371_v3
					})()
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _1372_v4 _dafny.Map
		_ = _1372_v4
		_1372_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _this.F19())
		var _1373_v5 T0
		_ = _1373_v5
		var _nw221 *C2 = New_C2_()
		_ = _nw221
		_nw221.Ctor__((_this.F17()).Union(_dafny.SetOf(_this.F19(), (func() _dafny.Array {
			if (_1372_v4).Contains(p1) {
				return (_1372_v4).Get(p1).(_dafny.Array)
			}
			return _this.F19()
		})(), _this.F19())), !(p1))
		_1373_v5 = _nw221
		_1373_v5 = _1373_v5
		(globalState).F9 = Companion_Default___.Fm13(globalState)
		var _1374_v6 _dafny.Set
		_ = _1374_v6
		_1374_v6 = _dafny.SetOf((_1373_v5).F18(), (_1373_v5).F18())
		var _1375_v7 *C4
		_ = _1375_v7
		var _nw222 *C4 = New_C4_()
		_ = _nw222
		_nw222.Ctor__((_1374_v6).Cardinality(), (_this).F20(), _1373_v5.F17(), (_this).F24())
		_1375_v7 = _nw222
		r1 = (_this).F24()
		var _rhs222 bool = (_1373_v5).F18()
		_ = _rhs222
		var _rhs223 bool = (_1373_v5).F18()
		_ = _rhs223
		var _lhs185 *GlobalState = globalState
		_ = _lhs185
		var _lhs186 *GlobalState = globalState
		_ = _lhs186
		_lhs185.F9 = _rhs222
		_lhs186.F9 = _rhs223
		r0 = p1
		r1 = ((_1374_v6).Union(_dafny.SetOf(p1))).IsDisjointFrom((_1374_v6).Intersection(_1374_v6))
		var _1376_v8 _dafny.Array
		_ = _1376_v8
		var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
		_ = _nw223
		_1376_v8 = _nw223
		var _1377_v9 _dafny.MultiSet
		_ = _1377_v9
		_1377_v9 = _dafny.MultiSetOf(_1376_v8)
		r2 = (((_1377_v9).Intersection(_1377_v9)).Cardinality()).Minus((_this).F20())
		return r0, r1, r2
	}
}
func (_this *C8) M1(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		(globalState).F14 = Companion_Default___.SafeDivisionInt(p2, (_dafny.Zero).Minus(_dafny.IntOfInt64(-181)))
		for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_this.F19()), 0))); ; {
			_guard_loop_4, _ok58 := _iter58()
			if !_ok58 {
				break
			}
			var _1378_i0 _dafny.Int
			_1378_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1378_i0).Sign() != -1) && ((_1378_i0).Cmp(_dafny.ArrayLen((_this.F19()), 0)) < 0)) {
				var _arr43 _dafny.Array = _this.F19()
				_ = _arr43
				(_arr43).ArraySet1((_this).F24(), (_1378_i0).Int())
			}
		}
		var _1379_i1 _dafny.Int
		_ = _1379_i1
		_1379_i1 = _dafny.Zero
		{
			for !((_this).F24()) {
				{
					if (_1379_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1379_i1 = (_1379_i1).Plus(_dafny.One)
					var _1380_v0 _dafny.Sequence
					_ = _1380_v0
					_1380_v0 = _dafny.UnicodeSeqOfUtf8Bytes("t")
					_1380_v0 = _1380_v0
					var _1381_v1 T2
					_ = _1381_v1
					var _nw224 *C7 = New_C7_()
					_ = _nw224
					_nw224.Ctor__((_this).F20(), _this.F21(), _this.F19(), _this.F17(), !((_this).F18()))
					_1381_v1 = _nw224
					var _1382_v2 _dafny.Map
					_ = _1382_v2
					_1382_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1381_v1)
					var _1383_v3 _dafny.Sequence
					_ = _1383_v3
					_1383_v3 = _dafny.SeqOf(_1382_v2, _1382_v2, _1382_v2, _1382_v2, _1382_v2)
					var _1384_v4 _dafny.Map
					_ = _1384_v4
					_1384_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1383_v3).Cardinality()), _1381_v1.F19())
					var _1385_v5 D13
					_ = _1385_v5
					_1385_v5 = Companion_D13_.Create_DC31_(_1384_v4)
					var _1386_v6 _dafny.Array
					_ = _1386_v6
					var _nwElement0_45 _dafny.Map = _1384_v4
					_ = _nwElement0_45
					var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(10))
					_ = _nw225
					(_nw225).ArraySet1(_nwElement0_45, 0)
					(_nw225).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1381_v1.F19()), 1)
					(_nw225).ArraySet1(_1384_v4, 2)
					(_nw225).ArraySet1(_1384_v4, 3)
					(_nw225).ArraySet1((_1384_v4).Merge(_1384_v4), 4)
					(_nw225).ArraySet1((_1385_v5).Dtor_cf60(), 5)
					(_nw225).ArraySet1((_1384_v4).Merge(_1384_v4), 6)
					(_nw225).ArraySet1(_1384_v4, 7)
					(_nw225).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1381_v1).F20(), _1381_v1.F19()), 8)
					(_nw225).ArraySet1((_1384_v4).Merge(_1384_v4), 9)
					_1386_v6 = _nw225
					var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1386_v6), 0))
					_ = _index230
					(_1386_v6).ArraySet1(_1384_v4, (_index230).Int())
					var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1386_v6), 0))
					_ = _index231
					(_1386_v6).ArraySet1(_1384_v4, (_index231).Int())
					var _1387_v7 _dafny.Set
					_ = _1387_v7
					_1387_v7 = _dafny.SetOf((_this).F18())
					var _1388_v8 _dafny.Sequence
					_ = _1388_v8
					_1388_v8 = _dafny.SeqOf(_1387_v7, _1387_v7, _1387_v7, _1387_v7)
					var _1389_v9 _dafny.Map
					_ = _1389_v9
					_1389_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1381_v1).F18(), p0)
					var _1390_v10 _dafny.Map
					_ = _1390_v10
					_1390_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1381_v1).F20(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vxjbxgjo")).Cardinality()))
					var _1391_v11 _dafny.Sequence
					_ = _1391_v11
					_1391_v11 = _dafny.SeqOf((_this).F18(), p1)
					var _1392_v12 _dafny.MultiSet
					_ = _1392_v12
					_1392_v12 = _dafny.MultiSetOf(_1389_v9)
					var _1393_v13 _dafny.Array
					_ = _1393_v13
					var _nwElement0_46 _dafny.Int = p0
					_ = _nwElement0_46
					var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(15))
					_ = _nw226
					(_nw226).ArraySet1(_nwElement0_46, 0)
					(_nw226).ArraySet1((_1390_v10).Cardinality(), 1)
					(_nw226).ArraySet1(_dafny.IntOfUint32((_1391_v11).Cardinality()), 2)
					(_nw226).ArraySet1(_dafny.IntOfInt64(-285), 3)
					(_nw226).ArraySet1(p0, 4)
					(_nw226).ArraySet1((_1392_v12).Cardinality(), 5)
					(_nw226).ArraySet1((_dafny.Zero).Minus((_this).F20()), 6)
					(_nw226).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p1, (_this).F24(), (_this).F18())).Cardinality()), 7)
					(_nw226).ArraySet1((_this).F20(), 8)
					(_nw226).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bvsdc")).Cardinality()))), 9)
					(_nw226).ArraySet1(p0, 10)
					(_nw226).ArraySet1(p0, 11)
					(_nw226).ArraySet1(p2, 12)
					(_nw226).ArraySet1(_this.F23, 13)
					(_nw226).ArraySet1((_this).F20(), 14)
					_1393_v13 = _nw226
					var _1394_v14 D0
					_ = _1394_v14
					_1394_v14 = Companion_D0_.Create_DC2_(p1, p0, _1393_v13)
					var _1395_v16 _dafny.MultiSet
					_ = _1395_v16
					_1395_v16 = _dafny.MultiSetOf((_1381_v1).F20())
					var _1396_v17 _dafny.Array
					_ = _1396_v17
					var _nwElement0_47 _dafny.Int = Companion_Default___.Fm7(_dafny.IntOfInt64(916), _1389_v9, Companion_D0_.Create_DC0_((_1394_v14).Dtor_cf5()), globalState)
					_ = _nwElement0_47
					var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(8))
					_ = _nw227
					(_nw227).ArraySet1(_nwElement0_47, 0)
					(_nw227).ArraySet1((_this).F20(), 1)
					(_nw227).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(p1, p3)).Cardinality())).Minus((_dafny.Zero).Minus((_1389_v9).Cardinality())), 2)
					(_nw227).ArraySet1(p0, 3)
					(_nw227).ArraySet1(p2, 4)
					(_nw227).ArraySet1(((func() _dafny.Set {
						var _coll54 = _dafny.NewBuilder()
						_ = _coll54
						for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(60), _dafny.IntOfInt64(957))); ; {
							_compr_54, _ok59 := _iter59()
							if !_ok59 {
								break
							}
							var _1397_v15 _dafny.Int
							_1397_v15 = interface{}(_compr_54).(_dafny.Int)
							if ((_dafny.IntOfInt64(60)).Cmp(_1397_v15) <= 0) && ((_1397_v15).Cmp(_dafny.IntOfInt64(957)) < 0) {
								_coll54.Add((_1397_v15).Times((_dafny.SetOf(p2, (_this).F20())).Cardinality()))
							}
						}
						return _coll54.ToSet()
					}()).Difference(_dafny.SetOf(((_1395_v16).Update((_this).F20(), Companion_Default___.Abs(_this.F23))).Cardinality()))).Cardinality(), 5)
					(_nw227).ArraySet1((_this.F23).Plus((_this).F20()), 6)
					(_nw227).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p2), 7)
					_1396_v17 = _nw227
					var _rhs224 _dafny.Set = ((func() _dafny.Set {
						if (_this).F24() {
							return _1387_v7
						}
						return (_1388_v8).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_1388_v8).Cardinality()))).Uint32()).(_dafny.Set)
					})()).Union(_1387_v7)
					_ = _rhs224
					var _rhs225 bool = (_this).F18()
					_ = _rhs225
					var _rhs226 _dafny.Array = _1396_v17
					_ = _rhs226
					var _rhs227 bool = (_this).F18()
					_ = _rhs227
					var _lhs187 *GlobalState = globalState
					_ = _lhs187
					var _lhs188 *GlobalState = globalState
					_ = _lhs188
					var _lhs189 *GlobalState = globalState
					_ = _lhs189
					_1387_v7 = _rhs224
					_lhs187.F9 = _rhs225
					_lhs188.F7 = _rhs226
					_lhs189.F9 = _rhs227
					var _1398_v18 *C0
					_ = _1398_v18
					var _nw228 *C0 = New_C0_()
					_ = _nw228
					_nw228.Ctor__()
					_1398_v18 = _nw228
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _arr44 _dafny.Array = _this.F19()
		_ = _arr44
		var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index232
		(_arr44).ArraySet1((Companion_Default___.Fm29(globalState)).IsSubsetOf(_dafny.MultiSetOf(p3)), (_index232).Int())
		var _1399_v19 D5
		_ = _1399_v19
		_1399_v19 = Companion_D5_.Create_DC14_(!(false), _this.F23, p2)
		var _1400_v20 _dafny.Sequence
		_ = _1400_v20
		_1400_v20 = _dafny.SeqOf(p0, p0)
		var _1401_v21 _dafny.Map
		_ = _1401_v21
		_1401_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), ((_1399_v19).Dtor_cf27()).Plus(_dafny.IntOfUint32((_1400_v20).Cardinality())))
		var _1402_v22 _dafny.Sequence
		_ = _1402_v22
		_1402_v22 = _dafny.SeqOf(false)
		var _1403_v23 _dafny.Map
		_ = _1403_v23
		_1403_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
		var _arr45 _dafny.Array = _this.F19()
		_ = _arr45
		var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index233
		var _rhs228 bool = ((_this).F20()).Cmp((_dafny.IntOfUint32((_1402_v22).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dn")).Cardinality()))) < 0
		_ = _rhs228
		var _rhs229 bool = (func() bool {
			if p3 {
				return (p2).Cmp(_this.F23) >= 0
			}
			return (_this).F24()
		})()
		_ = _rhs229
		var _rhs230 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_1403_v23).Contains(_this.F23) {
				return (_1403_v23).Get(_this.F23).(bool)
			}
			return p3
		})(), _this.F23)).Merge(_1401_v21)
		_ = _rhs230
		var _lhs190 _dafny.Array = _this.F19()
		_ = _lhs190
		var _lhs191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))
		_ = _lhs191
		var _lhs192 *GlobalState = globalState
		_ = _lhs192
		(_lhs190).ArraySet1(_rhs228, (_lhs191).Int())
		_lhs192.F9 = _rhs229
		_1401_v21 = _rhs230
		var _1404_i2 _dafny.Int
		_ = _1404_i2
		_1404_i2 = _dafny.Zero
		{
			for (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool) {
				{
					if (_1404_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_1404_i2 = (_1404_i2).Plus(_dafny.One)
					if (_this).F18() {
						var _1405_v24 _dafny.CodePoint
						_ = _1405_v24
						_1405_v24 = _dafny.CodePoint('d')
						_1405_v24 = _1405_v24
						var _1406_v25 _dafny.Array
						_ = _1406_v25
						var _len0_42 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_42
						var _nw229 _dafny.Array
						_ = _nw229
						if _len0_42.Cmp(_dafny.Zero) == 0 {
							_nw229 = _dafny.NewArray(_len0_42)
						} else {
							var _init42 func(_dafny.Int) bool = func(_1407_i3 _dafny.Int) bool {
								return (_this).F18()
							}
							_ = _init42
							var _element0_42 = _init42(_dafny.Zero)
							_ = _element0_42
							_nw229 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
							(_nw229).ArraySet1(_element0_42, 0)
							var _nativeLen0_42 = (_len0_42).Int()
							_ = _nativeLen0_42
							for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
								(_nw229).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
							}
						}
						_1406_v25 = _nw229
						var _1408_v26 *C7
						_ = _1408_v26
						var _nw230 *C7 = New_C7_()
						_ = _nw230
						_nw230.Ctor__((_this).F20(), _this.F21(), _1406_v25, _this.F17(), (_this).F24())
						_1408_v26 = _nw230
						_1408_v26 = _1408_v26
						var _arr46 _dafny.Array = _this.F19()
						_ = _arr46
						var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))
						_ = _index234
						(_arr46).ArraySet1(p1, (_index234).Int())
						var _1409_v27 _dafny.Sequence
						_ = _1409_v27
						_1409_v27 = _dafny.UnicodeSeqOfUtf8Bytes("kme")
						var _1410_v28 T0
						_ = _1410_v28
						var _nw231 *C4 = New_C4_()
						_ = _nw231
						_nw231.Ctor__((_1403_v23).Cardinality(), p2, _this.F17(), p1)
						_1410_v28 = _nw231
						var _1411_v29 _dafny.Sequence
						_ = _1411_v29
						_1411_v29 = _dafny.SeqOf(_1410_v28, _1410_v28)
						var _1412_v30 _dafny.Array
						_ = _1412_v30
						var _nwElement0_48 _dafny.Int = (p2).Minus(p2)
						_ = _nwElement0_48
						var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(4))
						_ = _nw232
						(_nw232).ArraySet1(_nwElement0_48, 0)
						(_nw232).ArraySet1((p2).Minus(_dafny.IntOfUint32((_1409_v27).Cardinality())), 1)
						(_nw232).ArraySet1((_1400_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1411_v29).Cardinality()), _dafny.IntOfUint32((_1400_v20).Cardinality()))).Uint32()).(_dafny.Int), 2)
						(_nw232).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20()), 3)
						_1412_v30 = _nw232
						var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1412_v30), 0))
						_ = _index235
						(_1412_v30).ArraySet1(p2, (_index235).Int())
						var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1412_v30), 0))
						_ = _index236
						(_1412_v30).ArraySet1((_dafny.Zero).Minus(p0), (_index236).Int())
						_1403_v23 = (_1403_v23).Update((_1412_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1412_v30), 0))).Int()).(_dafny.Int), (func() bool {
							if (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool) {
								return p3
							}
							return p3
						})())
					} else {
						var _1413_v31 _dafny.CodePoint
						_ = _1413_v31
						_1413_v31 = _dafny.CodePoint('w')
						var _rhs231 _dafny.Int = ((_this).Fm0((_dafny.Zero).Minus(p0), _1413_v31, globalState)).Times(p0)
						_ = _rhs231
						var _rhs232 _dafny.Int = p0
						_ = _rhs232
						var _lhs193 *GlobalState = globalState
						_ = _lhs193
						var _lhs194 *GlobalState = globalState
						_ = _lhs194
						_lhs193.F14 = _rhs231
						_lhs194.F14 = _rhs232
						(globalState).F14 = ((_this).F20()).Plus(p2)
						var _1414_v32 _dafny.Array
						_ = _1414_v32
						var _len0_43 _dafny.Int = _dafny.IntOfInt64(16)
						_ = _len0_43
						var _nw233 _dafny.Array
						_ = _nw233
						if _len0_43.Cmp(_dafny.Zero) == 0 {
							_nw233 = _dafny.NewArray(_len0_43)
						} else {
							var _init43 func(_dafny.Int) _dafny.Int = (func(_1415_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1416_i4 _dafny.Int) _dafny.Int {
									return (_1416_i4).Minus(_1415_p2)
								}
							})(p2)
							_ = _init43
							var _element0_43 = _init43(_dafny.Zero)
							_ = _element0_43
							_nw233 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
							(_nw233).ArraySet1(_element0_43, 0)
							var _nativeLen0_43 = (_len0_43).Int()
							_ = _nativeLen0_43
							for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
								(_nw233).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
							}
						}
						_1414_v32 = _nw233
						var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_1414_v32), 0))
						_ = _index237
						(_1414_v32).ArraySet1(_this.F23, (_index237).Int())
						var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_1414_v32), 0))
						_ = _index238
						(_1414_v32).ArraySet1((((_this).F20()).Minus(p0)).Times(_dafny.IntOfInt64(586)), (_index238).Int())
						var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_1414_v32), 0))
						_ = _index239
						(_1414_v32).ArraySet1(p2, (_index239).Int())
						(globalState).F3 = !_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm16(globalState), true)
					}
					r0 = Companion_Default___.SafeModuloInt((_this).F20(), ((_dafny.Zero).Minus((_1403_v23).Cardinality())).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))).Cardinality()))
					var _1417_v33 _dafny.MultiSet
					_ = _1417_v33
					_1417_v33 = _dafny.MultiSetOf(p1)
					var _1418_v34 _dafny.Sequence
					_ = _1418_v34
					_1418_v34 = _dafny.UnicodeSeqOfUtf8Bytes("pinsphqf")
					var _1419_v35 _dafny.Array
					_ = _1419_v35
					var _nwElement0_49 bool = _dafny.Companion_Sequence_.Contains(_1418_v34, _dafny.CodePoint('g'))
					_ = _nwElement0_49
					var _nw234 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(4))
					_ = _nw234
					(_nw234).ArraySet1(_nwElement0_49, 0)
					(_nw234).ArraySet1(!((_this).F24()), 1)
					(_nw234).ArraySet1(Companion_Default___.Fm13(globalState), 2)
					(_nw234).ArraySet1(p1, 3)
					_1419_v35 = _nw234
					var _1420_v36 _dafny.Map
					_ = _1420_v36
					_1420_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
					var _rhs233 _dafny.MultiSet = Companion_Default___.Fm29(globalState)
					_ = _rhs233
					var _rhs234 _dafny.Int = (_this).F20()
					_ = _rhs234
					var _rhs235 bool = !((func() bool {
						if (_1420_v36).Contains(p1) {
							return (_1420_v36).Get(p1).(bool)
						}
						return (_this).F18()
					})())
					_ = _rhs235
					var _rhs236 _dafny.Array = _1419_v35
					_ = _rhs236
					var _lhs195 *GlobalState = globalState
					_ = _lhs195
					_1417_v33 = _rhs233
					r0 = _rhs234
					_lhs195.F3 = _rhs235
					_1419_v35 = _rhs236
					var _1421_v37 *C4
					_ = _1421_v37
					var _nw235 *C4 = New_C4_()
					_ = _nw235
					_nw235.Ctor__(p2, ((_1401_v21).Update(p1, p2)).Cardinality(), _this.F17(), p3)
					_1421_v37 = _nw235
					var _1422_v38 _dafny.Map
					_ = _1422_v38
					_1422_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1421_v37)
					_1422_v38 = (_1422_v38).Update((p0).Times((_this).F20()), _1421_v37)
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _1423_v39 _dafny.MultiSet
		_ = _1423_v39
		_1423_v39 = _dafny.MultiSetOf((_1401_v21).Cardinality())
		var _1424_v40 _dafny.CodePoint
		_ = _1424_v40
		_1424_v40 = _dafny.CodePoint('i')
		var _1425_v41 _dafny.Sequence
		_ = _1425_v41
		_1425_v41 = _dafny.SeqOf(_1424_v40)
		var _1426_v42 D3
		_ = _1426_v42
		_1426_v42 = Companion_D3_.Create_DC9_()
		var _1427_v44 _dafny.Set
		_ = _1427_v44
		_1427_v44 = _dafny.SetOf(_1403_v23, func() _dafny.Map {
			var _coll55 = _dafny.NewMapBuilder()
			_ = _coll55
			for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(662), _dafny.IntOfInt64(960))); ; {
				_compr_55, _ok60 := _iter60()
				if !_ok60 {
					break
				}
				var _1428_v43 _dafny.Int
				_1428_v43 = interface{}(_compr_55).(_dafny.Int)
				if ((_dafny.IntOfInt64(662)).Cmp(_1428_v43) <= 0) && ((_1428_v43).Cmp(_dafny.IntOfInt64(960)) < 0) {
					_coll55.Add((_1428_v43).Minus(p0), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
				}
			}
			return _coll55.ToMap()
		}())
		var _1429_v45 D0
		_ = _1429_v45
		_1429_v45 = Companion_D0_.Create_DC0_(true)
		var _1430_v46 D1
		_ = _1430_v46
		_1430_v46 = Companion_D1_.Create_DC4_(!(true), _1427_v44, Companion_Default___.Fm7((_this).F20(), (_1401_v21).Update((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), p2), _1429_v45, globalState))
		var _1431_v47 _dafny.Map
		_ = _1431_v47
		_1431_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1402_v22).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_1402_v22).Cardinality()))).Uint32()).(bool))
		var _1432_v48 _dafny.Map
		_ = _1432_v48
		_1432_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1424_v40)
		var _1433_v49 _dafny.Set
		_ = _1433_v49
		_1433_v49 = _dafny.SetOf(p3)
		var _1434_v50 _dafny.MultiSet
		_ = _1434_v50
		_1434_v50 = _dafny.MultiSetOf(_1424_v40, (func() _dafny.CodePoint {
			if (_1432_v48).Contains((_1433_v49).Cardinality()) {
				return (_1432_v48).Get((_1433_v49).Cardinality()).(_dafny.CodePoint)
			}
			return _1424_v40
		})())
		var _1435_v51 _dafny.Map
		_ = _1435_v51
		_1435_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1402_v22)
		var _1436_v52 _dafny.Set
		_ = _1436_v52
		_1436_v52 = _dafny.SetOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1435_v51).Contains(_dafny.IntOfInt64(123)) {
				return (_1435_v51).Get(_dafny.IntOfInt64(123)).(_dafny.Sequence)
			}
			return _1402_v22
		})()).Cardinality()))
		var _pat_let_tv17 = p1
		_ = _pat_let_tv17
		var _1437_v53 _dafny.Array
		_ = _1437_v53
		var _nwElement0_50 _dafny.Int = ((_1423_v39).Intersection(_1423_v39)).Cardinality()
		_ = _nwElement0_50
		var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(24))
		_ = _nw236
		(_nw236).ArraySet1(_nwElement0_50, 0)
		(_nw236).ArraySet1(_dafny.IntOfUint32((_1402_v22).Cardinality()), 1)
		(_nw236).ArraySet1(_dafny.IntOfUint32((_1425_v41).Cardinality()), 2)
		(_nw236).ArraySet1(_this.F23, 3)
		(_nw236).ArraySet1(p0, 4)
		(_nw236).ArraySet1(p0, 5)
		(_nw236).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm34(_1426_v42, (_this).F18(), _1424_v40, p1, globalState)).Cardinality()), 6)
		(_nw236).ArraySet1(((_dafny.Zero).Minus(p2)).Times((func(_pat_let24_0 D1) D1 {
			return func(_1438_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let25_0 bool) D1 {
					return func(_1439_dt__update_hcf9_h0 bool) D1 {
						return Companion_D1_.Create_DC4_(_1439_dt__update_hcf9_h0, (_1438_dt__update__tmp_h0).Dtor_cf10(), (_1438_dt__update__tmp_h0).Dtor_cf11())
					}(_pat_let25_0)
				}(_pat_let_tv17)
			}(_pat_let24_0)
		}(_1430_v46)).Dtor_cf11()), 7)
		(_nw236).ArraySet1((_dafny.Zero).Minus(_this.F23), 8)
		(_nw236).ArraySet1(p2, 9)
		(_nw236).ArraySet1((_1431_v47).Cardinality(), 10)
		(_nw236).ArraySet1((_1434_v50).Cardinality(), 11)
		(_nw236).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), !((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)))).Cardinality(), 12)
		(_nw236).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1424_v40)).Merge(_1432_v48)).Cardinality(), 13)
		(_nw236).ArraySet1(_this.F23, 14)
		(_nw236).ArraySet1((_this).F20(), 15)
		(_nw236).ArraySet1(p2, 16)
		(_nw236).ArraySet1((_this).F20(), 17)
		(_nw236).ArraySet1((_this).F20(), 18)
		(_nw236).ArraySet1(_this.F23, 19)
		(_nw236).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if !((_this).F18()) {
				return _1400_v20
			}
			return _1400_v20
		})()).Cardinality()), 20)
		(_nw236).ArraySet1((_1436_v52).Cardinality(), 21)
		(_nw236).ArraySet1((p0).Times((_this).F20()), 22)
		(_nw236).ArraySet1(p0, 23)
		_1437_v53 = _nw236
		for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1437_v53), 0))); ; {
			_guard_loop_5, _ok61 := _iter61()
			if !_ok61 {
				break
			}
			var _1440_i5 _dafny.Int
			_1440_i5 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1440_i5).Sign() != -1) && ((_1440_i5).Cmp(_dafny.ArrayLen((_1437_v53), 0)) < 0)) {
				(_1437_v53).ArraySet1((_1440_i5).Times(p2), (_1440_i5).Int())
			}
		}
		r0 = p2
		return r0
	}
}
func (_this *C8) M2(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1441_v0 _dafny.Set
		_ = _1441_v0
		_1441_v0 = _dafny.SetOf(_this.F23)
		var _1442_v1 _dafny.Map
		_ = _1442_v1
		_1442_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, (_this).F24())
		var _1443_v2 T1
		_ = _1443_v2
		var _nw237 *C3 = New_C3_()
		_ = _nw237
		_nw237.Ctor__(_this.F19(), _dafny.SetOf(_this.F19(), _this.F19()), !((func() bool {
			if (_1442_v1).Contains(p0) {
				return (_1442_v1).Get(p0).(bool)
			}
			return false
		})()))
		_1443_v2 = _nw237
		var _1444_v3 _dafny.MultiSet
		_ = _1444_v3
		_1444_v3 = _dafny.MultiSetOf(_1443_v2, _1443_v2)
		var _1445_v4 _dafny.Sequence
		_ = _1445_v4
		_1445_v4 = _dafny.SeqOf((_1443_v2).F18(), Companion_Default___.Fm13(globalState))
		var _1446_v5 _dafny.Sequence
		_ = _1446_v5
		_1446_v5 = _dafny.UnicodeSeqOfUtf8Bytes("bwffq")
		var _1447_v6 _dafny.Map
		_ = _1447_v6
		_1447_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _dafny.IntOfUint32((_1446_v5).Cardinality()))
		var _1448_v7 _dafny.Map
		_ = _1448_v7
		_1448_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1444_v3).Cardinality(), (_1445_v4).Select((Companion_Default___.SafeIndex((_1447_v6).Cardinality(), _dafny.IntOfUint32((_1445_v4).Cardinality()))).Uint32()).(bool))
		var _arr47 _dafny.Array = _this.F19()
		_ = _arr47
		var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index240
		(_arr47).ArraySet1(!(_1441_v0).Equals(_dafny.SetOf((_dafny.Zero).Minus((_1442_v1).Cardinality()))), (_index240).Int())
		var _arr48 _dafny.Array = _this.F19()
		_ = _arr48
		var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index241
		(_arr48).ArraySet1((_1443_v2).F18(), (_index241).Int())
		var _1449_i0 _dafny.Int
		_ = _1449_i0
		_1449_i0 = _dafny.Zero
		{
			for true {
				{
					if (_1449_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L15
					}
					_1449_i0 = (_1449_i0).Plus(_dafny.One)
					var _len0_44 _dafny.Int = _dafny.IntOfInt64(3)
					_ = _len0_44
					var _nw238 _dafny.Array
					_ = _nw238
					if _len0_44.Cmp(_dafny.Zero) == 0 {
						_nw238 = _dafny.NewArray(_len0_44)
					} else {
						var _init44 func(_dafny.Int) _dafny.CodePoint = func(_1450_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('r')
						}
						_ = _init44
						var _element0_44 = _init44(_dafny.Zero)
						_ = _element0_44
						_nw238 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
						(_nw238).ArraySet1CodePoint(_element0_44, 0)
						var _nativeLen0_44 = (_len0_44).Int()
						_ = _nativeLen0_44
						for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
							(_nw238).ArraySet1CodePoint(_init44(_dafny.IntOf(_i0_44)), _i0_44)
						}
					}
					(globalState).F16 = _nw238
					if (_1443_v2).F18() {
						var _1451_v8 _dafny.Array
						_ = _1451_v8
						var _nw239 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
						_ = _nw239
						_1451_v8 = _nw239
						var _1452_v9 _dafny.CodePoint
						_ = _1452_v9
						_1452_v9 = _dafny.CodePoint('y')
						var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1451_v8), 0))
						_ = _index242
						(_1451_v8).ArraySet1CodePoint(_1452_v9, (_index242).Int())
						var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1451_v8), 0))
						_ = _index243
						(_1451_v8).ArraySet1CodePoint(_1452_v9, (_index243).Int())
						var _1453_v10 *C3
						_ = _1453_v10
						var _nw240 *C3 = New_C3_()
						_ = _nw240
						_nw240.Ctor__(_1443_v2.F19(), _this.F17(), true)
						_1453_v10 = _nw240
						var _1454_v11 _dafny.Map
						_ = _1454_v11
						_1454_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _1453_v10)
						var _1455_v12 D3
						_ = _1455_v12
						_1455_v12 = Companion_D3_.Create_DC9_()
						var _1456_v13 _dafny.Sequence
						_ = _1456_v13
						_1456_v13 = _dafny.SeqOf(_1455_v12, _1455_v12, _1455_v12)
						var _1457_v14 D5
						_ = _1457_v14
						_1457_v14 = Companion_D5_.Create_DC13_(!((_1443_v2).F18()), (_1454_v11).Cardinality(), _1456_v13)
						_1457_v14 = _1457_v14
						var _1458_v15 *C2
						_ = _1458_v15
						var _nw241 *C2 = New_C2_()
						_ = _nw241
						_nw241.Ctor__(_1443_v2.F17(), true)
						_1458_v15 = _nw241
						var _1459_v16 _dafny.Map
						_ = _1459_v16
						_1459_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F23).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oed")).Cardinality())), _1458_v15)
						_1459_v16 = _1459_v16
						var _1460_v17 _dafny.Array
						_ = _1460_v17
						var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
						_ = _nw242
						_1460_v17 = _nw242
						var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_1460_v17), 0))
						_ = _index244
						(_1460_v17).ArraySet1(_this.F23, (_index244).Int())
						var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_1460_v17), 0))
						_ = _index245
						(_1460_v17).ArraySet1((_dafny.Zero).Minus(p0), (_index245).Int())
						_1460_v17 = _1460_v17
					} else {
						(globalState).F14 = _this.F23
						var _1461_v18 *C2
						_ = _1461_v18
						var _nw243 *C2 = New_C2_()
						_ = _nw243
						_nw243.Ctor__((_this.F17()).Intersection(_1443_v2.F17()), ((_this).F20()).Cmp((_this).F20()) == 0)
						_1461_v18 = _nw243
						var _1462_v19 _dafny.Map
						_ = _1462_v19
						_1462_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_this).F20())
						var _1463_v20 _dafny.Sequence
						_ = _1463_v20
						_1463_v20 = _dafny.SeqOf((_1462_v19).Merge(_1462_v19))
						var _1464_v21 _dafny.Map
						_ = _1464_v21
						_1464_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v6, (_this).F20())
						var _rhs237 bool = (_this).Fm1(_1464_v21, globalState)
						_ = _rhs237
						var _rhs238 _dafny.Sequence = _1463_v20
						_ = _rhs238
						var _lhs196 *GlobalState = globalState
						_ = _lhs196
						_lhs196.F9 = _rhs237
						_1463_v20 = _rhs238
						(_1443_v2).F19_set_(_1443_v2.F19())
						var _1465_v22 _dafny.MultiSet
						_ = _1465_v22
						_1465_v22 = _dafny.MultiSetOf((_this).F24())
						var _1466_v23 _dafny.Map
						_ = _1466_v23
						_1466_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F18())
						var _1467_v24 D3
						_ = _1467_v24
						_1467_v24 = Companion_D3_.Create_DC10_(_1446_v5, (_1443_v2).F18(), _dafny.UnicodeSeqOfUtf8Bytes("te"), _1446_v5, (_1466_v23).Cardinality())
						var _1468_v25 D11
						_ = _1468_v25
						_1468_v25 = Companion_D11_.Create_DC26_((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), _dafny.IntOfUint32((_1446_v5).Cardinality()))
						var _1469_v26 _dafny.Array
						_ = _1469_v26
						var _nwElement0_51 _dafny.MultiSet = _dafny.MultiSetOf((_1443_v2).F18())
						_ = _nwElement0_51
						var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(9))
						_ = _nw244
						(_nw244).ArraySet1(_nwElement0_51, 0)
						(_nw244).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Union(_1465_v22), 1)
						(_nw244).ArraySet1(_dafny.MultiSetOf((_1467_v24).Dtor_cf17(), true, (_1443_v2).F18(), (_1443_v2).F18()), 2)
						(_nw244).ArraySet1(_1465_v22, 3)
						(_nw244).ArraySet1(_1465_v22, 4)
						(_nw244).ArraySet1((_dafny.MultiSetFromSeq(_1445_v4)).Update((_1443_v2).F18(), Companion_Default___.Abs(_this.F23)), 5)
						(_nw244).ArraySet1((Companion_Default___.Fm29(globalState)).Intersection(_1465_v22), 6)
						(_nw244).ArraySet1((_1465_v22).Difference(_1465_v22), 7)
						(_nw244).ArraySet1(_dafny.MultiSetOf((_1468_v25).Dtor_cf48()), 8)
						_1469_v26 = _nw244
						var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1469_v26), 0))
						_ = _index246
						(_1469_v26).ArraySet1(_1465_v22, (_index246).Int())
						var _1470_v27 _dafny.Sequence
						_ = _1470_v27
						_1470_v27 = _dafny.SeqOf(Companion_Default___.Fm29(globalState), Companion_Default___.Fm29(globalState))
						var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1469_v26), 0))
						_ = _index247
						(_1469_v26).ArraySet1((_1470_v27).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_1470_v27).Cardinality()))).Uint32()).(_dafny.MultiSet), (_index247).Int())
					}
					_1442_v1 = (_1442_v1).Update((_this).F20(), (_1443_v2).F18())
					var _1471_v28 _dafny.CodePoint
					_ = _1471_v28
					_1471_v28 = _dafny.CodePoint('h')
					var _1472_v29 _dafny.Map
					_ = _1472_v29
					_1472_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.SetOf((_1447_v6).Cardinality(), _this.F23))
					var _1473_v30 _dafny.Set
					_ = _1473_v30
					_1473_v30 = _dafny.SetOf((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
					var _1474_v31 _dafny.Sequence
					_ = _1474_v31
					_1474_v31 = _dafny.SeqOf((_this).Fm0((_1447_v6).Cardinality(), _1471_v28, globalState), Companion_Default___.SafeDivisionInt((_this).F20(), (_this).F20()), Companion_Default___.SafeModuloInt((_this).F20(), ((func() _dafny.Set {
						if (_1472_v29).Contains(p0) {
							return (_1472_v29).Get(p0).(_dafny.Set)
						}
						return _dafny.SetOf((_1473_v30).Cardinality())
					})()).Cardinality()))
					var _1475_v32 _dafny.Map
					_ = _1475_v32
					_1475_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), p0)
					var _1476_v33 D0
					_ = _1476_v33
					_1476_v33 = Companion_D0_.Create_DC0_(Companion_Default___.Fm13(globalState))
					var _1477_v34 D3
					_ = _1477_v34
					_1477_v34 = Companion_D3_.Create_DC9_()
					var _1478_v35 _dafny.Sequence
					_ = _1478_v35
					_1478_v35 = _dafny.SeqOf(_1477_v34, _1477_v34)
					var _1479_v36 _dafny.Sequence
					_ = _1479_v36
					_1479_v36 = _1478_v35
					var _1480_v38 _dafny.Array
					_ = _1480_v38
					var _nwElement0_52 _dafny.Int = _this.F23
					_ = _nwElement0_52
					var _nw245 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(16))
					_ = _nw245
					(_nw245).ArraySet1(_nwElement0_52, 0)
					(_nw245).ArraySet1(Companion_Default___.Fm7((_this).F20(), _1475_v32, _1476_v33, globalState), 1)
					(_nw245).ArraySet1((_this).F20(), 2)
					(_nw245).ArraySet1((p0).Minus(_dafny.IntOfInt64(696)), 3)
					(_nw245).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F20()), (_this).F20()), 4)
					(_nw245).ArraySet1(Companion_Default___.SafeModuloInt(_this.F23, _dafny.IntOfInt64(440)), 5)
					(_nw245).ArraySet1((_dafny.Zero).Minus(p0), 6)
					(_nw245).ArraySet1((_1473_v30).Cardinality(), 7)
					(_nw245).ArraySet1((Companion_D5_.Create_DC13_((_this).F24(), p0, _1479_v36)).Dtor_cf24(), 8)
					(_nw245).ArraySet1(p0, 9)
					(_nw245).ArraySet1(p0, 10)
					(_nw245).ArraySet1((func() _dafny.Int {
						if (_1475_v32).Contains((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)) {
							return (_1475_v32).Get((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)).(_dafny.Int)
						}
						return p0
					})(), 11)
					(_nw245).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Map {
						var _coll56 = _dafny.NewMapBuilder()
						_ = _coll56
						for _iter62 := _dafny.Iterate((_1442_v1).Keys().Elements()); ; {
							_compr_56, _ok62 := _iter62()
							if !_ok62 {
								break
							}
							var _1481_v37 _dafny.Int
							_1481_v37 = interface{}(_compr_56).(_dafny.Int)
							if (_1442_v1).Contains(_1481_v37) {
								_coll56.Add((_1481_v37).Minus((_dafny.Zero).Minus(p0)), _1471_v28)
							}
						}
						return _coll56.ToMap()
					}()).Cardinality(), _dafny.IntOfUint32((_1445_v4).Cardinality())), 12)
					(_nw245).ArraySet1(Companion_Default___.SafeModuloInt(p0, _this.F23), 13)
					(_nw245).ArraySet1((_this).F20(), 14)
					(_nw245).ArraySet1((func() _dafny.Int {
						if (_1475_v32).Contains((_1443_v2).F18()) {
							return (_1475_v32).Get((_1443_v2).F18()).(_dafny.Int)
						}
						return (_dafny.Zero).Minus(_this.F23)
					})(), 15)
					_1480_v38 = _nw245
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_1480_v38), 0))
					_ = _index248
					(_1480_v38).ArraySet1((_this.F23).Plus(_this.F23), (_index248).Int())
					var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_1480_v38), 0))
					_ = _index249
					var _rhs239 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1474_v31, _1474_v31)
					_ = _rhs239
					var _rhs240 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F20()), (_this).F20())
					_ = _rhs240
					var _rhs241 _dafny.Int = p0
					_ = _rhs241
					var _lhs197 _dafny.Array = _1480_v38
					_ = _lhs197
					var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_1480_v38), 0))
					_ = _lhs198
					var _lhs199 *GlobalState = globalState
					_ = _lhs199
					_1474_v31 = _rhs239
					(_lhs197).ArraySet1(_rhs240, (_lhs198).Int())
					_lhs199.F14 = _rhs241
					goto C15
				}
			C15:
			}
			goto L15
		}
	L15:
		var _1482_v39 D11
		_ = _1482_v39
		_1482_v39 = Companion_D11_.Create_DC27_(_dafny.Companion_Sequence_.Update(_1445_v4, (Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_1445_v4).Cardinality()))).Uint32(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)), p0, (_this).F20())
		var _1483_v40 _dafny.Sequence
		_ = _1483_v40
		_1483_v40 = _dafny.SeqOf(_1448_v7)
		var _1484_v41 _dafny.MultiSet
		_ = _1484_v41
		_1484_v41 = _dafny.MultiSetOf(p0)
		var _1485_v42 _dafny.CodePoint
		_ = _1485_v42
		_1485_v42 = _dafny.CodePoint('b')
		var _1486_v43 _dafny.Set
		_ = _1486_v43
		_1486_v43 = _dafny.SetOf(_1442_v1)
		var _1487_v44 _dafny.Map
		_ = _1487_v44
		_1487_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), false)
		var _1488_v45 D1
		_ = _1488_v45
		_1488_v45 = Companion_D1_.Create_DC4_((_this).F18(), _1486_v43, (_1487_v44).Cardinality())
		var _pat_let_tv18 = _1486_v43
		_ = _pat_let_tv18
		var _1489_v46 _dafny.Array
		_ = _1489_v46
		var _nwElement0_53 _dafny.Int = (p0).Plus((_dafny.Zero).Minus(p0))
		_ = _nwElement0_53
		var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(16))
		_ = _nw246
		(_nw246).ArraySet1(_nwElement0_53, 0)
		(_nw246).ArraySet1(p0, 1)
		(_nw246).ArraySet1(_this.F23, 2)
		(_nw246).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if (_1447_v6).Contains(_this.F23) {
				return (_1447_v6).Get(_this.F23).(_dafny.Int)
			}
			return (_this).F20()
		})()), 3)
		(_nw246).ArraySet1((_1482_v39).Dtor_cf52(), 4)
		(_nw246).ArraySet1((_this).F20(), 5)
		(_nw246).ArraySet1((_this).F20(), 6)
		(_nw246).ArraySet1((p0).Plus((_this).F20()), 7)
		(_nw246).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(588), _dafny.IntOfInt64(-132)), 8)
		(_nw246).ArraySet1(_dafny.IntOfInt64(323), 9)
		(_nw246).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1483_v40).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1483_v40).Cardinality()))).Uint32()).(_dafny.Map), (_this).F20())).Cardinality(), 10)
		(_nw246).ArraySet1(Companion_Default___.SafeModuloInt((_1484_v41).Cardinality(), (_this).F20()), 11)
		(_nw246).ArraySet1(p0, 12)
		(_nw246).ArraySet1(_this.F23, 13)
		(_nw246).ArraySet1((_this).Fm0(_this.F23, _1485_v42, globalState), 14)
		(_nw246).ArraySet1((func() _dafny.Int {
			if (func(_pat_let26_0 D1) D1 {
				return func(_1490_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let27_0 _dafny.Set) D1 {
						return func(_1491_dt__update_hcf10_h0 _dafny.Set) D1 {
							return Companion_D1_.Create_DC4_((_1490_dt__update__tmp_h0).Dtor_cf9(), _1491_dt__update_hcf10_h0, (_1490_dt__update__tmp_h0).Dtor_cf11())
						}(_pat_let27_0)
					}(_pat_let_tv18)
				}(_pat_let26_0)
			}(_1488_v45)).Dtor_cf9() {
				return (_this).F20()
			}
			return (_1441_v0).Cardinality()
		})(), 15)
		_1489_v46 = _nw246
		var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_1489_v46), 0))
		_ = _index250
		(_1489_v46).ArraySet1(_dafny.IntOfInt64(358), (_index250).Int())
		var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_1489_v46), 0))
		_ = _index251
		(_1489_v46).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1446_v5, _1446_v5)).Cardinality())), (_index251).Int())
		var _rhs242 bool = (_this).F24()
		_ = _rhs242
		var _rhs243 _dafny.Sequence = _1445_v4
		_ = _rhs243
		var _lhs200 *GlobalState = globalState
		_ = _lhs200
		_lhs200.F3 = _rhs242
		_1445_v4 = _rhs243
		var _1492_v47 _dafny.Map
		_ = _1492_v47
		_1492_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1485_v42, (_1489_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_1489_v46), 0))).Int()).(_dafny.Int))
		var _1493_v48 _dafny.Map
		_ = _1493_v48
		_1493_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1492_v47).Merge(_1492_v47), _1487_v44)
		var _1494_v50 _dafny.Set
		_ = _1494_v50
		_1494_v50 = _dafny.SetOf(_1485_v42)
		_1493_v48 = (_1493_v48).Update(func() _dafny.Map {
			var _coll57 = _dafny.NewMapBuilder()
			_ = _coll57
			for _iter63 := _dafny.Iterate((_1494_v50).Elements()); ; {
				_compr_57, _ok63 := _iter63()
				if !_ok63 {
					break
				}
				var _1495_v49 _dafny.CodePoint
				_1495_v49 = interface{}(_compr_57).(_dafny.CodePoint)
				if (_1494_v50).Contains(_1495_v49) {
					_coll57.Add(_1495_v49, (Companion_D5_.Create_DC14_((_1443_v2).F18(), _dafny.IntOfUint32((_1446_v5).Cardinality()), p0)).Dtor_cf28())
				}
			}
			return _coll57.ToMap()
		}(), _1487_v44)
		var _1496_v51 _dafny.Array
		_ = _1496_v51
		var _nw247 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw247
		_1496_v51 = _nw247
		var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1496_v51), 0))
		_ = _index252
		(_1496_v51).ArraySet1(_1489_v46, (_index252).Int())
		var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1496_v51), 0))
		_ = _index253
		(_1496_v51).ArraySet1(_1489_v46, (_index253).Int())
		r0 = _this.F23
		r1 = !((_this).F18())
		r2 = Companion_Default___.Fm4(((_1489_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_1489_v46), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(649))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg79 _dafny.Int) interface{} {
				return coer79(arg79)
			}
		}((func(_1497_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1498_i2 _dafny.Int) _dafny.CodePoint {
				return _1497_v42
			}
		})(_1485_v42)))).Cardinality())), globalState)
		return r0, r1, r2
	}
}
func (_this *C8) F24() bool {
	{
		return _this._f24
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f17 _dafny.Set
	_f19 _dafny.Array
	_f21 _dafny.Map
	_f20 _dafny.Int
	_f18 bool
	F22  _dafny.Int
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f17 = _dafny.EmptySet
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f21 = _dafny.EmptyMap
	_this._f20 = _dafny.Zero
	_this._f18 = false
	_this.F22 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C9{}
var _ T1 = &C9{}
var _ T0 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) F17() _dafny.Set {
	return _this._f17
}
func (_this *C9) F17_set_(value _dafny.Set) {
	_this._f17 = value
}
func (_this *C9) F19() _dafny.Array {
	return _this._f19
}
func (_this *C9) F19_set_(value _dafny.Array) {
	_this._f19 = value
}
func (_this *C9) F21() _dafny.Map {
	return _this._f21
}
func (_this *C9) F21_set_(value _dafny.Map) {
	_this._f21 = value
}
func (_this *C9) F20() _dafny.Int {
	return _this._f20
}
func (_this *C9) F18() bool {
	return _this._f18
}
func (_this *C9) Ctor__(f22 _dafny.Int, f20 _dafny.Int, f21 _dafny.Map, f19 _dafny.Array, f17 _dafny.Set, f18 bool) {
	{
		(_this).F22 = f22
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f19 = f19
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *C9) Fm1(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		return (_this).F18()
	}
}
func (_this *C9) Fm0(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F18())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())))).Cardinality())
	}
}
func (_this *C9) Fm2(p0 bool, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(770)).Minus((_this).F20())
	}
}
func (_this *C9) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1499_v0 _dafny.CodePoint
		_ = _1499_v0
		_1499_v0 = _dafny.CodePoint('t')
		var _1500_v1 _dafny.Set
		_ = _1500_v1
		_1500_v1 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("x"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("x")).Cardinality()))).Uint32(), _1499_v0))
		_1500_v1 = _1500_v1
		var _1501_v2 _dafny.Sequence
		_ = _1501_v2
		_1501_v2 = _dafny.SeqOf(!((_this).F18()))
		var _1502_v3 _dafny.Sequence
		_ = _1502_v3
		_1502_v3 = _dafny.SeqOf(_1501_v2, _1501_v2)
		var _1503_v4 _dafny.Array
		_ = _1503_v4
		var _nwElement0_54 _dafny.Sequence = _1502_v3
		_ = _nwElement0_54
		var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(6))
		_ = _nw248
		(_nw248).ArraySet1(_nwElement0_54, 0)
		(_nw248).ArraySet1(_1502_v3, 1)
		(_nw248).ArraySet1(_1502_v3, 2)
		(_nw248).ArraySet1(_1502_v3, 3)
		(_nw248).ArraySet1(_1502_v3, 4)
		(_nw248).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1502_v3, _1502_v3), 5)
		_1503_v4 = _nw248
		var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1503_v4), 0))
		_ = _index254
		(_1503_v4).ArraySet1(_1502_v3, (_index254).Int())
		var _arr49 _dafny.Array = _this.F19()
		_ = _arr49
		var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index255
		(_arr49).ArraySet1(p1, (_index255).Int())
		var _1504_v5 _dafny.MultiSet
		_ = _1504_v5
		_1504_v5 = _dafny.MultiSetOf(!(p1))
		var _1505_v7 _dafny.Map
		_ = _1505_v7
		_1505_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, (_this).F18())
		var _1506_v8 _dafny.Sequence
		_ = _1506_v8
		_1506_v8 = _dafny.SeqOf(p2, _this.F22, p0)
		var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1503_v4), 0))
		_ = _index256
		var _arr50 _dafny.Array = _this.F19()
		_ = _arr50
		var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
		_ = _index257
		var _rhs244 bool = !((func() bool {
			if (_this).F18() {
				return p1
			}
			return !((_this).F18())
		})())
		_ = _rhs244
		var _rhs245 bool = !(((func() _dafny.MultiSet {
			if (_this).F18() {
				return _dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F18()))
			}
			return _1504_v5
		})()).IsDisjointFrom(_1504_v5))
		_ = _rhs245
		var _rhs246 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-140))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg80 _dafny.Int) interface{} {
				return coer80(arg80)
			}
		}((func(_1507_p1 bool) func(_dafny.Int) _dafny.Sequence {
			return func(_1508_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_1507_p1, (_this).F18(), _1507_p1, !(_1507_p1), (_this).F18())
			}
		})(p1)))
		_ = _rhs246
		var _rhs247 bool = (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll58 = _dafny.NewMapBuilder()
			_ = _coll58
			for _iter64 := _dafny.Iterate((_1505_v7).Keys().Elements()); ; {
				_compr_58, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _1509_v6 _dafny.Int
				_1509_v6 = interface{}(_compr_58).(_dafny.Int)
				if (_1505_v7).Contains(_1509_v6) {
					_coll58.Add((_1509_v6).Times((_dafny.Zero).Minus(_this.F22)), _this.F22)
				}
			}
			return _coll58.ToMap()
		}(), _dafny.IntOfUint32((_1506_v8).Cardinality())), globalState)
		_ = _rhs247
		var _lhs201 _dafny.Array = _1503_v4
		_ = _lhs201
		var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1503_v4), 0))
		_ = _lhs202
		var _lhs203 _dafny.Array = _this.F19()
		_ = _lhs203
		var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
		_ = _lhs204
		r0 = _rhs244
		r0 = _rhs245
		(_lhs201).ArraySet1(_rhs246, (_lhs202).Int())
		(_lhs203).ArraySet1(_rhs247, (_lhs204).Int())
		var _1510_i1 _dafny.Int
		_ = _1510_i1
		_1510_i1 = _dafny.Zero
		{
			for (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool) {
				{
					if (_1510_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L16
					}
					_1510_i1 = (_1510_i1).Plus(_dafny.One)
					var _1511_v9 _dafny.Array
					_ = _1511_v9
					var _len0_45 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_45
					var _nw249 _dafny.Array
					_ = _nw249
					if _len0_45.Cmp(_dafny.Zero) == 0 {
						_nw249 = _dafny.NewArray(_len0_45)
					} else {
						var _init45 func(_dafny.Int) _dafny.CodePoint = (func(_1512_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1513_i2 _dafny.Int) _dafny.CodePoint {
								return _1512_v0
							}
						})(_1499_v0)
						_ = _init45
						var _element0_45 = _init45(_dafny.Zero)
						_ = _element0_45
						_nw249 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
						(_nw249).ArraySet1CodePoint(_element0_45, 0)
						var _nativeLen0_45 = (_len0_45).Int()
						_ = _nativeLen0_45
						for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
							(_nw249).ArraySet1CodePoint(_init45(_dafny.IntOf(_i0_45)), _i0_45)
						}
					}
					_1511_v9 = _nw249
					var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))
					_ = _index258
					(_1511_v9).ArraySet1CodePoint(_1499_v0, (_index258).Int())
					var _1514_v10 _dafny.Sequence
					_ = _1514_v10
					_1514_v10 = _dafny.UnicodeSeqOfUtf8Bytes("bmfhaq")
					var _1515_v11 _dafny.MultiSet
					_ = _1515_v11
					_1515_v11 = _dafny.MultiSetOf(_1499_v0)
					var _1516_v12 _dafny.Map
					_ = _1516_v12
					_1516_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1515_v11).Cardinality(), p2)
					var _1517_v13 _dafny.Map
					_ = _1517_v13
					_1517_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1516_v12, p0)
					var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))
					_ = _index259
					var _rhs248 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1514_v10, _1514_v10)).Cardinality())
					_ = _rhs248
					var _rhs249 bool = (_this).Fm1(_1517_v13, globalState)
					_ = _rhs249
					var _rhs250 _dafny.CodePoint = _1499_v0
					_ = _rhs250
					var _rhs251 _dafny.Int = _dafny.IntOfInt64(795)
					_ = _rhs251
					var _lhs205 *C9 = _this
					_ = _lhs205
					var _lhs206 _dafny.Array = _1511_v9
					_ = _lhs206
					var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))
					_ = _lhs207
					_lhs205.F22 = _rhs248
					r0 = _rhs249
					(_lhs206).ArraySet1CodePoint(_rhs250, (_lhs207).Int())
					r2 = _rhs251
					var _1518_v14 _dafny.Map
					_ = _1518_v14
					_1518_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), true)
					if (_1518_v14).Contains(!((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))) {
						(globalState).F3 = p1
						var _arr51 _dafny.Array = _this.F19()
						_ = _arr51
						var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
						_ = _index260
						var _rhs252 _dafny.Int = _this.F22
						_ = _rhs252
						var _rhs253 bool = !_dafny.Companion_Sequence_.Contains(_1514_v10, (_1511_v9).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))).Int()))
						_ = _rhs253
						var _rhs254 bool = (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
						_ = _rhs254
						var _rhs255 _dafny.Int = (_this).Fm0((_1506_v8).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1506_v8).Cardinality()))).Uint32()).(_dafny.Int), _1499_v0, globalState)
						_ = _rhs255
						var _rhs256 bool = (_1516_v12).Contains((_this).F20())
						_ = _rhs256
						var _lhs208 *GlobalState = globalState
						_ = _lhs208
						var _lhs209 *GlobalState = globalState
						_ = _lhs209
						var _lhs210 _dafny.Array = _this.F19()
						_ = _lhs210
						var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
						_ = _lhs211
						r2 = _rhs252
						_lhs208.F3 = _rhs253
						r1 = _rhs254
						_lhs209.F14 = _rhs255
						(_lhs210).ArraySet1(_rhs256, (_lhs211).Int())
						_1514_v10 = _dafny.UnicodeSeqOfUtf8Bytes("jdrmxxf")
						r0 = (!(p1)) || ((func() bool {
							if p1 {
								return (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
							}
							return (_1501_v2).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32()).(bool)
						})())
						var _arr52 _dafny.Array = _this.F19()
						_ = _arr52
						var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
						_ = _index261
						(_arr52).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_1514_v10, _1514_v10), _1499_v0), (_index261).Int())
					} else {
						var _1519_v15 _dafny.Array
						_ = _1519_v15
						var _len0_46 _dafny.Int = _dafny.IntOfInt64(15)
						_ = _len0_46
						var _nw250 _dafny.Array
						_ = _nw250
						if _len0_46.Cmp(_dafny.Zero) == 0 {
							_nw250 = _dafny.NewArray(_len0_46)
						} else {
							var _init46 func(_dafny.Int) bool = func(_1520_i3 _dafny.Int) bool {
								return (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
							}
							_ = _init46
							var _element0_46 = _init46(_dafny.Zero)
							_ = _element0_46
							_nw250 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
							(_nw250).ArraySet1(_element0_46, 0)
							var _nativeLen0_46 = (_len0_46).Int()
							_ = _nativeLen0_46
							for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
								(_nw250).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
							}
						}
						_1519_v15 = _nw250
						var _1521_v16 T2
						_ = _1521_v16
						var _nw251 *C6 = New_C6_()
						_ = _nw251
						_nw251.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_this).F20()), (_this).F20(), (_this).F20(), _this.F21(), _1519_v15, _this.F17(), (_this).F18())
						_1521_v16 = _nw251
						var _1522_v17 _dafny.Set
						_ = _1522_v17
						_1522_v17 = _dafny.SetOf(_1521_v16, _1521_v16)
						(globalState).F14 = (func() _dafny.Int {
							if (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool) {
								return _dafny.IntOfUint32((_1514_v10).Cardinality())
							}
							return ((_this).F20()).Times((_1522_v17).Cardinality())
						})()
						(globalState).F14 = _this.F22
						var _1523_v18 *C3
						_ = _1523_v18
						var _nw252 *C3 = New_C3_()
						_ = _nw252
						_nw252.Ctor__(_1519_v15, (_dafny.SetOf(_1519_v15)).Union(_this.F17()), !(_1505_v7).Contains((_1521_v16).F20()))
						_1523_v18 = _nw252
						var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))
						_ = _index262
						(_1511_v9).ArraySet1CodePoint((_1511_v9).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))).Int()), (_index262).Int())
						var _arr53 _dafny.Array = _this.F19()
						_ = _arr53
						var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))
						_ = _index263
						(_arr53).ArraySet1(!((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)), (_index263).Int())
					}
					if p1 {
						var _1524_v19 _dafny.Array
						_ = _1524_v19
						var _nw253 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
						_ = _nw253
						_1524_v19 = _nw253
						var _1525_v20 _dafny.Map
						_ = _1525_v20
						_1525_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1524_v19, _1524_v19)
						_1525_v20 = ((_1525_v20).Merge(_1525_v20)).Merge(_1525_v20)
						var _1526_v21 _dafny.Map
						_ = _1526_v21
						_1526_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm21(_dafny.SeqOf((_this).F20(), _dafny.IntOfInt64(231)), globalState), p2)
						_1526_v21 = (_1526_v21).Update(_1505_v7, p0)
						var _1527_v22 *C2
						_ = _1527_v22
						var _nw254 *C2 = New_C2_()
						_ = _nw254
						_nw254.Ctor__(_this.F17(), Companion_Default___.Fm13(globalState))
						_1527_v22 = _nw254
						var _1528_v23 D11
						_ = _1528_v23
						_1528_v23 = Companion_D11_.Create_DC27_(_dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), (_this).F18()), p0, p2)
						var _1529_v24 D0
						_ = _1529_v24
						_1529_v24 = Companion_D0_.Create_DC1_(_dafny.Companion_Sequence_.Concatenate(_1514_v10, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("hcpfqafh"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hcpfqafh")).Cardinality()))).Uint32(), (_1511_v9).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1511_v9), 0))).Int()))), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_1528_v23).Dtor_cf51()), p0), p0, _1499_v0)
						_1529_v24 = _1529_v24
						(globalState).F9 = p1
					} else {
						var _1530_v25 _dafny.Array
						_ = _1530_v25
						var _nwElement0_55 bool = (_this).F18()
						_ = _nwElement0_55
						var _nw255 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(11))
						_ = _nw255
						(_nw255).ArraySet1(_nwElement0_55, 0)
						(_nw255).ArraySet1((_this).F18(), 1)
						(_nw255).ArraySet1((_this).F18(), 2)
						(_nw255).ArraySet1(Companion_Default___.Fm13(globalState), 3)
						(_nw255).ArraySet1(!((_this).Fm1(_1517_v13, globalState)), 4)
						(_nw255).ArraySet1((_this).F18(), 5)
						(_nw255).ArraySet1(false, 6)
						(_nw255).ArraySet1((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), 7)
						(_nw255).ArraySet1((_this).F18(), 8)
						(_nw255).ArraySet1(p1, 9)
						(_nw255).ArraySet1((_this).F18(), 10)
						_1530_v25 = _nw255
						var _1531_v27 *C8
						_ = _1531_v27
						var _nw256 *C8 = New_C8_()
						_ = _nw256
						_nw256.Ctor__(p0, !((p1) == ((func() bool {
							if (_1505_v7).Contains(_this.F22) {
								return (_1505_v7).Get(_this.F22).(bool)
							}
							return (_this).F18()
						})())), _1530_v25, _this.F17(), (_this).Fm1(_1517_v13, globalState), p2, (func() _dafny.Map {
							var _coll59 = _dafny.NewMapBuilder()
							_ = _coll59
							for _iter65 := _dafny.Iterate((_1515_v11).Elements()); ; {
								_compr_59, _ok65 := _iter65()
								if !_ok65 {
									break
								}
								var _1532_v26 _dafny.CodePoint
								_1532_v26 = interface{}(_compr_59).(_dafny.CodePoint)
								if (_1515_v11).Contains(_1532_v26) {
									_coll59.Add(_1532_v26, p1)
								}
							}
							return _coll59.ToMap()
						}()).Merge(_this.F21()))
						_1531_v27 = _nw256
						_1531_v27 = _1531_v27
						_1530_v25 = _1530_v25
						var _1533_v28 *C0
						_ = _1533_v28
						var _nw257 *C0 = New_C0_()
						_ = _nw257
						_nw257.Ctor__()
						_1533_v28 = _nw257
						_1533_v28 = _1533_v28
						var _1534_v29 _dafny.Set
						_ = _1534_v29
						_1534_v29 = _dafny.SetOf((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
						var _1535_v30 _dafny.Sequence
						_ = _1535_v30
						_1535_v30 = _1501_v2
						var _1536_v31 _dafny.MultiSet
						_ = _1536_v31
						_1536_v31 = _dafny.MultiSetOf((_this).F20())
						var _1537_v32 _dafny.Array
						_ = _1537_v32
						var _nwElement0_56 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1501_v2, _dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), false))
						_ = _nwElement0_56
						var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(17))
						_ = _nw258
						(_nw258).ArraySet1(_nwElement0_56, 0)
						(_nw258).ArraySet1((func() _dafny.Sequence {
							if (_1531_v27).F24() {
								return _1501_v2
							}
							return _1501_v2
						})(), 1)
						(_nw258).ArraySet1(_1501_v2, 2)
						(_nw258).ArraySet1(_dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), (_1531_v27).F24()), 3)
						(_nw258).ArraySet1((_1535_v30), 4)
						(_nw258).ArraySet1(_1501_v2, 5)
						(_nw258).ArraySet1(_1501_v2, 6)
						(_nw258).ArraySet1(_dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), false), 7)
						(_nw258).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1501_v2, Companion_Default___.Fm16(globalState)), 8)
						(_nw258).ArraySet1(_1501_v2, 9)
						(_nw258).ArraySet1(Companion_Default___.Fm16(globalState), 10)
						(_nw258).ArraySet1(_1501_v2, 11)
						(_nw258).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F18(), (_this).F18(), (_this).F18(), p1, p1), _1501_v2), 12)
						(_nw258).ArraySet1(_1501_v2, 13)
						(_nw258).ArraySet1(Companion_Default___.Fm16(globalState), 14)
						(_nw258).ArraySet1(_1501_v2, 15)
						(_nw258).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex((_1536_v31).Cardinality(), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), (_this).F18()), _1501_v2), 16)
						_1537_v32 = _nw258
						var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1537_v32), 0))
						_ = _index264
						(_1537_v32).ArraySet1(_1501_v2, (_index264).Int())
						var _1538_v33 _dafny.Sequence
						_ = _1538_v33
						_1538_v33 = _dafny.SeqOf(Companion_Default___.Fm37((_this).F18(), _this.F22, (_this).F18(), (_1531_v27).F24(), globalState))
						var _1539_v34 _dafny.MultiSet
						_ = _1539_v34
						_1539_v34 = _dafny.MultiSetOf(((_1538_v33).Select((Companion_Default___.SafeIndex((_1518_v14).Cardinality(), _dafny.IntOfUint32((_1538_v33).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_1534_v29))
						var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1537_v32), 0))
						_ = _index265
						var _rhs257 _dafny.Set = _dafny.SetOf((_dafny.IntOfUint32((_1514_v10).Cardinality())).Cmp(p2) == 0)
						_ = _rhs257
						var _rhs258 _dafny.Int = (_dafny.IntOfUint32((_1506_v8).Cardinality())).Times(_dafny.IntOfUint32((Companion_Default___.Fm40(p1, (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), globalState)).Cardinality()))
						_ = _rhs258
						var _rhs259 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1501_v2, _1501_v2)
						_ = _rhs259
						var _rhs260 _dafny.MultiSet = _1539_v34
						_ = _rhs260
						var _lhs212 *C9 = _this
						_ = _lhs212
						var _lhs213 _dafny.Array = _1537_v32
						_ = _lhs213
						var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1537_v32), 0))
						_ = _lhs214
						_1534_v29 = _rhs257
						_lhs212.F22 = _rhs258
						(_lhs213).ArraySet1(_rhs259, (_lhs214).Int())
						_1539_v34 = _rhs260
						var _1540_v35 _dafny.Map
						_ = _1540_v35
						_1540_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v25, (_this).F20())
						var _rhs261 _dafny.Map = _1540_v35
						_ = _rhs261
						var _rhs262 _dafny.Int = _this.F22
						_ = _rhs262
						var _rhs263 bool = true
						_ = _rhs263
						var _rhs264 _dafny.Int = (_this).F20()
						_ = _rhs264
						var _rhs265 _dafny.Int = (_dafny.Zero).Minus((_this).F20())
						_ = _rhs265
						var _lhs215 *GlobalState = globalState
						_ = _lhs215
						var _lhs216 *GlobalState = globalState
						_ = _lhs216
						var _lhs217 *GlobalState = globalState
						_ = _lhs217
						_1540_v35 = _rhs261
						_lhs215.F14 = _rhs262
						_lhs216.F9 = _rhs263
						_lhs217.F14 = _rhs264
						r2 = _rhs265
					}
					r1 = !((_this).Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1516_v12, p2)).Merge(_1517_v13), globalState))
					goto C16
				}
			C16:
			}
			goto L16
		}
	L16:
		var _1541_v36 _dafny.Sequence
		_ = _1541_v36
		_1541_v36 = _dafny.SeqOf(_1499_v0)
		var _1542_v37 _dafny.Array
		_ = _1542_v37
		var _nw259 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw259
		_1542_v37 = _nw259
		var _1543_v38 _dafny.Map
		_ = _1543_v38
		_1543_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v0, _1542_v37)
		var _1544_v39 _dafny.Array
		_ = _1544_v39
		var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
		_ = _nw260
		_1544_v39 = _nw260
		var _1545_v40 _dafny.Map
		_ = _1545_v40
		_1545_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _1546_v41 _dafny.Set
		_ = _1546_v41
		_1546_v41 = _dafny.SetOf(_1505_v7, _1505_v7)
		var _1547_v42 D1
		_ = _1547_v42
		_1547_v42 = Companion_D1_.Create_DC4_((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), _1546_v41, _dafny.IntOfInt64(-445))
		var _1548_v43 _dafny.Sequence
		_ = _1548_v43
		_1548_v43 = _dafny.SeqOf(_1547_v42)
		var _1549_v44 _dafny.MultiSet
		_ = _1549_v44
		_1549_v44 = _dafny.MultiSetOf((_this).F20(), (_this).F20(), (func() _dafny.Int {
			if (_1545_v40).Contains(_this.F22) {
				return (_1545_v40).Get(_this.F22).(_dafny.Int)
			}
			return p2
		})(), _dafny.IntOfUint32((_1548_v43).Cardinality()), p2)
		var _1550_v45 _dafny.Set
		_ = _1550_v45
		_1550_v45 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D14_.Create_DC34_(_1544_v39)).Dtor_cf64(), _1499_v0)).Cardinality(), (func() _dafny.Int {
			if (_1549_v44).Contains(_dafny.IntOfInt64(346)) {
				return (_1549_v44).Multiplicity(_dafny.IntOfInt64(346))
			}
			return (_this).F20()
		})())
		var _1551_v46 _dafny.Map
		_ = _1551_v46
		_1551_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v0, (_1550_v45).Cardinality())
		var _1552_v47 _dafny.Sequence
		_ = _1552_v47
		_1552_v47 = _dafny.SeqOf(_1551_v46)
		var _1553_v48 _dafny.Map
		_ = _1553_v48
		_1553_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1552_v47, _dafny.UnicodeSeqOfUtf8Bytes("ymwc"))
		var _rhs266 _dafny.Sequence = _1541_v36
		_ = _rhs266
		var _rhs267 _dafny.Map = _1543_v38
		_ = _rhs267
		var _rhs268 _dafny.Int = p0
		_ = _rhs268
		var _rhs269 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1541_v36, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1541_v36).Cardinality()))).Uint32(), _1499_v0), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1553_v48).Contains(_1552_v47) {
				return (_1553_v48).Get(_1552_v47).(_dafny.Sequence)
			}
			return _1541_v36
		})(), _1541_v36)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1501_v2, _dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1541_v36, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1541_v36).Cardinality()))).Uint32(), _1499_v0), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1553_v48).Contains(_1552_v47) {
				return (_1553_v48).Get(_1552_v47).(_dafny.Sequence)
			}
			return _1541_v36
		})(), _1541_v36))).Cardinality()))).Uint32(), _1499_v0), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1541_v36).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1541_v36, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1541_v36).Cardinality()))).Uint32(), _1499_v0), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1553_v48).Contains(_1552_v47) {
				return (_1553_v48).Get(_1552_v47).(_dafny.Sequence)
			}
			return _1541_v36
		})(), _1541_v36)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1501_v2, _dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1541_v36, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1541_v36).Cardinality()))).Uint32(), _1499_v0), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1553_v48).Contains(_1552_v47) {
				return (_1553_v48).Get(_1552_v47).(_dafny.Sequence)
			}
			return _1541_v36
		})(), _1541_v36))).Cardinality()))).Uint32(), _1499_v0)).Cardinality()))).Uint32(), _1499_v0)
		_ = _rhs269
		var _rhs270 _dafny.Int = p0
		_ = _rhs270
		var _lhs218 *GlobalState = globalState
		_ = _lhs218
		_1541_v36 = _rhs266
		_1543_v38 = _rhs267
		_lhs218.F14 = _rhs268
		_1541_v36 = _rhs269
		r2 = _rhs270
		var _1554_i4 _dafny.Int
		_ = _1554_i4
		_1554_i4 = _dafny.Zero
		{
			for !(p1) || ((_this).F18()) {
				{
					if (_1554_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L17
					}
					_1554_i4 = (_1554_i4).Plus(_dafny.One)
					(_this).F22 = p2
					_1551_v46 = (_1551_v46).Update(_1499_v0, (_this).F20())
					var _1555_v49 _dafny.Map
					_ = _1555_v49
					_1555_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1501_v2, _1499_v0)
					var _1556_v50 _dafny.Array
					_ = _1556_v50
					var _nwElement0_57 _dafny.Map = _1555_v49
					_ = _nwElement0_57
					var _nw261 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(3))
					_ = _nw261
					(_nw261).ArraySet1(_nwElement0_57, 0)
					(_nw261).ArraySet1(_1555_v49, 1)
					(_nw261).ArraySet1(_1555_v49, 2)
					_1556_v50 = _nw261
					var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1556_v50), 0))
					_ = _index266
					(_1556_v50).ArraySet1(_1555_v49, (_index266).Int())
					var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1556_v50), 0))
					_ = _index267
					var _rhs271 bool = (_this).F18()
					_ = _rhs271
					var _rhs272 _dafny.Set = (_1550_v45).Difference(_1550_v45)
					_ = _rhs272
					var _rhs273 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1501_v2, _1499_v0)
					_ = _rhs273
					var _rhs274 bool = (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
					_ = _rhs274
					var _lhs219 *GlobalState = globalState
					_ = _lhs219
					var _lhs220 _dafny.Array = _1556_v50
					_ = _lhs220
					var _lhs221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1556_v50), 0))
					_ = _lhs221
					var _lhs222 *GlobalState = globalState
					_ = _lhs222
					_lhs219.F9 = _rhs271
					_1550_v45 = _rhs272
					(_lhs220).ArraySet1(_rhs273, (_lhs221).Int())
					_lhs222.F3 = _rhs274
					(globalState).F3 = (p1) && (p1)
					goto C17
				}
			C17:
			}
			goto L17
		}
	L17:
		var _1557_i5 _dafny.Int
		_ = _1557_i5
		_1557_i5 = _dafny.Zero
		{
			for (_this).F18() {
				{
					if (_1557_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L18
					}
					_1557_i5 = (_1557_i5).Plus(_dafny.One)
					r0 = !((_this).F18())
					if false {
						var _1558_v51 _dafny.Map
						_ = _1558_v51
						_1558_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)) || ((_this).F18()), !((_1547_v42).Dtor_cf9()))
						var _1559_v52 T0
						_ = _1559_v52
						var _nw262 *C2 = New_C2_()
						_ = _nw262
						_nw262.Ctor__(_this.F17(), p1)
						_1559_v52 = _nw262
						var _1560_v53 _dafny.MultiSet
						_ = _1560_v53
						_1560_v53 = _dafny.MultiSetOf(_1559_v52)
						_1558_v51 = (_1558_v51).Update((_1547_v42).Dtor_cf9(), (_dafny.MultiSetOf(_1559_v52)).IsDisjointFrom(_1560_v53))
						_1499_v0 = _1499_v0
						_1501_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1501_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1541_v36, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_1541_v36).Cardinality()))).Uint32(), _1499_v0)).Cardinality()), _dafny.IntOfUint32((_1501_v2).Cardinality()))).Uint32(), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)), (Companion_Default___.SafeIndex(_this.F22, _dafny.IntOfUint32((_dafny.SeqOf((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))).Cardinality()))).Uint32(), (_1559_v52).F18()))
						var _1561_v54 *C4
						_ = _1561_v54
						var _nw263 *C4 = New_C4_()
						_ = _nw263
						_nw263.Ctor__(((_1505_v7).Merge(_1505_v7)).Cardinality(), p0, _this.F17(), (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1545_v40, p0), globalState))
						_1561_v54 = _nw263
						(_1561_v54).F26 = _dafny.IntOfUint32(((func() _dafny.Sequence {
							if ((_this).Fm2((_1559_v52).F18(), _dafny.IntOfUint32((_1541_v36).Cardinality()), Companion_Default___.Fm43(globalState), globalState)).Cmp(_dafny.IntOfUint32((_1541_v36).Cardinality())) >= 0 {
								return _dafny.Companion_Sequence_.Concatenate(_1501_v2, Companion_Default___.Fm16(globalState))
							}
							return _1501_v2
						})()).Cardinality())
					} else {
						var _1562_v55 D0
						_ = _1562_v55
						_1562_v55 = Companion_D0_.Create_DC0_((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
						_1499_v0 = Companion_Default___.Fm33(Companion_Default___.Fm7(p2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F22), func(_pat_let28_0 D0) D0 {
							return func(_1563_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let29_0 bool) D0 {
									return func(_1564_dt__update_hcf0_h0 bool) D0 {
										return Companion_D0_.Create_DC0_(_1564_dt__update_hcf0_h0)
									}(_pat_let29_0)
								}((_this).F18())
							}(_pat_let28_0)
						}(_1562_v55), globalState), (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), globalState)
						var _1565_v56 _dafny.Array
						_ = _1565_v56
						var _nw264 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
						_ = _nw264
						_1565_v56 = _nw264
						var _1566_v57 *C7
						_ = _1566_v57
						var _nw265 *C7 = New_C7_()
						_ = _nw265
						_nw265.Ctor__(_this.F22, _this.F21(), _1565_v56, _this.F17(), p1)
						_1566_v57 = _nw265
						var _1567_v58 _dafny.Map
						_ = _1567_v58
						_1567_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1566_v57, (_this).F20())
						_1567_v58 = (_1567_v58).Update(_1566_v57, _dafny.IntOfUint32((_1541_v36).Cardinality()))
						(globalState).F3 = p1
						var _1568_v59 _dafny.Array
						_ = _1568_v59
						var _nw266 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
						_ = _nw266
						_1568_v59 = _nw266
						var _1569_v60 _dafny.Map
						_ = _1569_v60
						_1569_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v0, _1565_v56)
						var _1570_v61 *C5
						_ = _1570_v61
						var _nw267 *C5 = New_C5_()
						_ = _nw267
						_nw267.Ctor__(_1541_v36, (_this).F18(), p0, _this.F21(), (func() _dafny.Array {
							if (_1569_v60).Contains(_1499_v0) {
								return (_1569_v60).Get(_1499_v0).(_dafny.Array)
							}
							return _1565_v56
						})(), _dafny.SetOf(_1565_v56), false)
						_1570_v61 = _nw267
						var _1571_v62 _dafny.Map
						_ = _1571_v62
						_1571_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool))
						var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1568_v59), 0))
						_ = _index268
						(_1568_v59).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1570_v61, _1571_v62)).Update(_1570_v61, _1571_v62), (_index268).Int())
						var _1572_v63 _dafny.Map
						_ = _1572_v63
						_1572_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v5, _1570_v61)
						var _1573_v64 _dafny.Map
						_ = _1573_v64
						_1573_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C5 {
							if (_1572_v63).Contains(_1504_v5) {
								return (_1572_v63).Get(_1504_v5).(*C5)
							}
							return _1570_v61
						})(), _1571_v62)
						var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1568_v59), 0))
						_ = _index269
						(_1568_v59).ArraySet1(_1573_v64, (_index269).Int())
						var _1574_v67 _dafny.Map
						_ = _1574_v67
						_1574_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
							var _coll60 = _dafny.NewMapBuilder()
							_ = _coll60
							for _iter66 := _dafny.Iterate((_1506_v8).Elements()); ; {
								_compr_60, _ok66 := _iter66()
								if !_ok66 {
									break
								}
								var _1575_v66 _dafny.Int
								_1575_v66 = interface{}(_compr_60).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_1506_v8, _1575_v66) {
									_coll60.Add((_1575_v66).Plus(p2), (_dafny.Zero).Minus((_this).F20()))
								}
							}
							return _coll60.ToMap()
						}(), (_this).F20())
						(globalState).F3 = (_this).Fm1(func() _dafny.Map {
							var _coll61 = _dafny.NewMapBuilder()
							_ = _coll61
							for _iter67 := _dafny.Iterate((_1574_v67).Keys().Elements()); ; {
								_compr_61, _ok67 := _iter67()
								if !_ok67 {
									break
								}
								var _1576_v65 _dafny.Map
								_1576_v65 = interface{}(_compr_61).(_dafny.Map)
								if (_1574_v67).Contains(_1576_v65) {
									_coll61.Add(_1576_v65, p0)
								}
							}
							return _coll61.ToMap()
						}(), globalState)
					}
					var _1577_v68 _dafny.Array
					_ = _1577_v68
					var _len0_47 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_47
					var _nw268 _dafny.Array
					_ = _nw268
					if _len0_47.Cmp(_dafny.Zero) == 0 {
						_nw268 = _dafny.NewArray(_len0_47)
					} else {
						var _init47 func(_dafny.Int) bool = func(_1578_i6 _dafny.Int) bool {
							return (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool)
						}
						_ = _init47
						var _element0_47 = _init47(_dafny.Zero)
						_ = _element0_47
						_nw268 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
						(_nw268).ArraySet1(_element0_47, 0)
						var _nativeLen0_47 = (_len0_47).Int()
						_ = _nativeLen0_47
						for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
							(_nw268).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
						}
					}
					_1577_v68 = _nw268
					var _1579_v70 *C8
					_ = _1579_v70
					var _nw269 *C8 = New_C8_()
					_ = _nw269
					_nw269.Ctor__(p0, (_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), _1577_v68, _this.F17(), (_this).F18(), p2, func() _dafny.Map {
						var _coll62 = _dafny.NewMapBuilder()
						_ = _coll62
						for _iter68 := _dafny.Iterate((_1551_v46).Keys().Elements()); ; {
							_compr_62, _ok68 := _iter68()
							if !_ok68 {
								break
							}
							var _1580_v69 _dafny.CodePoint
							_1580_v69 = interface{}(_compr_62).(_dafny.CodePoint)
							if (_1551_v46).Contains(_1580_v69) {
								_coll62.Add(_1580_v69, p1)
							}
						}
						return _coll62.ToMap()
					}())
					_1579_v70 = _nw269
					var _1581_v71 _dafny.Array
					_ = _1581_v71
					var _nwElement0_58 *C8 = _1579_v70
					_ = _nwElement0_58
					var _nw270 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(10))
					_ = _nw270
					(_nw270).ArraySet1(_nwElement0_58, 0)
					(_nw270).ArraySet1(_1579_v70, 1)
					(_nw270).ArraySet1(_1579_v70, 2)
					(_nw270).ArraySet1(_1579_v70, 3)
					(_nw270).ArraySet1(_1579_v70, 4)
					(_nw270).ArraySet1(_1579_v70, 5)
					(_nw270).ArraySet1(_1579_v70, 6)
					(_nw270).ArraySet1(_1579_v70, 7)
					(_nw270).ArraySet1(_1579_v70, 8)
					(_nw270).ArraySet1(_1579_v70, 9)
					_1581_v71 = _nw270
					var _1582_v72 _dafny.Set
					_ = _1582_v72
					_1582_v72 = _dafny.SetOf(_1581_v71)
					_1582_v72 = _1582_v72
					(_this).F22 = p2
					goto C18
				}
			C18:
			}
			goto L18
		}
	L18:
		var _1583_v73 _dafny.Set
		_ = _1583_v73
		_1583_v73 = _dafny.SetOf((_this.F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_this.F19()), 0))).Int()).(bool), (_this).F18())
		r0 = (_dafny.IntOfUint32((_1541_v36).Cardinality())).Cmp((_1583_v73).Cardinality()) > 0
		r1 = !(((func() _dafny.MultiSet {
			if !((_this).F18()) {
				return _dafny.MultiSetOf(p0)
			}
			return _1549_v44
		})()).IsDisjointFrom(_1549_v44))
		r2 = (p2).Times(p0)
		return r0, r1, r2
	}
}

// End of class C9
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
