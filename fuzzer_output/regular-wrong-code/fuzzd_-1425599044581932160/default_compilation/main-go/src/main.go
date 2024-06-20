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
func (_static *CompanionStruct_Default___) Fm6(p0 bool, globalState *GlobalState) D0 {
	if false {
		return Companion_D0_.Create_DC0_(true)
	} else {
		return Companion_D0_.Create_DC0_(true)
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D3 {
	if false {
		return Companion_D3_.Create_DC8_(_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(615))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))))
	} else {
		return Companion_D3_.Create_DC8_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer1 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i1 _dafny.Int) D1 {
			return Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-926))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_2_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qeexk"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(105))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tcrjsmo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(995))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	}))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("acxmpij"), _dafny.UnicodeSeqOfUtf8Bytes("xt"))).Cardinality()), (_dafny.Zero).Minus((_dafny.SetOf(true, !(false), true, false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(500))).Cardinality())))).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(410))).Cardinality()))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_dafny.IntOfInt64(115)))
	})())
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	var _source0 D1 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("kbtb"))
	_ = _source0
	if _source0.Is_DC4() {
		var _5___mcc_h0 _dafny.Int = _source0.Get_().(D1_DC4).Cf5
		_ = _5___mcc_h0
		var _6___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC4).Cf6
		_ = _6___mcc_h1
		var _7_cf6 _dafny.Int = _6___mcc_h1
		_ = _7_cf6
		var _8_cf5 _dafny.Int = _5___mcc_h0
		_ = _8_cf5
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_8_cf5), _dafny.SeqOf(_8_cf5, _7_cf6, _8_cf5, _8_cf5))
	} else if _source0.Is_DC5() {
		var _9___mcc_h2 _dafny.Int = _source0.Get_().(D1_DC5).Cf7
		_ = _9___mcc_h2
		var _10___mcc_h3 _dafny.Int = _source0.Get_().(D1_DC5).Cf8
		_ = _10___mcc_h3
		var _11___mcc_h4 bool = _source0.Get_().(D1_DC5).Cf9
		_ = _11___mcc_h4
		var _12___mcc_h5 _dafny.Map = _source0.Get_().(D1_DC5).Cf10
		_ = _12___mcc_h5
		var _13_cf10 _dafny.Map = _12___mcc_h5
		_ = _13_cf10
		var _14_cf9 bool = _11___mcc_h4
		_ = _14_cf9
		var _15_cf8 _dafny.Int = _10___mcc_h3
		_ = _15_cf8
		var _16_cf7 _dafny.Int = _9___mcc_h2
		_ = _16_cf7
		if !(_14_cf9) {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_17_cf7 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_18_i0 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_17_cf7)
				}
			})(_16_cf7)))
		} else {
			return _dafny.SeqOf(_15_cf8)
		}
	} else {
		var _19___mcc_h6 _dafny.Sequence = _source0.Get_().(D1_DC3).Cf4
		_ = _19___mcc_h6
		var _20_cf4 _dafny.Sequence = _19___mcc_h6
		_ = _20_cf4
		return _dafny.SeqOf(_dafny.IntOfInt64(767))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if false {
			return _dafny.SetOf(_dafny.IntOfInt64(492))
		}
		return _dafny.SetOf(_dafny.IntOfInt64(135), (func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(688), _dafny.IntOfInt64(805))).Cardinality())), _dafny.IntOfInt64(52))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _21_v0 _dafny.Int
				_21_v0 = interface{}(_compr_0).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(688), _dafny.IntOfInt64(805))).Cardinality())), _dafny.IntOfInt64(52)), _21_v0) {
					_coll0.Add((_21_v0).Times(_dafny.IntOfInt64(82)), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
				}
			}
			return _coll0.ToMap()
		}()).Cardinality())
	})()).Intersection((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lql")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	}))).Cardinality()), _dafny.IntOfInt64(765), _dafny.IntOfInt64(79))).Intersection((Companion_D10_.Create_DC28_(_dafny.SetOf(_dafny.IntOfInt64(701)))).Dtor_cf44()))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(583), _dafny.IntOfInt64(764))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(583)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(764)) < 0) {
				_coll1.Add((_23_v0).Times((_dafny.SetOf(_dafny.IntOfInt64(37))).Cardinality()), _dafny.IntOfInt64(350))
			}
		}
		return _coll1.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-272), _dafny.IntOfInt64(931))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _24_v1 _dafny.Int
			_24_v1 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-272)).Cmp(_24_v1) <= 0) && ((_24_v1).Cmp(_dafny.IntOfInt64(931)) < 0) {
				_coll2.Add((_24_v1).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_25_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality())), _dafny.IntOfInt64(569))
			}
		}
		return _coll2.ToMap()
	}())).Merge((func() _dafny.Map {
		if true {
			return func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(900), _dafny.IntOfInt64(332))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _26_v2 _dafny.Int
					_26_v2 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(900)).Cmp(_26_v2) <= 0) && ((_26_v2).Cmp(_dafny.IntOfInt64(332)) < 0) {
						_coll3.Add(Companion_Default___.SafeDivisionInt(_26_v2, _dafny.IntOfInt64(574)), _dafny.IntOfInt64(492))
					}
				}
				return _coll3.ToMap()
			}()
		}
		return func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(410), false)).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _27_v3 _dafny.Int
				_27_v3 = interface{}(_compr_4).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(410), false)).Contains(_27_v3) {
					_coll4.Add(Companion_Default___.SafeDivisionInt(_27_v3, _dafny.IntOfInt64(610)), _dafny.IntOfInt64(745))
				}
			}
			return _coll4.ToMap()
		}()
	})())
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Map, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SetOf(false))).Intersection(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(!(!(true))), _dafny.SetOf(true, !(false)))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _28_v0 _dafny.Set
			_28_v0 = interface{}(_compr_5).(_dafny.Set)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(!(!(true))), _dafny.SetOf(true, !(false))), _28_v0) {
				_coll5.Add(_28_v0)
			}
		}
		return _coll5.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(804))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_29_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))).Cardinality())).Cmp(_dafny.IntOfInt64(-960)) == 0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xvbhie")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(724), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-52))).Cardinality()))).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _30_v0 _dafny.Int
			_30_v0 = interface{}(_compr_6).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(724), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-52))).Cardinality()))).Contains(_30_v0) {
				_coll6.Add((_30_v0).Times(_dafny.IntOfInt64(628)), false)
			}
		}
		return _coll6.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-925), false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(47), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(623), true)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(33))).Union(_dafny.SetOf(_dafny.IntOfInt64(-138))))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(true))), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), false, (Companion_D1_.Create_DC5_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kmamusrti")).Cardinality())), _dafny.IntOfInt64(983), false, func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(871), _dafny.IntOfInt64(44))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _31_v0 _dafny.Int
			_31_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(871)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(44)) < 0) {
				_coll7.Add((_31_v0).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))).Cardinality()), _dafny.SeqOf(_dafny.CodePoint('q')))
			}
		}
		return _coll7.ToMap()
	}())).Dtor_cf9(), (Companion_D0_.Create_DC0_(!(true))).Dtor_cf0()), _dafny.SeqOf(false, false, false)))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((func() _dafny.Int {
		if !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(false), _dafny.SeqOf(true, false)) {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-775))).Uint32(), func(coer9 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_32_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(false)
			})), _dafny.SeqOf(_dafny.MultiSetOf(true)))).Cardinality())
		}
		return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ltc")).Cardinality())), _dafny.IntOfInt64(236))
	})())
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('d')
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Map, globalState *GlobalState) bool {
	var _source1 D1 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("yohqnt"))
	_ = _source1
	if _source1.Is_DC4() {
		var _33___mcc_h0 _dafny.Int = _source1.Get_().(D1_DC4).Cf5
		_ = _33___mcc_h0
		var _34___mcc_h1 _dafny.Int = _source1.Get_().(D1_DC4).Cf6
		_ = _34___mcc_h1
		var _35_cf6 _dafny.Int = _34___mcc_h1
		_ = _35_cf6
		var _36_cf5 _dafny.Int = _33___mcc_h0
		_ = _36_cf5
		return false
	} else if _source1.Is_DC5() {
		var _37___mcc_h2 _dafny.Int = _source1.Get_().(D1_DC5).Cf7
		_ = _37___mcc_h2
		var _38___mcc_h3 _dafny.Int = _source1.Get_().(D1_DC5).Cf8
		_ = _38___mcc_h3
		var _39___mcc_h4 bool = _source1.Get_().(D1_DC5).Cf9
		_ = _39___mcc_h4
		var _40___mcc_h5 _dafny.Map = _source1.Get_().(D1_DC5).Cf10
		_ = _40___mcc_h5
		var _41_cf10 _dafny.Map = _40___mcc_h5
		_ = _41_cf10
		var _42_cf9 bool = _39___mcc_h4
		_ = _42_cf9
		var _43_cf8 _dafny.Int = _38___mcc_h3
		_ = _43_cf8
		var _44_cf7 _dafny.Int = _37___mcc_h2
		_ = _44_cf7
		return ((_dafny.MultiSetOf(_43_cf8, _43_cf8, _43_cf8, _44_cf7)).Cardinality()).Cmp((_dafny.MultiSetOf(_dafny.IntOfInt64(237), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_44_cf7, (_dafny.Zero).Minus(_43_cf8))).Cardinality())).Cardinality()) == 0
	} else {
		var _45___mcc_h6 _dafny.Sequence = _source1.Get_().(D1_DC3).Cf4
		_ = _45___mcc_h6
		var _46_cf4 _dafny.Sequence = _45___mcc_h6
		_ = _46_cf4
		return (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_46_cf4).Cardinality()), _dafny.IntOfInt64(390), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality(), _dafny.IntOfUint32((_46_cf4).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())).Cmp(_dafny.IntOfInt64(-744)) <= 0
	}
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(true) || (true))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) D3 {
	var _source2 D5 = Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(738)))
	_ = _source2
	if _source2.Is_DC14() {
		var _47___mcc_h0 _dafny.Int = _source2.Get_().(D5_DC14).Cf25
		_ = _47___mcc_h0
		var _48___mcc_h1 _dafny.Array = _source2.Get_().(D5_DC14).Cf26
		_ = _48___mcc_h1
		var _49___mcc_h2 bool = _source2.Get_().(D5_DC14).Cf27
		_ = _49___mcc_h2
		var _50_cf27 bool = _49___mcc_h2
		_ = _50_cf27
		var _51_cf26 _dafny.Array = _48___mcc_h1
		_ = _51_cf26
		var _52_cf25 _dafny.Int = _47___mcc_h0
		_ = _52_cf25
		return Companion_D3_.Create_DC9_(_52_cf25)
	} else if _source2.Is_DC13() {
		var _53___mcc_h3 _dafny.Map = _source2.Get_().(D5_DC13).Cf24
		_ = _53___mcc_h3
		var _54_cf24 _dafny.Map = _53___mcc_h3
		_ = _54_cf24
		return Companion_D3_.Create_DC9_((_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality())))).Cardinality()))
	} else {
		var _55___mcc_h4 D5 = _source2.Get_().(D5_DC15).Cf28
		_ = _55___mcc_h4
		var _56_cf28 D5 = _55___mcc_h4
		_ = _56_cf28
		return Companion_D3_.Create_DC9_(_dafny.IntOfInt64(218))
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Map, globalState *GlobalState) {
	var _57_v0 _dafny.Sequence
	_ = _57_v0
	_57_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ir")
	var _58_v1 _dafny.Int
	_ = _58_v1
	_58_v1 = _dafny.IntOfInt64(671)
	_57_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("drtgu"), Companion_Default___.Fm9((_dafny.Zero).Minus(_58_v1), globalState)), _dafny.Companion_Sequence_.Concatenate(_57_v0, _57_v0))
	var _59_v2 bool
	_ = _59_v2
	_59_v2 = true
	var _60_v3 _dafny.MultiSet
	_ = _60_v3
	_60_v3 = _dafny.MultiSetOf(!(_59_v2))
	var _61_v4 _dafny.Map
	_ = _61_v4
	_61_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v2, _60_v3)
	var _62_v5 _dafny.Sequence
	_ = _62_v5
	_62_v5 = _dafny.SeqOf(((func() _dafny.MultiSet {
		if (_61_v4).Contains(_59_v2) {
			return (_61_v4).Get(_59_v2).(_dafny.MultiSet)
		}
		return _60_v3
	})()).Update(false, Companion_Default___.Abs(_58_v1)))
	var _63_v6 _dafny.MultiSet
	_ = _63_v6
	_63_v6 = _dafny.MultiSetOf(_dafny.IntOfUint32((_57_v0).Cardinality()))
	var _64_v7 _dafny.Map
	_ = _64_v7
	_64_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v6, false)
	var _65_v8 _dafny.Map
	_ = _65_v8
	_65_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v1, (_64_v7).Cardinality())
	var _66_v9 _dafny.Sequence
	_ = _66_v9
	_66_v9 = _dafny.SeqOf(_59_v2)
	var _67_v10 D2
	_ = _67_v10
	_67_v10 = Companion_D2_.Create_DC7_(_59_v2, false, _58_v1)
	var _68_v11 _dafny.Array
	_ = _68_v11
	var _nwElement0_0 _dafny.MultiSet = Companion_Default___.Fm22(globalState)
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(20))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_60_v3, 1)
	(_nw0).ArraySet1((Companion_Default___.Fm22(globalState)).Update(_59_v2, Companion_Default___.Abs(_58_v1)), 2)
	(_nw0).ArraySet1((_62_v5).Select((Companion_Default___.SafeIndex(_58_v1, _dafny.IntOfUint32((_62_v5).Cardinality()))).Uint32()).(_dafny.MultiSet), 3)
	(_nw0).ArraySet1(_dafny.MultiSetOf(_59_v2, _59_v2, !(Companion_Default___.Fm21(_58_v1, _59_v2, _59_v2, _65_v8, globalState))), 4)
	(_nw0).ArraySet1(_60_v3, 5)
	(_nw0).ArraySet1(_60_v3, 6)
	(_nw0).ArraySet1(_60_v3, 7)
	(_nw0).ArraySet1(_60_v3, 8)
	(_nw0).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_59_v2, _59_v2)), 9)
	(_nw0).ArraySet1(_60_v3, 10)
	(_nw0).ArraySet1(_60_v3, 11)
	(_nw0).ArraySet1((_60_v3).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_66_v9, (Companion_Default___.SafeIndex(_58_v1, _dafny.IntOfUint32((_66_v9).Cardinality()))).Uint32(), (_67_v10).Dtor_cf12()), (Companion_Default___.SafeIndex(_58_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_66_v9, (Companion_Default___.SafeIndex(_58_v1, _dafny.IntOfUint32((_66_v9).Cardinality()))).Uint32(), (_67_v10).Dtor_cf12())).Cardinality()))).Uint32(), _59_v2))), 12)
	(_nw0).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(_59_v2))), 13)
	(_nw0).ArraySet1((_60_v3).Difference(Companion_Default___.Fm22(globalState)), 14)
	(_nw0).ArraySet1((_dafny.MultiSetOf(_59_v2, _59_v2, _59_v2)).Union((_60_v3).Update(_59_v2, Companion_Default___.Abs(_dafny.IntOfUint32((_57_v0).Cardinality())))), 15)
	(_nw0).ArraySet1((_60_v3).Update(_59_v2, Companion_Default___.Abs(_58_v1)), 16)
	(_nw0).ArraySet1(_60_v3, 17)
	(_nw0).ArraySet1(_60_v3, 18)
	(_nw0).ArraySet1(_dafny.MultiSetFromSeq(_66_v9), 19)
	_68_v11 = _nw0
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_68_v11), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _69_i0 _dafny.Int
		_69_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_69_i0).Sign() != -1) && ((_69_i0).Cmp(_dafny.ArrayLen((_68_v11), 0)) < 0)) {
			(_68_v11).ArraySet1((func() _dafny.MultiSet {
				if _59_v2 {
					return _60_v3
				}
				return (_60_v3).Update(_59_v2, Companion_Default___.Abs(_58_v1))
			})(), (_69_i0).Int())
		}
	}
	var _70_v12 T1
	_ = _70_v12
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__(_59_v2, _59_v2, _dafny.UnicodeSeqOfUtf8Bytes("jtxlh"), _58_v1, _59_v2)
	_70_v12 = _nw1
	var _71_v13 _dafny.Map
	_ = _71_v13
	_71_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v12, _58_v1)
	(globalState).F11 = (((func() _dafny.Int {
		if (_71_v13).Contains(_70_v12) {
			return (_71_v13).Get(_70_v12).(_dafny.Int)
		}
		return _70_v12.F14()
	})()).Times(_70_v12.F14())).Plus(_70_v12.F14())
	_59_v2 = _59_v2
	var _source3 D0 = Companion_D0_.Create_DC2_(_70_v12.F14(), (_dafny.MultiSetOf(true, _59_v2, (_70_v12).F15(), _59_v2)).Difference(_dafny.MultiSetFromSeq(_66_v9)))
	_ = _source3
	if _source3.Is_DC1() {
		var _72___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
		_ = _72___mcc_h0
		var _73_cf1 _dafny.Int = _72___mcc_h0
		_ = _73_cf1
		var _74_v14 _dafny.Map
		_ = _74_v14
		_74_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_70_v12).F15())
		var _75_v15 *C0
		_ = _75_v15
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__(_59_v2, ((func() bool {
			if (_74_v14).Contains(_59_v2) {
				return (_74_v14).Get(_59_v2).(bool)
			}
			return (_70_v12).F15()
		})()) == ((_70_v12).F15()), _70_v12.F16(), _58_v1, (_70_v12.F14()).Cmp(_70_v12.F14()) >= 0)
		_75_v15 = _nw2
		_75_v15 = _75_v15
		_58_v1 = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_dafny.MultiSetOf(_58_v1)).IsSubsetOf(_63_v6) {
				return (_dafny.Zero).Minus(_73_cf1)
			}
			return _70_v12.F14()
		})())
		var _76_v16 _dafny.Map
		_ = _76_v16
		_76_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_70_v12).F15(), _70_v12)
		var _77_v17 _dafny.Array
		_ = _77_v17
		var _nwElement0_1 *C0 = _75_v15
		_ = _nwElement0_1
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_1, 0)
		_77_v17 = _nw3
		var _78_v18 _dafny.MultiSet
		_ = _78_v18
		_78_v18 = _dafny.MultiSetOf(_77_v17, _77_v17, _77_v17, _77_v17)
		(globalState).F11 = (func() _dafny.Int {
			if _59_v2 {
				return Companion_Default___.SafeModuloInt(_70_v12.F14(), (_76_v16).Cardinality())
			}
			return (_dafny.Zero).Minus(((_78_v18).Update(_77_v17, Companion_Default___.Abs(_70_v12.F14()))).Cardinality())
		})()
		_73_cf1 = (func() _dafny.Int {
			if (_60_v3).Contains((_75_v15).F17()) {
				return (_60_v3).Multiplicity((_75_v15).F17())
			}
			return _73_cf1
		})()
	} else if _source3.Is_DC2() {
		var _79___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC2).Cf2
		_ = _79___mcc_h1
		var _80___mcc_h2 _dafny.MultiSet = _source3.Get_().(D0_DC2).Cf3
		_ = _80___mcc_h2
		var _81_cf3 _dafny.MultiSet = _80___mcc_h2
		_ = _81_cf3
		var _82_cf2 _dafny.Int = _79___mcc_h1
		_ = _82_cf2
		var _83_v19 _dafny.Array
		_ = _83_v19
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw4
		_83_v19 = _nw4
		var _84_v20 _dafny.Map
		_ = _84_v20
		_84_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_70_v12).F15(), _83_v19)
		var _85_v21 _dafny.Map
		_ = _85_v21
		_85_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_70_v12).Fm0((_70_v12).F15(), globalState), _59_v2)
		if Companion_Default___.Fm21((_84_v20).Cardinality(), (func() bool {
			if (_85_v21).Contains((_70_v12).F15()) {
				return (_85_v21).Get((_70_v12).F15()).(bool)
			}
			return (_70_v12).F15()
		})(), (_70_v12).F15(), _65_v8, globalState) {
			(globalState).F11 = (_70_v12.F14()).Minus(_58_v1)
			var _86_v22 D3
			_ = _86_v22
			_86_v22 = Companion_D3_.Create_DC9_(_82_cf2)
			var _87_v23 *C1
			_ = _87_v23
			var _nw5 *C1 = New_C1_()
			_ = _nw5
			_nw5.Ctor__(_70_v12.F16(), _66_v9, (_60_v3).Cardinality(), (_70_v12).F15())
			_87_v23 = _nw5
			var _88_v24 _dafny.Map
			_ = _88_v24
			_88_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_70_v12).F15(), _87_v23)
			_86_v22 = Companion_Default___.Fm23(_dafny.SeqOf(_70_v12.F14(), (_88_v24).Cardinality(), _70_v12.F14()), !(_59_v2), globalState)
			var _89_v25 _dafny.Array
			_ = _89_v25
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_0
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_90_v12 T1, _91_cf3 _dafny.MultiSet, _92_cf2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_93_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((func() _dafny.Int {
							if (_91_cf3).Contains((_90_v12).F15()) {
								return (_91_cf3).Multiplicity((_90_v12).F15())
							}
							return _90_v12.F14()
						})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(834), _92_cf2)).Cardinality(), _90_v12.F14())
					}
				})(_70_v12, _81_cf3, _82_cf2)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw6 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw6).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw6).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_89_v25 = _nw6
			var _rhs0 _dafny.Int = _70_v12.F14()
			_ = _rhs0
			var _rhs1 bool = !(_63_v6).Contains((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yitpb")).Cardinality())).Minus(_58_v1))
			_ = _rhs1
			var _rhs2 bool = !((_70_v12.F14()).Cmp((_dafny.IntOfInt64(949)).Plus(_dafny.IntOfInt64(-463))) < 0)
			_ = _rhs2
			var _rhs3 _dafny.Sequence = _57_v0
			_ = _rhs3
			var _rhs4 _dafny.Array = (func() _dafny.Array {
				if _59_v2 {
					return _89_v25
				}
				return _89_v25
			})()
			_ = _rhs4
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			_lhs0.F11 = _rhs0
			_lhs1.F0 = _rhs1
			_59_v2 = _rhs2
			_57_v0 = _rhs3
			_89_v25 = _rhs4
			var _94_v26 _dafny.Map
			_ = _94_v26
			_94_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(562), _59_v2)
			var _95_v27 D8
			_ = _95_v27
			_95_v27 = Companion_D8_.Create_DC21_(_94_v26)
			_94_v26 = ((_95_v27).Dtor_cf36()).Merge(_94_v26)
			var _96_v28 _dafny.Array
			_ = _96_v28
			var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw7
			_96_v28 = _nw7
			_96_v28 = _96_v28
		} else {
			(_70_v12).F16_set_(_dafny.Companion_Sequence_.Concatenate(_70_v12.F16(), _70_v12.F16()))
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_83_v19), 0))
			_ = _index0
			(_83_v19).ArraySet1(Companion_Default___.Fm21(_70_v12.F14(), false, (_70_v12).F15(), _65_v8, globalState), (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_83_v19), 0))
			_ = _index1
			(_83_v19).ArraySet1((_81_cf3).IsDisjointFrom(_60_v3), (_index1).Int())
			var _97_v29 _dafny.Array
			_ = _97_v29
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_1
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = (func(_98_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_99_i2 _dafny.Int) _dafny.Sequence {
						return _98_v9
					}
				})(_66_v9)
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
			_97_v29 = _nw8
			var _100_v30 D5
			_ = _100_v30
			_100_v30 = Companion_D5_.Create_DC14_(_70_v12.F14(), _97_v29, true)
			var _101_v31 _dafny.Array
			_ = _101_v31
			var _nwElement0_2 bool = (_70_v12).F15()
			_ = _nwElement0_2
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(17))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_2, 0)
			(_nw9).ArraySet1(false, 1)
			(_nw9).ArraySet1((_70_v12).F15(), 2)
			(_nw9).ArraySet1((_70_v12).F15(), 3)
			(_nw9).ArraySet1((_70_v12).F15(), 4)
			(_nw9).ArraySet1((_70_v12).F15(), 5)
			(_nw9).ArraySet1((_70_v12).F15(), 6)
			(_nw9).ArraySet1((_70_v12).F15(), 7)
			(_nw9).ArraySet1((_70_v12).F15(), 8)
			(_nw9).ArraySet1((_70_v12).F15(), 9)
			(_nw9).ArraySet1((_70_v12).F15(), 10)
			(_nw9).ArraySet1(!((_100_v30).Dtor_cf27()), 11)
			(_nw9).ArraySet1(!(!((_70_v12).F15())), 12)
			(_nw9).ArraySet1((_70_v12).F15(), 13)
			(_nw9).ArraySet1((_70_v12).F15(), 14)
			(_nw9).ArraySet1(!(!(_59_v2)), 15)
			(_nw9).ArraySet1((_70_v12).F15(), 16)
			_101_v31 = _nw9
			var _102_v32 _dafny.Map
			_ = _102_v32
			_102_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v31, _dafny.IntOfInt64(420))
			_102_v32 = (_102_v32).Update(_101_v31, (_63_v6).Cardinality())
			var _103_v33 _dafny.Set
			_ = _103_v33
			_103_v33 = _dafny.SetOf((_70_v12).F15(), _59_v2)
			var _104_v34 _dafny.CodePoint
			_ = _104_v34
			_104_v34 = _dafny.CodePoint('y')
			var _105_v35 _dafny.MultiSet
			_ = _105_v35
			_105_v35 = _dafny.MultiSetOf(_104_v34, _104_v34)
			var _pat_let_tv0 = _70_v12
			_ = _pat_let_tv0
			var _pat_let_tv1 = _83_v19
			_ = _pat_let_tv1
			var _pat_let_tv2 = _83_v19
			_ = _pat_let_tv2
			var _pat_let_tv3 = _83_v19
			_ = _pat_let_tv3
			var _pat_let_tv4 = _83_v19
			_ = _pat_let_tv4
			var _pat_let_tv5 = globalState
			_ = _pat_let_tv5
			var _pat_let_tv6 = globalState
			_ = _pat_let_tv6
			var _pat_let_tv7 = _70_v12
			_ = _pat_let_tv7
			var _106_v36 *C2
			_ = _106_v36
			var _nw10 *C2 = New_C2_()
			_ = _nw10
			_nw10.Ctor__(func(_pat_let0_0 D2) D2 {
				return func(_107_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let1_0 bool) D2 {
						return func(_108_dt__update_hcf12_h0 bool) D2 {
							return func(_pat_let2_0 bool) D2 {
								return func(_109_dt__update_hcf13_h0 bool) D2 {
									return Companion_D2_.Create_DC7_(_108_dt__update_hcf12_h0, _109_dt__update_hcf13_h0, (_107_dt__update__tmp_h0).Dtor_cf14())
								}(_pat_let2_0)
							}((_pat_let_tv7).F15())
						}(_pat_let1_0)
					}(Companion_Default___.Fm21(_pat_let_tv0.F14(), (_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(bool), (_pat_let_tv4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_pat_let_tv3), 0))).Int()).(bool), Companion_Default___.Fm13(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ffbfp")).Cardinality()), _pat_let_tv5), _pat_let_tv6))
				}(_pat_let0_0)
			}(Companion_D2_.Create_DC7_(Companion_Default___.Fm21((_103_v33).Cardinality(), (_70_v12).F15(), (_70_v12).F15(), _65_v8, globalState), (_70_v12).F15(), _70_v12.F14())), (func() _dafny.Int {
				if (_105_v35).Contains(_104_v34) {
					return (_105_v35).Multiplicity(_104_v34)
				}
				return _82_cf2
			})(), (_70_v12).F15())
			_106_v36 = _nw10
			var _110_v37 *C0
			_ = _110_v37
			var _nw11 *C0 = New_C0_()
			_ = _nw11
			_nw11.Ctor__(!((_70_v12).F15()), (_70_v12).F15(), _57_v0, _82_cf2, (_70_v12).F15())
			_110_v37 = _nw11
			var _111_v38 _dafny.Map
			_ = _111_v38
			_111_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_66_v9).Select((Companion_Default___.SafeIndex(_82_cf2, _dafny.IntOfUint32((_66_v9).Cardinality()))).Uint32()).(bool), _110_v37)
			(globalState).F11 = (((_111_v38).Merge(_111_v38)).Merge(_111_v38)).Cardinality()
		}
		var _112_v39 *C1
		_ = _112_v39
		var _nw12 *C1 = New_C1_()
		_ = _nw12
		_nw12.Ctor__(_57_v0, _dafny.SeqOf(_59_v2), _58_v1, (_70_v12).F15())
		_112_v39 = _nw12
		var _113_v40 _dafny.Map
		_ = _113_v40
		_113_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v39, (_70_v12).F15())
		var _114_v41 _dafny.Map
		_ = _114_v41
		_114_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_70_v12).F15(), (_113_v40).Cardinality())
		_114_v41 = (_114_v41).Update((_70_v12.F14()).Cmp(_70_v12.F14()) <= 0, _70_v12.F14())
		var _115_v42 _dafny.Set
		_ = _115_v42
		_115_v42 = _dafny.SetOf(_112_v39, _112_v39, _112_v39)
		var _116_v43 D1
		_ = _116_v43
		_116_v43 = Companion_D1_.Create_DC4_(Companion_Default___.SafeDivisionInt((_115_v42).Cardinality(), _dafny.IntOfUint32((_70_v12.F16()).Cardinality())), _70_v12.F14())
		_116_v43 = _116_v43
		var _117_v44 _dafny.Array
		_ = _117_v44
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
		_ = _nw13
		_117_v44 = _nw13
		var _118_v45 _dafny.CodePoint
		_ = _118_v45
		_118_v45 = _dafny.CodePoint('s')
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_117_v44), 0))
		_ = _index2
		(_117_v44).ArraySet1CodePoint(_118_v45, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_117_v44), 0))
		_ = _index3
		(_117_v44).ArraySet1CodePoint(_118_v45, (_index3).Int())
	} else {
		var _119___mcc_h3 bool = _source3.Get_().(D0_DC0).Cf0
		_ = _119___mcc_h3
		var _120_cf0 bool = _119___mcc_h3
		_ = _120_cf0
		_66_v9 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, (_70_v12).F15(), (_70_v12).Fm0((_70_v12).F15(), globalState)), _66_v9)
		var _121_v46 *C0
		_ = _121_v46
		var _nw14 *C0 = New_C0_()
		_ = _nw14
		_nw14.Ctor__(_59_v2, !((_70_v12).F15()), _57_v0, _70_v12.F14(), (_70_v12).F15())
		_121_v46 = _nw14
		var _122_v47 _dafny.Map
		_ = _122_v47
		_122_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_63_v6, _121_v46)
		var _123_v48 _dafny.Map
		_ = _123_v48
		_123_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v2, _70_v12.F14())
		var _124_v49 _dafny.Map
		_ = _124_v49
		_124_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v1, _123_v48)
		var _125_v50 _dafny.Array
		_ = _125_v50
		var _nwElement0_3 *C0 = _121_v46
		_ = _nwElement0_3
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(12))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_3, 0)
		(_nw15).ArraySet1(_121_v46, 1)
		(_nw15).ArraySet1(_121_v46, 2)
		(_nw15).ArraySet1(_121_v46, 3)
		(_nw15).ArraySet1(_121_v46, 4)
		(_nw15).ArraySet1(_121_v46, 5)
		(_nw15).ArraySet1(_121_v46, 6)
		(_nw15).ArraySet1(_121_v46, 7)
		(_nw15).ArraySet1((func() *C0 {
			if (_70_v12).F15() {
				return _121_v46
			}
			return _121_v46
		})(), 8)
		(_nw15).ArraySet1(_121_v46, 9)
		(_nw15).ArraySet1(_121_v46, 10)
		(_nw15).ArraySet1((func() *C0 {
			if (_122_v47).Contains(_dafny.MultiSetOf(((func() _dafny.Map {
				if (_124_v49).Contains(_dafny.IntOfInt64(512)) {
					return (_124_v49).Get(_dafny.IntOfInt64(512)).(_dafny.Map)
				}
				return _123_v48
			})()).Cardinality())) {
				return (_122_v47).Get(_dafny.MultiSetOf(((func() _dafny.Map {
					if (_124_v49).Contains(_dafny.IntOfInt64(512)) {
						return (_124_v49).Get(_dafny.IntOfInt64(512)).(_dafny.Map)
					}
					return _123_v48
				})()).Cardinality())).(*C0)
			}
			return _121_v46
		})(), 11)
		_125_v50 = _nw15
		var _126_v51 D9
		_ = _126_v51
		_126_v51 = Companion_D9_.Create_DC25_(_125_v50)
		_125_v50 = (_126_v51).Dtor_cf41()
		var _127_v52 _dafny.Map
		_ = _127_v52
		_127_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(globalState), (_70_v12.F14()).Cmp(_70_v12.F14()) == 0)
		var _128_v53 _dafny.MultiSet
		_ = _128_v53
		_128_v53 = _dafny.MultiSetOf(_121_v46, _121_v46)
		var _129_v54 _dafny.Sequence
		_ = _129_v54
		_129_v54 = _dafny.SeqOf(_58_v1)
		_120_cf0 = (func() bool {
			if (_127_v52).Contains((func() _dafny.Int {
				if _120_cf0 {
					return _58_v1
				}
				return (func() _dafny.Int {
					if (_128_v53).Contains(_121_v46) {
						return (_128_v53).Multiplicity(_121_v46)
					}
					return _dafny.IntOfUint32((_129_v54).Cardinality())
				})()
			})()) {
				return (_127_v52).Get((func() _dafny.Int {
					if _120_cf0 {
						return _58_v1
					}
					return (func() _dafny.Int {
						if (_128_v53).Contains(_121_v46) {
							return (_128_v53).Multiplicity(_121_v46)
						}
						return _dafny.IntOfUint32((_129_v54).Cardinality())
					})()
				})()).(bool)
			}
			return _120_cf0
		})()
		var _130_v55 _dafny.Array
		_ = _130_v55
		var _len0_2 _dafny.Int = _dafny.One
		_ = _len0_2
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_131_v46 *C0) func(_dafny.Int) bool {
				return func(_132_i3 _dafny.Int) bool {
					return (_131_v46).F18()
				}
			})(_121_v46)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw16 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw16).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw16).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_130_v55 = _nw16
		_130_v55 = _130_v55
	}
	var _133_v56 _dafny.Array
	_ = _133_v56
	var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
	_ = _nw17
	_133_v56 = _nw17
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_133_v56), 0))
	_ = _index4
	(_133_v56).ArraySet1((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!((_70_v12).F15()), _59_v2), _dafny.SeqOf(_59_v2, (_66_v9).Select((Companion_Default___.SafeIndex(_58_v1, _dafny.IntOfUint32((_66_v9).Cardinality()))).Uint32()).(bool), (_70_v12).F15(), _59_v2)))).Cardinality(), (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_133_v56), 0))
	_ = _index5
	(_133_v56).ArraySet1(((_dafny.Zero).Minus((_70_v12.F14()).Plus(_dafny.IntOfUint32((_66_v9).Cardinality())))).Plus((_dafny.Zero).Minus(_58_v1)), (_index5).Int())
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _134_v0 bool
	_ = _134_v0
	_134_v0 = false
	var _135_v1 _dafny.Sequence
	_ = _135_v1
	_135_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ogvv")
	var _136_v2 _dafny.Sequence
	_ = _136_v2
	_136_v2 = _dafny.SeqOf(_134_v0)
	var _137_v3 _dafny.Int
	_ = _137_v3
	_137_v3 = _dafny.IntOfInt64(329)
	var _138_v4 _dafny.Sequence
	_ = _138_v4
	_138_v4 = _dafny.SeqOf(_134_v0, _134_v0, _134_v0, (_136_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_137_v3), _dafny.IntOfUint32((_136_v2).Cardinality()))).Uint32()).(bool), !(_134_v0))
	var _139_globalState *GlobalState
	_ = _139_globalState
	var _nw18 *GlobalState = New_GlobalState_()
	_ = _nw18
	_nw18.Ctor__(false, true, _dafny.IntOfInt64(140), true, (func() _dafny.Sequence {
		if _134_v0 {
			return _135_v1
		}
		return _135_v1
	})(), false, _dafny.IntOfInt64(188), true, _dafny.IntOfInt64(38), _dafny.IntOfInt64(937), _dafny.IntOfInt64(439), _dafny.IntOfInt64(8), false, _136_v2)
	_139_globalState = _nw18
	if _134_v0 {
		var _140_v5 _dafny.Set
		_ = _140_v5
		_140_v5 = _dafny.SetOf(_134_v0, _134_v0)
		var _141_v6 _dafny.Sequence
		_ = _141_v6
		_141_v6 = _dafny.SeqOf(_140_v5)
		(_139_globalState).F0 = !_dafny.Companion_Sequence_.Equal(_141_v6, _dafny.SeqOf(_140_v5, _140_v5, _140_v5, _140_v5))
		var _142_v7 _dafny.Map
		_ = _142_v7
		_142_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, !(_134_v0))
		var _143_v8 _dafny.Array
		_ = _143_v8
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw19
		_143_v8 = _nw19
		Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v7, _143_v8), _139_globalState)
		var _144_v9 _dafny.Map
		_ = _144_v9
		_144_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, _135_v1)
		_144_v9 = (_144_v9).Update(_137_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hkrlymcrd"), _135_v1))
		var _145_v10 _dafny.Array
		_ = _145_v10
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw20
		_145_v10 = _nw20
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_145_v10), 0))
		_ = _index6
		(_145_v10).ArraySet1(_137_v3, (_index6).Int())
		var _146_v11 _dafny.MultiSet
		_ = _146_v11
		_146_v11 = _dafny.MultiSetOf(_137_v3, _dafny.IntOfInt64(378))
		var _147_v12 _dafny.Map
		_ = _147_v12
		_147_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, ((_146_v11).Intersection((_146_v11).Update(_dafny.IntOfUint32((_dafny.SeqOf(_137_v3)).Cardinality()), Companion_Default___.Abs(_137_v3)))).Cardinality())
		var _148_v13 _dafny.Map
		_ = _148_v13
		_148_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _147_v12)
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_145_v10), 0))
		_ = _index7
		var _rhs5 bool = true
		_ = _rhs5
		var _rhs6 _dafny.Set = _140_v5
		_ = _rhs6
		var _rhs7 _dafny.Int = _137_v3
		_ = _rhs7
		var _rhs8 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _137_v3)).Merge(((func() _dafny.Map {
			if (_148_v13).Contains(!(_134_v0)) {
				return (_148_v13).Get(!(_134_v0)).(_dafny.Map)
			}
			return _147_v12
		})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _dafny.IntOfUint32((_135_v1).Cardinality()))))
		_ = _rhs8
		var _lhs2 _dafny.Array = _145_v10
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_145_v10), 0))
		_ = _lhs3
		_134_v0 = _rhs5
		_140_v5 = _rhs6
		(_lhs2).ArraySet1(_rhs7, (_lhs3).Int())
		_147_v12 = _rhs8
		var _149_v14 _dafny.Map
		_ = _149_v14
		_149_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_145_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_145_v10), 0))).Int()).(_dafny.Int), _134_v0), _143_v8)
		Companion_Default___.M0(_149_v14, _139_globalState)
	} else {
		var _150_v15 _dafny.Map
		_ = _150_v15
		_150_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_137_v3), _134_v0)
		var _151_v16 _dafny.Array
		_ = _151_v16
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_3
		var _nw21 _dafny.Array
		_ = _nw21
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw21 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_152_v0 bool) func(_dafny.Int) bool {
				return func(_153_i0 _dafny.Int) bool {
					return !(_152_v0)
				}
			})(_134_v0)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw21 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw21).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw21).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_151_v16 = _nw21
		var _154_v17 _dafny.Map
		_ = _154_v17
		_154_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v15, _151_v16)
		Companion_Default___.M0(_154_v17, _139_globalState)
		(_139_globalState).F11 = _137_v3
		if (func() bool {
			if (_150_v15).Contains(_137_v3) {
				return (_150_v15).Get(_137_v3).(bool)
			}
			return _134_v0
		})() {
			var _155_v18 _dafny.MultiSet
			_ = _155_v18
			_155_v18 = _dafny.MultiSetOf(_137_v3, _137_v3, (_dafny.Zero).Minus(_137_v3))
			_134_v0 = (func() bool {
				if (_150_v15).Contains((_155_v18).Cardinality()) {
					return (_150_v15).Get((_155_v18).Cardinality()).(bool)
				}
				return _134_v0
			})()
			Companion_Default___.M0((_154_v17).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, _134_v0), _151_v16)), _139_globalState)
			var _156_v19 *C0
			_ = _156_v19
			var _nw22 *C0 = New_C0_()
			_ = _nw22
			_nw22.Ctor__(_134_v0, !((_134_v0) || (_134_v0)), _135_v1, _dafny.IntOfInt64(385), true)
			_156_v19 = _nw22
			var _157_v21 _dafny.Map
			_ = _157_v21
			_157_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _137_v3)
			(_139_globalState).F11 = (func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter9 := _dafny.Iterate((_157_v21).Keys().Elements()); ; {
					_compr_8, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _158_v20 _dafny.CodePoint
					_158_v20 = interface{}(_compr_8).(_dafny.CodePoint)
					if (_157_v21).Contains(_158_v20) {
						_coll8.Add(_158_v20, (_156_v19).F17())
					}
				}
				return _coll8.ToMap()
			}()).Cardinality()
			var _159_v22 D0
			_ = _159_v22
			_159_v22 = Companion_D0_.Create_DC0_(true)
			(_139_globalState).F0 = (_159_v22).Dtor_cf0()
		} else {
			var _160_v23 _dafny.CodePoint
			_ = _160_v23
			_160_v23 = _dafny.CodePoint('r')
			var _161_v24 _dafny.Sequence
			_ = _161_v24
			_161_v24 = _dafny.SeqOf(_dafny.IntOfInt64(382))
			var _162_v25 _dafny.Map
			_ = _162_v25
			_162_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _160_v23)
			_160_v23 = (_135_v1).Select((Companion_Default___.SafeIndex((_161_v24).Select((Companion_Default___.SafeIndex((_162_v25).Cardinality(), _dafny.IntOfUint32((_161_v24).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_135_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			Companion_Default___.M0((_154_v17).Merge(_154_v17), _139_globalState)
			var _163_v26 _dafny.MultiSet
			_ = _163_v26
			_163_v26 = _dafny.MultiSetOf(_134_v0)
			var _rhs9 bool = !(_134_v0)
			_ = _rhs9
			var _rhs10 _dafny.Int = (_137_v3).Minus((func() _dafny.Int {
				if _134_v0 {
					return _dafny.IntOfUint32((_136_v2).Cardinality())
				}
				return _137_v3
			})())
			_ = _rhs10
			var _rhs11 _dafny.Int = _137_v3
			_ = _rhs11
			var _rhs12 bool = !(((_dafny.MultiSetOf(_134_v0)).Difference(_163_v26)).IsSubsetOf(_163_v26))
			_ = _rhs12
			var _lhs4 *GlobalState = _139_globalState
			_ = _lhs4
			var _lhs5 *GlobalState = _139_globalState
			_ = _lhs5
			var _lhs6 *GlobalState = _139_globalState
			_ = _lhs6
			_lhs4.F0 = _rhs9
			_137_v3 = _rhs10
			_lhs5.F11 = _rhs11
			_lhs6.F3 = _rhs12
			var _164_v27 _dafny.Map
			_ = _164_v27
			_164_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_137_v3).Plus((_dafny.Zero).Minus(_137_v3)), (_dafny.MultiSetOf(_134_v0, _134_v0)).Update(_134_v0, Companion_Default___.Abs(_137_v3)))
			_164_v27 = (Companion_D2_.Create_DC6_(_164_v27)).Dtor_cf11()
			Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v15, _151_v16), _139_globalState)
		}
		(_139_globalState).F3 = true
		(_139_globalState).F3 = !(_134_v0)
	}
	var _165_v28 _dafny.Array
	_ = _165_v28
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(7)
	_ = _len0_4
	var _nw23 _dafny.Array
	_ = _nw23
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw23 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) bool = (func(_166_v0 bool) func(_dafny.Int) bool {
			return func(_167_i2 _dafny.Int) bool {
				return (func() bool {
					if _166_v0 {
						return _166_v0
					}
					return _166_v0
				})()
			}
		})(_134_v0)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw23 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw23).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw23).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_165_v28 = _nw23
	for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_165_v28), 0))); ; {
		_guard_loop_1, _ok10 := _iter10()
		if !_ok10 {
			break
		}
		var _168_i1 _dafny.Int
		_168_i1 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_168_i1).Sign() != -1) && ((_168_i1).Cmp(_dafny.ArrayLen((_165_v28), 0)) < 0)) {
			(_165_v28).ArraySet1(_134_v0, (_168_i1).Int())
		}
	}
	var _169_v29 _dafny.Sequence
	_ = _169_v29
	_169_v29 = _dafny.SeqOf(_dafny.SeqOf(_134_v0), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_138_v4, _138_v4), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_137_v3), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_138_v4, _138_v4)).Cardinality()))).Uint32(), _134_v0), _136_v2)
	_169_v29 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(659))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}((func(_170_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
		return func(_171_i3 _dafny.Int) _dafny.Sequence {
			return _170_v4
		}
	})(_138_v4))), _dafny.SeqOf(_138_v4, _dafny.Companion_Sequence_.Update(_136_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tknb")).Cardinality()), _dafny.IntOfUint32((_136_v2).Cardinality()))).Uint32(), _134_v0), _138_v4)), _169_v29)
	_137_v3 = (Companion_Default___.SafeModuloInt(_137_v3, _137_v3)).Plus(_137_v3)
	if (_137_v3).Cmp(_137_v3) >= 0 {
		var _172_v30 _dafny.Array
		_ = _172_v30
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_5
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = (func(_173_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_174_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_174_i4, _dafny.IntOfUint32((_173_v4).Cardinality()))
				}
			})(_138_v4)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw24 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw24).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw24).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_172_v30 = _nw24
		var _175_v31 T0
		_ = _175_v31
		var _nw25 *C1 = New_C1_()
		_ = _nw25
		_nw25.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("qsgjrgt"), _138_v4, _137_v3, _134_v0)
		_175_v31 = _nw25
		var _176_v32 _dafny.Map
		_ = _176_v32
		_176_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v31, (_175_v31).F15())
		var _177_v33 T1
		_ = _177_v33
		var _nw26 *C0 = New_C0_()
		_ = _nw26
		_nw26.Ctor__(true, _134_v0, _135_v1, _137_v3, (_138_v4).Select((Companion_Default___.SafeIndex((_176_v32).Cardinality(), _dafny.IntOfUint32((_138_v4).Cardinality()))).Uint32()).(bool))
		_177_v33 = _nw26
		_137_v3 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v30, _177_v33)).Cardinality()
		var _178_v34 _dafny.MultiSet
		_ = _178_v34
		_178_v34 = _dafny.MultiSetOf((_175_v31.F14()).Times(_137_v3))
		_178_v34 = _dafny.MultiSetOf(_177_v33.F14(), (_175_v31.F14()).Times((_175_v31).Fm1(_178_v34, !((_175_v31).F15()), _139_globalState)), _175_v31.F14(), _dafny.IntOfInt64(297))
		var _179_v35 D1
		_ = _179_v35
		_179_v35 = Companion_D1_.Create_DC4_(_137_v3, (_dafny.Zero).Minus(_175_v31.F14()))
		var _source4 D1 = _179_v35
		_ = _source4
		if _source4.Is_DC4() {
			var _180___mcc_h0 _dafny.Int = _source4.Get_().(D1_DC4).Cf5
			_ = _180___mcc_h0
			var _181___mcc_h1 _dafny.Int = _source4.Get_().(D1_DC4).Cf6
			_ = _181___mcc_h1
			var _182_cf6 _dafny.Int = _181___mcc_h1
			_ = _182_cf6
			var _183_cf5 _dafny.Int = _180___mcc_h0
			_ = _183_cf5
			var _184_v36 D0
			_ = _184_v36
			_184_v36 = Companion_D0_.Create_DC0_(_134_v0)
			var _185_v37 D2
			_ = _185_v37
			_185_v37 = Companion_D2_.Create_DC7_(!((_177_v33).F15()), (_177_v33).F15(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_184_v36, _184_v36, _184_v36)).Cardinality()))
			var _186_v38 *C2
			_ = _186_v38
			var _nw27 *C2 = New_C2_()
			_ = _nw27
			_nw27.Ctor__(_185_v37, (_177_v33.F14()).Times(_175_v31.F14()), !(true))
			_186_v38 = _nw27
			var _187_v39 *C2
			_ = _187_v39
			var _nw28 *C2 = New_C2_()
			_ = _nw28
			_nw28.Ctor__(Companion_D2_.Create_DC7_(_134_v0, _134_v0, _175_v31.F14()), Companion_Default___.SafeDivisionInt(_177_v33.F14(), _182_cf6), (_186_v38).Fm0((_136_v2).Select((Companion_Default___.SafeIndex(_183_cf5, _dafny.IntOfUint32((_136_v2).Cardinality()))).Uint32()).(bool), _139_globalState))
			_187_v39 = _nw28
			var _188_v40 _dafny.Array
			_ = _188_v40
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(24))
			_ = _nw29
			_188_v40 = _nw29
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_188_v40), 0))
			_ = _index8
			(_188_v40).ArraySet1((_186_v38).F19(), (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_188_v40), 0))
			_ = _index9
			(_188_v40).ArraySet1((_187_v39).F19(), (_index9).Int())
			_138_v4 = _138_v4
		} else if _source4.Is_DC5() {
			var _189___mcc_h2 _dafny.Int = _source4.Get_().(D1_DC5).Cf7
			_ = _189___mcc_h2
			var _190___mcc_h3 _dafny.Int = _source4.Get_().(D1_DC5).Cf8
			_ = _190___mcc_h3
			var _191___mcc_h4 bool = _source4.Get_().(D1_DC5).Cf9
			_ = _191___mcc_h4
			var _192___mcc_h5 _dafny.Map = _source4.Get_().(D1_DC5).Cf10
			_ = _192___mcc_h5
			var _193_cf10 _dafny.Map = _192___mcc_h5
			_ = _193_cf10
			var _194_cf9 bool = _191___mcc_h4
			_ = _194_cf9
			var _195_cf8 _dafny.Int = _190___mcc_h3
			_ = _195_cf8
			var _196_cf7 _dafny.Int = _189___mcc_h2
			_ = _196_cf7
			var _197_v41 _dafny.Map
			_ = _197_v41
			_197_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, (_177_v33).F15())
			_197_v41 = (_197_v41).Update((_175_v31).F15(), (_177_v33).F15())
			var _198_v42 _dafny.Array
			_ = _198_v42
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
			_ = _nw30
			_198_v42 = _nw30
			var _199_v43 D5
			_ = _199_v43
			_199_v43 = Companion_D5_.Create_DC14_(_177_v33.F14(), _198_v42, _134_v0)
			var _200_v44 _dafny.MultiSet
			_ = _200_v44
			_200_v44 = _dafny.MultiSetOf(_199_v43, _199_v43, _199_v43, _199_v43, _199_v43)
			_200_v44 = ((_200_v44).Update(_199_v43, Companion_Default___.Abs((_dafny.Zero).Minus(_195_cf8)))).Union(_200_v44)
			(_139_globalState).F3 = _134_v0
			var _201_v45 D6
			_ = _201_v45
			_201_v45 = Companion_D6_.Create_DC17_(true, _177_v33)
			(_139_globalState).F0 = (_201_v45).Dtor_cf30()
		} else {
			var _202___mcc_h6 _dafny.Sequence = _source4.Get_().(D1_DC3).Cf4
			_ = _202___mcc_h6
			var _203_cf4 _dafny.Sequence = _202___mcc_h6
			_ = _203_cf4
			var _204_v46 _dafny.Array
			_ = _204_v46
			var _nw31 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
			_ = _nw31
			_204_v46 = _nw31
			var _205_v47 *C1
			_ = _205_v47
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__(_203_cf4, _138_v4, _177_v33.F14(), _134_v0)
			_205_v47 = _nw32
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_204_v46), 0))
			_ = _index10
			(_204_v46).ArraySet1(_205_v47, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_165_v28), 0))
			_ = _index11
			(_165_v28).ArraySet1((_177_v33).F15(), (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_204_v46), 0))
			_ = _index12
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_165_v28), 0))
			_ = _index13
			var _rhs13 *C1 = _205_v47
			_ = _rhs13
			var _rhs14 bool = (_175_v31).F15()
			_ = _rhs14
			var _rhs15 _dafny.Int = _175_v31.F14()
			_ = _rhs15
			var _lhs7 _dafny.Array = _204_v46
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_204_v46), 0))
			_ = _lhs8
			var _lhs9 _dafny.Array = _165_v28
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_165_v28), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = _139_globalState
			_ = _lhs11
			(_lhs7).ArraySet1(_rhs13, (_lhs8).Int())
			(_lhs9).ArraySet1(_rhs14, (_lhs10).Int())
			_lhs11.F11 = _rhs15
			var _206_v48 _dafny.Sequence
			_ = _206_v48
			_206_v48 = _dafny.SeqOf(_177_v33)
			var _207_v49 _dafny.Map
			_ = _207_v49
			_207_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _dafny.IntOfUint32((_206_v48).Cardinality()))
			(_175_v31).F14_set_(Companion_Default___.SafeModuloInt((_207_v49).Cardinality(), _177_v33.F14()))
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_172_v30), 0))
			_ = _index14
			(_172_v30).ArraySet1(_177_v33.F14(), (_index14).Int())
			var _208_v50 _dafny.CodePoint
			_ = _208_v50
			_208_v50 = _dafny.CodePoint('c')
			var _209_v51 _dafny.Set
			_ = _209_v51
			_209_v51 = _dafny.SetOf(_208_v50, _208_v50)
			var _210_v52 _dafny.Array
			_ = _210_v52
			var _nwElement0_4 _dafny.Set = _dafny.SetOf(_208_v50)
			_ = _nwElement0_4
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(2))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_4, 0)
			(_nw33).ArraySet1(_209_v51, 1)
			_210_v52 = _nw33
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_210_v52), 0))
			_ = _index15
			(_210_v52).ArraySet1(_209_v51, (_index15).Int())
			var _211_v53 _dafny.Array
			_ = _211_v53
			var _nw34 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw34
			_211_v53 = _nw34
			var _212_v54 *C0
			_ = _212_v54
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(_134_v0, (_175_v31).F15(), _203_cf4, _175_v31.F14(), _134_v0)
			_212_v54 = _nw35
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_211_v53), 0))
			_ = _index16
			(_211_v53).ArraySet1(_212_v54, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_172_v30), 0))
			_ = _index17
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_210_v52), 0))
			_ = _index18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_211_v53), 0))
			_ = _index19
			var _rhs16 _dafny.Int = _137_v3
			_ = _rhs16
			var _rhs17 _dafny.Set = (_dafny.SetOf(_208_v50)).Union(_209_v51)
			_ = _rhs17
			var _rhs18 *C0 = _212_v54
			_ = _rhs18
			var _lhs12 _dafny.Array = _172_v30
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_172_v30), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _210_v52
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_210_v52), 0))
			_ = _lhs15
			var _lhs16 _dafny.Array = _211_v53
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_211_v53), 0))
			_ = _lhs17
			(_lhs12).ArraySet1(_rhs16, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs17, (_lhs15).Int())
			(_lhs16).ArraySet1(_rhs18, (_lhs17).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_165_v28), 0))
			_ = _index20
			(_165_v28).ArraySet1(_134_v0, (_index20).Int())
		}
		var _213_v55 _dafny.Array
		_ = _213_v55
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_6
		var _nw36 _dafny.Array
		_ = _nw36
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw36 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.CodePoint = func(_214_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw36 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw36).ArraySet1CodePoint(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw36).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_213_v55 = _nw36
		var _215_v56 _dafny.CodePoint
		_ = _215_v56
		_215_v56 = _dafny.CodePoint('y')
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_213_v55), 0))
		_ = _index21
		(_213_v55).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_175_v31).F15() {
				return _215_v56
			}
			return _dafny.CodePoint('t')
		})(), (_index21).Int())
		var _216_v57 _dafny.Array
		_ = _216_v57
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_7
		var _nw37 _dafny.Array
		_ = _nw37
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw37 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Sequence = func(_217_i6 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("esgshe")
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw37 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw37).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw37).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_216_v57 = _nw37
		var _218_v58 _dafny.Array
		_ = _218_v58
		_218_v58 = _216_v57
		var _219_v59 D6
		_ = _219_v59
		_219_v59 = Companion_D6_.Create_DC18_(_175_v31.F14(), _215_v56)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_172_v30), 0))
		_ = _index22
		(_172_v30).ArraySet1((_219_v59).Dtor_cf32(), (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_213_v55), 0))
		_ = _index23
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_172_v30), 0))
		_ = _index24
		var _rhs19 bool = (_175_v31).F15()
		_ = _rhs19
		var _rhs20 _dafny.CodePoint = _215_v56
		_ = _rhs20
		var _rhs21 _dafny.Array = _216_v57
		_ = _rhs21
		var _rhs22 _dafny.Int = _177_v33.F14()
		_ = _rhs22
		var _rhs23 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-341))
		_ = _rhs23
		var _lhs18 _dafny.Array = _213_v55
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_213_v55), 0))
		_ = _lhs19
		var _lhs20 T0 = _175_v31
		_ = _lhs20
		var _lhs21 _dafny.Array = _172_v30
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_172_v30), 0))
		_ = _lhs22
		_134_v0 = _rhs19
		(_lhs18).ArraySet1CodePoint(_rhs20, (_lhs19).Int())
		_218_v58 = _rhs21
		_lhs20.F14_set_(_rhs22)
		(_lhs21).ArraySet1(_rhs23, (_lhs22).Int())
		(_139_globalState).F3 = _134_v0
	} else {
		(_139_globalState).F11 = _137_v3
		var _220_v60 *C1
		_ = _220_v60
		var _nw38 *C1 = New_C1_()
		_ = _nw38
		_nw38.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("daadh"), _136_v2, _137_v3, _134_v0)
		_220_v60 = _nw38
		var _221_v61 _dafny.Map
		_ = _221_v61
		_221_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v60, _134_v0)
		_221_v61 = (_221_v61).Update(_220_v60, _134_v0)
		if (_134_v0) || ((_137_v3).Cmp(_dafny.IntOfInt64(612)) != 0) {
			var _222_v62 _dafny.CodePoint
			_ = _222_v62
			_222_v62 = _dafny.CodePoint('v')
			var _223_v63 _dafny.Set
			_ = _223_v63
			_223_v63 = _dafny.SetOf((_dafny.Zero).Minus(_137_v3))
			var _224_v64 _dafny.Sequence
			_ = _224_v64
			_224_v64 = _dafny.SeqOf(Companion_Default___.Fm12(_137_v3, (_220_v60).Fm0(_134_v0, _139_globalState), _dafny.IntOfInt64(699), _222_v62, _139_globalState), _223_v63)
			var _225_v65 _dafny.Array
			_ = _225_v65
			var _nwElement0_5 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm17(_137_v3, _139_globalState), _224_v64)
			_ = _nwElement0_5
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(5))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_5, 0)
			(_nw39).ArraySet1(_224_v64, 1)
			(_nw39).ArraySet1(_224_v64, 2)
			(_nw39).ArraySet1(_224_v64, 3)
			(_nw39).ArraySet1(_224_v64, 4)
			_225_v65 = _nw39
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_225_v65), 0))
			_ = _index25
			(_225_v65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(188))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_226_v3 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_227_i7 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_226_v3)
				}
			})(_137_v3))), (_index25).Int())
			var _228_v66 _dafny.Sequence
			_ = _228_v66
			_228_v66 = _dafny.SeqOf(_224_v64, _224_v64, _224_v64)
			var _229_v67 _dafny.Map
			_ = _229_v67
			_229_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(595), _134_v0)
			var _230_v68 _dafny.MultiSet
			_ = _230_v68
			_230_v68 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_220_v60).Fm8(_134_v0, _139_globalState)).Cardinality()), _137_v3, (_dafny.Zero).Minus(_137_v3), _137_v3, _137_v3)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_225_v65), 0))
			_ = _index26
			(_225_v65).ArraySet1((_228_v66).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (func() bool {
					if (_229_v67).Contains(_137_v3) {
						return (_229_v67).Get(_137_v3).(bool)
					}
					return !(_134_v0)
				})() {
					return _137_v3
				}
				return (_220_v60).Fm1(_230_v68, _134_v0, _139_globalState)
			})(), _dafny.IntOfUint32((_228_v66).Cardinality()))).Uint32()).(_dafny.Sequence), (_index26).Int())
			var _231_v69 T1
			_ = _231_v69
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__(!(_134_v0), _134_v0, (_220_v60).F20(), _dafny.IntOfInt64(224), _134_v0)
			_231_v69 = _nw40
			var _232_v70 D2
			_ = _232_v70
			_232_v70 = Companion_D2_.Create_DC7_(false, false, _dafny.IntOfInt64(-533))
			var _233_v71 *C2
			_ = _233_v71
			var _nw41 *C2 = New_C2_()
			_ = _nw41
			_nw41.Ctor__(_232_v70, (_dafny.SetOf(_137_v3)).Cardinality(), (_231_v69).F15())
			_233_v71 = _nw41
			var _234_v72 _dafny.Map
			_ = _234_v72
			_234_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v69, _233_v71)
			(_139_globalState).F11 = ((_234_v72).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v69, _233_v71))).Cardinality()
			_229_v67 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rhwepn")).Cardinality()), !((_231_v69).F15()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_230_v68).Cardinality(), _134_v0))
			(_139_globalState).F11 = (_dafny.Zero).Minus(_137_v3)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_165_v28), 0))
			_ = _index27
			(_165_v28).ArraySet1((_231_v69).F15(), (_index27).Int())
			var _235_v73 _dafny.Map
			_ = _235_v73
			_235_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v62, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(50))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_236_v62 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_237_i8 _dafny.Int) _dafny.CodePoint {
					return _236_v62
				}
			})(_222_v62))))
			var _238_v74 _dafny.Map
			_ = _238_v74
			_238_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_231_v69).F15(), _134_v0)
			var _239_v75 _dafny.Map
			_ = _239_v75
			_239_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_220_v60).Fm0(_134_v0, _139_globalState) {
					return _238_v74
				}
				return _238_v74
			})(), _134_v0)
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_165_v28), 0))
			_ = _index28
			var _rhs24 bool = _134_v0
			_ = _rhs24
			var _rhs25 _dafny.Map = (_235_v73).Update(_222_v62, _135_v1)
			_ = _rhs25
			var _rhs26 _dafny.Map = ((_239_v75).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_238_v74, _134_v0))).Merge(func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter11 := _dafny.Iterate((_dafny.SeqOf(_238_v74, _238_v74, _238_v74)).Elements()); ; {
					_compr_9, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _240_v76 _dafny.Map
					_240_v76 = interface{}(_compr_9).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_238_v74, _238_v74, _238_v74), _240_v76) {
						_coll9.Add(_240_v76, _134_v0)
					}
				}
				return _coll9.ToMap()
			}())
			_ = _rhs26
			var _rhs27 *C2 = _233_v71
			_ = _rhs27
			var _rhs28 _dafny.Int = _137_v3
			_ = _rhs28
			var _lhs23 _dafny.Array = _165_v28
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_165_v28), 0))
			_ = _lhs24
			var _lhs25 *GlobalState = _139_globalState
			_ = _lhs25
			(_lhs23).ArraySet1(_rhs24, (_lhs24).Int())
			_235_v73 = _rhs25
			_239_v75 = _rhs26
			_233_v71 = _rhs27
			_lhs25.F11 = _rhs28
		} else {
			var _241_v77 _dafny.CodePoint
			_ = _241_v77
			_241_v77 = _dafny.CodePoint('y')
			var _242_v78 _dafny.Set
			_ = _242_v78
			_242_v78 = _dafny.SetOf(_241_v77)
			var _243_v79 T0
			_ = _243_v79
			var _nw42 *C1 = New_C1_()
			_ = _nw42
			_nw42.Ctor__(_dafny.Companion_Sequence_.Update(_135_v1, (Companion_Default___.SafeIndex((_242_v78).Cardinality(), _dafny.IntOfUint32((_135_v1).Cardinality()))).Uint32(), _241_v77), (_220_v60).F21(), _137_v3, (_220_v60).Fm0(true, _139_globalState))
			_243_v79 = _nw42
			_243_v79 = _243_v79
			(_139_globalState).F11 = (_dafny.Zero).Minus(_243_v79.F14())
			var _244_v80 T1
			_ = _244_v80
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__(_134_v0, _134_v0, (_220_v60).F20(), _137_v3, _134_v0)
			_244_v80 = _nw43
			var _245_v81 _dafny.Map
			_ = _245_v81
			_245_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v1, _244_v80)
			var _246_v82 _dafny.Array
			_ = _246_v82
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_8
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = (func(_247_v80 T1) func(_dafny.Int) _dafny.Int {
					return func(_248_i9 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_248_i9, _247_v80.F14())
					}
				})(_244_v80)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw44 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw44).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw44).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_246_v82 = _nw44
			var _249_v83 _dafny.Set
			_ = _249_v83
			_249_v83 = _dafny.SetOf(_246_v82, _246_v82, _246_v82)
			var _250_v84 _dafny.Map
			_ = _250_v84
			_250_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_137_v3, (_245_v81).Cardinality()), (_249_v83).Cardinality())
			_250_v84 = (_250_v84).Update(_244_v80.F14(), _244_v80.F14())
			var _251_v85 _dafny.Map
			_ = _251_v85
			_251_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, (_243_v79).F15())
			var _252_v86 _dafny.Map
			_ = _252_v86
			_252_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_251_v85).Update(_244_v80.F14(), _134_v0), _165_v28)
			Companion_Default___.M0(_252_v86, _139_globalState)
			(_139_globalState).F11 = (_243_v79.F14()).Minus(_dafny.IntOfInt64(-435))
		}
		var _253_v87 _dafny.MultiSet
		_ = _253_v87
		_253_v87 = _dafny.MultiSetOf((_dafny.Zero).Minus(_137_v3))
		var _254_v88 T0
		_ = _254_v88
		var _nw45 *C1 = New_C1_()
		_ = _nw45
		_nw45.Ctor__(_dafny.Companion_Sequence_.Concatenate((_220_v60).F20(), (_220_v60).F20()), Companion_Default___.Fm18(_139_globalState), (func() _dafny.Int {
			if _134_v0 {
				return _dafny.IntOfUint32((_136_v2).Cardinality())
			}
			return (_253_v87).Cardinality()
		})(), !(_134_v0))
		_254_v88 = _nw45
		var _255_v89 _dafny.Sequence
		_ = _255_v89
		_255_v89 = _dafny.SeqOf(_254_v88)
		_254_v88 = (_255_v89).Select((Companion_Default___.SafeIndex((_137_v3).Minus(_254_v88.F14()), _dafny.IntOfUint32((_255_v89).Cardinality()))).Uint32()).(T0)
		var _256_v90 D2
		_ = _256_v90
		_256_v90 = Companion_D2_.Create_DC7_((_254_v88).F15(), _134_v0, _137_v3)
		var _nw46 *C2 = New_C2_()
		_ = _nw46
		_nw46.Ctor__(_256_v90, _254_v88.F14(), (_254_v88).F15())
		_254_v88 = _nw46
	}
	var _hi0 _dafny.Int = _dafny.IntOfInt64(583)
	_ = _hi0
	for _257_i10 := Companion_Default___.SafeDivisionInt(_137_v3, _137_v3); _257_i10.Cmp(_hi0) < 0; _257_i10 = _257_i10.Plus(_dafny.One) {
		var _258_v91 _dafny.Map
		_ = _258_v91
		_258_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_i10, _257_i10)
		var _259_v92 D2
		_ = _259_v92
		_259_v92 = Companion_D2_.Create_DC7_(_134_v0, false, _257_i10)
		var _260_v93 _dafny.Set
		_ = _260_v93
		_260_v93 = _dafny.SetOf(Companion_D2_.Create_DC7_(false, _134_v0, (func() _dafny.Int {
			if (_258_v91).Contains(_137_v3) {
				return (_258_v91).Get(_137_v3).(_dafny.Int)
			}
			return _137_v3
		})()), _259_v92)
		var _261_v94 D6
		_ = _261_v94
		_261_v94 = Companion_D6_.Create_DC16_(_dafny.SetOf(_260_v93))
		var _262_v95 _dafny.Set
		_ = _262_v95
		_262_v95 = _dafny.SetOf(_260_v93)
		var _pat_let_tv8 = _262_v95
		_ = _pat_let_tv8
		_261_v94 = func(_pat_let3_0 D6) D6 {
			return func(_263_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let4_0 _dafny.Set) D6 {
					return func(_264_dt__update_hcf29_h0 _dafny.Set) D6 {
						return Companion_D6_.Create_DC16_(_264_dt__update_hcf29_h0)
					}(_pat_let4_0)
				}(_pat_let_tv8)
			}(_pat_let3_0)
		}(Companion_D6_.Create_DC16_(_262_v95))
		var _265_v96 _dafny.CodePoint
		_ = _265_v96
		_265_v96 = _dafny.CodePoint('q')
		var _266_v97 T0
		_ = _266_v97
		var _nw47 *C1 = New_C1_()
		_ = _nw47
		_nw47.Ctor__(_135_v1, _136_v2, _257_i10, _134_v0)
		_266_v97 = _nw47
		var _267_v98 _dafny.Map
		_ = _267_v98
		_267_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v97, (_266_v97).F15())
		var _268_v99 _dafny.Sequence
		_ = _268_v99
		_268_v99 = _dafny.SeqOf((_267_v98).Cardinality())
		var _269_v100 T1
		_ = _269_v100
		var _nw48 *C0 = New_C0_()
		_ = _nw48
		_nw48.Ctor__((_266_v97).F15(), (_266_v97).F15(), _135_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-637))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_270_v96 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_271_i11 _dafny.Int) _dafny.CodePoint {
				return _270_v96
			}
		})(_265_v96)))).Cardinality()), _134_v0)
		_269_v100 = _nw48
		var _272_v101 _dafny.Map
		_ = _272_v101
		_272_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, _269_v100)
		var _273_v102 _dafny.Map
		_ = _273_v102
		_273_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm15(_265_v96, _257_i10, _dafny.IntOfUint32((_268_v99).Cardinality()), _139_globalState), Companion_Default___.Fm20(_dafny.IntOfInt64(647), (_272_v101).Cardinality(), _134_v0, _139_globalState))
		var _274_v103 D6
		_ = _274_v103
		_274_v103 = Companion_D6_.Create_DC17_(_134_v0, _269_v100)
		var _275_v104 _dafny.Sequence
		_ = _275_v104
		_275_v104 = _dafny.SeqOf((_274_v103).Dtor_cf31())
		var _276_v106 _dafny.Map
		_ = _276_v106
		_276_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(418), _dafny.IntOfInt64(104))); ; {
				_compr_10, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _277_v105 _dafny.Int
				_277_v105 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(418)).Cmp(_277_v105) <= 0) && ((_277_v105).Cmp(_dafny.IntOfInt64(104)) < 0) {
					_coll10.Add((_277_v105).Minus(_266_v97.F14()), _134_v0)
				}
			}
			return _coll10.ToMap()
		}(), (_266_v97).F15())
		var _278_v107 _dafny.Array
		_ = _278_v107
		var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw49
		_278_v107 = _nw49
		var _279_v108 _dafny.Sequence
		_ = _279_v108
		_279_v108 = _dafny.SeqOf(_278_v107, _278_v107)
		var _280_v109 _dafny.Set
		_ = _280_v109
		_280_v109 = _dafny.SetOf((_269_v100).F15())
		var _281_v110 _dafny.Map
		_ = _281_v110
		_281_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v28, (_280_v109).Cardinality())
		var _282_v111 _dafny.Map
		_ = _282_v111
		_282_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _137_v3)
		var _283_v112 _dafny.Map
		_ = _283_v112
		_283_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_282_v111).Cardinality(), (_266_v97).F15())
		var _284_v113 _dafny.Array
		_ = _284_v113
		var _nwElement0_6 _dafny.Int = _257_i10
		_ = _nwElement0_6
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(27))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_6, 0)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ninu")).Cardinality()), 1)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_275_v104).Cardinality()), 2)
		(_nw50).ArraySet1(_dafny.IntOfInt64(-917), 3)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()), 4)
		(_nw50).ArraySet1(_137_v3, 5)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_135_v1).Cardinality()), 6)
		(_nw50).ArraySet1(_137_v3, 7)
		(_nw50).ArraySet1((_276_v106).Cardinality(), 8)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_279_v108).Cardinality()), 9)
		(_nw50).ArraySet1(_257_i10, 10)
		(_nw50).ArraySet1(_257_i10, 11)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_136_v2).Cardinality()), 12)
		(_nw50).ArraySet1(_137_v3, 13)
		(_nw50).ArraySet1(_257_i10, 14)
		(_nw50).ArraySet1(_137_v3, 15)
		(_nw50).ArraySet1(_137_v3, 16)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_135_v1).Cardinality()), 17)
		(_nw50).ArraySet1((_268_v99).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_136_v2).Cardinality()), _dafny.IntOfUint32((_268_v99).Cardinality()))).Uint32()).(_dafny.Int), 18)
		(_nw50).ArraySet1(_dafny.IntOfInt64(413), 19)
		(_nw50).ArraySet1(_257_i10, 20)
		(_nw50).ArraySet1((_281_v110).Cardinality(), 21)
		(_nw50).ArraySet1(_269_v100.F14(), 22)
		(_nw50).ArraySet1((_dafny.Zero).Minus((_283_v112).Cardinality()), 23)
		(_nw50).ArraySet1(_269_v100.F14(), 24)
		(_nw50).ArraySet1(_257_i10, 25)
		(_nw50).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, (_269_v100).F15())).Cardinality(), 26)
		_284_v113 = _nw50
		var _285_v114 _dafny.MultiSet
		_ = _285_v114
		_285_v114 = _dafny.MultiSetOf(_284_v113, _284_v113, _284_v113)
		var _286_v115 _dafny.Sequence
		_ = _286_v115
		_286_v115 = _dafny.SeqOf(Companion_Default___.Fm19(_139_globalState), (_273_v102).Cardinality(), (_285_v114).Cardinality(), _137_v3)
		_286_v115 = _268_v99
		var _287_v116 _dafny.Map
		_ = _287_v116
		_287_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v96, Companion_D1_.Create_DC3_(_135_v1))
		var _288_v117 _dafny.Map
		_ = _288_v117
		_288_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_284_v113, _287_v116)
		_137_v3 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_266_v97.F14()).Plus(_266_v97.F14()), (func() _dafny.Map {
			if (_288_v117).Contains(_284_v113) {
				return (_288_v117).Get(_284_v113).(_dafny.Map)
			}
			return _287_v116
		})())).Cardinality()
		(_266_v97).F14_set_(Companion_Default___.Fm19(_139_globalState))
	}
	_137_v3 = _137_v3
	var _289_v118 D3
	_ = _289_v118
	_289_v118 = Companion_D3_.Create_DC10_(_135_v1, _165_v28, _134_v0, _137_v3, _dafny.IntOfInt64(-289))
	var _290_v119 _dafny.Array
	_ = _290_v119
	var _nwElement0_7 _dafny.Sequence = _135_v1
	_ = _nwElement0_7
	var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(21))
	_ = _nw51
	(_nw51).ArraySet1(_nwElement0_7, 0)
	(_nw51).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_289_v118).Dtor_cf17(), _135_v1), 1)
	(_nw51).ArraySet1(_135_v1, 2)
	(_nw51).ArraySet1(_135_v1, 3)
	(_nw51).ArraySet1(_135_v1, 4)
	(_nw51).ArraySet1(_135_v1, 5)
	(_nw51).ArraySet1(_135_v1, 6)
	(_nw51).ArraySet1(_135_v1, 7)
	(_nw51).ArraySet1(_135_v1, 8)
	(_nw51).ArraySet1(_135_v1, 9)
	(_nw51).ArraySet1(Companion_Default___.Fm9(_137_v3, _139_globalState), 10)
	(_nw51).ArraySet1(_135_v1, 11)
	(_nw51).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qwdbsyk"), 12)
	(_nw51).ArraySet1(_135_v1, 13)
	(_nw51).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_135_v1, _dafny.UnicodeSeqOfUtf8Bytes("iiayqcdiq")), 14)
	(_nw51).ArraySet1(_135_v1, 15)
	(_nw51).ArraySet1(_135_v1, 16)
	(_nw51).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hbwlhhvq"), 17)
	(_nw51).ArraySet1(_135_v1, 18)
	(_nw51).ArraySet1(_135_v1, 19)
	(_nw51).ArraySet1(_135_v1, 20)
	_290_v119 = _nw51
	for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_290_v119), 0))); ; {
		_guard_loop_2, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		var _291_i12 _dafny.Int
		_291_i12 = interface{}(_guard_loop_2).(_dafny.Int)
		if (true) && (((_291_i12).Sign() != -1) && ((_291_i12).Cmp(_dafny.ArrayLen((_290_v119), 0)) < 0)) {
			(_290_v119).ArraySet1((func() _dafny.Sequence {
				if (false) == (_134_v0) {
					return _135_v1
				}
				return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}(func(_292_i13 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				})), (Companion_Default___.SafeIndex((_dafny.SetOf(false, _134_v0)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}(func(_293_i13 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))).Cardinality()))).Uint32(), _dafny.CodePoint('g'))
			})(), (_291_i12).Int())
		}
	}
	var _294_v120 _dafny.Map
	_ = _294_v120
	_294_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, _dafny.IntOfInt64(622))
	var _295_i14 _dafny.Int
	_ = _295_i14
	_295_i14 = _dafny.Zero
	{
		for (func() bool {
			if (_134_v0) || (_134_v0) {
				return _134_v0
			}
			return (_134_v0) == (Companion_Default___.Fm21(_137_v3, !(false), !(_134_v0), _294_v120, _139_globalState))
		})() {
			{
				if (_295_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_295_i14 = (_295_i14).Plus(_dafny.One)
				var _296_v121 _dafny.Sequence
				_ = _296_v121
				_296_v121 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_135_v1).Cardinality()), _dafny.IntOfUint32((_138_v4).Cardinality())))
				_296_v121 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_296_v121, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm19(_139_globalState))), _dafny.IntOfUint32((_296_v121).Cardinality()))).Uint32(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(864), _137_v3)), (Companion_Default___.SafeIndex(_137_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_296_v121, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm19(_139_globalState))), _dafny.IntOfUint32((_296_v121).Cardinality()))).Uint32(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(864), _137_v3))).Cardinality()))).Uint32(), _dafny.IntOfInt64(366))
				_135_v1 = _dafny.UnicodeSeqOfUtf8Bytes("cuplutb")
				var _297_v122 _dafny.CodePoint
				_ = _297_v122
				_297_v122 = _dafny.CodePoint('u')
				if (!((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_135_v1, (Companion_Default___.SafeIndex(_137_v3, _dafny.IntOfUint32((_135_v1).Cardinality()))).Uint32(), _297_v122)).Cardinality())).Cmp(_137_v3) < 0)) || ((_137_v3).Cmp(_dafny.IntOfInt64(911)) != 0) {
					var _298_v123 *C0
					_ = _298_v123
					var _nw52 *C0 = New_C0_()
					_ = _nw52
					_nw52.Ctor__((_137_v3).Cmp(_dafny.IntOfInt64(79)) == 0, (_134_v0) || (_134_v0), _135_v1, _137_v3, (func() bool {
						if _134_v0 {
							return _134_v0
						}
						return false
					})())
					_298_v123 = _nw52
					var _299_v124 _dafny.Map
					_ = _299_v124
					_299_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, (_137_v3).Cmp(_137_v3) > 0)
					_299_v124 = _299_v124
					var _300_v125 *C1
					_ = _300_v125
					var _nw53 *C1 = New_C1_()
					_ = _nw53
					_nw53.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-172))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}((func(_301_v122 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_302_i15 _dafny.Int) _dafny.CodePoint {
							return _301_v122
						}
					})(_297_v122))), _138_v4, _dafny.IntOfUint32((_135_v1).Cardinality()), (_298_v123).F17())
					_300_v125 = _nw53
					_300_v125 = _300_v125
					var _303_v126 _dafny.Map
					_ = _303_v126
					_303_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_299_v124, _165_v28)
					var _304_v127 _dafny.Sequence
					_ = _304_v127
					_304_v127 = _dafny.SeqOf(_303_v126)
					Companion_Default___.M0(((_304_v127).Select((Companion_Default___.SafeIndex(_137_v3, _dafny.IntOfUint32((_304_v127).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_303_v126), _139_globalState)
					var _305_v128 _dafny.MultiSet
					_ = _305_v128
					_305_v128 = _dafny.MultiSetOf(_dafny.IntOfUint32((_135_v1).Cardinality()), _137_v3)
					(_139_globalState).F3 = (func() bool {
						if (_299_v124).Contains((Companion_Default___.Fm13((_305_v128).Cardinality(), _139_globalState)).Cardinality()) {
							return (_299_v124).Get((Companion_Default___.Fm13((_305_v128).Cardinality(), _139_globalState)).Cardinality()).(bool)
						}
						return (_298_v123).F17()
					})()
				} else {
					var _306_v129 _dafny.Array
					_ = _306_v129
					var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
					_ = _nw54
					_306_v129 = _nw54
					var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_306_v129), 0))
					_ = _index29
					(_306_v129).ArraySet1(_137_v3, (_index29).Int())
					var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_306_v129), 0))
					_ = _index30
					(_306_v129).ArraySet1(_137_v3, (_index30).Int())
					_137_v3 = (_306_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_306_v129), 0))).Int()).(_dafny.Int)
					var _307_v130 D3
					_ = _307_v130
					_307_v130 = Companion_D3_.Create_DC8_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(309))).Uint32(), func(coer17 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}(func(_308_i16 _dafny.Int) D1 {
						return Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("psxioq"))
					})))
					var _309_v131 _dafny.MultiSet
					_ = _309_v131
					_309_v131 = _dafny.MultiSetOf(_134_v0)
					_307_v130 = (func() D3 {
						if (_309_v131).IsSubsetOf(_309_v131) {
							return _307_v130
						}
						return _307_v130
					})()
					(_139_globalState).F0 = (!(_134_v0)) || (_134_v0)
					var _310_v132 _dafny.Map
					_ = _310_v132
					_310_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_306_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_306_v129), 0))).Int()).(_dafny.Int), _134_v0)
					var _311_v133 _dafny.Map
					_ = _311_v133
					_311_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v132, _165_v28)
					Companion_Default___.M0(_311_v133, _139_globalState)
				}
				var _312_v134 _dafny.Map
				_ = _312_v134
				_312_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _137_v3)
				var _313_v135 _dafny.Map
				_ = _313_v135
				_313_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_135_v1).Cardinality()), (_312_v134).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, (_294_v120).Cardinality())))
				_313_v135 = (_313_v135).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(935), _312_v134))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _314_v136 _dafny.MultiSet
	_ = _314_v136
	_314_v136 = _dafny.MultiSetOf(_137_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_315_i17 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))).Cardinality()), _137_v3)
	var _316_v137 _dafny.Sequence
	_ = _316_v137
	_316_v137 = _dafny.SeqOf(_137_v3, _137_v3)
	var _317_v138 *C0
	_ = _317_v138
	var _nw55 *C0 = New_C0_()
	_ = _nw55
	_nw55.Ctor__((_137_v3).Cmp(_137_v3) < 0, (_314_v136).IsProperSubsetOf(_314_v136), _135_v1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_137_v3), _316_v137)).Cardinality())), !(Companion_Default___.Fm21(_137_v3, _134_v0, _134_v0, _294_v120, _139_globalState)))
	_317_v138 = _nw55
	var _318_v139 _dafny.Array
	_ = _318_v139
	var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
	_ = _nw56
	_318_v139 = _nw56
	_318_v139 = _318_v139
	for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_318_v139), 0))); ; {
		_guard_loop_3, _ok14 := _iter14()
		if !_ok14 {
			break
		}
		var _319_i18 _dafny.Int
		_319_i18 = interface{}(_guard_loop_3).(_dafny.Int)
		if (true) && (((_319_i18).Sign() != -1) && ((_319_i18).Cmp(_dafny.ArrayLen((_318_v139), 0)) < 0)) {
			(_318_v139).ArraySet1(Companion_Default___.SafeModuloInt(_319_i18, _137_v3), (_319_i18).Int())
		}
	}
	var _320_v140 _dafny.MultiSet
	_ = _320_v140
	_320_v140 = _dafny.MultiSetOf(true)
	(_139_globalState).F0 = (_320_v140).IsDisjointFrom(_320_v140)
	if !((_317_v138).F18()) {
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))
		_ = _index31
		(_165_v28).ArraySet1((_317_v138).F17(), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))
		_ = _index32
		(_165_v28).ArraySet1((_317_v138).F18(), (_index32).Int())
		var _321_v141 _dafny.Array
		_ = _321_v141
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_9
		var _nw57 _dafny.Array
		_ = _nw57
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw57 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Set = (func(_322_v3 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_323_i19 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_322_v3)
				}
			})(_137_v3)
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
		_321_v141 = _nw57
		var _324_v142 _dafny.Set
		_ = _324_v142
		_324_v142 = _dafny.SetOf((func() _dafny.Int {
			if (_314_v136).Contains((func() _dafny.Int {
				if (_320_v140).Contains(_134_v0) {
					return (_320_v140).Multiplicity(_134_v0)
				}
				return _137_v3
			})()) {
				return (_314_v136).Multiplicity((func() _dafny.Int {
					if (_320_v140).Contains(_134_v0) {
						return (_320_v140).Multiplicity(_134_v0)
					}
					return _137_v3
				})())
			}
			return _137_v3
		})())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_321_v141), 0))
		_ = _index33
		(_321_v141).ArraySet1(_324_v142, (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_318_v139), 0))
		_ = _index34
		(_318_v139).ArraySet1(_137_v3, (_index34).Int())
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_321_v141), 0))
		_ = _index35
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_318_v139), 0))
		_ = _index36
		var _rhs29 _dafny.Set = _324_v142
		_ = _rhs29
		var _rhs30 _dafny.Int = (_dafny.Zero).Minus(_137_v3)
		_ = _rhs30
		var _rhs31 _dafny.Int = _137_v3
		_ = _rhs31
		var _rhs32 _dafny.Int = ((func() _dafny.Int {
			if (_314_v136).Contains(_137_v3) {
				return (_314_v136).Multiplicity(_137_v3)
			}
			return (_317_v138).Fm2(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(241))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_325_i20 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), (_317_v138).F18(), (_165_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))).Int()).(bool), _137_v3, _139_globalState)
		})()).Times(_137_v3)
		_ = _rhs32
		var _lhs26 _dafny.Array = _321_v141
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_321_v141), 0))
		_ = _lhs27
		var _lhs28 *GlobalState = _139_globalState
		_ = _lhs28
		var _lhs29 *GlobalState = _139_globalState
		_ = _lhs29
		var _lhs30 _dafny.Array = _318_v139
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_318_v139), 0))
		_ = _lhs31
		(_lhs26).ArraySet1(_rhs29, (_lhs27).Int())
		_lhs28.F11 = _rhs30
		_lhs29.F11 = _rhs31
		(_lhs30).ArraySet1(_rhs32, (_lhs31).Int())
		var _326_v143 D2
		_ = _326_v143
		_326_v143 = Companion_D2_.Create_DC7_((_317_v138).F18(), _134_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_316_v137).Cardinality()), (_165_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))).Int()).(bool))).Cardinality())
		var _327_v144 *C2
		_ = _327_v144
		var _nw58 *C2 = New_C2_()
		_ = _nw58
		_nw58.Ctor__(_326_v143, (_320_v140).Cardinality(), true)
		_327_v144 = _nw58
		var _328_v145 _dafny.Map
		_ = _328_v145
		_328_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm21((_318_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_318_v139), 0))).Int()).(_dafny.Int), (_317_v138).F18(), _134_v0, _294_v120, _139_globalState), _327_v144)
		var _329_v146 _dafny.Array
		_ = _329_v146
		var _nwElement0_8 *C2 = _327_v144
		_ = _nwElement0_8
		var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(24))
		_ = _nw59
		(_nw59).ArraySet1(_nwElement0_8, 0)
		(_nw59).ArraySet1((func() *C2 {
			if (_317_v138).F18() {
				return _327_v144
			}
			return _327_v144
		})(), 1)
		(_nw59).ArraySet1(_327_v144, 2)
		(_nw59).ArraySet1(_327_v144, 3)
		(_nw59).ArraySet1(_327_v144, 4)
		(_nw59).ArraySet1(_327_v144, 5)
		(_nw59).ArraySet1(_327_v144, 6)
		(_nw59).ArraySet1(_327_v144, 7)
		(_nw59).ArraySet1(_327_v144, 8)
		(_nw59).ArraySet1((func() *C2 {
			if (_328_v145).Contains(_134_v0) {
				return (_328_v145).Get(_134_v0).(*C2)
			}
			return _327_v144
		})(), 9)
		(_nw59).ArraySet1(_327_v144, 10)
		(_nw59).ArraySet1(_327_v144, 11)
		(_nw59).ArraySet1(_327_v144, 12)
		(_nw59).ArraySet1(_327_v144, 13)
		(_nw59).ArraySet1(_327_v144, 14)
		(_nw59).ArraySet1(_327_v144, 15)
		(_nw59).ArraySet1(_327_v144, 16)
		(_nw59).ArraySet1(_327_v144, 17)
		(_nw59).ArraySet1(_327_v144, 18)
		(_nw59).ArraySet1(_327_v144, 19)
		(_nw59).ArraySet1(_327_v144, 20)
		(_nw59).ArraySet1(_327_v144, 21)
		(_nw59).ArraySet1(_327_v144, 22)
		(_nw59).ArraySet1(_327_v144, 23)
		_329_v146 = _nw59
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_329_v146), 0))
		_ = _index37
		(_329_v146).ArraySet1(_327_v144, (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_329_v146), 0))
		_ = _index38
		(_329_v146).ArraySet1(_327_v144, (_index38).Int())
		var _330_v147 *C0
		_ = _330_v147
		var _nw60 *C0 = New_C0_()
		_ = _nw60
		_nw60.Ctor__((_317_v138).F17(), _134_v0, _135_v1, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("syla")).Cardinality())).Minus(_137_v3), _134_v0)
		_330_v147 = _nw60
		var _331_v148 _dafny.Map
		_ = _331_v148
		_331_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v1, (_165_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))).Int()).(bool))
		var _332_v149 _dafny.Map
		_ = _332_v149
		_332_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_331_v148, (_317_v138).F17())
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))
		_ = _index39
		var _rhs33 bool = (_317_v138).F17()
		_ = _rhs33
		var _rhs34 bool = (_330_v147).F18()
		_ = _rhs34
		var _rhs35 bool = (_165_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))).Int()).(bool)
		_ = _rhs35
		var _rhs36 bool = (func() bool {
			if (_332_v149).Contains(_331_v148) {
				return (_332_v149).Get(_331_v148).(bool)
			}
			return (_330_v147).F17()
		})()
		_ = _rhs36
		var _rhs37 bool = (func() bool {
			if false {
				return Companion_Default___.Fm21(_137_v3, !((_317_v138).F17()), _134_v0, _294_v120, _139_globalState)
			}
			return !((_317_v138).F17())
		})()
		_ = _rhs37
		var _lhs32 _dafny.Array = _165_v28
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_165_v28), 0))
		_ = _lhs33
		var _lhs34 *GlobalState = _139_globalState
		_ = _lhs34
		var _lhs35 *GlobalState = _139_globalState
		_ = _lhs35
		(_lhs32).ArraySet1(_rhs33, (_lhs33).Int())
		_134_v0 = _rhs34
		_lhs34.F3 = _rhs35
		_134_v0 = _rhs36
		_lhs35.F0 = _rhs37
	} else {
		(_139_globalState).F3 = (_317_v138).F18()
		if (_317_v138).F18() {
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_165_v28), 0))
			_ = _index40
			(_165_v28).ArraySet1((_317_v138).F17(), (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_165_v28), 0))
			_ = _index41
			(_165_v28).ArraySet1((_317_v138).F18(), (_index41).Int())
			var _333_v150 _dafny.Map
			_ = _333_v150
			_333_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, (_317_v138).F17())
			var _334_v151 _dafny.Array
			_ = _334_v151
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw61
			_334_v151 = _nw61
			var _335_v152 _dafny.Map
			_ = _335_v152
			_335_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v150, _334_v151)
			Companion_Default___.M0((_335_v152).Merge(_335_v152), _139_globalState)
			var _336_v153 D2
			_ = _336_v153
			_336_v153 = Companion_D2_.Create_DC7_((_317_v138).F17(), (_317_v138).F18(), _137_v3)
			var _337_v154 *C1
			_ = _337_v154
			var _nw62 *C1 = New_C1_()
			_ = _nw62
			_nw62.Ctor__(_135_v1, _138_v4, _137_v3, (_317_v138).F18())
			_337_v154 = _nw62
			var _338_v155 _dafny.Map
			_ = _338_v155
			_338_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v154, !((_317_v138).F17()))
			var _339_v156 *C2
			_ = _339_v156
			var _nw63 *C2 = New_C2_()
			_ = _nw63
			_nw63.Ctor__(_336_v153, (_338_v155).Cardinality(), (_dafny.IntOfUint32((_136_v2).Cardinality())).Cmp(_137_v3) >= 0)
			_339_v156 = _nw63
			(_139_globalState).F0 = (func() bool {
				if (_317_v138).F17() {
					return !(_314_v136).Contains(_137_v3)
				}
				return (_317_v138).F18()
			})()
			var _340_v157 _dafny.Map
			_ = _340_v157
			_340_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v3, _320_v140)
			var _341_v158 D2
			_ = _341_v158
			_341_v158 = Companion_D2_.Create_DC6_(_340_v157)
			var _342_v159 _dafny.Int
			_ = _342_v159
			var _343_v160 _dafny.Int
			_ = _343_v160
			var _out0 _dafny.Int
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			_out0, _out1 = (_339_v156).M2(_341_v158, (_317_v138).F17(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_137_v3)), _137_v3, _139_globalState)
			_342_v159 = _out0
			_343_v160 = _out1
		} else {
			var _344_v161 _dafny.Array
			_ = _344_v161
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_10
			var _nw64 _dafny.Array
			_ = _nw64
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw64 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.CodePoint = func(_345_i21 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw64 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw64).ArraySet1CodePoint(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw64).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_344_v161 = _nw64
			var _346_v162 _dafny.CodePoint
			_ = _346_v162
			_346_v162 = _dafny.CodePoint('g')
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_344_v161), 0))
			_ = _index42
			(_344_v161).ArraySet1CodePoint(_346_v162, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_344_v161), 0))
			_ = _index43
			(_344_v161).ArraySet1CodePoint(_346_v162, (_index43).Int())
			var _347_v163 D1
			_ = _347_v163
			_347_v163 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("tcir"))
			_347_v163 = Companion_D1_.Create_DC3_(_135_v1)
			(_139_globalState).F11 = _137_v3
			(_139_globalState).F11 = (_317_v138).Fm1(_314_v136, _134_v0, _139_globalState)
			_165_v28 = _165_v28
		}
		_316_v137 = _dafny.Companion_Sequence_.Concatenate(_316_v137, _316_v137)
		(_139_globalState).F11 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_137_v3, _137_v3), _137_v3)
		var _348_v164 _dafny.Set
		_ = _348_v164
		_348_v164 = _dafny.SetOf((_317_v138).F17())
		var _rhs38 _dafny.Int = _137_v3
		_ = _rhs38
		var _rhs39 _dafny.Int = (_137_v3).Minus(_137_v3)
		_ = _rhs39
		var _rhs40 _dafny.Set = _348_v164
		_ = _rhs40
		_137_v3 = _rhs38
		_137_v3 = _rhs39
		_348_v164 = _rhs40
	}
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_165_v28), 0))
	_ = _index44
	(_165_v28).ArraySet1((_317_v138).Fm0((_317_v138).F17(), _139_globalState), (_index44).Int())
	var _349_v165 _dafny.Map
	_ = _349_v165
	_349_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(255), (_317_v138).F18())
	var _350_v166 D3
	_ = _350_v166
	_350_v166 = Companion_D3_.Create_DC9_((_349_v165).Cardinality())
	var _pat_let_tv9 = _317_v138
	_ = _pat_let_tv9
	var _pat_let_tv10 = _317_v138
	_ = _pat_let_tv10
	var _pat_let_tv11 = _317_v138
	_ = _pat_let_tv11
	var _pat_let_tv12 = _317_v138
	_ = _pat_let_tv12
	var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_165_v28), 0))
	_ = _index45
	(_165_v28).ArraySet1(func(_source5 D3) bool {
		if _source5.Is_DC9() {
			var _351___mcc_h7 _dafny.Int = _source5.Get_().(D3_DC9).Cf16
			_ = _351___mcc_h7
			var _352_cf16 _dafny.Int = _351___mcc_h7
			_ = _352_cf16
			return !((_pat_let_tv9).F18())
		} else if _source5.Is_DC10() {
			var _353___mcc_h8 _dafny.Sequence = _source5.Get_().(D3_DC10).Cf17
			_ = _353___mcc_h8
			var _354___mcc_h9 _dafny.Array = _source5.Get_().(D3_DC10).Cf18
			_ = _354___mcc_h9
			var _355___mcc_h10 bool = _source5.Get_().(D3_DC10).Cf19
			_ = _355___mcc_h10
			var _356___mcc_h11 _dafny.Int = _source5.Get_().(D3_DC10).Cf20
			_ = _356___mcc_h11
			var _357___mcc_h12 _dafny.Int = _source5.Get_().(D3_DC10).Cf21
			_ = _357___mcc_h12
			var _358_cf21 _dafny.Int = _357___mcc_h12
			_ = _358_cf21
			var _359_cf20 _dafny.Int = _356___mcc_h11
			_ = _359_cf20
			var _360_cf19 bool = _355___mcc_h10
			_ = _360_cf19
			var _361_cf18 _dafny.Array = _354___mcc_h9
			_ = _361_cf18
			var _362_cf17 _dafny.Sequence = _353___mcc_h8
			_ = _362_cf17
			return (_pat_let_tv10).F18()
		} else if _source5.Is_DC8() {
			var _363___mcc_h13 _dafny.Sequence = _source5.Get_().(D3_DC8).Cf15
			_ = _363___mcc_h13
			var _364_cf15 _dafny.Sequence = _363___mcc_h13
			_ = _364_cf15
			return (_pat_let_tv11).F18()
		} else {
			var _365___mcc_h14 D3 = _source5.Get_().(D3_DC11).Cf22
			_ = _365___mcc_h14
			var _366_cf22 D3 = _365___mcc_h14
			_ = _366_cf22
			return (_pat_let_tv12).F18()
		}
	}(_350_v166), (_index45).Int())
	var _367_v167 _dafny.Array
	_ = _367_v167
	var _nw65 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
	_ = _nw65
	_367_v167 = _nw65
	_367_v167 = _367_v167
	_dafny.Print(_134_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_135_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_136_v2, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_138_v4, _dafny.SeqOf(false, false, false, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_139_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_139_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_139_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_139_globalState).F13(), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v28).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_169_v29, _dafny.SeqOf(_dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false), _dafny.SeqOf(false, false, false, false, true), _dafny.SeqOf(false), _dafny.SeqOf(false, false, false, false, true, false, false, false, false, true), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v118).Dtor_cf17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_v118).Dtor_cf18()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v118).Dtor_cf19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v118).Dtor_cf20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v118).Dtor_cf21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v119).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_294_v120).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(622))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_295_i14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_314_v136).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.One, _dafny.IntOfInt64(168))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_316_v137, _dafny.SeqOf(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_317_v138).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_317_v138).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v139).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_320_v140).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v165).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(255), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_350_v166).Dtor_cf16())
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
	Cf2 _dafny.Int
	Cf3 _dafny.MultiSet
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 _dafny.Int, Cf3 _dafny.MultiSet) D0 {
	return D0{D0_DC2{Cf2, Cf3}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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
	return Companion_D0_.Create_DC1_(_dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.MultiSet {
	return _this.Get_().(D0_DC2).Cf3
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
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Equals(data2.Cf3)
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

type D1_DC4 struct {
	Cf5 _dafny.Int
	Cf6 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 _dafny.Int, Cf6 _dafny.Int) D1 {
	return D1{D1_DC4{Cf5, Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf7  _dafny.Int
	Cf8  _dafny.Int
	Cf9  bool
	Cf10 _dafny.Map
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 _dafny.Int, Cf8 _dafny.Int, Cf9 bool, Cf10 _dafny.Map) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf4
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
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf4.VerbatimString(true) + ")"
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
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9 && data1.Cf10.Equals(data2.Cf10)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4)
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
	Cf12 bool
	Cf13 bool
	Cf14 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 bool, Cf13 bool, Cf14 _dafny.Int) D2 {
	return D2{D2_DC7{Cf12, Cf13, Cf14}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf11 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Map) D2 {
	return D2{D2_DC6{Cf11}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false, false, _dafny.Zero)
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0
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
	Cf16 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf16 _dafny.Int) D3 {
	return D3{D3_DC9{Cf16}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf17 _dafny.Sequence
	Cf18 _dafny.Array
	Cf19 bool
	Cf20 _dafny.Int
	Cf21 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf17 _dafny.Sequence, Cf18 _dafny.Array, Cf19 bool, Cf20 _dafny.Int, Cf21 _dafny.Int) D3 {
	return D3{D3_DC10{Cf17, Cf18, Cf19, Cf20, Cf21}}
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

type D3_DC11 struct {
	Cf22 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf22 D3) D3 {
	return D3{D3_DC11{Cf22}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.Zero)
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf20
}

func (_this D3) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf21
}

func (_this D3) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) Dtor_cf22() D3 {
	return _this.Get_().(D3_DC11).Cf22
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + data.Cf17.VerbatimString(true) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf17.Equals(data2.Cf17) && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf15.Equals(data2.Cf15)
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf22.Equals(data2.Cf22)
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
	Cf23 _dafny.Array
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf23 _dafny.Array) D4 {
	return D4{D4_DC12{Cf23}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D4) Dtor_cf23() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf23
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf23 == data2.Cf23
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
	Cf25 _dafny.Int
	Cf26 _dafny.Array
	Cf27 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf25 _dafny.Int, Cf26 _dafny.Array, Cf27 bool) D5 {
	return D5{D5_DC14{Cf25, Cf26, Cf27}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf24 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf24 _dafny.Map) D5 {
	return D5{D5_DC13{Cf24}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC15 struct {
	Cf28 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf28 D5) D5 {
	return D5{D5_DC15{Cf28}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D5) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf26
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC14).Cf27
}

func (_this D5) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf24
}

func (_this D5) Dtor_cf28() D5 {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf24.Equals(data2.Cf24)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
	Cf31 T1
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf30 bool, Cf31 T1) D6 {
	return D6{D6_DC17{Cf30, Cf31}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf32 _dafny.Int
	Cf33 _dafny.CodePoint
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf32 _dafny.Int, Cf33 _dafny.CodePoint) D6 {
	return D6{D6_DC18{Cf32, Cf33}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC16 struct {
	Cf29 _dafny.Set
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf29 _dafny.Set) D6 {
	return D6{D6_DC16{Cf29}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC19 struct {
	Cf34 D6
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf34 D6) D6 {
	return D6{D6_DC19{Cf34}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false, (T1)(nil))
}

func (_this D6) Dtor_cf30() bool {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) Dtor_cf31() T1 {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf32
}

func (_this D6) Dtor_cf33() _dafny.CodePoint {
	return _this.Get_().(D6_DC18).Cf33
}

func (_this D6) Dtor_cf29() _dafny.Set {
	return _this.Get_().(D6_DC16).Cf29
}

func (_this D6) Dtor_cf34() D6 {
	return _this.Get_().(D6_DC19).Cf34
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf34) + ")"
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
			return ok && data1.Cf30 == data2.Cf30 && _dafny.AreEqual(data1.Cf31, data2.Cf31)
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33 == data2.Cf33
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf34.Equals(data2.Cf34)
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
	Cf35 _dafny.Array
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf35 _dafny.Array) D7 {
	return D7{D7_DC20{Cf35}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D7) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D7_DC20).Cf35
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf35 == data2.Cf35
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
	Cf37 _dafny.CodePoint
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf37 _dafny.CodePoint) D8 {
	return D8{D8_DC22{Cf37}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC23 struct {
	Cf38 *C2
	Cf39 bool
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf38 *C2, Cf39 bool) D8 {
	return D8{D8_DC23{Cf38, Cf39}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC21 struct {
	Cf36 _dafny.Map
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf36 _dafny.Map) D8 {
	return D8{D8_DC21{Cf36}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC24 struct {
	Cf40 D8
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf40 D8) D8 {
	return D8{D8_DC24{Cf40}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC22_(_dafny.CodePoint('D'))
}

func (_this D8) Dtor_cf37() _dafny.CodePoint {
	return _this.Get_().(D8_DC22).Cf37
}

func (_this D8) Dtor_cf38() *C2 {
	return _this.Get_().(D8_DC23).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC23).Cf39
}

func (_this D8) Dtor_cf36() _dafny.Map {
	return _this.Get_().(D8_DC21).Cf36
}

func (_this D8) Dtor_cf40() D8 {
	return _this.Get_().(D8_DC24).Cf40
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf40) + ")"
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
			return ok && data1.Cf37 == data2.Cf37
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf36.Equals(data2.Cf36)
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf40.Equals(data2.Cf40)
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
	Cf42 _dafny.Int
	Cf43 bool
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf42 _dafny.Int, Cf43 bool) D9 {
	return D9{D9_DC26{Cf42, Cf43}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_() D9 {
	return D9{D9_DC27{}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC25 struct {
	Cf41 _dafny.Array
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf41 _dafny.Array) D9 {
	return D9{D9_DC25{Cf41}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC26_(_dafny.Zero, false)
}

func (_this D9) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf42
}

func (_this D9) Dtor_cf43() bool {
	return _this.Get_().(D9_DC26).Cf43
}

func (_this D9) Dtor_cf41() _dafny.Array {
	return _this.Get_().(D9_DC25).Cf41
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf41) + ")"
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
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43 == data2.Cf43
		}
	case D9_DC27:
		{
			_, ok := other.Get_().(D9_DC27)
			return ok
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf41 == data2.Cf41
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
	Cf45 bool
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf45 bool) D10 {
	return D10{D10_DC29{Cf45}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf44 _dafny.Set
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf44 _dafny.Set) D10 {
	return D10{D10_DC28{Cf44}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(false)
}

func (_this D10) Dtor_cf45() bool {
	return _this.Get_().(D10_DC29).Cf45
}

func (_this D10) Dtor_cf44() _dafny.Set {
	return _this.Get_().(D10_DC28).Cf44
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf44) + ")"
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
			return ok && data1.Cf45 == data2.Cf45
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf44.Equals(data2.Cf44)
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
	F14() _dafny.Int
	F14_set_(value _dafny.Int)
	Fm0(p0 bool, globalState *GlobalState) bool
	Fm1(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Int
	F15() bool
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
	F14() _dafny.Int
	F14_set_(value _dafny.Int)
	Fm0(p0 bool, globalState *GlobalState) bool
	Fm1(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Int
	F15() bool
	F16() _dafny.Sequence
	F16_set_(value _dafny.Sequence)
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
	F0   bool
	F3   bool
	F11  _dafny.Int
	_f1  bool
	_f2  _dafny.Int
	_f4  _dafny.Sequence
	_f5  bool
	_f6  _dafny.Int
	_f7  bool
	_f8  _dafny.Int
	_f9  _dafny.Int
	_f10 _dafny.Int
	_f12 bool
	_f13 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F3 = false
	_this.F11 = _dafny.Zero
	_this._f1 = false
	_this._f2 = _dafny.Zero
	_this._f4 = _dafny.EmptySeq
	_this._f5 = false
	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f12 = false
	_this._f13 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 _dafny.Int, f3 bool, f4 _dafny.Sequence, f5 bool, f6 _dafny.Int, f7 bool, f8 _dafny.Int, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Int, f12 bool, f13 _dafny.Sequence) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
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
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Sequence {
	{
		return _this._f13
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f14 _dafny.Int
	_f16 _dafny.Sequence
	_f15 bool
	_f17 bool
	_f18 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f14 = _dafny.Zero
	_this._f16 = _dafny.EmptySeq
	_this._f15 = false
	_this._f17 = false
	_this._f18 = false
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

func (_this *C0) F14() _dafny.Int {
	return _this._f14
}
func (_this *C0) F14_set_(value _dafny.Int) {
	_this._f14 = value
}
func (_this *C0) F16() _dafny.Sequence {
	return _this._f16
}
func (_this *C0) F16_set_(value _dafny.Sequence) {
	_this._f16 = value
}
func (_this *C0) F15() bool {
	return _this._f15
}
func (_this *C0) Ctor__(f17 bool, f18 bool, f16 _dafny.Sequence, f14 _dafny.Int, f15 bool) {
	{
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f16 = f16
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C0) Fm0(p0 bool, globalState *GlobalState) bool {
	{
		return !(!((Companion_D0_.Create_DC0_((_this).F18())).Dtor_cf0()))
	}
}
func (_this *C0) Fm1(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F14()
	}
}
func (_this *C0) Fm2(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((Companion_D1_.Create_DC3_(_this.F16())).Dtor_cf4(), (Companion_Default___.SafeIndex((Companion_D1_.Create_DC4_((func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(404), _dafny.IntOfInt64(615))); ; {
				_compr_11, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _368_v0 _dafny.Int
				_368_v0 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(404)).Cmp(_368_v0) <= 0) && ((_368_v0).Cmp(_dafny.IntOfInt64(615)) < 0) {
					_coll11.Add((_368_v0).Plus(_this.F14()), _this.F14())
				}
			}
			return _coll11.ToMap()
		}()).Cardinality(), _this.F14())).Dtor_cf5(), _dafny.IntOfUint32(((Companion_D1_.Create_DC3_(_this.F16())).Dtor_cf4()).Cardinality()))).Uint32(), (_this.F16()).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_this.F16()).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality())
	}
}
func (_this *C0) Fm3(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf(_this.F14(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(523))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_369_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F18(), (_this).F18(), (_this).F15()), _dafny.SeqOf((_this).F15()))).Cardinality())))
	}
}
func (_this *C0) F17() bool {
	{
		return _this._f17
	}
}
func (_this *C0) F18() bool {
	{
		return _this._f18
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f14 _dafny.Int
	_f15 bool
	_f20 _dafny.Sequence
	_f21 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f14 = _dafny.Zero
	_this._f15 = false
	_this._f20 = _dafny.EmptySeq
	_this._f21 = _dafny.EmptySeq
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

func (_this *C1) F14() _dafny.Int {
	return _this._f14
}
func (_this *C1) F14_set_(value _dafny.Int) {
	_this._f14 = value
}
func (_this *C1) F15() bool {
	return _this._f15
}
func (_this *C1) Ctor__(f20 _dafny.Sequence, f21 _dafny.Sequence, f14 _dafny.Int, f15 bool) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C1) Fm0(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F15()
	}
}
func (_this *C1) Fm1(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(244), _this.F14())).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfUint32(((_this).F20()).Cardinality())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F14())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F14())))).Cardinality())
	}
}
func (_this *C1) Fm8(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		var _source6 D0 = Companion_D0_.Create_DC2_(_this.F14(), _dafny.MultiSetOf((_this).F15()))
		_ = _source6
		if _source6.Is_DC1() {
			var _370___mcc_h0 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
			_ = _370___mcc_h0
			var _371_cf1 _dafny.Int = _370___mcc_h0
			_ = _371_cf1
			return _dafny.UnicodeSeqOfUtf8Bytes("uuycpn")
		} else if _source6.Is_DC2() {
			var _372___mcc_h1 _dafny.Int = _source6.Get_().(D0_DC2).Cf2
			_ = _372___mcc_h1
			var _373___mcc_h2 _dafny.MultiSet = _source6.Get_().(D0_DC2).Cf3
			_ = _373___mcc_h2
			var _374_cf3 _dafny.MultiSet = _373___mcc_h2
			_ = _374_cf3
			var _375_cf2 _dafny.Int = _372___mcc_h1
			_ = _375_cf2
			return (_this).F20()
		} else {
			var _376___mcc_h3 bool = _source6.Get_().(D0_DC0).Cf0
			_ = _376___mcc_h3
			var _377_cf0 bool = _376___mcc_h3
			_ = _377_cf0
			return (_this).F20()
		}
	}
}
func (_this *C1) F20() _dafny.Sequence {
	{
		return _this._f20
	}
}
func (_this *C1) F21() _dafny.Sequence {
	{
		return _this._f21
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f14 _dafny.Int
	_f15 bool
	_f19 D2
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f14 = _dafny.Zero
	_this._f15 = false
	_this._f19 = Companion_D2_.Default()
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

func (_this *C2) F14() _dafny.Int {
	return _this._f14
}
func (_this *C2) F14_set_(value _dafny.Int) {
	_this._f14 = value
}
func (_this *C2) F15() bool {
	return _this._f15
}
func (_this *C2) Ctor__(f19 D2, f14 _dafny.Int, f15 bool) {
	{
		(_this)._f19 = f19
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C2) Fm0(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F15()
	}
}
func (_this *C2) Fm1(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(292)
	}
}
func (_this *C2) Fm4(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('m')
	}
}
func (_this *C2) Fm5(p0 D1, globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf((_this).F15())).Union((_dafny.SetOf((_this).F15())).Difference(_dafny.SetOf((_this).F15())))
	}
}
func (_this *C2) M1(globalState *GlobalState) (_dafny.Sequence, D1, bool, _dafny.Int) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D1 = Companion_D1_.Default()
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _378_v0 _dafny.Array
		_ = _378_v0
		var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw66
		_378_v0 = _nw66
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
		_ = _index46
		(_378_v0).ArraySet1((_this).F15(), (_index46).Int())
		var _379_v1 _dafny.Array
		_ = _379_v1
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_11
		var _nw67 _dafny.Array
		_ = _nw67
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw67 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Sequence = func(_380_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eafbti"), _dafny.UnicodeSeqOfUtf8Bytes("i"))
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw67 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw67).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw67).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_379_v1 = _nw67
		var _381_v2 _dafny.Sequence
		_ = _381_v2
		_381_v2 = _dafny.UnicodeSeqOfUtf8Bytes("kaueyjj")
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))
		_ = _index47
		(_379_v1).ArraySet1(_381_v2, (_index47).Int())
		var _382_v3 _dafny.CodePoint
		_ = _382_v3
		_382_v3 = _dafny.CodePoint('q')
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
		_ = _index48
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))
		_ = _index49
		var _rhs41 bool = (_this).F15()
		_ = _rhs41
		var _rhs42 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_381_v2, _381_v2), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F14()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_381_v2, _381_v2)).Cardinality()))).Uint32(), _382_v3)
		_ = _rhs42
		var _lhs36 _dafny.Array = _378_v0
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
		_ = _lhs37
		var _lhs38 _dafny.Array = _379_v1
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))
		_ = _lhs39
		(_lhs36).ArraySet1(_rhs41, (_lhs37).Int())
		(_lhs38).ArraySet1(_rhs42, (_lhs39).Int())
		if (_dafny.MultiSetOf(_this.F14())).Contains(_this.F14()) {
			(globalState).F11 = _this.F14()
			(globalState).F3 = ((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)) == ((_this).F15())
			var _383_v4 _dafny.Map
			_ = _383_v4
			_383_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this.F14()).Minus(_this.F14()))
			var _384_v5 _dafny.MultiSet
			_ = _384_v5
			_384_v5 = _dafny.MultiSetOf((_this).F15())
			_383_v4 = (_383_v4).Update((_dafny.IntOfUint32((_381_v2).Cardinality())).Cmp((_384_v5).Cardinality()) > 0, _this.F14())
			var _385_v6 _dafny.Map
			_ = _385_v6
			_385_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), !((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)) || ((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)))
			r2 = (func() bool {
				if (_385_v6).Contains((_383_v4).Equals(_383_v4)) {
					return (_385_v6).Get((_383_v4).Equals(_383_v4)).(bool)
				}
				return (_this).Fm0((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), globalState)
			})()
			var _386_v7 D0
			_ = _386_v7
			_386_v7 = Companion_D0_.Create_DC0_(true)
			_386_v7 = Companion_Default___.Fm6(false, globalState)
		} else {
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))
			_ = _index50
			(_379_v1).ArraySet1(_381_v2, (_index50).Int())
			var _387_v8 _dafny.Map
			_ = _387_v8
			_387_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_382_v3, (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence))
			_387_v8 = (_387_v8).Update(_dafny.CodePoint('r'), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence))
			var _388_v9 _dafny.Sequence
			_ = _388_v9
			_388_v9 = _dafny.SeqOf((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_this).F15())
			var _389_v10 T0
			_ = _389_v10
			var _nw68 *C1 = New_C1_()
			_ = _nw68
			_nw68.Ctor__(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(823))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_390_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), (Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(823))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_391_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))).Cardinality()))).Uint32(), _382_v3), _dafny.Companion_Sequence_.Update(_388_v9, (Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_388_v9).Cardinality()))).Uint32(), (_this).F15()), _this.F14(), (_this).F15())
			_389_v10 = _nw68
			var _392_v11 _dafny.Sequence
			_ = _392_v11
			_392_v11 = _dafny.SeqOf(_389_v10, _389_v10)
			var _393_v12 _dafny.Map
			_ = _393_v12
			_393_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(376), (_this.F14()).Times(_dafny.IntOfUint32(((Companion_Default___.Fm7((_this).F15(), (_this).F15(), _dafny.IntOfUint32((_392_v11).Cardinality()), globalState)).Dtor_cf15()).Cardinality())))
			_393_v12 = (_393_v12).Update(_this.F14(), _389_v10.F14())
			(globalState).F0 = (_this).F15()
			(globalState).F11 = _dafny.IntOfInt64(727)
		}
		var _hi1 _dafny.Int = (_this.F14()).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-936)))
		_ = _hi1
		for _394_i2 := _this.F14(); _394_i2.Cmp(_hi1) < 0; _394_i2 = _394_i2.Plus(_dafny.One) {
			var _395_v13 _dafny.Set
			_ = _395_v13
			_395_v13 = _dafny.SetOf(_382_v3)
			var _396_v15 _dafny.Map
			_ = _396_v15
			_396_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(-98)), func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter16 := _dafny.Iterate(((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)).Elements()); ; {
					_compr_12, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _397_v14 _dafny.CodePoint
					_397_v14 = interface{}(_compr_12).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence), _397_v14) {
						_coll12.Add(_397_v14)
					}
				}
				return _coll12.ToSet()
			}())
			if (_395_v13).Equals((func() _dafny.Set {
				if (_396_v15).Contains(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_382_v3, _394_i2)).Cardinality())) {
					return (_396_v15).Get(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_382_v3, _394_i2)).Cardinality())).(_dafny.Set)
				}
				return _395_v13
			})()) {
				var _398_v16 *C0
				_ = _398_v16
				var _nw69 *C0 = New_C0_()
				_ = _nw69
				_nw69.Ctor__(false, (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence), _dafny.IntOfUint32((_381_v2).Cardinality()), !((_this).F15()))
				_398_v16 = _nw69
				var _399_v17 _dafny.Array
				_ = _399_v17
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_12
				var _nw70 _dafny.Array
				_ = _nw70
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw70 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) D3 = func(_400_i3 _dafny.Int) D3 {
						return Companion_D3_.Create_DC9_(_this.F14())
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw70 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw70).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw70).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_399_v17 = _nw70
				var _401_v18 _dafny.Array
				_ = _401_v18
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_13
				var _nw71 _dafny.Array
				_ = _nw71
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw71 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Int = (func(_402_v16 *C0) func(_dafny.Int) _dafny.Int {
						return func(_403_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_403_i4, (_dafny.MultiSetOf((_402_v16).F17())).Cardinality())
						}
					})(_398_v16)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw71 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw71).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw71).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_401_v18 = _nw71
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_401_v18), 0))
				_ = _index51
				(_401_v18).ArraySet1(_this.F14(), (_index51).Int())
				var _404_v19 _dafny.Sequence
				_ = _404_v19
				_404_v19 = _dafny.SeqOf(_399_v17)
				var _405_v20 _dafny.Sequence
				_ = _405_v20
				_405_v20 = _dafny.SeqOf((func() _dafny.Array {
					if (_398_v16).F18() {
						return _399_v17
					}
					return (_404_v19).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_404_v19).Cardinality()))).Uint32()).(_dafny.Array)
				})())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_401_v18), 0))
				_ = _index52
				var _rhs43 bool = (_this).F15()
				_ = _rhs43
				var _rhs44 *C0 = _398_v16
				_ = _rhs44
				var _rhs45 _dafny.Array = (_405_v20).Select((Companion_Default___.SafeIndex(_394_i2, _dafny.IntOfUint32((_405_v20).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs45
				var _rhs46 _dafny.Int = _dafny.IntOfInt64(474)
				_ = _rhs46
				var _lhs40 _dafny.Array = _401_v18
				_ = _lhs40
				var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_401_v18), 0))
				_ = _lhs41
				r2 = _rhs43
				_398_v16 = _rhs44
				_399_v17 = _rhs45
				(_lhs40).ArraySet1(_rhs46, (_lhs41).Int())
				var _406_v21 _dafny.Sequence
				_ = _406_v21
				_406_v21 = _dafny.SeqOf((_398_v16).F18())
				var _407_v22 _dafny.MultiSet
				_ = _407_v22
				_407_v22 = _dafny.MultiSetOf(_406_v21)
				var _408_v23 _dafny.Map
				_ = _408_v23
				_408_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_398_v16).F18(), (_398_v16).F17())
				var _409_v24 _dafny.Map
				_ = _409_v24
				_409_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(360))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}((func(_410_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_411_i5 _dafny.Int) _dafny.CodePoint {
						return _410_v3
					}
				})(_382_v3)))).Cardinality()), _this.F14())
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
				_ = _index53
				var _rhs47 bool = (_398_v16).F17()
				_ = _rhs47
				var _rhs48 _dafny.Int = (func() _dafny.Int {
					if (func() bool {
						if (_408_v23).Contains(true) {
							return (_408_v23).Get(true).(bool)
						}
						return (_398_v16).F17()
					})() {
						return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F14())))
					}
					return (_dafny.IntOfUint32((Companion_Default___.Fm9(_this.F14(), globalState)).Cardinality())).Plus(_394_i2)
				})()
				_ = _rhs48
				var _rhs49 _dafny.MultiSet = _407_v22
				_ = _rhs49
				var _rhs50 bool = (_406_v21).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_409_v24).Contains(_dafny.IntOfInt64(769)) {
						return (_409_v24).Get(_dafny.IntOfInt64(769)).(_dafny.Int)
					}
					return _this.F14()
				})(), _dafny.IntOfUint32((_406_v21).Cardinality()))).Uint32()).(bool)
				_ = _rhs50
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				var _lhs43 *GlobalState = globalState
				_ = _lhs43
				var _lhs44 _dafny.Array = _378_v0
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
				_ = _lhs45
				_lhs42.F0 = _rhs47
				_lhs43.F11 = _rhs48
				_407_v22 = _rhs49
				(_lhs44).ArraySet1(_rhs50, (_lhs45).Int())
				var _412_v25 _dafny.Array
				_ = _412_v25
				_412_v25 = _379_v1
				_379_v1 = (_412_v25)
				var _413_v26 _dafny.Map
				_ = _413_v26
				_413_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_398_v16).F18())), _378_v0)
				_378_v0 = (func() _dafny.Array {
					if (_413_v26).Contains((_this).F15()) {
						return (_413_v26).Get((_this).F15()).(_dafny.Array)
					}
					return _378_v0
				})()
				var _414_v27 *C0
				_ = _414_v27
				var _nw72 *C0 = New_C0_()
				_ = _nw72
				_nw72.Ctor__((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_398_v16).F17(), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence), _this.F14(), (_398_v16).F18())
				_414_v27 = _nw72
			} else {
				(globalState).F11 = _394_i2
				var _415_v28 _dafny.Sequence
				_ = _415_v28
				_415_v28 = _dafny.SeqOf(_this.F14())
				var _416_v29 _dafny.Sequence
				_ = _416_v29
				_416_v29 = _dafny.SeqOf((_this).F15())
				var _417_v30 *C1
				_ = _417_v30
				var _nw73 *C1 = New_C1_()
				_ = _nw73
				_nw73.Ctor__(_381_v2, _dafny.SeqOf((_this).Fm0((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), globalState), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)), (_394_i2).Plus((_415_v28).Select((Companion_Default___.SafeIndex(_394_i2, _dafny.IntOfUint32((_415_v28).Cardinality()))).Uint32()).(_dafny.Int)), (_416_v29).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_416_v29).Cardinality()))).Uint32()).(bool))
				_417_v30 = _nw73
				(globalState).F11 = (_dafny.Zero).Minus(_394_i2)
				(globalState).F3 = (_this).F15()
				var _418_v31 _dafny.Array
				_ = _418_v31
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_14
				var _nw74 _dafny.Array
				_ = _nw74
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw74 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.Map = (func(_419_v0 _dafny.Array, _420_v28 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
						return func(_421_i6 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_419_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_419_v0), 0))).Int()).(bool), _420_v28)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _420_v28))
						}
					})(_378_v0, _415_v28)
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw74 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw74).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw74).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_418_v31 = _nw74
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v31), 0))
				_ = _index54
				(_418_v31).ArraySet1(Companion_Default___.Fm10((_this).F15(), globalState), (_index54).Int())
				var _422_v32 _dafny.MultiSet
				_ = _422_v32
				_422_v32 = _dafny.MultiSetOf(_394_i2)
				var _423_v33 _dafny.Map
				_ = _423_v33
				_423_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_415_v28, (Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_415_v28)).Cardinality(), _dafny.IntOfUint32((_415_v28).Cardinality()))).Uint32(), _394_i2), (Companion_Default___.SafeIndex((_417_v30).Fm1(_422_v32, (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_415_v28, (Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_415_v28)).Cardinality(), _dafny.IntOfUint32((_415_v28).Cardinality()))).Uint32(), _394_i2)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_381_v2).Cardinality())))
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v31), 0))
				_ = _index55
				(_418_v31).ArraySet1((_423_v33).Update(false, _415_v28), (_index55).Int())
			}
			var _424_v34 _dafny.Map
			_ = _424_v34
			_424_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			var _425_v35 _dafny.Sequence
			_ = _425_v35
			_425_v35 = _dafny.SeqOf((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_this).F15())
			var _426_v36 *C0
			_ = _426_v36
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__((_this).Fm0((_this).Fm0((_this).F15(), globalState), globalState), !((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(101))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}(func(_427_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), _381_v2), ((_424_v34).Update((_dafny.Zero).Minus(_dafny.IntOfUint32((_425_v35).Cardinality())), (func() bool {
				if (_424_v34).Contains(_dafny.IntOfInt64(660)) {
					return (_424_v34).Get(_dafny.IntOfInt64(660)).(bool)
				}
				return (_this).Fm0((_this).F15(), globalState)
			})())).Cardinality(), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			_426_v36 = _nw75
			var _428_v37 _dafny.Sequence
			_ = _428_v37
			_428_v37 = _dafny.SeqOf(_dafny.IntOfInt64(-45), _this.F14(), _dafny.IntOfInt64(-309), _this.F14(), _this.F14())
			var _429_v38 _dafny.MultiSet
			_ = _429_v38
			_429_v38 = _dafny.MultiSetOf(_381_v2)
			var _430_v39 _dafny.MultiSet
			_ = _430_v39
			_430_v39 = _dafny.MultiSetOf(_this.F14())
			var _431_v40 D3
			_ = _431_v40
			_431_v40 = Companion_D3_.Create_DC10_(_381_v2, _378_v0, (_430_v39).IsSubsetOf((_dafny.MultiSetFromSeq(_428_v37)).Update((func() _dafny.Int {
				if (_429_v38).Contains(_dafny.UnicodeSeqOfUtf8Bytes("bxayt")) {
					return (_429_v38).Multiplicity(_dafny.UnicodeSeqOfUtf8Bytes("bxayt"))
				}
				return _dafny.IntOfUint32(((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())
			})(), Companion_Default___.Abs(_dafny.IntOfUint32(((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())))), (_dafny.Zero).Minus(_394_i2), _394_i2)
			_431_v40 = _431_v40
			_424_v34 = (_424_v34).Update(_this.F14(), (_394_i2).Cmp(_this.F14()) <= 0)
		}
		var _432_v41 _dafny.Sequence
		_ = _432_v41
		_432_v41 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence), _381_v2), (func() _dafny.Sequence {
			if true {
				return (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)
			}
			return (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)
		})(), _dafny.UnicodeSeqOfUtf8Bytes("urdc"), _dafny.Companion_Sequence_.Concatenate((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence))
		var _433_v42 _dafny.Map
		_ = _433_v42
		_433_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(828), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(471))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_434_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_435_i8 _dafny.Int) _dafny.Sequence {
				return _434_v2
			}
		})(_381_v2))))
		_432_v41 = (func() _dafny.Sequence {
			if (_433_v42).Contains(_this.F14()) {
				return (_433_v42).Get(_this.F14()).(_dafny.Sequence)
			}
			return _432_v41
		})()
		var _436_v43 _dafny.Array
		_ = _436_v43
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_15
		var _nw76 _dafny.Array
		_ = _nw76
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw76 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.Int = (func(_437_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_438_i9 _dafny.Int) _dafny.Int {
					return (_438_i9).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}((func(_439_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_440_i10 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqOf((_this).F15(), (_439_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_439_v0), 0))).Int()).(bool))).Cardinality())
						}
					})(_437_v0)))).Cardinality()))
				}
			})(_378_v0)
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw76 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw76).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw76).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_436_v43 = _nw76
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))
		_ = _index56
		(_436_v43).ArraySet1(((_dafny.Zero).Minus(_this.F14())).Plus(_this.F14()), (_index56).Int())
		var _441_v44 _dafny.Sequence
		_ = _441_v44
		_441_v44 = _dafny.SeqOf(_this.F14())
		var _442_v45 _dafny.Set
		_ = _442_v45
		_442_v45 = _dafny.SetOf((_441_v44).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_441_v44).Cardinality()))).Uint32()).(_dafny.Int))
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_436_v43), 0))
		_ = _index57
		(_436_v43).ArraySet1((_442_v45).Cardinality(), (_index57).Int())
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))
		_ = _index58
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_436_v43), 0))
		_ = _index59
		var _rhs51 _dafny.Int = _dafny.IntOfInt64(262)
		_ = _rhs51
		var _rhs52 _dafny.Int = _this.F14()
		_ = _rhs52
		var _rhs53 _dafny.Int = _this.F14()
		_ = _rhs53
		var _rhs54 _dafny.Array = _436_v43
		_ = _rhs54
		var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_441_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_443_i11 _dafny.Int) _dafny.Int {
			return _this.F14()
		})))
		_ = _rhs55
		var _lhs46 _dafny.Array = _436_v43
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))
		_ = _lhs47
		var _lhs48 *C2 = _this
		_ = _lhs48
		var _lhs49 _dafny.Array = _436_v43
		_ = _lhs49
		var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_436_v43), 0))
		_ = _lhs50
		(_lhs46).ArraySet1(_rhs51, (_lhs47).Int())
		_lhs48.F14_set_(_rhs52)
		(_lhs49).ArraySet1(_rhs53, (_lhs50).Int())
		_436_v43 = _rhs54
		_441_v44 = _rhs55
		var _444_v46 _dafny.MultiSet
		_ = _444_v46
		_444_v46 = _dafny.MultiSetOf((_this).F15(), (_this).F15(), (_this).F15(), (_this).F15())
		var _445_v47 D0
		_ = _445_v47
		_445_v47 = Companion_D0_.Create_DC2_(_this.F14(), _444_v46)
		var _pat_let_tv13 = _436_v43
		_ = _pat_let_tv13
		var _pat_let_tv14 = _436_v43
		_ = _pat_let_tv14
		var _source7 D3 = func(_source8 D0) D3 {
			if _source8.Is_DC1() {
				var _446___mcc_h8 _dafny.Int = _source8.Get_().(D0_DC1).Cf1
				_ = _446___mcc_h8
				var _447_cf1 _dafny.Int = _446___mcc_h8
				_ = _447_cf1
				return Companion_D3_.Create_DC9_(_447_cf1)
			} else if _source8.Is_DC2() {
				var _448___mcc_h9 _dafny.Int = _source8.Get_().(D0_DC2).Cf2
				_ = _448___mcc_h9
				var _449___mcc_h10 _dafny.MultiSet = _source8.Get_().(D0_DC2).Cf3
				_ = _449___mcc_h10
				var _450_cf3 _dafny.MultiSet = _449___mcc_h10
				_ = _450_cf3
				var _451_cf2 _dafny.Int = _448___mcc_h9
				_ = _451_cf2
				return Companion_D3_.Create_DC9_(_451_cf2)
			} else {
				var _452___mcc_h11 bool = _source8.Get_().(D0_DC0).Cf0
				_ = _452___mcc_h11
				var _453_cf0 bool = _452___mcc_h11
				_ = _453_cf0
				return Companion_D3_.Create_DC9_((_pat_let_tv14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_pat_let_tv13), 0))).Int()).(_dafny.Int))
			}
		}(_445_v47)
		_ = _source7
		if _source7.Is_DC9() {
			var _454___mcc_h0 _dafny.Int = _source7.Get_().(D3_DC9).Cf16
			_ = _454___mcc_h0
			var _455_cf16 _dafny.Int = _454___mcc_h0
			_ = _455_cf16
			var _456_v48 _dafny.Set
			_ = _456_v48
			_456_v48 = _dafny.SetOf((_this).F15())
			var _457_v49 _dafny.MultiSet
			_ = _457_v49
			_457_v49 = _dafny.MultiSetOf(_456_v48, _456_v48, _456_v48)
			var _458_v50 _dafny.Map
			_ = _458_v50
			_458_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _457_v49)
			(_this).F14_set_(_dafny.IntOfUint32((Companion_Default___.Fm11((func() _dafny.MultiSet {
				if (_458_v50).Contains((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)) {
					return (_458_v50).Get((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)).(_dafny.MultiSet)
				}
				return _457_v49
			})(), globalState)).Cardinality()))
			var _459_v51 _dafny.MultiSet
			_ = _459_v51
			_459_v51 = _dafny.MultiSetOf(_dafny.IntOfInt64(170))
			var _460_v52 _dafny.Map
			_ = _460_v52
			_460_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F15())
			var _461_v53 _dafny.Map
			_ = _461_v53
			_461_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_459_v51, _460_v52)
			_461_v53 = (_461_v53).Update((_459_v51).Update(_this.F14(), Companion_Default___.Abs((_dafny.Zero).Minus((func() _dafny.Int {
				if (_444_v46).Contains((_this).F15()) {
					return (_444_v46).Multiplicity((_this).F15())
				}
				return _this.F14()
			})()))), (_460_v52).Merge(_460_v52))
			var _462_v54 D1
			_ = _462_v54
			_462_v54 = Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(746))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_463_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_464_i12 _dafny.Int) _dafny.CodePoint {
					return _463_v3
				}
			})(_382_v3))))
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))
			_ = _index60
			var _rhs56 _dafny.Int = (_455_cf16).Plus(_455_cf16)
			_ = _rhs56
			var _rhs57 _dafny.Set = ((_456_v48).Difference(_456_v48)).Difference((_this).Fm5(_462_v54, globalState))
			_ = _rhs57
			var _lhs51 _dafny.Array = _436_v43
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))
			_ = _lhs52
			(_lhs51).ArraySet1(_rhs56, (_lhs52).Int())
			_456_v48 = _rhs57
			var _465_v55 _dafny.Sequence
			_ = _465_v55
			_465_v55 = _dafny.SeqOf((_this).F15(), true, (_this).F15(), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			var _466_v56 _dafny.Map
			_ = _466_v56
			_466_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int), _dafny.MultiSetFromSeq(_465_v55))
			var _467_v57 D2
			_ = _467_v57
			_467_v57 = Companion_D2_.Create_DC6_(_466_v56)
			var _468_v58 _dafny.Map
			_ = _468_v58
			_468_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_467_v57, ((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)) && ((_this).F15()))
			var _469_v59 *C0
			_ = _469_v59
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_this).F15(), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence), _dafny.IntOfInt64(670), (_this).F15())
			_469_v59 = _nw77
			var _470_v60 _dafny.Map
			_ = _470_v60
			_470_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C0 {
				if (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool) {
					return _469_v59
				}
				return _469_v59
			})(), _468_v58)
			_468_v58 = (func() _dafny.Map {
				if (_470_v60).Contains(_469_v59) {
					return (_470_v60).Get(_469_v59).(_dafny.Map)
				}
				return _468_v58
			})()
		} else if _source7.Is_DC10() {
			var _471___mcc_h1 _dafny.Sequence = _source7.Get_().(D3_DC10).Cf17
			_ = _471___mcc_h1
			var _472___mcc_h2 _dafny.Array = _source7.Get_().(D3_DC10).Cf18
			_ = _472___mcc_h2
			var _473___mcc_h3 bool = _source7.Get_().(D3_DC10).Cf19
			_ = _473___mcc_h3
			var _474___mcc_h4 _dafny.Int = _source7.Get_().(D3_DC10).Cf20
			_ = _474___mcc_h4
			var _475___mcc_h5 _dafny.Int = _source7.Get_().(D3_DC10).Cf21
			_ = _475___mcc_h5
			var _476_cf21 _dafny.Int = _475___mcc_h5
			_ = _476_cf21
			var _477_cf20 _dafny.Int = _474___mcc_h4
			_ = _477_cf20
			var _478_cf19 bool = _473___mcc_h3
			_ = _478_cf19
			var _479_cf18 _dafny.Array = _472___mcc_h2
			_ = _479_cf18
			var _480_cf17 _dafny.Sequence = _471___mcc_h1
			_ = _480_cf17
			_381_v2 = (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)
			(globalState).F3 = (_this).F15()
			var _481_v61 _dafny.Set
			_ = _481_v61
			_481_v61 = _dafny.SetOf(_479_cf18)
			var _482_v62 _dafny.Map
			_ = _482_v62
			_482_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			var _483_v63 *C0
			_ = _483_v63
			var _nw78 *C0 = New_C0_()
			_ = _nw78
			_nw78.Ctor__((_481_v61).IsDisjointFrom(_481_v61), (_this).F15(), _dafny.Companion_Sequence_.Concatenate(_381_v2, _480_cf17), _476_cf21, (func() bool {
				if (_482_v62).Contains(_476_cf21) {
					return (_482_v62).Get(_476_cf21).(bool)
				}
				return !(true)
			})())
			_483_v63 = _nw78
			var _484_v64 _dafny.Map
			_ = _484_v64
			_484_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_cf19, _this.F14())
			var _485_v65 D5
			_ = _485_v65
			_485_v65 = Companion_D5_.Create_DC13_(_484_v64)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
			_ = _index61
			var _rhs58 _dafny.Int = (_476_cf21).Minus((func() _dafny.Int {
				if (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool) {
					return (_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)
				}
				return _477_cf20
			})())
			_ = _rhs58
			var _rhs59 bool = true
			_ = _rhs59
			var _rhs60 _dafny.Int = (_483_v63).Fm2(_dafny.Companion_Sequence_.Concatenate(_381_v2, _381_v2), (_483_v63).F17(), !(_484_v64).Equals((_485_v65).Dtor_cf24()), (_dafny.Zero).Minus(_this.F14()), globalState)
			_ = _rhs60
			var _rhs61 *C0 = _483_v63
			_ = _rhs61
			var _rhs62 _dafny.Int = _477_cf20
			_ = _rhs62
			var _lhs53 _dafny.Array = _378_v0
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))
			_ = _lhs54
			var _lhs55 *C2 = _this
			_ = _lhs55
			r3 = _rhs58
			(_lhs53).ArraySet1(_rhs59, (_lhs54).Int())
			r3 = _rhs60
			_483_v63 = _rhs61
			_lhs55.F14_set_(_rhs62)
			_476_cf21 = _dafny.IntOfUint32((_480_cf17).Cardinality())
		} else if _source7.Is_DC8() {
			var _486___mcc_h6 _dafny.Sequence = _source7.Get_().(D3_DC8).Cf15
			_ = _486___mcc_h6
			var _487_cf15 _dafny.Sequence = _486___mcc_h6
			_ = _487_cf15
			var _488_v66 _dafny.Map
			_ = _488_v66
			_488_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_442_v45).Cardinality())
			_488_v66 = (_488_v66).Update(_this.F14(), _dafny.IntOfInt64(855))
			_442_v45 = (func() _dafny.Set {
				if (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool) {
					return (_442_v45).Union(_442_v45)
				}
				return Companion_Default___.Fm12(_this.F14(), (_this).F15(), (_dafny.Zero).Minus((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)), _382_v3, globalState)
			})()
			var _489_v67 _dafny.Map
			_ = _489_v67
			_489_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(935), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			var _490_v68 *C0
			_ = _490_v68
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__((func() bool {
				if (_489_v67).Contains(_dafny.IntOfInt64(239)) {
					return (_489_v67).Get(_dafny.IntOfInt64(239)).(bool)
				}
				return !((_this).F15())
			})(), (func() bool {
				if (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool) {
					return (_this).F15()
				}
				return (_this).F15()
			})(), _381_v2, (_dafny.Zero).Minus((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)), (_this).F15())
			_490_v68 = _nw79
			(globalState).F11 = (_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)
		} else {
			var _491___mcc_h7 D3 = _source7.Get_().(D3_DC11).Cf22
			_ = _491___mcc_h7
			var _492_cf22 D3 = _491___mcc_h7
			_ = _492_cf22
			var _493_v69 _dafny.Sequence
			_ = _493_v69
			_493_v69 = _dafny.SeqOf((_this).Fm0((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), globalState))
			var _494_v70 _dafny.Map
			_ = _494_v70
			_494_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _493_v69)
			(globalState).F11 = ((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_494_v70).Contains(_this.F14()) {
					return (_494_v70).Get(_this.F14()).(_dafny.Sequence)
				}
				return _dafny.SeqOf((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			})()).Cardinality()))
			var _495_v71 D0
			_ = _495_v71
			_495_v71 = Companion_D0_.Create_DC0_((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool))
			_495_v71 = _495_v71
			var _496_v72 _dafny.Map
			_ = _496_v72
			_496_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F15())
			_496_v72 = ((_496_v72).Update((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), !((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool)))).Merge(_496_v72)
			var _497_v73 _dafny.Array
			_ = _497_v73
			var _nwElement0_9 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_this).F15() {
					return _382_v3
				}
				return (_381_v2).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_381_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
			})()
			_ = _nwElement0_9
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(20))
			_ = _nw80
			(_nw80).ArraySet1CodePoint(_nwElement0_9, 0)
			(_nw80).ArraySet1CodePoint(_dafny.CodePoint('h'), 1)
			(_nw80).ArraySet1CodePoint(_dafny.CodePoint('y'), 2)
			(_nw80).ArraySet1CodePoint(_382_v3, 3)
			(_nw80).ArraySet1CodePoint(_382_v3, 4)
			(_nw80).ArraySet1CodePoint(((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32(((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), 5)
			(_nw80).ArraySet1CodePoint(_382_v3, 6)
			(_nw80).ArraySet1CodePoint(_382_v3, 7)
			(_nw80).ArraySet1CodePoint(_382_v3, 8)
			(_nw80).ArraySet1CodePoint((_381_v2).Select((Companion_Default___.SafeIndex((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_381_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
			(_nw80).ArraySet1CodePoint(_382_v3, 10)
			(_nw80).ArraySet1CodePoint(_382_v3, 11)
			(_nw80).ArraySet1CodePoint(_382_v3, 12)
			(_nw80).ArraySet1CodePoint(_382_v3, 13)
			(_nw80).ArraySet1CodePoint(_dafny.CodePoint('w'), 14)
			(_nw80).ArraySet1CodePoint(_382_v3, 15)
			(_nw80).ArraySet1CodePoint(_382_v3, 16)
			(_nw80).ArraySet1CodePoint(_382_v3, 17)
			(_nw80).ArraySet1CodePoint(_382_v3, 18)
			(_nw80).ArraySet1CodePoint(_382_v3, 19)
			_497_v73 = _nw80
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_497_v73), 0))
			_ = _index62
			(_497_v73).ArraySet1CodePoint(_382_v3, (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_497_v73), 0))
			_ = _index63
			(_497_v73).ArraySet1CodePoint(_382_v3, (_index63).Int())
		}
		r0 = (_432_v41).Select((Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_432_v41).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _498_v74 _dafny.Map
		_ = _498_v74
		_498_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int), (_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence))
		var _499_v75 _dafny.Map
		_ = _499_v75
		_499_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_this).F15())
		r1 = Companion_D1_.Create_DC5_((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int), (_dafny.SetOf((_379_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_379_v1), 0))).Int()).(_dafny.Sequence))).Cardinality(), (_378_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_378_v0), 0))).Int()).(bool), (_498_v74).Update((_499_v75).Cardinality(), _dafny.SeqOf(_382_v3, _382_v3, _dafny.CodePoint('d'))))
		r2 = ((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)).Cmp((_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)) != 0
		r3 = (_436_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_436_v43), 0))).Int()).(_dafny.Int)
		return r0, r1, r2, r3
	}
}
func (_this *C2) M2(p0 D2, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _500_v0 _dafny.Array
		_ = _500_v0
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
		_ = _nw81
		_500_v0 = _nw81
		var _501_v1 D5
		_ = _501_v1
		_501_v1 = Companion_D5_.Create_DC14_(p3, _500_v0, (_this).F15())
		var _502_v2 _dafny.Sequence
		_ = _502_v2
		_502_v2 = _dafny.SeqOf((_this).F15())
		var _503_v3 *C0
		_ = _503_v3
		var _nw82 *C0 = New_C0_()
		_ = _nw82
		_nw82.Ctor__((_this).F15(), p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_504_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _dafny.IntOfUint32((_502_v2).Cardinality()), p1)
		_503_v3 = _nw82
		var _505_v4 _dafny.MultiSet
		_ = _505_v4
		_505_v4 = _dafny.MultiSetOf(_503_v3, _503_v3)
		var _506_v5 _dafny.Map
		_ = _506_v5
		_506_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _505_v4)
		var _507_v6 _dafny.Sequence
		_ = _507_v6
		_507_v6 = _dafny.SeqOf(_506_v5)
		var _508_v7 _dafny.Sequence
		_ = _508_v7
		_508_v7 = _dafny.UnicodeSeqOfUtf8Bytes("ejencq")
		var _509_v8 T0
		_ = _509_v8
		var _nw83 *C1 = New_C1_()
		_ = _nw83
		_nw83.Ctor__(_508_v7, _502_v2, _dafny.IntOfInt64(335), (_503_v3).F18())
		_509_v8 = _nw83
		var _510_v9 _dafny.Set
		_ = _510_v9
		_510_v9 = _dafny.SetOf(_509_v8, _509_v8)
		var _511_v10 _dafny.Array
		_ = _511_v10
		var _nwElement0_10 bool = (_this).F15()
		_ = _nwElement0_10
		var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(16))
		_ = _nw84
		(_nw84).ArraySet1(_nwElement0_10, 0)
		(_nw84).ArraySet1(true, 1)
		(_nw84).ArraySet1((func() bool {
			if (_this).F15() {
				return p1
			}
			return !((_501_v1).Dtor_cf27())
		})(), 2)
		(_nw84).ArraySet1((_this).F15(), 3)
		(_nw84).ArraySet1((_this).F15(), 4)
		(_nw84).ArraySet1((_this).F15(), 5)
		(_nw84).ArraySet1((((_507_v6).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_507_v6).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Cmp(p3) == 0, 6)
		(_nw84).ArraySet1((p2).Cmp((_dafny.Zero).Minus(p3)) <= 0, 7)
		(_nw84).ArraySet1((_503_v3).F17(), 8)
		(_nw84).ArraySet1((_502_v2).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_502_v2).Cardinality()))).Uint32()).(bool), 9)
		(_nw84).ArraySet1(!((_503_v3).F18()), 10)
		(_nw84).ArraySet1((_this).F15(), 11)
		(_nw84).ArraySet1((_510_v9).IsSubsetOf(_dafny.SetOf(_509_v8, _509_v8, _509_v8, _509_v8)), 12)
		(_nw84).ArraySet1((_503_v3).F17(), 13)
		(_nw84).ArraySet1(p1, 14)
		(_nw84).ArraySet1(p1, 15)
		_511_v10 = _nw84
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
		_ = _index64
		(_511_v10).ArraySet1((_509_v8).F15(), (_index64).Int())
		var _512_v11 _dafny.Array
		_ = _512_v11
		var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
		_ = _nw85
		_512_v11 = _nw85
		var _513_v12 _dafny.MultiSet
		_ = _513_v12
		_513_v12 = _dafny.MultiSetOf(_this.F14(), p2)
		var _514_v13 _dafny.Map
		_ = _514_v13
		_514_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_512_v11, (_513_v12).Cardinality())
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
		_ = _index65
		var _rhs63 _dafny.Int = p3
		_ = _rhs63
		var _rhs64 _dafny.Int = (func() _dafny.Int {
			if (_514_v13).Contains(_512_v11) {
				return (_514_v13).Get(_512_v11).(_dafny.Int)
			}
			return p2
		})()
		_ = _rhs64
		var _rhs65 bool = false
		_ = _rhs65
		var _lhs56 _dafny.Array = _511_v10
		_ = _lhs56
		var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
		_ = _lhs57
		r0 = _rhs63
		r1 = _rhs64
		(_lhs56).ArraySet1(_rhs65, (_lhs57).Int())
		var _hi2 _dafny.Int = _dafny.IntOfUint32((_508_v7).Cardinality())
		_ = _hi2
		for _515_i1 := (Companion_Default___.Fm13(p3, globalState)).Cardinality(); _515_i1.Cmp(_hi2) < 0; _515_i1 = _515_i1.Plus(_dafny.One) {
			var _516_v14 _dafny.Array
			_ = _516_v14
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
			_ = _nw86
			_516_v14 = _nw86
			var _517_v15 _dafny.Array
			_ = _517_v15
			var _nwElement0_11 _dafny.Array = _516_v14
			_ = _nwElement0_11
			var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(8))
			_ = _nw87
			(_nw87).ArraySet1(_nwElement0_11, 0)
			(_nw87).ArraySet1(_516_v14, 1)
			(_nw87).ArraySet1(_516_v14, 2)
			(_nw87).ArraySet1(_516_v14, 3)
			(_nw87).ArraySet1(_516_v14, 4)
			(_nw87).ArraySet1(_516_v14, 5)
			(_nw87).ArraySet1(_516_v14, 6)
			(_nw87).ArraySet1(_516_v14, 7)
			_517_v15 = _nw87
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_517_v15), 0))
			_ = _index66
			(_517_v15).ArraySet1(_516_v14, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_517_v15), 0))
			_ = _index67
			(_517_v15).ArraySet1(_516_v14, (_index67).Int())
			var _518_v16 _dafny.Set
			_ = _518_v16
			_518_v16 = _dafny.SetOf((_503_v3).F18(), (_503_v3).F18())
			var _519_v17 *C0
			_ = _519_v17
			var _nw88 *C0 = New_C0_()
			_ = _nw88
			_nw88.Ctor__(p1, p1, _508_v7, (_513_v12).Cardinality(), (_dafny.SetOf((_503_v3).F18())).IsSubsetOf(_518_v16))
			_519_v17 = _nw88
			_508_v7 = _508_v7
			(globalState).F0 = p1
		}
		(globalState).F11 = p2
		var _520_v18 D3
		_ = _520_v18
		_520_v18 = Companion_D3_.Create_DC11_(Companion_D3_.Create_DC10_(_508_v7, _511_v10, (_503_v3).F18(), _dafny.IntOfInt64(502), _509_v8.F14()))
		var _521_v19 _dafny.Set
		_ = _521_v19
		_521_v19 = _dafny.SetOf(_520_v18)
		var _522_v20 _dafny.Map
		_ = _522_v20
		_522_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_521_v19, (_dafny.Zero).Minus(p2))
		var _523_v21 _dafny.Map
		_ = _523_v21
		_523_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Times((_dafny.Zero).Minus(p2)), !(_522_v20).Equals(_522_v20))
		var _524_v22 _dafny.Map
		_ = _524_v22
		_524_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-673))
		var _525_v23 _dafny.MultiSet
		_ = _525_v23
		_525_v23 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool), _509_v8.F14()), _524_v22)
		_523_v21 = (_523_v21).Update((_525_v23).Cardinality(), (_503_v3).F17())
		var _526_v24 D0
		_ = _526_v24
		_526_v24 = Companion_D0_.Create_DC1_((_523_v21).Cardinality())
		var _source9 D0 = _526_v24
		_ = _source9
		if _source9.Is_DC1() {
			var _527___mcc_h0 _dafny.Int = _source9.Get_().(D0_DC1).Cf1
			_ = _527___mcc_h0
			var _528_cf1 _dafny.Int = _527___mcc_h0
			_ = _528_cf1
			var _529_v25 _dafny.Set
			_ = _529_v25
			_529_v25 = _dafny.SetOf((_this).F15())
			var _530_v26 _dafny.Map
			_ = _530_v26
			_530_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_529_v25, _509_v8.F14())
			var _rhs66 bool = (_528_cf1).Cmp(_this.F14()) >= 0
			_ = _rhs66
			var _rhs67 bool = (_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool)
			_ = _rhs67
			var _rhs68 _dafny.Int = (p3).Plus(p2)
			_ = _rhs68
			var _rhs69 _dafny.Map = (_530_v26).Merge(_530_v26)
			_ = _rhs69
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			_lhs58.F0 = _rhs66
			_lhs59.F0 = _rhs67
			_lhs60.F11 = _rhs68
			_530_v26 = _rhs69
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
			_ = _index68
			(_511_v10).ArraySet1(p1, (_index68).Int())
			_524_v22 = (_524_v22).Update(!((_503_v3).F17()), p3)
			var _531_v27 *C0
			_ = _531_v27
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__((func() bool {
				if (_523_v21).Contains(_528_cf1) {
					return (_523_v21).Get(_528_cf1).(bool)
				}
				return (_509_v8).F15()
			})(), false, _508_v7, _dafny.IntOfUint32((_508_v7).Cardinality()), !(!((_503_v3).F17())))
			_531_v27 = _nw89
		} else if _source9.Is_DC2() {
			var _532___mcc_h1 _dafny.Int = _source9.Get_().(D0_DC2).Cf2
			_ = _532___mcc_h1
			var _533___mcc_h2 _dafny.MultiSet = _source9.Get_().(D0_DC2).Cf3
			_ = _533___mcc_h2
			var _534_cf3 _dafny.MultiSet = _533___mcc_h2
			_ = _534_cf3
			var _535_cf2 _dafny.Int = _532___mcc_h1
			_ = _535_cf2
			r1 = _dafny.IntOfInt64(-780)
			if (_535_cf2).Cmp((_dafny.IntOfInt64(219)).Minus(_509_v8.F14())) != 0 {
				(globalState).F11 = ((_dafny.IntOfUint32((_502_v2).Cardinality())).Minus(_this.F14())).Times(p3)
				var _536_v28 _dafny.Array
				_ = _536_v28
				var _nwElement0_12 _dafny.Sequence = _508_v7
				_ = _nwElement0_12
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(15))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_12, 0)
				(_nw90).ArraySet1(_508_v7, 1)
				(_nw90).ArraySet1(_508_v7, 2)
				(_nw90).ArraySet1(_508_v7, 3)
				(_nw90).ArraySet1(_508_v7, 4)
				(_nw90).ArraySet1(_508_v7, 5)
				(_nw90).ArraySet1(_508_v7, 6)
				(_nw90).ArraySet1(_508_v7, 7)
				(_nw90).ArraySet1(_508_v7, 8)
				(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_508_v7, Companion_Default___.Fm9(p2, globalState)), 9)
				(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cwdgami"), _508_v7), 10)
				(_nw90).ArraySet1(_508_v7, 11)
				(_nw90).ArraySet1(_508_v7, 12)
				(_nw90).ArraySet1(_508_v7, 13)
				(_nw90).ArraySet1((func() _dafny.Sequence {
					if false {
						return _dafny.UnicodeSeqOfUtf8Bytes("gamojfa")
					}
					return _508_v7
				})(), 14)
				_536_v28 = _nw90
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_536_v28), 0))
				_ = _index69
				(_536_v28).ArraySet1((func() _dafny.Sequence {
					if (_503_v3).F17() {
						return _508_v7
					}
					return _508_v7
				})(), (_index69).Int())
				var _537_v30 _dafny.Array
				_ = _537_v30
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_16
				var _nw91 _dafny.Array
				_ = _nw91
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw91 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Set = (func(_538_v3 *C0, _539_v8 T0, _540_p3 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_541_i2 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf(func() _dafny.Set {
								var _coll13 = _dafny.NewBuilder()
								_ = _coll13
								for _iter17 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D2_.Create_DC7_((_538_v3).F17(), (_539_v8).F15(), _540_p3), (_this).F19())).Elements()); ; {
									_compr_13, _ok17 := _iter17()
									if !_ok17 {
										break
									}
									var _542_v29 D2
									_542_v29 = interface{}(_compr_13).(D2)
									if (_dafny.MultiSetOf(Companion_D2_.Create_DC7_((_538_v3).F17(), (_539_v8).F15(), _540_p3), (_this).F19())).Contains(_542_v29) {
										_coll13.Add(_542_v29)
									}
								}
								return _coll13.ToSet()
							}())).Difference(_dafny.SetOf(_dafny.SetOf((_this).F19())))
						}
					})(_503_v3, _509_v8, p3)
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
				_537_v30 = _nw91
				var _543_v31 _dafny.Set
				_ = _543_v31
				_543_v31 = _dafny.SetOf(_dafny.SetOf(Companion_D2_.Create_DC7_(p1, (_503_v3).F17(), (_dafny.Zero).Minus(_509_v8.F14())), (_this).F19()))
				var _544_v32 D6
				_ = _544_v32
				_544_v32 = Companion_D6_.Create_DC16_(_543_v31)
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_537_v30), 0))
				_ = _index70
				(_537_v30).ArraySet1((_544_v32).Dtor_cf29(), (_index70).Int())
				var _545_v33 _dafny.Map
				_ = _545_v33
				_545_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_509_v8.F14()), _511_v10)
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_536_v28), 0))
				_ = _index71
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_537_v30), 0))
				_ = _index72
				var _rhs70 _dafny.Int = (_545_v33).Cardinality()
				_ = _rhs70
				var _rhs71 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("kykbqtx")
				_ = _rhs71
				var _rhs72 _dafny.Set = (_543_v31).Union(_543_v31)
				_ = _rhs72
				var _rhs73 _dafny.MultiSet = _513_v12
				_ = _rhs73
				var _lhs61 *GlobalState = globalState
				_ = _lhs61
				var _lhs62 _dafny.Array = _536_v28
				_ = _lhs62
				var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_536_v28), 0))
				_ = _lhs63
				var _lhs64 _dafny.Array = _537_v30
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_537_v30), 0))
				_ = _lhs65
				_lhs61.F11 = _rhs70
				(_lhs62).ArraySet1(_rhs71, (_lhs63).Int())
				(_lhs64).ArraySet1(_rhs72, (_lhs65).Int())
				_513_v12 = _rhs73
				var _546_v34 _dafny.Array
				_ = _546_v34
				var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(23))
				_ = _nw92
				_546_v34 = _nw92
				var _547_v35 _dafny.MultiSet
				_ = _547_v35
				_547_v35 = _dafny.MultiSetOf(_546_v34, _546_v34)
				var _548_v36 _dafny.Map
				_ = _548_v36
				_548_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), _509_v8.F14())
				var _549_v37 _dafny.Map
				_ = _549_v37
				_549_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _548_v36)
				(_509_v8).F14_set_((func() _dafny.Int {
					if (_547_v35).Contains(_546_v34) {
						return (_547_v35).Multiplicity(_546_v34)
					}
					return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_503_v3).F18(), false)).Cardinality()), (_549_v37).Cardinality())
				})())
				var _550_v38 _dafny.Set
				_ = _550_v38
				_550_v38 = _dafny.SetOf((_dafny.Zero).Minus(_509_v8.F14()))
				(globalState).F0 = (_550_v38).IsProperSubsetOf((func() _dafny.Set {
					if (_503_v3).F18() {
						return _550_v38
					}
					return _550_v38
				})())
				var _551_v39 D1
				_ = _551_v39
				_551_v39 = Companion_D1_.Create_DC3_(_508_v7)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_536_v28), 0))
				_ = _index73
				(_536_v28).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_536_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_536_v28), 0))).Int()).(_dafny.Sequence), (_551_v39).Dtor_cf4()), (_index73).Int())
			} else {
				var _552_v40 *C1
				_ = _552_v40
				var _nw93 *C1 = New_C1_()
				_ = _nw93
				_nw93.Ctor__(_dafny.Companion_Sequence_.Concatenate(_508_v7, _508_v7), _502_v2, _this.F14(), false)
				_552_v40 = _nw93
				_552_v40 = _552_v40
				var _553_v41 _dafny.Array
				_ = _553_v41
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw94
				_553_v41 = _nw94
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_553_v41), 0))
				_ = _index74
				(_553_v41).ArraySet1((func() _dafny.Int {
					if (_534_cf3).Contains((_509_v8).F15()) {
						return (_534_cf3).Multiplicity((_509_v8).F15())
					}
					return p3
				})(), (_index74).Int())
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_553_v41), 0))
				_ = _index75
				(_553_v41).ArraySet1((_dafny.Zero).Minus(_this.F14()), (_index75).Int())
				var _554_v42 _dafny.Map
				_ = _554_v42
				_554_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v8.F14(), Companion_D5_.Create_DC13_((_524_v22).Update((_509_v8).F15(), _509_v8.F14())))
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
				_ = _index76
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_553_v41), 0))
				_ = _index77
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_553_v41), 0))
				_ = _index78
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
				_ = _index79
				var _rhs74 bool = (_dafny.IntOfInt64(65)).Cmp((_554_v42).Cardinality()) != 0
				_ = _rhs74
				var _rhs75 _dafny.Int = Companion_Default___.SafeModuloInt((_this.F14()).Plus(p2), Companion_Default___.SafeModuloInt(p3, _dafny.IntOfInt64(109)))
				_ = _rhs75
				var _rhs76 _dafny.Int = _509_v8.F14()
				_ = _rhs76
				var _rhs77 bool = !((_503_v3).F18()) || ((_503_v3).F18())
				_ = _rhs77
				var _rhs78 _dafny.Int = (_535_cf2).Minus((func() _dafny.Int {
					if (_503_v3).F18() {
						return (_503_v3).Fm2((_552_v40).F20(), (_503_v3).F18(), (_503_v3).F18(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_509_v8.F14())), globalState)
					}
					return (_dafny.Zero).Minus((_523_v21).Cardinality())
				})())
				_ = _rhs78
				var _lhs66 _dafny.Array = _511_v10
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
				_ = _lhs67
				var _lhs68 _dafny.Array = _553_v41
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_553_v41), 0))
				_ = _lhs69
				var _lhs70 _dafny.Array = _553_v41
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_553_v41), 0))
				_ = _lhs71
				var _lhs72 _dafny.Array = _511_v10
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
				_ = _lhs73
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				(_lhs66).ArraySet1(_rhs74, (_lhs67).Int())
				(_lhs68).ArraySet1(_rhs75, (_lhs69).Int())
				(_lhs70).ArraySet1(_rhs76, (_lhs71).Int())
				(_lhs72).ArraySet1(_rhs77, (_lhs73).Int())
				_lhs74.F11 = _rhs78
				var _555_v43 _dafny.CodePoint
				_ = _555_v43
				_555_v43 = _dafny.CodePoint('m')
				var _556_v44 _dafny.Map
				_ = _556_v44
				_556_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v8.F14(), (_552_v40).F20())
				var _557_v45 _dafny.Array
				_ = _557_v45
				var _nwElement0_13 _dafny.Sequence = (_552_v40).F20()
				_ = _nwElement0_13
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(29))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_13, 0)
				(_nw95).ArraySet1((_552_v40).F20(), 1)
				(_nw95).ArraySet1(_508_v7, 2)
				(_nw95).ArraySet1((_552_v40).F20(), 3)
				(_nw95).ArraySet1((_552_v40).F20(), 4)
				(_nw95).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_558_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_559_i3 _dafny.Int) _dafny.CodePoint {
						return _558_v43
					}
				})(_555_v43))), 5)
				(_nw95).ArraySet1((_552_v40).F20(), 6)
				(_nw95).ArraySet1((_552_v40).F20(), 7)
				(_nw95).ArraySet1(_508_v7, 8)
				(_nw95).ArraySet1((_552_v40).F20(), 9)
				(_nw95).ArraySet1(_508_v7, 10)
				(_nw95).ArraySet1(_508_v7, 11)
				(_nw95).ArraySet1((_552_v40).F20(), 12)
				(_nw95).ArraySet1((_552_v40).F20(), 13)
				(_nw95).ArraySet1((_552_v40).Fm8((_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool), globalState), 14)
				(_nw95).ArraySet1(Companion_Default___.Fm9(_509_v8.F14(), globalState), 15)
				(_nw95).ArraySet1(_508_v7, 16)
				(_nw95).ArraySet1((_552_v40).F20(), 17)
				(_nw95).ArraySet1(_508_v7, 18)
				(_nw95).ArraySet1(_508_v7, 19)
				(_nw95).ArraySet1(_508_v7, 20)
				(_nw95).ArraySet1(_508_v7, 21)
				(_nw95).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gb"), 22)
				(_nw95).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("h"), 23)
				(_nw95).ArraySet1((_552_v40).F20(), 24)
				(_nw95).ArraySet1(_508_v7, 25)
				(_nw95).ArraySet1((_552_v40).F20(), 26)
				(_nw95).ArraySet1(_dafny.Companion_Sequence_.Update(_508_v7, (Companion_Default___.SafeIndex((_556_v44).Cardinality(), _dafny.IntOfUint32((_508_v7).Cardinality()))).Uint32(), _555_v43), 27)
				(_nw95).ArraySet1((_552_v40).F20(), 28)
				_557_v45 = _nw95
				var _560_v46 _dafny.Map
				_ = _560_v46
				_560_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_555_v43, _557_v45)
				_560_v46 = (_560_v46).Update(_555_v43, _557_v45)
				(globalState).F3 = ((_dafny.Zero).Minus(p3)).Cmp(Companion_Default___.SafeDivisionInt(_509_v8.F14(), (_dafny.Zero).Minus((_524_v22).Cardinality()))) >= 0
				var _561_v47 D3
				_ = _561_v47
				_561_v47 = Companion_D3_.Create_DC9_((_this.F14()).Plus((_553_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_553_v41), 0))).Int()).(_dafny.Int)))
				_561_v47 = _561_v47
			}
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_17
			var _nw96 _dafny.Array
			_ = _nw96
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw96 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) bool = (func(_562_v10 _dafny.Array) func(_dafny.Int) bool {
					return func(_563_i4 _dafny.Int) bool {
						return (_562_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_562_v10), 0))).Int()).(bool)
					}
				})(_511_v10)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw96 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw96).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw96).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_511_v10 = _nw96
			var _564_v48 _dafny.Array
			_ = _564_v48
			var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw97
			_564_v48 = _nw97
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_512_v11), 0))
			_ = _index80
			(_512_v11).ArraySet1(_564_v48, (_index80).Int())
			var _565_v49 _dafny.CodePoint
			_ = _565_v49
			_565_v49 = _dafny.CodePoint('r')
			var _566_v50 _dafny.Array
			_ = _566_v50
			_566_v50 = _564_v48
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_512_v11), 0))
			_ = _index81
			var _rhs79 _dafny.Int = _dafny.IntOfUint32((_508_v7).Cardinality())
			_ = _rhs79
			var _rhs80 _dafny.Array = (func() _dafny.Array {
				if (_dafny.IntOfInt64(887)).Cmp(_535_cf2) != 0 {
					return (_566_v50)
				}
				return _564_v48
			})()
			_ = _rhs80
			var _rhs81 bool = (_503_v3).F18()
			_ = _rhs81
			var _rhs82 _dafny.CodePoint = (_this).Fm4(_dafny.Companion_Sequence_.Equal(_502_v2, _502_v2), globalState)
			_ = _rhs82
			var _rhs83 bool = (p3).Cmp(_509_v8.F14()) > 0
			_ = _rhs83
			var _lhs75 _dafny.Array = _512_v11
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_512_v11), 0))
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			r0 = _rhs79
			(_lhs75).ArraySet1(_rhs80, (_lhs76).Int())
			_lhs77.F0 = _rhs81
			_565_v49 = _rhs82
			_lhs78.F0 = _rhs83
		} else {
			var _567___mcc_h3 bool = _source9.Get_().(D0_DC0).Cf0
			_ = _567___mcc_h3
			var _568_cf0 bool = _567___mcc_h3
			_ = _568_cf0
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
			_ = _index82
			(_511_v10).ArraySet1(_568_cf0, (_index82).Int())
			var _569_v51 _dafny.Set
			_ = _569_v51
			_569_v51 = _dafny.SetOf((_503_v3).F18())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
			_ = _index83
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
			_ = _index84
			var _rhs84 _dafny.Set = (func() _dafny.Set {
				if (_this).F15() {
					return _569_v51
				}
				return _569_v51
			})()
			_ = _rhs84
			var _rhs85 bool = ((_this).F15()) && ((_509_v8).F15())
			_ = _rhs85
			var _rhs86 bool = p1
			_ = _rhs86
			var _lhs79 _dafny.Array = _511_v10
			_ = _lhs79
			var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
			_ = _lhs80
			var _lhs81 _dafny.Array = _511_v10
			_ = _lhs81
			var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
			_ = _lhs82
			_569_v51 = _rhs84
			(_lhs79).ArraySet1(_rhs85, (_lhs80).Int())
			(_lhs81).ArraySet1(_rhs86, (_lhs82).Int())
			var _570_v52 T1
			_ = _570_v52
			var _nw98 *C0 = New_C0_()
			_ = _nw98
			_nw98.Ctor__((_509_v8).F15(), (_this).F15(), _508_v7, _509_v8.F14(), (_503_v3).F18())
			_570_v52 = _nw98
			var _571_v53 D6
			_ = _571_v53
			_571_v53 = Companion_D6_.Create_DC17_((_509_v8).F15(), _570_v52)
			_571_v53 = _571_v53
			var _572_v54 _dafny.CodePoint
			_ = _572_v54
			_572_v54 = _dafny.CodePoint('i')
			var _573_v55 *C1
			_ = _573_v55
			var _nw99 *C1 = New_C1_()
			_ = _nw99
			_nw99.Ctor__(_dafny.Companion_Sequence_.Update(_508_v7, (Companion_Default___.SafeIndex(_509_v8.F14(), _dafny.IntOfUint32((_508_v7).Cardinality()))).Uint32(), _572_v54), _502_v2, (_524_v22).Cardinality(), (_503_v3).F18())
			_573_v55 = _nw99
			var _574_v56 _dafny.Set
			_ = _574_v56
			_574_v56 = _dafny.SetOf(_570_v52)
			var _rhs87 _dafny.Array = _500_v0
			_ = _rhs87
			var _rhs88 _dafny.Int = (_574_v56).Cardinality()
			_ = _rhs88
			var _rhs89 _dafny.Int = _dafny.IntOfUint32(((_573_v55).F21()).Cardinality())
			_ = _rhs89
			var _rhs90 _dafny.Sequence = (func() _dafny.Sequence {
				if (_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool) {
					return (_573_v55).F21()
				}
				return (_573_v55).F21()
			})()
			_ = _rhs90
			var _rhs91 *C1 = _573_v55
			_ = _rhs91
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			var _lhs84 *GlobalState = globalState
			_ = _lhs84
			_500_v0 = _rhs87
			_lhs83.F11 = _rhs88
			_lhs84.F11 = _rhs89
			_502_v2 = _rhs90
			_573_v55 = _rhs91
		}
		var _575_v57 _dafny.Sequence
		_ = _575_v57
		_575_v57 = _dafny.SeqOf(_509_v8.F14())
		var _576_v58 _dafny.Sequence
		_ = _576_v58
		_576_v58 = _dafny.SeqOf(_575_v57, _575_v57, _575_v57)
		var _577_v59 _dafny.Set
		_ = _577_v59
		_577_v59 = _dafny.SetOf((_dafny.Zero).Minus((_this).Fm1(_513_v12, (_this).F15(), globalState)))
		var _578_v60 _dafny.Map
		_ = _578_v60
		_578_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf((_576_v58).Select((Companion_Default___.SafeIndex(_509_v8.F14(), _dafny.IntOfUint32((_576_v58).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_575_v57, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_575_v57).Cardinality()))).Uint32(), _509_v8.F14()), (Companion_Default___.SafeIndex(_this.F14(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_575_v57, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_575_v57).Cardinality()))).Uint32(), _509_v8.F14())).Cardinality()))).Uint32(), p3)), ((_577_v59).Cardinality()).Cmp(_dafny.IntOfInt64(-239)) > 0)
		if (func() bool {
			if (_578_v60).Contains(!(false)) {
				return (_578_v60).Get(!(false)).(bool)
			}
			return (_509_v8).F15()
		})() {
			var _579_v61 _dafny.Map
			_ = _579_v61
			_579_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_523_v21, _511_v10)
			Companion_Default___.M0(_579_v61, globalState)
			var _580_v62 _dafny.Array
			_ = _580_v62
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_18
			var _nw100 _dafny.Array
			_ = _nw100
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw100 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) _dafny.Int = (func(_581_v8 T0) func(_dafny.Int) _dafny.Int {
					return func(_582_i5 _dafny.Int) _dafny.Int {
						return (_582_i5).Times(_581_v8.F14())
					}
				})(_509_v8)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw100 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw100).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw100).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_580_v62 = _nw100
			var _583_v63 _dafny.Array
			_ = _583_v63
			_583_v63 = _580_v62
			_583_v63 = (func() _dafny.Array {
				if (_503_v3).F18() {
					return _583_v63
				}
				return _583_v63
			})()
			(globalState).F0 = !((_503_v3).F18())
			var _rhs92 _dafny.Int = _509_v8.F14()
			_ = _rhs92
			var _rhs93 _dafny.Int = p2
			_ = _rhs93
			var _lhs85 T0 = _509_v8
			_ = _lhs85
			var _lhs86 *GlobalState = globalState
			_ = _lhs86
			_lhs85.F14_set_(_rhs92)
			_lhs86.F11 = _rhs93
			var _584_v64 _dafny.Set
			_ = _584_v64
			_584_v64 = _dafny.SetOf((_503_v3).F18(), p1, (_503_v3).F18())
			(globalState).F0 = (_584_v64).IsProperSubsetOf(_584_v64)
		} else {
			if !((_503_v3).F17()) {
				var _585_v65 _dafny.Array
				_ = _585_v65
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_19
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Int = (func(_586_v3 *C0, _587_v8 T0, _588_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_589_i6 _dafny.Int) _dafny.Int {
							return (_589_i6).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_586_v3).F17(), (_586_v3).F17(), (_587_v8).F15(), (_586_v3).F18(), (_587_v8).F15()), _588_v2)).Cardinality()))
						}
					})(_503_v3, _509_v8, _502_v2)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw101 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw101).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw101).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_585_v65 = _nw101
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_585_v65), 0))
				_ = _index85
				(_585_v65).ArraySet1(_509_v8.F14(), (_index85).Int())
				var _590_v66 _dafny.Sequence
				_ = _590_v66
				_590_v66 = _dafny.SeqOf(_511_v10)
				var _591_v67 *C1
				_ = _591_v67
				var _nw102 *C1 = New_C1_()
				_ = _nw102
				_nw102.Ctor__(_508_v7, _502_v2, (_dafny.Zero).Minus(p3), (_this).F15())
				_591_v67 = _nw102
				var _592_v68 _dafny.Map
				_ = _592_v68
				_592_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_591_v67, (_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool))
				var _593_v69 _dafny.Sequence
				_ = _593_v69
				_593_v69 = _dafny.SeqOf(_592_v68, _592_v68, _592_v68, _592_v68, _592_v68)
				var _594_v70 _dafny.Set
				_ = _594_v70
				_594_v70 = _dafny.SetOf((_503_v3).F18())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_585_v65), 0))
				_ = _index86
				var _rhs94 _dafny.Int = (_509_v8.F14()).Minus(_509_v8.F14())
				_ = _rhs94
				var _rhs95 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_590_v66, _dafny.SeqOf(_511_v10, _511_v10, _511_v10)), _590_v66)).Cardinality())
				_ = _rhs95
				var _rhs96 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(((_593_v69).Select((Companion_Default___.SafeIndex((_594_v70).Cardinality(), _dafny.IntOfUint32((_593_v69).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), Companion_Default___.SafeDivisionInt(_this.F14(), _this.F14())))))
				_ = _rhs96
				var _rhs97 bool = false
				_ = _rhs97
				var _lhs87 _dafny.Array = _585_v65
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_585_v65), 0))
				_ = _lhs88
				var _lhs89 T0 = _509_v8
				_ = _lhs89
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				(_lhs87).ArraySet1(_rhs94, (_lhs88).Int())
				r0 = _rhs95
				_lhs89.F14_set_(_rhs96)
				_lhs90.F0 = _rhs97
				var _595_v71 _dafny.Map
				_ = _595_v71
				_595_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(809)).Plus(p2), _523_v21)
				var _596_v73 _dafny.Map
				_ = _596_v73
				_596_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_509_v8.F14()), (_585_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_585_v65), 0))).Int()).(_dafny.Int))
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
				_ = _index87
				var _rhs98 _dafny.Map = (func() _dafny.Map {
					if (_595_v71).Contains((func() _dafny.Set {
						var _coll14 = _dafny.NewBuilder()
						_ = _coll14
						for _iter18 := _dafny.Iterate((_577_v59).Elements()); ; {
							_compr_14, _ok18 := _iter18()
							if !_ok18 {
								break
							}
							var _597_v72 _dafny.Int
							_597_v72 = interface{}(_compr_14).(_dafny.Int)
							if (_577_v59).Contains(_597_v72) {
								_coll14.Add(Companion_Default___.SafeDivisionInt(_597_v72, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(253))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg31 _dafny.Int) interface{} {
										return coer31(arg31)
									}
								}(func(_598_i7 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('p')
								}))).Cardinality()))).Cardinality()))
							}
						}
						return _coll14.ToSet()
					}()).Cardinality()) {
						return (_595_v71).Get((func() _dafny.Set {
							var _coll15 = _dafny.NewBuilder()
							_ = _coll15
							for _iter19 := _dafny.Iterate((_577_v59).Elements()); ; {
								_compr_15, _ok19 := _iter19()
								if !_ok19 {
									break
								}
								var _599_v72 _dafny.Int
								_599_v72 = interface{}(_compr_15).(_dafny.Int)
								if (_577_v59).Contains(_599_v72) {
									_coll15.Add(Companion_Default___.SafeDivisionInt(_599_v72, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(253))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg32 _dafny.Int) interface{} {
											return coer32(arg32)
										}
									}(func(_600_i7 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('p')
									}))).Cardinality()))).Cardinality()))
								}
							}
							return _coll15.ToSet()
						}()).Cardinality()).(_dafny.Map)
					}
					return _523_v21
				})()
				_ = _rhs98
				var _rhs99 bool = (((_596_v73).Merge(_596_v73)).Cardinality()).Cmp(_dafny.IntOfInt64(19)) >= 0
				_ = _rhs99
				var _lhs91 _dafny.Array = _511_v10
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))
				_ = _lhs92
				_523_v21 = _rhs98
				(_lhs91).ArraySet1(_rhs99, (_lhs92).Int())
				var _rhs100 _dafny.Int = _509_v8.F14()
				_ = _rhs100
				var _rhs101 _dafny.Set = _577_v59
				_ = _rhs101
				var _rhs102 bool = (_503_v3).F17()
				_ = _rhs102
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				_lhs93.F11 = _rhs100
				_577_v59 = _rhs101
				_lhs94.F3 = _rhs102
				r1 = Companion_Default___.SafeDivisionInt(p2, _this.F14())
				var _601_v74 _dafny.CodePoint
				_ = _601_v74
				_601_v74 = _dafny.CodePoint('d')
				var _602_v75 _dafny.Map
				_ = _602_v75
				_602_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(681), _601_v74)
				r1 = (Companion_Default___.Fm14((Companion_Default___.Fm15(_601_v74, p3, (_602_v75).Cardinality(), globalState)).Update((_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool), p2), _dafny.CodePoint('c'), !(!_dafny.Companion_Sequence_.Equal((_591_v67).F20(), _dafny.Companion_Sequence_.Update(_508_v7, (Companion_Default___.SafeIndex((_591_v67).Fm1(_513_v12, true, globalState), _dafny.IntOfUint32((_508_v7).Cardinality()))).Uint32(), _601_v74))), globalState)).Cardinality()
			} else {
				var _603_v76 _dafny.Map
				_ = _603_v76
				_603_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_508_v7, _508_v7), _575_v57)
				_603_v76 = _603_v76
				var _604_v77 _dafny.Sequence
				_ = _604_v77
				_604_v77 = _dafny.SeqOf(_502_v2)
				var _605_v78 *C1
				_ = _605_v78
				var _nw103 *C1 = New_C1_()
				_ = _nw103
				_nw103.Ctor__(_508_v7, (_604_v77).Select((Companion_Default___.SafeIndex((_526_v24).Dtor_cf1(), _dafny.IntOfUint32((_604_v77).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.IntOfInt64(169), true)
				_605_v78 = _nw103
				var _606_v79 _dafny.CodePoint
				_ = _606_v79
				_606_v79 = _dafny.CodePoint('h')
				var _607_v80 _dafny.Map
				_ = _607_v80
				_607_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_503_v3).F18(), _606_v79)
				var _608_v81 D6
				_ = _608_v81
				_608_v81 = Companion_D6_.Create_DC18_(_509_v8.F14(), (func() _dafny.CodePoint {
					if (_607_v80).Contains((_503_v3).F18()) {
						return (_607_v80).Get((_503_v3).F18()).(_dafny.CodePoint)
					}
					return _606_v79
				})())
				var _609_v82 D6
				_ = _609_v82
				_609_v82 = Companion_D6_.Create_DC19_(_608_v81)
				var _610_v83 D6
				_ = _610_v83
				_610_v83 = Companion_D6_.Create_DC19_(_609_v82)
				var _611_v84 _dafny.Map
				_ = _611_v84
				_611_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v78, _610_v83)
				_611_v84 = (_611_v84).Update((func() *C1 {
					if (_509_v8).F15() {
						return _605_v78
					}
					return _605_v78
				})(), _610_v83)
				var _612_v85 *C0
				_ = _612_v85
				var _nw104 *C0 = New_C0_()
				_ = _nw104
				_nw104.Ctor__((_503_v3).F17(), (_503_v3).F18(), (_605_v78).F20(), _this.F14(), (_503_v3).F17())
				_612_v85 = _nw104
				var _613_v86 _dafny.MultiSet
				_ = _613_v86
				_613_v86 = _dafny.MultiSetOf((_503_v3).F18(), (_612_v85).Fm0((_503_v3).F18(), globalState), (_509_v8).F15(), (_503_v3).F18(), (_this).F15())
				var _614_v87 _dafny.Map
				_ = _614_v87
				_614_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v8.F14(), _613_v86)
				var _615_v88 D2
				_ = _615_v88
				_615_v88 = Companion_D2_.Create_DC6_(_614_v87)
				_615_v88 = _615_v88
				var _616_v89 _dafny.Map
				_ = _616_v89
				_616_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm16((_577_v59).Cardinality(), globalState), _511_v10)
				Companion_Default___.M0(_616_v89, globalState)
			}
			(globalState).F0 = (_502_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if (_513_v12).Contains(_dafny.IntOfUint32((_508_v7).Cardinality())) {
					return (_513_v12).Multiplicity(_dafny.IntOfUint32((_508_v7).Cardinality()))
				}
				return _509_v8.F14()
			})()), _dafny.IntOfUint32((_502_v2).Cardinality()))).Uint32()).(bool)
			r1 = (_509_v8.F14()).Minus(_this.F14())
			(globalState).F0 = (_503_v3).Fm0((_511_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_511_v10), 0))).Int()).(bool), globalState)
			var _617_v90 _dafny.Sequence
			_ = _617_v90
			_617_v90 = _dafny.SeqOf(_502_v2, _502_v2, _502_v2, _dafny.SeqOf((_503_v3).F18()), _dafny.SeqOf((_509_v8).F15(), true))
			var _618_v91 *C1
			_ = _618_v91
			var _nw105 *C1 = New_C1_()
			_ = _nw105
			_nw105.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("aqvurq"), (_617_v90).Select((Companion_Default___.SafeIndex(_509_v8.F14(), _dafny.IntOfUint32((_617_v90).Cardinality()))).Uint32()).(_dafny.Sequence), _509_v8.F14(), (_503_v3).F18())
			_618_v91 = _nw105
		}
		r0 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_509_v8).Fm1(_513_v12, (_this).F15(), globalState))), p3)
		var _619_v92 _dafny.Sequence
		_ = _619_v92
		_619_v92 = _dafny.SeqOf(_dafny.SetOf(p3, p2))
		r1 = Companion_Default___.SafeModuloInt(_509_v8.F14(), (_dafny.IntOfUint32((_619_v92).Cardinality())).Minus(_dafny.IntOfInt64(-105)))
		return r0, r1
	}
}
func (_this *C2) F19() D2 {
	{
		return _this._f19
	}
}

// End of class C2
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
