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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(4), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false)).Cardinality())).Plus(_dafny.IntOfInt64(592))))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) bool {
	return !(!(((func() _dafny.Set {
		if !(true) {
			return _dafny.SetOf(_dafny.CodePoint('t'), _dafny.CodePoint('s'), _dafny.CodePoint('f'))
		}
		return _dafny.SetOf(_dafny.CodePoint('r'), _dafny.CodePoint('k'))
	})()).IsProperSubsetOf(_dafny.SetOf(_dafny.CodePoint('k')))))
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.Set {
	if !(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("anivmhual"), _dafny.CodePoint('x'))) {
		return _dafny.SetOf(false)
	} else {
		return _dafny.SetOf(true, !(true), true)
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	if _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("ujvkkcrk"), _dafny.CodePoint('w')) {
		return _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(16))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rpe")).Cardinality())))).Cardinality()))).Cardinality())
	} else {
		return _dafny.SetOf(_dafny.IntOfInt64(233))
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(24))).Cardinality())).Cardinality(), true)).Cardinality())).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _1_v1 _dafny.CodePoint
				_1_v1 = interface{}(_compr_1).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(24))).Cardinality())).Cardinality(), true)).Cardinality())).Contains(_1_v1) {
					_coll1.Add(_1_v1)
				}
			}
			return _coll1.ToSet()
		}()).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.CodePoint
			_2_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if (func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(24))).Cardinality())).Cardinality(), true)).Cardinality())).Keys().Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _3_v1 _dafny.CodePoint
					_3_v1 = interface{}(_compr_2).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(24))).Cardinality())).Cardinality(), true)).Cardinality())).Contains(_3_v1) {
						_coll2.Add(_3_v1)
					}
				}
				return _coll2.ToSet()
			}()).Contains(_2_v0) {
				_coll0.Add(_2_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('v'))).Cardinality())
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))).Cardinality())))).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o'))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _5_v3 _dafny.CodePoint
				_5_v3 = interface{}(_compr_4).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o')), _5_v3) {
					_coll4.Add(_5_v3, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)), _dafny.MultiSetFromSeq(_dafny.SeqOf(true)), _dafny.MultiSetOf(false), _dafny.MultiSetOf(true))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(629), _dafny.IntOfInt64(834))).Cardinality()))).Cardinality()))
				}
			}
			return _coll4.ToMap()
		}()).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _6_v2 _dafny.CodePoint
			_6_v2 = interface{}(_compr_3).(_dafny.CodePoint)
			if (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o'))).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _5_v3 _dafny.CodePoint
					_5_v3 = interface{}(_compr_5).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('o')), _5_v3) {
						_coll5.Add(_5_v3, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)), _dafny.MultiSetFromSeq(_dafny.SeqOf(true)), _dafny.MultiSetOf(false), _dafny.MultiSetOf(true))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(629), _dafny.IntOfInt64(834))).Cardinality()))).Cardinality()))
					}
				}
				return _coll5.ToMap()
			}()).Contains(_6_v2) {
				_coll3.Add(_6_v2, (func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(309), _dafny.IntOfInt64(772))); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _7_v4 _dafny.Int
						_7_v4 = interface{}(_compr_6).(_dafny.Int)
						if ((_dafny.IntOfInt64(309)).Cmp(_7_v4) <= 0) && ((_7_v4).Cmp(_dafny.IntOfInt64(772)) < 0) {
							_coll6.Add(Companion_Default___.SafeDivisionInt(_7_v4, _dafny.IntOfInt64(-942)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wg")).Cardinality())))
						}
					}
					return _coll6.ToMap()
				}()).Cardinality())
			}
		}
		return _coll3.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("kh")
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqOf(false))).Union((_dafny.MultiSetOf(_dafny.SeqOf(true))).Union(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(722))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(false)
	})))))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('a')
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(35))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("nsov"), false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("k"), !(true))).Merge(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.UnicodeSeqOfUtf8Bytes("v"))).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _10_v0 _dafny.Sequence
			_10_v0 = interface{}(_compr_7).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.UnicodeSeqOfUtf8Bytes("v"))).Contains(_10_v0) {
				_coll7.Add(_10_v0, true)
			}
		}
		return _coll7.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), ((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pess")).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pj")).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(501), true)).Cardinality())).Cardinality()).Times(_dafny.IntOfInt64(143)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, p1 _dafny.Int, p2 D3, globalState *GlobalState) _dafny.MultiSet {
	if (Companion_D1_.Create_DC6_(!(false), false)).Dtor_cf8() {
		return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("u"))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ddchqe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(788))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})))))
	} else {
		return (_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(597))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_12_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))))
	}
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true, false, false, true, true)).Cardinality()), _dafny.IntOfInt64(698)))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.IntOfInt64(472)).Minus(_dafny.IntOfInt64(-750)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D12_.Create_DC31_(true, _dafny.UnicodeSeqOfUtf8Bytes("wyhk"), _dafny.IntOfInt64(716), true)).Dtor_cf47(), (Companion_D5_.Create_DC14_(!(false), _dafny.IntOfInt64(560))).Dtor_cf17())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("qgxmchpc"), _dafny.IntOfInt64(411)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-378), _dafny.IntOfInt64(-241))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _13_v0 _dafny.Int
			_13_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(-378)).Cmp(_13_v0) <= 0) && ((_13_v0).Cmp(_dafny.IntOfInt64(-241)) < 0) {
				_coll8.Add((_13_v0).Times(_dafny.IntOfInt64(-530)), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_14_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(690)
				}))).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality())))
			}
		}
		return _coll8.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D3 = Companion_D3_.Create_DC10_(!(false))
	_ = _source0
	if _source0.Is_DC9() {
		var _15___mcc_h0 bool = _source0.Get_().(D3_DC9).Cf11
		_ = _15___mcc_h0
		var _16___mcc_h1 _dafny.Int = _source0.Get_().(D3_DC9).Cf12
		_ = _16___mcc_h1
		var _17_cf12 _dafny.Int = _16___mcc_h1
		_ = _17_cf12
		var _18_cf11 bool = _15___mcc_h0
		_ = _18_cf11
		return _dafny.CodePoint('u')
	} else if _source0.Is_DC10() {
		var _19___mcc_h2 bool = _source0.Get_().(D3_DC10).Cf13
		_ = _19___mcc_h2
		var _20_cf13 bool = _19___mcc_h2
		_ = _20_cf13
		return _dafny.CodePoint('e')
	} else {
		var _21___mcc_h3 _dafny.CodePoint = _source0.Get_().(D3_DC8).Cf10
		_ = _21___mcc_h3
		var _22_cf10 _dafny.CodePoint = _21___mcc_h3
		_ = _22_cf10
		return _22_cf10
	}
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ihvn"), _dafny.UnicodeSeqOfUtf8Bytes("nvbkgws")), _dafny.UnicodeSeqOfUtf8Bytes("qjxpor"))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(282)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(266)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-107))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(571)))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-604)))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(453)))))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true)))
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_(!(true))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Map, p1 D9, p2 bool, p3 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(-304))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(597)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-128)))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(false, true)
		}
		return _dafny.SeqOf(false, true, true)
	})(), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-633), (func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-876), _dafny.IntOfInt64(398))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(-876)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(398)) < 0) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_23_v0, _dafny.IntOfInt64(378)))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(80))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_24_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()))))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(52)))).Union((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_25_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
			}))).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(168)))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.SetOf(_dafny.CodePoint('a'))).Cardinality()), false)).Cardinality(), _dafny.IntOfInt64(569))).Cardinality())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(976))).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(_dafny.IntOfInt64(295), Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(!(true))).Cardinality(), _dafny.IntOfInt64(420)), (_dafny.IntOfInt64(151)).Cmp(_dafny.IntOfInt64(-521)) <= 0)
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(-734))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) D1 {
	var _source1 D12 = Companion_D12_.Create_DC29_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(595))).Cardinality()), _dafny.IntOfInt64(463)))
	_ = _source1
	if _source1.Is_DC30() {
		var _26___mcc_h0 bool = _source1.Get_().(D12_DC30).Cf44
		_ = _26___mcc_h0
		var _27___mcc_h1 bool = _source1.Get_().(D12_DC30).Cf45
		_ = _27___mcc_h1
		var _28_cf45 bool = _27___mcc_h1
		_ = _28_cf45
		var _29_cf44 bool = _26___mcc_h0
		_ = _29_cf44
		return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(993))
	} else if _source1.Is_DC31() {
		var _30___mcc_h2 bool = _source1.Get_().(D12_DC31).Cf46
		_ = _30___mcc_h2
		var _31___mcc_h3 _dafny.Sequence = _source1.Get_().(D12_DC31).Cf47
		_ = _31___mcc_h3
		var _32___mcc_h4 _dafny.Int = _source1.Get_().(D12_DC31).Cf48
		_ = _32___mcc_h4
		var _33___mcc_h5 bool = _source1.Get_().(D12_DC31).Cf49
		_ = _33___mcc_h5
		var _34_cf49 bool = _33___mcc_h5
		_ = _34_cf49
		var _35_cf48 _dafny.Int = _32___mcc_h4
		_ = _35_cf48
		var _36_cf47 _dafny.Sequence = _31___mcc_h3
		_ = _36_cf47
		var _37_cf46 bool = _30___mcc_h2
		_ = _37_cf46
		return Companion_D1_.Create_DC5_((_dafny.Zero).Minus(_35_cf48))
	} else if _source1.Is_DC29() {
		var _38___mcc_h6 _dafny.Set = _source1.Get_().(D12_DC29).Cf43
		_ = _38___mcc_h6
		var _39_cf43 _dafny.Set = _38___mcc_h6
		_ = _39_cf43
		return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(215))
	} else {
		var _40___mcc_h7 D12 = _source1.Get_().(D12_DC32).Cf50
		_ = _40___mcc_h7
		var _41_cf50 D12 = _40___mcc_h7
		_ = _41_cf50
		if false {
			return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(-264))
		} else {
			return Companion_D1_.Create_DC5_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(_dafny.IntOfInt64(784))).Cardinality())).Cardinality())
		}
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool, _dafny.Int) {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	if !(p1) {
		var _42_v0 _dafny.CodePoint
		_ = _42_v0
		_42_v0 = _dafny.CodePoint('d')
		var _43_v1 _dafny.Sequence
		_ = _43_v1
		_43_v1 = _dafny.UnicodeSeqOfUtf8Bytes("srfvaskfl")
		var _44_v2 D0
		_ = _44_v2
		_44_v2 = Companion_D0_.Create_DC0_(p2)
		var _45_v3 _dafny.Map
		_ = _45_v3
		_45_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_43_v1, (_44_v2).Dtor_cf0())
		var _46_v4 *C3
		_ = _46_v4
		var _nw0 *C3 = New_C3_()
		_ = _nw0
		_nw0.Ctor__(_42_v0, _45_v3)
		_46_v4 = _nw0
		var _47_v5 _dafny.Set
		_ = _47_v5
		_47_v5 = _dafny.SetOf(_46_v4, _46_v4)
		var _48_v6 _dafny.Sequence
		_ = _48_v6
		_48_v6 = _dafny.SeqOf(_47_v5)
		var _49_v7 D1
		_ = _49_v7
		_49_v7 = Companion_D1_.Create_DC6_(p1, p1)
		var _50_v8 _dafny.Sequence
		_ = _50_v8
		_50_v8 = _dafny.SeqOf(p1)
		var _51_v9 _dafny.MultiSet
		_ = _51_v9
		_51_v9 = _dafny.MultiSetOf(p2, p0, (_dafny.Zero).Minus(p2), (_dafny.Zero).Minus(p0), p0)
		var _52_v10 _dafny.Array
		_ = _52_v10
		var _nwElement0_0 bool = p1
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(17))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1(p1, 1)
		(_nw1).ArraySet1(p1, 2)
		(_nw1).ArraySet1(((_48_v6).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_48_v6).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf(_47_v5), 3)
		(_nw1).ArraySet1((_49_v7).Dtor_cf8(), 4)
		(_nw1).ArraySet1(p1, 5)
		(_nw1).ArraySet1(!(((_dafny.Zero).Minus(p0)).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm21(globalState)).Cardinality())) < 0), 6)
		(_nw1).ArraySet1(p1, 7)
		(_nw1).ArraySet1(!(false), 8)
		(_nw1).ArraySet1((_50_v8).Select((Companion_Default___.SafeIndex((_51_v9).Cardinality(), _dafny.IntOfUint32((_50_v8).Cardinality()))).Uint32()).(bool), 9)
		(_nw1).ArraySet1(p1, 10)
		(_nw1).ArraySet1(p1, 11)
		(_nw1).ArraySet1(p1, 12)
		(_nw1).ArraySet1(p1, 13)
		(_nw1).ArraySet1((p0).Cmp(p0) <= 0, 14)
		(_nw1).ArraySet1(p1, 15)
		(_nw1).ArraySet1(p1, 16)
		_52_v10 = _nw1
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_52_v10), 0))
		_ = _index0
		(_52_v10).ArraySet1(!(p1), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_52_v10), 0))
		_ = _index1
		(_52_v10).ArraySet1(p1, (_index1).Int())
		var _53_v11 _dafny.Map
		_ = _53_v11
		_53_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_52_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_52_v10), 0))).Int()).(bool), _42_v0)
		_53_v11 = (_53_v11).Update(p1, _42_v0)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_52_v10), 0))
		_ = _index2
		(_52_v10).ArraySet1(true, (_index2).Int())
		(globalState).F6 = _dafny.IntOfUint32((_43_v1).Cardinality())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_52_v10), 0))
		_ = _index3
		(_52_v10).ArraySet1(!((_52_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_52_v10), 0))).Int()).(bool)) || (p1), (_index3).Int())
	} else {
		var _54_v12 _dafny.CodePoint
		_ = _54_v12
		_54_v12 = _dafny.CodePoint('j')
		var _55_v13 _dafny.Sequence
		_ = _55_v13
		_55_v13 = _dafny.UnicodeSeqOfUtf8Bytes("vqatmsp")
		var _56_v14 _dafny.Array
		_ = _56_v14
		var _nwElement0_1 _dafny.Int = p0
		_ = _nwElement0_1
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(6))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_1, 0)
		(_nw2).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_57_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality()))).Uint32(), _54_v12)).Cardinality()), 1)
		(_nw2).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_55_v13).Cardinality()))), 2)
		(_nw2).ArraySet1(p0, 3)
		(_nw2).ArraySet1(_dafny.IntOfInt64(-689), 4)
		(_nw2).ArraySet1(p2, 5)
		_56_v14 = _nw2
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))
		_ = _index4
		(_56_v14).ArraySet1(p0, (_index4).Int())
		var _59_v15 _dafny.MultiSet
		_ = _59_v15
		_59_v15 = _dafny.MultiSetOf(p0)
		var _60_v16 _dafny.Array
		_ = _60_v16
		var _nwElement0_2 _dafny.MultiSet = _59_v15
		_ = _nwElement0_2
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_2, 0)
		(_nw3).ArraySet1((_59_v15).Difference(_59_v15), 1)
		(_nw3).ArraySet1(_59_v15, 2)
		(_nw3).ArraySet1((_59_v15).Intersection(_59_v15), 3)
		_60_v16 = _nw3
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_60_v16), 0))
		_ = _index5
		(_60_v16).ArraySet1(Companion_Default___.Fm28(p1, p2, p1, globalState), (_index5).Int())
		var _61_v17 _dafny.MultiSet
		_ = _61_v17
		_61_v17 = _dafny.MultiSetOf(p1, p1, p1, !(!(p1)))
		var _62_v18 _dafny.Array
		_ = _62_v18
		var _nwElement0_3 bool = p1
		_ = _nwElement0_3
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_3, 0)
		(_nw4).ArraySet1(p1, 1)
		(_nw4).ArraySet1(p1, 2)
		(_nw4).ArraySet1(true, 3)
		(_nw4).ArraySet1(p1, 4)
		(_nw4).ArraySet1(p1, 5)
		(_nw4).ArraySet1(p1, 6)
		(_nw4).ArraySet1(p1, 7)
		(_nw4).ArraySet1(p1, 8)
		(_nw4).ArraySet1((func() bool {
			if p1 {
				return p1
			}
			return p1
		})(), 9)
		(_nw4).ArraySet1(p1, 10)
		(_nw4).ArraySet1(true, 11)
		(_nw4).ArraySet1(Companion_Default___.Fm1(p1, _54_v12, _dafny.IntOfInt64(245), globalState), 12)
		(_nw4).ArraySet1(!(!(p1)), 13)
		(_nw4).ArraySet1((_61_v17).IsDisjointFrom(_61_v17), 14)
		(_nw4).ArraySet1(Companion_Default___.Fm1(true, _54_v12, p2, globalState), 15)
		(_nw4).ArraySet1(p1, 16)
		_62_v18 = _nw4
		var _63_v19 _dafny.Map
		_ = _63_v19
		_63_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("fjijcedg"), p2)
		var _64_v20 *C2
		_ = _64_v20
		var _nw5 *C2 = New_C2_()
		_ = _nw5
		_nw5.Ctor__(_54_v12, _63_v19)
		_64_v20 = _nw5
		var _65_v21 _dafny.Set
		_ = _65_v21
		_65_v21 = _dafny.SetOf(_64_v20)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))
		_ = _index6
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_60_v16), 0))
		_ = _index7
		var _rhs0 _dafny.Int = Companion_Default___.SafeDivisionInt(p2, p2)
		_ = _rhs0
		var _rhs1 _dafny.MultiSet = (_59_v15).Difference(_59_v15)
		_ = _rhs1
		var _rhs2 _dafny.Array = _62_v18
		_ = _rhs2
		var _rhs3 _dafny.Int = _dafny.IntOfInt64(664)
		_ = _rhs3
		var _rhs4 bool = !(!(_65_v21).Contains(_64_v20))
		_ = _rhs4
		var _lhs0 _dafny.Array = _56_v14
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))
		_ = _lhs1
		var _lhs2 _dafny.Array = _60_v16
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_60_v16), 0))
		_ = _lhs3
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		(_lhs2).ArraySet1(_rhs1, (_lhs3).Int())
		r0 = _rhs2
		_lhs4.F6 = _rhs3
		r1 = _rhs4
		var _66_v22 *C3
		_ = _66_v22
		var _nw6 *C3 = New_C3_()
		_ = _nw6
		_nw6.Ctor__(_54_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("umuxo"), (_56_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))).Int()).(_dafny.Int)))
		_66_v22 = _nw6
		var _67_v23 _dafny.Array
		_ = _67_v23
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_0
		var _nw7 _dafny.Array
		_ = _nw7
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw7 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = (func(_68_v13 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_69_i1 _dafny.Int) _dafny.Sequence {
					return _68_v13
				}
			})(_55_v13)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw7 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw7).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw7).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_67_v23 = _nw7
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_67_v23), 0))
		_ = _index8
		(_67_v23).ArraySet1(_55_v13, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_67_v23), 0))
		_ = _index9
		var _rhs5 bool = p1
		_ = _rhs5
		var _rhs6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_70_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_71_i2 _dafny.Int) _dafny.CodePoint {
				return _70_v12
			}
		})(_54_v12))), _55_v13), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_55_v13, (Companion_Default___.SafeIndex((_56_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_55_v13).Cardinality()))).Uint32(), _54_v12), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_55_v13, (Companion_Default___.SafeIndex((_56_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_55_v13).Cardinality()))).Uint32(), _54_v12)).Cardinality()))).Uint32(), _dafny.CodePoint('q')))
		_ = _rhs6
		var _rhs7 *C3 = _66_v22
		_ = _rhs7
		var _rhs8 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("estdt"), _dafny.UnicodeSeqOfUtf8Bytes("mfn"))
		_ = _rhs8
		var _lhs5 _dafny.Array = _67_v23
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_67_v23), 0))
		_ = _lhs6
		r1 = _rhs5
		_55_v13 = _rhs6
		_66_v22 = _rhs7
		(_lhs5).ArraySet1(_rhs8, (_lhs6).Int())
		var _72_v24 D0
		_ = _72_v24
		_72_v24 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(p2))
		var _73_v25 D11
		_ = _73_v25
		_73_v25 = Companion_D11_.Create_DC27_(_63_v19)
		var _74_v26 *C1
		_ = _74_v26
		var _nw8 *C1 = New_C1_()
		_ = _nw8
		_nw8.Ctor__(_dafny.IntOfInt64(687), p1, Companion_Default___.Fm20(Companion_Default___.Fm1(p1, _54_v12, (_59_v15).Cardinality(), globalState), globalState), ((_73_v25).Dtor_cf39()).Merge((_63_v19).Update((_67_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_67_v23), 0))).Int()).(_dafny.Sequence), (_56_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))).Int()).(_dafny.Int))))
		_74_v26 = _nw8
		var _75_v27 _dafny.Map
		_ = _75_v27
		_75_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v12, p0)
		var _76_v28 T0
		_ = _76_v28
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__((_67_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_67_v23), 0))).Int()).(_dafny.Sequence), _75_v27, _54_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v13, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), (_56_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))).Int()).(_dafny.Int))).Cardinality()))
		_76_v28 = _nw9
		var _77_v29 _dafny.Map
		_ = _77_v29
		_77_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v28, p2)
		var _rhs9 _dafny.Array = _62_v18
		_ = _rhs9
		var _rhs10 D0 = _72_v24
		_ = _rhs10
		var _rhs11 *C1 = _74_v26
		_ = _rhs11
		var _rhs12 _dafny.Int = (func() _dafny.Int {
			if (_77_v29).Contains(_76_v28) {
				return (_77_v29).Get(_76_v28).(_dafny.Int)
			}
			return ((_56_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_56_v14), 0))).Int()).(_dafny.Int)).Minus(p2)
		})()
		_ = _rhs12
		var _rhs13 _dafny.Array = _62_v18
		_ = _rhs13
		_62_v18 = _rhs9
		_72_v24 = _rhs10
		_74_v26 = _rhs11
		r2 = _rhs12
		_62_v18 = _rhs13
		(_76_v28).F18_set_(_54_v12)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_62_v18), 0))
		_ = _index10
		(_62_v18).ArraySet1((_74_v26).F23(), (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_62_v18), 0))
		_ = _index11
		(_62_v18).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_74_v26.F22), _dafny.IntOfInt64(316)), (_index11).Int())
	}
	var _78_v30 _dafny.Sequence
	_ = _78_v30
	_78_v30 = _dafny.UnicodeSeqOfUtf8Bytes("lgcj")
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_78_v30, Companion_Default___.Fm21(globalState))).Cardinality())
	_ = _hi0
	for _79_i3 := p0; _79_i3.Cmp(_hi0) < 0; _79_i3 = _79_i3.Plus(_dafny.One) {
		var _80_v31 _dafny.Map
		_ = _80_v31
		_80_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _79_i3)
		(globalState).F6 = (func() _dafny.Int {
			if (_80_v31).Contains(false) {
				return (_80_v31).Get(false).(_dafny.Int)
			}
			return _79_i3
		})()
		var _81_v32 _dafny.Map
		_ = _81_v32
		_81_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), _dafny.Companion_Sequence_.Concatenate(_78_v30, _78_v30))
		var _82_v33 _dafny.CodePoint
		_ = _82_v33
		_82_v33 = _dafny.CodePoint('y')
		_81_v32 = (_81_v32).Update(Companion_Default___.Fm1(p1, _82_v33, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(538))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_83_v33 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_84_i4 _dafny.Int) _dafny.CodePoint {
				return _83_v33
			}
		})(_82_v33))), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(538))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_85_v33 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_86_i4 _dafny.Int) _dafny.CodePoint {
				return _85_v33
			}
		})(_82_v33)))).Cardinality()))).Uint32(), Companion_Default___.Fm20(Companion_Default___.Fm1(p1, _82_v33, p2, globalState), globalState))).Cardinality()), globalState), _78_v30)
		var _87_v34 _dafny.Map
		_ = _87_v34
		_87_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v33, p2)
		var _88_v35 _dafny.Map
		_ = _88_v35
		_88_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, p0)
		var _89_v36 *C0
		_ = _89_v36
		var _nw10 *C0 = New_C0_()
		_ = _nw10
		_nw10.Ctor__(_dafny.Companion_Sequence_.Update(_78_v30, (Companion_Default___.SafeIndex(_79_i3, _dafny.IntOfUint32((_78_v30).Cardinality()))).Uint32(), _82_v33), (_87_v34).Merge(_87_v34), _dafny.CodePoint('r'), (_88_v35).Merge(_88_v35))
		_89_v36 = _nw10
		var _90_v37 _dafny.Array
		_ = _90_v37
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(3))
		_ = _nw11
		_90_v37 = _nw11
		var _91_v38 _dafny.Map
		_ = _91_v38
		_91_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _90_v37)
		var _92_v39 _dafny.Array
		_ = _92_v39
		var _nwElement0_4 _dafny.Array = _90_v37
		_ = _nwElement0_4
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(24))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_4, 0)
		(_nw12).ArraySet1(_90_v37, 1)
		(_nw12).ArraySet1(_90_v37, 2)
		(_nw12).ArraySet1(_90_v37, 3)
		(_nw12).ArraySet1(_90_v37, 4)
		(_nw12).ArraySet1(_90_v37, 5)
		(_nw12).ArraySet1(_90_v37, 6)
		(_nw12).ArraySet1(_90_v37, 7)
		(_nw12).ArraySet1(_90_v37, 8)
		(_nw12).ArraySet1(_90_v37, 9)
		(_nw12).ArraySet1((func() _dafny.Array {
			if (_91_v38).Contains(p0) {
				return (_91_v38).Get(p0).(_dafny.Array)
			}
			return _90_v37
		})(), 10)
		(_nw12).ArraySet1(_90_v37, 11)
		(_nw12).ArraySet1(_90_v37, 12)
		(_nw12).ArraySet1(_90_v37, 13)
		(_nw12).ArraySet1(_90_v37, 14)
		(_nw12).ArraySet1(_90_v37, 15)
		(_nw12).ArraySet1(_90_v37, 16)
		(_nw12).ArraySet1(_90_v37, 17)
		(_nw12).ArraySet1(_90_v37, 18)
		(_nw12).ArraySet1(_90_v37, 19)
		(_nw12).ArraySet1(_90_v37, 20)
		(_nw12).ArraySet1(_90_v37, 21)
		(_nw12).ArraySet1(_90_v37, 22)
		(_nw12).ArraySet1(_90_v37, 23)
		_92_v39 = _nw12
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_92_v39), 0))
		_ = _index12
		(_92_v39).ArraySet1(_90_v37, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_92_v39), 0))
		_ = _index13
		(_92_v39).ArraySet1(_90_v37, (_index13).Int())
	}
	var _93_v40 _dafny.Array
	_ = _93_v40
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
	_ = _nw13
	_93_v40 = _nw13
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
	_ = _index14
	(_93_v40).ArraySet1(Companion_Default___.SafeDivisionInt(p2, p0), (_index14).Int())
	var _94_v41 _dafny.Array
	_ = _94_v41
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
	_ = _nw14
	_94_v41 = _nw14
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
	_ = _index15
	(_94_v41).ArraySet1(((_dafny.Zero).Minus((_dafny.Zero).Minus(p2))).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))) >= 0, (_index15).Int())
	var _95_v42 _dafny.CodePoint
	_ = _95_v42
	_95_v42 = _dafny.CodePoint('j')
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
	_ = _index16
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
	_ = _index17
	var _rhs14 _dafny.Int = _dafny.IntOfInt64(715)
	_ = _rhs14
	var _rhs15 bool = !(p1)
	_ = _rhs15
	var _rhs16 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_95_v42, Companion_Default___.Fm10(globalState)), Companion_Default___.Fm8(globalState))
	_ = _rhs16
	var _lhs7 _dafny.Array = _93_v40
	_ = _lhs7
	var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
	_ = _lhs8
	var _lhs9 _dafny.Array = _94_v41
	_ = _lhs9
	var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
	_ = _lhs10
	(_lhs7).ArraySet1(_rhs14, (_lhs8).Int())
	(_lhs9).ArraySet1(_rhs15, (_lhs10).Int())
	_78_v30 = _rhs16
	var _96_v43 _dafny.Map
	_ = _96_v43
	_96_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v42, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
	var _97_v45 _dafny.Map
	_ = _97_v45
	_97_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v42, _dafny.IntOfInt64(979))
	_96_v43 = (func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_97_v45).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _98_v44 _dafny.CodePoint
			_98_v44 = interface{}(_compr_10).(_dafny.CodePoint)
			if (_97_v45).Contains(_98_v44) {
				_coll10.Add(_98_v44, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
			}
		}
		return _coll10.ToMap()
	}()).Merge((_96_v43).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v42, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))))
	var _99_v46 _dafny.Map
	_ = _99_v46
	_99_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p1)
	if ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), (_99_v46).Cardinality()))).Cmp(p2) < 0 {
		var _100_v47 _dafny.Map
		_ = _100_v47
		_100_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
		var _101_v48 _dafny.Sequence
		_ = _101_v48
		_101_v48 = _dafny.SeqOf(_100_v47)
		if !_dafny.Companion_Sequence_.Equal(_101_v48, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_102_p2 _dafny.Int, _103_v41 _dafny.Array) func(_dafny.Int) _dafny.Map {
			return func(_104_i5 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll11 = _dafny.NewMapBuilder()
					_ = _coll11
					for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-992), _dafny.IntOfInt64(954))); ; {
						_compr_11, _ok11 := _iter11()
						if !_ok11 {
							break
						}
						var _105_v49 _dafny.Int
						_105_v49 = interface{}(_compr_11).(_dafny.Int)
						if ((_dafny.IntOfInt64(-992)).Cmp(_105_v49) <= 0) && ((_105_v49).Cmp(_dafny.IntOfInt64(954)) < 0) {
							_coll11.Add(Companion_Default___.SafeDivisionInt(_105_v49, _102_p2), (_103_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_103_v41), 0))).Int()).(bool))
						}
					}
					return _coll11.ToMap()
				}()
			}
		})(p2, _94_v41)))) {
			var _106_v50 D9
			_ = _106_v50
			_106_v50 = Companion_D9_.Create_DC23_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(865))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_107_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_108_i6 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_107_p0)
				}
			})(p0))), p2, p1, p2)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index19
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index20
			var _rhs17 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_dafny.SeqOf((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))).Cardinality()))).Uint32(), (_106_v50).Dtor_cf31())
			_ = _rhs17
			var _rhs18 bool = (func() bool {
				if p1 {
					return (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)
				}
				return !(p1)
			})()
			_ = _rhs18
			var _rhs19 _dafny.Int = Companion_Default___.SafeModuloInt((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), p0)
			_ = _rhs19
			var _rhs20 bool = (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)
			_ = _rhs20
			var _rhs21 bool = (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)
			_ = _rhs21
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 _dafny.Array = _94_v41
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _93_v40
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _lhs15
			var _lhs16 _dafny.Array = _94_v41
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _lhs17
			_lhs11.F1 = _rhs17
			(_lhs12).ArraySet1(_rhs18, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs19, (_lhs15).Int())
			(_lhs16).ArraySet1(_rhs20, (_lhs17).Int())
			r1 = _rhs21
			var _109_v51 _dafny.Map
			_ = _109_v51
			_109_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, p0)
			var _110_v52 _dafny.Set
			_ = _110_v52
			_110_v52 = _dafny.SetOf(true, true, !(p1))
			var _111_v53 D0
			_ = _111_v53
			_111_v53 = Companion_D0_.Create_DC0_(p0)
			_109_v51 = (_109_v51).Update(_dafny.Companion_Sequence_.Update(_78_v30, (Companion_Default___.SafeIndex((Companion_Default___.Fm31(_110_v52, p0, globalState)).Dtor_cf6(), _dafny.IntOfUint32((_78_v30).Cardinality()))).Uint32(), _95_v42), (_111_v53).Dtor_cf0())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index21
			(_94_v41).ArraySet1(p1, (_index21).Int())
			var _112_v54 _dafny.Map
			_ = _112_v54
			_112_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), p2)
			_112_v54 = (_112_v54).Update((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p2), _dafny.IntOfInt64(119))))
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index22
			(_93_v40).ArraySet1(_dafny.IntOfInt64(-852), (_index22).Int())
		} else {
			var _113_v55 D10
			_ = _113_v55
			_113_v55 = Companion_D10_.Create_DC25_(_93_v40)
			var _114_v56 _dafny.Sequence
			_ = _114_v56
			_114_v56 = _dafny.SeqOf((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
			var _pat_let_tv0 = _93_v40
			_ = _pat_let_tv0
			var _115_v57 _dafny.Map
			_ = _115_v57
			_115_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let0_0 D10) D10 {
				return func(_116_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let1_0 _dafny.Array) D10 {
						return func(_117_dt__update_hcf33_h0 _dafny.Array) D10 {
							return Companion_D10_.Create_DC25_(_117_dt__update_hcf33_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_113_v55)).Dtor_cf33(), _114_v56)
			_115_v57 = _115_v57
			var _118_v58 _dafny.Sequence
			_ = _118_v58
			_118_v58 = _dafny.SeqOf(p0, p0)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index23
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index24
			var _rhs22 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_119_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_120_i7 _dafny.Int) _dafny.CodePoint {
					return _119_v42
				}
			})(_95_v42)))).Cardinality())
			_ = _rhs22
			var _rhs23 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(p2)), p2)
			_ = _rhs23
			var _rhs24 _dafny.Sequence = _118_v58
			_ = _rhs24
			var _rhs25 bool = false
			_ = _rhs25
			var _lhs18 *GlobalState = globalState
			_ = _lhs18
			var _lhs19 _dafny.Array = _93_v40
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _lhs20
			var _lhs21 *GlobalState = globalState
			_ = _lhs21
			var _lhs22 _dafny.Array = _94_v41
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _lhs23
			_lhs18.F6 = _rhs22
			(_lhs19).ArraySet1(_rhs23, (_lhs20).Int())
			_lhs21.F1 = _rhs24
			(_lhs22).ArraySet1(_rhs25, (_lhs23).Int())
			r1 = (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)
			_95_v42 = _95_v42
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index25
			(_94_v41).ArraySet1(!(p1), (_index25).Int())
		}
		if true {
			var _121_v59 _dafny.MultiSet
			_ = _121_v59
			_121_v59 = _dafny.MultiSetOf(p1)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index26
			(_94_v41).ArraySet1((func() bool {
				if (_121_v59).IsProperSubsetOf(_121_v59) {
					return p1
				}
				return (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)
			})(), (_index26).Int())
			var _122_v60 _dafny.Sequence
			_ = _122_v60
			_122_v60 = _dafny.SeqOf((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
			var _123_v61 _dafny.Map
			_ = _123_v61
			_123_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, _dafny.IntOfUint32((_122_v60).Cardinality()))
			var _124_v62 *C1
			_ = _124_v62
			var _nw15 *C1 = New_C1_()
			_ = _nw15
			_nw15.Ctor__(p2, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), _95_v42, (_123_v61).Update(_78_v30, p0))
			_124_v62 = _nw15
			var _125_v63 _dafny.MultiSet
			_ = _125_v63
			_125_v63 = _dafny.MultiSetOf(_124_v62, _124_v62, _124_v62)
			var _126_v64 _dafny.Map
			_ = _126_v64
			_126_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), _125_v63)
			var _127_v65 _dafny.MultiSet
			_ = _127_v65
			_127_v65 = _dafny.MultiSetOf((_126_v64).Cardinality())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index27
			(_94_v41).ArraySet1(!((_127_v65).Contains((_dafny.Zero).Minus((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int)))), (_index27).Int())
			var _128_v66 _dafny.Map
			_ = _128_v66
			_128_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _93_v40)
			var _129_v67 _dafny.Set
			_ = _129_v67
			_129_v67 = _dafny.SetOf(p0)
			var _130_v68 _dafny.MultiSet
			_ = _130_v68
			_130_v68 = _dafny.MultiSetOf(_93_v40, _93_v40, (func() _dafny.Array {
				if (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool) {
					return (func() _dafny.Array {
						if (_128_v66).Contains((_129_v67).Cardinality()) {
							return (_128_v66).Get((_129_v67).Cardinality()).(_dafny.Array)
						}
						return _93_v40
					})()
				}
				return _93_v40
			})())
			_130_v68 = _130_v68
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index28
			(_93_v40).ArraySet1(Companion_Default___.SafeModuloInt((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), (p0).Times((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))), (_index28).Int())
			var _131_v69 *C0
			_ = _131_v69
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__(_78_v30, Companion_Default___.Fm14((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), globalState), _95_v42, _123_v61)
			_131_v69 = _nw16
		} else {
			var _132_v70 _dafny.Set
			_ = _132_v70
			_132_v70 = _dafny.SetOf(p2)
			var _133_v71 _dafny.Sequence
			_ = _133_v71
			_133_v71 = _dafny.SeqOf((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
			var _134_v72 _dafny.MultiSet
			_ = _134_v72
			_134_v72 = _dafny.MultiSetOf(_78_v30)
			var _135_v73 _dafny.Sequence
			_ = _135_v73
			_135_v73 = _dafny.SeqOf((_133_v71).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_78_v30, (Companion_Default___.SafeIndex((_134_v72).Cardinality(), _dafny.IntOfUint32((_78_v30).Cardinality()))).Uint32(), _95_v42)).Cardinality()), (func() bool {
				if (_100_v47).Contains(p2) {
					return (_100_v47).Get(p2).(bool)
				}
				return p1
			})())).Cardinality(), _dafny.IntOfUint32((_133_v71).Cardinality()))).Uint32()).(bool))
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index29
			var _rhs26 bool = ((_132_v70).Intersection(_132_v70)).IsDisjointFrom((_132_v70).Union(_132_v70))
			_ = _rhs26
			var _rhs27 bool = (_135_v73).Select((Companion_Default___.SafeIndex((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_135_v73).Cardinality()))).Uint32()).(bool)
			_ = _rhs27
			var _lhs24 _dafny.Array = _94_v41
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _lhs25
			r1 = _rhs26
			(_lhs24).ArraySet1(_rhs27, (_lhs25).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index30
			(_93_v40).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_78_v30).Cardinality()), p0), (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index31
			(_93_v40).ArraySet1(p2, (_index31).Int())
			_95_v42 = _95_v42
			var _136_v74 D1
			_ = _136_v74
			_136_v74 = Companion_D1_.Create_DC3_(p1)
			var _137_v75 D0
			_ = _137_v75
			_137_v75 = Companion_D0_.Create_DC1_()
			var _138_v76 _dafny.Array
			_ = _138_v76
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_1
			var _nw17 _dafny.Array
			_ = _nw17
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw17 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = (func(_139_v71 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_140_i8 _dafny.Int) _dafny.Sequence {
						return _139_v71
					}
				})(_133_v71)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw17 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw17).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw17).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_138_v76 = _nw17
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_138_v76), 0))
			_ = _index32
			(_138_v76).ArraySet1(_133_v71, (_index32).Int())
			var _141_v77 _dafny.Array
			_ = _141_v77
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw18
			_141_v77 = _nw18
			var _142_v78 _dafny.Array
			_ = _142_v78
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(17))
			_ = _nw19
			_142_v78 = _nw19
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_141_v77), 0))
			_ = _index33
			(_141_v77).ArraySet1(_142_v78, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_138_v76), 0))
			_ = _index34
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_141_v77), 0))
			_ = _index35
			var _rhs28 D1 = _136_v74
			_ = _rhs28
			var _rhs29 _dafny.Int = _dafny.IntOfInt64(-121)
			_ = _rhs29
			var _rhs30 D0 = (func() D0 {
				if p1 {
					return _137_v75
				}
				return _137_v75
			})()
			_ = _rhs30
			var _rhs31 _dafny.Sequence = _133_v71
			_ = _rhs31
			var _rhs32 _dafny.Array = _142_v78
			_ = _rhs32
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _138_v76
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_138_v76), 0))
			_ = _lhs28
			var _lhs29 _dafny.Array = _141_v77
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_141_v77), 0))
			_ = _lhs30
			_136_v74 = _rhs28
			_lhs26.F6 = _rhs29
			_137_v75 = _rhs30
			(_lhs27).ArraySet1(_rhs31, (_lhs28).Int())
			(_lhs29).ArraySet1(_rhs32, (_lhs30).Int())
		}
		var _143_v79 D6
		_ = _143_v79
		_143_v79 = Companion_D6_.Create_DC15_(_97_v45)
		var _source2 D6 = _143_v79
		_ = _source2
		if _source2.Is_DC16() {
			var _144___mcc_h0 bool = _source2.Get_().(D6_DC16).Cf19
			_ = _144___mcc_h0
			var _145___mcc_h1 _dafny.Array = _source2.Get_().(D6_DC16).Cf20
			_ = _145___mcc_h1
			var _146_cf20 _dafny.Array = _145___mcc_h1
			_ = _146_cf20
			var _147_cf19 bool = _144___mcc_h0
			_ = _147_cf19
			var _148_v80 _dafny.Map
			_ = _148_v80
			_148_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(613), _78_v30)
			_148_v80 = ((_148_v80).Merge(_148_v80)).Merge((_148_v80).Merge(_148_v80))
			var _149_v82 D9
			_ = _149_v82
			_149_v82 = Companion_D9_.Create_DC22_(_78_v30)
			var _150_v83 _dafny.Sequence
			_ = _150_v83
			_150_v83 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qribx"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qribx")).Cardinality()))).Uint32(), _95_v42), (_149_v82).Dtor_cf27())
			var _151_v84 *C2
			_ = _151_v84
			var _nw20 *C2 = New_C2_()
			_ = _nw20
			_nw20.Ctor__(_95_v42, func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_150_v83).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _152_v81 _dafny.Sequence
					_152_v81 = interface{}(_compr_12).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_150_v83, _152_v81) {
						_coll12.Add(_152_v81, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p2, p0)).Cardinality(), p1)).Cardinality())
					}
				}
				return _coll12.ToMap()
			}())
			_151_v84 = _nw20
			var _153_v85 _dafny.Set
			_ = _153_v85
			_153_v85 = _dafny.SetOf(_100_v47)
			var _154_v87 _dafny.MultiSet
			_ = _154_v87
			_154_v87 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(763), _dafny.IntOfInt64(-498))); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _155_v86 _dafny.Int
					_155_v86 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(763)).Cmp(_155_v86) <= 0) && ((_155_v86).Cmp(_dafny.IntOfInt64(-498)) < 0) {
						_coll13.Add((_155_v86).Times((_100_v47).Cardinality()), (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
					}
				}
				return _coll13.ToMap()
			}())
			_153_v85 = ((_153_v85).Difference(func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_154_v87).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _156_v88 _dafny.Map
					_156_v88 = interface{}(_compr_14).(_dafny.Map)
					if (_154_v87).Contains(_156_v88) {
						_coll14.Add(_156_v88)
					}
				}
				return _coll14.ToSet()
			}())).Difference(_153_v85)
			var _157_v89 _dafny.Map
			_ = _157_v89
			_157_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, p2)
			var _158_v90 *C1
			_ = _158_v90
			var _nw21 *C1 = New_C1_()
			_ = _nw21
			_nw21.Ctor__(p2, _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm30(p1, globalState), _dafny.SeqOf(p2, (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))), _dafny.CodePoint('u'), _157_v89)
			_158_v90 = _nw21
			var _159_v91 *C3
			_ = _159_v91
			var _nw22 *C3 = New_C3_()
			_ = _nw22
			_nw22.Ctor__(_95_v42, (_157_v89).Merge(_157_v89))
			_159_v91 = _nw22
			var _160_v92 _dafny.Array
			_ = _160_v92
			var _nw23 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw23
			_160_v92 = _nw23
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_160_v92), 0))
			_ = _index36
			(_160_v92).ArraySet1(_159_v91, (_index36).Int())
			var _161_v93 _dafny.MultiSet
			_ = _161_v93
			_161_v93 = _dafny.MultiSetOf((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), true)
			var _pat_let_tv1 = _93_v40
			_ = _pat_let_tv1
			var _pat_let_tv2 = _93_v40
			_ = _pat_let_tv2
			var _pat_let_tv3 = p1
			_ = _pat_let_tv3
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_160_v92), 0))
			_ = _index38
			var _rhs33 *C1 = _158_v90
			_ = _rhs33
			var _rhs34 *C3 = _159_v91
			_ = _rhs34
			var _rhs35 _dafny.Int = ((_dafny.IntOfUint32((_78_v30).Cardinality())).Times((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))).Plus((p2).Minus((_161_v93).Cardinality()))
			_ = _rhs35
			var _rhs36 _dafny.Int = (func(_pat_let2_0 D10) D10 {
				return func(_162_dt__update__tmp_h1 D10) D10 {
					return func(_pat_let3_0 _dafny.Int) D10 {
						return func(_163_dt__update_hcf38_h0 _dafny.Int) D10 {
							return func(_pat_let4_0 bool) D10 {
								return func(_164_dt__update_hcf37_h0 bool) D10 {
									return func(_pat_let5_0 _dafny.Int) D10 {
										return func(_165_dt__update_hcf34_h0 _dafny.Int) D10 {
											return Companion_D10_.Create_DC26_(_165_dt__update_hcf34_h0, (_162_dt__update__tmp_h1).Dtor_cf35(), (_162_dt__update__tmp_h1).Dtor_cf36(), _164_dt__update_hcf37_h0, _163_dt__update_hcf38_h0)
										}(_pat_let5_0)
									}(_dafny.IntOfInt64(-203))
								}(_pat_let4_0)
							}(!(_pat_let_tv3))
						}(_pat_let3_0)
					}((_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(_dafny.Int))
				}(_pat_let2_0)
			}(Companion_D10_.Create_DC26_(p2, _94_v41, (_158_v90).F23(), (_158_v90).F23(), p2))).Dtor_cf34()
			_ = _rhs36
			var _rhs37 *C3 = _159_v91
			_ = _rhs37
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			var _lhs32 _dafny.Array = _93_v40
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))
			_ = _lhs33
			var _lhs34 _dafny.Array = _160_v92
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_160_v92), 0))
			_ = _lhs35
			_158_v90 = _rhs33
			_159_v91 = _rhs34
			_lhs31.F6 = _rhs35
			(_lhs32).ArraySet1(_rhs36, (_lhs33).Int())
			(_lhs34).ArraySet1(_rhs37, (_lhs35).Int())
		} else if _source2.Is_DC15() {
			var _166___mcc_h2 _dafny.Map = _source2.Get_().(D6_DC15).Cf18
			_ = _166___mcc_h2
			var _167_cf18 _dafny.Map = _166___mcc_h2
			_ = _167_cf18
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index39
			(_94_v41).ArraySet1((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), (_index39).Int())
			_78_v30 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_78_v30, _78_v30), _78_v30)
			var _168_v94 _dafny.Array
			_ = _168_v94
			var _nw24 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
			_ = _nw24
			_168_v94 = _nw24
			var _169_v95 *C3
			_ = _169_v95
			var _nw25 *C3 = New_C3_()
			_ = _nw25
			_nw25.Ctor__(_95_v42, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, p2))
			_169_v95 = _nw25
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_168_v94), 0))
			_ = _index40
			(_168_v94).ArraySet1(_169_v95, (_index40).Int())
			var _170_v97 _dafny.Map
			_ = _170_v97
			_170_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, ((Companion_D12_.Create_DC29_(_dafny.SetOf(p2, p2))).Dtor_cf43()).Cardinality())
			var _171_v98 *C0
			_ = _171_v98
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__(_78_v30, func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate((_78_v30).Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _172_v96 _dafny.CodePoint
					_172_v96 = interface{}(_compr_15).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_78_v30, _172_v96) {
						_coll15.Add(_172_v96, p0)
					}
				}
				return _coll15.ToMap()
			}(), _95_v42, (_170_v97).Merge(_170_v97))
			_171_v98 = _nw26
			var _173_v99 _dafny.Array
			_ = _173_v99
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_2
			var _nw27 _dafny.Array
			_ = _nw27
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw27 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Map = (func(_174_v47 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_175_i9 _dafny.Int) _dafny.Map {
						return _174_v47
					}
				})(_100_v47)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw27 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw27).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw27).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_173_v99 = _nw27
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_168_v94), 0))
			_ = _index41
			var _rhs38 *C3 = _169_v95
			_ = _rhs38
			var _rhs39 *C0 = _171_v98
			_ = _rhs39
			var _rhs40 _dafny.Array = _173_v99
			_ = _rhs40
			var _lhs36 _dafny.Array = _168_v94
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_168_v94), 0))
			_ = _lhs37
			(_lhs36).ArraySet1(_rhs38, (_lhs37).Int())
			_171_v98 = _rhs39
			_173_v99 = _rhs40
			var _176_v100 _dafny.Set
			_ = _176_v100
			_176_v100 = _dafny.SetOf(p1)
			var _177_v101 _dafny.Map
			_ = _177_v101
			_177_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _178_v102 _dafny.Map
			_ = _178_v102
			_178_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), _177_v101)
			(globalState).F6 = ((func() _dafny.Map {
				if (_176_v100).IsSubsetOf(_176_v100) {
					return _177_v101
				}
				return ((func() _dafny.Map {
					if (_178_v102).Contains(p1) {
						return (_178_v102).Get(p1).(_dafny.Map)
					}
					return _177_v101
				})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))
			})()).Cardinality()
		} else {
			var _179___mcc_h3 D6 = _source2.Get_().(D6_DC17).Cf21
			_ = _179___mcc_h3
			var _180_cf21 D6 = _179___mcc_h3
			_ = _180_cf21
			var _181_v103 _dafny.Map
			_ = _181_v103
			_181_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if true {
					return p1
				}
				return p1
			})(), (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))
			_181_v103 = _181_v103
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index42
			(_94_v41).ArraySet1((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index43
			(_94_v41).ArraySet1((p2).Cmp(p2) < 0, (_index43).Int())
			(globalState).F6 = p0
		}
		var _rhs41 _dafny.Int = (func() _dafny.Int {
			if (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool) {
				return Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(362))
			}
			return (p0).Plus(p2)
		})()
		_ = _rhs41
		var _rhs42 _dafny.Array = _94_v41
		_ = _rhs42
		var _lhs38 *GlobalState = globalState
		_ = _lhs38
		_lhs38.F6 = _rhs41
		_94_v41 = _rhs42
		var _182_v104 _dafny.Array
		_ = _182_v104
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
		_ = _nw28
		_182_v104 = _nw28
		var _183_v105 D6
		_ = _183_v105
		_183_v105 = Companion_D6_.Create_DC16_(p1, _182_v104)
		var _184_v106 D6
		_ = _184_v106
		_184_v106 = Companion_D6_.Create_DC17_(_183_v105)
		var _185_v107 D6
		_ = _185_v107
		_185_v107 = Companion_D6_.Create_DC17_((func() D6 {
			if p1 {
				return _183_v105
			}
			return _184_v106
		})())
		_185_v107 = _185_v107
	} else {
		var _186_v108 _dafny.Map
		_ = _186_v108
		_186_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v30, (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))
		var _187_v109 *C2
		_ = _187_v109
		var _nw29 *C2 = New_C2_()
		_ = _nw29
		_nw29.Ctor__(_dafny.CodePoint('f'), _186_v108)
		_187_v109 = _nw29
		var _188_v110 _dafny.Sequence
		_ = _188_v110
		_188_v110 = _dafny.SeqOf(_187_v109)
		var _189_v111 _dafny.Set
		_ = _189_v111
		_189_v111 = _dafny.SetOf(_dafny.IntOfUint32((_188_v110).Cardinality()), (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))
		if (_189_v111).IsDisjointFrom(_189_v111) {
			var _190_v112 _dafny.Map
			_ = _190_v112
			_190_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)) == ((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)), p0)
			var _191_v113 T0
			_ = _191_v113
			var _nw30 *C0 = New_C0_()
			_ = _nw30
			_nw30.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_192_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_193_i10 _dafny.Int) _dafny.CodePoint {
					return _192_v42
				}
			})(_95_v42))), _97_v45, _95_v42, Companion_Default___.Fm18((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), globalState))
			_191_v113 = _nw30
			var _194_v114 _dafny.Sequence
			_ = _194_v114
			_194_v114 = _dafny.SeqOf(_191_v113)
			_190_v112 = (_190_v112).Update(!_dafny.Companion_Sequence_.Equal(_194_v114, _194_v114), Companion_Default___.SafeDivisionInt((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), p0))
			var _195_v115 _dafny.Sequence
			_ = _195_v115
			_195_v115 = _dafny.SeqOf(p1, p1)
			var _196_v116 _dafny.Sequence
			_ = _196_v116
			_196_v116 = _dafny.SeqOf(p1, (_195_v115).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-484), _dafny.IntOfUint32((_195_v115).Cardinality()))).Uint32()).(bool), p1)
			var _197_v117 _dafny.Map
			_ = _197_v117
			_197_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_195_v115).Cardinality()), p2)
			_197_v117 = _197_v117
			var _198_v118 *C3
			_ = _198_v118
			var _nw31 *C3 = New_C3_()
			_ = _nw31
			_nw31.Ctor__(_dafny.CodePoint('p'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jfumgxqqe"), p0))
			_198_v118 = _nw31
			var _199_v119 D9
			_ = _199_v119
			_199_v119 = Companion_D9_.Create_DC24_(false)
			var _200_v120 _dafny.Map
			_ = _200_v120
			_200_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v118, (p1) == (!((_199_v119).Dtor_cf32())))
			_200_v120 = (_200_v120).Update(_198_v118, p1)
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _index44
			var _rhs43 bool = (_187_v109).Fm3(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_78_v30, _78_v30)).Cardinality()), globalState)
			_ = _rhs43
			var _rhs44 _dafny.CodePoint = _95_v42
			_ = _rhs44
			var _rhs45 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfUint32((_dafny.SeqOf(_94_v41, _94_v41)).Cardinality())), (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))
			_ = _rhs45
			var _lhs39 _dafny.Array = _94_v41
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))
			_ = _lhs40
			var _lhs41 T0 = _191_v113
			_ = _lhs41
			var _lhs42 *GlobalState = globalState
			_ = _lhs42
			(_lhs39).ArraySet1(_rhs43, (_lhs40).Int())
			_lhs41.F18_set_(_rhs44)
			_lhs42.F6 = _rhs45
			var _201_v122 _dafny.Sequence
			_ = _201_v122
			_201_v122 = _dafny.SeqOf(_dafny.SetOf(p0), (func() _dafny.Set {
				var _coll16 = _dafny.NewBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-685), _dafny.IntOfInt64(204))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _202_v121 _dafny.Int
					_202_v121 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(-685)).Cmp(_202_v121) <= 0) && ((_202_v121).Cmp(_dafny.IntOfInt64(204)) < 0) {
						_coll16.Add(Companion_Default___.SafeModuloInt(_202_v121, (_dafny.SetOf((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))).Cardinality()))
					}
				}
				return _coll16.ToSet()
			}()).Intersection(_189_v111))
			_201_v122 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_189_v111), _dafny.SeqOf(_189_v111, _189_v111))
		} else {
			r2 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus(p0)).Minus(p2), (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int)))
			var _203_v123 _dafny.Array
			_ = _203_v123
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_3
			var _nw32 _dafny.Array
			_ = _nw32
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw32 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_204_v30 _dafny.Sequence, _205_p2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_206_i11 _dafny.Int) _dafny.CodePoint {
						return (_204_v30).Select((Companion_Default___.SafeIndex(_205_p2, _dafny.IntOfUint32((_204_v30).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_78_v30, p2)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw32 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw32).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw32).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_203_v123 = _nw32
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_203_v123), 0))
			_ = _index45
			(_203_v123).ArraySet1CodePoint(_95_v42, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_203_v123), 0))
			_ = _index46
			(_203_v123).ArraySet1CodePoint((Companion_D3_.Create_DC8_(_95_v42)).Dtor_cf10(), (_index46).Int())
			var _207_v124 _dafny.Map
			_ = _207_v124
			_207_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v111, _dafny.SeqOf(p1))
			var _208_v125 _dafny.Sequence
			_ = _208_v125
			_208_v125 = _dafny.SeqOf(_189_v111)
			var _209_v126 D9
			_ = _209_v126
			_209_v126 = Companion_D9_.Create_DC23_(_208_v125, p2, p1, (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))
			var _210_v127 _dafny.Sequence
			_ = _210_v127
			_210_v127 = _dafny.SeqOf((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), (_209_v126).Dtor_cf30())
			var _211_v128 _dafny.Sequence
			_ = _211_v128
			_211_v128 = _dafny.SeqOf(_210_v127)
			var _212_v129 _dafny.Map
			_ = _212_v129
			_212_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v127, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
			var _213_v130 _dafny.Map
			_ = _213_v130
			_213_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v129, _210_v127)
			var _214_v132 _dafny.Array
			_ = _214_v132
			var _nwElement0_5 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_207_v124).Contains(_189_v111) {
					return (_207_v124).Get(_189_v111).(_dafny.Sequence)
				}
				return _210_v127
			})(), _210_v127)
			_ = _nwElement0_5
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(6))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_5, 0)
			(_nw33).ArraySet1(_210_v127, 1)
			(_nw33).ArraySet1(_210_v127, 2)
			(_nw33).ArraySet1((_211_v128).Select((Companion_Default___.SafeIndex((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_211_v128).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
			(_nw33).ArraySet1(_210_v127, 4)
			(_nw33).ArraySet1((func() _dafny.Sequence {
				if (_213_v130).Contains(func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter17 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-229))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}((func(_215_v127 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_216_i12 _dafny.Int) _dafny.Sequence {
							return _215_v127
						}
					})(_210_v127)))).Elements()); ; {
						_compr_17, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _217_v131 _dafny.Sequence
						_217_v131 = interface{}(_compr_17).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-229))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg19 _dafny.Int) interface{} {
								return coer19(arg19)
							}
						}((func(_218_v127 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_216_i12 _dafny.Int) _dafny.Sequence {
								return _218_v127
							}
						})(_210_v127))), _217_v131) {
							_coll17.Add(_217_v131, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
						}
					}
					return _coll17.ToMap()
				}()) {
					return (_213_v130).Get(func() _dafny.Map {
						var _coll18 = _dafny.NewMapBuilder()
						_ = _coll18
						for _iter18 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-229))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg20 _dafny.Int) interface{} {
								return coer20(arg20)
							}
						}((func(_219_v127 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_220_i12 _dafny.Int) _dafny.Sequence {
								return _219_v127
							}
						})(_210_v127)))).Elements()); ; {
							_compr_18, _ok18 := _iter18()
							if !_ok18 {
								break
							}
							var _221_v131 _dafny.Sequence
							_221_v131 = interface{}(_compr_18).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-229))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg21 _dafny.Int) interface{} {
									return coer21(arg21)
								}
							}((func(_222_v127 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_220_i12 _dafny.Int) _dafny.Sequence {
									return _222_v127
								}
							})(_210_v127))), _221_v131) {
								_coll18.Add(_221_v131, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
							}
						}
						return _coll18.ToMap()
					}()).(_dafny.Sequence)
				}
				return _dafny.SeqOf(p1, (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool))
			})(), 5)
			_214_v132 = _nw33
			_214_v132 = _214_v132
			var _223_v133 _dafny.Set
			_ = _223_v133
			_223_v133 = _dafny.SetOf(p1)
			var _224_v134 _dafny.Sequence
			_ = _224_v134
			_224_v134 = _dafny.SeqOf((_223_v133).Cardinality())
			var _225_v135 _dafny.Sequence
			_ = _225_v135
			_225_v135 = _dafny.SeqOf(_186_v108)
			var _226_v136 *C1
			_ = _226_v136
			var _nw34 *C1 = New_C1_()
			_ = _nw34
			_nw34.Ctor__((_dafny.Zero).Minus((_224_v134).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_224_v134).Cardinality()))).Uint32()).(_dafny.Int)), p1, (_203_v123).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_203_v123), 0))).Int()), (_225_v135).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_210_v127).Cardinality()), _dafny.IntOfUint32((_225_v135).Cardinality()))).Uint32()).(_dafny.Map))
			_226_v136 = _nw34
			r2 = p2
		}
		r1 = (func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate((_189_v111).Elements()); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _227_v137 _dafny.Int
				_227_v137 = interface{}(_compr_19).(_dafny.Int)
				if (_189_v111).Contains(_227_v137) {
					_coll19.Add(Companion_Default___.SafeModuloInt(_227_v137, _dafny.IntOfInt64(36)))
				}
			}
			return _coll19.ToSet()
		}()).IsDisjointFrom(_189_v111)
		var _228_v138 _dafny.Array
		_ = _228_v138
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
		_ = _nw35
		_228_v138 = _nw35
		var _229_v139 _dafny.MultiSet
		_ = _229_v139
		_229_v139 = _dafny.MultiSetOf(_94_v41)
		var _230_v140 _dafny.Map
		_ = _230_v140
		_230_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _229_v139)
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_228_v138), 0))
		_ = _index47
		(_228_v138).ArraySet1((func() _dafny.MultiSet {
			if (_230_v140).Contains(p1) {
				return (_230_v140).Get(p1).(_dafny.MultiSet)
			}
			return _229_v139
		})(), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_228_v138), 0))
		_ = _index48
		(_228_v138).ArraySet1(_229_v139, (_index48).Int())
		var _231_v141 _dafny.Sequence
		_ = _231_v141
		_231_v141 = _dafny.SeqOf((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(globalState))
		(globalState).F1 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_187_v109).Fm3(p1, p0, globalState) {
				return _231_v141
			}
			return _231_v141
		})(), _231_v141)
		var _232_v142 _dafny.Set
		_ = _232_v142
		_232_v142 = _dafny.SetOf(p1)
		var _233_v143 D12
		_ = _233_v143
		_233_v143 = Companion_D12_.Create_DC31_(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(758))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_234_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_235_i13 _dafny.Int) _dafny.CodePoint {
				return _234_v42
			}
		})(_95_v42))), _dafny.IntOfUint32((Companion_Default___.Fm30((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), globalState)).Cardinality()), p1)
		var _rhs46 _dafny.Int = (func() _dafny.Int {
			if !((_232_v142).IsDisjointFrom(_232_v142)) {
				return (_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_233_v143).Dtor_cf47(), _78_v30)).Cardinality())
		})()
		_ = _rhs46
		var _rhs47 _dafny.Int = (func() _dafny.Int {
			if p1 {
				return ((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int)), _94_v41)).Cardinality())
			}
			return p0
		})()
		_ = _rhs47
		var _lhs43 *GlobalState = globalState
		_ = _lhs43
		_lhs43.F6 = _rhs46
		r2 = _rhs47
	}
	var _236_v144 _dafny.Map
	_ = _236_v144
	_236_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_78_v30, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_78_v30).Cardinality()))).Uint32(), _95_v42), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_78_v30, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_78_v30).Cardinality()))).Uint32(), _95_v42)).Cardinality()))).Uint32(), _95_v42), p2)
	var _237_v145 *C2
	_ = _237_v145
	var _nw36 *C2 = New_C2_()
	_ = _nw36
	_nw36.Ctor__(_95_v42, _236_v144)
	_237_v145 = _nw36
	r0 = _94_v41
	r1 = (func() bool {
		if p1 {
			return (_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool)
		}
		return p1
	})()
	var _238_v146 _dafny.Set
	_ = _238_v146
	_238_v146 = _dafny.SetOf((_93_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_93_v40), 0))).Int()).(_dafny.Int))
	r2 = (_237_v145).Fm4((_94_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_94_v41), 0))).Int()).(bool), _95_v42, Companion_Default___.Fm10(globalState), (_238_v146).IsProperSubsetOf(_238_v146), globalState)
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _239_v0 _dafny.Int
	_ = _239_v0
	_239_v0 = _dafny.IntOfInt64(140)
	var _240_v1 _dafny.Sequence
	_ = _240_v1
	_240_v1 = _dafny.SeqOf((_dafny.Zero).Minus(_239_v0))
	var _241_v2 bool
	_ = _241_v2
	_241_v2 = true
	var _242_v3 _dafny.Map
	_ = _242_v3
	_242_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _241_v2)
	var _243_v4 _dafny.CodePoint
	_ = _243_v4
	_243_v4 = _dafny.CodePoint('q')
	var _244_v5 _dafny.Map
	_ = _244_v5
	_244_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v4, (_dafny.Zero).Minus(_239_v0))
	var _245_v7 _dafny.Array
	_ = _245_v7
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_4
	var _nw37 _dafny.Array
	_ = _nw37
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw37 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Map = (func(_246_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_247_i0 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(399), _dafny.IntOfInt64(-703))); ; {
						_compr_20, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _248_v6 _dafny.Int
						_248_v6 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(399)).Cmp(_248_v6) <= 0) && ((_248_v6).Cmp(_dafny.IntOfInt64(-703)) < 0) {
							_coll20.Add(Companion_Default___.SafeModuloInt(_248_v6, _dafny.IntOfInt64(668)), _246_v0)
						}
					}
					return _coll20.ToMap()
				}()
			}
		})(_239_v0)
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
	_245_v7 = _nw37
	var _249_v8 _dafny.Set
	_ = _249_v8
	_249_v8 = _dafny.SetOf(_241_v2)
	var _250_v9 _dafny.Array
	_ = _250_v9
	var _nwElement0_6 _dafny.Int = _239_v0
	_ = _nwElement0_6
	var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(12))
	_ = _nw38
	(_nw38).ArraySet1(_nwElement0_6, 0)
	(_nw38).ArraySet1(_239_v0, 1)
	(_nw38).ArraySet1(_239_v0, 2)
	(_nw38).ArraySet1(_239_v0, 3)
	(_nw38).ArraySet1(_239_v0, 4)
	(_nw38).ArraySet1(_dafny.IntOfInt64(44), 5)
	(_nw38).ArraySet1(_239_v0, 6)
	(_nw38).ArraySet1(_239_v0, 7)
	(_nw38).ArraySet1(_239_v0, 8)
	(_nw38).ArraySet1(_239_v0, 9)
	(_nw38).ArraySet1(_239_v0, 10)
	(_nw38).ArraySet1((_249_v8).Cardinality(), 11)
	_250_v9 = _nw38
	var _251_v10 _dafny.Array
	_ = _251_v10
	var _nw39 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
	_ = _nw39
	_251_v10 = _nw39
	var _252_v11 _dafny.Map
	_ = _252_v11
	_252_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_250_v9, _251_v10)
	var _253_v12 _dafny.Sequence
	_ = _253_v12
	_253_v12 = _dafny.UnicodeSeqOfUtf8Bytes("oalh")
	var _254_v13 _dafny.Sequence
	_ = _254_v13
	_254_v13 = _dafny.SeqOf(_250_v9)
	var _255_v14 _dafny.Sequence
	_ = _255_v14
	_255_v14 = _dafny.SeqOf(false, _241_v2)
	var _256_v15 _dafny.Map
	_ = _256_v15
	_256_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v0, _241_v2)
	var _257_v16 _dafny.Map
	_ = _257_v16
	_257_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v0, (func() bool {
		if (_256_v15).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qgoah")).Cardinality())) {
			return (_256_v15).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qgoah")).Cardinality())).(bool)
		}
		return _241_v2
	})())
	var _258_v17 _dafny.Sequence
	_ = _258_v17
	_258_v17 = _dafny.SeqOf(_257_v16)
	var _259_globalState *GlobalState
	_ = _259_globalState
	var _nw40 *GlobalState = New_GlobalState_()
	_ = _nw40
	_nw40.Ctor__(_dafny.IntOfInt64(515), _240_v1, _240_v1, _242_v3, _dafny.IntOfInt64(523), _244_v5, _dafny.IntOfInt64(914), _245_v7, _252_v11, _253_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v10, _241_v2), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_250_v9, _250_v9, _250_v9, _250_v9, _250_v9), _254_v13), false, _255_v14, false, _dafny.IntOfInt64(682), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_257_v16), _258_v17), true)
	_259_globalState = _nw40
	var _hi1 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(902), _239_v0))
	_ = _hi1
	for _260_i1 := _239_v0; _260_i1.Cmp(_hi1) < 0; _260_i1 = _260_i1.Plus(_dafny.One) {
		_241_v2 = ((func() _dafny.Int {
			if _241_v2 {
				return _260_i1
			}
			return (_dafny.Zero).Minus(_260_i1)
		})()).Cmp(((_dafny.Zero).Minus(_239_v0)).Times(_239_v0)) < 0
		var _261_v18 _dafny.Sequence
		_ = _261_v18
		_261_v18 = _dafny.SeqOf(_253_v12, _253_v12)
		_261_v18 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_253_v12), _261_v18)
		var _262_v19 _dafny.Set
		_ = _262_v19
		_262_v19 = _dafny.SetOf(_250_v9)
		var _rhs48 bool = _241_v2
		_ = _rhs48
		var _rhs49 bool = _241_v2
		_ = _rhs49
		var _rhs50 _dafny.Sequence = _240_v1
		_ = _rhs50
		var _rhs51 _dafny.Int = Companion_Default___.SafeDivisionInt(((_262_v19).Intersection(_262_v19)).Cardinality(), _dafny.IntOfUint32((_240_v1).Cardinality()))
		_ = _rhs51
		var _lhs44 *GlobalState = _259_globalState
		_ = _lhs44
		_241_v2 = _rhs48
		_241_v2 = _rhs49
		_lhs44.F1 = _rhs50
		_239_v0 = _rhs51
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_251_v10), 0))
		_ = _index49
		(_251_v10).ArraySet1(_241_v2, (_index49).Int())
		var _263_v20 _dafny.Map
		_ = _263_v20
		_263_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_259_globalState), (func() _dafny.CodePoint {
			if _241_v2 {
				return _243_v4
			}
			return _243_v4
		})())
		var _264_v21 D0
		_ = _264_v21
		_264_v21 = Companion_D0_.Create_DC0_(Companion_Default___.Fm0(_259_globalState))
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_251_v10), 0))
		_ = _index50
		var _rhs52 _dafny.CodePoint = (func() _dafny.CodePoint {
			if (_263_v20).Contains(_dafny.IntOfInt64(288)) {
				return (_263_v20).Get(_dafny.IntOfInt64(288)).(_dafny.CodePoint)
			}
			return _243_v4
		})()
		_ = _rhs52
		var _rhs53 _dafny.Int = ((_264_v21).Dtor_cf0()).Minus(_239_v0)
		_ = _rhs53
		var _rhs54 bool = _241_v2
		_ = _rhs54
		var _rhs55 bool = (func() bool {
			if false {
				return false
			}
			return _241_v2
		})()
		_ = _rhs55
		var _rhs56 _dafny.Int = (_264_v21).Dtor_cf0()
		_ = _rhs56
		var _lhs45 *GlobalState = _259_globalState
		_ = _lhs45
		var _lhs46 _dafny.Array = _251_v10
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_251_v10), 0))
		_ = _lhs47
		_243_v4 = _rhs52
		_lhs45.F6 = _rhs53
		(_lhs46).ArraySet1(_rhs54, (_lhs47).Int())
		_241_v2 = _rhs55
		_239_v0 = _rhs56
	}
	var _265_i2 _dafny.Int
	_ = _265_i2
	_265_i2 = _dafny.Zero
	{
		for _241_v2 {
			{
				if (_265_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_265_i2 = (_265_i2).Plus(_dafny.One)
				_242_v3 = (_242_v3).Update(!((_dafny.IntOfInt64(904)).Cmp(_239_v0) <= 0), _241_v2)
				var _rhs57 _dafny.CodePoint = _243_v4
				_ = _rhs57
				var _rhs58 _dafny.Int = _239_v0
				_ = _rhs58
				var _lhs48 *GlobalState = _259_globalState
				_ = _lhs48
				_243_v4 = _rhs57
				_lhs48.F6 = _rhs58
				var _266_v22 _dafny.Map
				_ = _266_v22
				_266_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v10, true)
				(_259_globalState).F6 = (func() _dafny.Int {
					if _241_v2 {
						return _239_v0
					}
					return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-533))).Times((_266_v22).Cardinality())
				})()
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_250_v9), 0))
				_ = _index51
				(_250_v9).ArraySet1(_239_v0, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_250_v9), 0))
				_ = _index52
				(_250_v9).ArraySet1((_dafny.IntOfInt64(-29)).Minus(_239_v0), (_index52).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
	_ = _index53
	(_250_v9).ArraySet1(_239_v0, (_index53).Int())
	var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
	_ = _index54
	(_250_v9).ArraySet1(Companion_Default___.SafeModuloInt(_239_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nxertem")).Cardinality()))), (_index54).Int())
	var _267_v23 _dafny.Array
	_ = _267_v23
	var _268_v24 bool
	_ = _268_v24
	var _269_v25 _dafny.Int
	_ = _269_v25
	var _out0 _dafny.Array
	_ = _out0
	var _out1 bool
	_ = _out1
	var _out2 _dafny.Int
	_ = _out2
	_out0, _out1, _out2 = Companion_Default___.M0((_dafny.Zero).Minus((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)), Companion_Default___.Fm1((func() bool {
		if (_256_v15).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mjgtu")).Cardinality())) {
			return (_256_v15).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mjgtu")).Cardinality())).(bool)
		}
		return _241_v2
	})(), _243_v4, _dafny.IntOfInt64(-261), _259_globalState), (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_253_v12, (Companion_Default___.SafeIndex(_239_v0, _dafny.IntOfUint32((_253_v12).Cardinality()))).Uint32(), _243_v4)).Cardinality())).Times((_dafny.Zero).Minus(_239_v0)), _259_globalState)
	_267_v23 = _out0
	_268_v24 = _out1
	_269_v25 = _out2
	_253_v12 = _253_v12
	var _270_v26 _dafny.Map
	_ = _270_v26
	_270_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v24, _255_v14)
	var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
	_ = _index55
	var _rhs59 _dafny.Int = Companion_Default___.Fm0(_259_globalState)
	_ = _rhs59
	var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if (_270_v26).Contains(_268_v24) {
			return (_270_v26).Get(_268_v24).(_dafny.Sequence)
		}
		return _255_v14
	})(), _255_v14)
	_ = _rhs60
	var _lhs49 _dafny.Array = _250_v9
	_ = _lhs49
	var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
	_ = _lhs50
	(_lhs49).ArraySet1(_rhs59, (_lhs50).Int())
	_255_v14 = _rhs60
	_268_v24 = _241_v2
	var _271_v27 _dafny.Array
	_ = _271_v27
	var _len0_5 _dafny.Int = _dafny.One
	_ = _len0_5
	var _nw41 _dafny.Array
	_ = _nw41
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw41 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_272_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_273_i3 _dafny.Int) _dafny.CodePoint {
				return _272_v4
			}
		})(_243_v4)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw41 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw41).ArraySet1CodePoint(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw41).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_271_v27 = _nw41
	var _274_v28 _dafny.Map
	_ = _274_v28
	_274_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _271_v27)
	if (_dafny.SetOf(Companion_Default___.Fm1(_241_v2, _243_v4, (_274_v28).Cardinality(), _259_globalState), _268_v24, (func() bool {
		if (_242_v3).Contains(_268_v24) {
			return (_242_v3).Get(_268_v24).(bool)
		}
		return _268_v24
	})())).IsSubsetOf(Companion_Default___.Fm2(_259_globalState)) {
		var _275_v29 D0
		_ = _275_v29
		_275_v29 = Companion_D0_.Create_DC1_()
		_275_v29 = _275_v29
		var _276_v30 _dafny.Map
		_ = _276_v30
		_276_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v2, _dafny.IntOfInt64(531))
		var _277_v31 *C3
		_ = _277_v31
		var _nw42 *C3 = New_C3_()
		_ = _nw42
		_nw42.Ctor__(Companion_Default___.Fm20(false, _259_globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, (_276_v30).Cardinality()))
		_277_v31 = _nw42
		(_259_globalState).F6 = ((func() _dafny.Map {
			var _coll21 = _dafny.NewMapBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(584), _dafny.IntOfInt64(334))); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _278_v32 _dafny.Int
				_278_v32 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(584)).Cmp(_278_v32) <= 0) && ((_278_v32).Cmp(_dafny.IntOfInt64(334)) < 0) {
					_coll21.Add((_278_v32).Minus(_239_v0), Companion_D5_.Create_DC12_(_dafny.MultiSetOf(_241_v2)))
				}
			}
			return _coll21.ToMap()
		}()).Cardinality()).Times((_269_v25).Plus((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)))
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_267_v23), 0))
		_ = _index56
		(_267_v23).ArraySet1(!((func() bool {
			if _268_v24 {
				return _268_v24
			}
			return _241_v2
		})()), (_index56).Int())
		var _279_v33 _dafny.Set
		_ = _279_v33
		_279_v33 = _dafny.SetOf(_269_v25)
		var _280_v34 _dafny.Sequence
		_ = _280_v34
		_280_v34 = _dafny.SeqOf(_279_v33)
		var _281_v35 _dafny.Map
		_ = _281_v35
		_281_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v0, _280_v34)
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_267_v23), 0))
		_ = _index57
		var _rhs61 bool = !(_268_v24)
		_ = _rhs61
		var _rhs62 bool = _241_v2
		_ = _rhs62
		var _rhs63 D0 = Companion_Default___.Fm25(_242_v3, Companion_D9_.Create_DC23_((func() _dafny.Sequence {
			if (_281_v35).Contains(_239_v0) {
				return (_281_v35).Get(_239_v0).(_dafny.Sequence)
			}
			return _280_v34
		})(), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _241_v2, _239_v0), Companion_Default___.Fm1(false, _243_v4, _239_v0, _259_globalState), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _259_globalState)
		_ = _rhs63
		var _rhs64 bool = !(_241_v2) || ((_269_v25).Cmp((Companion_Default___.Fm26(_259_globalState)).Cardinality()) < 0)
		_ = _rhs64
		var _rhs65 bool = !(_268_v24)
		_ = _rhs65
		var _lhs51 _dafny.Array = _267_v23
		_ = _lhs51
		var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_267_v23), 0))
		_ = _lhs52
		(_lhs51).ArraySet1(_rhs61, (_lhs52).Int())
		_241_v2 = _rhs62
		_275_v29 = _rhs63
		_241_v2 = _rhs64
		_268_v24 = _rhs65
		_241_v2 = (_269_v25).Cmp(_269_v25) != 0
	} else {
		var _282_v36 _dafny.Array
		_ = _282_v36
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
		_ = _nw43
		_282_v36 = _nw43
		var _283_v37 _dafny.Map
		_ = _283_v37
		_283_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_v36, _251_v10)
		_283_v37 = (_283_v37).Update(_282_v36, _267_v23)
		var _284_v38 _dafny.Array
		_ = _284_v38
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_6
		var _nw44 _dafny.Array
		_ = _nw44
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw44 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Sequence = (func(_285_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_286_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_285_v14, _285_v14)
				}
			})(_255_v14)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw44 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw44).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw44).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_284_v38 = _nw44
		_284_v38 = _284_v38
		var _287_v39 _dafny.Array
		_ = _287_v39
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw45
		_287_v39 = _nw45
		var _288_v40 D6
		_ = _288_v40
		_288_v40 = Companion_D6_.Create_DC16_(_241_v2, _287_v39)
		var _source3 D6 = _288_v40
		_ = _source3
		if _source3.Is_DC16() {
			var _289___mcc_h0 bool = _source3.Get_().(D6_DC16).Cf19
			_ = _289___mcc_h0
			var _290___mcc_h1 _dafny.Array = _source3.Get_().(D6_DC16).Cf20
			_ = _290___mcc_h1
			var _291_cf20 _dafny.Array = _290___mcc_h1
			_ = _291_cf20
			var _292_cf19 bool = _289___mcc_h0
			_ = _292_cf19
			var _293_v41 D7
			_ = _293_v41
			_293_v41 = Companion_D7_.Create_DC18_(_249_v8)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
			_ = _index58
			(_250_v9).ArraySet1((_269_v25).Minus((((_293_v41).Dtor_cf22()).Cardinality()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_292_cf19, _269_v25)).Cardinality())), (_index58).Int())
			_292_cf19 = _268_v24
			var _294_v42 _dafny.MultiSet
			_ = _294_v42
			_294_v42 = _dafny.MultiSetOf((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_251_v10), 0))
			_ = _index59
			(_251_v10).ArraySet1(!(!(_294_v42).Contains(_dafny.IntOfUint32((_253_v12).Cardinality()))), (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_251_v10), 0))
			_ = _index60
			(_251_v10).ArraySet1(_241_v2, (_index60).Int())
			_239_v0 = Companion_Default___.Fm0(_259_globalState)
		} else if _source3.Is_DC15() {
			var _295___mcc_h2 _dafny.Map = _source3.Get_().(D6_DC15).Cf18
			_ = _295___mcc_h2
			var _296_cf18 _dafny.Map = _295___mcc_h2
			_ = _296_cf18
			var _297_v43 _dafny.Array
			_ = _297_v43
			var _298_v44 bool
			_ = _298_v44
			var _299_v45 _dafny.Int
			_ = _299_v45
			var _out3 _dafny.Array
			_ = _out3
			var _out4 bool
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out3, _out4, _out5 = Companion_Default___.M0(_dafny.IntOfUint32((_253_v12).Cardinality()), _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_300_i5 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(629)
			})), Companion_Default___.Fm0(_259_globalState)), _dafny.IntOfUint32((_253_v12).Cardinality()), _259_globalState)
			_297_v43 = _out3
			_298_v44 = _out4
			_299_v45 = _out5
			var _301_v46 *C3
			_ = _301_v46
			var _nw46 *C3 = New_C3_()
			_ = _nw46
			_nw46.Ctor__(_243_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, _239_v0))
			_301_v46 = _nw46
			_301_v46 = _301_v46
			var _302_v47 _dafny.MultiSet
			_ = _302_v47
			_302_v47 = _dafny.MultiSetOf(_298_v44, true)
			var _303_v48 D3
			_ = _303_v48
			_303_v48 = Companion_D3_.Create_DC10_((_dafny.MultiSetFromSeq(_dafny.SeqOf(_241_v2))).IsProperSubsetOf(_302_v47))
			_303_v48 = _303_v48
			var _304_v49 _dafny.Array
			_ = _304_v49
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw47
			_304_v49 = _nw47
			var _305_v50 _dafny.Map
			_ = _305_v50
			_305_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v2, _282_v36)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_304_v49), 0))
			_ = _index61
			(_304_v49).ArraySet1((func() _dafny.Array {
				if (_305_v50).Contains(_298_v44) {
					return (_305_v50).Get(_298_v44).(_dafny.Array)
				}
				return _282_v36
			})(), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_304_v49), 0))
			_ = _index62
			(_304_v49).ArraySet1(_282_v36, (_index62).Int())
		} else {
			var _306___mcc_h3 D6 = _source3.Get_().(D6_DC17).Cf21
			_ = _306___mcc_h3
			var _307_cf21 D6 = _306___mcc_h3
			_ = _307_cf21
			var _308_v51 _dafny.Map
			_ = _308_v51
			_308_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, _dafny.IntOfUint32(((Companion_D9_.Create_DC22_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(932))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_309_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_310_i6 _dafny.Int) _dafny.CodePoint {
					return _309_v4
				}
			})(_243_v4))))).Dtor_cf27()).Cardinality()))
			var _311_v52 *C3
			_ = _311_v52
			var _nw48 *C3 = New_C3_()
			_ = _nw48
			_nw48.Ctor__(_243_v4, _308_v51)
			_311_v52 = _nw48
			_311_v52 = _311_v52
			var _312_v53 _dafny.Map
			_ = _312_v53
			_312_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v24, (func() _dafny.Array {
				if _241_v2 {
					return _250_v9
				}
				return (_254_v13).Select((Companion_Default___.SafeIndex(_239_v0, _dafny.IntOfUint32((_254_v13).Cardinality()))).Uint32()).(_dafny.Array)
			})())
			_250_v9 = (func() _dafny.Array {
				if (_312_v53).Contains((_311_v52).Fm3(_241_v2, (_dafny.Zero).Minus(_239_v0), _259_globalState)) {
					return (_312_v53).Get((_311_v52).Fm3(_241_v2, (_dafny.Zero).Minus(_239_v0), _259_globalState)).(_dafny.Array)
				}
				return _250_v9
			})()
			var _313_v54 *C3
			_ = _313_v54
			var _nw49 *C3 = New_C3_()
			_ = _nw49
			_nw49.Ctor__(_243_v4, (_308_v51).Merge(_308_v51))
			_313_v54 = _nw49
			var _314_v55 D0
			_ = _314_v55
			var _315_v56 _dafny.Int
			_ = _315_v56
			var _316_v57 bool
			_ = _316_v57
			var _317_v58 bool
			_ = _317_v58
			var _out6 D0
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 bool
			_ = _out8
			var _out9 bool
			_ = _out9
			_out6, _out7, _out8, _out9 = (_311_v52).M1(Companion_Default___.Fm1(_241_v2, _dafny.CodePoint('m'), (_312_v53).Cardinality(), _259_globalState), _259_globalState)
			_314_v55 = _out6
			_315_v56 = _out7
			_316_v57 = _out8
			_317_v58 = _out9
		}
		var _318_v59 _dafny.Map
		_ = _318_v59
		_318_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, _239_v0)
		var _319_v60 *C1
		_ = _319_v60
		var _nw50 *C1 = New_C1_()
		_ = _nw50
		_nw50.Ctor__((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _268_v24, _243_v4, _318_v59)
		_319_v60 = _nw50
		var _320_v61 _dafny.Map
		_ = _320_v61
		_320_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_319_v60, (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_239_v0, _269_v25)))
		_320_v61 = (_320_v61).Update(_319_v60, _269_v25)
		var _321_v62 _dafny.MultiSet
		_ = _321_v62
		_321_v62 = _dafny.MultiSetOf(_268_v24, true)
		var _322_v63 _dafny.Map
		_ = _322_v63
		_322_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v4, (_dafny.MultiSetFromSeq(_dafny.SeqOf((_319_v60).F23()))).Intersection(_321_v62))
		_322_v63 = (_322_v63).Update(_243_v4, _321_v62)
	}
	var _323_v64 _dafny.Array
	_ = _323_v64
	var _nw51 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(18))
	_ = _nw51
	_323_v64 = _nw51
	for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_323_v64), 0))); ; {
		_guard_loop_0, _ok22 := _iter22()
		if !_ok22 {
			break
		}
		var _324_i7 _dafny.Int
		_324_i7 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_324_i7).Sign() != -1) && ((_324_i7).Cmp(_dafny.ArrayLen((_323_v64), 0)) < 0)) {
			(_323_v64).ArraySet1(Companion_D7_.Create_DC20_(_241_v2, _241_v2), (_324_i7).Int())
		}
	}
	var _325_v65 _dafny.Map
	_ = _325_v65
	_325_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v25, _dafny.IntOfInt64(458))
	var _326_i8 _dafny.Int
	_ = _326_i8
	_326_i8 = _dafny.Zero
	{
		for (((_dafny.Zero).Minus((_325_v65).Cardinality())).Times(_dafny.IntOfUint32((_253_v12).Cardinality()))).Cmp(_269_v25) >= 0 {
			{
				if (_326_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_326_i8 = (_326_i8).Plus(_dafny.One)
				_269_v25 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_240_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.IntOfUint32((_240_v1).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_269_v25)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_269_v25, _239_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_240_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.IntOfUint32((_240_v1).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_269_v25))).Cardinality()))).Uint32(), _dafny.IntOfInt64(123))).Cardinality())
				var _327_v66 _dafny.Set
				_ = _327_v66
				_327_v66 = _dafny.SetOf((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))
				var _328_v67 _dafny.Sequence
				_ = _328_v67
				_328_v67 = _dafny.SeqOf(_327_v66, _327_v66)
				_328_v67 = (func() _dafny.Sequence {
					if _241_v2 {
						return _dafny.SeqOf(_327_v66, _327_v66, _327_v66, _327_v66, _dafny.SetOf(_239_v0))
					}
					return _dafny.SeqOf(func() _dafny.Set {
						var _coll22 = _dafny.NewBuilder()
						_ = _coll22
						for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-240), _dafny.IntOfInt64(336))); ; {
							_compr_22, _ok23 := _iter23()
							if !_ok23 {
								break
							}
							var _329_v68 _dafny.Int
							_329_v68 = interface{}(_compr_22).(_dafny.Int)
							if ((_dafny.IntOfInt64(-240)).Cmp(_329_v68) <= 0) && ((_329_v68).Cmp(_dafny.IntOfInt64(336)) < 0) {
								_coll22.Add((_329_v68).Minus((_256_v15).Cardinality()))
							}
						}
						return _coll22.ToSet()
					}(), _dafny.SetOf(_dafny.IntOfInt64(90)), _327_v66, _327_v66, _327_v66)
				})()
				var _330_v70 _dafny.Sequence
				_ = _330_v70
				_330_v70 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}((func(_331_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_332_i9 _dafny.Int) _dafny.CodePoint {
						return _331_v4
					}
				})(_243_v4))))
				var _333_v71 *C2
				_ = _333_v71
				var _nw52 *C2 = New_C2_()
				_ = _nw52
				_nw52.Ctor__(_243_v4, func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter24 := _dafny.Iterate((_330_v70).Elements()); ; {
						_compr_23, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _334_v69 _dafny.Sequence
						_334_v69 = interface{}(_compr_23).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_330_v70, _334_v69) {
							_coll23.Add(_334_v69, (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll23.ToMap()
				}())
				_333_v71 = _nw52
				_325_v65 = (_325_v65).Update((_dafny.Zero).Minus((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(378), _269_v25))
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _335_v72 _dafny.Array
	_ = _335_v72
	var _nw53 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
	_ = _nw53
	_335_v72 = _nw53
	var _336_v73 _dafny.Map
	_ = _336_v73
	_336_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v72, _240_v1)
	_336_v73 = (_336_v73).Update(_335_v72, _240_v1)
	var _337_v74 D1
	_ = _337_v74
	_337_v74 = Companion_D1_.Create_DC6_(_dafny.Companion_Sequence_.IsPrefixOf(_253_v12, _253_v12), _268_v24)
	var _source4 D1 = _337_v74
	_ = _source4
	if _source4.Is_DC4() {
		var _338___mcc_h4 _dafny.Int = _source4.Get_().(D1_DC4).Cf3
		_ = _338___mcc_h4
		var _339___mcc_h5 _dafny.Int = _source4.Get_().(D1_DC4).Cf4
		_ = _339___mcc_h5
		var _340___mcc_h6 bool = _source4.Get_().(D1_DC4).Cf5
		_ = _340___mcc_h6
		var _341_cf5 bool = _340___mcc_h6
		_ = _341_cf5
		var _342_cf4 _dafny.Int = _339___mcc_h5
		_ = _342_cf4
		var _343_cf3 _dafny.Int = _338___mcc_h4
		_ = _343_cf3
		var _344_v75 _dafny.Map
		_ = _344_v75
		_344_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, _342_cf4)
		var _345_v76 T0
		_ = _345_v76
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("ikrmgew"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)), _243_v4, _344_v75)
		_345_v76 = _nw54
		var _346_v77 _dafny.MultiSet
		_ = _346_v77
		_346_v77 = _dafny.MultiSetOf(_345_v76)
		var _347_v78 _dafny.Sequence
		_ = _347_v78
		_347_v78 = _dafny.SeqOf(_dafny.MultiSetOf(_345_v76), _dafny.MultiSetOf(_345_v76), (_346_v77).Update(_345_v76, Companion_Default___.Abs(Companion_Default___.Fm0(_259_globalState))), _dafny.MultiSetOf(_345_v76, _345_v76))
		_347_v78 = _347_v78
		var _348_v79 _dafny.Set
		_ = _348_v79
		_348_v79 = _dafny.SetOf(_dafny.IntOfInt64(788), _269_v25, _dafny.IntOfInt64(979), _343_cf3, (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))
		var _349_v80 _dafny.Sequence
		_ = _349_v80
		_349_v80 = _dafny.SeqOf(_348_v79, _348_v79)
		var _350_v81 _dafny.Array
		_ = _350_v81
		var _351_v82 bool
		_ = _351_v82
		var _352_v83 _dafny.Int
		_ = _352_v83
		var _out10 _dafny.Array
		_ = _out10
		var _out11 bool
		_ = _out11
		var _out12 _dafny.Int
		_ = _out12
		_out10, _out11, _out12 = Companion_Default___.M0((_dafny.Zero).Minus((Companion_D9_.Create_DC23_(_349_v80, _269_v25, _241_v2, _dafny.IntOfInt64(277))).Dtor_cf31()), _241_v2, _239_v0, _259_globalState)
		_350_v81 = _out10
		_351_v82 = _out11
		_352_v83 = _out12
		if _241_v2 {
			_351_v82 = !((_341_cf5) == (Companion_Default___.Fm1(_268_v24, _345_v76.F18(), _239_v0, _259_globalState)))
			(_259_globalState).F6 = _239_v0
			var _353_v86 _dafny.Sequence
			_ = _353_v86
			_353_v86 = _dafny.SeqOf(_251_v10)
			var _354_v87 _dafny.Array
			_ = _354_v87
			var _nwElement0_7 _dafny.Array = _350_v81
			_ = _nwElement0_7
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(29))
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_7, 0)
			(_nw55).ArraySet1(_350_v81, 1)
			(_nw55).ArraySet1(_267_v23, 2)
			(_nw55).ArraySet1(_251_v10, 3)
			(_nw55).ArraySet1(_350_v81, 4)
			(_nw55).ArraySet1(_350_v81, 5)
			(_nw55).ArraySet1(_251_v10, 6)
			(_nw55).ArraySet1(_251_v10, 7)
			(_nw55).ArraySet1(_350_v81, 8)
			(_nw55).ArraySet1(_350_v81, 9)
			(_nw55).ArraySet1(_350_v81, 10)
			(_nw55).ArraySet1(_251_v10, 11)
			(_nw55).ArraySet1(_267_v23, 12)
			(_nw55).ArraySet1(_267_v23, 13)
			(_nw55).ArraySet1(_267_v23, 14)
			(_nw55).ArraySet1(_350_v81, 15)
			(_nw55).ArraySet1(_350_v81, 16)
			(_nw55).ArraySet1(_251_v10, 17)
			(_nw55).ArraySet1(_251_v10, 18)
			(_nw55).ArraySet1(_267_v23, 19)
			(_nw55).ArraySet1(_267_v23, 20)
			(_nw55).ArraySet1(_251_v10, 21)
			(_nw55).ArraySet1(_267_v23, 22)
			(_nw55).ArraySet1(_267_v23, 23)
			(_nw55).ArraySet1(_251_v10, 24)
			(_nw55).ArraySet1(_267_v23, 25)
			(_nw55).ArraySet1(_267_v23, 26)
			(_nw55).ArraySet1(_350_v81, 27)
			(_nw55).ArraySet1((_353_v86).Select((Companion_Default___.SafeIndex((_249_v8).Cardinality(), _dafny.IntOfUint32((_353_v86).Cardinality()))).Uint32()).(_dafny.Array), 28)
			_354_v87 = _nw55
			var _355_v88 D6
			_ = _355_v88
			_355_v88 = Companion_D6_.Create_DC16_(_341_cf5, _354_v87)
			var _356_v89 D6
			_ = _356_v89
			_356_v89 = Companion_D6_.Create_DC17_(_355_v88)
			var _357_v90 _dafny.Map
			_ = _357_v90
			_357_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_356_v89, _341_cf5)
			var _358_v93 _dafny.Array
			_ = _358_v93
			var _nwElement0_8 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_253_v12).Cardinality())), _241_v2)
			_ = _nwElement0_8
			var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(18))
			_ = _nw56
			(_nw56).ArraySet1(_nwElement0_8, 0)
			(_nw56).ArraySet1(_256_v15, 1)
			(_nw56).ArraySet1(_257_v16, 2)
			(_nw56).ArraySet1(_257_v16, 3)
			(_nw56).ArraySet1(_257_v16, 4)
			(_nw56).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_249_v8).Cardinality(), Companion_Default___.Fm1(true, _243_v4, _352_v83, _259_globalState)), 5)
			(_nw56).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_cf4, true), 6)
			(_nw56).ArraySet1(func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(919), _dafny.IntOfInt64(570))); ; {
					_compr_24, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _359_v84 _dafny.Int
					_359_v84 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(919)).Cmp(_359_v84) <= 0) && ((_359_v84).Cmp(_dafny.IntOfInt64(570)) < 0) {
						_coll24.Add(Companion_Default___.SafeModuloInt(_359_v84, _239_v0), _341_cf5)
					}
				}
				return _coll24.ToMap()
			}(), 7)
			(_nw56).ArraySet1(_257_v16, 8)
			(_nw56).ArraySet1(_256_v15, 9)
			(_nw56).ArraySet1(_257_v16, 10)
			(_nw56).ArraySet1((func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(699), _dafny.IntOfInt64(764))); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _360_v85 _dafny.Int
					_360_v85 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(699)).Cmp(_360_v85) <= 0) && ((_360_v85).Cmp(_dafny.IntOfInt64(764)) < 0) {
						_coll25.Add((_360_v85).Plus(_269_v25), _341_cf5)
					}
				}
				return _coll25.ToMap()
			}()).Merge(_256_v15), 11)
			(_nw56).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v0, !((func() bool {
				if (_357_v90).Contains(_356_v89) {
					return (_357_v90).Get(_356_v89).(bool)
				}
				return _241_v2
			})()))).Merge(_257_v16), 12)
			(_nw56).ArraySet1(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(58), _dafny.IntOfInt64(424))); ; {
					_compr_26, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _361_v91 _dafny.Int
					_361_v91 = interface{}(_compr_26).(_dafny.Int)
					if ((_dafny.IntOfInt64(58)).Cmp(_361_v91) <= 0) && ((_361_v91).Cmp(_dafny.IntOfInt64(424)) < 0) {
						_coll26.Add((_361_v91).Minus(_dafny.IntOfInt64(335)), true)
					}
				}
				return _coll26.ToMap()
			}(), 13)
			(_nw56).ArraySet1(_256_v15, 14)
			(_nw56).ArraySet1(_257_v16, 15)
			(_nw56).ArraySet1(func() _dafny.Map {
				var _coll27 = _dafny.NewMapBuilder()
				_ = _coll27
				for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(230), _dafny.IntOfInt64(64))); ; {
					_compr_27, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _362_v92 _dafny.Int
					_362_v92 = interface{}(_compr_27).(_dafny.Int)
					if ((_dafny.IntOfInt64(230)).Cmp(_362_v92) <= 0) && ((_362_v92).Cmp(_dafny.IntOfInt64(64)) < 0) {
						_coll27.Add((_362_v92).Minus(_239_v0), _268_v24)
					}
				}
				return _coll27.ToMap()
			}(), 16)
			(_nw56).ArraySet1(_257_v16, 17)
			_358_v93 = _nw56
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_358_v93), 0))
			_ = _index63
			(_358_v93).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_253_v12).Cardinality()), !(_268_v24)), (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_358_v93), 0))
			_ = _index64
			(_358_v93).ArraySet1((func() _dafny.Map {
				if Companion_Default___.Fm1(_351_v82, _345_v76.F18(), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _259_globalState) {
					return _256_v15
				}
				return _257_v16
			})(), (_index64).Int())
			var _363_v94 D10
			_ = _363_v94
			_363_v94 = Companion_D10_.Create_DC25_(_250_v9)
			var _364_v95 _dafny.Array
			_ = _364_v95
			var _nwElement0_9 _dafny.Array = _250_v9
			_ = _nwElement0_9
			var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(9))
			_ = _nw57
			(_nw57).ArraySet1(_nwElement0_9, 0)
			(_nw57).ArraySet1(_250_v9, 1)
			(_nw57).ArraySet1(_250_v9, 2)
			(_nw57).ArraySet1((_363_v94).Dtor_cf33(), 3)
			(_nw57).ArraySet1(_250_v9, 4)
			(_nw57).ArraySet1(_250_v9, 5)
			(_nw57).ArraySet1(_250_v9, 6)
			(_nw57).ArraySet1((_254_v13).Select((Companion_Default___.SafeIndex(_352_v83, _dafny.IntOfUint32((_254_v13).Cardinality()))).Uint32()).(_dafny.Array), 7)
			(_nw57).ArraySet1((func() _dafny.Array {
				if _268_v24 {
					return _250_v9
				}
				return _250_v9
			})(), 8)
			_364_v95 = _nw57
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_364_v95), 0))
			_ = _index65
			(_364_v95).ArraySet1(_250_v9, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_364_v95), 0))
			_ = _index66
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
			_ = _index67
			var _rhs66 _dafny.Sequence = _240_v1
			_ = _rhs66
			var _rhs67 _dafny.Array = _250_v9
			_ = _rhs67
			var _rhs68 _dafny.Int = _269_v25
			_ = _rhs68
			var _lhs53 *GlobalState = _259_globalState
			_ = _lhs53
			var _lhs54 _dafny.Array = _364_v95
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_364_v95), 0))
			_ = _lhs55
			var _lhs56 _dafny.Array = _250_v9
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
			_ = _lhs57
			_lhs53.F1 = _rhs66
			(_lhs54).ArraySet1(_rhs67, (_lhs55).Int())
			(_lhs56).ArraySet1(_rhs68, (_lhs57).Int())
			_268_v24 = _351_v82
		} else {
			var _365_v96 _dafny.Array
			_ = _365_v96
			var _366_v97 bool
			_ = _366_v97
			var _367_v98 _dafny.Int
			_ = _367_v98
			var _out13 _dafny.Array
			_ = _out13
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			_out13, _out14, _out15 = Companion_Default___.M0(_343_cf3, (_269_v25).Cmp((Companion_Default___.Fm2(_259_globalState)).Cardinality()) == 0, (_348_v79).Cardinality(), _259_globalState)
			_365_v96 = _out13
			_366_v97 = _out14
			_367_v98 = _out15
			_268_v24 = false
			var _368_v99 T1
			_ = _368_v99
			var _nw58 *C3 = New_C3_()
			_ = _nw58
			_nw58.Ctor__(_345_v76.F18(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, _269_v25))
			_368_v99 = _nw58
			var _369_v100 *C2
			_ = _369_v100
			var _nw59 *C2 = New_C2_()
			_ = _nw59
			_nw59.Ctor__(_345_v76.F18(), _344_v75)
			_369_v100 = _nw59
			var _370_v101 _dafny.Map
			_ = _370_v101
			_370_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v99, _369_v100)
			var _371_v102 _dafny.MultiSet
			_ = _371_v102
			_371_v102 = _dafny.MultiSetOf((_370_v101).Merge(_370_v101))
			var _372_v103 _dafny.Sequence
			_ = _372_v103
			_372_v103 = _dafny.SeqOf(_370_v101)
			_371_v102 = (_dafny.MultiSetFromSeq(_372_v103)).Difference(_371_v102)
			_352_v83 = _239_v0
			_239_v0 = _367_v98
		}
		_243_v4 = _243_v4
	} else if _source4.Is_DC5() {
		var _373___mcc_h7 _dafny.Int = _source4.Get_().(D1_DC5).Cf6
		_ = _373___mcc_h7
		var _374_cf6 _dafny.Int = _373___mcc_h7
		_ = _374_cf6
		_249_v8 = (func() _dafny.Set {
			if (!(_268_v24)) && (_241_v2) {
				return (func() _dafny.Set {
					if _241_v2 {
						return Companion_Default___.Fm2(_259_globalState)
					}
					return _249_v8
				})()
			}
			return _249_v8
		})()
		var _375_v104 T0
		_ = _375_v104
		var _nw60 *C0 = New_C0_()
		_ = _nw60
		_nw60.Ctor__(_253_v12, _244_v5, _243_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vf"), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)))
		_375_v104 = _nw60
		_375_v104 = _375_v104
		var _376_v105 _dafny.Array
		_ = _376_v105
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
		_ = _nw61
		_376_v105 = _nw61
		_376_v105 = _376_v105
		var _377_v106 _dafny.MultiSet
		_ = _377_v106
		_377_v106 = _dafny.MultiSetOf((func() bool {
			if (_242_v3).Contains(false) {
				return (_242_v3).Get(false).(bool)
			}
			return _241_v2
		})(), !(_241_v2))
		var _378_v107 D1
		_ = _378_v107
		_378_v107 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus((_377_v106).Cardinality()), (_dafny.Zero).Minus(_269_v25), _241_v2)
		var _379_v108 _dafny.Sequence
		_ = _379_v108
		_379_v108 = _dafny.SeqOf(_249_v8)
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
		_ = _index68
		var _rhs69 _dafny.Int = (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)
		_ = _rhs69
		var _rhs70 D1 = Companion_D1_.Create_DC4_(_239_v0, _269_v25, true)
		_ = _rhs70
		var _rhs71 _dafny.Sequence = _379_v108
		_ = _rhs71
		var _rhs72 _dafny.Sequence = Companion_Default___.Fm27(_259_globalState)
		_ = _rhs72
		var _lhs58 _dafny.Array = _250_v9
		_ = _lhs58
		var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))
		_ = _lhs59
		var _lhs60 *GlobalState = _259_globalState
		_ = _lhs60
		(_lhs58).ArraySet1(_rhs69, (_lhs59).Int())
		_378_v107 = _rhs70
		_379_v108 = _rhs71
		_lhs60.F13 = _rhs72
	} else if _source4.Is_DC6() {
		var _380___mcc_h8 bool = _source4.Get_().(D1_DC6).Cf7
		_ = _380___mcc_h8
		var _381___mcc_h9 bool = _source4.Get_().(D1_DC6).Cf8
		_ = _381___mcc_h9
		var _382_cf8 bool = _381___mcc_h9
		_ = _382_cf8
		var _383_cf7 bool = _380___mcc_h8
		_ = _383_cf7
		var _384_v109 _dafny.Array
		_ = _384_v109
		var _385_v110 bool
		_ = _385_v110
		var _386_v111 _dafny.Int
		_ = _386_v111
		var _out16 _dafny.Array
		_ = _out16
		var _out17 bool
		_ = _out17
		var _out18 _dafny.Int
		_ = _out18
		_out16, _out17, _out18 = Companion_Default___.M0(_239_v0, (_383_cf7) && (_241_v2), _239_v0, _259_globalState)
		_384_v109 = _out16
		_385_v110 = _out17
		_386_v111 = _out18
		var _387_v112 _dafny.Map
		_ = _387_v112
		_387_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v25, _383_cf7)).Cardinality()))
		var _388_v113 _dafny.Map
		_ = _388_v113
		_388_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_269_v25), _387_v112)
		var _389_v114 *C1
		_ = _389_v114
		var _nw62 *C1 = New_C1_()
		_ = _nw62
		_nw62.Ctor__((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_269_v25, _386_v111)), true, _243_v4, (func() _dafny.Map {
			if (_388_v113).Contains(_239_v0) {
				return (_388_v113).Get(_239_v0).(_dafny.Map)
			}
			return _387_v112
		})())
		_389_v114 = _nw62
		var _390_v115 _dafny.Int
		_ = _390_v115
		var _out19 _dafny.Int
		_ = _out19
		_out19 = (_389_v114).M2(!(_382_cf8), _389_v114.F22, (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _386_v111, _259_globalState)
		_390_v115 = _out19
		var _391_v116 *C0
		_ = _391_v116
		var _nw63 *C0 = New_C0_()
		_ = _nw63
		_nw63.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(102))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_392_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_393_i10 _dafny.Int) _dafny.CodePoint {
				return _392_v4
			}
		})(_243_v4))), _244_v5, _243_v4, _387_v112)
		_391_v116 = _nw63
	} else {
		var _394___mcc_h10 bool = _source4.Get_().(D1_DC3).Cf2
		_ = _394___mcc_h10
		var _395_cf2 bool = _394___mcc_h10
		_ = _395_cf2
		var _396_v117 _dafny.Map
		_ = _396_v117
		_396_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))
		var _397_v118 _dafny.MultiSet
		_ = _397_v118
		_397_v118 = _dafny.MultiSetOf(_243_v4)
		var _398_v119 *C0
		_ = _398_v119
		var _nw64 *C0 = New_C0_()
		_ = _nw64
		_nw64.Ctor__(_dafny.Companion_Sequence_.Concatenate(_253_v12, _253_v12), _244_v5, _243_v4, (_396_v117).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, ((_397_v118).Update(_243_v4, Companion_Default___.Abs(_269_v25))).Cardinality())).Update(_dafny.UnicodeSeqOfUtf8Bytes("rtoijma"), (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))))
		_398_v119 = _nw64
		var _399_v120 D5
		_ = _399_v120
		_399_v120 = Companion_D5_.Create_DC14_(_268_v24, (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))
		var _400_v121 _dafny.Map
		_ = _400_v121
		_400_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_399_v120, _241_v2), _269_v25)
		var _401_v122 _dafny.Map
		_ = _401_v122
		_401_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_399_v120, _268_v24)
		_400_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_v122, (func() _dafny.Int {
			if (_325_v65).Contains((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)) {
				return (_325_v65).Get((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)).(_dafny.Int)
			}
			return _269_v25
		})())
		var _rhs73 _dafny.CodePoint = _243_v4
		_ = _rhs73
		var _rhs74 bool = (func() bool {
			if (_242_v3).Contains((_255_v14).Select((Companion_Default___.SafeIndex(_269_v25, _dafny.IntOfUint32((_255_v14).Cardinality()))).Uint32()).(bool)) {
				return (_242_v3).Get((_255_v14).Select((Companion_Default___.SafeIndex(_269_v25, _dafny.IntOfUint32((_255_v14).Cardinality()))).Uint32()).(bool)).(bool)
			}
			return _268_v24
		})()
		_ = _rhs74
		_243_v4 = _rhs73
		_395_cf2 = _rhs74
		var _402_v123 _dafny.Array
		_ = _402_v123
		var _nwElement0_10 _dafny.Sequence = _253_v12
		_ = _nwElement0_10
		var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(12))
		_ = _nw65
		(_nw65).ArraySet1(_nwElement0_10, 0)
		(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(416))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_403_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_404_i11 _dafny.Int) _dafny.CodePoint {
				return _403_v4
			}
		})(_243_v4))), 1)
		(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("p"), 2)
		(_nw65).ArraySet1(_253_v12, 3)
		(_nw65).ArraySet1(Companion_Default___.Fm8(_259_globalState), 4)
		(_nw65).ArraySet1((_398_v119).F20(), 5)
		(_nw65).ArraySet1((func() _dafny.Sequence {
			if !(_268_v24) {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}((func(_405_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_406_i12 _dafny.Int) _dafny.CodePoint {
						return _405_v4
					}
				})(_243_v4)))
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("f")
		})(), 6)
		(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}((func(_407_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_408_i13 _dafny.Int) _dafny.CodePoint {
				return _407_v4
			}
		})(_243_v4))), 7)
		(_nw65).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("l"), (_398_v119).F20()), 8)
		(_nw65).ArraySet1(_253_v12, 9)
		(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("art"), 10)
		(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fhfctskoe"), 11)
		_402_v123 = _nw65
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_402_v123), 0))
		_ = _index69
		(_402_v123).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_253_v12, _dafny.Companion_Sequence_.Update(_253_v12, (Companion_Default___.SafeIndex(_239_v0, _dafny.IntOfUint32((_253_v12).Cardinality()))).Uint32(), _243_v4)), (_index69).Int())
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_402_v123), 0))
		_ = _index70
		(_402_v123).ArraySet1(_253_v12, (_index70).Int())
	}
	var _409_i14 _dafny.Int
	_ = _409_i14
	_409_i14 = _dafny.Zero
	{
		for false {
			{
				if (_409_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_409_i14 = (_409_i14).Plus(_dafny.One)
				var _410_v124 *C0
				_ = _410_v124
				var _nw66 *C0 = New_C0_()
				_ = _nw66
				_nw66.Ctor__(_253_v12, _244_v5, _243_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, (_dafny.SetOf(_dafny.IntOfUint32((_255_v14).Cardinality()))).Cardinality()))
				_410_v124 = _nw66
				_410_v124 = _410_v124
				var _411_v125 _dafny.Map
				_ = _411_v125
				_411_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, _dafny.IntOfUint32((_253_v12).Cardinality()))
				var _412_v126 *C0
				_ = _412_v126
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__(_253_v12, (_410_v124).F21(), (func() _dafny.CodePoint {
					if _268_v24 {
						return _243_v4
					}
					return _243_v4
				})(), (_411_v125).Merge(_411_v125))
				_412_v126 = _nw67
				var _413_v127 _dafny.Array
				_ = _413_v127
				var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw68
				_413_v127 = _nw68
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_413_v127), 0))
				_ = _index71
				(_413_v127).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_268_v24), _255_v14), (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_413_v127), 0))
				_ = _index72
				(_413_v127).ArraySet1(_255_v14, (_index72).Int())
				var _414_v128 _dafny.Map
				_ = _414_v128
				_414_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_250_v9, _411_v125)
				var _415_v129 *C3
				_ = _415_v129
				var _nw69 *C3 = New_C3_()
				_ = _nw69
				_nw69.Ctor__(_243_v4, ((func() _dafny.Map {
					if (_414_v128).Contains(_250_v9) {
						return (_414_v128).Get(_250_v9).(_dafny.Map)
					}
					return _411_v125
				})()).Merge(_411_v125))
				_415_v129 = _nw69
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _416_v130 _dafny.Map
	_ = _416_v130
	_416_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v24, _239_v0)
	var _hi2 _dafny.Int = Companion_Default___.Fm0(_259_globalState)
	_ = _hi2
	for _417_i15 := (_416_v130).Cardinality(); _417_i15.Cmp(_hi2) < 0; _417_i15 = _417_i15.Plus(_dafny.One) {
		_241_v2 = (func() bool {
			if (_239_v0).Cmp(_417_i15) >= 0 {
				return !(_241_v2)
			}
			return _268_v24
		})()
		var _418_v131 _dafny.Map
		_ = _418_v131
		_418_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_i15, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_416_v130).Contains(_241_v2) {
				return (_416_v130).Get(_241_v2).(_dafny.Int)
			}
			return _417_i15
		})(), (_249_v8).Cardinality()))
		_418_v131 = _418_v131
		_242_v3 = _242_v3
		_250_v9 = _250_v9
	}
	_269_v25 = _239_v0
	if _268_v24 {
		var _419_v132 D3
		_ = _419_v132
		_419_v132 = Companion_D3_.Create_DC8_(Companion_Default___.Fm20(_268_v24, _259_globalState))
		var _420_v134 _dafny.Sequence
		_ = _420_v134
		_420_v134 = _dafny.SeqOf(_253_v12)
		var _421_v135 *C3
		_ = _421_v135
		var _nw70 *C3 = New_C3_()
		_ = _nw70
		_nw70.Ctor__((_419_v132).Dtor_cf10(), func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter29 := _dafny.Iterate((_420_v134).Elements()); ; {
				_compr_28, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _422_v133 _dafny.Sequence
				_422_v133 = interface{}(_compr_28).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_420_v134, _422_v133) {
					_coll28.Add(_422_v133, _269_v25)
				}
			}
			return _coll28.ToMap()
		}())
		_421_v135 = _nw70
		var _423_v136 _dafny.Map
		_ = _423_v136
		_423_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_421_v135, (func() bool {
			if _241_v2 {
				return _241_v2
			}
			return !((_421_v135).Fm3(_268_v24, _dafny.IntOfUint32((_253_v12).Cardinality()), _259_globalState))
		})())
		var _424_v137 _dafny.Sequence
		_ = _424_v137
		_424_v137 = _dafny.SeqOf(Companion_Default___.Fm28(_241_v2, (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int), _268_v24, _259_globalState))
		_423_v136 = (_423_v136).Update(_421_v135, (((_424_v137).Select((Companion_Default___.SafeIndex(_269_v25, _dafny.IntOfUint32((_424_v137).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()).Cmp(_239_v0) != 0)
		(_259_globalState).F6 = (_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)
		var _425_v138 _dafny.Map
		_ = _425_v138
		_425_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v12, Companion_Default___.Fm0(_259_globalState))
		var _426_v139 *C1
		_ = _426_v139
		var _nw71 *C1 = New_C1_()
		_ = _nw71
		_nw71.Ctor__(_269_v25, _241_v2, _243_v4, _425_v138)
		_426_v139 = _nw71
		var _427_v140 _dafny.Array
		_ = _427_v140
		var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
		_ = _nw72
		_427_v140 = _nw72
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_427_v140), 0))
		_ = _index73
		(_427_v140).ArraySet1(_250_v9, (_index73).Int())
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_427_v140), 0))
		_ = _index74
		(_427_v140).ArraySet1(_250_v9, (_index74).Int())
		_253_v12 = Companion_Default___.Fm8(_259_globalState)
	} else {
		(_259_globalState).F6 = _269_v25
		var _428_v141 _dafny.Array
		_ = _428_v141
		var _nwElement0_11 _dafny.Array = _250_v9
		_ = _nwElement0_11
		var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(10))
		_ = _nw73
		(_nw73).ArraySet1(_nwElement0_11, 0)
		(_nw73).ArraySet1(_250_v9, 1)
		(_nw73).ArraySet1(_250_v9, 2)
		(_nw73).ArraySet1(_250_v9, 3)
		(_nw73).ArraySet1(_250_v9, 4)
		(_nw73).ArraySet1(_250_v9, 5)
		(_nw73).ArraySet1(_250_v9, 6)
		(_nw73).ArraySet1(_250_v9, 7)
		(_nw73).ArraySet1(_250_v9, 8)
		(_nw73).ArraySet1(_250_v9, 9)
		_428_v141 = _nw73
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_428_v141), 0))
		_ = _index75
		(_428_v141).ArraySet1(_250_v9, (_index75).Int())
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_428_v141), 0))
		_ = _index76
		(_428_v141).ArraySet1(_250_v9, (_index76).Int())
		var _429_v142 D5
		_ = _429_v142
		_429_v142 = Companion_D5_.Create_DC14_(false, _dafny.IntOfInt64(51))
		var _rhs75 bool = (Companion_Default___.Fm29(_259_globalState)).Dtor_cf5()
		_ = _rhs75
		var _rhs76 _dafny.Int = (((_dafny.Zero).Minus((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int))).Minus(_dafny.IntOfUint32((_253_v12).Cardinality()))).Minus((_429_v142).Dtor_cf17())
		_ = _rhs76
		var _rhs77 bool = ((_250_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_250_v9), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus(_269_v25)) < 0
		_ = _rhs77
		_268_v24 = _rhs75
		_239_v0 = _rhs76
		_241_v2 = _rhs77
		var _430_v143 _dafny.MultiSet
		_ = _430_v143
		_430_v143 = _dafny.MultiSetOf(_268_v24, _268_v24, _241_v2)
		var _431_v144 _dafny.Array
		_ = _431_v144
		var _432_v145 bool
		_ = _432_v145
		var _433_v146 _dafny.Int
		_ = _433_v146
		var _out20 _dafny.Array
		_ = _out20
		var _out21 bool
		_ = _out21
		var _out22 _dafny.Int
		_ = _out22
		_out20, _out21, _out22 = Companion_Default___.M0((_430_v143).Cardinality(), _268_v24, (_dafny.Zero).Minus((_269_v25).Times(_dafny.IntOfUint32((Companion_Default___.Fm30(_241_v2, _259_globalState)).Cardinality()))), _259_globalState)
		_431_v144 = _out20
		_432_v145 = _out21
		_433_v146 = _out22
		_432_v145 = _268_v24
	}
	_dafny.Print(_239_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_240_v1, _dafny.SeqOf(_dafny.IntOfInt64(-140))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_241_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true).UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_243_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfInt64(-140))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v7).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_249_v8).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v11).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_253_v12.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_254_v13).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_255_v14, _dafny.SeqOf(false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(140), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(140), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_258_v17, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(140), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_259_globalState.F1, _dafny.SeqOf(_dafny.IntOfInt64(-140))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_259_globalState).F2(), _dafny.SeqOf(_dafny.IntOfInt64(-140))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_globalState).F3()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_globalState).F5()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfInt64(-140))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_259_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_259_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_globalState).F8()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F9().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_259_globalState).F10()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_259_globalState).F11()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_259_globalState.F13, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_259_globalState).F16(), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(140), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(140), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_265_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v23).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v23).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_268_v24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_269_v25)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v26).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_271_v27).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v28).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.Zero).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.Zero).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.One).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.One).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(D7)).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_323_v64).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(D7)).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_325_v65).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-976), _dafny.IntOfInt64(458)).UpdateUnsafe(_dafny.IntOfInt64(-592), _dafny.IntOfInt64(378))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_326_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_336_v73).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_337_v74).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_337_v74).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_409_i14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_416_v130).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(140))))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
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

