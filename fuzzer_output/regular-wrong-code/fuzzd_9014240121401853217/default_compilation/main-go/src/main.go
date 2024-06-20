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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(false, !(true)), true) {
		return ((_dafny.MultiSetOf(_dafny.IntOfInt64(-675))).Cardinality()).Times((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()), _dafny.IntOfInt64(873), _dafny.IntOfInt64(829), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), true)).Cardinality())).Cardinality())
	} else {
		return ((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(431))).Cardinality())).Times(_dafny.IntOfInt64(590))
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, globalState *GlobalState) bool {
	return !(false)
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(914))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-832))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.SetOf(true, false)).IsSubsetOf(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(716), _dafny.IntOfInt64(849))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(716)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(849)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_1_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ja")).Cardinality())), _dafny.IntOfInt64(-249))
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(125), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality())))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true)).Union((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false), true))).Difference(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("khrapsoyo"), _dafny.UnicodeSeqOfUtf8Bytes("l"))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('r')
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Set, p1 _dafny.CodePoint, p2 D2, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(410), _dafny.IntOfInt64(725))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(410)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(725)) < 0) {
				_coll1.Add((_2_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tsnmoy")).Cardinality())))
			}
		}
		return _coll1.ToSet()
	}()).IsSubsetOf(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(291), _dafny.IntOfInt64(256))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _3_v1 _dafny.Int
				_3_v1 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(291)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(256)) < 0) {
					_coll3.Add(Companion_Default___.SafeDivisionInt(_3_v1, (_dafny.SetOf(_dafny.IntOfInt64(-295), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-593), _dafny.IntOfInt64(360))).Cardinality()))).Cardinality()), true)
				}
			}
			return _coll3.ToMap()
		}()).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v2 _dafny.Int
			_4_v2 = interface{}(_compr_2).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(291), _dafny.IntOfInt64(256))); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _3_v1 _dafny.Int
					_3_v1 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(291)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(256)) < 0) {
						_coll4.Add(Companion_Default___.SafeDivisionInt(_3_v1, (_dafny.SetOf(_dafny.IntOfInt64(-295), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-593), _dafny.IntOfInt64(360))).Cardinality()))).Cardinality()), true)
					}
				}
				return _coll4.ToMap()
			}()).Contains(_4_v2) {
				_coll2.Add((_4_v2).Times(_dafny.IntOfInt64(208)))
			}
		}
		return _coll2.ToSet()
	}()), false, !(_dafny.SetOf(_dafny.IntOfInt64(-756), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-709))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_5_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdw")).Cardinality()), _dafny.CodePoint('m'))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(497))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_6_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("smixm")).Cardinality())
		}))).Cardinality()), _dafny.IntOfInt64(338))).Cardinality())
	}))).Cardinality()))).Contains(_dafny.IntOfInt64(478)), true)
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Union((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Difference(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _7_v0 _dafny.Map
				_7_v0 = interface{}(_compr_6).(_dafny.Map)
				if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Contains(_7_v0) {
					_coll6.Add(_7_v0, true)
				}
			}
			return _coll6.ToMap()
		}()).Keys().Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _8_v1 _dafny.Map
			_8_v1 = interface{}(_compr_5).(_dafny.Map)
			if (func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _7_v0 _dafny.Map
					_7_v0 = interface{}(_compr_7).(_dafny.Map)
					if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Contains(_7_v0) {
						_coll7.Add(_7_v0, true)
					}
				}
				return _coll7.ToMap()
			}()).Contains(_8_v1) {
				_coll5.Add(_8_v1)
			}
		}
		return _coll5.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dybqgrbwg"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wxkb"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(756))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-520)
	}))), _dafny.SeqOf(_dafny.IntOfInt64(421), (_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(469)))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.CodePoint {
	var _source0 D13 = Companion_D13_.Create_DC38_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(874), _dafny.IntOfInt64(620))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _11_v0 _dafny.Int
			_11_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(874)).Cmp(_11_v0) <= 0) && ((_11_v0).Cmp(_dafny.IntOfInt64(620)) < 0) {
				_coll8.Add((_11_v0).Times(_dafny.IntOfInt64(751)), _dafny.IntOfInt64(-828))
			}
		}
		return _coll8.ToMap()
	}()).Cardinality()))
	_ = _source0
	if _source0.Is_DC39() {
		if false {
			return _dafny.CodePoint('h')
		} else {
			return _dafny.CodePoint('t')
		}
	} else if _source0.Is_DC40() {
		var _12___mcc_h0 bool = _source0.Get_().(D13_DC40).Cf56
		_ = _12___mcc_h0
		var _13___mcc_h1 *C0 = _source0.Get_().(D13_DC40).Cf57
		_ = _13___mcc_h1
		var _14___mcc_h2 _dafny.Map = _source0.Get_().(D13_DC40).Cf58
		_ = _14___mcc_h2
		var _15_cf58 _dafny.Map = _14___mcc_h2
		_ = _15_cf58
		var _16_cf57 *C0 = _13___mcc_h1
		_ = _16_cf57
		var _17_cf56 bool = _12___mcc_h0
		_ = _17_cf56
		return _dafny.CodePoint('p')
	} else {
		var _18___mcc_h3 _dafny.Map = _source0.Get_().(D13_DC38).Cf55
		_ = _18___mcc_h3
		var _19_cf55 _dafny.Map = _18___mcc_h3
		_ = _19_cf55
		return _dafny.CodePoint('u')
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Difference(_dafny.SetOf(true))).Difference((_dafny.SetOf(false)).Union(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(822)
		}
		return _dafny.IntOfInt64(-43)
	})(), _dafny.CodePoint('t'))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.SetOf(!(false)), _dafny.SetOf(false), _dafny.SetOf(true), _dafny.SetOf(!(false))), _dafny.SeqOf(_dafny.SetOf(true, !(!(true))))) {
		return _dafny.CodePoint('a')
	} else {
		return _dafny.CodePoint('p')
	}
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(791))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.UnicodeSeqOfUtf8Bytes("curdiyset")), _dafny.UnicodeSeqOfUtf8Bytes("o"))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, globalState *GlobalState) _dafny.Map {
	var _source1 D12 = Companion_D12_.Create_DC37_(Companion_D12_.Create_DC36_(true, _dafny.IntOfInt64(-997), false))
	_ = _source1
	if _source1.Is_DC34() {
		var _21___mcc_h0 D0 = _source1.Get_().(D12_DC34).Cf48
		_ = _21___mcc_h0
		var _22___mcc_h1 bool = _source1.Get_().(D12_DC34).Cf49
		_ = _22___mcc_h1
		var _23___mcc_h2 bool = _source1.Get_().(D12_DC34).Cf50
		_ = _23___mcc_h2
		var _24_cf50 bool = _23___mcc_h2
		_ = _24_cf50
		var _25_cf49 bool = _22___mcc_h1
		_ = _25_cf49
		var _26_cf48 D0 = _21___mcc_h0
		_ = _26_cf48
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(26))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality()), _dafny.IntOfInt64(192))).Cardinality()), true)
	} else if _source1.Is_DC35() {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(882), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(144), true))
	} else if _source1.Is_DC36() {
		var _28___mcc_h3 bool = _source1.Get_().(D12_DC36).Cf51
		_ = _28___mcc_h3
		var _29___mcc_h4 _dafny.Int = _source1.Get_().(D12_DC36).Cf52
		_ = _29___mcc_h4
		var _30___mcc_h5 bool = _source1.Get_().(D12_DC36).Cf53
		_ = _30___mcc_h5
		var _31_cf53 bool = _30___mcc_h5
		_ = _31_cf53
		var _32_cf52 _dafny.Int = _29___mcc_h4
		_ = _32_cf52
		var _33_cf51 bool = _28___mcc_h3
		_ = _33_cf51
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_cf52, _33_cf51)
	} else if _source1.Is_DC33() {
		var _34___mcc_h6 _dafny.Set = _source1.Get_().(D12_DC33).Cf47
		_ = _34___mcc_h6
		var _35_cf47 _dafny.Set = _34___mcc_h6
		_ = _35_cf47
		return (func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(997), _dafny.IntOfInt64(164))); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _36_v0 _dafny.Int
				_36_v0 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(997)).Cmp(_36_v0) <= 0) && ((_36_v0).Cmp(_dafny.IntOfInt64(164)) < 0) {
					_coll9.Add((_36_v0).Minus(_dafny.IntOfInt64(492)), false)
				}
			}
			return _coll9.ToMap()
		}()).Merge(func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(766), _dafny.IntOfInt64(670))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _37_v1 _dafny.Int
				_37_v1 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(766)).Cmp(_37_v1) <= 0) && ((_37_v1).Cmp(_dafny.IntOfInt64(670)) < 0) {
					_coll10.Add((_37_v1).Times((_35_cf47).Cardinality()), !(!(true)))
				}
			}
			return _coll10.ToMap()
		}())
	} else {
		var _38___mcc_h7 D12 = _source1.Get_().(D12_DC37).Cf54
		_ = _38___mcc_h7
		var _39_cf54 D12 = _38___mcc_h7
		_ = _39_cf54
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("slrpor")).Cardinality()), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(765), false))
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qgn")).Cardinality()))).Difference((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-41))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfInt64(458))).Difference(_dafny.MultiSetOf((func() _dafny.Set {
		var _coll11 = _dafny.NewBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())), (_dafny.SetOf(false, false, false, false)).Cardinality()), true)).Keys().Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _40_v0 _dafny.Map
			_40_v0 = interface{}(_compr_11).(_dafny.Map)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())), (_dafny.SetOf(false, false, false, false)).Cardinality()), true)).Contains(_40_v0) {
				_coll11.Add(_40_v0)
			}
		}
		return _coll11.ToSet()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)) || ((Companion_D9_.Create_DC24_(!(!(true)), false, true)).Dtor_cf35()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(855)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true), _dafny.SeqOf((Companion_D12_.Create_DC34_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false)), false, true)).Dtor_cf50(), true, !(true), false)))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) D2 {
	var _source2 D3 = Companion_D3_.Create_DC10_((Companion_D3_.Create_DC10_(Companion_D3_.Create_DC9_(_dafny.IntOfInt64(528)))).Dtor_cf12())
	_ = _source2
	if _source2.Is_DC9() {
		var _41___mcc_h0 _dafny.Int = _source2.Get_().(D3_DC9).Cf11
		_ = _41___mcc_h0
		var _42_cf11 _dafny.Int = _41___mcc_h0
		_ = _42_cf11
		return Companion_D2_.Create_DC7_(_dafny.IntOfInt64(-394), func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true))).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _43_v0 _dafny.Map
				_43_v0 = interface{}(_compr_12).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)), _43_v0) {
					_coll12.Add(_43_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_cf11, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true, true, true)).Cardinality())))).Cardinality())
				}
			}
			return _coll12.ToMap()
		}(), _dafny.IntOfInt64(890))
	} else if _source2.Is_DC8() {
		var _44___mcc_h1 _dafny.Sequence = _source2.Get_().(D3_DC8).Cf10
		_ = _44___mcc_h1
		var _45_cf10 _dafny.Sequence = _44___mcc_h1
		_ = _45_cf10
		return Companion_D2_.Create_DC7_(_dafny.IntOfInt64(338), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.IntOfInt64(-77)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(373))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_46_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality()))
	} else {
		var _47___mcc_h2 D3 = _source2.Get_().(D3_DC10).Cf12
		_ = _47___mcc_h2
		var _48_cf12 D3 = _47___mcc_h2
		_ = _48_cf12
		return Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(115))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_49_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		}))).Cardinality()), func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _50_v1 _dafny.Map
				_50_v1 = interface{}(_compr_13).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)), _50_v1) {
					_coll13.Add(_50_v1, _dafny.IntOfInt64(762))
				}
			}
			return _coll13.ToMap()
		}(), _dafny.IntOfInt64(107))
	}
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Sequence {
	if !(false) || (!(false)) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("j"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(455))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_51_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))
	} else {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(838))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_52_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('r'))
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.CodePoint {
	if !(!((!(false)) && (true))) {
		return _dafny.CodePoint('i')
	} else {
		return _dafny.CodePoint('n')
	}
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Set {
	var _source3 D11 = (func() D11 {
		if false {
			return Companion_D11_.Create_DC30_()
		}
		return Companion_D11_.Create_DC30_()
	})()
	_ = _source3
	if _source3.Is_DC30() {
		return _dafny.SetOf(_dafny.IntOfInt64(-94), _dafny.IntOfInt64(271))
	} else if _source3.Is_DC31() {
		var _53___mcc_h0 _dafny.Int = _source3.Get_().(D11_DC31).Cf44
		_ = _53___mcc_h0
		var _54___mcc_h1 _dafny.Int = _source3.Get_().(D11_DC31).Cf45
		_ = _54___mcc_h1
		var _55_cf45 _dafny.Int = _54___mcc_h1
		_ = _55_cf45
		var _56_cf44 _dafny.Int = _53___mcc_h0
		_ = _56_cf44
		return (_dafny.SetOf(_56_cf44)).Intersection(_dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())))
	} else if _source3.Is_DC29() {
		var _57___mcc_h2 T0 = _source3.Get_().(D11_DC29).Cf43
		_ = _57___mcc_h2
		var _58_cf43 T0 = _57___mcc_h2
		_ = _58_cf43
		return (_dafny.SetOf(_dafny.IntOfInt64(413))).Union(_dafny.SetOf(_dafny.IntOfInt64(-448), (func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(591), _dafny.IntOfInt64(801))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _59_v0 _dafny.Int
				_59_v0 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(591)).Cmp(_59_v0) <= 0) && ((_59_v0).Cmp(_dafny.IntOfInt64(801)) < 0) {
					_coll14.Add((_59_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aydffd")).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ikjyjyifl")).Cardinality()))
				}
			}
			return _coll14.ToMap()
		}()).Cardinality()))
	} else {
		var _60___mcc_h3 D11 = _source3.Get_().(D11_DC32).Cf46
		_ = _60___mcc_h3
		var _61_cf46 D11 = _60___mcc_h3
		_ = _61_cf46
		return (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ba")).Cardinality()))).Intersection(_dafny.SetOf((_dafny.SetOf(false, false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ryykijq")).Cardinality()), _dafny.CodePoint('t'))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-475)))).Cardinality(), false))), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jgwt")).Cardinality()), true)))), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-981))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_62_i0 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_63_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			}))).Cardinality()))).Cardinality()), !(false))).Cardinality()
		}))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _64_v0 _dafny.Int
			_64_v0 = interface{}(_compr_15).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-981))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_62_i0 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}(func(_63_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))).Cardinality()))).Cardinality()), !(false))).Cardinality()
			})), _64_v0) {
				_coll15.Add((_64_v0).Plus(_dafny.IntOfInt64(806)), true)
			}
		}
		return _coll15.ToMap()
	}()))))).Intersection(_dafny.MultiSetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), false)))))).Intersection((_dafny.MultiSetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(true))), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(false)), Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(true))))).Intersection(_dafny.MultiSetOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(true))))))))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D3_.Create_DC9_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(!(true))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(754))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_65_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality()), false)).Cardinality())).Cardinality()))), Companion_D3_.Create_DC9_((_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfInt64(678), _dafny.IntOfInt64(813), (_dafny.SetOf(_dafny.IntOfInt64(973), _dafny.IntOfInt64(62), (_dafny.Zero).Minus(_dafny.IntOfInt64(-221)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())).Cardinality()))), _dafny.SeqOf(Companion_D3_.Create_DC9_(_dafny.IntOfInt64(-407))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pptit")).Cardinality()), Companion_D8_.Create_DC21_(false, false, _dafny.SeqOf(false, true), true, _dafny.IntOfInt64(249)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-630), _dafny.IntOfInt64(-844))).Cardinality(), Companion_D8_.Create_DC21_(true, true, _dafny.SeqOf(false), !(false), _dafny.IntOfInt64(-398))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(123), Companion_D8_.Create_DC21_(true, true, _dafny.SeqOf(false), false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(48))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ska")).Cardinality()), Companion_D8_.Create_DC21_(true, !(!(false)), _dafny.SeqOf(false, true, false, false, true), true, _dafny.IntOfInt64(-541)))))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, true))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(!(true), !(false), false, false, !(true)), _dafny.IntOfInt64(299)))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) {
	var _66_v0 _dafny.Map
	_ = _66_v0
	_66_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.MultiSetOf(p0, p0, p0, false, p0))
	var _67_v1 D0
	_ = _67_v1
	_67_v1 = Companion_D0_.Create_DC1_(p0)
	var _68_v2 _dafny.MultiSet
	_ = _68_v2
	_68_v2 = _dafny.MultiSetOf(p0, (_67_v1).Dtor_cf1())
	var _69_v3 D10
	_ = _69_v3
	_69_v3 = Companion_D10_.Create_DC27_(p0, _dafny.IntOfInt64(921), p1, (func() _dafny.MultiSet {
		if (_66_v0).Contains(!(!(true))) {
			return (_66_v0).Get(!(!(true))).(_dafny.MultiSet)
		}
		return _68_v2
	})())
	var _70_v4 _dafny.Sequence
	_ = _70_v4
	_70_v4 = _dafny.SeqOf(p0)
	var _71_v5 _dafny.Set
	_ = _71_v5
	_71_v5 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_70_v4, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_70_v4).Cardinality()))).Uint32(), !(p0))).Cardinality())), p2, p2, _dafny.IntOfInt64(947))
	if (_71_v5).Contains(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((func() _dafny.Int {
		if (p3).Contains(true) {
			return (p3).Get(true).(_dafny.Int)
		}
		return ((_69_v3).Dtor_cf41()).Cardinality()
	})()), p2)) {
		(globalState).F0 = p0
		var _72_v6 *C2
		_ = _72_v6
		var _nw0 *C2 = New_C2_()
		_ = _nw0
		_nw0.Ctor__(p0)
		_72_v6 = _nw0
		_72_v6 = _72_v6
		var _73_v7 _dafny.Sequence
		_ = _73_v7
		_73_v7 = _dafny.UnicodeSeqOfUtf8Bytes("osmkakap")
		_73_v7 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_74_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_75_i0 _dafny.Int) _dafny.CodePoint {
				return _74_p1
			}
		})(p1))), _dafny.UnicodeSeqOfUtf8Bytes("baefwmqnu"))
		var _76_v8 _dafny.MultiSet
		_ = _76_v8
		_76_v8 = _dafny.MultiSetOf(p1, p1)
		var _77_v9 _dafny.Array
		_ = _77_v9
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_78_v6 *C2) func(_dafny.Int) bool {
				return func(_79_i1 _dafny.Int) bool {
					return !((_78_v6).F13())
				}
			})(_72_v6)
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
		_77_v9 = _nw1
		var _80_v10 *C0
		_ = _80_v10
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__((p2).Cmp(p2) != 0, (Companion_D4_.Create_DC13_((_76_v8).Cardinality(), _77_v9)).Dtor_cf20())
		_80_v10 = _nw2
		var _81_v11 _dafny.Array
		_ = _81_v11
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw3
		_81_v11 = _nw3
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_81_v11), 0))
		_ = _index0
		(_81_v11).ArraySet1(Companion_Default___.Fm14(_73_v7, p2, globalState), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_81_v11), 0))
		_ = _index1
		(_81_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_73_v7, _73_v7), (_index1).Int())
	} else {
		var _82_v12 _dafny.Map
		_ = _82_v12
		_82_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !((_69_v3).Dtor_cf38()) || (p0))
		_82_v12 = _82_v12
		var _83_v13 _dafny.Sequence
		_ = _83_v13
		_83_v13 = _dafny.SeqOf(p2)
		var _84_v14 _dafny.Sequence
		_ = _84_v14
		_84_v14 = _dafny.UnicodeSeqOfUtf8Bytes("yxwm")
		var _rhs0 _dafny.Sequence = _83_v13
		_ = _rhs0
		var _rhs1 bool = !(p0)
		_ = _rhs1
		var _rhs2 bool = p0
		_ = _rhs2
		var _rhs3 _dafny.Int = _dafny.IntOfUint32((_84_v14).Cardinality())
		_ = _rhs3
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		_83_v13 = _rhs0
		_lhs0.F0 = _rhs1
		_lhs1.F0 = _rhs2
		_lhs2.F5 = _rhs3
		var _85_v15 _dafny.Map
		_ = _85_v15
		_85_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm0(_dafny.IntOfUint32((_84_v14).Cardinality()), p2, _dafny.IntOfUint32((_84_v14).Cardinality()), p0, globalState)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_86_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_87_i2 _dafny.Int) _dafny.Int {
				return _86_p2
			}
		})(p2)))).Cardinality()))
		_85_v15 = (_85_v15).Update((_dafny.SetOf(false, p0)).Cardinality(), (_82_v12).Cardinality())
		var _88_v16 _dafny.Array
		_ = _88_v16
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_1
		var _nw4 _dafny.Array
		_ = _nw4
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw4 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_89_p0 bool) func(_dafny.Int) bool {
				return func(_90_i3 _dafny.Int) bool {
					return _89_p0
				}
			})(p0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw4 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw4).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw4).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_88_v16 = _nw4
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_88_v16), 0))
		_ = _index2
		(_88_v16).ArraySet1(p0, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_88_v16), 0))
		_ = _index3
		(_88_v16).ArraySet1((p2).Cmp(p2) > 0, (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_88_v16), 0))
		_ = _index4
		(_88_v16).ArraySet1(!_dafny.Companion_Sequence_.Contains(_70_v4, p0), (_index4).Int())
	}
	(globalState).F2 = p0
	if Companion_Default___.Fm1(p0, globalState) {
		var _91_v17 _dafny.Array
		_ = _91_v17
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw5
		_91_v17 = _nw5
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))
		_ = _index5
		(_91_v17).ArraySet1(false, (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))
		_ = _index6
		(_91_v17).ArraySet1(!(false), (_index6).Int())
		if p0 {
			var _92_v18 _dafny.Array
			_ = _92_v18
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_2
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_93_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_94_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_94_i4, _93_p2)
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
			_92_v18 = _nw6
			var _95_v19 _dafny.Map
			_ = _95_v19
			_95_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _92_v18)
			var _96_v20 _dafny.Map
			_ = _96_v20
			_96_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
			var _97_v21 _dafny.Sequence
			_ = _97_v21
			_97_v21 = _dafny.UnicodeSeqOfUtf8Bytes("ikfpgwiid")
			var _98_v22 T1
			_ = _98_v22
			var _nw7 *C3 = New_C3_()
			_ = _nw7
			_nw7.Ctor__(p2)
			_98_v22 = _nw7
			var _99_v23 _dafny.Map
			_ = _99_v23
			_99_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v22, p2)
			var _100_v24 *C2
			_ = _100_v24
			var _nw8 *C2 = New_C2_()
			_ = _nw8
			_nw8.Ctor__(p0)
			_100_v24 = _nw8
			var _101_v25 _dafny.Map
			_ = _101_v25
			_101_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _100_v24)
			var _102_v26 _dafny.Map
			_ = _102_v26
			_102_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), (_101_v25).Cardinality())
			var _103_v27 _dafny.Map
			_ = _103_v27
			_103_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v17, _dafny.IntOfInt64(806))
			var _104_v28 _dafny.Set
			_ = _104_v28
			_104_v28 = _dafny.SetOf((_91_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))).Int()).(bool), p0)
			var _105_v29 _dafny.Array
			_ = _105_v29
			var _nwElement0_0 _dafny.Int = p2
			_ = _nwElement0_0
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(21))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_0, 0)
			(_nw9).ArraySet1(p2, 1)
			(_nw9).ArraySet1((p2).Minus((_95_v19).Cardinality()), 2)
			(_nw9).ArraySet1(p2, 3)
			(_nw9).ArraySet1(p2, 4)
			(_nw9).ArraySet1(p2, 5)
			(_nw9).ArraySet1(p2, 6)
			(_nw9).ArraySet1(p2, 7)
			(_nw9).ArraySet1((func() _dafny.Int {
				if (p3).Contains(p0) {
					return (p3).Get(p0).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(Companion_Default___.Fm0((_96_v20).Cardinality(), p2, _dafny.IntOfUint32((_97_v21).Cardinality()), p0, globalState))
			})(), 8)
			(_nw9).ArraySet1((_99_v23).Cardinality(), 9)
			(_nw9).ArraySet1((_dafny.Zero).Minus(p2), 10)
			(_nw9).ArraySet1(p2, 11)
			(_nw9).ArraySet1(p2, 12)
			(_nw9).ArraySet1(Companion_Default___.SafeDivisionInt(p2, (_dafny.Zero).Minus(p2)), 13)
			(_nw9).ArraySet1((func() _dafny.Int {
				if (_102_v26).Contains(p1) {
					return (_102_v26).Get(p1).(_dafny.Int)
				}
				return p2
			})(), 14)
			(_nw9).ArraySet1(p2, 15)
			(_nw9).ArraySet1((_dafny.Zero).Minus(p2), 16)
			(_nw9).ArraySet1(p2, 17)
			(_nw9).ArraySet1((_dafny.Zero).Minus(((_103_v27).Merge(_103_v27)).Cardinality()), 18)
			(_nw9).ArraySet1((_104_v28).Cardinality(), 19)
			(_nw9).ArraySet1(p2, 20)
			_105_v29 = _nw9
			var _106_v30 *C0
			_ = _106_v30
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__((_91_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))).Int()).(bool), _91_v17)
			_106_v30 = _nw10
			var _107_v31 _dafny.Set
			_ = _107_v31
			_107_v31 = _dafny.SetOf(_106_v30)
			var _108_v32 _dafny.Map
			_ = _108_v32
			_108_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(p2, (_107_v31).Cardinality())).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_109_i5 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pbcudgtk")).Cardinality())
			})))
			var _110_v33 _dafny.Sequence
			_ = _110_v33
			_110_v33 = _dafny.SeqOf(p2)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))
			_ = _index7
			(_105_v29).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_108_v32).Contains(p2) {
					return (_108_v32).Get(p2).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_dafny.IntOfInt64(552))
			})(), _110_v33)).Cardinality()), (_index7).Int())
			var _111_v34 _dafny.Array
			_ = _111_v34
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_3
			var _nw11 _dafny.Array
			_ = _nw11
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw11 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_112_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_113_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg19 _dafny.Int) interface{} {
								return coer19(arg19)
							}
						}((func(_114_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_115_i7 _dafny.Int) _dafny.CodePoint {
								return _114_p1
							}
						})(_112_p1)))
					}
				})(p1)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw11 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw11).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw11).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_111_v34 = _nw11
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))
			_ = _index8
			var _rhs4 _dafny.Int = p2
			_ = _rhs4
			var _rhs5 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("agjcjgg"), _dafny.Companion_Sequence_.Concatenate(_97_v21, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(944))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_116_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_117_i8 _dafny.Int) _dafny.CodePoint {
					return _116_p1
				}
			})(p1)))))
			_ = _rhs5
			var _rhs6 _dafny.Int = p2
			_ = _rhs6
			var _rhs7 _dafny.Array = _111_v34
			_ = _rhs7
			var _lhs3 _dafny.Array = _105_v29
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			(_lhs3).ArraySet1(_rhs4, (_lhs4).Int())
			_lhs5.F2 = _rhs5
			_lhs6.F5 = _rhs6
			_111_v34 = _rhs7
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))
			_ = _index9
			var _rhs8 bool = p0
			_ = _rhs8
			var _rhs9 _dafny.Int = Companion_Default___.Fm0((_105_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))).Int()).(_dafny.Int), (_105_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))).Int()).(_dafny.Int), p2, p0, globalState)
			_ = _rhs9
			var _rhs10 _dafny.Int = ((_105_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))).Int()).(_dafny.Int)).Times((_105_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))).Int()).(_dafny.Int))
			_ = _rhs10
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			var _lhs9 _dafny.Array = _105_v29
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))
			_ = _lhs10
			_lhs7.F2 = _rhs8
			_lhs8.F5 = _rhs9
			(_lhs9).ArraySet1(_rhs10, (_lhs10).Int())
			var _118_v35 *C4
			_ = _118_v35
			var _nw12 *C4 = New_C4_()
			_ = _nw12
			_nw12.Ctor__((_105_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_105_v29), 0))).Int()).(_dafny.Int))
			_118_v35 = _nw12
			var _119_v36 _dafny.Map
			_ = _119_v36
			_119_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _118_v35)
			var _120_v37 _dafny.Sequence
			_ = _120_v37
			_120_v37 = _dafny.SeqOf(_118_v35, (func() *C4 {
				if (_119_v36).Contains((_104_v28).Cardinality()) {
					return (_119_v36).Get((_104_v28).Cardinality()).(*C4)
				}
				return _118_v35
			})(), _118_v35)
			_120_v37 = _dafny.Companion_Sequence_.Concatenate(_120_v37, _dafny.Companion_Sequence_.Update(_120_v37, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_120_v37).Cardinality()))).Uint32(), _118_v35))
			var _121_v38 _dafny.CodePoint
			_ = _121_v38
			_121_v38 = _dafny.CodePoint('s')
			_121_v38 = (func() _dafny.CodePoint {
				if (_100_v24).F13() {
					return p1
				}
				return p1
			})()
			(globalState).F0 = !(p3).Contains(true)
		} else {
			var _122_v39 _dafny.Array
			_ = _122_v39
			var _nw13 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
			_ = _nw13
			_122_v39 = _nw13
			var _123_v40 _dafny.Sequence
			_ = _123_v40
			_123_v40 = _dafny.SeqOf(p2)
			var _124_v41 *C2
			_ = _124_v41
			var _nw14 *C2 = New_C2_()
			_ = _nw14
			_nw14.Ctor__((_70_v4).Select((Companion_Default___.SafeIndex((_123_v40).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_123_v40).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_70_v4).Cardinality()))).Uint32()).(bool))
			_124_v41 = _nw14
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_122_v39), 0))
			_ = _index10
			(_122_v39).ArraySet1(_124_v41, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_122_v39), 0))
			_ = _index11
			var _nw15 *C2 = New_C2_()
			_ = _nw15
			_nw15.Ctor__((_124_v41).F13())
			(_122_v39).ArraySet1(_nw15, (_index11).Int())
			(globalState).F2 = p0
			var _125_v42 T0
			_ = _125_v42
			var _nw16 *C2 = New_C2_()
			_ = _nw16
			_nw16.Ctor__(p0)
			_125_v42 = _nw16
			var _126_v43 _dafny.Map
			_ = _126_v43
			_126_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_125_v42, p2)
			var _127_v44 _dafny.Map
			_ = _127_v44
			_127_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v43).Cardinality(), (_91_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))).Int()).(bool))
			var _128_v45 D0
			_ = _128_v45
			_128_v45 = Companion_D0_.Create_DC0_(_127_v44)
			_68_v2 = _dafny.MultiSetOf(!((false) && ((_91_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))).Int()).(bool))), !((_124_v41).Fm4(_128_v45, _128_v45, _68_v2, globalState)))
			var _129_v46 _dafny.Sequence
			_ = _129_v46
			_129_v46 = _dafny.UnicodeSeqOfUtf8Bytes("colybgq")
			var _130_v47 _dafny.Map
			_ = _130_v47
			_130_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v46, _dafny.Companion_Sequence_.Concatenate(_123_v40, _dafny.SeqOf(_dafny.IntOfInt64(652))))
			var _131_v48 _dafny.Sequence
			_ = _131_v48
			_131_v48 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(131))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_132_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_133_i9 _dafny.Int) _dafny.Int {
					return _132_p2
				}
			})(p2))))
			_130_v47 = (_130_v47).Update(_dafny.UnicodeSeqOfUtf8Bytes("tqy"), (_131_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-912), _dafny.IntOfUint32((_131_v48).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _134_v49 *C3
			_ = _134_v49
			var _nw17 *C3 = New_C3_()
			_ = _nw17
			_nw17.Ctor__(p2)
			_134_v49 = _nw17
		}
		var _135_v50 *C1
		_ = _135_v50
		var _nw18 *C1 = New_C1_()
		_ = _nw18
		_nw18.Ctor__()
		_135_v50 = _nw18
		(globalState).F0 = (_68_v2).IsSubsetOf(_68_v2)
		var _136_v51 D7
		_ = _136_v51
		_136_v51 = Companion_D7_.Create_DC18_()
		var _137_v52 _dafny.Sequence
		_ = _137_v52
		_137_v52 = _dafny.SeqOf(_136_v51)
		var _138_v53 _dafny.Sequence
		_ = _138_v53
		_138_v53 = _dafny.SeqOf(_137_v52)
		var _139_v54 _dafny.Map
		_ = _139_v54
		_139_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_91_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))).Int()).(bool), _138_v53)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_91_v17), 0))
		_ = _index12
		(_91_v17).ArraySet1((_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_139_v54).Contains(p0) {
				return (_139_v54).Get(p0).(_dafny.Sequence)
			}
			return _138_v53
		})()).Cardinality())).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
			if true {
				return _dafny.IntOfInt64(-279)
			}
			return p2
		})())) <= 0, (_index12).Int())
	} else {
		var _140_v55 _dafny.Map
		_ = _140_v55
		_140_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _141_v56 *C4
		_ = _141_v56
		var _nw19 *C4 = New_C4_()
		_ = _nw19
		_nw19.Ctor__((_140_v55).Cardinality())
		_141_v56 = _nw19
		var _142_v57 _dafny.Array
		_ = _142_v57
		var _nwElement0_1 *C4 = _141_v56
		_ = _nwElement0_1
		var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(15))
		_ = _nw20
		(_nw20).ArraySet1(_nwElement0_1, 0)
		(_nw20).ArraySet1(_141_v56, 1)
		(_nw20).ArraySet1(_141_v56, 2)
		(_nw20).ArraySet1((func() *C4 {
			if !(p0) {
				return _141_v56
			}
			return _141_v56
		})(), 3)
		(_nw20).ArraySet1(_141_v56, 4)
		(_nw20).ArraySet1(_141_v56, 5)
		(_nw20).ArraySet1(_141_v56, 6)
		(_nw20).ArraySet1(_141_v56, 7)
		(_nw20).ArraySet1(_141_v56, 8)
		(_nw20).ArraySet1(_141_v56, 9)
		(_nw20).ArraySet1(_141_v56, 10)
		(_nw20).ArraySet1(_141_v56, 11)
		(_nw20).ArraySet1(_141_v56, 12)
		(_nw20).ArraySet1((func() *C4 {
			if p0 {
				return _141_v56
			}
			return _141_v56
		})(), 13)
		(_nw20).ArraySet1(_141_v56, 14)
		_142_v57 = _nw20
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_142_v57), 0))
		_ = _index13
		(_142_v57).ArraySet1(_141_v56, (_index13).Int())
		var _143_v58 _dafny.Sequence
		_ = _143_v58
		_143_v58 = _dafny.UnicodeSeqOfUtf8Bytes("tlsst")
		var _144_v59 _dafny.Map
		_ = _144_v59
		_144_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_143_v58).Cardinality()))
		var _145_v60 _dafny.Array
		_ = _145_v60
		var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw21
		_145_v60 = _nw21
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_145_v60), 0))
		_ = _index14
		(_145_v60).ArraySet1((_dafny.Zero).Minus((_141_v56).F9()), (_index14).Int())
		var _146_v61 _dafny.Array
		_ = _146_v61
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_4
		var _nw22 _dafny.Array
		_ = _nw22
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw22 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) bool = (func(_147_p0 bool) func(_dafny.Int) bool {
				return func(_148_i10 _dafny.Int) bool {
					return _147_p0
				}
			})(p0)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw22 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw22).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw22).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_146_v61 = _nw22
		var _149_v62 D4
		_ = _149_v62
		_149_v62 = Companion_D4_.Create_DC13_(p2, _146_v61)
		var _150_v63 *C4
		_ = _150_v63
		_150_v63 = _141_v56
		var _151_v64 *C4
		_ = _151_v64
		_151_v64 = (_150_v63)
		var _152_v65 _dafny.Map
		_ = _152_v65
		_152_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('y'))
		var _153_v66 D2
		_ = _153_v66
		_153_v66 = Companion_D2_.Create_DC6_(p0)
		var _154_v67 *C2
		_ = _154_v67
		var _nw23 *C2 = New_C2_()
		_ = _nw23
		_nw23.Ctor__(p0)
		_154_v67 = _nw23
		var _155_v68 _dafny.Set
		_ = _155_v68
		_155_v68 = _dafny.SetOf(_154_v67, _154_v67, _154_v67)
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_142_v57), 0))
		_ = _index15
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_145_v60), 0))
		_ = _index16
		var _rhs11 *C4 = (_150_v63)
		_ = _rhs11
		var _rhs12 _dafny.Map = _144_v59
		_ = _rhs12
		var _rhs13 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm8(p0, _152_v65, (_153_v66).Dtor_cf6(), globalState)).Cardinality())
		_ = _rhs13
		var _rhs14 D4 = _149_v62
		_ = _rhs14
		var _rhs15 bool = ((_155_v68).IsDisjointFrom(_155_v68)) && ((_154_v67).F13())
		_ = _rhs15
		var _lhs11 _dafny.Array = _142_v57
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.ArrayLen((_142_v57), 0))
		_ = _lhs12
		var _lhs13 _dafny.Array = _145_v60
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_145_v60), 0))
		_ = _lhs14
		var _lhs15 *GlobalState = globalState
		_ = _lhs15
		(_lhs11).ArraySet1(_rhs11, (_lhs12).Int())
		_144_v59 = _rhs12
		(_lhs13).ArraySet1(_rhs13, (_lhs14).Int())
		_149_v62 = _rhs14
		_lhs15.F2 = _rhs15
		var _156_v69 *C0
		_ = _156_v69
		var _nw24 *C0 = New_C0_()
		_ = _nw24
		_nw24.Ctor__(p0, _146_v61)
		_156_v69 = _nw24
		var _157_v70 T0
		_ = _157_v70
		var _nw25 *C1 = New_C1_()
		_ = _nw25
		_nw25.Ctor__()
		_157_v70 = _nw25
		var _158_v71 _dafny.Set
		_ = _158_v71
		_158_v71 = _dafny.SetOf(_157_v70, _157_v70, _157_v70, _157_v70)
		var _159_v72 _dafny.Sequence
		_ = _159_v72
		_159_v72 = _dafny.SeqOf((_158_v71).Cardinality(), p2)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_145_v60), 0))
		_ = _index17
		(_145_v60).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_141_v56).F9(), p2), _dafny.IntOfUint32((_159_v72).Cardinality()))), (_index17).Int())
		_140_v55 = _140_v55
		var _160_v74 _dafny.Sequence
		_ = _160_v74
		_160_v74 = _dafny.SeqOf(_dafny.SeqOf((_156_v69).F10(), (_154_v67).F13()))
		var _161_v75 _dafny.Map
		_ = _161_v75
		_161_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((_160_v74).Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _162_v73 _dafny.Sequence
				_162_v73 = interface{}(_compr_16).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_160_v74, _162_v73) {
					_coll16.Add(_162_v73, _dafny.IntOfUint32((_159_v72).Cardinality()))
				}
			}
			return _coll16.ToMap()
		}(), _145_v60)
		var _163_v76 _dafny.Map
		_ = _163_v76
		_163_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v4, (_145_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_145_v60), 0))).Int()).(_dafny.Int))
		_145_v60 = (func() _dafny.Array {
			if (_161_v75).Contains((_163_v76).Merge(_163_v76)) {
				return (_161_v75).Get((_163_v76).Merge(_163_v76)).(_dafny.Array)
			}
			return _145_v60
		})()
	}
	var _164_v77 _dafny.Array
	_ = _164_v77
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(20)
	_ = _len0_5
	var _nw26 _dafny.Array
	_ = _nw26
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw26 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Int = (func(_165_p0 bool) func(_dafny.Int) _dafny.Int {
			return func(_166_i11 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_166_i11, (_dafny.SetOf(_165_p0)).Cardinality())
			}
		})(p0)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw26 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw26).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw26).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_164_v77 = _nw26
	var _167_v78 _dafny.Map
	_ = _167_v78
	_167_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _164_v77)
	(globalState).F5 = Companion_Default___.SafeModuloInt((p2).Times((_167_v78).Cardinality()), Companion_Default___.Fm0(p2, Companion_Default___.Fm0(_dafny.IntOfInt64(818), p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dmpivcli")).Cardinality()), p0, globalState), p2, p0, globalState))
	var _168_v79 *C4
	_ = _168_v79
	var _nw27 *C4 = New_C4_()
	_ = _nw27
	_nw27.Ctor__((p2).Minus(p2))
	_168_v79 = _nw27
	(globalState).F0 = false
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _169_v0 _dafny.Int
	_ = _169_v0
	_169_v0 = _dafny.IntOfInt64(292)
	var _170_v1 _dafny.Sequence
	_ = _170_v1
	_170_v1 = _dafny.SeqOf(_169_v0, _169_v0)
	var _171_v2 bool
	_ = _171_v2
	_171_v2 = false
	var _172_v3 _dafny.Array
	_ = _172_v3
	var _nwElement0_2 _dafny.Int = _169_v0
	_ = _nwElement0_2
	var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(3))
	_ = _nw28
	(_nw28).ArraySet1(_nwElement0_2, 0)
	(_nw28).ArraySet1((_170_v1).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _171_v2)).Cardinality(), _dafny.IntOfUint32((_170_v1).Cardinality()))).Uint32()).(_dafny.Int), 1)
	(_nw28).ArraySet1(_169_v0, 2)
	_172_v3 = _nw28
	var _173_v4 _dafny.MultiSet
	_ = _173_v4
	_173_v4 = _dafny.MultiSetOf(_172_v3)
	var _174_v5 _dafny.Sequence
	_ = _174_v5
	_174_v5 = _dafny.SeqOf(true)
	var _175_v6 _dafny.Map
	_ = _175_v6
	_175_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_174_v5).Cardinality()), _171_v2)
	var _176_v7 _dafny.Map
	_ = _176_v7
	_176_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v0, _dafny.MultiSetOf(_175_v6))
	var _177_v9 D0
	_ = _177_v9
	_177_v9 = Companion_D0_.Create_DC0_(func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(388), _dafny.IntOfInt64(693))); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _178_v8 _dafny.Int
			_178_v8 = interface{}(_compr_17).(_dafny.Int)
			if ((_dafny.IntOfInt64(388)).Cmp(_178_v8) <= 0) && ((_178_v8).Cmp(_dafny.IntOfInt64(693)) < 0) {
				_coll17.Add(Companion_Default___.SafeDivisionInt(_178_v8, _169_v0), _171_v2)
			}
		}
		return _coll17.ToMap()
	}())
	var _179_v10 _dafny.MultiSet
	_ = _179_v10
	_179_v10 = _dafny.MultiSetOf(_175_v6, (_177_v9).Dtor_cf0())
	var _180_v11 _dafny.Set
	_ = _180_v11
	_180_v11 = _dafny.SetOf(_171_v2, _171_v2)
	var _181_globalState *GlobalState
	_ = _181_globalState
	var _nw29 *GlobalState = New_GlobalState_()
	_ = _nw29
	_nw29.Ctor__(true, _dafny.IntOfInt64(-616), false, _173_v4, true, _dafny.IntOfInt64(663), ((func() _dafny.MultiSet {
		if (_176_v7).Contains(_169_v0) {
			return (_176_v7).Get(_169_v0).(_dafny.MultiSet)
		}
		return _179_v10
	})()).Union(_179_v10), false, _180_v11)
	_181_globalState = _nw29
	var _182_v12 _dafny.Sequence
	_ = _182_v12
	_182_v12 = _dafny.UnicodeSeqOfUtf8Bytes("aiwqmuo")
	(_181_globalState).F2 = (Companion_Default___.SafeModuloInt(_169_v0, (_dafny.Zero).Minus(_169_v0))).Cmp(Companion_Default___.Fm0(_dafny.IntOfUint32((_182_v12).Cardinality()), _169_v0, _169_v0, _171_v2, _181_globalState)) < 0
	var _183_i0 _dafny.Int
	_ = _183_i0
	_183_i0 = _dafny.Zero
	{
		for _171_v2 {
			{
				if (_183_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_183_i0 = (_183_i0).Plus(_dafny.One)
				(_181_globalState).F0 = Companion_Default___.Fm1((_169_v0).Cmp(_dafny.IntOfInt64(311)) <= 0, _181_globalState)
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_172_v3), 0))
				_ = _index18
				(_172_v3).ArraySet1(_169_v0, (_index18).Int())
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_172_v3), 0))
				_ = _index19
				var _rhs16 _dafny.Int = _169_v0
				_ = _rhs16
				var _rhs17 _dafny.Int = _dafny.IntOfInt64(618)
				_ = _rhs17
				var _lhs16 _dafny.Array = _172_v3
				_ = _lhs16
				var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_172_v3), 0))
				_ = _lhs17
				(_lhs16).ArraySet1(_rhs16, (_lhs17).Int())
				_169_v0 = _rhs17
				var _184_v13 _dafny.CodePoint
				_ = _184_v13
				_184_v13 = _dafny.CodePoint('y')
				var _185_v14 _dafny.Map
				_ = _185_v14
				_185_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _dafny.IntOfUint32((_182_v12).Cardinality()))
				Companion_Default___.M0(!((_171_v2) == (_171_v2)), _184_v13, (_dafny.Zero).Minus((func() _dafny.Int {
					if (_185_v14).Contains(_171_v2) {
						return (_185_v14).Get(_171_v2).(_dafny.Int)
					}
					return (_172_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_172_v3), 0))).Int()).(_dafny.Int)
				})()), _185_v14, _181_globalState)
				var _186_v15 _dafny.Map
				_ = _186_v15
				_186_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_169_v0), _170_v1)
				_186_v15 = _186_v15
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _187_v17 _dafny.Map
	_ = _187_v17
	_187_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _dafny.IntOfUint32((_182_v12).Cardinality()))
	Companion_Default___.M0(true, (_182_v12).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0((func() _dafny.Set {
		var _coll18 = _dafny.NewBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((_182_v12).Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _188_v16 _dafny.CodePoint
			_188_v16 = interface{}(_compr_18).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_182_v12, _188_v16) {
				_coll18.Add(_188_v16)
			}
		}
		return _coll18.ToSet()
	}()).Cardinality(), (_dafny.Zero).Minus(_169_v0), _169_v0, _171_v2, _181_globalState), _dafny.IntOfUint32((_182_v12).Cardinality()))).Uint32()).(_dafny.CodePoint), (Companion_Default___.Fm0(_169_v0, _169_v0, (func() _dafny.Int {
		if (_187_v17).Contains(_171_v2) {
			return (_187_v17).Get(_171_v2).(_dafny.Int)
		}
		return _dafny.IntOfUint32((_182_v12).Cardinality())
	})(), _171_v2, _181_globalState)).Minus(_169_v0), Companion_Default___.Fm2(_171_v2, _171_v2, _171_v2, _187_v17, _181_globalState), _181_globalState)
	var _189_v18 _dafny.CodePoint
	_ = _189_v18
	_189_v18 = _dafny.CodePoint('n')
	Companion_Default___.M0(_171_v2, _189_v18, Companion_Default___.Fm0(_169_v0, _169_v0, _169_v0, _171_v2, _181_globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, Companion_Default___.Fm0(_169_v0, _169_v0, _dafny.IntOfUint32((_182_v12).Cardinality()), true, _181_globalState)), _181_globalState)
	var _hi0 _dafny.Int = _dafny.IntOfInt64(-752)
	_ = _hi0
	for _190_i1 := _169_v0; _190_i1.Cmp(_hi0) < 0; _190_i1 = _190_i1.Plus(_dafny.One) {
		_169_v0 = _190_i1
		(_181_globalState).F0 = _171_v2
		Companion_Default___.M0(_171_v2, _189_v18, _190_i1, _187_v17, _181_globalState)
		var _191_v19 _dafny.MultiSet
		_ = _191_v19
		_191_v19 = _dafny.MultiSetOf(_190_i1, _169_v0, _190_i1)
		_191_v19 = _191_v19
	}
	var _192_v20 D0
	_ = _192_v20
	_192_v20 = Companion_D0_.Create_DC1_(_171_v2)
	var _193_v21 _dafny.Map
	_ = _193_v21
	_193_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _192_v20)
	_193_v21 = (_193_v21).Update(_171_v2, Companion_D0_.Create_DC1_(_171_v2))
	var _source4 D0 = _192_v20
	_ = _source4
	if _source4.Is_DC1() {
		var _194___mcc_h0 bool = _source4.Get_().(D0_DC1).Cf1
		_ = _194___mcc_h0
		var _195_cf1 bool = _194___mcc_h0
		_ = _195_cf1
		var _196_v22 _dafny.Map
		_ = _196_v22
		_196_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("qrvcjj"), _169_v0)
		var _197_v23 _dafny.MultiSet
		_ = _197_v23
		_197_v23 = _dafny.MultiSetOf((_196_v22).Cardinality())
		var _198_v25 D0
		_ = _198_v25
		_198_v25 = Companion_D0_.Create_DC0_(_175_v6)
		var _199_v26 D0
		_ = _199_v26
		_199_v26 = Companion_D0_.Create_DC2_(_198_v25)
		var _200_v27 _dafny.Map
		_ = _200_v27
		_200_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_v26, _182_v12)
		var _rhs18 _dafny.Int = (func() _dafny.Int {
			if (_197_v23).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Map {
				var _coll19 = _dafny.NewMapBuilder()
				_ = _coll19
				for _iter19 := _dafny.Iterate((_200_v27).Keys().Elements()); ; {
					_compr_19, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _201_v24 D0
					_201_v24 = interface{}(_compr_19).(D0)
					if (_200_v27).Contains(_201_v24) {
						_coll19.Add(_201_v24, (_dafny.Zero).Minus(_169_v0))
					}
				}
				return _coll19.ToMap()
			}()).Cardinality()))) {
				return (_197_v23).Multiplicity((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter20 := _dafny.Iterate((_200_v27).Keys().Elements()); ; {
						_compr_20, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _202_v24 D0
						_202_v24 = interface{}(_compr_20).(D0)
						if (_200_v27).Contains(_202_v24) {
							_coll20.Add(_202_v24, (_dafny.Zero).Minus(_169_v0))
						}
					}
					return _coll20.ToMap()
				}()).Cardinality())))
			}
			return Companion_Default___.Fm0(_169_v0, _169_v0, _169_v0, _195_cf1, _181_globalState)
		})()
		_ = _rhs18
		var _rhs19 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(550), Companion_Default___.SafeDivisionInt(_169_v0, _169_v0))
		_ = _rhs19
		var _lhs18 *GlobalState = _181_globalState
		_ = _lhs18
		_lhs18.F5 = _rhs18
		_169_v0 = _rhs19
		var _203_v28 _dafny.Sequence
		_ = _203_v28
		_203_v28 = _dafny.SeqOf(_180_v11)
		(_181_globalState).F0 = !(((_180_v11).Intersection((_203_v28).Select((Companion_Default___.SafeIndex(_169_v0, _dafny.IntOfUint32((_203_v28).Cardinality()))).Uint32()).(_dafny.Set))).Contains((_195_cf1) || (Companion_Default___.Fm1(_195_cf1, _181_globalState))))
		_182_v12 = _182_v12
		(_181_globalState).F2 = _171_v2
	} else if _source4.Is_DC0() {
		var _204___mcc_h1 _dafny.Map = _source4.Get_().(D0_DC0).Cf0
		_ = _204___mcc_h1
		var _205_cf0 _dafny.Map = _204___mcc_h1
		_ = _205_cf0
		Companion_Default___.M0(!(_171_v2), _189_v18, _169_v0, _187_v17, _181_globalState)
		Companion_Default___.M0((!(_171_v2)) == (_171_v2), _189_v18, _169_v0, _187_v17, _181_globalState)
		_182_v12 = _dafny.Companion_Sequence_.Concatenate(_182_v12, _182_v12)
		(_181_globalState).F5 = _169_v0
	} else {
		var _206___mcc_h2 D0 = _source4.Get_().(D0_DC2).Cf2
		_ = _206___mcc_h2
		var _207_cf2 D0 = _206___mcc_h2
		_ = _207_cf2
		Companion_Default___.M0(_171_v2, _189_v18, (_dafny.Zero).Minus((_169_v0).Minus(_169_v0)), _187_v17, _181_globalState)
		var _pat_let_tv0 = _171_v2
		_ = _pat_let_tv0
		var _208_v29 _dafny.MultiSet
		_ = _208_v29
		_208_v29 = _dafny.MultiSetOf(_192_v20, _192_v20, func(_pat_let0_0 D0) D0 {
			return func(_209_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 bool) D0 {
					return func(_210_dt__update_hcf1_h0 bool) D0 {
						return Companion_D0_.Create_DC1_(_210_dt__update_hcf1_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_192_v20))
		var _211_v30 _dafny.Set
		_ = _211_v30
		_211_v30 = _dafny.SetOf(_172_v3)
		(_181_globalState).F2 = ((func() _dafny.MultiSet {
			if _171_v2 {
				return _dafny.MultiSetOf(Companion_D0_.Create_DC1_((_174_v5).Select((Companion_Default___.SafeIndex((_187_v17).Cardinality(), _dafny.IntOfUint32((_174_v5).Cardinality()))).Uint32()).(bool)))
			}
			return _208_v29
		})()).Equals((func() _dafny.MultiSet {
			if false {
				return ((_208_v29).Update(_192_v20, Companion_Default___.Abs(_169_v0))).Update(Companion_Default___.Fm3((_192_v20).Dtor_cf1(), _169_v0, (_211_v30).Cardinality(), _169_v0, _181_globalState), Companion_Default___.Abs(_169_v0))
			}
			return _208_v29
		})())
		_175_v6 = (_175_v6).Update(_169_v0, ((_dafny.Zero).Minus(_169_v0)).Cmp(_169_v0) < 0)
		var _212_v31 *C1
		_ = _212_v31
		var _nw30 *C1 = New_C1_()
		_ = _nw30
		_nw30.Ctor__()
		_212_v31 = _nw30
	}
	if false {
		(_181_globalState).F2 = _171_v2
		(_181_globalState).F0 = !(_171_v2)
		if _171_v2 {
			var _213_v32 _dafny.Map
			_ = _213_v32
			_213_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v0, _172_v3)
			var _214_v33 _dafny.MultiSet
			_ = _214_v33
			_214_v33 = _dafny.MultiSetOf(_171_v2, Companion_Default___.Fm1(_171_v2, _181_globalState), _171_v2)
			_213_v32 = (_213_v32).Update((_214_v33).Cardinality(), _172_v3)
			_170_v1 = _170_v1
			var _215_v34 D4
			_ = _215_v34
			_215_v34 = Companion_D4_.Create_DC12_(_dafny.IntOfUint32((_182_v12).Cardinality()), _169_v0, _182_v12, _dafny.IntOfInt64(350), _dafny.IntOfUint32((_174_v5).Cardinality()))
			_187_v17 = (_187_v17).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_182_v12, _dafny.UnicodeSeqOfUtf8Bytes("fvye")), (_215_v34).Dtor_cf17())
			var _216_v35 _dafny.Array
			_ = _216_v35
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw31
			_216_v35 = _nw31
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_216_v35), 0))
			_ = _index20
			(_216_v35).ArraySet1(_171_v2, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_216_v35), 0))
			_ = _index21
			(_216_v35).ArraySet1(_171_v2, (_index21).Int())
			(_181_globalState).F2 = _171_v2
		} else {
			(_181_globalState).F2 = _171_v2
			_169_v0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_182_v12, _dafny.UnicodeSeqOfUtf8Bytes("poavktqf"))).Cardinality())
			var _217_v36 _dafny.Array
			_ = _217_v36
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_6
			var _nw32 _dafny.Array
			_ = _nw32
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw32 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Map = (func(_218_v6 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_219_i2 _dafny.Int) _dafny.Map {
						return _218_v6
					}
				})(_175_v6)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw32 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw32).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw32).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_217_v36 = _nw32
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_217_v36), 0))
			_ = _index22
			(_217_v36).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v0, _171_v2), (_index22).Int())
			var _220_v37 D8
			_ = _220_v37
			_220_v37 = Companion_D8_.Create_DC21_(_171_v2, _171_v2, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_174_v5, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wcd")).Cardinality()), _dafny.IntOfUint32((_174_v5).Cardinality()))).Uint32(), _171_v2), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-577), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_174_v5, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wcd")).Cardinality()), _dafny.IntOfUint32((_174_v5).Cardinality()))).Uint32(), _171_v2)).Cardinality()))).Uint32(), _171_v2), _171_v2, _169_v0)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_217_v36), 0))
			_ = _index23
			(_217_v36).ArraySet1((_175_v6).Update(_169_v0, (_220_v37).Dtor_cf27()), (_index23).Int())
			_174_v5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_174_v5, (Companion_Default___.SafeIndex(_169_v0, _dafny.IntOfUint32((_174_v5).Cardinality()))).Uint32(), _171_v2), _174_v5), _174_v5)
			var _221_v38 T1
			_ = _221_v38
			var _nw33 *C4 = New_C4_()
			_ = _nw33
			_nw33.Ctor__(_169_v0)
			_221_v38 = _nw33
			var _222_v39 _dafny.Sequence
			_ = _222_v39
			_222_v39 = _dafny.SeqOf(_221_v38, _221_v38)
			Companion_Default___.M0(_171_v2, _189_v18, Companion_Default___.SafeDivisionInt(_169_v0, Companion_Default___.Fm0(_169_v0, _169_v0, _169_v0, _171_v2, _181_globalState)), (_187_v17).Update(_171_v2, _dafny.IntOfUint32((_222_v39).Cardinality())), _181_globalState)
		}
		var _223_v40 _dafny.Map
		_ = _223_v40
		_223_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v0, Companion_D8_.Create_DC21_(_171_v2, _171_v2, _174_v5, _171_v2, _169_v0))
		_223_v40 = Companion_Default___.Fm33(_171_v2, _189_v18, _181_globalState)
		var _224_v41 _dafny.Array
		_ = _224_v41
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_7
		var _nw34 _dafny.Array
		_ = _nw34
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw34 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.CodePoint = (func(_225_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_226_i3 _dafny.Int) _dafny.CodePoint {
					return _225_v18
				}
			})(_189_v18)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw34 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw34).ArraySet1CodePoint(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw34).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_224_v41 = _nw34
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_224_v41), 0))
		_ = _index24
		(_224_v41).ArraySet1CodePoint(_189_v18, (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_224_v41), 0))
		_ = _index25
		(_224_v41).ArraySet1CodePoint(_189_v18, (_index25).Int())
	} else {
		var _227_v42 _dafny.Map
		_ = _227_v42
		_227_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v0, _169_v0)
		_227_v42 = (_227_v42).Update(_169_v0, _169_v0)
		_169_v0 = (_dafny.Zero).Minus((_187_v17).Cardinality())
		var _228_v43 *C1
		_ = _228_v43
		var _nw35 *C1 = New_C1_()
		_ = _nw35
		_nw35.Ctor__()
		_228_v43 = _nw35
		_228_v43 = _228_v43
		(_181_globalState).F2 = _171_v2
		(_181_globalState).F2 = _171_v2
	}
	(_181_globalState).F2 = (_169_v0).Cmp(_169_v0) < 0
	var _229_v44 _dafny.Map
	_ = _229_v44
	_229_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v12, _169_v0)
	_229_v44 = (_229_v44).Update(_182_v12, Companion_Default___.SafeDivisionInt(_169_v0, _169_v0))
	Companion_Default___.M0(_171_v2, _189_v18, _169_v0, (_187_v17).Merge(_187_v17), _181_globalState)
	var _230_v45 _dafny.MultiSet
	_ = _230_v45
	_230_v45 = _dafny.MultiSetOf(_171_v2)
	(_181_globalState).F2 = (_230_v45).IsDisjointFrom(_230_v45)
	if _171_v2 {
		var _231_v46 _dafny.MultiSet
		_ = _231_v46
		_231_v46 = _dafny.MultiSetOf(_189_v18, _189_v18)
		if (Companion_Default___.Fm0((func() _dafny.Int {
			if (_231_v46).Contains(_189_v18) {
				return (_231_v46).Multiplicity(_189_v18)
			}
			return _169_v0
		})(), _dafny.IntOfUint32((_170_v1).Cardinality()), _169_v0, _171_v2, _181_globalState)).Cmp(_169_v0) <= 0 {
			Companion_Default___.M0(_171_v2, _189_v18, _169_v0, _187_v17, _181_globalState)
			var _232_v47 T1
			_ = _232_v47
			var _nw36 *C3 = New_C3_()
			_ = _nw36
			_nw36.Ctor__(_dafny.IntOfInt64(997))
			_232_v47 = _nw36
			var _233_v48 _dafny.Set
			_ = _233_v48
			_233_v48 = _dafny.SetOf(_232_v47)
			var _234_v49 D8
			_ = _234_v49
			_234_v49 = Companion_D8_.Create_DC20_((_233_v48).Intersection(_233_v48))
			_234_v49 = (func() D8 {
				if (_169_v0).Cmp(_169_v0) < 0 {
					return _234_v49
				}
				return _234_v49
			})()
			var _235_v50 _dafny.Array
			_ = _235_v50
			var _nwElement0_3 bool = _171_v2
			_ = _nwElement0_3
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(25))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_3, 0)
			(_nw37).ArraySet1(_171_v2, 1)
			(_nw37).ArraySet1(_171_v2, 2)
			(_nw37).ArraySet1(_171_v2, 3)
			(_nw37).ArraySet1(_171_v2, 4)
			(_nw37).ArraySet1(_171_v2, 5)
			(_nw37).ArraySet1(false, 6)
			(_nw37).ArraySet1(_171_v2, 7)
			(_nw37).ArraySet1(true, 8)
			(_nw37).ArraySet1(true, 9)
			(_nw37).ArraySet1(_171_v2, 10)
			(_nw37).ArraySet1(_171_v2, 11)
			(_nw37).ArraySet1(true, 12)
			(_nw37).ArraySet1(_171_v2, 13)
			(_nw37).ArraySet1(_171_v2, 14)
			(_nw37).ArraySet1(_171_v2, 15)
			(_nw37).ArraySet1((_232_v47).Fm4(_177_v9, Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(793), _171_v2)), _dafny.MultiSetFromSeq(_dafny.SeqOf(_171_v2)), _181_globalState), 16)
			(_nw37).ArraySet1(_171_v2, 17)
			(_nw37).ArraySet1(_171_v2, 18)
			(_nw37).ArraySet1(_171_v2, 19)
			(_nw37).ArraySet1(_171_v2, 20)
			(_nw37).ArraySet1(_171_v2, 21)
			(_nw37).ArraySet1(_171_v2, 22)
			(_nw37).ArraySet1(_171_v2, 23)
			(_nw37).ArraySet1(_171_v2, 24)
			_235_v50 = _nw37
			var _236_v51 _dafny.Map
			_ = _236_v51
			_236_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v1, _235_v50)
			var _237_v52 _dafny.Map
			_ = _237_v52
			_237_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v51, (Companion_D9_.Create_DC24_(_171_v2, false, _171_v2)).Dtor_cf34())
			_237_v52 = (_237_v52).Update(_236_v51, (_180_v11).Contains(_171_v2))
			(_181_globalState).F0 = true
			(_181_globalState).F2 = Companion_Default___.Fm1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_238_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_239_i4 _dafny.Int) _dafny.CodePoint {
					return _238_v18
				}
			})(_189_v18))), (Companion_Default___.SafeIndex(_169_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_240_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_241_i4 _dafny.Int) _dafny.CodePoint {
					return _240_v18
				}
			})(_189_v18)))).Cardinality()))).Uint32(), _189_v18), _182_v12), _181_globalState)
		} else {
			var _rhs20 _dafny.Int = Companion_Default___.Fm0(_169_v0, _169_v0, (_dafny.Zero).Minus(_169_v0), !(_171_v2), _181_globalState)
			_ = _rhs20
			var _rhs21 bool = ((func() _dafny.MultiSet {
				if !(_171_v2) {
					return _231_v46
				}
				return _231_v46
			})()).IsSubsetOf(_dafny.MultiSetOf(_189_v18))
			_ = _rhs21
			var _lhs19 *GlobalState = _181_globalState
			_ = _lhs19
			_169_v0 = _rhs20
			_lhs19.F2 = _rhs21
			var _242_v53 _dafny.Map
			_ = _242_v53
			_242_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _171_v2)
			var _243_v54 D2
			_ = _243_v54
			_243_v54 = Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_182_v12).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v53, _169_v0), _169_v0)
			var _244_v55 _dafny.MultiSet
			_ = _244_v55
			_244_v55 = _dafny.MultiSetOf(_243_v54)
			var _245_v56 _dafny.Map
			_ = _245_v56
			_245_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_244_v55).Contains(_243_v54) {
					return (_244_v55).Multiplicity(_243_v54)
				}
				return _169_v0
			})(), _169_v0)
			var _246_v57 T0
			_ = _246_v57
			var _nw38 *C1 = New_C1_()
			_ = _nw38
			_nw38.Ctor__()
			_246_v57 = _nw38
			var _247_v58 _dafny.Set
			_ = _247_v58
			_247_v58 = _dafny.SetOf(_246_v57)
			_245_v56 = (_245_v56).Update(Companion_Default___.SafeModuloInt(_169_v0, _169_v0), ((_247_v58).Intersection(_247_v58)).Cardinality())
			_172_v3 = _172_v3
			(_181_globalState).F5 = _169_v0
			var _248_v59 T1
			_ = _248_v59
			var _nw39 *C3 = New_C3_()
			_ = _nw39
			_nw39.Ctor__(_169_v0)
			_248_v59 = _nw39
			var _249_v60 D2
			_ = _249_v60
			_249_v60 = Companion_D2_.Create_DC5_(_248_v59)
			var _250_v61 _dafny.Map
			_ = _250_v61
			_250_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_174_v5), (_249_v60).Dtor_cf5())
			var _251_v62 D9
			_ = _251_v62
			_251_v62 = Companion_D9_.Create_DC24_(_171_v2, _171_v2, _171_v2)
			var _252_v63 _dafny.Array
			_ = _252_v63
			var _nwElement0_4 T1 = _248_v59
			_ = _nwElement0_4
			var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(15))
			_ = _nw40
			(_nw40).ArraySet1(_nwElement0_4, 0)
			(_nw40).ArraySet1((func() T1 {
				if _171_v2 {
					return _248_v59
				}
				return _248_v59
			})(), 1)
			(_nw40).ArraySet1(_248_v59, 2)
			(_nw40).ArraySet1((func() T1 {
				if (_250_v61).Contains(_230_v45) {
					return (_250_v61).Get(_230_v45).(T1)
				}
				return _248_v59
			})(), 3)
			(_nw40).ArraySet1(_248_v59, 4)
			(_nw40).ArraySet1(_248_v59, 5)
			(_nw40).ArraySet1(_248_v59, 6)
			(_nw40).ArraySet1((func() T1 {
				if (_251_v62).Dtor_cf35() {
					return _248_v59
				}
				return _248_v59
			})(), 7)
			(_nw40).ArraySet1(_248_v59, 8)
			(_nw40).ArraySet1(_248_v59, 9)
			(_nw40).ArraySet1(_248_v59, 10)
			(_nw40).ArraySet1(_248_v59, 11)
			(_nw40).ArraySet1(_248_v59, 12)
			(_nw40).ArraySet1(_248_v59, 13)
			(_nw40).ArraySet1(_248_v59, 14)
			_252_v63 = _nw40
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_252_v63), 0))
			_ = _index26
			(_252_v63).ArraySet1(_248_v59, (_index26).Int())
			var _253_v64 D13
			_ = _253_v64
			_253_v64 = Companion_D13_.Create_DC38_(_187_v17)
			var _254_v65 _dafny.MultiSet
			_ = _254_v65
			_254_v65 = _dafny.MultiSetOf(_253_v64, _253_v64)
			var _255_v66 _dafny.Array
			_ = _255_v66
			var _nwElement0_5 bool = !(_171_v2)
			_ = _nwElement0_5
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(3))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_5, 0)
			(_nw41).ArraySet1((_254_v65).IsProperSubsetOf(_254_v65), 1)
			(_nw41).ArraySet1(_171_v2, 2)
			_255_v66 = _nw41
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_255_v66), 0))
			_ = _index27
			(_255_v66).ArraySet1(_171_v2, (_index27).Int())
			var _256_v67 *C3
			_ = _256_v67
			var _nw42 *C3 = New_C3_()
			_ = _nw42
			_nw42.Ctor__(_169_v0)
			_256_v67 = _nw42
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_252_v63), 0))
			_ = _index28
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_255_v66), 0))
			_ = _index29
			var _rhs22 T1 = _248_v59
			_ = _rhs22
			var _rhs23 bool = _171_v2
			_ = _rhs23
			var _rhs24 *C3 = _256_v67
			_ = _rhs24
			var _lhs20 _dafny.Array = _252_v63
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_252_v63), 0))
			_ = _lhs21
			var _lhs22 _dafny.Array = _255_v66
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_255_v66), 0))
			_ = _lhs23
			(_lhs20).ArraySet1(_rhs22, (_lhs21).Int())
			(_lhs22).ArraySet1(_rhs23, (_lhs23).Int())
			_256_v67 = _rhs24
		}
		var _257_v68 _dafny.Array
		_ = _257_v68
		var _nw43 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
		_ = _nw43
		_257_v68 = _nw43
		var _258_v69 T0
		_ = _258_v69
		var _nw44 *C2 = New_C2_()
		_ = _nw44
		_nw44.Ctor__(_171_v2)
		_258_v69 = _nw44
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_257_v68), 0))
		_ = _index30
		(_257_v68).ArraySet1(_258_v69, (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_257_v68), 0))
		_ = _index31
		var _rhs25 _dafny.Sequence = (func() _dafny.Sequence {
			if _171_v2 {
				return _182_v12
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(938))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_259_v2 bool, _260_v0 _dafny.Int, _261_v18 _dafny.CodePoint, _262_v45 _dafny.MultiSet) func(_dafny.Int) _dafny.CodePoint {
				return func(_263_i5 _dafny.Int) _dafny.CodePoint {
					return (Companion_D10_.Create_DC27_(_259_v2, _260_v0, _261_v18, _262_v45)).Dtor_cf40()
				}
			})(_171_v2, _169_v0, _189_v18, _230_v45)))
		})()
		_ = _rhs25
		var _rhs26 T0 = _258_v69
		_ = _rhs26
		var _lhs24 _dafny.Array = _257_v68
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_257_v68), 0))
		_ = _lhs25
		_182_v12 = _rhs25
		(_lhs24).ArraySet1(_rhs26, (_lhs25).Int())
		(_181_globalState).F0 = (func() bool {
			if (_169_v0).Cmp(_169_v0) <= 0 {
				return _171_v2
			}
			return _171_v2
		})()
		(_181_globalState).F5 = (_dafny.Zero).Minus(_169_v0)
		(_181_globalState).F5 = ((Companion_Default___.Fm34(_169_v0, _169_v0, !(_171_v2), _181_globalState)).Update(_180_v11, _dafny.IntOfInt64(-582))).Cardinality()
	} else {
		Companion_Default___.M0(!(_171_v2) || (_171_v2), _189_v18, _169_v0, _187_v17, _181_globalState)
		_171_v2 = !_dafny.Companion_Sequence_.Contains(_170_v1, _169_v0)
		var _264_v70 _dafny.Array
		_ = _264_v70
		var _nwElement0_6 bool = false
		_ = _nwElement0_6
		var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(28))
		_ = _nw45
		(_nw45).ArraySet1(_nwElement0_6, 0)
		(_nw45).ArraySet1(_171_v2, 1)
		(_nw45).ArraySet1(_171_v2, 2)
		(_nw45).ArraySet1(_171_v2, 3)
		(_nw45).ArraySet1(_171_v2, 4)
		(_nw45).ArraySet1(_171_v2, 5)
		(_nw45).ArraySet1(_171_v2, 6)
		(_nw45).ArraySet1(_171_v2, 7)
		(_nw45).ArraySet1(_171_v2, 8)
		(_nw45).ArraySet1(_171_v2, 9)
		(_nw45).ArraySet1(_171_v2, 10)
		(_nw45).ArraySet1(_171_v2, 11)
		(_nw45).ArraySet1(_171_v2, 12)
		(_nw45).ArraySet1(!(_171_v2), 13)
		(_nw45).ArraySet1(Companion_Default___.Fm1(!(_171_v2), _181_globalState), 14)
		(_nw45).ArraySet1(_171_v2, 15)
		(_nw45).ArraySet1(_171_v2, 16)
		(_nw45).ArraySet1(_171_v2, 17)
		(_nw45).ArraySet1(_171_v2, 18)
		(_nw45).ArraySet1(_171_v2, 19)
		(_nw45).ArraySet1(_171_v2, 20)
		(_nw45).ArraySet1(_171_v2, 21)
		(_nw45).ArraySet1(_171_v2, 22)
		(_nw45).ArraySet1(!(_171_v2), 23)
		(_nw45).ArraySet1(_171_v2, 24)
		(_nw45).ArraySet1(false, 25)
		(_nw45).ArraySet1(true, 26)
		(_nw45).ArraySet1(_171_v2, 27)
		_264_v70 = _nw45
		var _265_v71 *C0
		_ = _265_v71
		var _nw46 *C0 = New_C0_()
		_ = _nw46
		_nw46.Ctor__(true, _264_v70)
		_265_v71 = _nw46
		var _266_v72 *C2
		_ = _266_v72
		var _nw47 *C2 = New_C2_()
		_ = _nw47
		_nw47.Ctor__((_265_v71).F10())
		_266_v72 = _nw47
		var _267_v73 _dafny.Array
		_ = _267_v73
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw48
		_267_v73 = _nw48
		var _268_v74 bool
		_ = _268_v74
		var _269_v75 _dafny.Int
		_ = _269_v75
		var _270_v76 bool
		_ = _270_v76
		var _271_v77 bool
		_ = _271_v77
		var _out0 bool
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		var _out2 bool
		_ = _out2
		var _out3 bool
		_ = _out3
		_out0, _out1, _out2, _out3 = (_266_v72).M1(_169_v0, _267_v73, ((_dafny.MultiSetFromSeq(_174_v5)).Union(_dafny.MultiSetFromSeq(_174_v5))).Cardinality(), _169_v0, _181_globalState)
		_268_v74 = _out0
		_269_v75 = _out1
		_270_v76 = _out2
		_271_v77 = _out3
	}
	var _272_v78 _dafny.Array
	_ = _272_v78
	var _nw49 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
	_ = _nw49
	_272_v78 = _nw49
	var _273_v79 *C0
	_ = _273_v79
	var _nw50 *C0 = New_C0_()
	_ = _nw50
	_nw50.Ctor__(_171_v2, _272_v78)
	_273_v79 = _nw50
	var _274_v80 _dafny.Set
	_ = _274_v80
	_274_v80 = _dafny.SetOf(_273_v79)
	var _275_v81 _dafny.Map
	_ = _275_v81
	_275_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _274_v80)
	var _276_v82 _dafny.Map
	_ = _276_v82
	_276_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Set {
		if (_275_v81).Contains((_273_v79).Fm4(_177_v9, Companion_D0_.Create_DC0_(_175_v6), _230_v45, _181_globalState)) {
			return (_275_v81).Get((_273_v79).Fm4(_177_v9, Companion_D0_.Create_DC0_(_175_v6), _230_v45, _181_globalState)).(_dafny.Set)
		}
		return _274_v80
	})()).Cardinality(), _169_v0)
	var _277_v83 T1
	_ = _277_v83
	var _nw51 *C0 = New_C0_()
	_ = _nw51
	_nw51.Ctor__((_273_v79).F10(), (_273_v79).F11())
	_277_v83 = _nw51
	var _278_v84 _dafny.Map
	_ = _278_v84
	_278_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v83, _169_v0)
	var _279_v85 _dafny.Map
	_ = _279_v85
	_279_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v2, _278_v84)
	_276_v82 = (_276_v82).Update((func() _dafny.Int {
		if (_273_v79).F10() {
			return _169_v0
		}
		return Companion_Default___.Fm0(_169_v0, _169_v0, (_279_v85).Cardinality(), (_273_v79).F10(), _181_globalState)
	})(), _169_v0)
	var _280_v86 _dafny.Array
	_ = _280_v86
	var _len0_8 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_8
	var _nw52 _dafny.Array
	_ = _nw52
	if _len0_8.Cmp(_dafny.Zero) == 0 {
		_nw52 = _dafny.NewArray(_len0_8)
	} else {
		var _init8 func(_dafny.Int) D13 = (func(_281_v2 bool, _282_v0 _dafny.Int) func(_dafny.Int) D13 {
			return func(_283_i6 _dafny.Int) D13 {
				return Companion_D13_.Create_DC38_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v2, _282_v0))
			}
		})(_171_v2, _169_v0)
		_ = _init8
		var _element0_8 = _init8(_dafny.Zero)
		_ = _element0_8
		_nw52 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
		(_nw52).ArraySet1(_element0_8, 0)
		var _nativeLen0_8 = (_len0_8).Int()
		_ = _nativeLen0_8
		for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
			(_nw52).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
		}
	}
	_280_v86 = _nw52
	var _284_v87 _dafny.Set
	_ = _284_v87
	_284_v87 = _dafny.SetOf(_280_v86)
	if (_284_v87).IsDisjointFrom(_284_v87) {
		_180_v11 = Companion_Default___.Fm17(!((_273_v79).F10()), (_dafny.Zero).Minus(_169_v0), (_273_v79).F10(), _169_v0, _181_globalState)
		var _285_v88 *C4
		_ = _285_v88
		var _nw53 *C4 = New_C4_()
		_ = _nw53
		_nw53.Ctor__(_169_v0)
		_285_v88 = _nw53
		var _286_v89 _dafny.Sequence
		_ = _286_v89
		_286_v89 = _dafny.SeqOf(_285_v88)
		var _287_v90 _dafny.Set
		_ = _287_v90
		_287_v90 = _dafny.SetOf(_dafny.IntOfUint32((_286_v89).Cardinality()), (_285_v88).F9(), (_285_v88).F9())
		var _288_v91 _dafny.Map
		_ = _288_v91
		_288_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_273_v79, (_287_v90).Cardinality())
		_288_v91 = (_288_v91).Update(_273_v79, (_dafny.Zero).Minus(_169_v0))
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw54
		_272_v78 = _nw54
		var _289_v92 D3
		_ = _289_v92
		_289_v92 = Companion_D3_.Create_DC8_(_182_v12)
		(_181_globalState).F2 = _dafny.Companion_Sequence_.IsPrefixOf((_289_v92).Dtor_cf10(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(956))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_290_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _182_v12))
		var _291_v93 *C1
		_ = _291_v93
		var _nw55 *C1 = New_C1_()
		_ = _nw55
		_nw55.Ctor__()
		_291_v93 = _nw55
		var _292_v94 _dafny.Map
		_ = _292_v94
		_292_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v93, _171_v2)
		_169_v0 = ((_292_v94).Cardinality()).Times((_285_v88).F9())
	} else {
		(_181_globalState).F2 = (_169_v0).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_182_v12, (Companion_Default___.SafeIndex(_169_v0, _dafny.IntOfUint32((_182_v12).Cardinality()))).Uint32(), _189_v18)).Cardinality())) != 0
		var _293_v95 *C4
		_ = _293_v95
		var _nw56 *C4 = New_C4_()
		_ = _nw56
		_nw56.Ctor__(_169_v0)
		_293_v95 = _nw56
		_171_v2 = _171_v2
		var _294_v96 *C1
		_ = _294_v96
		var _nw57 *C1 = New_C1_()
		_ = _nw57
		_nw57.Ctor__()
		_294_v96 = _nw57
		(_181_globalState).F5 = (_169_v0).Plus(_169_v0)
	}
	var _295_v97 *C1
	_ = _295_v97
	var _nw58 *C1 = New_C1_()
	_ = _nw58
	_nw58.Ctor__()
	_295_v97 = _nw58
	var _296_v98 _dafny.Sequence
	_ = _296_v98
	_296_v98 = _dafny.SeqOf(_295_v97, _295_v97)
	_295_v97 = (_296_v98).Select((Companion_Default___.SafeIndex(_169_v0, _dafny.IntOfUint32((_296_v98).Cardinality()))).Uint32()).(*C1)
	_dafny.Print(_169_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_170_v1, _dafny.SeqOf(_dafny.IntOfInt64(292), _dafny.IntOfInt64(292))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_171_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_v4).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_174_v5, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(292), _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_177_v9).Dtor_cf0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false).UpdateUnsafe(_dafny.IntOfInt64(2), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v10).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false).UpdateUnsafe(_dafny.IntOfInt64(2), false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v11).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_181_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_181_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_globalState.F3).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_181_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_181_globalState).F6()).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false).UpdateUnsafe(_dafny.IntOfInt64(2), false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_globalState.F8).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_v12.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_183_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(7))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_189_v18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v20).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_193_v21).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D0_.Create_DC1_(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v44).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("aiwqmuo"), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v45).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v79).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v80).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v81).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v82).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(-1)).UpdateUnsafe(_dafny.IntOfInt64(-1), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_278_v84).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v85).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.Zero).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.One).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_280_v86).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D13)).Dtor_cf55()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v87).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_296_v98).Cardinality()))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Map
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf2 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 D0) D0 {
	return D0{D0_DC2{Cf2}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf0() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf2() D0 {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
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
	Cf4 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.Int) D1 {
	return D1{D1_DC4{Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf3 _dafny.Array
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Array) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero)
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3 == data2.Cf3
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
	Cf6 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 bool) D2 {
	return D2{D2_DC6{Cf6}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf7 _dafny.Int
	Cf8 _dafny.Map
	Cf9 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf7 _dafny.Int, Cf8 _dafny.Map, Cf9 _dafny.Int) D2 {
	return D2{D2_DC7{Cf7, Cf8, Cf9}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf5 T1
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 T1) D2 {
	return D2{D2_DC5{Cf5}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(false)
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf5() T1 {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6 == data2.Cf6
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Equals(data2.Cf8) && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && _dafny.AreEqual(data1.Cf5, data2.Cf5)
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
	Cf11 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 _dafny.Int) D3 {
	return D3{D3_DC9{Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf10 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf10 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf10}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC10 struct {
	Cf12 D3
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf12 D3) D3 {
	return D3{D3_DC10{Cf12}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.Zero)
}

func (_this D3) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf10
}

func (_this D3) Dtor_cf12() D3 {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + data.Cf10.VerbatimString(true) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf14 _dafny.Int
	Cf15 _dafny.Int
	Cf16 _dafny.Sequence
	Cf17 _dafny.Int
	Cf18 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf14 _dafny.Int, Cf15 _dafny.Int, Cf16 _dafny.Sequence, Cf17 _dafny.Int, Cf18 _dafny.Int) D4 {
	return D4{D4_DC12{Cf14, Cf15, Cf16, Cf17, Cf18}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf19 _dafny.Int
	Cf20 _dafny.Array
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf19 _dafny.Int, Cf20 _dafny.Array) D4 {
	return D4{D4_DC13{Cf19, Cf20}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_() D4 {
	return D4{D4_DC14{}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC11 struct {
	Cf13 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf13 _dafny.Array) D4 {
	return D4{D4_DC11{Cf13}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(_dafny.Zero, _dafny.Zero, _dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf13
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + data.Cf16.VerbatimString(true) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Equals(data2.Cf16) && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D4_DC14:
		{
			_, ok := other.Get_().(D4_DC14)
			return ok
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf13 == data2.Cf13
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

type D5_DC15 struct {
	Cf21 _dafny.Array
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf21 _dafny.Array) D5 {
	return D5{D5_DC15{Cf21}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf21 == data2.Cf21
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
	Cf22 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D6) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf22.Equals(data2.Cf22)
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

type D7_DC18 struct {
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_() D7 {
	return D7{D7_DC18{}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC17 struct {
	Cf23 _dafny.Sequence
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf23 _dafny.Sequence) D7 {
	return D7{D7_DC17{Cf23}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC19 struct {
	Cf24 D7
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf24 D7) D7 {
	return D7{D7_DC19{Cf24}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_()
}

func (_this D7) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D7_DC17).Cf23
}

func (_this D7) Dtor_cf24() D7 {
	return _this.Get_().(D7_DC19).Cf24
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC18:
		{
			_, ok := other.Get_().(D7_DC18)
			return ok
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf26 bool
	Cf27 bool
	Cf28 _dafny.Sequence
	Cf29 bool
	Cf30 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf26 bool, Cf27 bool, Cf28 _dafny.Sequence, Cf29 bool, Cf30 _dafny.Int) D8 {
	return D8{D8_DC21{Cf26, Cf27, Cf28, Cf29, Cf30}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC20 struct {
	Cf25 _dafny.Set
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf25 _dafny.Set) D8 {
	return D8{D8_DC20{Cf25}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC22 struct {
	Cf31 D8
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf31 D8) D8 {
	return D8{D8_DC22{Cf31}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(false, false, _dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D8) Dtor_cf26() bool {
	return _this.Get_().(D8_DC21).Cf26
}

func (_this D8) Dtor_cf27() bool {
	return _this.Get_().(D8_DC21).Cf27
}

func (_this D8) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D8_DC21).Cf28
}

func (_this D8) Dtor_cf29() bool {
	return _this.Get_().(D8_DC21).Cf29
}

func (_this D8) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf30
}

func (_this D8) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D8_DC20).Cf25
}

func (_this D8) Dtor_cf31() D8 {
	return _this.Get_().(D8_DC22).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28.Equals(data2.Cf28) && data1.Cf29 == data2.Cf29 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf33 bool
	Cf34 bool
	Cf35 bool
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf33 bool, Cf34 bool, Cf35 bool) D9 {
	return D9{D9_DC24{Cf33, Cf34, Cf35}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC23 struct {
	Cf32 _dafny.Map
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf32 _dafny.Map) D9 {
	return D9{D9_DC23{Cf32}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC25 struct {
	Cf36 D9
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf36 D9) D9 {
	return D9{D9_DC25{Cf36}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(false, false, false)
}

func (_this D9) Dtor_cf33() bool {
	return _this.Get_().(D9_DC24).Cf33
}

func (_this D9) Dtor_cf34() bool {
	return _this.Get_().(D9_DC24).Cf34
}

func (_this D9) Dtor_cf35() bool {
	return _this.Get_().(D9_DC24).Cf35
}

func (_this D9) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D9_DC23).Cf32
}

func (_this D9) Dtor_cf36() D9 {
	return _this.Get_().(D9_DC25).Cf36
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D10_DC27 struct {
	Cf38 bool
	Cf39 _dafny.Int
	Cf40 _dafny.CodePoint
	Cf41 _dafny.MultiSet
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf38 bool, Cf39 _dafny.Int, Cf40 _dafny.CodePoint, Cf41 _dafny.MultiSet) D10 {
	return D10{D10_DC27{Cf38, Cf39, Cf40, Cf41}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

type D10_DC26 struct {
	Cf37 _dafny.Sequence
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf37 _dafny.Sequence) D10 {
	return D10{D10_DC26{Cf37}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC28 struct {
	Cf42 D10
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf42 D10) D10 {
	return D10{D10_DC28{Cf42}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC27_(false, _dafny.Zero, _dafny.CodePoint('D'), _dafny.EmptyMultiSet)
}

func (_this D10) Dtor_cf38() bool {
	return _this.Get_().(D10_DC27).Cf38
}

func (_this D10) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D10_DC27).Cf39
}

func (_this D10) Dtor_cf40() _dafny.CodePoint {
	return _this.Get_().(D10_DC27).Cf40
}

func (_this D10) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D10_DC27).Cf41
}

func (_this D10) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D10_DC26).Cf37
}

func (_this D10) Dtor_cf42() D10 {
	return _this.Get_().(D10_DC28).Cf42
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39.Cmp(data2.Cf39) == 0 && data1.Cf40 == data2.Cf40 && data1.Cf41.Equals(data2.Cf41)
		}
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
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

type D11_DC30 struct {
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_() D11 {
	return D11{D11_DC30{}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC31 struct {
	Cf44 _dafny.Int
	Cf45 _dafny.Int
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf44 _dafny.Int, Cf45 _dafny.Int) D11 {
	return D11{D11_DC31{Cf44, Cf45}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC29 struct {
	Cf43 T0
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf43 T0) D11 {
	return D11{D11_DC29{Cf43}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC32 struct {
	Cf46 D11
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf46 D11) D11 {
	return D11{D11_DC32{Cf46}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_()
}

func (_this D11) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf44
}

func (_this D11) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf45
}

func (_this D11) Dtor_cf43() T0 {
	return _this.Get_().(D11_DC29).Cf43
}

func (_this D11) Dtor_cf46() D11 {
	return _this.Get_().(D11_DC32).Cf46
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			_, ok := other.Get_().(D11_DC30)
			return ok
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45.Cmp(data2.Cf45) == 0
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && _dafny.AreEqual(data1.Cf43, data2.Cf43)
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D12_DC34 struct {
	Cf48 D0
	Cf49 bool
	Cf50 bool
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf48 D0, Cf49 bool, Cf50 bool) D12 {
	return D12{D12_DC34{Cf48, Cf49, Cf50}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

type D12_DC35 struct {
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_() D12 {
	return D12{D12_DC35{}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC36 struct {
	Cf51 bool
	Cf52 _dafny.Int
	Cf53 bool
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf51 bool, Cf52 _dafny.Int, Cf53 bool) D12 {
	return D12{D12_DC36{Cf51, Cf52, Cf53}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC33 struct {
	Cf47 _dafny.Set
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf47 _dafny.Set) D12 {
	return D12{D12_DC33{Cf47}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC37 struct {
	Cf54 D12
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf54 D12) D12 {
	return D12{D12_DC37{Cf54}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC34_(Companion_D0_.Default(), false, false)
}

func (_this D12) Dtor_cf48() D0 {
	return _this.Get_().(D12_DC34).Cf48
}

func (_this D12) Dtor_cf49() bool {
	return _this.Get_().(D12_DC34).Cf49
}

func (_this D12) Dtor_cf50() bool {
	return _this.Get_().(D12_DC34).Cf50
}

func (_this D12) Dtor_cf51() bool {
	return _this.Get_().(D12_DC36).Cf51
}

func (_this D12) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf52
}

func (_this D12) Dtor_cf53() bool {
	return _this.Get_().(D12_DC36).Cf53
}

func (_this D12) Dtor_cf47() _dafny.Set {
	return _this.Get_().(D12_DC33).Cf47
}

func (_this D12) Dtor_cf54() D12 {
	return _this.Get_().(D12_DC37).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D12_DC35:
		{
			return "D12.DC35"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf48.Equals(data2.Cf48) && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50
		}
	case D12_DC35:
		{
			_, ok := other.Get_().(D12_DC35)
			return ok
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53 == data2.Cf53
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf47.Equals(data2.Cf47)
		}
	case D12_DC37:
		{
			data2, ok := other.Get_().(D12_DC37)
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

type D13_DC39 struct {
}

func (D13_DC39) isD13() {}

func (CompanionStruct_D13_) Create_DC39_() D13 {
	return D13{D13_DC39{}}
}

func (_this D13) Is_DC39() bool {
	_, ok := _this.Get_().(D13_DC39)
	return ok
}

type D13_DC40 struct {
	Cf56 bool
	Cf57 *C0
	Cf58 _dafny.Map
}

func (D13_DC40) isD13() {}

func (CompanionStruct_D13_) Create_DC40_(Cf56 bool, Cf57 *C0, Cf58 _dafny.Map) D13 {
	return D13{D13_DC40{Cf56, Cf57, Cf58}}
}

func (_this D13) Is_DC40() bool {
	_, ok := _this.Get_().(D13_DC40)
	return ok
}

type D13_DC38 struct {
	Cf55 _dafny.Map
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf55 _dafny.Map) D13 {
	return D13{D13_DC38{Cf55}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC39_()
}

func (_this D13) Dtor_cf56() bool {
	return _this.Get_().(D13_DC40).Cf56
}

func (_this D13) Dtor_cf57() *C0 {
	return _this.Get_().(D13_DC40).Cf57
}

func (_this D13) Dtor_cf58() _dafny.Map {
	return _this.Get_().(D13_DC40).Cf58
}

func (_this D13) Dtor_cf55() _dafny.Map {
	return _this.Get_().(D13_DC38).Cf55
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC39:
		{
			return "D13.DC39"
		}
	case D13_DC40:
		{
			return "D13.DC40" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC39:
		{
			_, ok := other.Get_().(D13_DC39)
			return ok
		}
	case D13_DC40:
		{
			data2, ok := other.Get_().(D13_DC40)
			return ok && data1.Cf56 == data2.Cf56 && data1.Cf57 == data2.Cf57 && data1.Cf58.Equals(data2.Cf58)
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D14_DC41 struct {
	Cf59 *C4
}

func (D14_DC41) isD14() {}

func (CompanionStruct_D14_) Create_DC41_(Cf59 *C4) D14 {
	return D14{D14_DC41{Cf59}}
}

func (_this D14) Is_DC41() bool {
	_, ok := _this.Get_().(D14_DC41)
	return ok
}

func (CompanionStruct_D14_) Default() *C4 {
	return (*C4)(nil)
}

func (_this D14) Dtor_cf59() *C4 {
	return _this.Get_().(D14_DC41).Cf59
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC41:
		{
			return "D14.DC41" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC41:
		{
			data2, ok := other.Get_().(D14_DC41)
			return ok && data1.Cf59 == data2.Cf59
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
	Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool
	M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool)
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
	Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool
	M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool)
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
	F0  bool
	F2  bool
	F3  _dafny.MultiSet
	F5  _dafny.Int
	F8  _dafny.Set
	_f1 _dafny.Int
	_f4 bool
	_f6 _dafny.MultiSet
	_f7 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F2 = false
	_this.F3 = _dafny.EmptyMultiSet
	_this.F5 = _dafny.Zero
	_this.F8 = _dafny.EmptySet
	_this._f1 = _dafny.Zero
	_this._f4 = false
	_this._f6 = _dafny.EmptyMultiSet
	_this._f7 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 _dafny.MultiSet, f4 bool, f5 _dafny.Int, f6 _dafny.MultiSet, f7 bool, f8 _dafny.Set) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.MultiSet {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f10 bool
	_f11 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f10 = false
	_this._f11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f10 bool, f11 _dafny.Array) {
	{
		(_this)._f10 = f10
		(_this)._f11 = f11
	}
}
func (_this *C0) Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return (_this).F10()
	}
}
func (_this *C0) M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		(globalState).F2 = (_this).F10()
		var _297_v0 _dafny.Sequence
		_ = _297_v0
		_297_v0 = _dafny.UnicodeSeqOfUtf8Bytes("asdhidii")
		var _298_i0 _dafny.Int
		_ = _298_i0
		_298_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsPrefixOf(_297_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sqahgg"), _297_v0)) {
				{
					if (_298_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_298_i0 = (_298_i0).Plus(_dafny.One)
					var _299_v1 D0
					_ = _299_v1
					_299_v1 = Companion_D0_.Create_DC1_((_this).F10())
					var _300_v2 _dafny.Map
					_ = _300_v2
					_300_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_299_v1).Dtor_cf1()) == ((_this).F10()), (_this).F10())
					_300_v2 = (_300_v2).Update((_this).F10(), (_this).F10())
					var _301_v3 _dafny.Array
					_ = _301_v3
					var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
					_ = _nw59
					_301_v3 = _nw59
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_301_v3), 0))
					_ = _index32
					(_301_v3).ArraySet1(p2, (_index32).Int())
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_301_v3), 0))
					_ = _index33
					(_301_v3).ArraySet1(p3, (_index33).Int())
					var _302_v4 _dafny.Sequence
					_ = _302_v4
					_302_v4 = _dafny.SeqOf((_this).F10())
					var _rhs27 bool = (_this).F10()
					_ = _rhs27
					var _rhs28 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_302_v4, _302_v4)
					_ = _rhs28
					var _rhs29 bool = (_this).F10()
					_ = _rhs29
					var _rhs30 bool = (_this).F10()
					_ = _rhs30
					var _lhs26 *GlobalState = globalState
					_ = _lhs26
					var _lhs27 *GlobalState = globalState
					_ = _lhs27
					r0 = _rhs27
					_302_v4 = _rhs28
					_lhs26.F2 = _rhs29
					_lhs27.F2 = _rhs30
					(globalState).F5 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_301_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_301_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_297_v0).Cardinality())), p3)
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _303_v5 _dafny.Array
		_ = _303_v5
		var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
		_ = _nw60
		_303_v5 = _nw60
		var _304_v6 D1
		_ = _304_v6
		_304_v6 = Companion_D1_.Create_DC3_(_303_v5)
		var _305_v7 _dafny.Map
		_ = _305_v7
		_305_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _303_v5)
		var _306_v8 _dafny.Sequence
		_ = _306_v8
		_306_v8 = _dafny.SeqOf(_303_v5, _303_v5)
		var _307_v9 _dafny.Array
		_ = _307_v9
		var _nwElement0_7 _dafny.Array = _303_v5
		_ = _nwElement0_7
		var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(21))
		_ = _nw61
		(_nw61).ArraySet1(_nwElement0_7, 0)
		(_nw61).ArraySet1(_303_v5, 1)
		(_nw61).ArraySet1(_303_v5, 2)
		(_nw61).ArraySet1(_303_v5, 3)
		(_nw61).ArraySet1(_303_v5, 4)
		(_nw61).ArraySet1(_303_v5, 5)
		(_nw61).ArraySet1(_303_v5, 6)
		(_nw61).ArraySet1((_304_v6).Dtor_cf3(), 7)
		(_nw61).ArraySet1(_303_v5, 8)
		(_nw61).ArraySet1(_303_v5, 9)
		(_nw61).ArraySet1(_303_v5, 10)
		(_nw61).ArraySet1(_303_v5, 11)
		(_nw61).ArraySet1((func() _dafny.Array {
			if (_305_v7).Contains((_this).F10()) {
				return (_305_v7).Get((_this).F10()).(_dafny.Array)
			}
			return (_306_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.IntOfUint32((_306_v8).Cardinality()))).Uint32()).(_dafny.Array)
		})(), 12)
		(_nw61).ArraySet1(_303_v5, 13)
		(_nw61).ArraySet1((func() _dafny.Array {
			if (_this).F10() {
				return _303_v5
			}
			return _303_v5
		})(), 14)
		(_nw61).ArraySet1(_303_v5, 15)
		(_nw61).ArraySet1(_303_v5, 16)
		(_nw61).ArraySet1(_303_v5, 17)
		(_nw61).ArraySet1(_303_v5, 18)
		(_nw61).ArraySet1(_303_v5, 19)
		(_nw61).ArraySet1(_303_v5, 20)
		_307_v9 = _nw61
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_307_v9), 0))
		_ = _index34
		(_307_v9).ArraySet1(_303_v5, (_index34).Int())
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_307_v9), 0))
		_ = _index35
		var _rhs31 _dafny.Array = _303_v5
		_ = _rhs31
		var _rhs32 _dafny.Int = Companion_Default___.Fm0(_dafny.IntOfUint32((_297_v0).Cardinality()), p2, p0, ((_this).F10()) == (!((_this).F10())), globalState)
		_ = _rhs32
		var _rhs33 _dafny.Int = p0
		_ = _rhs33
		var _lhs28 _dafny.Array = _307_v9
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_307_v9), 0))
		_ = _lhs29
		var _lhs30 *GlobalState = globalState
		_ = _lhs30
		(_lhs28).ArraySet1(_rhs31, (_lhs29).Int())
		_lhs30.F5 = _rhs32
		r1 = _rhs33
		var _308_v10 _dafny.Set
		_ = _308_v10
		_308_v10 = _dafny.SetOf((_this).F10())
		var _309_v11 _dafny.Set
		_ = _309_v11
		_309_v11 = _dafny.SetOf((_308_v10).Cardinality())
		_309_v11 = Companion_Default___.Fm6(globalState)
		if !(true) {
			(globalState).F0 = (func() bool {
				if (_this).F10() {
					return (_this).F10()
				}
				return !(false)
			})()
			var _310_v12 _dafny.Map
			_ = _310_v12
			_310_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), !((_this).F10()))
			var _311_v13 _dafny.Sequence
			_ = _311_v13
			_311_v13 = _dafny.SeqOf(((func() bool {
				if (_310_v12).Contains((_this).F10()) {
					return (_310_v12).Get((_this).F10()).(bool)
				}
				return (_this).F10()
			})()) == ((_this).F10()))
			_311_v13 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F10()), _dafny.Companion_Sequence_.Concatenate(_311_v13, _311_v13))
			var _312_v14 _dafny.Array
			_ = _312_v14
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_9
			var _nw62 _dafny.Array
			_ = _nw62
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw62 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = (func(_313_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_314_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_314_i1, _313_p3)
					}
				})(p3)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw62 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw62).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw62).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_312_v14 = _nw62
			_312_v14 = _312_v14
			r1 = p0
			(globalState).F0 = (_this).F10()
		} else {
			var _315_v15 _dafny.MultiSet
			_ = _315_v15
			_315_v15 = _dafny.MultiSetOf((_this).F10())
			(globalState).F2 = (_315_v15).IsProperSubsetOf((Companion_Default___.Fm7(globalState)).Union((_315_v15).Update((_this).F10(), Companion_Default___.Abs(p3))))
			r0 = true
			var _316_v16 _dafny.Array
			_ = _316_v16
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_10
			var _nw63 _dafny.Array
			_ = _nw63
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw63 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Int = (func(_317_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_318_i2 _dafny.Int) _dafny.Int {
						return (_318_i2).Minus(_317_p0)
					}
				})(p0)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw63 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw63).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw63).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_316_v16 = _nw63
			_316_v16 = _316_v16
			var _319_v17 _dafny.Sequence
			_ = _319_v17
			_319_v17 = _dafny.SeqOf(false)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index36
			((_this).F11()).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_319_v17, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_319_v17).Cardinality()))).Uint32(), (_319_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_319_v17).Cardinality()), _dafny.IntOfUint32((_319_v17).Cardinality()))).Uint32()).(bool)), _319_v17), (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index37
			((_this).F11()).ArraySet1((p3).Cmp(p2) >= 0, (_index37).Int())
			if (_this).F10() {
				(globalState).F5 = Companion_Default___.Fm0(p3, p3, p2, ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(bool), globalState)
				(globalState).F0 = Companion_Default___.Fm1(false, globalState)
				var _320_v18 _dafny.Map
				_ = _320_v18
				_320_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(bool))
				var _321_v19 _dafny.Map
				_ = _321_v19
				_321_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_320_v18).Cardinality(), Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfInt64(682)))
				_321_v19 = (_321_v19).Update(p0, (func() _dafny.Int {
					if (_321_v19).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("weoxwqp")).Cardinality()))) {
						return (_321_v19).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("weoxwqp")).Cardinality()))).(_dafny.Int)
					}
					return p2
				})())
				(globalState).F2 = (_this).F10()
				var _322_v20 _dafny.Map
				_ = _322_v20
				_322_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v10, (p3).Cmp(p3) == 0)
				_322_v20 = _322_v20
			} else {
				var _323_v21 _dafny.CodePoint
				_ = _323_v21
				_323_v21 = _dafny.CodePoint('h')
				_323_v21 = _323_v21
				var _324_v22 _dafny.Map
				_ = _324_v22
				_324_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
				var _325_v23 _dafny.Map
				_ = _325_v23
				_325_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(bool), (_this).F10())
				var _326_v24 _dafny.Map
				_ = _326_v24
				_326_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(bool), ((func() _dafny.Int {
					if (_324_v22).Contains(Companion_Default___.Fm0(p2, (_325_v23).Cardinality(), p2, (_this).F10(), globalState)) {
						return (_324_v22).Get(Companion_Default___.Fm0(p2, (_325_v23).Cardinality(), p2, (_this).F10(), globalState)).(_dafny.Int)
					}
					return p0
				})()).Times((_dafny.Zero).Minus(p0)))
				_326_v24 = (_326_v24).Update((_this).F10(), (p0).Minus(p0))
				var _327_v25 _dafny.Map
				_ = _327_v25
				_327_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F10())
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_316_v16), 0))
				_ = _index38
				(_316_v16).ArraySet1(_dafny.IntOfInt64(-985), (_index38).Int())
				var _328_v26 D0
				_ = _328_v26
				_328_v26 = Companion_D0_.Create_DC0_(_327_v25)
				var _329_v27 _dafny.Set
				_ = _329_v27
				_329_v27 = _dafny.SetOf(_323_v21)
				var _pat_let_tv1 = _307_v9
				_ = _pat_let_tv1
				var _pat_let_tv2 = _307_v9
				_ = _pat_let_tv2
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_316_v16), 0))
				_ = _index39
				var _rhs34 _dafny.Map = (_328_v26).Dtor_cf0()
				_ = _rhs34
				var _rhs35 _dafny.Int = _dafny.IntOfInt64(699)
				_ = _rhs35
				var _rhs36 D1 = func(_pat_let2_0 D1) D1 {
					return func(_330_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let3_0 _dafny.Array) D1 {
							return func(_331_dt__update_hcf3_h0 _dafny.Array) D1 {
								return Companion_D1_.Create_DC3_(_331_dt__update_hcf3_h0)
							}(_pat_let3_0)
						}(_dafny.ArrayCastTo((_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_pat_let_tv1), 0))).Int())))
					}(_pat_let2_0)
				}(_304_v6)
				_ = _rhs36
				var _rhs37 bool = (func() bool {
					if (_329_v27).IsProperSubsetOf(_329_v27) {
						return (_this).F10()
					}
					return ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(bool)
				})()
				_ = _rhs37
				var _lhs31 _dafny.Array = _316_v16
				_ = _lhs31
				var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_316_v16), 0))
				_ = _lhs32
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				_327_v25 = _rhs34
				(_lhs31).ArraySet1(_rhs35, (_lhs32).Int())
				_304_v6 = _rhs36
				_lhs33.F0 = _rhs37
				var _332_v28 _dafny.Map
				_ = _332_v28
				_332_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(p3), p2)
				var _333_v29 D1
				_ = _333_v29
				_333_v29 = Companion_D1_.Create_DC4_((_316_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_316_v16), 0))).Int()).(_dafny.Int))
				var _334_v30 _dafny.MultiSet
				_ = _334_v30
				_334_v30 = _dafny.MultiSetOf(_dafny.IntOfInt64(675))
				_332_v28 = (_332_v28).Update(_333_v29, Companion_Default___.Fm0(_dafny.IntOfInt64(492), p2, (_334_v30).Cardinality(), false, globalState))
				(globalState).F2 = true
			}
		}
		var _335_v31 _dafny.Sequence
		_ = _335_v31
		_335_v31 = _dafny.SeqOf(p0)
		var _336_v32 _dafny.Map
		_ = _336_v32
		_336_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_335_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_335_v31).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(p0))
		r1 = ((_336_v32).Cardinality()).Minus(((_dafny.Zero).Minus(p3)).Minus(p3))
		r0 = (_this).F10()
		var _337_v33 _dafny.Map
		_ = _337_v33
		_337_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
		var _338_v34 D0
		_ = _338_v34
		_338_v34 = Companion_D0_.Create_DC0_(_337_v33)
		var _339_v35 _dafny.MultiSet
		_ = _339_v35
		_339_v35 = _dafny.MultiSetOf((_this).F10(), (_this).F10())
		var _340_v36 _dafny.MultiSet
		_ = _340_v36
		_340_v36 = _dafny.MultiSetOf((_this).F10(), (_this).Fm4(_338_v34, _338_v34, _339_v35, globalState), (_this).F10())
		var _341_v37 _dafny.Map
		_ = _341_v37
		_341_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(_338_v34, _338_v34, (_339_v35).Update((_this).F10(), Companion_Default___.Abs(p0)), globalState), p3)
		r1 = (func() _dafny.Int {
			if (_341_v37).Contains((_this).F10()) {
				return (_341_v37).Get((_this).F10()).(_dafny.Int)
			}
			return p3
		})()
		r2 = ((_this).F10()) && ((func() bool {
			if (_337_v33).Contains(p3) {
				return (_337_v33).Get(p3).(bool)
			}
			return (_this).Fm4(_338_v34, _338_v34, _340_v36, globalState)
		})())
		var _342_v38 _dafny.CodePoint
		_ = _342_v38
		_342_v38 = _dafny.CodePoint('w')
		var _343_v39 _dafny.Map
		_ = _343_v39
		_343_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _342_v38)
		r3 = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm8((_this).F10(), _343_v39, (_this).F10(), globalState), _297_v0)
		return r0, r1, r2, r3
	}
}
func (_this *C0) F10() bool {
	{
		return _this._f10
	}
}
func (_this *C0) F11() _dafny.Array {
	{
		return _this._f11
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		var _source5 D1 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(-738))
		_ = _source5
		if _source5.Is_DC4() {
			var _344___mcc_h0 _dafny.Int = _source5.Get_().(D1_DC4).Cf4
			_ = _344___mcc_h0
			var _345_cf4 _dafny.Int = _344___mcc_h0
			_ = _345_cf4
			return true
		} else {
			var _346___mcc_h1 _dafny.Array = _source5.Get_().(D1_DC3).Cf3
			_ = _346___mcc_h1
			var _347_cf3 _dafny.Array = _346___mcc_h1
			_ = _347_cf3
			return true
		}
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _348_v0 bool
		_ = _348_v0
		_348_v0 = true
		var _349_v1 _dafny.Array
		_ = _349_v1
		var _nw64 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw64
		_349_v1 = _nw64
		var _350_v2 *C0
		_ = _350_v2
		var _nw65 *C0 = New_C0_()
		_ = _nw65
		_nw65.Ctor__(_348_v0, _349_v1)
		_350_v2 = _nw65
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))
		_ = _index40
		((_350_v2).F11()).ArraySet1(true, (_index40).Int())
		var _351_v3 _dafny.Map
		_ = _351_v3
		_351_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_350_v2).F10())
		var _352_v4 D0
		_ = _352_v4
		_352_v4 = Companion_D0_.Create_DC0_(_351_v3)
		var _353_v5 D0
		_ = _353_v5
		_353_v5 = Companion_D0_.Create_DC0_((_352_v4).Dtor_cf0())
		var _354_v6 _dafny.MultiSet
		_ = _354_v6
		_354_v6 = _dafny.MultiSetOf(Companion_Default___.Fm1(_348_v0, globalState))
		var _355_v7 _dafny.Map
		_ = _355_v7
		_355_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _354_v6)
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))
		_ = _index41
		((_350_v2).F11()).ArraySet1((_350_v2).Fm4(_352_v4, _353_v5, (func() _dafny.MultiSet {
			if (_355_v7).Contains(p3) {
				return (_355_v7).Get(p3).(_dafny.MultiSet)
			}
			return _354_v6
		})(), globalState), (_index41).Int())
		var _356_v8 _dafny.Sequence
		_ = _356_v8
		_356_v8 = _dafny.SeqOf(p0)
		var _357_v9 _dafny.Sequence
		_ = _357_v9
		_357_v9 = _356_v8
		(globalState).F5 = Companion_Default___.Fm0(p3, p2, _dafny.IntOfInt64(474), _dafny.Companion_Sequence_.IsPrefixOf((_357_v9), _356_v8), globalState)
		var _hi1 _dafny.Int = p0
		_ = _hi1
		for _358_i0 := p2; _358_i0.Cmp(_hi1) < 0; _358_i0 = _358_i0.Plus(_dafny.One) {
			var _359_v10 D7
			_ = _359_v10
			_359_v10 = Companion_D7_.Create_DC17_(_dafny.SeqOf((_350_v2).F10()))
			var _360_v11 _dafny.Sequence
			_ = _360_v11
			_360_v11 = _dafny.SeqOf(((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool), (_350_v2).F10(), ((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool), (_this).Fm4(Companion_D0_.Create_DC0_(Companion_Default___.Fm22((_350_v2).F10(), globalState)), _352_v4, _354_v6, globalState), true)
			var _361_v12 _dafny.Set
			_ = _361_v12
			_361_v12 = _dafny.SetOf(_360_v11)
			if (_361_v12).Contains((_359_v10).Dtor_cf23()) {
				(globalState).F0 = (Companion_D2_.Create_DC6_((_350_v2).F10())).Dtor_cf6()
				r2 = (p3).Cmp(p0) >= 0
				var _362_v13 _dafny.Map
				_ = _362_v13
				_362_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_v0, _348_v0), p0)
				var _363_v14 D2
				_ = _363_v14
				_363_v14 = Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_360_v11).Cardinality()), _362_v13, _358_i0)
				(globalState).F5 = (_363_v14).Dtor_cf9()
				var _364_v15 _dafny.MultiSet
				_ = _364_v15
				_364_v15 = _dafny.MultiSetOf(_dafny.IntOfInt64(493))
				_364_v15 = Companion_Default___.Fm23((_364_v15).Update(Companion_Default___.Fm0(p0, p3, p2, (_350_v2).F10(), globalState), Companion_Default___.Abs(p3)), p3, globalState)
				var _rhs38 _dafny.Int = p2
				_ = _rhs38
				var _rhs39 bool = _348_v0
				_ = _rhs39
				var _rhs40 _dafny.Array = (_350_v2).F11()
				_ = _rhs40
				var _rhs41 _dafny.Array = _349_v1
				_ = _rhs41
				var _rhs42 _dafny.Int = ((_351_v3).Merge(_351_v3)).Cardinality()
				_ = _rhs42
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				r1 = _rhs38
				_lhs34.F0 = _rhs39
				_349_v1 = _rhs40
				_349_v1 = _rhs41
				r1 = _rhs42
			} else {
				var _365_v16 _dafny.Map
				_ = _365_v16
				_365_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _349_v1)
				_365_v16 = ((_365_v16).Merge(_365_v16)).Merge(_365_v16)
				r2 = _348_v0
				var _366_v17 T1
				_ = _366_v17
				var _nw66 *C0 = New_C0_()
				_ = _nw66
				_nw66.Ctor__((_350_v2).F10(), (_350_v2).F11())
				_366_v17 = _nw66
				var _367_v18 _dafny.Set
				_ = _367_v18
				_367_v18 = _dafny.SetOf(_366_v17)
				var _368_v19 D8
				_ = _368_v19
				_368_v19 = Companion_D8_.Create_DC20_(_367_v18)
				var _rhs43 _dafny.Array = (func() _dafny.Array {
					if !((_350_v2).F10()) {
						return (_350_v2).F11()
					}
					return (_350_v2).F11()
				})()
				_ = _rhs43
				var _rhs44 bool = (func() bool {
					if !((_350_v2).F10()) {
						return (_367_v18).IsSubsetOf((_368_v19).Dtor_cf25())
					}
					return (_350_v2).F10()
				})()
				_ = _rhs44
				var _rhs45 bool = (_350_v2).F10()
				_ = _rhs45
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				_349_v1 = _rhs43
				_lhs35.F2 = _rhs44
				_lhs36.F2 = _rhs45
				var _369_v20 *C0
				_ = _369_v20
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__(((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool), (_350_v2).F11())
				_369_v20 = _nw67
				r0 = ((_350_v2).F10()) && (true)
			}
			var _370_v21 *C0
			_ = _370_v21
			var _nw68 *C0 = New_C0_()
			_ = _nw68
			_nw68.Ctor__(((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool), (_350_v2).F11())
			_370_v21 = _nw68
			(globalState).F5 = (_dafny.MultiSetOf((_350_v2).F10())).Cardinality()
			var _371_v22 _dafny.Array
			_ = _371_v22
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw69
			_371_v22 = _nw69
			var _372_v23 _dafny.Array
			_ = _372_v23
			_372_v23 = _371_v22
			var _373_v24 _dafny.Map
			_ = _373_v24
			_373_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_350_v2).F10() {
					return _372_v23
				}
				return _372_v23
			})(), (Companion_D4_.Create_DC13_(p3, _349_v1)).Dtor_cf19())
			_373_v24 = (_373_v24).Update(_372_v23, Companion_Default___.SafeDivisionInt(p0, p3))
		}
		var _374_v25 D2
		_ = _374_v25
		_374_v25 = Companion_D2_.Create_DC6_(false)
		var _375_i1 _dafny.Int
		_ = _375_i1
		_375_i1 = _dafny.Zero
		{
			for (_374_v25).Dtor_cf6() {
				{
					if (_375_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_375_i1 = (_375_i1).Plus(_dafny.One)
					var _376_v26 *C0
					_ = _376_v26
					var _nw70 *C0 = New_C0_()
					_ = _nw70
					_nw70.Ctor__(_348_v0, _349_v1)
					_376_v26 = _nw70
					(globalState).F2 = (Companion_Default___.Fm0(p0, p0, p2, (_350_v2).F10(), globalState)).Cmp((_351_v3).Cardinality()) != 0
					_351_v3 = (_351_v3).Update(p3, ((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool))
					var _377_v27 _dafny.CodePoint
					_ = _377_v27
					_377_v27 = _dafny.CodePoint('m')
					_377_v27 = _377_v27
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _378_v28 _dafny.Sequence
		_ = _378_v28
		_378_v28 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_356_v8))
		var _pat_let_tv3 = _348_v0
		_ = _pat_let_tv3
		var _pat_let_tv4 = p0
		_ = _pat_let_tv4
		var _pat_let_tv5 = _350_v2
		_ = _pat_let_tv5
		var _pat_let_tv6 = _350_v2
		_ = _pat_let_tv6
		var _pat_let_tv7 = _350_v2
		_ = _pat_let_tv7
		var _pat_let_tv8 = _350_v2
		_ = _pat_let_tv8
		var _pat_let_tv9 = p0
		_ = _pat_let_tv9
		var _rhs46 bool = !(!(_348_v0) || (_348_v0))
		_ = _rhs46
		var _rhs47 _dafny.Int = Companion_Default___.SafeDivisionInt(p3, _dafny.IntOfInt64(652))
		_ = _rhs47
		var _rhs48 bool = ((_354_v6).Union(_354_v6)).IsDisjointFrom((_354_v6).Update(_348_v0, Companion_Default___.Abs((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_378_v28)).Cardinality()))))
		_ = _rhs48
		var _rhs49 _dafny.Int = func(_source6 D0) _dafny.Int {
			if _source6.Is_DC1() {
				var _379___mcc_h0 bool = _source6.Get_().(D0_DC1).Cf1
				_ = _379___mcc_h0
				var _380_cf1 bool = _379___mcc_h0
				_ = _380_cf1
				return _dafny.IntOfInt64(527)
			} else if _source6.Is_DC0() {
				var _381___mcc_h1 _dafny.Map = _source6.Get_().(D0_DC0).Cf0
				_ = _381___mcc_h1
				var _382_cf0 _dafny.Map = _381___mcc_h1
				_ = _382_cf0
				return ((_dafny.SetOf(_pat_let_tv3)).Cardinality()).Times(_pat_let_tv4)
			} else {
				var _383___mcc_h2 D0 = _source6.Get_().(D0_DC2).Cf2
				_ = _383___mcc_h2
				var _384_cf2 D0 = _383___mcc_h2
				_ = _384_cf2
				return Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_pat_let_tv6).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_pat_let_tv5).F11()), 0))).Int()).(bool), ((_pat_let_tv8).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_pat_let_tv7).F11()), 0))).Int()).(bool))).Cardinality(), _pat_let_tv9)
			}
		}(_352_v4)
		_ = _rhs49
		var _rhs50 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(p0, p2), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p3, p2)))
		_ = _rhs50
		var _lhs37 *GlobalState = globalState
		_ = _lhs37
		r3 = _rhs46
		_lhs37.F5 = _rhs47
		_348_v0 = _rhs48
		r1 = _rhs49
		r1 = _rhs50
		var _385_v29 _dafny.Map
		_ = _385_v29
		_385_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool), false)
		var _386_v30 _dafny.Map
		_ = _386_v30
		_386_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_385_v29, p0)
		var _387_v31 _dafny.Sequence
		_ = _387_v31
		_387_v31 = _dafny.SeqOf(((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool))
		var _388_v32 D2
		_ = _388_v32
		_388_v32 = Companion_D2_.Create_DC7_(p3, _386_v30, _dafny.IntOfUint32((_387_v31).Cardinality()))
		r0 = (p3).Cmp((_388_v32).Dtor_cf9()) > 0
		r1 = p2
		r2 = !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_348_v0, _dafny.IntOfUint32((_387_v31).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(Companion_D0_.Create_DC0_(_351_v3), _353_v5, _354_v6, globalState), Companion_Default___.Fm0(p2, p0, p2, ((_350_v2).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen(((_350_v2).F11()), 0))).Int()).(bool), globalState)))).Contains(true)
		r3 = _348_v0
		return r0, r1, r2, r3
	}
}
func (_this *C1) M6(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) (_dafny.Int, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		(globalState).F0 = p2
		var _389_v0 _dafny.Sequence
		_ = _389_v0
		_389_v0 = _dafny.UnicodeSeqOfUtf8Bytes("igpdhpy")
		var _390_v1 _dafny.Set
		_ = _390_v1
		_390_v1 = _dafny.SetOf(_389_v0)
		_390_v1 = _390_v1
		if p2 {
			var _391_v2 D2
			_ = _391_v2
			_391_v2 = Companion_D2_.Create_DC6_(p2)
			var _392_v3 _dafny.Sequence
			_ = _392_v3
			_392_v3 = _dafny.SeqOf((_391_v2).Dtor_cf6(), p2, p0, p2, !(p0))
			var _393_v4 _dafny.MultiSet
			_ = _393_v4
			_393_v4 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_392_v3).Cardinality()), p1))
			var _394_v5 _dafny.Map
			_ = _394_v5
			_394_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _395_v6 _dafny.Array
			_ = _395_v6
			var _nwElement0_8 bool = p0
			_ = _nwElement0_8
			var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(11))
			_ = _nw71
			(_nw71).ArraySet1(_nwElement0_8, 0)
			(_nw71).ArraySet1(((func() bool {
				if (_394_v5).Contains(p2) {
					return (_394_v5).Get(p2).(bool)
				}
				return p0
			})()) && (p2), 1)
			(_nw71).ArraySet1((func() bool {
				if p2 {
					return p0
				}
				return p2
			})(), 2)
			(_nw71).ArraySet1(true, 3)
			(_nw71).ArraySet1(p2, 4)
			(_nw71).ArraySet1(p2, 5)
			(_nw71).ArraySet1(p0, 6)
			(_nw71).ArraySet1(p0, 7)
			(_nw71).ArraySet1(p0, 8)
			(_nw71).ArraySet1(p2, 9)
			(_nw71).ArraySet1(p2, 10)
			_395_v6 = _nw71
			var _396_v7 _dafny.Map
			_ = _396_v7
			_396_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_392_v3).Select((Companion_Default___.SafeIndex((_393_v4).Cardinality(), _dafny.IntOfUint32((_392_v3).Cardinality()))).Uint32()).(bool), _392_v3)
			var _rhs51 _dafny.MultiSet = _393_v4
			_ = _rhs51
			var _rhs52 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p0 {
					return _389_v0
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("inuynen")
			})()).Cardinality())
			_ = _rhs52
			var _rhs53 bool = !((!(p0)) && (p2))
			_ = _rhs53
			var _rhs54 bool = ((p1).Times(p1)).Cmp((_396_v7).Cardinality()) >= 0
			_ = _rhs54
			var _rhs55 _dafny.Array = _395_v6
			_ = _rhs55
			var _lhs38 *GlobalState = globalState
			_ = _lhs38
			var _lhs39 *GlobalState = globalState
			_ = _lhs39
			var _lhs40 *GlobalState = globalState
			_ = _lhs40
			_393_v4 = _rhs51
			_lhs38.F5 = _rhs52
			_lhs39.F0 = _rhs53
			_lhs40.F0 = _rhs54
			_395_v6 = _rhs55
			var _397_v8 _dafny.Array
			_ = _397_v8
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_11
			var _nw72 _dafny.Array
			_ = _nw72
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw72 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Int = (func(_398_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_399_i0 _dafny.Int) _dafny.Int {
						return (_399_i0).Times(_398_p1)
					}
				})(p1)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw72 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw72).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw72).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_397_v8 = _nw72
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
			_ = _index42
			(_397_v8).ArraySet1(p1, (_index42).Int())
			var _400_v10 _dafny.Array
			_ = _400_v10
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_12
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Map = (func(_401_p0 bool, _402_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_403_i1 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_p0, (_dafny.SetOf(_dafny.SetOf(_402_p1), func() _dafny.Set {
							var _coll21 = _dafny.NewBuilder()
							_ = _coll21
							for _iter21 := _dafny.Iterate((_dafny.SeqOf(_402_p1)).Elements()); ; {
								_compr_21, _ok21 := _iter21()
								if !_ok21 {
									break
								}
								var _404_v9 _dafny.Int
								_404_v9 = interface{}(_compr_21).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_402_p1), _404_v9) {
									_coll21.Add(Companion_Default___.SafeModuloInt(_404_v9, _dafny.IntOfInt64(718)))
								}
							}
							return _coll21.ToSet()
						}())).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_p0, _402_p1)))
					}
				})(p0, p1)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw73 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw73).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw73).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_400_v10 = _nw73
			var _405_v11 _dafny.Map
			_ = _405_v11
			_405_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_v6, p0)
			var _406_v12 D9
			_ = _406_v12
			_406_v12 = Companion_D9_.Create_DC23_(_405_v11)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_400_v10), 0))
			_ = _index43
			(_400_v10).ArraySet1(Companion_Default___.Fm24(((_406_v12).Dtor_cf32()).Cardinality(), globalState), (_index43).Int())
			var _407_v13 _dafny.Set
			_ = _407_v13
			_407_v13 = _dafny.SetOf(Companion_Default___.Fm25((func() bool {
				if (_394_v5).Contains(p2) {
					return (_394_v5).Get(p2).(bool)
				}
				return p2
			})(), _dafny.CodePoint('n'), p0, globalState), _392_v3, _392_v3)
			var _408_v14 _dafny.Map
			_ = _408_v14
			_408_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
			var _409_v15 _dafny.Map
			_ = _409_v15
			_409_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _408_v14)
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
			_ = _index44
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_400_v10), 0))
			_ = _index45
			var _rhs56 _dafny.Int = (_407_v13).Cardinality()
			_ = _rhs56
			var _rhs57 _dafny.Sequence = _392_v3
			_ = _rhs57
			var _rhs58 _dafny.Map = (_409_v15).Merge(_409_v15)
			_ = _rhs58
			var _lhs41 _dafny.Array = _397_v8
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
			_ = _lhs42
			var _lhs43 _dafny.Array = _400_v10
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_400_v10), 0))
			_ = _lhs44
			(_lhs41).ArraySet1(_rhs56, (_lhs42).Int())
			_392_v3 = _rhs57
			(_lhs43).ArraySet1(_rhs58, (_lhs44).Int())
			var _410_v16 _dafny.Sequence
			_ = _410_v16
			_410_v16 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bodi")).Cardinality()), p1, p1, _dafny.IntOfInt64(-446), (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int))
			var _411_v17 _dafny.Set
			_ = _411_v17
			_411_v17 = _dafny.SetOf(!(Companion_Default___.Fm1(p2, globalState)))
			if (((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int)).Minus((_410_v16).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_410_v16).Cardinality()))).Uint32()).(_dafny.Int))).Cmp((p1).Minus((_dafny.Zero).Minus((_411_v17).Cardinality()))) != 0 {
				var _412_v18 D4
				_ = _412_v18
				_412_v18 = Companion_D4_.Create_DC13_((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _395_v6)
				_412_v18 = Companion_D4_.Create_DC13_((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _395_v6)
				(globalState).F2 = ((_dafny.Zero).Minus(p1)).Cmp(((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int)).Times((_394_v5).Cardinality())) < 0
				var _413_v19 _dafny.Map
				_ = _413_v19
				_413_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _395_v6)
				_413_v19 = (_413_v19).Update(p1, _395_v6)
				var _414_v20 _dafny.Map
				_ = _414_v20
				_414_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), !(!((p2) && (p0))))
				var _pat_let_tv10 = p2
				_ = _pat_let_tv10
				var _rhs59 bool = (func() bool {
					if (_414_v20).Contains(Companion_Default___.Fm0((_411_v17).Cardinality(), (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-119), p2, globalState)) {
						return (_414_v20).Get(Companion_Default___.Fm0((_411_v17).Cardinality(), (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-119), p2, globalState)).(bool)
					}
					return p0
				})()
				_ = _rhs59
				var _rhs60 _dafny.Int = _dafny.IntOfInt64(437)
				_ = _rhs60
				var _rhs61 bool = (func(_pat_let4_0 D2) D2 {
					return func(_415_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let5_0 bool) D2 {
							return func(_416_dt__update_hcf6_h0 bool) D2 {
								return Companion_D2_.Create_DC6_(_416_dt__update_hcf6_h0)
							}(_pat_let5_0)
						}(_pat_let_tv10)
					}(_pat_let4_0)
				}(_391_v2)).Dtor_cf6()
				_ = _rhs61
				var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_389_v0, _dafny.Companion_Sequence_.Concatenate(_389_v0, _389_v0))
				_ = _rhs62
				var _lhs45 *GlobalState = globalState
				_ = _lhs45
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				_lhs45.F2 = _rhs59
				_lhs46.F5 = _rhs60
				_lhs47.F0 = _rhs61
				_389_v0 = _rhs62
				var _417_v21 *C0
				_ = _417_v21
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__(p0, _395_v6)
				_417_v21 = _nw74
			} else {
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_395_v6), 0))
				_ = _index46
				(_395_v6).ArraySet1(p2, (_index46).Int())
				var _418_v22 _dafny.Array
				_ = _418_v22
				_418_v22 = _397_v8
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_395_v6), 0))
				_ = _index47
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
				_ = _index48
				var _rhs63 _dafny.Array = _395_v6
				_ = _rhs63
				var _rhs64 _dafny.Int = (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int)
				_ = _rhs64
				var _rhs65 bool = p0
				_ = _rhs65
				var _rhs66 _dafny.Int = p1
				_ = _rhs66
				var _rhs67 _dafny.Array = (_418_v22)
				_ = _rhs67
				var _lhs48 *GlobalState = globalState
				_ = _lhs48
				var _lhs49 _dafny.Array = _395_v6
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_395_v6), 0))
				_ = _lhs50
				var _lhs51 _dafny.Array = _397_v8
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
				_ = _lhs52
				_395_v6 = _rhs63
				_lhs48.F5 = _rhs64
				(_lhs49).ArraySet1(_rhs65, (_lhs50).Int())
				(_lhs51).ArraySet1(_rhs66, (_lhs52).Int())
				_397_v8 = _rhs67
				var _419_v23 _dafny.MultiSet
				_ = _419_v23
				_419_v23 = _dafny.MultiSetOf((func() bool {
					if p0 {
						return (_391_v2).Dtor_cf6()
					}
					return (_392_v3).Select((Companion_Default___.SafeIndex((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_392_v3).Cardinality()))).Uint32()).(bool)
				})())
				var _420_v24 D0
				_ = _420_v24
				_420_v24 = Companion_D0_.Create_DC0_(Companion_Default___.Fm22(p2, globalState))
				var _rhs68 _dafny.Int = (_dafny.Zero).Minus((_419_v23).Cardinality())
				_ = _rhs68
				var _rhs69 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((Companion_Default___.Fm26((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(901), false, (_this).Fm4(_420_v24, _420_v24, _419_v23, globalState), globalState)).Dtor_cf9()), p1)
				_ = _rhs69
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				_lhs53.F5 = _rhs68
				_lhs54.F5 = _rhs69
				var _421_v26 _dafny.Map
				_ = _421_v26
				_421_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_389_v0).Cardinality()), _393_v4)
				var _422_v27 _dafny.Map
				_ = _422_v27
				_422_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
					if (_393_v4).Contains(p1) {
						return (_393_v4).Multiplicity(p1)
					}
					return (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int)
				})()).Times((func() _dafny.Map {
					var _coll22 = _dafny.NewMapBuilder()
					_ = _coll22
					for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(850), _dafny.IntOfInt64(991))); ; {
						_compr_22, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _423_v25 _dafny.Int
						_423_v25 = interface{}(_compr_22).(_dafny.Int)
						if ((_dafny.IntOfInt64(850)).Cmp(_423_v25) <= 0) && ((_423_v25).Cmp(_dafny.IntOfInt64(991)) < 0) {
							_coll22.Add(Companion_Default___.SafeModuloInt(_423_v25, _dafny.IntOfInt64(416)), _dafny.IntOfInt64(132))
						}
					}
					return _coll22.ToMap()
				}()).Cardinality()), ((func() _dafny.MultiSet {
					if (_421_v26).Contains((_dafny.MultiSetOf((_395_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_395_v6), 0))).Int()).(bool), p0)).Cardinality()) {
						return (_421_v26).Get((_dafny.MultiSetOf((_395_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_395_v6), 0))).Int()).(bool), p0)).Cardinality()).(_dafny.MultiSet)
					}
					return _393_v4
				})()).Intersection(_393_v4))
				_422_v27 = (_422_v27).Update((p1).Minus(p1), _393_v4)
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_395_v6), 0))
				_ = _index49
				(_395_v6).ArraySet1((Companion_Default___.SafeModuloInt(p1, p1)).Cmp(((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int)).Plus(p1)) > 0, (_index49).Int())
				var _424_v28 _dafny.Array
				_ = _424_v28
				var _len0_13 _dafny.Int = _dafny.One
				_ = _len0_13
				var _nw75 _dafny.Array
				_ = _nw75
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw75 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Map = (func(_425_p2 bool, _426_p1 _dafny.Int, _427_v14 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_428_i2 _dafny.Int) _dafny.Map {
							return (func() _dafny.Map {
								if false {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_p2, _426_p1)
								}
								return _427_v14
							})()
						}
					})(p2, p1, _408_v14)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw75 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw75).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw75).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_424_v28 = _nw75
				var _429_v29 _dafny.Array
				_ = _429_v29
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw76
				_429_v29 = _nw76
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
				_ = _index50
				var _rhs70 _dafny.Array = _424_v28
				_ = _rhs70
				var _rhs71 _dafny.Int = (Companion_Default___.Fm0((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), p1, _dafny.IntOfUint32((_389_v0).Cardinality()), !(p0), globalState)).Plus(_dafny.IntOfUint32((_389_v0).Cardinality()))
				_ = _rhs71
				var _rhs72 _dafny.Int = (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int)
				_ = _rhs72
				var _rhs73 _dafny.Int = (p1).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_410_v16, _410_v16)).Cardinality()))
				_ = _rhs73
				var _rhs74 _dafny.Array = _429_v29
				_ = _rhs74
				var _lhs55 _dafny.Array = _397_v8
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))
				_ = _lhs56
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				_424_v28 = _rhs70
				(_lhs55).ArraySet1(_rhs71, (_lhs56).Int())
				r0 = _rhs72
				_lhs57.F5 = _rhs73
				_429_v29 = _rhs74
			}
			var _430_v30 _dafny.Map
			_ = _430_v30
			_430_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p2), (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int))
			var _431_v31 D2
			_ = _431_v31
			_431_v31 = Companion_D2_.Create_DC7_(p1, _430_v30, (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int))
			var _432_v32 _dafny.Array
			_ = _432_v32
			var _nwElement0_9 D2 = _431_v31
			_ = _nwElement0_9
			var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
			_ = _nw77
			(_nw77).ArraySet1(_nwElement0_9, 0)
			(_nw77).ArraySet1(_431_v31, 1)
			(_nw77).ArraySet1(_431_v31, 2)
			(_nw77).ArraySet1(Companion_Default___.Fm26((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), (_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), p2, !(true), globalState), 3)
			(_nw77).ArraySet1(Companion_D2_.Create_DC7_(p1, _430_v30, _dafny.IntOfInt64(587)), 4)
			(_nw77).ArraySet1(Companion_D2_.Create_DC7_((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _430_v30, _dafny.IntOfUint32((_389_v0).Cardinality())), 5)
			(_nw77).ArraySet1(Companion_Default___.Fm26(p1, p1, !(!(p0)), p2, globalState), 6)
			_432_v32 = _nw77
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_432_v32), 0))
			_ = _index51
			(_432_v32).ArraySet1(Companion_D2_.Create_DC7_((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), _430_v30, p1), (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_432_v32), 0))
			_ = _index52
			(_432_v32).ArraySet1(_431_v31, (_index52).Int())
			var _433_v33 D4
			_ = _433_v33
			_433_v33 = Companion_D4_.Create_DC14_()
			var _434_v34 _dafny.Sequence
			_ = _434_v34
			_434_v34 = _dafny.SeqOf(_433_v33, Companion_D4_.Create_DC14_(), _433_v33)
			var _435_v35 _dafny.CodePoint
			_ = _435_v35
			_435_v35 = _dafny.CodePoint('s')
			var _436_v36 _dafny.Map
			_ = _436_v36
			_436_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _395_v6)
			var _437_v37 _dafny.Set
			_ = _437_v37
			_437_v37 = _dafny.SetOf((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int), ((_436_v36).Cardinality()).Times((_dafny.Zero).Minus((_397_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_397_v8), 0))).Int()).(_dafny.Int))), _dafny.IntOfInt64(350), p1, _dafny.IntOfInt64(-914))
			var _rhs75 _dafny.Sequence = _dafny.SeqOf(_433_v33, _433_v33, Companion_D4_.Create_DC14_(), _433_v33, _433_v33)
			_ = _rhs75
			var _rhs76 _dafny.Sequence = Companion_Default___.Fm27(globalState)
			_ = _rhs76
			var _rhs77 _dafny.CodePoint = _435_v35
			_ = _rhs77
			var _rhs78 _dafny.Set = (_437_v37).Difference((_437_v37).Difference(_437_v37))
			_ = _rhs78
			var _rhs79 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("cps")
			_ = _rhs79
			_434_v34 = _rhs75
			_389_v0 = _rhs76
			_435_v35 = _rhs77
			_437_v37 = _rhs78
			_389_v0 = _rhs79
		} else {
			var _438_v38 _dafny.Map
			_ = _438_v38
			_438_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _439_v39 D0
			_ = _439_v39
			_439_v39 = Companion_D0_.Create_DC1_(p2)
			var _440_v40 _dafny.Map
			_ = _440_v40
			_440_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_438_v38).Cardinality(), _439_v39)
			var _441_v41 D0
			_ = _441_v41
			_441_v41 = Companion_D0_.Create_DC2_((func() D0 {
				if (_440_v40).Contains(p1) {
					return (_440_v40).Get(p1).(D0)
				}
				return _439_v39
			})())
			_441_v41 = _441_v41
			var _442_v42 _dafny.Map
			_ = _442_v42
			_442_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.IntOfInt64(941)).Plus(p1)), (_dafny.Zero).Minus(p1))
			var _443_v43 _dafny.Set
			_ = _443_v43
			_443_v43 = _dafny.SetOf(p0, p2, p0, p2)
			_442_v42 = (_442_v42).Update((_443_v43).Cardinality(), (_dafny.SetOf(p2)).Cardinality())
			var _444_v45 _dafny.Set
			_ = _444_v45
			_444_v45 = _dafny.SetOf(p1, _dafny.IntOfInt64(584))
			r2 = ((_444_v45).Union(_444_v45)).IsProperSubsetOf((func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(508), _dafny.IntOfInt64(914))); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _445_v44 _dafny.Int
					_445_v44 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(508)).Cmp(_445_v44) <= 0) && ((_445_v44).Cmp(_dafny.IntOfInt64(914)) < 0) {
						_coll23.Add((_445_v44).Plus((_dafny.Zero).Minus(p1)))
					}
				}
				return _coll23.ToSet()
			}()).Union(_444_v45))
			r0 = p1
			var _446_v47 _dafny.Sequence
			_ = _446_v47
			_446_v47 = _dafny.SeqOf((func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(807), _dafny.IntOfInt64(152))); ; {
					_compr_24, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _447_v46 _dafny.Int
					_447_v46 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(807)).Cmp(_447_v46) <= 0) && ((_447_v46).Cmp(_dafny.IntOfInt64(152)) < 0) {
						_coll24.Add((_447_v46).Times(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality())))
					}
				}
				return _coll24.ToSet()
			}()).Difference(_444_v45))
			(globalState).F5 = ((_446_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_389_v0).Cardinality()), _dafny.IntOfUint32((_446_v47).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
		}
		(globalState).F0 = (func() bool {
			if p2 {
				return p0
			}
			return p0
		})()
		var _448_v48 _dafny.Array
		_ = _448_v48
		var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw78
		_448_v48 = _nw78
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_448_v48), 0))
		_ = _index53
		(_448_v48).ArraySet1(p1, (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_448_v48), 0))
		_ = _index54
		(_448_v48).ArraySet1(p1, (_index54).Int())
		r0 = p1
		var _449_v49 _dafny.Map
		_ = _449_v49
		_449_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_448_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_448_v48), 0))).Int()).(_dafny.Int))
		var _450_v50 _dafny.Sequence
		_ = _450_v50
		_450_v50 = _dafny.SeqOf(_dafny.IntOfInt64(321))
		var _451_v51 _dafny.Map
		_ = _451_v51
		_451_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_450_v50).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_450_v50).Cardinality()))).Uint32()).(_dafny.Int), false)
		r0 = (Companion_Default___.Fm2(p0, true, p2, (_449_v49).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_451_v51).Contains((_448_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_448_v48), 0))).Int()).(_dafny.Int)) {
				return (_451_v51).Get((_448_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_448_v48), 0))).Int()).(_dafny.Int)).(bool)
			}
			return p0
		})(), p1)), globalState)).Cardinality()
		r1 = p0
		r2 = !((p2) || ((func() bool {
			if p0 {
				return p2
			}
			return p0
		})()))
		return r0, r1, r2
	}
}
func (_this *C1) M7(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, _dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _452_v0 _dafny.CodePoint
		_ = _452_v0
		_452_v0 = _dafny.CodePoint('u')
		var _453_v1 D0
		_ = _453_v1
		_453_v1 = Companion_D0_.Create_DC1_(p1)
		var _454_v2 D2
		_ = _454_v2
		_454_v2 = Companion_D2_.Create_DC6_((_453_v1).Dtor_cf1())
		var _455_v3 _dafny.Map
		_ = _455_v3
		_455_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, false)
		var _456_v4 _dafny.Map
		_ = _456_v4
		_456_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() bool {
			if (_455_v3).Contains(p3) {
				return (_455_v3).Get(p3).(bool)
			}
			return p1
		})())
		var _457_v5 _dafny.Map
		_ = _457_v5
		_457_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v4, p3)
		var _458_v6 _dafny.MultiSet
		_ = _458_v6
		_458_v6 = _dafny.MultiSetOf((_dafny.Zero).Minus(p3))
		var _459_v7 _dafny.Array
		_ = _459_v7
		var _nwElement0_10 bool = p0
		_ = _nwElement0_10
		var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(17))
		_ = _nw79
		(_nw79).ArraySet1(_nwElement0_10, 0)
		(_nw79).ArraySet1(p1, 1)
		(_nw79).ArraySet1(p1, 2)
		(_nw79).ArraySet1(p0, 3)
		(_nw79).ArraySet1(p0, 4)
		(_nw79).ArraySet1((_dafny.SetOf(_452_v0)).IsProperSubsetOf(_dafny.SetOf(_452_v0, _452_v0)), 5)
		(_nw79).ArraySet1(true, 6)
		(_nw79).ArraySet1(p0, 7)
		(_nw79).ArraySet1(p1, 8)
		(_nw79).ArraySet1((p0) && ((_454_v2).Dtor_cf6()), 9)
		(_nw79).ArraySet1((_457_v5).Contains(_456_v4), 10)
		(_nw79).ArraySet1(false, 11)
		(_nw79).ArraySet1(((_458_v6).Update(p3, Companion_Default___.Abs(p3))).IsDisjointFrom(_dafny.MultiSetOf(p3)), 12)
		(_nw79).ArraySet1(p1, 13)
		(_nw79).ArraySet1(p1, 14)
		(_nw79).ArraySet1(p0, 15)
		(_nw79).ArraySet1(p0, 16)
		_459_v7 = _nw79
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))
		_ = _index55
		(_459_v7).ArraySet1(p0, (_index55).Int())
		var _460_v8 D3
		_ = _460_v8
		_460_v8 = Companion_D3_.Create_DC9_(p3)
		var _pat_let_tv11 = p1
		_ = _pat_let_tv11
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _pat_let_tv13 = p1
		_ = _pat_let_tv13
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))
		_ = _index56
		(_459_v7).ArraySet1(func(_source7 D3) bool {
			if _source7.Is_DC9() {
				var _461___mcc_h0 _dafny.Int = _source7.Get_().(D3_DC9).Cf11
				_ = _461___mcc_h0
				var _462_cf11 _dafny.Int = _461___mcc_h0
				_ = _462_cf11
				return _pat_let_tv11
			} else if _source7.Is_DC8() {
				var _463___mcc_h1 _dafny.Sequence = _source7.Get_().(D3_DC8).Cf10
				_ = _463___mcc_h1
				var _464_cf10 _dafny.Sequence = _463___mcc_h1
				_ = _464_cf10
				return _pat_let_tv12
			} else {
				var _465___mcc_h2 D3 = _source7.Get_().(D3_DC10).Cf12
				_ = _465___mcc_h2
				var _466_cf12 D3 = _465___mcc_h2
				_ = _466_cf12
				return _pat_let_tv13
			}
		}(_460_v8), (_index56).Int())
		var _467_v9 D4
		_ = _467_v9
		_467_v9 = Companion_D4_.Create_DC11_(_459_v7)
		var _source8 D4 = _467_v9
		_ = _source8
		if _source8.Is_DC12() {
			var _468___mcc_h3 _dafny.Int = _source8.Get_().(D4_DC12).Cf14
			_ = _468___mcc_h3
			var _469___mcc_h4 _dafny.Int = _source8.Get_().(D4_DC12).Cf15
			_ = _469___mcc_h4
			var _470___mcc_h5 _dafny.Sequence = _source8.Get_().(D4_DC12).Cf16
			_ = _470___mcc_h5
			var _471___mcc_h6 _dafny.Int = _source8.Get_().(D4_DC12).Cf17
			_ = _471___mcc_h6
			var _472___mcc_h7 _dafny.Int = _source8.Get_().(D4_DC12).Cf18
			_ = _472___mcc_h7
			var _473_cf18 _dafny.Int = _472___mcc_h7
			_ = _473_cf18
			var _474_cf17 _dafny.Int = _471___mcc_h6
			_ = _474_cf17
			var _475_cf16 _dafny.Sequence = _470___mcc_h5
			_ = _475_cf16
			var _476_cf15 _dafny.Int = _469___mcc_h4
			_ = _476_cf15
			var _477_cf14 _dafny.Int = _468___mcc_h3
			_ = _477_cf14
			_475_cf16 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_475_cf16, _475_cf16), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("taufebrr"), _475_cf16))
			r3 = _477_cf14
			var _478_v10 _dafny.Set
			_ = _478_v10
			_478_v10 = _dafny.SetOf(_452_v0)
			var _479_v13 _dafny.Map
			_ = _479_v13
			_479_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_452_v0, _477_cf14)
			var _480_v15 _dafny.MultiSet
			_ = _480_v15
			_480_v15 = _dafny.MultiSetOf(_452_v0, _452_v0)
			var _481_v17 _dafny.Map
			_ = _481_v17
			_481_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_452_v0, p1)
			var _482_v19 _dafny.Array
			_ = _482_v19
			var _nwElement0_11 _dafny.Set = (func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate((func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter26 := _dafny.Iterate((_478_v10).Elements()); ; {
						_compr_26, _ok26 := _iter26()
						if !_ok26 {
							break
						}
						var _483_v11 _dafny.CodePoint
						_483_v11 = interface{}(_compr_26).(_dafny.CodePoint)
						if (_478_v10).Contains(_483_v11) {
							_coll26.Add(_483_v11)
						}
					}
					return _coll26.ToSet()
				}()).Elements()); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _484_v12 _dafny.CodePoint
					_484_v12 = interface{}(_compr_25).(_dafny.CodePoint)
					if (func() _dafny.Set {
						var _coll27 = _dafny.NewBuilder()
						_ = _coll27
						for _iter27 := _dafny.Iterate((_478_v10).Elements()); ; {
							_compr_27, _ok27 := _iter27()
							if !_ok27 {
								break
							}
							var _485_v11 _dafny.CodePoint
							_485_v11 = interface{}(_compr_27).(_dafny.CodePoint)
							if (_478_v10).Contains(_485_v11) {
								_coll27.Add(_485_v11)
							}
						}
						return _coll27.ToSet()
					}()).Contains(_484_v12) {
						_coll25.Add(_484_v12)
					}
				}
				return _coll25.ToSet()
			}()).Union(_478_v10)
			_ = _nwElement0_11
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(22))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_11, 0)
			(_nw80).ArraySet1(_478_v10, 1)
			(_nw80).ArraySet1((_478_v10).Intersection(Companion_Default___.Fm28((_458_v6).Cardinality(), _474_cf17, !(p1), globalState)), 2)
			(_nw80).ArraySet1(_dafny.SetOf(_452_v0, Companion_Default___.Fm29(globalState), _452_v0, _452_v0), 3)
			(_nw80).ArraySet1(_478_v10, 4)
			(_nw80).ArraySet1((_dafny.SetOf(_452_v0, _452_v0)).Intersection(_478_v10), 5)
			(_nw80).ArraySet1((func() _dafny.Set {
				if true {
					return func() _dafny.Set {
						var _coll28 = _dafny.NewBuilder()
						_ = _coll28
						for _iter28 := _dafny.Iterate((_479_v13).Keys().Elements()); ; {
							_compr_28, _ok28 := _iter28()
							if !_ok28 {
								break
							}
							var _486_v14 _dafny.CodePoint
							_486_v14 = interface{}(_compr_28).(_dafny.CodePoint)
							if (_479_v13).Contains(_486_v14) {
								_coll28.Add(_486_v14)
							}
						}
						return _coll28.ToSet()
					}()
				}
				return _dafny.SetOf(_452_v0)
			})(), 6)
			(_nw80).ArraySet1(_478_v10, 7)
			(_nw80).ArraySet1(_478_v10, 8)
			(_nw80).ArraySet1((_478_v10).Intersection(_478_v10), 9)
			(_nw80).ArraySet1(_478_v10, 10)
			(_nw80).ArraySet1((func() _dafny.Set {
				var _coll29 = _dafny.NewBuilder()
				_ = _coll29
				for _iter29 := _dafny.Iterate((_480_v15).Elements()); ; {
					_compr_29, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _487_v16 _dafny.CodePoint
					_487_v16 = interface{}(_compr_29).(_dafny.CodePoint)
					if (_480_v15).Contains(_487_v16) {
						_coll29.Add(_487_v16)
					}
				}
				return _coll29.ToSet()
			}()).Difference(_478_v10), 11)
			(_nw80).ArraySet1(_478_v10, 12)
			(_nw80).ArraySet1(_478_v10, 13)
			(_nw80).ArraySet1(_478_v10, 14)
			(_nw80).ArraySet1(_478_v10, 15)
			(_nw80).ArraySet1(_478_v10, 16)
			(_nw80).ArraySet1(_478_v10, 17)
			(_nw80).ArraySet1((_478_v10).Intersection(func() _dafny.Set {
				var _coll30 = _dafny.NewBuilder()
				_ = _coll30
				for _iter30 := _dafny.Iterate((_481_v17).Keys().Elements()); ; {
					_compr_30, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _488_v18 _dafny.CodePoint
					_488_v18 = interface{}(_compr_30).(_dafny.CodePoint)
					if (_481_v17).Contains(_488_v18) {
						_coll30.Add(_488_v18)
					}
				}
				return _coll30.ToSet()
			}()), 18)
			(_nw80).ArraySet1(_478_v10, 19)
			(_nw80).ArraySet1(_dafny.SetOf(_452_v0), 20)
			(_nw80).ArraySet1(_dafny.SetOf(_452_v0), 21)
			_482_v19 = _nw80
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_482_v19), 0))
			_ = _index57
			(_482_v19).ArraySet1(Companion_Default___.Fm28(_474_cf17, _dafny.IntOfUint32((_475_cf16).Cardinality()), (_459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))).Int()).(bool), globalState), (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_482_v19), 0))
			_ = _index58
			(_482_v19).ArraySet1(_478_v10, (_index58).Int())
			r3 = _dafny.IntOfInt64(340)
		} else if _source8.Is_DC13() {
			var _489___mcc_h8 _dafny.Int = _source8.Get_().(D4_DC13).Cf19
			_ = _489___mcc_h8
			var _490___mcc_h9 _dafny.Array = _source8.Get_().(D4_DC13).Cf20
			_ = _490___mcc_h9
			var _491_cf20 _dafny.Array = _490___mcc_h9
			_ = _491_cf20
			var _492_cf19 _dafny.Int = _489___mcc_h8
			_ = _492_cf19
			var _493_v20 *C0
			_ = _493_v20
			var _nw81 *C0 = New_C0_()
			_ = _nw81
			_nw81.Ctor__((p2).IsSubsetOf(p2), _459_v7)
			_493_v20 = _nw81
			var _494_v21 _dafny.Array
			_ = _494_v21
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw82
			_494_v21 = _nw82
			_494_v21 = _494_v21
			_492_cf19 = _492_cf19
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_494_v21), 0))
			_ = _index59
			(_494_v21).ArraySet1(p3, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_494_v21), 0))
			_ = _index60
			(_494_v21).ArraySet1(_492_cf19, (_index60).Int())
		} else if _source8.Is_DC14() {
			var _495_v22 _dafny.Sequence
			_ = _495_v22
			_495_v22 = _dafny.UnicodeSeqOfUtf8Bytes("guejok")
			var _496_v23 _dafny.Array
			_ = _496_v23
			var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw83
			_496_v23 = _nw83
			var _497_v24 _dafny.Map
			_ = _497_v24
			_497_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _496_v23)
			var _498_v25 D4
			_ = _498_v25
			_498_v25 = Companion_D4_.Create_DC12_(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("isuftd")).Cardinality()), p3, p3, true, globalState), p3, _495_v22, (_dafny.Zero).Minus(((_497_v24).Update(true, _496_v23)).Cardinality()), _dafny.IntOfInt64(612))
			r3 = _dafny.IntOfUint32(((_498_v25).Dtor_cf16()).Cardinality())
			var _499_v26 _dafny.Array
			_ = _499_v26
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_14
			var _nw84 _dafny.Array
			_ = _nw84
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw84 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Sequence = (func(_500_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_501_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-92))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg26 _dafny.Int) interface{} {
								return coer26(arg26)
							}
						}((func(_502_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_503_i1 _dafny.Int) _dafny.CodePoint {
								return _502_v0
							}
						})(_500_v0))), _dafny.UnicodeSeqOfUtf8Bytes("elvverckl"))
					}
				})(_452_v0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw84 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw84).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw84).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_499_v26 = _nw84
			var _504_v27 _dafny.Sequence
			_ = _504_v27
			_504_v27 = _dafny.SeqOf(_495_v22, _495_v22, _495_v22)
			var _505_v28 _dafny.Sequence
			_ = _505_v28
			_505_v28 = _dafny.SeqOf(p0, p0)
			var _506_v29 D8
			_ = _506_v29
			_506_v29 = Companion_D8_.Create_DC21_(p0, p1, _505_v28, true, p3)
			var _507_v30 _dafny.Set
			_ = _507_v30
			_507_v30 = _dafny.SetOf((_505_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_506_v29).Dtor_cf28()).Cardinality()), _dafny.IntOfUint32((_505_v28).Cardinality()))).Uint32()).(bool))
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_499_v26), 0))
			_ = _index61
			(_499_v26).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(638))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_508_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), (_504_v27).Select((Companion_Default___.SafeIndex((_507_v30).Cardinality(), _dafny.IntOfUint32((_504_v27).Cardinality()))).Uint32()).(_dafny.Sequence)), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_499_v26), 0))
			_ = _index62
			(_499_v26).ArraySet1(_495_v22, (_index62).Int())
			var _509_v31 D0
			_ = _509_v31
			_509_v31 = Companion_D0_.Create_DC0_(_455_v3)
			var _510_v32 _dafny.Sequence
			_ = _510_v32
			_510_v32 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, p0)))
			var _511_v33 *C0
			_ = _511_v33
			var _nw85 *C0 = New_C0_()
			_ = _nw85
			_nw85.Ctor__((_this).Fm4(_509_v31, Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(728), (_459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))).Int()).(bool))), (_510_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.IntOfUint32((_510_v32).Cardinality()))).Uint32()).(_dafny.MultiSet), globalState), _459_v7)
			_511_v33 = _nw85
			var _512_v34 *C0
			_ = _512_v34
			var _nw86 *C0 = New_C0_()
			_ = _nw86
			_nw86.Ctor__((_454_v2).Dtor_cf6(), (_511_v33).F11())
			_512_v34 = _nw86
		} else {
			var _513___mcc_h10 _dafny.Array = _source8.Get_().(D4_DC11).Cf13
			_ = _513___mcc_h10
			var _514_cf13 _dafny.Array = _513___mcc_h10
			_ = _514_cf13
			(globalState).F5 = p3
			var _515_v35 *C0
			_ = _515_v35
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__(p0, _514_cf13)
			_515_v35 = _nw87
			var _516_v36 _dafny.Array
			_ = _516_v36
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_15
			var _nw88 _dafny.Array
			_ = _nw88
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw88 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = (func(_517_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_518_i3 _dafny.Int) _dafny.Int {
						return (_518_i3).Minus(_517_p3)
					}
				})(p3)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw88 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw88).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw88).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_516_v36 = _nw88
			var _519_v37 _dafny.Sequence
			_ = _519_v37
			_519_v37 = _dafny.SeqOf(_516_v36)
			var _520_v38 D10
			_ = _520_v38
			_520_v38 = Companion_D10_.Create_DC26_(_519_v37)
			_519_v37 = _dafny.Companion_Sequence_.Concatenate((_520_v38).Dtor_cf37(), _519_v37)
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__((p3).Cmp(p3) >= 0, (_515_v35).F11())
			_515_v35 = _nw89
		}
		r3 = Companion_Default___.SafeDivisionInt(p3, p3)
		var _521_v39 _dafny.Array
		_ = _521_v39
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw90
		_521_v39 = _nw90
		var _522_v40 _dafny.Array
		_ = _522_v40
		_522_v40 = _521_v39
		var _source9 _dafny.Array = _522_v40
		_ = _source9
		var _523___mcc_h11 _dafny.Array = _source9
		_ = _523___mcc_h11
		var _524_cf21 _dafny.Array = _523___mcc_h11
		_ = _524_cf21
		var _525_v41 _dafny.MultiSet
		_ = _525_v41
		_525_v41 = _dafny.MultiSetOf(_452_v0)
		_525_v41 = ((_525_v41).Union(_525_v41)).Union((_525_v41).Update(_452_v0, Companion_Default___.Abs(p3)))
		(globalState).F5 = p3
		var _526_v43 _dafny.Sequence
		_ = _526_v43
		_526_v43 = _dafny.SeqOf(p3)
		var _527_v44 _dafny.Sequence
		_ = _527_v44
		_527_v44 = _dafny.SeqOf((_455_v3).Merge(func() _dafny.Map {
			var _coll31 = _dafny.NewMapBuilder()
			_ = _coll31
			for _iter31 := _dafny.Iterate((_526_v43).Elements()); ; {
				_compr_31, _ok31 := _iter31()
				if !_ok31 {
					break
				}
				var _528_v42 _dafny.Int
				_528_v42 = interface{}(_compr_31).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_526_v43, _528_v42) {
					_coll31.Add(Companion_Default___.SafeDivisionInt(_528_v42, _dafny.IntOfInt64(-970)), p1)
				}
			}
			return _coll31.ToMap()
		}()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))).Int()).(bool)), _455_v3)
		_527_v44 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_527_v44, _527_v44), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(p3, p3, p3, (_459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_527_v44, _527_v44)).Cardinality()))).Uint32(), Companion_Default___.Fm22(p0, globalState))
		var _529_v45 _dafny.Array
		_ = _529_v45
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_16
		var _nw91 _dafny.Array
		_ = _nw91
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw91 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Sequence = func(_530_i4 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ipgn"), _dafny.UnicodeSeqOfUtf8Bytes("uhyhwmg"))
			}
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
		_529_v45 = _nw91
		var _531_v46 _dafny.Sequence
		_ = _531_v46
		_531_v46 = _dafny.UnicodeSeqOfUtf8Bytes("knmngv")
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_529_v45), 0))
		_ = _index63
		(_529_v45).ArraySet1(_531_v46, (_index63).Int())
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_529_v45), 0))
		_ = _index64
		(_529_v45).ArraySet1(Companion_Default___.Fm27(globalState), (_index64).Int())
		var _532_v47 _dafny.Array
		_ = _532_v47
		var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
		_ = _nw92
		_532_v47 = _nw92
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_532_v47), 0))
		_ = _index65
		(_532_v47).ArraySet1CodePoint(_452_v0, (_index65).Int())
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_532_v47), 0))
		_ = _index66
		(_532_v47).ArraySet1CodePoint(_452_v0, (_index66).Int())
		var _533_v48 _dafny.Map
		_ = _533_v48
		_533_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(-115))
		var _534_v49 D0
		_ = _534_v49
		_534_v49 = Companion_D0_.Create_DC0_(_455_v3)
		var _rhs80 _dafny.Map = (func() _dafny.Map {
			if p0 {
				return _533_v48
			}
			return _533_v48
		})()
		_ = _rhs80
		var _rhs81 bool = (func() bool {
			if (_456_v4).Contains(!((p1) == (p1))) {
				return (_456_v4).Get(!((p1) == (p1))).(bool)
			}
			return (_this).Fm4(_534_v49, _534_v49, p2, globalState)
		})()
		_ = _rhs81
		var _rhs82 _dafny.Array = _459_v7
		_ = _rhs82
		_533_v48 = _rhs80
		r1 = _rhs81
		_459_v7 = _rhs82
		var _535_v50 _dafny.Map
		_ = _535_v50
		_535_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))).Int()).(bool), _dafny.IntOfInt64(-467))
		r0 = (_535_v50).Update(p0, p3)
		r1 = (_459_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_459_v7), 0))).Int()).(bool)
		var _536_v51 _dafny.Array
		_ = _536_v51
		var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
		_ = _nw93
		_536_v51 = _nw93
		var _537_v52 _dafny.Sequence
		_ = _537_v52
		_537_v52 = _dafny.SeqOf(_536_v51, _536_v51, _536_v51)
		r2 = (_537_v52).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_537_v52).Cardinality()))).Uint32()).(_dafny.Array)
		r3 = p3
		return r0, r1, r2, r3
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

func (_this *C2) Ctor__(f13 bool) {
	{
		(_this)._f13 = f13
	}
}
func (_this *C2) Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return (_this).F13()
	}
}
func (_this *C2) Fm18(p0 D2, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(65)
	}
}
func (_this *C2) M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _538_v0 _dafny.Sequence
		_ = _538_v0
		_538_v0 = _dafny.SeqOf((_this).F13(), (_this).F13())
		r2 = !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_538_v0, _538_v0), _dafny.Companion_Sequence_.Concatenate(_538_v0, _538_v0)))
		var _539_v1 _dafny.CodePoint
		_ = _539_v1
		_539_v1 = _dafny.CodePoint('j')
		var _540_v2 _dafny.Map
		_ = _540_v2
		_540_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _539_v1)
		var _541_v3 _dafny.Sequence
		_ = _541_v3
		_541_v3 = _dafny.UnicodeSeqOfUtf8Bytes("bnmkw")
		var _542_v4 D3
		_ = _542_v4
		_542_v4 = Companion_D3_.Create_DC8_(_541_v3)
		var _543_v5 D3
		_ = _543_v5
		_543_v5 = Companion_D3_.Create_DC10_(_542_v4)
		var _544_v7 _dafny.Map
		_ = _544_v7
		_544_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_v5, func() _dafny.Map {
			var _coll32 = _dafny.NewMapBuilder()
			_ = _coll32
			for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(225), _dafny.IntOfInt64(316))); ; {
				_compr_32, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _545_v6 _dafny.Int
				_545_v6 = interface{}(_compr_32).(_dafny.Int)
				if ((_dafny.IntOfInt64(225)).Cmp(_545_v6) <= 0) && ((_545_v6).Cmp(_dafny.IntOfInt64(316)) < 0) {
					_coll32.Add((_545_v6).Times(p2), _539_v1)
				}
			}
			return _coll32.ToMap()
		}())
		var _546_v9 _dafny.Set
		_ = _546_v9
		_546_v9 = _dafny.SetOf(p0, p2)
		var _547_v10 _dafny.Map
		_ = _547_v10
		_547_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter33 := _dafny.Iterate((_546_v9).Elements()); ; {
				_compr_33, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _548_v8 _dafny.Int
				_548_v8 = interface{}(_compr_33).(_dafny.Int)
				if (_546_v9).Contains(_548_v8) {
					_coll33.Add(Companion_Default___.SafeDivisionInt(_548_v8, p2), _539_v1)
				}
			}
			return _coll33.ToMap()
		}())
		var _549_v11 _dafny.Set
		_ = _549_v11
		_549_v11 = _dafny.SetOf((_this).F13(), (_this).F13())
		var _550_v12 _dafny.Map
		_ = _550_v12
		_550_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_549_v11).Cardinality())
		var _551_v14 _dafny.Sequence
		_ = _551_v14
		_551_v14 = _dafny.SeqOf(_540_v2)
		var _552_v15 _dafny.Sequence
		_ = _552_v15
		_552_v15 = _dafny.SeqOf(p0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_541_v3).Cardinality())))
		var _553_v17 _dafny.Array
		_ = _553_v17
		var _nwElement0_12 _dafny.Map = (_540_v2).Update(p2, _539_v1)
		_ = _nwElement0_12
		var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(28))
		_ = _nw94
		(_nw94).ArraySet1(_nwElement0_12, 0)
		(_nw94).ArraySet1((_540_v2).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(952), _539_v1)), 1)
		(_nw94).ArraySet1(((_540_v2).Update(p0, _539_v1)).Merge(_540_v2), 2)
		(_nw94).ArraySet1((_540_v2).Merge(_540_v2), 3)
		(_nw94).ArraySet1(((func() _dafny.Map {
			if (_544_v7).Contains(_543_v5) {
				return (_544_v7).Get(_543_v5).(_dafny.Map)
			}
			return _540_v2
		})()).Merge(_540_v2), 4)
		(_nw94).ArraySet1(Companion_Default___.Fm19(globalState), 5)
		(_nw94).ArraySet1(_540_v2, 6)
		(_nw94).ArraySet1((_540_v2).Merge(_540_v2), 7)
		(_nw94).ArraySet1(_540_v2, 8)
		(_nw94).ArraySet1((_540_v2).Update(p2, _dafny.CodePoint('d')), 9)
		(_nw94).ArraySet1((func() _dafny.Map {
			if (_this).F13() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), _539_v1)
			}
			return _540_v2
		})(), 10)
		(_nw94).ArraySet1(_540_v2, 11)
		(_nw94).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_541_v3).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_541_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)), 12)
		(_nw94).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())).Cardinality(), _539_v1)).Update(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_554_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_555_i1 _dafny.Int) _dafny.CodePoint {
				return _554_v1
			}
		})(_539_v1)))).Cardinality()), _539_v1), 13)
		(_nw94).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _539_v1), 14)
		(_nw94).ArraySet1(_540_v2, 15)
		(_nw94).ArraySet1((_540_v2).Update(p3, _539_v1), 16)
		(_nw94).ArraySet1(_540_v2, 17)
		(_nw94).ArraySet1(_540_v2, 18)
		(_nw94).ArraySet1(_540_v2, 19)
		(_nw94).ArraySet1((func() _dafny.Map {
			if (_547_v10).Contains(!((_this).F13())) {
				return (_547_v10).Get(!((_this).F13())).(_dafny.Map)
			}
			return _540_v2
		})(), 20)
		(_nw94).ArraySet1(_540_v2, 21)
		(_nw94).ArraySet1(_540_v2, 22)
		(_nw94).ArraySet1(_540_v2, 23)
		(_nw94).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm20((_dafny.Zero).Minus((_550_v12).Cardinality()), (_this).F13(), p3, globalState)), 24)
		(_nw94).ArraySet1(func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter34 := _dafny.Iterate((_546_v9).Elements()); ; {
				_compr_34, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _556_v13 _dafny.Int
				_556_v13 = interface{}(_compr_34).(_dafny.Int)
				if (_546_v9).Contains(_556_v13) {
					_coll34.Add(Companion_Default___.SafeDivisionInt(_556_v13, p3), _539_v1)
				}
			}
			return _coll34.ToMap()
		}(), 25)
		(_nw94).ArraySet1((_551_v14).Select((Companion_Default___.SafeIndex((_552_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_541_v3).Cardinality()), _dafny.IntOfUint32((_552_v15).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_551_v14).Cardinality()))).Uint32()).(_dafny.Map), 26)
		(_nw94).ArraySet1((func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(921), _dafny.IntOfInt64(-898))); ; {
				_compr_35, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _557_v16 _dafny.Int
				_557_v16 = interface{}(_compr_35).(_dafny.Int)
				if ((_dafny.IntOfInt64(921)).Cmp(_557_v16) <= 0) && ((_557_v16).Cmp(_dafny.IntOfInt64(-898)) < 0) {
					_coll35.Add((_557_v16).Times(p2), _dafny.CodePoint('s'))
				}
			}
			return _coll35.ToMap()
		}()).Update(p2, _539_v1), 27)
		_553_v17 = _nw94
		for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_553_v17), 0))); ; {
			_guard_loop_0, _ok36 := _iter36()
			if !_ok36 {
				break
			}
			var _558_i0 _dafny.Int
			_558_i0 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_558_i0).Sign() != -1) && ((_558_i0).Cmp(_dafny.ArrayLen((_553_v17), 0)) < 0)) {
				(_553_v17).ArraySet1((_540_v2).Merge((_540_v2).Update(_dafny.IntOfInt64(74), _539_v1)), (_558_i0).Int())
			}
		}
		var _559_i2 _dafny.Int
		_ = _559_i2
		_559_i2 = _dafny.Zero
		{
			for (_this).F13() {
				{
					if (_559_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_559_i2 = (_559_i2).Plus(_dafny.One)
					(globalState).F5 = (_dafny.Zero).Minus((_dafny.IntOfInt64(65)).Times(p0))
					r1 = (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F13())).Cardinality(), p2)).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mow")).Cardinality())).Plus((_dafny.Zero).Minus(p3)))
					if (_this).F13() {
						var _560_v18 _dafny.Array
						_ = _560_v18
						var _len0_17 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_17
						var _nw95 _dafny.Array
						_ = _nw95
						if _len0_17.Cmp(_dafny.Zero) == 0 {
							_nw95 = _dafny.NewArray(_len0_17)
						} else {
							var _init17 func(_dafny.Int) _dafny.Int = (func(_561_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_562_i3 _dafny.Int) _dafny.Int {
									return (_562_i3).Plus(_561_p2)
								}
							})(p2)
							_ = _init17
							var _element0_17 = _init17(_dafny.Zero)
							_ = _element0_17
							_nw95 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
							(_nw95).ArraySet1(_element0_17, 0)
							var _nativeLen0_17 = (_len0_17).Int()
							_ = _nativeLen0_17
							for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
								(_nw95).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
							}
						}
						_560_v18 = _nw95
						var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v18), 0))
						_ = _index67
						(_560_v18).ArraySet1((p0).Plus(p3), (_index67).Int())
						var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v18), 0))
						_ = _index68
						(_560_v18).ArraySet1(p0, (_index68).Int())
						var _563_v19 _dafny.Array
						_ = _563_v19
						var _nwElement0_13 bool = (_this).F13()
						_ = _nwElement0_13
						var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
						_ = _nw96
						(_nw96).ArraySet1(_nwElement0_13, 0)
						(_nw96).ArraySet1(false, 1)
						(_nw96).ArraySet1((func() bool {
							if (_this).F13() {
								return (_538_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_541_v3).Cardinality()), _dafny.IntOfUint32((_538_v0).Cardinality()))).Uint32()).(bool)
							}
							return (_this).F13()
						})(), 2)
						(_nw96).ArraySet1((_this).F13(), 3)
						(_nw96).ArraySet1(false, 4)
						(_nw96).ArraySet1((_this).F13(), 5)
						(_nw96).ArraySet1((_this).F13(), 6)
						_563_v19 = _nw96
						_563_v19 = _563_v19
						var _564_v20 T1
						_ = _564_v20
						var _nw97 *C0 = New_C0_()
						_ = _nw97
						_nw97.Ctor__(true, _563_v19)
						_564_v20 = _nw97
						var _565_v21 _dafny.Set
						_ = _565_v21
						_565_v21 = _dafny.SetOf(Companion_D2_.Create_DC5_(_564_v20))
						var _566_v22 D2
						_ = _566_v22
						_566_v22 = Companion_D2_.Create_DC5_(_564_v20)
						var _pat_let_tv14 = _564_v20
						_ = _pat_let_tv14
						_565_v21 = (func() _dafny.Set {
							if (_this).F13() {
								return _dafny.SetOf(func(_pat_let6_0 D2) D2 {
									return func(_567_dt__update__tmp_h0 D2) D2 {
										return func(_pat_let7_0 T1) D2 {
											return func(_568_dt__update_hcf5_h0 T1) D2 {
												return Companion_D2_.Create_DC5_(_568_dt__update_hcf5_h0)
											}(_pat_let7_0)
										}(_pat_let_tv14)
									}(_pat_let6_0)
								}(Companion_D2_.Create_DC5_(_564_v20)), _566_v22)
							}
							return _dafny.SetOf(Companion_D2_.Create_DC5_(_564_v20), _566_v22, _566_v22, _566_v22, _566_v22)
						})()
						_550_v12 = (_550_v12).Update(true, Companion_Default___.SafeDivisionInt((_560_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v18), 0))).Int()).(_dafny.Int), p2))
						var _569_v23 D2
						_ = _569_v23
						_569_v23 = Companion_D2_.Create_DC6_((_this).F13())
						var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v18), 0))
						_ = _index69
						var _rhs83 _dafny.Int = (p3).Times((_this).Fm18(_569_v23, (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm21(globalState)).Cardinality())), p3, globalState))
						_ = _rhs83
						var _rhs84 bool = (_this).F13()
						_ = _rhs84
						var _rhs85 _dafny.Array = _560_v18
						_ = _rhs85
						var _lhs58 _dafny.Array = _560_v18
						_ = _lhs58
						var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v18), 0))
						_ = _lhs59
						var _lhs60 *GlobalState = globalState
						_ = _lhs60
						(_lhs58).ArraySet1(_rhs83, (_lhs59).Int())
						_lhs60.F2 = _rhs84
						_560_v18 = _rhs85
					} else {
						var _570_v24 D2
						_ = _570_v24
						_570_v24 = Companion_D2_.Create_DC6_((_this).F13())
						var _571_v25 _dafny.Array
						_ = _571_v25
						var _nwElement0_14 bool = (_this).F13()
						_ = _nwElement0_14
						var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(13))
						_ = _nw98
						(_nw98).ArraySet1(_nwElement0_14, 0)
						(_nw98).ArraySet1((_this).F13(), 1)
						(_nw98).ArraySet1((Companion_Default___.Fm1((_this).F13(), globalState)) || ((_this).F13()), 2)
						(_nw98).ArraySet1((func() bool {
							if (_this).F13() {
								return (_this).F13()
							}
							return (_this).F13()
						})(), 3)
						(_nw98).ArraySet1(false, 4)
						(_nw98).ArraySet1((_this).F13(), 5)
						(_nw98).ArraySet1((_this).F13(), 6)
						(_nw98).ArraySet1((_this).F13(), 7)
						(_nw98).ArraySet1((_this).F13(), 8)
						(_nw98).ArraySet1((_this).F13(), 9)
						(_nw98).ArraySet1((_570_v24).Dtor_cf6(), 10)
						(_nw98).ArraySet1((_this).F13(), 11)
						(_nw98).ArraySet1((_this).F13(), 12)
						_571_v25 = _nw98
						var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_571_v25), 0))
						_ = _index70
						(_571_v25).ArraySet1((_this).F13(), (_index70).Int())
						var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_571_v25), 0))
						_ = _index71
						(_571_v25).ArraySet1(!((p3).Cmp(p3) != 0), (_index71).Int())
						var _572_v26 _dafny.Array
						_ = _572_v26
						var _len0_18 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_18
						var _nw99 _dafny.Array
						_ = _nw99
						if _len0_18.Cmp(_dafny.Zero) == 0 {
							_nw99 = _dafny.NewArray(_len0_18)
						} else {
							var _init18 func(_dafny.Int) _dafny.CodePoint = (func(_573_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_574_i4 _dafny.Int) _dafny.CodePoint {
									return _573_v1
								}
							})(_539_v1)
							_ = _init18
							var _element0_18 = _init18(_dafny.Zero)
							_ = _element0_18
							_nw99 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
							(_nw99).ArraySet1CodePoint(_element0_18, 0)
							var _nativeLen0_18 = (_len0_18).Int()
							_ = _nativeLen0_18
							for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
								(_nw99).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
							}
						}
						_572_v26 = _nw99
						var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_572_v26), 0))
						_ = _index72
						(_572_v26).ArraySet1CodePoint((func() _dafny.CodePoint {
							if (_this).F13() {
								return _539_v1
							}
							return _539_v1
						})(), (_index72).Int())
						var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_571_v25), 0))
						_ = _index73
						var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_572_v26), 0))
						_ = _index74
						var _rhs86 bool = (_571_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_571_v25), 0))).Int()).(bool)
						_ = _rhs86
						var _rhs87 _dafny.CodePoint = _539_v1
						_ = _rhs87
						var _lhs61 _dafny.Array = _571_v25
						_ = _lhs61
						var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_571_v25), 0))
						_ = _lhs62
						var _lhs63 _dafny.Array = _572_v26
						_ = _lhs63
						var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_572_v26), 0))
						_ = _lhs64
						(_lhs61).ArraySet1(_rhs86, (_lhs62).Int())
						(_lhs63).ArraySet1CodePoint(_rhs87, (_lhs64).Int())
						(globalState).F2 = (_this).F13()
						(globalState).F5 = p3
						var _575_v27 D4
						_ = _575_v27
						_575_v27 = Companion_D4_.Create_DC12_(p0, p3, _541_v3, p3, _dafny.IntOfUint32((_538_v0).Cardinality()))
						var _576_v28 _dafny.Map
						_ = _576_v28
						_576_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _575_v27)
						_576_v28 = (_576_v28).Update(p2, _575_v27)
					}
					var _577_v29 T0
					_ = _577_v29
					var _nw100 *C1 = New_C1_()
					_ = _nw100
					_nw100.Ctor__()
					_577_v29 = _nw100
					var _578_v30 _dafny.Array
					_ = _578_v30
					var _nwElement0_15 T0 = _577_v29
					_ = _nwElement0_15
					var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(7))
					_ = _nw101
					(_nw101).ArraySet1(_nwElement0_15, 0)
					(_nw101).ArraySet1(_577_v29, 1)
					(_nw101).ArraySet1(_577_v29, 2)
					(_nw101).ArraySet1(_577_v29, 3)
					(_nw101).ArraySet1(_577_v29, 4)
					(_nw101).ArraySet1(_577_v29, 5)
					(_nw101).ArraySet1(_577_v29, 6)
					_578_v30 = _nw101
					var _579_v31 D11
					_ = _579_v31
					_579_v31 = Companion_D11_.Create_DC29_(_577_v29)
					var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_578_v30), 0))
					_ = _index75
					(_578_v30).ArraySet1((_579_v31).Dtor_cf43(), (_index75).Int())
					var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_578_v30), 0))
					_ = _index76
					(_578_v30).ArraySet1(_577_v29, (_index76).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _580_v32 _dafny.Array
		_ = _580_v32
		var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw102
		_580_v32 = _nw102
		_580_v32 = _580_v32
		var _581_v33 T0
		_ = _581_v33
		var _nw103 *C1 = New_C1_()
		_ = _nw103
		_nw103.Ctor__()
		_581_v33 = _nw103
		_581_v33 = _581_v33
		var _582_v34 _dafny.Map
		_ = _582_v34
		_582_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F13())
		var _583_v36 D2
		_ = _583_v36
		_583_v36 = Companion_D2_.Create_DC6_((_this).F13())
		var _584_v38 _dafny.Array
		_ = _584_v38
		var _nwElement0_16 _dafny.Map = _582_v34
		_ = _nwElement0_16
		var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(17))
		_ = _nw104
		(_nw104).ArraySet1(_nwElement0_16, 0)
		(_nw104).ArraySet1((func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-111), _dafny.IntOfInt64(985))); ; {
				_compr_36, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _585_v35 _dafny.Int
				_585_v35 = interface{}(_compr_36).(_dafny.Int)
				if ((_dafny.IntOfInt64(-111)).Cmp(_585_v35) <= 0) && ((_585_v35).Cmp(_dafny.IntOfInt64(985)) < 0) {
					_coll36.Add((_585_v35).Minus(p0), false)
				}
			}
			return _coll36.ToMap()
		}()).Merge(_582_v34), 1)
		(_nw104).ArraySet1((_582_v34).Merge(_582_v34), 2)
		(_nw104).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm18(_583_v36, _dafny.IntOfInt64(510), p2, globalState), (_this).F13()), 3)
		(_nw104).ArraySet1((_582_v34).Update((_dafny.Zero).Minus(p0), (_this).F13()), 4)
		(_nw104).ArraySet1((_582_v34).Merge(func() _dafny.Map {
			var _coll37 = _dafny.NewMapBuilder()
			_ = _coll37
			for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(278), _dafny.IntOfInt64(751))); ; {
				_compr_37, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _586_v37 _dafny.Int
				_586_v37 = interface{}(_compr_37).(_dafny.Int)
				if ((_dafny.IntOfInt64(278)).Cmp(_586_v37) <= 0) && ((_586_v37).Cmp(_dafny.IntOfInt64(751)) < 0) {
					_coll37.Add((_586_v37).Times(p0), (_this).F13())
				}
			}
			return _coll37.ToMap()
		}()), 5)
		(_nw104).ArraySet1(_582_v34, 6)
		(_nw104).ArraySet1(_582_v34, 7)
		(_nw104).ArraySet1(_582_v34, 8)
		(_nw104).ArraySet1((_582_v34).Update(p0, (_this).F13()), 9)
		(_nw104).ArraySet1(_582_v34, 10)
		(_nw104).ArraySet1(_582_v34, 11)
		(_nw104).ArraySet1(_582_v34, 12)
		(_nw104).ArraySet1(_582_v34, 13)
		(_nw104).ArraySet1(_582_v34, 14)
		(_nw104).ArraySet1((_582_v34).Merge(_582_v34), 15)
		(_nw104).ArraySet1((_582_v34).Merge(_582_v34), 16)
		_584_v38 = _nw104
		var _587_v39 _dafny.Set
		_ = _587_v39
		_587_v39 = _dafny.SetOf((_541_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.IntOfUint32((_541_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.CodePoint('j'))
		var _588_v40 _dafny.Map
		_ = _588_v40
		_588_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_587_v39).Difference(_587_v39), _584_v38)
		var _rhs88 bool = (_this).F13()
		_ = _rhs88
		var _rhs89 _dafny.Sequence = _541_v3
		_ = _rhs89
		var _rhs90 _dafny.Int = p3
		_ = _rhs90
		var _rhs91 _dafny.Array = (func() _dafny.Array {
			if (_588_v40).Contains(_587_v39) {
				return (_588_v40).Get(_587_v39).(_dafny.Array)
			}
			return _584_v38
		})()
		_ = _rhs91
		var _lhs65 *GlobalState = globalState
		_ = _lhs65
		r3 = _rhs88
		_541_v3 = _rhs89
		_lhs65.F5 = _rhs90
		_584_v38 = _rhs91
		var _589_v41 _dafny.Map
		_ = _589_v41
		_589_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _538_v0)
		r0 = (_dafny.IntOfInt64(757)).Cmp((_589_v41).Cardinality()) < 0
		r1 = Companion_Default___.SafeDivisionInt(p0, p0)
		r2 = !(!(false))
		var _590_v42 D0
		_ = _590_v42
		_590_v42 = Companion_D0_.Create_DC0_(_582_v34)
		var _591_v43 _dafny.MultiSet
		_ = _591_v43
		_591_v43 = _dafny.MultiSetOf((_this).F13())
		var _592_v44 D8
		_ = _592_v44
		_592_v44 = Companion_D8_.Create_DC21_((_this).F13(), (_this).F13(), _538_v0, (_581_v33).Fm4(_590_v42, _590_v42, _591_v43, globalState), p0)
		r3 = !((_592_v44).Dtor_cf29())
		return r0, r1, r2, r3
	}
}
func (_this *C2) M4(globalState *GlobalState) (bool, _dafny.MultiSet, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _593_v0 _dafny.Int
		_ = _593_v0
		_593_v0 = _dafny.IntOfInt64(945)
		var _hi2 _dafny.Int = _593_v0
		_ = _hi2
		for _594_i0 := _593_v0; _594_i0.Cmp(_hi2) < 0; _594_i0 = _594_i0.Plus(_dafny.One) {
			_593_v0 = _594_i0
			var _595_v1 *C1
			_ = _595_v1
			var _nw105 *C1 = New_C1_()
			_ = _nw105
			_nw105.Ctor__()
			_595_v1 = _nw105
			var _596_v2 _dafny.Array
			_ = _596_v2
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_19
			var _nw106 _dafny.Array
			_ = _nw106
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw106 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) bool = func(_597_i1 _dafny.Int) bool {
					return (_this).F13()
				}
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw106 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw106).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw106).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_596_v2 = _nw106
			var _598_v3 T1
			_ = _598_v3
			var _nw107 *C0 = New_C0_()
			_ = _nw107
			_nw107.Ctor__((_this).F13(), _596_v2)
			_598_v3 = _nw107
			var _599_v4 D2
			_ = _599_v4
			_599_v4 = Companion_D2_.Create_DC5_(_598_v3)
			var _600_v5 _dafny.Set
			_ = _600_v5
			_600_v5 = _dafny.SetOf(_599_v4, _599_v4)
			if (_600_v5).IsProperSubsetOf((_600_v5).Union(_600_v5)) {
				var _601_v6 _dafny.Map
				_ = _601_v6
				_601_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _594_i0)
				var _602_v7 D0
				_ = _602_v7
				_602_v7 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_601_v6).Cardinality(), !((_this).F13())))
				(globalState).F2 = (func() bool {
					if (_this).F13() {
						return (_this).F13()
					}
					return (_this).Fm4(_602_v7, _602_v7, _dafny.MultiSetOf(false), globalState)
				})()
				var _603_v8 _dafny.Set
				_ = _603_v8
				_603_v8 = _dafny.SetOf(_594_i0, _593_v0, _594_i0, _593_v0)
				var _604_v9 _dafny.Sequence
				_ = _604_v9
				_604_v9 = _dafny.SeqOf(((_dafny.Zero).Minus(Companion_Default___.Fm0((_dafny.MultiSetOf(_594_i0, _593_v0)).Cardinality(), _593_v0, _593_v0, (_this).F13(), globalState))).Cmp((_603_v8).Cardinality()) >= 0, (_this).F13(), true, (_this).F13(), (_this).F13())
				_604_v9 = _dafny.Companion_Sequence_.Concatenate(_604_v9, _604_v9)
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_596_v2), 0))
				_ = _index77
				(_596_v2).ArraySet1((_this).F13(), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_596_v2), 0))
				_ = _index78
				(_596_v2).ArraySet1((_this).F13(), (_index78).Int())
				_603_v8 = func() _dafny.Set {
					var _coll38 = _dafny.NewBuilder()
					_ = _coll38
					for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(676), _dafny.IntOfInt64(277))); ; {
						_compr_38, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _605_v10 _dafny.Int
						_605_v10 = interface{}(_compr_38).(_dafny.Int)
						if ((_dafny.IntOfInt64(676)).Cmp(_605_v10) <= 0) && ((_605_v10).Cmp(_dafny.IntOfInt64(277)) < 0) {
							_coll38.Add((_605_v10).Times((Companion_D2_.Create_DC7_(_593_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_596_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_596_v2), 0))).Int()).(bool)), _594_i0), _593_v0)).Dtor_cf7()))
						}
					}
					return _coll38.ToSet()
				}()
				var _606_v11 _dafny.Array
				_ = _606_v11
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw108
				_606_v11 = _nw108
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_606_v11), 0))
				_ = _index79
				(_606_v11).ArraySet1(_593_v0, (_index79).Int())
				var _607_v12 _dafny.Map
				_ = _607_v12
				_607_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_596_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_596_v2), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(702))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}(func(_608_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})))
				var _609_v13 _dafny.Sequence
				_ = _609_v13
				_609_v13 = _dafny.UnicodeSeqOfUtf8Bytes("d")
				var _610_v14 _dafny.Sequence
				_ = _610_v14
				_610_v14 = _dafny.SeqOf(_609_v13)
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_596_v2), 0))
				_ = _index80
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_606_v11), 0))
				_ = _index81
				var _rhs92 _dafny.Int = (func() _dafny.Int {
					if (_this).F13() {
						return (_593_v0).Times(_593_v0)
					}
					return Companion_Default___.SafeModuloInt(_593_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_607_v12).Contains(false) {
							return (_607_v12).Get(false).(_dafny.Sequence)
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("smmjmydnm")
					})()).Cardinality()))
				})()
				_ = _rhs92
				var _rhs93 bool = !_dafny.Companion_Sequence_.Equal(_610_v14, _610_v14)
				_ = _rhs93
				var _rhs94 _dafny.Int = _594_i0
				_ = _rhs94
				var _lhs66 *GlobalState = globalState
				_ = _lhs66
				var _lhs67 _dafny.Array = _596_v2
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_596_v2), 0))
				_ = _lhs68
				var _lhs69 _dafny.Array = _606_v11
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_606_v11), 0))
				_ = _lhs70
				_lhs66.F5 = _rhs92
				(_lhs67).ArraySet1(_rhs93, (_lhs68).Int())
				(_lhs69).ArraySet1(_rhs94, (_lhs70).Int())
			} else {
				var _611_v15 _dafny.Array
				_ = _611_v15
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_20
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Sequence = func(_612_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("m")
					}
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw109 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw109).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw109).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_611_v15 = _nw109
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_611_v15), 0))
				_ = _index82
				(_611_v15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fyghxaonv"), (_index82).Int())
				var _613_v16 _dafny.Sequence
				_ = _613_v16
				_613_v16 = _dafny.UnicodeSeqOfUtf8Bytes("trl")
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_611_v15), 0))
				_ = _index83
				(_611_v15).ArraySet1(_613_v16, (_index83).Int())
				var _614_v17 _dafny.Sequence
				_ = _614_v17
				_614_v17 = _dafny.SeqOf((_this).F13(), (_this).F13(), (_this).F13(), false, (_this).F13())
				var _615_v18 _dafny.MultiSet
				_ = _615_v18
				_615_v18 = _dafny.MultiSetOf((_this).F13(), true)
				r0 = ((func() _dafny.MultiSet {
					if (_this).F13() {
						return _615_v18
					}
					return _615_v18
				})()).IsProperSubsetOf(_dafny.MultiSetFromSeq(_614_v17))
				var _616_v19 _dafny.CodePoint
				_ = _616_v19
				_616_v19 = _dafny.CodePoint('w')
				var _617_v20 _dafny.Map
				_ = _617_v20
				_617_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _616_v19)
				var _618_v21 _dafny.Set
				_ = _618_v21
				_618_v21 = _dafny.SetOf((_this).F13())
				var _619_v22 _dafny.Map
				_ = _619_v22
				_619_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_617_v20).Cardinality()).Cmp((_dafny.Zero).Minus(_594_i0)) == 0, _618_v21)
				_619_v22 = _619_v22
				var _620_v23 _dafny.Array
				_ = _620_v23
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw110
				_620_v23 = _nw110
				var _621_v24 D10
				_ = _621_v24
				_621_v24 = Companion_D10_.Create_DC26_(_dafny.SeqOf(_620_v23, _620_v23))
				_621_v24 = _621_v24
				_593_v0 = _593_v0
			}
			var _622_v25 _dafny.Array
			_ = _622_v25
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
			_ = _nw111
			_622_v25 = _nw111
			var _623_v26 _dafny.MultiSet
			_ = _623_v26
			_623_v26 = _dafny.MultiSetOf((_this).F13())
			var _624_v27 _dafny.Map
			_ = _624_v27
			_624_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _594_i0)
			var _625_v28 _dafny.CodePoint
			_ = _625_v28
			_625_v28 = _dafny.CodePoint('s')
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_622_v25), 0))
			_ = _index84
			(_622_v25).ArraySet1((_623_v26).Intersection((Companion_D10_.Create_DC27_(true, (func() _dafny.Int {
				if (_624_v27).Contains((_this).F13()) {
					return (_624_v27).Get((_this).F13()).(_dafny.Int)
				}
				return _594_i0
			})(), _625_v28, _623_v26)).Dtor_cf41()), (_index84).Int())
			var _626_v29 _dafny.Set
			_ = _626_v29
			_626_v29 = _dafny.SetOf(!((_this).F13()))
			var _627_v30 D12
			_ = _627_v30
			_627_v30 = Companion_D12_.Create_DC33_(_626_v29)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_622_v25), 0))
			_ = _index85
			(_622_v25).ArraySet1(_dafny.MultiSetOf((_this).F13(), true, true, ((_627_v30).Dtor_cf47()).IsSubsetOf(_626_v29)), (_index85).Int())
		}
		var _628_v31 _dafny.Array
		_ = _628_v31
		var _nw112 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(6))
		_ = _nw112
		_628_v31 = _nw112
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_628_v31), 0))); ; {
			_guard_loop_1, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _629_i4 _dafny.Int
			_629_i4 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_629_i4).Sign() != -1) && ((_629_i4).Cmp(_dafny.ArrayLen((_628_v31), 0)) < 0)) {
				(_628_v31).ArraySet1(Companion_D2_.Create_DC7_((_dafny.Zero).Minus(_593_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13()), _593_v0), (func() _dafny.Int {
					if (_this).F13() {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wlpnjgxrs")).Cardinality())
					}
					return _593_v0
				})()), (_629_i4).Int())
			}
		}
		var _rhs95 _dafny.Int = _593_v0
		_ = _rhs95
		var _rhs96 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(177), _593_v0)
		_ = _rhs96
		var _lhs71 *GlobalState = globalState
		_ = _lhs71
		_593_v0 = _rhs95
		_lhs71.F5 = _rhs96
		var _630_v32 *C1
		_ = _630_v32
		var _nw113 *C1 = New_C1_()
		_ = _nw113
		_nw113.Ctor__()
		_630_v32 = _nw113
		var _hi3 _dafny.Int = _593_v0
		_ = _hi3
		for _631_i5 := Companion_Default___.SafeDivisionInt(_593_v0, _593_v0); _631_i5.Cmp(_hi3) < 0; _631_i5 = _631_i5.Plus(_dafny.One) {
			var _632_v33 _dafny.Map
			_ = _632_v33
			_632_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v0, (_this).F13())
			var _633_v34 D0
			_ = _633_v34
			_633_v34 = Companion_D0_.Create_DC0_(_632_v33)
			var _634_v35 _dafny.MultiSet
			_ = _634_v35
			_634_v35 = _dafny.MultiSetOf((_this).F13(), false)
			var _635_v36 _dafny.Array
			_ = _635_v36
			var _nwElement0_17 bool = (_this).F13()
			_ = _nwElement0_17
			var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(12))
			_ = _nw114
			(_nw114).ArraySet1(_nwElement0_17, 0)
			(_nw114).ArraySet1(true, 1)
			(_nw114).ArraySet1((_this).F13(), 2)
			(_nw114).ArraySet1((_this).F13(), 3)
			(_nw114).ArraySet1((_this).F13(), 4)
			(_nw114).ArraySet1((_this).F13(), 5)
			(_nw114).ArraySet1((_this).F13(), 6)
			(_nw114).ArraySet1((_this).F13(), 7)
			(_nw114).ArraySet1((_this).F13(), 8)
			(_nw114).ArraySet1((_this).F13(), 9)
			(_nw114).ArraySet1((_this).F13(), 10)
			(_nw114).ArraySet1(false, 11)
			_635_v36 = _nw114
			var _636_v37 _dafny.Sequence
			_ = _636_v37
			_636_v37 = _dafny.SeqOf(_635_v36)
			var _637_v38 _dafny.Set
			_ = _637_v38
			_637_v38 = _dafny.SetOf(_632_v33, _632_v33)
			var _638_v39 _dafny.Sequence
			_ = _638_v39
			_638_v39 = _dafny.UnicodeSeqOfUtf8Bytes("vgk")
			var _639_v40 _dafny.Sequence
			_ = _639_v40
			_639_v40 = _dafny.SeqOf((_this).F13(), (_this).F13())
			var _640_v41 D12
			_ = _640_v41
			_640_v41 = Companion_D12_.Create_DC36_((_this).F13(), _dafny.IntOfInt64(75), (_this).F13())
			var _641_v42 *C0
			_ = _641_v42
			var _nw115 *C0 = New_C0_()
			_ = _nw115
			_nw115.Ctor__((_this).F13(), _635_v36)
			_641_v42 = _nw115
			var _642_v43 _dafny.Sequence
			_ = _642_v43
			_642_v43 = _dafny.SeqOf(_641_v42)
			var _643_v44 _dafny.Array
			_ = _643_v44
			var _nwElement0_18 bool = Companion_Default___.Fm1((_this).F13(), globalState)
			_ = _nwElement0_18
			var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(29))
			_ = _nw116
			(_nw116).ArraySet1(_nwElement0_18, 0)
			(_nw116).ArraySet1(false, 1)
			(_nw116).ArraySet1(!((_this).F13()), 2)
			(_nw116).ArraySet1((_this).F13(), 3)
			(_nw116).ArraySet1((_this).F13(), 4)
			(_nw116).ArraySet1((_this).Fm4(_633_v34, Companion_D0_.Create_DC0_(_632_v33), _634_v35, globalState), 5)
			(_nw116).ArraySet1(!_dafny.Companion_Sequence_.Equal(_636_v37, _636_v37), 6)
			(_nw116).ArraySet1((_637_v38).IsProperSubsetOf(_637_v38), 7)
			(_nw116).ArraySet1((_this).F13(), 8)
			(_nw116).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf(_638_v39, _638_v39)), 9)
			(_nw116).ArraySet1((_634_v35).IsProperSubsetOf(_634_v35), 10)
			(_nw116).ArraySet1((_this).F13(), 11)
			(_nw116).ArraySet1((_this).F13(), 12)
			(_nw116).ArraySet1((_this).F13(), 13)
			(_nw116).ArraySet1((_this).F13(), 14)
			(_nw116).ArraySet1((_631_i5).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_639_v40, (Companion_Default___.SafeIndex(_593_v0, _dafny.IntOfUint32((_639_v40).Cardinality()))).Uint32(), !((_this).F13()))).Cardinality())) > 0, 15)
			(_nw116).ArraySet1(true, 16)
			(_nw116).ArraySet1(!((!(!((_this).F13()))) || ((_640_v41).Dtor_cf53())), 17)
			(_nw116).ArraySet1((_this).F13(), 18)
			(_nw116).ArraySet1(_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm21(globalState), _638_v39), 19)
			(_nw116).ArraySet1((_this).F13(), 20)
			(_nw116).ArraySet1(Companion_Default___.Fm1((_this).F13(), globalState), 21)
			(_nw116).ArraySet1(!((_this).F13()), 22)
			(_nw116).ArraySet1(!((_this).F13()), 23)
			(_nw116).ArraySet1((_this).F13(), 24)
			(_nw116).ArraySet1((_this).F13(), 25)
			(_nw116).ArraySet1((_this).F13(), 26)
			(_nw116).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_642_v43, _642_v43), 27)
			(_nw116).ArraySet1((func() bool {
				if true {
					return (_this).F13()
				}
				return (_641_v42).F10()
			})(), 28)
			_643_v44 = _nw116
			_643_v44 = (_641_v42).F11()
			var _644_v45 _dafny.Array
			_ = _644_v45
			var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
			_ = _nw117
			_644_v45 = _nw117
			_644_v45 = _644_v45
			_641_v42 = _641_v42
			(globalState).F5 = Companion_Default___.SafeDivisionInt(_631_i5, _631_i5)
		}
		if (_this).F13() {
			var _645_v46 _dafny.Sequence
			_ = _645_v46
			_645_v46 = _dafny.SeqOf(_593_v0)
			var _646_v47 _dafny.Sequence
			_ = _646_v47
			_646_v47 = _645_v46
			_646_v47 = _646_v47
			var _647_v48 _dafny.Array
			_ = _647_v48
			var _nwElement0_19 _dafny.Int = _593_v0
			_ = _nwElement0_19
			var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(17))
			_ = _nw118
			(_nw118).ArraySet1(_nwElement0_19, 0)
			(_nw118).ArraySet1(_593_v0, 1)
			(_nw118).ArraySet1(_dafny.IntOfInt64(158), 2)
			(_nw118).ArraySet1((_593_v0).Times(_593_v0), 3)
			(_nw118).ArraySet1((_645_v46).Select((Companion_Default___.SafeIndex(_593_v0, _dafny.IntOfUint32((_645_v46).Cardinality()))).Uint32()).(_dafny.Int), 4)
			(_nw118).ArraySet1(_593_v0, 5)
			(_nw118).ArraySet1(_dafny.IntOfInt64(262), 6)
			(_nw118).ArraySet1(_593_v0, 7)
			(_nw118).ArraySet1(_593_v0, 8)
			(_nw118).ArraySet1(_593_v0, 9)
			(_nw118).ArraySet1(_dafny.IntOfInt64(-672), 10)
			(_nw118).ArraySet1(_593_v0, 11)
			(_nw118).ArraySet1((_dafny.IntOfUint32((_645_v46).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-606))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_648_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}))).Cardinality())), 12)
			(_nw118).ArraySet1(_593_v0, 13)
			(_nw118).ArraySet1(_593_v0, 14)
			(_nw118).ArraySet1(_593_v0, 15)
			(_nw118).ArraySet1(_593_v0, 16)
			_647_v48 = _nw118
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_647_v48), 0))
			_ = _index86
			(_647_v48).ArraySet1((_dafny.Zero).Minus(_593_v0), (_index86).Int())
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_647_v48), 0))
			_ = _index87
			var _rhs97 _dafny.Int = _593_v0
			_ = _rhs97
			var _rhs98 _dafny.Int = (_dafny.Zero).Minus(_593_v0)
			_ = _rhs98
			var _lhs72 _dafny.Array = _647_v48
			_ = _lhs72
			var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_647_v48), 0))
			_ = _lhs73
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			(_lhs72).ArraySet1(_rhs97, (_lhs73).Int())
			_lhs74.F5 = _rhs98
			var _649_v49 _dafny.Map
			_ = _649_v49
			_649_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v0, !(!((_this).F13())))
			var _650_v50 _dafny.Sequence
			_ = _650_v50
			_650_v50 = _dafny.SeqOf((_this).F13())
			var _651_v51 D0
			_ = _651_v51
			_651_v51 = Companion_D0_.Create_DC0_(_649_v49)
			var _652_v52 _dafny.MultiSet
			_ = _652_v52
			_652_v52 = _dafny.MultiSetOf((_this).F13())
			var _pat_let_tv15 = _647_v48
			_ = _pat_let_tv15
			var _pat_let_tv16 = _647_v48
			_ = _pat_let_tv16
			var _653_v53 _dafny.Array
			_ = _653_v53
			var _nwElement0_20 bool = (_this).F13()
			_ = _nwElement0_20
			var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(22))
			_ = _nw119
			(_nw119).ArraySet1(_nwElement0_20, 0)
			(_nw119).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((_this).F13()), _dafny.SeqOf(Companion_Default___.Fm1((_this).F13(), globalState))), 1)
			(_nw119).ArraySet1((_this).F13(), 2)
			(_nw119).ArraySet1((func() bool {
				if (_649_v49).Contains(_593_v0) {
					return (_649_v49).Get(_593_v0).(bool)
				}
				return !((_this).F13())
			})(), 3)
			(_nw119).ArraySet1((_this).F13(), 4)
			(_nw119).ArraySet1(((_647_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_647_v48), 0))).Int()).(_dafny.Int)).Cmp(_593_v0) > 0, 5)
			(_nw119).ArraySet1((_this).F13(), 6)
			(_nw119).ArraySet1(!((_this).F13()), 7)
			(_nw119).ArraySet1(!((_this).F13()), 8)
			(_nw119).ArraySet1((_this).F13(), 9)
			(_nw119).ArraySet1((_this).F13(), 10)
			(_nw119).ArraySet1((_this).F13(), 11)
			(_nw119).ArraySet1((_this).F13(), 12)
			(_nw119).ArraySet1((_this).F13(), 13)
			(_nw119).ArraySet1(!((func() bool {
				if (_649_v49).Contains(_593_v0) {
					return (_649_v49).Get(_593_v0).(bool)
				}
				return (_this).F13()
			})()), 14)
			(_nw119).ArraySet1(!_dafny.Companion_Sequence_.Equal(_650_v50, _dafny.SeqOf((_this).F13(), (_this).F13())), 15)
			(_nw119).ArraySet1(false, 16)
			(_nw119).ArraySet1((_this).F13(), 17)
			(_nw119).ArraySet1((_this).F13(), 18)
			(_nw119).ArraySet1((_this).F13(), 19)
			(_nw119).ArraySet1(!((_this).F13()), 20)
			(_nw119).ArraySet1((func() bool {
				if (_630_v32).Fm4(func(_pat_let8_0 D0) D0 {
					return func(_654_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let9_0 _dafny.Map) D0 {
							return func(_655_dt__update_hcf0_h0 _dafny.Map) D0 {
								return Companion_D0_.Create_DC0_(_655_dt__update_hcf0_h0)
							}(_pat_let9_0)
						}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_pat_let_tv15), 0))).Int()).(_dafny.Int), (_this).F13()))
					}(_pat_let8_0)
				}(_651_v51), _651_v51, _652_v52, globalState) {
					return true
				}
				return (_this).F13()
			})(), 21)
			_653_v53 = _nw119
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_653_v53), 0))
			_ = _index88
			(_653_v53).ArraySet1((_this).F13(), (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_653_v53), 0))
			_ = _index89
			(_653_v53).ArraySet1(!((_this).F13()), (_index89).Int())
			var _656_v54 _dafny.CodePoint
			_ = _656_v54
			_656_v54 = _dafny.CodePoint('p')
			_656_v54 = _656_v54
			var _657_v55 _dafny.Array
			_ = _657_v55
			var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
			_ = _nw120
			_657_v55 = _nw120
			_657_v55 = _657_v55
		} else {
			var _658_v56 *C1
			_ = _658_v56
			var _nw121 *C1 = New_C1_()
			_ = _nw121
			_nw121.Ctor__()
			_658_v56 = _nw121
			(globalState).F2 = (_this).F13()
			var _659_v57 _dafny.Map
			_ = _659_v57
			_659_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v0, _dafny.IntOfInt64(-311))
			var _660_v58 _dafny.Set
			_ = _660_v58
			_660_v58 = _dafny.SetOf((_659_v57).Cardinality())
			if ((Companion_Default___.Fm30(globalState)).Intersection(_660_v58)).IsProperSubsetOf(_660_v58) {
				var _661_v59 _dafny.Array
				_ = _661_v59
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_21
				var _nw122 _dafny.Array
				_ = _nw122
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw122 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) bool = func(_662_i7 _dafny.Int) bool {
						return false
					}
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw122 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw122).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw122).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_661_v59 = _nw122
				_661_v59 = _661_v59
				var _663_v60 *C1
				_ = _663_v60
				var _nw123 *C1 = New_C1_()
				_ = _nw123
				_nw123.Ctor__()
				_663_v60 = _nw123
				(globalState).F0 = (_this).F13()
				var _664_v61 *C0
				_ = _664_v61
				var _nw124 *C0 = New_C0_()
				_ = _nw124
				_nw124.Ctor__((_this).F13(), _661_v59)
				_664_v61 = _nw124
				_664_v61 = _664_v61
				var _665_v62 _dafny.Array
				_ = _665_v62
				var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw125
				_665_v62 = _nw125
				var _666_v63 _dafny.MultiSet
				_ = _666_v63
				_666_v63 = _dafny.MultiSetOf(_665_v62, _665_v62)
				var _667_v64 _dafny.Map
				_ = _667_v64
				_667_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_630_v32, _666_v63)
				var _668_v65 _dafny.Sequence
				_ = _668_v65
				_668_v65 = _dafny.SeqOf(_593_v0)
				var _669_v66 _dafny.Array
				_ = _669_v66
				_669_v66 = _665_v62
				var _670_v67 _dafny.Map
				_ = _670_v67
				_670_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_664_v61).F10()), _dafny.CodePoint('v'))
				var _671_v68 _dafny.Sequence
				_ = _671_v68
				_671_v68 = _dafny.SeqOf(_666_v63, _666_v63, _666_v63, ((_666_v63).Update(_665_v62, Companion_Default___.Abs((_670_v67).Cardinality()))).Update(_665_v62, Companion_Default___.Abs(_593_v0)), _666_v63)
				var _672_v69 D2
				_ = _672_v69
				_672_v69 = Companion_D2_.Create_DC6_((_this).F13())
				var _673_v70 _dafny.Map
				_ = _673_v70
				_673_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_664_v61).F10(), _593_v0)
				var _674_v71 _dafny.CodePoint
				_ = _674_v71
				_674_v71 = _dafny.CodePoint('q')
				var _675_v72 _dafny.Map
				_ = _675_v72
				_675_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm18(_672_v69, (_660_v58).Cardinality(), (_673_v70).Cardinality(), globalState), _674_v71)
				var _676_v73 _dafny.Array
				_ = _676_v73
				var _nwElement0_21 _dafny.MultiSet = _666_v63
				_ = _nwElement0_21
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(19))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_21, 0)
				(_nw126).ArraySet1(_dafny.MultiSetOf(_665_v62), 1)
				(_nw126).ArraySet1(((func() _dafny.MultiSet {
					if (_667_v64).Contains(_658_v56) {
						return (_667_v64).Get(_658_v56).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf(_665_v62)
				})()).Update(_665_v62, Companion_Default___.Abs(_dafny.IntOfUint32((_668_v65).Cardinality()))), 2)
				(_nw126).ArraySet1((_666_v63).Update((_669_v66), Companion_Default___.Abs(_dafny.IntOfInt64(89))), 3)
				(_nw126).ArraySet1((_666_v63).Union(_666_v63), 4)
				(_nw126).ArraySet1((_666_v63).Difference(_666_v63), 5)
				(_nw126).ArraySet1(_dafny.MultiSetOf(_665_v62, _665_v62), 6)
				(_nw126).ArraySet1(_666_v63, 7)
				(_nw126).ArraySet1((_671_v68).Select((Companion_Default___.SafeIndex(_593_v0, _dafny.IntOfUint32((_671_v68).Cardinality()))).Uint32()).(_dafny.MultiSet), 8)
				(_nw126).ArraySet1((_666_v63).Difference((_666_v63).Update(_665_v62, Companion_Default___.Abs((_675_v72).Cardinality()))), 9)
				(_nw126).ArraySet1((func() _dafny.MultiSet {
					if Companion_Default___.Fm1((_this).F13(), globalState) {
						return _666_v63
					}
					return _666_v63
				})(), 10)
				(_nw126).ArraySet1(_666_v63, 11)
				(_nw126).ArraySet1(_666_v63, 12)
				(_nw126).ArraySet1(_666_v63, 13)
				(_nw126).ArraySet1(_666_v63, 14)
				(_nw126).ArraySet1(_dafny.MultiSetOf(_665_v62, _665_v62, _665_v62), 15)
				(_nw126).ArraySet1((_666_v63).Union(_666_v63), 16)
				(_nw126).ArraySet1((_666_v63).Intersection(_dafny.MultiSetOf(_665_v62)), 17)
				(_nw126).ArraySet1((_671_v68).Select((Companion_Default___.SafeIndex(_593_v0, _dafny.IntOfUint32((_671_v68).Cardinality()))).Uint32()).(_dafny.MultiSet), 18)
				_676_v73 = _nw126
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_676_v73), 0))
				_ = _index90
				(_676_v73).ArraySet1(_666_v63, (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_676_v73), 0))
				_ = _index91
				(_676_v73).ArraySet1(_666_v63, (_index91).Int())
			} else {
				var _677_v74 _dafny.Set
				_ = _677_v74
				_677_v74 = _dafny.SetOf(_659_v57, _659_v57)
				var _678_v75 _dafny.Array
				_ = _678_v75
				var _nwElement0_22 _dafny.Int = _593_v0
				_ = _nwElement0_22
				var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(5))
				_ = _nw127
				(_nw127).ArraySet1(_nwElement0_22, 0)
				(_nw127).ArraySet1((func() _dafny.Int {
					if (_this).F13() {
						return _593_v0
					}
					return _593_v0
				})(), 1)
				(_nw127).ArraySet1(_593_v0, 2)
				(_nw127).ArraySet1(_593_v0, 3)
				(_nw127).ArraySet1((_677_v74).Cardinality(), 4)
				_678_v75 = _nw127
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_678_v75), 0))
				_ = _index92
				(_678_v75).ArraySet1(_593_v0, (_index92).Int())
				var _679_v76 _dafny.Sequence
				_ = _679_v76
				_679_v76 = _dafny.SeqOf((_this).F13())
				var _680_v77 _dafny.Sequence
				_ = _680_v77
				_680_v77 = _dafny.SeqOf(_679_v76)
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_678_v75), 0))
				_ = _index93
				var _rhs99 _dafny.Int = (_dafny.Zero).Minus(_593_v0)
				_ = _rhs99
				var _rhs100 _dafny.Array = _678_v75
				_ = _rhs100
				var _rhs101 _dafny.Set = Companion_Default___.Fm30(globalState)
				_ = _rhs101
				var _rhs102 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_680_v77, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-310))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_681_i8 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F13())
				})))).Cardinality()))
				_ = _rhs102
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				var _lhs76 _dafny.Array = _678_v75
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_678_v75), 0))
				_ = _lhs77
				_lhs75.F5 = _rhs99
				_678_v75 = _rhs100
				_660_v58 = _rhs101
				(_lhs76).ArraySet1(_rhs102, (_lhs77).Int())
				var _682_v78 _dafny.Map
				_ = _682_v78
				_682_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
				var _683_v79 D2
				_ = _683_v79
				_683_v79 = Companion_D2_.Create_DC7_(_593_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_682_v78, _593_v0), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fjam")).Cardinality())))
				var _684_v80 _dafny.MultiSet
				_ = _684_v80
				_684_v80 = _dafny.MultiSetOf((func() D2 {
					if !((_this).F13()) {
						return _683_v79
					}
					return _683_v79
				})())
				_684_v80 = (_684_v80).Union(_684_v80)
				var _685_v81 _dafny.CodePoint
				_ = _685_v81
				_685_v81 = _dafny.CodePoint('p')
				_685_v81 = _685_v81
				var _686_v82 *C1
				_ = _686_v82
				var _nw128 *C1 = New_C1_()
				_ = _nw128
				_nw128.Ctor__()
				_686_v82 = _nw128
				(globalState).F0 = (_this).F13()
			}
			var _687_v83 _dafny.Array
			_ = _687_v83
			var _nwElement0_23 bool = (_this).F13()
			_ = _nwElement0_23
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(24))
			_ = _nw129
			(_nw129).ArraySet1(_nwElement0_23, 0)
			(_nw129).ArraySet1((_this).F13(), 1)
			(_nw129).ArraySet1((_this).F13(), 2)
			(_nw129).ArraySet1((_this).F13(), 3)
			(_nw129).ArraySet1(Companion_Default___.Fm1((_this).F13(), globalState), 4)
			(_nw129).ArraySet1(false, 5)
			(_nw129).ArraySet1((_this).F13(), 6)
			(_nw129).ArraySet1(!((_this).F13()), 7)
			(_nw129).ArraySet1((_this).F13(), 8)
			(_nw129).ArraySet1((_this).F13(), 9)
			(_nw129).ArraySet1((_this).F13(), 10)
			(_nw129).ArraySet1((_this).F13(), 11)
			(_nw129).ArraySet1((_this).F13(), 12)
			(_nw129).ArraySet1(!((_this).F13()), 13)
			(_nw129).ArraySet1((_this).F13(), 14)
			(_nw129).ArraySet1((_this).F13(), 15)
			(_nw129).ArraySet1((_this).F13(), 16)
			(_nw129).ArraySet1((_this).F13(), 17)
			(_nw129).ArraySet1((_this).F13(), 18)
			(_nw129).ArraySet1(!((_this).F13()), 19)
			(_nw129).ArraySet1((_this).F13(), 20)
			(_nw129).ArraySet1((_this).F13(), 21)
			(_nw129).ArraySet1((_this).F13(), 22)
			(_nw129).ArraySet1((_this).F13(), 23)
			_687_v83 = _nw129
			var _688_v84 *C0
			_ = _688_v84
			var _nw130 *C0 = New_C0_()
			_ = _nw130
			_nw130.Ctor__((_this).F13(), _687_v83)
			_688_v84 = _nw130
			var _689_v85 _dafny.Sequence
			_ = _689_v85
			_689_v85 = _dafny.UnicodeSeqOfUtf8Bytes("pswtnsaun")
			(globalState).F0 = !(!(false) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_689_v85, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-527))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_690_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})))))
		}
		var _691_v86 _dafny.Array
		_ = _691_v86
		var _len0_22 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_22
		var _nw131 _dafny.Array
		_ = _nw131
		if _len0_22.Cmp(_dafny.Zero) == 0 {
			_nw131 = _dafny.NewArray(_len0_22)
		} else {
			var _init22 func(_dafny.Int) bool = func(_692_i11 _dafny.Int) bool {
				return false
			}
			_ = _init22
			var _element0_22 = _init22(_dafny.Zero)
			_ = _element0_22
			_nw131 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
			(_nw131).ArraySet1(_element0_22, 0)
			var _nativeLen0_22 = (_len0_22).Int()
			_ = _nativeLen0_22
			for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
				(_nw131).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
			}
		}
		_691_v86 = _nw131
		var _693_v87 _dafny.Sequence
		_ = _693_v87
		_693_v87 = _dafny.UnicodeSeqOfUtf8Bytes("gvhxsqh")
		var _694_v88 _dafny.Map
		_ = _694_v88
		_694_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v86, _693_v87)
		var _695_v89 _dafny.Sequence
		_ = _695_v89
		_695_v89 = _dafny.SeqOf((_694_v88).Cardinality())
		r0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(691))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}((func(_696_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_697_i10 _dafny.Int) _dafny.Int {
				return _696_v0
			}
		})(_593_v0))), _695_v89)
		var _698_v90 _dafny.MultiSet
		_ = _698_v90
		_698_v90 = _dafny.MultiSetOf(_dafny.IntOfInt64(507))
		var _699_v91 _dafny.Map
		_ = _699_v91
		_699_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _698_v90)
		r1 = (func() _dafny.MultiSet {
			if (_699_v91).Contains((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v0, !((_this).F13()))).Update(_593_v0, (_this).F13())).Cardinality()) != 0) {
				return (_699_v91).Get((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v0, !((_this).F13()))).Update(_593_v0, (_this).F13())).Cardinality()) != 0).(_dafny.MultiSet)
			}
			return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_700_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_701_i12 _dafny.Int) _dafny.Int {
					return _700_v0
				}
			})(_593_v0))))
		})()
		var _702_v92 _dafny.Map
		_ = _702_v92
		_702_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v87, (_this).F13())
		r2 = _702_v92
		return r0, r1, r2
	}
}
func (_this *C2) M5(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _703_v0 _dafny.Array
		_ = _703_v0
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_23
		var _nw132 _dafny.Array
		_ = _nw132
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw132 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) bool = func(_704_i0 _dafny.Int) bool {
				return (_this).F13()
			}
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw132 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw132).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw132).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_703_v0 = _nw132
		var _705_v1 *C0
		_ = _705_v1
		var _nw133 *C0 = New_C0_()
		_ = _nw133
		_nw133.Ctor__((_this).F13(), _703_v0)
		_705_v1 = _nw133
		var _706_v2 D2
		_ = _706_v2
		_706_v2 = Companion_D2_.Create_DC6_((_705_v1).F10())
		var _707_v3 _dafny.Map
		_ = _707_v3
		_707_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm18(_706_v2, p1, p1, globalState), !((_this).F13()))
		var _708_v4 _dafny.Set
		_ = _708_v4
		_708_v4 = _dafny.SetOf((_this).F13(), (_705_v1).F10())
		_707_v3 = (_707_v3).Update(((_708_v4).Cardinality()).Minus(p1), !((_this).F13()))
		var _709_v5 _dafny.CodePoint
		_ = _709_v5
		_709_v5 = _dafny.CodePoint('o')
		_709_v5 = Companion_Default___.Fm20(Companion_Default___.SafeDivisionInt(p1, p1), (_this).F13(), p1, globalState)
		var _710_i1 _dafny.Int
		_ = _710_i1
		_710_i1 = _dafny.Zero
		{
			for false {
				{
					if (_710_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_710_i1 = (_710_i1).Plus(_dafny.One)
					var _711_v6 _dafny.Array
					_ = _711_v6
					var _len0_24 _dafny.Int = _dafny.IntOfInt64(15)
					_ = _len0_24
					var _nw134 _dafny.Array
					_ = _nw134
					if _len0_24.Cmp(_dafny.Zero) == 0 {
						_nw134 = _dafny.NewArray(_len0_24)
					} else {
						var _init24 func(_dafny.Int) _dafny.Int = func(_712_i2 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_712_i2, _dafny.IntOfInt64(349))
						}
						_ = _init24
						var _element0_24 = _init24(_dafny.Zero)
						_ = _element0_24
						_nw134 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
						(_nw134).ArraySet1(_element0_24, 0)
						var _nativeLen0_24 = (_len0_24).Int()
						_ = _nativeLen0_24
						for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
							(_nw134).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
						}
					}
					_711_v6 = _nw134
					var _713_v7 _dafny.Map
					_ = _713_v7
					_713_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _711_v6)
					var _714_v8 _dafny.Map
					_ = _714_v8
					_714_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _711_v6)
					var _715_v9 _dafny.Array
					_ = _715_v9
					_715_v9 = _711_v6
					var _716_v10 _dafny.Array
					_ = _716_v10
					var _nwElement0_24 _dafny.Array = _711_v6
					_ = _nwElement0_24
					var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(18))
					_ = _nw135
					(_nw135).ArraySet1(_nwElement0_24, 0)
					(_nw135).ArraySet1(_711_v6, 1)
					(_nw135).ArraySet1(_711_v6, 2)
					(_nw135).ArraySet1(_711_v6, 3)
					(_nw135).ArraySet1(_711_v6, 4)
					(_nw135).ArraySet1(_711_v6, 5)
					(_nw135).ArraySet1(_711_v6, 6)
					(_nw135).ArraySet1(_711_v6, 7)
					(_nw135).ArraySet1(_711_v6, 8)
					(_nw135).ArraySet1(_711_v6, 9)
					(_nw135).ArraySet1(_711_v6, 10)
					(_nw135).ArraySet1(_711_v6, 11)
					(_nw135).ArraySet1(_711_v6, 12)
					(_nw135).ArraySet1(_711_v6, 13)
					(_nw135).ArraySet1((func() _dafny.Array {
						if (_713_v7).Contains((_705_v1).F10()) {
							return (_713_v7).Get((_705_v1).F10()).(_dafny.Array)
						}
						return (func() _dafny.Array {
							if (_714_v8).Contains(p1) {
								return (_714_v8).Get(p1).(_dafny.Array)
							}
							return _711_v6
						})()
					})(), 14)
					(_nw135).ArraySet1(_711_v6, 15)
					(_nw135).ArraySet1((_715_v9), 16)
					(_nw135).ArraySet1(_711_v6, 17)
					_716_v10 = _nw135
					var _rhs103 _dafny.Array = _716_v10
					_ = _rhs103
					var _rhs104 _dafny.Array = _703_v0
					_ = _rhs104
					_716_v10 = _rhs103
					_703_v0 = _rhs104
					var _717_v11 _dafny.Sequence
					_ = _717_v11
					_717_v11 = _dafny.SeqOf(_709_v5, _dafny.CodePoint('y'), _709_v5)
					_709_v5 = (_717_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.IntOfUint32((_717_v11).Cardinality()))).Uint32()).(_dafny.CodePoint)
					var _718_v12 _dafny.Sequence
					_ = _718_v12
					_718_v12 = _dafny.SeqOf((_705_v1).F10())
					(globalState).F0 = (func() bool {
						if (_this).F13() {
							return (_718_v12).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_718_v12).Cardinality()))).Uint32()).(bool)
						}
						return !(_dafny.MultiSetOf((_705_v1).F10(), (_705_v1).F10())).Contains((_705_v1).F10())
					})()
					(globalState).F0 = (_this).F13()
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _719_v13 _dafny.Sequence
		_ = _719_v13
		_719_v13 = _dafny.SeqOf((_dafny.Zero).Minus(p1), p1, p1)
		var _720_v14 _dafny.Map
		_ = _720_v14
		_720_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(364), (_dafny.Zero).Minus(p1))
		var _721_v15 _dafny.Map
		_ = _721_v15
		_721_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_720_v14).Cardinality(), p1)
		var _722_i3 _dafny.Int
		_ = _722_i3
		_722_i3 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.IsPrefixOf(_719_v13, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfInt64(599)), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_721_v15).Contains(_dafny.IntOfInt64(820)) {
					return (_721_v15).Get(_dafny.IntOfInt64(820)).(_dafny.Int)
				}
				return (_719_v13).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_719_v13).Cardinality()))).Uint32()).(_dafny.Int)
			})(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(599))).Cardinality()))).Uint32(), p1))) {
				{
					if (_722_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_722_i3 = (_722_i3).Plus(_dafny.One)
					var _723_v16 *C1
					_ = _723_v16
					var _nw136 *C1 = New_C1_()
					_ = _nw136
					_nw136.Ctor__()
					_723_v16 = _nw136
					var _724_v17 _dafny.MultiSet
					_ = _724_v17
					_724_v17 = _dafny.MultiSetOf((_705_v1).F10())
					_724_v17 = (_724_v17).Union(_724_v17)
					(globalState).F5 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, _dafny.IntOfInt64(538)))
					var _725_v18 _dafny.Map
					_ = _725_v18
					_725_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false) || ((_this).F13()), (p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Array))
					_725_v18 = _725_v18
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		_721_v15 = (_721_v15).Update(p1, p1)
		var _726_v19 _dafny.Map
		_ = _726_v19
		_726_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_705_v1).F10()), _dafny.IntOfInt64(381))
		r0 = Companion_Default___.SafeDivisionInt((p1).Minus(p1), (func() _dafny.Int {
			if (_726_v19).Contains((_this).F13()) {
				return (_726_v19).Get((_this).F13()).(_dafny.Int)
			}
			return p1
		})())
		r1 = p1
		var _727_v20 D12
		_ = _727_v20
		_727_v20 = Companion_D12_.Create_DC36_((_705_v1).F10(), p1, false)
		var _pat_let_tv17 = _705_v1
		_ = _pat_let_tv17
		var _pat_let_tv18 = _719_v13
		_ = _pat_let_tv18
		var _pat_let_tv19 = _719_v13
		_ = _pat_let_tv19
		var _pat_let_tv20 = _705_v1
		_ = _pat_let_tv20
		var _pat_let_tv21 = _705_v1
		_ = _pat_let_tv21
		var _pat_let_tv22 = _705_v1
		_ = _pat_let_tv22
		r2 = func(_source10 D12) bool {
			if _source10.Is_DC34() {
				var _728___mcc_h0 D0 = _source10.Get_().(D12_DC34).Cf48
				_ = _728___mcc_h0
				var _729___mcc_h1 bool = _source10.Get_().(D12_DC34).Cf49
				_ = _729___mcc_h1
				var _730___mcc_h2 bool = _source10.Get_().(D12_DC34).Cf50
				_ = _730___mcc_h2
				var _731_cf50 bool = _730___mcc_h2
				_ = _731_cf50
				var _732_cf49 bool = _729___mcc_h1
				_ = _732_cf49
				var _733_cf48 D0 = _728___mcc_h0
				_ = _733_cf48
				return (_pat_let_tv17).F10()
			} else if _source10.Is_DC35() {
				return _dafny.Companion_Sequence_.IsPrefixOf(_pat_let_tv18, _pat_let_tv19)
			} else if _source10.Is_DC36() {
				var _734___mcc_h3 bool = _source10.Get_().(D12_DC36).Cf51
				_ = _734___mcc_h3
				var _735___mcc_h4 _dafny.Int = _source10.Get_().(D12_DC36).Cf52
				_ = _735___mcc_h4
				var _736___mcc_h5 bool = _source10.Get_().(D12_DC36).Cf53
				_ = _736___mcc_h5
				var _737_cf53 bool = _736___mcc_h5
				_ = _737_cf53
				var _738_cf52 _dafny.Int = _735___mcc_h4
				_ = _738_cf52
				var _739_cf51 bool = _734___mcc_h3
				_ = _739_cf51
				return (_738_cf52).Cmp((_dafny.MultiSetOf(Companion_D0_.Create_DC1_((_pat_let_tv20).F10()), Companion_D0_.Create_DC1_(_739_cf51), Companion_D0_.Create_DC1_(_739_cf51))).Cardinality()) != 0
			} else if _source10.Is_DC33() {
				var _740___mcc_h6 _dafny.Set = _source10.Get_().(D12_DC33).Cf47
				_ = _740___mcc_h6
				var _741_cf47 _dafny.Set = _740___mcc_h6
				_ = _741_cf47
				return (_pat_let_tv21).F10()
			} else {
				var _742___mcc_h7 D12 = _source10.Get_().(D12_DC37).Cf54
				_ = _742___mcc_h7
				var _743_cf54 D12 = _742___mcc_h7
				_ = _743_cf54
				return (_pat_let_tv22).F10()
			}
		}(_727_v20)
		return r0, r1, r2
	}
}
func (_this *C2) F13() bool {
	{
		return _this._f13
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f12 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f12 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C3{}
var _ T1 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__(f12 _dafny.Int) {
	{
		(_this)._f12 = f12
	}
}
func (_this *C3) Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return ((_this).F12()).Cmp((_this).F12()) <= 0
	}
}
func (_this *C3) Fm11(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vd")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F12()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(!(!(false)))))), _dafny.IntOfInt64(-214))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F12())))
	}
}
func (_this *C3) Fm12(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.SeqOf((_this).F12()), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F12())).Cardinality(), (_this).F12())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf((_this).F12(), (_this).F12(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cvi")).Cardinality()), (_dafny.Zero).Minus((_this).F12())), _dafny.SeqOf((_this).F12()), _dafny.SeqOf((_this).F12()), _dafny.SeqOf((_this).F12(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-877))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_744_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf((_this).F12(), (_this).F12())
		}))))
	}
}
func (_this *C3) M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _745_v0 bool
		_ = _745_v0
		_745_v0 = true
		if !(_745_v0) || (_745_v0) {
			var _746_v1 _dafny.Array
			_ = _746_v1
			var _len0_25 _dafny.Int = _dafny.One
			_ = _len0_25
			var _nw137 _dafny.Array
			_ = _nw137
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw137 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_747_v0 bool) func(_dafny.Int) bool {
					return func(_748_i0 _dafny.Int) bool {
						return (Companion_D0_.Create_DC1_(_747_v0)).Dtor_cf1()
					}
				})(_745_v0)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw137 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw137).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw137).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_746_v1 = _nw137
			_746_v1 = _746_v1
			if _745_v0 {
				var _749_v2 _dafny.Map
				_ = _749_v2
				_749_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _745_v0)
				var _750_v3 _dafny.Array
				_ = _750_v3
				var _nwElement0_25 _dafny.Int = p0
				_ = _nwElement0_25
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(6))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_25, 0)
				(_nw138).ArraySet1((_this).F12(), 1)
				(_nw138).ArraySet1((_749_v2).Cardinality(), 2)
				(_nw138).ArraySet1(p2, 3)
				(_nw138).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqcgspfpy")).Cardinality())), 4)
				(_nw138).ArraySet1((_this).F12(), 5)
				_750_v3 = _nw138
				var _751_v4 _dafny.Array
				_ = _751_v4
				_751_v4 = _750_v3
				var _752_v5 _dafny.Sequence
				_ = _752_v5
				_752_v5 = _dafny.SeqOf(_750_v3, (_750_v3))
				var _753_v6 _dafny.Set
				_ = _753_v6
				_753_v6 = _dafny.SetOf(_745_v0, _745_v0)
				var _754_v7 _dafny.Set
				_ = _754_v7
				_754_v7 = _dafny.SetOf(_750_v3, _750_v3, (_752_v5).Select((Companion_Default___.SafeIndex((_753_v6).Cardinality(), _dafny.IntOfUint32((_752_v5).Cardinality()))).Uint32()).(_dafny.Array))
				_754_v7 = _754_v7
				var _755_v8 _dafny.MultiSet
				_ = _755_v8
				_755_v8 = _dafny.MultiSetOf(Companion_Default___.Fm1(_745_v0, globalState), _745_v0, _745_v0)
				var _756_v9 _dafny.Sequence
				_ = _756_v9
				_756_v9 = _dafny.UnicodeSeqOfUtf8Bytes("axeend")
				var _rhs105 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_756_v9, _756_v9), _dafny.Companion_Sequence_.Concatenate(_756_v9, _756_v9))
				_ = _rhs105
				var _rhs106 _dafny.MultiSet = _755_v8
				_ = _rhs106
				r0 = _rhs105
				_755_v8 = _rhs106
				var _757_v10 _dafny.CodePoint
				_ = _757_v10
				_757_v10 = _dafny.CodePoint('u')
				_757_v10 = _757_v10
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((p1), 0))
				_ = _index94
				(p1).ArraySet1(_750_v3, (_index94).Int())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((p1), 0))
				_ = _index95
				(p1).ArraySet1(_750_v3, (_index95).Int())
				var _758_v11 D1
				_ = _758_v11
				_758_v11 = Companion_D1_.Create_DC4_(p3)
				var _759_v12 _dafny.Map
				_ = _759_v12
				_759_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_758_v11, p3)
				var _760_v13 _dafny.Map
				_ = _760_v13
				_760_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _759_v12)
				_760_v13 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_758_v11, _dafny.IntOfInt64(-438)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _759_v12))
			} else {
				var _761_v14 _dafny.Sequence
				_ = _761_v14
				_761_v14 = _dafny.SeqOf(_745_v0, !(!(_745_v0)))
				var _762_v15 *C0
				_ = _762_v15
				var _nw139 *C0 = New_C0_()
				_ = _nw139
				_nw139.Ctor__((_761_v14).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_761_v14).Cardinality()))).Uint32()).(bool), _746_v1)
				_762_v15 = _nw139
				var _763_v16 _dafny.Array
				_ = _763_v16
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_26
				var _nw140 _dafny.Array
				_ = _nw140
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw140 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Int = (func(_764_v0 bool) func(_dafny.Int) _dafny.Int {
						return func(_765_i1 _dafny.Int) _dafny.Int {
							return (_765_i1).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_764_v0, true)).Cardinality())
						}
					})(_745_v0)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw140 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw140).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw140).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_763_v16 = _nw140
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_763_v16), 0))
				_ = _index96
				(_763_v16).ArraySet1(p0, (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_763_v16), 0))
				_ = _index97
				(_763_v16).ArraySet1((_this).F12(), (_index97).Int())
				var _766_v17 _dafny.Set
				_ = _766_v17
				_766_v17 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, (_762_v15).F10()))
				_766_v17 = Companion_Default___.Fm13(globalState)
				var _767_v18 D4
				_ = _767_v18
				_767_v18 = Companion_D4_.Create_DC14_()
				var _rhs107 _dafny.Int = (func() _dafny.Int {
					if (_762_v15).F10() {
						return (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, p2))
					}
					return Companion_Default___.SafeDivisionInt((_this).F12(), p2)
				})()
				_ = _rhs107
				var _rhs108 D4 = _767_v18
				_ = _rhs108
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				_lhs78.F5 = _rhs107
				_767_v18 = _rhs108
				var _768_v19 _dafny.Sequence
				_ = _768_v19
				_768_v19 = _dafny.UnicodeSeqOfUtf8Bytes("uqxmp")
				_768_v19 = (func() _dafny.Sequence {
					if _745_v0 {
						return _dafny.Companion_Sequence_.Concatenate(_768_v19, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-716))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg36 _dafny.Int) interface{} {
								return coer36(arg36)
							}
						}(func(_769_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						})))
					}
					return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14(_768_v19, p3, globalState), _dafny.UnicodeSeqOfUtf8Bytes("fraq"))
				})()
			}
			(globalState).F5 = (_dafny.Zero).Minus((p0).Times(_dafny.IntOfUint32((_dafny.SeqOf(_745_v0, _745_v0, _745_v0)).Cardinality())))
			r1 = p2
			var _770_v20 D4
			_ = _770_v20
			_770_v20 = Companion_D4_.Create_DC11_(_746_v1)
			var _rhs109 D4 = _770_v20
			_ = _rhs109
			var _rhs110 bool = false
			_ = _rhs110
			var _lhs79 *GlobalState = globalState
			_ = _lhs79
			_770_v20 = _rhs109
			_lhs79.F0 = _rhs110
		} else {
			var _771_v21 _dafny.Map
			_ = _771_v21
			_771_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _745_v0)
			var _772_v22 _dafny.Set
			_ = _772_v22
			_772_v22 = _dafny.SetOf(_745_v0, _745_v0)
			_771_v21 = (_771_v21).Update((_772_v22).IsDisjointFrom(_772_v22), _745_v0)
			var _773_v23 _dafny.MultiSet
			_ = _773_v23
			_773_v23 = _dafny.MultiSetOf(p3, (_this).F12())
			var _774_v24 _dafny.Sequence
			_ = _774_v24
			_774_v24 = _dafny.SeqOf(p0)
			var _775_v25 _dafny.Sequence
			_ = _775_v25
			_775_v25 = _dafny.SeqOf(_745_v0, _745_v0, _745_v0, _745_v0)
			var _776_v26 _dafny.Map
			_ = _776_v26
			_776_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _745_v0)
			var _777_v27 D0
			_ = _777_v27
			_777_v27 = Companion_D0_.Create_DC0_(_776_v26)
			var _778_v28 _dafny.MultiSet
			_ = _778_v28
			_778_v28 = _dafny.MultiSetOf(false)
			var _779_v29 _dafny.Sequence
			_ = _779_v29
			_779_v29 = _dafny.SeqOf(_745_v0, (_775_v25).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_775_v25).Cardinality()))).Uint32()).(bool), (_this).Fm4(_777_v27, Companion_D0_.Create_DC0_(_776_v26), _778_v28, globalState), true)
			var _rhs111 bool = (_dafny.MultiSetFromSeq(_774_v24)).IsSubsetOf(_773_v23)
			_ = _rhs111
			var _rhs112 _dafny.Int = (Companion_Default___.SafeModuloInt(p3, (_this).F12())).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_779_v29, _779_v29)).Cardinality()))
			_ = _rhs112
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			_lhs80.F0 = _rhs111
			_lhs81.F5 = _rhs112
			var _780_v30 D0
			_ = _780_v30
			_780_v30 = Companion_D0_.Create_DC1_(_745_v0)
			var _source11 D0 = _780_v30
			_ = _source11
			if _source11.Is_DC1() {
				var _781___mcc_h0 bool = _source11.Get_().(D0_DC1).Cf1
				_ = _781___mcc_h0
				var _782_cf1 bool = _781___mcc_h0
				_ = _782_cf1
				(globalState).F0 = _782_cf1
				var _783_v31 _dafny.Map
				_ = _783_v31
				_783_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(52), _dafny.IntOfUint32((_779_v29).Cardinality()))
				var _784_v32 _dafny.Array
				_ = _784_v32
				var _nwElement0_26 bool = !((_780_v30).Dtor_cf1())
				_ = _nwElement0_26
				var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(19))
				_ = _nw141
				(_nw141).ArraySet1(_nwElement0_26, 0)
				(_nw141).ArraySet1(_782_cf1, 1)
				(_nw141).ArraySet1(_745_v0, 2)
				(_nw141).ArraySet1(_745_v0, 3)
				(_nw141).ArraySet1(_782_cf1, 4)
				(_nw141).ArraySet1(_745_v0, 5)
				(_nw141).ArraySet1(true, 6)
				(_nw141).ArraySet1((_this).Fm4(_777_v27, _777_v27, _778_v28, globalState), 7)
				(_nw141).ArraySet1(_782_cf1, 8)
				(_nw141).ArraySet1(_745_v0, 9)
				(_nw141).ArraySet1(_745_v0, 10)
				(_nw141).ArraySet1(_782_cf1, 11)
				(_nw141).ArraySet1(_782_cf1, 12)
				(_nw141).ArraySet1(_745_v0, 13)
				(_nw141).ArraySet1(_782_cf1, 14)
				(_nw141).ArraySet1(_745_v0, 15)
				(_nw141).ArraySet1(_782_cf1, 16)
				(_nw141).ArraySet1((_this).Fm4(_777_v27, _777_v27, _778_v28, globalState), 17)
				(_nw141).ArraySet1(_782_cf1, 18)
				_784_v32 = _nw141
				var _785_v33 *C0
				_ = _785_v33
				var _nw142 *C0 = New_C0_()
				_ = _nw142
				_nw142.Ctor__((p2).Cmp((func() _dafny.Int {
					if (_783_v31).Contains(p2) {
						return (_783_v31).Get(p2).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-979)
				})()) > 0, _784_v32)
				_785_v33 = _nw142
				var _786_v34 _dafny.Sequence
				_ = _786_v34
				_786_v34 = _dafny.UnicodeSeqOfUtf8Bytes("gvwujy")
				_786_v34 = _786_v34
				var _787_v35 _dafny.Array
				_ = _787_v35
				var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
				_ = _nw143
				_787_v35 = _nw143
				_787_v35 = _787_v35
			} else if _source11.Is_DC0() {
				var _788___mcc_h1 _dafny.Map = _source11.Get_().(D0_DC0).Cf0
				_ = _788___mcc_h1
				var _789_cf0 _dafny.Map = _788___mcc_h1
				_ = _789_cf0
				var _790_v36 _dafny.Array
				_ = _790_v36
				var _nwElement0_27 bool = _745_v0
				_ = _nwElement0_27
				var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(29))
				_ = _nw144
				(_nw144).ArraySet1(_nwElement0_27, 0)
				(_nw144).ArraySet1(_745_v0, 1)
				(_nw144).ArraySet1(_745_v0, 2)
				(_nw144).ArraySet1(_745_v0, 3)
				(_nw144).ArraySet1(_745_v0, 4)
				(_nw144).ArraySet1(_745_v0, 5)
				(_nw144).ArraySet1(_745_v0, 6)
				(_nw144).ArraySet1(_745_v0, 7)
				(_nw144).ArraySet1(_745_v0, 8)
				(_nw144).ArraySet1(_745_v0, 9)
				(_nw144).ArraySet1(_745_v0, 10)
				(_nw144).ArraySet1(_745_v0, 11)
				(_nw144).ArraySet1(false, 12)
				(_nw144).ArraySet1(_745_v0, 13)
				(_nw144).ArraySet1(_745_v0, 14)
				(_nw144).ArraySet1(_745_v0, 15)
				(_nw144).ArraySet1(true, 16)
				(_nw144).ArraySet1(_745_v0, 17)
				(_nw144).ArraySet1(_745_v0, 18)
				(_nw144).ArraySet1(_745_v0, 19)
				(_nw144).ArraySet1(!(_745_v0), 20)
				(_nw144).ArraySet1(_745_v0, 21)
				(_nw144).ArraySet1(_745_v0, 22)
				(_nw144).ArraySet1(_745_v0, 23)
				(_nw144).ArraySet1(_745_v0, 24)
				(_nw144).ArraySet1(_745_v0, 25)
				(_nw144).ArraySet1((_775_v25).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.IntOfUint32((_775_v25).Cardinality()))).Uint32()).(bool), 26)
				(_nw144).ArraySet1(_745_v0, 27)
				(_nw144).ArraySet1(_745_v0, 28)
				_790_v36 = _nw144
				var _791_v37 _dafny.Map
				_ = _791_v37
				_791_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_790_v36, _745_v0)
				_791_v37 = (_791_v37).Merge(_791_v37)
				var _792_v38 _dafny.CodePoint
				_ = _792_v38
				_792_v38 = _dafny.CodePoint('x')
				var _793_v40 _dafny.Map
				_ = _793_v40
				_793_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, (_this).F12())
				var _794_v41 _dafny.Sequence
				_ = _794_v41
				_794_v41 = _dafny.SeqOf(_793_v40)
				Companion_Default___.M0(_745_v0, (func() _dafny.CodePoint {
					if _745_v0 {
						return _792_v38
					}
					return _792_v38
				})(), (func() _dafny.Set {
					var _coll39 = _dafny.NewBuilder()
					_ = _coll39
					for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(364), _dafny.IntOfInt64(751))); ; {
						_compr_39, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _795_v39 _dafny.Int
						_795_v39 = interface{}(_compr_39).(_dafny.Int)
						if ((_dafny.IntOfInt64(364)).Cmp(_795_v39) <= 0) && ((_795_v39).Cmp(_dafny.IntOfInt64(751)) < 0) {
							_coll39.Add((_795_v39).Plus(_dafny.IntOfInt64(165)))
						}
					}
					return _coll39.ToSet()
				}()).Cardinality(), (_793_v40).Merge((_794_v41).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_794_v41).Cardinality()))).Uint32()).(_dafny.Map)), globalState)
				var _796_v42 _dafny.Array
				_ = _796_v42
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw145
				_796_v42 = _nw145
				var _797_v43 D1
				_ = _797_v43
				_797_v43 = Companion_D1_.Create_DC3_(_796_v42)
				var _pat_let_tv23 = _796_v42
				_ = _pat_let_tv23
				_797_v43 = func(_pat_let10_0 D1) D1 {
					return func(_798_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let11_0 _dafny.Array) D1 {
							return func(_799_dt__update_hcf3_h0 _dafny.Array) D1 {
								return Companion_D1_.Create_DC3_(_799_dt__update_hcf3_h0)
							}(_pat_let11_0)
						}(_pat_let_tv23)
					}(_pat_let10_0)
				}(_797_v43)
				(globalState).F0 = (func() bool {
					if _745_v0 {
						return _745_v0
					}
					return _745_v0
				})()
			} else {
				var _800___mcc_h2 D0 = _source11.Get_().(D0_DC2).Cf2
				_ = _800___mcc_h2
				var _801_cf2 D0 = _800___mcc_h2
				_ = _801_cf2
				var _802_v44 _dafny.Array
				_ = _802_v44
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(23))
				_ = _nw146
				_802_v44 = _nw146
				var _803_v45 _dafny.MultiSet
				_ = _803_v45
				_803_v45 = _dafny.MultiSetOf(_779_v29, _779_v29, _775_v25)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_802_v44), 0))
				_ = _index98
				(_802_v44).ArraySet1((_803_v45).Intersection((_dafny.MultiSetFromSeq(_dafny.SeqOf(_779_v29, _dafny.SeqOf(_745_v0)))).Update(_775_v25, Companion_Default___.Abs(p0))), (_index98).Int())
				var _804_v46 _dafny.Sequence
				_ = _804_v46
				_804_v46 = _dafny.UnicodeSeqOfUtf8Bytes("ofcufvnr")
				var _805_v47 _dafny.Map
				_ = _805_v47
				_805_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_804_v46).Cardinality()), p0)
				var _806_v48 _dafny.Sequence
				_ = _806_v48
				_806_v48 = _dafny.SeqOf(_805_v47)
				var _807_v49 _dafny.Array
				_ = _807_v49
				var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw147
				_807_v49 = _nw147
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))
				_ = _index99
				(p1).ArraySet1(_807_v49, (_index99).Int())
				var _808_v50 _dafny.Map
				_ = _808_v50
				_808_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, p2)
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_802_v44), 0))
				_ = _index100
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))
				_ = _index101
				var _rhs113 _dafny.Int = (_808_v50).Cardinality()
				_ = _rhs113
				var _rhs114 _dafny.MultiSet = _803_v45
				_ = _rhs114
				var _rhs115 _dafny.Sequence = _dafny.SeqOf(_805_v47)
				_ = _rhs115
				var _rhs116 _dafny.Array = _807_v49
				_ = _rhs116
				var _rhs117 _dafny.Int = (_this).F12()
				_ = _rhs117
				var _lhs82 _dafny.Array = _802_v44
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_802_v44), 0))
				_ = _lhs83
				var _lhs84 _dafny.Array = p1
				_ = _lhs84
				var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				r1 = _rhs113
				(_lhs82).ArraySet1(_rhs114, (_lhs83).Int())
				_806_v48 = _rhs115
				(_lhs84).ArraySet1(_rhs116, (_lhs85).Int())
				_lhs86.F5 = _rhs117
				var _809_v51 _dafny.Array
				_ = _809_v51
				var _nwElement0_28 _dafny.Sequence = _774_v24
				_ = _nwElement0_28
				var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(23))
				_ = _nw148
				(_nw148).ArraySet1(_nwElement0_28, 0)
				(_nw148).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_774_v24, _774_v24), 1)
				(_nw148).ArraySet1((func() _dafny.Sequence {
					if _745_v0 {
						return _774_v24
					}
					return Companion_Default___.Fm15(_745_v0, _745_v0, globalState)
				})(), 2)
				(_nw148).ArraySet1(_774_v24, 3)
				(_nw148).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_774_v24, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-462))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_810_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_811_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg38 _dafny.Int) interface{} {
								return coer38(arg38)
							}
						}((func(_812_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_813_i4 _dafny.Int) _dafny.Int {
								return _812_p0
							}
						})(_810_p0)))).Cardinality())
					}
				})(p0)))), 4)
				(_nw148).ArraySet1(_dafny.SeqOf((_this).F12()), 5)
				(_nw148).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(56))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_814_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_815_i5 _dafny.Int) _dafny.Int {
						return _814_p3
					}
				})(p3))), 6)
				(_nw148).ArraySet1(_774_v24, 7)
				(_nw148).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(679))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_816_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_817_i6 _dafny.Int) _dafny.Int {
						return _816_p0
					}
				})(p0))), _774_v24), 8)
				(_nw148).ArraySet1(_dafny.SeqOf((_this).F12(), p3, Companion_Default___.Fm0(_dafny.IntOfInt64(499), (_805_v47).Cardinality(), p3, _745_v0, globalState)), 9)
				(_nw148).ArraySet1(_774_v24, 10)
				(_nw148).ArraySet1(_774_v24, 11)
				(_nw148).ArraySet1(_774_v24, 12)
				(_nw148).ArraySet1(_774_v24, 13)
				(_nw148).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(985))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_818_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_819_i7 _dafny.Int) _dafny.Int {
						return _818_p3
					}
				})(p3))), 14)
				(_nw148).ArraySet1(_774_v24, 15)
				(_nw148).ArraySet1(_774_v24, 16)
				(_nw148).ArraySet1(_dafny.SeqOf((_this).F12(), p2), 17)
				(_nw148).ArraySet1(_774_v24, 18)
				(_nw148).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_820_v50 _dafny.Map, _821_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_822_i8 _dafny.Int) _dafny.Int {
						return (func() _dafny.Int {
							if (_820_v50).Contains(false) {
								return (_820_v50).Get(false).(_dafny.Int)
							}
							return _821_p2
						})()
					}
				})(_808_v50, p2))), _774_v24), 19)
				(_nw148).ArraySet1(_774_v24, 20)
				(_nw148).ArraySet1(_774_v24, 21)
				(_nw148).ArraySet1(_774_v24, 22)
				_809_v51 = _nw148
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_809_v51), 0))
				_ = _index102
				(_809_v51).ArraySet1(_774_v24, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_809_v51), 0))
				_ = _index103
				(_809_v51).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(939))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_823_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_824_i9 _dafny.Int) _dafny.Int {
						return _823_p2
					}
				})(p2))), (_index103).Int())
				var _825_v52 _dafny.Array
				_ = _825_v52
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_27
				var _nw149 _dafny.Array
				_ = _nw149
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw149 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = (func(_826_v0 bool) func(_dafny.Int) bool {
						return func(_827_i10 _dafny.Int) bool {
							return _826_v0
						}
					})(_745_v0)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw149 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw149).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw149).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_825_v52 = _nw149
				_825_v52 = _825_v52
				var _828_v53 _dafny.Map
				_ = _828_v53
				_828_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int())), _745_v0)
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_807_v49), 0))
				_ = _index104
				(_807_v49).ArraySet1(Companion_Default___.Fm0(p2, p3, p0, (func() bool {
					if (_828_v53).Contains(_807_v49) {
						return (_828_v53).Get(_807_v49).(bool)
					}
					return _745_v0
				})(), globalState), (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_807_v49), 0))
				_ = _index105
				(_807_v49).ArraySet1((_this).F12(), (_index105).Int())
			}
			var _829_v54 D0
			_ = _829_v54
			_829_v54 = Companion_D0_.Create_DC0_(_776_v26)
			var _830_v55 D0
			_ = _830_v55
			_830_v55 = Companion_D0_.Create_DC2_(_829_v54)
			var _831_v56 _dafny.MultiSet
			_ = _831_v56
			_831_v56 = _dafny.MultiSetOf(_830_v55)
			var _832_v57 _dafny.Sequence
			_ = _832_v57
			_832_v57 = _dafny.UnicodeSeqOfUtf8Bytes("sfcvgsxqa")
			var _833_v58 _dafny.Sequence
			_ = _833_v58
			_833_v58 = _dafny.SeqOf(_775_v25, _dafny.SeqOf(_745_v0, _745_v0), _779_v29, _779_v29, _779_v29)
			var _834_v59 D2
			_ = _834_v59
			_834_v59 = Companion_D2_.Create_DC6_(_745_v0)
			var _835_v60 _dafny.Array
			_ = _835_v60
			var _nwElement0_29 bool = (_831_v56).IsDisjointFrom(_831_v56)
			_ = _nwElement0_29
			var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(18))
			_ = _nw150
			(_nw150).ArraySet1(_nwElement0_29, 0)
			(_nw150).ArraySet1(_745_v0, 1)
			(_nw150).ArraySet1(_745_v0, 2)
			(_nw150).ArraySet1(_745_v0, 3)
			(_nw150).ArraySet1(_dafny.Companion_Sequence_.Equal(_774_v24, _dafny.SeqOf(_dafny.IntOfUint32((_779_v29).Cardinality()), p3, _dafny.IntOfUint32((_832_v57).Cardinality()))), 4)
			(_nw150).ArraySet1(_745_v0, 5)
			(_nw150).ArraySet1(!(_745_v0), 6)
			(_nw150).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_833_v58, _dafny.SeqOf(_775_v25, _779_v29)), 7)
			(_nw150).ArraySet1(false, 8)
			(_nw150).ArraySet1(_745_v0, 9)
			(_nw150).ArraySet1(_745_v0, 10)
			(_nw150).ArraySet1(_745_v0, 11)
			(_nw150).ArraySet1(_745_v0, 12)
			(_nw150).ArraySet1(_745_v0, 13)
			(_nw150).ArraySet1(_745_v0, 14)
			(_nw150).ArraySet1((_834_v59).Dtor_cf6(), 15)
			(_nw150).ArraySet1(_745_v0, 16)
			(_nw150).ArraySet1(_745_v0, 17)
			_835_v60 = _nw150
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_835_v60), 0))
			_ = _index106
			(_835_v60).ArraySet1((_745_v0) || (_745_v0), (_index106).Int())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_835_v60), 0))
			_ = _index107
			var _rhs118 bool = _745_v0
			_ = _rhs118
			var _rhs119 _dafny.Int = p3
			_ = _rhs119
			var _rhs120 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((p2).Minus((p0).Times((_this).F12()))))
			_ = _rhs120
			var _rhs121 _dafny.Int = (func() _dafny.Int {
				if (_834_v59).Dtor_cf6() {
					return p0
				}
				return p2
			})()
			_ = _rhs121
			var _lhs87 _dafny.Array = _835_v60
			_ = _lhs87
			var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_835_v60), 0))
			_ = _lhs88
			var _lhs89 *GlobalState = globalState
			_ = _lhs89
			var _lhs90 *GlobalState = globalState
			_ = _lhs90
			var _lhs91 *GlobalState = globalState
			_ = _lhs91
			(_lhs87).ArraySet1(_rhs118, (_lhs88).Int())
			_lhs89.F5 = _rhs119
			_lhs90.F5 = _rhs120
			_lhs91.F5 = _rhs121
			var _836_v61 _dafny.Map
			_ = _836_v61
			_836_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_835_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_835_v60), 0))).Int()).(bool), p0)
			var _837_v62 _dafny.Set
			_ = _837_v62
			_837_v62 = _dafny.SetOf((_771_v21).Cardinality())
			var _838_v63 _dafny.Map
			_ = _838_v63
			_838_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_837_v62).Cardinality()), _dafny.IntOfInt64(-557))
			var _839_v64 _dafny.Array
			_ = _839_v64
			var _nwElement0_30 _dafny.Int = ((_this).F12()).Times(p2)
			_ = _nwElement0_30
			var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(13))
			_ = _nw151
			(_nw151).ArraySet1(_nwElement0_30, 0)
			(_nw151).ArraySet1((_this).F12(), 1)
			(_nw151).ArraySet1((_773_v23).Cardinality(), 2)
			(_nw151).ArraySet1(p0, 3)
			(_nw151).ArraySet1(p0, 4)
			(_nw151).ArraySet1(p2, 5)
			(_nw151).ArraySet1((func() _dafny.Int {
				if (_836_v61).Contains(_745_v0) {
					return (_836_v61).Get(_745_v0).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_838_v63).Contains((_776_v26).Cardinality()) {
						return (_838_v63).Get((_776_v26).Cardinality()).(_dafny.Int)
					}
					return p0
				})()
			})(), 6)
			(_nw151).ArraySet1(p0, 7)
			(_nw151).ArraySet1(_dafny.IntOfInt64(687), 8)
			(_nw151).ArraySet1(_dafny.IntOfInt64(448), 9)
			(_nw151).ArraySet1(p3, 10)
			(_nw151).ArraySet1(_dafny.IntOfInt64(952), 11)
			(_nw151).ArraySet1((_dafny.Zero).Minus(p2), 12)
			_839_v64 = _nw151
			var _rhs122 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, p0)
			_ = _rhs122
			var _rhs123 _dafny.Int = (func() _dafny.Int {
				if (_835_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_835_v60), 0))).Int()).(bool) {
					return _dafny.IntOfInt64(807)
				}
				return (p3).Plus((_dafny.MultiSetOf((_835_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_835_v60), 0))).Int()).(bool))).Cardinality())
			})()
			_ = _rhs123
			var _rhs124 _dafny.Array = _839_v64
			_ = _rhs124
			var _rhs125 bool = !(_745_v0)
			_ = _rhs125
			var _lhs92 *GlobalState = globalState
			_ = _lhs92
			var _lhs93 *GlobalState = globalState
			_ = _lhs93
			_lhs92.F5 = _rhs122
			_lhs93.F5 = _rhs123
			_839_v64 = _rhs124
			_745_v0 = _rhs125
		}
		var _840_v65 _dafny.CodePoint
		_ = _840_v65
		_840_v65 = _dafny.CodePoint('u')
		_840_v65 = Companion_Default___.Fm16(globalState)
		var _841_v66 _dafny.MultiSet
		_ = _841_v66
		_841_v66 = _dafny.MultiSetOf(true, false)
		r1 = ((_this).F12()).Minus((_841_v66).Cardinality())
		var _842_i11 _dafny.Int
		_ = _842_i11
		_842_i11 = _dafny.Zero
		{
			for true {
				{
					if (_842_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_842_i11 = (_842_i11).Plus(_dafny.One)
					if false {
						var _843_v67 _dafny.Map
						_ = _843_v67
						_843_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _745_v0)
						var _844_v68 _dafny.Sequence
						_ = _844_v68
						_844_v68 = _dafny.SeqOf(false)
						var _845_v69 _dafny.Sequence
						_ = _845_v69
						_845_v69 = _dafny.UnicodeSeqOfUtf8Bytes("kdak")
						var _846_v70 _dafny.Map
						_ = _846_v70
						_846_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _845_v69), false)
						var _847_v72 _dafny.Array
						_ = _847_v72
						var _nwElement0_31 bool = true
						_ = _nwElement0_31
						var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(28))
						_ = _nw152
						(_nw152).ArraySet1(_nwElement0_31, 0)
						(_nw152).ArraySet1(_745_v0, 1)
						(_nw152).ArraySet1(false, 2)
						(_nw152).ArraySet1(_745_v0, 3)
						(_nw152).ArraySet1(_745_v0, 4)
						(_nw152).ArraySet1(_745_v0, 5)
						(_nw152).ArraySet1((func() bool {
							if (_843_v67).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_844_v68, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_844_v68).Cardinality()))).Uint32(), _745_v0)).Cardinality())) {
								return (_843_v67).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_844_v68, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_844_v68).Cardinality()))).Uint32(), _745_v0)).Cardinality())).(bool)
							}
							return _745_v0
						})(), 6)
						(_nw152).ArraySet1(_745_v0, 7)
						(_nw152).ArraySet1(Companion_Default___.Fm1(_745_v0, globalState), 8)
						(_nw152).ArraySet1(_745_v0, 9)
						(_nw152).ArraySet1(_745_v0, 10)
						(_nw152).ArraySet1(_745_v0, 11)
						(_nw152).ArraySet1(_745_v0, 12)
						(_nw152).ArraySet1(_745_v0, 13)
						(_nw152).ArraySet1(true, 14)
						(_nw152).ArraySet1(_745_v0, 15)
						(_nw152).ArraySet1(_745_v0, 16)
						(_nw152).ArraySet1(_745_v0, 17)
						(_nw152).ArraySet1(_745_v0, 18)
						(_nw152).ArraySet1(_745_v0, 19)
						(_nw152).ArraySet1(_745_v0, 20)
						(_nw152).ArraySet1(!(_745_v0), 21)
						(_nw152).ArraySet1(_745_v0, 22)
						(_nw152).ArraySet1(_745_v0, 23)
						(_nw152).ArraySet1(_745_v0, 24)
						(_nw152).ArraySet1(Companion_Default___.Fm1(_745_v0, globalState), 25)
						(_nw152).ArraySet1((func() bool {
							if (_846_v70).Contains(func() _dafny.Map {
								var _coll40 = _dafny.NewMapBuilder()
								_ = _coll40
								for _iter42 := _dafny.Iterate((_843_v67).Keys().Elements()); ; {
									_compr_40, _ok42 := _iter42()
									if !_ok42 {
										break
									}
									var _848_v71 _dafny.Int
									_848_v71 = interface{}(_compr_40).(_dafny.Int)
									if (_843_v67).Contains(_848_v71) {
										_coll40.Add((_848_v71).Times(p0), _845_v69)
									}
								}
								return _coll40.ToMap()
							}()) {
								return (_846_v70).Get(func() _dafny.Map {
									var _coll41 = _dafny.NewMapBuilder()
									_ = _coll41
									for _iter43 := _dafny.Iterate((_843_v67).Keys().Elements()); ; {
										_compr_41, _ok43 := _iter43()
										if !_ok43 {
											break
										}
										var _849_v71 _dafny.Int
										_849_v71 = interface{}(_compr_41).(_dafny.Int)
										if (_843_v67).Contains(_849_v71) {
											_coll41.Add((_849_v71).Times(p0), _845_v69)
										}
									}
									return _coll41.ToMap()
								}()).(bool)
							}
							return _745_v0
						})(), 26)
						(_nw152).ArraySet1(_745_v0, 27)
						_847_v72 = _nw152
						var _850_v73 *C0
						_ = _850_v73
						var _nw153 *C0 = New_C0_()
						_ = _nw153
						_nw153.Ctor__(true, _847_v72)
						_850_v73 = _nw153
						var _851_v75 D0
						_ = _851_v75
						_851_v75 = Companion_D0_.Create_DC0_(func() _dafny.Map {
							var _coll42 = _dafny.NewMapBuilder()
							_ = _coll42
							for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(725), _dafny.IntOfInt64(-252))); ; {
								_compr_42, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _852_v74 _dafny.Int
								_852_v74 = interface{}(_compr_42).(_dafny.Int)
								if ((_dafny.IntOfInt64(725)).Cmp(_852_v74) <= 0) && ((_852_v74).Cmp(_dafny.IntOfInt64(-252)) < 0) {
									_coll42.Add(Companion_Default___.SafeDivisionInt(_852_v74, p3), (_850_v73).F10())
								}
							}
							return _coll42.ToMap()
						}())
						var _pat_let_tv24 = _843_v67
						_ = _pat_let_tv24
						_847_v72 = (func() _dafny.Array {
							if (_this).Fm4(func(_pat_let12_0 D0) D0 {
								return func(_853_dt__update__tmp_h3 D0) D0 {
									return func(_pat_let13_0 _dafny.Map) D0 {
										return func(_854_dt__update_hcf0_h0 _dafny.Map) D0 {
											return Companion_D0_.Create_DC0_(_854_dt__update_hcf0_h0)
										}(_pat_let13_0)
									}(_pat_let_tv24)
								}(_pat_let12_0)
							}(_851_v75), Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _745_v0)), _dafny.MultiSetFromSeq(_844_v68), globalState) {
								return (_850_v73).F11()
							}
							return _847_v72
						})()
						var _855_v76 _dafny.Set
						_ = _855_v76
						_855_v76 = _dafny.SetOf((_850_v73).F10())
						(globalState).F5 = ((_855_v76).Cardinality()).Times(p3)
						r0 = ((_855_v76).Difference(_855_v76)).IsSubsetOf(_855_v76)
						var _856_v77 *C0
						_ = _856_v77
						var _nw154 *C0 = New_C0_()
						_ = _nw154
						_nw154.Ctor__(_745_v0, (_850_v73).F11())
						_856_v77 = _nw154
					} else {
						var _857_v78 _dafny.Array
						_ = _857_v78
						var _len0_28 _dafny.Int = _dafny.IntOfInt64(4)
						_ = _len0_28
						var _nw155 _dafny.Array
						_ = _nw155
						if _len0_28.Cmp(_dafny.Zero) == 0 {
							_nw155 = _dafny.NewArray(_len0_28)
						} else {
							var _init28 func(_dafny.Int) bool = (func(_858_v0 bool) func(_dafny.Int) bool {
								return func(_859_i12 _dafny.Int) bool {
									return _858_v0
								}
							})(_745_v0)
							_ = _init28
							var _element0_28 = _init28(_dafny.Zero)
							_ = _element0_28
							_nw155 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
							(_nw155).ArraySet1(_element0_28, 0)
							var _nativeLen0_28 = (_len0_28).Int()
							_ = _nativeLen0_28
							for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
								(_nw155).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
							}
						}
						_857_v78 = _nw155
						var _860_v79 *C0
						_ = _860_v79
						var _nw156 *C0 = New_C0_()
						_ = _nw156
						_nw156.Ctor__(_745_v0, _857_v78)
						_860_v79 = _nw156
						var _861_v80 _dafny.Array
						_ = _861_v80
						var _nwElement0_32 _dafny.Int = p2
						_ = _nwElement0_32
						var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(4))
						_ = _nw157
						(_nw157).ArraySet1(_nwElement0_32, 0)
						(_nw157).ArraySet1(p0, 1)
						(_nw157).ArraySet1(p0, 2)
						(_nw157).ArraySet1(p0, 3)
						_861_v80 = _nw157
						var _862_v81 _dafny.Map
						_ = _862_v81
						_862_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _861_v80)
						var _863_v82 _dafny.Sequence
						_ = _863_v82
						_863_v82 = _dafny.UnicodeSeqOfUtf8Bytes("dypeuvr")
						var _864_v83 _dafny.Set
						_ = _864_v83
						_864_v83 = _dafny.SetOf((_863_v82).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("etmifa")).Cardinality()), _dafny.IntOfUint32((_863_v82).Cardinality()))).Uint32()).(_dafny.CodePoint))
						var _865_v84 _dafny.Map
						_ = _865_v84
						_865_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_863_v82, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_863_v82).Cardinality()))).Uint32(), _840_v65)).Cardinality()))
						var _866_v85 _dafny.Array
						_ = _866_v85
						var _nwElement0_33 _dafny.Int = _dafny.IntOfInt64(-539)
						_ = _nwElement0_33
						var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(23))
						_ = _nw158
						(_nw158).ArraySet1(_nwElement0_33, 0)
						(_nw158).ArraySet1(((_862_v81).Cardinality()).Times((_dafny.Zero).Minus(p0)), 1)
						(_nw158).ArraySet1(p2, 2)
						(_nw158).ArraySet1(p0, 3)
						(_nw158).ArraySet1(p0, 4)
						(_nw158).ArraySet1((func() _dafny.Int {
							if _745_v0 {
								return (_this).F12()
							}
							return p0
						})(), 5)
						(_nw158).ArraySet1(Companion_Default___.Fm0(Companion_Default___.Fm0(p0, p0, (Companion_Default___.Fm17(true, p3, _745_v0, p0, globalState)).Cardinality(), _745_v0, globalState), (_this).F12(), (_864_v83).Cardinality(), _745_v0, globalState), 6)
						(_nw158).ArraySet1((func() _dafny.Int {
							if _745_v0 {
								return (_865_v84).Cardinality()
							}
							return (_dafny.Zero).Minus((_this).F12())
						})(), 7)
						(_nw158).ArraySet1((_this).F12(), 8)
						(_nw158).ArraySet1((_this).F12(), 9)
						(_nw158).ArraySet1(p3, 10)
						(_nw158).ArraySet1(p0, 11)
						(_nw158).ArraySet1(p3, 12)
						(_nw158).ArraySet1((p2).Times(p2), 13)
						(_nw158).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_860_v79).F10() {
								return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfInt64(568)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(568))).Cardinality()))).Uint32(), _dafny.IntOfUint32((_863_v82).Cardinality()))
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg44 _dafny.Int) interface{} {
									return coer44(arg44)
								}
							}((func(_867_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_868_i13 _dafny.Int) _dafny.Int {
									return _867_p2
								}
							})(p2)))
						})()).Cardinality()), 14)
						(_nw158).ArraySet1(p0, 15)
						(_nw158).ArraySet1((_841_v66).Cardinality(), 16)
						(_nw158).ArraySet1(p3, 17)
						(_nw158).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_863_v82).Cardinality()), (_this).F12()), 18)
						(_nw158).ArraySet1(Companion_Default___.SafeDivisionInt(p2, p2), 19)
						(_nw158).ArraySet1((_dafny.MultiSetOf((_this).F12())).Cardinality(), 20)
						(_nw158).ArraySet1(Companion_Default___.SafeModuloInt((_this).F12(), (_this).F12()), 21)
						(_nw158).ArraySet1(_dafny.IntOfInt64(-487), 22)
						_866_v85 = _nw158
						_866_v85 = _866_v85
						(globalState).F5 = (_this).F12()
						(globalState).F0 = !(!(!((_this).Fm12(_863_v82, (_this).F12(), p0, globalState))) || ((_860_v79).F10()))
						(globalState).F0 = _745_v0
					}
					if _745_v0 {
						var _869_v86 _dafny.Array
						_ = _869_v86
						var _nw159 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
						_ = _nw159
						_869_v86 = _nw159
						_869_v86 = _869_v86
						var _870_v87 _dafny.Sequence
						_ = _870_v87
						_870_v87 = _dafny.UnicodeSeqOfUtf8Bytes("jqxxg")
						var _871_v88 _dafny.Array
						_ = _871_v88
						var _nw160 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
						_ = _nw160
						_871_v88 = _nw160
						var _872_v89 _dafny.Int
						_ = _872_v89
						var _873_v90 _dafny.Array
						_ = _873_v90
						var _out4 _dafny.Int
						_ = _out4
						var _out5 _dafny.Array
						_ = _out5
						_out4, _out5 = (_this).M3(Companion_Default___.Fm14(_870_v87, p2, globalState), _745_v0, _871_v88, p0, globalState)
						_872_v89 = _out4
						_873_v90 = _out5
						_870_v87 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _870_v87), _870_v87)
						var _874_v91 _dafny.Sequence
						_ = _874_v91
						_874_v91 = _dafny.SeqOf(_dafny.IntOfUint32((_870_v87).Cardinality()))
						var _875_v92 _dafny.Map
						_ = _875_v92
						_875_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_874_v91, p2)
						_875_v92 = ((_875_v92).Merge(_875_v92)).Merge(_875_v92)
						var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_869_v86), 0))
						_ = _index108
						(_869_v86).ArraySet1((func() bool {
							if _745_v0 {
								return _745_v0
							}
							return true
						})(), (_index108).Int())
						var _876_v93 T0
						_ = _876_v93
						var _nw161 *C2 = New_C2_()
						_ = _nw161
						_nw161.Ctor__(false)
						_876_v93 = _nw161
						var _877_v94 _dafny.Map
						_ = _877_v94
						_877_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _876_v93)
						var _878_v95 _dafny.MultiSet
						_ = _878_v95
						_878_v95 = _dafny.MultiSetOf(_872_v89, _dafny.IntOfInt64(-22), _dafny.IntOfInt64(181), (_877_v94).Cardinality(), _dafny.IntOfInt64(-526))
						var _879_v96 _dafny.Map
						_ = _879_v96
						_879_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(192))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}((func(_880_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_881_i14 _dafny.Int) _dafny.Int {
								return _880_p3
							}
						})(p3)))).Cardinality()), _745_v0)
						var _882_v97 D0
						_ = _882_v97
						_882_v97 = Companion_D0_.Create_DC0_(_879_v96)
						var _883_v98 D0
						_ = _883_v98
						_883_v98 = Companion_D0_.Create_DC2_(_882_v97)
						var _884_v99 D0
						_ = _884_v99
						_884_v99 = Companion_D0_.Create_DC0_(_879_v96)
						var _885_v100 D12
						_ = _885_v100
						_885_v100 = Companion_D12_.Create_DC34_(_883_v98, _745_v0, (_876_v93).Fm4(_884_v99, _884_v99, _841_v66, globalState))
						var _886_v101 D12
						_ = _886_v101
						_886_v101 = Companion_D12_.Create_DC34_(Companion_D0_.Create_DC2_(_882_v97), _745_v0, !((_885_v100).Dtor_cf50()))
						var _887_v102 _dafny.Sequence
						_ = _887_v102
						_887_v102 = _dafny.SeqOf(_745_v0, (_878_v95).IsDisjointFrom(_dafny.MultiSetFromSeq(_874_v91)), (_885_v100).Dtor_cf49(), true)
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_869_v86), 0))
						_ = _index109
						(_869_v86).ArraySet1((_887_v102).Select((Companion_Default___.SafeIndex(_872_v89, _dafny.IntOfUint32((_887_v102).Cardinality()))).Uint32()).(bool), (_index109).Int())
					} else {
						var _888_v103 _dafny.Map
						_ = _888_v103
						_888_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v0, _745_v0)
						var _889_v104 _dafny.Sequence
						_ = _889_v104
						_889_v104 = _dafny.SeqOf(_888_v103)
						var _890_v105 _dafny.Map
						_ = _890_v105
						_890_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_889_v104).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_889_v104).Cardinality()))).Uint32()).(_dafny.Map), _dafny.IntOfInt64(762))
						var _891_v106 D2
						_ = _891_v106
						_891_v106 = Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("heyc"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F12()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("heyc")).Cardinality()))).Uint32(), _840_v65)).Cardinality()), (_890_v105).Update(_888_v103, (_this).F12()), (_this).F12())
						(globalState).F5 = (_891_v106).Dtor_cf9()
						var _892_v107 _dafny.Sequence
						_ = _892_v107
						_892_v107 = _dafny.SeqOf(!(_745_v0))
						var _893_v108 D9
						_ = _893_v108
						_893_v108 = Companion_D9_.Create_DC24_(true, true, (_892_v107).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F12()), _dafny.IntOfUint32((_892_v107).Cardinality()))).Uint32()).(bool))
						var _894_v109 _dafny.Map
						_ = _894_v109
						_894_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _745_v0)
						var _895_v110 _dafny.Map
						_ = _895_v110
						_895_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_892_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.IntOfUint32((_892_v107).Cardinality()))).Uint32()).(bool), p0)
						var _rhs126 D9 = (func() D9 {
							if ((func() bool {
								if (_894_v109).Contains(_dafny.IntOfInt64(632)) {
									return (_894_v109).Get(_dafny.IntOfInt64(632)).(bool)
								}
								return (func() bool {
									if (_888_v103).Contains(false) {
										return (_888_v103).Get(false).(bool)
									}
									return _745_v0
								})()
							})()) == (_745_v0) {
								return Companion_D9_.Create_DC24_(false, _745_v0, _745_v0)
							}
							return _893_v108
						})()
						_ = _rhs126
						var _rhs127 _dafny.Int = Companion_Default___.SafeDivisionInt(((Companion_D13_.Create_DC38_(_895_v110)).Dtor_cf55()).Cardinality(), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_892_v107).Cardinality()), p3))
						_ = _rhs127
						_893_v108 = _rhs126
						r1 = _rhs127
						var _896_v111 _dafny.Map
						_ = _896_v111
						_896_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F12())
						var _897_v112 _dafny.Map
						_ = _897_v112
						_897_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_896_v111, _745_v0)
						(globalState).F0 = !(!(!((func() bool {
							if (_897_v112).Contains(_896_v111) {
								return (_897_v112).Get(_896_v111).(bool)
							}
							return _745_v0
						})())))
						var _898_v113 _dafny.Sequence
						_ = _898_v113
						_898_v113 = _dafny.SeqOf((_this).F12())
						var _899_v114 _dafny.Set
						_ = _899_v114
						_899_v114 = _dafny.SetOf(Companion_Default___.Fm0(_dafny.IntOfUint32((_898_v113).Cardinality()), p0, p0, _745_v0, globalState), p0)
						var _900_v116 _dafny.Array
						_ = _900_v116
						var _nwElement0_34 bool = _745_v0
						_ = _nwElement0_34
						var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(9))
						_ = _nw162
						(_nw162).ArraySet1(_nwElement0_34, 0)
						(_nw162).ArraySet1(true, 1)
						(_nw162).ArraySet1(_dafny.Companion_Sequence_.Equal(_898_v113, _898_v113), 2)
						(_nw162).ArraySet1(_745_v0, 3)
						(_nw162).ArraySet1(!(_745_v0), 4)
						(_nw162).ArraySet1(_745_v0, 5)
						(_nw162).ArraySet1(_745_v0, 6)
						(_nw162).ArraySet1(_745_v0, 7)
						(_nw162).ArraySet1((_899_v114).IsDisjointFrom(func() _dafny.Set {
							var _coll43 = _dafny.NewBuilder()
							_ = _coll43
							for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(750), _dafny.IntOfInt64(-18))); ; {
								_compr_43, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _901_v115 _dafny.Int
								_901_v115 = interface{}(_compr_43).(_dafny.Int)
								if ((_dafny.IntOfInt64(750)).Cmp(_901_v115) <= 0) && ((_901_v115).Cmp(_dafny.IntOfInt64(-18)) < 0) {
									_coll43.Add(Companion_Default___.SafeModuloInt(_901_v115, p0))
								}
							}
							return _coll43.ToSet()
						}()), 8)
						_900_v116 = _nw162
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_900_v116), 0))
						_ = _index110
						(_900_v116).ArraySet1(_745_v0, (_index110).Int())
						var _902_v117 _dafny.Array
						_ = _902_v117
						var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
						_ = _nw163
						_902_v117 = _nw163
						var _903_v118 _dafny.Array
						_ = _903_v118
						var _len0_29 _dafny.Int = _dafny.IntOfInt64(11)
						_ = _len0_29
						var _nw164 _dafny.Array
						_ = _nw164
						if _len0_29.Cmp(_dafny.Zero) == 0 {
							_nw164 = _dafny.NewArray(_len0_29)
						} else {
							var _init29 func(_dafny.Int) _dafny.Int = (func(_904_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_905_i15 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeModuloInt(_905_i15, _904_p3)
								}
							})(p3)
							_ = _init29
							var _element0_29 = _init29(_dafny.Zero)
							_ = _element0_29
							_nw164 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
							(_nw164).ArraySet1(_element0_29, 0)
							var _nativeLen0_29 = (_len0_29).Int()
							_ = _nativeLen0_29
							for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
								(_nw164).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
							}
						}
						_903_v118 = _nw164
						var _906_v119 _dafny.Array
						_ = _906_v119
						_906_v119 = _903_v118
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_902_v117), 0))
						_ = _index111
						(_902_v117).ArraySet1(_906_v119, (_index111).Int())
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_900_v116), 0))
						_ = _index112
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_902_v117), 0))
						_ = _index113
						var _rhs128 bool = !((_745_v0) == (!(_745_v0))) || (_745_v0)
						_ = _rhs128
						var _rhs129 bool = false
						_ = _rhs129
						var _rhs130 _dafny.Array = _906_v119
						_ = _rhs130
						var _lhs94 _dafny.Array = _900_v116
						_ = _lhs94
						var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_900_v116), 0))
						_ = _lhs95
						var _lhs96 _dafny.Array = _902_v117
						_ = _lhs96
						var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_902_v117), 0))
						_ = _lhs97
						r2 = _rhs128
						(_lhs94).ArraySet1(_rhs129, (_lhs95).Int())
						(_lhs96).ArraySet1(_rhs130, (_lhs97).Int())
						var _907_v120 D3
						_ = _907_v120
						_907_v120 = Companion_D3_.Create_DC9_(_dafny.IntOfInt64(409))
						var _908_v121 D3
						_ = _908_v121
						_908_v121 = Companion_D3_.Create_DC10_(_907_v120)
						var _909_v122 _dafny.Sequence
						_ = _909_v122
						_909_v122 = _dafny.UnicodeSeqOfUtf8Bytes("fxm")
						var _pat_let_tv25 = _909_v122
						_ = _pat_let_tv25
						var _910_v123 _dafny.Map
						_ = _910_v123
						_910_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, func(_pat_let14_0 D3) D3 {
							return func(_911_dt__update__tmp_h4 D3) D3 {
								return func(_pat_let15_0 D3) D3 {
									return func(_912_dt__update_hcf12_h0 D3) D3 {
										return Companion_D3_.Create_DC10_(_912_dt__update_hcf12_h0)
									}(_pat_let15_0)
								}(Companion_D3_.Create_DC8_(_pat_let_tv25))
							}(_pat_let14_0)
						}(_908_v121))
						(globalState).F0 = ((_910_v123).Cardinality()).Cmp(p0) > 0
					}
					var _913_v124 D11
					_ = _913_v124
					_913_v124 = Companion_D11_.Create_DC30_()
					var _914_v125 D11
					_ = _914_v125
					_914_v125 = Companion_D11_.Create_DC32_(_913_v124)
					var _915_v126 D11
					_ = _915_v126
					_915_v126 = Companion_D11_.Create_DC32_((_914_v125).Dtor_cf46())
					var _916_v128 _dafny.Array
					_ = _916_v128
					var _len0_30 _dafny.Int = _dafny.IntOfInt64(12)
					_ = _len0_30
					var _nw165 _dafny.Array
					_ = _nw165
					if _len0_30.Cmp(_dafny.Zero) == 0 {
						_nw165 = _dafny.NewArray(_len0_30)
					} else {
						var _init30 func(_dafny.Int) D4 = (func(_917_v66 _dafny.MultiSet, _918_p2 _dafny.Int) func(_dafny.Int) D4 {
							return func(_919_i16 _dafny.Int) D4 {
								return Companion_D4_.Create_DC12_((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F12(), (func() _dafny.Set {
									var _coll44 = _dafny.NewBuilder()
									_ = _coll44
									for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(349), _dafny.IntOfInt64(538))); ; {
										_compr_44, _ok46 := _iter46()
										if !_ok46 {
											break
										}
										var _920_v127 _dafny.Int
										_920_v127 = interface{}(_compr_44).(_dafny.Int)
										if ((_dafny.IntOfInt64(349)).Cmp(_920_v127) <= 0) && ((_920_v127).Cmp(_dafny.IntOfInt64(538)) < 0) {
											_coll44.Add(Companion_Default___.SafeDivisionInt(_920_v127, _918_p2))
										}
									}
									return _coll44.ToSet()
								}()).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_917_v66)).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("qxnwwoae"), (_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality()), _918_p2)
							}
						})(_841_v66, p2)
						_ = _init30
						var _element0_30 = _init30(_dafny.Zero)
						_ = _element0_30
						_nw165 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
						(_nw165).ArraySet1(_element0_30, 0)
						var _nativeLen0_30 = (_len0_30).Int()
						_ = _nativeLen0_30
						for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
							(_nw165).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
						}
					}
					_916_v128 = _nw165
					var _921_v129 _dafny.Map
					_ = _921_v129
					_921_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D11 {
						if _745_v0 {
							return _915_v126
						}
						return _915_v126
					})(), _916_v128)
					_921_v129 = (_921_v129).Update(_915_v126, _916_v128)
					var _922_v130 D11
					_ = _922_v130
					_922_v130 = Companion_D11_.Create_DC30_()
					var _923_v131 _dafny.Sequence
					_ = _923_v131
					_923_v131 = _dafny.SeqOf(Companion_D11_.Create_DC30_(), _922_v130, _922_v130, _922_v130)
					var _924_v132 _dafny.MultiSet
					_ = _924_v132
					_924_v132 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0))
					var _925_v133 _dafny.Sequence
					_ = _925_v133
					_925_v133 = _dafny.SeqOf(_dafny.SeqOf(p0, (_924_v132).Cardinality()))
					var _926_v134 _dafny.Sequence
					_ = _926_v134
					_926_v134 = _dafny.UnicodeSeqOfUtf8Bytes("sxyqh")
					var _927_v135 _dafny.Array
					_ = _927_v135
					var _nwElement0_35 _dafny.Int = p0
					_ = _nwElement0_35
					var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(19))
					_ = _nw166
					(_nw166).ArraySet1(_nwElement0_35, 0)
					(_nw166).ArraySet1(p2, 1)
					(_nw166).ArraySet1(((_this).F12()).Times(p3), 2)
					(_nw166).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(73))).Uint32(), func(coer46 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}(func(_928_i17 _dafny.Int) D11 {
						return Companion_D11_.Create_DC30_()
					})), _923_v131)).Cardinality())), 3)
					(_nw166).ArraySet1(p2, 4)
					(_nw166).ArraySet1(p3, 5)
					(_nw166).ArraySet1(p2, 6)
					(_nw166).ArraySet1(p3, 7)
					(_nw166).ArraySet1(((_this).F12()).Plus(p3), 8)
					(_nw166).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_925_v133).Cardinality()), p3), 9)
					(_nw166).ArraySet1((_this).F12(), 10)
					(_nw166).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p3), _dafny.IntOfUint32((_dafny.SeqOf(_745_v0, !(_745_v0), !(_745_v0), _745_v0, _745_v0)).Cardinality())), 11)
					(_nw166).ArraySet1(_dafny.IntOfInt64(-586), 12)
					(_nw166).ArraySet1((_this).F12(), 13)
					(_nw166).ArraySet1(p2, 14)
					(_nw166).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_this).F12()), 15)
					(_nw166).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_926_v134, _926_v134)).Cardinality()), 16)
					(_nw166).ArraySet1(p0, 17)
					(_nw166).ArraySet1(p2, 18)
					_927_v135 = _nw166
					var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_927_v135), 0))
					_ = _index114
					(_927_v135).ArraySet1((p2).Plus(_dafny.IntOfUint32((_926_v134).Cardinality())), (_index114).Int())
					var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_927_v135), 0))
					_ = _index115
					(_927_v135).ArraySet1(p2, (_index115).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _929_v136 _dafny.Array
		_ = _929_v136
		var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw167
		_929_v136 = _nw167
		for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_929_v136), 0))); ; {
			_guard_loop_2, _ok47 := _iter47()
			if !_ok47 {
				break
			}
			var _930_i18 _dafny.Int
			_930_i18 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_930_i18).Sign() != -1) && ((_930_i18).Cmp(_dafny.ArrayLen((_929_v136), 0)) < 0)) {
				(_929_v136).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_745_v0, _745_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_745_v0, _745_v0, _745_v0), _dafny.SeqOf(false))), (_930_i18).Int())
			}
		}
		var _931_v137 _dafny.Sequence
		_ = _931_v137
		_931_v137 = _dafny.UnicodeSeqOfUtf8Bytes("r")
		_931_v137 = _dafny.Companion_Sequence_.Concatenate(_931_v137, _931_v137)
		r0 = _745_v0
		r1 = p3
		var _932_v138 D11
		_ = _932_v138
		_932_v138 = Companion_D11_.Create_DC31_((_this).F12(), (_dafny.Zero).Minus(p2))
		var _pat_let_tv26 = _745_v0
		_ = _pat_let_tv26
		var _pat_let_tv27 = _745_v0
		_ = _pat_let_tv27
		var _pat_let_tv28 = _745_v0
		_ = _pat_let_tv28
		var _pat_let_tv29 = _745_v0
		_ = _pat_let_tv29
		r2 = func(_source12 D11) bool {
			if _source12.Is_DC30() {
				return _pat_let_tv26
			} else if _source12.Is_DC31() {
				var _933___mcc_h3 _dafny.Int = _source12.Get_().(D11_DC31).Cf44
				_ = _933___mcc_h3
				var _934___mcc_h4 _dafny.Int = _source12.Get_().(D11_DC31).Cf45
				_ = _934___mcc_h4
				var _935_cf45 _dafny.Int = _934___mcc_h4
				_ = _935_cf45
				var _936_cf44 _dafny.Int = _933___mcc_h3
				_ = _936_cf44
				return _pat_let_tv27
			} else if _source12.Is_DC29() {
				var _937___mcc_h5 T0 = _source12.Get_().(D11_DC29).Cf43
				_ = _937___mcc_h5
				var _938_cf43 T0 = _937___mcc_h5
				_ = _938_cf43
				return _pat_let_tv28
			} else {
				var _939___mcc_h6 D11 = _source12.Get_().(D11_DC32).Cf46
				_ = _939___mcc_h6
				var _940_cf46 D11 = _939___mcc_h6
				_ = _940_cf46
				return _pat_let_tv29
			}
		}(_932_v138)
		r3 = !(!(_745_v0) || (_745_v0))
		return r0, r1, r2, r3
	}
}
func (_this *C3) M3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		r0 = p3
		(globalState).F2 = !(!(p1))
		var _941_v0 D0
		_ = _941_v0
		_941_v0 = Companion_D0_.Create_DC1_(p1)
		var _942_v1 D0
		_ = _942_v1
		_942_v1 = Companion_D0_.Create_DC2_(_941_v0)
		var _943_v2 D12
		_ = _943_v2
		_943_v2 = Companion_D12_.Create_DC34_(_942_v1, p1, true)
		var _944_i0 _dafny.Int
		_ = _944_i0
		_944_i0 = _dafny.Zero
		{
			for (_943_v2).Dtor_cf50() {
				{
					if (_944_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_944_i0 = (_944_i0).Plus(_dafny.One)
					var _945_v3 *C2
					_ = _945_v3
					var _nw168 *C2 = New_C2_()
					_ = _nw168
					_nw168.Ctor__(p1)
					_945_v3 = _nw168
					var _946_v4 _dafny.Map
					_ = _946_v4
					_946_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _945_v3)
					var _947_v7 D0
					_ = _947_v7
					_947_v7 = Companion_D0_.Create_DC0_(func() _dafny.Map {
						var _coll45 = _dafny.NewMapBuilder()
						_ = _coll45
						for _iter48 := _dafny.Iterate((func() _dafny.Map {
							var _coll46 = _dafny.NewMapBuilder()
							_ = _coll46
							for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(695), _dafny.IntOfInt64(246))); ; {
								_compr_46, _ok49 := _iter49()
								if !_ok49 {
									break
								}
								var _948_v6 _dafny.Int
								_948_v6 = interface{}(_compr_46).(_dafny.Int)
								if ((_dafny.IntOfInt64(695)).Cmp(_948_v6) <= 0) && ((_948_v6).Cmp(_dafny.IntOfInt64(246)) < 0) {
									_coll46.Add(Companion_Default___.SafeModuloInt(_948_v6, (_dafny.Zero).Minus(p3)), (_945_v3).F13())
								}
							}
							return _coll46.ToMap()
						}()).Keys().Elements()); ; {
							_compr_45, _ok48 := _iter48()
							if !_ok48 {
								break
							}
							var _949_v5 _dafny.Int
							_949_v5 = interface{}(_compr_45).(_dafny.Int)
							if (func() _dafny.Map {
								var _coll47 = _dafny.NewMapBuilder()
								_ = _coll47
								for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(695), _dafny.IntOfInt64(246))); ; {
									_compr_47, _ok50 := _iter50()
									if !_ok50 {
										break
									}
									var _948_v6 _dafny.Int
									_948_v6 = interface{}(_compr_47).(_dafny.Int)
									if ((_dafny.IntOfInt64(695)).Cmp(_948_v6) <= 0) && ((_948_v6).Cmp(_dafny.IntOfInt64(246)) < 0) {
										_coll47.Add(Companion_Default___.SafeModuloInt(_948_v6, (_dafny.Zero).Minus(p3)), (_945_v3).F13())
									}
								}
								return _coll47.ToMap()
							}()).Contains(_949_v5) {
								_coll45.Add(Companion_Default___.SafeDivisionInt(_949_v5, p3), (_945_v3).F13())
							}
						}
						return _coll45.ToMap()
					}())
					var _950_v8 _dafny.MultiSet
					_ = _950_v8
					_950_v8 = _dafny.MultiSetOf((_945_v3).F13())
					var _951_v9 _dafny.Set
					_ = _951_v9
					_951_v9 = _dafny.SetOf(p1)
					var _952_v10 _dafny.Array
					_ = _952_v10
					var _nwElement0_36 bool = !(_946_v4).Contains(p1)
					_ = _nwElement0_36
					var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(17))
					_ = _nw169
					(_nw169).ArraySet1(_nwElement0_36, 0)
					(_nw169).ArraySet1((_945_v3).F13(), 1)
					(_nw169).ArraySet1(true, 2)
					(_nw169).ArraySet1(p1, 3)
					(_nw169).ArraySet1(((_945_v3).F13()) && (p1), 4)
					(_nw169).ArraySet1(p1, 5)
					(_nw169).ArraySet1((p3).Cmp((_this).F12()) >= 0, 6)
					(_nw169).ArraySet1(!(p1) || (!((_945_v3).F13())), 7)
					(_nw169).ArraySet1((_945_v3).F13(), 8)
					(_nw169).ArraySet1(((_945_v3).F13()) == (p1), 9)
					(_nw169).ArraySet1(Companion_Default___.Fm1((_945_v3).F13(), globalState), 10)
					(_nw169).ArraySet1((_945_v3).F13(), 11)
					(_nw169).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-224))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}(func(_953_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}(func(_954_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					}))), 12)
					(_nw169).ArraySet1((func() bool {
						if (_945_v3).F13() {
							return p1
						}
						return (_945_v3).F13()
					})(), 13)
					(_nw169).ArraySet1((_this).Fm4(_947_v7, _947_v7, _950_v8, globalState), 14)
					(_nw169).ArraySet1(((_this).F12()).Cmp((_951_v9).Cardinality()) == 0, 15)
					(_nw169).ArraySet1((_945_v3).F13(), 16)
					_952_v10 = _nw169
					var _955_v11 _dafny.Sequence
					_ = _955_v11
					_955_v11 = _dafny.UnicodeSeqOfUtf8Bytes("akjkn")
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_952_v10), 0))
					_ = _index116
					(_952_v10).ArraySet1((_945_v3).F13(), (_index116).Int())
					var _956_v12 _dafny.MultiSet
					_ = _956_v12
					_956_v12 = _dafny.MultiSetOf(_942_v1)
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_952_v10), 0))
					_ = _index117
					var _rhs131 _dafny.Array = _952_v10
					_ = _rhs131
					var _rhs132 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}((func(_957_p1 bool, _958_v3 *C2) func(_dafny.Int) _dafny.CodePoint {
						return func(_959_i3 _dafny.Int) _dafny.CodePoint {
							return (Companion_D10_.Create_DC27_(_957_p1, (_this).F12(), _dafny.CodePoint('n'), _dafny.MultiSetOf(false, (_958_v3).F13(), false, _957_p1))).Dtor_cf40()
						}
					})(p1, _945_v3)))
					_ = _rhs132
					var _rhs133 bool = (_956_v12).IsDisjointFrom((_956_v12).Intersection(Companion_Default___.Fm31(globalState)))
					_ = _rhs133
					var _lhs98 _dafny.Array = _952_v10
					_ = _lhs98
					var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_952_v10), 0))
					_ = _lhs99
					_952_v10 = _rhs131
					_955_v11 = _rhs132
					(_lhs98).ArraySet1(_rhs133, (_lhs99).Int())
					(globalState).F5 = (_this).F12()
					var _960_v13 _dafny.CodePoint
					_ = _960_v13
					_960_v13 = _dafny.CodePoint('s')
					_960_v13 = _dafny.CodePoint('f')
					var _961_v14 _dafny.Set
					_ = _961_v14
					_961_v14 = _dafny.SetOf((_this).F12(), (_this).F12())
					var _962_v15 D4
					_ = _962_v15
					_962_v15 = Companion_D4_.Create_DC13_(p3, _952_v10)
					var _963_v16 *C0
					_ = _963_v16
					var _nw170 *C0 = New_C0_()
					_ = _nw170
					_nw170.Ctor__((_961_v14).Equals(_961_v14), (_962_v15).Dtor_cf20())
					_963_v16 = _nw170
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _964_v17 _dafny.MultiSet
		_ = _964_v17
		_964_v17 = _dafny.MultiSetOf(p3, (_this).F12(), p3)
		var _965_v18 _dafny.Sequence
		_ = _965_v18
		_965_v18 = _dafny.SeqOf(p1, p1, p1, !(p1), Companion_Default___.Fm1(!(p1), globalState))
		var _hi4 _dafny.Int = (_this).F12()
		_ = _hi4
		for _966_i4 := ((_964_v17).Cardinality()).Plus(_dafny.IntOfUint32((_965_v18).Cardinality())); _966_i4.Cmp(_hi4) < 0; _966_i4 = _966_i4.Plus(_dafny.One) {
			_965_v18 = _965_v18
			var _967_v19 *C2
			_ = _967_v19
			var _nw171 *C2 = New_C2_()
			_ = _nw171
			_nw171.Ctor__(p1)
			_967_v19 = _nw171
			var _968_v20 _dafny.Array
			_ = _968_v20
			var _nw172 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw172
			_968_v20 = _nw172
			var _969_v21 *C0
			_ = _969_v21
			var _nw173 *C0 = New_C0_()
			_ = _nw173
			_nw173.Ctor__((_965_v18).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_965_v18).Cardinality()))).Uint32()).(bool), _968_v20)
			_969_v21 = _nw173
			var _970_v22 _dafny.Map
			_ = _970_v22
			_970_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), (_dafny.Zero).Minus(p3))
			_970_v22 = (_970_v22).Update((_967_v19).F13(), (_this).F12())
		}
		var _hi5 _dafny.Int = ((_this).F12()).Times(p3)
		_ = _hi5
		for _971_i5 := Companion_Default___.SafeDivisionInt((_964_v17).Cardinality(), p3); _971_i5.Cmp(_hi5) < 0; _971_i5 = _971_i5.Plus(_dafny.One) {
			(globalState).F5 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.UnicodeSeqOfUtf8Bytes("wjl"))).Cardinality())
			var _972_v23 D7
			_ = _972_v23
			_972_v23 = Companion_D7_.Create_DC18_()
			var _973_v24 D7
			_ = _973_v24
			_973_v24 = Companion_D7_.Create_DC19_(_972_v23)
			var _974_v25 D7
			_ = _974_v25
			_974_v25 = Companion_D7_.Create_DC19_(_973_v24)
			var _975_v26 D7
			_ = _975_v26
			_975_v26 = Companion_D7_.Create_DC19_((_974_v25).Dtor_cf24())
			var _976_v27 D7
			_ = _976_v27
			_976_v27 = Companion_D7_.Create_DC19_(_972_v23)
			var _977_v28 _dafny.Map
			_ = _977_v28
			_977_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_974_v25, true)
			_977_v28 = (_977_v28).Update(_976_v27, p1)
			var _978_v29 _dafny.Array
			_ = _978_v29
			var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw174
			_978_v29 = _nw174
			var _979_v30 *C0
			_ = _979_v30
			var _nw175 *C0 = New_C0_()
			_ = _nw175
			_nw175.Ctor__(!(p1) || (p1), _978_v29)
			_979_v30 = _nw175
			(globalState).F2 = !((_dafny.IntOfUint32((p0).Cardinality())).Cmp(_dafny.IntOfUint32((_965_v18).Cardinality())) >= 0)
		}
		var _980_v31 _dafny.Array
		_ = _980_v31
		var _nw176 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw176
		_980_v31 = _nw176
		_980_v31 = _980_v31
		var _981_v32 _dafny.Map
		_ = _981_v32
		_981_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
		var _982_v33 _dafny.Map
		_ = _982_v33
		_982_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_981_v32, (_this).F12())
		r0 = (Companion_D2_.Create_DC7_((_this).F12(), _982_v33, p3)).Dtor_cf9()
		var _983_v34 _dafny.Array
		_ = _983_v34
		var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw177
		_983_v34 = _nw177
		r1 = _983_v34
		return r0, r1
	}
}
func (_this *C3) F12() _dafny.Int {
	{
		return _this._f12
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f9 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f9 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f9 _dafny.Int) {
	{
		(_this)._f9 = f9
	}
}
func (_this *C4) Fm4(p0 D0, p1 D0, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeModuloInt((_this).F9(), (_this).F9())).Cmp(((_dafny.MultiSetOf(true)).Intersection(_dafny.MultiSetOf(false, true))).Cardinality()) > 0
	}
}
func (_this *C4) Fm5(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_dafny.SetOf((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("urcwm")).Cardinality())).Cmp((_this).F9()) == 0, _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("dwdsvbe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-918))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg50 _dafny.Int) interface{} {
				return coer50(arg50)
			}
		}(func(_984_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))), false)).Cardinality())
	}
}
func (_this *C4) M1(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		r1 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((p0).Times(p2)), p3)
		var _985_v0 _dafny.Array
		_ = _985_v0
		var _nw178 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw178
		_985_v0 = _nw178
		var _986_v1 *C0
		_ = _986_v1
		var _nw179 *C0 = New_C0_()
		_ = _nw179
		_nw179.Ctor__(true, _985_v0)
		_986_v1 = _nw179
		var _987_v2 *C0
		_ = _987_v2
		var _nw180 *C0 = New_C0_()
		_ = _nw180
		_nw180.Ctor__(((_986_v1).F10()) == ((_986_v1).F10()), _985_v0)
		_987_v2 = _nw180
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen(((_986_v1).F11()), 0))
		_ = _index118
		((_986_v1).F11()).ArraySet1((_986_v1).F10(), (_index118).Int())
		var _988_v3 _dafny.CodePoint
		_ = _988_v3
		_988_v3 = _dafny.CodePoint('m')
		var _989_v4 _dafny.Array
		_ = _989_v4
		var _len0_31 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_31
		var _nw181 _dafny.Array
		_ = _nw181
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw181 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) _dafny.Int = (func(_990_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_991_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_991_i0, _990_p0)
				}
			})(p0)
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw181 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw181).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw181).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_989_v4 = _nw181
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_989_v4), 0))
		_ = _index119
		(_989_v4).ArraySet1(p0, (_index119).Int())
		var _992_v5 _dafny.MultiSet
		_ = _992_v5
		_992_v5 = _dafny.MultiSetOf((_986_v1).F10())
		var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen(((_986_v1).F11()), 0))
		_ = _index120
		var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_989_v4), 0))
		_ = _index121
		var _rhs134 bool = !((_992_v5).Intersection(_dafny.MultiSetOf((_986_v1).F10(), (_986_v1).F10()))).Equals((_992_v5).Difference(_992_v5))
		_ = _rhs134
		var _rhs135 _dafny.CodePoint = _988_v3
		_ = _rhs135
		var _rhs136 _dafny.Int = Companion_Default___.SafeDivisionInt(p2, (_992_v5).Cardinality())
		_ = _rhs136
		var _lhs100 _dafny.Array = (_986_v1).F11()
		_ = _lhs100
		var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen(((_986_v1).F11()), 0))
		_ = _lhs101
		var _lhs102 _dafny.Array = _989_v4
		_ = _lhs102
		var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_989_v4), 0))
		_ = _lhs103
		(_lhs100).ArraySet1(_rhs134, (_lhs101).Int())
		_988_v3 = _rhs135
		(_lhs102).ArraySet1(_rhs136, (_lhs103).Int())
		var _993_v6 _dafny.Sequence
		_ = _993_v6
		_993_v6 = _dafny.SeqOf((_987_v2).F10(), (_986_v1).F10(), (_987_v2).F10())
		var _994_v7 _dafny.Map
		_ = _994_v7
		_994_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_993_v6).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_993_v6).Cardinality()))).Uint32()).(bool), (_989_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_989_v4), 0))).Int()).(_dafny.Int))
		_994_v7 = (_994_v7).Update(((_986_v1).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen(((_986_v1).F11()), 0))).Int()).(bool), (_989_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_989_v4), 0))).Int()).(_dafny.Int))
		r2 = ((_986_v1).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen(((_986_v1).F11()), 0))).Int()).(bool)
		r0 = (_986_v1).F10()
		r1 = p2
		r2 = Companion_Default___.Fm1((_987_v2).F10(), globalState)
		var _995_v8 D0
		_ = _995_v8
		_995_v8 = Companion_D0_.Create_DC1_((_987_v2).F10())
		var _pat_let_tv30 = _987_v2
		_ = _pat_let_tv30
		var _pat_let_tv31 = _986_v1
		_ = _pat_let_tv31
		var _pat_let_tv32 = _986_v1
		_ = _pat_let_tv32
		r3 = func(_source13 D0) bool {
			if _source13.Is_DC1() {
				var _996___mcc_h0 bool = _source13.Get_().(D0_DC1).Cf1
				_ = _996___mcc_h0
				var _997_cf1 bool = _996___mcc_h0
				_ = _997_cf1
				return (_dafny.SetOf(false)).IsDisjointFrom(_dafny.SetOf(true))
			} else if _source13.Is_DC0() {
				var _998___mcc_h1 _dafny.Map = _source13.Get_().(D0_DC0).Cf0
				_ = _998___mcc_h1
				var _999_cf0 _dafny.Map = _998___mcc_h1
				_ = _999_cf0
				return !(false) || ((_pat_let_tv30).F10())
			} else {
				var _1000___mcc_h2 D0 = _source13.Get_().(D0_DC2).Cf2
				_ = _1000___mcc_h2
				var _1001_cf2 D0 = _1000___mcc_h2
				_ = _1001_cf2
				return ((_pat_let_tv31).F10()) && ((_pat_let_tv32).F10())
			}
		}(_995_v8)
		return r0, r1, r2, r3
	}
}
func (_this *C4) M2(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Map, _dafny.Int, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1002_v0 _dafny.Sequence
		_ = _1002_v0
		_1002_v0 = _dafny.UnicodeSeqOfUtf8Bytes("vg")
		var _1003_v1 bool
		_ = _1003_v1
		_1003_v1 = false
		var _1004_v2 D0
		_ = _1004_v2
		_1004_v2 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1002_v0).Cardinality()), _1003_v1))
		var _pat_let_tv33 = _1003_v1
		_ = _pat_let_tv33
		var _pat_let_tv34 = _1003_v1
		_ = _pat_let_tv34
		var _pat_let_tv35 = _1003_v1
		_ = _pat_let_tv35
		var _source14 D0 = func(_source15 D0) D0 {
			if _source15.Is_DC1() {
				var _1005___mcc_h3 bool = _source15.Get_().(D0_DC1).Cf1
				_ = _1005___mcc_h3
				var _1006_cf1 bool = _1005___mcc_h3
				_ = _1006_cf1
				if true {
					return Companion_D0_.Create_DC1_(true)
				} else {
					return Companion_D0_.Create_DC1_(_pat_let_tv33)
				}
			} else if _source15.Is_DC0() {
				var _1007___mcc_h4 _dafny.Map = _source15.Get_().(D0_DC0).Cf0
				_ = _1007___mcc_h4
				var _1008_cf0 _dafny.Map = _1007___mcc_h4
				_ = _1008_cf0
				return Companion_D0_.Create_DC1_(_pat_let_tv34)
			} else {
				var _1009___mcc_h5 D0 = _source15.Get_().(D0_DC2).Cf2
				_ = _1009___mcc_h5
				var _1010_cf2 D0 = _1009___mcc_h5
				_ = _1010_cf2
				return Companion_D0_.Create_DC1_(!(_pat_let_tv35))
			}
		}(_1004_v2)
		_ = _source14
		if _source14.Is_DC1() {
			var _1011___mcc_h0 bool = _source14.Get_().(D0_DC1).Cf1
			_ = _1011___mcc_h0
			var _1012_cf1 bool = _1011___mcc_h0
			_ = _1012_cf1
			r1 = p0
			(globalState).F5 = (_dafny.Zero).Minus((_dafny.Zero).Minus(p0))
			var _1013_v3 _dafny.Sequence
			_ = _1013_v3
			_1013_v3 = _dafny.SeqOf(false)
			var _1014_v4 _dafny.Sequence
			_ = _1014_v4
			_1014_v4 = _dafny.SeqOf((_dafny.IntOfUint32((_1013_v3).Cardinality())).Cmp(p1) > 0)
			var _rhs137 bool = _1012_cf1
			_ = _rhs137
			var _rhs138 _dafny.Sequence = _1014_v4
			_ = _rhs138
			var _rhs139 _dafny.Int = _dafny.IntOfUint32((_1014_v4).Cardinality())
			_ = _rhs139
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			var _lhs105 *GlobalState = globalState
			_ = _lhs105
			_lhs104.F0 = _rhs137
			_1013_v3 = _rhs138
			_lhs105.F5 = _rhs139
			var _1015_v5 _dafny.CodePoint
			_ = _1015_v5
			_1015_v5 = _dafny.CodePoint('c')
			_1015_v5 = _1015_v5
		} else if _source14.Is_DC0() {
			var _1016___mcc_h1 _dafny.Map = _source14.Get_().(D0_DC0).Cf0
			_ = _1016___mcc_h1
			var _1017_cf0 _dafny.Map = _1016___mcc_h1
			_ = _1017_cf0
			var _1018_v6 _dafny.Map
			_ = _1018_v6
			_1018_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1003_v1, p1)
			var _1019_v7 _dafny.Array
			_ = _1019_v7
			var _nwElement0_37 bool = true
			_ = _nwElement0_37
			var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(26))
			_ = _nw182
			(_nw182).ArraySet1(_nwElement0_37, 0)
			(_nw182).ArraySet1(_1003_v1, 1)
			(_nw182).ArraySet1(_1003_v1, 2)
			(_nw182).ArraySet1(_1003_v1, 3)
			(_nw182).ArraySet1(!(_1003_v1), 4)
			(_nw182).ArraySet1(_1003_v1, 5)
			(_nw182).ArraySet1(_1003_v1, 6)
			(_nw182).ArraySet1(!(_1003_v1), 7)
			(_nw182).ArraySet1(_1003_v1, 8)
			(_nw182).ArraySet1(_1003_v1, 9)
			(_nw182).ArraySet1(false, 10)
			(_nw182).ArraySet1(_1003_v1, 11)
			(_nw182).ArraySet1(_1003_v1, 12)
			(_nw182).ArraySet1(_1003_v1, 13)
			(_nw182).ArraySet1(_1003_v1, 14)
			(_nw182).ArraySet1(_1003_v1, 15)
			(_nw182).ArraySet1(_1003_v1, 16)
			(_nw182).ArraySet1(_1003_v1, 17)
			(_nw182).ArraySet1(_1003_v1, 18)
			(_nw182).ArraySet1(_1003_v1, 19)
			(_nw182).ArraySet1(_1003_v1, 20)
			(_nw182).ArraySet1(_1003_v1, 21)
			(_nw182).ArraySet1(_1003_v1, 22)
			(_nw182).ArraySet1(!(_1003_v1), 23)
			(_nw182).ArraySet1(_1003_v1, 24)
			(_nw182).ArraySet1(_1003_v1, 25)
			_1019_v7 = _nw182
			var _1020_v8 T1
			_ = _1020_v8
			var _nw183 *C0 = New_C0_()
			_ = _nw183
			_nw183.Ctor__(_1003_v1, _1019_v7)
			_1020_v8 = _nw183
			var _1021_v9 _dafny.Map
			_ = _1021_v9
			_1021_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1018_v6, _1020_v8)
			var _1022_v10 D0
			_ = _1022_v10
			_1022_v10 = Companion_D0_.Create_DC0_(_1017_cf0)
			var _1023_v11 D0
			_ = _1023_v11
			_1023_v11 = Companion_D0_.Create_DC2_(_1022_v10)
			var _rhs140 _dafny.Map = ((_1021_v9).Merge(_1021_v9)).Update((_1018_v6).Merge(_1018_v6), (func() T1 {
				if _1003_v1 {
					return _1020_v8
				}
				return _1020_v8
			})())
			_ = _rhs140
			var _rhs141 D0 = _1023_v11
			_ = _rhs141
			_1021_v9 = _rhs140
			_1023_v11 = _rhs141
			if _1003_v1 {
				_1017_cf0 = (_1017_cf0).Update(p1, _1003_v1)
				var _1024_v12 _dafny.Set
				_ = _1024_v12
				_1024_v12 = _dafny.SetOf(false)
				var _1025_v13 _dafny.Sequence
				_ = _1025_v13
				_1025_v13 = _dafny.SeqOf(_1024_v12)
				var _1026_v14 _dafny.CodePoint
				_ = _1026_v14
				_1026_v14 = _dafny.CodePoint('f')
				var _1027_v15 _dafny.MultiSet
				_ = _1027_v15
				_1027_v15 = _dafny.MultiSetOf((Companion_D2_.Create_DC5_(_1020_v8)).Dtor_cf5(), _1020_v8, _1020_v8)
				(globalState).F5 = Companion_Default___.Fm0((_this).F9(), _dafny.IntOfUint32((_1002_v0).Cardinality()), (p1).Plus((_this).Fm5(_1025_v13, _1026_v14, _dafny.SeqOf((_1027_v15).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(835)), globalState)), _1003_v1, globalState)
				var _1028_v16 _dafny.MultiSet
				_ = _1028_v16
				_1028_v16 = _dafny.MultiSetOf(_1003_v1)
				var _1029_v17 _dafny.Map
				_ = _1029_v17
				_1029_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v16, _1003_v1)
				_1026_v14 = Companion_Default___.Fm9((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((p2).Cardinality()), (_1029_v17).Cardinality())), globalState)
				(globalState).F5 = p1
				_1017_cf0 = (_1017_cf0).Update((_dafny.Zero).Minus(p1), true)
			} else {
				var _1030_v18 _dafny.Array
				_ = _1030_v18
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_32
				var _nw184 _dafny.Array
				_ = _nw184
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw184 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Int = func(_1031_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1031_i0, (_this).F9())
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw184 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw184).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw184).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1030_v18 = _nw184
				var _1032_v19 _dafny.Set
				_ = _1032_v19
				_1032_v19 = _dafny.SetOf(_dafny.IntOfUint32((_1002_v0).Cardinality()))
				var _rhs142 _dafny.Array = _1030_v18
				_ = _rhs142
				var _rhs143 bool = false
				_ = _rhs143
				var _rhs144 _dafny.Set = _1032_v19
				_ = _rhs144
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				_1030_v18 = _rhs142
				_lhs106.F0 = _rhs143
				_1032_v19 = _rhs144
				var _1033_v20 _dafny.MultiSet
				_ = _1033_v20
				_1033_v20 = _dafny.MultiSetOf(_1003_v1)
				(globalState).F5 = Companion_Default___.SafeDivisionInt(p1, (func() _dafny.Int {
					if (_1033_v20).Contains((_1020_v8).Fm4(_1004_v2, _1004_v2, _1033_v20, globalState)) {
						return (_1033_v20).Multiplicity((_1020_v8).Fm4(_1004_v2, _1004_v2, _1033_v20, globalState))
					}
					return _dafny.IntOfInt64(-158)
				})())
				var _rhs145 _dafny.Int = (p1).Plus((_this).F9())
				_ = _rhs145
				var _rhs146 _dafny.Int = (_this).F9()
				_ = _rhs146
				var _rhs147 _dafny.Int = p1
				_ = _rhs147
				var _lhs107 *GlobalState = globalState
				_ = _lhs107
				r1 = _rhs145
				_lhs107.F5 = _rhs146
				r1 = _rhs147
				var _1034_v21 _dafny.CodePoint
				_ = _1034_v21
				_1034_v21 = _dafny.CodePoint('i')
				var _1035_v22 _dafny.Map
				_ = _1035_v22
				_1035_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1003_v1, _1034_v21)
				var _1036_v23 D3
				_ = _1036_v23
				_1036_v23 = Companion_D3_.Create_DC8_(_1002_v0)
				var _1037_v24 _dafny.Array
				_ = _1037_v24
				var _nwElement0_38 _dafny.Sequence = _1002_v0
				_ = _nwElement0_38
				var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(21))
				_ = _nw185
				(_nw185).ArraySet1(_nwElement0_38, 0)
				(_nw185).ArraySet1(_1002_v0, 1)
				(_nw185).ArraySet1(_1002_v0, 2)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1002_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(667))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}(func(_1038_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				}))), 3)
				(_nw185).ArraySet1(_1002_v0, 4)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1002_v0, _1002_v0), 5)
				(_nw185).ArraySet1(_1002_v0, 6)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1002_v0, _1002_v0), 7)
				(_nw185).ArraySet1(_1002_v0, 8)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1002_v0, _1002_v0), 9)
				(_nw185).ArraySet1(_1002_v0, 10)
				(_nw185).ArraySet1(_1002_v0, 11)
				(_nw185).ArraySet1(_1002_v0, 12)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1002_v0, _1002_v0), 13)
				(_nw185).ArraySet1(_1002_v0, 14)
				(_nw185).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-366))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}(func(_1039_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})), 15)
				(_nw185).ArraySet1(_1002_v0, 16)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(_1003_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1034_v21), _1003_v1, globalState), _1002_v0), 17)
				(_nw185).ArraySet1(Companion_Default___.Fm8(_1003_v1, _1035_v22, _1003_v1, globalState), 18)
				(_nw185).ArraySet1(_1002_v0, 19)
				(_nw185).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(273))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_1040_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1041_i3 _dafny.Int) _dafny.CodePoint {
						return _1040_v21
					}
				})(_1034_v21))), (_1036_v23).Dtor_cf10()), 20)
				_1037_v24 = _nw185
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_1037_v24), 0))
				_ = _index122
				(_1037_v24).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ggh"), _1002_v0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ggh"), _1002_v0)).Cardinality()))).Uint32(), _1034_v21), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ggh"), _1002_v0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ggh"), _1002_v0)).Cardinality()))).Uint32(), _1034_v21)).Cardinality()))).Uint32(), _1034_v21), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_1037_v24), 0))
				_ = _index123
				(_1037_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(827))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1042_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1043_i4 _dafny.Int) _dafny.CodePoint {
						return _1042_v21
					}
				})(_1034_v21))), _1002_v0), (_index123).Int())
				var _1044_v25 _dafny.Array
				_ = _1044_v25
				var _nwElement0_39 _dafny.CodePoint = _dafny.CodePoint('g')
				_ = _nwElement0_39
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(28))
				_ = _nw186
				(_nw186).ArraySet1CodePoint(_nwElement0_39, 0)
				(_nw186).ArraySet1CodePoint(_1034_v21, 1)
				(_nw186).ArraySet1CodePoint(_1034_v21, 2)
				(_nw186).ArraySet1CodePoint(_dafny.CodePoint('l'), 3)
				(_nw186).ArraySet1CodePoint(_1034_v21, 4)
				(_nw186).ArraySet1CodePoint(_1034_v21, 5)
				(_nw186).ArraySet1CodePoint(_dafny.CodePoint('v'), 6)
				(_nw186).ArraySet1CodePoint(_1034_v21, 7)
				(_nw186).ArraySet1CodePoint(_1034_v21, 8)
				(_nw186).ArraySet1CodePoint(_1034_v21, 9)
				(_nw186).ArraySet1CodePoint(_1034_v21, 10)
				(_nw186).ArraySet1CodePoint(_1034_v21, 11)
				(_nw186).ArraySet1CodePoint(_1034_v21, 12)
				(_nw186).ArraySet1CodePoint(_1034_v21, 13)
				(_nw186).ArraySet1CodePoint(_dafny.CodePoint('l'), 14)
				(_nw186).ArraySet1CodePoint(_1034_v21, 15)
				(_nw186).ArraySet1CodePoint(_1034_v21, 16)
				(_nw186).ArraySet1CodePoint(_1034_v21, 17)
				(_nw186).ArraySet1CodePoint(_1034_v21, 18)
				(_nw186).ArraySet1CodePoint(((_1037_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_1037_v24), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1033_v20).Cardinality()), _dafny.IntOfUint32(((_1037_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_1037_v24), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), 19)
				(_nw186).ArraySet1CodePoint(_1034_v21, 20)
				(_nw186).ArraySet1CodePoint(_1034_v21, 21)
				(_nw186).ArraySet1CodePoint(_1034_v21, 22)
				(_nw186).ArraySet1CodePoint(_1034_v21, 23)
				(_nw186).ArraySet1CodePoint(_1034_v21, 24)
				(_nw186).ArraySet1CodePoint(_1034_v21, 25)
				(_nw186).ArraySet1CodePoint(_1034_v21, 26)
				(_nw186).ArraySet1CodePoint(_1034_v21, 27)
				_1044_v25 = _nw186
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1044_v25), 0))
				_ = _index124
				(_1044_v25).ArraySet1CodePoint(_1034_v21, (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1044_v25), 0))
				_ = _index125
				(_1044_v25).ArraySet1CodePoint(_1034_v21, (_index125).Int())
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1044_v25), 0))
				_ = _index126
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1044_v25), 0))
				_ = _index127
				var _rhs148 _dafny.Array = _1030_v18
				_ = _rhs148
				var _rhs149 _dafny.Int = (p0).Minus(p0)
				_ = _rhs149
				var _rhs150 _dafny.CodePoint = _1034_v21
				_ = _rhs150
				var _rhs151 _dafny.CodePoint = _1034_v21
				_ = _rhs151
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				var _lhs109 _dafny.Array = _1044_v25
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1044_v25), 0))
				_ = _lhs110
				var _lhs111 _dafny.Array = _1044_v25
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_1044_v25), 0))
				_ = _lhs112
				_1030_v18 = _rhs148
				_lhs108.F5 = _rhs149
				(_lhs109).ArraySet1CodePoint(_rhs150, (_lhs110).Int())
				(_lhs111).ArraySet1CodePoint(_rhs151, (_lhs112).Int())
			}
			var _1045_v26 _dafny.Array
			_ = _1045_v26
			var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw187
			_1045_v26 = _nw187
			var _1046_v27 _dafny.Array
			_ = _1046_v27
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_33
			var _nw188 _dafny.Array
			_ = _nw188
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw188 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Sequence = (func(_1047_cf0 _dafny.Map, _1048_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1049_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D3_.Create_DC10_(Companion_D3_.Create_DC8_(_dafny.UnicodeSeqOfUtf8Bytes("ewwnswuai"))), Companion_D3_.Create_DC10_(Companion_D3_.Create_DC9_((_1047_cf0).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(724))).Uint32(), func(coer55 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
							return func(arg55 _dafny.Int) interface{} {
								return coer55(arg55)
							}
						}((func(_1050_v0 _dafny.Sequence) func(_dafny.Int) D3 {
							return func(_1051_i6 _dafny.Int) D3 {
								return Companion_D3_.Create_DC10_((Companion_D3_.Create_DC10_(Companion_D3_.Create_DC8_(_1050_v0))).Dtor_cf12())
							}
						})(_1048_v0))))
					}
				})(_1017_cf0, _1002_v0)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw188 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw188).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw188).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1046_v27 = _nw188
			var _1052_v28 D3
			_ = _1052_v28
			_1052_v28 = Companion_D3_.Create_DC8_(_1002_v0)
			var _1053_v29 _dafny.Sequence
			_ = _1053_v29
			_1053_v29 = _dafny.SeqOf(Companion_D3_.Create_DC10_(_1052_v28))
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1046_v27), 0))
			_ = _index128
			(_1046_v27).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1053_v29, _1053_v29), (_index128).Int())
			var _1054_v30 _dafny.Sequence
			_ = _1054_v30
			_1054_v30 = _dafny.SeqOf((true) || (_1003_v1))
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1046_v27), 0))
			_ = _index129
			var _rhs152 bool = (_1054_v30).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1054_v30).Cardinality()))).Uint32()).(bool)
			_ = _rhs152
			var _rhs153 _dafny.Array = _1045_v26
			_ = _rhs153
			var _rhs154 _dafny.Sequence = _1053_v29
			_ = _rhs154
			var _rhs155 _dafny.Int = (_this).F9()
			_ = _rhs155
			var _rhs156 bool = ((_this).F9()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1054_v30, _dafny.SeqOf(_1003_v1))).Cardinality())) < 0
			_ = _rhs156
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			var _lhs114 _dafny.Array = _1046_v27
			_ = _lhs114
			var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1046_v27), 0))
			_ = _lhs115
			var _lhs116 *GlobalState = globalState
			_ = _lhs116
			_lhs113.F2 = _rhs152
			_1045_v26 = _rhs153
			(_lhs114).ArraySet1(_rhs154, (_lhs115).Int())
			_lhs116.F5 = _rhs155
			r2 = _rhs156
			(globalState).F5 = p0
		} else {
			var _1055___mcc_h2 D0 = _source14.Get_().(D0_DC2).Cf2
			_ = _1055___mcc_h2
			var _1056_cf2 D0 = _1055___mcc_h2
			_ = _1056_cf2
			var _rhs157 _dafny.Int = (Companion_Default___.SafeDivisionInt(p0, (_this).F9())).Minus((_this).F9())
			_ = _rhs157
			var _rhs158 bool = _1003_v1
			_ = _rhs158
			var _rhs159 _dafny.Int = (_this).F9()
			_ = _rhs159
			var _lhs117 *GlobalState = globalState
			_ = _lhs117
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			_lhs117.F5 = _rhs157
			_1003_v1 = _rhs158
			_lhs118.F5 = _rhs159
			var _1057_v31 _dafny.Array
			_ = _1057_v31
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw189
			_1057_v31 = _nw189
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1057_v31), 0))
			_ = _index130
			(_1057_v31).ArraySet1(_dafny.IntOfInt64(903), (_index130).Int())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1057_v31), 0))
			_ = _index131
			(_1057_v31).ArraySet1(Companion_Default___.SafeDivisionInt(p1, (p0).Plus(_dafny.IntOfInt64(717))), (_index131).Int())
			var _1058_v32 _dafny.Sequence
			_ = _1058_v32
			_1058_v32 = _dafny.SeqOf(_1057_v31)
			var _1059_v33 _dafny.Array
			_ = _1059_v33
			var _nwElement0_40 _dafny.Array = (_1058_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sbtwqs")).Cardinality()), _dafny.IntOfUint32((_1058_v32).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _nwElement0_40
			var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(2))
			_ = _nw190
			(_nw190).ArraySet1(_nwElement0_40, 0)
			(_nw190).ArraySet1(_1057_v31, 1)
			_1059_v33 = _nw190
			var _1060_v34 _dafny.Array
			_ = _1060_v34
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw191
			_1060_v34 = _nw191
			var _1061_v35 _dafny.Array
			_ = _1061_v35
			var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(20))
			_ = _nw192
			_1061_v35 = _nw192
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1060_v34), 0))
			_ = _index132
			(_1060_v34).ArraySet1(_1061_v35, (_index132).Int())
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1060_v34), 0))
			_ = _index133
			var _rhs160 _dafny.Array = _1059_v33
			_ = _rhs160
			var _rhs161 _dafny.Array = _1061_v35
			_ = _rhs161
			var _rhs162 _dafny.Int = p1
			_ = _rhs162
			var _rhs163 _dafny.Sequence = _1002_v0
			_ = _rhs163
			var _rhs164 bool = _1003_v1
			_ = _rhs164
			var _lhs119 _dafny.Array = _1060_v34
			_ = _lhs119
			var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1060_v34), 0))
			_ = _lhs120
			var _lhs121 *GlobalState = globalState
			_ = _lhs121
			var _lhs122 *GlobalState = globalState
			_ = _lhs122
			_1059_v33 = _rhs160
			(_lhs119).ArraySet1(_rhs161, (_lhs120).Int())
			_lhs121.F5 = _rhs162
			_1002_v0 = _rhs163
			_lhs122.F2 = _rhs164
			var _1062_v36 D3
			_ = _1062_v36
			_1062_v36 = Companion_D3_.Create_DC9_(p0)
			var _pat_let_tv36 = _1062_v36
			_ = _pat_let_tv36
			var _1063_v37 _dafny.MultiSet
			_ = _1063_v37
			_1063_v37 = _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_1062_v36), func(_pat_let16_0 D3) D3 {
				return func(_1064_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let17_0 D3) D3 {
						return func(_1065_dt__update_hcf12_h0 D3) D3 {
							return Companion_D3_.Create_DC10_(_1065_dt__update_hcf12_h0)
						}(_pat_let17_0)
					}(_pat_let_tv36)
				}(_pat_let16_0)
			}(Companion_D3_.Create_DC10_(_1062_v36)))
			var _1066_v38 D3
			_ = _1066_v38
			_1066_v38 = Companion_D3_.Create_DC10_(_1062_v36)
			var _1067_v39 D0
			_ = _1067_v39
			_1067_v39 = Companion_D0_.Create_DC1_(_1003_v1)
			var _1068_v40 _dafny.Array
			_ = _1068_v40
			var _nwElement0_41 bool = _1003_v1
			_ = _nwElement0_41
			var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(15))
			_ = _nw193
			(_nw193).ArraySet1(_nwElement0_41, 0)
			(_nw193).ArraySet1(_1003_v1, 1)
			(_nw193).ArraySet1(_1003_v1, 2)
			(_nw193).ArraySet1(_1003_v1, 3)
			(_nw193).ArraySet1(_1003_v1, 4)
			(_nw193).ArraySet1(_1003_v1, 5)
			(_nw193).ArraySet1(true, 6)
			(_nw193).ArraySet1(true, 7)
			(_nw193).ArraySet1(true, 8)
			(_nw193).ArraySet1(_1003_v1, 9)
			(_nw193).ArraySet1(_1003_v1, 10)
			(_nw193).ArraySet1(_1003_v1, 11)
			(_nw193).ArraySet1(_1003_v1, 12)
			(_nw193).ArraySet1(true, 13)
			(_nw193).ArraySet1((_1067_v39).Dtor_cf1(), 14)
			_1068_v40 = _nw193
			var _pat_let_tv37 = _1062_v36
			_ = _pat_let_tv37
			var _1069_v41 T1
			_ = _1069_v41
			var _nw194 *C0 = New_C0_()
			_ = _nw194
			_nw194.Ctor__(!(((_1063_v37).Update(func(_pat_let18_0 D3) D3 {
				return func(_1070_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let19_0 D3) D3 {
						return func(_1071_dt__update_hcf12_h1 D3) D3 {
							return Companion_D3_.Create_DC10_(_1071_dt__update_hcf12_h1)
						}(_pat_let19_0)
					}(_pat_let_tv37)
				}(_pat_let18_0)
			}(_1066_v38), Companion_Default___.Abs((_1057_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1057_v31), 0))).Int()).(_dafny.Int)))).IsSubsetOf(_1063_v37)), _1068_v40)
			_1069_v41 = _nw194
			_1069_v41 = _1069_v41
		}
		if false {
			var _1072_v42 _dafny.Map
			_ = _1072_v42
			_1072_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1003_v1)
			var _1073_v43 _dafny.Sequence
			_ = _1073_v43
			_1073_v43 = _dafny.SeqOf(_1072_v42)
			var _1074_v44 _dafny.MultiSet
			_ = _1074_v44
			_1074_v44 = _dafny.MultiSetOf(_1003_v1)
			var _1075_v45 _dafny.Set
			_ = _1075_v45
			_1075_v45 = _dafny.SetOf(_1003_v1, _1003_v1)
			var _1076_v46 _dafny.CodePoint
			_ = _1076_v46
			_1076_v46 = _dafny.CodePoint('u')
			var _1077_v47 _dafny.Map
			_ = _1077_v47
			_1077_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1003_v1, _1003_v1)
			var _1078_v48 _dafny.Map
			_ = _1078_v48
			_1078_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1077_v47, (_1075_v45).Cardinality())
			var _1079_v49 D2
			_ = _1079_v49
			_1079_v49 = Companion_D2_.Create_DC7_(p0, _1078_v48, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1002_v0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1002_v0).Cardinality()))).Uint32(), _1076_v46)).Cardinality()))
			var _1080_v50 _dafny.Map
			_ = _1080_v50
			_1080_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1074_v44, Companion_Default___.Fm10(_1075_v45, _1076_v46, _1079_v49, p0, globalState))
			var _rhs165 bool = _1003_v1
			_ = _rhs165
			var _rhs166 bool = (_1080_v50).Contains(_1074_v44)
			_ = _rhs166
			var _rhs167 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1072_v42).Update((_this).F9(), _1003_v1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1003_v1)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-788))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_1081_v42 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1082_i7 _dafny.Int) _dafny.Map {
					return _1081_v42
				}
			})(_1072_v42)))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(971))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_1083_v42 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1084_i8 _dafny.Int) _dafny.Map {
					return _1083_v42
				}
			})(_1072_v42))))
			_ = _rhs167
			_1003_v1 = _rhs165
			_1003_v1 = _rhs166
			_1073_v43 = _rhs167
			(globalState).F0 = _1003_v1
			var _1085_v51 D2
			_ = _1085_v51
			_1085_v51 = Companion_D2_.Create_DC6_(_1003_v1)
			var _source16 D2 = _1085_v51
			_ = _source16
			if _source16.Is_DC6() {
				var _1086___mcc_h6 bool = _source16.Get_().(D2_DC6).Cf6
				_ = _1086___mcc_h6
				var _1087_cf6 bool = _1086___mcc_h6
				_ = _1087_cf6
				var _1088_v52 _dafny.Array
				_ = _1088_v52
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_34
				var _nw195 _dafny.Array
				_ = _nw195
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw195 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) bool = (func(_1089_cf6 bool) func(_dafny.Int) bool {
						return func(_1090_i9 _dafny.Int) bool {
							return _1089_cf6
						}
					})(_1087_cf6)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw195 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw195).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw195).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1088_v52 = _nw195
				var _1091_v53 *C0
				_ = _1091_v53
				var _nw196 *C0 = New_C0_()
				_ = _nw196
				_nw196.Ctor__(_1087_cf6, _1088_v52)
				_1091_v53 = _nw196
				var _1092_v54 _dafny.Sequence
				_ = _1092_v54
				_1092_v54 = _dafny.SeqOf(_1091_v53, _1091_v53)
				var _1093_v55 _dafny.Array
				_ = _1093_v55
				var _nwElement0_42 bool = _1087_cf6
				_ = _nwElement0_42
				var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(15))
				_ = _nw197
				(_nw197).ArraySet1(_nwElement0_42, 0)
				(_nw197).ArraySet1(_1003_v1, 1)
				(_nw197).ArraySet1(!(true) || (_1003_v1), 2)
				(_nw197).ArraySet1(_1087_cf6, 3)
				(_nw197).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1092_v54, _dafny.SeqOf(_1091_v53)), 4)
				(_nw197).ArraySet1(_1003_v1, 5)
				(_nw197).ArraySet1(_1087_cf6, 6)
				(_nw197).ArraySet1((_1091_v53).F10(), 7)
				(_nw197).ArraySet1(!(_1003_v1) || (_1087_cf6), 8)
				(_nw197).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1002_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(130))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_1094_v46 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1095_i10 _dafny.Int) _dafny.CodePoint {
						return _1094_v46
					}
				})(_1076_v46)))), 9)
				(_nw197).ArraySet1(_1087_cf6, 10)
				(_nw197).ArraySet1(_1003_v1, 11)
				(_nw197).ArraySet1(_1003_v1, 12)
				(_nw197).ArraySet1(_1087_cf6, 13)
				(_nw197).ArraySet1(_1003_v1, 14)
				_1093_v55 = _nw197
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1093_v55), 0))
				_ = _index134
				(_1093_v55).ArraySet1(!(_1087_cf6), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1093_v55), 0))
				_ = _index135
				(_1093_v55).ArraySet1(((_this).F9()).Cmp(p0) > 0, (_index135).Int())
				(globalState).F2 = (_1091_v53).F10()
				var _1096_v56 _dafny.Array
				_ = _1096_v56
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw198
				_1096_v56 = _nw198
				_1096_v56 = (func() _dafny.Array {
					if _1087_cf6 {
						return _1096_v56
					}
					return _1096_v56
				})()
				var _1097_v57 _dafny.Sequence
				_ = _1097_v57
				_1097_v57 = _dafny.SeqOf((_1093_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1093_v55), 0))).Int()).(bool), (_1091_v53).F10())
				var _1098_v58 _dafny.Map
				_ = _1098_v58
				_1098_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1076_v46, (_1091_v53).F11())
				var _1099_v59 D4
				_ = _1099_v59
				_1099_v59 = Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_1097_v57).Cardinality()), (func() _dafny.Array {
					if (_1098_v58).Contains(_1076_v46) {
						return (_1098_v58).Get(_1076_v46).(_dafny.Array)
					}
					return _1093_v55
				})())
				var _1100_v60 _dafny.Map
				_ = _1100_v60
				_1100_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1099_v59).Dtor_cf20(), _1088_v52)
				var _1101_v61 T1
				_ = _1101_v61
				var _nw199 *C0 = New_C0_()
				_ = _nw199
				_nw199.Ctor__((_1091_v53).F10(), (func() _dafny.Array {
					if (_1100_v60).Contains(_1088_v52) {
						return (_1100_v60).Get(_1088_v52).(_dafny.Array)
					}
					return _1088_v52
				})())
				_1101_v61 = _nw199
				var _1102_v62 _dafny.Map
				_ = _1102_v62
				_1102_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1101_v61)
				var _1103_v63 _dafny.Sequence
				_ = _1103_v63
				_1103_v63 = _dafny.SeqOf(_1075_v45)
				var _1104_v64 _dafny.Sequence
				_ = _1104_v64
				_1104_v64 = _dafny.SeqOf(_dafny.IntOfUint32((_1097_v57).Cardinality()), p1)
				var _1105_v65 _dafny.Set
				_ = _1105_v65
				_1105_v65 = _dafny.SetOf(_dafny.IntOfUint32((_1097_v57).Cardinality()))
				var _1106_v66 _dafny.Array
				_ = _1106_v66
				var _nwElement0_43 _dafny.Map = _1102_v62
				_ = _nwElement0_43
				var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(16))
				_ = _nw200
				(_nw200).ArraySet1(_nwElement0_43, 0)
				(_nw200).ArraySet1(_1102_v62, 1)
				(_nw200).ArraySet1(_1102_v62, 2)
				(_nw200).ArraySet1(_1102_v62, 3)
				(_nw200).ArraySet1(_1102_v62, 4)
				(_nw200).ArraySet1((_1102_v62).Update((_this).F9(), _1101_v61), 5)
				(_nw200).ArraySet1(_1102_v62, 6)
				(_nw200).ArraySet1(_1102_v62, 7)
				(_nw200).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1101_v61), 8)
				(_nw200).ArraySet1(_1102_v62, 9)
				(_nw200).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))).Cardinality(), _1101_v61), 10)
				(_nw200).ArraySet1((_1102_v62).Merge(_1102_v62), 11)
				(_nw200).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1079_v49).Dtor_cf7(), _1101_v61), 12)
				(_nw200).ArraySet1((_1102_v62).Update((_this).Fm5(_1103_v63, _1076_v46, _1104_v64, _dafny.SeqOf((_dafny.Zero).Minus((_1105_v65).Cardinality()), p0), globalState), _1101_v61), 13)
				(_nw200).ArraySet1(_1102_v62, 14)
				(_nw200).ArraySet1((_1102_v62).Merge(_1102_v62), 15)
				_1106_v66 = _nw200
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_1106_v66), 0))
				_ = _index136
				(_1106_v66).ArraySet1(_1102_v62, (_index136).Int())
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_1106_v66), 0))
				_ = _index137
				(_1106_v66).ArraySet1(_1102_v62, (_index137).Int())
			} else if _source16.Is_DC7() {
				var _1107___mcc_h7 _dafny.Int = _source16.Get_().(D2_DC7).Cf7
				_ = _1107___mcc_h7
				var _1108___mcc_h8 _dafny.Map = _source16.Get_().(D2_DC7).Cf8
				_ = _1108___mcc_h8
				var _1109___mcc_h9 _dafny.Int = _source16.Get_().(D2_DC7).Cf9
				_ = _1109___mcc_h9
				var _1110_cf9 _dafny.Int = _1109___mcc_h9
				_ = _1110_cf9
				var _1111_cf8 _dafny.Map = _1108___mcc_h8
				_ = _1111_cf8
				var _1112_cf7 _dafny.Int = _1107___mcc_h7
				_ = _1112_cf7
				(globalState).F0 = ((func() bool {
					if _1003_v1 {
						return _1003_v1
					}
					return _1003_v1
				})()) && (_1003_v1)
				var _1113_v67 _dafny.Sequence
				_ = _1113_v67
				_1113_v67 = _dafny.SeqOf(_1112_cf7, (_1077_v47).Cardinality())
				_1113_v67 = _1113_v67
				_1079_v49 = _1079_v49
				var _1114_v68 D0
				_ = _1114_v68
				_1114_v68 = Companion_D0_.Create_DC1_(false)
				var _1115_v69 _dafny.Sequence
				_ = _1115_v69
				_1115_v69 = _dafny.SeqOf(_1003_v1, _1003_v1, _1003_v1, _1003_v1, (_1114_v68).Dtor_cf1())
				var _1116_v70 _dafny.Array
				_ = _1116_v70
				var _nw201 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw201
				_1116_v70 = _nw201
				var _1117_v71 D4
				_ = _1117_v71
				_1117_v71 = Companion_D4_.Create_DC13_(_1110_cf9, _1116_v70)
				(globalState).F0 = (_1115_v69).Select((Companion_Default___.SafeIndex((_1117_v71).Dtor_cf19(), _dafny.IntOfUint32((_1115_v69).Cardinality()))).Uint32()).(bool)
			} else {
				var _1118___mcc_h10 T1 = _source16.Get_().(D2_DC5).Cf5
				_ = _1118___mcc_h10
				var _1119_cf5 T1 = _1118___mcc_h10
				_ = _1119_cf5
				r2 = true
				var _1120_v72 _dafny.Sequence
				_ = _1120_v72
				_1120_v72 = _dafny.SeqOf(_1003_v1)
				var _1121_v73 D0
				_ = _1121_v73
				_1121_v73 = Companion_D0_.Create_DC1_((_1120_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1120_v72).Cardinality()))).Uint32()).(bool))
				var _1122_v74 _dafny.Map
				_ = _1122_v74
				_1122_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC9_(_dafny.IntOfInt64(610)), _1121_v73)
				_1122_v74 = _1122_v74
				var _1123_v75 _dafny.Array
				_ = _1123_v75
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_35
				var _nw202 _dafny.Array
				_ = _nw202
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw202 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) bool = (func(_1124_v1 bool) func(_dafny.Int) bool {
						return func(_1125_i11 _dafny.Int) bool {
							return _1124_v1
						}
					})(_1003_v1)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw202 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw202).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw202).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1123_v75 = _nw202
				var _1126_v76 *C0
				_ = _1126_v76
				var _nw203 *C0 = New_C0_()
				_ = _nw203
				_nw203.Ctor__(true, _1123_v75)
				_1126_v76 = _nw203
				var _1127_v77 _dafny.Array
				_ = _1127_v77
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_36
				var _nw204 _dafny.Array
				_ = _nw204
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw204 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = (func(_1128_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1129_i12 _dafny.Int) _dafny.Int {
							return (_1129_i12).Plus(_1128_p1)
						}
					})(p1)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw204 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw204).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw204).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1127_v77 = _nw204
				var _1130_v78 _dafny.Map
				_ = _1130_v78
				_1130_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1126_v76, _1127_v77)
				var _1131_v79 _dafny.Map
				_ = _1131_v79
				_1131_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1126_v76).F10(), _1077_v47)
				var _1132_v80 _dafny.Map
				_ = _1132_v80
				_1132_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_1131_v79).Cardinality(), p0, p1, (_1126_v76).F10(), globalState), p1)
				var _rhs168 bool = _1003_v1
				_ = _rhs168
				var _rhs169 _dafny.Int = (func() _dafny.Int {
					if _1003_v1 {
						return ((_1130_v78).Merge(_1130_v78)).Cardinality()
					}
					return (func() _dafny.Int {
						if (_1126_v76).F10() {
							return (_this).F9()
						}
						return p0
					})()
				})()
				_ = _rhs169
				var _rhs170 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), p0)).Merge((_1132_v80).Merge(_1132_v80))).Cardinality()
				_ = _rhs170
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				var _lhs124 *GlobalState = globalState
				_ = _lhs124
				_lhs123.F2 = _rhs168
				r1 = _rhs169
				_lhs124.F5 = _rhs170
				var _1133_v81 _dafny.MultiSet
				_ = _1133_v81
				_1133_v81 = _dafny.MultiSetOf(_1072_v42, (_1072_v42).Merge(_1072_v42), _1072_v42, _1072_v42)
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen(((_1126_v76).F11()), 0))
				_ = _index138
				((_1126_v76).F11()).ArraySet1((_1126_v76).F10(), (_index138).Int())
				var _1134_v82 _dafny.MultiSet
				_ = _1134_v82
				_1134_v82 = _dafny.MultiSetOf(_1002_v0)
				var _1135_v83 _dafny.Set
				_ = _1135_v83
				_1135_v83 = _dafny.SetOf(p1)
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen(((_1126_v76).F11()), 0))
				_ = _index139
				var _rhs171 _dafny.Int = p0
				_ = _rhs171
				var _rhs172 _dafny.Int = ((_1134_v82).Cardinality()).Times((func() _dafny.Int {
					if (_1126_v76).F10() {
						return (_this).F9()
					}
					return Companion_Default___.Fm0((_1135_v83).Cardinality(), p0, p1, (_1085_v51).Dtor_cf6(), globalState)
				})())
				_ = _rhs172
				var _rhs173 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1120_v72, _dafny.SeqOf((_1126_v76).F10())), _1120_v72)).Cardinality())
				_ = _rhs173
				var _rhs174 _dafny.MultiSet = _1133_v81
				_ = _rhs174
				var _rhs175 bool = _1003_v1
				_ = _rhs175
				var _lhs125 *GlobalState = globalState
				_ = _lhs125
				var _lhs126 *GlobalState = globalState
				_ = _lhs126
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 _dafny.Array = (_1126_v76).F11()
				_ = _lhs128
				var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen(((_1126_v76).F11()), 0))
				_ = _lhs129
				_lhs125.F5 = _rhs171
				_lhs126.F5 = _rhs172
				_lhs127.F5 = _rhs173
				_1133_v81 = _rhs174
				(_lhs128).ArraySet1(_rhs175, (_lhs129).Int())
			}
			if (p1).Cmp((_1075_v45).Cardinality()) <= 0 {
				_1003_v1 = _1003_v1
				var _1136_v84 D0
				_ = _1136_v84
				_1136_v84 = Companion_D0_.Create_DC1_(_1003_v1)
				var _1137_v85 _dafny.Array
				_ = _1137_v85
				var _nwElement0_44 bool = false
				_ = _nwElement0_44
				var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(3))
				_ = _nw205
				(_nw205).ArraySet1(_nwElement0_44, 0)
				(_nw205).ArraySet1(_1003_v1, 1)
				(_nw205).ArraySet1((_1136_v84).Dtor_cf1(), 2)
				_1137_v85 = _nw205
				var _1138_v86 *C0
				_ = _1138_v86
				var _nw206 *C0 = New_C0_()
				_ = _nw206
				_nw206.Ctor__(!(!(((_this).F9()).Cmp((_this).F9()) >= 0)), _1137_v85)
				_1138_v86 = _nw206
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1137_v85), 0))
				_ = _index140
				(_1137_v85).ArraySet1(!(_1003_v1) || ((_1138_v86).F10()), (_index140).Int())
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1137_v85), 0))
				_ = _index141
				(_1137_v85).ArraySet1(!(_1003_v1), (_index141).Int())
				var _1139_v87 _dafny.Map
				_ = _1139_v87
				_1139_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1074_v44, (_1137_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1137_v85), 0))).Int()).(bool))
				r1 = (((_1139_v87).Merge(_1139_v87)).Merge(_1139_v87)).Cardinality()
				var _1140_v88 _dafny.Array
				_ = _1140_v88
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
				_ = _nw207
				_1140_v88 = _nw207
				var _1141_v89 T0
				_ = _1141_v89
				var _nw208 *C2 = New_C2_()
				_ = _nw208
				_nw208.Ctor__(!((_1138_v86).F10()))
				_1141_v89 = _nw208
				var _1142_v90 _dafny.Sequence
				_ = _1142_v90
				_1142_v90 = _dafny.SeqOf((_1137_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1137_v85), 0))).Int()).(bool), (_1138_v86).F10(), (_1137_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1137_v85), 0))).Int()).(bool))
				var _1143_v91 D7
				_ = _1143_v91
				_1143_v91 = Companion_D7_.Create_DC17_(_1142_v90)
				var _1144_v92 _dafny.MultiSet
				_ = _1144_v92
				_1144_v92 = _dafny.MultiSetOf(Companion_D7_.Create_DC17_(_dafny.SeqOf((_1138_v86).F10())), _1143_v91)
				var _1145_v93 _dafny.MultiSet
				_ = _1145_v93
				_1145_v93 = _dafny.MultiSetOf(_1141_v89)
				var _rhs176 _dafny.Array = _1140_v88
				_ = _rhs176
				var _rhs177 bool = (_1137_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1137_v85), 0))).Int()).(bool)
				_ = _rhs177
				var _rhs178 bool = (_1138_v86).F10()
				_ = _rhs178
				var _rhs179 T0 = _1141_v89
				_ = _rhs179
				var _rhs180 _dafny.Set = _dafny.SetOf(_1003_v1, (_1144_v92).IsDisjointFrom(_1144_v92), (_1145_v93).IsDisjointFrom(_1145_v93))
				_ = _rhs180
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				_1140_v88 = _rhs176
				_lhs130.F2 = _rhs177
				r2 = _rhs178
				_1141_v89 = _rhs179
				_lhs131.F8 = _rhs180
			} else {
				var _rhs181 bool = _1003_v1
				_ = _rhs181
				var _rhs182 _dafny.Int = p1
				_ = _rhs182
				var _rhs183 bool = !(!(_1003_v1))
				_ = _rhs183
				var _rhs184 _dafny.Int = (_this).F9()
				_ = _rhs184
				var _lhs132 *GlobalState = globalState
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				var _lhs135 *GlobalState = globalState
				_ = _lhs135
				_lhs132.F2 = _rhs181
				_lhs133.F5 = _rhs182
				_lhs134.F2 = _rhs183
				_lhs135.F5 = _rhs184
				var _1146_v94 _dafny.Map
				_ = _1146_v94
				_1146_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Minus(_dafny.IntOfInt64(755)), (_dafny.Zero).Minus(p0))
				var _1147_v96 _dafny.Sequence
				_ = _1147_v96
				_1147_v96 = _dafny.SeqOf(_1003_v1, !(_1003_v1))
				_1146_v94 = (_1146_v94).Update(Companion_Default___.Fm0(p0, p1, (func() _dafny.Map {
					var _coll48 = _dafny.NewMapBuilder()
					_ = _coll48
					for _iter51 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)).Keys().Elements()); ; {
						_compr_48, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1148_v95 _dafny.Int
						_1148_v95 = interface{}(_compr_48).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)).Contains(_1148_v95) {
							_coll48.Add((_1148_v95).Minus((_dafny.Zero).Minus((_dafny.SetOf(p0, p0)).Cardinality())), _dafny.IntOfUint32((_1002_v0).Cardinality()))
						}
					}
					return _coll48.ToMap()
				}()).Cardinality(), true, globalState), _dafny.IntOfUint32((_1147_v96).Cardinality()))
				(globalState).F0 = true
				var _1149_v97 _dafny.Array
				_ = _1149_v97
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw209
				_1149_v97 = _nw209
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1149_v97), 0))
				_ = _index142
				(_1149_v97).ArraySet1(p0, (_index142).Int())
				var _1150_v98 _dafny.Array
				_ = _1150_v98
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw210
				_1150_v98 = _nw210
				var _1151_v99 *C0
				_ = _1151_v99
				var _nw211 *C0 = New_C0_()
				_ = _nw211
				_nw211.Ctor__(_1003_v1, _1150_v98)
				_1151_v99 = _nw211
				var _1152_v100 _dafny.MultiSet
				_ = _1152_v100
				_1152_v100 = _dafny.MultiSetOf(_1151_v99, _1151_v99, _1151_v99, _1151_v99, _1151_v99)
				var _1153_v101 _dafny.Sequence
				_ = _1153_v101
				_1153_v101 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(p1)).Cardinality(), (_this).F9())).Cardinality())
				var _1154_v102 _dafny.Sequence
				_ = _1154_v102
				_1154_v102 = _dafny.SeqOf(_1002_v0, _1002_v0)
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1149_v97), 0))
				_ = _index143
				var _rhs185 _dafny.Int = (func() _dafny.Int {
					if (_1152_v100).Contains(_1151_v99) {
						return (_1152_v100).Multiplicity(_1151_v99)
					}
					return (_dafny.Zero).Minus((func() _dafny.Int {
						if (_1151_v99).F10() {
							return (_1153_v101).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1147_v96).Cardinality()), _dafny.IntOfUint32((_1153_v101).Cardinality()))).Uint32()).(_dafny.Int)
						}
						return (_this).F9()
					})())
				})()
				_ = _rhs185
				var _rhs186 _dafny.Int = (_dafny.Zero).Minus(((p1).Plus(p1)).Times(p1))
				_ = _rhs186
				var _rhs187 _dafny.Int = (func() _dafny.Int {
					if Companion_Default___.Fm1((_1151_v99).F10(), globalState) {
						return _dafny.IntOfUint32((_1154_v102).Cardinality())
					}
					return (p1).Times((_this).F9())
				})()
				_ = _rhs187
				var _lhs136 *GlobalState = globalState
				_ = _lhs136
				var _lhs137 _dafny.Array = _1149_v97
				_ = _lhs137
				var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1149_v97), 0))
				_ = _lhs138
				r1 = _rhs185
				_lhs136.F5 = _rhs186
				(_lhs137).ArraySet1(_rhs187, (_lhs138).Int())
				var _1155_v103 *C3
				_ = _1155_v103
				var _nw212 *C3 = New_C3_()
				_ = _nw212
				_nw212.Ctor__((_dafny.Zero).Minus(p0))
				_1155_v103 = _nw212
				var _1156_v104 _dafny.Array
				_ = _1156_v104
				var _nwElement0_45 *C3 = _1155_v103
				_ = _nwElement0_45
				var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(8))
				_ = _nw213
				(_nw213).ArraySet1(_nwElement0_45, 0)
				(_nw213).ArraySet1(_1155_v103, 1)
				(_nw213).ArraySet1(_1155_v103, 2)
				(_nw213).ArraySet1(_1155_v103, 3)
				(_nw213).ArraySet1(_1155_v103, 4)
				(_nw213).ArraySet1(_1155_v103, 5)
				(_nw213).ArraySet1(_1155_v103, 6)
				(_nw213).ArraySet1(_1155_v103, 7)
				_1156_v104 = _nw213
				var _1157_v105 _dafny.Set
				_ = _1157_v105
				_1157_v105 = _dafny.SetOf(_1076_v46)
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1149_v97), 0))
				_ = _index144
				var _rhs188 bool = ((func() _dafny.Int {
					if _1003_v1 {
						return (_this).F9()
					}
					return (_1157_v105).Cardinality()
				})()).Cmp(_dafny.IntOfInt64(-348)) <= 0
				_ = _rhs188
				var _rhs189 bool = !_dafny.Companion_Sequence_.Equal(_1147_v96, _1147_v96)
				_ = _rhs189
				var _rhs190 _dafny.Array = _1156_v104
				_ = _rhs190
				var _rhs191 _dafny.Sequence = Companion_Default___.Fm27(globalState)
				_ = _rhs191
				var _rhs192 _dafny.Int = (_1149_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1149_v97), 0))).Int()).(_dafny.Int)
				_ = _rhs192
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				var _lhs140 _dafny.Array = _1149_v97
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1149_v97), 0))
				_ = _lhs141
				_lhs139.F2 = _rhs188
				_1003_v1 = _rhs189
				_1156_v104 = _rhs190
				_1002_v0 = _rhs191
				(_lhs140).ArraySet1(_rhs192, (_lhs141).Int())
			}
			var _1158_v106 _dafny.Array
			_ = _1158_v106
			var _len0_37 _dafny.Int = _dafny.One
			_ = _len0_37
			var _nw214 _dafny.Array
			_ = _nw214
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw214 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) bool = (func(_1159_v1 bool) func(_dafny.Int) bool {
					return func(_1160_i13 _dafny.Int) bool {
						return _1159_v1
					}
				})(_1003_v1)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw214 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw214).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw214).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1158_v106 = _nw214
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_1158_v106), 0))
			_ = _index145
			(_1158_v106).ArraySet1((func() bool {
				if _1003_v1 {
					return _1003_v1
				}
				return false
			})(), (_index145).Int())
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.ArrayLen((_1158_v106), 0))
			_ = _index146
			(_1158_v106).ArraySet1(_1003_v1, (_index146).Int())
		} else {
			var _1161_v107 _dafny.CodePoint
			_ = _1161_v107
			_1161_v107 = _dafny.CodePoint('k')
			var _1162_v108 _dafny.Map
			_ = _1162_v108
			_1162_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1161_v107)
			(globalState).F5 = ((_1162_v108).Cardinality()).Times(_dafny.IntOfInt64(535))
			r1 = (_dafny.Zero).Minus(p0)
			var _1163_v109 _dafny.Map
			_ = _1163_v109
			_1163_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1003_v1)
			var _1164_v110 D0
			_ = _1164_v110
			_1164_v110 = Companion_D0_.Create_DC0_(_1163_v109)
			var _1165_v111 D0
			_ = _1165_v111
			_1165_v111 = Companion_D0_.Create_DC2_(_1164_v110)
			var _1166_v112 D0
			_ = _1166_v112
			_1166_v112 = Companion_D0_.Create_DC2_(_1165_v111)
			var _source17 D0 = Companion_D0_.Create_DC2_((func() D0 {
				if _1003_v1 {
					return _1165_v111
				}
				return _1165_v111
			})())
			_ = _source17
			if _source17.Is_DC1() {
				var _1167___mcc_h11 bool = _source17.Get_().(D0_DC1).Cf1
				_ = _1167___mcc_h11
				var _1168_cf1 bool = _1167___mcc_h11
				_ = _1168_cf1
				var _1169_v113 _dafny.Set
				_ = _1169_v113
				_1169_v113 = _dafny.SetOf(_dafny.CodePoint('l'), _dafny.CodePoint('m'), _1161_v107)
				var _1170_v114 _dafny.Sequence
				_ = _1170_v114
				_1170_v114 = _dafny.SeqOf(p0)
				var _1171_v115 _dafny.Sequence
				_ = _1171_v115
				_1171_v115 = _dafny.SeqOf(_1168_cf1, false)
				var _1172_v116 _dafny.Set
				_ = _1172_v116
				_1172_v116 = _dafny.SetOf(_1003_v1, _1168_cf1)
				var _1173_v117 _dafny.Map
				_ = _1173_v117
				_1173_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1172_v116, _1168_cf1)
				var _1174_v118 _dafny.Sequence
				_ = _1174_v118
				_1174_v118 = _dafny.SeqOf((_1173_v117).Update(_dafny.SetOf(_1003_v1), true), _1173_v117, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1172_v116, _1168_cf1))
				var _1175_v119 _dafny.Map
				_ = _1175_v119
				_1175_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(889), p1)
				var _1176_v120 _dafny.Sequence
				_ = _1176_v120
				_1176_v120 = _dafny.SeqOf(_1002_v0)
				var _1177_v121 _dafny.Sequence
				_ = _1177_v121
				_1177_v121 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(406))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_1178_v107 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1179_i14 _dafny.Int) _dafny.CodePoint {
						return _1178_v107
					}
				})(_1161_v107))), (_1176_v120).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1176_v120).Cardinality()))).Uint32()).(_dafny.Sequence), _1002_v0, _dafny.UnicodeSeqOfUtf8Bytes("nikx"), _1002_v0)
				var _1180_v122 _dafny.Array
				_ = _1180_v122
				var _nwElement0_46 _dafny.Int = p0
				_ = _nwElement0_46
				var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(25))
				_ = _nw215
				(_nw215).ArraySet1(_nwElement0_46, 0)
				(_nw215).ArraySet1((func() _dafny.Int {
					if (func() bool {
						if (_1163_v109).Contains((_1169_v113).Cardinality()) {
							return (_1163_v109).Get((_1169_v113).Cardinality()).(bool)
						}
						return _1003_v1
					})() {
						return Companion_Default___.Fm0((_this).F9(), (_this).F9(), p0, _1168_cf1, globalState)
					}
					return (_this).F9()
				})(), 1)
				(_nw215).ArraySet1((_this).F9(), 2)
				(_nw215).ArraySet1((_1170_v114).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1170_v114).Cardinality()))).Uint32()).(_dafny.Int), 3)
				(_nw215).ArraySet1((func() _dafny.Int {
					if _1003_v1 {
						return (_this).F9()
					}
					return (_this).F9()
				})(), 4)
				(_nw215).ArraySet1((_this).F9(), 5)
				(_nw215).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfUint32((_1171_v115).Cardinality()), _dafny.IntOfInt64(-482), p1, true, globalState), 6)
				(_nw215).ArraySet1((_this).F9(), 7)
				(_nw215).ArraySet1(p1, 8)
				(_nw215).ArraySet1(p0, 9)
				(_nw215).ArraySet1((p0).Times((_dafny.SetOf((_this).F9())).Cardinality()), 10)
				(_nw215).ArraySet1(p0, 11)
				(_nw215).ArraySet1(((_1174_v118).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1170_v114).Cardinality())), _dafny.IntOfUint32((_1174_v118).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 12)
				(_nw215).ArraySet1(p0, 13)
				(_nw215).ArraySet1(((_1175_v119).Merge(_1175_v119)).Cardinality(), 14)
				(_nw215).ArraySet1((_this).F9(), 15)
				(_nw215).ArraySet1((_dafny.IntOfUint32((_1171_v115).Cardinality())).Times((_1172_v116).Cardinality()), 16)
				(_nw215).ArraySet1(p1, 17)
				(_nw215).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1177_v121).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_dafny.IntOfUint32((_1002_v0).Cardinality()), _dafny.IntOfInt64(378), (_this).F9(), true, globalState), _dafny.IntOfUint32((_1177_v121).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.IntOfUint32(((_1177_v121).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_dafny.IntOfUint32((_1002_v0).Cardinality()), _dafny.IntOfInt64(378), (_this).F9(), true, globalState), _dafny.IntOfUint32((_1177_v121).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1161_v107)).Cardinality()), 18)
				(_nw215).ArraySet1(p1, 19)
				(_nw215).ArraySet1(p0, 20)
				(_nw215).ArraySet1((_this).F9(), 21)
				(_nw215).ArraySet1(_dafny.IntOfInt64(21), 22)
				(_nw215).ArraySet1(p1, 23)
				(_nw215).ArraySet1(p1, 24)
				_1180_v122 = _nw215
				_1180_v122 = _1180_v122
				(globalState).F8 = ((_1172_v116).Difference(_1172_v116)).Union((_1172_v116).Union(_1172_v116))
				var _1181_v123 *C3
				_ = _1181_v123
				var _nw216 *C3 = New_C3_()
				_ = _nw216
				_nw216.Ctor__(p0)
				_1181_v123 = _nw216
				var _1182_v124 _dafny.Array
				_ = _1182_v124
				var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
				_ = _nw217
				_1182_v124 = _nw217
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_1182_v124), 0))
				_ = _index147
				(_1182_v124).ArraySet1(_1002_v0, (_index147).Int())
				var _1183_v125 _dafny.Array
				_ = _1183_v125
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_38
				var _nw218 _dafny.Array
				_ = _nw218
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw218 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) bool = (func(_1184_v1 bool) func(_dafny.Int) bool {
						return func(_1185_i15 _dafny.Int) bool {
							return _1184_v1
						}
					})(_1003_v1)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw218 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw218).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw218).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1183_v125 = _nw218
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1183_v125), 0))
				_ = _index148
				(_1183_v125).ArraySet1(_1003_v1, (_index148).Int())
				var _1186_v126 _dafny.Set
				_ = _1186_v126
				_1186_v126 = _dafny.SetOf((_1181_v123).F12())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_1182_v124), 0))
				_ = _index149
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1183_v125), 0))
				_ = _index150
				var _rhs193 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1002_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-283))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_1187_v107 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1188_i16 _dafny.Int) _dafny.CodePoint {
						return _1187_v107
					}
				})(_1161_v107)))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1002_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-283))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}((func(_1189_v107 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1190_i16 _dafny.Int) _dafny.CodePoint {
						return _1189_v107
					}
				})(_1161_v107))))).Cardinality()))).Uint32(), _1161_v107)
				_ = _rhs193
				var _rhs194 bool = _1003_v1
				_ = _rhs194
				var _rhs195 bool = _1003_v1
				_ = _rhs195
				var _rhs196 _dafny.Int = _dafny.IntOfInt64(292)
				_ = _rhs196
				var _rhs197 bool = ((_1186_v126).Difference(_1186_v126)).IsSubsetOf(_1186_v126)
				_ = _rhs197
				var _lhs142 _dafny.Array = _1182_v124
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_1182_v124), 0))
				_ = _lhs143
				var _lhs144 _dafny.Array = _1183_v125
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1183_v125), 0))
				_ = _lhs145
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				(_lhs142).ArraySet1(_rhs193, (_lhs143).Int())
				r2 = _rhs194
				(_lhs144).ArraySet1(_rhs195, (_lhs145).Int())
				_lhs146.F5 = _rhs196
				_lhs147.F2 = _rhs197
			} else if _source17.Is_DC0() {
				var _1191___mcc_h12 _dafny.Map = _source17.Get_().(D0_DC0).Cf0
				_ = _1191___mcc_h12
				var _1192_cf0 _dafny.Map = _1191___mcc_h12
				_ = _1192_cf0
				var _1193_v127 _dafny.Array
				_ = _1193_v127
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_39
				var _nw219 _dafny.Array
				_ = _nw219
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw219 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) bool = (func(_1194_v1 bool) func(_dafny.Int) bool {
						return func(_1195_i17 _dafny.Int) bool {
							return !(_1194_v1) || (false)
						}
					})(_1003_v1)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw219 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw219).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw219).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1193_v127 = _nw219
				var _1196_v128 D4
				_ = _1196_v128
				_1196_v128 = Companion_D4_.Create_DC11_(_1193_v127)
				var _pat_let_tv38 = _1193_v127
				_ = _pat_let_tv38
				_1193_v127 = (func(_pat_let20_0 D4) D4 {
					return func(_1197_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let21_0 _dafny.Array) D4 {
							return func(_1198_dt__update_hcf13_h0 _dafny.Array) D4 {
								return Companion_D4_.Create_DC11_(_1198_dt__update_hcf13_h0)
							}(_pat_let21_0)
						}(_pat_let_tv38)
					}(_pat_let20_0)
				}(_1196_v128)).Dtor_cf13()
				var _1199_v129 *C1
				_ = _1199_v129
				var _nw220 *C1 = New_C1_()
				_ = _nw220
				_nw220.Ctor__()
				_1199_v129 = _nw220
				var _1200_v130 *C3
				_ = _1200_v130
				var _nw221 *C3 = New_C3_()
				_ = _nw221
				_nw221.Ctor__((_dafny.Zero).Minus((_this).F9()))
				_1200_v130 = _nw221
				_1192_cf0 = (_1192_cf0).Update((_1200_v130).F12(), _1003_v1)
			} else {
				var _1201___mcc_h13 D0 = _source17.Get_().(D0_DC2).Cf2
				_ = _1201___mcc_h13
				var _1202_cf2 D0 = _1201___mcc_h13
				_ = _1202_cf2
				(globalState).F5 = (_this).F9()
				(globalState).F5 = p1
				r1 = (_this).F9()
				var _1203_v131 *C1
				_ = _1203_v131
				var _nw222 *C1 = New_C1_()
				_ = _nw222
				_nw222.Ctor__()
				_1203_v131 = _nw222
			}
			var _1204_v132 _dafny.Array
			_ = _1204_v132
			var _nwElement0_47 bool = !(true) || (_1003_v1)
			_ = _nwElement0_47
			var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.One)
			_ = _nw223
			(_nw223).ArraySet1(_nwElement0_47, 0)
			_1204_v132 = _nw223
			var _1205_v133 _dafny.Map
			_ = _1205_v133
			_1205_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1002_v0, _1003_v1)
			var _1206_v134 _dafny.MultiSet
			_ = _1206_v134
			_1206_v134 = _dafny.MultiSetOf(_1003_v1, _1003_v1)
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1204_v132), 0))
			_ = _index151
			(_1204_v132).ArraySet1((func() bool {
				if (_1205_v133).Contains(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rngu"), (Companion_Default___.SafeIndex((_1206_v134).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rngu")).Cardinality()))).Uint32(), _1161_v107)) {
					return (_1205_v133).Get(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rngu"), (Companion_Default___.SafeIndex((_1206_v134).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rngu")).Cardinality()))).Uint32(), _1161_v107)).(bool)
				}
				return _1003_v1
			})(), (_index151).Int())
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1204_v132), 0))
			_ = _index152
			(_1204_v132).ArraySet1(_1003_v1, (_index152).Int())
			r1 = (_1206_v134).Cardinality()
		}
		var _1207_v135 *C1
		_ = _1207_v135
		var _nw224 *C1 = New_C1_()
		_ = _nw224
		_nw224.Ctor__()
		_1207_v135 = _nw224
		var _1208_v136 _dafny.Map
		_ = _1208_v136
		_1208_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1207_v135, _1003_v1)
		var _1209_v137 _dafny.Map
		_ = _1209_v137
		_1209_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1003_v1, _1003_v1)
		var _1210_v138 _dafny.CodePoint
		_ = _1210_v138
		_1210_v138 = _dafny.CodePoint('o')
		var _1211_v139 D10
		_ = _1211_v139
		_1211_v139 = Companion_D10_.Create_DC27_(_1003_v1, p1, _1210_v138, _dafny.MultiSetOf(_1003_v1))
		var _1212_v140 _dafny.Sequence
		_ = _1212_v140
		_1212_v140 = _dafny.SeqOf(_1002_v0, _1002_v0, _1002_v0)
		var _1213_v141 _dafny.Set
		_ = _1213_v141
		_1213_v141 = _dafny.SetOf((_1212_v140).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1212_v140).Cardinality()))).Uint32()).(_dafny.Sequence), _1002_v0)
		var _1214_v142 _dafny.Array
		_ = _1214_v142
		var _nwElement0_48 bool = !(_1003_v1)
		_ = _nwElement0_48
		var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(27))
		_ = _nw225
		(_nw225).ArraySet1(_nwElement0_48, 0)
		(_nw225).ArraySet1(_1003_v1, 1)
		(_nw225).ArraySet1((func() bool {
			if !(!(_1003_v1)) {
				return (func() bool {
					if (_1208_v136).Contains(_1207_v135) {
						return (_1208_v136).Get(_1207_v135).(bool)
					}
					return _1003_v1
				})()
			}
			return false
		})(), 2)
		(_nw225).ArraySet1(Companion_Default___.Fm1(_1003_v1, globalState), 3)
		(_nw225).ArraySet1(!(_1003_v1) || (_1003_v1), 4)
		(_nw225).ArraySet1((_1003_v1) == (_1003_v1), 5)
		(_nw225).ArraySet1(!(true), 6)
		(_nw225).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1003_v1)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm1(_1003_v1, globalState))), 7)
		(_nw225).ArraySet1(!(_1003_v1) || ((func() bool {
			if (_1209_v137).Contains(_1003_v1) {
				return (_1209_v137).Get(_1003_v1).(bool)
			}
			return _1003_v1
		})()), 8)
		(_nw225).ArraySet1((Companion_D9_.Create_DC24_(_1003_v1, _1003_v1, _1003_v1)).Dtor_cf35(), 9)
		(_nw225).ArraySet1(_1003_v1, 10)
		(_nw225).ArraySet1(_1003_v1, 11)
		(_nw225).ArraySet1(_1003_v1, 12)
		(_nw225).ArraySet1(_1003_v1, 13)
		(_nw225).ArraySet1(!(false), 14)
		(_nw225).ArraySet1(!(false), 15)
		(_nw225).ArraySet1(_1003_v1, 16)
		(_nw225).ArraySet1(((_this).F9()).Cmp((_this).F9()) > 0, 17)
		(_nw225).ArraySet1((_1207_v135).Fm4(_1004_v2, _1004_v2, (_1211_v139).Dtor_cf41(), globalState), 18)
		(_nw225).ArraySet1(false, 19)
		(_nw225).ArraySet1((_1213_v141).IsSubsetOf(_dafny.SetOf(_1002_v0, _1002_v0, _1002_v0, _dafny.UnicodeSeqOfUtf8Bytes("q"))), 20)
		(_nw225).ArraySet1((p1).Cmp((_this).F9()) == 0, 21)
		(_nw225).ArraySet1(_1003_v1, 22)
		(_nw225).ArraySet1(!(_1003_v1) || (_1003_v1), 23)
		(_nw225).ArraySet1((_1003_v1) && (_1003_v1), 24)
		(_nw225).ArraySet1(_1003_v1, 25)
		(_nw225).ArraySet1(_1003_v1, 26)
		_1214_v142 = _nw225
		var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))
		_ = _index153
		(_1214_v142).ArraySet1((_1003_v1) && (_1003_v1), (_index153).Int())
		var _1215_v143 _dafny.Array
		_ = _1215_v143
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_40
		var _nw226 _dafny.Array
		_ = _nw226
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw226 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) _dafny.MultiSet = (func(_1216_v1 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_1217_i18 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(912))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg62 _dafny.Int) interface{} {
							return coer62(arg62)
						}
					}(func(_1218_i19 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(607)
					}))).Cardinality()), _dafny.SetOf(_1216_v1, _1216_v1))).Cardinality(), _dafny.IntOfInt64(241))
				}
			})(_1003_v1)
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw226 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw226).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw226).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1215_v143 = _nw226
		var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1215_v143), 0))
		_ = _index154
		(_1215_v143).ArraySet1(_dafny.MultiSetOf(p1), (_index154).Int())
		var _1219_v144 _dafny.MultiSet
		_ = _1219_v144
		_1219_v144 = _dafny.MultiSetOf((_dafny.MultiSetOf(_1002_v0, _1002_v0, _1002_v0)).Cardinality(), (_this).F9(), (_this).F9())
		var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))
		_ = _index155
		var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1215_v143), 0))
		_ = _index156
		var _rhs198 bool = _1003_v1
		_ = _rhs198
		var _rhs199 _dafny.MultiSet = _1219_v144
		_ = _rhs199
		var _lhs148 _dafny.Array = _1214_v142
		_ = _lhs148
		var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))
		_ = _lhs149
		var _lhs150 _dafny.Array = _1215_v143
		_ = _lhs150
		var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1215_v143), 0))
		_ = _lhs151
		(_lhs148).ArraySet1(_rhs198, (_lhs149).Int())
		(_lhs150).ArraySet1(_rhs199, (_lhs151).Int())
		var _1220_v145 *C2
		_ = _1220_v145
		var _nw227 *C2 = New_C2_()
		_ = _nw227
		_nw227.Ctor__(_1003_v1)
		_1220_v145 = _nw227
		var _1221_v146 D7
		_ = _1221_v146
		_1221_v146 = Companion_D7_.Create_DC18_()
		_1221_v146 = (func() D7 {
			if (_1214_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))).Int()).(bool) {
				return _1221_v146
			}
			return _1221_v146
		})()
		var _1222_v147 D11
		_ = _1222_v147
		_1222_v147 = Companion_D11_.Create_DC31_(p0, (_dafny.Zero).Minus((_this).F9()))
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))
		_ = _index157
		(_1214_v142).ArraySet1(((_1222_v147).Dtor_cf45()).Cmp(p1) < 0, (_index157).Int())
		var _1223_v148 _dafny.Sequence
		_ = _1223_v148
		_1223_v148 = _dafny.SeqOf((_1220_v145).F13())
		var _1224_v149 _dafny.Sequence
		_ = _1224_v149
		_1224_v149 = _dafny.SeqOf(_1223_v148, _1223_v148)
		var _1225_v150 _dafny.Map
		_ = _1225_v150
		_1225_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm32(_1224_v149, globalState), true)
		var _1226_v151 _dafny.Map
		_ = _1226_v151
		_1226_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1225_v150).Cardinality())
		r0 = _1226_v151
		var _1227_v152 _dafny.MultiSet
		_ = _1227_v152
		_1227_v152 = _dafny.MultiSetOf((_1214_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))).Int()).(bool))
		var _1228_v153 D7
		_ = _1228_v153
		_1228_v153 = Companion_D7_.Create_DC19_(Companion_D7_.Create_DC18_())
		var _1229_v154 _dafny.Map
		_ = _1229_v154
		_1229_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1228_v153, (_1214_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))).Int()).(bool))
		var _1230_v155 _dafny.Map
		_ = _1230_v155
		_1230_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1212_v140, (_this).F9())
		r1 = (Companion_Default___.SafeDivisionInt(p0, (func() _dafny.Int {
			if (_1227_v152).Contains((func() bool {
				if (_1229_v154).Contains(_1228_v153) {
					return (_1229_v154).Get(_1228_v153).(bool)
				}
				return (_1214_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))).Int()).(bool)
			})()) {
				return (_1227_v152).Multiplicity((func() bool {
					if (_1229_v154).Contains(_1228_v153) {
						return (_1229_v154).Get(_1228_v153).(bool)
					}
					return (_1214_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1214_v142), 0))).Int()).(bool)
				})())
			}
			return _dafny.IntOfInt64(-660)
		})())).Plus((func() _dafny.Int {
			if (_1230_v155).Contains(_dafny.SeqOf(_1002_v0)) {
				return (_1230_v155).Get(_dafny.SeqOf(_1002_v0)).(_dafny.Int)
			}
			return p0
		})())
		var _1231_v156 _dafny.Set
		_ = _1231_v156
		_1231_v156 = _dafny.SetOf(_1214_v142)
		r2 = (Companion_Default___.SafeModuloInt((_1231_v156).Cardinality(), _dafny.IntOfInt64(326))).Cmp(p1) == 0
		return r0, r1, r2
	}
}
func (_this *C4) F9() _dafny.Int {
	{
		return _this._f9
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
