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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) bool {
	var _source0 D0 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(801), (_dafny.MultiSetOf(_dafny.IntOfInt64(672), _dafny.IntOfInt64(917), (_dafny.MultiSetOf(false)).Cardinality())).Cardinality(), _dafny.IntOfInt64(458))
	_ = _source0
	if _source0.Is_DC1() {
		var _0___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _0___mcc_h0
		var _1___mcc_h1 _dafny.Int = _source0.Get_().(D0_DC1).Cf2
		_ = _1___mcc_h1
		var _2___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC1).Cf3
		_ = _2___mcc_h2
		var _3_cf3 _dafny.Int = _2___mcc_h2
		_ = _3_cf3
		var _4_cf2 _dafny.Int = _1___mcc_h1
		_ = _4_cf2
		var _5_cf1 _dafny.Int = _0___mcc_h0
		_ = _5_cf1
		return _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("bielrkk"), _dafny.UnicodeSeqOfUtf8Bytes("w"))
	} else if _source0.Is_DC2() {
		var _6___mcc_h3 _dafny.Sequence = _source0.Get_().(D0_DC2).Cf4
		_ = _6___mcc_h3
		var _7___mcc_h4 _dafny.Int = _source0.Get_().(D0_DC2).Cf5
		_ = _7___mcc_h4
		var _8___mcc_h5 _dafny.Int = _source0.Get_().(D0_DC2).Cf6
		_ = _8___mcc_h5
		var _9_cf6 _dafny.Int = _8___mcc_h5
		_ = _9_cf6
		var _10_cf5 _dafny.Int = _7___mcc_h4
		_ = _10_cf5
		var _11_cf4 _dafny.Sequence = _6___mcc_h3
		_ = _11_cf4
		return !(false)
	} else if _source0.Is_DC0() {
		var _12___mcc_h6 bool = _source0.Get_().(D0_DC0).Cf0
		_ = _12___mcc_h6
		var _13_cf0 bool = _12___mcc_h6
		_ = _13_cf0
		return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		}))).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("phiihuc")).Cardinality())) < 0
	} else {
		var _15___mcc_h7 D0 = _source0.Get_().(D0_DC3).Cf7
		_ = _15___mcc_h7
		var _16_cf7 D0 = _15___mcc_h7
		_ = _16_cf7
		return true
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-879))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return (func() _dafny.CodePoint {
			if true {
				return _dafny.CodePoint('p')
			}
			return _dafny.CodePoint('n')
		})()
	}))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((func() D5 {
		if true {
			return Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(641), false)))
		}
		return Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(219), false)))
	})()).Dtor_cf20()
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if !(true) || (false) {
		if false {
			return _dafny.CodePoint('j')
		} else {
			return _dafny.CodePoint('g')
		}
	} else {
		return _dafny.CodePoint('y')
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Union(_dafny.SetOf(!(!(!(!(true))))))).Union(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(521)
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((Companion_D1_.Create_DC5_(!(false), !(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false)))).Cardinality())).Cardinality())).Dtor_cf10())), true)).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))
	})())
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(523)).Minus((_dafny.SetOf(_dafny.IntOfInt64(185))).Cardinality()), (_dafny.SetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.SeqOf(false, !(true), !(false)))).Cardinality()))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, globalState *GlobalState) _dafny.Set {
	var _source1 D0 = Companion_D0_.Create_DC0_(false)
	_ = _source1
	if _source1.Is_DC1() {
		var _18___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC1).Cf1
		_ = _18___mcc_h0
		var _19___mcc_h1 _dafny.Int = _source1.Get_().(D0_DC1).Cf2
		_ = _19___mcc_h1
		var _20___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC1).Cf3
		_ = _20___mcc_h2
		var _21_cf3 _dafny.Int = _20___mcc_h2
		_ = _21_cf3
		var _22_cf2 _dafny.Int = _19___mcc_h1
		_ = _22_cf2
		var _23_cf1 _dafny.Int = _18___mcc_h0
		_ = _23_cf1
		return func() _dafny.Set {
			var _coll0 = _dafny.NewBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_(false))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _24_v0 D0
				_24_v0 = interface{}(_compr_0).(D0)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_(false)), _24_v0) {
					_coll0.Add(_24_v0)
				}
			}
			return _coll0.ToSet()
		}()
	} else if _source1.Is_DC2() {
		var _25___mcc_h3 _dafny.Sequence = _source1.Get_().(D0_DC2).Cf4
		_ = _25___mcc_h3
		var _26___mcc_h4 _dafny.Int = _source1.Get_().(D0_DC2).Cf5
		_ = _26___mcc_h4
		var _27___mcc_h5 _dafny.Int = _source1.Get_().(D0_DC2).Cf6
		_ = _27___mcc_h5
		var _28_cf6 _dafny.Int = _27___mcc_h5
		_ = _28_cf6
		var _29_cf5 _dafny.Int = _26___mcc_h4
		_ = _29_cf5
		var _30_cf4 _dafny.Sequence = _25___mcc_h3
		_ = _30_cf4
		if !(false) {
			return func() _dafny.Set {
				var _coll1 = _dafny.NewBuilder()
				_ = _coll1
				for _iter1 := _dafny.Iterate((_dafny.SetOf(Companion_D0_.Create_DC0_(true))).Elements()); ; {
					_compr_1, _ok1 := _iter1()
					if !_ok1 {
						break
					}
					var _31_v1 D0
					_31_v1 = interface{}(_compr_1).(D0)
					if (_dafny.SetOf(Companion_D0_.Create_DC0_(true))).Contains(_31_v1) {
						_coll1.Add(_31_v1)
					}
				}
				return _coll1.ToSet()
			}()
		} else {
			return func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC0_(false), Companion_D0_.Create_DC0_(false))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _32_v2 D0
					_32_v2 = interface{}(_compr_2).(D0)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC0_(false), Companion_D0_.Create_DC0_(false)), _32_v2) {
						_coll2.Add(_32_v2)
					}
				}
				return _coll2.ToSet()
			}()
		}
	} else if _source1.Is_DC0() {
		var _33___mcc_h6 bool = _source1.Get_().(D0_DC0).Cf0
		_ = _33___mcc_h6
		var _34_cf0 bool = _33___mcc_h6
		_ = _34_cf0
		return func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(226))).Uint32(), func(coer2 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}((func(_35_cf0 bool) func(_dafny.Int) D0 {
				return func(_36_i0 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_35_cf0)
				}
			})(_34_cf0)))).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _37_v3 D0
				_37_v3 = interface{}(_compr_3).(D0)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(226))).Uint32(), func(coer3 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}((func(_38_cf0 bool) func(_dafny.Int) D0 {
					return func(_36_i0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_38_cf0)
					}
				})(_34_cf0))), _37_v3) {
					_coll3.Add(_37_v3)
				}
			}
			return _coll3.ToSet()
		}()
	} else {
		var _39___mcc_h7 D0 = _source1.Get_().(D0_DC3).Cf7
		_ = _39___mcc_h7
		var _40_cf7 D0 = _39___mcc_h7
		_ = _40_cf7
		return _dafny.SetOf(Companion_D0_.Create_DC0_(false), Companion_D0_.Create_DC0_(!(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(697))).Merge(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_41_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _42_v0 _dafny.CodePoint
			_42_v0 = interface{}(_compr_4).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_41_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			})), _42_v0) {
				_coll4.Add(_42_v0, (_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.SetOf(Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(808), false))))).Cardinality())).Cardinality(), _dafny.IntOfInt64(720))).Cardinality())
			}
		}
		return _coll4.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm10(p0 D1, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _dafny.IntOfInt64(498))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.IntOfInt64(330))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(515))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_43_i0 _dafny.Int) _dafny.Map {
		return func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('t'))).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _44_v0 _dafny.CodePoint
				_44_v0 = interface{}(_compr_5).(_dafny.CodePoint)
				if (_dafny.SetOf(_dafny.CodePoint('t'))).Contains(_44_v0) {
					_coll5.Add(_44_v0, _dafny.IntOfInt64(357))
				}
			}
			return _coll5.ToMap()
		}()
	}))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC8_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(328))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_45_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_46_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('x'), _dafny.CodePoint('c'), _dafny.CodePoint('u')), _dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('f')))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yo")).Cardinality()))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-126))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) D1 {
	var _source2 D1 = Companion_D1_.Create_DC6_()
	_ = _source2
	if _source2.Is_DC5() {
		var _47___mcc_h0 bool = _source2.Get_().(D1_DC5).Cf9
		_ = _47___mcc_h0
		var _48___mcc_h1 bool = _source2.Get_().(D1_DC5).Cf10
		_ = _48___mcc_h1
		var _49___mcc_h2 _dafny.Int = _source2.Get_().(D1_DC5).Cf11
		_ = _49___mcc_h2
		var _50_cf11 _dafny.Int = _49___mcc_h2
		_ = _50_cf11
		var _51_cf10 bool = _48___mcc_h1
		_ = _51_cf10
		var _52_cf9 bool = _47___mcc_h0
		_ = _52_cf9
		if _52_cf9 {
			return Companion_D1_.Create_DC6_()
		} else {
			return Companion_D1_.Create_DC6_()
		}
	} else if _source2.Is_DC6() {
		return Companion_D1_.Create_DC6_()
	} else {
		var _53___mcc_h3 _dafny.Array = _source2.Get_().(D1_DC4).Cf8
		_ = _53___mcc_h3
		var _54_cf8 _dafny.Array = _53___mcc_h3
		_ = _54_cf8
		return Companion_D1_.Create_DC6_()
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, p3 D0, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC10_(_dafny.SeqOf(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(_dafny.UnicodeSeqOfUtf8Bytes("givcgweu"), _dafny.IntOfInt64(-147), _dafny.IntOfInt64(324)))))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _55_globalState *GlobalState
	_ = _55_globalState
	var _nw0 *GlobalState = New_GlobalState_()
	_ = _nw0
	_nw0.Ctor__(_dafny.IntOfInt64(825), _dafny.One, true)
	_55_globalState = _nw0
	var _56_v0 bool
	_ = _56_v0
	_56_v0 = true
	var _57_v1 _dafny.Array
	_ = _57_v1
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
	_ = _nw1
	_57_v1 = _nw1
	var _58_v2 _dafny.Int
	_ = _58_v2
	_58_v2 = _dafny.IntOfInt64(891)
	var _59_v3 _dafny.Map
	_ = _59_v3
	_59_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
		if _56_v0 {
			return _57_v1
		}
		return _57_v1
	})(), (_dafny.Zero).Minus(_58_v2))
	_59_v3 = _59_v3
	_56_v0 = !(Companion_Default___.Fm0(_58_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("if")).Cardinality()), _55_globalState))
	var _60_v4 _dafny.Array
	_ = _60_v4
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Map = (func(_61_v2 _dafny.Int, _62_v0 bool) func(_dafny.Int) _dafny.Map {
			return func(_63_i0 _dafny.Int) _dafny.Map {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v2, _62_v0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v2, _62_v0))
			}
		})(_58_v2, _56_v0)
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
	_60_v4 = _nw2
	var _64_v5 _dafny.Map
	_ = _64_v5
	_64_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v2, _56_v0)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_60_v4), 0))
	_ = _index0
	(_60_v4).ArraySet1(_64_v5, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_60_v4), 0))
	_ = _index1
	(_60_v4).ArraySet1(_64_v5, (_index1).Int())
	var _65_v6 _dafny.MultiSet
	_ = _65_v6
	_65_v6 = _dafny.MultiSetOf(false)
	_56_v0 = !(((_58_v2).Minus(_58_v2)).Cmp((_58_v2).Plus((_65_v6).Cardinality())) > 0)
	var _66_v7 D0
	_ = _66_v7
	_66_v7 = Companion_D0_.Create_DC0_(_56_v0)
	var _67_v8 *C0
	_ = _67_v8
	var _nw3 *C0 = New_C0_()
	_ = _nw3
	_nw3.Ctor__(_66_v7, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, _57_v1), _57_v1, _58_v2)
	_67_v8 = _nw3
	var _68_v9 _dafny.CodePoint
	_ = _68_v9
	_68_v9 = _dafny.CodePoint('q')
	var _69_v10 T0
	_ = _69_v10
	var _nw4 *C0 = New_C0_()
	_ = _nw4
	_nw4.Ctor__((_67_v8).F7(), (_67_v8).F8(), _57_v1, _58_v2)
	_69_v10 = _nw4
	var _70_v11 _dafny.Map
	_ = _70_v11
	_70_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _69_v10)
	var _71_v12 _dafny.Map
	_ = _71_v12
	_71_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v5, !(!(_70_v11).Contains(_68_v9)))
	_71_v12 = _71_v12
	var _72_v13 _dafny.Array
	_ = _72_v13
	var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(16))
	_ = _nw5
	_72_v13 = _nw5
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_72_v13), 0))); ; {
		_guard_loop_0, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _73_i1 _dafny.Int
		_73_i1 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_73_i1).Sign() != -1) && ((_73_i1).Cmp(_dafny.ArrayLen((_72_v13), 0)) < 0)) {
			(_72_v13).ArraySet1(((_dafny.SetOf(false, _56_v0)).Intersection(_dafny.SetOf(_56_v0, true, _56_v0))).Union(_dafny.SetOf(_56_v0)), (_73_i1).Int())
		}
	}
	var _rhs0 _dafny.Int = _dafny.IntOfInt64(551)
	_ = _rhs0
	var _rhs1 _dafny.Int = _dafny.IntOfInt64(-189)
	_ = _rhs1
	var _lhs0 *GlobalState = _55_globalState
	_ = _lhs0
	_lhs0.F1 = _rhs0
	_58_v2 = _rhs1
	_56_v0 = _56_v0
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
	_ = _index2
	(_57_v1).ArraySet1(_58_v2, (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
	_ = _index3
	(_57_v1).ArraySet1((func() _dafny.Int {
		if (_65_v6).Contains((_58_v2).Cmp(_58_v2) < 0) {
			return (_65_v6).Multiplicity((_58_v2).Cmp(_58_v2) < 0)
		}
		return (_69_v10).F4()
	})(), (_index3).Int())
	var _74_v14 _dafny.Array
	_ = _74_v14
	var _nw6 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
	_ = _nw6
	_74_v14 = _nw6
	var _75_v15 _dafny.Map
	_ = _75_v15
	_75_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v14, _56_v0)
	var _76_v16 _dafny.Map
	_ = _76_v16
	_76_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), (_69_v10).F4())
	var _77_v17 _dafny.Sequence
	_ = _77_v17
	_77_v17 = _dafny.SeqOf(_56_v0)
	var _78_v18 D0
	_ = _78_v18
	_78_v18 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_77_v17).Cardinality()), (_69_v10).F4(), _dafny.IntOfUint32((_dafny.SeqOf(_58_v2)).Cardinality()))
	var _79_v19 _dafny.Map
	_ = _79_v19
	_79_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v10, _78_v18)
	_75_v15 = (_75_v15).Update(_74_v14, Companion_Default___.Fm0((_76_v16).Cardinality(), (_79_v19).Cardinality(), _55_globalState))
	var _80_v20 _dafny.Array
	_ = _80_v20
	var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
	_ = _nw7
	_80_v20 = _nw7
	var _81_v21 _dafny.Map
	_ = _81_v21
	_81_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v10, _80_v20)
	_81_v21 = (_81_v21).Update(_69_v10, _80_v20)
	var _82_v22 D1
	_ = _82_v22
	_82_v22 = Companion_D1_.Create_DC4_(_57_v1)
	_82_v22 = _82_v22
	var _83_v23 _dafny.Sequence
	_ = _83_v23
	_83_v23 = _dafny.UnicodeSeqOfUtf8Bytes("gpuqdrl")
	var _84_v24 _dafny.Array
	_ = _84_v24
	var _nw8 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
	_ = _nw8
	_84_v24 = _nw8
	var _85_v25 *C1
	_ = _85_v25
	var _nw9 *C1 = New_C1_()
	_ = _nw9
	_nw9.Ctor__((_dafny.CodePoint('y')), _84_v24, (_69_v10).F3(), _58_v2)
	_85_v25 = _nw9
	var _rhs2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_83_v23, _83_v23)
	_ = _rhs2
	var _rhs3 *C1 = _85_v25
	_ = _rhs3
	_83_v23 = _rhs2
	_85_v25 = _rhs3
	var _86_v26 D0
	_ = _86_v26
	_86_v26 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(_83_v23, Companion_Default___.Fm5(false, _55_globalState), _58_v2))
	var _source3 D0 = _86_v26
	_ = _source3
	if _source3.Is_DC1() {
		var _87___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
		_ = _87___mcc_h0
		var _88___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
		_ = _88___mcc_h1
		var _89___mcc_h2 _dafny.Int = _source3.Get_().(D0_DC1).Cf3
		_ = _89___mcc_h2
		var _90_cf3 _dafny.Int = _89___mcc_h2
		_ = _90_cf3
		var _91_cf2 _dafny.Int = _88___mcc_h1
		_ = _91_cf2
		var _92_cf1 _dafny.Int = _87___mcc_h0
		_ = _92_cf1
		var _93_v27 _dafny.MultiSet
		_ = _93_v27
		_93_v27 = _dafny.MultiSetOf(_90_cf3, (_69_v10).F4())
		if (_93_v27).IsSubsetOf(_dafny.MultiSetOf((_69_v10).F4())) {
			var _94_v28 _dafny.Map
			_ = _94_v28
			_94_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_cf3, _85_v25)
			var _95_v29 _dafny.Map
			_ = _95_v29
			_95_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_94_v28).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _85_v25)), !(_56_v0))
			_95_v29 = _95_v29
			var _96_v30 _dafny.Sequence
			_ = _96_v30
			_96_v30 = _dafny.SeqOf(_86_v26)
			var _97_v31 D4
			_ = _97_v31
			_97_v31 = Companion_D4_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-50))).Uint32(), func(coer9 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_98_v10 T0, _99_v0 bool, _100_v1 _dafny.Array) func(_dafny.Int) D0 {
				return func(_101_i2 _dafny.Int) D0 {
					return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_((_98_v10).F4(), (_98_v10).F4(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v0, (_100_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_100_v1), 0))).Int()).(_dafny.Int))).Cardinality()))
				}
			})(_69_v10, _56_v0, _57_v1))))
			var _102_v32 _dafny.Map
			_ = _102_v32
			_102_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, _92_cf1)
			var _103_v33 _dafny.Sequence
			_ = _103_v33
			_103_v33 = _dafny.SeqOf(_dafny.IntOfUint32((_83_v23).Cardinality()))
			var _104_v34 _dafny.MultiSet
			_ = _104_v34
			_104_v34 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_96_v30, _dafny.Companion_Sequence_.Update((_97_v31).Dtor_cf16(), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_102_v32).Contains(_56_v0) {
					return (_102_v32).Get(_56_v0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_103_v33).Cardinality())
			})(), _dafny.IntOfUint32(((_97_v31).Dtor_cf16()).Cardinality()))).Uint32(), _86_v26)), _96_v30)
			var _105_v35 D2
			_ = _105_v35
			_105_v35 = Companion_D2_.Create_DC7_(_77_v17)
			var _rhs4 bool = _56_v0
			_ = _rhs4
			var _rhs5 _dafny.Int = (_69_v10).F4()
			_ = _rhs5
			var _rhs6 _dafny.MultiSet = _104_v34
			_ = _rhs6
			var _rhs7 _dafny.MultiSet = _dafny.MultiSetFromSeq((_105_v35).Dtor_cf12())
			_ = _rhs7
			_56_v0 = _rhs4
			_91_cf2 = _rhs5
			_104_v34 = _rhs6
			_65_v6 = _rhs7
			_56_v0 = _56_v0
			var _106_v36 bool
			_ = _106_v36
			var _107_v37 bool
			_ = _107_v37
			var _108_v38 bool
			_ = _108_v38
			var _out0 bool
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 bool
			_ = _out2
			_out0, _out1, _out2 = (_85_v25).M0(_55_globalState)
			_106_v36 = _out0
			_107_v37 = _out1
			_108_v38 = _out2
			_108_v38 = _106_v36
		} else {
			var _109_v39 _dafny.Sequence
			_ = _109_v39
			_109_v39 = _dafny.SeqOf(_69_v10, _69_v10)
			_69_v10 = (func() T0 {
				if _56_v0 {
					return (_109_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-846), _dafny.IntOfUint32((_109_v39).Cardinality()))).Uint32()).(T0)
				}
				return _69_v10
			})()
			var _110_v40 _dafny.Sequence
			_ = _110_v40
			_110_v40 = _dafny.SeqOf((_93_v27).Cardinality())
			var _111_v41 _dafny.Sequence
			_ = _111_v41
			_111_v41 = _dafny.SeqOf(_93_v27, _dafny.MultiSetOf(_dafny.IntOfUint32((_110_v40).Cardinality())))
			var _112_v42 _dafny.Map
			_ = _112_v42
			_112_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_91_cf2), (_111_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.IntOfUint32((_111_v41).Cardinality()))).Uint32()).(_dafny.MultiSet))
			_93_v27 = (func() _dafny.MultiSet {
				if (_112_v42).Contains(_dafny.IntOfInt64(855)) {
					return (_112_v42).Get(_dafny.IntOfInt64(855)).(_dafny.MultiSet)
				}
				return _93_v27
			})()
			var _113_v43 D2
			_ = _113_v43
			_113_v43 = Companion_D2_.Create_DC8_(_83_v23, _92_cf1)
			var _pat_let_tv0 = _83_v23
			_ = _pat_let_tv0
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index4
			var _rhs8 bool = Companion_Default___.Fm0((_69_v10).F4(), _90_cf3, _55_globalState)
			_ = _rhs8
			var _rhs9 D2 = func(_pat_let0_0 D2) D2 {
				return func(_114_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let1_0 _dafny.Sequence) D2 {
						return func(_115_dt__update_hcf13_h0 _dafny.Sequence) D2 {
							return Companion_D2_.Create_DC8_(_115_dt__update_hcf13_h0, (_114_dt__update__tmp_h0).Dtor_cf14())
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(Companion_Default___.Fm11(_56_v0, _55_globalState))
			_ = _rhs9
			var _rhs10 _dafny.CodePoint = Companion_Default___.Fm3(_68_v9, _92_cf1, _55_globalState)
			_ = _rhs10
			var _rhs11 _dafny.Int = (_69_v10).F4()
			_ = _rhs11
			var _lhs1 _dafny.Array = _57_v1
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _lhs2
			_56_v0 = _rhs8
			_113_v43 = _rhs9
			_68_v9 = _rhs10
			(_lhs1).ArraySet1(_rhs11, (_lhs2).Int())
			_84_v24 = (_85_v25).F6()
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index5
			(_57_v1).ArraySet1((_dafny.Zero).Minus((_69_v10).F4()), (_index5).Int())
		}
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
		_ = _index6
		(_57_v1).ArraySet1(((_69_v10).F4()).Minus((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)), (_index6).Int())
		if !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_56_v0), (Companion_Default___.SafeIndex(_58_v2, _dafny.IntOfUint32((_dafny.SeqOf(_56_v0)).Cardinality()))).Uint32(), _56_v0), (Companion_Default___.SafeIndex((_69_v10).F4(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_56_v0), (Companion_Default___.SafeIndex(_58_v2, _dafny.IntOfUint32((_dafny.SeqOf(_56_v0)).Cardinality()))).Uint32(), _56_v0)).Cardinality()))).Uint32(), _56_v0), _56_v0) {
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen(((_85_v25).F6()), 0))
			_ = _index7
			((_85_v25).F6()).ArraySet1(_56_v0, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen(((_85_v25).F6()), 0))
			_ = _index8
			((_85_v25).F6()).ArraySet1(Companion_Default___.Fm0(_91_cf2, ((_69_v10).F4()).Minus(_91_cf2), _55_globalState), (_index8).Int())
			_76_v16 = (_76_v16).Update(((_60_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_60_v4), 0))).Int()).(_dafny.Map)).Cardinality(), _92_cf1)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index9
			(_57_v1).ArraySet1(_dafny.IntOfInt64(656), (_index9).Int())
			var _116_v44 _dafny.Sequence
			_ = _116_v44
			var _out3 _dafny.Sequence
			_ = _out3
			_out3 = (_85_v25).M1(false, _58_v2, Companion_Default___.Fm5(!(_56_v0), _55_globalState), _55_globalState)
			_116_v44 = _out3
			var _117_v45 _dafny.MultiSet
			_ = _117_v45
			_117_v45 = _dafny.MultiSetOf((_69_v10).F3())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen(((_85_v25).F6()), 0))
			_ = _index10
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen(((_85_v25).F6()), 0))
			_ = _index11
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index12
			var _rhs12 bool = !(Companion_Default___.Fm0(_dafny.IntOfUint32((_83_v23).Cardinality()), (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _55_globalState))
			_ = _rhs12
			var _rhs13 bool = !(!((_117_v45).Difference(_117_v45)).Contains((_69_v10).F3()))
			_ = _rhs13
			var _rhs14 _dafny.Int = _dafny.IntOfUint32((_116_v44).Cardinality())
			_ = _rhs14
			var _lhs3 _dafny.Array = (_85_v25).F6()
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen(((_85_v25).F6()), 0))
			_ = _lhs4
			var _lhs5 _dafny.Array = (_85_v25).F6()
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen(((_85_v25).F6()), 0))
			_ = _lhs6
			var _lhs7 _dafny.Array = _57_v1
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _lhs8
			(_lhs3).ArraySet1(_rhs12, (_lhs4).Int())
			(_lhs5).ArraySet1(_rhs13, (_lhs6).Int())
			(_lhs7).ArraySet1(_rhs14, (_lhs8).Int())
		} else {
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index13
			(_57_v1).ArraySet1(Companion_Default___.SafeDivisionInt(_90_cf3, (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)), (_index13).Int())
			_56_v0 = _56_v0
			var _118_v46 _dafny.Sequence
			_ = _118_v46
			_118_v46 = _dafny.SeqOf(_83_v23, _83_v23, _83_v23, _83_v23)
			var _119_v47 _dafny.Map
			_ = _119_v47
			_119_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, (_69_v10).F4())
			var _120_v48 _dafny.Array
			_ = _120_v48
			var _nwElement0_0 _dafny.Sequence = (func() _dafny.Sequence {
				if _56_v0 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-279))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg10 _dafny.Int) interface{} {
							return coer10(arg10)
						}
					}((func(_121_v25 *C1) func(_dafny.Int) _dafny.CodePoint {
						return func(_122_i3 _dafny.Int) _dafny.CodePoint {
							return (_121_v25).F5()
						}
					})(_85_v25)))
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("avwjo")
			})()
			_ = _nwElement0_0
			var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(18))
			_ = _nw10
			(_nw10).ArraySet1(_nwElement0_0, 0)
			(_nw10).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ntbyypb"), 1)
			(_nw10).ArraySet1(_83_v23, 2)
			(_nw10).ArraySet1((_118_v46).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_119_v47).Contains(_56_v0) {
					return (_119_v47).Get(_56_v0).(_dafny.Int)
				}
				return Companion_Default___.Fm5(_56_v0, _55_globalState)
			})(), _dafny.IntOfUint32((_118_v46).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
			(_nw10).ArraySet1(Companion_Default___.Fm1(_56_v0, _56_v0, _55_globalState), 4)
			(_nw10).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dddi"), 5)
			(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_83_v23, _83_v23), 6)
			(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gnfds"), _83_v23), 7)
			(_nw10).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-336))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_123_v25 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_124_i4 _dafny.Int) _dafny.CodePoint {
					return (_123_v25).F5()
				}
			})(_85_v25))), 8)
			(_nw10).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("c"), 9)
			(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_83_v23, _83_v23), 10)
			(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_83_v23, _dafny.UnicodeSeqOfUtf8Bytes("ddtjgujl")), 11)
			(_nw10).ArraySet1(_83_v23, 12)
			(_nw10).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(308))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_125_v25 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_126_i5 _dafny.Int) _dafny.CodePoint {
					return (_125_v25).F5()
				}
			})(_85_v25))), 13)
			(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_83_v23, Companion_Default___.Fm1(_56_v0, !(_56_v0), _55_globalState)), 14)
			(_nw10).ArraySet1(_83_v23, 15)
			(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("my"), _83_v23), 16)
			(_nw10).ArraySet1(_83_v23, 17)
			_120_v48 = _nw10
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_120_v48), 0))
			_ = _index14
			(_120_v48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_83_v23, _83_v23), (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_120_v48), 0))
			_ = _index15
			(_120_v48).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_83_v23, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_127_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			}))), (Companion_Default___.SafeIndex(_91_cf2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_83_v23, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_128_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})))).Cardinality()))).Uint32(), (_85_v25).F5()), (_index15).Int())
			_69_v10 = _69_v10
			var _129_v49 _dafny.Sequence
			_ = _129_v49
			_129_v49 = _dafny.SeqOf(_92_cf1)
			var _130_v50 _dafny.Sequence
			_ = _130_v50
			var _out4 _dafny.Sequence
			_ = _out4
			_out4 = (_85_v25).M1(_56_v0, (_dafny.Zero).Minus((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)), (_129_v49).Select((Companion_Default___.SafeIndex(_91_cf2, _dafny.IntOfUint32((_129_v49).Cardinality()))).Uint32()).(_dafny.Int), _55_globalState)
			_130_v50 = _out4
		}
		var _131_v51 _dafny.Map
		_ = _131_v51
		_131_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v17, !_dafny.Companion_Sequence_.Equal(_83_v23, _dafny.SeqOf((_83_v23).Select((Companion_Default___.SafeIndex(_90_cf3, _dafny.IntOfUint32((_83_v23).Cardinality()))).Uint32()).(_dafny.CodePoint))))
		var _pat_let_tv1 = _56_v0
		_ = _pat_let_tv1
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
		_ = _index16
		var _rhs15 bool = (func() bool {
			if (_131_v51).Contains(_dafny.SeqOf(_56_v0)) {
				return (_131_v51).Get(_dafny.SeqOf(_56_v0)).(bool)
			}
			return (func(_pat_let2_0 D0) D0 {
				return func(_132_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let3_0 bool) D0 {
						return func(_133_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_133_dt__update_hcf0_h0)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let2_0)
			}(_66_v7)).Dtor_cf0()
		})()
		_ = _rhs15
		var _rhs16 _dafny.Int = Companion_Default___.Fm5(_56_v0, _55_globalState)
		_ = _rhs16
		var _lhs9 _dafny.Array = _57_v1
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
		_ = _lhs10
		_56_v0 = _rhs15
		(_lhs9).ArraySet1(_rhs16, (_lhs10).Int())
	} else if _source3.Is_DC2() {
		var _134___mcc_h3 _dafny.Sequence = _source3.Get_().(D0_DC2).Cf4
		_ = _134___mcc_h3
		var _135___mcc_h4 _dafny.Int = _source3.Get_().(D0_DC2).Cf5
		_ = _135___mcc_h4
		var _136___mcc_h5 _dafny.Int = _source3.Get_().(D0_DC2).Cf6
		_ = _136___mcc_h5
		var _137_cf6 _dafny.Int = _136___mcc_h5
		_ = _137_cf6
		var _138_cf5 _dafny.Int = _135___mcc_h4
		_ = _138_cf5
		var _139_cf4 _dafny.Sequence = _134___mcc_h3
		_ = _139_cf4
		var _140_v52 _dafny.Sequence
		_ = _140_v52
		_140_v52 = _dafny.SeqOf((_69_v10).F3())
		_57_v1 = (func() _dafny.Array {
			if ((_67_v8).F8()).Contains(!(_56_v0)) {
				return ((_67_v8).F8()).Get(!(_56_v0)).(_dafny.Array)
			}
			return (_140_v52).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_69_v10).F4()), _dafny.IntOfUint32((_140_v52).Cardinality()))).Uint32()).(_dafny.Array)
		})()
		if _56_v0 {
			var _141_v53 *C1
			_ = _141_v53
			var _nw11 *C1 = New_C1_()
			_ = _nw11
			_nw11.Ctor__((_85_v25).F5(), (_85_v25).F6(), (_69_v10).F3(), _dafny.IntOfUint32((_77_v17).Cardinality()))
			_141_v53 = _nw11
			_57_v1 = (_69_v10).F3()
			_139_cf4 = _83_v23
			var _142_v54 _dafny.Array
			_ = _142_v54
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_1
			var _nw12 _dafny.Array
			_ = _nw12
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw12 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = (func(_143_v25 *C1) func(_dafny.Int) _dafny.Sequence {
					return func(_144_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_143_v25).F5())
					}
				})(_85_v25)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw12 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw12).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw12).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_142_v54 = _nw12
			var _rhs17 _dafny.Int = (_dafny.Zero).Minus(_138_cf5)
			_ = _rhs17
			var _rhs18 _dafny.Array = _142_v54
			_ = _rhs18
			var _rhs19 _dafny.Int = (_69_v10).F4()
			_ = _rhs19
			var _rhs20 bool = _56_v0
			_ = _rhs20
			var _rhs21 bool = ((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)).Cmp((_69_v10).F4()) < 0
			_ = _rhs21
			var _lhs11 *GlobalState = _55_globalState
			_ = _lhs11
			_lhs11.F1 = _rhs17
			_142_v54 = _rhs18
			_58_v2 = _rhs19
			_56_v0 = _rhs20
			_56_v0 = _rhs21
			var _145_v55 D2
			_ = _145_v55
			_145_v55 = Companion_D2_.Create_DC8_(_83_v23, _138_cf5)
			var _146_v56 _dafny.MultiSet
			_ = _146_v56
			_146_v56 = _dafny.MultiSetOf((_145_v55).Dtor_cf14())
			var _147_v57 _dafny.Map
			_ = _147_v57
			_147_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v53, _146_v56)
			_56_v0 = !(_56_v0) || ((_dafny.SetOf((_147_v57).Cardinality())).IsDisjointFrom(_dafny.SetOf((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _138_cf5)))
		} else {
			_69_v10 = _69_v10
			_77_v17 = _dafny.Companion_Sequence_.Concatenate(_77_v17, _77_v17)
			var _148_v58 bool
			_ = _148_v58
			var _149_v59 bool
			_ = _149_v59
			var _150_v60 bool
			_ = _150_v60
			var _out5 bool
			_ = _out5
			var _out6 bool
			_ = _out6
			var _out7 bool
			_ = _out7
			_out5, _out6, _out7 = (_85_v25).M0(_55_globalState)
			_148_v58 = _out5
			_149_v59 = _out6
			_150_v60 = _out7
			_150_v60 = ((_dafny.Zero).Minus(_58_v2)).Cmp((func() _dafny.Int {
				if Companion_Default___.Fm0((_69_v10).F4(), (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _55_globalState) {
					return _137_cf6
				}
				return (_69_v10).F4()
			})()) == 0
			var _151_v61 _dafny.Sequence
			_ = _151_v61
			_151_v61 = _dafny.SeqOf((_69_v10).F4(), (_69_v10).F4(), _138_cf5)
			var _152_v62 _dafny.Map
			_ = _152_v62
			_152_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v61, _137_cf6)
			var _153_v64 _dafny.Array
			_ = _153_v64
			var _nwElement0_1 _dafny.CodePoint = Companion_Default___.Fm3(_68_v9, (func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter7 := _dafny.Iterate((_152_v62).Keys().Elements()); ; {
					_compr_6, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _154_v63 _dafny.Sequence
					_154_v63 = interface{}(_compr_6).(_dafny.Sequence)
					if (_152_v62).Contains(_154_v63) {
						_coll6.Add(_154_v63)
					}
				}
				return _coll6.ToSet()
			}()).Cardinality(), _55_globalState)
			_ = _nwElement0_1
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(16))
			_ = _nw13
			(_nw13).ArraySet1CodePoint(_nwElement0_1, 0)
			(_nw13).ArraySet1CodePoint(_68_v9, 1)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 2)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 3)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 4)
			(_nw13).ArraySet1CodePoint(_68_v9, 5)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 6)
			(_nw13).ArraySet1CodePoint(_dafny.CodePoint('q'), 7)
			(_nw13).ArraySet1CodePoint(_68_v9, 8)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 9)
			(_nw13).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _149_v59 {
					return (_85_v25).F5()
				}
				return _dafny.CodePoint('o')
			})(), 10)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 11)
			(_nw13).ArraySet1CodePoint(_68_v9, 12)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 13)
			(_nw13).ArraySet1CodePoint((_85_v25).F5(), 14)
			(_nw13).ArraySet1CodePoint(Companion_Default___.Fm3(_dafny.CodePoint('o'), _58_v2, _55_globalState), 15)
			_153_v64 = _nw13
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_153_v64), 0))
			_ = _index17
			(_153_v64).ArraySet1CodePoint(_dafny.CodePoint('t'), (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_153_v64), 0))
			_ = _index18
			(_153_v64).ArraySet1CodePoint(_68_v9, (_index18).Int())
		}
		if false {
			var _155_v65 _dafny.Map
			_ = _155_v65
			_155_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _138_cf5, _55_globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjmglnlgb")).Cardinality()))
			_137_cf6 = (func() _dafny.Int {
				if (func() bool {
					if _56_v0 {
						return _56_v0
					}
					return true
				})() {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_140_v52).Cardinality()))
				}
				return ((func() _dafny.Int {
					if (_155_v65).Contains(_56_v0) {
						return (_155_v65).Get(_56_v0).(_dafny.Int)
					}
					return _137_cf6
				})()).Plus((_69_v10).F4())
			})()
			var _156_v66 _dafny.Sequence
			_ = _156_v66
			_156_v66 = _dafny.SeqOf((_69_v10).F4(), (_69_v10).F4())
			var _157_v67 _dafny.MultiSet
			_ = _157_v67
			_157_v67 = _dafny.MultiSetOf(_dafny.IntOfUint32((_156_v66).Cardinality()))
			_56_v0 = !((_157_v67).Difference(_157_v67)).Equals((_157_v67).Difference(_157_v67))
			var _158_v68 _dafny.CodePoint
			_ = _158_v68
			_158_v68 = (_85_v25).F5()
			var _159_v69 _dafny.Map
			_ = _159_v69
			_159_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v68, true)
			_159_v69 = (_159_v69).Update(_68_v9, _56_v0)
			_56_v0 = (Companion_Default___.Fm0(_138_cf5, (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _55_globalState)) == (_56_v0)
			var _160_v70 D1
			_ = _160_v70
			_160_v70 = Companion_D1_.Create_DC6_()
			var _161_v71 _dafny.Map
			_ = _161_v71
			_161_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_83_v23).Select((Companion_Default___.SafeIndex((_69_v10).F4(), _dafny.IntOfUint32((_83_v23).Cardinality()))).Uint32()).(_dafny.CodePoint), _67_v8)
			var _162_v72 _dafny.Sequence
			_ = _162_v72
			_162_v72 = _dafny.SeqOf(_67_v8, _67_v8, (func() *C0 {
				if (_161_v71).Contains((_85_v25).F5()) {
					return (_161_v71).Get((_85_v25).F5()).(*C0)
				}
				return _67_v8
			})(), _67_v8, _67_v8)
			var _rhs22 D1 = _160_v70
			_ = _rhs22
			var _rhs23 _dafny.Sequence = _162_v72
			_ = _rhs23
			var _rhs24 bool = _56_v0
			_ = _rhs24
			_160_v70 = _rhs22
			_162_v72 = _rhs23
			_56_v0 = _rhs24
		} else {
			var _163_v73 *C0
			_ = _163_v73
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_66_v7, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, _57_v1), _57_v1, (_69_v10).F4())
			_163_v73 = _nw14
			var _164_v74 *C0
			_ = _164_v74
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__((_67_v8).F7(), (_67_v8).F8(), (_69_v10).F3(), _dafny.IntOfInt64(931))
			_164_v74 = _nw15
			var _165_v75 _dafny.Map
			_ = _165_v75
			_165_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v6, _56_v0)
			var _166_v76 _dafny.Map
			_ = _166_v76
			_166_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, _56_v0)
			_165_v75 = (_165_v75).Update(_65_v6, ((_166_v76).Cardinality()).Cmp((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)) != 0)
			_56_v0 = !(Companion_Default___.Fm0((_69_v10).F4(), _137_cf6, _55_globalState))
			var _167_v77 *C1
			_ = _167_v77
			var _nw16 *C1 = New_C1_()
			_ = _nw16
			_nw16.Ctor__(_68_v9, (_85_v25).F6(), (_140_v52).Select((Companion_Default___.SafeIndex(_58_v2, _dafny.IntOfUint32((_140_v52).Cardinality()))).Uint32()).(_dafny.Array), (_69_v10).F4())
			_167_v77 = _nw16
		}
		_83_v23 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-834))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_168_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))), (Companion_Default___.SafeIndex((_69_v10).F4(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-834))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_169_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})))).Cardinality()))).Uint32(), (_85_v25).F5())
	} else if _source3.Is_DC0() {
		var _170___mcc_h6 bool = _source3.Get_().(D0_DC0).Cf0
		_ = _170___mcc_h6
		var _171_cf0 bool = _170___mcc_h6
		_ = _171_cf0
		var _172_v78 *C0
		_ = _172_v78
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__((_67_v8).F7(), ((_67_v8).F8()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_cf0, (_69_v10).F3())), (_69_v10).F3(), _dafny.IntOfUint32((_83_v23).Cardinality()))
		_172_v78 = _nw17
		var _173_v79 _dafny.MultiSet
		_ = _173_v79
		_173_v79 = _dafny.MultiSetOf(_dafny.IntOfInt64(-708), _dafny.IntOfInt64(399), _58_v2, (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(153))
		var _174_v80 _dafny.Sequence
		_ = _174_v80
		_174_v80 = _dafny.SeqOf((_69_v10).F4())
		if ((_dafny.MultiSetFromSeq(_174_v80)).Intersection(_173_v79)).IsSubsetOf(_173_v79) {
			var _175_v81 _dafny.Sequence
			_ = _175_v81
			var _out8 _dafny.Sequence
			_ = _out8
			_out8 = (_85_v25).M1(_56_v0, (_69_v10).F4(), (_dafny.Zero).Minus((_69_v10).F4()), _55_globalState)
			_175_v81 = _out8
			(_55_globalState).F1 = _dafny.IntOfInt64(90)
			var _176_v82 _dafny.Sequence
			_ = _176_v82
			_176_v82 = _dafny.SeqOf(_86_v26)
			var _177_v83 D4
			_ = _177_v83
			_177_v83 = Companion_D4_.Create_DC10_(_176_v82)
			var _178_v84 _dafny.Sequence
			_ = _178_v84
			_178_v84 = _dafny.SeqOf(_177_v83, Companion_D4_.Create_DC10_(_176_v82), _177_v83)
			var _179_v85 _dafny.Map
			_ = _179_v85
			_179_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_173_v79, _dafny.Companion_Sequence_.Concatenate(_178_v84, _178_v84))
			_179_v85 = (_179_v85).Update(Companion_Default___.Fm12((Companion_Default___.Fm12(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ijxylmpsk"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dk")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ijxylmpsk")).Cardinality()))).Uint32(), (_85_v25).F5())).Cardinality()), (_69_v10).F4(), (_69_v10).F4(), _dafny.Companion_Sequence_.Update(_175_v81, (Companion_Default___.SafeIndex((_69_v10).F4(), _dafny.IntOfUint32((_175_v81).Cardinality()))).Uint32(), (_85_v25).F5()), _55_globalState)).Cardinality(), _58_v2, (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _175_v81, _55_globalState), _dafny.SeqOf(_177_v83))
			var _180_v86 D0
			_ = _180_v86
			_180_v86 = Companion_D0_.Create_DC2_(_83_v23, _58_v2, _dafny.IntOfInt64(785))
			var _rhs25 _dafny.Int = _58_v2
			_ = _rhs25
			var _rhs26 D0 = _180_v86
			_ = _rhs26
			var _lhs12 *GlobalState = _55_globalState
			_ = _lhs12
			_lhs12.F1 = _rhs25
			_180_v86 = _rhs26
			_171_cf0 = _56_v0
		} else {
			(_55_globalState).F1 = Companion_Default___.SafeModuloInt((Companion_Default___.Fm5(_56_v0, _55_globalState)).Plus((_69_v10).F4()), Companion_Default___.SafeDivisionInt((_69_v10).F4(), (_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)))
			var _181_v87 _dafny.Sequence
			_ = _181_v87
			_181_v87 = _dafny.SeqOf(_83_v23, _83_v23, _dafny.Companion_Sequence_.Update(_83_v23, (Companion_Default___.SafeIndex((_69_v10).F4(), _dafny.IntOfUint32((_83_v23).Cardinality()))).Uint32(), _68_v9))
			_83_v23 = (_181_v87).Select((Companion_Default___.SafeIndex((_69_v10).F4(), _dafny.IntOfUint32((_181_v87).Cardinality()))).Uint32()).(_dafny.Sequence)
			_57_v1 = (_69_v10).F3()
			var _182_v88 bool
			_ = _182_v88
			var _183_v89 bool
			_ = _183_v89
			var _184_v90 bool
			_ = _184_v90
			var _out9 bool
			_ = _out9
			var _out10 bool
			_ = _out10
			var _out11 bool
			_ = _out11
			_out9, _out10, _out11 = (_85_v25).M0(_55_globalState)
			_182_v88 = _out9
			_183_v89 = _out10
			_184_v90 = _out11
			var _185_v91 D1
			_ = _185_v91
			_185_v91 = Companion_D1_.Create_DC6_()
			var _186_v92 _dafny.Array
			_ = _186_v92
			var _nwElement0_2 D1 = _185_v91
			_ = _nwElement0_2
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(18))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_2, 0)
			(_nw18).ArraySet1(_185_v91, 1)
			(_nw18).ArraySet1(_185_v91, 2)
			(_nw18).ArraySet1(Companion_Default___.Fm13((_dafny.Zero).Minus((_69_v10).F4()), _55_globalState), 3)
			(_nw18).ArraySet1(_185_v91, 4)
			(_nw18).ArraySet1(_185_v91, 5)
			(_nw18).ArraySet1((func() D1 {
				if _183_v89 {
					return _185_v91
				}
				return _185_v91
			})(), 6)
			(_nw18).ArraySet1(_185_v91, 7)
			(_nw18).ArraySet1(_185_v91, 8)
			(_nw18).ArraySet1(Companion_D1_.Create_DC6_(), 9)
			(_nw18).ArraySet1(_185_v91, 10)
			(_nw18).ArraySet1(_185_v91, 11)
			(_nw18).ArraySet1(_185_v91, 12)
			(_nw18).ArraySet1(Companion_Default___.Fm13((_76_v16).Cardinality(), _55_globalState), 13)
			(_nw18).ArraySet1(_185_v91, 14)
			(_nw18).ArraySet1(_185_v91, 15)
			(_nw18).ArraySet1(_185_v91, 16)
			(_nw18).ArraySet1(_185_v91, 17)
			_186_v92 = _nw18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_186_v92), 0))
			_ = _index19
			(_186_v92).ArraySet1(_185_v91, (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_186_v92), 0))
			_ = _index20
			(_186_v92).ArraySet1(_185_v91, (_index20).Int())
		}
		var _187_v93 *C0
		_ = _187_v93
		var _nw19 *C0 = New_C0_()
		_ = _nw19
		_nw19.Ctor__((_67_v8).F7(), ((_172_v78).F8()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), _58_v2, _55_globalState), (_69_v10).F3())), _57_v1, (_dafny.IntOfUint32((_83_v23).Cardinality())).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-931))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_188_v25 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_189_i9 _dafny.Int) _dafny.CodePoint {
				return (_188_v25).F5()
			}
		})(_85_v25)))).Cardinality()))))
		_187_v93 = _nw19
		var _190_v94 _dafny.Map
		_ = _190_v94
		_190_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_85_v25).F5(), _171_cf0)
		var _source4 D4 = Companion_Default___.Fm14((func() _dafny.Map {
			if _56_v0 {
				return _190_v94
			}
			return _190_v94
		})(), ((_69_v10).F4()).Times((_69_v10).F4()), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(197), (_dafny.Zero).Minus((_69_v10).F4())), _78_v18, _55_globalState)
		_ = _source4
		if _source4.Is_DC11() {
			var _191___mcc_h8 _dafny.CodePoint = _source4.Get_().(D4_DC11).Cf17
			_ = _191___mcc_h8
			var _192_cf17 _dafny.CodePoint = _191___mcc_h8
			_ = _192_cf17
			var _193_v95 _dafny.Array
			_ = _193_v95
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_2
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_194_v25 *C1) func(_dafny.Int) _dafny.Sequence {
					return func(_195_i10 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(813))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}((func(_196_v25 *C1) func(_dafny.Int) _dafny.CodePoint {
							return func(_197_i11 _dafny.Int) _dafny.CodePoint {
								return (_196_v25).F5()
							}
						})(_194_v25)))
					}
				})(_85_v25)
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
			_193_v95 = _nw20
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_193_v95), 0))
			_ = _index21
			(_193_v95).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_198_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_199_i12 _dafny.Int) _dafny.CodePoint {
					return (_198_v9)
				}
			})(_68_v9))), _83_v23), (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_193_v95), 0))
			_ = _index22
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index23
			var _rhs27 _dafny.Int = _dafny.IntOfInt64(849)
			_ = _rhs27
			var _rhs28 _dafny.Sequence = _83_v23
			_ = _rhs28
			var _rhs29 _dafny.Int = ((_dafny.Zero).Minus((_69_v10).F4())).Times(Companion_Default___.SafeDivisionInt((_69_v10).F4(), (_69_v10).F4()))
			_ = _rhs29
			var _rhs30 _dafny.Int = ((_dafny.Zero).Minus(((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)).Plus((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int)))).Minus((_69_v10).F4())
			_ = _rhs30
			var _lhs13 *GlobalState = _55_globalState
			_ = _lhs13
			var _lhs14 _dafny.Array = _193_v95
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_193_v95), 0))
			_ = _lhs15
			var _lhs16 *GlobalState = _55_globalState
			_ = _lhs16
			var _lhs17 _dafny.Array = _57_v1
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _lhs18
			_lhs13.F1 = _rhs27
			(_lhs14).ArraySet1(_rhs28, (_lhs15).Int())
			_lhs16.F1 = _rhs29
			(_lhs17).ArraySet1(_rhs30, (_lhs18).Int())
			var _200_v96 _dafny.Array
			_ = _200_v96
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw21
			_200_v96 = _nw21
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_200_v96), 0))
			_ = _index24
			(_200_v96).ArraySet1((_85_v25).F6(), (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_200_v96), 0))
			_ = _index25
			var _nwElement0_3 bool = false
			_ = _nwElement0_3
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(9))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_3, 0)
			(_nw22).ArraySet1(_171_cf0, 1)
			(_nw22).ArraySet1(_56_v0, 2)
			(_nw22).ArraySet1(_56_v0, 3)
			(_nw22).ArraySet1(false, 4)
			(_nw22).ArraySet1(_56_v0, 5)
			(_nw22).ArraySet1(_171_cf0, 6)
			(_nw22).ArraySet1(_171_cf0, 7)
			(_nw22).ArraySet1(_56_v0, 8)
			(_200_v96).ArraySet1(_nw22, (_index25).Int())
			var _pat_let_tv2 = _171_cf0
			_ = _pat_let_tv2
			var _201_v97 *C0
			_ = _201_v97
			var _nw23 *C0 = New_C0_()
			_ = _nw23
			_nw23.Ctor__(func(_pat_let4_0 D0) D0 {
				return func(_202_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let5_0 bool) D0 {
						return func(_203_dt__update_hcf0_h1 bool) D0 {
							return Companion_D0_.Create_DC0_(_203_dt__update_hcf0_h1)
						}(_pat_let5_0)
					}(_pat_let_tv2)
				}(_pat_let4_0)
			}((_187_v93).F7()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, (_69_v10).F3()), (_69_v10).F3(), _dafny.IntOfUint32((_174_v80).Cardinality()))
			_201_v97 = _nw23
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index26
			(_57_v1).ArraySet1((_69_v10).F4(), (_index26).Int())
		} else if _source4.Is_DC12() {
			var _204___mcc_h9 *C0 = _source4.Get_().(D4_DC12).Cf18
			_ = _204___mcc_h9
			var _205___mcc_h10 bool = _source4.Get_().(D4_DC12).Cf19
			_ = _205___mcc_h10
			var _206_cf19 bool = _205___mcc_h10
			_ = _206_cf19
			var _207_cf18 *C0 = _204___mcc_h9
			_ = _207_cf18
			var _208_v98 _dafny.Set
			_ = _208_v98
			_208_v98 = _dafny.SetOf(_171_cf0)
			(_55_globalState).F1 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm5(!(Companion_Default___.Fm0((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), (_69_v10).F4(), _55_globalState)), _55_globalState), (_208_v98).Cardinality())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index27
			(_57_v1).ArraySet1(((_69_v10).F4()).Times(_58_v2), (_index27).Int())
			var _209_v99 _dafny.Set
			_ = _209_v99
			_209_v99 = _dafny.SetOf(_83_v23, _dafny.Companion_Sequence_.Concatenate(_83_v23, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(696))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_210_v23 _dafny.Sequence, _211_v2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_212_i13 _dafny.Int) _dafny.CodePoint {
					return (_210_v23).Select((Companion_Default___.SafeIndex(_211_v2, _dafny.IntOfUint32((_210_v23).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_83_v23, _58_v2)))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(974))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_213_v25 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_214_i14 _dafny.Int) _dafny.CodePoint {
					return (_213_v25).F5()
				}
			})(_85_v25))), _83_v23, _83_v23)
			_209_v99 = (_dafny.SetOf(_83_v23)).Intersection(_209_v99)
			_173_v79 = _173_v79
		} else {
			var _215___mcc_h11 _dafny.Sequence = _source4.Get_().(D4_DC10).Cf16
			_ = _215___mcc_h11
			var _216_cf16 _dafny.Sequence = _215___mcc_h11
			_ = _216_cf16
			var _217_v100 D0
			_ = _217_v100
			_217_v100 = Companion_D0_.Create_DC2_(_83_v23, (_69_v10).F4(), (_69_v10).F4())
			(_55_globalState).F1 = (_217_v100).Dtor_cf5()
			_187_v93 = _67_v8
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
			_ = _index28
			(_57_v1).ArraySet1(_58_v2, (_index28).Int())
			var _218_v101 *C1
			_ = _218_v101
			var _nw24 *C1 = New_C1_()
			_ = _nw24
			_nw24.Ctor__((_85_v25).F5(), _84_v24, (_69_v10).F3(), Companion_Default___.SafeDivisionInt((_69_v10).F4(), _58_v2))
			_218_v101 = _nw24
		}
	} else {
		var _219___mcc_h7 D0 = _source3.Get_().(D0_DC3).Cf7
		_ = _219___mcc_h7
		var _220_cf7 D0 = _219___mcc_h7
		_ = _220_cf7
		(_55_globalState).F1 = (_69_v10).F4()
		var _221_v102 bool
		_ = _221_v102
		var _222_v103 bool
		_ = _222_v103
		var _223_v104 bool
		_ = _223_v104
		var _out12 bool
		_ = _out12
		var _out13 bool
		_ = _out13
		var _out14 bool
		_ = _out14
		_out12, _out13, _out14 = (_85_v25).M0(_55_globalState)
		_221_v102 = _out12
		_222_v103 = _out13
		_223_v104 = _out14
		var _224_v105 _dafny.Sequence
		_ = _224_v105
		_224_v105 = _dafny.SeqOf(_dafny.IntOfUint32((_83_v23).Cardinality()), (_69_v10).F4())
		_56_v0 = Companion_Default___.Fm0((_dafny.IntOfInt64(918)).Minus(Companion_Default___.Fm5(Companion_Default___.Fm0(_dafny.IntOfUint32((_224_v105).Cardinality()), _58_v2, _55_globalState), _55_globalState)), (_58_v2).Times(_dafny.IntOfInt64(-630)), _55_globalState)
		var _225_v106 _dafny.Sequence
		_ = _225_v106
		_225_v106 = _dafny.SeqOf(_69_v10, _69_v10)
		var _rhs31 bool = !(Companion_Default___.Fm0((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), (_69_v10).F4(), _55_globalState))
		_ = _rhs31
		var _rhs32 _dafny.Int = (_dafny.Zero).Minus(_58_v2)
		_ = _rhs32
		var _rhs33 _dafny.Sequence = _83_v23
		_ = _rhs33
		var _rhs34 T0 = (_225_v106).Select((Companion_Default___.SafeIndex(_58_v2, _dafny.IntOfUint32((_225_v106).Cardinality()))).Uint32()).(T0)
		_ = _rhs34
		var _rhs35 bool = (Companion_D1_.Create_DC5_(_223_v104, _223_v104, (_69_v10).F4())).Dtor_cf10()
		_ = _rhs35
		var _lhs19 *GlobalState = _55_globalState
		_ = _lhs19
		_56_v0 = _rhs31
		_lhs19.F1 = _rhs32
		_83_v23 = _rhs33
		_69_v10 = _rhs34
		_222_v103 = _rhs35
	}
	var _226_v107 _dafny.Map
	_ = _226_v107
	_226_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v10, _56_v0)
	var _227_v108 _dafny.Sequence
	_ = _227_v108
	_227_v108 = _dafny.SeqOf((_226_v107).Cardinality())
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))
	_ = _index29
	(_57_v1).ArraySet1(Companion_Default___.SafeDivisionInt((_57_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_57_v1), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt(_58_v2, (_dafny.Zero).Minus((_227_v108).Select((Companion_Default___.SafeIndex(_58_v2, _dafny.IntOfUint32((_227_v108).Cardinality()))).Uint32()).(_dafny.Int)))), (_index29).Int())
	_dafny.Print((_55_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_55_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_55_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_56_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v1).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_58_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_v3).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v4).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_64_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_65_v6).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v7).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_67_v8).F7()).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_67_v8).F8()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_69_v10).F3()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_69_v10).F3()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_69_v10).F3()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_69_v10).F3()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v10).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_70_v11).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), true), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v13).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_v15).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(891))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_77_v17, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_v18).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_v18).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_v18).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v19).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v21).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_82_v22).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_82_v22).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_82_v22).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_82_v22).Dtor_cf8()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_83_v23.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_84_v24).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_84_v24).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v25).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_v25).F6()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_v25).F6()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_86_v26).Dtor_cf7()).Dtor_cf4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_86_v26).Dtor_cf7()).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_86_v26).Dtor_cf7()).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_226_v107).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_227_v108, _dafny.SeqOf(_dafny.One)))
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
	Cf2 _dafny.Int
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.Sequence
	Cf5 _dafny.Int
	Cf6 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.Sequence, Cf5 _dafny.Int, Cf6 _dafny.Int) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6}}
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