type D0_DC2 struct {
	Cf1 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf1 D0) D0 {
	return D0{D0_DC2{Cf1}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() D0 {
	return _this.Get_().(D0_DC2).Cf1
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf1) + ")"
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
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf1.Equals(data2.Cf1)
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
	Cf3 _dafny.Int
	Cf4 _dafny.Int
	Cf5 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 bool) D1 {
	return D1{D1_DC4{Cf3, Cf4, Cf5}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf6 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf6 _dafny.Int) D1 {
	return D1{D1_DC5{Cf6}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf7 bool
	Cf8 bool
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf7 bool, Cf8 bool) D1 {
	return D1{D1_DC6{Cf7, Cf8}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC3 struct {
	Cf2 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 bool) D1 {
	return D1{D1_DC3{Cf2}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC6).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC6).Cf8
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
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

type D2_DC7 struct {
	Cf9 _dafny.Sequence
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf9 _dafny.Sequence) D2 {
	return D2{D2_DC7{Cf9}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf11 bool
	Cf12 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 bool, Cf12 _dafny.Int) D3 {
	return D3{D3_DC9{Cf11, Cf12}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf13 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf13 bool) D3 {
	return D3{D3_DC10{Cf13}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf10 _dafny.CodePoint
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf10 _dafny.CodePoint) D3 {
	return D3{D3_DC8{Cf10}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(false, _dafny.Zero)
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D3_DC8).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf13 == data2.Cf13
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf10 == data2.Cf10
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
	Cf14 T0
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf14 T0) D4 {
	return D4{D4_DC11{Cf14}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() T0 {
	return (T0)(nil)
}

func (_this D4) Dtor_cf14() T0 {
	return _this.Get_().(D4_DC11).Cf14
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf14, data2.Cf14)
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
	Cf16 bool
	Cf17 _dafny.Int
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf16 bool, Cf17 _dafny.Int) D5 {
	return D5{D5_DC14{Cf16, Cf17}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf15 _dafny.MultiSet
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf15 _dafny.MultiSet) D5 {
	return D5{D5_DC12{Cf15}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_()
}

func (_this D5) Dtor_cf16() bool {
	return _this.Get_().(D5_DC14).Cf16
}

func (_this D5) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf17
}

func (_this D5) Dtor_cf15() _dafny.MultiSet {
	return _this.Get_().(D5_DC12).Cf15
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
			return "D5.DC14" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf19 bool
	Cf20 _dafny.Array
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf19 bool, Cf20 _dafny.Array) D6 {
	return D6{D6_DC16{Cf19, Cf20}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf18 _dafny.Map
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf18 _dafny.Map) D6 {
	return D6{D6_DC15{Cf18}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC17 struct {
	Cf21 D6
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf21 D6) D6 {
	return D6{D6_DC17{Cf21}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D6) Dtor_cf19() bool {
	return _this.Get_().(D6_DC16).Cf19
}

func (_this D6) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D6_DC16).Cf20
}

func (_this D6) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D6_DC15).Cf18
}

func (_this D6) Dtor_cf21() D6 {
	return _this.Get_().(D6_DC17).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
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

type D7_DC19 struct {
	Cf23 bool
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf23 bool) D7 {
	return D7{D7_DC19{Cf23}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf24 bool
	Cf25 bool
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf24 bool, Cf25 bool) D7 {
	return D7{D7_DC20{Cf24, Cf25}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC18 struct {
	Cf22 _dafny.Set
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf22 _dafny.Set) D7 {
	return D7{D7_DC18{Cf22}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(false)
}

func (_this D7) Dtor_cf23() bool {
	return _this.Get_().(D7_DC19).Cf23
}

func (_this D7) Dtor_cf24() bool {
	return _this.Get_().(D7_DC20).Cf24
}

func (_this D7) Dtor_cf25() bool {
	return _this.Get_().(D7_DC20).Cf25
}

func (_this D7) Dtor_cf22() _dafny.Set {
	return _this.Get_().(D7_DC18).Cf22
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D8_DC21 struct {
	Cf26 _dafny.Sequence
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf26 _dafny.Sequence) D8 {
	return D8{D8_DC21{Cf26}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D8_DC21).Cf26
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf26.Equals(data2.Cf26)
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
	Cf28 _dafny.Sequence
	Cf29 _dafny.Int
	Cf30 bool
	Cf31 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf28 _dafny.Sequence, Cf29 _dafny.Int, Cf30 bool, Cf31 _dafny.Int) D9 {
	return D9{D9_DC23{Cf28, Cf29, Cf30, Cf31}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC24 struct {
	Cf32 bool
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf32 bool) D9 {
	return D9{D9_DC24{Cf32}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC22 struct {
	Cf27 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf27 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf27}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(_dafny.EmptySeq, _dafny.Zero, false, _dafny.Zero)
}

func (_this D9) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D9_DC23).Cf28
}

func (_this D9) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf29
}

func (_this D9) Dtor_cf30() bool {
	return _this.Get_().(D9_DC23).Cf30
}

func (_this D9) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf31
}

func (_this D9) Dtor_cf32() bool {
	return _this.Get_().(D9_DC24).Cf32
}

func (_this D9) Dtor_cf27() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf27
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + data.Cf27.VerbatimString(true) + ")"
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
			return ok && data1.Cf28.Equals(data2.Cf28) && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf32 == data2.Cf32
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	Cf34 _dafny.Int
	Cf35 _dafny.Array
	Cf36 bool
	Cf37 bool
	Cf38 _dafny.Int
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf34 _dafny.Int, Cf35 _dafny.Array, Cf36 bool, Cf37 bool, Cf38 _dafny.Int) D10 {
	return D10{D10_DC26{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC25 struct {
	Cf33 _dafny.Array
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf33 _dafny.Array) D10 {
	return D10{D10_DC25{Cf33}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC26_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, false, _dafny.Zero)
}

func (_this D10) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D10_DC26).Cf34
}

func (_this D10) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D10_DC26).Cf35
}

func (_this D10) Dtor_cf36() bool {
	return _this.Get_().(D10_DC26).Cf36
}

func (_this D10) Dtor_cf37() bool {
	return _this.Get_().(D10_DC26).Cf37
}

func (_this D10) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D10_DC26).Cf38
}

func (_this D10) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D10_DC25).Cf33
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf33) + ")"
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
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf33 == data2.Cf33
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
	Cf40 *C2
	Cf41 _dafny.Int
	Cf42 bool
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf40 *C2, Cf41 _dafny.Int, Cf42 bool) D11 {
	return D11{D11_DC28{Cf40, Cf41, Cf42}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC27 struct {
	Cf39 _dafny.Map
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf39 _dafny.Map) D11 {
	return D11{D11_DC27{Cf39}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC28_((*C2)(nil), _dafny.Zero, false)
}

func (_this D11) Dtor_cf40() *C2 {
	return _this.Get_().(D11_DC28).Cf40
}

func (_this D11) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D11_DC28).Cf41
}

func (_this D11) Dtor_cf42() bool {
	return _this.Get_().(D11_DC28).Cf42
}

func (_this D11) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D11_DC27).Cf39
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42 == data2.Cf42
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf44 bool
	Cf45 bool
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf44 bool, Cf45 bool) D12 {
	return D12{D12_DC30{Cf44, Cf45}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

type D12_DC31 struct {
	Cf46 bool
	Cf47 _dafny.Sequence
	Cf48 _dafny.Int
	Cf49 bool
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf46 bool, Cf47 _dafny.Sequence, Cf48 _dafny.Int, Cf49 bool) D12 {
	return D12{D12_DC31{Cf46, Cf47, Cf48, Cf49}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC29 struct {
	Cf43 _dafny.Set
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf43 _dafny.Set) D12 {
	return D12{D12_DC29{Cf43}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC32 struct {
	Cf50 D12
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf50 D12) D12 {
	return D12{D12_DC32{Cf50}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC30_(false, false)
}

func (_this D12) Dtor_cf44() bool {
	return _this.Get_().(D12_DC30).Cf44
}

func (_this D12) Dtor_cf45() bool {
	return _this.Get_().(D12_DC30).Cf45
}

func (_this D12) Dtor_cf46() bool {
	return _this.Get_().(D12_DC31).Cf46
}

func (_this D12) Dtor_cf47() _dafny.Sequence {
	return _this.Get_().(D12_DC31).Cf47
}

func (_this D12) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D12_DC31).Cf48
}

func (_this D12) Dtor_cf49() bool {
	return _this.Get_().(D12_DC31).Cf49
}

func (_this D12) Dtor_cf43() _dafny.Set {
	return _this.Get_().(D12_DC29).Cf43
}

func (_this D12) Dtor_cf50() D12 {
	return _this.Get_().(D12_DC32).Cf50
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf46) + ", " + data.Cf47.VerbatimString(true) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf50) + ")"
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
			return ok && data1.Cf44 == data2.Cf44 && data1.Cf45 == data2.Cf45
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Equals(data2.Cf47) && data1.Cf48.Cmp(data2.Cf48) == 0 && data1.Cf49 == data2.Cf49
		}
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D13_DC33 struct {
	Cf51 _dafny.MultiSet
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf51 _dafny.MultiSet) D13 {
	return D13{D13_DC33{Cf51}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D13) Dtor_cf51() _dafny.MultiSet {
	return _this.Get_().(D13_DC33).Cf51
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
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

// Definition of trait T0
type T0 interface {
	String() string
	F18() _dafny.CodePoint
	F18_set_(value _dafny.CodePoint)
	F19() _dafny.Map
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
	F18() _dafny.CodePoint
	F18_set_(value _dafny.CodePoint)
	F19() _dafny.Map
	Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool
	Fm4(p0 bool, p1 _dafny.CodePoint, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Int
	M1(p0 bool, globalState *GlobalState) (D0, _dafny.Int, bool, bool)
	M2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int
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
	F1   _dafny.Sequence
	F6   _dafny.Int
	F13  _dafny.Sequence
	_f0  _dafny.Int
	_f2  _dafny.Sequence
	_f3  _dafny.Map
	_f4  _dafny.Int
	_f5  _dafny.Map
	_f7  _dafny.Array
	_f8  _dafny.Map
	_f9  _dafny.Sequence
	_f10 _dafny.Map
	_f11 _dafny.Sequence
	_f12 bool
	_f14 bool
	_f15 _dafny.Int
	_f16 _dafny.Sequence
	_f17 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.EmptySeq
	_this.F6 = _dafny.Zero
	_this.F13 = _dafny.EmptySeq
	_this._f0 = _dafny.Zero
	_this._f2 = _dafny.EmptySeq
	_this._f3 = _dafny.EmptyMap
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.EmptyMap
	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptyMap
	_this._f9 = _dafny.EmptySeq
	_this._f10 = _dafny.EmptyMap
	_this._f11 = _dafny.EmptySeq
	_this._f12 = false
	_this._f14 = false
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.EmptySeq
	_this._f17 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 _dafny.Sequence, f3 _dafny.Map, f4 _dafny.Int, f5 _dafny.Map, f6 _dafny.Int, f7 _dafny.Array, f8 _dafny.Map, f9 _dafny.Sequence, f10 _dafny.Map, f11 _dafny.Sequence, f12 bool, f13 _dafny.Sequence, f14 bool, f15 _dafny.Int, f16 _dafny.Sequence, f17 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Sequence {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Map {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Map {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F7() _dafny.Array {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Map {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Sequence {
	{
		return _this._f9
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
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Sequence {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f18 _dafny.CodePoint
	_f19 _dafny.Map
	_f20 _dafny.Sequence
	_f21 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f18 = _dafny.CodePoint('D')
	_this._f19 = _dafny.EmptyMap
	_this._f20 = _dafny.EmptySeq
	_this._f21 = _dafny.EmptyMap
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

func (_this *C0) F18() _dafny.CodePoint {
	return _this._f18
}
func (_this *C0) F18_set_(value _dafny.CodePoint) {
	_this._f18 = value
}
func (_this *C0) F19() _dafny.Map {
	return _this._f19
}
func (_this *C0) Ctor__(f20 _dafny.Sequence, f21 _dafny.Map, f18 _dafny.CodePoint, f19 _dafny.Map) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C0) Fm6(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) bool {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()).Minus(_dafny.IntOfUint32(((_this).F20()).Cardinality()))).Cmp((_dafny.IntOfInt64(-574)).Plus(_dafny.IntOfInt64(924))) != 0
	}
}
func (_this *C0) F20() _dafny.Sequence {
	{
		return _this._f20
	}
}
func (_this *C0) F21() _dafny.Map {
	{
		return _this._f21
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f18 _dafny.CodePoint
	_f19 _dafny.Map
	F22  _dafny.Int
	_f23 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f18 = _dafny.CodePoint('D')
	_this._f19 = _dafny.EmptyMap
	_this.F22 = _dafny.Zero
	_this._f23 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F18() _dafny.CodePoint {
	return _this._f18
}
func (_this *C1) F18_set_(value _dafny.CodePoint) {
	_this._f18 = value
}
func (_this *C1) F19() _dafny.Map {
	return _this._f19
}
func (_this *C1) Ctor__(f22 _dafny.Int, f23 bool, f18 _dafny.CodePoint, f19 _dafny.Map) {
	{
		(_this).F22 = f22
		(_this)._f23 = f23
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C1) Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F23()
	}
}
func (_this *C1) Fm4(p0 bool, p1 _dafny.CodePoint, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(390)).Times((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if !((_this).F23()) {
				return _dafny.IntOfInt64(851)
			}
			return _dafny.IntOfInt64(366)
		})())))
	}
}
func (_this *C1) Fm12(globalState *GlobalState) bool {
	{
		return (_this).F23()
	}
}
func (_this *C1) Fm13(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		if !(!((_this).F23())) {
			return _this.F18()
		} else {
			return _dafny.CodePoint('b')
		}
	}
}
func (_this *C1) M1(p0 bool, globalState *GlobalState) (D0, _dafny.Int, bool, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _434_v0 _dafny.Map
		_ = _434_v0
		_434_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _this.F22)
		var _435_v1 _dafny.Set
		_ = _435_v1
		_435_v1 = _dafny.SetOf(_434_v0, _434_v0)
		r3 = ((_dafny.SetOf(_434_v0)).Union(_435_v1)).IsSubsetOf(_435_v1)
		(globalState).F6 = (_this.F22).Plus(_this.F22)
		var _436_v2 _dafny.MultiSet
		_ = _436_v2
		_436_v2 = _dafny.MultiSetOf(_this.F22, _this.F22, _this.F22)
		var _437_i0 _dafny.Int
		_ = _437_i0
		_437_i0 = _dafny.Zero
		{
			for ((_436_v2).Update(_this.F22, Companion_Default___.Abs(_this.F22))).IsSubsetOf((func() _dafny.MultiSet {
				if p0 {
					return _436_v2
				}
				return _436_v2
			})()) {
				{
					if (_437_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_437_i0 = (_437_i0).Plus(_dafny.One)
					var _438_v3 D3
					_ = _438_v3
					_438_v3 = Companion_D3_.Create_DC10_(p0)
					var _source5 D3 = _438_v3
					_ = _source5
					if _source5.Is_DC9() {
						var _439___mcc_h0 bool = _source5.Get_().(D3_DC9).Cf11
						_ = _439___mcc_h0
						var _440___mcc_h1 _dafny.Int = _source5.Get_().(D3_DC9).Cf12
						_ = _440___mcc_h1
						var _441_cf12 _dafny.Int = _440___mcc_h1
						_ = _441_cf12
						var _442_cf11 bool = _439___mcc_h0
						_ = _442_cf11
						var _443_v4 _dafny.Sequence
						_ = _443_v4
						_443_v4 = _dafny.SeqOf(!(_442_cf11))
						var _444_v5 _dafny.MultiSet
						_ = _444_v5
						_444_v5 = _dafny.MultiSetOf((_this).F23(), (_this).F23(), p0)
						var _445_v6 _dafny.Map
						_ = _445_v6
						_445_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_cf12, (_444_v5).Cardinality())
						var _446_v7 D3
						_ = _446_v7
						_446_v7 = Companion_D3_.Create_DC8_(_this.F18())
						var _447_v8 _dafny.Sequence
						_ = _447_v8
						_447_v8 = _dafny.SeqOf(_this.F22)
						var _448_v9 _dafny.Set
						_ = _448_v9
						_448_v9 = _dafny.SetOf(_442_cf11)
						var _449_v10 _dafny.Set
						_ = _449_v10
						_449_v10 = _dafny.SetOf(_441_cf12)
						var _450_v11 _dafny.Map
						_ = _450_v11
						_450_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_442_cf11, _447_v8)
						var _451_v12 _dafny.Array
						_ = _451_v12
						var _nwElement0_12 bool = p0
						_ = _nwElement0_12
						var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(21))
						_ = _nw74
						(_nw74).ArraySet1(_nwElement0_12, 0)
						(_nw74).ArraySet1(Companion_Default___.Fm1(p0, _this.F18(), _this.F22, globalState), 1)
						(_nw74).ArraySet1(_442_cf11, 2)
						(_nw74).ArraySet1(false, 3)
						(_nw74).ArraySet1(_442_cf11, 4)
						(_nw74).ArraySet1(!((_this).F23()), 5)
						(_nw74).ArraySet1((_this).F23(), 6)
						(_nw74).ArraySet1(Companion_Default___.Fm1((_this).F23(), _this.F18(), _this.F22, globalState), 7)
						(_nw74).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_443_v4, _dafny.SeqOf((_this).F23(), p0, p0, _442_cf11)), 8)
						(_nw74).ArraySet1((_441_cf12).Cmp((_445_v6).Cardinality()) <= 0, 9)
						(_nw74).ArraySet1((_this.F22).Cmp(_this.F22) != 0, 10)
						(_nw74).ArraySet1(_442_cf11, 11)
						(_nw74).ArraySet1(Companion_Default___.Fm1((_this).F23(), (_446_v7).Dtor_cf10(), (_447_v8).Select((Companion_Default___.SafeIndex(_441_cf12, _dafny.IntOfUint32((_447_v8).Cardinality()))).Uint32()).(_dafny.Int), globalState), 12)
						(_nw74).ArraySet1((Companion_Default___.Fm2(globalState)).IsSubsetOf(_448_v9), 13)
						(_nw74).ArraySet1(!(false), 14)
						(_nw74).ArraySet1(p0, 15)
						(_nw74).ArraySet1(_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
							if (_450_v11).Contains((_this).F23()) {
								return (_450_v11).Get((_this).F23()).(_dafny.Sequence)
							}
							return _447_v8
						})(), (_449_v10).Cardinality()), 16)
						(_nw74).ArraySet1((_441_cf12).Cmp(_441_cf12) == 0, 17)
						(_nw74).ArraySet1((_this).F23(), 18)
						(_nw74).ArraySet1(_442_cf11, 19)
						(_nw74).ArraySet1(Companion_Default___.Fm1(false, _this.F18(), _dafny.IntOfInt64(-279), globalState), 20)
						_451_v12 = _nw74
						_451_v12 = _451_v12
						var _452_v13 D0
						_ = _452_v13
						_452_v13 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_())
						var _453_v14 D0
						_ = _453_v14
						_453_v14 = Companion_D0_.Create_DC2_(_452_v13)
						_453_v14 = _453_v14
						var _454_v15 _dafny.Sequence
						_ = _454_v15
						_454_v15 = _dafny.UnicodeSeqOfUtf8Bytes("hkdt")
						var _455_v16 _dafny.Set
						_ = _455_v16
						_455_v16 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ylwbg"), _454_v15)
						_455_v16 = (_455_v16).Union(_dafny.SetOf(_454_v15, _454_v15, _454_v15))
						var _456_v17 _dafny.Map
						_ = _456_v17
						_456_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), _this.F22)
						var _457_v19 _dafny.Sequence
						_ = _457_v19
						_457_v19 = _dafny.SeqOf(_454_v15, _454_v15, _454_v15)
						var _458_v20 _dafny.Sequence
						_ = _458_v20
						_458_v20 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(170))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg30 _dafny.Int) interface{} {
								return coer30(arg30)
							}
						}(func(_459_i1 _dafny.Int) _dafny.CodePoint {
							return _this.F18()
						})), _441_cf12), (_this).F19(), func() _dafny.Map {
							var _coll29 = _dafny.NewMapBuilder()
							_ = _coll29
							for _iter30 := _dafny.Iterate((_457_v19).Elements()); ; {
								_compr_29, _ok30 := _iter30()
								if !_ok30 {
									break
								}
								var _460_v18 _dafny.Sequence
								_460_v18 = interface{}(_compr_29).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_457_v19, _460_v18) {
									_coll29.Add(_460_v18, _this.F22)
								}
							}
							return _coll29.ToMap()
						}(), (_this).F19())
						var _461_v21 *C0
						_ = _461_v21
						var _nw75 *C0 = New_C0_()
						_ = _nw75
						_nw75.Ctor__(_454_v15, _456_v17, _dafny.CodePoint('x'), (_458_v20).Select((Companion_Default___.SafeIndex(_this.F22, _dafny.IntOfUint32((_458_v20).Cardinality()))).Uint32()).(_dafny.Map))
						_461_v21 = _nw75
					} else if _source5.Is_DC10() {
						var _462___mcc_h2 bool = _source5.Get_().(D3_DC10).Cf13
						_ = _462___mcc_h2
						var _463_cf13 bool = _462___mcc_h2
						_ = _463_cf13
						_463_cf13 = ((Companion_D3_.Create_DC9_(_463_cf13, Companion_Default___.Fm0(globalState))).Dtor_cf11()) == (_463_cf13)
						var _464_v22 _dafny.Array
						_ = _464_v22
						var _nw76 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
						_ = _nw76
						_464_v22 = _nw76
						var _465_v23 _dafny.Map
						_ = _465_v23
						_465_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v22, _this.F22)
						var _466_v24 _dafny.MultiSet
						_ = _466_v24
						_466_v24 = _dafny.MultiSetOf((_this).F23())
						var _467_v25 _dafny.Set
						_ = _467_v25
						_467_v25 = _dafny.SetOf(false)
						var _468_v26 _dafny.Array
						_ = _468_v26
						var _nwElement0_13 _dafny.Int = (_this.F22).Plus(_this.F22)
						_ = _nwElement0_13
						var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(8))
						_ = _nw77
						(_nw77).ArraySet1(_nwElement0_13, 0)
						(_nw77).ArraySet1(((func() _dafny.Int {
							if (_465_v23).Contains(_464_v22) {
								return (_465_v23).Get(_464_v22).(_dafny.Int)
							}
							return (_466_v24).Cardinality()
						})()).Times(_this.F22), 1)
						(_nw77).ArraySet1(_this.F22, 2)
						(_nw77).ArraySet1(_this.F22, 3)
						(_nw77).ArraySet1(_this.F22, 4)
						(_nw77).ArraySet1((_467_v25).Cardinality(), 5)
						(_nw77).ArraySet1((_dafny.Zero).Minus(_this.F22), 6)
						(_nw77).ArraySet1((_dafny.IntOfInt64(251)).Minus(_this.F22), 7)
						_468_v26 = _nw77
						var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_468_v26), 0))
						_ = _index77
						(_468_v26).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F22, (Companion_Default___.Fm14(Companion_Default___.Fm0(globalState), globalState)).Cardinality()), (_index77).Int())
						var _469_v27 _dafny.Sequence
						_ = _469_v27
						_469_v27 = _dafny.SeqOf(Companion_Default___.Fm0(globalState))
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_468_v26), 0))
						_ = _index78
						(_468_v26).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.MultiSetFromSeq(_469_v27)).Cardinality(), _this.F22), (_index78).Int())
						var _470_v28 _dafny.Array
						_ = _470_v28
						var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
						_ = _nw78
						_470_v28 = _nw78
						var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_470_v28), 0))
						_ = _index79
						(_470_v28).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(132))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}(func(_471_i2 _dafny.Int) _dafny.CodePoint {
							return _this.F18()
						})), (_index79).Int())
						var _472_v29 _dafny.Sequence
						_ = _472_v29
						_472_v29 = _dafny.UnicodeSeqOfUtf8Bytes("vphuxrsd")
						var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_470_v28), 0))
						_ = _index80
						(_470_v28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_472_v29, _dafny.UnicodeSeqOfUtf8Bytes("f")), (_index80).Int())
						var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_468_v26), 0))
						_ = _index81
						(_468_v26).ArraySet1((_dafny.IntOfInt64(249)).Times((_dafny.Zero).Minus(_this.F22)), (_index81).Int())
					} else {
						var _473___mcc_h3 _dafny.CodePoint = _source5.Get_().(D3_DC8).Cf10
						_ = _473___mcc_h3
						var _474_cf10 _dafny.CodePoint = _473___mcc_h3
						_ = _474_cf10
						var _475_v30 _dafny.MultiSet
						_ = _475_v30
						_475_v30 = _dafny.MultiSetOf((_this).F23(), p0, p0, (_this).F23(), (_this).F23())
						var _476_v31 D5
						_ = _476_v31
						_476_v31 = Companion_D5_.Create_DC12_(_475_v30)
						var _477_v32 _dafny.Map
						_ = _477_v32
						_477_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, (_476_v31).Dtor_cf15())
						var _478_v33 _dafny.Sequence
						_ = _478_v33
						_478_v33 = _dafny.SeqOf(p0)
						var _479_v34 _dafny.Array
						_ = _479_v34
						var _nwElement0_14 _dafny.MultiSet = _475_v30
						_ = _nwElement0_14
						var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(17))
						_ = _nw79
						(_nw79).ArraySet1(_nwElement0_14, 0)
						(_nw79).ArraySet1(_475_v30, 1)
						(_nw79).ArraySet1((func() _dafny.MultiSet {
							if (_477_v32).Contains((_dafny.MultiSetOf(p0)).Cardinality()) {
								return (_477_v32).Get((_dafny.MultiSetOf(p0)).Cardinality()).(_dafny.MultiSet)
							}
							return _dafny.MultiSetOf((_this).F23())
						})(), 2)
						(_nw79).ArraySet1(_475_v30, 3)
						(_nw79).ArraySet1(_475_v30, 4)
						(_nw79).ArraySet1(_475_v30, 5)
						(_nw79).ArraySet1(_475_v30, 6)
						(_nw79).ArraySet1((_475_v30).Update((_this).F23(), Companion_Default___.Abs(_this.F22)), 7)
						(_nw79).ArraySet1(_dafny.MultiSetFromSeq(_478_v33), 8)
						(_nw79).ArraySet1((_475_v30).Intersection(_475_v30), 9)
						(_nw79).ArraySet1(_475_v30, 10)
						(_nw79).ArraySet1(_475_v30, 11)
						(_nw79).ArraySet1(_475_v30, 12)
						(_nw79).ArraySet1(_475_v30, 13)
						(_nw79).ArraySet1(_475_v30, 14)
						(_nw79).ArraySet1(_475_v30, 15)
						(_nw79).ArraySet1((_475_v30).Difference(_dafny.MultiSetFromSeq(_478_v33)), 16)
						_479_v34 = _nw79
						var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_479_v34), 0))
						_ = _index82
						(_479_v34).ArraySet1((_475_v30).Union(_475_v30), (_index82).Int())
						var _480_v35 _dafny.Sequence
						_ = _480_v35
						_480_v35 = _dafny.UnicodeSeqOfUtf8Bytes("hoka")
						var _481_v36 _dafny.MultiSet
						_ = _481_v36
						_481_v36 = _dafny.MultiSetOf((func() _dafny.Sequence {
							if p0 {
								return _480_v35
							}
							return _480_v35
						})(), _dafny.Companion_Sequence_.Concatenate(_480_v35, _480_v35), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(874))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg32 _dafny.Int) interface{} {
								return coer32(arg32)
							}
						}(func(_482_i3 _dafny.Int) _dafny.CodePoint {
							return _this.F18()
						})), _480_v35, _dafny.Companion_Sequence_.Update(_480_v35, (Companion_Default___.SafeIndex(_this.F22, _dafny.IntOfUint32((_480_v35).Cardinality()))).Uint32(), _474_cf10))
						var _483_v37 _dafny.Map
						_ = _483_v37
						_483_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, (_475_v30).Cardinality())
						var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_479_v34), 0))
						_ = _index83
						var _rhs78 bool = p0
						_ = _rhs78
						var _rhs79 _dafny.MultiSet = (_475_v30).Union(_475_v30)
						_ = _rhs79
						var _rhs80 _dafny.MultiSet = Companion_Default___.Fm15(_this.F18(), (_dafny.Zero).Minus(_this.F22), _438_v3, globalState)
						_ = _rhs80
						var _rhs81 _dafny.Int = (_dafny.Zero).Minus(((_this.F22).Minus(_this.F22)).Times((func() _dafny.Int {
							if (_483_v37).Contains((_434_v0).Cardinality()) {
								return (_483_v37).Get((_434_v0).Cardinality()).(_dafny.Int)
							}
							return _this.F22
						})()))
						_ = _rhs81
						var _lhs61 _dafny.Array = _479_v34
						_ = _lhs61
						var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_479_v34), 0))
						_ = _lhs62
						var _lhs63 *GlobalState = globalState
						_ = _lhs63
						r3 = _rhs78
						(_lhs61).ArraySet1(_rhs79, (_lhs62).Int())
						_481_v36 = _rhs80
						_lhs63.F6 = _rhs81
						var _484_v38 _dafny.Map
						_ = _484_v38
						_484_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F23())
						var _485_v39 _dafny.Set
						_ = _485_v39
						_485_v39 = _dafny.SetOf(_this.F22, _this.F22, _dafny.IntOfInt64(295))
						var _486_v40 _dafny.Sequence
						_ = _486_v40
						_486_v40 = _dafny.SeqOf(_485_v39)
						_484_v38 = (_484_v38).Update((_485_v39).IsDisjointFrom((_486_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-737), _dafny.IntOfUint32((_486_v40).Cardinality()))).Uint32()).(_dafny.Set)), p0)
						var _487_v41 _dafny.Array
						_ = _487_v41
						var _len0_7 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_7
						var _nw80 _dafny.Array
						_ = _nw80
						if _len0_7.Cmp(_dafny.Zero) == 0 {
							_nw80 = _dafny.NewArray(_len0_7)
						} else {
							var _init7 func(_dafny.Int) _dafny.Int = func(_488_i4 _dafny.Int) _dafny.Int {
								return (_488_i4).Times(_this.F22)
							}
							_ = _init7
							var _element0_7 = _init7(_dafny.Zero)
							_ = _element0_7
							_nw80 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
							(_nw80).ArraySet1(_element0_7, 0)
							var _nativeLen0_7 = (_len0_7).Int()
							_ = _nativeLen0_7
							for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
								(_nw80).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
							}
						}
						_487_v41 = _nw80
						var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_487_v41), 0))
						_ = _index84
						(_487_v41).ArraySet1(_this.F22, (_index84).Int())
						var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_487_v41), 0))
						_ = _index85
						(_487_v41).ArraySet1((_dafny.Zero).Minus(_this.F22), (_index85).Int())
						var _489_v42 D5
						_ = _489_v42
						_489_v42 = Companion_D5_.Create_DC14_((func() bool {
							if p0 {
								return p0
							}
							return (_this).F23()
						})(), (_487_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_487_v41), 0))).Int()).(_dafny.Int))
						var _490_v43 _dafny.Sequence
						_ = _490_v43
						_490_v43 = _dafny.SeqOf(_dafny.IntOfUint32((_480_v35).Cardinality()))
						var _pat_let_tv4 = _487_v41
						_ = _pat_let_tv4
						var _pat_let_tv5 = _487_v41
						_ = _pat_let_tv5
						var _pat_let_tv6 = _487_v41
						_ = _pat_let_tv6
						var _pat_let_tv7 = _487_v41
						_ = _pat_let_tv7
						var _pat_let_tv8 = _490_v43
						_ = _pat_let_tv8
						var _pat_let_tv9 = _490_v43
						_ = _pat_let_tv9
						_489_v42 = func(_pat_let6_0 D5) D5 {
							return func(_495_dt__update__tmp_h2 D5) D5 {
								return func(_pat_let11_0 _dafny.Int) D5 {
									return func(_496_dt__update_hcf17_h2 _dafny.Int) D5 {
										return Companion_D5_.Create_DC14_((_495_dt__update__tmp_h2).Dtor_cf16(), _496_dt__update_hcf17_h2)
									}(_pat_let11_0)
								}(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv8, _pat_let_tv9)).Cardinality()))
							}(_pat_let6_0)
						}(func(_pat_let7_0 D5) D5 {
							return func(_493_dt__update__tmp_h1 D5) D5 {
								return func(_pat_let10_0 _dafny.Int) D5 {
									return func(_494_dt__update_hcf17_h1 _dafny.Int) D5 {
										return Companion_D5_.Create_DC14_((_493_dt__update__tmp_h1).Dtor_cf16(), _494_dt__update_hcf17_h1)
									}(_pat_let10_0)
								}((_pat_let_tv7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_pat_let_tv6), 0))).Int()).(_dafny.Int))
							}(_pat_let7_0)
						}(func(_pat_let8_0 D5) D5 {
							return func(_491_dt__update__tmp_h0 D5) D5 {
								return func(_pat_let9_0 _dafny.Int) D5 {
									return func(_492_dt__update_hcf17_h0 _dafny.Int) D5 {
										return Companion_D5_.Create_DC14_((_491_dt__update__tmp_h0).Dtor_cf16(), _492_dt__update_hcf17_h0)
									}(_pat_let9_0)
								}((_pat_let_tv5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_pat_let_tv4), 0))).Int()).(_dafny.Int))
							}(_pat_let8_0)
						}(_489_v42)))
					}
					var _497_v44 _dafny.Sequence
					_ = _497_v44
					_497_v44 = _dafny.SeqOf((_this).F23())
					var _498_v45 _dafny.Map
					_ = _498_v45
					_498_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, _this.F22)
					var _499_v46 _dafny.Sequence
					_ = _499_v46
					_499_v46 = _dafny.SeqOf(_498_v45, _498_v45)
					var _500_v47 _dafny.Sequence
					_ = _500_v47
					_500_v47 = _dafny.UnicodeSeqOfUtf8Bytes("gr")
					var _501_v48 _dafny.Sequence
					_ = _501_v48
					_501_v48 = _497_v44
					var _502_v49 _dafny.Array
					_ = _502_v49
					var _nwElement0_15 bool = false
					_ = _nwElement0_15
					var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(15))
					_ = _nw81
					(_nw81).ArraySet1(_nwElement0_15, 0)
					(_nw81).ArraySet1((_this).F23(), 1)
					(_nw81).ArraySet1((_497_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.IntOfUint32((_497_v44).Cardinality()))).Uint32()).(bool), 2)
					(_nw81).ArraySet1((_this).F23(), 3)
					(_nw81).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_499_v46, Companion_Default___.Fm16(globalState)), 4)
					(_nw81).ArraySet1((_this).F23(), 5)
					(_nw81).ArraySet1(p0, 6)
					(_nw81).ArraySet1(_dafny.Companion_Sequence_.Contains(_500_v47, _this.F18()), 7)
					(_nw81).ArraySet1(_dafny.Companion_Sequence_.Equal((_501_v48), _497_v44), 8)
					(_nw81).ArraySet1(p0, 9)
					(_nw81).ArraySet1(p0, 10)
					(_nw81).ArraySet1(!((_this).F23()) || (p0), 11)
					(_nw81).ArraySet1(!(p0), 12)
					(_nw81).ArraySet1(Companion_Default___.Fm1(p0, _this.F18(), _this.F22, globalState), 13)
					(_nw81).ArraySet1(p0, 14)
					_502_v49 = _nw81
					var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_502_v49), 0))
					_ = _index86
					(_502_v49).ArraySet1(p0, (_index86).Int())
					var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_502_v49), 0))
					_ = _index87
					(_502_v49).ArraySet1(Companion_Default___.Fm1(p0, (func() _dafny.CodePoint {
						if !(p0) {
							return _this.F18()
						}
						return _this.F18()
					})(), Companion_Default___.Fm0(globalState), globalState), (_index87).Int())
					var _503_v50 *C0
					_ = _503_v50
					var _nw82 *C0 = New_C0_()
					_ = _nw82
					_nw82.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("s"), _500_v47), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _this.F22), _this.F18(), ((_this).F19()).Merge((_this).F19()))
					_503_v50 = _nw82
					var _504_v51 _dafny.Array
					_ = _504_v51
					var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
					_ = _nw83
					_504_v51 = _nw83
					var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_504_v51), 0))
					_ = _index88
					(_504_v51).ArraySet1(_dafny.IntOfUint32(((_503_v50).F20()).Cardinality()), (_index88).Int())
					var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_504_v51), 0))
					_ = _index89
					(_504_v51).ArraySet1(_this.F22, (_index89).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _hi3 _dafny.Int = _this.F22
		_ = _hi3
		for _505_i5 := _this.F22; _505_i5.Cmp(_hi3) < 0; _505_i5 = _505_i5.Plus(_dafny.One) {
			var _506_v52 _dafny.Sequence
			_ = _506_v52
			_506_v52 = _dafny.SeqOf(_505_i5, _dafny.IntOfInt64(-684))
			var _507_v53 _dafny.Sequence
			_ = _507_v53
			_507_v53 = _dafny.SeqOf(_dafny.SeqOf(_505_i5), _506_v52)
			(globalState).F6 = (_505_i5).Times((_dafny.IntOfUint32((_507_v53).Cardinality())).Times(_505_i5))
			var _508_v54 _dafny.Sequence
			_ = _508_v54
			_508_v54 = _dafny.UnicodeSeqOfUtf8Bytes("isnmbr")
			var _rhs82 bool = (_this).F23()
			_ = _rhs82
			var _rhs83 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_508_v54).Cardinality()))
			_ = _rhs83
			r2 = _rhs82
			r1 = _rhs83
			r3 = p0
			r2 = !(p0)
		}
		var _509_v55 _dafny.Sequence
		_ = _509_v55
		_509_v55 = _dafny.UnicodeSeqOfUtf8Bytes("sk")
		var _510_i6 _dafny.Int
		_ = _510_i6
		_510_i6 = _dafny.Zero
		{
			for (func() bool {
				if (_this).F23() {
					return !_dafny.Companion_Sequence_.Equal(_509_v55, _509_v55)
				}
				return (_this).F23()
			})() {
				{
					if (_510_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_510_i6 = (_510_i6).Plus(_dafny.One)
					var _511_v56 _dafny.Map
					_ = _511_v56
					_511_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), _this.F22)
					var _512_v57 *C0
					_ = _512_v57
					var _nw84 *C0 = New_C0_()
					_ = _nw84
					_nw84.Ctor__(_509_v55, _511_v56, (func() _dafny.CodePoint {
						if (_this).F23() {
							return _this.F18()
						}
						return _this.F18()
					})(), ((_this).F19()).Merge((_this).F19()))
					_512_v57 = _nw84
					var _513_v58 _dafny.Sequence
					_ = _513_v58
					_513_v58 = _dafny.SeqOf(p0)
					var _514_v59 _dafny.Map
					_ = _514_v59
					_514_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, _513_v58)
					var _515_v60 _dafny.Sequence
					_ = _515_v60
					_515_v60 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if (_514_v59).Contains(_this.F22) {
							return (_514_v59).Get(_this.F22).(_dafny.Sequence)
						}
						return _dafny.SeqOf((_this).F23())
					})(), _513_v58)
					var _516_v61 _dafny.Set
					_ = _516_v61
					_516_v61 = _dafny.SetOf(p0, Companion_Default___.Fm1(!(!(p0)), _this.F18(), _dafny.IntOfInt64(169), globalState), (_this).F23())
					var _517_v62 _dafny.Map
					_ = _517_v62
					_517_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, (_this).F23())
					var _518_v63 _dafny.Array
					_ = _518_v63
					var _nwElement0_16 bool = true
					_ = _nwElement0_16
					var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(4))
					_ = _nw85
					(_nw85).ArraySet1(_nwElement0_16, 0)
					(_nw85).ArraySet1((_516_v61).IsDisjointFrom(_dafny.SetOf((_this).F23())), 1)
					(_nw85).ArraySet1((_this).F23(), 2)
					(_nw85).ArraySet1((func() bool {
						if (_517_v62).Contains(_this.F22) {
							return (_517_v62).Get(_this.F22).(bool)
						}
						return (_this).F23()
					})(), 3)
					_518_v63 = _nw85
					var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_518_v63), 0))
					_ = _index90
					(_518_v63).ArraySet1(p0, (_index90).Int())
					var _519_v64 _dafny.Sequence
					_ = _519_v64
					_519_v64 = _dafny.SeqOf(_517_v62, _517_v62)
					var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_518_v63), 0))
					_ = _index91
					var _rhs84 _dafny.Int = _this.F22
					_ = _rhs84
					var _rhs85 _dafny.Sequence = _515_v60
					_ = _rhs85
					var _rhs86 bool = true
					_ = _rhs86
					var _rhs87 _dafny.Int = _this.F22
					_ = _rhs87
					var _rhs88 bool = (_517_v62).Equals((_519_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_512_v57).F20()).Cardinality()), _dafny.IntOfUint32((_519_v64).Cardinality()))).Uint32()).(_dafny.Map))
					_ = _rhs88
					var _lhs64 *GlobalState = globalState
					_ = _lhs64
					var _lhs65 *GlobalState = globalState
					_ = _lhs65
					var _lhs66 _dafny.Array = _518_v63
					_ = _lhs66
					var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_518_v63), 0))
					_ = _lhs67
					_lhs64.F6 = _rhs84
					_515_v60 = _rhs85
					r3 = _rhs86
					_lhs65.F6 = _rhs87
					(_lhs66).ArraySet1(_rhs88, (_lhs67).Int())
					var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_518_v63), 0))
					_ = _index92
					(_518_v63).ArraySet1((_513_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.IntOfUint32((_513_v58).Cardinality()))).Uint32()).(bool), (_index92).Int())
					var _520_v65 D3
					_ = _520_v65
					_520_v65 = Companion_D3_.Create_DC8_((func() _dafny.CodePoint {
						if false {
							return _this.F18()
						}
						return _this.F18()
					})())
					var _source6 D3 = _520_v65
					_ = _source6
					if _source6.Is_DC9() {
						var _521___mcc_h4 bool = _source6.Get_().(D3_DC9).Cf11
						_ = _521___mcc_h4
						var _522___mcc_h5 _dafny.Int = _source6.Get_().(D3_DC9).Cf12
						_ = _522___mcc_h5
						var _523_cf12 _dafny.Int = _522___mcc_h5
						_ = _523_cf12
						var _524_cf11 bool = _521___mcc_h4
						_ = _524_cf11
						var _525_v66 *C0
						_ = _525_v66
						var _nw86 *C0 = New_C0_()
						_ = _nw86
						_nw86.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("isrngtcpi"), (func() _dafny.Map {
							if _524_cf11 {
								return (_512_v57).F21()
							}
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F22, _523_cf12)).Cardinality()))
						})(), _this.F18(), (_this).F19())
						_525_v66 = _nw86
						var _526_v67 _dafny.Array
						_ = _526_v67
						var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
						_ = _nw87
						_526_v67 = _nw87
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_526_v67), 0))
						_ = _index93
						(_526_v67).ArraySet1((_this.F22).Plus(_523_cf12), (_index93).Int())
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_526_v67), 0))
						_ = _index94
						(_526_v67).ArraySet1(_523_cf12, (_index94).Int())
						var _527_v68 _dafny.Array
						_ = _527_v68
						var _nwElement0_17 _dafny.CodePoint = _this.F18()
						_ = _nwElement0_17
						var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(11))
						_ = _nw88
						(_nw88).ArraySet1CodePoint(_nwElement0_17, 0)
						(_nw88).ArraySet1CodePoint(_dafny.CodePoint('w'), 1)
						(_nw88).ArraySet1CodePoint((_this).Fm13(_523_cf12, _this.F22, (_dafny.MultiSetOf(!(p0))).Cardinality(), (_this).F23(), globalState), 2)
						(_nw88).ArraySet1CodePoint(_this.F18(), 3)
						(_nw88).ArraySet1CodePoint(_this.F18(), 4)
						(_nw88).ArraySet1CodePoint((func() _dafny.CodePoint {
							if _524_cf11 {
								return _this.F18()
							}
							return _this.F18()
						})(), 5)
						(_nw88).ArraySet1CodePoint((_this).Fm13((_526_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_526_v67), 0))).Int()).(_dafny.Int), _523_cf12, _this.F22, !(false), globalState), 6)
						(_nw88).ArraySet1CodePoint(_this.F18(), 7)
						(_nw88).ArraySet1CodePoint(_this.F18(), 8)
						(_nw88).ArraySet1CodePoint(_this.F18(), 9)
						(_nw88).ArraySet1CodePoint(_this.F18(), 10)
						_527_v68 = _nw88
						_527_v68 = _527_v68
						r3 = _524_cf11
					} else if _source6.Is_DC10() {
						var _528___mcc_h6 bool = _source6.Get_().(D3_DC10).Cf13
						_ = _528___mcc_h6
						var _529_cf13 bool = _528___mcc_h6
						_ = _529_cf13
						var _530_v69 D5
						_ = _530_v69
						_530_v69 = Companion_D5_.Create_DC14_(((_518_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_518_v63), 0))).Int()).(bool)) == (_529_cf13), (_this.F22).Minus(_this.F22))
						_530_v69 = _530_v69
						r1 = (_dafny.Zero).Minus(_this.F22)
						var _531_v70 _dafny.Array
						_ = _531_v70
						var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
						_ = _nw89
						_531_v70 = _nw89
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_531_v70), 0))
						_ = _index95
						(_531_v70).ArraySet1((_512_v57).F20(), (_index95).Int())
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_531_v70), 0))
						_ = _index96
						(_531_v70).ArraySet1((_512_v57).F20(), (_index96).Int())
						var _532_v71 _dafny.Array
						_ = _532_v71
						var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
						_ = _nw90
						_532_v71 = _nw90
						var _len0_8 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_8
						var _nw91 _dafny.Array
						_ = _nw91
						if _len0_8.Cmp(_dafny.Zero) == 0 {
							_nw91 = _dafny.NewArray(_len0_8)
						} else {
							var _init8 func(_dafny.Int) _dafny.Sequence = (func(_533_v60 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_534_i7 _dafny.Int) _dafny.Sequence {
									return _533_v60
								}
							})(_515_v60)
							_ = _init8
							var _element0_8 = _init8(_dafny.Zero)
							_ = _element0_8
							_nw91 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
							(_nw91).ArraySet1(_element0_8, 0)
							var _nativeLen0_8 = (_len0_8).Int()
							_ = _nativeLen0_8
							for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
								(_nw91).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
							}
						}
						_532_v71 = _nw91
					} else {
						var _535___mcc_h7 _dafny.CodePoint = _source6.Get_().(D3_DC8).Cf10
						_ = _535___mcc_h7
						var _536_cf10 _dafny.CodePoint = _535___mcc_h7
						_ = _536_cf10
						_536_cf10 = _536_cf10
						(globalState).F6 = _this.F22
						var _537_v72 _dafny.Set
						_ = _537_v72
						_537_v72 = _dafny.SetOf(_this.F22, _dafny.IntOfInt64(513))
						var _538_v73 _dafny.Sequence
						_ = _538_v73
						_538_v73 = _dafny.SeqOf((_517_v62).Cardinality(), _this.F22)
						_537_v72 = (_dafny.SetOf((_538_v73).Select((Companion_Default___.SafeIndex(_this.F22, _dafny.IntOfUint32((_538_v73).Cardinality()))).Uint32()).(_dafny.Int))).Intersection(_537_v72)
						var _539_v74 _dafny.Sequence
						_ = _539_v74
						_539_v74 = _dafny.SeqOf(_436_v2)
						var _540_v75 _dafny.Set
						_ = _540_v75
						_540_v75 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(836), _this.F22), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(264))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}(func(_541_i8 _dafny.Int) _dafny.Int {
							return _this.F22
						}))), _dafny.SeqOf(_dafny.IntOfUint32((_539_v74).Cardinality()), _dafny.IntOfUint32((_538_v73).Cardinality())), _538_v73)
						_540_v75 = (func() _dafny.Set {
							var _coll30 = _dafny.NewBuilder()
							_ = _coll30
							for _iter31 := _dafny.Iterate((_dafny.SeqOf(_538_v73)).Elements()); ; {
								_compr_30, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _542_v76 _dafny.Sequence
								_542_v76 = interface{}(_compr_30).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_538_v73), _542_v76) {
									_coll30.Add(_542_v76)
								}
							}
							return _coll30.ToSet()
						}()).Union(_540_v75)
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _543_v77 _dafny.Array
		_ = _543_v77
		var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw92
		_543_v77 = _nw92
		var _544_v78 D3
		_ = _544_v78
		_544_v78 = Companion_D3_.Create_DC8_(_dafny.CodePoint('r'))
		var _545_v79 _dafny.MultiSet
		_ = _545_v79
		_545_v79 = _dafny.MultiSetOf(_509_v55, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-325))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_546_i9 _dafny.Int) _dafny.CodePoint {
			return _this.F18()
		})), _dafny.UnicodeSeqOfUtf8Bytes("xdamu"))
		var _547_v80 _dafny.Array
		_ = _547_v80
		var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw93
		_547_v80 = _nw93
		var _548_v81 _dafny.Sequence
		_ = _548_v81
		_548_v81 = _dafny.SeqOf(_this.F22, _this.F22, _this.F22, _this.F22, _this.F22)
		var _549_v82 _dafny.Map
		_ = _549_v82
		_549_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_547_v80, (_548_v81).Select((Companion_Default___.SafeIndex(_this.F22, _dafny.IntOfUint32((_548_v81).Cardinality()))).Uint32()).(_dafny.Int))
		var _550_v83 _dafny.Sequence
		_ = _550_v83
		_550_v83 = _dafny.SeqOf((_this).F23())
		var _rhs89 _dafny.Array = _543_v77
		_ = _rhs89
		var _rhs90 D3 = _544_v78
		_ = _rhs90
		var _rhs91 _dafny.MultiSet = _545_v79
		_ = _rhs91
		var _rhs92 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_547_v80, _this.F22)
		_ = _rhs92
		var _rhs93 _dafny.Sequence = _550_v83
		_ = _rhs93
		var _lhs68 *GlobalState = globalState
		_ = _lhs68
		_543_v77 = _rhs89
		_544_v78 = _rhs90
		_545_v79 = _rhs91
		_549_v82 = _rhs92
		_lhs68.F13 = _rhs93
		r0 = Companion_D0_.Create_DC0_(_this.F22)
		r1 = _this.F22
		r2 = !(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("yjptbhmhv"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_551_i10 _dafny.Int) _dafny.CodePoint {
			return _this.F18()
		}))))
		r3 = p0
		return r0, r1, r2, r3
	}
}
func (_this *C1) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _552_v0 _dafny.Array
		_ = _552_v0
		var _nwElement0_18 _dafny.CodePoint = _this.F18()
		_ = _nwElement0_18
		var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(15))
		_ = _nw94
		(_nw94).ArraySet1CodePoint(_nwElement0_18, 0)
		(_nw94).ArraySet1CodePoint(_dafny.CodePoint('q'), 1)
		(_nw94).ArraySet1CodePoint(_this.F18(), 2)
		(_nw94).ArraySet1CodePoint(_this.F18(), 3)
		(_nw94).ArraySet1CodePoint((_this).Fm13(p3, p2, _dafny.IntOfInt64(968), p0, globalState), 4)
		(_nw94).ArraySet1CodePoint(_this.F18(), 5)
		(_nw94).ArraySet1CodePoint(_this.F18(), 6)
		(_nw94).ArraySet1CodePoint(_this.F18(), 7)
		(_nw94).ArraySet1CodePoint(_this.F18(), 8)
		(_nw94).ArraySet1CodePoint(_this.F18(), 9)
		(_nw94).ArraySet1CodePoint(_this.F18(), 10)
		(_nw94).ArraySet1CodePoint(_this.F18(), 11)
		(_nw94).ArraySet1CodePoint(_this.F18(), 12)
		(_nw94).ArraySet1CodePoint(_this.F18(), 13)
		(_nw94).ArraySet1CodePoint(_this.F18(), 14)
		_552_v0 = _nw94
		var _553_v1 _dafny.Set
		_ = _553_v1
		_553_v1 = _dafny.SetOf(_552_v0, _552_v0)
		(globalState).F6 = (_this).Fm4(!((_this).F23()), (func() _dafny.CodePoint {
			if (_this).F23() {
				return _this.F18()
			}
			return _this.F18()
		})(), _this.F18(), (_553_v1).Contains(_552_v0), globalState)
		var _554_v2 _dafny.Sequence
		_ = _554_v2
		_554_v2 = _dafny.UnicodeSeqOfUtf8Bytes("jdgvd")
		var _555_v3 _dafny.Map
		_ = _555_v3
		_555_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_554_v2, _554_v2), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_554_v2, _554_v2)).Cardinality()))).Uint32(), _this.F18()), !(Companion_Default___.Fm1((_this).F23(), _this.F18(), _dafny.IntOfUint32((_554_v2).Cardinality()), globalState)))
		var _556_v4 _dafny.Map
		_ = _556_v4
		_556_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, _554_v2)
		_555_v3 = (_555_v3).Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_556_v4).Contains(p1) {
				return (_556_v4).Get(p1).(_dafny.Sequence)
			}
			return _554_v2
		})(), _554_v2), false)
		var _557_i0 _dafny.Int
		_ = _557_i0
		_557_i0 = _dafny.Zero
		{
			for (p1).Cmp(p2) >= 0 {
				{
					if (_557_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_557_i0 = (_557_i0).Plus(_dafny.One)
					(globalState).F6 = (p2).Minus(p1)
					var _558_v5 _dafny.Map
					_ = _558_v5
					_558_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _this.F18())
					var _559_v6 D1
					_ = _559_v6
					_559_v6 = Companion_D1_.Create_DC4_(p1, p2, (_this).F23())
					r0 = (_this).Fm4((_this).Fm12(globalState), _this.F18(), (func() _dafny.CodePoint {
						if (_558_v5).Contains((_559_v6).Dtor_cf4()) {
							return (_558_v5).Get((_559_v6).Dtor_cf4()).(_dafny.CodePoint)
						}
						return _dafny.CodePoint('y')
					})(), (_this).F23(), globalState)
					(_this).F22 = p1
					var _560_v7 _dafny.Array
					_ = _560_v7
					var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
					_ = _nw95
					_560_v7 = _nw95
					var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_560_v7), 0))
					_ = _index97
					(_560_v7).ArraySet1(Companion_Default___.SafeModuloInt(p2, (_558_v5).Cardinality()), (_index97).Int())
					var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_560_v7), 0))
					_ = _index98
					(_560_v7).ArraySet1(_dafny.IntOfInt64(365), (_index98).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _561_v8 _dafny.Sequence
		_ = _561_v8
		_561_v8 = _dafny.SeqOf(p0, false)
		if (p1).Cmp(Companion_Default___.SafeModuloInt(p1, _dafny.IntOfUint32((_561_v8).Cardinality()))) >= 0 {
			var _562_v9 _dafny.MultiSet
			_ = _562_v9
			_562_v9 = _dafny.MultiSetOf(false, (_this).F23())
			_562_v9 = (_562_v9).Union(_562_v9)
			_561_v8 = _dafny.Companion_Sequence_.Concatenate(_561_v8, _561_v8)
			var _563_v10 _dafny.Set
			_ = _563_v10
			_563_v10 = _dafny.SetOf(_dafny.IntOfInt64(299))
			_563_v10 = Companion_Default___.Fm17(globalState)
			var _564_v11 _dafny.Map
			_ = _564_v11
			_564_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, p3)
			var _565_v12 _dafny.Map
			_ = _565_v12
			_565_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), (_564_v11).Cardinality())
			var _566_v14 _dafny.Sequence
			_ = _566_v14
			_566_v14 = _dafny.SeqOf(_554_v2, _554_v2)
			var _567_v15 *C0
			_ = _567_v15
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__(_554_v2, _565_v12, _this.F18(), func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter32 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_566_v14, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_566_v14).Cardinality()))).Uint32(), _554_v2)).Elements()); ; {
					_compr_31, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _568_v13 _dafny.Sequence
					_568_v13 = interface{}(_compr_31).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_566_v14, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_566_v14).Cardinality()))).Uint32(), _554_v2), _568_v13) {
						_coll31.Add(_568_v13, (_dafny.MultiSetOf(_this.F18(), _this.F18(), _dafny.CodePoint('f'), _dafny.CodePoint('t'), _this.F18())).Cardinality())
					}
				}
				return _coll31.ToMap()
			}())
			_567_v15 = _nw96
			var _569_v16 _dafny.Map
			_ = _569_v16
			_569_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_567_v15, _567_v15)
			var _570_v17 _dafny.Map
			_ = _570_v17
			_570_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (func() *C0 {
				if (_569_v16).Contains(_567_v15) {
					return (_569_v16).Get(_567_v15).(*C0)
				}
				return _567_v15
			})())
			_570_v17 = (_570_v17).Update(!_dafny.Companion_Sequence_.Equal(_561_v8, _561_v8), _567_v15)
			var _571_v18 D6
			_ = _571_v18
			_571_v18 = Companion_D6_.Create_DC15_((_567_v15).F21())
			var _572_v19 _dafny.MultiSet
			_ = _572_v19
			_572_v19 = _dafny.MultiSetOf((_565_v12).Merge(_565_v12), (_571_v18).Dtor_cf18())
			_572_v19 = _572_v19
		} else {
			r0 = p1
			if (_this).F23() {
				var _573_v20 _dafny.Map
				_ = _573_v20
				_573_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _this.F18())
				_573_v20 = (_573_v20).Update(!((_this).F23()), _this.F18())
				var _574_v21 _dafny.Map
				_ = _574_v21
				_574_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC6_(!(!(false)), true), p0)
				var _575_v22 D1
				_ = _575_v22
				_575_v22 = Companion_D1_.Create_DC6_((_this).F23(), p0)
				_574_v21 = (_574_v21).Update(_575_v22, (_this).F23())
				(_this).F22 = p2
				(globalState).F6 = (_dafny.Zero).Minus(_this.F22)
				var _576_v23 _dafny.Map
				_ = _576_v23
				_576_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), p0)
				var _577_v24 _dafny.Array
				_ = _577_v24
				var _nwElement0_19 _dafny.Int = ((_576_v23).Cardinality()).Minus(_dafny.IntOfUint32((_561_v8).Cardinality()))
				_ = _nwElement0_19
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(11))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_19, 0)
				(_nw97).ArraySet1((_dafny.IntOfInt64(744)).Plus(p1), 1)
				(_nw97).ArraySet1(p1, 2)
				(_nw97).ArraySet1(_this.F22, 3)
				(_nw97).ArraySet1((_this).Fm4((_this).F23(), _this.F18(), _this.F18(), p0, globalState), 4)
				(_nw97).ArraySet1(Companion_Default___.Fm0(globalState), 5)
				(_nw97).ArraySet1(p1, 6)
				(_nw97).ArraySet1(p3, 7)
				(_nw97).ArraySet1(_dafny.IntOfInt64(-798), 8)
				(_nw97).ArraySet1(p2, 9)
				(_nw97).ArraySet1(p2, 10)
				_577_v24 = _nw97
				var _578_v25 _dafny.Sequence
				_ = _578_v25
				_578_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_561_v8).Cardinality()))
				var _579_v26 _dafny.MultiSet
				_ = _579_v26
				_579_v26 = _dafny.MultiSetOf((_this).F23(), p0)
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_577_v24), 0))
				_ = _index99
				(_577_v24).ArraySet1(Companion_Default___.SafeModuloInt((_578_v25).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_578_v25).Cardinality()))).Uint32()).(_dafny.Int), (_579_v26).Cardinality()), (_index99).Int())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_577_v24), 0))
				_ = _index100
				(_577_v24).ArraySet1(p3, (_index100).Int())
			} else {
				var _580_v27 T0
				_ = _580_v27
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__(_554_v2, Companion_Default___.Fm14(_dafny.IntOfInt64(356), globalState), _this.F18(), (_this).F19())
				_580_v27 = _nw98
				var _581_v28 _dafny.Sequence
				_ = _581_v28
				_581_v28 = _dafny.SeqOf(_580_v27)
				var _582_v29 _dafny.Map
				_ = _582_v29
				_582_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _dafny.Companion_Sequence_.Concatenate(_581_v28, _581_v28))
				var _583_v30 _dafny.Array
				_ = _583_v30
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw99
				_583_v30 = _nw99
				var _584_v31 _dafny.Map
				_ = _584_v31
				_584_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _this.F22)
				var _585_v32 *C0
				_ = _585_v32
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("v"), _584_v31, _this.F18(), (_580_v27).F19())
				_585_v32 = _nw100
				var _586_v33 _dafny.Map
				_ = _586_v33
				_586_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v32, !((_this).F23()))
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_583_v30), 0))
				_ = _index101
				(_583_v30).ArraySet1((func() bool {
					if (_586_v33).Contains(_585_v32) {
						return (_586_v33).Get(_585_v32).(bool)
					}
					return p0
				})(), (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_583_v30), 0))
				_ = _index102
				var _rhs94 _dafny.Map = (_582_v29).Update((_this).F23(), _581_v28)
				_ = _rhs94
				var _rhs95 bool = (_this).F23()
				_ = _rhs95
				var _lhs69 _dafny.Array = _583_v30
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_583_v30), 0))
				_ = _lhs70
				_582_v29 = _rhs94
				(_lhs69).ArraySet1(_rhs95, (_lhs70).Int())
				var _587_v34 _dafny.Map
				_ = _587_v34
				_587_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), ((_this).F23()) == ((_this).F23()))
				_587_v34 = (_587_v34).Update(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}((func(_588_v27 T0) func(_dafny.Int) _dafny.CodePoint {
					return func(_589_i1 _dafny.Int) _dafny.CodePoint {
						return _588_v27.F18()
					}
				})(_580_v27))), _554_v2), (_this).F23())
				var _590_v35 _dafny.Sequence
				_ = _590_v35
				_590_v35 = _dafny.SeqOf(_this.F22, p1, p2)
				var _591_v36 _dafny.Set
				_ = _591_v36
				_591_v36 = _dafny.SetOf(_dafny.IntOfInt64(304), _dafny.IntOfUint32((_590_v35).Cardinality()))
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_583_v30), 0))
				_ = _index103
				(_583_v30).ArraySet1(Companion_Default___.Fm1(p0, _580_v27.F18(), ((_591_v36).Cardinality()).Minus((_590_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-813))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_592_v27 T0) func(_dafny.Int) _dafny.CodePoint {
					return func(_593_i2 _dafny.Int) _dafny.CodePoint {
						return _592_v27.F18()
					}
				})(_580_v27)))).Cardinality()), _dafny.IntOfUint32((_590_v35).Cardinality()))).Uint32()).(_dafny.Int)), globalState), (_index103).Int())
				var _594_v37 _dafny.Sequence
				_ = _594_v37
				_594_v37 = _dafny.SeqOf((_585_v32).F20(), _554_v2, _dafny.Companion_Sequence_.Concatenate(_554_v2, _554_v2))
				var _rhs96 _dafny.Int = (_590_v35).Select((Companion_Default___.SafeIndex((_555_v3).Cardinality(), _dafny.IntOfUint32((_590_v35).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs96
				var _rhs97 _dafny.Sequence = _594_v37
				_ = _rhs97
				var _rhs98 _dafny.Int = Companion_Default___.SafeModuloInt((Companion_D5_.Create_DC14_(!(false), p1)).Dtor_cf17(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-221), _dafny.IntOfInt64(274)))
				_ = _rhs98
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				_lhs71.F6 = _rhs96
				_594_v37 = _rhs97
				_lhs72.F6 = _rhs98
				_554_v2 = _dafny.Companion_Sequence_.Update((_585_v32).F20(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_585_v32).F20()).Cardinality()))).Uint32(), _580_v27.F18())
			}
			if (_this).Fm3(!((p0) && ((_this).F23())), (_dafny.IntOfUint32((_554_v2).Cardinality())).Plus(_dafny.IntOfUint32((_554_v2).Cardinality())), globalState) {
				(_this).F18_set_(_this.F18())
				var _595_v38 bool
				_ = _595_v38
				_595_v38 = false
				_595_v38 = !(p0)
				(_this).F22 = p2
				var _596_v39 D1
				_ = _596_v39
				_596_v39 = Companion_D1_.Create_DC4_(p3, p1, _595_v38)
				var _597_v40 _dafny.Map
				_ = _597_v40
				_597_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_596_v39).Dtor_cf5(), (_this).F23())
				_597_v40 = (_597_v40).Update((_this).F23(), (_this).F23())
				(_this).F22 = (_dafny.Zero).Minus(p2)
			} else {
				var _598_v41 _dafny.MultiSet
				_ = _598_v41
				_598_v41 = _dafny.MultiSetOf(p1, p1, p1, p1)
				var _599_v42 D3
				_ = _599_v42
				_599_v42 = Companion_D3_.Create_DC10_((_598_v41).IsSubsetOf(_598_v41))
				_599_v42 = (func() D3 {
					if (_this).F23() {
						return Companion_D3_.Create_DC10_((_this).F23())
					}
					return _599_v42
				})()
				var _600_v43 _dafny.Array
				_ = _600_v43
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_9
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Int = (func(_601_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_602_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_602_i3, _601_p3)
						}
					})(p3)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw101 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw101).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw101).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_600_v43 = _nw101
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_600_v43), 0))
				_ = _index104
				(_600_v43).ArraySet1(p3, (_index104).Int())
				var _603_v44 _dafny.Set
				_ = _603_v44
				_603_v44 = _dafny.SetOf(p2, p3)
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_600_v43), 0))
				_ = _index105
				(_600_v43).ArraySet1((p2).Plus((_dafny.IntOfInt64(693)).Times((_603_v44).Cardinality())), (_index105).Int())
				var _604_v45 _dafny.Sequence
				_ = _604_v45
				_604_v45 = _dafny.SeqOf(_552_v0, _552_v0, _552_v0, _552_v0, _552_v0)
				var _605_v46 _dafny.Sequence
				_ = _605_v46
				_605_v46 = _dafny.SeqOf(_604_v45)
				var _606_v47 _dafny.Set
				_ = _606_v47
				_606_v47 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("tqhnwiyw"), _554_v2, _554_v2)
				var _607_v48 _dafny.Map
				_ = _607_v48
				_607_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_554_v2, (_605_v46).Select((Companion_Default___.SafeIndex((_606_v47).Cardinality(), _dafny.IntOfUint32((_605_v46).Cardinality()))).Uint32()).(_dafny.Sequence))
				_607_v48 = (_607_v48).Update(_dafny.Companion_Sequence_.Concatenate(_554_v2, _dafny.UnicodeSeqOfUtf8Bytes("jrv")), _604_v45)
				var _608_v49 bool
				_ = _608_v49
				_608_v49 = true
				var _609_v50 _dafny.MultiSet
				_ = _609_v50
				_609_v50 = _dafny.MultiSetOf((_554_v2).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_554_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F18())
				_608_v49 = (func() bool {
					if _608_v49 {
						return (_561_v8).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_609_v50).Contains(_this.F18()) {
								return (_609_v50).Multiplicity(_this.F18())
							}
							return (_598_v41).Cardinality()
						})(), _dafny.IntOfUint32((_561_v8).Cardinality()))).Uint32()).(bool)
					}
					return ((_609_v50).Cardinality()).Cmp((_600_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.ArrayLen((_600_v43), 0))).Int()).(_dafny.Int)) <= 0
				})()
				var _610_v51 _dafny.Map
				_ = _610_v51
				_610_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_608_v49, _this.F22)
				_610_v51 = (_610_v51).Update(p0, (_dafny.IntOfInt64(843)).Minus(p2))
			}
			if (_561_v8).Select((Companion_Default___.SafeIndex(_this.F22, _dafny.IntOfUint32((_561_v8).Cardinality()))).Uint32()).(bool) {
				var _611_v52 _dafny.Array
				_ = _611_v52
				var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw102
				_611_v52 = _nw102
				var _612_v53 D6
				_ = _612_v53
				_612_v53 = Companion_D6_.Create_DC16_(!((_this).F23()), _611_v52)
				var _613_v54 D6
				_ = _613_v54
				_613_v54 = Companion_D6_.Create_DC17_(_612_v53)
				var _614_v55 _dafny.Map
				_ = _614_v55
				_614_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_613_v54, _dafny.CodePoint('y'))
				_614_v55 = (_614_v55).Update(_613_v54, _this.F18())
				(_this).F22 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_615_p3 _dafny.Int, _616_p0 bool) func(_dafny.Int) _dafny.Int {
					return func(_617_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_615_p3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _616_p0)).Cardinality())
					}
				})(p3, p0)))).Cardinality())
				var _618_v56 _dafny.MultiSet
				_ = _618_v56
				_618_v56 = _dafny.MultiSetOf(_dafny.IntOfInt64(527))
				var _619_v57 _dafny.Set
				_ = _619_v57
				_619_v57 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F18()))
				var _620_v58 _dafny.Set
				_ = _620_v58
				_620_v58 = _dafny.SetOf((_this).F23())
				var _621_v59 _dafny.Array
				_ = _621_v59
				var _nwElement0_20 bool = p0
				_ = _nwElement0_20
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(27))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_20, 0)
				(_nw103).ArraySet1((_this).F23(), 1)
				(_nw103).ArraySet1((_this).F23(), 2)
				(_nw103).ArraySet1(!(true) || (!(Companion_Default___.Fm1(p0, _this.F18(), (_618_v56).Cardinality(), globalState))), 3)
				(_nw103).ArraySet1(!((_this).F23()), 4)
				(_nw103).ArraySet1(p0, 5)
				(_nw103).ArraySet1(true, 6)
				(_nw103).ArraySet1((_this).Fm3(false, _dafny.IntOfUint32((_dafny.SeqOf(_this.F22)).Cardinality()), globalState), 7)
				(_nw103).ArraySet1(p0, 8)
				(_nw103).ArraySet1(!((_this).F23()), 9)
				(_nw103).ArraySet1((_this).F23(), 10)
				(_nw103).ArraySet1((_this).F23(), 11)
				(_nw103).ArraySet1(true, 12)
				(_nw103).ArraySet1((_561_v8).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_561_v8).Cardinality()))).Uint32()).(bool), 13)
				(_nw103).ArraySet1(p0, 14)
				(_nw103).ArraySet1((_this.F22).Cmp(p2) >= 0, 15)
				(_nw103).ArraySet1((_619_v57).IsProperSubsetOf(_619_v57), 16)
				(_nw103).ArraySet1(p0, 17)
				(_nw103).ArraySet1(true, 18)
				(_nw103).ArraySet1(p0, 19)
				(_nw103).ArraySet1(!((_this).F23()), 20)
				(_nw103).ArraySet1((_this).F23(), 21)
				(_nw103).ArraySet1((_this).F23(), 22)
				(_nw103).ArraySet1((_620_v58).IsDisjointFrom(_620_v58), 23)
				(_nw103).ArraySet1(p0, 24)
				(_nw103).ArraySet1(p0, 25)
				(_nw103).ArraySet1((_this).F23(), 26)
				_621_v59 = _nw103
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_621_v59), 0))
				_ = _index106
				(_621_v59).ArraySet1(p0, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_621_v59), 0))
				_ = _index107
				(_621_v59).ArraySet1((func() bool {
					if (p1).Cmp(_dafny.IntOfUint32((_554_v2).Cardinality())) <= 0 {
						return p0
					}
					return (p2).Cmp(p2) != 0
				})(), (_index107).Int())
				var _622_v60 D1
				_ = _622_v60
				_622_v60 = Companion_D1_.Create_DC4_(p2, _dafny.IntOfInt64(192), (_this).F23())
				var _623_v61 _dafny.Sequence
				_ = _623_v61
				_623_v61 = _dafny.SeqOf((_this).Fm4(Companion_Default___.Fm1((_this).F23(), _this.F18(), p3, globalState), _this.F18(), _dafny.CodePoint('d'), !(p0), globalState))
				var _624_v62 _dafny.Set
				_ = _624_v62
				_624_v62 = _dafny.SetOf(p3, _this.F22, _dafny.IntOfInt64(241), p2, (_618_v56).Cardinality())
				var _625_v63 _dafny.Array
				_ = _625_v63
				var _nwElement0_21 D1 = _622_v60
				_ = _nwElement0_21
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(24))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_21, 0)
				(_nw104).ArraySet1(_622_v60, 1)
				(_nw104).ArraySet1(Companion_D1_.Create_DC4_(p1, p1, true), 2)
				(_nw104).ArraySet1(_622_v60, 3)
				(_nw104).ArraySet1(Companion_D1_.Create_DC4_(p1, p3, (_this).F23()), 4)
				(_nw104).ArraySet1(_622_v60, 5)
				(_nw104).ArraySet1(_622_v60, 6)
				(_nw104).ArraySet1(Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_623_v61).Cardinality()), (_624_v62).Cardinality(), false), 7)
				(_nw104).ArraySet1(_622_v60, 8)
				(_nw104).ArraySet1(_622_v60, 9)
				(_nw104).ArraySet1(_622_v60, 10)
				(_nw104).ArraySet1(_622_v60, 11)
				(_nw104).ArraySet1(_622_v60, 12)
				(_nw104).ArraySet1(_622_v60, 13)
				(_nw104).ArraySet1(Companion_D1_.Create_DC4_(p3, p3, (_this).F23()), 14)
				(_nw104).ArraySet1(_622_v60, 15)
				(_nw104).ArraySet1(_622_v60, 16)
				(_nw104).ArraySet1(_622_v60, 17)
				(_nw104).ArraySet1(_622_v60, 18)
				(_nw104).ArraySet1(_622_v60, 19)
				(_nw104).ArraySet1(Companion_D1_.Create_DC4_(_this.F22, p2, p0), 20)
				(_nw104).ArraySet1(_622_v60, 21)
				(_nw104).ArraySet1(Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_561_v8).Cardinality()), _this.F22, (_561_v8).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_561_v8).Cardinality()))).Uint32()).(bool)), 22)
				(_nw104).ArraySet1(_622_v60, 23)
				_625_v63 = _nw104
				var _626_v64 _dafny.MultiSet
				_ = _626_v64
				_626_v64 = _dafny.MultiSetOf(_625_v63, _625_v63, _625_v63)
				_626_v64 = _626_v64
				var _627_v65 _dafny.Map
				_ = _627_v65
				_627_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), p3)
				var _628_v66 *C0
				_ = _628_v66
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("scd"), _627_v65, _this.F18(), (_this).F19())
				_628_v66 = _nw105
			} else {
				(globalState).F6 = p3
				var _629_v67 _dafny.MultiSet
				_ = _629_v67
				_629_v67 = _dafny.MultiSetOf(_dafny.IntOfUint32((_561_v8).Cardinality()), p3, _dafny.IntOfInt64(542))
				var _rhs99 _dafny.Int = p2
				_ = _rhs99
				var _rhs100 _dafny.MultiSet = _629_v67
				_ = _rhs100
				var _lhs73 *C1 = _this
				_ = _lhs73
				_lhs73.F22 = _rhs99
				_629_v67 = _rhs100
				var _630_v68 D3
				_ = _630_v68
				_630_v68 = Companion_D3_.Create_DC8_((_this).Fm13((_dafny.Zero).Minus(p2), p1, p2, (_this).F23(), globalState))
				var _631_v69 _dafny.Map
				_ = _631_v69
				_631_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_630_v68, (_this).F23())
				var _632_v70 _dafny.Set
				_ = _632_v70
				_632_v70 = _dafny.SetOf(p0, !(p0), p0)
				_631_v69 = (_631_v69).Update(func(_pat_let12_0 D3) D3 {
					return func(_633_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let13_0 _dafny.CodePoint) D3 {
							return func(_634_dt__update_hcf10_h0 _dafny.CodePoint) D3 {
								return Companion_D3_.Create_DC8_(_634_dt__update_hcf10_h0)
							}(_pat_let13_0)
						}(_dafny.CodePoint('a'))
					}(_pat_let12_0)
				}(_630_v68), (_dafny.SetOf(p0)).IsProperSubsetOf(_632_v70))
				var _635_v71 bool
				_ = _635_v71
				_635_v71 = true
				_635_v71 = _635_v71
				var _636_v72 _dafny.Array
				_ = _636_v72
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw106
				_636_v72 = _nw106
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_636_v72), 0))
				_ = _index108
				(_636_v72).ArraySet1(Companion_Default___.SafeModuloInt(p3, p3), (_index108).Int())
				var _637_v73 _dafny.Map
				_ = _637_v73
				_637_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F22, _dafny.IntOfInt64(931))
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_636_v72), 0))
				_ = _index109
				var _rhs101 _dafny.Int = _this.F22
				_ = _rhs101
				var _rhs102 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_629_v67).Contains(p1) {
						return (_629_v67).Multiplicity(p1)
					}
					return (func() _dafny.Int {
						if (_637_v73).Contains(_this.F22) {
							return (_637_v73).Get(_this.F22).(_dafny.Int)
						}
						return _dafny.IntOfInt64(698)
					})()
				})(), (_dafny.Zero).Minus((func() _dafny.Int {
					if _635_v71 {
						return p1
					}
					return _this.F22
				})()))
				_ = _rhs102
				var _rhs103 bool = (p0) && (false)
				_ = _rhs103
				var _lhs74 _dafny.Array = _636_v72
				_ = _lhs74
				var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_636_v72), 0))
				_ = _lhs75
				var _lhs76 *GlobalState = globalState
				_ = _lhs76
				(_lhs74).ArraySet1(_rhs101, (_lhs75).Int())
				_lhs76.F6 = _rhs102
				_635_v71 = _rhs103
			}
			var _638_v74 bool
			_ = _638_v74
			_638_v74 = true
			_638_v74 = Companion_Default___.Fm1((_553_v1).IsSubsetOf(_553_v1), _dafny.CodePoint('y'), p3, globalState)
		}
		(globalState).F6 = _dafny.IntOfInt64(-101)
		var _639_i5 _dafny.Int
		_ = _639_i5
		_639_i5 = _dafny.Zero
		{
			for (_561_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.IntOfUint32((_561_v8).Cardinality()))).Uint32()).(bool) {
				{
					if (_639_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_639_i5 = (_639_i5).Plus(_dafny.One)
					var _640_v75 _dafny.Array
					_ = _640_v75
					var _len0_10 _dafny.Int = _dafny.IntOfInt64(4)
					_ = _len0_10
					var _nw107 _dafny.Array
					_ = _nw107
					if _len0_10.Cmp(_dafny.Zero) == 0 {
						_nw107 = _dafny.NewArray(_len0_10)
					} else {
						var _init10 func(_dafny.Int) _dafny.Int = (func(_641_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_642_i6 _dafny.Int) _dafny.Int {
								return (_642_i6).Plus(_641_p3)
							}
						})(p3)
						_ = _init10
						var _element0_10 = _init10(_dafny.Zero)
						_ = _element0_10
						_nw107 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
						(_nw107).ArraySet1(_element0_10, 0)
						var _nativeLen0_10 = (_len0_10).Int()
						_ = _nativeLen0_10
						for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
							(_nw107).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
						}
					}
					_640_v75 = _nw107
					var _643_v76 _dafny.Map
					_ = _643_v76
					_643_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), false)
					var _644_v77 _dafny.Map
					_ = _644_v77
					_644_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus((_643_v76).Cardinality()))
					var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_640_v75), 0))
					_ = _index110
					(_640_v75).ArraySet1(Companion_Default___.SafeDivisionInt(p1, (_644_v77).Cardinality()), (_index110).Int())
					var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_640_v75), 0))
					_ = _index111
					(_640_v75).ArraySet1(p1, (_index111).Int())
					var _645_v78 _dafny.MultiSet
					_ = _645_v78
					_645_v78 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(p0, (_this).F23())).Cardinality()), p1)
					var _646_v79 D3
					_ = _646_v79
					_646_v79 = Companion_D3_.Create_DC9_(p0, (_645_v78).Cardinality())
					var _647_v80 D1
					_ = _647_v80
					_647_v80 = Companion_D1_.Create_DC4_((_646_v79).Dtor_cf12(), _this.F22, (_this).F23())
					_647_v80 = _647_v80
					(globalState).F6 = _this.F22
					_640_v75 = _640_v75
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		r0 = (p1).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(599), _this.F22))
		return r0
	}
}
func (_this *C1) F23() bool {
	{
		return _this._f23
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f18 _dafny.CodePoint
	_f19 _dafny.Map
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f18 = _dafny.CodePoint('D')
	_this._f19 = _dafny.EmptyMap
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

func (_this *C2) F18() _dafny.CodePoint {
	return _this._f18
}
func (_this *C2) F18_set_(value _dafny.CodePoint) {
	_this._f18 = value
}
func (_this *C2) F19() _dafny.Map {
	return _this._f19
}
func (_this *C2) Ctor__(f18 _dafny.CodePoint, f19 _dafny.Map) {
	{
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C2) Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		if (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fywgn")).Cardinality())).Cmp(_dafny.IntOfInt64(896)) <= 0 {
			return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bnebdr")).Cardinality())).Cmp(_dafny.IntOfInt64(329)) <= 0
		} else {
			return true
		}
	}
}
func (_this *C2) Fm4(p0 bool, p1 _dafny.CodePoint, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false, false, !(true), false))).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(971))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_648_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true, !(!(false)), true, !(!(true)), !(false))).Cardinality())
		}))).Cardinality())))
	}
}
func (_this *C2) M1(p0 bool, globalState *GlobalState) (D0, _dafny.Int, bool, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _649_v0 _dafny.Array
		_ = _649_v0
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_11
		var _nw108 _dafny.Array
		_ = _nw108
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw108 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = func(_650_i0 _dafny.Int) _dafny.Int {
				return (_650_i0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fugltux")).Cardinality()))
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw108 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw108).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw108).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_649_v0 = _nw108
		var _651_v1 _dafny.Int
		_ = _651_v1
		_651_v1 = _dafny.IntOfInt64(-605)
		var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))
		_ = _index112
		(_649_v0).ArraySet1(_651_v1, (_index112).Int())
		var _652_v2 _dafny.Array
		_ = _652_v2
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_12
		var _nw109 _dafny.Array
		_ = _nw109
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw109 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Map = (func(_653_p0 bool, _654_v1 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_655_i1 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yxgl"), _653_p0)).Cardinality(), Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-553))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}((func(_656_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_657_i2 _dafny.Int) _dafny.Int {
							return _656_v1
						}
					})(_654_v1)))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v1, Companion_D0_.Create_DC0_(_654_v1)))
				}
			})(p0, _651_v1)
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
		_652_v2 = _nw109
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_649_v0), 0))
		_ = _index113
		(_649_v0).ArraySet1(_651_v1, (_index113).Int())
		var _658_v3 _dafny.Sequence
		_ = _658_v3
		_658_v3 = _dafny.UnicodeSeqOfUtf8Bytes("uuruhxs")
		var _659_v4 _dafny.Map
		_ = _659_v4
		_659_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_651_v1, _658_v3)
		var _660_v5 _dafny.Array
		_ = _660_v5
		var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw110
		_660_v5 = _nw110
		var _661_v6 _dafny.MultiSet
		_ = _661_v6
		_661_v6 = _dafny.MultiSetOf(_660_v5, _660_v5)
		var _662_v7 _dafny.Sequence
		_ = _662_v7
		_662_v7 = _dafny.SeqOf(!(p0))
		var _663_v8 _dafny.Sequence
		_ = _663_v8
		_663_v8 = _662_v7
		var _pat_let_tv10 = _651_v1
		_ = _pat_let_tv10
		var _pat_let_tv11 = _651_v1
		_ = _pat_let_tv11
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))
		_ = _index114
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_649_v0), 0))
		_ = _index115
		var _rhs104 bool = ((_659_v4).Cardinality()).Cmp((_dafny.IntOfInt64(579)).Times((_661_v6).Cardinality())) <= 0
		_ = _rhs104
		var _rhs105 _dafny.CodePoint = _this.F18()
		_ = _rhs105
		var _rhs106 _dafny.Int = _651_v1
		_ = _rhs106
		var _rhs107 _dafny.Array = _652_v2
		_ = _rhs107
		var _rhs108 _dafny.Int = func(_source7 _dafny.Sequence) _dafny.Int {
			var _664___mcc_h0 _dafny.Sequence = _source7
			_ = _664___mcc_h0
			var _665_cf9 _dafny.Sequence = _664___mcc_h0
			_ = _665_cf9
			return (_pat_let_tv10).Times(_pat_let_tv11)
		}(_663_v8)
		_ = _rhs108
		var _lhs77 *C2 = _this
		_ = _lhs77
		var _lhs78 _dafny.Array = _649_v0
		_ = _lhs78
		var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))
		_ = _lhs79
		var _lhs80 _dafny.Array = _649_v0
		_ = _lhs80
		var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_649_v0), 0))
		_ = _lhs81
		r2 = _rhs104
		_lhs77.F18_set_(_rhs105)
		(_lhs78).ArraySet1(_rhs106, (_lhs79).Int())
		_652_v2 = _rhs107
		(_lhs80).ArraySet1(_rhs108, (_lhs81).Int())
		var _666_v9 _dafny.Map
		_ = _666_v9
		_666_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int))
		var _667_v10 *C0
		_ = _667_v10
		var _nw111 *C0 = New_C0_()
		_ = _nw111
		_nw111.Ctor__(_658_v3, _666_v9, _this.F18(), (_this).F19())
		_667_v10 = _nw111
		_651_v1 = (_651_v1).Times(_dafny.IntOfInt64(988))
		var _668_v11 _dafny.Map
		_ = _668_v11
		_668_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int), (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int))
		var _669_v12 _dafny.Sequence
		_ = _669_v12
		_669_v12 = _dafny.SeqOf(((_668_v11).Cardinality()).Plus(_651_v1), (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int), (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int))
		r1 = (_669_v12).Select((Companion_Default___.SafeIndex(_651_v1, _dafny.IntOfUint32((_669_v12).Cardinality()))).Uint32()).(_dafny.Int)
		var _hi4 _dafny.Int = _651_v1
		_ = _hi4
		for _670_i3 := (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int); _670_i3.Cmp(_hi4) < 0; _670_i3 = _670_i3.Plus(_dafny.One) {
			var _671_v13 bool
			_ = _671_v13
			var _out23 bool
			_ = _out23
			_out23 = (_this).M4((_670_i3).Times(_651_v1), Companion_Default___.Fm0(globalState), _660_v5, globalState)
			_671_v13 = _out23
			r3 = _671_v13
			var _672_v14 D0
			_ = _672_v14
			_672_v14 = Companion_D0_.Create_DC1_()
			var _source8 D0 = _672_v14
			_ = _source8
			if _source8.Is_DC1() {
				var _673_v15 _dafny.MultiSet
				_ = _673_v15
				_673_v15 = _dafny.MultiSetOf(_671_v13)
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))
				_ = _index116
				(_649_v0).ArraySet1((_673_v15).Cardinality(), (_index116).Int())
				(globalState).F13 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_662_v7, _662_v7), _dafny.SeqOf(_671_v13))
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_660_v5), 0))
				_ = _index117
				(_660_v5).ArraySet1(false, (_index117).Int())
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_660_v5), 0))
				_ = _index118
				(_660_v5).ArraySet1(_671_v13, (_index118).Int())
				var _674_v16 _dafny.MultiSet
				_ = _674_v16
				_674_v16 = _dafny.MultiSetOf(_649_v0, _649_v0, _649_v0)
				var _675_v17 _dafny.Sequence
				_ = _675_v17
				_675_v17 = _dafny.SeqOf(_674_v16, _674_v16)
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_660_v5), 0))
				_ = _index119
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_660_v5), 0))
				_ = _index120
				var _rhs109 bool = _671_v13
				_ = _rhs109
				var _rhs110 _dafny.Int = Companion_Default___.Fm0(globalState)
				_ = _rhs110
				var _rhs111 bool = !((_675_v17).Select((Companion_Default___.SafeIndex(_651_v1, _dafny.IntOfUint32((_675_v17).Cardinality()))).Uint32()).(_dafny.MultiSet)).Equals((_674_v16).Update(_649_v0, Companion_Default___.Abs(_dafny.IntOfUint32(((_667_v10).F20()).Cardinality()))))
				_ = _rhs111
				var _rhs112 bool = (_670_i3).Cmp(Companion_Default___.SafeModuloInt((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int), _651_v1)) >= 0
				_ = _rhs112
				var _rhs113 bool = p0
				_ = _rhs113
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				var _lhs83 _dafny.Array = _660_v5
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_660_v5), 0))
				_ = _lhs84
				var _lhs85 _dafny.Array = _660_v5
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_660_v5), 0))
				_ = _lhs86
				_671_v13 = _rhs109
				_lhs82.F6 = _rhs110
				_671_v13 = _rhs111
				(_lhs83).ArraySet1(_rhs112, (_lhs84).Int())
				(_lhs85).ArraySet1(_rhs113, (_lhs86).Int())
				_671_v13 = _671_v13
			} else if _source8.Is_DC0() {
				var _676___mcc_h1 _dafny.Int = _source8.Get_().(D0_DC0).Cf0
				_ = _676___mcc_h1
				var _677_cf0 _dafny.Int = _676___mcc_h1
				_ = _677_cf0
				var _678_v18 _dafny.Set
				_ = _678_v18
				_678_v18 = _dafny.SetOf(_671_v13)
				_678_v18 = _678_v18
				(globalState).F6 = (_651_v1).Times(_651_v1)
				var _679_v19 _dafny.Array
				_ = _679_v19
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw112
				_679_v19 = _nw112
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_679_v19), 0))
				_ = _index121
				(_679_v19).ArraySet1(_dafny.SeqOf(_671_v13), (_index121).Int())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_679_v19), 0))
				_ = _index122
				(_679_v19).ArraySet1(_663_v8, (_index122).Int())
				var _680_v20 _dafny.Map
				_ = _680_v20
				_680_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v3, !(p0))
				var _681_v21 _dafny.Map
				_ = _681_v21
				_681_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if _671_v13 {
						return _dafny.IntOfInt64(-196)
					}
					return _670_i3
				})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_667_v10).F20(), _671_v13)).Merge(_680_v20))
				_681_v21 = (_681_v21).Update((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int), (func() _dafny.Map {
					if _671_v13 {
						return Companion_Default___.Fm11(globalState)
					}
					return (_680_v20).Update((_667_v10).F20(), true)
				})())
			} else {
				var _682___mcc_h2 D0 = _source8.Get_().(D0_DC2).Cf1
				_ = _682___mcc_h2
				var _683_cf1 D0 = _682___mcc_h2
				_ = _683_cf1
				var _684_v22 T0
				_ = _684_v22
				var _nw113 *C0 = New_C0_()
				_ = _nw113
				_nw113.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("qfwpi"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), _651_v1), _dafny.CodePoint('m'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_667_v10).F20(), (Companion_Default___.SafeIndex((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_667_v10).F20()).Cardinality()))).Uint32(), _this.F18()), (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)))
				_684_v22 = _nw113
				var _685_v23 T0
				_ = _685_v23
				_685_v23 = _684_v22
				var _686_v24 T0
				_ = _686_v24
				_686_v24 = (_685_v23)
				_684_v22 = (_684_v22)
				var _687_v26 _dafny.MultiSet
				_ = _687_v26
				_687_v26 = _dafny.MultiSetOf(_658_v3)
				var _688_v27 *C0
				_ = _688_v27
				var _nw114 *C0 = New_C0_()
				_ = _nw114
				_nw114.Ctor__((_667_v10).F20(), _666_v9, _dafny.CodePoint('l'), func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter33 := _dafny.Iterate((_687_v26).Elements()); ; {
						_compr_32, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _689_v25 _dafny.Sequence
						_689_v25 = interface{}(_compr_32).(_dafny.Sequence)
						if (_687_v26).Contains(_689_v25) {
							_coll32.Add(_689_v25, _dafny.IntOfUint32(((_667_v10).F20()).Cardinality()))
						}
					}
					return _coll32.ToMap()
				}())
				_688_v27 = _nw114
				var _690_v28 _dafny.Map
				_ = _690_v28
				_690_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_667_v10, (_dafny.Zero).Minus(_651_v1))
				var _rhs114 bool = ((func() _dafny.Int {
					if (_690_v28).Contains(_688_v27) {
						return (_690_v28).Get(_688_v27).(_dafny.Int)
					}
					return (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)
				})()).Cmp(_651_v1) < 0
				_ = _rhs114
				var _rhs115 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _rhs115
				_671_v13 = _rhs114
				_651_v1 = _rhs115
				var _691_v29 T1
				_ = _691_v29
				var _nw115 *C1 = New_C1_()
				_ = _nw115
				_nw115.Ctor__(_651_v1, _671_v13, _this.F18(), Companion_Default___.Fm18(_dafny.IntOfUint32(((_667_v10).F20()).Cardinality()), globalState))
				_691_v29 = _nw115
				var _692_v30 _dafny.Map
				_ = _692_v30
				_692_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_662_v7, _691_v29)
				var _693_v31 _dafny.Map
				_ = _693_v31
				_693_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_662_v7, _691_v29))
				var _rhs116 _dafny.Map = ((func() _dafny.Map {
					if (_693_v31).Contains(p0) {
						return (_693_v31).Get(p0).(_dafny.Map)
					}
					return _692_v30
				})()).Merge(_692_v30)
				_ = _rhs116
				var _rhs117 _dafny.Int = (_651_v1).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_694_v29 T1) func(_dafny.Int) _dafny.CodePoint {
					return func(_695_i4 _dafny.Int) _dafny.CodePoint {
						return _694_v29.F18()
					}
				})(_691_v29)))).Cardinality()))
				_ = _rhs117
				_692_v30 = _rhs116
				_651_v1 = _rhs117
			}
			r2 = !(p0)
		}
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_660_v5), 0))
		_ = _index123
		(_660_v5).ArraySet1(p0, (_index123).Int())
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))
		_ = _index124
		var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_660_v5), 0))
		_ = _index125
		var _rhs118 _dafny.Int = (_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)
		_ = _rhs118
		var _rhs119 bool = _dafny.Companion_Sequence_.Equal((_667_v10).F20(), (_667_v10).F20())
		_ = _rhs119
		var _rhs120 _dafny.CodePoint = _this.F18()
		_ = _rhs120
		var _rhs121 _dafny.Int = _651_v1
		_ = _rhs121
		var _lhs87 _dafny.Array = _649_v0
		_ = _lhs87
		var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))
		_ = _lhs88
		var _lhs89 _dafny.Array = _660_v5
		_ = _lhs89
		var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_660_v5), 0))
		_ = _lhs90
		var _lhs91 *C2 = _this
		_ = _lhs91
		(_lhs87).ArraySet1(_rhs118, (_lhs88).Int())
		(_lhs89).ArraySet1(_rhs119, (_lhs90).Int())
		_lhs91.F18_set_(_rhs120)
		_651_v1 = _rhs121
		r0 = Companion_D0_.Create_DC0_(((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)).Minus((_this).Fm4(p0, _this.F18(), _dafny.CodePoint('r'), (_660_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_660_v5), 0))).Int()).(bool), globalState)))
		r1 = (func() _dafny.Int {
			if (_668_v11).Contains((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)) {
				return (_668_v11).Get((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(380))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_696_v11 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_697_i5 _dafny.Int) _dafny.Map {
					return _696_v11
				}
			})(_668_v11)))).Cardinality())
		})()
		r2 = (func() bool {
			if (_660_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_660_v5), 0))).Int()).(bool) {
				return p0
			}
			return ((_649_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_649_v0), 0))).Int()).(_dafny.Int)).Cmp(_651_v1) < 0
		})()
		var _698_v32 _dafny.Set
		_ = _698_v32
		_698_v32 = _dafny.SetOf(!(false))
		r3 = (_698_v32).IsSubsetOf(_698_v32)
		return r0, r1, r2, r3
	}
}
func (_this *C2) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		r0 = p2
		if p0 {
			r0 = p3
			r0 = p3
			var _699_v0 _dafny.Array
			_ = _699_v0
			var _nw116 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw116
			_699_v0 = _nw116
			_699_v0 = _699_v0
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_699_v0), 0))
			_ = _index126
			(_699_v0).ArraySet1(p0, (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_699_v0), 0))
			_ = _index127
			(_699_v0).ArraySet1((p3).Cmp((_dafny.Zero).Minus(p3)) > 0, (_index127).Int())
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_699_v0), 0))
			_ = _index128
			(_699_v0).ArraySet1((true) == ((p2).Cmp(p3) > 0), (_index128).Int())
		} else {
			var _700_v1 bool
			_ = _700_v1
			_700_v1 = false
			var _701_v2 _dafny.Set
			_ = _701_v2
			_701_v2 = _dafny.SetOf(_700_v1)
			_700_v1 = (_701_v2).IsSubsetOf(_701_v2)
			var _702_v3 _dafny.Sequence
			_ = _702_v3
			_702_v3 = _dafny.UnicodeSeqOfUtf8Bytes("wrx")
			var _703_v4 _dafny.Map
			_ = _703_v4
			_703_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_700_v1, !(_700_v1))
			var _704_v5 _dafny.Map
			_ = _704_v5
			_704_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_702_v3).Select((Companion_Default___.SafeIndex((_703_v4).Cardinality(), _dafny.IntOfUint32((_702_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), p2)
			var _705_v6 D6
			_ = _705_v6
			_705_v6 = Companion_D6_.Create_DC15_((_704_v5).Merge(_704_v5))
			var _706_v7 _dafny.Array
			_ = _706_v7
			var _nwElement0_22 _dafny.Sequence = _702_v3
			_ = _nwElement0_22
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(8))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_22, 0)
			(_nw117).ArraySet1(_702_v3, 1)
			(_nw117).ArraySet1(_702_v3, 2)
			(_nw117).ArraySet1(_dafny.Companion_Sequence_.Update(_702_v3, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_702_v3).Cardinality()))).Uint32(), _dafny.CodePoint('p')), 3)
			(_nw117).ArraySet1(_702_v3, 4)
			(_nw117).ArraySet1(_702_v3, 5)
			(_nw117).ArraySet1(_702_v3, 6)
			(_nw117).ArraySet1(_702_v3, 7)
			_706_v7 = _nw117
			var _707_v8 _dafny.Array
			_ = _707_v8
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw118
			_707_v8 = _nw118
			var _708_v9 _dafny.MultiSet
			_ = _708_v9
			_708_v9 = _dafny.MultiSetOf(_700_v1, !(false), _700_v1, Companion_Default___.Fm1(_700_v1, _dafny.CodePoint('r'), p2, globalState))
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_707_v8), 0))
			_ = _index129
			(_707_v8).ArraySet1((_708_v9).IsSubsetOf(_708_v9), (_index129).Int())
			var _709_v10 D0
			_ = _709_v10
			_709_v10 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(p1))
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_707_v8), 0))
			_ = _index130
			var _rhs122 D6 = Companion_D6_.Create_DC15_((_704_v5).Merge(_704_v5))
			_ = _rhs122
			var _rhs123 _dafny.Array = _706_v7
			_ = _rhs123
			var _rhs124 bool = Companion_Default___.Fm1(_700_v1, _this.F18(), _dafny.IntOfInt64(-405), globalState)
			_ = _rhs124
			var _rhs125 D0 = Companion_D0_.Create_DC0_(p2)
			_ = _rhs125
			var _rhs126 _dafny.CodePoint = _this.F18()
			_ = _rhs126
			var _lhs92 _dafny.Array = _707_v8
			_ = _lhs92
			var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_707_v8), 0))
			_ = _lhs93
			var _lhs94 *C2 = _this
			_ = _lhs94
			_705_v6 = _rhs122
			_706_v7 = _rhs123
			(_lhs92).ArraySet1(_rhs124, (_lhs93).Int())
			_709_v10 = _rhs125
			_lhs94.F18_set_(_rhs126)
			var _710_v11 D1
			_ = _710_v11
			_710_v11 = Companion_D1_.Create_DC6_(true, p0)
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_707_v8), 0))
			_ = _index131
			(_707_v8).ArraySet1((_710_v11).Dtor_cf7(), (_index131).Int())
			var _711_v12 _dafny.Map
			_ = _711_v12
			_711_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !((_707_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_707_v8), 0))).Int()).(bool)))
			var _712_v13 _dafny.Map
			_ = _712_v13
			_712_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_702_v3).Cardinality()), (func() bool {
				if (_711_v12).Contains(p3) {
					return (_711_v12).Get(p3).(bool)
				}
				return p0
			})())
			_712_v13 = (_712_v13).Update(p1, _700_v1)
			(globalState).F6 = Companion_Default___.Fm0(globalState)
		}
		var _713_v14 bool
		_ = _713_v14
		_713_v14 = true
		var _714_v15 _dafny.Set
		_ = _714_v15
		_714_v15 = _dafny.SetOf(p0)
		var _715_v16 D7
		_ = _715_v16
		_715_v16 = Companion_D7_.Create_DC18_(_714_v15)
		_713_v14 = (_714_v15).Equals((_715_v16).Dtor_cf22())
		var _716_v17 _dafny.Map
		_ = _716_v17
		_716_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_713_v14, p0)
		var _717_v18 _dafny.Map
		_ = _717_v18
		_717_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((((_716_v17).Update(p0, _713_v14)).Merge(_716_v17)).Cardinality(), p2)
		var _718_v19 _dafny.Sequence
		_ = _718_v19
		_718_v19 = _dafny.UnicodeSeqOfUtf8Bytes("oiyohkc")
		var _719_v20 _dafny.Array
		_ = _719_v20
		var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw119
		_719_v20 = _nw119
		var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_719_v20), 0))
		_ = _index132
		(_719_v20).ArraySet1(p1, (_index132).Int())
		var _720_v21 _dafny.MultiSet
		_ = _720_v21
		_720_v21 = _dafny.MultiSetOf(_713_v14, _713_v14)
		var _721_v22 D5
		_ = _721_v22
		_721_v22 = Companion_D5_.Create_DC12_(_720_v21)
		var _pat_let_tv12 = p2
		_ = _pat_let_tv12
		var _pat_let_tv13 = p1
		_ = _pat_let_tv13
		var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_719_v20), 0))
		_ = _index133
		var _rhs127 _dafny.Map = (Companion_Default___.Fm19(p0, p0, globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2))
		_ = _rhs127
		var _rhs128 _dafny.Sequence = _718_v19
		_ = _rhs128
		var _rhs129 _dafny.Int = func(_source9 D5) _dafny.Int {
			if _source9.Is_DC13() {
				return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality())), (_dafny.SeqOf(_pat_let_tv12))))).Cardinality()
			} else if _source9.Is_DC14() {
				var _722___mcc_h0 bool = _source9.Get_().(D5_DC14).Cf16
				_ = _722___mcc_h0
				var _723___mcc_h1 _dafny.Int = _source9.Get_().(D5_DC14).Cf17
				_ = _723___mcc_h1
				var _724_cf17 _dafny.Int = _723___mcc_h1
				_ = _724_cf17
				var _725_cf16 bool = _722___mcc_h0
				_ = _725_cf16
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(378))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}(func(_726_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality()))
			} else {
				var _727___mcc_h2 _dafny.MultiSet = _source9.Get_().(D5_DC12).Cf15
				_ = _727___mcc_h2
				var _728_cf15 _dafny.MultiSet = _727___mcc_h2
				_ = _728_cf15
				return _pat_let_tv13
			}
		}(_721_v22)
		_ = _rhs129
		var _lhs95 _dafny.Array = _719_v20
		_ = _lhs95
		var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_719_v20), 0))
		_ = _lhs96
		_717_v18 = _rhs127
		_718_v19 = _rhs128
		(_lhs95).ArraySet1(_rhs129, (_lhs96).Int())
		var _729_v23 _dafny.Map
		_ = _729_v23
		_729_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_713_v14, p3)
		_729_v23 = (_729_v23).Update(p0, p2)
		var _730_v24 _dafny.MultiSet
		_ = _730_v24
		_730_v24 = _dafny.MultiSetOf(_dafny.IntOfInt64(296))
		var _731_v25 *C1
		_ = _731_v25
		var _nw120 *C1 = New_C1_()
		_ = _nw120
		_nw120.Ctor__(p3, (_730_v24).Contains((_719_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_719_v20), 0))).Int()).(_dafny.Int)), _dafny.CodePoint('y'), ((_this).F19()).Merge((_this).F19()))
		_731_v25 = _nw120
		var _732_v26 _dafny.Sequence
		_ = _732_v26
		_732_v26 = _dafny.SeqOf(_731_v25)
		_731_v25 = (_732_v26).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_732_v26).Cardinality()))).Uint32()).(*C1)
		r0 = p1
		return r0
	}
}
func (_this *C2) M4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Array, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _hi5 _dafny.Int = p1
		_ = _hi5
		for _733_i0 := (p1).Plus(p1); _733_i0.Cmp(_hi5) < 0; _733_i0 = _733_i0.Plus(_dafny.One) {
			var _734_v0 bool
			_ = _734_v0
			_734_v0 = false
			var _735_v1 _dafny.Map
			_ = _735_v1
			_735_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v0, p0)
			var _736_v2 _dafny.MultiSet
			_ = _736_v2
			_736_v2 = _dafny.MultiSetOf(_734_v0)
			var _737_v3 _dafny.Set
			_ = _737_v3
			_737_v3 = _dafny.SetOf((_735_v1).Cardinality(), (_736_v2).Cardinality())
			var _738_v4 _dafny.Map
			_ = _738_v4
			_738_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_737_v3, (p1).Times(p0))
			_738_v4 = (_738_v4).Update(_737_v3, Companion_Default___.Fm0(globalState))
			var _739_v5 _dafny.Sequence
			_ = _739_v5
			_739_v5 = _dafny.UnicodeSeqOfUtf8Bytes("gtyc")
			(globalState).F6 = (p1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_739_v5, _739_v5)).Cardinality()))
			_734_v0 = true
			var _740_v6 _dafny.Map
			_ = _740_v6
			_740_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v5, _734_v0)
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((p2), 0))
			_ = _index134
			(p2).ArraySet1((func() bool {
				if (_740_v6).Contains(_739_v5) {
					return (_740_v6).Get(_739_v5).(bool)
				}
				return _734_v0
			})(), (_index134).Int())
			var _741_v7 D5
			_ = _741_v7
			_741_v7 = Companion_D5_.Create_DC13_()
			var _742_v8 _dafny.Map
			_ = _742_v8
			_742_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_741_v7, _733_i0)
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((p2), 0))
			_ = _index135
			var _rhs130 bool = !(Companion_Default___.Fm2(globalState)).Contains(_734_v0)
			_ = _rhs130
			var _rhs131 _dafny.Map = _742_v8
			_ = _rhs131
			var _rhs132 bool = Companion_Default___.Fm1((_734_v0) || (_734_v0), _dafny.CodePoint('v'), (_dafny.Zero).Minus(p1), globalState)
			_ = _rhs132
			var _lhs97 _dafny.Array = p2
			_ = _lhs97
			var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((p2), 0))
			_ = _lhs98
			(_lhs97).ArraySet1(_rhs130, (_lhs98).Int())
			_742_v8 = _rhs131
			_734_v0 = _rhs132
		}
		var _743_v9 bool
		_ = _743_v9
		_743_v9 = false
		var _744_i1 _dafny.Int
		_ = _744_i1
		_744_i1 = _dafny.Zero
		{
			for _743_v9 {
				{
					if (_744_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_744_i1 = (_744_i1).Plus(_dafny.One)
					(globalState).F6 = (p0).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(p1)))
					var _745_v10 _dafny.Sequence
					_ = _745_v10
					_745_v10 = _dafny.UnicodeSeqOfUtf8Bytes("yesur")
					var _746_v11 _dafny.Map
					_ = _746_v11
					_746_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v9, (_745_v10).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_745_v10).Cardinality()))).Uint32()).(_dafny.CodePoint))
					var _747_v12 _dafny.Sequence
					_ = _747_v12
					_747_v12 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v9, _dafny.CodePoint('c')))
					var _748_v13 _dafny.Map
					_ = _748_v13
					_748_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _745_v10)
					_746_v11 = (_747_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_748_v13).Contains(_dafny.IntOfInt64(-768)) {
							return (_748_v13).Get(_dafny.IntOfInt64(-768)).(_dafny.Sequence)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}(func(_749_i2 _dafny.Int) _dafny.CodePoint {
							return _this.F18()
						}))
					})()).Cardinality()), _dafny.IntOfUint32((_747_v12).Cardinality()))).Uint32()).(_dafny.Map)
					r0 = Companion_Default___.Fm1(_743_v9, _this.F18(), p1, globalState)
					var _750_v14 _dafny.Map
					_ = _750_v14
					_750_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_743_v9), _743_v9)
					var _751_v15 _dafny.Sequence
					_ = _751_v15
					_751_v15 = _dafny.SeqOf(_743_v9)
					var _752_v16 _dafny.Map
					_ = _752_v16
					_752_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-623), _this.F18())
					var _753_v17 _dafny.Map
					_ = _753_v17
					_753_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_751_v15, _752_v16)
					var _754_v19 _dafny.Set
					_ = _754_v19
					_754_v19 = _dafny.SetOf(_751_v15, _751_v15, _751_v15)
					var _755_v20 _dafny.Array
					_ = _755_v20
					var _nwElement0_23 bool = _743_v9
					_ = _nwElement0_23
					var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(21))
					_ = _nw121
					(_nw121).ArraySet1(_nwElement0_23, 0)
					(_nw121).ArraySet1(_743_v9, 1)
					(_nw121).ArraySet1((func() bool {
						if _743_v9 {
							return true
						}
						return !((func() bool {
							if (_750_v14).Contains(_743_v9) {
								return (_750_v14).Get(_743_v9).(bool)
							}
							return _743_v9
						})())
					})(), 2)
					(_nw121).ArraySet1(_743_v9, 3)
					(_nw121).ArraySet1(_743_v9, 4)
					(_nw121).ArraySet1(_dafny.Companion_Sequence_.Contains(_751_v15, _743_v9), 5)
					(_nw121).ArraySet1(_743_v9, 6)
					(_nw121).ArraySet1(true, 7)
					(_nw121).ArraySet1(true, 8)
					(_nw121).ArraySet1(!((func() bool {
						if _743_v9 {
							return _743_v9
						}
						return true
					})()), 9)
					(_nw121).ArraySet1(_743_v9, 10)
					(_nw121).ArraySet1((_751_v15).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)).Cardinality(), _dafny.IntOfUint32((_751_v15).Cardinality()))).Uint32()).(bool), 11)
					(_nw121).ArraySet1(_743_v9, 12)
					(_nw121).ArraySet1(_743_v9, 13)
					(_nw121).ArraySet1(!(_743_v9) || (_743_v9), 14)
					(_nw121).ArraySet1(_743_v9, 15)
					(_nw121).ArraySet1((func() _dafny.Set {
						var _coll33 = _dafny.NewBuilder()
						_ = _coll33
						for _iter34 := _dafny.Iterate((_753_v17).Keys().Elements()); ; {
							_compr_33, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _756_v18 _dafny.Sequence
							_756_v18 = interface{}(_compr_33).(_dafny.Sequence)
							if (_753_v17).Contains(_756_v18) {
								_coll33.Add(_756_v18)
							}
						}
						return _coll33.ToSet()
					}()).IsDisjointFrom(_754_v19), 16)
					(_nw121).ArraySet1(_743_v9, 17)
					(_nw121).ArraySet1(_743_v9, 18)
					(_nw121).ArraySet1(_743_v9, 19)
					(_nw121).ArraySet1(_743_v9, 20)
					_755_v20 = _nw121
					var _757_v21 _dafny.Map
					_ = _757_v21
					_757_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_745_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.IntOfUint32((_745_v10).Cardinality()))).Uint32()).(_dafny.CodePoint), p1)
					var _758_v22 T0
					_ = _758_v22
					var _nw122 *C0 = New_C0_()
					_ = _nw122
					_nw122.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}(func(_759_i3 _dafny.Int) _dafny.CodePoint {
						return _this.F18()
					})), _757_v21, _this.F18(), (_this).F19())
					_758_v22 = _nw122
					var _760_v23 _dafny.Map
					_ = _760_v23
					_760_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v9, _755_v20)
					var _rhs133 _dafny.Array = p2
					_ = _rhs133
					var _rhs134 _dafny.Int = ((func() _dafny.Map {
						if _743_v9 {
							return _760_v23
						}
						return (_760_v23).Merge(_760_v23)
					})()).Cardinality()
					_ = _rhs134
					var _rhs135 T0 = (func() T0 {
						if _743_v9 {
							return _758_v22
						}
						return _758_v22
					})()
					_ = _rhs135
					var _rhs136 _dafny.CodePoint = Companion_Default___.Fm20(_743_v9, globalState)
					_ = _rhs136
					var _lhs99 *GlobalState = globalState
					_ = _lhs99
					var _lhs100 *C2 = _this
					_ = _lhs100
					_755_v20 = _rhs133
					_lhs99.F6 = _rhs134
					_758_v22 = _rhs135
					_lhs100.F18_set_(_rhs136)
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _761_v24 _dafny.Array
		_ = _761_v24
		var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw123
		_761_v24 = _nw123
		var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_761_v24), 0))
		_ = _index136
		(_761_v24).ArraySet1(p2, (_index136).Int())
		var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_761_v24), 0))
		_ = _index137
		(_761_v24).ArraySet1(p2, (_index137).Int())
		var _762_v25 _dafny.MultiSet
		_ = _762_v25
		_762_v25 = _dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-174), p1))
		var _763_v26 _dafny.Map
		_ = _763_v26
		_763_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v9, p0)
		var _764_v27 _dafny.Sequence
		_ = _764_v27
		_764_v27 = _dafny.UnicodeSeqOfUtf8Bytes("jcyiogjr")
		var _765_v28 _dafny.Set
		_ = _765_v28
		_765_v28 = _dafny.SetOf((_763_v26).Cardinality(), _dafny.IntOfUint32((_764_v27).Cardinality()))
		var _766_v29 _dafny.Sequence
		_ = _766_v29
		_766_v29 = _dafny.SeqOf(p0)
		_762_v25 = (_dafny.MultiSetOf(_765_v28)).Update((_765_v28).Union(func() _dafny.Set {
			var _coll34 = _dafny.NewBuilder()
			_ = _coll34
			for _iter35 := _dafny.Iterate((_766_v29).Elements()); ; {
				_compr_34, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _767_v30 _dafny.Int
				_767_v30 = interface{}(_compr_34).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_766_v29, _767_v30) {
					_coll34.Add(Companion_Default___.SafeDivisionInt(_767_v30, (func() _dafny.Set {
						var _coll35 = _dafny.NewBuilder()
						_ = _coll35
						for _iter36 := _dafny.Iterate((func() _dafny.Map {
							var _coll36 = _dafny.NewMapBuilder()
							_ = _coll36
							for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(814), _dafny.IntOfInt64(478))); ; {
								_compr_36, _ok37 := _iter37()
								if !_ok37 {
									break
								}
								var _768_v31 _dafny.Int
								_768_v31 = interface{}(_compr_36).(_dafny.Int)
								if ((_dafny.IntOfInt64(814)).Cmp(_768_v31) <= 0) && ((_768_v31).Cmp(_dafny.IntOfInt64(478)) < 0) {
									_coll36.Add((_768_v31).Plus(_dafny.IntOfInt64(-657)), _dafny.IntOfInt64(846))
								}
							}
							return _coll36.ToMap()
						}()).Keys().Elements()); ; {
							_compr_35, _ok36 := _iter36()
							if !_ok36 {
								break
							}
							var _769_v32 _dafny.Int
							_769_v32 = interface{}(_compr_35).(_dafny.Int)
							if (func() _dafny.Map {
								var _coll37 = _dafny.NewMapBuilder()
								_ = _coll37
								for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(814), _dafny.IntOfInt64(478))); ; {
									_compr_37, _ok38 := _iter38()
									if !_ok38 {
										break
									}
									var _768_v31 _dafny.Int
									_768_v31 = interface{}(_compr_37).(_dafny.Int)
									if ((_dafny.IntOfInt64(814)).Cmp(_768_v31) <= 0) && ((_768_v31).Cmp(_dafny.IntOfInt64(478)) < 0) {
										_coll37.Add((_768_v31).Plus(_dafny.IntOfInt64(-657)), _dafny.IntOfInt64(846))
									}
								}
								return _coll37.ToMap()
							}()).Contains(_769_v32) {
								_coll35.Add((_769_v32).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(883))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg46 _dafny.Int) interface{} {
										return coer46(arg46)
									}
								}(func(_770_i4 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('t')
								}))).Cardinality())))
							}
						}
						return _coll35.ToSet()
					}()).Cardinality()))
				}
			}
			return _coll34.ToSet()
		}()), Companion_Default___.Abs(Companion_Default___.Fm0(globalState)))
		var _771_i5 _dafny.Int
		_ = _771_i5
		_771_i5 = _dafny.Zero
		{
			for (_765_v28).IsDisjointFrom(_765_v28) {
				{
					if (_771_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_771_i5 = (_771_i5).Plus(_dafny.One)
					(globalState).F6 = p0
					(globalState).F6 = p0
					var _772_v33 _dafny.Map
					_ = _772_v33
					_772_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), p0)
					var _773_v34 T0
					_ = _773_v34
					var _nw124 *C0 = New_C0_()
					_ = _nw124
					_nw124.Ctor__(_764_v27, _772_v33, _this.F18(), (_this).F19())
					_773_v34 = _nw124
					var _774_v35 T0
					_ = _774_v35
					_774_v35 = _773_v34
					var _source10 T0 = _774_v35
					_ = _source10
					var _775___mcc_h0 T0 = _source10
					_ = _775___mcc_h0
					var _776_cf14 T0 = _775___mcc_h0
					_ = _776_cf14
					var _777_v36 _dafny.Sequence
					_ = _777_v36
					_777_v36 = _dafny.SeqOf(_743_v9)
					_764_v27 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm21(globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm21(globalState)).Cardinality()))).Uint32(), _this.F18()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm21(globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm21(globalState)).Cardinality()))).Uint32(), _this.F18())).Cardinality()))).Uint32(), _776_cf14.F18()), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_764_v27, _764_v27), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_777_v36).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_764_v27, _764_v27)).Cardinality()))).Uint32(), _this.F18()))
					var _778_v37 _dafny.Sequence
					_ = _778_v37
					_778_v37 = _dafny.SeqOf(_763_v26)
					var _779_v38 _dafny.Map
					_ = _779_v38
					_779_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm22(globalState))
					var _780_v39 _dafny.Sequence
					_ = _780_v39
					_780_v39 = _dafny.SeqOf((func() _dafny.Sequence {
						if (_779_v38).Contains(p1) {
							return (_779_v38).Get(p1).(_dafny.Sequence)
						}
						return _778_v37
					})())
					var _pat_let_tv14 = _743_v9
					_ = _pat_let_tv14
					var _pat_let_tv15 = p0
					_ = _pat_let_tv15
					_778_v37 = (_780_v39).Select((Companion_Default___.SafeIndex((func(_pat_let14_0 D5) D5 {
						return func(_783_dt__update__tmp_h1 D5) D5 {
							return func(_pat_let17_0 _dafny.Int) D5 {
								return func(_784_dt__update_hcf17_h0 _dafny.Int) D5 {
									return Companion_D5_.Create_DC14_((_783_dt__update__tmp_h1).Dtor_cf16(), _784_dt__update_hcf17_h0)
								}(_pat_let17_0)
							}(_pat_let_tv15)
						}(_pat_let14_0)
					}(func(_pat_let15_0 D5) D5 {
						return func(_781_dt__update__tmp_h0 D5) D5 {
							return func(_pat_let16_0 bool) D5 {
								return func(_782_dt__update_hcf16_h0 bool) D5 {
									return Companion_D5_.Create_DC14_(_782_dt__update_hcf16_h0, (_781_dt__update__tmp_h0).Dtor_cf17())
								}(_pat_let16_0)
							}(!(_pat_let_tv14))
						}(_pat_let15_0)
					}(Companion_D5_.Create_DC14_(_743_v9, p1)))).Dtor_cf17(), _dafny.IntOfUint32((_780_v39).Cardinality()))).Uint32()).(_dafny.Sequence)
					var _785_v40 _dafny.Map
					_ = _785_v40
					_785_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_764_v27, _743_v9)
					_785_v40 = (_785_v40).Update(_764_v27, _743_v9)
					(globalState).F6 = (_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))
					var _786_v41 _dafny.Sequence
					_ = _786_v41
					_786_v41 = _766_v29
					_786_v41 = _786_v41
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _787_v42 _dafny.Array
		_ = _787_v42
		var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw125
		_787_v42 = _nw125
		for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_787_v42), 0))); ; {
			_guard_loop_1, _ok39 := _iter39()
			if !_ok39 {
				break
			}
			var _788_i6 _dafny.Int
			_788_i6 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_788_i6).Sign() != -1) && ((_788_i6).Cmp(_dafny.ArrayLen((_787_v42), 0)) < 0)) {
				(_787_v42).ArraySet1((_788_i6).Plus(Companion_Default___.SafeModuloInt(p0, p1)), (_788_i6).Int())
			}
		}
		r0 = !(_743_v9)
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f18 _dafny.CodePoint
	_f19 _dafny.Map
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f18 = _dafny.CodePoint('D')
	_this._f19 = _dafny.EmptyMap
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

