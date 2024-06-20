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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) bool {
	return (Companion_D7_.Create_DC16_(_dafny.IntOfInt64(163), _dafny.IntOfInt64(125), true, false, _dafny.IntOfInt64(276))).Dtor_cf26()
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(387))).Intersection(_dafny.SetOf(_dafny.IntOfInt64(-805)))).Union((_dafny.SetOf(_dafny.IntOfInt64(624), (Companion_D6_.Create_DC13_(_dafny.IntOfInt64(510))).Dtor_cf19())))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Map, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('f')
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)), _dafny.SeqOf(true, false))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(549)
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, true, !(true), !(false), false)).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(192)))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false)))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true))))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(315), false)).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(525), _dafny.IntOfInt64(56))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(525)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(56)) < 0) {
				_coll0.Add((_0_v0).Minus((_dafny.MultiSetOf(false)).Cardinality()), !(!(false)))
			}
		}
		return _coll0.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("drfsgt"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.UnicodeSeqOfUtf8Bytes("wtfucath")))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) D3 {
	if false {
		return Companion_D3_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(864))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})))
	} else {
		return Companion_D3_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-403))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source0 D7 = (func() D7 {
		if false {
			return Companion_D7_.Create_DC15_((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(983), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(true))).Cardinality())
		}
		return Companion_D7_.Create_DC15_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(!(false)))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ujipcla")).Cardinality()))
	})()
	_ = _source0
	if _source0.Is_DC15() {
		var _3___mcc_h0 _dafny.Int = _source0.Get_().(D7_DC15).Cf21
		_ = _3___mcc_h0
		var _4___mcc_h1 _dafny.Int = _source0.Get_().(D7_DC15).Cf22
		_ = _4___mcc_h1
		var _5_cf22 _dafny.Int = _4___mcc_h1
		_ = _5_cf22
		var _6_cf21 _dafny.Int = _3___mcc_h0
		_ = _6_cf21
		return func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("htgdju")).Cardinality()), _6_cf21), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_cf21, _5_cf22), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_cf21, _6_cf21)))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _7_v0 _dafny.Map
				_7_v0 = interface{}(_compr_1).(_dafny.Map)
				if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("htgdju")).Cardinality()), _6_cf21), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_cf21, _5_cf22), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_cf21, _6_cf21)))).Contains(_7_v0) {
					_coll1.Add(_7_v0, true)
				}
			}
			return _coll1.ToMap()
		}()
	} else if _source0.Is_DC16() {
		var _8___mcc_h2 _dafny.Int = _source0.Get_().(D7_DC16).Cf23
		_ = _8___mcc_h2
		var _9___mcc_h3 _dafny.Int = _source0.Get_().(D7_DC16).Cf24
		_ = _9___mcc_h3
		var _10___mcc_h4 bool = _source0.Get_().(D7_DC16).Cf25
		_ = _10___mcc_h4
		var _11___mcc_h5 bool = _source0.Get_().(D7_DC16).Cf26
		_ = _11___mcc_h5
		var _12___mcc_h6 _dafny.Int = _source0.Get_().(D7_DC16).Cf27
		_ = _12___mcc_h6
		var _13_cf27 _dafny.Int = _12___mcc_h6
		_ = _13_cf27
		var _14_cf26 bool = _11___mcc_h5
		_ = _14_cf26
		var _15_cf25 bool = _10___mcc_h4
		_ = _15_cf25
		var _16_cf24 _dafny.Int = _9___mcc_h3
		_ = _16_cf24
		var _17_cf23 _dafny.Int = _8___mcc_h2
		_ = _17_cf23
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_cf23, _16_cf24), _15_cf25)
	} else if _source0.Is_DC14() {
		var _18___mcc_h7 _dafny.Map = _source0.Get_().(D7_DC14).Cf20
		_ = _18___mcc_h7
		var _19_cf20 _dafny.Map = _18___mcc_h7
		_ = _19_cf20
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(51), true)).Keys().Elements()); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _20_v1 _dafny.Int
				_20_v1 = interface{}(_compr_2).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(51), true)).Contains(_20_v1) {
					_coll2.Add((_20_v1).Times((_dafny.SetOf(false, false)).Cardinality()), _dafny.IntOfInt64(420))
				}
			}
			return _coll2.ToMap()
		}(), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.MultiSetOf((func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _21_v3 _dafny.Sequence
					_21_v3 = interface{}(_compr_4).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Contains(_21_v3) {
						_coll4.Add(_21_v3, _dafny.IntOfInt64(-542))
					}
				}
				return _coll4.ToMap()
			}()).Cardinality(), (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))).Cardinality()))).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _23_v4 _dafny.Int
					_23_v4 = interface{}(_compr_5).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg3 _dafny.Int) interface{} {
							return coer3(arg3)
						}
					}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality())), _23_v4) {
						_coll5.Add((_23_v4).Plus(_dafny.IntOfInt64(405)), _dafny.IntOfInt64(190))
					}
				}
				return _coll5.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(344))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_24_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(83)
			}))).Cardinality()), _dafny.IntOfInt64(350))).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _25_v2 _dafny.Int
				_25_v2 = interface{}(_compr_3).(_dafny.Int)
				if (_dafny.MultiSetOf((func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Elements()); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _21_v3 _dafny.Sequence
						_21_v3 = interface{}(_compr_6).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Contains(_21_v3) {
							_coll6.Add(_21_v3, _dafny.IntOfInt64(-542))
						}
					}
					return _coll6.ToMap()
				}()).Cardinality(), (func() _dafny.Map {
					var _coll7 = _dafny.NewMapBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg5 _dafny.Int) interface{} {
							return coer5(arg5)
						}
					}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality()))).Elements()); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _23_v4 _dafny.Int
						_23_v4 = interface{}(_compr_7).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg6 _dafny.Int) interface{} {
								return coer6(arg6)
							}
						}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('a')
						}))).Cardinality())), _23_v4) {
							_coll7.Add((_23_v4).Plus(_dafny.IntOfInt64(405)), _dafny.IntOfInt64(190))
						}
					}
					return _coll7.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(344))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_24_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(83)
				}))).Cardinality()), _dafny.IntOfInt64(350))).Contains(_25_v2) {
					_coll3.Add(Companion_Default___.SafeDivisionInt(_25_v2, (_dafny.SetOf(!(true))).Cardinality()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(true, false, true)).Cardinality()))).Cardinality()))
				}
			}
			return _coll3.ToMap()
		}(), !(true)))
	} else {
		var _26___mcc_h8 D7 = _source0.Get_().(D7_DC17).Cf28
		_ = _26___mcc_h8
		var _27_cf28 D7 = _26___mcc_h8
		_ = _27_cf28
		return func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(658), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_28_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Cardinality()), _dafny.IntOfInt64(-688), _dafny.IntOfInt64(-524), _dafny.IntOfInt64(-555))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvgbw")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(394))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-341)))).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _29_v5 _dafny.Map
				_29_v5 = interface{}(_compr_8).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(658), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_28_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()), _dafny.IntOfInt64(-688), _dafny.IntOfInt64(-524), _dafny.IntOfInt64(-555))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvgbw")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(394))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-341))), _29_v5) {
					_coll8.Add(_29_v5, false)
				}
			}
			return _coll8.ToMap()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(492), _dafny.IntOfInt64(841))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _30_v0 _dafny.Int
			_30_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(492)).Cmp(_30_v0) <= 0) && ((_30_v0).Cmp(_dafny.IntOfInt64(841)) < 0) {
				_coll9.Add((_30_v0).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(256))).Cardinality()))))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("orpdwcbhu")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(893))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))).Cardinality()))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.SeqOf((_dafny.MultiSetOf(false, false)).Cardinality())).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _32_v1 _dafny.Int
			_32_v1 = interface{}(_compr_10).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.MultiSetOf(false, false)).Cardinality()), _32_v1) {
				_coll10.Add((_32_v1).Plus(_dafny.IntOfInt64(445)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(908))).Cardinality())
			}
		}
		return _coll10.ToMap()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) {
	var _33_i0 _dafny.Int
	_ = _33_i0
	_33_i0 = _dafny.Zero
	{
		for p3 {
			{
				if (_33_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_33_i0 = (_33_i0).Plus(_dafny.One)
				var _34_v0 _dafny.Map
				_ = _34_v0
				_34_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
				(globalState).F8 = ((_34_v0).Update(p3, p0)).Cardinality()
				var _35_v1 *C0
				_ = _35_v1
				var _nw0 *C0 = New_C0_()
				_ = _nw0
				_nw0.Ctor__()
				_35_v1 = _nw0
				var _36_v2 *C0
				_ = _36_v2
				var _nw1 *C0 = New_C0_()
				_ = _nw1
				_nw1.Ctor__()
				_36_v2 = _nw1
				var _37_v3 D3
				_ = _37_v3
				_37_v3 = Companion_D3_.Create_DC5_(p3, p1)
				var _38_v4 _dafny.Sequence
				_ = _38_v4
				_38_v4 = _dafny.UnicodeSeqOfUtf8Bytes("bkerwni")
				var _39_v5 _dafny.Set
				_ = _39_v5
				_39_v5 = _dafny.SetOf(false, p0)
				var _40_v6 _dafny.Map
				_ = _40_v6
				_40_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v2, p3)
				var _41_v7 _dafny.Sequence
				_ = _41_v7
				_41_v7 = _dafny.SeqOf(_40_v6)
				var _42_v8 _dafny.Array
				_ = _42_v8
				var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw2
				_42_v8 = _nw2
				var _43_v9 _dafny.Map
				_ = _43_v9
				_43_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p0, p3)).Cardinality()), _42_v8)
				var _44_v10 _dafny.Array
				_ = _44_v10
				var _nwElement0_0 _dafny.Int = p1
				_ = _nwElement0_0
				var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(21))
				_ = _nw3
				(_nw3).ArraySet1(_nwElement0_0, 0)
				(_nw3).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), 1)
				(_nw3).ArraySet1((_37_v3).Dtor_cf7(), 2)
				(_nw3).ArraySet1(p1, 3)
				(_nw3).ArraySet1((_dafny.Zero).Minus(p1), 4)
				(_nw3).ArraySet1(p1, 5)
				(_nw3).ArraySet1(p1, 6)
				(_nw3).ArraySet1(_dafny.IntOfUint32((_38_v4).Cardinality()), 7)
				(_nw3).ArraySet1((_39_v5).Cardinality(), 8)
				(_nw3).ArraySet1(p1, 9)
				(_nw3).ArraySet1(_dafny.IntOfUint32((_38_v4).Cardinality()), 10)
				(_nw3).ArraySet1(p1, 11)
				(_nw3).ArraySet1(_dafny.IntOfInt64(526), 12)
				(_nw3).ArraySet1(p1, 13)
				(_nw3).ArraySet1(p1, 14)
				(_nw3).ArraySet1(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_41_v7, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_41_v7).Cardinality()))).Uint32(), _40_v6)).Cardinality())), 15)
				(_nw3).ArraySet1(p1, 16)
				(_nw3).ArraySet1(_dafny.IntOfUint32((_38_v4).Cardinality()), 17)
				(_nw3).ArraySet1(p1, 18)
				(_nw3).ArraySet1(((_43_v9).Merge(_43_v9)).Cardinality(), 19)
				(_nw3).ArraySet1(p1, 20)
				_44_v10 = _nw3
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_44_v10), 0))
				_ = _index0
				(_44_v10).ArraySet1(((_34_v0).Update(p3, true)).Cardinality(), (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_44_v10), 0))
				_ = _index1
				(_44_v10).ArraySet1(p1, (_index1).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _45_v11 _dafny.Sequence
	_ = _45_v11
	_45_v11 = _dafny.SeqOf((_dafny.Zero).Minus(p1))
	var _46_v12 _dafny.Sequence
	_ = _46_v12
	_46_v12 = _dafny.UnicodeSeqOfUtf8Bytes("owdvbmnf")
	var _47_v13 *C0
	_ = _47_v13
	var _nw4 *C0 = New_C0_()
	_ = _nw4
	_nw4.Ctor__()
	_47_v13 = _nw4
	var _48_v14 _dafny.Sequence
	_ = _48_v14
	_48_v14 = _dafny.SeqOf(_47_v13, _47_v13)
	var _49_v15 _dafny.Sequence
	_ = _49_v15
	_49_v15 = _dafny.SeqOf(_48_v14, _48_v14, _48_v14, _dafny.SeqOf(_47_v13, _47_v13, _47_v13))
	var _50_v16 _dafny.Map
	_ = _50_v16
	_50_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_45_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_46_v12).Cardinality()), _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32()).(_dafny.Int), _dafny.Companion_Sequence_.Contains(_49_v15, _48_v14))
	var _51_i1 _dafny.Int
	_ = _51_i1
	_51_i1 = _dafny.Zero
	{
		for (func() bool {
			if (_50_v16).Contains(p1) {
				return (_50_v16).Get(p1).(bool)
			}
			return p3
		})() {
			{
				if (_51_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_51_i1 = (_51_i1).Plus(_dafny.One)
				if !(!(!(p0))) {
					var _52_v17 _dafny.Sequence
					_ = _52_v17
					_52_v17 = _dafny.SeqOf(p0, p0)
					var _53_v18 _dafny.Map
					_ = _53_v18
					_53_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_52_v17, _47_v13)
					var _54_v19 _dafny.Map
					_ = _54_v19
					_54_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C0 {
						if (_53_v18).Contains(_52_v17) {
							return (_53_v18).Get(_52_v17).(*C0)
						}
						return _47_v13
					})(), _dafny.Companion_Sequence_.IsProperPrefixOf(_45_v11, _45_v11))
					_54_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v13, p3)
					(globalState).F8 = (Companion_Default___.Fm4(p3, p1, globalState)).Times(p1)
					var _55_v20 _dafny.Array
					_ = _55_v20
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_0
					var _nw5 _dafny.Array
					_ = _nw5
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw5 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) bool = (func(_56_p3 bool, _57_v11 _dafny.Sequence, _58_p1 _dafny.Int) func(_dafny.Int) bool {
							return func(_59_i2 _dafny.Int) bool {
								return ((_dafny.MultiSetOf(_56_p3)).Cardinality()).Cmp((_57_v11).Select((Companion_Default___.SafeIndex(_58_p1, _dafny.IntOfUint32((_57_v11).Cardinality()))).Uint32()).(_dafny.Int)) < 0
							}
						})(p3, _45_v11, p1)
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
					_55_v20 = _nw5
					var _60_v21 _dafny.MultiSet
					_ = _60_v21
					_60_v21 = _dafny.MultiSetOf(p3)
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_55_v20), 0))
					_ = _index2
					(_55_v20).ArraySet1(!(_60_v21).Equals(_60_v21), (_index2).Int())
					var _61_v22 _dafny.MultiSet
					_ = _61_v22
					_61_v22 = _dafny.MultiSetOf(p1)
					var _62_v23 _dafny.Map
					_ = _62_v23
					_62_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-270), p1)
					var _63_v24 _dafny.Set
					_ = _63_v24
					_63_v24 = _dafny.SetOf(_dafny.IntOfInt64(357))
					var _64_v25 _dafny.Map
					_ = _64_v25
					_64_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v13, _63_v24)
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_55_v20), 0))
					_ = _index3
					(_55_v20).ArraySet1(((_dafny.Zero).Minus(((_45_v11).Select((Companion_Default___.SafeIndex((_61_v22).Cardinality(), _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_dafny.Zero).Minus((func() _dafny.Int {
						if (_62_v23).Contains(p1) {
							return (_62_v23).Get(p1).(_dafny.Int)
						}
						return (_64_v25).Cardinality()
					})())))).Cmp(_dafny.IntOfUint32((_52_v17).Cardinality())) > 0, (_index3).Int())
					(globalState).F8 = _dafny.IntOfInt64(182)
					var _65_v26 _dafny.Set
					_ = _65_v26
					_65_v26 = _dafny.SetOf(_47_v13, _47_v13)
					var _66_v27 _dafny.Map
					_ = _66_v27
					_66_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _65_v26)
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_55_v20), 0))
					_ = _index4
					(_55_v20).ArraySet1((_65_v26).IsDisjointFrom((func() _dafny.Set {
						if (_66_v27).Contains(p1) {
							return (_66_v27).Get(p1).(_dafny.Set)
						}
						return _65_v26
					})()), (_index4).Int())
				} else {
					var _67_v28 bool
					_ = _67_v28
					_67_v28 = false
					var _rhs0 _dafny.Int = (_45_v11).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(136)).Plus(p1), _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs0
					var _rhs1 bool = !(_67_v28)
					_ = _rhs1
					var _rhs2 _dafny.Sequence = _46_v12
					_ = _rhs2
					var _rhs3 bool = _dafny.Companion_Sequence_.Contains(_46_v12, p2)
					_ = _rhs3
					var _lhs0 *GlobalState = globalState
					_ = _lhs0
					_lhs0.F8 = _rhs0
					_67_v28 = _rhs1
					_46_v12 = _rhs2
					_67_v28 = _rhs3
					var _68_v29 _dafny.Map
					_ = _68_v29
					_68_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p0)
					var _69_v30 _dafny.Sequence
					_ = _69_v30
					_69_v30 = _dafny.SeqOf(_68_v29, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _67_v28))
					_69_v30 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_69_v30, _69_v30), _69_v30)
					var _70_v31 _dafny.Map
					_ = _70_v31
					_70_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v13, p2)
					_70_v31 = (_70_v31).Update(_47_v13, p2)
					var _71_v32 *C0
					_ = _71_v32
					var _nw6 *C0 = New_C0_()
					_ = _nw6
					_nw6.Ctor__()
					_71_v32 = _nw6
					var _72_v33 _dafny.MultiSet
					_ = _72_v33
					_72_v33 = _dafny.MultiSetOf(p2, _dafny.CodePoint('q'), p2, _dafny.CodePoint('f'), p2)
					var _73_v34 _dafny.Set
					_ = _73_v34
					_73_v34 = _dafny.SetOf(p1)
					var _74_v35 _dafny.Map
					_ = _74_v35
					_74_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
					var _75_v36 _dafny.Array
					_ = _75_v36
					var _nwElement0_1 _dafny.Int = p1
					_ = _nwElement0_1
					var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(19))
					_ = _nw7
					(_nw7).ArraySet1(_nwElement0_1, 0)
					(_nw7).ArraySet1(_dafny.IntOfInt64(293), 1)
					(_nw7).ArraySet1(p1, 2)
					(_nw7).ArraySet1((func() _dafny.Int {
						if (_72_v33).Contains(p2) {
							return (_72_v33).Multiplicity(p2)
						}
						return p1
					})(), 3)
					(_nw7).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_45_v11, _45_v11)).Cardinality()), 4)
					(_nw7).ArraySet1(p1, 5)
					(_nw7).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm4(p0, p1, globalState)), 6)
					(_nw7).ArraySet1(p1, 7)
					(_nw7).ArraySet1(p1, 8)
					(_nw7).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v32, _dafny.IntOfUint32((_45_v11).Cardinality()))).Cardinality(), 9)
					(_nw7).ArraySet1((_dafny.Zero).Minus((Companion_D3_.Create_DC5_(false, _dafny.IntOfInt64(662))).Dtor_cf7()), 10)
					(_nw7).ArraySet1(p1, 11)
					(_nw7).ArraySet1(_dafny.IntOfInt64(185), 12)
					(_nw7).ArraySet1(((_73_v34).Cardinality()).Plus(p1), 13)
					(_nw7).ArraySet1(Companion_Default___.SafeModuloInt(p1, p1), 14)
					(_nw7).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(862))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}(func(_76_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(326)
					}))).Cardinality()), 15)
					(_nw7).ArraySet1(p1, 16)
					(_nw7).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-257), p1), 17)
					(_nw7).ArraySet1(((_74_v35).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1))).Cardinality(), 18)
					_75_v36 = _nw7
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_75_v36), 0))
					_ = _index5
					(_75_v36).ArraySet1(_dafny.IntOfInt64(506), (_index5).Int())
					var _77_v37 _dafny.Map
					_ = _77_v37
					_77_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _46_v12)
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_75_v36), 0))
					_ = _index6
					(_75_v36).ArraySet1((p1).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(15), (_77_v37).Cardinality())), (_index6).Int())
				}
				var _78_v38 bool
				_ = _78_v38
				_78_v38 = false
				_78_v38 = p0
				var _79_v39 *C0
				_ = _79_v39
				var _nw8 *C0 = New_C0_()
				_ = _nw8
				_nw8.Ctor__()
				_79_v39 = _nw8
				var _80_v40 _dafny.Array
				_ = _80_v40
				var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw9
				_80_v40 = _nw9
				var _81_v41 _dafny.Map
				_ = _81_v41
				_81_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v40, _dafny.Companion_Sequence_.Concatenate(_46_v12, _dafny.Companion_Sequence_.Update(_46_v12, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_46_v12).Cardinality()))).Uint32(), _dafny.CodePoint('e'))))
				var _82_v42 D3
				_ = _82_v42
				_82_v42 = Companion_D3_.Create_DC5_(true, p1)
				var _rhs4 _dafny.Map = (_81_v41).Merge(_81_v41)
				_ = _rhs4
				var _rhs5 bool = p3
				_ = _rhs5
				var _rhs6 _dafny.Int = (_82_v42).Dtor_cf7()
				_ = _rhs6
				var _rhs7 _dafny.Int = Companion_Default___.Fm4(_78_v38, (_dafny.Zero).Minus(p1), globalState)
				_ = _rhs7
				var _lhs1 *GlobalState = globalState
				_ = _lhs1
				var _lhs2 *GlobalState = globalState
				_ = _lhs2
				_81_v41 = _rhs4
				_78_v38 = _rhs5
				_lhs1.F0 = _rhs6
				_lhs2.F0 = _rhs7
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _source1 D3 = Companion_Default___.Fm9(globalState)
	_ = _source1
	if _source1.Is_DC5() {
		var _83___mcc_h0 bool = _source1.Get_().(D3_DC5).Cf6
		_ = _83___mcc_h0
		var _84___mcc_h1 _dafny.Int = _source1.Get_().(D3_DC5).Cf7
		_ = _84___mcc_h1
		var _85_cf7 _dafny.Int = _84___mcc_h1
		_ = _85_cf7
		var _86_cf6 bool = _83___mcc_h0
		_ = _86_cf6
		var _87_v43 *C0
		_ = _87_v43
		var _nw10 *C0 = New_C0_()
		_ = _nw10
		_nw10.Ctor__()
		_87_v43 = _nw10
		var _88_v44 _dafny.CodePoint
		_ = _88_v44
		_88_v44 = _dafny.CodePoint('g')
		_88_v44 = (func() _dafny.CodePoint {
			if p0 {
				return Companion_Default___.Fm2(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(p1)), !(!(_86_cf6)), _85_cf7, globalState)
			}
			return _88_v44
		})()
		_86_cf6 = p0
		_86_cf6 = (func() bool {
			if !(_86_cf6) {
				return p3
			}
			return p0
		})()
	} else {
		var _89___mcc_h2 _dafny.Sequence = _source1.Get_().(D3_DC4).Cf5
		_ = _89___mcc_h2
		var _90_cf5 _dafny.Sequence = _89___mcc_h2
		_ = _90_cf5
		var _91_v45 _dafny.Map
		_ = _91_v45
		_91_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v11, _90_cf5)
		var _92_v46 D6
		_ = _92_v46
		_92_v46 = Companion_D6_.Create_DC10_(_dafny.SeqOf(p1))
		_91_v45 = (_91_v45).Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_45_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_92_v46).Dtor_cf14()).Cardinality()), _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32(), (_45_v11).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32()).(_dafny.Int)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_45_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_92_v46).Dtor_cf14()).Cardinality()), _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32(), (_45_v11).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality()))).Uint32(), p1), _dafny.UnicodeSeqOfUtf8Bytes("imvdwrgjx"))
		var _93_v47 *C0
		_ = _93_v47
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__()
		_93_v47 = _nw11
		(globalState).F0 = p1
		var _94_v48 _dafny.Array
		_ = _94_v48
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw12
		_94_v48 = _nw12
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_94_v48), 0))
		_ = _index7
		(_94_v48).ArraySet1((_dafny.IntOfUint32((_90_cf5).Cardinality())).Plus((func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(400), _dafny.IntOfInt64(602))); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _95_v49 _dafny.Int
				_95_v49 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(400)).Cmp(_95_v49) <= 0) && ((_95_v49).Cmp(_dafny.IntOfInt64(602)) < 0) {
					_coll11.Add(Companion_Default___.SafeDivisionInt(_95_v49, p1), p2)
				}
			}
			return _coll11.ToMap()
		}()).Cardinality()), (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_94_v48), 0))
		_ = _index8
		(_94_v48).ArraySet1(p1, (_index8).Int())
	}
	var _96_v50 _dafny.MultiSet
	_ = _96_v50
	_96_v50 = _dafny.MultiSetOf(p1)
	var _97_v51 _dafny.Sequence
	_ = _97_v51
	_97_v51 = _dafny.SeqOf((func() bool {
		if !(p3) {
			return !(p3)
		}
		return p0
	})(), (_96_v50).IsProperSubsetOf(_96_v50))
	_97_v51 = _97_v51
	var _98_v52 *C0
	_ = _98_v52
	var _nw13 *C0 = New_C0_()
	_ = _nw13
	_nw13.Ctor__()
	_98_v52 = _nw13
	var _hi0 _dafny.Int = (_dafny.Zero).Minus((p1).Plus(p1))
	_ = _hi0
	for _99_i4 := p1; _99_i4.Cmp(_hi0) < 0; _99_i4 = _99_i4.Plus(_dafny.One) {
		var _100_v53 _dafny.Array
		_ = _100_v53
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_1
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = (func(_101_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_102_i5 _dafny.Int) _dafny.Int {
					return (_102_i5).Times(_101_p1)
				}
			})(p1)
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
		_100_v53 = _nw14
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))
		_ = _index9
		(_100_v53).ArraySet1((_dafny.IntOfInt64(751)).Minus(p1), (_index9).Int())
		var _103_v54 _dafny.Array
		_ = _103_v54
		var _nw15 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw15
		_103_v54 = _nw15
		var _104_v55 _dafny.Map
		_ = _104_v55
		_104_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v54, (_97_v51).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_97_v51).Cardinality()))).Uint32()).(bool))
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))
		_ = _index10
		var _rhs8 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.Int {
			if p0 {
				return p1
			}
			return _99_i4
		})()).Plus(Companion_Default___.SafeModuloInt(_99_i4, _99_i4)))
		_ = _rhs8
		var _rhs9 _dafny.Sequence = (func() _dafny.Sequence {
			if p3 {
				return _46_v12
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("br"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("br")).Cardinality()))).Uint32(), p2), _46_v12)
		})()
		_ = _rhs9
		var _rhs10 _dafny.Int = Companion_Default___.SafeModuloInt(_99_i4, p1)
		_ = _rhs10
		var _rhs11 _dafny.Map = _104_v55
		_ = _rhs11
		var _rhs12 _dafny.Int = Companion_Default___.SafeDivisionInt(_99_i4, (p1).Minus(p1))
		_ = _rhs12
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		var _lhs5 _dafny.Array = _100_v53
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))
		_ = _lhs6
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		_lhs3.F0 = _rhs8
		_lhs4.F2 = _rhs9
		(_lhs5).ArraySet1(_rhs10, (_lhs6).Int())
		_104_v55 = _rhs11
		_lhs7.F8 = _rhs12
		if false {
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_103_v54), 0))
			_ = _index11
			(_103_v54).ArraySet1(p3, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_103_v54), 0))
			_ = _index12
			(_103_v54).ArraySet1((func() bool {
				if !(!(p3)) {
					return !((p1).Cmp(_99_i4) != 0)
				}
				return (false) && (p0)
			})(), (_index12).Int())
			var _105_v56 _dafny.Map
			_ = _105_v56
			_105_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_103_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_103_v54), 0))).Int()).(bool))
			var _106_v57 _dafny.Map
			_ = _106_v57
			_106_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_100_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))).Int()).(_dafny.Int)), _105_v56)
			_106_v57 = (_106_v57).Update((_dafny.Zero).Minus(_99_i4), _105_v56)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_103_v54), 0))
			_ = _index13
			(_103_v54).ArraySet1(p3, (_index13).Int())
			var _107_v58 _dafny.Map
			_ = _107_v58
			_107_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(p0, _99_i4, _dafny.IntOfInt64(-914), globalState), p1)
			var _108_v60 _dafny.Map
			_ = _108_v60
			_108_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(736), p1)
			var _109_v61 _dafny.Set
			_ = _109_v61
			_109_v61 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("dxrmfw"), _46_v12)
			var _110_v62 D1
			_ = _110_v62
			_110_v62 = Companion_D1_.Create_DC2_(_109_v61, p2)
			var _111_v63 _dafny.Map
			_ = _111_v63
			_111_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v60, _110_v62)
			_107_v58 = (_107_v58).Update(func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_111_v63).Keys().Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _112_v59 _dafny.Map
					_112_v59 = interface{}(_compr_12).(_dafny.Map)
					if (_111_v63).Contains(_112_v59) {
						_coll12.Add(_112_v59, p0)
					}
				}
				return _coll12.ToMap()
			}(), p1)
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_103_v54), 0))
			_ = _index14
			(_103_v54).ArraySet1(p3, (_index14).Int())
		} else {
			var _113_v64 bool
			_ = _113_v64
			_113_v64 = true
			_113_v64 = ((_100_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))).Int()).(_dafny.Int)).Cmp(_99_i4) != 0
			_113_v64 = _dafny.Companion_Sequence_.IsProperPrefixOf(_45_v11, _dafny.Companion_Sequence_.Concatenate(_45_v11, _dafny.Companion_Sequence_.Update(_45_v11, (Companion_Default___.SafeIndex((_100_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_45_v11).Cardinality()))).Uint32(), p1)))
			var _114_v65 _dafny.Sequence
			_ = _114_v65
			_114_v65 = _dafny.SeqOf((_dafny.MultiSetFromSeq(_48_v14)).Intersection(_dafny.MultiSetOf(_98_v52)), _dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if p0 {
					return _48_v14
				}
				return _48_v14
			})()), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if p0 {
					return _dafny.SeqOf(_98_v52)
				}
				return _48_v14
			})(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p0 {
					return _dafny.SeqOf(_98_v52)
				}
				return _48_v14
			})()).Cardinality()))).Uint32(), _47_v13)))
			var _rhs13 *C0 = _47_v13
			_ = _rhs13
			var _rhs14 *C0 = _47_v13
			_ = _rhs14
			var _rhs15 _dafny.Sequence = _114_v65
			_ = _rhs15
			_47_v13 = _rhs13
			_98_v52 = _rhs14
			_114_v65 = _rhs15
			var _115_v66 D4
			_ = _115_v66
			_115_v66 = Companion_D4_.Create_DC8_()
			_115_v66 = _115_v66
			var _116_v67 _dafny.Array
			_ = _116_v67
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw16
			_116_v67 = _nw16
			var _117_v68 _dafny.Map
			_ = _117_v68
			_117_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.MultiSetOf((_97_v51).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_97_v51).Cardinality()))).Uint32()).(bool))).Cardinality())
			var _118_v69 _dafny.Sequence
			_ = _118_v69
			_118_v69 = _dafny.SeqOf(Companion_Default___.Fm11(_dafny.IntOfUint32((_45_v11).Cardinality()), (_100_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))).Int()).(_dafny.Int), _117_v68, globalState))
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_116_v67), 0))
			_ = _index15
			(_116_v67).ArraySet1(_118_v69, (_index15).Int())
			var _119_v70 _dafny.Set
			_ = _119_v70
			_119_v70 = _dafny.SetOf(p1)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_116_v67), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))
			_ = _index17
			var _rhs16 bool = (func() bool {
				if _dafny.Companion_Sequence_.IsPrefixOf(_46_v12, _46_v12) {
					return _113_v64
				}
				return p0
			})()
			_ = _rhs16
			var _rhs17 _dafny.Sequence = _118_v69
			_ = _rhs17
			var _rhs18 _dafny.Int = (_100_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))).Int()).(_dafny.Int)
			_ = _rhs18
			var _rhs19 _dafny.Int = Companion_Default___.SafeDivisionInt((_100_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))).Int()).(_dafny.Int), (_119_v70).Cardinality())
			_ = _rhs19
			var _lhs8 _dafny.Array = _116_v67
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_116_v67), 0))
			_ = _lhs9
			var _lhs10 _dafny.Array = _100_v53
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_100_v53), 0))
			_ = _lhs11
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			_113_v64 = _rhs16
			(_lhs8).ArraySet1(_rhs17, (_lhs9).Int())
			(_lhs10).ArraySet1(_rhs18, (_lhs11).Int())
			_lhs12.F0 = _rhs19
		}
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_103_v54), 0))
		_ = _index18
		(_103_v54).ArraySet1((_97_v51).Select((Companion_Default___.SafeIndex(_99_i4, _dafny.IntOfUint32((_97_v51).Cardinality()))).Uint32()).(bool), (_index18).Int())
		var _120_v71 _dafny.MultiSet
		_ = _120_v71
		_120_v71 = _dafny.MultiSetOf(_100_v53, _100_v53, _100_v53)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_103_v54), 0))
		_ = _index19
		(_103_v54).ArraySet1((_120_v71).IsDisjointFrom(_120_v71), (_index19).Int())
		var _121_v72 _dafny.Map
		_ = _121_v72
		_121_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0) == (p3), (_96_v50).IsProperSubsetOf(_96_v50))
		var _122_v73 _dafny.Map
		_ = _122_v73
		_122_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v13, (_103_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_103_v54), 0))).Int()).(bool))
		var _123_v74 D7
		_ = _123_v74
		_123_v74 = Companion_D7_.Create_DC14_(_122_v73)
		var _124_v75 _dafny.Map
		_ = _124_v75
		_124_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v50, (((_123_v74).Dtor_cf20()).Update(_47_v13, (_103_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_103_v54), 0))).Int()).(bool))).Cardinality())
		_121_v72 = (_121_v72).Update(((_124_v75).Cardinality()).Cmp((_dafny.Zero).Minus(p1)) < 0, (_103_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_103_v54), 0))).Int()).(bool))
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _125_v0 _dafny.Sequence
	_ = _125_v0
	_125_v0 = _dafny.UnicodeSeqOfUtf8Bytes("si")
	var _126_v1 _dafny.Array
	_ = _126_v1
	var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
	_ = _nw17
	_126_v1 = _nw17
	var _127_v2 bool
	_ = _127_v2
	_127_v2 = false
	var _128_v3 _dafny.Array
	_ = _128_v3
	var _nwElement0_2 bool = !(_127_v2)
	_ = _nwElement0_2
	var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(3))
	_ = _nw18
	(_nw18).ArraySet1(_nwElement0_2, 0)
	(_nw18).ArraySet1(!(_127_v2), 1)
	(_nw18).ArraySet1(_127_v2, 2)
	_128_v3 = _nw18
	var _129_v4 _dafny.CodePoint
	_ = _129_v4
	_129_v4 = _dafny.CodePoint('j')
	var _130_v5 _dafny.Set
	_ = _130_v5
	_130_v5 = _dafny.SetOf(_129_v4)
	var _131_globalState *GlobalState
	_ = _131_globalState
	var _nw19 *GlobalState = New_GlobalState_()
	_ = _nw19
	_nw19.Ctor__(_dafny.IntOfInt64(-164), _dafny.IntOfInt64(378), _125_v0, _126_v1, true, true, _128_v3, _dafny.IntOfInt64(223), _dafny.IntOfInt64(583), _130_v5, _dafny.IntOfInt64(782))
	_131_globalState = _nw19
	_127_v2 = !((!(_127_v2)) && (_127_v2))
	var _132_v6 _dafny.Int
	_ = _132_v6
	_132_v6 = _dafny.IntOfInt64(520)
	var _133_v7 _dafny.Map
	_ = _133_v7
	_133_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
		if _127_v2 {
			return _126_v1
		}
		return _126_v1
	})(), (func() _dafny.Int {
		if Companion_Default___.Fm0(_131_globalState) {
			return _132_v6
		}
		return _132_v6
	})())
	var _134_v8 _dafny.Sequence
	_ = _134_v8
	_134_v8 = _dafny.SeqOf(_127_v2)
	var _135_v9 _dafny.Set
	_ = _135_v9
	_135_v9 = _dafny.SetOf(_132_v6, _132_v6, _dafny.IntOfUint32((_134_v8).Cardinality()))
	_133_v7 = (_133_v7).Update(_126_v1, ((_135_v9).Cardinality()).Minus(_132_v6))
	for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_128_v3), 0))); ; {
		_guard_loop_0, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		var _136_i0 _dafny.Int
		_136_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_136_i0).Sign() != -1) && ((_136_i0).Cmp(_dafny.ArrayLen((_128_v3), 0)) < 0)) {
			(_128_v3).ArraySet1((_dafny.MultiSetOf(false, _127_v2, false, _127_v2)).IsSubsetOf((_dafny.MultiSetOf(false))), (_136_i0).Int())
		}
	}
	var _137_v10 _dafny.MultiSet
	_ = _137_v10
	_137_v10 = _dafny.MultiSetOf(!(_127_v2), _127_v2)
	var _138_v11 _dafny.MultiSet
	_ = _138_v11
	_138_v11 = _137_v10
	var _source2 _dafny.MultiSet = _138_v11
	_ = _source2
	var _139___mcc_h0 _dafny.MultiSet = _source2
	_ = _139___mcc_h0
	var _140_cf0 _dafny.MultiSet = _139___mcc_h0
	_ = _140_cf0
	if !(!(_127_v2) || (!((_135_v9).IsProperSubsetOf(_135_v9)))) {
		Companion_Default___.M0((func() bool {
			if _127_v2 {
				return Companion_Default___.Fm0(_131_globalState)
			}
			return (_134_v8).Select((Companion_Default___.SafeIndex(_132_v6, _dafny.IntOfUint32((_134_v8).Cardinality()))).Uint32()).(bool)
		})(), (_dafny.MultiSetFromSeq(_134_v8)).Cardinality(), _129_v4, _127_v2, _131_globalState)
		var _141_v12 _dafny.Set
		_ = _141_v12
		_141_v12 = _dafny.SetOf(false)
		var _142_v13 _dafny.Sequence
		_ = _142_v13
		_142_v13 = _dafny.SeqOf(_dafny.IntOfUint32((_134_v8).Cardinality()))
		var _143_v14 _dafny.Map
		_ = _143_v14
		_143_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_135_v9).Difference(_dafny.SetOf((_141_v12).Cardinality(), _dafny.IntOfInt64(-308))), ((_142_v13).Select((Companion_Default___.SafeIndex(_132_v6, _dafny.IntOfUint32((_142_v13).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_132_v6))
		var _rhs20 bool = Companion_Default___.Fm0(_131_globalState)
		_ = _rhs20
		var _rhs21 _dafny.Int = _132_v6
		_ = _rhs21
		var _rhs22 _dafny.Int = _132_v6
		_ = _rhs22
		var _rhs23 _dafny.Map = (_143_v14).Update(Companion_Default___.Fm1(_dafny.CodePoint('d'), _127_v2, _127_v2, _131_globalState), _dafny.IntOfInt64(80))
		_ = _rhs23
		var _lhs13 *GlobalState = _131_globalState
		_ = _lhs13
		var _lhs14 *GlobalState = _131_globalState
		_ = _lhs14
		_127_v2 = _rhs20
		_lhs13.F8 = _rhs21
		_lhs14.F8 = _rhs22
		_143_v14 = _rhs23
		_127_v2 = _127_v2
		(_131_globalState).F0 = _132_v6
		var _144_v15 _dafny.MultiSet
		_ = _144_v15
		_144_v15 = _dafny.MultiSetOf(_134_v8, _dafny.SeqOf(_127_v2, _127_v2))
		var _145_v16 _dafny.Map
		_ = _145_v16
		_145_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v15, _125_v0)
		var _146_v17 _dafny.Sequence
		_ = _146_v17
		_146_v17 = _dafny.SeqOf(_144_v15, _144_v15, _144_v15)
		_145_v16 = (_145_v16).Update((_146_v17).Select((Companion_Default___.SafeIndex(_132_v6, _dafny.IntOfUint32((_146_v17).Cardinality()))).Uint32()).(_dafny.MultiSet), _125_v0)
	} else {
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_126_v1), 0))
		_ = _index20
		(_126_v1).ArraySet1CodePoint(_129_v4, (_index20).Int())
		var _147_v18 _dafny.Map
		_ = _147_v18
		_147_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _132_v6)
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_126_v1), 0))
		_ = _index21
		(_126_v1).ArraySet1CodePoint(Companion_Default___.Fm2(_147_v18, _127_v2, _132_v6, _131_globalState), (_index21).Int())
		var _148_v19 _dafny.Set
		_ = _148_v19
		_148_v19 = _dafny.SetOf(_127_v2)
		var _149_v20 _dafny.Map
		_ = _149_v20
		_149_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(430))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_150_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_151_i1 _dafny.Int) _dafny.CodePoint {
				return _150_v4
			}
		})(_129_v4))))
		var _152_v21 _dafny.Map
		_ = _152_v21
		_152_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v19, !(_149_v20).Equals(_149_v20))
		_152_v21 = (_152_v21).Update(_148_v19, _127_v2)
		(_131_globalState).F0 = _132_v6
		(_131_globalState).F2 = _125_v0
		Companion_Default___.M0((_134_v8).Select((Companion_Default___.SafeIndex(_132_v6, _dafny.IntOfUint32((_134_v8).Cardinality()))).Uint32()).(bool), _132_v6, _129_v4, _127_v2, _131_globalState)
	}
	(_131_globalState).F0 = _132_v6
	var _153_v22 D1
	_ = _153_v22
	_153_v22 = Companion_D1_.Create_DC1_(true)
	if (_153_v22).Dtor_cf1() {
		var _154_v23 _dafny.Array
		_ = _154_v23
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw20
		_154_v23 = _nw20
		var _155_v24 _dafny.Map
		_ = _155_v24
		_155_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v3, _127_v2)
		var _156_v25 _dafny.Map
		_ = _156_v25
		_156_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v23, (_132_v6).Times((_155_v24).Cardinality()))
		_156_v25 = (_156_v25).Update(_154_v23, _132_v6)
		_134_v8 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm3(_131_globalState), _134_v8)
		(_131_globalState).F0 = (_132_v6).Minus(_132_v6)
		var _157_v26 *C0
		_ = _157_v26
		var _nw21 *C0 = New_C0_()
		_ = _nw21
		_nw21.Ctor__()
		_157_v26 = _nw21
		var _158_v27 _dafny.Array
		_ = _158_v27
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
		_ = _nw22
		_158_v27 = _nw22
		var _159_v28 _dafny.Map
		_ = _159_v28
		_159_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v2, _125_v0)
		var _160_v29 _dafny.Map
		_ = _160_v29
		_160_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v26, (func() _dafny.Sequence {
			if (_159_v28).Contains(!(_127_v2)) {
				return (_159_v28).Get(!(_127_v2)).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(23))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_161_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_162_i2 _dafny.Int) _dafny.CodePoint {
					return _161_v4
				}
			})(_129_v4)))
		})())
		var _163_v30 _dafny.Array
		_ = _163_v30
		var _nwElement0_3 _dafny.Sequence = _dafny.SeqOf(_129_v4, _dafny.CodePoint('g'), _129_v4, _129_v4)
		_ = _nwElement0_3
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(8))
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_3, 0)
		(_nw23).ArraySet1((func() _dafny.Sequence {
			if (_160_v29).Contains(_157_v26) {
				return (_160_v29).Get(_157_v26).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-93))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_164_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_165_i3 _dafny.Int) _dafny.CodePoint {
					return _164_v4
				}
			})(_129_v4)))
		})(), 1)
		(_nw23).ArraySet1(_125_v0, 2)
		(_nw23).ArraySet1(_125_v0, 3)
		(_nw23).ArraySet1(_125_v0, 4)
		(_nw23).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_166_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		})), 5)
		(_nw23).ArraySet1(_125_v0, 6)
		(_nw23).ArraySet1(_125_v0, 7)
		_163_v30 = _nw23
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_158_v27), 0))
		_ = _index22
		(_158_v27).ArraySet1(_163_v30, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_126_v1), 0))
		_ = _index23
		(_126_v1).ArraySet1CodePoint(_129_v4, (_index23).Int())
		var _167_v31 _dafny.Array
		_ = _167_v31
		var _nw24 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
		_ = _nw24
		_167_v31 = _nw24
		var _168_v32 _dafny.Map
		_ = _168_v32
		_168_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v31, _163_v30)
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_158_v27), 0))
		_ = _index24
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_126_v1), 0))
		_ = _index25
		var _rhs24 _dafny.Array = (func() _dafny.Array {
			if (_168_v32).Contains(_167_v31) {
				return (_168_v32).Get(_167_v31).(_dafny.Array)
			}
			return _163_v30
		})()
		_ = _rhs24
		var _rhs25 _dafny.Array = _154_v23
		_ = _rhs25
		var _rhs26 _dafny.CodePoint = _129_v4
		_ = _rhs26
		var _rhs27 _dafny.Int = _132_v6
		_ = _rhs27
		var _lhs15 _dafny.Array = _158_v27
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_158_v27), 0))
		_ = _lhs16
		var _lhs17 _dafny.Array = _126_v1
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(325), _dafny.ArrayLen((_126_v1), 0))
		_ = _lhs18
		var _lhs19 *GlobalState = _131_globalState
		_ = _lhs19
		(_lhs15).ArraySet1(_rhs24, (_lhs16).Int())
		_154_v23 = _rhs25
		(_lhs17).ArraySet1CodePoint(_rhs26, (_lhs18).Int())
		_lhs19.F0 = _rhs27
	} else {
		var _169_v33 _dafny.Array
		_ = _169_v33
		var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw25
		_169_v33 = _nw25
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_169_v33), 0))
		_ = _index26
		(_169_v33).ArraySet1(_126_v1, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_169_v33), 0))
		_ = _index27
		var _rhs28 _dafny.Array = _126_v1
		_ = _rhs28
		var _rhs29 _dafny.Int = (_dafny.Zero).Minus((_132_v6).Plus(_132_v6))
		_ = _rhs29
		var _lhs20 _dafny.Array = _169_v33
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_169_v33), 0))
		_ = _lhs21
		var _lhs22 *GlobalState = _131_globalState
		_ = _lhs22
		(_lhs20).ArraySet1(_rhs28, (_lhs21).Int())
		_lhs22.F8 = _rhs29
		var _170_v34 _dafny.Map
		_ = _170_v34
		_170_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v2, _132_v6)
		_170_v34 = (_170_v34).Update(_127_v2, _132_v6)
		var _171_v35 _dafny.Map
		_ = _171_v35
		_171_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v2, _127_v2)
		var _172_v36 _dafny.Sequence
		_ = _172_v36
		_172_v36 = _dafny.SeqOf(_135_v9)
		var _173_v37 _dafny.Map
		_ = _173_v37
		_173_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _dafny.IntOfUint32((_125_v0).Cardinality()))
		var _174_v38 _dafny.MultiSet
		_ = _174_v38
		_174_v38 = _dafny.MultiSetOf(_129_v4)
		var _rhs30 _dafny.Set = (_172_v36).Select((Companion_Default___.SafeIndex((_132_v6).Times(_132_v6), _dafny.IntOfUint32((_172_v36).Cardinality()))).Uint32()).(_dafny.Set)
		_ = _rhs30
		var _rhs31 _dafny.Int = Companion_Default___.Fm4(_127_v2, _132_v6, _131_globalState)
		_ = _rhs31
		var _rhs32 _dafny.Int = Companion_Default___.SafeDivisionInt(_132_v6, _132_v6)
		_ = _rhs32
		var _rhs33 bool = !(((_173_v37).Cardinality()).Cmp((_174_v38).Cardinality()) == 0) || ((_127_v2) && (_127_v2))
		_ = _rhs33
		var _rhs34 _dafny.Map = ((_171_v35).Merge((_171_v35).Update(_127_v2, true))).Update(Companion_Default___.Fm0(_131_globalState), _127_v2)
		_ = _rhs34
		var _lhs23 *GlobalState = _131_globalState
		_ = _lhs23
		_135_v9 = _rhs30
		_132_v6 = _rhs31
		_lhs23.F0 = _rhs32
		_127_v2 = _rhs33
		_171_v35 = _rhs34
		var _175_v39 _dafny.Sequence
		_ = _175_v39
		_175_v39 = _dafny.SeqOf(_173_v37)
		var _176_v40 _dafny.Map
		_ = _176_v40
		_176_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _127_v2)
		var _177_v41 _dafny.Map
		_ = _177_v41
		_177_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((func() bool {
			if (_176_v40).Contains(_132_v6) {
				return (_176_v40).Get(_132_v6).(bool)
			}
			return _127_v2
		})())).Cardinality(), _127_v2)
		var _178_v43 _dafny.Array
		_ = _178_v43
		var _nwElement0_4 _dafny.Map = _173_v37
		_ = _nwElement0_4
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(10))
		_ = _nw26
		(_nw26).ArraySet1(_nwElement0_4, 0)
		(_nw26).ArraySet1((_175_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.IntOfUint32((_175_v39).Cardinality()))).Uint32()).(_dafny.Map), 1)
		(_nw26).ArraySet1(Companion_Default___.Fm5(_dafny.IntOfInt64(183), _127_v2, _131_globalState), 2)
		(_nw26).ArraySet1((_173_v37).Merge(_173_v37), 3)
		(_nw26).ArraySet1(_173_v37, 4)
		(_nw26).ArraySet1(_173_v37, 5)
		(_nw26).ArraySet1((Companion_Default___.Fm5(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aywqiw")).Cardinality()), _127_v2, _131_globalState)).Merge(_173_v37), 6)
		(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(377), (func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter14 := _dafny.Iterate((_177_v41).Keys().Elements()); ; {
				_compr_13, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _179_v42 _dafny.Int
				_179_v42 = interface{}(_compr_13).(_dafny.Int)
				if (_177_v41).Contains(_179_v42) {
					_coll13.Add((_179_v42).Times(_dafny.IntOfInt64(51)))
				}
			}
			return _coll13.ToSet()
		}()).Cardinality()), 7)
		(_nw26).ArraySet1((_173_v37).Merge(Companion_Default___.Fm5(_132_v6, _127_v2, _131_globalState)), 8)
		(_nw26).ArraySet1(_173_v37, 9)
		_178_v43 = _nw26
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_178_v43), 0))
		_ = _index28
		(_178_v43).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _dafny.IntOfUint32((_134_v8).Cardinality())), (_index28).Int())
		var _180_v44 _dafny.Set
		_ = _180_v44
		_180_v44 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("nagnmac"))
		var _181_v45 D1
		_ = _181_v45
		_181_v45 = Companion_D1_.Create_DC2_(_180_v44, _129_v4)
		var _182_v46 _dafny.Array
		_ = _182_v46
		var _nw27 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
		_ = _nw27
		_182_v46 = _nw27
		var _183_v47 _dafny.Map
		_ = _183_v47
		_183_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_cf0, (func() _dafny.Array {
			if _127_v2 {
				return _182_v46
			}
			return _182_v46
		})())
		var _184_v48 _dafny.Map
		_ = _184_v48
		_184_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _173_v37)
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_178_v43), 0))
		_ = _index29
		var _rhs35 _dafny.Map = ((_173_v37).Update(_132_v6, _132_v6)).Merge((func() _dafny.Map {
			if (_184_v48).Contains(_132_v6) {
				return (_184_v48).Get(_132_v6).(_dafny.Map)
			}
			return _173_v37
		})())
		_ = _rhs35
		var _rhs36 D1 = _181_v45
		_ = _rhs36
		var _rhs37 _dafny.Map = (_183_v47).Merge((_183_v47).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v11, _182_v46)))
		_ = _rhs37
		var _rhs38 bool = (_132_v6).Cmp(_132_v6) >= 0
		_ = _rhs38
		var _lhs24 _dafny.Array = _178_v43
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_178_v43), 0))
		_ = _lhs25
		(_lhs24).ArraySet1(_rhs35, (_lhs25).Int())
		_181_v45 = _rhs36
		_183_v47 = _rhs37
		_127_v2 = _rhs38
		_127_v2 = false
	}
	Companion_Default___.M0(_127_v2, (_dafny.Zero).Minus(_132_v6), _129_v4, (_dafny.IntOfInt64(-278)).Cmp(_132_v6) >= 0, _131_globalState)
	_127_v2 = _127_v2
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
	_ = _index30
	(_128_v3).ArraySet1(_127_v2, (_index30).Int())
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
	_ = _index31
	(_128_v3).ArraySet1(!(_127_v2), (_index31).Int())
	var _185_v49 _dafny.Sequence
	_ = _185_v49
	_185_v49 = _dafny.SeqOf(Companion_Default___.Fm4((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool), _132_v6, _131_globalState))
	var _186_v50 _dafny.Map
	_ = _186_v50
	_186_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v49, _132_v6)
	var _187_v51 _dafny.Map
	_ = _187_v51
	_187_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_125_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.IntOfUint32((_125_v0).Cardinality()))).Uint32(), _129_v4)).Cardinality()))
	var _188_v52 _dafny.Map
	_ = _188_v52
	_188_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v50, _187_v51)
	_188_v52 = (_188_v52).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v49, _132_v6)).Merge(func() _dafny.Map {
		var _coll14 = _dafny.NewMapBuilder()
		_ = _coll14
		for _iter15 := _dafny.Iterate((_186_v50).Keys().Elements()); ; {
			_compr_14, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _189_v53 _dafny.Sequence
			_189_v53 = interface{}(_compr_14).(_dafny.Sequence)
			if (_186_v50).Contains(_189_v53) {
				_coll14.Add(_189_v53, (_135_v9).Cardinality())
			}
		}
		return _coll14.ToMap()
	}()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool), _132_v6))
	if (_127_v2) || ((_132_v6).Cmp(_132_v6) >= 0) {
		(_131_globalState).F8 = _132_v6
		var _190_v54 _dafny.Map
		_ = _190_v54
		_190_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v3, _127_v2)
		_190_v54 = ((_190_v54).Merge(_190_v54)).Update(_128_v3, true)
		var _191_v55 _dafny.Sequence
		_ = _191_v55
		_191_v55 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pnjsvda"))
		var _192_v56 _dafny.Sequence
		_ = _192_v56
		_192_v56 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_134_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_191_v55).Cardinality()), _dafny.IntOfUint32((_134_v8).Cardinality()))).Uint32(), _127_v2), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_125_v0).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_134_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_191_v55).Cardinality()), _dafny.IntOfUint32((_134_v8).Cardinality()))).Uint32(), _127_v2)).Cardinality()))).Uint32(), false)
		_134_v8 = (_192_v56)
		var _193_v57 *C0
		_ = _193_v57
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__()
		_193_v57 = _nw28
		var _194_v58 _dafny.Sequence
		_ = _194_v58
		_194_v58 = _dafny.SeqOf(_193_v57)
		var _195_v59 _dafny.Set
		_ = _195_v59
		_195_v59 = _dafny.SetOf(_193_v57)
		var _196_v60 _dafny.Map
		_ = _196_v60
		_196_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _193_v57)
		var _197_v61 D3
		_ = _197_v61
		_197_v61 = Companion_D3_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(244))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_198_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_199_i5 _dafny.Int) _dafny.CodePoint {
				return _198_v4
			}
		})(_129_v4))))
		var _200_v62 _dafny.Sequence
		_ = _200_v62
		_200_v62 = _dafny.SeqOf(_138_v11, Companion_Default___.Fm6(_dafny.CodePoint('i'), (_197_v61).Dtor_cf5(), _127_v2, _131_globalState))
		Companion_Default___.M0((_dafny.SetOf((_194_v58).Select((Companion_Default___.SafeIndex(_132_v6, _dafny.IntOfUint32((_194_v58).Cardinality()))).Uint32()).(*C0), _193_v57, _193_v57, _193_v57, _193_v57)).IsDisjointFrom(_195_v59), ((_196_v60).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _193_v57))).Cardinality(), _129_v4, !_dafny.Companion_Sequence_.Equal(_200_v62, _dafny.SeqOf(_138_v11, _138_v11)), _131_globalState)
		var _201_v63 _dafny.Array
		_ = _201_v63
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_2
		var _nw29 _dafny.Array
		_ = _nw29
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw29 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_202_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_203_i6 _dafny.Int) _dafny.Int {
					return (_203_i6).Times(_202_v6)
				}
			})(_132_v6)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw29 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw29).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw29).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_201_v63 = _nw29
		_201_v63 = _201_v63
	} else {
		var _204_v64 _dafny.Map
		_ = _204_v64
		_204_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _132_v6)
		_204_v64 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_125_v0).Cardinality()), _132_v6)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _132_v6))
		var _205_v65 D3
		_ = _205_v65
		_205_v65 = Companion_D3_.Create_DC4_(_125_v0)
		var _206_v66 _dafny.Sequence
		_ = _206_v66
		_206_v66 = _dafny.SeqOf(_205_v65, _205_v65)
		var _207_v67 _dafny.Map
		_ = _207_v67
		_207_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_131_globalState), _206_v66)
		var _208_v68 _dafny.Map
		_ = _208_v68
		_208_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)), (func() _dafny.Sequence {
			if (_207_v67).Contains((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)) {
				return (_207_v67).Get((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_205_v65)
		})())
		var _209_v69 D4
		_ = _209_v69
		_209_v69 = Companion_D4_.Create_DC6_(_206_v66)
		_207_v67 = (_207_v67).Update((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool), (_209_v69).Dtor_cf8())
		var _210_v70 _dafny.Map
		_ = _210_v70
		_210_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _127_v2)
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _index32
		(_128_v3).ArraySet1(!(_210_v70).Equals(Companion_Default___.Fm7(_131_globalState)), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _index33
		(_128_v3).ArraySet1(!(_127_v2), (_index33).Int())
		_135_v9 = _135_v9
	}
	_127_v2 = _127_v2
	var _211_v71 _dafny.Array
	_ = _211_v71
	var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
	_ = _nw30
	_211_v71 = _nw30
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_211_v71), 0))); ; {
		_guard_loop_1, _ok16 := _iter16()
		if !_ok16 {
			break
		}
		var _212_i7 _dafny.Int
		_212_i7 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_212_i7).Sign() != -1) && ((_212_i7).Cmp(_dafny.ArrayLen((_211_v71), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_211_v71, (_212_i7).Int(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v8, (_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool))).Merge(func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter17 := _dafny.Iterate((func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter18 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(862))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}(func(_213_i8 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(false)
					}))).Elements()); ; {
						_compr_16, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _214_v73 _dafny.Sequence
						_214_v73 = interface{}(_compr_16).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(862))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}(func(_213_i8 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(false)
						})), _214_v73) {
							_coll16.Add(_214_v73, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _132_v6))
						}
					}
					return _coll16.ToMap()
				}()).Keys().Elements()); ; {
					_compr_15, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _215_v72 _dafny.Sequence
					_215_v72 = interface{}(_compr_15).(_dafny.Sequence)
					if (func() _dafny.Map {
						var _coll17 = _dafny.NewMapBuilder()
						_ = _coll17
						for _iter19 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(862))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg19 _dafny.Int) interface{} {
								return coer19(arg19)
							}
						}(func(_213_i8 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(false)
						}))).Elements()); ; {
							_compr_17, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _214_v73 _dafny.Sequence
							_214_v73 = interface{}(_compr_17).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(862))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg20 _dafny.Int) interface{} {
									return coer20(arg20)
								}
							}(func(_213_i8 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqOf(false)
							})), _214_v73) {
								_coll17.Add(_214_v73, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _132_v6))
							}
						}
						return _coll17.ToMap()
					}()).Contains(_215_v72) {
						_coll15.Add(_215_v72, (_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool))
					}
				}
				return _coll15.ToMap()
			}())).Merge((func() _dafny.Map {
				if (_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool) {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v8, !((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)))
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v8, (_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool))
			})())))
		}
	}
	for _iter20 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok20 := _iter20()
		if !_ok20 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	var _216_v74 *C0
	_ = _216_v74
	var _nw31 *C0 = New_C0_()
	_ = _nw31
	_nw31.Ctor__()
	_216_v74 = _nw31
	var _217_v75 _dafny.Map
	_ = _217_v75
	_217_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v6, _132_v6)
	_125_v0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(_132_v6, _217_v75, _dafny.IntOfUint32((_185_v49).Cardinality()), _131_globalState), _dafny.Companion_Sequence_.Concatenate(_125_v0, _125_v0))
	var _218_v76 _dafny.Map
	_ = _218_v76
	_218_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v74, _132_v6)
	_218_v76 = (_218_v76).Update(_216_v74, _132_v6)
	_127_v2 = (_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)
	var _219_v77 _dafny.Array
	_ = _219_v77
	var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
	_ = _nw32
	_219_v77 = _nw32
	_219_v77 = _219_v77
	var _220_v78 _dafny.MultiSet
	_ = _220_v78
	_220_v78 = _dafny.MultiSetOf(_132_v6)
	var _hi1 _dafny.Int = ((_220_v78).Cardinality()).Times(_132_v6)
	_ = _hi1
	for _221_i9 := Companion_Default___.Fm4(false, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_134_v8, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.IntOfUint32((_134_v8).Cardinality()))).Uint32(), (Companion_D1_.Create_DC1_((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool))).Dtor_cf1())).Cardinality()), _131_globalState); _221_i9.Cmp(_hi1) < 0; _221_i9 = _221_i9.Plus(_dafny.One) {
		var _222_v79 _dafny.Sequence
		_ = _222_v79
		_222_v79 = _dafny.SeqOf(_128_v3, _128_v3, _128_v3)
		_222_v79 = _dafny.SeqOf(_128_v3)
		var _223_v80 _dafny.Array
		_ = _223_v80
		_223_v80 = _219_v77
		var _224_v81 _dafny.Sequence
		_ = _224_v81
		_224_v81 = _dafny.SeqOf(_219_v77, (_223_v80), _219_v77, _219_v77)
		_224_v81 = _224_v81
		var _225_v82 D3
		_ = _225_v82
		_225_v82 = Companion_D3_.Create_DC5_(_127_v2, (_185_v49).Select((Companion_Default___.SafeIndex(_221_i9, _dafny.IntOfUint32((_185_v49).Cardinality()))).Uint32()).(_dafny.Int))
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _index34
		(_128_v3).ArraySet1((_132_v6).Cmp(((_225_v82).Dtor_cf7()).Times(_dafny.IntOfInt64(-792))) > 0, (_index34).Int())
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _index35
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _index36
		var _rhs39 bool = !(_127_v2)
		_ = _rhs39
		var _rhs40 bool = ((_137_v10).Intersection(_137_v10)).IsProperSubsetOf(_137_v10)
		_ = _rhs40
		var _rhs41 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)), _134_v8), _dafny.SeqOf(_127_v2, (_128_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))).Int()).(bool)))).Cardinality())
		_ = _rhs41
		var _rhs42 _dafny.Set = (_dafny.SetOf(_132_v6)).Difference(_135_v9)
		_ = _rhs42
		var _rhs43 *C0 = _216_v74
		_ = _rhs43
		var _lhs26 _dafny.Array = _128_v3
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _lhs27
		var _lhs28 _dafny.Array = _128_v3
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_128_v3), 0))
		_ = _lhs29
		var _lhs30 *GlobalState = _131_globalState
		_ = _lhs30
		(_lhs26).ArraySet1(_rhs39, (_lhs27).Int())
		(_lhs28).ArraySet1(_rhs40, (_lhs29).Int())
		_lhs30.F8 = _rhs41
		_135_v9 = _rhs42
		_216_v74 = _rhs43
	}
	_dafny.Print(_125_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v1).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v1).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v3).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v3).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_129_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v5).Equals(_dafny.SetOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_globalState.F2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F3()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F3()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F6()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F6()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F9()).Equals(_dafny.SetOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_132_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v7).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_134_v8, _dafny.SeqOf(true, true, false, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v9).Equals(_dafny.SetOf(_dafny.IntOfInt64(520), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v10).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v11).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_185_v49, _dafny.SeqOf(_dafny.IntOfInt64(549))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v50).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(549)), _dafny.IntOfInt64(520))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v51).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v52).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(549)), _dafny.IntOfInt64(520)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(2))).UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(549)), _dafny.IntOfInt64(2)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(520)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_211_v71).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true, false, false, true), false).UpdateUnsafe(_dafny.SeqOf(false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_211_v71).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true, false, false, true), false).UpdateUnsafe(_dafny.SeqOf(false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_211_v71).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true, false, false, true), false).UpdateUnsafe(_dafny.SeqOf(false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_211_v71).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true, false, false, true), false).UpdateUnsafe(_dafny.SeqOf(false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_217_v75).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(520), _dafny.IntOfInt64(520))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v76).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v78).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(520))))
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

