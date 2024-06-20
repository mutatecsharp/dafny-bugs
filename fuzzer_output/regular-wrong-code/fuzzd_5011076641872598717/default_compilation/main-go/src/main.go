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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (((_dafny.MultiSetOf(_dafny.IntOfInt64(699), _dafny.IntOfInt64(128))).Intersection(_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.CodePoint('y'))).Cardinality(), _dafny.IntOfInt64(-261)))).Difference((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(64), _dafny.IntOfInt64(565))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(64)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(565)) < 0) {
				_coll0.Add((_0_v0).Minus((_dafny.SetOf(_dafny.CodePoint('u'))).Cardinality()), true)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Cardinality())).Difference(_dafny.MultiSetOf((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(556), _dafny.IntOfInt64(176))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _1_v1 _dafny.Int
			_1_v1 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(556)).Cmp(_1_v1) <= 0) && ((_1_v1).Cmp(_dafny.IntOfInt64(176)) < 0) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_1_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(992))).Cardinality())), true)
			}
		}
		return _coll1.ToMap()
	}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(402), _dafny.IntOfInt64(682))).Cardinality())))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("be"), _dafny.UnicodeSeqOfUtf8Bytes("mfixa"))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D5 = Companion_D5_.Create_DC17_(_dafny.IntOfInt64(400), false, !(false), !(false), (_dafny.MultiSetOf(false, false)).Cardinality())
	_ = _source0
	if _source0.Is_DC17() {
		var _2___mcc_h0 _dafny.Int = _source0.Get_().(D5_DC17).Cf27
		_ = _2___mcc_h0
		var _3___mcc_h1 bool = _source0.Get_().(D5_DC17).Cf28
		_ = _3___mcc_h1
		var _4___mcc_h2 bool = _source0.Get_().(D5_DC17).Cf29
		_ = _4___mcc_h2
		var _5___mcc_h3 bool = _source0.Get_().(D5_DC17).Cf30
		_ = _5___mcc_h3
		var _6___mcc_h4 _dafny.Int = _source0.Get_().(D5_DC17).Cf31
		_ = _6___mcc_h4
		var _7_cf31 _dafny.Int = _6___mcc_h4
		_ = _7_cf31
		var _8_cf30 bool = _5___mcc_h3
		_ = _8_cf30
		var _9_cf29 bool = _4___mcc_h2
		_ = _9_cf29
		var _10_cf28 bool = _3___mcc_h1
		_ = _10_cf28
		var _11_cf27 _dafny.Int = _2___mcc_h0
		_ = _11_cf27
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(662))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _dafny.UnicodeSeqOfUtf8Bytes("clhe"), _dafny.UnicodeSeqOfUtf8Bytes("lh")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ilyfkpfb"), _dafny.UnicodeSeqOfUtf8Bytes("xvy"), _dafny.UnicodeSeqOfUtf8Bytes("tygu")))
	} else if _source0.Is_DC18() {
		var _13___mcc_h5 _dafny.Int = _source0.Get_().(D5_DC18).Cf32
		_ = _13___mcc_h5
		var _14___mcc_h6 _dafny.Int = _source0.Get_().(D5_DC18).Cf33
		_ = _14___mcc_h6
		var _15_cf33 _dafny.Int = _14___mcc_h6
		_ = _15_cf33
		var _16_cf32 _dafny.Int = _13___mcc_h5
		_ = _16_cf32
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wfng")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jovnaak"), _dafny.UnicodeSeqOfUtf8Bytes("gepycipjg")))
	} else {
		var _17___mcc_h7 _dafny.Sequence = _source0.Get_().(D5_DC16).Cf26
		_ = _17___mcc_h7
		var _18_cf26 _dafny.Sequence = _17___mcc_h7
		_ = _18_cf26
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_18_cf26, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-90))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_19_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))), _dafny.SeqOf(_18_cf26))
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	if (_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(2), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kdpoord")).Cardinality())))).IsSubsetOf(_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(806))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_20_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())
		}))).Cardinality()), _dafny.IntOfInt64(625))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _21_v0 _dafny.Int
			_21_v0 = interface{}(_compr_2).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(806))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_20_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())
			}))).Cardinality()), _dafny.IntOfInt64(625)), _21_v0) {
				_coll2.Add(Companion_Default___.SafeDivisionInt(_21_v0, _dafny.IntOfInt64(237)))
			}
		}
		return _coll2.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(547))))) {
		return _dafny.SetOf(_dafny.CodePoint('t'), _dafny.CodePoint('i'))
	} else {
		return (func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('e'))).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _22_v1 _dafny.CodePoint
					_22_v1 = interface{}(_compr_4).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('e')), _22_v1) {
						_coll4.Add(_22_v1)
					}
				}
				return _coll4.ToSet()
			}()).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _23_v2 _dafny.CodePoint
				_23_v2 = interface{}(_compr_3).(_dafny.CodePoint)
				if (func() _dafny.Set {
					var _coll5 = _dafny.NewBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('e'))).Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _24_v1 _dafny.CodePoint
						_24_v1 = interface{}(_compr_5).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('e')), _24_v1) {
							_coll5.Add(_24_v1)
						}
					}
					return _coll5.ToSet()
				}()).Contains(_23_v2) {
					_coll3.Add(_23_v2)
				}
			}
			return _coll3.ToSet()
		}()).Difference(_dafny.SetOf(_dafny.CodePoint('q')))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC3_((_dafny.MultiSetOf(true, true)).IsSubsetOf(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.IntOfInt64(-108)).Plus((_dafny.SetOf(false)).Cardinality()), (_dafny.IntOfUint32((_dafny.SeqOf(false, !(!(false)), false)).Cardinality())).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("domjx")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bgwkdeg")).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, p2 D0, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(!(false))).Intersection(_dafny.MultiSetOf(true))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) D0 {
	if true {
		if true {
			return Companion_D0_.Create_DC0_((func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('t'))).Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _25_v0 _dafny.CodePoint
					_25_v0 = interface{}(_compr_6).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('t')), _25_v0) {
						_coll6.Add(_25_v0)
					}
				}
				return _coll6.ToSet()
			}()).Cardinality())
		} else {
			return Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(589), false)).Cardinality())
		}
	} else if !(true) {
		return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqOf(false, true, true)).Cardinality()))
	} else {
		return Companion_D0_.Create_DC0_((_dafny.MultiSetOf(_dafny.IntOfInt64(-76), _dafny.IntOfInt64(-81))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(98))).Uint32(), func(coer4 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_26_i0 _dafny.Int) D0 {
		return Companion_D0_.Create_DC2_(false)
	})), _dafny.SeqOf(Companion_D0_.Create_DC2_(true), Companion_D0_.Create_DC2_(true), Companion_D0_.Create_DC2_(false), Companion_D0_.Create_DC2_(false))), _dafny.SeqOf(Companion_D0_.Create_DC2_(true)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(392))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()))).Cardinality())).Cardinality())).Cardinality(), _dafny.SeqOf(Companion_D5_.Create_DC18_(_dafny.IntOfInt64(20), _dafny.IntOfInt64(354)), Companion_D5_.Create_DC18_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), _dafny.IntOfInt64(87))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(748), _dafny.SeqOf(Companion_D5_.Create_DC18_(_dafny.IntOfInt64(73), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), Companion_D5_.Create_DC18_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dsqcv")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))).Cardinality())).Cardinality()), Companion_D5_.Create_DC18_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_28_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-780), true)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(354))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_29_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()), !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfInt64(339), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-818))).Uint32(), func(coer8 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_30_i1 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(false, false)
	}))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(740))))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) D2 {
	if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l'))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(141))).Cardinality()))).Cardinality())) > 0 {
		return Companion_D2_.Create_DC8_(_dafny.IntOfInt64(895), _dafny.IntOfInt64(72))
	} else {
		return Companion_D2_.Create_DC8_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), _dafny.CodePoint('b'))).Cardinality()), _dafny.IntOfInt64(-93))
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, true, true)).Difference(_dafny.SetOf(false, !(!(true))))).Difference(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(200), _dafny.IntOfInt64(666))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _31_v0 _dafny.Int
			_31_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(200)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(666)) < 0) {
				_coll7.Add(Companion_Default___.SafeDivisionInt(_31_v0, _dafny.IntOfInt64(261)))
			}
		}
		return _coll7.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfInt64(884)))).Union(_dafny.SetOf((_dafny.MultiSetOf(true, false, !(!(false)), true, false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Map, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC20_(_dafny.SeqOf(true), !(false) || (!(!(true))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 _dafny.Int, p2 _dafny.Set, p3 bool, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC25_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('x'))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvyxx")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.MultiSetOf(false))).Cardinality())).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(265), _dafny.IntOfInt64(982))))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kromkpsyu")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("le")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Map, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.MultiSetOf(!(false), true)).IsProperSubsetOf(_dafny.MultiSetOf(false, false)) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(true, true)), _dafny.SeqOf((Companion_D1_.Create_DC4_(_dafny.MultiSetOf(false))).Dtor_cf4(), _dafny.MultiSetOf(!(false), !(true)), _dafny.MultiSetOf(false), _dafny.MultiSetOf(false, false)))
	} else {
		return _dafny.SeqOf(_dafny.MultiSetOf(false, !(true)))
	}
}
func (_static *CompanionStruct_Default___) M2(p0 _dafny.Set, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, bool, bool, bool) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 bool = false
	_ = r2
	var r3 bool = false
	_ = r3
	(globalState).F3 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfInt64(102))), (_dafny.Zero).Minus(p2))
	var _32_v0 bool
	_ = _32_v0
	_32_v0 = false
	if _32_v0 {
		var _33_v1 _dafny.MultiSet
		_ = _33_v1
		_33_v1 = _dafny.MultiSetOf(_dafny.IntOfInt64(416), p2)
		var _34_v2 _dafny.Map
		_ = _34_v2
		_34_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_33_v1).Cardinality(), true)
		var _35_v3 _dafny.Sequence
		_ = _35_v3
		_35_v3 = _dafny.SeqOf((_34_v2).Cardinality(), (Companion_Default___.Fm16(globalState)).Cardinality())
		var _36_v4 _dafny.Map
		_ = _36_v4
		_36_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_35_v3).Cardinality()), _32_v0)
		r3 = (func() bool {
			if (_34_v2).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(12))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_37_p3 _dafny.Sequence, _38_p2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_39_i0 _dafny.Int) _dafny.CodePoint {
					return (_37_p3).Select((Companion_Default___.SafeIndex(_38_p2, _dafny.IntOfUint32((_37_p3).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(p3, p2)))).Cardinality())) {
				return (_34_v2).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(12))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_40_p3 _dafny.Sequence, _41_p2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_42_i0 _dafny.Int) _dafny.CodePoint {
						return (_40_p3).Select((Companion_Default___.SafeIndex(_41_p2, _dafny.IntOfUint32((_40_p3).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(p3, p2)))).Cardinality())).(bool)
			}
			return _32_v0
		})()
		var _43_v5 _dafny.Array
		_ = _43_v5
		var _nwElement0_0 _dafny.Int = p2
		_ = _nwElement0_0
		var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(13))
		_ = _nw0
		(_nw0).ArraySet1(_nwElement0_0, 0)
		(_nw0).ArraySet1(_dafny.IntOfInt64(-94), 1)
		(_nw0).ArraySet1(p2, 2)
		(_nw0).ArraySet1(_dafny.IntOfInt64(119), 3)
		(_nw0).ArraySet1(_dafny.IntOfInt64(30), 4)
		(_nw0).ArraySet1((_dafny.Zero).Minus(p2), 5)
		(_nw0).ArraySet1(p2, 6)
		(_nw0).ArraySet1((p2).Times(p2), 7)
		(_nw0).ArraySet1(p2, 8)
		(_nw0).ArraySet1((_dafny.Zero).Minus(p2), 9)
		(_nw0).ArraySet1(_dafny.IntOfInt64(340), 10)
		(_nw0).ArraySet1((_dafny.Zero).Minus(p2), 11)
		(_nw0).ArraySet1((_35_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_35_v3).Cardinality()))).Uint32()).(_dafny.Int), 12)
		_43_v5 = _nw0
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_43_v5), 0))
		_ = _index0
		(_43_v5).ArraySet1(p2, (_index0).Int())
		var _44_v6 _dafny.Array
		_ = _44_v6
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_45_v0 bool) func(_dafny.Int) bool {
				return func(_46_i1 _dafny.Int) bool {
					return (_dafny.SetOf(_45_v0)).Contains(_45_v0)
				}
			})(_32_v0)
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
		_44_v6 = _nw1
		var _47_v7 _dafny.Sequence
		_ = _47_v7
		_47_v7 = _dafny.SeqOf(_44_v6, _44_v6)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_43_v5), 0))
		_ = _index1
		var _rhs0 bool = !((_32_v0) && (_32_v0)) || (_32_v0)
		_ = _rhs0
		var _rhs1 _dafny.Int = p2
		_ = _rhs1
		var _rhs2 _dafny.Int = ((Companion_D3_.Create_DC11_(p2)).Dtor_cf14()).Times(p2)
		_ = _rhs2
		var _rhs3 _dafny.Array = _44_v6
		_ = _rhs3
		var _rhs4 _dafny.Array = (_47_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p3).Cardinality()), _dafny.IntOfUint32((_47_v7).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 _dafny.Array = _43_v5
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_43_v5), 0))
		_ = _lhs3
		_lhs0.F0 = _rhs0
		_lhs1.F3 = _rhs1
		(_lhs2).ArraySet1(_rhs2, (_lhs3).Int())
		_44_v6 = _rhs3
		_44_v6 = _rhs4
		(globalState).F1 = (_43_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_43_v5), 0))).Int()).(_dafny.Int)
		var _48_v8 _dafny.Array
		_ = _48_v8
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_1
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Sequence = (func(_49_v0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_50_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_49_v0, false, _49_v0, _49_v0)), _dafny.MultiSetOf(_49_v0)), _dafny.SeqOf(_dafny.MultiSetOf(_49_v0)))
				}
			})(_32_v0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw2 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw2).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw2).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_48_v8 = _nw2
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_48_v8), 0))
		_ = _index2
		(_48_v8).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(337))).Uint32(), func(coer11 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_51_v0 bool) func(_dafny.Int) _dafny.MultiSet {
			return func(_52_i3 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_51_v0, _51_v0, _51_v0)
			}
		})(_32_v0))), (_index2).Int())
		var _53_v9 _dafny.Map
		_ = _53_v9
		_53_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), p3)
		var _54_v10 _dafny.CodePoint
		_ = _54_v10
		_54_v10 = _dafny.CodePoint('j')
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_48_v8), 0))
		_ = _index3
		(_48_v8).ArraySet1(Companion_Default___.Fm23(_dafny.IntOfInt64(49), _dafny.CodePoint('y'), (_53_v9).Update(_54_v10, p3), p3, globalState), (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_43_v5), 0))
		_ = _index4
		(_43_v5).ArraySet1(p2, (_index4).Int())
	} else {
		var _55_v11 _dafny.Sequence
		_ = _55_v11
		_55_v11 = _dafny.SeqOf(p2)
		var _56_v12 _dafny.Map
		_ = _56_v12
		_56_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v11, p2)
		_56_v12 = (_56_v12).Update(_55_v11, p2)
		var _57_v13 D1
		_ = _57_v13
		_57_v13 = Companion_D1_.Create_DC5_(_32_v0)
		var _source1 D1 = _57_v13
		_ = _source1
		if _source1.Is_DC5() {
			var _58___mcc_h0 bool = _source1.Get_().(D1_DC5).Cf5
			_ = _58___mcc_h0
			var _59_cf5 bool = _58___mcc_h0
			_ = _59_cf5
			(globalState).F1 = p2
			var _60_v14 _dafny.Array
			_ = _60_v14
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw3
			_60_v14 = _nw3
			var _61_v15 _dafny.Sequence
			_ = _61_v15
			_61_v15 = _dafny.SeqOf(_60_v14)
			var _62_v16 *C0
			_ = _62_v16
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__((_61_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-78), _dafny.IntOfUint32((_61_v15).Cardinality()))).Uint32()).(_dafny.Array))
			_62_v16 = _nw4
			var _63_v17 _dafny.MultiSet
			_ = _63_v17
			_63_v17 = _dafny.MultiSetOf(p2)
			var _64_v18 D0
			_ = _64_v18
			_64_v18 = Companion_D0_.Create_DC3_(_32_v0)
			var _65_v19 _dafny.CodePoint
			_ = _65_v19
			_65_v19 = _dafny.CodePoint('k')
			var _66_v20 _dafny.Sequence
			_ = _66_v20
			_66_v20 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v19, _59_cf5), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v19, _59_cf5))
			var _67_v21 D3
			_ = _67_v21
			_67_v21 = Companion_D3_.Create_DC11_(p2)
			var _68_v22 _dafny.Sequence
			_ = _68_v22
			_68_v22 = _dafny.SeqOf(_67_v21, _67_v21)
			var _69_v23 _dafny.Array
			_ = _69_v23
			var _nwElement0_1 _dafny.Int = p2
			_ = _nwElement0_1
			var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
			_ = _nw5
			(_nw5).ArraySet1(_nwElement0_1, 0)
			(_nw5).ArraySet1(p2, 1)
			(_nw5).ArraySet1((_dafny.Zero).Minus(p2), 2)
			(_nw5).ArraySet1(p2, 3)
			(_nw5).ArraySet1(p2, 4)
			(_nw5).ArraySet1((_63_v17).Cardinality(), 5)
			(_nw5).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm11(_59_cf5, _32_v0, _64_v18, globalState)).Cardinality()), 6)
			(_nw5).ArraySet1(_dafny.IntOfInt64(473), 7)
			(_nw5).ArraySet1(p2, 8)
			(_nw5).ArraySet1((_dafny.Zero).Minus(p2), 9)
			(_nw5).ArraySet1(p2, 10)
			(_nw5).ArraySet1(p2, 11)
			(_nw5).ArraySet1(p2, 12)
			(_nw5).ArraySet1(p2, 13)
			(_nw5).ArraySet1(_dafny.IntOfInt64(-114), 14)
			(_nw5).ArraySet1(p2, 15)
			(_nw5).ArraySet1((_dafny.Zero).Minus(p2), 16)
			(_nw5).ArraySet1(((_66_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(p2, _dafny.IntOfInt64(697))).Cardinality()), _dafny.IntOfUint32((_66_v20).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 17)
			(_nw5).ArraySet1(p2, 18)
			(_nw5).ArraySet1(p2, 19)
			(_nw5).ArraySet1(_dafny.IntOfUint32((_68_v22).Cardinality()), 20)
			(_nw5).ArraySet1(p2, 21)
			(_nw5).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((p3).Cardinality()), globalState), 22)
			(_nw5).ArraySet1((_55_v11).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_55_v11).Cardinality()))).Uint32()).(_dafny.Int), 23)
			(_nw5).ArraySet1(p2, 24)
			(_nw5).ArraySet1(p2, 25)
			_69_v23 = _nw5
			var _70_v24 _dafny.Sequence
			_ = _70_v24
			_70_v24 = _dafny.SeqOf(_69_v23)
			_70_v24 = _dafny.Companion_Sequence_.Concatenate(_70_v24, _dafny.Companion_Sequence_.Concatenate(_70_v24, _70_v24))
			(globalState).F1 = (_dafny.Zero).Minus((_dafny.IntOfInt64(427)).Plus(p2))
		} else if _source1.Is_DC4() {
			var _71___mcc_h1 _dafny.MultiSet = _source1.Get_().(D1_DC4).Cf4
			_ = _71___mcc_h1
			var _72_cf4 _dafny.MultiSet = _71___mcc_h1
			_ = _72_cf4
			var _73_v25 _dafny.Array
			_ = _73_v25
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_2
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_74_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_75_i4 _dafny.Int) _dafny.Int {
						return (_75_i4).Minus((Companion_D6_.Create_DC22_(_74_p2, _74_p2)).Dtor_cf42())
					}
				})(p2)
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
			_73_v25 = _nw6
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))
			_ = _index5
			(_73_v25).ArraySet1(_dafny.IntOfInt64(183), (_index5).Int())
			var _76_v26 _dafny.Array
			_ = _76_v26
			var _nwElement0_2 bool = _32_v0
			_ = _nwElement0_2
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(27))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_2, 0)
			(_nw7).ArraySet1(_32_v0, 1)
			(_nw7).ArraySet1(_32_v0, 2)
			(_nw7).ArraySet1(false, 3)
			(_nw7).ArraySet1(_32_v0, 4)
			(_nw7).ArraySet1(_32_v0, 5)
			(_nw7).ArraySet1(_32_v0, 6)
			(_nw7).ArraySet1(!(_32_v0), 7)
			(_nw7).ArraySet1(_32_v0, 8)
			(_nw7).ArraySet1(_32_v0, 9)
			(_nw7).ArraySet1(_32_v0, 10)
			(_nw7).ArraySet1(_32_v0, 11)
			(_nw7).ArraySet1(_32_v0, 12)
			(_nw7).ArraySet1(_32_v0, 13)
			(_nw7).ArraySet1(_32_v0, 14)
			(_nw7).ArraySet1(_32_v0, 15)
			(_nw7).ArraySet1(_32_v0, 16)
			(_nw7).ArraySet1(_32_v0, 17)
			(_nw7).ArraySet1(_32_v0, 18)
			(_nw7).ArraySet1(_32_v0, 19)
			(_nw7).ArraySet1(!(_32_v0), 20)
			(_nw7).ArraySet1(false, 21)
			(_nw7).ArraySet1(_32_v0, 22)
			(_nw7).ArraySet1(_32_v0, 23)
			(_nw7).ArraySet1(_32_v0, 24)
			(_nw7).ArraySet1(_32_v0, 25)
			(_nw7).ArraySet1(_32_v0, 26)
			_76_v26 = _nw7
			var _77_v27 _dafny.Map
			_ = _77_v27
			_77_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v26, p3)
			var _78_v28 _dafny.Map
			_ = _78_v28
			_78_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v27, false)
			var _79_v29 _dafny.Map
			_ = _79_v29
			_79_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v11).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_55_v11).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((p3).Cardinality()))
			var _80_v30 _dafny.Map
			_ = _80_v30
			_80_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v29, _77_v27)
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))
			_ = _index6
			var _rhs5 _dafny.Int = (func() _dafny.Int {
				if _32_v0 {
					return p2
				}
				return ((p0).Cardinality()).Minus(p2)
			})()
			_ = _rhs5
			var _rhs6 _dafny.Int = p2
			_ = _rhs6
			var _rhs7 bool = !((func() bool {
				if (_78_v28).Contains((func() _dafny.Map {
					if (_80_v30).Contains(_79_v29) {
						return (_80_v30).Get(_79_v29).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v26, _dafny.UnicodeSeqOfUtf8Bytes("ft"))
				})()) {
					return (_78_v28).Get((func() _dafny.Map {
						if (_80_v30).Contains(_79_v29) {
							return (_80_v30).Get(_79_v29).(_dafny.Map)
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v26, _dafny.UnicodeSeqOfUtf8Bytes("ft"))
					})()).(bool)
				}
				return true
			})())
			_ = _rhs7
			var _rhs8 _dafny.Int = p2
			_ = _rhs8
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			var _lhs6 _dafny.Array = _73_v25
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))
			_ = _lhs7
			_lhs4.F1 = _rhs5
			_lhs5.F3 = _rhs6
			r3 = _rhs7
			(_lhs6).ArraySet1(_rhs8, (_lhs7).Int())
			var _81_v31 _dafny.Set
			_ = _81_v31
			_81_v31 = _dafny.SetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(347), Companion_Default___.Fm1(p2, globalState)), (_73_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))).Int()).(_dafny.Int), (_73_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))).Int()).(_dafny.Int))
			var _82_v32 _dafny.Set
			_ = _82_v32
			_82_v32 = _dafny.SetOf(_32_v0, _32_v0, _32_v0)
			var _rhs9 _dafny.Int = _dafny.IntOfInt64(362)
			_ = _rhs9
			var _rhs10 _dafny.Int = (_dafny.IntOfInt64(-108)).Minus(p2)
			_ = _rhs10
			var _rhs11 _dafny.Set = p0
			_ = _rhs11
			var _rhs12 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)).Cardinality(), (_dafny.Zero).Minus((_82_v32).Cardinality()))
			_ = _rhs12
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			r0 = _rhs9
			_lhs8.F3 = _rhs10
			_81_v31 = _rhs11
			r0 = _rhs12
			var _83_v33 _dafny.Map
			_ = _83_v33
			_83_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_dafny.Zero).Minus((_73_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))).Int()).(_dafny.Int)), globalState), _dafny.Companion_Sequence_.Concatenate(p3, p3))
			var _84_v34 _dafny.Sequence
			_ = _84_v34
			_84_v34 = _dafny.SeqOf(p3)
			_83_v33 = (_83_v33).Update((_dafny.Zero).Minus((_dafny.IntOfInt64(-273)).Minus((_dafny.Zero).Minus(p2))), (_84_v34).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_32_v0, _32_v0, _32_v0, _32_v0, Companion_Default___.Fm0((_73_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_73_v25), 0))).Int()).(_dafny.Int), globalState))).Cardinality(), _dafny.IntOfUint32((_84_v34).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_76_v26), 0))
			_ = _index7
			(_76_v26).ArraySet1(_32_v0, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_76_v26), 0))
			_ = _index8
			(_76_v26).ArraySet1(!(_32_v0), (_index8).Int())
		} else {
			var _85___mcc_h2 D1 = _source1.Get_().(D1_DC6).Cf6
			_ = _85___mcc_h2
			var _86_cf6 D1 = _85___mcc_h2
			_ = _86_cf6
			(globalState).F1 = (p2).Minus(p2)
			var _87_v35 _dafny.Set
			_ = _87_v35
			_87_v35 = _dafny.SetOf(p2, p2, p2)
			_87_v35 = _87_v35
			r3 = _32_v0
			(globalState).F1 = p2
		}
		var _88_v36 *C1
		_ = _88_v36
		var _nw8 *C1 = New_C1_()
		_ = _nw8
		_nw8.Ctor__()
		_88_v36 = _nw8
		var _89_v37 _dafny.CodePoint
		_ = _89_v37
		_89_v37 = _dafny.CodePoint('j')
		var _90_v38 _dafny.Map
		_ = _90_v38
		_90_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v36, _89_v37)
		_90_v38 = (_90_v38).Update(_88_v36, _89_v37)
		_88_v36 = _88_v36
		var _91_v39 *C1
		_ = _91_v39
		var _nw9 *C1 = New_C1_()
		_ = _nw9
		_nw9.Ctor__()
		_91_v39 = _nw9
	}
	var _92_v40 _dafny.Map
	_ = _92_v40
	_92_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
	var _93_v41 _dafny.Map
	_ = _93_v41
	_93_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, (p3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((p3).Cardinality()))).Uint32()).(_dafny.CodePoint))
	var _94_v42 _dafny.Sequence
	_ = _94_v42
	_94_v42 = _dafny.SeqOf(_32_v0)
	var _95_v43 _dafny.Set
	_ = _95_v43
	_95_v43 = _dafny.SetOf(false)
	var _96_v44 _dafny.Array
	_ = _96_v44
	var _nwElement0_3 _dafny.Set = _dafny.SetOf(_32_v0)
	_ = _nwElement0_3
	var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
	_ = _nw10
	(_nw10).ArraySet1(_nwElement0_3, 0)
	(_nw10).ArraySet1((_dafny.SetOf(!(Companion_Default___.Fm0(_dafny.IntOfUint32((p3).Cardinality()), globalState)), _32_v0, (func() bool {
		if (_92_v40).Contains((_93_v41).Cardinality()) {
			return (_92_v40).Get((_93_v41).Cardinality()).(bool)
		}
		return _32_v0
	})(), _32_v0)).Union(_dafny.SetOf(_32_v0, (_94_v42).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_94_v42).Cardinality()))).Uint32()).(bool), _32_v0)), 1)
	(_nw10).ArraySet1((_95_v43).Difference(_95_v43), 2)
	(_nw10).ArraySet1((_95_v43).Union(_95_v43), 3)
	(_nw10).ArraySet1(_dafny.SetOf(_32_v0, _32_v0, _32_v0), 4)
	(_nw10).ArraySet1(_95_v43, 5)
	(_nw10).ArraySet1(_95_v43, 6)
	(_nw10).ArraySet1((_dafny.SetOf(_32_v0, _32_v0)).Intersection(_95_v43), 7)
	(_nw10).ArraySet1(Companion_Default___.Fm18(p2, globalState), 8)
	(_nw10).ArraySet1((_dafny.SetOf(_32_v0)).Union(_95_v43), 9)
	(_nw10).ArraySet1(_95_v43, 10)
	_96_v44 = _nw10
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_96_v44), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _97_i5 _dafny.Int
		_97_i5 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_97_i5).Sign() != -1) && ((_97_i5).Cmp(_dafny.ArrayLen((_96_v44), 0)) < 0)) {
			(_96_v44).ArraySet1(_dafny.SetOf(_32_v0), (_97_i5).Int())
		}
	}
	var _98_v45 _dafny.CodePoint
	_ = _98_v45
	_98_v45 = _dafny.CodePoint('i')
	var _99_v46 D5
	_ = _99_v46
	_99_v46 = Companion_D5_.Create_DC16_(_dafny.Companion_Sequence_.Update((Companion_D6_.Create_DC21_(p2, _32_v0, p3, _32_v0)).Dtor_cf39(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p3).Cardinality()), _dafny.IntOfUint32(((Companion_D6_.Create_DC21_(p2, _32_v0, p3, _32_v0)).Dtor_cf39()).Cardinality()))).Uint32(), _98_v45))
	_99_v46 = _99_v46
	(globalState).F3 = Companion_Default___.Fm1((p2).Minus(p2), globalState)
	var _100_i6 _dafny.Int
	_ = _100_i6
	_100_i6 = _dafny.Zero
	{
		for Companion_Default___.Fm0(p2, globalState) {
			{
				if (_100_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_100_i6 = (_100_i6).Plus(_dafny.One)
				var _101_v47 _dafny.Sequence
				_ = _101_v47
				_101_v47 = _dafny.SeqOf(p2, _dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality()), p2)
				var _102_v48 D5
				_ = _102_v48
				_102_v48 = Companion_D5_.Create_DC17_((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_101_v47, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_101_v47).Cardinality()))).Uint32(), (_95_v43).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(342))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}((func(_103_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_104_i7 _dafny.Int) _dafny.Int {
						return _103_p2
					}
				})(p2)))))).Cardinality(), _32_v0, (func() bool {
					if _32_v0 {
						return !(_32_v0)
					}
					return _32_v0
				})(), (p2).Cmp(p2) == 0, (p2).Times(p2))
				var _105_v49 _dafny.Map
				_ = _105_v49
				_105_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_101_v47).Cardinality()))
				var _rhs13 _dafny.Int = p2
				_ = _rhs13
				var _rhs14 D5 = _102_v48
				_ = _rhs14
				var _rhs15 _dafny.Map = _105_v49
				_ = _rhs15
				var _lhs9 *GlobalState = globalState
				_ = _lhs9
				_lhs9.F3 = _rhs13
				_102_v48 = _rhs14
				_105_v49 = _rhs15
				if _32_v0 {
					var _106_v50 _dafny.Array
					_ = _106_v50
					var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
					_ = _nw11
					_106_v50 = _nw11
					_106_v50 = _106_v50
					var _107_v51 *C1
					_ = _107_v51
					var _nw12 *C1 = New_C1_()
					_ = _nw12
					_nw12.Ctor__()
					_107_v51 = _nw12
					var _108_v52 _dafny.Set
					_ = _108_v52
					_108_v52 = _dafny.SetOf(_107_v51, _107_v51)
					_101_v47 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(16))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}((func(_109_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_110_i8 _dafny.Int) _dafny.Int {
							return _109_p2
						}
					})(p2))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_108_v52).Cardinality()), _101_v47))
					var _111_v53 _dafny.Map
					_ = _111_v53
					_111_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v51, p2)
					_111_v53 = (_111_v53).Update(_107_v51, _dafny.IntOfUint32((p3).Cardinality()))
					(globalState).F1 = (p2).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg14 _dafny.Int) interface{} {
							return coer14(arg14)
						}
					}(func(_112_i9 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality()))
					r3 = _32_v0
				} else {
					var _113_v54 _dafny.Map
					_ = _113_v54
					_113_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-486), _94_v42)
					var _114_v55 D6
					_ = _114_v55
					_114_v55 = Companion_D6_.Create_DC20_(_94_v42, _32_v0)
					var _115_v56 _dafny.Sequence
					_ = _115_v56
					_115_v56 = _dafny.SeqOf(_dafny.SeqOf((func() _dafny.Sequence {
						if (_113_v54).Contains(p2) {
							return (_113_v54).Get(p2).(_dafny.Sequence)
						}
						return _94_v42
					})(), _94_v42, (_114_v55).Dtor_cf35()))
					(globalState).F1 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_115_v56).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_115_v56).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), _98_v45)).Cardinality()))
					var _116_v57 *C1
					_ = _116_v57
					var _nw13 *C1 = New_C1_()
					_ = _nw13
					_nw13.Ctor__()
					_116_v57 = _nw13
					var _117_v58 _dafny.Array
					_ = _117_v58
					var _len0_3 _dafny.Int = _dafny.IntOfInt64(15)
					_ = _len0_3
					var _nw14 _dafny.Array
					_ = _nw14
					if _len0_3.Cmp(_dafny.Zero) == 0 {
						_nw14 = _dafny.NewArray(_len0_3)
					} else {
						var _init3 func(_dafny.Int) bool = (func(_118_v0 bool) func(_dafny.Int) bool {
							return func(_119_i10 _dafny.Int) bool {
								return _118_v0
							}
						})(_32_v0)
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
					_117_v58 = _nw14
					var _120_v59 _dafny.Set
					_ = _120_v59
					_120_v59 = _dafny.SetOf(_117_v58)
					var _121_v60 _dafny.Map
					_ = _121_v60
					_121_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, _dafny.IntOfInt64(-399))
					var _122_v61 D6
					_ = _122_v61
					_122_v61 = Companion_D6_.Create_DC22_(p2, p2)
					var _rhs16 bool = _32_v0
					_ = _rhs16
					var _rhs17 _dafny.CodePoint = (p3).Select((Companion_Default___.SafeIndex(((_120_v59).Union(_120_v59)).Cardinality(), _dafny.IntOfUint32((p3).Cardinality()))).Uint32()).(_dafny.CodePoint)
					_ = _rhs17
					var _rhs18 bool = ((_dafny.IntOfInt64(649)).Plus((func() _dafny.Int {
						if (_121_v60).Contains(false) {
							return (_121_v60).Get(false).(_dafny.Int)
						}
						return p2
					})())).Cmp((_122_v61).Dtor_cf42()) < 0
					_ = _rhs18
					r1 = _rhs16
					_98_v45 = _rhs17
					r1 = _rhs18
					var _123_v62 *C1
					_ = _123_v62
					var _nw15 *C1 = New_C1_()
					_ = _nw15
					_nw15.Ctor__()
					_123_v62 = _nw15
					var _124_v63 _dafny.Array
					_ = _124_v63
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_4
					var _nw16 _dafny.Array
					_ = _nw16
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw16 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) _dafny.Sequence = (func(_125_p3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_126_i11 _dafny.Int) _dafny.Sequence {
								return _125_p3
							}
						})(p3)
						_ = _init4
						var _element0_4 = _init4(_dafny.Zero)
						_ = _element0_4
						_nw16 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
						(_nw16).ArraySet1(_element0_4, 0)
						var _nativeLen0_4 = (_len0_4).Int()
						_ = _nativeLen0_4
						for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
							(_nw16).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
						}
					}
					_124_v63 = _nw16
					var _127_v64 *C0
					_ = _127_v64
					var _nw17 *C0 = New_C0_()
					_ = _nw17
					_nw17.Ctor__(_124_v63)
					_127_v64 = _nw17
				}
				(globalState).F1 = Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ywah")).Cardinality()), globalState)
				if (_95_v43).IsProperSubsetOf((Companion_Default___.Fm18(p2, globalState)).Intersection(_95_v43)) {
					r2 = (_dafny.IntOfInt64(5)).Cmp(_dafny.IntOfInt64(-131)) > 0
					var _128_v65 *C1
					_ = _128_v65
					var _nw18 *C1 = New_C1_()
					_ = _nw18
					_nw18.Ctor__()
					_128_v65 = _nw18
					var _129_v66 _dafny.Array
					_ = _129_v66
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_5
					var _nw19 _dafny.Array
					_ = _nw19
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw19 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) bool = (func(_130_v47 _dafny.Sequence, _131_p2 _dafny.Int, _132_p3 _dafny.Sequence, _133_v45 _dafny.CodePoint) func(_dafny.Int) bool {
							return func(_134_i12 _dafny.Int) bool {
								return (_dafny.MultiSetOf(_dafny.IntOfInt64(448), _131_p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_132_p3, (Companion_Default___.SafeIndex(_131_p2, _dafny.IntOfUint32((_132_p3).Cardinality()))).Uint32(), _133_v45)).Cardinality()), _131_p2, _131_p2)).IsSubsetOf(_dafny.MultiSetFromSeq(_130_v47))
							}
						})(_101_v47, p2, p3, _98_v45)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw19 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw19).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw19).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_129_v66 = _nw19
					var _135_v67 _dafny.MultiSet
					_ = _135_v67
					_135_v67 = _dafny.MultiSetOf(p3, _dafny.Companion_Sequence_.Concatenate(p3, Companion_Default___.Fm4(true, p2, p2, _dafny.CodePoint('e'), globalState)))
					var _136_v68 _dafny.Array
					_ = _136_v68
					var _nw20 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
					_ = _nw20
					_136_v68 = _nw20
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_136_v68), 0))
					_ = _index9
					(_136_v68).ArraySet1(_128_v65, (_index9).Int())
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_136_v68), 0))
					_ = _index10
					var _rhs19 _dafny.Array = _129_v66
					_ = _rhs19
					var _rhs20 _dafny.MultiSet = _135_v67
					_ = _rhs20
					var _rhs21 *C1 = _128_v65
					_ = _rhs21
					var _lhs10 _dafny.Array = _136_v68
					_ = _lhs10
					var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_136_v68), 0))
					_ = _lhs11
					_129_v66 = _rhs19
					_135_v67 = _rhs20
					(_lhs10).ArraySet1(_rhs21, (_lhs11).Int())
					(globalState).F0 = _dafny.Companion_Sequence_.IsPrefixOf(p3, _dafny.Companion_Sequence_.Concatenate(p3, _dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), _98_v45)))
					var _137_v69 _dafny.Array
					_ = _137_v69
					var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
					_ = _nw21
					_137_v69 = _nw21
					var _138_v70 *C0
					_ = _138_v70
					var _nw22 *C0 = New_C0_()
					_ = _nw22
					_nw22.Ctor__(_137_v69)
					_138_v70 = _nw22
					var _139_v71 _dafny.MultiSet
					_ = _139_v71
					_139_v71 = _dafny.MultiSetOf(p2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_138_v70), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf(_138_v70)).Cardinality()))).Uint32(), _138_v70)).Cardinality())), p2, p2)
					var _140_v72 D0
					_ = _140_v72
					_140_v72 = Companion_D0_.Create_DC3_(_32_v0)
					var _141_v73 _dafny.Int
					_ = _141_v73
					var _out0 _dafny.Int
					_ = _out0
					_out0 = (_128_v65).M1((Companion_D0_.Create_DC3_(_32_v0)).Dtor_cf3(), _dafny.SetOf((_128_v65).Fm2(_139_v71, globalState)), _dafny.Companion_Sequence_.Concatenate(p3, p3), _140_v72, globalState)
					_141_v73 = _out0
				} else {
					var _142_v74 _dafny.Array
					_ = _142_v74
					var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
					_ = _nw23
					_142_v74 = _nw23
					var _143_v75 _dafny.Map
					_ = _143_v75
					_143_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v74, (_dafny.IntOfInt64(236)).Minus(p2))
					_143_v75 = (_143_v75).Update(_142_v74, p2)
					var _144_v76 _dafny.Array
					_ = _144_v76
					var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
					_ = _nw24
					_144_v76 = _nw24
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_144_v76), 0))
					_ = _index11
					(_144_v76).ArraySet1(p2, (_index11).Int())
					var _145_v77 D1
					_ = _145_v77
					_145_v77 = Companion_D1_.Create_DC5_(!(!(_32_v0)))
					var _146_v78 _dafny.Array
					_ = _146_v78
					var _nw25 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
					_ = _nw25
					_146_v78 = _nw25
					var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_146_v78), 0))
					_ = _index12
					(_146_v78).ArraySet1(_32_v0, (_index12).Int())
					var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_144_v76), 0))
					_ = _index13
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_146_v78), 0))
					_ = _index14
					var _rhs22 _dafny.Int = p2
					_ = _rhs22
					var _rhs23 _dafny.Int = (func() _dafny.Int {
						if _32_v0 {
							return _dafny.IntOfInt64(273)
						}
						return _dafny.IntOfInt64(-363)
					})()
					_ = _rhs23
					var _rhs24 D1 = _145_v77
					_ = _rhs24
					var _rhs25 bool = !((p2).Cmp(_dafny.IntOfInt64(375)) < 0)
					_ = _rhs25
					var _lhs12 _dafny.Array = _144_v76
					_ = _lhs12
					var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_144_v76), 0))
					_ = _lhs13
					var _lhs14 *GlobalState = globalState
					_ = _lhs14
					var _lhs15 _dafny.Array = _146_v78
					_ = _lhs15
					var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_146_v78), 0))
					_ = _lhs16
					(_lhs12).ArraySet1(_rhs22, (_lhs13).Int())
					_lhs14.F3 = _rhs23
					_145_v77 = _rhs24
					(_lhs15).ArraySet1(_rhs25, (_lhs16).Int())
					var _147_v79 D5
					_ = _147_v79
					_147_v79 = Companion_D5_.Create_DC18_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(994))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg15 _dafny.Int) interface{} {
							return coer15(arg15)
						}
					}((func(_148_v45 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_149_i13 _dafny.Int) _dafny.CodePoint {
							return _148_v45
						}
					})(_98_v45)))).Cardinality()), _dafny.IntOfUint32((p3).Cardinality()))
					var _150_v80 _dafny.Map
					_ = _150_v80
					_150_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _147_v79)
					_150_v80 = (_150_v80).Update(p3, _147_v79)
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_144_v76), 0))
					_ = _index15
					(_144_v76).ArraySet1(Companion_Default___.SafeModuloInt(p2, (func() _dafny.Int {
						if _32_v0 {
							return (_144_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_144_v76), 0))).Int()).(_dafny.Int)
						}
						return p2
					})()), (_index15).Int())
					var _151_v81 *C0
					_ = _151_v81
					var _nw26 *C0 = New_C0_()
					_ = _nw26
					_nw26.Ctor__(_142_v74)
					_151_v81 = _nw26
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _152_v82 *C1
	_ = _152_v82
	var _nw27 *C1 = New_C1_()
	_ = _nw27
	_nw27.Ctor__()
	_152_v82 = _nw27
	var _153_v83 _dafny.Set
	_ = _153_v83
	_153_v83 = _dafny.SetOf(_152_v82, _152_v82)
	r0 = (_153_v83).Cardinality()
	var _154_v84 _dafny.Sequence
	_ = _154_v84
	_154_v84 = _dafny.SeqOf(_92_v40)
	var _155_v85 _dafny.Sequence
	_ = _155_v85
	_155_v85 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex(((_154_v84).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.IntOfUint32((_154_v84).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), _98_v45)).Cardinality()), p2)
	var _156_v86 _dafny.Set
	_ = _156_v86
	_156_v86 = _dafny.SetOf(_155_v85)
	var _157_v87 _dafny.Sequence
	_ = _157_v87
	_157_v87 = _dafny.SeqOf(_155_v85)
	r1 = ((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(73))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}((func(_158_v41 _dafny.Map) func(_dafny.Int) _dafny.Int {
		return func(_159_i14 _dafny.Int) _dafny.Int {
			return (_158_v41).Cardinality()
		}
	})(_93_v41))))).Difference(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter9 := _dafny.Iterate((_157_v87).Elements()); ; {
			_compr_8, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _160_v88 _dafny.Sequence
			_160_v88 = interface{}(_compr_8).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_157_v87, _160_v88) {
				_coll8.Add(_160_v88)
			}
		}
		return _coll8.ToSet()
	}())).IsSubsetOf(_156_v86)
	r2 = _32_v0
	var _161_v89 _dafny.Array
	_ = _161_v89
	var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
	_ = _nw28
	_161_v89 = _nw28
	var _162_v90 *C0
	_ = _162_v90
	var _nw29 *C0 = New_C0_()
	_ = _nw29
	_nw29.Ctor__(_161_v89)
	_162_v90 = _nw29
	var _163_v91 _dafny.Sequence
	_ = _163_v91
	_163_v91 = _dafny.SeqOf(_162_v90)
	r3 = Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfUint32((_163_v91).Cardinality())), globalState)
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _164_v0 _dafny.Array
	_ = _164_v0
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(13)
	_ = _len0_6
	var _nw30 _dafny.Array
	_ = _nw30
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw30 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) _dafny.MultiSet = func(_165_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(false, true))
		}
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw30 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw30).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw30).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_164_v0 = _nw30
	var _166_v1 bool
	_ = _166_v1
	_166_v1 = false
	var _167_v2 _dafny.Sequence
	_ = _167_v2
	_167_v2 = _dafny.SeqOf(_166_v1, !(_166_v1))
	var _168_globalState *GlobalState
	_ = _168_globalState
	var _nw31 *GlobalState = New_GlobalState_()
	_ = _nw31
	_nw31.Ctor__(true, _dafny.IntOfInt64(827), true, _dafny.IntOfInt64(-651), _164_v0, _dafny.Companion_Sequence_.Concatenate(_167_v2, _167_v2))
	_168_globalState = _nw31
	var _169_v3 _dafny.Int
	_ = _169_v3
	_169_v3 = _dafny.IntOfInt64(272)
	var _170_v4 _dafny.Map
	_ = _170_v4
	_170_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v3, _169_v3)
	var _hi0 _dafny.Int = (_169_v3).Minus((func() _dafny.Int {
		if (_170_v4).Contains((_dafny.Zero).Minus(_169_v3)) {
			return (_170_v4).Get((_dafny.Zero).Minus(_169_v3)).(_dafny.Int)
		}
		return _169_v3
	})())
	_ = _hi0
	for _171_i1 := _169_v3; _171_i1.Cmp(_hi0) < 0; _171_i1 = _171_i1.Plus(_dafny.One) {
		var _172_v5 _dafny.CodePoint
		_ = _172_v5
		_172_v5 = _dafny.CodePoint('k')
		var _173_v6 _dafny.MultiSet
		_ = _173_v6
		_173_v6 = _dafny.MultiSetOf(_172_v5, _172_v5)
		var _174_v7 _dafny.Sequence
		_ = _174_v7
		_174_v7 = _dafny.SeqOf(_173_v6, _173_v6, _dafny.MultiSetOf(_dafny.CodePoint('x')))
		var _175_v8 D0
		_ = _175_v8
		_175_v8 = Companion_D0_.Create_DC0_(_169_v3)
		_174_v7 = _dafny.Companion_Sequence_.Concatenate(_174_v7, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer17 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_176_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.MultiSet {
			return func(_177_i2 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_176_v5, _176_v5)
			}
		})(_172_v5))), (Companion_Default___.SafeIndex((_175_v8).Dtor_cf0(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer18 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_178_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.MultiSet {
			return func(_179_i2 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_178_v5, _178_v5)
			}
		})(_172_v5)))).Cardinality()))).Uint32(), _173_v6))
		_170_v4 = (_170_v4).Update(_171_i1, (_171_i1).Minus((_dafny.Zero).Minus(_171_i1)))
		var _180_v9 _dafny.Set
		_ = _180_v9
		_180_v9 = _dafny.SetOf(Companion_Default___.Fm0(_dafny.IntOfUint32((_167_v2).Cardinality()), _168_globalState), _166_v1)
		var _181_v10 _dafny.MultiSet
		_ = _181_v10
		_181_v10 = _dafny.MultiSetOf(_180_v9, _180_v9)
		var _182_v11 D0
		_ = _182_v11
		_182_v11 = Companion_D0_.Create_DC2_((_181_v10).IsProperSubsetOf(_181_v10))
		var _183_v12 _dafny.Sequence
		_ = _183_v12
		_183_v12 = _dafny.UnicodeSeqOfUtf8Bytes("g")
		_182_v11 = (func() D0 {
			if (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_183_v12, (Companion_Default___.SafeIndex(_169_v3, _dafny.IntOfUint32((_183_v12).Cardinality()))).Uint32(), _172_v5)).Cardinality())).Cmp(_169_v3) != 0 {
				return _182_v11
			}
			return _182_v11
		})()
		(_168_globalState).F3 = _171_i1
	}
	var _184_v13 _dafny.Map
	_ = _184_v13
	_184_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v1, (_dafny.Zero).Minus(_169_v3))
	var _185_v14 D0
	_ = _185_v14
	_185_v14 = Companion_D0_.Create_DC2_(true)
	_184_v13 = (_184_v13).Update((_185_v14).Dtor_cf2(), _169_v3)
	_166_v1 = _166_v1
	var _186_v15 _dafny.Array
	_ = _186_v15
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(20)
	_ = _len0_7
	var _nw32 _dafny.Array
	_ = _nw32
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw32 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Int = (func(_187_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_188_i3 _dafny.Int) _dafny.Int {
				return (_188_i3).Plus((_dafny.SetOf(_187_v3)).Cardinality())
			}
		})(_169_v3)
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw32 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw32).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw32).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_186_v15 = _nw32
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
	_ = _index16
	(_186_v15).ArraySet1(_169_v3, (_index16).Int())
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
	_ = _index17
	(_186_v15).ArraySet1(_169_v3, (_index17).Int())
	var _189_v16 _dafny.Map
	_ = _189_v16
	_189_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v15, _dafny.IntOfInt64(-954))
	var _hi1 _dafny.Int = (_189_v16).Cardinality()
	_ = _hi1
	for _190_i4 := Companion_Default___.Fm1(_169_v3, _168_globalState); _190_i4.Cmp(_hi1) < 0; _190_i4 = _190_i4.Plus(_dafny.One) {
		var _191_v17 _dafny.Sequence
		_ = _191_v17
		_191_v17 = _dafny.UnicodeSeqOfUtf8Bytes("ldovmaed")
		var _192_v18 _dafny.Map
		_ = _192_v18
		_192_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v3, _dafny.Companion_Sequence_.Concatenate(_191_v17, _191_v17))
		_192_v18 = (_192_v18).Update(Companion_Default___.SafeDivisionInt(_169_v3, Companion_Default___.Fm1(_169_v3, _168_globalState)), _191_v17)
		var _193_v19 _dafny.Map
		_ = _193_v19
		_193_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_170_v4).Cardinality(), _166_v1)
		var _194_v20 _dafny.Set
		_ = _194_v20
		_194_v20 = _dafny.SetOf(_166_v1, _166_v1)
		_166_v1 = (func() bool {
			if (_193_v19).Contains(_190_i4) {
				return (_193_v19).Get(_190_i4).(bool)
			}
			return (_194_v20).IsProperSubsetOf(_194_v20)
		})()
		if false {
			var _195_v21 *C1
			_ = _195_v21
			var _nw33 *C1 = New_C1_()
			_ = _nw33
			_nw33.Ctor__()
			_195_v21 = _nw33
			var _196_v22 _dafny.MultiSet
			_ = _196_v22
			_196_v22 = _dafny.MultiSetOf((func() _dafny.Sequence {
				if (_192_v18).Contains((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)) {
					return (_192_v18).Get((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _191_v17
			})())
			var _197_v23 _dafny.Array
			_ = _197_v23
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw34
			_197_v23 = _nw34
			var _198_v24 D3
			_ = _198_v24
			_198_v24 = Companion_D3_.Create_DC12_(_169_v3, _196_v22, _dafny.IntOfUint32((_167_v2).Cardinality()), _197_v23)
			var _199_v25 _dafny.MultiSet
			_ = _199_v25
			_199_v25 = _dafny.MultiSetOf(_198_v24, _198_v24)
			var _200_v26 _dafny.Array
			_ = _200_v26
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw35
			_200_v26 = _nw35
			var _201_v27 _dafny.Map
			_ = _201_v27
			_201_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v3, _200_v26)
			var _202_v28 _dafny.Map
			_ = _202_v28
			_202_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_v25, (func() _dafny.Array {
				if (_201_v27).Contains(_169_v3) {
					return (_201_v27).Get(_169_v3).(_dafny.Array)
				}
				return _200_v26
			})())
			_202_v28 = (_202_v28).Update(_199_v25, _200_v26)
			var _203_v29 _dafny.Map
			_ = _203_v29
			_203_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)).Minus(_169_v3), (func() _dafny.Array {
				if _166_v1 {
					return _186_v15
				}
				return _186_v15
			})())
			_186_v15 = (func() _dafny.Array {
				if (_203_v29).Contains(_190_i4) {
					return (_203_v29).Get(_190_i4).(_dafny.Array)
				}
				return _186_v15
			})()
			var _204_v30 _dafny.Int
			_ = _204_v30
			var _205_v31 _dafny.Array
			_ = _205_v31
			var _206_v32 D0
			_ = _206_v32
			var _out1 _dafny.Int
			_ = _out1
			var _out2 _dafny.Array
			_ = _out2
			var _out3 D0
			_ = _out3
			_out1, _out2, _out3 = (_195_v21).M0(_168_globalState)
			_204_v30 = _out1
			_205_v31 = _out2
			_206_v32 = _out3
			var _207_v33 _dafny.Array
			_ = _207_v33
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_8
			var _nw36 _dafny.Array
			_ = _nw36
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw36 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = (func(_208_v17 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_209_i5 _dafny.Int) _dafny.Sequence {
						return _208_v17
					}
				})(_191_v17)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw36 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw36).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw36).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_207_v33 = _nw36
			var _210_v34 *C0
			_ = _210_v34
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__(_207_v33)
			_210_v34 = _nw37
			var _211_v35 _dafny.Array
			_ = _211_v35
			var _nwElement0_4 *C0 = _210_v34
			_ = _nwElement0_4
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_4, 0)
			(_nw38).ArraySet1(_210_v34, 1)
			(_nw38).ArraySet1(_210_v34, 2)
			(_nw38).ArraySet1(_210_v34, 3)
			(_nw38).ArraySet1(_210_v34, 4)
			(_nw38).ArraySet1(_210_v34, 5)
			(_nw38).ArraySet1(_210_v34, 6)
			(_nw38).ArraySet1(_210_v34, 7)
			(_nw38).ArraySet1(_210_v34, 8)
			(_nw38).ArraySet1(_210_v34, 9)
			(_nw38).ArraySet1(_210_v34, 10)
			(_nw38).ArraySet1(_210_v34, 11)
			(_nw38).ArraySet1(_210_v34, 12)
			(_nw38).ArraySet1(_210_v34, 13)
			(_nw38).ArraySet1(_210_v34, 14)
			(_nw38).ArraySet1(_210_v34, 15)
			(_nw38).ArraySet1(_210_v34, 16)
			(_nw38).ArraySet1(_210_v34, 17)
			(_nw38).ArraySet1((func() *C0 {
				if _166_v1 {
					return _210_v34
				}
				return _210_v34
			})(), 18)
			(_nw38).ArraySet1(_210_v34, 19)
			(_nw38).ArraySet1(_210_v34, 20)
			(_nw38).ArraySet1(_210_v34, 21)
			(_nw38).ArraySet1(_210_v34, 22)
			(_nw38).ArraySet1(_210_v34, 23)
			(_nw38).ArraySet1(_210_v34, 24)
			(_nw38).ArraySet1(_210_v34, 25)
			(_nw38).ArraySet1(_210_v34, 26)
			(_nw38).ArraySet1(_210_v34, 27)
			_211_v35 = _nw38
			_211_v35 = _211_v35
		} else {
			(_168_globalState).F0 = !(!(_166_v1) || ((func() bool {
				if true {
					return _166_v1
				}
				return _166_v1
			})()))
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
			_ = _index18
			(_186_v15).ArraySet1(((_193_v19).Merge(_193_v19)).Cardinality(), (_index18).Int())
			_166_v1 = !(_166_v1) || (_dafny.Companion_Sequence_.IsPrefixOf(_167_v2, _167_v2))
			(_168_globalState).F0 = Companion_Default___.Fm0((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _168_globalState)
			var _212_v36 _dafny.Array
			_ = _212_v36
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_9
			var _nw39 _dafny.Array
			_ = _nw39
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw39 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) bool = (func(_213_v1 bool) func(_dafny.Int) bool {
					return func(_214_i6 _dafny.Int) bool {
						return _213_v1
					}
				})(_166_v1)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw39 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw39).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw39).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_212_v36 = _nw39
			_212_v36 = _212_v36
		}
		var _215_v37 _dafny.Sequence
		_ = _215_v37
		_215_v37 = _dafny.SeqOf(_170_v4, (_170_v4).Update(_dafny.IntOfUint32((_191_v17).Cardinality()), _190_i4))
		var _216_v38 D3
		_ = _216_v38
		_216_v38 = Companion_D3_.Create_DC9_(_215_v37)
		var _217_v39 _dafny.Sequence
		_ = _217_v39
		_217_v39 = _dafny.SeqOf(_216_v38)
		var _218_v40 _dafny.Map
		_ = _218_v40
		_218_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v39, _191_v17)
		var _rhs26 _dafny.Map = _218_v40
		_ = _rhs26
		var _rhs27 _dafny.Array = _186_v15
		_ = _rhs27
		_218_v40 = _rhs26
		_186_v15 = _rhs27
	}
	var _219_v41 _dafny.Array
	_ = _219_v41
	var _nw40 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
	_ = _nw40
	_219_v41 = _nw40
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))
	_ = _index19
	(_219_v41).ArraySet1(Companion_Default___.Fm0((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _168_globalState), (_index19).Int())
	var _220_v42 _dafny.MultiSet
	_ = _220_v42
	_220_v42 = _dafny.MultiSetOf(_166_v1)
	var _pat_let_tv0 = _169_v3
	_ = _pat_let_tv0
	var _pat_let_tv1 = _186_v15
	_ = _pat_let_tv1
	var _pat_let_tv2 = _186_v15
	_ = _pat_let_tv2
	var _pat_let_tv3 = _220_v42
	_ = _pat_let_tv3
	var _pat_let_tv4 = _169_v3
	_ = _pat_let_tv4
	var _pat_let_tv5 = _166_v1
	_ = _pat_let_tv5
	var _pat_let_tv6 = _169_v3
	_ = _pat_let_tv6
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))
	_ = _index20
	(_219_v41).ArraySet1(func(_source2 D1) bool {
		if _source2.Is_DC5() {
			var _221___mcc_h0 bool = _source2.Get_().(D1_DC5).Cf5
			_ = _221___mcc_h0
			var _222_cf5 bool = _221___mcc_h0
			_ = _222_cf5
			return (_dafny.MultiSetOf(_pat_let_tv0, (_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(_dafny.Int), (_pat_let_tv3).Cardinality())).IsDisjointFrom(_dafny.MultiSetOf(_pat_let_tv4))
		} else if _source2.Is_DC4() {
			var _223___mcc_h1 _dafny.MultiSet = _source2.Get_().(D1_DC4).Cf4
			_ = _223___mcc_h1
			var _224_cf4 _dafny.MultiSet = _223___mcc_h1
			_ = _224_cf4
			return _pat_let_tv5
		} else {
			var _225___mcc_h2 D1 = _source2.Get_().(D1_DC6).Cf6
			_ = _225___mcc_h2
			var _226_cf6 D1 = _225___mcc_h2
			_ = _226_cf6
			return (_dafny.IntOfInt64(371)).Cmp(_pat_let_tv6) <= 0
		}
	}(Companion_D1_.Create_DC4_(_220_v42)), (_index20).Int())
	var _227_i7 _dafny.Int
	_ = _227_i7
	_227_i7 = _dafny.Zero
	{
		for (_169_v3).Cmp(_169_v3) < 0 {
			{
				if (_227_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_227_i7 = (_227_i7).Plus(_dafny.One)
				var _228_v43 _dafny.Array
				_ = _228_v43
				var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw41
				_228_v43 = _nw41
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_228_v43), 0))
				_ = _index21
				(_228_v43).ArraySet1(_219_v41, (_index21).Int())
				var _229_v44 _dafny.CodePoint
				_ = _229_v44
				_229_v44 = _dafny.CodePoint('w')
				var _230_v45 _dafny.MultiSet
				_ = _230_v45
				_230_v45 = _dafny.MultiSetOf(_229_v44, _229_v44, _229_v44, _229_v44)
				var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_228_v43), 0))
				_ = _index22
				var _nwElement0_5 bool = (func() bool {
					if _166_v1 {
						return (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)
					}
					return (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)
				})()
				_ = _nwElement0_5
				var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(12))
				_ = _nw42
				(_nw42).ArraySet1(_nwElement0_5, 0)
				(_nw42).ArraySet1(_166_v1, 1)
				(_nw42).ArraySet1((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), 2)
				(_nw42).ArraySet1(_166_v1, 3)
				(_nw42).ArraySet1((_230_v45).IsSubsetOf(_230_v45), 4)
				(_nw42).ArraySet1((func() bool {
					if _166_v1 {
						return (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)
					}
					return _166_v1
				})(), 5)
				(_nw42).ArraySet1((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), 6)
				(_nw42).ArraySet1((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), 7)
				(_nw42).ArraySet1(_166_v1, 8)
				(_nw42).ArraySet1(_166_v1, 9)
				(_nw42).ArraySet1((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), 10)
				(_nw42).ArraySet1(_166_v1, 11)
				(_228_v43).ArraySet1(_nw42, (_index22).Int())
				if false {
					(_168_globalState).F0 = (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)
					var _231_v46 D1
					_ = _231_v46
					_231_v46 = Companion_D1_.Create_DC4_(_220_v42)
					var _232_v47 _dafny.Sequence
					_ = _232_v47
					_232_v47 = _dafny.SeqOf(_169_v3, _169_v3, _169_v3, _169_v3)
					var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
					_ = _index23
					var _rhs28 D1 = _231_v46
					_ = _rhs28
					var _rhs29 _dafny.Int = Companion_Default___.Fm1((_dafny.Zero).Minus(_dafny.IntOfUint32((_232_v47).Cardinality())), _168_globalState)
					_ = _rhs29
					var _rhs30 _dafny.Int = _169_v3
					_ = _rhs30
					var _lhs17 *GlobalState = _168_globalState
					_ = _lhs17
					var _lhs18 _dafny.Array = _186_v15
					_ = _lhs18
					var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
					_ = _lhs19
					_231_v46 = _rhs28
					_lhs17.F3 = _rhs29
					(_lhs18).ArraySet1(_rhs30, (_lhs19).Int())
					var _233_v48 _dafny.Sequence
					_ = _233_v48
					_233_v48 = _dafny.UnicodeSeqOfUtf8Bytes("tabouo")
					var _234_v49 *C1
					_ = _234_v49
					var _nw43 *C1 = New_C1_()
					_ = _nw43
					_nw43.Ctor__()
					_234_v49 = _nw43
					var _235_v50 _dafny.MultiSet
					_ = _235_v50
					_235_v50 = _dafny.MultiSetOf(_234_v49)
					var _236_v51 _dafny.Map
					_ = _236_v51
					_236_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_233_v48).Cardinality()), (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool))
					var _rhs31 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_233_v48, _233_v48)
					_ = _rhs31
					var _rhs32 bool = (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)
					_ = _rhs32
					var _rhs33 bool = (_169_v3).Cmp(((_dafny.Zero).Minus(_169_v3)).Times((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))) > 0
					_ = _rhs33
					var _rhs34 bool = (_236_v51).Contains(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_233_v48).Cardinality()), (_235_v50).Cardinality()))
					_ = _rhs34
					var _rhs35 _dafny.Int = (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)
					_ = _rhs35
					var _lhs20 *GlobalState = _168_globalState
					_ = _lhs20
					var _lhs21 *GlobalState = _168_globalState
					_ = _lhs21
					_233_v48 = _rhs31
					_lhs20.F0 = _rhs32
					_166_v1 = _rhs33
					_166_v1 = _rhs34
					_lhs21.F1 = _rhs35
					var _237_v52 *C1
					_ = _237_v52
					var _nw44 *C1 = New_C1_()
					_ = _nw44
					_nw44.Ctor__()
					_237_v52 = _nw44
					var _238_v53 _dafny.Int
					_ = _238_v53
					var _239_v54 _dafny.Array
					_ = _239_v54
					var _240_v55 D0
					_ = _240_v55
					var _out4 _dafny.Int
					_ = _out4
					var _out5 _dafny.Array
					_ = _out5
					var _out6 D0
					_ = _out6
					_out4, _out5, _out6 = (_234_v49).M0(_168_globalState)
					_238_v53 = _out4
					_239_v54 = _out5
					_240_v55 = _out6
				} else {
					var _241_v56 _dafny.Set
					_ = _241_v56
					_241_v56 = _dafny.SetOf(_166_v1)
					var _242_v57 _dafny.Set
					_ = _242_v57
					_242_v57 = _241_v56
					var _243_v58 _dafny.Array
					_ = _243_v58
					var _nwElement0_6 _dafny.Set = _241_v56
					_ = _nwElement0_6
					var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(9))
					_ = _nw45
					(_nw45).ArraySet1(_nwElement0_6, 0)
					(_nw45).ArraySet1((_242_v57), 1)
					(_nw45).ArraySet1(_241_v56, 2)
					(_nw45).ArraySet1(_241_v56, 3)
					(_nw45).ArraySet1(Companion_Default___.Fm18(_dafny.IntOfInt64(914), _168_globalState), 4)
					(_nw45).ArraySet1(_241_v56, 5)
					(_nw45).ArraySet1((_dafny.SetOf(!((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), _166_v1)).Intersection(_dafny.SetOf(_166_v1)), 6)
					(_nw45).ArraySet1((_241_v56).Intersection(_241_v56), 7)
					(_nw45).ArraySet1(_241_v56, 8)
					_243_v58 = _nw45
					var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_243_v58), 0))
					_ = _index24
					(_243_v58).ArraySet1(Companion_Default___.Fm18(_dafny.IntOfInt64(359), _168_globalState), (_index24).Int())
					var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_243_v58), 0))
					_ = _index25
					(_243_v58).ArraySet1(_dafny.SetOf((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), (_index25).Int())
					var _244_v59 _dafny.Array
					_ = _244_v59
					var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
					_ = _nw46
					_244_v59 = _nw46
					var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_244_v59), 0))
					_ = _index26
					(_244_v59).ArraySet1CodePoint(_229_v44, (_index26).Int())
					var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_244_v59), 0))
					_ = _index27
					(_244_v59).ArraySet1CodePoint(_229_v44, (_index27).Int())
					var _245_v60 _dafny.Set
					_ = _245_v60
					_245_v60 = _dafny.SetOf(_169_v3, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _169_v3, _dafny.IntOfInt64(221))
					var _246_v61 _dafny.Array
					_ = _246_v61
					var _len0_10 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_10
					var _nw47 _dafny.Array
					_ = _nw47
					if _len0_10.Cmp(_dafny.Zero) == 0 {
						_nw47 = _dafny.NewArray(_len0_10)
					} else {
						var _init10 func(_dafny.Int) D6 = (func(_247_v2 _dafny.Sequence, _248_v1 bool) func(_dafny.Int) D6 {
							return func(_249_i8 _dafny.Int) D6 {
								return Companion_D6_.Create_DC20_(_247_v2, _248_v1)
							}
						})(_167_v2, _166_v1)
						_ = _init10
						var _element0_10 = _init10(_dafny.Zero)
						_ = _element0_10
						_nw47 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
						(_nw47).ArraySet1(_element0_10, 0)
						var _nativeLen0_10 = (_len0_10).Int()
						_ = _nativeLen0_10
						for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
							(_nw47).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
						}
					}
					_246_v61 = _nw47
					var _250_v62 _dafny.Sequence
					_ = _250_v62
					_250_v62 = _dafny.UnicodeSeqOfUtf8Bytes("gdhmm")
					var _251_v63 _dafny.Int
					_ = _251_v63
					var _252_v64 bool
					_ = _252_v64
					var _253_v65 bool
					_ = _253_v65
					var _254_v66 bool
					_ = _254_v66
					var _out7 _dafny.Int
					_ = _out7
					var _out8 bool
					_ = _out8
					var _out9 bool
					_ = _out9
					var _out10 bool
					_ = _out10
					_out7, _out8, _out9, _out10 = Companion_Default___.M2(_245_v60, _246_v61, (func() _dafny.Int {
						if (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool) {
							return ((_243_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_243_v58), 0))).Int()).(_dafny.Set)).Cardinality()
						}
						return _169_v3
					})(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_250_v62, _250_v62), (Companion_Default___.SafeIndex((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_250_v62, _250_v62)).Cardinality()))).Uint32(), (_244_v59).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_244_v59), 0))).Int())), _168_globalState)
					_251_v63 = _out7
					_252_v64 = _out8
					_253_v65 = _out9
					_254_v66 = _out10
					var _255_v67 _dafny.Int
					_ = _255_v67
					var _256_v68 bool
					_ = _256_v68
					var _257_v69 bool
					_ = _257_v69
					var _258_v70 bool
					_ = _258_v70
					var _out11 _dafny.Int
					_ = _out11
					var _out12 bool
					_ = _out12
					var _out13 bool
					_ = _out13
					var _out14 bool
					_ = _out14
					_out11, _out12, _out13, _out14 = Companion_Default___.M2((_245_v60).Difference(Companion_Default___.Fm19((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), _168_globalState)), _246_v61, ((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus(_169_v3)), _250_v62, _168_globalState)
					_255_v67 = _out11
					_256_v68 = _out12
					_257_v69 = _out13
					_258_v70 = _out14
					_258_v70 = (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)
				}
				if _166_v1 {
					var _259_v71 _dafny.Map
					_ = _259_v71
					_259_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v3, !(_166_v1))
					_259_v71 = (_259_v71).Update((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _166_v1)
					var _260_v72 *C1
					_ = _260_v72
					var _nw48 *C1 = New_C1_()
					_ = _nw48
					_nw48.Ctor__()
					_260_v72 = _nw48
					var _261_v73 _dafny.Map
					_ = _261_v73
					_261_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), true)).Cardinality(), _260_v72)
					var _262_v74 D6
					_ = _262_v74
					_262_v74 = Companion_D6_.Create_DC20_(_167_v2, _166_v1)
					var _263_v75 _dafny.Array
					_ = _263_v75
					var _nwElement0_7 D6 = _262_v74
					_ = _nwElement0_7
					var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(24))
					_ = _nw49
					(_nw49).ArraySet1(_nwElement0_7, 0)
					(_nw49).ArraySet1(_262_v74, 1)
					(_nw49).ArraySet1(_262_v74, 2)
					(_nw49).ArraySet1(Companion_D6_.Create_DC20_(_167_v2, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), 3)
					(_nw49).ArraySet1(_262_v74, 4)
					(_nw49).ArraySet1(_262_v74, 5)
					(_nw49).ArraySet1(_262_v74, 6)
					(_nw49).ArraySet1(_262_v74, 7)
					(_nw49).ArraySet1(_262_v74, 8)
					(_nw49).ArraySet1(_262_v74, 9)
					(_nw49).ArraySet1(_262_v74, 10)
					(_nw49).ArraySet1(_262_v74, 11)
					(_nw49).ArraySet1(_262_v74, 12)
					(_nw49).ArraySet1(_262_v74, 13)
					(_nw49).ArraySet1(_262_v74, 14)
					(_nw49).ArraySet1(_262_v74, 15)
					(_nw49).ArraySet1(_262_v74, 16)
					(_nw49).ArraySet1(_262_v74, 17)
					(_nw49).ArraySet1(Companion_D6_.Create_DC20_(_167_v2, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), 18)
					(_nw49).ArraySet1(_262_v74, 19)
					(_nw49).ArraySet1(_262_v74, 20)
					(_nw49).ArraySet1(_262_v74, 21)
					(_nw49).ArraySet1(Companion_D6_.Create_DC20_(_167_v2, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), 22)
					(_nw49).ArraySet1(_262_v74, 23)
					_263_v75 = _nw49
					var _264_v76 _dafny.Sequence
					_ = _264_v76
					_264_v76 = _dafny.UnicodeSeqOfUtf8Bytes("quqguci")
					var _265_v77 _dafny.Int
					_ = _265_v77
					var _266_v78 bool
					_ = _266_v78
					var _267_v79 bool
					_ = _267_v79
					var _268_v80 bool
					_ = _268_v80
					var _out15 _dafny.Int
					_ = _out15
					var _out16 bool
					_ = _out16
					var _out17 bool
					_ = _out17
					var _out18 bool
					_ = _out18
					_out15, _out16, _out17, _out18 = Companion_Default___.M2(_dafny.SetOf((_261_v73).Cardinality(), (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _169_v3, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)), (func() _dafny.Array {
						if (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool) {
							return _263_v75
						}
						return _263_v75
					})(), _169_v3, _264_v76, _168_globalState)
					_265_v77 = _out15
					_266_v78 = _out16
					_267_v79 = _out17
					_268_v80 = _out18
					var _269_v81 _dafny.Sequence
					_ = _269_v81
					_269_v81 = _dafny.SeqOf(_dafny.IntOfInt64(295), _265_v77, _265_v77)
					_267_v79 = !(!_dafny.Companion_Sequence_.Contains(_269_v81, (_169_v3).Plus((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))))
					var _270_v82 _dafny.Set
					_ = _270_v82
					_270_v82 = _dafny.SetOf(_265_v77, _169_v3)
					var _271_v83 _dafny.Map
					_ = _271_v83
					_271_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v79, _270_v82)
					var _nwElement0_8 _dafny.Int = _265_v77
					_ = _nwElement0_8
					var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(24))
					_ = _nw50
					(_nw50).ArraySet1(_nwElement0_8, 0)
					(_nw50).ArraySet1(_265_v77, 1)
					(_nw50).ArraySet1(((func() _dafny.Set {
						if (_271_v83).Contains(_266_v78) {
							return (_271_v83).Get(_266_v78).(_dafny.Set)
						}
						return Companion_Default___.Fm19((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), _168_globalState)
					})()).Cardinality(), 2)
					(_nw50).ArraySet1((_169_v3).Plus(_169_v3), 3)
					(_nw50).ArraySet1(_265_v77, 4)
					(_nw50).ArraySet1((_dafny.IntOfInt64(338)).Minus(_dafny.IntOfInt64(39)), 5)
					(_nw50).ArraySet1((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), 6)
					(_nw50).ArraySet1((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), 7)
					(_nw50).ArraySet1((_169_v3).Times(_265_v77), 8)
					(_nw50).ArraySet1(_169_v3, 9)
					(_nw50).ArraySet1((_265_v77).Minus(_265_v77), 10)
					(_nw50).ArraySet1((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), 11)
					(_nw50).ArraySet1(_169_v3, 12)
					(_nw50).ArraySet1(_dafny.IntOfInt64(525), 13)
					(_nw50).ArraySet1((_169_v3).Minus(_265_v77), 14)
					(_nw50).ArraySet1((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), 15)
					(_nw50).ArraySet1((_265_v77).Minus((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)), 16)
					(_nw50).ArraySet1((_dafny.Zero).Minus(_169_v3), 17)
					(_nw50).ArraySet1((func() _dafny.Int {
						if (_170_v4).Contains((_220_v42).Cardinality()) {
							return (_170_v4).Get((_220_v42).Cardinality()).(_dafny.Int)
						}
						return (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)
					})(), 18)
					(_nw50).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v78, _266_v78)).Cardinality(), 19)
					(_nw50).ArraySet1(_169_v3, 20)
					(_nw50).ArraySet1(_265_v77, 21)
					(_nw50).ArraySet1(_265_v77, 22)
					(_nw50).ArraySet1(_dafny.IntOfUint32((_167_v2).Cardinality()), 23)
					_186_v15 = _nw50
					_268_v80 = !(Companion_Default___.Fm0(Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(344))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg19 _dafny.Int) interface{} {
							return coer19(arg19)
						}
					}((func(_272_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_273_i9 _dafny.Int) _dafny.CodePoint {
							return _272_v44
						}
					})(_229_v44)))).Cardinality()), _168_globalState), _168_globalState))
				} else {
					var _274_v84 _dafny.MultiSet
					_ = _274_v84
					_274_v84 = _dafny.MultiSetOf(_219_v41, _219_v41)
					(_168_globalState).F0 = !(_274_v84).Equals(_274_v84)
					var _275_v85 D2
					_ = _275_v85
					_275_v85 = Companion_D2_.Create_DC8_(Companion_Default___.SafeModuloInt(_169_v3, (_dafny.MultiSetOf(_dafny.IntOfInt64(31))).Cardinality()), _dafny.IntOfInt64(-850))
					_275_v85 = _275_v85
					var _276_v86 D4
					_ = _276_v86
					_276_v86 = Companion_D4_.Create_DC13_(_dafny.ArrayCastTo((_228_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_228_v43), 0))).Int())))
					var _277_v87 _dafny.Map
					_ = _277_v87
					_277_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v1, _276_v86)
					var _278_v88 _dafny.Array
					_ = _278_v88
					var _len0_11 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_11
					var _nw51 _dafny.Array
					_ = _nw51
					if _len0_11.Cmp(_dafny.Zero) == 0 {
						_nw51 = _dafny.NewArray(_len0_11)
					} else {
						var _init11 func(_dafny.Int) _dafny.Sequence = (func(_279_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
							return func(_280_i10 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg20 _dafny.Int) interface{} {
										return coer20(arg20)
									}
								}((func(_281_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_282_i11 _dafny.Int) _dafny.CodePoint {
										return _281_v44
									}
								})(_279_v44)))
							}
						})(_229_v44)
						_ = _init11
						var _element0_11 = _init11(_dafny.Zero)
						_ = _element0_11
						_nw51 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
						(_nw51).ArraySet1(_element0_11, 0)
						var _nativeLen0_11 = (_len0_11).Int()
						_ = _nativeLen0_11
						for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
							(_nw51).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
						}
					}
					_278_v88 = _nw51
					var _283_v89 *C0
					_ = _283_v89
					var _nw52 *C0 = New_C0_()
					_ = _nw52
					_nw52.Ctor__(_278_v88)
					_283_v89 = _nw52
					var _284_v90 _dafny.Map
					_ = _284_v90
					_284_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v87, _283_v89)
					_284_v90 = (_284_v90).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v1, _276_v86)).Merge(_277_v87), _283_v89)
					var _285_v91 _dafny.Set
					_ = _285_v91
					_285_v91 = _dafny.SetOf((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(_169_v3)).Cardinality()), (_dafny.Zero).Minus(_169_v3), (_dafny.Zero).Minus(_169_v3))
					var _286_v92 D6
					_ = _286_v92
					_286_v92 = Companion_D6_.Create_DC20_(_167_v2, _166_v1)
					var _287_v93 _dafny.Map
					_ = _287_v93
					_287_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v44, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool))
					var _pat_let_tv7 = _219_v41
					_ = _pat_let_tv7
					var _pat_let_tv8 = _219_v41
					_ = _pat_let_tv8
					var _288_v94 _dafny.Array
					_ = _288_v94
					var _nwElement0_9 D6 = _286_v92
					_ = _nwElement0_9
					var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(28))
					_ = _nw53
					(_nw53).ArraySet1(_nwElement0_9, 0)
					(_nw53).ArraySet1(_286_v92, 1)
					(_nw53).ArraySet1(_286_v92, 2)
					(_nw53).ArraySet1(_286_v92, 3)
					(_nw53).ArraySet1(_286_v92, 4)
					(_nw53).ArraySet1(_286_v92, 5)
					(_nw53).ArraySet1(func(_pat_let0_0 D6) D6 {
						return func(_289_dt__update__tmp_h0 D6) D6 {
							return func(_pat_let1_0 _dafny.Sequence) D6 {
								return func(_290_dt__update_hcf35_h0 _dafny.Sequence) D6 {
									return Companion_D6_.Create_DC20_(_290_dt__update_hcf35_h0, (_289_dt__update__tmp_h0).Dtor_cf36())
								}(_pat_let1_0)
							}(_dafny.SeqOf((_pat_let_tv8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_pat_let_tv7), 0))).Int()).(bool)))
						}(_pat_let0_0)
					}(_286_v92), 6)
					(_nw53).ArraySet1(_286_v92, 7)
					(_nw53).ArraySet1(_286_v92, 8)
					(_nw53).ArraySet1(_286_v92, 9)
					(_nw53).ArraySet1(_286_v92, 10)
					(_nw53).ArraySet1(_286_v92, 11)
					(_nw53).ArraySet1(_286_v92, 12)
					(_nw53).ArraySet1(_286_v92, 13)
					(_nw53).ArraySet1(_286_v92, 14)
					(_nw53).ArraySet1(Companion_D6_.Create_DC20_(_167_v2, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), 15)
					(_nw53).ArraySet1(_286_v92, 16)
					(_nw53).ArraySet1(_286_v92, 17)
					(_nw53).ArraySet1(Companion_Default___.Fm20(_287_v93, _166_v1, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _168_globalState), 18)
					(_nw53).ArraySet1(_286_v92, 19)
					(_nw53).ArraySet1(_286_v92, 20)
					(_nw53).ArraySet1(_286_v92, 21)
					(_nw53).ArraySet1(_286_v92, 22)
					(_nw53).ArraySet1(Companion_D6_.Create_DC20_(_dafny.SeqOf(_166_v1, false, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), (Companion_D5_.Create_DC17_((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _166_v1, _166_v1, false, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))).Dtor_cf29(), (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), _166_v1), 23)
					(_nw53).ArraySet1(_286_v92, 24)
					(_nw53).ArraySet1(_286_v92, 25)
					(_nw53).ArraySet1(_286_v92, 26)
					(_nw53).ArraySet1(Companion_D6_.Create_DC20_(_167_v2, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool)), 27)
					_288_v94 = _nw53
					var _291_v95 _dafny.Sequence
					_ = _291_v95
					_291_v95 = _dafny.UnicodeSeqOfUtf8Bytes("sbgug")
					var _292_v96 _dafny.Int
					_ = _292_v96
					var _293_v97 bool
					_ = _293_v97
					var _294_v98 bool
					_ = _294_v98
					var _295_v99 bool
					_ = _295_v99
					var _out19 _dafny.Int
					_ = _out19
					var _out20 bool
					_ = _out20
					var _out21 bool
					_ = _out21
					var _out22 bool
					_ = _out22
					_out19, _out20, _out21, _out22 = Companion_Default___.M2(_285_v91, _288_v94, _169_v3, _291_v95, _168_globalState)
					_292_v96 = _out19
					_293_v97 = _out20
					_294_v98 = _out21
					_295_v99 = _out22
					var _296_v100 _dafny.MultiSet
					_ = _296_v100
					_296_v100 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_169_v3, _169_v3))
					_296_v100 = _296_v100
				}
				_229_v44 = _229_v44
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _297_v101 *C1
	_ = _297_v101
	var _nw54 *C1 = New_C1_()
	_ = _nw54
	_nw54.Ctor__()
	_297_v101 = _nw54
	var _298_v102 _dafny.Sequence
	_ = _298_v102
	_298_v102 = _dafny.SeqOf(_297_v101, _297_v101, _297_v101, _297_v101)
	_297_v101 = (_298_v102).Select((Companion_Default___.SafeIndex(((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)).Times((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_298_v102).Cardinality()))).Uint32()).(*C1)
	var _299_v103 _dafny.Map
	_ = _299_v103
	_299_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v3, _167_v2)
	var _300_v104 _dafny.Sequence
	_ = _300_v104
	_300_v104 = _dafny.UnicodeSeqOfUtf8Bytes("fb")
	var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
	_ = _index28
	(_186_v15).ArraySet1((func() _dafny.Int {
		if _166_v1 {
			return ((_299_v103).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, false, _166_v1, false)).Cardinality()), _167_v2))).Cardinality()
		}
		return _dafny.IntOfUint32((_300_v104).Cardinality())
	})(), (_index28).Int())
	var _301_v105 _dafny.CodePoint
	_ = _301_v105
	_301_v105 = _dafny.CodePoint('d')
	var _302_v106 _dafny.Set
	_ = _302_v106
	_302_v106 = _dafny.SetOf((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))
	var _rhs36 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_169_v3, Companion_Default___.Fm1(_dafny.IntOfInt64(-808), _168_globalState)))
	_ = _rhs36
	var _rhs37 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_300_v104, (Companion_Default___.SafeIndex(_169_v3, _dafny.IntOfUint32((_300_v104).Cardinality()))).Uint32(), _301_v105)
	_ = _rhs37
	var _rhs38 _dafny.Map = (Companion_Default___.Fm21((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _302_v106, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), _168_globalState)).Dtor_cf45()
	_ = _rhs38
	_169_v3 = _rhs36
	_300_v104 = _rhs37
	_170_v4 = _rhs38
	var _303_v107 _dafny.MultiSet
	_ = _303_v107
	_303_v107 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xhuyl"))
	var _304_v108 _dafny.Sequence
	_ = _304_v108
	_304_v108 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}((func(_305_v105 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_306_i12 _dafny.Int) _dafny.CodePoint {
			return _305_v105
		}
	})(_301_v105))))
	_303_v107 = ((_dafny.MultiSetOf(_300_v104, _300_v104)).Intersection(_303_v107)).Union(_dafny.MultiSetFromSeq(_304_v108))
	var _307_v109 _dafny.Int
	_ = _307_v109
	var _308_v110 _dafny.Array
	_ = _308_v110
	var _309_v111 D0
	_ = _309_v111
	var _out23 _dafny.Int
	_ = _out23
	var _out24 _dafny.Array
	_ = _out24
	var _out25 D0
	_ = _out25
	_out23, _out24, _out25 = (_297_v101).M0(_168_globalState)
	_307_v109 = _out23
	_308_v110 = _out24
	_309_v111 = _out25
	var _310_v112 _dafny.Int
	_ = _310_v112
	var _311_v113 _dafny.Array
	_ = _311_v113
	var _312_v114 D0
	_ = _312_v114
	var _out26 _dafny.Int
	_ = _out26
	var _out27 _dafny.Array
	_ = _out27
	var _out28 D0
	_ = _out28
	_out26, _out27, _out28 = (_297_v101).M0(_168_globalState)
	_310_v112 = _out26
	_311_v113 = _out27
	_312_v114 = _out28
	if _166_v1 {
		_169_v3 = Companion_Default___.SafeDivisionInt((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _169_v3)
		_297_v101 = _297_v101
		var _313_v115 _dafny.Array
		_ = _313_v115
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
		_ = _nw55
		_313_v115 = _nw55
		var _314_v116 _dafny.Set
		_ = _314_v116
		_314_v116 = _dafny.SetOf(true)
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_313_v115), 0))
		_ = _index29
		(_313_v115).ArraySet1((_314_v116).Difference(_314_v116), (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_313_v115), 0))
		_ = _index30
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
		_ = _index31
		var _rhs39 _dafny.Set = (Companion_Default___.Fm18(_169_v3, _168_globalState)).Intersection(_314_v116)
		_ = _rhs39
		var _rhs40 _dafny.Int = Companion_Default___.SafeModuloInt((Companion_Default___.Fm1(Companion_Default___.Fm1((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int), _168_globalState), _168_globalState)).Plus((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)), ((_314_v116).Cardinality()).Minus((_184_v13).Cardinality()))
		_ = _rhs40
		var _rhs41 _dafny.Int = Companion_Default___.Fm1(_310_v112, _168_globalState)
		_ = _rhs41
		var _lhs22 _dafny.Array = _313_v115
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_313_v115), 0))
		_ = _lhs23
		var _lhs24 _dafny.Array = _186_v15
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))
		_ = _lhs25
		var _lhs26 *GlobalState = _168_globalState
		_ = _lhs26
		(_lhs22).ArraySet1(_rhs39, (_lhs23).Int())
		(_lhs24).ArraySet1(_rhs40, (_lhs25).Int())
		_lhs26.F3 = _rhs41
		var _315_v117 _dafny.Array
		_ = _315_v117
		var _nwElement0_10 _dafny.Map = (_184_v13).Update(_166_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(563))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_316_v105 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_317_i13 _dafny.Int) _dafny.CodePoint {
				return _316_v105
			}
		})(_301_v105)))).Cardinality()))
		_ = _nwElement0_10
		var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(17))
		_ = _nw56
		(_nw56).ArraySet1(_nwElement0_10, 0)
		(_nw56).ArraySet1(_184_v13, 1)
		(_nw56).ArraySet1(Companion_Default___.Fm22(_166_v1, _301_v105, _dafny.UnicodeSeqOfUtf8Bytes("gdnjcodfx"), !(!(_166_v1)), _168_globalState), 2)
		(_nw56).ArraySet1(_184_v13, 3)
		(_nw56).ArraySet1((_184_v13).Merge(_184_v13), 4)
		(_nw56).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), (_dafny.Zero).Minus((_170_v4).Cardinality())), 5)
		(_nw56).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v1, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))).Merge(_184_v13), 6)
		(_nw56).ArraySet1(_184_v13, 7)
		(_nw56).ArraySet1(_184_v13, 8)
		(_nw56).ArraySet1((func() _dafny.Map {
			if (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool) {
				return _184_v13
			}
			return _184_v13
		})(), 9)
		(_nw56).ArraySet1((_184_v13).Merge(_184_v13), 10)
		(_nw56).ArraySet1(_184_v13, 11)
		(_nw56).ArraySet1((_184_v13).Merge(_184_v13), 12)
		(_nw56).ArraySet1(_184_v13, 13)
		(_nw56).ArraySet1(_184_v13, 14)
		(_nw56).ArraySet1((_184_v13).Merge(_184_v13), 15)
		(_nw56).ArraySet1(_184_v13, 16)
		_315_v117 = _nw56
		var _rhs42 _dafny.Map = ((func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter10 := _dafny.Iterate((_170_v4).Keys().Elements()); ; {
				_compr_9, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _318_v118 _dafny.Int
				_318_v118 = interface{}(_compr_9).(_dafny.Int)
				if (_170_v4).Contains(_318_v118) {
					_coll9.Add((_318_v118).Times((_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_169_v3)).Cardinality(), _310_v112, _307_v109, _169_v3, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))).Cardinality()))
				}
			}
			return _coll9.ToMap()
		}()).Merge(_170_v4)).Merge(_170_v4)
		_ = _rhs42
		var _rhs43 _dafny.Array = _315_v117
		_ = _rhs43
		_170_v4 = _rhs42
		_315_v117 = _rhs43
		_169_v3 = (func() _dafny.Int {
			if (_307_v109).Cmp(_307_v109) <= 0 {
				return (_169_v3).Times(_dafny.IntOfInt64(-212))
			}
			return _310_v112
		})()
	} else {
		var _319_v119 D3
		_ = _319_v119
		_319_v119 = Companion_D3_.Create_DC11_(_dafny.IntOfInt64(174))
		var _320_v120 D0
		_ = _320_v120
		_320_v120 = Companion_D0_.Create_DC0_((_319_v119).Dtor_cf14())
		var _321_v121 _dafny.Sequence
		_ = _321_v121
		_321_v121 = _dafny.SeqOf(_310_v112)
		var _322_v122 _dafny.Sequence
		_ = _322_v122
		_322_v122 = _dafny.SeqOf(_321_v121)
		var _323_v123 D4
		_ = _323_v123
		_323_v123 = Companion_D4_.Create_DC14_(_310_v112, _320_v120, (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), _166_v1, (_322_v122).Select((Companion_Default___.SafeIndex(_307_v109, _dafny.IntOfUint32((_322_v122).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _324_v124 D4
		_ = _324_v124
		_324_v124 = Companion_D4_.Create_DC15_(Companion_D4_.Create_DC15_(_323_v123))
		_324_v124 = _324_v124
		(_168_globalState).F1 = _310_v112
		var _325_v125 _dafny.Sequence
		_ = _325_v125
		_325_v125 = _dafny.SeqOf(_186_v15, (func() _dafny.Array {
			if true {
				return _186_v15
			}
			return _186_v15
		})())
		var _326_v126 _dafny.MultiSet
		_ = _326_v126
		_326_v126 = _dafny.MultiSetOf(_310_v112)
		var _327_v127 D6
		_ = _327_v127
		_327_v127 = Companion_D6_.Create_DC19_(_325_v125)
		var _rhs44 *C1 = _297_v101
		_ = _rhs44
		var _rhs45 _dafny.Int = Companion_Default___.Fm1((func() _dafny.Int {
			if (_326_v126).Contains(_169_v3) {
				return (_326_v126).Multiplicity(_169_v3)
			}
			return _310_v112
		})(), _168_globalState)
		_ = _rhs45
		var _rhs46 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_325_v125, (_327_v127).Dtor_cf34()), _dafny.SeqOf(_186_v15, _186_v15))
		_ = _rhs46
		var _rhs47 _dafny.Array = _311_v113
		_ = _rhs47
		_297_v101 = _rhs44
		_307_v109 = _rhs45
		_325_v125 = _rhs46
		_311_v113 = _rhs47
		var _328_v128 _dafny.Array
		_ = _328_v128
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
		_ = _nw57
		_328_v128 = _nw57
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_328_v128), 0))
		_ = _index32
		(_328_v128).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(622))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_329_v105 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_330_i14 _dafny.Int) _dafny.CodePoint {
				return (Companion_D0_.Create_DC1_(_329_v105)).Dtor_cf1()
			}
		})(_301_v105))), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_328_v128), 0))
		_ = _index33
		(_328_v128).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_300_v104, _300_v104), _300_v104), (_index33).Int())
		var _331_v129 _dafny.Array
		_ = _331_v129
		var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw58
		_331_v129 = _nw58
		var _332_v130 _dafny.Map
		_ = _332_v130
		_332_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), _331_v129)
		var _rhs48 _dafny.Int = (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int)
		_ = _rhs48
		var _rhs49 _dafny.Array = (func() _dafny.Array {
			if (_332_v130).Contains(_dafny.Companion_Sequence_.IsPrefixOf((_328_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_328_v128), 0))).Int()).(_dafny.Sequence), (_328_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_328_v128), 0))).Int()).(_dafny.Sequence))) {
				return (_332_v130).Get(_dafny.Companion_Sequence_.IsPrefixOf((_328_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_328_v128), 0))).Int()).(_dafny.Sequence), (_328_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_328_v128), 0))).Int()).(_dafny.Sequence))).(_dafny.Array)
			}
			return _331_v129
		})()
		_ = _rhs49
		_307_v109 = _rhs48
		_331_v129 = _rhs49
	}
	_170_v4 = (_170_v4).Update(_310_v112, _307_v109)
	var _333_v131 _dafny.Map
	_ = _333_v131
	_333_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v105, _307_v109)
	var _334_v132 _dafny.Sequence
	_ = _334_v132
	_334_v132 = _dafny.SeqOf(_333_v131)
	var _hi2 _dafny.Int = ((_334_v132).Select((Companion_Default___.SafeIndex(_310_v112, _dafny.IntOfUint32((_334_v132).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()
	_ = _hi2
	for _335_i15 := (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jywtmvy")).Cardinality())).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-98))); _335_i15.Cmp(_hi2) < 0; _335_i15 = _335_i15.Plus(_dafny.One) {
		var _336_v133 _dafny.MultiSet
		_ = _336_v133
		_336_v133 = _dafny.MultiSetOf((_302_v106).Cardinality(), _dafny.IntOfUint32((_300_v104).Cardinality()), _335_i15)
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))
		_ = _index34
		(_219_v41).ArraySet1((_297_v101).Fm2(_336_v133, _168_globalState), (_index34).Int())
		var _337_v134 _dafny.Array
		_ = _337_v134
		var _nwElement0_11 _dafny.Map = _184_v13
		_ = _nwElement0_11
		var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(5))
		_ = _nw59
		(_nw59).ArraySet1(_nwElement0_11, 0)
		(_nw59).ArraySet1((func() _dafny.Map {
			if _166_v1 {
				return _184_v13
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v1, (_186_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_186_v15), 0))).Int()).(_dafny.Int))
		})(), 1)
		(_nw59).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool), _307_v109), 2)
		(_nw59).ArraySet1(_184_v13, 3)
		(_nw59).ArraySet1(_184_v13, 4)
		_337_v134 = _nw59
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_337_v134), 0))
		_ = _index35
		(_337_v134).ArraySet1(_184_v13, (_index35).Int())
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_337_v134), 0))
		_ = _index36
		(_337_v134).ArraySet1(_184_v13, (_index36).Int())
		_166_v1 = !(!(_166_v1))
		var _338_v135 _dafny.Array
		_ = _338_v135
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_12
		var _nw60 _dafny.Array
		_ = _nw60
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw60 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Sequence = (func(_339_v104 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_340_i16 _dafny.Int) _dafny.Sequence {
					return _339_v104
				}
			})(_300_v104)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw60 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw60).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw60).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_338_v135 = _nw60
		var _341_v136 *C0
		_ = _341_v136
		var _nw61 *C0 = New_C0_()
		_ = _nw61
		_nw61.Ctor__(_338_v135)
		_341_v136 = _nw61
		var _342_v137 _dafny.Map
		_ = _342_v137
		_342_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v136, _297_v101)
		var _343_v138 _dafny.Map
		_ = _343_v138
		_343_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_341_v136).F6(), _297_v101)
		var _344_v139 _dafny.Array
		_ = _344_v139
		var _nwElement0_12 *C1 = _297_v101
		_ = _nwElement0_12
		var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(28))
		_ = _nw62
		(_nw62).ArraySet1(_nwElement0_12, 0)
		(_nw62).ArraySet1(_297_v101, 1)
		(_nw62).ArraySet1(_297_v101, 2)
		(_nw62).ArraySet1(_297_v101, 3)
		(_nw62).ArraySet1(_297_v101, 4)
		(_nw62).ArraySet1((func() *C1 {
			if (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool) {
				return (func() *C1 {
					if (_342_v137).Contains(_341_v136) {
						return (_342_v137).Get(_341_v136).(*C1)
					}
					return _297_v101
				})()
			}
			return _297_v101
		})(), 5)
		(_nw62).ArraySet1((func() *C1 {
			if (_343_v138).Contains(_338_v135) {
				return (_343_v138).Get(_338_v135).(*C1)
			}
			return _297_v101
		})(), 6)
		(_nw62).ArraySet1(_297_v101, 7)
		(_nw62).ArraySet1(_297_v101, 8)
		(_nw62).ArraySet1((_298_v102).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.IntOfUint32((_298_v102).Cardinality()))).Uint32()).(*C1), 9)
		(_nw62).ArraySet1(_297_v101, 10)
		(_nw62).ArraySet1(_297_v101, 11)
		(_nw62).ArraySet1(_297_v101, 12)
		(_nw62).ArraySet1(_297_v101, 13)
		(_nw62).ArraySet1(_297_v101, 14)
		(_nw62).ArraySet1(_297_v101, 15)
		(_nw62).ArraySet1(_297_v101, 16)
		(_nw62).ArraySet1(_297_v101, 17)
		(_nw62).ArraySet1(_297_v101, 18)
		(_nw62).ArraySet1(_297_v101, 19)
		(_nw62).ArraySet1(_297_v101, 20)
		(_nw62).ArraySet1(_297_v101, 21)
		(_nw62).ArraySet1(_297_v101, 22)
		(_nw62).ArraySet1(_297_v101, 23)
		(_nw62).ArraySet1(_297_v101, 24)
		(_nw62).ArraySet1((func() *C1 {
			if (_219_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_219_v41), 0))).Int()).(bool) {
				return _297_v101
			}
			return _297_v101
		})(), 25)
		(_nw62).ArraySet1(_297_v101, 26)
		(_nw62).ArraySet1(_297_v101, 27)
		_344_v139 = _nw62
		_344_v139 = _344_v139
	}
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_164_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_166_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_167_v2, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_168_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_168_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_168_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_168_globalState.F4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_168_globalState).F5(), _dafny.SeqOf(false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_169_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(5)).UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(2)).UpdateUnsafe(_dafny.IntOfInt64(265), _dafny.IntOfInt64(982)).UpdateUnsafe(_dafny.IntOfInt64(534), _dafny.IntOfInt64(534))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-272)).UpdateUnsafe(true, _dafny.IntOfInt64(272))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_v14).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v15).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v16).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v42).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_227_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_298_v102).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_299_v103).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(272), _dafny.SeqOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_300_v104.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_301_v105)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_302_v106).Equals(_dafny.SetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_303_v107).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_304_v108, _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_307_v109)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_v111).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_310_v112)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_312_v114).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_333_v131).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfInt64(534))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_334_v132, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfInt64(534)))))
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
	Cf1 _dafny.CodePoint
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.CodePoint) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool) D0 {
	return D0{D0_DC2{Cf2}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf3 bool
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf3 bool) D0 {
	return D0{D0_DC3{Cf3}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.CodePoint('D'))
}

func (_this D0) Dtor_cf1() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC3).Cf3
}

