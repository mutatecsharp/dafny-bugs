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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Map, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(773), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(506))).Cardinality())).Cardinality()))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(773), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(506))).Cardinality())).Cardinality())), _0_v0) {
				_coll0.Add((_0_v0).Plus(_dafny.IntOfInt64(357)))
			}
		}
		return _coll0.ToSet()
	}()).Difference((func() _dafny.Set {
		if true {
			return _dafny.SetOf(_dafny.IntOfInt64(766))
		}
		return _dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(412))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(98))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Int {
	return (((_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D8_.Create_DC18_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khb")).Cardinality()))))).Union(_dafny.MultiSetOf(Companion_D8_.Create_DC18_(_dafny.IntOfInt64(-677))))).Union((_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D8_.Create_DC18_(_dafny.IntOfInt64(-488))))).Intersection(_dafny.MultiSetOf(Companion_D8_.Create_DC18_((_dafny.SetOf(_dafny.CodePoint('d'))).Cardinality()))))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) bool {
	if true {
		return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("bhudtnhvn"), _dafny.CodePoint('w'))
	} else {
		return (_dafny.IntOfInt64(-773)).Cmp(_dafny.IntOfInt64(-371)) != 0
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(4)
}
func (_static *CompanionStruct_Default___) Fm8(p0 D1, globalState *GlobalState) D3 {
	var _source0 D7 = Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("xtkckav")))
	_ = _source0
	if _source0.Is_DC16() {
		var _2___mcc_h0 bool = _source0.Get_().(D7_DC16).Cf26
		_ = _2___mcc_h0
		var _3___mcc_h1 _dafny.Map = _source0.Get_().(D7_DC16).Cf27
		_ = _3___mcc_h1
		var _4___mcc_h2 _dafny.CodePoint = _source0.Get_().(D7_DC16).Cf28
		_ = _4___mcc_h2
		var _5___mcc_h3 bool = _source0.Get_().(D7_DC16).Cf29
		_ = _5___mcc_h3
		var _6_cf29 bool = _5___mcc_h3
		_ = _6_cf29
		var _7_cf28 _dafny.CodePoint = _4___mcc_h2
		_ = _7_cf28
		var _8_cf27 _dafny.Map = _3___mcc_h1
		_ = _8_cf27
		var _9_cf26 bool = _2___mcc_h0
		_ = _9_cf26
		return Companion_D3_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(26), _dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_9_cf26))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wlkdfem")).Cardinality())), _dafny.IntOfInt64(610))))
	} else {
		var _10___mcc_h4 _dafny.Map = _source0.Get_().(D7_DC15).Cf25
		_ = _10___mcc_h4
		var _11_cf25 _dafny.Map = _10___mcc_h4
		_ = _11_cf25
		if false {
			return Companion_D3_.Create_DC9_(func() _dafny.Map {
				var _coll1 = _dafny.NewMapBuilder()
				_ = _coll1
				for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(974), _dafny.IntOfInt64(89))); ; {
					_compr_1, _ok1 := _iter1()
					if !_ok1 {
						break
					}
					var _12_v0 _dafny.Int
					_12_v0 = interface{}(_compr_1).(_dafny.Int)
					if ((_dafny.IntOfInt64(974)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(89)) < 0) {
						_coll1.Add((_12_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(409))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg1 _dafny.Int) interface{} {
								return coer1(arg1)
							}
						}(func(_13_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						}))).Cardinality()))), _dafny.SetOf(_dafny.IntOfInt64(74), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), _dafny.IntOfInt64(799), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()))).Cardinality())))
					}
				}
				return _coll1.ToMap()
			}())
		} else {
			return Companion_D3_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-232), _dafny.SetOf(_dafny.IntOfInt64(822))))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("njtrptc"), _dafny.UnicodeSeqOfUtf8Bytes("l"))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.CodePoint, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("yn"), (func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("uegdnut")
	})())
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	if (_dafny.MultiSetOf(false)).IsDisjointFrom(_dafny.MultiSetOf(false, true)) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
	}
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Map, globalState *GlobalState) bool {
	return (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-719), _dafny.IntOfInt64(-755))).Cmp(((func() _dafny.Set {
		if false {
			return _dafny.SetOf(!(false))
		}
		return _dafny.SetOf(true, !(false))
	})()).Cardinality()) == 0
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(122))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_15_i0 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(false)).Cardinality()
	})), _dafny.SeqOf((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m'))).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _16_v2 _dafny.CodePoint
					_16_v2 = interface{}(_compr_4).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')), _16_v2) {
						_coll4.Add(_16_v2, (Companion_D9_.Create_DC23_(true)).Dtor_cf38())
					}
				}
				return _coll4.ToMap()
			}()).Keys().Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _17_v1 _dafny.CodePoint
				_17_v1 = interface{}(_compr_3).(_dafny.CodePoint)
				if (func() _dafny.Map {
					var _coll5 = _dafny.NewMapBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m'))).Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _16_v2 _dafny.CodePoint
						_16_v2 = interface{}(_compr_5).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')), _16_v2) {
							_coll5.Add(_16_v2, (Companion_D9_.Create_DC23_(true)).Dtor_cf38())
						}
					}
					return _coll5.ToMap()
				}()).Contains(_17_v1) {
					_coll3.Add(_17_v1, (_dafny.SetOf(_dafny.IntOfInt64(-244), _dafny.IntOfInt64(369))).Cardinality())
				}
			}
			return _coll3.ToMap()
		}()).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _18_v0 _dafny.CodePoint
			_18_v0 = interface{}(_compr_2).(_dafny.CodePoint)
			if (func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((func() _dafny.Map {
					var _coll7 = _dafny.NewMapBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m'))).Elements()); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _16_v2 _dafny.CodePoint
						_16_v2 = interface{}(_compr_7).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')), _16_v2) {
							_coll7.Add(_16_v2, (Companion_D9_.Create_DC23_(true)).Dtor_cf38())
						}
					}
					return _coll7.ToMap()
				}()).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _17_v1 _dafny.CodePoint
					_17_v1 = interface{}(_compr_6).(_dafny.CodePoint)
					if (func() _dafny.Map {
						var _coll8 = _dafny.NewMapBuilder()
						_ = _coll8
						for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m'))).Elements()); ; {
							_compr_8, _ok8 := _iter8()
							if !_ok8 {
								break
							}
							var _16_v2 _dafny.CodePoint
							_16_v2 = interface{}(_compr_8).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('m')), _16_v2) {
								_coll8.Add(_16_v2, (Companion_D9_.Create_DC23_(true)).Dtor_cf38())
							}
						}
						return _coll8.ToMap()
					}()).Contains(_17_v1) {
						_coll6.Add(_17_v1, (_dafny.SetOf(_dafny.IntOfInt64(-244), _dafny.IntOfInt64(369))).Cardinality())
					}
				}
				return _coll6.ToMap()
			}()).Contains(_18_v0) {
				_coll2.Add(_18_v0, _dafny.IntOfInt64(866))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(567), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cpapaci")).Cardinality())))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((Companion_D0_.Create_DC2_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(739), _dafny.IntOfInt64(-766)), Companion_D1_.Create_DC7_(_dafny.CodePoint('q'), _dafny.UnicodeSeqOfUtf8Bytes("p")))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-981))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_19_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	})))).Dtor_cf6(), _dafny.UnicodeSeqOfUtf8Bytes("rvlldbjbx"))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-240))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()))).Cardinality()))).Union(func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(951), _dafny.IntOfInt64(-986))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _21_v0 _dafny.Int
				_21_v0 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(951)).Cmp(_21_v0) <= 0) && ((_21_v0).Cmp(_dafny.IntOfInt64(-986)) < 0) {
					_coll10.Add(Companion_Default___.SafeDivisionInt(_21_v0, _dafny.IntOfInt64(28)))
				}
			}
			return _coll10.ToSet()
		}()).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _22_v1 _dafny.Int
			_22_v1 = interface{}(_compr_9).(_dafny.Int)
			if (func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(951), _dafny.IntOfInt64(-986))); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _23_v0 _dafny.Int
					_23_v0 = interface{}(_compr_11).(_dafny.Int)
					if ((_dafny.IntOfInt64(951)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(-986)) < 0) {
						_coll11.Add(Companion_Default___.SafeDivisionInt(_23_v0, _dafny.IntOfInt64(28)))
					}
				}
				return _coll11.ToSet()
			}()).Contains(_22_v1) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_22_v1, _dafny.IntOfInt64(-818)))
			}
		}
		return _coll9.ToSet()
	}()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_24_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.UnicodeSeqOfUtf8Bytes("oiuswc")))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	return (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("sxdj"), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ypeb")).Cardinality()))).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())))).Cmp((func() _dafny.Int {
		if !(!(false)) {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cn")).Cardinality())
		}
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(800))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_25_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(708))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_26_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality())
		}))).Cardinality())
	})()) == 0
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC6_((_dafny.SetOf(false)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(Companion_D1_.Create_DC6_(_dafny.IntOfInt64(-922)))).Union(_dafny.SetOf(Companion_D1_.Create_DC6_(_dafny.IntOfInt64(-649)), Companion_D1_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(922))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality())), Companion_D1_.Create_DC6_((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(656), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qjdg")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())))).Cardinality())))).Intersection(func() _dafny.Set {
		var _coll12 = _dafny.NewBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf(Companion_D1_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), false)).Cardinality()))).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _28_v0 D1
				_28_v0 = interface{}(_compr_13).(D1)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), false)).Cardinality())), _28_v0) {
					_coll13.Add(_28_v0)
				}
			}
			return _coll13.ToSet()
		}()).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _29_v1 D1
			_29_v1 = interface{}(_compr_12).(D1)
			if (func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.SeqOf(Companion_D1_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), false)).Cardinality()))).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _30_v0 D1
					_30_v0 = interface{}(_compr_14).(D1)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), false)).Cardinality())), _30_v0) {
						_coll14.Add(_30_v0)
					}
				}
				return _coll14.ToSet()
			}()).Contains(_29_v1) {
				_coll12.Add(_29_v1)
			}
		}
		return _coll12.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), true, true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(!(true), false)))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	if _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(false), _dafny.SeqOf(false, !(false))) {
		return (_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(false, true, true, false))
	} else {
		return (_dafny.MultiSetOf(true)).Union(_dafny.MultiSetOf(false, false))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(456), _dafny.IntOfInt64(-349))); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _31_v0 _dafny.Int
			_31_v0 = interface{}(_compr_15).(_dafny.Int)
			if ((_dafny.IntOfInt64(456)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(-349)) < 0) {
				_coll15.Add((_31_v0).Plus(_dafny.IntOfInt64(131)), Companion_D1_.Create_DC7_(_dafny.CodePoint('e'), _dafny.UnicodeSeqOfUtf8Bytes("gstwlmab")))
			}
		}
		return _coll15.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false, !(!(false))))).Cardinality(), Companion_D1_.Create_DC7_(_dafny.CodePoint('a'), _dafny.UnicodeSeqOfUtf8Bytes("leq"))))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-714)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-198))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_32_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("emwqfdps"))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Map {
	var _source1 D1 = Companion_D1_.Create_DC6_(_dafny.IntOfInt64(135))
	_ = _source1
	if _source1.Is_DC6() {
		var _33___mcc_h0 _dafny.Int = _source1.Get_().(D1_DC6).Cf12
		_ = _33___mcc_h0
		var _34_cf12 _dafny.Int = _33___mcc_h0
		_ = _34_cf12
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(271), true)
	} else if _source1.Is_DC7() {
		var _35___mcc_h1 _dafny.CodePoint = _source1.Get_().(D1_DC7).Cf13
		_ = _35___mcc_h1
		var _36___mcc_h2 _dafny.Sequence = _source1.Get_().(D1_DC7).Cf14
		_ = _36___mcc_h2
		var _37_cf14 _dafny.Sequence = _36___mcc_h2
		_ = _37_cf14
		var _38_cf13 _dafny.CodePoint = _35___mcc_h1
		_ = _38_cf13
		return (func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(556), _dafny.IntOfInt64(947))); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _39_v0 _dafny.Int
				_39_v0 = interface{}(_compr_16).(_dafny.Int)
				if ((_dafny.IntOfInt64(556)).Cmp(_39_v0) <= 0) && ((_39_v0).Cmp(_dafny.IntOfInt64(947)) < 0) {
					_coll16.Add((_39_v0).Plus(_dafny.IntOfInt64(-346)), true)
				}
			}
			return _coll16.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-892), true))
	} else {
		var _40___mcc_h3 _dafny.Sequence = _source1.Get_().(D1_DC5).Cf11
		_ = _40___mcc_h3
		var _41_cf11 _dafny.Sequence = _40___mcc_h3
		_ = _41_cf11
		return (func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-751), _dafny.IntOfInt64(772))); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _42_v1 _dafny.Int
				_42_v1 = interface{}(_compr_17).(_dafny.Int)
				if ((_dafny.IntOfInt64(-751)).Cmp(_42_v1) <= 0) && ((_42_v1).Cmp(_dafny.IntOfInt64(772)) < 0) {
					_coll17.Add((_42_v1).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("fwcp"))).Cardinality()), false)
				}
			}
			return _coll17.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kkb")).Cardinality()), false))
	}
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Intersection(_dafny.SetOf(!(!(true)), true, true, false))).Union((func() _dafny.Set {
		if !(true) {
			return _dafny.SetOf(true)
		}
		return _dafny.SetOf(true)
	})())
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("rxmfcp")
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(840))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_43_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kyandbjqq")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
	}))
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(-592)
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D0_.Create_DC3_(_dafny.IntOfInt64(330), true, false)).Dtor_cf8()), false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Difference(_dafny.SetOf(true))).Difference(_dafny.SetOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) _dafny.Sequence {
	var _source2 D11 = Companion_D11_.Create_DC28_()
	_ = _source2
	if _source2.Is_DC28() {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer12 func(_dafny.Int) D7) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_44_i0 _dafny.Int) D7 {
			return Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(995), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(89))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_45_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			}))))
		}))
	} else {
		var _46___mcc_h0 _dafny.Array = _source2.Get_().(D11_DC27).Cf42
		_ = _46___mcc_h0
		var _47_cf42 _dafny.Array = _46___mcc_h0
		_ = _47_cf42
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(560))).Uint32(), func(coer14 func(_dafny.Int) D7) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_48_i2 _dafny.Int) D7 {
			return Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(881), _dafny.UnicodeSeqOfUtf8Bytes("rqaiipv")))
		})), _dafny.SeqOf(Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("tnwkwu"))), Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-651))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_49_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_50_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})))), Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qdadesapq")).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("mdilqo"))), Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(false, false, true)).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("cmgqbmpj")))))
	}
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) bool {
	return !(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(22)), (_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32(((Companion_D14_.Create_DC31_(_dafny.SeqOf(true))).Dtor_cf45()).Cardinality())))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Sequence, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC28_()
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('q')
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, p1 D0, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
		var _coll18 = _dafny.NewBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)), false), false)).Keys().Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _51_v0 _dafny.Map
			_51_v0 = interface{}(_compr_18).(_dafny.Map)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)), false), false)).Contains(_51_v0) {
				_coll18.Add(_51_v0)
			}
		}
		return _coll18.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(-731))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(87))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_52_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()))).Merge((Companion_D15_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(923), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.IntOfInt64(496)))).Dtor_cf48())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('e'))).Cardinality()))), _dafny.IntOfInt64(417))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), _dafny.IntOfInt64(310))))
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.CodePoint('c'))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('p'))))).Union((func() _dafny.MultiSet {
		if !(!(true)) {
			return (Companion_D16_.Create_DC36_(_dafny.MultiSetOf(_dafny.CodePoint('d')))).Dtor_cf50()
		}
		return _dafny.MultiSetOf(_dafny.CodePoint('i'))
	})())
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yvjprk")).Cardinality())), _dafny.IntOfInt64(800))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-433), _dafny.IntOfInt64(576)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-714))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_53_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality())).Cardinality())).Cardinality())), (func() _dafny.Map {
		var _coll19 = _dafny.NewMapBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(270))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_54_i1 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-971))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_55_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()))
		})))).Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _56_v0 _dafny.Sequence
			_56_v0 = interface{}(_compr_19).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(270))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_54_i1 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-971))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_55_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))).Cardinality()))
			}))), _56_v0) {
				_coll19.Add(_56_v0, _dafny.IntOfInt64(460))
			}
		}
		return _coll19.ToMap()
	}()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.Map, p1 _dafny.MultiSet, p2 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(315), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))).Cardinality(), _dafny.IntOfInt64(228), _dafny.IntOfInt64(809), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ywclxcu")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 D8, globalState *GlobalState) D7 {
	if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-356)))).Cardinality()).Cmp(_dafny.IntOfInt64(90)) < 0 {
		return Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()))).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("khcfpyqr")))
	} else {
		return Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((Companion_D18_.Create_DC40_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(699)))).Dtor_cf57()).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("pqprsunx")))
	}
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_57_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(true, false, true, false)
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(251))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_58_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(false)
	}))), _dafny.SeqOf(_dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(!(false))))), _dafny.IntOfInt64(615))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-21)))
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source3 _dafny.Sequence = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)))
	_ = _source3
	var _59___mcc_h0 _dafny.Sequence = _source3
	_ = _59___mcc_h0
	var _60_cf19 _dafny.Sequence = _59___mcc_h0
	_ = _60_cf19
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('r'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('y')))
}
func (_static *CompanionStruct_Default___) Fm52(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)), _dafny.MultiSetOf(false, false), _dafny.MultiSetOf(true, true), _dafny.MultiSetOf(false))
		}
		return _dafny.SeqOf(_dafny.MultiSetOf(true, false))
	})()
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, p1 D0, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, true, false)), !(true))).Merge(func() _dafny.Map {
		var _coll20 = _dafny.NewMapBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, !(true)), (_dafny.MultiSetOf(_dafny.IntOfInt64(-591))).Cardinality())).Keys().Elements()); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _61_v0 _dafny.MultiSet
			_61_v0 = interface{}(_compr_20).(_dafny.MultiSet)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, !(true)), (_dafny.MultiSetOf(_dafny.IntOfInt64(-591))).Cardinality())).Contains(_61_v0) {
				_coll20.Add(_61_v0, false)
			}
		}
		return _coll20.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, false), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true)))
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source4 D7 = (func() D7 {
		if false {
			return Companion_D7_.Create_DC16_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(330), !(!(true))), _dafny.CodePoint('h'), true)
		}
		return Companion_D7_.Create_DC16_(false, func() _dafny.Map {
			var _coll21 = _dafny.NewMapBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate((func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(313), _dafny.IntOfInt64(-556))); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _62_v1 _dafny.Int
					_62_v1 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(313)).Cmp(_62_v1) <= 0) && ((_62_v1).Cmp(_dafny.IntOfInt64(-556)) < 0) {
						_coll22.Add((_62_v1).Times(_dafny.IntOfInt64(-215)), false)
					}
				}
				return _coll22.ToMap()
			}()).Keys().Elements()); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _63_v0 _dafny.Int
				_63_v0 = interface{}(_compr_21).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(313), _dafny.IntOfInt64(-556))); ; {
						_compr_23, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _62_v1 _dafny.Int
						_62_v1 = interface{}(_compr_23).(_dafny.Int)
						if ((_dafny.IntOfInt64(313)).Cmp(_62_v1) <= 0) && ((_62_v1).Cmp(_dafny.IntOfInt64(-556)) < 0) {
							_coll23.Add((_62_v1).Times(_dafny.IntOfInt64(-215)), false)
						}
					}
					return _coll23.ToMap()
				}()).Contains(_63_v0) {
					_coll21.Add(Companion_Default___.SafeModuloInt(_63_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-921))).Cardinality()), false)
				}
			}
			return _coll21.ToMap()
		}(), _dafny.CodePoint('r'), true)
	})()
	_ = _source4
	if _source4.Is_DC16() {
		var _64___mcc_h0 bool = _source4.Get_().(D7_DC16).Cf26
		_ = _64___mcc_h0
		var _65___mcc_h1 _dafny.Map = _source4.Get_().(D7_DC16).Cf27
		_ = _65___mcc_h1
		var _66___mcc_h2 _dafny.CodePoint = _source4.Get_().(D7_DC16).Cf28
		_ = _66___mcc_h2
		var _67___mcc_h3 bool = _source4.Get_().(D7_DC16).Cf29
		_ = _67___mcc_h3
		var _68_cf29 bool = _67___mcc_h3
		_ = _68_cf29
		var _69_cf28 _dafny.CodePoint = _66___mcc_h2
		_ = _69_cf28
		var _70_cf27 _dafny.Map = _65___mcc_h1
		_ = _70_cf27
		var _71_cf26 bool = _64___mcc_h0
		_ = _71_cf26
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-156)), Companion_D8_.Create_DC20_(Companion_D8_.Create_DC19_())), _dafny.SetOf(_69_cf28))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(259), _dafny.IntOfInt64(610))); ; {
				_compr_24, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _72_v2 _dafny.Int
				_72_v2 = interface{}(_compr_24).(_dafny.Int)
				if ((_dafny.IntOfInt64(259)).Cmp(_72_v2) <= 0) && ((_72_v2).Cmp(_dafny.IntOfInt64(610)) < 0) {
					_coll24.Add((_72_v2).Plus(_dafny.IntOfInt64(216)), Companion_D8_.Create_DC20_(Companion_D8_.Create_DC18_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-865), _71_cf26)).Cardinality())))
				}
			}
			return _coll24.ToMap()
		}(), _dafny.SetOf(_69_cf28)))
	} else {
		var _73___mcc_h4 _dafny.Map = _source4.Get_().(D7_DC15).Cf25
		_ = _73___mcc_h4
		var _74_cf25 _dafny.Map = _73___mcc_h4
		_ = _74_cf25
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(35), Companion_D8_.Create_DC20_(Companion_D8_.Create_DC20_(Companion_D8_.Create_DC19_()))), _dafny.SetOf(_dafny.CodePoint('d')))
	}
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, globalState *GlobalState) _dafny.Sequence {
	var _source5 D16 = Companion_D16_.Create_DC36_(_dafny.MultiSetOf(_dafny.CodePoint('j'), _dafny.CodePoint('m')))
	_ = _source5
	if _source5.Is_DC37() {
		var _75___mcc_h0 bool = _source5.Get_().(D16_DC37).Cf51
		_ = _75___mcc_h0
		var _76___mcc_h1 bool = _source5.Get_().(D16_DC37).Cf52
		_ = _76___mcc_h1
		var _77_cf52 bool = _76___mcc_h1
		_ = _77_cf52
		var _78_cf51 bool = _75___mcc_h0
		_ = _78_cf51
		return _dafny.SeqOf(_dafny.IntOfInt64(616))
	} else if _source5.Is_DC38() {
		var _79___mcc_h2 bool = _source5.Get_().(D16_DC38).Cf53
		_ = _79___mcc_h2
		var _80___mcc_h3 bool = _source5.Get_().(D16_DC38).Cf54
		_ = _80___mcc_h3
		var _81___mcc_h4 _dafny.Int = _source5.Get_().(D16_DC38).Cf55
		_ = _81___mcc_h4
		var _82_cf55 _dafny.Int = _81___mcc_h4
		_ = _82_cf55
		var _83_cf54 bool = _80___mcc_h3
		_ = _83_cf54
		var _84_cf53 bool = _79___mcc_h2
		_ = _84_cf53
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_82_cf55), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-104))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_85_cf55 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_86_i0 _dafny.Int) _dafny.Int {
				return _85_cf55
			}
		})(_82_cf55))))
	} else {
		var _87___mcc_h5 _dafny.MultiSet = _source5.Get_().(D16_DC36).Cf50
		_ = _87___mcc_h5
		var _88_cf50 _dafny.MultiSet = _87___mcc_h5
		_ = _88_cf50
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-5))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_89_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm56(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll25 = _dafny.NewMapBuilder()
		_ = _coll25
		for _iter25 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, !(false)), true))).Keys().Elements()); ; {
			_compr_25, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _90_v0 _dafny.Set
			_90_v0 = interface{}(_compr_25).(_dafny.Set)
			if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, !(false)), true))).Contains(_90_v0) {
				_coll25.Add(_90_v0, (Companion_D3_.Create_DC10_(true, false)).Dtor_cf17())
			}
		}
		return _coll25.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (_dafny.Map, _dafny.Int) {
	var r0 _dafny.Map = _dafny.EmptyMap
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var _91_v0 _dafny.Int
	_ = _91_v0
	_91_v0 = _dafny.IntOfInt64(805)
	var _92_v1 bool
	_ = _92_v1
	_92_v1 = false
	var _93_v2 _dafny.MultiSet
	_ = _93_v2
	_93_v2 = _dafny.MultiSetOf(_92_v1, _92_v1)
	var _94_v3 D10
	_ = _94_v3
	_94_v3 = Companion_D10_.Create_DC25_(_93_v2)
	var _95_v4 _dafny.Sequence
	_ = _95_v4
	_95_v4 = _dafny.SeqOf(_94_v3)
	var _rhs0 _dafny.Int = _91_v0
	_ = _rhs0
	var _rhs1 _dafny.Int = _91_v0
	_ = _rhs1
	var _rhs2 bool = (_91_v0).Cmp(_dafny.IntOfUint32((_95_v4).Cardinality())) == 0
	_ = _rhs2
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	r1 = _rhs0
	_lhs0.F4 = _rhs1
	_lhs1.F6 = _rhs2
	var _96_v5 *C4
	_ = _96_v5
	var _nw0 *C4 = New_C4_()
	_ = _nw0
	_nw0.Ctor__()
	_96_v5 = _nw0
	var _97_v6 _dafny.Map
	_ = _97_v6
	_97_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v1, true)
	var _98_v7 _dafny.MultiSet
	_ = _98_v7
	_98_v7 = _dafny.MultiSetOf(_91_v0)
	var _99_v8 _dafny.MultiSet
	_ = _99_v8
	_99_v8 = _dafny.MultiSetOf((_98_v7).Cardinality())
	var _100_v9 _dafny.Sequence
	_ = _100_v9
	_100_v9 = _dafny.SeqOf((_99_v8).Cardinality(), _dafny.IntOfInt64(32), _91_v0)
	var _101_v10 _dafny.Sequence
	_ = _101_v10
	_101_v10 = _dafny.UnicodeSeqOfUtf8Bytes("sxtwt")
	var _102_v11 _dafny.Array
	_ = _102_v11
	var _nwElement0_0 _dafny.Int = _dafny.IntOfUint32((_100_v9).Cardinality())
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(6))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1((func() _dafny.Int {
		if (_98_v7).Contains(_91_v0) {
			return (_98_v7).Multiplicity(_91_v0)
		}
		return _91_v0
	})(), 1)
	(_nw1).ArraySet1(_91_v0, 2)
	(_nw1).ArraySet1(_dafny.IntOfUint32((_101_v10).Cardinality()), 3)
	(_nw1).ArraySet1(_91_v0, 4)
	(_nw1).ArraySet1(_91_v0, 5)
	_102_v11 = _nw1
	var _103_v12 _dafny.Array
	_ = _103_v12
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_104_v1 bool) func(_dafny.Int) bool {
			return func(_105_i2 _dafny.Int) bool {
				return _104_v1
			}
		})(_92_v1)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw2 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw2).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw2).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_103_v12 = _nw2
	var _106_v13 _dafny.Array
	_ = _106_v13
	var _nwElement0_1 _dafny.Array = _103_v12
	_ = _nwElement0_1
	var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(9))
	_ = _nw3
	(_nw3).ArraySet1(_nwElement0_1, 0)
	(_nw3).ArraySet1(_103_v12, 1)
	(_nw3).ArraySet1(_103_v12, 2)
	(_nw3).ArraySet1(_103_v12, 3)
	(_nw3).ArraySet1(_103_v12, 4)
	(_nw3).ArraySet1(_103_v12, 5)
	(_nw3).ArraySet1(_103_v12, 6)
	(_nw3).ArraySet1(_103_v12, 7)
	(_nw3).ArraySet1(_103_v12, 8)
	_106_v13 = _nw3
	var _107_v14 D9
	_ = _107_v14
	_107_v14 = Companion_D9_.Create_DC22_(_102_v11, _92_v1, _106_v13, _101_v10)
	var _108_v15 _dafny.Sequence
	_ = _108_v15
	_108_v15 = _dafny.SeqOf(Companion_Default___.Fm16(Companion_Default___.Fm11((_107_v14).Dtor_cf35(), _92_v1, _92_v1, globalState), globalState), !(!(true)), _92_v1)
	var _109_v16 _dafny.Array
	_ = _109_v16
	var _nwElement0_2 bool = (Companion_Default___.Fm16(_97_v6, globalState)) || (_92_v1)
	_ = _nwElement0_2
	var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(18))
	_ = _nw4
	(_nw4).ArraySet1(_nwElement0_2, 0)
	(_nw4).ArraySet1(_92_v1, 1)
	(_nw4).ArraySet1(!(_92_v1), 2)
	(_nw4).ArraySet1(_92_v1, 3)
	(_nw4).ArraySet1(_92_v1, 4)
	(_nw4).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(563))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}((func(_110_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_111_i1 _dafny.Int) _dafny.Int {
			return _110_v0
		}
	})(_91_v0))), _100_v9), 5)
	(_nw4).ArraySet1(false, 6)
	(_nw4).ArraySet1((_108_v15).Select((Companion_Default___.SafeIndex(_91_v0, _dafny.IntOfUint32((_108_v15).Cardinality()))).Uint32()).(bool), 7)
	(_nw4).ArraySet1((_dafny.IntOfInt64(-904)).Cmp((_99_v8).Cardinality()) < 0, 8)
	(_nw4).ArraySet1(_92_v1, 9)
	(_nw4).ArraySet1(_92_v1, 10)
	(_nw4).ArraySet1(((_99_v8).Update(_91_v0, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(29))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}((func(_112_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_113_i4 _dafny.Int) _dafny.Int {
			return _112_v0
		}
	})(_91_v0)))).Cardinality())))).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}((func(_114_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_115_i3 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v0, _114_v0)).Cardinality())
		}
	})(_91_v0))))), 11)
	(_nw4).ArraySet1(true, 12)
	(_nw4).ArraySet1(_92_v1, 13)
	(_nw4).ArraySet1(_92_v1, 14)
	(_nw4).ArraySet1(_92_v1, 15)
	(_nw4).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_108_v15, _108_v15), 16)
	(_nw4).ArraySet1((func() bool {
		if _92_v1 {
			return _92_v1
		}
		return Companion_Default___.Fm22(false, _91_v0, globalState)
	})(), 17)
	_109_v16 = _nw4
	for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_109_v16), 0))); ; {
		_guard_loop_0, _ok26 := _iter26()
		if !_ok26 {
			break
		}
		var _116_i0 _dafny.Int
		_116_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_116_i0).Sign() != -1) && ((_116_i0).Cmp(_dafny.ArrayLen((_109_v16), 0)) < 0)) {
			(_109_v16).ArraySet1(_92_v1, (_116_i0).Int())
		}
	}
	var _117_v17 T1
	_ = _117_v17
	var _nw5 *C7 = New_C7_()
	_ = _nw5
	_nw5.Ctor__(_109_v16, _92_v1, _91_v0)
	_117_v17 = _nw5
	var _118_v18 _dafny.Map
	_ = _118_v18
	_118_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_101_v10).Cardinality()), _117_v17)
	var _119_v19 _dafny.MultiSet
	_ = _119_v19
	_119_v19 = _dafny.MultiSetOf(_118_v18)
	if ((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_117_v17.F29())).Cardinality()), _117_v17)).Update(_91_v0, _117_v17))).Difference(_119_v19)).IsSubsetOf(_dafny.MultiSetOf(_118_v18, _118_v18)) {
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_102_v11), 0))
		_ = _index0
		(_102_v11).ArraySet1(_117_v17.F29(), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_102_v11), 0))
		_ = _index1
		(_102_v11).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-587), _117_v17.F29()), (_index1).Int())
		var _120_v20 _dafny.CodePoint
		_ = _120_v20
		_120_v20 = _dafny.CodePoint('p')
		(globalState).F25 = _120_v20
		_91_v0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-565), (func() _dafny.Int {
			if (_117_v17).F28() {
				return Companion_Default___.Fm1(globalState)
			}
			return (_102_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_102_v11), 0))).Int()).(_dafny.Int)
		})())))
		(globalState).F6 = (_117_v17).F28()
		(globalState).F17 = _dafny.Companion_Sequence_.Equal(_101_v10, _101_v10)
	} else {
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))
		_ = _index2
		(_102_v11).ArraySet1(_117_v17.F29(), (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))
		_ = _index3
		(_102_v11).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_117_v17.F29(), (_117_v17.F29()).Minus(_91_v0))), (_index3).Int())
		_91_v0 = ((_102_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))).Int()).(_dafny.Int)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_117_v17).F28(), _97_v6)).Cardinality())
		var _121_v21 _dafny.CodePoint
		_ = _121_v21
		_121_v21 = _dafny.CodePoint('u')
		var _122_v22 *C8
		_ = _122_v22
		var _nw6 *C8 = New_C8_()
		_ = _nw6
		_nw6.Ctor__(_121_v21, _109_v16, (_117_v17).F28())
		_122_v22 = _nw6
		var _123_v23 D1
		_ = _123_v23
		_123_v23 = Companion_D1_.Create_DC6_((_102_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))).Int()).(_dafny.Int))
		var _124_v24 _dafny.MultiSet
		_ = _124_v24
		_124_v24 = _dafny.MultiSetOf(_123_v23, _123_v23, _123_v23, _123_v23)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))
		_ = _index4
		var _rhs3 _dafny.Int = _91_v0
		_ = _rhs3
		var _rhs4 *C8 = _122_v22
		_ = _rhs4
		var _rhs5 bool = (Companion_Default___.Fm4(_124_v24, globalState)).Cmp(_117_v17.F29()) >= 0
		_ = _rhs5
		var _lhs2 _dafny.Array = _102_v11
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))
		_ = _lhs3
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
		_122_v22 = _rhs4
		_lhs4.F17 = _rhs5
		var _125_v25 *C0
		_ = _125_v25
		var _nw7 *C0 = New_C0_()
		_ = _nw7
		_nw7.Ctor__(_101_v10, (_102_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))).Int()).(_dafny.Int), _109_v16, (_117_v17).F28())
		_125_v25 = _nw7
		_125_v25 = _125_v25
		var _126_v26 _dafny.Set
		_ = _126_v26
		_126_v26 = _dafny.SetOf((_122_v22).F32(), Companion_Default___.Fm42(_dafny.IntOfUint32((_108_v15).Cardinality()), globalState))
		(globalState).F4 = ((_102_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_102_v11), 0))).Int()).(_dafny.Int)).Times(Companion_Default___.SafeModuloInt(_117_v17.F29(), (_126_v26).Cardinality()))
	}
	(globalState).F6 = (Companion_D5_.Create_DC13_(true, _102_v11, false)).Dtor_cf23()
	(globalState).F4 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg30 _dafny.Int) interface{} {
			return coer30(arg30)
		}
	}((func(_127_v17 T1) func(_dafny.Int) _dafny.Set {
		return func(_128_i5 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_127_v17.F29(), _127_v17.F29())
		}
	})(_117_v17)))).Cardinality())
	var _129_v27 _dafny.Map
	_ = _129_v27
	_129_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_117_v17.F29()), (_117_v17).F27())
	r0 = _129_v27
	r1 = _dafny.IntOfInt64(820)
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _130_v0 _dafny.Array
	_ = _130_v0
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_1
	var _nw8 _dafny.Array
	_ = _nw8
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw8 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) bool = func(_131_i0 _dafny.Int) bool {
			return false
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw8 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw8).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw8).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_130_v0 = _nw8
	var _132_v1 _dafny.Int
	_ = _132_v1
	_132_v1 = _dafny.IntOfInt64(900)
	var _133_v2 bool
	_ = _133_v2
	_133_v2 = false
	var _134_v3 _dafny.MultiSet
	_ = _134_v3
	_134_v3 = _dafny.MultiSetOf(_133_v2)
	var _135_v4 _dafny.Map
	_ = _135_v4
	_135_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_132_v1), _134_v3)
	var _136_v5 _dafny.Array
	_ = _136_v5
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
	_ = _nw9
	_136_v5 = _nw9
	var _137_v6 _dafny.Map
	_ = _137_v6
	_137_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v3, _136_v5)
	var _138_v7 _dafny.Array
	_ = _138_v7
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
	_ = _nw10
	_138_v7 = _nw10
	var _139_v8 _dafny.Array
	_ = _139_v8
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(3)
	_ = _len0_2
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Int = (func(_140_v1 _dafny.Int, _141_v2 bool) func(_dafny.Int) _dafny.Int {
			return func(_142_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_142_i1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_140_v1, _140_v1)).Cardinality(), _141_v2))).Cardinality()))
			}
		})(_132_v1, _133_v2)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw11 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw11).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw11).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_139_v8 = _nw11
	var _143_v9 _dafny.Sequence
	_ = _143_v9
	_143_v9 = _dafny.SeqOf(_139_v8, _139_v8)
	var _144_v10 _dafny.Map
	_ = _144_v10
	_144_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v2, _136_v5)
	var _145_v11 _dafny.Sequence
	_ = _145_v11
	_145_v11 = _dafny.SeqOf(_133_v2, _133_v2)
	var _146_v12 _dafny.Sequence
	_ = _146_v12
	_146_v12 = _dafny.SeqOf(_132_v1)
	var _147_v13 _dafny.Sequence
	_ = _147_v13
	_147_v13 = _dafny.SeqOf(_136_v5, (func() _dafny.Array {
		if (_144_v10).Contains((_145_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_146_v12).Cardinality()), _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool)) {
			return (_144_v10).Get((_145_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_146_v12).Cardinality()), _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool)).(_dafny.Array)
		}
		return _136_v5
	})())
	var _148_v14 _dafny.Sequence
	_ = _148_v14
	_148_v14 = _dafny.UnicodeSeqOfUtf8Bytes("prvhys")
	var _149_v15 _dafny.Sequence
	_ = _149_v15
	_149_v15 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("y"), _148_v14)
	var _150_v16 _dafny.Map
	_ = _150_v16
	_150_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(76), _139_v8)
	var _151_globalState *GlobalState
	_ = _151_globalState
	var _nw12 *GlobalState = New_GlobalState_()
	_ = _nw12
	_nw12.Ctor__(false, true, _dafny.IntOfInt64(888), true, _dafny.IntOfInt64(277), _dafny.IntOfInt64(131), true, _dafny.IntOfInt64(306), _130_v0, _135_v4, _dafny.MultiSetOf(_132_v1), (func() _dafny.Array {
		if (_137_v6).Contains(_134_v3) {
			return (_137_v6).Get(_134_v3).(_dafny.Array)
		}
		return _136_v5
	})(), _dafny.IntOfInt64(321), _dafny.IntOfInt64(529), _138_v7, false, _dafny.Companion_Sequence_.Concatenate(_143_v9, _143_v9), false, (_147_v13).Select((Companion_Default___.SafeIndex(_132_v1, _dafny.IntOfUint32((_147_v13).Cardinality()))).Uint32()).(_dafny.Array), _dafny.Companion_Sequence_.Concatenate(_149_v15, _149_v15), _dafny.IntOfInt64(778), _dafny.IntOfInt64(481), _150_v16, _dafny.IntOfInt64(726), _dafny.IntOfInt64(300), _dafny.CodePoint('u'), _dafny.IntOfInt64(792))
	_151_globalState = _nw12
	var _152_v17 _dafny.Map
	_ = _152_v17
	_152_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v2, _132_v1)
	var _153_v18 _dafny.MultiSet
	_ = _153_v18
	_153_v18 = _dafny.MultiSetOf(_152_v17)
	if (_153_v18).IsSubsetOf(_dafny.MultiSetOf(_152_v17)) {
		var _154_v19 _dafny.Map
		_ = _154_v19
		var _155_v20 _dafny.Int
		_ = _155_v20
		var _out0 _dafny.Map
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		_out0, _out1 = Companion_Default___.M0(_151_globalState)
		_154_v19 = _out0
		_155_v20 = _out1
		var _156_v21 _dafny.MultiSet
		_ = _156_v21
		_156_v21 = _dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Int {
			if (_134_v3).Contains(_133_v2) {
				return (_134_v3).Multiplicity(_133_v2)
			}
			return _155_v20
		})()), _dafny.IntOfUint32((_148_v14).Cardinality()))
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_130_v0), 0))
		_ = _index5
		(_130_v0).ArraySet1(_133_v2, (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_130_v0), 0))
		_ = _index6
		(_130_v0).ArraySet1(_133_v2, (_index6).Int())
		var _157_v22 _dafny.Map
		_ = _157_v22
		_157_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_133_v2) || (_133_v2), _133_v2)
		var _158_v23 _dafny.Set
		_ = _158_v23
		_158_v23 = _dafny.SetOf(_132_v1, _132_v1, (_dafny.Zero).Minus(_132_v1), (_134_v3).Cardinality())
		var _159_v24 _dafny.Sequence
		_ = _159_v24
		_159_v24 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v20, _158_v23))
		var _160_v25 _dafny.Map
		_ = _160_v25
		_160_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v0, true)
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_130_v0), 0))
		_ = _index7
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_130_v0), 0))
		_ = _index8
		var _rhs6 _dafny.MultiSet = _156_v21
		_ = _rhs6
		var _rhs7 bool = _133_v2
		_ = _rhs7
		var _rhs8 _dafny.Int = (Companion_Default___.Fm0((_159_v24).Select((Companion_Default___.SafeIndex(_155_v20, _dafny.IntOfUint32((_159_v24).Cardinality()))).Uint32()).(_dafny.Map), (func() bool {
			if (_160_v25).Contains(_130_v0) {
				return (_160_v25).Get(_130_v0).(bool)
			}
			return _133_v2
		})(), _133_v2, _151_globalState)).Cardinality()
		_ = _rhs8
		var _rhs9 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_146_v12, _146_v12)).Cardinality())).Cmp((_155_v20).Minus(_dafny.IntOfUint32((_148_v14).Cardinality()))) <= 0
		_ = _rhs9
		var _rhs10 _dafny.Map = _157_v22
		_ = _rhs10
		var _lhs5 _dafny.Array = _130_v0
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_130_v0), 0))
		_ = _lhs6
		var _lhs7 _dafny.Array = _130_v0
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_130_v0), 0))
		_ = _lhs8
		_156_v21 = _rhs6
		(_lhs5).ArraySet1(_rhs7, (_lhs6).Int())
		_132_v1 = _rhs8
		(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
		_157_v22 = _rhs10
		var _161_v26 *C4
		_ = _161_v26
		var _nw13 *C4 = New_C4_()
		_ = _nw13
		_nw13.Ctor__()
		_161_v26 = _nw13
		_148_v14 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_148_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_162_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))), _148_v14)
		_152_v17 = (_152_v17).Update(_133_v2, _155_v20)
	} else {
		_148_v14 = _148_v14
		var _163_v27 _dafny.Array
		_ = _163_v27
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
		_ = _nw14
		_163_v27 = _nw14
		_163_v27 = _163_v27
		var _164_v28 _dafny.Map
		_ = _164_v28
		_164_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v8, _148_v14)
		_164_v28 = _164_v28
		var _rhs11 _dafny.MultiSet = _134_v3
		_ = _rhs11
		var _rhs12 bool = _133_v2
		_ = _rhs12
		var _lhs9 *GlobalState = _151_globalState
		_ = _lhs9
		_134_v3 = _rhs11
		_lhs9.F6 = _rhs12
		var _165_v29 _dafny.Set
		_ = _165_v29
		_165_v29 = _dafny.SetOf((func() _dafny.Int {
			if (_134_v3).Contains(_133_v2) {
				return (_134_v3).Multiplicity(_133_v2)
			}
			return _132_v1
		})())
		_165_v29 = (func() _dafny.Set {
			if _133_v2 {
				return (func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(667), _dafny.IntOfInt64(735))); ; {
						_compr_26, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _166_v30 _dafny.Int
						_166_v30 = interface{}(_compr_26).(_dafny.Int)
						if ((_dafny.IntOfInt64(667)).Cmp(_166_v30) <= 0) && ((_166_v30).Cmp(_dafny.IntOfInt64(735)) < 0) {
							_coll26.Add(Companion_Default___.SafeModuloInt(_166_v30, _132_v1))
						}
					}
					return _coll26.ToSet()
				}()).Union(_165_v29)
			}
			return (func() _dafny.Set {
				if (_145_v11).Select((Companion_Default___.SafeIndex(_132_v1, _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool) {
					return _165_v29
				}
				return _165_v29
			})()
		})()
	}
	var _167_v31 _dafny.Set
	_ = _167_v31
	_167_v31 = _dafny.SetOf(true, false)
	var _168_v32 _dafny.Sequence
	_ = _168_v32
	_168_v32 = _dafny.SeqOf(_167_v31)
	var _169_v33 _dafny.Sequence
	_ = _169_v33
	_169_v33 = _dafny.SeqOf(Companion_D5_.Create_DC12_(_168_v32))
	var _170_v34 _dafny.Sequence
	_ = _170_v34
	_170_v34 = _dafny.SeqOf(_169_v33)
	var _171_v35 _dafny.Sequence
	_ = _171_v35
	_171_v35 = _170_v34
	var _172_v36 *C6
	_ = _172_v36
	var _nw15 *C6 = New_C6_()
	_ = _nw15
	_nw15.Ctor__(_dafny.IntOfInt64(609), (func() _dafny.Array {
		if _133_v2 {
			return _130_v0
		}
		return _130_v0
	})(), _dafny.Companion_Sequence_.Contains((_171_v35), _169_v33))
	_172_v36 = _nw15
	if _133_v2 {
		(_151_globalState).F4 = _132_v1
		var _173_v37 _dafny.Set
		_ = _173_v37
		_173_v37 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-185)))
		var _rhs13 _dafny.Set = (_173_v37).Intersection(_173_v37)
		_ = _rhs13
		var _rhs14 _dafny.Array = _130_v0
		_ = _rhs14
		_173_v37 = _rhs13
		_130_v0 = _rhs14
		_152_v17 = (_152_v17).Update(_133_v2, _132_v1)
		(_151_globalState).F13 = _132_v1
		(_151_globalState).F5 = _132_v1
	} else {
		var _174_v38 _dafny.Set
		_ = _174_v38
		_174_v38 = _dafny.SetOf(_132_v1)
		var _175_v39 D0
		_ = _175_v39
		_175_v39 = Companion_D0_.Create_DC2_(_174_v38, _148_v14)
		var _176_v40 _dafny.Map
		_ = _176_v40
		_176_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v1, _132_v1)
		var _177_v41 _dafny.MultiSet
		_ = _177_v41
		_177_v41 = _dafny.MultiSetOf(_132_v1, ((_175_v39).Dtor_cf5()).Cardinality(), ((_146_v12).Select((Companion_Default___.SafeIndex(_132_v1, _dafny.IntOfUint32((_146_v12).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_176_v40).Cardinality()), _132_v1)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_139_v8), 0))
		_ = _index9
		(_139_v8).ArraySet1((_132_v1).Plus((_dafny.Zero).Minus(_132_v1)), (_index9).Int())
		var _178_v42 _dafny.Map
		_ = _178_v42
		_178_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v1, _174_v38)
		var _179_v43 _dafny.Map
		_ = _179_v43
		_179_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v1, _145_v11)
		var _pat_let_tv0 = _132_v1
		_ = _pat_let_tv0
		var _pat_let_tv1 = _132_v1
		_ = _pat_let_tv1
		var _pat_let_tv2 = _132_v1
		_ = _pat_let_tv2
		var _pat_let_tv3 = _178_v42
		_ = _pat_let_tv3
		var _pat_let_tv4 = _133_v2
		_ = _pat_let_tv4
		var _pat_let_tv5 = _133_v2
		_ = _pat_let_tv5
		var _pat_let_tv6 = _151_globalState
		_ = _pat_let_tv6
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_139_v8), 0))
		_ = _index10
		var _rhs15 D0 = func(_pat_let0_0 D0) D0 {
			return func(_182_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let3_0 _dafny.Set) D0 {
					return func(_183_dt__update_hcf5_h1 _dafny.Set) D0 {
						return Companion_D0_.Create_DC2_(_183_dt__update_hcf5_h1, (_182_dt__update__tmp_h1).Dtor_cf6())
					}(_pat_let3_0)
				}(Companion_Default___.Fm0(_pat_let_tv3, _pat_let_tv4, _pat_let_tv5, _pat_let_tv6))
			}(_pat_let0_0)
		}(func(_pat_let1_0 D0) D0 {
			return func(_180_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let2_0 _dafny.Set) D0 {
					return func(_181_dt__update_hcf5_h0 _dafny.Set) D0 {
						return Companion_D0_.Create_DC2_(_181_dt__update_hcf5_h0, (_180_dt__update__tmp_h0).Dtor_cf6())
					}(_pat_let2_0)
				}(_dafny.SetOf(_pat_let_tv0, (_dafny.Zero).Minus(_pat_let_tv1), _pat_let_tv2))
			}(_pat_let1_0)
		}(_175_v39))
		_ = _rhs15
		var _rhs16 _dafny.MultiSet = _177_v41
		_ = _rhs16
		var _rhs17 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_179_v43).Contains(_132_v1) {
				return (_179_v43).Get(_132_v1).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_133_v2)
		})()).Cardinality()), _dafny.IntOfUint32((_145_v11).Cardinality()))).Times(((_167_v31).Cardinality()).Plus(_132_v1)))
		_ = _rhs17
		var _rhs18 bool = _133_v2
		_ = _rhs18
		var _lhs10 _dafny.Array = _139_v8
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_139_v8), 0))
		_ = _lhs11
		var _lhs12 *GlobalState = _151_globalState
		_ = _lhs12
		_175_v39 = _rhs15
		_177_v41 = _rhs16
		(_lhs10).ArraySet1(_rhs17, (_lhs11).Int())
		_lhs12.F15 = _rhs18
		var _184_v44 _dafny.Array
		_ = _184_v44
		var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
		_ = _nw16
		_184_v44 = _nw16
		_184_v44 = _184_v44
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_139_v8), 0))
		_ = _index11
		(_139_v8).ArraySet1(Companion_Default___.SafeModuloInt(_132_v1, (_139_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_139_v8), 0))).Int()).(_dafny.Int)), (_index11).Int())
		var _185_v45 D8
		_ = _185_v45
		_185_v45 = Companion_D8_.Create_DC19_()
		var _186_v46 _dafny.CodePoint
		_ = _186_v46
		_186_v46 = _dafny.CodePoint('a')
		var _187_v47 _dafny.Set
		_ = _187_v47
		_187_v47 = _dafny.SetOf(_186_v46, _186_v46, _dafny.CodePoint('p'), _186_v46, _186_v46)
		var _188_v48 _dafny.Map
		_ = _188_v48
		_188_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v1, Companion_D8_.Create_DC20_(_185_v45)), (func() _dafny.Set {
			if true {
				return _187_v47
			}
			return _187_v47
		})())
		var _189_v49 _dafny.Map
		_ = _189_v49
		_189_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v2, _167_v31)
		var _rhs19 bool = ((func() _dafny.Map {
			if _133_v2 {
				return _189_v49
			}
			return _189_v49
		})()).Contains(_133_v2)
		_ = _rhs19
		var _rhs20 _dafny.Map = Companion_Default___.Fm54((_132_v1).Cmp((_139_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_139_v8), 0))).Int()).(_dafny.Int)) < 0, false, (_133_v2) || ((_145_v11).Select((Companion_Default___.SafeIndex(_132_v1, _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool)), _133_v2, _151_globalState)
		_ = _rhs20
		var _lhs13 *GlobalState = _151_globalState
		_ = _lhs13
		_lhs13.F17 = _rhs19
		_188_v48 = _rhs20
		var _190_v50 D0
		_ = _190_v50
		_190_v50 = Companion_D0_.Create_DC3_(Companion_Default___.Fm36(_151_globalState), Companion_Default___.Fm2(Companion_Default___.Fm42(_132_v1, _151_globalState), (_167_v31).Cardinality(), _151_globalState), !((_133_v2) == (_133_v2)))
		var _pat_let_tv7 = _133_v2
		_ = _pat_let_tv7
		var _pat_let_tv8 = _139_v8
		_ = _pat_let_tv8
		var _pat_let_tv9 = _139_v8
		_ = _pat_let_tv9
		var _pat_let_tv10 = _152_v17
		_ = _pat_let_tv10
		_190_v50 = func(_pat_let4_0 D0) D0 {
			return func(_191_dt__update__tmp_h2 D0) D0 {
				return func(_pat_let5_0 bool) D0 {
					return func(_192_dt__update_hcf9_h0 bool) D0 {
						return func(_pat_let6_0 bool) D0 {
							return func(_193_dt__update_hcf8_h0 bool) D0 {
								return Companion_D0_.Create_DC3_((_191_dt__update__tmp_h2).Dtor_cf7(), _193_dt__update_hcf8_h0, _192_dt__update_hcf9_h0)
							}(_pat_let6_0)
						}(((_pat_let_tv9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_pat_let_tv8), 0))).Int()).(_dafny.Int)).Cmp((_pat_let_tv10).Cardinality()) < 0)
					}(_pat_let5_0)
				}(_pat_let_tv7)
			}(_pat_let4_0)
		}(_190_v50)
	}
	(_151_globalState).F6 = _133_v2
	(_151_globalState).F6 = _133_v2
	var _194_v51 _dafny.CodePoint
	_ = _194_v51
	_194_v51 = _dafny.CodePoint('e')
	(_151_globalState).F25 = _194_v51
	var _195_v52 D8
	_ = _195_v52
	_195_v52 = Companion_D8_.Create_DC18_(_132_v1)
	var _196_v53 D8
	_ = _196_v53
	_196_v53 = Companion_D8_.Create_DC20_(_195_v52)
	var _197_v54 D8
	_ = _197_v54
	_197_v54 = Companion_D8_.Create_DC20_(_195_v52)
	var _198_v55 D8
	_ = _198_v55
	_198_v55 = Companion_D8_.Create_DC20_(_197_v54)
	_198_v55 = Companion_D8_.Create_DC20_(_196_v53)
	var _199_v56 D0
	_ = _199_v56
	_199_v56 = Companion_D0_.Create_DC3_(_132_v1, _133_v2, _133_v2)
	var _200_v57 _dafny.Map
	_ = _200_v57
	_200_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_v56, _dafny.IntOfUint32((_148_v14).Cardinality()))
	var _201_v58 *C3
	_ = _201_v58
	var _nw17 *C3 = New_C3_()
	_ = _nw17
	_nw17.Ctor__(_dafny.IntOfUint32((_148_v14).Cardinality()), _200_v57, _132_v1, _130_v0, (_132_v1).Cmp(_132_v1) != 0)
	_201_v58 = _nw17
	var _202_v59 D10
	_ = _202_v59
	_202_v59 = Companion_D10_.Create_DC25_(_dafny.MultiSetOf(_133_v2))
	var _203_v60 _dafny.Set
	_ = _203_v60
	_203_v60 = _dafny.SetOf((_202_v59).Dtor_cf40())
	var _204_i3 _dafny.Int
	_ = _204_i3
	_204_i3 = _dafny.Zero
	{
		for (_203_v60).Contains(_134_v3) {
			{
				if (_204_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_204_i3 = (_204_i3).Plus(_dafny.One)
				(_151_globalState).F25 = _194_v51
				(_172_v36).M6(_133_v2, (_133_v2) && (_133_v2), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_201_v58).F34()), _dafny.IntOfUint32((_145_v11).Cardinality())), _151_globalState)
				var _205_v61 _dafny.MultiSet
				_ = _205_v61
				_205_v61 = _dafny.MultiSetOf(_132_v1)
				var _206_v62 *C9
				_ = _206_v62
				var _nw18 *C9 = New_C9_()
				_ = _nw18
				_nw18.Ctor__(_133_v2, _205_v61, _132_v1, _130_v0, _133_v2)
				_206_v62 = _nw18
				var _207_v63 _dafny.Map
				_ = _207_v63
				_207_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v1, _206_v62)
				var _208_v64 _dafny.Sequence
				_ = _208_v64
				_208_v64 = _dafny.SeqOf(_152_v17)
				_207_v63 = (_207_v63).Update(_dafny.IntOfUint32((_208_v64).Cardinality()), _206_v62)
				var _209_v65 _dafny.Map
				_ = _209_v65
				_209_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v1, _133_v2)
				var _210_v67 _dafny.Set
				_ = _210_v67
				_210_v67 = _dafny.SetOf(_209_v65, func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter28 := _dafny.Iterate(((_206_v62).F31()).Elements()); ; {
						_compr_27, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _211_v66 _dafny.Int
						_211_v66 = interface{}(_compr_27).(_dafny.Int)
						if ((_206_v62).F31()).Contains(_211_v66) {
							_coll27.Add((_211_v66).Plus((_201_v58).F34()), !(_133_v2))
						}
					}
					return _coll27.ToMap()
				}())
				(_151_globalState).F6 = ((_210_v67).Intersection(_210_v67)).IsSubsetOf((_210_v67).Difference(_210_v67))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _212_v68 _dafny.Map
	_ = _212_v68
	_212_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(498))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg32 _dafny.Int) interface{} {
			return coer32(arg32)
		}
	}((func(_213_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_214_i4 _dafny.Int) _dafny.CodePoint {
			return _213_v51
		}
	})(_194_v51)))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_201_v58).F34(), _133_v2)).Cardinality())
	var _215_v69 _dafny.Sequence
	_ = _215_v69
	_215_v69 = _dafny.SeqOf(_172_v36, _172_v36)
	var _216_v70 bool
	_ = _216_v70
	var _217_v71 bool
	_ = _217_v71
	var _218_v72 _dafny.CodePoint
	_ = _218_v72
	var _219_v73 _dafny.Int
	_ = _219_v73
	var _out2 bool
	_ = _out2
	var _out3 bool
	_ = _out3
	var _out4 _dafny.CodePoint
	_ = _out4
	var _out5 _dafny.Int
	_ = _out5
	_out2, _out3, _out4, _out5 = (_201_v58).M10(((func() _dafny.Int {
		if (_212_v68).Contains(_132_v1) {
			return (_212_v68).Get(_132_v1).(_dafny.Int)
		}
		return _dafny.IntOfUint32((_215_v69).Cardinality())
	})()).Cmp((_201_v58).Fm14((_201_v58).F34(), _203_v60, _151_globalState)) > 0, _151_globalState)
	_216_v70 = _out2
	_217_v71 = _out3
	_218_v72 = _out4
	_219_v73 = _out5
	var _220_v74 _dafny.Set
	_ = _220_v74
	_220_v74 = _dafny.SetOf(_130_v0, _130_v0, _130_v0)
	var _221_v75 _dafny.Map
	_ = _221_v75
	_221_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v74, (_167_v31).IsProperSubsetOf(_167_v31))
	var _222_v76 _dafny.Set
	_ = _222_v76
	_222_v76 = _dafny.SetOf(_dafny.IntOfInt64(-854))
	var _223_v77 _dafny.Map
	_ = _223_v77
	_223_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v71, _222_v76)
	var _224_v78 _dafny.Map
	_ = _224_v78
	_224_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v73, _222_v76)
	_221_v75 = (_221_v75).Update((_220_v74).Intersection(_220_v74), (Companion_Default___.Fm0(_224_v78, _216_v70, true, _151_globalState)).IsSubsetOf((func() _dafny.Set {
		if (_223_v77).Contains(_133_v2) {
			return (_223_v77).Get(_133_v2).(_dafny.Set)
		}
		return _222_v76
	})()))
	if (_145_v11).Select((Companion_Default___.SafeIndex(_219_v73, _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool) {
		_132_v1 = Companion_Default___.SafeDivisionInt((_201_v58).F34(), (_dafny.Zero).Minus((_201_v58).F34()))
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))
		_ = _index12
		(_139_v8).ArraySet1(_219_v73, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))
		_ = _index13
		(_139_v8).ArraySet1(Companion_Default___.Fm36(_151_globalState), (_index13).Int())
		if false {
			var _225_v79 _dafny.Map
			_ = _225_v79
			_225_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_201_v58).F34(), _216_v70)
			_225_v79 = (func() _dafny.Map {
				if (_dafny.IntOfInt64(564)).Cmp(_219_v73) != 0 {
					return (_225_v79).Merge(Companion_Default___.Fm30((func() bool {
						if (_225_v79).Contains((_139_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))).Int()).(_dafny.Int)) {
							return (_225_v79).Get((_139_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))).Int()).(_dafny.Int)).(bool)
						}
						return _216_v70
					})(), _148_v14, _133_v2, _151_globalState))
				}
				return _225_v79
			})()
			var _226_v80 _dafny.Map
			_ = _226_v80
			_226_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v2, _130_v0)
			var _227_v81 T0
			_ = _227_v81
			var _nw19 *C7 = New_C7_()
			_ = _nw19
			_nw19.Ctor__(_130_v0, _217_v71, _132_v1)
			_227_v81 = _nw19
			var _228_v82 _dafny.Map
			_ = _228_v82
			_228_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v81, (_227_v81).F27())
			var _229_v83 _dafny.Sequence
			_ = _229_v83
			_229_v83 = _dafny.SeqOf((func() _dafny.Array {
				if (_226_v80).Contains(!(_216_v70)) {
					return (_226_v80).Get(!(_216_v70)).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if (_228_v82).Contains(_227_v81) {
						return (_228_v82).Get(_227_v81).(_dafny.Array)
					}
					return _130_v0
				})()
			})(), _130_v0, (_227_v81).F27(), (_227_v81).F27())
			_130_v0 = (_229_v83).Select((Companion_Default___.SafeIndex(_219_v73, _dafny.IntOfUint32((_229_v83).Cardinality()))).Uint32()).(_dafny.Array)
			(_151_globalState).F15 = _216_v70
			var _230_v84 _dafny.Array
			_ = _230_v84
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw20
			_230_v84 = _nw20
			var _231_v85 _dafny.Sequence
			_ = _231_v85
			_231_v85 = _dafny.SeqOf(_dafny.SeqOf(_230_v84))
			var _232_v86 *C5
			_ = _232_v86
			var _nw21 *C5 = New_C5_()
			_ = _nw21
			_nw21.Ctor__(_dafny.Companion_Sequence_.Concatenate((_231_v85).Select((Companion_Default___.SafeIndex((_201_v58).Fm14(_dafny.IntOfInt64(83), _203_v60, _151_globalState), _dafny.IntOfUint32((_231_v85).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update(_143_v9, (Companion_Default___.SafeIndex((_201_v58).F34(), _dafny.IntOfUint32((_143_v9).Cardinality()))).Uint32(), _230_v84)), (func() _dafny.Int {
				if _216_v70 {
					return (_201_v58).F34()
				}
				return _dafny.IntOfInt64(-87)
			})(), _130_v0, _216_v70)
			_232_v86 = _nw21
			_219_v73 = _219_v73
		} else {
			_218_v72 = _218_v72
			_133_v2 = !(_133_v2)
			var _233_v87 *C4
			_ = _233_v87
			var _nw22 *C4 = New_C4_()
			_ = _nw22
			_nw22.Ctor__()
			_233_v87 = _nw22
			var _234_v88 _dafny.Map
			_ = _234_v88
			_234_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_146_v12, _146_v12), _233_v87)
			_234_v88 = (_234_v88).Update(Companion_Default___.Fm55(_216_v70, _151_globalState), _233_v87)
			(_151_globalState).F21 = (_201_v58).F34()
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_130_v0), 0))
			_ = _index14
			(_130_v0).ArraySet1(_217_v71, (_index14).Int())
			var _235_v89 _dafny.Array
			_ = _235_v89
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw23
			_235_v89 = _nw23
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_235_v89), 0))
			_ = _index15
			(_235_v89).ArraySet1(_130_v0, (_index15).Int())
			var _236_v90 D1
			_ = _236_v90
			_236_v90 = Companion_D1_.Create_DC6_(_dafny.IntOfInt64(318))
			var _237_v91 _dafny.MultiSet
			_ = _237_v91
			_237_v91 = _dafny.MultiSetOf(Companion_D1_.Create_DC6_(((_212_v68).Update((_201_v58).F34(), _dafny.IntOfInt64(117))).Cardinality()), _236_v90, Companion_D1_.Create_DC6_(_132_v1))
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_130_v0), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))
			_ = _index17
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))
			_ = _index18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_235_v89), 0))
			_ = _index19
			var _rhs21 bool = ((_201_v58).F34()).Cmp(_dafny.IntOfInt64(956)) == 0
			_ = _rhs21
			var _rhs22 _dafny.Sequence = _146_v12
			_ = _rhs22
			var _rhs23 _dafny.Int = Companion_Default___.Fm4(_237_v91, _151_globalState)
			_ = _rhs23
			var _rhs24 _dafny.Int = (_201_v58).F34()
			_ = _rhs24
			var _rhs25 _dafny.Array = _130_v0
			_ = _rhs25
			var _lhs14 _dafny.Array = _130_v0
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_130_v0), 0))
			_ = _lhs15
			var _lhs16 _dafny.Array = _139_v8
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))
			_ = _lhs17
			var _lhs18 _dafny.Array = _139_v8
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_139_v8), 0))
			_ = _lhs19
			var _lhs20 _dafny.Array = _235_v89
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_235_v89), 0))
			_ = _lhs21
			(_lhs14).ArraySet1(_rhs21, (_lhs15).Int())
			_146_v12 = _rhs22
			(_lhs16).ArraySet1(_rhs23, (_lhs17).Int())
			(_lhs18).ArraySet1(_rhs24, (_lhs19).Int())
			(_lhs20).ArraySet1(_rhs25, (_lhs21).Int())
		}
		_132_v1 = _132_v1
		var _238_v92 _dafny.MultiSet
		_ = _238_v92
		_238_v92 = _dafny.MultiSetOf(_132_v1)
		var _239_v93 _dafny.MultiSet
		_ = _239_v93
		_239_v93 = _dafny.MultiSetOf(_238_v92, _238_v92)
		var _240_v94 T0
		_ = _240_v94
		var _nw24 *C2 = New_C2_()
		_ = _nw24
		_nw24.Ctor__(Companion_Default___.Fm42(_dafny.IntOfInt64(992), _151_globalState), _217_v71, _130_v0, _216_v70)
		_240_v94 = _nw24
		var _241_v95 _dafny.Sequence
		_ = _241_v95
		_241_v95 = _dafny.SeqOf(_240_v94, _240_v94)
		var _242_v96 _dafny.Map
		_ = _242_v96
		_242_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_239_v93).Update(_238_v92, Companion_Default___.Abs(_dafny.IntOfUint32((_241_v95).Cardinality()))), _218_v72)
		_242_v96 = (_242_v96).Update(_239_v93, _194_v51)
	} else {
		var _243_v97 _dafny.MultiSet
		_ = _243_v97
		_243_v97 = _dafny.MultiSetOf(_148_v14)
		(_151_globalState).F17 = (_243_v97).IsSubsetOf(_243_v97)
		_194_v51 = (func() _dafny.CodePoint {
			if _133_v2 {
				return _218_v72
			}
			return _194_v51
		})()
		var _244_v98 _dafny.Array
		_ = _244_v98
		var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
		_ = _nw25
		_244_v98 = _nw25
		var _245_v99 _dafny.Map
		_ = _245_v99
		_245_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v71, _244_v98)
		var _246_v100 D5
		_ = _246_v100
		_246_v100 = Companion_D5_.Create_DC13_(false, _139_v8, _133_v2)
		_244_v98 = (func() _dafny.Array {
			if (_245_v99).Contains((_246_v100).Dtor_cf21()) {
				return (_245_v99).Get((_246_v100).Dtor_cf21()).(_dafny.Array)
			}
			return _244_v98
		})()
		if Companion_Default___.Fm22(!((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-179)), (_201_v58).F34())).IsSubsetOf(_222_v76)), _dafny.IntOfUint32(((_149_v15).Select((Companion_Default___.SafeIndex(_219_v73, _dafny.IntOfUint32((_149_v15).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _151_globalState) {
			var _247_v101 T0
			_ = _247_v101
			var _nw26 *C7 = New_C7_()
			_ = _nw26
			_nw26.Ctor__(_130_v0, Companion_Default___.Fm2(_218_v72, (_201_v58).F34(), _151_globalState), _219_v73)
			_247_v101 = _nw26
			var _248_v102 T1
			_ = _248_v102
			var _nw27 *C9 = New_C9_()
			_ = _nw27
			_nw27.Ctor__(_217_v71, _dafny.MultiSetOf((_dafny.Zero).Minus(_219_v73)), _132_v1, _130_v0, _133_v2)
			_248_v102 = _nw27
			var _249_v103 _dafny.Map
			_ = _249_v103
			_249_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v102, _247_v101)
			var _rhs26 T0 = (func() T0 {
				if (_249_v103).Contains(_248_v102) {
					return (_249_v103).Get(_248_v102).(T0)
				}
				return _247_v101
			})()
			_ = _rhs26
			var _rhs27 *C6 = _172_v36
			_ = _rhs27
			_247_v101 = _rhs26
			_172_v36 = _rhs27
			var _250_v104 _dafny.Map
			_ = _250_v104
			_250_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v70, !(_217_v71))
			_250_v104 = (_250_v104).Update((_247_v101).F28(), _216_v70)
			var _251_v105 D8
			_ = _251_v105
			_251_v105 = Companion_D8_.Create_DC19_()
			var _252_v106 _dafny.Map
			_ = _252_v106
			_252_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v105, Companion_Default___.Fm40(_151_globalState))
			var _253_v107 _dafny.Map
			_ = _253_v107
			_253_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(860), _252_v106)
			var _rhs28 _dafny.Int = _248_v102.F29()
			_ = _rhs28
			var _rhs29 bool = !(_217_v71)
			_ = _rhs29
			var _rhs30 _dafny.Map = ((_252_v106).Merge(_252_v106)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v105, _133_v2)).Merge(((func() _dafny.Map {
				if (_253_v107).Contains(_248_v102.F29()) {
					return (_253_v107).Get(_248_v102.F29()).(_dafny.Map)
				}
				return _252_v106
			})()).Update(_251_v105, _216_v70)))
			_ = _rhs30
			var _rhs31 _dafny.CodePoint = _218_v72
			_ = _rhs31
			var _lhs22 *GlobalState = _151_globalState
			_ = _lhs22
			var _lhs23 *GlobalState = _151_globalState
			_ = _lhs23
			var _lhs24 *GlobalState = _151_globalState
			_ = _lhs24
			_lhs22.F4 = _rhs28
			_lhs23.F15 = _rhs29
			_252_v106 = _rhs30
			_lhs24.F25 = _rhs31
			_149_v15 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(990))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_254_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_255_i5 _dafny.Int) _dafny.Sequence {
					return _254_v14
				}
			})(_148_v14))), _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(808))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_256_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			})), _148_v14, _dafny.UnicodeSeqOfUtf8Bytes("fqiqust"), _dafny.UnicodeSeqOfUtf8Bytes("uapyvflv")))
			(_151_globalState).F15 = _217_v71
		} else {
			var _257_v108 *C6
			_ = _257_v108
			var _nw28 *C6 = New_C6_()
			_ = _nw28
			_nw28.Ctor__(_dafny.IntOfUint32((_146_v12).Cardinality()), _130_v0, _217_v71)
			_257_v108 = _nw28
			_148_v14 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_148_v14, _148_v14), _148_v14)
			(_151_globalState).F6 = _133_v2
			var _258_v109 _dafny.Set
			_ = _258_v109
			_258_v109 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v70, (_201_v58).F34()), _152_v17, _152_v17, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v70, _132_v1), Companion_Default___.Fm50(_133_v2, (Companion_Default___.Fm31(_151_globalState)).Cardinality(), _151_globalState))
			var _259_v110 _dafny.Map
			_ = _259_v110
			_259_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v109, _139_v8)
			_259_v110 = (_259_v110).Update(_258_v109, _139_v8)
			var _260_v111 *C4
			_ = _260_v111
			var _nw29 *C4 = New_C4_()
			_ = _nw29
			_nw29.Ctor__()
			_260_v111 = _nw29
		}
		_132_v1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_148_v14, (_149_v15).Select((Companion_Default___.SafeIndex(_132_v1, _dafny.IntOfUint32((_149_v15).Cardinality()))).Uint32()).(_dafny.Sequence)), _dafny.Companion_Sequence_.Concatenate(_148_v14, Companion_Default___.Fm32((_243_v97).Cardinality(), _216_v70, _151_globalState)))).Cardinality())
	}
	var _261_v112 _dafny.Map
	_ = _261_v112
	_261_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v70, _133_v2)
	_261_v112 = (_261_v112).Update(!(!(_217_v71)) || (_216_v70), _217_v71)
	var _262_v113 *C0
	_ = _262_v113
	var _nw30 *C0 = New_C0_()
	_ = _nw30
	_nw30.Ctor__(_148_v14, _132_v1, _130_v0, _216_v70)
	_262_v113 = _nw30
	(_151_globalState).F21 = (Companion_Default___.SafeDivisionInt(_219_v73, _132_v1)).Plus(_132_v1)
	if !((_133_v2) || (_217_v71)) {
		var _263_v115 _dafny.MultiSet
		_ = _263_v115
		_263_v115 = _dafny.MultiSetOf(_132_v1)
		var _264_v116 _dafny.Sequence
		_ = _264_v116
		_264_v116 = _dafny.SeqOf(func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter29 := _dafny.Iterate((_263_v115).Elements()); ; {
				_compr_28, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _265_v114 _dafny.Int
				_265_v114 = interface{}(_compr_28).(_dafny.Int)
				if (_263_v115).Contains(_265_v114) {
					_coll28.Add((_265_v114).Minus((_201_v58).F34()), (_145_v11).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_201_v58).F34()), _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool))
				}
			}
			return _coll28.ToMap()
		}())
		_264_v116 = _264_v116
		_212_v68 = (_212_v68).Update(_219_v73, _dafny.IntOfUint32(((_262_v113).F36()).Cardinality()))
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_130_v0), 0))
		_ = _index20
		(_130_v0).ArraySet1(!(_133_v2), (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_130_v0), 0))
		_ = _index21
		var _rhs32 _dafny.Sequence = (_149_v15).Select((Companion_Default___.SafeIndex((_201_v58).F34(), _dafny.IntOfUint32((_149_v15).Cardinality()))).Uint32()).(_dafny.Sequence)
		_ = _rhs32
		var _rhs33 bool = _133_v2
		_ = _rhs33
		var _rhs34 bool = (_132_v1).Cmp(((_263_v115).Cardinality()).Times(_dafny.IntOfUint32((_145_v11).Cardinality()))) >= 0
		_ = _rhs34
		var _rhs35 bool = _dafny.Companion_Sequence_.Contains((_262_v113).F36(), _dafny.CodePoint('e'))
		_ = _rhs35
		var _lhs25 _dafny.Array = _130_v0
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_130_v0), 0))
		_ = _lhs26
		var _lhs27 *GlobalState = _151_globalState
		_ = _lhs27
		_148_v14 = _rhs32
		(_lhs25).ArraySet1(_rhs33, (_lhs26).Int())
		_217_v71 = _rhs34
		_lhs27.F15 = _rhs35
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_139_v8), 0))
		_ = _index22
		(_139_v8).ArraySet1(_219_v73, (_index22).Int())
		var _266_v117 D1
		_ = _266_v117
		_266_v117 = Companion_D1_.Create_DC6_((_201_v58).F34())
		var _267_v118 _dafny.MultiSet
		_ = _267_v118
		_267_v118 = _dafny.MultiSetOf(Companion_D1_.Create_DC6_((_263_v115).Cardinality()), _266_v117, _266_v117, _266_v117)
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_139_v8), 0))
		_ = _index23
		(_139_v8).ArraySet1((Companion_Default___.Fm4(_267_v118, _151_globalState)).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_145_v11).Cardinality()), _219_v73)), (_index23).Int())
		var _268_v119 _dafny.Array
		_ = _268_v119
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(11))
		_ = _nw31
		_268_v119 = _nw31
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_268_v119), 0))
		_ = _index24
		(_268_v119).ArraySet1((func() D9 {
			if (_130_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_130_v0), 0))).Int()).(bool) {
				return Companion_D9_.Create_DC21_(_138_v7)
			}
			return Companion_D9_.Create_DC21_(_138_v7)
		})(), (_index24).Int())
		var _pat_let_tv11 = _138_v7
		_ = _pat_let_tv11
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_268_v119), 0))
		_ = _index25
		var _rhs36 bool = _133_v2
		_ = _rhs36
		var _rhs37 D9 = func(_pat_let7_0 D9) D9 {
			return func(_269_dt__update__tmp_h3 D9) D9 {
				return func(_pat_let8_0 _dafny.Array) D9 {
					return func(_270_dt__update_hcf33_h0 _dafny.Array) D9 {
						return Companion_D9_.Create_DC21_(_270_dt__update_hcf33_h0)
					}(_pat_let8_0)
				}(_pat_let_tv11)
			}(_pat_let7_0)
		}(Companion_D9_.Create_DC21_(_138_v7))
		_ = _rhs37
		var _rhs38 bool = _217_v71
		_ = _rhs38
		var _lhs28 *GlobalState = _151_globalState
		_ = _lhs28
		var _lhs29 _dafny.Array = _268_v119
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_268_v119), 0))
		_ = _lhs30
		_lhs28.F6 = _rhs36
		(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
		_133_v2 = _rhs38
	} else {
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))
		_ = _index26
		(_130_v0).ArraySet1(_217_v71, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))
		_ = _index27
		(_130_v0).ArraySet1(_217_v71, (_index27).Int())
		var _271_v120 _dafny.Map
		_ = _271_v120
		_271_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(892), (_201_v58).F34()), _216_v70)
		_271_v120 = (_271_v120).Merge(_271_v120)
		(_151_globalState).F5 = (_201_v58).F34()
		var _272_v121 D0
		_ = _272_v121
		_272_v121 = Companion_D0_.Create_DC2_((_222_v76).Intersection(_dafny.SetOf(_132_v1, (_201_v58).F34())), _148_v14)
		_272_v121 = _272_v121
		var _273_v122 _dafny.Array
		_ = _273_v122
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
		_ = _nw32
		_273_v122 = _nw32
		var _274_v123 _dafny.Map
		_ = _274_v123
		_274_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v31, _133_v2)
		var _275_v125 _dafny.MultiSet
		_ = _275_v125
		_275_v125 = _dafny.MultiSetOf(_dafny.SetOf(_217_v71), _167_v31)
		var _276_v129 _dafny.Map
		_ = _276_v129
		_276_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v31, (_201_v58).F34())
		var _277_v130 _dafny.Array
		_ = _277_v130
		var _nwElement0_3 _dafny.Map = _274_v123
		_ = _nwElement0_3
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(21))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_3, 0)
		(_nw33).ArraySet1(_274_v123, 1)
		(_nw33).ArraySet1(_274_v123, 2)
		(_nw33).ArraySet1(_274_v123, 3)
		(_nw33).ArraySet1(_274_v123, 4)
		(_nw33).ArraySet1((func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter30 := _dafny.Iterate((_275_v125).Elements()); ; {
				_compr_29, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _278_v124 _dafny.Set
				_278_v124 = interface{}(_compr_29).(_dafny.Set)
				if (_275_v125).Contains(_278_v124) {
					_coll29.Add(_278_v124, (_130_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))).Int()).(bool))
				}
			}
			return _coll29.ToMap()
		}()).Merge(_274_v123), 5)
		(_nw33).ArraySet1(Companion_Default___.Fm56(_217_v71, _218_v72, _151_globalState), 6)
		(_nw33).ArraySet1(_274_v123, 7)
		(_nw33).ArraySet1(_274_v123, 8)
		(_nw33).ArraySet1((_274_v123).Merge(_274_v123), 9)
		(_nw33).ArraySet1(func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter31 := _dafny.Iterate((_275_v125).Elements()); ; {
				_compr_30, _ok31 := _iter31()
				if !_ok31 {
					break
				}
				var _279_v126 _dafny.Set
				_279_v126 = interface{}(_compr_30).(_dafny.Set)
				if (_275_v125).Contains(_279_v126) {
					_coll30.Add(_279_v126, _133_v2)
				}
			}
			return _coll30.ToMap()
		}(), 10)
		(_nw33).ArraySet1(func() _dafny.Map {
			var _coll31 = _dafny.NewMapBuilder()
			_ = _coll31
			for _iter32 := _dafny.Iterate((_274_v123).Keys().Elements()); ; {
				_compr_31, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _280_v127 _dafny.Set
				_280_v127 = interface{}(_compr_31).(_dafny.Set)
				if (_274_v123).Contains(_280_v127) {
					_coll31.Add(_280_v127, (_145_v11).Select((Companion_Default___.SafeIndex(_219_v73, _dafny.IntOfUint32((_145_v11).Cardinality()))).Uint32()).(bool))
				}
			}
			return _coll31.ToMap()
		}(), 11)
		(_nw33).ArraySet1(_274_v123, 12)
		(_nw33).ArraySet1(_274_v123, 13)
		(_nw33).ArraySet1((_274_v123).Merge(func() _dafny.Map {
			var _coll32 = _dafny.NewMapBuilder()
			_ = _coll32
			for _iter33 := _dafny.Iterate((_276_v129).Keys().Elements()); ; {
				_compr_32, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _281_v128 _dafny.Set
				_281_v128 = interface{}(_compr_32).(_dafny.Set)
				if (_276_v129).Contains(_281_v128) {
					_coll32.Add(_281_v128, (_130_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))).Int()).(bool))
				}
			}
			return _coll32.ToMap()
		}()), 14)
		(_nw33).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v31, _133_v2)).Update(_dafny.SetOf((_130_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))).Int()).(bool)), true), 15)
		(_nw33).ArraySet1(_274_v123, 16)
		(_nw33).ArraySet1((Companion_Default___.Fm56(_216_v70, ((_262_v113).F36()).Select((Companion_Default___.SafeIndex((_199_v56).Dtor_cf7(), _dafny.IntOfUint32(((_262_v113).F36()).Cardinality()))).Uint32()).(_dafny.CodePoint), _151_globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_130_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))).Int()).(bool), _216_v70, _216_v70), true)), 17)
		(_nw33).ArraySet1(_274_v123, 18)
		(_nw33).ArraySet1(_274_v123, 19)
		(_nw33).ArraySet1(_274_v123, 20)
		_277_v130 = _nw33
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_277_v130), 0))
		_ = _index28
		(_277_v130).ArraySet1(_274_v123, (_index28).Int())
		var _282_v131 _dafny.MultiSet
		_ = _282_v131
		_282_v131 = _dafny.MultiSetOf(_132_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_148_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}((func(_283_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_284_i7 _dafny.Int) _dafny.CodePoint {
				return _283_v51
			}
		})(_194_v51))))).Cardinality()))
		var _285_v132 _dafny.Array
		_ = _285_v132
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw34
		_285_v132 = _nw34
		var _286_v133 _dafny.Map
		_ = _286_v133
		_286_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_285_v132, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-663)), (_201_v58).F34()))
		var _287_v134 D9
		_ = _287_v134
		_287_v134 = Companion_D9_.Create_DC22_(_139_v8, _133_v2, _285_v132, (_262_v113).F36())
		var _288_v135 _dafny.Sequence
		_ = _288_v135
		_288_v135 = _dafny.SeqOf(_274_v123, _274_v123)
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_277_v130), 0))
		_ = _index29
		var _rhs39 _dafny.Array = (func() _dafny.Array {
			if _217_v71 {
				return _273_v122
			}
			return _273_v122
		})()
		_ = _rhs39
		var _rhs40 _dafny.Int = (func() _dafny.Int {
			if (_286_v133).Contains((_287_v134).Dtor_cf36()) {
				return (_286_v133).Get((_287_v134).Dtor_cf36()).(_dafny.Int)
			}
			return _219_v73
		})()
		_ = _rhs40
		var _rhs41 _dafny.Int = _132_v1
		_ = _rhs41
		var _rhs42 _dafny.Map = ((_274_v123).Merge((_288_v135).Select((Companion_Default___.SafeIndex(_132_v1, _dafny.IntOfUint32((_288_v135).Cardinality()))).Uint32()).(_dafny.Map))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_216_v70), (_130_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_130_v0), 0))).Int()).(bool)))
		_ = _rhs42
		var _rhs43 _dafny.MultiSet = (_282_v131).Difference(_282_v131)
		_ = _rhs43
		var _lhs31 *GlobalState = _151_globalState
		_ = _lhs31
		var _lhs32 _dafny.Array = _277_v130
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_277_v130), 0))
		_ = _lhs33
		_273_v122 = _rhs39
		_lhs31.F13 = _rhs40
		_132_v1 = _rhs41
		(_lhs32).ArraySet1(_rhs42, (_lhs33).Int())
		_282_v131 = _rhs43
	}
	_dafny.Print((_130_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v0).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_132_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v3).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-900), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v8).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v8).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v8).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_143_v9).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_144_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_145_v11, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_146_v12, _dafny.SeqOf(_dafny.IntOfInt64(900))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_147_v13).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_148_v14.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_149_v15, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("y"), _dafny.UnicodeSeqOfUtf8Bytes("prvhys"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v16).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState.F9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-900), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(900))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_151_globalState).F16()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_151_globalState).F19(), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("y"), _dafny.UnicodeSeqOfUtf8Bytes("prvhys"), _dafny.UnicodeSeqOfUtf8Bytes("y"), _dafny.UnicodeSeqOfUtf8Bytes("prvhys"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_globalState).F22()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_151_globalState.F25)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_globalState).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(820))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v18).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(900)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v31).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_168_v32, _dafny.SeqOf(_dafny.SetOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_169_v33, _dafny.SeqOf(Companion_D5_.Create_DC12_(_dafny.SeqOf(_dafny.SetOf(true, false))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_170_v34, _dafny.SeqOf(_dafny.SeqOf(Companion_D5_.Create_DC12_(_dafny.SeqOf(_dafny.SetOf(true, false)))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_171_v35), _dafny.SeqOf(_dafny.SeqOf(Companion_D5_.Create_DC12_(_dafny.SeqOf(_dafny.SetOf(true, false)))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_194_v51)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v52).Dtor_cf31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_196_v53).Dtor_cf32()).Dtor_cf31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_197_v54).Dtor_cf32()).Dtor_cf31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_198_v55).Dtor_cf32()).Dtor_cf32()).Dtor_cf31())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v56).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v56).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v56).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_200_v57).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC3_(_dafny.One, false, false), _dafny.IntOfInt64(763))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_201_v58).F34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v58).F35()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC3_(_dafny.One, false, false), _dafny.IntOfInt64(763))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_202_v59).Dtor_cf40()).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v60).Equals(_dafny.SetOf(_dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_204_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v68).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(498), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_215_v69).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_216_v70)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_217_v71)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_218_v72)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_219_v73)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v74).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_221_v75).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v76).Equals(_dafny.SetOf(_dafny.IntOfInt64(-854))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v77).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(_dafny.IntOfInt64(-854)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v78).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(763), _dafny.SetOf(_dafny.IntOfInt64(-854)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_261_v112).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v113).F36().VerbatimString(false))
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
	Cf0 _dafny.Map
	Cf1 bool
	Cf2 _dafny.Int
	Cf3 _dafny.CodePoint
	Cf4 bool
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map, Cf1 bool, Cf2 _dafny.Int, Cf3 _dafny.CodePoint, Cf4 bool) D0 {
	return D0{D0_DC0{Cf0, Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 _dafny.Set
	Cf6 _dafny.Sequence
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 _dafny.Set, Cf6 _dafny.Sequence) D0 {
	return D0{D0_DC2{Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf7 _dafny.Int
	Cf8 bool
	Cf9 bool
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf7 _dafny.Int, Cf8 bool, Cf9 bool) D0 {
	return D0{D0_DC3{Cf7, Cf8, Cf9}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

type D0_DC4 struct {
	Cf10 D0
}

func (D0_DC4) isD0() {}

func (CompanionStruct_D0_) Create_DC4_(Cf10 D0) D0 {
	return D0{D0_DC4{Cf10}}
}

func (_this D0) Is_DC4() bool {
	_, ok := _this.Get_().(D0_DC4)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.EmptyMap, false, _dafny.Zero, _dafny.CodePoint('D'), false)
}

func (_this D0) Dtor_cf0() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf2
}

func (_this D0) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D0_DC0).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC0).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Set {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf7
}

func (_this D0) Dtor_cf8() bool {
	return _this.Get_().(D0_DC3).Cf8
}

func (_this D0) Dtor_cf9() bool {
	return _this.Get_().(D0_DC3).Cf9
}

func (_this D0) Dtor_cf10() D0 {
	return _this.Get_().(D0_DC4).Cf10
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ", " + data.Cf6.VerbatimString(true) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D0_DC4:
		{
			return "D0.DC4" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf0.Equals(data2.Cf0) && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4
		}
	case D0_DC1:
		{
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6.Equals(data2.Cf6)
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9
		}
	case D0_DC4:
		{
			data2, ok := other.Get_().(D0_DC4)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D1_DC6 struct {
	Cf12 _dafny.Int
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf12 _dafny.Int) D1 {
	return D1{D1_DC6{Cf12}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC7 struct {
	Cf13 _dafny.CodePoint
	Cf14 _dafny.Sequence
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_(Cf13 _dafny.CodePoint, Cf14 _dafny.Sequence) D1 {
	return D1{D1_DC7{Cf13, Cf14}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

type D1_DC5 struct {
	Cf11 _dafny.Sequence
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf11 _dafny.Sequence) D1 {
	return D1{D1_DC5{Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC6_(_dafny.Zero)
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) Dtor_cf13() _dafny.CodePoint {
	return _this.Get_().(D1_DC7).Cf13
}

func (_this D1) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D1_DC7).Cf14
}

func (_this D1) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7" + "(" + _dafny.String(data.Cf13) + ", " + data.Cf14.VerbatimString(true) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D1_DC7:
		{
			data2, ok := other.Get_().(D1_DC7)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Equals(data2.Cf14)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D2_DC8 struct {
	Cf15 _dafny.Sequence
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf15 _dafny.Sequence) D2 {
	return D2{D2_DC8{Cf15}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D2) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D3_DC10 struct {
	Cf17 bool
	Cf18 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf17 bool, Cf18 bool) D3 {
	return D3{D3_DC10{Cf17, Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

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

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, false)
}

func (_this D3) Dtor_cf17() bool {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
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
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18
		}
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
	Cf19 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf19 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf19}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf21 bool
	Cf22 _dafny.Array
	Cf23 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf21 bool, Cf22 _dafny.Array, Cf23 bool) D5 {
	return D5{D5_DC13{Cf21, Cf22, Cf23}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC12 struct {
	Cf20 _dafny.Sequence
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf20 _dafny.Sequence) D5 {
	return D5{D5_DC12{Cf20}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf22
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC13).Cf23
}

func (_this D5) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf20
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
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
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf20.Equals(data2.Cf20)
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
	Cf24 _dafny.Array
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf24 _dafny.Array) D6 {
	return D6{D6_DC14{Cf24}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D6) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D6_DC14).Cf24
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf24 == data2.Cf24
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

type D7_DC16 struct {
	Cf26 bool
	Cf27 _dafny.Map
	Cf28 _dafny.CodePoint
	Cf29 bool
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf26 bool, Cf27 _dafny.Map, Cf28 _dafny.CodePoint, Cf29 bool) D7 {
	return D7{D7_DC16{Cf26, Cf27, Cf28, Cf29}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC15 struct {
	Cf25 _dafny.Map
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf25 _dafny.Map) D7 {
	return D7{D7_DC15{Cf25}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC16_(false, _dafny.EmptyMap, _dafny.CodePoint('D'), false)
}

func (_this D7) Dtor_cf26() bool {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D7_DC16).Cf27
}

func (_this D7) Dtor_cf28() _dafny.CodePoint {
	return _this.Get_().(D7_DC16).Cf28
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC16).Cf29
}

func (_this D7) Dtor_cf25() _dafny.Map {
	return _this.Get_().(D7_DC15).Cf25
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27.Equals(data2.Cf27) && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf31 _dafny.Int
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf31 _dafny.Int) D8 {
	return D8{D8_DC18{Cf31}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

type D8_DC19 struct {
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_() D8 {
	return D8{D8_DC19{}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC17 struct {
	Cf30 _dafny.Map
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf30 _dafny.Map) D8 {
	return D8{D8_DC17{Cf30}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC20 struct {
	Cf32 D8
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf32 D8) D8 {
	return D8{D8_DC20{Cf32}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC18_(_dafny.Zero)
}

func (_this D8) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D8_DC18).Cf31
}

func (_this D8) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D8_DC17).Cf30
}

func (_this D8) Dtor_cf32() D8 {
	return _this.Get_().(D8_DC20).Cf32
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19"
		}
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf32) + ")"
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
			return ok && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D8_DC19:
		{
			_, ok := other.Get_().(D8_DC19)
			return ok
		}
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf30.Equals(data2.Cf30)
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D9_DC22 struct {
	Cf34 _dafny.Array
	Cf35 bool
	Cf36 _dafny.Array
	Cf37 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf34 _dafny.Array, Cf35 bool, Cf36 _dafny.Array, Cf37 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf34, Cf35, Cf36, Cf37}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC23 struct {
	Cf38 bool
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf38 bool) D9 {
	return D9{D9_DC23{Cf38}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC21 struct {
	Cf33 _dafny.Array
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf33 _dafny.Array) D9 {
	return D9{D9_DC21{Cf33}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

type D9_DC24 struct {
	Cf39 D9
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf39 D9) D9 {
	return D9{D9_DC24{Cf39}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC22_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq)
}

func (_this D9) Dtor_cf34() _dafny.Array {
	return _this.Get_().(D9_DC22).Cf34
}

func (_this D9) Dtor_cf35() bool {
	return _this.Get_().(D9_DC22).Cf35
}

func (_this D9) Dtor_cf36() _dafny.Array {
	return _this.Get_().(D9_DC22).Cf36
}

func (_this D9) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf37
}

func (_this D9) Dtor_cf38() bool {
	return _this.Get_().(D9_DC23).Cf38
}

func (_this D9) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D9_DC21).Cf33
}

func (_this D9) Dtor_cf39() D9 {
	return _this.Get_().(D9_DC24).Cf39
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + data.Cf37.VerbatimString(true) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36 && data1.Cf37.Equals(data2.Cf37)
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf38 == data2.Cf38
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf33 == data2.Cf33
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf41 bool
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf41 bool) D10 {
	return D10{D10_DC26{Cf41}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC25 struct {
	Cf40 _dafny.MultiSet
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf40 _dafny.MultiSet) D10 {
	return D10{D10_DC25{Cf40}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC26_(false)
}

func (_this D10) Dtor_cf41() bool {
	return _this.Get_().(D10_DC26).Cf41
}

func (_this D10) Dtor_cf40() _dafny.MultiSet {
	return _this.Get_().(D10_DC25).Cf40
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf40) + ")"
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
			return ok && data1.Cf41 == data2.Cf41
		}
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf40.Equals(data2.Cf40)
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
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_() D11 {
	return D11{D11_DC28{}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC27 struct {
	Cf42 _dafny.Array
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf42 _dafny.Array) D11 {
	return D11{D11_DC27{Cf42}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC28_()
}

func (_this D11) Dtor_cf42() _dafny.Array {
	return _this.Get_().(D11_DC27).Cf42
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC28:
		{
			return "D11.DC28"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf42) + ")"
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
			_, ok := other.Get_().(D11_DC28)
			return ok
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf42 == data2.Cf42
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
	Cf43 _dafny.Map
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf43 _dafny.Map) D12 {
	return D12{D12_DC29{Cf43}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D12) Dtor_cf43() _dafny.Map {
	return _this.Get_().(D12_DC29).Cf43
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf43) + ")"
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

type D13_DC30 struct {
	Cf44 _dafny.Sequence
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf44 _dafny.Sequence) D13 {
	return D13{D13_DC30{Cf44}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D13) Dtor_cf44() _dafny.Sequence {
	return _this.Get_().(D13_DC30).Cf44
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf44) + ")"
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

type D14_DC32 struct {
}

func (D14_DC32) isD14() {}

func (CompanionStruct_D14_) Create_DC32_() D14 {
	return D14{D14_DC32{}}
}

func (_this D14) Is_DC32() bool {
	_, ok := _this.Get_().(D14_DC32)
	return ok
}

type D14_DC33 struct {
	Cf46 _dafny.Int
	Cf47 _dafny.CodePoint
}

func (D14_DC33) isD14() {}

func (CompanionStruct_D14_) Create_DC33_(Cf46 _dafny.Int, Cf47 _dafny.CodePoint) D14 {
	return D14{D14_DC33{Cf46, Cf47}}
}

func (_this D14) Is_DC33() bool {
	_, ok := _this.Get_().(D14_DC33)
	return ok
}

type D14_DC31 struct {
	Cf45 _dafny.Sequence
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf45 _dafny.Sequence) D14 {
	return D14{D14_DC31{Cf45}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC32_()
}

func (_this D14) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D14_DC33).Cf46
}

func (_this D14) Dtor_cf47() _dafny.CodePoint {
	return _this.Get_().(D14_DC33).Cf47
}

func (_this D14) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D14_DC31).Cf45
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC32:
		{
			return "D14.DC32"
		}
	case D14_DC33:
		{
			return "D14.DC33" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf45) + ")"
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
			_, ok := other.Get_().(D14_DC32)
			return ok
		}
	case D14_DC33:
		{
			data2, ok := other.Get_().(D14_DC33)
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47 == data2.Cf47
		}
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D15_DC35 struct {
	Cf49 _dafny.Int
}

func (D15_DC35) isD15() {}

func (CompanionStruct_D15_) Create_DC35_(Cf49 _dafny.Int) D15 {
	return D15{D15_DC35{Cf49}}
}

func (_this D15) Is_DC35() bool {
	_, ok := _this.Get_().(D15_DC35)
	return ok
}

type D15_DC34 struct {
	Cf48 _dafny.Map
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_(Cf48 _dafny.Map) D15 {
	return D15{D15_DC34{Cf48}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC35_(_dafny.Zero)
}

func (_this D15) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D15_DC35).Cf49
}

func (_this D15) Dtor_cf48() _dafny.Map {
	return _this.Get_().(D15_DC34).Cf48
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC35:
		{
			return "D15.DC35" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D15_DC34:
		{
			return "D15.DC34" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC35:
		{
			data2, ok := other.Get_().(D15_DC35)
			return ok && data1.Cf49.Cmp(data2.Cf49) == 0
		}
	case D15_DC34:
		{
			data2, ok := other.Get_().(D15_DC34)
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

type D16_DC37 struct {
	Cf51 bool
	Cf52 bool
}

func (D16_DC37) isD16() {}

func (CompanionStruct_D16_) Create_DC37_(Cf51 bool, Cf52 bool) D16 {
	return D16{D16_DC37{Cf51, Cf52}}
}

func (_this D16) Is_DC37() bool {
	_, ok := _this.Get_().(D16_DC37)
	return ok
}

type D16_DC38 struct {
	Cf53 bool
	Cf54 bool
	Cf55 _dafny.Int
}

func (D16_DC38) isD16() {}

func (CompanionStruct_D16_) Create_DC38_(Cf53 bool, Cf54 bool, Cf55 _dafny.Int) D16 {
	return D16{D16_DC38{Cf53, Cf54, Cf55}}
}

func (_this D16) Is_DC38() bool {
	_, ok := _this.Get_().(D16_DC38)
	return ok
}

type D16_DC36 struct {
	Cf50 _dafny.MultiSet
}

func (D16_DC36) isD16() {}

func (CompanionStruct_D16_) Create_DC36_(Cf50 _dafny.MultiSet) D16 {
	return D16{D16_DC36{Cf50}}
}

func (_this D16) Is_DC36() bool {
	_, ok := _this.Get_().(D16_DC36)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC37_(false, false)
}

func (_this D16) Dtor_cf51() bool {
	return _this.Get_().(D16_DC37).Cf51
}

func (_this D16) Dtor_cf52() bool {
	return _this.Get_().(D16_DC37).Cf52
}

func (_this D16) Dtor_cf53() bool {
	return _this.Get_().(D16_DC38).Cf53
}

func (_this D16) Dtor_cf54() bool {
	return _this.Get_().(D16_DC38).Cf54
}

func (_this D16) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D16_DC38).Cf55
}

func (_this D16) Dtor_cf50() _dafny.MultiSet {
	return _this.Get_().(D16_DC36).Cf50
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC37:
		{
			return "D16.DC37" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D16_DC38:
		{
			return "D16.DC38" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D16_DC36:
		{
			return "D16.DC36" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC37:
		{
			data2, ok := other.Get_().(D16_DC37)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52 == data2.Cf52
		}
	case D16_DC38:
		{
			data2, ok := other.Get_().(D16_DC38)
			return ok && data1.Cf53 == data2.Cf53 && data1.Cf54 == data2.Cf54 && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D16_DC36:
		{
			data2, ok := other.Get_().(D16_DC36)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D17_DC39 struct {
	Cf56 _dafny.Map
}

func (D17_DC39) isD17() {}

func (CompanionStruct_D17_) Create_DC39_(Cf56 _dafny.Map) D17 {
	return D17{D17_DC39{Cf56}}
}

func (_this D17) Is_DC39() bool {
	_, ok := _this.Get_().(D17_DC39)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D17) Dtor_cf56() _dafny.Map {
	return _this.Get_().(D17_DC39).Cf56
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC39:
		{
			return "D17.DC39" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC39:
		{
			data2, ok := other.Get_().(D17_DC39)
			return ok && data1.Cf56.Equals(data2.Cf56)
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

type D18_DC41 struct {
	Cf58 bool
}

func (D18_DC41) isD18() {}

func (CompanionStruct_D18_) Create_DC41_(Cf58 bool) D18 {
	return D18{D18_DC41{Cf58}}
}

func (_this D18) Is_DC41() bool {
	_, ok := _this.Get_().(D18_DC41)
	return ok
}

type D18_DC40 struct {
	Cf57 _dafny.Map
}

func (D18_DC40) isD18() {}

func (CompanionStruct_D18_) Create_DC40_(Cf57 _dafny.Map) D18 {
	return D18{D18_DC40{Cf57}}
}

func (_this D18) Is_DC40() bool {
	_, ok := _this.Get_().(D18_DC40)
	return ok
}

type D18_DC42 struct {
	Cf59 D18
}

func (D18_DC42) isD18() {}

func (CompanionStruct_D18_) Create_DC42_(Cf59 D18) D18 {
	return D18{D18_DC42{Cf59}}
}

func (_this D18) Is_DC42() bool {
	_, ok := _this.Get_().(D18_DC42)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC41_(false)
}

func (_this D18) Dtor_cf58() bool {
	return _this.Get_().(D18_DC41).Cf58
}

func (_this D18) Dtor_cf57() _dafny.Map {
	return _this.Get_().(D18_DC40).Cf57
}

func (_this D18) Dtor_cf59() D18 {
	return _this.Get_().(D18_DC42).Cf59
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC41:
		{
			return "D18.DC41" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D18_DC40:
		{
			return "D18.DC40" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D18_DC42:
		{
			return "D18.DC42" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC41:
		{
			data2, ok := other.Get_().(D18_DC41)
			return ok && data1.Cf58 == data2.Cf58
		}
	case D18_DC40:
		{
			data2, ok := other.Get_().(D18_DC40)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	case D18_DC42:
		{
			data2, ok := other.Get_().(D18_DC42)
			return ok && data1.Cf59.Equals(data2.Cf59)
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
	F27() _dafny.Array
	F28() bool
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
	F27() _dafny.Array
	F28() bool
	F29() _dafny.Int
	F29_set_(value _dafny.Int)
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
	F4   _dafny.Int
	F5   _dafny.Int
	F6   bool
	F9   _dafny.Map
	F13  _dafny.Int
	F15  bool
	F17  bool
	F18  _dafny.Array
	F21  _dafny.Int
	F25  _dafny.CodePoint
	_f0  bool
	_f1  bool
	_f2  _dafny.Int
	_f3  bool
	_f7  _dafny.Int
	_f8  _dafny.Array
	_f10 _dafny.MultiSet
	_f11 _dafny.Array
	_f12 _dafny.Int
	_f14 _dafny.Array
	_f16 _dafny.Sequence
	_f19 _dafny.Sequence
	_f20 _dafny.Int
	_f22 _dafny.Map
	_f23 _dafny.Int
	_f24 _dafny.Int
	_f26 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.Zero
	_this.F5 = _dafny.Zero
	_this.F6 = false
	_this.F9 = _dafny.EmptyMap
	_this.F13 = _dafny.Zero
	_this.F15 = false
	_this.F17 = false
	_this.F18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F21 = _dafny.Zero
	_this.F25 = _dafny.CodePoint('D')
	_this._f0 = false
	_this._f1 = false
	_this._f2 = _dafny.Zero
	_this._f3 = false
	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f10 = _dafny.EmptyMultiSet
	_this._f11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f12 = _dafny.Zero
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f16 = _dafny.EmptySeq
	_this._f19 = _dafny.EmptySeq
	_this._f20 = _dafny.Zero
	_this._f22 = _dafny.EmptyMap
	_this._f23 = _dafny.Zero
	_this._f24 = _dafny.Zero
	_this._f26 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 _dafny.Int, f3 bool, f4 _dafny.Int, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 _dafny.Array, f9 _dafny.Map, f10 _dafny.MultiSet, f11 _dafny.Array, f12 _dafny.Int, f13 _dafny.Int, f14 _dafny.Array, f15 bool, f16 _dafny.Sequence, f17 bool, f18 _dafny.Array, f19 _dafny.Sequence, f20 _dafny.Int, f21 _dafny.Int, f22 _dafny.Map, f23 _dafny.Int, f24 _dafny.Int, f25 _dafny.CodePoint, f26 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this).F17 = f17
		(_this).F18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this).F21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this).F25 = f25
		(_this)._f26 = f26
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
func (_this *GlobalState) F8() _dafny.Array {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F10() _dafny.MultiSet {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Array {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Array {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() _dafny.Sequence {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F19() _dafny.Sequence {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() _dafny.Int {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F22() _dafny.Map {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Int {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() _dafny.Int {
	{
		return _this._f24
	}
}
func (_this *GlobalState) F26() _dafny.Int {
	{
		return _this._f26
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f29 _dafny.Int
	_f27 _dafny.Array
	_f28 bool
	_f36 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f29 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
	_this._f36 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C0{}
var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F29() _dafny.Int {
	return _this._f29
}
func (_this *C0) F29_set_(value _dafny.Int) {
	_this._f29 = value
}
func (_this *C0) F27() _dafny.Array {
	return _this._f27
}
func (_this *C0) F28() bool {
	return _this._f28
}
func (_this *C0) Ctor__(f36 _dafny.Sequence, f29 _dafny.Int, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f36 = f36
		(_this)._f29 = f29
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C0) Fm15(globalState *GlobalState) _dafny.Int {
	{
		return (_this.F29()).Plus((_dafny.Zero).Minus(_this.F29()))
	}
}
func (_this *C0) F36() _dafny.Sequence {
	{
		return _this._f36
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f27 _dafny.Array
	_f28 bool
	_f39 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F27() _dafny.Array {
	return _this._f27
}
func (_this *C1) F28() bool {
	return _this._f28
}
func (_this *C1) Ctor__(f39 bool, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f39 = f39
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C1) Fm34(p0 bool, p1 _dafny.Int, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("rhictu")
	}
}
func (_this *C1) Fm35(p0 bool, p1 bool, globalState *GlobalState) D3 {
	{
		return Companion_D3_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(801), _dafny.SetOf(_dafny.IntOfInt64(527), _dafny.IntOfInt64(65), _dafny.IntOfInt64(543))))
	}
}
func (_this *C1) M14(p0 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Sequence, _dafny.Sequence, _dafny.Map) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Map = _dafny.EmptyMap
		_ = r3
		var _289_v0 _dafny.Sequence
		_ = _289_v0
		_289_v0 = _dafny.UnicodeSeqOfUtf8Bytes("usplmirjv")
		(globalState).F13 = _dafny.IntOfUint32((_289_v0).Cardinality())
		var _290_v1 _dafny.Int
		_ = _290_v1
		_290_v1 = _dafny.IntOfInt64(511)
		var _291_v2 _dafny.MultiSet
		_ = _291_v2
		_291_v2 = _dafny.MultiSetOf(_290_v1, _290_v1, _290_v1)
		var _292_v3 _dafny.Sequence
		_ = _292_v3
		_292_v3 = _dafny.SeqOf(_290_v1)
		if ((_dafny.MultiSetFromSeq(_292_v3)).Difference(_291_v2)).IsSubsetOf((_291_v2).Update(_290_v1, Companion_Default___.Abs(_290_v1))) {
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index30
			((_this).F27()).ArraySet1((_this).F39(), (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index31
			((_this).F27()).ArraySet1(p0, (_index31).Int())
			var _293_v4 _dafny.Map
			_ = _293_v4
			_293_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_291_v2).Cardinality()), _dafny.IntOfUint32((_289_v0).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("jj"))
			_293_v4 = _293_v4
			var _294_v5 _dafny.Sequence
			_ = _294_v5
			_294_v5 = _dafny.SeqOf(_289_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(231))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_295_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})))
			var _296_v6 _dafny.Array
			_ = _296_v6
			var _nwElement0_4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_294_v5).Select((Companion_Default___.SafeIndex(_290_v1, _dafny.IntOfUint32((_294_v5).Cardinality()))).Uint32()).(_dafny.Sequence), _289_v0)
			_ = _nwElement0_4
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(4))
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_4, 0)
			(_nw35).ArraySet1(_289_v0, 1)
			(_nw35).ArraySet1(_289_v0, 2)
			(_nw35).ArraySet1(_289_v0, 3)
			_296_v6 = _nw35
			var _297_v7 D9
			_ = _297_v7
			_297_v7 = Companion_D9_.Create_DC21_(_296_v6)
			_296_v6 = (_297_v7).Dtor_cf33()
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index32
			((_this).F27()).ArraySet1((_290_v1).Cmp(Companion_Default___.Fm36(globalState)) <= 0, (_index32).Int())
			var _298_v8 _dafny.Sequence
			_ = _298_v8
			_298_v8 = _dafny.SeqOf(p0, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool))
			(globalState).F5 = Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_298_v8).Cardinality())).Minus(_dafny.IntOfUint32((_292_v3).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_289_v0, _289_v0)).Cardinality()))
		} else {
			var _299_v9 D8
			_ = _299_v9
			_299_v9 = Companion_D8_.Create_DC18_(_290_v1)
			var _300_v10 _dafny.Set
			_ = _300_v10
			_300_v10 = _dafny.SetOf(_299_v9, _299_v9, _299_v9, Companion_D8_.Create_DC18_(_290_v1), _299_v9)
			_300_v10 = (_300_v10).Intersection((_300_v10).Difference(_300_v10))
			var _301_v11 _dafny.Set
			_ = _301_v11
			_301_v11 = _dafny.SetOf(false)
			(globalState).F5 = (_dafny.IntOfInt64(783)).Times((_301_v11).Cardinality())
			(globalState).F5 = (_290_v1).Times(_290_v1)
			(globalState).F13 = (_290_v1).Times((_290_v1).Plus(_290_v1))
			(globalState).F6 = p0
		}
		(globalState).F15 = (_this).F39()
		var _302_v12 *C0
		_ = _302_v12
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__(_289_v0, _290_v1, (_this).F27(), p0)
		_302_v12 = _nw36
		(globalState).F15 = !(!(p0))
		var _303_v13 _dafny.MultiSet
		_ = _303_v13
		_303_v13 = _dafny.MultiSetOf((_this).F39())
		var _304_v14 D10
		_ = _304_v14
		_304_v14 = Companion_D10_.Create_DC25_(_303_v13)
		var _hi0 _dafny.Int = (_dafny.MultiSetOf((func() _dafny.Int {
			if (_303_v13).Contains((_this).F39()) {
				return (_303_v13).Multiplicity((_this).F39())
			}
			return _290_v1
		})())).Cardinality()
		_ = _hi0
		for _305_i1 := ((_304_v14).Dtor_cf40()).Cardinality(); _305_i1.Cmp(_hi0) < 0; _305_i1 = _305_i1.Plus(_dafny.One) {
			var _306_v16 _dafny.Array
			_ = _306_v16
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_3
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Map = (func(_307_i1 _dafny.Int, _308_v1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_309_i2 _dafny.Int) _dafny.Map {
						return func() _dafny.Map {
							var _coll33 = _dafny.NewMapBuilder()
							_ = _coll33
							for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-710), _dafny.IntOfInt64(610))); ; {
								_compr_33, _ok34 := _iter34()
								if !_ok34 {
									break
								}
								var _310_v15 _dafny.Int
								_310_v15 = interface{}(_compr_33).(_dafny.Int)
								if ((_dafny.IntOfInt64(-710)).Cmp(_310_v15) <= 0) && ((_310_v15).Cmp(_dafny.IntOfInt64(610)) < 0) {
									_coll33.Add((_310_v15).Times((_dafny.SetOf(_307_i1, _dafny.IntOfInt64(954))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v1, _dafny.IntOfInt64(947)))
								}
							}
							return _coll33.ToMap()
						}()
					}
				})(_305_i1, _290_v1)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw37 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw37).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw37).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_306_v16 = _nw37
			var _311_v18 _dafny.Map
			_ = _311_v18
			_311_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_i1, ((func() _dafny.Map {
				var _coll34 = _dafny.NewMapBuilder()
				_ = _coll34
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-859), _dafny.IntOfInt64(455))); ; {
					_compr_34, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _312_v17 _dafny.Int
					_312_v17 = interface{}(_compr_34).(_dafny.Int)
					if ((_dafny.IntOfInt64(-859)).Cmp(_312_v17) <= 0) && ((_312_v17).Cmp(_dafny.IntOfInt64(455)) < 0) {
						_coll34.Add(Companion_Default___.SafeModuloInt(_312_v17, _305_i1), _305_i1)
					}
				}
				return _coll34.ToMap()
			}()).Update(_290_v1, _290_v1)).Update((func() _dafny.Int {
				if (_291_v2).Contains(_305_i1) {
					return (_291_v2).Multiplicity(_305_i1)
				}
				return _290_v1
			})(), _305_i1))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_306_v16), 0))
			_ = _index33
			(_306_v16).ArraySet1(_311_v18, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_306_v16), 0))
			_ = _index34
			(_306_v16).ArraySet1(_311_v18, (_index34).Int())
			(globalState).F5 = (func() _dafny.Int {
				if (_this).F39() {
					return (_303_v13).Cardinality()
				}
				return _290_v1
			})()
			var _313_v19 _dafny.Array
			_ = _313_v19
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_4
			var _nw38 _dafny.Array
			_ = _nw38
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw38 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Int = (func(_314_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_315_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_315_i3, _314_i1)
					}
				})(_305_i1)
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
			_313_v19 = _nw38
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_313_v19), 0))
			_ = _index35
			(_313_v19).ArraySet1(_290_v1, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_313_v19), 0))
			_ = _index36
			(_313_v19).ArraySet1((Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_292_v3).Cardinality()), _290_v1)).Plus(_290_v1), (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_313_v19), 0))
			_ = _index37
			(_313_v19).ArraySet1((_313_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_313_v19), 0))).Int()).(_dafny.Int), (_index37).Int())
		}
		r0 = _dafny.Companion_Sequence_.Concatenate((_302_v12).F36(), _289_v0)
		var _316_v20 D9
		_ = _316_v20
		_316_v20 = Companion_D9_.Create_DC23_(p0)
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _pat_let_tv13 = p0
		_ = _pat_let_tv13
		var _pat_let_tv14 = _290_v1
		_ = _pat_let_tv14
		var _pat_let_tv15 = p0
		_ = _pat_let_tv15
		var _pat_let_tv16 = p0
		_ = _pat_let_tv16
		var _pat_let_tv17 = p0
		_ = _pat_let_tv17
		var _pat_let_tv18 = p0
		_ = _pat_let_tv18
		r1 = func(_source6 D9) _dafny.Sequence {
			if _source6.Is_DC22() {
				var _317___mcc_h0 _dafny.Array = _source6.Get_().(D9_DC22).Cf34
				_ = _317___mcc_h0
				var _318___mcc_h1 bool = _source6.Get_().(D9_DC22).Cf35
				_ = _318___mcc_h1
				var _319___mcc_h2 _dafny.Array = _source6.Get_().(D9_DC22).Cf36
				_ = _319___mcc_h2
				var _320___mcc_h3 _dafny.Sequence = _source6.Get_().(D9_DC22).Cf37
				_ = _320___mcc_h3
				var _321_cf37 _dafny.Sequence = _320___mcc_h3
				_ = _321_cf37
				var _322_cf36 _dafny.Array = _319___mcc_h2
				_ = _322_cf36
				var _323_cf35 bool = _318___mcc_h1
				_ = _323_cf35
				var _324_cf34 _dafny.Array = _317___mcc_h0
				_ = _324_cf34
				return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv12, _pat_let_tv13), _dafny.SeqOf((_this).F28())), (Companion_Default___.SafeIndex(_pat_let_tv14, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv15, _pat_let_tv16), _dafny.SeqOf((_this).F28()))).Cardinality()))).Uint32(), (_this).F28())
			} else if _source6.Is_DC23() {
				var _325___mcc_h4 bool = _source6.Get_().(D9_DC23).Cf38
				_ = _325___mcc_h4
				var _326_cf38 bool = _325___mcc_h4
				_ = _326_cf38
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv17, _326_cf38), _dafny.SeqOf(_326_cf38, (_this).F39()))
			} else if _source6.Is_DC21() {
				var _327___mcc_h5 _dafny.Array = _source6.Get_().(D9_DC21).Cf33
				_ = _327___mcc_h5
				var _328_cf33 _dafny.Array = _327___mcc_h5
				_ = _328_cf33
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F28()), _dafny.SeqOf(false))
			} else {
				var _329___mcc_h6 D9 = _source6.Get_().(D9_DC24).Cf39
				_ = _329___mcc_h6
				var _330_cf39 D9 = _329___mcc_h6
				_ = _330_cf39
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv18), _dafny.SeqOf((_this).F28()))
			}
		}(_316_v20)
		var _331_v21 _dafny.Sequence
		_ = _331_v21
		_331_v21 = _dafny.SeqOf((_302_v12).F36())
		r2 = (_331_v21).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(Companion_Default___.Fm36(globalState), _290_v1), _dafny.IntOfUint32((_331_v21).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _332_v22 _dafny.Sequence
		_ = _332_v22
		_332_v22 = _dafny.SeqOf((_this).F28())
		var _333_v23 _dafny.Sequence
		_ = _333_v23
		_333_v23 = _dafny.SeqOf(_332_v22, _332_v22)
		r3 = Companion_Default___.Fm37(p0, false, _290_v1, _dafny.IntOfUint32(((_333_v23).Select((Companion_Default___.SafeIndex(_290_v1, _dafny.IntOfUint32((_333_v23).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), globalState)
		return r0, r1, r2, r3
	}
}
func (_this *C1) M15(p0 _dafny.Set, p1 _dafny.Int, p2 bool, globalState *GlobalState) (_dafny.MultiSet, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 bool = false
		_ = r1
		var _334_v0 _dafny.Array
		_ = _334_v0
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
		_ = _nw39
		_334_v0 = _nw39
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_334_v0), 0))
		_ = _index38
		(_334_v0).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p1), (_index38).Int())
		var _335_v1 _dafny.Set
		_ = _335_v1
		_335_v1 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ubsto"))
		var _336_v2 _dafny.Map
		_ = _336_v2
		_336_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_this).F39())), (p1).Plus((_dafny.Zero).Minus((_335_v1).Cardinality())))
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_334_v0), 0))
		_ = _index39
		(_334_v0).ArraySet1(_336_v2, (_index39).Int())
		var _337_v3 _dafny.Array
		_ = _337_v3
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
		_ = _nw40
		_337_v3 = _nw40
		var _338_v4 D11
		_ = _338_v4
		_338_v4 = Companion_D11_.Create_DC27_(_337_v3)
		var _pat_let_tv19 = _337_v3
		_ = _pat_let_tv19
		var _339_v5 _dafny.Array
		_ = _339_v5
		var _nwElement0_5 _dafny.Array = _337_v3
		_ = _nwElement0_5
		var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(11))
		_ = _nw41
		(_nw41).ArraySet1(_nwElement0_5, 0)
		(_nw41).ArraySet1(_337_v3, 1)
		(_nw41).ArraySet1(_337_v3, 2)
		(_nw41).ArraySet1(_337_v3, 3)
		(_nw41).ArraySet1(_337_v3, 4)
		(_nw41).ArraySet1((func(_pat_let9_0 D11) D11 {
			return func(_340_dt__update__tmp_h0 D11) D11 {
				return func(_pat_let10_0 _dafny.Array) D11 {
					return func(_341_dt__update_hcf42_h0 _dafny.Array) D11 {
						return Companion_D11_.Create_DC27_(_341_dt__update_hcf42_h0)
					}(_pat_let10_0)
				}(_pat_let_tv19)
			}(_pat_let9_0)
		}(_338_v4)).Dtor_cf42(), 5)
		(_nw41).ArraySet1(_337_v3, 6)
		(_nw41).ArraySet1(_337_v3, 7)
		(_nw41).ArraySet1(_337_v3, 8)
		(_nw41).ArraySet1(_337_v3, 9)
		(_nw41).ArraySet1(_337_v3, 10)
		_339_v5 = _nw41
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_339_v5), 0))
		_ = _index40
		(_339_v5).ArraySet1(_337_v3, (_index40).Int())
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_339_v5), 0))
		_ = _index41
		(_339_v5).ArraySet1(_337_v3, (_index41).Int())
		var _342_v6 _dafny.Sequence
		_ = _342_v6
		_342_v6 = _dafny.SeqOf((_this).F27(), (_this).F27(), (_this).F27())
		var _343_v7 _dafny.Map
		_ = _343_v7
		_343_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SeqOf((_this).F27()))
		_342_v6 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_343_v7).Contains(p1) {
				return (_343_v7).Get(p1).(_dafny.Sequence)
			}
			return _342_v6
		})(), _dafny.SeqOf((_this).F27()))
		if (_this).F28() {
			(globalState).F6 = (p1).Cmp(p1) == 0
			var _344_v8 _dafny.Sequence
			_ = _344_v8
			_344_v8 = _dafny.SeqOf(p1, p1)
			var _345_v9 _dafny.Sequence
			_ = _345_v9
			_345_v9 = _dafny.Companion_Sequence_.Concatenate(_344_v8, _344_v8)
			_345_v9 = _345_v9
			(globalState).F15 = ((p0).Intersection(p0)).IsProperSubsetOf((p0).Union(p0))
			var _346_v10 _dafny.Array
			_ = _346_v10
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
			_ = _nw42
			_346_v10 = _nw42
			var _347_v11 _dafny.Map
			_ = _347_v11
			_347_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_346_v10), 0))
			_ = _index42
			(_346_v10).ArraySet1(_347_v11, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_346_v10), 0))
			_ = _index43
			(_346_v10).ArraySet1((_347_v11).Merge((_347_v11).Merge(_347_v11)), (_index43).Int())
			var _348_v12 _dafny.Array
			_ = _348_v12
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw43
			_348_v12 = _nw43
			var _349_v13 _dafny.Sequence
			_ = _349_v13
			_349_v13 = _dafny.SeqOf(_348_v12)
			var _350_v14 _dafny.Sequence
			_ = _350_v14
			_350_v14 = _dafny.UnicodeSeqOfUtf8Bytes("xuh")
			var _351_v15 _dafny.Map
			_ = _351_v15
			_351_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(p1)).Times(p1), (_dafny.IntOfUint32((_349_v13).Cardinality())).Cmp(_dafny.IntOfUint32((_350_v14).Cardinality())) > 0)
			var _352_v16 _dafny.Sequence
			_ = _352_v16
			_352_v16 = _dafny.SeqOf((_this).F28(), (_this).F39())
			var _353_v17 _dafny.MultiSet
			_ = _353_v17
			_353_v17 = _dafny.MultiSetOf(p0, Companion_Default___.Fm38(p1, !((_this).F28()), (_this).F28(), (_this).F39(), globalState))
			_351_v15 = (_351_v15).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_352_v16, _352_v16)).Cardinality()), (_353_v17).Equals(_353_v17))
		} else {
			var _354_v18 _dafny.Array
			_ = _354_v18
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw44
			_354_v18 = _nw44
			var _355_v19 _dafny.MultiSet
			_ = _355_v19
			_355_v19 = _dafny.MultiSetOf(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _dafny.IntOfInt64(-418))).Cardinality())
			var _356_v20 _dafny.Set
			_ = _356_v20
			_356_v20 = _dafny.SetOf(_355_v19, _dafny.MultiSetOf(p1, p1), _355_v19)
			var _357_v21 _dafny.Map
			_ = _357_v21
			_357_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v18, _356_v20)
			_357_v21 = (_357_v21).Update(_354_v18, (_356_v20).Union(_356_v20))
			if (_this).F28() {
				var _358_v22 _dafny.Array
				_ = _358_v22
				var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw45
				_358_v22 = _nw45
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))
				_ = _index44
				(_358_v22).ArraySet1(_354_v18, (_index44).Int())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))
				_ = _index45
				(_358_v22).ArraySet1(_354_v18, (_index45).Int())
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))
				_ = _arr0
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))), 0))
				_ = _index46
				(_arr0).ArraySet1(p1, (_index46).Int())
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))
				_ = _arr1
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))), 0))
				_ = _index47
				(_arr1).ArraySet1(((_dafny.SetOf(p1, (func() _dafny.Map {
					var _coll35 = _dafny.NewMapBuilder()
					_ = _coll35
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(580), _dafny.IntOfInt64(746))); ; {
						_compr_35, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _359_v23 _dafny.Int
						_359_v23 = interface{}(_compr_35).(_dafny.Int)
						if ((_dafny.IntOfInt64(580)).Cmp(_359_v23) <= 0) && ((_359_v23).Cmp(_dafny.IntOfInt64(746)) < 0) {
							_coll35.Add(Companion_Default___.SafeModuloInt(_359_v23, _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality())), (_this).F39())
						}
					}
					return _coll35.ToMap()
				}()).Cardinality(), p1, (_dafny.Zero).Minus(p1))).Cardinality()).Times(p1), (_index47).Int())
				var _360_v24 _dafny.Sequence
				_ = _360_v24
				_360_v24 = _dafny.UnicodeSeqOfUtf8Bytes("txp")
				var _361_v25 _dafny.Sequence
				_ = _361_v25
				_361_v25 = _dafny.SeqOf((_this).F39())
				_360_v24 = (_this).Fm34((false) && (!((_this).F28())), _dafny.IntOfInt64(222), _dafny.MultiSetOf((_this).F39(), (_361_v25).Select((Companion_Default___.SafeIndex((_dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_361_v25).Cardinality()))).Uint32()).(bool), p2, (_this).F28(), (_this).F28()), (_dafny.Zero).Minus((_dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_dafny.ArrayCastTo((_358_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_358_v22), 0))).Int()))), 0))).Int()).(_dafny.Int)), globalState)
				_360_v24 = _dafny.Companion_Sequence_.Concatenate(_360_v24, _360_v24)
				var _362_v26 _dafny.Map
				_ = _362_v26
				_362_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F28())
				var _363_v27 _dafny.Sequence
				_ = _363_v27
				_363_v27 = _dafny.SeqOf(_355_v19, (_355_v19).Update(p1, Companion_Default___.Abs((_362_v26).Cardinality())), _355_v19)
				_363_v27 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-501))).Uint32(), func(coer37 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_364_v19 _dafny.MultiSet, _365_p1 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_366_i0 _dafny.Int) _dafny.MultiSet {
						return (_364_v19).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(159), _dafny.MultiSetOf((_this).F39()))).Cardinality(), Companion_Default___.Abs(_365_p1))
					}
				})(_355_v19, p1))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(453))).Uint32(), func(coer38 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_367_v19 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_368_i1 _dafny.Int) _dafny.MultiSet {
						return _367_v19
					}
				})(_355_v19))))
			} else {
				var _369_v28 _dafny.Array
				_ = _369_v28
				var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
				_ = _nw46
				_369_v28 = _nw46
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))
				_ = _index48
				(_354_v18).ArraySet1(p1, (_index48).Int())
				var _370_v29 _dafny.Map
				_ = _370_v29
				_370_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F39()), _354_v18)
				var _371_v30 _dafny.Map
				_ = _371_v30
				_371_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))
				_ = _index49
				var _rhs44 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, p1)
				_ = _rhs44
				var _rhs45 _dafny.Array = _369_v28
				_ = _rhs45
				var _rhs46 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.Map {
					if p2 {
						return _370_v29
					}
					return (_370_v29).Merge(_370_v29)
				})()).Cardinality())
				_ = _rhs46
				var _rhs47 bool = (func() bool {
					if (_371_v30).Contains(!(p2) || ((_this).F39())) {
						return (_371_v30).Get(!(p2) || ((_this).F39())).(bool)
					}
					return p2
				})()
				_ = _rhs47
				var _rhs48 bool = p2
				_ = _rhs48
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				var _lhs35 _dafny.Array = _354_v18
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))
				_ = _lhs36
				var _lhs37 *GlobalState = globalState
				_ = _lhs37
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				_lhs34.F5 = _rhs44
				_369_v28 = _rhs45
				(_lhs35).ArraySet1(_rhs46, (_lhs36).Int())
				_lhs37.F15 = _rhs47
				_lhs38.F17 = _rhs48
				var _372_v31 _dafny.MultiSet
				_ = _372_v31
				_372_v31 = _dafny.MultiSetOf((_this).F39())
				var _373_v32 _dafny.Sequence
				_ = _373_v32
				_373_v32 = _dafny.UnicodeSeqOfUtf8Bytes("i")
				var _374_v33 _dafny.Map
				_ = _374_v33
				_374_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_354_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))).Int()).(_dafny.Int))
				var _375_v34 _dafny.Map
				_ = _375_v34
				_375_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_374_v33).Contains(_dafny.SetOf(p2)) {
						return (_374_v33).Get(_dafny.SetOf(p2)).(_dafny.Int)
					}
					return p1
				})(), (_this).F39())
				var _376_v35 _dafny.CodePoint
				_ = _376_v35
				_376_v35 = _dafny.CodePoint('i')
				var _377_v36 *C0
				_ = _377_v36
				var _nw47 *C0 = New_C0_()
				_ = _nw47
				_nw47.Ctor__((_this).Fm34((_this).F39(), (_354_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))).Int()).(_dafny.Int), _372_v31, p1, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_373_v32, (Companion_Default___.SafeIndex((_375_v34).Cardinality(), _dafny.IntOfUint32((_373_v32).Cardinality()))).Uint32(), _376_v35)).Cardinality()), (_this).F27(), !(p2))
				_377_v36 = _nw47
				var _378_v37 _dafny.Array
				_ = _378_v37
				var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw48
				_378_v37 = _nw48
				var _379_v38 _dafny.Map
				_ = _379_v38
				_379_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v37, p2)
				_379_v38 = (_379_v38).Update(_378_v37, (_this).F28())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index50
				((_this).F27()).ArraySet1((_this).F39(), (_index50).Int())
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index51
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))
				_ = _index52
				var _rhs49 _dafny.CodePoint = (_373_v32).Select((Companion_Default___.SafeIndex((_354_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_373_v32).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs49
				var _rhs50 bool = (func() bool {
					if (_this).F28() {
						return p2
					}
					return !((_this).F28())
				})()
				_ = _rhs50
				var _rhs51 bool = (_this).F28()
				_ = _rhs51
				var _rhs52 _dafny.Int = (p1).Minus((func() _dafny.Int {
					if (_this).F28() {
						return (_dafny.Zero).Minus(Companion_Default___.Fm36(globalState))
					}
					return p1
				})())
				_ = _rhs52
				var _lhs39 _dafny.Array = (_this).F27()
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _lhs40
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 _dafny.Array = _354_v18
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))
				_ = _lhs43
				_376_v35 = _rhs49
				(_lhs39).ArraySet1(_rhs50, (_lhs40).Int())
				_lhs41.F6 = _rhs51
				(_lhs42).ArraySet1(_rhs52, (_lhs43).Int())
				var _380_v39 _dafny.Array
				_ = _380_v39
				var _nwElement0_6 _dafny.Int = p1
				_ = _nwElement0_6
				var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
				_ = _nw49
				(_nw49).ArraySet1(_nwElement0_6, 0)
				_380_v39 = _nw49
				var _381_v40 _dafny.Map
				_ = _381_v40
				_381_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_354_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_354_v18), 0))).Int()).(_dafny.Int)), _376_v35)
				var _382_v41 _dafny.Sequence
				_ = _382_v41
				_382_v41 = _dafny.SeqOf(p1, p1, (_381_v40).Cardinality())
				var _383_v42 _dafny.Map
				_ = _383_v42
				_383_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_382_v41, _380_v39)
				var _384_v43 _dafny.Sequence
				_ = _384_v43
				_384_v43 = _dafny.SeqOf(_380_v39, _380_v39, _380_v39)
				var _385_v44 _dafny.Set
				_ = _385_v44
				_385_v44 = _dafny.SetOf(_dafny.CodePoint('k'), _376_v35)
				_380_v39 = (func() _dafny.Array {
					if (_383_v42).Contains(_382_v41) {
						return (_383_v42).Get(_382_v41).(_dafny.Array)
					}
					return (_384_v43).Select((Companion_Default___.SafeIndex((_385_v44).Cardinality(), _dafny.IntOfUint32((_384_v43).Cardinality()))).Uint32()).(_dafny.Array)
				})()
			}
			(globalState).F21 = p1
			var _386_v45 _dafny.MultiSet
			_ = _386_v45
			_386_v45 = _dafny.MultiSetOf((_this).F39())
			if (p1).Cmp((func() _dafny.Int {
				if (_this).F28() {
					return (func() _dafny.Int {
						if (_386_v45).Contains(p2) {
							return (_386_v45).Multiplicity(p2)
						}
						return p1
					})()
				}
				return p1
			})()) != 0 {
				var _387_v46 _dafny.Sequence
				_ = _387_v46
				_387_v46 = _dafny.UnicodeSeqOfUtf8Bytes("qouh")
				var _388_v47 _dafny.CodePoint
				_ = _388_v47
				_388_v47 = _dafny.CodePoint('c')
				var _389_v48 *C0
				_ = _389_v48
				var _nw50 *C0 = New_C0_()
				_ = _nw50
				_nw50.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nulmgbv"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_387_v46).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nulmgbv")).Cardinality()))).Uint32(), _388_v47), _387_v46), p1, (_this).F27(), !((_this).F28()) || ((_this).F28()))
				_389_v48 = _nw50
				var _390_v49 D0
				_ = _390_v49
				_390_v49 = Companion_D0_.Create_DC1_()
				var _391_v50 D0
				_ = _391_v50
				_391_v50 = Companion_D0_.Create_DC4_(_390_v49)
				_391_v50 = _391_v50
				(globalState).F21 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-95), p1))
				var _392_v51 _dafny.Sequence
				_ = _392_v51
				_392_v51 = _dafny.SeqOf(p2, (_this).F28())
				var _393_v52 _dafny.Map
				_ = _393_v52
				_393_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_392_v51).Cardinality()), (_389_v48).F36())
				var _394_v53 D7
				_ = _394_v53
				_394_v53 = Companion_D7_.Create_DC15_(_393_v52)
				var _395_v54 _dafny.Sequence
				_ = _395_v54
				_395_v54 = _dafny.SeqOf(_394_v53)
				_395_v54 = _dafny.Companion_Sequence_.Concatenate(_395_v54, _dafny.Companion_Sequence_.Concatenate(_395_v54, Companion_Default___.Fm39(globalState)))
				var _396_v55 _dafny.Map
				_ = _396_v55
				_396_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				var _397_v56 _dafny.Set
				_ = _397_v56
				_397_v56 = _dafny.SetOf(p1, ((_396_v55).Merge(_396_v55)).Cardinality(), p1)
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index53
				((_this).F27()).ArraySet1(true, (_index53).Int())
				var _398_v57 _dafny.Map
				_ = _398_v57
				_398_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_397_v56).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_387_v46).Cardinality()))))
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index54
				var _rhs53 _dafny.Set = (func() _dafny.Set {
					if (_398_v57).Contains(p1) {
						return (_398_v57).Get(p1).(_dafny.Set)
					}
					return _397_v56
				})()
				_ = _rhs53
				var _rhs54 bool = (_this).F39()
				_ = _rhs54
				var _rhs55 bool = false
				_ = _rhs55
				var _lhs44 _dafny.Array = (_this).F27()
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _lhs45
				_397_v56 = _rhs53
				(_lhs44).ArraySet1(_rhs54, (_lhs45).Int())
				r1 = _rhs55
			} else {
				var _399_v58 _dafny.Sequence
				_ = _399_v58
				_399_v58 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				_399_v58 = _dafny.Companion_Sequence_.Concatenate(_399_v58, _399_v58)
				var _400_v59 D8
				_ = _400_v59
				_400_v59 = Companion_D8_.Create_DC18_(Companion_Default___.SafeModuloInt(p1, p1))
				_400_v59 = _400_v59
				(globalState).F17 = Companion_Default___.Fm40(globalState)
				(globalState).F5 = p1
				var _401_v60 _dafny.Map
				_ = _401_v60
				_401_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(337), Companion_Default___.SafeDivisionInt((_355_v19).Cardinality(), p1))
				var _402_v61 D11
				_ = _402_v61
				_402_v61 = Companion_D11_.Create_DC28_()
				var _403_v62 T1
				_ = _403_v62
				var _nw51 *C0 = New_C0_()
				_ = _nw51
				_nw51.Ctor__(_399_v58, p1, (_this).F27(), p2)
				_403_v62 = _nw51
				var _404_v63 _dafny.Map
				_ = _404_v63
				_404_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v62, _403_v62.F29())
				var _405_v64 _dafny.Map
				_ = _405_v64
				_405_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_404_v63, p2)
				var _406_v65 _dafny.CodePoint
				_ = _406_v65
				_406_v65 = _dafny.CodePoint('c')
				var _407_v66 D0
				_ = _407_v66
				_407_v66 = Companion_D0_.Create_DC0_(_405_v64, (_this).F39(), _dafny.IntOfInt64(-710), _406_v65, (_this).F28())
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index55
				((_this).F27()).ArraySet1((_407_v66).Dtor_cf1(), (_index55).Int())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index56
				var _rhs56 _dafny.Map = _401_v60
				_ = _rhs56
				var _rhs57 bool = (_this).F28()
				_ = _rhs57
				var _rhs58 _dafny.Int = p1
				_ = _rhs58
				var _rhs59 D11 = Companion_Default___.Fm41(_399_v58, globalState)
				_ = _rhs59
				var _rhs60 bool = (_this).F39()
				_ = _rhs60
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				var _lhs48 _dafny.Array = (_this).F27()
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _lhs49
				_401_v60 = _rhs56
				_lhs46.F17 = _rhs57
				_lhs47.F4 = _rhs58
				_402_v61 = _rhs59
				(_lhs48).ArraySet1(_rhs60, (_lhs49).Int())
			}
			_336_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p1)
		}
		var _408_v67 _dafny.Map
		_ = _408_v67
		_408_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm42(p1, globalState), p1)
		var _409_v68 _dafny.CodePoint
		_ = _409_v68
		_409_v68 = _dafny.CodePoint('r')
		var _410_v69 _dafny.Sequence
		_ = _410_v69
		_410_v69 = _dafny.UnicodeSeqOfUtf8Bytes("dpjwj")
		var _hi1 _dafny.Int = (func() _dafny.Int {
			if (_408_v67).Contains(_409_v68) {
				return (_408_v67).Get(_409_v68).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_410_v69).Cardinality())
		})()
		_ = _hi1
		for _411_i2 := p1; _411_i2.Cmp(_hi1) < 0; _411_i2 = _411_i2.Plus(_dafny.One) {
			var _412_v70 _dafny.MultiSet
			_ = _412_v70
			_412_v70 = _dafny.MultiSetOf(_411_i2)
			var _413_v71 _dafny.Sequence
			_ = _413_v71
			_413_v71 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("oom"))
			var _414_v72 _dafny.Map
			_ = _414_v72
			_414_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _409_v68)
			r0 = ((_412_v70).Difference(_dafny.MultiSetOf(_411_i2, p1, _411_i2, _dafny.IntOfInt64(-612), p1))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_413_v71).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F39(), (_dafny.Zero).Minus(p1))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("my")).Cardinality()), _411_i2, (_414_v72).Cardinality()))
			var _415_v73 _dafny.Array
			_ = _415_v73
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_5
			var _nw52 _dafny.Array
			_ = _nw52
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw52 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = func(_416_i3 _dafny.Int) bool {
					return true
				}
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw52 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw52).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw52).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_415_v73 = _nw52
			_415_v73 = (_this).F27()
			var _417_v74 _dafny.Sequence
			_ = _417_v74
			_417_v74 = _dafny.SeqOf(_411_i2, _dafny.IntOfInt64(403))
			r1 = ((func() _dafny.Int {
				if (_this).F39() {
					return (p0).Cardinality()
				}
				return _411_i2
			})()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_417_v74, _417_v74)).Cardinality())) < 0
			(globalState).F21 = _dafny.IntOfInt64(935)
		}
		var _418_v75 D1
		_ = _418_v75
		_418_v75 = Companion_D1_.Create_DC7_(_409_v68, _410_v69)
		var _419_v76 _dafny.Array
		_ = _419_v76
		var _nwElement0_7 _dafny.CodePoint = _409_v68
		_ = _nwElement0_7
		var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(18))
		_ = _nw53
		(_nw53).ArraySet1CodePoint(_nwElement0_7, 0)
		(_nw53).ArraySet1CodePoint((_418_v75).Dtor_cf13(), 1)
		(_nw53).ArraySet1CodePoint(_409_v68, 2)
		(_nw53).ArraySet1CodePoint(Companion_Default___.Fm42(p1, globalState), 3)
		(_nw53).ArraySet1CodePoint(_dafny.CodePoint('k'), 4)
		(_nw53).ArraySet1CodePoint(_409_v68, 5)
		(_nw53).ArraySet1CodePoint(_dafny.CodePoint('g'), 6)
		(_nw53).ArraySet1CodePoint(_dafny.CodePoint('y'), 7)
		(_nw53).ArraySet1CodePoint(_409_v68, 8)
		(_nw53).ArraySet1CodePoint(_409_v68, 9)
		(_nw53).ArraySet1CodePoint(_409_v68, 10)
		(_nw53).ArraySet1CodePoint(_409_v68, 11)
		(_nw53).ArraySet1CodePoint(_409_v68, 12)
		(_nw53).ArraySet1CodePoint(_409_v68, 13)
		(_nw53).ArraySet1CodePoint(_409_v68, 14)
		(_nw53).ArraySet1CodePoint(_409_v68, 15)
		(_nw53).ArraySet1CodePoint(_409_v68, 16)
		(_nw53).ArraySet1CodePoint(_409_v68, 17)
		_419_v76 = _nw53
		var _420_v77 _dafny.Map
		_ = _420_v77
		_420_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v76, (func() _dafny.Sequence {
			if (_this).F39() {
				return _410_v69
			}
			return _410_v69
		})())
		_420_v77 = _420_v77
		var _421_v78 _dafny.MultiSet
		_ = _421_v78
		_421_v78 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(487), p1))
		r0 = _421_v78
		r1 = p2
		return r0, r1
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
	_f27 _dafny.Array
	_f28 bool
	F38  bool
	_f37 _dafny.CodePoint
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
	_this.F38 = false
	_this._f37 = _dafny.CodePoint('D')
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

func (_this *C2) F27() _dafny.Array {
	return _this._f27
}
func (_this *C2) F28() bool {
	return _this._f28
}
func (_this *C2) Ctor__(f37 _dafny.CodePoint, f38 bool, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f37 = f37
		(_this).F38 = f38
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C2) Fm19(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_this.F38)).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf((_this).F28(), _this.F38)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, (func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter37 := _dafny.Iterate((_dafny.SetOf((_this).F37(), (_this).F37())).Elements()); ; {
				_compr_36, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _422_v0 _dafny.CodePoint
				_422_v0 = interface{}(_compr_36).(_dafny.CodePoint)
				if (_dafny.SetOf((_this).F37(), (_this).F37())).Contains(_422_v0) {
					_coll36.Add(_422_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F38, _this.F38)).Cardinality()))).Cardinality())
				}
			}
			return _coll36.ToMap()
		}()).Cardinality())))
	}
}
func (_this *C2) Fm20(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.MultiSetOf(!(false) || (_this.F38), (_this).F28())).Cardinality()
	}
}
func (_this *C2) M12(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		if _this.F38 {
			var _423_v0 _dafny.Map
			_ = _423_v0
			_423_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F38)
			_423_v0 = (_423_v0).Update((_dafny.Zero).Minus((p0).Minus(p0)), (_this).F28())
			var _424_v1 _dafny.Array
			_ = _424_v1
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw54
			_424_v1 = _nw54
			var _425_v2 _dafny.Sequence
			_ = _425_v2
			_425_v2 = _dafny.UnicodeSeqOfUtf8Bytes("tnt")
			var _426_v3 T1
			_ = _426_v3
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__(_425_v2, p3, (_this).F27(), !(_this.F38))
			_426_v3 = _nw55
			var _427_v4 _dafny.Sequence
			_ = _427_v4
			_427_v4 = _dafny.SeqOf(_426_v3)
			var _428_v5 _dafny.Map
			_ = _428_v5
			_428_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_424_v1, _dafny.Companion_Sequence_.Concatenate(_427_v4, _427_v4))
			_428_v5 = (_428_v5).Update(_424_v1, _427_v4)
			(globalState).F17 = (_426_v3).F28()
			(globalState).F5 = (_dafny.IntOfUint32((_425_v2).Cardinality())).Plus(_426_v3.F29())
			_425_v2 = _425_v2
		} else {
			var _429_v6 _dafny.Map
			_ = _429_v6
			_429_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())
			var _430_v7 _dafny.Sequence
			_ = _430_v7
			_430_v7 = _dafny.UnicodeSeqOfUtf8Bytes("gntq")
			var _431_v8 _dafny.Sequence
			_ = _431_v8
			_431_v8 = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), p0)
			var _432_v9 _dafny.Array
			_ = _432_v9
			var _nwElement0_8 _dafny.Int = ((_429_v6).Merge(_429_v6)).Cardinality()
			_ = _nwElement0_8
			var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(18))
			_ = _nw56
			(_nw56).ArraySet1(_nwElement0_8, 0)
			(_nw56).ArraySet1(_dafny.IntOfInt64(607), 1)
			(_nw56).ArraySet1(p3, 2)
			(_nw56).ArraySet1((_dafny.IntOfInt64(-478)).Times(p0), 3)
			(_nw56).ArraySet1((p3).Plus(p0), 4)
			(_nw56).ArraySet1(p0, 5)
			(_nw56).ArraySet1(_dafny.IntOfInt64(68), 6)
			(_nw56).ArraySet1(_dafny.IntOfInt64(101), 7)
			(_nw56).ArraySet1(_dafny.IntOfInt64(996), 8)
			(_nw56).ArraySet1(_dafny.IntOfUint32((_430_v7).Cardinality()), 9)
			(_nw56).ArraySet1(p3, 10)
			(_nw56).ArraySet1((func() _dafny.Int {
				if _this.F38 {
					return p0
				}
				return p0
			})(), 11)
			(_nw56).ArraySet1(_dafny.IntOfInt64(43), 12)
			(_nw56).ArraySet1(p3, 13)
			(_nw56).ArraySet1((_this).Fm20(globalState), 14)
			(_nw56).ArraySet1(p3, 15)
			(_nw56).ArraySet1(p3, 16)
			(_nw56).ArraySet1(_dafny.IntOfUint32((_431_v8).Cardinality()), 17)
			_432_v9 = _nw56
			_432_v9 = _432_v9
			(globalState).F21 = p0
			var _433_v10 D1
			_ = _433_v10
			_433_v10 = Companion_D1_.Create_DC7_(p1, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fkuwjpj"), _430_v7))
			var _source7 D1 = _433_v10
			_ = _source7
			if _source7.Is_DC6() {
				var _434___mcc_h0 _dafny.Int = _source7.Get_().(D1_DC6).Cf12
				_ = _434___mcc_h0
				var _435_cf12 _dafny.Int = _434___mcc_h0
				_ = _435_cf12
				var _436_v11 D0
				_ = _436_v11
				_436_v11 = Companion_D0_.Create_DC1_()
				var _437_v12 D0
				_ = _437_v12
				_437_v12 = Companion_D0_.Create_DC4_(_436_v11)
				_437_v12 = (func() D0 {
					if (_dafny.IntOfUint32((_430_v7).Cardinality())).Cmp(p3) >= 0 {
						return _437_v12
					}
					return _437_v12
				})()
				var _438_v13 *C0
				_ = _438_v13
				var _nw57 *C0 = New_C0_()
				_ = _nw57
				_nw57.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(192))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}(func(_439_i0 _dafny.Int) _dafny.CodePoint {
					return (_this).F37()
				})), _435_cf12, (_this).F27(), _this.F38)
				_438_v13 = _nw57
				var _440_v14 _dafny.Sequence
				_ = _440_v14
				_440_v14 = _dafny.SeqOf((_435_cf12).Cmp(p0) < 0, (_this).F28(), _dafny.Companion_Sequence_.IsProperPrefixOf(_431_v8, _431_v8), (_this).F28(), Companion_Default___.Fm22(true, _dafny.IntOfInt64(979), globalState))
				var _rhs61 bool = true
				_ = _rhs61
				var _rhs62 _dafny.Sequence = (Companion_Default___.Fm21((_this).F28(), p3, globalState)).Dtor_cf6()
				_ = _rhs62
				var _rhs63 bool = true
				_ = _rhs63
				var _rhs64 bool = (_dafny.IntOfUint32((_dafny.SeqOf((_this).F28())).Cardinality())).Cmp(p0) == 0
				_ = _rhs64
				var _rhs65 bool = (_440_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_440_v14).Cardinality()))).Uint32()).(bool)
				_ = _rhs65
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				var _lhs52 *GlobalState = globalState
				_ = _lhs52
				_lhs50.F15 = _rhs61
				_430_v7 = _rhs62
				r0 = _rhs63
				_lhs51.F6 = _rhs64
				_lhs52.F17 = _rhs65
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_432_v9), 0))
				_ = _index57
				(_432_v9).ArraySet1(((_dafny.Zero).Minus(_435_cf12)).Times(_435_cf12), (_index57).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_432_v9), 0))
				_ = _index58
				(_432_v9).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_435_cf12), (_this).Fm20(globalState)), (_index58).Int())
			} else if _source7.Is_DC7() {
				var _441___mcc_h1 _dafny.CodePoint = _source7.Get_().(D1_DC7).Cf13
				_ = _441___mcc_h1
				var _442___mcc_h2 _dafny.Sequence = _source7.Get_().(D1_DC7).Cf14
				_ = _442___mcc_h2
				var _443_cf14 _dafny.Sequence = _442___mcc_h2
				_ = _443_cf14
				var _444_cf13 _dafny.CodePoint = _441___mcc_h1
				_ = _444_cf13
				(globalState).F25 = p1
				(globalState).F21 = p3
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_432_v9), 0))
				_ = _index59
				(_432_v9).ArraySet1(p0, (_index59).Int())
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_432_v9), 0))
				_ = _index60
				(_432_v9).ArraySet1(Companion_Default___.SafeDivisionInt(p3, _dafny.IntOfInt64(887)), (_index60).Int())
				var _445_v15 _dafny.Map
				_ = _445_v15
				_445_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_Default___.Fm23(globalState))
				var _446_v16 _dafny.Set
				_ = _446_v16
				_446_v16 = _dafny.SetOf((_this).F28(), (_this).F28())
				var _447_v17 D1
				_ = _447_v17
				_447_v17 = Companion_D1_.Create_DC6_(p3)
				_445_v15 = (_445_v15).Update((_446_v16).IsProperSubsetOf(_dafny.SetOf(_this.F38, _this.F38)), _447_v17)
			} else {
				var _448___mcc_h3 _dafny.Sequence = _source7.Get_().(D1_DC5).Cf11
				_ = _448___mcc_h3
				var _449_cf11 _dafny.Sequence = _448___mcc_h3
				_ = _449_cf11
				(globalState).F13 = (_this).Fm20(globalState)
				var _450_v18 _dafny.Array
				_ = _450_v18
				var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
				_ = _nw58
				_450_v18 = _nw58
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_450_v18), 0))
				_ = _index61
				(_450_v18).ArraySet1((_this).F27(), (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_450_v18), 0))
				_ = _index62
				(_450_v18).ArraySet1((_this).F27(), (_index62).Int())
				var _451_v19 _dafny.MultiSet
				_ = _451_v19
				_451_v19 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(338), p3)
				(globalState).F5 = ((_dafny.Zero).Minus((p0).Minus(p0))).Plus((_451_v19).Cardinality())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index63
				((_this).F27()).ArraySet1(_this.F38, (_index63).Int())
				var _452_v20 _dafny.Map
				_ = _452_v20
				_452_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p3)
				var _453_v21 _dafny.Map
				_ = _453_v21
				_453_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v20).Cardinality(), p3)
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index64
				((_this).F27()).ArraySet1((func() bool {
					if (_this).F28() {
						return Companion_Default___.Fm22(true, _dafny.IntOfUint32((_dafny.SeqOf(_this.F38, _this.F38)).Cardinality()), globalState)
					}
					return ((_453_v21).Cardinality()).Cmp((_dafny.SetOf(true)).Cardinality()) > 0
				})(), (_index64).Int())
			}
			var _454_v22 _dafny.Sequence
			_ = _454_v22
			_454_v22 = _dafny.SeqOf(false, (_this).F28(), (_this).F28(), _this.F38)
			var _rhs66 bool = _this.F38
			_ = _rhs66
			var _rhs67 _dafny.Int = p0
			_ = _rhs67
			var _rhs68 bool = ((_this).Fm20(globalState)).Cmp(_dafny.IntOfUint32((_454_v22).Cardinality())) >= 0
			_ = _rhs68
			var _lhs53 *GlobalState = globalState
			_ = _lhs53
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			var _lhs55 *GlobalState = globalState
			_ = _lhs55
			_lhs53.F6 = _rhs66
			_lhs54.F21 = _rhs67
			_lhs55.F17 = _rhs68
			var _455_v23 _dafny.MultiSet
			_ = _455_v23
			_455_v23 = _dafny.MultiSetOf(_432_v9)
			_455_v23 = _455_v23
		}
		if (_this).F28() {
			var _456_v24 _dafny.Map
			_ = _456_v24
			_456_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if !(true) {
					return (_this).F28()
				}
				return _this.F38
			})(), p0)
			var _457_v25 _dafny.Set
			_ = _457_v25
			_457_v25 = _dafny.SetOf(p1)
			_456_v24 = (_456_v24).Update((_457_v25).IsProperSubsetOf(_457_v25), p3)
			var _458_v26 _dafny.Array
			_ = _458_v26
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_6
			var _nw59 _dafny.Array
			_ = _nw59
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw59 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = func(_459_i1 _dafny.Int) _dafny.Int {
					return (_459_i1).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chc")).Cardinality()))
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw59 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw59).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw59).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_458_v26 = _nw59
			_458_v26 = _458_v26
			var _460_v27 D0
			_ = _460_v27
			_460_v27 = Companion_D0_.Create_DC1_()
			var _source8 D0 = _460_v27
			_ = _source8
			if _source8.Is_DC0() {
				var _461___mcc_h4 _dafny.Map = _source8.Get_().(D0_DC0).Cf0
				_ = _461___mcc_h4
				var _462___mcc_h5 bool = _source8.Get_().(D0_DC0).Cf1
				_ = _462___mcc_h5
				var _463___mcc_h6 _dafny.Int = _source8.Get_().(D0_DC0).Cf2
				_ = _463___mcc_h6
				var _464___mcc_h7 _dafny.CodePoint = _source8.Get_().(D0_DC0).Cf3
				_ = _464___mcc_h7
				var _465___mcc_h8 bool = _source8.Get_().(D0_DC0).Cf4
				_ = _465___mcc_h8
				var _466_cf4 bool = _465___mcc_h8
				_ = _466_cf4
				var _467_cf3 _dafny.CodePoint = _464___mcc_h7
				_ = _467_cf3
				var _468_cf2 _dafny.Int = _463___mcc_h6
				_ = _468_cf2
				var _469_cf1 bool = _462___mcc_h5
				_ = _469_cf1
				var _470_cf0 _dafny.Map = _461___mcc_h4
				_ = _470_cf0
				r0 = _469_cf1
				var _471_v28 _dafny.Set
				_ = _471_v28
				_471_v28 = _dafny.SetOf((_dafny.IntOfInt64(-739)).Minus(p3))
				var _472_v29 _dafny.Map
				_ = _472_v29
				_472_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_cf2, _dafny.IntOfInt64(24))
				var _473_v30 _dafny.Map
				_ = _473_v30
				_473_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_472_v29, _471_v28)
				var _474_v31 _dafny.MultiSet
				_ = _474_v31
				_474_v31 = _dafny.MultiSetOf(_466_cf4)
				var _475_v32 _dafny.Sequence
				_ = _475_v32
				_475_v32 = _dafny.UnicodeSeqOfUtf8Bytes("qqvo")
				var _476_v33 _dafny.Map
				_ = _476_v33
				_476_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm22(!(_466_cf4), p0, globalState), _475_v32)
				var _477_v34 D0
				_ = _477_v34
				_477_v34 = Companion_D0_.Create_DC2_(_471_v28, _475_v32)
				var _rhs69 bool = (_this).F28()
				_ = _rhs69
				var _rhs70 _dafny.Set = _dafny.SetOf((p3).Minus(p0), p3, (_dafny.Zero).Minus(((_474_v31).Update(_466_cf4, Companion_Default___.Abs(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_476_v33).Contains(_466_cf4) {
						return (_476_v33).Get(_466_cf4).(_dafny.Sequence)
					}
					return (_477_v34).Dtor_cf6()
				})()).Cardinality())))).Cardinality()))
				_ = _rhs70
				var _rhs71 _dafny.Map = ((_473_v30).Update(_472_v29, _471_v28)).Update(_472_v29, func() _dafny.Set {
					var _coll37 = _dafny.NewBuilder()
					_ = _coll37
					for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(904), _dafny.IntOfInt64(560))); ; {
						_compr_37, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _478_v35 _dafny.Int
						_478_v35 = interface{}(_compr_37).(_dafny.Int)
						if ((_dafny.IntOfInt64(904)).Cmp(_478_v35) <= 0) && ((_478_v35).Cmp(_dafny.IntOfInt64(560)) < 0) {
							_coll37.Add(Companion_Default___.SafeModuloInt(_478_v35, _468_cf2))
						}
					}
					return _coll37.ToSet()
				}())
				_ = _rhs71
				var _lhs56 *GlobalState = globalState
				_ = _lhs56
				_lhs56.F15 = _rhs69
				_471_v28 = _rhs70
				_473_v30 = _rhs71
				(globalState).F21 = _468_cf2
				var _479_v36 D1
				_ = _479_v36
				_479_v36 = Companion_D1_.Create_DC6_(p0)
				(globalState).F6 = ((Companion_Default___.Fm24(globalState)).Difference(_dafny.SetOf(_479_v36, _479_v36, _479_v36, _479_v36))).Equals(_dafny.SetOf(_479_v36))
			} else if _source8.Is_DC1() {
				var _480_v37 _dafny.Sequence
				_ = _480_v37
				_480_v37 = _dafny.SeqOf(p0)
				var _rhs72 bool = _this.F38
				_ = _rhs72
				var _rhs73 bool = Companion_Default___.Fm22((_this).F28(), (_dafny.Zero).Minus((_480_v37).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_480_v37).Cardinality()))).Uint32()).(_dafny.Int)), globalState)
				_ = _rhs73
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				_lhs57.F15 = _rhs72
				_lhs58.F6 = _rhs73
				var _481_v38 _dafny.Map
				_ = _481_v38
				_481_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
				var _482_v39 D3
				_ = _482_v39
				_482_v39 = Companion_D3_.Create_DC10_(_this.F38, _this.F38)
				var _483_v40 _dafny.Set
				_ = _483_v40
				_483_v40 = _dafny.SetOf(_482_v39)
				var _484_v41 _dafny.Set
				_ = _484_v41
				_484_v41 = _dafny.SetOf(_483_v40, _483_v40)
				var _485_v42 _dafny.Map
				_ = _485_v42
				_485_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_481_v38).Contains(_this.F38) {
						return (_481_v38).Get(_this.F38).(bool)
					}
					return _this.F38
				})(), (func() _dafny.Set {
					if _this.F38 {
						return _484_v41
					}
					return _dafny.SetOf(_483_v40, _dafny.SetOf(Companion_D3_.Create_DC10_((_this).F28(), _this.F38)), _483_v40, _dafny.SetOf(_482_v39, _482_v39, _482_v39, func(_pat_let11_0 D3) D3 {
						return func(_486_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let12_0 bool) D3 {
								return func(_487_dt__update_hcf18_h0 bool) D3 {
									return Companion_D3_.Create_DC10_((_486_dt__update__tmp_h0).Dtor_cf17(), _487_dt__update_hcf18_h0)
								}(_pat_let12_0)
							}((_this).F28())
						}(_pat_let11_0)
					}(Companion_D3_.Create_DC10_(!(_this.F38), (_this).F28())), _482_v39), _dafny.SetOf(_482_v39))
				})())
				_485_v42 = (_485_v42).Update((_this).F28(), (_484_v41).Difference(_484_v41))
				(globalState).F5 = Companion_Default___.SafeModuloInt(p3, p3)
				var _488_v43 _dafny.Sequence
				_ = _488_v43
				_488_v43 = _dafny.SeqOf(false, _this.F38)
				_488_v43 = Companion_Default___.Fm25(p0, (Companion_Default___.Fm26((_this).F28(), p0, _dafny.CodePoint('n'), globalState)).Cardinality(), Companion_Default___.SafeDivisionInt(p0, p3), globalState)
			} else if _source8.Is_DC2() {
				var _489___mcc_h9 _dafny.Set = _source8.Get_().(D0_DC2).Cf5
				_ = _489___mcc_h9
				var _490___mcc_h10 _dafny.Sequence = _source8.Get_().(D0_DC2).Cf6
				_ = _490___mcc_h10
				var _491_cf6 _dafny.Sequence = _490___mcc_h10
				_ = _491_cf6
				var _492_cf5 _dafny.Set = _489___mcc_h9
				_ = _492_cf5
				var _493_v44 _dafny.Sequence
				_ = _493_v44
				_493_v44 = _dafny.SeqOf((_this).F28(), (_this).F28())
				_493_v44 = (func() _dafny.Sequence {
					if true {
						return _493_v44
					}
					return Companion_Default___.Fm25(_dafny.IntOfUint32((_491_cf6).Cardinality()), _dafny.IntOfInt64(689), p0, globalState)
				})()
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index65
				((_this).F27()).ArraySet1((_this).F28(), (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index66
				((_this).F27()).ArraySet1(((_this).F28()) && (!(_dafny.SetOf(_dafny.IntOfInt64(557))).Equals(_492_cf5)), (_index66).Int())
				var _494_v45 _dafny.Set
				_ = _494_v45
				_494_v45 = _dafny.SetOf(_this.F38)
				var _495_v46 _dafny.Sequence
				_ = _495_v46
				_495_v46 = _dafny.SeqOf(_494_v45, _494_v45, (_494_v45).Union(_dafny.SetOf(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), !(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool)), _this.F38)))
				var _496_v47 D5
				_ = _496_v47
				_496_v47 = Companion_D5_.Create_DC12_(_495_v46)
				_495_v46 = (_496_v47).Dtor_cf20()
				(globalState).F5 = (_dafny.Zero).Minus(p3)
			} else if _source8.Is_DC3() {
				var _497___mcc_h11 _dafny.Int = _source8.Get_().(D0_DC3).Cf7
				_ = _497___mcc_h11
				var _498___mcc_h12 bool = _source8.Get_().(D0_DC3).Cf8
				_ = _498___mcc_h12
				var _499___mcc_h13 bool = _source8.Get_().(D0_DC3).Cf9
				_ = _499___mcc_h13
				var _500_cf9 bool = _499___mcc_h13
				_ = _500_cf9
				var _501_cf8 bool = _498___mcc_h12
				_ = _501_cf8
				var _502_cf7 _dafny.Int = _497___mcc_h11
				_ = _502_cf7
				var _503_v48 _dafny.Set
				_ = _503_v48
				_503_v48 = _dafny.SetOf(_502_cf7, p3, p3)
				var _504_v49 _dafny.Map
				_ = _504_v49
				_504_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_502_cf7, _503_v48)
				var _rhs74 bool = Companion_Default___.Fm22(_this.F38, (_dafny.Zero).Minus((p0).Minus(p3)), globalState)
				_ = _rhs74
				var _rhs75 bool = Companion_Default___.Fm22((Companion_Default___.Fm0(_504_v49, false, false, globalState)).IsDisjointFrom(_503_v48), Companion_Default___.SafeDivisionInt(p0, _502_cf7), globalState)
				_ = _rhs75
				var _rhs76 bool = _this.F38
				_ = _rhs76
				var _lhs59 *GlobalState = globalState
				_ = _lhs59
				var _lhs60 *GlobalState = globalState
				_ = _lhs60
				var _lhs61 *C2 = _this
				_ = _lhs61
				_lhs59.F17 = _rhs74
				_lhs60.F17 = _rhs75
				_lhs61.F38 = _rhs76
				var _505_v50 _dafny.Set
				_ = _505_v50
				_505_v50 = _dafny.SetOf(_501_cf8, _500_cf9)
				var _506_v51 D0
				_ = _506_v51
				_506_v51 = Companion_D0_.Create_DC1_()
				var _507_v52 D0
				_ = _507_v52
				_507_v52 = Companion_D0_.Create_DC4_(_506_v51)
				var _508_v53 _dafny.Map
				_ = _508_v53
				_508_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_505_v50, _507_v52)
				_508_v53 = (_508_v53).Update(_505_v50, _507_v52)
				(globalState).F13 = (func() _dafny.Int {
					if _501_cf8 {
						return _dafny.IntOfInt64(175)
					}
					return p0
				})()
				var _509_v54 _dafny.Sequence
				_ = _509_v54
				_509_v54 = _dafny.UnicodeSeqOfUtf8Bytes("v")
				var _510_v55 T1
				_ = _510_v55
				var _nw60 *C0 = New_C0_()
				_ = _nw60
				_nw60.Ctor__(_509_v54, _dafny.IntOfInt64(735), (_this).F27(), (_this).F28())
				_510_v55 = _nw60
				var _511_v56 _dafny.Map
				_ = _511_v56
				_511_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_v55, _510_v55.F29())
				var _512_v57 _dafny.Map
				_ = _512_v57
				_512_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_511_v56, !(_this.F38))
				var _513_v58 D0
				_ = _513_v58
				_513_v58 = Companion_D0_.Create_DC0_(_512_v57, _500_cf9, _510_v55.F29(), (_this).F37(), (_this).F28())
				(globalState).F17 = (_513_v58).Dtor_cf1()
			} else {
				var _514___mcc_h14 D0 = _source8.Get_().(D0_DC4).Cf10
				_ = _514___mcc_h14
				var _515_cf10 D0 = _514___mcc_h14
				_ = _515_cf10
				var _516_v59 D5
				_ = _516_v59
				_516_v59 = Companion_D5_.Create_DC13_((_this).F28(), _458_v26, _this.F38)
				var _517_v60 _dafny.Set
				_ = _517_v60
				_517_v60 = _dafny.SetOf(_516_v59)
				var _518_v61 _dafny.Array
				_ = _518_v61
				var _nwElement0_9 _dafny.Set = _517_v60
				_ = _nwElement0_9
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(10))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_9, 0)
				(_nw61).ArraySet1(_517_v60, 1)
				(_nw61).ArraySet1(_517_v60, 2)
				(_nw61).ArraySet1((_517_v60).Union(_517_v60), 3)
				(_nw61).ArraySet1(_dafny.SetOf(Companion_D5_.Create_DC13_(_this.F38, _458_v26, (_this).F28())), 4)
				(_nw61).ArraySet1(_517_v60, 5)
				(_nw61).ArraySet1(_517_v60, 6)
				(_nw61).ArraySet1(_517_v60, 7)
				(_nw61).ArraySet1(_dafny.SetOf(_516_v59, _516_v59), 8)
				(_nw61).ArraySet1((_dafny.SetOf(_516_v59, _516_v59, _516_v59)).Union(_517_v60), 9)
				_518_v61 = _nw61
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_518_v61), 0))
				_ = _index67
				(_518_v61).ArraySet1(_517_v60, (_index67).Int())
				var _519_v62 _dafny.Map
				_ = _519_v62
				_519_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.MultiSetOf((_this).F27(), (_this).F27()))
				var _520_v63 _dafny.MultiSet
				_ = _520_v63
				_520_v63 = _dafny.MultiSetOf((_this).F27())
				var _521_v64 _dafny.Set
				_ = _521_v64
				_521_v64 = _dafny.SetOf(p0, p3)
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_518_v61), 0))
				_ = _index68
				var _rhs77 bool = ((func() _dafny.MultiSet {
					if (_519_v62).Contains(p0) {
						return (_519_v62).Get(p0).(_dafny.MultiSet)
					}
					return (_520_v63).Update((_this).F27(), Companion_Default___.Abs((_521_v64).Cardinality()))
				})()).IsSubsetOf(_dafny.MultiSetOf((_this).F27()))
				_ = _rhs77
				var _rhs78 _dafny.Set = _dafny.SetOf(_516_v59)
				_ = _rhs78
				var _rhs79 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _rhs79
				var _rhs80 _dafny.CodePoint = (_this).F37()
				_ = _rhs80
				var _rhs81 bool = (_this).F28()
				_ = _rhs81
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 _dafny.Array = _518_v61
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_518_v61), 0))
				_ = _lhs64
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				var _lhs66 *GlobalState = globalState
				_ = _lhs66
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				_lhs62.F17 = _rhs77
				(_lhs63).ArraySet1(_rhs78, (_lhs64).Int())
				_lhs65.F5 = _rhs79
				_lhs66.F25 = _rhs80
				_lhs67.F17 = _rhs81
				var _522_v65 _dafny.MultiSet
				_ = _522_v65
				_522_v65 = _dafny.MultiSetOf(p1)
				var _523_v66 _dafny.Map
				_ = _523_v66
				_523_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_522_v65).Cardinality(), _this.F38)
				var _524_v68 _dafny.Array
				_ = _524_v68
				var _nwElement0_10 _dafny.Map = _523_v66
				_ = _nwElement0_10
				var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(5))
				_ = _nw62
				(_nw62).ArraySet1(_nwElement0_10, 0)
				(_nw62).ArraySet1((func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(236), _dafny.IntOfInt64(663))); ; {
						_compr_38, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _525_v67 _dafny.Int
						_525_v67 = interface{}(_compr_38).(_dafny.Int)
						if ((_dafny.IntOfInt64(236)).Cmp(_525_v67) <= 0) && ((_525_v67).Cmp(_dafny.IntOfInt64(663)) < 0) {
							_coll38.Add(Companion_Default___.SafeModuloInt(_525_v67, p3), _this.F38)
						}
					}
					return _coll38.ToMap()
				}()).Merge(_523_v66), 1)
				(_nw62).ArraySet1(_523_v66, 2)
				(_nw62).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F28()), 3)
				(_nw62).ArraySet1(_523_v66, 4)
				_524_v68 = _nw62
				var _526_v69 _dafny.Sequence
				_ = _526_v69
				_526_v69 = _dafny.UnicodeSeqOfUtf8Bytes("osp")
				var _527_v70 _dafny.Map
				_ = _527_v70
				_527_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_D1_.Create_DC7_((_this).F37(), _526_v69))
				var _528_v71 _dafny.Sequence
				_ = _528_v71
				_528_v71 = _dafny.SeqOf(_527_v70)
				var _rhs82 _dafny.Set = func() _dafny.Set {
					var _coll39 = _dafny.NewBuilder()
					_ = _coll39
					for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(637), _dafny.IntOfInt64(773))); ; {
						_compr_39, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _529_v72 _dafny.Int
						_529_v72 = interface{}(_compr_39).(_dafny.Int)
						if ((_dafny.IntOfInt64(637)).Cmp(_529_v72) <= 0) && ((_529_v72).Cmp(_dafny.IntOfInt64(773)) < 0) {
							_coll39.Add((_529_v72).Plus(p3))
						}
					}
					return _coll39.ToSet()
				}()
				_ = _rhs82
				var _rhs83 _dafny.Int = p0
				_ = _rhs83
				var _rhs84 bool = (_this).F28()
				_ = _rhs84
				var _rhs85 _dafny.Array = _524_v68
				_ = _rhs85
				var _rhs86 _dafny.Sequence = _dafny.SeqOf(Companion_Default___.Fm27(_this.F38, p3, _this.F38, globalState))
				_ = _rhs86
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				_521_v64 = _rhs82
				_lhs68.F21 = _rhs83
				_lhs69.F15 = _rhs84
				_524_v68 = _rhs85
				_528_v71 = _rhs86
				_526_v69 = _526_v69
				var _530_v73 D0
				_ = _530_v73
				_530_v73 = Companion_D0_.Create_DC2_(_521_v64, _526_v69)
				var _531_v74 _dafny.Set
				_ = _531_v74
				_531_v74 = _dafny.SetOf(_this.F38, ((_this).F28()) && ((_this).F28()), (_dafny.IntOfUint32(((_530_v73).Dtor_cf6()).Cardinality())).Cmp(p0) == 0, (_this).F28(), !(false))
				_531_v74 = _531_v74
			}
			(globalState).F21 = (_dafny.Zero).Minus((_dafny.Zero).Minus(p3))
			var _532_v75 _dafny.Map
			_ = _532_v75
			var _533_v76 _dafny.Int
			_ = _533_v76
			var _out6 _dafny.Map
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out6, _out7 = Companion_Default___.M0(globalState)
			_532_v75 = _out6
			_533_v76 = _out7
		} else {
			if (_this).F28() {
				var _534_v77 _dafny.Sequence
				_ = _534_v77
				_534_v77 = _dafny.SeqOf(_this.F38)
				var _535_v78 _dafny.MultiSet
				_ = _535_v78
				_535_v78 = _dafny.MultiSetOf(((_dafny.Zero).Minus(p3)).Cmp(p3) == 0, (_534_v77).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_534_v77).Cardinality()))).Uint32()).(bool), _this.F38)
				var _536_v79 _dafny.Sequence
				_ = _536_v79
				_536_v79 = _dafny.UnicodeSeqOfUtf8Bytes("gio")
				var _537_v80 _dafny.Array
				_ = _537_v80
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw63
				_537_v80 = _nw63
				var _538_v81 _dafny.Map
				_ = _538_v81
				_538_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F38) || ((_534_v77).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_534_v77).Cardinality()))).Uint32()).(bool)), _dafny.Companion_Sequence_.Update(_536_v79, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.IntOfUint32((_536_v79).Cardinality()))).Uint32(), (_this).F37()))
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))
				_ = _index69
				(_537_v80).ArraySet1(p0, (_index69).Int())
				var _539_v82 _dafny.Set
				_ = _539_v82
				_539_v82 = _dafny.SetOf((_this).Fm20(globalState))
				var _540_v83 _dafny.Sequence
				_ = _540_v83
				_540_v83 = _dafny.SeqOf(_539_v82, _539_v82, _dafny.SetOf(p0), _539_v82, _539_v82)
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))
				_ = _index70
				var _rhs87 _dafny.MultiSet = (_535_v78).Difference((_535_v78).Intersection(Companion_Default___.Fm26((_534_v77).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_534_v77).Cardinality()))).Uint32()).(bool), p0, p1, globalState)))
				_ = _rhs87
				var _rhs88 _dafny.Sequence = _536_v79
				_ = _rhs88
				var _rhs89 _dafny.Array = _537_v80
				_ = _rhs89
				var _rhs90 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _536_v79)
				_ = _rhs90
				var _rhs91 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_540_v83, _dafny.SeqOf(_dafny.SetOf(p0)))).Cardinality())).Plus(p0)
				_ = _rhs91
				var _lhs70 _dafny.Array = _537_v80
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))
				_ = _lhs71
				_535_v78 = _rhs87
				_536_v79 = _rhs88
				_537_v80 = _rhs89
				_538_v81 = _rhs90
				(_lhs70).ArraySet1(_rhs91, (_lhs71).Int())
				var _541_v84 *C0
				_ = _541_v84
				var _nw64 *C0 = New_C0_()
				_ = _nw64
				_nw64.Ctor__(_536_v79, (_dafny.Zero).Minus(_dafny.IntOfInt64(-36)), (_this).F27(), (_this).F28())
				_541_v84 = _nw64
				(globalState).F21 = (_537_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))).Int()).(_dafny.Int)
				var _542_v85 _dafny.Array
				_ = _542_v85
				_542_v85 = (_this).F27()
				var _543_v86 *C0
				_ = _543_v86
				var _nw65 *C0 = New_C0_()
				_ = _nw65
				_nw65.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fervo"), (_541_v84).F36()), (_this).Fm20(globalState), (_542_v85), false)
				_543_v86 = _nw65
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index71
				((_this).F27()).ArraySet1((_this).F28(), (_index71).Int())
				var _544_v87 _dafny.Map
				_ = _544_v87
				_544_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.UnicodeSeqOfUtf8Bytes("vo"))
				var _545_v89 _dafny.Map
				_ = _545_v89
				_545_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _544_v87)
				var _546_v90 _dafny.Sequence
				_ = _546_v90
				_546_v90 = _dafny.SeqOf(Companion_Default___.Fm28(p0, globalState))
				var _547_v92 _dafny.Array
				_ = _547_v92
				var _nwElement0_11 _dafny.Map = _544_v87
				_ = _nwElement0_11
				var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(15))
				_ = _nw66
				(_nw66).ArraySet1(_nwElement0_11, 0)
				(_nw66).ArraySet1(_544_v87, 1)
				(_nw66).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_537_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("je")), 2)
				(_nw66).ArraySet1(_544_v87, 3)
				(_nw66).ArraySet1(func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(547), _dafny.IntOfInt64(-19))); ; {
						_compr_40, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _548_v88 _dafny.Int
						_548_v88 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(547)).Cmp(_548_v88) <= 0) && ((_548_v88).Cmp(_dafny.IntOfInt64(-19)) < 0) {
							_coll40.Add(Companion_Default___.SafeModuloInt(_548_v88, _dafny.IntOfInt64(510)), _536_v79)
						}
					}
					return _coll40.ToMap()
				}(), 4)
				(_nw66).ArraySet1((func() _dafny.Map {
					if (_545_v89).Contains(p3) {
						return (_545_v89).Get(p3).(_dafny.Map)
					}
					return (_546_v90).Select((Companion_Default___.SafeIndex((_537_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_546_v90).Cardinality()))).Uint32()).(_dafny.Map)
				})(), 5)
				(_nw66).ArraySet1(_544_v87, 6)
				(_nw66).ArraySet1((_544_v87).Merge((Companion_D7_.Create_DC15_(_544_v87)).Dtor_cf25()), 7)
				(_nw66).ArraySet1(Companion_Default___.Fm28(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, _this.F38)).Cardinality(), p3, p3)).Cardinality()), globalState), 8)
				(_nw66).ArraySet1((_544_v87).Merge(_544_v87), 9)
				(_nw66).ArraySet1(func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(568), _dafny.IntOfInt64(385))); ; {
						_compr_41, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _549_v91 _dafny.Int
						_549_v91 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(568)).Cmp(_549_v91) <= 0) && ((_549_v91).Cmp(_dafny.IntOfInt64(385)) < 0) {
							_coll41.Add((_549_v91).Plus(p0), _dafny.UnicodeSeqOfUtf8Bytes("fnogundpk"))
						}
					}
					return _coll41.ToMap()
				}(), 10)
				(_nw66).ArraySet1(_544_v87, 11)
				(_nw66).ArraySet1((_544_v87).Update(p0, _536_v79), 12)
				(_nw66).ArraySet1((func() _dafny.Map {
					if (_this).F28() {
						return _544_v87
					}
					return _544_v87
				})(), 13)
				(_nw66).ArraySet1(_544_v87, 14)
				_547_v92 = _nw66
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_547_v92), 0))
				_ = _index72
				(_547_v92).ArraySet1(_544_v87, (_index72).Int())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index73
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_547_v92), 0))
				_ = _index74
				var _rhs92 bool = _this.F38
				_ = _rhs92
				var _rhs93 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), (_537_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_537_v80), 0))).Int()).(_dafny.Int))
				_ = _rhs93
				var _rhs94 _dafny.Map = _544_v87
				_ = _rhs94
				var _rhs95 bool = !((_this).F28()) || (_this.F38)
				_ = _rhs95
				var _lhs72 _dafny.Array = (_this).F27()
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _lhs73
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 _dafny.Array = _547_v92
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_547_v92), 0))
				_ = _lhs76
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				(_lhs72).ArraySet1(_rhs92, (_lhs73).Int())
				_lhs74.F13 = _rhs93
				(_lhs75).ArraySet1(_rhs94, (_lhs76).Int())
				_lhs77.F15 = _rhs95
			} else {
				var _550_v93 D0
				_ = _550_v93
				_550_v93 = Companion_D0_.Create_DC3_(p3, (_this).F28(), (_this).F28())
				var _551_v94 D0
				_ = _551_v94
				_551_v94 = Companion_D0_.Create_DC4_(_550_v93)
				var _552_v95 _dafny.MultiSet
				_ = _552_v95
				_552_v95 = _dafny.MultiSetOf(_551_v94)
				var _553_v96 _dafny.Map
				_ = _553_v96
				_553_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p0)
				var _rhs96 _dafny.Int = (func() _dafny.Int {
					if (_this).F28() {
						return (_553_v96).Cardinality()
					}
					return p3
				})()
				_ = _rhs96
				var _rhs97 bool = _this.F38
				_ = _rhs97
				var _rhs98 _dafny.MultiSet = _dafny.MultiSetOf(_551_v94, _551_v94)
				_ = _rhs98
				var _rhs99 bool = _this.F38
				_ = _rhs99
				var _rhs100 _dafny.Int = (_dafny.Zero).Minus(p3)
				_ = _rhs100
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				_lhs78.F4 = _rhs96
				_lhs79.F17 = _rhs97
				_552_v95 = _rhs98
				_lhs80.F15 = _rhs99
				_lhs81.F13 = _rhs100
				var _554_v97 _dafny.Array
				_ = _554_v97
				var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
				_ = _nw67
				_554_v97 = _nw67
				var _555_v98 _dafny.Map
				_ = _555_v98
				_555_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _556_v99 _dafny.Sequence
				_ = _556_v99
				_556_v99 = _dafny.SeqOf(_555_v98)
				var _557_v100 _dafny.Map
				_ = _557_v100
				_557_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _556_v99)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_554_v97), 0))
				_ = _index75
				(_554_v97).ArraySet1((func() _dafny.Sequence {
					if (_557_v100).Contains(p0) {
						return (_557_v100).Get(p0).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}((func(_558_v98 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_559_i2 _dafny.Int) _dafny.Map {
							return _558_v98
						}
					})(_555_v98)))
				})(), (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_554_v97), 0))
				_ = _index76
				(_554_v97).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_556_v99, _556_v99), (_index76).Int())
				var _560_v101 _dafny.Sequence
				_ = _560_v101
				_560_v101 = _dafny.UnicodeSeqOfUtf8Bytes("jjcirftt")
				var _561_v102 D0
				_ = _561_v102
				var _out8 D0
				_ = _out8
				_out8 = (_this).M13(_this.F38, (_this).F28(), (Companion_Default___.Fm29(_this.F38, true, globalState)).Cardinality(), _560_v101, globalState)
				_561_v102 = _out8
				var _562_v103 _dafny.Map
				_ = _562_v103
				var _563_v104 _dafny.Int
				_ = _563_v104
				var _out9 _dafny.Map
				_ = _out9
				var _out10 _dafny.Int
				_ = _out10
				_out9, _out10 = Companion_Default___.M0(globalState)
				_562_v103 = _out9
				_563_v104 = _out10
				var _564_v105 *C0
				_ = _564_v105
				var _nw68 *C0 = New_C0_()
				_ = _nw68
				_nw68.Ctor__(_dafny.Companion_Sequence_.Concatenate(_560_v101, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(628))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_565_i3 _dafny.Int) _dafny.CodePoint {
					return (_this).F37()
				}))), p3, (func() _dafny.Array {
					if !((_this).F28()) {
						return (_this).F27()
					}
					return (_this).F27()
				})(), Companion_Default___.Fm22(false, p0, globalState))
				_564_v105 = _nw68
			}
			var _566_v106 _dafny.Sequence
			_ = _566_v106
			_566_v106 = _dafny.UnicodeSeqOfUtf8Bytes("nke")
			var _567_v107 *C0
			_ = _567_v107
			var _nw69 *C0 = New_C0_()
			_ = _nw69
			_nw69.Ctor__(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("unybyg"), _566_v106), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("unybyg"), _566_v106)).Cardinality()))).Uint32(), (_this).F37()), _dafny.IntOfUint32((_dafny.SeqOf((_this).Fm20(globalState))).Cardinality()), (_this).F27(), (_this).F28())
			_567_v107 = _nw69
			r0 = Companion_Default___.Fm22(_this.F38, p3, globalState)
			var _568_v108 _dafny.Map
			_ = _568_v108
			_568_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.SetOf(!((_this).F28()))).Cardinality())
			(globalState).F4 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_568_v108).Contains(p3) {
					return (_568_v108).Get(p3).(_dafny.Int)
				}
				return p0
			})())
			if !(Companion_Default___.Fm22(_this.F38, p3, globalState)) {
				(globalState).F4 = ((_this).Fm20(globalState)).Times(p3)
				(globalState).F17 = (_this).F28()
				var _569_v109 D0
				_ = _569_v109
				_569_v109 = Companion_D0_.Create_DC3_((_this).Fm20(globalState), _this.F38, (_this).F28())
				var _570_v110 _dafny.Map
				_ = _570_v110
				_570_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_569_v109).Dtor_cf7(), (_this).F28())
				_570_v110 = (_570_v110).Merge(Companion_Default___.Fm30(_this.F38, _dafny.UnicodeSeqOfUtf8Bytes("a"), _this.F38, globalState))
				var _571_v111 _dafny.Map
				_ = _571_v111
				_571_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, !(_this.F38))
				r0 = (func() bool {
					if (_571_v111).Contains(_this.F38) {
						return (_571_v111).Get(_this.F38).(bool)
					}
					return (p3).Cmp(p3) >= 0
				})()
				var _572_v112 _dafny.Sequence
				_ = _572_v112
				_572_v112 = _dafny.SeqOf((!((_this).F28())) || (_this.F38), _this.F38, (!(true)) || (!(!((_this).F28()))))
				var _573_v113 _dafny.Sequence
				_ = _573_v113
				_573_v113 = _dafny.SeqOf(_572_v112, Companion_Default___.Fm25(p0, (_this).Fm20(globalState), p3, globalState))
				_572_v112 = _dafny.Companion_Sequence_.Concatenate((_573_v113).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_573_v113).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(_this.F38, (_572_v112).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_572_v112).Cardinality()))).Uint32()).(bool), !(_this.F38)))
			} else {
				var _574_v114 _dafny.Map
				_ = _574_v114
				_574_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, !((_this).F28()))
				var _575_v115 _dafny.MultiSet
				_ = _575_v115
				_575_v115 = _dafny.MultiSetOf(p0, p3)
				_574_v114 = (_574_v114).Update(p0, ((_575_v115).Update(p3, Companion_Default___.Abs(p3))).IsDisjointFrom(_575_v115))
				var _576_v116 _dafny.Array
				_ = _576_v116
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw70
				_576_v116 = _nw70
				var _577_v117 _dafny.MultiSet
				_ = _577_v117
				_577_v117 = _dafny.MultiSetOf(_576_v116, _576_v116, _576_v116)
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index77
				((_this).F27()).ArraySet1(!(_577_v117).Contains(_576_v116), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index78
				((_this).F27()).ArraySet1(_this.F38, (_index78).Int())
				(globalState).F4 = (Companion_Default___.SafeModuloInt(p0, p0)).Minus(p3)
				var _578_v118 _dafny.Sequence
				_ = _578_v118
				_578_v118 = _dafny.SeqOf(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), (_this).F28())
				var _579_v119 _dafny.Array
				_ = _579_v119
				var _nwElement0_12 _dafny.CodePoint = p1
				_ = _nwElement0_12
				var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(28))
				_ = _nw71
				(_nw71).ArraySet1CodePoint(_nwElement0_12, 0)
				(_nw71).ArraySet1CodePoint(p1, 1)
				(_nw71).ArraySet1CodePoint(p1, 2)
				(_nw71).ArraySet1CodePoint((_this).F37(), 3)
				(_nw71).ArraySet1CodePoint((_this).F37(), 4)
				(_nw71).ArraySet1CodePoint((_this).F37(), 5)
				(_nw71).ArraySet1CodePoint((_this).F37(), 6)
				(_nw71).ArraySet1CodePoint(_dafny.CodePoint('h'), 7)
				(_nw71).ArraySet1CodePoint(p1, 8)
				(_nw71).ArraySet1CodePoint((_this).F37(), 9)
				(_nw71).ArraySet1CodePoint(p1, 10)
				(_nw71).ArraySet1CodePoint((_this).F37(), 11)
				(_nw71).ArraySet1CodePoint((_this).F37(), 12)
				(_nw71).ArraySet1CodePoint((_this).F37(), 13)
				(_nw71).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _this.F38 {
						return (_this).F37()
					}
					return (_this).F37()
				})(), 14)
				(_nw71).ArraySet1CodePoint(p1, 15)
				(_nw71).ArraySet1CodePoint(((_567_v107).F36()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_578_v118).Cardinality()), _dafny.IntOfUint32(((_567_v107).F36()).Cardinality()))).Uint32()).(_dafny.CodePoint), 16)
				(_nw71).ArraySet1CodePoint((_this).F37(), 17)
				(_nw71).ArraySet1CodePoint(p1, 18)
				(_nw71).ArraySet1CodePoint((_this).F37(), 19)
				(_nw71).ArraySet1CodePoint(_dafny.CodePoint('y'), 20)
				(_nw71).ArraySet1CodePoint(p1, 21)
				(_nw71).ArraySet1CodePoint((_this).F37(), 22)
				(_nw71).ArraySet1CodePoint((_this).F37(), 23)
				(_nw71).ArraySet1CodePoint((_this).F37(), 24)
				(_nw71).ArraySet1CodePoint((_this).F37(), 25)
				(_nw71).ArraySet1CodePoint((_this).F37(), 26)
				(_nw71).ArraySet1CodePoint((_this).F37(), 27)
				_579_v119 = _nw71
				_579_v119 = _579_v119
				(globalState).F4 = (func() _dafny.Int {
					if (_this).F28() {
						return (_this).Fm20(globalState)
					}
					return p0
				})()
			}
		}
		var _580_v120 D0
		_ = _580_v120
		_580_v120 = Companion_D0_.Create_DC3_(p3, !((_this).F28()), (_this).F28())
		var _581_v121 _dafny.MultiSet
		_ = _581_v121
		_581_v121 = _dafny.MultiSetOf((_580_v120).Dtor_cf8(), _this.F38)
		var _582_v122 _dafny.Map
		_ = _582_v122
		_582_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_581_v121).Cardinality(), (_dafny.Zero).Minus(p0))
		_582_v122 = ((_582_v122).Merge(_582_v122)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm20(globalState), p0))
		if _this.F38 {
			var _583_v123 _dafny.Sequence
			_ = _583_v123
			_583_v123 = _dafny.UnicodeSeqOfUtf8Bytes("vkegrqm")
			(globalState).F21 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(314), _dafny.IntOfUint32((_583_v123).Cardinality()))
			var _584_v124 _dafny.Sequence
			_ = _584_v124
			_584_v124 = _dafny.SeqOf((_this).Fm20(globalState))
			(globalState).F13 = (p0).Minus(_dafny.IntOfUint32((_584_v124).Cardinality()))
			var _585_v125 _dafny.Array
			_ = _585_v125
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw72
			_585_v125 = _nw72
			_585_v125 = p2
			var _586_v126 _dafny.Sequence
			_ = _586_v126
			_586_v126 = _dafny.SeqOf((_this).F28())
			var _587_v127 _dafny.Sequence
			_ = _587_v127
			_587_v127 = _dafny.SeqOf(_586_v126)
			var _588_v128 _dafny.Map
			_ = _588_v128
			_588_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_587_v127, _587_v127), (p3).Times(_dafny.IntOfUint32((_583_v123).Cardinality())))
			var _589_v129 _dafny.Map
			_ = _589_v129
			_589_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _587_v127)
			var _590_v130 _dafny.Array
			_ = _590_v130
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_7
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = func(_591_i4 _dafny.Int) _dafny.Int {
					return (_591_i4).Plus(_dafny.IntOfInt64(199))
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw73 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw73).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw73).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_590_v130 = _nw73
			var _592_v131 _dafny.MultiSet
			_ = _592_v131
			_592_v131 = _dafny.MultiSetOf(_590_v130, _590_v130, _590_v130)
			var _593_v132 _dafny.Sequence
			_ = _593_v132
			_593_v132 = _dafny.SeqOf(_592_v131)
			_588_v128 = (_588_v128).Update(_dafny.Companion_Sequence_.Concatenate(_587_v127, (func() _dafny.Sequence {
				if (_589_v129).Contains(_this.F38) {
					return (_589_v129).Get(_this.F38).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_586_v126, _586_v126)
			})()), ((_dafny.MultiSetOf(_590_v130, _590_v130)).Union((_593_v132).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_583_v123).Cardinality()), _dafny.IntOfUint32((_593_v132).Cardinality()))).Uint32()).(_dafny.MultiSet))).Cardinality())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_590_v130), 0))
			_ = _index79
			(_590_v130).ArraySet1(p3, (_index79).Int())
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_590_v130), 0))
			_ = _index80
			(_590_v130).ArraySet1(Companion_Default___.SafeDivisionInt(p3, _dafny.IntOfInt64(-240)), (_index80).Int())
		} else {
			var _594_v133 _dafny.Array
			_ = _594_v133
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_8
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = (func(_595_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_596_i5 _dafny.Int) _dafny.Int {
						return (_596_i5).Times(_595_p0)
					}
				})(p0)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw74 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw74).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw74).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_594_v133 = _nw74
			_594_v133 = _594_v133
			var _597_v134 _dafny.Array
			_ = _597_v134
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
			_ = _nw75
			_597_v134 = _nw75
			_597_v134 = _597_v134
			var _598_v135 _dafny.Array
			_ = _598_v135
			var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw76
			_598_v135 = _nw76
			_598_v135 = _598_v135
			(globalState).F4 = p0
			var _599_v136 _dafny.Sequence
			_ = _599_v136
			_599_v136 = _dafny.SeqOf(_this.F38)
			if (_dafny.IntOfUint32((_599_v136).Cardinality())).Cmp(p3) != 0 {
				var _600_v137 _dafny.Map
				_ = _600_v137
				_600_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !((_this).F28()) || (_this.F38))
				_600_v137 = (_600_v137).Update(false, (_this).F28())
				var _601_v138 _dafny.Array
				_ = _601_v138
				var _nwElement0_13 _dafny.Sequence = (func() _dafny.Sequence {
					if Companion_Default___.Fm22((_this).F28(), p0, globalState) {
						return _599_v136
					}
					return _599_v136
				})()
				_ = _nwElement0_13
				var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
				_ = _nw77
				(_nw77).ArraySet1(_nwElement0_13, 0)
				(_nw77).ArraySet1(_dafny.SeqOf(_this.F38, (_this).F28()), 1)
				(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_599_v136, _599_v136), 2)
				(_nw77).ArraySet1(_dafny.SeqOf(_this.F38, _this.F38), 3)
				(_nw77).ArraySet1(_599_v136, 4)
				(_nw77).ArraySet1(_599_v136, 5)
				(_nw77).ArraySet1(_599_v136, 6)
				_601_v138 = _nw77
				var _602_v139 _dafny.Sequence
				_ = _602_v139
				_602_v139 = _dafny.UnicodeSeqOfUtf8Bytes("kkbyosc")
				var _nwElement0_14 _dafny.Sequence = _599_v136
				_ = _nwElement0_14
				var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(15))
				_ = _nw78
				(_nw78).ArraySet1(_nwElement0_14, 0)
				(_nw78).ArraySet1(_dafny.SeqOf(_this.F38), 1)
				(_nw78).ArraySet1(_599_v136, 2)
				(_nw78).ArraySet1(_599_v136, 3)
				(_nw78).ArraySet1(Companion_Default___.Fm25((_dafny.Zero).Minus(p3), (_dafny.Zero).Minus(p0), p0, globalState), 4)
				(_nw78).ArraySet1(_599_v136, 5)
				(_nw78).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_599_v136, _599_v136), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_602_v139).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_599_v136, _599_v136)).Cardinality()))).Uint32(), (_this).F28()), 6)
				(_nw78).ArraySet1(_599_v136, 7)
				(_nw78).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_599_v136, _599_v136), 8)
				(_nw78).ArraySet1(_599_v136, 9)
				(_nw78).ArraySet1(_599_v136, 10)
				(_nw78).ArraySet1(_599_v136, 11)
				(_nw78).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_599_v136, _599_v136), 12)
				(_nw78).ArraySet1(_599_v136, 13)
				(_nw78).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), (_this).F28(), (_this).F28()), _dafny.SeqOf(_this.F38, false)), 14)
				_601_v138 = _nw78
				var _603_v140 _dafny.MultiSet
				_ = _603_v140
				_603_v140 = _dafny.MultiSetOf(p0, p3)
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_594_v133), 0))
				_ = _index81
				(_594_v133).ArraySet1((_603_v140).Cardinality(), (_index81).Int())
				var _604_v141 _dafny.Array
				_ = _604_v141
				var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw79
				_604_v141 = _nw79
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_604_v141), 0))
				_ = _index82
				(_604_v141).ArraySet1(_594_v133, (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_594_v133), 0))
				_ = _index83
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_604_v141), 0))
				_ = _index84
				var _rhs101 _dafny.Int = p0
				_ = _rhs101
				var _rhs102 _dafny.Int = p3
				_ = _rhs102
				var _rhs103 _dafny.Int = p3
				_ = _rhs103
				var _rhs104 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if Companion_Default___.Fm22((_this).F28(), p0, globalState) {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F28(), (_599_v136).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.IntOfUint32((_599_v136).Cardinality()))).Uint32()).(bool), !(Companion_Default___.Fm22((_this).F28(), p0, globalState))), _599_v136)
					}
					return _599_v136
				})()).Cardinality())
				_ = _rhs104
				var _rhs105 _dafny.Array = _594_v133
				_ = _rhs105
				var _lhs82 _dafny.Array = _594_v133
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_594_v133), 0))
				_ = _lhs83
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = _604_v141
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_604_v141), 0))
				_ = _lhs88
				(_lhs82).ArraySet1(_rhs101, (_lhs83).Int())
				_lhs84.F13 = _rhs102
				_lhs85.F21 = _rhs103
				_lhs86.F5 = _rhs104
				(_lhs87).ArraySet1(_rhs105, (_lhs88).Int())
				(globalState).F5 = (_594_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_594_v133), 0))).Int()).(_dafny.Int)
				var _605_v142 T1
				_ = _605_v142
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__(_602_v139, (_594_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_594_v133), 0))).Int()).(_dafny.Int), (_this).F27(), (_this).F28())
				_605_v142 = _nw80
				var _606_v143 _dafny.Map
				_ = _606_v143
				_606_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v142, _dafny.IntOfInt64(542))
				var _607_v144 _dafny.Map
				_ = _607_v144
				_607_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_v143, _this.F38)
				var _608_v145 _dafny.Sequence
				_ = _608_v145
				_608_v145 = _dafny.SeqOf(p3)
				var _609_v146 D0
				_ = _609_v146
				_609_v146 = Companion_D0_.Create_DC4_(Companion_D0_.Create_DC0_(_607_v144, (_this).F28(), (_608_v145).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.IntOfUint32((_608_v145).Cardinality()))).Uint32()).(_dafny.Int), (_this).F37(), (_605_v142).F28()))
				var _rhs106 _dafny.Int = p3
				_ = _rhs106
				var _rhs107 D0 = _609_v146
				_ = _rhs107
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				_lhs89.F5 = _rhs106
				_609_v146 = _rhs107
			} else {
				var _610_v147 _dafny.Map
				_ = _610_v147
				_610_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F28()), (_this).F28())
				(globalState).F21 = Companion_Default___.SafeModuloInt((_610_v147).Cardinality(), p0)
				var _611_v148 _dafny.Sequence
				_ = _611_v148
				_611_v148 = _dafny.SeqOf((_this).F27(), (_this).F27(), (_this).F27(), (_this).F27())
				var _612_v149 _dafny.Array
				_ = _612_v149
				var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
				_ = _nw81
				_612_v149 = _nw81
				var _rhs108 _dafny.Sequence = _611_v148
				_ = _rhs108
				var _rhs109 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("syuird")).Cardinality()), p0)).Cardinality()), p0)).Minus(p3)
				_ = _rhs109
				var _rhs110 _dafny.Int = (_dafny.Zero).Minus(p3)
				_ = _rhs110
				var _rhs111 _dafny.Array = _612_v149
				_ = _rhs111
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				_611_v148 = _rhs108
				_lhs90.F21 = _rhs109
				_lhs91.F5 = _rhs110
				_612_v149 = _rhs111
				(globalState).F13 = p0
				(_this).F38 = _this.F38
				(globalState).F6 = _this.F38
			}
		}
		var _613_v150 _dafny.Array
		_ = _613_v150
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_9
		var _nw82 _dafny.Array
		_ = _nw82
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw82 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Sequence = (func(_614_p0 _dafny.Int, _615_p3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_616_i7 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqOf(_614_p0)))), _dafny.SeqOf(Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(90))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}((func(_617_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_618_i8 _dafny.Int) _dafny.Int {
							return _617_p3
						}
					})(_615_p3))), _dafny.SeqOf(_615_p3, _615_p3))), Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(325))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg43 _dafny.Int) interface{} {
							return coer43(arg43)
						}
					}((func(_619_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_620_i9 _dafny.Int) _dafny.Int {
							return _619_p0
						}
					})(_614_p0))), _dafny.SeqOf(_614_p0, _614_p0, _615_p3, _615_p3, _615_p3))), Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqOf(_614_p0, _614_p0)))))
				}
			})(p0, p3)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw82 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw82).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw82).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_613_v150 = _nw82
		for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_613_v150), 0))); ; {
			_guard_loop_1, _ok43 := _iter43()
			if !_ok43 {
				break
			}
			var _621_i6 _dafny.Int
			_621_i6 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_621_i6).Sign() != -1) && ((_621_i6).Cmp(_dafny.ArrayLen((_613_v150), 0)) < 0)) {
				(_613_v150).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer44 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_622_p0 _dafny.Int, _623_p3 _dafny.Int) func(_dafny.Int) D1 {
					return func(_624_i10 _dafny.Int) D1 {
						return Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqOf(_622_p0, _622_p0, _623_p3)))
					}
				})(p0, p3))), (_621_i6).Int())
			}
		}
		if !((_this).F28()) {
			(globalState).F5 = p3
			var _625_v151 _dafny.Set
			_ = _625_v151
			_625_v151 = _dafny.SetOf(p3)
			if (func() _dafny.Set {
				var _coll42 = _dafny.NewBuilder()
				_ = _coll42
				for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-827), _dafny.IntOfInt64(751))); ; {
					_compr_42, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _626_v153 _dafny.Int
					_626_v153 = interface{}(_compr_42).(_dafny.Int)
					if ((_dafny.IntOfInt64(-827)).Cmp(_626_v153) <= 0) && ((_626_v153).Cmp(_dafny.IntOfInt64(751)) < 0) {
						_coll42.Add(Companion_Default___.SafeDivisionInt(_626_v153, p3))
					}
				}
				return _coll42.ToSet()
			}()).IsProperSubsetOf(func() _dafny.Set {
				var _coll43 = _dafny.NewBuilder()
				_ = _coll43
				for _iter45 := _dafny.Iterate((_625_v151).Elements()); ; {
					_compr_43, _ok45 := _iter45()
					if !_ok45 {
						break
					}
					var _627_v152 _dafny.Int
					_627_v152 = interface{}(_compr_43).(_dafny.Int)
					if (_625_v151).Contains(_627_v152) {
						_coll43.Add((_627_v152).Plus(_dafny.IntOfInt64(-414)))
					}
				}
				return _coll43.ToSet()
			}()) {
				var _628_v154 _dafny.Array
				_ = _628_v154
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_10
				var _nw83 _dafny.Array
				_ = _nw83
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw83 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = func(_629_i11 _dafny.Int) _dafny.Int {
						return (_629_i11).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}(func(_630_i12 _dafny.Int) _dafny.CodePoint {
							return (_this).F37()
						}))).Cardinality()))
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw83 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw83).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw83).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_628_v154 = _nw83
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_628_v154), 0))
				_ = _index85
				(_628_v154).ArraySet1(_dafny.IntOfInt64(29), (_index85).Int())
				var _631_v155 _dafny.Sequence
				_ = _631_v155
				_631_v155 = _dafny.UnicodeSeqOfUtf8Bytes("mrefi")
				var _632_v156 _dafny.Map
				_ = _632_v156
				_632_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_631_v155).Cardinality()), _this.F38)
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_628_v154), 0))
				_ = _index86
				(_628_v154).ArraySet1((_632_v156).Cardinality(), (_index86).Int())
				var _633_v157 _dafny.Array
				_ = _633_v157
				_633_v157 = (_this).F27()
				var _634_v158 _dafny.Map
				_ = _634_v158
				_634_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_633_v157, _580_v120)
				var _635_v159 _dafny.Map
				_ = _635_v159
				_635_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_633_v157, _580_v120)).Merge(_634_v158), (p3).Minus(p3))
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_628_v154), 0))
				_ = _index87
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_628_v154), 0))
				_ = _index88
				var _rhs112 _dafny.Int = p3
				_ = _rhs112
				var _rhs113 _dafny.Int = p0
				_ = _rhs113
				var _rhs114 _dafny.Int = p0
				_ = _rhs114
				var _rhs115 _dafny.Map = (_635_v159).Update(_634_v158, p3)
				_ = _rhs115
				var _rhs116 _dafny.Int = p0
				_ = _rhs116
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 _dafny.Array = _628_v154
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_628_v154), 0))
				_ = _lhs94
				var _lhs95 _dafny.Array = _628_v154
				_ = _lhs95
				var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_628_v154), 0))
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				_lhs92.F4 = _rhs112
				(_lhs93).ArraySet1(_rhs113, (_lhs94).Int())
				(_lhs95).ArraySet1(_rhs114, (_lhs96).Int())
				_635_v159 = _rhs115
				_lhs97.F4 = _rhs116
				var _636_v160 _dafny.Map
				_ = _636_v160
				var _637_v161 _dafny.Int
				_ = _637_v161
				var _out11 _dafny.Map
				_ = _out11
				var _out12 _dafny.Int
				_ = _out12
				_out11, _out12 = Companion_Default___.M0(globalState)
				_636_v160 = _out11
				_637_v161 = _out12
				(globalState).F21 = (p3).Times((_628_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_628_v154), 0))).Int()).(_dafny.Int))
				var _638_v162 _dafny.Map
				_ = _638_v162
				_638_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC6_(_637_v161), false)
				var _639_v163 D1
				_ = _639_v163
				_639_v163 = Companion_D1_.Create_DC6_(p0)
				var _640_v164 _dafny.Set
				_ = _640_v164
				_640_v164 = _dafny.SetOf(_this.F38)
				(globalState).F17 = (func() bool {
					if (_638_v162).Contains(_639_v163) {
						return (_638_v162).Get(_639_v163).(bool)
					}
					return (Companion_Default___.Fm31(globalState)).Equals(_640_v164)
				})()
				var _641_v165 _dafny.Array
				_ = _641_v165
				var _len0_11 _dafny.Int = _dafny.One
				_ = _len0_11
				var _nw84 _dafny.Array
				_ = _nw84
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw84 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = func(_642_i13 _dafny.Int) bool {
						return (_this).F28()
					}
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw84 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw84).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw84).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_641_v165 = _nw84
				_641_v165 = (_this).F27()
			} else {
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index89
				((_this).F27()).ArraySet1((_this).F28(), (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index90
				((_this).F27()).ArraySet1((Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p0))).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
					if (_582_v122).Contains(p3) {
						return (_582_v122).Get(p3).(_dafny.Int)
					}
					return p3
				})())) < 0, (_index90).Int())
				var _643_v166 _dafny.MultiSet
				_ = _643_v166
				_643_v166 = _dafny.MultiSetOf(_dafny.CodePoint('m'))
				var _644_v167 _dafny.Sequence
				_ = _644_v167
				_644_v167 = _dafny.SeqOf(_643_v166, _643_v166)
				var _645_v168 _dafny.Array
				_ = _645_v168
				var _nwElement0_15 bool = _this.F38
				_ = _nwElement0_15
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(13))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_15, 0)
				(_nw85).ArraySet1(_this.F38, 1)
				(_nw85).ArraySet1((_this).F28(), 2)
				(_nw85).ArraySet1(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), 3)
				(_nw85).ArraySet1(!(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool)), 4)
				(_nw85).ArraySet1(_this.F38, 5)
				(_nw85).ArraySet1(_this.F38, 6)
				(_nw85).ArraySet1(true, 7)
				(_nw85).ArraySet1(!((_this).F28()), 8)
				(_nw85).ArraySet1(!((_this).F28()), 9)
				(_nw85).ArraySet1(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), 10)
				(_nw85).ArraySet1((_this).F28(), 11)
				(_nw85).ArraySet1(_this.F38, 12)
				_645_v168 = _nw85
				var _646_v169 *C0
				_ = _646_v169
				var _nw86 *C0 = New_C0_()
				_ = _nw86
				_nw86.Ctor__(Companion_Default___.Fm32(p0, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), globalState), (_dafny.Zero).Minus((_dafny.IntOfUint32((_644_v167).Cardinality())).Minus((_dafny.Zero).Minus(p0))), _645_v168, (_this).F28())
				_646_v169 = _nw86
				(globalState).F13 = p0
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p2), 0))
				_ = _index91
				(p2).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("k"), (_646_v169).F36()), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("k"), (_646_v169).F36())).Cardinality()))).Uint32(), p1), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p2), 0))
				_ = _index92
				(p2).ArraySet1((_646_v169).F36(), (_index92).Int())
				(globalState).F6 = (func() bool {
					if ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool) {
						return _this.F38
					}
					return _this.F38
				})()
			}
			var _647_v170 _dafny.Sequence
			_ = _647_v170
			_647_v170 = _dafny.UnicodeSeqOfUtf8Bytes("pteeog")
			var _648_v171 *C0
			_ = _648_v171
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__(_647_v170, Companion_Default___.SafeDivisionInt(p3, p3), (_this).F27(), _this.F38)
			_648_v171 = _nw87
			var _649_v172 _dafny.Map
			_ = _649_v172
			_649_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())
			var _650_v173 _dafny.Sequence
			_ = _650_v173
			_650_v173 = _dafny.SeqOf(Companion_Default___.Fm22((func() bool {
				if (_649_v172).Contains(p0) {
					return (_649_v172).Get(p0).(bool)
				}
				return _this.F38
			})(), p0, globalState), (_this).F28())
			if ((_dafny.IntOfUint32((_650_v173).Cardinality())).Cmp(_dafny.IntOfInt64(695)) > 0) || ((_this).F28()) {
				var _651_v174 _dafny.Array
				_ = _651_v174
				var _nwElement0_16 _dafny.MultiSet = (_581_v121).Difference(_581_v121)
				_ = _nwElement0_16
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(14))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_16, 0)
				(_nw88).ArraySet1(_581_v121, 1)
				(_nw88).ArraySet1(_581_v121, 2)
				(_nw88).ArraySet1(_581_v121, 3)
				(_nw88).ArraySet1(Companion_Default___.Fm26((_this).F28(), p3, p1, globalState), 4)
				(_nw88).ArraySet1(_dafny.MultiSetFromSeq(_650_v173), 5)
				(_nw88).ArraySet1(_581_v121, 6)
				(_nw88).ArraySet1(_581_v121, 7)
				(_nw88).ArraySet1(_581_v121, 8)
				(_nw88).ArraySet1((_dafny.MultiSetFromSeq(_650_v173)).Intersection(_581_v121), 9)
				(_nw88).ArraySet1(_581_v121, 10)
				(_nw88).ArraySet1(_581_v121, 11)
				(_nw88).ArraySet1(_581_v121, 12)
				(_nw88).ArraySet1(_dafny.MultiSetOf(_this.F38), 13)
				_651_v174 = _nw88
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_651_v174), 0))
				_ = _index93
				(_651_v174).ArraySet1(_dafny.MultiSetOf(_this.F38), (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_651_v174), 0))
				_ = _index94
				(_651_v174).ArraySet1((func() _dafny.MultiSet {
					if _this.F38 {
						return _581_v121
					}
					return (func() _dafny.MultiSet {
						if (_650_v173).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_650_v173).Cardinality()))).Uint32()).(bool) {
							return _dafny.MultiSetFromSeq(Companion_Default___.Fm25(p0, p0, p0, globalState))
						}
						return _dafny.MultiSetOf(false, _this.F38, (_this).F28(), (_this).F28(), (_this).F28())
					})()
				})(), (_index94).Int())
				var _652_v175 _dafny.Array
				_ = _652_v175
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw89
				_652_v175 = _nw89
				var _653_v176 _dafny.Array
				_ = _653_v176
				var _nwElement0_17 _dafny.Array = _652_v175
				_ = _nwElement0_17
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(2))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_17, 0)
				(_nw90).ArraySet1(_652_v175, 1)
				_653_v176 = _nw90
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_653_v176), 0))
				_ = _index95
				(_653_v176).ArraySet1(_652_v175, (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_653_v176), 0))
				_ = _index96
				(_653_v176).ArraySet1(_652_v175, (_index96).Int())
				var _654_v177 _dafny.Map
				_ = _654_v177
				_654_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm22(Companion_Default___.Fm22((_this).F28(), (_625_v151).Cardinality(), globalState), p3, globalState))
				var _655_v178 _dafny.Set
				_ = _655_v178
				_655_v178 = _dafny.SetOf(_654_v177, _654_v177, _654_v177)
				(globalState).F4 = ((_655_v178).Difference((_dafny.SetOf(_654_v177, _654_v177)).Union(_655_v178))).Cardinality()
				(globalState).F17 = (p0).Cmp(p0) <= 0
				(globalState).F5 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(515))).Uint32(), func(coer46 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_656_v121 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_657_i14 _dafny.Int) _dafny.MultiSet {
						return _656_v121
					}
				})(_581_v121)))).Cardinality())
			} else {
				var _658_v179 _dafny.Array
				_ = _658_v179
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw91
				_658_v179 = _nw91
				var _659_v180 _dafny.Map
				_ = _659_v180
				_659_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _658_v179)
				_659_v180 = (_659_v180).Update(Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(p3)), _658_v179)
				var _660_v181 _dafny.Sequence
				_ = _660_v181
				_660_v181 = _dafny.SeqOf(p3, (_dafny.Zero).Minus(p3), p3)
				var _661_v182 _dafny.MultiSet
				_ = _661_v182
				_661_v182 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_660_v181)).Cardinality(), Companion_Default___.SafeModuloInt(p3, p3))
				var _662_v183 _dafny.Array
				_ = _662_v183
				var _nw92 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw92
				_662_v183 = _nw92
				var _663_v184 _dafny.Sequence
				_ = _663_v184
				_663_v184 = _dafny.SeqOf(_581_v121)
				var _664_v185 _dafny.Array
				_ = _664_v185
				var _nwElement0_18 _dafny.MultiSet = _581_v121
				_ = _nwElement0_18
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(18))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_18, 0)
				(_nw93).ArraySet1(_581_v121, 1)
				(_nw93).ArraySet1((_581_v121).Difference(_581_v121), 2)
				(_nw93).ArraySet1((_581_v121).Union(_581_v121), 3)
				(_nw93).ArraySet1((_581_v121).Union(_581_v121), 4)
				(_nw93).ArraySet1((_663_v184).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_647_v170).Cardinality()), _dafny.IntOfUint32((_663_v184).Cardinality()))).Uint32()).(_dafny.MultiSet), 5)
				(_nw93).ArraySet1((_dafny.MultiSetOf(_this.F38, (_this).F28(), Companion_Default___.Fm22((_this).F28(), p0, globalState), (_this).F28())).Intersection(_dafny.MultiSetFromSeq(_650_v173)), 6)
				(_nw93).ArraySet1((func() _dafny.MultiSet {
					if !((_this).F28()) {
						return (_dafny.MultiSetFromSeq(_650_v173)).Update((_this).F28(), Companion_Default___.Abs(p0))
					}
					return _581_v121
				})(), 7)
				(_nw93).ArraySet1((func() _dafny.MultiSet {
					if _this.F38 {
						return _581_v121
					}
					return _581_v121
				})(), 8)
				(_nw93).ArraySet1(_581_v121, 9)
				(_nw93).ArraySet1(_dafny.MultiSetOf(_this.F38), 10)
				(_nw93).ArraySet1((_581_v121).Intersection(_581_v121), 11)
				(_nw93).ArraySet1((_581_v121).Union(_581_v121), 12)
				(_nw93).ArraySet1(_581_v121, 13)
				(_nw93).ArraySet1((_581_v121).Difference(_581_v121), 14)
				(_nw93).ArraySet1(_581_v121, 15)
				(_nw93).ArraySet1(_581_v121, 16)
				(_nw93).ArraySet1((_581_v121).Difference(_581_v121), 17)
				_664_v185 = _nw93
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_664_v185), 0))
				_ = _index97
				(_664_v185).ArraySet1(_581_v121, (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_664_v185), 0))
				_ = _index98
				var _rhs117 bool = ((_this).F28()) && (_this.F38)
				_ = _rhs117
				var _rhs118 _dafny.MultiSet = _661_v182
				_ = _rhs118
				var _rhs119 _dafny.Int = (_648_v171).Fm15(globalState)
				_ = _rhs119
				var _rhs120 _dafny.Array = (_this).F27()
				_ = _rhs120
				var _rhs121 _dafny.MultiSet = _581_v121
				_ = _rhs121
				var _lhs98 *C2 = _this
				_ = _lhs98
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				var _lhs100 _dafny.Array = _664_v185
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_664_v185), 0))
				_ = _lhs101
				_lhs98.F38 = _rhs117
				_661_v182 = _rhs118
				_lhs99.F4 = _rhs119
				_662_v183 = _rhs120
				(_lhs100).ArraySet1(_rhs121, (_lhs101).Int())
				var _rhs122 _dafny.Int = p3
				_ = _rhs122
				var _rhs123 _dafny.Map = _649_v172
				_ = _rhs123
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				_lhs102.F4 = _rhs122
				_649_v172 = _rhs123
				(globalState).F13 = Companion_Default___.SafeModuloInt((p0).Minus(p0), (_dafny.Zero).Minus(p3))
				var _665_v186 _dafny.Array
				_ = _665_v186
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_12
				var _nw94 _dafny.Array
				_ = _nw94
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw94 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.CodePoint = (func(_666_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_667_i15 _dafny.Int) _dafny.CodePoint {
							return _666_p1
						}
					})(p1)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw94 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw94).ArraySet1CodePoint(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw94).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_665_v186 = _nw94
				var _668_v187 _dafny.Sequence
				_ = _668_v187
				_668_v187 = _dafny.SeqOf(_665_v186)
				(globalState).F21 = _dafny.IntOfUint32((_668_v187).Cardinality())
			}
			var _669_v188 _dafny.Array
			_ = _669_v188
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_13
			var _nw95 _dafny.Array
			_ = _nw95
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw95 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Sequence = func(_670_i16 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer47 func(_dafny.Int) D5) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}(func(_671_i17 _dafny.Int) D5 {
						return Companion_D5_.Create_DC12_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(458))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
							return func(arg48 _dafny.Int) interface{} {
								return coer48(arg48)
							}
						}(func(_672_i18 _dafny.Int) _dafny.Set {
							return _dafny.SetOf((_this).F28(), (_this).F28(), _this.F38)
						})))
					}))
				}
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
			_669_v188 = _nw95
			_669_v188 = _669_v188
		} else {
			var _673_v189 _dafny.Sequence
			_ = _673_v189
			_673_v189 = _dafny.SeqOf(_this.F38, (_this).F28())
			var _674_v190 _dafny.MultiSet
			_ = _674_v190
			_674_v190 = _dafny.MultiSetOf(_dafny.IntOfUint32((_673_v189).Cardinality()), p3)
			var _675_v191 _dafny.Map
			_ = _675_v191
			_675_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, p3)
			var _676_v192 _dafny.Sequence
			_ = _676_v192
			_676_v192 = _dafny.SeqOf((_674_v190).Update((_675_v191).Cardinality(), Companion_Default___.Abs(p0)), _674_v190)
			(globalState).F6 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_674_v190, (_676_v192).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_676_v192).Cardinality()))).Uint32()).(_dafny.MultiSet)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(259))).Uint32(), func(coer49 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_677_v190 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_678_i19 _dafny.Int) _dafny.MultiSet {
					return _677_v190
				}
			})(_674_v190)))), _676_v192)
			var _679_v193 _dafny.Sequence
			_ = _679_v193
			_679_v193 = _dafny.UnicodeSeqOfUtf8Bytes("iiaprsa")
			var _680_v194 *C0
			_ = _680_v194
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__(_dafny.Companion_Sequence_.Concatenate(_679_v193, _679_v193), p0, (_this).F27(), _this.F38)
			_680_v194 = _nw96
			var _681_v195 _dafny.Array
			_ = _681_v195
			var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw97
			_681_v195 = _nw97
			var _682_v196 _dafny.Sequence
			_ = _682_v196
			_682_v196 = _dafny.SeqOf(_681_v195)
			var _683_v197 _dafny.Array
			_ = _683_v197
			var _nwElement0_19 _dafny.Array = _681_v195
			_ = _nwElement0_19
			var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(20))
			_ = _nw98
			(_nw98).ArraySet1(_nwElement0_19, 0)
			(_nw98).ArraySet1(_681_v195, 1)
			(_nw98).ArraySet1(_681_v195, 2)
			(_nw98).ArraySet1(_681_v195, 3)
			(_nw98).ArraySet1(_681_v195, 4)
			(_nw98).ArraySet1(_681_v195, 5)
			(_nw98).ArraySet1(_681_v195, 6)
			(_nw98).ArraySet1(_681_v195, 7)
			(_nw98).ArraySet1(_681_v195, 8)
			(_nw98).ArraySet1(_681_v195, 9)
			(_nw98).ArraySet1(_681_v195, 10)
			(_nw98).ArraySet1(_681_v195, 11)
			(_nw98).ArraySet1(_681_v195, 12)
			(_nw98).ArraySet1(_681_v195, 13)
			(_nw98).ArraySet1(_681_v195, 14)
			(_nw98).ArraySet1((_682_v196).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_682_v196).Cardinality()))).Uint32()).(_dafny.Array), 15)
			(_nw98).ArraySet1((func() _dafny.Array {
				if !(_this.F38) {
					return (_682_v196).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_682_v196).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _681_v195
			})(), 16)
			(_nw98).ArraySet1(_681_v195, 17)
			(_nw98).ArraySet1(_681_v195, 18)
			(_nw98).ArraySet1(_681_v195, 19)
			_683_v197 = _nw98
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_683_v197), 0))
			_ = _index99
			(_683_v197).ArraySet1(_681_v195, (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_683_v197), 0))
			_ = _index100
			var _rhs124 bool = (_this).F28()
			_ = _rhs124
			var _rhs125 _dafny.Array = _681_v195
			_ = _rhs125
			var _rhs126 bool = _this.F38
			_ = _rhs126
			var _lhs103 *GlobalState = globalState
			_ = _lhs103
			var _lhs104 _dafny.Array = _683_v197
			_ = _lhs104
			var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_683_v197), 0))
			_ = _lhs105
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			_lhs103.F17 = _rhs124
			(_lhs104).ArraySet1(_rhs125, (_lhs105).Int())
			_lhs106.F6 = _rhs126
			var _684_v198 _dafny.Sequence
			_ = _684_v198
			_684_v198 = _dafny.SeqOf((_580_v120).Dtor_cf7(), p3, p0)
			var _685_v199 _dafny.Sequence
			_ = _685_v199
			_685_v199 = _684_v198
			var _686_v200 _dafny.Sequence
			_ = _686_v200
			_686_v200 = (func() _dafny.Sequence {
				if _this.F38 {
					return _684_v198
				}
				return (_685_v199)
			})()
			var _source9 _dafny.Sequence = _685_v199
			_ = _source9
			var _687___mcc_h15 _dafny.Sequence = _source9
			_ = _687___mcc_h15
			var _688_cf15 _dafny.Sequence = _687___mcc_h15
			_ = _688_cf15
			var _689_v201 _dafny.Array
			_ = _689_v201
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_14
			var _nw99 _dafny.Array
			_ = _nw99
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw99 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.CodePoint = (func(_690_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_691_i20 _dafny.Int) _dafny.CodePoint {
						return _690_p1
					}
				})(p1)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw99 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw99).ArraySet1CodePoint(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw99).ArraySet1CodePoint(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_689_v201 = _nw99
			var _692_v202 _dafny.Map
			_ = _692_v202
			_692_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if false {
					return _689_v201
				}
				return _689_v201
			})(), _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm32(p3, !(_this.F38), globalState), (_680_v194).F36()))
			(globalState).F17 = (func() bool {
				if (_692_v202).Contains(_689_v201) {
					return (_692_v202).Get(_689_v201).(bool)
				}
				return (func() bool {
					if (_this).F28() {
						return false
					}
					return (_this).F28()
				})()
			})()
			var _693_v203 _dafny.Set
			_ = _693_v203
			_693_v203 = _dafny.SetOf((_this).F37())
			var _694_v204 _dafny.Map
			_ = _694_v204
			_694_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v203, (_this).F28())
			var _695_v205 _dafny.Map
			_ = _695_v205
			_695_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p3).Cmp(p3) >= 0, (_694_v204).Merge(_694_v204))
			_695_v205 = (_695_v205).Update(((_dafny.Zero).Minus((_this).Fm20(globalState))).Cmp(p3) == 0, _694_v204)
			var _696_v206 _dafny.Map
			_ = _696_v206
			_696_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sgpyl")).Cardinality()), p1)
			var _rhs127 _dafny.Int = Companion_Default___.SafeModuloInt(p3, _dafny.IntOfUint32((_688_cf15).Cardinality()))
			_ = _rhs127
			var _rhs128 _dafny.Int = (func() _dafny.Int {
				if _this.F38 {
					return p0
				}
				return ((_dafny.Zero).Minus((_this).Fm20(globalState))).Times((_696_v206).Cardinality())
			})()
			_ = _rhs128
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			var _lhs108 *GlobalState = globalState
			_ = _lhs108
			_lhs107.F5 = _rhs127
			_lhs108.F13 = _rhs128
			(globalState).F21 = p0
			var _rhs129 _dafny.Int = p0
			_ = _rhs129
			var _rhs130 bool = (_673_v189).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.IntOfUint32((_673_v189).Cardinality()))).Uint32()).(bool)
			_ = _rhs130
			var _lhs109 *GlobalState = globalState
			_ = _lhs109
			_lhs109.F13 = _rhs129
			r0 = _rhs130
		}
		r0 = (p3).Cmp(p3) == 0
		return r0
	}
}
func (_this *C2) M13(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) D0 {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var _697_i0 _dafny.Int
		_ = _697_i0
		_697_i0 = _dafny.Zero
		{
			for (p1) || (!(p0)) {
				{
					if (_697_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_697_i0 = (_697_i0).Plus(_dafny.One)
					var _698_v0 T1
					_ = _698_v0
					var _nw100 *C0 = New_C0_()
					_ = _nw100
					_nw100.Ctor__(p3, p2, (_this).F27(), p1)
					_698_v0 = _nw100
					var _699_v1 _dafny.MultiSet
					_ = _699_v1
					_699_v1 = _dafny.MultiSetOf(_dafny.IntOfInt64(548))
					var _700_v2 _dafny.Map
					_ = _700_v2
					_700_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_698_v0, (_dafny.Zero).Minus((_699_v1).Cardinality()))
					var _701_v3 _dafny.Map
					_ = _701_v3
					_701_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_700_v2, p0)
					var _702_v4 D0
					_ = _702_v4
					_702_v4 = Companion_D0_.Create_DC0_(_701_v3, (_699_v1).IsProperSubsetOf(_699_v1), (_dafny.SetOf(p2)).Cardinality(), (_this).F37(), (_this).F28())
					var _703_v5 _dafny.Set
					_ = _703_v5
					_703_v5 = _dafny.SetOf(_dafny.IntOfUint32((p3).Cardinality()))
					var _704_v8 _dafny.Map
					_ = _704_v8
					_704_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_698_v0.F29(), func() _dafny.Set {
						var _coll44 = _dafny.NewBuilder()
						_ = _coll44
						for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(524), _dafny.IntOfInt64(662))); ; {
							_compr_44, _ok46 := _iter46()
							if !_ok46 {
								break
							}
							var _705_v7 _dafny.Int
							_705_v7 = interface{}(_compr_44).(_dafny.Int)
							if ((_dafny.IntOfInt64(524)).Cmp(_705_v7) <= 0) && ((_705_v7).Cmp(_dafny.IntOfInt64(662)) < 0) {
								_coll44.Add((_705_v7).Plus(p2))
							}
						}
						return _coll44.ToSet()
					}())
					var _706_v9 _dafny.Array
					_ = _706_v9
					var _nwElement0_20 _dafny.Set = _703_v5
					_ = _nwElement0_20
					var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(16))
					_ = _nw101
					(_nw101).ArraySet1(_nwElement0_20, 0)
					(_nw101).ArraySet1((_dafny.SetOf(p2)).Union(_703_v5), 1)
					(_nw101).ArraySet1(_703_v5, 2)
					(_nw101).ArraySet1(func() _dafny.Set {
						var _coll45 = _dafny.NewBuilder()
						_ = _coll45
						for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(182), _dafny.IntOfInt64(243))); ; {
							_compr_45, _ok47 := _iter47()
							if !_ok47 {
								break
							}
							var _707_v6 _dafny.Int
							_707_v6 = interface{}(_compr_45).(_dafny.Int)
							if ((_dafny.IntOfInt64(182)).Cmp(_707_v6) <= 0) && ((_707_v6).Cmp(_dafny.IntOfInt64(243)) < 0) {
								_coll45.Add((_707_v6).Minus(_698_v0.F29()))
							}
						}
						return _coll45.ToSet()
					}(), 3)
					(_nw101).ArraySet1(_703_v5, 4)
					(_nw101).ArraySet1((Companion_Default___.Fm0(_704_v8, p0, _this.F38, globalState)).Difference(_703_v5), 5)
					(_nw101).ArraySet1(_703_v5, 6)
					(_nw101).ArraySet1(_703_v5, 7)
					(_nw101).ArraySet1(_703_v5, 8)
					(_nw101).ArraySet1(_703_v5, 9)
					(_nw101).ArraySet1(_703_v5, 10)
					(_nw101).ArraySet1(_703_v5, 11)
					(_nw101).ArraySet1(_703_v5, 12)
					(_nw101).ArraySet1(_dafny.SetOf(_dafny.IntOfInt64(-582)), 13)
					(_nw101).ArraySet1(_703_v5, 14)
					(_nw101).ArraySet1(_703_v5, 15)
					_706_v9 = _nw101
					var _708_v10 _dafny.Sequence
					_ = _708_v10
					_708_v10 = _dafny.SeqOf(_703_v5)
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_706_v9), 0))
					_ = _index101
					(_706_v9).ArraySet1((_708_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.IntOfUint32((_708_v10).Cardinality()))).Uint32()).(_dafny.Set), (_index101).Int())
					var _709_v11 _dafny.Array
					_ = _709_v11
					var _len0_15 _dafny.Int = _dafny.IntOfInt64(11)
					_ = _len0_15
					var _nw102 _dafny.Array
					_ = _nw102
					if _len0_15.Cmp(_dafny.Zero) == 0 {
						_nw102 = _dafny.NewArray(_len0_15)
					} else {
						var _init15 func(_dafny.Int) _dafny.Sequence = func(_710_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(503))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg50 _dafny.Int) interface{} {
									return coer50(arg50)
								}
							}(func(_711_i2 _dafny.Int) _dafny.CodePoint {
								return (_this).F37()
							}))
						}
						_ = _init15
						var _element0_15 = _init15(_dafny.Zero)
						_ = _element0_15
						_nw102 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
						(_nw102).ArraySet1(_element0_15, 0)
						var _nativeLen0_15 = (_len0_15).Int()
						_ = _nativeLen0_15
						for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
							(_nw102).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
						}
					}
					_709_v11 = _nw102
					var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_709_v11), 0))
					_ = _index102
					(_709_v11).ArraySet1(p3, (_index102).Int())
					var _712_v12 _dafny.Map
					_ = _712_v12
					_712_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(p2))
					var _713_v13 _dafny.Sequence
					_ = _713_v13
					_713_v13 = _dafny.SeqOf(Companion_Default___.Fm22(p0, (func() _dafny.Int {
						if (_712_v12).Contains((_this).F28()) {
							return (_712_v12).Get((_this).F28()).(_dafny.Int)
						}
						return _698_v0.F29()
					})(), globalState), (_698_v0).F28())
					var _714_v14 _dafny.Array
					_ = _714_v14
					var _len0_16 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_16
					var _nw103 _dafny.Array
					_ = _nw103
					if _len0_16.Cmp(_dafny.Zero) == 0 {
						_nw103 = _dafny.NewArray(_len0_16)
					} else {
						var _init16 func(_dafny.Int) _dafny.Int = (func(_715_v0 T1) func(_dafny.Int) _dafny.Int {
							return func(_716_i3 _dafny.Int) _dafny.Int {
								return (_716_i3).Times(_715_v0.F29())
							}
						})(_698_v0)
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
					_714_v14 = _nw103
					var _717_v15 _dafny.MultiSet
					_ = _717_v15
					_717_v15 = _dafny.MultiSetOf(_714_v14)
					var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_706_v9), 0))
					_ = _index103
					var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_709_v11), 0))
					_ = _index104
					var _rhs131 D0 = _702_v4
					_ = _rhs131
					var _rhs132 _dafny.Set = (_703_v5).Difference(_703_v5)
					_ = _rhs132
					var _rhs133 _dafny.Sequence = _dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex((_712_v12).Cardinality(), _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), (_this).F37())
					_ = _rhs133
					var _rhs134 bool = (_713_v13).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(p2, (_717_v15).Cardinality()), _dafny.IntOfUint32((_713_v13).Cardinality()))).Uint32()).(bool)
					_ = _rhs134
					var _lhs110 _dafny.Array = _706_v9
					_ = _lhs110
					var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_706_v9), 0))
					_ = _lhs111
					var _lhs112 _dafny.Array = _709_v11
					_ = _lhs112
					var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_709_v11), 0))
					_ = _lhs113
					var _lhs114 *GlobalState = globalState
					_ = _lhs114
					_702_v4 = _rhs131
					(_lhs110).ArraySet1(_rhs132, (_lhs111).Int())
					(_lhs112).ArraySet1(_rhs133, (_lhs113).Int())
					_lhs114.F15 = _rhs134
					_702_v4 = _702_v4
					(globalState).F5 = (_dafny.Zero).Minus(p2)
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_709_v11), 0))
					_ = _index105
					(_709_v11).ArraySet1(p3, (_index105).Int())
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		if p0 {
			if p0 {
				var _718_v16 _dafny.Sequence
				_ = _718_v16
				_718_v16 = _dafny.SeqOf(p2)
				var _719_v17 _dafny.Sequence
				_ = _719_v17
				_719_v17 = _dafny.SeqOf(_dafny.IntOfUint32((_718_v16).Cardinality()))
				var _720_v18 _dafny.Map
				_ = _720_v18
				_720_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_719_v17).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_719_v17).Cardinality()))).Uint32()).(_dafny.Int))
				(globalState).F15 = Companion_Default___.Fm22(true, (_dafny.Zero).Minus((func() _dafny.Int {
					if (_720_v18).Contains(true) {
						return (_720_v18).Get(true).(_dafny.Int)
					}
					return p2
				})()), globalState)
				(globalState).F17 = !_dafny.Companion_Sequence_.Equal(_718_v16, _718_v16)
				(globalState).F15 = (func() _dafny.Map {
					var _coll46 = _dafny.NewMapBuilder()
					_ = _coll46
					for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-895), _dafny.IntOfInt64(596))); ; {
						_compr_46, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _721_v19 _dafny.Int
						_721_v19 = interface{}(_compr_46).(_dafny.Int)
						if ((_dafny.IntOfInt64(-895)).Cmp(_721_v19) <= 0) && ((_721_v19).Cmp(_dafny.IntOfInt64(596)) < 0) {
							_coll46.Add((_721_v19).Minus(p2), _dafny.IntOfInt64(647))
						}
					}
					return _coll46.ToMap()
				}()).Contains(p2)
				var _722_v20 *C0
				_ = _722_v20
				var _nw104 *C0 = New_C0_()
				_ = _nw104
				_nw104.Ctor__(p3, (p2).Times(p2), (_this).F27(), false)
				_722_v20 = _nw104
				(globalState).F15 = false
			} else {
				var _723_v21 *C0
				_ = _723_v21
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__(p3, _dafny.IntOfInt64(-948), (_this).F27(), p1)
				_723_v21 = _nw105
				var _724_v22 _dafny.Map
				_ = _724_v22
				_724_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _725_v23 _dafny.Sequence
				_ = _725_v23
				_725_v23 = _dafny.SeqOf(Companion_Default___.Fm22(true, p2, globalState), (p2).Cmp((func() _dafny.Int {
					if (_724_v22).Contains(p2) {
						return (_724_v22).Get(p2).(_dafny.Int)
					}
					return p2
				})()) == 0, true)
				_725_v23 = _dafny.Companion_Sequence_.Update(_725_v23, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_725_v23).Cardinality()))).Uint32(), p1)
				var _726_v24 _dafny.Sequence
				_ = _726_v24
				_726_v24 = _dafny.SeqOf(p2, p2)
				var _727_v25 _dafny.Map
				_ = _727_v25
				_727_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_726_v24).Cardinality()))
				var _728_v26 _dafny.Map
				_ = _728_v26
				_728_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_727_v25, p2)
				var _729_v27 D8
				_ = _729_v27
				_729_v27 = Companion_D8_.Create_DC17_(_728_v26)
				var _730_v28 _dafny.Sequence
				_ = _730_v28
				_730_v28 = _dafny.SeqOf(_728_v26, _728_v26)
				var _731_v29 _dafny.Map
				_ = _731_v29
				_731_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, p0)
				var _732_v30 _dafny.Array
				_ = _732_v30
				var _nwElement0_21 _dafny.Int = (_dafny.Zero).Minus(p2)
				_ = _nwElement0_21
				var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(25))
				_ = _nw106
				(_nw106).ArraySet1(_nwElement0_21, 0)
				(_nw106).ArraySet1((_723_v21).Fm15(globalState), 1)
				(_nw106).ArraySet1(p2, 2)
				(_nw106).ArraySet1(p2, 3)
				(_nw106).ArraySet1(p2, 4)
				(_nw106).ArraySet1(p2, 5)
				(_nw106).ArraySet1(p2, 6)
				(_nw106).ArraySet1((func() _dafny.Int {
					if false {
						return p2
					}
					return _dafny.IntOfInt64(144)
				})(), 7)
				(_nw106).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm33((_dafny.Zero).Minus(p2), (_this).F28(), globalState)).Cardinality()), 8)
				(_nw106).ArraySet1((p2).Plus(p2), 9)
				(_nw106).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sdbbk")).Cardinality())).Plus(p2)), 10)
				(_nw106).ArraySet1(p2, 11)
				(_nw106).ArraySet1(p2, 12)
				(_nw106).ArraySet1((((_729_v27).Dtor_cf30()).Merge((_730_v28).Select((Companion_Default___.SafeIndex((_723_v21).Fm15(globalState), _dafny.IntOfUint32((_730_v28).Cardinality()))).Uint32()).(_dafny.Map))).Cardinality(), 13)
				(_nw106).ArraySet1(p2, 14)
				(_nw106).ArraySet1(p2, 15)
				(_nw106).ArraySet1(p2, 16)
				(_nw106).ArraySet1(p2, 17)
				(_nw106).ArraySet1((_726_v24).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_726_v24).Cardinality()))).Uint32()).(_dafny.Int), 18)
				(_nw106).ArraySet1(p2, 19)
				(_nw106).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_725_v23).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_725_v23).Cardinality()))).Uint32()).(bool), (_this).F28())).Cardinality()), 20)
				(_nw106).ArraySet1(p2, 21)
				(_nw106).ArraySet1(p2, 22)
				(_nw106).ArraySet1(p2, 23)
				(_nw106).ArraySet1(((_731_v29).Update(_this.F38, (_this).F28())).Cardinality(), 24)
				_732_v30 = _nw106
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_732_v30), 0))
				_ = _index106
				(_732_v30).ArraySet1(p2, (_index106).Int())
				var _733_v31 _dafny.MultiSet
				_ = _733_v31
				_733_v31 = _dafny.MultiSetOf(_dafny.IntOfUint32((_726_v24).Cardinality()), p2)
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_732_v30), 0))
				_ = _index107
				(_732_v30).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), (_733_v31).Cardinality()), Companion_Default___.SafeModuloInt(p2, p2)), (_index107).Int())
				(globalState).F4 = (_732_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_732_v30), 0))).Int()).(_dafny.Int)
				var _734_v32 T0
				_ = _734_v32
				var _nw107 *C1 = New_C1_()
				_ = _nw107
				_nw107.Ctor__(_this.F38, (_this).F27(), p1)
				_734_v32 = _nw107
				var _735_v33 _dafny.Map
				_ = _735_v33
				_735_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v32, (_731_v29).Update(false, p0))
				var _736_v34 _dafny.Set
				_ = _736_v34
				_736_v34 = _dafny.SetOf((_732_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_732_v30), 0))).Int()).(_dafny.Int), p2, (_735_v33).Cardinality(), p2)
				(globalState).F15 = (_dafny.SetOf(p2, p2, (_dafny.Zero).Minus((_732_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_732_v30), 0))).Int()).(_dafny.Int)))).IsSubsetOf(_736_v34)
			}
			var _737_v35 _dafny.Map
			_ = _737_v35
			_737_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), p2), p0)
			_737_v35 = (_737_v35).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iivinjgue")).Cardinality()), (_this).F28())
			var _738_v36 *C0
			_ = _738_v36
			var _nw108 *C0 = New_C0_()
			_ = _nw108
			_nw108.Ctor__(p3, p2, (_this).F27(), (_this).F28())
			_738_v36 = _nw108
			_738_v36 = _738_v36
			var _739_v37 _dafny.Array
			_ = _739_v37
			var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw109
			_739_v37 = _nw109
			var _nwElement0_22 _dafny.Array = _739_v37
			_ = _nwElement0_22
			var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(9))
			_ = _nw110
			(_nw110).ArraySet1(_nwElement0_22, 0)
			(_nw110).ArraySet1(_739_v37, 1)
			(_nw110).ArraySet1(_739_v37, 2)
			(_nw110).ArraySet1(_739_v37, 3)
			(_nw110).ArraySet1(_739_v37, 4)
			(_nw110).ArraySet1(_739_v37, 5)
			(_nw110).ArraySet1(_739_v37, 6)
			(_nw110).ArraySet1(_739_v37, 7)
			(_nw110).ArraySet1(_739_v37, 8)
			(globalState).F18 = _nw110
			var _740_v38 _dafny.Sequence
			_ = _740_v38
			_740_v38 = _dafny.SeqOf(_dafny.SetOf((_this).F28()))
			var _741_v39 _dafny.Map
			_ = _741_v39
			_741_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _740_v38)
			var _742_v40 _dafny.Array
			_ = _742_v40
			_742_v40 = (_this).F27()
			var _pat_let_tv20 = _741_v39
			_ = _pat_let_tv20
			var _pat_let_tv21 = _742_v40
			_ = _pat_let_tv21
			var _pat_let_tv22 = _741_v39
			_ = _pat_let_tv22
			var _pat_let_tv23 = _742_v40
			_ = _pat_let_tv23
			var _pat_let_tv24 = _740_v38
			_ = _pat_let_tv24
			var _source10 D5 = func(_pat_let13_0 D5) D5 {
				return func(_743_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let14_0 _dafny.Sequence) D5 {
						return func(_744_dt__update_hcf20_h0 _dafny.Sequence) D5 {
							return Companion_D5_.Create_DC12_(_744_dt__update_hcf20_h0)
						}(_pat_let14_0)
					}((func() _dafny.Sequence {
						if (_pat_let_tv20).Contains(_pat_let_tv21) {
							return (_pat_let_tv22).Get(_pat_let_tv23).(_dafny.Sequence)
						}
						return _pat_let_tv24
					})())
				}(_pat_let13_0)
			}(Companion_D5_.Create_DC12_(_740_v38))
			_ = _source10
			if _source10.Is_DC13() {
				var _745___mcc_h0 bool = _source10.Get_().(D5_DC13).Cf21
				_ = _745___mcc_h0
				var _746___mcc_h1 _dafny.Array = _source10.Get_().(D5_DC13).Cf22
				_ = _746___mcc_h1
				var _747___mcc_h2 bool = _source10.Get_().(D5_DC13).Cf23
				_ = _747___mcc_h2
				var _748_cf23 bool = _747___mcc_h2
				_ = _748_cf23
				var _749_cf22 _dafny.Array = _746___mcc_h1
				_ = _749_cf22
				var _750_cf21 bool = _745___mcc_h0
				_ = _750_cf21
				var _751_v41 _dafny.Sequence
				_ = _751_v41
				_751_v41 = _dafny.SeqOf(p0)
				var _752_v42 _dafny.Sequence
				_ = _752_v42
				_752_v42 = _dafny.SeqOf(_751_v41)
				var _rhs135 _dafny.Sequence = _752_v42
				_ = _rhs135
				var _rhs136 _dafny.Int = p2
				_ = _rhs136
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				_752_v42 = _rhs135
				_lhs115.F21 = _rhs136
				var _753_v43 _dafny.MultiSet
				_ = _753_v43
				_753_v43 = _dafny.MultiSetOf(_this.F38)
				var _754_v44 _dafny.Sequence
				_ = _754_v44
				_754_v44 = _dafny.SeqOf(p2)
				var _rhs137 _dafny.Int = p2
				_ = _rhs137
				var _rhs138 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_754_v44).Cardinality()), p2)
				_ = _rhs138
				var _rhs139 _dafny.MultiSet = _753_v43
				_ = _rhs139
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				var _lhs117 *GlobalState = globalState
				_ = _lhs117
				_lhs116.F13 = _rhs137
				_lhs117.F5 = _rhs138
				_753_v43 = _rhs139
				var _755_v45 _dafny.Sequence
				_ = _755_v45
				_755_v45 = _dafny.SeqOf((_this).F27(), (_this).F27(), (_this).F27())
				var _756_v46 _dafny.Sequence
				_ = _756_v46
				_756_v46 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_755_v45, _755_v45), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_755_v45, _755_v45)).Cardinality()))).Uint32(), (_this).F27()), _dafny.Companion_Sequence_.Concatenate(_755_v45, _755_v45), _755_v45, _dafny.Companion_Sequence_.Concatenate(_755_v45, _755_v45))
				_755_v45 = (_756_v46).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_756_v46).Cardinality()))).Uint32()).(_dafny.Sequence)
				(globalState).F6 = _748_cf23
			} else {
				var _757___mcc_h3 _dafny.Sequence = _source10.Get_().(D5_DC12).Cf20
				_ = _757___mcc_h3
				var _758_cf20 _dafny.Sequence = _757___mcc_h3
				_ = _758_cf20
				var _759_v47 _dafny.Array
				_ = _759_v47
				var _nwElement0_23 _dafny.Array = _739_v37
				_ = _nwElement0_23
				var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(27))
				_ = _nw111
				(_nw111).ArraySet1(_nwElement0_23, 0)
				(_nw111).ArraySet1(_739_v37, 1)
				(_nw111).ArraySet1(_739_v37, 2)
				(_nw111).ArraySet1(_739_v37, 3)
				(_nw111).ArraySet1(_739_v37, 4)
				(_nw111).ArraySet1(_739_v37, 5)
				(_nw111).ArraySet1(_739_v37, 6)
				(_nw111).ArraySet1(_739_v37, 7)
				(_nw111).ArraySet1((func() _dafny.Array {
					if p0 {
						return _739_v37
					}
					return _739_v37
				})(), 8)
				(_nw111).ArraySet1(_739_v37, 9)
				(_nw111).ArraySet1(_739_v37, 10)
				(_nw111).ArraySet1(_739_v37, 11)
				(_nw111).ArraySet1(_739_v37, 12)
				(_nw111).ArraySet1(_739_v37, 13)
				(_nw111).ArraySet1(_739_v37, 14)
				(_nw111).ArraySet1(_739_v37, 15)
				(_nw111).ArraySet1(_739_v37, 16)
				(_nw111).ArraySet1(_739_v37, 17)
				(_nw111).ArraySet1(_739_v37, 18)
				(_nw111).ArraySet1(_739_v37, 19)
				(_nw111).ArraySet1(_739_v37, 20)
				(_nw111).ArraySet1(_739_v37, 21)
				(_nw111).ArraySet1(_739_v37, 22)
				(_nw111).ArraySet1(_739_v37, 23)
				(_nw111).ArraySet1(_739_v37, 24)
				(_nw111).ArraySet1(_739_v37, 25)
				(_nw111).ArraySet1(_739_v37, 26)
				_759_v47 = _nw111
				var _760_v48 _dafny.Sequence
				_ = _760_v48
				_760_v48 = _dafny.SeqOf(_739_v37)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_759_v47), 0))
				_ = _index108
				(_759_v47).ArraySet1((_760_v48).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_760_v48).Cardinality()))).Uint32()).(_dafny.Array), (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_759_v47), 0))
				_ = _index109
				(_759_v47).ArraySet1(_739_v37, (_index109).Int())
				var _761_v49 *C1
				_ = _761_v49
				var _nw112 *C1 = New_C1_()
				_ = _nw112
				_nw112.Ctor__(p0, (_this).F27(), p1)
				_761_v49 = _nw112
				var _762_v50 _dafny.MultiSet
				_ = _762_v50
				_762_v50 = _dafny.MultiSetOf(p1)
				var _763_v51 _dafny.Set
				_ = _763_v51
				_763_v51 = _dafny.SetOf((_761_v49).F39())
				var _764_v52 *C0
				_ = _764_v52
				var _nw113 *C0 = New_C0_()
				_ = _nw113
				_nw113.Ctor__((_761_v49).Fm34((_761_v49).F39(), p2, _762_v50, p2, globalState), (p2).Times(p2), (_this).F27(), (_763_v51).Contains(true))
				_764_v52 = _nw113
				var _765_v53 _dafny.Array
				_ = _765_v53
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
				_ = _nw114
				_765_v53 = _nw114
				_765_v53 = _765_v53
			}
		} else {
			(globalState).F21 = (p2).Minus((func() _dafny.Int {
				if p1 {
					return p2
				}
				return p2
			})())
			var _766_v54 *C1
			_ = _766_v54
			var _nw115 *C1 = New_C1_()
			_ = _nw115
			_nw115.Ctor__(!(_this.F38), (_this).F27(), (_this).F28())
			_766_v54 = _nw115
			if (_766_v54).F39() {
				var _767_v55 _dafny.Set
				_ = _767_v55
				_767_v55 = _dafny.SetOf((_766_v54).F39())
				var _rhs140 bool = (_dafny.MultiSetOf((_this).F28(), (_this).F28())).IsProperSubsetOf(_dafny.MultiSetOf(p0, (_this).F28()))
				_ = _rhs140
				var _rhs141 _dafny.Int = p2
				_ = _rhs141
				var _rhs142 _dafny.Int = Companion_Default___.Fm36(globalState)
				_ = _rhs142
				var _rhs143 _dafny.Int = p2
				_ = _rhs143
				var _rhs144 bool = !(!((_767_v55).IsSubsetOf(_767_v55)) || (p0))
				_ = _rhs144
				var _lhs118 *GlobalState = globalState
				_ = _lhs118
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				var _lhs121 *GlobalState = globalState
				_ = _lhs121
				var _lhs122 *GlobalState = globalState
				_ = _lhs122
				_lhs118.F6 = _rhs140
				_lhs119.F13 = _rhs141
				_lhs120.F13 = _rhs142
				_lhs121.F21 = _rhs143
				_lhs122.F15 = _rhs144
				var _768_v56 _dafny.Set
				_ = _768_v56
				_768_v56 = _dafny.SetOf(_dafny.IntOfInt64(840))
				var _769_v57 _dafny.Sequence
				_ = _769_v57
				_769_v57 = _dafny.SeqOf(p2, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vxsf")).Cardinality()), p2), (_768_v56).Cardinality(), p2)
				var _770_v58 _dafny.Sequence
				_ = _770_v58
				_770_v58 = _769_v57
				_769_v57 = (_770_v58)
				var _771_v59 D8
				_ = _771_v59
				_771_v59 = Companion_D8_.Create_DC18_(p2)
				(globalState).F25 = Companion_Default___.Fm42((_771_v59).Dtor_cf31(), globalState)
				(globalState).F15 = _this.F38
				var _772_v60 _dafny.Map
				_ = _772_v60
				_772_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_766_v54).F39()), (_this).F27())
				_772_v60 = (_772_v60).Update((_766_v54).F39(), (_this).F27())
			} else {
				(_this).F38 = true
				var _773_v61 _dafny.Sequence
				_ = _773_v61
				_773_v61 = _dafny.SeqOf(_dafny.IntOfInt64(863))
				var _774_v62 _dafny.Sequence
				_ = _774_v62
				_774_v62 = _773_v61
				var _775_v63 _dafny.Map
				_ = _775_v63
				_775_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v62, (_773_v61).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.IntOfUint32((_773_v61).Cardinality()))).Uint32()).(_dafny.Int))
				var _776_v64 _dafny.Map
				_ = _776_v64
				_776_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _775_v63)
				var _777_v65 _dafny.Sequence
				_ = _777_v65
				_777_v65 = _dafny.SeqOf((func() _dafny.Map {
					if (_776_v64).Contains(false) {
						return (_776_v64).Get(false).(_dafny.Map)
					}
					return _775_v63
				})())
				var _778_v66 D0
				_ = _778_v66
				_778_v66 = Companion_D0_.Create_DC3_(p2, (_this).F28(), p0)
				var _779_v68 _dafny.MultiSet
				_ = _779_v68
				_779_v68 = _dafny.MultiSetOf(_774_v62)
				var _780_v69 _dafny.Sequence
				_ = _780_v69
				_780_v69 = _dafny.SeqOf(p0, p0)
				var _781_v70 _dafny.Map
				_ = _781_v70
				_781_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_766_v54).F39(), (_778_v66).Dtor_cf7())).Update(true, p2))).Cardinality()), _775_v63)
				var _782_v72 _dafny.Array
				_ = _782_v72
				var _nwElement0_24 _dafny.Map = (_777_v65).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_777_v65).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _nwElement0_24
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(25))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_24, 0)
				(_nw116).ArraySet1(_775_v63, 1)
				(_nw116).ArraySet1(_775_v63, 2)
				(_nw116).ArraySet1(_775_v63, 3)
				(_nw116).ArraySet1((_775_v63).Merge((_775_v63).Update(_773_v61, p2)), 4)
				(_nw116).ArraySet1(Companion_Default___.Fm43(p2, _778_v66, p2, _this.F38, globalState), 5)
				(_nw116).ArraySet1(_775_v63, 6)
				(_nw116).ArraySet1((func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter49 := _dafny.Iterate((_779_v68).Elements()); ; {
						_compr_47, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _783_v67 _dafny.Sequence
						_783_v67 = interface{}(_compr_47).(_dafny.Sequence)
						if (_779_v68).Contains(_783_v67) {
							_coll47.Add(_783_v67, p2)
						}
					}
					return _coll47.ToMap()
				}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v62, p2)), 7)
				(_nw116).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v62, _dafny.IntOfInt64(592)), 8)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 9)
				(_nw116).ArraySet1(_775_v63, 10)
				(_nw116).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v62, p2)).Merge(_775_v63), 11)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 12)
				(_nw116).ArraySet1(_775_v63, 13)
				(_nw116).ArraySet1(_775_v63, 14)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 15)
				(_nw116).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((p3).Cardinality()), p2, _dafny.IntOfUint32((_773_v61).Cardinality())), _dafny.IntOfUint32((_780_v69).Cardinality()))).Merge(_775_v63), 16)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 17)
				(_nw116).ArraySet1((_775_v63).Update(_774_v62, p2), 18)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 19)
				(_nw116).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p2, _dafny.IntOfUint32((p3).Cardinality())), p2)).Merge(_775_v63), 20)
				(_nw116).ArraySet1(_775_v63, 21)
				(_nw116).ArraySet1((func() _dafny.Map {
					if (_781_v70).Contains(p2) {
						return (_781_v70).Get(p2).(_dafny.Map)
					}
					return func() _dafny.Map {
						var _coll48 = _dafny.NewMapBuilder()
						_ = _coll48
						for _iter50 := _dafny.Iterate((_779_v68).Elements()); ; {
							_compr_48, _ok50 := _iter50()
							if !_ok50 {
								break
							}
							var _784_v71 _dafny.Sequence
							_784_v71 = interface{}(_compr_48).(_dafny.Sequence)
							if (_779_v68).Contains(_784_v71) {
								_coll48.Add(_784_v71, p2)
							}
						}
						return _coll48.ToMap()
					}()
				})(), 22)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 23)
				(_nw116).ArraySet1((_775_v63).Merge(_775_v63), 24)
				_782_v72 = _nw116
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_782_v72), 0))
				_ = _index110
				(_782_v72).ArraySet1(_775_v63, (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_782_v72), 0))
				_ = _index111
				(_782_v72).ArraySet1((_775_v63).Merge(_775_v63), (_index111).Int())
				var _785_v73 _dafny.MultiSet
				_ = _785_v73
				_785_v73 = _dafny.MultiSetOf(p0, true)
				var _786_v74 _dafny.Map
				_ = _786_v74
				_786_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p2)
				var _787_v75 T1
				_ = _787_v75
				var _nw117 *C0 = New_C0_()
				_ = _nw117
				_nw117.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("at"), (_786_v74).Cardinality(), (_this).F27(), _this.F38)
				_787_v75 = _nw117
				var _788_v76 _dafny.Map
				_ = _788_v76
				_788_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_787_v75, p2)
				var _789_v77 _dafny.Map
				_ = _789_v77
				_789_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v76, (_766_v54).F39())
				var _790_v78 _dafny.Sequence
				_ = _790_v78
				_790_v78 = _dafny.SeqOf((_787_v75).F27())
				var _791_v79 _dafny.Map
				_ = _791_v79
				_791_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_787_v75.F29(), p2)
				var _792_v80 _dafny.Array
				_ = _792_v80
				var _nwElement0_25 _dafny.Int = ((_dafny.Zero).Minus(p2)).Plus(_dafny.IntOfInt64(922))
				_ = _nwElement0_25
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(19))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_25, 0)
				(_nw118).ArraySet1((func() _dafny.Int {
					if (_this).F28() {
						return p2
					}
					return p2
				})(), 1)
				(_nw118).ArraySet1(p2, 2)
				(_nw118).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_773_v61).Cardinality()), (_785_v73).Cardinality()), 3)
				(_nw118).ArraySet1(p2, 4)
				(_nw118).ArraySet1(p2, 5)
				(_nw118).ArraySet1((Companion_D0_.Create_DC0_(_789_v77, (_this).F28(), _dafny.IntOfUint32((_790_v78).Cardinality()), (_this).F37(), p1)).Dtor_cf2(), 6)
				(_nw118).ArraySet1(_787_v75.F29(), 7)
				(_nw118).ArraySet1((_787_v75.F29()).Plus((_791_v79).Cardinality()), 8)
				(_nw118).ArraySet1(p2, 9)
				(_nw118).ArraySet1(_787_v75.F29(), 10)
				(_nw118).ArraySet1(_787_v75.F29(), 11)
				(_nw118).ArraySet1(p2, 12)
				(_nw118).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), 13)
				(_nw118).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((p3).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((p3).Cardinality()))), 14)
				(_nw118).ArraySet1(p2, 15)
				(_nw118).ArraySet1((p2).Minus(_787_v75.F29()), 16)
				(_nw118).ArraySet1(((func() _dafny.Int {
					if (_791_v79).Contains((_dafny.Zero).Minus(_787_v75.F29())) {
						return (_791_v79).Get((_dafny.Zero).Minus(_787_v75.F29())).(_dafny.Int)
					}
					return _787_v75.F29()
				})()).Minus((_786_v74).Cardinality()), 17)
				(_nw118).ArraySet1(_dafny.IntOfUint32((_790_v78).Cardinality()), 18)
				_792_v80 = _nw118
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_792_v80), 0))
				_ = _index112
				(_792_v80).ArraySet1(p2, (_index112).Int())
				var _793_v81 _dafny.Array
				_ = _793_v81
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(16))
				_ = _nw119
				_793_v81 = _nw119
				var _794_v82 _dafny.Set
				_ = _794_v82
				_794_v82 = _dafny.SetOf(Companion_Default___.Fm40(globalState), p1, !(p1))
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_793_v81), 0))
				_ = _index113
				(_793_v81).ArraySet1(_794_v82, (_index113).Int())
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_792_v80), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_793_v81), 0))
				_ = _index115
				var _rhs145 bool = (_766_v54).F39()
				_ = _rhs145
				var _rhs146 bool = p0
				_ = _rhs146
				var _rhs147 _dafny.Int = p2
				_ = _rhs147
				var _rhs148 _dafny.Set = _dafny.SetOf(p1)
				_ = _rhs148
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				var _lhs124 *GlobalState = globalState
				_ = _lhs124
				var _lhs125 _dafny.Array = _792_v80
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_792_v80), 0))
				_ = _lhs126
				var _lhs127 _dafny.Array = _793_v81
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_793_v81), 0))
				_ = _lhs128
				_lhs123.F17 = _rhs145
				_lhs124.F6 = _rhs146
				(_lhs125).ArraySet1(_rhs147, (_lhs126).Int())
				(_lhs127).ArraySet1(_rhs148, (_lhs128).Int())
				var _795_v83 D3
				_ = _795_v83
				_795_v83 = Companion_D3_.Create_DC10_(true, _this.F38)
				_795_v83 = _795_v83
				(globalState).F15 = !(p1)
			}
			var _796_v84 _dafny.Sequence
			_ = _796_v84
			var _797_v85 _dafny.Sequence
			_ = _797_v85
			var _798_v86 _dafny.Sequence
			_ = _798_v86
			var _799_v87 _dafny.Map
			_ = _799_v87
			var _out13 _dafny.Sequence
			_ = _out13
			var _out14 _dafny.Sequence
			_ = _out14
			var _out15 _dafny.Sequence
			_ = _out15
			var _out16 _dafny.Map
			_ = _out16
			_out13, _out14, _out15, _out16 = (_766_v54).M14(_this.F38, globalState)
			_796_v84 = _out13
			_797_v85 = _out14
			_798_v86 = _out15
			_799_v87 = _out16
			var _800_v88 _dafny.Map
			_ = _800_v88
			_800_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
			_800_v88 = (_800_v88).Update(p2, (_766_v54).F39())
		}
		if !(!(!(p1))) {
			if (_this).F28() {
				var _801_v89 _dafny.Array
				_ = _801_v89
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_17
				var _nw120 _dafny.Array
				_ = _nw120
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw120 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) bool = func(_802_i4 _dafny.Int) bool {
						return _this.F38
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw120 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw120).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw120).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_801_v89 = _nw120
				_801_v89 = _801_v89
				var _803_v90 *C0
				_ = _803_v90
				var _nw121 *C0 = New_C0_()
				_ = _nw121
				_nw121.Ctor__(p3, p2, _801_v89, _this.F38)
				_803_v90 = _nw121
				var _804_v91 _dafny.Array
				_ = _804_v91
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw122
				_804_v91 = _nw122
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_804_v91), 0))
				_ = _index116
				(_804_v91).ArraySet1(p2, (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_804_v91), 0))
				_ = _index117
				(_804_v91).ArraySet1((p2).Times(p2), (_index117).Int())
				(globalState).F4 = (_804_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_804_v91), 0))).Int()).(_dafny.Int)
				var _805_v92 _dafny.Map
				_ = _805_v92
				var _806_v93 _dafny.Int
				_ = _806_v93
				var _out17 _dafny.Map
				_ = _out17
				var _out18 _dafny.Int
				_ = _out18
				_out17, _out18 = Companion_Default___.M0(globalState)
				_805_v92 = _out17
				_806_v93 = _out18
			} else {
				var _807_v94 _dafny.Set
				_ = _807_v94
				_807_v94 = _dafny.SetOf(p1, (_this).F28())
				var _808_v95 *C0
				_ = _808_v95
				var _nw123 *C0 = New_C0_()
				_ = _nw123
				_nw123.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(257))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}(func(_809_i5 _dafny.Int) _dafny.CodePoint {
					return (_this).F37()
				})), p2, (_this).F27(), (_dafny.IntOfUint32((p3).Cardinality())).Cmp((_807_v94).Cardinality()) < 0)
				_808_v95 = _nw123
				_808_v95 = _808_v95
				var _810_v96 _dafny.Sequence
				_ = _810_v96
				_810_v96 = _dafny.UnicodeSeqOfUtf8Bytes("cplehhjuc")
				var _811_v97 _dafny.Array
				_ = _811_v97
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw124
				_811_v97 = _nw124
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_811_v97), 0))
				_ = _index118
				(_811_v97).ArraySet1(_dafny.IntOfInt64(855), (_index118).Int())
				var _812_v98 _dafny.Map
				_ = _812_v98
				_812_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_807_v94).IsSubsetOf(_807_v94), (_808_v95).Fm15(globalState))
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_811_v97), 0))
				_ = _index119
				var _rhs149 _dafny.Int = p2
				_ = _rhs149
				var _rhs150 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p3, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}(func(_813_i6 _dafny.Int) _dafny.CodePoint {
					return (_this).F37()
				})), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}(func(_814_i6 _dafny.Int) _dafny.CodePoint {
					return (_this).F37()
				}))).Cardinality()))).Uint32(), (_this).F37()))
				_ = _rhs150
				var _rhs151 _dafny.Int = (func() _dafny.Int {
					if (_812_v98).Contains(p1) {
						return (_812_v98).Get(p1).(_dafny.Int)
					}
					return p2
				})()
				_ = _rhs151
				var _rhs152 _dafny.Int = p2
				_ = _rhs152
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				var _lhs131 _dafny.Array = _811_v97
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_811_v97), 0))
				_ = _lhs132
				_lhs129.F4 = _rhs149
				_810_v96 = _rhs150
				_lhs130.F13 = _rhs151
				(_lhs131).ArraySet1(_rhs152, (_lhs132).Int())
				var _815_v99 D9
				_ = _815_v99
				_815_v99 = Companion_D9_.Create_DC23_(_this.F38)
				var _816_v100 _dafny.MultiSet
				_ = _816_v100
				_816_v100 = _dafny.MultiSetOf(Companion_D9_.Create_DC23_(p0), _815_v99, _815_v99)
				var _817_v101 _dafny.Sequence
				_ = _817_v101
				_817_v101 = _dafny.SeqOf(_this.F38)
				var _818_v102 _dafny.Map
				_ = _818_v102
				_818_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, (_this).F28())
				var _819_v103 *C0
				_ = _819_v103
				var _nw125 *C0 = New_C0_()
				_ = _nw125
				_nw125.Ctor__((_808_v95).F36(), (p2).Minus(((_dafny.MultiSetOf(_816_v100, _816_v100, _dafny.MultiSetOf(_815_v99, Companion_D9_.Create_DC23_((_817_v101).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_817_v101).Cardinality()))).Uint32()).(bool))))).Update(_816_v100, Companion_Default___.Abs((_811_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_811_v97), 0))).Int()).(_dafny.Int)))).Cardinality()), (_this).F27(), ((func() bool {
					if (_818_v102).Contains(_this.F38) {
						return (_818_v102).Get(_this.F38).(bool)
					}
					return (_this).F28()
				})()) == ((_this).F28()))
				_819_v103 = _nw125
				(globalState).F21 = p2
				var _820_v104 _dafny.Sequence
				_ = _820_v104
				_820_v104 = _dafny.SeqOf((_this).F27(), (_this).F27())
				(globalState).F17 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_820_v104, _820_v104), (_this).F27())
			}
			var _821_v105 D8
			_ = _821_v105
			_821_v105 = Companion_D8_.Create_DC20_(Companion_D8_.Create_DC19_())
			var _822_v106 _dafny.Map
			_ = _822_v106
			_822_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((p3).Cardinality())), p2)
			var _823_v107 D8
			_ = _823_v107
			_823_v107 = Companion_D8_.Create_DC17_(_822_v106)
			var _pat_let_tv25 = _823_v107
			_ = _pat_let_tv25
			var _824_v108 _dafny.Map
			_ = _824_v108
			_824_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let15_0 D8) D8 {
				return func(_825_dt__update__tmp_h1 D8) D8 {
					return func(_pat_let16_0 D8) D8 {
						return func(_826_dt__update_hcf32_h0 D8) D8 {
							return Companion_D8_.Create_DC20_(_826_dt__update_hcf32_h0)
						}(_pat_let16_0)
					}(_pat_let_tv25)
				}(_pat_let15_0)
			}(_821_v105), _dafny.Companion_Sequence_.Concatenate(p3, _dafny.UnicodeSeqOfUtf8Bytes("dxkphh")))
			(globalState).F4 = (_824_v108).Cardinality()
			if p1 {
				var _827_v109 _dafny.Map
				_ = _827_v109
				var _828_v110 _dafny.Int
				_ = _828_v110
				var _out19 _dafny.Map
				_ = _out19
				var _out20 _dafny.Int
				_ = _out20
				_out19, _out20 = Companion_Default___.M0(globalState)
				_827_v109 = _out19
				_828_v110 = _out20
				var _829_v111 _dafny.Sequence
				_ = _829_v111
				_829_v111 = _dafny.SeqOf(_828_v110)
				var _830_v112 _dafny.Array
				_ = _830_v112
				var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw126
				_830_v112 = _nw126
				var _831_v113 _dafny.Sequence
				_ = _831_v113
				_831_v113 = _dafny.SeqOf(_830_v112, _830_v112)
				var _832_v114 _dafny.MultiSet
				_ = _832_v114
				_832_v114 = _dafny.MultiSetOf((_dafny.Zero).Minus((_829_v111).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_829_v111).Cardinality()))).Uint32()).(_dafny.Int)), (_dafny.Zero).Minus(_828_v110), _dafny.IntOfInt64(-504), (p2).Times(p2), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F28() {
						return _831_v113
					}
					return _831_v113
				})()).Cardinality()))
				(globalState).F21 = (func() _dafny.Int {
					if (_832_v114).Contains(_828_v110) {
						return (_832_v114).Multiplicity(_828_v110)
					}
					return _828_v110
				})()
				(globalState).F15 = p0
				var _833_v115 _dafny.Map
				_ = _833_v115
				_833_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p3).Cardinality()), (_this).F28())
				(globalState).F21 = ((_833_v115).Cardinality()).Times(Companion_Default___.Fm36(globalState))
				var _834_v116 _dafny.Array
				_ = _834_v116
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_18
				var _nw127 _dafny.Array
				_ = _nw127
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw127 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.CodePoint = func(_835_i7 _dafny.Int) _dafny.CodePoint {
						return (_this).F37()
					}
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw127 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw127).ArraySet1CodePoint(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw127).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_834_v116 = _nw127
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_834_v116), 0))
				_ = _index120
				(_834_v116).ArraySet1CodePoint((_this).F37(), (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_834_v116), 0))
				_ = _index121
				(_834_v116).ArraySet1CodePoint((_this).F37(), (_index121).Int())
			} else {
				var _836_v117 _dafny.Sequence
				_ = _836_v117
				_836_v117 = _dafny.SeqOf(false)
				var _837_v118 _dafny.Map
				_ = _837_v118
				_837_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_836_v117, (_this).F28())
				var _838_v119 _dafny.Map
				_ = _838_v119
				_838_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, (_837_v118).Merge(_837_v118))
				_838_v119 = (_838_v119).Update((_this).F28(), _837_v118)
				(globalState).F4 = _dafny.IntOfInt64(884)
				(globalState).F15 = _this.F38
				var _839_v120 _dafny.Array
				_ = _839_v120
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw128
				_839_v120 = _nw128
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_839_v120), 0))
				_ = _index122
				(_839_v120).ArraySet1((Companion_Default___.Fm36(globalState)).Minus(p2), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_839_v120), 0))
				_ = _index123
				(_839_v120).ArraySet1((_dafny.IntOfUint32((_836_v117).Cardinality())).Plus(p2), (_index123).Int())
				(globalState).F13 = (_839_v120).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_839_v120), 0))).Int()).(_dafny.Int)
			}
			(globalState).F21 = (p2).Plus(p2)
			var _840_v121 _dafny.Array
			_ = _840_v121
			var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw129
			_840_v121 = _nw129
			var _841_v122 _dafny.Sequence
			_ = _841_v122
			_841_v122 = _dafny.SeqOf(p2)
			var _842_v123 _dafny.MultiSet
			_ = _842_v123
			_842_v123 = _dafny.MultiSetOf(_841_v122)
			var _843_v124 _dafny.Sequence
			_ = _843_v124
			_843_v124 = _dafny.SeqOf(p0, p0, p1)
			var _844_v125 _dafny.Set
			_ = _844_v125
			_844_v125 = _dafny.SetOf((_843_v124).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_843_v124).Cardinality()))).Uint32()).(bool))
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_840_v121), 0))
			_ = _index124
			(_840_v121).ArraySet1(((_842_v123).Cardinality()).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_844_v125).Cardinality())).Cardinality()), (_index124).Int())
			var _845_v126 _dafny.Map
			_ = _845_v126
			_845_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F37(), _dafny.IntOfUint32((p3).Cardinality()))
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_840_v121), 0))
			_ = _index125
			(_840_v121).ArraySet1((func() _dafny.Int {
				if (_845_v126).Contains((_this).F37()) {
					return (_845_v126).Get((_this).F37()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(p2)
			})(), (_index125).Int())
		} else {
			var _846_v127 _dafny.Array
			_ = _846_v127
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
			_ = _nw130
			_846_v127 = _nw130
			var _847_v128 D0
			_ = _847_v128
			_847_v128 = Companion_D0_.Create_DC3_(p2, p1, p0)
			var _pat_let_tv26 = p1
			_ = _pat_let_tv26
			var _848_v129 _dafny.Array
			_ = _848_v129
			var _nwElement0_26 D0 = _847_v128
			_ = _nwElement0_26
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(26))
			_ = _nw131
			(_nw131).ArraySet1(_nwElement0_26, 0)
			(_nw131).ArraySet1(_847_v128, 1)
			(_nw131).ArraySet1(_847_v128, 2)
			(_nw131).ArraySet1(_847_v128, 3)
			(_nw131).ArraySet1(_847_v128, 4)
			(_nw131).ArraySet1(func(_pat_let17_0 D0) D0 {
				return func(_849_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let18_0 bool) D0 {
						return func(_850_dt__update_hcf8_h0 bool) D0 {
							return Companion_D0_.Create_DC3_((_849_dt__update__tmp_h2).Dtor_cf7(), _850_dt__update_hcf8_h0, (_849_dt__update__tmp_h2).Dtor_cf9())
						}(_pat_let18_0)
					}(_pat_let_tv26)
				}(_pat_let17_0)
			}(_847_v128), 5)
			(_nw131).ArraySet1(_847_v128, 6)
			(_nw131).ArraySet1(Companion_D0_.Create_DC3_(p2, p1, p1), 7)
			(_nw131).ArraySet1(_847_v128, 8)
			(_nw131).ArraySet1(_847_v128, 9)
			(_nw131).ArraySet1(_847_v128, 10)
			(_nw131).ArraySet1(_847_v128, 11)
			(_nw131).ArraySet1(_847_v128, 12)
			(_nw131).ArraySet1(_847_v128, 13)
			(_nw131).ArraySet1(_847_v128, 14)
			(_nw131).ArraySet1(_847_v128, 15)
			(_nw131).ArraySet1(Companion_D0_.Create_DC3_(p2, p0, p1), 16)
			(_nw131).ArraySet1(_847_v128, 17)
			(_nw131).ArraySet1(Companion_D0_.Create_DC3_(p2, p0, p0), 18)
			(_nw131).ArraySet1(_847_v128, 19)
			(_nw131).ArraySet1(_847_v128, 20)
			(_nw131).ArraySet1(_847_v128, 21)
			(_nw131).ArraySet1(_847_v128, 22)
			(_nw131).ArraySet1(_847_v128, 23)
			(_nw131).ArraySet1(Companion_D0_.Create_DC3_(p2, (_this).F28(), p0), 24)
			(_nw131).ArraySet1(_847_v128, 25)
			_848_v129 = _nw131
			var _851_v130 _dafny.Map
			_ = _851_v130
			_851_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(p2)), p1), _848_v129)
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_846_v127), 0))
			_ = _index126
			(_846_v127).ArraySet1(_851_v130, (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_846_v127), 0))
			_ = _index127
			(_846_v127).ArraySet1(_851_v130, (_index127).Int())
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index128
			((_this).F27()).ArraySet1((_this.F38) && ((_this).F28()), (_index128).Int())
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index129
			((_this).F27()).ArraySet1((p2).Cmp((_dafny.Zero).Minus((_dafny.IntOfInt64(829)).Times(p2))) == 0, (_index129).Int())
			var _852_v131 _dafny.Sequence
			_ = _852_v131
			_852_v131 = _dafny.UnicodeSeqOfUtf8Bytes("cqp")
			var _853_v132 _dafny.Array
			_ = _853_v132
			var _nwElement0_27 bool = p1
			_ = _nwElement0_27
			var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(8))
			_ = _nw132
			(_nw132).ArraySet1(_nwElement0_27, 0)
			(_nw132).ArraySet1((_this).F28(), 1)
			(_nw132).ArraySet1((_this).F28(), 2)
			(_nw132).ArraySet1(p1, 3)
			(_nw132).ArraySet1(p1, 4)
			(_nw132).ArraySet1(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), 5)
			(_nw132).ArraySet1(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), 6)
			(_nw132).ArraySet1(p0, 7)
			_853_v132 = _nw132
			var _854_v133 T1
			_ = _854_v133
			var _nw133 *C0 = New_C0_()
			_ = _nw133
			_nw133.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("p"), p2, _853_v132, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool))
			_854_v133 = _nw133
			var _855_v134 _dafny.Map
			_ = _855_v134
			_855_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _854_v133)
			var _rhs153 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return p2
				}
				return (func() _dafny.Int {
					if true {
						return _dafny.IntOfInt64(-641)
					}
					return (_855_v134).Cardinality()
				})()
			})()
			_ = _rhs153
			var _rhs154 _dafny.Int = _dafny.IntOfInt64(187)
			_ = _rhs154
			var _rhs155 _dafny.Sequence = p3
			_ = _rhs155
			var _rhs156 _dafny.Int = _dafny.IntOfInt64(997)
			_ = _rhs156
			var _lhs133 *GlobalState = globalState
			_ = _lhs133
			var _lhs134 *GlobalState = globalState
			_ = _lhs134
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			_lhs133.F5 = _rhs153
			_lhs134.F13 = _rhs154
			_852_v131 = _rhs155
			_lhs135.F21 = _rhs156
			var _856_v135 _dafny.Array
			_ = _856_v135
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_19
			var _nw134 _dafny.Array
			_ = _nw134
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw134 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = (func(_857_v133 T1) func(_dafny.Int) _dafny.Int {
					return func(_858_i8 _dafny.Int) _dafny.Int {
						return (_858_i8).Minus((_dafny.Zero).Minus(_857_v133.F29()))
					}
				})(_854_v133)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw134 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw134).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw134).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_856_v135 = _nw134
			_856_v135 = _856_v135
			var _859_v136 D8
			_ = _859_v136
			_859_v136 = Companion_D8_.Create_DC18_(_854_v133.F29())
			var _source11 D8 = _859_v136
			_ = _source11
			if _source11.Is_DC18() {
				var _860___mcc_h4 _dafny.Int = _source11.Get_().(D8_DC18).Cf31
				_ = _860___mcc_h4
				var _861_cf31 _dafny.Int = _860___mcc_h4
				_ = _861_cf31
				(globalState).F6 = (_847_v128).Dtor_cf8()
				var _862_v137 _dafny.MultiSet
				_ = _862_v137
				_862_v137 = _dafny.MultiSetOf((_854_v133).F28())
				var _863_v138 _dafny.Map
				_ = _863_v138
				_863_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_862_v137).Cardinality())
				var _864_v139 _dafny.Set
				_ = _864_v139
				_864_v139 = _dafny.SetOf(((_854_v133).F27()), _853_v132)
				_863_v138 = (_863_v138).Update(_854_v133.F29(), (_864_v139).Cardinality())
				(globalState).F21 = Companion_Default___.Fm36(globalState)
				var _865_v140 _dafny.Sequence
				_ = _865_v140
				_865_v140 = _dafny.SeqOf((_854_v133).F28(), !((_854_v133).F28()), ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), true)
				var _866_v141 _dafny.Map
				_ = _866_v141
				_866_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.Companion_Sequence_.Update(_865_v140, (Companion_Default___.SafeIndex(_854_v133.F29(), _dafny.IntOfUint32((_865_v140).Cardinality()))).Uint32(), p1))
				var _867_v142 _dafny.MultiSet
				_ = _867_v142
				_867_v142 = _dafny.MultiSetOf((_854_v133).F27())
				var _rhs157 bool = (_867_v142).IsSubsetOf((_867_v142).Intersection(_867_v142))
				_ = _rhs157
				var _rhs158 _dafny.Map = _866_v141
				_ = _rhs158
				var _rhs159 _dafny.Int = (((_dafny.Zero).Minus(_854_v133.F29())).Minus(_854_v133.F29())).Times(_861_cf31)
				_ = _rhs159
				var _lhs136 *GlobalState = globalState
				_ = _lhs136
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				_lhs136.F17 = _rhs157
				_866_v141 = _rhs158
				_lhs137.F21 = _rhs159
			} else if _source11.Is_DC19() {
				(globalState).F13 = p2
				(_854_v133).F29_set_(p2)
				var _868_v143 *C1
				_ = _868_v143
				var _nw135 *C1 = New_C1_()
				_ = _nw135
				_nw135.Ctor__((_854_v133).F28(), (_854_v133).F27(), p0)
				_868_v143 = _nw135
				(globalState).F4 = _854_v133.F29()
			} else if _source11.Is_DC17() {
				var _869___mcc_h5 _dafny.Map = _source11.Get_().(D8_DC17).Cf30
				_ = _869___mcc_h5
				var _870_cf30 _dafny.Map = _869___mcc_h5
				_ = _870_cf30
				var _871_v144 _dafny.Array
				_ = _871_v144
				var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw136
				_871_v144 = _nw136
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v144), 0))
				_ = _index130
				(_871_v144).ArraySet1(_856_v135, (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v144), 0))
				_ = _index131
				(_871_v144).ArraySet1(_856_v135, (_index131).Int())
				var _872_v145 _dafny.Map
				_ = _872_v145
				_872_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.ArrayCastTo((_871_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v144), 0))).Int())))
				var _873_v146 _dafny.Map
				_ = _873_v146
				_873_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_v133.F29(), (func() _dafny.Array {
					if (_872_v145).Contains(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool)) {
						return (_872_v145).Get(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool)).(_dafny.Array)
					}
					return _856_v135
				})())
				_873_v146 = (_873_v146).Update((_this).Fm20(globalState), _dafny.ArrayCastTo((_871_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v144), 0))).Int())))
				var _874_v147 _dafny.Sequence
				_ = _874_v147
				_874_v147 = _dafny.SeqOf((_854_v133).F28())
				var _875_v148 _dafny.Set
				_ = _875_v148
				_875_v148 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_874_v147, _874_v147)).Cardinality()), _854_v133.F29())
				var _rhs160 _dafny.Set = _875_v148
				_ = _rhs160
				var _rhs161 _dafny.Int = (_dafny.IntOfInt64(690)).Minus(p2)
				_ = _rhs161
				var _rhs162 _dafny.CodePoint = (_852_v131).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_852_v131).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs162
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				_875_v148 = _rhs160
				_lhs138.F5 = _rhs161
				_lhs139.F25 = _rhs162
				(globalState).F13 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_v133, _dafny.IntOfUint32((p3).Cardinality()))).Cardinality()).Minus(p2)
			} else {
				var _876___mcc_h6 D8 = _source11.Get_().(D8_DC20).Cf32
				_ = _876___mcc_h6
				var _877_cf32 D8 = _876___mcc_h6
				_ = _877_cf32
				var _878_v149 T0
				_ = _878_v149
				var _nw137 *C1 = New_C1_()
				_ = _nw137
				_nw137.Ctor__(true, (_854_v133).F27(), p1)
				_878_v149 = _nw137
				var _879_v150 _dafny.Map
				_ = _879_v150
				_879_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _878_v149)
				var _880_v151 _dafny.Map
				_ = _880_v151
				_880_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p2, _854_v133.F29()), (func() T0 {
					if (_879_v150).Contains(p1) {
						return (_879_v150).Get(p1).(T0)
					}
					return _878_v149
				})())
				_880_v151 = (_880_v151).Update(_854_v133.F29(), _878_v149)
				var _881_v152 _dafny.Array
				_ = _881_v152
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw138
				_881_v152 = _nw138
				var _882_v153 _dafny.Set
				_ = _882_v153
				_882_v153 = _dafny.SetOf(_854_v133.F29())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_881_v152), 0))
				_ = _index132
				(_881_v152).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v153, _854_v133.F29()), (_index132).Int())
				var _883_v154 _dafny.Map
				_ = _883_v154
				_883_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F28())
				var _884_v155 _dafny.Map
				_ = _884_v155
				_884_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_883_v154).Cardinality(), _854_v133.F29())
				var _885_v156 _dafny.Map
				_ = _885_v156
				_885_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _854_v133.F29())
				var _886_v157 _dafny.Map
				_ = _886_v157
				_886_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(826), p2, (_884_v155).Cardinality(), p2, (_dafny.Zero).Minus(p2))).Union(_882_v153), Companion_Default___.SafeModuloInt((_885_v156).Cardinality(), p2))
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_881_v152), 0))
				_ = _index133
				(_881_v152).ArraySet1(_886_v157, (_index133).Int())
				var _887_v158 _dafny.Array
				_ = _887_v158
				var _nwElement0_28 _dafny.Array = (_878_v149).F27()
				_ = _nwElement0_28
				var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(11))
				_ = _nw139
				(_nw139).ArraySet1(_nwElement0_28, 0)
				(_nw139).ArraySet1((_854_v133).F27(), 1)
				(_nw139).ArraySet1((_854_v133).F27(), 2)
				(_nw139).ArraySet1((_854_v133).F27(), 3)
				(_nw139).ArraySet1((_854_v133).F27(), 4)
				(_nw139).ArraySet1((_854_v133).F27(), 5)
				(_nw139).ArraySet1((_878_v149).F27(), 6)
				(_nw139).ArraySet1((_878_v149).F27(), 7)
				(_nw139).ArraySet1((_854_v133).F27(), 8)
				(_nw139).ArraySet1((_878_v149).F27(), 9)
				(_nw139).ArraySet1((_854_v133).F27(), 10)
				_887_v158 = _nw139
				var _888_v159 D9
				_ = _888_v159
				_888_v159 = Companion_D9_.Create_DC22_(_856_v135, false, _887_v158, _852_v131)
				_885_v156 = (_885_v156).Update((_888_v159).Dtor_cf35(), (_854_v133.F29()).Times(_854_v133.F29()))
				var _889_v160 _dafny.Map
				_ = _889_v160
				_889_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_878_v149).F28(), p3)
				var _890_v161 _dafny.Sequence
				_ = _890_v161
				_890_v161 = _dafny.SeqOf(_852_v131, (func() _dafny.Sequence {
					if (_889_v160).Contains((_854_v133).F28()) {
						return (_889_v160).Get((_854_v133).F28()).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-306))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}(func(_891_i9 _dafny.Int) _dafny.CodePoint {
						return (_this).F37()
					}))
				})())
				var _892_v162 _dafny.Sequence
				_ = _892_v162
				_892_v162 = _dafny.SeqOf(p2)
				var _893_v163 D1
				_ = _893_v163
				_893_v163 = Companion_D1_.Create_DC7_((_this).F37(), _852_v131)
				var _894_v164 _dafny.Array
				_ = _894_v164
				var _nwElement0_29 _dafny.Sequence = _852_v131
				_ = _nwElement0_29
				var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(25))
				_ = _nw140
				(_nw140).ArraySet1(_nwElement0_29, 0)
				(_nw140).ArraySet1(Companion_Default___.Fm32((_dafny.Zero).Minus(_854_v133.F29()), (_854_v133).F28(), globalState), 1)
				(_nw140).ArraySet1((_890_v161).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_892_v162).Cardinality()), _dafny.IntOfUint32((_890_v161).Cardinality()))).Uint32()).(_dafny.Sequence), 2)
				(_nw140).ArraySet1(p3, 3)
				(_nw140).ArraySet1(_852_v131, 4)
				(_nw140).ArraySet1(p3, 5)
				(_nw140).ArraySet1(_852_v131, 6)
				(_nw140).ArraySet1(_852_v131, 7)
				(_nw140).ArraySet1((_893_v163).Dtor_cf14(), 8)
				(_nw140).ArraySet1(_852_v131, 9)
				(_nw140).ArraySet1(Companion_Default___.Fm32(_854_v133.F29(), p0, globalState), 10)
				(_nw140).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p3, p3), 11)
				(_nw140).ArraySet1(p3, 12)
				(_nw140).ArraySet1(p3, 13)
				(_nw140).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm32(_854_v133.F29(), (_this).F28(), globalState), _852_v131), 14)
				(_nw140).ArraySet1(p3, 15)
				(_nw140).ArraySet1(_852_v131, 16)
				(_nw140).ArraySet1(p3, 17)
				(_nw140).ArraySet1(_852_v131, 18)
				(_nw140).ArraySet1(p3, 19)
				(_nw140).ArraySet1((_890_v161).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.IntOfUint32((_890_v161).Cardinality()))).Uint32()).(_dafny.Sequence), 20)
				(_nw140).ArraySet1(_852_v131, 21)
				(_nw140).ArraySet1(p3, 22)
				(_nw140).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_852_v131, p3), 23)
				(_nw140).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_852_v131, _852_v131), 24)
				_894_v164 = _nw140
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_894_v164), 0))
				_ = _index134
				(_894_v164).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("obsvgeys"), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_894_v164), 0))
				_ = _index135
				(_894_v164).ArraySet1(_852_v131, (_index135).Int())
			}
		}
		var _895_v165 _dafny.Sequence
		_ = _895_v165
		_895_v165 = _dafny.SeqOf(_dafny.IntOfInt64(8), p2, p2)
		var _896_v166 _dafny.Sequence
		_ = _896_v166
		_896_v166 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(378))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg55 _dafny.Int) interface{} {
				return coer55(arg55)
			}
		}((func(_897_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_898_i10 _dafny.Int) _dafny.Int {
				return _897_p2
			}
		})(p2)))
		_895_v165 = _dafny.Companion_Sequence_.Concatenate(_895_v165, (func() _dafny.Sequence {
			if !((_this).F28()) {
				return _895_v165
			}
			return (_896_v166)
		})())
		(globalState).F6 = !(_this.F38) || ((p0) || ((_this).F28()))
		(globalState).F4 = p2
		r0 = Companion_Default___.Fm44(globalState)
		return r0
	}
}
func (_this *C2) F37() _dafny.CodePoint {
	{
		return _this._f37
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f29 _dafny.Int
	_f27 _dafny.Array
	_f28 bool
	_f34 _dafny.Int
	_f35 _dafny.Map
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f29 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
	_this._f34 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F29() _dafny.Int {
	return _this._f29
}
func (_this *C3) F29_set_(value _dafny.Int) {
	_this._f29 = value
}
func (_this *C3) F27() _dafny.Array {
	return _this._f27
}
func (_this *C3) F28() bool {
	return _this._f28
}
func (_this *C3) Ctor__(f34 _dafny.Int, f35 _dafny.Map, f29 _dafny.Int, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f29 = f29
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C3) Fm14(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return ((_this.F29()).Times((_this).F34())).Minus(_this.F29())
	}
}
func (_this *C3) M10(p0 bool, globalState *GlobalState) (bool, bool, _dafny.CodePoint, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _899_v0 _dafny.Sequence
		_ = _899_v0
		_899_v0 = _dafny.SeqOf((_this).F27())
		(globalState).F17 = !(_dafny.Companion_Sequence_.IsPrefixOf(_899_v0, _dafny.Companion_Sequence_.Concatenate(_899_v0, _899_v0)))
		var _900_v1 _dafny.Map
		_ = _900_v1
		_900_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
		var _901_v2 *C0
		_ = _901_v2
		var _nw141 *C0 = New_C0_()
		_ = _nw141
		_nw141.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("phmmcxi"), (_this).F34(), (_this).F27(), Companion_Default___.Fm16(_900_v1, globalState))
		_901_v2 = _nw141
		var _902_v3 _dafny.Sequence
		_ = _902_v3
		_902_v3 = _dafny.SeqOf((_this).F28(), (_this).F28(), p0)
		var _903_v4 _dafny.Sequence
		_ = _903_v4
		_903_v4 = _dafny.SeqOf((_902_v3).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_902_v3).Cardinality()))).Uint32()).(bool))
		var _904_v5 _dafny.Sequence
		_ = _904_v5
		_904_v5 = _dafny.SeqOf(_903_v4)
		var _905_v6 _dafny.Sequence
		_ = _905_v6
		_905_v6 = _dafny.SeqOf((_this).F34())
		var _906_v7 _dafny.CodePoint
		_ = _906_v7
		_906_v7 = _dafny.CodePoint('p')
		var _907_v8 T1
		_ = _907_v8
		var _nw142 *C0 = New_C0_()
		_ = _nw142
		_nw142.Ctor__((_901_v2).F36(), _this.F29(), (_this).F27(), p0)
		_907_v8 = _nw142
		var _908_v9 _dafny.Map
		_ = _908_v9
		_908_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_901_v2).F36(), (Companion_Default___.SafeIndex((_905_v6).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_905_v6).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32(((_901_v2).F36()).Cardinality()))).Uint32(), _906_v7), _907_v8)
		var _909_v10 _dafny.Set
		_ = _909_v10
		_909_v10 = _dafny.SetOf((func() T1 {
			if (_908_v9).Contains((_901_v2).F36()) {
				return (_908_v9).Get((_901_v2).F36()).(T1)
			}
			return _907_v8
		})())
		var _910_v11 _dafny.Map
		_ = _910_v11
		_910_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_904_v5, _909_v10)
		_910_v11 = (_910_v11).Update(_904_v5, _909_v10)
		var _911_v12 _dafny.Sequence
		_ = _911_v12
		_911_v12 = _905_v6
		var _912_v13 _dafny.Sequence
		_ = _912_v13
		_912_v13 = _dafny.SeqOf(_905_v6, _905_v6, (_911_v12), _905_v6)
		var _913_v14 D1
		_ = _913_v14
		_913_v14 = Companion_D1_.Create_DC5_(_912_v13)
		var _914_v15 _dafny.MultiSet
		_ = _914_v15
		_914_v15 = _dafny.MultiSetOf(Companion_D1_.Create_DC5_(Companion_Default___.Fm17(p0, (_907_v8).F28(), _dafny.IntOfUint32(((_901_v2).F36()).Cardinality()), p0, globalState)), Companion_D1_.Create_DC5_(_912_v13), _913_v14, Companion_D1_.Create_DC5_(_912_v13), _913_v14)
		var _915_v16 _dafny.MultiSet
		_ = _915_v16
		_915_v16 = _dafny.MultiSetOf(_914_v15)
		_915_v16 = ((_915_v16).Update(_914_v15, Companion_Default___.Abs((_dafny.Zero).Minus(_dafny.IntOfUint32(((_901_v2).F36()).Cardinality()))))).Intersection(_915_v16)
		var _916_v17 D0
		_ = _916_v17
		_916_v17 = Companion_D0_.Create_DC3_(_907_v8.F29(), !((_this).F28()), true)
		var _917_v18 _dafny.Sequence
		_ = _917_v18
		_917_v18 = _dafny.SeqOf(_916_v17)
		var _918_v19 _dafny.Array
		_ = _918_v19
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_20
		var _nw143 _dafny.Array
		_ = _nw143
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw143 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = (func(_919_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_920_i0 _dafny.Int) _dafny.Int {
					return (_920_i0).Plus(_dafny.IntOfUint32((_919_v4).Cardinality()))
				}
			})(_903_v4)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw143 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw143).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw143).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_918_v19 = _nw143
		var _921_v20 _dafny.Map
		_ = _921_v20
		_921_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_917_v18).Cardinality())).Times((_this).F34()), _918_v19)
		_921_v20 = _921_v20
		r0 = !((_this).F28())
		var _922_v21 _dafny.MultiSet
		_ = _922_v21
		_922_v21 = _dafny.MultiSetOf((_907_v8).F28(), (_this).F28())
		var _923_v22 _dafny.Sequence
		_ = _923_v22
		_923_v22 = _dafny.SeqOf(_922_v21)
		var _924_v23 _dafny.Sequence
		_ = _924_v23
		_924_v23 = _923_v22
		r0 = ((_this).Fm14(_dafny.IntOfUint32((_905_v6).Cardinality()), func() _dafny.Set {
			var _coll49 = _dafny.NewBuilder()
			_ = _coll49
			for _iter51 := _dafny.Iterate((_924_v23).Elements()); ; {
				_compr_49, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _925_v24 _dafny.MultiSet
				_925_v24 = interface{}(_compr_49).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains((_924_v23), _925_v24) {
					_coll49.Add(_925_v24)
				}
			}
			return _coll49.ToSet()
		}(), globalState)).Cmp((_905_v6).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_905_v6).Cardinality()))).Uint32()).(_dafny.Int)) <= 0
		var _926_v25 _dafny.Map
		_ = _926_v25
		_926_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F34())
		var _927_v26 _dafny.Map
		_ = _927_v26
		_927_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_926_v25).Contains(p0) {
				return (_926_v25).Get(p0).(_dafny.Int)
			}
			return _907_v8.F29()
		})(), (_this).F34())
		r1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_903_v4, _902_v3)).Cardinality())).Cmp((func() _dafny.Int {
			if (_927_v26).Contains(_dafny.IntOfInt64(159)) {
				return (_927_v26).Get(_dafny.IntOfInt64(159)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(241)
		})()) <= 0
		r2 = _906_v7
		r3 = (_this).F34()
		return r0, r1, r2, r3
	}
}
func (_this *C3) M11(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 bool = false
		_ = r2
		var _928_v0 _dafny.Map
		_ = _928_v0
		_928_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_this).F28())
		var _929_v1 _dafny.Map
		_ = _929_v1
		_929_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_928_v0).Update((_this).F34(), p0), (_this).F34())
		var _930_v2 _dafny.MultiSet
		_ = _930_v2
		_930_v2 = _dafny.MultiSetOf(true)
		var _931_v3 _dafny.Set
		_ = _931_v3
		_931_v3 = _dafny.SetOf(_930_v2)
		(globalState).F13 = (_this).Fm14((_this).Fm14((func() _dafny.Int {
			if (_929_v1).Contains(_928_v0) {
				return (_929_v1).Get(_928_v0).(_dafny.Int)
			}
			return _this.F29()
		})(), _931_v3, globalState), _931_v3, globalState)
		var _hi2 _dafny.Int = (_this).F34()
		_ = _hi2
		for _932_i0 := (_this).F34(); _932_i0.Cmp(_hi2) < 0; _932_i0 = _932_i0.Plus(_dafny.One) {
			var _933_v4 _dafny.Sequence
			_ = _933_v4
			_933_v4 = _dafny.SeqOf(_932_i0, (_this).F34(), _dafny.IntOfInt64(137), _dafny.IntOfInt64(-189), _932_i0)
			var _934_v5 _dafny.Sequence
			_ = _934_v5
			_934_v5 = _933_v4
			var _source12 _dafny.Sequence = _934_v5
			_ = _source12
			var _935___mcc_h0 _dafny.Sequence = _source12
			_ = _935___mcc_h0
			var _936_cf15 _dafny.Sequence = _935___mcc_h0
			_ = _936_cf15
			var _937_v6 _dafny.Array
			_ = _937_v6
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_21
			var _nw144 _dafny.Array
			_ = _nw144
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw144 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_938_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_939_i1 _dafny.Int) _dafny.Int {
						return (_939_i1).Minus(_dafny.IntOfUint32((_938_v4).Cardinality()))
					}
				})(_933_v4)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw144 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw144).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw144).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_937_v6 = _nw144
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_937_v6), 0))
			_ = _index136
			(_937_v6).ArraySet1((_this).F34(), (_index136).Int())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_937_v6), 0))
			_ = _index137
			(_937_v6).ArraySet1(_this.F29(), (_index137).Int())
			(globalState).F25 = _dafny.CodePoint('h')
			var _940_v7 D1
			_ = _940_v7
			_940_v7 = Companion_D1_.Create_DC6_(_932_i0)
			(globalState).F21 = (_dafny.Zero).Minus((_940_v7).Dtor_cf12())
			var _941_v8 _dafny.Map
			_ = _941_v8
			_941_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_940_v7, _932_i0)
			_941_v8 = (_941_v8).Update(_940_v7, _this.F29())
			(globalState).F6 = (_932_i0).Cmp(_this.F29()) > 0
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index138
			((_this).F27()).ArraySet1(p0, (_index138).Int())
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index139
			((_this).F27()).ArraySet1((_this).F28(), (_index139).Int())
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index140
			((_this).F27()).ArraySet1(false, (_index140).Int())
		}
		var _942_v9 _dafny.Array
		_ = _942_v9
		var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
		_ = _nw145
		_942_v9 = _nw145
		var _943_v10 _dafny.Map
		_ = _943_v10
		_943_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_942_v9, _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(false), (func() bool {
			if (_928_v0).Contains(_this.F29()) {
				return (_928_v0).Get(_this.F29()).(bool)
			}
			return (_this).F28()
		})()))
		_943_v10 = (_943_v10).Update(_942_v9, !((_this).F28()))
		var _944_v11 _dafny.Map
		_ = _944_v11
		_944_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
		var _945_i2 _dafny.Int
		_ = _945_i2
		_945_i2 = _dafny.Zero
		{
			for !((func() bool {
				if Companion_Default___.Fm16(_944_v11, globalState) {
					return p0
				}
				return (_this).F28()
			})()) {
				{
					if (_945_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_945_i2 = (_945_i2).Plus(_dafny.One)
					if !(p0) {
						var _946_v12 D3
						_ = _946_v12
						_946_v12 = Companion_D3_.Create_DC10_((_this).F28(), p0)
						_928_v0 = (_928_v0).Update((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_946_v12).Dtor_cf18(), _dafny.MultiSetOf((_this).F34()))).Cardinality()), p0)
						var _947_v13 _dafny.CodePoint
						_ = _947_v13
						_947_v13 = _dafny.CodePoint('e')
						var _948_v14 *C0
						_ = _948_v14
						var _nw146 *C0 = New_C0_()
						_ = _nw146
						_nw146.Ctor__(Companion_Default___.Fm18(p0, (_this).F34(), _947_v13, globalState), (_dafny.Zero).Minus((_this).F34()), (_this).F27(), !((_this).F28()) || (p0))
						_948_v14 = _nw146
						var _949_v15 _dafny.Sequence
						_ = _949_v15
						_949_v15 = _dafny.SeqOf(p0)
						var _950_v16 _dafny.Map
						_ = _950_v16
						_950_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_949_v15).Cardinality()))).Cardinality())
						var _951_v17 _dafny.Set
						_ = _951_v17
						_951_v17 = _dafny.SetOf(_950_v16, _950_v16, _950_v16, _950_v16)
						var _952_v18 _dafny.Map
						_ = _952_v18
						_952_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(704), _951_v17)
						_951_v17 = ((func() _dafny.Set {
							if (_952_v18).Contains(_this.F29()) {
								return (_952_v18).Get(_this.F29()).(_dafny.Set)
							}
							return _dafny.SetOf(_950_v16)
						})()).Union((_dafny.SetOf(_950_v16, _950_v16, _950_v16)).Union(_951_v17))
						var _953_v19 _dafny.Sequence
						_ = _953_v19
						_953_v19 = _dafny.SeqOf((_this).F27(), (_this).F27())
						var _954_v20 T0
						_ = _954_v20
						var _nw147 *C1 = New_C1_()
						_ = _nw147
						_nw147.Ctor__((_this).F28(), (_953_v19).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_953_v19).Cardinality()))).Uint32()).(_dafny.Array), p0)
						_954_v20 = _nw147
						var _955_v21 _dafny.Array
						_ = _955_v21
						_955_v21 = (_this).F27()
						var _956_v22 _dafny.Map
						_ = _956_v22
						_956_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _955_v21)
						var _957_v23 _dafny.Set
						_ = _957_v23
						_957_v23 = _dafny.SetOf(_956_v22, _956_v22)
						var _rhs163 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_949_v15, _dafny.SeqOf(p0)), _dafny.SeqOf((_this).F28(), p0))
						_ = _rhs163
						var _rhs164 bool = !(((_957_v23).Difference(_957_v23)).IsSubsetOf(_957_v23))
						_ = _rhs164
						var _rhs165 _dafny.Int = _dafny.IntOfUint32(((_948_v14).F36()).Cardinality())
						_ = _rhs165
						var _rhs166 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(908))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg56 _dafny.Int) interface{} {
								return coer56(arg56)
							}
						}(func(_958_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('o')
						}))
						_ = _rhs166
						var _rhs167 T0 = (func() T0 {
							if !(p0) || ((_this).F28()) {
								return _954_v20
							}
							return _954_v20
						})()
						_ = _rhs167
						var _lhs140 *C3 = _this
						_ = _lhs140
						_949_v15 = _rhs163
						r2 = _rhs164
						_lhs140.F29_set_(_rhs165)
						r1 = _rhs166
						_954_v20 = _rhs167
						var _959_v24 _dafny.MultiSet
						_ = _959_v24
						_959_v24 = _dafny.MultiSetOf((_this).F34())
						_959_v24 = (_959_v24).Intersection(_959_v24)
					} else {
						(globalState).F21 = ((_this).F34()).Plus((_dafny.IntOfInt64(783)).Plus(_this.F29()))
						(globalState).F6 = (_this).F28()
						var _960_v25 _dafny.Array
						_ = _960_v25
						var _len0_22 _dafny.Int = _dafny.IntOfInt64(22)
						_ = _len0_22
						var _nw148 _dafny.Array
						_ = _nw148
						if _len0_22.Cmp(_dafny.Zero) == 0 {
							_nw148 = _dafny.NewArray(_len0_22)
						} else {
							var _init22 func(_dafny.Int) _dafny.Set = (func(_961_v11 _dafny.Map, _962_p0 bool) func(_dafny.Int) _dafny.Set {
								return func(_963_i4 _dafny.Int) _dafny.Set {
									return (_dafny.SetOf((func() bool {
										if (_961_v11).Contains((_this).F28()) {
											return (_961_v11).Get((_this).F28()).(bool)
										}
										return _962_p0
									})(), (func() bool {
										if (_961_v11).Contains((_this).F28()) {
											return (_961_v11).Get((_this).F28()).(bool)
										}
										return (_this).F28()
									})())).Union(_dafny.SetOf((_this).F28()))
								}
							})(_944_v11, p0)
							_ = _init22
							var _element0_22 = _init22(_dafny.Zero)
							_ = _element0_22
							_nw148 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
							(_nw148).ArraySet1(_element0_22, 0)
							var _nativeLen0_22 = (_len0_22).Int()
							_ = _nativeLen0_22
							for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
								(_nw148).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
							}
						}
						_960_v25 = _nw148
						var _964_v26 _dafny.Set
						_ = _964_v26
						_964_v26 = _dafny.SetOf((_this).F28())
						var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_960_v25), 0))
						_ = _index141
						(_960_v25).ArraySet1(_964_v26, (_index141).Int())
						var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_960_v25), 0))
						_ = _index142
						(_960_v25).ArraySet1(_964_v26, (_index142).Int())
						(globalState).F6 = p0
						(globalState).F13 = _this.F29()
					}
					r2 = !((_this).F28())
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _index143
					((_this).F27()).ArraySet1((func() bool {
						if p0 {
							return p0
						}
						return (_this).F28()
					})(), (_index143).Int())
					var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _index144
					((_this).F27()).ArraySet1(p0, (_index144).Int())
					var _965_v27 _dafny.Array
					_ = _965_v27
					var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
					_ = _nw149
					_965_v27 = _nw149
					var _966_v28 _dafny.Array
					_ = _966_v28
					var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
					_ = _nw150
					_966_v28 = _nw150
					var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_965_v27), 0))
					_ = _index145
					(_965_v27).ArraySet1(_966_v28, (_index145).Int())
					var _967_v29 _dafny.Sequence
					_ = _967_v29
					_967_v29 = _dafny.UnicodeSeqOfUtf8Bytes("djlpjj")
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_965_v27), 0))
					_ = _index146
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _index147
					var _rhs168 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_967_v29).Cardinality()))
					_ = _rhs168
					var _rhs169 _dafny.Array = _966_v28
					_ = _rhs169
					var _rhs170 bool = ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool)
					_ = _rhs170
					var _lhs141 *GlobalState = globalState
					_ = _lhs141
					var _lhs142 _dafny.Array = _965_v27
					_ = _lhs142
					var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_965_v27), 0))
					_ = _lhs143
					var _lhs144 _dafny.Array = (_this).F27()
					_ = _lhs144
					var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _lhs145
					_lhs141.F5 = _rhs168
					(_lhs142).ArraySet1(_rhs169, (_lhs143).Int())
					(_lhs144).ArraySet1(_rhs170, (_lhs145).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _968_v30 _dafny.Sequence
		_ = _968_v30
		_968_v30 = _dafny.SeqOf(p0)
		var _969_v31 _dafny.Array
		_ = _969_v31
		var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F28(), p0), _968_v30)
		_ = _nwElement0_30
		var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.One)
		_ = _nw151
		(_nw151).ArraySet1(_nwElement0_30, 0)
		_969_v31 = _nw151
		var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_969_v31), 0))
		_ = _index148
		(_969_v31).ArraySet1(_968_v30, (_index148).Int())
		var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_969_v31), 0))
		_ = _index149
		(_969_v31).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_968_v30, (Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_968_v30).Cardinality()))).Uint32(), (_this).F28()), _968_v30), _968_v30), (_index149).Int())
		var _970_v32 _dafny.Sequence
		_ = _970_v32
		_970_v32 = _dafny.UnicodeSeqOfUtf8Bytes("vn")
		var _971_v33 _dafny.Sequence
		_ = _971_v33
		_971_v33 = _dafny.SeqOf(_this.F29())
		var _972_v34 _dafny.Map
		_ = _972_v34
		_972_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_970_v32).Cardinality())), _971_v33)
		var _973_v35 _dafny.Sequence
		_ = _973_v35
		_973_v35 = (func() _dafny.Sequence {
			if (_972_v34).Contains(_dafny.IntOfInt64(93)) {
				return (_972_v34).Get(_dafny.IntOfInt64(93)).(_dafny.Sequence)
			}
			return _971_v33
		})()
		var _source13 D8 = func(_source14 _dafny.Sequence) D8 {
			var _974___mcc_h4 _dafny.Sequence = _source14
			_ = _974___mcc_h4
			var _975_cf15 _dafny.Sequence = _974___mcc_h4
			_ = _975_cf15
			return Companion_D8_.Create_DC20_(Companion_D8_.Create_DC20_(Companion_D8_.Create_DC20_(Companion_D8_.Create_DC18_((_this).F34()))))
		}(_973_v35)
		_ = _source13
		if _source13.Is_DC18() {
			var _976___mcc_h1 _dafny.Int = _source13.Get_().(D8_DC18).Cf31
			_ = _976___mcc_h1
			var _977_cf31 _dafny.Int = _976___mcc_h1
			_ = _977_cf31
			var _978_v36 D1
			_ = _978_v36
			_978_v36 = Companion_D1_.Create_DC6_((_this).F34())
			var _979_v37 _dafny.Map
			_ = _979_v37
			_979_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F29())
			var _980_v38 _dafny.MultiSet
			_ = _980_v38
			_980_v38 = _dafny.MultiSetOf(_977_cf31)
			var _981_v39 _dafny.Set
			_ = _981_v39
			_981_v39 = _dafny.SetOf(_977_cf31, (_dafny.Zero).Minus(_dafny.IntOfUint32((_970_v32).Cardinality())))
			var _982_v40 _dafny.Map
			_ = _982_v40
			_982_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_981_v39, p0)
			var _983_v41 *C0
			_ = _983_v41
			var _nw152 *C0 = New_C0_()
			_ = _nw152
			_nw152.Ctor__(_970_v32, _977_cf31, (_this).F27(), (_this).F28())
			_983_v41 = _nw152
			var _984_v42 _dafny.Sequence
			_ = _984_v42
			_984_v42 = _dafny.SeqOf(_983_v41, _983_v41)
			var _985_v43 _dafny.CodePoint
			_ = _985_v43
			_985_v43 = _dafny.CodePoint('y')
			var _986_v44 _dafny.Map
			_ = _986_v44
			_986_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_985_v43, (_this).F28())
			var _987_v45 _dafny.Array
			_ = _987_v45
			var _nwElement0_31 _dafny.Int = (_this).F34()
			_ = _nwElement0_31
			var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(25))
			_ = _nw153
			(_nw153).ArraySet1(_nwElement0_31, 0)
			(_nw153).ArraySet1((_978_v36).Dtor_cf12(), 1)
			(_nw153).ArraySet1((_this).Fm14(((_979_v37).Update((_980_v38).Cardinality(), _this.F29())).Cardinality(), _931_v3, globalState), 2)
			(_nw153).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_970_v32).Cardinality()), _dafny.IntOfUint32((_971_v33).Cardinality())), 3)
			(_nw153).ArraySet1(_977_cf31, 4)
			(_nw153).ArraySet1((_dafny.Zero).Minus(_977_cf31), 5)
			(_nw153).ArraySet1(_977_cf31, 6)
			(_nw153).ArraySet1((_this).F34(), 7)
			(_nw153).ArraySet1((_dafny.Zero).Minus((_928_v0).Cardinality()), 8)
			(_nw153).ArraySet1(Companion_Default___.Fm36(globalState), 9)
			(_nw153).ArraySet1(_977_cf31, 10)
			(_nw153).ArraySet1(_dafny.IntOfUint32((_970_v32).Cardinality()), 11)
			(_nw153).ArraySet1((_this).F34(), 12)
			(_nw153).ArraySet1((_dafny.Zero).Minus((_this).F34()), 13)
			(_nw153).ArraySet1(Companion_Default___.Fm36(globalState), 14)
			(_nw153).ArraySet1(_dafny.IntOfInt64(-535), 15)
			(_nw153).ArraySet1((_982_v40).Cardinality(), 16)
			(_nw153).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_970_v32).Cardinality())), 17)
			(_nw153).ArraySet1(_977_cf31, 18)
			(_nw153).ArraySet1((_this).F34(), 19)
			(_nw153).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F29(), (_944_v11).Cardinality()), 20)
			(_nw153).ArraySet1(_977_cf31, 21)
			(_nw153).ArraySet1((_this).F34(), 22)
			(_nw153).ArraySet1(_dafny.IntOfUint32((_984_v42).Cardinality()), 23)
			(_nw153).ArraySet1((Companion_Default___.Fm45(_dafny.IntOfInt64(877), p0, _this.F29(), (func() bool {
				if (_986_v44).Contains(_985_v43) {
					return (_986_v44).Get(_985_v43).(bool)
				}
				return !(p0)
			})(), globalState)).Cardinality(), 24)
			_987_v45 = _nw153
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_987_v45), 0))
			_ = _index150
			(_987_v45).ArraySet1(_977_cf31, (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_987_v45), 0))
			_ = _index151
			(_987_v45).ArraySet1(_this.F29(), (_index151).Int())
			(globalState).F15 = (_this).F28()
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_987_v45), 0))
			_ = _index152
			(_987_v45).ArraySet1((_this.F29()).Minus((_this).F34()), (_index152).Int())
			(globalState).F15 = (_this).F28()
		} else if _source13.Is_DC19() {
			(globalState).F5 = Companion_Default___.Fm36(globalState)
			(globalState).F15 = false
			var _988_v46 _dafny.MultiSet
			_ = _988_v46
			_988_v46 = _dafny.MultiSetOf(_970_v32, _970_v32)
			var _989_v47 _dafny.Map
			_ = _989_v47
			_989_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_988_v46).Cardinality(), (_this).F34())
			var _990_v48 _dafny.Sequence
			_ = _990_v48
			_990_v48 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("uk"))
			var _991_v49 _dafny.CodePoint
			_ = _991_v49
			_991_v49 = _dafny.CodePoint('l')
			_989_v47 = Companion_Default___.Fm46(_dafny.Companion_Sequence_.IsProperPrefixOf(_990_v48, _dafny.SeqOf(_970_v32, Companion_Default___.Fm18(p0, _dafny.IntOfInt64(22), _991_v49, globalState))), globalState)
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_969_v31), 0))
			_ = _index153
			(_969_v31).ArraySet1((_969_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_969_v31), 0))).Int()).(_dafny.Sequence), (_index153).Int())
		} else if _source13.Is_DC17() {
			var _992___mcc_h2 _dafny.Map = _source13.Get_().(D8_DC17).Cf30
			_ = _992___mcc_h2
			var _993_cf30 _dafny.Map = _992___mcc_h2
			_ = _993_cf30
			(globalState).F25 = Companion_Default___.Fm42((_this.F29()).Minus(_dafny.IntOfUint32((_970_v32).Cardinality())), globalState)
			(globalState).F17 = Companion_Default___.Fm40(globalState)
			if !_dafny.Companion_Sequence_.Equal(_970_v32, _970_v32) {
				var _994_v50 _dafny.Set
				_ = _994_v50
				_994_v50 = _dafny.SetOf((_this).F28(), p0, Companion_Default___.Fm40(globalState))
				var _995_v51 _dafny.Set
				_ = _995_v51
				_995_v51 = _dafny.SetOf((_994_v50).IsDisjointFrom(_994_v50))
				_995_v51 = _995_v51
				(globalState).F17 = Companion_Default___.Fm16(_944_v11, globalState)
				(globalState).F15 = (_this).F28()
				(globalState).F5 = (func() _dafny.Int {
					if p0 {
						return _this.F29()
					}
					return _this.F29()
				})()
				var _996_v52 _dafny.Array
				_ = _996_v52
				var _nw154 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(24))
				_ = _nw154
				_996_v52 = _nw154
				var _997_v53 _dafny.CodePoint
				_ = _997_v53
				_997_v53 = _dafny.CodePoint('a')
				var _998_v54 _dafny.Set
				_ = _998_v54
				_998_v54 = _dafny.SetOf(_997_v53)
				var _999_v55 D1
				_ = _999_v55
				_999_v55 = Companion_D1_.Create_DC6_((_998_v54).Cardinality())
				var _1000_v56 _dafny.Map
				_ = _1000_v56
				_1000_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_996_v52, _999_v55)
				var _1001_v57 _dafny.Map
				_ = _1001_v57
				_1001_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(597)).Times((_dafny.Zero).Minus((_this).F34())), (_970_v32).Select((Companion_Default___.SafeIndex((_1000_v56).Cardinality(), _dafny.IntOfUint32((_970_v32).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_1001_v57 = (_1001_v57).Update((_this).F34(), (_970_v32).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_970_v32).Cardinality()))).Uint32()).(_dafny.CodePoint))
			} else {
				var _1002_v58 _dafny.Map
				_ = _1002_v58
				_1002_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _970_v32)
				_970_v32 = _dafny.Companion_Sequence_.Concatenate(_970_v32, _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1002_v58).Contains(_dafny.IntOfUint32((_968_v30).Cardinality())) {
						return (_1002_v58).Get(_dafny.IntOfUint32((_968_v30).Cardinality())).(_dafny.Sequence)
					}
					return _970_v32
				})(), _970_v32))
				var _1003_v59 _dafny.Sequence
				_ = _1003_v59
				_1003_v59 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F28()))
				var _1004_v60 _dafny.Sequence
				_ = _1004_v60
				_1004_v60 = _dafny.SeqOf(_970_v32)
				var _1005_v61 _dafny.Map
				_ = _1005_v61
				_1005_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1003_v59).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1003_v59).Cardinality()))).Uint32()).(_dafny.Map), _1004_v60)
				_1005_v61 = (_1005_v61).Update((_928_v0).Merge(_928_v0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-231))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1006_v32 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1007_i5 _dafny.Int) _dafny.Sequence {
						return _1006_v32
					}
				})(_970_v32))))
				_930_v2 = _930_v2
				(globalState).F13 = _this.F29()
				(globalState).F21 = (_971_v33).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_971_v33).Cardinality()))).Uint32()).(_dafny.Int)
			}
			var _1008_v62 _dafny.Map
			_ = _1008_v62
			_1008_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F29())
			var _1009_v63 *C0
			_ = _1009_v63
			var _nw155 *C0 = New_C0_()
			_ = _nw155
			_nw155.Ctor__(_970_v32, ((_this).F34()).Times((_1008_v62).Cardinality()), (_this).F27(), p0)
			_1009_v63 = _nw155
		} else {
			var _1010___mcc_h3 D8 = _source13.Get_().(D8_DC20).Cf32
			_ = _1010___mcc_h3
			var _1011_cf32 D8 = _1010___mcc_h3
			_ = _1011_cf32
			var _1012_v64 *C1
			_ = _1012_v64
			var _nw156 *C1 = New_C1_()
			_ = _nw156
			_nw156.Ctor__(p0, (_this).F27(), (_this).F28())
			_1012_v64 = _nw156
			var _1013_v65 _dafny.Sequence
			_ = _1013_v65
			_1013_v65 = _dafny.SeqOf(_1012_v64)
			(_this).F29_set_((func() _dafny.Int {
				if !(_dafny.Companion_Sequence_.Contains(_1013_v65, _1012_v64)) {
					return Companion_Default___.SafeModuloInt(_this.F29(), _this.F29())
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_969_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_969_v31), 0))).Int()).(_dafny.Sequence), _968_v30)).Cardinality())
			})())
			(globalState).F15 = (_1012_v64).F39()
			(globalState).F15 = p0
			var _1014_v66 *C0
			_ = _1014_v66
			var _nw157 *C0 = New_C0_()
			_ = _nw157
			_nw157.Ctor__(_970_v32, (_this).F34(), (_this).F27(), (_this).F28())
			_1014_v66 = _nw157
			var _1015_v67 _dafny.Sequence
			_ = _1015_v67
			_1015_v67 = _dafny.SeqOf(_1014_v66)
			var _1016_v68 _dafny.Set
			_ = _1016_v68
			_1016_v68 = _dafny.SetOf((_1015_v67).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1015_v67).Cardinality()))).Uint32()).(*C0), _1014_v66)
			(globalState).F15 = (_dafny.SetOf(_1014_v66)).IsSubsetOf(_1016_v68)
		}
		var _1017_v69 _dafny.Set
		_ = _1017_v69
		_1017_v69 = _dafny.SetOf(_this.F29())
		r0 = (_1017_v69).Cardinality()
		var _1018_v70 _dafny.CodePoint
		_ = _1018_v70
		_1018_v70 = _dafny.CodePoint('r')
		var _1019_v71 D1
		_ = _1019_v71
		_1019_v71 = Companion_D1_.Create_DC7_(_1018_v70, _970_v32)
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_970_v32, _dafny.UnicodeSeqOfUtf8Bytes("dyjwr")), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1019_v71, (_this).F27())).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_970_v32, _dafny.UnicodeSeqOfUtf8Bytes("dyjwr"))).Cardinality()))).Uint32(), _1018_v70)
		r2 = (_this).F28()
		return r0, r1, r2
	}
}
func (_this *C3) F34() _dafny.Int {
	{
		return _this._f34
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
	dummy byte
}

func New_C4_() *C4 {
	_this := C4{}

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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__() {
	{
	}
}
func (_this *C4) Fm12(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		var _source15 D1 = Companion_D1_.Create_DC6_((_dafny.MultiSetOf(_dafny.IntOfInt64(187), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(55))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}(func(_1020_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-962)
		}))).Cardinality()))).Cardinality())
		_ = _source15
		if _source15.Is_DC6() {
			var _1021___mcc_h0 _dafny.Int = _source15.Get_().(D1_DC6).Cf12
			_ = _1021___mcc_h0
			var _1022_cf12 _dafny.Int = _1021___mcc_h0
			_ = _1022_cf12
			return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))).Intersection(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
		} else if _source15.Is_DC7() {
			var _1023___mcc_h1 _dafny.CodePoint = _source15.Get_().(D1_DC7).Cf13
			_ = _1023___mcc_h1
			var _1024___mcc_h2 _dafny.Sequence = _source15.Get_().(D1_DC7).Cf14
			_ = _1024___mcc_h2
			var _1025_cf14 _dafny.Sequence = _1024___mcc_h2
			_ = _1025_cf14
			var _1026_cf13 _dafny.CodePoint = _1023___mcc_h1
			_ = _1026_cf13
			return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(true))))
		} else {
			var _1027___mcc_h3 _dafny.Sequence = _source15.Get_().(D1_DC5).Cf11
			_ = _1027___mcc_h3
			var _1028_cf11 _dafny.Sequence = _1027___mcc_h3
			_ = _1028_cf11
			return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
		}
	}
}
func (_this *C4) Fm13(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(true)
	}
}
func (_this *C4) M8(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.MultiSet, p3 D3, globalState *GlobalState) (_dafny.Int, bool, _dafny.Set, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _1029_v0 _dafny.Sequence
		_ = _1029_v0
		_1029_v0 = _dafny.UnicodeSeqOfUtf8Bytes("xxd")
		var _1030_v1 bool
		_ = _1030_v1
		_1030_v1 = true
		var _1031_v2 _dafny.Array
		_ = _1031_v2
		var _nwElement0_32 bool = _1030_v1
		_ = _nwElement0_32
		var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(9))
		_ = _nw158
		(_nw158).ArraySet1(_nwElement0_32, 0)
		(_nw158).ArraySet1(_1030_v1, 1)
		(_nw158).ArraySet1(_1030_v1, 2)
		(_nw158).ArraySet1(_1030_v1, 3)
		(_nw158).ArraySet1(_1030_v1, 4)
		(_nw158).ArraySet1(_1030_v1, 5)
		(_nw158).ArraySet1(_1030_v1, 6)
		(_nw158).ArraySet1(_1030_v1, 7)
		(_nw158).ArraySet1(_1030_v1, 8)
		_1031_v2 = _nw158
		var _1032_v3 T1
		_ = _1032_v3
		var _nw159 *C0 = New_C0_()
		_ = _nw159
		_nw159.Ctor__(_1029_v0, p1, _1031_v2, _1030_v1)
		_1032_v3 = _nw159
		var _1033_v4 _dafny.Map
		_ = _1033_v4
		_1033_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1032_v3, p1)
		var _1034_v5 _dafny.Map
		_ = _1034_v5
		_1034_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1033_v4, (_1032_v3).F28())
		var _1035_v6 _dafny.Map
		_ = _1035_v6
		_1035_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1032_v3.F29(), Companion_Default___.Fm36(globalState))
		var _1036_v7 _dafny.CodePoint
		_ = _1036_v7
		_1036_v7 = _dafny.CodePoint('e')
		var _1037_v8 D0
		_ = _1037_v8
		_1037_v8 = Companion_D0_.Create_DC0_(_1034_v5, false, (_1035_v6).Cardinality(), _1036_v7, !(_1030_v1))
		r0 = (_1037_v8).Dtor_cf2()
		var _1038_v9 *C1
		_ = _1038_v9
		var _nw160 *C1 = New_C1_()
		_ = _nw160
		_nw160.Ctor__((func() bool {
			if (_1032_v3).F28() {
				return _1030_v1
			}
			return _1030_v1
		})(), (_1032_v3).F27(), _1030_v1)
		_1038_v9 = _nw160
		var _hi3 _dafny.Int = (func() _dafny.Int {
			if (_1038_v9).F39() {
				return _1032_v3.F29()
			}
			return (p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
		})()
		_ = _hi3
		for _1039_i0 := _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1029_v0).Cardinality()), p1, _1032_v3.F29())).Cardinality()); _1039_i0.Cmp(_hi3) < 0; _1039_i0 = _1039_i0.Plus(_dafny.One) {
			var _1040_v10 D0
			_ = _1040_v10
			_1040_v10 = Companion_D0_.Create_DC3_(Companion_Default___.Fm36(globalState), (_1038_v9).F39(), (_1032_v3).F28())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen(((_1032_v3).F27()), 0))
			_ = _index154
			((_1032_v3).F27()).ArraySet1((_1040_v10).Dtor_cf8(), (_index154).Int())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen(((_1032_v3).F27()), 0))
			_ = _index155
			((_1032_v3).F27()).ArraySet1(!(_1030_v1), (_index155).Int())
			var _1041_v11 _dafny.Set
			_ = _1041_v11
			_1041_v11 = _dafny.SetOf((_1038_v9).F39())
			var _1042_v12 _dafny.MultiSet
			_ = _1042_v12
			_1042_v12 = _dafny.MultiSetOf((func() _dafny.Set {
				if ((_1032_v3).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen(((_1032_v3).F27()), 0))).Int()).(bool) {
					return _dafny.SetOf(_1030_v1, ((_1032_v3).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen(((_1032_v3).F27()), 0))).Int()).(bool))
				}
				return _1041_v11
			})())
			_1042_v12 = (func() _dafny.MultiSet {
				if (Companion_D0_.Create_DC3_(_1032_v3.F29(), true, false)).Dtor_cf9() {
					return _1042_v12
				}
				return _1042_v12
			})()
			(globalState).F25 = (_1029_v0).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm36(globalState), _dafny.IntOfUint32((_1029_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen(((_1032_v3).F27()), 0))
			_ = _index156
			((_1032_v3).F27()).ArraySet1(((_1041_v11).Intersection(_1041_v11)).IsProperSubsetOf(_1041_v11), (_index156).Int())
		}
		var _1043_i1 _dafny.Int
		_ = _1043_i1
		_1043_i1 = _dafny.Zero
		{
			for (_1032_v3).F28() {
				{
					if (_1043_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_1043_i1 = (_1043_i1).Plus(_dafny.One)
					var _1044_v13 _dafny.MultiSet
					_ = _1044_v13
					_1044_v13 = _dafny.MultiSetOf(!(!((_1032_v3).F28())), (_1038_v9).F39())
					var _1045_v14 _dafny.Sequence
					_ = _1045_v14
					_1045_v14 = _dafny.SeqOf(_1044_v13)
					var _1046_v15 D9
					_ = _1046_v15
					_1046_v15 = Companion_D9_.Create_DC23_((_1038_v9).F39())
					var _1047_v16 _dafny.Map
					_ = _1047_v16
					_1047_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1038_v9).F39(), _1046_v15)
					var _1048_v17 _dafny.Map
					_ = _1048_v17
					_1048_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1032_v3.F29(), _1047_v16)
					var _1049_v18 _dafny.Map
					_ = _1049_v18
					_1049_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
						if (_1048_v17).Contains((_dafny.Zero).Minus(p1)) {
							return (_1048_v17).Get((_dafny.Zero).Minus(p1)).(_dafny.Map)
						}
						return _1047_v16
					})(), _1030_v1)
					var _1050_v20 _dafny.Array
					_ = _1050_v20
					var _len0_23 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_23
					var _nw161 _dafny.Array
					_ = _nw161
					if _len0_23.Cmp(_dafny.Zero) == 0 {
						_nw161 = _dafny.NewArray(_len0_23)
					} else {
						var _init23 func(_dafny.Int) _dafny.Set = (func(_1051_v3 T1, _1052_p1 _dafny.Int, _1053_v1 bool) func(_dafny.Int) _dafny.Set {
							return func(_1054_i2 _dafny.Int) _dafny.Set {
								return (_dafny.SetOf(_1051_v3.F29(), _1052_p1, _1051_v3.F29())).Intersection(func() _dafny.Set {
									var _coll50 = _dafny.NewBuilder()
									_ = _coll50
									for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-97), _dafny.IntOfInt64(757))); ; {
										_compr_50, _ok52 := _iter52()
										if !_ok52 {
											break
										}
										var _1055_v19 _dafny.Int
										_1055_v19 = interface{}(_compr_50).(_dafny.Int)
										if ((_dafny.IntOfInt64(-97)).Cmp(_1055_v19) <= 0) && ((_1055_v19).Cmp(_dafny.IntOfInt64(757)) < 0) {
											_coll50.Add((_1055_v19).Plus((_dafny.MultiSetOf(_1053_v1)).Cardinality()))
										}
									}
									return _coll50.ToSet()
								}())
							}
						})(_1032_v3, p1, _1030_v1)
						_ = _init23
						var _element0_23 = _init23(_dafny.Zero)
						_ = _element0_23
						_nw161 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
						(_nw161).ArraySet1(_element0_23, 0)
						var _nativeLen0_23 = (_len0_23).Int()
						_ = _nativeLen0_23
						for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
							(_nw161).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
						}
					}
					_1050_v20 = _nw161
					var _1056_v21 _dafny.Set
					_ = _1056_v21
					_1056_v21 = _dafny.SetOf(_dafny.IntOfUint32((_1029_v0).Cardinality()))
					var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1050_v20), 0))
					_ = _index157
					(_1050_v20).ArraySet1((_1056_v21).Intersection(_1056_v21), (_index157).Int())
					var _1057_v22 D0
					_ = _1057_v22
					_1057_v22 = Companion_D0_.Create_DC3_(Companion_Default___.Fm36(globalState), (_1038_v9).F39(), true)
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1050_v20), 0))
					_ = _index158
					var _rhs171 _dafny.Int = _1032_v3.F29()
					_ = _rhs171
					var _rhs172 _dafny.Sequence = _1045_v14
					_ = _rhs172
					var _rhs173 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1047_v16, (_1032_v3.F29()).Cmp(_1032_v3.F29()) != 0)
					_ = _rhs173
					var _rhs174 _dafny.Set = _1056_v21
					_ = _rhs174
					var _rhs175 bool = (_1057_v22).Dtor_cf8()
					_ = _rhs175
					var _lhs146 *GlobalState = globalState
					_ = _lhs146
					var _lhs147 _dafny.Array = _1050_v20
					_ = _lhs147
					var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1050_v20), 0))
					_ = _lhs148
					var _lhs149 *GlobalState = globalState
					_ = _lhs149
					_lhs146.F13 = _rhs171
					_1045_v14 = _rhs172
					_1049_v18 = _rhs173
					(_lhs147).ArraySet1(_rhs174, (_lhs148).Int())
					_lhs149.F17 = _rhs175
					var _1058_v23 _dafny.Array
					_ = _1058_v23
					var _nw162 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(10))
					_ = _nw162
					_1058_v23 = _nw162
					var _1059_v24 _dafny.MultiSet
					_ = _1059_v24
					_1059_v24 = _dafny.MultiSetOf(_1058_v23)
					_1059_v24 = _1059_v24
					_1057_v22 = Companion_D0_.Create_DC3_(_1032_v3.F29(), (_1038_v9).F39(), (_1032_v3).F28())
					(globalState).F4 = (p1).Plus((p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		(globalState).F5 = _1032_v3.F29()
		var _1060_v25 _dafny.Array
		_ = _1060_v25
		var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw163
		_1060_v25 = _nw163
		var _1061_v26 _dafny.Map
		_ = _1061_v26
		_1061_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(417), _1036_v7)
		var _1062_v27 _dafny.Map
		_ = _1062_v27
		_1062_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1032_v3).F28(), (_1038_v9).F39())
		var _1063_v28 _dafny.Array
		_ = _1063_v28
		var _nwElement0_33 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), _1036_v7)
		_ = _nwElement0_33
		var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(16))
		_ = _nw164
		(_nw164).ArraySet1(_nwElement0_33, 0)
		(_nw164).ArraySet1(_1061_v26, 1)
		(_nw164).ArraySet1(_1061_v26, 2)
		(_nw164).ArraySet1(_1061_v26, 3)
		(_nw164).ArraySet1(_1061_v26, 4)
		(_nw164).ArraySet1(_1061_v26, 5)
		(_nw164).ArraySet1(_1061_v26, 6)
		(_nw164).ArraySet1(_1061_v26, 7)
		(_nw164).ArraySet1(_1061_v26, 8)
		(_nw164).ArraySet1(_1061_v26, 9)
		(_nw164).ArraySet1(_1061_v26, 10)
		(_nw164).ArraySet1(_1061_v26, 11)
		(_nw164).ArraySet1(_1061_v26, 12)
		(_nw164).ArraySet1(_1061_v26, 13)
		(_nw164).ArraySet1(_1061_v26, 14)
		(_nw164).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1062_v27).Cardinality(), _1036_v7), 15)
		_1063_v28 = _nw164
		var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_1060_v25), 0))
		_ = _index159
		(_1060_v25).ArraySet1(_1063_v28, (_index159).Int())
		var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen(((_1032_v3).F27()), 0))
		_ = _index160
		((_1032_v3).F27()).ArraySet1(_1030_v1, (_index160).Int())
		var _1064_v29 _dafny.Map
		_ = _1064_v29
		_1064_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1030_v1, _1029_v0)
		var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_1060_v25), 0))
		_ = _index161
		var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen(((_1032_v3).F27()), 0))
		_ = _index162
		var _rhs176 _dafny.Array = _1063_v28
		_ = _rhs176
		var _rhs177 _dafny.Int = (p0).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1064_v29).Contains((_1032_v3).F28()) {
				return (_1064_v29).Get((_1032_v3).F28()).(_dafny.Sequence)
			}
			return _1029_v0
		})()).Cardinality())), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs177
		var _rhs178 bool = (_1032_v3).F28()
		_ = _rhs178
		var _lhs150 _dafny.Array = _1060_v25
		_ = _lhs150
		var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_1060_v25), 0))
		_ = _lhs151
		var _lhs152 *GlobalState = globalState
		_ = _lhs152
		var _lhs153 _dafny.Array = (_1032_v3).F27()
		_ = _lhs153
		var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen(((_1032_v3).F27()), 0))
		_ = _lhs154
		(_lhs150).ArraySet1(_rhs176, (_lhs151).Int())
		_lhs152.F13 = _rhs177
		(_lhs153).ArraySet1(_rhs178, (_lhs154).Int())
		r0 = p1
		r1 = (_1038_v9).F39()
		r2 = func() _dafny.Set {
			var _coll51 = _dafny.NewBuilder()
			_ = _coll51
			for _iter53 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-443), (_1038_v9).F39())).Update(_dafny.IntOfUint32((_1029_v0).Cardinality()), ((_1032_v3).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen(((_1032_v3).F27()), 0))).Int()).(bool))).Keys().Elements()); ; {
				_compr_51, _ok53 := _iter53()
				if !_ok53 {
					break
				}
				var _1065_v30 _dafny.Int
				_1065_v30 = interface{}(_compr_51).(_dafny.Int)
				if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-443), (_1038_v9).F39())).Update(_dafny.IntOfUint32((_1029_v0).Cardinality()), ((_1032_v3).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen(((_1032_v3).F27()), 0))).Int()).(bool))).Contains(_1065_v30) {
					_coll51.Add((_1065_v30).Times((_dafny.Zero).Minus(((_dafny.Zero).Minus((_dafny.SetOf(true, false, true, !(false), true)).Cardinality())).Minus((_dafny.MultiSetOf(_dafny.SeqOf(_dafny.CodePoint('r')))).Cardinality()))))
				}
			}
			return _coll51.ToSet()
		}()
		r3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_1032_v3.F29(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), p1), p0)
		return r0, r1, r2, r3
	}
}
func (_this *C4) M9(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 _dafny.Int, globalState *GlobalState) {
	{
		var _1066_v0 _dafny.Sequence
		_ = _1066_v0
		_1066_v0 = _dafny.UnicodeSeqOfUtf8Bytes("islesyilu")
		var _1067_v1 D8
		_ = _1067_v1
		_1067_v1 = Companion_D8_.Create_DC18_(p3)
		var _1068_v2 _dafny.Array
		_ = _1068_v2
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_24
		var _nw165 _dafny.Array
		_ = _nw165
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw165 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) bool = (func(_1069_p2 bool) func(_dafny.Int) bool {
				return func(_1070_i0 _dafny.Int) bool {
					return _1069_p2
				}
			})(p2)
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw165 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw165).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw165).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_1068_v2 = _nw165
		var _1071_v3 *C0
		_ = _1071_v3
		var _nw166 *C0 = New_C0_()
		_ = _nw166
		_nw166.Ctor__(_1066_v0, (_1067_v1).Dtor_cf31(), _1068_v2, false)
		_1071_v3 = _nw166
		var _1072_v5 _dafny.Set
		_ = _1072_v5
		_1072_v5 = _dafny.SetOf(p0, p3)
		if (func() bool {
			if _dafny.Companion_Sequence_.IsPrefixOf((_1071_v3).F36(), _1066_v0) {
				return (_1072_v5).IsProperSubsetOf(func() _dafny.Set {
					var _coll52 = _dafny.NewBuilder()
					_ = _coll52
					for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(623), _dafny.IntOfInt64(207))); ; {
						_compr_52, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1073_v4 _dafny.Int
						_1073_v4 = interface{}(_compr_52).(_dafny.Int)
						if ((_dafny.IntOfInt64(623)).Cmp(_1073_v4) <= 0) && ((_1073_v4).Cmp(_dafny.IntOfInt64(207)) < 0) {
							_coll52.Add(Companion_Default___.SafeDivisionInt(_1073_v4, p3))
						}
					}
					return _coll52.ToSet()
				}())
			}
			return !(p2)
		})() {
			var _1074_v6 _dafny.CodePoint
			_ = _1074_v6
			_1074_v6 = _dafny.CodePoint('r')
			var _1075_v7 D1
			_ = _1075_v7
			_1075_v7 = Companion_D1_.Create_DC7_(_1074_v6, (_1071_v3).F36())
			var _source16 D1 = (func() D1 {
				if p2 {
					return _1075_v7
				}
				return _1075_v7
			})()
			_ = _source16
			if _source16.Is_DC6() {
				var _1076___mcc_h0 _dafny.Int = _source16.Get_().(D1_DC6).Cf12
				_ = _1076___mcc_h0
				var _1077_cf12 _dafny.Int = _1076___mcc_h0
				_ = _1077_cf12
				var _1078_v8 *C1
				_ = _1078_v8
				var _nw167 *C1 = New_C1_()
				_ = _nw167
				_nw167.Ctor__(p2, _1068_v2, p2)
				_1078_v8 = _nw167
				var _1079_v9 D1
				_ = _1079_v9
				_1079_v9 = Companion_D1_.Create_DC5_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_1080_cf12 _dafny.Int, _1081_p3 _dafny.Int, _1082_v3 *C0, _1083_p0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1084_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_1080_cf12, _1080_cf12, _1081_p3, _dafny.IntOfUint32(((_1082_v3).F36()).Cardinality()), _1083_p0)
					}
				})(_1077_cf12, p3, _1071_v3, p0))))
				var _1085_v10 _dafny.Map
				_ = _1085_v10
				_1085_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus(p3))
				var _1086_v11 _dafny.MultiSet
				_ = _1086_v11
				_1086_v11 = _dafny.MultiSetOf(p3)
				_1079_v9 = Companion_Default___.Fm47(_1085_v10, _1086_v11, p3, globalState)
				(globalState).F15 = (_1078_v8).F39()
				(globalState).F25 = _1074_v6
			} else if _source16.Is_DC7() {
				var _1087___mcc_h1 _dafny.CodePoint = _source16.Get_().(D1_DC7).Cf13
				_ = _1087___mcc_h1
				var _1088___mcc_h2 _dafny.Sequence = _source16.Get_().(D1_DC7).Cf14
				_ = _1088___mcc_h2
				var _1089_cf14 _dafny.Sequence = _1088___mcc_h2
				_ = _1089_cf14
				var _1090_cf13 _dafny.CodePoint = _1087___mcc_h1
				_ = _1090_cf13
				var _1091_v12 *C1
				_ = _1091_v12
				var _nw168 *C1 = New_C1_()
				_ = _nw168
				_nw168.Ctor__(false, _1068_v2, p2)
				_1091_v12 = _nw168
				var _1092_v13 _dafny.Set
				_ = _1092_v13
				_1092_v13 = _dafny.SetOf(_1074_v6, Companion_Default___.Fm42(p0, globalState))
				var _1093_v14 *C0
				_ = _1093_v14
				var _nw169 *C0 = New_C0_()
				_ = _nw169
				_nw169.Ctor__(_1066_v0, (_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality())).Minus((_1092_v13).Cardinality()), _1068_v2, p2)
				_1093_v14 = _nw169
				var _1094_v15 _dafny.Sequence
				_ = _1094_v15
				_1094_v15 = _dafny.SeqOf(p3)
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))
				_ = _index163
				(p1).ArraySet1((_1094_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1094_v15).Cardinality()))).Uint32()).(_dafny.Int), (_index163).Int())
				var _1095_v16 _dafny.Map
				_ = _1095_v16
				_1095_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), p0)
				var _1096_v17 _dafny.Sequence
				_ = _1096_v17
				_1096_v17 = _dafny.SeqOf((_1095_v16).Merge(Companion_Default___.Fm46(!(p2), globalState)), _1095_v16)
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((p1), 0))
				_ = _index164
				(p1).ArraySet1(p3, (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))
				_ = _index165
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((p1), 0))
				_ = _index166
				var _rhs179 _dafny.Int = Companion_Default___.SafeModuloInt(p0, p0)
				_ = _rhs179
				var _rhs180 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1096_v17, _dafny.Companion_Sequence_.Concatenate(_1096_v17, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-711)))))
				_ = _rhs180
				var _rhs181 _dafny.Int = p3
				_ = _rhs181
				var _lhs155 _dafny.Array = p1
				_ = _lhs155
				var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))
				_ = _lhs156
				var _lhs157 _dafny.Array = p1
				_ = _lhs157
				var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((p1), 0))
				_ = _lhs158
				(_lhs155).ArraySet1(_rhs179, (_lhs156).Int())
				_1096_v17 = _rhs180
				(_lhs157).ArraySet1(_rhs181, (_lhs158).Int())
				_1090_cf13 = _1074_v6
			} else {
				var _1097___mcc_h3 _dafny.Sequence = _source16.Get_().(D1_DC5).Cf11
				_ = _1097___mcc_h3
				var _1098_cf11 _dafny.Sequence = _1097___mcc_h3
				_ = _1098_cf11
				var _1099_v18 _dafny.MultiSet
				_ = _1099_v18
				_1099_v18 = _dafny.MultiSetOf(p2, p2, p2, p2, (p3).Cmp(p3) < 0)
				_1099_v18 = ((_1099_v18).Difference(_1099_v18)).Difference(_1099_v18)
				(globalState).F15 = (_1099_v18).IsSubsetOf((_1099_v18).Intersection(_1099_v18))
				_1066_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_1071_v3).F36(), (_1071_v3).F36()), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ppri"), _dafny.UnicodeSeqOfUtf8Bytes("clfofo")))
				var _1100_v19 _dafny.Set
				_ = _1100_v19
				_1100_v19 = _dafny.SetOf(p2, p2)
				var _1101_v20 _dafny.Map
				_ = _1101_v20
				_1101_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1100_v19).Intersection(_1100_v19), false)
				var _1102_v21 _dafny.Sequence
				_ = _1102_v21
				_1102_v21 = _dafny.SeqOf(p2)
				_1101_v20 = (_1101_v20).Update((_1100_v19).Union(_dafny.SetOf(p2)), !((_1102_v21).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1102_v21).Cardinality()))).Uint32()).(bool)))
			}
			var _1103_v22 _dafny.Sequence
			_ = _1103_v22
			_1103_v22 = _dafny.SeqOf(p0)
			var _1104_v23 _dafny.Set
			_ = _1104_v23
			_1104_v23 = _dafny.SetOf(p2)
			var _1105_v24 _dafny.MultiSet
			_ = _1105_v24
			_1105_v24 = _dafny.MultiSetOf((_1103_v22).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1104_v23)).Cardinality()), _dafny.IntOfUint32((_1103_v22).Cardinality()))).Uint32()).(_dafny.Int))
			var _1106_v25 _dafny.Map
			_ = _1106_v25
			_1106_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1105_v24).Update(p0, Companion_Default___.Abs(p3))).Update(p3, Companion_Default___.Abs(p0)), _1074_v6)
			var _1107_v27 _dafny.Map
			_ = _1107_v27
			_1107_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1105_v24, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wydr")).Cardinality()))
			_1106_v25 = (_1106_v25).Merge(func() _dafny.Map {
				var _coll53 = _dafny.NewMapBuilder()
				_ = _coll53
				for _iter55 := _dafny.Iterate((_1107_v27).Keys().Elements()); ; {
					_compr_53, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _1108_v26 _dafny.MultiSet
					_1108_v26 = interface{}(_compr_53).(_dafny.MultiSet)
					if (_1107_v27).Contains(_1108_v26) {
						_coll53.Add(_1108_v26, _1074_v6)
					}
				}
				return _coll53.ToMap()
			}())
			var _1109_v28 _dafny.Sequence
			_ = _1109_v28
			_1109_v28 = _dafny.SeqOf(p2)
			var _1110_v29 D0
			_ = _1110_v29
			_1110_v29 = Companion_D0_.Create_DC3_(_dafny.IntOfUint32((_1109_v28).Cardinality()), !(p2), p2)
			var _1111_v30 _dafny.Map
			_ = _1111_v30
			_1111_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1110_v29, p0)
			var _1112_v31 _dafny.Array
			_ = _1112_v31
			_1112_v31 = _1068_v2
			var _1113_v32 *C3
			_ = _1113_v32
			var _nw170 *C3 = New_C3_()
			_ = _nw170
			_nw170.Ctor__(p0, (_1111_v30).Merge(_1111_v30), (_dafny.Zero).Minus((p0).Plus(p3)), (_1068_v2), p2)
			_1113_v32 = _nw170
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((p1), 0))
			_ = _index167
			(p1).ArraySet1(((_1104_v23).Cardinality()).Minus(p3), (_index167).Int())
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((p1), 0))
			_ = _index168
			(p1).ArraySet1((_1113_v32).F34(), (_index168).Int())
			var _1114_v33 D1
			_ = _1114_v33
			_1114_v33 = Companion_D1_.Create_DC6_(p0)
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((p1), 0))
			_ = _index169
			(p1).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1114_v33).Dtor_cf12()), _1103_v22)).Cardinality())).Minus((_1113_v32).F34()), (_index169).Int())
		} else {
			var _1115_v34 D8
			_ = _1115_v34
			_1115_v34 = Companion_D8_.Create_DC18_((_dafny.Zero).Minus(p3))
			var _1116_v35 D8
			_ = _1116_v35
			_1116_v35 = Companion_D8_.Create_DC20_(_1115_v34)
			var _source17 D8 = _1116_v35
			_ = _source17
			if _source17.Is_DC18() {
				var _1117___mcc_h4 _dafny.Int = _source17.Get_().(D8_DC18).Cf31
				_ = _1117___mcc_h4
				var _1118_cf31 _dafny.Int = _1117___mcc_h4
				_ = _1118_cf31
				var _1119_v36 _dafny.Sequence
				_ = _1119_v36
				_1119_v36 = _dafny.SeqOf(p2, !(!(p2)), !(p2), p2, p2)
				var _1120_v37 _dafny.Map
				_ = _1120_v37
				_1120_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1071_v3).F36()).Cardinality()), p2)
				(globalState).F6 = (func() bool {
					if p2 {
						return (_1119_v36).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p3), _dafny.IntOfUint32((_1119_v36).Cardinality()))).Uint32()).(bool)
					}
					return !((func() bool {
						if (_1120_v37).Contains(_dafny.IntOfInt64(635)) {
							return (_1120_v37).Get(_dafny.IntOfInt64(635)).(bool)
						}
						return p2
					})())
				})()
				var _1121_v38 _dafny.Sequence
				_ = _1121_v38
				_1121_v38 = _dafny.SeqOf(_1072_v5, _1072_v5, _1072_v5)
				var _1122_v39 _dafny.MultiSet
				_ = _1122_v39
				_1122_v39 = _dafny.MultiSetOf(p0, p0, p3)
				(globalState).F17 = ((_1121_v38).Select((Companion_Default___.SafeIndex((_1122_v39).Cardinality(), _dafny.IntOfUint32((_1121_v38).Cardinality()))).Uint32()).(_dafny.Set)).IsSubsetOf(_dafny.SetOf(_1118_cf31))
				var _1123_v40 _dafny.Map
				_ = _1123_v40
				var _1124_v41 _dafny.Int
				_ = _1124_v41
				var _out21 _dafny.Map
				_ = _out21
				var _out22 _dafny.Int
				_ = _out22
				_out21, _out22 = Companion_Default___.M0(globalState)
				_1123_v40 = _out21
				_1124_v41 = _out22
				var _1125_v42 _dafny.Sequence
				_ = _1125_v42
				_1125_v42 = _dafny.SeqOf(p0)
				var _1126_v43 _dafny.Sequence
				_ = _1126_v43
				_1126_v43 = _dafny.SeqOf(_1125_v42, _1125_v42, _1125_v42, _1125_v42, _1125_v42)
				var _1127_v44 _dafny.Sequence
				_ = _1127_v44
				_1127_v44 = (_1126_v43).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1126_v43).Cardinality()))).Uint32()).(_dafny.Sequence)
				var _1128_v45 _dafny.CodePoint
				_ = _1128_v45
				_1128_v45 = _dafny.CodePoint('p')
				var _1129_v46 _dafny.Set
				_ = _1129_v46
				_1129_v46 = _dafny.SetOf(Companion_Default___.Fm26(p2, _1118_cf31, _1128_v45, globalState))
				var _1130_v47 D0
				_ = _1130_v47
				_1130_v47 = Companion_D0_.Create_DC3_((_1071_v3).Fm15(globalState), p2, p2)
				var _rhs182 _dafny.Int = p0
				_ = _rhs182
				var _rhs183 bool = p2
				_ = _rhs183
				var _rhs184 bool = ((_1129_v46).Union(_1129_v46)).IsProperSubsetOf(_1129_v46)
				_ = _rhs184
				var _rhs185 _dafny.Sequence = _1127_v44
				_ = _rhs185
				var _rhs186 bool = (_1130_v47).Dtor_cf9()
				_ = _rhs186
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				_lhs159.F5 = _rhs182
				_lhs160.F6 = _rhs183
				_lhs161.F6 = _rhs184
				_1127_v44 = _rhs185
				_lhs162.F15 = _rhs186
			} else if _source17.Is_DC19() {
				var _1131_v48 _dafny.CodePoint
				_ = _1131_v48
				_1131_v48 = _dafny.CodePoint('r')
				_1066_v0 = _dafny.Companion_Sequence_.Update((_1071_v3).F36(), (Companion_Default___.SafeIndex(Companion_Default___.Fm36(globalState), _dafny.IntOfUint32(((_1071_v3).F36()).Cardinality()))).Uint32(), _1131_v48)
				var _1132_v50 _dafny.Sequence
				_ = _1132_v50
				_1132_v50 = _dafny.SeqOf(_1072_v5)
				var _rhs187 _dafny.Sequence = Companion_Default___.Fm32(p0, p2, globalState)
				_ = _rhs187
				var _rhs188 bool = p2
				_ = _rhs188
				var _rhs189 bool = !((func() _dafny.Map {
					var _coll54 = _dafny.NewMapBuilder()
					_ = _coll54
					for _iter56 := _dafny.Iterate(((_1132_v50).Select((Companion_Default___.SafeIndex((_1072_v5).Cardinality(), _dafny.IntOfUint32((_1132_v50).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
						_compr_54, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1133_v49 _dafny.Int
						_1133_v49 = interface{}(_compr_54).(_dafny.Int)
						if ((_1132_v50).Select((Companion_Default___.SafeIndex((_1072_v5).Cardinality(), _dafny.IntOfUint32((_1132_v50).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_1133_v49) {
							_coll54.Add((_1133_v49).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())), p0)
						}
					}
					return _coll54.ToMap()
				}()).Equals((func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(615), _dafny.IntOfInt64(805))); ; {
						_compr_55, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1134_v51 _dafny.Int
						_1134_v51 = interface{}(_compr_55).(_dafny.Int)
						if ((_dafny.IntOfInt64(615)).Cmp(_1134_v51) <= 0) && ((_1134_v51).Cmp(_dafny.IntOfInt64(805)) < 0) {
							_coll55.Add(Companion_Default___.SafeDivisionInt(_1134_v51, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(37))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg60 _dafny.Int) interface{} {
									return coer60(arg60)
								}
							}((func(_1135_v48 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1136_i2 _dafny.Int) _dafny.CodePoint {
									return _1135_v48
								}
							})(_1131_v48)))).Cardinality()))), _dafny.IntOfInt64(-578))
						}
					}
					return _coll55.ToMap()
				}()).Merge(func() _dafny.Map {
					var _coll56 = _dafny.NewMapBuilder()
					_ = _coll56
					for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(787), _dafny.IntOfInt64(-73))); ; {
						_compr_56, _ok58 := _iter58()
						if !_ok58 {
							break
						}
						var _1137_v52 _dafny.Int
						_1137_v52 = interface{}(_compr_56).(_dafny.Int)
						if ((_dafny.IntOfInt64(787)).Cmp(_1137_v52) <= 0) && ((_1137_v52).Cmp(_dafny.IntOfInt64(-73)) < 0) {
							_coll56.Add(Companion_Default___.SafeModuloInt(_1137_v52, (_dafny.Zero).Minus(p3)), p3)
						}
					}
					return _coll56.ToMap()
				}())))
				_ = _rhs189
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				_1066_v0 = _rhs187
				_lhs163.F17 = _rhs188
				_lhs164.F15 = _rhs189
				var _1138_v53 _dafny.MultiSet
				_ = _1138_v53
				_1138_v53 = _dafny.MultiSetOf(p2, p2)
				var _1139_v54 _dafny.Map
				_ = _1139_v54
				_1139_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1138_v53).Cardinality())
				var _1140_v56 *C1
				_ = _1140_v56
				var _nw171 *C1 = New_C1_()
				_ = _nw171
				_nw171.Ctor__((_1072_v5).IsProperSubsetOf(func() _dafny.Set {
					var _coll57 = _dafny.NewBuilder()
					_ = _coll57
					for _iter59 := _dafny.Iterate((_1139_v54).Keys().Elements()); ; {
						_compr_57, _ok59 := _iter59()
						if !_ok59 {
							break
						}
						var _1141_v55 _dafny.Int
						_1141_v55 = interface{}(_compr_57).(_dafny.Int)
						if (_1139_v54).Contains(_1141_v55) {
							_coll57.Add((_1141_v55).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(665))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg61 _dafny.Int) interface{} {
									return coer61(arg61)
								}
							}(func(_1142_i3 _dafny.Int) _dafny.Int {
								return (_dafny.MultiSetOf(true, false)).Cardinality()
							}))).Cardinality())))
						}
					}
					return _coll57.ToSet()
				}()), _1068_v2, p2)
				_1140_v56 = _nw171
				var _rhs190 *C1 = _1140_v56
				_ = _rhs190
				var _rhs191 bool = p2
				_ = _rhs191
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				_1140_v56 = _rhs190
				_lhs165.F6 = _rhs191
				(globalState).F5 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm36(globalState), Companion_Default___.SafeDivisionInt((_1071_v3).Fm15(globalState), p0))
			} else if _source17.Is_DC17() {
				var _1143___mcc_h5 _dafny.Map = _source17.Get_().(D8_DC17).Cf30
				_ = _1143___mcc_h5
				var _1144_cf30 _dafny.Map = _1143___mcc_h5
				_ = _1144_cf30
				var _1145_v57 _dafny.MultiSet
				_ = _1145_v57
				_1145_v57 = _dafny.MultiSetOf(Companion_D7_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1072_v5).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("sgx"))))
				var _1146_v58 D8
				_ = _1146_v58
				_1146_v58 = Companion_D8_.Create_DC17_(_1144_cf30)
				(globalState).F4 = (func() _dafny.Int {
					if (_1145_v57).Contains(Companion_Default___.Fm48(p3, p2, p0, _1146_v58, globalState)) {
						return (_1145_v57).Multiplicity(Companion_Default___.Fm48(p3, p2, p0, _1146_v58, globalState))
					}
					return p0
				})()
				(globalState).F4 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(p3, p3), p0)
				(globalState).F5 = Companion_Default___.SafeDivisionInt((Companion_Default___.Fm36(globalState)).Minus(p3), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(625), _dafny.IntOfInt64(647)))
				var _1147_v59 _dafny.Array
				_ = _1147_v59
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_25
				var _nw172 _dafny.Array
				_ = _nw172
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw172 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.CodePoint = func(_1148_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('r')
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw172 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw172).ArraySet1CodePoint(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw172).ArraySet1CodePoint(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1147_v59 = _nw172
				var _1149_v60 _dafny.CodePoint
				_ = _1149_v60
				_1149_v60 = _dafny.CodePoint('u')
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1147_v59), 0))
				_ = _index170
				(_1147_v59).ArraySet1CodePoint(_1149_v60, (_index170).Int())
				var _1150_v61 _dafny.Sequence
				_ = _1150_v61
				_1150_v61 = _dafny.SeqOf(_dafny.IntOfInt64(145))
				var _1151_v62 _dafny.MultiSet
				_ = _1151_v62
				_1151_v62 = _dafny.MultiSetOf(((_1150_v61).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1150_v61).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfUint32(((_1071_v3).F36()).Cardinality())) < 0)
				var _1152_v63 _dafny.Sequence
				_ = _1152_v63
				_1152_v63 = _dafny.SeqOf(p2)
				var _1153_v64 _dafny.Map
				_ = _1153_v64
				_1153_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1154_v65 _dafny.Sequence
				_ = _1154_v65
				_1154_v65 = _dafny.SeqOf(_1150_v61)
				var _1155_v66 D0
				_ = _1155_v66
				_1155_v66 = Companion_D0_.Create_DC3_(_dafny.IntOfUint32((_1150_v61).Cardinality()), p2, p2)
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1147_v59), 0))
				_ = _index171
				var _rhs192 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1152_v63, _1152_v63)).Cardinality()), p0)
				_ = _rhs192
				var _rhs193 _dafny.Int = (func() _dafny.Int {
					if (_1153_v64).Contains(Companion_Default___.Fm36(globalState)) {
						return (_1153_v64).Get(Companion_Default___.Fm36(globalState)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1154_v65).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1154_v65).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(p0))).Cardinality())
				})()
				_ = _rhs193
				var _rhs194 _dafny.CodePoint = (func() _dafny.CodePoint {
					if p2 {
						return _1149_v60
					}
					return _1149_v60
				})()
				_ = _rhs194
				var _rhs195 _dafny.MultiSet = (_1151_v62).Union(_dafny.MultiSetOf(p2))
				_ = _rhs195
				var _rhs196 _dafny.Int = Companion_Default___.SafeModuloInt(((Companion_Default___.Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1072_v5), p2, (_1155_v66).Dtor_cf8(), globalState)).Cardinality()).Plus(p3), p0)
				_ = _rhs196
				var _lhs166 *GlobalState = globalState
				_ = _lhs166
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				var _lhs168 _dafny.Array = _1147_v59
				_ = _lhs168
				var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1147_v59), 0))
				_ = _lhs169
				var _lhs170 *GlobalState = globalState
				_ = _lhs170
				_lhs166.F13 = _rhs192
				_lhs167.F21 = _rhs193
				(_lhs168).ArraySet1CodePoint(_rhs194, (_lhs169).Int())
				_1151_v62 = _rhs195
				_lhs170.F5 = _rhs196
			} else {
				var _1156___mcc_h6 D8 = _source17.Get_().(D8_DC20).Cf32
				_ = _1156___mcc_h6
				var _1157_cf32 D8 = _1156___mcc_h6
				_ = _1157_cf32
				_1068_v2 = _1068_v2
				(globalState).F6 = !(p2)
				var _1158_v67 _dafny.Sequence
				_ = _1158_v67
				_1158_v67 = _dafny.SeqOf(p1)
				var _1159_v68 _dafny.Set
				_ = _1159_v68
				_1159_v68 = _dafny.SetOf((_1158_v67).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1158_v67).Cardinality()))).Uint32()).(_dafny.Array), p1)
				var _1160_v69 D0
				_ = _1160_v69
				_1160_v69 = Companion_D0_.Create_DC3_(p0, !(p2), p2)
				var _1161_v70 _dafny.Map
				_ = _1161_v70
				_1161_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1160_v69, p3)
				var _1162_v71 T1
				_ = _1162_v71
				var _nw173 *C3 = New_C3_()
				_ = _nw173
				_nw173.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vfeldy")).Cardinality()), _1161_v70, p0, _1068_v2, p2)
				_1162_v71 = _nw173
				var _1163_v72 _dafny.Map
				_ = _1163_v72
				_1163_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1162_v71, _1162_v71.F29())
				var _1164_v73 _dafny.Map
				_ = _1164_v73
				_1164_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1165_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1166_i5 _dafny.Int) _dafny.Int {
						return _1165_p0
					}
				})(p0))))).Cardinality(), p2)
				var _1167_v74 _dafny.Map
				_ = _1167_v74
				_1167_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v72, (func() bool {
					if (_1164_v73).Contains((_dafny.Zero).Minus(_1162_v71.F29())) {
						return (_1164_v73).Get((_dafny.Zero).Minus(_1162_v71.F29())).(bool)
					}
					return (_1162_v71).F28()
				})())
				var _1168_v75 _dafny.CodePoint
				_ = _1168_v75
				_1168_v75 = _dafny.CodePoint('v')
				var _1169_v76 D0
				_ = _1169_v76
				_1169_v76 = Companion_D0_.Create_DC0_(_1167_v74, p2, _1162_v71.F29(), _1168_v75, !(p2))
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_1068_v2), 0))
				_ = _index172
				(_1068_v2).ArraySet1(!((_1169_v76).Dtor_cf1()), (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_1068_v2), 0))
				_ = _index173
				(_1068_v2).ArraySet1((_1162_v71).F28(), (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen(((_1162_v71).F27()), 0))
				_ = _index174
				((_1162_v71).F27()).ArraySet1(!(p2), (_index174).Int())
				var _1170_v77 _dafny.Map
				_ = _1170_v77
				_1170_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_1068_v2), 0))
				_ = _index175
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_1068_v2), 0))
				_ = _index176
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen(((_1162_v71).F27()), 0))
				_ = _index177
				var _rhs197 _dafny.Int = _dafny.IntOfInt64(720)
				_ = _rhs197
				var _rhs198 _dafny.Set = (func() _dafny.Set {
					if false {
						return _1159_v68
					}
					return _1159_v68
				})()
				_ = _rhs198
				var _rhs199 bool = (_1162_v71).F28()
				_ = _rhs199
				var _rhs200 bool = Companion_Default___.Fm40(globalState)
				_ = _rhs200
				var _rhs201 bool = (((_1170_v77).Merge(_1170_v77)).Cardinality()).Cmp(Companion_Default___.SafeModuloInt(_1162_v71.F29(), p3)) >= 0
				_ = _rhs201
				var _lhs171 *GlobalState = globalState
				_ = _lhs171
				var _lhs172 _dafny.Array = _1068_v2
				_ = _lhs172
				var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_1068_v2), 0))
				_ = _lhs173
				var _lhs174 _dafny.Array = _1068_v2
				_ = _lhs174
				var _lhs175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_1068_v2), 0))
				_ = _lhs175
				var _lhs176 _dafny.Array = (_1162_v71).F27()
				_ = _lhs176
				var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen(((_1162_v71).F27()), 0))
				_ = _lhs177
				_lhs171.F5 = _rhs197
				_1159_v68 = _rhs198
				(_lhs172).ArraySet1(_rhs199, (_lhs173).Int())
				(_lhs174).ArraySet1(_rhs200, (_lhs175).Int())
				(_lhs176).ArraySet1(_rhs201, (_lhs177).Int())
				var _1171_v78 _dafny.MultiSet
				_ = _1171_v78
				_1171_v78 = _dafny.MultiSetOf((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool))
				_1171_v78 = _1171_v78
			}
			(globalState).F15 = true
			(globalState).F6 = !(false)
			_1066_v0 = _dafny.Companion_Sequence_.Concatenate(_1066_v0, _dafny.Companion_Sequence_.Concatenate((_1071_v3).F36(), _1066_v0))
			var _1172_v79 _dafny.Sequence
			_ = _1172_v79
			_1172_v79 = _dafny.SeqOf(p2)
			(globalState).F15 = _dafny.Companion_Sequence_.Contains(_1172_v79, p2)
		}
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
		_ = _index178
		(_1068_v2).ArraySet1(p2, (_index178).Int())
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
		_ = _index179
		var _rhs202 _dafny.Int = (p0).Times(p0)
		_ = _rhs202
		var _rhs203 _dafny.Int = p0
		_ = _rhs203
		var _rhs204 _dafny.Array = _1068_v2
		_ = _rhs204
		var _rhs205 bool = false
		_ = _rhs205
		var _lhs178 *GlobalState = globalState
		_ = _lhs178
		var _lhs179 *GlobalState = globalState
		_ = _lhs179
		var _lhs180 _dafny.Array = _1068_v2
		_ = _lhs180
		var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
		_ = _lhs181
		_lhs178.F5 = _rhs202
		_lhs179.F4 = _rhs203
		_1068_v2 = _rhs204
		(_lhs180).ArraySet1(_rhs205, (_lhs181).Int())
		var _1173_v80 _dafny.Map
		_ = _1173_v80
		_1173_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool), false), p0)
		var _1174_v81 _dafny.Map
		_ = _1174_v81
		_1174_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool), !((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool)))
		var _hi4 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-130))
		_ = _hi4
		for _1175_i6 := (func() _dafny.Int {
			if (_1173_v80).Contains(_1174_v81) {
				return (_1173_v80).Get(_1174_v81).(_dafny.Int)
			}
			return _dafny.IntOfUint32(((_1071_v3).F36()).Cardinality())
		})(); _1175_i6.Cmp(_hi4) < 0; _1175_i6 = _1175_i6.Plus(_dafny.One) {
			var _1176_v82 _dafny.Map
			_ = _1176_v82
			_1176_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1072_v5)
			var _1177_v83 *C2
			_ = _1177_v83
			var _nw174 *C2 = New_C2_()
			_ = _nw174
			_nw174.Ctor__(_dafny.CodePoint('t'), p2, _1068_v2, !((_1072_v5).IsDisjointFrom((func() _dafny.Set {
				if (_1176_v82).Contains((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool)) {
					return (_1176_v82).Get((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool)).(_dafny.Set)
				}
				return _1072_v5
			})())))
			_1177_v83 = _nw174
			var _1178_v84 _dafny.Map
			_ = _1178_v84
			_1178_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p3)
			var _1179_v85 _dafny.Map
			_ = _1179_v85
			_1179_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_1178_v84))
			var _1180_v86 _dafny.Map
			_ = _1180_v86
			_1180_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1072_v5).Cardinality(), (func() _dafny.Sequence {
				if (_1179_v85).Contains((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool)) {
					return (_1179_v85).Get((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool)).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_1178_v84, _1178_v84)
			})())
			var _1181_v87 _dafny.Sequence
			_ = _1181_v87
			_1181_v87 = _dafny.SeqOf(_1178_v84)
			_1180_v86 = (_1180_v86).Update(_1175_i6, _1181_v87)
			var _1182_v88 _dafny.Sequence
			_ = _1182_v88
			_1182_v88 = _dafny.SeqOf(_1177_v83.F38, _1177_v83.F38, _1177_v83.F38)
			var _1183_v89 _dafny.MultiSet
			_ = _1183_v89
			_1183_v89 = _dafny.MultiSetOf(_1177_v83.F38)
			if (_1183_v89).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1182_v88, _1182_v88), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1071_v3).F36()).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1182_v88, _1182_v88)).Cardinality()))).Uint32(), (_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool)))) {
				var _1184_v90 _dafny.Sequence
				_ = _1184_v90
				_1184_v90 = _dafny.SeqOf(_1068_v2)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
				_ = _index180
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
				_ = _index181
				var _rhs206 bool = !_dafny.Companion_Sequence_.Equal(_1184_v90, _1184_v90)
				_ = _rhs206
				var _rhs207 bool = (_dafny.IntOfInt64(-605)).Cmp(p0) == 0
				_ = _rhs207
				var _rhs208 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm18(p2, p3, (_1177_v83).F37(), globalState)).Cardinality())
				_ = _rhs208
				var _lhs182 _dafny.Array = _1068_v2
				_ = _lhs182
				var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
				_ = _lhs183
				var _lhs184 _dafny.Array = _1068_v2
				_ = _lhs184
				var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))
				_ = _lhs185
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				(_lhs182).ArraySet1(_rhs206, (_lhs183).Int())
				(_lhs184).ArraySet1(_rhs207, (_lhs185).Int())
				_lhs186.F5 = _rhs208
				var _1185_v91 _dafny.Map
				_ = _1185_v91
				_1185_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1178_v84, _1175_i6)
				var _1186_v92 D8
				_ = _1186_v92
				_1186_v92 = Companion_D8_.Create_DC17_(_1185_v91)
				var _1187_v93 D8
				_ = _1187_v93
				_1187_v93 = Companion_D8_.Create_DC20_(_1186_v92)
				var _rhs209 _dafny.Int = _dafny.IntOfInt64(894)
				_ = _rhs209
				var _rhs210 D8 = _1187_v93
				_ = _rhs210
				var _lhs187 *GlobalState = globalState
				_ = _lhs187
				_lhs187.F13 = _rhs209
				_1187_v93 = _rhs210
				(globalState).F4 = Companion_Default___.SafeModuloInt(p0, p3)
				_1066_v0 = (_1071_v3).F36()
				var _1188_v94 _dafny.Sequence
				_ = _1188_v94
				_1188_v94 = _dafny.SeqOf(_1066_v0)
				var _1189_v95 _dafny.Map
				_ = _1189_v95
				_1189_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1188_v94, _1188_v94)).Cardinality()), _1068_v2)
				_1068_v2 = (func() _dafny.Array {
					if (_1189_v95).Contains(_1175_i6) {
						return (_1189_v95).Get(_1175_i6).(_dafny.Array)
					}
					return _1068_v2
				})()
			} else {
				var _1190_v96 _dafny.Set
				_ = _1190_v96
				_1190_v96 = _dafny.SetOf(true)
				_1190_v96 = _1190_v96
				var _1191_v97 _dafny.Array
				_ = _1191_v97
				var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
				_ = _nw175
				_1191_v97 = _nw175
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1191_v97), 0))
				_ = _index182
				(_1191_v97).ArraySet1CodePoint((_1177_v83).F37(), (_index182).Int())
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1191_v97), 0))
				_ = _index183
				(_1191_v97).ArraySet1CodePoint((_1177_v83).F37(), (_index183).Int())
				var _1192_v98 _dafny.Map
				_ = _1192_v98
				_1192_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1175_i6, _1190_v96)
				_1192_v98 = (_1192_v98).Update((func() _dafny.Int {
					if (_1178_v84).Contains(_1177_v83.F38) {
						return (_1178_v84).Get(_1177_v83.F38).(_dafny.Int)
					}
					return _1175_i6
				})(), (_1190_v96).Difference(_1190_v96))
				(globalState).F21 = (Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(p3)))).Minus((_1177_v83).Fm20(globalState))
				var _1193_v99 _dafny.Map
				_ = _1193_v99
				_1193_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1178_v84, p0)
				var _1194_v100 D8
				_ = _1194_v100
				_1194_v100 = Companion_D8_.Create_DC20_(Companion_D8_.Create_DC17_(_1193_v99))
				_1194_v100 = _1194_v100
			}
			var _1195_v101 _dafny.Array
			_ = _1195_v101
			var _nwElement0_34 _dafny.Sequence = (_1071_v3).F36()
			_ = _nwElement0_34
			var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(14))
			_ = _nw176
			(_nw176).ArraySet1(_nwElement0_34, 0)
			(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("meyjff"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(595))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_1196_v83 *C2) func(_dafny.Int) _dafny.CodePoint {
				return func(_1197_i7 _dafny.Int) _dafny.CodePoint {
					return (_1196_v83).F37()
				}
			})(_1177_v83)))), 1)
			(_nw176).ArraySet1(_1066_v0, 2)
			(_nw176).ArraySet1((_1071_v3).F36(), 3)
			(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kqcgi"), (_1071_v3).F36()), 4)
			(_nw176).ArraySet1((_1071_v3).F36(), 5)
			(_nw176).ArraySet1(_dafny.Companion_Sequence_.Update((_1071_v3).F36(), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32(((_1071_v3).F36()).Cardinality()))).Uint32(), _dafny.CodePoint('t')), 6)
			(_nw176).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("krlvulblo"), 7)
			(_nw176).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("b"), 8)
			(_nw176).ArraySet1((_1071_v3).F36(), 9)
			(_nw176).ArraySet1(_1066_v0, 10)
			(_nw176).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}((func(_1198_v83 *C2) func(_dafny.Int) _dafny.CodePoint {
				return func(_1199_i8 _dafny.Int) _dafny.CodePoint {
					return (_1198_v83).F37()
				}
			})(_1177_v83))), 11)
			(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1066_v0, (_1071_v3).F36()), 12)
			(_nw176).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(733))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}((func(_1200_v83 *C2) func(_dafny.Int) _dafny.CodePoint {
				return func(_1201_i9 _dafny.Int) _dafny.CodePoint {
					return (_1200_v83).F37()
				}
			})(_1177_v83))), 13)
			_1195_v101 = _nw176
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_1195_v101), 0))
			_ = _index184
			(_1195_v101).ArraySet1((_1071_v3).F36(), (_index184).Int())
			var _1202_v102 D9
			_ = _1202_v102
			_1202_v102 = Companion_D9_.Create_DC21_(_1195_v101)
			var _pat_let_tv27 = _1202_v102
			_ = _pat_let_tv27
			var _1203_v103 _dafny.Map
			_ = _1203_v103
			_1203_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let19_0 D9) D9 {
				return func(_1204_dt__update__tmp_h1 D9) D9 {
					return func(_pat_let20_0 D9) D9 {
						return func(_1205_dt__update_hcf39_h0 D9) D9 {
							return Companion_D9_.Create_DC24_(_1205_dt__update_hcf39_h0)
						}(_pat_let20_0)
					}(_pat_let_tv27)
				}(_pat_let19_0)
			}(Companion_D9_.Create_DC24_(_1202_v102)), _dafny.UnicodeSeqOfUtf8Bytes("wgbvpofn"))
			var _1206_v104 D9
			_ = _1206_v104
			_1206_v104 = Companion_D9_.Create_DC24_(_1202_v102)
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_1195_v101), 0))
			_ = _index185
			(_1195_v101).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1071_v3).F36(), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1203_v103).Contains(_1206_v104) {
					return (_1203_v103).Get(_1206_v104).(_dafny.Sequence)
				}
				return _1066_v0
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1203_v103).Contains(_1206_v104) {
					return (_1203_v103).Get(_1206_v104).(_dafny.Sequence)
				}
				return _1066_v0
			})()).Cardinality()))).Uint32(), (_1177_v83).F37())), (_index185).Int())
		}
		var _1207_v105 _dafny.Sequence
		_ = _1207_v105
		_1207_v105 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2))
		(globalState).F5 = (_dafny.IntOfUint32((_1207_v105).Cardinality())).Minus((p0).Plus(p0))
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((p1), 0))
		_ = _index186
		(p1).ArraySet1(p3, (_index186).Int())
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((p1), 0))
		_ = _index187
		(p1).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool), (_1068_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_1068_v2), 0))).Int()).(bool))).Cardinality())), (_index187).Int())
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f29 _dafny.Int
	_f27 _dafny.Array
	_f28 bool
	_f33 _dafny.Sequence
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f29 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
	_this._f33 = _dafny.EmptySeq
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

func (_this *C5) F29() _dafny.Int {
	return _this._f29
}
func (_this *C5) F29_set_(value _dafny.Int) {
	_this._f29 = value
}
func (_this *C5) F27() _dafny.Array {
	return _this._f27
}
func (_this *C5) F28() bool {
	return _this._f28
}
func (_this *C5) Ctor__(f33 _dafny.Sequence, f29 _dafny.Int, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f33 = f33
		(_this)._f29 = f29
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C5) Fm6(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F29()
	}
}
func (_this *C5) Fm7(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_this.F29()).Times(_this.F29())
	}
}
func (_this *C5) M7(p0 bool, p1 _dafny.Map, p2 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1208_v0 _dafny.Array
		_ = _1208_v0
		var _len0_26 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_26
		var _nw177 _dafny.Array
		_ = _nw177
		if _len0_26.Cmp(_dafny.Zero) == 0 {
			_nw177 = _dafny.NewArray(_len0_26)
		} else {
			var _init26 func(_dafny.Int) _dafny.Int = func(_1209_i0 _dafny.Int) _dafny.Int {
				return (_1209_i0).Times(_this.F29())
			}
			_ = _init26
			var _element0_26 = _init26(_dafny.Zero)
			_ = _element0_26
			_nw177 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
			(_nw177).ArraySet1(_element0_26, 0)
			var _nativeLen0_26 = (_len0_26).Int()
			_ = _nativeLen0_26
			for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
				(_nw177).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
			}
		}
		_1208_v0 = _nw177
		var _1210_v1 _dafny.MultiSet
		_ = _1210_v1
		_1210_v1 = _dafny.MultiSetOf(p0, true)
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))
		_ = _index188
		(_1208_v0).ArraySet1((_1210_v1).Cardinality(), (_index188).Int())
		var _1211_v2 D3
		_ = _1211_v2
		_1211_v2 = Companion_D3_.Create_DC10_(p0, (_this).F28())
		var _1212_v3 _dafny.Set
		_ = _1212_v3
		_1212_v3 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(!((_this).F28()))).Cardinality()))
		var _1213_v4 _dafny.Map
		_ = _1213_v4
		_1213_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1211_v2, _1212_v3)
		var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))
		_ = _index189
		(_1208_v0).ArraySet1(((func() _dafny.Map {
			if p0 {
				return _1213_v4
			}
			return (_1213_v4).Merge(_1213_v4)
		})()).Cardinality(), (_index189).Int())
		var _hi5 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F29(), (_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int)))
		_ = _hi5
		for _1214_i1 := (_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int); _1214_i1.Cmp(_hi5) < 0; _1214_i1 = _1214_i1.Plus(_dafny.One) {
			(globalState).F17 = p0
			var _1215_v8 _dafny.MultiSet
			_ = _1215_v8
			_1215_v8 = _dafny.MultiSetOf(func() _dafny.Set {
				var _coll58 = _dafny.NewBuilder()
				_ = _coll58
				for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-224), _dafny.IntOfInt64(926))); ; {
					_compr_58, _ok60 := _iter60()
					if !_ok60 {
						break
					}
					var _1216_v5 _dafny.Int
					_1216_v5 = interface{}(_compr_58).(_dafny.Int)
					if ((_dafny.IntOfInt64(-224)).Cmp(_1216_v5) <= 0) && ((_1216_v5).Cmp(_dafny.IntOfInt64(926)) < 0) {
						_coll58.Add((_1216_v5).Times((func() _dafny.Set {
							var _coll59 = _dafny.NewBuilder()
							_ = _coll59
							for _iter61 := _dafny.Iterate((func() _dafny.Map {
								var _coll60 = _dafny.NewMapBuilder()
								_ = _coll60
								for _iter62 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _1214_i1)).Keys().Elements()); ; {
									_compr_60, _ok62 := _iter62()
									if !_ok62 {
										break
									}
									var _1217_v6 _dafny.CodePoint
									_1217_v6 = interface{}(_compr_60).(_dafny.CodePoint)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _1214_i1)).Contains(_1217_v6) {
										_coll60.Add(_1217_v6, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(853))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg66 _dafny.Int) interface{} {
												return coer66(arg66)
											}
										}(func(_1218_i2 _dafny.Int) _dafny.CodePoint {
											return _dafny.CodePoint('l')
										}))).Cardinality()))
									}
								}
								return _coll60.ToMap()
							}()).Keys().Elements()); ; {
								_compr_59, _ok61 := _iter61()
								if !_ok61 {
									break
								}
								var _1219_v7 _dafny.CodePoint
								_1219_v7 = interface{}(_compr_59).(_dafny.CodePoint)
								if (func() _dafny.Map {
									var _coll61 = _dafny.NewMapBuilder()
									_ = _coll61
									for _iter63 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _1214_i1)).Keys().Elements()); ; {
										_compr_61, _ok63 := _iter63()
										if !_ok63 {
											break
										}
										var _1217_v6 _dafny.CodePoint
										_1217_v6 = interface{}(_compr_61).(_dafny.CodePoint)
										if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _1214_i1)).Contains(_1217_v6) {
											_coll61.Add(_1217_v6, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(853))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
												return func(arg67 _dafny.Int) interface{} {
													return coer67(arg67)
												}
											}(func(_1218_i2 _dafny.Int) _dafny.CodePoint {
												return _dafny.CodePoint('l')
											}))).Cardinality()))
										}
									}
									return _coll61.ToMap()
								}()).Contains(_1219_v7) {
									_coll59.Add(_1219_v7)
								}
							}
							return _coll59.ToSet()
						}()).Cardinality()))
					}
				}
				return _coll58.ToSet()
			}(), _1212_v3, (_1212_v3).Intersection(_1212_v3))
			var _1220_v9 _dafny.Sequence
			_ = _1220_v9
			_1220_v9 = _dafny.SeqOf(_1212_v3)
			_1215_v8 = ((_dafny.MultiSetFromSeq(_1220_v9)).Difference(_1215_v8)).Update(_1212_v3, Companion_Default___.Abs(((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int)).Times(_this.F29())))
			var _1221_v10 _dafny.CodePoint
			_ = _1221_v10
			_1221_v10 = _dafny.CodePoint('g')
			var _1222_v11 D1
			_ = _1222_v11
			_1222_v11 = Companion_D1_.Create_DC7_(_1221_v10, Companion_Default___.Fm9((_this).F28(), globalState))
			var _source18 D3 = Companion_Default___.Fm8((func() D1 {
				if true {
					return _1222_v11
				}
				return _1222_v11
			})(), globalState)
			_ = _source18
			if _source18.Is_DC10() {
				var _1223___mcc_h0 bool = _source18.Get_().(D3_DC10).Cf17
				_ = _1223___mcc_h0
				var _1224___mcc_h1 bool = _source18.Get_().(D3_DC10).Cf18
				_ = _1224___mcc_h1
				var _1225_cf18 bool = _1224___mcc_h1
				_ = _1225_cf18
				var _1226_cf17 bool = _1223___mcc_h0
				_ = _1226_cf17
				_1226_cf17 = (_this).F28()
				var _1227_v12 _dafny.Sequence
				_ = _1227_v12
				_1227_v12 = _dafny.UnicodeSeqOfUtf8Bytes("d")
				var _1228_v14 _dafny.MultiSet
				_ = _1228_v14
				_1228_v14 = _dafny.MultiSetOf(_1227_v12, _1227_v12)
				(globalState).F4 = ((func() _dafny.Map {
					if true {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1227_v12, _1222_v11)
					}
					return func() _dafny.Map {
						var _coll62 = _dafny.NewMapBuilder()
						_ = _coll62
						for _iter64 := _dafny.Iterate((_1228_v14).Elements()); ; {
							_compr_62, _ok64 := _iter64()
							if !_ok64 {
								break
							}
							var _1229_v13 _dafny.Sequence
							_1229_v13 = interface{}(_compr_62).(_dafny.Sequence)
							if (_1228_v14).Contains(_1229_v13) {
								_coll62.Add(_1229_v13, _1222_v11)
							}
						}
						return _coll62.ToMap()
					}()
				})()).Cardinality()
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))
				_ = _index190
				(_1208_v0).ArraySet1(_1214_i1, (_index190).Int())
				var _1230_v15 _dafny.Array
				_ = _1230_v15
				var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw178
				_1230_v15 = _nw178
				var _1231_v16 _dafny.Map
				_ = _1231_v16
				_1231_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1230_v15, _dafny.IntOfInt64(594))
				(globalState).F21 = (func() _dafny.Int {
					if (_1231_v16).Contains(_1230_v15) {
						return (_1231_v16).Get(_1230_v15).(_dafny.Int)
					}
					return _this.F29()
				})()
			} else {
				var _1232___mcc_h2 _dafny.Map = _source18.Get_().(D3_DC9).Cf16
				_ = _1232___mcc_h2
				var _1233_cf16 _dafny.Map = _1232___mcc_h2
				_ = _1233_cf16
				(globalState).F25 = _1221_v10
				var _1234_v17 _dafny.Map
				_ = _1234_v17
				_1234_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(_dafny.CodePoint('u'), globalState), (_this).F28())
				var _1235_v18 _dafny.Sequence
				_ = _1235_v18
				_1235_v18 = _dafny.UnicodeSeqOfUtf8Bytes("jrlr")
				_1234_v17 = Companion_Default___.Fm11((_1210_v1).IsProperSubsetOf(_1210_v1), !(_dafny.Companion_Sequence_.IsPrefixOf(_1235_v18, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-321))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}((func(_1236_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1237_i3 _dafny.Int) _dafny.CodePoint {
						return _1236_v10
					}
				})(_1221_v10))))), p0, globalState)
				var _1238_v19 D0
				_ = _1238_v19
				_1238_v19 = Companion_D0_.Create_DC3_(_this.F29(), p0, p0)
				var _1239_v20 _dafny.Map
				_ = _1239_v20
				_1239_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1238_v19, _1214_i1)
				var _1240_v21 *C3
				_ = _1240_v21
				var _nw179 *C3 = New_C3_()
				_ = _nw179
				_nw179.Ctor__((_this.F29()).Times((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int)), _1239_v20, _this.F29(), (_this).F27(), p0)
				_1240_v21 = _nw179
				var _1241_v22 _dafny.Array
				_ = _1241_v22
				var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw180
				_1241_v22 = _nw180
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1241_v22), 0))
				_ = _index191
				(_1241_v22).ArraySet1(_1235_v18, (_index191).Int())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1241_v22), 0))
				_ = _index192
				(_1241_v22).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1235_v18, (func() _dafny.Sequence {
					if (p2).Contains(p0) {
						return (p2).Get(p0).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Update(_1235_v18, (Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1235_v18).Cardinality()))).Uint32(), _1221_v10)
				})()), (_index192).Int())
			}
			(globalState).F4 = (_dafny.Zero).Minus(_this.F29())
		}
		var _1242_i4 _dafny.Int
		_ = _1242_i4
		_1242_i4 = _dafny.Zero
		{
			for p0 {
				{
					if (_1242_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1242_i4 = (_1242_i4).Plus(_dafny.One)
					var _1243_v23 _dafny.MultiSet
					_ = _1243_v23
					_1243_v23 = _dafny.MultiSetOf((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int))
					var _1244_v24 _dafny.Map
					_ = _1244_v24
					_1244_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int), _1208_v0)
					var _1245_v25 _dafny.Map
					_ = _1245_v25
					_1245_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1244_v24).Cardinality())
					var _1246_v26 _dafny.Map
					_ = _1246_v26
					_1246_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _1245_v25)
					var _1247_v27 _dafny.MultiSet
					_ = _1247_v27
					_1247_v27 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int))).Update(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1243_v23)).Cardinality()), _1245_v25, (func() _dafny.Map {
						if (_1246_v26).Contains((_this).F28()) {
							return (_1246_v26).Get((_this).F28()).(_dafny.Map)
						}
						return _1245_v25
					})())
					(globalState).F6 = (_1247_v27).IsDisjointFrom((_1247_v27).Update(_1245_v25, Companion_Default___.Abs(_dafny.IntOfInt64(922))))
					var _1248_v28 _dafny.Sequence
					_ = _1248_v28
					_1248_v28 = _dafny.SeqOf((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int), (_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int), _this.F29(), _dafny.IntOfInt64(441))
					var _1249_v29 _dafny.Sequence
					_ = _1249_v29
					_1249_v29 = _dafny.UnicodeSeqOfUtf8Bytes("mcumuyosc")
					var _1250_v30 _dafny.Sequence
					_ = _1250_v30
					_1250_v30 = _dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm25((_1248_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1249_v29).Cardinality()), _dafny.IntOfUint32((_1248_v28).Cardinality()))).Uint32()).(_dafny.Int), (_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int), _this.F29(), globalState)).Cardinality()))
					var _1251_v31 _dafny.Map
					_ = _1251_v31
					_1251_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())
					var _1252_v32 _dafny.CodePoint
					_ = _1252_v32
					_1252_v32 = _dafny.CodePoint('w')
					var _1253_v33 *C2
					_ = _1253_v33
					var _nw181 *C2 = New_C2_()
					_ = _nw181
					_nw181.Ctor__(_1252_v32, p0, (_this).F27(), true)
					_1253_v33 = _nw181
					var _1254_v34 _dafny.Map
					_ = _1254_v34
					_1254_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_v33, _1252_v32)
					(globalState).F13 = (_dafny.IntOfUint32((_1248_v28).Cardinality())).Minus(((func() _dafny.Map {
						if Companion_Default___.Fm16(_1251_v31, globalState) {
							return _1254_v34
						}
						return _1254_v34
					})()).Cardinality())
					_1250_v30 = _1250_v30
					var _1255_v35 *C4
					_ = _1255_v35
					var _nw182 *C4 = New_C4_()
					_ = _nw182
					_nw182.Ctor__()
					_1255_v35 = _nw182
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1256_v36 _dafny.Sequence
		_ = _1256_v36
		_1256_v36 = _dafny.SeqOf((_dafny.Zero).Minus((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int)))
		var _1257_v37 _dafny.Map
		_ = _1257_v37
		_1257_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_1256_v36))).IsSubsetOf(_dafny.SetOf(_1256_v36, _1256_v36)), (_dafny.IntOfInt64(243)).Cmp((_1208_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1208_v0), 0))).Int()).(_dafny.Int)) != 0)
		_1257_v37 = _1257_v37
		var _1258_v38 _dafny.Sequence
		_ = _1258_v38
		_1258_v38 = _dafny.UnicodeSeqOfUtf8Bytes("g")
		(globalState).F5 = _dafny.IntOfUint32((_1258_v38).Cardinality())
		(globalState).F15 = !(!(p0) || (p0))
		r0 = _this.F29()
		return r0
	}
}
func (_this *C5) F33() _dafny.Sequence {
	{
		return _this._f33
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f29 _dafny.Int
	_f27 _dafny.Array
	_f28 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f29 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
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

func (_this *C6) F29() _dafny.Int {
	return _this._f29
}
func (_this *C6) F29_set_(value _dafny.Int) {
	_this._f29 = value
}
func (_this *C6) F27() _dafny.Array {
	return _this._f27
}
func (_this *C6) F28() bool {
	return _this._f28
}
func (_this *C6) Ctor__(f29 _dafny.Int, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f29 = f29
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C6) Fm5(globalState *GlobalState) _dafny.Int {
	{
		return _this.F29()
	}
}
func (_this *C6) M6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) {
	{
		(globalState).F6 = !(p0)
		var _1259_v0 _dafny.Set
		_ = _1259_v0
		_1259_v0 = _dafny.SetOf(p2, _dafny.IntOfInt64(-885), p2)
		var _1260_v1 _dafny.CodePoint
		_ = _1260_v1
		_1260_v1 = _dafny.CodePoint('g')
		var _1261_v2 _dafny.Map
		_ = _1261_v2
		_1261_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _1260_v1)
		var _1262_v4 _dafny.Map
		_ = _1262_v4
		_1262_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F29()), func() _dafny.Set {
			var _coll63 = _dafny.NewBuilder()
			_ = _coll63
			for _iter65 := _dafny.Iterate((_1261_v2).Keys().Elements()); ; {
				_compr_63, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1263_v3 _dafny.Int
				_1263_v3 = interface{}(_compr_63).(_dafny.Int)
				if (_1261_v2).Contains(_1263_v3) {
					_coll63.Add((_1263_v3).Minus(_dafny.IntOfInt64(-91)))
				}
			}
			return _coll63.ToSet()
		}())
		var _1264_v5 D3
		_ = _1264_v5
		_1264_v5 = Companion_D3_.Create_DC9_(_1262_v4)
		_1259_v0 = Companion_Default___.Fm0((_1264_v5).Dtor_cf16(), p0, !(((_this).F28()) == (p1)), globalState)
		var _1265_v6 _dafny.Map
		_ = _1265_v6
		var _1266_v7 _dafny.Int
		_ = _1266_v7
		var _out23 _dafny.Map
		_ = _out23
		var _out24 _dafny.Int
		_ = _out24
		_out23, _out24 = Companion_Default___.M0(globalState)
		_1265_v6 = _out23
		_1266_v7 = _out24
		var _1267_v8 _dafny.Sequence
		_ = _1267_v8
		_1267_v8 = _dafny.SeqOf(_1259_v0)
		var _1268_v9 _dafny.Sequence
		_ = _1268_v9
		_1268_v9 = _dafny.UnicodeSeqOfUtf8Bytes("kwuoanirg")
		var _1269_v10 D0
		_ = _1269_v10
		_1269_v10 = Companion_D0_.Create_DC2_((_1267_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1268_v9).Cardinality()), _dafny.IntOfUint32((_1267_v8).Cardinality()))).Uint32()).(_dafny.Set), _1268_v9)
		var _source19 D0 = _1269_v10
		_ = _source19
		if _source19.Is_DC0() {
			var _1270___mcc_h0 _dafny.Map = _source19.Get_().(D0_DC0).Cf0
			_ = _1270___mcc_h0
			var _1271___mcc_h1 bool = _source19.Get_().(D0_DC0).Cf1
			_ = _1271___mcc_h1
			var _1272___mcc_h2 _dafny.Int = _source19.Get_().(D0_DC0).Cf2
			_ = _1272___mcc_h2
			var _1273___mcc_h3 _dafny.CodePoint = _source19.Get_().(D0_DC0).Cf3
			_ = _1273___mcc_h3
			var _1274___mcc_h4 bool = _source19.Get_().(D0_DC0).Cf4
			_ = _1274___mcc_h4
			var _1275_cf4 bool = _1274___mcc_h4
			_ = _1275_cf4
			var _1276_cf3 _dafny.CodePoint = _1273___mcc_h3
			_ = _1276_cf3
			var _1277_cf2 _dafny.Int = _1272___mcc_h2
			_ = _1277_cf2
			var _1278_cf1 bool = _1271___mcc_h1
			_ = _1278_cf1
			var _1279_cf0 _dafny.Map = _1270___mcc_h0
			_ = _1279_cf0
			var _1280_v11 _dafny.Array
			_ = _1280_v11
			var _nw183 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw183
			_1280_v11 = _nw183
			_1280_v11 = (_this).F27()
			var _1281_v12 _dafny.Array
			_ = _1281_v12
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_27
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Int = (func(_1282_p1 bool, _1283_cf4 bool) func(_dafny.Int) _dafny.Int {
					return func(_1284_i0 _dafny.Int) _dafny.Int {
						return (_1284_i0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_p1, _1283_cf4)).Cardinality())
					}
				})(p1, _1275_cf4)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw184 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw184).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw184).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1281_v12 = _nw184
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1281_v12), 0))
			_ = _index193
			(_1281_v12).ArraySet1(_this.F29(), (_index193).Int())
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1281_v12), 0))
			_ = _index194
			(_1281_v12).ArraySet1(_this.F29(), (_index194).Int())
			var _1285_v13 D0
			_ = _1285_v13
			_1285_v13 = Companion_D0_.Create_DC1_()
			_1285_v13 = _1285_v13
			var _1286_v14 *C2
			_ = _1286_v14
			var _nw185 *C2 = New_C2_()
			_ = _nw185
			_nw185.Ctor__(_1276_cf3, true, _1280_v11, false)
			_1286_v14 = _nw185
		} else if _source19.Is_DC1() {
			var _1287_v15 _dafny.Map
			_ = _1287_v15
			_1287_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1266_v7, _this.F29())
			(globalState).F4 = (func() _dafny.Int {
				if !((func() bool {
					if p1 {
						return p1
					}
					return (_this).F28()
				})()) {
					return _1266_v7
				}
				return (_this.F29()).Minus((_1287_v15).Cardinality())
			})()
			(globalState).F5 = _dafny.IntOfInt64(-700)
			var _1288_v16 D0
			_ = _1288_v16
			_1288_v16 = Companion_D0_.Create_DC3_(_this.F29(), p0, !(p1))
			var _1289_v17 _dafny.Sequence
			_ = _1289_v17
			_1289_v17 = _dafny.SeqOf(p1, p1)
			var _1290_v18 _dafny.Map
			_ = _1290_v18
			_1290_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1288_v16, _dafny.IntOfUint32((_1289_v17).Cardinality()))
			var _1291_v19 *C3
			_ = _1291_v19
			var _nw186 *C3 = New_C3_()
			_ = _nw186
			_nw186.Ctor__(_this.F29(), (func() _dafny.Map {
				if p1 {
					return _1290_v18
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1288_v16, (_dafny.Zero).Minus(_1266_v7))
			})(), ((_dafny.Zero).Minus(p2)).Plus(p2), (_this).F27(), (_this).F28())
			_1291_v19 = _nw186
			var _1292_v20 _dafny.Array
			_ = _1292_v20
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_28
			var _nw187 _dafny.Array
			_ = _nw187
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw187 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = (func(_1293_p0 bool, _1294_p1 bool) func(_dafny.Int) _dafny.Int {
					return func(_1295_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1295_i1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1293_p0, _1294_p1)).Cardinality())
					}
				})(p0, p1)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw187 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw187).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw187).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1292_v20 = _nw187
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1292_v20), 0))
			_ = _index195
			(_1292_v20).ArraySet1((_this.F29()).Minus((_1291_v19).F34()), (_index195).Int())
			var _1296_v21 _dafny.Sequence
			_ = _1296_v21
			_1296_v21 = _dafny.SeqOf(_1292_v20)
			var _1297_v22 _dafny.Map
			_ = _1297_v22
			_1297_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F27())
			var _1298_v23 *C5
			_ = _1298_v23
			var _nw188 *C5 = New_C5_()
			_ = _nw188
			_nw188.Ctor__(_1296_v21, _this.F29(), (func() _dafny.Array {
				if (_1297_v22).Contains(true) {
					return (_1297_v22).Get(true).(_dafny.Array)
				}
				return (_this).F27()
			})(), (_this).F28())
			_1298_v23 = _nw188
			var _1299_v24 _dafny.MultiSet
			_ = _1299_v24
			_1299_v24 = _dafny.MultiSetOf((_1296_v21).Select((Companion_Default___.SafeIndex((_1291_v19).F34(), _dafny.IntOfUint32((_1296_v21).Cardinality()))).Uint32()).(_dafny.Array))
			var _1300_v25 _dafny.Map
			_ = _1300_v25
			_1300_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1298_v23, (_1299_v24).Cardinality())
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1292_v20), 0))
			_ = _index196
			(_1292_v20).ArraySet1((Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1300_v25).Contains(_1298_v23) {
					return (_1300_v25).Get(_1298_v23).(_dafny.Int)
				}
				return _1266_v7
			})()), (_1291_v19).F34())).Minus(_1266_v7), (_index196).Int())
		} else if _source19.Is_DC2() {
			var _1301___mcc_h5 _dafny.Set = _source19.Get_().(D0_DC2).Cf5
			_ = _1301___mcc_h5
			var _1302___mcc_h6 _dafny.Sequence = _source19.Get_().(D0_DC2).Cf6
			_ = _1302___mcc_h6
			var _1303_cf6 _dafny.Sequence = _1302___mcc_h6
			_ = _1303_cf6
			var _1304_cf5 _dafny.Set = _1301___mcc_h5
			_ = _1304_cf5
			var _1305_v26 _dafny.Map
			_ = _1305_v26
			var _1306_v27 _dafny.Int
			_ = _1306_v27
			var _out25 _dafny.Map
			_ = _out25
			var _out26 _dafny.Int
			_ = _out26
			_out25, _out26 = Companion_Default___.M0(globalState)
			_1305_v26 = _out25
			_1306_v27 = _out26
			var _1307_v28 _dafny.MultiSet
			_ = _1307_v28
			_1307_v28 = _dafny.MultiSetOf(Companion_Default___.Fm46(p0, globalState))
			var _rhs211 _dafny.Sequence = _1268_v9
			_ = _rhs211
			var _rhs212 _dafny.MultiSet = ((_1307_v28).Union(_1307_v28)).Difference(_1307_v28)
			_ = _rhs212
			var _rhs213 bool = p0
			_ = _rhs213
			var _lhs188 *GlobalState = globalState
			_ = _lhs188
			_1303_cf6 = _rhs211
			_1307_v28 = _rhs212
			_lhs188.F6 = _rhs213
			var _1308_v29 _dafny.Array
			_ = _1308_v29
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw189
			_1308_v29 = _nw189
			var _1309_v30 _dafny.MultiSet
			_ = _1309_v30
			_1309_v30 = _dafny.MultiSetOf(_1266_v7)
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1308_v29), 0))
			_ = _index197
			(_1308_v29).ArraySet1(Companion_Default___.Fm49(_1309_v30, globalState), (_index197).Int())
			var _1310_v31 _dafny.Sequence
			_ = _1310_v31
			_1310_v31 = _dafny.SeqOf((_this).F28())
			var _1311_v32 _dafny.Sequence
			_ = _1311_v32
			_1311_v32 = _dafny.SeqOf(_1310_v31)
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1308_v29), 0))
			_ = _index198
			(_1308_v29).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1311_v32, _dafny.Companion_Sequence_.Concatenate(_1311_v32, _1311_v32)), (_index198).Int())
			(globalState).F6 = p0
		} else if _source19.Is_DC3() {
			var _1312___mcc_h7 _dafny.Int = _source19.Get_().(D0_DC3).Cf7
			_ = _1312___mcc_h7
			var _1313___mcc_h8 bool = _source19.Get_().(D0_DC3).Cf8
			_ = _1313___mcc_h8
			var _1314___mcc_h9 bool = _source19.Get_().(D0_DC3).Cf9
			_ = _1314___mcc_h9
			var _1315_cf9 bool = _1314___mcc_h9
			_ = _1315_cf9
			var _1316_cf8 bool = _1313___mcc_h8
			_ = _1316_cf8
			var _1317_cf7 _dafny.Int = _1312___mcc_h7
			_ = _1317_cf7
			(globalState).F13 = (Companion_Default___.Fm36(globalState)).Plus(_1266_v7)
			var _1318_v33 D3
			_ = _1318_v33
			_1318_v33 = Companion_D3_.Create_DC10_(p1, _1315_cf9)
			var _1319_v34 _dafny.Set
			_ = _1319_v34
			_1319_v34 = _dafny.SetOf(_1318_v33)
			var _1320_v35 _dafny.MultiSet
			_ = _1320_v35
			_1320_v35 = _dafny.MultiSetOf((_this).F28(), !(_1316_cf8), p1)
			var _1321_v36 _dafny.Sequence
			_ = _1321_v36
			_1321_v36 = _dafny.SeqOf((func() _dafny.Int {
				if (_1320_v35).Contains(false) {
					return (_1320_v35).Multiplicity(false)
				}
				return _this.F29()
			})(), _1266_v7)
			var _rhs214 _dafny.Int = _1266_v7
			_ = _rhs214
			var _rhs215 _dafny.Set = (func() _dafny.Set {
				if true {
					return _dafny.SetOf(_1318_v33)
				}
				return _dafny.SetOf(Companion_D3_.Create_DC10_(_1315_cf9, (_this).F28()))
			})()
			_ = _rhs215
			var _rhs216 _dafny.Sequence = _1321_v36
			_ = _rhs216
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			_lhs189.F21 = _rhs214
			_1319_v34 = _rhs215
			_1321_v36 = _rhs216
			(_this).F29_set_(Companion_Default___.Fm36(globalState))
			var _1322_v37 *C4
			_ = _1322_v37
			var _nw190 *C4 = New_C4_()
			_ = _nw190
			_nw190.Ctor__()
			_1322_v37 = _nw190
		} else {
			var _1323___mcc_h10 D0 = _source19.Get_().(D0_DC4).Cf10
			_ = _1323___mcc_h10
			var _1324_cf10 D0 = _1323___mcc_h10
			_ = _1324_cf10
			var _1325_v38 _dafny.Map
			_ = _1325_v38
			_1325_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
			var _1326_v39 *C0
			_ = _1326_v39
			var _nw191 *C0 = New_C0_()
			_ = _nw191
			_nw191.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1268_v9, _1268_v9), (_dafny.Zero).Minus((func() _dafny.Int {
				if p1 {
					return (_1325_v38).Cardinality()
				}
				return p2
			})()), (_this).F27(), (_this).F28())
			_1326_v39 = _nw191
			var _1327_v40 _dafny.Set
			_ = _1327_v40
			_1327_v40 = _dafny.SetOf(p1, Companion_Default___.Fm40(globalState))
			var _1328_v41 _dafny.Sequence
			_ = _1328_v41
			_1328_v41 = _dafny.SeqOf((_this).F28(), p0)
			var _1329_v42 _dafny.Array
			_ = _1329_v42
			var _nwElement0_35 _dafny.Int = (_dafny.Zero).Minus(((_1327_v40).Intersection(_1327_v40)).Cardinality())
			_ = _nwElement0_35
			var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(7))
			_ = _nw192
			(_nw192).ArraySet1(_nwElement0_35, 0)
			(_nw192).ArraySet1((_dafny.Zero).Minus(_1266_v7), 1)
			(_nw192).ArraySet1((_dafny.Zero).Minus(p2), 2)
			(_nw192).ArraySet1(_this.F29(), 3)
			(_nw192).ArraySet1(_this.F29(), 4)
			(_nw192).ArraySet1(_dafny.IntOfUint32((_1328_v41).Cardinality()), 5)
			(_nw192).ArraySet1(p2, 6)
			_1329_v42 = _nw192
			var _1330_v43 _dafny.Sequence
			_ = _1330_v43
			_1330_v43 = _dafny.SeqOf(_1329_v42)
			var _1331_v44 _dafny.MultiSet
			_ = _1331_v44
			_1331_v44 = _dafny.MultiSetOf(_1329_v42, (_1330_v43).Select((Companion_Default___.SafeIndex(_1266_v7, _dafny.IntOfUint32((_1330_v43).Cardinality()))).Uint32()).(_dafny.Array), _1329_v42)
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1329_v42), 0))
			_ = _index199
			(_1329_v42).ArraySet1((_1331_v44).Cardinality(), (_index199).Int())
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1329_v42), 0))
			_ = _index200
			(_1329_v42).ArraySet1((_this).Fm5(globalState), (_index200).Int())
			var _1332_v45 _dafny.Map
			_ = _1332_v45
			_1332_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1260_v1, _1330_v43)
			var _1333_v46 *C5
			_ = _1333_v46
			var _nw193 *C5 = New_C5_()
			_ = _nw193
			_nw193.Ctor__((func() _dafny.Sequence {
				if (_1332_v45).Contains(_dafny.CodePoint('q')) {
					return (_1332_v45).Get(_dafny.CodePoint('q')).(_dafny.Sequence)
				}
				return _1330_v43
			})(), _1266_v7, (_this).F27(), (_1328_v41).Select((Companion_Default___.SafeIndex((_1329_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1329_v42), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1328_v41).Cardinality()))).Uint32()).(bool))
			_1333_v46 = _nw193
			var _1334_v47 _dafny.Sequence
			_ = _1334_v47
			_1334_v47 = _dafny.SeqOf(_1266_v7, (_1326_v39).Fm15(globalState), _dafny.IntOfInt64(406), (_dafny.Zero).Minus(_dafny.IntOfInt64(-368)), Companion_Default___.Fm36(globalState))
			var _rhs217 _dafny.Int = ((_1334_v47).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1334_v47).Cardinality()))).Uint32()).(_dafny.Int)).Times(Companion_Default___.SafeModuloInt(_1266_v7, Companion_Default___.Fm36(globalState)))
			_ = _rhs217
			var _rhs218 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1328_v41, _1328_v41)
			_ = _rhs218
			var _rhs219 _dafny.Sequence = (_1333_v46).F33()
			_ = _rhs219
			var _rhs220 bool = !(!(p0)) || (p1)
			_ = _rhs220
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			var _lhs191 *GlobalState = globalState
			_ = _lhs191
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			_lhs190.F13 = _rhs217
			_lhs191.F15 = _rhs218
			_1330_v43 = _rhs219
			_lhs192.F6 = _rhs220
		}
		var _1335_v48 _dafny.Map
		_ = _1335_v48
		_1335_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F29()).Minus(_1266_v7), p0)
		_1335_v48 = (_1335_v48).Update(_dafny.IntOfUint32((_1268_v9).Cardinality()), (_this).F28())
		var _1336_v49 _dafny.Sequence
		_ = _1336_v49
		_1336_v49 = _dafny.SeqOf(p1)
		var _1337_v50 _dafny.MultiSet
		_ = _1337_v50
		_1337_v50 = _dafny.MultiSetOf(p0, !(p0), (_this).F28(), false, true)
		var _1338_v51 _dafny.Sequence
		_ = _1338_v51
		_1338_v51 = _dafny.SeqOf(_1337_v50)
		var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen(((_this).F27()), 0))
		_ = _index201
		((_this).F27()).ArraySet1((_1336_v49).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_1338_v51)).Cardinality(), _dafny.IntOfUint32((_1336_v49).Cardinality()))).Uint32()).(bool), (_index201).Int())
		var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen(((_this).F27()), 0))
		_ = _index202
		((_this).F27()).ArraySet1((p2).Cmp(_1266_v7) == 0, (_index202).Int())
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f29 _dafny.Int
	_f27 _dafny.Array
	_f28 bool
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f29 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C7{}
var _ T1 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F29() _dafny.Int {
	return _this._f29
}
func (_this *C7) F29_set_(value _dafny.Int) {
	_this._f29 = value
}
func (_this *C7) F27() _dafny.Array {
	return _this._f27
}
func (_this *C7) F28() bool {
	return _this._f28
}
func (_this *C7) Ctor__(f27 _dafny.Array, f28 bool, f29 _dafny.Int) {
	{
		(_this)._f27 = f27
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C7) M5(globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1339_v0 _dafny.CodePoint
		_ = _1339_v0
		_1339_v0 = _dafny.CodePoint('e')
		var _1340_v1 _dafny.Map
		_ = _1340_v1
		_1340_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_dafny.Zero).Minus(_this.F29()))
		var _1341_v2 _dafny.Sequence
		_ = _1341_v2
		_1341_v2 = _dafny.SeqOf((_dafny.Zero).Minus((func() _dafny.Int {
			if (_1340_v1).Contains(true) {
				return (_1340_v1).Get(true).(_dafny.Int)
			}
			return _this.F29()
		})()), _this.F29())
		var _rhs221 _dafny.Int = _this.F29()
		_ = _rhs221
		var _rhs222 _dafny.CodePoint = _1339_v0
		_ = _rhs222
		var _rhs223 bool = !(!((_this).F28()))
		_ = _rhs223
		var _rhs224 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F29(), (_1341_v2).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1341_v2).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _rhs224
		var _lhs193 *GlobalState = globalState
		_ = _lhs193
		var _lhs194 *GlobalState = globalState
		_ = _lhs194
		r0 = _rhs221
		_lhs193.F25 = _rhs222
		_lhs194.F17 = _rhs223
		r0 = _rhs224
		var _1342_v3 _dafny.Sequence
		_ = _1342_v3
		_1342_v3 = _dafny.UnicodeSeqOfUtf8Bytes("yaenitp")
		(globalState).F17 = !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("rhmgaits"), _dafny.Companion_Sequence_.Concatenate(_1342_v3, _1342_v3))
		(globalState).F6 = (func() bool {
			if (_this).F28() {
				return (_this).F28()
			}
			return (func() bool {
				if (_this).F28() {
					return (_this).F28()
				}
				return (_this).F28()
			})()
		})()
		var _1343_v4 _dafny.Sequence
		_ = _1343_v4
		_1343_v4 = _1341_v2
		var _source20 _dafny.Sequence = _1343_v4
		_ = _source20
		var _1344___mcc_h0 _dafny.Sequence = _source20
		_ = _1344___mcc_h0
		var _1345_cf15 _dafny.Sequence = _1344___mcc_h0
		_ = _1345_cf15
		(globalState).F25 = (_1342_v3).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1342_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
		_1339_v0 = _1339_v0
		var _1346_v5 D1
		_ = _1346_v5
		_1346_v5 = Companion_D1_.Create_DC6_(_this.F29())
		var _1347_v6 _dafny.MultiSet
		_ = _1347_v6
		_1347_v6 = _dafny.MultiSetOf(_1346_v5, _1346_v5)
		var _1348_v7 _dafny.Sequence
		_ = _1348_v7
		_1348_v7 = _dafny.SeqOf(_1347_v6)
		var _1349_v8 _dafny.Sequence
		_ = _1349_v8
		_1349_v8 = _dafny.SeqOf(true)
		var _1350_v9 _dafny.Map
		_ = _1350_v9
		_1350_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _1339_v0)
		var _1351_v10 _dafny.Map
		_ = _1351_v10
		_1351_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _1350_v9)
		var _1352_v11 _dafny.Sequence
		_ = _1352_v11
		_1352_v11 = _dafny.SeqOf((func() _dafny.Map {
			if (_1351_v10).Contains(false) {
				return (_1351_v10).Get(false).(_dafny.Map)
			}
			return _1350_v9
		})())
		var _1353_v12 _dafny.Map
		_ = _1353_v12
		_1353_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _dafny.MultiSetOf(_1346_v5, _1346_v5))
		var _1354_v13 _dafny.Map
		_ = _1354_v13
		_1354_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F28())
		var _1355_v14 _dafny.Array
		_ = _1355_v14
		var _nwElement0_36 _dafny.Int = _this.F29()
		_ = _nwElement0_36
		var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(29))
		_ = _nw194
		(_nw194).ArraySet1(_nwElement0_36, 0)
		(_nw194).ArraySet1(_this.F29(), 1)
		(_nw194).ArraySet1(Companion_Default___.Fm4((_1348_v7).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1348_v7).Cardinality()))).Uint32()).(_dafny.MultiSet), globalState), 2)
		(_nw194).ArraySet1(_this.F29(), 3)
		(_nw194).ArraySet1(_dafny.IntOfUint32((_1342_v3).Cardinality()), 4)
		(_nw194).ArraySet1(_this.F29(), 5)
		(_nw194).ArraySet1(_this.F29(), 6)
		(_nw194).ArraySet1(_dafny.IntOfUint32((_1349_v8).Cardinality()), 7)
		(_nw194).ArraySet1(_this.F29(), 8)
		(_nw194).ArraySet1(_this.F29(), 9)
		(_nw194).ArraySet1((func() _dafny.Int {
			if (_1340_v1).Contains(true) {
				return (_1340_v1).Get(true).(_dafny.Int)
			}
			return _this.F29()
		})(), 10)
		(_nw194).ArraySet1(_this.F29(), 11)
		(_nw194).ArraySet1(_dafny.IntOfInt64(220), 12)
		(_nw194).ArraySet1(_this.F29(), 13)
		(_nw194).ArraySet1(((_1352_v11).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1352_v11).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 14)
		(_nw194).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F29())), 15)
		(_nw194).ArraySet1(_this.F29(), 16)
		(_nw194).ArraySet1(_this.F29(), 17)
		(_nw194).ArraySet1(_dafny.IntOfInt64(824), 18)
		(_nw194).ArraySet1(_this.F29(), 19)
		(_nw194).ArraySet1(_this.F29(), 20)
		(_nw194).ArraySet1(_dafny.IntOfInt64(20), 21)
		(_nw194).ArraySet1(_dafny.IntOfInt64(360), 22)
		(_nw194).ArraySet1(Companion_Default___.Fm4((func() _dafny.MultiSet {
			if (_1353_v12).Contains(_this.F29()) {
				return (_1353_v12).Get(_this.F29()).(_dafny.MultiSet)
			}
			return _1347_v6
		})(), globalState), 23)
		(_nw194).ArraySet1(_this.F29(), 24)
		(_nw194).ArraySet1(_dafny.IntOfUint32((_1345_cf15).Cardinality()), 25)
		(_nw194).ArraySet1(_dafny.IntOfInt64(660), 26)
		(_nw194).ArraySet1(_dafny.IntOfUint32((_1342_v3).Cardinality()), 27)
		(_nw194).ArraySet1((_dafny.Zero).Minus((_1354_v13).Cardinality()), 28)
		_1355_v14 = _nw194
		var _1356_v15 _dafny.Map
		_ = _1356_v15
		_1356_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_this).F28() {
				return true
			}
			return (_this).F28()
		})(), _1355_v14)
		_1356_v15 = (_1356_v15).Update(!((_this).F28()), _1355_v14)
		var _1357_v16 _dafny.Sequence
		_ = _1357_v16
		_1357_v16 = _dafny.SeqOf(_1355_v14)
		var _1358_v17 T1
		_ = _1358_v17
		var _nw195 *C5 = New_C5_()
		_ = _nw195
		_nw195.Ctor__(_1357_v16, _this.F29(), (_this).F27(), !((_this).F28()))
		_1358_v17 = _nw195
		var _1359_v18 _dafny.Map
		_ = _1359_v18
		_1359_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1358_v17, _this.F29())
		var _1360_v19 _dafny.Map
		_ = _1360_v19
		_1360_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1359_v18).Update(_1358_v17, _dafny.IntOfUint32((_1342_v3).Cardinality())), (_this).F28())
		var _1361_v20 D0
		_ = _1361_v20
		_1361_v20 = Companion_D0_.Create_DC0_(_1360_v19, (_this).F28(), _1358_v17.F29(), _1339_v0, (_1358_v17).F28())
		var _1362_v21 _dafny.Map
		_ = _1362_v21
		_1362_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_1358_v17).F28())
		var _1363_v22 _dafny.Array
		_ = _1363_v22
		var _nwElement0_37 bool = (_dafny.IntOfInt64(902)).Cmp(Companion_Default___.Fm4(_1347_v6, globalState)) > 0
		_ = _nwElement0_37
		var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(24))
		_ = _nw196
		(_nw196).ArraySet1(_nwElement0_37, 0)
		(_nw196).ArraySet1(!((_this).F28()), 1)
		(_nw196).ArraySet1((_this).F28(), 2)
		(_nw196).ArraySet1(!((_this).F28()), 3)
		(_nw196).ArraySet1(!((_this).F28()), 4)
		(_nw196).ArraySet1((_1361_v20).Dtor_cf1(), 5)
		(_nw196).ArraySet1((_this).F28(), 6)
		(_nw196).ArraySet1((_1358_v17).F28(), 7)
		(_nw196).ArraySet1(!((_1358_v17).F28()), 8)
		(_nw196).ArraySet1((_1358_v17).F28(), 9)
		(_nw196).ArraySet1(!(Companion_Default___.Fm10(_1339_v0, globalState)), 10)
		(_nw196).ArraySet1((_1358_v17).F28(), 11)
		(_nw196).ArraySet1(((_this).F28()) == ((func() bool {
			if (_1362_v21).Contains((_this).F28()) {
				return (_1362_v21).Get((_this).F28()).(bool)
			}
			return (_1358_v17).F28()
		})()), 12)
		(_nw196).ArraySet1((_1358_v17).F28(), 13)
		(_nw196).ArraySet1((_1358_v17).F28(), 14)
		(_nw196).ArraySet1((_1358_v17).F28(), 15)
		(_nw196).ArraySet1(Companion_Default___.Fm22((_1358_v17).F28(), _1358_v17.F29(), globalState), 16)
		(_nw196).ArraySet1((func() bool {
			if true {
				return (_this).F28()
			}
			return true
		})(), 17)
		(_nw196).ArraySet1((_this).F28(), 18)
		(_nw196).ArraySet1(!((_this).F28()), 19)
		(_nw196).ArraySet1((_this).F28(), 20)
		(_nw196).ArraySet1((_1358_v17).F28(), 21)
		(_nw196).ArraySet1((_1358_v17).F28(), 22)
		(_nw196).ArraySet1((_this).F28(), 23)
		_1363_v22 = _nw196
		_1363_v22 = (_1358_v17).F27()
		var _1364_v23 _dafny.Map
		_ = _1364_v23
		_1364_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC26_(true), (_this).F28())
		var _1365_v24 *C1
		_ = _1365_v24
		var _nw197 *C1 = New_C1_()
		_ = _nw197
		_nw197.Ctor__((func() bool {
			if (_1364_v23).Contains(Companion_D10_.Create_DC26_((_this).F28())) {
				return (_1364_v23).Get(Companion_D10_.Create_DC26_((_this).F28())).(bool)
			}
			return (_this).F28()
		})(), (_this).F27(), (_dafny.IntOfInt64(-978)).Cmp(_this.F29()) > 0)
		_1365_v24 = _nw197
		var _hi6 _dafny.Int = _this.F29()
		_ = _hi6
		for _1366_i0 := _this.F29(); _1366_i0.Cmp(_hi6) < 0; _1366_i0 = _1366_i0.Plus(_dafny.One) {
			r0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(894), _this.F29())).Minus(Companion_Default___.Fm36(globalState))
			var _1367_v25 _dafny.Array
			_ = _1367_v25
			var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
			_ = _nw198
			_1367_v25 = _nw198
			var _source21 D11 = Companion_D11_.Create_DC27_(_1367_v25)
			_ = _source21
			if _source21.Is_DC28() {
				_1342_v3 = _1342_v3
				_1342_v3 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bdswexrbt"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("am"), _1342_v3))
				var _1368_v26 _dafny.Set
				_ = _1368_v26
				_1368_v26 = _dafny.SetOf((_1365_v24).F39(), false, (_1365_v24).F39())
				var _1369_v27 D8
				_ = _1369_v27
				_1369_v27 = Companion_D8_.Create_DC18_((_1368_v26).Cardinality())
				_1369_v27 = _1369_v27
				var _1370_v28 _dafny.MultiSet
				_ = _1370_v28
				_1370_v28 = _dafny.MultiSetOf((_this).F28())
				var _1371_v29 _dafny.MultiSet
				_ = _1371_v29
				_1371_v29 = _dafny.MultiSetOf((_1370_v28).Difference(_1370_v28), (func() _dafny.MultiSet {
					if (_1365_v24).F39() {
						return _1370_v28
					}
					return _1370_v28
				})(), _1370_v28)
				var _rhs225 _dafny.Int = (_dafny.Zero).Minus(_1366_i0)
				_ = _rhs225
				var _rhs226 _dafny.Int = (_1371_v29).Cardinality()
				_ = _rhs226
				var _lhs195 *GlobalState = globalState
				_ = _lhs195
				var _lhs196 *GlobalState = globalState
				_ = _lhs196
				_lhs195.F5 = _rhs225
				_lhs196.F13 = _rhs226
			} else {
				var _1372___mcc_h1 _dafny.Array = _source21.Get_().(D11_DC27).Cf42
				_ = _1372___mcc_h1
				var _1373_cf42 _dafny.Array = _1372___mcc_h1
				_ = _1373_cf42
				(globalState).F25 = _dafny.CodePoint('m')
				var _1374_v30 _dafny.MultiSet
				_ = _1374_v30
				_1374_v30 = _dafny.MultiSetOf(_this.F29())
				var _1375_v31 _dafny.MultiSet
				_ = _1375_v31
				_1375_v31 = _dafny.MultiSetOf(true)
				var _1376_v32 _dafny.Map
				_ = _1376_v32
				_1376_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _1375_v31)
				var _1377_v33 _dafny.Sequence
				_ = _1377_v33
				_1377_v33 = _dafny.SeqOf((_1375_v31).Update((_1365_v24).F39(), Companion_Default___.Abs(_this.F29())), (func() _dafny.MultiSet {
					if (_1376_v32).Contains(_1366_i0) {
						return (_1376_v32).Get(_1366_i0).(_dafny.MultiSet)
					}
					return _1375_v31
				})(), _1375_v31)
				(globalState).F6 = ((func() _dafny.Int {
					if (_1374_v30).Contains(_this.F29()) {
						return (_1374_v30).Multiplicity(_this.F29())
					}
					return ((_1377_v33).Select((Companion_Default___.SafeIndex(_1366_i0, _dafny.IntOfUint32((_1377_v33).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()
				})()).Cmp(_1366_i0) <= 0
				var _1378_v34 _dafny.Array
				_ = _1378_v34
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_29
				var _nw199 _dafny.Array
				_ = _nw199
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw199 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Map = (func(_1379_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
						return func(_1380_i1 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1379_v3).Cardinality()), (_this).F28())
						}
					})(_1342_v3)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw199 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw199).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw199).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1378_v34 = _nw199
				var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw200
				_1378_v34 = _nw200
				var _1381_v35 _dafny.Array
				_ = _1381_v35
				var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw201
				_1381_v35 = _nw201
				var _1382_v36 _dafny.Sequence
				_ = _1382_v36
				_1382_v36 = _dafny.SeqOf(_1381_v35, _1381_v35, _1381_v35)
				var _1383_v37 _dafny.MultiSet
				_ = _1383_v37
				_1383_v37 = _dafny.MultiSetOf(_1381_v35, _1381_v35, (_1382_v36).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1382_v36).Cardinality()))).Uint32()).(_dafny.Array), _1381_v35, _1381_v35)
				(globalState).F17 = (_1383_v37).IsProperSubsetOf((_1383_v37).Update(_1381_v35, Companion_Default___.Abs(_this.F29())))
			}
			var _1384_v38 _dafny.Map
			_ = _1384_v38
			_1384_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm50((_this).F28(), _this.F29(), globalState), (_dafny.Zero).Minus(_1366_i0))
			_1384_v38 = (_1384_v38).Update(_1340_v1, _dafny.IntOfInt64(370))
			(globalState).F13 = _dafny.IntOfUint32((_1341_v2).Cardinality())
		}
		r0 = (_dafny.Zero).Minus(_this.F29())
		r1 = (_this.F29()).Times(_this.F29())
		return r0, r1
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f27 _dafny.Array
	_f28 bool
	_f32 _dafny.CodePoint
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
	_this._f32 = _dafny.CodePoint('D')
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F27() _dafny.Array {
	return _this._f27
}
func (_this *C8) F28() bool {
	return _this._f28
}
func (_this *C8) Ctor__(f32 _dafny.CodePoint, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f32 = f32
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C8) Fm3(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate((Companion_D1_.Create_DC5_(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(729))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_1385_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-103)
		}))))).Dtor_cf11(), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(323)), _dafny.SeqOf(_dafny.IntOfInt64(-823), (_dafny.MultiSetOf(_dafny.IntOfInt64(164))).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F28())).Cardinality())), (_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-361))).Cardinality()))))))
	}
}
func (_this *C8) M3(p0 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _1386_v0 T1
		_ = _1386_v0
		var _nw202 *C0 = New_C0_()
		_ = _nw202
		_nw202.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("yovobx"), _dafny.IntOfInt64(139), (_this).F27(), (_this).F28())
		_1386_v0 = _nw202
		var _1387_v1 _dafny.Map
		_ = _1387_v1
		_1387_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1386_v0, _1386_v0.F29())
		var _1388_v2 _dafny.Map
		_ = _1388_v2
		_1388_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1387_v1, (_1386_v0).F28())
		var _1389_v3 D0
		_ = _1389_v3
		_1389_v3 = Companion_D0_.Create_DC0_(_1388_v2, (_1386_v0).F28(), _dafny.IntOfInt64(130), (_this).F32(), (_1386_v0).F28())
		var _1390_v4 _dafny.Map
		_ = _1390_v4
		_1390_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_1389_v3).Dtor_cf1())
		var _1391_v5 D1
		_ = _1391_v5
		_1391_v5 = Companion_D1_.Create_DC6_(_1386_v0.F29())
		var _1392_v6 _dafny.Set
		_ = _1392_v6
		_1392_v6 = _dafny.SetOf(Companion_Default___.Fm4(_dafny.MultiSetOf(_1391_v5, Companion_D1_.Create_DC6_(p0), _1391_v5), globalState))
		var _1393_v7 _dafny.Sequence
		_ = _1393_v7
		_1393_v7 = _dafny.UnicodeSeqOfUtf8Bytes("upehw")
		var _1394_v8 D0
		_ = _1394_v8
		_1394_v8 = Companion_D0_.Create_DC2_(_1392_v6, (func() _dafny.Sequence {
			if (_1386_v0).F28() {
				return _1393_v7
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("ojwsghveo")
		})())
		var _1395_v9 _dafny.Sequence
		_ = _1395_v9
		_1395_v9 = _dafny.SeqOf((_1386_v0).F28(), (_this).F28(), (_1386_v0).F28())
		var _rhs227 _dafny.Map = _1390_v4
		_ = _rhs227
		var _rhs228 bool = (func() bool {
			if true {
				return (_this).F28()
			}
			return (_1395_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1395_v9).Cardinality()))).Uint32()).(bool)
		})()
		_ = _rhs228
		var _rhs229 D0 = _1394_v8
		_ = _rhs229
		_1390_v4 = _rhs227
		r1 = _rhs228
		_1394_v8 = _rhs229
		_1390_v4 = (_1390_v4).Update((_this).F28(), (_1386_v0).F28())
		var _1396_i0 _dafny.Int
		_ = _1396_i0
		_1396_i0 = _dafny.Zero
		{
			for ((_1386_v0).F28()) || ((_this).F28()) {
				{
					if (_1396_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1396_i0 = (_1396_i0).Plus(_dafny.One)
					var _1397_v10 _dafny.Array
					_ = _1397_v10
					var _len0_30 _dafny.Int = _dafny.IntOfInt64(14)
					_ = _len0_30
					var _nw203 _dafny.Array
					_ = _nw203
					if _len0_30.Cmp(_dafny.Zero) == 0 {
						_nw203 = _dafny.NewArray(_len0_30)
					} else {
						var _init30 func(_dafny.Int) _dafny.Int = func(_1398_i1 _dafny.Int) _dafny.Int {
							return (_1398_i1).Minus(_dafny.IntOfInt64(124))
						}
						_ = _init30
						var _element0_30 = _init30(_dafny.Zero)
						_ = _element0_30
						_nw203 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
						(_nw203).ArraySet1(_element0_30, 0)
						var _nativeLen0_30 = (_len0_30).Int()
						_ = _nativeLen0_30
						for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
							(_nw203).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
						}
					}
					_1397_v10 = _nw203
					_1397_v10 = _1397_v10
					(_1386_v0).F29_set_(_1386_v0.F29())
					var _1399_v11 *C4
					_ = _1399_v11
					var _nw204 *C4 = New_C4_()
					_ = _nw204
					_nw204.Ctor__()
					_1399_v11 = _nw204
					_1399_v11 = _1399_v11
					var _1400_v12 _dafny.Sequence
					_ = _1400_v12
					_1400_v12 = _dafny.SeqOf(p0)
					var _1401_v13 _dafny.Sequence
					_ = _1401_v13
					_1401_v13 = _1400_v12
					var _rhs230 _dafny.Set = _1392_v6
					_ = _rhs230
					var _rhs231 _dafny.Sequence = _1401_v13
					_ = _rhs231
					var _rhs232 _dafny.Int = _1386_v0.F29()
					_ = _rhs232
					var _rhs233 bool = (_this).F28()
					_ = _rhs233
					var _lhs197 *GlobalState = globalState
					_ = _lhs197
					_1392_v6 = _rhs230
					_1401_v13 = _rhs231
					_lhs197.F13 = _rhs232
					r1 = _rhs233
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1402_v14 _dafny.MultiSet
		_ = _1402_v14
		_1402_v14 = _dafny.MultiSetOf(_1386_v0.F29())
		var _1403_v16 _dafny.Map
		_ = _1403_v16
		_1403_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1386_v0).F28(), p0)
		var _1404_v17 _dafny.Map
		_ = _1404_v17
		_1404_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1403_v16, (_1386_v0).F28())
		var _1405_v18 _dafny.Set
		_ = _1405_v18
		_1405_v18 = _dafny.SetOf(Companion_Default___.Fm16(_1390_v4, globalState), ((_1402_v14).Cardinality()).Cmp((func() _dafny.Map {
			var _coll64 = _dafny.NewMapBuilder()
			_ = _coll64
			for _iter66 := _dafny.Iterate((_1404_v17).Keys().Elements()); ; {
				_compr_64, _ok66 := _iter66()
				if !_ok66 {
					break
				}
				var _1406_v15 _dafny.Map
				_1406_v15 = interface{}(_compr_64).(_dafny.Map)
				if (_1404_v17).Contains(_1406_v15) {
					_coll64.Add(_1406_v15, _1403_v16)
				}
			}
			return _coll64.ToMap()
		}()).Cardinality()) <= 0)
		(globalState).F13 = (_dafny.Zero).Minus((_1405_v18).Cardinality())
		var _rhs234 bool = (func() bool {
			if (_1386_v0).F28() {
				return !(!((_this).F28()))
			}
			return Companion_Default___.Fm40(globalState)
		})()
		_ = _rhs234
		var _rhs235 bool = (_this).F28()
		_ = _rhs235
		var _rhs236 _dafny.Int = _1386_v0.F29()
		_ = _rhs236
		var _lhs198 *GlobalState = globalState
		_ = _lhs198
		var _lhs199 *GlobalState = globalState
		_ = _lhs199
		r1 = _rhs234
		_lhs198.F17 = _rhs235
		_lhs199.F21 = _rhs236
		var _1407_i2 _dafny.Int
		_ = _1407_i2
		_1407_i2 = _dafny.Zero
		{
			for (_this).F28() {
				{
					if (_1407_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1407_i2 = (_1407_i2).Plus(_dafny.One)
					var _1408_v19 *C1
					_ = _1408_v19
					var _nw205 *C1 = New_C1_()
					_ = _nw205
					_nw205.Ctor__((_this).F28(), (_1386_v0).F27(), (_this).F28())
					_1408_v19 = _nw205
					_1405_v18 = _1405_v18
					var _1409_v20 _dafny.MultiSet
					_ = _1409_v20
					_1409_v20 = _dafny.MultiSetOf(_1391_v5)
					(_1386_v0).F29_set_(Companion_Default___.Fm4(_1409_v20, globalState))
					if (_1408_v19).F39() {
						var _1410_v21 _dafny.Map
						_ = _1410_v21
						_1410_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cdotqlbut")).Cardinality()), (_1386_v0).F27())
						var _1411_v22 _dafny.Array
						_ = _1411_v22
						_1411_v22 = (_1386_v0).F27()
						var _1412_v23 _dafny.Array
						_ = _1412_v23
						var _nwElement0_38 _dafny.Array = (_1386_v0).F27()
						_ = _nwElement0_38
						var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(26))
						_ = _nw206
						(_nw206).ArraySet1(_nwElement0_38, 0)
						(_nw206).ArraySet1((_1386_v0).F27(), 1)
						(_nw206).ArraySet1((_1386_v0).F27(), 2)
						(_nw206).ArraySet1((_1386_v0).F27(), 3)
						(_nw206).ArraySet1((_1386_v0).F27(), 4)
						(_nw206).ArraySet1((_1386_v0).F27(), 5)
						(_nw206).ArraySet1((_1386_v0).F27(), 6)
						(_nw206).ArraySet1((_1386_v0).F27(), 7)
						(_nw206).ArraySet1((_this).F27(), 8)
						(_nw206).ArraySet1((_1386_v0).F27(), 9)
						(_nw206).ArraySet1((_1386_v0).F27(), 10)
						(_nw206).ArraySet1((_1386_v0).F27(), 11)
						(_nw206).ArraySet1((_this).F27(), 12)
						(_nw206).ArraySet1((_1386_v0).F27(), 13)
						(_nw206).ArraySet1((_1386_v0).F27(), 14)
						(_nw206).ArraySet1((_1386_v0).F27(), 15)
						(_nw206).ArraySet1((_1386_v0).F27(), 16)
						(_nw206).ArraySet1((_1386_v0).F27(), 17)
						(_nw206).ArraySet1((_this).F27(), 18)
						(_nw206).ArraySet1((_1386_v0).F27(), 19)
						(_nw206).ArraySet1((func() _dafny.Array {
							if (_1410_v21).Contains(_1386_v0.F29()) {
								return (_1410_v21).Get(_1386_v0.F29()).(_dafny.Array)
							}
							return (_1386_v0).F27()
						})(), 20)
						(_nw206).ArraySet1((_1386_v0).F27(), 21)
						(_nw206).ArraySet1((_this).F27(), 22)
						(_nw206).ArraySet1((_1386_v0).F27(), 23)
						(_nw206).ArraySet1((_1411_v22), 24)
						(_nw206).ArraySet1((_this).F27(), 25)
						_1412_v23 = _nw206
						var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1412_v23), 0))
						_ = _index203
						(_1412_v23).ArraySet1((_1386_v0).F27(), (_index203).Int())
						var _1413_v24 *C4
						_ = _1413_v24
						var _nw207 *C4 = New_C4_()
						_ = _nw207
						_nw207.Ctor__()
						_1413_v24 = _nw207
						var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1412_v23), 0))
						_ = _index204
						var _rhs237 bool = (_1408_v19).F39()
						_ = _rhs237
						var _rhs238 bool = (_1386_v0).F28()
						_ = _rhs238
						var _rhs239 _dafny.Array = (func() _dafny.Array {
							if (func() bool {
								if (_1386_v0).F28() {
									return (_1408_v19).F39()
								}
								return (_1386_v0).F28()
							})() {
								return (_1386_v0).F27()
							}
							return (_1386_v0).F27()
						})()
						_ = _rhs239
						var _rhs240 *C4 = _1413_v24
						_ = _rhs240
						var _lhs200 *GlobalState = globalState
						_ = _lhs200
						var _lhs201 _dafny.Array = _1412_v23
						_ = _lhs201
						var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1412_v23), 0))
						_ = _lhs202
						r1 = _rhs237
						_lhs200.F6 = _rhs238
						(_lhs201).ArraySet1(_rhs239, (_lhs202).Int())
						_1413_v24 = _rhs240
						(globalState).F17 = (func() bool {
							if (_this).F28() {
								return (_1395_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1395_v9).Cardinality()))).Uint32()).(bool)
							}
							return (func() bool {
								if (_1390_v4).Contains((_1386_v0).F28()) {
									return (_1390_v4).Get((_1386_v0).F28()).(bool)
								}
								return !((_1408_v19).F39())
							})()
						})()
						_1393_v7 = Companion_Default___.Fm18((func() bool {
							if (_1408_v19).F39() {
								return (_1408_v19).F39()
							}
							return (_1408_v19).F39()
						})(), (_dafny.Zero).Minus(_1386_v0.F29()), (_this).F32(), globalState)
						var _1414_v25 _dafny.Array
						_ = _1414_v25
						var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
						_ = _nw208
						_1414_v25 = _nw208
						_1414_v25 = _1414_v25
						r1 = (_1386_v0).F28()
					} else {
						var _1415_v26 _dafny.Array
						_ = _1415_v26
						var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
						_ = _nw209
						_1415_v26 = _nw209
						var _1416_v27 _dafny.Sequence
						_ = _1416_v27
						_1416_v27 = _dafny.SeqOf(_1415_v26, _1415_v26)
						var _1417_v28 D8
						_ = _1417_v28
						_1417_v28 = Companion_D8_.Create_DC18_(_dafny.IntOfInt64(-681))
						var _1418_v29 _dafny.Map
						_ = _1418_v29
						_1418_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1417_v28, (_1386_v0).F28())
						var _1419_v30 *C5
						_ = _1419_v30
						var _nw210 *C5 = New_C5_()
						_ = _nw210
						_nw210.Ctor__(_1416_v27, (_1418_v29).Cardinality(), (_this).F27(), (_this).F28())
						_1419_v30 = _nw210
						_1402_v14 = (_1402_v14).Union((_1402_v14).Union(_dafny.MultiSetOf(p0)))
						(globalState).F5 = Companion_Default___.SafeModuloInt(_1386_v0.F29(), (_1419_v30).Fm7((_this).F32(), globalState))
						(globalState).F15 = (_1386_v0).F28()
						var _1420_v31 _dafny.Array
						_ = _1420_v31
						var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
						_ = _nw211
						_1420_v31 = _nw211
						var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_1420_v31), 0))
						_ = _index205
						(_1420_v31).ArraySet1CodePoint((_this).F32(), (_index205).Int())
						var _1421_v32 _dafny.Map
						_ = _1421_v32
						_1421_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1386_v0.F29(), (_1386_v0).F28())
						var _1422_v33 D3
						_ = _1422_v33
						_1422_v33 = Companion_D3_.Create_DC10_((_1408_v19).F39(), (_this).F28())
						var _1423_v34 D5
						_ = _1423_v34
						_1423_v34 = Companion_D5_.Create_DC13_((_1386_v0).F28(), _1415_v26, (Companion_D7_.Create_DC16_((_1408_v19).F39(), _1421_v32, (_this).F32(), (_1422_v33).Dtor_cf17())).Dtor_cf29())
						var _pat_let_tv28 = _1386_v0
						_ = _pat_let_tv28
						var _1424_v35 _dafny.Array
						_ = _1424_v35
						var _nwElement0_39 D5 = func(_pat_let21_0 D5) D5 {
							return func(_1425_dt__update__tmp_h0 D5) D5 {
								return func(_pat_let22_0 bool) D5 {
									return func(_1426_dt__update_hcf21_h0 bool) D5 {
										return Companion_D5_.Create_DC13_(_1426_dt__update_hcf21_h0, (_1425_dt__update__tmp_h0).Dtor_cf22(), (_1425_dt__update__tmp_h0).Dtor_cf23())
									}(_pat_let22_0)
								}((_pat_let_tv28).F28())
							}(_pat_let21_0)
						}(_1423_v34)
						_ = _nwElement0_39
						var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(2))
						_ = _nw212
						(_nw212).ArraySet1(_nwElement0_39, 0)
						(_nw212).ArraySet1(_1423_v34, 1)
						_1424_v35 = _nw212
						var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_1424_v35), 0))
						_ = _index206
						(_1424_v35).ArraySet1(_1423_v34, (_index206).Int())
						var _1427_v36 _dafny.Sequence
						_ = _1427_v36
						_1427_v36 = _dafny.SeqOf(_1393_v7, _1393_v7, _1393_v7, _1393_v7)
						var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_1420_v31), 0))
						_ = _index207
						var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_1424_v35), 0))
						_ = _index208
						var _rhs241 _dafny.CodePoint = (_this).F32()
						_ = _rhs241
						var _rhs242 D5 = _1423_v34
						_ = _rhs242
						var _rhs243 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1427_v36).Select((Companion_Default___.SafeIndex((_1421_v32).Cardinality(), _dafny.IntOfUint32((_1427_v36).Cardinality()))).Uint32()).(_dafny.Sequence), _1393_v7)
						_ = _rhs243
						var _rhs244 _dafny.Array = (_1423_v34).Dtor_cf22()
						_ = _rhs244
						var _rhs245 bool = (_1405_v18).IsDisjointFrom(_1405_v18)
						_ = _rhs245
						var _lhs203 _dafny.Array = _1420_v31
						_ = _lhs203
						var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_1420_v31), 0))
						_ = _lhs204
						var _lhs205 _dafny.Array = _1424_v35
						_ = _lhs205
						var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_1424_v35), 0))
						_ = _lhs206
						var _lhs207 *GlobalState = globalState
						_ = _lhs207
						(_lhs203).ArraySet1CodePoint(_rhs241, (_lhs204).Int())
						(_lhs205).ArraySet1(_rhs242, (_lhs206).Int())
						_1393_v7 = _rhs243
						_1415_v26 = _rhs244
						_lhs207.F15 = _rhs245
					}
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1428_v37 _dafny.Array
		_ = _1428_v37
		var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw213
		_1428_v37 = _nw213
		var _1429_v38 _dafny.Sequence
		_ = _1429_v38
		_1429_v38 = _dafny.SeqOf(_1428_v37)
		r0 = _1429_v38
		r1 = (_dafny.IntOfInt64(225)).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1393_v7).Cardinality()), (_dafny.Zero).Minus(p0))) == 0
		return r0, r1
	}
}
func (_this *C8) M4(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _1430_v0 _dafny.Set
		_ = _1430_v0
		_1430_v0 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("krl")).Cardinality()))
		var _1431_v1 D0
		_ = _1431_v1
		_1431_v1 = Companion_D0_.Create_DC2_(_1430_v0, Companion_Default___.Fm32(p0, (_this).F28(), globalState))
		var _1432_v2 _dafny.Map
		_ = _1432_v2
		_1432_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F32())
		var _1433_v3 _dafny.Array
		_ = _1433_v3
		var _nwElement0_40 _dafny.Int = _dafny.IntOfUint32(((_1431_v1).Dtor_cf6()).Cardinality())
		_ = _nwElement0_40
		var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(6))
		_ = _nw214
		(_nw214).ArraySet1(_nwElement0_40, 0)
		(_nw214).ArraySet1((_dafny.Zero).Minus(p0), 1)
		(_nw214).ArraySet1(p0, 2)
		(_nw214).ArraySet1((_1432_v2).Cardinality(), 3)
		(_nw214).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}(func(_1434_i0 _dafny.Int) _dafny.CodePoint {
			return (_this).F32()
		}))).Cardinality()), 4)
		(_nw214).ArraySet1(p0, 5)
		_1433_v3 = _nw214
		var _1435_v4 _dafny.Sequence
		_ = _1435_v4
		_1435_v4 = _dafny.UnicodeSeqOfUtf8Bytes("wuevn")
		var _1436_v5 _dafny.Map
		_ = _1436_v5
		_1436_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1435_v4, _1433_v3)
		var _1437_v6 _dafny.Sequence
		_ = _1437_v6
		_1437_v6 = _dafny.SeqOf(_1433_v3, _1433_v3, _1433_v3, _1433_v3, _1433_v3)
		var _1438_v7 _dafny.Array
		_ = _1438_v7
		var _nwElement0_41 _dafny.Array = (func() _dafny.Array {
			if !((_this).F28()) {
				return _1433_v3
			}
			return _1433_v3
		})()
		_ = _nwElement0_41
		var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(16))
		_ = _nw215
		(_nw215).ArraySet1(_nwElement0_41, 0)
		(_nw215).ArraySet1(_1433_v3, 1)
		(_nw215).ArraySet1(_1433_v3, 2)
		(_nw215).ArraySet1((func() _dafny.Array {
			if (_1436_v5).Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1435_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1435_v4).Cardinality()))).Uint32(), Companion_Default___.Fm42(p0, globalState)), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1435_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1435_v4).Cardinality()))).Uint32(), Companion_Default___.Fm42(p0, globalState))).Cardinality()))).Uint32(), (_this).F32())) {
				return (_1436_v5).Get(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1435_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1435_v4).Cardinality()))).Uint32(), Companion_Default___.Fm42(p0, globalState)), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1435_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1435_v4).Cardinality()))).Uint32(), Companion_Default___.Fm42(p0, globalState))).Cardinality()))).Uint32(), (_this).F32())).(_dafny.Array)
			}
			return _1433_v3
		})(), 3)
		(_nw215).ArraySet1((_1437_v6).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1437_v6).Cardinality()))).Uint32()).(_dafny.Array), 4)
		(_nw215).ArraySet1((func() _dafny.Array {
			if (_this).F28() {
				return _1433_v3
			}
			return _1433_v3
		})(), 5)
		(_nw215).ArraySet1(_1433_v3, 6)
		(_nw215).ArraySet1(_1433_v3, 7)
		(_nw215).ArraySet1(_1433_v3, 8)
		(_nw215).ArraySet1(_1433_v3, 9)
		(_nw215).ArraySet1(_1433_v3, 10)
		(_nw215).ArraySet1(_1433_v3, 11)
		(_nw215).ArraySet1(_1433_v3, 12)
		(_nw215).ArraySet1(_1433_v3, 13)
		(_nw215).ArraySet1(_1433_v3, 14)
		(_nw215).ArraySet1(_1433_v3, 15)
		_1438_v7 = _nw215
		var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))
		_ = _index209
		(_1438_v7).ArraySet1(_1433_v3, (_index209).Int())
		var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))
		_ = _index210
		(_1438_v7).ArraySet1(_1433_v3, (_index210).Int())
		if (_this).F28() {
			var _1439_v8 _dafny.Set
			_ = _1439_v8
			_1439_v8 = _dafny.SetOf((_this).F28())
			var _1440_v9 _dafny.Sequence
			_ = _1440_v9
			_1440_v9 = _dafny.SeqOf(_1439_v8)
			var _1441_v10 D5
			_ = _1441_v10
			_1441_v10 = Companion_D5_.Create_DC12_(_1440_v9)
			var _1442_v11 _dafny.MultiSet
			_ = _1442_v11
			_1442_v11 = _dafny.MultiSetOf(false, Companion_Default___.Fm40(globalState))
			var _1443_v12 _dafny.Sequence
			_ = _1443_v12
			_1443_v12 = _dafny.SeqOf(_1442_v11)
			var _1444_v13 _dafny.Sequence
			_ = _1444_v13
			_1444_v13 = _1443_v12
			var _1445_v14 _dafny.Sequence
			_ = _1445_v14
			_1445_v14 = _dafny.SeqOf((_this).F28())
			var _1446_v15 _dafny.Map
			_ = _1446_v15
			_1446_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf((_this).F28()))
			var _rhs246 D5 = _1441_v10
			_ = _rhs246
			var _rhs247 bool = (_this).F28()
			_ = _rhs247
			var _rhs248 _dafny.Sequence = _1444_v13
			_ = _rhs248
			var _rhs249 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_1446_v15).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(167))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}(func(_1447_i1 _dafny.Int) _dafny.CodePoint {
					return (_this).F32()
				}))).Cardinality())) {
					return (_1446_v15).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(167))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg72 _dafny.Int) interface{} {
							return coer72(arg72)
						}
					}(func(_1448_i1 _dafny.Int) _dafny.CodePoint {
						return (_this).F32()
					}))).Cardinality())).(_dafny.Sequence)
				}
				return _1445_v14
			})(), _1445_v14), _dafny.Companion_Sequence_.Concatenate(_1445_v14, _1445_v14))
			_ = _rhs249
			var _lhs208 *GlobalState = globalState
			_ = _lhs208
			_1441_v10 = _rhs246
			_lhs208.F15 = _rhs247
			_1444_v13 = _rhs248
			_1445_v14 = _rhs249
			var _1449_v16 *C0
			_ = _1449_v16
			var _nw216 *C0 = New_C0_()
			_ = _nw216
			_nw216.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg73 _dafny.Int) interface{} {
					return coer73(arg73)
				}
			}(func(_1450_i2 _dafny.Int) _dafny.CodePoint {
				return (_this).F32()
			})), Companion_Default___.SafeDivisionInt(p0, p0), (_this).F27(), (_this).F28())
			_1449_v16 = _nw216
			(globalState).F17 = (_this).F28()
			(globalState).F21 = (Companion_Default___.Fm51((func() _dafny.Int {
				if (_this).F28() {
					return p0
				}
				return (_1446_v15).Cardinality()
			})(), globalState)).Cardinality()
			var _1451_v17 *C1
			_ = _1451_v17
			var _nw217 *C1 = New_C1_()
			_ = _nw217
			_nw217.Ctor__((_this).F28(), (_this).F27(), (p0).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(p0))) <= 0)
			_1451_v17 = _nw217
		} else {
			var _1452_v18 *C4
			_ = _1452_v18
			var _nw218 *C4 = New_C4_()
			_ = _nw218
			_nw218.Ctor__()
			_1452_v18 = _nw218
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))
			_ = _arr2
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))), 0))
			_ = _index211
			(_arr2).ArraySet1(p0, (_index211).Int())
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))
			_ = _arr3
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))), 0))
			_ = _index212
			(_arr3).ArraySet1(p0, (_index212).Int())
			var _1453_v19 _dafny.Sequence
			_ = _1453_v19
			_1453_v19 = _dafny.SeqOf(true)
			var _1454_v20 _dafny.Map
			_ = _1454_v20
			_1454_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1453_v19)
			var _1455_v21 _dafny.Map
			_ = _1455_v21
			_1455_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1454_v20, _dafny.IntOfInt64(-922))
			_1455_v21 = (_1455_v21).Update(_1454_v20, (_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))), 0))).Int()).(_dafny.Int))
			var _1456_v22 _dafny.MultiSet
			_ = _1456_v22
			_1456_v22 = _dafny.MultiSetOf(true, !(!((_this).F28())))
			var _1457_v23 _dafny.Sequence
			_ = _1457_v23
			_1457_v23 = _dafny.SeqOf(_1456_v22, _1456_v22)
			var _1458_v24 _dafny.Sequence
			_ = _1458_v24
			_1458_v24 = _1457_v23
			var _1459_v25 _dafny.Array
			_ = _1459_v25
			var _nwElement0_42 _dafny.Sequence = _1457_v23
			_ = _nwElement0_42
			var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(25))
			_ = _nw219
			(_nw219).ArraySet1(_nwElement0_42, 0)
			(_nw219).ArraySet1(_1458_v24, 1)
			(_nw219).ArraySet1(_1458_v24, 2)
			(_nw219).ArraySet1(_1458_v24, 3)
			(_nw219).ArraySet1(_1458_v24, 4)
			(_nw219).ArraySet1(_1458_v24, 5)
			(_nw219).ArraySet1(_1458_v24, 6)
			(_nw219).ArraySet1((func() _dafny.Sequence {
				if !((_this).F28()) {
					return _1458_v24
				}
				return _1458_v24
			})(), 7)
			(_nw219).ArraySet1(_1458_v24, 8)
			(_nw219).ArraySet1(_1457_v23, 9)
			(_nw219).ArraySet1((func() _dafny.Sequence {
				if true {
					return _1458_v24
				}
				return _1458_v24
			})(), 10)
			(_nw219).ArraySet1(_1458_v24, 11)
			(_nw219).ArraySet1(_1458_v24, 12)
			(_nw219).ArraySet1(Companion_Default___.Fm52(p0, p0, (_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))), 0))).Int()).(_dafny.Int), globalState), 13)
			(_nw219).ArraySet1(_1458_v24, 14)
			(_nw219).ArraySet1(_1458_v24, 15)
			(_nw219).ArraySet1((func() _dafny.Sequence {
				if (_this).F28() {
					return _1458_v24
				}
				return _1457_v23
			})(), 16)
			(_nw219).ArraySet1(_1458_v24, 17)
			(_nw219).ArraySet1(_1458_v24, 18)
			(_nw219).ArraySet1(_1458_v24, 19)
			(_nw219).ArraySet1(_1458_v24, 20)
			(_nw219).ArraySet1(_1458_v24, 21)
			(_nw219).ArraySet1(_1458_v24, 22)
			(_nw219).ArraySet1(_1457_v23, 23)
			(_nw219).ArraySet1(_1458_v24, 24)
			_1459_v25 = _nw219
			_1459_v25 = _1459_v25
			(globalState).F4 = (_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int()))), 0))).Int()).(_dafny.Int)
		}
		var _1460_v26 _dafny.Map
		_ = _1460_v26
		_1460_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F28()), p0)
		var _1461_v27 _dafny.Map
		_ = _1461_v27
		_1461_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1460_v26, _dafny.CodePoint('i'))
		var _1462_v28 _dafny.Sequence
		_ = _1462_v28
		_1462_v28 = _dafny.SeqOf((_this).F28(), (_this).F28())
		_1461_v27 = (_1461_v27).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1462_v28).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1462_v28).Cardinality()))).Uint32()).(bool), p0), (_this).F32())
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg74 _dafny.Int) interface{} {
				return coer74(arg74)
			}
		}((func(_1463_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1464_i4 _dafny.Int) _dafny.Int {
				return _1463_p0
			}
		})(p0)))).Cardinality())
		_ = _hi7
		for _1465_i3 := (p0).Times(p0); _1465_i3.Cmp(_hi7) < 0; _1465_i3 = _1465_i3.Plus(_dafny.One) {
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))
			_ = _index213
			var _rhs250 _dafny.Array = _1433_v3
			_ = _rhs250
			var _rhs251 _dafny.Int = (func() _dafny.Int {
				if !(true) {
					return _1465_i3
				}
				return (p0).Times(p0)
			})()
			_ = _rhs251
			var _lhs209 _dafny.Array = _1438_v7
			_ = _lhs209
			var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))
			_ = _lhs210
			var _lhs211 *GlobalState = globalState
			_ = _lhs211
			(_lhs209).ArraySet1(_rhs250, (_lhs210).Int())
			_lhs211.F13 = _rhs251
			var _1466_v29 _dafny.Map
			_ = _1466_v29
			_1466_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_i3, (_this).F28())
			var _1467_v31 _dafny.Set
			_ = _1467_v31
			_1467_v31 = _dafny.SetOf(_1466_v29, _1466_v29, (func() _dafny.Map {
				var _coll65 = _dafny.NewMapBuilder()
				_ = _coll65
				for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(515), _dafny.IntOfInt64(829))); ; {
					_compr_65, _ok67 := _iter67()
					if !_ok67 {
						break
					}
					var _1468_v30 _dafny.Int
					_1468_v30 = interface{}(_compr_65).(_dafny.Int)
					if ((_dafny.IntOfInt64(515)).Cmp(_1468_v30) <= 0) && ((_1468_v30).Cmp(_dafny.IntOfInt64(829)) < 0) {
						_coll65.Add((_1468_v30).Plus(p0), (_this).F28())
					}
				}
				return _coll65.ToMap()
			}()).Update(_1465_i3, (_this).F28()), _1466_v29, _1466_v29)
			_1460_v26 = (_1460_v26).Update((_this).F28(), (_1467_v31).Cardinality())
			var _1469_v32 D1
			_ = _1469_v32
			_1469_v32 = Companion_D1_.Create_DC7_((_this).F32(), _1435_v4)
			var _source22 D1 = _1469_v32
			_ = _source22
			if _source22.Is_DC6() {
				var _1470___mcc_h0 _dafny.Int = _source22.Get_().(D1_DC6).Cf12
				_ = _1470___mcc_h0
				var _1471_cf12 _dafny.Int = _1470___mcc_h0
				_ = _1471_cf12
				var _1472_v33 _dafny.Map
				_ = _1472_v33
				_1472_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1460_v26).Merge(_1460_v26), (_1471_cf12).Times((_dafny.Zero).Minus(_1465_i3)))
				_1472_v33 = _1472_v33
				var _1473_v34 _dafny.Map
				_ = _1473_v34
				_1473_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _1471_cf12)
				var _rhs252 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1435_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1435_v4).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_this).F28() {
						return (_1469_v32).Dtor_cf13()
					}
					return (_this).F32()
				})()), (Companion_Default___.SafeIndex(_1471_cf12, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1435_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1435_v4).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_this).F28() {
						return (_1469_v32).Dtor_cf13()
					}
					return (_this).F32()
				})())).Cardinality()))).Uint32(), (_this).F32())
				_ = _rhs252
				var _rhs253 bool = !((_1473_v34).Equals(_1473_v34))
				_ = _rhs253
				var _lhs212 *GlobalState = globalState
				_ = _lhs212
				_1435_v4 = _rhs252
				_lhs212.F15 = _rhs253
				var _1474_v35 _dafny.Map
				_ = _1474_v35
				var _1475_v36 _dafny.Int
				_ = _1475_v36
				var _out27 _dafny.Map
				_ = _out27
				var _out28 _dafny.Int
				_ = _out28
				_out27, _out28 = Companion_Default___.M0(globalState)
				_1474_v35 = _out27
				_1475_v36 = _out28
				var _1476_v37 *C4
				_ = _1476_v37
				var _nw220 *C4 = New_C4_()
				_ = _nw220
				_nw220.Ctor__()
				_1476_v37 = _nw220
				_1476_v37 = _1476_v37
			} else if _source22.Is_DC7() {
				var _1477___mcc_h1 _dafny.CodePoint = _source22.Get_().(D1_DC7).Cf13
				_ = _1477___mcc_h1
				var _1478___mcc_h2 _dafny.Sequence = _source22.Get_().(D1_DC7).Cf14
				_ = _1478___mcc_h2
				var _1479_cf14 _dafny.Sequence = _1478___mcc_h2
				_ = _1479_cf14
				var _1480_cf13 _dafny.CodePoint = _1477___mcc_h1
				_ = _1480_cf13
				var _1481_v38 _dafny.Map
				_ = _1481_v38
				_1481_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1482_v39 *C5
				_ = _1482_v39
				var _nw221 *C5 = New_C5_()
				_ = _nw221
				_nw221.Ctor__(_1437_v6, (_1481_v38).Cardinality(), (_this).F27(), (_this).F28())
				_1482_v39 = _nw221
				var _1483_v40 _dafny.MultiSet
				_ = _1483_v40
				_1483_v40 = _dafny.MultiSetOf(_1482_v39, _1482_v39)
				var _1484_v41 _dafny.Sequence
				_ = _1484_v41
				_1484_v41 = _dafny.SeqOf(_1479_cf14, _1479_cf14)
				var _rhs254 bool = !((_dafny.IntOfUint32((_1435_v4).Cardinality())).Cmp(p0) > 0) || ((_1483_v40).Equals(_1483_v40))
				_ = _rhs254
				var _rhs255 _dafny.Int = (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(796), _dafny.IntOfUint32((_1484_v41).Cardinality()))).Plus(p0)
				_ = _rhs255
				var _rhs256 _dafny.Int = _1465_i3
				_ = _rhs256
				var _rhs257 _dafny.Int = (_dafny.Zero).Minus(_1465_i3)
				_ = _rhs257
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				var _lhs215 *GlobalState = globalState
				_ = _lhs215
				var _lhs216 *GlobalState = globalState
				_ = _lhs216
				_lhs213.F15 = _rhs254
				_lhs214.F5 = _rhs255
				_lhs215.F5 = _rhs256
				_lhs216.F4 = _rhs257
				var _1485_v42 _dafny.MultiSet
				_ = _1485_v42
				_1485_v42 = _dafny.MultiSetOf(_1435_v4)
				var _1486_v43 _dafny.Map
				_ = _1486_v43
				_1486_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm30((_this).F28(), _dafny.UnicodeSeqOfUtf8Bytes("kcwjncj"), (_this).F28(), globalState)).Cardinality(), _1485_v42)
				var _1487_v44 _dafny.Sequence
				_ = _1487_v44
				_1487_v44 = _dafny.SeqOf(_1485_v42)
				_1486_v43 = (_1486_v43).Update(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}(func(_1488_i5 _dafny.Int) _dafny.CodePoint {
					return (_this).F32()
				}))).Cardinality()), (_1487_v44).Select((Companion_Default___.SafeIndex(_1465_i3, _dafny.IntOfUint32((_1487_v44).Cardinality()))).Uint32()).(_dafny.MultiSet))
				var _1489_v45 _dafny.Sequence
				_ = _1489_v45
				var _1490_v46 bool
				_ = _1490_v46
				var _out29 _dafny.Sequence
				_ = _out29
				var _out30 bool
				_ = _out30
				_out29, _out30 = (_this).M3((p0).Times((_dafny.Zero).Minus(p0)), globalState)
				_1489_v45 = _out29
				_1490_v46 = _out30
				var _1491_v47 *C1
				_ = _1491_v47
				var _nw222 *C1 = New_C1_()
				_ = _nw222
				_nw222.Ctor__((_this).F28(), (_this).F27(), (_dafny.IntOfInt64(998)).Cmp(_dafny.IntOfInt64(65)) >= 0)
				_1491_v47 = _nw222
			} else {
				var _1492___mcc_h3 _dafny.Sequence = _source22.Get_().(D1_DC5).Cf11
				_ = _1492___mcc_h3
				var _1493_cf11 _dafny.Sequence = _1492___mcc_h3
				_ = _1493_cf11
				(globalState).F15 = !((_this).F28())
				(globalState).F6 = (_this).F28()
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1433_v3), 0))
				_ = _index214
				(_1433_v3).ArraySet1(p0, (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1433_v3), 0))
				_ = _index215
				(_1433_v3).ArraySet1(((_1465_i3).Plus(_1465_i3)).Minus(_1465_i3), (_index215).Int())
				var _1494_v48 *C4
				_ = _1494_v48
				var _nw223 *C4 = New_C4_()
				_ = _nw223
				_nw223.Ctor__()
				_1494_v48 = _nw223
			}
			var _1495_v49 D7
			_ = _1495_v49
			_1495_v49 = Companion_D7_.Create_DC16_((_this).F28(), _1466_v29, (_this).F32(), (_this).F28())
			var _1496_v50 _dafny.Map
			_ = _1496_v50
			_1496_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), false)
			if Companion_Default___.Fm16((func() _dafny.Map {
				if (_1495_v49).Dtor_cf26() {
					return _1496_v50
				}
				return _1496_v50
			})(), globalState) {
				var _1497_v52 _dafny.Map
				_ = _1497_v52
				_1497_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F32())
				var _1498_v53 _dafny.MultiSet
				_ = _1498_v53
				_1498_v53 = _dafny.MultiSetOf((func() _dafny.CodePoint {
					if (_1497_v52).Contains(_1465_i3) {
						return (_1497_v52).Get(_1465_i3).(_dafny.CodePoint)
					}
					return (_this).F32()
				})())
				var _1499_v54 _dafny.MultiSet
				_ = _1499_v54
				_1499_v54 = _dafny.MultiSetOf((func() _dafny.Map {
					var _coll66 = _dafny.NewMapBuilder()
					_ = _coll66
					for _iter68 := _dafny.Iterate(((_1498_v53).Update((_this).F32(), Companion_Default___.Abs(_1465_i3))).Elements()); ; {
						_compr_66, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1500_v51 _dafny.CodePoint
						_1500_v51 = interface{}(_compr_66).(_dafny.CodePoint)
						if ((_1498_v53).Update((_this).F32(), Companion_Default___.Abs(_1465_i3))).Contains(_1500_v51) {
							_coll66.Add(_1500_v51, _1465_i3)
						}
					}
					return _coll66.ToMap()
				}()).Cardinality(), _1465_i3, _1465_i3, _1465_i3)
				var _1501_v55 _dafny.Map
				_ = _1501_v55
				_1501_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1432_v2, (_1499_v54).Cardinality())
				(globalState).F5 = (func() _dafny.Int {
					if (_1501_v55).Contains(_1432_v2) {
						return (_1501_v55).Get(_1432_v2).(_dafny.Int)
					}
					return (_dafny.IntOfUint32((_1462_v28).Cardinality())).Minus(_1465_i3)
				})()
				var _1502_v56 _dafny.Map
				_ = _1502_v56
				_1502_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(_1465_i3))
				(globalState).F21 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_1502_v56).Contains(_1465_i3) {
						return (_1502_v56).Get(_1465_i3).(_dafny.Int)
					}
					return p0
				})(), p0)
				var _1503_v57 D9
				_ = _1503_v57
				_1503_v57 = Companion_D9_.Create_DC23_(!((_this).F28()))
				var _1504_v58 D9
				_ = _1504_v58
				_1504_v58 = Companion_D9_.Create_DC24_(_1503_v57)
				var _1505_v59 D11
				_ = _1505_v59
				_1505_v59 = Companion_D11_.Create_DC28_()
				var _1506_v60 _dafny.Array
				_ = _1506_v60
				var _nw224 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
				_ = _nw224
				_1506_v60 = _nw224
				var _1507_v65 _dafny.Array
				_ = _1507_v65
				var _nwElement0_43 _dafny.Map = (_1466_v29).Merge(_1466_v29)
				_ = _nwElement0_43
				var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(25))
				_ = _nw225
				(_nw225).ArraySet1(_nwElement0_43, 0)
				(_nw225).ArraySet1(_1466_v29, 1)
				(_nw225).ArraySet1(_1466_v29, 2)
				(_nw225).ArraySet1((func() _dafny.Map {
					var _coll67 = _dafny.NewMapBuilder()
					_ = _coll67
					for _iter69 := _dafny.Iterate((func() _dafny.Set {
						var _coll68 = _dafny.NewBuilder()
						_ = _coll68
						for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(347), _dafny.IntOfInt64(714))); ; {
							_compr_68, _ok70 := _iter70()
							if !_ok70 {
								break
							}
							var _1508_v62 _dafny.Int
							_1508_v62 = interface{}(_compr_68).(_dafny.Int)
							if ((_dafny.IntOfInt64(347)).Cmp(_1508_v62) <= 0) && ((_1508_v62).Cmp(_dafny.IntOfInt64(714)) < 0) {
								_coll68.Add((_1508_v62).Minus(p0))
							}
						}
						return _coll68.ToSet()
					}()).Elements()); ; {
						_compr_67, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _1509_v61 _dafny.Int
						_1509_v61 = interface{}(_compr_67).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll69 = _dafny.NewBuilder()
							_ = _coll69
							for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(347), _dafny.IntOfInt64(714))); ; {
								_compr_69, _ok71 := _iter71()
								if !_ok71 {
									break
								}
								var _1510_v62 _dafny.Int
								_1510_v62 = interface{}(_compr_69).(_dafny.Int)
								if ((_dafny.IntOfInt64(347)).Cmp(_1510_v62) <= 0) && ((_1510_v62).Cmp(_dafny.IntOfInt64(714)) < 0) {
									_coll69.Add((_1510_v62).Minus(p0))
								}
							}
							return _coll69.ToSet()
						}()).Contains(_1509_v61) {
							_coll67.Add((_1509_v61).Minus(_1465_i3), (_this).F28())
						}
					}
					return _coll67.ToMap()
				}()).Merge(_1466_v29), 3)
				(_nw225).ArraySet1(_1466_v29, 4)
				(_nw225).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_i3, (_this).F28()), 5)
				(_nw225).ArraySet1(_1466_v29, 6)
				(_nw225).ArraySet1(_1466_v29, 7)
				(_nw225).ArraySet1(_1466_v29, 8)
				(_nw225).ArraySet1(_1466_v29, 9)
				(_nw225).ArraySet1((_1466_v29).Merge(_1466_v29), 10)
				(_nw225).ArraySet1((_1466_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())), 11)
				(_nw225).ArraySet1(_1466_v29, 12)
				(_nw225).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_i3, (_this).F28())).Merge((_1466_v29).Update(p0, (_this).F28())), 13)
				(_nw225).ArraySet1(Companion_Default___.Fm30((_this).F28(), _1435_v4, (_this).F28(), globalState), 14)
				(_nw225).ArraySet1(_1466_v29, 15)
				(_nw225).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_i3, (_this).F28())).Merge(_1466_v29), 16)
				(_nw225).ArraySet1(_1466_v29, 17)
				(_nw225).ArraySet1(Companion_Default___.Fm30((_this).F28(), _dafny.UnicodeSeqOfUtf8Bytes("vtw"), (_this).F28(), globalState), 18)
				(_nw225).ArraySet1((_1495_v49).Dtor_cf27(), 19)
				(_nw225).ArraySet1((_1466_v29).Merge(func() _dafny.Map {
					var _coll70 = _dafny.NewMapBuilder()
					_ = _coll70
					for _iter72 := _dafny.Iterate((_1466_v29).Keys().Elements()); ; {
						_compr_70, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _1511_v63 _dafny.Int
						_1511_v63 = interface{}(_compr_70).(_dafny.Int)
						if (_1466_v29).Contains(_1511_v63) {
							_coll70.Add((_1511_v63).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _1465_i3)).Cardinality()), (_this).F28())
						}
					}
					return _coll70.ToMap()
				}()), 20)
				(_nw225).ArraySet1((func() _dafny.Map {
					var _coll71 = _dafny.NewMapBuilder()
					_ = _coll71
					for _iter73 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(174), _dafny.IntOfInt64(816))); ; {
						_compr_71, _ok73 := _iter73()
						if !_ok73 {
							break
						}
						var _1512_v64 _dafny.Int
						_1512_v64 = interface{}(_compr_71).(_dafny.Int)
						if ((_dafny.IntOfInt64(174)).Cmp(_1512_v64) <= 0) && ((_1512_v64).Cmp(_dafny.IntOfInt64(816)) < 0) {
							_coll71.Add((_1512_v64).Minus((Companion_D1_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F28(), (_this).F28(), (_this).F28()), (_this).F28())).Cardinality())).Dtor_cf12()), (Companion_D9_.Create_DC23_(true)).Dtor_cf38())
						}
					}
					return _coll71.ToMap()
				}()).Merge(_1466_v29), 21)
				(_nw225).ArraySet1(((_1466_v29).Update(p0, (_this).F28())).Merge(_1466_v29), 22)
				(_nw225).ArraySet1(_1466_v29, 23)
				(_nw225).ArraySet1((_1495_v49).Dtor_cf27(), 24)
				_1507_v65 = _nw225
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_1507_v65), 0))
				_ = _index216
				(_1507_v65).ArraySet1(_1466_v29, (_index216).Int())
				var _1513_v66 _dafny.Array
				_ = _1513_v66
				_1513_v66 = (_this).F27()
				var _1514_v67 _dafny.Array
				_ = _1514_v67
				var _nwElement0_44 _dafny.Array = (_this).F27()
				_ = _nwElement0_44
				var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(28))
				_ = _nw226
				(_nw226).ArraySet1(_nwElement0_44, 0)
				(_nw226).ArraySet1((_this).F27(), 1)
				(_nw226).ArraySet1((_this).F27(), 2)
				(_nw226).ArraySet1((_this).F27(), 3)
				(_nw226).ArraySet1((_this).F27(), 4)
				(_nw226).ArraySet1((_this).F27(), 5)
				(_nw226).ArraySet1((_this).F27(), 6)
				(_nw226).ArraySet1((_1513_v66), 7)
				(_nw226).ArraySet1((_this).F27(), 8)
				(_nw226).ArraySet1((_this).F27(), 9)
				(_nw226).ArraySet1((_1513_v66), 10)
				(_nw226).ArraySet1((_1513_v66), 11)
				(_nw226).ArraySet1((_this).F27(), 12)
				(_nw226).ArraySet1((_this).F27(), 13)
				(_nw226).ArraySet1((_this).F27(), 14)
				(_nw226).ArraySet1((_this).F27(), 15)
				(_nw226).ArraySet1((_this).F27(), 16)
				(_nw226).ArraySet1((_this).F27(), 17)
				(_nw226).ArraySet1((_this).F27(), 18)
				(_nw226).ArraySet1((_this).F27(), 19)
				(_nw226).ArraySet1((_this).F27(), 20)
				(_nw226).ArraySet1((_this).F27(), 21)
				(_nw226).ArraySet1((_this).F27(), 22)
				(_nw226).ArraySet1((_this).F27(), 23)
				(_nw226).ArraySet1((_this).F27(), 24)
				(_nw226).ArraySet1((_1513_v66), 25)
				(_nw226).ArraySet1((_this).F27(), 26)
				(_nw226).ArraySet1((_this).F27(), 27)
				_1514_v67 = _nw226
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_1507_v65), 0))
				_ = _index217
				var _rhs258 D9 = Companion_D9_.Create_DC24_(Companion_D9_.Create_DC22_(_dafny.ArrayCastTo((_1438_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1438_v7), 0))).Int())), !(false), _1514_v67, _1435_v4))
				_ = _rhs258
				var _rhs259 D11 = _1505_v59
				_ = _rhs259
				var _rhs260 _dafny.Array = _1506_v60
				_ = _rhs260
				var _rhs261 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())).Merge((_1466_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F28())))
				_ = _rhs261
				var _lhs217 _dafny.Array = _1507_v65
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_1507_v65), 0))
				_ = _lhs218
				_1504_v58 = _rhs258
				_1505_v59 = _rhs259
				_1506_v60 = _rhs260
				(_lhs217).ArraySet1(_rhs261, (_lhs218).Int())
				(globalState).F25 = Companion_Default___.Fm42(p0, globalState)
				var _1515_v68 D0
				_ = _1515_v68
				_1515_v68 = Companion_D0_.Create_DC3_(p0, false, (_this).F28())
				var _1516_v69 _dafny.Map
				_ = _1516_v69
				_1516_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1515_v68, p0)
				var _1517_v70 _dafny.MultiSet
				_ = _1517_v70
				_1517_v70 = _dafny.MultiSetOf((_this).F28(), (_this).F28(), (_this).F28())
				var _1518_v71 T1
				_ = _1518_v71
				var _nw227 *C3 = New_C3_()
				_ = _nw227
				_nw227.Ctor__(_1465_i3, _1516_v69, (_1517_v70).Cardinality(), (_this).F27(), true)
				_1518_v71 = _nw227
				var _1519_v72 _dafny.Set
				_ = _1519_v72
				_1519_v72 = _dafny.SetOf(_1518_v71)
				(globalState).F15 = ((_1519_v72).Union(_dafny.SetOf(_1518_v71))).IsDisjointFrom(_1519_v72)
			} else {
				(globalState).F15 = !((_this).F28()) || ((_this).F28())
				(globalState).F15 = false
				var _1520_v73 *C1
				_ = _1520_v73
				var _nw228 *C1 = New_C1_()
				_ = _nw228
				_nw228.Ctor__((_this).F28(), (_this).F27(), (_this).F28())
				_1520_v73 = _nw228
				(globalState).F21 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p0)
				var _1521_v74 _dafny.MultiSet
				_ = _1521_v74
				_1521_v74 = _dafny.MultiSetOf(Companion_D1_.Create_DC6_(_1465_i3))
				(globalState).F13 = Companion_Default___.Fm4(_1521_v74, globalState)
			}
		}
		var _1522_v75 _dafny.Sequence
		_ = _1522_v75
		_1522_v75 = _dafny.SeqOf(_1430_v0)
		var _hi8 _dafny.Int = p0
		_ = _hi8
		for _1523_i6 := _dafny.IntOfUint32((_1522_v75).Cardinality()); _1523_i6.Cmp(_hi8) < 0; _1523_i6 = _1523_i6.Plus(_dafny.One) {
			if (_this).F28() {
				var _1524_v76 D1
				_ = _1524_v76
				_1524_v76 = Companion_D1_.Create_DC6_(_dafny.IntOfInt64(60))
				var _1525_v77 _dafny.MultiSet
				_ = _1525_v77
				_1525_v77 = _dafny.MultiSetOf(_1524_v76, _1524_v76)
				(globalState).F21 = Companion_Default___.Fm4(_1525_v77, globalState)
				var _1526_v78 _dafny.Map
				_ = _1526_v78
				_1526_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1460_v26, p0)
				var _1527_v79 D8
				_ = _1527_v79
				_1527_v79 = Companion_D8_.Create_DC17_(_1526_v78)
				var _1528_v80 _dafny.Sequence
				_ = _1528_v80
				_1528_v80 = _dafny.SeqOf(_1527_v79)
				(globalState).F21 = _dafny.IntOfUint32((_1528_v80).Cardinality())
				var _1529_v81 *C2
				_ = _1529_v81
				var _nw229 *C2 = New_C2_()
				_ = _nw229
				_nw229.Ctor__((_this).F32(), (_this).F28(), (_this).F27(), (_this).F28())
				_1529_v81 = _nw229
				var _rhs262 *C2 = _1529_v81
				_ = _rhs262
				var _rhs263 _dafny.Int = _1523_i6
				_ = _rhs263
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				_1529_v81 = _rhs262
				_lhs219.F13 = _rhs263
				var _1530_v82 *C5
				_ = _1530_v82
				var _nw230 *C5 = New_C5_()
				_ = _nw230
				_nw230.Ctor__(_1437_v6, _1523_i6, (_this).F27(), ((_this).F28()) || (true))
				_1530_v82 = _nw230
				var _1531_v83 _dafny.Map
				_ = _1531_v83
				_1531_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F28())), _1529_v81.F38)
				var _1532_v84 _dafny.MultiSet
				_ = _1532_v84
				_1532_v84 = _dafny.MultiSetOf(_1529_v81.F38)
				var _1533_v86 _dafny.Map
				_ = _1533_v86
				_1533_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v84, p0)
				var _1534_v87 _dafny.Map
				_ = _1534_v87
				_1534_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1523_i6, (_this).F28())
				var _1535_v88 _dafny.Map
				_ = _1535_v88
				_1535_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfInt64(-288))
				var _1536_v89 _dafny.Array
				_ = _1536_v89
				var _nwElement0_45 _dafny.Map = _1531_v83
				_ = _nwElement0_45
				var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(24))
				_ = _nw231
				(_nw231).ArraySet1(_nwElement0_45, 0)
				(_nw231).ArraySet1(_1531_v83, 1)
				(_nw231).ArraySet1((_1531_v83).Merge(_1531_v83), 2)
				(_nw231).ArraySet1((_1531_v83).Merge(_1531_v83), 3)
				(_nw231).ArraySet1(((_1531_v83).Update(_1532_v84, false)).Merge(func() _dafny.Map {
					var _coll72 = _dafny.NewMapBuilder()
					_ = _coll72
					for _iter74 := _dafny.Iterate((_1533_v86).Keys().Elements()); ; {
						_compr_72, _ok74 := _iter74()
						if !_ok74 {
							break
						}
						var _1537_v85 _dafny.MultiSet
						_1537_v85 = interface{}(_compr_72).(_dafny.MultiSet)
						if (_1533_v86).Contains(_1537_v85) {
							_coll72.Add(_1537_v85, _1529_v81.F38)
						}
					}
					return _coll72.ToMap()
				}()), 4)
				(_nw231).ArraySet1(_1531_v83, 5)
				(_nw231).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v84, (Companion_D7_.Create_DC16_(_1529_v81.F38, _1534_v87, (_1529_v81).F37(), _1529_v81.F38)).Dtor_cf29())).Update(_1532_v84, Companion_Default___.Fm22(_1529_v81.F38, _1523_i6, globalState)), 6)
				(_nw231).ArraySet1(_1531_v83, 7)
				(_nw231).ArraySet1((func() _dafny.Map {
					if !(_1529_v81.F38) {
						return _1531_v83
					}
					return _1531_v83
				})(), 8)
				(_nw231).ArraySet1((_1531_v83).Update((_1532_v84).Update((_this).F28(), Companion_Default___.Abs(p0)), _1529_v81.F38), 9)
				(_nw231).ArraySet1((func() _dafny.Map {
					if !((_this).F28()) {
						return (_1531_v83).Update(_dafny.MultiSetOf((_this).F28()), (_this).F28())
					}
					return _1531_v83
				})(), 10)
				(_nw231).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v84, !(_1529_v81.F38))).Merge((_1531_v83)), 11)
				(_nw231).ArraySet1(_1531_v83, 12)
				(_nw231).ArraySet1(_1531_v83, 13)
				(_nw231).ArraySet1((_1531_v83).Update(_1532_v84, (_this).F28()), 14)
				(_nw231).ArraySet1(_1531_v83, 15)
				(_nw231).ArraySet1(_1531_v83, 16)
				(_nw231).ArraySet1(_1531_v83, 17)
				(_nw231).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v84, !(true)), 18)
				(_nw231).ArraySet1(_1531_v83, 19)
				(_nw231).ArraySet1(_1531_v83, 20)
				(_nw231).ArraySet1(Companion_Default___.Fm53(p0, _1431_v1, globalState), 21)
				(_nw231).ArraySet1(_1531_v83, 22)
				(_nw231).ArraySet1(Companion_Default___.Fm53((_1535_v88).Cardinality(), _1431_v1, globalState), 23)
				_1536_v89 = _nw231
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_1536_v89), 0))
				_ = _index218
				(_1536_v89).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1532_v84).Update((_this).F28(), Companion_Default___.Abs(_1523_i6)), _1529_v81.F38), (_index218).Int())
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_1536_v89), 0))
				_ = _index219
				(_1536_v89).ArraySet1((_1531_v83).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v84, (_this).F28())).Merge(_1531_v83)), (_index219).Int())
			} else {
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index220
				((_this).F27()).ArraySet1((_this).F28(), (_index220).Int())
				var _1538_v90 _dafny.Set
				_ = _1538_v90
				_1538_v90 = _dafny.SetOf(_1462_v28)
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index221
				((_this).F27()).ArraySet1(!(_1538_v90).Contains(_1462_v28), (_index221).Int())
				var _1539_v91 D0
				_ = _1539_v91
				_1539_v91 = Companion_D0_.Create_DC3_(p0, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool), (_this).F28())
				_1539_v91 = _1539_v91
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1433_v3), 0))
				_ = _index222
				(_1433_v3).ArraySet1(p0, (_index222).Int())
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1433_v3), 0))
				_ = _index223
				(_1433_v3).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(470), (_dafny.Zero).Minus((_dafny.Zero).Minus((_1523_i6).Times((_dafny.Zero).Minus(p0))))), (_index223).Int())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1433_v3), 0))
				_ = _index224
				(_1433_v3).ArraySet1((_1433_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1433_v3), 0))).Int()).(_dafny.Int), (_index224).Int())
				(globalState).F4 = (_dafny.Zero).Minus(p0)
			}
			(globalState).F6 = !(Companion_Default___.Fm40(globalState))
			var _1540_v92 *C5
			_ = _1540_v92
			var _nw232 *C5 = New_C5_()
			_ = _nw232
			_nw232.Ctor__(_1437_v6, (_dafny.Zero).Minus((Companion_Default___.Fm30((_this).F28(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(875))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg76 _dafny.Int) interface{} {
					return coer76(arg76)
				}
			}(func(_1541_i7 _dafny.Int) _dafny.CodePoint {
				return (_this).F32()
			})), (_this).F28(), globalState)).Cardinality()), (func() _dafny.Array {
				if (_this).F28() {
					return (_this).F27()
				}
				return (_this).F27()
			})(), (_this).F28())
			_1540_v92 = _nw232
			var _1542_v93 _dafny.Map
			_ = _1542_v93
			_1542_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), p0)
			var _1543_v94 D8
			_ = _1543_v94
			_1543_v94 = Companion_D8_.Create_DC18_(_1523_i6)
			var _1544_v95 _dafny.Sequence
			_ = _1544_v95
			_1544_v95 = _dafny.SeqOf(p0)
			var _1545_v96 *C5
			_ = _1545_v96
			var _nw233 *C5 = New_C5_()
			_ = _nw233
			_nw233.Ctor__(_1437_v6, (func() _dafny.Int {
				if (_1542_v93).Contains((_1543_v94).Dtor_cf31()) {
					return (_1542_v93).Get((_1543_v94).Dtor_cf31()).(_dafny.Int)
				}
				return (_1544_v95).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.IntOfUint32((_1544_v95).Cardinality()))).Uint32()).(_dafny.Int)
			})(), (_this).F27(), (_this).F28())
			_1545_v96 = _nw233
		}
		var _hi9 _dafny.Int = p0
		_ = _hi9
		for _1546_i8 := Companion_Default___.SafeDivisionInt(p0, p0); _1546_i8.Cmp(_hi9) < 0; _1546_i8 = _1546_i8.Plus(_dafny.One) {
			(globalState).F4 = _1546_i8
			(globalState).F15 = !((_this).F28())
			var _1547_v97 D0
			_ = _1547_v97
			_1547_v97 = Companion_D0_.Create_DC3_((_1430_v0).Cardinality(), (_this).F28(), (_this).F28())
			var _1548_v98 _dafny.Map
			_ = _1548_v98
			_1548_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1547_v97, p0)
			var _1549_v99 D5
			_ = _1549_v99
			_1549_v99 = Companion_D5_.Create_DC13_(true, _1433_v3, (_this).F28())
			var _1550_v100 _dafny.Map
			_ = _1550_v100
			_1550_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1549_v99).Dtor_cf23())
			var _1551_v101 D3
			_ = _1551_v101
			_1551_v101 = Companion_D3_.Create_DC10_((_this).F28(), (func() bool {
				if (_1550_v100).Contains(p0) {
					return (_1550_v100).Get(p0).(bool)
				}
				return (_this).F28()
			})())
			var _1552_v102 *C3
			_ = _1552_v102
			var _nw234 *C3 = New_C3_()
			_ = _nw234
			_nw234.Ctor__((func() _dafny.Int {
				if (_this).F28() {
					return p0
				}
				return p0
			})(), (_1548_v98).Merge(_1548_v98), Companion_Default___.SafeModuloInt(p0, _1546_i8), (_this).F27(), (_1551_v101).Dtor_cf17())
			_1552_v102 = _nw234
			var _rhs264 _dafny.Sequence = _1435_v4
			_ = _rhs264
			var _rhs265 _dafny.Int = (Companion_Default___.SafeModuloInt(_1546_i8, p0)).Plus(((_1552_v102).F34()).Times(_1546_i8))
			_ = _rhs265
			var _lhs220 *GlobalState = globalState
			_ = _lhs220
			_1435_v4 = _rhs264
			_lhs220.F4 = _rhs265
		}
	}
}
func (_this *C8) F32() _dafny.CodePoint {
	{
		return _this._f32
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f29 _dafny.Int
	_f27 _dafny.Array
	_f28 bool
	_f30 bool
	_f31 _dafny.MultiSet
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f29 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f28 = false
	_this._f30 = false
	_this._f31 = _dafny.EmptyMultiSet
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

func (_this *C9) F29() _dafny.Int {
	return _this._f29
}
func (_this *C9) F29_set_(value _dafny.Int) {
	_this._f29 = value
}
func (_this *C9) F27() _dafny.Array {
	return _this._f27
}
func (_this *C9) F28() bool {
	return _this._f28
}
func (_this *C9) Ctor__(f30 bool, f31 _dafny.MultiSet, f29 _dafny.Int, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f29 = f29
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C9) M1(globalState *GlobalState) (_dafny.Int, bool, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _1553_v0 _dafny.Array
		_ = _1553_v0
		var _nw235 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw235
		_1553_v0 = _nw235
		var _1554_v1 _dafny.Map
		_ = _1554_v1
		_1554_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1553_v0, _1553_v0)
		_1554_v1 = (_1554_v1).Update(_1553_v0, _1553_v0)
		var _1555_v2 _dafny.Map
		_ = _1555_v2
		_1555_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F30())
		var _1556_v3 _dafny.Set
		_ = _1556_v3
		_1556_v3 = _dafny.SetOf(Companion_Default___.Fm1(globalState))
		var _1557_v4 _dafny.Sequence
		_ = _1557_v4
		_1557_v4 = _dafny.UnicodeSeqOfUtf8Bytes("hulb")
		var _1558_v5 D0
		_ = _1558_v5
		_1558_v5 = Companion_D0_.Create_DC3_((_1555_v2).Cardinality(), (_1556_v3).Contains(_this.F29()), Companion_Default___.Fm2((_1557_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F29())), _dafny.IntOfUint32((_1557_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F29(), globalState))
		var _source23 D0 = _1558_v5
		_ = _source23
		if _source23.Is_DC0() {
			var _1559___mcc_h0 _dafny.Map = _source23.Get_().(D0_DC0).Cf0
			_ = _1559___mcc_h0
			var _1560___mcc_h1 bool = _source23.Get_().(D0_DC0).Cf1
			_ = _1560___mcc_h1
			var _1561___mcc_h2 _dafny.Int = _source23.Get_().(D0_DC0).Cf2
			_ = _1561___mcc_h2
			var _1562___mcc_h3 _dafny.CodePoint = _source23.Get_().(D0_DC0).Cf3
			_ = _1562___mcc_h3
			var _1563___mcc_h4 bool = _source23.Get_().(D0_DC0).Cf4
			_ = _1563___mcc_h4
			var _1564_cf4 bool = _1563___mcc_h4
			_ = _1564_cf4
			var _1565_cf3 _dafny.CodePoint = _1562___mcc_h3
			_ = _1565_cf3
			var _1566_cf2 _dafny.Int = _1561___mcc_h2
			_ = _1566_cf2
			var _1567_cf1 bool = _1560___mcc_h1
			_ = _1567_cf1
			var _1568_cf0 _dafny.Map = _1559___mcc_h0
			_ = _1568_cf0
			(globalState).F5 = _this.F29()
			var _1569_v6 _dafny.Map
			_ = _1569_v6
			_1569_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1567_cf1, _1565_cf3)
			var _1570_v7 _dafny.MultiSet
			_ = _1570_v7
			_1570_v7 = _dafny.MultiSetOf(_1567_cf1, (_this).F28(), (_this).F28())
			(globalState).F15 = Companion_Default___.Fm2((func() _dafny.CodePoint {
				if (_1569_v6).Contains(false) {
					return (_1569_v6).Get(false).(_dafny.CodePoint)
				}
				return _1565_cf3
			})(), (func() _dafny.Int {
				if (_1570_v7).Contains((_this).F28()) {
					return (_1570_v7).Multiplicity((_this).F28())
				}
				return _this.F29()
			})(), globalState)
			(globalState).F5 = _this.F29()
			var _1571_v8 T0
			_ = _1571_v8
			var _nw236 *C2 = New_C2_()
			_ = _nw236
			_nw236.Ctor__(_1565_cf3, !(_1564_cf4), (_this).F27(), (_this).F30())
			_1571_v8 = _nw236
			var _1572_v9 _dafny.Map
			_ = _1572_v9
			_1572_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v8, (_this).F28())
			(globalState).F13 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_this.F29(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efreopb")).Cardinality())), (_1572_v9).Cardinality())
		} else if _source23.Is_DC1() {
			var _1573_v10 _dafny.Set
			_ = _1573_v10
			_1573_v10 = _dafny.SetOf(!((_this).F30()), (_this).F30(), Companion_Default___.Fm40(globalState), (_this).F30(), Companion_Default___.Fm16(_1555_v2, globalState))
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))
			_ = _index225
			(_1553_v0).ArraySet1((_1573_v10).Cardinality(), (_index225).Int())
			var _1574_v11 _dafny.Sequence
			_ = _1574_v11
			_1574_v11 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_this.F29(), _this.F29()))
			var _1575_v12 *C0
			_ = _1575_v12
			var _nw237 *C0 = New_C0_()
			_ = _nw237
			_nw237.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(501))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg77 _dafny.Int) interface{} {
					return coer77(arg77)
				}
			}(func(_1576_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), _this.F29(), (_this).F27(), (_this).F30())
			_1575_v12 = _nw237
			var _1577_v13 _dafny.Map
			_ = _1577_v13
			_1577_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1575_v12)
			var _1578_v14 _dafny.Set
			_ = _1578_v14
			_1578_v14 = _dafny.SetOf(_1577_v13, _1577_v13)
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))
			_ = _index226
			var _rhs266 _dafny.Int = (func() _dafny.Int {
				if (_1578_v14).Equals(_1578_v14) {
					return (_this.F29()).Minus(_this.F29())
				}
				return _this.F29()
			})()
			_ = _rhs266
			var _rhs267 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1574_v11, (Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1574_v11).Cardinality()))).Uint32(), _this.F29())
			_ = _rhs267
			var _lhs221 _dafny.Array = _1553_v0
			_ = _lhs221
			var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))
			_ = _lhs222
			(_lhs221).ArraySet1(_rhs266, (_lhs222).Int())
			_1574_v11 = _rhs267
			var _1579_v15 _dafny.Sequence
			_ = _1579_v15
			_1579_v15 = _dafny.SeqOf(true)
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))
			_ = _index227
			(_1553_v0).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_1579_v15).Cardinality())).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(928), _this.F29()))), (_index227).Int())
			var _1580_v16 D1
			_ = _1580_v16
			_1580_v16 = Companion_D1_.Create_DC6_((_1555_v2).Cardinality())
			(globalState).F13 = Companion_Default___.SafeModuloInt(_this.F29(), Companion_Default___.Fm4(_dafny.MultiSetOf(_1580_v16), globalState))
			var _1581_v17 _dafny.Map
			_ = _1581_v17
			_1581_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_this).F30() {
					return (func() _dafny.Int {
						if ((_this).F31()).Contains((_1553_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))).Int()).(_dafny.Int)) {
							return ((_this).F31()).Multiplicity((_1553_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))).Int()).(_dafny.Int))
						}
						return (_1553_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_1553_v0), 0))).Int()).(_dafny.Int)
					})()
				}
				return _this.F29()
			})(), Companion_Default___.Fm1(globalState))
			(globalState).F13 = (func() _dafny.Int {
				if (_1581_v17).Contains(_this.F29()) {
					return (_1581_v17).Get(_this.F29()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_this.F29())
			})()
		} else if _source23.Is_DC2() {
			var _1582___mcc_h5 _dafny.Set = _source23.Get_().(D0_DC2).Cf5
			_ = _1582___mcc_h5
			var _1583___mcc_h6 _dafny.Sequence = _source23.Get_().(D0_DC2).Cf6
			_ = _1583___mcc_h6
			var _1584_cf6 _dafny.Sequence = _1583___mcc_h6
			_ = _1584_cf6
			var _1585_cf5 _dafny.Set = _1582___mcc_h5
			_ = _1585_cf5
			var _1586_v18 _dafny.Array
			_ = _1586_v18
			var _nw238 _dafny.Array = _dafny.NewArrayWithValue(Companion_D10_.Default(), _dafny.IntOfInt64(15))
			_ = _nw238
			_1586_v18 = _nw238
			var _rhs268 _dafny.Int = ((_this.F29()).Minus(_this.F29())).Plus((_this.F29()).Minus(_this.F29()))
			_ = _rhs268
			var _rhs269 _dafny.Array = _1586_v18
			_ = _rhs269
			var _lhs223 *GlobalState = globalState
			_ = _lhs223
			_lhs223.F13 = _rhs268
			_1586_v18 = _rhs269
			var _1587_v19 _dafny.CodePoint
			_ = _1587_v19
			_1587_v19 = _dafny.CodePoint('v')
			var _1588_v20 _dafny.Sequence
			_ = _1588_v20
			_1588_v20 = _dafny.SeqOf((_this.F29()).Cmp(_this.F29()) < 0, (_this).F28(), (_this).F28(), (_dafny.IntOfUint32((_1557_v4).Cardinality())).Cmp(_this.F29()) < 0, !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_1557_v4, (Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1557_v4).Cardinality()))).Uint32(), _1587_v19), _1557_v4)))
			(globalState).F15 = (_1588_v20).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1588_v20).Cardinality()))).Uint32()).(bool)
			var _1589_v21 _dafny.Map
			_ = _1589_v21
			_1589_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())
			var _1590_v22 _dafny.Map
			_ = _1590_v22
			_1590_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_cf6, (_this).F27())
			_1589_v21 = (_1589_v21).Update((_this.F29()).Times((_1590_v22).Cardinality()), (_this).F28())
			var _1591_v23 _dafny.Array
			_ = _1591_v23
			var _nwElement0_46 _dafny.Array = (_this).F27()
			_ = _nwElement0_46
			var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(21))
			_ = _nw239
			(_nw239).ArraySet1(_nwElement0_46, 0)
			(_nw239).ArraySet1((_this).F27(), 1)
			(_nw239).ArraySet1((_this).F27(), 2)
			(_nw239).ArraySet1((_this).F27(), 3)
			(_nw239).ArraySet1((_this).F27(), 4)
			(_nw239).ArraySet1((_this).F27(), 5)
			(_nw239).ArraySet1((_this).F27(), 6)
			(_nw239).ArraySet1((_this).F27(), 7)
			(_nw239).ArraySet1((_this).F27(), 8)
			(_nw239).ArraySet1((_this).F27(), 9)
			(_nw239).ArraySet1((_this).F27(), 10)
			(_nw239).ArraySet1((_this).F27(), 11)
			(_nw239).ArraySet1((_this).F27(), 12)
			(_nw239).ArraySet1((_this).F27(), 13)
			(_nw239).ArraySet1((_this).F27(), 14)
			(_nw239).ArraySet1((_this).F27(), 15)
			(_nw239).ArraySet1((_this).F27(), 16)
			(_nw239).ArraySet1((_this).F27(), 17)
			(_nw239).ArraySet1((_this).F27(), 18)
			(_nw239).ArraySet1((_this).F27(), 19)
			(_nw239).ArraySet1((_this).F27(), 20)
			_1591_v23 = _nw239
			var _1592_v24 D9
			_ = _1592_v24
			_1592_v24 = Companion_D9_.Create_DC22_(_1553_v0, (_this).F28(), _1591_v23, _1584_cf6)
			_1557_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_1592_v24).Dtor_cf37(), _1584_cf6), _1557_v4)
		} else if _source23.Is_DC3() {
			var _1593___mcc_h7 _dafny.Int = _source23.Get_().(D0_DC3).Cf7
			_ = _1593___mcc_h7
			var _1594___mcc_h8 bool = _source23.Get_().(D0_DC3).Cf8
			_ = _1594___mcc_h8
			var _1595___mcc_h9 bool = _source23.Get_().(D0_DC3).Cf9
			_ = _1595___mcc_h9
			var _1596_cf9 bool = _1595___mcc_h9
			_ = _1596_cf9
			var _1597_cf8 bool = _1594___mcc_h8
			_ = _1597_cf8
			var _1598_cf7 _dafny.Int = _1593___mcc_h7
			_ = _1598_cf7
			var _1599_v25 _dafny.Array
			_ = _1599_v25
			var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw240
			_1599_v25 = _nw240
			_1599_v25 = _1599_v25
			(globalState).F17 = (func() bool {
				if (_this).F28() {
					return (_this).F30()
				}
				return _1596_cf9
			})()
			(globalState).F5 = Companion_Default___.Fm36(globalState)
			var _1600_v26 _dafny.Array
			_ = _1600_v26
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_31
			var _nw241 _dafny.Array
			_ = _nw241
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw241 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Sequence = func(_1601_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-313))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg78 _dafny.Int) interface{} {
							return coer78(arg78)
						}
					}(func(_1602_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					}))
				}
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw241 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw241).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw241).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1600_v26 = _nw241
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1600_v26), 0))
			_ = _index228
			(_1600_v26).ArraySet1(_1557_v4, (_index228).Int())
			var _1603_v27 _dafny.MultiSet
			_ = _1603_v27
			_1603_v27 = _dafny.MultiSetOf(_1553_v0)
			var _1604_v28 _dafny.CodePoint
			_ = _1604_v28
			_1604_v28 = _dafny.CodePoint('a')
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1600_v26), 0))
			_ = _index229
			var _rhs270 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fs"), _1557_v4), _dafny.Companion_Sequence_.Concatenate(_1557_v4, Companion_Default___.Fm18(_1596_cf9, _this.F29(), _1604_v28, globalState)))
			_ = _rhs270
			var _rhs271 _dafny.Int = _1598_cf7
			_ = _rhs271
			var _rhs272 bool = (_this).F28()
			_ = _rhs272
			var _rhs273 _dafny.MultiSet = (_1603_v27).Difference(_1603_v27)
			_ = _rhs273
			var _lhs224 _dafny.Array = _1600_v26
			_ = _lhs224
			var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1600_v26), 0))
			_ = _lhs225
			var _lhs226 *GlobalState = globalState
			_ = _lhs226
			(_lhs224).ArraySet1(_rhs270, (_lhs225).Int())
			_lhs226.F21 = _rhs271
			r2 = _rhs272
			_1603_v27 = _rhs273
		} else {
			var _1605___mcc_h10 D0 = _source23.Get_().(D0_DC4).Cf10
			_ = _1605___mcc_h10
			var _1606_cf10 D0 = _1605___mcc_h10
			_ = _1606_cf10
			var _1607_v29 _dafny.Sequence
			_ = _1607_v29
			_1607_v29 = _dafny.SeqOf((_this).F30(), (_this).F30(), (_this).F30())
			var _1608_v30 _dafny.Sequence
			_ = _1608_v30
			_1608_v30 = _dafny.SeqOf(_1553_v0)
			var _1609_v31 *C5
			_ = _1609_v31
			var _nw242 *C5 = New_C5_()
			_ = _nw242
			_nw242.Ctor__(_1608_v30, _this.F29(), (_this).F27(), (_this).F30())
			_1609_v31 = _nw242
			var _1610_v32 _dafny.Map
			_ = _1610_v32
			_1610_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1609_v31, (func() _dafny.Sequence {
				if (_this).F28() {
					return _dafny.SeqOf((_this).F30(), (_this).F28())
				}
				return _1607_v29
			})())
			_1607_v29 = (func() _dafny.Sequence {
				if (_1610_v32).Contains(_1609_v31) {
					return (_1610_v32).Get(_1609_v31).(_dafny.Sequence)
				}
				return _1607_v29
			})()
			(globalState).F15 = (_this).F28()
			var _1611_v33 T1
			_ = _1611_v33
			var _nw243 *C7 = New_C7_()
			_ = _nw243
			_nw243.Ctor__((_this).F27(), (_this).F28(), _this.F29())
			_1611_v33 = _nw243
			var _1612_v34 _dafny.Set
			_ = _1612_v34
			_1612_v34 = _dafny.SetOf(_1557_v4, _1557_v4)
			var _nw244 *C6 = New_C6_()
			_ = _nw244
			_nw244.Ctor__((_dafny.Zero).Minus(_1611_v33.F29()), (_this).F27(), (_1612_v34).IsSubsetOf(_1612_v34))
			_1611_v33 = _nw244
			(globalState).F13 = _1611_v33.F29()
		}
		var _hi10 _dafny.Int = _this.F29()
		_ = _hi10
		for _1613_i3 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1557_v4, _1557_v4)).Cardinality()); _1613_i3.Cmp(_hi10) < 0; _1613_i3 = _1613_i3.Plus(_dafny.One) {
			var _1614_v35 _dafny.Map
			_ = _1614_v35
			_1614_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1613_i3, _this.F29())
			var _1615_v36 *C0
			_ = _1615_v36
			var _nw245 *C0 = New_C0_()
			_ = _nw245
			_nw245.Ctor__(_1557_v4, (_1614_v35).Cardinality(), (_this).F27(), !((_this).F28()))
			_1615_v36 = _nw245
			var _1616_v37 _dafny.Array
			_ = _1616_v37
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_32
			var _nw246 _dafny.Array
			_ = _nw246
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw246 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Set = (func(_1617_v3 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_1618_i4 _dafny.Int) _dafny.Set {
						return _1617_v3
					}
				})(_1556_v3)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw246 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw246).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw246).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1616_v37 = _nw246
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1616_v37), 0))
			_ = _index230
			(_1616_v37).ArraySet1((_1556_v3).Intersection(_1556_v3), (_index230).Int())
			var _1619_v38 _dafny.Sequence
			_ = _1619_v38
			_1619_v38 = _dafny.SeqOf(_this.F29())
			var _1620_v39 _dafny.Sequence
			_ = _1620_v39
			_1620_v39 = _dafny.SeqOf(_1619_v38)
			var _1621_v40 _dafny.Map
			_ = _1621_v40
			_1621_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(_1620_v39), (_this).F28())
			var _1622_v41 _dafny.Sequence
			_ = _1622_v41
			_1622_v41 = _dafny.SeqOf((_this).F30(), (_this).F30(), (_this).F30())
			var _1623_v42 _dafny.Sequence
			_ = _1623_v42
			_1623_v42 = _dafny.SeqOf((_1621_v40).Cardinality(), _1613_i3, _dafny.IntOfUint32((_1622_v41).Cardinality()))
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1616_v37), 0))
			_ = _index231
			var _rhs274 *C0 = _1615_v36
			_ = _rhs274
			var _rhs275 _dafny.Set = _1556_v3
			_ = _rhs275
			var _rhs276 _dafny.Int = ((_1623_v42).Select((Companion_Default___.SafeIndex(_1613_i3, _dafny.IntOfUint32((_1623_v42).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_this.F29())
			_ = _rhs276
			var _lhs227 _dafny.Array = _1616_v37
			_ = _lhs227
			var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1616_v37), 0))
			_ = _lhs228
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			_1615_v36 = _rhs274
			(_lhs227).ArraySet1(_rhs275, (_lhs228).Int())
			_lhs229.F21 = _rhs276
			_1623_v42 = _1619_v38
			var _1624_v43 _dafny.Array
			_ = _1624_v43
			var _nwElement0_47 _dafny.Sequence = _1557_v4
			_ = _nwElement0_47
			var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(7))
			_ = _nw247
			(_nw247).ArraySet1(_nwElement0_47, 0)
			(_nw247).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uhkbfbg"), 1)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1557_v4, _dafny.UnicodeSeqOfUtf8Bytes("nnmatc")), 2)
			(_nw247).ArraySet1((_1615_v36).F36(), 3)
			(_nw247).ArraySet1(Companion_Default___.Fm9((_this).F28(), globalState), 4)
			(_nw247).ArraySet1((func() _dafny.Sequence {
				if Companion_Default___.Fm40(globalState) {
					return _1557_v4
				}
				return _1557_v4
			})(), 5)
			(_nw247).ArraySet1(_1557_v4, 6)
			_1624_v43 = _nw247
			_1624_v43 = _1624_v43
			if (_this).F30() {
				var _1625_v44 _dafny.Map
				_ = _1625_v44
				_1625_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1558_v5, _dafny.IntOfInt64(20))
				var _1626_v45 _dafny.Map
				_ = _1626_v45
				_1626_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _1625_v44)
				var _1627_v46 T1
				_ = _1627_v46
				var _nw248 *C3 = New_C3_()
				_ = _nw248
				_nw248.Ctor__(_dafny.IntOfUint32((_1557_v4).Cardinality()), (func() _dafny.Map {
					if (_1626_v45).Contains(!((_this).F28())) {
						return (_1626_v45).Get(!((_this).F28())).(_dafny.Map)
					}
					return _1625_v44
				})(), _dafny.IntOfInt64(568), (_this).F27(), (_this).F28())
				_1627_v46 = _nw248
				var _1628_v47 _dafny.Map
				_ = _1628_v47
				_1628_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1627_v46.F29(), _1627_v46)
				var _1629_v48 D1
				_ = _1629_v48
				_1629_v48 = Companion_D1_.Create_DC6_(_this.F29())
				var _1630_v49 _dafny.MultiSet
				_ = _1630_v49
				_1630_v49 = _dafny.MultiSetOf(_1629_v48)
				var _1631_v50 _dafny.Array
				_ = _1631_v50
				var _nwElement0_48 T1 = _1627_v46
				_ = _nwElement0_48
				var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(16))
				_ = _nw249
				(_nw249).ArraySet1(_nwElement0_48, 0)
				(_nw249).ArraySet1(_1627_v46, 1)
				(_nw249).ArraySet1((func() T1 {
					if (_this).F30() {
						return _1627_v46
					}
					return _1627_v46
				})(), 2)
				(_nw249).ArraySet1(_1627_v46, 3)
				(_nw249).ArraySet1(_1627_v46, 4)
				(_nw249).ArraySet1(_1627_v46, 5)
				(_nw249).ArraySet1(_1627_v46, 6)
				(_nw249).ArraySet1(_1627_v46, 7)
				(_nw249).ArraySet1(_1627_v46, 8)
				(_nw249).ArraySet1(_1627_v46, 9)
				(_nw249).ArraySet1((func() T1 {
					if (_1628_v47).Contains(Companion_Default___.Fm4(_1630_v49, globalState)) {
						return (_1628_v47).Get(Companion_Default___.Fm4(_1630_v49, globalState)).(T1)
					}
					return _1627_v46
				})(), 10)
				(_nw249).ArraySet1(_1627_v46, 11)
				(_nw249).ArraySet1(_1627_v46, 12)
				(_nw249).ArraySet1(_1627_v46, 13)
				(_nw249).ArraySet1(_1627_v46, 14)
				(_nw249).ArraySet1(_1627_v46, 15)
				_1631_v50 = _nw249
				var _1632_v51 _dafny.Array
				_ = _1632_v51
				var _nw250 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.One)
				_ = _nw250
				_1632_v51 = _nw250
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1632_v51), 0))
				_ = _index232
				(_1632_v51).ArraySet1(_1558_v5, (_index232).Int())
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1632_v51), 0))
				_ = _index233
				var _rhs277 _dafny.Array = _1631_v50
				_ = _rhs277
				var _rhs278 D0 = _1558_v5
				_ = _rhs278
				var _lhs230 _dafny.Array = _1632_v51
				_ = _lhs230
				var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_1632_v51), 0))
				_ = _lhs231
				_1631_v50 = _rhs277
				(_lhs230).ArraySet1(_rhs278, (_lhs231).Int())
				(globalState).F4 = ((_1613_i3).Times((_dafny.Zero).Minus(_1613_i3))).Plus(_1627_v46.F29())
				var _1633_v52 _dafny.Sequence
				_ = _1633_v52
				_1633_v52 = _1619_v38
				var _1634_v53 _dafny.MultiSet
				_ = _1634_v53
				_1634_v53 = _dafny.MultiSetOf((_this).F30())
				_1619_v38 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_1633_v52), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1615_v36).F36()).Cardinality())), _dafny.IntOfUint32((_1633_v52).Cardinality()))).Uint32(), (_1634_v53).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}((func(_1635_v36 *C0) func(_dafny.Int) _dafny.Int {
					return func(_1636_i5 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32(((_1635_v36).F36()).Cardinality())
					}
				})(_1615_v36))))
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_1553_v0), 0))
				_ = _index234
				(_1553_v0).ArraySet1(_dafny.IntOfInt64(310), (_index234).Int())
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_1553_v0), 0))
				_ = _index235
				(_1553_v0).ArraySet1(_this.F29(), (_index235).Int())
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1624_v43), 0))
				_ = _index236
				(_1624_v43).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}(func(_1637_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				})), (_index236).Int())
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1624_v43), 0))
				_ = _index237
				(_1624_v43).ArraySet1((_1615_v36).F36(), (_index237).Int())
			} else {
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index238
				((_this).F27()).ArraySet1((_this).F28(), (_index238).Int())
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index239
				((_this).F27()).ArraySet1((_this).F28(), (_index239).Int())
				(globalState).F4 = (_dafny.IntOfInt64(-30)).Plus(((_1616_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1616_v37), 0))).Int()).(_dafny.Set)).Cardinality())
				var _1638_v54 _dafny.Array
				_ = _1638_v54
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_33
				var _nw251 _dafny.Array
				_ = _nw251
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw251 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = func(_1639_i7 _dafny.Int) bool {
						return !(true)
					}
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw251 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw251).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw251).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1638_v54 = _nw251
				_1638_v54 = _1638_v54
				r1 = !(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool))
				var _1640_v55 _dafny.Sequence
				_ = _1640_v55
				_1640_v55 = _dafny.SeqOf(_1556_v3, _1556_v3, (_1556_v3).Intersection((_1616_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1616_v37), 0))).Int()).(_dafny.Set)))
				_1640_v55 = _dafny.SeqOf(_1556_v3, (_1616_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1616_v37), 0))).Int()).(_dafny.Set))
			}
		}
		var _1641_v56 _dafny.Set
		_ = _1641_v56
		_1641_v56 = _dafny.SetOf((_this).F28(), !((_this).F30()), (_this).F30(), (_this).F28(), (_this).F28())
		_1641_v56 = _1641_v56
		var _1642_v57 _dafny.CodePoint
		_ = _1642_v57
		_1642_v57 = _dafny.CodePoint('c')
		var _1643_v58 _dafny.Map
		_ = _1643_v58
		_1643_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1642_v57)
		var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen(((_this).F27()), 0))
		_ = _index240
		((_this).F27()).ArraySet1((_this).F30(), (_index240).Int())
		var _1644_v59 _dafny.Sequence
		_ = _1644_v59
		_1644_v59 = _dafny.SeqOf(!((_this).F30()) || ((_this).F30()), Companion_Default___.Fm16(_1555_v2, globalState), ((_dafny.Zero).Minus(_this.F29())).Cmp(_this.F29()) < 0)
		var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen(((_this).F27()), 0))
		_ = _index241
		var _rhs279 _dafny.Int = _this.F29()
		_ = _rhs279
		var _rhs280 _dafny.Map = _1643_v58
		_ = _rhs280
		var _rhs281 bool = (_1644_v59).Select((Companion_Default___.SafeIndex(_this.F29(), _dafny.IntOfUint32((_1644_v59).Cardinality()))).Uint32()).(bool)
		_ = _rhs281
		var _lhs232 *GlobalState = globalState
		_ = _lhs232
		var _lhs233 _dafny.Array = (_this).F27()
		_ = _lhs233
		var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen(((_this).F27()), 0))
		_ = _lhs234
		_lhs232.F5 = _rhs279
		_1643_v58 = _rhs280
		(_lhs233).ArraySet1(_rhs281, (_lhs234).Int())
		var _hi11 _dafny.Int = (_dafny.IntOfInt64(585)).Times(_this.F29())
		_ = _hi11
		for _1645_i8 := _dafny.IntOfUint32((_1557_v4).Cardinality()); _1645_i8.Cmp(_hi11) < 0; _1645_i8 = _1645_i8.Plus(_dafny.One) {
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index242
			var _rhs282 bool = (_this).F28()
			_ = _rhs282
			var _rhs283 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1557_v4, _1557_v4)
			_ = _rhs283
			var _rhs284 _dafny.Int = (_this.F29()).Plus(Companion_Default___.Fm36(globalState))
			_ = _rhs284
			var _rhs285 bool = ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(bool)
			_ = _rhs285
			var _lhs235 _dafny.Array = (_this).F27()
			_ = _lhs235
			var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _lhs236
			var _lhs237 *GlobalState = globalState
			_ = _lhs237
			var _lhs238 *GlobalState = globalState
			_ = _lhs238
			(_lhs235).ArraySet1(_rhs282, (_lhs236).Int())
			_1557_v4 = _rhs283
			_lhs237.F5 = _rhs284
			_lhs238.F17 = _rhs285
			var _1646_v60 _dafny.Array
			_ = _1646_v60
			var _nw252 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw252
			_1646_v60 = _nw252
			var _1647_v61 _dafny.Array
			_ = _1647_v61
			var _nwElement0_49 _dafny.Array = _1646_v60
			_ = _nwElement0_49
			var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(6))
			_ = _nw253
			(_nw253).ArraySet1(_nwElement0_49, 0)
			(_nw253).ArraySet1(_1646_v60, 1)
			(_nw253).ArraySet1(_1646_v60, 2)
			(_nw253).ArraySet1(_1646_v60, 3)
			(_nw253).ArraySet1(_1646_v60, 4)
			(_nw253).ArraySet1(_1646_v60, 5)
			_1647_v61 = _nw253
			var _1648_v62 D9
			_ = _1648_v62
			_1648_v62 = Companion_D9_.Create_DC22_(_1553_v0, !((_this).F30()), _1647_v61, _1557_v4)
			r2 = (_1648_v62).Dtor_cf35()
			r1 = (_this).F28()
			var _1649_v63 _dafny.Set
			_ = _1649_v63
			_1649_v63 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("odjwsypnq"), _1557_v4)
			var _1650_v64 _dafny.Map
			_ = _1650_v64
			_1650_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _1649_v63)
			_1650_v64 = (_1650_v64).Update((_this).F30(), _1649_v63)
		}
		r0 = (_this.F29()).Minus(_this.F29())
		r1 = (_this).F30()
		r2 = ((_this).F30()) == ((_this).F30())
		r3 = (_this).F30()
		return r0, r1, r2, r3
	}
}
func (_this *C9) M2(globalState *GlobalState) (bool, _dafny.Int, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1651_v0 _dafny.Sequence
		_ = _1651_v0
		_1651_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wfqmokab")
		_1651_v0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm32(_this.F29(), (_this).F30(), globalState), _1651_v0)
		(globalState).F13 = _this.F29()
		var _1652_v1 _dafny.Array
		_ = _1652_v1
		var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw254
		_1652_v1 = _nw254
		var _1653_v2 _dafny.Array
		_ = _1653_v2
		var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
		_ = _nw255
		_1653_v2 = _nw255
		var _1654_v3 D9
		_ = _1654_v3
		_1654_v3 = Companion_D9_.Create_DC22_(_1652_v1, (_this).F30(), _1653_v2, _1651_v0)
		var _1655_v4 _dafny.Map
		_ = _1655_v4
		_1655_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _this.F29())
		var _1656_v5 _dafny.Map
		_ = _1656_v5
		_1656_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _1655_v4)
		var _1657_v6 _dafny.Map
		_ = _1657_v6
		_1657_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1656_v5).Cardinality(), (_this).F28())
		var _1658_v7 _dafny.Array
		_ = _1658_v7
		var _nwElement0_50 bool = (_this).F30()
		_ = _nwElement0_50
		var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(17))
		_ = _nw256
		(_nw256).ArraySet1(_nwElement0_50, 0)
		(_nw256).ArraySet1((_this).F28(), 1)
		(_nw256).ArraySet1((_this).F28(), 2)
		(_nw256).ArraySet1((_1654_v3).Dtor_cf35(), 3)
		(_nw256).ArraySet1((_this).F28(), 4)
		(_nw256).ArraySet1(false, 5)
		(_nw256).ArraySet1(true, 6)
		(_nw256).ArraySet1((_this.F29()).Cmp(_dafny.IntOfInt64(-922)) != 0, 7)
		(_nw256).ArraySet1((_this).F30(), 8)
		(_nw256).ArraySet1((_this).F30(), 9)
		(_nw256).ArraySet1((_this).F28(), 10)
		(_nw256).ArraySet1((_this).F30(), 11)
		(_nw256).ArraySet1((Companion_D9_.Create_DC22_(_1652_v1, (_this).F28(), _1653_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-72))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1659_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})))).Dtor_cf35(), 12)
		(_nw256).ArraySet1((_this).F30(), 13)
		(_nw256).ArraySet1((_this).F30(), 14)
		(_nw256).ArraySet1((func() bool {
			if (_1657_v6).Contains(_this.F29()) {
				return (_1657_v6).Get(_this.F29()).(bool)
			}
			return !(false)
		})(), 15)
		(_nw256).ArraySet1(true, 16)
		_1658_v7 = _nw256
		for _iter75 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1658_v7), 0))); ; {
			_guard_loop_2, _ok75 := _iter75()
			if !_ok75 {
				break
			}
			var _1660_i0 _dafny.Int
			_1660_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1660_i0).Sign() != -1) && ((_1660_i0).Cmp(_dafny.ArrayLen((_1658_v7), 0)) < 0)) {
				(_1658_v7).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Set {
					var _coll73 = _dafny.NewBuilder()
					_ = _coll73
					for _iter76 := _dafny.Iterate((_dafny.SetOf(Companion_D7_.Create_DC16_((_this).F30(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(747), _this.F29(), _this.F29())).Cardinality())).Cardinality()), (_this).F28()), _dafny.CodePoint('q'), (_this).F28()), Companion_D7_.Create_DC16_((_this).F30(), _1657_v6, _dafny.CodePoint('k'), !((_this).F28())), Companion_D7_.Create_DC16_((_this).F30(), (_1657_v6).Update(_this.F29(), false), _dafny.CodePoint('g'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F28(), _1657_v6, _dafny.CodePoint('v'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F30(), _1657_v6, _dafny.CodePoint('b'), (_this).F28()))).Elements()); ; {
						_compr_73, _ok76 := _iter76()
						if !_ok76 {
							break
						}
						var _1661_v8 D7
						_1661_v8 = interface{}(_compr_73).(D7)
						if (_dafny.SetOf(Companion_D7_.Create_DC16_((_this).F30(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(747), _this.F29(), _this.F29())).Cardinality())).Cardinality()), (_this).F28()), _dafny.CodePoint('q'), (_this).F28()), Companion_D7_.Create_DC16_((_this).F30(), _1657_v6, _dafny.CodePoint('k'), !((_this).F28())), Companion_D7_.Create_DC16_((_this).F30(), (_1657_v6).Update(_this.F29(), false), _dafny.CodePoint('g'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F28(), _1657_v6, _dafny.CodePoint('v'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F30(), _1657_v6, _dafny.CodePoint('b'), (_this).F28()))).Contains(_1661_v8) {
							_coll73.Add(_1661_v8)
						}
					}
					return _coll73.ToSet()
				}()), _dafny.SeqOf(_dafny.SetOf(Companion_D7_.Create_DC16_((_this).F30(), (func() _dafny.Map {
					var _coll74 = _dafny.NewMapBuilder()
					_ = _coll74
					for _iter77 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg82 _dafny.Int) interface{} {
							return coer82(arg82)
						}
					}(func(_1662_i2 _dafny.Int) _dafny.Int {
						return _this.F29()
					}))).Elements()); ; {
						_compr_74, _ok77 := _iter77()
						if !_ok77 {
							break
						}
						var _1663_v9 _dafny.Int
						_1663_v9 = interface{}(_compr_74).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg83 _dafny.Int) interface{} {
								return coer83(arg83)
							}
						}(func(_1662_i2 _dafny.Int) _dafny.Int {
							return _this.F29()
						}))), _1663_v9) {
							_coll74.Add(Companion_Default___.SafeDivisionInt(_1663_v9, _this.F29()), (_this).F30())
						}
					}
					return _coll74.ToMap()
				}()).Update(_this.F29(), (_this).F30()), _dafny.CodePoint('k'), (_this).F28()), Companion_D7_.Create_DC16_((_this).F30(), _1657_v6, _dafny.CodePoint('j'), (_this).F28()), Companion_D7_.Create_DC16_((_this).F28(), _1657_v6, _dafny.CodePoint('h'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F28(), _1657_v6, _dafny.CodePoint('o'), (_this).F28())))), (_dafny.SetOf(Companion_D7_.Create_DC16_((_this).F30(), _1657_v6, _dafny.CodePoint('b'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F28(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30()), _dafny.CodePoint('f'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F28(), _1657_v6, _dafny.CodePoint('y'), (_this).F28()), Companion_D7_.Create_DC16_((_this).F28(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30()), _dafny.CodePoint('k'), (_this).F30()), Companion_D7_.Create_DC16_((_this).F28(), _1657_v6, _dafny.CodePoint('p'), (_this).F28()))).Difference(_dafny.SetOf(Companion_D7_.Create_DC16_(true, (_1657_v6).Update((_dafny.MultiSetOf((_this).F28(), (_this).F28())).Cardinality(), (_this).F28()), _dafny.CodePoint('d'), (_this).F28()), Companion_D7_.Create_DC16_((_this).F28(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F28())).Update(_dafny.IntOfInt64(558), (_this).F28()), _dafny.CodePoint('m'), true)))), (_1660_i0).Int())
			}
		}
		for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1652_v1), 0))); ; {
			_guard_loop_3, _ok78 := _iter78()
			if !_ok78 {
				break
			}
			var _1664_i3 _dafny.Int
			_1664_i3 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1664_i3).Sign() != -1) && ((_1664_i3).Cmp(_dafny.ArrayLen((_1652_v1), 0)) < 0)) {
				(_1652_v1).ArraySet1((_1664_i3).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1655_v4, (_this).F28())).Cardinality()), (_1664_i3).Int())
			}
		}
		r0 = (_this.F29()).Cmp(_this.F29()) == 0
		var _1665_v10 _dafny.Set
		_ = _1665_v10
		_1665_v10 = _dafny.SetOf(_this.F29())
		_1665_v10 = (_1665_v10).Union(func() _dafny.Set {
			var _coll75 = _dafny.NewBuilder()
			_ = _coll75
			for _iter79 := _dafny.Iterate((_dafny.SeqOf(_this.F29(), _dafny.IntOfInt64(250))).Elements()); ; {
				_compr_75, _ok79 := _iter79()
				if !_ok79 {
					break
				}
				var _1666_v11 _dafny.Int
				_1666_v11 = interface{}(_compr_75).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_this.F29(), _dafny.IntOfInt64(250)), _1666_v11) {
					_coll75.Add(Companion_Default___.SafeDivisionInt(_1666_v11, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfInt64(44))).Cardinality())))
				}
			}
			return _coll75.ToSet()
		}())
		r0 = false
		r1 = _this.F29()
		r2 = !((_this).F30()) || ((_this.F29()).Cmp(_this.F29()) != 0)
		var _1667_v12 _dafny.Map
		_ = _1667_v12
		_1667_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F30())
		var _1668_v13 _dafny.Map
		_ = _1668_v13
		_1668_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1667_v12, (_this).F28())
		var _1669_v14 _dafny.Sequence
		_ = _1669_v14
		_1669_v14 = _dafny.SeqOf(_this.F29())
		r3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1668_v13).Cardinality()), _1669_v14)).Cardinality())
		return r0, r1, r2, r3
	}
}
func (_this *C9) F30() bool {
	{
		return _this._f30
	}
}
func (_this *C9) F31() _dafny.MultiSet {
	{
		return _this._f31
	}
}

// End of class C9
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