func (CompanionStruct_D0_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D0) Dtor_cf0() _dafny.MultiSet {
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
			return ok && data1.Cf0.Equals(data2.Cf0)
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
	Cf3 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Set, Cf3 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf2, Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 bool
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 bool) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.EmptySet, _dafny.CodePoint('D'))
}

func (_this D1) Dtor_cf2() _dafny.Set {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf1() bool {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf2.Equals(data2.Cf2) && data1.Cf3 == data2.Cf3
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

type D2_DC3 struct {
	Cf4 _dafny.Sequence
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf4 _dafny.Sequence) D2 {
	return D2{D2_DC3{Cf4}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D2) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4)
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

type D3_DC5 struct {
	Cf6 bool
	Cf7 _dafny.Int
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf6 bool, Cf7 _dafny.Int) D3 {
	return D3{D3_DC5{Cf6, Cf7}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

type D3_DC4 struct {
	Cf5 _dafny.Sequence
}

func (D3_DC4) isD3() {}

func (CompanionStruct_D3_) Create_DC4_(Cf5 _dafny.Sequence) D3 {
	return D3{D3_DC4{Cf5}}
}

func (_this D3) Is_DC4() bool {
	_, ok := _this.Get_().(D3_DC4)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC5_(false, _dafny.Zero)
}

func (_this D3) Dtor_cf6() bool {
	return _this.Get_().(D3_DC5).Cf6
}

func (_this D3) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D3_DC5).Cf7
}

func (_this D3) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D3_DC4).Cf5
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D3_DC4:
		{
			return "D3.DC4" + "(" + data.Cf5.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0
		}
	case D3_DC4:
		{
			data2, ok := other.Get_().(D3_DC4)
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D4_DC7 struct {
	Cf9  bool
	Cf10 *C0
	Cf11 bool
	Cf12 bool
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf9 bool, Cf10 *C0, Cf11 bool, Cf12 bool) D4 {
	return D4{D4_DC7{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

type D4_DC8 struct {
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_() D4 {
	return D4{D4_DC8{}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC6 struct {
	Cf8 _dafny.Sequence
}

func (D4_DC6) isD4() {}

func (CompanionStruct_D4_) Create_DC6_(Cf8 _dafny.Sequence) D4 {
	return D4{D4_DC6{Cf8}}
}

func (_this D4) Is_DC6() bool {
	_, ok := _this.Get_().(D4_DC6)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC7_(false, (*C0)(nil), false, false)
}

func (_this D4) Dtor_cf9() bool {
	return _this.Get_().(D4_DC7).Cf9
}

func (_this D4) Dtor_cf10() *C0 {
	return _this.Get_().(D4_DC7).Cf10
}

func (_this D4) Dtor_cf11() bool {
	return _this.Get_().(D4_DC7).Cf11
}

func (_this D4) Dtor_cf12() bool {
	return _this.Get_().(D4_DC7).Cf12
}

func (_this D4) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D4_DC6).Cf8
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8"
		}
	case D4_DC6:
		{
			return "D4.DC6" + "(" + _dafny.String(data.Cf8) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D4_DC8:
		{
			_, ok := other.Get_().(D4_DC8)
			return ok
		}
	case D4_DC6:
		{
			data2, ok := other.Get_().(D4_DC6)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf13 _dafny.Array
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf13 _dafny.Array) D5 {
	return D5{D5_DC9{Cf13}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D5_DC9).Cf13
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC9:
		{
			return "D5.DC9" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf13 == data2.Cf13
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
	Cf16 _dafny.Int
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf15 _dafny.Int, Cf16 _dafny.Int) D6 {
	return D6{D6_DC11{Cf15, Cf16}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

type D6_DC12 struct {
	Cf17 *C0
	Cf18 bool
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf17 *C0, Cf18 bool) D6 {
	return D6{D6_DC12{Cf17, Cf18}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC13 struct {
	Cf19 _dafny.Int
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf19 _dafny.Int) D6 {
	return D6{D6_DC13{Cf19}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC10 struct {
	Cf14 _dafny.Sequence
}

func (D6_DC10) isD6() {}

func (CompanionStruct_D6_) Create_DC10_(Cf14 _dafny.Sequence) D6 {
	return D6{D6_DC10{Cf14}}
}

func (_this D6) Is_DC10() bool {
	_, ok := _this.Get_().(D6_DC10)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC11_(_dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf15
}

func (_this D6) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf16
}

func (_this D6) Dtor_cf17() *C0 {
	return _this.Get_().(D6_DC12).Cf17
}

func (_this D6) Dtor_cf18() bool {
	return _this.Get_().(D6_DC12).Cf18
}

func (_this D6) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf19
}

func (_this D6) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D6_DC10).Cf14
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0
		}
	case D6_DC10:
		{
			data2, ok := other.Get_().(D6_DC10)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D7_DC15 struct {
	Cf21 _dafny.Int
	Cf22 _dafny.Int
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf21 _dafny.Int, Cf22 _dafny.Int) D7 {
	return D7{D7_DC15{Cf21, Cf22}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC16 struct {
	Cf23 _dafny.Int
	Cf24 _dafny.Int
	Cf25 bool
	Cf26 bool
	Cf27 _dafny.Int
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf23 _dafny.Int, Cf24 _dafny.Int, Cf25 bool, Cf26 bool, Cf27 _dafny.Int) D7 {
	return D7{D7_DC16{Cf23, Cf24, Cf25, Cf26, Cf27}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC14 struct {
	Cf20 _dafny.Map
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf20 _dafny.Map) D7 {
	return D7{D7_DC14{Cf20}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

type D7_DC17 struct {
	Cf28 D7
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf28 D7) D7 {
	return D7{D7_DC17{Cf28}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC15_(_dafny.Zero, _dafny.Zero)
}

func (_this D7) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf21
}

func (_this D7) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf22
}

func (_this D7) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D7_DC16).Cf23
}

func (_this D7) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D7_DC16).Cf24
}

func (_this D7) Dtor_cf25() bool {
	return _this.Get_().(D7_DC16).Cf25
}

func (_this D7) Dtor_cf26() bool {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D7_DC16).Cf27
}

func (_this D7) Dtor_cf20() _dafny.Map {
	return _this.Get_().(D7_DC14).Cf20
}

func (_this D7) Dtor_cf28() D7 {
	return _this.Get_().(D7_DC17).Cf28
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf20) + ")"
		}
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
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf20.Equals(data2.Cf20)
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
	Cf29 _dafny.Set
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf29 _dafny.Set) D8 {
	return D8{D8_DC18{Cf29}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D8) Dtor_cf29() _dafny.Set {
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F2   _dafny.Sequence
	F8   _dafny.Int
	_f1  _dafny.Int
	_f3  _dafny.Array
	_f4  bool
	_f5  bool
	_f6  _dafny.Array
	_f7  _dafny.Int
	_f9  _dafny.Set
	_f10 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F2 = _dafny.EmptySeq
	_this.F8 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = false
	_this._f5 = false
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f7 = _dafny.Zero
	_this._f9 = _dafny.EmptySet
	_this._f10 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.Sequence, f3 _dafny.Array, f4 bool, f5 bool, f6 _dafny.Array, f7 _dafny.Int, f8 _dafny.Int, f9 _dafny.Set, f10 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
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
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() _dafny.Set {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
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
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