func (_this *C3) F18() _dafny.CodePoint {
	return _this._f18
}
func (_this *C3) F18_set_(value _dafny.CodePoint) {
	_this._f18 = value
}
func (_this *C3) F19() _dafny.Map {
	return _this._f19
}
func (_this *C3) Ctor__(f18 _dafny.CodePoint, f19 _dafny.Map) {
	{
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C3) Fm3(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg47 _dafny.Int) interface{} {
				return coer47(arg47)
			}
		}(func(_789_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(937), _dafny.IntOfInt64(678))).Cardinality())
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}(func(_790_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(307)
		}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(647))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}(func(_791_i2 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-125)
		})))
	}
}
func (_this *C3) Fm4(p0 bool, p1 _dafny.CodePoint, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(((_dafny.MultiSetOf(_dafny.IntOfInt64(265))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(373), _dafny.IntOfInt64(858)))).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hkdhj"), _dafny.UnicodeSeqOfUtf8Bytes("wtgkshg"))).Cardinality()))
	}
}
func (_this *C3) M1(p0 bool, globalState *GlobalState) (D0, _dafny.Int, bool, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _792_i0 _dafny.Int
		_ = _792_i0
		_792_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_792_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_792_i0 = (_792_i0).Plus(_dafny.One)
					var _793_v0 _dafny.Int
					_ = _793_v0
					_793_v0 = _dafny.IntOfInt64(652)
					var _794_v1 _dafny.MultiSet
					_ = _794_v1
					_794_v1 = _dafny.MultiSetOf(_793_v0)
					var _795_v2 _dafny.Sequence
					_ = _795_v2
					_795_v2 = _dafny.SeqOf(_this.F18())
					var _796_v3 _dafny.Set
					_ = _796_v3
					_796_v3 = _dafny.SetOf((_dafny.Zero).Minus(_793_v0), _793_v0)
					var _797_v4 _dafny.Sequence
					_ = _797_v4
					_797_v4 = _dafny.SeqOf(false, p0)
					var _798_v5 _dafny.Sequence
					_ = _798_v5
					_798_v5 = _797_v4
					var _799_v6 _dafny.Map
					_ = _799_v6
					_799_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v5), _793_v0)
					var _800_v7 _dafny.Map
					_ = _800_v7
					_800_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_799_v6).Cardinality(), p0)
					var _801_v8 _dafny.Map
					_ = _801_v8
					_801_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_800_v7).Cardinality(), p0)
					var _802_v9 _dafny.Array
					_ = _802_v9
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_13
					var _nw126 _dafny.Array
					_ = _nw126
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw126 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) _dafny.Int = (func(_803_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_804_i1 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_804_i1, _803_v0)
							}
						})(_793_v0)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw126 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw126).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw126).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_802_v9 = _nw126
					var _805_v10 _dafny.Map
					_ = _805_v10
					_805_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v9, (_this).Fm3(p0, _793_v0, globalState))
					var _806_v11 _dafny.Array
					_ = _806_v11
					var _nwElement0_24 bool = (p0) || (p0)
					_ = _nwElement0_24
					var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(23))
					_ = _nw127
					(_nw127).ArraySet1(_nwElement0_24, 0)
					(_nw127).ArraySet1((_794_v1).IsSubsetOf(_794_v1), 1)
					(_nw127).ArraySet1(false, 2)
					(_nw127).ArraySet1((Companion_D1_.Create_DC4_(_793_v0, _793_v0, true)).Dtor_cf5(), 3)
					(_nw127).ArraySet1(_dafny.Companion_Sequence_.Equal(_795_v2, _795_v2), 4)
					(_nw127).ArraySet1(false, 5)
					(_nw127).ArraySet1(p0, 6)
					(_nw127).ArraySet1(p0, 7)
					(_nw127).ArraySet1((p0) || (true), 8)
					(_nw127).ArraySet1((func() bool {
						if p0 {
							return p0
						}
						return p0
					})(), 9)
					(_nw127).ArraySet1(p0, 10)
					(_nw127).ArraySet1(p0, 11)
					(_nw127).ArraySet1(!(p0), 12)
					(_nw127).ArraySet1(!((Companion_Default___.Fm5(_793_v0, globalState)).IsDisjointFrom(_796_v3)), 13)
					(_nw127).ArraySet1((_dafny.IntOfUint32((_795_v2).Cardinality())).Cmp((_800_v7).Cardinality()) > 0, 14)
					(_nw127).ArraySet1((_793_v0).Cmp(_793_v0) < 0, 15)
					(_nw127).ArraySet1(p0, 16)
					(_nw127).ArraySet1(p0, 17)
					(_nw127).ArraySet1((func() bool {
						if (_805_v10).Contains(_802_v9) {
							return (_805_v10).Get(_802_v9).(bool)
						}
						return p0
					})(), 18)
					(_nw127).ArraySet1((_793_v0).Cmp(_793_v0) > 0, 19)
					(_nw127).ArraySet1((_794_v1).IsDisjointFrom(_794_v1), 20)
					(_nw127).ArraySet1(p0, 21)
					(_nw127).ArraySet1((p0) || (p0), 22)
					_806_v11 = _nw127
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_806_v11), 0))
					_ = _index138
					(_806_v11).ArraySet1(p0, (_index138).Int())
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_802_v9), 0))
					_ = _index139
					(_802_v9).ArraySet1(_793_v0, (_index139).Int())
					var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_806_v11), 0))
					_ = _index140
					var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_802_v9), 0))
					_ = _index141
					var _rhs137 bool = (_793_v0).Cmp(_793_v0) > 0
					_ = _rhs137
					var _rhs138 _dafny.Int = (_793_v0).Times(_793_v0)
					_ = _rhs138
					var _rhs139 bool = p0
					_ = _rhs139
					var _rhs140 _dafny.Int = _dafny.IntOfInt64(620)
					_ = _rhs140
					var _lhs101 *GlobalState = globalState
					_ = _lhs101
					var _lhs102 _dafny.Array = _806_v11
					_ = _lhs102
					var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_806_v11), 0))
					_ = _lhs103
					var _lhs104 _dafny.Array = _802_v9
					_ = _lhs104
					var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_802_v9), 0))
					_ = _lhs105
					r3 = _rhs137
					_lhs101.F6 = _rhs138
					(_lhs102).ArraySet1(_rhs139, (_lhs103).Int())
					(_lhs104).ArraySet1(_rhs140, (_lhs105).Int())
					var _807_v12 _dafny.Map
					_ = _807_v12
					_807_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_795_v2).Cardinality()), _793_v0)
					_800_v7 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_793_v0, p0)).Update((_dafny.Zero).Minus(((_802_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_802_v9), 0))).Int()).(_dafny.Int)).Times((_807_v12).Cardinality())), (_806_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_806_v11), 0))).Int()).(bool))
					var _808_v13 *C0
					_ = _808_v13
					var _nw128 *C0 = New_C0_()
					_ = _nw128
					_nw128.Ctor__(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if (Companion_D1_.Create_DC3_(p0)).Dtor_cf2() {
							return _795_v2
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("as")
					})(), (Companion_Default___.SafeIndex((_802_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_802_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (Companion_D1_.Create_DC3_(p0)).Dtor_cf2() {
							return _795_v2
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("as")
					})()).Cardinality()))).Uint32(), _this.F18()), Companion_Default___.Fm7((_806_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_806_v11), 0))).Int()).(bool), _793_v0, globalState), _this.F18(), ((_this).F19()).Merge((_this).F19()))
					_808_v13 = _nw128
					var _809_v14 *C0
					_ = _809_v14
					var _nw129 *C0 = New_C0_()
					_ = _nw129
					_nw129.Ctor__((_808_v13).F20(), (_808_v13).F21(), _this.F18(), (_this).F19())
					_809_v14 = _nw129
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _810_v15 _dafny.Int
		_ = _810_v15
		_810_v15 = _dafny.IntOfInt64(854)
		var _811_v16 _dafny.MultiSet
		_ = _811_v16
		_811_v16 = _dafny.MultiSetOf(_810_v15)
		var _812_v17 _dafny.Array
		_ = _812_v17
		var _nwElement0_25 _dafny.Int = _810_v15
		_ = _nwElement0_25
		var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(21))
		_ = _nw130
		(_nw130).ArraySet1(_nwElement0_25, 0)
		(_nw130).ArraySet1(_810_v15, 1)
		(_nw130).ArraySet1(_810_v15, 2)
		(_nw130).ArraySet1(_810_v15, 3)
		(_nw130).ArraySet1(_dafny.IntOfInt64(200), 4)
		(_nw130).ArraySet1(_810_v15, 5)
		(_nw130).ArraySet1(_810_v15, 6)
		(_nw130).ArraySet1(_810_v15, 7)
		(_nw130).ArraySet1((_dafny.Zero).Minus(_810_v15), 8)
		(_nw130).ArraySet1(_dafny.IntOfInt64(497), 9)
		(_nw130).ArraySet1(_810_v15, 10)
		(_nw130).ArraySet1(_810_v15, 11)
		(_nw130).ArraySet1(_810_v15, 12)
		(_nw130).ArraySet1((_811_v16).Cardinality(), 13)
		(_nw130).ArraySet1(_810_v15, 14)
		(_nw130).ArraySet1(_810_v15, 15)
		(_nw130).ArraySet1(_810_v15, 16)
		(_nw130).ArraySet1(_810_v15, 17)
		(_nw130).ArraySet1(_810_v15, 18)
		(_nw130).ArraySet1(_dafny.IntOfInt64(192), 19)
		(_nw130).ArraySet1(_dafny.IntOfInt64(-38), 20)
		_812_v17 = _nw130
		var _813_v18 _dafny.Map
		_ = _813_v18
		_813_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_812_v17, _this.F18())
		var _814_v19 _dafny.Sequence
		_ = _814_v19
		_814_v19 = _dafny.UnicodeSeqOfUtf8Bytes("kbgkabcen")
		var _hi6 _dafny.Int = (_dafny.IntOfUint32((_814_v19).Cardinality())).Minus(_810_v15)
		_ = _hi6
		for _815_i2 := (_813_v18).Cardinality(); _815_i2.Cmp(_hi6) < 0; _815_i2 = _815_i2.Plus(_dafny.One) {
			var _816_v20 D3
			_ = _816_v20
			_816_v20 = Companion_D3_.Create_DC8_(_this.F18())
			var _817_v21 _dafny.Map
			_ = _817_v21
			_817_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _818_v22 _dafny.Sequence
			_ = _818_v22
			_818_v22 = _dafny.SeqOf(p0)
			var _819_v23 _dafny.Map
			_ = _819_v23
			_819_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_817_v21, _dafny.IntOfUint32((_818_v22).Cardinality()))
			var _820_v24 _dafny.Sequence
			_ = _820_v24
			_820_v24 = _dafny.SeqOf(p0, !(p0) || (Companion_Default___.Fm1(p0, (_816_v20).Dtor_cf10(), (func() _dafny.Int {
				if (_819_v23).Contains(_817_v21) {
					return (_819_v23).Get(_817_v21).(_dafny.Int)
				}
				return _815_i2
			})(), globalState)), p0)
			r1 = _dafny.IntOfUint32((_818_v22).Cardinality())
			(globalState).F6 = _810_v15
			(globalState).F6 = _810_v15
			r1 = (func() _dafny.Int {
				if !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("pqpigsw"), _dafny.UnicodeSeqOfUtf8Bytes("mlp")) {
					return _810_v15
				}
				return _810_v15
			})()
		}
		var _hi7 _dafny.Int = _810_v15
		_ = _hi7
		for _821_i3 := _810_v15; _821_i3.Cmp(_hi7) < 0; _821_i3 = _821_i3.Plus(_dafny.One) {
			var _822_v25 _dafny.Array
			_ = _822_v25
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_14
			var _nw131 _dafny.Array
			_ = _nw131
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw131 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Sequence = (func(_823_v19 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_824_i4 _dafny.Int) _dafny.Sequence {
						return _823_v19
					}
				})(_814_v19)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw131 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw131).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw131).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_822_v25 = _nw131
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_822_v25), 0))
			_ = _index142
			(_822_v25).ArraySet1(_814_v19, (_index142).Int())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_822_v25), 0))
			_ = _index143
			(_822_v25).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_814_v19, _814_v19), (_index143).Int())
			_812_v17 = (func() _dafny.Array {
				if false {
					return _812_v17
				}
				return _812_v17
			})()
			var _825_v26 _dafny.Map
			_ = _825_v26
			_825_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_821_i3, p0)
			var _826_v27 _dafny.Set
			_ = _826_v27
			_826_v27 = _dafny.SetOf(p0)
			var _827_v28 _dafny.Sequence
			_ = _827_v28
			_827_v28 = _dafny.SeqOf((func() bool {
				if (_825_v26).Contains((_811_v16).Cardinality()) {
					return (_825_v26).Get((_811_v16).Cardinality()).(bool)
				}
				return p0
			})(), p0, p0, Companion_Default___.Fm1(p0, _this.F18(), (_826_v27).Cardinality(), globalState), p0)
			var _828_v29 _dafny.Sequence
			_ = _828_v29
			_828_v29 = _dafny.SeqOf(true, (_827_v28).Select((Companion_Default___.SafeIndex(_810_v15, _dafny.IntOfUint32((_827_v28).Cardinality()))).Uint32()).(bool))
			var _829_v30 _dafny.Map
			_ = _829_v30
			_829_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.IntOfUint32((_814_v19).Cardinality()))
			var _830_v31 T0
			_ = _830_v31
			var _nw132 *C0 = New_C0_()
			_ = _nw132
			_nw132.Ctor__((_822_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_822_v25), 0))).Int()).(_dafny.Sequence), _829_v30, _this.F18(), (_this).F19())
			_830_v31 = _nw132
			var _831_v32 _dafny.Set
			_ = _831_v32
			_831_v32 = _dafny.SetOf(_830_v31, _830_v31, _830_v31)
			var _rhs141 _dafny.Int = (func() _dafny.Int {
				if (_827_v28).Select((Companion_Default___.SafeIndex(_810_v15, _dafny.IntOfUint32((_827_v28).Cardinality()))).Uint32()).(bool) {
					return _810_v15
				}
				return _821_i3
			})()
			_ = _rhs141
			var _rhs142 bool = !(((_831_v32).Union(_831_v32)).IsProperSubsetOf(_831_v32))
			_ = _rhs142
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			_lhs106.F6 = _rhs141
			r3 = _rhs142
			(globalState).F6 = _821_i3
		}
		var _832_v33 _dafny.Array
		_ = _832_v33
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_15
		var _nw133 _dafny.Array
		_ = _nw133
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw133 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = (func(_833_p0 bool) func(_dafny.Int) bool {
				return func(_834_i5 _dafny.Int) bool {
					return _833_p0
				}
			})(p0)
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw133 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw133).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw133).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_832_v33 = _nw133
		_832_v33 = _832_v33
		var _835_v34 D1
		_ = _835_v34
		_835_v34 = Companion_D1_.Create_DC3_(false)
		r2 = (!((_835_v34).Dtor_cf2())) && (p0)
		var _836_v35 *C0
		_ = _836_v35
		var _nw134 *C0 = New_C0_()
		_ = _nw134
		_nw134.Ctor__((func() _dafny.Sequence {
			if p0 {
				return _814_v19
			}
			return _814_v19
		})(), Companion_Default___.Fm7(p0, _810_v15, globalState), _dafny.CodePoint('g'), (_this).F19())
		_836_v35 = _nw134
		var _837_v36 D0
		_ = _837_v36
		_837_v36 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(_810_v15))
		r0 = _837_v36
		var _838_v37 _dafny.Sequence
		_ = _838_v37
		_838_v37 = _dafny.SeqOf(_810_v15, _810_v15, _810_v15)
		r1 = (_810_v15).Times((_dafny.IntOfUint32((_838_v37).Cardinality())).Minus(_810_v15))
		r2 = p0
		r3 = p0
		return r0, r1, r2, r3
	}
}
func (_this *C3) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _839_v0 _dafny.Array
		_ = _839_v0
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_16
		var _nw135 _dafny.Array
		_ = _nw135
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw135 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = (func(_840_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_841_i0 _dafny.Int) _dafny.Int {
					return (_841_i0).Plus(_840_p3)
				}
			})(p3)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw135 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw135).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw135).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_839_v0 = _nw135
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
		_ = _index144
		(_839_v0).ArraySet1(p3, (_index144).Int())
		var _842_v1 bool
		_ = _842_v1
		_842_v1 = true
		var _843_v2 _dafny.Set
		_ = _843_v2
		_843_v2 = _dafny.SetOf(_842_v1)
		var _844_v3 _dafny.Map
		_ = _844_v3
		_844_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _843_v2)
		var _845_v4 D3
		_ = _845_v4
		_845_v4 = Companion_D3_.Create_DC9_(_842_v1, Companion_Default___.Fm0(globalState))
		var _pat_let_tv16 = p2
		_ = _pat_let_tv16
		var _pat_let_tv17 = p2
		_ = _pat_let_tv17
		var _pat_let_tv18 = p3
		_ = _pat_let_tv18
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
		_ = _index145
		var _rhs143 _dafny.Int = (_dafny.Zero).Minus(func(_source11 D3) _dafny.Int {
			if _source11.Is_DC9() {
				var _846___mcc_h0 bool = _source11.Get_().(D3_DC9).Cf11
				_ = _846___mcc_h0
				var _847___mcc_h1 _dafny.Int = _source11.Get_().(D3_DC9).Cf12
				_ = _847___mcc_h1
				var _848_cf12 _dafny.Int = _847___mcc_h1
				_ = _848_cf12
				var _849_cf11 bool = _846___mcc_h0
				_ = _849_cf11
				return _848_cf12
			} else if _source11.Is_DC10() {
				var _850___mcc_h2 bool = _source11.Get_().(D3_DC10).Cf13
				_ = _850___mcc_h2
				var _851_cf13 bool = _850___mcc_h2
				_ = _851_cf13
				return Companion_Default___.SafeModuloInt(_pat_let_tv16, _pat_let_tv17)
			} else {
				var _852___mcc_h3 _dafny.CodePoint = _source11.Get_().(D3_DC8).Cf10
				_ = _852___mcc_h3
				var _853_cf10 _dafny.CodePoint = _852___mcc_h3
				_ = _853_cf10
				return _pat_let_tv18
			}
		}(_845_v4))
		_ = _rhs143
		var _rhs144 bool = false
		_ = _rhs144
		var _rhs145 _dafny.Map = _844_v3
		_ = _rhs145
		var _lhs107 _dafny.Array = _839_v0
		_ = _lhs107
		var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
		_ = _lhs108
		(_lhs107).ArraySet1(_rhs143, (_lhs108).Int())
		_842_v1 = _rhs144
		_844_v3 = _rhs145
		var _854_i1 _dafny.Int
		_ = _854_i1
		_854_i1 = _dafny.Zero
		{
			for _842_v1 {
				{
					if (_854_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_854_i1 = (_854_i1).Plus(_dafny.One)
					(globalState).F6 = ((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int)).Times(p3)
					var _855_v5 _dafny.Sequence
					_ = _855_v5
					_855_v5 = _dafny.UnicodeSeqOfUtf8Bytes("qxtvrh")
					var _856_v6 _dafny.Map
					_ = _856_v6
					_856_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), _dafny.IntOfInt64(308))
					var _857_v7 *C0
					_ = _857_v7
					var _nw136 *C0 = New_C0_()
					_ = _nw136
					_nw136.Ctor__(_855_v5, _856_v6, _this.F18(), ((_this).F19()).Update(_855_v5, (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int)))
					_857_v7 = _nw136
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
					_ = _index146
					(_839_v0).ArraySet1(Companion_Default___.Fm0(globalState), (_index146).Int())
					if ((p3).Minus(p2)).Cmp((_dafny.Zero).Minus((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int))) <= 0 {
						var _858_v8 _dafny.Sequence
						_ = _858_v8
						_858_v8 = _dafny.SeqOf((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int))
						(globalState).F1 = _dafny.Companion_Sequence_.Concatenate(_858_v8, _dafny.SeqOf(p2))
						var _859_v9 _dafny.Map
						_ = _859_v9
						_859_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p2, p1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0))
						_859_v9 = (_859_v9).Update((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), p0))
						var _860_v10 D1
						_ = _860_v10
						_860_v10 = Companion_D1_.Create_DC3_(true)
						var _861_v11 _dafny.Map
						_ = _861_v11
						_861_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _860_v10)
						var _862_v12 T0
						_ = _862_v12
						var _nw137 *C0 = New_C0_()
						_ = _nw137
						_nw137.Ctor__(Companion_Default___.Fm8(globalState), (_857_v7).F21(), _this.F18(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_857_v7).F20(), (_861_v11).Cardinality())).Update(_dafny.UnicodeSeqOfUtf8Bytes("vnfqve"), (_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))))
						_862_v12 = _nw137
						var _863_v13 _dafny.Array
						_ = _863_v13
						var _len0_17 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_17
						var _nw138 _dafny.Array
						_ = _nw138
						if _len0_17.Cmp(_dafny.Zero) == 0 {
							_nw138 = _dafny.NewArray(_len0_17)
						} else {
							var _init17 func(_dafny.Int) bool = (func(_864_p0 bool) func(_dafny.Int) bool {
								return func(_865_i2 _dafny.Int) bool {
									return _864_p0
								}
							})(p0)
							_ = _init17
							var _element0_17 = _init17(_dafny.Zero)
							_ = _element0_17
							_nw138 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
							(_nw138).ArraySet1(_element0_17, 0)
							var _nativeLen0_17 = (_len0_17).Int()
							_ = _nativeLen0_17
							for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
								(_nw138).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
							}
						}
						_863_v13 = _nw138
						var _866_v14 _dafny.Map
						_ = _866_v14
						_866_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_862_v12, _863_v13)
						_866_v14 = (_866_v14).Update(_862_v12, _863_v13)
						var _867_v15 _dafny.Sequence
						_ = _867_v15
						_867_v15 = _dafny.SeqOf(true, _842_v1)
						var _868_v16 _dafny.Map
						_ = _868_v16
						_868_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
						var _869_v17 _dafny.Sequence
						_ = _869_v17
						_869_v17 = _dafny.Companion_Sequence_.Update(_867_v15, (Companion_Default___.SafeIndex((_868_v16).Cardinality(), _dafny.IntOfUint32((_867_v15).Cardinality()))).Uint32(), p0)
						var _870_v18 _dafny.MultiSet
						_ = _870_v18
						_870_v18 = _dafny.MultiSetOf(_869_v17)
						var _871_v19 _dafny.Sequence
						_ = _871_v19
						_871_v19 = _dafny.SeqOf(_867_v15)
						var _872_v20 _dafny.Sequence
						_ = _872_v20
						_872_v20 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_871_v19))
						var _873_v21 _dafny.Array
						_ = _873_v21
						var _nwElement0_26 _dafny.MultiSet = (_870_v18).Intersection(_870_v18)
						_ = _nwElement0_26
						var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(11))
						_ = _nw139
						(_nw139).ArraySet1(_nwElement0_26, 0)
						(_nw139).ArraySet1(_dafny.MultiSetOf(_869_v17, _869_v17, _869_v17, _867_v15, _dafny.SeqOf(p0, p0, false, _842_v1, p0)), 1)
						(_nw139).ArraySet1(_870_v18, 2)
						(_nw139).ArraySet1((_870_v18).Difference(_870_v18), 3)
						(_nw139).ArraySet1(_dafny.MultiSetOf(_869_v17, _869_v17), 4)
						(_nw139).ArraySet1(Companion_Default___.Fm9(globalState), 5)
						(_nw139).ArraySet1(_870_v18, 6)
						(_nw139).ArraySet1((_872_v20).Select((Companion_Default___.SafeIndex((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_872_v20).Cardinality()))).Uint32()).(_dafny.MultiSet), 7)
						(_nw139).ArraySet1((_872_v20).Select((Companion_Default___.SafeIndex((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_872_v20).Cardinality()))).Uint32()).(_dafny.MultiSet), 8)
						(_nw139).ArraySet1(_870_v18, 9)
						(_nw139).ArraySet1((_dafny.MultiSetOf(_869_v17)).Update(_869_v17, Companion_Default___.Abs((_this).Fm4(_842_v1, _this.F18(), _dafny.CodePoint('u'), _842_v1, globalState))), 10)
						_873_v21 = _nw139
						var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_873_v21), 0))
						_ = _index147
						(_873_v21).ArraySet1(_870_v18, (_index147).Int())
						var _874_v22 _dafny.MultiSet
						_ = _874_v22
						_874_v22 = _dafny.MultiSetOf(_842_v1)
						var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_873_v21), 0))
						_ = _index148
						(_873_v21).ArraySet1(((_870_v18).Update(_869_v17, Companion_Default___.Abs((_874_v22).Cardinality()))).Intersection(_870_v18), (_index148).Int())
						var _875_v23 _dafny.Set
						_ = _875_v23
						_875_v23 = _dafny.SetOf(_dafny.IntOfInt64(-192), (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int))
						_842_v1 = (_875_v23).IsSubsetOf((_875_v23).Union(_875_v23))
					} else {
						var _876_v24 _dafny.Map
						_ = _876_v24
						_876_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(globalState), (_857_v7).F20()), _842_v1)
						_842_v1 = (func() bool {
							if (_876_v24).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fvox"), (_857_v7).F20())) {
								return (_876_v24).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fvox"), (_857_v7).F20())).(bool)
							}
							return _842_v1
						})()
						var _877_v25 _dafny.Array
						_ = _877_v25
						var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
						_ = _nw140
						_877_v25 = _nw140
						var _878_v26 _dafny.Set
						_ = _878_v26
						_878_v26 = _dafny.SetOf(p3)
						var _879_v27 _dafny.Sequence
						_ = _879_v27
						_879_v27 = _dafny.SeqOf(p1, (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), p1, (_878_v26).Cardinality(), p1)
						var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_877_v25), 0))
						_ = _index149
						(_877_v25).ArraySet1(_879_v27, (_index149).Int())
						var _880_v28 _dafny.Array
						_ = _880_v28
						var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
						_ = _len0_18
						var _nw141 _dafny.Array
						_ = _nw141
						if _len0_18.Cmp(_dafny.Zero) == 0 {
							_nw141 = _dafny.NewArray(_len0_18)
						} else {
							var _init18 func(_dafny.Int) _dafny.Set = (func(_881_v2 _dafny.Set) func(_dafny.Int) _dafny.Set {
								return func(_882_i3 _dafny.Int) _dafny.Set {
									return _881_v2
								}
							})(_843_v2)
							_ = _init18
							var _element0_18 = _init18(_dafny.Zero)
							_ = _element0_18
							_nw141 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
							(_nw141).ArraySet1(_element0_18, 0)
							var _nativeLen0_18 = (_len0_18).Int()
							_ = _nativeLen0_18
							for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
								(_nw141).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
							}
						}
						_880_v28 = _nw141
						var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_880_v28), 0))
						_ = _index150
						(_880_v28).ArraySet1(_dafny.SetOf(_842_v1, _842_v1, _842_v1, p0, !(_842_v1)), (_index150).Int())
						var _883_v29 _dafny.Array
						_ = _883_v29
						var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
						_ = _nw142
						_883_v29 = _nw142
						var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_883_v29), 0))
						_ = _index151
						(_883_v29).ArraySet1(_842_v1, (_index151).Int())
						var _884_v30 _dafny.Sequence
						_ = _884_v30
						_884_v30 = _dafny.SeqOf(!(_842_v1))
						var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_877_v25), 0))
						_ = _index152
						var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_880_v28), 0))
						_ = _index153
						var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_883_v29), 0))
						_ = _index154
						var _rhs146 bool = (_884_v30).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_884_v30).Cardinality()))).Uint32()).(bool)
						_ = _rhs146
						var _rhs147 D3 = Companion_D3_.Create_DC9_(p0, p3)
						_ = _rhs147
						var _rhs148 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_879_v27, _879_v27), _879_v27)
						_ = _rhs148
						var _rhs149 _dafny.Set = ((_843_v2).Union(_843_v2)).Intersection(_dafny.SetOf(p0))
						_ = _rhs149
						var _rhs150 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
							if p0 {
								return _879_v27
							}
							return _879_v27
						})(), _dafny.Companion_Sequence_.Concatenate(_879_v27, _879_v27))
						_ = _rhs150
						var _lhs109 _dafny.Array = _877_v25
						_ = _lhs109
						var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_877_v25), 0))
						_ = _lhs110
						var _lhs111 _dafny.Array = _880_v28
						_ = _lhs111
						var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_880_v28), 0))
						_ = _lhs112
						var _lhs113 _dafny.Array = _883_v29
						_ = _lhs113
						var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_883_v29), 0))
						_ = _lhs114
						_842_v1 = _rhs146
						_845_v4 = _rhs147
						(_lhs109).ArraySet1(_rhs148, (_lhs110).Int())
						(_lhs111).ArraySet1(_rhs149, (_lhs112).Int())
						(_lhs113).ArraySet1(_rhs150, (_lhs114).Int())
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
						_ = _index155
						(_839_v0).ArraySet1(_dafny.IntOfInt64(600), (_index155).Int())
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_883_v29), 0))
						_ = _index156
						(_883_v29).ArraySet1((_884_v30).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_884_v30).Cardinality()))).Uint32()).(bool), (_index156).Int())
						_883_v29 = _883_v29
					}
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _885_v31 _dafny.Map
		_ = _885_v31
		_885_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
		var _886_v32 _dafny.Sequence
		_ = _886_v32
		_886_v32 = _dafny.SeqOf(_885_v31, _885_v31, _885_v31)
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
		_ = _index157
		(_839_v0).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if !(p0) || (!(_842_v1)) {
				return (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int)
			}
			return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-147))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_887_v1 bool) func(_dafny.Int) _dafny.Map {
				return func(_888_i4 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_this.F18(), _this.F18(), _this.F18())).Cardinality(), _887_v1)
				}
			})(_842_v1))), _886_v32))).Cardinality()
		})()), (_index157).Int())
		var _889_v33 _dafny.Map
		_ = _889_v33
		_889_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(135))
		var _hi8 _dafny.Int = (_889_v33).Cardinality()
		_ = _hi8
		for _890_i5 := p3; _890_i5.Cmp(_hi8) < 0; _890_i5 = _890_i5.Plus(_dafny.One) {
			var _891_v34 D1
			_ = _891_v34
			_891_v34 = Companion_D1_.Create_DC5_(p1)
			var _892_v35 _dafny.Array
			_ = _892_v35
			var _nwElement0_27 D1 = _891_v34
			_ = _nwElement0_27
			var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(6))
			_ = _nw143
			(_nw143).ArraySet1(_nwElement0_27, 0)
			(_nw143).ArraySet1((func() D1 {
				if _842_v1 {
					return _891_v34
				}
				return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(717))
			})(), 1)
			(_nw143).ArraySet1(_891_v34, 2)
			(_nw143).ArraySet1(_891_v34, 3)
			(_nw143).ArraySet1(_891_v34, 4)
			(_nw143).ArraySet1(_891_v34, 5)
			_892_v35 = _nw143
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_892_v35), 0))
			_ = _index158
			(_892_v35).ArraySet1(_891_v34, (_index158).Int())
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_892_v35), 0))
			_ = _index159
			(_892_v35).ArraySet1(Companion_D1_.Create_DC5_(p2), (_index159).Int())
			_842_v1 = (p0) && (_842_v1)
			var _893_v36 _dafny.Array
			_ = _893_v36
			var _nw144 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw144
			_893_v36 = _nw144
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_893_v36), 0))
			_ = _index160
			(_893_v36).ArraySet1(!(p0), (_index160).Int())
			var _894_v37 _dafny.MultiSet
			_ = _894_v37
			_894_v37 = _dafny.MultiSetOf((_dafny.Zero).Minus(p1), p3, p1)
			var _895_v38 _dafny.Map
			_ = _895_v38
			_895_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v1, p0)
			var _896_v39 T1
			_ = _896_v39
			var _nw145 *C2 = New_C2_()
			_ = _nw145
			_nw145.Ctor__(_this.F18(), (_this).F19())
			_896_v39 = _nw145
			var _897_v40 _dafny.Map
			_ = _897_v40
			_897_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_896_v39, _842_v1)
			var _898_v41 D1
			_ = _898_v41
			_898_v41 = Companion_D1_.Create_DC4_(p1, (_897_v40).Cardinality(), _842_v1)
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_893_v36), 0))
			_ = _index161
			var _rhs151 bool = ((_this).Fm3(_842_v1, (func() _dafny.Int {
				if (_894_v37).Contains((_dafny.Zero).Minus(p3)) {
					return (_894_v37).Multiplicity((_dafny.Zero).Minus(p3))
				}
				return (_895_v38).Cardinality()
			})(), globalState)) && ((func() bool {
				if p0 {
					return Companion_Default___.Fm1(p0, Companion_Default___.Fm10(globalState), _890_i5, globalState)
				}
				return _842_v1
			})())
			_ = _rhs151
			var _rhs152 _dafny.Set = _843_v2
			_ = _rhs152
			var _rhs153 bool = (_898_v41).Dtor_cf5()
			_ = _rhs153
			var _lhs115 _dafny.Array = _893_v36
			_ = _lhs115
			var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_893_v36), 0))
			_ = _lhs116
			(_lhs115).ArraySet1(_rhs151, (_lhs116).Int())
			_843_v2 = _rhs152
			_842_v1 = _rhs153
			var _899_v42 _dafny.Array
			_ = _899_v42
			var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw146
			_899_v42 = _nw146
			var _900_v43 _dafny.Array
			_ = _900_v43
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw147
			_900_v43 = _nw147
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_899_v42), 0))
			_ = _index162
			(_899_v42).ArraySet1(_900_v43, (_index162).Int())
			var _901_v44 _dafny.Sequence
			_ = _901_v44
			_901_v44 = _dafny.SeqOf(_842_v1)
			var _902_v45 _dafny.Set
			_ = _902_v45
			_902_v45 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_901_v44, _dafny.SeqOf((_893_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_893_v36), 0))).Int()).(bool), _842_v1, _842_v1, true, !(p0))), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_901_v44, _dafny.SeqOf((_893_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_893_v36), 0))).Int()).(bool), _842_v1, _842_v1, true, !(p0)))).Cardinality()))).Uint32(), _842_v1))
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_899_v42), 0))
			_ = _index163
			var _rhs154 bool = (_893_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_893_v36), 0))).Int()).(bool)
			_ = _rhs154
			var _rhs155 _dafny.Array = _900_v43
			_ = _rhs155
			var _rhs156 _dafny.Set = _902_v45
			_ = _rhs156
			var _lhs117 _dafny.Array = _899_v42
			_ = _lhs117
			var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_899_v42), 0))
			_ = _lhs118
			_842_v1 = _rhs154
			(_lhs117).ArraySet1(_rhs155, (_lhs118).Int())
			_902_v45 = _rhs156
		}
		var _903_v46 *C2
		_ = _903_v46
		var _nw148 *C2 = New_C2_()
		_ = _nw148
		_nw148.Ctor__(_dafny.CodePoint('q'), (_this).F19())
		_903_v46 = _nw148
		var _904_v47 _dafny.Map
		_ = _904_v47
		_904_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v46, p0)
		var _905_v48 _dafny.Set
		_ = _905_v48
		_905_v48 = _dafny.SetOf(_904_v47, _904_v47, _904_v47, _904_v47)
		var _906_v49 _dafny.Sequence
		_ = _906_v49
		_906_v49 = _dafny.SeqOf(_905_v48)
		var _907_v50 _dafny.MultiSet
		_ = _907_v50
		_907_v50 = _dafny.MultiSetOf(p0)
		if !(!((_906_v49).Select((Companion_Default___.SafeIndex((_907_v50).Cardinality(), _dafny.IntOfUint32((_906_v49).Cardinality()))).Uint32()).(_dafny.Set)).Equals((_905_v48).Difference(_905_v48))) {
			var _908_v51 _dafny.Sequence
			_ = _908_v51
			_908_v51 = _dafny.SeqOf(p0)
			var _909_v52 _dafny.Array
			_ = _909_v52
			var _nwElement0_28 bool = p0
			_ = _nwElement0_28
			var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(29))
			_ = _nw149
			(_nw149).ArraySet1(_nwElement0_28, 0)
			(_nw149).ArraySet1(p0, 1)
			(_nw149).ArraySet1(true, 2)
			(_nw149).ArraySet1(p0, 3)
			(_nw149).ArraySet1(_842_v1, 4)
			(_nw149).ArraySet1(_842_v1, 5)
			(_nw149).ArraySet1(_842_v1, 6)
			(_nw149).ArraySet1((_908_v51).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_908_v51).Cardinality()))).Uint32()).(bool), 7)
			(_nw149).ArraySet1(_842_v1, 8)
			(_nw149).ArraySet1(p0, 9)
			(_nw149).ArraySet1(_842_v1, 10)
			(_nw149).ArraySet1(p0, 11)
			(_nw149).ArraySet1(_842_v1, 12)
			(_nw149).ArraySet1(_842_v1, 13)
			(_nw149).ArraySet1(_842_v1, 14)
			(_nw149).ArraySet1(_842_v1, 15)
			(_nw149).ArraySet1(p0, 16)
			(_nw149).ArraySet1(p0, 17)
			(_nw149).ArraySet1(_842_v1, 18)
			(_nw149).ArraySet1(true, 19)
			(_nw149).ArraySet1(_842_v1, 20)
			(_nw149).ArraySet1(p0, 21)
			(_nw149).ArraySet1(_842_v1, 22)
			(_nw149).ArraySet1(p0, 23)
			(_nw149).ArraySet1(p0, 24)
			(_nw149).ArraySet1(_842_v1, 25)
			(_nw149).ArraySet1((_908_v51).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v1, _839_v0)).Cardinality()), _dafny.IntOfUint32((_908_v51).Cardinality()))).Uint32()).(bool), 26)
			(_nw149).ArraySet1(p0, 27)
			(_nw149).ArraySet1(p0, 28)
			_909_v52 = _nw149
			var _910_v53 _dafny.Map
			_ = _910_v53
			_910_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _909_v52)
			var _911_v54 _dafny.Array
			_ = _911_v54
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw150
			_911_v54 = _nw150
			var _912_v55 _dafny.Map
			_ = _912_v55
			_912_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_910_v53).Contains(_842_v1) {
					return (_910_v53).Get(_842_v1).(_dafny.Array)
				}
				return _909_v52
			})(), _911_v54)
			var _913_v56 _dafny.Map
			_ = _913_v56
			_913_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_912_v55).Contains(_909_v52) {
					return (_912_v55).Get(_909_v52).(_dafny.Array)
				}
				return _911_v54
			})(), p0)
			_913_v56 = (_913_v56).Update(_911_v54, p0)
			var _914_v57 _dafny.Map
			_ = _914_v57
			_914_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _915_v58 _dafny.Sequence
			_ = _915_v58
			_915_v58 = _dafny.SeqOf(_914_v57)
			var _916_v59 _dafny.MultiSet
			_ = _916_v59
			_916_v59 = _dafny.MultiSetOf(_914_v57, _914_v57, _914_v57, (_914_v57).Merge((_915_v58).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_915_v58).Cardinality()))).Uint32()).(_dafny.Map)), Companion_Default___.Fm23(p1, _842_v1, globalState))
			_916_v59 = _916_v59
			var _917_v60 *C1
			_ = _917_v60
			var _nw151 *C1 = New_C1_()
			_ = _nw151
			_nw151.Ctor__((p2).Plus(p2), p0, _this.F18(), Companion_Default___.Fm18((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), globalState))
			_917_v60 = _nw151
			var _918_v61 _dafny.Sequence
			_ = _918_v61
			_918_v61 = _dafny.SeqOf(p2)
			var _919_v62 _dafny.Map
			_ = _919_v62
			_919_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_918_v61, _842_v1)
			_842_v1 = !(_842_v1) || ((true) || ((func() bool {
				if (_919_v62).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-491))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}(func(_920_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(131)
				}))) {
					return (_919_v62).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-491))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg52 _dafny.Int) interface{} {
							return coer52(arg52)
						}
					}(func(_921_i6 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(131)
					}))).(bool)
				}
				return false
			})()))
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))
			_ = _index164
			(_839_v0).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_908_v51, (Companion_Default___.SafeIndex((_843_v2).Cardinality(), _dafny.IntOfUint32((_908_v51).Cardinality()))).Uint32(), !(_842_v1))).Cardinality()), (_885_v31).Cardinality()), (_index164).Int())
		} else {
			var _rhs157 _dafny.Int = p2
			_ = _rhs157
			var _rhs158 _dafny.Int = p2
			_ = _rhs158
			var _rhs159 bool = (func() bool {
				if (_843_v2).IsDisjointFrom(_843_v2) {
					return true
				}
				return !(false)
			})()
			_ = _rhs159
			var _lhs119 *GlobalState = globalState
			_ = _lhs119
			_lhs119.F6 = _rhs157
			r0 = _rhs158
			_842_v1 = _rhs159
			var _922_v63 _dafny.Sequence
			_ = _922_v63
			_922_v63 = _dafny.SeqOf(_839_v0, _839_v0, _839_v0, _839_v0, _839_v0)
			var _923_v64 D1
			_ = _923_v64
			_923_v64 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int)), (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), _842_v1)
			var _924_v65 _dafny.Map
			_ = _924_v65
			_924_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int))
			var _925_v66 _dafny.MultiSet
			_ = _925_v66
			_925_v66 = _dafny.MultiSetOf(p1)
			var _926_v67 _dafny.Sequence
			_ = _926_v67
			_926_v67 = _dafny.UnicodeSeqOfUtf8Bytes("wyt")
			var _927_v68 _dafny.Array
			_ = _927_v68
			var _nwElement0_29 bool = p0
			_ = _nwElement0_29
			var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(28))
			_ = _nw152
			(_nw152).ArraySet1(_nwElement0_29, 0)
			(_nw152).ArraySet1(_842_v1, 1)
			(_nw152).ArraySet1(_842_v1, 2)
			(_nw152).ArraySet1(_842_v1, 3)
			(_nw152).ArraySet1(!(true), 4)
			(_nw152).ArraySet1(p0, 5)
			(_nw152).ArraySet1(_842_v1, 6)
			(_nw152).ArraySet1(!_dafny.Companion_Sequence_.Contains(_922_v63, _839_v0), 7)
			(_nw152).ArraySet1(_842_v1, 8)
			(_nw152).ArraySet1(!(p0), 9)
			(_nw152).ArraySet1((func() bool {
				if (_885_v31).Contains((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int)) {
					return (_885_v31).Get((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int)).(bool)
				}
				return (_923_v64).Dtor_cf5()
			})(), 10)
			(_nw152).ArraySet1((_924_v65).Contains(p0), 11)
			(_nw152).ArraySet1(_842_v1, 12)
			(_nw152).ArraySet1((_925_v66).IsSubsetOf(_925_v66), 13)
			(_nw152).ArraySet1(_842_v1, 14)
			(_nw152).ArraySet1(((_845_v4).Dtor_cf11()) == (p0), 15)
			(_nw152).ArraySet1((func() bool {
				if (_885_v31).Contains(p2) {
					return (_885_v31).Get(p2).(bool)
				}
				return !(p0)
			})(), 16)
			(_nw152).ArraySet1(true, 17)
			(_nw152).ArraySet1(true, 18)
			(_nw152).ArraySet1(!(_842_v1) || (p0), 19)
			(_nw152).ArraySet1(!(p0), 20)
			(_nw152).ArraySet1(!_dafny.Companion_Sequence_.Equal(_926_v67, _926_v67), 21)
			(_nw152).ArraySet1(p0, 22)
			(_nw152).ArraySet1((func() bool {
				if _842_v1 {
					return false
				}
				return _842_v1
			})(), 23)
			(_nw152).ArraySet1(_842_v1, 24)
			(_nw152).ArraySet1(!(true), 25)
			(_nw152).ArraySet1((func() bool {
				if !(_842_v1) {
					return p0
				}
				return _842_v1
			})(), 26)
			(_nw152).ArraySet1(p0, 27)
			_927_v68 = _nw152
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))
			_ = _index165
			(_927_v68).ArraySet1(_842_v1, (_index165).Int())
			var _928_v69 _dafny.Map
			_ = _928_v69
			_928_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_907_v50, p2)
			var _929_v71 _dafny.Set
			_ = _929_v71
			_929_v71 = _dafny.SetOf(_907_v50, _907_v50)
			var _930_v72 _dafny.Map
			_ = _930_v72
			_930_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_907_v50, p0)
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))
			_ = _index166
			var _rhs160 bool = (func() _dafny.Set {
				var _coll38 = _dafny.NewBuilder()
				_ = _coll38
				for _iter40 := _dafny.Iterate(((_930_v72).Update(_907_v50, p0)).Keys().Elements()); ; {
					_compr_38, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _931_v73 _dafny.MultiSet
					_931_v73 = interface{}(_compr_38).(_dafny.MultiSet)
					if ((_930_v72).Update(_907_v50, p0)).Contains(_931_v73) {
						_coll38.Add(_931_v73)
					}
				}
				return _coll38.ToSet()
			}()).IsProperSubsetOf((func() _dafny.Set {
				var _coll39 = _dafny.NewBuilder()
				_ = _coll39
				for _iter41 := _dafny.Iterate((_928_v69).Keys().Elements()); ; {
					_compr_39, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _932_v70 _dafny.MultiSet
					_932_v70 = interface{}(_compr_39).(_dafny.MultiSet)
					if (_928_v69).Contains(_932_v70) {
						_coll39.Add(_932_v70)
					}
				}
				return _coll39.ToSet()
			}()).Union(_929_v71))
			_ = _rhs160
			var _rhs161 bool = _842_v1
			_ = _rhs161
			var _rhs162 _dafny.Sequence = _926_v67
			_ = _rhs162
			var _lhs120 _dafny.Array = _927_v68
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))
			_ = _lhs121
			(_lhs120).ArraySet1(_rhs160, (_lhs121).Int())
			_842_v1 = _rhs161
			_926_v67 = _rhs162
			if true {
				_842_v1 = !((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool)) || ((func() bool {
					if false {
						return false
					}
					return p0
				})())
				_889_v33 = _889_v33
				var _933_v75 _dafny.Sequence
				_ = _933_v75
				_933_v75 = _dafny.SeqOf((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool))
				var _934_v76 _dafny.Sequence
				_ = _934_v76
				_934_v76 = _933_v75
				var _935_v77 _dafny.Sequence
				_ = _935_v77
				_935_v77 = _dafny.SeqOf(_934_v76, _933_v75)
				var _936_v78 _dafny.Sequence
				_ = _936_v78
				_936_v78 = _dafny.SeqOf(_935_v77)
				var _937_v79 _dafny.Map
				_ = _937_v79
				_937_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(304), _dafny.IntOfInt64(-945))); ; {
						_compr_40, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _938_v74 _dafny.Int
						_938_v74 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(304)).Cmp(_938_v74) <= 0) && ((_938_v74).Cmp(_dafny.IntOfInt64(-945)) < 0) {
							_coll40.Add((_938_v74).Minus(p1), p0)
						}
					}
					return _coll40.ToMap()
				}(), (_936_v78).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_936_v78).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _939_v80 _dafny.Map
				_ = _939_v80
				_939_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v2, _937_v79)
				_939_v80 = (_939_v80).Update((func() _dafny.Set {
					if (_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool) {
						return _843_v2
					}
					return _dafny.SetOf((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool), p0, !((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool)))
				})(), _937_v79)
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))
				_ = _index167
				(_927_v68).ArraySet1(p0, (_index167).Int())
				r0 = p3
			} else {
				var _940_v81 D3
				_ = _940_v81
				_940_v81 = Companion_D3_.Create_DC8_(_dafny.CodePoint('d'))
				var _941_v82 _dafny.Map
				_ = _941_v82
				_941_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_940_v81).Dtor_cf10(), p3)
				var _942_v83 *C0
				_ = _942_v83
				var _nw153 *C0 = New_C0_()
				_ = _nw153
				_nw153.Ctor__(_926_v67, _941_v82, _this.F18(), ((_this).F19()).Merge((_this).F19()))
				_942_v83 = _nw153
				var _943_v84 *C0
				_ = _943_v84
				var _nw154 *C0 = New_C0_()
				_ = _nw154
				_nw154.Ctor__(Companion_Default___.Fm8(globalState), (_942_v83).F21(), _this.F18(), (_this).F19())
				_943_v84 = _nw154
				_943_v84 = _942_v83
				var _944_v85 D1
				_ = _944_v85
				_944_v85 = Companion_D1_.Create_DC6_(_842_v1, (_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool))
				var _945_v86 _dafny.Map
				_ = _945_v86
				_945_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _944_v85)
				_945_v86 = (_945_v86).Update((_839_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_839_v0), 0))).Int()).(_dafny.Int), _944_v85)
				var _946_v87 *C0
				_ = _946_v87
				var _nw155 *C0 = New_C0_()
				_ = _nw155
				_nw155.Ctor__((func() _dafny.Sequence {
					if !(!(true)) {
						return (_943_v84).F20()
					}
					return _926_v67
				})(), (_941_v82).Merge(_941_v82), _this.F18(), (_this).F19())
				_946_v87 = _nw155
				(globalState).F6 = Companion_Default___.SafeModuloInt((p3).Minus(p1), _dafny.IntOfInt64(259))
			}
			var _947_v88 _dafny.Map
			_ = _947_v88
			_947_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_842_v1, p0))
			var _948_v89 D3
			_ = _948_v89
			_948_v89 = Companion_D3_.Create_DC10_(!(p0))
			var _949_v90 _dafny.Map
			_ = _949_v90
			_949_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_947_v88, _948_v89)
			_949_v90 = (_949_v90).Update(_947_v88, Companion_Default___.Fm24(globalState))
			var _950_v92 _dafny.Map
			_ = _950_v92
			_950_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), _842_v1)
			var _951_v93 D6
			_ = _951_v93
			_951_v93 = Companion_D6_.Create_DC15_(func() _dafny.Map {
				var _coll41 = _dafny.NewMapBuilder()
				_ = _coll41
				for _iter43 := _dafny.Iterate((_950_v92).Keys().Elements()); ; {
					_compr_41, _ok43 := _iter43()
					if !_ok43 {
						break
					}
					var _952_v91 _dafny.CodePoint
					_952_v91 = interface{}(_compr_41).(_dafny.CodePoint)
					if (_950_v92).Contains(_952_v91) {
						_coll41.Add(_952_v91, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()))
					}
				}
				return _coll41.ToMap()
			}())
			var _953_v94 D6
			_ = _953_v94
			_953_v94 = Companion_D6_.Create_DC17_(_951_v93)
			var _954_v95 _dafny.Map
			_ = _954_v95
			_954_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_953_v94, (_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool))
			var _955_v96 _dafny.Sequence
			_ = _955_v96
			_955_v96 = _dafny.SeqOf(_954_v95)
			var _956_v97 _dafny.Array
			_ = _956_v97
			var _nwElement0_30 _dafny.Map = _954_v95
			_ = _nwElement0_30
			var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(29))
			_ = _nw156
			(_nw156).ArraySet1(_nwElement0_30, 0)
			(_nw156).ArraySet1((_954_v95).Merge((_954_v95).Update(_953_v94, p0)), 1)
			(_nw156).ArraySet1(_954_v95, 2)
			(_nw156).ArraySet1(_954_v95, 3)
			(_nw156).ArraySet1((_954_v95).Merge(_954_v95), 4)
			(_nw156).ArraySet1(_954_v95, 5)
			(_nw156).ArraySet1(_954_v95, 6)
			(_nw156).ArraySet1(_954_v95, 7)
			(_nw156).ArraySet1(_954_v95, 8)
			(_nw156).ArraySet1(_954_v95, 9)
			(_nw156).ArraySet1((_955_v96).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_955_v96).Cardinality()))).Uint32()).(_dafny.Map), 10)
			(_nw156).ArraySet1(_954_v95, 11)
			(_nw156).ArraySet1(_954_v95, 12)
			(_nw156).ArraySet1(_954_v95, 13)
			(_nw156).ArraySet1(_954_v95, 14)
			(_nw156).ArraySet1(_954_v95, 15)
			(_nw156).ArraySet1(_954_v95, 16)
			(_nw156).ArraySet1((func() _dafny.Map {
				if _842_v1 {
					return _954_v95
				}
				return _954_v95
			})(), 17)
			(_nw156).ArraySet1(_954_v95, 18)
			(_nw156).ArraySet1(_954_v95, 19)
			(_nw156).ArraySet1((_954_v95).Update(_953_v94, (_this).Fm3(!((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool)), (func() _dafny.Int {
				if (_907_v50).Contains((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool)) {
					return (_907_v50).Multiplicity((_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool))
				}
				return _dafny.IntOfUint32((_926_v67).Cardinality())
			})(), globalState)), 20)
			(_nw156).ArraySet1(_954_v95, 21)
			(_nw156).ArraySet1(_954_v95, 22)
			(_nw156).ArraySet1(_954_v95, 23)
			(_nw156).ArraySet1((_954_v95).Merge((_954_v95).Update(_953_v94, (_927_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_927_v68), 0))).Int()).(bool))), 24)
			(_nw156).ArraySet1((_954_v95).Merge(_954_v95), 25)
			(_nw156).ArraySet1(_954_v95, 26)
			(_nw156).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC17_(Companion_D6_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18(), p3))), _842_v1), 27)
			(_nw156).ArraySet1(_954_v95, 28)
			_956_v97 = _nw156
			_956_v97 = _956_v97
		}
		var _957_v98 _dafny.Sequence
		_ = _957_v98
		_957_v98 = _dafny.UnicodeSeqOfUtf8Bytes("dv")
		var _958_v99 D9
		_ = _958_v99
		_958_v99 = Companion_D9_.Create_DC22_(_957_v98)
		var _959_v100 _dafny.Array
		_ = _959_v100
		var _nwElement0_31 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(670))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg53 _dafny.Int) interface{} {
				return coer53(arg53)
			}
		}(func(_960_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))
		_ = _nwElement0_31
		var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(16))
		_ = _nw157
		(_nw157).ArraySet1(_nwElement0_31, 0)
		(_nw157).ArraySet1(_957_v98, 1)
		(_nw157).ArraySet1((_958_v99).Dtor_cf27(), 2)
		(_nw157).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ude"), 3)
		(_nw157).ArraySet1(_957_v98, 4)
		(_nw157).ArraySet1(_957_v98, 5)
		(_nw157).ArraySet1(_957_v98, 6)
		(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_957_v98, _957_v98), 7)
		(_nw157).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg54 _dafny.Int) interface{} {
				return coer54(arg54)
			}
		}(func(_961_i9 _dafny.Int) _dafny.CodePoint {
			return _this.F18()
		})), 8)
		(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_957_v98, _957_v98), 9)
		(_nw157).ArraySet1(_957_v98, 10)
		(_nw157).ArraySet1((_958_v99).Dtor_cf27(), 11)
		(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_957_v98, _957_v98), 12)
		(_nw157).ArraySet1(_957_v98, 13)
		(_nw157).ArraySet1(_957_v98, 14)
		(_nw157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("o"), _957_v98), 15)
		_959_v100 = _nw157
		for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_959_v100), 0))); ; {
			_guard_loop_2, _ok44 := _iter44()
			if !_ok44 {
				break
			}
			var _962_i7 _dafny.Int
			_962_i7 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_962_i7).Sign() != -1) && ((_962_i7).Cmp(_dafny.ArrayLen((_959_v100), 0)) < 0)) {
				(_959_v100).ArraySet1(_957_v98, (_962_i7).Int())
			}
		}
		r0 = p3
		return r0
	}
}
func (_this *C3) M3(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _963_v0 _dafny.Sequence
		_ = _963_v0
		_963_v0 = _dafny.UnicodeSeqOfUtf8Bytes("bfbxyfb")
		var _964_v1 _dafny.Map
		_ = _964_v1
		_964_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Cardinality(), p1)
		var _965_v2 _dafny.Sequence
		_ = _965_v2
		_965_v2 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(p1, p1)).Cardinality()), p0, p0, p0, _dafny.IntOfUint32((_963_v0).Cardinality()))
		var _hi9 _dafny.Int = ((_964_v1).Cardinality()).Minus((_965_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.IntOfUint32((_965_v2).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _hi9
		for _966_i0 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_963_v0).Cardinality())); _966_i0.Cmp(_hi9) < 0; _966_i0 = _966_i0.Plus(_dafny.One) {
			_963_v0 = _963_v0
			var _967_v3 _dafny.Array
			_ = _967_v3
			var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
			_ = _nw158
			_967_v3 = _nw158
			var _968_v4 _dafny.Map
			_ = _968_v4
			_968_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p2).Cardinality())
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_967_v3), 0))
			_ = _index168
			(_967_v3).ArraySet1(_968_v4, (_index168).Int())
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_967_v3), 0))
			_ = _index169
			(_967_v3).ArraySet1((func() _dafny.Map {
				if (p1) && (p1) {
					return _968_v4
				}
				return _968_v4
			})(), (_index169).Int())
			var _969_v5 _dafny.Array
			_ = _969_v5
			var _970_v6 bool
			_ = _970_v6
			var _971_v7 _dafny.Int
			_ = _971_v7
			var _out24 _dafny.Array
			_ = _out24
			var _out25 bool
			_ = _out25
			var _out26 _dafny.Int
			_ = _out26
			_out24, _out25, _out26 = Companion_Default___.M0((_this).Fm4(p1, _dafny.CodePoint('e'), _this.F18(), p1, globalState), p1, _966_i0, globalState)
			_969_v5 = _out24
			_970_v6 = _out25
			_971_v7 = _out26
			var _972_v8 _dafny.MultiSet
			_ = _972_v8
			_972_v8 = _dafny.MultiSetOf(p2, p2, _dafny.MultiSetOf(_971_v7, Companion_Default___.Fm0(globalState), (_dafny.Zero).Minus(_971_v7), p0, _dafny.IntOfUint32((_963_v0).Cardinality())), (_dafny.MultiSetOf(_966_i0)).Union(p2))
			var _973_v9 _dafny.Set
			_ = _973_v9
			_973_v9 = _dafny.SetOf(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_963_v0).Cardinality()), _dafny.IntOfUint32((_963_v0).Cardinality()))
			var _rhs163 _dafny.Int = _966_i0
			_ = _rhs163
			var _rhs164 _dafny.Int = (func() _dafny.Int {
				if (_972_v8).Contains(p2) {
					return (_972_v8).Multiplicity(p2)
				}
				return _971_v7
			})()
			_ = _rhs164
			var _rhs165 _dafny.Int = Companion_Default___.SafeDivisionInt((_973_v9).Cardinality(), Companion_Default___.Fm0(globalState))
			_ = _rhs165
			var _rhs166 _dafny.Sequence = Companion_Default___.Fm8(globalState)
			_ = _rhs166
			var _lhs122 *GlobalState = globalState
			_ = _lhs122
			var _lhs123 *GlobalState = globalState
			_ = _lhs123
			r0 = _rhs163
			_lhs122.F6 = _rhs164
			_lhs123.F6 = _rhs165
			_963_v0 = _rhs166
		}
		var _974_v10 *C2
		_ = _974_v10
		var _nw159 *C2 = New_C2_()
		_ = _nw159
		_nw159.Ctor__(_this.F18(), (_this).F19())
		_974_v10 = _nw159
		var _975_v11 _dafny.Sequence
		_ = _975_v11
		_975_v11 = _dafny.SeqOf(_974_v10)
		var _976_v12 _dafny.Map
		_ = _976_v12
		_976_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Sequence {
			if p1 {
				return _975_v11
			}
			return _975_v11
		})())
		_976_v12 = ((_976_v12).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _975_v11))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _975_v11))
		var _977_v13 _dafny.MultiSet
		_ = _977_v13
		_977_v13 = _dafny.MultiSetOf(false, false, !(p1), p1, _dafny.Companion_Sequence_.IsPrefixOf(_963_v0, _963_v0))
		_977_v13 = _977_v13
		var _978_v14 bool
		_ = _978_v14
		_978_v14 = true
		_978_v14 = !(_978_v14)
		_978_v14 = !(p1)
		var _979_i1 _dafny.Int
		_ = _979_i1
		_979_i1 = _dafny.Zero
		{
			for _978_v14 {
				{
					if (_979_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_979_i1 = (_979_i1).Plus(_dafny.One)
					_978_v14 = p1
					(_this).F18_set_((_963_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_963_v0, _963_v0)).Cardinality()), _dafny.IntOfUint32((_963_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
					_978_v14 = _978_v14
					var _980_v15 _dafny.Set
					_ = _980_v15
					_980_v15 = _dafny.SetOf((_this).Fm4(p1, (_963_v0).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_963_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F18(), _978_v14, globalState))
					_978_v14 = (_980_v15).IsSubsetOf(_980_v15)
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		r0 = p0
		return r0
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