type D0_DC3 struct {
	Cf7 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf7 D0) D0 {
	return D0{D0_DC3{Cf7}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf7() D0 {
	return _this.Get_().(D0_DC3).Cf7
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + data.Cf4.VerbatimString(true) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4) && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf9  bool
	Cf10 bool
	Cf11 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 bool, Cf10 bool, Cf11 _dafny.Int) D1 {
	return D1{D1_DC5{Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_() D1 {
	return D1{D1_DC6{}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC4 struct {
	Cf8 _dafny.Array
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.Array) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false, false, _dafny.Zero)
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC6:
		{
			_, ok := other.Get_().(D1_DC6)
			return ok
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf8 == data2.Cf8
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
	Cf13 _dafny.Sequence
	Cf14 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Sequence, Cf14 _dafny.Int) D2 {
	return D2{D2_DC8{Cf13, Cf14}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf12 _dafny.Sequence
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.Sequence) D2 {
	return D2{D2_DC7{Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D2) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + data.Cf13.VerbatimString(true) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13.Equals(data2.Cf13) && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf15 _dafny.CodePoint
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 _dafny.CodePoint) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D3) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf15 == data2.Cf15
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
	Cf17 _dafny.CodePoint
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 _dafny.CodePoint) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf18 *C0
	Cf19 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf18 *C0, Cf19 bool) D4 {
	return D4{D4_DC12{Cf18, Cf19}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC10 struct {
	Cf16 _dafny.Sequence
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 _dafny.Sequence) D4 {
	return D4{D4_DC10{Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.CodePoint('D'))
}

func (_this D4) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) Dtor_cf18() *C0 {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) Dtor_cf19() bool {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf17 == data2.Cf17
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
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

type D5_DC14 struct {
	Cf21 _dafny.Int
	Cf22 _dafny.Int
	Cf23 bool
	Cf24 _dafny.Int
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf21 _dafny.Int, Cf22 _dafny.Int, Cf23 bool, Cf24 _dafny.Int) D5 {
	return D5{D5_DC14{Cf21, Cf22, Cf23, Cf24}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf20 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf20 _dafny.Map) D5 {
	return D5{D5_DC13{Cf20}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC15 struct {
	Cf25 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf25 D5) D5 {
	return D5{D5_DC15{Cf25}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.Zero, _dafny.Zero, false, _dafny.Zero)
}

func (_this D5) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) Dtor_cf23() bool {
	return _this.Get_().(D5_DC14).Cf23
}

func (_this D5) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf24
}

func (_this D5) Dtor_cf20() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf25() D5 {
	return _this.Get_().(D5_DC15).Cf25
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf20.Equals(data2.Cf20)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

// Definition of trait T0
type T0 interface {
	String() string
	F3() _dafny.Array
	F4() _dafny.Int
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
	F1  _dafny.Int
	_f0 _dafny.Int
	_f2 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this._f0 = _dafny.Zero
	_this._f2 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f3 _dafny.Array
	_f4 _dafny.Int
	_f7 D0
	_f8 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = _dafny.Zero
	_this._f7 = Companion_D0_.Default()
	_this._f8 = _dafny.EmptyMap
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

func (_this *C0) F3() _dafny.Array {
	return _this._f3
}
func (_this *C0) F4() _dafny.Int {
	return _this._f4
}
func (_this *C0) Ctor__(f7 D0, f8 _dafny.Map, f3 _dafny.Array, f4 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C0) F7() D0 {
	{
		return _this._f7
	}
}
func (_this *C0) F8() _dafny.Map {
	{
		return _this._f8
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f3 _dafny.Array
	_f4 _dafny.Int
	_f5 _dafny.CodePoint
	_f6 _dafny.Array
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C1) F3() _dafny.Array {
	return _this._f3
}
func (_this *C1) F4() _dafny.Int {
	return _this._f4
}
func (_this *C1) Ctor__(f5 _dafny.CodePoint, f6 _dafny.Array, f3 _dafny.Array, f4 _dafny.Int) {
	{
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C1) M0(globalState *GlobalState) (bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _228_v0 _dafny.Sequence
		_ = _228_v0
		_228_v0 = _dafny.UnicodeSeqOfUtf8Bytes("lwcdnhtm")
		var _229_v1 _dafny.Set
		_ = _229_v1
		_229_v1 = _dafny.SetOf(_228_v0)
		var _230_v2 bool
		_ = _230_v2
		_230_v2 = false
		var _231_v3 D0
		_ = _231_v3
		_231_v3 = Companion_D0_.Create_DC0_(_230_v2)
		var _232_v4 _dafny.Map
		_ = _232_v4
		_232_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v1, (_231_v3).Dtor_cf0())
		_232_v4 = (_232_v4).Update((_229_v1).Union(_229_v1), _230_v2)
		var _233_v5 _dafny.Sequence
		_ = _233_v5
		_233_v5 = _dafny.SeqOf((_this).F4(), _dafny.IntOfInt64(930), (_this).F4(), (_this).F4(), (_this).F4())
		var _234_v6 _dafny.Map
		_ = _234_v6
		_234_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v2, (_dafny.Zero).Minus((_this).F4()))
		var _235_v7 _dafny.MultiSet
		_ = _235_v7
		_235_v7 = _dafny.MultiSetOf((_this).F4(), (_234_v6).Cardinality(), _dafny.IntOfUint32((_228_v0).Cardinality()), (_this).F4())
		(globalState).F1 = (func() _dafny.Int {
			if !(!(_230_v2)) {
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_233_v5).Cardinality()))
			}
			return Companion_Default___.SafeModuloInt((_this).F4(), (_235_v7).Cardinality())
		})()
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))
		_ = _index30
		((_this).F6()).ArraySet1(_230_v2, (_index30).Int())
		var _236_v8 _dafny.Map
		_ = _236_v8
		_236_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v2, !(_230_v2))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))
		_ = _index31
		((_this).F6()).ArraySet1(!_dafny.Companion_Sequence_.Contains(_233_v5, ((func() _dafny.Map {
			if _230_v2 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v2, _230_v2)
			}
			return _236_v8
		})()).Cardinality()), (_index31).Int())
		var _237_v9 D0
		_ = _237_v9
		_237_v9 = Companion_D0_.Create_DC2_(_228_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _230_v2))).Cardinality(), _dafny.IntOfUint32((_228_v0).Cardinality()))
		var _238_v10 _dafny.Sequence
		_ = _238_v10
		_238_v10 = _dafny.SeqOf((_237_v9).Dtor_cf4(), _228_v0)
		var _hi0 _dafny.Int = _dafny.IntOfUint32(((_238_v10).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_238_v10).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
		_ = _hi0
		for _239_i0 := _dafny.IntOfInt64(-768); _239_i0.Cmp(_hi0) < 0; _239_i0 = _239_i0.Plus(_dafny.One) {
			var _240_v11 _dafny.Sequence
			_ = _240_v11
			_240_v11 = _dafny.SeqOf(_230_v2)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen(((_this).F3()), 0))
			_ = _index32
			((_this).F3()).ArraySet1(((_235_v7).Cardinality()).Plus(_dafny.IntOfUint32((_240_v11).Cardinality())), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen(((_this).F3()), 0))
			_ = _index33
			var _rhs36 _dafny.Int = (_237_v9).Dtor_cf6()
			_ = _rhs36
			var _rhs37 bool = (Companion_Default___.SafeModuloInt((_this).F4(), (_this).F4())).Cmp(_dafny.IntOfInt64(288)) == 0
			_ = _rhs37
			var _rhs38 bool = _230_v2
			_ = _rhs38
			var _rhs39 _dafny.Int = _dafny.IntOfUint32((_228_v0).Cardinality())
			_ = _rhs39
			var _lhs20 *GlobalState = globalState
			_ = _lhs20
			var _lhs21 _dafny.Array = (_this).F3()
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen(((_this).F3()), 0))
			_ = _lhs22
			_lhs20.F1 = _rhs36
			r1 = _rhs37
			r1 = _rhs38
			(_lhs21).ArraySet1(_rhs39, (_lhs22).Int())
			(globalState).F1 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(-468)).Minus(((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int)), (_this).F4())
			var _241_v12 _dafny.Map
			_ = _241_v12
			_241_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _230_v2, globalState), (Companion_Default___.SafeIndex(((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm1(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _230_v2, globalState)).Cardinality()))).Uint32(), (_this).F5()))
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))
			_ = _index34
			((_this).F6()).ArraySet1((func() bool {
				if (_236_v8).Contains(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)) {
					return (_236_v8).Get(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)).(bool)
				}
				return !((_241_v12).Contains(_230_v2))
			})(), (_index34).Int())
			var _242_v13 _dafny.Array
			_ = _242_v13
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw25
			_242_v13 = _nw25
			var _243_v14 _dafny.Array
			_ = _243_v14
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_3
			var _nw26 _dafny.Array
			_ = _nw26
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw26 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_244_v3 D0) func(_dafny.Int) _dafny.Sequence {
					return func(_245_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_244_v3).Dtor_cf0())
					}
				})(_231_v3)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw26 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw26).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw26).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_243_v14 = _nw26
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_242_v13), 0))
			_ = _index35
			(_242_v13).ArraySet1(_243_v14, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_242_v13), 0))
			_ = _index36
			(_242_v13).ArraySet1(_243_v14, (_index36).Int())
		}
		if !(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)) {
			r2 = false
			var _246_v15 _dafny.Map
			_ = _246_v15
			_246_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F4()).Plus((_this).F4()), ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool))
			_246_v15 = (_246_v15).Update((_this).F4(), ((_this).F4()).Cmp((_this).F4()) <= 0)
			var _247_v16 _dafny.Sequence
			_ = _247_v16
			_247_v16 = _dafny.SeqOf(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _230_v2)
			var _248_v17 _dafny.Sequence
			_ = _248_v17
			_248_v17 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_247_v16).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_234_v6).Contains(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)) {
						return (_234_v6).Get(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)).(_dafny.Int)
					}
					return (_this).F4()
				})(), _dafny.IntOfUint32((_247_v16).Cardinality()))).Uint32()).(bool) {
					return _247_v16
				}
				return _dafny.SeqOf(false)
			})())
			_248_v17 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_248_v17, _248_v17), _dafny.Companion_Sequence_.Update(_248_v17, (Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_248_v17).Cardinality()))).Uint32(), _247_v16))
			(globalState).F1 = (_this).F4()
			if _230_v2 {
				(globalState).F1 = (_this).F4()
				var _249_v18 _dafny.Map
				_ = _249_v18
				_249_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _246_v15)
				_249_v18 = (_249_v18).Merge((_249_v18).Merge(Companion_Default___.Fm2((_this).F4(), !(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)), _230_v2, globalState)))
				_235_v7 = _235_v7
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index37
				((_this).F3()).ArraySet1((_this).F4(), (_index37).Int())
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index38
				((_this).F3()).ArraySet1((_this).F4(), (_index38).Int())
				(globalState).F1 = ((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int)
			} else {
				var _250_v19 _dafny.Set
				_ = _250_v19
				_250_v19 = _dafny.SetOf(_230_v2)
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))
				_ = _index39
				((_this).F6()).ArraySet1((_dafny.SetOf(false)).IsProperSubsetOf(_250_v19), (_index39).Int())
				var _251_v20 _dafny.Array
				_ = _251_v20
				var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
				_ = _nw27
				_251_v20 = _nw27
				var _252_v21 _dafny.MultiSet
				_ = _252_v21
				_252_v21 = _dafny.MultiSetOf(Companion_D0_.Create_DC0_(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)))
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_251_v20), 0))
				_ = _index40
				(_251_v20).ArraySet1(_252_v21, (_index40).Int())
				var _253_v22 _dafny.Map
				_ = _253_v22
				_253_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _246_v15)
				var _254_v23 _dafny.MultiSet
				_ = _254_v23
				_254_v23 = _dafny.MultiSetOf(Companion_Default___.Fm3((_this).F5(), (_this).F4(), globalState), Companion_Default___.Fm3((_this).F5(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_255_i2 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				}))).Cardinality()), globalState), (_this).F5(), (_this).F5())
				var _256_v24 _dafny.Map
				_ = _256_v24
				_256_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_253_v22).Equals(_253_v22), _254_v23)
				var _257_v25 _dafny.Map
				_ = _257_v25
				_257_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_251_v20), 0))
				_ = _index41
				var _rhs40 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.MultiSet {
					if (_256_v24).Contains(!(_230_v2)) {
						return (_256_v24).Get(!(_230_v2)).(_dafny.MultiSet)
					}
					return _254_v23
				})()).Cardinality())
				_ = _rhs40
				var _rhs41 _dafny.Int = (func() _dafny.Int {
					if (_257_v25).Contains((_this).F4()) {
						return (_257_v25).Get((_this).F4()).(_dafny.Int)
					}
					return Companion_Default___.SafeDivisionInt((_this).F4(), _dafny.IntOfInt64(-768))
				})()
				_ = _rhs41
				var _rhs42 _dafny.MultiSet = _252_v21
				_ = _rhs42
				var _lhs23 *GlobalState = globalState
				_ = _lhs23
				var _lhs24 *GlobalState = globalState
				_ = _lhs24
				var _lhs25 _dafny.Array = _251_v20
				_ = _lhs25
				var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_251_v20), 0))
				_ = _lhs26
				_lhs23.F1 = _rhs40
				_lhs24.F1 = _rhs41
				(_lhs25).ArraySet1(_rhs42, (_lhs26).Int())
				(globalState).F1 = _dafny.IntOfUint32((_228_v0).Cardinality())
				var _258_v26 _dafny.Array
				_ = _258_v26
				var _nw28 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw28
				_258_v26 = _nw28
				_258_v26 = _258_v26
				var _259_v27 _dafny.Map
				_ = _259_v27
				_259_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), (_this).F3())
				var _260_v28 *C0
				_ = _260_v28
				var _nw29 *C0 = New_C0_()
				_ = _nw29
				_nw29.Ctor__(Companion_D0_.Create_DC0_(true), (_259_v27).Merge(_259_v27), (_this).F3(), (_this).F4())
				_260_v28 = _nw29
			}
		} else {
			(globalState).F1 = (Companion_Default___.SafeDivisionInt((_this).F4(), _dafny.IntOfUint32((_dafny.SeqOf(_230_v2)).Cardinality()))).Minus((_dafny.Zero).Minus((_this).F4()))
			_228_v0 = _dafny.UnicodeSeqOfUtf8Bytes("blty")
			r1 = (true) && (!((_231_v3).Dtor_cf0()))
			var _261_v29 _dafny.Map
			_ = _261_v29
			_261_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), (_this).F3())
			var _262_v30 *C0
			_ = _262_v30
			var _nw30 *C0 = New_C0_()
			_ = _nw30
			_nw30.Ctor__(_231_v3, _261_v29, (_this).F3(), (_this).F4())
			_262_v30 = _nw30
			r0 = ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)
		}
		var _hi1 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hki")).Cardinality())
		_ = _hi1
		for _263_i3 := (Companion_Default___.Fm4((func() bool {
			if (_236_v8).Contains(_230_v2) {
				return (_236_v8).Get(_230_v2).(bool)
			}
			return _230_v2
		})(), _dafny.IntOfUint32((_dafny.SeqOf(_230_v2, _230_v2, _230_v2)).Cardinality()), (_this).F4(), globalState)).Cardinality(); _263_i3.Cmp(_hi1) < 0; _263_i3 = _263_i3.Plus(_dafny.One) {
			var _264_v31 _dafny.Array
			_ = _264_v31
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw31
			_264_v31 = _nw31
			var _265_v32 _dafny.Sequence
			_ = _265_v32
			_265_v32 = _dafny.SeqOf(_264_v31)
			var _rhs43 bool = !(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool))
			_ = _rhs43
			var _rhs44 _dafny.Sequence = _228_v0
			_ = _rhs44
			var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_265_v32, _265_v32)
			_ = _rhs45
			r0 = _rhs43
			_228_v0 = _rhs44
			_265_v32 = _rhs45
			var _266_v33 _dafny.Set
			_ = _266_v33
			_266_v33 = _dafny.SetOf(_264_v31, _264_v31, _264_v31, _264_v31, _264_v31)
			var _267_v34 _dafny.Array
			_ = _267_v34
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_4
			var _nw32 _dafny.Array
			_ = _nw32
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw32 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_268_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_269_i4 _dafny.Int) _dafny.Sequence {
						return _268_v0
					}
				})(_228_v0)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw32 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw32).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw32).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_267_v34 = _nw32
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_267_v34), 0))
			_ = _index42
			(_267_v34).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vcbn"), (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_267_v34), 0))
			_ = _index43
			var _rhs46 bool = _230_v2
			_ = _rhs46
			var _rhs47 _dafny.Set = _266_v33
			_ = _rhs47
			var _rhs48 bool = ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)
			_ = _rhs48
			var _rhs49 _dafny.Sequence = _228_v0
			_ = _rhs49
			var _lhs27 _dafny.Array = _267_v34
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_267_v34), 0))
			_ = _lhs28
			r0 = _rhs46
			_266_v33 = _rhs47
			r1 = _rhs48
			(_lhs27).ArraySet1(_rhs49, (_lhs28).Int())
			r0 = _230_v2
			var _source5 D0 = _237_v9
			_ = _source5
			if _source5.Is_DC1() {
				var _270___mcc_h0 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
				_ = _270___mcc_h0
				var _271___mcc_h1 _dafny.Int = _source5.Get_().(D0_DC1).Cf2
				_ = _271___mcc_h1
				var _272___mcc_h2 _dafny.Int = _source5.Get_().(D0_DC1).Cf3
				_ = _272___mcc_h2
				var _273_cf3 _dafny.Int = _272___mcc_h2
				_ = _273_cf3
				var _274_cf2 _dafny.Int = _271___mcc_h1
				_ = _274_cf2
				var _275_cf1 _dafny.Int = _270___mcc_h0
				_ = _275_cf1
				var _276_v35 _dafny.Array
				_ = _276_v35
				var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw33
				_276_v35 = _nw33
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_276_v35), 0))
				_ = _index44
				(_276_v35).ArraySet1(_264_v31, (_index44).Int())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))
				_ = _index45
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_276_v35), 0))
				_ = _index46
				var _rhs50 bool = _dafny.Companion_Sequence_.Contains(_228_v0, (_this).F5())
				_ = _rhs50
				var _rhs51 _dafny.Array = (_265_v32).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_265_v32).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs51
				var _rhs52 bool = _230_v2
				_ = _rhs52
				var _lhs29 _dafny.Array = (_this).F6()
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))
				_ = _lhs30
				var _lhs31 _dafny.Array = _276_v35
				_ = _lhs31
				var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_276_v35), 0))
				_ = _lhs32
				(_lhs29).ArraySet1(_rhs50, (_lhs30).Int())
				(_lhs31).ArraySet1(_rhs51, (_lhs32).Int())
				r2 = _rhs52
				var _277_v36 _dafny.Set
				_ = _277_v36
				_277_v36 = _dafny.SetOf((_this).F4(), _274_cf2)
				_277_v36 = (_277_v36).Union(_277_v36)
				var _278_v37 _dafny.Sequence
				_ = _278_v37
				_278_v37 = _dafny.SeqOf(_231_v3, _231_v3, Companion_D0_.Create_DC0_(_230_v2))
				var _279_v38 _dafny.CodePoint
				_ = _279_v38
				_279_v38 = _dafny.CodePoint('h')
				var _280_v39 D1
				_ = _280_v39
				_280_v39 = Companion_D1_.Create_DC4_((_this).F3())
				var _281_v40 _dafny.Sequence
				_ = _281_v40
				_281_v40 = _dafny.SeqOf((_this).F3())
				var _282_v41 _dafny.Array
				_ = _282_v41
				var _nwElement0_4 _dafny.Array = (_this).F3()
				_ = _nwElement0_4
				var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
				_ = _nw34
				(_nw34).ArraySet1(_nwElement0_4, 0)
				(_nw34).ArraySet1((_280_v39).Dtor_cf8(), 1)
				(_nw34).ArraySet1((_this).F3(), 2)
				(_nw34).ArraySet1((_this).F3(), 3)
				(_nw34).ArraySet1((_this).F3(), 4)
				(_nw34).ArraySet1((_this).F3(), 5)
				(_nw34).ArraySet1((_this).F3(), 6)
				(_nw34).ArraySet1((_this).F3(), 7)
				(_nw34).ArraySet1((_this).F3(), 8)
				(_nw34).ArraySet1((_this).F3(), 9)
				(_nw34).ArraySet1((_this).F3(), 10)
				(_nw34).ArraySet1((_this).F3(), 11)
				(_nw34).ArraySet1((_this).F3(), 12)
				(_nw34).ArraySet1((_this).F3(), 13)
				(_nw34).ArraySet1((_this).F3(), 14)
				(_nw34).ArraySet1((_this).F3(), 15)
				(_nw34).ArraySet1((_this).F3(), 16)
				(_nw34).ArraySet1((_this).F3(), 17)
				(_nw34).ArraySet1((_this).F3(), 18)
				(_nw34).ArraySet1((_this).F3(), 19)
				(_nw34).ArraySet1((_this).F3(), 20)
				(_nw34).ArraySet1((_this).F3(), 21)
				(_nw34).ArraySet1((_281_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.IntOfUint32((_281_v40).Cardinality()))).Uint32()).(_dafny.Array), 22)
				(_nw34).ArraySet1((_this).F3(), 23)
				(_nw34).ArraySet1((_this).F3(), 24)
				(_nw34).ArraySet1((_this).F3(), 25)
				(_nw34).ArraySet1((_281_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_233_v5).Cardinality()), _dafny.IntOfUint32((_281_v40).Cardinality()))).Uint32()).(_dafny.Array), 26)
				(_nw34).ArraySet1((_this).F3(), 27)
				_282_v41 = _nw34
				var _283_v42 _dafny.Sequence
				_ = _283_v42
				_283_v42 = _dafny.SeqOf(_282_v41, _282_v41, _282_v41, _282_v41, _282_v41)
				var _284_v43 _dafny.Array
				_ = _284_v43
				var _nwElement0_5 _dafny.Array = _282_v41
				_ = _nwElement0_5
				var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(18))
				_ = _nw35
				(_nw35).ArraySet1(_nwElement0_5, 0)
				(_nw35).ArraySet1(_282_v41, 1)
				(_nw35).ArraySet1(_282_v41, 2)
				(_nw35).ArraySet1(_282_v41, 3)
				(_nw35).ArraySet1(_282_v41, 4)
				(_nw35).ArraySet1((func() _dafny.Array {
					if Companion_Default___.Fm0((_this).F4(), _274_cf2, globalState) {
						return _282_v41
					}
					return _282_v41
				})(), 5)
				(_nw35).ArraySet1(_282_v41, 6)
				(_nw35).ArraySet1(_282_v41, 7)
				(_nw35).ArraySet1((_283_v42).Select((Companion_Default___.SafeIndex(_274_cf2, _dafny.IntOfUint32((_283_v42).Cardinality()))).Uint32()).(_dafny.Array), 8)
				(_nw35).ArraySet1(_282_v41, 9)
				(_nw35).ArraySet1(_282_v41, 10)
				(_nw35).ArraySet1(_282_v41, 11)
				(_nw35).ArraySet1(_282_v41, 12)
				(_nw35).ArraySet1(_282_v41, 13)
				(_nw35).ArraySet1((func() _dafny.Array {
					if ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool) {
						return _282_v41
					}
					return _282_v41
				})(), 14)
				(_nw35).ArraySet1(_282_v41, 15)
				(_nw35).ArraySet1(_282_v41, 16)
				(_nw35).ArraySet1(_282_v41, 17)
				_284_v43 = _nw35
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_284_v43), 0))
				_ = _index47
				(_284_v43).ArraySet1(_282_v41, (_index47).Int())
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_284_v43), 0))
				_ = _index48
				var _rhs53 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(Companion_D0_.Create_DC0_(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)), func(_pat_let6_0 D0) D0 {
					return func(_285_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let7_0 bool) D0 {
							return func(_286_dt__update_hcf0_h0 bool) D0 {
								return Companion_D0_.Create_DC0_(_286_dt__update_hcf0_h0)
							}(_pat_let7_0)
						}(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool))
					}(_pat_let6_0)
				}(Companion_D0_.Create_DC0_(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)))), (Companion_Default___.SafeIndex(_275_cf1, _dafny.IntOfUint32((_dafny.SeqOf(Companion_D0_.Create_DC0_(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)), func(_pat_let8_0 D0) D0 {
					return func(_287_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let9_0 bool) D0 {
							return func(_288_dt__update_hcf0_h1 bool) D0 {
								return Companion_D0_.Create_DC0_(_288_dt__update_hcf0_h1)
							}(_pat_let9_0)
						}(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool))
					}(_pat_let8_0)
				}(Companion_D0_.Create_DC0_(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool))))).Cardinality()))).Uint32(), _231_v3)
				_ = _rhs53
				var _rhs54 _dafny.CodePoint = _279_v38
				_ = _rhs54
				var _rhs55 _dafny.Array = _282_v41
				_ = _rhs55
				var _lhs33 _dafny.Array = _284_v43
				_ = _lhs33
				var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_284_v43), 0))
				_ = _lhs34
				_278_v37 = _rhs53
				_279_v38 = _rhs54
				(_lhs33).ArraySet1(_rhs55, (_lhs34).Int())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index49
				((_this).F3()).ArraySet1(_263_i3, (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index50
				((_this).F3()).ArraySet1((_dafny.Zero).Minus(_275_cf1), (_index50).Int())
			} else if _source5.Is_DC2() {
				var _289___mcc_h3 _dafny.Sequence = _source5.Get_().(D0_DC2).Cf4
				_ = _289___mcc_h3
				var _290___mcc_h4 _dafny.Int = _source5.Get_().(D0_DC2).Cf5
				_ = _290___mcc_h4
				var _291___mcc_h5 _dafny.Int = _source5.Get_().(D0_DC2).Cf6
				_ = _291___mcc_h5
				var _292_cf6 _dafny.Int = _291___mcc_h5
				_ = _292_cf6
				var _293_cf5 _dafny.Int = _290___mcc_h4
				_ = _293_cf5
				var _294_cf4 _dafny.Sequence = _289___mcc_h3
				_ = _294_cf4
				(globalState).F1 = _293_cf5
				var _295_v44 _dafny.Map
				_ = _295_v44
				_295_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v2, (_this).F3())
				var _296_v45 *C0
				_ = _296_v45
				var _nw36 *C0 = New_C0_()
				_ = _nw36
				_nw36.Ctor__(_231_v3, _295_v44, (_this).F3(), Companion_Default___.SafeDivisionInt((_this).F4(), _263_i3))
				_296_v45 = _nw36
				_296_v45 = _296_v45
				var _297_v46 _dafny.Sequence
				_ = _297_v46
				var _out15 _dafny.Sequence
				_ = _out15
				_out15 = (_this).M1(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _dafny.IntOfInt64(484), (_dafny.Zero).Minus(_263_i3), globalState)
				_297_v46 = _out15
				var _298_v47 *C0
				_ = _298_v47
				var _nw37 *C0 = New_C0_()
				_ = _nw37
				_nw37.Ctor__((_296_v45).F7(), (_295_v44).Merge((_296_v45).F8()), (_this).F3(), _263_i3)
				_298_v47 = _nw37
			} else if _source5.Is_DC0() {
				var _299___mcc_h6 bool = _source5.Get_().(D0_DC0).Cf0
				_ = _299___mcc_h6
				var _300_cf0 bool = _299___mcc_h6
				_ = _300_cf0
				var _301_v48 _dafny.Map
				_ = _301_v48
				_301_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_cf0, (_this).F3())
				var _302_v49 *C0
				_ = _302_v49
				var _nw38 *C0 = New_C0_()
				_ = _nw38
				_nw38.Ctor__(_231_v3, _301_v48, (_this).F3(), (func() _dafny.Int {
					if _230_v2 {
						return _dafny.IntOfUint32((_228_v0).Cardinality())
					}
					return (_this).F4()
				})())
				_302_v49 = _nw38
				var _303_v50 D1
				_ = _303_v50
				_303_v50 = Companion_D1_.Create_DC6_()
				var _304_v51 D0
				_ = _304_v51
				_304_v51 = Companion_D0_.Create_DC1_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v2, _263_i3)).Cardinality(), _dafny.IntOfUint32((_233_v5).Cardinality()), _dafny.IntOfInt64(852))
				var _305_v52 _dafny.Map
				_ = _305_v52
				_305_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v50, _304_v51)
				_305_v52 = (_305_v52).Update(Companion_D1_.Create_DC6_(), _304_v51)
				var _306_v53 _dafny.Map
				_ = _306_v53
				_306_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_267_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_267_v34), 0))).Int()).(_dafny.Sequence)).Cardinality()), true)
				_236_v8 = (_236_v8).Update((func() bool {
					if (func() bool {
						if (_236_v8).Contains(_300_cf0) {
							return (_236_v8).Get(_300_cf0).(bool)
						}
						return _300_cf0
					})() {
						return (func() bool {
							if (_306_v53).Contains(_263_i3) {
								return (_306_v53).Get(_263_i3).(bool)
							}
							return Companion_Default___.Fm0((_dafny.Zero).Minus(Companion_Default___.Fm5(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), globalState)), (_this).F4(), globalState)
						})()
					}
					return _230_v2
				})(), _300_cf0)
				var _307_v54 _dafny.Map
				_ = _307_v54
				_307_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), (_dafny.SetOf(_230_v2, _230_v2)).Cardinality())
				var _308_v55 _dafny.Map
				_ = _308_v55
				_308_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v2, _236_v8)
				var _rhs56 _dafny.Int = Companion_Default___.Fm5(false, globalState)
				_ = _rhs56
				var _rhs57 _dafny.Map = _307_v54
				_ = _rhs57
				var _rhs58 bool = (((Companion_Default___.Fm6(true, ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _263_i3, globalState)).Cardinality()).Cmp(_263_i3) >= 0) || (_230_v2)
				_ = _rhs58
				var _rhs59 _dafny.Map = (func() _dafny.Map {
					if ((_this).F4()).Cmp(_263_i3) <= 0 {
						return _308_v55
					}
					return (_308_v55).Merge(_308_v55)
				})()
				_ = _rhs59
				var _rhs60 D0 = _231_v3
				_ = _rhs60
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				_lhs35.F1 = _rhs56
				_307_v54 = _rhs57
				_300_cf0 = _rhs58
				_308_v55 = _rhs59
				_231_v3 = _rhs60
			} else {
				var _309___mcc_h7 D0 = _source5.Get_().(D0_DC3).Cf7
				_ = _309___mcc_h7
				var _310_cf7 D0 = _309___mcc_h7
				_ = _310_cf7
				var _311_v56 _dafny.Array
				_ = _311_v56
				var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw39
				_311_v56 = _nw39
				var _312_v57 _dafny.Array
				_ = _312_v57
				var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw40
				_312_v57 = _nw40
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_311_v56), 0))
				_ = _index51
				(_311_v56).ArraySet1(_312_v57, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_311_v56), 0))
				_ = _index52
				(_311_v56).ArraySet1(_312_v57, (_index52).Int())
				var _313_v58 _dafny.CodePoint
				_ = _313_v58
				_313_v58 = _dafny.CodePoint('r')
				_313_v58 = _313_v58
				(globalState).F1 = Companion_Default___.SafeDivisionInt((_this).F4(), (Companion_Default___.Fm7(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), globalState)).Cardinality())
				(globalState).F1 = _dafny.IntOfInt64(-569)
			}
		}
		var _314_v59 _dafny.Sequence
		_ = _314_v59
		_314_v59 = _dafny.SeqOf(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _230_v2)
		r0 = (_314_v59).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_314_v59).Cardinality()))).Uint32()).(bool)
		r1 = _230_v2
		r2 = ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool)
		return r0, r1, r2
	}
}
func (_this *C1) M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _315_i0 _dafny.Int
		_ = _315_i0
		_315_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_315_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_315_i0 = (_315_i0).Plus(_dafny.One)
					var _316_v0 _dafny.Array
					_ = _316_v0
					var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
					_ = _nw41
					_316_v0 = _nw41
					_316_v0 = _316_v0
					var _317_v1 _dafny.Map
					_ = _317_v1
					_317_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
					var _318_v2 _dafny.Sequence
					_ = _318_v2
					_318_v2 = _dafny.SeqOf(_317_v1, _317_v1, _317_v1)
					var _319_v3 _dafny.Map
					_ = _319_v3
					_319_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-227)), (_318_v2).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_318_v2).Cardinality()))).Uint32()).(_dafny.Map))
					var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen(((_this).F3()), 0))
					_ = _index53
					((_this).F3()).ArraySet1(((func() _dafny.Map {
						if (_319_v3).Contains(p1) {
							return (_319_v3).Get(p1).(_dafny.Map)
						}
						return _317_v1
					})()).Cardinality(), (_index53).Int())
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen(((_this).F6()), 0))
					_ = _index54
					((_this).F6()).ArraySet1(p0, (_index54).Int())
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen(((_this).F3()), 0))
					_ = _index55
					var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen(((_this).F6()), 0))
					_ = _index56
					var _rhs61 _dafny.Int = p1
					_ = _rhs61
					var _rhs62 bool = !(p0)
					_ = _rhs62
					var _lhs36 _dafny.Array = (_this).F3()
					_ = _lhs36
					var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen(((_this).F3()), 0))
					_ = _lhs37
					var _lhs38 _dafny.Array = (_this).F6()
					_ = _lhs38
					var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen(((_this).F6()), 0))
					_ = _lhs39
					(_lhs36).ArraySet1(_rhs61, (_lhs37).Int())
					(_lhs38).ArraySet1(_rhs62, (_lhs39).Int())
					var _320_v4 _dafny.Sequence
					_ = _320_v4
					_320_v4 = _dafny.UnicodeSeqOfUtf8Bytes("k")
					var _321_v5 _dafny.Sequence
					_ = _321_v5
					_321_v5 = _dafny.SeqOf(_320_v4, _320_v4)
					var _322_v6 D0
					_ = _322_v6
					_322_v6 = Companion_D0_.Create_DC0_(p0)
					var _323_v7 _dafny.Array
					_ = _323_v7
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(19)
					_ = _len0_5
					var _nw42 _dafny.Array
					_ = _nw42
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw42 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.Int = func(_324_i1 _dafny.Int) _dafny.Int {
							return (_324_i1).Times((_this).F4())
						}
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw42 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw42).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw42).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_323_v7 = _nw42
					var _325_v8 _dafny.Map
					_ = _325_v8
					_325_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _323_v7)
					var _326_v9 _dafny.Sequence
					_ = _326_v9
					_326_v9 = _dafny.SeqOf(p0)
					var _327_v10 T0
					_ = _327_v10
					var _nw43 *C0 = New_C0_()
					_ = _nw43
					_nw43.Ctor__(_322_v6, _325_v8, _323_v7, _dafny.IntOfUint32((_326_v9).Cardinality()))
					_327_v10 = _nw43
					var _328_v11 _dafny.Map
					_ = _328_v11
					_328_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_321_v5, _321_v5), (Companion_Default___.SafeIndex(((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_321_v5, _321_v5)).Cardinality()))).Uint32(), _320_v4), _327_v10)
					var _329_v12 _dafny.Map
					_ = _329_v12
					_329_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _321_v5)
					var _330_v13 _dafny.Map
					_ = _330_v13
					_330_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _327_v10)
					_328_v11 = (_328_v11).Update((func() _dafny.Sequence {
						if (_329_v12).Contains((_327_v10).F4()) {
							return (_329_v12).Get((_327_v10).F4()).(_dafny.Sequence)
						}
						return _321_v5
					})(), (func() T0 {
						if (_330_v13).Contains(!(p0)) {
							return (_330_v13).Get(!(p0)).(T0)
						}
						return _327_v10
					})())
					var _331_v14 _dafny.Map
					_ = _331_v14
					_331_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v4, _dafny.UnicodeSeqOfUtf8Bytes("hcftnkrjh"))
					_331_v14 = (_331_v14).Update(_320_v4, _320_v4)
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _332_v15 _dafny.Array
		_ = _332_v15
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw44
		_332_v15 = _nw44
		_332_v15 = _332_v15
		var _333_v16 _dafny.Array
		_ = _333_v16
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_6
		var _nw45 _dafny.Array
		_ = _nw45
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw45 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = func(_334_i3 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_334_i3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("noju")).Cardinality()))
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw45 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw45).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw45).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_333_v16 = _nw45
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_333_v16), 0))); ; {
			_guard_loop_1, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _335_i2 _dafny.Int
			_335_i2 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_335_i2).Sign() != -1) && ((_335_i2).Cmp(_dafny.ArrayLen((_333_v16), 0)) < 0)) {
				(_333_v16).ArraySet1((_335_i2).Plus((_dafny.Zero).Minus((_this).F4())), (_335_i2).Int())
			}
		}
		if p0 {
			var _336_v17 D1
			_ = _336_v17
			_336_v17 = Companion_D1_.Create_DC4_((_this).F3())
			var _337_v18 _dafny.MultiSet
			_ = _337_v18
			_337_v18 = _dafny.MultiSetOf((_336_v17).Dtor_cf8())
			var _338_v19 _dafny.Sequence
			_ = _338_v19
			_338_v19 = _dafny.SeqOf(_dafny.MultiSetOf(_333_v16), _337_v18, _337_v18)
			var _339_v20 D0
			_ = _339_v20
			_339_v20 = Companion_D0_.Create_DC0_(p0)
			var _340_v21 _dafny.Map
			_ = _340_v21
			_340_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _333_v16)
			var _341_v22 *C0
			_ = _341_v22
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__(_339_v20, _340_v21, (_this).F3(), p2)
			_341_v22 = _nw46
			var _342_v23 _dafny.Map
			_ = _342_v23
			_342_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_336_v17).Dtor_cf8())).IsProperSubsetOf((_338_v19).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_338_v19).Cardinality()))).Uint32()).(_dafny.MultiSet)), _341_v22)
			var _343_v24 bool
			_ = _343_v24
			_343_v24 = false
			var _344_v25 _dafny.CodePoint
			_ = _344_v25
			_344_v25 = _dafny.CodePoint('p')
			var _345_v26 _dafny.Map
			_ = _345_v26
			_345_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_344_v25, (_this).F5())
			var _346_v27 _dafny.Sequence
			_ = _346_v27
			_346_v27 = _dafny.SeqOf((func() _dafny.CodePoint {
				if (_345_v26).Contains(_344_v25) {
					return (_345_v26).Get(_344_v25).(_dafny.CodePoint)
				}
				return _344_v25
			})())
			var _347_v28 _dafny.Set
			_ = _347_v28
			_347_v28 = _dafny.SetOf(p1)
			var _348_v29 _dafny.Array
			_ = _348_v29
			var _nwElement0_6 _dafny.Sequence = _346_v27
			_ = _nwElement0_6
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(23))
			_ = _nw47
			(_nw47).ArraySet1(_nwElement0_6, 0)
			(_nw47).ArraySet1(_346_v27, 1)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if _343_v24 {
					return _346_v27
				}
				return _dafny.SeqOf(_344_v25, (_this).F5())
			})(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _343_v24 {
					return _346_v27
				}
				return _dafny.SeqOf(_344_v25, (_this).F5())
			})()).Cardinality()))).Uint32(), _dafny.CodePoint('l')), 2)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(940))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_349_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_350_i4 _dafny.Int) _dafny.CodePoint {
					return _349_v25
				}
			})(_344_v25))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(92))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_351_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_352_i5 _dafny.Int) _dafny.CodePoint {
					return _351_v25
				}
			})(_344_v25)))), 3)
			(_nw47).ArraySet1(_346_v27, 4)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_346_v27, _346_v27), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_346_v27, _346_v27)).Cardinality()))).Uint32(), (_346_v27).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_346_v27).Cardinality()))).Uint32()).(_dafny.CodePoint)), 5)
			(_nw47).ArraySet1(_346_v27, 6)
			(_nw47).ArraySet1(_346_v27, 7)
			(_nw47).ArraySet1(_346_v27, 8)
			(_nw47).ArraySet1(_346_v27, 9)
			(_nw47).ArraySet1(_346_v27, 10)
			(_nw47).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}(func(_353_i6 _dafny.Int) _dafny.CodePoint {
				return (_this).F5()
			})), 11)
			(_nw47).ArraySet1(_346_v27, 12)
			(_nw47).ArraySet1(_346_v27, 13)
			(_nw47).ArraySet1(_346_v27, 14)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_346_v27, _346_v27), 15)
			(_nw47).ArraySet1(_346_v27, 16)
			(_nw47).ArraySet1((func() _dafny.Sequence {
				if p0 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(938))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}(func(_354_i7 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					}))
				}
				return _dafny.SeqOf((_this).F5())
			})(), 17)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_346_v27, _346_v27), 18)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_346_v27, (Companion_Default___.SafeIndex((_347_v28).Cardinality(), _dafny.IntOfUint32((_346_v27).Cardinality()))).Uint32(), (_this).F5()), _346_v27), 19)
			(_nw47).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(516))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_355_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_356_i8 _dafny.Int) _dafny.CodePoint {
					return _355_v25
				}
			})(_344_v25))), 20)
			(_nw47).ArraySet1(_346_v27, 21)
			(_nw47).ArraySet1(_dafny.SeqOf((_this).F5()), 22)
			_348_v29 = _nw47
			var _357_v30 _dafny.Sequence
			_ = _357_v30
			_357_v30 = _dafny.SeqOf(_346_v27)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))
			_ = _index57
			(_348_v29).ArraySet1((_357_v30).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_357_v30).Cardinality()))).Uint32()).(_dafny.Sequence), (_index57).Int())
			var _358_v31 _dafny.Map
			_ = _358_v31
			_358_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
			_ = _index58
			((_this).F3()).ArraySet1((func() _dafny.Int {
				if (_358_v31).Contains(p0) {
					return (_358_v31).Get(p0).(_dafny.Int)
				}
				return p1
			})(), (_index58).Int())
			var _359_v32 _dafny.Sequence
			_ = _359_v32
			_359_v32 = _dafny.SeqOf(_343_v24, p0)
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))
			_ = _index59
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
			_ = _index60
			var _rhs63 _dafny.Map = (_342_v23).Update(_343_v24, _341_v22)
			_ = _rhs63
			var _rhs64 bool = p0
			_ = _rhs64
			var _rhs65 _dafny.CodePoint = _344_v25
			_ = _rhs65
			var _rhs66 _dafny.Sequence = Companion_Default___.Fm1(p0, (Companion_Default___.Fm5(_343_v24, globalState)).Cmp(_dafny.IntOfUint32((_359_v32).Cardinality())) >= 0, globalState)
			_ = _rhs66
			var _rhs67 _dafny.Int = _dafny.IntOfInt64(157)
			_ = _rhs67
			var _lhs40 _dafny.Array = _348_v29
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))
			_ = _lhs41
			var _lhs42 _dafny.Array = (_this).F3()
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
			_ = _lhs43
			_342_v23 = _rhs63
			_343_v24 = _rhs64
			_344_v25 = _rhs65
			(_lhs40).ArraySet1(_rhs66, (_lhs41).Int())
			(_lhs42).ArraySet1(_rhs67, (_lhs43).Int())
			if p0 {
				_343_v24 = _343_v24
				var _360_v33 _dafny.Array
				_ = _360_v33
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_7
				var _nw48 _dafny.Array
				_ = _nw48
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw48 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = (func(_361_p0 bool, _362_v24 bool) func(_dafny.Int) bool {
						return func(_363_i9 _dafny.Int) bool {
							return (func() bool {
								if _361_p0 {
									return _361_p0
								}
								return _362_v24
							})()
						}
					})(p0, _343_v24)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw48 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw48).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw48).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_360_v33 = _nw48
				var _364_v34 _dafny.Array
				_ = _364_v34
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_8
				var _nw49 _dafny.Array
				_ = _nw49
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw49 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Sequence = (func(_365_v24 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_366_i10 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_365_v24)
						}
					})(_343_v24)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw49 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw49).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw49).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_364_v34 = _nw49
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))
				_ = _index61
				var _rhs68 _dafny.Array = (_this).F6()
				_ = _rhs68
				var _rhs69 _dafny.CodePoint = (_this).F5()
				_ = _rhs69
				var _rhs70 _dafny.Array = _364_v34
				_ = _rhs70
				var _rhs71 _dafny.Sequence = (_348_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))).Int()).(_dafny.Sequence)
				_ = _rhs71
				var _rhs72 bool = Companion_Default___.Fm0(p2, p2, globalState)
				_ = _rhs72
				var _lhs44 _dafny.Array = _348_v29
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))
				_ = _lhs45
				_360_v33 = _rhs68
				_344_v25 = _rhs69
				_364_v34 = _rhs70
				(_lhs44).ArraySet1(_rhs71, (_lhs45).Int())
				_343_v24 = _rhs72
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index62
				((_this).F3()).ArraySet1((func() _dafny.Int {
					if (p0) || (_343_v24) {
						return Companion_Default___.Fm5(!(_343_v24), globalState)
					}
					return _dafny.IntOfInt64(469)
				})(), (_index62).Int())
				(globalState).F1 = _dafny.IntOfInt64(683)
				var _367_v35 _dafny.MultiSet
				_ = _367_v35
				_367_v35 = _dafny.MultiSetOf(p0)
				var _368_v36 _dafny.Map
				_ = _368_v36
				_368_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F4(), ((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int), globalState), _367_v35)
				_368_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v24, _367_v35)
			} else {
				(globalState).F1 = Companion_Default___.Fm5(p0, globalState)
				var _369_v37 _dafny.Map
				_ = _369_v37
				_369_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen(((_this).F6()), 0))
				_ = _index63
				((_this).F6()).ArraySet1((func() bool {
					if (_369_v37).Contains(p0) {
						return (_369_v37).Get(p0).(bool)
					}
					return p0
				})(), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen(((_this).F6()), 0))
				_ = _index64
				((_this).F6()).ArraySet1((_dafny.IntOfInt64(175)).Cmp(p1) == 0, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index65
				((_this).F3()).ArraySet1(p1, (_index65).Int())
				var _nw50 *C0 = New_C0_()
				_ = _nw50
				_nw50.Ctor__(_339_v20, ((_341_v22).F8()).Update(((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool), _333_v16), _333_v16, _dafny.IntOfInt64(-771))
				_341_v22 = _nw50
				(globalState).F1 = (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if ((_this).F6()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen(((_this).F6()), 0))).Int()).(bool) {
						return _dafny.UnicodeSeqOfUtf8Bytes("i")
					}
					return _346_v27
				})()).Cardinality()))
			}
			if p0 {
				var _370_v38 D0
				_ = _370_v38
				_370_v38 = Companion_D0_.Create_DC2_(_346_v27, ((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int), p1)
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))
				_ = _index66
				(_348_v29).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_348_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))).Int()).(_dafny.Sequence), (_370_v38).Dtor_cf4()), (_348_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))).Int()).(_dafny.Sequence)), (_index66).Int())
				var _371_v39 *C0
				_ = _371_v39
				var _nw51 *C0 = New_C0_()
				_ = _nw51
				_nw51.Ctor__(Companion_D0_.Create_DC0_(_343_v24), ((_341_v22).F8()).Merge((_341_v22).F8()), _333_v16, (_dafny.IntOfInt64(-897)).Minus(p1))
				_371_v39 = _nw51
				var _372_v40 *C0
				_ = _372_v40
				var _nw52 *C0 = New_C0_()
				_ = _nw52
				_nw52.Ctor__(Companion_D0_.Create_DC0_(_343_v24), (_371_v39).F8(), _333_v16, Companion_Default___.SafeDivisionInt(p2, p1))
				_372_v40 = _nw52
				(globalState).F1 = p1
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index67
				((_this).F3()).ArraySet1(p1, (_index67).Int())
			} else {
				_358_v31 = ((_358_v31).Merge(_358_v31)).Merge(_358_v31)
				var _373_v41 _dafny.Map
				_ = _373_v41
				_373_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v24, _343_v24)
				var _374_v42 _dafny.Set
				_ = _374_v42
				_374_v42 = _dafny.SetOf((_341_v22).F7(), _339_v20, (_341_v22).F7())
				var _375_v43 _dafny.Array
				_ = _375_v43
				var _nwElement0_7 bool = p0
				_ = _nwElement0_7
				var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(19))
				_ = _nw53
				(_nw53).ArraySet1(_nwElement0_7, 0)
				(_nw53).ArraySet1(_343_v24, 1)
				(_nw53).ArraySet1(p0, 2)
				(_nw53).ArraySet1(_343_v24, 3)
				(_nw53).ArraySet1(p0, 4)
				(_nw53).ArraySet1(_343_v24, 5)
				(_nw53).ArraySet1(_343_v24, 6)
				(_nw53).ArraySet1((((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int)).Cmp((_this).F4()) < 0, 7)
				(_nw53).ArraySet1(((_this).F4()).Cmp(p2) <= 0, 8)
				(_nw53).ArraySet1(p0, 9)
				(_nw53).ArraySet1(!(_343_v24), 10)
				(_nw53).ArraySet1(_343_v24, 11)
				(_nw53).ArraySet1((func() bool {
					if (_373_v41).Contains(p0) {
						return (_373_v41).Get(p0).(bool)
					}
					return p0
				})(), 12)
				(_nw53).ArraySet1((_374_v42).IsProperSubsetOf(Companion_Default___.Fm8(_343_v24, globalState)), 13)
				(_nw53).ArraySet1(p0, 14)
				(_nw53).ArraySet1(_343_v24, 15)
				(_nw53).ArraySet1(p0, 16)
				(_nw53).ArraySet1((((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
					if (_358_v31).Contains(Companion_Default___.Fm0(((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int), (_347_v28).Cardinality(), globalState)) {
						return (_358_v31).Get(Companion_Default___.Fm0(((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int), (_347_v28).Cardinality(), globalState)).(_dafny.Int)
					}
					return p1
				})()) < 0, 17)
				(_nw53).ArraySet1(!(true), 18)
				_375_v43 = _nw53
				var _len0_9 _dafny.Int = _dafny.One
				_ = _len0_9
				var _nw54 _dafny.Array
				_ = _nw54
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw54 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) bool = (func(_376_p0 bool) func(_dafny.Int) bool {
						return func(_377_i11 _dafny.Int) bool {
							return _376_p0
						}
					})(p0)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw54 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw54).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw54).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_375_v43 = _nw54
				var _378_v44 _dafny.Sequence
				_ = _378_v44
				_378_v44 = _dafny.SeqOf(_dafny.IntOfInt64(247), (_this).F4())
				var _379_v45 _dafny.Set
				_ = _379_v45
				_379_v45 = _dafny.SetOf(!(!(_343_v24)), true)
				var _380_v46 _dafny.Map
				_ = _380_v46
				_380_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(p0, globalState), _379_v45)
				var _381_v47 _dafny.Map
				_ = _381_v47
				_381_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_346_v27, _346_v27), _dafny.Companion_Sequence_.Concatenate(_378_v44, _dafny.Companion_Sequence_.Update(_378_v44, (Companion_Default___.SafeIndex((_380_v46).Cardinality(), _dafny.IntOfUint32((_378_v44).Cardinality()))).Uint32(), _dafny.IntOfInt64(32))))
				var _rhs73 bool = _343_v24
				_ = _rhs73
				var _rhs74 _dafny.Map = (_358_v31).Update(!(p0), ((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int))
				_ = _rhs74
				var _rhs75 _dafny.Sequence = (func() _dafny.Sequence {
					if (_381_v47).Contains(_346_v27) {
						return (_381_v47).Get(_346_v27).(_dafny.Sequence)
					}
					return _378_v44
				})()
				_ = _rhs75
				_343_v24 = _rhs73
				_358_v31 = _rhs74
				_378_v44 = _rhs75
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index68
				((_this).F3()).ArraySet1(((_this).F4()).Plus(p2), (_index68).Int())
				var _382_v48 D1
				_ = _382_v48
				_382_v48 = Companion_D1_.Create_DC5_((p0) && (_343_v24), true, ((_this).F3()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen(((_this).F3()), 0))).Int()).(_dafny.Int))
				var _pat_let_tv3 = _382_v48
				_ = _pat_let_tv3
				var _pat_let_tv4 = _343_v24
				_ = _pat_let_tv4
				_382_v48 = func(_pat_let10_0 D1) D1 {
					return func(_383_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let11_0 bool) D1 {
							return func(_384_dt__update_hcf9_h0 bool) D1 {
								return func(_pat_let12_0 bool) D1 {
									return func(_385_dt__update_hcf10_h0 bool) D1 {
										return Companion_D1_.Create_DC5_(_384_dt__update_hcf9_h0, _385_dt__update_hcf10_h0, (_383_dt__update__tmp_h0).Dtor_cf11())
									}(_pat_let12_0)
								}(_pat_let_tv4)
							}(_pat_let11_0)
						}((_pat_let_tv3).Dtor_cf9())
					}(_pat_let10_0)
				}(Companion_D1_.Create_DC5_(p0, p0, _dafny.IntOfInt64(-408)))
			}
			_343_v24 = (_dafny.IntOfInt64(-255)).Cmp(p2) >= 0
			_346_v27 = (_348_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_348_v29), 0))).Int()).(_dafny.Sequence)
		} else {
			var _386_v49 _dafny.Sequence
			_ = _386_v49
			_386_v49 = _dafny.UnicodeSeqOfUtf8Bytes("uaymukwrs")
			var _387_v50 _dafny.Array
			_ = _387_v50
			var _nwElement0_8 _dafny.CodePoint = (_this).F5()
			_ = _nwElement0_8
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(13))
			_ = _nw55
			(_nw55).ArraySet1CodePoint(_nwElement0_8, 0)
			(_nw55).ArraySet1CodePoint((_386_v49).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_386_v49).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
			(_nw55).ArraySet1CodePoint((_this).F5(), 2)
			(_nw55).ArraySet1CodePoint((_this).F5(), 3)
			(_nw55).ArraySet1CodePoint(Companion_Default___.Fm3((_this).F5(), p2, globalState), 4)
			(_nw55).ArraySet1CodePoint((_this).F5(), 5)
			(_nw55).ArraySet1CodePoint((_this).F5(), 6)
			(_nw55).ArraySet1CodePoint((_this).F5(), 7)
			(_nw55).ArraySet1CodePoint((_this).F5(), 8)
			(_nw55).ArraySet1CodePoint((_this).F5(), 9)
			(_nw55).ArraySet1CodePoint((_this).F5(), 10)
			(_nw55).ArraySet1CodePoint(_dafny.CodePoint('x'), 11)
			(_nw55).ArraySet1CodePoint((_this).F5(), 12)
			_387_v50 = _nw55
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))
			_ = _index69
			(_387_v50).ArraySet1CodePoint((_this).F5(), (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))
			_ = _index70
			(_387_v50).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return (_this).F5()
				}
				return (_this).F5()
			})(), (_index70).Int())
			var _388_v51 bool
			_ = _388_v51
			_388_v51 = false
			_388_v51 = p0
			var _389_v52 _dafny.Map
			_ = _389_v52
			_389_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(553), _386_v49)
			var _390_v53 _dafny.Set
			_ = _390_v53
			_390_v53 = _dafny.SetOf((_389_v52).Cardinality())
			var _391_v54 _dafny.Sequence
			_ = _391_v54
			_391_v54 = _dafny.SeqOf(_388_v51, _388_v51)
			var _392_v55 D0
			_ = _392_v55
			_392_v55 = Companion_D0_.Create_DC2_(_386_v49, (_390_v53).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_391_v54).Cardinality())))
			var _393_v56 D0
			_ = _393_v56
			_393_v56 = Companion_D0_.Create_DC3_(_392_v55)
			var _394_v57 D0
			_ = _394_v57
			_394_v57 = Companion_D0_.Create_DC3_(_393_v56)
			var _source6 D0 = _394_v57
			_ = _source6
			if _source6.Is_DC1() {
				var _395___mcc_h0 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
				_ = _395___mcc_h0
				var _396___mcc_h1 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
				_ = _396___mcc_h1
				var _397___mcc_h2 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
				_ = _397___mcc_h2
				var _398_cf3 _dafny.Int = _397___mcc_h2
				_ = _398_cf3
				var _399_cf2 _dafny.Int = _396___mcc_h1
				_ = _399_cf2
				var _400_cf1 _dafny.Int = _395___mcc_h0
				_ = _400_cf1
				var _401_v58 _dafny.Sequence
				_ = _401_v58
				_401_v58 = _dafny.SeqOf(Companion_D0_.Create_DC2_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}(func(_402_i12 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				})), (_this).F4(), p2))
				var _403_v59 _dafny.Sequence
				_ = _403_v59
				_403_v59 = _dafny.SeqOf(_dafny.IntOfUint32((_401_v58).Cardinality()))
				var _404_v60 _dafny.Array
				_ = _404_v60
				var _nwElement0_9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_391_v54, _391_v54)
				_ = _nwElement0_9
				var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(10))
				_ = _nw56
				(_nw56).ArraySet1(_nwElement0_9, 0)
				(_nw56).ArraySet1(_391_v54, 1)
				(_nw56).ArraySet1(_391_v54, 2)
				(_nw56).ArraySet1(_391_v54, 3)
				(_nw56).ArraySet1(_dafny.SeqOf(p0, _388_v51, p0, true), 4)
				(_nw56).ArraySet1(_391_v54, 5)
				(_nw56).ArraySet1(_391_v54, 6)
				(_nw56).ArraySet1((func() _dafny.Sequence {
					if p0 {
						return _391_v54
					}
					return _391_v54
				})(), 7)
				(_nw56).ArraySet1(_dafny.SeqOf(p0), 8)
				(_nw56).ArraySet1(_391_v54, 9)
				_404_v60 = _nw56
				var _405_v61 _dafny.MultiSet
				_ = _405_v61
				_405_v61 = _dafny.MultiSetOf(_403_v59)
				var _rhs76 _dafny.Sequence = (func() _dafny.Sequence {
					if true {
						return _403_v59
					}
					return _dafny.Companion_Sequence_.Update(_403_v59, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-229), _dafny.IntOfUint32((_403_v59).Cardinality()))).Uint32(), (_405_v61).Cardinality())
				})()
				_ = _rhs76
				var _rhs77 _dafny.Int = _398_cf3
				_ = _rhs77
				var _rhs78 _dafny.Array = _404_v60
				_ = _rhs78
				var _rhs79 _dafny.Int = (_dafny.Zero).Minus(p1)
				_ = _rhs79
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				_403_v59 = _rhs76
				_398_cf3 = _rhs77
				_404_v60 = _rhs78
				_lhs46.F1 = _rhs79
				var _406_v62 _dafny.Map
				_ = _406_v62
				_406_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F3())
				var _407_v63 *C0
				_ = _407_v63
				var _nw57 *C0 = New_C0_()
				_ = _nw57
				_nw57.Ctor__(Companion_D0_.Create_DC0_(false), _406_v62, _333_v16, _398_cf3)
				_407_v63 = _nw57
				var _408_v64 _dafny.Sequence
				_ = _408_v64
				_408_v64 = _dafny.SeqOf(_407_v63)
				_386_v49 = _dafny.Companion_Sequence_.Update(_386_v49, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_408_v64, _408_v64)).Cardinality()), _dafny.IntOfUint32((_386_v49).Cardinality()))).Uint32(), _dafny.CodePoint('b'))
				_394_v57 = _394_v57
				var _nwElement0_10 _dafny.CodePoint = (_386_v49).Select((Companion_Default___.SafeIndex(_398_cf3, _dafny.IntOfUint32((_386_v49).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _nwElement0_10
				var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
				_ = _nw58
				(_nw58).ArraySet1CodePoint(_nwElement0_10, 0)
				(_nw58).ArraySet1CodePoint((_this).F5(), 1)
				(_nw58).ArraySet1CodePoint((_this).F5(), 2)
				(_nw58).ArraySet1CodePoint((_this).F5(), 3)
				(_nw58).ArraySet1CodePoint((_this).F5(), 4)
				(_nw58).ArraySet1CodePoint((_this).F5(), 5)
				(_nw58).ArraySet1CodePoint((_this).F5(), 6)
				(_nw58).ArraySet1CodePoint((_387_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))).Int()), 7)
				(_nw58).ArraySet1CodePoint((_387_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))).Int()), 8)
				(_nw58).ArraySet1CodePoint((_387_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))).Int()), 9)
				(_nw58).ArraySet1CodePoint((_387_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))).Int()), 10)
				(_nw58).ArraySet1CodePoint(_dafny.CodePoint('b'), 11)
				(_nw58).ArraySet1CodePoint((_387_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))).Int()), 12)
				_387_v50 = _nw58
			} else if _source6.Is_DC2() {
				var _409___mcc_h3 _dafny.Sequence = _source6.Get_().(D0_DC2).Cf4
				_ = _409___mcc_h3
				var _410___mcc_h4 _dafny.Int = _source6.Get_().(D0_DC2).Cf5
				_ = _410___mcc_h4
				var _411___mcc_h5 _dafny.Int = _source6.Get_().(D0_DC2).Cf6
				_ = _411___mcc_h5
				var _412_cf6 _dafny.Int = _411___mcc_h5
				_ = _412_cf6
				var _413_cf5 _dafny.Int = _410___mcc_h4
				_ = _413_cf5
				var _414_cf4 _dafny.Sequence = _409___mcc_h3
				_ = _414_cf4
				var _415_v65 _dafny.Map
				_ = _415_v65
				_415_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F3())
				var _416_v66 *C0
				_ = _416_v66
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(Companion_D0_.Create_DC0_(p0), _415_v65, _333_v16, p1)
				_416_v66 = _nw59
				var _417_v67 _dafny.Array
				_ = _417_v67
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_10
				var _nw60 _dafny.Array
				_ = _nw60
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw60 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Sequence = (func(_418_v57 D0) func(_dafny.Int) _dafny.Sequence {
						return func(_419_i13 _dafny.Int) _dafny.Sequence {
							return (func() _dafny.Sequence {
								if true {
									return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer29 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
										return func(arg29 _dafny.Int) interface{} {
											return coer29(arg29)
										}
									}((func(_420_v57 D0) func(_dafny.Int) D0 {
										return func(_421_i14 _dafny.Int) D0 {
											return _420_v57
										}
									})(_418_v57)))
								}
								return _dafny.SeqOf(_418_v57, _418_v57)
							})()
						}
					})(_394_v57)
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
				_417_v67 = _nw60
				var _422_v68 D1
				_ = _422_v68
				_422_v68 = Companion_D1_.Create_DC5_(_388_v51, _388_v51, p1)
				var _rhs80 _dafny.Int = p1
				_ = _rhs80
				var _rhs81 _dafny.Array = _417_v67
				_ = _rhs81
				var _rhs82 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_414_cf4, _dafny.Companion_Sequence_.Update(_414_cf4, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_414_cf4).Cardinality()))).Uint32(), (_387_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_387_v50), 0))).Int())))).Cardinality()))
				_ = _rhs82
				var _rhs83 bool = (Companion_Default___.SafeModuloInt(p1, _dafny.IntOfInt64(856))).Cmp((_422_v68).Dtor_cf11()) >= 0
				_ = _rhs83
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				var _lhs48 *GlobalState = globalState
				_ = _lhs48
				_lhs47.F1 = _rhs80
				_417_v67 = _rhs81
				_lhs48.F1 = _rhs82
				_388_v51 = _rhs83
				var _423_v69 _dafny.Sequence
				_ = _423_v69
				_423_v69 = _dafny.SeqOf(Companion_Default___.Fm5(_388_v51, globalState))
				var _424_v70 _dafny.Map
				_ = _424_v70
				_424_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(true, globalState), _dafny.IntOfInt64(-217))
				var _425_v71 _dafny.Array
				_ = _425_v71
				var _nwElement0_11 bool = _388_v51
				_ = _nwElement0_11
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(19))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_11, 0)
				(_nw61).ArraySet1(((_this).F4()).Cmp(_412_cf6) < 0, 1)
				(_nw61).ArraySet1(_388_v51, 2)
				(_nw61).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_423_v69, _423_v69), 3)
				(_nw61).ArraySet1((p1).Cmp((_390_v53).Cardinality()) <= 0, 4)
				(_nw61).ArraySet1(p0, 5)
				(_nw61).ArraySet1((_390_v53).IsProperSubsetOf(_390_v53), 6)
				(_nw61).ArraySet1(p0, 7)
				(_nw61).ArraySet1((_388_v51) == (_388_v51), 8)
				(_nw61).ArraySet1(p0, 9)
				(_nw61).ArraySet1(p0, 10)
				(_nw61).ArraySet1(_388_v51, 11)
				(_nw61).ArraySet1((_391_v54).Select((Companion_Default___.SafeIndex(_412_cf6, _dafny.IntOfUint32((_391_v54).Cardinality()))).Uint32()).(bool), 12)
				(_nw61).ArraySet1(true, 13)
				(_nw61).ArraySet1(_388_v51, 14)
				(_nw61).ArraySet1(((func() _dafny.Int {
					if (_424_v70).Contains((_dafny.Zero).Minus(_412_cf6)) {
						return (_424_v70).Get((_dafny.Zero).Minus(_412_cf6)).(_dafny.Int)
					}
					return (_dafny.Zero).Minus((_this).F4())
				})()).Cmp((_this).F4()) != 0, 15)
				(_nw61).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(752))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_426_v50 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_427_i15 _dafny.Int) _dafny.CodePoint {
						return (_426_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_426_v50), 0))).Int())
					}
				})(_387_v50))), _386_v49), 16)
				(_nw61).ArraySet1(!(_388_v51), 17)
				(_nw61).ArraySet1(p0, 18)
				_425_v71 = _nw61
				_425_v71 = _425_v71
				(globalState).F1 = p1
			} else if _source6.Is_DC0() {
				var _428___mcc_h6 bool = _source6.Get_().(D0_DC0).Cf0
				_ = _428___mcc_h6
				var _429_cf0 bool = _428___mcc_h6
				_ = _429_cf0
				var _430_v72 D1
				_ = _430_v72
				_430_v72 = Companion_D1_.Create_DC4_(_333_v16)
				var _431_v73 _dafny.Map
				_ = _431_v73
				_431_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_429_cf0), Companion_Default___.Fm5(true, globalState))
				var _rhs84 _dafny.Int = ((func() _dafny.Int {
					if p0 {
						return (_dafny.Zero).Minus((_431_v73).Cardinality())
					}
					return (_this).F4()
				})()).Plus((_this).F4())
				_ = _rhs84
				var _rhs85 D1 = _430_v72
				_ = _rhs85
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				_lhs49.F1 = _rhs84
				_430_v72 = _rhs85
				var _432_v74 D0
				_ = _432_v74
				_432_v74 = Companion_D0_.Create_DC0_(_388_v51)
				var _433_v75 _dafny.Map
				_ = _433_v75
				_433_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_cf0, (_this).F3())
				var _434_v76 *C0
				_ = _434_v76
				var _nw62 *C0 = New_C0_()
				_ = _nw62
				_nw62.Ctor__(_432_v74, _433_v75, (_this).F3(), p2)
				_434_v76 = _nw62
				(globalState).F1 = (_this).F4()
				var _rhs86 _dafny.Int = p1
				_ = _rhs86
				var _rhs87 _dafny.Int = Companion_Default___.SafeModuloInt(p1, p1)
				_ = _rhs87
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				_lhs50.F1 = _rhs86
				_lhs51.F1 = _rhs87
			} else {
				var _435___mcc_h7 D0 = _source6.Get_().(D0_DC3).Cf7
				_ = _435___mcc_h7
				var _436_cf7 D0 = _435___mcc_h7
				_ = _436_cf7
				(globalState).F1 = p2
				var _437_v77 _dafny.Map
				_ = _437_v77
				_437_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
				var _438_v78 _dafny.Map
				_ = _438_v78
				_438_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_388_v51) && (false), (_386_v49).Select((Companion_Default___.SafeIndex((_437_v77).Cardinality(), _dafny.IntOfUint32((_386_v49).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_438_v78 = (_438_v78).Update(_388_v51, (_this).F5())
				var _439_v79 _dafny.Map
				_ = _439_v79
				_439_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), p1)
				_439_v79 = (_439_v79).Update(_333_v16, p2)
				var _440_v80 _dafny.Array
				_ = _440_v80
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw63
				_440_v80 = _nw63
				_440_v80 = (_this).F6()
			}
			var _441_v81 _dafny.Sequence
			_ = _441_v81
			_441_v81 = _dafny.SeqOf(_387_v50, _387_v50)
			var _rhs88 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_441_v81, _dafny.Companion_Sequence_.Concatenate(_441_v81, _dafny.SeqOf((_441_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_391_v54).Cardinality()), _dafny.IntOfUint32((_441_v81).Cardinality()))).Uint32()).(_dafny.Array), _387_v50, _387_v50)))
			_ = _rhs88
			var _rhs89 _dafny.Int = (_this).F4()
			_ = _rhs89
			var _rhs90 bool = !(!_dafny.Companion_Sequence_.Equal(_386_v49, _386_v49))
			_ = _rhs90
			var _lhs52 *GlobalState = globalState
			_ = _lhs52
			_388_v51 = _rhs88
			_lhs52.F1 = _rhs89
			_388_v51 = _rhs90
			_388_v51 = _388_v51
		}
		var _442_v82 _dafny.Map
		_ = _442_v82
		_442_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(p2))
		var _443_v83 D0
		_ = _443_v83
		_443_v83 = Companion_D0_.Create_DC1_((_this).F4(), (_dafny.Zero).Minus(p2), (_442_v82).Cardinality())
		var _444_v84 _dafny.Sequence
		_ = _444_v84
		_444_v84 = _dafny.UnicodeSeqOfUtf8Bytes("uohkfjswt")
		var _hi2 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F4(), Companion_Default___.Fm5(p0, globalState))
		_ = _hi2
		for _445_i16 := Companion_Default___.SafeModuloInt((_443_v83).Dtor_cf2(), _dafny.IntOfUint32((_444_v84).Cardinality())); _445_i16.Cmp(_hi2) < 0; _445_i16 = _445_i16.Plus(_dafny.One) {
			var _446_v85 _dafny.Sequence
			_ = _446_v85
			_446_v85 = _dafny.SeqOf(p1)
			var _447_v86 _dafny.Array
			_ = _447_v86
			var _nwElement0_12 _dafny.CodePoint = Companion_Default___.Fm3((_this).F5(), (_this).F4(), globalState)
			_ = _nwElement0_12
			var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(22))
			_ = _nw64
			(_nw64).ArraySet1CodePoint(_nwElement0_12, 0)
			(_nw64).ArraySet1CodePoint(_dafny.CodePoint('f'), 1)
			(_nw64).ArraySet1CodePoint((_this).F5(), 2)
			(_nw64).ArraySet1CodePoint((_this).F5(), 3)
			(_nw64).ArraySet1CodePoint(Companion_Default___.Fm3((_this).F5(), p1, globalState), 4)
			(_nw64).ArraySet1CodePoint((_this).F5(), 5)
			(_nw64).ArraySet1CodePoint((_this).F5(), 6)
			(_nw64).ArraySet1CodePoint((_this).F5(), 7)
			(_nw64).ArraySet1CodePoint(_dafny.CodePoint('a'), 8)
			(_nw64).ArraySet1CodePoint((_this).F5(), 9)
			(_nw64).ArraySet1CodePoint((_this).F5(), 10)
			(_nw64).ArraySet1CodePoint((_this).F5(), 11)
			(_nw64).ArraySet1CodePoint((_this).F5(), 12)
			(_nw64).ArraySet1CodePoint((_this).F5(), 13)
			(_nw64).ArraySet1CodePoint(_dafny.CodePoint('j'), 14)
			(_nw64).ArraySet1CodePoint((_this).F5(), 15)
			(_nw64).ArraySet1CodePoint((_this).F5(), 16)
			(_nw64).ArraySet1CodePoint((_this).F5(), 17)
			(_nw64).ArraySet1CodePoint((_444_v84).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_446_v85).Cardinality()), _dafny.IntOfUint32((_444_v84).Cardinality()))).Uint32()).(_dafny.CodePoint), 18)
			(_nw64).ArraySet1CodePoint((_this).F5(), 19)
			(_nw64).ArraySet1CodePoint((_this).F5(), 20)
			(_nw64).ArraySet1CodePoint((_this).F5(), 21)
			_447_v86 = _nw64
			var _448_v87 _dafny.Sequence
			_ = _448_v87
			_448_v87 = _dafny.SeqOf(_447_v86, _447_v86)
			var _449_v88 _dafny.Set
			_ = _449_v88
			_449_v88 = _dafny.SetOf((_this).F5())
			var _source7 D0 = Companion_D0_.Create_DC2_(_444_v84, _dafny.IntOfUint32((_448_v87).Cardinality()), Companion_Default___.SafeDivisionInt(p2, (_449_v88).Cardinality()))
			_ = _source7
			if _source7.Is_DC1() {
				var _450___mcc_h8 _dafny.Int = _source7.Get_().(D0_DC1).Cf1
				_ = _450___mcc_h8
				var _451___mcc_h9 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
				_ = _451___mcc_h9
				var _452___mcc_h10 _dafny.Int = _source7.Get_().(D0_DC1).Cf3
				_ = _452___mcc_h10
				var _453_cf3 _dafny.Int = _452___mcc_h10
				_ = _453_cf3
				var _454_cf2 _dafny.Int = _451___mcc_h9
				_ = _454_cf2
				var _455_cf1 _dafny.Int = _450___mcc_h8
				_ = _455_cf1
				var _456_v89 bool
				_ = _456_v89
				_456_v89 = false
				_456_v89 = !(!(!(_456_v89)))
				var _457_v90 _dafny.Set
				_ = _457_v90
				_457_v90 = _dafny.SetOf(_456_v89, _456_v89, true)
				_457_v90 = (_457_v90).Union(_457_v90)
				var _458_v91 D0
				_ = _458_v91
				_458_v91 = Companion_D0_.Create_DC0_(p0)
				var _459_v92 _dafny.Map
				_ = _459_v92
				_459_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v89, _333_v16)
				var _460_v93 T0
				_ = _460_v93
				var _nw65 *C0 = New_C0_()
				_ = _nw65
				_nw65.Ctor__(_458_v91, _459_v92, _333_v16, _453_cf3)
				_460_v93 = _nw65
				var _461_v94 _dafny.Array
				_ = _461_v94
				var _nwElement0_13 T0 = _460_v93
				_ = _nwElement0_13
				var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(21))
				_ = _nw66
				(_nw66).ArraySet1(_nwElement0_13, 0)
				(_nw66).ArraySet1(_460_v93, 1)
				(_nw66).ArraySet1(_460_v93, 2)
				(_nw66).ArraySet1(_460_v93, 3)
				(_nw66).ArraySet1(_460_v93, 4)
				(_nw66).ArraySet1(_460_v93, 5)
				(_nw66).ArraySet1(_460_v93, 6)
				(_nw66).ArraySet1(_460_v93, 7)
				(_nw66).ArraySet1(_460_v93, 8)
				(_nw66).ArraySet1(_460_v93, 9)
				(_nw66).ArraySet1(_460_v93, 10)
				(_nw66).ArraySet1(_460_v93, 11)
				(_nw66).ArraySet1(_460_v93, 12)
				(_nw66).ArraySet1(_460_v93, 13)
				(_nw66).ArraySet1(_460_v93, 14)
				(_nw66).ArraySet1(_460_v93, 15)
				(_nw66).ArraySet1(_460_v93, 16)
				(_nw66).ArraySet1(_460_v93, 17)
				(_nw66).ArraySet1(_460_v93, 18)
				(_nw66).ArraySet1(_460_v93, 19)
				(_nw66).ArraySet1(_460_v93, 20)
				_461_v94 = _nw66
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_461_v94), 0))
				_ = _index71
				(_461_v94).ArraySet1(_460_v93, (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_461_v94), 0))
				_ = _index72
				(_461_v94).ArraySet1(_460_v93, (_index72).Int())
				var _462_v95 _dafny.Map
				_ = _462_v95
				_462_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v89, _456_v89)
				var _463_v96 *C0
				_ = _463_v96
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__(_458_v91, _459_v92, (_this).F3(), (_462_v95).Cardinality())
				_463_v96 = _nw67
				var _464_v97 _dafny.Map
				_ = _464_v97
				_464_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_463_v96, (_460_v93).F3())
				var _465_v98 *C0
				_ = _465_v98
				var _nw68 *C0 = New_C0_()
				_ = _nw68
				_nw68.Ctor__(Companion_D0_.Create_DC0_(p0), _459_v92, (func() _dafny.Array {
					if (_464_v97).Contains(_463_v96) {
						return (_464_v97).Get(_463_v96).(_dafny.Array)
					}
					return _333_v16
				})(), _dafny.IntOfInt64(-693))
				_465_v98 = _nw68
			} else if _source7.Is_DC2() {
				var _466___mcc_h11 _dafny.Sequence = _source7.Get_().(D0_DC2).Cf4
				_ = _466___mcc_h11
				var _467___mcc_h12 _dafny.Int = _source7.Get_().(D0_DC2).Cf5
				_ = _467___mcc_h12
				var _468___mcc_h13 _dafny.Int = _source7.Get_().(D0_DC2).Cf6
				_ = _468___mcc_h13
				var _469_cf6 _dafny.Int = _468___mcc_h13
				_ = _469_cf6
				var _470_cf5 _dafny.Int = _467___mcc_h12
				_ = _470_cf5
				var _471_cf4 _dafny.Sequence = _466___mcc_h11
				_ = _471_cf4
				var _472_v99 _dafny.Array
				_ = _472_v99
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_11
				var _nw69 _dafny.Array
				_ = _nw69
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw69 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.MultiSet = func(_473_i17 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetFromSeq((Companion_D2_.Create_DC7_(_dafny.SeqOf(true))).Dtor_cf12())
					}
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw69 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw69).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw69).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_472_v99 = _nw69
				var _474_v100 _dafny.Map
				_ = _474_v100
				_474_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_469_cf6, _470_cf5)), _472_v99)
				_474_v100 = (_474_v100).Update(_469_cf6, _472_v99)
				var _475_v101 _dafny.MultiSet
				_ = _475_v101
				_475_v101 = _dafny.MultiSetOf((func() _dafny.Int {
					if !(p0) {
						return _dafny.IntOfUint32((_471_cf4).Cardinality())
					}
					return _470_cf5
				})(), Companion_Default___.SafeModuloInt(_445_i16, _470_cf5), _469_cf6, (_470_cf5).Plus(p2))
				(globalState).F1 = (_475_v101).Cardinality()
				var _476_v102 bool
				_ = _476_v102
				_476_v102 = true
				var _477_v103 _dafny.Map
				_ = _477_v103
				_477_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _476_v102)
				_476_v102 = !(!(!((_477_v103).Merge(_477_v103)).Equals((_477_v103).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), p0)))))
				var _478_v104 D1
				_ = _478_v104
				_478_v104 = Companion_D1_.Create_DC4_((_this).F3())
				var _479_v105 _dafny.Sequence
				_ = _479_v105
				_479_v105 = _dafny.SeqOf(_476_v102)
				var _480_v106 D1
				_ = _480_v106
				_480_v106 = Companion_D1_.Create_DC5_(false, p0, _dafny.IntOfUint32((_479_v105).Cardinality()))
				_478_v104 = Companion_D1_.Create_DC4_((func() _dafny.Array {
					if (_480_v106).Dtor_cf10() {
						return _333_v16
					}
					return _333_v16
				})())
			} else if _source7.Is_DC0() {
				var _481___mcc_h14 bool = _source7.Get_().(D0_DC0).Cf0
				_ = _481___mcc_h14
				var _482_cf0 bool = _481___mcc_h14
				_ = _482_cf0
				_482_cf0 = true
				(globalState).F1 = _445_i16
				_444_v84 = _444_v84
				var _483_v107 _dafny.Map
				_ = _483_v107
				_483_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_484_i18 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				})), _445_i16)
				_483_v107 = (_483_v107).Update(_444_v84, Companion_Default___.Fm5(p0, globalState))
			} else {
				var _485___mcc_h15 D0 = _source7.Get_().(D0_DC3).Cf7
				_ = _485___mcc_h15
				var _486_cf7 D0 = _485___mcc_h15
				_ = _486_cf7
				var _487_v108 _dafny.Array
				_ = _487_v108
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw70
				_487_v108 = _nw70
				var _488_v109 D1
				_ = _488_v109
				_488_v109 = Companion_D1_.Create_DC4_(_333_v16)
				var _489_v110 _dafny.Sequence
				_ = _489_v110
				_489_v110 = _dafny.SeqOf((_488_v109).Dtor_cf8())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_487_v108), 0))
				_ = _index73
				(_487_v108).ArraySet1(_489_v110, (_index73).Int())
				var _490_v111 bool
				_ = _490_v111
				_490_v111 = false
				var _491_v112 _dafny.Sequence
				_ = _491_v112
				_491_v112 = _dafny.SeqOf((_this).F6(), (_this).F6(), (_this).F6(), (_this).F6(), (_this).F6())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_487_v108), 0))
				_ = _index74
				var _rhs91 _dafny.Int = _dafny.IntOfInt64(957)
				_ = _rhs91
				var _rhs92 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_491_v112, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_491_v112).Cardinality()))).Uint32(), (_this).F6())).Cardinality())
				_ = _rhs92
				var _rhs93 _dafny.Sequence = _489_v110
				_ = _rhs93
				var _rhs94 bool = !((func() bool {
					if !(_490_v111) || (_490_v111) {
						return Companion_Default___.Fm0(_dafny.IntOfUint32((_446_v85).Cardinality()), _445_i16, globalState)
					}
					return _490_v111
				})())
				_ = _rhs94
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				var _lhs55 _dafny.Array = _487_v108
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_487_v108), 0))
				_ = _lhs56
				_lhs53.F1 = _rhs91
				_lhs54.F1 = _rhs92
				(_lhs55).ArraySet1(_rhs93, (_lhs56).Int())
				_490_v111 = _rhs94
				_490_v111 = (_490_v111) == ((Companion_D1_.Create_DC5_(_490_v111, _490_v111, (_442_v82).Cardinality())).Dtor_cf9())
				_490_v111 = _490_v111
				var _492_v113 _dafny.CodePoint
				_ = _492_v113
				_492_v113 = _dafny.CodePoint('j')
				_492_v113 = (_this).F5()
			}
			var _493_v114 D0
			_ = _493_v114
			_493_v114 = Companion_D0_.Create_DC0_(p0)
			var _494_v115 _dafny.Map
			_ = _494_v115
			_494_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _333_v16)
			var _495_v116 _dafny.Sequence
			_ = _495_v116
			_495_v116 = _dafny.SeqOf(p0, p0)
			var _pat_let_tv5 = p0
			_ = _pat_let_tv5
			var _496_v117 *C0
			_ = _496_v117
			var _nw71 *C0 = New_C0_()
			_ = _nw71
			_nw71.Ctor__(func(_pat_let13_0 D0) D0 {
				return func(_497_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let14_0 bool) D0 {
						return func(_498_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_498_dt__update_hcf0_h0)
						}(_pat_let14_0)
					}(_pat_let_tv5)
				}(_pat_let13_0)
			}(_493_v114), ((_494_v115).Update(p0, _333_v16)).Merge((_494_v115).Update((_495_v116).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm5(p0, globalState), _dafny.IntOfUint32((_495_v116).Cardinality()))).Uint32()).(bool), (_this).F3())), (_this).F3(), _dafny.IntOfUint32((_444_v84).Cardinality()))
			_496_v117 = _nw71
			var _499_v118 bool
			_ = _499_v118
			_499_v118 = true
			var _500_v119 _dafny.Sequence
			_ = _500_v119
			_500_v119 = _dafny.SeqOf(_444_v84)
			_499_v118 = _dafny.Companion_Sequence_.Equal((_500_v119).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.IntOfUint32((_500_v119).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_444_v84, _dafny.UnicodeSeqOfUtf8Bytes("hjsywnkxd")))
			var _501_v120 _dafny.Map
			_ = _501_v120
			_501_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _445_i16)
			var _502_v121 _dafny.MultiSet
			_ = _502_v121
			_502_v121 = _dafny.MultiSetOf(_501_v120, Companion_Default___.Fm9(globalState), _501_v120)
			_502_v121 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm10(Companion_D1_.Create_DC6_(), (_444_v84).Select((Companion_Default___.SafeIndex(_445_i16, _dafny.IntOfUint32((_444_v84).Cardinality()))).Uint32()).(_dafny.CodePoint), (_dafny.Zero).Minus(_445_i16), !(_499_v118), globalState), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((Companion_Default___.Fm10(Companion_D1_.Create_DC6_(), (_444_v84).Select((Companion_Default___.SafeIndex(_445_i16, _dafny.IntOfUint32((_444_v84).Cardinality()))).Uint32()).(_dafny.CodePoint), (_dafny.Zero).Minus(_445_i16), !(_499_v118), globalState)).Cardinality()))).Uint32(), _501_v120))
		}
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_333_v16), 0))); ; {
			_guard_loop_2, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _503_i19 _dafny.Int
			_503_i19 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_503_i19).Sign() != -1) && ((_503_i19).Cmp(_dafny.ArrayLen((_333_v16), 0)) < 0)) {
				(_333_v16).ArraySet1(Companion_Default___.SafeModuloInt(_503_i19, (func() _dafny.Int {
					if p0 {
						return (_this).F4()
					}
					return p2
				})()), (_503_i19).Int())
			}
		}
		var _504_v122 _dafny.Sequence
		_ = _504_v122
		_504_v122 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(870))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}(func(_505_i20 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})))
		var _506_v123 D0
		_ = _506_v123
		_506_v123 = Companion_D0_.Create_DC2_(_444_v84, (_this).F4(), (_dafny.Zero).Minus(p2))
		r0 = _dafny.Companion_Sequence_.Concatenate((_504_v122).Select((Companion_Default___.SafeIndex((_this).F4(), _dafny.IntOfUint32((_504_v122).Cardinality()))).Uint32()).(_dafny.Sequence), (_506_v123).Dtor_cf4())
		return r0
	}
}
func (_this *C1) F5() _dafny.CodePoint {
	{
		return _this._f5
	}
}
func (_this *C1) F6() _dafny.Array {
	{
		return _this._f6
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
