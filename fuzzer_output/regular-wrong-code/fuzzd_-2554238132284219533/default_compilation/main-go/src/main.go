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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Map, globalState *GlobalState) bool {
	return (true) == (!(!((func() bool {
		if true {
			return true
		}
		return false
	})())))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	var _source0 D10 = Companion_D10_.Create_DC23_((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(11), (_dafny.MultiSetOf(_dafny.IntOfInt64(473), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mwihgjxua")).Cardinality()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-810), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))).Cardinality())))).Cardinality(), !(true), true)
	_ = _source0
	if _source0.Is_DC23() {
		var _0___mcc_h0 _dafny.Int = _source0.Get_().(D10_DC23).Cf33
		_ = _0___mcc_h0
		var _1___mcc_h1 bool = _source0.Get_().(D10_DC23).Cf34
		_ = _1___mcc_h1
		var _2___mcc_h2 bool = _source0.Get_().(D10_DC23).Cf35
		_ = _2___mcc_h2
		var _3_cf35 bool = _2___mcc_h2
		_ = _3_cf35
		var _4_cf34 bool = _1___mcc_h1
		_ = _4_cf34
		var _5_cf33 _dafny.Int = _0___mcc_h0
		_ = _5_cf33
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_5_cf33, _dafny.IntOfInt64(841))).Cardinality()), _4_cf34)).Cardinality()
	} else {
		var _6___mcc_h3 *C1 = _source0.Get_().(D10_DC22).Cf32
		_ = _6___mcc_h3
		var _7_cf32 *C1 = _6___mcc_h3
		_ = _7_cf32
		if (_7_cf32).F16() {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-846))).Cardinality())
		} else {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dhf")).Cardinality())
		}
	}
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("g")
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Map, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(933), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(200), true)).Cardinality())).Intersection(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-918), !(false))).Cardinality()))).Union(_dafny.SetOf(_dafny.IntOfInt64(-219), _dafny.IntOfUint32((_dafny.SeqOf(false, !(!(true)))).Cardinality()), _dafny.IntOfInt64(623), _dafny.IntOfInt64(881)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Int {
	return (func(_source1 D1) _dafny.MultiSet {
		if _source1.Is_DC2() {
			var _8___mcc_h0 bool = _source1.Get_().(D1_DC2).Cf2
			_ = _8___mcc_h0
			var _9_cf2 bool = _8___mcc_h0
			_ = _9_cf2
			return (_dafny.MultiSetOf(_dafny.IntOfInt64(-14), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(541))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()), _dafny.IntOfInt64(958), (func() _dafny.Map {
				var _coll0 = _dafny.NewMapBuilder()
				_ = _coll0
				for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(892), _dafny.IntOfInt64(185))); ; {
					_compr_0, _ok0 := _iter0()
					if !_ok0 {
						break
					}
					var _11_v0 _dafny.Int
					_11_v0 = interface{}(_compr_0).(_dafny.Int)
					if ((_dafny.IntOfInt64(892)).Cmp(_11_v0) <= 0) && ((_11_v0).Cmp(_dafny.IntOfInt64(185)) < 0) {
						_coll0.Add((_11_v0).Plus(_dafny.IntOfInt64(712)), _dafny.IntOfUint32((_dafny.SeqOf(_9_cf2, _9_cf2)).Cardinality()))
					}
				}
				return _coll0.ToMap()
			}()).Cardinality())).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-142))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_12_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-701)
			}))))
		} else if _source1.Is_DC3() {
			var _13___mcc_h1 bool = _source1.Get_().(D1_DC3).Cf3
			_ = _13___mcc_h1
			var _14___mcc_h2 _dafny.Int = _source1.Get_().(D1_DC3).Cf4
			_ = _14___mcc_h2
			var _15_cf4 _dafny.Int = _14___mcc_h2
			_ = _15_cf4
			var _16_cf3 bool = _13___mcc_h1
			_ = _16_cf3
			return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_17_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(944)
			})), _dafny.SeqOf(_dafny.IntOfInt64(290), _dafny.IntOfInt64(674))))
		} else {
			var _18___mcc_h3 bool = _source1.Get_().(D1_DC1).Cf1
			_ = _18___mcc_h3
			var _19_cf1 bool = _18___mcc_h3
			_ = _19_cf1
			return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_19_cf1, _19_cf1, true)).Cardinality())))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_19_cf1)).Cardinality())))
		}
	}(Companion_D1_.Create_DC3_(!(!(false)), _dafny.IntOfInt64(18)))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(756))).Uint32(), func(coer3 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_20_i0 _dafny.Int) _dafny.MultiSet {
		return (_dafny.MultiSetOf(!(!(false)), !(!(true)))).Union(_dafny.MultiSetOf(true))
	}))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Map, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC8_(true, _dafny.IntOfInt64(257), _dafny.IntOfInt64(9)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC8_(true, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC8_(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-201)))))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(false, (_dafny.IntOfInt64(633)).Times((Companion_D15_.Create_DC39_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.SeqOf(Companion_D8_.Create_DC18_(), Companion_D8_.Create_DC18_()), _dafny.SeqOf(Companion_D8_.Create_DC18_()))).Cardinality(), true)).Cardinality(), !(!(false)), _dafny.IntOfInt64(981), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(470))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_21_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(263))).Cardinality())
	}))).Cardinality()))).Dtor_cf59()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lriyttbrc"), _dafny.UnicodeSeqOfUtf8Bytes("jftuf"))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(((_dafny.MultiSetOf(false, false, false, true, true)).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-45))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_22_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))))
	}))).Cardinality())), (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('j'))).Cardinality())).Plus(_dafny.IntOfInt64(271)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false, !((_dafny.IntOfInt64(-922)).Cmp(_dafny.IntOfInt64(583)) > 0), (false) || (!(false)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((Companion_D15_.Create_DC39_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Cardinality(), true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cynnxklv")).Cardinality()), _dafny.IntOfInt64(-334))).Dtor_cf59()).Plus(_dafny.IntOfInt64(-349))
}
func (_static *CompanionStruct_Default___) Fm20(globalState *GlobalState) _dafny.Map {
	if !(!(!(false))) || (false) {
		return func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(650), _dafny.IntOfInt64(968))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _23_v1 _dafny.Int
					_23_v1 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(650)).Cmp(_23_v1) <= 0) && ((_23_v1).Cmp(_dafny.IntOfInt64(968)) < 0) {
						_coll2.Add(Companion_Default___.SafeDivisionInt(_23_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ubretlxf")).Cardinality())), !(true))
					}
				}
				return _coll2.ToMap()
			}()).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _24_v0 _dafny.Int
				_24_v0 = interface{}(_compr_1).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll3 = _dafny.NewMapBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(650), _dafny.IntOfInt64(968))); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _23_v1 _dafny.Int
						_23_v1 = interface{}(_compr_3).(_dafny.Int)
						if ((_dafny.IntOfInt64(650)).Cmp(_23_v1) <= 0) && ((_23_v1).Cmp(_dafny.IntOfInt64(968)) < 0) {
							_coll3.Add(Companion_Default___.SafeDivisionInt(_23_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ubretlxf")).Cardinality())), !(true))
						}
					}
					return _coll3.ToMap()
				}()).Contains(_24_v0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_24_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(809)))
				}
			}
			return _coll1.ToMap()
		}()
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(!(false), !(false))).Cardinality(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true, true, false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kjccwlgyi")).Cardinality()), _dafny.IntOfInt64(12)))
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!((func() bool {
		if true {
			return !(false)
		}
		return true
	})()), false, !_dafny.Companion_Sequence_.Equal((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(309))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_25_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_26_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(false))))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Map, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.IntOfInt64(-93)).Cmp((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, !(false)))).Cardinality()) != 0, (_dafny.SetOf(_dafny.IntOfInt64(235), _dafny.IntOfInt64(10))).IsSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(-751))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	var _source2 D3 = Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ap"), (_dafny.UnicodeSeqOfUtf8Bytes("mthdwg")))).Cardinality(), _dafny.IntOfInt64(-666))).Cardinality()), true, false, (func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(188), _dafny.IntOfInt64(665))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _27_v0 _dafny.Int
			_27_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(188)).Cmp(_27_v0) <= 0) && ((_27_v0).Cmp(_dafny.IntOfInt64(665)) < 0) {
				_coll4.Add((_27_v0).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-960))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_29_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(954))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_28_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}))).Cardinality()))
			}
		}
		return _coll4.ToMap()
	}()).Cardinality(), false)
	_ = _source2
	if _source2.Is_DC6() {
		var _30___mcc_h0 _dafny.Int = _source2.Get_().(D3_DC6).Cf7
		_ = _30___mcc_h0
		var _31___mcc_h1 bool = _source2.Get_().(D3_DC6).Cf8
		_ = _31___mcc_h1
		var _32___mcc_h2 bool = _source2.Get_().(D3_DC6).Cf9
		_ = _32___mcc_h2
		var _33___mcc_h3 _dafny.Int = _source2.Get_().(D3_DC6).Cf10
		_ = _33___mcc_h3
		var _34___mcc_h4 bool = _source2.Get_().(D3_DC6).Cf11
		_ = _34___mcc_h4
		var _35_cf11 bool = _34___mcc_h4
		_ = _35_cf11
		var _36_cf10 _dafny.Int = _33___mcc_h3
		_ = _36_cf10
		var _37_cf9 bool = _32___mcc_h2
		_ = _37_cf9
		var _38_cf8 bool = _31___mcc_h1
		_ = _38_cf8
		var _39_cf7 _dafny.Int = _30___mcc_h0
		_ = _39_cf7
		return _dafny.MultiSetOf(_dafny.IntOfInt64(608))
	} else if _source2.Is_DC7() {
		var _40___mcc_h5 _dafny.Int = _source2.Get_().(D3_DC7).Cf12
		_ = _40___mcc_h5
		var _41_cf12 _dafny.Int = _40___mcc_h5
		_ = _41_cf12
		return (_dafny.MultiSetOf(_41_cf12, _41_cf12)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(321), _41_cf12)))
	} else if _source2.Is_DC8() {
		var _42___mcc_h6 bool = _source2.Get_().(D3_DC8).Cf13
		_ = _42___mcc_h6
		var _43___mcc_h7 _dafny.Int = _source2.Get_().(D3_DC8).Cf14
		_ = _43___mcc_h7
		var _44___mcc_h8 _dafny.Int = _source2.Get_().(D3_DC8).Cf15
		_ = _44___mcc_h8
		var _45_cf15 _dafny.Int = _44___mcc_h8
		_ = _45_cf15
		var _46_cf14 _dafny.Int = _43___mcc_h7
		_ = _46_cf14
		var _47_cf13 bool = _42___mcc_h6
		_ = _47_cf13
		return _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((Companion_D14_.Create_DC36_(_45_cf15, _47_cf13, _47_cf13, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gqwuhixo")).Cardinality()), _47_cf13)).Dtor_cf51())).Cardinality())))
	} else {
		var _48___mcc_h9 _dafny.Array = _source2.Get_().(D3_DC5).Cf6
		_ = _48___mcc_h9
		var _49_cf6 _dafny.Array = _48___mcc_h9
		_ = _49_cf6
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wiaxnjh")).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(841))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_50_i0 _dafny.Int) _dafny.Int {
		return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gmhrvhsnu")).Cardinality())).Plus(_dafny.IntOfInt64(207))
	}))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source3 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("r")
	_ = _source3
	var _51___mcc_h0 _dafny.Sequence = _source3
	_ = _51___mcc_h0
	var _52_cf20 _dafny.Sequence = _51___mcc_h0
	_ = _52_cf20
	return _dafny.CodePoint('w')
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-319), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(36))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_53_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})))).Cardinality())).Cardinality()), false))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("butfuvwqw")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(170), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, !(false), false, true)).Cardinality(), func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(((Companion_D17_.Create_DC44_(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.CodePoint('v'))).Cardinality())).Cardinality()), _dafny.IntOfInt64(728)))).Dtor_cf70()).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _54_v0 _dafny.Int
			_54_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((Companion_D17_.Create_DC44_(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.CodePoint('v'))).Cardinality())).Cardinality()), _dafny.IntOfInt64(728)))).Dtor_cf70()).Contains(_54_v0) {
				_coll5.Add((_54_v0).Plus(_dafny.IntOfInt64(866)), true)
			}
		}
		return _coll5.ToMap()
	}()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-349))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((Companion_D9_.Create_DC20_(true, _dafny.CodePoint('e'))).Dtor_cf29(), true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nwvavjx")).Cardinality()))).Cardinality(), false)))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D1 {
	var _source4 D8 = Companion_D8_.Create_DC18_()
	_ = _source4
	if _source4.Is_DC18() {
		return Companion_D1_.Create_DC1_(true)
	} else {
		var _55___mcc_h0 _dafny.Map = _source4.Get_().(D8_DC17).Cf27
		_ = _55___mcc_h0
		var _56_cf27 _dafny.Map = _55___mcc_h0
		_ = _56_cf27
		return Companion_D1_.Create_DC1_(true)
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC20_((func() bool {
		if !(false) {
			return true
		}
		return !(false)
	})(), _dafny.CodePoint('j'))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (_dafny.Int, bool, bool, _dafny.Map) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 bool = false
	_ = r2
	var r3 _dafny.Map = _dafny.EmptyMap
	_ = r3
	var _57_v0 bool
	_ = _57_v0
	_57_v0 = false
	var _58_v1 _dafny.Array
	_ = _58_v1
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
	_ = _nw0
	_58_v1 = _nw0
	var _59_v2 _dafny.Array
	_ = _59_v2
	var _nwElement0_0 _dafny.Array = _58_v1
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(6))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1(_58_v1, 1)
	(_nw1).ArraySet1(_58_v1, 2)
	(_nw1).ArraySet1(_58_v1, 3)
	(_nw1).ArraySet1(_58_v1, 4)
	(_nw1).ArraySet1(_58_v1, 5)
	_59_v2 = _nw1
	var _60_v3 D3
	_ = _60_v3
	_60_v3 = Companion_D3_.Create_DC5_(_59_v2)
	var _61_v4 *C1
	_ = _61_v4
	var _nw2 *C1 = New_C1_()
	_ = _nw2
	_nw2.Ctor__((func() bool {
		if _57_v0 {
			return _57_v0
		}
		return _57_v0
	})(), _60_v3)
	_61_v4 = _nw2
	var _62_v5 _dafny.Int
	_ = _62_v5
	_62_v5 = _dafny.IntOfInt64(188)
	if (func() bool {
		if _57_v0 {
			return _57_v0
		}
		return Companion_Default___.Fm0(_62_v5, _62_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer12 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_63_v4 *C1) func(_dafny.Int) _dafny.MultiSet {
			return func(_64_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf((_63_v4).F16())
			}
		})(_61_v4))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rkrihtkwb"), !((_61_v4).F16())), globalState)
	})() {
		var _65_v6 _dafny.MultiSet
		_ = _65_v6
		_65_v6 = _dafny.MultiSetOf(_57_v0, _57_v0)
		var _66_v7 _dafny.Sequence
		_ = _66_v7
		_66_v7 = _dafny.SeqOf(_65_v6)
		var _67_v8 _dafny.Sequence
		_ = _67_v8
		_67_v8 = _dafny.UnicodeSeqOfUtf8Bytes("gmqmsbnsb")
		var _68_v9 _dafny.CodePoint
		_ = _68_v9
		_68_v9 = _dafny.CodePoint('d')
		var _69_v10 _dafny.Map
		_ = _69_v10
		_69_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v8, (_61_v4).Fm4(_62_v5, true, _68_v9, globalState))
		if ((func() bool {
			if (_61_v4).F16() {
				return false
			}
			return Companion_Default___.Fm0(_62_v5, _62_v5, _66_v7, _69_v10, globalState)
		})()) == ((_61_v4).F16()) {
			r1 = !(_57_v0)
			_58_v1 = _58_v1
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_58_v1), 0))
			_ = _index0
			(_58_v1).ArraySet1(_57_v0, (_index0).Int())
			var _70_v11 _dafny.Sequence
			_ = _70_v11
			_70_v11 = _dafny.SeqOf((_61_v4).F16())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_58_v1), 0))
			_ = _index1
			var _rhs0 _dafny.CodePoint = _68_v9
			_ = _rhs0
			var _rhs1 bool = (_70_v11).Select((Companion_Default___.SafeIndex(_62_v5, _dafny.IntOfUint32((_70_v11).Cardinality()))).Uint32()).(bool)
			_ = _rhs1
			var _rhs2 bool = (_61_v4).F16()
			_ = _rhs2
			var _lhs0 _dafny.Array = _58_v1
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_58_v1), 0))
			_ = _lhs1
			_68_v9 = _rhs0
			(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
			_57_v0 = _rhs2
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_58_v1), 0))
			_ = _index2
			(_58_v1).ArraySet1((_61_v4).F16(), (_index2).Int())
			var _71_v12 _dafny.Array
			_ = _71_v12
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw3
			_71_v12 = _nw3
			_71_v12 = _71_v12
		} else {
			r2 = !(!((_61_v4).F16())) || ((_61_v4).F16())
			var _72_v13 _dafny.Map
			_ = _72_v13
			_72_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_62_v5), (_61_v4).F16())
			var _73_v14 _dafny.Sequence
			_ = _73_v14
			_73_v14 = _dafny.SeqOf((_72_v13).Cardinality())
			var _74_v15 _dafny.Set
			_ = _74_v15
			_74_v15 = _dafny.SetOf(_62_v5, _dafny.IntOfUint32((_73_v14).Cardinality()), _62_v5)
			var _75_v16 _dafny.Map
			_ = _75_v16
			_75_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v1, _67_v8)
			var _76_v17 _dafny.Map
			_ = _76_v17
			_76_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_74_v15).Cardinality(), _75_v16)
			_76_v17 = (_76_v17).Update(_dafny.IntOfInt64(184), _75_v16)
			_69_v10 = (_69_v10).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_77_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_78_i1 _dafny.Int) _dafny.CodePoint {
					return _77_v9
				}
			})(_68_v9))), (_61_v4).F16())
			var _79_v18 _dafny.MultiSet
			_ = _79_v18
			_79_v18 = _dafny.MultiSetOf(_68_v9)
			var _80_v19 D15
			_ = _80_v19
			_80_v19 = Companion_D15_.Create_DC39_((_61_v4).F16(), (_79_v18).Cardinality(), false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_81_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_82_i2 _dafny.Int) _dafny.CodePoint {
					return _81_v9
				}
			})(_68_v9)))).Cardinality()), _62_v5)
			var _83_v20 _dafny.Map
			_ = _83_v20
			_83_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v5, _80_v19)
			_83_v20 = (_83_v20).Update((_62_v5).Times(_62_v5), _80_v19)
			(globalState).F2 = _62_v5
		}
		var _84_v21 _dafny.Array
		_ = _84_v21
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_0
		var _nw4 _dafny.Array
		_ = _nw4
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw4 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_85_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_86_i3 _dafny.Int) _dafny.Int {
					return (_86_i3).Times(_85_v5)
				}
			})(_62_v5)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw4 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw4).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw4).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_84_v21 = _nw4
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
		_ = _index3
		(_84_v21).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(127), _62_v5), (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
		_ = _index4
		(_84_v21).ArraySet1(Companion_Default___.SafeModuloInt(_62_v5, _62_v5), (_index4).Int())
		r1 = (_61_v4).F16()
		if (_61_v4).F16() {
			(globalState).F4 = _84_v21
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))
			_ = _index5
			(_59_v2).ArraySet1(_58_v1, (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))
			_ = _index6
			(_59_v2).ArraySet1(_58_v1, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _index7
			(_84_v21).ArraySet1((_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int), (_index7).Int())
			var _87_v22 _dafny.Map
			_ = _87_v22
			_87_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v0, (_65_v6).Cardinality())
			var _88_v23 _dafny.Set
			_ = _88_v23
			_88_v23 = _dafny.SetOf(_dafny.IntOfUint32((_67_v8).Cardinality()))
			var _89_v24 _dafny.Sequence
			_ = _89_v24
			_89_v24 = _dafny.SeqOf(_88_v23, _88_v23)
			var _90_v25 *C0
			_ = _90_v25
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(Companion_Default___.Fm0(Companion_Default___.Fm19((_87_v22).Cardinality(), globalState), (_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int), _66_v7, _69_v10, globalState), (_dafny.IntOfUint32((_89_v24).Cardinality())).Minus(_62_v5))
			_90_v25 = _nw5
			var _91_v26 T0
			_ = _91_v26
			var _nw6 *C2 = New_C2_()
			_ = _nw6
			_nw6.Ctor__(_57_v0, ((_90_v25).F12()) && (true))
			_91_v26 = _nw6
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_59_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))).Int()))
			_ = _arr0
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_dafny.ArrayCastTo((_59_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))).Int()))), 0))
			_ = _index8
			(_arr0).ArraySet1((_61_v4).F16(), (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _index9
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_59_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))).Int()))
			_ = _arr1
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_dafny.ArrayCastTo((_59_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))).Int()))), 0))
			_ = _index10
			var _rhs3 _dafny.Int = (_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)
			_ = _rhs3
			var _rhs4 *C0 = _90_v25
			_ = _rhs4
			var _rhs5 T0 = _91_v26
			_ = _rhs5
			var _rhs6 bool = !((_61_v4).F16())
			_ = _rhs6
			var _rhs7 bool = (_91_v26).Fm4(_dafny.IntOfInt64(994), !(false), _dafny.CodePoint('h'), globalState)
			_ = _rhs7
			var _lhs2 _dafny.Array = _84_v21
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _lhs3
			var _lhs4 _dafny.Array = _dafny.ArrayCastTo((_59_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))).Int()))
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_dafny.ArrayCastTo((_59_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_59_v2), 0))).Int()))), 0))
			_ = _lhs5
			(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
			_90_v25 = _rhs4
			_91_v26 = _rhs5
			r1 = _rhs6
			(_lhs4).ArraySet1(_rhs7, (_lhs5).Int())
			var _92_v27 _dafny.Array
			_ = _92_v27
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
			_ = _nw7
			_92_v27 = _nw7
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_92_v27), 0))
			_ = _index11
			(_92_v27).ArraySet1(_69_v10, (_index11).Int())
			var _93_v28 _dafny.Sequence
			_ = _93_v28
			_93_v28 = _dafny.SeqOf((_90_v25).F13(), _62_v5)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _index12
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_92_v27), 0))
			_ = _index13
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _index14
			var _rhs8 _dafny.Int = (func() _dafny.Int {
				if (_90_v25).F12() {
					return (_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)
				}
				return (_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)
			})()
			_ = _rhs8
			var _rhs9 _dafny.Map = _69_v10
			_ = _rhs9
			var _rhs10 bool = (func() bool {
				if (func() bool {
					if false {
						return Companion_Default___.Fm0((_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int), _62_v5, _66_v7, _69_v10, globalState)
					}
					return (_61_v4).F16()
				})() {
					return (_61_v4).F16()
				}
				return !((_62_v5).Cmp((_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)) < 0)
			})()
			_ = _rhs10
			var _rhs11 _dafny.Int = ((_90_v25).F13()).Plus((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(204))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_94_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_95_i4 _dafny.Int) _dafny.CodePoint {
					return _94_v9
				}
			})(_68_v9)))).Cardinality())).Minus((_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)))
			_ = _rhs11
			var _rhs12 _dafny.Int = ((_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)).Plus((_93_v28).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_90_v25).F13()), _dafny.IntOfUint32((_93_v28).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs12
			var _lhs6 _dafny.Array = _84_v21
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _lhs7
			var _lhs8 _dafny.Array = _92_v27
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_92_v27), 0))
			_ = _lhs9
			var _lhs10 _dafny.Array = _84_v21
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))
			_ = _lhs11
			(_lhs6).ArraySet1(_rhs8, (_lhs7).Int())
			(_lhs8).ArraySet1(_rhs9, (_lhs9).Int())
			_57_v0 = _rhs10
			(_lhs10).ArraySet1(_rhs11, (_lhs11).Int())
			_62_v5 = _rhs12
		} else {
			_67_v8 = _67_v8
			var _96_v29 _dafny.Set
			_ = _96_v29
			_96_v29 = _dafny.SetOf(_58_v1)
			var _97_v30 _dafny.Sequence
			_ = _97_v30
			_97_v30 = _dafny.SeqOf(Companion_D7_.Create_DC14_(_96_v29))
			_97_v30 = _97_v30
			var _98_v31 *C2
			_ = _98_v31
			var _nw8 *C2 = New_C2_()
			_ = _nw8
			_nw8.Ctor__((func() bool {
				if (_61_v4).Fm4(_dafny.IntOfInt64(567), (_61_v4).F16(), _68_v9, globalState) {
					return !(false)
				}
				return _57_v0
			})(), true)
			_98_v31 = _nw8
			_98_v31 = _98_v31
			var _99_v32 _dafny.Sequence
			_ = _99_v32
			_99_v32 = _dafny.SeqOf(_67_v8)
			var _100_v33 *C1
			_ = _100_v33
			var _nw9 *C1 = New_C1_()
			_ = _nw9
			_nw9.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_67_v8, (_99_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_67_v8).Cardinality()), _dafny.IntOfUint32((_99_v32).Cardinality()))).Uint32()).(_dafny.Sequence)), _61_v4.F17)
			_100_v33 = _nw9
			_68_v9 = _68_v9
		}
		var _101_v34 _dafny.Map
		_ = _101_v34
		_101_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v0, ((_84_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_84_v21), 0))).Int()).(_dafny.Int)).Cmp(_62_v5) >= 0)
		_101_v34 = (_101_v34).Update(!((_61_v4).F16()) || (_57_v0), !(_57_v0) || ((_61_v4).F16()))
	} else {
		var _102_v35 *C2
		_ = _102_v35
		var _nw10 *C2 = New_C2_()
		_ = _nw10
		_nw10.Ctor__((_61_v4).F16(), (_61_v4).F16())
		_102_v35 = _nw10
		_102_v35 = _102_v35
		var _103_v36 _dafny.Map
		_ = _103_v36
		_103_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_61_v4).F16(), _102_v35.F14)
		var _104_v37 D9
		_ = _104_v37
		_104_v37 = Companion_D9_.Create_DC20_(!((func() bool {
			if (_103_v36).Contains((_102_v35).F15()) {
				return (_103_v36).Get((_102_v35).F15()).(bool)
			}
			return (_102_v35).F15()
		})()), _dafny.CodePoint('t'))
		var _105_v38 _dafny.Sequence
		_ = _105_v38
		_105_v38 = _dafny.UnicodeSeqOfUtf8Bytes("kplyjx")
		var _106_v39 _dafny.Sequence
		_ = _106_v39
		_106_v39 = _dafny.SeqOf(_dafny.IntOfInt64(647), _62_v5)
		var _107_v40 D16
		_ = _107_v40
		_107_v40 = Companion_D16_.Create_DC42_((_106_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.IntOfUint32((_106_v39).Cardinality()))).Uint32()).(_dafny.Int), (_61_v4).F16())
		var _108_v41 _dafny.CodePoint
		_ = _108_v41
		_108_v41 = _dafny.CodePoint('a')
		var _109_v42 _dafny.Map
		_ = _109_v42
		_109_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_110_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_111_i5 _dafny.Int) _dafny.CodePoint {
				return _110_v41
			}
		})(_108_v41))))
		var _rhs13 bool = !(false)
		_ = _rhs13
		var _rhs14 D9 = Companion_Default___.Fm30((_107_v40).Dtor_cf65(), _dafny.IntOfInt64(210), _108_v41, _62_v5, globalState)
		_ = _rhs14
		var _rhs15 _dafny.Sequence = (func() _dafny.Sequence {
			if (_109_v42).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_105_v38).Cardinality()))) {
				return (_109_v42).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_105_v38).Cardinality()))).(_dafny.Sequence)
			}
			return Companion_Default___.Fm2(_102_v35.F14, _62_v5, _dafny.IntOfInt64(875), true, globalState)
		})()
		_ = _rhs15
		var _rhs16 bool = ((_dafny.IntOfInt64(970)).Minus(_62_v5)).Cmp((_dafny.Zero).Minus(_62_v5)) < 0
		_ = _rhs16
		var _lhs12 *C2 = _102_v35
		_ = _lhs12
		r1 = _rhs13
		_104_v37 = _rhs14
		_105_v38 = _rhs15
		_lhs12.F14 = _rhs16
		r1 = (_61_v4).F16()
		if _57_v0 {
			var _112_v43 T0
			_ = _112_v43
			var _nw11 *C1 = New_C1_()
			_ = _nw11
			_nw11.Ctor__((_61_v4).F16(), _61_v4.F17)
			_112_v43 = _nw11
			var _113_v44 _dafny.Map
			_ = _113_v44
			_113_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v38, Companion_Default___.SafeDivisionInt(_62_v5, _62_v5))
			var _114_v45 _dafny.Sequence
			_ = _114_v45
			_114_v45 = _dafny.SeqOf(_112_v43, _112_v43, _112_v43)
			var _115_v47 _dafny.Set
			_ = _115_v47
			_115_v47 = _dafny.SetOf(_105_v38)
			var _rhs17 T0 = (_114_v45).Select((Companion_Default___.SafeIndex(_62_v5, _dafny.IntOfUint32((_114_v45).Cardinality()))).Uint32()).(T0)
			_ = _rhs17
			var _rhs18 _dafny.Int = _62_v5
			_ = _rhs18
			var _rhs19 D3 = _60_v3
			_ = _rhs19
			var _rhs20 _dafny.Int = _62_v5
			_ = _rhs20
			var _rhs21 _dafny.Map = (func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_115_v47).Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _116_v46 _dafny.Sequence
					_116_v46 = interface{}(_compr_6).(_dafny.Sequence)
					if (_115_v47).Contains(_116_v46) {
						_coll6.Add(_116_v46, (_dafny.Zero).Minus((_dafny.SetOf((_dafny.Zero).Minus(_62_v5))).Cardinality()))
					}
				}
				return _coll6.ToMap()
			}()).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-734))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_117_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_118_i6 _dafny.Int) _dafny.CodePoint {
					return _117_v41
				}
			})(_108_v41))), _62_v5)
			_ = _rhs21
			var _lhs13 *GlobalState = globalState
			_ = _lhs13
			var _lhs14 *GlobalState = globalState
			_ = _lhs14
			_112_v43 = _rhs17
			_lhs13.F2 = _rhs18
			_60_v3 = _rhs19
			_lhs14.F2 = _rhs20
			_113_v44 = _rhs21
			var _119_v48 _dafny.MultiSet
			_ = _119_v48
			_119_v48 = _dafny.MultiSetOf(_112_v43)
			var _120_v49 _dafny.Sequence
			_ = _120_v49
			_120_v49 = _dafny.SeqOf((_61_v4).F16(), (_102_v35).Fm12(true, (func() _dafny.Int {
				if (_119_v48).Contains(_112_v43) {
					return (_119_v48).Multiplicity(_112_v43)
				}
				return _62_v5
			})(), _dafny.MultiSetOf(_62_v5), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_121_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_122_i7 _dafny.Int) _dafny.CodePoint {
					return _121_v41
				}
			})(_108_v41)))).Cardinality()), globalState))
			var _123_v50 _dafny.Map
			_ = _123_v50
			_123_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v5, _57_v0)
			var _124_v51 T1
			_ = _124_v51
			var _nw12 *C3 = New_C3_()
			_ = _nw12
			_nw12.Ctor__(_62_v5, _123_v50)
			_124_v51 = _nw12
			var _125_v52 _dafny.Map
			_ = _125_v52
			_125_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v0, _dafny.IntOfUint32((_120_v49).Cardinality()))
			var _126_v53 _dafny.Set
			_ = _126_v53
			_126_v53 = _dafny.SetOf(_dafny.IntOfInt64(198))
			var _127_v54 _dafny.MultiSet
			_ = _127_v54
			_127_v54 = _dafny.MultiSetOf(_102_v35.F14, (_102_v35).F15())
			var _128_v55 D15
			_ = _128_v55
			_128_v55 = Companion_D15_.Create_DC39_(_102_v35.F14, _dafny.IntOfUint32((_105_v38).Cardinality()), (_61_v4).F16(), (_127_v54).Cardinality(), _124_v51.F10())
			var _129_v56 _dafny.Int
			_ = _129_v56
			_129_v56 = _124_v51.F10()
			var _130_v57 _dafny.Map
			_ = _130_v57
			_130_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v56, _102_v35.F14)
			var _131_v58 _dafny.Array
			_ = _131_v58
			var _nwElement0_1 _dafny.Int = _62_v5
			_ = _nwElement0_1
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(9))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_1, 0)
			(_nw13).ArraySet1((_103_v36).Cardinality(), 1)
			(_nw13).ArraySet1((_62_v5).Minus((func() _dafny.Int {
				if (_125_v52).Contains(_102_v35.F14) {
					return (_125_v52).Get(_102_v35.F14).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_105_v38).Cardinality())
			})()), 2)
			(_nw13).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm18((_126_v53).Cardinality(), (_128_v55).Dtor_cf58(), _124_v51.F10(), globalState)).Cardinality()), 3)
			(_nw13).ArraySet1(_62_v5, 4)
			(_nw13).ArraySet1((Companion_Default___.Fm10((_102_v35).F15(), globalState)).Minus(_dafny.IntOfUint32((Companion_Default___.Fm25(_62_v5, Companion_Default___.Fm17(_102_v35.F14, _124_v51.F10(), (_102_v35).F15(), (_130_v57).Cardinality(), globalState), _124_v51.F10(), _62_v5, globalState)).Cardinality())), 5)
			(_nw13).ArraySet1(_124_v51.F10(), 6)
			(_nw13).ArraySet1(_62_v5, 7)
			(_nw13).ArraySet1(_62_v5, 8)
			_131_v58 = _nw13
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_131_v58), 0))
			_ = _index15
			(_131_v58).ArraySet1(((_124_v51).F11()).Cardinality(), (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_131_v58), 0))
			_ = _index16
			var _rhs22 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_120_v49, (Companion_Default___.SafeIndex(_124_v51.F10(), _dafny.IntOfUint32((_120_v49).Cardinality()))).Uint32(), _102_v35.F14), (Companion_Default___.SafeIndex((_127_v54).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_120_v49, (Companion_Default___.SafeIndex(_124_v51.F10(), _dafny.IntOfUint32((_120_v49).Cardinality()))).Uint32(), _102_v35.F14)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.IsProperPrefixOf(_105_v38, _105_v38))
			_ = _rhs22
			var _rhs23 _dafny.Int = _124_v51.F10()
			_ = _rhs23
			var _rhs24 T1 = _124_v51
			_ = _rhs24
			var _rhs25 _dafny.Int = _62_v5
			_ = _rhs25
			var _rhs26 _dafny.Int = _124_v51.F10()
			_ = _rhs26
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _131_v58
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_131_v58), 0))
			_ = _lhs17
			_120_v49 = _rhs22
			_lhs15.F2 = _rhs23
			_124_v51 = _rhs24
			(_lhs16).ArraySet1(_rhs25, (_lhs17).Int())
			_62_v5 = _rhs26
			_106_v39 = _106_v39
			r0 = (_dafny.Zero).Minus((_106_v39).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_62_v5), _dafny.IntOfUint32((_106_v39).Cardinality()))).Uint32()).(_dafny.Int))
			var _132_v59 _dafny.MultiSet
			_ = _132_v59
			_132_v59 = _dafny.MultiSetOf(_131_v58, _131_v58)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_58_v1), 0))
			_ = _index17
			(_58_v1).ArraySet1(((_131_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_131_v58), 0))).Int()).(_dafny.Int)).Cmp((_132_v59).Cardinality()) == 0, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_58_v1), 0))
			_ = _index18
			var _rhs27 _dafny.Array = _58_v1
			_ = _rhs27
			var _rhs28 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D3_.Create_DC8_((_61_v4).F16(), _124_v51.F10(), (_dafny.Zero).Minus(_124_v51.F10()))).Dtor_cf13(), (_61_v4).F16())).Update(false, !((Companion_D3_.Create_DC6_((_131_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_131_v58), 0))).Int()).(_dafny.Int), _57_v0, (_61_v4).F16(), _124_v51.F10(), (_102_v35).F15())).Dtor_cf8()))).Cardinality()
			_ = _rhs28
			var _rhs29 bool = _102_v35.F14
			_ = _rhs29
			var _rhs30 _dafny.Int = Companion_Default___.SafeDivisionInt(_124_v51.F10(), _124_v51.F10())
			_ = _rhs30
			var _lhs18 _dafny.Array = _58_v1
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_58_v1), 0))
			_ = _lhs19
			_58_v1 = _rhs27
			_62_v5 = _rhs28
			(_lhs18).ArraySet1(_rhs29, (_lhs19).Int())
			r0 = _rhs30
		} else {
			(_102_v35).F14 = (_62_v5).Cmp((func() _dafny.Int {
				if (_61_v4).F16() {
					return _dafny.IntOfInt64(376)
				}
				return _dafny.IntOfUint32((_105_v38).Cardinality())
			})()) >= 0
			var _133_v60 _dafny.Set
			_ = _133_v60
			_133_v60 = _dafny.SetOf(_106_v39, _106_v39)
			r0 = (_dafny.Zero).Minus((_133_v60).Cardinality())
			var _134_v61 _dafny.Map
			_ = _134_v61
			_134_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_62_v5).Minus(_62_v5), _62_v5)
			var _rhs31 _dafny.Int = _62_v5
			_ = _rhs31
			var _rhs32 _dafny.Map = _134_v61
			_ = _rhs32
			var _lhs20 *GlobalState = globalState
			_ = _lhs20
			_lhs20.F2 = _rhs31
			_134_v61 = _rhs32
			var _135_v62 _dafny.Array
			_ = _135_v62
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw14
			_135_v62 = _nw14
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_135_v62), 0))
			_ = _index19
			(_135_v62).ArraySet1(_dafny.IntOfUint32((_106_v39).Cardinality()), (_index19).Int())
			var _136_v63 _dafny.Map
			_ = _136_v63
			_136_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_61_v4).F16(), (_dafny.Zero).Minus(_62_v5))
			var _137_v64 D9
			_ = _137_v64
			_137_v64 = Companion_D9_.Create_DC19_(_136_v63)
			var _138_v65 D9
			_ = _138_v65
			_138_v65 = Companion_D9_.Create_DC21_(_137_v64)
			var _139_v66 _dafny.Array
			_ = _139_v66
			var _nwElement0_2 D9 = Companion_D9_.Create_DC21_(_137_v64)
			_ = _nwElement0_2
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_2, 0)
			(_nw15).ArraySet1(_138_v65, 1)
			_139_v66 = _nw15
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_139_v66), 0))
			_ = _index20
			(_139_v66).ArraySet1(_138_v65, (_index20).Int())
			var _140_v67 _dafny.Map
			_ = _140_v67
			_140_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v5, (_61_v4).F16())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_135_v62), 0))
			_ = _index21
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_139_v66), 0))
			_ = _index22
			var _rhs33 *C1 = _61_v4
			_ = _rhs33
			var _rhs34 bool = (_102_v35.F14) == (!(_57_v0) || ((func() bool {
				if (_140_v67).Contains(_62_v5) {
					return (_140_v67).Get(_62_v5).(bool)
				}
				return _102_v35.F14
			})()))
			_ = _rhs34
			var _rhs35 _dafny.Int = Companion_Default___.Fm19(_62_v5, globalState)
			_ = _rhs35
			var _rhs36 D9 = _138_v65
			_ = _rhs36
			var _lhs21 _dafny.Array = _135_v62
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_135_v62), 0))
			_ = _lhs22
			var _lhs23 _dafny.Array = _139_v66
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_139_v66), 0))
			_ = _lhs24
			_61_v4 = _rhs33
			_57_v0 = _rhs34
			(_lhs21).ArraySet1(_rhs35, (_lhs22).Int())
			(_lhs23).ArraySet1(_rhs36, (_lhs24).Int())
			var _141_v68 _dafny.Sequence
			_ = _141_v68
			_141_v68 = _dafny.SeqOf((_102_v35).F15(), true)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_135_v62), 0))
			_ = _index23
			(_135_v62).ArraySet1(_dafny.IntOfUint32((_141_v68).Cardinality()), (_index23).Int())
		}
		_58_v1 = _58_v1
	}
	var _142_v69 *C0
	_ = _142_v69
	var _nw16 *C0 = New_C0_()
	_ = _nw16
	_nw16.Ctor__(_57_v0, _62_v5)
	_142_v69 = _nw16
	var _143_v70 _dafny.Map
	_ = _143_v70
	_143_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_61_v4).F16(), _142_v69)
	_143_v70 = (_143_v70).Update((_61_v4).F16(), _142_v69)
	r1 = false
	r2 = _57_v0
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_58_v1), 0))
	_ = _index24
	(_58_v1).ArraySet1(_57_v0, (_index24).Int())
	var _144_v71 _dafny.Array
	_ = _144_v71
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_1
	var _nw17 _dafny.Array
	_ = _nw17
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw17 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = (func(_145_v0 bool) func(_dafny.Int) _dafny.Int {
			return func(_146_i8 _dafny.Int) _dafny.Int {
				return (_146_i8).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_145_v0)).Cardinality()))
			}
		})(_57_v0)
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
	_144_v71 = _nw17
	var _147_v72 _dafny.Sequence
	_ = _147_v72
	_147_v72 = _dafny.SeqOf((_142_v69).F13())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_58_v1), 0))
	_ = _index25
	var _rhs37 _dafny.Array = _144_v71
	_ = _rhs37
	var _rhs38 _dafny.Int = (Companion_Default___.Fm10(_57_v0, globalState)).Plus(((_142_v69).F13()).Minus((_147_v72).Select((Companion_Default___.SafeIndex((_142_v69).F13(), _dafny.IntOfUint32((_147_v72).Cardinality()))).Uint32()).(_dafny.Int)))
	_ = _rhs38
	var _rhs39 bool = _57_v0
	_ = _rhs39
	var _lhs25 *GlobalState = globalState
	_ = _lhs25
	var _lhs26 _dafny.Array = _58_v1
	_ = _lhs26
	var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_58_v1), 0))
	_ = _lhs27
	_lhs25.F4 = _rhs37
	_62_v5 = _rhs38
	(_lhs26).ArraySet1(_rhs39, (_lhs27).Int())
	r0 = _dafny.IntOfInt64(809)
	var _148_v73 _dafny.CodePoint
	_ = _148_v73
	_148_v73 = _dafny.CodePoint('c')
	var _149_v74 _dafny.Sequence
	_ = _149_v74
	_149_v74 = _dafny.SeqOf(_148_v73)
	r1 = (_dafny.IntOfUint32((_149_v74).Cardinality())).Cmp(_62_v5) < 0
	r2 = !((func() bool {
		if (_142_v69).F12() {
			return ((_58_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_58_v1), 0))).Int()).(bool)) == (_57_v0)
		}
		return _57_v0
	})())
	var _150_v75 _dafny.Map
	_ = _150_v75
	_150_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_58_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_58_v1), 0))).Int()).(bool), (_142_v69).F13())
	r3 = (func() _dafny.Map {
		if (_142_v69).F12() {
			return (func() _dafny.Map {
				if (_61_v4).F16() {
					return _150_v75
				}
				return _150_v75
			})()
		}
		return _150_v75
	})()
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _151_v0 _dafny.Array
	_ = _151_v0
	var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
	_ = _nw18
	_151_v0 = _nw18
	var _152_v1 _dafny.Int
	_ = _152_v1
	_152_v1 = _dafny.IntOfInt64(581)
	var _153_v2 _dafny.MultiSet
	_ = _153_v2
	_153_v2 = _dafny.MultiSetOf(_152_v1, _152_v1)
	var _154_v3 _dafny.Sequence
	_ = _154_v3
	_154_v3 = _dafny.SeqOf((_153_v2).Cardinality())
	var _155_v4 bool
	_ = _155_v4
	_155_v4 = false
	var _156_v5 _dafny.Map
	_ = _156_v5
	_156_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _155_v4)
	var _157_v6 _dafny.Sequence
	_ = _157_v6
	_157_v6 = _dafny.SeqOf(_155_v4)
	var _158_v7 _dafny.Set
	_ = _158_v7
	_158_v7 = _dafny.SetOf(_dafny.IntOfUint32((_157_v6).Cardinality()))
	var _159_v8 _dafny.Int
	_ = _159_v8
	_159_v8 = _152_v1
	var _160_v9 _dafny.Sequence
	_ = _160_v9
	_160_v9 = _dafny.UnicodeSeqOfUtf8Bytes("joueqk")
	var _161_v10 _dafny.Array
	_ = _161_v10
	var _nwElement0_3 _dafny.Int = _dafny.IntOfInt64(-822)
	_ = _nwElement0_3
	var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(29))
	_ = _nw19
	(_nw19).ArraySet1(_nwElement0_3, 0)
	(_nw19).ArraySet1(_152_v1, 1)
	(_nw19).ArraySet1(_152_v1, 2)
	(_nw19).ArraySet1(_152_v1, 3)
	(_nw19).ArraySet1(_152_v1, 4)
	(_nw19).ArraySet1(_152_v1, 5)
	(_nw19).ArraySet1((_dafny.Zero).Minus(_152_v1), 6)
	(_nw19).ArraySet1(_dafny.IntOfInt64(288), 7)
	(_nw19).ArraySet1(_152_v1, 8)
	(_nw19).ArraySet1(_152_v1, 9)
	(_nw19).ArraySet1(_dafny.IntOfUint32((_154_v3).Cardinality()), 10)
	(_nw19).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_156_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(55), _155_v4), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _155_v4)), (Companion_Default___.SafeIndex((_158_v7).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_156_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(55), _155_v4), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _155_v4))).Cardinality()))).Uint32(), _156_v5)).Cardinality()), 11)
	(_nw19).ArraySet1(_152_v1, 12)
	(_nw19).ArraySet1(_152_v1, 13)
	(_nw19).ArraySet1((_156_v5).Cardinality(), 14)
	(_nw19).ArraySet1(_152_v1, 15)
	(_nw19).ArraySet1(_152_v1, 16)
	(_nw19).ArraySet1(_152_v1, 17)
	(_nw19).ArraySet1((_159_v8), 18)
	(_nw19).ArraySet1(_152_v1, 19)
	(_nw19).ArraySet1(_152_v1, 20)
	(_nw19).ArraySet1(_dafny.IntOfInt64(576), 21)
	(_nw19).ArraySet1(_152_v1, 22)
	(_nw19).ArraySet1(_dafny.IntOfUint32((_160_v9).Cardinality()), 23)
	(_nw19).ArraySet1(_152_v1, 24)
	(_nw19).ArraySet1(_dafny.IntOfInt64(-713), 25)
	(_nw19).ArraySet1(_152_v1, 26)
	(_nw19).ArraySet1(_152_v1, 27)
	(_nw19).ArraySet1(_152_v1, 28)
	_161_v10 = _nw19
	var _162_globalState *GlobalState
	_ = _162_globalState
	var _nw20 *GlobalState = New_GlobalState_()
	_ = _nw20
	_nw20.Ctor__(true, false, _dafny.IntOfInt64(572), _151_v0, _161_v10, _dafny.IntOfInt64(807), _dafny.IntOfInt64(286), _161_v10, _dafny.IntOfInt64(-297), _dafny.IntOfInt64(926))
	_162_globalState = _nw20
	var _163_v11 _dafny.Array
	_ = _163_v11
	var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(2))
	_ = _nw21
	_163_v11 = _nw21
	var _164_v12 _dafny.Map
	_ = _164_v12
	_164_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _163_v11)
	_164_v12 = (_164_v12).Update(_152_v1, _163_v11)
	var _165_v13 _dafny.MultiSet
	_ = _165_v13
	_165_v13 = _dafny.MultiSetOf(_155_v4, _155_v4, _155_v4, _155_v4)
	var _166_v14 _dafny.Sequence
	_ = _166_v14
	_166_v14 = _dafny.SeqOf(_165_v13)
	var _167_v15 _dafny.Map
	_ = _167_v15
	_167_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v9, _155_v4)
	if Companion_Default___.Fm0(_152_v1, _152_v1, _166_v14, _167_v15, _162_globalState) {
		var _168_v16 _dafny.Int
		_ = _168_v16
		var _169_v17 bool
		_ = _169_v17
		var _170_v18 bool
		_ = _170_v18
		var _171_v19 _dafny.Map
		_ = _171_v19
		var _out0 _dafny.Int
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 bool
		_ = _out2
		var _out3 _dafny.Map
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0(_162_globalState)
		_168_v16 = _out0
		_169_v17 = _out1
		_170_v18 = _out2
		_171_v19 = _out3
		var _172_v20 _dafny.Array
		_ = _172_v20
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_2
		var _nw22 _dafny.Array
		_ = _nw22
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw22 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_173_v17 bool) func(_dafny.Int) bool {
				return func(_174_i0 _dafny.Int) bool {
					return (Companion_D1_.Create_DC2_(_173_v17)).Dtor_cf2()
				}
			})(_169_v17)
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
		_172_v20 = _nw22
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_172_v20), 0))
		_ = _index26
		(_172_v20).ArraySet1((_152_v1).Cmp(_dafny.IntOfUint32((_157_v6).Cardinality())) <= 0, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_172_v20), 0))
		_ = _index27
		var _rhs40 bool = false
		_ = _rhs40
		var _rhs41 bool = _170_v18
		_ = _rhs41
		var _rhs42 _dafny.Int = _dafny.IntOfInt64(261)
		_ = _rhs42
		var _rhs43 bool = (_168_v16).Cmp(_152_v1) >= 0
		_ = _rhs43
		var _lhs28 *GlobalState = _162_globalState
		_ = _lhs28
		var _lhs29 _dafny.Array = _172_v20
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_172_v20), 0))
		_ = _lhs30
		_169_v17 = _rhs40
		_169_v17 = _rhs41
		_lhs28.F2 = _rhs42
		(_lhs29).ArraySet1(_rhs43, (_lhs30).Int())
		var _175_v21 D1
		_ = _175_v21
		_175_v21 = Companion_D1_.Create_DC2_(_169_v17)
		_175_v21 = Companion_D1_.Create_DC2_(true)
		_170_v18 = _dafny.Companion_Sequence_.Equal(_154_v3, _dafny.SeqOf(_168_v16))
		var _176_v22 _dafny.Int
		_ = _176_v22
		var _177_v23 bool
		_ = _177_v23
		var _178_v24 bool
		_ = _178_v24
		var _179_v25 _dafny.Map
		_ = _179_v25
		var _out4 _dafny.Int
		_ = _out4
		var _out5 bool
		_ = _out5
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Map
		_ = _out7
		_out4, _out5, _out6, _out7 = Companion_Default___.M0(_162_globalState)
		_176_v22 = _out4
		_177_v23 = _out5
		_178_v24 = _out6
		_179_v25 = _out7
	} else {
		var _180_v26 _dafny.Int
		_ = _180_v26
		var _181_v27 bool
		_ = _181_v27
		var _182_v28 bool
		_ = _182_v28
		var _183_v29 _dafny.Map
		_ = _183_v29
		var _out8 _dafny.Int
		_ = _out8
		var _out9 bool
		_ = _out9
		var _out10 bool
		_ = _out10
		var _out11 _dafny.Map
		_ = _out11
		_out8, _out9, _out10, _out11 = Companion_Default___.M0(_162_globalState)
		_180_v26 = _out8
		_181_v27 = _out9
		_182_v28 = _out10
		_183_v29 = _out11
		_155_v4 = Companion_Default___.Fm0((_152_v1).Minus(_152_v1), _180_v26, _dafny.Companion_Sequence_.Concatenate(_166_v14, _166_v14), _167_v15, _162_globalState)
		_161_v10 = (func() _dafny.Array {
			if _181_v27 {
				return _161_v10
			}
			return _161_v10
		})()
		if Companion_Default___.Fm0(_152_v1, _152_v1, _166_v14, (_167_v15).Merge(_167_v15), _162_globalState) {
			var _184_v30 _dafny.Array
			_ = _184_v30
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_3
			var _nw23 _dafny.Array
			_ = _nw23
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw23 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Map = (func(_185_v1 _dafny.Int, _186_v28 bool) func(_dafny.Int) _dafny.Map {
					return func(_187_i1 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v1, !(_186_v28))
					}
				})(_152_v1, _182_v28)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw23 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw23).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw23).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_184_v30 = _nw23
			_184_v30 = _184_v30
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_161_v10), 0))
			_ = _index28
			(_161_v10).ArraySet1(_180_v26, (_index28).Int())
			var _188_v31 _dafny.Array
			_ = _188_v31
			var _nwElement0_4 _dafny.Int = _159_v8
			_ = _nwElement0_4
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(21))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_4, 0)
			(_nw24).ArraySet1(_159_v8, 1)
			(_nw24).ArraySet1(_180_v26, 2)
			(_nw24).ArraySet1(_159_v8, 3)
			(_nw24).ArraySet1(_dafny.IntOfInt64(386), 4)
			(_nw24).ArraySet1(_159_v8, 5)
			(_nw24).ArraySet1(Companion_Default___.Fm1(_181_v27, _155_v4, _162_globalState), 6)
			(_nw24).ArraySet1(_dafny.IntOfInt64(911), 7)
			(_nw24).ArraySet1(_159_v8, 8)
			(_nw24).ArraySet1(_159_v8, 9)
			(_nw24).ArraySet1(_159_v8, 10)
			(_nw24).ArraySet1(_159_v8, 11)
			(_nw24).ArraySet1(_159_v8, 12)
			(_nw24).ArraySet1(_159_v8, 13)
			(_nw24).ArraySet1(_159_v8, 14)
			(_nw24).ArraySet1(_159_v8, 15)
			(_nw24).ArraySet1(_159_v8, 16)
			(_nw24).ArraySet1(_159_v8, 17)
			(_nw24).ArraySet1(_159_v8, 18)
			(_nw24).ArraySet1(_152_v1, 19)
			(_nw24).ArraySet1((_183_v29).Cardinality(), 20)
			_188_v31 = _nw24
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_188_v31), 0))
			_ = _index29
			(_188_v31).ArraySet1(_159_v8, (_index29).Int())
			var _189_v32 _dafny.Array
			_ = _189_v32
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw25
			_189_v32 = _nw25
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_189_v32), 0))
			_ = _index30
			(_189_v32).ArraySet1(_181_v27, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_161_v10), 0))
			_ = _index31
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_188_v31), 0))
			_ = _index32
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_189_v32), 0))
			_ = _index33
			var _rhs44 _dafny.Int = (_180_v26).Times(_180_v26)
			_ = _rhs44
			var _rhs45 _dafny.Int = _159_v8
			_ = _rhs45
			var _rhs46 bool = Companion_Default___.Fm0((_156_v5).Cardinality(), _152_v1, _166_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v9, _155_v4), _162_globalState)
			_ = _rhs46
			var _rhs47 _dafny.Int = _180_v26
			_ = _rhs47
			var _lhs31 _dafny.Array = _161_v10
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_161_v10), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _188_v31
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_188_v31), 0))
			_ = _lhs34
			var _lhs35 _dafny.Array = _189_v32
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_189_v32), 0))
			_ = _lhs36
			(_lhs31).ArraySet1(_rhs44, (_lhs32).Int())
			(_lhs33).ArraySet1(_rhs45, (_lhs34).Int())
			(_lhs35).ArraySet1(_rhs46, (_lhs36).Int())
			_152_v1 = _rhs47
			var _190_v33 _dafny.Sequence
			_ = _190_v33
			_190_v33 = _dafny.SeqOf(_160_v9, _160_v9)
			var _191_v34 _dafny.MultiSet
			_ = _191_v34
			_191_v34 = _dafny.MultiSetOf(_160_v9)
			_182_v28 = (_191_v34).IsSubsetOf(_dafny.MultiSetFromSeq(_190_v33))
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_189_v32), 0))
			_ = _index34
			(_189_v32).ArraySet1(_182_v28, (_index34).Int())
			var _192_v35 _dafny.Int
			_ = _192_v35
			var _193_v36 bool
			_ = _193_v36
			var _194_v37 bool
			_ = _194_v37
			var _195_v38 _dafny.Map
			_ = _195_v38
			var _out12 _dafny.Int
			_ = _out12
			var _out13 bool
			_ = _out13
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Map
			_ = _out15
			_out12, _out13, _out14, _out15 = Companion_Default___.M0(_162_globalState)
			_192_v35 = _out12
			_193_v36 = _out13
			_194_v37 = _out14
			_195_v38 = _out15
		} else {
			_182_v28 = _155_v4
			(_162_globalState).F2 = _180_v26
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_161_v10), 0))
			_ = _index35
			(_161_v10).ArraySet1(_152_v1, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_161_v10), 0))
			_ = _index36
			(_161_v10).ArraySet1((_152_v1).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_160_v9, _160_v9)).Cardinality())), (_index36).Int())
			_182_v28 = (func() bool {
				if false {
					return _182_v28
				}
				return _182_v28
			})()
			var _196_v39 _dafny.CodePoint
			_ = _196_v39
			_196_v39 = _dafny.CodePoint('r')
			var _197_v40 _dafny.Map
			_ = _197_v40
			_197_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_196_v39, _181_v27)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_161_v10), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_161_v10), 0))
			_ = _index38
			var _rhs48 bool = _155_v4
			_ = _rhs48
			var _rhs49 _dafny.Int = _dafny.IntOfInt64(137)
			_ = _rhs49
			var _rhs50 bool = ((_dafny.IntOfInt64(-485)).Times(_180_v26)).Cmp(_180_v26) <= 0
			_ = _rhs50
			var _rhs51 bool = true
			_ = _rhs51
			var _rhs52 _dafny.Int = ((_197_v40).Cardinality()).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality())))
			_ = _rhs52
			var _lhs37 _dafny.Array = _161_v10
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_161_v10), 0))
			_ = _lhs38
			var _lhs39 _dafny.Array = _161_v10
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_161_v10), 0))
			_ = _lhs40
			_182_v28 = _rhs48
			(_lhs37).ArraySet1(_rhs49, (_lhs38).Int())
			_155_v4 = _rhs50
			_182_v28 = _rhs51
			(_lhs39).ArraySet1(_rhs52, (_lhs40).Int())
		}
		_181_v27 = !(_155_v4) || ((_180_v26).Cmp(_152_v1) <= 0)
	}
	var _hi0 _dafny.Int = _152_v1
	_ = _hi0
	for _198_i2 := _152_v1; _198_i2.Cmp(_hi0) < 0; _198_i2 = _198_i2.Plus(_dafny.One) {
		var _source5 _dafny.Int = _159_v8
		_ = _source5
		var _199___mcc_h0 _dafny.Int = _source5
		_ = _199___mcc_h0
		var _200_cf0 _dafny.Int = _199___mcc_h0
		_ = _200_cf0
		_155_v4 = (false) || (_155_v4)
		var _201_v41 _dafny.Array
		_ = _201_v41
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw26
		_201_v41 = _nw26
		var _202_v42 _dafny.Sequence
		_ = _202_v42
		_202_v42 = _dafny.SeqOf(_161_v10)
		var _203_v43 _dafny.Map
		_ = _203_v43
		_203_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v4, _155_v4)
		var _rhs53 _dafny.Array = (_202_v42).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_157_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm2(!(_155_v4), _200_cf0, (_203_v43).Cardinality(), _155_v4, _162_globalState)).Cardinality()), _dafny.IntOfUint32((_157_v6).Cardinality()))).Uint32()).(bool), false)).Cardinality(), _dafny.IntOfUint32((_202_v42).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs53
		var _rhs54 _dafny.Array = _201_v41
		_ = _rhs54
		var _rhs55 bool = false
		_ = _rhs55
		var _lhs41 *GlobalState = _162_globalState
		_ = _lhs41
		_lhs41.F4 = _rhs53
		_201_v41 = _rhs54
		_155_v4 = _rhs55
		_203_v43 = (_203_v43).Update(_155_v4, _155_v4)
		_160_v9 = _160_v9
		var _204_v44 _dafny.Set
		_ = _204_v44
		_204_v44 = _dafny.SetOf(_160_v9)
		_204_v44 = (_204_v44).Difference(_204_v44)
		(_162_globalState).F2 = (_152_v1).Plus(Companion_Default___.SafeDivisionInt(_198_i2, _dafny.IntOfInt64(566)))
		var _205_v45 *C0
		_ = _205_v45
		var _nw27 *C0 = New_C0_()
		_ = _nw27
		_nw27.Ctor__(false, _152_v1)
		_205_v45 = _nw27
		var _206_v46 _dafny.MultiSet
		_ = _206_v46
		_206_v46 = _dafny.MultiSetOf(_205_v45, _205_v45)
		var _207_v47 _dafny.Array
		_ = _207_v47
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw28
		_207_v47 = _nw28
		var _208_v48 *C1
		_ = _208_v48
		var _nw29 *C1 = New_C1_()
		_ = _nw29
		_nw29.Ctor__((_206_v46).IsSubsetOf(_206_v46), Companion_D3_.Create_DC5_(_207_v47))
		_208_v48 = _nw29
	}
	var _209_v49 _dafny.Array
	_ = _209_v49
	var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
	_ = _nw30
	_209_v49 = _nw30
	var _210_v50 D3
	_ = _210_v50
	_210_v50 = Companion_D3_.Create_DC5_(_209_v49)
	var _211_v51 *C1
	_ = _211_v51
	var _nw31 *C1 = New_C1_()
	_ = _nw31
	_nw31.Ctor__(_155_v4, (func() D3 {
		if _155_v4 {
			return _210_v50
		}
		return _210_v50
	})())
	_211_v51 = _nw31
	if _155_v4 {
		_160_v9 = _160_v9
		var _212_v52 _dafny.Int
		_ = _212_v52
		var _213_v53 bool
		_ = _213_v53
		var _214_v54 bool
		_ = _214_v54
		var _215_v55 _dafny.Map
		_ = _215_v55
		var _out16 _dafny.Int
		_ = _out16
		var _out17 bool
		_ = _out17
		var _out18 bool
		_ = _out18
		var _out19 _dafny.Map
		_ = _out19
		_out16, _out17, _out18, _out19 = Companion_Default___.M0(_162_globalState)
		_212_v52 = _out16
		_213_v53 = _out17
		_214_v54 = _out18
		_215_v55 = _out19
		var _216_v56 _dafny.Array
		_ = _216_v56
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
		_ = _nw32
		_216_v56 = _nw32
		_216_v56 = _216_v56
		_211_v51 = _211_v51
		_214_v54 = false
	} else {
		var _217_v57 _dafny.Array
		_ = _217_v57
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
		_ = _nw33
		_217_v57 = _nw33
		_217_v57 = _217_v57
		_211_v51 = _211_v51
		var _218_v58 *C2
		_ = _218_v58
		var _nw34 *C2 = New_C2_()
		_ = _nw34
		_nw34.Ctor__((_158_v7).IsSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(414), _152_v1, _152_v1, _152_v1, _152_v1)), true)
		_218_v58 = _nw34
		_218_v58 = _218_v58
		var _219_v59 _dafny.CodePoint
		_ = _219_v59
		_219_v59 = _dafny.CodePoint('q')
		var _220_v60 _dafny.Map
		_ = _220_v60
		_220_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v59, _152_v1)
		var _221_v61 _dafny.Map
		_ = _221_v61
		_221_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_220_v60).Merge(_220_v60), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-347), _152_v1))
		_221_v61 = (_221_v61).Merge(_221_v61)
		var _222_v62 _dafny.Array
		_ = _222_v62
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(10))
		_ = _nw35
		_222_v62 = _nw35
		var _223_v63 D3
		_ = _223_v63
		_223_v63 = Companion_D3_.Create_DC8_((_211_v51).F16(), _152_v1, _152_v1)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_222_v62), 0))
		_ = _index39
		(_222_v62).ArraySet1(_223_v63, (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_222_v62), 0))
		_ = _index40
		(_222_v62).ArraySet1(_223_v63, (_index40).Int())
	}
	var _hi1 _dafny.Int = Companion_Default___.SafeModuloInt(_152_v1, _152_v1)
	_ = _hi1
	for _224_i3 := Companion_Default___.Fm17(!(false), _152_v1, _155_v4, _152_v1, _162_globalState); _224_i3.Cmp(_hi1) < 0; _224_i3 = _224_i3.Plus(_dafny.One) {
		var _225_v64 _dafny.Map
		_ = _225_v64
		_225_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_152_v1), _152_v1)
		var _226_v65 _dafny.Set
		_ = _226_v65
		_226_v65 = _dafny.SetOf(_160_v9, _160_v9)
		_225_v64 = (_225_v64).Update(((_226_v65).Difference(_226_v65)).Cardinality(), (_224_i3).Minus((_225_v64).Cardinality()))
		var _227_v66 _dafny.Array
		_ = _227_v66
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
		_ = _nw36
		_227_v66 = _nw36
		var _228_v67 _dafny.Map
		_ = _228_v67
		_228_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v66, _224_i3)
		_228_v67 = (_228_v67).Update(_227_v66, _152_v1)
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_161_v10), 0))
		_ = _index41
		(_161_v10).ArraySet1(_224_i3, (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_161_v10), 0))
		_ = _index42
		(_161_v10).ArraySet1(_224_i3, (_index42).Int())
		var _rhs56 _dafny.Sequence = Companion_Default___.Fm2(_155_v4, (_dafny.Zero).Minus((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)), _152_v1, !(_155_v4), _162_globalState)
		_ = _rhs56
		var _rhs57 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(897), _224_i3)
		_ = _rhs57
		var _rhs58 _dafny.Int = (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)
		_ = _rhs58
		var _lhs42 *GlobalState = _162_globalState
		_ = _lhs42
		var _lhs43 *GlobalState = _162_globalState
		_ = _lhs43
		_160_v9 = _rhs56
		_lhs42.F2 = _rhs57
		_lhs43.F2 = _rhs58
	}
	if _155_v4 {
		var _229_v68 T0
		_ = _229_v68
		var _nw37 *C2 = New_C2_()
		_ = _nw37
		_nw37.Ctor__((_211_v51).F16(), (_211_v51).F16())
		_229_v68 = _nw37
		var _230_v69 D4
		_ = _230_v69
		_230_v69 = Companion_D4_.Create_DC10_(_229_v68, _dafny.IntOfInt64(84), _152_v1)
		var _source6 D4 = _230_v69
		_ = _source6
		if _source6.Is_DC10() {
			var _231___mcc_h1 T0 = _source6.Get_().(D4_DC10).Cf17
			_ = _231___mcc_h1
			var _232___mcc_h2 _dafny.Int = _source6.Get_().(D4_DC10).Cf18
			_ = _232___mcc_h2
			var _233___mcc_h3 _dafny.Int = _source6.Get_().(D4_DC10).Cf19
			_ = _233___mcc_h3
			var _234_cf19 _dafny.Int = _233___mcc_h3
			_ = _234_cf19
			var _235_cf18 _dafny.Int = _232___mcc_h2
			_ = _235_cf18
			var _236_cf17 T0 = _231___mcc_h1
			_ = _236_cf17
			_160_v9 = _160_v9
			var _237_v70 D8
			_ = _237_v70
			_237_v70 = Companion_D8_.Create_DC18_()
			var _238_v71 _dafny.Map
			_ = _238_v71
			_238_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_235_cf18), _237_v70)
			var _239_v72 _dafny.CodePoint
			_ = _239_v72
			_239_v72 = _dafny.CodePoint('a')
			var _rhs59 bool = (func() bool {
				if _dafny.Companion_Sequence_.IsProperPrefixOf((_dafny.UnicodeSeqOfUtf8Bytes("bonxg")), _dafny.Companion_Sequence_.Update(_160_v9, (Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32((_160_v9).Cardinality()))).Uint32(), _239_v72)) {
					return (_211_v51).F16()
				}
				return (_211_v51).F16()
			})()
			_ = _rhs59
			var _rhs60 _dafny.Map = _238_v71
			_ = _rhs60
			var _rhs61 bool = (_152_v1).Cmp((_234_cf19).Minus(_152_v1)) >= 0
			_ = _rhs61
			_155_v4 = _rhs59
			_238_v71 = _rhs60
			_155_v4 = _rhs61
			var _240_v73 _dafny.Array
			_ = _240_v73
			var _nwElement0_5 *C1 = _211_v51
			_ = _nwElement0_5
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(26))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_5, 0)
			(_nw38).ArraySet1(_211_v51, 1)
			(_nw38).ArraySet1(_211_v51, 2)
			(_nw38).ArraySet1(_211_v51, 3)
			(_nw38).ArraySet1(_211_v51, 4)
			(_nw38).ArraySet1(_211_v51, 5)
			(_nw38).ArraySet1(_211_v51, 6)
			(_nw38).ArraySet1(_211_v51, 7)
			(_nw38).ArraySet1(_211_v51, 8)
			(_nw38).ArraySet1(_211_v51, 9)
			(_nw38).ArraySet1(_211_v51, 10)
			(_nw38).ArraySet1(_211_v51, 11)
			(_nw38).ArraySet1(_211_v51, 12)
			(_nw38).ArraySet1(_211_v51, 13)
			(_nw38).ArraySet1(_211_v51, 14)
			(_nw38).ArraySet1(_211_v51, 15)
			(_nw38).ArraySet1(_211_v51, 16)
			(_nw38).ArraySet1(_211_v51, 17)
			(_nw38).ArraySet1(_211_v51, 18)
			(_nw38).ArraySet1(_211_v51, 19)
			(_nw38).ArraySet1(_211_v51, 20)
			(_nw38).ArraySet1(_211_v51, 21)
			(_nw38).ArraySet1(_211_v51, 22)
			(_nw38).ArraySet1(_211_v51, 23)
			(_nw38).ArraySet1(_211_v51, 24)
			(_nw38).ArraySet1((func() *C1 {
				if _155_v4 {
					return _211_v51
				}
				return _211_v51
			})(), 25)
			_240_v73 = _nw38
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_240_v73), 0))
			_ = _index43
			(_240_v73).ArraySet1(_211_v51, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_240_v73), 0))
			_ = _index44
			(_240_v73).ArraySet1(_211_v51, (_index44).Int())
			var _241_v74 _dafny.Map
			_ = _241_v74
			_241_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_152_v1), _235_cf18)
			_241_v74 = (_241_v74).Update((_152_v1).Plus(_152_v1), _152_v1)
		} else {
			var _242___mcc_h4 _dafny.MultiSet = _source6.Get_().(D4_DC9).Cf16
			_ = _242___mcc_h4
			var _243_cf16 _dafny.MultiSet = _242___mcc_h4
			_ = _243_cf16
			_155_v4 = (_211_v51).F16()
			(_162_globalState).F2 = _dafny.IntOfInt64(412)
			var _244_v75 _dafny.Int
			_ = _244_v75
			var _245_v76 _dafny.Int
			_ = _245_v76
			var _out20 _dafny.Int
			_ = _out20
			var _out21 _dafny.Int
			_ = _out21
			_out20, _out21 = (_211_v51).M5(_152_v1, _162_globalState)
			_244_v75 = _out20
			_245_v76 = _out21
			var _246_v77 _dafny.Int
			_ = _246_v77
			var _247_v78 _dafny.Int
			_ = _247_v78
			var _out22 _dafny.Int
			_ = _out22
			var _out23 _dafny.Int
			_ = _out23
			_out22, _out23 = (_211_v51).M5((_154_v3).Select((Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32()).(_dafny.Int), _162_globalState)
			_246_v77 = _out22
			_247_v78 = _out23
		}
		var _248_v79 _dafny.Array
		_ = _248_v79
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw39
		_248_v79 = _nw39
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
		_ = _index45
		(_248_v79).ArraySet1((func() bool {
			if (_211_v51).F16() {
				return (_211_v51).F16()
			}
			return (_211_v51).F16()
		})(), (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
		_ = _index46
		(_248_v79).ArraySet1(!((_211_v51).F16()), (_index46).Int())
		(_162_globalState).F2 = _152_v1
		var _source7 D11 = Companion_D11_.Create_DC26_()
		_ = _source7
		if _source7.Is_DC25() {
			var _249___mcc_h5 _dafny.CodePoint = _source7.Get_().(D11_DC25).Cf37
			_ = _249___mcc_h5
			var _250_cf37 _dafny.CodePoint = _249___mcc_h5
			_ = _250_cf37
			var _251_v80 D6
			_ = _251_v80
			_251_v80 = Companion_D6_.Create_DC13_((_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool))
			var _252_v81 *C1
			_ = _252_v81
			var _nw40 *C1 = New_C1_()
			_ = _nw40
			_nw40.Ctor__((_251_v80).Dtor_cf22(), _211_v51.F17)
			_252_v81 = _nw40
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index47
			(_248_v79).ArraySet1((_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index48
			(_248_v79).ArraySet1((_211_v51).F16(), (_index48).Int())
			var _253_v82 _dafny.Array
			_ = _253_v82
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(7))
			_ = _nw41
			_253_v82 = _nw41
			var _254_v83 _dafny.Map
			_ = _254_v83
			_254_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v3, _152_v1)
			var _255_v84 _dafny.MultiSet
			_ = _255_v84
			_255_v84 = _dafny.MultiSetOf(_250_cf37)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_253_v82), 0))
			_ = _index49
			(_253_v82).ArraySet1(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(((_254_v83).Update(_154_v3, (_255_v84).Cardinality())).Keys().Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _256_v85 _dafny.Sequence
					_256_v85 = interface{}(_compr_7).(_dafny.Sequence)
					if ((_254_v83).Update(_154_v3, (_255_v84).Cardinality())).Contains(_256_v85) {
						_coll7.Add(_256_v85)
					}
				}
				return _coll7.ToSet()
			}(), (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_253_v82), 0))
			_ = _index50
			(_253_v82).ArraySet1(_dafny.SetOf(_154_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_154_v3, (Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_257_cf37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_258_i4 _dafny.Int) _dafny.CodePoint {
					return _257_cf37
				}
			})(_250_cf37)))).Cardinality())), _154_v3), Companion_Default___.Fm25(_152_v1, _152_v1, _152_v1, (_158_v7).Cardinality(), _162_globalState)), (_index50).Int())
		} else if _source7.Is_DC26() {
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_161_v10), 0))
			_ = _index51
			(_161_v10).ArraySet1(_dafny.IntOfUint32((_160_v9).Cardinality()), (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_161_v10), 0))
			_ = _index52
			(_161_v10).ArraySet1(_152_v1, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index53
			(_248_v79).ArraySet1((_211_v51).F16(), (_index53).Int())
			var _259_v86 _dafny.CodePoint
			_ = _259_v86
			_259_v86 = _dafny.CodePoint('b')
			_155_v4 = (func() bool {
				if (_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool) {
					return !(!((_211_v51).Fm4(_152_v1, (_211_v51).F16(), _259_v86, _162_globalState)))
				}
				return (_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool)
			})()
			var _260_v87 _dafny.Array
			_ = _260_v87
			var _nwElement0_6 _dafny.Sequence = _160_v9
			_ = _nwElement0_6
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_6, 0)
			_260_v87 = _nw42
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_260_v87), 0))
			_ = _index54
			(_260_v87).ArraySet1(_160_v9, (_index54).Int())
			var _261_v88 *C2
			_ = _261_v88
			var _nw43 *C2 = New_C2_()
			_ = _nw43
			_nw43.Ctor__((_211_v51).F16(), (_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool))
			_261_v88 = _nw43
			var _262_v89 _dafny.Map
			_ = _262_v89
			_262_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v88, _261_v88.F14)
			var _263_v90 _dafny.Sequence
			_ = _263_v90
			_263_v90 = _160_v9
			var _264_v91 D12
			_ = _264_v91
			_264_v91 = Companion_D12_.Create_DC28_(_262_v89)
			var _265_v92 _dafny.Map
			_ = _265_v92
			_265_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _211_v51)
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_260_v87), 0))
			_ = _index55
			var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_160_v9, (_263_v90)), _160_v9)
			_ = _rhs62
			var _rhs63 _dafny.Map = (_264_v91).Dtor_cf39()
			_ = _rhs63
			var _rhs64 _dafny.Int = (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)
			_ = _rhs64
			var _rhs65 *C1 = (func() *C1 {
				if _155_v4 {
					return _211_v51
				}
				return (func() *C1 {
					if (_265_v92).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_157_v6).Cardinality())))) {
						return (_265_v92).Get((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_157_v6).Cardinality())))).(*C1)
					}
					return _211_v51
				})()
			})()
			_ = _rhs65
			var _rhs66 _dafny.CodePoint = _259_v86
			_ = _rhs66
			var _lhs44 _dafny.Array = _260_v87
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_260_v87), 0))
			_ = _lhs45
			var _lhs46 *GlobalState = _162_globalState
			_ = _lhs46
			(_lhs44).ArraySet1(_rhs62, (_lhs45).Int())
			_262_v89 = _rhs63
			_lhs46.F2 = _rhs64
			_211_v51 = _rhs65
			_259_v86 = _rhs66
		} else if _source7.Is_DC24() {
			var _266___mcc_h6 _dafny.Array = _source7.Get_().(D11_DC24).Cf36
			_ = _266___mcc_h6
			var _267_cf36 _dafny.Array = _266___mcc_h6
			_ = _267_cf36
			var _268_v93 _dafny.Set
			_ = _268_v93
			_268_v93 = _dafny.SetOf((_211_v51).F16())
			var _269_v94 _dafny.Map
			_ = _269_v94
			_269_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _152_v1)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index56
			(_248_v79).ArraySet1(((Companion_Default___.Fm23(_269_v94, (_211_v51).F16(), _162_globalState)).Union(_dafny.SetOf(true, (_211_v51).F16()))).IsProperSubsetOf(_268_v93), (_index56).Int())
			var _270_v95 _dafny.Map
			_ = _270_v95
			_270_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _152_v1)
			_270_v95 = (_270_v95).Update((func() bool {
				if (_211_v51).F16() {
					return _155_v4
				}
				return (_211_v51).F16()
			})(), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_152_v1)), _152_v1))
			(_162_globalState).F2 = (_156_v5).Cardinality()
			(_162_globalState).F2 = _152_v1
		} else {
			var _271___mcc_h7 D11 = _source7.Get_().(D11_DC27).Cf38
			_ = _271___mcc_h7
			var _272_cf38 D11 = _271___mcc_h7
			_ = _272_cf38
			var _273_v96 _dafny.CodePoint
			_ = _273_v96
			_273_v96 = _dafny.CodePoint('y')
			_273_v96 = _273_v96
			var _274_v97 T1
			_ = _274_v97
			var _nw44 *C3 = New_C3_()
			_ = _nw44
			_nw44.Ctor__((func() _dafny.Int {
				if (_211_v51).F16() {
					return _152_v1
				}
				return (_152_v1)
			})(), Companion_Default___.Fm27(_162_globalState))
			_274_v97 = _nw44
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index57
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index58
			var _rhs67 bool = (Companion_Default___.Fm19(_274_v97.F10(), _162_globalState)).Cmp(_274_v97.F10()) == 0
			_ = _rhs67
			var _rhs68 bool = (_274_v97.F10()).Cmp(_152_v1) >= 0
			_ = _rhs68
			var _rhs69 T1 = _274_v97
			_ = _rhs69
			var _lhs47 _dafny.Array = _248_v79
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _lhs48
			var _lhs49 _dafny.Array = _248_v79
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _lhs50
			(_lhs47).ArraySet1(_rhs67, (_lhs48).Int())
			(_lhs49).ArraySet1(_rhs68, (_lhs50).Int())
			_274_v97 = _rhs69
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
			_ = _index59
			(_248_v79).ArraySet1((_211_v51).F16(), (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_161_v10), 0))
			_ = _index60
			(_161_v10).ArraySet1(_152_v1, (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_161_v10), 0))
			_ = _index61
			(_161_v10).ArraySet1(_274_v97.F10(), (_index61).Int())
		}
		var _275_v98 _dafny.Map
		_ = _275_v98
		_275_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_154_v3, (Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32(), _152_v1)).Cardinality()))
		var _276_v99 _dafny.CodePoint
		_ = _276_v99
		_276_v99 = _dafny.CodePoint('e')
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
		_ = _index62
		var _rhs70 bool = (_152_v1).Cmp(_152_v1) <= 0
		_ = _rhs70
		var _rhs71 bool = (_248_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))).Int()).(bool)
		_ = _rhs71
		var _rhs72 _dafny.Sequence = _160_v9
		_ = _rhs72
		var _rhs73 bool = (_229_v68).Fm4((_dafny.Zero).Minus(_152_v1), _dafny.Companion_Sequence_.Equal(_160_v9, _160_v9), _276_v99, _162_globalState)
		_ = _rhs73
		var _rhs74 _dafny.Map = _275_v98
		_ = _rhs74
		var _lhs51 _dafny.Array = _248_v79
		_ = _lhs51
		var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_248_v79), 0))
		_ = _lhs52
		_155_v4 = _rhs70
		_155_v4 = _rhs71
		_160_v9 = _rhs72
		(_lhs51).ArraySet1(_rhs73, (_lhs52).Int())
		_275_v98 = _rhs74
	} else {
		var _277_v100 _dafny.Array
		_ = _277_v100
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw45
		_277_v100 = _nw45
		var _278_v101 D12
		_ = _278_v101
		_278_v101 = Companion_D12_.Create_DC29_(_277_v100, _155_v4)
		_155_v4 = (_278_v101).Dtor_cf41()
		var _279_v102 T0
		_ = _279_v102
		var _nw46 *C2 = New_C2_()
		_ = _nw46
		_nw46.Ctor__((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_154_v3, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_160_v9).Cardinality()), _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32(), _152_v1)).Cardinality())).Cmp(_152_v1) >= 0, (_152_v1).Cmp(_152_v1) <= 0)
		_279_v102 = _nw46
		var _280_v103 _dafny.Map
		_ = _280_v103
		_280_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v1, _156_v5)
		var _rhs75 T0 = _279_v102
		_ = _rhs75
		var _rhs76 _dafny.Map = Companion_Default___.Fm28(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_157_v6, _dafny.SeqOf((_211_v51).F16(), Companion_Default___.Fm0(_dafny.IntOfInt64(891), _152_v1, _166_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v9, (_211_v51).F16()), _162_globalState)))).Cardinality()), true, _162_globalState)
		_ = _rhs76
		_279_v102 = _rhs75
		_280_v103 = _rhs76
		_155_v4 = (Companion_D1_.Create_DC1_(_155_v4)).Dtor_cf1()
		var _281_v104 T1
		_ = _281_v104
		var _nw47 *C3 = New_C3_()
		_ = _nw47
		_nw47.Ctor__(_dafny.IntOfUint32((_154_v3).Cardinality()), _156_v5)
		_281_v104 = _nw47
		var _282_v105 *C2
		_ = _282_v105
		var _nw48 *C2 = New_C2_()
		_ = _nw48
		_nw48.Ctor__(_155_v4, (_157_v6).Select((Companion_Default___.SafeIndex(_281_v104.F10(), _dafny.IntOfUint32((_157_v6).Cardinality()))).Uint32()).(bool))
		_282_v105 = _nw48
		var _283_v106 _dafny.Map
		_ = _283_v106
		_283_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v104, _282_v105)
		var _284_v107 _dafny.Sequence
		_ = _284_v107
		_284_v107 = _dafny.SeqOf(_283_v106)
		var _285_v108 *C0
		_ = _285_v108
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__((_211_v51).F16(), _dafny.IntOfUint32((_284_v107).Cardinality()))
		_285_v108 = _nw49
		if !((_211_v51).F16()) {
			var _rhs77 bool = (func() bool {
				if (_211_v51).F16() {
					return (_211_v51).F16()
				}
				return _282_v105.F14
			})()
			_ = _rhs77
			var _rhs78 _dafny.Int = ((_dafny.IntOfInt64(771)).Times(_dafny.IntOfInt64(265))).Minus((_285_v108).F13())
			_ = _rhs78
			var _rhs79 _dafny.Array = _277_v100
			_ = _rhs79
			var _rhs80 *C0 = _285_v108
			_ = _rhs80
			var _lhs53 *C2 = _282_v105
			_ = _lhs53
			var _lhs54 *GlobalState = _162_globalState
			_ = _lhs54
			_lhs53.F14 = _rhs77
			_lhs54.F2 = _rhs78
			_277_v100 = _rhs79
			_285_v108 = _rhs80
			var _286_v109 _dafny.CodePoint
			_ = _286_v109
			_286_v109 = _dafny.CodePoint('s')
			var _rhs81 _dafny.Int = _281_v104.F10()
			_ = _rhs81
			var _rhs82 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_287_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), Companion_Default___.Fm2((_211_v51).F16(), (_285_v108).F13(), (_dafny.Zero).Minus((_285_v108).F13()), _282_v105.F14, _162_globalState)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_157_v6).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_288_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), Companion_Default___.Fm2((_211_v51).F16(), (_285_v108).F13(), (_dafny.Zero).Minus((_285_v108).F13()), _282_v105.F14, _162_globalState))).Cardinality()))).Uint32(), _286_v109)
			_ = _rhs82
			var _rhs83 _dafny.Sequence = _160_v9
			_ = _rhs83
			var _lhs55 *GlobalState = _162_globalState
			_ = _lhs55
			_lhs55.F2 = _rhs81
			_160_v9 = _rhs82
			_160_v9 = _rhs83
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_277_v100), 0))
			_ = _index63
			(_277_v100).ArraySet1(!(true), (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_277_v100), 0))
			_ = _index64
			(_277_v100).ArraySet1((_282_v105).F15(), (_index64).Int())
			var _289_v110 *C3
			_ = _289_v110
			var _nw50 *C3 = New_C3_()
			_ = _nw50
			_nw50.Ctor__((_154_v3).Select((Companion_Default___.SafeIndex((_285_v108).F13(), _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32()).(_dafny.Int), ((_281_v104).F11()).Update((_285_v108).F13(), _155_v4))
			_289_v110 = _nw50
			_289_v110 = _289_v110
			var _290_v111 *C0
			_ = _290_v111
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__(!((_285_v108).F12()), _281_v104.F10())
			_290_v111 = _nw51
		} else {
			(_282_v105).F14 = (_282_v105).F15()
			var _291_v112 D10
			_ = _291_v112
			_291_v112 = Companion_D10_.Create_DC23_((_285_v108).F13(), _282_v105.F14, (_211_v51).F16())
			var _292_v113 _dafny.Map
			_ = _292_v113
			_292_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_157_v6).Cardinality())).Cmp((_291_v112).Dtor_cf33()) >= 0, (_282_v105.F14) && ((_285_v108).F12()))
			_292_v113 = (_292_v113).Merge(Companion_Default___.Fm22(_162_globalState))
			(_162_globalState).F2 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_152_v1, _152_v1))
			(_281_v104).F10_set_((_285_v108).F13())
			var _293_v114 _dafny.Int
			_ = _293_v114
			var _294_v115 bool
			_ = _294_v115
			var _295_v116 bool
			_ = _295_v116
			var _296_v117 _dafny.Map
			_ = _296_v117
			var _out24 _dafny.Int
			_ = _out24
			var _out25 bool
			_ = _out25
			var _out26 bool
			_ = _out26
			var _out27 _dafny.Map
			_ = _out27
			_out24, _out25, _out26, _out27 = Companion_Default___.M0(_162_globalState)
			_293_v114 = _out24
			_294_v115 = _out25
			_295_v116 = _out26
			_296_v117 = _out27
		}
	}
	var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
	_ = _index65
	(_161_v10).ArraySet1(_dafny.IntOfInt64(192), (_index65).Int())
	var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
	_ = _index66
	(_161_v10).ArraySet1(_152_v1, (_index66).Int())
	var _297_v118 *C3
	_ = _297_v118
	var _nw52 *C3 = New_C3_()
	_ = _nw52
	_nw52.Ctor__(_152_v1, _156_v5)
	_297_v118 = _nw52
	var _298_v119 _dafny.Map
	_ = _298_v119
	_298_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _297_v118)
	var _299_v120 _dafny.CodePoint
	_ = _299_v120
	_299_v120 = _dafny.CodePoint('j')
	_298_v119 = (_298_v119).Update(_299_v120, _297_v118)
	var _300_v121 _dafny.Array
	_ = _300_v121
	var _nwElement0_7 bool = _155_v4
	_ = _nwElement0_7
	var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(5))
	_ = _nw53
	(_nw53).ArraySet1(_nwElement0_7, 0)
	(_nw53).ArraySet1((func() bool {
		if true {
			return _155_v4
		}
		return _155_v4
	})(), 1)
	(_nw53).ArraySet1((_211_v51).F16(), 2)
	(_nw53).ArraySet1((_211_v51).F16(), 3)
	(_nw53).ArraySet1((_155_v4) || (_155_v4), 4)
	_300_v121 = _nw53
	var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))
	_ = _index67
	(_300_v121).ArraySet1(true, (_index67).Int())
	var _301_v122 D9
	_ = _301_v122
	_301_v122 = Companion_D9_.Create_DC20_(_155_v4, _299_v120)
	var _302_v123 _dafny.MultiSet
	_ = _302_v123
	_302_v123 = _dafny.MultiSetOf(_301_v122, _301_v122, _301_v122, _301_v122)
	var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))
	_ = _index68
	(_300_v121).ArraySet1((_302_v123).IsDisjointFrom(_302_v123), (_index68).Int())
	var _303_i6 _dafny.Int
	_ = _303_i6
	_303_i6 = _dafny.Zero
	{
		for _155_v4 {
			{
				if (_303_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_303_i6 = (_303_i6).Plus(_dafny.One)
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
				_ = _index69
				(_161_v10).ArraySet1(_dafny.IntOfUint32((_160_v9).Cardinality()), (_index69).Int())
				var _304_v125 _dafny.MultiSet
				_ = _304_v125
				_304_v125 = _dafny.MultiSetOf(_299_v120, _299_v120)
				var _305_v126 _dafny.Map
				_ = _305_v126
				_305_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_211_v51).F16(), Companion_Default___.Fm29(_152_v1, (func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate((_304_v125).Elements()); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _306_v124 _dafny.CodePoint
						_306_v124 = interface{}(_compr_8).(_dafny.CodePoint)
						if (_304_v125).Contains(_306_v124) {
							_coll8.Add(_306_v124, _152_v1)
						}
					}
					return _coll8.ToMap()
				}()).Update(_299_v120, _dafny.IntOfUint32((_154_v3).Cardinality())), _dafny.IntOfUint32((_160_v9).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_211_v51).F16()))).Cardinality(), _162_globalState))
				var _307_v127 D3
				_ = _307_v127
				_307_v127 = Companion_D3_.Create_DC7_((_156_v5).Cardinality())
				var _308_v128 _dafny.Map
				_ = _308_v128
				_308_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_297_v118).Fm7(_305_v126, true, _152_v1, _162_globalState), _307_v127)
				var _309_v129 _dafny.Map
				_ = _309_v129
				_309_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_211_v51).F16(), _dafny.IntOfUint32((_160_v9).Cardinality()))
				var _310_v130 D4
				_ = _310_v130
				_310_v130 = Companion_D4_.Create_DC9_(Companion_Default___.Fm14(_308_v128, _155_v4, _309_v129, (_300_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))).Int()).(bool), _162_globalState))
				_310_v130 = _310_v130
				var _311_v131 _dafny.Set
				_ = _311_v131
				_311_v131 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_157_v6, _157_v6), _157_v6)
				var _312_v132 _dafny.Sequence
				_ = _312_v132
				_312_v132 = _dafny.SeqOf(_309_v129)
				var _313_v133 _dafny.Map
				_ = _313_v133
				_313_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v4, _dafny.Companion_Sequence_.Update(_157_v6, (Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32((_157_v6).Cardinality()))).Uint32(), (_300_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))).Int()).(bool)))
				var _314_v134 _dafny.Sequence
				_ = _314_v134
				_314_v134 = _dafny.SeqOf((func() _dafny.Sequence {
					if (_313_v133).Contains(false) {
						return (_313_v133).Get(false).(_dafny.Sequence)
					}
					return _157_v6
				})(), _157_v6)
				var _315_v136 _dafny.Sequence
				_ = _315_v136
				_315_v136 = _dafny.SeqOf((_dafny.SetOf(_dafny.SeqOf((_211_v51).F16()), _dafny.SeqOf(_155_v4, (_211_v51).F16()), _157_v6, _dafny.SeqOf((_297_v118).Fm5(_299_v120, _312_v132, _155_v4, _162_globalState)), _157_v6)).Union(_311_v131), ((Companion_D13_.Create_DC31_(_311_v131)).Dtor_cf43()).Intersection(func() _dafny.Set {
					var _coll9 = _dafny.NewBuilder()
					_ = _coll9
					for _iter9 := _dafny.Iterate((_314_v134).Elements()); ; {
						_compr_9, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _316_v135 _dafny.Sequence
						_316_v135 = interface{}(_compr_9).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_314_v134, _316_v135) {
							_coll9.Add(_316_v135)
						}
					}
					return _coll9.ToSet()
				}()), _311_v131, _311_v131)
				var _317_v137 _dafny.Map
				_ = _317_v137
				_317_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int))
				var _318_v138 _dafny.Map
				_ = _318_v138
				_318_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_211_v51).F16(), _317_v137)
				_311_v131 = (_315_v136).Select((Companion_Default___.SafeIndex(((_318_v138).Merge(_318_v138)).Cardinality(), _dafny.IntOfUint32((_315_v136).Cardinality()))).Uint32()).(_dafny.Set)
				var _319_v139 _dafny.MultiSet
				_ = _319_v139
				_319_v139 = _dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_154_v3).Select((Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32()).(_dafny.Int), _299_v120)).Cardinality()), _153_v2, _153_v2)
				_319_v139 = _319_v139
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))
	_ = _index70
	(_300_v121).ArraySet1(true, (_index70).Int())
	var _320_v140 _dafny.Array
	_ = _320_v140
	var _nw54 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
	_ = _nw54
	_320_v140 = _nw54
	var _321_v141 _dafny.Sequence
	_ = _321_v141
	_321_v141 = _dafny.SeqOf(_320_v140)
	if (_dafny.IntOfUint32((_321_v141).Cardinality())).Cmp(_152_v1) <= 0 {
		var _322_v142 D14
		_ = _322_v142
		_322_v142 = Companion_D14_.Create_DC35_(_dafny.SeqOf(_211_v51))
		var _323_v143 _dafny.Sequence
		_ = _323_v143
		_323_v143 = _dafny.SeqOf(_211_v51)
		var _pat_let_tv0 = _323_v143
		_ = _pat_let_tv0
		var _pat_let_tv1 = _323_v143
		_ = _pat_let_tv1
		_155_v4 = !(((_158_v7).Difference(_158_v7)).IsSubsetOf((_158_v7).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func(_pat_let0_0 D14) D14 {
			return func(_324_dt__update__tmp_h5 D14) D14 {
				return func(_pat_let1_0 _dafny.Sequence) D14 {
					return func(_325_dt__update_hcf50_h0 _dafny.Sequence) D14 {
						return Companion_D14_.Create_DC35_(_325_dt__update_hcf50_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_322_v142)).Dtor_cf50(), (Companion_Default___.SafeIndex(_152_v1, _dafny.IntOfUint32(((func(_pat_let2_0 D14) D14 {
			return func(_326_dt__update__tmp_h6 D14) D14 {
				return func(_pat_let3_0 _dafny.Sequence) D14 {
					return func(_327_dt__update_hcf50_h1 _dafny.Sequence) D14 {
						return Companion_D14_.Create_DC35_(_327_dt__update_hcf50_h1)
					}(_pat_let3_0)
				}(_pat_let_tv1)
			}(_pat_let2_0)
		}(_322_v142)).Dtor_cf50()).Cardinality()))).Uint32(), _211_v51)).Cardinality())))))
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
		_ = _index71
		(_161_v10).ArraySet1((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(63), _dafny.IntOfInt64(992))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _328_v144 _dafny.Int
				_328_v144 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(63)).Cmp(_328_v144) <= 0) && ((_328_v144).Cmp(_dafny.IntOfInt64(992)) < 0) {
					_coll10.Add(Companion_Default___.SafeModuloInt(_328_v144, _dafny.IntOfUint32((_160_v9).Cardinality())))
				}
			}
			return _coll10.ToSet()
		}()).Cardinality(), (_index71).Int())
		var _329_v145 D11
		_ = _329_v145
		_329_v145 = Companion_D11_.Create_DC26_()
		_329_v145 = _329_v145
		var _330_v146 _dafny.Int
		_ = _330_v146
		var _331_v147 _dafny.Int
		_ = _331_v147
		var _out28 _dafny.Int
		_ = _out28
		var _out29 _dafny.Int
		_ = _out29
		_out28, _out29 = (_211_v51).M5((_dafny.Zero).Minus((_152_v1).Minus(_dafny.IntOfUint32((_160_v9).Cardinality()))), _162_globalState)
		_330_v146 = _out28
		_331_v147 = _out29
		if (_211_v51).F16() {
			var _332_v148 _dafny.Array
			_ = _332_v148
			var _nw55 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
			_ = _nw55
			_332_v148 = _nw55
			_332_v148 = (func() _dafny.Array {
				if true {
					return _332_v148
				}
				return _332_v148
			})()
			var _rhs84 _dafny.Int = _331_v147
			_ = _rhs84
			var _rhs85 bool = !((_211_v51).F16())
			_ = _rhs85
			var _rhs86 _dafny.Int = (_dafny.Zero).Minus((_152_v1).Minus(_152_v1))
			_ = _rhs86
			_331_v147 = _rhs84
			_155_v4 = _rhs85
			_152_v1 = _rhs86
			var _333_v149 _dafny.Sequence
			_ = _333_v149
			_333_v149 = _dafny.SeqOf(_297_v118)
			var _334_v150 _dafny.Map
			_ = _334_v150
			_334_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), (_333_v149).Select((Companion_Default___.SafeIndex((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_333_v149).Cardinality()))).Uint32()).(*C3))
			_334_v150 = (_334_v150).Update(Companion_Default___.Fm10((_211_v51).F16(), _162_globalState), _297_v118)
			_155_v4 = true
			var _335_v151 _dafny.Sequence
			_ = _335_v151
			_335_v151 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_160_v9, (Companion_Default___.SafeIndex(_330_v146, _dafny.IntOfUint32((_160_v9).Cardinality()))).Uint32(), _299_v120))
			_152_v1 = (func() _dafny.Int {
				if (_153_v2).Contains(_152_v1) {
					return (_153_v2).Multiplicity(_152_v1)
				}
				return _dafny.IntOfUint32(((_335_v151).Select((Companion_Default___.SafeIndex(_330_v146, _dafny.IntOfUint32((_335_v151).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			})()
		} else {
			var _336_v152 _dafny.Int
			_ = _336_v152
			var _337_v153 _dafny.Int
			_ = _337_v153
			var _out30 _dafny.Int
			_ = _out30
			var _out31 _dafny.Int
			_ = _out31
			_out30, _out31 = (_211_v51).M5((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _162_globalState)
			_336_v152 = _out30
			_337_v153 = _out31
			var _338_v154 T1
			_ = _338_v154
			var _nw56 *C3 = New_C3_()
			_ = _nw56
			_nw56.Ctor__(_337_v153, _156_v5)
			_338_v154 = _nw56
			var _339_v155 _dafny.MultiSet
			_ = _339_v155
			_339_v155 = _dafny.MultiSetOf(_338_v154)
			var _340_v156 D15
			_ = _340_v156
			_340_v156 = Companion_D15_.Create_DC38_(_338_v154)
			var _341_v157 _dafny.MultiSet
			_ = _341_v157
			_341_v157 = _dafny.MultiSetOf(_339_v155, _339_v155, _dafny.MultiSetFromSeq(_dafny.SeqOf((_340_v156).Dtor_cf57(), _338_v154, _338_v154)), _339_v155)
			_341_v157 = _341_v157
			var _342_v159 D7
			_ = _342_v159
			_342_v159 = Companion_D7_.Create_DC15_((func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(941), _dafny.IntOfInt64(43))); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _343_v158 _dafny.Int
					_343_v158 = interface{}(_compr_11).(_dafny.Int)
					if ((_dafny.IntOfInt64(941)).Cmp(_343_v158) <= 0) && ((_343_v158).Cmp(_dafny.IntOfInt64(43)) < 0) {
						_coll11.Add((_343_v158).Plus(_331_v147))
					}
				}
				return _coll11.ToSet()
			}()).Cardinality(), (_211_v51).F16())
			var _344_v160 T0
			_ = _344_v160
			var _nw57 *C2 = New_C2_()
			_ = _nw57
			_nw57.Ctor__((func() bool {
				if (_211_v51).F16() {
					return (_342_v159).Dtor_cf25()
				}
				return (_211_v51).F16()
			})(), !(_155_v4))
			_344_v160 = _nw57
			var _345_v161 _dafny.Map
			_ = _345_v161
			_345_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v4, (_211_v51).F16())
			var _346_v162 _dafny.Map
			_ = _346_v162
			_346_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_v121, !((func() bool {
				if (_345_v161).Contains(_155_v4) {
					return (_345_v161).Get(_155_v4).(bool)
				}
				return (_211_v51).F16()
			})()))
			var _rhs87 bool = (_211_v51).F16()
			_ = _rhs87
			var _rhs88 T0 = _344_v160
			_ = _rhs88
			var _rhs89 _dafny.Map = (_346_v162).Merge(_346_v162)
			_ = _rhs89
			_155_v4 = _rhs87
			_344_v160 = _rhs88
			_346_v162 = _rhs89
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
			_ = _index72
			(_161_v10).ArraySet1(_331_v147, (_index72).Int())
			var _347_v163 _dafny.Set
			_ = _347_v163
			_347_v163 = _dafny.SetOf(_299_v120)
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))
			_ = _index73
			(_300_v121).ArraySet1(!((_dafny.SetOf(_299_v120)).IsSubsetOf((_347_v163).Union(_347_v163))), (_index73).Int())
		}
	} else {
		var _348_v164 D12
		_ = _348_v164
		_348_v164 = Companion_D12_.Create_DC29_(_300_v121, _155_v4)
		var _349_v165 D1
		_ = _349_v165
		_349_v165 = Companion_D1_.Create_DC1_((_211_v51).F16())
		var _350_v166 _dafny.Map
		_ = _350_v166
		_350_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_211_v51).F16(), _349_v165)
		_348_v164 = (func() D12 {
			if ((_297_v118).Fm7(_350_v166, (_211_v51).F16(), _dafny.IntOfInt64(575), _162_globalState)) && (true) {
				return Companion_D12_.Create_DC29_(_300_v121, (_211_v51).F16())
			}
			return _348_v164
		})()
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))
		_ = _index74
		(_300_v121).ArraySet1(_155_v4, (_index74).Int())
		var _351_v167 _dafny.Set
		_ = _351_v167
		_351_v167 = _dafny.SetOf(_156_v5)
		_155_v4 = !(_351_v167).Equals((_351_v167).Intersection(_351_v167))
		if (_157_v6).Select((Companion_Default___.SafeIndex((_152_v1).Times((_dafny.MultiSetOf(!(Companion_Default___.Fm0((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _152_v1, _166_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v9, Companion_Default___.Fm0((_154_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-797), _dafny.IntOfUint32((_154_v3).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_160_v9).Cardinality()), _166_v14, _167_v15, _162_globalState)), _162_globalState)), (_211_v51).F16(), _155_v4)).Cardinality()), _dafny.IntOfUint32((_157_v6).Cardinality()))).Uint32()).(bool) {
			var _352_v168 *C1
			_ = _352_v168
			var _nw58 *C1 = New_C1_()
			_ = _nw58
			_nw58.Ctor__(!(!(!(_153_v2).Equals(_dafny.MultiSetOf(_152_v1)))), _210_v50)
			_352_v168 = _nw58
			var _353_v169 _dafny.Array
			_ = _353_v169
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
			_ = _nw59
			_353_v169 = _nw59
			var _354_v170 _dafny.Set
			_ = _354_v170
			_354_v170 = _dafny.SetOf(_160_v9)
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_353_v169), 0))
			_ = _index75
			(_353_v169).ArraySet1((_354_v170).Intersection(_dafny.SetOf(_160_v9)), (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_353_v169), 0))
			_ = _index76
			(_353_v169).ArraySet1(_354_v170, (_index76).Int())
			_152_v1 = Companion_Default___.Fm17(_155_v4, _152_v1, (_211_v51).F16(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), false)).Cardinality(), _162_globalState)
			_300_v121 = _300_v121
			var _355_v171 _dafny.Sequence
			_ = _355_v171
			_355_v171 = _dafny.SeqOf(_297_v118, _297_v118, _297_v118, _297_v118, _297_v118)
			var _356_v172 _dafny.Set
			_ = _356_v172
			_356_v172 = _dafny.SetOf(_355_v171)
			var _357_v173 _dafny.Map
			_ = _357_v173
			_357_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(_152_v1))), (_dafny.Zero).Minus((_356_v172).Cardinality()))
			_357_v173 = func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_154_v3).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _358_v174 _dafny.Int
					_358_v174 = interface{}(_compr_12).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_154_v3, _358_v174) {
						_coll12.Add((_358_v174).Minus((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)), _dafny.IntOfInt64(39))
					}
				}
				return _coll12.ToMap()
			}()
		} else {
			_155_v4 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("yx"), _160_v9)
			_155_v4 = (_300_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))).Int()).(bool)
			var _359_v175 D16
			_ = _359_v175
			_359_v175 = Companion_D16_.Create_DC41_(_166_v14)
			(_162_globalState).F2 = _dafny.IntOfUint32(((_359_v175).Dtor_cf64()).Cardinality())
			var _360_v176 _dafny.Map
			_ = _360_v176
			_360_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_211_v51).F16(), (_211_v51).F16())
			var _361_v177 _dafny.Map
			_ = _361_v177
			_361_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), ((_360_v176).Cardinality()).Minus(_dafny.IntOfUint32((_154_v3).Cardinality())))
			_361_v177 = (_361_v177).Update(((_361_v177).Merge(_361_v177)).Cardinality(), _dafny.IntOfInt64(881))
			var _362_v178 bool
			_ = _362_v178
			var _363_v179 _dafny.Int
			_ = _363_v179
			var _364_v180 bool
			_ = _364_v180
			var _365_v181 bool
			_ = _365_v181
			var _out32 bool
			_ = _out32
			var _out33 _dafny.Int
			_ = _out33
			var _out34 bool
			_ = _out34
			var _out35 bool
			_ = _out35
			_out32, _out33, _out34, _out35 = (_297_v118).M2(_162_globalState)
			_362_v178 = _out32
			_363_v179 = _out33
			_364_v180 = _out34
			_365_v181 = _out35
		}
		var _366_v182 *C2
		_ = _366_v182
		var _nw60 *C2 = New_C2_()
		_ = _nw60
		_nw60.Ctor__((_211_v51).F16(), (_211_v51).F16())
		_366_v182 = _nw60
		var _367_v183 _dafny.Map
		_ = _367_v183
		_367_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_366_v182, _155_v4)
		var _source8 D12 = Companion_D12_.Create_DC28_((_367_v183).Merge(_367_v183))
		_ = _source8
		if _source8.Is_DC29() {
			var _368___mcc_h8 _dafny.Array = _source8.Get_().(D12_DC29).Cf40
			_ = _368___mcc_h8
			var _369___mcc_h9 bool = _source8.Get_().(D12_DC29).Cf41
			_ = _369___mcc_h9
			var _370_cf41 bool = _369___mcc_h9
			_ = _370_cf41
			var _371_cf40 _dafny.Array = _368___mcc_h8
			_ = _371_cf40
			var _rhs90 _dafny.Int = _152_v1
			_ = _rhs90
			var _rhs91 _dafny.Int = (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)
			_ = _rhs91
			var _lhs56 *GlobalState = _162_globalState
			_ = _lhs56
			var _lhs57 *GlobalState = _162_globalState
			_ = _lhs57
			_lhs56.F2 = _rhs90
			_lhs57.F2 = _rhs91
			_297_v118 = _297_v118
			var _372_v184 _dafny.Int
			_ = _372_v184
			var _373_v185 _dafny.Int
			_ = _373_v185
			var _out36 _dafny.Int
			_ = _out36
			var _out37 _dafny.Int
			_ = _out37
			_out36, _out37 = (_211_v51).M5(_dafny.IntOfInt64(-863), _162_globalState)
			_372_v184 = _out36
			_373_v185 = _out37
			(_366_v182).M4((_366_v182).F15(), _372_v184, _162_globalState)
		} else if _source8.Is_DC28() {
			var _374___mcc_h10 _dafny.Map = _source8.Get_().(D12_DC28).Cf39
			_ = _374___mcc_h10
			var _375_cf39 _dafny.Map = _374___mcc_h10
			_ = _375_cf39
			var _376_v186 *C1
			_ = _376_v186
			var _nw61 *C1 = New_C1_()
			_ = _nw61
			_nw61.Ctor__((_dafny.IntOfInt64(-981)).Cmp(Companion_Default___.Fm17(!((_366_v182).F15()), _dafny.IntOfInt64(-403), (_211_v51).F16(), (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _162_globalState)) == 0, _211_v51.F17)
			_376_v186 = _nw61
			_366_v182 = _366_v182
			(_162_globalState).F2 = _152_v1
			var _377_v187 _dafny.Array
			_ = _377_v187
			_377_v187 = _161_v10
			var _378_v188 _dafny.Array
			_ = _378_v188
			var _nwElement0_8 _dafny.Array = _377_v187
			_ = _nwElement0_8
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(2))
			_ = _nw62
			(_nw62).ArraySet1(_nwElement0_8, 0)
			(_nw62).ArraySet1(_161_v10, 1)
			_378_v188 = _nw62
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_378_v188), 0))
			_ = _index77
			(_378_v188).ArraySet1(_377_v187, (_index77).Int())
			var _379_v190 T0
			_ = _379_v190
			var _nw63 *C2 = New_C2_()
			_ = _nw63
			_nw63.Ctor__(!(_158_v7).Equals(func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10((_376_v186).F16(), _162_globalState), _153_v2)).Keys().Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _380_v189 _dafny.Int
					_380_v189 = interface{}(_compr_13).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10((_376_v186).F16(), _162_globalState), _153_v2)).Contains(_380_v189) {
						_coll13.Add((_380_v189).Minus((_dafny.SetOf(false)).Cardinality()))
					}
				}
				return _coll13.ToSet()
			}()), true)
			_379_v190 = _nw63
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_378_v188), 0))
			_ = _index78
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
			_ = _index79
			var _rhs92 _dafny.Array = (func() _dafny.Array {
				if !((_366_v182).F15()) {
					return _377_v187
				}
				return _377_v187
			})()
			_ = _rhs92
			var _rhs93 _dafny.Int = _dafny.IntOfInt64(286)
			_ = _rhs93
			var _rhs94 T0 = _379_v190
			_ = _rhs94
			var _rhs95 _dafny.MultiSet = _165_v13
			_ = _rhs95
			var _rhs96 _dafny.Int = _152_v1
			_ = _rhs96
			var _lhs58 _dafny.Array = _378_v188
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_378_v188), 0))
			_ = _lhs59
			var _lhs60 _dafny.Array = _161_v10
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))
			_ = _lhs61
			var _lhs62 *GlobalState = _162_globalState
			_ = _lhs62
			(_lhs58).ArraySet1(_rhs92, (_lhs59).Int())
			(_lhs60).ArraySet1(_rhs93, (_lhs61).Int())
			_379_v190 = _rhs94
			_165_v13 = _rhs95
			_lhs62.F2 = _rhs96
		} else {
			var _381___mcc_h11 D12 = _source8.Get_().(D12_DC30).Cf42
			_ = _381___mcc_h11
			var _382_cf42 D12 = _381___mcc_h11
			_ = _382_cf42
			var _383_v191 T0
			_ = _383_v191
			var _nw64 *C1 = New_C1_()
			_ = _nw64
			_nw64.Ctor__((func() bool {
				if (_156_v5).Contains(_dafny.IntOfInt64(128)) {
					return (_156_v5).Get(_dafny.IntOfInt64(128)).(bool)
				}
				return true
			})(), _211_v51.F17)
			_383_v191 = _nw64
			var _rhs97 T0 = _383_v191
			_ = _rhs97
			var _rhs98 _dafny.Int = Companion_Default___.SafeDivisionInt(_152_v1, (_dafny.Zero).Minus((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)))
			_ = _rhs98
			var _rhs99 _dafny.Int = ((_152_v1).Minus((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int))).Times(_152_v1)
			_ = _rhs99
			var _rhs100 _dafny.Set = _158_v7
			_ = _rhs100
			var _lhs63 *GlobalState = _162_globalState
			_ = _lhs63
			var _lhs64 *GlobalState = _162_globalState
			_ = _lhs64
			_383_v191 = _rhs97
			_lhs63.F2 = _rhs98
			_lhs64.F2 = _rhs99
			_158_v7 = _rhs100
			var _384_v192 *C0
			_ = _384_v192
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__(_366_v182.F14, _152_v1)
			_384_v192 = _nw65
			var _385_v193 _dafny.Map
			_ = _385_v193
			_385_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_366_v182.F14, _384_v192)
			var _386_v194 D13
			_ = _386_v194
			_386_v194 = Companion_D13_.Create_DC33_(_300_v121, _366_v182, _385_v193, (_211_v51).F16())
			var _387_v195 _dafny.Map
			_ = _387_v195
			_387_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_153_v2).Contains((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)) {
					return (_153_v2).Multiplicity((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int))
				}
				return Companion_Default___.Fm19(_152_v1, _162_globalState)
			})())), (_386_v194).Dtor_cf45())
			_387_v195 = (_387_v195).Update(_152_v1, _300_v121)
			var _388_v196 bool
			_ = _388_v196
			var _389_v197 _dafny.Int
			_ = _389_v197
			var _390_v198 bool
			_ = _390_v198
			var _391_v199 bool
			_ = _391_v199
			var _out38 bool
			_ = _out38
			var _out39 _dafny.Int
			_ = _out39
			var _out40 bool
			_ = _out40
			var _out41 bool
			_ = _out41
			_out38, _out39, _out40, _out41 = (_297_v118).M2(_162_globalState)
			_388_v196 = _out38
			_389_v197 = _out39
			_390_v198 = _out40
			_391_v199 = _out41
			(_162_globalState).F2 = (_dafny.Zero).Minus((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int))
		}
	}
	var _392_i7 _dafny.Int
	_ = _392_i7
	_392_i7 = _dafny.Zero
	{
		for (_300_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))).Int()).(bool) {
			{
				if (_392_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_392_i7 = (_392_i7).Plus(_dafny.One)
				var _393_v200 _dafny.Array
				_ = _393_v200
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(13))
				_ = _nw66
				_393_v200 = _nw66
				var _394_v201 D3
				_ = _394_v201
				_394_v201 = Companion_D3_.Create_DC6_((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), (_211_v51).F16(), (_211_v51).F16(), _152_v1, (_211_v51).F16())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_393_v200), 0))
				_ = _index80
				(_393_v200).ArraySet1(_394_v201, (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_393_v200), 0))
				_ = _index81
				(_393_v200).ArraySet1(_394_v201, (_index81).Int())
				_155_v4 = (_300_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))).Int()).(bool)
				var _395_v202 _dafny.Set
				_ = _395_v202
				_395_v202 = _dafny.SetOf((_211_v51).F16())
				var _396_v203 _dafny.Sequence
				_ = _396_v203
				_396_v203 = _dafny.SeqOf((_395_v202).Union(_395_v202), _395_v202)
				_396_v203 = _dafny.Companion_Sequence_.Concatenate(_396_v203, _dafny.SeqOf(_dafny.SetOf((_211_v51).F16()), _395_v202, _395_v202, _395_v202))
				_155_v4 = _dafny.Companion_Sequence_.IsProperPrefixOf(_160_v9, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yecri"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_397_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))))
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_155_v4 = (_300_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_300_v121), 0))).Int()).(bool)
	_155_v4 = !((_157_v6).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_152_v1), _dafny.IntOfUint32((_157_v6).Cardinality()))).Uint32()).(bool))
	_dafny.Print(_152_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v2).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(581), _dafny.IntOfInt64(581))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_154_v3, _dafny.SeqOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_155_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(581), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_157_v6, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v7).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v8))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_160_v9.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState.F4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState.F4).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState.F4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState.F4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F7()).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v12).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v13).Equals(_dafny.MultiSetOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_166_v14, _dafny.SeqOf(_dafny.MultiSetOf(false, false, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("joueqk"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_211_v51).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_298_v119).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_299_v120)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_300_v121).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_300_v121).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_300_v121).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_300_v121).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_300_v121).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v122).Dtor_cf29())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v122).Dtor_cf30())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_302_v123).Equals(_dafny.MultiSetOf(Companion_D9_.Create_DC20_(false, _dafny.CodePoint('j')), Companion_D9_.Create_DC20_(false, _dafny.CodePoint('j')), Companion_D9_.Create_DC20_(false, _dafny.CodePoint('j')), Companion_D9_.Create_DC20_(false, _dafny.CodePoint('j')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_303_i6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_321_v141).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_392_i7)
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

func (CompanionStruct_D0_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D0) Dtor_cf0() _dafny.Int {
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

type D1_DC2 struct {
	Cf2 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 bool) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf3 bool
	Cf4 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 bool, Cf4 _dafny.Int) D1 {
	return D1{D1_DC3{Cf3, Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
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
	return Companion_D1_.Create_DC2_(false)
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf4
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
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf2 == data2.Cf2
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf5 _dafny.Array
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.Array) D2 {
	return D2{D2_DC4{Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf5 == data2.Cf5
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
	Cf7  _dafny.Int
	Cf8  bool
	Cf9  bool
	Cf10 _dafny.Int
	Cf11 bool
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf7 _dafny.Int, Cf8 bool, Cf9 bool, Cf10 _dafny.Int, Cf11 bool) D3 {
	return D3{D3_DC6{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC7 struct {
	Cf12 _dafny.Int
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf12 _dafny.Int) D3 {
	return D3{D3_DC7{Cf12}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC8 struct {
	Cf13 bool
	Cf14 _dafny.Int
	Cf15 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf13 bool, Cf14 _dafny.Int, Cf15 _dafny.Int) D3 {
	return D3{D3_DC8{Cf13, Cf14, Cf15}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC5 struct {
	Cf6 _dafny.Array
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf6 _dafny.Array) D3 {
	return D3{D3_DC5{Cf6}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(_dafny.Zero, false, false, _dafny.Zero, false)
}

func (_this D3) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf7
}

func (_this D3) Dtor_cf8() bool {
	return _this.Get_().(D3_DC6).Cf8
}

func (_this D3) Dtor_cf9() bool {
	return _this.Get_().(D3_DC6).Cf9
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf10
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC6).Cf11
}

func (_this D3) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D3_DC7).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) Dtor_cf6() _dafny.Array {
	return _this.Get_().(D3_DC5).Cf6
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf6 == data2.Cf6
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
	Cf17 T0
	Cf18 _dafny.Int
	Cf19 _dafny.Int
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf17 T0, Cf18 _dafny.Int, Cf19 _dafny.Int) D4 {
	return D4{D4_DC10{Cf17, Cf18, Cf19}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf16 _dafny.MultiSet
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf16 _dafny.MultiSet) D4 {
	return D4{D4_DC9{Cf16}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_((T0)(nil), _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf17() T0 {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf19
}

func (_this D4) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D4_DC9).Cf16
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf17, data2.Cf17) && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Cmp(data2.Cf19) == 0
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D5_DC11 struct {
	Cf20 _dafny.Sequence
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf20 _dafny.Sequence) D5 {
	return D5{D5_DC11{Cf20}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D5) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D5_DC11).Cf20
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + data.Cf20.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
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

type D6_DC13 struct {
	Cf22 bool
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf22 bool) D6 {
	return D6{D6_DC13{Cf22}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC12 struct {
	Cf21 _dafny.Map
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf21 _dafny.Map) D6 {
	return D6{D6_DC12{Cf21}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC13_(false)
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) Dtor_cf21() _dafny.Map {
	return _this.Get_().(D6_DC12).Cf21
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf22 == data2.Cf22
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
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

type D7_DC15 struct {
	Cf24 _dafny.Int
	Cf25 bool
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf24 _dafny.Int, Cf25 bool) D7 {
	return D7{D7_DC15{Cf24, Cf25}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC14 struct {
	Cf23 _dafny.Set
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf23 _dafny.Set) D7 {
	return D7{D7_DC14{Cf23}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

type D7_DC16 struct {
	Cf26 D7
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf26 D7) D7 {
	return D7{D7_DC16{Cf26}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC15_(_dafny.Zero, false)
}

func (_this D7) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf24
}

func (_this D7) Dtor_cf25() bool {
	return _this.Get_().(D7_DC15).Cf25
}

func (_this D7) Dtor_cf23() _dafny.Set {
	return _this.Get_().(D7_DC14).Cf23
}

func (_this D7) Dtor_cf26() D7 {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf23) + ")"
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
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25 == data2.Cf25
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D8_DC18 struct {
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_() D8 {
	return D8{D8_DC18{}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

type D8_DC17 struct {
	Cf27 _dafny.Map
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf27 _dafny.Map) D8 {
	return D8{D8_DC17{Cf27}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC18_()
}

func (_this D8) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D8_DC17).Cf27
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18"
		}
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf27) + ")"
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
			_, ok := other.Get_().(D8_DC18)
			return ok
		}
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	Cf29 bool
	Cf30 _dafny.CodePoint
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf29 bool, Cf30 _dafny.CodePoint) D9 {
	return D9{D9_DC20{Cf29, Cf30}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC19 struct {
	Cf28 _dafny.Map
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf28 _dafny.Map) D9 {
	return D9{D9_DC19{Cf28}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC21 struct {
	Cf31 D9
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf31 D9) D9 {
	return D9{D9_DC21{Cf31}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC20_(false, _dafny.CodePoint('D'))
}

func (_this D9) Dtor_cf29() bool {
	return _this.Get_().(D9_DC20).Cf29
}

func (_this D9) Dtor_cf30() _dafny.CodePoint {
	return _this.Get_().(D9_DC20).Cf30
}

func (_this D9) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D9_DC19).Cf28
}

func (_this D9) Dtor_cf31() D9 {
	return _this.Get_().(D9_DC21).Cf31
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf33 _dafny.Int
	Cf34 bool
	Cf35 bool
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf33 _dafny.Int, Cf34 bool, Cf35 bool) D10 {
	return D10{D10_DC23{Cf33, Cf34, Cf35}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

type D10_DC22 struct {
	Cf32 *C1
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf32 *C1) D10 {
	return D10{D10_DC22{Cf32}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC23_(_dafny.Zero, false, false)
}

func (_this D10) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D10_DC23).Cf33
}

func (_this D10) Dtor_cf34() bool {
	return _this.Get_().(D10_DC23).Cf34
}

func (_this D10) Dtor_cf35() bool {
	return _this.Get_().(D10_DC23).Cf35
}

func (_this D10) Dtor_cf32() *C1 {
	return _this.Get_().(D10_DC22).Cf32
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf32) + ")"
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
			return ok && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35
		}
	case D10_DC22:
		{
			data2, ok := other.Get_().(D10_DC22)
			return ok && data1.Cf32 == data2.Cf32
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

type D11_DC25 struct {
	Cf37 _dafny.CodePoint
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf37 _dafny.CodePoint) D11 {
	return D11{D11_DC25{Cf37}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC26 struct {
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_() D11 {
	return D11{D11_DC26{}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC24 struct {
	Cf36 _dafny.Array
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf36 _dafny.Array) D11 {
	return D11{D11_DC24{Cf36}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC27 struct {
	Cf38 D11
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf38 D11) D11 {
	return D11{D11_DC27{Cf38}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC25_(_dafny.CodePoint('D'))
}

func (_this D11) Dtor_cf37() _dafny.CodePoint {
	return _this.Get_().(D11_DC25).Cf37
}

func (_this D11) Dtor_cf36() _dafny.Array {
	return _this.Get_().(D11_DC24).Cf36
}

func (_this D11) Dtor_cf38() D11 {
	return _this.Get_().(D11_DC27).Cf38
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D11_DC26:
		{
			return "D11.DC26"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf37 == data2.Cf37
		}
	case D11_DC26:
		{
			_, ok := other.Get_().(D11_DC26)
			return ok
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf36 == data2.Cf36
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf38.Equals(data2.Cf38)
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
	Cf40 _dafny.Array
	Cf41 bool
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf40 _dafny.Array, Cf41 bool) D12 {
	return D12{D12_DC29{Cf40, Cf41}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC28 struct {
	Cf39 _dafny.Map
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf39 _dafny.Map) D12 {
	return D12{D12_DC28{Cf39}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

type D12_DC30 struct {
	Cf42 D12
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf42 D12) D12 {
	return D12{D12_DC30{Cf42}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC29_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D12) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D12_DC29).Cf40
}

func (_this D12) Dtor_cf41() bool {
	return _this.Get_().(D12_DC29).Cf41
}

func (_this D12) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D12_DC28).Cf39
}

func (_this D12) Dtor_cf42() D12 {
	return _this.Get_().(D12_DC30).Cf42
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41 == data2.Cf41
		}
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
			return ok && data1.Cf39.Equals(data2.Cf39)
		}
	case D12_DC30:
		{
			data2, ok := other.Get_().(D12_DC30)
			return ok && data1.Cf42.Equals(data2.Cf42)
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
	Cf44 _dafny.Int
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf44 _dafny.Int) D13 {
	return D13{D13_DC32{Cf44}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

type D13_DC33 struct {
	Cf45 _dafny.Array
	Cf46 *C2
	Cf47 _dafny.Map
	Cf48 bool
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf45 _dafny.Array, Cf46 *C2, Cf47 _dafny.Map, Cf48 bool) D13 {
	return D13{D13_DC33{Cf45, Cf46, Cf47, Cf48}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

type D13_DC31 struct {
	Cf43 _dafny.Set
}

func (D13_DC31) isD13() {}

func (CompanionStruct_D13_) Create_DC31_(Cf43 _dafny.Set) D13 {
	return D13{D13_DC31{Cf43}}
}

func (_this D13) Is_DC31() bool {
	_, ok := _this.Get_().(D13_DC31)
	return ok
}

type D13_DC34 struct {
	Cf49 D13
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf49 D13) D13 {
	return D13{D13_DC34{Cf49}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC32_(_dafny.Zero)
}

func (_this D13) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D13_DC32).Cf44
}

func (_this D13) Dtor_cf45() _dafny.Array {
	return _this.Get_().(D13_DC33).Cf45
}

func (_this D13) Dtor_cf46() *C2 {
	return _this.Get_().(D13_DC33).Cf46
}

func (_this D13) Dtor_cf47() _dafny.Map {
	return _this.Get_().(D13_DC33).Cf47
}

func (_this D13) Dtor_cf48() bool {
	return _this.Get_().(D13_DC33).Cf48
}

func (_this D13) Dtor_cf43() _dafny.Set {
	return _this.Get_().(D13_DC31).Cf43
}

func (_this D13) Dtor_cf49() D13 {
	return _this.Get_().(D13_DC34).Cf49
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC32:
		{
			return "D13.DC32" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D13_DC31:
		{
			return "D13.DC31" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf44.Cmp(data2.Cf44) == 0
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46 == data2.Cf46 && data1.Cf47.Equals(data2.Cf47) && data1.Cf48 == data2.Cf48
		}
	case D13_DC31:
		{
			data2, ok := other.Get_().(D13_DC31)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D14_DC36 struct {
	Cf51 _dafny.Int
	Cf52 bool
	Cf53 bool
	Cf54 _dafny.Int
	Cf55 bool
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf51 _dafny.Int, Cf52 bool, Cf53 bool, Cf54 _dafny.Int, Cf55 bool) D14 {
	return D14{D14_DC36{Cf51, Cf52, Cf53, Cf54, Cf55}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC37 struct {
	Cf56 bool
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf56 bool) D14 {
	return D14{D14_DC37{Cf56}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

type D14_DC35 struct {
	Cf50 _dafny.Sequence
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf50 _dafny.Sequence) D14 {
	return D14{D14_DC35{Cf50}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC36_(_dafny.Zero, false, false, _dafny.Zero, false)
}

func (_this D14) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf51
}

func (_this D14) Dtor_cf52() bool {
	return _this.Get_().(D14_DC36).Cf52
}

func (_this D14) Dtor_cf53() bool {
	return _this.Get_().(D14_DC36).Cf53
}

func (_this D14) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf54
}

func (_this D14) Dtor_cf55() bool {
	return _this.Get_().(D14_DC36).Cf55
}

func (_this D14) Dtor_cf56() bool {
	return _this.Get_().(D14_DC37).Cf56
}

func (_this D14) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D14_DC35).Cf50
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0 && data1.Cf55 == data2.Cf55
		}
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf56 == data2.Cf56
		}
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D15_DC39 struct {
	Cf58 bool
	Cf59 _dafny.Int
	Cf60 bool
	Cf61 _dafny.Int
	Cf62 _dafny.Int
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf58 bool, Cf59 _dafny.Int, Cf60 bool, Cf61 _dafny.Int, Cf62 _dafny.Int) D15 {
	return D15{D15_DC39{Cf58, Cf59, Cf60, Cf61, Cf62}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

type D15_DC38 struct {
	Cf57 T1
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf57 T1) D15 {
	return D15{D15_DC38{Cf57}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC40 struct {
	Cf63 D15
}

func (D15_DC40) isD15() {}

func (CompanionStruct_D15_) Create_DC40_(Cf63 D15) D15 {
	return D15{D15_DC40{Cf63}}
}

func (_this D15) Is_DC40() bool {
	_, ok := _this.Get_().(D15_DC40)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC39_(false, _dafny.Zero, false, _dafny.Zero, _dafny.Zero)
}

func (_this D15) Dtor_cf58() bool {
	return _this.Get_().(D15_DC39).Cf58
}

func (_this D15) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf59
}

func (_this D15) Dtor_cf60() bool {
	return _this.Get_().(D15_DC39).Cf60
}

func (_this D15) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf61
}

func (_this D15) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf62
}

func (_this D15) Dtor_cf57() T1 {
	return _this.Get_().(D15_DC38).Cf57
}

func (_this D15) Dtor_cf63() D15 {
	return _this.Get_().(D15_DC40).Cf63
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ")"
		}
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D15_DC40:
		{
			return "D15.DC40" + "(" + _dafny.String(data.Cf63) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC39:
		{
			data2, ok := other.Get_().(D15_DC39)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60 == data2.Cf60 && data1.Cf61.Cmp(data2.Cf61) == 0 && data1.Cf62.Cmp(data2.Cf62) == 0
		}
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && _dafny.AreEqual(data1.Cf57, data2.Cf57)
		}
	case D15_DC40:
		{
			data2, ok := other.Get_().(D15_DC40)
			return ok && data1.Cf63.Equals(data2.Cf63)
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

type D16_DC42 struct {
	Cf65 _dafny.Int
	Cf66 bool
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf65 _dafny.Int, Cf66 bool) D16 {
	return D16{D16_DC42{Cf65, Cf66}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC43 struct {
	Cf67 _dafny.Array
	Cf68 _dafny.Int
	Cf69 _dafny.MultiSet
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf67 _dafny.Array, Cf68 _dafny.Int, Cf69 _dafny.MultiSet) D16 {
	return D16{D16_DC43{Cf67, Cf68, Cf69}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

type D16_DC41 struct {
	Cf64 _dafny.Sequence
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf64 _dafny.Sequence) D16 {
	return D16{D16_DC41{Cf64}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC42_(_dafny.Zero, false)
}

func (_this D16) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf65
}

func (_this D16) Dtor_cf66() bool {
	return _this.Get_().(D16_DC42).Cf66
}

func (_this D16) Dtor_cf67() _dafny.Array {
	return _this.Get_().(D16_DC43).Cf67
}

func (_this D16) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D16_DC43).Cf68
}

func (_this D16) Dtor_cf69() _dafny.MultiSet {
	return _this.Get_().(D16_DC43).Cf69
}

func (_this D16) Dtor_cf64() _dafny.Sequence {
	return _this.Get_().(D16_DC41).Cf64
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ")"
		}
	case D16_DC43:
		{
			return "D16.DC43" + "(" + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf65.Cmp(data2.Cf65) == 0 && data1.Cf66 == data2.Cf66
		}
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf67 == data2.Cf67 && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69.Equals(data2.Cf69)
		}
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf64.Equals(data2.Cf64)
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

type D17_DC45 struct {
	Cf71 _dafny.Map
	Cf72 bool
}

func (D17_DC45) isD17() {}

func (CompanionStruct_D17_) Create_DC45_(Cf71 _dafny.Map, Cf72 bool) D17 {
	return D17{D17_DC45{Cf71, Cf72}}
}

func (_this D17) Is_DC45() bool {
	_, ok := _this.Get_().(D17_DC45)
	return ok
}

type D17_DC44 struct {
	Cf70 _dafny.MultiSet
}

func (D17_DC44) isD17() {}

func (CompanionStruct_D17_) Create_DC44_(Cf70 _dafny.MultiSet) D17 {
	return D17{D17_DC44{Cf70}}
}

func (_this D17) Is_DC44() bool {
	_, ok := _this.Get_().(D17_DC44)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC45_(_dafny.EmptyMap, false)
}

func (_this D17) Dtor_cf71() _dafny.Map {
	return _this.Get_().(D17_DC45).Cf71
}

func (_this D17) Dtor_cf72() bool {
	return _this.Get_().(D17_DC45).Cf72
}

func (_this D17) Dtor_cf70() _dafny.MultiSet {
	return _this.Get_().(D17_DC44).Cf70
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC45:
		{
			return "D17.DC45" + "(" + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D17_DC44:
		{
			return "D17.DC44" + "(" + _dafny.String(data.Cf70) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC45:
		{
			data2, ok := other.Get_().(D17_DC45)
			return ok && data1.Cf71.Equals(data2.Cf71) && data1.Cf72 == data2.Cf72
		}
	case D17_DC44:
		{
			data2, ok := other.Get_().(D17_DC44)
			return ok && data1.Cf70.Equals(data2.Cf70)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm3(globalState *GlobalState) _dafny.Map
	Fm4(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) bool
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
	Fm3(globalState *GlobalState) _dafny.Map
	Fm4(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) bool
	F10() _dafny.Int
	F10_set_(value _dafny.Int)
	Fm5(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) bool
	Fm6(p0 _dafny.Set, globalState *GlobalState) D1
	M1(globalState *GlobalState) _dafny.Int
	M2(globalState *GlobalState) (bool, _dafny.Int, bool, bool)
	F11() _dafny.Map
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
	F2  _dafny.Int
	F4  _dafny.Array
	_f0 bool
	_f1 bool
	_f3 _dafny.Array
	_f5 _dafny.Int
	_f6 _dafny.Int
	_f7 _dafny.Array
	_f8 _dafny.Int
	_f9 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this.F4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = false
	_this._f1 = false
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 _dafny.Int, f3 _dafny.Array, f4 _dafny.Array, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.Array, f8 _dafny.Int, f9 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
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
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
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
func (_this *GlobalState) F7() _dafny.Array {
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 bool
	_f13 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = false
	_this._f13 = _dafny.Zero
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

func (_this *C0) Ctor__(f12 bool, f13 _dafny.Int) {
	{
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C0) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aqbqv"), _dafny.UnicodeSeqOfUtf8Bytes("x")), _dafny.UnicodeSeqOfUtf8Bytes("g"))
	}
}
func (_this *C0) F12() bool {
	{
		return _this._f12
	}
}
func (_this *C0) F13() _dafny.Int {
	{
		return _this._f13
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F17  D3
	_f16 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F17 = Companion_D3_.Default()
	_this._f16 = false
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

func (_this *C1) Ctor__(f16 bool, f17 D3) {
	{
		(_this)._f16 = f16
		(_this).F17 = f17
	}
}
func (_this *C1) Fm3(globalState *GlobalState) _dafny.Map {
	{
		var _source9 _dafny.Int = _dafny.IntOfInt64(310)
		_ = _source9
		var _398___mcc_h0 _dafny.Int = _source9
		_ = _398___mcc_h0
		var _399_cf0 _dafny.Int = _398___mcc_h0
		_ = _399_cf0
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-795)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _399_cf0))
	}
}
func (_this *C1) Fm4(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (_this).F16()
	}
}
func (_this *C1) M5(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _400_v0 D1
		_ = _400_v0
		_400_v0 = Companion_D1_.Create_DC3_((_this).F16(), Companion_Default___.Fm1((_this).F16(), true, globalState))
		var _source10 D4 = func(_source11 D1) D4 {
			if _source11.Is_DC2() {
				var _401___mcc_h4 bool = _source11.Get_().(D1_DC2).Cf2
				_ = _401___mcc_h4
				var _402_cf2 bool = _401___mcc_h4
				_ = _402_cf2
				return Companion_D4_.Create_DC9_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_402_cf2)))
			} else if _source11.Is_DC3() {
				var _403___mcc_h5 bool = _source11.Get_().(D1_DC3).Cf3
				_ = _403___mcc_h5
				var _404___mcc_h6 _dafny.Int = _source11.Get_().(D1_DC3).Cf4
				_ = _404___mcc_h6
				var _405_cf4 _dafny.Int = _404___mcc_h6
				_ = _405_cf4
				var _406_cf3 bool = _403___mcc_h5
				_ = _406_cf3
				return Companion_D4_.Create_DC9_(_dafny.MultiSetOf(_406_cf3, (_this).F16()))
			} else {
				var _407___mcc_h7 bool = _source11.Get_().(D1_DC1).Cf1
				_ = _407___mcc_h7
				var _408_cf1 bool = _407___mcc_h7
				_ = _408_cf1
				return Companion_D4_.Create_DC9_(_dafny.MultiSetOf(true))
			}
		}(_400_v0)
		_ = _source10
		if _source10.Is_DC10() {
			var _409___mcc_h0 T0 = _source10.Get_().(D4_DC10).Cf17
			_ = _409___mcc_h0
			var _410___mcc_h1 _dafny.Int = _source10.Get_().(D4_DC10).Cf18
			_ = _410___mcc_h1
			var _411___mcc_h2 _dafny.Int = _source10.Get_().(D4_DC10).Cf19
			_ = _411___mcc_h2
			var _412_cf19 _dafny.Int = _411___mcc_h2
			_ = _412_cf19
			var _413_cf18 _dafny.Int = _410___mcc_h1
			_ = _413_cf18
			var _414_cf17 T0 = _409___mcc_h0
			_ = _414_cf17
			var _415_v1 _dafny.Map
			_ = _415_v1
			_415_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (func() _dafny.Int {
				if (_this).F16() {
					return _412_cf19
				}
				return _413_cf18
			})())
			_415_v1 = (_415_v1).Update((_this).F16(), _412_cf19)
			var _416_v2 D3
			_ = _416_v2
			_416_v2 = Companion_D3_.Create_DC7_((p0).Minus(_413_cf18))
			var _417_v3 _dafny.Map
			_ = _417_v3
			_417_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_this).F16())
			_416_v2 = (func() D3 {
				if ((_417_v3).Cardinality()).Cmp(p0) >= 0 {
					return _416_v2
				}
				return (func() D3 {
					if (_this).F16() {
						return _416_v2
					}
					return _416_v2
				})()
			})()
			var _418_v4 _dafny.Sequence
			_ = _418_v4
			_418_v4 = _dafny.SeqOf(_dafny.IntOfInt64(980))
			var _419_v5 _dafny.Sequence
			_ = _419_v5
			_419_v5 = _dafny.SeqOf(_dafny.IntOfInt64(-191), (_418_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_418_v4).Cardinality()))).Uint32()).(_dafny.Int))
			_419_v5 = _dafny.Companion_Sequence_.Update(_418_v4, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_418_v4).Cardinality()))).Uint32(), p0)
			var _420_v6 _dafny.Sequence
			_ = _420_v6
			_420_v6 = _dafny.UnicodeSeqOfUtf8Bytes("raf")
			(globalState).F2 = _dafny.IntOfUint32((Companion_Default___.Fm18((func() _dafny.Int {
				if (_this).F16() {
					return _dafny.IntOfUint32((_420_v6).Cardinality())
				}
				return _413_cf18
			})(), (_this).F16(), _412_cf19, globalState)).Cardinality())
		} else {
			var _421___mcc_h3 _dafny.MultiSet = _source10.Get_().(D4_DC9).Cf16
			_ = _421___mcc_h3
			var _422_cf16 _dafny.MultiSet = _421___mcc_h3
			_ = _422_cf16
			var _423_v7 bool
			_ = _423_v7
			_423_v7 = true
			_423_v7 = (_this).F16()
			var _424_v8 _dafny.Sequence
			_ = _424_v8
			_424_v8 = _dafny.UnicodeSeqOfUtf8Bytes("fg")
			var _425_v9 _dafny.CodePoint
			_ = _425_v9
			_425_v9 = _dafny.CodePoint('t')
			var _426_v10 _dafny.MultiSet
			_ = _426_v10
			_426_v10 = _dafny.MultiSetOf(_425_v9, _dafny.CodePoint('h'), _425_v9)
			var _427_v11 _dafny.MultiSet
			_ = _427_v11
			_427_v11 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0))
			var _428_v12 _dafny.Set
			_ = _428_v12
			_428_v12 = _dafny.SetOf((_427_v11).Cardinality())
			var _429_v14 _dafny.Sequence
			_ = _429_v14
			_429_v14 = _dafny.SeqOf(_424_v8)
			var _430_v15 _dafny.Sequence
			_ = _430_v15
			_430_v15 = _dafny.SeqOf(Companion_Default___.Fm0(p0, p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer23 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_431_cf16 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_432_i0 _dafny.Int) _dafny.MultiSet {
					return _431_cf16
				}
			})(_422_cf16))), func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_429_v14, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_429_v14).Cardinality()))).Uint32(), _424_v8)).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _433_v13 _dafny.Sequence
					_433_v13 = interface{}(_compr_14).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_429_v14, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_429_v14).Cardinality()))).Uint32(), _424_v8), _433_v13) {
						_coll14.Add(_433_v13, (_this).F16())
					}
				}
				return _coll14.ToMap()
			}(), globalState), _423_v7)
			var _434_v16 _dafny.Array
			_ = _434_v16
			var _nwElement0_9 bool = (_this).F16()
			_ = _nwElement0_9
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(22))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_9, 0)
			(_nw67).ArraySet1((func() bool {
				if !((_this).F16()) {
					return (_this).F16()
				}
				return (_this).F16()
			})(), 1)
			(_nw67).ArraySet1((_this).F16(), 2)
			(_nw67).ArraySet1((_this).F16(), 3)
			(_nw67).ArraySet1(_dafny.Companion_Sequence_.Equal(_424_v8, _dafny.UnicodeSeqOfUtf8Bytes("qmuiq")), 4)
			(_nw67).ArraySet1((_426_v10).IsDisjointFrom(_426_v10), 5)
			(_nw67).ArraySet1((_this).F16(), 6)
			(_nw67).ArraySet1(!(_423_v7), 7)
			(_nw67).ArraySet1((_428_v12).IsSubsetOf(_428_v12), 8)
			(_nw67).ArraySet1((_423_v7) == (_423_v7), 9)
			(_nw67).ArraySet1(_423_v7, 10)
			(_nw67).ArraySet1(_423_v7, 11)
			(_nw67).ArraySet1(_423_v7, 12)
			(_nw67).ArraySet1(_423_v7, 13)
			(_nw67).ArraySet1(!(!(((_this).F16()) && (_423_v7))), 14)
			(_nw67).ArraySet1(_423_v7, 15)
			(_nw67).ArraySet1(_423_v7, 16)
			(_nw67).ArraySet1((_this).F16(), 17)
			(_nw67).ArraySet1(_423_v7, 18)
			(_nw67).ArraySet1((_this).F16(), 19)
			(_nw67).ArraySet1(_dafny.Companion_Sequence_.Equal(_430_v15, _dafny.Companion_Sequence_.Update(_430_v15, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_430_v15).Cardinality()))).Uint32(), (_this).F16())), 20)
			(_nw67).ArraySet1(_423_v7, 21)
			_434_v16 = _nw67
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_434_v16), 0))
			_ = _index82
			(_434_v16).ArraySet1((_this).F16(), (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_434_v16), 0))
			_ = _index83
			(_434_v16).ArraySet1(true, (_index83).Int())
			var _435_v17 _dafny.Sequence
			_ = _435_v17
			_435_v17 = _dafny.SeqOf(_422_cf16)
			var _436_v18 _dafny.Map
			_ = _436_v18
			_436_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_424_v8, _423_v7)
			var _437_v19 _dafny.Map
			_ = _437_v19
			_437_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_v9, _436_v18)
			var _438_v20 _dafny.Map
			_ = _438_v20
			_438_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_434_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_434_v16), 0))).Int()).(bool), Companion_Default___.Fm0(p0, _dafny.IntOfUint32((Companion_Default___.Fm18(p0, (_this).F16(), p0, globalState)).Cardinality()), _435_v17, (func() _dafny.Map {
				if (_437_v19).Contains(_dafny.CodePoint('b')) {
					return (_437_v19).Get(_dafny.CodePoint('b')).(_dafny.Map)
				}
				return _436_v18
			})(), globalState))
			var _439_v21 _dafny.Set
			_ = _439_v21
			_439_v21 = _dafny.SetOf((_434_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_434_v16), 0))).Int()).(bool))
			var _440_v22 _dafny.Sequence
			_ = _440_v22
			_440_v22 = _dafny.SeqOf(_dafny.SetOf((_434_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_434_v16), 0))).Int()).(bool), false), _439_v21, _dafny.SetOf((_434_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_434_v16), 0))).Int()).(bool), !((_this).F16())))
			_438_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(((_440_v22).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_440_v22).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), p0, _435_v17, _436_v18, globalState), (_this).F16())
			var _441_v23 _dafny.Array
			_ = _441_v23
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw68
			_441_v23 = _nw68
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_441_v23), 0))
			_ = _index84
			(_441_v23).ArraySet1(_430_v15, (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_441_v23), 0))
			_ = _index85
			(_441_v23).ArraySet1(_430_v15, (_index85).Int())
		}
		var _442_i1 _dafny.Int
		_ = _442_i1
		_442_i1 = _dafny.Zero
		{
			for (_dafny.IntOfInt64(-271)).Cmp(p0) >= 0 {
				{
					if (_442_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_442_i1 = (_442_i1).Plus(_dafny.One)
					var _443_v24 bool
					_ = _443_v24
					_443_v24 = false
					_443_v24 = (_this).F16()
					r1 = p0
					var _444_v25 _dafny.MultiSet
					_ = _444_v25
					_444_v25 = _dafny.MultiSetOf((_dafny.Zero).Minus(Companion_Default___.Fm19(p0, globalState)))
					var _445_v26 _dafny.Map
					_ = _445_v26
					_445_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _446_v27 _dafny.Array
					_ = _446_v27
					var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
					_ = _nw69
					_446_v27 = _nw69
					var _447_v28 _dafny.Map
					_ = _447_v28
					_447_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_443_v24), _446_v27)
					var _448_v29 _dafny.Array
					_ = _448_v29
					var _nwElement0_10 bool = (_this).F16()
					_ = _nwElement0_10
					var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(11))
					_ = _nw70
					(_nw70).ArraySet1(_nwElement0_10, 0)
					(_nw70).ArraySet1(_443_v24, 1)
					(_nw70).ArraySet1(true, 2)
					(_nw70).ArraySet1((_this).F16(), 3)
					(_nw70).ArraySet1(_443_v24, 4)
					(_nw70).ArraySet1(_443_v24, 5)
					(_nw70).ArraySet1((_this).F16(), 6)
					(_nw70).ArraySet1((_this).F16(), 7)
					(_nw70).ArraySet1((_this).F16(), 8)
					(_nw70).ArraySet1(true, 9)
					(_nw70).ArraySet1((_this).F16(), 10)
					_448_v29 = _nw70
					var _449_v30 _dafny.Set
					_ = _449_v30
					_449_v30 = _dafny.SetOf(_448_v29, _448_v29)
					var _450_v31 D7
					_ = _450_v31
					_450_v31 = Companion_D7_.Create_DC14_(_449_v30)
					var _451_v32 *C0
					_ = _451_v32
					var _nw71 *C0 = New_C0_()
					_ = _nw71
					_nw71.Ctor__(_443_v24, p0)
					_451_v32 = _nw71
					var _452_v33 _dafny.Map
					_ = _452_v33
					_452_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_451_v32, Companion_Default___.Fm19(p0, globalState))
					var _453_v34 _dafny.Sequence
					_ = _453_v34
					_453_v34 = _dafny.SeqOf(_443_v24, (_451_v32).F12(), (_451_v32).F12(), (_this).F16(), _443_v24)
					var _454_v35 _dafny.Sequence
					_ = _454_v35
					_454_v35 = _dafny.UnicodeSeqOfUtf8Bytes("ie")
					var _455_v36 _dafny.Map
					_ = _455_v36
					_455_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_451_v32).F12(), _454_v35)
					var _456_v37 _dafny.Map
					_ = _456_v37
					_456_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_443_v24, (_455_v36).Cardinality())
					var _457_v38 _dafny.Array
					_ = _457_v38
					var _nwElement0_11 _dafny.Int = p0
					_ = _nwElement0_11
					var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(24))
					_ = _nw72
					(_nw72).ArraySet1(_nwElement0_11, 0)
					(_nw72).ArraySet1((func() _dafny.Int {
						if (_445_v26).Contains(_dafny.IntOfInt64(452)) {
							return (_445_v26).Get(_dafny.IntOfInt64(452)).(_dafny.Int)
						}
						return (_447_v28).Cardinality()
					})(), 1)
					(_nw72).ArraySet1(p0, 2)
					(_nw72).ArraySet1(p0, 3)
					(_nw72).ArraySet1(p0, 4)
					(_nw72).ArraySet1((_dafny.Zero).Minus(p0), 5)
					(_nw72).ArraySet1(p0, 6)
					(_nw72).ArraySet1(p0, 7)
					(_nw72).ArraySet1(_dafny.IntOfInt64(689), 8)
					(_nw72).ArraySet1(((_450_v31).Dtor_cf23()).Cardinality(), 9)
					(_nw72).ArraySet1(p0, 10)
					(_nw72).ArraySet1(p0, 11)
					(_nw72).ArraySet1(p0, 12)
					(_nw72).ArraySet1((_452_v33).Cardinality(), 13)
					(_nw72).ArraySet1(p0, 14)
					(_nw72).ArraySet1((Companion_Default___.Fm20(globalState)).Cardinality(), 15)
					(_nw72).ArraySet1(_dafny.IntOfInt64(499), 16)
					(_nw72).ArraySet1(p0, 17)
					(_nw72).ArraySet1((_451_v32).F13(), 18)
					(_nw72).ArraySet1(p0, 19)
					(_nw72).ArraySet1(_dafny.IntOfUint32((_453_v34).Cardinality()), 20)
					(_nw72).ArraySet1((func() _dafny.Int {
						if (_456_v37).Contains(true) {
							return (_456_v37).Get(true).(_dafny.Int)
						}
						return Companion_Default___.Fm19(p0, globalState)
					})(), 21)
					(_nw72).ArraySet1((_451_v32).F13(), 22)
					(_nw72).ArraySet1(p0, 23)
					_457_v38 = _nw72
					var _458_v39 _dafny.Array
					_ = _458_v39
					_458_v39 = _457_v38
					var _459_v40 _dafny.Sequence
					_ = _459_v40
					_459_v40 = _dafny.SeqOf(_dafny.IntOfUint32((_454_v35).Cardinality()))
					var _rhs101 bool = _443_v24
					_ = _rhs101
					var _rhs102 bool = false
					_ = _rhs102
					var _rhs103 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_459_v40)).Difference(((_444_v25).Update(_dafny.IntOfUint32((_454_v35).Cardinality()), Companion_Default___.Abs(p0))).Intersection(_444_v25))
					_ = _rhs103
					var _rhs104 _dafny.Array = _446_v27
					_ = _rhs104
					_443_v24 = _rhs101
					_443_v24 = _rhs102
					_444_v25 = _rhs103
					_458_v39 = _rhs104
					(globalState).F2 = p0
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _460_i2 _dafny.Int
		_ = _460_i2
		_460_i2 = _dafny.Zero
		{
			for !(((_this).F16()) && (!((_this).F16()))) {
				{
					if (_460_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_460_i2 = (_460_i2).Plus(_dafny.One)
					(globalState).F2 = p0
					var _461_v41 *C0
					_ = _461_v41
					var _nw73 *C0 = New_C0_()
					_ = _nw73
					_nw73.Ctor__((_this).F16(), p0)
					_461_v41 = _nw73
					if (_this).F16() {
						var _462_v42 bool
						_ = _462_v42
						_462_v42 = false
						var _463_v43 _dafny.MultiSet
						_ = _463_v43
						_463_v43 = _dafny.MultiSetOf((_461_v41).F13())
						var _464_v44 _dafny.CodePoint
						_ = _464_v44
						_464_v44 = _dafny.CodePoint('o')
						_462_v42 = (_this).Fm4(((_461_v41).F13()).Times((_463_v43).Cardinality()), (p0).Cmp((_dafny.Zero).Minus((_461_v41).F13())) <= 0, _464_v44, globalState)
						_462_v42 = !((_461_v41).F12())
						var _465_v45 _dafny.Sequence
						_ = _465_v45
						_465_v45 = _dafny.SeqOf((_461_v41).F13(), (_461_v41).F13())
						var _466_v46 _dafny.Sequence
						_ = _466_v46
						_466_v46 = _dafny.SeqOf((_465_v45).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_465_v45).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm19((_461_v41).F13(), globalState), p0, Companion_Default___.Fm19((_461_v41).F13(), globalState))
						var _467_v47 _dafny.Array
						_ = _467_v47
						var _len0_4 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_4
						var _nw74 _dafny.Array
						_ = _nw74
						if _len0_4.Cmp(_dafny.Zero) == 0 {
							_nw74 = _dafny.NewArray(_len0_4)
						} else {
							var _init4 func(_dafny.Int) _dafny.Int = (func(_468_v41 *C0) func(_dafny.Int) _dafny.Int {
								return func(_469_i3 _dafny.Int) _dafny.Int {
									return (_469_i3).Plus((_468_v41).F13())
								}
							})(_461_v41)
							_ = _init4
							var _element0_4 = _init4(_dafny.Zero)
							_ = _element0_4
							_nw74 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
							(_nw74).ArraySet1(_element0_4, 0)
							var _nativeLen0_4 = (_len0_4).Int()
							_ = _nativeLen0_4
							for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
								(_nw74).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
							}
						}
						_467_v47 = _nw74
						var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_467_v47), 0))
						_ = _index86
						(_467_v47).ArraySet1((_461_v41).F13(), (_index86).Int())
						var _470_v48 _dafny.Array
						_ = _470_v48
						var _len0_5 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_5
						var _nw75 _dafny.Array
						_ = _nw75
						if _len0_5.Cmp(_dafny.Zero) == 0 {
							_nw75 = _dafny.NewArray(_len0_5)
						} else {
							var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_471_v42 bool, _472_v41 *C0) func(_dafny.Int) _dafny.MultiSet {
								return func(_473_i4 _dafny.Int) _dafny.MultiSet {
									return (_dafny.MultiSetOf(_471_v42)).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf((_472_v41).F12())))
								}
							})(_462_v42, _461_v41)
							_ = _init5
							var _element0_5 = _init5(_dafny.Zero)
							_ = _element0_5
							_nw75 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
							(_nw75).ArraySet1(_element0_5, 0)
							var _nativeLen0_5 = (_len0_5).Int()
							_ = _nativeLen0_5
							for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
								(_nw75).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
							}
						}
						_470_v48 = _nw75
						var _474_v49 _dafny.MultiSet
						_ = _474_v49
						_474_v49 = _dafny.MultiSetOf((_461_v41).F12())
						var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_470_v48), 0))
						_ = _index87
						(_470_v48).ArraySet1((_474_v49).Union(Companion_Default___.Fm21((_461_v41).F13(), false, (_461_v41).F13(), globalState)), (_index87).Int())
						var _475_v50 _dafny.Set
						_ = _475_v50
						_475_v50 = _dafny.SetOf(_dafny.IntOfInt64(53))
						var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_467_v47), 0))
						_ = _index88
						var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_470_v48), 0))
						_ = _index89
						var _rhs105 bool = !((_475_v50).IsDisjointFrom(_475_v50))
						_ = _rhs105
						var _rhs106 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_465_v45, _465_v45), _466_v46)
						_ = _rhs106
						var _rhs107 _dafny.Int = Companion_Default___.Fm19(((_463_v43).Cardinality()).Minus((_474_v49).Cardinality()), globalState)
						_ = _rhs107
						var _rhs108 _dafny.MultiSet = _474_v49
						_ = _rhs108
						var _lhs65 _dafny.Array = _467_v47
						_ = _lhs65
						var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_467_v47), 0))
						_ = _lhs66
						var _lhs67 _dafny.Array = _470_v48
						_ = _lhs67
						var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_470_v48), 0))
						_ = _lhs68
						_462_v42 = _rhs105
						_465_v45 = _rhs106
						(_lhs65).ArraySet1(_rhs107, (_lhs66).Int())
						(_lhs67).ArraySet1(_rhs108, (_lhs68).Int())
						var _476_v51 *C0
						_ = _476_v51
						var _nw76 *C0 = New_C0_()
						_ = _nw76
						_nw76.Ctor__((_461_v41).F12(), (_dafny.Zero).Minus((_467_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_467_v47), 0))).Int()).(_dafny.Int)))
						_476_v51 = _nw76
						var _477_v52 _dafny.Map
						_ = _477_v52
						_477_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v42, (_this).F16())
						var _478_v53 _dafny.Array
						_ = _478_v53
						var _nwElement0_12 _dafny.Map = Companion_Default___.Fm22(globalState)
						_ = _nwElement0_12
						var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(29))
						_ = _nw77
						(_nw77).ArraySet1(_nwElement0_12, 0)
						(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_461_v41).F12()), 1)
						(_nw77).ArraySet1((Companion_Default___.Fm22(globalState)).Merge(_477_v52), 2)
						(_nw77).ArraySet1((Companion_Default___.Fm22(globalState)).Merge(_477_v52), 3)
						(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_461_v41).F12()), 4)
						(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_476_v51).F12(), (_this).F16()), 5)
						(_nw77).ArraySet1(_477_v52, 6)
						(_nw77).ArraySet1(_477_v52, 7)
						(_nw77).ArraySet1(_477_v52, 8)
						(_nw77).ArraySet1(_477_v52, 9)
						(_nw77).ArraySet1((_477_v52).Merge(_477_v52), 10)
						(_nw77).ArraySet1(_477_v52, 11)
						(_nw77).ArraySet1(_477_v52, 12)
						(_nw77).ArraySet1(_477_v52, 13)
						(_nw77).ArraySet1(_477_v52, 14)
						(_nw77).ArraySet1((_477_v52).Merge(_477_v52), 15)
						(_nw77).ArraySet1(_477_v52, 16)
						(_nw77).ArraySet1(_477_v52, 17)
						(_nw77).ArraySet1(_477_v52, 18)
						(_nw77).ArraySet1(_477_v52, 19)
						(_nw77).ArraySet1((func() _dafny.Map {
							if (_476_v51).F12() {
								return _477_v52
							}
							return _477_v52
						})(), 20)
						(_nw77).ArraySet1(_477_v52, 21)
						(_nw77).ArraySet1((_477_v52).Merge(_477_v52), 22)
						(_nw77).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v41).F12(), (_461_v41).F12())).Merge(_477_v52), 23)
						(_nw77).ArraySet1((_477_v52).Merge(_477_v52), 24)
						(_nw77).ArraySet1((_477_v52).Update(_462_v42, (_this).F16()), 25)
						(_nw77).ArraySet1(_477_v52, 26)
						(_nw77).ArraySet1(_477_v52, 27)
						(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v41).F12(), (_476_v51).F12()), 28)
						_478_v53 = _nw77
						var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_478_v53), 0))
						_ = _index90
						(_478_v53).ArraySet1((_477_v52).Merge(_477_v52), (_index90).Int())
						var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_478_v53), 0))
						_ = _index91
						(_478_v53).ArraySet1(_477_v52, (_index91).Int())
					} else {
						r1 = p0
						var _479_v54 _dafny.Array
						_ = _479_v54
						var _nw78 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
						_ = _nw78
						_479_v54 = _nw78
						_479_v54 = _479_v54
						_479_v54 = _479_v54
						var _480_v55 bool
						_ = _480_v55
						_480_v55 = true
						var _481_v56 _dafny.CodePoint
						_ = _481_v56
						_481_v56 = _dafny.CodePoint('m')
						var _482_v57 _dafny.Sequence
						_ = _482_v57
						_482_v57 = _dafny.UnicodeSeqOfUtf8Bytes("m")
						_480_v55 = _dafny.Companion_Sequence_.Contains(_482_v57, _481_v56)
						var _483_v58 *C0
						_ = _483_v58
						var _nw79 *C0 = New_C0_()
						_ = _nw79
						_nw79.Ctor__(((_461_v41).F13()).Cmp(_dafny.IntOfInt64(734)) != 0, (p0).Plus(p0))
						_483_v58 = _nw79
					}
					if (_461_v41).F12() {
						var _484_v59 bool
						_ = _484_v59
						_484_v59 = true
						var _485_v60 _dafny.CodePoint
						_ = _485_v60
						_485_v60 = _dafny.CodePoint('v')
						var _486_v61 _dafny.Sequence
						_ = _486_v61
						_486_v61 = _dafny.SeqOf(_485_v60, _485_v60, _485_v60, _485_v60, _485_v60)
						_484_v59 = _dafny.Companion_Sequence_.IsPrefixOf(_486_v61, _486_v61)
						_485_v60 = _485_v60
						var _487_v62 _dafny.Array
						_ = _487_v62
						var _len0_6 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_6
						var _nw80 _dafny.Array
						_ = _nw80
						if _len0_6.Cmp(_dafny.Zero) == 0 {
							_nw80 = _dafny.NewArray(_len0_6)
						} else {
							var _init6 func(_dafny.Int) D7 = (func(_488_v61 _dafny.Sequence, _489_v41 *C0) func(_dafny.Int) D7 {
								return func(_490_i5 _dafny.Int) D7 {
									return Companion_D7_.Create_DC15_(_dafny.IntOfUint32((_488_v61).Cardinality()), (_489_v41).F12())
								}
							})(_486_v61, _461_v41)
							_ = _init6
							var _element0_6 = _init6(_dafny.Zero)
							_ = _element0_6
							_nw80 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
							(_nw80).ArraySet1(_element0_6, 0)
							var _nativeLen0_6 = (_len0_6).Int()
							_ = _nativeLen0_6
							for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
								(_nw80).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
							}
						}
						_487_v62 = _nw80
						var _491_v63 D7
						_ = _491_v63
						_491_v63 = Companion_D7_.Create_DC15_((_461_v41).F13(), (_this).F16())
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_487_v62), 0))
						_ = _index92
						(_487_v62).ArraySet1(_491_v63, (_index92).Int())
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_487_v62), 0))
						_ = _index93
						(_487_v62).ArraySet1(_491_v63, (_index93).Int())
						_484_v59 = (_461_v41).F12()
						var _rhs109 _dafny.Int = (p0).Times((_461_v41).F13())
						_ = _rhs109
						var _rhs110 bool = (Companion_Default___.Fm19((_461_v41).F13(), globalState)).Cmp(p0) >= 0
						_ = _rhs110
						var _rhs111 _dafny.Int = (_461_v41).F13()
						_ = _rhs111
						var _rhs112 _dafny.Int = (Companion_Default___.SafeModuloInt((_461_v41).F13(), (_461_v41).F13())).Plus((_461_v41).F13())
						_ = _rhs112
						var _lhs69 *GlobalState = globalState
						_ = _lhs69
						var _lhs70 *GlobalState = globalState
						_ = _lhs70
						_lhs69.F2 = _rhs109
						_484_v59 = _rhs110
						_lhs70.F2 = _rhs111
						r1 = _rhs112
					} else {
						var _492_v64 bool
						_ = _492_v64
						_492_v64 = true
						_492_v64 = _492_v64
						var _493_v65 _dafny.Map
						_ = _493_v65
						_493_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v41).F12(), (_461_v41).F12())
						var _494_v66 _dafny.Set
						_ = _494_v66
						_494_v66 = _dafny.SetOf(p0, (_461_v41).F13())
						var _495_v67 D3
						_ = _495_v67
						_495_v67 = Companion_D3_.Create_DC8_((_this).F16(), (_494_v66).Cardinality(), p0)
						var _496_v68 _dafny.MultiSet
						_ = _496_v68
						_496_v68 = _dafny.MultiSetOf((_495_v67).Dtor_cf15())
						var _497_v69 _dafny.Map
						_ = _497_v69
						_497_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v41).F13(), (_this).F16())
						var _498_v70 D3
						_ = _498_v70
						_498_v70 = Companion_D3_.Create_DC8_((_461_v41).F12(), (_461_v41).F13(), (func() _dafny.Int {
							if (_496_v68).Contains(p0) {
								return (_496_v68).Multiplicity(p0)
							}
							return (_497_v69).Cardinality()
						})())
						var _499_v71 _dafny.MultiSet
						_ = _499_v71
						_499_v71 = _dafny.MultiSetOf(((_493_v65).Cardinality()).Times((_498_v70).Dtor_cf14()), (_461_v41).F13())
						var _rhs113 _dafny.MultiSet = (_499_v71).Union(_dafny.MultiSetOf((_461_v41).F13()))
						_ = _rhs113
						var _rhs114 _dafny.Int = p0
						_ = _rhs114
						var _rhs115 _dafny.Int = Companion_Default___.Fm19((_461_v41).F13(), globalState)
						_ = _rhs115
						var _rhs116 bool = (_this).F16()
						_ = _rhs116
						var _lhs71 *GlobalState = globalState
						_ = _lhs71
						_499_v71 = _rhs113
						_lhs71.F2 = _rhs114
						r0 = _rhs115
						_492_v64 = _rhs116
						var _500_v72 _dafny.Int
						_ = _500_v72
						_500_v72 = (_461_v41).F13()
						var _501_v73 _dafny.MultiSet
						_ = _501_v73
						_501_v73 = _dafny.MultiSetOf(Companion_D1_.Create_DC3_((_461_v41).F12(), _500_v72), _400_v0)
						r0 = ((func() _dafny.Int {
							if (_501_v73).Contains(_400_v0) {
								return (_501_v73).Multiplicity(_400_v0)
							}
							return Companion_Default___.Fm19(p0, globalState)
						})()).Minus((_dafny.Zero).Minus((_461_v41).F13()))
						var _502_v74 _dafny.Array
						_ = _502_v74
						var _nw81 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
						_ = _nw81
						_502_v74 = _nw81
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_502_v74), 0))
						_ = _index94
						(_502_v74).ArraySet1((_461_v41).F12(), (_index94).Int())
						var _503_v75 _dafny.Sequence
						_ = _503_v75
						_503_v75 = _dafny.UnicodeSeqOfUtf8Bytes("eotb")
						var _504_v76 _dafny.Map
						_ = _504_v76
						_504_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_503_v75).Cardinality()), _495_v67)
						var _505_v78 _dafny.CodePoint
						_ = _505_v78
						_505_v78 = _dafny.CodePoint('d')
						var _506_v79 _dafny.Map
						_ = _506_v79
						_506_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v41).F13(), _505_v78)
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_502_v74), 0))
						_ = _index95
						var _rhs117 bool = (p0).Cmp(p0) < 0
						_ = _rhs117
						var _rhs118 _dafny.Map = ((func() _dafny.Map {
							if _492_v64 {
								return _504_v76
							}
							return _504_v76
						})()).Merge((_504_v76).Merge(func() _dafny.Map {
							var _coll15 = _dafny.NewMapBuilder()
							_ = _coll15
							for _iter15 := _dafny.Iterate((_506_v79).Keys().Elements()); ; {
								_compr_15, _ok15 := _iter15()
								if !_ok15 {
									break
								}
								var _507_v77 _dafny.Int
								_507_v77 = interface{}(_compr_15).(_dafny.Int)
								if (_506_v79).Contains(_507_v77) {
									_coll15.Add((_507_v77).Plus(_dafny.IntOfInt64(748)), _495_v67)
								}
							}
							return _coll15.ToMap()
						}()))
						_ = _rhs118
						var _lhs72 _dafny.Array = _502_v74
						_ = _lhs72
						var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_502_v74), 0))
						_ = _lhs73
						(_lhs72).ArraySet1(_rhs117, (_lhs73).Int())
						_504_v76 = _rhs118
						var _508_v80 *C0
						_ = _508_v80
						var _nw82 *C0 = New_C0_()
						_ = _nw82
						_nw82.Ctor__((_461_v41).F12(), (_461_v41).F13())
						_508_v80 = _nw82
					}
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _509_v81 bool
		_ = _509_v81
		_509_v81 = true
		var _rhs119 bool = (_this).F16()
		_ = _rhs119
		var _rhs120 _dafny.Int = p0
		_ = _rhs120
		var _lhs74 *GlobalState = globalState
		_ = _lhs74
		_509_v81 = _rhs119
		_lhs74.F2 = _rhs120
		var _hi2 _dafny.Int = p0
		_ = _hi2
		for _510_i6 := p0; _510_i6.Cmp(_hi2) < 0; _510_i6 = _510_i6.Plus(_dafny.One) {
			var _511_v82 _dafny.Int
			_ = _511_v82
			var _512_v83 bool
			_ = _512_v83
			var _513_v84 bool
			_ = _513_v84
			var _514_v85 _dafny.Map
			_ = _514_v85
			var _out42 _dafny.Int
			_ = _out42
			var _out43 bool
			_ = _out43
			var _out44 bool
			_ = _out44
			var _out45 _dafny.Map
			_ = _out45
			_out42, _out43, _out44, _out45 = Companion_Default___.M0(globalState)
			_511_v82 = _out42
			_512_v83 = _out43
			_513_v84 = _out44
			_514_v85 = _out45
			var _515_v86 _dafny.Array
			_ = _515_v86
			var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw83
			_515_v86 = _nw83
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_515_v86), 0))
			_ = _index96
			(_515_v86).ArraySet1(Companion_Default___.Fm19((_dafny.Zero).Minus(p0), globalState), (_index96).Int())
			var _516_v87 _dafny.Array
			_ = _516_v87
			var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw84
			_516_v87 = _nw84
			var _517_v88 _dafny.Sequence
			_ = _517_v88
			_517_v88 = _dafny.UnicodeSeqOfUtf8Bytes("jasbi")
			var _518_v89 _dafny.Sequence
			_ = _518_v89
			_518_v89 = _517_v88
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_516_v87), 0))
			_ = _index97
			(_516_v87).ArraySet1(_518_v89, (_index97).Int())
			var _519_v90 _dafny.Map
			_ = _519_v90
			_519_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_i6, (_this).F16())
			var _520_v91 _dafny.Sequence
			_ = _520_v91
			_520_v91 = _dafny.SeqOf(_512_v83)
			var _521_v92 _dafny.Map
			_ = _521_v92
			_521_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_i6, p0)
			var _522_v93 _dafny.MultiSet
			_ = _522_v93
			_522_v93 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_521_v92).Contains(_510_i6) {
					return (_521_v92).Get(_510_i6).(_dafny.Int)
				}
				return _dafny.IntOfInt64(267)
			})())
			var _523_v94 _dafny.MultiSet
			_ = _523_v94
			_523_v94 = _dafny.MultiSetOf(_509_v81)
			var _524_v95 _dafny.Sequence
			_ = _524_v95
			_524_v95 = _dafny.SeqOf(_523_v94)
			var _525_v96 _dafny.Map
			_ = _525_v96
			_525_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_517_v88, _509_v81)
			var _526_v97 D3
			_ = _526_v97
			_526_v97 = Companion_D3_.Create_DC6_(_511_v82, (func() bool {
				if (_519_v90).Contains(_dafny.IntOfUint32((_520_v91).Cardinality())) {
					return (_519_v90).Get(_dafny.IntOfUint32((_520_v91).Cardinality())).(bool)
				}
				return false
			})(), !(Companion_Default___.Fm0((_522_v93).Cardinality(), _dafny.IntOfUint32((_520_v91).Cardinality()), _524_v95, _525_v96, globalState)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_511_v82, (_this).F16())).Cardinality(), true)
			var _527_v98 _dafny.Sequence
			_ = _527_v98
			_527_v98 = _dafny.SeqOf(_511_v82, _510_i6)
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_515_v86), 0))
			_ = _index98
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_516_v87), 0))
			_ = _index99
			var _rhs121 _dafny.Int = (_526_v97).Dtor_cf7()
			_ = _rhs121
			var _rhs122 _dafny.Int = _510_i6
			_ = _rhs122
			var _rhs123 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_527_v98, (func() _dafny.Sequence {
				if _513_v84 {
					return _527_v98
				}
				return _527_v98
			})())).Cardinality())
			_ = _rhs123
			var _rhs124 _dafny.Int = p0
			_ = _rhs124
			var _rhs125 _dafny.Sequence = _518_v89
			_ = _rhs125
			var _lhs75 _dafny.Array = _515_v86
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_515_v86), 0))
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			var _lhs79 _dafny.Array = _516_v87
			_ = _lhs79
			var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_516_v87), 0))
			_ = _lhs80
			_511_v82 = _rhs121
			(_lhs75).ArraySet1(_rhs122, (_lhs76).Int())
			_lhs77.F2 = _rhs123
			_lhs78.F2 = _rhs124
			(_lhs79).ArraySet1(_rhs125, (_lhs80).Int())
			_513_v84 = (p0).Cmp(Companion_Default___.SafeDivisionInt(_511_v82, _dafny.IntOfInt64(965))) >= 0
			var _528_v99 _dafny.CodePoint
			_ = _528_v99
			_528_v99 = _dafny.CodePoint('e')
			var _529_v100 _dafny.Map
			_ = _529_v100
			_529_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_509_v81), false)
			_528_v99 = (func() _dafny.CodePoint {
				if (func() bool {
					if (_529_v100).Contains(_513_v84) {
						return (_529_v100).Get(_513_v84).(bool)
					}
					return true
				})() {
					return _528_v99
				}
				return _dafny.CodePoint('m')
			})()
		}
		var _530_v101 _dafny.Sequence
		_ = _530_v101
		_530_v101 = _dafny.UnicodeSeqOfUtf8Bytes("tri")
		r0 = _dafny.IntOfUint32((_530_v101).Cardinality())
		r0 = p0
		r1 = p0
		return r0, r1
	}
}
func (_this *C1) F16() bool {
	{
		return _this._f16
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F14  bool
	_f15 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F14 = false
	_this._f15 = false
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

func (_this *C2) Ctor__(f14 bool, f15 bool) {
	{
		(_this).F14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C2) Fm3(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(531)).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F15(), _this.F14, false)).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_531_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _dafny.IntOfInt64(730))))
	}
}
func (_this *C2) Fm4(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return _this.F14
	}
}
func (_this *C2) Fm12(p0 bool, p1 _dafny.Int, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf((_dafny.MultiSetOf(_this.F14)).Cardinality(), _dafny.IntOfInt64(73))).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_532_i0 _dafny.Int) _dafny.Int {
			return (_dafny.SetOf(_dafny.SeqOf((_dafny.SetOf(_this.F14)).Cardinality()))).Cardinality()
		}))).Cardinality())))
	}
}
func (_this *C2) Fm13(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(406)).Cmp((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lgislaabn")).Cardinality())).Minus(_dafny.IntOfInt64(169))) > 0
	}
}
func (_this *C2) M3(p0 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _533_v0 _dafny.Int
		_ = _533_v0
		_533_v0 = _dafny.IntOfInt64(55)
		var _534_v1 *C0
		_ = _534_v1
		var _nw85 *C0 = New_C0_()
		_ = _nw85
		_nw85.Ctor__((_this).F15(), _533_v0)
		_534_v1 = _nw85
		var _535_v2 _dafny.Sequence
		_ = _535_v2
		_535_v2 = _dafny.SeqOf(_534_v1, _534_v1)
		var _536_v3 _dafny.MultiSet
		_ = _536_v3
		_536_v3 = _dafny.MultiSetOf(!((_this).F15()))
		var _537_v4 D3
		_ = _537_v4
		_537_v4 = Companion_D3_.Create_DC7_(_533_v0)
		var _538_v5 _dafny.Map
		_ = _538_v5
		_538_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _537_v4)
		var _539_v6 _dafny.Map
		_ = _539_v6
		_539_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _533_v0)
		var _540_v7 _dafny.Sequence
		_ = _540_v7
		_540_v7 = _dafny.UnicodeSeqOfUtf8Bytes("vedhrd")
		var _541_v8 _dafny.Map
		_ = _541_v8
		_541_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_534_v1).F13(), (_539_v6).Cardinality())
		var _542_v10 _dafny.Set
		_ = _542_v10
		_542_v10 = _dafny.SetOf((_534_v1).F13(), (_534_v1).F13())
		var _543_v11 _dafny.CodePoint
		_ = _543_v11
		_543_v11 = _dafny.CodePoint('i')
		var _544_v12 _dafny.Sequence
		_ = _544_v12
		_544_v12 = _dafny.SeqOf((_542_v10).Cardinality(), (_534_v1).F13(), (_dafny.MultiSetOf(_543_v11, _543_v11)).Cardinality(), _533_v0)
		var _545_v13 _dafny.Array
		_ = _545_v13
		var _nwElement0_13 _dafny.Int = (_533_v0).Minus(_533_v0)
		_ = _nwElement0_13
		var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(26))
		_ = _nw86
		(_nw86).ArraySet1(_nwElement0_13, 0)
		(_nw86).ArraySet1(_533_v0, 1)
		(_nw86).ArraySet1(_533_v0, 2)
		(_nw86).ArraySet1(_dafny.IntOfInt64(-457), 3)
		(_nw86).ArraySet1(_533_v0, 4)
		(_nw86).ArraySet1((_dafny.Zero).Minus(_533_v0), 5)
		(_nw86).ArraySet1(_533_v0, 6)
		(_nw86).ArraySet1((_533_v0).Minus(_dafny.IntOfInt64(452)), 7)
		(_nw86).ArraySet1(_533_v0, 8)
		(_nw86).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_535_v2).Select((Companion_Default___.SafeIndex((_534_v1).F13(), _dafny.IntOfUint32((_535_v2).Cardinality()))).Uint32()).(*C0), (_534_v1).F13())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v1, _533_v0))).Cardinality(), 9)
		(_nw86).ArraySet1(((_536_v3).Update((_534_v1).F12(), Companion_Default___.Abs(_533_v0))).Cardinality(), 10)
		(_nw86).ArraySet1((_534_v1).F13(), 11)
		(_nw86).ArraySet1((Companion_Default___.Fm14(_538_v5, _this.F14, _539_v6, (_this).F15(), globalState)).Cardinality(), 12)
		(_nw86).ArraySet1(_533_v0, 13)
		(_nw86).ArraySet1(((_dafny.MultiSetOf((_this).F15(), (_this).F15())).Cardinality()).Plus(_dafny.IntOfUint32((_540_v7).Cardinality())), 14)
		(_nw86).ArraySet1((_dafny.Zero).Minus(_533_v0), 15)
		(_nw86).ArraySet1((_534_v1).F13(), 16)
		(_nw86).ArraySet1((((_541_v8).Update((_534_v1).F13(), (_534_v1).F13())).Update((func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_534_v1).F13())).Keys().Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _546_v9 _dafny.Sequence
				_546_v9 = interface{}(_compr_16).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_534_v1).F13())).Contains(_546_v9) {
					_coll16.Add(_546_v9, false)
				}
			}
			return _coll16.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(-304))).Cardinality(), 17)
		(_nw86).ArraySet1((_534_v1).F13(), 18)
		(_nw86).ArraySet1((_534_v1).F13(), 19)
		(_nw86).ArraySet1(((_534_v1).F13()).Minus(_dafny.IntOfUint32((_544_v12).Cardinality())), 20)
		(_nw86).ArraySet1(_dafny.IntOfInt64(-752), 21)
		(_nw86).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(117), (_534_v1).F13()), 22)
		(_nw86).ArraySet1(((func() _dafny.MultiSet {
			if (_this).F15() {
				return _536_v3
			}
			return _536_v3
		})()).Cardinality(), 23)
		(_nw86).ArraySet1(_533_v0, 24)
		(_nw86).ArraySet1((func() _dafny.Int {
			if _this.F14 {
				return _dafny.IntOfInt64(293)
			}
			return _533_v0
		})(), 25)
		_545_v13 = _nw86
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_545_v13), 0))); ; {
			_guard_loop_0, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _547_i0 _dafny.Int
			_547_i0 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_547_i0).Sign() != -1) && ((_547_i0).Cmp(_dafny.ArrayLen((_545_v13), 0)) < 0)) {
				(_545_v13).ArraySet1((_547_i0).Minus((_dafny.Zero).Minus(_533_v0)), (_547_i0).Int())
			}
		}
		var _548_v14 _dafny.Array
		_ = _548_v14
		var _nw87 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
		_ = _nw87
		_548_v14 = _nw87
		_548_v14 = _548_v14
		var _549_v15 *C0
		_ = _549_v15
		var _nw88 *C0 = New_C0_()
		_ = _nw88
		_nw88.Ctor__(true, Companion_Default___.SafeDivisionInt(_533_v0, _533_v0))
		_549_v15 = _nw88
		var _550_v16 _dafny.Map
		_ = _550_v16
		_550_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_549_v15).F13(), false)
		var _551_v17 _dafny.Sequence
		_ = _551_v17
		_551_v17 = _dafny.SeqOf(_550_v16, (_550_v16).Merge(_550_v16), _550_v16)
		_551_v17 = _551_v17
		var _552_v18 _dafny.Map
		_ = _552_v18
		_552_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), Companion_D3_.Create_DC8_(_this.F14, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_544_v12, (Companion_Default___.SafeIndex((_549_v15).F13(), _dafny.IntOfUint32((_544_v12).Cardinality()))).Uint32(), (_534_v1).F13())).Cardinality()), _533_v0))
		var _553_v19 D3
		_ = _553_v19
		_553_v19 = Companion_D3_.Create_DC8_((_this).Fm4((_549_v15).F13(), (_549_v15).F12(), _543_v11, globalState), _533_v0, (_549_v15).F13())
		var _554_v20 _dafny.Map
		_ = _554_v20
		_554_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, (_549_v15).F12())
		var _555_v21 _dafny.Set
		_ = _555_v21
		_555_v21 = _dafny.SetOf(true, (_549_v15).F12())
		var _556_v22 _dafny.Array
		_ = _556_v22
		var _nwElement0_14 _dafny.Map = (_552_v18).Merge(_552_v18)
		_ = _nwElement0_14
		var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(25))
		_ = _nw89
		(_nw89).ArraySet1(_nwElement0_14, 0)
		(_nw89).ArraySet1((_552_v18).Merge(_552_v18), 1)
		(_nw89).ArraySet1(_552_v18, 2)
		(_nw89).ArraySet1(_552_v18, 3)
		(_nw89).ArraySet1(_552_v18, 4)
		(_nw89).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _553_v19)).Merge(_552_v18), 5)
		(_nw89).ArraySet1(((_552_v18).Update(!((_534_v1).F12()), _553_v19)).Merge(_552_v18), 6)
		(_nw89).ArraySet1((Companion_Default___.Fm15(globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_534_v1).F12(), Companion_D3_.Create_DC8_(_this.F14, (_549_v15).F13(), _dafny.IntOfInt64(-421)))), 7)
		(_nw89).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_549_v15).F12(), Companion_Default___.Fm16(globalState)), 8)
		(_nw89).ArraySet1(_552_v18, 9)
		(_nw89).ArraySet1(_552_v18, 10)
		(_nw89).ArraySet1(_552_v18, 11)
		(_nw89).ArraySet1((_552_v18).Merge((_552_v18).Update(true, Companion_D3_.Create_DC8_((_549_v15).F12(), _533_v0, _533_v0))), 12)
		(_nw89).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_549_v15).F12(), Companion_D3_.Create_DC8_(false, Companion_Default___.Fm17((func() bool {
			if (_554_v20).Contains(_this.F14) {
				return (_554_v20).Get(_this.F14).(bool)
			}
			return _this.F14
		})(), _533_v0, (_549_v15).F12(), (_534_v1).F13(), globalState), (_555_v21).Cardinality())), 13)
		(_nw89).ArraySet1((_552_v18).Merge(_552_v18), 14)
		(_nw89).ArraySet1(_552_v18, 15)
		(_nw89).ArraySet1((_552_v18).Merge(_552_v18), 16)
		(_nw89).ArraySet1((_552_v18).Merge(_552_v18), 17)
		(_nw89).ArraySet1(((Companion_D6_.Create_DC12_(_552_v18)).Dtor_cf21()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_534_v1).F12(), _553_v19)), 18)
		(_nw89).ArraySet1(_552_v18, 19)
		(_nw89).ArraySet1((_552_v18).Update((_this).F15(), _553_v19), 20)
		(_nw89).ArraySet1(_552_v18, 21)
		(_nw89).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_534_v1).F12(), _553_v19)).Update((_this).F15(), _553_v19), 22)
		(_nw89).ArraySet1(_552_v18, 23)
		(_nw89).ArraySet1(_552_v18, 24)
		_556_v22 = _nw89
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_556_v22), 0))); ; {
			_guard_loop_1, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _557_i1 _dafny.Int
			_557_i1 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_557_i1).Sign() != -1) && ((_557_i1).Cmp(_dafny.ArrayLen((_556_v22), 0)) < 0)) {
				(_556_v22).ArraySet1(_552_v18, (_557_i1).Int())
			}
		}
		var _558_v23 _dafny.Sequence
		_ = _558_v23
		_558_v23 = _540_v7
		_558_v23 = _540_v7
		r0 = (_534_v1).F12()
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw90
		r1 = _nw90
		return r0, r1
	}
}
func (_this *C2) M4(p0 bool, p1 _dafny.Int, globalState *GlobalState) {
	{
		(globalState).F2 = (p1).Plus(p1)
		var _559_v0 _dafny.CodePoint
		_ = _559_v0
		_559_v0 = _dafny.CodePoint('v')
		var _560_v1 _dafny.Map
		_ = _560_v1
		_560_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
		var _561_v2 _dafny.Map
		_ = _561_v2
		_561_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_560_v1).Contains(p0) {
				return (_560_v1).Get(p0).(_dafny.Int)
			}
			return p1
		})(), _559_v0)
		var _562_v3 _dafny.MultiSet
		_ = _562_v3
		_562_v3 = _dafny.MultiSetOf((_this).F15())
		_559_v0 = (func() _dafny.CodePoint {
			if (_561_v2).Contains((_dafny.IntOfInt64(692)).Times((func() _dafny.Int {
				if (_562_v3).Contains(_this.F14) {
					return (_562_v3).Multiplicity(_this.F14)
				}
				return _dafny.IntOfInt64(648)
			})())) {
				return (_561_v2).Get((_dafny.IntOfInt64(692)).Times((func() _dafny.Int {
					if (_562_v3).Contains(_this.F14) {
						return (_562_v3).Multiplicity(_this.F14)
					}
					return _dafny.IntOfInt64(648)
				})())).(_dafny.CodePoint)
			}
			return _559_v0
		})()
		var _563_v4 _dafny.Int
		_ = _563_v4
		_563_v4 = p1
		var _564_i0 _dafny.Int
		_ = _564_i0
		_564_i0 = _dafny.Zero
		{
			for func(_source12 _dafny.Int) bool {
				var _574___mcc_h0 _dafny.Int = _source12
				_ = _574___mcc_h0
				var _575_cf0 _dafny.Int = _574___mcc_h0
				_ = _575_cf0
				return _this.F14
			}(_563_v4) {
				{
					if (_564_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_564_i0 = (_564_i0).Plus(_dafny.One)
					var _565_v5 _dafny.MultiSet
					_ = _565_v5
					_565_v5 = _dafny.MultiSetOf(_dafny.IntOfInt64(685), (_562_v3).Cardinality(), _dafny.IntOfInt64(525), p1)
					var _566_v6 _dafny.Array
					_ = _566_v6
					var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
					_ = _nw91
					_566_v6 = _nw91
					var _567_v7 D3
					_ = _567_v7
					_567_v7 = Companion_D3_.Create_DC5_(_566_v6)
					var _568_v8 T0
					_ = _568_v8
					var _nw92 *C1 = New_C1_()
					_ = _nw92
					_nw92.Ctor__(_this.F14, _567_v7)
					_568_v8 = _nw92
					var _569_v9 _dafny.Map
					_ = _569_v9
					_569_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _568_v8)
					(globalState).F2 = (((func() _dafny.Int {
						if (_565_v5).Contains((_569_v9).Cardinality()) {
							return (_565_v5).Multiplicity((_569_v9).Cardinality())
						}
						return p1
					})()).Plus(p1)).Times((_dafny.Zero).Minus(p1))
					(_this).F14 = p0
					var _570_v10 _dafny.Sequence
					_ = _570_v10
					_570_v10 = _dafny.UnicodeSeqOfUtf8Bytes("e")
					var _571_v11 _dafny.Map
					_ = _571_v11
					_571_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0) || (p0), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_570_v10).Cardinality()))).Cmp(p1) != 0)
					_571_v11 = Companion_Default___.Fm22(globalState)
					var _572_v12 _dafny.Array
					_ = _572_v12
					var _nwElement0_15 bool = (_this).F15()
					_ = _nwElement0_15
					var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(3))
					_ = _nw93
					(_nw93).ArraySet1(_nwElement0_15, 0)
					(_nw93).ArraySet1(!(p0), 1)
					(_nw93).ArraySet1(p0, 2)
					_572_v12 = _nw93
					var _573_v13 _dafny.MultiSet
					_ = _573_v13
					_573_v13 = _dafny.MultiSetOf(_572_v12)
					_573_v13 = (_573_v13).Difference((_573_v13).Update(_572_v12, Companion_Default___.Abs((_560_v1).Cardinality())))
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _576_i1 _dafny.Int
		_ = _576_i1
		_576_i1 = _dafny.Zero
		{
			for p0 {
				{
					if (_576_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_576_i1 = (_576_i1).Plus(_dafny.One)
					(_this).F14 = (_this).F15()
					var _577_v14 _dafny.Set
					_ = _577_v14
					_577_v14 = _dafny.SetOf(_559_v0, _dafny.CodePoint('y'))
					var _578_v15 _dafny.Map
					_ = _578_v15
					_578_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), p0)
					var _579_v16 *C0
					_ = _579_v16
					var _nw94 *C0 = New_C0_()
					_ = _nw94
					_nw94.Ctor__(p0, (_578_v15).Cardinality())
					_579_v16 = _nw94
					var _580_v17 _dafny.Sequence
					_ = _580_v17
					_580_v17 = _dafny.SeqOf(_579_v16)
					var _581_v18 D3
					_ = _581_v18
					_581_v18 = Companion_D3_.Create_DC8_((_this).F15(), (_579_v16).F13(), (_579_v16).F13())
					var _582_v19 _dafny.Array
					_ = _582_v19
					var _nwElement0_16 bool = (_this).F15()
					_ = _nwElement0_16
					var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(20))
					_ = _nw95
					(_nw95).ArraySet1(_nwElement0_16, 0)
					(_nw95).ArraySet1(_this.F14, 1)
					(_nw95).ArraySet1(true, 2)
					(_nw95).ArraySet1(p0, 3)
					(_nw95).ArraySet1(p0, 4)
					(_nw95).ArraySet1((p1).Cmp(p1) > 0, 5)
					(_nw95).ArraySet1((_this).F15(), 6)
					(_nw95).ArraySet1((_this).F15(), 7)
					(_nw95).ArraySet1((_this).Fm4(p1, p0, _559_v0, globalState), 8)
					(_nw95).ArraySet1(false, 9)
					(_nw95).ArraySet1((_577_v14).IsDisjointFrom(_577_v14), 10)
					(_nw95).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_579_v16, _579_v16), _580_v17), 11)
					(_nw95).ArraySet1(p0, 12)
					(_nw95).ArraySet1((true) && ((_579_v16).F12()), 13)
					(_nw95).ArraySet1(true, 14)
					(_nw95).ArraySet1((_this).F15(), 15)
					(_nw95).ArraySet1((_579_v16).F12(), 16)
					(_nw95).ArraySet1((_579_v16).F12(), 17)
					(_nw95).ArraySet1((func() bool {
						if (_581_v18).Dtor_cf13() {
							return (_this).F15()
						}
						return (_579_v16).F12()
					})(), 18)
					(_nw95).ArraySet1(_this.F14, 19)
					_582_v19 = _nw95
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_582_v19), 0))
					_ = _index100
					(_582_v19).ArraySet1((_this).F15(), (_index100).Int())
					var _583_v20 D7
					_ = _583_v20
					_583_v20 = Companion_D7_.Create_DC15_((_579_v16).F13(), _this.F14)
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_582_v19), 0))
					_ = _index101
					(_582_v19).ArraySet1((func() bool {
						if !(_this.F14) {
							return _this.F14
						}
						return ((_583_v20).Dtor_cf24()).Cmp(p1) <= 0
					})(), (_index101).Int())
					var _584_v21 _dafny.Map
					_ = _584_v21
					_584_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_579_v16).F13(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-457))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}((func(_585_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_586_i2 _dafny.Int) _dafny.CodePoint {
							return _585_v0
						}
					})(_559_v0)))).Cardinality()))
					var _587_v22 _dafny.Set
					_ = _587_v22
					_587_v22 = _dafny.SetOf((_579_v16).F13(), (_dafny.Zero).Minus((func() _dafny.Int {
						if (_584_v21).Contains((_579_v16).F13()) {
							return (_584_v21).Get((_579_v16).F13()).(_dafny.Int)
						}
						return _dafny.IntOfInt64(-421)
					})()))
					var _588_v23 _dafny.Array
					_ = _588_v23
					var _nwElement0_17 _dafny.Array = _582_v19
					_ = _nwElement0_17
					var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(19))
					_ = _nw96
					(_nw96).ArraySet1(_nwElement0_17, 0)
					(_nw96).ArraySet1(_582_v19, 1)
					(_nw96).ArraySet1(_582_v19, 2)
					(_nw96).ArraySet1(_582_v19, 3)
					(_nw96).ArraySet1(_582_v19, 4)
					(_nw96).ArraySet1(_582_v19, 5)
					(_nw96).ArraySet1(_582_v19, 6)
					(_nw96).ArraySet1(_582_v19, 7)
					(_nw96).ArraySet1(_582_v19, 8)
					(_nw96).ArraySet1(_582_v19, 9)
					(_nw96).ArraySet1(_582_v19, 10)
					(_nw96).ArraySet1(_582_v19, 11)
					(_nw96).ArraySet1(_582_v19, 12)
					(_nw96).ArraySet1(_582_v19, 13)
					(_nw96).ArraySet1(_582_v19, 14)
					(_nw96).ArraySet1(_582_v19, 15)
					(_nw96).ArraySet1(_582_v19, 16)
					(_nw96).ArraySet1(_582_v19, 17)
					(_nw96).ArraySet1(_582_v19, 18)
					_588_v23 = _nw96
					var _589_v24 D3
					_ = _589_v24
					_589_v24 = Companion_D3_.Create_DC5_(_588_v23)
					var _590_v25 T0
					_ = _590_v25
					var _nw97 *C1 = New_C1_()
					_ = _nw97
					_nw97.Ctor__(false, _589_v24)
					_590_v25 = _nw97
					var _591_v26 _dafny.Map
					_ = _591_v26
					_591_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_590_v25, _587_v22)
					_587_v22 = (func() _dafny.Set {
						if (_591_v26).Contains(_590_v25) {
							return (_591_v26).Get(_590_v25).(_dafny.Set)
						}
						return (_587_v22).Intersection(_587_v22)
					})()
					var _592_v28 _dafny.Sequence
					_ = _592_v28
					_592_v28 = _dafny.SeqOf(p1)
					var _593_v29 _dafny.Sequence
					_ = _593_v29
					_593_v29 = _dafny.SeqOf(_562_v3)
					var _594_v30 _dafny.Sequence
					_ = _594_v30
					_594_v30 = _dafny.UnicodeSeqOfUtf8Bytes("leigq")
					var _595_v31 _dafny.Map
					_ = _595_v31
					_595_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_594_v30, (_582_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_582_v19), 0))).Int()).(bool))
					var _596_v32 _dafny.Set
					_ = _596_v32
					_596_v32 = _dafny.SetOf(_this.F14, Companion_Default___.Fm0(p1, (func() _dafny.Map {
						var _coll17 = _dafny.NewMapBuilder()
						_ = _coll17
						for _iter19 := _dafny.Iterate((_592_v28).Elements()); ; {
							_compr_17, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _597_v27 _dafny.Int
							_597_v27 = interface{}(_compr_17).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_592_v28, _597_v27) {
								_coll17.Add((_597_v27).Times((_579_v16).F13()), (_this).F15())
							}
						}
						return _coll17.ToMap()
					}()).Cardinality(), _593_v29, _595_v31, globalState))
					var _598_v33 _dafny.Sequence
					_ = _598_v33
					_598_v33 = _dafny.SeqOf(_596_v32)
					var _599_v34 _dafny.MultiSet
					_ = _599_v34
					_599_v34 = _dafny.MultiSetOf(((_562_v3).Cardinality()).Times((_596_v32).Cardinality()), p1, p1, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(((_598_v33).Select((Companion_Default___.SafeIndex((_579_v16).F13(), _dafny.IntOfUint32((_598_v33).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()), p1), (p1).Minus(p1))
					(globalState).F2 = (_dafny.Zero).Minus((func() _dafny.Int {
						if (_599_v34).Contains(p1) {
							return (_599_v34).Multiplicity(p1)
						}
						return p1
					})())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _600_v35 _dafny.Sequence
		_ = _600_v35
		_600_v35 = _dafny.SeqOf(_this.F14, _this.F14)
		var _601_v36 bool
		_ = _601_v36
		var _602_v37 _dafny.Array
		_ = _602_v37
		var _out46 bool
		_ = _out46
		var _out47 _dafny.Array
		_ = _out47
		_out46, _out47 = (_this).M3(_600_v35, globalState)
		_601_v36 = _out46
		_602_v37 = _out47
		var _603_v38 D4
		_ = _603_v38
		_603_v38 = Companion_D4_.Create_DC9_((_562_v3).Update(_601_v36, Companion_Default___.Abs(p1)))
		_603_v38 = (func() D4 {
			if p0 {
				return _603_v38
			}
			return _603_v38
		})()
	}
}
func (_this *C2) F15() bool {
	{
		return _this._f15
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f10 _dafny.Int
	_f11 _dafny.Map
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.EmptyMap
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

func (_this *C3) F10() _dafny.Int {
	return _this._f10
}
func (_this *C3) F10_set_(value _dafny.Int) {
	_this._f10 = value
}
func (_this *C3) F11() _dafny.Map {
	return _this._f11
}
func (_this *C3) Ctor__(f10 _dafny.Int, f11 _dafny.Map) {
	{
		(_this)._f10 = f10
		(_this)._f11 = f11
	}
}
func (_this *C3) Fm5(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) bool {
	{
		return !(true) || ((_dafny.SetOf(_dafny.IntOfInt64(246))).IsProperSubsetOf(func() _dafny.Set {
			var _coll18 = _dafny.NewBuilder()
			_ = _coll18
			for _iter20 := _dafny.Iterate((_dafny.SetOf(_this.F10())).Elements()); ; {
				_compr_18, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _604_v0 _dafny.Int
				_604_v0 = interface{}(_compr_18).(_dafny.Int)
				if (_dafny.SetOf(_this.F10())).Contains(_604_v0) {
					_coll18.Add(Companion_Default___.SafeModuloInt(_604_v0, _dafny.IntOfInt64(777)))
				}
			}
			return _coll18.ToSet()
		}()))
	}
}
func (_this *C3) Fm6(p0 _dafny.Set, globalState *GlobalState) D1 {
	{
		var _source13 D1 = Companion_D1_.Create_DC2_(!(true))
		_ = _source13
		if _source13.Is_DC2() {
			var _605___mcc_h0 bool = _source13.Get_().(D1_DC2).Cf2
			_ = _605___mcc_h0
			var _606_cf2 bool = _605___mcc_h0
			_ = _606_cf2
			return Companion_D1_.Create_DC2_(_606_cf2)
		} else if _source13.Is_DC3() {
			var _607___mcc_h1 bool = _source13.Get_().(D1_DC3).Cf3
			_ = _607___mcc_h1
			var _608___mcc_h2 _dafny.Int = _source13.Get_().(D1_DC3).Cf4
			_ = _608___mcc_h2
			var _609_cf4 _dafny.Int = _608___mcc_h2
			_ = _609_cf4
			var _610_cf3 bool = _607___mcc_h1
			_ = _610_cf3
			return Companion_D1_.Create_DC2_(_610_cf3)
		} else {
			var _611___mcc_h3 bool = _source13.Get_().(D1_DC1).Cf1
			_ = _611___mcc_h3
			var _612_cf1 bool = _611___mcc_h3
			_ = _612_cf1
			return Companion_D1_.Create_DC2_(false)
		}
	}
}
func (_this *C3) Fm3(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F10()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F11()).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F10())))
	}
}
func (_this *C3) Fm4(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (_this.F10()).Cmp((_dafny.Zero).Minus(_this.F10())) >= 0
	}
}
func (_this *C3) Fm7(p0 _dafny.Map, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((func() bool {
			if true {
				return (func() bool {
					if !(true) {
						return !(true)
					}
					return false
				})()
			}
			return false
		})())
	}
}
func (_this *C3) M1(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _613_v0 bool
		_ = _613_v0
		_613_v0 = true
		_613_v0 = (func() bool {
			if _613_v0 {
				return _613_v0
			}
			return _613_v0
		})()
		var _614_v1 _dafny.Sequence
		_ = _614_v1
		_614_v1 = _dafny.SeqOf(_this.F10())
		var _615_v2 _dafny.Set
		_ = _615_v2
		_615_v2 = _dafny.SetOf(_this.F10(), _this.F10(), _this.F10())
		if (func() bool {
			if _613_v0 {
				return ((_614_v1).Select((Companion_Default___.SafeIndex((_615_v2).Cardinality(), _dafny.IntOfUint32((_614_v1).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_this.F10()) <= 0
			}
			return _613_v0
		})() {
			var _616_v3 *C0
			_ = _616_v3
			var _nw98 *C0 = New_C0_()
			_ = _nw98
			_nw98.Ctor__((_this.F10()).Cmp((_dafny.Zero).Minus(_this.F10())) >= 0, (_dafny.Zero).Minus(_this.F10()))
			_616_v3 = _nw98
			if (_613_v0) || ((_616_v3).F12()) {
				(_this).F10_set_((_616_v3).F13())
				var _617_v4 _dafny.Int
				_ = _617_v4
				var _618_v5 bool
				_ = _618_v5
				var _619_v6 bool
				_ = _619_v6
				var _620_v7 _dafny.Map
				_ = _620_v7
				var _out48 _dafny.Int
				_ = _out48
				var _out49 bool
				_ = _out49
				var _out50 bool
				_ = _out50
				var _out51 _dafny.Map
				_ = _out51
				_out48, _out49, _out50, _out51 = Companion_Default___.M0(globalState)
				_617_v4 = _out48
				_618_v5 = _out49
				_619_v6 = _out50
				_620_v7 = _out51
				var _621_v8 _dafny.Array
				_ = _621_v8
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_7
				var _nw99 _dafny.Array
				_ = _nw99
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw99 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = (func(_622_v3 *C0) func(_dafny.Int) bool {
						return func(_623_i0 _dafny.Int) bool {
							return (_dafny.IntOfUint32((_dafny.SeqOf((_622_v3).F12())).Cardinality())).Cmp((_622_v3).F13()) == 0
						}
					})(_616_v3)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw99 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw99).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw99).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_621_v8 = _nw99
				_621_v8 = _621_v8
				var _624_v9 _dafny.Array
				_ = _624_v9
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_8
				var _nw100 _dafny.Array
				_ = _nw100
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw100 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = (func(_625_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_626_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_626_i1, _dafny.IntOfUint32((_625_v1).Cardinality()))
						}
					})(_614_v1)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw100 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw100).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw100).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_624_v9 = _nw100
				var _627_v10 _dafny.Array
				_ = _627_v10
				var _nw101 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
				_ = _nw101
				_627_v10 = _nw101
				var _628_v11 _dafny.Map
				_ = _628_v11
				_628_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_627_v10, _624_v9)
				var _629_v12 _dafny.Array
				_ = _629_v12
				_629_v12 = _624_v9
				var _630_v13 _dafny.Sequence
				_ = _630_v13
				_630_v13 = _dafny.SeqOf(_624_v9, _624_v9, _624_v9, _624_v9, _624_v9)
				var _631_v14 _dafny.Array
				_ = _631_v14
				var _nwElement0_18 _dafny.Array = (func() _dafny.Array {
					if _619_v6 {
						return _624_v9
					}
					return _624_v9
				})()
				_ = _nwElement0_18
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(29))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_18, 0)
				(_nw102).ArraySet1(_624_v9, 1)
				(_nw102).ArraySet1(_624_v9, 2)
				(_nw102).ArraySet1(_624_v9, 3)
				(_nw102).ArraySet1(_624_v9, 4)
				(_nw102).ArraySet1(_624_v9, 5)
				(_nw102).ArraySet1(_624_v9, 6)
				(_nw102).ArraySet1(_624_v9, 7)
				(_nw102).ArraySet1(_624_v9, 8)
				(_nw102).ArraySet1((func() _dafny.Array {
					if (_628_v11).Contains(_627_v10) {
						return (_628_v11).Get(_627_v10).(_dafny.Array)
					}
					return _624_v9
				})(), 9)
				(_nw102).ArraySet1(_624_v9, 10)
				(_nw102).ArraySet1((func() _dafny.Array {
					if (_616_v3).F12() {
						return _624_v9
					}
					return _624_v9
				})(), 11)
				(_nw102).ArraySet1(_624_v9, 12)
				(_nw102).ArraySet1(_624_v9, 13)
				(_nw102).ArraySet1(_624_v9, 14)
				(_nw102).ArraySet1(_624_v9, 15)
				(_nw102).ArraySet1(_624_v9, 16)
				(_nw102).ArraySet1(_624_v9, 17)
				(_nw102).ArraySet1((_629_v12), 18)
				(_nw102).ArraySet1(_624_v9, 19)
				(_nw102).ArraySet1(_624_v9, 20)
				(_nw102).ArraySet1(_624_v9, 21)
				(_nw102).ArraySet1(_624_v9, 22)
				(_nw102).ArraySet1(_624_v9, 23)
				(_nw102).ArraySet1(_624_v9, 24)
				(_nw102).ArraySet1(_624_v9, 25)
				(_nw102).ArraySet1(_624_v9, 26)
				(_nw102).ArraySet1(_624_v9, 27)
				(_nw102).ArraySet1((_630_v13).Select((Companion_Default___.SafeIndex(_617_v4, _dafny.IntOfUint32((_630_v13).Cardinality()))).Uint32()).(_dafny.Array), 28)
				_631_v14 = _nw102
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_631_v14), 0))
				_ = _index102
				(_631_v14).ArraySet1(_624_v9, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_631_v14), 0))
				_ = _index103
				(_631_v14).ArraySet1(_624_v9, (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_621_v8), 0))
				_ = _index104
				(_621_v8).ArraySet1(_618_v5, (_index104).Int())
				var _632_v15 _dafny.Array
				_ = _632_v15
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(13))
				_ = _nw103
				_632_v15 = _nw103
				var _633_v16 _dafny.Sequence
				_ = _633_v16
				_633_v16 = _dafny.SeqOf((_616_v3).F12())
				var _634_v17 _dafny.Map
				_ = _634_v17
				_634_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_633_v16, _this.F10())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_632_v15), 0))
				_ = _index105
				(_632_v15).ArraySet1(Companion_Default___.Fm9(_613_v0, _634_v17, _617_v4, _dafny.CodePoint('u'), globalState), (_index105).Int())
				var _635_v18 _dafny.Array
				_ = _635_v18
				var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw104
				_635_v18 = _nw104
				var _636_v19 _dafny.Sequence
				_ = _636_v19
				_636_v19 = _dafny.SeqOf(_635_v18)
				var _637_v20 _dafny.Array
				_ = _637_v20
				var _nwElement0_19 _dafny.Array = (Companion_D3_.Create_DC5_(_635_v18)).Dtor_cf6()
				_ = _nwElement0_19
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(12))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_19, 0)
				(_nw105).ArraySet1(_635_v18, 1)
				(_nw105).ArraySet1(_635_v18, 2)
				(_nw105).ArraySet1((_636_v19).Select((Companion_Default___.SafeIndex((_616_v3).F13(), _dafny.IntOfUint32((_636_v19).Cardinality()))).Uint32()).(_dafny.Array), 3)
				(_nw105).ArraySet1(_635_v18, 4)
				(_nw105).ArraySet1(_635_v18, 5)
				(_nw105).ArraySet1(_635_v18, 6)
				(_nw105).ArraySet1(_635_v18, 7)
				(_nw105).ArraySet1(_635_v18, 8)
				(_nw105).ArraySet1(_635_v18, 9)
				(_nw105).ArraySet1((func() _dafny.Array {
					if true {
						return _635_v18
					}
					return _635_v18
				})(), 10)
				(_nw105).ArraySet1(_635_v18, 11)
				_637_v20 = _nw105
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_637_v20), 0))
				_ = _index106
				(_637_v20).ArraySet1(_635_v18, (_index106).Int())
				var _638_v21 _dafny.CodePoint
				_ = _638_v21
				_638_v21 = _dafny.CodePoint('m')
				var _639_v22 _dafny.Sequence
				_ = _639_v22
				_639_v22 = _dafny.SeqOf(_638_v21, _638_v21)
				var _640_v23 _dafny.Map
				_ = _640_v23
				_640_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(_618_v5, globalState), _dafny.IntOfInt64(-133))
				var _641_v24 _dafny.MultiSet
				_ = _641_v24
				_641_v24 = _dafny.MultiSetOf(_613_v0)
				var _642_v25 D4
				_ = _642_v25
				_642_v25 = Companion_D4_.Create_DC9_(_641_v24)
				var _643_v26 _dafny.Array
				_ = _643_v26
				var _nwElement0_20 _dafny.CodePoint = _638_v21
				_ = _nwElement0_20
				var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(20))
				_ = _nw106
				(_nw106).ArraySet1CodePoint(_nwElement0_20, 0)
				(_nw106).ArraySet1CodePoint(_638_v21, 1)
				(_nw106).ArraySet1CodePoint(_638_v21, 2)
				(_nw106).ArraySet1CodePoint((_639_v22).Select((Companion_Default___.SafeIndex((_616_v3).F13(), _dafny.IntOfUint32((_639_v22).Cardinality()))).Uint32()).(_dafny.CodePoint), 3)
				(_nw106).ArraySet1CodePoint(_638_v21, 4)
				(_nw106).ArraySet1CodePoint((_639_v22).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_640_v23).Contains((func() _dafny.Int {
						if (_620_v7).Contains((_616_v3).F12()) {
							return (_620_v7).Get((_616_v3).F12()).(_dafny.Int)
						}
						return ((_642_v25).Dtor_cf16()).Cardinality()
					})()) {
						return (_640_v23).Get((func() _dafny.Int {
							if (_620_v7).Contains((_616_v3).F12()) {
								return (_620_v7).Get((_616_v3).F12()).(_dafny.Int)
							}
							return ((_642_v25).Dtor_cf16()).Cardinality()
						})()).(_dafny.Int)
					}
					return (_616_v3).F13()
				})(), _dafny.IntOfUint32((_639_v22).Cardinality()))).Uint32()).(_dafny.CodePoint), 5)
				(_nw106).ArraySet1CodePoint(_638_v21, 6)
				(_nw106).ArraySet1CodePoint(_638_v21, 7)
				(_nw106).ArraySet1CodePoint(_638_v21, 8)
				(_nw106).ArraySet1CodePoint(_638_v21, 9)
				(_nw106).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _618_v5 {
						return _638_v21
					}
					return _638_v21
				})(), 10)
				(_nw106).ArraySet1CodePoint(_dafny.CodePoint('p'), 11)
				(_nw106).ArraySet1CodePoint(_dafny.CodePoint('y'), 12)
				(_nw106).ArraySet1CodePoint(_638_v21, 13)
				(_nw106).ArraySet1CodePoint(_638_v21, 14)
				(_nw106).ArraySet1CodePoint(_638_v21, 15)
				(_nw106).ArraySet1CodePoint(_638_v21, 16)
				(_nw106).ArraySet1CodePoint(_638_v21, 17)
				(_nw106).ArraySet1CodePoint(_dafny.CodePoint('c'), 18)
				(_nw106).ArraySet1CodePoint(_638_v21, 19)
				_643_v26 = _nw106
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_643_v26), 0))
				_ = _index107
				(_643_v26).ArraySet1CodePoint(_638_v21, (_index107).Int())
				var _644_v27 _dafny.Sequence
				_ = _644_v27
				_644_v27 = _dafny.SeqOf(_641_v24, _dafny.MultiSetFromSeq(_dafny.SeqOf(_613_v0)))
				var _645_v28 _dafny.Map
				_ = _645_v28
				_645_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_639_v22, _619_v6)
				var _646_v29 _dafny.Set
				_ = _646_v29
				_646_v29 = _dafny.SetOf(_613_v0, Companion_Default___.Fm0(_dafny.IntOfUint32((_614_v1).Cardinality()), _this.F10(), _644_v27, _645_v28, globalState), (_616_v3).F12(), _613_v0)
				var _647_v31 _dafny.Map
				_ = _647_v31
				_647_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_639_v22, (_616_v3).F13())
				var _648_v32 _dafny.Sequence
				_ = _648_v32
				_648_v32 = _dafny.SeqOf(func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter21 := _dafny.Iterate((_647_v31).Keys().Elements()); ; {
						_compr_19, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _649_v30 _dafny.Sequence
						_649_v30 = interface{}(_compr_19).(_dafny.Sequence)
						if (_647_v31).Contains(_649_v30) {
							_coll19.Add(_649_v30, _618_v5)
						}
					}
					return _coll19.ToMap()
				}())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_621_v8), 0))
				_ = _index108
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_632_v15), 0))
				_ = _index109
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_637_v20), 0))
				_ = _index110
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_643_v26), 0))
				_ = _index111
				var _rhs126 bool = (func() bool {
					if _618_v5 {
						return Companion_Default___.Fm0((_dafny.Zero).Minus((_646_v29).Cardinality()), _this.F10(), Companion_Default___.Fm11(globalState), (_648_v32).Select((Companion_Default___.SafeIndex((_616_v3).F13(), _dafny.IntOfUint32((_648_v32).Cardinality()))).Uint32()).(_dafny.Map), globalState)
					}
					return _613_v0
				})()
				_ = _rhs126
				var _rhs127 _dafny.Int = ((_616_v3).F13()).Times(_this.F10())
				_ = _rhs127
				var _rhs128 _dafny.Set = _615_v2
				_ = _rhs128
				var _rhs129 _dafny.Array = _635_v18
				_ = _rhs129
				var _rhs130 _dafny.CodePoint = (_639_v22).Select((Companion_Default___.SafeIndex(_617_v4, _dafny.IntOfUint32((_639_v22).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs130
				var _lhs81 _dafny.Array = _621_v8
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_621_v8), 0))
				_ = _lhs82
				var _lhs83 _dafny.Array = _632_v15
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_632_v15), 0))
				_ = _lhs84
				var _lhs85 _dafny.Array = _637_v20
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_637_v20), 0))
				_ = _lhs86
				var _lhs87 _dafny.Array = _643_v26
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_643_v26), 0))
				_ = _lhs88
				(_lhs81).ArraySet1(_rhs126, (_lhs82).Int())
				_617_v4 = _rhs127
				(_lhs83).ArraySet1(_rhs128, (_lhs84).Int())
				(_lhs85).ArraySet1(_rhs129, (_lhs86).Int())
				(_lhs87).ArraySet1CodePoint(_rhs130, (_lhs88).Int())
			} else {
				_614_v1 = _614_v1
				r0 = ((_616_v3).F13()).Minus(_dafny.IntOfInt64(-909))
				var _650_v33 _dafny.Array
				_ = _650_v33
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw107
				_650_v33 = _nw107
				var _rhs131 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_616_v3).F13(), (_616_v3).F13()), Companion_Default___.SafeModuloInt((_616_v3).F13(), _this.F10()))
				_ = _rhs131
				var _rhs132 _dafny.Array = _650_v33
				_ = _rhs132
				r0 = _rhs131
				_650_v33 = _rhs132
				var _651_v34 _dafny.Sequence
				_ = _651_v34
				_651_v34 = _dafny.UnicodeSeqOfUtf8Bytes("ajenxyshb")
				r0 = Companion_Default___.SafeModuloInt((_614_v1).Select((Companion_Default___.SafeIndex((_616_v3).F13(), _dafny.IntOfUint32((_614_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_651_v34).Cardinality()))
				var _652_v35 _dafny.Sequence
				_ = _652_v35
				_652_v35 = _651_v34
				var _653_v36 _dafny.Sequence
				_ = _653_v36
				_653_v36 = _dafny.SeqOf(Companion_Default___.Fm2(false, _dafny.IntOfInt64(996), (_616_v3).F13(), !(_613_v0), globalState), _dafny.Companion_Sequence_.Concatenate((_651_v34), _651_v34))
				_653_v36 = _653_v36
			}
			var _654_v37 _dafny.Sequence
			_ = _654_v37
			_654_v37 = _dafny.UnicodeSeqOfUtf8Bytes("adcgwephx")
			_654_v37 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lndvj"), _654_v37)
			if (Companion_D3_.Create_DC8_(!((_616_v3).F12()), (_615_v2).Cardinality(), (_616_v3).F13())).Dtor_cf13() {
				_613_v0 = (_616_v3).F12()
				var _655_v38 _dafny.Array
				_ = _655_v38
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_9
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) bool = (func(_656_v0 bool) func(_dafny.Int) bool {
						return func(_657_i2 _dafny.Int) bool {
							return _656_v0
						}
					})(_613_v0)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw108 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw108).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw108).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_655_v38 = _nw108
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_655_v38), 0))
				_ = _index112
				(_655_v38).ArraySet1(_613_v0, (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_655_v38), 0))
				_ = _index113
				(_655_v38).ArraySet1(_613_v0, (_index113).Int())
				var _658_v39 D3
				_ = _658_v39
				_658_v39 = Companion_D3_.Create_DC8_((_655_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_655_v38), 0))).Int()).(bool), _this.F10(), _dafny.IntOfInt64(704))
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_655_v38), 0))
				_ = _index114
				(_655_v38).ArraySet1((_dafny.IntOfUint32((_614_v1).Cardinality())).Cmp(Companion_Default___.SafeModuloInt((_658_v39).Dtor_cf15(), _dafny.IntOfUint32((_654_v37).Cardinality()))) > 0, (_index114).Int())
				_613_v0 = !(_613_v0)
				var _659_v40 _dafny.Sequence
				_ = _659_v40
				_659_v40 = _654_v37
				var _660_v41 _dafny.Map
				_ = _660_v41
				_660_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_659_v40), (_dafny.Zero).Minus((_616_v3).F13()))
				_660_v41 = (_660_v41).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wayojdhbc"), _dafny.UnicodeSeqOfUtf8Bytes("kwne")), (_616_v3).F13())
			} else {
				var _661_v42 _dafny.Array
				_ = _661_v42
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_10
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = (func(_662_v0 bool) func(_dafny.Int) bool {
						return func(_663_i3 _dafny.Int) bool {
							return _662_v0
						}
					})(_613_v0)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw109 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw109).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw109).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_661_v42 = _nw109
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))
				_ = _index115
				(_661_v42).ArraySet1(!(!(!((_616_v3).F12()))), (_index115).Int())
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))
				_ = _index116
				(_661_v42).ArraySet1(false, (_index116).Int())
				_616_v3 = _616_v3
				_613_v0 = (_661_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))).Int()).(bool)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))
				_ = _index117
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))
				_ = _index118
				var _rhs133 bool = (_616_v3).F12()
				_ = _rhs133
				var _rhs134 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(377))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}(func(_664_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				})), _654_v37)
				_ = _rhs134
				var _rhs135 bool = (_616_v3).F12()
				_ = _rhs135
				var _rhs136 bool = (((_616_v3).F12()) && ((_661_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))).Int()).(bool))) || ((_616_v3).F12())
				_ = _rhs136
				var _lhs89 _dafny.Array = _661_v42
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))
				_ = _lhs90
				var _lhs91 _dafny.Array = _661_v42
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))
				_ = _lhs92
				_613_v0 = _rhs133
				_654_v37 = _rhs134
				(_lhs89).ArraySet1(_rhs135, (_lhs90).Int())
				(_lhs91).ArraySet1(_rhs136, (_lhs92).Int())
				var _665_v43 _dafny.Array
				_ = _665_v43
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_11
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Int = func(_666_i5 _dafny.Int) _dafny.Int {
						return (_666_i5).Minus(_this.F10())
					}
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw110 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw110).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw110).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_665_v43 = _nw110
				var _667_v44 _dafny.Map
				_ = _667_v44
				_667_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_661_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))).Int()).(bool), (_661_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_661_v42), 0))).Int()).(bool))
				var _668_v45 D1
				_ = _668_v45
				_668_v45 = Companion_D1_.Create_DC3_(_613_v0, Companion_Default___.Fm1((func() bool {
					if (_667_v44).Contains((_616_v3).F12()) {
						return (_667_v44).Get((_616_v3).F12()).(bool)
					}
					return true
				})(), _613_v0, globalState))
				var _669_v46 _dafny.Map
				_ = _669_v46
				_669_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_665_v43, func(_pat_let4_0 D1) D1 {
					return func(_670_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let5_0 bool) D1 {
							return func(_671_dt__update_hcf3_h0 bool) D1 {
								return Companion_D1_.Create_DC3_(_671_dt__update_hcf3_h0, (_670_dt__update__tmp_h1).Dtor_cf4())
							}(_pat_let5_0)
						}(!(false))
					}(_pat_let4_0)
				}(_668_v45))
				(globalState).F2 = (_669_v46).Cardinality()
			}
			var _672_v47 D1
			_ = _672_v47
			_672_v47 = Companion_D1_.Create_DC2_(_613_v0)
			var _source14 D1 = _672_v47
			_ = _source14
			if _source14.Is_DC2() {
				var _673___mcc_h0 bool = _source14.Get_().(D1_DC2).Cf2
				_ = _673___mcc_h0
				var _674_cf2 bool = _673___mcc_h0
				_ = _674_cf2
				(_this).F10_set_((_dafny.Zero).Minus((func() _dafny.Int {
					if (_615_v2).IsProperSubsetOf(_615_v2) {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_614_v1, _dafny.SeqOf((_614_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-325), _dafny.IntOfUint32((_614_v1).Cardinality()))).Uint32()).(_dafny.Int)))).Cardinality())
					}
					return Companion_Default___.Fm10((_616_v3).F12(), globalState)
				})()))
				var _675_v48 _dafny.Map
				_ = _675_v48
				_675_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_674_cf2, Companion_Default___.Fm10((_616_v3).F12(), globalState))
				var _676_v49 _dafny.Map
				_ = _676_v49
				_676_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v37, _674_cf2)
				var _677_v50 _dafny.Map
				_ = _677_v50
				_677_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_616_v3).F12(), Companion_Default___.Fm0((_675_v48).Cardinality(), (_616_v3).F13(), Companion_Default___.Fm11(globalState), _676_v49, globalState))
				(_this).F10_set_((_dafny.Zero).Minus((_677_v50).Cardinality()))
				_675_v48 = (_675_v48).Update(true, (_616_v3).F13())
				var _678_v51 _dafny.Map
				_ = _678_v51
				_678_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_613_v0), (_616_v3).F13()), (_616_v3).F12())
				_678_v51 = (_678_v51).Update(_675_v48, !(_613_v0) || ((_616_v3).F12()))
			} else if _source14.Is_DC3() {
				var _679___mcc_h1 bool = _source14.Get_().(D1_DC3).Cf3
				_ = _679___mcc_h1
				var _680___mcc_h2 _dafny.Int = _source14.Get_().(D1_DC3).Cf4
				_ = _680___mcc_h2
				var _681_cf4 _dafny.Int = _680___mcc_h2
				_ = _681_cf4
				var _682_cf3 bool = _679___mcc_h1
				_ = _682_cf3
				var _683_v52 _dafny.Sequence
				_ = _683_v52
				_683_v52 = _dafny.SeqOf(_682_cf3, _613_v0)
				var _684_v53 _dafny.MultiSet
				_ = _684_v53
				_684_v53 = _dafny.MultiSetOf((_616_v3).F13(), _this.F10())
				var _685_v54 D3
				_ = _685_v54
				_685_v54 = Companion_D3_.Create_DC8_(_dafny.Companion_Sequence_.IsPrefixOf(_683_v52, _683_v52), (_616_v3).F13(), (func() _dafny.Int {
					if (_684_v53).Contains(_this.F10()) {
						return (_684_v53).Multiplicity(_this.F10())
					}
					return _this.F10()
				})())
				_685_v54 = _685_v54
				(globalState).F2 = _this.F10()
				var _686_v55 _dafny.MultiSet
				_ = _686_v55
				_686_v55 = _dafny.MultiSetOf(_613_v0, _682_cf3)
				var _687_v56 _dafny.Map
				_ = _687_v56
				_687_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_686_v55, (_616_v3).F12())
				var _688_v57 D3
				_ = _688_v57
				_688_v57 = Companion_D3_.Create_DC6_((_616_v3).F13(), _613_v0, (_616_v3).F12(), _this.F10(), (_616_v3).F12())
				_687_v56 = (_687_v56).Update((_686_v55).Update((_616_v3).F12(), Companion_Default___.Abs((_615_v2).Cardinality())), (func() bool {
					if _613_v0 {
						return (_688_v57).Dtor_cf8()
					}
					return (_616_v3).F12()
				})())
				_613_v0 = (_616_v3).F12()
			} else {
				var _689___mcc_h3 bool = _source14.Get_().(D1_DC1).Cf1
				_ = _689___mcc_h3
				var _690_cf1 bool = _689___mcc_h3
				_ = _690_cf1
				var _691_v58 _dafny.Map
				_ = _691_v58
				_691_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_616_v3).F13(), !((_616_v3).F12()))
				_691_v58 = (_691_v58).Update((_616_v3).F13(), (_616_v3).F12())
				var _692_v59 D1
				_ = _692_v59
				_692_v59 = Companion_D1_.Create_DC1_((_616_v3).F12())
				var _693_v60 _dafny.Sequence
				_ = _693_v60
				_693_v60 = _dafny.SeqOf(_613_v0, (_616_v3).F12())
				var _694_v61 _dafny.Map
				_ = _694_v61
				_694_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_693_v60).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F10()), _dafny.IntOfUint32((_693_v60).Cardinality()))).Uint32()).(bool), !((_616_v3).F12()))
				_690_cf1 = ((_this).Fm7(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_616_v3).F12(), _692_v59), _690_cf1, (_616_v3).F13(), globalState)) || (((_694_v61).Cardinality()).Cmp(_this.F10()) != 0)
				(globalState).F2 = _dafny.IntOfUint32((_654_v37).Cardinality())
				var _695_v62 _dafny.Array
				_ = _695_v62
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_12
				var _nw111 _dafny.Array
				_ = _nw111
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw111 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Int = func(_696_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_696_i6, _dafny.IntOfInt64(850))
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw111 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw111).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw111).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_695_v62 = _nw111
				var _697_v63 _dafny.Map
				_ = _697_v63
				_697_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10(), _695_v62)
				var _698_v64 _dafny.Array
				_ = _698_v64
				var _nwElement0_21 _dafny.Array = _695_v62
				_ = _nwElement0_21
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(11))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_21, 0)
				(_nw112).ArraySet1(_695_v62, 1)
				(_nw112).ArraySet1(_695_v62, 2)
				(_nw112).ArraySet1(_695_v62, 3)
				(_nw112).ArraySet1(_695_v62, 4)
				(_nw112).ArraySet1(_695_v62, 5)
				(_nw112).ArraySet1((func() _dafny.Array {
					if (_697_v63).Contains(_this.F10()) {
						return (_697_v63).Get(_this.F10()).(_dafny.Array)
					}
					return _695_v62
				})(), 6)
				(_nw112).ArraySet1(_695_v62, 7)
				(_nw112).ArraySet1(_695_v62, 8)
				(_nw112).ArraySet1(_695_v62, 9)
				(_nw112).ArraySet1(_695_v62, 10)
				_698_v64 = _nw112
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_698_v64), 0))
				_ = _index119
				(_698_v64).ArraySet1(_695_v62, (_index119).Int())
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_698_v64), 0))
				_ = _index120
				(_698_v64).ArraySet1(_695_v62, (_index120).Int())
			}
		} else {
			var _699_v65 _dafny.Array
			_ = _699_v65
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_13
			var _nw113 _dafny.Array
			_ = _nw113
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw113 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Int = (func(_700_v0 bool) func(_dafny.Int) _dafny.Int {
					return func(_701_i7 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_701_i7, _dafny.IntOfUint32((_dafny.SeqOf(false, _700_v0)).Cardinality()))
					}
				})(_613_v0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw113 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw113).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw113).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_699_v65 = _nw113
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))
			_ = _index121
			(_699_v65).ArraySet1(_this.F10(), (_index121).Int())
			var _702_v66 _dafny.Set
			_ = _702_v66
			_702_v66 = _dafny.SetOf(_613_v0, _613_v0, _613_v0)
			var _703_v67 _dafny.Map
			_ = _703_v67
			_703_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_613_v0, _this.F10())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))
			_ = _index122
			var _rhs137 bool = _613_v0
			_ = _rhs137
			var _rhs138 bool = (_702_v66).IsProperSubsetOf(_702_v66)
			_ = _rhs138
			var _rhs139 _dafny.Int = (_this.F10()).Plus(_dafny.IntOfInt64(-305))
			_ = _rhs139
			var _rhs140 _dafny.Int = (_dafny.Zero).Minus((_703_v67).Cardinality())
			_ = _rhs140
			var _lhs93 _dafny.Array = _699_v65
			_ = _lhs93
			var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))
			_ = _lhs94
			var _lhs95 *GlobalState = globalState
			_ = _lhs95
			_613_v0 = _rhs137
			_613_v0 = _rhs138
			(_lhs93).ArraySet1(_rhs139, (_lhs94).Int())
			_lhs95.F2 = _rhs140
			var _704_v68 _dafny.Array
			_ = _704_v68
			var _nw114 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw114
			_704_v68 = _nw114
			var _705_v69 _dafny.Sequence
			_ = _705_v69
			_705_v69 = _dafny.UnicodeSeqOfUtf8Bytes("cirg")
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))
			_ = _index123
			(_704_v68).ArraySet1(_dafny.Companion_Sequence_.Equal(_705_v69, _705_v69), (_index123).Int())
			var _706_v70 _dafny.Array
			_ = _706_v70
			var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw115
			_706_v70 = _nw115
			var _707_v71 _dafny.Array
			_ = _707_v71
			var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw116
			_707_v71 = _nw116
			var _708_v72 T0
			_ = _708_v72
			var _nw117 *C1 = New_C1_()
			_ = _nw117
			_nw117.Ctor__(_613_v0, Companion_D3_.Create_DC5_(_707_v71))
			_708_v72 = _nw117
			var _709_v73 _dafny.MultiSet
			_ = _709_v73
			_709_v73 = _dafny.MultiSetOf(_708_v72)
			var _710_v74 _dafny.MultiSet
			_ = _710_v74
			_710_v74 = _dafny.MultiSetOf(_613_v0)
			var _711_v75 _dafny.Map
			_ = _711_v75
			_711_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_710_v74, _613_v0)
			var _712_v76 _dafny.Sequence
			_ = _712_v76
			_712_v76 = _dafny.SeqOf(_613_v0, _613_v0, _613_v0, _613_v0, _613_v0)
			var _713_v77 _dafny.Map
			_ = _713_v77
			_713_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_711_v75).Cardinality(), _dafny.MultiSetOf((_699_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_712_v76).Cardinality())))
			var _714_v78 _dafny.Map
			_ = _714_v78
			_714_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_713_v77, _706_v70)
			var _715_v81 _dafny.Map
			_ = _715_v81
			_715_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10(), _dafny.IntOfUint32((_705_v69).Cardinality()))
			var _716_v82 _dafny.CodePoint
			_ = _716_v82
			_716_v82 = _dafny.CodePoint('a')
			var _717_v83 _dafny.Set
			_ = _717_v83
			_717_v83 = _dafny.SetOf(_716_v82)
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))
			_ = _index124
			var _rhs141 bool = (_615_v2).IsDisjointFrom(_615_v2)
			_ = _rhs141
			var _rhs142 _dafny.Int = (func() _dafny.Int {
				if (_709_v73).Contains(_708_v72) {
					return (_709_v73).Multiplicity(_708_v72)
				}
				return _dafny.IntOfInt64(265)
			})()
			_ = _rhs142
			var _rhs143 _dafny.Array = (func() _dafny.Array {
				if (_714_v78).Contains((func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-398), _dafny.IntOfInt64(496))); ; {
						_compr_20, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _718_v79 _dafny.Int
						_718_v79 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(-398)).Cmp(_718_v79) <= 0) && ((_718_v79).Cmp(_dafny.IntOfInt64(496)) < 0) {
							_coll20.Add((_718_v79).Minus(_this.F10()), _dafny.MultiSetOf((_702_v66).Cardinality(), (_699_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))).Int()).(_dafny.Int)))
						}
					}
					return _coll20.ToMap()
				}()).Merge(func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter23 := _dafny.Iterate((_715_v81).Keys().Elements()); ; {
						_compr_21, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _719_v80 _dafny.Int
						_719_v80 = interface{}(_compr_21).(_dafny.Int)
						if (_715_v81).Contains(_719_v80) {
							_coll21.Add(Companion_Default___.SafeModuloInt(_719_v80, _dafny.IntOfInt64(-284)), _dafny.MultiSetOf(_this.F10()))
						}
					}
					return _coll21.ToMap()
				}())) {
					return (_714_v78).Get((func() _dafny.Map {
						var _coll22 = _dafny.NewMapBuilder()
						_ = _coll22
						for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-398), _dafny.IntOfInt64(496))); ; {
							_compr_22, _ok24 := _iter24()
							if !_ok24 {
								break
							}
							var _720_v79 _dafny.Int
							_720_v79 = interface{}(_compr_22).(_dafny.Int)
							if ((_dafny.IntOfInt64(-398)).Cmp(_720_v79) <= 0) && ((_720_v79).Cmp(_dafny.IntOfInt64(496)) < 0) {
								_coll22.Add((_720_v79).Minus(_this.F10()), _dafny.MultiSetOf((_702_v66).Cardinality(), (_699_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))).Int()).(_dafny.Int)))
							}
						}
						return _coll22.ToMap()
					}()).Merge(func() _dafny.Map {
						var _coll23 = _dafny.NewMapBuilder()
						_ = _coll23
						for _iter25 := _dafny.Iterate((_715_v81).Keys().Elements()); ; {
							_compr_23, _ok25 := _iter25()
							if !_ok25 {
								break
							}
							var _721_v80 _dafny.Int
							_721_v80 = interface{}(_compr_23).(_dafny.Int)
							if (_715_v81).Contains(_721_v80) {
								_coll23.Add(Companion_Default___.SafeModuloInt(_721_v80, _dafny.IntOfInt64(-284)), _dafny.MultiSetOf(_this.F10()))
							}
						}
						return _coll23.ToMap()
					}())).(_dafny.Array)
				}
				return _706_v70
			})()
			_ = _rhs143
			var _rhs144 bool = ((func() _dafny.Set {
				if _613_v0 {
					return _717_v83
				}
				return _dafny.SetOf(_716_v82, _dafny.CodePoint('i'))
			})()).IsProperSubsetOf(_717_v83)
			_ = _rhs144
			var _rhs145 bool = (Companion_Default___.Fm23(_715_v81, _613_v0, globalState)).IsDisjointFrom(_702_v66)
			_ = _rhs145
			var _lhs96 _dafny.Array = _704_v68
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))
			_ = _lhs97
			(_lhs96).ArraySet1(_rhs141, (_lhs97).Int())
			r0 = _rhs142
			_706_v70 = _rhs143
			_613_v0 = _rhs144
			_613_v0 = _rhs145
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))
			_ = _index125
			(_704_v68).ArraySet1(((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ethwsioyp")).Cardinality())).Minus((_699_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_699_v65), 0))).Int()).(_dafny.Int))).Cmp(((Companion_Default___.Fm24(_613_v0, _712_v76, (_704_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))).Int()).(bool), globalState)).Union(_dafny.MultiSetFromSeq(_614_v1))).Cardinality()) > 0, (_index125).Int())
			var _722_v84 _dafny.Map
			_ = _722_v84
			_722_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))
			_ = _index126
			(_704_v68).ArraySet1((_this.F10()).Cmp((_722_v84).Cardinality()) > 0, (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))
			_ = _index127
			(_704_v68).ArraySet1((_704_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_704_v68), 0))).Int()).(bool), (_index127).Int())
		}
		(globalState).F2 = Companion_Default___.Fm10(_613_v0, globalState)
		var _723_v85 _dafny.Sequence
		_ = _723_v85
		_723_v85 = _dafny.SeqOf(_613_v0, true)
		var _724_v86 _dafny.Map
		_ = _724_v86
		_724_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_v85, (_dafny.Zero).Minus(_this.F10()))
		var _725_v88 _dafny.Sequence
		_ = _725_v88
		_725_v88 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_723_v85))
		var _726_v89 _dafny.Sequence
		_ = _726_v89
		_726_v89 = _dafny.UnicodeSeqOfUtf8Bytes("yjyrhbv")
		var _727_v90 _dafny.Map
		_ = _727_v90
		_727_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_v89, true)
		var _728_v91 *C2
		_ = _728_v91
		var _nw118 *C2 = New_C2_()
		_ = _nw118
		_nw118.Ctor__(_613_v0, Companion_Default___.Fm0((func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter26 := _dafny.Iterate((_615_v2).Elements()); ; {
				_compr_24, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _729_v87 _dafny.Int
				_729_v87 = interface{}(_compr_24).(_dafny.Int)
				if (_615_v2).Contains(_729_v87) {
					_coll24.Add(Companion_Default___.SafeDivisionInt(_729_v87, _this.F10()), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('g'))).Cardinality()))).Cardinality())
				}
			}
			return _coll24.ToMap()
		}()).Cardinality(), _this.F10(), _725_v88, _727_v90, globalState))
		_728_v91 = _nw118
		var _730_v92 _dafny.Sequence
		_ = _730_v92
		_730_v92 = _dafny.SeqOf(_728_v91, _728_v91)
		_724_v86 = (_724_v86).Update(_723_v85, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_730_v92, _730_v92)).Cardinality()))
		var _731_v93 _dafny.MultiSet
		_ = _731_v93
		_731_v93 = _dafny.MultiSetOf(_dafny.IntOfUint32((_723_v85).Cardinality()))
		var _732_v94 _dafny.MultiSet
		_ = _732_v94
		_732_v94 = _dafny.MultiSetOf(_this.F10(), _this.F10(), _dafny.IntOfUint32((_726_v89).Cardinality()), (_731_v93).Cardinality())
		var _733_v95 _dafny.MultiSet
		_ = _733_v95
		_733_v95 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_734_i9 _dafny.Int) _dafny.Int {
			return (_dafny.SetOf(_dafny.CodePoint('i'))).Cardinality()
		}))))
		var _735_v96 _dafny.Array
		_ = _735_v96
		var _nwElement0_22 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(339))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_736_i8 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(22)
		})), _614_v1))
		_ = _nwElement0_22
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(15))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_22, 0)
		(_nw119).ArraySet1(_732_v94, 1)
		(_nw119).ArraySet1(_732_v94, 2)
		(_nw119).ArraySet1(_dafny.MultiSetOf(_this.F10(), ((_this).F11()).Cardinality(), (_dafny.MultiSetOf((_728_v91).F15())).Cardinality(), (_615_v2).Cardinality()), 3)
		(_nw119).ArraySet1((func() _dafny.MultiSet {
			if false {
				return _732_v94
			}
			return _732_v94
		})(), 4)
		(_nw119).ArraySet1((_731_v93).Update(_this.F10(), Companion_Default___.Abs(_this.F10())), 5)
		(_nw119).ArraySet1((_732_v94).Intersection(_731_v93), 6)
		(_nw119).ArraySet1(Companion_Default___.Fm24((_728_v91).F15(), _723_v85, _728_v91.F14, globalState), 7)
		(_nw119).ArraySet1((_732_v94).Intersection((_dafny.MultiSetOf(_this.F10())).Update((_733_v95).Cardinality(), Companion_Default___.Abs(_this.F10()))), 8)
		(_nw119).ArraySet1((_732_v94).Intersection(_731_v93), 9)
		(_nw119).ArraySet1((Companion_Default___.Fm24((_728_v91).F15(), Companion_Default___.Fm18(_this.F10(), (_728_v91).F15(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_728_v91).F15()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.IntOfUint32((_dafny.SeqOf((_728_v91).F15())).Cardinality()))).Uint32(), true)).Cardinality()), globalState), (_728_v91).F15(), globalState)).Intersection(_732_v94), 10)
		(_nw119).ArraySet1(_dafny.MultiSetOf(_this.F10(), (_dafny.Zero).Minus(_this.F10())), 11)
		(_nw119).ArraySet1(_731_v93, 12)
		(_nw119).ArraySet1(_731_v93, 13)
		(_nw119).ArraySet1(_731_v93, 14)
		_735_v96 = _nw119
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_735_v96), 0))
		_ = _index128
		(_735_v96).ArraySet1(_dafny.MultiSetFromSeq(_614_v1), (_index128).Int())
		var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_735_v96), 0))
		_ = _index129
		(_735_v96).ArraySet1(((_dafny.MultiSetOf(_this.F10(), _dafny.IntOfInt64(552))).Intersection(_732_v94)).Intersection(_731_v93), (_index129).Int())
		var _737_v97 _dafny.Sequence
		_ = _737_v97
		_737_v97 = _726_v89
		var _738_v98 _dafny.Array
		_ = _738_v98
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_14
		var _nw120 _dafny.Array
		_ = _nw120
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw120 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Int = func(_739_i10 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_739_i10, _this.F10())
			}
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw120 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw120).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw120).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_738_v98 = _nw120
		var _740_v99 _dafny.Map
		_ = _740_v99
		_740_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_737_v97, _738_v98)
		var _741_v100 D8
		_ = _741_v100
		_741_v100 = Companion_D8_.Create_DC17_(_740_v99)
		(globalState).F2 = (((_740_v99).Merge((_741_v100).Dtor_cf27())).Merge(_740_v99)).Cardinality()
		r0 = _this.F10()
		return r0
	}
}
func (_this *C3) M2(globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _742_v0 D6
		_ = _742_v0
		_742_v0 = Companion_D6_.Create_DC13_((_this.F10()).Cmp((_dafny.Zero).Minus(_this.F10())) > 0)
		var _source15 D6 = _742_v0
		_ = _source15
		if _source15.Is_DC13() {
			var _743___mcc_h0 bool = _source15.Get_().(D6_DC13).Cf22
			_ = _743___mcc_h0
			var _744_cf22 bool = _743___mcc_h0
			_ = _744_cf22
			r0 = _744_cf22
			var _745_v1 _dafny.Array
			_ = _745_v1
			var _nw121 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw121
			_745_v1 = _nw121
			var _746_v2 D7
			_ = _746_v2
			_746_v2 = Companion_D7_.Create_DC14_(_dafny.SetOf(_745_v1, _745_v1, _745_v1, _745_v1))
			var _747_v3 _dafny.Set
			_ = _747_v3
			_747_v3 = _dafny.SetOf(_745_v1)
			var _pat_let_tv2 = _747_v3
			_ = _pat_let_tv2
			var _source16 D7 = func(_pat_let6_0 D7) D7 {
				return func(_748_dt__update__tmp_h0 D7) D7 {
					return func(_pat_let7_0 _dafny.Set) D7 {
						return func(_749_dt__update_hcf23_h0 _dafny.Set) D7 {
							return Companion_D7_.Create_DC14_(_749_dt__update_hcf23_h0)
						}(_pat_let7_0)
					}(_pat_let_tv2)
				}(_pat_let6_0)
			}(_746_v2)
			_ = _source16
			if _source16.Is_DC15() {
				var _750___mcc_h2 _dafny.Int = _source16.Get_().(D7_DC15).Cf24
				_ = _750___mcc_h2
				var _751___mcc_h3 bool = _source16.Get_().(D7_DC15).Cf25
				_ = _751___mcc_h3
				var _752_cf25 bool = _751___mcc_h3
				_ = _752_cf25
				var _753_cf24 _dafny.Int = _750___mcc_h2
				_ = _753_cf24
				(globalState).F2 = _753_cf24
				(_this).F10_set_((Companion_Default___.SafeModuloInt(_this.F10(), _dafny.IntOfInt64(317))).Minus(((_dafny.Zero).Minus((_dafny.Zero).Minus(_753_cf24))).Times(_this.F10())))
				var _754_v4 _dafny.Map
				_ = _754_v4
				_754_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_752_cf25, _753_cf24)
				var _755_v5 _dafny.MultiSet
				_ = _755_v5
				_755_v5 = _dafny.MultiSetOf(_753_cf24, _this.F10())
				var _756_v6 _dafny.MultiSet
				_ = _756_v6
				_756_v6 = _dafny.MultiSetOf(_752_cf25)
				var _757_v7 D9
				_ = _757_v7
				_757_v7 = Companion_D9_.Create_DC19_(_754_v4)
				var _758_v8 _dafny.MultiSet
				_ = _758_v8
				_758_v8 = _dafny.MultiSetOf(_754_v4, (_754_v4).Merge(_754_v4), (_754_v4).Update(_752_cf25, _this.F10()), _754_v4, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_744_cf22, (func() _dafny.Int {
					if (_755_v5).Contains((func() _dafny.Int {
						if (_756_v6).Contains(_744_cf22) {
							return (_756_v6).Multiplicity(_744_cf22)
						}
						return _753_cf24
					})()) {
						return (_755_v5).Multiplicity((func() _dafny.Int {
							if (_756_v6).Contains(_744_cf22) {
								return (_756_v6).Multiplicity(_744_cf22)
							}
							return _753_cf24
						})())
					}
					return _753_cf24
				})())).Merge((_757_v7).Dtor_cf28()))
				var _759_v9 _dafny.Sequence
				_ = _759_v9
				_759_v9 = _dafny.UnicodeSeqOfUtf8Bytes("wiwrmiodu")
				var _rhs146 _dafny.MultiSet = _758_v8
				_ = _rhs146
				var _rhs147 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("qjdplvq")
				_ = _rhs147
				_758_v8 = _rhs146
				_759_v9 = _rhs147
				var _760_v10 _dafny.Int
				_ = _760_v10
				var _761_v11 bool
				_ = _761_v11
				var _762_v12 bool
				_ = _762_v12
				var _763_v13 _dafny.Map
				_ = _763_v13
				var _out52 _dafny.Int
				_ = _out52
				var _out53 bool
				_ = _out53
				var _out54 bool
				_ = _out54
				var _out55 _dafny.Map
				_ = _out55
				_out52, _out53, _out54, _out55 = Companion_Default___.M0(globalState)
				_760_v10 = _out52
				_761_v11 = _out53
				_762_v12 = _out54
				_763_v13 = _out55
			} else if _source16.Is_DC14() {
				var _764___mcc_h4 _dafny.Set = _source16.Get_().(D7_DC14).Cf23
				_ = _764___mcc_h4
				var _765_cf23 _dafny.Set = _764___mcc_h4
				_ = _765_cf23
				var _766_v14 _dafny.CodePoint
				_ = _766_v14
				_766_v14 = _dafny.CodePoint('j')
				var _767_v15 D9
				_ = _767_v15
				_767_v15 = Companion_D9_.Create_DC20_(_744_cf22, _766_v14)
				var _768_v16 _dafny.Array
				_ = _768_v16
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw122
				_768_v16 = _nw122
				var _769_v17 D3
				_ = _769_v17
				_769_v17 = Companion_D3_.Create_DC5_(_768_v16)
				var _770_v18 *C1
				_ = _770_v18
				var _nw123 *C1 = New_C1_()
				_ = _nw123
				_nw123.Ctor__(((_767_v15).Dtor_cf29()) || (_744_cf22), _769_v17)
				_770_v18 = _nw123
				var _771_v19 D10
				_ = _771_v19
				_771_v19 = Companion_D10_.Create_DC22_(_770_v18)
				var _pat_let_tv3 = _770_v18
				_ = _pat_let_tv3
				_770_v18 = (func(_pat_let8_0 D10) D10 {
					return func(_772_dt__update__tmp_h1 D10) D10 {
						return func(_pat_let9_0 *C1) D10 {
							return func(_773_dt__update_hcf32_h0 *C1) D10 {
								return Companion_D10_.Create_DC22_(_773_dt__update_hcf32_h0)
							}(_pat_let9_0)
						}(_pat_let_tv3)
					}(_pat_let8_0)
				}(_771_v19)).Dtor_cf32()
				var _774_v20 _dafny.Map
				_ = _774_v20
				_774_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10(), (_this).F11())
				var _775_v21 _dafny.Sequence
				_ = _775_v21
				_775_v21 = _dafny.SeqOf(_this.F10(), _this.F10())
				var _776_v22 _dafny.Map
				_ = _776_v22
				_776_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v20, (_775_v21).Select((Companion_Default___.SafeIndex(_this.F10(), _dafny.IntOfUint32((_775_v21).Cardinality()))).Uint32()).(_dafny.Int))
				_776_v22 = (_776_v22).Update(func() _dafny.Map {
					var _coll25 = _dafny.NewMapBuilder()
					_ = _coll25
					for _iter27 := _dafny.Iterate(((_this).F11()).Keys().Elements()); ; {
						_compr_25, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _777_v23 _dafny.Int
						_777_v23 = interface{}(_compr_25).(_dafny.Int)
						if ((_this).F11()).Contains(_777_v23) {
							_coll25.Add((_777_v23).Minus(_this.F10()), (_this).F11())
						}
					}
					return _coll25.ToMap()
				}(), _this.F10())
				var _778_v24 _dafny.Sequence
				_ = _778_v24
				_778_v24 = _dafny.UnicodeSeqOfUtf8Bytes("d")
				var _779_v25 _dafny.Array
				_ = _779_v25
				var _nwElement0_23 _dafny.Sequence = _775_v21
				_ = _nwElement0_23
				var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(18))
				_ = _nw124
				(_nw124).ArraySet1(_nwElement0_23, 0)
				(_nw124).ArraySet1(Companion_Default___.Fm25(_this.F10(), _this.F10(), _dafny.IntOfUint32((_778_v24).Cardinality()), _this.F10(), globalState), 1)
				(_nw124).ArraySet1(_775_v21, 2)
				(_nw124).ArraySet1(_775_v21, 3)
				(_nw124).ArraySet1(_775_v21, 4)
				(_nw124).ArraySet1(_775_v21, 5)
				(_nw124).ArraySet1(_dafny.SeqOf(_this.F10()), 6)
				(_nw124).ArraySet1(_dafny.SeqOf(_this.F10()), 7)
				(_nw124).ArraySet1(_775_v21, 8)
				(_nw124).ArraySet1(_775_v21, 9)
				(_nw124).ArraySet1(_775_v21, 10)
				(_nw124).ArraySet1(_775_v21, 11)
				(_nw124).ArraySet1(_dafny.SeqOf(_this.F10(), _dafny.IntOfInt64(758)), 12)
				(_nw124).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(903))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}(func(_780_i0 _dafny.Int) _dafny.Int {
					return _this.F10()
				})), 13)
				(_nw124).ArraySet1(_775_v21, 14)
				(_nw124).ArraySet1(_dafny.SeqOf(_this.F10()), 15)
				(_nw124).ArraySet1(_dafny.SeqOf(_this.F10()), 16)
				(_nw124).ArraySet1(_775_v21, 17)
				_779_v25 = _nw124
				var _781_v26 *C2
				_ = _781_v26
				var _nw125 *C2 = New_C2_()
				_ = _nw125
				_nw125.Ctor__((_770_v18).F16(), (func() bool {
					if ((_this).F11()).Contains((_dafny.Zero).Minus(_this.F10())) {
						return ((_this).F11()).Get((_dafny.Zero).Minus(_this.F10())).(bool)
					}
					return (_770_v18).F16()
				})())
				_781_v26 = _nw125
				var _782_v27 _dafny.Map
				_ = _782_v27
				_782_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_779_v25, _781_v26)
				_782_v27 = (_782_v27).Update(_779_v25, _781_v26)
				var _783_v28 _dafny.Array
				_ = _783_v28
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_15
				var _nw126 _dafny.Array
				_ = _nw126
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw126 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Int = (func(_784_v21 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_785_i1 _dafny.Int) _dafny.Int {
							return (_785_i1).Minus((_784_v21).Select((Companion_Default___.SafeIndex(_this.F10(), _dafny.IntOfUint32((_784_v21).Cardinality()))).Uint32()).(_dafny.Int))
						}
					})(_775_v21)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw126 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw126).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw126).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_783_v28 = _nw126
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_783_v28), 0))
				_ = _index130
				(_783_v28).ArraySet1((_this.F10()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dq")).Cardinality())), (_index130).Int())
				var _786_v29 _dafny.Sequence
				_ = _786_v29
				_786_v29 = _dafny.SeqOf(_744_cf22)
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_783_v28), 0))
				_ = _index131
				(_783_v28).ArraySet1((func() _dafny.Int {
					if ((_775_v21).Select((Companion_Default___.SafeIndex(_this.F10(), _dafny.IntOfUint32((_775_v21).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_dafny.Zero).Minus(_this.F10())) < 0 {
						return (Companion_Default___.Fm10((_781_v26).F15(), globalState)).Plus(_dafny.IntOfUint32((_786_v29).Cardinality()))
					}
					return _dafny.IntOfUint32((_778_v24).Cardinality())
				})(), (_index131).Int())
			} else {
				var _787___mcc_h5 D7 = _source16.Get_().(D7_DC16).Cf26
				_ = _787___mcc_h5
				var _788_cf26 D7 = _787___mcc_h5
				_ = _788_cf26
				r2 = !((_744_cf22) && (_744_cf22))
				(globalState).F2 = (_dafny.IntOfInt64(187)).Times(_this.F10())
				var _789_v30 *C2
				_ = _789_v30
				var _nw127 *C2 = New_C2_()
				_ = _nw127
				_nw127.Ctor__(_744_cf22, false)
				_789_v30 = _nw127
				var _790_v31 _dafny.Sequence
				_ = _790_v31
				_790_v31 = _dafny.SeqOf((_789_v30).F15())
				var _791_v32 bool
				_ = _791_v32
				var _792_v33 _dafny.Array
				_ = _792_v33
				var _out56 bool
				_ = _out56
				var _out57 _dafny.Array
				_ = _out57
				_out56, _out57 = (_789_v30).M3(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_790_v31, (Companion_Default___.SafeIndex(_this.F10(), _dafny.IntOfUint32((_790_v31).Cardinality()))).Uint32(), _744_cf22), _dafny.SeqOf(!((_789_v30).F15()))), globalState)
				_791_v32 = _out56
				_792_v33 = _out57
			}
			var _793_v34 _dafny.Sequence
			_ = _793_v34
			_793_v34 = _dafny.SeqOf(_744_cf22)
			var _794_v35 D9
			_ = _794_v35
			_794_v35 = Companion_D9_.Create_DC20_(_dafny.Companion_Sequence_.Equal(_793_v34, _793_v34), Companion_Default___.Fm26((func() bool {
				if ((_this).F11()).Contains(_this.F10()) {
					return ((_this).F11()).Get(_this.F10()).(bool)
				}
				return _744_cf22
			})(), globalState))
			var _795_v36 _dafny.MultiSet
			_ = _795_v36
			_795_v36 = _dafny.MultiSetOf((Companion_Default___.Fm27(globalState)).Cardinality())
			var _796_v37 _dafny.CodePoint
			_ = _796_v37
			_796_v37 = _dafny.CodePoint('f')
			_794_v35 = Companion_D9_.Create_DC20_((_dafny.MultiSetOf(_this.F10())).IsSubsetOf(_795_v36), _796_v37)
			var _797_v38 _dafny.Sequence
			_ = _797_v38
			_797_v38 = _dafny.UnicodeSeqOfUtf8Bytes("bicxihw")
			var _798_v39 _dafny.Map
			_ = _798_v39
			_798_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_744_cf22, _this.F10())
			var _799_v40 _dafny.Map
			_ = _799_v40
			_799_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-765))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_800_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_801_i2 _dafny.Int) _dafny.CodePoint {
					return _800_v37
				}
			})(_796_v37)))).Cardinality()))
			r1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_797_v38, _dafny.Companion_Sequence_.Update(_797_v38, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.IntOfUint32((_797_v38).Cardinality()))).Uint32(), _796_v37))).Cardinality())).Times(((func() _dafny.Int {
				if (_798_v39).Contains(_744_cf22) {
					return (_798_v39).Get(_744_cf22).(_dafny.Int)
				}
				return _this.F10()
			})()).Times((func() _dafny.Int {
				if (_799_v40).Contains((_795_v36).Cardinality()) {
					return (_799_v40).Get((_795_v36).Cardinality()).(_dafny.Int)
				}
				return _this.F10()
			})()))
		} else {
			var _802___mcc_h1 _dafny.Map = _source15.Get_().(D6_DC12).Cf21
			_ = _802___mcc_h1
			var _803_cf21 _dafny.Map = _802___mcc_h1
			_ = _803_cf21
			(globalState).F2 = _dafny.IntOfInt64(447)
			(globalState).F2 = _this.F10()
			var _804_v41 bool
			_ = _804_v41
			_804_v41 = false
			var _805_v42 _dafny.Array
			_ = _805_v42
			var _nwElement0_24 bool = _804_v41
			_ = _nwElement0_24
			var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(3))
			_ = _nw128
			(_nw128).ArraySet1(_nwElement0_24, 0)
			(_nw128).ArraySet1(_804_v41, 1)
			(_nw128).ArraySet1(_804_v41, 2)
			_805_v42 = _nw128
			var _806_v43 _dafny.Array
			_ = _806_v43
			var _nwElement0_25 _dafny.Array = _805_v42
			_ = _nwElement0_25
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(15))
			_ = _nw129
			(_nw129).ArraySet1(_nwElement0_25, 0)
			(_nw129).ArraySet1(_805_v42, 1)
			(_nw129).ArraySet1(_805_v42, 2)
			(_nw129).ArraySet1(_805_v42, 3)
			(_nw129).ArraySet1(_805_v42, 4)
			(_nw129).ArraySet1(_805_v42, 5)
			(_nw129).ArraySet1(_805_v42, 6)
			(_nw129).ArraySet1(_805_v42, 7)
			(_nw129).ArraySet1(_805_v42, 8)
			(_nw129).ArraySet1(_805_v42, 9)
			(_nw129).ArraySet1(_805_v42, 10)
			(_nw129).ArraySet1((Companion_D11_.Create_DC24_(_805_v42)).Dtor_cf36(), 11)
			(_nw129).ArraySet1(_805_v42, 12)
			(_nw129).ArraySet1(_805_v42, 13)
			(_nw129).ArraySet1(_805_v42, 14)
			_806_v43 = _nw129
			var _rhs148 _dafny.Array = _806_v43
			_ = _rhs148
			var _rhs149 _dafny.Int = _this.F10()
			_ = _rhs149
			var _lhs98 *GlobalState = globalState
			_ = _lhs98
			_806_v43 = _rhs148
			_lhs98.F2 = _rhs149
			var _807_v44 _dafny.Int
			_ = _807_v44
			var _808_v45 bool
			_ = _808_v45
			var _809_v46 bool
			_ = _809_v46
			var _810_v47 _dafny.Map
			_ = _810_v47
			var _out58 _dafny.Int
			_ = _out58
			var _out59 bool
			_ = _out59
			var _out60 bool
			_ = _out60
			var _out61 _dafny.Map
			_ = _out61
			_out58, _out59, _out60, _out61 = Companion_Default___.M0(globalState)
			_807_v44 = _out58
			_808_v45 = _out59
			_809_v46 = _out60
			_810_v47 = _out61
		}
		var _811_v48 bool
		_ = _811_v48
		_811_v48 = false
		var _812_v49 *C0
		_ = _812_v49
		var _nw130 *C0 = New_C0_()
		_ = _nw130
		_nw130.Ctor__(_811_v48, _this.F10())
		_812_v49 = _nw130
		var _813_v50 D1
		_ = _813_v50
		_813_v50 = Companion_D1_.Create_DC1_(_811_v48)
		var _814_v51 *C2
		_ = _814_v51
		var _nw131 *C2 = New_C2_()
		_ = _nw131
		_nw131.Ctor__((_812_v49).F12(), (_812_v49).F12())
		_814_v51 = _nw131
		var _815_v52 _dafny.Map
		_ = _815_v52
		_815_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_813_v50).Dtor_cf1(), _814_v51)
		var _816_v53 _dafny.Sequence
		_ = _816_v53
		_816_v53 = _dafny.UnicodeSeqOfUtf8Bytes("ntqkq")
		var _817_v54 _dafny.Set
		_ = _817_v54
		_817_v54 = _dafny.SetOf((_815_v52).Cardinality(), (_812_v49).F13(), (_812_v49).F13(), _dafny.IntOfUint32((_816_v53).Cardinality()), _this.F10())
		var _818_v55 _dafny.Map
		_ = _818_v55
		_818_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_812_v49).F12()), (_817_v54).Cardinality())
		var _819_v56 _dafny.Sequence
		_ = _819_v56
		_819_v56 = _dafny.SeqOf(_818_v55)
		var _820_v57 _dafny.Array
		_ = _820_v57
		var _nwElement0_26 _dafny.Map = _818_v55
		_ = _nwElement0_26
		var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(13))
		_ = _nw132
		(_nw132).ArraySet1(_nwElement0_26, 0)
		(_nw132).ArraySet1(_818_v55, 1)
		(_nw132).ArraySet1(_818_v55, 2)
		(_nw132).ArraySet1(_818_v55, 3)
		(_nw132).ArraySet1(_818_v55, 4)
		(_nw132).ArraySet1(((_819_v56).Select((Companion_Default___.SafeIndex((_812_v49).F13(), _dafny.IntOfUint32((_819_v56).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_818_v55), 5)
		(_nw132).ArraySet1(_818_v55, 6)
		(_nw132).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_811_v48, (_812_v49).F13()), 7)
		(_nw132).ArraySet1((_818_v55).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_811_v48), _this.F10())), 8)
		(_nw132).ArraySet1((_818_v55).Merge(_818_v55), 9)
		(_nw132).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_814_v51).F15(), _dafny.IntOfUint32((_816_v53).Cardinality())), 10)
		(_nw132).ArraySet1((_818_v55).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_814_v51).F15(), (_812_v49).F13())), 11)
		(_nw132).ArraySet1(_818_v55, 12)
		_820_v57 = _nw132
		var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_820_v57), 0))
		_ = _index132
		(_820_v57).ArraySet1(_818_v55, (_index132).Int())
		var _821_v58 _dafny.Array
		_ = _821_v58
		var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
		_ = _nw133
		_821_v58 = _nw133
		var _822_v59 _dafny.CodePoint
		_ = _822_v59
		_822_v59 = _dafny.CodePoint('g')
		var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_821_v58), 0))
		_ = _index133
		(_821_v58).ArraySet1CodePoint(_822_v59, (_index133).Int())
		var _823_v60 D9
		_ = _823_v60
		_823_v60 = Companion_D9_.Create_DC19_(_818_v55)
		var _pat_let_tv4 = _814_v51
		_ = _pat_let_tv4
		var _pat_let_tv5 = _812_v49
		_ = _pat_let_tv5
		var _pat_let_tv6 = _812_v49
		_ = _pat_let_tv6
		var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_820_v57), 0))
		_ = _index134
		var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_821_v58), 0))
		_ = _index135
		var _rhs150 bool = func(_source17 D9) bool {
			if _source17.Is_DC20() {
				var _824___mcc_h6 bool = _source17.Get_().(D9_DC20).Cf29
				_ = _824___mcc_h6
				var _825___mcc_h7 _dafny.CodePoint = _source17.Get_().(D9_DC20).Cf30
				_ = _825___mcc_h7
				var _826_cf30 _dafny.CodePoint = _825___mcc_h7
				_ = _826_cf30
				var _827_cf29 bool = _824___mcc_h6
				_ = _827_cf29
				return (_pat_let_tv4).F15()
			} else if _source17.Is_DC19() {
				var _828___mcc_h8 _dafny.Map = _source17.Get_().(D9_DC19).Cf28
				_ = _828___mcc_h8
				var _829_cf28 _dafny.Map = _828___mcc_h8
				_ = _829_cf28
				return (_pat_let_tv5).F12()
			} else {
				var _830___mcc_h9 D9 = _source17.Get_().(D9_DC21).Cf31
				_ = _830___mcc_h9
				var _831_cf31 D9 = _830___mcc_h9
				_ = _831_cf31
				return (_this.F10()).Cmp((_pat_let_tv6).F13()) != 0
			}
		}((func() D9 {
			if _811_v48 {
				return _823_v60
			}
			return _823_v60
		})())
		_ = _rhs150
		var _rhs151 _dafny.Map = (_818_v55).Merge(_818_v55)
		_ = _rhs151
		var _rhs152 _dafny.CodePoint = _822_v59
		_ = _rhs152
		var _rhs153 _dafny.Int = ((_this.F10()).Plus(_this.F10())).Minus(Companion_Default___.Fm19(_dafny.IntOfUint32((_816_v53).Cardinality()), globalState))
		_ = _rhs153
		var _lhs99 _dafny.Array = _820_v57
		_ = _lhs99
		var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_820_v57), 0))
		_ = _lhs100
		var _lhs101 _dafny.Array = _821_v58
		_ = _lhs101
		var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_821_v58), 0))
		_ = _lhs102
		var _lhs103 *C3 = _this
		_ = _lhs103
		r0 = _rhs150
		(_lhs99).ArraySet1(_rhs151, (_lhs100).Int())
		(_lhs101).ArraySet1CodePoint(_rhs152, (_lhs102).Int())
		_lhs103.F10_set_(_rhs153)
		(globalState).F2 = (_812_v49).F13()
		var _832_v61 _dafny.Sequence
		_ = _832_v61
		_832_v61 = _dafny.SeqOf(_816_v53, _dafny.UnicodeSeqOfUtf8Bytes("hjimyu"))
		var _833_v62 _dafny.Int
		_ = _833_v62
		_833_v62 = (_dafny.Zero).Minus((_812_v49).F13())
		var _source18 D1 = Companion_D1_.Create_DC3_(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_834_v58 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_835_i3 _dafny.Int) _dafny.CodePoint {
				return (Companion_D11_.Create_DC25_((_834_v58).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_834_v58), 0))).Int()))).Dtor_cf37()
			}
		})(_821_v58))), (_832_v61).Select((Companion_Default___.SafeIndex(_this.F10(), _dafny.IntOfUint32((_832_v61).Cardinality()))).Uint32()).(_dafny.Sequence)), _833_v62)
		_ = _source18
		if _source18.Is_DC2() {
			var _836___mcc_h10 bool = _source18.Get_().(D1_DC2).Cf2
			_ = _836___mcc_h10
			var _837_cf2 bool = _836___mcc_h10
			_ = _837_cf2
			_811_v48 = (_812_v49).F12()
			(globalState).F2 = (func() _dafny.Int {
				if (_812_v49).F12() {
					return (_812_v49).F13()
				}
				return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_812_v49).F13(), _814_v51.F14)).Cardinality()).Minus((_812_v49).F13())
			})()
			(globalState).F2 = (_812_v49).F13()
			var _838_v63 _dafny.Sequence
			_ = _838_v63
			_838_v63 = _dafny.SeqOf((_814_v51).F15(), _814_v51.F14, (_814_v51).F15(), _811_v48)
			var _839_v64 *C2
			_ = _839_v64
			var _nw134 *C2 = New_C2_()
			_ = _nw134
			_nw134.Ctor__(!_dafny.Companion_Sequence_.Equal(_838_v63, _838_v63), (_812_v49).F12())
			_839_v64 = _nw134
		} else if _source18.Is_DC3() {
			var _840___mcc_h11 bool = _source18.Get_().(D1_DC3).Cf3
			_ = _840___mcc_h11
			var _841___mcc_h12 _dafny.Int = _source18.Get_().(D1_DC3).Cf4
			_ = _841___mcc_h12
			var _842_cf4 _dafny.Int = _841___mcc_h12
			_ = _842_cf4
			var _843_cf3 bool = _840___mcc_h11
			_ = _843_cf3
			var _844_v65 *C0
			_ = _844_v65
			var _nw135 *C0 = New_C0_()
			_ = _nw135
			_nw135.Ctor__(_843_cf3, _this.F10())
			_844_v65 = _nw135
			var _845_v66 D10
			_ = _845_v66
			_845_v66 = Companion_D10_.Create_DC23_((_812_v49).F13(), (_812_v49).F12(), _843_cf3)
			_845_v66 = _845_v66
			_822_v59 = _822_v59
			_811_v48 = _843_cf3
		} else {
			var _846___mcc_h13 bool = _source18.Get_().(D1_DC1).Cf1
			_ = _846___mcc_h13
			var _847_cf1 bool = _846___mcc_h13
			_ = _847_cf1
			var _848_v67 _dafny.Set
			_ = _848_v67
			_848_v67 = _dafny.SetOf((_812_v49).F12())
			(_814_v51).F14 = (!(_847_cf1)) || ((_848_v67).IsDisjointFrom(_848_v67))
			r3 = !((_814_v51).F15())
			var _849_v68 _dafny.Array
			_ = _849_v68
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_16
			var _nw136 _dafny.Array
			_ = _nw136
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw136 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) bool = (func(_850_v49 *C0) func(_dafny.Int) bool {
					return func(_851_i4 _dafny.Int) bool {
						return !((_850_v49).F12())
					}
				})(_812_v49)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw136 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw136).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw136).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_849_v68 = _nw136
			_849_v68 = _849_v68
			var _852_v69 _dafny.Int
			_ = _852_v69
			var _853_v70 bool
			_ = _853_v70
			var _854_v71 bool
			_ = _854_v71
			var _855_v72 _dafny.Map
			_ = _855_v72
			var _out62 _dafny.Int
			_ = _out62
			var _out63 bool
			_ = _out63
			var _out64 bool
			_ = _out64
			var _out65 _dafny.Map
			_ = _out65
			_out62, _out63, _out64, _out65 = Companion_Default___.M0(globalState)
			_852_v69 = _out62
			_853_v70 = _out63
			_854_v71 = _out64
			_855_v72 = _out65
		}
		(_814_v51).F14 = (_814_v51).F15()
		r0 = ((_812_v49).F13()).Cmp((_812_v49).F13()) <= 0
		r1 = (_812_v49).F13()
		r2 = _811_v48
		var _856_v73 _dafny.MultiSet
		_ = _856_v73
		_856_v73 = _dafny.MultiSetOf(_811_v48)
		r3 = (_814_v51).Fm13(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_814_v51.F14), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_811_v48, !(_814_v51.F14), !((_812_v49).F12()), (_814_v51).F15()), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_812_v49).F13()), _dafny.IntOfUint32((_dafny.SeqOf(_811_v48, !(_814_v51.F14), !((_812_v49).F12()), (_814_v51).F15())).Cardinality()))).Uint32(), (_814_v51).F15()))).Cardinality()), (_856_v73).IsSubsetOf(_856_v73), _dafny.IntOfInt64(22), globalState)
		return r0, r1, r2, r3
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