func (_this D0) Dtor_cf0() _dafny.Int {
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
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf3) + ")"
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
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf3 == data2.Cf3
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
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
	Cf5 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf5 bool) D1 {
	return D1{D1_DC5{Cf5}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf4 _dafny.MultiSet
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.MultiSet) D1 {
	return D1{D1_DC4{Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC6 struct {
	Cf6 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf6 D1) D1 {
	return D1{D1_DC6{Cf6}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC5).Cf5
}

func (_this D1) Dtor_cf4() _dafny.MultiSet {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf6() D1 {
	return _this.Get_().(D1_DC6).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf5 == data2.Cf5
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf8 _dafny.Int
	Cf9 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf8 _dafny.Int, Cf9 _dafny.Int) D2 {
	return D2{D2_DC8{Cf8, Cf9}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf7 *C0
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf7 *C0) D2 {
	return D2{D2_DC7{Cf7}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf9
}

func (_this D2) Dtor_cf7() *C0 {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf7 == data2.Cf7
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
	Cf11 bool
	Cf12 bool
	Cf13 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf11 bool, Cf12 bool, Cf13 bool) D3 {
	return D3{D3_DC10{Cf11, Cf12, Cf13}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf14 _dafny.Int
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf14 _dafny.Int) D3 {
	return D3{D3_DC11{Cf14}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC12 struct {
	Cf15 _dafny.Int
	Cf16 _dafny.MultiSet
	Cf17 _dafny.Int
	Cf18 _dafny.Array
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf15 _dafny.Int, Cf16 _dafny.MultiSet, Cf17 _dafny.Int, Cf18 _dafny.Array) D3 {
	return D3{D3_DC12{Cf15, Cf16, Cf17, Cf18}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC9 struct {
	Cf10 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf10 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf10}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, false, false)
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC10).Cf11
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC12).Cf15
}

func (_this D3) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D3_DC12).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC12).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Array {
	return _this.Get_().(D3_DC12).Cf18
}

func (_this D3) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Equals(data2.Cf16) && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18 == data2.Cf18
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D4_DC14 struct {
	Cf20 _dafny.Int
	Cf21 D0
	Cf22 bool
	Cf23 bool
	Cf24 _dafny.Sequence
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf20 _dafny.Int, Cf21 D0, Cf22 bool, Cf23 bool, Cf24 _dafny.Sequence) D4 {
	return D4{D4_DC14{Cf20, Cf21, Cf22, Cf23, Cf24}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf19 _dafny.Array
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf19 _dafny.Array) D4 {
	return D4{D4_DC13{Cf19}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC15 struct {
	Cf25 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf25 D4) D4 {
	return D4{D4_DC15{Cf25}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(_dafny.Zero, Companion_D0_.Default(), false, false, _dafny.EmptySeq)
}

func (_this D4) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf20
}

func (_this D4) Dtor_cf21() D0 {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC14).Cf22
}

func (_this D4) Dtor_cf23() bool {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D4_DC14).Cf24
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) Dtor_cf25() D4 {
	return _this.Get_().(D4_DC15).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Equals(data2.Cf21) && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23 && data1.Cf24.Equals(data2.Cf24)
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf19 == data2.Cf19
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

type D5_DC17 struct {
	Cf27 _dafny.Int
	Cf28 bool
	Cf29 bool
	Cf30 bool
	Cf31 _dafny.Int
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf27 _dafny.Int, Cf28 bool, Cf29 bool, Cf30 bool, Cf31 _dafny.Int) D5 {
	return D5{D5_DC17{Cf27, Cf28, Cf29, Cf30, Cf31}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC18 struct {
	Cf32 _dafny.Int
	Cf33 _dafny.Int
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf32 _dafny.Int, Cf33 _dafny.Int) D5 {
	return D5{D5_DC18{Cf32, Cf33}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC16 struct {
	Cf26 _dafny.Sequence
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf26 _dafny.Sequence) D5 {
	return D5{D5_DC16{Cf26}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(_dafny.Zero, false, false, false, _dafny.Zero)
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf27
}

func (_this D5) Dtor_cf28() bool {
	return _this.Get_().(D5_DC17).Cf28
}

func (_this D5) Dtor_cf29() bool {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) Dtor_cf30() bool {
	return _this.Get_().(D5_DC17).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf31
}

func (_this D5) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf32
}

func (_this D5) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf33
}

func (_this D5) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D5_DC16).Cf26
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + data.Cf26.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D6_DC20 struct {
	Cf35 _dafny.Sequence
	Cf36 bool
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf35 _dafny.Sequence, Cf36 bool) D6 {
	return D6{D6_DC20{Cf35, Cf36}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC21 struct {
	Cf37 _dafny.Int
	Cf38 bool
	Cf39 _dafny.Sequence
	Cf40 bool
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf37 _dafny.Int, Cf38 bool, Cf39 _dafny.Sequence, Cf40 bool) D6 {
	return D6{D6_DC21{Cf37, Cf38, Cf39, Cf40}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC22 struct {
	Cf41 _dafny.Int
	Cf42 _dafny.Int
}

func (D6_DC22) isD6() {}

func (CompanionStruct_D6_) Create_DC22_(Cf41 _dafny.Int, Cf42 _dafny.Int) D6 {
	return D6{D6_DC22{Cf41, Cf42}}
}

func (_this D6) Is_DC22() bool {
	_, ok := _this.Get_().(D6_DC22)
	return ok
}

type D6_DC19 struct {
	Cf34 _dafny.Sequence
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf34 _dafny.Sequence) D6 {
	return D6{D6_DC19{Cf34}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC23 struct {
	Cf43 D6
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf43 D6) D6 {
	return D6{D6_DC23{Cf43}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC20_(_dafny.EmptySeq, false)
}

func (_this D6) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D6_DC20).Cf35
}

func (_this D6) Dtor_cf36() bool {
	return _this.Get_().(D6_DC20).Cf36
}

func (_this D6) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D6_DC21).Cf37
}

func (_this D6) Dtor_cf38() bool {
	return _this.Get_().(D6_DC21).Cf38
}

func (_this D6) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D6_DC21).Cf39
}

func (_this D6) Dtor_cf40() bool {
	return _this.Get_().(D6_DC21).Cf40
}

func (_this D6) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D6_DC22).Cf41
}

func (_this D6) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D6_DC22).Cf42
}

func (_this D6) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D6_DC19).Cf34
}

func (_this D6) Dtor_cf43() D6 {
	return _this.Get_().(D6_DC23).Cf43
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + data.Cf39.VerbatimString(true) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D6_DC22:
		{
			return "D6.DC22" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf35.Equals(data2.Cf35) && data1.Cf36 == data2.Cf36
		}
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38 && data1.Cf39.Equals(data2.Cf39) && data1.Cf40 == data2.Cf40
		}
	case D6_DC22:
		{
			data2, ok := other.Get_().(D6_DC22)
			return ok && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42.Cmp(data2.Cf42) == 0
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D7_DC24 struct {
	Cf44 _dafny.Set
}

func (D7_DC24) isD7() {}

func (CompanionStruct_D7_) Create_DC24_(Cf44 _dafny.Set) D7 {
	return D7{D7_DC24{Cf44}}
}

func (_this D7) Is_DC24() bool {
	_, ok := _this.Get_().(D7_DC24)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D7) Dtor_cf44() _dafny.Set {
	return _this.Get_().(D7_DC24).Cf44
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC24:
		{
			return "D7.DC24" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC24:
		{
			data2, ok := other.Get_().(D7_DC24)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D8_DC26 struct {
	Cf46 bool
	Cf47 _dafny.Array
	Cf48 bool
	Cf49 bool
	Cf50 bool
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf46 bool, Cf47 _dafny.Array, Cf48 bool, Cf49 bool, Cf50 bool) D8 {
	return D8{D8_DC26{Cf46, Cf47, Cf48, Cf49, Cf50}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC25 struct {
	Cf45 _dafny.Map
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf45 _dafny.Map) D8 {
	return D8{D8_DC25{Cf45}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC27 struct {
	Cf51 D8
}

func (D8_DC27) isD8() {}

func (CompanionStruct_D8_) Create_DC27_(Cf51 D8) D8 {
	return D8{D8_DC27{Cf51}}
}

func (_this D8) Is_DC27() bool {
	_, ok := _this.Get_().(D8_DC27)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC26_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, false, false)
}

func (_this D8) Dtor_cf46() bool {
	return _this.Get_().(D8_DC26).Cf46
}

func (_this D8) Dtor_cf47() _dafny.Array {
	return _this.Get_().(D8_DC26).Cf47
}

func (_this D8) Dtor_cf48() bool {
	return _this.Get_().(D8_DC26).Cf48
}

func (_this D8) Dtor_cf49() bool {
	return _this.Get_().(D8_DC26).Cf49
}

func (_this D8) Dtor_cf50() bool {
	return _this.Get_().(D8_DC26).Cf50
}

func (_this D8) Dtor_cf45() _dafny.Map {
	return _this.Get_().(D8_DC25).Cf45
}

func (_this D8) Dtor_cf51() D8 {
	return _this.Get_().(D8_DC27).Cf51
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D8_DC27:
		{
			return "D8.DC27" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48 && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D8_DC27:
		{
			data2, ok := other.Get_().(D8_DC27)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D9_DC28 struct {
	Cf52 _dafny.Map
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf52 _dafny.Map) D9 {
	return D9{D9_DC28{Cf52}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D9) Dtor_cf52() _dafny.Map {
	return _this.Get_().(D9_DC28).Cf52
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf52) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf52.Equals(data2.Cf52)
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

// Definition of class GlobalState
type GlobalState struct {
	F0  bool
	F1  _dafny.Int
	F3  _dafny.Int
	F4  _dafny.Array
	_f2 bool
	_f5 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F1 = _dafny.Zero
	_this.F3 = _dafny.Zero
	_this.F4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = false
	_this._f5 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 _dafny.Int, f4 _dafny.Array, f5 _dafny.Sequence) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F5() _dafny.Sequence {
	{
		return _this._f5
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f6 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f6 _dafny.Array) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C0) Fm6(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("uy")
	}
}
func (_this *C0) Fm7(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(725)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, false), true)).Cardinality())
	}
}
func (_this *C0) F6() _dafny.Array {
	{
		return _this._f6
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
func (_this *C1) Fm2(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !(func(_source3 D0) bool {
			if _source3.Is_DC1() {
				var _345___mcc_h0 _dafny.CodePoint = _source3.Get_().(D0_DC1).Cf1
				_ = _345___mcc_h0
				var _346_cf1 _dafny.CodePoint = _345___mcc_h0
				_ = _346_cf1
				return false
			} else if _source3.Is_DC2() {
				var _347___mcc_h1 bool = _source3.Get_().(D0_DC2).Cf2
				_ = _347___mcc_h1
				var _348_cf2 bool = _347___mcc_h1
				_ = _348_cf2
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(58), _dafny.IntOfInt64(-449)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(773))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}(func(_349_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(3)
				})))
			} else if _source3.Is_DC3() {
				var _350___mcc_h2 bool = _source3.Get_().(D0_DC3).Cf3
				_ = _350___mcc_h2
				var _351_cf3 bool = _350___mcc_h2
				_ = _351_cf3
				if _351_cf3 {
					return false
				} else {
					return _351_cf3
				}
			} else {
				var _352___mcc_h3 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
				_ = _352___mcc_h3
				var _353_cf0 _dafny.Int = _352___mcc_h3
				_ = _353_cf0
				return true
			}
		}(Companion_D0_.Create_DC0_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hw")).Cardinality())))))
	}
}
func (_this *C1) Fm3(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('c')
	}
}
func (_this *C1) M0(globalState *GlobalState) (_dafny.Int, _dafny.Array, D0) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 D0 = Companion_D0_.Default()
		_ = r2
		var _354_v0 _dafny.Array
		_ = _354_v0
		var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw63
		_354_v0 = _nw63
		r1 = _354_v0
		r1 = _354_v0
		var _355_v1 bool
		_ = _355_v1
		_355_v1 = false
		if _355_v1 {
			var _356_v2 _dafny.Array
			_ = _356_v2
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_13
			var _nw64 _dafny.Array
			_ = _nw64
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw64 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Sequence = func(_357_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pl"), _dafny.UnicodeSeqOfUtf8Bytes("lg"))
				}
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
			_356_v2 = _nw64
			var _358_v3 _dafny.Int
			_ = _358_v3
			_358_v3 = _dafny.IntOfInt64(640)
			var _359_v4 _dafny.Array
			_ = _359_v4
			var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw65
			_359_v4 = _nw65
			var _360_v5 _dafny.Sequence
			_ = _360_v5
			_360_v5 = _dafny.SeqOf(_359_v4, _359_v4, _359_v4)
			var _361_v6 _dafny.Set
			_ = _361_v6
			_361_v6 = _dafny.SetOf(_358_v3, _dafny.IntOfUint32((_360_v5).Cardinality()))
			var _362_v7 _dafny.CodePoint
			_ = _362_v7
			_362_v7 = _dafny.CodePoint('b')
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))
			_ = _index37
			(_356_v2).ArraySet1(Companion_Default___.Fm4(_355_v1, _358_v3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_361_v6, _355_v1)).Cardinality(), _362_v7, globalState), (_index37).Int())
			var _363_v8 _dafny.Sequence
			_ = _363_v8
			_363_v8 = _dafny.UnicodeSeqOfUtf8Bytes("eottg")
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))
			_ = _index38
			(_356_v2).ArraySet1(_363_v8, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_354_v0), 0))
			_ = _index39
			(_354_v0).ArraySet1(_355_v1, (_index39).Int())
			var _364_v9 _dafny.Sequence
			_ = _364_v9
			_364_v9 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_363_v8, _363_v8), _dafny.UnicodeSeqOfUtf8Bytes("srkfki"), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence), _363_v8), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence), _363_v8)).Cardinality()))).Uint32(), _362_v7), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-93))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}(func(_365_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})), (_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence))
			var _366_v10 D0
			_ = _366_v10
			_366_v10 = Companion_D0_.Create_DC1_(_362_v7)
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_354_v0), 0))
			_ = _index40
			var _rhs50 bool = _355_v1
			_ = _rhs50
			var _rhs51 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm4(_355_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v0, !(_355_v1))).Cardinality(), _dafny.IntOfInt64(-828), _362_v7, globalState), (_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence)), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm5(_358_v3, _358_v3, globalState), (Companion_Default___.SafeIndex(_358_v3, _dafny.IntOfUint32((Companion_Default___.Fm5(_358_v3, _358_v3, globalState)).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("xkrtmgfeb")), (Companion_Default___.SafeIndex(_358_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm5(_358_v3, _358_v3, globalState), (Companion_Default___.SafeIndex(_358_v3, _dafny.IntOfUint32((Companion_Default___.Fm5(_358_v3, _358_v3, globalState)).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("xkrtmgfeb"))).Cardinality()))).Uint32(), _363_v8)), _364_v9)
			_ = _rhs51
			var _rhs52 D0 = func(_pat_let2_0 D0) D0 {
				return func(_367_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let3_0 _dafny.CodePoint) D0 {
						return func(_368_dt__update_hcf1_h0 _dafny.CodePoint) D0 {
							return Companion_D0_.Create_DC1_(_368_dt__update_hcf1_h0)
						}(_pat_let3_0)
					}(_dafny.CodePoint('g'))
				}(_pat_let2_0)
			}(_366_v10)
			_ = _rhs52
			var _lhs27 _dafny.Array = _354_v0
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_354_v0), 0))
			_ = _lhs28
			(_lhs27).ArraySet1(_rhs50, (_lhs28).Int())
			_364_v9 = _rhs51
			_366_v10 = _rhs52
			var _369_v11 D0
			_ = _369_v11
			_369_v11 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-402))
			var _370_v12 _dafny.Sequence
			_ = _370_v12
			_370_v12 = _dafny.SeqOf(_369_v11, _369_v11, _369_v11, Companion_D0_.Create_DC0_(_358_v3))
			(globalState).F3 = _dafny.IntOfUint32((_370_v12).Cardinality())
			(globalState).F1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(272), _358_v3)
			var _371_v13 _dafny.Map
			_ = _371_v13
			_371_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_362_v7, (_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence))
			_371_v13 = (_371_v13).Update(_362_v7, _dafny.Companion_Sequence_.Update((_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_358_v3, _dafny.IntOfUint32(((_356_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_356_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _362_v7))
		} else {
			var _372_v14 _dafny.Int
			_ = _372_v14
			_372_v14 = _dafny.IntOfInt64(382)
			(globalState).F1 = ((_dafny.Zero).Minus(_372_v14)).Times(_372_v14)
			var _373_v15 _dafny.Sequence
			_ = _373_v15
			_373_v15 = _dafny.UnicodeSeqOfUtf8Bytes("kp")
			if !(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(181))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}(func(_374_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), _373_v15)) || (_355_v1) {
				(globalState).F1 = _372_v14
				var _375_v16 _dafny.CodePoint
				_ = _375_v16
				_375_v16 = _dafny.CodePoint('v')
				var _376_v17 _dafny.Map
				_ = _376_v17
				_376_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC1_(_375_v16)), Companion_D0_.Create_DC1_(_375_v16)), _372_v14)
				var _377_v18 _dafny.Map
				_ = _377_v18
				_377_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_v1, _dafny.SeqOf(_355_v1))
				var _378_v19 _dafny.Sequence
				_ = _378_v19
				_378_v19 = _dafny.SeqOf(_355_v1)
				_376_v17 = (_376_v17).Update(_355_v1, (_dafny.Zero).Minus(((_377_v18).Update(_355_v1, _378_v19)).Cardinality()))
				var _379_v20 _dafny.Array
				_ = _379_v20
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw66
				_379_v20 = _nw66
				var _380_v21 *C0
				_ = _380_v21
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__(_379_v20)
				_380_v21 = _nw67
				var _381_v22 _dafny.Array
				_ = _381_v22
				var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw68
				_381_v22 = _nw68
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_381_v22), 0))
				_ = _index41
				(_381_v22).ArraySet1((_380_v21).F6(), (_index41).Int())
				var _382_v23 _dafny.Sequence
				_ = _382_v23
				_382_v23 = _dafny.SeqOf((_380_v21).F6())
				var _383_v24 _dafny.Sequence
				_ = _383_v24
				_383_v24 = _dafny.SeqOf((_380_v21).F6(), (_380_v21).F6(), _379_v20, (_382_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_378_v19).Cardinality()), _dafny.IntOfUint32((_382_v23).Cardinality()))).Uint32()).(_dafny.Array), (_380_v21).F6())
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_381_v22), 0))
				_ = _index42
				(_381_v22).ArraySet1((_382_v23).Select((Companion_Default___.SafeIndex(_372_v14, _dafny.IntOfUint32((_382_v23).Cardinality()))).Uint32()).(_dafny.Array), (_index42).Int())
				(globalState).F1 = _372_v14
			} else {
				var _384_v25 _dafny.Array
				_ = _384_v25
				var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
				_ = _nw69
				_384_v25 = _nw69
				var _385_v26 *C0
				_ = _385_v26
				var _nw70 *C0 = New_C0_()
				_ = _nw70
				_nw70.Ctor__(_384_v25)
				_385_v26 = _nw70
				var _386_v27 *C0
				_ = _386_v27
				var _nw71 *C0 = New_C0_()
				_ = _nw71
				_nw71.Ctor__(_384_v25)
				_386_v27 = _nw71
				var _387_v28 _dafny.Map
				_ = _387_v28
				_387_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm1(_dafny.IntOfUint32((_373_v15).Cardinality()), globalState)).Cmp(_372_v14) >= 0, _372_v14)
				var _388_v29 _dafny.CodePoint
				_ = _388_v29
				_388_v29 = _dafny.CodePoint('o')
				_387_v28 = (_387_v28).Update(!(!((Companion_Default___.Fm8(_355_v1, _355_v1, _dafny.IntOfUint32((_373_v15).Cardinality()), globalState)).IsProperSubsetOf(_dafny.SetOf(_388_v29)))), _372_v14)
				(globalState).F0 = _355_v1
				r0 = _372_v14
			}
			(globalState).F3 = _dafny.IntOfInt64(240)
			var _389_v30 _dafny.CodePoint
			_ = _389_v30
			_389_v30 = _dafny.CodePoint('n')
			var _390_v31 _dafny.Array
			_ = _390_v31
			var _nwElement0_13 _dafny.Sequence = Companion_Default___.Fm4(_355_v1, _372_v14, _372_v14, _389_v30, globalState)
			_ = _nwElement0_13
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(8))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_13, 0)
			(_nw72).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jcokmtins"), 1)
			(_nw72).ArraySet1(_373_v15, 2)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Update(_373_v15, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_373_v15).Cardinality()), _dafny.IntOfUint32((_373_v15).Cardinality()))).Uint32(), _389_v30), 3)
			(_nw72).ArraySet1(_373_v15, 4)
			(_nw72).ArraySet1(_373_v15, 5)
			(_nw72).ArraySet1(_373_v15, 6)
			(_nw72).ArraySet1(_373_v15, 7)
			_390_v31 = _nw72
			var _391_v32 *C0
			_ = _391_v32
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__(_390_v31)
			_391_v32 = _nw73
			var _392_v33 _dafny.Array
			_ = _392_v33
			var _nwElement0_14 *C0 = _391_v32
			_ = _nwElement0_14
			var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(15))
			_ = _nw74
			(_nw74).ArraySet1(_nwElement0_14, 0)
			(_nw74).ArraySet1(_391_v32, 1)
			(_nw74).ArraySet1(_391_v32, 2)
			(_nw74).ArraySet1(_391_v32, 3)
			(_nw74).ArraySet1(_391_v32, 4)
			(_nw74).ArraySet1(_391_v32, 5)
			(_nw74).ArraySet1(_391_v32, 6)
			(_nw74).ArraySet1(_391_v32, 7)
			(_nw74).ArraySet1(_391_v32, 8)
			(_nw74).ArraySet1(_391_v32, 9)
			(_nw74).ArraySet1(_391_v32, 10)
			(_nw74).ArraySet1(_391_v32, 11)
			(_nw74).ArraySet1(_391_v32, 12)
			(_nw74).ArraySet1(_391_v32, 13)
			(_nw74).ArraySet1(_391_v32, 14)
			_392_v33 = _nw74
			var _393_v34 _dafny.Set
			_ = _393_v34
			_393_v34 = _dafny.SetOf(_392_v33)
			_393_v34 = (_393_v34).Intersection(_393_v34)
			var _394_v35 _dafny.Array
			_ = _394_v35
			var _nwElement0_15 _dafny.Int = _372_v14
			_ = _nwElement0_15
			var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.One)
			_ = _nw75
			(_nw75).ArraySet1(_nwElement0_15, 0)
			_394_v35 = _nw75
			var _395_v36 _dafny.Set
			_ = _395_v36
			_395_v36 = _dafny.SetOf(_372_v14)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_394_v35), 0))
			_ = _index43
			(_394_v35).ArraySet1(((_395_v36).Difference(_395_v36)).Cardinality(), (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_394_v35), 0))
			_ = _index44
			(_394_v35).ArraySet1(_372_v14, (_index44).Int())
		}
		var _396_v37 _dafny.Int
		_ = _396_v37
		_396_v37 = _dafny.IntOfInt64(534)
		var _397_v38 _dafny.MultiSet
		_ = _397_v38
		_397_v38 = _dafny.MultiSetOf(_355_v1)
		var _398_v39 _dafny.Set
		_ = _398_v39
		_398_v39 = _dafny.SetOf(Companion_D0_.Create_DC2_(_355_v1), Companion_D0_.Create_DC2_(_355_v1))
		var _399_v40 _dafny.Sequence
		_ = _399_v40
		_399_v40 = _dafny.SeqOf(_398_v39, _398_v39)
		var _400_v41 _dafny.Map
		_ = _400_v41
		_400_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_v1, (_dafny.Zero).Minus(_396_v37))
		var _401_v42 _dafny.Sequence
		_ = _401_v42
		_401_v42 = _dafny.UnicodeSeqOfUtf8Bytes("cqusjevmf")
		var _402_v43 _dafny.Array
		_ = _402_v43
		var _nwElement0_16 _dafny.Int = _dafny.IntOfInt64(195)
		_ = _nwElement0_16
		var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(20))
		_ = _nw76
		(_nw76).ArraySet1(_nwElement0_16, 0)
		(_nw76).ArraySet1(_396_v37, 1)
		(_nw76).ArraySet1(_396_v37, 2)
		(_nw76).ArraySet1(_396_v37, 3)
		(_nw76).ArraySet1(Companion_Default___.Fm1(_396_v37, globalState), 4)
		(_nw76).ArraySet1(_396_v37, 5)
		(_nw76).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_397_v38, _396_v37)).Cardinality(), 6)
		(_nw76).ArraySet1(_396_v37, 7)
		(_nw76).ArraySet1(_396_v37, 8)
		(_nw76).ArraySet1(_396_v37, 9)
		(_nw76).ArraySet1(_dafny.IntOfUint32((_399_v40).Cardinality()), 10)
		(_nw76).ArraySet1((_396_v37).Plus(_396_v37), 11)
		(_nw76).ArraySet1(_396_v37, 12)
		(_nw76).ArraySet1(_396_v37, 13)
		(_nw76).ArraySet1(_396_v37, 14)
		(_nw76).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_400_v41).Contains(_355_v1) {
				return (_400_v41).Get(_355_v1).(_dafny.Int)
			}
			return _396_v37
		})(), _dafny.IntOfUint32((_401_v42).Cardinality())), 15)
		(_nw76).ArraySet1(_396_v37, 16)
		(_nw76).ArraySet1(_396_v37, 17)
		(_nw76).ArraySet1(_396_v37, 18)
		(_nw76).ArraySet1(_396_v37, 19)
		_402_v43 = _nw76
		_402_v43 = _402_v43
		var _403_v44 _dafny.Array
		_ = _403_v44
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_14
		var _nw77 _dafny.Array
		_ = _nw77
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw77 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) D0 = (func(_404_v1 bool) func(_dafny.Int) D0 {
				return func(_405_i3 _dafny.Int) D0 {
					return Companion_D0_.Create_DC3_(_404_v1)
				}
			})(_355_v1)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw77 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw77).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw77).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_403_v44 = _nw77
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_403_v44), 0))
		_ = _index45
		(_403_v44).ArraySet1(Companion_Default___.Fm9((_dafny.MultiSetOf(_396_v37, _396_v37, _dafny.IntOfUint32((_401_v42).Cardinality()), _396_v37, _396_v37)).Cardinality(), globalState), (_index45).Int())
		var _406_v45 _dafny.Map
		_ = _406_v45
		_406_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_396_v37, _355_v1)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_403_v44), 0))
		_ = _index46
		var _rhs53 bool = _355_v1
		_ = _rhs53
		var _rhs54 D0 = Companion_D0_.Create_DC3_((_355_v1) || (_355_v1))
		_ = _rhs54
		var _rhs55 _dafny.Int = (_dafny.Zero).Minus(_396_v37)
		_ = _rhs55
		var _rhs56 _dafny.Int = (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(342), _355_v1)).Merge(_406_v45)).Cardinality())
		_ = _rhs56
		var _lhs29 *GlobalState = globalState
		_ = _lhs29
		var _lhs30 _dafny.Array = _403_v44
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_403_v44), 0))
		_ = _lhs31
		var _lhs32 *GlobalState = globalState
		_ = _lhs32
		var _lhs33 *GlobalState = globalState
		_ = _lhs33
		_lhs29.F0 = _rhs53
		(_lhs30).ArraySet1(_rhs54, (_lhs31).Int())
		_lhs32.F1 = _rhs55
		_lhs33.F1 = _rhs56
		_396_v37 = _396_v37
		r0 = _396_v37
		r1 = _354_v0
		r2 = (_403_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_403_v44), 0))).Int()).(D0)
		return r0, r1, r2
	}
}
func (_this *C1) M1(p0 bool, p1 _dafny.Set, p2 _dafny.Sequence, p3 D0, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _407_v0 _dafny.Array
		_ = _407_v0
		var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
		_ = _nw78
		_407_v0 = _nw78
		var _408_v1 *C0
		_ = _408_v1
		var _nw79 *C0 = New_C0_()
		_ = _nw79
		_nw79.Ctor__(_407_v0)
		_408_v1 = _nw79
		var _409_v2 _dafny.Int
		_ = _409_v2
		_409_v2 = _dafny.IntOfInt64(652)
		var _410_v3 _dafny.Map
		_ = _410_v3
		_410_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _409_v2)
		var _411_v4 _dafny.Sequence
		_ = _411_v4
		_411_v4 = _dafny.SeqOf(_408_v1)
		_410_v3 = (_410_v3).Update(_408_v1, (func() _dafny.Int {
			if p0 {
				return _dafny.IntOfInt64(824)
			}
			return _dafny.IntOfUint32((_411_v4).Cardinality())
		})())
		var _412_i0 _dafny.Int
		_ = _412_i0
		_412_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_412_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_412_i0 = (_412_i0).Plus(_dafny.One)
					var _413_v5 _dafny.Map
					_ = _413_v5
					_413_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_409_v2), Companion_D0_.Create_DC0_(_409_v2))
					var _414_v6 _dafny.Sequence
					_ = _414_v6
					_414_v6 = _dafny.SeqOf(_413_v5)
					var _415_v7 _dafny.Sequence
					_ = _415_v7
					_415_v7 = _dafny.SeqOf(_409_v2, _409_v2)
					var _416_v8 _dafny.MultiSet
					_ = _416_v8
					_416_v8 = _dafny.MultiSetOf((_414_v6).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_415_v7).Select((Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_415_v7).Cardinality()))).Uint32()).(_dafny.Int), _409_v2)).Cardinality(), _dafny.IntOfUint32((_414_v6).Cardinality()))).Uint32()).(_dafny.Map))
					_416_v8 = _416_v8
					(globalState).F1 = (_409_v2).Times(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-537))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}((func(_417_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_418_i1 _dafny.Int) _dafny.Int {
							return _417_v2
						}
					})(_409_v2)))).Cardinality())), _409_v2))
					var _419_v9 *C0
					_ = _419_v9
					var _nw80 *C0 = New_C0_()
					_ = _nw80
					_nw80.Ctor__(_407_v0)
					_419_v9 = _nw80
					var _420_v10 _dafny.Array
					_ = _420_v10
					var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
					_ = _nw81
					_420_v10 = _nw81
					var _421_v11 _dafny.Sequence
					_ = _421_v11
					_421_v11 = _dafny.SeqOf(p0, p0, false, p0)
					var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_420_v10), 0))
					_ = _index47
					(_420_v10).ArraySet1(_dafny.IntOfUint32((_421_v11).Cardinality()), (_index47).Int())
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_420_v10), 0))
					_ = _index48
					(_420_v10).ArraySet1(_dafny.IntOfUint32((_421_v11).Cardinality()), (_index48).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _422_v12 _dafny.Array
		_ = _422_v12
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_15
		var _nw82 _dafny.Array
		_ = _nw82
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw82 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = (func(_423_p0 bool) func(_dafny.Int) bool {
				return func(_424_i2 _dafny.Int) bool {
					return _423_p0
				}
			})(p0)
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
		_422_v12 = _nw82
		var _425_v13 _dafny.Map
		_ = _425_v13
		_425_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality()), _422_v12)
		_425_v13 = (_425_v13).Update(_409_v2, (func() _dafny.Array {
			if p0 {
				return _422_v12
			}
			return _422_v12
		})())
		var _426_v14 _dafny.CodePoint
		_ = _426_v14
		_426_v14 = _dafny.CodePoint('x')
		var _427_v15 D0
		_ = _427_v15
		_427_v15 = Companion_D0_.Create_DC1_(_426_v14)
		var _428_v16 _dafny.Sequence
		_ = _428_v16
		_428_v16 = _dafny.SeqOf(Companion_Default___.Fm1((_dafny.MultiSetOf(Companion_D0_.Create_DC1_(_426_v14), _427_v15)).Cardinality(), globalState))
		if _dafny.Companion_Sequence_.Contains(_428_v16, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bbijgu")).Cardinality())).Times(_dafny.IntOfUint32((p2).Cardinality()))) {
			var _429_v17 _dafny.Sequence
			_ = _429_v17
			_429_v17 = _dafny.SeqOf(p0, p0, false)
			_429_v17 = _dafny.Companion_Sequence_.Concatenate(_429_v17, (func() _dafny.Sequence {
				if !(p0) {
					return _429_v17
				}
				return _429_v17
			})())
			var _430_v18 D0
			_ = _430_v18
			_430_v18 = Companion_D0_.Create_DC2_(p0)
			(globalState).F0 = (_430_v18).Dtor_cf2()
			var _source4 D0 = Companion_D0_.Create_DC3_(Companion_Default___.Fm0(_dafny.IntOfUint32((_428_v16).Cardinality()), globalState))
			_ = _source4
			if _source4.Is_DC1() {
				var _431___mcc_h0 _dafny.CodePoint = _source4.Get_().(D0_DC1).Cf1
				_ = _431___mcc_h0
				var _432_cf1 _dafny.CodePoint = _431___mcc_h0
				_ = _432_cf1
				var _433_v19 _dafny.Array
				_ = _433_v19
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_16
				var _nw83 _dafny.Array
				_ = _nw83
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw83 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = (func(_434_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_435_i3 _dafny.Int) _dafny.Int {
							return (_435_i3).Plus(_434_v2)
						}
					})(_409_v2)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw83 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw83).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw83).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_433_v19 = _nw83
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_433_v19), 0))
				_ = _index49
				(_433_v19).ArraySet1((_dafny.Zero).Minus((_409_v2).Minus(_dafny.IntOfInt64(-624))), (_index49).Int())
				var _436_v20 _dafny.Sequence
				_ = _436_v20
				_436_v20 = _dafny.SeqOf(_433_v19)
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_433_v19), 0))
				_ = _index50
				(_433_v19).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_433_v19), _dafny.Companion_Sequence_.Concatenate(_436_v20, _436_v20))).Cardinality()), (_index50).Int())
				var _437_v21 _dafny.Map
				_ = _437_v21
				_437_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_410_v3, p0)
				_437_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_410_v3, p0)
				r0 = Companion_Default___.SafeModuloInt((_433_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_433_v19), 0))).Int()).(_dafny.Int), _409_v2)
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_422_v12), 0))
				_ = _index51
				(_422_v12).ArraySet1(p0, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_422_v12), 0))
				_ = _index52
				var _rhs57 _dafny.CodePoint = (_this).Fm3(((_433_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_433_v19), 0))).Int()).(_dafny.Int)).Cmp((_433_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_433_v19), 0))).Int()).(_dafny.Int)) != 0, globalState)
				_ = _rhs57
				var _rhs58 bool = p0
				_ = _rhs58
				var _rhs59 bool = p0
				_ = _rhs59
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				var _lhs35 _dafny.Array = _422_v12
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_422_v12), 0))
				_ = _lhs36
				_432_cf1 = _rhs57
				_lhs34.F0 = _rhs58
				(_lhs35).ArraySet1(_rhs59, (_lhs36).Int())
			} else if _source4.Is_DC2() {
				var _438___mcc_h1 bool = _source4.Get_().(D0_DC2).Cf2
				_ = _438___mcc_h1
				var _439_cf2 bool = _438___mcc_h1
				_ = _439_cf2
				var _440_v22 _dafny.Sequence
				_ = _440_v22
				_440_v22 = _dafny.SeqOf(Companion_Default___.Fm11(p0, p0, Companion_D0_.Create_DC3_(p0), globalState), _429_v17, _429_v17)
				var _441_v23 _dafny.MultiSet
				_ = _441_v23
				_441_v23 = _dafny.MultiSetOf(_dafny.IntOfUint32((_440_v22).Cardinality()))
				var _442_v24 _dafny.Map
				_ = _442_v24
				_442_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v2, (_441_v23).Cardinality())
				_442_v24 = (_442_v24).Update(((_dafny.Zero).Minus(_409_v2)).Plus(_409_v2), Companion_Default___.Fm1(_409_v2, globalState))
				var _443_v25 *C0
				_ = _443_v25
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__((_408_v1).F6())
				_443_v25 = _nw84
				_426_v14 = _dafny.CodePoint('v')
				var _rhs60 _dafny.Int = _409_v2
				_ = _rhs60
				var _rhs61 _dafny.Int = (_409_v2).Minus((_409_v2).Times(_409_v2))
				_ = _rhs61
				var _lhs37 *GlobalState = globalState
				_ = _lhs37
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				_lhs37.F3 = _rhs60
				_lhs38.F1 = _rhs61
			} else if _source4.Is_DC3() {
				var _444___mcc_h2 bool = _source4.Get_().(D0_DC3).Cf3
				_ = _444___mcc_h2
				var _445_cf3 bool = _444___mcc_h2
				_ = _445_cf3
				(globalState).F0 = p0
				var _446_v26 _dafny.Set
				_ = _446_v26
				_446_v26 = _dafny.SetOf(_408_v1)
				var _447_v27 _dafny.Set
				_ = _447_v27
				_447_v27 = _dafny.SetOf(_446_v26)
				var _448_v28 _dafny.MultiSet
				_ = _448_v28
				_448_v28 = _dafny.MultiSetOf(_445_cf3)
				var _449_v29 _dafny.Map
				_ = _449_v29
				_449_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _448_v28)
				var _450_v30 D1
				_ = _450_v30
				_450_v30 = Companion_D1_.Create_DC4_(_448_v28)
				var _451_v31 D2
				_ = _451_v31
				_451_v31 = Companion_D2_.Create_DC7_(_408_v1)
				var _452_v32 _dafny.Map
				_ = _452_v32
				_452_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_426_v14, (_451_v31).Dtor_cf7())
				var _453_v33 _dafny.Array
				_ = _453_v33
				var _nwElement0_17 _dafny.Int = _409_v2
				_ = _nwElement0_17
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(27))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_17, 0)
				(_nw85).ArraySet1((_409_v2).Plus((Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("reo")).Cardinality()))).Dtor_cf0()), 1)
				(_nw85).ArraySet1(_409_v2, 2)
				(_nw85).ArraySet1(_409_v2, 3)
				(_nw85).ArraySet1(_409_v2, 4)
				(_nw85).ArraySet1(_409_v2, 5)
				(_nw85).ArraySet1(_409_v2, 6)
				(_nw85).ArraySet1(((_447_v27).Intersection(_447_v27)).Cardinality(), 7)
				(_nw85).ArraySet1(_409_v2, 8)
				(_nw85).ArraySet1(_409_v2, 9)
				(_nw85).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_409_v2), _409_v2), 10)
				(_nw85).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_428_v16).Cardinality()), _409_v2), 11)
				(_nw85).ArraySet1(((_428_v16).Select((Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_428_v16).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_409_v2), 12)
				(_nw85).ArraySet1(Companion_Default___.SafeModuloInt(_409_v2, (func() _dafny.Int {
					if (_448_v28).Contains(_445_cf3) {
						return (_448_v28).Multiplicity(_445_cf3)
					}
					return _409_v2
				})()), 13)
				(_nw85).ArraySet1(_409_v2, 14)
				(_nw85).ArraySet1(_409_v2, 15)
				(_nw85).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
					if (_449_v29).Contains(_408_v1) {
						return (_449_v29).Get(_408_v1).(_dafny.MultiSet)
					}
					return (_450_v30).Dtor_cf4()
				})(), _422_v12)).Cardinality(), 16)
				(_nw85).ArraySet1(Companion_Default___.Fm1(_409_v2, globalState), 17)
				(_nw85).ArraySet1(Companion_Default___.SafeDivisionInt(_409_v2, _409_v2), 18)
				(_nw85).ArraySet1(_409_v2, 19)
				(_nw85).ArraySet1(_409_v2, 20)
				(_nw85).ArraySet1(((_452_v32).Merge(_452_v32)).Cardinality(), 21)
				(_nw85).ArraySet1(_409_v2, 22)
				(_nw85).ArraySet1(_409_v2, 23)
				(_nw85).ArraySet1(_409_v2, 24)
				(_nw85).ArraySet1(_409_v2, 25)
				(_nw85).ArraySet1(_dafny.IntOfInt64(814), 26)
				_453_v33 = _nw85
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_453_v33), 0))
				_ = _index53
				(_453_v33).ArraySet1(_409_v2, (_index53).Int())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_453_v33), 0))
				_ = _index54
				(_453_v33).ArraySet1(_409_v2, (_index54).Int())
				var _454_v34 _dafny.Sequence
				_ = _454_v34
				_454_v34 = _dafny.UnicodeSeqOfUtf8Bytes("hkcpig")
				_454_v34 = (func() _dafny.Sequence {
					if (_409_v2).Cmp((_453_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_453_v33), 0))).Int()).(_dafny.Int)) <= 0 {
						return _454_v34
					}
					return _454_v34
				})()
				var _455_v35 _dafny.MultiSet
				_ = _455_v35
				_455_v35 = _dafny.MultiSetOf(_428_v16, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm10(globalState), (Companion_Default___.SafeIndex((_453_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_453_v33), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality()))).Uint32(), _409_v2), _428_v16)
				_455_v35 = ((_455_v35).Difference(_455_v35)).Intersection((func() _dafny.MultiSet {
					if p0 {
						return _455_v35
					}
					return _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_428_v16, (Companion_Default___.SafeIndex((p1).Cardinality(), _dafny.IntOfUint32((_428_v16).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(_445_cf3)).Cardinality())), _dafny.SeqOf((_453_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_453_v33), 0))).Int()).(_dafny.Int), _409_v2), _428_v16)
				})())
			} else {
				var _456___mcc_h3 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
				_ = _456___mcc_h3
				var _457_cf0 _dafny.Int = _456___mcc_h3
				_ = _457_cf0
				var _458_v36 _dafny.Set
				_ = _458_v36
				_458_v36 = _dafny.SetOf(_457_cf0, _409_v2)
				var _459_v37 _dafny.Map
				_ = _459_v37
				_459_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_cf0, (_458_v36).Cardinality())
				var _460_v38 _dafny.Sequence
				_ = _460_v38
				_460_v38 = _dafny.SeqOf(_459_v37)
				var _rhs62 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(903), _409_v2)
				_ = _rhs62
				var _rhs63 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_408_v1), _411_v4)
				_ = _rhs63
				var _rhs64 _dafny.Sequence = (Companion_D3_.Create_DC9_(_460_v38)).Dtor_cf10()
				_ = _rhs64
				var _rhs65 _dafny.Array = _422_v12
				_ = _rhs65
				var _lhs39 *GlobalState = globalState
				_ = _lhs39
				_lhs39.F1 = _rhs62
				_411_v4 = _rhs63
				_460_v38 = _rhs64
				_422_v12 = _rhs65
				var _461_v39 _dafny.Array
				_ = _461_v39
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
				_ = _nw86
				_461_v39 = _nw86
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_461_v39), 0))
				_ = _index55
				(_461_v39).ArraySet1CodePoint(_426_v14, (_index55).Int())
				var _462_v40 _dafny.Map
				_ = _462_v40
				_462_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _426_v14)
				var _463_v41 D2
				_ = _463_v41
				_463_v41 = Companion_D2_.Create_DC7_(_408_v1)
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_461_v39), 0))
				_ = _index56
				(_461_v39).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_462_v40).Contains((_463_v41).Dtor_cf7()) {
						return (_462_v40).Get((_463_v41).Dtor_cf7()).(_dafny.CodePoint)
					}
					return _426_v14
				})(), (_index56).Int())
				(globalState).F1 = _409_v2
				var _464_v42 _dafny.MultiSet
				_ = _464_v42
				_464_v42 = _dafny.MultiSetOf(p0, p0, !(p0))
				var _465_v43 D1
				_ = _465_v43
				_465_v43 = Companion_D1_.Create_DC4_(_dafny.MultiSetFromSeq(_429_v17))
				var _466_v44 _dafny.Array
				_ = _466_v44
				var _nwElement0_18 _dafny.MultiSet = (_464_v42).Difference(_dafny.MultiSetOf(false))
				_ = _nwElement0_18
				var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(22))
				_ = _nw87
				(_nw87).ArraySet1(_nwElement0_18, 0)
				(_nw87).ArraySet1((_464_v42).Union(_dafny.MultiSetFromSeq(_429_v17)), 1)
				(_nw87).ArraySet1((func() _dafny.MultiSet {
					if p0 {
						return _464_v42
					}
					return _464_v42
				})(), 2)
				(_nw87).ArraySet1(_dafny.MultiSetOf(p0, p0, Companion_Default___.Fm0(_409_v2, globalState)), 3)
				(_nw87).ArraySet1(_464_v42, 4)
				(_nw87).ArraySet1((_465_v43).Dtor_cf4(), 5)
				(_nw87).ArraySet1(_464_v42, 6)
				(_nw87).ArraySet1(_dafny.MultiSetOf(p0), 7)
				(_nw87).ArraySet1(_464_v42, 8)
				(_nw87).ArraySet1((_dafny.MultiSetFromSeq(_429_v17)).Intersection(_464_v42), 9)
				(_nw87).ArraySet1(_464_v42, 10)
				(_nw87).ArraySet1(_464_v42, 11)
				(_nw87).ArraySet1(_dafny.MultiSetOf(p0, p0), 12)
				(_nw87).ArraySet1(_464_v42, 13)
				(_nw87).ArraySet1((_464_v42).Intersection(_464_v42), 14)
				(_nw87).ArraySet1((_464_v42).Union(_464_v42), 15)
				(_nw87).ArraySet1((_464_v42).Difference(_dafny.MultiSetOf(true)), 16)
				(_nw87).ArraySet1(Companion_Default___.Fm12(_409_v2, globalState), 17)
				(_nw87).ArraySet1(_dafny.MultiSetOf(true), 18)
				(_nw87).ArraySet1(_464_v42, 19)
				(_nw87).ArraySet1((_464_v42).Difference(_464_v42), 20)
				(_nw87).ArraySet1((_dafny.MultiSetOf(false, !(p0))).Union(_464_v42), 21)
				_466_v44 = _nw87
				var _467_v45 _dafny.Map
				_ = _467_v45
				_467_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _464_v42)
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_466_v44), 0))
				_ = _index57
				(_466_v44).ArraySet1((func() _dafny.MultiSet {
					if (_467_v45).Contains(_408_v1) {
						return (_467_v45).Get(_408_v1).(_dafny.MultiSet)
					}
					return _464_v42
				})(), (_index57).Int())
				var _468_v46 D0
				_ = _468_v46
				_468_v46 = Companion_D0_.Create_DC0_(_457_cf0)
				var _pat_let_tv9 = _457_cf0
				_ = _pat_let_tv9
				var _469_v47 _dafny.Array
				_ = _469_v47
				var _nwElement0_19 D0 = Companion_D0_.Create_DC0_(_409_v2)
				_ = _nwElement0_19
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(11))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_19, 0)
				(_nw88).ArraySet1(_468_v46, 1)
				(_nw88).ArraySet1(func(_pat_let4_0 D0) D0 {
					return func(_470_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let5_0 _dafny.Int) D0 {
							return func(_471_dt__update_hcf0_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC0_(_471_dt__update_hcf0_h0)
							}(_pat_let5_0)
						}(_pat_let_tv9)
					}(_pat_let4_0)
				}(_468_v46), 2)
				(_nw88).ArraySet1(_468_v46, 3)
				(_nw88).ArraySet1(Companion_Default___.Fm13(false, p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v2, _457_cf0), p0, globalState), 4)
				(_nw88).ArraySet1(_468_v46, 5)
				(_nw88).ArraySet1(_468_v46, 6)
				(_nw88).ArraySet1(_468_v46, 7)
				(_nw88).ArraySet1(_468_v46, 8)
				(_nw88).ArraySet1(_468_v46, 9)
				(_nw88).ArraySet1(Companion_D0_.Create_DC0_(_409_v2), 10)
				_469_v47 = _nw88
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_469_v47), 0))
				_ = _index58
				(_469_v47).ArraySet1(_468_v46, (_index58).Int())
				var _472_v48 D0
				_ = _472_v48
				_472_v48 = Companion_D0_.Create_DC3_((_409_v2).Cmp(_457_cf0) > 0)
				var _473_v49 _dafny.Map
				_ = _473_v49
				_473_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v2, _dafny.MultiSetOf(!(p0), p0))
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_466_v44), 0))
				_ = _index59
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_469_v47), 0))
				_ = _index60
				var _rhs66 _dafny.MultiSet = (func() _dafny.MultiSet {
					if (_473_v49).Contains(_409_v2) {
						return (_473_v49).Get(_409_v2).(_dafny.MultiSet)
					}
					return (_464_v42).Difference(_464_v42)
				})()
				_ = _rhs66
				var _rhs67 _dafny.Int = _457_cf0
				_ = _rhs67
				var _rhs68 D0 = _468_v46
				_ = _rhs68
				var _rhs69 D0 = Companion_D0_.Create_DC3_(p0)
				_ = _rhs69
				var _rhs70 _dafny.Int = Companion_Default___.Fm1(_409_v2, globalState)
				_ = _rhs70
				var _lhs40 _dafny.Array = _466_v44
				_ = _lhs40
				var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_466_v44), 0))
				_ = _lhs41
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				var _lhs43 _dafny.Array = _469_v47
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_469_v47), 0))
				_ = _lhs44
				(_lhs40).ArraySet1(_rhs66, (_lhs41).Int())
				_lhs42.F3 = _rhs67
				(_lhs43).ArraySet1(_rhs68, (_lhs44).Int())
				_472_v48 = _rhs69
				r0 = _rhs70
			}
			if (_dafny.IntOfUint32((_429_v17).Cardinality())).Cmp(_409_v2) >= 0 {
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_422_v12), 0))
				_ = _index61
				(_422_v12).ArraySet1(p0, (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_422_v12), 0))
				_ = _index62
				(_422_v12).ArraySet1(p0, (_index62).Int())
				var _474_v50 _dafny.Array
				_ = _474_v50
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_17
				var _nw89 _dafny.Array
				_ = _nw89
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw89 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = (func(_475_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_476_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_476_i4, _475_v2)
						}
					})(_409_v2)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw89 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw89).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw89).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_474_v50 = _nw89
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_474_v50), 0))
				_ = _index63
				(_474_v50).ArraySet1(_409_v2, (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_474_v50), 0))
				_ = _index64
				(_474_v50).ArraySet1((_409_v2).Minus(_409_v2), (_index64).Int())
				(globalState).F0 = !(true)
				_407_v0 = (_408_v1).F6()
				(globalState).F0 = p0
			} else {
				(globalState).F0 = p0
				(globalState).F0 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_429_v17, _429_v17), _429_v17)
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_422_v12), 0))
				_ = _index65
				(_422_v12).ArraySet1(p0, (_index65).Int())
				var _477_v51 D4
				_ = _477_v51
				_477_v51 = Companion_D4_.Create_DC13_(_422_v12)
				var _pat_let_tv10 = _477_v51
				_ = _pat_let_tv10
				var _pat_let_tv11 = _422_v12
				_ = _pat_let_tv11
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_422_v12), 0))
				_ = _index66
				var _rhs71 _dafny.Array = (func(_pat_let6_0 D4) D4 {
					return func(_478_dt__update__tmp_h1 D4) D4 {
						return func(_pat_let7_0 _dafny.Array) D4 {
							return func(_481_dt__update_hcf19_h0 _dafny.Array) D4 {
								return Companion_D4_.Create_DC13_(_481_dt__update_hcf19_h0)
							}(_pat_let7_0)
						}((func(_pat_let8_0 D4) D4 {
							return func(_479_dt__update__tmp_h2 D4) D4 {
								return func(_pat_let9_0 _dafny.Array) D4 {
									return func(_480_dt__update_hcf19_h1 _dafny.Array) D4 {
										return Companion_D4_.Create_DC13_(_480_dt__update_hcf19_h1)
									}(_pat_let9_0)
								}(_pat_let_tv11)
							}(_pat_let8_0)
						}(_pat_let_tv10)).Dtor_cf19())
					}(_pat_let6_0)
				}(_477_v51)).Dtor_cf19()
				_ = _rhs71
				var _rhs72 bool = (_409_v2).Cmp(_409_v2) > 0
				_ = _rhs72
				var _rhs73 bool = p0
				_ = _rhs73
				var _lhs45 _dafny.Array = _422_v12
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_422_v12), 0))
				_ = _lhs46
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				_422_v12 = _rhs71
				(_lhs45).ArraySet1(_rhs72, (_lhs46).Int())
				_lhs47.F0 = _rhs73
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_422_v12), 0))
				_ = _index67
				var _rhs74 _dafny.Array = _407_v0
				_ = _rhs74
				var _rhs75 bool = (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool)
				_ = _rhs75
				var _lhs48 _dafny.Array = _422_v12
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_422_v12), 0))
				_ = _lhs49
				_407_v0 = _rhs74
				(_lhs48).ArraySet1(_rhs75, (_lhs49).Int())
				(globalState).F0 = p0
			}
			var _482_v52 _dafny.Sequence
			_ = _482_v52
			_482_v52 = _dafny.SeqOf(_426_v14)
			_482_v52 = _dafny.SeqOf(_426_v14)
		} else {
			(globalState).F1 = ((_dafny.Zero).Minus(_409_v2)).Minus(_409_v2)
			(globalState).F0 = p0
			var _483_v53 _dafny.Map
			_ = _483_v53
			_483_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v2, _dafny.IntOfInt64(547))
			r0 = (func() _dafny.Int {
				if (_483_v53).Contains(_409_v2) {
					return (_483_v53).Get(_409_v2).(_dafny.Int)
				}
				return _409_v2
			})()
			var _484_v54 _dafny.Sequence
			_ = _484_v54
			_484_v54 = _dafny.SeqOf(p0)
			var _485_v55 _dafny.Map
			_ = _485_v55
			_485_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_v54, Companion_Default___.Fm4(p0, _409_v2, _409_v2, _426_v14, globalState))
			_485_v55 = (_485_v55).Update(Companion_Default___.Fm11(p0, p0, p3, globalState), _dafny.Companion_Sequence_.Concatenate(p2, p2))
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_422_v12), 0))
			_ = _index68
			(_422_v12).ArraySet1(true, (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_422_v12), 0))
			_ = _index69
			(_422_v12).ArraySet1(!(p0), (_index69).Int())
		}
		var _486_v56 D0
		_ = _486_v56
		_486_v56 = Companion_D0_.Create_DC2_(p0)
		var _487_v57 _dafny.Sequence
		_ = _487_v57
		_487_v57 = _dafny.SeqOf(_486_v56, _486_v56)
		var _488_v58 _dafny.Array
		_ = _488_v58
		var _nwElement0_20 _dafny.Sequence = _487_v57
		_ = _nwElement0_20
		var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(10))
		_ = _nw90
		(_nw90).ArraySet1(_nwElement0_20, 0)
		(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14(p0, globalState), _dafny.SeqOf(_486_v56)), 1)
		(_nw90).ArraySet1(_487_v57, 2)
		(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_487_v57, _487_v57), 3)
		(_nw90).ArraySet1(Companion_Default___.Fm14(p0, globalState), 4)
		(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(69))).Uint32(), func(coer28 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_489_v56 D0) func(_dafny.Int) D0 {
			return func(_490_i5 _dafny.Int) D0 {
				return _489_v56
			}
		})(_486_v56))), _487_v57), 5)
		(_nw90).ArraySet1(_487_v57, 6)
		(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_487_v57, _487_v57), 7)
		(_nw90).ArraySet1(_dafny.SeqOf(_486_v56, _486_v56), 8)
		(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_487_v57, _487_v57), 9)
		_488_v58 = _nw90
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_488_v58), 0))
		_ = _index70
		(_488_v58).ArraySet1(_487_v57, (_index70).Int())
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
		_ = _index71
		(_422_v12).ArraySet1((_dafny.IntOfUint32((Companion_Default___.Fm4(p0, _409_v2, _409_v2, _426_v14, globalState)).Cardinality())).Cmp((_dafny.Zero).Minus(_409_v2)) <= 0, (_index71).Int())
		var _491_v59 _dafny.Map
		_ = _491_v59
		_491_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _487_v57)
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_488_v58), 0))
		_ = _index72
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
		_ = _index73
		var _rhs76 _dafny.Int = Companion_Default___.SafeModuloInt(_409_v2, _409_v2)
		_ = _rhs76
		var _rhs77 _dafny.Sequence = (func() _dafny.Sequence {
			if (_491_v59).Contains(!(p0)) {
				return (_491_v59).Get(!(p0)).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(50))).Uint32(), func(coer29 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_492_v56 D0) func(_dafny.Int) D0 {
				return func(_493_i6 _dafny.Int) D0 {
					return _492_v56
				}
			})(_486_v56))), _dafny.SeqOf(_486_v56))
		})()
		_ = _rhs77
		var _rhs78 bool = p0
		_ = _rhs78
		var _rhs79 bool = (p0) || (_dafny.Companion_Sequence_.IsProperPrefixOf(p2, _dafny.UnicodeSeqOfUtf8Bytes("y")))
		_ = _rhs79
		var _lhs50 *GlobalState = globalState
		_ = _lhs50
		var _lhs51 _dafny.Array = _488_v58
		_ = _lhs51
		var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_488_v58), 0))
		_ = _lhs52
		var _lhs53 _dafny.Array = _422_v12
		_ = _lhs53
		var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
		_ = _lhs54
		var _lhs55 *GlobalState = globalState
		_ = _lhs55
		_lhs50.F3 = _rhs76
		(_lhs51).ArraySet1(_rhs77, (_lhs52).Int())
		(_lhs53).ArraySet1(_rhs78, (_lhs54).Int())
		_lhs55.F0 = _rhs79
		if (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool) {
			var _494_v60 _dafny.MultiSet
			_ = _494_v60
			_494_v60 = _dafny.MultiSetOf(p0)
			var _pat_let_tv12 = _409_v2
			_ = _pat_let_tv12
			var _pat_let_tv13 = _494_v60
			_ = _pat_let_tv13
			_486_v56 = func(_pat_let10_0 D0) D0 {
				return func(_495_dt__update__tmp_h3 D0) D0 {
					return func(_pat_let11_0 bool) D0 {
						return func(_496_dt__update_hcf2_h0 bool) D0 {
							return Companion_D0_.Create_DC2_(_496_dt__update_hcf2_h0)
						}(_pat_let11_0)
					}(!((_pat_let_tv12).Cmp((_pat_let_tv13).Cardinality()) >= 0))
				}(_pat_let10_0)
			}(_486_v56)
			var _497_v61 _dafny.Array
			_ = _497_v61
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_18
			var _nw91 _dafny.Array
			_ = _nw91
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw91 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) _dafny.Sequence = (func(_498_p0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_499_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_498_p0, _498_p0)
					}
				})(p0)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw91 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw91).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw91).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_497_v61 = _nw91
			var _500_v62 _dafny.Map
			_ = _500_v62
			_500_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v2, _497_v61)
			_497_v61 = (func() _dafny.Array {
				if (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool) {
					return _497_v61
				}
				return (func() _dafny.Array {
					if (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool) {
						return (func() _dafny.Array {
							if (_500_v62).Contains(_409_v2) {
								return (_500_v62).Get(_409_v2).(_dafny.Array)
							}
							return _497_v61
						})()
					}
					return _497_v61
				})()
			})()
			if (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool) {
				(globalState).F1 = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(p2, p2))).Cardinality())
				var _501_v63 D2
				_ = _501_v63
				_501_v63 = Companion_D2_.Create_DC7_(_408_v1)
				var _rhs80 *C0 = _408_v1
				_ = _rhs80
				var _rhs81 D2 = _501_v63
				_ = _rhs81
				_408_v1 = _rhs80
				_501_v63 = _rhs81
				var _502_v64 _dafny.Sequence
				_ = _502_v64
				_502_v64 = _dafny.SeqOf(true, p0)
				var _503_v65 _dafny.Map
				_ = _503_v65
				_503_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_502_v64).Cardinality()), (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool))
				var _504_v66 _dafny.Array
				_ = _504_v66
				var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw92
				_504_v66 = _nw92
				var _505_v67 _dafny.Map
				_ = _505_v67
				_505_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_503_v65).Cardinality(), _504_v66)
				(globalState).F0 = ((_505_v67).Cardinality()).Cmp(_dafny.IntOfInt64(993)) < 0
				var _506_v68 D5
				_ = _506_v68
				_506_v68 = Companion_D5_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("xhrajd"))
				(globalState).F3 = ((func() _dafny.Int {
					if (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool) {
						return _409_v2
					}
					return _dafny.IntOfUint32(((_506_v68).Dtor_cf26()).Cardinality())
				})()).Minus(_409_v2)
				var _507_v69 *C0
				_ = _507_v69
				var _nw93 *C0 = New_C0_()
				_ = _nw93
				_nw93.Ctor__((_408_v1).F6())
				_507_v69 = _nw93
			} else {
				(globalState).F3 = Companion_Default___.SafeDivisionInt(_409_v2, _dafny.IntOfUint32((Companion_Default___.Fm4((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool), (_494_v60).Cardinality(), _409_v2, _426_v14, globalState)).Cardinality()))
				var _508_v70 _dafny.Sequence
				_ = _508_v70
				_508_v70 = _dafny.UnicodeSeqOfUtf8Bytes("fcqa")
				var _509_v71 D5
				_ = _509_v71
				_509_v71 = Companion_D5_.Create_DC18_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vcxcktchu")).Cardinality()), _409_v2)
				var _510_v72 _dafny.Sequence
				_ = _510_v72
				_510_v72 = _dafny.SeqOf(Companion_D5_.Create_DC18_(_409_v2, _409_v2), _509_v71, _509_v71, Companion_D5_.Create_DC18_(_409_v2, _dafny.IntOfInt64(-180)), _509_v71)
				var _511_v73 _dafny.Map
				_ = _511_v73
				_511_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm1(Companion_Default___.Fm1(_409_v2, globalState), globalState)).Minus(_409_v2), _510_v72)
				var _512_v74 _dafny.Sequence
				_ = _512_v74
				_512_v74 = _dafny.SeqOf(p0, p0)
				var _513_v75 _dafny.MultiSet
				_ = _513_v75
				_513_v75 = _dafny.MultiSetOf(_409_v2)
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _index74
				var _rhs82 _dafny.Int = (_408_v1).Fm7(_dafny.CodePoint('j'), globalState)
				_ = _rhs82
				var _rhs83 _dafny.Sequence = p2
				_ = _rhs83
				var _rhs84 _dafny.Int = _dafny.IntOfInt64(174)
				_ = _rhs84
				var _rhs85 _dafny.Map = Companion_Default___.Fm15(((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool)) == ((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool)), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_512_v74, (Companion_Default___.SafeIndex((_513_v75).Cardinality(), _dafny.IntOfUint32((_512_v74).Cardinality()))).Uint32(), true), _dafny.SeqOf(p0, p0, p0, (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool), !(p0))), _dafny.Companion_Sequence_.Concatenate(_512_v74, _512_v74), globalState)
				_ = _rhs85
				var _rhs86 bool = p0
				_ = _rhs86
				var _lhs56 *GlobalState = globalState
				_ = _lhs56
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				var _lhs58 _dafny.Array = _422_v12
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _lhs59
				_lhs56.F3 = _rhs82
				_508_v70 = _rhs83
				_lhs57.F3 = _rhs84
				_511_v73 = _rhs85
				(_lhs58).ArraySet1(_rhs86, (_lhs59).Int())
				_497_v61 = _497_v61
				var _514_v76 _dafny.Map
				_ = _514_v76
				_514_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v2, (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool))
				var _515_v77 _dafny.Map
				_ = _515_v77
				_515_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_514_v76).Cardinality())
				var _516_v78 _dafny.Array
				_ = _516_v78
				var _nwElement0_21 _dafny.Int = (_515_v77).Cardinality()
				_ = _nwElement0_21
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(19))
				_ = _nw94
				(_nw94).ArraySet1(_nwElement0_21, 0)
				(_nw94).ArraySet1((_dafny.Zero).Minus(_409_v2), 1)
				(_nw94).ArraySet1(_409_v2, 2)
				(_nw94).ArraySet1(_409_v2, 3)
				(_nw94).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), 4)
				(_nw94).ArraySet1(_409_v2, 5)
				(_nw94).ArraySet1(_409_v2, 6)
				(_nw94).ArraySet1((_513_v75).Cardinality(), 7)
				(_nw94).ArraySet1(_409_v2, 8)
				(_nw94).ArraySet1((_dafny.Zero).Minus(_409_v2), 9)
				(_nw94).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality()), globalState), 10)
				(_nw94).ArraySet1(_409_v2, 11)
				(_nw94).ArraySet1(_409_v2, 12)
				(_nw94).ArraySet1(_409_v2, 13)
				(_nw94).ArraySet1(_409_v2, 14)
				(_nw94).ArraySet1(_409_v2, 15)
				(_nw94).ArraySet1(_409_v2, 16)
				(_nw94).ArraySet1(_dafny.IntOfInt64(655), 17)
				(_nw94).ArraySet1(_409_v2, 18)
				_516_v78 = _nw94
				var _517_v79 _dafny.Sequence
				_ = _517_v79
				_517_v79 = _dafny.SeqOf(_516_v78)
				var _518_v80 _dafny.Set
				_ = _518_v80
				_518_v80 = _dafny.SetOf(_dafny.SetOf((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool), !(p0)))
				var _519_v81 _dafny.Map
				_ = _519_v81
				_519_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_518_v80).Cardinality(), _516_v78)
				var _520_v82 _dafny.Array
				_ = _520_v82
				var _nwElement0_22 _dafny.Sequence = _517_v79
				_ = _nwElement0_22
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(28))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_22, 0)
				(_nw95).ArraySet1(_517_v79, 1)
				(_nw95).ArraySet1(_517_v79, 2)
				(_nw95).ArraySet1(_517_v79, 3)
				(_nw95).ArraySet1(_517_v79, 4)
				(_nw95).ArraySet1(_517_v79, 5)
				(_nw95).ArraySet1(_517_v79, 6)
				(_nw95).ArraySet1(_dafny.SeqOf(_516_v78), 7)
				(_nw95).ArraySet1(_517_v79, 8)
				(_nw95).ArraySet1(_517_v79, 9)
				(_nw95).ArraySet1(_517_v79, 10)
				(_nw95).ArraySet1(_517_v79, 11)
				(_nw95).ArraySet1(_517_v79, 12)
				(_nw95).ArraySet1(_517_v79, 13)
				(_nw95).ArraySet1(_dafny.SeqOf(_516_v78, _516_v78), 14)
				(_nw95).ArraySet1(_517_v79, 15)
				(_nw95).ArraySet1(_517_v79, 16)
				(_nw95).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_517_v79, _dafny.SeqOf(_516_v78, _516_v78)), 17)
				(_nw95).ArraySet1(_517_v79, 18)
				(_nw95).ArraySet1(_517_v79, 19)
				(_nw95).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79), (Companion_Default___.SafeIndex((Companion_Default___.Fm16(globalState)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79)).Cardinality()))).Uint32(), (func() _dafny.Array {
					if (_519_v81).Contains(_dafny.IntOfInt64(992)) {
						return (_519_v81).Get(_dafny.IntOfInt64(992)).(_dafny.Array)
					}
					return _516_v78
				})()), 20)
				(_nw95).ArraySet1(_dafny.Companion_Sequence_.Update(_517_v79, (Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_517_v79).Cardinality()))).Uint32(), _516_v78), 21)
				(_nw95).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_512_v74).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79)).Cardinality()))).Uint32(), _516_v78), 22)
				(_nw95).ArraySet1(_dafny.SeqOf(_516_v78, _516_v78), 23)
				(_nw95).ArraySet1(_517_v79, 24)
				(_nw95).ArraySet1(_517_v79, 25)
				(_nw95).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79), 26)
				(_nw95).ArraySet1(_517_v79, 27)
				_520_v82 = _nw95
				var _521_v83 D6
				_ = _521_v83
				_521_v83 = Companion_D6_.Create_DC19_(_517_v79)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_520_v82), 0))
				_ = _index75
				(_520_v82).ArraySet1((_521_v83).Dtor_cf34(), (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_520_v82), 0))
				_ = _index76
				(_520_v82).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79), _517_v79), (Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_517_v79, _517_v79), _517_v79)).Cardinality()))).Uint32(), _516_v78), (_index76).Int())
				_515_v77 = (_515_v77).Update(!(p0), Companion_Default___.SafeModuloInt(_409_v2, _409_v2))
			}
			var _522_v84 _dafny.Sequence
			_ = _522_v84
			_522_v84 = _dafny.SeqOf(p0, (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool))
			var _523_v85 D6
			_ = _523_v85
			_523_v85 = Companion_D6_.Create_DC20_(_522_v84, p0)
			var _source5 D6 = _523_v85
			_ = _source5
			if _source5.Is_DC20() {
				var _524___mcc_h4 _dafny.Sequence = _source5.Get_().(D6_DC20).Cf35
				_ = _524___mcc_h4
				var _525___mcc_h5 bool = _source5.Get_().(D6_DC20).Cf36
				_ = _525___mcc_h5
				var _526_cf36 bool = _525___mcc_h5
				_ = _526_cf36
				var _527_cf35 _dafny.Sequence = _524___mcc_h4
				_ = _527_cf35
				var _528_v86 D6
				_ = _528_v86
				_528_v86 = Companion_D6_.Create_DC22_((func() _dafny.Int {
					if _526_cf36 {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nock")).Cardinality())
					}
					return _409_v2
				})(), Companion_Default___.Fm1(_dafny.IntOfUint32((p2).Cardinality()), globalState))
				var _529_v87 _dafny.Set
				_ = _529_v87
				_529_v87 = _dafny.SetOf(_408_v1)
				var _rhs87 _dafny.Int = (_529_v87).Cardinality()
				_ = _rhs87
				var _rhs88 bool = (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool)
				_ = _rhs88
				var _rhs89 D6 = Companion_D6_.Create_DC22_(Companion_Default___.SafeDivisionInt(_409_v2, _dafny.IntOfInt64(-337)), (_dafny.Zero).Minus(((_425_v13).Cardinality()).Plus(_409_v2)))
				_ = _rhs89
				_409_v2 = _rhs87
				_526_cf36 = _rhs88
				_528_v86 = _rhs89
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _index77
				(_422_v12).ArraySet1(!((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool)), (_index77).Int())
				_409_v2 = _409_v2
				(globalState).F1 = _409_v2
			} else if _source5.Is_DC21() {
				var _530___mcc_h6 _dafny.Int = _source5.Get_().(D6_DC21).Cf37
				_ = _530___mcc_h6
				var _531___mcc_h7 bool = _source5.Get_().(D6_DC21).Cf38
				_ = _531___mcc_h7
				var _532___mcc_h8 _dafny.Sequence = _source5.Get_().(D6_DC21).Cf39
				_ = _532___mcc_h8
				var _533___mcc_h9 bool = _source5.Get_().(D6_DC21).Cf40
				_ = _533___mcc_h9
				var _534_cf40 bool = _533___mcc_h9
				_ = _534_cf40
				var _535_cf39 _dafny.Sequence = _532___mcc_h8
				_ = _535_cf39
				var _536_cf38 bool = _531___mcc_h7
				_ = _536_cf38
				var _537_cf37 _dafny.Int = _530___mcc_h6
				_ = _537_cf37
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _index78
				(_422_v12).ArraySet1(_534_cf40, (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _index79
				(_422_v12).ArraySet1(_534_cf40, (_index79).Int())
				_534_cf40 = p0
				var _538_v88 _dafny.MultiSet
				_ = _538_v88
				_538_v88 = _dafny.MultiSetOf(_535_cf39, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("wa"), (Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wa")).Cardinality()))).Uint32(), _dafny.CodePoint('o')))
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _index80
				(_422_v12).ArraySet1(((_538_v88).Union(_538_v88)).IsSubsetOf(_538_v88), (_index80).Int())
			} else if _source5.Is_DC22() {
				var _539___mcc_h10 _dafny.Int = _source5.Get_().(D6_DC22).Cf41
				_ = _539___mcc_h10
				var _540___mcc_h11 _dafny.Int = _source5.Get_().(D6_DC22).Cf42
				_ = _540___mcc_h11
				var _541_cf42 _dafny.Int = _540___mcc_h11
				_ = _541_cf42
				var _542_cf41 _dafny.Int = _539___mcc_h10
				_ = _542_cf41
				var _543_v89 _dafny.Set
				_ = _543_v89
				_543_v89 = _dafny.SetOf(_541_cf42)
				var _544_v90 _dafny.Array
				_ = _544_v90
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_19
				var _nw96 _dafny.Array
				_ = _nw96
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw96 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Int = (func(_545_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_546_i8 _dafny.Int) _dafny.Int {
							return (_546_i8).Plus(_545_v2)
						}
					})(_409_v2)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw96 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw96).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw96).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_544_v90 = _nw96
				var _547_v91 _dafny.Set
				_ = _547_v91
				_547_v91 = _dafny.SetOf(_544_v90, _544_v90, _544_v90)
				var _548_v92 _dafny.Map
				_ = _548_v92
				_548_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_v89, (_547_v91).Cardinality())
				_548_v92 = (_548_v92).Update((_543_v89).Intersection(_543_v89), _409_v2)
				var _549_v93 _dafny.Array
				_ = _549_v93
				var _nwElement0_23 _dafny.Map = _410_v3
				_ = _nwElement0_23
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(26))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_23, 0)
				(_nw97).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _409_v2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _409_v2)), 1)
				(_nw97).ArraySet1((_410_v3).Update(_408_v1, _409_v2), 2)
				(_nw97).ArraySet1(_410_v3, 3)
				(_nw97).ArraySet1(_410_v3, 4)
				(_nw97).ArraySet1(_410_v3, 5)
				(_nw97).ArraySet1(_410_v3, 6)
				(_nw97).ArraySet1(_410_v3, 7)
				(_nw97).ArraySet1((_410_v3).Merge(_410_v3), 8)
				(_nw97).ArraySet1(_410_v3, 9)
				(_nw97).ArraySet1(_410_v3, 10)
				(_nw97).ArraySet1(_410_v3, 11)
				(_nw97).ArraySet1(_410_v3, 12)
				(_nw97).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_411_v4).Select((Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_411_v4).Cardinality()))).Uint32()).(*C0), (_dafny.Zero).Minus(_dafny.IntOfUint32((_428_v16).Cardinality()))), 13)
				(_nw97).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _541_cf42)).Merge((_410_v3).Update(_408_v1, _541_cf42)), 14)
				(_nw97).ArraySet1(_410_v3, 15)
				(_nw97).ArraySet1((_410_v3).Update(_408_v1, _542_cf41), 16)
				(_nw97).ArraySet1(_410_v3, 17)
				(_nw97).ArraySet1(_410_v3, 18)
				(_nw97).ArraySet1(_410_v3, 19)
				(_nw97).ArraySet1(((_410_v3).Update(_408_v1, Companion_Default___.Fm1(_409_v2, globalState))).Merge(_410_v3), 20)
				(_nw97).ArraySet1(_410_v3, 21)
				(_nw97).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, (p1).Cardinality()), 22)
				(_nw97).ArraySet1(_410_v3, 23)
				(_nw97).ArraySet1(_410_v3, 24)
				(_nw97).ArraySet1((_410_v3).Merge(_410_v3), 25)
				_549_v93 = _nw97
				_549_v93 = _549_v93
				var _550_v94 *C0
				_ = _550_v94
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__((func() _dafny.Array {
					if false {
						return (_408_v1).F6()
					}
					return (_408_v1).F6()
				})())
				_550_v94 = _nw98
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_544_v90), 0))
				_ = _index81
				(_544_v90).ArraySet1((_dafny.IntOfInt64(-10)).Times(_542_cf41), (_index81).Int())
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_544_v90), 0))
				_ = _index82
				(_544_v90).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_522_v84).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-978), _dafny.IntOfUint32((_522_v84).Cardinality()))).Uint32()).(bool) {
						return _409_v2
					}
					return _409_v2
				})()), (_index82).Int())
			} else if _source5.Is_DC19() {
				var _551___mcc_h12 _dafny.Sequence = _source5.Get_().(D6_DC19).Cf34
				_ = _551___mcc_h12
				var _552_cf34 _dafny.Sequence = _551___mcc_h12
				_ = _552_cf34
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
				_ = _index83
				(_422_v12).ArraySet1(p0, (_index83).Int())
				_494_v60 = _494_v60
				var _553_v95 _dafny.Set
				_ = _553_v95
				_553_v95 = _dafny.SetOf(p1)
				(globalState).F1 = ((_409_v2).Minus(_409_v2)).Times((_dafny.Zero).Minus((_409_v2).Plus((_553_v95).Cardinality())))
				var _554_v96 D2
				_ = _554_v96
				_554_v96 = Companion_D2_.Create_DC8_((_dafny.Zero).Minus(_409_v2), _dafny.IntOfUint32((_522_v84).Cardinality()))
				var _555_v97 _dafny.Sequence
				_ = _555_v97
				_555_v97 = _dafny.SeqOf(Companion_Default___.Fm17(globalState), Companion_D2_.Create_DC8_(_409_v2, _dafny.IntOfUint32((p2).Cardinality())), _554_v96)
				var _556_v98 _dafny.Set
				_ = _556_v98
				_556_v98 = _dafny.SetOf(_408_v1)
				var _557_v99 _dafny.Sequence
				_ = _557_v99
				_557_v99 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_555_v97, (Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((_555_v97).Cardinality()))).Uint32(), _554_v96), _dafny.SeqOf(_554_v96, Companion_D2_.Create_DC8_((_dafny.Zero).Minus((_556_v98).Cardinality()), _409_v2), _554_v96), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_554_v96), _555_v97), _dafny.SeqOf(_554_v96, _554_v96, _554_v96))
				_557_v99 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(297))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_558_v97 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_559_i9 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_558_v97, _558_v97)
					}
				})(_555_v97)))
			} else {
				var _560___mcc_h13 D6 = _source5.Get_().(D6_DC23).Cf43
				_ = _560___mcc_h13
				var _561_cf43 D6 = _560___mcc_h13
				_ = _561_cf43
				(globalState).F3 = _409_v2
				var _562_v100 *C0
				_ = _562_v100
				var _nw99 *C0 = New_C0_()
				_ = _nw99
				_nw99.Ctor__((func() _dafny.Array {
					if false {
						return (_408_v1).F6()
					}
					return (_408_v1).F6()
				})())
				_562_v100 = _nw99
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_497_v61), 0))
				_ = _index84
				(_497_v61).ArraySet1(_522_v84, (_index84).Int())
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_497_v61), 0))
				_ = _index85
				(_497_v61).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_522_v84, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm11(p0, (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool), p3, globalState), (Companion_Default___.SafeIndex(_409_v2, _dafny.IntOfUint32((Companion_Default___.Fm11(p0, (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool), p3, globalState)).Cardinality()))).Uint32(), true)), (_index85).Int())
				(globalState).F1 = _409_v2
			}
			var _563_v101 _dafny.Int
			_ = _563_v101
			var _564_v102 _dafny.Array
			_ = _564_v102
			var _565_v103 D0
			_ = _565_v103
			var _out29 _dafny.Int
			_ = _out29
			var _out30 _dafny.Array
			_ = _out30
			var _out31 D0
			_ = _out31
			_out29, _out30, _out31 = (_this).M0(globalState)
			_563_v101 = _out29
			_564_v102 = _out30
			_565_v103 = _out31
		} else {
			var _566_v104 _dafny.Map
			_ = _566_v104
			_566_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _426_v14)
			_566_v104 = (_566_v104).Update((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool), _426_v14)
			_426_v14 = _426_v14
			var _567_v105 _dafny.MultiSet
			_ = _567_v105
			_567_v105 = _dafny.MultiSetOf(!((_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool)), (_422_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))).Int()).(bool))
			_567_v105 = _567_v105
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_422_v12), 0))
			_ = _index86
			(_422_v12).ArraySet1(Companion_Default___.Fm0(_409_v2, globalState), (_index86).Int())
			var _568_v106 _dafny.Array
			_ = _568_v106
			var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
			_ = _nw100
			_568_v106 = _nw100
			_568_v106 = _568_v106
		}
		var _569_v107 D3
		_ = _569_v107
		_569_v107 = Companion_D3_.Create_DC11_(_409_v2)
		r0 = (_409_v2).Plus((_569_v107).Dtor_cf14())
		return r0
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
