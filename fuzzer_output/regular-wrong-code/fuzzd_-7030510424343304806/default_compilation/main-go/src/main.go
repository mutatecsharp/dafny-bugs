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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfUint32((_dafny.SeqOf(!(!(false)))).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-917))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality())) < 0 {
		return (_dafny.CodePoint('h'))
	} else {
		return _dafny.CodePoint('c')
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Set, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) _dafny.Int {
	return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), true)).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(490)))).Cardinality())).Cardinality()))).Cardinality()))).Times(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("wtycevgkj"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (_dafny.UnicodeSeqOfUtf8Bytes("lqmpa"))))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(229), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(838))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(294))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(720))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("auo")).Cardinality())
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-51))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i3 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(318)
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true, !(!(!(true))), true)).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(497))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_5_i4 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(245)
	}))))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("p"), _dafny.UnicodeSeqOfUtf8Bytes("rjwgaoud"))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D7 = Companion_D7_.Create_DC14_(Companion_D7_.Create_DC14_(Companion_D7_.Create_DC14_(Companion_D7_.Create_DC14_(Companion_D7_.Create_DC12_(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(989), (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(46), _dafny.IntOfInt64(483))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _6_v0 _dafny.Int
				_6_v0 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(46)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(483)) < 0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_6_v0, (func() _dafny.Map {
						var _coll2 = _dafny.NewMapBuilder()
						_ = _coll2
						for _iter2 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-312), (_dafny.MultiSetOf(_dafny.IntOfInt64(220))).Cardinality())).Elements()); ; {
							_compr_2, _ok2 := _iter2()
							if !_ok2 {
								break
							}
							var _7_v1 _dafny.Int
							_7_v1 = interface{}(_compr_2).(_dafny.Int)
							if (_dafny.MultiSetOf(_dafny.IntOfInt64(-312), (_dafny.MultiSetOf(_dafny.IntOfInt64(220))).Cardinality())).Contains(_7_v1) {
								_coll2.Add(Companion_Default___.SafeDivisionInt(_7_v1, _dafny.IntOfInt64(-651)), _dafny.IntOfInt64(555))
							}
						}
						return _coll2.ToMap()
					}()).Cardinality()), _dafny.IntOfInt64(899))
				}
			}
			return _coll1.ToMap()
		}()).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _8_v2 _dafny.Int
			_8_v2 = interface{}(_compr_0).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(46), _dafny.IntOfInt64(483))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _6_v0 _dafny.Int
					_6_v0 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(46)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(483)) < 0) {
						_coll3.Add(Companion_Default___.SafeModuloInt(_6_v0, (func() _dafny.Map {
							var _coll4 = _dafny.NewMapBuilder()
							_ = _coll4
							for _iter4 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-312), (_dafny.MultiSetOf(_dafny.IntOfInt64(220))).Cardinality())).Elements()); ; {
								_compr_4, _ok4 := _iter4()
								if !_ok4 {
									break
								}
								var _7_v1 _dafny.Int
								_7_v1 = interface{}(_compr_4).(_dafny.Int)
								if (_dafny.MultiSetOf(_dafny.IntOfInt64(-312), (_dafny.MultiSetOf(_dafny.IntOfInt64(220))).Cardinality())).Contains(_7_v1) {
									_coll4.Add(Companion_Default___.SafeDivisionInt(_7_v1, _dafny.IntOfInt64(-651)), _dafny.IntOfInt64(555))
								}
							}
							return _coll4.ToMap()
						}()).Cardinality()), _dafny.IntOfInt64(899))
					}
				}
				return _coll3.ToMap()
			}()).Contains(_8_v2) {
				_coll0.Add((_8_v2).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.SetOf(_dafny.IntOfInt64(899), _dafny.IntOfInt64(703), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("llm")).Cardinality()))).Cardinality())).Cardinality(), _dafny.IntOfInt64(-706))).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality())))))))
	_ = _source0
	if _source0.Is_DC13() {
		var _9___mcc_h0 bool = _source0.Get_().(D7_DC13).Cf17
		_ = _9___mcc_h0
		var _10_cf17 bool = _9___mcc_h0
		_ = _10_cf17
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(37))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		}))
	} else if _source0.Is_DC12() {
		var _12___mcc_h1 _dafny.MultiSet = _source0.Get_().(D7_DC12).Cf16
		_ = _12___mcc_h1
		var _13_cf16 _dafny.MultiSet = _12___mcc_h1
		_ = _13_cf16
		return _dafny.UnicodeSeqOfUtf8Bytes("y")
	} else {
		var _14___mcc_h2 D7 = _source0.Get_().(D7_DC14).Cf18
		_ = _14___mcc_h2
		var _15_cf18 D7 = _14___mcc_h2
		_ = _15_cf18
		return _dafny.UnicodeSeqOfUtf8Bytes("hvccdvgrp")
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(916), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-611))).Cardinality(), false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(991), !(true)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D6 = (func() D6 {
		if !(true) {
			return Companion_D6_.Create_DC11_(_dafny.IntOfInt64(-673))
		}
		return Companion_D6_.Create_DC11_((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gqh")).Cardinality()))).Cardinality())
	})()
	_ = _source1
	if _source1.Is_DC11() {
		var _16___mcc_h0 _dafny.Int = _source1.Get_().(D6_DC11).Cf15
		_ = _16___mcc_h0
		var _17_cf15 _dafny.Int = _16___mcc_h0
		_ = _17_cf15
		return _dafny.CodePoint('g')
	} else {
		var _18___mcc_h1 *C0 = _source1.Get_().(D6_DC10).Cf14
		_ = _18___mcc_h1
		var _19_cf14 *C0 = _18___mcc_h1
		_ = _19_cf14
		return _dafny.CodePoint('r')
	}
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(_dafny.CodePoint('i')))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(_dafny.CodePoint('g'))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.IntOfInt64(-514))).Keys().Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _20_v0 _dafny.CodePoint
			_20_v0 = interface{}(_compr_5).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.IntOfInt64(-514))).Contains(_20_v0) {
				_coll5.Add(_20_v0)
			}
		}
		return _coll5.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(134))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_21_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(943)
			})))
		}
		return _dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(7), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(520), !(true))).Cardinality()), _dafny.IntOfInt64(496), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-47)))))
	})()).Union(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if !(true) {
			return _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(975)))
		}
		return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(907))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_22_i1 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-786)))
		})), _dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(811), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_23_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()), _dafny.IntOfInt64(948))).Cardinality(), (func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(481), _dafny.IntOfInt64(593))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _24_v1 _dafny.Int
					_24_v1 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(481)).Cmp(_24_v1) <= 0) && ((_24_v1).Cmp(_dafny.IntOfInt64(593)) < 0) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_24_v1, _dafny.IntOfInt64(843)), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
					}
				}
				return _coll7.ToMap()
			}()).Keys().Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _25_v0 _dafny.Int
				_25_v0 = interface{}(_compr_6).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(481), _dafny.IntOfInt64(593))); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _24_v1 _dafny.Int
						_24_v1 = interface{}(_compr_8).(_dafny.Int)
						if ((_dafny.IntOfInt64(481)).Cmp(_24_v1) <= 0) && ((_24_v1).Cmp(_dafny.IntOfInt64(593)) < 0) {
							_coll8.Add(Companion_Default___.SafeModuloInt(_24_v1, _dafny.IntOfInt64(843)), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
						}
					}
					return _coll8.ToMap()
				}()).Contains(_25_v0) {
					_coll6.Add(Companion_Default___.SafeDivisionInt(_25_v0, _dafny.IntOfInt64(-620)), _dafny.UnicodeSeqOfUtf8Bytes("vyeoofmh"))
				}
			}
			return _coll6.ToMap()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lmjx")).Cardinality())))
	})()))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Intersection(_dafny.SetOf(false, false))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-173))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_26_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((_dafny.SetOf(!(false), true)).Cardinality())
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vetmckar")).Cardinality()))).Cardinality())).Cardinality()))), _dafny.SeqOf((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf(true, true), _dafny.MultiSetOf(true))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _27_v0 _dafny.MultiSet
			_27_v0 = interface{}(_compr_9).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf(true, true), _dafny.MultiSetOf(true)), _27_v0) {
				_coll9.Add(_27_v0, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()))).Cardinality()))
			}
		}
		return _coll9.ToMap()
	}()).Cardinality(), (_dafny.SetOf(!(!(!(true))))).Cardinality(), _dafny.IntOfInt64(670)), _dafny.SeqOf(_dafny.IntOfInt64(818), _dafny.IntOfUint32(((Companion_D10_.Create_DC21_(_dafny.SeqOf(!(true)))).Dtor_cf33()).Cardinality()), _dafny.IntOfInt64(699)))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-782), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-698))).Cardinality())).Cardinality(), !(true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-381), !(false)))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false, true))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) D1 {
	if (_dafny.SetOf(_dafny.SeqOf(false))).IsDisjointFrom(_dafny.SetOf(_dafny.SeqOf(false, true))) {
		if true {
			return Companion_D1_.Create_DC3_(_dafny.MultiSetOf(!(false), false, false), false, true)
		} else {
			return Companion_D1_.Create_DC3_(_dafny.MultiSetOf(true), false, false)
		}
	} else {
		return Companion_D1_.Create_DC3_(_dafny.MultiSetOf(!(!(true)), true), !((Companion_D7_.Create_DC13_(true)).Dtor_cf17()), true)
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Map, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int, _dafny.Set) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 _dafny.Set = _dafny.EmptySet
	_ = r3
	var _28_v0 _dafny.Int
	_ = _28_v0
	_28_v0 = _dafny.IntOfInt64(232)
	var _29_v1 _dafny.Sequence
	_ = _29_v1
	_29_v1 = _dafny.SeqOf(_28_v0)
	var _30_v2 D8
	_ = _30_v2
	_30_v2 = Companion_D8_.Create_DC15_(_29_v1)
	_29_v1 = (_30_v2).Dtor_cf19()
	var _31_v3 D7
	_ = _31_v3
	_31_v3 = Companion_D7_.Create_DC13_(p1)
	var _32_v4 D7
	_ = _32_v4
	_32_v4 = Companion_D7_.Create_DC14_(_31_v3)
	var _33_v5 _dafny.Map
	_ = _33_v5
	_33_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v4, p1)
	_33_v5 = (_33_v5).Update(_32_v4, false)
	var _34_v6 _dafny.Map
	_ = _34_v6
	_34_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _28_v0)
	var _35_v7 _dafny.Set
	_ = _35_v7
	_35_v7 = _dafny.SetOf(_34_v6, _34_v6)
	var _36_i0 _dafny.Int
	_ = _36_i0
	_36_i0 = _dafny.Zero
	{
		for (_35_v7).IsProperSubsetOf(func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_35_v7).Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _53_v8 _dafny.Map
				_53_v8 = interface{}(_compr_10).(_dafny.Map)
				if (_35_v7).Contains(_53_v8) {
					_coll10.Add(_53_v8)
				}
			}
			return _coll10.ToSet()
		}()) {
			{
				if (_36_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_36_i0 = (_36_i0).Plus(_dafny.One)
				var _37_v9 _dafny.Array
				_ = _37_v9
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_0
				var _nw0 _dafny.Array
				_ = _nw0
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw0 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.Int = (func(_38_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_39_i1 _dafny.Int) _dafny.Int {
							return (_39_i1).Times(_38_v0)
						}
					})(_28_v0)
					_ = _init0
					var _element0_0 = _init0(_dafny.Zero)
					_ = _element0_0
					_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
					(_nw0).ArraySet1(_element0_0, 0)
					var _nativeLen0_0 = (_len0_0).Int()
					_ = _nativeLen0_0
					for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
						(_nw0).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
					}
				}
				_37_v9 = _nw0
				_37_v9 = (Companion_D1_.Create_DC1_(_37_v9)).Dtor_cf1()
				var _40_v10 _dafny.Array
				_ = _40_v10
				var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_1
				var _nw1 _dafny.Array
				_ = _nw1
				if _len0_1.Cmp(_dafny.Zero) == 0 {
					_nw1 = _dafny.NewArray(_len0_1)
				} else {
					var _init1 func(_dafny.Int) _dafny.MultiSet = (func(_41_v1 _dafny.Sequence, _42_v0 _dafny.Int, _43_p1 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_44_i2 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_41_v1)).Difference(_dafny.MultiSetOf(_41_v1, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_42_v0, _dafny.IntOfUint32((_41_v1).Cardinality())), (Companion_Default___.SafeIndex(_42_v0, _dafny.IntOfUint32((_dafny.SeqOf(_42_v0, _dafny.IntOfUint32((_41_v1).Cardinality()))).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_43_p1)).Cardinality())))))
						}
					})(_29_v1, _28_v0, p1)
					_ = _init1
					var _element0_1 = _init1(_dafny.Zero)
					_ = _element0_1
					_nw1 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
					(_nw1).ArraySet1(_element0_1, 0)
					var _nativeLen0_1 = (_len0_1).Int()
					_ = _nativeLen0_1
					for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
						(_nw1).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
					}
				}
				_40_v10 = _nw1
				var _45_v11 _dafny.Sequence
				_ = _45_v11
				_45_v11 = _dafny.SeqOf(_29_v1)
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_40_v10), 0))
				_ = _index0
				(_40_v10).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_28_v0)), _45_v11)), (_index0).Int())
				var _46_v12 *C0
				_ = _46_v12
				var _nw2 *C0 = New_C0_()
				_ = _nw2
				_nw2.Ctor__(p1)
				_46_v12 = _nw2
				var _47_v13 _dafny.Map
				_ = _47_v13
				_47_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v12, _29_v1)
				var _48_v14 _dafny.MultiSet
				_ = _48_v14
				_48_v14 = _dafny.MultiSetOf(_dafny.SeqOf(_28_v0, _28_v0), (func() _dafny.Sequence {
					if (_47_v13).Contains(_46_v12) {
						return (_47_v13).Get(_46_v12).(_dafny.Sequence)
					}
					return _29_v1
				})())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_40_v10), 0))
				_ = _index1
				(_40_v10).ArraySet1((Companion_D7_.Create_DC12_(_48_v14)).Dtor_cf16(), (_index1).Int())
				var _49_v15 bool
				_ = _49_v15
				_49_v15 = false
				var _50_v16 _dafny.CodePoint
				_ = _50_v16
				_50_v16 = _dafny.CodePoint('r')
				var _51_v17 _dafny.CodePoint
				_ = _51_v17
				_51_v17 = _50_v16
				var _52_v18 _dafny.Sequence
				_ = _52_v18
				_52_v18 = _dafny.SeqOf((_51_v17))
				var _rhs0 bool = !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm7(_dafny.IntOfUint32((_52_v18).Cardinality()), _28_v0, globalState), (func() _dafny.Sequence {
					if p1 {
						return _dafny.UnicodeSeqOfUtf8Bytes("wrtno")
					}
					return _52_v18
				})())
				_ = _rhs0
				var _rhs1 *C0 = _46_v12
				_ = _rhs1
				_49_v15 = _rhs0
				_46_v12 = _rhs1
				_29_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_28_v0), _29_v1)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _54_v19 bool
	_ = _54_v19
	_54_v19 = true
	_54_v19 = p1
	var _55_i3 _dafny.Int
	_ = _55_i3
	_55_i3 = _dafny.Zero
	{
		for p1 {
			{
				if (_55_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_55_i3 = (_55_i3).Plus(_dafny.One)
				_54_v19 = p1
				r0 = _28_v0
				var _56_v20 _dafny.Sequence
				_ = _56_v20
				_56_v20 = _dafny.SeqOf(false)
				var _57_v21 _dafny.Sequence
				_ = _57_v21
				_57_v21 = _dafny.SeqOf(_32_v4, Companion_D7_.Create_DC14_(_31_v3))
				var _58_v22 _dafny.Map
				_ = _58_v22
				_58_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v0, _28_v0)
				var _59_v23 D8
				_ = _59_v23
				_59_v23 = Companion_D8_.Create_DC16_(_54_v19, _dafny.MultiSetFromSeq(_56_v20), _57_v21, _58_v22)
				var _60_v24 D8
				_ = _60_v24
				_60_v24 = Companion_D8_.Create_DC18_(_59_v23)
				var _61_v25 D8
				_ = _61_v25
				_61_v25 = Companion_D8_.Create_DC18_(_59_v23)
				var _62_v26 D8
				_ = _62_v26
				_62_v26 = Companion_D8_.Create_DC18_(_59_v23)
				var _63_v27 D8
				_ = _63_v27
				_63_v27 = Companion_D8_.Create_DC18_(_62_v26)
				var _64_v28 D8
				_ = _64_v28
				_64_v28 = Companion_D8_.Create_DC18_(_62_v26)
				var _source2 D8 = _64_v28
				_ = _source2
				if _source2.Is_DC16() {
					var _65___mcc_h0 bool = _source2.Get_().(D8_DC16).Cf20
					_ = _65___mcc_h0
					var _66___mcc_h1 _dafny.MultiSet = _source2.Get_().(D8_DC16).Cf21
					_ = _66___mcc_h1
					var _67___mcc_h2 _dafny.Sequence = _source2.Get_().(D8_DC16).Cf22
					_ = _67___mcc_h2
					var _68___mcc_h3 _dafny.Map = _source2.Get_().(D8_DC16).Cf23
					_ = _68___mcc_h3
					var _69_cf23 _dafny.Map = _68___mcc_h3
					_ = _69_cf23
					var _70_cf22 _dafny.Sequence = _67___mcc_h2
					_ = _70_cf22
					var _71_cf21 _dafny.MultiSet = _66___mcc_h1
					_ = _71_cf21
					var _72_cf20 bool = _65___mcc_h0
					_ = _72_cf20
					var _73_v29 _dafny.Array
					_ = _73_v29
					var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
					_ = _nw3
					_73_v29 = _nw3
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_73_v29), 0))
					_ = _index2
					(_73_v29).ArraySet1(Companion_Default___.Fm15(globalState), (_index2).Int())
					var _74_v30 _dafny.Map
					_ = _74_v30
					_74_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v0, _72_cf20)
					var _75_v31 _dafny.Sequence
					_ = _75_v31
					_75_v31 = _dafny.SeqOf(_74_v30, _74_v30)
					var _76_v32 _dafny.Sequence
					_ = _76_v32
					_76_v32 = _dafny.UnicodeSeqOfUtf8Bytes("nixsur")
					var _77_v33 _dafny.Map
					_ = _77_v33
					_77_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v32, _28_v0)
					var _78_v34 _dafny.Sequence
					_ = _78_v34
					_78_v34 = _dafny.SeqOf(_77_v33, _77_v33)
					var _79_v35 _dafny.Set
					_ = _79_v35
					_79_v35 = _dafny.SetOf(_54_v19, _54_v19, p1)
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_73_v29), 0))
					_ = _index3
					(_73_v29).ArraySet1(((_75_v31).Select((Companion_Default___.SafeIndex(((_78_v34).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_28_v0, _28_v0)).Cardinality(), _dafny.IntOfUint32((_78_v34).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32((_75_v31).Cardinality()))).Uint32()).(_dafny.Map)).Merge(((_74_v30).Update((_dafny.Zero).Minus(_28_v0), p1)).Update(_28_v0, Companion_Default___.Fm1(_dafny.UnicodeSeqOfUtf8Bytes("hmgxqq"), _28_v0, _79_v35, globalState))), (_index3).Int())
					var _80_v36 T0
					_ = _80_v36
					var _nw4 *C0 = New_C0_()
					_ = _nw4
					_nw4.Ctor__(p1)
					_80_v36 = _nw4
					_80_v36 = _80_v36
					var _81_v37 _dafny.Set
					_ = _81_v37
					_81_v37 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_74_v30).Update(_28_v0, _72_cf20), _54_v19)).Cardinality()))
					var _82_v38 _dafny.Array
					_ = _82_v38
					var _nwElement0_0 _dafny.Set = _81_v37
					_ = _nwElement0_0
					var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(4))
					_ = _nw5
					(_nw5).ArraySet1(_nwElement0_0, 0)
					(_nw5).ArraySet1(_81_v37, 1)
					(_nw5).ArraySet1(_dafny.SetOf(_28_v0, _28_v0), 2)
					(_nw5).ArraySet1(_81_v37, 3)
					_82_v38 = _nw5
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_82_v38), 0))
					_ = _index4
					(_82_v38).ArraySet1(_81_v37, (_index4).Int())
					var _83_v39 _dafny.Array
					_ = _83_v39
					var _len0_2 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_2
					var _nw6 _dafny.Array
					_ = _nw6
					if _len0_2.Cmp(_dafny.Zero) == 0 {
						_nw6 = _dafny.NewArray(_len0_2)
					} else {
						var _init2 func(_dafny.Int) _dafny.Int = (func(_84_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_85_i4 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeDivisionInt(_85_i4, _84_v0)
							}
						})(_28_v0)
						_ = _init2
						var _element0_2 = _init2(_dafny.Zero)
						_ = _element0_2
						_nw6 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
						(_nw6).ArraySet1(_element0_2, 0)
						var _nativeLen0_2 = (_len0_2).Int()
						_ = _nativeLen0_2
						for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
							(_nw6).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
						}
					}
					_83_v39 = _nw6
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_83_v39), 0))
					_ = _index5
					(_83_v39).ArraySet1(Companion_Default___.SafeModuloInt(_28_v0, (_58_v22).Cardinality()), (_index5).Int())
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_83_v39), 0))
					_ = _index6
					(_83_v39).ArraySet1(_dafny.IntOfUint32((_29_v1).Cardinality()), (_index6).Int())
					var _86_v40 *C0
					_ = _86_v40
					var _nw7 *C0 = New_C0_()
					_ = _nw7
					_nw7.Ctor__(p1)
					_86_v40 = _nw7
					var _87_v41 _dafny.Sequence
					_ = _87_v41
					_87_v41 = _dafny.SeqOf(_86_v40)
					var _88_v42 _dafny.Sequence
					_ = _88_v42
					_88_v42 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_87_v41).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_87_v41).Cardinality()))).Uint32()).(*C0)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_87_v41).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_87_v41).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_87_v41).Cardinality()))).Uint32()).(*C0))).Cardinality()))).Uint32(), _86_v40), _87_v41)
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_82_v38), 0))
					_ = _index7
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_83_v39), 0))
					_ = _index8
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_83_v39), 0))
					_ = _index9
					var _rhs2 _dafny.Set = ((_81_v37).Difference(Companion_Default___.Fm16(globalState))).Intersection(_81_v37)
					_ = _rhs2
					var _rhs3 _dafny.Sequence = _76_v32
					_ = _rhs3
					var _rhs4 _dafny.Int = (_29_v1).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_29_v1).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs4
					var _rhs5 _dafny.Int = (_28_v0).Times(_dafny.IntOfUint32((_76_v32).Cardinality()))
					_ = _rhs5
					var _rhs6 _dafny.Sequence = _88_v42
					_ = _rhs6
					var _lhs0 _dafny.Array = _82_v38
					_ = _lhs0
					var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_82_v38), 0))
					_ = _lhs1
					var _lhs2 _dafny.Array = _83_v39
					_ = _lhs2
					var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_83_v39), 0))
					_ = _lhs3
					var _lhs4 _dafny.Array = _83_v39
					_ = _lhs4
					var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_83_v39), 0))
					_ = _lhs5
					(_lhs0).ArraySet1(_rhs2, (_lhs1).Int())
					_76_v32 = _rhs3
					(_lhs2).ArraySet1(_rhs4, (_lhs3).Int())
					(_lhs4).ArraySet1(_rhs5, (_lhs5).Int())
					_88_v42 = _rhs6
					var _89_v43 _dafny.MultiSet
					_ = _89_v43
					_89_v43 = _dafny.MultiSetOf(_dafny.SeqOf(_28_v0, (_83_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_83_v39), 0))).Int()).(_dafny.Int)), _29_v1, _29_v1, _29_v1)
					var _90_v44 D7
					_ = _90_v44
					_90_v44 = Companion_D7_.Create_DC12_(_89_v43)
					var _pat_let_tv0 = _89_v43
					_ = _pat_let_tv0
					var _rhs7 bool = (func() bool {
						if _80_v36.F0() {
							return _72_cf20
						}
						return _80_v36.F0()
					})()
					_ = _rhs7
					var _rhs8 _dafny.Int = (_83_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_83_v39), 0))).Int()).(_dafny.Int)
					_ = _rhs8
					var _rhs9 D7 = func(_pat_let0_0 D7) D7 {
						return func(_91_dt__update__tmp_h0 D7) D7 {
							return func(_pat_let1_0 _dafny.MultiSet) D7 {
								return func(_92_dt__update_hcf16_h0 _dafny.MultiSet) D7 {
									return Companion_D7_.Create_DC12_(_92_dt__update_hcf16_h0)
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_90_v44)
					_ = _rhs9
					_72_cf20 = _rhs7
					r2 = _rhs8
					_90_v44 = _rhs9
				} else if _source2.Is_DC17() {
					var _93___mcc_h4 _dafny.Sequence = _source2.Get_().(D8_DC17).Cf24
					_ = _93___mcc_h4
					var _94___mcc_h5 _dafny.Int = _source2.Get_().(D8_DC17).Cf25
					_ = _94___mcc_h5
					var _95___mcc_h6 _dafny.Int = _source2.Get_().(D8_DC17).Cf26
					_ = _95___mcc_h6
					var _96___mcc_h7 _dafny.Int = _source2.Get_().(D8_DC17).Cf27
					_ = _96___mcc_h7
					var _97_cf27 _dafny.Int = _96___mcc_h7
					_ = _97_cf27
					var _98_cf26 _dafny.Int = _95___mcc_h6
					_ = _98_cf26
					var _99_cf25 _dafny.Int = _94___mcc_h5
					_ = _99_cf25
					var _100_cf24 _dafny.Sequence = _93___mcc_h4
					_ = _100_cf24
					var _101_v45 _dafny.CodePoint
					_ = _101_v45
					_101_v45 = _dafny.CodePoint('f')
					var _102_v46 D3
					_ = _102_v46
					_102_v46 = Companion_D3_.Create_DC6_(_28_v0, _98_cf26, p1)
					var _103_v47 _dafny.Set
					_ = _103_v47
					_103_v47 = _dafny.SetOf((_102_v46).Dtor_cf10(), p1, p1)
					var _104_v48 _dafny.Map
					_ = _104_v48
					_104_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_29_v1).Select((Companion_Default___.SafeIndex(_99_cf25, _dafny.IntOfUint32((_29_v1).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("efgeywy"), (Companion_Default___.SafeIndex(_99_cf25, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efgeywy")).Cardinality()))).Uint32(), _101_v45), _28_v0, _103_v47, globalState))
					_104_v48 = (_104_v48).Update(_97_cf27, _54_v19)
					_54_v19 = (_99_cf25).Cmp((_29_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((Companion_D6_.Create_DC11_(_97_cf27)).Dtor_cf15()), _dafny.IntOfUint32((_29_v1).Cardinality()))).Uint32()).(_dafny.Int)) <= 0
					var _105_v49 *C0
					_ = _105_v49
					var _nw8 *C0 = New_C0_()
					_ = _nw8
					_nw8.Ctor__(_54_v19)
					_105_v49 = _nw8
					_105_v49 = _105_v49
					var _106_v50 _dafny.Array
					_ = _106_v50
					var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
					_ = _nw9
					_106_v50 = _nw9
					_106_v50 = _106_v50
				} else if _source2.Is_DC15() {
					var _107___mcc_h8 _dafny.Sequence = _source2.Get_().(D8_DC15).Cf19
					_ = _107___mcc_h8
					var _108_cf19 _dafny.Sequence = _107___mcc_h8
					_ = _108_cf19
					var _109_v51 _dafny.Int
					_ = _109_v51
					_109_v51 = _28_v0
					var _110_v52 _dafny.Map
					_ = _110_v52
					_110_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v19, _109_v51)
					var _111_v53 _dafny.Map
					_ = _111_v53
					_111_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _54_v19)
					var _112_v54 _dafny.Map
					_ = _112_v54
					_112_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v52, (func() bool {
						if (_111_v53).Contains(_54_v19) {
							return (_111_v53).Get(_54_v19).(bool)
						}
						return p1
					})())
					_112_v54 = (_112_v54).Update(_110_v52, p1)
					_54_v19 = (func() bool {
						if (_111_v53).Contains(_54_v19) {
							return (_111_v53).Get(_54_v19).(bool)
						}
						return !(!(p1)) || (p1)
					})()
					_54_v19 = p1
					var _113_v55 _dafny.Set
					_ = _113_v55
					_113_v55 = _dafny.SetOf(_54_v19, _54_v19)
					_54_v19 = (_113_v55).IsProperSubsetOf((_113_v55).Intersection(_113_v55))
				} else {
					var _114___mcc_h9 D8 = _source2.Get_().(D8_DC18).Cf28
					_ = _114___mcc_h9
					var _115_cf28 D8 = _114___mcc_h9
					_ = _115_cf28
					var _116_v56 _dafny.MultiSet
					_ = _116_v56
					_116_v56 = _dafny.MultiSetOf((_56_v20).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_56_v20).Cardinality()))).Uint32()).(bool), p1)
					var _117_v57 _dafny.Set
					_ = _117_v57
					_117_v57 = _dafny.SetOf(p1, p1)
					_54_v19 = !(Companion_Default___.Fm1(Companion_Default___.Fm7((func() _dafny.Int {
						if (_116_v56).Contains(_54_v19) {
							return (_116_v56).Multiplicity(_54_v19)
						}
						return _28_v0
					})(), _28_v0, globalState), _28_v0, _117_v57, globalState))
					_29_v1 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if p1 {
							return _29_v1
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(907))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg11 _dafny.Int) interface{} {
								return coer11(arg11)
							}
						}((func(_118_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_119_i5 _dafny.Int) _dafny.Int {
								return _118_v0
							}
						})(_28_v0)))
					})(), (Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if p1 {
							return _29_v1
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(907))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg12 _dafny.Int) interface{} {
								return coer12(arg12)
							}
						}((func(_120_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_121_i5 _dafny.Int) _dafny.Int {
								return _120_v0
							}
						})(_28_v0)))
					})()).Cardinality()))).Uint32(), _28_v0)
					var _122_v58 *C0
					_ = _122_v58
					var _nw10 *C0 = New_C0_()
					_ = _nw10
					_nw10.Ctor__(p1)
					_122_v58 = _nw10
					var _123_v59 _dafny.Map
					_ = _123_v59
					_123_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v58, p1)
					var _124_v60 _dafny.Map
					_ = _124_v60
					_124_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_123_v59, p1)
					var _125_v61 *C0
					_ = _125_v61
					var _nw11 *C0 = New_C0_()
					_ = _nw11
					_nw11.Ctor__((func() bool {
						if (_124_v60).Contains(_123_v59) {
							return (_124_v60).Get(_123_v59).(bool)
						}
						return p1
					})())
					_125_v61 = _nw11
					var _126_v62 _dafny.Sequence
					_ = _126_v62
					_126_v62 = _dafny.UnicodeSeqOfUtf8Bytes("owkgj")
					var _127_v63 _dafny.Map
					_ = _127_v63
					_127_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v0, _126_v62)
					var _128_v64 _dafny.CodePoint
					_ = _128_v64
					_128_v64 = _dafny.CodePoint('d')
					var _129_v65 _dafny.Map
					_ = _129_v65
					_129_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v19, _128_v64)
					var _130_v66 _dafny.Map
					_ = _130_v66
					_130_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if _54_v19 {
							return (func() _dafny.Sequence {
								if (_127_v63).Contains(_28_v0) {
									return (_127_v63).Get(_28_v0).(_dafny.Sequence)
								}
								return _126_v62
							})()
						}
						return _dafny.Companion_Sequence_.Update(_126_v62, (Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_126_v62).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
							if (_129_v65).Contains(p1) {
								return (_129_v65).Get(p1).(_dafny.CodePoint)
							}
							return _128_v64
						})())
					})()).Cardinality()), _125_v61)
					var _131_v67 D6
					_ = _131_v67
					_131_v67 = Companion_D6_.Create_DC10_(_125_v61)
					_125_v61 = (func() *C0 {
						if (_130_v66).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_56_v20, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.IntOfUint32((_56_v20).Cardinality()))).Uint32(), _54_v19), _dafny.SeqOf(_54_v19, p1, true, _54_v19))).Cardinality())) {
							return (_130_v66).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_56_v20, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.IntOfUint32((_56_v20).Cardinality()))).Uint32(), _54_v19), _dafny.SeqOf(_54_v19, p1, true, _54_v19))).Cardinality())).(*C0)
						}
						return (_131_v67).Dtor_cf14()
					})()
					_54_v19 = p1
				}
				var _132_v68 _dafny.Sequence
				_ = _132_v68
				_132_v68 = _dafny.UnicodeSeqOfUtf8Bytes("emnf")
				if Companion_Default___.Fm1(_132_v68, _28_v0, Companion_Default___.Fm13(_54_v19, _54_v19, p1, globalState), globalState) {
					_54_v19 = _54_v19
					var _133_v69 _dafny.Set
					_ = _133_v69
					_133_v69 = _dafny.SetOf(_54_v19, _54_v19)
					var _134_v70 *C0
					_ = _134_v70
					var _nw12 *C0 = New_C0_()
					_ = _nw12
					_nw12.Ctor__((p1) || (Companion_Default___.Fm1(_132_v68, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ertqblss")).Cardinality()), _133_v69, globalState)))
					_134_v70 = _nw12
					var _135_v71 _dafny.Array
					_ = _135_v71
					var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
					_ = _nw13
					_135_v71 = _nw13
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_135_v71), 0))
					_ = _index10
					(_135_v71).ArraySet1(_28_v0, (_index10).Int())
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_135_v71), 0))
					_ = _index11
					var _rhs10 bool = _dafny.Companion_Sequence_.IsPrefixOf(_29_v1, _29_v1)
					_ = _rhs10
					var _rhs11 _dafny.Sequence = _132_v68
					_ = _rhs11
					var _rhs12 _dafny.Int = _28_v0
					_ = _rhs12
					var _lhs6 _dafny.Array = _135_v71
					_ = _lhs6
					var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_135_v71), 0))
					_ = _lhs7
					_54_v19 = _rhs10
					_132_v68 = _rhs11
					(_lhs6).ArraySet1(_rhs12, (_lhs7).Int())
					var _136_v72 _dafny.Array
					_ = _136_v72
					var _len0_3 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_3
					var _nw14 _dafny.Array
					_ = _nw14
					if _len0_3.Cmp(_dafny.Zero) == 0 {
						_nw14 = _dafny.NewArray(_len0_3)
					} else {
						var _init3 func(_dafny.Int) bool = func(_137_i6 _dafny.Int) bool {
							return false
						}
						_ = _init3
						var _element0_3 = _init3(_dafny.Zero)
						_ = _element0_3
						_nw14 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
						(_nw14).ArraySet1(_element0_3, 0)
						var _nativeLen0_3 = (_len0_3).Int()
						_ = _nativeLen0_3
						for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
							(_nw14).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
						}
					}
					_136_v72 = _nw14
					var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_136_v72), 0))
					_ = _index12
					(_136_v72).ArraySet1(p1, (_index12).Int())
					var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_136_v72), 0))
					_ = _index13
					(_136_v72).ArraySet1(p1, (_index13).Int())
					var _138_v73 *C0
					_ = _138_v73
					var _nw15 *C0 = New_C0_()
					_ = _nw15
					_nw15.Ctor__((Companion_Default___.Fm17(_28_v0, _dafny.IntOfInt64(755), p1, globalState)).Dtor_cf5())
					_138_v73 = _nw15
				} else {
					var _139_v74 _dafny.Array
					_ = _139_v74
					var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
					_ = _nw16
					_139_v74 = _nw16
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_139_v74), 0))
					_ = _index14
					(_139_v74).ArraySet1(_28_v0, (_index14).Int())
					var _140_v75 _dafny.CodePoint
					_ = _140_v75
					_140_v75 = _dafny.CodePoint('d')
					var _141_v76 _dafny.Sequence
					_ = _141_v76
					_141_v76 = _dafny.SeqOf(_132_v68)
					var _142_v77 _dafny.Int
					_ = _142_v77
					_142_v77 = _28_v0
					var _143_v78 _dafny.Map
					_ = _143_v78
					_143_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v0, _141_v76)
					var _144_v79 _dafny.Sequence
					_ = _144_v79
					_144_v79 = _dafny.SeqOf((func() _dafny.Sequence {
						if (_143_v78).Contains(_dafny.IntOfInt64(-272)) {
							return (_143_v78).Get(_dafny.IntOfInt64(-272)).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(_141_v76, (Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_141_v76).Cardinality()))).Uint32(), _132_v68)
					})())
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_139_v74), 0))
					_ = _index15
					var _rhs13 _dafny.Int = (_142_v77)
					_ = _rhs13
					var _rhs14 _dafny.CodePoint = _140_v75
					_ = _rhs14
					var _rhs15 _dafny.CodePoint = (_132_v68).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_132_v68).Cardinality()))).Uint32()).(_dafny.CodePoint)
					_ = _rhs15
					var _rhs16 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_144_v79).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2(p1, globalState), _dafny.IntOfUint32((_144_v79).Cardinality()))).Uint32()).(_dafny.Sequence), _141_v76)
					_ = _rhs16
					var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt(_28_v0, (_29_v1).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_29_v1).Cardinality()))).Uint32()).(_dafny.Int))
					_ = _rhs17
					var _lhs8 _dafny.Array = _139_v74
					_ = _lhs8
					var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_139_v74), 0))
					_ = _lhs9
					(_lhs8).ArraySet1(_rhs13, (_lhs9).Int())
					_140_v75 = _rhs14
					_140_v75 = _rhs15
					_141_v76 = _rhs16
					_28_v0 = _rhs17
					var _145_v80 _dafny.Array
					_ = _145_v80
					var _nw17 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
					_ = _nw17
					_145_v80 = _nw17
					var _nw18 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
					_ = _nw18
					_145_v80 = _nw18
					var _146_v82 D9
					_ = _146_v82
					_146_v82 = Companion_D9_.Create_DC19_(func() _dafny.Map {
						var _coll11 = _dafny.NewMapBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-359), _dafny.IntOfInt64(203))); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _147_v81 _dafny.Int
							_147_v81 = interface{}(_compr_11).(_dafny.Int)
							if ((_dafny.IntOfInt64(-359)).Cmp(_147_v81) <= 0) && ((_147_v81).Cmp(_dafny.IntOfInt64(203)) < 0) {
								_coll11.Add(Companion_Default___.SafeModuloInt(_147_v81, _28_v0), _56_v20)
							}
						}
						return _coll11.ToMap()
					}())
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_139_v74), 0))
					_ = _index16
					(_139_v74).ArraySet1(((_146_v82).Dtor_cf29()).Cardinality(), (_index16).Int())
					_54_v19 = _54_v19
					var _148_v83 *C0
					_ = _148_v83
					var _nw19 *C0 = New_C0_()
					_ = _nw19
					_nw19.Ctor__(_54_v19)
					_148_v83 = _nw19
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _149_i7 _dafny.Int
	_ = _149_i7
	_149_i7 = _dafny.Zero
	{
		for !(!(_54_v19)) {
			{
				if (_149_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_149_i7 = (_149_i7).Plus(_dafny.One)
				var _150_v84 _dafny.Sequence
				_ = _150_v84
				_150_v84 = _dafny.SeqOf(!(p1), _54_v19)
				var _151_v85 _dafny.MultiSet
				_ = _151_v85
				_151_v85 = _dafny.MultiSetOf((_150_v84).Select((Companion_Default___.SafeIndex(_28_v0, _dafny.IntOfUint32((_150_v84).Cardinality()))).Uint32()).(bool))
				var _152_v86 D1
				_ = _152_v86
				_152_v86 = Companion_D1_.Create_DC3_(_151_v85, _54_v19, _54_v19)
				var _153_v87 _dafny.Set
				_ = _153_v87
				_153_v87 = _dafny.SetOf(p1, _54_v19, p1, (_152_v86).Dtor_cf4())
				_54_v19 = Companion_Default___.Fm1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(253))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}(func(_154_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				})), _28_v0, _153_v87, globalState)
				if (_dafny.IntOfInt64(448)).Cmp(_28_v0) == 0 {
					var _155_v88 _dafny.Sequence
					_ = _155_v88
					_155_v88 = _dafny.UnicodeSeqOfUtf8Bytes("nlwntcgy")
					var _156_v89 *C0
					_ = _156_v89
					var _nw20 *C0 = New_C0_()
					_ = _nw20
					_nw20.Ctor__(p1)
					_156_v89 = _nw20
					var _157_v90 _dafny.CodePoint
					_ = _157_v90
					_157_v90 = _dafny.CodePoint('m')
					_155_v88 = _dafny.Companion_Sequence_.Update(_155_v88, (Companion_Default___.SafeIndex((_28_v0).Times((_dafny.SetOf(_156_v89)).Cardinality()), _dafny.IntOfUint32((_155_v88).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
						if p1 {
							return _157_v90
						}
						return _dafny.CodePoint('j')
					})())
					var _158_v91 _dafny.Array
					_ = _158_v91
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(19)
					_ = _len0_4
					var _nw21 _dafny.Array
					_ = _nw21
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw21 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) D8 = (func(_159_v2 D8) func(_dafny.Int) D8 {
							return func(_160_i9 _dafny.Int) D8 {
								return _159_v2
							}
						})(_30_v2)
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw21 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw21).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw21).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_158_v91 = _nw21
					var _161_v92 _dafny.Set
					_ = _161_v92
					_161_v92 = _dafny.SetOf(_158_v91)
					_54_v19 = ((func() _dafny.Int {
						if _54_v19 {
							return _dafny.IntOfInt64(521)
						}
						return _28_v0
					})()).Cmp(((_161_v92).Intersection(_161_v92)).Cardinality()) != 0
					_29_v1 = _29_v1
					var _162_v93 _dafny.Array
					_ = _162_v93
					var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
					_ = _nw22
					_162_v93 = _nw22
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_162_v93), 0))
					_ = _index17
					(_162_v93).ArraySet1(_54_v19, (_index17).Int())
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_162_v93), 0))
					_ = _index18
					(_162_v93).ArraySet1(false, (_index18).Int())
					_155_v88 = Companion_Default___.Fm7(_28_v0, _28_v0, globalState)
				} else {
					var _163_v94 _dafny.Map
					_ = _163_v94
					_163_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _54_v19)
					_163_v94 = (_163_v94).Update(_54_v19, !(p1) || (p1))
					r1 = _28_v0
					var _164_v95 _dafny.Array
					_ = _164_v95
					var _nw23 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
					_ = _nw23
					_164_v95 = _nw23
					var _165_v96 *C0
					_ = _165_v96
					var _nw24 *C0 = New_C0_()
					_ = _nw24
					_nw24.Ctor__(!(p1))
					_165_v96 = _nw24
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_164_v95), 0))
					_ = _index19
					(_164_v95).ArraySet1(_165_v96, (_index19).Int())
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_164_v95), 0))
					_ = _index20
					(_164_v95).ArraySet1(_165_v96, (_index20).Int())
					var _166_v97 *C0
					_ = _166_v97
					var _nw25 *C0 = New_C0_()
					_ = _nw25
					_nw25.Ctor__(p1)
					_166_v97 = _nw25
					var _167_v98 _dafny.Map
					_ = _167_v98
					_167_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v0, !(p1))
					var _168_v99 _dafny.Set
					_ = _168_v99
					_168_v99 = _dafny.SetOf(_29_v1, _29_v1)
					var _169_v100 D1
					_ = _169_v100
					_169_v100 = Companion_D1_.Create_DC2_(_168_v99)
					var _170_v101 _dafny.CodePoint
					_ = _170_v101
					_170_v101 = _dafny.CodePoint('w')
					var _171_v102 _dafny.Sequence
					_ = _171_v102
					_171_v102 = _dafny.SeqOf(_163_v94)
					var _172_v103 _dafny.MultiSet
					_ = _172_v103
					_172_v103 = _dafny.MultiSetOf(_170_v101, (func() _dafny.CodePoint {
						if (_166_v97).Fm6(_28_v0, _dafny.MultiSetFromSeq(_171_v102), _28_v0, globalState) {
							return _170_v101
						}
						return _170_v101
					})(), _170_v101, _170_v101, _170_v101)
					var _rhs18 bool = (_167_v98).Contains(_28_v0)
					_ = _rhs18
					var _rhs19 _dafny.Map = (func() _dafny.Map {
						var _coll12 = _dafny.NewMapBuilder()
						_ = _coll12
						for _iter12 := _dafny.Iterate((Companion_Default___.Fm3(globalState)).Elements()); ; {
							_compr_12, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _173_v104 _dafny.Int
							_173_v104 = interface{}(_compr_12).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm3(globalState), _173_v104) {
								_coll12.Add((_173_v104).Minus((_dafny.Zero).Minus(_28_v0)), false)
							}
						}
						return _coll12.ToMap()
					}()).Merge((_167_v98).Update(_dafny.IntOfInt64(556), p1))
					_ = _rhs19
					var _rhs20 D1 = _169_v100
					_ = _rhs20
					var _rhs21 _dafny.Array = _164_v95
					_ = _rhs21
					var _rhs22 _dafny.MultiSet = (_172_v103).Intersection(_172_v103)
					_ = _rhs22
					_54_v19 = _rhs18
					_167_v98 = _rhs19
					_169_v100 = _rhs20
					_164_v95 = _rhs21
					_172_v103 = _rhs22
				}
				var _174_v105 _dafny.Sequence
				_ = _174_v105
				_174_v105 = _dafny.UnicodeSeqOfUtf8Bytes("hr")
				_28_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(20)).Plus(_dafny.IntOfInt64(655)), Companion_Default___.SafeModuloInt(_28_v0, _dafny.IntOfUint32((_174_v105).Cardinality()))))
				_54_v19 = (p1) && ((func() bool {
					if _54_v19 {
						return _54_v19
					}
					return false
				})())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	r0 = _28_v0
	r1 = _28_v0
	r2 = _28_v0
	var _175_v106 _dafny.Set
	_ = _175_v106
	_175_v106 = _dafny.SetOf(true)
	var _176_v107 _dafny.Set
	_ = _176_v107
	_176_v107 = _dafny.SetOf(_175_v106, _175_v106, _175_v106)
	r3 = _176_v107
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _177_globalState *GlobalState
	_ = _177_globalState
	var _nw26 *GlobalState = New_GlobalState_()
	_ = _nw26
	_nw26.Ctor__()
	_177_globalState = _nw26
	var _178_v0 _dafny.Array
	_ = _178_v0
	var _nw27 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
	_ = _nw27
	_178_v0 = _nw27
	var _179_v1 _dafny.Int
	_ = _179_v1
	_179_v1 = _dafny.IntOfInt64(365)
	var _180_v2 bool
	_ = _180_v2
	_180_v2 = true
	var _181_v3 _dafny.Int
	_ = _181_v3
	var _182_v4 _dafny.Int
	_ = _182_v4
	var _183_v5 _dafny.Int
	_ = _183_v5
	var _184_v6 _dafny.Set
	_ = _184_v6
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	var _out2 _dafny.Int
	_ = _out2
	var _out3 _dafny.Set
	_ = _out3
	_out0, _out1, _out2, _out3 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _179_v1), _180_v2, _177_globalState)
	_181_v3 = _out0
	_182_v4 = _out1
	_183_v5 = _out2
	_184_v6 = _out3
	var _185_v7 _dafny.Sequence
	_ = _185_v7
	_185_v7 = _dafny.UnicodeSeqOfUtf8Bytes("gbehkj")
	_182_v4 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_182_v4), _dafny.IntOfUint32((_185_v7).Cardinality())), _dafny.IntOfInt64(813))
	var _186_v8 _dafny.CodePoint
	_ = _186_v8
	_186_v8 = _dafny.CodePoint('k')
	var _187_v9 _dafny.Map
	_ = _187_v9
	_187_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _179_v1)
	var _188_v10 _dafny.CodePoint
	_ = _188_v10
	_188_v10 = _dafny.CodePoint('u')
	var _189_v11 _dafny.Array
	_ = _189_v11
	var _nwElement0_1 _dafny.CodePoint = _186_v8
	_ = _nwElement0_1
	var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(20))
	_ = _nw28
	(_nw28).ArraySet1CodePoint(_nwElement0_1, 0)
	(_nw28).ArraySet1CodePoint(_186_v8, 1)
	(_nw28).ArraySet1CodePoint(_186_v8, 2)
	(_nw28).ArraySet1CodePoint(_186_v8, 3)
	(_nw28).ArraySet1CodePoint(Companion_Default___.Fm0(_185_v7, (_dafny.Zero).Minus(_183_v5), _183_v5, _177_globalState), 4)
	(_nw28).ArraySet1CodePoint(Companion_Default___.Fm0(_185_v7, _dafny.IntOfUint32((_dafny.SeqOf((_187_v9).Cardinality(), _183_v5, _179_v1)).Cardinality()), _179_v1, _177_globalState), 5)
	(_nw28).ArraySet1CodePoint(_186_v8, 6)
	(_nw28).ArraySet1CodePoint(_186_v8, 7)
	(_nw28).ArraySet1CodePoint(_186_v8, 8)
	(_nw28).ArraySet1CodePoint(_186_v8, 9)
	(_nw28).ArraySet1CodePoint(_dafny.CodePoint('c'), 10)
	(_nw28).ArraySet1CodePoint(_186_v8, 11)
	(_nw28).ArraySet1CodePoint((_188_v10), 12)
	(_nw28).ArraySet1CodePoint(Companion_Default___.Fm0(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(441))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}((func(_190_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_191_i0 _dafny.Int) _dafny.CodePoint {
			return _190_v8
		}
	})(_186_v8))), _183_v5, _181_v3, _177_globalState), 13)
	(_nw28).ArraySet1CodePoint(_186_v8, 14)
	(_nw28).ArraySet1CodePoint(_186_v8, 15)
	(_nw28).ArraySet1CodePoint((_185_v7).Select((Companion_Default___.SafeIndex(_181_v3, _dafny.IntOfUint32((_185_v7).Cardinality()))).Uint32()).(_dafny.CodePoint), 16)
	(_nw28).ArraySet1CodePoint(_186_v8, 17)
	(_nw28).ArraySet1CodePoint(Companion_Default___.Fm0(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(916))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_192_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})), _181_v3, _181_v3, _177_globalState), 18)
	(_nw28).ArraySet1CodePoint(_dafny.CodePoint('x'), 19)
	_189_v11 = _nw28
	_189_v11 = _189_v11
	var _193_v12 _dafny.Map
	_ = _193_v12
	_193_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v4, _182_v4)
	if ((func() _dafny.Int {
		if (_193_v12).Contains(_182_v4) {
			return (_193_v12).Get(_182_v4).(_dafny.Int)
		}
		return _182_v4
	})()).Cmp(_183_v5) != 0 {
		if true {
			var _194_v13 _dafny.Array
			_ = _194_v13
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw29
			_194_v13 = _nw29
			_194_v13 = _194_v13
			var _195_v14 _dafny.Sequence
			_ = _195_v14
			_195_v14 = _dafny.SeqOf(_182_v4, _182_v4)
			_187_v9 = (_187_v9).Update(_180_v2, _dafny.IntOfUint32((_195_v14).Cardinality()))
			var _196_v15 D1
			_ = _196_v15
			_196_v15 = Companion_D1_.Create_DC1_(_194_v13)
			_194_v13 = (func() _dafny.Array {
				if !(_180_v2) {
					return _194_v13
				}
				return (_196_v15).Dtor_cf1()
			})()
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_178_v0), 0))
			_ = _index21
			(_178_v0).ArraySet1(_180_v2, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_178_v0), 0))
			_ = _index22
			(_178_v0).ArraySet1(_180_v2, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_178_v0), 0))
			_ = _index23
			(_178_v0).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), (_index23).Int())
		} else {
			var _197_v16 _dafny.Map
			_ = _197_v16
			_197_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _182_v4)
			var _198_v17 _dafny.Int
			_ = _198_v17
			var _199_v18 _dafny.Int
			_ = _199_v18
			var _200_v19 _dafny.Int
			_ = _200_v19
			var _201_v20 _dafny.Set
			_ = _201_v20
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Set
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0((_197_v16).Merge(_197_v16), true, _177_globalState)
			_198_v17 = _out4
			_199_v18 = _out5
			_200_v19 = _out6
			_201_v20 = _out7
			_180_v2 = _180_v2
			var _202_v21 _dafny.Sequence
			_ = _202_v21
			_202_v21 = _dafny.SeqOf(_180_v2, _180_v2)
			var _203_v22 _dafny.Set
			_ = _203_v22
			_203_v22 = _dafny.SetOf(_dafny.IntOfUint32((_202_v21).Cardinality()))
			var _204_v23 _dafny.Int
			_ = _204_v23
			var _205_v24 _dafny.Int
			_ = _205_v24
			var _206_v25 _dafny.Int
			_ = _206_v25
			var _207_v26 _dafny.Set
			_ = _207_v26
			var _out8 _dafny.Int
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Set
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0(_197_v16, (_203_v22).Equals(_203_v22), _177_globalState)
			_204_v23 = _out8
			_205_v24 = _out9
			_206_v25 = _out10
			_207_v26 = _out11
			_204_v23 = Companion_Default___.SafeDivisionInt(_206_v25, _206_v25)
			var _208_v27 D1
			_ = _208_v27
			_208_v27 = Companion_D1_.Create_DC3_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_180_v2, _180_v2)), _180_v2, _180_v2)
			var _pat_let_tv1 = _180_v2
			_ = _pat_let_tv1
			_180_v2 = ((func(_pat_let2_0 D1) D1 {
				return func(_209_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let3_0 bool) D1 {
						return func(_210_dt__update_hcf5_h0 bool) D1 {
							return Companion_D1_.Create_DC3_((_209_dt__update__tmp_h0).Dtor_cf3(), (_209_dt__update__tmp_h0).Dtor_cf4(), _210_dt__update_hcf5_h0)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let2_0)
			}(_208_v27)).Dtor_cf4()) && (_180_v2)
		}
		var _211_v28 _dafny.MultiSet
		_ = _211_v28
		_211_v28 = _dafny.MultiSetOf(_183_v5, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer16 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_212_i2 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(_dafny.IntOfInt64(31))
		}))).Cardinality()), _182_v4, _dafny.IntOfInt64(-248))
		_180_v2 = (_211_v28).IsDisjointFrom((_211_v28).Union(_211_v28))
		var _213_v29 _dafny.MultiSet
		_ = _213_v29
		_213_v29 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(259))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_214_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_215_i3 _dafny.Int) _dafny.CodePoint {
				return _214_v8
			}
		})(_186_v8))))
		var _216_v30 _dafny.Sequence
		_ = _216_v30
		_216_v30 = _dafny.SeqOf(_180_v2)
		var _217_v31 _dafny.Int
		_ = _217_v31
		_217_v31 = _dafny.IntOfUint32((_216_v30).Cardinality())
		var _218_v32 _dafny.Set
		_ = _218_v32
		_218_v32 = _dafny.SetOf(_180_v2, _180_v2)
		var _219_v33 _dafny.Set
		_ = _219_v33
		_219_v33 = _dafny.SetOf(!(_180_v2), _180_v2, Companion_Default___.Fm1(_185_v7, (_217_v31), _218_v32, _177_globalState), !(!(_180_v2)))
		var _rhs23 bool = _180_v2
		_ = _rhs23
		var _rhs24 bool = !((_dafny.MultiSetOf(_185_v7)).IsSubsetOf((_213_v29).Intersection(_213_v29)))
		_ = _rhs24
		var _rhs25 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(232))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_220_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_221_i4 _dafny.Int) _dafny.CodePoint {
				return _220_v8
			}
		})(_186_v8)))
		_ = _rhs25
		var _rhs26 _dafny.CodePoint = _186_v8
		_ = _rhs26
		var _rhs27 _dafny.Int = (_218_v32).Cardinality()
		_ = _rhs27
		_180_v2 = _rhs23
		_180_v2 = _rhs24
		_185_v7 = _rhs25
		_186_v8 = _rhs26
		_181_v3 = _rhs27
		var _222_v34 _dafny.Map
		_ = _222_v34
		_222_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _182_v4)
		var _223_v35 D3
		_ = _223_v35
		_223_v35 = Companion_D3_.Create_DC5_(_222_v34)
		var _224_v36 _dafny.Int
		_ = _224_v36
		var _225_v37 _dafny.Int
		_ = _225_v37
		var _226_v38 _dafny.Int
		_ = _226_v38
		var _227_v39 _dafny.Set
		_ = _227_v39
		var _out12 _dafny.Int
		_ = _out12
		var _out13 _dafny.Int
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		var _out15 _dafny.Set
		_ = _out15
		_out12, _out13, _out14, _out15 = Companion_Default___.M0((_223_v35).Dtor_cf7(), !(true), _177_globalState)
		_224_v36 = _out12
		_225_v37 = _out13
		_226_v38 = _out14
		_227_v39 = _out15
		var _228_v40 _dafny.Array
		_ = _228_v40
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
		_ = _nw30
		_228_v40 = _nw30
		var _229_v41 _dafny.Sequence
		_ = _229_v41
		_229_v41 = _dafny.SeqOf(_185_v7, _185_v7, _185_v7)
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_228_v40), 0))
		_ = _index24
		(_228_v40).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_229_v41, _229_v41), (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_228_v40), 0))
		_ = _index25
		(_228_v40).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_185_v7, _185_v7), _dafny.SeqOf(_185_v7)), (_index25).Int())
	} else {
		_179_v1 = _179_v1
		_180_v2 = false
		_193_v12 = (_193_v12).Update(_181_v3, (_dafny.IntOfInt64(265)).Plus(Companion_Default___.Fm2(_180_v2, _177_globalState)))
		var _rhs28 bool = _180_v2
		_ = _rhs28
		var _rhs29 _dafny.CodePoint = _186_v8
		_ = _rhs29
		var _rhs30 _dafny.Int = _181_v3
		_ = _rhs30
		_180_v2 = _rhs28
		_186_v8 = _rhs29
		_182_v4 = _rhs30
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_178_v0), 0))
		_ = _index26
		(_178_v0).ArraySet1(_180_v2, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_178_v0), 0))
		_ = _index27
		(_178_v0).ArraySet1(((_182_v4).Cmp(_183_v5) < 0) == (_180_v2), (_index27).Int())
	}
	var _230_v42 _dafny.Set
	_ = _230_v42
	_230_v42 = _dafny.SetOf(_180_v2)
	if (func() bool {
		if true {
			return Companion_Default___.Fm1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("pkkborlrx"), (Companion_Default___.SafeIndex(_183_v5, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pkkborlrx")).Cardinality()))).Uint32(), _186_v8), Companion_Default___.Fm2(_180_v2, _177_globalState), _230_v42, _177_globalState)
		}
		return (_230_v42).Equals(_230_v42)
	})() {
		var _231_v43 _dafny.MultiSet
		_ = _231_v43
		_231_v43 = _dafny.MultiSetOf(_180_v2)
		var _232_v44 _dafny.Map
		_ = _232_v44
		_232_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, (_231_v43).Cardinality())
		var _233_v45 D3
		_ = _233_v45
		_233_v45 = Companion_D3_.Create_DC5_(_232_v44)
		var _234_v46 _dafny.Int
		_ = _234_v46
		var _235_v47 _dafny.Int
		_ = _235_v47
		var _236_v48 _dafny.Int
		_ = _236_v48
		var _237_v49 _dafny.Set
		_ = _237_v49
		var _out16 _dafny.Int
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		var _out18 _dafny.Int
		_ = _out18
		var _out19 _dafny.Set
		_ = _out19
		_out16, _out17, _out18, _out19 = Companion_Default___.M0((func() _dafny.Map {
			if _180_v2 {
				return _232_v44
			}
			return (_233_v45).Dtor_cf7()
		})(), _180_v2, _177_globalState)
		_234_v46 = _out16
		_235_v47 = _out17
		_236_v48 = _out18
		_237_v49 = _out19
		var _238_v50 _dafny.Int
		_ = _238_v50
		var _239_v51 _dafny.Int
		_ = _239_v51
		var _240_v52 _dafny.Int
		_ = _240_v52
		var _241_v53 _dafny.Set
		_ = _241_v53
		var _out20 _dafny.Int
		_ = _out20
		var _out21 _dafny.Int
		_ = _out21
		var _out22 _dafny.Int
		_ = _out22
		var _out23 _dafny.Set
		_ = _out23
		_out20, _out21, _out22, _out23 = Companion_Default___.M0(_232_v44, _180_v2, _177_globalState)
		_238_v50 = _out20
		_239_v51 = _out21
		_240_v52 = _out22
		_241_v53 = _out23
		var _242_v54 _dafny.Int
		_ = _242_v54
		var _243_v55 _dafny.Int
		_ = _243_v55
		var _244_v56 _dafny.Int
		_ = _244_v56
		var _245_v57 _dafny.Set
		_ = _245_v57
		var _out24 _dafny.Int
		_ = _out24
		var _out25 _dafny.Int
		_ = _out25
		var _out26 _dafny.Int
		_ = _out26
		var _out27 _dafny.Set
		_ = _out27
		_out24, _out25, _out26, _out27 = Companion_Default___.M0((_232_v44).Merge(_232_v44), _180_v2, _177_globalState)
		_242_v54 = _out24
		_243_v55 = _out25
		_244_v56 = _out26
		_245_v57 = _out27
		var _246_v58 _dafny.Int
		_ = _246_v58
		var _247_v59 _dafny.Int
		_ = _247_v59
		var _248_v60 _dafny.Int
		_ = _248_v60
		var _249_v61 _dafny.Set
		_ = _249_v61
		var _out28 _dafny.Int
		_ = _out28
		var _out29 _dafny.Int
		_ = _out29
		var _out30 _dafny.Int
		_ = _out30
		var _out31 _dafny.Set
		_ = _out31
		_out28, _out29, _out30, _out31 = Companion_Default___.M0((_232_v44).Merge(_232_v44), _180_v2, _177_globalState)
		_246_v58 = _out28
		_247_v59 = _out29
		_248_v60 = _out30
		_249_v61 = _out31
		var _rhs31 bool = !((_180_v2) || (_180_v2))
		_ = _rhs31
		var _rhs32 bool = _180_v2
		_ = _rhs32
		var _rhs33 bool = _180_v2
		_ = _rhs33
		var _rhs34 _dafny.Array = (func() _dafny.Array {
			if _180_v2 {
				return _189_v11
			}
			return _189_v11
		})()
		_ = _rhs34
		var _rhs35 bool = _180_v2
		_ = _rhs35
		_180_v2 = _rhs31
		_180_v2 = _rhs32
		_180_v2 = _rhs33
		_189_v11 = _rhs34
		_180_v2 = _rhs35
	} else {
		var _250_v62 _dafny.Map
		_ = _250_v62
		_250_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _179_v1)
		var _251_v63 _dafny.Int
		_ = _251_v63
		var _252_v64 _dafny.Int
		_ = _252_v64
		var _253_v65 _dafny.Int
		_ = _253_v65
		var _254_v66 _dafny.Set
		_ = _254_v66
		var _out32 _dafny.Int
		_ = _out32
		var _out33 _dafny.Int
		_ = _out33
		var _out34 _dafny.Int
		_ = _out34
		var _out35 _dafny.Set
		_ = _out35
		_out32, _out33, _out34, _out35 = Companion_Default___.M0(_250_v62, _180_v2, _177_globalState)
		_251_v63 = _out32
		_252_v64 = _out33
		_253_v65 = _out34
		_254_v66 = _out35
		_253_v65 = (_181_v3).Times((_187_v9).Cardinality())
		_180_v2 = !(!_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm3(_177_globalState), (func() _dafny.Int {
			if _180_v2 {
				return _183_v5
			}
			return (func() _dafny.Int {
				if (_187_v9).Contains(_180_v2) {
					return (_187_v9).Get(_180_v2).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.SeqOf(_180_v2, _180_v2)).Cardinality())
			})()
		})()))
		var _255_v67 _dafny.Sequence
		_ = _255_v67
		_255_v67 = _dafny.SeqOf(_185_v7, _185_v7, _185_v7)
		var _256_v68 *C0
		_ = _256_v68
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__(Companion_Default___.Fm1((_255_v67).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_181_v3)).Cardinality()), _dafny.IntOfUint32((_255_v67).Cardinality()))).Uint32()).(_dafny.Sequence), _179_v1, _230_v42, _177_globalState))
		_256_v68 = _nw31
		var _257_v69 _dafny.Array
		_ = _257_v69
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
		_ = _nw32
		_257_v69 = _nw32
		var _258_v70 _dafny.Sequence
		_ = _258_v70
		_258_v70 = _dafny.SeqOf(Companion_Default___.Fm2(_180_v2, _177_globalState))
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_257_v69), 0))
		_ = _index28
		(_257_v69).ArraySet1(_258_v70, (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_257_v69), 0))
		_ = _index29
		(_257_v69).ArraySet1(_258_v70, (_index29).Int())
	}
	var _259_v71 *C0
	_ = _259_v71
	var _nw33 *C0 = New_C0_()
	_ = _nw33
	_nw33.Ctor__(_180_v2)
	_259_v71 = _nw33
	var _260_v72 _dafny.Sequence
	_ = _260_v72
	_260_v72 = _dafny.SeqOf(true)
	var _hi0 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
		if false {
			return _dafny.IntOfUint32((_dafny.SeqOf(_259_v71)).Cardinality())
		}
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_260_v72).Cardinality()))
	})())
	_ = _hi0
	for _261_i5 := _dafny.IntOfInt64(796); _261_i5.Cmp(_hi0) < 0; _261_i5 = _261_i5.Plus(_dafny.One) {
		var _262_v73 _dafny.Map
		_ = _262_v73
		_262_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _182_v4)
		var _263_v74 _dafny.Int
		_ = _263_v74
		var _264_v75 _dafny.Int
		_ = _264_v75
		var _265_v76 _dafny.Int
		_ = _265_v76
		var _266_v77 _dafny.Set
		_ = _266_v77
		var _out36 _dafny.Int
		_ = _out36
		var _out37 _dafny.Int
		_ = _out37
		var _out38 _dafny.Int
		_ = _out38
		var _out39 _dafny.Set
		_ = _out39
		_out36, _out37, _out38, _out39 = Companion_Default___.M0((_262_v73).Merge(_262_v73), Companion_Default___.Fm1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-795))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_267_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_268_i6 _dafny.Int) _dafny.CodePoint {
				return _267_v8
			}
		})(_186_v8))), _181_v3, _230_v42, _177_globalState), _177_globalState)
		_263_v74 = _out36
		_264_v75 = _out37
		_265_v76 = _out38
		_266_v77 = _out39
		var _269_v78 _dafny.Int
		_ = _269_v78
		var _270_v79 _dafny.Int
		_ = _270_v79
		var _271_v80 _dafny.Int
		_ = _271_v80
		var _272_v81 _dafny.Set
		_ = _272_v81
		var _out40 _dafny.Int
		_ = _out40
		var _out41 _dafny.Int
		_ = _out41
		var _out42 _dafny.Int
		_ = _out42
		var _out43 _dafny.Set
		_ = _out43
		_out40, _out41, _out42, _out43 = Companion_Default___.M0((func() _dafny.Map {
			if _180_v2 {
				return _262_v73
			}
			return _262_v73
		})(), (_230_v42).IsDisjointFrom(_dafny.SetOf(_180_v2, _180_v2, false, _180_v2, _180_v2)), _177_globalState)
		_269_v78 = _out40
		_270_v79 = _out41
		_271_v80 = _out42
		_272_v81 = _out43
		_180_v2 = !(false)
		var _273_v82 _dafny.Sequence
		_ = _273_v82
		_273_v82 = _dafny.SeqOf(_259_v71, _259_v71)
		var _274_v83 _dafny.Array
		_ = _274_v83
		var _nwElement0_2 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_273_v82, (Companion_Default___.SafeIndex(_270_v79, _dafny.IntOfUint32((_273_v82).Cardinality()))).Uint32(), _259_v71)
		_ = _nwElement0_2
		var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(3))
		_ = _nw34
		(_nw34).ArraySet1(_nwElement0_2, 0)
		(_nw34).ArraySet1(_273_v82, 1)
		(_nw34).ArraySet1(_273_v82, 2)
		_274_v83 = _nw34
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_274_v83), 0))
		_ = _index30
		(_274_v83).ArraySet1(_273_v82, (_index30).Int())
		var _275_v84 _dafny.Map
		_ = _275_v84
		_275_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _180_v2)
		var _276_v85 D1
		_ = _276_v85
		_276_v85 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_180_v2, (func() bool {
			if (_275_v84).Contains(_180_v2) {
				return (_275_v84).Get(_180_v2).(bool)
			}
			return _180_v2
		})()), true, _180_v2)
		var _277_v86 _dafny.Set
		_ = _277_v86
		_277_v86 = _dafny.SetOf(_dafny.IntOfInt64(38))
		var _278_v87 _dafny.Array
		_ = _278_v87
		var _nwElement0_3 _dafny.Int = _dafny.IntOfInt64(826)
		_ = _nwElement0_3
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(7))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_3, 0)
		(_nw35).ArraySet1(_181_v3, 1)
		(_nw35).ArraySet1((func() _dafny.Int {
			if (_187_v9).Contains((_276_v85).Dtor_cf4()) {
				return (_187_v9).Get((_276_v85).Dtor_cf4()).(_dafny.Int)
			}
			return _265_v76
		})(), 2)
		(_nw35).ArraySet1((_277_v86).Cardinality(), 3)
		(_nw35).ArraySet1(_181_v3, 4)
		(_nw35).ArraySet1(_265_v76, 5)
		(_nw35).ArraySet1(_271_v80, 6)
		_278_v87 = _nw35
		var _279_v88 _dafny.Map
		_ = _279_v88
		_279_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v87, _273_v82)
		var _280_v89 _dafny.Sequence
		_ = _280_v89
		_280_v89 = _dafny.SeqOf(_278_v87)
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_274_v83), 0))
		_ = _index31
		(_274_v83).ArraySet1((func() _dafny.Sequence {
			if (_279_v88).Contains((_280_v89).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_185_v7).Cardinality()), _dafny.IntOfUint32((_280_v89).Cardinality()))).Uint32()).(_dafny.Array)) {
				return (_279_v88).Get((_280_v89).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_185_v7).Cardinality()), _dafny.IntOfUint32((_280_v89).Cardinality()))).Uint32()).(_dafny.Array)).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_259_v71), _273_v82)
		})(), (_index31).Int())
	}
	var _281_v90 *C0
	_ = _281_v90
	var _nw36 *C0 = New_C0_()
	_ = _nw36
	_nw36.Ctor__((_dafny.IntOfUint32((_260_v72).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_260_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-989), _dafny.IntOfUint32((_260_v72).Cardinality()))).Uint32()).(bool)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_182_v4, _179_v1)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_260_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-989), _dafny.IntOfUint32((_260_v72).Cardinality()))).Uint32()).(bool))).Cardinality()))).Uint32(), _180_v2)).Cardinality())) < 0)
	_281_v90 = _nw36
	var _hi1 _dafny.Int = _183_v5
	_ = _hi1
	for _282_i7 := Companion_Default___.Fm2(_180_v2, _177_globalState); _282_i7.Cmp(_hi1) < 0; _282_i7 = _282_i7.Plus(_dafny.One) {
		var _283_v91 _dafny.Sequence
		_ = _283_v91
		_283_v91 = _dafny.SeqOf(_dafny.IntOfUint32((_185_v7).Cardinality()))
		var _284_v92 _dafny.Map
		_ = _284_v92
		_284_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _dafny.IntOfUint32((_283_v91).Cardinality()))
		var _285_v93 _dafny.Int
		_ = _285_v93
		var _286_v94 _dafny.Int
		_ = _286_v94
		var _287_v95 _dafny.Int
		_ = _287_v95
		var _288_v96 _dafny.Set
		_ = _288_v96
		var _out44 _dafny.Int
		_ = _out44
		var _out45 _dafny.Int
		_ = _out45
		var _out46 _dafny.Int
		_ = _out46
		var _out47 _dafny.Set
		_ = _out47
		_out44, _out45, _out46, _out47 = Companion_Default___.M0((_284_v92).Merge(_284_v92), _180_v2, _177_globalState)
		_285_v93 = _out44
		_286_v94 = _out45
		_287_v95 = _out46
		_288_v96 = _out47
		_180_v2 = _180_v2
		var _289_v97 _dafny.Array
		_ = _289_v97
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
		_ = _nw37
		_289_v97 = _nw37
		var _290_v98 _dafny.Sequence
		_ = _290_v98
		_290_v98 = _dafny.SeqOf(_259_v71, _281_v90)
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_289_v97), 0))
		_ = _index32
		(_289_v97).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_290_v98, _290_v98), (_index32).Int())
		var _291_v99 _dafny.Sequence
		_ = _291_v99
		_291_v99 = _dafny.SeqOf(_259_v71)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_289_v97), 0))
		_ = _index33
		(_289_v97).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_290_v98, (_291_v99)), (_index33).Int())
		_259_v71 = _259_v71
	}
	var _pat_let_tv2 = _183_v5
	_ = _pat_let_tv2
	if func(_source3 _dafny.CodePoint) bool {
		var _292___mcc_h0 _dafny.CodePoint = _source3
		_ = _292___mcc_h0
		var _293_cf0 _dafny.CodePoint = _292___mcc_h0
		_ = _293_cf0
		return (_pat_let_tv2).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(45), _293_cf0)).Cardinality()) >= 0
	}((func() _dafny.CodePoint {
		if _180_v2 {
			return _186_v8
		}
		return _186_v8
	})()) {
		var _294_v100 T0
		_ = _294_v100
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__(_180_v2)
		_294_v100 = _nw38
		var _295_v101 _dafny.Map
		_ = _295_v101
		_295_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _294_v100)
		var _296_v102 _dafny.Array
		_ = _296_v102
		var _nwElement0_4 T0 = _294_v100
		_ = _nwElement0_4
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(18))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_4, 0)
		(_nw39).ArraySet1(_294_v100, 1)
		(_nw39).ArraySet1(_294_v100, 2)
		(_nw39).ArraySet1(_294_v100, 3)
		(_nw39).ArraySet1(_294_v100, 4)
		(_nw39).ArraySet1(_294_v100, 5)
		(_nw39).ArraySet1(_294_v100, 6)
		(_nw39).ArraySet1(_294_v100, 7)
		(_nw39).ArraySet1(_294_v100, 8)
		(_nw39).ArraySet1(_294_v100, 9)
		(_nw39).ArraySet1(_294_v100, 10)
		(_nw39).ArraySet1(_294_v100, 11)
		(_nw39).ArraySet1(_294_v100, 12)
		(_nw39).ArraySet1((func() T0 {
			if true {
				return _294_v100
			}
			return _294_v100
		})(), 13)
		(_nw39).ArraySet1(_294_v100, 14)
		(_nw39).ArraySet1((func() T0 {
			if (_295_v101).Contains(!(_294_v100.F0())) {
				return (_295_v101).Get(!(_294_v100.F0())).(T0)
			}
			return _294_v100
		})(), 15)
		(_nw39).ArraySet1(_294_v100, 16)
		(_nw39).ArraySet1(_294_v100, 17)
		_296_v102 = _nw39
		var _rhs36 _dafny.Array = _296_v102
		_ = _rhs36
		var _rhs37 bool = (func() bool {
			if (func() bool {
				if _294_v100.F0() {
					return !(_294_v100.F0())
				}
				return _294_v100.F0()
			})() {
				return _294_v100.F0()
			}
			return _294_v100.F0()
		})()
		_ = _rhs37
		_296_v102 = _rhs36
		_180_v2 = _rhs37
		_182_v4 = _182_v4
		var _297_v103 _dafny.Sequence
		_ = _297_v103
		_297_v103 = _dafny.SeqOf(_185_v7, _185_v7, _185_v7, _dafny.UnicodeSeqOfUtf8Bytes("hucmr"))
		var _298_v104 _dafny.Sequence
		_ = _298_v104
		_298_v104 = _dafny.SeqOf(_dafny.IntOfUint32((_260_v72).Cardinality()), _dafny.IntOfInt64(479))
		var _299_v105 _dafny.Map
		_ = _299_v105
		_299_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_297_v103, _dafny.Companion_Sequence_.Update(_297_v103, (Companion_Default___.SafeIndex((_298_v104).Select((Companion_Default___.SafeIndex(_182_v4, _dafny.IntOfUint32((_298_v104).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_297_v103).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_300_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_301_i8 _dafny.Int) _dafny.CodePoint {
				return _300_v8
			}
		})(_186_v8))))), _186_v8)
		_299_v105 = (_299_v105).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm7(_181_v3, _182_v4, _177_globalState), _185_v7, _185_v7), _dafny.SeqOf(_185_v7)), _dafny.CodePoint('e'))
		var _302_v106 _dafny.Array
		_ = _302_v106
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_5
		var _nw40 _dafny.Array
		_ = _nw40
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw40 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = func(_303_i9 _dafny.Int) _dafny.Int {
				return (_303_i9).Times(_dafny.IntOfInt64(439))
			}
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw40 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw40).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw40).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_302_v106 = _nw40
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_302_v106), 0))
		_ = _index34
		(_302_v106).ArraySet1(_182_v4, (_index34).Int())
		var _304_v107 _dafny.Int
		_ = _304_v107
		_304_v107 = _181_v3
		var _305_v108 _dafny.Set
		_ = _305_v108
		_305_v108 = _dafny.SetOf((func() _dafny.Int {
			if _294_v100.F0() {
				return _179_v1
			}
			return _304_v107
		})(), _304_v107)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_302_v106), 0))
		_ = _index35
		(_302_v106).ArraySet1((_305_v108).Cardinality(), (_index35).Int())
		var _306_v109 _dafny.Sequence
		_ = _306_v109
		_306_v109 = _185_v7
		_185_v7 = ((func() _dafny.Sequence {
			if false {
				return _306_v109
			}
			return Companion_Default___.Fm8(_180_v2, !(_180_v2), _179_v1, _177_globalState)
		})())
	} else {
		var _307_v110 *C0
		_ = _307_v110
		var _nw41 *C0 = New_C0_()
		_ = _nw41
		_nw41.Ctor__((_182_v4).Cmp(_179_v1) > 0)
		_307_v110 = _nw41
		var _308_v111 D1
		_ = _308_v111
		_308_v111 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_180_v2, _180_v2), _180_v2, _180_v2)
		var _309_v112 D6
		_ = _309_v112
		_309_v112 = Companion_D6_.Create_DC10_(_281_v90)
		var _310_v113 _dafny.Sequence
		_ = _310_v113
		_310_v113 = _dafny.SeqOf(_259_v71, _281_v90, (func() *C0 {
			if _180_v2 {
				return _259_v71
			}
			return _259_v71
		})(), _281_v90, (func() *C0 {
			if (_308_v111).Dtor_cf4() {
				return _307_v110
			}
			return (_309_v112).Dtor_cf14()
		})())
		var _311_v114 _dafny.Array
		_ = _311_v114
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
		_ = _nw42
		_311_v114 = _nw42
		var _312_v115 _dafny.Map
		_ = _312_v115
		_312_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v1, _180_v2)
		var _313_v116 _dafny.Set
		_ = _313_v116
		_313_v116 = _dafny.SetOf(_312_v115)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_311_v114), 0))
		_ = _index36
		(_311_v114).ArraySet1(_313_v116, (_index36).Int())
		var _314_v117 T0
		_ = _314_v117
		var _nw43 *C0 = New_C0_()
		_ = _nw43
		_nw43.Ctor__(_180_v2)
		_314_v117 = _nw43
		var _315_v118 _dafny.Map
		_ = _315_v118
		_315_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v7, _314_v117)
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_311_v114), 0))
		_ = _index37
		var _rhs38 _dafny.Int = _dafny.IntOfInt64(721)
		_ = _rhs38
		var _rhs39 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_310_v113, _dafny.Companion_Sequence_.Concatenate(_310_v113, _310_v113)), (Companion_Default___.SafeIndex((_315_v118).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_310_v113, _dafny.Companion_Sequence_.Concatenate(_310_v113, _310_v113))).Cardinality()))).Uint32(), _307_v110)
		_ = _rhs39
		var _rhs40 _dafny.Set = Companion_Default___.Fm9(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_185_v7, _dafny.UnicodeSeqOfUtf8Bytes("icsjbibk"))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_260_v72, (Companion_Default___.SafeIndex(_182_v4, _dafny.IntOfUint32((_260_v72).Cardinality()))).Uint32(), _314_v117.F0())).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_314_v117.F0(), _dafny.IntOfInt64(-841))).Cardinality()).Plus((_314_v117).Fm5(_181_v3, true, _177_globalState)), _186_v8, _177_globalState)
		_ = _rhs40
		var _lhs10 _dafny.Array = _311_v114
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_311_v114), 0))
		_ = _lhs11
		_182_v4 = _rhs38
		_310_v113 = _rhs39
		(_lhs10).ArraySet1(_rhs40, (_lhs11).Int())
		var _316_v119 _dafny.MultiSet
		_ = _316_v119
		_316_v119 = _dafny.MultiSetOf(_180_v2)
		var _317_v120 _dafny.MultiSet
		_ = _317_v120
		_317_v120 = _dafny.MultiSetOf(_316_v119)
		_182_v4 = (_317_v120).Cardinality()
		_179_v1 = Companion_Default___.Fm2((_314_v117.F0()) && (!(false)), _177_globalState)
		if _180_v2 {
			var _318_v121 _dafny.Sequence
			_ = _318_v121
			_318_v121 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(-218), _182_v4, _dafny.IntOfInt64(-520), (Companion_D6_.Create_DC11_(_179_v1)).Dtor_cf15(), _179_v1))
			(_314_v117).F0_set_(_dafny.Companion_Sequence_.Equal(_318_v121, _318_v121))
			var _319_v122 _dafny.Sequence
			_ = _319_v122
			_319_v122 = _dafny.SeqOf(_181_v3, (_dafny.Zero).Minus(_181_v3))
			_319_v122 = _319_v122
			_186_v8 = _dafny.CodePoint('o')
			var _320_v123 *C0
			_ = _320_v123
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__(_180_v2)
			_320_v123 = _nw44
			_180_v2 = _180_v2
		} else {
			var _321_v124 *C0
			_ = _321_v124
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__(_314_v117.F0())
			_321_v124 = _nw45
			var _322_v125 *C0
			_ = _322_v125
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__(_180_v2)
			_322_v125 = _nw46
			var _323_v126 _dafny.Array
			_ = _323_v126
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw47
			_323_v126 = _nw47
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_323_v126), 0))
			_ = _index38
			(_323_v126).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(347))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_324_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_325_i10 _dafny.Int) _dafny.CodePoint {
					return _324_v8
				}
			})(_186_v8))), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_323_v126), 0))
			_ = _index39
			(_323_v126).ArraySet1(_185_v7, (_index39).Int())
			(_314_v117).F0_set_(_314_v117.F0())
			var _326_v127 *C0
			_ = _326_v127
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__(_180_v2)
			_326_v127 = _nw48
		}
	}
	var _327_v128 *C0
	_ = _327_v128
	var _nw49 *C0 = New_C0_()
	_ = _nw49
	_nw49.Ctor__(_180_v2)
	_327_v128 = _nw49
	if (_179_v1).Cmp((_182_v4).Minus(_183_v5)) < 0 {
		var _328_v129 _dafny.Map
		_ = _328_v129
		_328_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _180_v2)
		var _329_v130 _dafny.MultiSet
		_ = _329_v130
		_329_v130 = _dafny.MultiSetOf(_328_v129)
		_180_v2 = (_259_v71).Fm6((_dafny.Zero).Minus(_183_v5), (_329_v130).Union(_329_v130), (_dafny.Zero).Minus((_dafny.Zero).Minus(_182_v4)), _177_globalState)
		var _rhs41 _dafny.CodePoint = (_188_v10)
		_ = _rhs41
		var _rhs42 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-426))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_330_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_331_i11 _dafny.Int) _dafny.CodePoint {
				return _330_v8
			}
		})(_186_v8)))
		_ = _rhs42
		var _rhs43 _dafny.Int = _dafny.IntOfInt64(406)
		_ = _rhs43
		_186_v8 = _rhs41
		_185_v7 = _rhs42
		_182_v4 = _rhs43
		var _332_v131 _dafny.Map
		_ = _332_v131
		_332_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_185_v7, _181_v3, (_187_v9).Cardinality(), _177_globalState), _dafny.IntOfUint32((_185_v7).Cardinality()))
		var _333_v132 _dafny.Set
		_ = _333_v132
		_333_v132 = _dafny.SetOf(_332_v131)
		_333_v132 = _333_v132
		_180_v2 = _180_v2
		var _334_v133 _dafny.Map
		_ = _334_v133
		_334_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _dafny.IntOfInt64(940))
		var _335_v134 _dafny.MultiSet
		_ = _335_v134
		_335_v134 = _dafny.MultiSetOf(_180_v2, _180_v2)
		var _pat_let_tv3 = _335_v134
		_ = _pat_let_tv3
		var _336_v135 _dafny.Int
		_ = _336_v135
		var _337_v136 _dafny.Int
		_ = _337_v136
		var _338_v137 _dafny.Int
		_ = _338_v137
		var _339_v138 _dafny.Set
		_ = _339_v138
		var _out48 _dafny.Int
		_ = _out48
		var _out49 _dafny.Int
		_ = _out49
		var _out50 _dafny.Int
		_ = _out50
		var _out51 _dafny.Set
		_ = _out51
		_out48, _out49, _out50, _out51 = Companion_Default___.M0(_334_v133, (func(_pat_let4_0 D1) D1 {
			return func(_340_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let5_0 _dafny.MultiSet) D1 {
					return func(_341_dt__update_hcf3_h0 _dafny.MultiSet) D1 {
						return Companion_D1_.Create_DC3_(_341_dt__update_hcf3_h0, (_340_dt__update__tmp_h1).Dtor_cf4(), (_340_dt__update__tmp_h1).Dtor_cf5())
					}(_pat_let5_0)
				}(_pat_let_tv3)
			}(_pat_let4_0)
		}(Companion_D1_.Create_DC3_(_335_v134, !(_180_v2), _180_v2))).Dtor_cf5(), _177_globalState)
		_336_v135 = _out48
		_337_v136 = _out49
		_338_v137 = _out50
		_339_v138 = _out51
	} else {
		var _342_v139 _dafny.Array
		_ = _342_v139
		var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
		_ = _nw50
		_342_v139 = _nw50
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_6
		var _nw51 _dafny.Array
		_ = _nw51
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw51 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Map = (func(_343_v2 bool) func(_dafny.Int) _dafny.Map {
				return func(_344_i12 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v2, _343_v2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v2, _343_v2))
				}
			})(_180_v2)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw51 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw51).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw51).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_342_v139 = _nw51
		var _345_v140 _dafny.Sequence
		_ = _345_v140
		_345_v140 = _dafny.SeqOf(_230_v42)
		var _346_v141 _dafny.Map
		_ = _346_v141
		_346_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v5, _345_v140)
		_346_v141 = (_346_v141).Update(Companion_Default___.SafeModuloInt(_182_v4, _182_v4), _dafny.Companion_Sequence_.Concatenate(_345_v140, _345_v140))
		var _347_v142 _dafny.Map
		_ = _347_v142
		_347_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v3, _185_v7)
		var _348_v143 _dafny.Sequence
		_ = _348_v143
		_348_v143 = _dafny.SeqOf(_185_v7, _185_v7, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ffchstvvi"), _185_v7))
		var _rhs44 _dafny.Sequence = _260_v72
		_ = _rhs44
		var _rhs45 _dafny.Map = _347_v142
		_ = _rhs45
		var _rhs46 _dafny.Int = (_327_v128).Fm5(_183_v5, false, _177_globalState)
		_ = _rhs46
		var _rhs47 _dafny.Sequence = _185_v7
		_ = _rhs47
		var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-882))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_349_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_350_i13 _dafny.Int) _dafny.Sequence {
				return _349_v7
			}
		})(_185_v7))), _348_v143)
		_ = _rhs48
		_260_v72 = _rhs44
		_347_v142 = _rhs45
		_182_v4 = _rhs46
		_185_v7 = _rhs47
		_348_v143 = _rhs48
		_281_v90 = _281_v90
		var _351_v144 _dafny.Array
		_ = _351_v144
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_7
		var _nw52 _dafny.Array
		_ = _nw52
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw52 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Int = (func(_352_v12 _dafny.Map) func(_dafny.Int) _dafny.Int {
				return func(_353_i14 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_353_i14, (_352_v12).Cardinality())
				}
			})(_193_v12)
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
		_351_v144 = _nw52
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_351_v144), 0))
		_ = _index40
		(_351_v144).ArraySet1((func() _dafny.Int {
			if false {
				return (_dafny.Zero).Minus(_182_v4)
			}
			return _183_v5
		})(), (_index40).Int())
		var _354_v145 _dafny.MultiSet
		_ = _354_v145
		_354_v145 = _dafny.MultiSetOf(_185_v7, (_dafny.UnicodeSeqOfUtf8Bytes("stnwqxv")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(61))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_355_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_356_i15 _dafny.Int) _dafny.CodePoint {
				return _355_v8
			}
		})(_186_v8))))
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_351_v144), 0))
		_ = _index41
		var _rhs49 bool = (_354_v145).IsSubsetOf((_354_v145).Difference(_354_v145))
		_ = _rhs49
		var _rhs50 _dafny.Int = (_183_v5).Times(_181_v3)
		_ = _rhs50
		var _rhs51 _dafny.Int = (_281_v90).Fm5(_182_v4, !(_180_v2), _177_globalState)
		_ = _rhs51
		var _lhs12 _dafny.Array = _351_v144
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_351_v144), 0))
		_ = _lhs13
		_180_v2 = _rhs49
		(_lhs12).ArraySet1(_rhs50, (_lhs13).Int())
		_183_v5 = _rhs51
	}
	var _357_i16 _dafny.Int
	_ = _357_i16
	_357_i16 = _dafny.Zero
	{
		for (_260_v72).Select((Companion_Default___.SafeIndex(_179_v1, _dafny.IntOfUint32((_260_v72).Cardinality()))).Uint32()).(bool) {
			{
				if (_357_i16).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_357_i16 = (_357_i16).Plus(_dafny.One)
				_327_v128 = _281_v90
				var _358_v146 _dafny.Map
				_ = _358_v146
				_358_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_182_v4).Times((_327_v128).Fm5((_230_v42).Cardinality(), _180_v2, _177_globalState)), _180_v2)
				_358_v146 = _358_v146
				var _359_v147 _dafny.MultiSet
				_ = _359_v147
				_359_v147 = _dafny.MultiSetOf(_182_v4, _dafny.IntOfInt64(122))
				var _360_v148 _dafny.Sequence
				_ = _360_v148
				_360_v148 = _dafny.SeqOf(_dafny.IntOfInt64(176), _183_v5, _dafny.IntOfInt64(-170), Companion_Default___.Fm2(Companion_Default___.Fm1(_185_v7, _dafny.IntOfInt64(943), _dafny.SetOf(_180_v2), _177_globalState), _177_globalState))
				var _rhs52 bool = (_dafny.MultiSetFromSeq(_360_v148)).IsSubsetOf(_359_v147)
				_ = _rhs52
				var _rhs53 _dafny.CodePoint = _186_v8
				_ = _rhs53
				var _rhs54 _dafny.Int = (func() _dafny.Int {
					if (!(false)) || (_180_v2) {
						return Companion_Default___.SafeDivisionInt(_183_v5, _183_v5)
					}
					return _dafny.IntOfUint32((_185_v7).Cardinality())
				})()
				_ = _rhs54
				_180_v2 = _rhs52
				_186_v8 = _rhs53
				_182_v4 = _rhs54
				_183_v5 = _dafny.IntOfUint32((_185_v7).Cardinality())
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	_183_v5 = (_dafny.Zero).Minus(_179_v1)
	var _361_v149 _dafny.MultiSet
	_ = _361_v149
	_361_v149 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _180_v2))
	var _362_v150 _dafny.Map
	_ = _362_v150
	_362_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _180_v2)
	if (_327_v128).Fm6(_dafny.IntOfUint32((_185_v7).Cardinality()), (_361_v149).Union((_361_v149).Update(_362_v150, Companion_Default___.Abs(_182_v4))), _dafny.IntOfInt64(342), _177_globalState) {
		var _rhs55 _dafny.Array = _178_v0
		_ = _rhs55
		var _rhs56 _dafny.Int = (_dafny.Zero).Minus((_183_v5).Plus(_183_v5))
		_ = _rhs56
		var _rhs57 bool = _180_v2
		_ = _rhs57
		_178_v0 = _rhs55
		_181_v3 = _rhs56
		_180_v2 = _rhs57
		_185_v7 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(522))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_363_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_364_i17 _dafny.Int) _dafny.CodePoint {
				return _363_v8
			}
		})(_186_v8)))
		if _180_v2 {
			_178_v0 = _178_v0
			_181_v3 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-845), _182_v4)
			_185_v7 = _185_v7
			var _365_v151 *C0
			_ = _365_v151
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__(!(_180_v2))
			_365_v151 = _nw53
			var _366_v152 _dafny.Array
			_ = _366_v152
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_8
			var _nw54 _dafny.Array
			_ = _nw54
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw54 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.CodePoint = (func(_367_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_368_i18 _dafny.Int) _dafny.CodePoint {
						return _367_v10
					}
				})(_188_v10)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw54 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw54).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw54).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_366_v152 = _nw54
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_366_v152), 0))
			_ = _index42
			(_366_v152).ArraySet1(_188_v10, (_index42).Int())
			var _369_v153 _dafny.Sequence
			_ = _369_v153
			_369_v153 = _dafny.SeqOf(_183_v5)
			var _370_v154 _dafny.Sequence
			_ = _370_v154
			_370_v154 = _dafny.SeqOf(_369_v153)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_366_v152), 0))
			_ = _index43
			(_366_v152).ArraySet1(_186_v8, (_index43).Int())
		} else {
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))
			_ = _index44
			(_178_v0).ArraySet1(_180_v2, (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))
			_ = _index45
			(_178_v0).ArraySet1(_180_v2, (_index45).Int())
			var _371_v155 *C0
			_ = _371_v155
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__(!((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)) || (_180_v2))
			_371_v155 = _nw55
			var _372_v156 _dafny.Set
			_ = _372_v156
			_372_v156 = _dafny.SetOf(_186_v8)
			var _373_v157 _dafny.Map
			_ = _373_v157
			_373_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if _180_v2 {
					return (_260_v72).Select((Companion_Default___.SafeIndex(_179_v1, _dafny.IntOfUint32((_260_v72).Cardinality()))).Uint32()).(bool)
				}
				return _180_v2
			})(), (_dafny.SetOf(_186_v8, _186_v8, _186_v8)).Union(_372_v156))
			var _374_v158 _dafny.Set
			_ = _374_v158
			_374_v158 = _dafny.SetOf(_183_v5)
			var _375_v159 _dafny.MultiSet
			_ = _375_v159
			_375_v159 = _dafny.MultiSetOf((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool))
			var _376_v160 D1
			_ = _376_v160
			_376_v160 = Companion_D1_.Create_DC3_(_375_v159, _180_v2, _180_v2)
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))
			_ = _index46
			var _rhs58 _dafny.Int = Companion_Default___.SafeDivisionInt(_179_v1, (_dafny.IntOfInt64(-967)).Minus(_179_v1))
			_ = _rhs58
			var _rhs59 bool = (func() bool {
				if _180_v2 {
					return _180_v2
				}
				return (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)
			})()
			_ = _rhs59
			var _rhs60 _dafny.Int = Companion_Default___.SafeModuloInt(_182_v4, ((_374_v158).Cardinality()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v7, (_376_v160).Dtor_cf4())).Cardinality()))
			_ = _rhs60
			var _rhs61 _dafny.Map = (_373_v157).Merge((Companion_Default___.Fm11(_177_globalState)).Merge(_373_v157))
			_ = _rhs61
			var _lhs14 _dafny.Array = _178_v0
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))
			_ = _lhs15
			_182_v4 = _rhs58
			(_lhs14).ArraySet1(_rhs59, (_lhs15).Int())
			_183_v5 = _rhs60
			_373_v157 = _rhs61
			var _377_v161 _dafny.Array
			_ = _377_v161
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw56
			_377_v161 = _nw56
			var _378_v162 _dafny.Map
			_ = _378_v162
			_378_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v161, (_371_v155).Fm5(_183_v5, (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), _177_globalState))
			var _379_v163 _dafny.Map
			_ = _379_v163
			_379_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v128, _378_v162)
			var _380_v164 _dafny.Int
			_ = _380_v164
			var _381_v165 _dafny.Int
			_ = _381_v165
			var _382_v166 _dafny.Int
			_ = _382_v166
			var _383_v167 _dafny.Set
			_ = _383_v167
			var _out52 _dafny.Int
			_ = _out52
			var _out53 _dafny.Int
			_ = _out53
			var _out54 _dafny.Int
			_ = _out54
			var _out55 _dafny.Set
			_ = _out55
			_out52, _out53, _out54, _out55 = Companion_Default___.M0((func() _dafny.Map {
				if (_379_v163).Contains(_281_v90) {
					return (_379_v163).Get(_281_v90).(_dafny.Map)
				}
				return _378_v162
			})(), _180_v2, _177_globalState)
			_380_v164 = _out52
			_381_v165 = _out53
			_382_v166 = _out54
			_383_v167 = _out55
			var _384_v168 _dafny.Array
			_ = _384_v168
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_9
			var _nw57 _dafny.Array
			_ = _nw57
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw57 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = (func(_385_v72 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_386_i19 _dafny.Int) _dafny.Sequence {
						return _385_v72
					}
				})(_260_v72)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw57 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw57).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw57).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_384_v168 = _nw57
			var _387_v169 _dafny.MultiSet
			_ = _387_v169
			_387_v169 = _dafny.MultiSetOf(_327_v128)
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))
			_ = _index47
			var _rhs62 bool = ((_dafny.MultiSetOf(_327_v128, _327_v128)).Difference(_387_v169)).IsSubsetOf(_387_v169)
			_ = _rhs62
			var _rhs63 _dafny.Array = (func() _dafny.Array {
				if !((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)) {
					return _384_v168
				}
				return _384_v168
			})()
			_ = _rhs63
			var _rhs64 _dafny.Int = _182_v4
			_ = _rhs64
			var _lhs16 _dafny.Array = _178_v0
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_178_v0), 0))
			_ = _lhs17
			(_lhs16).ArraySet1(_rhs62, (_lhs17).Int())
			_384_v168 = _rhs63
			_181_v3 = _rhs64
		}
		if (_327_v128).Fm6((_230_v42).Cardinality(), _361_v149, _182_v4, _177_globalState) {
			_183_v5 = _dafny.IntOfInt64(826)
			var _388_v170 _dafny.MultiSet
			_ = _388_v170
			_388_v170 = _dafny.MultiSetOf(_182_v4)
			var _389_v171 _dafny.Map
			_ = _389_v171
			_389_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _178_v0)
			var _390_v172 _dafny.Array
			_ = _390_v172
			var _nwElement0_5 _dafny.Int = (_183_v5).Minus(_179_v1)
			_ = _nwElement0_5
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(18))
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_5, 0)
			(_nw58).ArraySet1((_193_v12).Cardinality(), 1)
			(_nw58).ArraySet1(((_dafny.MultiSetOf(_182_v4)).Update(_181_v3, Companion_Default___.Abs(_182_v4))).Cardinality(), 2)
			(_nw58).ArraySet1(_183_v5, 3)
			(_nw58).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_185_v7).Cardinality()), (_230_v42).Cardinality()), 4)
			(_nw58).ArraySet1(_182_v4, 5)
			(_nw58).ArraySet1((_dafny.Zero).Minus((_181_v3).Times(_183_v5)), 6)
			(_nw58).ArraySet1((func() _dafny.Int {
				if _180_v2 {
					return (_dafny.Zero).Minus((_388_v170).Cardinality())
				}
				return _183_v5
			})(), 7)
			(_nw58).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _178_v0)).Merge(_389_v171)).Cardinality(), 8)
			(_nw58).ArraySet1(_181_v3, 9)
			(_nw58).ArraySet1(_182_v4, 10)
			(_nw58).ArraySet1(_182_v4, 11)
			(_nw58).ArraySet1(_182_v4, 12)
			(_nw58).ArraySet1(_dafny.IntOfInt64(-876), 13)
			(_nw58).ArraySet1(_183_v5, 14)
			(_nw58).ArraySet1((_dafny.SetOf(_180_v2, _180_v2, _180_v2)).Cardinality(), 15)
			(_nw58).ArraySet1((_dafny.Zero).Minus((_281_v90).Fm5(_dafny.IntOfInt64(947), _180_v2, _177_globalState)), 16)
			(_nw58).ArraySet1(_182_v4, 17)
			_390_v172 = _nw58
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_390_v172), 0))
			_ = _index48
			(_390_v172).ArraySet1(_183_v5, (_index48).Int())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_390_v172), 0))
			_ = _index49
			(_390_v172).ArraySet1(_dafny.IntOfInt64(-73), (_index49).Int())
			var _391_v173 _dafny.Int
			_ = _391_v173
			var _392_v174 _dafny.Int
			_ = _392_v174
			var _393_v175 _dafny.Int
			_ = _393_v175
			var _394_v176 _dafny.Set
			_ = _394_v176
			var _out56 _dafny.Int
			_ = _out56
			var _out57 _dafny.Int
			_ = _out57
			var _out58 _dafny.Int
			_ = _out58
			var _out59 _dafny.Set
			_ = _out59
			_out56, _out57, _out58, _out59 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _183_v5), _180_v2, _177_globalState)
			_391_v173 = _out56
			_392_v174 = _out57
			_393_v175 = _out58
			_394_v176 = _out59
			var _395_v177 _dafny.MultiSet
			_ = _395_v177
			_395_v177 = _dafny.MultiSetOf(_180_v2)
			_181_v3 = (_392_v174).Times((_395_v177).Cardinality())
			_179_v1 = (_327_v128).Fm5(((_390_v172).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_390_v172), 0))).Int()).(_dafny.Int)).Times(_393_v175), false, _177_globalState)
		} else {
			_183_v5 = (_183_v5).Plus(_182_v4)
			var _396_v178 D3
			_ = _396_v178
			_396_v178 = Companion_D3_.Create_DC6_((_dafny.Zero).Minus(_dafny.IntOfUint32((_260_v72).Cardinality())), _182_v4, _180_v2)
			_180_v2 = (_396_v178).Dtor_cf10()
			var _397_v179 _dafny.Map
			_ = _397_v179
			_397_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _181_v3)
			var _398_v180 _dafny.Int
			_ = _398_v180
			var _399_v181 _dafny.Int
			_ = _399_v181
			var _400_v182 _dafny.Int
			_ = _400_v182
			var _401_v183 _dafny.Set
			_ = _401_v183
			var _out60 _dafny.Int
			_ = _out60
			var _out61 _dafny.Int
			_ = _out61
			var _out62 _dafny.Int
			_ = _out62
			var _out63 _dafny.Set
			_ = _out63
			_out60, _out61, _out62, _out63 = Companion_Default___.M0((_397_v179).Merge(_397_v179), (func() bool {
				if (_362_v150).Contains(_180_v2) {
					return (_362_v150).Get(_180_v2).(bool)
				}
				return _180_v2
			})(), _177_globalState)
			_398_v180 = _out60
			_399_v181 = _out61
			_400_v182 = _out62
			_401_v183 = _out63
			var _402_v185 _dafny.Sequence
			_ = _402_v185
			_402_v185 = _185_v7
			var _403_v186 _dafny.Map
			_ = _403_v186
			_403_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(417), _402_v185)
			var _404_v187 _dafny.Map
			_ = _404_v187
			_404_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_260_v72).Cardinality()), _186_v8)
			var _405_v188 _dafny.Sequence
			_ = _405_v188
			_405_v188 = _dafny.SeqOf(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_403_v186).Keys().Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _406_v184 _dafny.Int
					_406_v184 = interface{}(_compr_13).(_dafny.Int)
					if (_403_v186).Contains(_406_v184) {
						_coll13.Add((_406_v184).Times(_399_v181), _186_v8)
					}
				}
				return _coll13.ToMap()
			}(), (_404_v187).Merge(_404_v187), (_404_v187).Merge(_404_v187), _404_v187)
			var _407_v189 _dafny.Array
			_ = _407_v189
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw59
			_407_v189 = _nw59
			var _408_v190 _dafny.Array
			_ = _408_v190
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_10
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Int = (func(_409_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_410_i20 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_410_i20, _409_v1)
					}
				})(_179_v1)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw60 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw60).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw60).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_408_v190 = _nw60
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_407_v189), 0))
			_ = _index50
			(_407_v189).ArraySet1(_408_v190, (_index50).Int())
			var _411_v191 _dafny.Map
			_ = _411_v191
			_411_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v42, _180_v2)
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_407_v189), 0))
			_ = _index51
			var _rhs65 _dafny.Sequence = (func() _dafny.Sequence {
				if false {
					return _405_v188
				}
				return _405_v188
			})()
			_ = _rhs65
			var _rhs66 *C0 = _281_v90
			_ = _rhs66
			var _rhs67 _dafny.Array = _408_v190
			_ = _rhs67
			var _rhs68 _dafny.Int = (_399_v181).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_185_v7).Cardinality())))
			_ = _rhs68
			var _rhs69 bool = (_411_v191).Equals(_411_v191)
			_ = _rhs69
			var _lhs18 _dafny.Array = _407_v189
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_407_v189), 0))
			_ = _lhs19
			_405_v188 = _rhs65
			_259_v71 = _rhs66
			(_lhs18).ArraySet1(_rhs67, (_lhs19).Int())
			_400_v182 = _rhs68
			_180_v2 = _rhs69
			var _rhs70 _dafny.Int = ((_dafny.Zero).Minus(_181_v3)).Plus(_dafny.IntOfInt64(229))
			_ = _rhs70
			var _rhs71 _dafny.Int = (_396_v178).Dtor_cf8()
			_ = _rhs71
			var _rhs72 _dafny.CodePoint = _186_v8
			_ = _rhs72
			var _rhs73 _dafny.Array = _178_v0
			_ = _rhs73
			_399_v181 = _rhs70
			_398_v180 = _rhs71
			_186_v8 = _rhs72
			_178_v0 = _rhs73
		}
		var _412_v192 _dafny.Map
		_ = _412_v192
		_412_v192 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _dafny.IntOfUint32((_260_v72).Cardinality()))
		var _413_v193 _dafny.Int
		_ = _413_v193
		var _414_v194 _dafny.Int
		_ = _414_v194
		var _415_v195 _dafny.Int
		_ = _415_v195
		var _416_v196 _dafny.Set
		_ = _416_v196
		var _out64 _dafny.Int
		_ = _out64
		var _out65 _dafny.Int
		_ = _out65
		var _out66 _dafny.Int
		_ = _out66
		var _out67 _dafny.Set
		_ = _out67
		_out64, _out65, _out66, _out67 = Companion_Default___.M0(_412_v192, _180_v2, _177_globalState)
		_413_v193 = _out64
		_414_v194 = _out65
		_415_v195 = _out66
		_416_v196 = _out67
	} else {
		_180_v2 = _180_v2
		var _417_v197 _dafny.Array
		_ = _417_v197
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw61
		_417_v197 = _nw61
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_417_v197), 0))
		_ = _index52
		(_417_v197).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("odeh"), _185_v7)).Cardinality()), (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_417_v197), 0))
		_ = _index53
		(_417_v197).ArraySet1(_183_v5, (_index53).Int())
		var _418_v198 D7
		_ = _418_v198
		_418_v198 = Companion_D7_.Create_DC12_(Companion_Default___.Fm12(true, _186_v8, _180_v2, _179_v1, _177_globalState))
		var _419_v199 _dafny.Map
		_ = _419_v199
		_419_v199 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, ((_418_v198).Dtor_cf16()).Cardinality())
		var _420_v200 _dafny.Int
		_ = _420_v200
		var _421_v201 _dafny.Int
		_ = _421_v201
		var _422_v202 _dafny.Int
		_ = _422_v202
		var _423_v203 _dafny.Set
		_ = _423_v203
		var _out68 _dafny.Int
		_ = _out68
		var _out69 _dafny.Int
		_ = _out69
		var _out70 _dafny.Int
		_ = _out70
		var _out71 _dafny.Set
		_ = _out71
		_out68, _out69, _out70, _out71 = Companion_Default___.M0(_419_v199, _180_v2, _177_globalState)
		_420_v200 = _out68
		_421_v201 = _out69
		_422_v202 = _out70
		_423_v203 = _out71
		_180_v2 = !(_180_v2)
		var _424_v204 _dafny.Array
		_ = _424_v204
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw62
		_424_v204 = _nw62
		var _rhs74 _dafny.Int = _179_v1
		_ = _rhs74
		var _rhs75 _dafny.Array = _424_v204
		_ = _rhs75
		var _rhs76 bool = !(_180_v2) || (_180_v2)
		_ = _rhs76
		var _rhs77 bool = _180_v2
		_ = _rhs77
		_179_v1 = _rhs74
		_424_v204 = _rhs75
		_180_v2 = _rhs76
		_180_v2 = _rhs77
	}
	var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
	_ = _index54
	(_178_v0).ArraySet1(_180_v2, (_index54).Int())
	var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
	_ = _index55
	(_178_v0).ArraySet1(Companion_Default___.Fm1(_185_v7, _dafny.IntOfUint32((_185_v7).Cardinality()), _230_v42, _177_globalState), (_index55).Int())
	var _425_v205 _dafny.Sequence
	_ = _425_v205
	_425_v205 = _dafny.SeqOf(_dafny.IntOfUint32((_185_v7).Cardinality()), _dafny.IntOfUint32((_260_v72).Cardinality()))
	var _426_v206 _dafny.MultiSet
	_ = _426_v206
	_426_v206 = _dafny.MultiSetOf(_425_v205)
	var _427_v207 D7
	_ = _427_v207
	_427_v207 = Companion_D7_.Create_DC12_(_426_v206)
	var _428_v208 D7
	_ = _428_v208
	_428_v208 = Companion_D7_.Create_DC12_((_427_v207).Dtor_cf16())
	var _429_v209 D7
	_ = _429_v209
	_429_v209 = Companion_D7_.Create_DC14_(_428_v208)
	var _pat_let_tv4 = _428_v208
	_ = _pat_let_tv4
	var _source4 D7 = func(_pat_let6_0 D7) D7 {
		return func(_430_dt__update__tmp_h4 D7) D7 {
			return func(_pat_let7_0 D7) D7 {
				return func(_431_dt__update_hcf18_h0 D7) D7 {
					return Companion_D7_.Create_DC14_(_431_dt__update_hcf18_h0)
				}(_pat_let7_0)
			}(_pat_let_tv4)
		}(_pat_let6_0)
	}(_429_v209)
	_ = _source4
	if _source4.Is_DC13() {
		var _432___mcc_h1 bool = _source4.Get_().(D7_DC13).Cf17
		_ = _432___mcc_h1
		var _433_cf17 bool = _432___mcc_h1
		_ = _433_cf17
		_182_v4 = (_dafny.Zero).Minus(_181_v3)
		_433_cf17 = _433_cf17
		var _434_v210 _dafny.Map
		_ = _434_v210
		_434_v210 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_433_cf17, _259_v71)
		var _435_v211 _dafny.Array
		_ = _435_v211
		var _nwElement0_6 *C0 = _259_v71
		_ = _nwElement0_6
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(18))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_6, 0)
		(_nw63).ArraySet1(_281_v90, 1)
		(_nw63).ArraySet1(_281_v90, 2)
		(_nw63).ArraySet1(_281_v90, 3)
		(_nw63).ArraySet1(_327_v128, 4)
		(_nw63).ArraySet1(_259_v71, 5)
		(_nw63).ArraySet1(_259_v71, 6)
		(_nw63).ArraySet1(_327_v128, 7)
		(_nw63).ArraySet1(_259_v71, 8)
		(_nw63).ArraySet1(_259_v71, 9)
		(_nw63).ArraySet1(_327_v128, 10)
		(_nw63).ArraySet1((func() *C0 {
			if (_434_v210).Contains(_180_v2) {
				return (_434_v210).Get(_180_v2).(*C0)
			}
			return _327_v128
		})(), 11)
		(_nw63).ArraySet1(_327_v128, 12)
		(_nw63).ArraySet1(_327_v128, 13)
		(_nw63).ArraySet1(_281_v90, 14)
		(_nw63).ArraySet1(_281_v90, 15)
		(_nw63).ArraySet1(_281_v90, 16)
		(_nw63).ArraySet1(_259_v71, 17)
		_435_v211 = _nw63
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_435_v211), 0))
		_ = _index56
		(_435_v211).ArraySet1(_281_v90, (_index56).Int())
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_435_v211), 0))
		_ = _index57
		(_435_v211).ArraySet1(_327_v128, (_index57).Int())
		_179_v1 = (Companion_Default___.SafeDivisionInt(_181_v3, _183_v5)).Times((_dafny.IntOfUint32((_260_v72).Cardinality())).Minus(_dafny.IntOfInt64(283)))
	} else if _source4.Is_DC12() {
		var _436___mcc_h2 _dafny.MultiSet = _source4.Get_().(D7_DC12).Cf16
		_ = _436___mcc_h2
		var _437_cf16 _dafny.MultiSet = _436___mcc_h2
		_ = _437_cf16
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
		_ = _index58
		var _rhs78 _dafny.Int = _183_v5
		_ = _rhs78
		var _rhs79 _dafny.Int = _182_v4
		_ = _rhs79
		var _rhs80 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xya"), _dafny.UnicodeSeqOfUtf8Bytes("kbshdbhf")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-774))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_438_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_439_i21 _dafny.Int) _dafny.CodePoint {
				return _438_v8
			}
		})(_186_v8))))
		_ = _rhs80
		var _rhs81 _dafny.CodePoint = _186_v8
		_ = _rhs81
		var _rhs82 bool = _180_v2
		_ = _rhs82
		var _lhs20 _dafny.Array = _178_v0
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
		_ = _lhs21
		_182_v4 = _rhs78
		_181_v3 = _rhs79
		_180_v2 = _rhs80
		_186_v8 = _rhs81
		(_lhs20).ArraySet1(_rhs82, (_lhs21).Int())
		var _440_v212 _dafny.Set
		_ = _440_v212
		_440_v212 = _dafny.SetOf(_181_v3)
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
		_ = _index59
		(_178_v0).ArraySet1((_259_v71).Fm6(Companion_Default___.SafeDivisionInt(_183_v5, _dafny.IntOfUint32((_185_v7).Cardinality())), _361_v149, (_dafny.Zero).Minus((_183_v5).Minus((_440_v212).Cardinality())), _177_globalState), (_index59).Int())
		var _441_v213 _dafny.MultiSet
		_ = _441_v213
		_441_v213 = _dafny.MultiSetOf(_180_v2)
		var _442_v214 D1
		_ = _442_v214
		_442_v214 = Companion_D1_.Create_DC3_(_441_v213, _180_v2, (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool))
		var _source5 D1 = (func() D1 {
			if true {
				return _442_v214
			}
			return _442_v214
		})()
		_ = _source5
		if _source5.Is_DC2() {
			var _443___mcc_h4 _dafny.Set = _source5.Get_().(D1_DC2).Cf2
			_ = _443___mcc_h4
			var _444_cf2 _dafny.Set = _443___mcc_h4
			_ = _444_cf2
			_183_v5 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_182_v4))
			var _445_v215 _dafny.Int
			_ = _445_v215
			var _446_v216 _dafny.Int
			_ = _446_v216
			var _447_v217 _dafny.Int
			_ = _447_v217
			var _448_v218 _dafny.Set
			_ = _448_v218
			var _out72 _dafny.Int
			_ = _out72
			var _out73 _dafny.Int
			_ = _out73
			var _out74 _dafny.Int
			_ = _out74
			var _out75 _dafny.Set
			_ = _out75
			_out72, _out73, _out74, _out75 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _181_v3), _180_v2, _177_globalState)
			_445_v215 = _out72
			_446_v216 = _out73
			_447_v217 = _out74
			_448_v218 = _out75
			var _449_v219 _dafny.Map
			_ = _449_v219
			_449_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _dafny.IntOfUint32((_425_v205).Cardinality()))
			var _450_v220 _dafny.Int
			_ = _450_v220
			var _451_v221 _dafny.Int
			_ = _451_v221
			var _452_v222 _dafny.Int
			_ = _452_v222
			var _453_v223 _dafny.Set
			_ = _453_v223
			var _out76 _dafny.Int
			_ = _out76
			var _out77 _dafny.Int
			_ = _out77
			var _out78 _dafny.Int
			_ = _out78
			var _out79 _dafny.Set
			_ = _out79
			_out76, _out77, _out78, _out79 = Companion_Default___.M0((_449_v219).Merge(_449_v219), _180_v2, _177_globalState)
			_450_v220 = _out76
			_451_v221 = _out77
			_452_v222 = _out78
			_453_v223 = _out79
			var _454_v224 D3
			_ = _454_v224
			_454_v224 = Companion_D3_.Create_DC6_(Companion_Default___.SafeDivisionInt((_441_v213).Cardinality(), _183_v5), (_dafny.IntOfInt64(225)).Minus((_dafny.Zero).Minus(_182_v4)), (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool))
			var _pat_let_tv5 = _185_v7
			_ = _pat_let_tv5
			var _pat_let_tv6 = _185_v7
			_ = _pat_let_tv6
			_454_v224 = func(_pat_let8_0 D3) D3 {
				return func(_455_dt__update__tmp_h5 D3) D3 {
					return func(_pat_let9_0 _dafny.Int) D3 {
						return func(_456_dt__update_hcf9_h0 _dafny.Int) D3 {
							return Companion_D3_.Create_DC6_((_455_dt__update__tmp_h5).Dtor_cf8(), _456_dt__update_hcf9_h0, (_455_dt__update__tmp_h5).Dtor_cf10())
						}(_pat_let9_0)
					}(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv5, _pat_let_tv6)).Cardinality()))
				}(_pat_let8_0)
			}(_454_v224)
		} else if _source5.Is_DC3() {
			var _457___mcc_h5 _dafny.MultiSet = _source5.Get_().(D1_DC3).Cf3
			_ = _457___mcc_h5
			var _458___mcc_h6 bool = _source5.Get_().(D1_DC3).Cf4
			_ = _458___mcc_h6
			var _459___mcc_h7 bool = _source5.Get_().(D1_DC3).Cf5
			_ = _459___mcc_h7
			var _460_cf5 bool = _459___mcc_h7
			_ = _460_cf5
			var _461_cf4 bool = _458___mcc_h6
			_ = _461_cf4
			var _462_cf3 _dafny.MultiSet = _457___mcc_h5
			_ = _462_cf3
			_260_v72 = _dafny.Companion_Sequence_.Concatenate(_260_v72, _260_v72)
			var _463_v225 _dafny.Map
			_ = _463_v225
			_463_v225 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _dafny.IntOfInt64(691))
			var _464_v226 _dafny.Sequence
			_ = _464_v226
			_464_v226 = _dafny.SeqOf(_259_v71)
			var _465_v227 _dafny.Map
			_ = _465_v227
			_465_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v226, (_259_v71).Fm6(_183_v5, _361_v149, _179_v1, _177_globalState))
			var _466_v228 _dafny.Map
			_ = _466_v228
			_466_v228 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v227, false)
			var _467_v229 _dafny.Int
			_ = _467_v229
			var _468_v230 _dafny.Int
			_ = _468_v230
			var _469_v231 _dafny.Int
			_ = _469_v231
			var _470_v232 _dafny.Set
			_ = _470_v232
			var _out80 _dafny.Int
			_ = _out80
			var _out81 _dafny.Int
			_ = _out81
			var _out82 _dafny.Int
			_ = _out82
			var _out83 _dafny.Set
			_ = _out83
			_out80, _out81, _out82, _out83 = Companion_Default___.M0(_463_v225, (func() bool {
				if (_466_v228).Contains(_465_v227) {
					return (_466_v228).Get(_465_v227).(bool)
				}
				return (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)
			})(), _177_globalState)
			_467_v229 = _out80
			_468_v230 = _out81
			_469_v231 = _out82
			_470_v232 = _out83
			_185_v7 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7(_179_v1, _467_v229, _177_globalState), _185_v7)
			_182_v4 = (_dafny.Zero).Minus(_182_v4)
		} else {
			var _471___mcc_h8 _dafny.Array = _source5.Get_().(D1_DC1).Cf1
			_ = _471___mcc_h8
			var _472_cf1 _dafny.Array = _471___mcc_h8
			_ = _472_cf1
			var _473_v233 _dafny.Map
			_ = _473_v233
			_473_v233 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _dafny.IntOfUint32((_185_v7).Cardinality()))
			var _474_v234 _dafny.Map
			_ = _474_v234
			_474_v234 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v90, _182_v4)
			var _475_v235 _dafny.Int
			_ = _475_v235
			var _476_v236 _dafny.Int
			_ = _476_v236
			var _477_v237 _dafny.Int
			_ = _477_v237
			var _478_v238 _dafny.Set
			_ = _478_v238
			var _out84 _dafny.Int
			_ = _out84
			var _out85 _dafny.Int
			_ = _out85
			var _out86 _dafny.Int
			_ = _out86
			var _out87 _dafny.Set
			_ = _out87
			_out84, _out85, _out86, _out87 = Companion_Default___.M0((_473_v233).Merge((_473_v233).Update(_178_v0, _182_v4)), (_179_v1).Cmp((_474_v234).Cardinality()) < 0, _177_globalState)
			_475_v235 = _out84
			_476_v236 = _out85
			_477_v237 = _out86
			_478_v238 = _out87
			var _479_v239 _dafny.Int
			_ = _479_v239
			var _480_v240 _dafny.Int
			_ = _480_v240
			var _481_v241 _dafny.Int
			_ = _481_v241
			var _482_v242 _dafny.Set
			_ = _482_v242
			var _out88 _dafny.Int
			_ = _out88
			var _out89 _dafny.Int
			_ = _out89
			var _out90 _dafny.Int
			_ = _out90
			var _out91 _dafny.Set
			_ = _out91
			_out88, _out89, _out90, _out91 = Companion_Default___.M0(_473_v233, (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), _177_globalState)
			_479_v239 = _out88
			_480_v240 = _out89
			_481_v241 = _out90
			_482_v242 = _out91
			_230_v42 = Companion_Default___.Fm13(_dafny.Companion_Sequence_.IsPrefixOf(_185_v7, _185_v7), !((_441_v213).Contains(_180_v2)), false, _177_globalState)
			var _483_v243 D6
			_ = _483_v243
			_483_v243 = Companion_D6_.Create_DC11_(_481_v241)
			var _484_v244 _dafny.Set
			_ = _484_v244
			_484_v244 = _dafny.SetOf(_483_v243, Companion_D6_.Create_DC11_(_480_v240), _483_v243)
			var _485_v245 T0
			_ = _485_v245
			var _nw64 *C0 = New_C0_()
			_ = _nw64
			_nw64.Ctor__(_180_v2)
			_485_v245 = _nw64
			var _486_v246 _dafny.Sequence
			_ = _486_v246
			_486_v246 = _dafny.SeqOf(_485_v245, _485_v245)
			var _487_v247 _dafny.Array
			_ = _487_v247
			var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
			_ = _nw65
			_487_v247 = _nw65
			var _rhs83 _dafny.Set = _484_v244
			_ = _rhs83
			var _rhs84 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_485_v245, _485_v245), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_485_v245), _486_v246))
			_ = _rhs84
			var _rhs85 _dafny.Array = _487_v247
			_ = _rhs85
			_484_v244 = _rhs83
			_486_v246 = _rhs84
			_487_v247 = _rhs85
		}
		var _488_v248 _dafny.Map
		_ = _488_v248
		_488_v248 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v2, _178_v0)
		_488_v248 = (_488_v248).Update(_dafny.Companion_Sequence_.Contains(_185_v7, _dafny.CodePoint('m')), _178_v0)
	} else {
		var _489___mcc_h3 D7 = _source4.Get_().(D7_DC14).Cf18
		_ = _489___mcc_h3
		var _490_cf18 D7 = _489___mcc_h3
		_ = _490_cf18
		if (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool) {
			_182_v4 = _183_v5
			var _491_v249 _dafny.Set
			_ = _491_v249
			_491_v249 = _dafny.SetOf(_425_v205, _425_v205, _425_v205, _425_v205)
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
			_ = _index60
			(_178_v0).ArraySet1(!((_491_v249).IsSubsetOf(Companion_Default___.Fm14((_187_v9).Cardinality(), _177_globalState))), (_index60).Int())
			_186_v8 = _186_v8
			var _492_v250 _dafny.Map
			_ = _492_v250
			_492_v250 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_v0, _183_v5)
			var _493_v251 _dafny.Int
			_ = _493_v251
			var _494_v252 _dafny.Int
			_ = _494_v252
			var _495_v253 _dafny.Int
			_ = _495_v253
			var _496_v254 _dafny.Set
			_ = _496_v254
			var _out92 _dafny.Int
			_ = _out92
			var _out93 _dafny.Int
			_ = _out93
			var _out94 _dafny.Int
			_ = _out94
			var _out95 _dafny.Set
			_ = _out95
			_out92, _out93, _out94, _out95 = Companion_Default___.M0(_492_v250, (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), _177_globalState)
			_493_v251 = _out92
			_494_v252 = _out93
			_495_v253 = _out94
			_496_v254 = _out95
			var _497_v255 _dafny.Set
			_ = _497_v255
			_497_v255 = _dafny.SetOf((func() _dafny.Int {
				if (_187_v9).Contains(!(!(_180_v2))) {
					return (_187_v9).Get(!(!(_180_v2))).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_185_v7).Cardinality()))
			})())
			var _498_v256 _dafny.Map
			_ = _498_v256
			_498_v256 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)), (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), _180_v2), (_497_v255).Cardinality())
			var _499_v257 _dafny.Sequence
			_ = _499_v257
			_499_v257 = _dafny.SeqOf(_260_v72)
			_498_v256 = (_498_v256).Update(_dafny.Companion_Sequence_.Concatenate((_499_v257).Select((Companion_Default___.SafeIndex(_493_v251, _dafny.IntOfUint32((_499_v257).Cardinality()))).Uint32()).(_dafny.Sequence), _260_v72), _183_v5)
		} else {
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))
			_ = _index61
			(_178_v0).ArraySet1(_180_v2, (_index61).Int())
			var _500_v258 _dafny.MultiSet
			_ = _500_v258
			_500_v258 = _dafny.MultiSetOf(_185_v7, _185_v7)
			var _501_v259 _dafny.Map
			_ = _501_v259
			_501_v259 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_502_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_503_i22 _dafny.Int) _dafny.Int {
					return _502_v1
				}
			})(_179_v1))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(23))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_504_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_505_i23 _dafny.Int) _dafny.CodePoint {
					return _504_v8
				}
			})(_186_v8))))
			var _nwElement0_7 bool = (func() bool {
				if (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool) {
					return _180_v2
				}
				return (_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)
			})()
			_ = _nwElement0_7
			var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(19))
			_ = _nw66
			(_nw66).ArraySet1(_nwElement0_7, 0)
			(_nw66).ArraySet1(_180_v2, 1)
			(_nw66).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_260_v72, _dafny.SeqOf(_180_v2)), 2)
			(_nw66).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), 3)
			(_nw66).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), 4)
			(_nw66).ArraySet1(_180_v2, 5)
			(_nw66).ArraySet1((_260_v72).Select((Companion_Default___.SafeIndex((_500_v258).Cardinality(), _dafny.IntOfUint32((_260_v72).Cardinality()))).Uint32()).(bool), 6)
			(_nw66).ArraySet1((true) == ((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)), 7)
			(_nw66).ArraySet1(_180_v2, 8)
			(_nw66).ArraySet1(_180_v2, 9)
			(_nw66).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), 10)
			(_nw66).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), 11)
			(_nw66).ArraySet1(_180_v2, 12)
			(_nw66).ArraySet1(_180_v2, 13)
			(_nw66).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), 14)
			(_nw66).ArraySet1(Companion_Default___.Fm1((func() _dafny.Sequence {
				if (_501_v259).Contains(_425_v205) {
					return (_501_v259).Get(_425_v205).(_dafny.Sequence)
				}
				return _185_v7
			})(), _181_v3, _230_v42, _177_globalState), 15)
			(_nw66).ArraySet1((Companion_D7_.Create_DC13_((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool))).Dtor_cf17(), 16)
			(_nw66).ArraySet1((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool), 17)
			(_nw66).ArraySet1(!((_178_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_178_v0), 0))).Int()).(bool)), 18)
			_178_v0 = _nw66
			_180_v2 = !(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_185_v7, _185_v7)))
			_186_v8 = _186_v8
			var _506_v260 _dafny.Set
			_ = _506_v260
			_506_v260 = _dafny.SetOf(_dafny.IntOfInt64(324), _179_v1)
			var _507_v261 *C0
			_ = _507_v261
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__((_506_v260).IsDisjointFrom(_506_v260))
			_507_v261 = _nw67
		}
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_11
		var _nw68 _dafny.Array
		_ = _nw68
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw68 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = func(_508_i24 _dafny.Int) bool {
				return false
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw68 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw68).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw68).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_178_v0 = _nw68
		var _509_v262 _dafny.Array
		_ = _509_v262
		var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
		_ = _nw69
		_509_v262 = _nw69
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_509_v262), 0))
		_ = _index62
		(_509_v262).ArraySet1((_362_v150).Merge(_362_v150), (_index62).Int())
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_509_v262), 0))
		_ = _index63
		(_509_v262).ArraySet1((_362_v150).Merge(_362_v150), (_index63).Int())
		var _510_v263 _dafny.Map
		_ = _510_v263
		_510_v263 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_425_v205, _425_v205), _183_v5)
		_510_v263 = _510_v263
	}
	_dafny.Print((_178_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_179_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_180_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_181_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_183_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v6).Equals(_dafny.SetOf(_dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_185_v7, _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'), _dafny.CodePoint('k'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_186_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(365))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v10))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_193_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(265))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v42).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_260_v72, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_357_i16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_361_v149).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_362_v150).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_425_v205, _dafny.SeqOf(_dafny.IntOfInt64(426), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_426_v206).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(426), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_427_v207).Dtor_cf16()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(426), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_428_v208).Dtor_cf16()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(426), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_429_v209).Dtor_cf18()).Dtor_cf16()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(426), _dafny.One))))
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
	Cf0 _dafny.CodePoint
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.CodePoint) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D0) Dtor_cf0() _dafny.CodePoint {
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
	Cf2 _dafny.Set
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Set) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf3 _dafny.MultiSet
	Cf4 bool
	Cf5 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.MultiSet, Cf4 bool, Cf5 bool) D1 {
	return D1{D1_DC3{Cf3, Cf4, Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Array
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Array) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.EmptySet)
}

func (_this D1) Dtor_cf2() _dafny.Set {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.MultiSet {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf2.Equals(data2.Cf2)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3) && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5
		}
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

func (CompanionStruct_D2_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf6) + ")"
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
	Cf8  _dafny.Int
	Cf9  _dafny.Int
	Cf10 bool
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf8 _dafny.Int, Cf9 _dafny.Int, Cf10 bool) D3 {
	return D3{D3_DC6{Cf8, Cf9, Cf10}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC5 struct {
	Cf7 _dafny.Map
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf7 _dafny.Map) D3 {
	return D3{D3_DC5{Cf7}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

type D3_DC7 struct {
	Cf11 D3
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 D3) D3 {
	return D3{D3_DC7{Cf11}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D3) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf8
}

func (_this D3) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf9
}

func (_this D3) Dtor_cf10() bool {
	return _this.Get_().(D3_DC6).Cf10
}

func (_this D3) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D3_DC5).Cf7
}

func (_this D3) Dtor_cf11() D3 {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D4_DC8 struct {
	Cf12 _dafny.Sequence
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf12 _dafny.Sequence) D4 {
	return D4{D4_DC8{Cf12}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D4_DC8).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
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
	Cf13 _dafny.Sequence
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf13 _dafny.Sequence) D5 {
	return D5{D5_DC9{Cf13}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D5) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D5_DC9).Cf13
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC9:
		{
			return "D5.DC9" + "(" + data.Cf13.VerbatimString(true) + ")"
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

type D6_DC11 struct {
	Cf15 _dafny.Int
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf15 _dafny.Int) D6 {
	return D6{D6_DC11{Cf15}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

type D6_DC10 struct {
	Cf14 *C0
}

func (D6_DC10) isD6() {}

func (CompanionStruct_D6_) Create_DC10_(Cf14 *C0) D6 {
	return D6{D6_DC10{Cf14}}
}

func (_this D6) Is_DC10() bool {
	_, ok := _this.Get_().(D6_DC10)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC11_(_dafny.Zero)
}

func (_this D6) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf15
}

func (_this D6) Dtor_cf14() *C0 {
	return _this.Get_().(D6_DC10).Cf14
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D6_DC10:
		{
			return "D6.DC10" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D6_DC10:
		{
			data2, ok := other.Get_().(D6_DC10)
			return ok && data1.Cf14 == data2.Cf14
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

type D7_DC13 struct {
	Cf17 bool
}

func (D7_DC13) isD7() {}

func (CompanionStruct_D7_) Create_DC13_(Cf17 bool) D7 {
	return D7{D7_DC13{Cf17}}
}

func (_this D7) Is_DC13() bool {
	_, ok := _this.Get_().(D7_DC13)
	return ok
}

type D7_DC12 struct {
	Cf16 _dafny.MultiSet
}

func (D7_DC12) isD7() {}

func (CompanionStruct_D7_) Create_DC12_(Cf16 _dafny.MultiSet) D7 {
	return D7{D7_DC12{Cf16}}
}

func (_this D7) Is_DC12() bool {
	_, ok := _this.Get_().(D7_DC12)
	return ok
}

type D7_DC14 struct {
	Cf18 D7
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf18 D7) D7 {
	return D7{D7_DC14{Cf18}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC13_(false)
}

func (_this D7) Dtor_cf17() bool {
	return _this.Get_().(D7_DC13).Cf17
}

func (_this D7) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D7_DC12).Cf16
}

func (_this D7) Dtor_cf18() D7 {
	return _this.Get_().(D7_DC14).Cf18
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC13:
		{
			return "D7.DC13" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D7_DC12:
		{
			return "D7.DC12" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC13:
		{
			data2, ok := other.Get_().(D7_DC13)
			return ok && data1.Cf17 == data2.Cf17
		}
	case D7_DC12:
		{
			data2, ok := other.Get_().(D7_DC12)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 bool
	Cf21 _dafny.MultiSet
	Cf22 _dafny.Sequence
	Cf23 _dafny.Map
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf20 bool, Cf21 _dafny.MultiSet, Cf22 _dafny.Sequence, Cf23 _dafny.Map) D8 {
	return D8{D8_DC16{Cf20, Cf21, Cf22, Cf23}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

type D8_DC17 struct {
	Cf24 _dafny.Sequence
	Cf25 _dafny.Int
	Cf26 _dafny.Int
	Cf27 _dafny.Int
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf24 _dafny.Sequence, Cf25 _dafny.Int, Cf26 _dafny.Int, Cf27 _dafny.Int) D8 {
	return D8{D8_DC17{Cf24, Cf25, Cf26, Cf27}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC15 struct {
	Cf19 _dafny.Sequence
}

func (D8_DC15) isD8() {}

func (CompanionStruct_D8_) Create_DC15_(Cf19 _dafny.Sequence) D8 {
	return D8{D8_DC15{Cf19}}
}

func (_this D8) Is_DC15() bool {
	_, ok := _this.Get_().(D8_DC15)
	return ok
}

type D8_DC18 struct {
	Cf28 D8
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf28 D8) D8 {
	return D8{D8_DC18{Cf28}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC16_(false, _dafny.EmptyMultiSet, _dafny.EmptySeq, _dafny.EmptyMap)
}

func (_this D8) Dtor_cf20() bool {
	return _this.Get_().(D8_DC16).Cf20
}

func (_this D8) Dtor_cf21() _dafny.MultiSet {
	return _this.Get_().(D8_DC16).Cf21
}

func (_this D8) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D8_DC16).Cf22
}

func (_this D8) Dtor_cf23() _dafny.Map {
	return _this.Get_().(D8_DC16).Cf23
}

func (_this D8) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D8_DC17).Cf24
}

func (_this D8) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D8_DC17).Cf25
}

func (_this D8) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D8_DC17).Cf26
}

func (_this D8) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D8_DC17).Cf27
}

func (_this D8) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D8_DC15).Cf19
}

func (_this D8) Dtor_cf28() D8 {
	return _this.Get_().(D8_DC18).Cf28
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D8_DC17:
		{
			return "D8.DC17" + "(" + data.Cf24.VerbatimString(true) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D8_DC15:
		{
			return "D8.DC15" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Equals(data2.Cf21) && data1.Cf22.Equals(data2.Cf22) && data1.Cf23.Equals(data2.Cf23)
		}
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf24.Equals(data2.Cf24) && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D8_DC15:
		{
			data2, ok := other.Get_().(D8_DC15)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
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

type D9_DC20 struct {
	Cf30 _dafny.Array
	Cf31 bool
	Cf32 _dafny.Int
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf30 _dafny.Array, Cf31 bool, Cf32 _dafny.Int) D9 {
	return D9{D9_DC20{Cf30, Cf31, Cf32}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC19 struct {
	Cf29 _dafny.Map
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf29 _dafny.Map) D9 {
	return D9{D9_DC19{Cf29}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC20_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero)
}

func (_this D9) Dtor_cf30() _dafny.Array {
	return _this.Get_().(D9_DC20).Cf30
}

func (_this D9) Dtor_cf31() bool {
	return _this.Get_().(D9_DC20).Cf31
}

func (_this D9) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D9_DC20).Cf32
}

func (_this D9) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D9_DC19).Cf29
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
	Cf34 bool
	Cf35 _dafny.Int
	Cf36 bool
	Cf37 _dafny.Int
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf34 bool, Cf35 _dafny.Int, Cf36 bool, Cf37 _dafny.Int) D10 {
	return D10{D10_DC22{Cf34, Cf35, Cf36, Cf37}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC21 struct {
	Cf33 _dafny.Sequence
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf33 _dafny.Sequence) D10 {
	return D10{D10_DC21{Cf33}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

type D10_DC23 struct {
	Cf38 D10
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf38 D10) D10 {
	return D10{D10_DC23{Cf38}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_(false, _dafny.Zero, false, _dafny.Zero)
}

func (_this D10) Dtor_cf34() bool {
	return _this.Get_().(D10_DC22).Cf34
}

func (_this D10) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D10_DC22).Cf35
}

func (_this D10) Dtor_cf36() bool {
	return _this.Get_().(D10_DC22).Cf36
}

func (_this D10) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D10_DC22).Cf37
}

func (_this D10) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D10_DC21).Cf33
}

func (_this D10) Dtor_cf38() D10 {
	return _this.Get_().(D10_DC23).Cf38
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf38) + ")"
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
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf33.Equals(data2.Cf33)
		}
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf38.Equals(data2.Cf38)
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

// Definition of trait T0
type T0 interface {
	String() string
	F0() bool
	F0_set_(value bool)
	Fm4(globalState *GlobalState) _dafny.Int
	Fm5(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int
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
	_f0 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f0 = false
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

func (_this *C0) F0() bool {
	return _this._f0
}
func (_this *C0) F0_set_(value bool) {
	_this._f0 = value
}
func (_this *C0) Ctor__(f0 bool) {
	{
		(_this)._f0 = f0
	}
}
func (_this *C0) Fm4(globalState *GlobalState) _dafny.Int {
	{
		var _source6 D1 = Companion_D1_.Create_DC2_(_dafny.SetOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(840), _this.F0())).Cardinality())))
		_ = _source6
		if _source6.Is_DC2() {
			var _511___mcc_h0 _dafny.Set = _source6.Get_().(D1_DC2).Cf2
			_ = _511___mcc_h0
			var _512_cf2 _dafny.Set = _511___mcc_h0
			_ = _512_cf2
			return _dafny.IntOfInt64(-299)
		} else if _source6.Is_DC3() {
			var _513___mcc_h1 _dafny.MultiSet = _source6.Get_().(D1_DC3).Cf3
			_ = _513___mcc_h1
			var _514___mcc_h2 bool = _source6.Get_().(D1_DC3).Cf4
			_ = _514___mcc_h2
			var _515___mcc_h3 bool = _source6.Get_().(D1_DC3).Cf5
			_ = _515___mcc_h3
			var _516_cf5 bool = _515___mcc_h3
			_ = _516_cf5
			var _517_cf4 bool = _514___mcc_h2
			_ = _517_cf4
			var _518_cf3 _dafny.MultiSet = _513___mcc_h1
			_ = _518_cf3
			return (_dafny.SetOf(_517_cf4)).Cardinality()
		} else {
			var _519___mcc_h4 _dafny.Array = _source6.Get_().(D1_DC1).Cf1
			_ = _519___mcc_h4
			var _520_cf1 _dafny.Array = _519___mcc_h4
			_ = _520_cf1
			return (_dafny.Zero).Minus((func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(617), _dafny.IntOfInt64(586))); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _521_v0 _dafny.Int
					_521_v0 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(617)).Cmp(_521_v0) <= 0) && ((_521_v0).Cmp(_dafny.IntOfInt64(586)) < 0) {
						_coll14.Add((_521_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vo")).Cardinality())), _dafny.IntOfInt64(129))
					}
				}
				return _coll14.ToMap()
			}()).Cardinality())
		}
	}
}
func (_this *C0) Fm5(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_dafny.IntOfInt64(39)).Times((_dafny.IntOfInt64(452))))
	}
}
func (_this *C0) Fm6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(222))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_522_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-394)
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-784))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_523_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hpmca")).Cardinality())
		})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hs")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gjhdka")).Cardinality()), _dafny.IntOfInt64(-765)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(511))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_524_i2 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(240)
		})), _dafny.SeqOf(_dafny.IntOfInt64(815), _dafny.IntOfInt64(633), ((_dafny.SetOf(_dafny.IntOfInt64(-347))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), _dafny.IntOfInt64(390)))).IsSubsetOf(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(583))))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
