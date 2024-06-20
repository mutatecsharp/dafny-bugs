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
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("jfyngxuaa")
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.SeqOf(Companion_D6_.Create_DC24_(Companion_D6_.Create_DC21_(_dafny.MultiSetOf(_dafny.IntOfInt64(198)))), Companion_D6_.Create_DC24_(Companion_D6_.Create_DC21_(_dafny.MultiSetOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(155))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("crncpbwhy")
		}))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Sequence
			_1_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(155))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_0_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("crncpbwhy")
			})), _1_v0) {
				_coll0.Add(_1_v0, _dafny.IntOfInt64(275))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()))), Companion_D6_.Create_DC24_(Companion_D6_.Create_DC22_(_dafny.IntOfInt64(-556))), Companion_D6_.Create_DC24_(Companion_D6_.Create_DC21_(_dafny.MultiSetOf(_dafny.IntOfInt64(245), _dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(754), _dafny.IntOfInt64(-590), _dafny.IntOfInt64(382), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('t'))).Cardinality(), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality())).Cardinality()))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(754), _dafny.IntOfInt64(-590), _dafny.IntOfInt64(382), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('t'))).Cardinality(), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality())).Cardinality())), _2_v1) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_2_v1, _dafny.IntOfInt64(-199)), false)
			}
		}
		return _coll1.ToMap()
	}())).Cardinality())))))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(!(false), false)).Cardinality())).Cardinality()), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('x'), _dafny.CodePoint('u'))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _3_v0 _dafny.CodePoint
			_3_v0 = interface{}(_compr_2).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('x'), _dafny.CodePoint('u'))).Contains(_3_v0) {
				_coll2.Add(_3_v0)
			}
		}
		return _coll2.ToSet()
	}()).Cardinality()), true))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Sequence {
	if (!(true)) || (true) {
		return _dafny.SeqOf(_dafny.IntOfInt64(276))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(142))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_4_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_5_i1 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(44))).Cardinality()))).Cardinality()
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Union(_dafny.SetOf(!(false)))).Intersection(_dafny.SetOf(false, true, true))
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return (Companion_D10_.Create_DC34_(_dafny.CodePoint('h'), false, false, true, false)).Dtor_cf55()
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(547))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), _dafny.UnicodeSeqOfUtf8Bytes("luposje")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(624))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_7_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	})), _dafny.UnicodeSeqOfUtf8Bytes("yiwbhd")))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("thuvag"), _dafny.IntOfInt64(-296))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("njhdgt"), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sgehlao")).Cardinality())))).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(990))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("em"))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _9_v0 _dafny.Sequence
			_9_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(990))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})), _dafny.UnicodeSeqOfUtf8Bytes("em")), _9_v0) {
				_coll3.Add(_9_v0, _dafny.IntOfInt64(322))
			}
		}
		return _coll3.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source0 D6 = (func() D6 {
		if true {
			return Companion_D6_.Create_DC23_(true, _dafny.IntOfInt64(-860), !(true))
		}
		return Companion_D6_.Create_DC23_(true, _dafny.IntOfInt64(663), true)
	})()
	_ = _source0
	if _source0.Is_DC22() {
		var _10___mcc_h0 _dafny.Int = _source0.Get_().(D6_DC22).Cf36
		_ = _10___mcc_h0
		var _11_cf36 _dafny.Int = _10___mcc_h0
		_ = _11_cf36
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_cf36, true)).Merge(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(180), _dafny.IntOfInt64(854))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _12_v0 _dafny.Int
				_12_v0 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(180)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(854)) < 0) {
					_coll4.Add(Companion_Default___.SafeModuloInt(_12_v0, _11_cf36), false)
				}
			}
			return _coll4.ToMap()
		}())
	} else if _source0.Is_DC23() {
		var _13___mcc_h1 bool = _source0.Get_().(D6_DC23).Cf37
		_ = _13___mcc_h1
		var _14___mcc_h2 _dafny.Int = _source0.Get_().(D6_DC23).Cf38
		_ = _14___mcc_h2
		var _15___mcc_h3 bool = _source0.Get_().(D6_DC23).Cf39
		_ = _15___mcc_h3
		var _16_cf39 bool = _15___mcc_h3
		_ = _16_cf39
		var _17_cf38 _dafny.Int = _14___mcc_h2
		_ = _17_cf38
		var _18_cf37 bool = _13___mcc_h1
		_ = _18_cf37
		return (func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_cf38, _18_cf37)).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _19_v1 _dafny.Int
				_19_v1 = interface{}(_compr_5).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_cf38, _18_cf37)).Contains(_19_v1) {
					_coll5.Add((_19_v1).Times(_17_cf38), _18_cf37)
				}
			}
			return _coll5.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_cf38, _18_cf37))
	} else if _source0.Is_DC21() {
		var _20___mcc_h4 _dafny.MultiSet = _source0.Get_().(D6_DC21).Cf35
		_ = _20___mcc_h4
		var _21_cf35 _dafny.MultiSet = _20___mcc_h4
		_ = _21_cf35
		return func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate((_21_cf35).Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _22_v3 _dafny.Int
					_22_v3 = interface{}(_compr_7).(_dafny.Int)
					if (_21_cf35).Contains(_22_v3) {
						_coll7.Add((_22_v3).Minus((_dafny.MultiSetOf(true, true, true)).Cardinality()))
					}
				}
				return _coll7.ToSet()
			}()).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _23_v2 _dafny.Int
				_23_v2 = interface{}(_compr_6).(_dafny.Int)
				if (func() _dafny.Set {
					var _coll8 = _dafny.NewBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate((_21_cf35).Elements()); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _24_v3 _dafny.Int
						_24_v3 = interface{}(_compr_8).(_dafny.Int)
						if (_21_cf35).Contains(_24_v3) {
							_coll8.Add((_24_v3).Minus((_dafny.MultiSetOf(true, true, true)).Cardinality()))
						}
					}
					return _coll8.ToSet()
				}()).Contains(_23_v2) {
					_coll6.Add(Companion_Default___.SafeModuloInt(_23_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("udnqqbuso")).Cardinality())), !(!(false)))
				}
			}
			return _coll6.ToMap()
		}()
	} else {
		var _25___mcc_h5 D6 = _source0.Get_().(D6_DC24).Cf40
		_ = _25___mcc_h5
		var _26_cf40 D6 = _25___mcc_h5
		_ = _26_cf40
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), true))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('k')
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xunsdlrj")).Cardinality())))).Cardinality())).Cardinality()))).Intersection(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality()))))).Union((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pqwo")).Cardinality()), false)).Cardinality())).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _27_v0 _dafny.Int
			_27_v0 = interface{}(_compr_9).(_dafny.Int)
			if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pqwo")).Cardinality()), false)).Cardinality())).Contains(_27_v0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_27_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(296), (_dafny.SetOf(_dafny.IntOfInt64(276), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(157), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pvxtdfslh")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mlkshnmc")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(968), _dafny.IntOfInt64(660), _dafny.IntOfInt64(940)))).Cardinality())).Cardinality())).Cardinality()))
			}
		}
		return _coll9.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfInt64(723))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_28_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality()))), _dafny.SeqOf(_dafny.IntOfInt64(743), _dafny.IntOfInt64(349), _dafny.IntOfInt64(228), _dafny.IntOfInt64(-108))), _dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gw")).Cardinality()), _dafny.IntOfInt64(-30))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC10_(func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587))).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _29_v1 _dafny.Int
					_29_v1 = interface{}(_compr_12).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587)), _29_v1) {
						_coll12.Add(Companion_Default___.SafeDivisionInt(_29_v1, _dafny.IntOfInt64(890)), true)
					}
				}
				return _coll12.ToMap()
			}()).Keys().Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _30_v0 _dafny.Int
				_30_v0 = interface{}(_compr_11).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll13 = _dafny.NewMapBuilder()
					_ = _coll13
					for _iter13 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587))).Elements()); ; {
						_compr_13, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _29_v1 _dafny.Int
						_29_v1 = interface{}(_compr_13).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587)), _29_v1) {
							_coll13.Add(Companion_Default___.SafeDivisionInt(_29_v1, _dafny.IntOfInt64(890)), true)
						}
					}
					return _coll13.ToMap()
				}()).Contains(_30_v0) {
					_coll11.Add(Companion_Default___.SafeModuloInt(_30_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), false)).Cardinality())))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality())))).Cardinality()))
				}
			}
			return _coll11.ToMap()
		}()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality()))).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _32_v2 D3
			_32_v2 = interface{}(_compr_10).(D3)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC10_(func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((func() _dafny.Map {
					var _coll15 = _dafny.NewMapBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587))).Elements()); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _29_v1 _dafny.Int
						_29_v1 = interface{}(_compr_15).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587)), _29_v1) {
							_coll15.Add(Companion_Default___.SafeDivisionInt(_29_v1, _dafny.IntOfInt64(890)), true)
						}
					}
					return _coll15.ToMap()
				}()).Keys().Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _30_v0 _dafny.Int
					_30_v0 = interface{}(_compr_14).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll16 = _dafny.NewMapBuilder()
						_ = _coll16
						for _iter16 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587))).Elements()); ; {
							_compr_16, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _29_v1 _dafny.Int
							_29_v1 = interface{}(_compr_16).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(67), _dafny.IntOfInt64(98))).Cardinality()), _dafny.IntOfInt64(587)), _29_v1) {
								_coll16.Add(Companion_Default___.SafeDivisionInt(_29_v1, _dafny.IntOfInt64(890)), true)
							}
						}
						return _coll16.ToMap()
					}()).Contains(_30_v0) {
						_coll14.Add(Companion_Default___.SafeModuloInt(_30_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), false)).Cardinality())))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality())))).Cardinality()))
					}
				}
				return _coll14.ToMap()
			}()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))).Cardinality()))).Contains(_32_v2) {
				_coll10.Add(_32_v2)
			}
		}
		return _coll10.ToSet()
	}()).Union(_dafny.SetOf(Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(575), _dafny.IntOfInt64(899)))))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(466), _dafny.IntOfInt64(203)), _dafny.SeqOf(_dafny.IntOfInt64(107), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("udohgqy")).Cardinality())), (func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(_dafny.IntOfInt64(591))
		}
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(326))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_33_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (Companion_D6_.Create_DC23_(false, _dafny.IntOfInt64(174), false)).Dtor_cf38(), _dafny.IntOfInt64(942))
	})(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(239))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_34_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-517)
	}))).Cardinality()), (_dafny.SetOf(false)).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(685))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_35_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(855)
	})), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mogyijfrl")).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(553))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_36_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("sf")
	}))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(-295))))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 bool, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Map {
	var _source1 D5 = Companion_D5_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("eql"), (_dafny.Zero).Minus((_dafny.MultiSetOf(true, true, false, true, !(true))).Cardinality())))
	_ = _source1
	if _source1.Is_DC18() {
		var _37___mcc_h0 _dafny.Int = _source1.Get_().(D5_DC18).Cf30
		_ = _37___mcc_h0
		var _38___mcc_h1 _dafny.Int = _source1.Get_().(D5_DC18).Cf31
		_ = _38___mcc_h1
		var _39_cf31 _dafny.Int = _38___mcc_h1
		_ = _39_cf31
		var _40_cf30 _dafny.Int = _37___mcc_h0
		_ = _40_cf30
		return func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(936), _dafny.IntOfInt64(244))); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _41_v0 _dafny.Int
				_41_v0 = interface{}(_compr_17).(_dafny.Int)
				if ((_dafny.IntOfInt64(936)).Cmp(_41_v0) <= 0) && ((_41_v0).Cmp(_dafny.IntOfInt64(244)) < 0) {
					_coll17.Add((_41_v0).Plus((_dafny.Zero).Minus(_40_cf30)), _40_cf30)
				}
			}
			return _coll17.ToMap()
		}()
	} else if _source1.Is_DC19() {
		var _42___mcc_h2 _dafny.Map = _source1.Get_().(D5_DC19).Cf32
		_ = _42___mcc_h2
		var _43___mcc_h3 bool = _source1.Get_().(D5_DC19).Cf33
		_ = _43___mcc_h3
		var _44_cf33 bool = _43___mcc_h3
		_ = _44_cf33
		var _45_cf32 _dafny.Map = _42___mcc_h2
		_ = _45_cf32
		return _45_cf32
	} else if _source1.Is_DC17() {
		var _46___mcc_h4 _dafny.Map = _source1.Get_().(D5_DC17).Cf29
		_ = _46___mcc_h4
		var _47_cf29 _dafny.Map = _46___mcc_h4
		_ = _47_cf29
		return (Companion_D5_.Create_DC19_(func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(213), _dafny.IntOfInt64(935))); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _48_v1 _dafny.Int
				_48_v1 = interface{}(_compr_18).(_dafny.Int)
				if ((_dafny.IntOfInt64(213)).Cmp(_48_v1) <= 0) && ((_48_v1).Cmp(_dafny.IntOfInt64(935)) < 0) {
					_coll18.Add((_48_v1).Minus(_dafny.IntOfInt64(-704)), _dafny.IntOfInt64(766))
				}
			}
			return _coll18.ToMap()
		}(), !(false))).Dtor_cf32()
	} else {
		var _49___mcc_h5 D5 = _source1.Get_().(D5_DC20).Cf34
		_ = _49___mcc_h5
		var _50_cf34 D5 = _49___mcc_h5
		_ = _50_cf34
		return func() _dafny.Map {
			var _coll19 = _dafny.NewMapBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-975), _dafny.IntOfInt64(332))); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _51_v2 _dafny.Int
				_51_v2 = interface{}(_compr_19).(_dafny.Int)
				if ((_dafny.IntOfInt64(-975)).Cmp(_51_v2) <= 0) && ((_51_v2).Cmp(_dafny.IntOfInt64(332)) < 0) {
					_coll19.Add(Companion_Default___.SafeDivisionInt(_51_v2, _dafny.IntOfInt64(-974)), _dafny.IntOfInt64(-753))
				}
			}
			return _coll19.ToMap()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(_dafny.IntOfInt64(626))
		}
		return _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
	})()).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-947))))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_52_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("wuqb")
	})(), _dafny.UnicodeSeqOfUtf8Bytes("nfp"))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, globalState *GlobalState) D4 {
	if true {
		return Companion_D4_.Create_DC15_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (Companion_D12_.Create_DC42_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nekfj")).Cardinality()))).Dtor_cf69())).Cardinality(), _dafny.IntOfInt64(237))
	} else {
		return Companion_D4_.Create_DC15_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cpwd")).Cardinality()), true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(!(true))).Cardinality())).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC25_((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yhjvyhhbk")).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(270)), _dafny.SeqOf(_dafny.IntOfInt64(281), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ujoxgmcix")).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(240)), _dafny.SeqOf(_dafny.IntOfInt64(-985)))).Union(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(815), (_dafny.MultiSetOf(_dafny.IntOfInt64(439))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(381), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('c'))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(53))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_53_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), _dafny.UnicodeSeqOfUtf8Bytes("qsupnh"))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Sequence, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) bool {
	return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ujybiveij"), _dafny.UnicodeSeqOfUtf8Bytes("vvlkj"))) || ((func() bool {
		if true {
			return true
		}
		return true
	})())
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('x')
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (Companion_D19_.Create_DC59_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetFromSeq(_dafny.SeqOf(false))))).Dtor_cf89()
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	if ((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-647)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(442))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_54_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-834), false)).Cardinality(), _dafny.IntOfInt64(258))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())) > 0 {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(337), true)
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-765), true)
	}
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Set, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm42(globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("icyaksiiv"), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hf")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.CodePoint, p1 D0, p2 bool, p3 _dafny.MultiSet, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC36_((func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(876))).Uint32(), func(coer18 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_55_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_dafny.IntOfInt64(-360))
			}))
		}
		return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-909)))
	})())
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(78), !(true))).Cardinality(), _dafny.CodePoint('y'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nsq")).Cardinality())), _dafny.CodePoint('c')))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-871)), _dafny.CodePoint('d')))
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_((_dafny.IntOfInt64(424)).Times(_dafny.IntOfInt64(-232)))
}
func (_static *CompanionStruct_Default___) Fm46(globalState *GlobalState) _dafny.MultiSet {
	if (false) && (false) {
		return _dafny.MultiSetOf(Companion_D6_.Create_DC22_(_dafny.IntOfInt64(-148)))
	} else {
		return _dafny.MultiSetOf(Companion_D6_.Create_DC22_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mko")).Cardinality())), Companion_D6_.Create_DC22_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(858))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_56_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality())), Companion_D6_.Create_DC22_((_dafny.MultiSetOf(false, false, true)).Cardinality()), Companion_D6_.Create_DC22_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xh")).Cardinality())), Companion_D6_.Create_DC22_(_dafny.IntOfInt64(936)))
	}
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(21), _dafny.IntOfInt64(456))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aridp")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func() _dafny.Map {
		var _coll20 = _dafny.NewMapBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(329), _dafny.IntOfInt64(-646))); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _57_v0 _dafny.Int
			_57_v0 = interface{}(_compr_20).(_dafny.Int)
			if ((_dafny.IntOfInt64(329)).Cmp(_57_v0) <= 0) && ((_57_v0).Cmp(_dafny.IntOfInt64(-646)) < 0) {
				_coll20.Add(Companion_Default___.SafeDivisionInt(_57_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Cardinality()), _dafny.IntOfInt64(-100))
			}
		}
		return _coll20.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm48(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(true, false, true, false)), _dafny.SeqOf(_dafny.SeqOf((Companion_D2_.Create_DC7_(_dafny.SeqOf((Companion_D2_.Create_DC6_((_dafny.Zero).Minus(_dafny.IntOfInt64(-604)))).Dtor_cf9(), ((Companion_D20_.Create_DC61_(func() _dafny.Map {
		var _coll21 = _dafny.NewMapBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("me"), _dafny.UnicodeSeqOfUtf8Bytes("uwx"))).Elements()); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _58_v0 _dafny.Sequence
			_58_v0 = interface{}(_compr_21).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("me"), _dafny.UnicodeSeqOfUtf8Bytes("uwx")), _58_v0) {
				_coll21.Add(_58_v0, true)
			}
		}
		return _coll21.ToMap()
	}())).Dtor_cf93()).Cardinality()), !(false), !(!(!(!(false)))), (_dafny.MultiSetOf((_dafny.SetOf(true, !(false))).Cardinality(), _dafny.IntOfInt64(-983), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality(), _dafny.IntOfInt64(480), (func() _dafny.Map {
		var _coll22 = _dafny.NewMapBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-807), _dafny.IntOfInt64(704))); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _59_v1 _dafny.Int
			_59_v1 = interface{}(_compr_22).(_dafny.Int)
			if ((_dafny.IntOfInt64(-807)).Cmp(_59_v1) <= 0) && ((_59_v1).Cmp(_dafny.IntOfInt64(704)) < 0) {
				_coll22.Add(Companion_Default___.SafeDivisionInt(_59_v1, _dafny.IntOfInt64(861)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}(func(_60_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality()))
			}
		}
		return _coll22.ToMap()
	}()).Cardinality())).Cardinality())).Dtor_cf11()), _dafny.SeqOf(!(false)), _dafny.SeqOf(true), _dafny.SeqOf(!(false)), _dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm49(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), true)
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(func(_source2 D4) _dafny.Sequence {
		if _source2.Is_DC13() {
			var _61___mcc_h0 D0 = _source2.Get_().(D4_DC13).Cf22
			_ = _61___mcc_h0
			var _62_cf22 D0 = _61___mcc_h0
			_ = _62_cf22
			return _dafny.SeqOf(_dafny.SeqOf(true))
		} else if _source2.Is_DC14() {
			var _63___mcc_h1 _dafny.CodePoint = _source2.Get_().(D4_DC14).Cf23
			_ = _63___mcc_h1
			var _64___mcc_h2 _dafny.Array = _source2.Get_().(D4_DC14).Cf24
			_ = _64___mcc_h2
			var _65___mcc_h3 _dafny.Sequence = _source2.Get_().(D4_DC14).Cf25
			_ = _65___mcc_h3
			var _66_cf25 _dafny.Sequence = _65___mcc_h3
			_ = _66_cf25
			var _67_cf24 _dafny.Array = _64___mcc_h2
			_ = _67_cf24
			var _68_cf23 _dafny.CodePoint = _63___mcc_h1
			_ = _68_cf23
			return _dafny.SeqOf(_dafny.SeqOf(!(false)), _dafny.SeqOf(true, true), _dafny.SeqOf(false), _dafny.SeqOf(false))
		} else if _source2.Is_DC15() {
			var _69___mcc_h4 _dafny.Int = _source2.Get_().(D4_DC15).Cf26
			_ = _69___mcc_h4
			var _70___mcc_h5 _dafny.Int = _source2.Get_().(D4_DC15).Cf27
			_ = _70___mcc_h5
			var _71_cf27 _dafny.Int = _70___mcc_h5
			_ = _71_cf27
			var _72_cf26 _dafny.Int = _69___mcc_h4
			_ = _72_cf26
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(false, (Companion_D15_.Create_DC51_((Companion_D6_.Create_DC23_(false, _71_cf27, true)).Dtor_cf39(), !(false))).Dtor_cf83()), _dafny.SeqOf(false), _dafny.SeqOf(true, true), _dafny.SeqOf(false, true)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(565))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_73_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(false)
			})))
		} else if _source2.Is_DC12() {
			var _74___mcc_h6 _dafny.Map = _source2.Get_().(D4_DC12).Cf21
			_ = _74___mcc_h6
			var _75_cf21 _dafny.Map = _74___mcc_h6
			_ = _75_cf21
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(!(false), false, !(false), true, true), _dafny.SeqOf(true, !(true)), _dafny.SeqOf(true)), _dafny.SeqOf(_dafny.SeqOf(!(false))))
		} else {
			var _76___mcc_h7 D4 = _source2.Get_().(D4_DC16).Cf28
			_ = _76___mcc_h7
			var _77_cf28 D4 = _76___mcc_h7
			_ = _77_cf28
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(false)), _dafny.SeqOf(_dafny.SeqOf(true, true, false, false, false)))
		}
	}(Companion_D4_.Create_DC15_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(473), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-197))).Cardinality())).Cardinality())).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.IntOfInt64(50))))
}
func (_static *CompanionStruct_Default___) Fm51(globalState *GlobalState) D4 {
	if (!(!(true))) || (false) {
		return Companion_D4_.Create_DC12_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(933)))))
	} else {
		return Companion_D4_.Create_DC12_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-648), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfInt64(132))))
	}
}
func (_static *CompanionStruct_Default___) Fm52(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((Companion_D13_.Create_DC44_(_dafny.SeqOf(true))).Dtor_cf71(), _dafny.SeqOf(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm53(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_78_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(296))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_79_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))
	}))
}
func (_static *CompanionStruct_Default___) Fm54(p0 D15, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}(func(_80_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vslhotx")).Cardinality()))
			}))
		}
		return _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ctbyue")).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(554))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_81_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(767)
		})))
	})(), (func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(754))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}(func(_82_i2 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_dafny.IntOfInt64(-507))
			}))
		}
		return _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rvii")).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(-259), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(803))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_83_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality())))).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqnhvr")).Cardinality())))
	})())
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	if (func() _dafny.Set {
		var _coll23 = _dafny.NewBuilder()
		_ = _coll23
		for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(73), _dafny.IntOfInt64(703))); ; {
			_compr_23, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _84_v0 _dafny.Int
			_84_v0 = interface{}(_compr_23).(_dafny.Int)
			if ((_dafny.IntOfInt64(73)).Cmp(_84_v0) <= 0) && ((_84_v0).Cmp(_dafny.IntOfInt64(703)) < 0) {
				_coll23.Add((_84_v0).Times(_dafny.IntOfInt64(697)))
			}
		}
		return _coll23.ToSet()
	}()).IsSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(952), _dafny.IntOfInt64(130), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())) {
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("lnfrahgri"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-769))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_85_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})), _dafny.UnicodeSeqOfUtf8Bytes("as"), _dafny.UnicodeSeqOfUtf8Bytes("jkaix"), _dafny.UnicodeSeqOfUtf8Bytes("ujnwkx"))
	} else {
		return _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_86_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_87_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-605), _dafny.IntOfInt64(-71)), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sg")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xxkfrh")).Cardinality())))), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), _dafny.CodePoint('k'))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _88_v0 _dafny.MultiSet
	_ = _88_v0
	_88_v0 = _dafny.MultiSetOf(false)
	var _89_v1 _dafny.MultiSet
	_ = _89_v1
	_89_v1 = _dafny.MultiSetOf(_88_v0, _88_v0)
	var _90_v2 bool
	_ = _90_v2
	_90_v2 = true
	var _91_v3 _dafny.Sequence
	_ = _91_v3
	_91_v3 = _dafny.SeqOf(_90_v2, _90_v2)
	var _92_v4 _dafny.Int
	_ = _92_v4
	_92_v4 = _dafny.IntOfInt64(166)
	var _93_v5 _dafny.Sequence
	_ = _93_v5
	_93_v5 = _dafny.UnicodeSeqOfUtf8Bytes("c")
	var _94_v6 _dafny.MultiSet
	_ = _94_v6
	_94_v6 = _dafny.MultiSetOf(_dafny.IntOfUint32((_93_v5).Cardinality()), _dafny.IntOfInt64(476))
	var _95_v7 _dafny.MultiSet
	_ = _95_v7
	_95_v7 = _dafny.MultiSetOf((_94_v6).Cardinality())
	var _96_v8 _dafny.Sequence
	_ = _96_v8
	_96_v8 = _dafny.SeqOf(_95_v7)
	var _97_v9 _dafny.Sequence
	_ = _97_v9
	_97_v9 = _dafny.SeqOf((_96_v8).Select((Companion_Default___.SafeIndex(_92_v4, _dafny.IntOfUint32((_96_v8).Cardinality()))).Uint32()).(_dafny.MultiSet))
	var _98_v10 _dafny.Array
	_ = _98_v10
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = func(_99_i0 _dafny.Int) bool {
			return true
		}
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
	_98_v10 = _nw0
	var _100_v11 _dafny.Array
	_ = _100_v11
	var _nwElement0_0 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_92_v4))
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(7))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(656))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg31 _dafny.Int) interface{} {
			return coer31(arg31)
		}
	}(func(_101_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))).Cardinality()), 1)
	(_nw1).ArraySet1(_92_v4, 2)
	(_nw1).ArraySet1(_92_v4, 3)
	(_nw1).ArraySet1(_92_v4, 4)
	(_nw1).ArraySet1(_92_v4, 5)
	(_nw1).ArraySet1(_92_v4, 6)
	_100_v11 = _nw1
	var _102_globalState *GlobalState
	_ = _102_globalState
	var _nw2 *GlobalState = New_GlobalState_()
	_ = _nw2
	_nw2.Ctor__(true, _89_v1, _dafny.SeqOf(_90_v2), _dafny.Companion_Sequence_.Concatenate(_91_v3, _91_v3), false, _dafny.SeqOf(_90_v2), _dafny.IntOfInt64(301), true, _dafny.IntOfInt64(794), _dafny.IntOfInt64(616), false, _dafny.IntOfInt64(333), _dafny.IntOfInt64(349), _dafny.SetOf(_92_v4), _dafny.IntOfInt64(578), (_97_v9).Select((Companion_Default___.SafeIndex(_92_v4, _dafny.IntOfUint32((_97_v9).Cardinality()))).Uint32()).(_dafny.MultiSet), _dafny.IntOfInt64(982), _dafny.IntOfInt64(-565), _98_v10, _100_v11, _95_v7, _98_v10)
	_102_globalState = _nw2
	for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_98_v10), 0))); ; {
		_guard_loop_0, _ok24 := _iter24()
		if !_ok24 {
			break
		}
		var _103_i2 _dafny.Int
		_103_i2 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_103_i2).Sign() != -1) && ((_103_i2).Cmp(_dafny.ArrayLen((_98_v10), 0)) < 0)) {
			(_98_v10).ArraySet1(_90_v2, (_103_i2).Int())
		}
	}
	var _104_v12 _dafny.Sequence
	_ = _104_v12
	_104_v12 = _dafny.SeqOf(_98_v10, _98_v10)
	var _105_v13 *C4
	_ = _105_v13
	var _nw3 *C4 = New_C4_()
	_ = _nw3
	_nw3.Ctor__(_90_v2, _104_v12)
	_105_v13 = _nw3
	var _106_v14 D2
	_ = _106_v14
	_106_v14 = Companion_D2_.Create_DC8_(Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("ufigqlgw")), _100_v11, _90_v2)
	var _107_v15 _dafny.CodePoint
	_ = _107_v15
	_107_v15 = _dafny.CodePoint('m')
	var _108_v16 _dafny.Map
	_ = _108_v16
	_108_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v15, _90_v2)
	var _109_v17 _dafny.Set
	_ = _109_v17
	_109_v17 = _dafny.SetOf((_106_v14).Dtor_cf16(), (func() bool {
		if (_108_v16).Contains(_107_v15) {
			return (_108_v16).Get(_107_v15).(bool)
		}
		return _90_v2
	})(), _90_v2)
	var _110_v18 _dafny.Sequence
	_ = _110_v18
	_110_v18 = _dafny.SeqOf(_dafny.IntOfInt64(-999), _92_v4)
	var _111_v19 _dafny.Map
	_ = _111_v19
	_111_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_109_v17).Contains(_105_v13.F28), (_105_v13).Fm13(_dafny.Companion_Sequence_.Update(_110_v18, (Companion_Default___.SafeIndex(_92_v4, _dafny.IntOfUint32((_110_v18).Cardinality()))).Uint32(), _dafny.IntOfInt64(731)), _90_v2, _dafny.IntOfUint32((_110_v18).Cardinality()), _102_globalState))
	var _112_v20 _dafny.Array
	_ = _112_v20
	var _nw4 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
	_ = _nw4
	_112_v20 = _nw4
	var _113_v21 D11
	_ = _113_v21
	_113_v21 = Companion_D11_.Create_DC37_(_112_v20)
	var _114_v22 _dafny.Map
	_ = _114_v22
	_114_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v21, _90_v2)
	_111_v19 = (_111_v19).Update((func() bool {
		if (_114_v22).Contains(_113_v21) {
			return (_114_v22).Get(_113_v21).(bool)
		}
		return _90_v2
	})(), !((_105_v13.F28) && (_90_v2)))
	var _115_v23 _dafny.Set
	_ = _115_v23
	_115_v23 = _dafny.SetOf(_dafny.IntOfUint32((Companion_Default___.Fm6(_102_globalState)).Cardinality()), _92_v4, _92_v4, _92_v4, _92_v4)
	var _116_v24 D17
	_ = _116_v24
	_116_v24 = Companion_D17_.Create_DC55_(_115_v23)
	var _pat_let_tv0 = _92_v4
	_ = _pat_let_tv0
	var _pat_let_tv1 = _105_v13
	_ = _pat_let_tv1
	var _pat_let_tv2 = _93_v5
	_ = _pat_let_tv2
	var _pat_let_tv3 = _92_v4
	_ = _pat_let_tv3
	var _pat_let_tv4 = _92_v4
	_ = _pat_let_tv4
	var _pat_let_tv5 = _92_v4
	_ = _pat_let_tv5
	var _source3 D5 = func(_source4 D17) D5 {
		if _source4.Is_DC56() {
			var _117___mcc_h6 bool = _source4.Get_().(D17_DC56).Cf87
			_ = _117___mcc_h6
			var _118_cf87 bool = _117___mcc_h6
			_ = _118_cf87
			return Companion_D5_.Create_DC20_(Companion_D5_.Create_DC18_(_pat_let_tv0, _dafny.IntOfInt64(280)))
		} else {
			var _119___mcc_h7 _dafny.Set = _source4.Get_().(D17_DC55).Cf86
			_ = _119___mcc_h7
			var _120_cf86 _dafny.Set = _119___mcc_h7
			_ = _120_cf86
			if _pat_let_tv1.F28 {
				return Companion_D5_.Create_DC20_(Companion_D5_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv2, _pat_let_tv3)))
			} else {
				return Companion_D5_.Create_DC20_(Companion_D5_.Create_DC18_(_pat_let_tv4, _pat_let_tv5))
			}
		}
	}(_116_v24)
	_ = _source3
	if _source3.Is_DC18() {
		var _121___mcc_h0 _dafny.Int = _source3.Get_().(D5_DC18).Cf30
		_ = _121___mcc_h0
		var _122___mcc_h1 _dafny.Int = _source3.Get_().(D5_DC18).Cf31
		_ = _122___mcc_h1
		var _123_cf31 _dafny.Int = _122___mcc_h1
		_ = _123_cf31
		var _124_cf30 _dafny.Int = _121___mcc_h0
		_ = _124_cf30
		var _125_v25 _dafny.Map
		_ = _125_v25
		_125_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v2, _124_cf30)
		_125_v25 = _125_v25
		var _126_v26 *C4
		_ = _126_v26
		var _nw5 *C4 = New_C4_()
		_ = _nw5
		_nw5.Ctor__(_90_v2, _104_v12)
		_126_v26 = _nw5
		var _127_v27 _dafny.Array
		_ = _127_v27
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
		_ = _nw6
		_127_v27 = _nw6
		var _128_v28 _dafny.Map
		_ = _128_v28
		_128_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v13.F28, _127_v27)
		var _129_v29 D10
		_ = _129_v29
		_129_v29 = Companion_D10_.Create_DC34_(_107_v15, true, _105_v13.F28, !(_105_v13.F28), _126_v26.F28)
		var _130_v30 _dafny.Array
		_ = _130_v30
		var _nwElement0_1 _dafny.Array = (func() _dafny.Array {
			if (_128_v28).Contains((_129_v29).Dtor_cf57()) {
				return (_128_v28).Get((_129_v29).Dtor_cf57()).(_dafny.Array)
			}
			return _127_v27
		})()
		_ = _nwElement0_1
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(10))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_1, 0)
		(_nw7).ArraySet1(_127_v27, 1)
		(_nw7).ArraySet1(_127_v27, 2)
		(_nw7).ArraySet1(_127_v27, 3)
		(_nw7).ArraySet1(_127_v27, 4)
		(_nw7).ArraySet1(_127_v27, 5)
		(_nw7).ArraySet1((func() _dafny.Array {
			if _105_v13.F28 {
				return _127_v27
			}
			return _127_v27
		})(), 6)
		(_nw7).ArraySet1(_127_v27, 7)
		(_nw7).ArraySet1(_127_v27, 8)
		(_nw7).ArraySet1((func() _dafny.Array {
			if _90_v2 {
				return _127_v27
			}
			return _127_v27
		})(), 9)
		_130_v30 = _nw7
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_130_v30), 0))
		_ = _index0
		(_130_v30).ArraySet1(_127_v27, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_130_v30), 0))
		_ = _index1
		(_130_v30).ArraySet1(_127_v27, (_index1).Int())
		var _131_v31 D6
		_ = _131_v31
		_131_v31 = Companion_D6_.Create_DC22_(_dafny.IntOfInt64(590))
		var _132_v32 D6
		_ = _132_v32
		_132_v32 = Companion_D6_.Create_DC24_(_131_v31)
		var _133_v33 _dafny.Map
		_ = _133_v33
		_133_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v11, _132_v32)
		_133_v33 = _133_v33
	} else if _source3.Is_DC19() {
		var _134___mcc_h2 _dafny.Map = _source3.Get_().(D5_DC19).Cf32
		_ = _134___mcc_h2
		var _135___mcc_h3 bool = _source3.Get_().(D5_DC19).Cf33
		_ = _135___mcc_h3
		var _136_cf33 bool = _135___mcc_h3
		_ = _136_cf33
		var _137_cf32 _dafny.Map = _134___mcc_h2
		_ = _137_cf32
		(_102_globalState).F11 = _92_v4
		var _138_v34 *C5
		_ = _138_v34
		var _nw8 *C5 = New_C5_()
		_ = _nw8
		_nw8.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v4, _105_v13.F28), _92_v4)
		_138_v34 = _nw8
		var _139_v35 _dafny.Sequence
		_ = _139_v35
		_139_v35 = _dafny.SeqOf(_115_v23)
		(_102_globalState).F11 = (_92_v4).Minus((_dafny.Zero).Minus((_105_v13).Fm0(_95_v7, (_dafny.SetOf((_138_v34).F26())).Cardinality(), _139_v35, _105_v13.F28, _102_globalState)))
		var _140_v36 _dafny.Array
		_ = _140_v36
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_1
		var _nw9 _dafny.Array
		_ = _nw9
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw9 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Map = (func(_141_v13 *C4, _142_v34 *C5) func(_dafny.Int) _dafny.Map {
				return func(_143_i3 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v13.F28, (_142_v34).F27())
				}
			})(_105_v13, _138_v34)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw9 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw9).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw9).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_140_v36 = _nw9
		_140_v36 = _140_v36
	} else if _source3.Is_DC17() {
		var _144___mcc_h4 _dafny.Map = _source3.Get_().(D5_DC17).Cf29
		_ = _144___mcc_h4
		var _145_cf29 _dafny.Map = _144___mcc_h4
		_ = _145_cf29
		var _146_v37 _dafny.Array
		_ = _146_v37
		var _nw10 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
		_ = _nw10
		_146_v37 = _nw10
		(_102_globalState).F11 = ((func() _dafny.Int {
			if _105_v13.F28 {
				return _92_v4
			}
			return _92_v4
		})()).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v13.F28, _146_v37)).Cardinality())
		var _147_v38 _dafny.Array
		_ = _147_v38
		var _nwElement0_2 _dafny.Sequence = _93_v5
		_ = _nwElement0_2
		var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.One)
		_ = _nw11
		(_nw11).ArraySet1(_nwElement0_2, 0)
		_147_v38 = _nw11
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_147_v38), 0))
		_ = _index2
		(_147_v38).ArraySet1(_93_v5, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_147_v38), 0))
		_ = _index3
		(_147_v38).ArraySet1(_93_v5, (_index3).Int())
		(_105_v13).F28 = _105_v13.F28
		var _148_v39 bool
		_ = _148_v39
		var _149_v40 _dafny.Array
		_ = _149_v40
		var _150_v41 _dafny.Map
		_ = _150_v41
		var _151_v42 _dafny.Int
		_ = _151_v42
		var _out0 bool
		_ = _out0
		var _out1 _dafny.Array
		_ = _out1
		var _out2 _dafny.Map
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out0, _out1, _out2, _out3 = (_105_v13).M0(_105_v13.F28, _102_globalState)
		_148_v39 = _out0
		_149_v40 = _out1
		_150_v41 = _out2
		_151_v42 = _out3
	} else {
		var _152___mcc_h5 D5 = _source3.Get_().(D5_DC20).Cf34
		_ = _152___mcc_h5
		var _153_cf34 D5 = _152___mcc_h5
		_ = _153_cf34
		var _154_v43 _dafny.Map
		_ = _154_v43
		_154_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v10, _92_v4)
		var _155_v44 D14
		_ = _155_v44
		_155_v44 = Companion_D14_.Create_DC48_(_92_v4, _154_v43, _105_v13.F28)
		var _source5 D14 = _155_v44
		_ = _source5
		if _source5.Is_DC48() {
			var _156___mcc_h8 _dafny.Int = _source5.Get_().(D14_DC48).Cf78
			_ = _156___mcc_h8
			var _157___mcc_h9 _dafny.Map = _source5.Get_().(D14_DC48).Cf79
			_ = _157___mcc_h9
			var _158___mcc_h10 bool = _source5.Get_().(D14_DC48).Cf80
			_ = _158___mcc_h10
			var _159_cf80 bool = _158___mcc_h10
			_ = _159_cf80
			var _160_cf79 _dafny.Map = _157___mcc_h9
			_ = _160_cf79
			var _161_cf78 _dafny.Int = _156___mcc_h8
			_ = _161_cf78
			(_105_v13).F28 = !((_105_v13).Fm13(_dafny.Companion_Sequence_.Concatenate(_110_v18, _110_v18), _105_v13.F28, _dafny.IntOfInt64(201), _102_globalState))
			var _rhs0 bool = _105_v13.F28
			_ = _rhs0
			var _rhs1 bool = _105_v13.F28
			_ = _rhs1
			var _rhs2 bool = (_161_cf78).Cmp((func() _dafny.Int {
				if _159_cf80 {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_93_v5).Cardinality()))
				}
				return _161_cf78
			})()) == 0
			_ = _rhs2
			var _lhs0 *C4 = _105_v13
			_ = _lhs0
			var _lhs1 *C4 = _105_v13
			_ = _lhs1
			var _lhs2 *GlobalState = _102_globalState
			_ = _lhs2
			_lhs0.F28 = _rhs0
			_lhs1.F28 = _rhs1
			_lhs2.F10 = _rhs2
			_93_v5 = _dafny.Companion_Sequence_.Concatenate(_93_v5, _dafny.Companion_Sequence_.Update(_93_v5, (Companion_Default___.SafeIndex((_115_v23).Cardinality(), _dafny.IntOfUint32((_93_v5).Cardinality()))).Uint32(), _107_v15))
			var _162_v45 bool
			_ = _162_v45
			var _163_v46 _dafny.Array
			_ = _163_v46
			var _164_v47 _dafny.Map
			_ = _164_v47
			var _165_v48 _dafny.Int
			_ = _165_v48
			var _out4 bool
			_ = _out4
			var _out5 _dafny.Array
			_ = _out5
			var _out6 _dafny.Map
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out4, _out5, _out6, _out7 = (_105_v13).M0(_105_v13.F28, _102_globalState)
			_162_v45 = _out4
			_163_v46 = _out5
			_164_v47 = _out6
			_165_v48 = _out7
		} else {
			var _166___mcc_h11 _dafny.MultiSet = _source5.Get_().(D14_DC47).Cf77
			_ = _166___mcc_h11
			var _167_cf77 _dafny.MultiSet = _166___mcc_h11
			_ = _167_cf77
			var _168_v49 _dafny.Map
			_ = _168_v49
			_168_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D3_.Create_DC11_((_91_v3).Select((Companion_Default___.SafeIndex(_92_v4, _dafny.IntOfUint32((_91_v3).Cardinality()))).Uint32()).(bool), _105_v13.F28)).Dtor_cf19(), (_dafny.Zero).Minus(_92_v4))
			var _169_v50 _dafny.Sequence
			_ = _169_v50
			_169_v50 = _dafny.SeqOf(_115_v23)
			var _170_v51 _dafny.Sequence
			_ = _170_v51
			_170_v51 = _dafny.SeqOf(_169_v50)
			_168_v49 = (_168_v49).Update(true, (_105_v13).Fm0(_95_v7, _92_v4, (_170_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.IntOfUint32((_170_v51).Cardinality()))).Uint32()).(_dafny.Sequence), _105_v13.F28, _102_globalState))
			var _171_v52 *C5
			_ = _171_v52
			var _nw12 *C5 = New_C5_()
			_ = _nw12
			_nw12.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v4, _105_v13.F28), _92_v4)
			_171_v52 = _nw12
			var _172_v53 *C2
			_ = _172_v53
			var _nw13 *C2 = New_C2_()
			_ = _nw13
			_nw13.Ctor__((Companion_D12_.Create_DC42_(_92_v4)).Dtor_cf69(), _104_v12)
			_172_v53 = _nw13
			var _173_v54 _dafny.Set
			_ = _173_v54
			_173_v54 = _dafny.SetOf(_93_v5, _93_v5)
			var _174_v56 T0
			_ = _174_v56
			var _nw14 *C7 = New_C7_()
			_ = _nw14
			_nw14.Ctor__(_115_v23, _dafny.SeqOf(_98_v10, _98_v10, _98_v10, _98_v10, (_104_v12).Select((Companion_Default___.SafeIndex((_167_cf77).Cardinality(), _dafny.IntOfUint32((_104_v12).Cardinality()))).Uint32()).(_dafny.Array)))
			_174_v56 = _nw14
			var _175_v57 _dafny.Map
			_ = _175_v57
			_175_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v56, _173_v54)
			var _176_v58 D13
			_ = _176_v58
			_176_v58 = Companion_D13_.Create_DC46_((_171_v52).Fm5(!(_105_v13.F28), _102_globalState), _174_v56, _92_v4, _dafny.SetOf(_100_v11, _100_v11))
			var _177_v59 _dafny.Map
			_ = _177_v59
			_177_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v5, _105_v13.F28)
			var _178_v62 _dafny.Sequence
			_ = _178_v62
			_178_v62 = _dafny.SeqOf(_173_v54, _173_v54)
			var _179_v63 _dafny.Array
			_ = _179_v63
			var _nwElement0_3 _dafny.Set = _173_v54
			_ = _nwElement0_3
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(15))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_3, 0)
			(_nw15).ArraySet1(_173_v54, 1)
			(_nw15).ArraySet1((func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter25 := _dafny.Iterate((_173_v54).Elements()); ; {
					_compr_24, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _180_v55 _dafny.Sequence
					_180_v55 = interface{}(_compr_24).(_dafny.Sequence)
					if (_173_v54).Contains(_180_v55) {
						_coll24.Add(_180_v55)
					}
				}
				return _coll24.ToSet()
			}()).Difference((func() _dafny.Set {
				if (_175_v57).Contains((_176_v58).Dtor_cf74()) {
					return (_175_v57).Get((_176_v58).Dtor_cf74()).(_dafny.Set)
				}
				return _173_v54
			})()), 2)
			(_nw15).ArraySet1(_173_v54, 3)
			(_nw15).ArraySet1(_173_v54, 4)
			(_nw15).ArraySet1(func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate((_177_v59).Keys().Elements()); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _181_v60 _dafny.Sequence
					_181_v60 = interface{}(_compr_25).(_dafny.Sequence)
					if (_177_v59).Contains(_181_v60) {
						_coll25.Add(_181_v60)
					}
				}
				return _coll25.ToSet()
			}(), 5)
			(_nw15).ArraySet1(_173_v54, 6)
			(_nw15).ArraySet1((Companion_Default___.Fm55(_105_v13.F28, _dafny.IntOfUint32((_91_v3).Cardinality()), _107_v15, _102_globalState)).Union(_173_v54), 7)
			(_nw15).ArraySet1(_dafny.SetOf(_93_v5, _93_v5, _93_v5), 8)
			(_nw15).ArraySet1(_173_v54, 9)
			(_nw15).ArraySet1(_173_v54, 10)
			(_nw15).ArraySet1((_173_v54).Intersection(_173_v54), 11)
			(_nw15).ArraySet1(Companion_Default___.Fm55(_105_v13.F28, _dafny.IntOfUint32((_93_v5).Cardinality()), _107_v15, _102_globalState), 12)
			(_nw15).ArraySet1((func() _dafny.Set {
				var _coll26 = _dafny.NewBuilder()
				_ = _coll26
				for _iter27 := _dafny.Iterate((_173_v54).Elements()); ; {
					_compr_26, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _182_v61 _dafny.Sequence
					_182_v61 = interface{}(_compr_26).(_dafny.Sequence)
					if (_173_v54).Contains(_182_v61) {
						_coll26.Add(_182_v61)
					}
				}
				return _coll26.ToSet()
			}()).Union((_178_v62).Select((Companion_Default___.SafeIndex((_172_v53).F29(), _dafny.IntOfUint32((_178_v62).Cardinality()))).Uint32()).(_dafny.Set)), 13)
			(_nw15).ArraySet1(_173_v54, 14)
			_179_v63 = _nw15
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_179_v63), 0))
			_ = _index4
			(_179_v63).ArraySet1((_173_v54).Union(_173_v54), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_179_v63), 0))
			_ = _index5
			(_179_v63).ArraySet1((_173_v54).Difference(_173_v54), (_index5).Int())
		}
		var _183_v64 _dafny.Array
		_ = _183_v64
		var _nwElement0_4 _dafny.CodePoint = Companion_Default___.Fm38(_90_v2, _92_v4, _102_globalState)
		_ = _nwElement0_4
		var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
		_ = _nw16
		(_nw16).ArraySet1CodePoint(_nwElement0_4, 0)
		(_nw16).ArraySet1CodePoint(_107_v15, 1)
		(_nw16).ArraySet1CodePoint(_107_v15, 2)
		(_nw16).ArraySet1CodePoint(_107_v15, 3)
		(_nw16).ArraySet1CodePoint(_107_v15, 4)
		(_nw16).ArraySet1CodePoint(_107_v15, 5)
		(_nw16).ArraySet1CodePoint(_107_v15, 6)
		(_nw16).ArraySet1CodePoint(_dafny.CodePoint('n'), 7)
		(_nw16).ArraySet1CodePoint(_107_v15, 8)
		(_nw16).ArraySet1CodePoint(_107_v15, 9)
		(_nw16).ArraySet1CodePoint(_107_v15, 10)
		(_nw16).ArraySet1CodePoint(_107_v15, 11)
		_183_v64 = _nw16
		var _184_v65 _dafny.Sequence
		_ = _184_v65
		_184_v65 = _dafny.SeqOf(_183_v64)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_98_v10), 0))
		_ = _index6
		(_98_v10).ArraySet1(!_dafny.Companion_Sequence_.Equal(_184_v65, _184_v65), (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_98_v10), 0))
		_ = _index7
		var _rhs3 bool = (_dafny.IntOfUint32((_93_v5).Cardinality())).Cmp((_92_v4).Plus(_dafny.IntOfUint32((_93_v5).Cardinality()))) == 0
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _93_v5
		_ = _rhs4
		var _rhs5 bool = (_109_v17).Equals((_109_v17).Union(_109_v17))
		_ = _rhs5
		var _rhs6 _dafny.Int = _92_v4
		_ = _rhs6
		var _rhs7 _dafny.CodePoint = _107_v15
		_ = _rhs7
		var _lhs3 *GlobalState = _102_globalState
		_ = _lhs3
		var _lhs4 _dafny.Array = _98_v10
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_98_v10), 0))
		_ = _lhs5
		var _lhs6 *GlobalState = _102_globalState
		_ = _lhs6
		_lhs3.F0 = _rhs3
		_93_v5 = _rhs4
		(_lhs4).ArraySet1(_rhs5, (_lhs5).Int())
		_lhs6.F16 = _rhs6
		_107_v15 = _rhs7
		var _185_v66 *C3
		_ = _185_v66
		var _nw17 *C3 = New_C3_()
		_ = _nw17
		_nw17.Ctor__(_104_v12)
		_185_v66 = _nw17
		_185_v66 = (func() *C3 {
			if (_98_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_98_v10), 0))).Int()).(bool) {
				return _185_v66
			}
			return _185_v66
		})()
		var _186_v67 _dafny.Array
		_ = _186_v67
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_2
		var _nw18 _dafny.Array
		_ = _nw18
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw18 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) D15 = func(_187_i4 _dafny.Int) D15 {
				return Companion_D15_.Create_DC52_()
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw18 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw18).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw18).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_186_v67 = _nw18
		var _188_v68 D15
		_ = _188_v68
		_188_v68 = Companion_D15_.Create_DC52_()
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_186_v67), 0))
		_ = _index8
		(_186_v67).ArraySet1(_188_v68, (_index8).Int())
		var _189_v69 _dafny.Array
		_ = _189_v69
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_3
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_190_v13 *C4) func(_dafny.Int) bool {
				return func(_191_i5 _dafny.Int) bool {
					return _190_v13.F28
				}
			})(_105_v13)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw19 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw19).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw19).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_189_v69 = _nw19
		var _192_v70 *C2
		_ = _192_v70
		var _nw20 *C2 = New_C2_()
		_ = _nw20
		_nw20.Ctor__((_dafny.Zero).Minus(_92_v4), _dafny.SeqOf(_189_v69, _189_v69))
		_192_v70 = _nw20
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_186_v67), 0))
		_ = _index9
		var _rhs8 bool = ((_88_v0).Cardinality()).Cmp((_192_v70).F29()) >= 0
		_ = _rhs8
		var _rhs9 D15 = _188_v68
		_ = _rhs9
		var _rhs10 *C2 = _192_v70
		_ = _rhs10
		var _lhs7 *C4 = _105_v13
		_ = _lhs7
		var _lhs8 _dafny.Array = _186_v67
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_186_v67), 0))
		_ = _lhs9
		_lhs7.F28 = _rhs8
		(_lhs8).ArraySet1(_rhs9, (_lhs9).Int())
		_192_v70 = _rhs10
	}
	var _193_v71 D18
	_ = _193_v71
	_193_v71 = Companion_D18_.Create_DC57_(_105_v13)
	var _194_v72 _dafny.Map
	_ = _194_v72
	_194_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v2, _105_v13)
	var _195_v73 _dafny.Array
	_ = _195_v73
	var _nwElement0_5 *C4 = (_193_v71).Dtor_cf88()
	_ = _nwElement0_5
	var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(13))
	_ = _nw21
	(_nw21).ArraySet1(_nwElement0_5, 0)
	(_nw21).ArraySet1(_105_v13, 1)
	(_nw21).ArraySet1((func() *C4 {
		if (_194_v72).Contains(_90_v2) {
			return (_194_v72).Get(_90_v2).(*C4)
		}
		return _105_v13
	})(), 2)
	(_nw21).ArraySet1(_105_v13, 3)
	(_nw21).ArraySet1(_105_v13, 4)
	(_nw21).ArraySet1(_105_v13, 5)
	(_nw21).ArraySet1(_105_v13, 6)
	(_nw21).ArraySet1(_105_v13, 7)
	(_nw21).ArraySet1(_105_v13, 8)
	(_nw21).ArraySet1((_193_v71).Dtor_cf88(), 9)
	(_nw21).ArraySet1(_105_v13, 10)
	(_nw21).ArraySet1(_105_v13, 11)
	(_nw21).ArraySet1(_105_v13, 12)
	_195_v73 = _nw21
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_195_v73), 0))
	_ = _index10
	(_195_v73).ArraySet1(_105_v13, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_195_v73), 0))
	_ = _index11
	(_195_v73).ArraySet1(_105_v13, (_index11).Int())
	if (_92_v4).Cmp(_92_v4) == 0 {
		var _196_v74 _dafny.Map
		_ = _196_v74
		_196_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v13.F28, _92_v4)
		var _197_v75 _dafny.Sequence
		_ = _197_v75
		_197_v75 = _dafny.SeqOf(_93_v5)
		_196_v74 = (_196_v74).Update(_90_v2, (_dafny.IntOfUint32(((_197_v75).Select((Companion_Default___.SafeIndex(_92_v4, _dafny.IntOfUint32((_197_v75).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Plus(_92_v4))
		var _198_v76 *C3
		_ = _198_v76
		var _nw22 *C3 = New_C3_()
		_ = _nw22
		_nw22.Ctor__(_dafny.SeqOf(_98_v10, _98_v10, _98_v10, _98_v10, _98_v10))
		_198_v76 = _nw22
		var _199_v77 _dafny.Map
		_ = _199_v77
		_199_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_115_v23).Cardinality(), _98_v10)
		var _200_v78 _dafny.Map
		_ = _200_v78
		_200_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if (_199_v77).Contains(_92_v4) {
				return (_199_v77).Get(_92_v4).(_dafny.Array)
			}
			return _98_v10
		})(), _100_v11)
		_200_v78 = (_200_v78).Update(_98_v10, _100_v11)
		(_102_globalState).F16 = Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_93_v5).Cardinality())).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v4, _90_v2)).Cardinality()), _92_v4)
		var _201_v79 _dafny.Sequence
		_ = _201_v79
		_201_v79 = _dafny.SeqOf(_109_v17, _109_v17, _dafny.SetOf(_105_v13.F28, _105_v13.F28))
		var _202_v80 D12
		_ = _202_v80
		_202_v80 = Companion_D12_.Create_DC42_((_115_v23).Cardinality())
		var _203_v81 _dafny.Map
		_ = _203_v81
		_203_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_202_v80, _109_v17)
		var _204_v82 _dafny.Array
		_ = _204_v82
		var _nwElement0_6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_201_v79, _dafny.SeqOf((func() _dafny.Set {
			if (_203_v81).Contains(Companion_D12_.Create_DC42_(_92_v4)) {
				return (_203_v81).Get(Companion_D12_.Create_DC42_(_92_v4)).(_dafny.Set)
			}
			return _109_v17
		})(), _109_v17))
		_ = _nwElement0_6
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_6, 0)
		_204_v82 = _nw23
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_204_v82), 0))
		_ = _index12
		(_204_v82).ArraySet1(_201_v79, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_204_v82), 0))
		_ = _index13
		(_204_v82).ArraySet1(_201_v79, (_index13).Int())
	} else {
		var _205_v83 *C0
		_ = _205_v83
		var _nw24 *C0 = New_C0_()
		_ = _nw24
		_nw24.Ctor__()
		_205_v83 = _nw24
		var _206_v84 _dafny.MultiSet
		_ = _206_v84
		_206_v84 = _dafny.MultiSetOf(_205_v83)
		var _207_v85 _dafny.Map
		_ = _207_v85
		_207_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(113), Companion_D12_.Create_DC41_(_206_v84))
		_207_v85 = _207_v85
		var _208_v86 *C3
		_ = _208_v86
		var _nw25 *C3 = New_C3_()
		_ = _nw25
		_nw25.Ctor__(_104_v12)
		_208_v86 = _nw25
		var _rhs11 _dafny.Int = (Companion_Default___.SafeDivisionInt(_92_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_93_v5, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-282), _dafny.IntOfUint32((_93_v5).Cardinality()))).Uint32(), _107_v15)).Cardinality()))).Times(_dafny.IntOfInt64(208))
		_ = _rhs11
		var _rhs12 *C3 = _208_v86
		_ = _rhs12
		_92_v4 = _rhs11
		_208_v86 = _rhs12
		_93_v5 = _93_v5
		var _209_v87 _dafny.Map
		_ = _209_v87
		_209_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v3, _92_v4)
		var _210_v88 _dafny.Sequence
		_ = _210_v88
		_210_v88 = _dafny.SeqOf(_209_v87, _209_v87, _209_v87)
		(_102_globalState).F8 = (_dafny.IntOfInt64(287)).Times(((_210_v88).Select((Companion_Default___.SafeIndex(_92_v4, _dafny.IntOfUint32((_210_v88).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_98_v10), 0))
		_ = _index14
		(_98_v10).ArraySet1(Companion_Default___.Fm15(_90_v2, _102_globalState), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_98_v10), 0))
		_ = _index15
		(_98_v10).ArraySet1(_105_v13.F28, (_index15).Int())
	}
	var _211_v89 bool
	_ = _211_v89
	var _212_v90 _dafny.Array
	_ = _212_v90
	var _213_v91 _dafny.Map
	_ = _213_v91
	var _214_v92 _dafny.Int
	_ = _214_v92
	var _out8 bool
	_ = _out8
	var _out9 _dafny.Array
	_ = _out9
	var _out10 _dafny.Map
	_ = _out10
	var _out11 _dafny.Int
	_ = _out11
	_out8, _out9, _out10, _out11 = (_105_v13).M0(_90_v2, _102_globalState)
	_211_v89 = _out8
	_212_v90 = _out9
	_213_v91 = _out10
	_214_v92 = _out11
	var _hi0 _dafny.Int = _dafny.IntOfInt64(465)
	_ = _hi0
	for _215_i6 := (_dafny.Zero).Minus(_92_v4); _215_i6.Cmp(_hi0) < 0; _215_i6 = _215_i6.Plus(_dafny.One) {
		var _216_v93 D12
		_ = _216_v93
		_216_v93 = Companion_D12_.Create_DC42_(_92_v4)
		_216_v93 = func(_pat_let0_0 D12) D12 {
			return func(_217_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let1_0 _dafny.Int) D12 {
					return func(_218_dt__update_hcf69_h0 _dafny.Int) D12 {
						return Companion_D12_.Create_DC42_(_218_dt__update_hcf69_h0)
					}(_pat_let1_0)
				}(_215_i6)
			}(_pat_let0_0)
		}(_216_v93)
		if true {
			var _219_v94 *C1
			_ = _219_v94
			var _nw26 *C1 = New_C1_()
			_ = _nw26
			_nw26.Ctor__(_214_v92, _104_v12)
			_219_v94 = _nw26
			var _220_v95 _dafny.Sequence
			_ = _220_v95
			_220_v95 = _dafny.SeqOf(_219_v94, _219_v94)
			_211_v89 = ((_dafny.Zero).Minus(_92_v4)).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_220_v95).Cardinality()), _92_v4)) <= 0
			var _221_v96 *C4
			_ = _221_v96
			var _nw27 *C4 = New_C4_()
			_ = _nw27
			_nw27.Ctor__(_211_v89, _dafny.Companion_Sequence_.Concatenate(_104_v12, _104_v12))
			_221_v96 = _nw27
			_88_v0 = _88_v0
			var _222_v97 _dafny.Array
			_ = _222_v97
			var _nw28 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw28
			_222_v97 = _nw28
			var _223_v98 _dafny.Map
			_ = _223_v98
			_223_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_i6, true)
			var _224_v99 T0
			_ = _224_v99
			var _nw29 *C1 = New_C1_()
			_ = _nw29
			_nw29.Ctor__((_223_v98).Cardinality(), _104_v12)
			_224_v99 = _nw29
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_222_v97), 0))
			_ = _index16
			(_222_v97).ArraySet1(_224_v99, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_222_v97), 0))
			_ = _index17
			(_222_v97).ArraySet1(_224_v99, (_index17).Int())
			var _225_v100 _dafny.Array
			_ = _225_v100
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
			_ = _nw30
			_225_v100 = _nw30
			var _226_v101 _dafny.MultiSet
			_ = _226_v101
			_226_v101 = _dafny.MultiSetOf(_225_v100)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_98_v10), 0))
			_ = _index18
			(_98_v10).ArraySet1((_226_v101).IsProperSubsetOf(_dafny.MultiSetOf(_225_v100, _225_v100)), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_98_v10), 0))
			_ = _index19
			(_98_v10).ArraySet1(!(_105_v13.F28), (_index19).Int())
		} else {
			(_102_globalState).F0 = _105_v13.F28
			_110_v18 = _110_v18
			var _227_v102 _dafny.Map
			_ = _227_v102
			_227_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_214_v92, _215_i6)
			var _228_v103 D5
			_ = _228_v103
			_228_v103 = Companion_D5_.Create_DC19_(_227_v102, _90_v2)
			var _229_v104 bool
			_ = _229_v104
			var _230_v105 _dafny.Array
			_ = _230_v105
			var _231_v106 _dafny.Map
			_ = _231_v106
			var _232_v107 _dafny.Int
			_ = _232_v107
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Array
			_ = _out13
			var _out14 _dafny.Map
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			_out12, _out13, _out14, _out15 = (_105_v13).M0((func() bool {
				if _90_v2 {
					return !(true)
				}
				return (_228_v103).Dtor_cf33()
			})(), _102_globalState)
			_229_v104 = _out12
			_230_v105 = _out13
			_231_v106 = _out14
			_232_v107 = _out15
			var _233_v108 _dafny.Map
			_ = _233_v108
			_233_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_232_v107, _90_v2)
			var _234_v109 *C5
			_ = _234_v109
			var _nw31 *C5 = New_C5_()
			_ = _nw31
			_nw31.Ctor__(_233_v108, _92_v4)
			_234_v109 = _nw31
			var _235_v110 bool
			_ = _235_v110
			var _236_v111 _dafny.Array
			_ = _236_v111
			var _237_v112 _dafny.Map
			_ = _237_v112
			var _238_v113 _dafny.Int
			_ = _238_v113
			var _out16 bool
			_ = _out16
			var _out17 _dafny.Array
			_ = _out17
			var _out18 _dafny.Map
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			_out16, _out17, _out18, _out19 = (_105_v13).M0((Companion_Default___.Fm56(_232_v107, _102_globalState)).Dtor_cf1(), _102_globalState)
			_235_v110 = _out16
			_236_v111 = _out17
			_237_v112 = _out18
			_238_v113 = _out19
		}
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_100_v11), 0))
		_ = _index20
		(_100_v11).ArraySet1(_215_i6, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_100_v11), 0))
		_ = _index21
		(_100_v11).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(628), (_92_v4).Plus(_215_i6)), (_index21).Int())
		var _239_v114 *C7
		_ = _239_v114
		var _nw32 *C7 = New_C7_()
		_ = _nw32
		_nw32.Ctor__(_115_v23, _dafny.SeqOf(_98_v10))
		_239_v114 = _nw32
	}
	if !((_91_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.IntOfUint32((_91_v3).Cardinality()))).Uint32()).(bool)) {
		_93_v5 = _dafny.Companion_Sequence_.Update(_93_v5, (Companion_Default___.SafeIndex((_214_v92).Times(_dafny.IntOfInt64(-876)), _dafny.IntOfUint32((_93_v5).Cardinality()))).Uint32(), _107_v15)
		var _240_v115 bool
		_ = _240_v115
		var _241_v116 _dafny.Array
		_ = _241_v116
		var _242_v117 _dafny.Map
		_ = _242_v117
		var _243_v118 _dafny.Int
		_ = _243_v118
		var _out20 bool
		_ = _out20
		var _out21 _dafny.Array
		_ = _out21
		var _out22 _dafny.Map
		_ = _out22
		var _out23 _dafny.Int
		_ = _out23
		_out20, _out21, _out22, _out23 = (_105_v13).M0(_211_v89, _102_globalState)
		_240_v115 = _out20
		_241_v116 = _out21
		_242_v117 = _out22
		_243_v118 = _out23
		var _244_v119 bool
		_ = _244_v119
		var _245_v120 _dafny.Array
		_ = _245_v120
		var _246_v121 _dafny.Map
		_ = _246_v121
		var _247_v122 _dafny.Int
		_ = _247_v122
		var _out24 bool
		_ = _out24
		var _out25 _dafny.Array
		_ = _out25
		var _out26 _dafny.Map
		_ = _out26
		var _out27 _dafny.Int
		_ = _out27
		_out24, _out25, _out26, _out27 = (_105_v13).M0((_95_v7).IsProperSubsetOf(_95_v7), _102_globalState)
		_244_v119 = _out24
		_245_v120 = _out25
		_246_v121 = _out26
		_247_v122 = _out27
		var _248_v123 _dafny.Set
		_ = _248_v123
		_248_v123 = _dafny.SetOf(_91_v3, _91_v3, _91_v3)
		_240_v115 = !((_248_v123).Union(_dafny.SetOf(_91_v3, _91_v3))).Equals(_248_v123)
		var _249_v124 D17
		_ = _249_v124
		_249_v124 = Companion_D17_.Create_DC56_(_90_v2)
		(_102_globalState).F0 = (((_249_v124).Dtor_cf87()) || (_90_v2)) || (_105_v13.F28)
	} else {
		var _250_v125 _dafny.Set
		_ = _250_v125
		_250_v125 = _dafny.SetOf(_111_v19, (_111_v19).Merge(_111_v19))
		_250_v125 = _250_v125
		var _251_v126 _dafny.Map
		_ = _251_v126
		_251_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if !(_105_v13.F28) {
				return _105_v13.F28
			}
			return _211_v89
		})(), Companion_Default___.Fm18(_90_v2, _102_globalState))
		_251_v126 = _251_v126
		var _252_v127 _dafny.Map
		_ = _252_v127
		_252_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(475), _214_v92)
		var _253_v128 _dafny.Map
		_ = _253_v128
		_253_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_214_v92, _252_v127)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_100_v11), 0))
		_ = _index22
		(_100_v11).ArraySet1(((_252_v127).Merge((func() _dafny.Map {
			if (_253_v128).Contains(_214_v92) {
				return (_253_v128).Get(_214_v92).(_dafny.Map)
			}
			return _252_v127
		})())).Cardinality(), (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_100_v11), 0))
		_ = _index23
		(_100_v11).ArraySet1(_dafny.IntOfInt64(170), (_index23).Int())
		var _254_v129 *C4
		_ = _254_v129
		var _nw33 *C4 = New_C4_()
		_ = _nw33
		_nw33.Ctor__(_211_v89, _dafny.Companion_Sequence_.Concatenate(_104_v12, _104_v12))
		_254_v129 = _nw33
		(_102_globalState).F8 = Companion_Default___.Fm7((_115_v23).Cardinality(), _102_globalState)
	}
	var _255_v130 bool
	_ = _255_v130
	var _256_v131 _dafny.Array
	_ = _256_v131
	var _257_v132 _dafny.Map
	_ = _257_v132
	var _258_v133 _dafny.Int
	_ = _258_v133
	var _out28 bool
	_ = _out28
	var _out29 _dafny.Array
	_ = _out29
	var _out30 _dafny.Map
	_ = _out30
	var _out31 _dafny.Int
	_ = _out31
	_out28, _out29, _out30, _out31 = (_105_v13).M0(_211_v89, _102_globalState)
	_255_v130 = _out28
	_256_v131 = _out29
	_257_v132 = _out30
	_258_v133 = _out31
	_93_v5 = _93_v5
	var _259_v134 *C3
	_ = _259_v134
	var _nw34 *C3 = New_C3_()
	_ = _nw34
	_nw34.Ctor__(_104_v12)
	_259_v134 = _nw34
	_90_v2 = (_92_v4).Cmp((_214_v92).Times((_dafny.MultiSetOf(_259_v134)).Cardinality())) >= 0
	var _260_i7 _dafny.Int
	_ = _260_i7
	_260_i7 = _dafny.Zero
	{
		for !(_105_v13.F28) {
			{
				if (_260_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_260_i7 = (_260_i7).Plus(_dafny.One)
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_100_v11), 0))
				_ = _index24
				(_100_v11).ArraySet1(_92_v4, (_index24).Int())
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_100_v11), 0))
				_ = _index25
				(_100_v11).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xxxok")).Cardinality())).Times(Companion_Default___.SafeModuloInt(_258_v133, _258_v133)), (_index25).Int())
				(_102_globalState).F11 = (_100_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_100_v11), 0))).Int()).(_dafny.Int)
				var _261_v135 _dafny.Sequence
				_ = _261_v135
				_261_v135 = _dafny.SeqOf(_115_v23, _115_v23)
				(_102_globalState).F0 = (_dafny.IntOfInt64(318)).Cmp((_105_v13).Fm0(_94_v6, (_100_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_100_v11), 0))).Int()).(_dafny.Int), _261_v135, _105_v13.F28, _102_globalState)) == 0
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_195_v73), 0))
				_ = _index26
				(_195_v73).ArraySet1((_195_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_195_v73), 0))).Int()).(*C4), (_index26).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _262_v136 _dafny.Array
	_ = _262_v136
	var _263_v137 _dafny.Int
	_ = _263_v137
	var _out32 _dafny.Array
	_ = _out32
	var _out33 _dafny.Int
	_ = _out33
	_out32, _out33 = (_259_v134).M4(_102_globalState)
	_262_v136 = _out32
	_263_v137 = _out33
	var _264_v138 _dafny.Map
	_ = _264_v138
	_264_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(188))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg32 _dafny.Int) interface{} {
			return coer32(arg32)
		}
	}((func(_265_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_266_i8 _dafny.Int) _dafny.CodePoint {
			return _265_v15
		}
	})(_107_v15)))).Cardinality()), _211_v89)
	var _267_v139 *C5
	_ = _267_v139
	var _nw35 *C5 = New_C5_()
	_ = _nw35
	_nw35.Ctor__(_264_v138, (_92_v4).Times((_dafny.Zero).Minus(_258_v133)))
	_267_v139 = _nw35
	var _268_v140 _dafny.MultiSet
	_ = _268_v140
	_268_v140 = _dafny.MultiSetOf(_107_v15, _dafny.CodePoint('d'))
	var _269_v141 _dafny.Map
	_ = _269_v141
	_269_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v140, _214_v92)
	_269_v141 = (_269_v141).Update((_268_v140).Update(_107_v15, Companion_Default___.Abs((_267_v139).F27())), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_110_v18).Cardinality()), _263_v137))
	_dafny.Print((_88_v0).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_89_v1).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_90_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_91_v3, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_92_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_93_v5.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_94_v6).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(476))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_95_v7).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_96_v8, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_97_v9, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_v10).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v11).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState.F1).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_102_globalState).F2(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_102_globalState).F3(), _dafny.SeqOf(true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_102_globalState.F5, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState.F13).Equals(_dafny.SetOf(_dafny.IntOfInt64(166))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F15()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F19()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F20()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_102_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_104_v12).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_v13.F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf14()).Dtor_cf6().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v14).Dtor_cf15()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_106_v14).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_107_v15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_108_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v17).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_110_v18, _dafny.SeqOf(_dafny.IntOfInt64(-999), _dafny.IntOfInt64(166))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v22).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_115_v23).Equals(_dafny.SetOf(_dafny.IntOfInt64(9), _dafny.IntOfInt64(166))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_v24).Dtor_cf86()).Equals(_dafny.SetOf(_dafny.IntOfInt64(9), _dafny.IntOfInt64(166))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_193_v71).Dtor_cf88().F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_193_v71).Dtor_cf88()).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_194_v72).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.Zero).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.Zero).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.One).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.One).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v73).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(*C4).F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_195_v73).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(*C4)).F22()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_211_v89)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_213_v91).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(509), _dafny.IntOfInt64(508), _dafny.IntOfInt64(508), _dafny.IntOfInt64(508), _dafny.IntOfInt64(509)), _dafny.CodePoint('l'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_214_v92)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_255_v130)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v131).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v132).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(509), _dafny.IntOfInt64(508), _dafny.IntOfInt64(508), _dafny.IntOfInt64(508), _dafny.IntOfInt64(509)), _dafny.CodePoint('l'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_258_v133)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_260_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v136).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_263_v137)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v138).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(188), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_267_v139).F26()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(188), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v139).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v140).Equals(_dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v141).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('d')), _dafny.IntOfInt64(4)).UpdateUnsafe(_dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('d')), _dafny.Zero)))
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
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_() D0 {
	return D0{D0_DC0{}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf0 _dafny.Int
	Cf1 bool
	Cf2 _dafny.Map
	Cf3 _dafny.Int
	Cf4 _dafny.CodePoint
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf0 _dafny.Int, Cf1 bool, Cf2 _dafny.Map, Cf3 _dafny.Int, Cf4 _dafny.CodePoint) D0 {
	return D0{D0_DC1{Cf0, Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_()
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf0
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Map {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			_, ok := other.Get_().(D0_DC0)
			return ok
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1 == data2.Cf1 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
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

type D1_DC3 struct {
	Cf6 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf5 _dafny.Sequence
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf5 _dafny.Sequence) D1 {
	return D1{D1_DC2{Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC4 struct {
	Cf7 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 D1) D1 {
	return D1{D1_DC4{Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.EmptySeq)
}

func (_this D1) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf7() D1 {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf6.VerbatimString(true) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + data.Cf5.VerbatimString(true) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf9 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf9 _dafny.Int) D2 {
	return D2{D2_DC6{Cf9}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf10 _dafny.Sequence
	Cf11 bool
	Cf12 bool
	Cf13 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf10 _dafny.Sequence, Cf11 bool, Cf12 bool, Cf13 _dafny.Int) D2 {
	return D2{D2_DC7{Cf10, Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf14 D1
	Cf15 _dafny.Array
	Cf16 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf14 D1, Cf15 _dafny.Array, Cf16 bool) D2 {
	return D2{D2_DC8{Cf14, Cf15, Cf16}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC5 struct {
	Cf8 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 _dafny.Array) D2 {
	return D2{D2_DC5{Cf8}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC9 struct {
	Cf17 D2
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf17 D2) D2 {
	return D2{D2_DC9{Cf17}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero)
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() D1 {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf16() bool {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf17() D2 {
	return _this.Get_().(D2_DC9).Cf17
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf17) + ")"
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
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf10.Equals(data2.Cf10) && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12 && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf14.Equals(data2.Cf14) && data1.Cf15 == data2.Cf15 && data1.Cf16 == data2.Cf16
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf8 == data2.Cf8
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf17.Equals(data2.Cf17)
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

type D3_DC11 struct {
	Cf19 bool
	Cf20 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf19 bool, Cf20 bool) D3 {
	return D3{D3_DC11{Cf19, Cf20}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf18 _dafny.Map
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf18 _dafny.Map) D3 {
	return D3{D3_DC10{Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_(false, false)
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D4_DC13 struct {
	Cf22 D0
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf22 D0) D4 {
	return D4{D4_DC13{Cf22}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf23 _dafny.CodePoint
	Cf24 _dafny.Array
	Cf25 _dafny.Sequence
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf23 _dafny.CodePoint, Cf24 _dafny.Array, Cf25 _dafny.Sequence) D4 {
	return D4{D4_DC14{Cf23, Cf24, Cf25}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC15 struct {
	Cf26 _dafny.Int
	Cf27 _dafny.Int
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf26 _dafny.Int, Cf27 _dafny.Int) D4 {
	return D4{D4_DC15{Cf26, Cf27}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC12 struct {
	Cf21 _dafny.Map
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf21 _dafny.Map) D4 {
	return D4{D4_DC12{Cf21}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC16 struct {
	Cf28 D4
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf28 D4) D4 {
	return D4{D4_DC16{Cf28}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(Companion_D0_.Default())
}

func (_this D4) Dtor_cf22() D0 {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() _dafny.CodePoint {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D4_DC14).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf26
}

func (_this D4) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf27
}

func (_this D4) Dtor_cf21() _dafny.Map {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf28() D4 {
	return _this.Get_().(D4_DC16).Cf28
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + data.Cf25.VerbatimString(true) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24 && data1.Cf25.Equals(data2.Cf25)
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
			return ok && data1.Cf28.Equals(data2.Cf28)
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

type D5_DC18 struct {
	Cf30 _dafny.Int
	Cf31 _dafny.Int
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf30 _dafny.Int, Cf31 _dafny.Int) D5 {
	return D5{D5_DC18{Cf30, Cf31}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC19 struct {
	Cf32 _dafny.Map
	Cf33 bool
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf32 _dafny.Map, Cf33 bool) D5 {
	return D5{D5_DC19{Cf32, Cf33}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC17 struct {
	Cf29 _dafny.Map
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf29 _dafny.Map) D5 {
	return D5{D5_DC17{Cf29}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC20 struct {
	Cf34 D5
}

func (D5_DC20) isD5() {}

func (CompanionStruct_D5_) Create_DC20_(Cf34 D5) D5 {
	return D5{D5_DC20{Cf34}}
}

func (_this D5) Is_DC20() bool {
	_, ok := _this.Get_().(D5_DC20)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC18_(_dafny.Zero, _dafny.Zero)
}

func (_this D5) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf31
}

func (_this D5) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D5_DC19).Cf32
}

func (_this D5) Dtor_cf33() bool {
	return _this.Get_().(D5_DC19).Cf33
}

func (_this D5) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) Dtor_cf34() D5 {
	return _this.Get_().(D5_DC20).Cf34
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC20:
		{
			return "D5.DC20" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf32.Equals(data2.Cf32) && data1.Cf33 == data2.Cf33
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D5_DC20:
		{
			data2, ok := other.Get_().(D5_DC20)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D6_DC22 struct {
	Cf36 _dafny.Int
}

func (D6_DC22) isD6() {}

func (CompanionStruct_D6_) Create_DC22_(Cf36 _dafny.Int) D6 {
	return D6{D6_DC22{Cf36}}
}

func (_this D6) Is_DC22() bool {
	_, ok := _this.Get_().(D6_DC22)
	return ok
}

type D6_DC23 struct {
	Cf37 bool
	Cf38 _dafny.Int
	Cf39 bool
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf37 bool, Cf38 _dafny.Int, Cf39 bool) D6 {
	return D6{D6_DC23{Cf37, Cf38, Cf39}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

type D6_DC21 struct {
	Cf35 _dafny.MultiSet
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf35 _dafny.MultiSet) D6 {
	return D6{D6_DC21{Cf35}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC24 struct {
	Cf40 D6
}

func (D6_DC24) isD6() {}

func (CompanionStruct_D6_) Create_DC24_(Cf40 D6) D6 {
	return D6{D6_DC24{Cf40}}
}

func (_this D6) Is_DC24() bool {
	_, ok := _this.Get_().(D6_DC24)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC22_(_dafny.Zero)
}

func (_this D6) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D6_DC22).Cf36
}

func (_this D6) Dtor_cf37() bool {
	return _this.Get_().(D6_DC23).Cf37
}

func (_this D6) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D6_DC23).Cf38
}

func (_this D6) Dtor_cf39() bool {
	return _this.Get_().(D6_DC23).Cf39
}

func (_this D6) Dtor_cf35() _dafny.MultiSet {
	return _this.Get_().(D6_DC21).Cf35
}

func (_this D6) Dtor_cf40() D6 {
	return _this.Get_().(D6_DC24).Cf40
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC22:
		{
			return "D6.DC22" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D6_DC24:
		{
			return "D6.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC22:
		{
			data2, ok := other.Get_().(D6_DC22)
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0
		}
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39
		}
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	case D6_DC24:
		{
			data2, ok := other.Get_().(D6_DC24)
			return ok && data1.Cf40.Equals(data2.Cf40)
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

type D7_DC26 struct {
	Cf42 _dafny.Sequence
	Cf43 _dafny.Int
	Cf44 _dafny.Int
}

func (D7_DC26) isD7() {}

func (CompanionStruct_D7_) Create_DC26_(Cf42 _dafny.Sequence, Cf43 _dafny.Int, Cf44 _dafny.Int) D7 {
	return D7{D7_DC26{Cf42, Cf43, Cf44}}
}

func (_this D7) Is_DC26() bool {
	_, ok := _this.Get_().(D7_DC26)
	return ok
}

type D7_DC25 struct {
	Cf41 _dafny.Set
}

func (D7_DC25) isD7() {}

func (CompanionStruct_D7_) Create_DC25_(Cf41 _dafny.Set) D7 {
	return D7{D7_DC25{Cf41}}
}

func (_this D7) Is_DC25() bool {
	_, ok := _this.Get_().(D7_DC25)
	return ok
}

type D7_DC27 struct {
	Cf45 D7
}

func (D7_DC27) isD7() {}

func (CompanionStruct_D7_) Create_DC27_(Cf45 D7) D7 {
	return D7{D7_DC27{Cf45}}
}

func (_this D7) Is_DC27() bool {
	_, ok := _this.Get_().(D7_DC27)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC26_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D7) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D7_DC26).Cf42
}

func (_this D7) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D7_DC26).Cf43
}

func (_this D7) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D7_DC26).Cf44
}

func (_this D7) Dtor_cf41() _dafny.Set {
	return _this.Get_().(D7_DC25).Cf41
}

func (_this D7) Dtor_cf45() D7 {
	return _this.Get_().(D7_DC27).Cf45
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC26:
		{
			return "D7.DC26" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D7_DC25:
		{
			return "D7.DC25" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D7_DC27:
		{
			return "D7.DC27" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC26:
		{
			data2, ok := other.Get_().(D7_DC26)
			return ok && data1.Cf42.Equals(data2.Cf42) && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Cmp(data2.Cf44) == 0
		}
	case D7_DC25:
		{
			data2, ok := other.Get_().(D7_DC25)
			return ok && data1.Cf41.Equals(data2.Cf41)
		}
	case D7_DC27:
		{
			data2, ok := other.Get_().(D7_DC27)
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D8_DC29 struct {
	Cf47 bool
	Cf48 bool
	Cf49 _dafny.CodePoint
}

func (D8_DC29) isD8() {}

func (CompanionStruct_D8_) Create_DC29_(Cf47 bool, Cf48 bool, Cf49 _dafny.CodePoint) D8 {
	return D8{D8_DC29{Cf47, Cf48, Cf49}}
}

func (_this D8) Is_DC29() bool {
	_, ok := _this.Get_().(D8_DC29)
	return ok
}

type D8_DC28 struct {
	Cf46 _dafny.Sequence
}

func (D8_DC28) isD8() {}

func (CompanionStruct_D8_) Create_DC28_(Cf46 _dafny.Sequence) D8 {
	return D8{D8_DC28{Cf46}}
}

func (_this D8) Is_DC28() bool {
	_, ok := _this.Get_().(D8_DC28)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC29_(false, false, _dafny.CodePoint('D'))
}

func (_this D8) Dtor_cf47() bool {
	return _this.Get_().(D8_DC29).Cf47
}

func (_this D8) Dtor_cf48() bool {
	return _this.Get_().(D8_DC29).Cf48
}

func (_this D8) Dtor_cf49() _dafny.CodePoint {
	return _this.Get_().(D8_DC29).Cf49
}

func (_this D8) Dtor_cf46() _dafny.Sequence {
	return _this.Get_().(D8_DC28).Cf46
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC29:
		{
			return "D8.DC29" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D8_DC28:
		{
			return "D8.DC28" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC29:
		{
			data2, ok := other.Get_().(D8_DC29)
			return ok && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48 && data1.Cf49 == data2.Cf49
		}
	case D8_DC28:
		{
			data2, ok := other.Get_().(D8_DC28)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D9_DC31 struct {
	Cf51 _dafny.Int
	Cf52 _dafny.Array
}

func (D9_DC31) isD9() {}

func (CompanionStruct_D9_) Create_DC31_(Cf51 _dafny.Int, Cf52 _dafny.Array) D9 {
	return D9{D9_DC31{Cf51, Cf52}}
}

func (_this D9) Is_DC31() bool {
	_, ok := _this.Get_().(D9_DC31)
	return ok
}

type D9_DC30 struct {
	Cf50 _dafny.Map
}

func (D9_DC30) isD9() {}

func (CompanionStruct_D9_) Create_DC30_(Cf50 _dafny.Map) D9 {
	return D9{D9_DC30{Cf50}}
}

func (_this D9) Is_DC30() bool {
	_, ok := _this.Get_().(D9_DC30)
	return ok
}

type D9_DC32 struct {
	Cf53 D9
}

func (D9_DC32) isD9() {}

func (CompanionStruct_D9_) Create_DC32_(Cf53 D9) D9 {
	return D9{D9_DC32{Cf53}}
}

func (_this D9) Is_DC32() bool {
	_, ok := _this.Get_().(D9_DC32)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC31_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D9) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D9_DC31).Cf51
}

func (_this D9) Dtor_cf52() _dafny.Array {
	return _this.Get_().(D9_DC31).Cf52
}

func (_this D9) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D9_DC30).Cf50
}

func (_this D9) Dtor_cf53() D9 {
	return _this.Get_().(D9_DC32).Cf53
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC31:
		{
			return "D9.DC31" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D9_DC30:
		{
			return "D9.DC30" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D9_DC32:
		{
			return "D9.DC32" + "(" + _dafny.String(data.Cf53) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC31:
		{
			data2, ok := other.Get_().(D9_DC31)
			return ok && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52
		}
	case D9_DC30:
		{
			data2, ok := other.Get_().(D9_DC30)
			return ok && data1.Cf50.Equals(data2.Cf50)
		}
	case D9_DC32:
		{
			data2, ok := other.Get_().(D9_DC32)
			return ok && data1.Cf53.Equals(data2.Cf53)
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

type D10_DC34 struct {
	Cf55 _dafny.CodePoint
	Cf56 bool
	Cf57 bool
	Cf58 bool
	Cf59 bool
}

func (D10_DC34) isD10() {}

func (CompanionStruct_D10_) Create_DC34_(Cf55 _dafny.CodePoint, Cf56 bool, Cf57 bool, Cf58 bool, Cf59 bool) D10 {
	return D10{D10_DC34{Cf55, Cf56, Cf57, Cf58, Cf59}}
}

func (_this D10) Is_DC34() bool {
	_, ok := _this.Get_().(D10_DC34)
	return ok
}

type D10_DC35 struct {
	Cf60 _dafny.Int
	Cf61 _dafny.Int
}

func (D10_DC35) isD10() {}

func (CompanionStruct_D10_) Create_DC35_(Cf60 _dafny.Int, Cf61 _dafny.Int) D10 {
	return D10{D10_DC35{Cf60, Cf61}}
}

func (_this D10) Is_DC35() bool {
	_, ok := _this.Get_().(D10_DC35)
	return ok
}

type D10_DC33 struct {
	Cf54 _dafny.Map
}

func (D10_DC33) isD10() {}

func (CompanionStruct_D10_) Create_DC33_(Cf54 _dafny.Map) D10 {
	return D10{D10_DC33{Cf54}}
}

func (_this D10) Is_DC33() bool {
	_, ok := _this.Get_().(D10_DC33)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC34_(_dafny.CodePoint('D'), false, false, false, false)
}

func (_this D10) Dtor_cf55() _dafny.CodePoint {
	return _this.Get_().(D10_DC34).Cf55
}

func (_this D10) Dtor_cf56() bool {
	return _this.Get_().(D10_DC34).Cf56
}

func (_this D10) Dtor_cf57() bool {
	return _this.Get_().(D10_DC34).Cf57
}

func (_this D10) Dtor_cf58() bool {
	return _this.Get_().(D10_DC34).Cf58
}

func (_this D10) Dtor_cf59() bool {
	return _this.Get_().(D10_DC34).Cf59
}

func (_this D10) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D10_DC35).Cf60
}

func (_this D10) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D10_DC35).Cf61
}

func (_this D10) Dtor_cf54() _dafny.Map {
	return _this.Get_().(D10_DC33).Cf54
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC34:
		{
			return "D10.DC34" + "(" + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D10_DC35:
		{
			return "D10.DC35" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D10_DC33:
		{
			return "D10.DC33" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC34:
		{
			data2, ok := other.Get_().(D10_DC34)
			return ok && data1.Cf55 == data2.Cf55 && data1.Cf56 == data2.Cf56 && data1.Cf57 == data2.Cf57 && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59
		}
	case D10_DC35:
		{
			data2, ok := other.Get_().(D10_DC35)
			return ok && data1.Cf60.Cmp(data2.Cf60) == 0 && data1.Cf61.Cmp(data2.Cf61) == 0
		}
	case D10_DC33:
		{
			data2, ok := other.Get_().(D10_DC33)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D11_DC37 struct {
	Cf63 _dafny.Array
}

func (D11_DC37) isD11() {}

func (CompanionStruct_D11_) Create_DC37_(Cf63 _dafny.Array) D11 {
	return D11{D11_DC37{Cf63}}
}

func (_this D11) Is_DC37() bool {
	_, ok := _this.Get_().(D11_DC37)
	return ok
}

type D11_DC38 struct {
	Cf64 _dafny.Int
	Cf65 _dafny.Int
}

func (D11_DC38) isD11() {}

func (CompanionStruct_D11_) Create_DC38_(Cf64 _dafny.Int, Cf65 _dafny.Int) D11 {
	return D11{D11_DC38{Cf64, Cf65}}
}

func (_this D11) Is_DC38() bool {
	_, ok := _this.Get_().(D11_DC38)
	return ok
}

type D11_DC39 struct {
	Cf66 bool
}

func (D11_DC39) isD11() {}

func (CompanionStruct_D11_) Create_DC39_(Cf66 bool) D11 {
	return D11{D11_DC39{Cf66}}
}

func (_this D11) Is_DC39() bool {
	_, ok := _this.Get_().(D11_DC39)
	return ok
}

type D11_DC36 struct {
	Cf62 _dafny.Sequence
}

func (D11_DC36) isD11() {}

func (CompanionStruct_D11_) Create_DC36_(Cf62 _dafny.Sequence) D11 {
	return D11{D11_DC36{Cf62}}
}

func (_this D11) Is_DC36() bool {
	_, ok := _this.Get_().(D11_DC36)
	return ok
}

type D11_DC40 struct {
	Cf67 D11
}

func (D11_DC40) isD11() {}

func (CompanionStruct_D11_) Create_DC40_(Cf67 D11) D11 {
	return D11{D11_DC40{Cf67}}
}

func (_this D11) Is_DC40() bool {
	_, ok := _this.Get_().(D11_DC40)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC37_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D11) Dtor_cf63() _dafny.Array {
	return _this.Get_().(D11_DC37).Cf63
}

func (_this D11) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D11_DC38).Cf64
}

func (_this D11) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D11_DC38).Cf65
}

func (_this D11) Dtor_cf66() bool {
	return _this.Get_().(D11_DC39).Cf66
}

func (_this D11) Dtor_cf62() _dafny.Sequence {
	return _this.Get_().(D11_DC36).Cf62
}

func (_this D11) Dtor_cf67() D11 {
	return _this.Get_().(D11_DC40).Cf67
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC37:
		{
			return "D11.DC37" + "(" + _dafny.String(data.Cf63) + ")"
		}
	case D11_DC38:
		{
			return "D11.DC38" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D11_DC39:
		{
			return "D11.DC39" + "(" + _dafny.String(data.Cf66) + ")"
		}
	case D11_DC36:
		{
			return "D11.DC36" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D11_DC40:
		{
			return "D11.DC40" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC37:
		{
			data2, ok := other.Get_().(D11_DC37)
			return ok && data1.Cf63 == data2.Cf63
		}
	case D11_DC38:
		{
			data2, ok := other.Get_().(D11_DC38)
			return ok && data1.Cf64.Cmp(data2.Cf64) == 0 && data1.Cf65.Cmp(data2.Cf65) == 0
		}
	case D11_DC39:
		{
			data2, ok := other.Get_().(D11_DC39)
			return ok && data1.Cf66 == data2.Cf66
		}
	case D11_DC36:
		{
			data2, ok := other.Get_().(D11_DC36)
			return ok && data1.Cf62.Equals(data2.Cf62)
		}
	case D11_DC40:
		{
			data2, ok := other.Get_().(D11_DC40)
			return ok && data1.Cf67.Equals(data2.Cf67)
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

type D12_DC42 struct {
	Cf69 _dafny.Int
}

func (D12_DC42) isD12() {}

func (CompanionStruct_D12_) Create_DC42_(Cf69 _dafny.Int) D12 {
	return D12{D12_DC42{Cf69}}
}

func (_this D12) Is_DC42() bool {
	_, ok := _this.Get_().(D12_DC42)
	return ok
}

type D12_DC41 struct {
	Cf68 _dafny.MultiSet
}

func (D12_DC41) isD12() {}

func (CompanionStruct_D12_) Create_DC41_(Cf68 _dafny.MultiSet) D12 {
	return D12{D12_DC41{Cf68}}
}

func (_this D12) Is_DC41() bool {
	_, ok := _this.Get_().(D12_DC41)
	return ok
}

type D12_DC43 struct {
	Cf70 D12
}

func (D12_DC43) isD12() {}

func (CompanionStruct_D12_) Create_DC43_(Cf70 D12) D12 {
	return D12{D12_DC43{Cf70}}
}

func (_this D12) Is_DC43() bool {
	_, ok := _this.Get_().(D12_DC43)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC42_(_dafny.Zero)
}

func (_this D12) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D12_DC42).Cf69
}

func (_this D12) Dtor_cf68() _dafny.MultiSet {
	return _this.Get_().(D12_DC41).Cf68
}

func (_this D12) Dtor_cf70() D12 {
	return _this.Get_().(D12_DC43).Cf70
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC42:
		{
			return "D12.DC42" + "(" + _dafny.String(data.Cf69) + ")"
		}
	case D12_DC41:
		{
			return "D12.DC41" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D12_DC43:
		{
			return "D12.DC43" + "(" + _dafny.String(data.Cf70) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC42:
		{
			data2, ok := other.Get_().(D12_DC42)
			return ok && data1.Cf69.Cmp(data2.Cf69) == 0
		}
	case D12_DC41:
		{
			data2, ok := other.Get_().(D12_DC41)
			return ok && data1.Cf68.Equals(data2.Cf68)
		}
	case D12_DC43:
		{
			data2, ok := other.Get_().(D12_DC43)
			return ok && data1.Cf70.Equals(data2.Cf70)
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

type D13_DC45 struct {
	Cf72 _dafny.Int
}

func (D13_DC45) isD13() {}

func (CompanionStruct_D13_) Create_DC45_(Cf72 _dafny.Int) D13 {
	return D13{D13_DC45{Cf72}}
}

func (_this D13) Is_DC45() bool {
	_, ok := _this.Get_().(D13_DC45)
	return ok
}

type D13_DC46 struct {
	Cf73 bool
	Cf74 T0
	Cf75 _dafny.Int
	Cf76 _dafny.Set
}

func (D13_DC46) isD13() {}

func (CompanionStruct_D13_) Create_DC46_(Cf73 bool, Cf74 T0, Cf75 _dafny.Int, Cf76 _dafny.Set) D13 {
	return D13{D13_DC46{Cf73, Cf74, Cf75, Cf76}}
}

func (_this D13) Is_DC46() bool {
	_, ok := _this.Get_().(D13_DC46)
	return ok
}

type D13_DC44 struct {
	Cf71 _dafny.Sequence
}

func (D13_DC44) isD13() {}

func (CompanionStruct_D13_) Create_DC44_(Cf71 _dafny.Sequence) D13 {
	return D13{D13_DC44{Cf71}}
}

func (_this D13) Is_DC44() bool {
	_, ok := _this.Get_().(D13_DC44)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC45_(_dafny.Zero)
}

func (_this D13) Dtor_cf72() _dafny.Int {
	return _this.Get_().(D13_DC45).Cf72
}

func (_this D13) Dtor_cf73() bool {
	return _this.Get_().(D13_DC46).Cf73
}

func (_this D13) Dtor_cf74() T0 {
	return _this.Get_().(D13_DC46).Cf74
}

func (_this D13) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D13_DC46).Cf75
}

func (_this D13) Dtor_cf76() _dafny.Set {
	return _this.Get_().(D13_DC46).Cf76
}

func (_this D13) Dtor_cf71() _dafny.Sequence {
	return _this.Get_().(D13_DC44).Cf71
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC45:
		{
			return "D13.DC45" + "(" + _dafny.String(data.Cf72) + ")"
		}
	case D13_DC46:
		{
			return "D13.DC46" + "(" + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D13_DC44:
		{
			return "D13.DC44" + "(" + _dafny.String(data.Cf71) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC45:
		{
			data2, ok := other.Get_().(D13_DC45)
			return ok && data1.Cf72.Cmp(data2.Cf72) == 0
		}
	case D13_DC46:
		{
			data2, ok := other.Get_().(D13_DC46)
			return ok && data1.Cf73 == data2.Cf73 && _dafny.AreEqual(data1.Cf74, data2.Cf74) && data1.Cf75.Cmp(data2.Cf75) == 0 && data1.Cf76.Equals(data2.Cf76)
		}
	case D13_DC44:
		{
			data2, ok := other.Get_().(D13_DC44)
			return ok && data1.Cf71.Equals(data2.Cf71)
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

type D14_DC48 struct {
	Cf78 _dafny.Int
	Cf79 _dafny.Map
	Cf80 bool
}

func (D14_DC48) isD14() {}

func (CompanionStruct_D14_) Create_DC48_(Cf78 _dafny.Int, Cf79 _dafny.Map, Cf80 bool) D14 {
	return D14{D14_DC48{Cf78, Cf79, Cf80}}
}

func (_this D14) Is_DC48() bool {
	_, ok := _this.Get_().(D14_DC48)
	return ok
}

type D14_DC47 struct {
	Cf77 _dafny.MultiSet
}

func (D14_DC47) isD14() {}

func (CompanionStruct_D14_) Create_DC47_(Cf77 _dafny.MultiSet) D14 {
	return D14{D14_DC47{Cf77}}
}

func (_this D14) Is_DC47() bool {
	_, ok := _this.Get_().(D14_DC47)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC48_(_dafny.Zero, _dafny.EmptyMap, false)
}

func (_this D14) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D14_DC48).Cf78
}

func (_this D14) Dtor_cf79() _dafny.Map {
	return _this.Get_().(D14_DC48).Cf79
}

func (_this D14) Dtor_cf80() bool {
	return _this.Get_().(D14_DC48).Cf80
}

func (_this D14) Dtor_cf77() _dafny.MultiSet {
	return _this.Get_().(D14_DC47).Cf77
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC48:
		{
			return "D14.DC48" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ")"
		}
	case D14_DC47:
		{
			return "D14.DC47" + "(" + _dafny.String(data.Cf77) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC48:
		{
			data2, ok := other.Get_().(D14_DC48)
			return ok && data1.Cf78.Cmp(data2.Cf78) == 0 && data1.Cf79.Equals(data2.Cf79) && data1.Cf80 == data2.Cf80
		}
	case D14_DC47:
		{
			data2, ok := other.Get_().(D14_DC47)
			return ok && data1.Cf77.Equals(data2.Cf77)
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

type D15_DC50 struct {
}

func (D15_DC50) isD15() {}

func (CompanionStruct_D15_) Create_DC50_() D15 {
	return D15{D15_DC50{}}
}

func (_this D15) Is_DC50() bool {
	_, ok := _this.Get_().(D15_DC50)
	return ok
}

type D15_DC51 struct {
	Cf82 bool
	Cf83 bool
}

func (D15_DC51) isD15() {}

func (CompanionStruct_D15_) Create_DC51_(Cf82 bool, Cf83 bool) D15 {
	return D15{D15_DC51{Cf82, Cf83}}
}

func (_this D15) Is_DC51() bool {
	_, ok := _this.Get_().(D15_DC51)
	return ok
}

type D15_DC52 struct {
}

func (D15_DC52) isD15() {}

func (CompanionStruct_D15_) Create_DC52_() D15 {
	return D15{D15_DC52{}}
}

func (_this D15) Is_DC52() bool {
	_, ok := _this.Get_().(D15_DC52)
	return ok
}

type D15_DC49 struct {
	Cf81 _dafny.Set
}

func (D15_DC49) isD15() {}

func (CompanionStruct_D15_) Create_DC49_(Cf81 _dafny.Set) D15 {
	return D15{D15_DC49{Cf81}}
}

func (_this D15) Is_DC49() bool {
	_, ok := _this.Get_().(D15_DC49)
	return ok
}

type D15_DC53 struct {
	Cf84 D15
}

func (D15_DC53) isD15() {}

func (CompanionStruct_D15_) Create_DC53_(Cf84 D15) D15 {
	return D15{D15_DC53{Cf84}}
}

func (_this D15) Is_DC53() bool {
	_, ok := _this.Get_().(D15_DC53)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC50_()
}

func (_this D15) Dtor_cf82() bool {
	return _this.Get_().(D15_DC51).Cf82
}

func (_this D15) Dtor_cf83() bool {
	return _this.Get_().(D15_DC51).Cf83
}

func (_this D15) Dtor_cf81() _dafny.Set {
	return _this.Get_().(D15_DC49).Cf81
}

func (_this D15) Dtor_cf84() D15 {
	return _this.Get_().(D15_DC53).Cf84
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC50:
		{
			return "D15.DC50"
		}
	case D15_DC51:
		{
			return "D15.DC51" + "(" + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ")"
		}
	case D15_DC52:
		{
			return "D15.DC52"
		}
	case D15_DC49:
		{
			return "D15.DC49" + "(" + _dafny.String(data.Cf81) + ")"
		}
	case D15_DC53:
		{
			return "D15.DC53" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC50:
		{
			_, ok := other.Get_().(D15_DC50)
			return ok
		}
	case D15_DC51:
		{
			data2, ok := other.Get_().(D15_DC51)
			return ok && data1.Cf82 == data2.Cf82 && data1.Cf83 == data2.Cf83
		}
	case D15_DC52:
		{
			_, ok := other.Get_().(D15_DC52)
			return ok
		}
	case D15_DC49:
		{
			data2, ok := other.Get_().(D15_DC49)
			return ok && data1.Cf81.Equals(data2.Cf81)
		}
	case D15_DC53:
		{
			data2, ok := other.Get_().(D15_DC53)
			return ok && data1.Cf84.Equals(data2.Cf84)
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

type D16_DC54 struct {
	Cf85 _dafny.Map
}

func (D16_DC54) isD16() {}

func (CompanionStruct_D16_) Create_DC54_(Cf85 _dafny.Map) D16 {
	return D16{D16_DC54{Cf85}}
}

func (_this D16) Is_DC54() bool {
	_, ok := _this.Get_().(D16_DC54)
	return ok
}

func (CompanionStruct_D16_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D16) Dtor_cf85() _dafny.Map {
	return _this.Get_().(D16_DC54).Cf85
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC54:
		{
			return "D16.DC54" + "(" + _dafny.String(data.Cf85) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC54:
		{
			data2, ok := other.Get_().(D16_DC54)
			return ok && data1.Cf85.Equals(data2.Cf85)
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

type D17_DC56 struct {
	Cf87 bool
}

func (D17_DC56) isD17() {}

func (CompanionStruct_D17_) Create_DC56_(Cf87 bool) D17 {
	return D17{D17_DC56{Cf87}}
}

func (_this D17) Is_DC56() bool {
	_, ok := _this.Get_().(D17_DC56)
	return ok
}

type D17_DC55 struct {
	Cf86 _dafny.Set
}

func (D17_DC55) isD17() {}

func (CompanionStruct_D17_) Create_DC55_(Cf86 _dafny.Set) D17 {
	return D17{D17_DC55{Cf86}}
}

func (_this D17) Is_DC55() bool {
	_, ok := _this.Get_().(D17_DC55)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC56_(false)
}

func (_this D17) Dtor_cf87() bool {
	return _this.Get_().(D17_DC56).Cf87
}

func (_this D17) Dtor_cf86() _dafny.Set {
	return _this.Get_().(D17_DC55).Cf86
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC56:
		{
			return "D17.DC56" + "(" + _dafny.String(data.Cf87) + ")"
		}
	case D17_DC55:
		{
			return "D17.DC55" + "(" + _dafny.String(data.Cf86) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC56:
		{
			data2, ok := other.Get_().(D17_DC56)
			return ok && data1.Cf87 == data2.Cf87
		}
	case D17_DC55:
		{
			data2, ok := other.Get_().(D17_DC55)
			return ok && data1.Cf86.Equals(data2.Cf86)
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

type D18_DC58 struct {
}

func (D18_DC58) isD18() {}

func (CompanionStruct_D18_) Create_DC58_() D18 {
	return D18{D18_DC58{}}
}

func (_this D18) Is_DC58() bool {
	_, ok := _this.Get_().(D18_DC58)
	return ok
}

type D18_DC57 struct {
	Cf88 *C4
}

func (D18_DC57) isD18() {}

func (CompanionStruct_D18_) Create_DC57_(Cf88 *C4) D18 {
	return D18{D18_DC57{Cf88}}
}

func (_this D18) Is_DC57() bool {
	_, ok := _this.Get_().(D18_DC57)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC58_()
}

func (_this D18) Dtor_cf88() *C4 {
	return _this.Get_().(D18_DC57).Cf88
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC58:
		{
			return "D18.DC58"
		}
	case D18_DC57:
		{
			return "D18.DC57" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC58:
		{
			_, ok := other.Get_().(D18_DC58)
			return ok
		}
	case D18_DC57:
		{
			data2, ok := other.Get_().(D18_DC57)
			return ok && data1.Cf88 == data2.Cf88
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

// Definition of datatype D19
type D19 struct {
	Data_D19_
}

func (_this D19) Get_() Data_D19_ {
	return _this.Data_D19_
}

type Data_D19_ interface {
	isD19()
}

type CompanionStruct_D19_ struct {
}

var Companion_D19_ = CompanionStruct_D19_{}

type D19_DC60 struct {
	Cf90 _dafny.Int
	Cf91 _dafny.Int
	Cf92 _dafny.Int
}

func (D19_DC60) isD19() {}

func (CompanionStruct_D19_) Create_DC60_(Cf90 _dafny.Int, Cf91 _dafny.Int, Cf92 _dafny.Int) D19 {
	return D19{D19_DC60{Cf90, Cf91, Cf92}}
}

func (_this D19) Is_DC60() bool {
	_, ok := _this.Get_().(D19_DC60)
	return ok
}

type D19_DC59 struct {
	Cf89 _dafny.Map
}

func (D19_DC59) isD19() {}

func (CompanionStruct_D19_) Create_DC59_(Cf89 _dafny.Map) D19 {
	return D19{D19_DC59{Cf89}}
}

func (_this D19) Is_DC59() bool {
	_, ok := _this.Get_().(D19_DC59)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC60_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D19) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D19_DC60).Cf90
}

func (_this D19) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D19_DC60).Cf91
}

func (_this D19) Dtor_cf92() _dafny.Int {
	return _this.Get_().(D19_DC60).Cf92
}

func (_this D19) Dtor_cf89() _dafny.Map {
	return _this.Get_().(D19_DC59).Cf89
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC60:
		{
			return "D19.DC60" + "(" + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ")"
		}
	case D19_DC59:
		{
			return "D19.DC59" + "(" + _dafny.String(data.Cf89) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC60:
		{
			data2, ok := other.Get_().(D19_DC60)
			return ok && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91.Cmp(data2.Cf91) == 0 && data1.Cf92.Cmp(data2.Cf92) == 0
		}
	case D19_DC59:
		{
			data2, ok := other.Get_().(D19_DC59)
			return ok && data1.Cf89.Equals(data2.Cf89)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D19) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D19)
	return ok && _this.Equals(typed)
}

func Type_D19_() _dafny.TypeDescriptor {
	return type_D19_{}
}

type type_D19_ struct {
}

func (_this type_D19_) Default() interface{} {
	return Companion_D19_.Default()
}

func (_this type_D19_) String() string {
	return "main.D19"
}
func (_this D19) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D19{}

// End of datatype D19

// Definition of datatype D20
type D20 struct {
	Data_D20_
}

func (_this D20) Get_() Data_D20_ {
	return _this.Data_D20_
}

type Data_D20_ interface {
	isD20()
}

type CompanionStruct_D20_ struct {
}

var Companion_D20_ = CompanionStruct_D20_{}

type D20_DC62 struct {
	Cf94 _dafny.Int
	Cf95 bool
	Cf96 _dafny.Int
}

func (D20_DC62) isD20() {}

func (CompanionStruct_D20_) Create_DC62_(Cf94 _dafny.Int, Cf95 bool, Cf96 _dafny.Int) D20 {
	return D20{D20_DC62{Cf94, Cf95, Cf96}}
}

func (_this D20) Is_DC62() bool {
	_, ok := _this.Get_().(D20_DC62)
	return ok
}

type D20_DC61 struct {
	Cf93 _dafny.Map
}

func (D20_DC61) isD20() {}

func (CompanionStruct_D20_) Create_DC61_(Cf93 _dafny.Map) D20 {
	return D20{D20_DC61{Cf93}}
}

func (_this D20) Is_DC61() bool {
	_, ok := _this.Get_().(D20_DC61)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC62_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D20) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D20_DC62).Cf94
}

func (_this D20) Dtor_cf95() bool {
	return _this.Get_().(D20_DC62).Cf95
}

func (_this D20) Dtor_cf96() _dafny.Int {
	return _this.Get_().(D20_DC62).Cf96
}

func (_this D20) Dtor_cf93() _dafny.Map {
	return _this.Get_().(D20_DC61).Cf93
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC62:
		{
			return "D20.DC62" + "(" + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ")"
		}
	case D20_DC61:
		{
			return "D20.DC61" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC62:
		{
			data2, ok := other.Get_().(D20_DC62)
			return ok && data1.Cf94.Cmp(data2.Cf94) == 0 && data1.Cf95 == data2.Cf95 && data1.Cf96.Cmp(data2.Cf96) == 0
		}
	case D20_DC61:
		{
			data2, ok := other.Get_().(D20_DC61)
			return ok && data1.Cf93.Equals(data2.Cf93)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D20) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D20)
	return ok && _this.Equals(typed)
}

func Type_D20_() _dafny.TypeDescriptor {
	return type_D20_{}
}

type type_D20_ struct {
}

func (_this type_D20_) Default() interface{} {
	return Companion_D20_.Default()
}

func (_this type_D20_) String() string {
	return "main.D20"
}
func (_this D20) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D20{}

// End of datatype D20

// Definition of trait T0
type T0 interface {
	String() string
	Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int
	M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int)
	F22() _dafny.Sequence
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
	F0   bool
	F1   _dafny.MultiSet
	F5   _dafny.Sequence
	F6   _dafny.Int
	F8   _dafny.Int
	F10  bool
	F11  _dafny.Int
	F13  _dafny.Set
	F16  _dafny.Int
	_f2  _dafny.Sequence
	_f3  _dafny.Sequence
	_f4  bool
	_f7  bool
	_f9  _dafny.Int
	_f12 _dafny.Int
	_f14 _dafny.Int
	_f15 _dafny.MultiSet
	_f17 _dafny.Int
	_f18 _dafny.Array
	_f19 _dafny.Array
	_f20 _dafny.MultiSet
	_f21 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F1 = _dafny.EmptyMultiSet
	_this.F5 = _dafny.EmptySeq
	_this.F6 = _dafny.Zero
	_this.F8 = _dafny.Zero
	_this.F10 = false
	_this.F11 = _dafny.Zero
	_this.F13 = _dafny.EmptySet
	_this.F16 = _dafny.Zero
	_this._f2 = _dafny.EmptySeq
	_this._f3 = _dafny.EmptySeq
	_this._f4 = false
	_this._f7 = false
	_this._f9 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.EmptyMultiSet
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f20 = _dafny.EmptyMultiSet
	_this._f21 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.MultiSet, f2 _dafny.Sequence, f3 _dafny.Sequence, f4 bool, f5 _dafny.Sequence, f6 _dafny.Int, f7 bool, f8 _dafny.Int, f9 _dafny.Int, f10 bool, f11 _dafny.Int, f12 _dafny.Int, f13 _dafny.Set, f14 _dafny.Int, f15 _dafny.MultiSet, f16 _dafny.Int, f17 _dafny.Int, f18 _dafny.Array, f19 _dafny.Array, f20 _dafny.MultiSet, f21 _dafny.Array) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *GlobalState) F2() _dafny.Sequence {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.MultiSet {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Array {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() _dafny.Array {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() _dafny.MultiSet {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() _dafny.Array {
	{
		return _this._f21
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
func (_this *C0) Fm16(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(false)))
	}
}
func (_this *C0) Fm17(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(712))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_270_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("kdpncg")), _dafny.UnicodeSeqOfUtf8Bytes("ujpfdfbpa"))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f22 _dafny.Sequence
	_f30 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f22 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
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

func (_this *C1) F22() _dafny.Sequence {
	return _this._f22
}
func (_this *C1) Ctor__(f30 _dafny.Int, f22 _dafny.Sequence) {
	{
		(_this)._f30 = f30
		(_this)._f22 = f22
	}
}
func (_this *C1) Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F30()
	}
}
func (_this *C1) Fm34(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(259)).Times((_this).F30())
	}
}
func (_this *C1) Fm35(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), (_this).F30())).Merge((func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter28 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_271_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))).Elements()); ; {
				_compr_27, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _272_v0 _dafny.CodePoint
				_272_v0 = interface{}(_compr_27).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}(func(_271_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				})), _272_v0) {
					_coll27.Add(_272_v0, (_this).F30())
				}
			}
			return _coll27.ToMap()
		}()).Merge(func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter29 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('n'))).Elements()); ; {
				_compr_28, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _273_v1 _dafny.CodePoint
				_273_v1 = interface{}(_compr_28).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('n')), _273_v1) {
					_coll28.Add(_273_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, true, true))).Cardinality(), false)).Cardinality())
				}
			}
			return _coll28.ToMap()
		}()))).Cardinality()
	}
}
func (_this *C1) M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		(globalState).F11 = _dafny.IntOfInt64(664)
		var _274_v0 _dafny.Array
		_ = _274_v0
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw36
		_274_v0 = _nw36
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))
		_ = _index27
		(_274_v0).ArraySet1(true, (_index27).Int())
		var _275_v1 D3
		_ = _275_v1
		_275_v1 = Companion_D3_.Create_DC11_(p0, p0)
		var _276_v2 _dafny.Map
		_ = _276_v2
		_276_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)).Cardinality())
		var _277_v3 _dafny.Map
		_ = _277_v3
		_277_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _276_v2)
		var _278_v4 _dafny.CodePoint
		_ = _278_v4
		_278_v4 = _dafny.CodePoint('o')
		var _279_v5 D0
		_ = _279_v5
		_279_v5 = Companion_D0_.Create_DC1_((_this).F30(), p0, _277_v3, (_this).F30(), _278_v4)
		var _280_v6 _dafny.Sequence
		_ = _280_v6
		_280_v6 = _dafny.SeqOf(p0, p0)
		var _pat_let_tv6 = p0
		_ = _pat_let_tv6
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))
		_ = _index28
		var _rhs13 bool = !(p0)
		_ = _rhs13
		var _rhs14 bool = func(_source6 D3) bool {
			if _source6.Is_DC11() {
				var _281___mcc_h0 bool = _source6.Get_().(D3_DC11).Cf19
				_ = _281___mcc_h0
				var _282___mcc_h1 bool = _source6.Get_().(D3_DC11).Cf20
				_ = _282___mcc_h1
				var _283_cf20 bool = _282___mcc_h1
				_ = _283_cf20
				var _284_cf19 bool = _281___mcc_h0
				_ = _284_cf19
				return _283_cf20
			} else {
				var _285___mcc_h2 _dafny.Map = _source6.Get_().(D3_DC10).Cf18
				_ = _285___mcc_h2
				var _286_cf18 _dafny.Map = _285___mcc_h2
				_ = _286_cf18
				return _pat_let_tv6
			}
		}(_275_v1)
		_ = _rhs14
		var _rhs15 bool = ((func() _dafny.Int {
			if p0 {
				return (_this).F30()
			}
			return (_279_v5).Dtor_cf3()
		})()).Cmp(_dafny.IntOfUint32((_280_v6).Cardinality())) != 0
		_ = _rhs15
		var _lhs10 *GlobalState = globalState
		_ = _lhs10
		var _lhs11 *GlobalState = globalState
		_ = _lhs11
		var _lhs12 _dafny.Array = _274_v0
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))
		_ = _lhs13
		_lhs10.F0 = _rhs13
		_lhs11.F0 = _rhs14
		(_lhs12).ArraySet1(_rhs15, (_lhs13).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))
		_ = _index29
		(_274_v0).ArraySet1((func() bool {
			if p0 {
				return false
			}
			return ((_274_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))).Int()).(bool)) && (p0)
		})(), (_index29).Int())
		var _287_v7 D8
		_ = _287_v7
		_287_v7 = Companion_D8_.Create_DC29_((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jn")).Cardinality())).Cmp((_276_v2).Cardinality()) <= 0, ((_this).F30()).Cmp((_this).F30()) <= 0, _278_v4)
		_287_v7 = (func() D8 {
			if (_274_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))).Int()).(bool) {
				return Companion_D8_.Create_DC29_(p0, (_274_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))).Int()).(bool), _278_v4)
			}
			return _287_v7
		})()
		(globalState).F16 = (Companion_Default___.SafeModuloInt((_this).F30(), (_this).F30())).Times((_this).F30())
		var _288_v8 _dafny.Set
		_ = _288_v8
		_288_v8 = _dafny.SetOf((_this).F30())
		var _289_v9 _dafny.MultiSet
		_ = _289_v9
		_289_v9 = _dafny.MultiSetOf(_274_v0)
		(globalState).F11 = (((_288_v8).Cardinality()).Plus((_this).F30())).Times(((_289_v9).Difference(_289_v9)).Cardinality())
		r0 = (_274_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))).Int()).(bool)
		r1 = _274_v0
		var _290_v10 _dafny.Sequence
		_ = _290_v10
		_290_v10 = _dafny.SeqOf((_this).F30())
		var _291_v11 _dafny.Sequence
		_ = _291_v11
		_291_v11 = _dafny.SeqOf(_280_v6)
		var _292_v12 _dafny.Sequence
		_ = _292_v12
		_292_v12 = _dafny.SeqOf(_dafny.IntOfUint32((_290_v10).Cardinality()), (_this).F30(), _dafny.IntOfUint32((_291_v11).Cardinality()))
		var _293_v13 _dafny.Sequence
		_ = _293_v13
		_293_v13 = _dafny.UnicodeSeqOfUtf8Bytes("aipcn")
		var _294_v14 _dafny.Map
		_ = _294_v14
		_294_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_292_v12, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_293_v13).Cardinality()), _dafny.IntOfUint32((_292_v12).Cardinality()))).Uint32(), (_this).F30()), (func() _dafny.CodePoint {
			if (_274_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_274_v0), 0))).Int()).(bool) {
				return _278_v4
			}
			return _dafny.CodePoint('y')
		})())
		r2 = _294_v14
		var _295_v15 D6
		_ = _295_v15
		_295_v15 = Companion_D6_.Create_DC22_((_dafny.Zero).Minus((_this).F30()))
		r3 = (_this).Fm35((_this).F30(), (_dafny.Zero).Minus((_this).F30()), p0, (_dafny.Zero).Minus((_295_v15).Dtor_cf36()), globalState)
		return r0, r1, r2, r3
	}
}
func (_this *C1) M6(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Set, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		(globalState).F0 = (func() bool {
			if true {
				return p1
			}
			return p1
		})()
		var _296_v0 _dafny.Map
		_ = _296_v0
		_296_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F30())
		var _297_v1 _dafny.Array
		_ = _297_v1
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw37
		_297_v1 = _nw37
		var _298_v2 _dafny.MultiSet
		_ = _298_v2
		_298_v2 = _dafny.MultiSetOf(_297_v1, _297_v1)
		var _299_v3 _dafny.Sequence
		_ = _299_v3
		_299_v3 = _dafny.SeqOf(p1, false, p1, p1, p1)
		_296_v0 = (_296_v0).Update(Companion_Default___.SafeDivisionInt((_298_v2).Cardinality(), (_dafny.MultiSetFromSeq(_299_v3)).Cardinality()), (p2).Minus((_this).Fm35(_dafny.IntOfInt64(220), p2, false, (_dafny.Zero).Minus(p0), globalState)))
		var _300_v4 _dafny.MultiSet
		_ = _300_v4
		_300_v4 = _dafny.MultiSetOf(p2)
		(globalState).F0 = !((func() bool {
			if false {
				return (_300_v4).IsSubsetOf(_300_v4)
			}
			return !((func() bool {
				if p1 {
					return p1
				}
				return !(p1)
			})())
		})())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_297_v1), 0))
		_ = _index30
		(_297_v1).ArraySet1(p1, (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_297_v1), 0))
		_ = _index31
		(_297_v1).ArraySet1(p1, (_index31).Int())
		var _301_v5 _dafny.Map
		_ = _301_v5
		_301_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_297_v1, false)
		var _302_v6 _dafny.Set
		_ = _302_v6
		_302_v6 = _dafny.SetOf(_299_v3, _dafny.SeqOf(p1, !(p1), !((func() bool {
			if (_301_v5).Contains(_297_v1) {
				return (_301_v5).Get(_297_v1).(bool)
			}
			return (_297_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_297_v1), 0))).Int()).(bool)
		})())), _299_v3, _299_v3)
		_302_v6 = ((_302_v6).Union(_302_v6)).Difference((_302_v6).Union(_dafny.SetOf(_299_v3, _dafny.Companion_Sequence_.Update(_299_v3, (Companion_Default___.SafeIndex((_this).Fm0(_300_v4, p3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(773))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}((func(_303_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
			return func(_304_i0 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_303_p0)
			}
		})(p0))), p1, globalState), _dafny.IntOfUint32((_299_v3).Cardinality()))).Uint32(), p1), _299_v3)))
		var _305_v7 _dafny.Map
		_ = _305_v7
		_305_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_297_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_297_v1), 0))).Int()).(bool), (_this).F30())
		_305_v7 = (Companion_D9_.Create_DC30_(_305_v7)).Dtor_cf50()
		r0 = _dafny.IntOfInt64(-425)
		var _306_v8 _dafny.Set
		_ = _306_v8
		_306_v8 = _dafny.SetOf((_this).F30(), p0, p0)
		r1 = _306_v8
		r2 = p0
		r3 = p2
		return r0, r1, r2, r3
	}
}
func (_this *C1) M7(p0 bool, globalState *GlobalState) (bool, _dafny.Int, bool, _dafny.Set) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Set = _dafny.EmptySet
		_ = r3
		(globalState).F8 = (_this).F30()
		r2 = p0
		var _307_v0 _dafny.MultiSet
		_ = _307_v0
		_307_v0 = _dafny.MultiSetOf(_dafny.IntOfInt64(214), (_this).F30())
		var _308_v1 _dafny.Map
		_ = _308_v1
		_308_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _307_v0)
		var _309_v2 D4
		_ = _309_v2
		_309_v2 = Companion_D4_.Create_DC12_(_308_v1)
		var _310_v3 _dafny.Map
		_ = _310_v3
		_310_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _309_v2)
		var _311_v4 _dafny.Sequence
		_ = _311_v4
		_311_v4 = _dafny.SeqOf(_dafny.IntOfInt64(112), (_310_v3).Cardinality(), (_this).F30())
		var _312_v5 _dafny.Sequence
		_ = _312_v5
		_312_v5 = _dafny.SeqOf(_311_v4, _dafny.Companion_Sequence_.Concatenate(_311_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_313_i0 _dafny.Int) _dafny.Int {
			return (_this).F30()
		}))), _311_v4)
		_311_v4 = (_312_v5).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_312_v5).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _314_v6 _dafny.CodePoint
		_ = _314_v6
		_314_v6 = _dafny.CodePoint('i')
		var _315_v7 _dafny.Map
		_ = _315_v7
		_315_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _314_v6)
		var _316_i1 _dafny.Int
		_ = _316_i1
		_316_i1 = _dafny.Zero
		{
			for ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
				if (_315_v7).Contains(p0) {
					return (_315_v7).Get(p0).(_dafny.CodePoint)
				}
				return _314_v6
			})(), (_this).F30())).Cardinality()).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality()) > 0 {
				{
					if (_316_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_316_i1 = (_316_i1).Plus(_dafny.One)
					if p0 {
						var _317_v8 _dafny.Array
						_ = _317_v8
						var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
						_ = _nw38
						_317_v8 = _nw38
						var _318_v9 _dafny.Map
						_ = _318_v9
						_318_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_314_v6, _317_v8)
						_318_v9 = (_318_v9).Update(_314_v6, _317_v8)
						_314_v6 = _314_v6
						r0 = ((_this).F30()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_311_v4, _311_v4)).Cardinality()))) >= 0
						var _319_v10 _dafny.Sequence
						_ = _319_v10
						_319_v10 = _dafny.UnicodeSeqOfUtf8Bytes("frwriukby")
						_319_v10 = _dafny.Companion_Sequence_.Concatenate(_319_v10, Companion_Default___.Fm36(p0, _dafny.SetOf((_this).F30(), (_this).F30()), (_this).F30(), (_this).F30(), globalState))
						var _320_v11 _dafny.Array
						_ = _320_v11
						var _nw39 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
						_ = _nw39
						_320_v11 = _nw39
						var _rhs16 _dafny.Sequence = _319_v10
						_ = _rhs16
						var _rhs17 _dafny.CodePoint = _314_v6
						_ = _rhs17
						var _rhs18 _dafny.Array = _320_v11
						_ = _rhs18
						var _rhs19 _dafny.Sequence = _319_v10
						_ = _rhs19
						_319_v10 = _rhs16
						_314_v6 = _rhs17
						_320_v11 = _rhs18
						_319_v10 = _rhs19
					} else {
						(globalState).F6 = (_this).F30()
						var _321_v12 D2
						_ = _321_v12
						_321_v12 = Companion_D2_.Create_DC7_(_dafny.Companion_Sequence_.Update(_311_v4, (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_311_v4).Cardinality()))).Uint32(), (_this).F30()), p0, p0, _dafny.IntOfInt64(27))
						(globalState).F0 = (_321_v12).Dtor_cf11()
						_314_v6 = _314_v6
						(globalState).F8 = ((_this).F30()).Times((_this).F30())
						var _322_v13 _dafny.Sequence
						_ = _322_v13
						_322_v13 = _dafny.SeqOf(p0)
						var _323_v14 _dafny.Map
						_ = _323_v14
						_323_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), p0)
						var _324_v15 _dafny.Sequence
						_ = _324_v15
						_324_v15 = _dafny.UnicodeSeqOfUtf8Bytes("xw")
						(globalState).F10 = Companion_Default___.Fm37(_322_v13, _323_v14, _dafny.Companion_Sequence_.Concatenate(_324_v15, _324_v15), globalState)
					}
					(globalState).F16 = (_311_v4).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_311_v4).Cardinality()))).Uint32()).(_dafny.Int)
					var _325_v16 _dafny.Sequence
					_ = _325_v16
					_325_v16 = _dafny.UnicodeSeqOfUtf8Bytes("kdib")
					_325_v16 = _325_v16
					r2 = p0
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _326_v17 D3
		_ = _326_v17
		_326_v17 = Companion_D3_.Create_DC11_(p0, true)
		var _327_i2 _dafny.Int
		_ = _327_i2
		_327_i2 = _dafny.Zero
		{
			for !((func() bool {
				if (p0) && (p0) {
					return !(!(p0)) || (p0)
				}
				return (_326_v17).Dtor_cf20()
			})()) {
				{
					if (_327_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_327_i2 = (_327_i2).Plus(_dafny.One)
					(globalState).F6 = Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pcye")).Cardinality())).Minus((_dafny.Zero).Minus((_this).F30())), ((_this).F30()).Minus((_this).F30()))
					var _328_v18 _dafny.Array
					_ = _328_v18
					var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
					_ = _nw40
					_328_v18 = _nw40
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_328_v18), 0))
					_ = _index32
					(_328_v18).ArraySet1CodePoint(Companion_Default___.Fm38(p0, (_this).F30(), globalState), (_index32).Int())
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_328_v18), 0))
					_ = _index33
					(_328_v18).ArraySet1CodePoint(_314_v6, (_index33).Int())
					var _329_v19 _dafny.Sequence
					_ = _329_v19
					_329_v19 = _dafny.UnicodeSeqOfUtf8Bytes("whqyh")
					_329_v19 = _329_v19
					var _330_v20 _dafny.MultiSet
					_ = _330_v20
					_330_v20 = _dafny.MultiSetOf(p0)
					var _331_v21 _dafny.Set
					_ = _331_v21
					_331_v21 = _dafny.SetOf((_330_v20).Cardinality())
					var _332_v22 _dafny.Map
					_ = _332_v22
					_332_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_329_v19).Cardinality()))
					var _333_v23 _dafny.Array
					_ = _333_v23
					var _nwElement0_7 _dafny.Sequence = _329_v19
					_ = _nwElement0_7
					var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(24))
					_ = _nw41
					(_nw41).ArraySet1(_nwElement0_7, 0)
					(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mx"), 1)
					(_nw41).ArraySet1(_329_v19, 2)
					(_nw41).ArraySet1((func() _dafny.Sequence {
						if p0 {
							return _329_v19
						}
						return _329_v19
					})(), 3)
					(_nw41).ArraySet1(_329_v19, 4)
					(_nw41).ArraySet1(_329_v19, 5)
					(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_329_v19, _dafny.UnicodeSeqOfUtf8Bytes("foa")), 6)
					(_nw41).ArraySet1(Companion_Default___.Fm36(p0, _331_v21, _dafny.IntOfUint32((_329_v19).Cardinality()), (_this).F30(), globalState), 7)
					(_nw41).ArraySet1(_329_v19, 8)
					(_nw41).ArraySet1(_329_v19, 9)
					(_nw41).ArraySet1(_329_v19, 10)
					(_nw41).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("k"), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()))).Uint32(), (_328_v18).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_328_v18), 0))).Int())), 11)
					(_nw41).ArraySet1(_329_v19, 12)
					(_nw41).ArraySet1(_329_v19, 13)
					(_nw41).ArraySet1(Companion_Default___.Fm36(p0, _331_v21, (func() _dafny.Int {
						if (_332_v22).Contains(p0) {
							return (_332_v22).Get(p0).(_dafny.Int)
						}
						return (_this).F30()
					})(), (_this).F30(), globalState), 14)
					(_nw41).ArraySet1(_329_v19, 15)
					(_nw41).ArraySet1(_329_v19, 16)
					(_nw41).ArraySet1((func() _dafny.Sequence {
						if p0 {
							return _dafny.UnicodeSeqOfUtf8Bytes("ohl")
						}
						return _329_v19
					})(), 17)
					(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_329_v19, _329_v19), 18)
					(_nw41).ArraySet1(_329_v19, 19)
					(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_329_v19, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("wdespalc"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F30()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wdespalc")).Cardinality()))).Uint32(), (_328_v18).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_328_v18), 0))).Int()))), 20)
					(_nw41).ArraySet1(_329_v19, 21)
					(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("n"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg38 _dafny.Int) interface{} {
							return coer38(arg38)
						}
					}((func(_334_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_335_i3 _dafny.Int) _dafny.CodePoint {
							return _334_v6
						}
					})(_314_v6)))), 22)
					(_nw41).ArraySet1(_329_v19, 23)
					_333_v23 = _nw41
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_333_v23), 0))
					_ = _index34
					(_333_v23).ArraySet1(_329_v19, (_index34).Int())
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_333_v23), 0))
					_ = _index35
					(_333_v23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_329_v19, _329_v19), (_index35).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		(globalState).F16 = (_this).F30()
		var _336_v24 _dafny.Sequence
		_ = _336_v24
		_336_v24 = _dafny.SeqOf(p0, false)
		var _337_v25 _dafny.Map
		_ = _337_v25
		_337_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), p0)
		var _338_v26 _dafny.Sequence
		_ = _338_v26
		_338_v26 = _dafny.UnicodeSeqOfUtf8Bytes("gewwn")
		r0 = !(Companion_Default___.Fm37(_336_v24, _337_v25, _dafny.Companion_Sequence_.Update(_338_v26, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_336_v24).Cardinality()), _dafny.IntOfUint32((_338_v26).Cardinality()))).Uint32(), _314_v6), globalState)) || (!(p0))
		r1 = (_this).F30()
		r2 = !(p0)
		r3 = (_dafny.SetOf(_dafny.IntOfUint32((_338_v26).Cardinality()), _dafny.IntOfInt64(169))).Intersection(_dafny.SetOf((_this).F30()))
		return r0, r1, r2, r3
	}
}
func (_this *C1) F30() _dafny.Int {
	{
		return _this._f30
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f22 _dafny.Sequence
	_f29 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f22 = _dafny.EmptySeq
	_this._f29 = _dafny.Zero
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

func (_this *C2) F22() _dafny.Sequence {
	return _this._f22
}
func (_this *C2) Ctor__(f29 _dafny.Int, f22 _dafny.Sequence) {
	{
		(_this)._f29 = f29
		(_this)._f22 = f22
	}
}
func (_this *C2) Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F29()
	}
}
func (_this *C2) Fm21(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("qomukech"), _dafny.UnicodeSeqOfUtf8Bytes("eevpwlb"))
	}
}
func (_this *C2) M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _339_v0 D3
		_ = _339_v0
		_339_v0 = Companion_D3_.Create_DC11_(p0, (func() bool {
			if !(p0) {
				return p0
			}
			return p0
		})())
		var _source7 D3 = _339_v0
		_ = _source7
		if _source7.Is_DC11() {
			var _340___mcc_h0 bool = _source7.Get_().(D3_DC11).Cf19
			_ = _340___mcc_h0
			var _341___mcc_h1 bool = _source7.Get_().(D3_DC11).Cf20
			_ = _341___mcc_h1
			var _342_cf20 bool = _341___mcc_h1
			_ = _342_cf20
			var _343_cf19 bool = _340___mcc_h0
			_ = _343_cf19
			var _344_v1 *C0
			_ = _344_v1
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__()
			_344_v1 = _nw42
			var _345_v2 _dafny.Map
			_ = _345_v2
			_345_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
			var _346_v3 _dafny.Map
			_ = _346_v3
			_346_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if (_345_v2).Contains(_342_cf20) {
					return (_345_v2).Get(_342_cf20).(bool)
				}
				return _343_cf19
			})()), p0)
			var _347_v4 _dafny.Map
			_ = _347_v4
			_347_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29()))
			var _348_v5 _dafny.CodePoint
			_ = _348_v5
			_348_v5 = _dafny.CodePoint('h')
			var _349_v6 D0
			_ = _349_v6
			_349_v6 = Companion_D0_.Create_DC1_((_345_v2).Cardinality(), _343_cf19, _347_v4, (_this).F29(), _348_v5)
			var _350_v7 D4
			_ = _350_v7
			_350_v7 = Companion_D4_.Create_DC13_(_349_v6)
			var _source8 D4 = _350_v7
			_ = _source8
			if _source8.Is_DC13() {
				var _351___mcc_h3 D0 = _source8.Get_().(D4_DC13).Cf22
				_ = _351___mcc_h3
				var _352_cf22 D0 = _351___mcc_h3
				_ = _352_cf22
				var _353_v8 _dafny.Sequence
				_ = _353_v8
				_353_v8 = _dafny.UnicodeSeqOfUtf8Bytes("bgkie")
				var _354_v9 _dafny.Map
				_ = _354_v9
				_354_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _342_cf20)
				var _355_v10 _dafny.Map
				_ = _355_v10
				_355_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm22(_342_cf20, !(_342_cf20), (_this).F29(), (Companion_Default___.Fm23((_this).F29(), _dafny.IntOfUint32((_353_v8).Cardinality()), !(false), !(_343_cf19), globalState)).Cardinality(), globalState), ((func() bool {
					if (_354_v9).Contains((_dafny.Zero).Minus((_this).F29())) {
						return (_354_v9).Get((_dafny.Zero).Minus((_this).F29())).(bool)
					}
					return false
				})()) && (true))
				var _356_v11 _dafny.Set
				_ = _356_v11
				_356_v11 = _dafny.SetOf(_342_cf20)
				var _rhs20 bool = (_dafny.SetOf(false)).IsDisjointFrom((_dafny.SetOf(_342_cf20)).Intersection(_356_v11))
				_ = _rhs20
				var _rhs21 _dafny.Map = _355_v10
				_ = _rhs21
				var _rhs22 _dafny.Int = (_this).F29()
				_ = _rhs22
				var _lhs14 *GlobalState = globalState
				_ = _lhs14
				_343_cf19 = _rhs20
				_355_v10 = _rhs21
				_lhs14.F6 = _rhs22
				var _357_v12 _dafny.Array
				_ = _357_v12
				var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw43
				_357_v12 = _nw43
				var _358_v13 _dafny.Sequence
				_ = _358_v13
				_358_v13 = _dafny.SeqOf((_this).F29())
				var _359_v14 _dafny.MultiSet
				_ = _359_v14
				_359_v14 = _dafny.MultiSetOf(_342_cf20, _343_cf19)
				var _360_v15 _dafny.Sequence
				_ = _360_v15
				_360_v15 = _dafny.SeqOf(false, p0, p0, _343_cf19)
				var _361_v16 _dafny.Sequence
				_ = _361_v16
				_361_v16 = _dafny.SeqOf(_358_v13, Companion_Default___.Fm24(_359_v14, _dafny.IntOfInt64(948), (_this).F29(), (_360_v15).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_360_v15).Cardinality()))).Uint32()).(bool), globalState), _358_v13)
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_357_v12), 0))
				_ = _index36
				(_357_v12).ArraySet1(_dafny.IntOfUint32((_361_v16).Cardinality()), (_index36).Int())
				var _362_v17 _dafny.Array
				_ = _362_v17
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_4
				var _nw44 _dafny.Array
				_ = _nw44
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw44 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) bool = func(_363_i0 _dafny.Int) bool {
						return false
					}
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw44 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw44).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw44).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_362_v17 = _nw44
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_362_v17), 0))
				_ = _index37
				(_362_v17).ArraySet1(!(_342_cf20), (_index37).Int())
				var _364_v18 D3
				_ = _364_v18
				_364_v18 = Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29()))
				var _365_v20 _dafny.Set
				_ = _365_v20
				_365_v20 = _dafny.SetOf(_364_v18, _364_v18, Companion_D3_.Create_DC10_(func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(5), _dafny.IntOfInt64(124))); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _366_v19 _dafny.Int
						_366_v19 = interface{}(_compr_29).(_dafny.Int)
						if ((_dafny.IntOfInt64(5)).Cmp(_366_v19) <= 0) && ((_366_v19).Cmp(_dafny.IntOfInt64(124)) < 0) {
							_coll29.Add(Companion_Default___.SafeModuloInt(_366_v19, (_this).F29()), (_this).F29())
						}
					}
					return _coll29.ToMap()
				}()))
				var _367_v21 _dafny.MultiSet
				_ = _367_v21
				_367_v21 = _dafny.MultiSetOf((_this).F29())
				var _368_v22 D6
				_ = _368_v22
				_368_v22 = Companion_D6_.Create_DC21_(_367_v21)
				var _369_v23 D1
				_ = _369_v23
				_369_v23 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("hk"))
				var _370_v24 D1
				_ = _370_v24
				_370_v24 = Companion_D1_.Create_DC4_(_369_v23)
				var _371_v25 _dafny.Map
				_ = _371_v25
				_371_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_370_v24, (_this).F29())
				var _372_v26 _dafny.Map
				_ = _372_v26
				_372_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("aiu"), (func() _dafny.Int {
					if (_371_v25).Contains(Companion_D1_.Create_DC4_(_369_v23)) {
						return (_371_v25).Get(Companion_D1_.Create_DC4_(_369_v23)).(_dafny.Int)
					}
					return (_dafny.Zero).Minus((_this).F29())
				})())
				var _373_v27 _dafny.Map
				_ = _373_v27
				_373_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), Companion_D5_.Create_DC17_(_372_v26))
				var _374_v28 _dafny.Set
				_ = _374_v28
				_374_v28 = _dafny.SetOf(_349_v6)
				var _375_v29 _dafny.Set
				_ = _375_v29
				_375_v29 = _dafny.SetOf(_358_v13, _358_v13, _358_v13)
				var _376_v30 D7
				_ = _376_v30
				_376_v30 = Companion_D7_.Create_DC25_(_375_v29)
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_357_v12), 0))
				_ = _index38
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_362_v17), 0))
				_ = _index39
				var _rhs23 _dafny.Int = ((_358_v13).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_358_v13).Cardinality()))).Uint32()).(_dafny.Int)).Minus((func() _dafny.Int {
					if (_359_v14).Contains(_342_cf20) {
						return (_359_v14).Multiplicity(_342_cf20)
					}
					return (_this).F29()
				})())
				_ = _rhs23
				var _rhs24 bool = (Companion_Default___.Fm25((_this).F29(), (_this).F29(), globalState)).IsSubsetOf(_365_v20)
				_ = _rhs24
				var _rhs25 bool = ((Companion_Default___.Fm26(((_368_v22).Dtor_cf35()).Cardinality(), _373_v27, p0, _343_cf19, globalState)).Difference(Companion_Default___.Fm26((_this).F29(), _373_v27, _343_cf19, (_this).Fm21(_374_v28, (_this).F29(), globalState), globalState))).IsDisjointFrom((_376_v30).Dtor_cf41())
				_ = _rhs25
				var _rhs26 _dafny.Int = _dafny.IntOfInt64(860)
				_ = _rhs26
				var _lhs15 _dafny.Array = _357_v12
				_ = _lhs15
				var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_357_v12), 0))
				_ = _lhs16
				var _lhs17 _dafny.Array = _362_v17
				_ = _lhs17
				var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_362_v17), 0))
				_ = _lhs18
				var _lhs19 *GlobalState = globalState
				_ = _lhs19
				(_lhs15).ArraySet1(_rhs23, (_lhs16).Int())
				_343_cf19 = _rhs24
				(_lhs17).ArraySet1(_rhs25, (_lhs18).Int())
				_lhs19.F6 = _rhs26
				var _377_v31 D1
				_ = _377_v31
				_377_v31 = Companion_D1_.Create_DC3_(_353_v8)
				var _pat_let_tv7 = _353_v8
				_ = _pat_let_tv7
				var _pat_let_tv8 = _353_v8
				_ = _pat_let_tv8
				var _378_v32 _dafny.Array
				_ = _378_v32
				var _nwElement0_8 D1 = (func() D1 {
					if p0 {
						return _377_v31
					}
					return _377_v31
				})()
				_ = _nwElement0_8
				var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(26))
				_ = _nw45
				(_nw45).ArraySet1(_nwElement0_8, 0)
				(_nw45).ArraySet1(_377_v31, 1)
				(_nw45).ArraySet1(Companion_D1_.Create_DC3_(_353_v8), 2)
				(_nw45).ArraySet1(_377_v31, 3)
				(_nw45).ArraySet1(_377_v31, 4)
				(_nw45).ArraySet1(_377_v31, 5)
				(_nw45).ArraySet1(_377_v31, 6)
				(_nw45).ArraySet1(_377_v31, 7)
				(_nw45).ArraySet1(Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-589))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_379_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_380_i1 _dafny.Int) _dafny.CodePoint {
						return _379_v5
					}
				})(_348_v5))), (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-589))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_381_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_382_i1 _dafny.Int) _dafny.CodePoint {
						return _381_v5
					}
				})(_348_v5)))).Cardinality()))).Uint32(), _348_v5)), 8)
				(_nw45).ArraySet1(Companion_D1_.Create_DC3_((_344_v1).Fm17(_353_v8, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uwe")).Cardinality()), _343_cf19, (Companion_Default___.Fm27(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gu")).Cardinality()), globalState)).Cardinality(), globalState)), 9)
				(_nw45).ArraySet1(_377_v31, 10)
				(_nw45).ArraySet1(_377_v31, 11)
				(_nw45).ArraySet1(_377_v31, 12)
				(_nw45).ArraySet1(_377_v31, 13)
				(_nw45).ArraySet1(_377_v31, 14)
				(_nw45).ArraySet1(_377_v31, 15)
				(_nw45).ArraySet1(func(_pat_let2_0 D1) D1 {
					return func(_383_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let3_0 _dafny.Sequence) D1 {
							return func(_384_dt__update_hcf6_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_384_dt__update_hcf6_h0)
							}(_pat_let3_0)
						}(_pat_let_tv7)
					}(_pat_let2_0)
				}(_377_v31), 16)
				(_nw45).ArraySet1(_377_v31, 17)
				(_nw45).ArraySet1(_377_v31, 18)
				(_nw45).ArraySet1(_377_v31, 19)
				(_nw45).ArraySet1(_377_v31, 20)
				(_nw45).ArraySet1(_377_v31, 21)
				(_nw45).ArraySet1(func(_pat_let4_0 D1) D1 {
					return func(_385_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let5_0 _dafny.Sequence) D1 {
							return func(_386_dt__update_hcf6_h1 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_386_dt__update_hcf6_h1)
							}(_pat_let5_0)
						}(_pat_let_tv8)
					}(_pat_let4_0)
				}(_377_v31), 22)
				(_nw45).ArraySet1(_377_v31, 23)
				(_nw45).ArraySet1(_377_v31, 24)
				(_nw45).ArraySet1(_377_v31, 25)
				_378_v32 = _nw45
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_378_v32), 0))
				_ = _index40
				(_378_v32).ArraySet1(_377_v31, (_index40).Int())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_378_v32), 0))
				_ = _index41
				(_378_v32).ArraySet1(_377_v31, (_index41).Int())
				var _387_v33 *C0
				_ = _387_v33
				var _nw46 *C0 = New_C0_()
				_ = _nw46
				_nw46.Ctor__()
				_387_v33 = _nw46
			} else if _source8.Is_DC14() {
				var _388___mcc_h4 _dafny.CodePoint = _source8.Get_().(D4_DC14).Cf23
				_ = _388___mcc_h4
				var _389___mcc_h5 _dafny.Array = _source8.Get_().(D4_DC14).Cf24
				_ = _389___mcc_h5
				var _390___mcc_h6 _dafny.Sequence = _source8.Get_().(D4_DC14).Cf25
				_ = _390___mcc_h6
				var _391_cf25 _dafny.Sequence = _390___mcc_h6
				_ = _391_cf25
				var _392_cf24 _dafny.Array = _389___mcc_h5
				_ = _392_cf24
				var _393_cf23 _dafny.CodePoint = _388___mcc_h4
				_ = _393_cf23
				_346_v3 = (_346_v3).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_394_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})), _391_cf25), p0)
				var _395_v34 _dafny.Map
				_ = _395_v34
				_395_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_396_cf23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_397_i3 _dafny.Int) _dafny.CodePoint {
						return _396_cf23
					}
				})(_393_cf23))), _dafny.UnicodeSeqOfUtf8Bytes("xeuvryom")))
				var _398_v35 _dafny.MultiSet
				_ = _398_v35
				_398_v35 = _dafny.MultiSetOf(!(p0), p0)
				_395_v34 = (_395_v34).Update((func() _dafny.Int {
					if false {
						return (_this).F29()
					}
					return (_this).F29()
				})(), (_398_v35).IsProperSubsetOf(_398_v35))
				var _399_v36 _dafny.MultiSet
				_ = _399_v36
				_399_v36 = _dafny.MultiSetOf((_this).F29(), (_this).F29(), (_this).F29(), (_this).F29())
				var _400_v37 _dafny.Set
				_ = _400_v37
				_400_v37 = _dafny.SetOf((_this).F29())
				var _401_v38 _dafny.Sequence
				_ = _401_v38
				_401_v38 = _dafny.SeqOf(_400_v37, _400_v37, _400_v37, _400_v37)
				var _402_v39 _dafny.Array
				_ = _402_v39
				var _nwElement0_9 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_403_cf23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_404_i4 _dafny.Int) _dafny.CodePoint {
						return _403_cf23
					}
				})(_393_cf23))), _391_cf25)).Cardinality())
				_ = _nwElement0_9
				var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
				_ = _nw47
				(_nw47).ArraySet1(_nwElement0_9, 0)
				(_nw47).ArraySet1(Companion_Default___.SafeModuloInt((_this).Fm0(_399_v36, (_this).F29(), _401_v38, p0, globalState), (_dafny.Zero).Minus((_this).F29())), 1)
				(_nw47).ArraySet1((_this).F29(), 2)
				(_nw47).ArraySet1(_dafny.IntOfInt64(-527), 3)
				(_nw47).ArraySet1((func() _dafny.Int {
					if _342_cf20 {
						return (_this).F29()
					}
					return (_this).F29()
				})(), 4)
				(_nw47).ArraySet1((_399_v36).Cardinality(), 5)
				(_nw47).ArraySet1((_this).F29(), 6)
				_402_v39 = _nw47
				var _405_v40 _dafny.Map
				_ = _405_v40
				_405_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_cf20, (_this).F29())
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_402_v39), 0))
				_ = _index42
				(_402_v39).ArraySet1((func() _dafny.Int {
					if (_405_v40).Contains(_343_cf19) {
						return (_405_v40).Get(_343_cf19).(_dafny.Int)
					}
					return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}((func(_406_cf25 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_407_i5 _dafny.Int) _dafny.Sequence {
							return _406_cf25
						}
					})(_391_cf25))), (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}((func(_408_cf25 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_409_i5 _dafny.Int) _dafny.Sequence {
							return _408_cf25
						}
					})(_391_cf25)))).Cardinality()))).Uint32(), _391_cf25))).Cardinality()
				})(), (_index42).Int())
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_402_v39), 0))
				_ = _index43
				var _rhs27 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F29(), (_dafny.Zero).Minus((_this).F29()))
				_ = _rhs27
				var _rhs28 bool = _342_cf20
				_ = _rhs28
				var _rhs29 bool = !(_343_cf19) || (p0)
				_ = _rhs29
				var _rhs30 _dafny.Int = (_this).F29()
				_ = _rhs30
				var _rhs31 _dafny.Array = _392_cf24
				_ = _rhs31
				var _lhs20 *GlobalState = globalState
				_ = _lhs20
				var _lhs21 *GlobalState = globalState
				_ = _lhs21
				var _lhs22 *GlobalState = globalState
				_ = _lhs22
				var _lhs23 _dafny.Array = _402_v39
				_ = _lhs23
				var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_402_v39), 0))
				_ = _lhs24
				_lhs20.F11 = _rhs27
				_lhs21.F10 = _rhs28
				_lhs22.F10 = _rhs29
				(_lhs23).ArraySet1(_rhs30, (_lhs24).Int())
				_392_cf24 = _rhs31
				var _410_v41 _dafny.Int
				_ = _410_v41
				var _411_v42 bool
				_ = _411_v42
				var _412_v43 bool
				_ = _412_v43
				var _413_v44 _dafny.Int
				_ = _413_v44
				var _out34 _dafny.Int
				_ = _out34
				var _out35 bool
				_ = _out35
				var _out36 bool
				_ = _out36
				var _out37 _dafny.Int
				_ = _out37
				_out34, _out35, _out36, _out37 = (_this).M5((_dafny.Zero).Minus((_this).F29()), (_this).F29(), globalState)
				_410_v41 = _out34
				_411_v42 = _out35
				_412_v43 = _out36
				_413_v44 = _out37
			} else if _source8.Is_DC15() {
				var _414___mcc_h7 _dafny.Int = _source8.Get_().(D4_DC15).Cf26
				_ = _414___mcc_h7
				var _415___mcc_h8 _dafny.Int = _source8.Get_().(D4_DC15).Cf27
				_ = _415___mcc_h8
				var _416_cf27 _dafny.Int = _415___mcc_h8
				_ = _416_cf27
				var _417_cf26 _dafny.Int = _414___mcc_h7
				_ = _417_cf26
				var _418_v45 _dafny.Map
				_ = _418_v45
				_418_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29())
				var _419_v46 _dafny.Map
				_ = _419_v46
				_419_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_416_cf27, Companion_D6_.Create_DC23_(_342_cf20, (func() _dafny.Int {
					if (_418_v45).Contains(_dafny.IntOfInt64(42)) {
						return (_418_v45).Get(_dafny.IntOfInt64(42)).(_dafny.Int)
					}
					return (_this).F29()
				})(), p0))
				var _420_v47 D6
				_ = _420_v47
				_420_v47 = Companion_D6_.Create_DC23_(p0, _416_cf27, false)
				_419_v46 = (_419_v46).Update(_dafny.IntOfInt64(144), _420_v47)
				(globalState).F16 = _417_cf26
				(globalState).F6 = Companion_Default___.SafeDivisionInt(_417_cf26, Companion_Default___.SafeModuloInt(_417_cf26, _417_cf26))
				var _421_v48 _dafny.MultiSet
				_ = _421_v48
				_421_v48 = _dafny.MultiSetOf(_417_cf26)
				var _422_v49 _dafny.Array
				_ = _422_v49
				var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw48
				_422_v49 = _nw48
				var _423_v50 _dafny.Map
				_ = _423_v50
				_423_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_cf19, (_this).F29())
				var _424_v51 D4
				_ = _424_v51
				_424_v51 = Companion_D4_.Create_DC14_(_348_v5, _422_v49, _dafny.Companion_Sequence_.Update((_344_v1).Fm17(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(897))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_425_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_426_i6 _dafny.Int) _dafny.CodePoint {
						return _425_v5
					}
				})(_348_v5))), _dafny.IntOfInt64(496), _342_cf20, _417_cf26, globalState), (Companion_Default___.SafeIndex((_423_v50).Cardinality(), _dafny.IntOfUint32(((_344_v1).Fm17(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(897))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_427_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_428_i6 _dafny.Int) _dafny.CodePoint {
						return _427_v5
					}
				})(_348_v5))), _dafny.IntOfInt64(496), _342_cf20, _417_cf26, globalState)).Cardinality()))).Uint32(), _348_v5))
				var _429_v52 D4
				_ = _429_v52
				_429_v52 = Companion_D4_.Create_DC16_(_424_v51)
				var _430_v53 _dafny.Map
				_ = _430_v53
				_430_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _429_v52)
				var _431_v55 _dafny.Set
				_ = _431_v55
				_431_v55 = _dafny.SetOf(_dafny.IntOfInt64(124))
				(globalState).F11 = (_this).Fm0(_421_v48, (_430_v53).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(576))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_432_cf26 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_433_i7 _dafny.Int) _dafny.Set {
						return func() _dafny.Set {
							var _coll30 = _dafny.NewBuilder()
							_ = _coll30
							for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(742), _dafny.IntOfInt64(-148))); ; {
								_compr_30, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _434_v54 _dafny.Int
								_434_v54 = interface{}(_compr_30).(_dafny.Int)
								if ((_dafny.IntOfInt64(742)).Cmp(_434_v54) <= 0) && ((_434_v54).Cmp(_dafny.IntOfInt64(-148)) < 0) {
									_coll30.Add(Companion_Default___.SafeModuloInt(_434_v54, _432_cf26))
								}
							}
							return _coll30.ToSet()
						}()
					}
				})(_417_cf26))), _dafny.SeqOf(_431_v55, _431_v55)), p0, globalState)
			} else if _source8.Is_DC12() {
				var _435___mcc_h9 _dafny.Map = _source8.Get_().(D4_DC12).Cf21
				_ = _435___mcc_h9
				var _436_cf21 _dafny.Map = _435___mcc_h9
				_ = _436_cf21
				(globalState).F10 = _343_cf19
				var _437_v56 _dafny.Map
				_ = _437_v56
				_437_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(960))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}(func(_438_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})), (_this).F29())
				var _439_v57 _dafny.Sequence
				_ = _439_v57
				_439_v57 = _dafny.UnicodeSeqOfUtf8Bytes("delahctx")
				var _440_v58 D5
				_ = _440_v58
				_440_v58 = Companion_D5_.Create_DC17_((_437_v56).Update(_439_v57, (_this).F29()))
				var _441_v59 D5
				_ = _441_v59
				_441_v59 = Companion_D5_.Create_DC20_(_440_v58)
				_441_v59 = _441_v59
				var _442_v60 _dafny.Array
				_ = _442_v60
				var _nw49 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw49
				_442_v60 = _nw49
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_442_v60), 0))
				_ = _index44
				(_442_v60).ArraySet1(false, (_index44).Int())
				var _443_v61 _dafny.MultiSet
				_ = _443_v61
				_443_v61 = _dafny.MultiSetOf((_this).F29(), (_this).F29())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_442_v60), 0))
				_ = _index45
				(_442_v60).ArraySet1((_443_v61).IsDisjointFrom(_443_v61), (_index45).Int())
				var _444_v62 _dafny.Array
				_ = _444_v62
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(12))
				_ = _nw50
				_444_v62 = _nw50
				var _445_v63 _dafny.Map
				_ = _445_v63
				_445_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29())
				var _446_v64 D3
				_ = _446_v64
				_446_v64 = Companion_D3_.Create_DC10_(_445_v63)
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_444_v62), 0))
				_ = _index46
				(_444_v62).ArraySet1(_446_v64, (_index46).Int())
				var _447_v65 _dafny.Set
				_ = _447_v65
				_447_v65 = _dafny.SetOf(_349_v6)
				var _448_v66 _dafny.Map
				_ = _448_v66
				_448_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_439_v57, (_442_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_442_v60), 0))).Int()).(bool))
				var _449_v67 _dafny.Set
				_ = _449_v67
				_449_v67 = _dafny.SetOf((_448_v66).Cardinality())
				var _450_v68 _dafny.Map
				_ = _450_v68
				_450_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F29())
				var _451_v69 _dafny.Sequence
				_ = _451_v69
				_451_v69 = _dafny.SeqOf(p0, true)
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_444_v62), 0))
				_ = _index47
				var _rhs32 D3 = Companion_D3_.Create_DC10_(Companion_Default___.Fm28(_342_cf20, _343_cf19, !((_this).Fm21(_447_v65, (_449_v67).Cardinality(), globalState)), _450_v68, globalState))
				_ = _rhs32
				var _rhs33 bool = p0
				_ = _rhs33
				var _rhs34 bool = !((_451_v69).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_451_v69).Cardinality()))).Uint32()).(bool))
				_ = _rhs34
				var _lhs25 _dafny.Array = _444_v62
				_ = _lhs25
				var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_444_v62), 0))
				_ = _lhs26
				var _lhs27 *GlobalState = globalState
				_ = _lhs27
				(_lhs25).ArraySet1(_rhs32, (_lhs26).Int())
				_lhs27.F0 = _rhs33
				_343_cf19 = _rhs34
			} else {
				var _452___mcc_h10 D4 = _source8.Get_().(D4_DC16).Cf28
				_ = _452___mcc_h10
				var _453_cf28 D4 = _452___mcc_h10
				_ = _453_cf28
				var _454_v70 _dafny.Array
				_ = _454_v70
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_5
				var _nw51 _dafny.Array
				_ = _nw51
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw51 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Int = func(_455_i9 _dafny.Int) _dafny.Int {
						return (_455_i9).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F29())).Cardinality()))
					}
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw51 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw51).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw51).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_454_v70 = _nw51
				_454_v70 = (func() _dafny.Array {
					if !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("roqfhujdu"), _348_v5) {
						return _454_v70
					}
					return _454_v70
				})()
				var _456_v71 _dafny.Array
				_ = _456_v71
				var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw52
				_456_v71 = _nw52
				var _457_v72 _dafny.Map
				_ = _457_v72
				_457_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vsdcwur"), _456_v71)
				_457_v72 = (_457_v72).Update(_dafny.UnicodeSeqOfUtf8Bytes("tlmwfrvcm"), _456_v71)
				(globalState).F8 = (_this).F29()
				var _458_v73 _dafny.Sequence
				_ = _458_v73
				_458_v73 = _dafny.UnicodeSeqOfUtf8Bytes("frayqgv")
				var _rhs35 bool = ((_this).F29()).Cmp((_this).F29()) < 0
				_ = _rhs35
				var _rhs36 _dafny.Array = _454_v70
				_ = _rhs36
				var _rhs37 bool = _342_cf20
				_ = _rhs37
				var _rhs38 _dafny.Sequence = (_344_v1).Fm17(_458_v73, (_this).F29(), _dafny.Companion_Sequence_.IsPrefixOf(_458_v73, _458_v73), (_this).F29(), globalState)
				_ = _rhs38
				r0 = _rhs35
				_454_v70 = _rhs36
				_343_cf19 = _rhs37
				_458_v73 = _rhs38
			}
			var _459_v74 _dafny.Array
			_ = _459_v74
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw53
			_459_v74 = _nw53
			var _460_v75 _dafny.Sequence
			_ = _460_v75
			_460_v75 = _dafny.SeqOf(p0)
			var _461_v76 _dafny.Map
			_ = _461_v76
			_461_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_459_v74, !(!((_460_v75).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_460_v75).Cardinality()))).Uint32()).(bool)) || (p0)))
			_461_v76 = (_461_v76).Update(_459_v74, p0)
			var _462_v77 _dafny.Array
			_ = _462_v77
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw54
			_462_v77 = _nw54
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_462_v77), 0))
			_ = _index48
			(_462_v77).ArraySet1((func() bool {
				if p0 {
					return _342_cf20
				}
				return false
			})(), (_index48).Int())
			var _463_v78 _dafny.Sequence
			_ = _463_v78
			_463_v78 = _dafny.UnicodeSeqOfUtf8Bytes("xljlu")
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_462_v77), 0))
			_ = _index49
			var _rhs39 _dafny.Int = _dafny.IntOfUint32((_463_v78).Cardinality())
			_ = _rhs39
			var _rhs40 bool = _343_cf19
			_ = _rhs40
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			var _lhs29 _dafny.Array = _462_v77
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_462_v77), 0))
			_ = _lhs30
			_lhs28.F11 = _rhs39
			(_lhs29).ArraySet1(_rhs40, (_lhs30).Int())
		} else {
			var _464___mcc_h2 _dafny.Map = _source7.Get_().(D3_DC10).Cf18
			_ = _464___mcc_h2
			var _465_cf18 _dafny.Map = _464___mcc_h2
			_ = _465_cf18
			var _466_v79 D2
			_ = _466_v79
			_466_v79 = Companion_D2_.Create_DC6_((_this).F29())
			var _467_v80 D2
			_ = _467_v80
			_467_v80 = Companion_D2_.Create_DC9_(_466_v79)
			var _468_v81 _dafny.CodePoint
			_ = _468_v81
			_468_v81 = _dafny.CodePoint('j')
			var _469_v82 _dafny.Array
			_ = _469_v82
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(18))
			_ = _nw55
			_469_v82 = _nw55
			var _470_v83 _dafny.MultiSet
			_ = _470_v83
			_470_v83 = _dafny.MultiSetOf(_469_v82)
			var _471_v84 _dafny.Map
			_ = _471_v84
			_471_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v81, _470_v83)
			var _472_v85 _dafny.Map
			_ = _472_v85
			_472_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_467_v80, (func() _dafny.MultiSet {
				if (_471_v84).Contains(_468_v81) {
					return (_471_v84).Get(_468_v81).(_dafny.MultiSet)
				}
				return (_470_v83).Update(_469_v82, Companion_Default___.Abs((_this).F29()))
			})())
			var _473_v86 _dafny.Sequence
			_ = _473_v86
			_473_v86 = _dafny.SeqOf(_470_v83)
			_472_v85 = (_472_v85).Update(_467_v80, ((_473_v86).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_473_v86).Cardinality()))).Uint32()).(_dafny.MultiSet)).Intersection(_470_v83))
			var _474_v87 _dafny.Map
			_ = _474_v87
			_474_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p0)
			var _475_v88 _dafny.Sequence
			_ = _475_v88
			_475_v88 = _dafny.SeqOf((func() bool {
				if (_474_v87).Contains((_this).F29()) {
					return (_474_v87).Get((_this).F29()).(bool)
				}
				return p0
			})())
			var _476_v89 _dafny.Sequence
			_ = _476_v89
			_476_v89 = _dafny.UnicodeSeqOfUtf8Bytes("lklrnmhq")
			var _477_v90 _dafny.Array
			_ = _477_v90
			var _nwElement0_10 bool = p0
			_ = _nwElement0_10
			var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(14))
			_ = _nw56
			(_nw56).ArraySet1(_nwElement0_10, 0)
			(_nw56).ArraySet1((_475_v88).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_dafny.IntOfUint32((_476_v89).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_475_v88).Cardinality()))).Uint32()).(bool), 1)
			(_nw56).ArraySet1(!(p0), 2)
			(_nw56).ArraySet1(true, 3)
			(_nw56).ArraySet1(p0, 4)
			(_nw56).ArraySet1((_475_v88).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_475_v88).Cardinality()))).Uint32()).(bool), 5)
			(_nw56).ArraySet1(p0, 6)
			(_nw56).ArraySet1(p0, 7)
			(_nw56).ArraySet1(true, 8)
			(_nw56).ArraySet1(p0, 9)
			(_nw56).ArraySet1(p0, 10)
			(_nw56).ArraySet1(p0, 11)
			(_nw56).ArraySet1(p0, 12)
			(_nw56).ArraySet1(p0, 13)
			_477_v90 = _nw56
			var _478_v91 _dafny.Array
			_ = _478_v91
			var _nwElement0_11 _dafny.Array = _477_v90
			_ = _nwElement0_11
			var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(6))
			_ = _nw57
			(_nw57).ArraySet1(_nwElement0_11, 0)
			(_nw57).ArraySet1(_477_v90, 1)
			(_nw57).ArraySet1(_477_v90, 2)
			(_nw57).ArraySet1((func() _dafny.Array {
				if p0 {
					return _477_v90
				}
				return _477_v90
			})(), 3)
			(_nw57).ArraySet1(_477_v90, 4)
			(_nw57).ArraySet1(_477_v90, 5)
			_478_v91 = _nw57
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_478_v91), 0))
			_ = _index50
			(_478_v91).ArraySet1(_477_v90, (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_478_v91), 0))
			_ = _index51
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_6
			var _nw58 _dafny.Array
			_ = _nw58
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw58 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) bool = (func(_479_p0 bool) func(_dafny.Int) bool {
					return func(_480_i10 _dafny.Int) bool {
						return _479_p0
					}
				})(p0)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw58 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw58).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw58).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			(_478_v91).ArraySet1(_nw58, (_index51).Int())
			var _481_v92 _dafny.MultiSet
			_ = _481_v92
			_481_v92 = _dafny.MultiSetOf(p0)
			var _482_v93 _dafny.Int
			_ = _482_v93
			var _483_v94 bool
			_ = _483_v94
			var _484_v95 bool
			_ = _484_v95
			var _485_v96 _dafny.Int
			_ = _485_v96
			var _out38 _dafny.Int
			_ = _out38
			var _out39 bool
			_ = _out39
			var _out40 bool
			_ = _out40
			var _out41 _dafny.Int
			_ = _out41
			_out38, _out39, _out40, _out41 = (_this).M5((_dafny.Zero).Minus((_this).F29()), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_465_cf18).Contains((_481_v92).Cardinality()) {
					return (_465_cf18).Get((_481_v92).Cardinality()).(_dafny.Int)
				}
				return (_dafny.MultiSetOf(p0)).Cardinality()
			})(), (_this).F29())), globalState)
			_482_v93 = _out38
			_483_v94 = _out39
			_484_v95 = _out40
			_485_v96 = _out41
			var _486_v97 _dafny.Array
			_ = _486_v97
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_7
			var _nw59 _dafny.Array
			_ = _nw59
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw59 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = (func(_487_v96 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_488_i11 _dafny.Int) _dafny.Int {
						return (_488_i11).Times(_487_v96)
					}
				})(_485_v96)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw59 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw59).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw59).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_486_v97 = _nw59
			var _489_v98 _dafny.Sequence
			_ = _489_v98
			_489_v98 = _dafny.SeqOf(_486_v97, _486_v97, _486_v97)
			var _490_v99 *C0
			_ = _490_v99
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__()
			_490_v99 = _nw60
			_486_v97 = (_489_v98).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_482_v93, _490_v99)).Cardinality(), _dafny.IntOfUint32((_489_v98).Cardinality()))).Uint32()).(_dafny.Array)
		}
		var _491_v100 _dafny.Array
		_ = _491_v100
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
		_ = _nw61
		_491_v100 = _nw61
		var _492_v101 _dafny.Array
		_ = _492_v101
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw62
		_492_v101 = _nw62
		var _nwElement0_12 _dafny.Array = _492_v101
		_ = _nwElement0_12
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(7))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_12, 0)
		(_nw63).ArraySet1(_492_v101, 1)
		(_nw63).ArraySet1(_492_v101, 2)
		(_nw63).ArraySet1(_492_v101, 3)
		(_nw63).ArraySet1(((_this).F22()).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32(((_this).F22()).Cardinality()))).Uint32()).(_dafny.Array), 4)
		(_nw63).ArraySet1(_492_v101, 5)
		(_nw63).ArraySet1(_492_v101, 6)
		_491_v100 = _nw63
		r3 = (Companion_Default___.SafeDivisionInt((_this).F29(), (Companion_Default___.Fm29(p0, _dafny.IntOfInt64(277), p0, (_this).F29(), globalState)).Cardinality())).Minus(_dafny.IntOfInt64(638))
		var _hi1 _dafny.Int = _dafny.IntOfInt64(388)
		_ = _hi1
		for _493_i12 := (func() _dafny.Int {
			if p0 {
				return (_this).F29()
			}
			return (_this).F29()
		})(); _493_i12.Cmp(_hi1) < 0; _493_i12 = _493_i12.Plus(_dafny.One) {
			var _494_v102 _dafny.CodePoint
			_ = _494_v102
			_494_v102 = _dafny.CodePoint('u')
			var _495_v103 _dafny.Sequence
			_ = _495_v103
			_495_v103 = _dafny.SeqOf(_494_v102)
			_495_v103 = _dafny.Companion_Sequence_.Concatenate(_495_v103, _495_v103)
			(globalState).F16 = _493_i12
			(globalState).F6 = _493_i12
			var _496_v104 _dafny.MultiSet
			_ = _496_v104
			_496_v104 = _dafny.MultiSetOf((_this).F29(), _dafny.IntOfInt64(680))
			var _497_v106 _dafny.Sequence
			_ = _497_v106
			_497_v106 = _dafny.SeqOf(p0, p0)
			var _498_v107 _dafny.Sequence
			_ = _498_v107
			_498_v107 = _dafny.SeqOf(_dafny.IntOfUint32((_497_v106).Cardinality()), (_this).F29())
			var _499_v108 D4
			_ = _499_v108
			_499_v108 = Companion_D4_.Create_DC14_(_494_v102, _492_v101, _dafny.UnicodeSeqOfUtf8Bytes("gxovnuk"))
			var _500_v109 _dafny.MultiSet
			_ = _500_v109
			_500_v109 = _dafny.MultiSetOf(_499_v108)
			var _501_v110 _dafny.Map
			_ = _501_v110
			_501_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_500_v109).Cardinality())
			var _502_v111 _dafny.Set
			_ = _502_v111
			_502_v111 = _dafny.SetOf(_dafny.IntOfInt64(12), _493_i12, (_498_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_498_v107).Cardinality()), _dafny.IntOfUint32((_498_v107).Cardinality()))).Uint32()).(_dafny.Int))
			var _503_v112 _dafny.Map
			_ = _503_v112
			_503_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _502_v111)
			var _504_v113 _dafny.Sequence
			_ = _504_v113
			_504_v113 = _dafny.SeqOf((func() _dafny.Set {
				if (_503_v112).Contains(p0) {
					return (_503_v112).Get(p0).(_dafny.Set)
				}
				return _502_v111
			})(), _502_v111)
			var _505_v114 _dafny.Set
			_ = _505_v114
			_505_v114 = _dafny.SetOf((_this).Fm0(_dafny.MultiSetFromSeq(_498_v107), (_501_v110).Cardinality(), _504_v113, false, globalState))
			var _506_v115 _dafny.Sequence
			_ = _506_v115
			_506_v115 = _dafny.SeqOf(func() _dafny.Set {
				var _coll31 = _dafny.NewBuilder()
				_ = _coll31
				for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(984), _dafny.IntOfInt64(889))); ; {
					_compr_31, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _507_v105 _dafny.Int
					_507_v105 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(984)).Cmp(_507_v105) <= 0) && ((_507_v105).Cmp(_dafny.IntOfInt64(889)) < 0) {
						_coll31.Add((_507_v105).Plus(_493_i12))
					}
				}
				return _coll31.ToSet()
			}(), _505_v114)
			var _508_v116 D6
			_ = _508_v116
			_508_v116 = Companion_D6_.Create_DC22_((_this).Fm0(_496_v104, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekh")).Cardinality()), _504_v113, p0, globalState))
			var _509_v117 _dafny.Map
			_ = _509_v117
			_509_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xxitwh")).Cardinality()), (_this).F29())
			var _510_v118 _dafny.Map
			_ = _510_v118
			_510_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _509_v117)
			var _511_v119 D0
			_ = _511_v119
			_511_v119 = Companion_D0_.Create_DC1_((_508_v116).Dtor_cf36(), p0, _510_v118, _493_i12, _dafny.CodePoint('h'))
			var _512_v120 D4
			_ = _512_v120
			_512_v120 = Companion_D4_.Create_DC13_(_511_v119)
			_512_v120 = _512_v120
		}
		var _513_v121 _dafny.Set
		_ = _513_v121
		_513_v121 = _dafny.SetOf(p0)
		if (_513_v121).IsSubsetOf(_513_v121) {
			var _514_v122 _dafny.Sequence
			_ = _514_v122
			_514_v122 = _dafny.UnicodeSeqOfUtf8Bytes("bknlragrj")
			var _515_v123 _dafny.Map
			_ = _515_v123
			_515_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_514_v122).Cardinality()))
			(globalState).F6 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(69), (_515_v123).Cardinality())
			var _516_v124 _dafny.Map
			_ = _516_v124
			_516_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(globalState), false)
			(globalState).F0 = ((func() _dafny.Int {
				if (_515_v123).Contains(p0) {
					return (_515_v123).Get(p0).(_dafny.Int)
				}
				return (_this).F29()
			})()).Cmp((_dafny.Zero).Minus(((_516_v124).Merge(_516_v124)).Cardinality())) < 0
			var _517_v125 _dafny.Array
			_ = _517_v125
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw64
			_517_v125 = _nw64
			var _518_v127 _dafny.Map
			_ = _518_v127
			_518_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.SetOf((_this).F29()))).Contains(func() _dafny.Set {
				var _coll32 = _dafny.NewBuilder()
				_ = _coll32
				for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-266), _dafny.IntOfInt64(398))); ; {
					_compr_32, _ok33 := _iter33()
					if !_ok33 {
						break
					}
					var _519_v126 _dafny.Int
					_519_v126 = interface{}(_compr_32).(_dafny.Int)
					if ((_dafny.IntOfInt64(-266)).Cmp(_519_v126) <= 0) && ((_519_v126).Cmp(_dafny.IntOfInt64(398)) < 0) {
						_coll32.Add(Companion_Default___.SafeModuloInt(_519_v126, (_dafny.Zero).Minus((_this).F29())))
					}
				}
				return _coll32.ToSet()
			}()), _517_v125)
			_517_v125 = (func() _dafny.Array {
				if (_518_v127).Contains(p0) {
					return (_518_v127).Get(p0).(_dafny.Array)
				}
				return _517_v125
			})()
			var _520_v128 D2
			_ = _520_v128
			_520_v128 = Companion_D2_.Create_DC5_(_517_v125)
			var _pat_let_tv9 = _517_v125
			_ = _pat_let_tv9
			var _521_v129 _dafny.Map
			_ = _521_v129
			_521_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_492_v101, (func(_pat_let6_0 D2) D2 {
				return func(_522_dt__update__tmp_h2 D2) D2 {
					return func(_pat_let7_0 _dafny.Array) D2 {
						return func(_523_dt__update_hcf8_h0 _dafny.Array) D2 {
							return Companion_D2_.Create_DC5_(_523_dt__update_hcf8_h0)
						}(_pat_let7_0)
					}(_pat_let_tv9)
				}(_pat_let6_0)
			}(_520_v128)).Dtor_cf8())
			(globalState).F6 = (_dafny.Zero).Minus(((_521_v129).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_492_v101, _517_v125)).Update(_492_v101, _517_v125))).Cardinality())
			if !(p0) || (p0) {
				var _524_v130 _dafny.MultiSet
				_ = _524_v130
				_524_v130 = _dafny.MultiSetOf((_this).F29())
				(globalState).F10 = ((_524_v130).IsDisjointFrom(_524_v130)) == (p0)
				var _525_v131 _dafny.Set
				_ = _525_v131
				_525_v131 = _dafny.SetOf(_492_v101, _492_v101, _492_v101)
				_525_v131 = (func() _dafny.Set {
					if p0 {
						return _525_v131
					}
					return (_dafny.SetOf(_492_v101)).Intersection(_525_v131)
				})()
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_517_v125), 0))
				_ = _index52
				(_517_v125).ArraySet1((_this).F29(), (_index52).Int())
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_492_v101), 0))
				_ = _index53
				(_492_v101).ArraySet1(p0, (_index53).Int())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_517_v125), 0))
				_ = _index54
				(_517_v125).ArraySet1((_this).F29(), (_index54).Int())
				var _526_v132 _dafny.Sequence
				_ = _526_v132
				_526_v132 = _dafny.SeqOf(p0)
				var _527_v133 _dafny.Set
				_ = _527_v133
				_527_v133 = _dafny.SetOf(_dafny.IntOfInt64(686), _dafny.IntOfInt64(316))
				var _528_v134 _dafny.Sequence
				_ = _528_v134
				_528_v134 = _dafny.SeqOf((_this).Fm0(_524_v130, _dafny.IntOfUint32((_526_v132).Cardinality()), _dafny.SeqOf(_527_v133), p0, globalState))
				var _529_v135 D6
				_ = _529_v135
				_529_v135 = Companion_D6_.Create_DC22_((_this).F29())
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_517_v125), 0))
				_ = _index55
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_492_v101), 0))
				_ = _index56
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_517_v125), 0))
				_ = _index57
				var _rhs41 _dafny.Int = _dafny.IntOfUint32((_528_v134).Cardinality())
				_ = _rhs41
				var _rhs42 _dafny.Int = (Companion_Default___.Fm31(false, globalState)).Dtor_cf27()
				_ = _rhs42
				var _rhs43 bool = p0
				_ = _rhs43
				var _rhs44 _dafny.Int = ((func() D6 {
					if p0 {
						return _529_v135
					}
					return _529_v135
				})()).Dtor_cf36()
				_ = _rhs44
				var _lhs31 _dafny.Array = _517_v125
				_ = _lhs31
				var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_517_v125), 0))
				_ = _lhs32
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				var _lhs34 _dafny.Array = _492_v101
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_492_v101), 0))
				_ = _lhs35
				var _lhs36 _dafny.Array = _517_v125
				_ = _lhs36
				var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_517_v125), 0))
				_ = _lhs37
				(_lhs31).ArraySet1(_rhs41, (_lhs32).Int())
				_lhs33.F11 = _rhs42
				(_lhs34).ArraySet1(_rhs43, (_lhs35).Int())
				(_lhs36).ArraySet1(_rhs44, (_lhs37).Int())
				var _530_v136 _dafny.Sequence
				_ = _530_v136
				_530_v136 = _dafny.SeqOf((_527_v133).Union(_527_v133))
				_530_v136 = _dafny.Companion_Sequence_.Concatenate(_530_v136, _530_v136)
				var _531_v137 _dafny.Array
				_ = _531_v137
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw65
				_531_v137 = _nw65
				var _532_v138 _dafny.Sequence
				_ = _532_v138
				_532_v138 = _dafny.SeqOf(_531_v137)
				var _533_v139 D8
				_ = _533_v139
				_533_v139 = Companion_D8_.Create_DC28_(_532_v138)
				var _534_v140 D8
				_ = _534_v140
				_534_v140 = Companion_D8_.Create_DC28_((_533_v139).Dtor_cf46())
				var _535_v141 _dafny.Array
				_ = _535_v141
				var _nwElement0_13 _dafny.Sequence = _532_v138
				_ = _nwElement0_13
				var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(21))
				_ = _nw66
				(_nw66).ArraySet1(_nwElement0_13, 0)
				(_nw66).ArraySet1(_532_v138, 1)
				(_nw66).ArraySet1((func() _dafny.Sequence {
					if p0 {
						return _dafny.SeqOf(_531_v137)
					}
					return _dafny.SeqOf(_531_v137)
				})(), 2)
				(_nw66).ArraySet1(_532_v138, 3)
				(_nw66).ArraySet1(_dafny.SeqOf(_531_v137, _531_v137, _531_v137), 4)
				(_nw66).ArraySet1(_532_v138, 5)
				(_nw66).ArraySet1(_532_v138, 6)
				(_nw66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_532_v138, _532_v138), 7)
				(_nw66).ArraySet1(_532_v138, 8)
				(_nw66).ArraySet1((_534_v140).Dtor_cf46(), 9)
				(_nw66).ArraySet1(_532_v138, 10)
				(_nw66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_532_v138, _532_v138), 11)
				(_nw66).ArraySet1(_532_v138, 12)
				(_nw66).ArraySet1(_dafny.SeqOf(_531_v137), 13)
				(_nw66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_532_v138, _dafny.SeqOf(_531_v137, _531_v137)), 14)
				(_nw66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_532_v138, _532_v138), 15)
				(_nw66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_532_v138, _532_v138), 16)
				(_nw66).ArraySet1(_532_v138, 17)
				(_nw66).ArraySet1(_532_v138, 18)
				(_nw66).ArraySet1(_532_v138, 19)
				(_nw66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_532_v138, _532_v138), 20)
				_535_v141 = _nw66
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_535_v141), 0))
				_ = _index58
				(_535_v141).ArraySet1(_532_v138, (_index58).Int())
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_535_v141), 0))
				_ = _index59
				(_535_v141).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_532_v138, _dafny.SeqOf(_531_v137, _531_v137, _531_v137, _531_v137, _531_v137)), _dafny.SeqOf(_531_v137, _531_v137, _531_v137, _531_v137, (func() _dafny.Array {
					if (_518_v127).Contains(p0) {
						return (_518_v127).Get(p0).(_dafny.Array)
					}
					return _531_v137
				})())), (_index59).Int())
			} else {
				(globalState).F8 = (_this).F29()
				var _536_v142 _dafny.Array
				_ = _536_v142
				var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw67
				_536_v142 = _nw67
				var _537_v143 _dafny.Sequence
				_ = _537_v143
				_537_v143 = _dafny.SeqOf((_this).F29(), (_this).F29())
				var _538_v144 _dafny.Set
				_ = _538_v144
				_538_v144 = _dafny.SetOf(_537_v143)
				var _539_v145 D7
				_ = _539_v145
				_539_v145 = Companion_D7_.Create_DC25_(_538_v144)
				var _pat_let_tv10 = _537_v143
				_ = _pat_let_tv10
				var _pat_let_tv11 = _537_v143
				_ = _pat_let_tv11
				var _540_v146 _dafny.Array
				_ = _540_v146
				var _nwElement0_14 D7 = _539_v145
				_ = _nwElement0_14
				var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(24))
				_ = _nw68
				(_nw68).ArraySet1(_nwElement0_14, 0)
				(_nw68).ArraySet1(_539_v145, 1)
				(_nw68).ArraySet1(Companion_Default___.Fm32(globalState), 2)
				(_nw68).ArraySet1(Companion_D7_.Create_DC25_(_dafny.SetOf(_537_v143)), 3)
				(_nw68).ArraySet1(_539_v145, 4)
				(_nw68).ArraySet1(_539_v145, 5)
				(_nw68).ArraySet1(Companion_D7_.Create_DC25_(_538_v144), 6)
				(_nw68).ArraySet1(_539_v145, 7)
				(_nw68).ArraySet1(Companion_D7_.Create_DC25_(_538_v144), 8)
				(_nw68).ArraySet1(_539_v145, 9)
				(_nw68).ArraySet1(_539_v145, 10)
				(_nw68).ArraySet1(_539_v145, 11)
				(_nw68).ArraySet1(Companion_D7_.Create_DC25_(_538_v144), 12)
				(_nw68).ArraySet1(_539_v145, 13)
				(_nw68).ArraySet1(_539_v145, 14)
				(_nw68).ArraySet1(_539_v145, 15)
				(_nw68).ArraySet1(func(_pat_let8_0 D7) D7 {
					return func(_541_dt__update__tmp_h3 D7) D7 {
						return func(_pat_let9_0 _dafny.Set) D7 {
							return func(_542_dt__update_hcf41_h0 _dafny.Set) D7 {
								return Companion_D7_.Create_DC25_(_542_dt__update_hcf41_h0)
							}(_pat_let9_0)
						}(_dafny.SetOf(_pat_let_tv10, _pat_let_tv11))
					}(_pat_let8_0)
				}(_539_v145), 16)
				(_nw68).ArraySet1(_539_v145, 17)
				(_nw68).ArraySet1(_539_v145, 18)
				(_nw68).ArraySet1(_539_v145, 19)
				(_nw68).ArraySet1(_539_v145, 20)
				(_nw68).ArraySet1(_539_v145, 21)
				(_nw68).ArraySet1(Companion_D7_.Create_DC25_(_538_v144), 22)
				(_nw68).ArraySet1(_539_v145, 23)
				_540_v146 = _nw68
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_536_v142), 0))
				_ = _index60
				(_536_v142).ArraySet1(_540_v146, (_index60).Int())
				var _543_v147 _dafny.Sequence
				_ = _543_v147
				_543_v147 = _dafny.SeqOf(_540_v146, _540_v146)
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_536_v142), 0))
				_ = _index61
				(_536_v142).ArraySet1((_543_v147).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_515_v123).Contains(p0) {
						return (_515_v123).Get(p0).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-113)
				})(), _dafny.IntOfUint32((_543_v147).Cardinality()))).Uint32()).(_dafny.Array), (_index61).Int())
				var _544_v148 D2
				_ = _544_v148
				_544_v148 = Companion_D2_.Create_DC6_((_this).F29())
				var _545_v149 _dafny.Map
				_ = _545_v149
				_545_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F29()).Times(_dafny.IntOfInt64(-50)), ((_544_v148).Dtor_cf9()).Times((_this).F29()))
				_545_v149 = (_545_v149).Update(Companion_Default___.SafeModuloInt((_this).F29(), _dafny.IntOfUint32((_514_v122).Cardinality())), (_this).F29())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_492_v101), 0))
				_ = _index62
				(_492_v101).ArraySet1(p0, (_index62).Int())
				var _546_v150 _dafny.MultiSet
				_ = _546_v150
				_546_v150 = _dafny.MultiSetOf((_this).F29(), (_this).F29())
				var _547_v151 _dafny.Sequence
				_ = _547_v151
				_547_v151 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_514_v122).Cardinality())))
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_492_v101), 0))
				_ = _index63
				(_492_v101).ArraySet1(((_dafny.IntOfUint32((_514_v122).Cardinality())).Cmp((_this).Fm0(_546_v150, _dafny.IntOfUint32((_514_v122).Cardinality()), _dafny.Companion_Sequence_.Update(_547_v151, (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_547_v151).Cardinality()))).Uint32(), _dafny.SetOf((Companion_Default___.Fm33(globalState)).Cardinality())), !(p0), globalState)) >= 0) == (true), (_index63).Int())
				var _548_v152 _dafny.Array
				_ = _548_v152
				var _nw69 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw69
				_548_v152 = _nw69
				var _549_v153 T0
				_ = _549_v153
				var _nw70 *C1 = New_C1_()
				_ = _nw70
				_nw70.Ctor__((_dafny.Zero).Minus((_this).F29()), _dafny.SeqOf(_548_v152))
				_549_v153 = _nw70
				var _550_v154 _dafny.Array
				_ = _550_v154
				var _nwElement0_15 T0 = _549_v153
				_ = _nwElement0_15
				var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(16))
				_ = _nw71
				(_nw71).ArraySet1(_nwElement0_15, 0)
				(_nw71).ArraySet1(_549_v153, 1)
				(_nw71).ArraySet1(_549_v153, 2)
				(_nw71).ArraySet1(_549_v153, 3)
				(_nw71).ArraySet1(_549_v153, 4)
				(_nw71).ArraySet1(_549_v153, 5)
				(_nw71).ArraySet1(_549_v153, 6)
				(_nw71).ArraySet1(_549_v153, 7)
				(_nw71).ArraySet1(_549_v153, 8)
				(_nw71).ArraySet1(_549_v153, 9)
				(_nw71).ArraySet1(_549_v153, 10)
				(_nw71).ArraySet1(_549_v153, 11)
				(_nw71).ArraySet1(_549_v153, 12)
				(_nw71).ArraySet1(_549_v153, 13)
				(_nw71).ArraySet1(_549_v153, 14)
				(_nw71).ArraySet1(_549_v153, 15)
				_550_v154 = _nw71
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_550_v154), 0))
				_ = _index64
				(_550_v154).ArraySet1(_549_v153, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_550_v154), 0))
				_ = _index65
				(_550_v154).ArraySet1(_549_v153, (_index65).Int())
			}
		} else {
			var _551_v155 _dafny.Array
			_ = _551_v155
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw72
			_551_v155 = _nw72
			var _552_v156 _dafny.Sequence
			_ = _552_v156
			_552_v156 = _dafny.UnicodeSeqOfUtf8Bytes("rmsqxsyv")
			var _553_v157 _dafny.Sequence
			_ = _553_v157
			_553_v157 = _dafny.SeqOf(_552_v156, _552_v156)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))
			_ = _index66
			(_551_v155).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32(((_553_v157).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_553_v157).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), (_index66).Int())
			var _554_v158 _dafny.Map
			_ = _554_v158
			_554_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if p0 {
					return !(p0)
				}
				return p0
			})(), (_this).F29())
			var _555_v159 _dafny.Map
			_ = _555_v159
			_555_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), (_this).F29())
			var _556_v160 _dafny.Set
			_ = _556_v160
			_556_v160 = _dafny.SetOf((_this).F29())
			var _557_v161 _dafny.Sequence
			_ = _557_v161
			_557_v161 = _dafny.SeqOf((_556_v160).Cardinality())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))
			_ = _index67
			(_551_v155).ArraySet1((func() _dafny.Int {
				if (_554_v158).Contains(p0) {
					return (_554_v158).Get(p0).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_555_v159).Contains((_this).F29()) {
						return (_555_v159).Get((_this).F29()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_557_v161).Cardinality())
				})(), (_this).F29()))
			})(), (_index67).Int())
			var _558_v162 _dafny.Map
			_ = _558_v162
			_558_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.Companion_Sequence_.Update(_557_v161, (Companion_Default___.SafeIndex((_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_557_v161).Cardinality()))).Uint32(), (_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int))), _557_v161)
			var _559_v163 _dafny.MultiSet
			_ = _559_v163
			_559_v163 = _dafny.MultiSetOf(false)
			var _560_v164 _dafny.Sequence
			_ = _560_v164
			_560_v164 = _dafny.SeqOf(Companion_Default___.Fm24(_559_v163, (_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int))).Cardinality()), p0, globalState), _557_v161, _dafny.SeqOf((_this).F29()), _557_v161, _557_v161)
			var _561_v165 _dafny.Map
			_ = _561_v165
			_561_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_560_v164).Select((Companion_Default___.SafeIndex((_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_560_v164).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _562_v166 _dafny.Map
			_ = _562_v166
			_562_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _561_v165)
			_557_v161 = (func() _dafny.Sequence {
				if (_558_v162).Contains(((func() _dafny.Map {
					if (_562_v166).Contains((_this).F29()) {
						return (_562_v166).Get((_this).F29()).(_dafny.Map)
					}
					return _561_v165
				})()).Merge(_561_v165)) {
					return (_558_v162).Get(((func() _dafny.Map {
						if (_562_v166).Contains((_this).F29()) {
							return (_562_v166).Get((_this).F29()).(_dafny.Map)
						}
						return _561_v165
					})()).Merge(_561_v165)).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_557_v161, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(244))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}(func(_563_i13 _dafny.Int) _dafny.Int {
					return (_this).F29()
				})))
			})()
			var _564_v167 _dafny.Array
			_ = _564_v167
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_8
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) D3 = (func(_565_v159 _dafny.Map) func(_dafny.Int) D3 {
					return func(_566_i14 _dafny.Int) D3 {
						return Companion_D3_.Create_DC10_(_565_v159)
					}
				})(_555_v159)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw73 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw73).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw73).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_564_v167 = _nw73
			var _567_v168 D3
			_ = _567_v168
			_567_v168 = Companion_D3_.Create_DC10_(_555_v159)
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_564_v167), 0))
			_ = _index68
			(_564_v167).ArraySet1(_567_v168, (_index68).Int())
			var _568_v170 _dafny.MultiSet
			_ = _568_v170
			_568_v170 = _dafny.MultiSetOf(_dafny.IntOfUint32((_552_v156).Cardinality()))
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_564_v167), 0))
			_ = _index69
			(_564_v167).ArraySet1(Companion_D3_.Create_DC10_((_555_v159).Merge(func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter34 := _dafny.Iterate((_568_v170).Elements()); ; {
					_compr_33, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _569_v169 _dafny.Int
					_569_v169 = interface{}(_compr_33).(_dafny.Int)
					if (_568_v170).Contains(_569_v169) {
						_coll33.Add(Companion_Default___.SafeDivisionInt(_569_v169, (_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int)), (_this).F29())
					}
				}
				return _coll33.ToMap()
			}())), (_index69).Int())
			var _570_v171 _dafny.CodePoint
			_ = _570_v171
			_570_v171 = _dafny.CodePoint('s')
			var _571_v172 _dafny.Map
			_ = _571_v172
			_571_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-994))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_572_v171 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_573_i15 _dafny.Int) _dafny.CodePoint {
					return _572_v171
				}
			})(_570_v171))))
			var _574_v173 _dafny.Array
			_ = _574_v173
			var _nwElement0_16 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_552_v156, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_552_v156).Cardinality()), _dafny.IntOfUint32((_552_v156).Cardinality()))).Uint32(), _570_v171)
			_ = _nwElement0_16
			var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(6))
			_ = _nw74
			(_nw74).ArraySet1(_nwElement0_16, 0)
			(_nw74).ArraySet1(_552_v156, 1)
			(_nw74).ArraySet1((func() _dafny.Sequence {
				if (_571_v172).Contains(p0) {
					return (_571_v172).Get(p0).(_dafny.Sequence)
				}
				return _552_v156
			})(), 2)
			(_nw74).ArraySet1(_552_v156, 3)
			(_nw74).ArraySet1(_552_v156, 4)
			(_nw74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_552_v156, _dafny.UnicodeSeqOfUtf8Bytes("jwtvlaj")), 5)
			_574_v173 = _nw74
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_574_v173), 0))
			_ = _index70
			(_574_v173).ArraySet1(_552_v156, (_index70).Int())
			var _575_v174 _dafny.Map
			_ = _575_v174
			_575_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _555_v159)
			var _576_v175 _dafny.Map
			_ = _576_v175
			_576_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p0)
			var _577_v176 D0
			_ = _577_v176
			_577_v176 = Companion_D0_.Create_DC1_((_551_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_551_v155), 0))).Int()).(_dafny.Int), p0, _575_v174, (_576_v175).Cardinality(), _570_v171)
			var _578_v177 D4
			_ = _578_v177
			_578_v177 = Companion_D4_.Create_DC13_(_577_v176)
			var _579_v178 D4
			_ = _579_v178
			_579_v178 = Companion_D4_.Create_DC16_(_578_v177)
			var _580_v179 _dafny.Map
			_ = _580_v179
			_580_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
			var _581_v180 _dafny.Map
			_ = _581_v180
			_581_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_579_v178, (func() bool {
				if (_580_v179).Contains(false) {
					return (_580_v179).Get(false).(bool)
				}
				return p0
			})())
			var _582_v181 D10
			_ = _582_v181
			_582_v181 = Companion_D10_.Create_DC33_(_581_v180)
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_574_v173), 0))
			_ = _index71
			(_574_v173).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm30(globalState), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_582_v181).Dtor_cf54()).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm30(globalState)).Cardinality()))).Uint32(), _570_v171), (_index71).Int())
			r3 = (((_dafny.SetOf(_551_v155)).Cardinality()).Minus((_this).F29())).Plus((_dafny.Zero).Minus(((_this).F29()).Minus((_this).F29())))
		}
		var _583_v182 _dafny.Map
		_ = _583_v182
		_583_v182 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p0)
		var _584_v183 _dafny.Sequence
		_ = _584_v183
		_584_v183 = _dafny.UnicodeSeqOfUtf8Bytes("kdkvyfx")
		var _585_v184 _dafny.Map
		_ = _585_v184
		_585_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm37(_dafny.SeqOf(p0), _583_v182, _584_v183, globalState), _dafny.IntOfUint32((_584_v183).Cardinality()))
		var _586_v185 _dafny.Sequence
		_ = _586_v185
		_586_v185 = _dafny.SeqOf(_585_v184, _585_v184)
		var _587_v186 D9
		_ = _587_v186
		_587_v186 = Companion_D9_.Create_DC30_((_586_v185).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_586_v185).Cardinality()))).Uint32()).(_dafny.Map))
		var _588_v187 D9
		_ = _588_v187
		_588_v187 = Companion_D9_.Create_DC32_(_587_v186)
		var _source9 D9 = _588_v187
		_ = _source9
		if _source9.Is_DC31() {
			var _589___mcc_h11 _dafny.Int = _source9.Get_().(D9_DC31).Cf51
			_ = _589___mcc_h11
			var _590___mcc_h12 _dafny.Array = _source9.Get_().(D9_DC31).Cf52
			_ = _590___mcc_h12
			var _591_cf52 _dafny.Array = _590___mcc_h12
			_ = _591_cf52
			var _592_cf51 _dafny.Int = _589___mcc_h11
			_ = _592_cf51
			var _593_v188 _dafny.MultiSet
			_ = _593_v188
			_593_v188 = _dafny.MultiSetOf((_513_v121).Cardinality())
			var _594_v189 _dafny.Sequence
			_ = _594_v189
			_594_v189 = _dafny.SeqOf(_592_cf51)
			var _595_v190 _dafny.Set
			_ = _595_v190
			_595_v190 = _dafny.SetOf(_dafny.IntOfUint32((_594_v189).Cardinality()))
			var _596_v191 _dafny.Map
			_ = _596_v191
			_596_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _595_v190)
			var _597_v193 _dafny.Sequence
			_ = _597_v193
			_597_v193 = _dafny.SeqOf(_595_v190, _595_v190, _595_v190, _595_v190, (func() _dafny.Set {
				if (_596_v191).Contains(p0) {
					return (_596_v191).Get(p0).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter35 := _dafny.Iterate((_595_v190).Elements()); ; {
						_compr_34, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _598_v192 _dafny.Int
						_598_v192 = interface{}(_compr_34).(_dafny.Int)
						if (_595_v190).Contains(_598_v192) {
							_coll34.Add((_598_v192).Plus(_dafny.IntOfInt64(492)))
						}
					}
					return _coll34.ToSet()
				}()
			})())
			var _599_v194 _dafny.Sequence
			_ = _599_v194
			_599_v194 = _dafny.SeqOf(_592_cf51, (_this).Fm0(_593_v188, (_this).F29(), _597_v193, p0, globalState))
			var _600_v195 D2
			_ = _600_v195
			_600_v195 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_599_v194).Cardinality()))
			var _601_v196 D2
			_ = _601_v196
			_601_v196 = Companion_D2_.Create_DC9_(_600_v195)
			var _602_v197 _dafny.Set
			_ = _602_v197
			_602_v197 = _dafny.SetOf(_601_v196, _601_v196)
			_602_v197 = _602_v197
			var _603_v198 _dafny.Map
			_ = _603_v198
			_603_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_513_v121, _dafny.CodePoint('u'))
			var _604_v199 _dafny.CodePoint
			_ = _604_v199
			_604_v199 = _dafny.CodePoint('t')
			_603_v198 = (_603_v198).Update(_513_v121, _604_v199)
			var _605_v200 _dafny.Map
			_ = _605_v200
			_605_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_513_v121, p0)
			var _606_v201 D8
			_ = _606_v201
			_606_v201 = Companion_D8_.Create_DC29_(p0, (func() bool {
				if (_605_v200).Contains(_513_v121) {
					return (_605_v200).Get(_513_v121).(bool)
				}
				return p0
			})(), (_584_v183).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_584_v183).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_606_v201 = _606_v201
			_599_v194 = _599_v194
		} else if _source9.Is_DC30() {
			var _607___mcc_h13 _dafny.Map = _source9.Get_().(D9_DC30).Cf50
			_ = _607___mcc_h13
			var _608_cf50 _dafny.Map = _607___mcc_h13
			_ = _608_cf50
			r0 = p0
			var _609_v202 _dafny.Int
			_ = _609_v202
			var _610_v203 bool
			_ = _610_v203
			var _611_v204 bool
			_ = _611_v204
			var _612_v205 _dafny.Int
			_ = _612_v205
			var _out42 _dafny.Int
			_ = _out42
			var _out43 bool
			_ = _out43
			var _out44 bool
			_ = _out44
			var _out45 _dafny.Int
			_ = _out45
			_out42, _out43, _out44, _out45 = (_this).M5((_this).F29(), (_this).F29(), globalState)
			_609_v202 = _out42
			_610_v203 = _out43
			_611_v204 = _out44
			_612_v205 = _out45
			var _613_v206 _dafny.MultiSet
			_ = _613_v206
			_613_v206 = _dafny.MultiSetOf(_610_v203)
			var _614_v207 _dafny.Map
			_ = _614_v207
			_614_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_v203, _613_v206)
			var _615_v208 _dafny.Map
			_ = _615_v208
			_615_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(993), _614_v207)
			var _616_v209 _dafny.Sequence
			_ = _616_v209
			_616_v209 = _dafny.SeqOf(_610_v203)
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
			_ = _index72
			(_492_v101).ArraySet1(!((func() _dafny.Map {
				if (_615_v208).Contains(_609_v202) {
					return (_615_v208).Get(_609_v202).(_dafny.Map)
				}
				return _614_v207
			})()).Equals(Companion_Default___.Fm39((_this).F29(), _616_v209, globalState)), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
			_ = _index73
			var _rhs45 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_617_p0 bool) func(_dafny.Int) _dafny.CodePoint {
				return func(_618_i16 _dafny.Int) _dafny.CodePoint {
					return (func() _dafny.CodePoint {
						if _617_p0 {
							return _dafny.CodePoint('t')
						}
						return _dafny.CodePoint('l')
					})()
				}
			})(p0)))
			_ = _rhs45
			var _rhs46 _dafny.Int = ((_dafny.MultiSetOf(_dafny.IntOfInt64(737), _dafny.IntOfInt64(722), (_this).F29(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('b'))).Cardinality()))).Update(_609_v202, Companion_Default___.Abs(_dafny.IntOfInt64(-639)))).Cardinality()
			_ = _rhs46
			var _rhs47 _dafny.Set = _513_v121
			_ = _rhs47
			var _rhs48 bool = !(_611_v204)
			_ = _rhs48
			var _rhs49 _dafny.Sequence = _584_v183
			_ = _rhs49
			var _lhs38 *GlobalState = globalState
			_ = _lhs38
			var _lhs39 _dafny.Array = _492_v101
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
			_ = _lhs40
			_584_v183 = _rhs45
			_lhs38.F11 = _rhs46
			_513_v121 = _rhs47
			(_lhs39).ArraySet1(_rhs48, (_lhs40).Int())
			_584_v183 = _rhs49
			if Companion_Default___.Fm37(_616_v209, _583_v182, _584_v183, globalState) {
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
				_ = _index74
				(_492_v101).ArraySet1(p0, (_index74).Int())
				var _619_v210 _dafny.Set
				_ = _619_v210
				_619_v210 = _dafny.SetOf(_dafny.IntOfInt64(-516), (_this).F29())
				var _620_v212 _dafny.Sequence
				_ = _620_v212
				_620_v212 = _dafny.SeqOf(_dafny.IntOfUint32((_584_v183).Cardinality()))
				var _621_v213 _dafny.Array
				_ = _621_v213
				var _nwElement0_17 _dafny.Set = _619_v210
				_ = _nwElement0_17
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(22))
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_17, 0)
				(_nw75).ArraySet1((_619_v210).Difference(_dafny.SetOf(_dafny.IntOfUint32((_584_v183).Cardinality()))), 1)
				(_nw75).ArraySet1(_619_v210, 2)
				(_nw75).ArraySet1(_619_v210, 3)
				(_nw75).ArraySet1(_dafny.SetOf((_this).F29(), _612_v205, _dafny.IntOfInt64(60), _dafny.IntOfUint32((_584_v183).Cardinality()), _609_v202), 4)
				(_nw75).ArraySet1(_619_v210, 5)
				(_nw75).ArraySet1(_dafny.SetOf(_609_v202), 6)
				(_nw75).ArraySet1(func() _dafny.Set {
					var _coll35 = _dafny.NewBuilder()
					_ = _coll35
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-907), _dafny.IntOfInt64(-883))); ; {
						_compr_35, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _622_v211 _dafny.Int
						_622_v211 = interface{}(_compr_35).(_dafny.Int)
						if ((_dafny.IntOfInt64(-907)).Cmp(_622_v211) <= 0) && ((_622_v211).Cmp(_dafny.IntOfInt64(-883)) < 0) {
							_coll35.Add((_622_v211).Times((_513_v121).Cardinality()))
						}
					}
					return _coll35.ToSet()
				}(), 7)
				(_nw75).ArraySet1((_619_v210).Difference(_619_v210), 8)
				(_nw75).ArraySet1((_dafny.SetOf(_612_v205)).Union(_619_v210), 9)
				(_nw75).ArraySet1(_619_v210, 10)
				(_nw75).ArraySet1((_619_v210).Difference(_619_v210), 11)
				(_nw75).ArraySet1(_619_v210, 12)
				(_nw75).ArraySet1(Companion_Default___.Fm23(_609_v202, _609_v202, p0, (_492_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))).Int()).(bool), globalState), 13)
				(_nw75).ArraySet1(_619_v210, 14)
				(_nw75).ArraySet1((_619_v210).Difference(_619_v210), 15)
				(_nw75).ArraySet1(_619_v210, 16)
				(_nw75).ArraySet1(_619_v210, 17)
				(_nw75).ArraySet1((Companion_Default___.Fm23((_619_v210).Cardinality(), (_620_v212).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_620_v212).Cardinality()))).Uint32()).(_dafny.Int), true, _611_v204, globalState)).Intersection(_619_v210), 18)
				(_nw75).ArraySet1(_619_v210, 19)
				(_nw75).ArraySet1(_619_v210, 20)
				(_nw75).ArraySet1(_619_v210, 21)
				_621_v213 = _nw75
				var _623_v214 _dafny.Map
				_ = _623_v214
				_623_v214 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_584_v183, (_this).F29())
				var _624_v215 _dafny.Map
				_ = _624_v215
				_624_v215 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v202, _623_v214)
				var _625_v216 D5
				_ = _625_v216
				_625_v216 = Companion_D5_.Create_DC17_((func() _dafny.Map {
					if (_624_v215).Contains(_612_v205) {
						return (_624_v215).Get(_612_v205).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_584_v183, (_this).F29())
				})())
				var _626_v217 D5
				_ = _626_v217
				_626_v217 = Companion_D5_.Create_DC20_(_625_v216)
				var _627_v218 _dafny.CodePoint
				_ = _627_v218
				_627_v218 = _dafny.CodePoint('h')
				var _628_v219 D10
				_ = _628_v219
				_628_v219 = Companion_D10_.Create_DC34_(_627_v218, (_492_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))).Int()).(bool), _611_v204, p0, false)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
				_ = _index75
				var _rhs50 _dafny.Array = _621_v213
				_ = _rhs50
				var _rhs51 _dafny.Int = Companion_Default___.SafeModuloInt(_612_v205, _609_v202)
				_ = _rhs51
				var _rhs52 D5 = _626_v217
				_ = _rhs52
				var _rhs53 bool = (_628_v219).Dtor_cf57()
				_ = _rhs53
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 _dafny.Array = _492_v101
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
				_ = _lhs43
				_621_v213 = _rhs50
				_lhs41.F6 = _rhs51
				_626_v217 = _rhs52
				(_lhs42).ArraySet1(_rhs53, (_lhs43).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))
				_ = _index76
				(_492_v101).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("f"), _584_v183), (_index76).Int())
				var _629_v220 _dafny.Map
				_ = _629_v220
				_629_v220 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v204, (_492_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))).Int()).(bool))
				(globalState).F10 = (_609_v202).Cmp((_629_v220).Cardinality()) < 0
				(globalState).F8 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_620_v212).Cardinality()), _609_v202), (func() _dafny.Int {
					if (_608_cf50).Contains(p0) {
						return (_608_cf50).Get(p0).(_dafny.Int)
					}
					return (_this).F29()
				})())
			} else {
				_584_v183 = _584_v183
				_610_v203 = _611_v204
				(globalState).F0 = !(true) || ((_492_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_492_v101), 0))).Int()).(bool))
				var _630_v221 *C0
				_ = _630_v221
				var _nw76 *C0 = New_C0_()
				_ = _nw76
				_nw76.Ctor__()
				_630_v221 = _nw76
				var _631_v222 _dafny.Array
				_ = _631_v222
				var _nw77 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw77
				_631_v222 = _nw77
				var _632_v223 _dafny.Map
				_ = _632_v223
				_632_v223 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v202, _dafny.SeqOf(_631_v222))
				var _633_v224 *C1
				_ = _633_v224
				var _nw78 *C1 = New_C1_()
				_ = _nw78
				_nw78.Ctor__((_this).F29(), _dafny.Companion_Sequence_.Concatenate((_this).F22(), (func() _dafny.Sequence {
					if (_632_v223).Contains(_612_v205) {
						return (_632_v223).Get(_612_v205).(_dafny.Sequence)
					}
					return _dafny.SeqOf(_631_v222, _631_v222, _631_v222, _631_v222)
				})()))
				_633_v224 = _nw78
			}
		} else {
			var _634___mcc_h14 D9 = _source9.Get_().(D9_DC32).Cf53
			_ = _634___mcc_h14
			var _635_cf53 D9 = _634___mcc_h14
			_ = _635_cf53
			(globalState).F16 = (func() _dafny.Int {
				if !(p0) {
					return (_this).F29()
				}
				return (_dafny.MultiSetOf(!(p0))).Cardinality()
			})()
			(globalState).F11 = (_this).F29()
			var _636_v225 _dafny.Map
			_ = _636_v225
			_636_v225 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
			var _637_v226 _dafny.Map
			_ = _637_v226
			_637_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_636_v225).Merge(_636_v225), p0)
			_637_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if p0 {
					return _636_v225
				}
				return _636_v225
			})(), ((_this).F29()).Cmp(_dafny.IntOfUint32((_584_v183).Cardinality())) >= 0)
			var _638_v227 D9
			_ = _638_v227
			_638_v227 = Companion_D9_.Create_DC31_((_this).F29(), _492_v101)
			var _639_v228 _dafny.Map
			_ = _639_v228
			_639_v228 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_636_v225).Contains(p0) {
					return (_636_v225).Get(p0).(bool)
				}
				return p0
			})(), _638_v227)
			var _640_v229 _dafny.MultiSet
			_ = _640_v229
			_640_v229 = _dafny.MultiSetOf((_639_v228).Cardinality())
			var _641_v230 _dafny.Map
			_ = _641_v230
			_641_v230 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_492_v101, _640_v229)
			var _642_v231 _dafny.CodePoint
			_ = _642_v231
			_642_v231 = _dafny.CodePoint('q')
			var _643_v232 _dafny.Sequence
			_ = _643_v232
			_643_v232 = _dafny.SeqOf(_640_v229)
			var _644_v233 D11
			_ = _644_v233
			_644_v233 = Companion_D11_.Create_DC36_(_643_v232)
			var _645_v234 _dafny.Sequence
			_ = _645_v234
			_645_v234 = _dafny.SeqOf(p0, p0)
			var _rhs54 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_644_v233).Dtor_cf62(), _643_v232), _643_v232)).Cardinality())
			_ = _rhs54
			var _rhs55 _dafny.Map = (_641_v230).Update(_492_v101, (_640_v229).Update((_this).F29(), Companion_Default___.Abs((_this).F29())))
			_ = _rhs55
			var _rhs56 bool = p0
			_ = _rhs56
			var _rhs57 _dafny.CodePoint = Companion_Default___.Fm22(p0, (_645_v234).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_645_v234).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(667), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F29(), _dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality()))), globalState)
			_ = _rhs57
			var _rhs58 bool = !(!(p0)) || (p0)
			_ = _rhs58
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			_lhs44.F6 = _rhs54
			_641_v230 = _rhs55
			_lhs45.F0 = _rhs56
			_642_v231 = _rhs57
			_lhs46.F10 = _rhs58
		}
		r0 = p0
		r1 = (func() _dafny.Array {
			if !(p0) {
				return _492_v101
			}
			return _492_v101
		})()
		var _646_v235 T0
		_ = _646_v235
		var _nw79 *C1 = New_C1_()
		_ = _nw79
		_nw79.Ctor__((_this).F29(), (_this).F22())
		_646_v235 = _nw79
		var _647_v236 _dafny.Map
		_ = _647_v236
		_647_v236 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _646_v235)
		var _648_v237 _dafny.Array
		_ = _648_v237
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
		_ = _nw80
		_648_v237 = _nw80
		var _649_v238 _dafny.Map
		_ = _649_v238
		_649_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _648_v237)
		var _650_v239 _dafny.Sequence
		_ = _650_v239
		_650_v239 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F29()), (_647_v236).Cardinality(), (_649_v238).Cardinality(), (_this).F29())
		var _651_v240 _dafny.CodePoint
		_ = _651_v240
		_651_v240 = _dafny.CodePoint('y')
		var _652_v241 _dafny.Map
		_ = _652_v241
		_652_v241 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_650_v239, _651_v240)
		r2 = _652_v241
		var _653_v242 _dafny.MultiSet
		_ = _653_v242
		_653_v242 = _dafny.MultiSetOf((_this).F29(), (_this).F29(), _dafny.IntOfInt64(-818))
		var _654_v243 _dafny.Map
		_ = _654_v243
		_654_v243 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _655_v244 _dafny.Set
		_ = _655_v244
		_655_v244 = _dafny.SetOf((_this).F29(), (_this).F29())
		var _656_v245 _dafny.Sequence
		_ = _656_v245
		_656_v245 = _dafny.SeqOf(_655_v244)
		r3 = Companion_Default___.SafeDivisionInt((_this).F29(), (_646_v235).Fm0(_653_v242, (_654_v243).Cardinality(), _656_v245, p0, globalState))
		return r0, r1, r2, r3
	}
}
func (_this *C2) M5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _657_v0 bool
		_ = _657_v0
		_657_v0 = true
		var _658_v1 _dafny.MultiSet
		_ = _658_v1
		_658_v1 = _dafny.MultiSetOf(_657_v0)
		var _659_v2 _dafny.Sequence
		_ = _659_v2
		_659_v2 = _dafny.SeqOf(_658_v1)
		r3 = ((_659_v2).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrquxsrl")).Cardinality()), p0), _dafny.IntOfUint32((_659_v2).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()
		var _660_v3 _dafny.MultiSet
		_ = _660_v3
		_660_v3 = _dafny.MultiSetOf(p1)
		var _661_v4 _dafny.Set
		_ = _661_v4
		_661_v4 = _dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(804))).Cardinality(), p1)
		var _662_v5 _dafny.Sequence
		_ = _662_v5
		_662_v5 = _dafny.SeqOf(_661_v4, _dafny.SetOf((_this).F29()), _661_v4)
		var _663_v6 _dafny.Sequence
		_ = _663_v6
		_663_v6 = _dafny.SeqOf((_this).Fm0(_660_v3, (_this).F29(), _662_v5, _657_v0, globalState))
		var _664_v7 _dafny.MultiSet
		_ = _664_v7
		_664_v7 = _dafny.MultiSetOf(_dafny.MultiSetOf(p1), _dafny.MultiSetFromSeq(_663_v6))
		var _665_v8 _dafny.Array
		_ = _665_v8
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw81
		_665_v8 = _nw81
		var _666_v9 D9
		_ = _666_v9
		_666_v9 = Companion_D9_.Create_DC31_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_664_v7).Cardinality(), _dafny.IntOfInt64(927))).Cardinality(), _665_v8)
		var _667_v10 D9
		_ = _667_v10
		_667_v10 = Companion_D9_.Create_DC32_(_666_v9)
		_667_v10 = _667_v10
		var _668_v11 _dafny.Sequence
		_ = _668_v11
		_668_v11 = _dafny.UnicodeSeqOfUtf8Bytes("mcwffodmy")
		var _rhs59 _dafny.Sequence = _668_v11
		_ = _rhs59
		var _rhs60 _dafny.MultiSet = _660_v3
		_ = _rhs60
		_668_v11 = _rhs59
		_660_v3 = _rhs60
		if _657_v0 {
			var _669_v13 _dafny.Sequence
			_ = _669_v13
			_669_v13 = _dafny.SeqOf(_657_v0, true)
			var _670_v14 _dafny.Sequence
			_ = _670_v14
			_670_v14 = _dafny.SeqOf(_669_v13, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_657_v0, _657_v0, _657_v0, _657_v0, _657_v0), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf(_657_v0, _657_v0, _657_v0, _657_v0, _657_v0)).Cardinality()))).Uint32(), _657_v0), _669_v13, _669_v13)
			var _671_v15 _dafny.Map
			_ = _671_v15
			_671_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v3, _657_v0)
			var _672_v16 _dafny.Array
			_ = _672_v16
			var _nwElement0_18 _dafny.Int = (func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter37 := _dafny.Iterate((_661_v4).Elements()); ; {
					_compr_36, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _673_v12 _dafny.Int
					_673_v12 = interface{}(_compr_36).(_dafny.Int)
					if (_661_v4).Contains(_673_v12) {
						_coll36.Add((_673_v12).Times(p0), (_668_v11).Select((Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_668_v11).Cardinality()))).Uint32()).(_dafny.CodePoint))
					}
				}
				return _coll36.ToMap()
			}()).Cardinality()
			_ = _nwElement0_18
			var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(20))
			_ = _nw82
			(_nw82).ArraySet1(_nwElement0_18, 0)
			(_nw82).ArraySet1((_dafny.IntOfInt64(275)).Plus(p0), 1)
			(_nw82).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_670_v14, _dafny.SeqOf(_669_v13))).Cardinality()), 2)
			(_nw82).ArraySet1((func() _dafny.Int {
				if (_658_v1).Contains(_657_v0) {
					return (_658_v1).Multiplicity(_657_v0)
				}
				return (_658_v1).Cardinality()
			})(), 3)
			(_nw82).ArraySet1((p1).Plus(p1), 4)
			(_nw82).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wudmpyivk")).Cardinality())).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_674_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))).Cardinality()))), 5)
			(_nw82).ArraySet1(p1, 6)
			(_nw82).ArraySet1(p1, 7)
			(_nw82).ArraySet1(p0, 8)
			(_nw82).ArraySet1((func() _dafny.Int {
				if _657_v0 {
					return p1
				}
				return p0
			})(), 9)
			(_nw82).ArraySet1((_this).F29(), 10)
			(_nw82).ArraySet1((_660_v3).Cardinality(), 11)
			(_nw82).ArraySet1(p0, 12)
			(_nw82).ArraySet1(p1, 13)
			(_nw82).ArraySet1(_dafny.IntOfInt64(983), 14)
			(_nw82).ArraySet1(p1, 15)
			(_nw82).ArraySet1(p1, 16)
			(_nw82).ArraySet1((_671_v15).Cardinality(), 17)
			(_nw82).ArraySet1((_this).F29(), 18)
			(_nw82).ArraySet1((_dafny.IntOfInt64(100)).Minus(p1), 19)
			_672_v16 = _nw82
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_672_v16), 0))
			_ = _index77
			(_672_v16).ArraySet1((_this).F29(), (_index77).Int())
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_672_v16), 0))
			_ = _index78
			(_672_v16).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(350))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}((func(_675_v0 bool, _676_v11 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
				return func(_677_i1 _dafny.Int) _dafny.CodePoint {
					return (func() _dafny.CodePoint {
						if !(_675_v0) {
							return (_676_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.IntOfUint32((_676_v11).Cardinality()))).Uint32()).(_dafny.CodePoint)
						}
						return _dafny.CodePoint('h')
					})()
				}
			})(_657_v0, _668_v11)))).Cardinality()), (_index78).Int())
			var _678_v17 T0
			_ = _678_v17
			var _nw83 *C1 = New_C1_()
			_ = _nw83
			_nw83.Ctor__((_672_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_672_v16), 0))).Int()).(_dafny.Int), (_this).F22())
			_678_v17 = _nw83
			var _679_v18 _dafny.MultiSet
			_ = _679_v18
			_679_v18 = _dafny.MultiSetOf(_678_v17)
			_679_v18 = _679_v18
			var _680_v19 _dafny.MultiSet
			_ = _680_v19
			_680_v19 = _dafny.MultiSetOf(_665_v8)
			_680_v19 = ((_680_v19).Union(_680_v19)).Intersection(_dafny.MultiSetOf(_665_v8))
			(globalState).F6 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _657_v0 {
					return _668_v11
				}
				return _668_v11
			})(), _668_v11)).Cardinality())
			_668_v11 = _668_v11
		} else {
			var _681_v20 _dafny.Sequence
			_ = _681_v20
			_681_v20 = _dafny.SeqOf(_657_v0, !(_657_v0), _657_v0)
			var _682_v21 _dafny.Map
			_ = _682_v21
			_682_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.SafeModuloInt((_661_v4).Cardinality(), _dafny.IntOfUint32((_681_v20).Cardinality())))
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))
			_ = _index79
			(_665_v8).ArraySet1(_657_v0, (_index79).Int())
			var _683_v23 _dafny.Map
			_ = _683_v23
			_683_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_660_v3).Cardinality(), _657_v0)
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))
			_ = _index80
			var _rhs61 _dafny.Int = _dafny.IntOfInt64(499)
			_ = _rhs61
			var _rhs62 _dafny.Map = func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-244), _dafny.IntOfInt64(-210))); ; {
					_compr_37, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _684_v22 _dafny.Int
					_684_v22 = interface{}(_compr_37).(_dafny.Int)
					if ((_dafny.IntOfInt64(-244)).Cmp(_684_v22) <= 0) && ((_684_v22).Cmp(_dafny.IntOfInt64(-210)) < 0) {
						_coll37.Add(Companion_Default___.SafeModuloInt(_684_v22, (_this).F29()), (_682_v21).Cardinality())
					}
				}
				return _coll37.ToMap()
			}()
			_ = _rhs62
			var _rhs63 bool = (func() bool {
				if _657_v0 {
					return (_681_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.IntOfUint32((_681_v20).Cardinality()))).Uint32()).(bool)
				}
				return ((_661_v4).Cardinality()).Cmp(_dafny.IntOfUint32((_681_v20).Cardinality())) > 0
			})()
			_ = _rhs63
			var _rhs64 bool = (func() bool {
				if (_683_v23).Contains((_this).F29()) {
					return (_683_v23).Get((_this).F29()).(bool)
				}
				return !(_657_v0)
			})()
			_ = _rhs64
			var _lhs47 *GlobalState = globalState
			_ = _lhs47
			var _lhs48 _dafny.Array = _665_v8
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))
			_ = _lhs49
			_lhs47.F6 = _rhs61
			_682_v21 = _rhs62
			r1 = _rhs63
			(_lhs48).ArraySet1(_rhs64, (_lhs49).Int())
			(globalState).F16 = _dafny.IntOfInt64(614)
			if ((func() _dafny.Map {
				if !(false) {
					return Companion_Default___.Fm40((_this).F29(), p0, (_668_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_668_v11).Cardinality()))).Uint32()).(_dafny.CodePoint), globalState)
				}
				return _683_v23
			})()).Contains((_this).F29()) {
				(globalState).F0 = _657_v0
				var _685_v24 _dafny.MultiSet
				_ = _685_v24
				_685_v24 = _dafny.MultiSetOf(_663_v6, _dafny.SeqOf((_this).F29()))
				var _686_v25 _dafny.Set
				_ = _686_v25
				_686_v25 = _dafny.SetOf((_665_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))).Int()).(bool), _657_v0)
				var _rhs65 _dafny.MultiSet = _685_v24
				_ = _rhs65
				var _rhs66 _dafny.MultiSet = _658_v1
				_ = _rhs66
				var _rhs67 _dafny.MultiSet = Companion_Default___.Fm41(p0, _668_v11, _686_v25, globalState)
				_ = _rhs67
				_685_v24 = _rhs65
				_658_v1 = _rhs66
				_658_v1 = _rhs67
				var _687_v26 _dafny.Array
				_ = _687_v26
				var _nwElement0_19 _dafny.Int = p1
				_ = _nwElement0_19
				var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(4))
				_ = _nw84
				(_nw84).ArraySet1(_nwElement0_19, 0)
				(_nw84).ArraySet1(p0, 1)
				(_nw84).ArraySet1(p1, 2)
				(_nw84).ArraySet1(p0, 3)
				_687_v26 = _nw84
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_9
				var _nw85 _dafny.Array
				_ = _nw85
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw85 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Int = (func(_688_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_689_i2 _dafny.Int) _dafny.Int {
							return (_689_i2).Minus(_688_p1)
						}
					})(p1)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw85 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw85).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw85).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_687_v26 = _nw85
				var _690_v27 _dafny.Array
				_ = _690_v27
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw86
				_690_v27 = _nw86
				var _rhs68 _dafny.Array = _690_v27
				_ = _rhs68
				var _rhs69 bool = (_665_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))).Int()).(bool)
				_ = _rhs69
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				_690_v27 = _rhs68
				_lhs50.F10 = _rhs69
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))
				_ = _index81
				(_665_v8).ArraySet1(_657_v0, (_index81).Int())
			} else {
				var _691_v28 _dafny.Array
				_ = _691_v28
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_10
				var _nw87 _dafny.Array
				_ = _nw87
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw87 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = (func(_692_v0 bool) func(_dafny.Int) bool {
						return func(_693_i3 _dafny.Int) bool {
							return _692_v0
						}
					})(_657_v0)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw87 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw87).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw87).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_691_v28 = _nw87
				var _694_v29 _dafny.Map
				_ = _694_v29
				_694_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v28, p1)
				var _695_v30 _dafny.Map
				_ = _695_v30
				_695_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_663_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_681_v20).Cardinality()), _dafny.IntOfUint32((_663_v6).Cardinality()))).Uint32()).(_dafny.Int), _694_v29)
				_694_v29 = (func() _dafny.Map {
					if (_695_v30).Contains((p1).Times(_dafny.IntOfInt64(613))) {
						return (_695_v30).Get((p1).Times(_dafny.IntOfInt64(613))).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v28, p0)
				})()
				(globalState).F16 = ((_this).F29()).Plus((_this).F29())
				var _696_v31 *C1
				_ = _696_v31
				var _nw88 *C1 = New_C1_()
				_ = _nw88
				_nw88.Ctor__(Companion_Default___.SafeModuloInt(p0, (_this).F29()), _dafny.SeqOf(_691_v28, _691_v28, _691_v28))
				_696_v31 = _nw88
				var _697_v32 _dafny.Array
				_ = _697_v32
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
				_ = _nw89
				_697_v32 = _nw89
				_697_v32 = _697_v32
				var _698_v33 _dafny.Array
				_ = _698_v33
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_11
				var _nw90 _dafny.Array
				_ = _nw90
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw90 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Int = (func(_699_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_700_i4 _dafny.Int) _dafny.Int {
							return (_700_i4).Minus(_699_p1)
						}
					})(p1)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw90 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw90).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw90).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_698_v33 = _nw90
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_698_v33), 0))
				_ = _index82
				(_698_v33).ArraySet1((p0).Minus(p0), (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_698_v33), 0))
				_ = _index83
				(_698_v33).ArraySet1(_dafny.IntOfInt64(917), (_index83).Int())
			}
			var _701_v34 *C0
			_ = _701_v34
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__()
			_701_v34 = _nw91
			var _702_v35 _dafny.MultiSet
			_ = _702_v35
			_702_v35 = _dafny.MultiSetOf(_701_v34)
			var _703_v36 D12
			_ = _703_v36
			_703_v36 = Companion_D12_.Create_DC41_(_702_v35)
			var _704_v37 *C1
			_ = _704_v37
			var _nw92 *C1 = New_C1_()
			_ = _nw92
			_nw92.Ctor__((func() _dafny.Int {
				if (_665_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))).Int()).(bool) {
					return (_this).Fm0(_660_v3, ((_703_v36).Dtor_cf68()).Cardinality(), _662_v5, false, globalState)
				}
				return p1
			})(), (_this).F22())
			_704_v37 = _nw92
			_657_v0 = (func() bool {
				if (_683_v23).Contains((p0).Minus((_this).F29())) {
					return (_683_v23).Get((p0).Minus((_this).F29())).(bool)
				}
				return (_665_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_665_v8), 0))).Int()).(bool)
			})()
		}
		var _705_i5 _dafny.Int
		_ = _705_i5
		_705_i5 = _dafny.Zero
		{
			for !(Companion_Default___.Fm37(_dafny.SeqOf(_657_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(_657_v0)), _668_v11, globalState)) {
				{
					if (_705_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_705_i5 = (_705_i5).Plus(_dafny.One)
					var _706_v38 D12
					_ = _706_v38
					_706_v38 = Companion_D12_.Create_DC42_((_this).F29())
					var _707_v39 _dafny.Sequence
					_ = _707_v39
					_707_v39 = _dafny.SeqOf(!(true), _657_v0, _657_v0)
					var _708_v40 _dafny.Set
					_ = _708_v40
					_708_v40 = _dafny.SetOf(_706_v38, _706_v38, (func() D12 {
						if (_707_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_668_v11).Cardinality()), _dafny.IntOfUint32((_707_v39).Cardinality()))).Uint32()).(bool) {
							return _706_v38
						}
						return Companion_D12_.Create_DC42_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("agnqcr")).Cardinality()))
					})())
					_708_v40 = _708_v40
					(globalState).F0 = !((_707_v39).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_707_v39).Cardinality()))).Uint32()).(bool)) || (!((func() bool {
						if _657_v0 {
							return _657_v0
						}
						return _657_v0
					})()))
					var _709_v41 _dafny.Map
					_ = _709_v41
					_709_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg55 _dafny.Int) interface{} {
							return coer55(arg55)
						}
					}(func(_710_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					})), p0)
					var _711_v42 D5
					_ = _711_v42
					_711_v42 = Companion_D5_.Create_DC17_((_709_v41).Merge(_709_v41))
					var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_665_v8), 0))
					_ = _index84
					(_665_v8).ArraySet1(_657_v0, (_index84).Int())
					var _712_v43 _dafny.Map
					_ = _712_v43
					_712_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _657_v0)
					var _713_v44 _dafny.Map
					_ = _713_v44
					_713_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_712_v43).Cardinality(), _657_v0)
					var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_665_v8), 0))
					_ = _index85
					var _rhs70 D5 = Companion_Default___.Fm42(globalState)
					_ = _rhs70
					var _rhs71 bool = !(((_dafny.Zero).Minus(p1)).Cmp(p1) < 0) || (Companion_Default___.Fm37(_dafny.SeqOf(_657_v0), _713_v44, _668_v11, globalState))
					_ = _rhs71
					var _rhs72 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_663_v6, _dafny.SeqOf(_dafny.IntOfInt64(696)))
					_ = _rhs72
					var _rhs73 bool = !(_657_v0)
					_ = _rhs73
					var _lhs51 *GlobalState = globalState
					_ = _lhs51
					var _lhs52 _dafny.Array = _665_v8
					_ = _lhs52
					var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_665_v8), 0))
					_ = _lhs53
					_711_v42 = _rhs70
					_lhs51.F10 = _rhs71
					_663_v6 = _rhs72
					(_lhs52).ArraySet1(_rhs73, (_lhs53).Int())
					(globalState).F11 = (_dafny.IntOfInt64(-412)).Minus(_dafny.IntOfInt64(621))
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _714_v45 _dafny.Sequence
		_ = _714_v45
		_714_v45 = _dafny.SeqOf(_657_v0)
		(globalState).F5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_714_v45, (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32((_714_v45).Cardinality()))).Uint32(), _657_v0), _714_v45), _714_v45)
		var _715_v46 _dafny.CodePoint
		_ = _715_v46
		_715_v46 = _dafny.CodePoint('c')
		var _716_v47 _dafny.Map
		_ = _716_v47
		_716_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _715_v46)
		r0 = ((_663_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_663_v6).Cardinality()))).Uint32()).(_dafny.Int)).Times(((_716_v47).Cardinality()).Minus(p0))
		r1 = _657_v0
		r2 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(499))).Uint32(), func(coer56 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}((func(_717_v1 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
			return func(_718_i7 _dafny.Int) _dafny.MultiSet {
				return _717_v1
			}
		})(_658_v1))), (_658_v1).Intersection(_658_v1))
		r3 = (func() _dafny.Int {
			if _657_v0 {
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(p0))
			}
			return (_this).F29()
		})()
		return r0, r1, r2, r3
	}
}
func (_this *C2) F29() _dafny.Int {
	{
		return _this._f29
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f22 _dafny.Sequence
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f22 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F22() _dafny.Sequence {
	return _this._f22
}
func (_this *C3) Ctor__(f22 _dafny.Sequence) {
	{
		(_this)._f22 = f22
	}
}
func (_this *C3) Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-883), _dafny.IntOfInt64(290))).Minus(_dafny.IntOfInt64(-935))
	}
}
func (_this *C3) M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _719_v0 _dafny.Int
		_ = _719_v0
		_719_v0 = _dafny.IntOfInt64(-487)
		var _720_v1 _dafny.Sequence
		_ = _720_v1
		_720_v1 = _dafny.UnicodeSeqOfUtf8Bytes("f")
		var _721_v2 _dafny.Sequence
		_ = _721_v2
		_721_v2 = _dafny.SeqOf(_719_v0, _dafny.IntOfUint32((_720_v1).Cardinality()))
		var _722_v3 _dafny.MultiSet
		_ = _722_v3
		_722_v3 = _dafny.MultiSetOf(Companion_Default___.Fm14(_dafny.IntOfInt64(-870), _719_v0, globalState), _721_v2)
		var _723_v4 _dafny.Sequence
		_ = _723_v4
		_723_v4 = _dafny.SeqOf((_722_v3).IsProperSubsetOf(_722_v3), p0)
		var _724_v5 _dafny.Set
		_ = _724_v5
		_724_v5 = _dafny.SetOf(_719_v0, _719_v0)
		var _725_v6 _dafny.Sequence
		_ = _725_v6
		_725_v6 = _dafny.SeqOf(_724_v5, _724_v5, _724_v5)
		(globalState).F10 = (_723_v4).Select((Companion_Default___.SafeIndex((_this).Fm0((_dafny.MultiSetOf(_719_v0)).Update(_719_v0, Companion_Default___.Abs(_719_v0)), _719_v0, _725_v6, p0, globalState), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool)
		var _726_i0 _dafny.Int
		_ = _726_i0
		_726_i0 = _dafny.Zero
		{
			for Companion_Default___.Fm15((_719_v0).Cmp(_719_v0) >= 0, globalState) {
				{
					if (_726_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_726_i0 = (_726_i0).Plus(_dafny.One)
					r0 = (_719_v0).Cmp((_dafny.IntOfInt64(433)).Plus(_dafny.IntOfInt64(-409))) >= 0
					(globalState).F10 = !((_723_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool))
					var _727_v7 _dafny.MultiSet
					_ = _727_v7
					_727_v7 = _dafny.MultiSetOf(_dafny.IntOfInt64(100))
					(globalState).F11 = (((_727_v7).Cardinality()).Minus(_719_v0)).Plus(_719_v0)
					var _728_v8 _dafny.Map
					_ = _728_v8
					_728_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
					_728_v8 = _728_v8
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _729_v9 _dafny.Set
		_ = _729_v9
		_729_v9 = _dafny.SetOf(p0)
		(globalState).F0 = (_729_v9).IsProperSubsetOf((_729_v9).Intersection(_729_v9))
		var _730_i1 _dafny.Int
		_ = _730_i1
		_730_i1 = _dafny.Zero
		{
			for !((_723_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool)) {
				{
					if (_730_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_730_i1 = (_730_i1).Plus(_dafny.One)
					var _731_v10 _dafny.Map
					_ = _731_v10
					_731_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_721_v2).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_721_v2).Cardinality()))).Uint32()).(_dafny.Int), _723_v4)
					(globalState).F11 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _723_v4)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(65), _723_v4)).Merge(_731_v10))).Cardinality()
					var _732_v11 _dafny.CodePoint
					_ = _732_v11
					_732_v11 = _dafny.CodePoint('s')
					var _733_v12 D4
					_ = _733_v12
					_733_v12 = Companion_D4_.Create_DC14_(_732_v11, ((_this).F22()).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32(((_this).F22()).Cardinality()))).Uint32()).(_dafny.Array), _720_v1)
					var _734_v13 D4
					_ = _734_v13
					_734_v13 = Companion_D4_.Create_DC16_(_733_v12)
					var _735_v14 D4
					_ = _735_v14
					_735_v14 = Companion_D4_.Create_DC16_(_733_v12)
					_735_v14 = _735_v14
					(globalState).F10 = p0
					if ((func() bool {
						if p0 {
							return p0
						}
						return p0
					})()) && ((_723_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool)) {
						_720_v1 = _dafny.UnicodeSeqOfUtf8Bytes("hdmeaxjx")
						r3 = _dafny.IntOfInt64(284)
						var _736_v15 _dafny.Array
						_ = _736_v15
						var _nw93 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(21))
						_ = _nw93
						_736_v15 = _nw93
						var _737_v16 _dafny.Sequence
						_ = _737_v16
						_737_v16 = _dafny.SeqOf(_736_v15)
						_736_v15 = (_737_v16).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_719_v0), _dafny.IntOfUint32((_737_v16).Cardinality()))).Uint32()).(_dafny.Array)
						var _738_v17 _dafny.Array
						_ = _738_v17
						var _nwElement0_20 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("swkonbn"), _720_v1)
						_ = _nwElement0_20
						var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(3))
						_ = _nw94
						(_nw94).ArraySet1(_nwElement0_20, 0)
						(_nw94).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ldxmpoy"), 1)
						(_nw94).ArraySet1(_720_v1, 2)
						_738_v17 = _nw94
						var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_738_v17), 0))
						_ = _index86
						(_738_v17).ArraySet1(_720_v1, (_index86).Int())
						var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_738_v17), 0))
						_ = _index87
						(_738_v17).ArraySet1(_720_v1, (_index87).Int())
						var _739_v18 _dafny.Map
						_ = _739_v18
						_739_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_723_v4).Cardinality()), p0)
						_739_v18 = (_739_v18).Update(_719_v0, !(p0))
					} else {
						r0 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_723_v4, _723_v4), p0)
						var _740_v19 *C0
						_ = _740_v19
						var _nw95 *C0 = New_C0_()
						_ = _nw95
						_nw95.Ctor__()
						_740_v19 = _nw95
						var _741_v20 _dafny.Array
						_ = _741_v20
						var _len0_12 _dafny.Int = _dafny.IntOfInt64(26)
						_ = _len0_12
						var _nw96 _dafny.Array
						_ = _nw96
						if _len0_12.Cmp(_dafny.Zero) == 0 {
							_nw96 = _dafny.NewArray(_len0_12)
						} else {
							var _init12 func(_dafny.Int) _dafny.Int = func(_742_i2 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeDivisionInt(_742_i2, _dafny.IntOfInt64(144))
							}
							_ = _init12
							var _element0_12 = _init12(_dafny.Zero)
							_ = _element0_12
							_nw96 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
							(_nw96).ArraySet1(_element0_12, 0)
							var _nativeLen0_12 = (_len0_12).Int()
							_ = _nativeLen0_12
							for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
								(_nw96).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
							}
						}
						_741_v20 = _nw96
						var _743_v21 _dafny.Array
						_ = _743_v21
						var _nw97 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
						_ = _nw97
						_743_v21 = _nw97
						var _744_v22 _dafny.Map
						_ = _744_v22
						_744_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))
						_ = _index88
						(_743_v21).ArraySet1((func() bool {
							if (_744_v22).Contains(p0) {
								return (_744_v22).Get(p0).(bool)
							}
							return false
						})(), (_index88).Int())
						var _745_v23 _dafny.Array
						_ = _745_v23
						var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
						_ = _nw98
						_745_v23 = _nw98
						var _746_v24 _dafny.MultiSet
						_ = _746_v24
						_746_v24 = _dafny.MultiSetOf(_720_v1, _720_v1)
						var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_745_v23), 0))
						_ = _index89
						(_745_v23).ArraySet1(_746_v24, (_index89).Int())
						var _747_v25 _dafny.Map
						_ = _747_v25
						_747_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _741_v20)
						var _748_v26 _dafny.Map
						_ = _748_v26
						_748_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(348)).Plus(_719_v0), (func() _dafny.Array {
							if (_747_v25).Contains(p0) {
								return (_747_v25).Get(p0).(_dafny.Array)
							}
							return _741_v20
						})())
						var _749_v27 _dafny.Map
						_ = _749_v27
						_749_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _719_v0)
						var _750_v28 _dafny.Map
						_ = _750_v28
						_750_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v27, (_724_v5).IsSubsetOf(_724_v5))
						var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))
						_ = _index90
						var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_745_v23), 0))
						_ = _index91
						var _rhs74 _dafny.Array = (func() _dafny.Array {
							if (_748_v26).Contains(Companion_Default___.SafeModuloInt(_719_v0, _dafny.IntOfUint32((_721_v2).Cardinality()))) {
								return (_748_v26).Get(Companion_Default___.SafeModuloInt(_719_v0, _dafny.IntOfUint32((_721_v2).Cardinality()))).(_dafny.Array)
							}
							return _741_v20
						})()
						_ = _rhs74
						var _rhs75 bool = (func() bool {
							if (_750_v28).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _719_v0)) {
								return (_750_v28).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _719_v0)).(bool)
							}
							return p0
						})()
						_ = _rhs75
						var _rhs76 _dafny.MultiSet = _746_v24
						_ = _rhs76
						var _lhs54 _dafny.Array = _743_v21
						_ = _lhs54
						var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))
						_ = _lhs55
						var _lhs56 _dafny.Array = _745_v23
						_ = _lhs56
						var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_745_v23), 0))
						_ = _lhs57
						_741_v20 = _rhs74
						(_lhs54).ArraySet1(_rhs75, (_lhs55).Int())
						(_lhs56).ArraySet1(_rhs76, (_lhs57).Int())
						var _751_v29 _dafny.MultiSet
						_ = _751_v29
						_751_v29 = _dafny.MultiSetOf(p0, p0)
						var _752_v30 _dafny.Sequence
						_ = _752_v30
						_752_v30 = _dafny.SeqOf(_751_v29)
						var _753_v31 D4
						_ = _753_v31
						_753_v31 = Companion_D4_.Create_DC14_(_732_v11, _743_v21, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(736))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg57 _dafny.Int) interface{} {
								return coer57(arg57)
							}
						}((func(_754_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_755_i3 _dafny.Int) _dafny.CodePoint {
								return _754_v11
							}
						})(_732_v11))), (Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(736))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg58 _dafny.Int) interface{} {
								return coer58(arg58)
							}
						}((func(_756_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_757_i3 _dafny.Int) _dafny.CodePoint {
								return _756_v11
							}
						})(_732_v11)))).Cardinality()))).Uint32(), _732_v11), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(736))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}((func(_758_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_759_i3 _dafny.Int) _dafny.CodePoint {
								return _758_v11
							}
						})(_732_v11))), (Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(736))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg60 _dafny.Int) interface{} {
								return coer60(arg60)
							}
						}((func(_760_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_761_i3 _dafny.Int) _dafny.CodePoint {
								return _760_v11
							}
						})(_732_v11)))).Cardinality()))).Uint32(), _732_v11)).Cardinality()))).Uint32(), _732_v11))
						var _762_v32 _dafny.Map
						_ = _762_v32
						_762_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_753_v31, (_dafny.Zero).Minus(_719_v0))
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))
						_ = _index92
						(_743_v21).ArraySet1(((func() _dafny.MultiSet {
							if !(Companion_Default___.Fm15(p0, globalState)) {
								return _dafny.MultiSetOf(p0, (_743_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))).Int()).(bool), (_743_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))).Int()).(bool))
							}
							return _751_v29
						})()).IsDisjointFrom((_752_v30).Select((Companion_Default___.SafeIndex((_762_v32).Cardinality(), _dafny.IntOfUint32((_752_v30).Cardinality()))).Uint32()).(_dafny.MultiSet)), (_index92).Int())
						var _763_v33 _dafny.Map
						_ = _763_v33
						_763_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_740_v19, (_743_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_743_v21), 0))).Int()).(bool))
						_763_v33 = (_763_v33).Merge(_763_v33)
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _764_v34 D1
		_ = _764_v34
		_764_v34 = Companion_D1_.Create_DC3_(_720_v1)
		var _765_v35 _dafny.Array
		_ = _765_v35
		var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw99
		_765_v35 = _nw99
		var _766_v36 D2
		_ = _766_v36
		_766_v36 = Companion_D2_.Create_DC8_(_764_v34, _765_v35, p0)
		if ((func() D2 {
			if p0 {
				return Companion_D2_.Create_DC8_(_764_v34, _765_v35, p0)
			}
			return _766_v36
		})()).Dtor_cf16() {
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
			_ = _index93
			(_765_v35).ArraySet1(_719_v0, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
			_ = _index94
			(_765_v35).ArraySet1(_719_v0, (_index94).Int())
			var _767_v37 _dafny.CodePoint
			_ = _767_v37
			_767_v37 = _dafny.CodePoint('q')
			var _768_v38 _dafny.Map
			_ = _768_v38
			_768_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_720_v1).Cardinality()), false)
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
			_ = _index95
			var _rhs77 _dafny.Int = _719_v0
			_ = _rhs77
			var _rhs78 _dafny.CodePoint = (func() _dafny.CodePoint {
				if p0 {
					return _767_v37
				}
				return _767_v37
			})()
			_ = _rhs78
			var _rhs79 bool = ((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)).Cmp((_768_v38).Cardinality()) <= 0
			_ = _rhs79
			var _rhs80 _dafny.Int = _719_v0
			_ = _rhs80
			var _lhs58 _dafny.Array = _765_v35
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
			_ = _lhs59
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			var _lhs61 *GlobalState = globalState
			_ = _lhs61
			(_lhs58).ArraySet1(_rhs77, (_lhs59).Int())
			_767_v37 = _rhs78
			_lhs60.F0 = _rhs79
			_lhs61.F11 = _rhs80
			var _769_v39 D2
			_ = _769_v39
			_769_v39 = Companion_D2_.Create_DC6_(Companion_Default___.SafeModuloInt(_719_v0, _719_v0))
			_769_v39 = _769_v39
			if (_723_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_721_v2).Cardinality()), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool) {
				var _rhs81 bool = !(_724_v5).Contains((_dafny.Zero).Minus(_719_v0))
				_ = _rhs81
				var _rhs82 _dafny.Sequence = Companion_Default___.Fm18(p0, globalState)
				_ = _rhs82
				var _rhs83 bool = Companion_Default___.Fm15(((_dafny.Zero).Minus((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int))).Cmp(_719_v0) > 0, globalState)
				_ = _rhs83
				var _rhs84 bool = p0
				_ = _rhs84
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				_lhs62.F10 = _rhs81
				_720_v1 = _rhs82
				_lhs63.F10 = _rhs83
				r0 = _rhs84
				(globalState).F16 = (_dafny.Zero).Minus((_719_v0).Minus((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)))
				var _770_v40 _dafny.Array
				_ = _770_v40
				var _nwElement0_21 _dafny.CodePoint = _767_v37
				_ = _nwElement0_21
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(15))
				_ = _nw100
				(_nw100).ArraySet1CodePoint(_nwElement0_21, 0)
				(_nw100).ArraySet1CodePoint(_dafny.CodePoint('e'), 1)
				(_nw100).ArraySet1CodePoint(_767_v37, 2)
				(_nw100).ArraySet1CodePoint(_767_v37, 3)
				(_nw100).ArraySet1CodePoint(_767_v37, 4)
				(_nw100).ArraySet1CodePoint(_767_v37, 5)
				(_nw100).ArraySet1CodePoint((func() _dafny.CodePoint {
					if true {
						return _767_v37
					}
					return _767_v37
				})(), 6)
				(_nw100).ArraySet1CodePoint((func() _dafny.CodePoint {
					if p0 {
						return _767_v37
					}
					return _767_v37
				})(), 7)
				(_nw100).ArraySet1CodePoint(_767_v37, 8)
				(_nw100).ArraySet1CodePoint((_720_v1).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_720_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
				(_nw100).ArraySet1CodePoint(_767_v37, 10)
				(_nw100).ArraySet1CodePoint(_767_v37, 11)
				(_nw100).ArraySet1CodePoint(_767_v37, 12)
				(_nw100).ArraySet1CodePoint(_dafny.CodePoint('g'), 13)
				(_nw100).ArraySet1CodePoint(_767_v37, 14)
				_770_v40 = _nw100
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_770_v40), 0))
				_ = _index96
				(_770_v40).ArraySet1CodePoint(_767_v37, (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_770_v40), 0))
				_ = _index97
				(_770_v40).ArraySet1CodePoint(_767_v37, (_index97).Int())
				var _771_v41 _dafny.Map
				_ = _771_v41
				_771_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
				r3 = ((_771_v41).Cardinality()).Times(_719_v0)
				var _772_v42 _dafny.Map
				_ = _772_v42
				_772_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v1, _dafny.Companion_Sequence_.IsPrefixOf(_723_v4, _723_v4))
				_772_v42 = (_772_v42).Update(_720_v1, _dafny.Companion_Sequence_.IsPrefixOf(_723_v4, _723_v4))
			} else {
				_767_v37 = _dafny.CodePoint('l')
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
				_ = _index98
				(_765_v35).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm18(p0, globalState)).Cardinality()), (_index98).Int())
				var _773_v43 _dafny.Map
				_ = _773_v43
				_773_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v1, (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int))
				var _774_v45 _dafny.Set
				_ = _774_v45
				_774_v45 = _dafny.SetOf(_720_v1)
				var _775_v46 D5
				_ = _775_v46
				_775_v46 = Companion_D5_.Create_DC17_(_773_v43)
				var _776_v47 _dafny.Array
				_ = _776_v47
				var _nwElement0_22 _dafny.Map = _773_v43
				_ = _nwElement0_22
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(20))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_22, 0)
				(_nw101).ArraySet1(_773_v43, 1)
				(_nw101).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v1, _719_v0), 2)
				(_nw101).ArraySet1(_773_v43, 3)
				(_nw101).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ucufv"), (Companion_Default___.SafeIndex((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ucufv")).Cardinality()))).Uint32(), _767_v37), _719_v0), 4)
				(_nw101).ArraySet1(_773_v43, 5)
				(_nw101).ArraySet1((func() _dafny.Map {
					if (_723_v4).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool) {
						return (func() _dafny.Map {
							var _coll38 = _dafny.NewMapBuilder()
							_ = _coll38
							for _iter39 := _dafny.Iterate((_774_v45).Elements()); ; {
								_compr_38, _ok39 := _iter39()
								if !_ok39 {
									break
								}
								var _777_v44 _dafny.Sequence
								_777_v44 = interface{}(_compr_38).(_dafny.Sequence)
								if (_774_v45).Contains(_777_v44) {
									_coll38.Add(_777_v44, (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int))
								}
							}
							return _coll38.ToMap()
						}()).Update(_720_v1, _dafny.IntOfUint32((_721_v2).Cardinality()))
					}
					return _773_v43
				})(), 6)
				(_nw101).ArraySet1(Companion_Default___.Fm19(p0, _719_v0, (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), true, globalState), 7)
				(_nw101).ArraySet1(_773_v43, 8)
				(_nw101).ArraySet1(_773_v43, 9)
				(_nw101).ArraySet1(_773_v43, 10)
				(_nw101).ArraySet1(_773_v43, 11)
				(_nw101).ArraySet1((_773_v43).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v1, _dafny.IntOfUint32((_723_v4).Cardinality()))), 12)
				(_nw101).ArraySet1(_773_v43, 13)
				(_nw101).ArraySet1(_773_v43, 14)
				(_nw101).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v1, _719_v0), 15)
				(_nw101).ArraySet1(_773_v43, 16)
				(_nw101).ArraySet1(_773_v43, 17)
				(_nw101).ArraySet1((_775_v46).Dtor_cf29(), 18)
				(_nw101).ArraySet1((func() _dafny.Map {
					if p0 {
						return _773_v43
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v1, _719_v0)
				})(), 19)
				_776_v47 = _nw101
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_776_v47), 0))
				_ = _index99
				(_776_v47).ArraySet1(_773_v43, (_index99).Int())
				var _778_v48 _dafny.MultiSet
				_ = _778_v48
				_778_v48 = _dafny.MultiSetOf(p0, false)
				var _779_v50 _dafny.Map
				_ = _779_v50
				_779_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, func() _dafny.Map {
					var _coll39 = _dafny.NewMapBuilder()
					_ = _coll39
					for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(733), _dafny.IntOfInt64(-844))); ; {
						_compr_39, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _780_v49 _dafny.Int
						_780_v49 = interface{}(_compr_39).(_dafny.Int)
						if ((_dafny.IntOfInt64(733)).Cmp(_780_v49) <= 0) && ((_780_v49).Cmp(_dafny.IntOfInt64(-844)) < 0) {
							_coll39.Add((_780_v49).Times((_721_v2).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_721_v2).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfInt64(859))
						}
					}
					return _coll39.ToMap()
				}())
				var _781_v51 D0
				_ = _781_v51
				_781_v51 = Companion_D0_.Create_DC1_((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), p0, _779_v50, _719_v0, _767_v37)
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_776_v47), 0))
				_ = _index100
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
				_ = _index101
				var _rhs85 _dafny.Map = Companion_Default___.Fm19(((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
					if (_778_v48).Contains(p0) {
						return (_778_v48).Multiplicity(p0)
					}
					return (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)
				})()) > 0, (_719_v0).Times(_719_v0), _dafny.IntOfInt64(829), (_781_v51).Dtor_cf1(), globalState)
				_ = _rhs85
				var _rhs86 _dafny.Int = Companion_Default___.SafeDivisionInt((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(557), (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)))
				_ = _rhs86
				var _rhs87 _dafny.Int = _dafny.IntOfInt64(-230)
				_ = _rhs87
				var _lhs64 _dafny.Array = _776_v47
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_776_v47), 0))
				_ = _lhs65
				var _lhs66 _dafny.Array = _765_v35
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
				_ = _lhs67
				(_lhs64).ArraySet1(_rhs85, (_lhs65).Int())
				(_lhs66).ArraySet1(_rhs86, (_lhs67).Int())
				r3 = _rhs87
				var _782_v52 _dafny.Array
				_ = _782_v52
				var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw102
				_782_v52 = _nw102
				var _783_v53 _dafny.Array
				_ = _783_v53
				var _nwElement0_23 _dafny.Int = (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)
				_ = _nwElement0_23
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(9))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_23, 0)
				(_nw103).ArraySet1((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), 1)
				(_nw103).ArraySet1(_719_v0, 2)
				(_nw103).ArraySet1((Companion_Default___.Fm20(p0, _767_v37, _719_v0, p0, globalState)).Cardinality(), 3)
				(_nw103).ArraySet1(_719_v0, 4)
				(_nw103).ArraySet1(_719_v0, 5)
				(_nw103).ArraySet1(_719_v0, 6)
				(_nw103).ArraySet1((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), 7)
				(_nw103).ArraySet1((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), 8)
				_783_v53 = _nw103
				var _784_v54 _dafny.Map
				_ = _784_v54
				_784_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v53, p0)
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_782_v52), 0))
				_ = _index102
				(_782_v52).ArraySet1(_784_v54, (_index102).Int())
				var _785_v55 *C0
				_ = _785_v55
				var _nw104 *C0 = New_C0_()
				_ = _nw104
				_nw104.Ctor__()
				_785_v55 = _nw104
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_782_v52), 0))
				_ = _index103
				var _rhs88 _dafny.Map = _784_v54
				_ = _rhs88
				var _rhs89 *C0 = _785_v55
				_ = _rhs89
				var _lhs68 _dafny.Array = _782_v52
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_782_v52), 0))
				_ = _lhs69
				(_lhs68).ArraySet1(_rhs88, (_lhs69).Int())
				_785_v55 = _rhs89
				var _786_v56 _dafny.Map
				_ = _786_v56
				_786_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				(globalState).F0 = (_723_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_786_v56).Cardinality(), (_dafny.Zero).Minus(_719_v0))), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool)
			}
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
			_ = _index104
			var _rhs90 _dafny.Int = _719_v0
			_ = _rhs90
			var _rhs91 bool = p0
			_ = _rhs91
			var _lhs70 _dafny.Array = _765_v35
			_ = _lhs70
			var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_765_v35), 0))
			_ = _lhs71
			var _lhs72 *GlobalState = globalState
			_ = _lhs72
			(_lhs70).ArraySet1(_rhs90, (_lhs71).Int())
			_lhs72.F0 = _rhs91
		} else {
			(globalState).F10 = Companion_Default___.Fm15(p0, globalState)
			(globalState).F0 = p0
			var _787_v57 _dafny.Array
			_ = _787_v57
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_13
			var _nw105 _dafny.Array
			_ = _nw105
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw105 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) bool = (func(_788_p0 bool) func(_dafny.Int) bool {
					return func(_789_i4 _dafny.Int) bool {
						return _788_p0
					}
				})(p0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw105 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw105).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw105).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_787_v57 = _nw105
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_787_v57), 0))
			_ = _index105
			(_787_v57).ArraySet1(false, (_index105).Int())
			var _790_v58 _dafny.Set
			_ = _790_v58
			_790_v58 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, p0))
			var _791_v59 _dafny.Map
			_ = _791_v59
			_791_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(205), Companion_Default___.Fm15(false, globalState))
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_787_v57), 0))
			_ = _index106
			var _rhs92 bool = (_dafny.SetOf(_791_v59)).IsProperSubsetOf(_790_v58)
			_ = _rhs92
			var _rhs93 bool = !((_724_v5).Contains(_719_v0)) || (Companion_Default___.Fm15(p0, globalState))
			_ = _rhs93
			var _lhs73 _dafny.Array = _787_v57
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_787_v57), 0))
			_ = _lhs74
			r0 = _rhs92
			(_lhs73).ArraySet1(_rhs93, (_lhs74).Int())
			var _792_v60 *C0
			_ = _792_v60
			var _nw106 *C0 = New_C0_()
			_ = _nw106
			_nw106.Ctor__()
			_792_v60 = _nw106
			var _793_v61 D2
			_ = _793_v61
			_793_v61 = Companion_D2_.Create_DC9_(Companion_D2_.Create_DC6_(_719_v0))
			var _794_v62 D2
			_ = _794_v62
			_794_v62 = Companion_D2_.Create_DC9_(_793_v61)
			var _795_v63 _dafny.Set
			_ = _795_v63
			_795_v63 = _dafny.SetOf(_794_v62)
			var _796_v64 _dafny.MultiSet
			_ = _796_v64
			_796_v64 = _dafny.MultiSetOf(_787_v57)
			var _797_v65 _dafny.MultiSet
			_ = _797_v65
			_797_v65 = _dafny.MultiSetOf((_796_v64).Cardinality(), _dafny.IntOfInt64(715))
			var _798_v66 _dafny.Map
			_ = _798_v66
			_798_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _797_v65)
			var _799_v67 D4
			_ = _799_v67
			_799_v67 = Companion_D4_.Create_DC12_(_798_v66)
			var _800_v68 D4
			_ = _800_v68
			_800_v68 = Companion_D4_.Create_DC16_(_799_v67)
			var _801_v69 D4
			_ = _801_v69
			_801_v69 = Companion_D4_.Create_DC16_((_800_v68).Dtor_cf28())
			var _802_v70 D4
			_ = _802_v70
			_802_v70 = Companion_D4_.Create_DC16_(_799_v67)
			var _803_v71 _dafny.Sequence
			_ = _803_v71
			_803_v71 = _dafny.SeqOf(_800_v68, _800_v68, _800_v68)
			var _804_v72 _dafny.Map
			_ = _804_v72
			_804_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_794_v62)).Intersection(_795_v63), _803_v71)
			_804_v72 = (_804_v72).Update(_795_v63, _803_v71)
		}
		var _805_v73 _dafny.CodePoint
		_ = _805_v73
		_805_v73 = _dafny.CodePoint('o')
		var _806_v74 _dafny.Array
		_ = _806_v74
		var _nw107 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw107
		_806_v74 = _nw107
		var _807_v75 D4
		_ = _807_v75
		_807_v75 = Companion_D4_.Create_DC14_(_805_v73, _806_v74, _720_v1)
		var _source10 D4 = _807_v75
		_ = _source10
		if _source10.Is_DC13() {
			var _808___mcc_h0 D0 = _source10.Get_().(D4_DC13).Cf22
			_ = _808___mcc_h0
			var _809_cf22 D0 = _808___mcc_h0
			_ = _809_cf22
			if p0 {
				(globalState).F0 = (_723_v4).Select((Companion_Default___.SafeIndex((_729_v9).Cardinality(), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool)
				var _810_v76 _dafny.Array
				_ = _810_v76
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_14
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.Sequence = (func(_811_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_812_i5 _dafny.Int) _dafny.Sequence {
							return _811_v2
						}
					})(_721_v2)
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw108 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw108).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw108).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_810_v76 = _nw108
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_810_v76), 0))
				_ = _index107
				(_810_v76).ArraySet1(_721_v2, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_765_v35), 0))
				_ = _index108
				(_765_v35).ArraySet1(_719_v0, (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_810_v76), 0))
				_ = _index109
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_765_v35), 0))
				_ = _index110
				var _rhs94 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}((func(_813_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_814_i6 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_813_v0)
					}
				})(_719_v0))), _dafny.Companion_Sequence_.Concatenate(_721_v2, _721_v2)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_719_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_815_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_816_i6 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_815_v0)
					}
				})(_719_v0))), _dafny.Companion_Sequence_.Concatenate(_721_v2, _721_v2))).Cardinality()))).Uint32(), _719_v0)
				_ = _rhs94
				var _rhs95 _dafny.Int = (_719_v0).Plus((_719_v0).Minus(_719_v0))
				_ = _rhs95
				var _lhs75 _dafny.Array = _810_v76
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_810_v76), 0))
				_ = _lhs76
				var _lhs77 _dafny.Array = _765_v35
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_765_v35), 0))
				_ = _lhs78
				(_lhs75).ArraySet1(_rhs94, (_lhs76).Int())
				(_lhs77).ArraySet1(_rhs95, (_lhs78).Int())
				var _817_v77 _dafny.Array
				_ = _817_v77
				var _818_v78 _dafny.Int
				_ = _818_v78
				var _out46 _dafny.Array
				_ = _out46
				var _out47 _dafny.Int
				_ = _out47
				_out46, _out47 = (_this).M4(globalState)
				_817_v77 = _out46
				_818_v78 = _out47
				(globalState).F5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, p0), _723_v4), _723_v4)
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_806_v74), 0))
				_ = _index111
				(_806_v74).ArraySet1(!(!(((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)).Cmp(_719_v0) < 0)), (_index111).Int())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_806_v74), 0))
				_ = _index112
				(_806_v74).ArraySet1(Companion_Default___.Fm15(true, globalState), (_index112).Int())
			} else {
				var _819_v79 _dafny.Map
				_ = _819_v79
				_819_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _806_v74)
				(globalState).F6 = (_819_v79).Cardinality()
				_765_v35 = _765_v35
				_805_v73 = _805_v73
				var _820_v80 *C0
				_ = _820_v80
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__()
				_820_v80 = _nw109
				var _rhs96 _dafny.Array = _765_v35
				_ = _rhs96
				var _rhs97 _dafny.Int = _719_v0
				_ = _rhs97
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_765_v35 = _rhs96
				_lhs79.F11 = _rhs97
			}
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))
			_ = _index113
			(_765_v35).ArraySet1(_719_v0, (_index113).Int())
			var _821_v81 _dafny.MultiSet
			_ = _821_v81
			_821_v81 = _dafny.MultiSetOf(_719_v0)
			var _822_v82 _dafny.Map
			_ = _822_v82
			_822_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("lgi"), p0)
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))
			_ = _index114
			(_765_v35).ArraySet1((_this).Fm0(_821_v81, (_dafny.Zero).Minus((_822_v82).Cardinality()), _dafny.SeqOf(func() _dafny.Set {
				var _coll40 = _dafny.NewBuilder()
				_ = _coll40
				for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(623), _dafny.IntOfInt64(366))); ; {
					_compr_40, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _823_v83 _dafny.Int
					_823_v83 = interface{}(_compr_40).(_dafny.Int)
					if ((_dafny.IntOfInt64(623)).Cmp(_823_v83) <= 0) && ((_823_v83).Cmp(_dafny.IntOfInt64(366)) < 0) {
						_coll40.Add(Companion_Default___.SafeDivisionInt(_823_v83, _719_v0))
					}
				}
				return _coll40.ToSet()
			}()), p0, globalState), (_index114).Int())
			var _824_v84 _dafny.Map
			_ = _824_v84
			_824_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
			var _825_v85 _dafny.Sequence
			_ = _825_v85
			_825_v85 = _dafny.SeqOf(_824_v84)
			(globalState).F6 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_825_v85).Cardinality()))
			if p0 {
				(globalState).F10 = (func() bool {
					if !_dafny.Companion_Sequence_.Contains(_720_v1, _805_v73) {
						return p0
					}
					return p0
				})()
				(globalState).F0 = (_dafny.IntOfInt64(-252)).Cmp(_719_v0) < 0
				var _826_v86 T0
				_ = _826_v86
				var _nw110 *C1 = New_C1_()
				_ = _nw110
				_nw110.Ctor__(_719_v0, (_this).F22())
				_826_v86 = _nw110
				var _827_v87 _dafny.Sequence
				_ = _827_v87
				_827_v87 = _dafny.SeqOf(_826_v86)
				var _828_v88 _dafny.Set
				_ = _828_v88
				_828_v88 = _dafny.SetOf(_827_v87, _dafny.SeqOf(_826_v86))
				var _829_v89 _dafny.Map
				_ = _829_v89
				_829_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_723_v4).Select((Companion_Default___.SafeIndex((_this).Fm0(_dafny.MultiSetOf((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), _719_v0, _719_v0), (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), _725_v6, p0, globalState), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool), _828_v88)
				var _rhs98 _dafny.Int = (((_828_v88).Difference((func() _dafny.Set {
					if (_829_v89).Contains(p0) {
						return (_829_v89).Get(p0).(_dafny.Set)
					}
					return _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_826_v86), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_826_v86)).Cardinality()))).Uint32(), _826_v86))
				})())).Cardinality()).Plus((_719_v0).Times((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)))
				_ = _rhs98
				var _rhs99 _dafny.Int = Companion_Default___.SafeDivisionInt(_719_v0, _719_v0)
				_ = _rhs99
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				_lhs80.F11 = _rhs98
				r3 = _rhs99
				var _830_v90 _dafny.MultiSet
				_ = _830_v90
				_830_v90 = _dafny.MultiSetOf(p0)
				var _831_v91 _dafny.Map
				_ = _831_v91
				_831_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), p0)
				var _832_v92 _dafny.Map
				_ = _832_v92
				_832_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if p0 {
						return (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_830_v90).Contains((func() bool {
							if (_831_v91).Contains(_719_v0) {
								return (_831_v91).Get(_719_v0).(bool)
							}
							return p0
						})()) {
							return (_830_v90).Multiplicity((func() bool {
								if (_831_v91).Contains(_719_v0) {
									return (_831_v91).Get(_719_v0).(bool)
								}
								return p0
							})())
						}
						return _719_v0
					})()
				})(), Companion_Default___.SafeDivisionInt((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)))
				var _rhs100 _dafny.CodePoint = _805_v73
				_ = _rhs100
				var _rhs101 _dafny.Int = (func() _dafny.Int {
					if (_832_v92).Contains(_dafny.IntOfUint32((_720_v1).Cardinality())) {
						return (_832_v92).Get(_dafny.IntOfUint32((_720_v1).Cardinality())).(_dafny.Int)
					}
					return (_dafny.MultiSetOf(_720_v1)).Cardinality()
				})()
				_ = _rhs101
				var _rhs102 bool = p0
				_ = _rhs102
				var _rhs103 bool = p0
				_ = _rhs103
				var _rhs104 T0 = _826_v86
				_ = _rhs104
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				var _lhs83 *GlobalState = globalState
				_ = _lhs83
				_805_v73 = _rhs100
				_lhs81.F6 = _rhs101
				_lhs82.F0 = _rhs102
				_lhs83.F10 = _rhs103
				_826_v86 = _rhs104
				var _833_v93 _dafny.MultiSet
				_ = _833_v93
				_833_v93 = _dafny.MultiSetOf((_766_v36).Dtor_cf15())
				var _834_v94 *C1
				_ = _834_v94
				var _nw111 *C1 = New_C1_()
				_ = _nw111
				_nw111.Ctor__((_833_v93).Cardinality(), (_826_v86).F22())
				_834_v94 = _nw111
			} else {
				var _835_v95 D1
				_ = _835_v95
				_835_v95 = Companion_D1_.Create_DC2_(_720_v1)
				var _836_v96 _dafny.Set
				_ = _836_v96
				_836_v96 = _dafny.SetOf(_835_v95, _835_v95)
				(globalState).F10 = !((_dafny.SetOf(_835_v95, _835_v95)).IsDisjointFrom(_836_v96))
				var _837_v98 _dafny.Map
				_ = _837_v98
				_837_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int), p0)
				var _838_v100 _dafny.Map
				_ = _838_v100
				_838_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter42 := _dafny.Iterate((_837_v98).Keys().Elements()); ; {
						_compr_41, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _839_v97 _dafny.Int
						_839_v97 = interface{}(_compr_41).(_dafny.Int)
						if (_837_v98).Contains(_839_v97) {
							_coll41.Add(Companion_Default___.SafeDivisionInt(_839_v97, (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)), false)
						}
					}
					return _coll41.ToMap()
				}()).Merge((func() _dafny.Map {
					var _coll42 = _dafny.NewMapBuilder()
					_ = _coll42
					for _iter43 := _dafny.Iterate((_837_v98).Keys().Elements()); ; {
						_compr_42, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _840_v99 _dafny.Int
						_840_v99 = interface{}(_compr_42).(_dafny.Int)
						if (_837_v98).Contains(_840_v99) {
							_coll42.Add((_840_v99).Minus(_719_v0), p0)
						}
					}
					return _coll42.ToMap()
				}()).Update(_719_v0, true)))
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_765_v35), 0))
				_ = _index115
				(_765_v35).ArraySet1(((func() _dafny.Map {
					if (_838_v100).Contains(!((true) && (p0))) {
						return (_838_v100).Get(!((true) && (p0))).(_dafny.Map)
					}
					return (_837_v98).Merge(_837_v98)
				})()).Cardinality(), (_index115).Int())
				_837_v98 = ((_837_v98).Merge(_837_v98)).Merge(_837_v98)
				var _841_v101 _dafny.Array
				_ = _841_v101
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_15
				var _nw112 _dafny.Array
				_ = _nw112
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw112 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Sequence = (func(_842_v1 _dafny.Sequence, _843_v95 D1) func(_dafny.Int) _dafny.Sequence {
						return func(_844_i7 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_842_v1, (_843_v95).Dtor_cf5())
						}
					})(_720_v1, _835_v95)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw112 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw112).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw112).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_841_v101 = _nw112
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_841_v101), 0))
				_ = _index116
				(_841_v101).ArraySet1(_720_v1, (_index116).Int())
				var _845_v102 D12
				_ = _845_v102
				_845_v102 = Companion_D12_.Create_DC42_(_719_v0)
				var _846_v103 _dafny.Map
				_ = _846_v103
				_846_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('f'))
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_841_v101), 0))
				_ = _index117
				var _rhs105 bool = p0
				_ = _rhs105
				var _rhs106 _dafny.Int = (_dafny.Zero).Minus((_845_v102).Dtor_cf69())
				_ = _rhs106
				var _rhs107 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_720_v1, _720_v1)
				_ = _rhs107
				var _rhs108 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_720_v1, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_719_v0), _dafny.IntOfUint32((_720_v1).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_846_v103).Contains(p0) {
						return (_846_v103).Get(p0).(_dafny.CodePoint)
					}
					return _805_v73
				})())
				_ = _rhs108
				var _rhs109 bool = p0
				_ = _rhs109
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 _dafny.Array = _841_v101
				_ = _lhs86
				var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_841_v101), 0))
				_ = _lhs87
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				_lhs84.F0 = _rhs105
				_lhs85.F11 = _rhs106
				_720_v1 = _rhs107
				(_lhs86).ArraySet1(_rhs108, (_lhs87).Int())
				_lhs88.F10 = _rhs109
				var _847_v104 _dafny.Array
				_ = _847_v104
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw113
				_847_v104 = _nw113
				var _848_v105 _dafny.Array
				_ = _848_v105
				var _nwElement0_24 _dafny.Array = _847_v104
				_ = _nwElement0_24
				var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(21))
				_ = _nw114
				(_nw114).ArraySet1(_nwElement0_24, 0)
				(_nw114).ArraySet1(_847_v104, 1)
				(_nw114).ArraySet1(_847_v104, 2)
				(_nw114).ArraySet1(_847_v104, 3)
				(_nw114).ArraySet1(_847_v104, 4)
				(_nw114).ArraySet1(_847_v104, 5)
				(_nw114).ArraySet1(_847_v104, 6)
				(_nw114).ArraySet1(_847_v104, 7)
				(_nw114).ArraySet1(_847_v104, 8)
				(_nw114).ArraySet1(_847_v104, 9)
				(_nw114).ArraySet1(_847_v104, 10)
				(_nw114).ArraySet1(_847_v104, 11)
				(_nw114).ArraySet1(_847_v104, 12)
				(_nw114).ArraySet1(_847_v104, 13)
				(_nw114).ArraySet1(_847_v104, 14)
				(_nw114).ArraySet1(_847_v104, 15)
				(_nw114).ArraySet1(_847_v104, 16)
				(_nw114).ArraySet1(_847_v104, 17)
				(_nw114).ArraySet1(_847_v104, 18)
				(_nw114).ArraySet1(_847_v104, 19)
				(_nw114).ArraySet1(_847_v104, 20)
				_848_v105 = _nw114
				_848_v105 = _848_v105
			}
		} else if _source10.Is_DC14() {
			var _849___mcc_h1 _dafny.CodePoint = _source10.Get_().(D4_DC14).Cf23
			_ = _849___mcc_h1
			var _850___mcc_h2 _dafny.Array = _source10.Get_().(D4_DC14).Cf24
			_ = _850___mcc_h2
			var _851___mcc_h3 _dafny.Sequence = _source10.Get_().(D4_DC14).Cf25
			_ = _851___mcc_h3
			var _852_cf25 _dafny.Sequence = _851___mcc_h3
			_ = _852_cf25
			var _853_cf24 _dafny.Array = _850___mcc_h2
			_ = _853_cf24
			var _854_cf23 _dafny.CodePoint = _849___mcc_h1
			_ = _854_cf23
			var _855_v106 _dafny.Array
			_ = _855_v106
			var _856_v107 _dafny.Int
			_ = _856_v107
			var _out48 _dafny.Array
			_ = _out48
			var _out49 _dafny.Int
			_ = _out49
			_out48, _out49 = (_this).M4(globalState)
			_855_v106 = _out48
			_856_v107 = _out49
			var _857_v109 _dafny.Map
			_ = _857_v109
			_857_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter44 := _dafny.Iterate((_720_v1).Elements()); ; {
					_compr_43, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _858_v108 _dafny.CodePoint
					_858_v108 = interface{}(_compr_43).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_720_v1, _858_v108) {
						_coll43.Add(_858_v108, (_724_v5).Cardinality())
					}
				}
				return _coll43.ToMap()
			}(), _856_v107)
			var _859_v110 _dafny.Map
			_ = _859_v110
			_859_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_cf23, (_724_v5).Cardinality())
			(globalState).F11 = (_dafny.Zero).Minus((((_857_v109).Merge((_857_v109).Update(_859_v110, _719_v0))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_859_v110, _719_v0))).Cardinality())
			if !((_719_v0).Cmp(Companion_Default___.SafeDivisionInt(_719_v0, _856_v107)) <= 0) {
				(globalState).F0 = p0
				_720_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_852_cf25, _720_v1), _852_cf25)
				var _860_v111 _dafny.Map
				_ = _860_v111
				_860_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(960), p0)
				var _861_v112 _dafny.Map
				_ = _861_v112
				_861_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _719_v0)
				var _862_v113 D9
				_ = _862_v113
				_862_v113 = Companion_D9_.Create_DC30_(_861_v112)
				var _863_v114 D9
				_ = _863_v114
				_863_v114 = Companion_D9_.Create_DC32_(_862_v113)
				var _864_v115 D9
				_ = _864_v115
				_864_v115 = Companion_D9_.Create_DC32_((func() D9 {
					if (func() bool {
						if (_860_v111).Contains(_719_v0) {
							return (_860_v111).Get(_719_v0).(bool)
						}
						return p0
					})() {
						return _862_v113
					}
					return _863_v114
				})())
				var _pat_let_tv12 = _862_v113
				_ = _pat_let_tv12
				_864_v115 = func(_pat_let10_0 D9) D9 {
					return func(_865_dt__update__tmp_h0 D9) D9 {
						return func(_pat_let11_0 D9) D9 {
							return func(_866_dt__update_hcf53_h0 D9) D9 {
								return Companion_D9_.Create_DC32_(_866_dt__update_hcf53_h0)
							}(_pat_let11_0)
						}(_pat_let_tv12)
					}(_pat_let10_0)
				}(_864_v115)
				r0 = false
				var _867_v116 _dafny.Sequence
				_ = _867_v116
				_867_v116 = _dafny.SeqOf(Companion_Default___.Fm29(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_868_v73 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_869_i8 _dafny.Int) _dafny.CodePoint {
						return _868_v73
					}
				})(_805_v73)))).Cardinality()), p0, _719_v0, globalState))
				var _870_v117 D11
				_ = _870_v117
				_870_v117 = Companion_D11_.Create_DC36_(_867_v116)
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_806_v74), 0))
				_ = _index118
				(_806_v74).ArraySet1(p0, (_index118).Int())
				var _871_v118 _dafny.MultiSet
				_ = _871_v118
				_871_v118 = _dafny.MultiSetOf(_856_v107, _856_v107, _856_v107)
				var _872_v119 _dafny.Map
				_ = _872_v119
				_872_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _871_v118)
				var _873_v120 _dafny.Map
				_ = _873_v120
				_873_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _852_cf25)
				var _874_v121 _dafny.Map
				_ = _874_v121
				_874_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-940))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}(func(_875_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Cardinality()), (_873_v120).Cardinality())
				var _876_v122 _dafny.Map
				_ = _876_v122
				_876_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _874_v121)
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_806_v74), 0))
				_ = _index119
				var _rhs110 _dafny.Int = (_this).Fm0((func() _dafny.MultiSet {
					if (_872_v119).Contains(p0) {
						return (_872_v119).Get(p0).(_dafny.MultiSet)
					}
					return _871_v118
				})(), (_dafny.MultiSetOf(!(p0), p0, p0)).Cardinality(), _dafny.SeqOf(_dafny.SetOf(_856_v107), _dafny.SetOf(_719_v0, _856_v107)), (_723_v4).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool), globalState)
				_ = _rhs110
				var _rhs111 D11 = Companion_Default___.Fm43(_dafny.CodePoint('v'), Companion_D0_.Create_DC1_(_719_v0, p0, _876_v122, _719_v0, _854_cf23), p0, (_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs(_856_v107)), globalState)
				_ = _rhs111
				var _rhs112 bool = p0
				_ = _rhs112
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 _dafny.Array = _806_v74
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_806_v74), 0))
				_ = _lhs91
				_lhs89.F6 = _rhs110
				_870_v117 = _rhs111
				(_lhs90).ArraySet1(_rhs112, (_lhs91).Int())
			} else {
				(globalState).F0 = (func() bool {
					if _dafny.Companion_Sequence_.IsProperPrefixOf(_723_v4, _dafny.SeqOf(p0)) {
						return p0
					}
					return p0
				})()
				var _877_v123 D9
				_ = _877_v123
				_877_v123 = Companion_D9_.Create_DC31_(_719_v0, _855_v106)
				var _878_v124 _dafny.Map
				_ = _878_v124
				_878_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _719_v0)
				var _879_v125 _dafny.Map
				_ = _879_v125
				_879_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _878_v124)
				var _880_v126 _dafny.Array
				_ = _880_v126
				var _nwElement0_25 _dafny.Sequence = _dafny.SeqOf((_877_v123).Dtor_cf52())
				_ = _nwElement0_25
				var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(8))
				_ = _nw115
				(_nw115).ArraySet1(_nwElement0_25, 0)
				(_nw115).ArraySet1((_this).F22(), 1)
				(_nw115).ArraySet1((_this).F22(), 2)
				(_nw115).ArraySet1(_dafny.SeqOf(_853_cf24), 3)
				(_nw115).ArraySet1((_this).F22(), 4)
				(_nw115).ArraySet1((_this).F22(), 5)
				(_nw115).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()), 6)
				(_nw115).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F22(), (Companion_Default___.SafeIndex((_879_v125).Cardinality(), _dafny.IntOfUint32(((_this).F22()).Cardinality()))).Uint32(), _853_cf24), (_this).F22()), 7)
				_880_v126 = _nw115
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_880_v126), 0))
				_ = _index120
				(_880_v126).ArraySet1(_dafny.SeqOf(_855_v106), (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_880_v126), 0))
				_ = _index121
				(_880_v126).ArraySet1((_this).F22(), (_index121).Int())
				var _881_v127 *C1
				_ = _881_v127
				var _nw116 *C1 = New_C1_()
				_ = _nw116
				_nw116.Ctor__((_dafny.Zero).Minus(_719_v0), _dafny.Companion_Sequence_.Concatenate((_this).F22(), (_880_v126).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_880_v126), 0))).Int()).(_dafny.Sequence)))
				_881_v127 = _nw116
				var _882_v128 _dafny.Array
				_ = _882_v128
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
				_ = _nw117
				_882_v128 = _nw117
				_882_v128 = (func() _dafny.Array {
					if p0 {
						return _882_v128
					}
					return _882_v128
				})()
				(globalState).F8 = Companion_Default___.SafeModuloInt(_719_v0, _719_v0)
			}
			var _883_v129 _dafny.Array
			_ = _883_v129
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw118
			_883_v129 = _nw118
			var _rhs113 _dafny.Array = _883_v129
			_ = _rhs113
			var _rhs114 _dafny.Int = _856_v107
			_ = _rhs114
			var _lhs92 *GlobalState = globalState
			_ = _lhs92
			_883_v129 = _rhs113
			_lhs92.F6 = _rhs114
		} else if _source10.Is_DC15() {
			var _884___mcc_h4 _dafny.Int = _source10.Get_().(D4_DC15).Cf26
			_ = _884___mcc_h4
			var _885___mcc_h5 _dafny.Int = _source10.Get_().(D4_DC15).Cf27
			_ = _885___mcc_h5
			var _886_cf27 _dafny.Int = _885___mcc_h5
			_ = _886_cf27
			var _887_cf26 _dafny.Int = _884___mcc_h4
			_ = _887_cf26
			r1 = _806_v74
			(globalState).F16 = _dafny.IntOfUint32((_720_v1).Cardinality())
			(globalState).F0 = (_719_v0).Cmp(_719_v0) < 0
			_805_v73 = (Companion_D10_.Create_DC34_(_dafny.CodePoint('b'), p0, p0, p0, p0)).Dtor_cf55()
		} else if _source10.Is_DC12() {
			var _888___mcc_h6 _dafny.Map = _source10.Get_().(D4_DC12).Cf21
			_ = _888___mcc_h6
			var _889_cf21 _dafny.Map = _888___mcc_h6
			_ = _889_cf21
			(globalState).F5 = _723_v4
			var _890_v130 _dafny.MultiSet
			_ = _890_v130
			_890_v130 = _dafny.MultiSetOf(true, !(p0), p0)
			_890_v130 = _dafny.MultiSetOf((_723_v4).Select((Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool), p0)
			var _891_v131 D7
			_ = _891_v131
			_891_v131 = Companion_D7_.Create_DC26_(_721_v2, _719_v0, _719_v0)
			var _pat_let_tv13 = _721_v2
			_ = _pat_let_tv13
			var _source11 D7 = func(_pat_let12_0 D7) D7 {
				return func(_892_dt__update__tmp_h1 D7) D7 {
					return func(_pat_let13_0 _dafny.Sequence) D7 {
						return func(_893_dt__update_hcf42_h0 _dafny.Sequence) D7 {
							return Companion_D7_.Create_DC26_(_893_dt__update_hcf42_h0, (_892_dt__update__tmp_h1).Dtor_cf43(), (_892_dt__update__tmp_h1).Dtor_cf44())
						}(_pat_let13_0)
					}(_pat_let_tv13)
				}(_pat_let12_0)
			}(_891_v131)
			_ = _source11
			if _source11.Is_DC26() {
				var _894___mcc_h8 _dafny.Sequence = _source11.Get_().(D7_DC26).Cf42
				_ = _894___mcc_h8
				var _895___mcc_h9 _dafny.Int = _source11.Get_().(D7_DC26).Cf43
				_ = _895___mcc_h9
				var _896___mcc_h10 _dafny.Int = _source11.Get_().(D7_DC26).Cf44
				_ = _896___mcc_h10
				var _897_cf44 _dafny.Int = _896___mcc_h10
				_ = _897_cf44
				var _898_cf43 _dafny.Int = _895___mcc_h9
				_ = _898_cf43
				var _899_cf42 _dafny.Sequence = _894___mcc_h8
				_ = _899_cf42
				(globalState).F6 = _719_v0
				var _900_v132 *C0
				_ = _900_v132
				var _nw119 *C0 = New_C0_()
				_ = _nw119
				_nw119.Ctor__()
				_900_v132 = _nw119
				var _rhs115 bool = !(p0)
				_ = _rhs115
				var _rhs116 bool = _dafny.Companion_Sequence_.Equal(_723_v4, _723_v4)
				_ = _rhs116
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				_lhs93.F10 = _rhs115
				_lhs94.F10 = _rhs116
				(globalState).F10 = Companion_Default___.Fm15(!(p0), globalState)
			} else if _source11.Is_DC25() {
				var _901___mcc_h11 _dafny.Set = _source11.Get_().(D7_DC25).Cf41
				_ = _901___mcc_h11
				var _902_cf41 _dafny.Set = _901___mcc_h11
				_ = _902_cf41
				(globalState).F6 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_719_v0)), _719_v0)
				var _903_v133 D9
				_ = _903_v133
				_903_v133 = Companion_D9_.Create_DC31_(Companion_Default___.SafeDivisionInt(_719_v0, _719_v0), _806_v74)
				_903_v133 = _903_v133
				(globalState).F11 = _719_v0
				(globalState).F6 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_729_v9).Cardinality()), Companion_Default___.SafeModuloInt(_719_v0, _719_v0))
			} else {
				var _904___mcc_h12 D7 = _source11.Get_().(D7_DC27).Cf45
				_ = _904___mcc_h12
				var _905_cf45 D7 = _904___mcc_h12
				_ = _905_cf45
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))
				_ = _index122
				(_806_v74).ArraySet1(p0, (_index122).Int())
				var _906_v134 _dafny.MultiSet
				_ = _906_v134
				_906_v134 = _dafny.MultiSetOf(_806_v74, _806_v74)
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))
				_ = _index123
				(_806_v74).ArraySet1(!(((_906_v134).Intersection(_906_v134)).IsSubsetOf((_906_v134).Union(_906_v134))), (_index123).Int())
				(globalState).F11 = _719_v0
				var _907_v135 *C2
				_ = _907_v135
				var _nw120 *C2 = New_C2_()
				_ = _nw120
				_nw120.Ctor__(_719_v0, (_this).F22())
				_907_v135 = _nw120
				var _908_v136 _dafny.Map
				_ = _908_v136
				_908_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_907_v135).F29(), _719_v0)
				var _909_v137 D0
				_ = _909_v137
				_909_v137 = Companion_D0_.Create_DC1_((_907_v135).F29(), (_806_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))).Int()).(bool), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_806_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))).Int()).(bool)), _908_v136), _719_v0, _805_v73)
				var _910_v138 _dafny.Set
				_ = _910_v138
				_910_v138 = _dafny.SetOf(_909_v137)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))
				_ = _index124
				var _rhs117 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_805_v73, _805_v73), _720_v1), (Companion_Default___.SafeIndex(_719_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_805_v73, _805_v73), _720_v1)).Cardinality()))).Uint32(), _805_v73)).Cardinality())).Plus(_719_v0)
				_ = _rhs117
				var _rhs118 bool = (_723_v4).Select((Companion_Default___.SafeIndex((_907_v135).F29(), _dafny.IntOfUint32((_723_v4).Cardinality()))).Uint32()).(bool)
				_ = _rhs118
				var _rhs119 *C2 = _907_v135
				_ = _rhs119
				var _rhs120 bool = (_907_v135).Fm21(_910_v138, _719_v0, globalState)
				_ = _rhs120
				var _rhs121 bool = p0
				_ = _rhs121
				var _lhs95 *GlobalState = globalState
				_ = _lhs95
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				var _lhs97 _dafny.Array = _806_v74
				_ = _lhs97
				var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))
				_ = _lhs98
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				_lhs95.F11 = _rhs117
				_lhs96.F10 = _rhs118
				_907_v135 = _rhs119
				(_lhs97).ArraySet1(_rhs120, (_lhs98).Int())
				_lhs99.F10 = _rhs121
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_806_v74), 0))
				_ = _index125
				(_806_v74).ArraySet1(p0, (_index125).Int())
			}
			var _911_v139 _dafny.Map
			_ = _911_v139
			_911_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _805_v73)
			var _912_v140 _dafny.Map
			_ = _912_v140
			_912_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_765_v35, _719_v0)
			var _913_v141 _dafny.Map
			_ = _913_v141
			_913_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_912_v140, _911_v139)
			var _914_v142 _dafny.Sequence
			_ = _914_v142
			_914_v142 = _dafny.SeqOf(_765_v35)
			var _915_v146 _dafny.Map
			_ = _915_v146
			_915_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _911_v139)
			var _916_v147 _dafny.Array
			_ = _916_v147
			var _nwElement0_26 _dafny.Map = _911_v139
			_ = _nwElement0_26
			var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(26))
			_ = _nw121
			(_nw121).ArraySet1(_nwElement0_26, 0)
			(_nw121).ArraySet1(_911_v139, 1)
			(_nw121).ArraySet1((func() _dafny.Map {
				if (_913_v141).Contains(_912_v140) {
					return (_913_v141).Get(_912_v140).(_dafny.Map)
				}
				return _911_v139
			})(), 2)
			(_nw121).ArraySet1(Companion_Default___.Fm44(p0, !(!(true)), p0, _720_v1, globalState), 3)
			(_nw121).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_914_v142).Cardinality()), _dafny.CodePoint('i'))).Merge(_911_v139), 4)
			(_nw121).ArraySet1(_911_v139, 5)
			(_nw121).ArraySet1(func() _dafny.Map {
				var _coll44 = _dafny.NewMapBuilder()
				_ = _coll44
				for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(47), _dafny.IntOfInt64(90))); ; {
					_compr_44, _ok45 := _iter45()
					if !_ok45 {
						break
					}
					var _917_v143 _dafny.Int
					_917_v143 = interface{}(_compr_44).(_dafny.Int)
					if ((_dafny.IntOfInt64(47)).Cmp(_917_v143) <= 0) && ((_917_v143).Cmp(_dafny.IntOfInt64(90)) < 0) {
						_coll44.Add((_917_v143).Times(_719_v0), _805_v73)
					}
				}
				return _coll44.ToMap()
			}(), 6)
			(_nw121).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(175), _805_v73), 7)
			(_nw121).ArraySet1(_911_v139, 8)
			(_nw121).ArraySet1(_911_v139, 9)
			(_nw121).ArraySet1((func() _dafny.Map {
				var _coll45 = _dafny.NewMapBuilder()
				_ = _coll45
				for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-152), _dafny.IntOfInt64(201))); ; {
					_compr_45, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _918_v144 _dafny.Int
					_918_v144 = interface{}(_compr_45).(_dafny.Int)
					if ((_dafny.IntOfInt64(-152)).Cmp(_918_v144) <= 0) && ((_918_v144).Cmp(_dafny.IntOfInt64(201)) < 0) {
						_coll45.Add((_918_v144).Plus(_719_v0), (func() _dafny.CodePoint {
							if (_911_v139).Contains(_719_v0) {
								return (_911_v139).Get(_719_v0).(_dafny.CodePoint)
							}
							return _805_v73
						})())
					}
				}
				return _coll45.ToMap()
			}()).Merge((_911_v139).Update((_dafny.Zero).Minus(_719_v0), _805_v73)), 10)
			(_nw121).ArraySet1((func() _dafny.Map {
				var _coll46 = _dafny.NewMapBuilder()
				_ = _coll46
				for _iter47 := _dafny.Iterate((_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(978)))).Cardinality(), _719_v0, _dafny.IntOfInt64(314))).Elements()); ; {
					_compr_46, _ok47 := _iter47()
					if !_ok47 {
						break
					}
					var _919_v145 _dafny.Int
					_919_v145 = interface{}(_compr_46).(_dafny.Int)
					if (_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(978)))).Cardinality(), _719_v0, _dafny.IntOfInt64(314))).Contains(_919_v145) {
						_coll46.Add(Companion_Default___.SafeModuloInt(_919_v145, _719_v0), _805_v73)
					}
				}
				return _coll46.ToMap()
			}()).Merge(_911_v139), 11)
			(_nw121).ArraySet1(_911_v139, 12)
			(_nw121).ArraySet1(_911_v139, 13)
			(_nw121).ArraySet1(_911_v139, 14)
			(_nw121).ArraySet1((_911_v139).Merge(_911_v139), 15)
			(_nw121).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _805_v73)).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _dafny.IntOfInt64(242))).Cardinality(), _805_v73), 16)
			(_nw121).ArraySet1((_911_v139).Update(_dafny.IntOfInt64(780), _805_v73), 17)
			(_nw121).ArraySet1((_911_v139).Merge(_911_v139), 18)
			(_nw121).ArraySet1(_911_v139, 19)
			(_nw121).ArraySet1((_911_v139).Merge(_911_v139), 20)
			(_nw121).ArraySet1(_911_v139, 21)
			(_nw121).ArraySet1((func() _dafny.Map {
				if (_915_v146).Contains(true) {
					return (_915_v146).Get(true).(_dafny.Map)
				}
				return (_911_v139).Update(_719_v0, _805_v73)
			})(), 22)
			(_nw121).ArraySet1((_911_v139).Merge(_911_v139), 23)
			(_nw121).ArraySet1((_911_v139).Update(_719_v0, _805_v73), 24)
			(_nw121).ArraySet1(_911_v139, 25)
			_916_v147 = _nw121
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_916_v147), 0))
			_ = _index126
			(_916_v147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _805_v73), (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_916_v147), 0))
			_ = _index127
			(_916_v147).ArraySet1((_911_v139).Merge(_911_v139), (_index127).Int())
		} else {
			var _920___mcc_h7 D4 = _source10.Get_().(D4_DC16).Cf28
			_ = _920___mcc_h7
			var _921_cf28 D4 = _920___mcc_h7
			_ = _921_cf28
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))
			_ = _index128
			(_765_v35).ArraySet1(_719_v0, (_index128).Int())
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))
			_ = _index129
			(_765_v35).ArraySet1((_719_v0).Plus((func() _dafny.Int {
				if p0 {
					return _719_v0
				}
				return _719_v0
			})()), (_index129).Int())
			(globalState).F0 = p0
			var _922_v148 _dafny.MultiSet
			_ = _922_v148
			_922_v148 = _dafny.MultiSetOf(_719_v0, Companion_Default___.SafeDivisionInt(_719_v0, _719_v0))
			_922_v148 = _922_v148
			var _923_v149 _dafny.Array
			_ = _923_v149
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(14))
			_ = _nw122
			_923_v149 = _nw122
			var _924_v150 *C0
			_ = _924_v150
			var _nw123 *C0 = New_C0_()
			_ = _nw123
			_nw123.Ctor__()
			_924_v150 = _nw123
			var _925_v151 _dafny.MultiSet
			_ = _925_v151
			_925_v151 = _dafny.MultiSetOf(_924_v150, _924_v150)
			var _926_v152 D12
			_ = _926_v152
			_926_v152 = Companion_D12_.Create_DC41_(_925_v151)
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_923_v149), 0))
			_ = _index130
			(_923_v149).ArraySet1(_926_v152, (_index130).Int())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))
			_ = _index131
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_923_v149), 0))
			_ = _index132
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))
			_ = _index133
			var _rhs122 _dafny.Int = (_765_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))).Int()).(_dafny.Int)
			_ = _rhs122
			var _rhs123 D12 = _926_v152
			_ = _rhs123
			var _rhs124 bool = p0
			_ = _rhs124
			var _rhs125 bool = p0
			_ = _rhs125
			var _rhs126 _dafny.Int = (_this).Fm0((_922_v148).Intersection(_922_v148), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_720_v1, _720_v1)).Cardinality()), _725_v6, p0, globalState)
			_ = _rhs126
			var _lhs100 _dafny.Array = _765_v35
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))
			_ = _lhs101
			var _lhs102 _dafny.Array = _923_v149
			_ = _lhs102
			var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_923_v149), 0))
			_ = _lhs103
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			var _lhs105 *GlobalState = globalState
			_ = _lhs105
			var _lhs106 _dafny.Array = _765_v35
			_ = _lhs106
			var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_765_v35), 0))
			_ = _lhs107
			(_lhs100).ArraySet1(_rhs122, (_lhs101).Int())
			(_lhs102).ArraySet1(_rhs123, (_lhs103).Int())
			_lhs104.F0 = _rhs124
			_lhs105.F0 = _rhs125
			(_lhs106).ArraySet1(_rhs126, (_lhs107).Int())
		}
		r0 = !((func() bool {
			if !(!(p0)) {
				return p0
			}
			return false
		})())
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_16
		var _nw124 _dafny.Array
		_ = _nw124
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw124 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) bool = (func(_927_p0 bool) func(_dafny.Int) bool {
				return func(_928_i10 _dafny.Int) bool {
					return (_927_p0) == (_927_p0)
				}
			})(p0)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw124 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw124).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw124).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		r1 = _nw124
		var _929_v153 _dafny.Map
		_ = _929_v153
		_929_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v0, _719_v0)
		var _930_v154 D5
		_ = _930_v154
		_930_v154 = Companion_D5_.Create_DC19_(_929_v153, p0)
		var _pat_let_tv14 = _721_v2
		_ = _pat_let_tv14
		var _pat_let_tv15 = _805_v73
		_ = _pat_let_tv15
		var _pat_let_tv16 = _721_v2
		_ = _pat_let_tv16
		var _pat_let_tv17 = _805_v73
		_ = _pat_let_tv17
		var _pat_let_tv18 = _719_v0
		_ = _pat_let_tv18
		var _pat_let_tv19 = p0
		_ = _pat_let_tv19
		var _pat_let_tv20 = _805_v73
		_ = _pat_let_tv20
		var _pat_let_tv21 = _721_v2
		_ = _pat_let_tv21
		var _pat_let_tv22 = _805_v73
		_ = _pat_let_tv22
		var _pat_let_tv23 = _721_v2
		_ = _pat_let_tv23
		var _pat_let_tv24 = _805_v73
		_ = _pat_let_tv24
		var _pat_let_tv25 = _721_v2
		_ = _pat_let_tv25
		var _pat_let_tv26 = _805_v73
		_ = _pat_let_tv26
		var _pat_let_tv27 = _721_v2
		_ = _pat_let_tv27
		var _pat_let_tv28 = _805_v73
		_ = _pat_let_tv28
		var _pat_let_tv29 = _721_v2
		_ = _pat_let_tv29
		var _pat_let_tv30 = _805_v73
		_ = _pat_let_tv30
		var _pat_let_tv31 = _721_v2
		_ = _pat_let_tv31
		var _pat_let_tv32 = _719_v0
		_ = _pat_let_tv32
		var _pat_let_tv33 = _721_v2
		_ = _pat_let_tv33
		var _pat_let_tv34 = _719_v0
		_ = _pat_let_tv34
		var _pat_let_tv35 = _805_v73
		_ = _pat_let_tv35
		r2 = func(_source12 D5) _dafny.Map {
			if _source12.Is_DC18() {
				var _931___mcc_h13 _dafny.Int = _source12.Get_().(D5_DC18).Cf30
				_ = _931___mcc_h13
				var _932___mcc_h14 _dafny.Int = _source12.Get_().(D5_DC18).Cf31
				_ = _932___mcc_h14
				var _933_cf31 _dafny.Int = _932___mcc_h14
				_ = _933_cf31
				var _934_cf30 _dafny.Int = _931___mcc_h13
				_ = _934_cf30
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv14, _pat_let_tv15)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv16, _pat_let_tv17)).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_935_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_936_i11 _dafny.Int) _dafny.Int {
						return _935_v0
					}
				})(_pat_let_tv18))), _dafny.CodePoint('t')))
			} else if _source12.Is_DC19() {
				var _937___mcc_h15 _dafny.Map = _source12.Get_().(D5_DC19).Cf32
				_ = _937___mcc_h15
				var _938___mcc_h16 bool = _source12.Get_().(D5_DC19).Cf33
				_ = _938___mcc_h16
				var _939_cf33 bool = _938___mcc_h16
				_ = _939_cf33
				var _940_cf32 _dafny.Map = _937___mcc_h15
				_ = _940_cf32
				if _pat_let_tv19 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg66 _dafny.Int) interface{} {
							return coer66(arg66)
						}
					}(func(_941_i12 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality())
					})), _pat_let_tv20)
				} else {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv21, _pat_let_tv22)
				}
			} else if _source12.Is_DC17() {
				var _942___mcc_h17 _dafny.Map = _source12.Get_().(D5_DC17).Cf29
				_ = _942___mcc_h17
				var _943_cf29 _dafny.Map = _942___mcc_h17
				_ = _943_cf29
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv23, _pat_let_tv24)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv25, _pat_let_tv26))
			} else {
				var _944___mcc_h18 D5 = _source12.Get_().(D5_DC20).Cf34
				_ = _944___mcc_h18
				var _945_cf34 D5 = _944___mcc_h18
				_ = _945_cf34
				return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv27, _pat_let_tv28)).Update(_pat_let_tv29, _pat_let_tv30)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_pat_let_tv31, (Companion_Default___.SafeIndex(_pat_let_tv32, _dafny.IntOfUint32((_pat_let_tv33).Cardinality()))).Uint32(), _pat_let_tv34), _pat_let_tv35))
			}
		}(_930_v154)
		r3 = _719_v0
		return r0, r1, r2, r3
	}
}
func (_this *C3) M4(globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _946_v0 bool
		_ = _946_v0
		_946_v0 = true
		var _947_v1 _dafny.Set
		_ = _947_v1
		_947_v1 = _dafny.SetOf(Companion_Default___.Fm15(_946_v0, globalState))
		var _948_v2 _dafny.Int
		_ = _948_v2
		_948_v2 = _dafny.IntOfInt64(54)
		var _949_v3 _dafny.MultiSet
		_ = _949_v3
		_949_v3 = _dafny.MultiSetOf(_948_v2, _948_v2, _948_v2)
		var _950_v4 _dafny.MultiSet
		_ = _950_v4
		_950_v4 = _dafny.MultiSetOf(_949_v3)
		var _951_v5 _dafny.Sequence
		_ = _951_v5
		_951_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ftnyfkc")
		var _952_v6 _dafny.Set
		_ = _952_v6
		_952_v6 = _dafny.SetOf(_dafny.IntOfUint32((_951_v5).Cardinality()))
		var _953_v7 _dafny.Map
		_ = _953_v7
		_953_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, _948_v2)
		var _954_v8 _dafny.Sequence
		_ = _954_v8
		_954_v8 = _dafny.SeqOf(_946_v0)
		var _955_v9 _dafny.Sequence
		_ = _955_v9
		_955_v9 = _dafny.SeqOf(_952_v6)
		var _956_v10 _dafny.Array
		_ = _956_v10
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_17
		var _nw125 _dafny.Array
		_ = _nw125
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw125 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) bool = (func(_957_v0 bool) func(_dafny.Int) bool {
				return func(_958_i0 _dafny.Int) bool {
					return _957_v0
				}
			})(_946_v0)
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw125 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw125).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw125).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_956_v10 = _nw125
		var _959_v11 _dafny.Map
		_ = _959_v11
		_959_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, _956_v10)
		var _960_v12 _dafny.Array
		_ = _960_v12
		var _nwElement0_27 _dafny.Int = ((_dafny.MultiSetOf(_949_v3)).Intersection(_950_v4)).Cardinality()
		_ = _nwElement0_27
		var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(12))
		_ = _nw126
		(_nw126).ArraySet1(_nwElement0_27, 0)
		(_nw126).ArraySet1((_952_v6).Cardinality(), 1)
		(_nw126).ArraySet1((func() _dafny.Int {
			if (_953_v7).Contains(_946_v0) {
				return (_953_v7).Get(_946_v0).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_954_v8).Cardinality())
		})(), 2)
		(_nw126).ArraySet1((func() _dafny.Int {
			if !(_946_v0) {
				return _dafny.IntOfUint32((_954_v8).Cardinality())
			}
			return _948_v2
		})(), 3)
		(_nw126).ArraySet1((_this).Fm0(_949_v3, _dafny.IntOfUint32((_951_v5).Cardinality()), _955_v9, _946_v0, globalState), 4)
		(_nw126).ArraySet1((_this).Fm0(_dafny.MultiSetOf(_dafny.IntOfUint32((_954_v8).Cardinality())), _dafny.IntOfInt64(-673), _955_v9, false, globalState), 5)
		(_nw126).ArraySet1(Companion_Default___.SafeModuloInt(_948_v2, _948_v2), 6)
		(_nw126).ArraySet1(((func() _dafny.Map {
			if _946_v0 {
				return _959_v11
			}
			return _959_v11
		})()).Cardinality(), 7)
		(_nw126).ArraySet1(Companion_Default___.SafeModuloInt(_948_v2, _948_v2), 8)
		(_nw126).ArraySet1(_dafny.IntOfUint32((_951_v5).Cardinality()), 9)
		(_nw126).ArraySet1(_948_v2, 10)
		(_nw126).ArraySet1(_948_v2, 11)
		_960_v12 = _nw126
		var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
		_ = _index134
		(_960_v12).ArraySet1((_948_v2).Times(_948_v2), (_index134).Int())
		var _961_v13 _dafny.MultiSet
		_ = _961_v13
		_961_v13 = _dafny.MultiSetOf(_946_v0)
		var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
		_ = _index135
		(_956_v10).ArraySet1((_961_v13).IsProperSubsetOf(_961_v13), (_index135).Int())
		var _962_v14 _dafny.Map
		_ = _962_v14
		_962_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("dkfdq"), _951_v5), _946_v0)
		var _963_v15 _dafny.Sequence
		_ = _963_v15
		_963_v15 = _dafny.SeqOf(_951_v5)
		var _964_v16 _dafny.CodePoint
		_ = _964_v16
		_964_v16 = _dafny.CodePoint('i')
		var _965_v17 D8
		_ = _965_v17
		_965_v17 = Companion_D8_.Create_DC29_(_946_v0, _946_v0, _964_v16)
		var _966_v18 D13
		_ = _966_v18
		_966_v18 = Companion_D13_.Create_DC44_(_954_v8)
		var _967_v19 _dafny.Map
		_ = _967_v19
		_967_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_948_v2, _946_v0)
		var _pat_let_tv36 = _946_v0
		_ = _pat_let_tv36
		var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
		_ = _index136
		var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
		_ = _index137
		var _rhs127 _dafny.Set = _dafny.SetOf((func() bool {
			if (_962_v14).Contains(_963_v15) {
				return (_962_v14).Get(_963_v15).(bool)
			}
			return _946_v0
		})())
		_ = _rhs127
		var _rhs128 _dafny.Int = _dafny.IntOfInt64(971)
		_ = _rhs128
		var _rhs129 bool = _dafny.Companion_Sequence_.Contains(_954_v8, (_948_v2).Cmp((_dafny.SetOf(_965_v17, func(_pat_let14_0 D8) D8 {
			return func(_968_dt__update__tmp_h0 D8) D8 {
				return func(_pat_let15_0 bool) D8 {
					return func(_969_dt__update_hcf47_h0 bool) D8 {
						return Companion_D8_.Create_DC29_(_969_dt__update_hcf47_h0, (_968_dt__update__tmp_h0).Dtor_cf48(), (_968_dt__update__tmp_h0).Dtor_cf49())
					}(_pat_let15_0)
				}(_pat_let_tv36)
			}(_pat_let14_0)
		}(_965_v17), _965_v17)).Cardinality()) <= 0)
		_ = _rhs129
		var _rhs130 bool = _dafny.Companion_Sequence_.Equal(_951_v5, Companion_Default___.Fm18(_946_v0, globalState))
		_ = _rhs130
		var _rhs131 bool = Companion_Default___.Fm37((_966_v18).Dtor_cf71(), (_967_v19).Merge(Companion_Default___.Fm40(_948_v2, (_dafny.Zero).Minus(_948_v2), Companion_Default___.Fm38(true, _948_v2, globalState), globalState)), _951_v5, globalState)
		_ = _rhs131
		var _lhs108 _dafny.Array = _960_v12
		_ = _lhs108
		var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
		_ = _lhs109
		var _lhs110 _dafny.Array = _956_v10
		_ = _lhs110
		var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
		_ = _lhs111
		var _lhs112 *GlobalState = globalState
		_ = _lhs112
		_947_v1 = _rhs127
		(_lhs108).ArraySet1(_rhs128, (_lhs109).Int())
		(_lhs110).ArraySet1(_rhs129, (_lhs111).Int())
		_lhs112.F0 = _rhs130
		_946_v0 = _rhs131
		var _970_v20 _dafny.MultiSet
		_ = _970_v20
		_970_v20 = _dafny.MultiSetOf(_965_v17)
		var _971_v21 D11
		_ = _971_v21
		_971_v21 = Companion_D11_.Create_DC36_(_dafny.SeqOf(_949_v3))
		var _972_v22 _dafny.Sequence
		_ = _972_v22
		_972_v22 = _dafny.SeqOf((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
		var _973_v23 D2
		_ = _973_v23
		_973_v23 = Companion_D2_.Create_DC7_(_972_v22, (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), _948_v2)
		var _pat_let_tv37 = _946_v0
		_ = _pat_let_tv37
		var _pat_let_tv38 = _956_v10
		_ = _pat_let_tv38
		var _pat_let_tv39 = _956_v10
		_ = _pat_let_tv39
		var _pat_let_tv40 = _946_v0
		_ = _pat_let_tv40
		var _pat_let_tv41 = _956_v10
		_ = _pat_let_tv41
		var _pat_let_tv42 = _956_v10
		_ = _pat_let_tv42
		var _pat_let_tv43 = _946_v0
		_ = _pat_let_tv43
		var _pat_let_tv44 = _956_v10
		_ = _pat_let_tv44
		var _pat_let_tv45 = _956_v10
		_ = _pat_let_tv45
		var _pat_let_tv46 = _956_v10
		_ = _pat_let_tv46
		var _pat_let_tv47 = _956_v10
		_ = _pat_let_tv47
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
		_ = _index138
		var _rhs132 bool = !(func(_source13 D11) bool {
			if _source13.Is_DC37() {
				var _974___mcc_h0 _dafny.Array = _source13.Get_().(D11_DC37).Cf63
				_ = _974___mcc_h0
				var _975_cf63 _dafny.Array = _974___mcc_h0
				_ = _975_cf63
				return (true) || (_pat_let_tv37)
			} else if _source13.Is_DC38() {
				var _976___mcc_h1 _dafny.Int = _source13.Get_().(D11_DC38).Cf64
				_ = _976___mcc_h1
				var _977___mcc_h2 _dafny.Int = _source13.Get_().(D11_DC38).Cf65
				_ = _977___mcc_h2
				var _978_cf65 _dafny.Int = _977___mcc_h2
				_ = _978_cf65
				var _979_cf64 _dafny.Int = _976___mcc_h1
				_ = _979_cf64
				if (_pat_let_tv39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_pat_let_tv38), 0))).Int()).(bool) {
					return _pat_let_tv40
				} else {
					return (_pat_let_tv42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_pat_let_tv41), 0))).Int()).(bool)
				}
			} else if _source13.Is_DC39() {
				var _980___mcc_h3 bool = _source13.Get_().(D11_DC39).Cf66
				_ = _980___mcc_h3
				var _981_cf66 bool = _980___mcc_h3
				_ = _981_cf66
				return _pat_let_tv43
			} else if _source13.Is_DC36() {
				var _982___mcc_h4 _dafny.Sequence = _source13.Get_().(D11_DC36).Cf62
				_ = _982___mcc_h4
				var _983_cf62 _dafny.Sequence = _982___mcc_h4
				_ = _983_cf62
				return !(false) || ((_pat_let_tv45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_pat_let_tv44), 0))).Int()).(bool))
			} else {
				var _984___mcc_h5 D11 = _source13.Get_().(D11_DC40).Cf67
				_ = _984___mcc_h5
				var _985_cf67 D11 = _984___mcc_h5
				_ = _985_cf67
				return (_pat_let_tv47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_pat_let_tv46), 0))).Int()).(bool)
			}
		}(_971_v21))
		_ = _rhs132
		var _rhs133 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
		_ = _rhs133
		var _rhs134 _dafny.CodePoint = Companion_Default___.Fm22((_948_v2).Cmp((_973_v23).Dtor_cf13()) < 0, (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), _948_v2, _948_v2, globalState)
		_ = _rhs134
		var _rhs135 _dafny.MultiSet = _970_v20
		_ = _rhs135
		var _lhs113 _dafny.Array = _956_v10
		_ = _lhs113
		var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
		_ = _lhs114
		(_lhs113).ArraySet1(_rhs132, (_lhs114).Int())
		_948_v2 = _rhs133
		_964_v16 = _rhs134
		_970_v20 = _rhs135
		var _986_v24 D9
		_ = _986_v24
		_986_v24 = Companion_D9_.Create_DC31_(_dafny.IntOfInt64(618), _956_v10)
		var _987_v25 D9
		_ = _987_v25
		_987_v25 = Companion_D9_.Create_DC32_(_986_v24)
		var _source14 D9 = _987_v25
		_ = _source14
		if _source14.Is_DC31() {
			var _988___mcc_h6 _dafny.Int = _source14.Get_().(D9_DC31).Cf51
			_ = _988___mcc_h6
			var _989___mcc_h7 _dafny.Array = _source14.Get_().(D9_DC31).Cf52
			_ = _989___mcc_h7
			var _990_cf52 _dafny.Array = _989___mcc_h7
			_ = _990_cf52
			var _991_cf51 _dafny.Int = _988___mcc_h6
			_ = _991_cf51
			var _rhs136 bool = !(_946_v0)
			_ = _rhs136
			var _rhs137 bool = (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
			_ = _rhs137
			var _rhs138 _dafny.Int = _dafny.IntOfInt64(837)
			_ = _rhs138
			var _lhs115 *GlobalState = globalState
			_ = _lhs115
			var _lhs116 *GlobalState = globalState
			_ = _lhs116
			_lhs115.F0 = _rhs136
			_lhs116.F0 = _rhs137
			r1 = _rhs138
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
			_ = _index139
			(_956_v10).ArraySet1((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), (_index139).Int())
			var _992_v26 _dafny.Map
			_ = _992_v26
			_992_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_953_v7, _946_v0)
			_992_v26 = (_992_v26).Update(_953_v7, _946_v0)
			var _993_v27 _dafny.Map
			_ = _993_v27
			_993_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("y"), (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))
			if (func() bool {
				if (_993_v27).Contains(_951_v5) {
					return (_993_v27).Get(_951_v5).(bool)
				}
				return (func() bool {
					if (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool) {
						return !(!((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)))
					}
					return _946_v0
				})()
			})() {
				_954_v8 = _954_v8
				var _994_v28 _dafny.Map
				_ = _994_v28
				_994_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				var _995_v29 D5
				_ = _995_v29
				_995_v29 = Companion_D5_.Create_DC19_(_994_v28, _946_v0)
				var _996_v30 _dafny.Sequence
				_ = _996_v30
				_996_v30 = _dafny.SeqOf(_995_v29, _995_v29, _995_v29)
				var _997_v31 _dafny.Map
				_ = _997_v31
				_997_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_946_v0, (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)), _996_v30)
				var _rhs139 _dafny.Array = _960_v12
				_ = _rhs139
				var _rhs140 _dafny.Array = _956_v10
				_ = _rhs140
				var _rhs141 bool = ((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)) || (_946_v0)
				_ = _rhs141
				var _rhs142 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_947_v1, _996_v30)).Merge(_997_v31)
				_ = _rhs142
				_960_v12 = _rhs139
				r0 = _rhs140
				_946_v0 = _rhs141
				_997_v31 = _rhs142
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index140
				(_956_v10).ArraySet1(_946_v0, (_index140).Int())
				var _998_v32 _dafny.Array
				_ = _998_v32
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw127
				_998_v32 = _nw127
				var _999_v33 _dafny.Map
				_ = _999_v33
				_999_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, _951_v5)
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_998_v32), 0))
				_ = _index141
				(_998_v32).ArraySet1((func() _dafny.Sequence {
					if (_999_v33).Contains(_946_v0) {
						return (_999_v33).Get(_946_v0).(_dafny.Sequence)
					}
					return _951_v5
				})(), (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_998_v32), 0))
				_ = _index142
				(_998_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_951_v5, _951_v5), _dafny.Companion_Sequence_.Update(_951_v5, (Companion_Default___.SafeIndex(_948_v2, _dafny.IntOfUint32((_951_v5).Cardinality()))).Uint32(), _964_v16)), (_index142).Int())
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_998_v32), 0))
				_ = _index143
				(_998_v32).ArraySet1(Companion_Default___.Fm18(_946_v0, globalState), (_index143).Int())
			} else {
				(globalState).F0 = (_946_v0) || (true)
				(globalState).F10 = _946_v0
				var _1000_v34 _dafny.Map
				_ = _1000_v34
				_1000_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_972_v22, _991_cf51)
				var _1001_v35 _dafny.Map
				_ = _1001_v35
				_1001_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, _967_v19)
				var _1002_v36 _dafny.Map
				_ = _1002_v36
				_1002_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_951_v5, (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index144
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index145
				var _rhs143 bool = (_1001_v35).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_972_v22, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_951_v5, (Companion_Default___.SafeIndex((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_951_v5).Cardinality()))).Uint32(), _dafny.CodePoint('b'))).Cardinality()))).Equals(_1000_v34))
				_ = _rhs143
				var _rhs144 bool = _946_v0
				_ = _rhs144
				var _rhs145 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_1002_v36).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg67 _dafny.Int) interface{} {
							return coer67(arg67)
						}
					}((func(_1003_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1004_i1 _dafny.Int) _dafny.CodePoint {
							return _1003_v16
						}
					})(_964_v16)))) {
						return (_1002_v36).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg68 _dafny.Int) interface{} {
								return coer68(arg68)
							}
						}((func(_1005_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1006_i1 _dafny.Int) _dafny.CodePoint {
								return _1005_v16
							}
						})(_964_v16)))).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(_948_v2)
				})(), (func() _dafny.Int {
					if _946_v0 {
						return _991_cf51
					}
					return _991_cf51
				})())
				_ = _rhs145
				var _lhs117 _dafny.Array = _956_v10
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _lhs118
				var _lhs119 _dafny.Array = _960_v12
				_ = _lhs119
				var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _lhs120
				_946_v0 = _rhs143
				(_lhs117).ArraySet1(_rhs144, (_lhs118).Int())
				(_lhs119).ArraySet1(_rhs145, (_lhs120).Int())
				var _1007_v37 D1
				_ = _1007_v37
				_1007_v37 = Companion_D1_.Create_DC3_(_951_v5)
				var _1008_v38 D2
				_ = _1008_v38
				_1008_v38 = Companion_D2_.Create_DC8_(_1007_v37, _960_v12, (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))
				var _1009_v39 _dafny.Sequence
				_ = _1009_v39
				_1009_v39 = _dafny.SeqOf(_960_v12, _960_v12, _960_v12, _960_v12)
				var _1010_v40 _dafny.Array
				_ = _1010_v40
				var _nwElement0_28 _dafny.Array = _960_v12
				_ = _nwElement0_28
				var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(20))
				_ = _nw128
				(_nw128).ArraySet1(_nwElement0_28, 0)
				(_nw128).ArraySet1(_960_v12, 1)
				(_nw128).ArraySet1(_960_v12, 2)
				(_nw128).ArraySet1(_960_v12, 3)
				(_nw128).ArraySet1(_960_v12, 4)
				(_nw128).ArraySet1(_960_v12, 5)
				(_nw128).ArraySet1(_960_v12, 6)
				(_nw128).ArraySet1((_1008_v38).Dtor_cf15(), 7)
				(_nw128).ArraySet1(_960_v12, 8)
				(_nw128).ArraySet1(_960_v12, 9)
				(_nw128).ArraySet1(_960_v12, 10)
				(_nw128).ArraySet1(_960_v12, 11)
				(_nw128).ArraySet1(_960_v12, 12)
				(_nw128).ArraySet1(_960_v12, 13)
				(_nw128).ArraySet1((_1009_v39).Select((Companion_Default___.SafeIndex((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1009_v39).Cardinality()))).Uint32()).(_dafny.Array), 14)
				(_nw128).ArraySet1(_960_v12, 15)
				(_nw128).ArraySet1(_960_v12, 16)
				(_nw128).ArraySet1(_960_v12, 17)
				(_nw128).ArraySet1(_960_v12, 18)
				(_nw128).ArraySet1(_960_v12, 19)
				_1010_v40 = _nw128
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1010_v40), 0))
				_ = _index146
				(_1010_v40).ArraySet1((_1008_v38).Dtor_cf15(), (_index146).Int())
				var _1011_v41 _dafny.Array
				_ = _1011_v41
				var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
				_ = _nw129
				_1011_v41 = _nw129
				var _1012_v42 _dafny.Map
				_ = _1012_v42
				_1012_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_991_cf51), _960_v12)
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1011_v41), 0))
				_ = _index147
				(_1011_v41).ArraySet1(_1012_v42, (_index147).Int())
				var _1013_v43 _dafny.Map
				_ = _1013_v43
				_1013_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_cf51, (_1009_v39).Select((Companion_Default___.SafeIndex(_991_cf51, _dafny.IntOfUint32((_1009_v39).Cardinality()))).Uint32()).(_dafny.Array)))
				var _1014_v44 _dafny.Map
				_ = _1014_v44
				_1014_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, (_954_v8).Select((Companion_Default___.SafeIndex(_991_cf51, _dafny.IntOfUint32((_954_v8).Cardinality()))).Uint32()).(bool))
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1010_v40), 0))
				_ = _index148
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1011_v41), 0))
				_ = _index149
				var _rhs146 _dafny.Array = _960_v12
				_ = _rhs146
				var _rhs147 bool = _946_v0
				_ = _rhs147
				var _rhs148 _dafny.Map = (_1012_v42).Merge(((func() _dafny.Map {
					if (_1013_v43).Contains((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)) {
						return (_1013_v43).Get((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_cf51, _960_v12)
				})()).Merge(_1012_v42))
				_ = _rhs148
				var _rhs149 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_991_cf51, _dafny.IntOfInt64(47)), (_1014_v44).Cardinality())))
				_ = _rhs149
				var _lhs121 _dafny.Array = _1010_v40
				_ = _lhs121
				var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1010_v40), 0))
				_ = _lhs122
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				var _lhs124 _dafny.Array = _1011_v41
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_1011_v41), 0))
				_ = _lhs125
				(_lhs121).ArraySet1(_rhs146, (_lhs122).Int())
				_lhs123.F0 = _rhs147
				(_lhs124).ArraySet1(_rhs148, (_lhs125).Int())
				_991_cf51 = _rhs149
				(globalState).F8 = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
			}
		} else if _source14.Is_DC30() {
			var _1015___mcc_h8 _dafny.Map = _source14.Get_().(D9_DC30).Cf50
			_ = _1015___mcc_h8
			var _1016_cf50 _dafny.Map = _1015___mcc_h8
			_ = _1016_cf50
			r1 = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
			(globalState).F0 = (_954_v8).Select((Companion_Default___.SafeIndex(_948_v2, _dafny.IntOfUint32((_954_v8).Cardinality()))).Uint32()).(bool)
			var _1017_v45 _dafny.Map
			_ = _1017_v45
			_1017_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_948_v2, _949_v3)
			var _1018_v46 D4
			_ = _1018_v46
			_1018_v46 = Companion_D4_.Create_DC12_(_1017_v45)
			_1018_v46 = _1018_v46
			var _1019_v47 _dafny.Array
			_ = _1019_v47
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw130
			_1019_v47 = _nw130
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1019_v47), 0))
			_ = _index150
			(_1019_v47).ArraySet1(_956_v10, (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1019_v47), 0))
			_ = _index151
			(_1019_v47).ArraySet1(_956_v10, (_index151).Int())
		} else {
			var _1020___mcc_h9 D9 = _source14.Get_().(D9_DC32).Cf53
			_ = _1020___mcc_h9
			var _1021_cf53 D9 = _1020___mcc_h9
			_ = _1021_cf53
			_948_v2 = _948_v2
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
			_ = _index152
			var _rhs150 _dafny.Int = ((_952_v6).Cardinality()).Minus(_dafny.IntOfUint32((_972_v22).Cardinality()))
			_ = _rhs150
			var _rhs151 bool = _946_v0
			_ = _rhs151
			var _lhs126 *GlobalState = globalState
			_ = _lhs126
			var _lhs127 _dafny.Array = _956_v10
			_ = _lhs127
			var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
			_ = _lhs128
			_lhs126.F16 = _rhs150
			(_lhs127).ArraySet1(_rhs151, (_lhs128).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
			_ = _index153
			var _rhs152 bool = _946_v0
			_ = _rhs152
			var _rhs153 _dafny.Int = ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Times((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
			_ = _rhs153
			var _lhs129 *GlobalState = globalState
			_ = _lhs129
			var _lhs130 _dafny.Array = _960_v12
			_ = _lhs130
			var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
			_ = _lhs131
			_lhs129.F10 = _rhs152
			(_lhs130).ArraySet1(_rhs153, (_lhs131).Int())
			_951_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ia")
		}
		var _1022_v48 _dafny.Map
		_ = _1022_v48
		_1022_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_948_v2).Times(_948_v2), Companion_Default___.SafeModuloInt((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _948_v2))
		_1022_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_948_v2, (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
		if (_948_v2).Cmp(((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(-544))) >= 0 {
			var _1023_v49 _dafny.Array
			_ = _1023_v49
			var _nw131 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
			_ = _nw131
			_1023_v49 = _nw131
			var _1024_v50 D11
			_ = _1024_v50
			_1024_v50 = Companion_D11_.Create_DC37_(_1023_v49)
			var _source15 D11 = _1024_v50
			_ = _source15
			if _source15.Is_DC37() {
				var _1025___mcc_h10 _dafny.Array = _source15.Get_().(D11_DC37).Cf63
				_ = _1025___mcc_h10
				var _1026_cf63 _dafny.Array = _1025___mcc_h10
				_ = _1026_cf63
				var _rhs154 bool = (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
				_ = _rhs154
				var _rhs155 bool = _946_v0
				_ = _rhs155
				var _rhs156 _dafny.Int = (_dafny.IntOfInt64(212)).Times(_948_v2)
				_ = _rhs156
				var _rhs157 _dafny.Int = (func() _dafny.Int {
					if (_949_v3).Contains((_this).Fm0(_949_v3, _dafny.IntOfInt64(467), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-839))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
						return func(arg69 _dafny.Int) interface{} {
							return coer69(arg69)
						}
					}((func(_1027_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1028_i2 _dafny.Int) _dafny.Set {
							return _1027_v6
						}
					})(_952_v6))), !(_946_v0), globalState)) {
						return (_949_v3).Multiplicity((_this).Fm0(_949_v3, _dafny.IntOfInt64(467), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-839))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
							return func(arg70 _dafny.Int) interface{} {
								return coer70(arg70)
							}
						}((func(_1029_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
							return func(_1030_i2 _dafny.Int) _dafny.Set {
								return _1029_v6
							}
						})(_952_v6))), !(_946_v0), globalState))
					}
					return _dafny.IntOfInt64(750)
				})()
				_ = _rhs157
				var _rhs158 _dafny.Int = (_948_v2).Minus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				_ = _rhs158
				var _lhs132 *GlobalState = globalState
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				var _lhs135 *GlobalState = globalState
				_ = _lhs135
				var _lhs136 *GlobalState = globalState
				_ = _lhs136
				_lhs132.F0 = _rhs154
				_lhs133.F0 = _rhs155
				_lhs134.F6 = _rhs156
				_lhs135.F6 = _rhs157
				_lhs136.F11 = _rhs158
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index154
				var _rhs159 _dafny.Int = _948_v2
				_ = _rhs159
				var _rhs160 bool = ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Cmp(_948_v2) == 0
				_ = _rhs160
				var _rhs161 _dafny.Int = ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Times((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				_ = _rhs161
				var _lhs137 _dafny.Array = _956_v10
				_ = _lhs137
				var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				r1 = _rhs159
				(_lhs137).ArraySet1(_rhs160, (_lhs138).Int())
				_lhs139.F16 = _rhs161
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index155
				(_956_v10).ArraySet1(Companion_Default___.Fm37(_954_v8, (_967_v19).Update(_948_v2, _946_v0), _951_v5, globalState), (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index156
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index157
				var _rhs162 bool = (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
				_ = _rhs162
				var _rhs163 _dafny.Int = _948_v2
				_ = _rhs163
				var _rhs164 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
				_ = _rhs164
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				var _lhs141 _dafny.Array = _960_v12
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _lhs142
				var _lhs143 _dafny.Array = _960_v12
				_ = _lhs143
				var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _lhs144
				_lhs140.F10 = _rhs162
				(_lhs141).ArraySet1(_rhs163, (_lhs142).Int())
				(_lhs143).ArraySet1(_rhs164, (_lhs144).Int())
			} else if _source15.Is_DC38() {
				var _1031___mcc_h11 _dafny.Int = _source15.Get_().(D11_DC38).Cf64
				_ = _1031___mcc_h11
				var _1032___mcc_h12 _dafny.Int = _source15.Get_().(D11_DC38).Cf65
				_ = _1032___mcc_h12
				var _1033_cf65 _dafny.Int = _1032___mcc_h12
				_ = _1033_cf65
				var _1034_cf64 _dafny.Int = _1031___mcc_h11
				_ = _1034_cf64
				_967_v19 = func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter48 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(925), _1033_cf65)).Elements()); ; {
						_compr_47, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1035_v51 _dafny.Int
						_1035_v51 = interface{}(_compr_47).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(925), _1033_cf65), _1035_v51) {
							_coll47.Add((_1035_v51).Plus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)), (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))
						}
					}
					return _coll47.ToMap()
				}()
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index158
				(_956_v10).ArraySet1(_946_v0, (_index158).Int())
				_959_v11 = (_959_v11).Update((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), _956_v10)
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index159
				var _rhs165 bool = (_948_v2).Cmp(_1034_cf64) > 0
				_ = _rhs165
				var _rhs166 _dafny.Int = _1033_cf65
				_ = _rhs166
				var _rhs167 _dafny.Int = (func() _dafny.Int {
					if (_1022_v48).Contains(_1034_cf64) {
						return (_1022_v48).Get(_1034_cf64).(_dafny.Int)
					}
					return ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Minus(_1034_cf64)
				})()
				_ = _rhs167
				var _rhs168 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
				_ = _rhs168
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				var _lhs147 _dafny.Array = _960_v12
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _lhs148
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				_lhs145.F10 = _rhs165
				_lhs146.F11 = _rhs166
				(_lhs147).ArraySet1(_rhs167, (_lhs148).Int())
				_lhs149.F6 = _rhs168
			} else if _source15.Is_DC39() {
				var _1036___mcc_h13 bool = _source15.Get_().(D11_DC39).Cf66
				_ = _1036___mcc_h13
				var _1037_cf66 bool = _1036___mcc_h13
				_ = _1037_cf66
				var _1038_v52 _dafny.Map
				_ = _1038_v52
				_1038_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_cf66, _946_v0)
				var _1039_v53 _dafny.Sequence
				_ = _1039_v53
				_1039_v53 = _dafny.SeqOf(_954_v8)
				var _1040_v54 _dafny.Map
				_ = _1040_v54
				_1040_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1038_v52, !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_954_v8), _1039_v53))
				var _1041_v55 _dafny.Sequence
				_ = _1041_v55
				_1041_v55 = _dafny.SeqOf(Companion_D6_.Create_DC24_(Companion_D6_.Create_DC21_(_dafny.MultiSetOf((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_954_v8).Cardinality()), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)))))
				var _1042_v56 T0
				_ = _1042_v56
				var _nw132 *C2 = New_C2_()
				_ = _nw132
				_nw132.Ctor__(_dafny.IntOfUint32((_951_v5).Cardinality()), (_this).F22())
				_1042_v56 = _nw132
				var _1043_v57 D6
				_ = _1043_v57
				_1043_v57 = Companion_D6_.Create_DC23_(_1037_cf66, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(35), _1042_v56)).Cardinality(), _946_v0)
				var _1044_v58 D6
				_ = _1044_v58
				_1044_v58 = Companion_D6_.Create_DC24_(_1043_v57)
				var _rhs169 bool = _dafny.Companion_Sequence_.Equal(_1041_v55, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1041_v55, _1041_v55), (Companion_Default___.SafeIndex((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1041_v55, _1041_v55)).Cardinality()))).Uint32(), _1044_v58))
				_ = _rhs169
				var _rhs170 _dafny.Int = (func() _dafny.Int {
					if (func() bool {
						if (_967_v19).Contains(_948_v2) {
							return (_967_v19).Get(_948_v2).(bool)
						}
						return (_954_v8).Select((Companion_Default___.SafeIndex(_948_v2, _dafny.IntOfUint32((_954_v8).Cardinality()))).Uint32()).(bool)
					})() {
						return (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
					}
					return _948_v2
				})()
				_ = _rhs170
				var _rhs171 _dafny.Sequence = _972_v22
				_ = _rhs171
				var _rhs172 _dafny.Map = func() _dafny.Map {
					var _coll48 = _dafny.NewMapBuilder()
					_ = _coll48
					for _iter49 := _dafny.Iterate((_1040_v54).Keys().Elements()); ; {
						_compr_48, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1045_v59 _dafny.Map
						_1045_v59 = interface{}(_compr_48).(_dafny.Map)
						if (_1040_v54).Contains(_1045_v59) {
							_coll48.Add(_1045_v59, _946_v0)
						}
					}
					return _coll48.ToMap()
				}()
				_ = _rhs172
				var _rhs173 _dafny.Sequence = _951_v5
				_ = _rhs173
				var _lhs150 *GlobalState = globalState
				_ = _lhs150
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				_lhs150.F0 = _rhs169
				_lhs151.F8 = _rhs170
				_972_v22 = _rhs171
				_1040_v54 = _rhs172
				_951_v5 = _rhs173
				_1038_v52 = (_1038_v52).Update((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), !((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)) || ((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)))
				var _1046_v60 _dafny.Array
				_ = _1046_v60
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(28))
				_ = _nw133
				_1046_v60 = _nw133
				var _1047_v61 _dafny.Sequence
				_ = _1047_v61
				_1047_v61 = _dafny.SeqOf(_1046_v60, _1046_v60, _1046_v60)
				_1046_v60 = (_1047_v61).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_948_v2, (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_1047_v61).Cardinality()))).Uint32()).(_dafny.Array)
				var _1048_v62 D5
				_ = _1048_v62
				_1048_v62 = Companion_D5_.Create_DC18_(_948_v2, (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				var _1049_v63 _dafny.Sequence
				_ = _1049_v63
				_1049_v63 = _dafny.SeqOf(_dafny.SeqOf(_1048_v62, _1048_v62, _1048_v62))
				(globalState).F0 = !(_dafny.Companion_Sequence_.IsPrefixOf(_1049_v63, _1049_v63))
			} else if _source15.Is_DC36() {
				var _1050___mcc_h14 _dafny.Sequence = _source15.Get_().(D11_DC36).Cf62
				_ = _1050___mcc_h14
				var _1051_cf62 _dafny.Sequence = _1050___mcc_h14
				_ = _1051_cf62
				r1 = (func() _dafny.Int {
					if _946_v0 {
						return (_dafny.Zero).Minus(_948_v2)
					}
					return (_this).Fm0(_dafny.MultiSetOf((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _948_v2), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_946_v0, false)).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-26))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
						return func(arg71 _dafny.Int) interface{} {
							return coer71(arg71)
						}
					}((func(_1052_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1053_i3 _dafny.Int) _dafny.Set {
							return _1052_v6
						}
					})(_952_v6))), _946_v0, globalState)
				})()
				(globalState).F6 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_948_v2, _948_v2))
				r0 = _956_v10
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index160
				(_960_v12).ArraySet1((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), (_index160).Int())
			} else {
				var _1054___mcc_h15 D11 = _source15.Get_().(D11_DC40).Cf67
				_ = _1054___mcc_h15
				var _1055_cf67 D11 = _1054___mcc_h15
				_ = _1055_cf67
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index161
				(_956_v10).ArraySet1((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), (_index161).Int())
				(globalState).F0 = _946_v0
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index162
				(_960_v12).ArraySet1(_dafny.IntOfInt64(411), (_index162).Int())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index163
				var _rhs174 bool = (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
				_ = _rhs174
				var _rhs175 _dafny.Int = _948_v2
				_ = _rhs175
				var _rhs176 _dafny.Set = (func() _dafny.Set {
					if (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool) {
						return _947_v1
					}
					return (func() _dafny.Set {
						if _946_v0 {
							return _dafny.SetOf((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))
						}
						return _947_v1
					})()
				})()
				_ = _rhs176
				var _rhs177 _dafny.Sequence = _dafny.SeqOf((func() _dafny.Int {
					if _946_v0 {
						return (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
					}
					return _948_v2
				})(), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				_ = _rhs177
				var _rhs178 bool = (_954_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-182), _dafny.IntOfUint32((_954_v8).Cardinality()))).Uint32()).(bool)
				_ = _rhs178
				var _lhs152 _dafny.Array = _956_v10
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _lhs153
				var _lhs154 *GlobalState = globalState
				_ = _lhs154
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				(_lhs152).ArraySet1(_rhs174, (_lhs153).Int())
				_lhs154.F16 = _rhs175
				_947_v1 = _rhs176
				_972_v22 = _rhs177
				_lhs155.F0 = _rhs178
			}
			if ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Cmp(_948_v2) >= 0 {
				(globalState).F11 = (_948_v2).Minus(((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Plus((_this).Fm0(_949_v3, _dafny.IntOfUint32((_954_v8).Cardinality()), _955_v9, Companion_Default___.Fm15(_946_v0, globalState), globalState)))
				var _1056_v64 *C2
				_ = _1056_v64
				var _nw134 *C2 = New_C2_()
				_ = _nw134
				_nw134.Ctor__((_dafny.Zero).Minus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Concatenate((_this).F22(), _dafny.SeqOf(_956_v10)))
				_1056_v64 = _nw134
				_951_v5 = _951_v5
				var _1057_v65 _dafny.Array
				_ = _1057_v65
				var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
				_ = _nw135
				_1057_v65 = _nw135
				_1057_v65 = _1057_v65
				var _1058_v66 *C2
				_ = _1058_v66
				var _nw136 *C2 = New_C2_()
				_ = _nw136
				_nw136.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1059_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1060_i4 _dafny.Int) _dafny.CodePoint {
						return _1059_v16
					}
				})(_964_v16)))).Cardinality()), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)), (_this).F22())
				_1058_v66 = _nw136
			} else {
				_951_v5 = _951_v5
				(globalState).F11 = _948_v2
				(globalState).F16 = _dafny.IntOfInt64(-93)
				(globalState).F0 = (_dafny.SetOf(_dafny.IntOfInt64(606))).IsSubsetOf(_952_v6)
				var _1061_v67 _dafny.Map
				_ = _1061_v67
				_1061_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, _947_v1)
				_1061_v67 = _1061_v67
			}
			(globalState).F0 = _946_v0
			_948_v2 = _dafny.IntOfInt64(484)
			_951_v5 = _dafny.Companion_Sequence_.Concatenate(_951_v5, (_963_v15).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_948_v2), _dafny.IntOfUint32((_963_v15).Cardinality()))).Uint32()).(_dafny.Sequence))
		} else {
			if (func() bool {
				if ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_951_v5).Cardinality())) < 0 {
					return (_948_v2).Cmp((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)) > 0
				}
				return (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
			})() {
				(globalState).F6 = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index164
				(_960_v12).ArraySet1((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
				_ = _index165
				(_960_v12).ArraySet1(Companion_Default___.SafeDivisionInt((_this).Fm0(_949_v3, _948_v2, _955_v9, true, globalState), _dafny.IntOfInt64(511)), (_index165).Int())
				var _1062_v68 _dafny.Array
				_ = _1062_v68
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_18
				var _nw137 _dafny.Array
				_ = _nw137
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw137 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.CodePoint = (func(_1063_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1064_i5 _dafny.Int) _dafny.CodePoint {
							return _1063_v16
						}
					})(_964_v16)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw137 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw137).ArraySet1CodePoint(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw137).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_1062_v68 = _nw137
				var _1065_v69 _dafny.MultiSet
				_ = _1065_v69
				_1065_v69 = _dafny.MultiSetOf(_1062_v68, _1062_v68)
				var _rhs179 bool = _946_v0
				_ = _rhs179
				var _rhs180 bool = _946_v0
				_ = _rhs180
				var _rhs181 _dafny.Int = (_dafny.Zero).Minus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				_ = _rhs181
				var _rhs182 bool = ((func() _dafny.MultiSet {
					if (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool) {
						return _1065_v69
					}
					return ((_1065_v69).Update(_1062_v68, Companion_Default___.Abs(_dafny.IntOfInt64(990)))).Update(_1062_v68, Companion_Default___.Abs((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)))
				})()).IsDisjointFrom((_1065_v69).Update(_1062_v68, Companion_Default___.Abs((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))))
				_ = _rhs182
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				_lhs156.F10 = _rhs179
				_lhs157.F0 = _rhs180
				_lhs158.F8 = _rhs181
				_lhs159.F10 = _rhs182
				_967_v19 = (_967_v19).Update((_949_v3).Cardinality(), _946_v0)
			} else {
				var _1066_v70 D2
				_ = _1066_v70
				_1066_v70 = Companion_D2_.Create_DC6_(_948_v2)
				var _1067_v71 _dafny.Map
				_ = _1067_v71
				_1067_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), _dafny.MultiSetOf((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(928)))
				var _1068_v72 _dafny.Map
				_ = _1068_v72
				_1068_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_961_v13).Update((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), Companion_Default___.Abs((_this).Fm0((func() _dafny.MultiSet {
					if (_1067_v71).Contains(_946_v0) {
						return (_1067_v71).Get(_946_v0).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_972_v22)
				})(), _948_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1069_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_1070_i6 _dafny.Int) _dafny.Set {
						return _1069_v6
					}
				})(_952_v6))), (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), globalState))), true)
				var _pat_let_tv48 = _972_v22
				_ = _pat_let_tv48
				var _rhs183 _dafny.Int = _dafny.IntOfInt64(130)
				_ = _rhs183
				var _rhs184 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
				_ = _rhs184
				var _rhs185 _dafny.Int = (_1068_v72).Cardinality()
				_ = _rhs185
				var _rhs186 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
				_ = _rhs186
				var _rhs187 D2 = func(_pat_let16_0 D2) D2 {
					return func(_1071_dt__update__tmp_h1 D2) D2 {
						return func(_pat_let17_0 _dafny.Int) D2 {
							return func(_1072_dt__update_hcf9_h0 _dafny.Int) D2 {
								return Companion_D2_.Create_DC6_(_1072_dt__update_hcf9_h0)
							}(_pat_let17_0)
						}(_dafny.IntOfUint32((_pat_let_tv48).Cardinality()))
					}(_pat_let16_0)
				}(Companion_Default___.Fm45((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), globalState))
				_ = _rhs187
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				_lhs160.F11 = _rhs183
				_lhs161.F16 = _rhs184
				_lhs162.F11 = _rhs185
				_lhs163.F6 = _rhs186
				_1066_v70 = _rhs187
				_961_v13 = (_dafny.MultiSetOf(true)).Intersection(_961_v13)
				var _1073_v73 *C1
				_ = _1073_v73
				var _nw138 *C1 = New_C1_()
				_ = _nw138
				_nw138.Ctor__(((_1022_v48).Cardinality()).Plus(_dafny.IntOfUint32((_951_v5).Cardinality())), (_this).F22())
				_1073_v73 = _nw138
				var _1074_v74 *C0
				_ = _1074_v74
				var _nw139 *C0 = New_C0_()
				_ = _nw139
				_nw139.Ctor__()
				_1074_v74 = _nw139
				var _1075_v75 *C0
				_ = _1075_v75
				var _nw140 *C0 = New_C0_()
				_ = _nw140
				_nw140.Ctor__()
				_1075_v75 = _nw140
			}
			var _1076_v76 D5
			_ = _1076_v76
			_1076_v76 = Companion_D5_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_951_v5, (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)))
			var _source16 D5 = _1076_v76
			_ = _source16
			if _source16.Is_DC18() {
				var _1077___mcc_h16 _dafny.Int = _source16.Get_().(D5_DC18).Cf30
				_ = _1077___mcc_h16
				var _1078___mcc_h17 _dafny.Int = _source16.Get_().(D5_DC18).Cf31
				_ = _1078___mcc_h17
				var _1079_cf31 _dafny.Int = _1078___mcc_h17
				_ = _1079_cf31
				var _1080_cf30 _dafny.Int = _1077___mcc_h16
				_ = _1080_cf30
				(globalState).F8 = _1080_cf30
				var _1081_v77 *C2
				_ = _1081_v77
				var _nw141 *C2 = New_C2_()
				_ = _nw141
				_nw141.Ctor__(_948_v2, (_this).F22())
				_1081_v77 = _nw141
				(globalState).F16 = _948_v2
				var _1082_v78 _dafny.Map
				_ = _1082_v78
				_1082_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v12, !(true) || (_946_v0))
				_1082_v78 = (_1082_v78).Update(_960_v12, ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(417)) < 0)
			} else if _source16.Is_DC19() {
				var _1083___mcc_h18 _dafny.Map = _source16.Get_().(D5_DC19).Cf32
				_ = _1083___mcc_h18
				var _1084___mcc_h19 bool = _source16.Get_().(D5_DC19).Cf33
				_ = _1084___mcc_h19
				var _1085_cf33 bool = _1084___mcc_h19
				_ = _1085_cf33
				var _1086_cf32 _dafny.Map = _1083___mcc_h18
				_ = _1086_cf32
				_947_v1 = _947_v1
				var _1087_v79 *C0
				_ = _1087_v79
				var _nw142 *C0 = New_C0_()
				_ = _nw142
				_nw142.Ctor__()
				_1087_v79 = _nw142
				(globalState).F10 = _1085_cf33
				_948_v2 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_951_v5, _951_v5)).Cardinality())).Minus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
			} else if _source16.Is_DC17() {
				var _1088___mcc_h20 _dafny.Map = _source16.Get_().(D5_DC17).Cf29
				_ = _1088___mcc_h20
				var _1089_cf29 _dafny.Map = _1088___mcc_h20
				_ = _1089_cf29
				(globalState).F0 = (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
				_946_v0 = false
				(globalState).F0 = _946_v0
				var _1090_v80 _dafny.Set
				_ = _1090_v80
				_1090_v80 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("xbtvyab"))
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index166
				(_956_v10).ArraySet1((_1090_v80).IsProperSubsetOf(_1090_v80), (_index166).Int())
			} else {
				var _1091___mcc_h21 D5 = _source16.Get_().(D5_DC20).Cf34
				_ = _1091___mcc_h21
				var _1092_cf34 D5 = _1091___mcc_h21
				_ = _1092_cf34
				(globalState).F11 = (_dafny.Zero).Minus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				var _1093_v81 D10
				_ = _1093_v81
				_1093_v81 = Companion_D10_.Create_DC34_(_964_v16, _946_v0, _946_v0, false, _946_v0)
				var _1094_v82 _dafny.Map
				_ = _1094_v82
				_1094_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(164))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}((func(_1095_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1096_i7 _dafny.Int) _dafny.CodePoint {
						return _1095_v16
					}
				})(_964_v16))), (_1093_v81).Dtor_cf56())
				(globalState).F6 = Companion_Default___.SafeModuloInt((_1094_v82).Cardinality(), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
				_951_v5 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1097_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1098_i8 _dafny.Int) _dafny.CodePoint {
						return _1097_v16
					}
				})(_964_v16))), (Companion_Default___.SafeIndex(((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Plus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1099_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1100_i8 _dafny.Int) _dafny.CodePoint {
						return _1099_v16
					}
				})(_964_v16)))).Cardinality()))).Uint32(), _964_v16), (Companion_Default___.SafeIndex((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1101_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1102_i8 _dafny.Int) _dafny.CodePoint {
						return _1101_v16
					}
				})(_964_v16))), (Companion_Default___.SafeIndex(((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Plus((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1103_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1104_i8 _dafny.Int) _dafny.CodePoint {
						return _1103_v16
					}
				})(_964_v16)))).Cardinality()))).Uint32(), _964_v16)).Cardinality()))).Uint32(), _964_v16)
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _index167
				var _rhs188 _dafny.Array = _960_v12
				_ = _rhs188
				var _rhs189 bool = !(!((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))) || (_946_v0)
				_ = _rhs189
				var _lhs164 _dafny.Array = _956_v10
				_ = _lhs164
				var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))
				_ = _lhs165
				_960_v12 = _rhs188
				(_lhs164).ArraySet1(_rhs189, (_lhs165).Int())
			}
			var _1105_v83 _dafny.Map
			_ = _1105_v83
			_1105_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_954_v8).Cardinality()), _953_v7)
			(globalState).F8 = (((_1105_v83).Update(_dafny.IntOfUint32((_951_v5).Cardinality()), _953_v7)).Merge(_1105_v83)).Cardinality()
			var _1106_v84 D6
			_ = _1106_v84
			_1106_v84 = Companion_D6_.Create_DC23_(!((_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))
			_1106_v84 = _1106_v84
			var _1107_v85 _dafny.Sequence
			_ = _1107_v85
			_1107_v85 = _dafny.SeqOf(_960_v12, _960_v12)
			var _1108_v86 D8
			_ = _1108_v86
			_1108_v86 = Companion_D8_.Create_DC28_(_dafny.Companion_Sequence_.Concatenate(_1107_v85, _dafny.SeqOf(_960_v12)))
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
			_ = _index168
			var _rhs190 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
			_ = _rhs190
			var _rhs191 _dafny.Int = (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)
			_ = _rhs191
			var _rhs192 D8 = _1108_v86
			_ = _rhs192
			var _rhs193 _dafny.Int = ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Times((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int))
			_ = _rhs193
			var _lhs166 *GlobalState = globalState
			_ = _lhs166
			var _lhs167 *GlobalState = globalState
			_ = _lhs167
			var _lhs168 _dafny.Array = _960_v12
			_ = _lhs168
			var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))
			_ = _lhs169
			_lhs166.F16 = _rhs190
			_lhs167.F16 = _rhs191
			_1108_v86 = _rhs192
			(_lhs168).ArraySet1(_rhs193, (_lhs169).Int())
		}
		var _1109_i9 _dafny.Int
		_ = _1109_i9
		_1109_i9 = _dafny.Zero
		{
			for Companion_Default___.Fm37(_954_v8, _967_v19, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(964))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg79 _dafny.Int) interface{} {
					return coer79(arg79)
				}
			}(func(_1113_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(964))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg80 _dafny.Int) interface{} {
					return coer80(arg80)
				}
			}(func(_1114_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			}))).Cardinality()))).Uint32(), _964_v16), globalState) {
				{
					if (_1109_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1109_i9 = (_1109_i9).Plus(_dafny.One)
					var _1110_v87 _dafny.Array
					_ = _1110_v87
					var _nwElement0_29 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_951_v5, _951_v5)
					_ = _nwElement0_29
					var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(2))
					_ = _nw143
					(_nw143).ArraySet1(_nwElement0_29, 0)
					(_nw143).ArraySet1(_951_v5, 1)
					_1110_v87 = _nw143
					_1110_v87 = _1110_v87
					(globalState).F0 = true
					(globalState).F10 = (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool)
					var _1111_v88 _dafny.Map
					_ = _1111_v88
					_1111_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool))
					var _1112_v89 D12
					_ = _1112_v89
					_1112_v89 = Companion_D12_.Create_DC42_(_948_v2)
					_1111_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v0, ((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Cmp((_this).Fm0(_dafny.MultiSetOf((_1112_v89).Dtor_cf69()), (_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int), _955_v9, (_956_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_956_v10), 0))).Int()).(bool), globalState)) > 0)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		r0 = _956_v10
		r1 = (((_960_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_960_v12), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(65))).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nlauvedc")).Cardinality()))
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f22 _dafny.Sequence
	F28  bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f22 = _dafny.EmptySeq
	_this.F28 = false
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

func (_this *C4) F22() _dafny.Sequence {
	return _this._f22
}
func (_this *C4) Ctor__(f28 bool, f22 _dafny.Sequence) {
	{
		(_this).F28 = f28
		(_this)._f22 = f22
	}
}
func (_this *C4) Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(955)
	}
}
func (_this *C4) Fm13(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C4) M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1115_v0 _dafny.Sequence
		_ = _1115_v0
		_1115_v0 = _dafny.SeqOf(!(p0))
		var _1116_v1 _dafny.Int
		_ = _1116_v1
		_1116_v1 = _dafny.IntOfInt64(508)
		var _1117_v2 _dafny.MultiSet
		_ = _1117_v2
		_1117_v2 = _dafny.MultiSetOf(p0)
		(globalState).F5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1115_v0, _dafny.Companion_Sequence_.Update(_1115_v0, (Companion_Default___.SafeIndex(_1116_v1, _dafny.IntOfUint32((_1115_v0).Cardinality()))).Uint32(), (_this).Fm13(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(170))).Uint32(), func(coer81 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}((func(_1118_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1119_i0 _dafny.Int) _dafny.Int {
				return _1118_v1
			}
		})(_1116_v1))), p0, (_1117_v2).Cardinality(), globalState))), (func() _dafny.Sequence {
			if _this.F28 {
				return _1115_v0
			}
			return _dafny.SeqOf(_this.F28)
		})())
		var _1120_v3 _dafny.Array
		_ = _1120_v3
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_19
		var _nw144 _dafny.Array
		_ = _nw144
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw144 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) bool = func(_1121_i2 _dafny.Int) bool {
				return !(false)
			}
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw144 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw144).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw144).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_1120_v3 = _nw144
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1120_v3), 0))); ; {
			_guard_loop_1, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _1122_i1 _dafny.Int
			_1122_i1 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_1122_i1).Sign() != -1) && ((_1122_i1).Cmp(_dafny.ArrayLen((_1120_v3), 0)) < 0)) {
				(_1120_v3).ArraySet1(p0, (_1122_i1).Int())
			}
		}
		_1117_v2 = _1117_v2
		var _1123_v4 *C1
		_ = _1123_v4
		var _nw145 *C1 = New_C1_()
		_ = _nw145
		_nw145.Ctor__((_dafny.IntOfUint32((_1115_v0).Cardinality())).Minus((_dafny.Zero).Minus(_1116_v1)), (_this).F22())
		_1123_v4 = _nw145
		var _1124_v5 _dafny.Sequence
		_ = _1124_v5
		_1124_v5 = _dafny.UnicodeSeqOfUtf8Bytes("kqfjidb")
		_1124_v5 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-549))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1125_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _1124_v5)
		_1120_v3 = (func() _dafny.Array {
			if p0 {
				return _1120_v3
			}
			return _1120_v3
		})()
		r0 = _this.F28
		r1 = _1120_v3
		var _1126_v6 _dafny.Sequence
		_ = _1126_v6
		_1126_v6 = _dafny.SeqOf(_1116_v1, _1116_v1, (_1123_v4).F30())
		var _1127_v7 _dafny.Sequence
		_ = _1127_v7
		_1127_v7 = _dafny.SeqOf((_1123_v4).F30(), _dafny.IntOfUint32((_1126_v6).Cardinality()), (_1123_v4).F30(), (_1123_v4).F30(), _1116_v1)
		var _1128_v8 _dafny.CodePoint
		_ = _1128_v8
		_1128_v8 = _dafny.CodePoint('l')
		var _1129_v9 D10
		_ = _1129_v9
		_1129_v9 = Companion_D10_.Create_DC34_(_1128_v8, _this.F28, p0, _this.F28, p0)
		var _1130_v10 _dafny.Map
		_ = _1130_v10
		_1130_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1123_v4).F30(), _1116_v1), _1126_v6), (_1129_v9).Dtor_cf55())
		r2 = _1130_v10
		r3 = (_1123_v4).Fm35((_1123_v4).F30(), (_1123_v4).F30(), p0, (_1123_v4).F30(), globalState)
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f26 _dafny.Map
	_f27 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f26 = _dafny.EmptyMap
	_this._f27 = _dafny.Zero
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f26 _dafny.Map, f27 _dafny.Int) {
	{
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *C5) Fm4(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(!(false)))
	}
}
func (_this *C5) Fm5(p0 bool, globalState *GlobalState) bool {
	{
		return !((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(511)))).Contains((_this).F27())) || (((_this).F27()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf((_this).F27())).Cardinality(), (_this).F27(), (_this).F27(), (_this).F27(), ((_this).F26()).Cardinality())).Cardinality())) != 0)
	}
}
func (_this *C5) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _1131_v0 D0
		_ = _1131_v0
		_1131_v0 = Companion_D0_.Create_DC0_()
		var _source17 D0 = _1131_v0
		_ = _source17
		if _source17.Is_DC0() {
			var _1132_v1 bool
			_ = _1132_v1
			_1132_v1 = true
			if (_this).Fm5(!(!(true) || (_1132_v1)), globalState) {
				var _1133_v2 _dafny.CodePoint
				_ = _1133_v2
				_1133_v2 = _dafny.CodePoint('s')
				var _1134_v3 _dafny.Sequence
				_ = _1134_v3
				_1134_v3 = _dafny.SeqOf(_1132_v1, _1132_v1, _1132_v1)
				var _1135_v4 _dafny.Map
				_ = _1135_v4
				_1135_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1133_v2, _1134_v3)
				var _rhs194 bool = ((_1134_v3).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1134_v3).Cardinality()))).Uint32()).(bool)) || (_1132_v1)
				_ = _rhs194
				var _rhs195 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1133_v2, _1134_v3)).Merge(_1135_v4)
				_ = _rhs195
				var _rhs196 _dafny.Int = p2
				_ = _rhs196
				var _lhs170 *GlobalState = globalState
				_ = _lhs170
				var _lhs171 *GlobalState = globalState
				_ = _lhs171
				_lhs170.F0 = _rhs194
				_1135_v4 = _rhs195
				_lhs171.F8 = _rhs196
				var _1136_v5 _dafny.Array
				_ = _1136_v5
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw146
				_1136_v5 = _nw146
				var _1137_v6 _dafny.Map
				_ = _1137_v6
				_1137_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v5, _dafny.CodePoint('q'))
				var _1138_v7 _dafny.Array
				_ = _1138_v7
				var _nwElement0_30 _dafny.Map = _1137_v6
				_ = _nwElement0_30
				var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(17))
				_ = _nw147
				(_nw147).ArraySet1(_nwElement0_30, 0)
				(_nw147).ArraySet1(_1137_v6, 1)
				(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v5, _1133_v2), 2)
				(_nw147).ArraySet1(_1137_v6, 3)
				(_nw147).ArraySet1((_1137_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v5, _1133_v2)), 4)
				(_nw147).ArraySet1((_1137_v6).Merge(_1137_v6), 5)
				(_nw147).ArraySet1((_1137_v6).Merge(_1137_v6), 6)
				(_nw147).ArraySet1(_1137_v6, 7)
				(_nw147).ArraySet1(_1137_v6, 8)
				(_nw147).ArraySet1(_1137_v6, 9)
				(_nw147).ArraySet1(_1137_v6, 10)
				(_nw147).ArraySet1((_1137_v6).Update(_1136_v5, _1133_v2), 11)
				(_nw147).ArraySet1(_1137_v6, 12)
				(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v5, _1133_v2), 13)
				(_nw147).ArraySet1((_1137_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v5, _1133_v2)), 14)
				(_nw147).ArraySet1((_1137_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v5, _1133_v2)), 15)
				(_nw147).ArraySet1(_1137_v6, 16)
				_1138_v7 = _nw147
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1138_v7), 0))
				_ = _index169
				(_1138_v7).ArraySet1(_1137_v6, (_index169).Int())
				var _1139_v8 _dafny.Map
				_ = _1139_v8
				_1139_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(636))
				var _1140_v9 _dafny.Map
				_ = _1140_v9
				_1140_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_v1, _1139_v8)
				var _1141_v10 _dafny.Map
				_ = _1141_v10
				_1141_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf(_dafny.IntOfInt64(988)))
				var _1142_v11 D0
				_ = _1142_v11
				_1142_v11 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(496))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1143_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1144_i0 _dafny.Int) _dafny.CodePoint {
						return _1143_v2
					}
				})(_1133_v2)))).Cardinality()), p2)), (func() bool {
					if _1132_v1 {
						return _1132_v1
					}
					return _1132_v1
				})(), (_1140_v9).Merge(_1140_v9), (_1141_v10).Cardinality(), _1133_v2)
				var _1145_v12 _dafny.Array
				_ = _1145_v12
				var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw148
				_1145_v12 = _nw148
				var _1146_v13 _dafny.Array
				_ = _1146_v13
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
				_ = _nw149
				_1146_v13 = _nw149
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_1145_v12), 0))
				_ = _index170
				(_1145_v12).ArraySet1(_1146_v13, (_index170).Int())
				var _pat_let_tv49 = _1132_v1
				_ = _pat_let_tv49
				var _pat_let_tv50 = p2
				_ = _pat_let_tv50
				var _pat_let_tv51 = p0
				_ = _pat_let_tv51
				var _pat_let_tv52 = _1132_v1
				_ = _pat_let_tv52
				var _pat_let_tv53 = _1139_v8
				_ = _pat_let_tv53
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1138_v7), 0))
				_ = _index171
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_1145_v12), 0))
				_ = _index172
				var _rhs197 bool = (Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality()))).Cmp(p2) != 0
				_ = _rhs197
				var _rhs198 _dafny.Map = _1137_v6
				_ = _rhs198
				var _rhs199 D0 = func(_pat_let18_0 D0) D0 {
					return func(_1147_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let19_0 _dafny.Int) D0 {
							return func(_1148_dt__update_hcf3_h0 _dafny.Int) D0 {
								return func(_pat_let20_0 _dafny.Int) D0 {
									return func(_1149_dt__update_hcf0_h0 _dafny.Int) D0 {
										return func(_pat_let21_0 _dafny.Map) D0 {
											return func(_1150_dt__update_hcf2_h0 _dafny.Map) D0 {
												return Companion_D0_.Create_DC1_(_1149_dt__update_hcf0_h0, (_1147_dt__update__tmp_h0).Dtor_cf1(), _1150_dt__update_hcf2_h0, _1148_dt__update_hcf3_h0, (_1147_dt__update__tmp_h0).Dtor_cf4())
											}(_pat_let21_0)
										}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv52, _pat_let_tv53))
									}(_pat_let20_0)
								}(_pat_let_tv51)
							}(_pat_let19_0)
						}((func() _dafny.Int {
							if _pat_let_tv49 {
								return _pat_let_tv50
							}
							return _dafny.IntOfInt64(-302)
						})())
					}(_pat_let18_0)
				}(_1142_v11)
				_ = _rhs199
				var _rhs200 _dafny.Array = _1146_v13
				_ = _rhs200
				var _lhs172 *GlobalState = globalState
				_ = _lhs172
				var _lhs173 _dafny.Array = _1138_v7
				_ = _lhs173
				var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1138_v7), 0))
				_ = _lhs174
				var _lhs175 _dafny.Array = _1145_v12
				_ = _lhs175
				var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_1145_v12), 0))
				_ = _lhs176
				_lhs172.F10 = _rhs197
				(_lhs173).ArraySet1(_rhs198, (_lhs174).Int())
				_1142_v11 = _rhs199
				(_lhs175).ArraySet1(_rhs200, (_lhs176).Int())
				var _1151_v14 _dafny.Sequence
				_ = _1151_v14
				_1151_v14 = _dafny.UnicodeSeqOfUtf8Bytes("hcvvsgggf")
				var _1152_v15 _dafny.MultiSet
				_ = _1152_v15
				_1152_v15 = _dafny.MultiSetOf((_this).F27(), p0)
				var _rhs201 _dafny.Sequence = _1151_v14
				_ = _rhs201
				var _rhs202 _dafny.Int = (Companion_Default___.SafeModuloInt((_1152_v15).Cardinality(), p0)).Times(p2)
				_ = _rhs202
				var _rhs203 bool = _1132_v1
				_ = _rhs203
				var _lhs177 *GlobalState = globalState
				_ = _lhs177
				var _lhs178 *GlobalState = globalState
				_ = _lhs178
				_1151_v14 = _rhs201
				_lhs177.F8 = _rhs202
				_lhs178.F0 = _rhs203
				var _1153_v16 _dafny.Int
				_ = _1153_v16
				var _1154_v17 D0
				_ = _1154_v17
				var _out50 _dafny.Int
				_ = _out50
				var _out51 D0
				_ = _out51
				_out50, _out51 = (_this).M3(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(882))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1155_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1156_i1 _dafny.Int) _dafny.CodePoint {
						return _1155_v2
					}
				})(_1133_v2))), !(!(_1132_v1) || (_1132_v1)), globalState)
				_1153_v16 = _out50
				_1154_v17 = _out51
				var _1157_v18 _dafny.Map
				_ = _1157_v18
				_1157_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_1132_v1)) || (true), _1136_v5)
				_1157_v18 = (_1157_v18).Update(_1132_v1, _1136_v5)
			} else {
				var _1158_v19 _dafny.Sequence
				_ = _1158_v19
				_1158_v19 = _dafny.UnicodeSeqOfUtf8Bytes("w")
				_1158_v19 = (func() _dafny.Sequence {
					if false {
						return _dafny.UnicodeSeqOfUtf8Bytes("tggp")
					}
					return _dafny.Companion_Sequence_.Concatenate(p1, p1)
				})()
				_1132_v1 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hkfl"), _1158_v19), Companion_Default___.Fm6(globalState))
				var _1159_v20 _dafny.CodePoint
				_ = _1159_v20
				_1159_v20 = _dafny.CodePoint('e')
				var _1160_v21 _dafny.Set
				_ = _1160_v21
				_1160_v21 = _dafny.SetOf(p1, _dafny.Companion_Sequence_.Update(_1158_v19, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1158_v19).Cardinality()))).Uint32(), _1159_v20))
				var _1161_v22 _dafny.Sequence
				_ = _1161_v22
				_1161_v22 = _dafny.SeqOf(_1158_v19)
				var _rhs204 _dafny.Set = (func() _dafny.Set {
					var _coll49 = _dafny.NewBuilder()
					_ = _coll49
					for _iter51 := _dafny.Iterate((_1161_v22).Elements()); ; {
						_compr_49, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1162_v23 _dafny.Sequence
						_1162_v23 = interface{}(_compr_49).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_1161_v22, _1162_v23) {
							_coll49.Add(_1162_v23)
						}
					}
					return _coll49.ToSet()
				}()).Union(_1160_v21)
				_ = _rhs204
				var _rhs205 _dafny.Int = ((_this).F27()).Times(_dafny.IntOfInt64(893))
				_ = _rhs205
				var _rhs206 bool = _1132_v1
				_ = _rhs206
				var _rhs207 _dafny.Int = ((_this).F27()).Minus((_dafny.Zero).Minus(p0))
				_ = _rhs207
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				_1160_v21 = _rhs204
				_lhs179.F11 = _rhs205
				_1132_v1 = _rhs206
				_lhs180.F11 = _rhs207
				var _1163_v24 _dafny.Sequence
				_ = _1163_v24
				_1163_v24 = _dafny.SeqOf(_1132_v1, true)
				var _1164_v25 _dafny.MultiSet
				_ = _1164_v25
				_1164_v25 = _dafny.MultiSetOf(_1132_v1, _1132_v1)
				var _1165_v26 _dafny.Map
				_ = _1165_v26
				_1165_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1164_v25).Contains(_1132_v1) {
						return (_1164_v25).Multiplicity(_1132_v1)
					}
					return _dafny.IntOfUint32((_1158_v19).Cardinality())
				})(), p0)
				var _1166_v27 _dafny.Map
				_ = _1166_v27
				_1166_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1163_v24).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1163_v24).Cardinality()))).Uint32()).(bool)), _1165_v26)
				var _1167_v28 _dafny.Map
				_ = _1167_v28
				_1167_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1158_v19, _1132_v1)
				var _1168_v29 D0
				_ = _1168_v29
				_1168_v29 = Companion_D0_.Create_DC1_(((_this).F26()).Cardinality(), false, _1166_v27, (_1167_v28).Cardinality(), _1159_v20)
				(globalState).F0 = (func() bool {
					if (_1168_v29).Dtor_cf1() {
						return !((_this).Fm5(_1132_v1, globalState))
					}
					return _1132_v1
				})()
				_1132_v1 = !((_1164_v25).IsSubsetOf(_1164_v25))
			}
			if _1132_v1 {
				_1131_v0 = _1131_v0
				var _1169_v30 _dafny.Array
				_ = _1169_v30
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
				_ = _nw150
				_1169_v30 = _nw150
				var _1170_v31 _dafny.Array
				_ = _1170_v31
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_20
				var _nw151 _dafny.Array
				_ = _nw151
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw151 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) bool = (func(_1171_v1 bool) func(_dafny.Int) bool {
						return func(_1172_i2 _dafny.Int) bool {
							return _1171_v1
						}
					})(_1132_v1)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw151 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw151).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw151).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_1170_v31 = _nw151
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1169_v30), 0))
				_ = _index173
				(_1169_v30).ArraySet1(_1170_v31, (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_1169_v30), 0))
				_ = _index174
				var _nwElement0_31 bool = _1132_v1
				_ = _nwElement0_31
				var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(7))
				_ = _nw152
				(_nw152).ArraySet1(_nwElement0_31, 0)
				(_nw152).ArraySet1(false, 1)
				(_nw152).ArraySet1((Companion_Default___.Fm7(p0, globalState)).Cmp((_this).F27()) == 0, 2)
				(_nw152).ArraySet1(_1132_v1, 3)
				(_nw152).ArraySet1(_1132_v1, 4)
				(_nw152).ArraySet1(_1132_v1, 5)
				(_nw152).ArraySet1(_1132_v1, 6)
				(_1169_v30).ArraySet1(_nw152, (_index174).Int())
				var _1173_v32 _dafny.Sequence
				_ = _1173_v32
				_1173_v32 = _dafny.SeqOf((_this).F26(), (_this).F26(), (_this).F26())
				var _1174_v35 _dafny.Sequence
				_ = _1174_v35
				_1174_v35 = _dafny.SeqOf(_dafny.IntOfInt64(738))
				var _1175_v36 _dafny.Array
				_ = _1175_v36
				var _nwElement0_32 _dafny.Map = (_this).F26()
				_ = _nwElement0_32
				var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(20))
				_ = _nw153
				(_nw153).ArraySet1(_nwElement0_32, 0)
				(_nw153).ArraySet1((_this).F26(), 1)
				(_nw153).ArraySet1((_1173_v32).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1173_v32).Cardinality()))).Uint32()).(_dafny.Map), 2)
				(_nw153).ArraySet1((_this).F26(), 3)
				(_nw153).ArraySet1((((_this).F26()).Update(p0, false)).Merge((_this).F26()), 4)
				(_nw153).ArraySet1((_this).F26(), 5)
				(_nw153).ArraySet1((_this).F26(), 6)
				(_nw153).ArraySet1(((_this).F26()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_1132_v1, !(_1132_v1))).Cardinality(), _1132_v1)), 7)
				(_nw153).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1132_v1), 8)
				(_nw153).ArraySet1((_this).F26(), 9)
				(_nw153).ArraySet1((_this).F26(), 10)
				(_nw153).ArraySet1((func() _dafny.Map {
					if _1132_v1 {
						return (_1173_v32).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1173_v32).Cardinality()))).Uint32()).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1132_v1)
				})(), 11)
				(_nw153).ArraySet1(((_this).F26()).Merge((_this).F26()), 12)
				(_nw153).ArraySet1((_this).F26(), 13)
				(_nw153).ArraySet1((_this).F26(), 14)
				(_nw153).ArraySet1(func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(243), _dafny.IntOfInt64(987))); ; {
						_compr_50, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1176_v33 _dafny.Int
						_1176_v33 = interface{}(_compr_50).(_dafny.Int)
						if ((_dafny.IntOfInt64(243)).Cmp(_1176_v33) <= 0) && ((_1176_v33).Cmp(_dafny.IntOfInt64(987)) < 0) {
							_coll50.Add((_1176_v33).Minus(_dafny.IntOfInt64(-492)), _1132_v1)
						}
					}
					return _coll50.ToMap()
				}(), 15)
				(_nw153).ArraySet1((_this).F26(), 16)
				(_nw153).ArraySet1((func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter53 := _dafny.Iterate((_1174_v35).Elements()); ; {
						_compr_51, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _1177_v34 _dafny.Int
						_1177_v34 = interface{}(_compr_51).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1174_v35, _1177_v34) {
							_coll51.Add((_1177_v34).Plus((_this).F27()), _1132_v1)
						}
					}
					return _coll51.ToMap()
				}()).Merge(Companion_Default___.Fm8(p1, _1132_v1, p2, globalState)), 17)
				(_nw153).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(653), _1132_v1), 18)
				(_nw153).ArraySet1((_this).F26(), 19)
				_1175_v36 = _nw153
				_1175_v36 = _1175_v36
				var _1178_v37 _dafny.Set
				_ = _1178_v37
				_1178_v37 = _dafny.SetOf(_1132_v1)
				var _1179_v38 _dafny.MultiSet
				_ = _1179_v38
				_1179_v38 = _dafny.MultiSetOf(_1178_v37)
				(globalState).F0 = (_dafny.Companion_Sequence_.Equal(_1174_v35, _1174_v35)) == ((_1179_v38).IsProperSubsetOf(_1179_v38))
				var _1180_v39 _dafny.Array
				_ = _1180_v39
				var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw154
				_1180_v39 = _nw154
				var _1181_v40 _dafny.MultiSet
				_ = _1181_v40
				_1181_v40 = _dafny.MultiSetOf(_dafny.IntOfUint32((p1).Cardinality()))
				var _1182_v41 _dafny.Array
				_ = _1182_v41
				var _nwElement0_33 _dafny.Int = _dafny.IntOfInt64(740)
				_ = _nwElement0_33
				var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(8))
				_ = _nw155
				(_nw155).ArraySet1(_nwElement0_33, 0)
				(_nw155).ArraySet1(p0, 1)
				(_nw155).ArraySet1(p2, 2)
				(_nw155).ArraySet1((_this).F27(), 3)
				(_nw155).ArraySet1(p0, 4)
				(_nw155).ArraySet1(_dafny.IntOfInt64(23), 5)
				(_nw155).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 6)
				(_nw155).ArraySet1((_1181_v40).Cardinality(), 7)
				_1182_v41 = _nw155
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1180_v39), 0))
				_ = _index175
				(_1180_v39).ArraySet1(_1182_v41, (_index175).Int())
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1180_v39), 0))
				_ = _index176
				(_1180_v39).ArraySet1(_1182_v41, (_index176).Int())
			} else {
				var _1183_v42 _dafny.Array
				_ = _1183_v42
				var _len0_21 _dafny.Int = _dafny.One
				_ = _len0_21
				var _nw156 _dafny.Array
				_ = _nw156
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw156 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = (func(_1184_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1185_i3 _dafny.Int) _dafny.Int {
							return (_1185_i3).Plus(_1184_p2)
						}
					})(p2)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw156 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw156).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw156).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_1183_v42 = _nw156
				_1183_v42 = _1183_v42
				var _1186_v43 _dafny.Set
				_ = _1186_v43
				_1186_v43 = _dafny.SetOf(_1132_v1)
				var _1187_v44 _dafny.Map
				_ = _1187_v44
				_1187_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(-518))
				var _1188_v45 _dafny.CodePoint
				_ = _1188_v45
				_1188_v45 = _dafny.CodePoint('l')
				var _1189_v46 D0
				_ = _1189_v46
				_1189_v46 = Companion_D0_.Create_DC1_(((_1186_v43).Cardinality()).Plus((_this).F27()), _1132_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1132_v1), _1187_v44), (_this).F27(), (func() _dafny.CodePoint {
					if _1132_v1 {
						return _1188_v45
					}
					return _1188_v45
				})())
				var _1190_v47 _dafny.Sequence
				_ = _1190_v47
				_1190_v47 = _dafny.SeqOf(Companion_Default___.Fm7((_this).F27(), globalState), _dafny.IntOfInt64(-995), _dafny.IntOfUint32((p1).Cardinality()))
				var _1191_v48 _dafny.MultiSet
				_ = _1191_v48
				_1191_v48 = _dafny.MultiSetOf(p0, (_dafny.MultiSetOf((_this).F26())).Cardinality(), p2, _dafny.IntOfUint32((_1190_v47).Cardinality()), (_this).F27())
				var _rhs208 D0 = _1189_v46
				_ = _rhs208
				var _rhs209 _dafny.MultiSet = (_dafny.MultiSetOf(p0)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(p2)))
				_ = _rhs209
				var _rhs210 _dafny.Int = (p2).Plus(p2)
				_ = _rhs210
				var _lhs181 *GlobalState = globalState
				_ = _lhs181
				_1189_v46 = _rhs208
				_1191_v48 = _rhs209
				_lhs181.F8 = _rhs210
				_1188_v45 = _1188_v45
				(globalState).F0 = _1132_v1
				var _1192_v49 _dafny.Sequence
				_ = _1192_v49
				_1192_v49 = _dafny.UnicodeSeqOfUtf8Bytes("lx")
				var _1193_v50 D1
				_ = _1193_v50
				_1193_v50 = Companion_D1_.Create_DC3_(p1)
				_1192_v49 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p1, (_1193_v50).Dtor_cf6()), p1)
			}
			var _1194_v51 _dafny.Map
			_ = _1194_v51
			_1194_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0))
			var _1195_v52 _dafny.Map
			_ = _1195_v52
			_1195_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_v1, _1132_v1)
			var _1196_v53 _dafny.CodePoint
			_ = _1196_v53
			_1196_v53 = _dafny.CodePoint('k')
			var _1197_v54 D0
			_ = _1197_v54
			_1197_v54 = Companion_D0_.Create_DC1_(((_this).F27()).Plus((_this).F27()), _1132_v1, (_1194_v51).Merge(_1194_v51), Companion_Default___.Fm7((_1195_v52).Cardinality(), globalState), _1196_v53)
			var _source18 D0 = _1197_v54
			_ = _source18
			if _source18.Is_DC0() {
				(globalState).F0 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-337))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1198_v53 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1199_i4 _dafny.Int) _dafny.CodePoint {
						return _1198_v53
					}
				})(_1196_v53))), _dafny.UnicodeSeqOfUtf8Bytes("qewbkgcot")), _1196_v53)
				(globalState).F11 = Companion_Default___.Fm7(_dafny.IntOfInt64(93), globalState)
				var _1200_v55 _dafny.Sequence
				_ = _1200_v55
				_1200_v55 = _dafny.SeqOf(_1132_v1)
				var _1201_v56 _dafny.Sequence
				_ = _1201_v56
				_1201_v56 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()), p2)
				var _1202_v57 _dafny.Sequence
				_ = _1202_v57
				_1202_v57 = _dafny.SeqOf(_1200_v55, _dafny.Companion_Sequence_.Concatenate(_1200_v55, (_this).Fm4(_1201_v56, p1, _dafny.SetOf(!(_1132_v1)), p2, globalState)))
				_1202_v57 = _1202_v57
				var _1203_v58 _dafny.Array
				_ = _1203_v58
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_22
				var _nw157 _dafny.Array
				_ = _nw157
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw157 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) bool = (func(_1204_v1 bool) func(_dafny.Int) bool {
						return func(_1205_i5 _dafny.Int) bool {
							return (func() bool {
								if _1204_v1 {
									return _1204_v1
								}
								return _1204_v1
							})()
						}
					})(_1132_v1)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw157 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw157).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw157).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_1203_v58 = _nw157
				_1203_v58 = _1203_v58
			} else {
				var _1206___mcc_h5 _dafny.Int = _source18.Get_().(D0_DC1).Cf0
				_ = _1206___mcc_h5
				var _1207___mcc_h6 bool = _source18.Get_().(D0_DC1).Cf1
				_ = _1207___mcc_h6
				var _1208___mcc_h7 _dafny.Map = _source18.Get_().(D0_DC1).Cf2
				_ = _1208___mcc_h7
				var _1209___mcc_h8 _dafny.Int = _source18.Get_().(D0_DC1).Cf3
				_ = _1209___mcc_h8
				var _1210___mcc_h9 _dafny.CodePoint = _source18.Get_().(D0_DC1).Cf4
				_ = _1210___mcc_h9
				var _1211_cf4 _dafny.CodePoint = _1210___mcc_h9
				_ = _1211_cf4
				var _1212_cf3 _dafny.Int = _1209___mcc_h8
				_ = _1212_cf3
				var _1213_cf2 _dafny.Map = _1208___mcc_h7
				_ = _1213_cf2
				var _1214_cf1 bool = _1207___mcc_h6
				_ = _1214_cf1
				var _1215_cf0 _dafny.Int = _1206___mcc_h5
				_ = _1215_cf0
				var _1216_v59 _dafny.Int
				_ = _1216_v59
				var _1217_v60 D0
				_ = _1217_v60
				var _out52 _dafny.Int
				_ = _out52
				var _out53 D0
				_ = _out53
				_out52, _out53 = (_this).M3(_dafny.UnicodeSeqOfUtf8Bytes("bubrnhav"), _1132_v1, globalState)
				_1216_v59 = _out52
				_1217_v60 = _out53
				var _1218_v61 _dafny.Array
				_ = _1218_v61
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_23
				var _nw158 _dafny.Array
				_ = _nw158
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw158 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_1219_cf0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1220_i6 _dafny.Int) _dafny.Int {
							return (_1220_i6).Minus(_1219_cf0)
						}
					})(_1215_cf0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw158 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw158).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw158).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1218_v61 = _nw158
				_1218_v61 = _1218_v61
				var _1221_v62 _dafny.Sequence
				_ = _1221_v62
				_1221_v62 = _dafny.UnicodeSeqOfUtf8Bytes("kgfnnee")
				_1221_v62 = _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(Companion_Default___.Fm9(globalState))).Cardinality(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1211_cf4)
				var _1222_v63 _dafny.Map
				_ = _1222_v63
				_1222_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1214_cf1, _1211_cf4)
				var _1223_v64 _dafny.Map
				_ = _1223_v64
				_1223_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _1222_v63)
				var _1224_v65 _dafny.Array
				_ = _1224_v65
				var _nwElement0_34 _dafny.Map = (_1222_v63).Merge(_1222_v63)
				_ = _nwElement0_34
				var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(8))
				_ = _nw159
				(_nw159).ArraySet1(_nwElement0_34, 0)
				(_nw159).ArraySet1((func() _dafny.Map {
					if (_1223_v64).Contains(_1215_cf0) {
						return (_1223_v64).Get(_1215_cf0).(_dafny.Map)
					}
					return _1222_v63
				})(), 1)
				(_nw159).ArraySet1(_1222_v63, 2)
				(_nw159).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1214_cf1, _1196_v53), 3)
				(_nw159).ArraySet1((_1222_v63).Update(_1214_cf1, _1196_v53), 4)
				(_nw159).ArraySet1(_1222_v63, 5)
				(_nw159).ArraySet1((_1222_v63).Merge(_1222_v63), 6)
				(_nw159).ArraySet1((_1222_v63).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1214_cf1, _1211_cf4)), 7)
				_1224_v65 = _nw159
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1224_v65), 0))
				_ = _index177
				(_1224_v65).ArraySet1((_1222_v63).Update(_1132_v1, _dafny.CodePoint('e')), (_index177).Int())
				var _1225_v66 _dafny.Array
				_ = _1225_v66
				var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw160
				_1225_v66 = _nw160
				var _1226_v67 _dafny.Sequence
				_ = _1226_v67
				_1226_v67 = _dafny.SeqOf(_1132_v1)
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_1225_v66), 0))
				_ = _index178
				(_1225_v66).ArraySet1(_1226_v67, (_index178).Int())
				var _1227_v68 _dafny.Sequence
				_ = _1227_v68
				_1227_v68 = _dafny.SeqOf((_this).F27())
				var _1228_v69 _dafny.MultiSet
				_ = _1228_v69
				_1228_v69 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1227_v68).Cardinality()), _1215_cf0)
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1224_v65), 0))
				_ = _index179
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_1225_v66), 0))
				_ = _index180
				var _rhs211 bool = _1132_v1
				_ = _rhs211
				var _rhs212 bool = _1132_v1
				_ = _rhs212
				var _rhs213 _dafny.Map = (_1222_v63).Merge(_1222_v63)
				_ = _rhs213
				var _rhs214 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1132_v1), _dafny.SeqOf((Companion_D0_.Create_DC1_((func() _dafny.Int {
					if (_1228_v69).Contains(_1212_cf3) {
						return (_1228_v69).Multiplicity(_1212_cf3)
					}
					return (Companion_Default___.Fm10(_dafny.IntOfUint32((_1221_v62).Cardinality()), _1132_v1, _1212_cf3, globalState)).Cardinality()
				})(), false, _1213_cf2, Companion_Default___.Fm7(p0, globalState), _1196_v53)).Dtor_cf1(), _1214_cf1, _1132_v1))
				_ = _rhs214
				var _lhs182 *GlobalState = globalState
				_ = _lhs182
				var _lhs183 _dafny.Array = _1224_v65
				_ = _lhs183
				var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1224_v65), 0))
				_ = _lhs184
				var _lhs185 _dafny.Array = _1225_v66
				_ = _lhs185
				var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_1225_v66), 0))
				_ = _lhs186
				_lhs182.F10 = _rhs211
				_1214_cf1 = _rhs212
				(_lhs183).ArraySet1(_rhs213, (_lhs184).Int())
				(_lhs185).ArraySet1(_rhs214, (_lhs186).Int())
			}
			_1196_v53 = _1196_v53
		} else {
			var _1229___mcc_h0 _dafny.Int = _source17.Get_().(D0_DC1).Cf0
			_ = _1229___mcc_h0
			var _1230___mcc_h1 bool = _source17.Get_().(D0_DC1).Cf1
			_ = _1230___mcc_h1
			var _1231___mcc_h2 _dafny.Map = _source17.Get_().(D0_DC1).Cf2
			_ = _1231___mcc_h2
			var _1232___mcc_h3 _dafny.Int = _source17.Get_().(D0_DC1).Cf3
			_ = _1232___mcc_h3
			var _1233___mcc_h4 _dafny.CodePoint = _source17.Get_().(D0_DC1).Cf4
			_ = _1233___mcc_h4
			var _1234_cf4 _dafny.CodePoint = _1233___mcc_h4
			_ = _1234_cf4
			var _1235_cf3 _dafny.Int = _1232___mcc_h3
			_ = _1235_cf3
			var _1236_cf2 _dafny.Map = _1231___mcc_h2
			_ = _1236_cf2
			var _1237_cf1 bool = _1230___mcc_h1
			_ = _1237_cf1
			var _1238_cf0 _dafny.Int = _1229___mcc_h0
			_ = _1238_cf0
			var _1239_v70 _dafny.Map
			_ = _1239_v70
			_1239_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_Default___.Fm11(_1238_cf0, _1234_cf4, _1237_cf1, globalState)).Cardinality())
			var _1240_v71 _dafny.Sequence
			_ = _1240_v71
			_1240_v71 = _dafny.SeqOf(_1238_cf0, _dafny.IntOfInt64(-250), (func() _dafny.Int {
				if (_1239_v70).Contains((_this).F27()) {
					return (_1239_v70).Get((_this).F27()).(_dafny.Int)
				}
				return (_this).F27()
			})(), (_dafny.Zero).Minus((_this).F27()), p2)
			var _1241_v72 _dafny.Map
			_ = _1241_v72
			_1241_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_cf1, _1238_cf0)
			var _1242_v73 _dafny.Set
			_ = _1242_v73
			_1242_v73 = _dafny.SetOf((_this).F27(), (func() _dafny.Int {
				if (_1241_v72).Contains(_1237_cf1) {
					return (_1241_v72).Get(_1237_cf1).(_dafny.Int)
				}
				return _1235_cf3
			})())
			var _1243_v74 _dafny.Map
			_ = _1243_v74
			_1243_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_cf1, _1242_v73)
			var _1244_v75 D0
			_ = _1244_v75
			_1244_v75 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_1240_v71).Cardinality()), _1237_cf1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_cf1, _1239_v70), p2, _1234_cf4)
			var _1245_v76 _dafny.Map
			_ = _1245_v76
			_1245_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v73, _1238_cf0)
			var _1246_v77 _dafny.Array
			_ = _1246_v77
			var _nwElement0_35 _dafny.Int = _dafny.IntOfUint32((_1240_v71).Cardinality())
			_ = _nwElement0_35
			var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(25))
			_ = _nw161
			(_nw161).ArraySet1(_nwElement0_35, 0)
			(_nw161).ArraySet1((_dafny.Zero).Minus((_1243_v74).Cardinality()), 1)
			(_nw161).ArraySet1(_dafny.IntOfInt64(-82), 2)
			(_nw161).ArraySet1(_1235_cf3, 3)
			(_nw161).ArraySet1(_dafny.IntOfInt64(191), 4)
			(_nw161).ArraySet1(_1235_cf3, 5)
			(_nw161).ArraySet1(_1235_cf3, 6)
			(_nw161).ArraySet1(p2, 7)
			(_nw161).ArraySet1(p0, 8)
			(_nw161).ArraySet1(p2, 9)
			(_nw161).ArraySet1((_1240_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.IntOfUint32((_1240_v71).Cardinality()))).Uint32()).(_dafny.Int), 10)
			(_nw161).ArraySet1(_1238_cf0, 11)
			(_nw161).ArraySet1(p0, 12)
			(_nw161).ArraySet1(p0, 13)
			(_nw161).ArraySet1(_1238_cf0, 14)
			(_nw161).ArraySet1((_this).F27(), 15)
			(_nw161).ArraySet1((_this).F27(), 16)
			(_nw161).ArraySet1(_dafny.IntOfInt64(65), 17)
			(_nw161).ArraySet1(_1235_cf3, 18)
			(_nw161).ArraySet1(_dafny.IntOfInt64(750), 19)
			(_nw161).ArraySet1((_1244_v75).Dtor_cf3(), 20)
			(_nw161).ArraySet1((_this).F27(), 21)
			(_nw161).ArraySet1((func() _dafny.Int {
				if (_1245_v76).Contains(_1242_v73) {
					return (_1245_v76).Get(_1242_v73).(_dafny.Int)
				}
				return p2
			})(), 22)
			(_nw161).ArraySet1((_1241_v72).Cardinality(), 23)
			(_nw161).ArraySet1(p0, 24)
			_1246_v77 = _nw161
			var _1247_v78 D2
			_ = _1247_v78
			_1247_v78 = Companion_D2_.Create_DC5_(_1246_v77)
			var _1248_v79 _dafny.Array
			_ = _1248_v79
			var _nwElement0_36 _dafny.Array = _1246_v77
			_ = _nwElement0_36
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(14))
			_ = _nw162
			(_nw162).ArraySet1(_nwElement0_36, 0)
			(_nw162).ArraySet1(_1246_v77, 1)
			(_nw162).ArraySet1(_1246_v77, 2)
			(_nw162).ArraySet1((func() _dafny.Array {
				if _1237_cf1 {
					return _1246_v77
				}
				return _1246_v77
			})(), 3)
			(_nw162).ArraySet1(_1246_v77, 4)
			(_nw162).ArraySet1(_1246_v77, 5)
			(_nw162).ArraySet1(_1246_v77, 6)
			(_nw162).ArraySet1(_1246_v77, 7)
			(_nw162).ArraySet1(_1246_v77, 8)
			(_nw162).ArraySet1(_1246_v77, 9)
			(_nw162).ArraySet1((func() _dafny.Array {
				if _1237_cf1 {
					return _1246_v77
				}
				return (_1247_v78).Dtor_cf8()
			})(), 10)
			(_nw162).ArraySet1(_1246_v77, 11)
			(_nw162).ArraySet1(_1246_v77, 12)
			(_nw162).ArraySet1(_1246_v77, 13)
			_1248_v79 = _nw162
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_1248_v79), 0))
			_ = _index181
			(_1248_v79).ArraySet1(_1246_v77, (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_1248_v79), 0))
			_ = _index182
			var _rhs215 _dafny.Array = _1246_v77
			_ = _rhs215
			var _rhs216 _dafny.Int = (_dafny.IntOfUint32((Companion_Default___.Fm6(globalState)).Cardinality())).Plus((_dafny.Zero).Minus(_1235_cf3))
			_ = _rhs216
			var _rhs217 _dafny.CodePoint = _1234_cf4
			_ = _rhs217
			var _rhs218 _dafny.Int = (_this).F27()
			_ = _rhs218
			var _lhs187 _dafny.Array = _1248_v79
			_ = _lhs187
			var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_1248_v79), 0))
			_ = _lhs188
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			(_lhs187).ArraySet1(_rhs215, (_lhs188).Int())
			_lhs189.F16 = _rhs216
			_1234_cf4 = _rhs217
			_lhs190.F16 = _rhs218
			var _1249_v80 _dafny.Array
			_ = _1249_v80
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_24
			var _nw163 _dafny.Array
			_ = _nw163
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw163 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) bool = (func(_1250_v75 D0) func(_dafny.Int) bool {
					return func(_1251_i7 _dafny.Int) bool {
						return (_1250_v75).Dtor_cf1()
					}
				})(_1244_v75)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw163 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw163).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw163).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1249_v80 = _nw163
			_1249_v80 = _1249_v80
			(globalState).F10 = false
			_1237_cf1 = false
		}
		var _hi2 _dafny.Int = p0
		_ = _hi2
		for _1252_i8 := (_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())); _1252_i8.Cmp(_hi2) < 0; _1252_i8 = _1252_i8.Plus(_dafny.One) {
			var _1253_v81 _dafny.Sequence
			_ = _1253_v81
			_1253_v81 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(617), Companion_Default___.Fm7(_1252_i8, globalState)))
			var _1254_v82 _dafny.Map
			_ = _1254_v82
			_1254_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Plus((_this).F27()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0, (_dafny.Zero).Minus((_this).F27()), _1252_i8, (_this).F27()), _1253_v81))
			_1253_v81 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1254_v82).Contains((_dafny.Zero).Minus(p2)) {
					return (_1254_v82).Get((_dafny.Zero).Minus(p2)).(_dafny.Sequence)
				}
				return _1253_v81
			})(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1254_v82).Contains((_dafny.Zero).Minus(p2)) {
					return (_1254_v82).Get((_dafny.Zero).Minus(p2)).(_dafny.Sequence)
				}
				return _1253_v81
			})()).Cardinality()))).Uint32(), (p0).Minus(p2))
			var _1255_v83 _dafny.Array
			_ = _1255_v83
			var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
			_ = _nw164
			_1255_v83 = _nw164
			_1255_v83 = _1255_v83
			var _1256_v84 bool
			_ = _1256_v84
			_1256_v84 = true
			if ((_this).Fm5(_1256_v84, globalState)) == (_1256_v84) {
				(globalState).F10 = (_1256_v84) || (_1256_v84)
				var _1257_v85 _dafny.Array
				_ = _1257_v85
				var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw165
				_1257_v85 = _nw165
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw166
				_1257_v85 = _nw166
				(globalState).F0 = _1256_v84
				var _1258_v86 _dafny.Array
				_ = _1258_v86
				var _nw167 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(16))
				_ = _nw167
				_1258_v86 = _nw167
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1258_v86), 0))
				_ = _index183
				(_1258_v86).ArraySet1(Companion_D2_.Create_DC5_(_1257_v85), (_index183).Int())
				var _pat_let_tv54 = _1257_v85
				_ = _pat_let_tv54
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1258_v86), 0))
				_ = _index184
				(_1258_v86).ArraySet1(func(_pat_let22_0 D2) D2 {
					return func(_1259_dt__update__tmp_h1 D2) D2 {
						return func(_pat_let23_0 _dafny.Array) D2 {
							return func(_1260_dt__update_hcf8_h0 _dafny.Array) D2 {
								return Companion_D2_.Create_DC5_(_1260_dt__update_hcf8_h0)
							}(_pat_let23_0)
						}(_pat_let_tv54)
					}(_pat_let22_0)
				}(Companion_D2_.Create_DC5_(_1257_v85)), (_index184).Int())
				(globalState).F11 = Companion_Default___.Fm7((_this).F27(), globalState)
			} else {
				var _1261_v87 _dafny.CodePoint
				_ = _1261_v87
				_1261_v87 = _dafny.CodePoint('h')
				_1261_v87 = Companion_Default___.Fm12((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1256_v84), (_1253_v81).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1253_v81).Cardinality()))).Uint32()).(_dafny.Int))).Contains(_1256_v84), _dafny.IntOfInt64(287), globalState)
				var _1262_v88 _dafny.Int
				_ = _1262_v88
				var _1263_v89 D0
				_ = _1263_v89
				var _out54 _dafny.Int
				_ = _out54
				var _out55 D0
				_ = _out55
				_out54, _out55 = (_this).M3(p1, _1256_v84, globalState)
				_1262_v88 = _out54
				_1263_v89 = _out55
				var _1264_v90 _dafny.Map
				_ = _1264_v90
				_1264_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6(globalState), _dafny.IntOfUint32((p1).Cardinality()))
				_1264_v90 = (_1264_v90).Update(p1, p2)
				var _1265_v91 _dafny.Array
				_ = _1265_v91
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_25
				var _nw168 _dafny.Array
				_ = _nw168
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw168 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) bool = (func(_1266_p1 _dafny.Sequence, _1267_p2 _dafny.Int) func(_dafny.Int) bool {
						return func(_1268_i9 _dafny.Int) bool {
							return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1266_p1, (Companion_Default___.SafeIndex(_1267_p2, _dafny.IntOfUint32((_1266_p1).Cardinality()))).Uint32(), _dafny.CodePoint('u'))).Cardinality())).Cmp(_1267_p2) == 0
						}
					})(p1, p2)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw168 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw168).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw168).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1265_v91 = _nw168
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1265_v91), 0))
				_ = _index185
				(_1265_v91).ArraySet1(_1256_v84, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_1265_v91), 0))
				_ = _index186
				(_1265_v91).ArraySet1(_1256_v84, (_index186).Int())
				var _1269_v92 _dafny.Sequence
				_ = _1269_v92
				_1269_v92 = _dafny.UnicodeSeqOfUtf8Bytes("a")
				var _1270_v93 _dafny.Map
				_ = _1270_v93
				_1270_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1256_v84, _1256_v84)
				var _1271_v94 D2
				_ = _1271_v94
				_1271_v94 = Companion_D2_.Create_DC7_(_1253_v81, _1256_v84, _1256_v84, (_1270_v93).Cardinality())
				var _1272_v95 _dafny.MultiSet
				_ = _1272_v95
				_1272_v95 = _dafny.MultiSetOf(_1271_v94)
				var _1273_v96 _dafny.Map
				_ = _1273_v96
				_1273_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1252_i8, (_this).F27())
				var _1274_v97 D3
				_ = _1274_v97
				_1274_v97 = Companion_D3_.Create_DC10_(_1273_v96)
				var _1275_v98 _dafny.Map
				_ = _1275_v98
				_1275_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1256_v84, (_1274_v97).Dtor_cf18())
				var _1276_v99 D0
				_ = _1276_v99
				_1276_v99 = Companion_D0_.Create_DC1_(p0, _1256_v84, _1275_v98, p0, _1261_v87)
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1265_v91), 0))
				_ = _index187
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_1265_v91), 0))
				_ = _index188
				var _rhs219 bool = (_1272_v95).IsProperSubsetOf(_1272_v95)
				_ = _rhs219
				var _rhs220 bool = (_1276_v99).Dtor_cf1()
				_ = _rhs220
				var _rhs221 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _rhs221
				var _rhs222 _dafny.CodePoint = (p1).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs222
				var _rhs223 _dafny.Sequence = p1
				_ = _rhs223
				var _lhs191 _dafny.Array = _1265_v91
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1265_v91), 0))
				_ = _lhs192
				var _lhs193 _dafny.Array = _1265_v91
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_1265_v91), 0))
				_ = _lhs194
				var _lhs195 *GlobalState = globalState
				_ = _lhs195
				(_lhs191).ArraySet1(_rhs219, (_lhs192).Int())
				(_lhs193).ArraySet1(_rhs220, (_lhs194).Int())
				_lhs195.F11 = _rhs221
				_1261_v87 = _rhs222
				_1269_v92 = _rhs223
				var _1277_v100 _dafny.Map
				_ = _1277_v100
				_1277_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1261_v87, _1262_v88)
				_1277_v100 = (_1277_v100).Update(Companion_Default___.Fm12((_1265_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1265_v91), 0))).Int()).(bool), _dafny.IntOfUint32((_1269_v92).Cardinality()), globalState), _1252_i8)
			}
			var _1278_v101 _dafny.Map
			_ = _1278_v101
			_1278_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1252_i8, p0)
			var _1279_v102 D3
			_ = _1279_v102
			_1279_v102 = Companion_D3_.Create_DC10_(_1278_v101)
			var _source19 D3 = _1279_v102
			_ = _source19
			if _source19.Is_DC11() {
				var _1280___mcc_h10 bool = _source19.Get_().(D3_DC11).Cf19
				_ = _1280___mcc_h10
				var _1281___mcc_h11 bool = _source19.Get_().(D3_DC11).Cf20
				_ = _1281___mcc_h11
				var _1282_cf20 bool = _1281___mcc_h11
				_ = _1282_cf20
				var _1283_cf19 bool = _1280___mcc_h10
				_ = _1283_cf19
				_1283_cf19 = _1256_v84
				(globalState).F0 = _1256_v84
				(globalState).F0 = !(_1283_cf19)
				var _1284_v103 _dafny.Sequence
				_ = _1284_v103
				_1284_v103 = _dafny.UnicodeSeqOfUtf8Bytes("bkhdpyee")
				var _1285_v104 _dafny.Sequence
				_ = _1285_v104
				_1285_v104 = _dafny.SeqOf(p1)
				var _1286_v105 _dafny.Map
				_ = _1286_v105
				_1286_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1283_cf19, _1278_v101)
				var _1287_v106 D0
				_ = _1287_v106
				_1287_v106 = Companion_D0_.Create_DC1_(_1252_i8, _1283_cf19, _1286_v105, (_dafny.Zero).Minus(p2), _dafny.CodePoint('e'))
				_1284_v103 = _dafny.Companion_Sequence_.Concatenate((_1285_v104).Select((Companion_Default___.SafeIndex((_1287_v106).Dtor_cf0(), _dafny.IntOfUint32((_1285_v104).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-122))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}(func(_1288_i10 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				})))
			} else {
				var _1289___mcc_h12 _dafny.Map = _source19.Get_().(D3_DC10).Cf18
				_ = _1289___mcc_h12
				var _1290_cf18 _dafny.Map = _1289___mcc_h12
				_ = _1290_cf18
				var _1291_v107 _dafny.Sequence
				_ = _1291_v107
				_1291_v107 = _dafny.SeqOf(_1256_v84)
				var _1292_v108 D3
				_ = _1292_v108
				_1292_v108 = Companion_D3_.Create_DC11_(_1256_v84, (_1291_v107).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_1291_v107).Cardinality()))).Uint32()).(bool))
				(globalState).F10 = !((_1292_v108).Dtor_cf19())
				var _1293_v109 _dafny.CodePoint
				_ = _1293_v109
				_1293_v109 = _dafny.CodePoint('h')
				var _1294_v110 _dafny.Map
				_ = _1294_v110
				_1294_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.SeqOf(_1293_v109))
				var _1295_v111 _dafny.Array
				_ = _1295_v111
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw169
				_1295_v111 = _nw169
				var _1296_v112 _dafny.MultiSet
				_ = _1296_v112
				_1296_v112 = _dafny.MultiSetOf(_1295_v111, _1295_v111)
				var _1297_v113 _dafny.Map
				_ = _1297_v113
				_1297_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1294_v110).Contains(_1256_v84) {
						return (_1294_v110).Get(_1256_v84).(_dafny.Sequence)
					}
					return p1
				})()).Cardinality()), (_dafny.SetOf((func() _dafny.Int {
					if (_1296_v112).Contains(_1295_v111) {
						return (_1296_v112).Multiplicity(_1295_v111)
					}
					return _1252_i8
				})())).Cardinality())), true)
				_1297_v113 = (_this).F26()
				_1295_v111 = _1295_v111
				var _1298_v114 _dafny.Set
				_ = _1298_v114
				_1298_v114 = _dafny.SetOf(_1293_v109, _1293_v109)
				var _1299_v115 _dafny.Sequence
				_ = _1299_v115
				_1299_v115 = _dafny.SeqOf(_1297_v113, (_1297_v113).Merge(_1297_v113), (_this).F26(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _1256_v84)).Update((_1298_v114).Cardinality(), _1256_v84))
				var _1300_v116 _dafny.MultiSet
				_ = _1300_v116
				_1300_v116 = _dafny.MultiSetOf(Companion_Default___.Fm7((_this).F27(), globalState), _dafny.IntOfInt64(122), _dafny.IntOfUint32((_1291_v107).Cardinality()), _dafny.IntOfInt64(107))
				var _1301_v117 _dafny.Map
				_ = _1301_v117
				_1301_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1300_v116).IsProperSubsetOf(_1300_v116), !(!(_1256_v84)))
				var _rhs224 bool = (p2).Cmp(p0) < 0
				_ = _rhs224
				var _rhs225 bool = _1256_v84
				_ = _rhs225
				var _rhs226 _dafny.Int = _1252_i8
				_ = _rhs226
				var _rhs227 _dafny.Sequence = _1299_v115
				_ = _rhs227
				var _rhs228 _dafny.Map = _1301_v117
				_ = _rhs228
				var _lhs196 *GlobalState = globalState
				_ = _lhs196
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				var _lhs198 *GlobalState = globalState
				_ = _lhs198
				_lhs196.F0 = _rhs224
				_lhs197.F0 = _rhs225
				_lhs198.F8 = _rhs226
				_1299_v115 = _rhs227
				_1301_v117 = _rhs228
			}
		}
		var _1302_v119 _dafny.Map
		_ = _1302_v119
		_1302_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p2)
		var _1303_v120 bool
		_ = _1303_v120
		_1303_v120 = false
		var _1304_v121 _dafny.Map
		_ = _1304_v121
		_1304_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1303_v120, p2)
		var _1305_v122 _dafny.MultiSet
		_ = _1305_v122
		_1305_v122 = _dafny.MultiSetOf((_1304_v121).Cardinality(), (_this).F27())
		var _1306_v123 _dafny.Map
		_ = _1306_v123
		_1306_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _1305_v122)
		var _1307_v124 _dafny.Map
		_ = _1307_v124
		_1307_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll52 = _dafny.NewMapBuilder()
			_ = _coll52
			for _iter54 := _dafny.Iterate(((_1302_v119).Update((_this).F27(), p0)).Keys().Elements()); ; {
				_compr_52, _ok54 := _iter54()
				if !_ok54 {
					break
				}
				var _1308_v118 _dafny.Int
				_1308_v118 = interface{}(_compr_52).(_dafny.Int)
				if ((_1302_v119).Update((_this).F27(), p0)).Contains(_1308_v118) {
					_coll52.Add((_1308_v118).Times((_this).F27()), _dafny.MultiSetOf((_this).F27()))
				}
			}
			return _coll52.ToMap()
		}()).Merge((Companion_D4_.Create_DC12_(_1306_v123)).Dtor_cf21()), _1303_v120)
		var _1309_v125 _dafny.Sequence
		_ = _1309_v125
		_1309_v125 = _dafny.SeqOf(_1306_v123)
		_1307_v124 = (_1307_v124).Update((_1309_v125).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1309_v125).Cardinality()))).Uint32()).(_dafny.Map), _1303_v120)
		if _1303_v120 {
			(globalState).F6 = _dafny.IntOfUint32((p1).Cardinality())
			var _1310_v126 _dafny.Map
			_ = _1310_v126
			_1310_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1303_v120, _1303_v120)
			var _1311_v127 _dafny.Array
			_ = _1311_v127
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_26
			var _nw170 _dafny.Array
			_ = _nw170
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw170 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) bool = (func(_1312_v120 bool) func(_dafny.Int) bool {
					return func(_1313_i11 _dafny.Int) bool {
						return _1312_v120
					}
				})(_1303_v120)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw170 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw170).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw170).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1311_v127 = _nw170
			var _1314_v128 _dafny.Sequence
			_ = _1314_v128
			_1314_v128 = _dafny.SeqOf(_1311_v127)
			var _1315_v129 *C2
			_ = _1315_v129
			var _nw171 *C2 = New_C2_()
			_ = _nw171
			_nw171.Ctor__(((func() _dafny.Map {
				if _1303_v120 {
					return _1310_v126
				}
				return _1310_v126
			})()).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1311_v127, _1311_v127), _1314_v128))
			_1315_v129 = _nw171
			var _1316_v130 _dafny.Sequence
			_ = _1316_v130
			_1316_v130 = _dafny.SeqOf(_1305_v122)
			var _1317_v131 D11
			_ = _1317_v131
			_1317_v131 = Companion_D11_.Create_DC36_(_1316_v130)
			var _rhs229 bool = _1303_v120
			_ = _rhs229
			var _rhs230 D11 = _1317_v131
			_ = _rhs230
			var _lhs199 *GlobalState = globalState
			_ = _lhs199
			_lhs199.F10 = _rhs229
			_1317_v131 = _rhs230
			_1304_v121 = (_1304_v121).Update(_1303_v120, _dafny.IntOfUint32((p1).Cardinality()))
			var _source20 D0 = _1131_v0
			_ = _source20
			if _source20.Is_DC0() {
				var _1318_v132 *C3
				_ = _1318_v132
				var _nw172 *C3 = New_C3_()
				_ = _nw172
				_nw172.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1314_v128, _1314_v128))
				_1318_v132 = _nw172
				(globalState).F8 = (_this).F27()
				var _1319_v133 _dafny.MultiSet
				_ = _1319_v133
				_1319_v133 = _dafny.MultiSetOf(_1303_v120, _1303_v120)
				var _1320_v134 _dafny.MultiSet
				_ = _1320_v134
				_1320_v134 = _dafny.MultiSetOf(_1319_v133, _1319_v133, _1319_v133)
				(globalState).F1 = (_1320_v134).Difference(_1320_v134)
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_1311_v127), 0))
				_ = _index189
				(_1311_v127).ArraySet1(false, (_index189).Int())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_1311_v127), 0))
				_ = _index190
				(_1311_v127).ArraySet1(!(!(_1303_v120) || (true)), (_index190).Int())
			} else {
				var _1321___mcc_h13 _dafny.Int = _source20.Get_().(D0_DC1).Cf0
				_ = _1321___mcc_h13
				var _1322___mcc_h14 bool = _source20.Get_().(D0_DC1).Cf1
				_ = _1322___mcc_h14
				var _1323___mcc_h15 _dafny.Map = _source20.Get_().(D0_DC1).Cf2
				_ = _1323___mcc_h15
				var _1324___mcc_h16 _dafny.Int = _source20.Get_().(D0_DC1).Cf3
				_ = _1324___mcc_h16
				var _1325___mcc_h17 _dafny.CodePoint = _source20.Get_().(D0_DC1).Cf4
				_ = _1325___mcc_h17
				var _1326_cf4 _dafny.CodePoint = _1325___mcc_h17
				_ = _1326_cf4
				var _1327_cf3 _dafny.Int = _1324___mcc_h16
				_ = _1327_cf3
				var _1328_cf2 _dafny.Map = _1323___mcc_h15
				_ = _1328_cf2
				var _1329_cf1 bool = _1322___mcc_h14
				_ = _1329_cf1
				var _1330_cf0 _dafny.Int = _1321___mcc_h13
				_ = _1330_cf0
				var _1331_v135 _dafny.Sequence
				_ = _1331_v135
				_1331_v135 = _dafny.UnicodeSeqOfUtf8Bytes("riagx")
				var _rhs231 _dafny.Sequence = p1
				_ = _rhs231
				var _rhs232 _dafny.Int = _1327_cf3
				_ = _rhs232
				_1331_v135 = _rhs231
				_1327_cf3 = _rhs232
				var _1332_v136 *C0
				_ = _1332_v136
				var _nw173 *C0 = New_C0_()
				_ = _nw173
				_nw173.Ctor__()
				_1332_v136 = _nw173
				var _1333_v137 *C4
				_ = _1333_v137
				var _nw174 *C4 = New_C4_()
				_ = _nw174
				_nw174.Ctor__(!(!(_1303_v120)), _1314_v128)
				_1333_v137 = _nw174
				var _1334_v138 _dafny.Set
				_ = _1334_v138
				_1334_v138 = _dafny.SetOf(_1333_v137)
				(globalState).F8 = ((p2).Times((_1334_v138).Cardinality())).Minus(_dafny.IntOfUint32((p1).Cardinality()))
				var _1335_v139 _dafny.Sequence
				_ = _1335_v139
				_1335_v139 = _dafny.SeqOf((_1315_v129).F29(), (_1315_v129).F29())
				var _1336_v140 D2
				_ = _1336_v140
				_1336_v140 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_1335_v139).Cardinality()))
				(globalState).F0 = (func() bool {
					if ((_this).F26()).Contains((_1330_cf0).Plus(_dafny.IntOfUint32((Companion_Default___.Fm14(p2, (_1336_v140).Dtor_cf9(), globalState)).Cardinality()))) {
						return ((_this).F26()).Get((_1330_cf0).Plus(_dafny.IntOfUint32((Companion_Default___.Fm14(p2, (_1336_v140).Dtor_cf9(), globalState)).Cardinality()))).(bool)
					}
					return ((_this).F27()).Cmp((_1315_v129).F29()) <= 0
				})()
			}
		} else {
			var _1337_v141 _dafny.Map
			_ = _1337_v141
			_1337_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _1338_v142 _dafny.Sequence
			_ = _1338_v142
			_1338_v142 = _dafny.SeqOf((_this).F27())
			_1337_v141 = (_1337_v141).Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("kceepx")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("kceepx"))).Cardinality()))).Uint32(), _dafny.CodePoint('j')), (_dafny.MultiSetOf(p2, _dafny.IntOfUint32((_1338_v142).Cardinality()), p0)).Cardinality())
			var _1339_v143 D4
			_ = _1339_v143
			_1339_v143 = Companion_D4_.Create_DC15_(p2, (_1304_v121).Cardinality())
			(globalState).F10 = !((Companion_Default___.Fm7((_1339_v143).Dtor_cf26(), globalState)).Cmp(p2) == 0)
			var _1340_v144 _dafny.Sequence
			_ = _1340_v144
			_1340_v144 = _dafny.UnicodeSeqOfUtf8Bytes("bxtwiau")
			_1340_v144 = _1340_v144
			var _1341_v145 _dafny.Map
			_ = _1341_v145
			_1341_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1340_v144)
			var _1342_v146 _dafny.Set
			_ = _1342_v146
			_1342_v146 = _dafny.SetOf(_1303_v120)
			var _1343_v147 _dafny.Set
			_ = _1343_v147
			_1343_v147 = _dafny.SetOf((_1342_v146).Cardinality(), p0)
			var _1344_v148 _dafny.Map
			_ = _1344_v148
			_1344_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1343_v147, _1303_v120)
			_1341_v145 = (_1341_v145).Update((p2).Plus((_1344_v148).Cardinality()), (func() _dafny.Sequence {
				if !(_1303_v120) {
					return p1
				}
				return p1
			})())
			(globalState).F0 = _1303_v120
		}
		var _1345_v149 _dafny.CodePoint
		_ = _1345_v149
		_1345_v149 = _dafny.CodePoint('q')
		var _1346_v150 _dafny.Array
		_ = _1346_v150
		var _nw175 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw175
		_1346_v150 = _nw175
		var _pat_let_tv55 = _1345_v149
		_ = _pat_let_tv55
		var _1347_v151 *C2
		_ = _1347_v151
		var _nw176 *C2 = New_C2_()
		_ = _nw176
		_nw176.Ctor__(p2, _dafny.SeqOf((func(_pat_let24_0 D4) D4 {
			return func(_1348_dt__update__tmp_h2 D4) D4 {
				return func(_pat_let25_0 _dafny.CodePoint) D4 {
					return func(_1349_dt__update_hcf23_h0 _dafny.CodePoint) D4 {
						return Companion_D4_.Create_DC14_(_1349_dt__update_hcf23_h0, (_1348_dt__update__tmp_h2).Dtor_cf24(), (_1348_dt__update__tmp_h2).Dtor_cf25())
					}(_pat_let25_0)
				}(_pat_let_tv55)
			}(_pat_let24_0)
		}(Companion_D4_.Create_DC14_(_1345_v149, _1346_v150, _dafny.UnicodeSeqOfUtf8Bytes("jixw")))).Dtor_cf24()))
		_1347_v151 = _nw176
		var _1350_v152 _dafny.MultiSet
		_ = _1350_v152
		_1350_v152 = _dafny.MultiSetOf(false, _1303_v120)
		var _1351_v153 _dafny.Sequence
		_ = _1351_v153
		_1351_v153 = _dafny.SeqOf(_1303_v120)
		(globalState).F11 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm24((_1350_v152).Difference(_1350_v152), p2, (_dafny.IntOfUint32((_1351_v153).Cardinality())).Minus(p2), !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jrtm"), p1), globalState), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1347_v151).F29()), _dafny.IntOfUint32((Companion_Default___.Fm24((_1350_v152).Difference(_1350_v152), p2, (_dafny.IntOfUint32((_1351_v153).Cardinality())).Minus(p2), !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jrtm"), p1), globalState)).Cardinality()))).Uint32(), p2)).Cardinality())
	}
}
func (_this *C5) M3(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) (_dafny.Int, D0) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var _1352_v0 _dafny.Array
		_ = _1352_v0
		var _nw177 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw177
		_1352_v0 = _nw177
		for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1352_v0), 0))); ; {
			_guard_loop_2, _ok55 := _iter55()
			if !_ok55 {
				break
			}
			var _1353_i0 _dafny.Int
			_1353_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1353_i0).Sign() != -1) && ((_1353_i0).Cmp(_dafny.ArrayLen((_1352_v0), 0)) < 0)) {
				(_1352_v0).ArraySet1(!(!((Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mtnfgu")).Cardinality()), (_this).F27())).Cmp((_this).F27()) > 0)), (_1353_i0).Int())
			}
		}
		var _1354_v1 _dafny.Sequence
		_ = _1354_v1
		_1354_v1 = _dafny.SeqOf((_this).F27())
		var _1355_v2 _dafny.Sequence
		_ = _1355_v2
		_1355_v2 = _dafny.SeqOf(_1354_v1)
		var _1356_v4 D7
		_ = _1356_v4
		_1356_v4 = Companion_D7_.Create_DC25_(func() _dafny.Set {
			var _coll53 = _dafny.NewBuilder()
			_ = _coll53
			for _iter56 := _dafny.Iterate((_1355_v2).Elements()); ; {
				_compr_53, _ok56 := _iter56()
				if !_ok56 {
					break
				}
				var _1357_v3 _dafny.Sequence
				_1357_v3 = interface{}(_compr_53).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_1355_v2, _1357_v3) {
					_coll53.Add(_1357_v3)
				}
			}
			return _coll53.ToSet()
		}())
		var _1358_v5 _dafny.Map
		_ = _1358_v5
		_1358_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1356_v4, (_this).F27())
		var _1359_v7 _dafny.Map
		_ = _1359_v7
		_1359_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, func() _dafny.Map {
			var _coll54 = _dafny.NewMapBuilder()
			_ = _coll54
			for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-409), _dafny.IntOfInt64(24))); ; {
				_compr_54, _ok57 := _iter57()
				if !_ok57 {
					break
				}
				var _1360_v6 _dafny.Int
				_1360_v6 = interface{}(_compr_54).(_dafny.Int)
				if ((_dafny.IntOfInt64(-409)).Cmp(_1360_v6) <= 0) && ((_1360_v6).Cmp(_dafny.IntOfInt64(24)) < 0) {
					_coll54.Add((_1360_v6).Times((_this).F27()), _dafny.IntOfUint32((p0).Cardinality()))
				}
			}
			return _coll54.ToMap()
		}())
		var _1361_v8 _dafny.Sequence
		_ = _1361_v8
		_1361_v8 = _dafny.SeqOf(p1, p1)
		var _1362_v9 _dafny.CodePoint
		_ = _1362_v9
		_1362_v9 = _dafny.CodePoint('u')
		var _1363_v10 D0
		_ = _1363_v10
		_1363_v10 = Companion_D0_.Create_DC1_((_1358_v5).Cardinality(), p1, _1359_v7, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1361_v8).Cardinality()))).Cardinality()), _1362_v9)
		var _1364_v11 _dafny.MultiSet
		_ = _1364_v11
		_1364_v11 = _dafny.MultiSetOf(_1363_v10, _1363_v10)
		var _1365_v12 _dafny.Array
		_ = _1365_v12
		var _nwElement0_37 _dafny.MultiSet = (_dafny.MultiSetOf(_1363_v10)).Difference(_1364_v11)
		_ = _nwElement0_37
		var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.One)
		_ = _nw178
		(_nw178).ArraySet1(_nwElement0_37, 0)
		_1365_v12 = _nw178
		var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1365_v12), 0))
		_ = _index191
		(_1365_v12).ArraySet1(_1364_v11, (_index191).Int())
		var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_1365_v12), 0))
		_ = _index192
		(_1365_v12).ArraySet1(_1364_v11, (_index192).Int())
		var _1366_v13 D6
		_ = _1366_v13
		_1366_v13 = Companion_D6_.Create_DC22_((_this).F27())
		var _hi3 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F27(), (_1366_v13).Dtor_cf36())
		_ = _hi3
		for _1367_i1 := (_this).F27(); _1367_i1.Cmp(_hi3) < 0; _1367_i1 = _1367_i1.Plus(_dafny.One) {
			var _1368_v14 *C0
			_ = _1368_v14
			var _nw179 *C0 = New_C0_()
			_ = _nw179
			_nw179.Ctor__()
			_1368_v14 = _nw179
			var _1369_v15 _dafny.Sequence
			_ = _1369_v15
			_1369_v15 = _dafny.SeqOf(_1368_v14, _1368_v14)
			var _1370_v16 _dafny.Array
			_ = _1370_v16
			var _nwElement0_38 *C0 = _1368_v14
			_ = _nwElement0_38
			var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(21))
			_ = _nw180
			(_nw180).ArraySet1(_nwElement0_38, 0)
			(_nw180).ArraySet1(_1368_v14, 1)
			(_nw180).ArraySet1(_1368_v14, 2)
			(_nw180).ArraySet1(_1368_v14, 3)
			(_nw180).ArraySet1(_1368_v14, 4)
			(_nw180).ArraySet1(_1368_v14, 5)
			(_nw180).ArraySet1((_1369_v15).Select((Companion_Default___.SafeIndex(_1367_i1, _dafny.IntOfUint32((_1369_v15).Cardinality()))).Uint32()).(*C0), 6)
			(_nw180).ArraySet1(_1368_v14, 7)
			(_nw180).ArraySet1(_1368_v14, 8)
			(_nw180).ArraySet1(_1368_v14, 9)
			(_nw180).ArraySet1(_1368_v14, 10)
			(_nw180).ArraySet1(_1368_v14, 11)
			(_nw180).ArraySet1(_1368_v14, 12)
			(_nw180).ArraySet1(_1368_v14, 13)
			(_nw180).ArraySet1(_1368_v14, 14)
			(_nw180).ArraySet1(_1368_v14, 15)
			(_nw180).ArraySet1(_1368_v14, 16)
			(_nw180).ArraySet1(_1368_v14, 17)
			(_nw180).ArraySet1(_1368_v14, 18)
			(_nw180).ArraySet1(_1368_v14, 19)
			(_nw180).ArraySet1(_1368_v14, 20)
			_1370_v16 = _nw180
			var _1371_v17 _dafny.Array
			_ = _1371_v17
			var _nwElement0_39 _dafny.Array = _1370_v16
			_ = _nwElement0_39
			var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(4))
			_ = _nw181
			(_nw181).ArraySet1(_nwElement0_39, 0)
			(_nw181).ArraySet1(_1370_v16, 1)
			(_nw181).ArraySet1(_1370_v16, 2)
			(_nw181).ArraySet1(_1370_v16, 3)
			_1371_v17 = _nw181
			_1371_v17 = _1371_v17
			var _1372_v18 _dafny.Sequence
			_ = _1372_v18
			_1372_v18 = _dafny.SeqOf(_1352_v0, _1352_v0, _1352_v0, _1352_v0)
			var _1373_v19 *C4
			_ = _1373_v19
			var _nw182 *C4 = New_C4_()
			_ = _nw182
			_nw182.Ctor__(p1, _1372_v18)
			_1373_v19 = _nw182
			(_1373_v19).F28 = p1
			var _1374_v20 _dafny.Sequence
			_ = _1374_v20
			_1374_v20 = _dafny.UnicodeSeqOfUtf8Bytes("myxqtmm")
			_1374_v20 = _1374_v20
		}
		var _1375_v21 _dafny.Array
		_ = _1375_v21
		var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw183
		_1375_v21 = _nw183
		var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1375_v21), 0))
		_ = _index193
		(_1375_v21).ArraySet1(_dafny.IntOfInt64(361), (_index193).Int())
		var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1375_v21), 0))
		_ = _index194
		(_1375_v21).ArraySet1((_this).F27(), (_index194).Int())
		var _1376_v22 D2
		_ = _1376_v22
		_1376_v22 = Companion_D2_.Create_DC6_((_1375_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1375_v21), 0))).Int()).(_dafny.Int))
		var _1377_v23 _dafny.Map
		_ = _1377_v23
		_1377_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1376_v22).Dtor_cf9(), (_this).F27())
		_1377_v23 = (_1377_v23).Update(Companion_Default___.SafeDivisionInt((_this).F27(), (_this).F27()), (_this).F27())
		var _rhs233 D6 = Companion_D6_.Create_DC22_((_this).F27())
		_ = _rhs233
		var _rhs234 bool = p1
		_ = _rhs234
		var _lhs200 *GlobalState = globalState
		_ = _lhs200
		_1366_v13 = _rhs233
		_lhs200.F10 = _rhs234
		r0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1361_v8).Cardinality()), (_dafny.Zero).Minus((func() _dafny.Int {
			if p1 {
				return _dafny.IntOfUint32((p0).Cardinality())
			}
			return (_1375_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1375_v21), 0))).Int()).(_dafny.Int)
		})()))
		r1 = Companion_D0_.Create_DC0_()
		return r0, r1
	}
}
func (_this *C5) F26() _dafny.Map {
	{
		return _this._f26
	}
}
func (_this *C5) F27() _dafny.Int {
	{
		return _this._f27
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f22 _dafny.Sequence
	_f24 bool
	_f25 _dafny.Sequence
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f22 = _dafny.EmptySeq
	_this._f24 = false
	_this._f25 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F22() _dafny.Sequence {
	return _this._f22
}
func (_this *C6) Ctor__(f24 bool, f25 _dafny.Sequence, f22 _dafny.Sequence) {
	{
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this)._f22 = f22
	}
}
func (_this *C6) Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(378), _dafny.IntOfInt64(-343))).Plus(_dafny.IntOfInt64(-858))
	}
}
func (_this *C6) Fm3(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return (_this).F24()
	}
}
func (_this *C6) M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1378_v0 _dafny.Array
		_ = _1378_v0
		var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
		_ = _nw184
		_1378_v0 = _nw184
		for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1378_v0), 0))); ; {
			_guard_loop_3, _ok58 := _iter58()
			if !_ok58 {
				break
			}
			var _1379_i0 _dafny.Int
			_1379_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1379_i0).Sign() != -1) && ((_1379_i0).Cmp(_dafny.ArrayLen((_1378_v0), 0)) < 0)) {
				(_1378_v0).ArraySet1((func() _dafny.Sequence {
					if p0 {
						return _dafny.UnicodeSeqOfUtf8Bytes("tri")
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(268))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg87 _dafny.Int) interface{} {
							return coer87(arg87)
						}
					}(func(_1380_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					}))
				})(), (_1379_i0).Int())
			}
		}
		(globalState).F0 = (_this).F24()
		var _1381_v1 D0
		_ = _1381_v1
		_1381_v1 = Companion_D0_.Create_DC0_()
		var _source21 D0 = _1381_v1
		_ = _source21
		if _source21.Is_DC0() {
			var _1382_v2 _dafny.Int
			_ = _1382_v2
			_1382_v2 = _dafny.IntOfInt64(468)
			(globalState).F16 = _1382_v2
			var _1383_v3 *C3
			_ = _1383_v3
			var _nw185 *C3 = New_C3_()
			_ = _nw185
			_nw185.Ctor__(_dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()))
			_1383_v3 = _nw185
			if p0 {
				var _1384_v4 *C2
				_ = _1384_v4
				var _nw186 *C2 = New_C2_()
				_ = _nw186
				_nw186.Ctor__(_1382_v2, _dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()))
				_1384_v4 = _nw186
				var _1385_v5 _dafny.Array
				_ = _1385_v5
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_27
				var _nw187 _dafny.Array
				_ = _nw187
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw187 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) D11 = func(_1386_i2 _dafny.Int) D11 {
						return Companion_D11_.Create_DC39_((_this).F24())
					}
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw187 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw187).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw187).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1385_v5 = _nw187
				var _1387_v6 D11
				_ = _1387_v6
				_1387_v6 = Companion_D11_.Create_DC39_((_this).F24())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1385_v5), 0))
				_ = _index195
				(_1385_v5).ArraySet1(_1387_v6, (_index195).Int())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1385_v5), 0))
				_ = _index196
				(_1385_v5).ArraySet1(_1387_v6, (_index196).Int())
				var _1388_v7 _dafny.Sequence
				_ = _1388_v7
				_1388_v7 = _dafny.UnicodeSeqOfUtf8Bytes("kfs")
				var _1389_v8 _dafny.Array
				_ = _1389_v8
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_28
				var _nw188 _dafny.Array
				_ = _nw188
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw188 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Int = (func(_1390_v4 *C2) func(_dafny.Int) _dafny.Int {
						return func(_1391_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1391_i3, (_1390_v4).F29())
						}
					})(_1384_v4)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw188 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw188).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw188).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1389_v8 = _nw188
				var _1392_v9 _dafny.Map
				_ = _1392_v9
				_1392_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1388_v7).Cardinality()), _1389_v8)
				var _1393_v10 _dafny.Sequence
				_ = _1393_v10
				_1393_v10 = _dafny.SeqOf(_dafny.IntOfUint32((_1388_v7).Cardinality()), (_1384_v4).F29())
				var _1394_v11 _dafny.Map
				_ = _1394_v11
				_1394_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1384_v4).F29(), _dafny.IntOfUint32((_1393_v10).Cardinality()))
				var _1395_v12 _dafny.Sequence
				_ = _1395_v12
				_1395_v12 = _dafny.SeqOf((func() _dafny.Int {
					if (_1394_v11).Contains(_dafny.IntOfUint32((_1393_v10).Cardinality())) {
						return (_1394_v11).Get(_dafny.IntOfUint32((_1393_v10).Cardinality())).(_dafny.Int)
					}
					return Companion_Default___.Fm7(_1382_v2, globalState)
				})())
				var _1396_v13 D2
				_ = _1396_v13
				_1396_v13 = Companion_D2_.Create_DC5_(_1389_v8)
				_1392_v9 = (_1392_v9).Update((_dafny.SetOf(_1382_v2, _1382_v2, (_1395_v12).Select((Companion_Default___.SafeIndex(_1382_v2, _dafny.IntOfUint32((_1395_v12).Cardinality()))).Uint32()).(_dafny.Int), (Companion_Default___.Fm33(globalState)).Cardinality())).Cardinality(), (_1396_v13).Dtor_cf8())
				var _1397_v14 _dafny.CodePoint
				_ = _1397_v14
				_1397_v14 = _dafny.CodePoint('q')
				var _1398_v15 _dafny.MultiSet
				_ = _1398_v15
				_1398_v15 = _dafny.MultiSetOf((_this).F24())
				var _1399_v16 _dafny.Map
				_ = _1399_v16
				_1399_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0) || ((_this).F24()), (_dafny.MultiSetOf(_dafny.IntOfInt64(502), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1397_v14, (_1384_v4).F29())).Cardinality())).Update((_1398_v15).Cardinality(), Companion_Default___.Abs(_1382_v2)))
				_1399_v16 = (_1399_v16).Update((_this).F24(), (_dafny.MultiSetFromSeq(_1393_v10)).Intersection(_dafny.MultiSetOf((_1384_v4).F29())))
				var _1400_v17 _dafny.MultiSet
				_ = _1400_v17
				_1400_v17 = _dafny.MultiSetOf((_1384_v4).F29(), _1382_v2, _dafny.IntOfUint32((_1388_v7).Cardinality()), _1382_v2)
				var _1401_v21 _dafny.Sequence
				_ = _1401_v21
				_1401_v21 = _dafny.SeqOf(func() _dafny.Set {
					var _coll55 = _dafny.NewBuilder()
					_ = _coll55
					for _iter59 := _dafny.Iterate((func() _dafny.Map {
						var _coll56 = _dafny.NewMapBuilder()
						_ = _coll56
						for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(325), _dafny.IntOfInt64(-867))); ; {
							_compr_56, _ok60 := _iter60()
							if !_ok60 {
								break
							}
							var _1402_v18 _dafny.Int
							_1402_v18 = interface{}(_compr_56).(_dafny.Int)
							if ((_dafny.IntOfInt64(325)).Cmp(_1402_v18) <= 0) && ((_1402_v18).Cmp(_dafny.IntOfInt64(-867)) < 0) {
								_coll56.Add((_1402_v18).Times(_1382_v2), p0)
							}
						}
						return _coll56.ToMap()
					}()).Keys().Elements()); ; {
						_compr_55, _ok59 := _iter59()
						if !_ok59 {
							break
						}
						var _1403_v19 _dafny.Int
						_1403_v19 = interface{}(_compr_55).(_dafny.Int)
						if (func() _dafny.Map {
							var _coll57 = _dafny.NewMapBuilder()
							_ = _coll57
							for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(325), _dafny.IntOfInt64(-867))); ; {
								_compr_57, _ok61 := _iter61()
								if !_ok61 {
									break
								}
								var _1402_v18 _dafny.Int
								_1402_v18 = interface{}(_compr_57).(_dafny.Int)
								if ((_dafny.IntOfInt64(325)).Cmp(_1402_v18) <= 0) && ((_1402_v18).Cmp(_dafny.IntOfInt64(-867)) < 0) {
									_coll57.Add((_1402_v18).Times(_1382_v2), p0)
								}
							}
							return _coll57.ToMap()
						}()).Contains(_1403_v19) {
							_coll55.Add((_1403_v19).Times((func() _dafny.Map {
								var _coll58 = _dafny.NewMapBuilder()
								_ = _coll58
								for _iter62 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(219), (_dafny.SetOf(!(false), false, false, true, false)).Cardinality())).Keys().Elements()); ; {
									_compr_58, _ok62 := _iter62()
									if !_ok62 {
										break
									}
									var _1404_v20 _dafny.Int
									_1404_v20 = interface{}(_compr_58).(_dafny.Int)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(219), (_dafny.SetOf(!(false), false, false, true, false)).Cardinality())).Contains(_1404_v20) {
										_coll58.Add(Companion_Default___.SafeDivisionInt(_1404_v20, (_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false), true))).Cardinality())).Cardinality()), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-324)))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kqmipstkb")).Cardinality()), _dafny.CodePoint('g'))).Cardinality()))
									}
								}
								return _coll58.ToMap()
							}()).Cardinality()))
						}
					}
					return _coll55.ToSet()
				}())
				var _1405_v22 _dafny.Set
				_ = _1405_v22
				_1405_v22 = _dafny.SetOf(_1382_v2)
				(globalState).F11 = (_1384_v4).Fm0(_1400_v17, _dafny.IntOfInt64(574), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1401_v21, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.IntOfUint32((_1401_v21).Cardinality()))).Uint32(), _dafny.SetOf((_1384_v4).F29(), _dafny.IntOfUint32((_1395_v12).Cardinality()), _1382_v2, (_1384_v4).F29())), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1384_v4).F29()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1401_v21, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.IntOfUint32((_1401_v21).Cardinality()))).Uint32(), _dafny.SetOf((_1384_v4).F29(), _dafny.IntOfUint32((_1395_v12).Cardinality()), _1382_v2, (_1384_v4).F29()))).Cardinality()))).Uint32(), _1405_v22), p0, globalState)
			} else {
				var _1406_v23 _dafny.Sequence
				_ = _1406_v23
				_1406_v23 = _dafny.SeqOf(_1382_v2)
				var _1407_v24 _dafny.Sequence
				_ = _1407_v24
				_1407_v24 = _dafny.SeqOf(_1406_v23)
				var _1408_v25 _dafny.Set
				_ = _1408_v25
				_1408_v25 = _dafny.SetOf(_1382_v2, _1382_v2, _1382_v2, _1382_v2)
				r0 = !(_dafny.Companion_Sequence_.Equal(_1406_v23, (_1407_v24).Select((Companion_Default___.SafeIndex((_1408_v25).Cardinality(), _dafny.IntOfUint32((_1407_v24).Cardinality()))).Uint32()).(_dafny.Sequence)))
				_1378_v0 = _1378_v0
				(globalState).F6 = _1382_v2
				r3 = (func() _dafny.Int {
					if p0 {
						return _1382_v2
					}
					return _1382_v2
				})()
				(globalState).F0 = (_this).F24()
			}
			var _1409_v26 _dafny.Map
			_ = _1409_v26
			_1409_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), (_this).F24())
			var _1410_v27 _dafny.CodePoint
			_ = _1410_v27
			_1410_v27 = _dafny.CodePoint('m')
			var _1411_v28 _dafny.Array
			_ = _1411_v28
			var _nwElement0_40 bool = (_this).F24()
			_ = _nwElement0_40
			var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(9))
			_ = _nw189
			(_nw189).ArraySet1(_nwElement0_40, 0)
			(_nw189).ArraySet1(p0, 1)
			(_nw189).ArraySet1(!((func() bool {
				if (_1409_v26).Contains(_1410_v27) {
					return (_1409_v26).Get(_1410_v27).(bool)
				}
				return (_this).F24()
			})()), 2)
			(_nw189).ArraySet1(p0, 3)
			(_nw189).ArraySet1((_this).F24(), 4)
			(_nw189).ArraySet1(p0, 5)
			(_nw189).ArraySet1((_this).F24(), 6)
			(_nw189).ArraySet1(p0, 7)
			(_nw189).ArraySet1(p0, 8)
			_1411_v28 = _nw189
			var _1412_v29 *C3
			_ = _1412_v29
			var _nw190 *C3 = New_C3_()
			_ = _nw190
			_nw190.Ctor__(_dafny.SeqOf(_1411_v28, _1411_v28, _1411_v28, _1411_v28, _1411_v28))
			_1412_v29 = _nw190
		} else {
			var _1413___mcc_h0 _dafny.Int = _source21.Get_().(D0_DC1).Cf0
			_ = _1413___mcc_h0
			var _1414___mcc_h1 bool = _source21.Get_().(D0_DC1).Cf1
			_ = _1414___mcc_h1
			var _1415___mcc_h2 _dafny.Map = _source21.Get_().(D0_DC1).Cf2
			_ = _1415___mcc_h2
			var _1416___mcc_h3 _dafny.Int = _source21.Get_().(D0_DC1).Cf3
			_ = _1416___mcc_h3
			var _1417___mcc_h4 _dafny.CodePoint = _source21.Get_().(D0_DC1).Cf4
			_ = _1417___mcc_h4
			var _1418_cf4 _dafny.CodePoint = _1417___mcc_h4
			_ = _1418_cf4
			var _1419_cf3 _dafny.Int = _1416___mcc_h3
			_ = _1419_cf3
			var _1420_cf2 _dafny.Map = _1415___mcc_h2
			_ = _1420_cf2
			var _1421_cf1 bool = _1414___mcc_h1
			_ = _1421_cf1
			var _1422_cf0 _dafny.Int = _1413___mcc_h0
			_ = _1422_cf0
			var _1423_v30 bool
			_ = _1423_v30
			var _1424_v31 _dafny.Int
			_ = _1424_v31
			var _out56 bool
			_ = _out56
			var _out57 _dafny.Int
			_ = _out57
			_out56, _out57 = (_this).M1(globalState)
			_1423_v30 = _out56
			_1424_v31 = _out57
			var _1425_v32 _dafny.Array
			_ = _1425_v32
			var _len0_29 _dafny.Int = _dafny.One
			_ = _len0_29
			var _nw191 _dafny.Array
			_ = _nw191
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw191 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Int = (func(_1426_cf0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1427_i4 _dafny.Int) _dafny.Int {
						return (_1427_i4).Times(_1426_cf0)
					}
				})(_1422_cf0)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw191 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw191).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw191).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1425_v32 = _nw191
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1425_v32), 0))
			_ = _index197
			(_1425_v32).ArraySet1(_1419_cf3, (_index197).Int())
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1425_v32), 0))
			_ = _index198
			(_1425_v32).ArraySet1(((_dafny.Zero).Minus(_1422_cf0)).Minus(_1424_v31), (_index198).Int())
			if _1423_v30 {
				var _1428_v33 _dafny.Array
				_ = _1428_v33
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
				_ = _nw192
				_1428_v33 = _nw192
				_1428_v33 = _1428_v33
				var _1429_v34 *C0
				_ = _1429_v34
				var _nw193 *C0 = New_C0_()
				_ = _nw193
				_nw193.Ctor__()
				_1429_v34 = _nw193
				(globalState).F6 = Companion_Default___.SafeModuloInt(_1422_cf0, _1422_cf0)
				var _1430_v35 _dafny.Array
				_ = _1430_v35
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_30
				var _nw194 _dafny.Array
				_ = _nw194
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw194 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Set = (func(_1431_cf0 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_1432_i5 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_1431_cf0, _1431_cf0)
						}
					})(_1422_cf0)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw194 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw194).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw194).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1430_v35 = _nw194
				var _1433_v36 _dafny.Map
				_ = _1433_v36
				_1433_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1419_cf3, _1424_v31)
				var _1434_v37 _dafny.Sequence
				_ = _1434_v37
				_1434_v37 = _dafny.SeqOf(_1433_v36, _1433_v36)
				var _1435_v38 _dafny.Map
				_ = _1435_v38
				_1435_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1424_v31)
				var _1436_v40 _dafny.Set
				_ = _1436_v40
				_1436_v40 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1434_v37, (Companion_Default___.SafeIndex((_1435_v38).Cardinality(), _dafny.IntOfUint32((_1434_v37).Cardinality()))).Uint32(), func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-383), _dafny.IntOfInt64(-369))); ; {
						_compr_59, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1437_v39 _dafny.Int
						_1437_v39 = interface{}(_compr_59).(_dafny.Int)
						if ((_dafny.IntOfInt64(-383)).Cmp(_1437_v39) <= 0) && ((_1437_v39).Cmp(_dafny.IntOfInt64(-369)) < 0) {
							_coll59.Add(Companion_Default___.SafeModuloInt(_1437_v39, (_1425_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1425_v32), 0))).Int()).(_dafny.Int)), _1422_cf0)
						}
					}
					return _coll59.ToMap()
				}())).Cardinality()), (_1425_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1425_v32), 0))).Int()).(_dafny.Int), _1422_cf0, (_dafny.Zero).Minus(Companion_Default___.Fm7(_1422_cf0, globalState)), _1419_cf3)
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1430_v35), 0))
				_ = _index199
				(_1430_v35).ArraySet1(_1436_v40, (_index199).Int())
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1430_v35), 0))
				_ = _index200
				(_1430_v35).ArraySet1((func() _dafny.Set {
					var _coll60 = _dafny.NewBuilder()
					_ = _coll60
					for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(460), _dafny.IntOfInt64(-440))); ; {
						_compr_60, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _1438_v41 _dafny.Int
						_1438_v41 = interface{}(_compr_60).(_dafny.Int)
						if ((_dafny.IntOfInt64(460)).Cmp(_1438_v41) <= 0) && ((_1438_v41).Cmp(_dafny.IntOfInt64(-440)) < 0) {
							_coll60.Add((_1438_v41).Times((_1425_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1425_v32), 0))).Int()).(_dafny.Int)))
						}
					}
					return _coll60.ToSet()
				}()).Union(_1436_v40), (_index200).Int())
				(globalState).F13 = (_1430_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1430_v35), 0))).Int()).(_dafny.Set)
			} else {
				(globalState).F6 = (_dafny.Zero).Minus(_1422_cf0)
				var _1439_v42 _dafny.Set
				_ = _1439_v42
				_1439_v42 = _dafny.SetOf(_1418_cf4, _dafny.CodePoint('w'), _1418_cf4, _1418_cf4)
				(globalState).F10 = (_dafny.SetOf(_1418_cf4)).IsProperSubsetOf(_1439_v42)
				var _1440_v43 _dafny.Sequence
				_ = _1440_v43
				_1440_v43 = _dafny.UnicodeSeqOfUtf8Bytes("mkn")
				(globalState).F11 = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if !((_this).F24()) {
						return _1440_v43
					}
					return _dafny.Companion_Sequence_.Concatenate(_1440_v43, _1440_v43)
				})()).Cardinality())
				var _1441_v44 _dafny.Sequence
				_ = _1441_v44
				_1441_v44 = _dafny.SeqOf((_this).F24())
				var _1442_v45 _dafny.Array
				_ = _1442_v45
				var _nwElement0_41 bool = Companion_Default___.Fm37(_1441_v44, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1422_cf0, _1423_v30), _dafny.UnicodeSeqOfUtf8Bytes("xqe"), globalState)
				_ = _nwElement0_41
				var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(4))
				_ = _nw195
				(_nw195).ArraySet1(_nwElement0_41, 0)
				(_nw195).ArraySet1(p0, 1)
				(_nw195).ArraySet1(p0, 2)
				(_nw195).ArraySet1(p0, 3)
				_1442_v45 = _nw195
				var _1443_v46 *C4
				_ = _1443_v46
				var _nw196 *C4 = New_C4_()
				_ = _nw196
				_nw196.Ctor__(_1421_cf1, _dafny.SeqOf(_1442_v45, _1442_v45, _1442_v45))
				_1443_v46 = _nw196
				var _1444_v47 _dafny.Sequence
				_ = _1444_v47
				_1444_v47 = _dafny.SeqOf(_1378_v0)
				var _1445_v48 _dafny.Array
				_ = _1445_v48
				var _nwElement0_42 _dafny.Array = _1378_v0
				_ = _nwElement0_42
				var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(20))
				_ = _nw197
				(_nw197).ArraySet1(_nwElement0_42, 0)
				(_nw197).ArraySet1((_1444_v47).Select((Companion_Default___.SafeIndex(_1419_cf3, _dafny.IntOfUint32((_1444_v47).Cardinality()))).Uint32()).(_dafny.Array), 1)
				(_nw197).ArraySet1(_1378_v0, 2)
				(_nw197).ArraySet1(_1378_v0, 3)
				(_nw197).ArraySet1((_1444_v47).Select((Companion_Default___.SafeIndex((_1425_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1425_v32), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1444_v47).Cardinality()))).Uint32()).(_dafny.Array), 4)
				(_nw197).ArraySet1(_1378_v0, 5)
				(_nw197).ArraySet1(_1378_v0, 6)
				(_nw197).ArraySet1(_1378_v0, 7)
				(_nw197).ArraySet1(_1378_v0, 8)
				(_nw197).ArraySet1(_1378_v0, 9)
				(_nw197).ArraySet1(_1378_v0, 10)
				(_nw197).ArraySet1(_1378_v0, 11)
				(_nw197).ArraySet1(_1378_v0, 12)
				(_nw197).ArraySet1(_1378_v0, 13)
				(_nw197).ArraySet1(_1378_v0, 14)
				(_nw197).ArraySet1(_1378_v0, 15)
				(_nw197).ArraySet1(_1378_v0, 16)
				(_nw197).ArraySet1(_1378_v0, 17)
				(_nw197).ArraySet1(_1378_v0, 18)
				(_nw197).ArraySet1(_1378_v0, 19)
				_1445_v48 = _nw197
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1445_v48), 0))
				_ = _index201
				(_1445_v48).ArraySet1(_1378_v0, (_index201).Int())
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1445_v48), 0))
				_ = _index202
				(_1445_v48).ArraySet1((func() _dafny.Array {
					if (func() bool {
						if _1423_v30 {
							return (_this).F24()
						}
						return _1423_v30
					})() {
						return _1378_v0
					}
					return _1378_v0
				})(), (_index202).Int())
			}
			var _1446_v49 _dafny.Sequence
			_ = _1446_v49
			_1446_v49 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			_1418_cf4 = (_1446_v49).Select((Companion_Default___.SafeIndex(_1419_cf3, _dafny.IntOfUint32((_1446_v49).Cardinality()))).Uint32()).(_dafny.CodePoint)
		}
		var _1447_v50 _dafny.Int
		_ = _1447_v50
		_1447_v50 = _dafny.IntOfInt64(-250)
		if Companion_Default___.Fm15((_dafny.MultiSetOf(_1447_v50)).IsSubsetOf(_dafny.MultiSetOf(_1447_v50)), globalState) {
			var _1448_v51 _dafny.Array
			_ = _1448_v51
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_31
			var _nw198 _dafny.Array
			_ = _nw198
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw198 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.MultiSet = (func(_1449_p0 bool) func(_dafny.Int) _dafny.MultiSet {
					return func(_1450_i6 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(_dafny.MultiSetOf((_this).F24(), (_this).F24()), _dafny.MultiSetOf(_1449_p0), _dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F24())))
					}
				})(p0)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw198 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw198).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw198).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1448_v51 = _nw198
			var _1451_v52 _dafny.Sequence
			_ = _1451_v52
			_1451_v52 = _dafny.UnicodeSeqOfUtf8Bytes("cnkxhr")
			var _1452_v53 _dafny.Map
			_ = _1452_v53
			_1452_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(_1451_v52, (_this).F24(), false, globalState), (_this).F24())
			var _1453_v54 _dafny.Set
			_ = _1453_v54
			_1453_v54 = _dafny.SetOf((func() bool {
				if (_1452_v53).Contains(false) {
					return (_1452_v53).Get(false).(bool)
				}
				return (_this).F24()
			})(), p0, (_this).F24(), (_this).F24(), (_this).F24())
			var _1454_v55 _dafny.MultiSet
			_ = _1454_v55
			_1454_v55 = _dafny.MultiSetOf((_this).F24(), (_this).F24())
			var _1455_v56 _dafny.MultiSet
			_ = _1455_v56
			_1455_v56 = _dafny.MultiSetOf(Companion_Default___.Fm41(_1447_v50, _1451_v52, _1453_v54, globalState), _1454_v55, _1454_v55)
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1448_v51), 0))
			_ = _index203
			(_1448_v51).ArraySet1(_1455_v56, (_index203).Int())
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1448_v51), 0))
			_ = _index204
			(_1448_v51).ArraySet1((_1455_v56).Difference(_1455_v56), (_index204).Int())
			var _1456_v57 _dafny.Array
			_ = _1456_v57
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_32
			var _nw199 _dafny.Array
			_ = _nw199
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw199 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) bool = (func(_1457_p0 bool) func(_dafny.Int) bool {
					return func(_1458_i7 _dafny.Int) bool {
						return _1457_p0
					}
				})(p0)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw199 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw199).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw199).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1456_v57 = _nw199
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))
			_ = _index205
			(_1456_v57).ArraySet1(!(p0), (_index205).Int())
			var _1459_v58 _dafny.CodePoint
			_ = _1459_v58
			_1459_v58 = _dafny.CodePoint('i')
			var _1460_v59 *C1
			_ = _1460_v59
			var _nw200 *C1 = New_C1_()
			_ = _nw200
			_nw200.Ctor__(_1447_v50, (_this).F22())
			_1460_v59 = _nw200
			var _1461_v60 _dafny.Map
			_ = _1461_v60
			_1461_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1460_v59)
			var _1462_v61 _dafny.Map
			_ = _1462_v61
			_1462_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1459_v58, _1461_v60)
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))
			_ = _index206
			(_1456_v57).ArraySet1((func() bool {
				if (func() bool {
					if false {
						return false
					}
					return (_this).F24()
				})() {
					return !((_this).F24()) || (false)
				}
				return (_1447_v50).Cmp((_1462_v61).Cardinality()) < 0
			})(), (_index206).Int())
			var _1463_v62 D4
			_ = _1463_v62
			_1463_v62 = Companion_D4_.Create_DC15_(((_1460_v59).F30()).Times(_dafny.IntOfUint32((_1451_v52).Cardinality())), (_1460_v59).F30())
			var _source22 D4 = _1463_v62
			_ = _source22
			if _source22.Is_DC13() {
				var _1464___mcc_h5 D0 = _source22.Get_().(D4_DC13).Cf22
				_ = _1464___mcc_h5
				var _1465_cf22 D0 = _1464___mcc_h5
				_ = _1465_cf22
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))
				_ = _index207
				(_1456_v57).ArraySet1((_this).F24(), (_index207).Int())
				_1459_v58 = _1459_v58
				_1460_v59 = _1460_v59
				var _1466_v63 _dafny.Array
				_ = _1466_v63
				var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
				_ = _nw201
				_1466_v63 = _nw201
				var _1467_v64 _dafny.MultiSet
				_ = _1467_v64
				_1467_v64 = _dafny.MultiSetOf(Companion_D6_.Create_DC22_((_1460_v59).F30()))
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1466_v63), 0))
				_ = _index208
				(_1466_v63).ArraySet1(_1467_v64, (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1466_v63), 0))
				_ = _index209
				(_1466_v63).ArraySet1(Companion_Default___.Fm46(globalState), (_index209).Int())
			} else if _source22.Is_DC14() {
				var _1468___mcc_h6 _dafny.CodePoint = _source22.Get_().(D4_DC14).Cf23
				_ = _1468___mcc_h6
				var _1469___mcc_h7 _dafny.Array = _source22.Get_().(D4_DC14).Cf24
				_ = _1469___mcc_h7
				var _1470___mcc_h8 _dafny.Sequence = _source22.Get_().(D4_DC14).Cf25
				_ = _1470___mcc_h8
				var _1471_cf25 _dafny.Sequence = _1470___mcc_h8
				_ = _1471_cf25
				var _1472_cf24 _dafny.Array = _1469___mcc_h7
				_ = _1472_cf24
				var _1473_cf23 _dafny.CodePoint = _1468___mcc_h6
				_ = _1473_cf23
				_1473_cf23 = _1473_cf23
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))
				_ = _index210
				(_1456_v57).ArraySet1(p0, (_index210).Int())
				var _1474_v65 _dafny.Array
				_ = _1474_v65
				var _nw202 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw202
				_1474_v65 = _nw202
				var _rhs235 _dafny.Array = _1474_v65
				_ = _rhs235
				var _rhs236 bool = (_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool)
				_ = _rhs236
				var _lhs201 *GlobalState = globalState
				_ = _lhs201
				_1474_v65 = _rhs235
				_lhs201.F10 = _rhs236
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1474_v65), 0))
				_ = _index211
				(_1474_v65).ArraySet1(Companion_Default___.SafeModuloInt((_1460_v59).F30(), (_dafny.Zero).Minus((_1460_v59).F30())), (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1474_v65), 0))
				_ = _index212
				(_1474_v65).ArraySet1((_1460_v59).F30(), (_index212).Int())
			} else if _source22.Is_DC15() {
				var _1475___mcc_h9 _dafny.Int = _source22.Get_().(D4_DC15).Cf26
				_ = _1475___mcc_h9
				var _1476___mcc_h10 _dafny.Int = _source22.Get_().(D4_DC15).Cf27
				_ = _1476___mcc_h10
				var _1477_cf27 _dafny.Int = _1476___mcc_h10
				_ = _1477_cf27
				var _1478_cf26 _dafny.Int = _1475___mcc_h9
				_ = _1478_cf26
				var _1479_v66 bool
				_ = _1479_v66
				var _1480_v67 _dafny.Int
				_ = _1480_v67
				var _out58 bool
				_ = _out58
				var _out59 _dafny.Int
				_ = _out59
				_out58, _out59 = (_this).M1(globalState)
				_1479_v66 = _out58
				_1480_v67 = _out59
				var _1481_v68 *C2
				_ = _1481_v68
				var _nw203 *C2 = New_C2_()
				_ = _nw203
				_nw203.Ctor__((_1460_v59).F30(), (_this).F22())
				_1481_v68 = _nw203
				(globalState).F6 = _1447_v50
				var _1482_v69 _dafny.Array
				_ = _1482_v69
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_33
				var _nw204 _dafny.Array
				_ = _nw204
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw204 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Int = (func(_1483_v59 *C1) func(_dafny.Int) _dafny.Int {
						return func(_1484_i8 _dafny.Int) _dafny.Int {
							return (_1484_i8).Minus((_1483_v59).F30())
						}
					})(_1460_v59)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw204 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw204).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw204).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1482_v69 = _nw204
				var _1485_v70 _dafny.Sequence
				_ = _1485_v70
				_1485_v70 = _dafny.SeqOf(_dafny.IntOfInt64(-441))
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_1482_v69), 0))
				_ = _index213
				(_1482_v69).ArraySet1(Companion_Default___.SafeDivisionInt(_1447_v50, (_1485_v70).Select((Companion_Default___.SafeIndex((_1481_v68).F29(), _dafny.IntOfUint32((_1485_v70).Cardinality()))).Uint32()).(_dafny.Int)), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_1482_v69), 0))
				_ = _index214
				(_1482_v69).ArraySet1((_1485_v70).Select((Companion_Default___.SafeIndex((_1460_v59).F30(), _dafny.IntOfUint32((_1485_v70).Cardinality()))).Uint32()).(_dafny.Int), (_index214).Int())
			} else if _source22.Is_DC12() {
				var _1486___mcc_h11 _dafny.Map = _source22.Get_().(D4_DC12).Cf21
				_ = _1486___mcc_h11
				var _1487_cf21 _dafny.Map = _1486___mcc_h11
				_ = _1487_cf21
				(globalState).F16 = (_1460_v59).F30()
				var _1488_v71 *C3
				_ = _1488_v71
				var _nw205 *C3 = New_C3_()
				_ = _nw205
				_nw205.Ctor__((_this).F22())
				_1488_v71 = _nw205
				var _1489_v72 _dafny.Array
				_ = _1489_v72
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw206
				_1489_v72 = _nw206
				var _1490_v73 _dafny.Array
				_ = _1490_v73
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
				_ = _nw207
				_1490_v73 = _nw207
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_1489_v72), 0))
				_ = _index215
				(_1489_v72).ArraySet1(_1490_v73, (_index215).Int())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_1489_v72), 0))
				_ = _index216
				(_1489_v72).ArraySet1(_1490_v73, (_index216).Int())
				var _1491_v74 D4
				_ = _1491_v74
				_1491_v74 = Companion_D4_.Create_DC14_(_1459_v58, _1456_v57, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(778))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}((func(_1492_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1493_i9 _dafny.Int) _dafny.CodePoint {
						return _1492_v58
					}
				})(_1459_v58))))
				var _pat_let_tv56 = _1451_v52
				_ = _pat_let_tv56
				var _rhs237 _dafny.Int = _1447_v50
				_ = _rhs237
				var _rhs238 _dafny.Array = (func(_pat_let26_0 D4) D4 {
					return func(_1494_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let27_0 _dafny.Sequence) D4 {
							return func(_1495_dt__update_hcf25_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC14_((_1494_dt__update__tmp_h0).Dtor_cf23(), (_1494_dt__update__tmp_h0).Dtor_cf24(), _1495_dt__update_hcf25_h0)
							}(_pat_let27_0)
						}(_pat_let_tv56)
					}(_pat_let26_0)
				}(_1491_v74)).Dtor_cf24()
				_ = _rhs238
				var _rhs239 _dafny.Int = (_1460_v59).F30()
				_ = _rhs239
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				r3 = _rhs237
				r1 = _rhs238
				_lhs202.F11 = _rhs239
			} else {
				var _1496___mcc_h12 D4 = _source22.Get_().(D4_DC16).Cf28
				_ = _1496___mcc_h12
				var _1497_cf28 D4 = _1496___mcc_h12
				_ = _1497_cf28
				(globalState).F6 = (_dafny.Zero).Minus((_1460_v59).F30())
				var _1498_v75 D14
				_ = _1498_v75
				_1498_v75 = Companion_D14_.Create_DC47_(_1454_v55)
				(globalState).F0 = ((_1498_v75).Dtor_cf77()).IsSubsetOf((_1454_v55).Difference(_1454_v55))
				var _1499_v76 _dafny.Array
				_ = _1499_v76
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw208
				_1499_v76 = _nw208
				var _1500_v77 _dafny.Map
				_ = _1500_v77
				_1500_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v76, (_1460_v59).F30())
				var _1501_v78 _dafny.Map
				_ = _1501_v78
				_1501_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1447_v50)
				var _1502_v79 _dafny.Map
				_ = _1502_v79
				_1502_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, !((_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool)))
				var _1503_v80 _dafny.Map
				_ = _1503_v80
				_1503_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1502_v79, !((_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool)))
				var _1504_v81 _dafny.Map
				_ = _1504_v81
				_1504_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1454_v55)
				var _1505_v82 _dafny.Sequence
				_ = _1505_v82
				_1505_v82 = _dafny.SeqOf(_1447_v50, (_1460_v59).F30(), (_1460_v59).F30(), (_1460_v59).F30(), (_1504_v81).Cardinality())
				var _1506_v83 _dafny.Sequence
				_ = _1506_v83
				_1506_v83 = _dafny.SeqOf(_1502_v79)
				var _1507_v84 _dafny.Array
				_ = _1507_v84
				var _nwElement0_43 _dafny.Int = (_1460_v59).F30()
				_ = _nwElement0_43
				var _nw209 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(24))
				_ = _nw209
				(_nw209).ArraySet1(_nwElement0_43, 0)
				(_nw209).ArraySet1((_1460_v59).Fm35((_dafny.Zero).Minus((_1460_v59).F30()), (_1500_v77).Cardinality(), p0, _dafny.IntOfInt64(665), globalState), 1)
				(_nw209).ArraySet1((_1501_v78).Cardinality(), 2)
				(_nw209).ArraySet1((_1460_v59).F30(), 3)
				(_nw209).ArraySet1((_1460_v59).F30(), 4)
				(_nw209).ArraySet1((_1460_v59).F30(), 5)
				(_nw209).ArraySet1((_1460_v59).F30(), 6)
				(_nw209).ArraySet1((_1460_v59).F30(), 7)
				(_nw209).ArraySet1((_1460_v59).F30(), 8)
				(_nw209).ArraySet1((_1460_v59).F30(), 9)
				(_nw209).ArraySet1((_1460_v59).F30(), 10)
				(_nw209).ArraySet1((_1460_v59).F30(), 11)
				(_nw209).ArraySet1(_1447_v50, 12)
				(_nw209).ArraySet1(_dafny.IntOfInt64(-563), 13)
				(_nw209).ArraySet1((_1460_v59).F30(), 14)
				(_nw209).ArraySet1((_1503_v80).Cardinality(), 15)
				(_nw209).ArraySet1((_1460_v59).F30(), 16)
				(_nw209).ArraySet1((_1505_v82).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_1460_v59).F30())), _dafny.IntOfUint32((_1505_v82).Cardinality()))).Uint32()).(_dafny.Int), 17)
				(_nw209).ArraySet1((_1460_v59).F30(), 18)
				(_nw209).ArraySet1((_1460_v59).F30(), 19)
				(_nw209).ArraySet1((_dafny.MultiSetFromSeq(_1505_v82)).Cardinality(), 20)
				(_nw209).ArraySet1(_dafny.IntOfUint32((_1506_v83).Cardinality()), 21)
				(_nw209).ArraySet1(_dafny.IntOfInt64(624), 22)
				(_nw209).ArraySet1((_1460_v59).F30(), 23)
				_1507_v84 = _nw209
				var _1508_v85 _dafny.Map
				_ = _1508_v85
				_1508_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v84, (_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool))
				var _1509_v86 D10
				_ = _1509_v86
				_1509_v86 = Companion_D10_.Create_DC34_(_1459_v58, (_this).F24(), (Companion_D8_.Create_DC29_(!(p0), p0, _dafny.CodePoint('l'))).Dtor_cf48(), (_this).F24(), true)
				var _1510_v87 T0
				_ = _1510_v87
				var _nw210 *C4 = New_C4_()
				_ = _nw210
				_nw210.Ctor__((_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool), (_this).F22())
				_1510_v87 = _nw210
				var _1511_v88 _dafny.Set
				_ = _1511_v88
				_1511_v88 = _dafny.SetOf(_1499_v76)
				var _1512_v89 D13
				_ = _1512_v89
				_1512_v89 = Companion_D13_.Create_DC46_((_this).F24(), _1510_v87, _1447_v50, _1511_v88)
				var _1513_v90 _dafny.Sequence
				_ = _1513_v90
				_1513_v90 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v84, (_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool)), _1508_v85)
				var _1514_v91 _dafny.Array
				_ = _1514_v91
				var _nwElement0_44 _dafny.Map = _1508_v85
				_ = _nwElement0_44
				var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(17))
				_ = _nw211
				(_nw211).ArraySet1(_nwElement0_44, 0)
				(_nw211).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v76, false), 1)
				(_nw211).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v76, Companion_Default___.Fm15((_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool), globalState))).Merge(_1508_v85), 2)
				(_nw211).ArraySet1(_1508_v85, 3)
				(_nw211).ArraySet1(_1508_v85, 4)
				(_nw211).ArraySet1(_1508_v85, 5)
				(_nw211).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v76, (_1509_v86).Dtor_cf57())).Merge(_1508_v85), 6)
				(_nw211).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1499_v76, (_this).F24()), 7)
				(_nw211).ArraySet1(_1508_v85, 8)
				(_nw211).ArraySet1(_1508_v85, 9)
				(_nw211).ArraySet1(_1508_v85, 10)
				(_nw211).ArraySet1((func() _dafny.Map {
					if (_1512_v89).Dtor_cf73() {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v84, (_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool))
					}
					return (_1513_v90).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1505_v82).Cardinality()), _dafny.IntOfUint32((_1513_v90).Cardinality()))).Uint32()).(_dafny.Map)
				})(), 11)
				(_nw211).ArraySet1((_1508_v85).Merge(_1508_v85), 12)
				(_nw211).ArraySet1(_1508_v85, 13)
				(_nw211).ArraySet1(_1508_v85, 14)
				(_nw211).ArraySet1(_1508_v85, 15)
				(_nw211).ArraySet1((_1508_v85).Update(_1499_v76, p0), 16)
				_1514_v91 = _nw211
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1514_v91), 0))
				_ = _index217
				(_1514_v91).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v84, (_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool)), (_index217).Int())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1514_v91), 0))
				_ = _index218
				(_1514_v91).ArraySet1(_1508_v85, (_index218).Int())
				var _1515_v92 _dafny.Sequence
				_ = _1515_v92
				_1515_v92 = _dafny.SeqOf(_1499_v76)
				_1499_v76 = (_1515_v92).Select((Companion_Default___.SafeIndex(_1447_v50, _dafny.IntOfUint32((_1515_v92).Cardinality()))).Uint32()).(_dafny.Array)
			}
			var _1516_v93 _dafny.Sequence
			_ = _1516_v93
			_1516_v93 = _dafny.SeqOf((_1460_v59).F30())
			_1452_v53 = (_1452_v53).Update(!(((_1460_v59).F30()).Cmp((_1516_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1516_v93).Cardinality()), _dafny.IntOfUint32((_1516_v93).Cardinality()))).Uint32()).(_dafny.Int)) >= 0), (_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool))
			var _1517_v94 _dafny.Map
			_ = _1517_v94
			_1517_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1456_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1456_v57), 0))).Int()).(bool), (_1460_v59).F30())
			var _1518_v95 D9
			_ = _1518_v95
			_1518_v95 = Companion_D9_.Create_DC30_((Companion_D9_.Create_DC30_(_1517_v94)).Dtor_cf50())
			_1518_v95 = _1518_v95
		} else {
			var _1519_v96 _dafny.Sequence
			_ = _1519_v96
			_1519_v96 = _dafny.SeqOf((_1447_v50).Cmp(_1447_v50) != 0)
			if (_1519_v96).Select((Companion_Default___.SafeIndex(_1447_v50, _dafny.IntOfUint32((_1519_v96).Cardinality()))).Uint32()).(bool) {
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_34
				var _nw212 _dafny.Array
				_ = _nw212
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw212 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) bool = func(_1520_i10 _dafny.Int) bool {
						return (_this).F24()
					}
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw212 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw212).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw212).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				r1 = _nw212
				var _1521_v97 *C1
				_ = _1521_v97
				var _nw213 *C1 = New_C1_()
				_ = _nw213
				_nw213.Ctor__(_1447_v50, (_this).F22())
				_1521_v97 = _nw213
				var _1522_v98 _dafny.Array
				_ = _1522_v98
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_35
				var _nw214 _dafny.Array
				_ = _nw214
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw214 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.Int = (func(_1523_v97 *C1) func(_dafny.Int) _dafny.Int {
						return func(_1524_i11 _dafny.Int) _dafny.Int {
							return (_1524_i11).Plus((_1523_v97).F30())
						}
					})(_1521_v97)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw214 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw214).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw214).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1522_v98 = _nw214
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _index219
				(_1522_v98).ArraySet1((_1521_v97).F30(), (_index219).Int())
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _index220
				(_1522_v98).ArraySet1(Companion_Default___.Fm7((_1521_v97).F30(), globalState), (_index220).Int())
				var _1525_v99 _dafny.CodePoint
				_ = _1525_v99
				_1525_v99 = _dafny.CodePoint('y')
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _index221
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _index222
				var _rhs240 _dafny.Int = (_1521_v97).F30()
				_ = _rhs240
				var _rhs241 _dafny.CodePoint = Companion_Default___.Fm38(true, _1447_v50, globalState)
				_ = _rhs241
				var _rhs242 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_1447_v50, _dafny.IntOfUint32((_1519_v96).Cardinality())), (_1447_v50).Times((_1521_v97).F30()))
				_ = _rhs242
				var _rhs243 _dafny.Int = ((_1521_v97).F30()).Times(_1447_v50)
				_ = _rhs243
				var _rhs244 _dafny.Int = (_1522_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))).Int()).(_dafny.Int)
				_ = _rhs244
				var _lhs203 _dafny.Array = _1522_v98
				_ = _lhs203
				var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _lhs204
				var _lhs205 *GlobalState = globalState
				_ = _lhs205
				var _lhs206 _dafny.Array = _1522_v98
				_ = _lhs206
				var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _lhs207
				(_lhs203).ArraySet1(_rhs240, (_lhs204).Int())
				_1525_v99 = _rhs241
				_lhs205.F6 = _rhs242
				r3 = _rhs243
				(_lhs206).ArraySet1(_rhs244, (_lhs207).Int())
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1522_v98), 0))
				_ = _index223
				(_1522_v98).ArraySet1(_1447_v50, (_index223).Int())
			} else {
				var _1526_v100 D2
				_ = _1526_v100
				_1526_v100 = Companion_D2_.Create_DC7_(_dafny.SeqOf(_1447_v50), (_this).F24(), !((_this).F24()), _1447_v50)
				var _1527_v101 _dafny.Set
				_ = _1527_v101
				_1527_v101 = _dafny.SetOf(_1526_v100, _1526_v100, _1526_v100)
				(globalState).F10 = (_1527_v101).Equals(_1527_v101)
				var _1528_v102 _dafny.Array
				_ = _1528_v102
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_36
				var _nw215 _dafny.Array
				_ = _nw215
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw215 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) bool = (func(_1529_p0 bool) func(_dafny.Int) bool {
						return func(_1530_i12 _dafny.Int) bool {
							return _1529_p0
						}
					})(p0)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw215 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw215).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw215).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1528_v102 = _nw215
				var _1531_v103 T0
				_ = _1531_v103
				var _nw216 *C2 = New_C2_()
				_ = _nw216
				_nw216.Ctor__(_1447_v50, _dafny.SeqOf(_1528_v102))
				_1531_v103 = _nw216
				var _1532_v104 _dafny.Map
				_ = _1532_v104
				_1532_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(584))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg89 _dafny.Int) interface{} {
						return coer89(arg89)
					}
				}((func(_1533_v50 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1534_i13 _dafny.Int) _dafny.Int {
						return _1533_v50
					}
				})(_1447_v50)))).Cardinality()))
				var _1535_v105 _dafny.Array
				_ = _1535_v105
				var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw217
				_1535_v105 = _nw217
				var _1536_v106 D13
				_ = _1536_v106
				_1536_v106 = Companion_D13_.Create_DC46_(p0, _1531_v103, (_1532_v104).Cardinality(), _dafny.SetOf(_1535_v105))
				var _1537_v107 D13
				_ = _1537_v107
				_1537_v107 = Companion_D13_.Create_DC46_((_this).F24(), (_1536_v106).Dtor_cf74(), _1447_v50, _dafny.SetOf(_1535_v105, _1535_v105, _1535_v105, _1535_v105))
				var _1538_v108 _dafny.MultiSet
				_ = _1538_v108
				_1538_v108 = _dafny.MultiSetOf(_1447_v50)
				var _1539_v111 _dafny.Set
				_ = _1539_v111
				_1539_v111 = _dafny.SetOf(func() _dafny.Set {
					var _coll61 = _dafny.NewBuilder()
					_ = _coll61
					for _iter65 := _dafny.Iterate((func() _dafny.Set {
						var _coll62 = _dafny.NewBuilder()
						_ = _coll62
						for _iter66 := _dafny.Iterate((_1538_v108).Elements()); ; {
							_compr_62, _ok66 := _iter66()
							if !_ok66 {
								break
							}
							var _1540_v109 _dafny.Int
							_1540_v109 = interface{}(_compr_62).(_dafny.Int)
							if (_1538_v108).Contains(_1540_v109) {
								_coll62.Add((_1540_v109).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eelqwwl")).Cardinality())))
							}
						}
						return _coll62.ToSet()
					}()).Elements()); ; {
						_compr_61, _ok65 := _iter65()
						if !_ok65 {
							break
						}
						var _1541_v110 _dafny.Int
						_1541_v110 = interface{}(_compr_61).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll63 = _dafny.NewBuilder()
							_ = _coll63
							for _iter67 := _dafny.Iterate((_1538_v108).Elements()); ; {
								_compr_63, _ok67 := _iter67()
								if !_ok67 {
									break
								}
								var _1542_v109 _dafny.Int
								_1542_v109 = interface{}(_compr_63).(_dafny.Int)
								if (_1538_v108).Contains(_1542_v109) {
									_coll63.Add((_1542_v109).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eelqwwl")).Cardinality())))
								}
							}
							return _coll63.ToSet()
						}()).Contains(_1541_v110) {
							_coll61.Add((_1541_v110).Times(_dafny.IntOfInt64(274)))
						}
					}
					return _coll61.ToSet()
				}())
				var _1543_v112 _dafny.Set
				_ = _1543_v112
				_1543_v112 = _dafny.SetOf(_1535_v105, _1535_v105, _1535_v105)
				var _1544_v113 _dafny.Map
				_ = _1544_v113
				_1544_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D13_.Create_DC46_(p0, _1531_v103, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1447_v50)).Cardinality(), _1543_v112), _dafny.IntOfInt64(718))
				var _1545_v114 _dafny.Sequence
				_ = _1545_v114
				_1545_v114 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1537_v107, (_1539_v111).Cardinality()), (_1544_v113).Update(_1537_v107, _dafny.IntOfInt64(-263)))
				(globalState).F11 = _dafny.IntOfUint32((_1545_v114).Cardinality())
				var _1546_v115 _dafny.Sequence
				_ = _1546_v115
				_1546_v115 = _dafny.UnicodeSeqOfUtf8Bytes("phpv")
				_1546_v115 = _1546_v115
				var _1547_v116 _dafny.CodePoint
				_ = _1547_v116
				_1547_v116 = _dafny.CodePoint('y')
				var _1548_v117 _dafny.Map
				_ = _1548_v117
				_1548_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, _1547_v116)
				var _1549_v119 _dafny.Map
				_ = _1549_v119
				_1549_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, !(_1548_v117).Equals((func() _dafny.Map {
					var _coll64 = _dafny.NewMapBuilder()
					_ = _coll64
					for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(667), _dafny.IntOfInt64(816))); ; {
						_compr_64, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1550_v118 _dafny.Int
						_1550_v118 = interface{}(_compr_64).(_dafny.Int)
						if ((_dafny.IntOfInt64(667)).Cmp(_1550_v118) <= 0) && ((_1550_v118).Cmp(_dafny.IntOfInt64(816)) < 0) {
							_coll64.Add((_1550_v118).Times((_1532_v104).Cardinality()), _1547_v116)
						}
					}
					return _coll64.ToMap()
				}()).Update(_1447_v50, _dafny.CodePoint('n'))))
				var _1551_v120 _dafny.Sequence
				_ = _1551_v120
				_1551_v120 = _dafny.SeqOf(_1447_v50)
				var _1552_v121 D13
				_ = _1552_v121
				_1552_v121 = Companion_D13_.Create_DC45_(_dafny.IntOfUint32((_1551_v120).Cardinality()))
				var _pat_let_tv57 = _1447_v50
				_ = _pat_let_tv57
				_1549_v119 = (_1549_v119).Update(_1447_v50, (_1447_v50).Cmp((func(_pat_let28_0 D13) D13 {
					return func(_1553_dt__update__tmp_h1 D13) D13 {
						return func(_pat_let29_0 _dafny.Int) D13 {
							return func(_1554_dt__update_hcf72_h0 _dafny.Int) D13 {
								return Companion_D13_.Create_DC45_(_1554_dt__update_hcf72_h0)
							}(_pat_let29_0)
						}(_pat_let_tv57)
					}(_pat_let28_0)
				}(_1552_v121)).Dtor_cf72()) == 0)
				(globalState).F16 = _1447_v50
			}
			var _1555_v122 D10
			_ = _1555_v122
			_1555_v122 = Companion_D10_.Create_DC35_((_dafny.Zero).Minus(_1447_v50), _1447_v50)
			var _source23 D10 = _1555_v122
			_ = _source23
			if _source23.Is_DC34() {
				var _1556___mcc_h13 _dafny.CodePoint = _source23.Get_().(D10_DC34).Cf55
				_ = _1556___mcc_h13
				var _1557___mcc_h14 bool = _source23.Get_().(D10_DC34).Cf56
				_ = _1557___mcc_h14
				var _1558___mcc_h15 bool = _source23.Get_().(D10_DC34).Cf57
				_ = _1558___mcc_h15
				var _1559___mcc_h16 bool = _source23.Get_().(D10_DC34).Cf58
				_ = _1559___mcc_h16
				var _1560___mcc_h17 bool = _source23.Get_().(D10_DC34).Cf59
				_ = _1560___mcc_h17
				var _1561_cf59 bool = _1560___mcc_h17
				_ = _1561_cf59
				var _1562_cf58 bool = _1559___mcc_h16
				_ = _1562_cf58
				var _1563_cf57 bool = _1558___mcc_h15
				_ = _1563_cf57
				var _1564_cf56 bool = _1557___mcc_h14
				_ = _1564_cf56
				var _1565_cf55 _dafny.CodePoint = _1556___mcc_h13
				_ = _1565_cf55
				var _1566_v123 _dafny.Sequence
				_ = _1566_v123
				_1566_v123 = _dafny.UnicodeSeqOfUtf8Bytes("n")
				var _1567_v124 _dafny.Map
				_ = _1567_v124
				_1567_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_1519_v96, (Companion_Default___.SafeIndex(_1447_v50, _dafny.IntOfUint32((_1519_v96).Cardinality()))).Uint32(), _1561_cf59), _1519_v96), (_1447_v50).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1566_v123).Cardinality()))))
				var _rhs245 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1447_v50), _1447_v50))))
				_ = _rhs245
				var _rhs246 bool = !(_dafny.Companion_Sequence_.IsPrefixOf(_1566_v123, _1566_v123))
				_ = _rhs246
				var _rhs247 _dafny.Map = ((_1567_v124).Merge(_1567_v124)).Merge(_1567_v124)
				_ = _rhs247
				var _lhs208 *GlobalState = globalState
				_ = _lhs208
				_lhs208.F8 = _rhs245
				_1562_cf58 = _rhs246
				_1567_v124 = _rhs247
				var _1568_v125 _dafny.Array
				_ = _1568_v125
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_37
				var _nw218 _dafny.Array
				_ = _nw218
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw218 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.Int = (func(_1569_v50 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1570_i14 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1570_i14, _1569_v50)
						}
					})(_1447_v50)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw218 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw218).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw218).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1568_v125 = _nw218
				var _nwElement0_45 _dafny.Int = _dafny.IntOfInt64(-204)
				_ = _nwElement0_45
				var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.One)
				_ = _nw219
				(_nw219).ArraySet1(_nwElement0_45, 0)
				_1568_v125 = _nw219
				(globalState).F16 = _1447_v50
				r3 = _1447_v50
			} else if _source23.Is_DC35() {
				var _1571___mcc_h18 _dafny.Int = _source23.Get_().(D10_DC35).Cf60
				_ = _1571___mcc_h18
				var _1572___mcc_h19 _dafny.Int = _source23.Get_().(D10_DC35).Cf61
				_ = _1572___mcc_h19
				var _1573_cf61 _dafny.Int = _1572___mcc_h19
				_ = _1573_cf61
				var _1574_cf60 _dafny.Int = _1571___mcc_h18
				_ = _1574_cf60
				var _1575_v126 _dafny.CodePoint
				_ = _1575_v126
				_1575_v126 = _dafny.CodePoint('l')
				var _1576_v127 _dafny.Sequence
				_ = _1576_v127
				_1576_v127 = _dafny.SeqOf(_1447_v50)
				var _1577_v128 D0
				_ = _1577_v128
				_1577_v128 = Companion_D0_.Create_DC1_(_1573_cf61, p0, Companion_Default___.Fm47(_1575_v126, false, globalState), _dafny.IntOfUint32((_1576_v127).Cardinality()), _1575_v126)
				var _1578_v129 D4
				_ = _1578_v129
				_1578_v129 = Companion_D4_.Create_DC13_(_1577_v128)
				var _1579_v130 _dafny.Sequence
				_ = _1579_v130
				_1579_v130 = _dafny.SeqOf(_1578_v129)
				var _1580_v131 _dafny.Map
				_ = _1580_v131
				_1580_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1579_v130).Cardinality()), p0)
				var _1581_v132 _dafny.Sequence
				_ = _1581_v132
				_1581_v132 = _dafny.UnicodeSeqOfUtf8Bytes("kmopkrvet")
				var _rhs248 bool = (!(p0)) || (Companion_Default___.Fm15((_this).F24(), globalState))
				_ = _rhs248
				var _rhs249 bool = Companion_Default___.Fm37(_dafny.SeqOf(p0, p0, p0), _1580_v131, _1581_v132, globalState)
				_ = _rhs249
				var _rhs250 _dafny.Int = _1573_cf61
				_ = _rhs250
				var _lhs209 *GlobalState = globalState
				_ = _lhs209
				var _lhs210 *GlobalState = globalState
				_ = _lhs210
				r0 = _rhs248
				_lhs209.F0 = _rhs249
				_lhs210.F16 = _rhs250
				var _1582_v133 _dafny.Map
				_ = _1582_v133
				_1582_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1581_v132).Cardinality()))
				var _1583_v134 D9
				_ = _1583_v134
				_1583_v134 = Companion_D9_.Create_DC30_(_1582_v133)
				var _1584_v135 _dafny.Sequence
				_ = _1584_v135
				_1584_v135 = _dafny.SeqOf(_1582_v133, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_1576_v127).Cardinality())))
				var _pat_let_tv58 = _1584_v135
				_ = _pat_let_tv58
				var _pat_let_tv59 = _1573_cf61
				_ = _pat_let_tv59
				var _pat_let_tv60 = _1584_v135
				_ = _pat_let_tv60
				var _1585_v136 _dafny.Map
				_ = _1585_v136
				_1585_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1581_v132, func(_pat_let30_0 D9) D9 {
					return func(_1586_dt__update__tmp_h2 D9) D9 {
						return func(_pat_let31_0 _dafny.Map) D9 {
							return func(_1587_dt__update_hcf50_h0 _dafny.Map) D9 {
								return Companion_D9_.Create_DC30_(_1587_dt__update_hcf50_h0)
							}(_pat_let31_0)
						}((_pat_let_tv58).Select((Companion_Default___.SafeIndex(_pat_let_tv59, _dafny.IntOfUint32((_pat_let_tv60).Cardinality()))).Uint32()).(_dafny.Map))
					}(_pat_let30_0)
				}(_1583_v134))
				_1585_v136 = _1585_v136
				var _1588_v137 _dafny.Array
				_ = _1588_v137
				var _nw220 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw220
				_1588_v137 = _nw220
				r1 = _1588_v137
				var _1589_v138 _dafny.Array
				_ = _1589_v138
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_38
				var _nw221 _dafny.Array
				_ = _nw221
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw221 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Int = (func(_1590_cf61 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1591_i15 _dafny.Int) _dafny.Int {
							return (_1591_i15).Times((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1590_cf61, _dafny.IntOfInt64(43))).Cardinality()))
						}
					})(_1573_cf61)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw221 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw221).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw221).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1589_v138 = _nw221
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_1589_v138), 0))
				_ = _index224
				(_1589_v138).ArraySet1(_dafny.IntOfInt64(-623), (_index224).Int())
				var _1592_v139 _dafny.MultiSet
				_ = _1592_v139
				_1592_v139 = _dafny.MultiSetOf(_1447_v50)
				var _1593_v140 _dafny.MultiSet
				_ = _1593_v140
				_1593_v140 = _dafny.MultiSetOf((_this).F24(), (_this).F24(), (_this).F24(), p0, p0)
				var _1594_v141 _dafny.Set
				_ = _1594_v141
				_1594_v141 = _dafny.SetOf(_dafny.IntOfInt64(726), _dafny.IntOfInt64(-786))
				var _1595_v142 _dafny.Sequence
				_ = _1595_v142
				_1595_v142 = _dafny.SeqOf(Companion_Default___.Fm23(_1447_v50, _1447_v50, p0, p0, globalState), _1594_v141)
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_1589_v138), 0))
				_ = _index225
				var _rhs251 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1581_v132, _dafny.UnicodeSeqOfUtf8Bytes("q")), _dafny.Companion_Sequence_.Concatenate(_1581_v132, _1581_v132))
				_ = _rhs251
				var _rhs252 bool = (_this).F24()
				_ = _rhs252
				var _rhs253 _dafny.Int = (_this).Fm0(_1592_v139, (func() _dafny.Int {
					if p0 {
						return _1447_v50
					}
					return (_dafny.Zero).Minus((_1593_v140).Cardinality())
				})(), _1595_v142, (_this).F24(), globalState)
				_ = _rhs253
				var _rhs254 _dafny.Int = (_1576_v127).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1573_cf61), _dafny.IntOfUint32((_1576_v127).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs254
				var _lhs211 *GlobalState = globalState
				_ = _lhs211
				var _lhs212 _dafny.Array = _1589_v138
				_ = _lhs212
				var _lhs213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_1589_v138), 0))
				_ = _lhs213
				_1581_v132 = _rhs251
				_lhs211.F10 = _rhs252
				r3 = _rhs253
				(_lhs212).ArraySet1(_rhs254, (_lhs213).Int())
			} else {
				var _1596___mcc_h20 _dafny.Map = _source23.Get_().(D10_DC33).Cf54
				_ = _1596___mcc_h20
				var _1597_cf54 _dafny.Map = _1596___mcc_h20
				_ = _1597_cf54
				var _1598_v143 _dafny.Map
				_ = _1598_v143
				_1598_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1599_v144 _dafny.Map
				_ = _1599_v144
				_1599_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1381_v1, _1598_v143)
				var _1600_v145 _dafny.Sequence
				_ = _1600_v145
				_1600_v145 = _dafny.SeqOf(_1599_v144)
				var _1601_v146 _dafny.Sequence
				_ = _1601_v146
				_1601_v146 = _dafny.SeqOf(_1447_v50, _dafny.IntOfInt64(306))
				_1599_v144 = (_1600_v145).Select((Companion_Default___.SafeIndex((_1601_v146).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1447_v50), _dafny.IntOfUint32((_1601_v146).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1600_v145).Cardinality()))).Uint32()).(_dafny.Map)
				(globalState).F11 = _1447_v50
				var _1602_v147 _dafny.MultiSet
				_ = _1602_v147
				_1602_v147 = _dafny.MultiSetOf(_1447_v50)
				(globalState).F0 = (_dafny.MultiSetOf(_1447_v50)).IsProperSubsetOf(_1602_v147)
				var _1603_v148 _dafny.Array
				_ = _1603_v148
				var _nw222 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw222
				_1603_v148 = _nw222
				var _1604_v149 _dafny.CodePoint
				_ = _1604_v149
				_1604_v149 = _dafny.CodePoint('s')
				var _1605_v150 _dafny.Map
				_ = _1605_v150
				_1605_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1603_v148, _1604_v149)
				_1605_v150 = (_1605_v150).Update(_1603_v148, _1604_v149)
			}
			var _1606_v151 _dafny.Array
			_ = _1606_v151
			var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
			_ = _nw223
			_1606_v151 = _nw223
			var _1607_v152 _dafny.Array
			_ = _1607_v152
			var _nw224 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
			_ = _nw224
			_1607_v152 = _nw224
			var _1608_v153 _dafny.Map
			_ = _1608_v153
			_1608_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, p0)
			var _1609_v154 _dafny.Sequence
			_ = _1609_v154
			_1609_v154 = _dafny.UnicodeSeqOfUtf8Bytes("lntn")
			var _1610_v155 _dafny.Map
			_ = _1610_v155
			_1610_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, Companion_Default___.Fm37(_1519_v96, _1608_v153, _1609_v154, globalState))
			var _1611_v156 _dafny.Sequence
			_ = _1611_v156
			_1611_v156 = _dafny.SeqOf(_1610_v155)
			var _rhs255 bool = !(p0)
			_ = _rhs255
			var _rhs256 _dafny.Array = _1606_v151
			_ = _rhs256
			var _rhs257 _dafny.Array = _1607_v152
			_ = _rhs257
			var _rhs258 _dafny.Sequence = _1611_v156
			_ = _rhs258
			var _rhs259 bool = p0
			_ = _rhs259
			var _lhs214 *GlobalState = globalState
			_ = _lhs214
			var _lhs215 *GlobalState = globalState
			_ = _lhs215
			_lhs214.F10 = _rhs255
			_1606_v151 = _rhs256
			_1607_v152 = _rhs257
			_1611_v156 = _rhs258
			_lhs215.F10 = _rhs259
			var _1612_v157 _dafny.MultiSet
			_ = _1612_v157
			_1612_v157 = _dafny.MultiSetOf(_1609_v154)
			var _1613_v158 D10
			_ = _1613_v158
			_1613_v158 = Companion_D10_.Create_DC34_(Companion_Default___.Fm38(p0, _1447_v50, globalState), p0, p0, !((_1612_v157).IsSubsetOf(_1612_v157)), p0)
			var _source24 D10 = _1613_v158
			_ = _source24
			if _source24.Is_DC34() {
				var _1614___mcc_h21 _dafny.CodePoint = _source24.Get_().(D10_DC34).Cf55
				_ = _1614___mcc_h21
				var _1615___mcc_h22 bool = _source24.Get_().(D10_DC34).Cf56
				_ = _1615___mcc_h22
				var _1616___mcc_h23 bool = _source24.Get_().(D10_DC34).Cf57
				_ = _1616___mcc_h23
				var _1617___mcc_h24 bool = _source24.Get_().(D10_DC34).Cf58
				_ = _1617___mcc_h24
				var _1618___mcc_h25 bool = _source24.Get_().(D10_DC34).Cf59
				_ = _1618___mcc_h25
				var _1619_cf59 bool = _1618___mcc_h25
				_ = _1619_cf59
				var _1620_cf58 bool = _1617___mcc_h24
				_ = _1620_cf58
				var _1621_cf57 bool = _1616___mcc_h23
				_ = _1621_cf57
				var _1622_cf56 bool = _1615___mcc_h22
				_ = _1622_cf56
				var _1623_cf55 _dafny.CodePoint = _1614___mcc_h21
				_ = _1623_cf55
				var _1624_v159 *C0
				_ = _1624_v159
				var _nw225 *C0 = New_C0_()
				_ = _nw225
				_nw225.Ctor__()
				_1624_v159 = _nw225
				var _1625_v160 *C3
				_ = _1625_v160
				var _nw226 *C3 = New_C3_()
				_ = _nw226
				_nw226.Ctor__((_this).F22())
				_1625_v160 = _nw226
				var _1626_v161 _dafny.MultiSet
				_ = _1626_v161
				_1626_v161 = _dafny.MultiSetOf(_1625_v160)
				var _1627_v162 _dafny.Map
				_ = _1627_v162
				_1627_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, _1626_v161)
				var _1628_v163 *C1
				_ = _1628_v163
				var _nw227 *C1 = New_C1_()
				_ = _nw227
				_nw227.Ctor__((_1627_v162).Cardinality(), _dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()))
				_1628_v163 = _nw227
				var _1629_v164 _dafny.Map
				_ = _1629_v164
				_1629_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1619_cf59, _dafny.IntOfInt64(-861))
				var _1630_v165 _dafny.Map
				_ = _1630_v165
				_1630_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1619_cf59, _1625_v160)
				var _1631_v166 _dafny.Set
				_ = _1631_v166
				_1631_v166 = _dafny.SetOf(_1630_v165, _1630_v165, _1630_v165, (_1630_v165).Update(p0, _1625_v160), (_1630_v165).Update(_1619_cf59, _1625_v160))
				var _1632_v167 _dafny.Sequence
				_ = _1632_v167
				_1632_v167 = _dafny.SeqOf(_dafny.MultiSetOf(_1619_cf59))
				var _rhs260 _dafny.Int = (_dafny.Zero).Minus(_1447_v50)
				_ = _rhs260
				var _rhs261 _dafny.Int = (_1628_v163).F30()
				_ = _rhs261
				var _rhs262 _dafny.Map = _1629_v164
				_ = _rhs262
				var _rhs263 _dafny.Set = _dafny.SetOf((_1630_v165).Merge(_1630_v165), _1630_v165, _1630_v165)
				_ = _rhs263
				var _rhs264 _dafny.Sequence = _1632_v167
				_ = _rhs264
				var _lhs216 *GlobalState = globalState
				_ = _lhs216
				_1447_v50 = _rhs260
				_lhs216.F16 = _rhs261
				_1629_v164 = _rhs262
				_1631_v166 = _rhs263
				_1632_v167 = _rhs264
				(globalState).F16 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_1628_v163).Fm34(true, globalState)))
			} else if _source24.Is_DC35() {
				var _1633___mcc_h26 _dafny.Int = _source24.Get_().(D10_DC35).Cf60
				_ = _1633___mcc_h26
				var _1634___mcc_h27 _dafny.Int = _source24.Get_().(D10_DC35).Cf61
				_ = _1634___mcc_h27
				var _1635_cf61 _dafny.Int = _1634___mcc_h27
				_ = _1635_cf61
				var _1636_cf60 _dafny.Int = _1633___mcc_h26
				_ = _1636_cf60
				var _1637_v168 _dafny.Array
				_ = _1637_v168
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw228
				_1637_v168 = _nw228
				var _1638_v169 _dafny.Array
				_ = _1638_v169
				var _nw229 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw229
				_1638_v169 = _nw229
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_1638_v169), 0))
				_ = _index226
				(_1638_v169).ArraySet1((_this).F24(), (_index226).Int())
				var _1639_v170 _dafny.Set
				_ = _1639_v170
				_1639_v170 = _dafny.SetOf(p0)
				var _1640_v171 _dafny.MultiSet
				_ = _1640_v171
				_1640_v171 = _dafny.MultiSetOf(_1639_v170, _1639_v170)
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_1638_v169), 0))
				_ = _index227
				var _rhs265 _dafny.Array = _1637_v168
				_ = _rhs265
				var _rhs266 _dafny.Int = _1447_v50
				_ = _rhs266
				var _rhs267 bool = (_1640_v171).IsSubsetOf(_1640_v171)
				_ = _rhs267
				var _rhs268 bool = !((_this).F24())
				_ = _rhs268
				var _lhs217 *GlobalState = globalState
				_ = _lhs217
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				var _lhs219 _dafny.Array = _1638_v169
				_ = _lhs219
				var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_1638_v169), 0))
				_ = _lhs220
				_1637_v168 = _rhs265
				_lhs217.F16 = _rhs266
				_lhs218.F10 = _rhs267
				(_lhs219).ArraySet1(_rhs268, (_lhs220).Int())
				(globalState).F5 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1519_v96, _1519_v96), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1638_v169).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_1638_v169), 0))).Int()).(bool), p0, (_1638_v169).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_1638_v169), 0))).Int()).(bool), (_this).F24(), (_this).F24()), _1519_v96))
				var _1641_v172 _dafny.Sequence
				_ = _1641_v172
				_1641_v172 = _dafny.SeqOf(_dafny.SeqOf(Companion_Default___.Fm15((_this).F24(), globalState)))
				var _1642_v173 _dafny.MultiSet
				_ = _1642_v173
				_1642_v173 = _dafny.MultiSetOf(p0)
				_1641_v172 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1641_v172, Companion_Default___.Fm48(globalState)), _dafny.SeqOf(_1519_v96)), (Companion_Default___.SafeIndex((_1447_v50).Times((_1642_v173).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1641_v172, Companion_Default___.Fm48(globalState)), _dafny.SeqOf(_1519_v96))).Cardinality()))).Uint32(), _1519_v96)
				var _1643_v174 _dafny.Sequence
				_ = _1643_v174
				_1643_v174 = _dafny.SeqOf(_1635_cf61)
				_1643_v174 = _1643_v174
			} else {
				var _1644___mcc_h28 _dafny.Map = _source24.Get_().(D10_DC33).Cf54
				_ = _1644___mcc_h28
				var _1645_cf54 _dafny.Map = _1644___mcc_h28
				_ = _1645_cf54
				var _rhs269 bool = !((func() bool {
					if p0 {
						return p0
					}
					return p0
				})()) || ((_this).Fm3(_1609_v154, p0, (_1519_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.IntOfUint32((_1519_v96).Cardinality()))).Uint32()).(bool), globalState))
				_ = _rhs269
				var _rhs270 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1519_v96, _1519_v96)
				_ = _rhs270
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				r0 = _rhs269
				_lhs221.F10 = _rhs270
				var _1646_v176 _dafny.Array
				_ = _1646_v176
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_39
				var _nw230 _dafny.Array
				_ = _nw230
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw230 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) bool = (func(_1647_v50 _dafny.Int) func(_dafny.Int) bool {
						return func(_1648_i16 _dafny.Int) bool {
							return (Companion_D5_.Create_DC19_(func() _dafny.Map {
								var _coll65 = _dafny.NewMapBuilder()
								_ = _coll65
								for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(205), _dafny.IntOfInt64(829))); ; {
									_compr_65, _ok69 := _iter69()
									if !_ok69 {
										break
									}
									var _1649_v175 _dafny.Int
									_1649_v175 = interface{}(_compr_65).(_dafny.Int)
									if ((_dafny.IntOfInt64(205)).Cmp(_1649_v175) <= 0) && ((_1649_v175).Cmp(_dafny.IntOfInt64(829)) < 0) {
										_coll65.Add((_1649_v175).Times(_1647_v50), _1647_v50)
									}
								}
								return _coll65.ToMap()
							}(), true)).Dtor_cf33()
						}
					})(_1447_v50)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw230 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw230).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw230).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1646_v176 = _nw230
				var _1650_v177 _dafny.Map
				_ = _1650_v177
				_1650_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, _1646_v176)
				var _1651_v178 _dafny.Sequence
				_ = _1651_v178
				_1651_v178 = _dafny.SeqOf((_1650_v177).Cardinality())
				var _1652_v179 _dafny.Map
				_ = _1652_v179
				_1652_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1447_v50)
				var _1653_v180 D2
				_ = _1653_v180
				_1653_v180 = Companion_D2_.Create_DC7_(_1651_v178, (_this).F24(), true, (_1652_v179).Cardinality())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1646_v176), 0))
				_ = _index228
				(_1646_v176).ArraySet1((_1653_v180).Dtor_cf11(), (_index228).Int())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1646_v176), 0))
				_ = _index229
				(_1646_v176).ArraySet1((_1519_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if false {
						return _dafny.UnicodeSeqOfUtf8Bytes("on")
					}
					return _1609_v154
				})()).Cardinality()), _dafny.IntOfUint32((_1519_v96).Cardinality()))).Uint32()).(bool), (_index229).Int())
				_1609_v154 = _dafny.UnicodeSeqOfUtf8Bytes("hrcsc")
				var _1654_v181 _dafny.Map
				_ = _1654_v181
				_1654_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1609_v154).Cardinality()), _dafny.IntOfInt64(-637))
				var _1655_v182 _dafny.Array
				_ = _1655_v182
				var _nwElement0_46 _dafny.Int = _1447_v50
				_ = _nwElement0_46
				var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(20))
				_ = _nw231
				(_nw231).ArraySet1(_nwElement0_46, 0)
				(_nw231).ArraySet1(_1447_v50, 1)
				(_nw231).ArraySet1(_dafny.IntOfUint32((_1609_v154).Cardinality()), 2)
				(_nw231).ArraySet1(_1447_v50, 3)
				(_nw231).ArraySet1(_1447_v50, 4)
				(_nw231).ArraySet1(_dafny.IntOfUint32((_1609_v154).Cardinality()), 5)
				(_nw231).ArraySet1(_1447_v50, 6)
				(_nw231).ArraySet1(_1447_v50, 7)
				(_nw231).ArraySet1(_1447_v50, 8)
				(_nw231).ArraySet1(_1447_v50, 9)
				(_nw231).ArraySet1(_dafny.IntOfInt64(605), 10)
				(_nw231).ArraySet1(_dafny.IntOfInt64(413), 11)
				(_nw231).ArraySet1((func() _dafny.Int {
					if (_1654_v181).Contains(Companion_Default___.Fm7(_1447_v50, globalState)) {
						return (_1654_v181).Get(Companion_Default___.Fm7(_1447_v50, globalState)).(_dafny.Int)
					}
					return _1447_v50
				})(), 12)
				(_nw231).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qbq")).Cardinality()), 13)
				(_nw231).ArraySet1(_1447_v50, 14)
				(_nw231).ArraySet1(_1447_v50, 15)
				(_nw231).ArraySet1(_1447_v50, 16)
				(_nw231).ArraySet1(_dafny.IntOfUint32((_1609_v154).Cardinality()), 17)
				(_nw231).ArraySet1(_dafny.IntOfInt64(402), 18)
				(_nw231).ArraySet1(_dafny.IntOfInt64(-262), 19)
				_1655_v182 = _nw231
				var _1656_v183 _dafny.Sequence
				_ = _1656_v183
				_1656_v183 = _dafny.SeqOf(_1655_v182)
				var _1657_v184 _dafny.Array
				_ = _1657_v184
				var _nwElement0_47 _dafny.Array = (_1656_v183).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1447_v50), _dafny.IntOfUint32((_1656_v183).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _nwElement0_47
				var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(29))
				_ = _nw232
				(_nw232).ArraySet1(_nwElement0_47, 0)
				(_nw232).ArraySet1(_1655_v182, 1)
				(_nw232).ArraySet1((func() _dafny.Array {
					if (_1646_v176).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1646_v176), 0))).Int()).(bool) {
						return _1655_v182
					}
					return _1655_v182
				})(), 2)
				(_nw232).ArraySet1(_1655_v182, 3)
				(_nw232).ArraySet1(_1655_v182, 4)
				(_nw232).ArraySet1(_1655_v182, 5)
				(_nw232).ArraySet1(_1655_v182, 6)
				(_nw232).ArraySet1(_1655_v182, 7)
				(_nw232).ArraySet1(_1655_v182, 8)
				(_nw232).ArraySet1(_1655_v182, 9)
				(_nw232).ArraySet1(_1655_v182, 10)
				(_nw232).ArraySet1(_1655_v182, 11)
				(_nw232).ArraySet1(_1655_v182, 12)
				(_nw232).ArraySet1(_1655_v182, 13)
				(_nw232).ArraySet1(_1655_v182, 14)
				(_nw232).ArraySet1(_1655_v182, 15)
				(_nw232).ArraySet1(_1655_v182, 16)
				(_nw232).ArraySet1(_1655_v182, 17)
				(_nw232).ArraySet1(_1655_v182, 18)
				(_nw232).ArraySet1(_1655_v182, 19)
				(_nw232).ArraySet1(_1655_v182, 20)
				(_nw232).ArraySet1(_1655_v182, 21)
				(_nw232).ArraySet1(_1655_v182, 22)
				(_nw232).ArraySet1(_1655_v182, 23)
				(_nw232).ArraySet1(_1655_v182, 24)
				(_nw232).ArraySet1(_1655_v182, 25)
				(_nw232).ArraySet1(_1655_v182, 26)
				(_nw232).ArraySet1(_1655_v182, 27)
				(_nw232).ArraySet1(_1655_v182, 28)
				_1657_v184 = _nw232
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_1657_v184), 0))
				_ = _index230
				(_1657_v184).ArraySet1(_1655_v182, (_index230).Int())
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_1657_v184), 0))
				_ = _index231
				(_1657_v184).ArraySet1(_1655_v182, (_index231).Int())
			}
			var _1658_v185 _dafny.CodePoint
			_ = _1658_v185
			_1658_v185 = _dafny.CodePoint('h')
			var _1659_v186 _dafny.Map
			_ = _1659_v186
			_1659_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v185, p0)
			var _rhs271 _dafny.Map = ((Companion_Default___.Fm49(p0, globalState)).Update(_1658_v185, false)).Merge(_1659_v186)
			_ = _rhs271
			var _rhs272 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1519_v96, _1519_v96), _1519_v96)
			_ = _rhs272
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			_1659_v186 = _rhs271
			_lhs222.F0 = _rhs272
		}
		var _1660_v187 _dafny.Sequence
		_ = _1660_v187
		_1660_v187 = _dafny.UnicodeSeqOfUtf8Bytes("emlcy")
		_1660_v187 = _dafny.UnicodeSeqOfUtf8Bytes("usnjcg")
		if !(!(!_dafny.Companion_Sequence_.Equal(_1660_v187, _1660_v187))) {
			_1660_v187 = Companion_Default___.Fm30(globalState)
			var _1661_v188 _dafny.MultiSet
			_ = _1661_v188
			_1661_v188 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1660_v187).Cardinality()))
			_1661_v188 = (_dafny.MultiSetOf(_1447_v50, _1447_v50, (_dafny.Zero).Minus(Companion_Default___.Fm7(_1447_v50, globalState)), _1447_v50, _1447_v50)).Difference((_1661_v188).Intersection(_1661_v188))
			var _1662_v189 *C4
			_ = _1662_v189
			var _nw233 *C4 = New_C4_()
			_ = _nw233
			_nw233.Ctor__(p0, (_this).F22())
			_1662_v189 = _nw233
			r0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_1660_v187, _dafny.Companion_Sequence_.Concatenate(_1660_v187, _1660_v187))
			var _1663_v190 _dafny.Map
			_ = _1663_v190
			_1663_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('x'))).Cardinality()))
			var _1664_v191 _dafny.Set
			_ = _1664_v191
			_1664_v191 = _dafny.SetOf(!(p0))
			var _1665_v192 _dafny.Set
			_ = _1665_v192
			_1665_v192 = _dafny.SetOf(_1447_v50, (_1663_v190).Cardinality(), (_dafny.Zero).Minus((_1664_v191).Cardinality()))
			var _1666_v194 _dafny.MultiSet
			_ = _1666_v194
			_1666_v194 = _dafny.MultiSetOf((func() _dafny.Set {
				if p0 {
					return _1665_v192
				}
				return func() _dafny.Set {
					var _coll66 = _dafny.NewBuilder()
					_ = _coll66
					for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(965), _dafny.IntOfInt64(848))); ; {
						_compr_66, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _1667_v193 _dafny.Int
						_1667_v193 = interface{}(_compr_66).(_dafny.Int)
						if ((_dafny.IntOfInt64(965)).Cmp(_1667_v193) <= 0) && ((_1667_v193).Cmp(_dafny.IntOfInt64(848)) < 0) {
							_coll66.Add(Companion_Default___.SafeModuloInt(_1667_v193, _1447_v50))
						}
					}
					return _coll66.ToSet()
				}()
			})(), _1665_v192, _1665_v192, (func() _dafny.Set {
				if p0 {
					return _1665_v192
				}
				return _1665_v192
			})())
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1378_v0), 0))
			_ = _index232
			(_1378_v0).ArraySet1(_1660_v187, (_index232).Int())
			var _1668_v196 _dafny.Map
			_ = _1668_v196
			_1668_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1660_v187, (func() _dafny.Set {
				var _coll67 = _dafny.NewBuilder()
				_ = _coll67
				for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(377), _dafny.IntOfInt64(-477))); ; {
					_compr_67, _ok71 := _iter71()
					if !_ok71 {
						break
					}
					var _1669_v195 _dafny.Int
					_1669_v195 = interface{}(_compr_67).(_dafny.Int)
					if ((_dafny.IntOfInt64(377)).Cmp(_1669_v195) <= 0) && ((_1669_v195).Cmp(_dafny.IntOfInt64(-477)) < 0) {
						_coll67.Add(Companion_Default___.SafeDivisionInt(_1669_v195, _1447_v50))
					}
				}
				return _coll67.ToSet()
			}()).Cardinality())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1378_v0), 0))
			_ = _index233
			var _rhs273 _dafny.MultiSet = _1666_v194
			_ = _rhs273
			var _rhs274 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(610))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg90 _dafny.Int) interface{} {
					return coer90(arg90)
				}
			}(func(_1670_i17 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))
			_ = _rhs274
			var _rhs275 bool = _1662_v189.F28
			_ = _rhs275
			var _rhs276 _dafny.Sequence = _1660_v187
			_ = _rhs276
			var _rhs277 _dafny.Int = (func() _dafny.Int {
				if (_this).F24() {
					return ((_1668_v196).Merge(_1668_v196)).Cardinality()
				}
				return (_1447_v50).Times(_dafny.IntOfInt64(841))
			})()
			_ = _rhs277
			var _lhs223 _dafny.Array = _1378_v0
			_ = _lhs223
			var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1378_v0), 0))
			_ = _lhs224
			var _lhs225 *GlobalState = globalState
			_ = _lhs225
			var _lhs226 *GlobalState = globalState
			_ = _lhs226
			_1666_v194 = _rhs273
			(_lhs223).ArraySet1(_rhs274, (_lhs224).Int())
			_lhs225.F0 = _rhs275
			_1660_v187 = _rhs276
			_lhs226.F16 = _rhs277
		} else {
			var _1671_v197 _dafny.CodePoint
			_ = _1671_v197
			_1671_v197 = _dafny.CodePoint('d')
			if !(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("emuwsan"), _1671_v197)) {
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_1378_v0), 0))
				_ = _index234
				(_1378_v0).ArraySet1(_1660_v187, (_index234).Int())
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_1378_v0), 0))
				_ = _index235
				(_1378_v0).ArraySet1(_1660_v187, (_index235).Int())
				var _1672_v198 _dafny.Array
				_ = _1672_v198
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_40
				var _nw234 _dafny.Array
				_ = _nw234
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw234 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Int = (func(_1673_v50 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1674_i18 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1674_i18, _1673_v50)
						}
					})(_1447_v50)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw234 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw234).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw234).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1672_v198 = _nw234
				var _1675_v199 _dafny.Set
				_ = _1675_v199
				_1675_v199 = _dafny.SetOf(_dafny.IntOfInt64(-974))
				var _1676_v200 _dafny.Map
				_ = _1676_v200
				_1676_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_1675_v199).Cardinality()))
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1672_v198), 0))
				_ = _index236
				(_1672_v198).ArraySet1(((_1676_v200).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-823)))).Cardinality(), (_index236).Int())
				var _1677_v201 _dafny.Set
				_ = _1677_v201
				_1677_v201 = _dafny.SetOf(p0, true)
				var _1678_v202 _dafny.Sequence
				_ = _1678_v202
				_1678_v202 = _dafny.SeqOf(_1447_v50, ((_1677_v201).Difference(_1677_v201)).Cardinality(), _1447_v50)
				var _1679_v203 _dafny.Array
				_ = _1679_v203
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_41
				var _nw235 _dafny.Array
				_ = _nw235
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw235 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = (func(_1680_v50 _dafny.Int) func(_dafny.Int) bool {
						return func(_1681_i19 _dafny.Int) bool {
							return (_dafny.IntOfInt64(417)).Cmp(_1680_v50) >= 0
						}
					})(_1447_v50)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw235 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw235).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw235).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1679_v203 = _nw235
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1672_v198), 0))
				_ = _index237
				var _rhs278 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1678_v202).Cardinality()))
				_ = _rhs278
				var _rhs279 _dafny.Array = _1679_v203
				_ = _rhs279
				var _rhs280 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1660_v187, _1660_v187)
				_ = _rhs280
				var _lhs227 _dafny.Array = _1672_v198
				_ = _lhs227
				var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1672_v198), 0))
				_ = _lhs228
				var _lhs229 *GlobalState = globalState
				_ = _lhs229
				(_lhs227).ArraySet1(_rhs278, (_lhs228).Int())
				r1 = _rhs279
				_lhs229.F0 = _rhs280
				var _1682_v204 _dafny.Sequence
				_ = _1682_v204
				_1682_v204 = _dafny.SeqOf(p0, (_this).F24())
				var _1683_v205 _dafny.Sequence
				_ = _1683_v205
				_1683_v205 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1682_v204, _1682_v204))
				_1683_v205 = _1683_v205
				var _1684_v206 _dafny.Array
				_ = _1684_v206
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
				_ = _nw236
				_1684_v206 = _nw236
				var _1685_v207 _dafny.MultiSet
				_ = _1685_v207
				_1685_v207 = _dafny.MultiSetOf((_1672_v198).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1672_v198), 0))).Int()).(_dafny.Int))
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1684_v206), 0))
				_ = _index238
				(_1684_v206).ArraySet1(_1685_v207, (_index238).Int())
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1684_v206), 0))
				_ = _index239
				(_1684_v206).ArraySet1(_1685_v207, (_index239).Int())
				var _1686_v208 D2
				_ = _1686_v208
				_1686_v208 = Companion_D2_.Create_DC6_((_1672_v198).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1672_v198), 0))).Int()).(_dafny.Int))
				var _1687_v209 D2
				_ = _1687_v209
				_1687_v209 = Companion_D2_.Create_DC9_(Companion_D2_.Create_DC9_(_1686_v208))
				var _1688_v210 _dafny.MultiSet
				_ = _1688_v210
				_1688_v210 = _dafny.MultiSetOf(_1687_v209)
				var _1689_v211 *C4
				_ = _1689_v211
				var _nw237 *C4 = New_C4_()
				_ = _nw237
				_nw237.Ctor__(!(_1688_v210).Equals(_1688_v210), (_this).F22())
				_1689_v211 = _nw237
			} else {
				var _1690_v212 D11
				_ = _1690_v212
				_1690_v212 = Companion_D11_.Create_DC39_((_this).F24())
				var _1691_v213 _dafny.Array
				_ = _1691_v213
				var _nwElement0_48 D11 = _1690_v212
				_ = _nwElement0_48
				var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.One)
				_ = _nw238
				(_nw238).ArraySet1(_nwElement0_48, 0)
				_1691_v213 = _nw238
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1691_v213), 0))
				_ = _index240
				(_1691_v213).ArraySet1(Companion_D11_.Create_DC39_((_this).F24()), (_index240).Int())
				var _1692_v214 D1
				_ = _1692_v214
				_1692_v214 = Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-352))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}((func(_1693_v197 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1694_i20 _dafny.Int) _dafny.CodePoint {
						return _1693_v197
					}
				})(_1671_v197))))
				var _1695_v215 _dafny.Sequence
				_ = _1695_v215
				_1695_v215 = _dafny.SeqOf(_1692_v214, func(_pat_let32_0 D1) D1 {
					return func(_1696_dt__update__tmp_h3 D1) D1 {
						return func(_pat_let33_0 _dafny.Sequence) D1 {
							return func(_1697_dt__update_hcf6_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_1697_dt__update_hcf6_h0)
							}(_pat_let33_0)
						}(_dafny.UnicodeSeqOfUtf8Bytes("ujsmsmjoj"))
					}(_pat_let32_0)
				}(Companion_D1_.Create_DC3_(_1660_v187)))
				var _1698_v216 _dafny.Array
				_ = _1698_v216
				var _nwElement0_49 _dafny.CodePoint = _dafny.CodePoint('i')
				_ = _nwElement0_49
				var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(4))
				_ = _nw239
				(_nw239).ArraySet1CodePoint(_nwElement0_49, 0)
				(_nw239).ArraySet1CodePoint(_1671_v197, 1)
				(_nw239).ArraySet1CodePoint((_1660_v187).Select((Companion_Default___.SafeIndex(_1447_v50, _dafny.IntOfUint32((_1660_v187).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
				(_nw239).ArraySet1CodePoint(_1671_v197, 3)
				_1698_v216 = _nw239
				var _1699_v217 _dafny.Sequence
				_ = _1699_v217
				_1699_v217 = _dafny.SeqOf(p0)
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1698_v216), 0))
				_ = _index241
				(_1698_v216).ArraySet1CodePoint((_1660_v187).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1699_v217).Cardinality())), _dafny.IntOfUint32((_1660_v187).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index241).Int())
				var _1700_v218 _dafny.Map
				_ = _1700_v218
				_1700_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1447_v50)
				var _1701_v219 _dafny.Map
				_ = _1701_v219
				_1701_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1700_v218).Cardinality(), _dafny.IntOfInt64(4))
				var _1702_v220 D0
				_ = _1702_v220
				_1702_v220 = Companion_D0_.Create_DC1_(_1447_v50, (p0) || (p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1701_v219), (func() _dafny.Int {
					if p0 {
						return _1447_v50
					}
					return _dafny.IntOfInt64(186)
				})(), _1671_v197)
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1691_v213), 0))
				_ = _index242
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1698_v216), 0))
				_ = _index243
				var _rhs281 _dafny.Int = _1447_v50
				_ = _rhs281
				var _rhs282 D11 = _1690_v212
				_ = _rhs282
				var _rhs283 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(90))).Uint32(), func(coer92 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}((func(_1703_v214 D1) func(_dafny.Int) D1 {
					return func(_1704_i21 _dafny.Int) D1 {
						return _1703_v214
					}
				})(_1692_v214)))
				_ = _rhs283
				var _rhs284 _dafny.CodePoint = _1671_v197
				_ = _rhs284
				var _rhs285 D0 = _1702_v220
				_ = _rhs285
				var _lhs230 _dafny.Array = _1691_v213
				_ = _lhs230
				var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1691_v213), 0))
				_ = _lhs231
				var _lhs232 _dafny.Array = _1698_v216
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1698_v216), 0))
				_ = _lhs233
				r3 = _rhs281
				(_lhs230).ArraySet1(_rhs282, (_lhs231).Int())
				_1695_v215 = _rhs283
				(_lhs232).ArraySet1CodePoint(_rhs284, (_lhs233).Int())
				_1702_v220 = _rhs285
				var _rhs286 _dafny.Int = _1447_v50
				_ = _rhs286
				var _rhs287 _dafny.CodePoint = (_1698_v216).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1698_v216), 0))).Int())
				_ = _rhs287
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				_lhs234.F11 = _rhs286
				_1671_v197 = _rhs287
				var _1705_v221 _dafny.Array
				_ = _1705_v221
				var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
				_ = _nw240
				_1705_v221 = _nw240
				var _1706_v222 T0
				_ = _1706_v222
				var _nw241 *C1 = New_C1_()
				_ = _nw241
				_nw241.Ctor__(_1447_v50, (_this).F22())
				_1706_v222 = _nw241
				var _1707_v223 _dafny.MultiSet
				_ = _1707_v223
				_1707_v223 = _dafny.MultiSetOf(p0)
				var _1708_v224 D14
				_ = _1708_v224
				_1708_v224 = Companion_D14_.Create_DC47_(_1707_v223)
				var _1709_v225 _dafny.Map
				_ = _1709_v225
				_1709_v225 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1706_v222, (_1708_v224).Dtor_cf77())
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_1705_v221), 0))
				_ = _index244
				(_1705_v221).ArraySet1(_1709_v225, (_index244).Int())
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_1705_v221), 0))
				_ = _index245
				(_1705_v221).ArraySet1(((_1709_v225).Merge((_1709_v225).Update(_1706_v222, _1707_v223))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1706_v222, _1707_v223)).Merge(_1709_v225)), (_index245).Int())
				var _1710_v226 _dafny.Array
				_ = _1710_v226
				var _nw242 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
				_ = _nw242
				_1710_v226 = _nw242
				var _1711_v227 D11
				_ = _1711_v227
				_1711_v227 = Companion_D11_.Create_DC37_(_1710_v226)
				var _1712_v228 _dafny.Array
				_ = _1712_v228
				var _nwElement0_50 _dafny.Int = _1447_v50
				_ = _nwElement0_50
				var _nw243 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(10))
				_ = _nw243
				(_nw243).ArraySet1(_nwElement0_50, 0)
				(_nw243).ArraySet1(_dafny.IntOfUint32((_1660_v187).Cardinality()), 1)
				(_nw243).ArraySet1(_1447_v50, 2)
				(_nw243).ArraySet1(_1447_v50, 3)
				(_nw243).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1660_v187, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1660_v187).Cardinality()), _dafny.IntOfUint32((_1660_v187).Cardinality()))).Uint32(), (_1698_v216).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_1698_v216), 0))).Int()))).Cardinality()), 4)
				(_nw243).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1699_v217, (Companion_Default___.SafeIndex((_1707_v223).Cardinality(), _dafny.IntOfUint32((_1699_v217).Cardinality()))).Uint32(), (_this).F24())).Cardinality()), 5)
				(_nw243).ArraySet1(_1447_v50, 6)
				(_nw243).ArraySet1(_1447_v50, 7)
				(_nw243).ArraySet1(_dafny.IntOfInt64(380), 8)
				(_nw243).ArraySet1(_1447_v50, 9)
				_1712_v228 = _nw243
				var _1713_v229 D2
				_ = _1713_v229
				_1713_v229 = Companion_D2_.Create_DC8_(_1692_v214, _1712_v228, p0)
				var _1714_v230 _dafny.Map
				_ = _1714_v230
				_1714_v230 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1711_v227, _1713_v229)
				var _1715_v231 _dafny.Map
				_ = _1715_v231
				_1715_v231 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1447_v50, _1710_v226)
				var _1716_v232 _dafny.Sequence
				_ = _1716_v232
				_1716_v232 = _dafny.SeqOf(_1712_v228)
				var _pat_let_tv61 = _1715_v231
				_ = _pat_let_tv61
				var _pat_let_tv62 = _1716_v232
				_ = _pat_let_tv62
				var _pat_let_tv63 = _1715_v231
				_ = _pat_let_tv63
				var _pat_let_tv64 = _1716_v232
				_ = _pat_let_tv64
				var _pat_let_tv65 = _1710_v226
				_ = _pat_let_tv65
				_1714_v230 = (_1714_v230).Update(func(_pat_let34_0 D11) D11 {
					return func(_1717_dt__update__tmp_h4 D11) D11 {
						return func(_pat_let35_0 _dafny.Array) D11 {
							return func(_1718_dt__update_hcf63_h0 _dafny.Array) D11 {
								return Companion_D11_.Create_DC37_(_1718_dt__update_hcf63_h0)
							}(_pat_let35_0)
						}((func() _dafny.Array {
							if (_pat_let_tv61).Contains(_dafny.IntOfUint32((_pat_let_tv62).Cardinality())) {
								return (_pat_let_tv63).Get(_dafny.IntOfUint32((_pat_let_tv64).Cardinality())).(_dafny.Array)
							}
							return _pat_let_tv65
						})())
					}(_pat_let34_0)
				}(_1711_v227), _1713_v229)
				var _1719_v233 _dafny.Array
				_ = _1719_v233
				var _nw244 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw244
				_1719_v233 = _nw244
				var _1720_v234 _dafny.MultiSet
				_ = _1720_v234
				_1720_v234 = _dafny.MultiSetOf(_1719_v233, _1719_v233, _1719_v233)
				_1720_v234 = ((_1720_v234).Difference(_dafny.MultiSetOf(_1719_v233, _1719_v233))).Intersection(_1720_v234)
			}
			var _1721_v235 _dafny.Array
			_ = _1721_v235
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw245
			_1721_v235 = _nw245
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1721_v235), 0))
			_ = _index246
			(_1721_v235).ArraySet1(p0, (_index246).Int())
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1721_v235), 0))
			_ = _index247
			(_1721_v235).ArraySet1(p0, (_index247).Int())
			var _1722_v236 T0
			_ = _1722_v236
			var _nw246 *C4 = New_C4_()
			_ = _nw246
			_nw246.Ctor__((_1721_v235).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1721_v235), 0))).Int()).(bool), (_this).F22())
			_1722_v236 = _nw246
			var _1723_v237 _dafny.Set
			_ = _1723_v237
			_1723_v237 = _dafny.SetOf(_1447_v50)
			var _1724_v238 _dafny.Sequence
			_ = _1724_v238
			_1724_v238 = _dafny.SeqOf(_1723_v237)
			var _1725_v239 _dafny.Set
			_ = _1725_v239
			_1725_v239 = _dafny.SetOf(true, (_this).F24(), p0)
			var _1726_v240 _dafny.Array
			_ = _1726_v240
			var _nwElement0_51 _dafny.Int = _dafny.IntOfInt64(-876)
			_ = _nwElement0_51
			var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(9))
			_ = _nw247
			(_nw247).ArraySet1(_nwElement0_51, 0)
			(_nw247).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1660_v187).Cardinality())), 1)
			(_nw247).ArraySet1((_dafny.Zero).Minus(_1447_v50), 2)
			(_nw247).ArraySet1(_1447_v50, 3)
			(_nw247).ArraySet1(_1447_v50, 4)
			(_nw247).ArraySet1(_1447_v50, 5)
			(_nw247).ArraySet1(_1447_v50, 6)
			(_nw247).ArraySet1((_1722_v236).Fm0(((_this).F25()).Select((Companion_Default___.SafeIndex(_1447_v50, _dafny.IntOfUint32(((_this).F25()).Cardinality()))).Uint32()).(_dafny.MultiSet), _1447_v50, _1724_v238, (_this).F24(), globalState), 7)
			(_nw247).ArraySet1((_1725_v239).Cardinality(), 8)
			_1726_v240 = _nw247
			var _1727_v241 _dafny.Set
			_ = _1727_v241
			_1727_v241 = _dafny.SetOf(_1726_v240)
			(globalState).F10 = (Companion_D13_.Create_DC46_((_this).F24(), _1722_v236, _1447_v50, _1727_v241)).Dtor_cf73()
			var _1728_v242 _dafny.Sequence
			_ = _1728_v242
			_1728_v242 = _dafny.SeqOf((_1721_v235).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1721_v235), 0))).Int()).(bool))
			var _1729_v243 _dafny.MultiSet
			_ = _1729_v243
			_1729_v243 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_1728_v242, _dafny.SeqOf((_this).F24())))
			var _1730_v244 _dafny.MultiSet
			_ = _1730_v244
			_1730_v244 = _dafny.MultiSetOf((_1721_v235).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1721_v235), 0))).Int()).(bool), (_this).F24())
			_1729_v243 = Companion_Default___.Fm50(((_1730_v244).Cardinality()).Times(_1447_v50), globalState)
			(globalState).F0 = (_1721_v235).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1721_v235), 0))).Int()).(bool)
		}
		var _1731_v245 _dafny.CodePoint
		_ = _1731_v245
		_1731_v245 = _dafny.CodePoint('k')
		var _1732_v246 _dafny.Set
		_ = _1732_v246
		_1732_v246 = _dafny.SetOf(_1731_v245)
		var _1733_v247 _dafny.Map
		_ = _1733_v247
		_1733_v247 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1731_v245, _dafny.SetOf(_1732_v246, Companion_Default___.Fm33(globalState), _1732_v246, _1732_v246))
		var _1734_v248 _dafny.Map
		_ = _1734_v248
		_1734_v248 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1731_v245), _dafny.IntOfUint32((_dafny.SeqOf(_1447_v50)).Cardinality()))
		r0 = (Companion_Default___.Fm7(_1447_v50, globalState)).Cmp(((func() _dafny.Set {
			if (_1733_v247).Contains(_1731_v245) {
				return (_1733_v247).Get(_1731_v245).(_dafny.Set)
			}
			return func() _dafny.Set {
				var _coll68 = _dafny.NewBuilder()
				_ = _coll68
				for _iter72 := _dafny.Iterate((_1734_v248).Keys().Elements()); ; {
					_compr_68, _ok72 := _iter72()
					if !_ok72 {
						break
					}
					var _1735_v249 _dafny.Set
					_1735_v249 = interface{}(_compr_68).(_dafny.Set)
					if (_1734_v248).Contains(_1735_v249) {
						_coll68.Add(_1735_v249)
					}
				}
				return _coll68.ToSet()
			}()
		})()).Cardinality()) < 0
		var _nw248 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
		_ = _nw248
		r1 = _nw248
		r2 = func() _dafny.Map {
			var _coll69 = _dafny.NewMapBuilder()
			_ = _coll69
			for _iter73 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}((func(_1736_v50 _dafny.Int, _1737_v245 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_1738_i22 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg94 _dafny.Int) interface{} {
							return coer94(arg94)
						}
					}((func(_1739_v50 _dafny.Int, _1740_v245 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
						return func(_1741_i23 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jfkcni"), (Companion_Default___.SafeIndex(_1739_v50, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jfkcni")).Cardinality()))).Uint32(), _1740_v245)).Cardinality())
						}
					})(_1736_v50, _1737_v245)))
				}
			})(_1447_v50, _1731_v245)))).Elements()); ; {
				_compr_69, _ok73 := _iter73()
				if !_ok73 {
					break
				}
				var _1742_v250 _dafny.Sequence
				_1742_v250 = interface{}(_compr_69).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}((func(_1743_v50 _dafny.Int, _1744_v245 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_1738_i22 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg96 _dafny.Int) interface{} {
								return coer96(arg96)
							}
						}((func(_1745_v50 _dafny.Int, _1746_v245 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
							return func(_1741_i23 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jfkcni"), (Companion_Default___.SafeIndex(_1745_v50, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jfkcni")).Cardinality()))).Uint32(), _1746_v245)).Cardinality())
							}
						})(_1743_v50, _1744_v245)))
					}
				})(_1447_v50, _1731_v245))), _1742_v250) {
					_coll69.Add(_1742_v250, _1731_v245)
				}
			}
			return _coll69.ToMap()
		}()
		r3 = Companion_Default___.SafeModuloInt(_1447_v50, _1447_v50)
		return r0, r1, r2, r3
	}
}
func (_this *C6) M1(globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1747_v0 _dafny.Int
		_ = _1747_v0
		_1747_v0 = _dafny.IntOfInt64(217)
		r1 = _1747_v0
		if (_this).F24() {
			var _1748_v1 D15
			_ = _1748_v1
			_1748_v1 = Companion_D15_.Create_DC49_(_dafny.SetOf((_this).F24()))
			var _1749_v2 _dafny.Set
			_ = _1749_v2
			_1749_v2 = _dafny.SetOf((_this).F24())
			(globalState).F0 = (((_1748_v1).Dtor_cf81()).Difference(_1749_v2)).Equals(_1749_v2)
			var _1750_v3 _dafny.Array
			_ = _1750_v3
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_42
			var _nw249 _dafny.Array
			_ = _nw249
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw249 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.Int = (func(_1751_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1752_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1752_i0, (_dafny.SetOf(_1751_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1751_v0)).Cardinality())).Cardinality())
					}
				})(_1747_v0)
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw249 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw249).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw249).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1750_v3 = _nw249
			var _1753_v4 _dafny.Set
			_ = _1753_v4
			_1753_v4 = _dafny.SetOf(_1750_v3, _1750_v3, _1750_v3, _1750_v3)
			var _1754_v5 _dafny.Sequence
			_ = _1754_v5
			_1754_v5 = _dafny.UnicodeSeqOfUtf8Bytes("uxel")
			var _1755_v6 D15
			_ = _1755_v6
			_1755_v6 = Companion_D15_.Create_DC51_((_this).F24(), (_this).F24())
			var _1756_v7 _dafny.Sequence
			_ = _1756_v7
			_1756_v7 = _dafny.SeqOf((_this).F24(), (_this).F24(), (_this).F24(), (_this).F24())
			var _1757_v8 _dafny.Map
			_ = _1757_v8
			_1757_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, !((_this).F24()))
			var _1758_v9 _dafny.Array
			_ = _1758_v9
			var _nwElement0_52 bool = (_this).F24()
			_ = _nwElement0_52
			var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(24))
			_ = _nw250
			(_nw250).ArraySet1(_nwElement0_52, 0)
			(_nw250).ArraySet1((_this).F24(), 1)
			(_nw250).ArraySet1((_this).F24(), 2)
			(_nw250).ArraySet1((_this).F24(), 3)
			(_nw250).ArraySet1((_this).F24(), 4)
			(_nw250).ArraySet1((_this).F24(), 5)
			(_nw250).ArraySet1((_this).F24(), 6)
			(_nw250).ArraySet1(!((_this).F24()), 7)
			(_nw250).ArraySet1((_this).F24(), 8)
			(_nw250).ArraySet1((_this).F24(), 9)
			(_nw250).ArraySet1(Companion_Default___.Fm15((_this).F24(), globalState), 10)
			(_nw250).ArraySet1(!((_this).F24()), 11)
			(_nw250).ArraySet1((_this).F24(), 12)
			(_nw250).ArraySet1((_this).F24(), 13)
			(_nw250).ArraySet1((_1755_v6).Dtor_cf82(), 14)
			(_nw250).ArraySet1((_this).F24(), 15)
			(_nw250).ArraySet1((_this).F24(), 16)
			(_nw250).ArraySet1((_this).F24(), 17)
			(_nw250).ArraySet1(Companion_Default___.Fm37(_1756_v7, _1757_v8, _1754_v5, globalState), 18)
			(_nw250).ArraySet1((_this).F24(), 19)
			(_nw250).ArraySet1((_this).F24(), 20)
			(_nw250).ArraySet1((_this).F24(), 21)
			(_nw250).ArraySet1(Companion_Default___.Fm15((_this).F24(), globalState), 22)
			(_nw250).ArraySet1(true, 23)
			_1758_v9 = _nw250
			var _1759_v10 D9
			_ = _1759_v10
			_1759_v10 = Companion_D9_.Create_DC31_(((_1753_v4).Cardinality()).Plus(_dafny.IntOfUint32((_1754_v5).Cardinality())), _1758_v9)
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))
			_ = _index248
			(_1750_v3).ArraySet1(_1747_v0, (_index248).Int())
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))
			_ = _index249
			var _rhs288 D9 = _1759_v10
			_ = _rhs288
			var _rhs289 _dafny.Array = _1750_v3
			_ = _rhs289
			var _rhs290 _dafny.Int = _1747_v0
			_ = _rhs290
			var _rhs291 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1756_v7, (Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32((_1756_v7).Cardinality()))).Uint32(), (_this).F24()), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1747_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1756_v7, (Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32((_1756_v7).Cardinality()))).Uint32(), (_this).F24())).Cardinality()))).Uint32(), (_this).F24()), _dafny.SeqOf((_this).F24(), (_this).F24(), (_this).F24()))
			_ = _rhs291
			var _rhs292 _dafny.Int = _1747_v0
			_ = _rhs292
			var _lhs235 _dafny.Array = _1750_v3
			_ = _lhs235
			var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))
			_ = _lhs236
			var _lhs237 *GlobalState = globalState
			_ = _lhs237
			_1759_v10 = _rhs288
			_1750_v3 = _rhs289
			(_lhs235).ArraySet1(_rhs290, (_lhs236).Int())
			_1756_v7 = _rhs291
			_lhs237.F8 = _rhs292
			var _1760_v11 _dafny.Map
			_ = _1760_v11
			_1760_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1750_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))).Int()).(_dafny.Int), Companion_Default___.Fm7((_dafny.Zero).Minus((_1749_v2).Cardinality()), globalState))
			var _1761_v12 _dafny.Map
			_ = _1761_v12
			_1761_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
			_1760_v11 = (_1760_v11).Update(Companion_Default___.SafeDivisionInt(_1747_v0, (_1750_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))).Int()).(_dafny.Int)), (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf(_1761_v12, _1761_v12)).Cardinality())).Minus((_1750_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))).Int()).(_dafny.Int))))
			var _1762_v13 _dafny.Map
			_ = _1762_v13
			_1762_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1754_v5, (_this).F24())
			var _1763_v14 _dafny.Map
			_ = _1763_v14
			_1763_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F24()), _1762_v13)
			var _rhs293 _dafny.Map = (_1762_v13).Merge((func() _dafny.Map {
				if (_1763_v14).Contains((_this).F24()) {
					return (_1763_v14).Get((_this).F24()).(_dafny.Map)
				}
				return _1762_v13
			})())
			_ = _rhs293
			var _rhs294 bool = (_dafny.IntOfUint32((_1754_v5).Cardinality())).Cmp((_1750_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_1750_v3), 0))).Int()).(_dafny.Int)) == 0
			_ = _rhs294
			var _lhs238 *GlobalState = globalState
			_ = _lhs238
			_1762_v13 = _rhs293
			_lhs238.F0 = _rhs294
			_1747_v0 = _1747_v0
		} else {
			var _1764_v15 _dafny.Array
			_ = _1764_v15
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_43
			var _nw251 _dafny.Array
			_ = _nw251
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw251 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) bool = func(_1765_i1 _dafny.Int) bool {
					return (_this).F24()
				}
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw251 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw251).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw251).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1764_v15 = _nw251
			var _1766_v16 *C1
			_ = _1766_v16
			var _nw252 *C1 = New_C1_()
			_ = _nw252
			_nw252.Ctor__(Companion_Default___.SafeDivisionInt(_1747_v0, _dafny.IntOfInt64(817)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1764_v15, _1764_v15, _1764_v15, _1764_v15, _1764_v15), (_this).F22()))
			_1766_v16 = _nw252
			if Companion_Default___.Fm15(true, globalState) {
				var _1767_v17 _dafny.Int
				_ = _1767_v17
				var _1768_v18 _dafny.Set
				_ = _1768_v18
				var _1769_v19 _dafny.Int
				_ = _1769_v19
				var _1770_v20 _dafny.Int
				_ = _1770_v20
				var _out60 _dafny.Int
				_ = _out60
				var _out61 _dafny.Set
				_ = _out61
				var _out62 _dafny.Int
				_ = _out62
				var _out63 _dafny.Int
				_ = _out63
				_out60, _out61, _out62, _out63 = (_1766_v16).M6(_1747_v0, ((_1766_v16).F30()).Cmp((_1766_v16).F30()) == 0, Companion_Default___.SafeModuloInt(_1747_v0, _1747_v0), (_1766_v16).F30(), globalState)
				_1767_v17 = _out60
				_1768_v18 = _out61
				_1769_v19 = _out62
				_1770_v20 = _out63
				var _1771_v21 _dafny.CodePoint
				_ = _1771_v21
				_1771_v21 = _dafny.CodePoint('q')
				_1771_v21 = _1771_v21
				var _1772_v22 _dafny.Set
				_ = _1772_v22
				_1772_v22 = _dafny.SetOf((_this).F24(), (_this).F24(), false)
				var _1773_v23 _dafny.Sequence
				_ = _1773_v23
				_1773_v23 = _dafny.SeqOf((_this).F24())
				var _1774_v24 _dafny.Array
				_ = _1774_v24
				var _nwElement0_53 _dafny.Set = (_1772_v22).Intersection(_1772_v22)
				_ = _nwElement0_53
				var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(23))
				_ = _nw253
				(_nw253).ArraySet1(_nwElement0_53, 0)
				(_nw253).ArraySet1(_1772_v22, 1)
				(_nw253).ArraySet1((func() _dafny.Set {
					if (_this).F24() {
						return _1772_v22
					}
					return _1772_v22
				})(), 2)
				(_nw253).ArraySet1(_1772_v22, 3)
				(_nw253).ArraySet1((_1772_v22).Intersection(_1772_v22), 4)
				(_nw253).ArraySet1(_dafny.SetOf((_this).F24()), 5)
				(_nw253).ArraySet1(_1772_v22, 6)
				(_nw253).ArraySet1(_1772_v22, 7)
				(_nw253).ArraySet1(_1772_v22, 8)
				(_nw253).ArraySet1(_1772_v22, 9)
				(_nw253).ArraySet1((_1772_v22).Intersection(_1772_v22), 10)
				(_nw253).ArraySet1(_dafny.SetOf((_this).F24(), (_this).F24(), (_this).F24()), 11)
				(_nw253).ArraySet1(_1772_v22, 12)
				(_nw253).ArraySet1(_1772_v22, 13)
				(_nw253).ArraySet1((_dafny.SetOf((_1773_v23).Select((Companion_Default___.SafeIndex((_1766_v16).F30(), _dafny.IntOfUint32((_1773_v23).Cardinality()))).Uint32()).(bool))).Difference(_1772_v22), 14)
				(_nw253).ArraySet1(_1772_v22, 15)
				(_nw253).ArraySet1((_1772_v22).Union(_1772_v22), 16)
				(_nw253).ArraySet1(_1772_v22, 17)
				(_nw253).ArraySet1(_1772_v22, 18)
				(_nw253).ArraySet1(_1772_v22, 19)
				(_nw253).ArraySet1(_1772_v22, 20)
				(_nw253).ArraySet1((Companion_Default___.Fm11(_1770_v20, _1771_v21, (_this).F24(), globalState)).Union(_1772_v22), 21)
				(_nw253).ArraySet1(_dafny.SetOf(!((_this).F24()), false), 22)
				_1774_v24 = _nw253
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1774_v24), 0))
				_ = _index250
				(_1774_v24).ArraySet1(_1772_v22, (_index250).Int())
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1774_v24), 0))
				_ = _index251
				(_1774_v24).ArraySet1(_1772_v22, (_index251).Int())
				var _1775_v25 D13
				_ = _1775_v25
				_1775_v25 = Companion_D13_.Create_DC45_(_1769_v19)
				var _pat_let_tv66 = _1767_v17
				_ = _pat_let_tv66
				_1775_v25 = func(_pat_let36_0 D13) D13 {
					return func(_1776_dt__update__tmp_h0 D13) D13 {
						return func(_pat_let37_0 _dafny.Int) D13 {
							return func(_1777_dt__update_hcf72_h0 _dafny.Int) D13 {
								return Companion_D13_.Create_DC45_(_1777_dt__update_hcf72_h0)
							}(_pat_let37_0)
						}(_pat_let_tv66)
					}(_pat_let36_0)
				}(_1775_v25)
				var _1778_v26 *C2
				_ = _1778_v26
				var _nw254 *C2 = New_C2_()
				_ = _nw254
				_nw254.Ctor__(_1770_v20, _dafny.SeqOf(_1764_v15, _1764_v15, _1764_v15, _1764_v15))
				_1778_v26 = _nw254
				var _1779_v27 _dafny.Map
				_ = _1779_v27
				_1779_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(296), _1767_v17)
				var _1780_v28 _dafny.Array
				_ = _1780_v28
				var _nwElement0_54 _dafny.Int = (_1766_v16).F30()
				_ = _nwElement0_54
				var _nw255 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(11))
				_ = _nw255
				(_nw255).ArraySet1(_nwElement0_54, 0)
				(_nw255).ArraySet1(_dafny.IntOfInt64(-755), 1)
				(_nw255).ArraySet1((_1779_v27).Cardinality(), 2)
				(_nw255).ArraySet1(_dafny.IntOfInt64(978), 3)
				(_nw255).ArraySet1((_1766_v16).F30(), 4)
				(_nw255).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(692))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}((func(_1781_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1782_i2 _dafny.Int) _dafny.CodePoint {
						return _1781_v21
					}
				})(_1771_v21)))).Cardinality()), 5)
				(_nw255).ArraySet1(_1769_v19, 6)
				(_nw255).ArraySet1((_1778_v26).F29(), 7)
				(_nw255).ArraySet1(_1769_v19, 8)
				(_nw255).ArraySet1((_1766_v16).F30(), 9)
				(_nw255).ArraySet1(_1747_v0, 10)
				_1780_v28 = _nw255
				var _1783_v29 _dafny.Map
				_ = _1783_v29
				_1783_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1778_v26, _1780_v28))
				var _1784_v30 _dafny.Map
				_ = _1784_v30
				_1784_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1778_v26, _1780_v28)
				_1783_v29 = (_1783_v29).Update(!((_this).F24()), (_1784_v30).Merge(_1784_v30))
			} else {
				var _1785_v31 _dafny.Sequence
				_ = _1785_v31
				_1785_v31 = _dafny.UnicodeSeqOfUtf8Bytes("iyldgbb")
				var _1786_v32 _dafny.CodePoint
				_ = _1786_v32
				_1786_v32 = _dafny.CodePoint('i')
				var _1787_v33 _dafny.Sequence
				_ = _1787_v33
				_1787_v33 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1785_v31, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.IntOfUint32((_1785_v31).Cardinality()))).Uint32(), _1786_v32), _1785_v31)).Cardinality()))
				var _1788_v34 D2
				_ = _1788_v34
				_1788_v34 = Companion_D2_.Create_DC7_(_1787_v33, (_this).F24(), (_this).F24(), (_1766_v16).F30())
				var _1789_v35 _dafny.Map
				_ = _1789_v35
				_1789_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-395), (_1788_v34).Dtor_cf10())
				_1787_v33 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1787_v33, (Companion_Default___.SafeIndex((_1766_v16).Fm34(false, globalState), _dafny.IntOfUint32((_1787_v33).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1785_v31).Cardinality())), (func() _dafny.Sequence {
					if (_1789_v35).Contains((_1766_v16).F30()) {
						return (_1789_v35).Get((_1766_v16).F30()).(_dafny.Sequence)
					}
					return _1787_v33
				})())
				var _1790_v36 *C1
				_ = _1790_v36
				var _nw256 *C1 = New_C1_()
				_ = _nw256
				_nw256.Ctor__((_1766_v16).F30(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F22(), (Companion_Default___.SafeIndex((_1766_v16).F30(), _dafny.IntOfUint32(((_this).F22()).Cardinality()))).Uint32(), _1764_v15), _dafny.SeqOf(_1764_v15)))
				_1790_v36 = _nw256
				var _1791_v37 _dafny.MultiSet
				_ = _1791_v37
				_1791_v37 = _dafny.MultiSetOf((_this).F24(), (_this).F24(), (_this).F24(), (_this).F24())
				var _1792_v38 _dafny.MultiSet
				_ = _1792_v38
				_1792_v38 = _dafny.MultiSetOf(_1747_v0)
				var _1793_v39 _dafny.Set
				_ = _1793_v39
				_1793_v39 = _dafny.SetOf((_1766_v16).F30())
				var _1794_v40 _dafny.Sequence
				_ = _1794_v40
				_1794_v40 = _dafny.SeqOf(_1793_v39)
				var _1795_v41 _dafny.Map
				_ = _1795_v41
				_1795_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1766_v16).Fm0(_1792_v38, _1747_v0, _1794_v40, (_this).F24(), globalState), (_this).F24())
				(globalState).F16 = ((func() _dafny.Int {
					if (_1791_v37).Contains((_this).F24()) {
						return (_1791_v37).Multiplicity((_this).F24())
					}
					return (_1795_v41).Cardinality()
				})()).Times(((_1790_v36).F30()).Plus((_1790_v36).F30()))
				r0 = (_this).Fm3(_1785_v31, (func() bool {
					if (_this).F24() {
						return (_this).F24()
					}
					return (_this).F24()
				})(), (_this).F24(), globalState)
				var _1796_v42 _dafny.Map
				_ = _1796_v42
				_1796_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vcllqkg"), _1787_v33)
				var _1797_v44 _dafny.Map
				_ = _1797_v44
				_1797_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1785_v31, (_1790_v36).F30())
				_1796_v42 = func() _dafny.Map {
					var _coll70 = _dafny.NewMapBuilder()
					_ = _coll70
					for _iter74 := _dafny.Iterate((_1797_v44).Keys().Elements()); ; {
						_compr_70, _ok74 := _iter74()
						if !_ok74 {
							break
						}
						var _1798_v43 _dafny.Sequence
						_1798_v43 = interface{}(_compr_70).(_dafny.Sequence)
						if (_1797_v44).Contains(_1798_v43) {
							_coll70.Add(_1798_v43, _dafny.Companion_Sequence_.Concatenate(_1787_v33, _1787_v33))
						}
					}
					return _coll70.ToMap()
				}()
			}
			r1 = _1747_v0
			var _1799_v45 _dafny.Sequence
			_ = _1799_v45
			_1799_v45 = _dafny.SeqOf(Companion_Default___.Fm51(globalState))
			var _1800_v46 _dafny.Array
			_ = _1800_v46
			var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw257
			_1800_v46 = _nw257
			var _1801_v47 _dafny.Sequence
			_ = _1801_v47
			_1801_v47 = _dafny.SeqOf(_1800_v46)
			var _1802_v48 _dafny.Map
			_ = _1802_v48
			_1802_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(573), _1800_v46)
			var _1803_v49 _dafny.Set
			_ = _1803_v49
			_1803_v49 = _dafny.SetOf((_1766_v16).F30())
			var _1804_v50 _dafny.Array
			_ = _1804_v50
			var _nwElement0_55 _dafny.Array = _1800_v46
			_ = _nwElement0_55
			var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(19))
			_ = _nw258
			(_nw258).ArraySet1(_nwElement0_55, 0)
			(_nw258).ArraySet1(_1800_v46, 1)
			(_nw258).ArraySet1(_1800_v46, 2)
			(_nw258).ArraySet1((_1801_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.IntOfUint32((_1801_v47).Cardinality()))).Uint32()).(_dafny.Array), 3)
			(_nw258).ArraySet1(_1800_v46, 4)
			(_nw258).ArraySet1(_1800_v46, 5)
			(_nw258).ArraySet1(_1800_v46, 6)
			(_nw258).ArraySet1(_1800_v46, 7)
			(_nw258).ArraySet1((func() _dafny.Array {
				if (_this).F24() {
					return _1800_v46
				}
				return _1800_v46
			})(), 8)
			(_nw258).ArraySet1(_1800_v46, 9)
			(_nw258).ArraySet1(_1800_v46, 10)
			(_nw258).ArraySet1(_1800_v46, 11)
			(_nw258).ArraySet1(_1800_v46, 12)
			(_nw258).ArraySet1(_1800_v46, 13)
			(_nw258).ArraySet1((func() _dafny.Array {
				if (_this).F24() {
					return _1800_v46
				}
				return _1800_v46
			})(), 14)
			(_nw258).ArraySet1(_1800_v46, 15)
			(_nw258).ArraySet1((func() _dafny.Array {
				if (_1802_v48).Contains((_1766_v16).F30()) {
					return (_1802_v48).Get((_1766_v16).F30()).(_dafny.Array)
				}
				return _1800_v46
			})(), 16)
			(_nw258).ArraySet1((func() _dafny.Array {
				if (_this).F24() {
					return (func() _dafny.Array {
						if (_1802_v48).Contains((_1803_v49).Cardinality()) {
							return (_1802_v48).Get((_1803_v49).Cardinality()).(_dafny.Array)
						}
						return _1800_v46
					})()
				}
				return _1800_v46
			})(), 17)
			(_nw258).ArraySet1(_1800_v46, 18)
			_1804_v50 = _nw258
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1804_v50), 0))
			_ = _index252
			(_1804_v50).ArraySet1(_1800_v46, (_index252).Int())
			var _1805_v51 D4
			_ = _1805_v51
			_1805_v51 = Companion_D4_.Create_DC15_(_dafny.IntOfInt64(238), _1747_v0)
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1804_v50), 0))
			_ = _index253
			var _rhs295 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1799_v45, _1799_v45)
			_ = _rhs295
			var _rhs296 _dafny.Int = _1747_v0
			_ = _rhs296
			var _rhs297 _dafny.Array = _1800_v46
			_ = _rhs297
			var _rhs298 bool = (((_1805_v51).Dtor_cf26()).Cmp(_1747_v0) >= 0) == ((_this).F24())
			_ = _rhs298
			var _lhs239 *GlobalState = globalState
			_ = _lhs239
			var _lhs240 _dafny.Array = _1804_v50
			_ = _lhs240
			var _lhs241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1804_v50), 0))
			_ = _lhs241
			var _lhs242 *GlobalState = globalState
			_ = _lhs242
			_1799_v45 = _rhs295
			_lhs239.F6 = _rhs296
			(_lhs240).ArraySet1(_rhs297, (_lhs241).Int())
			_lhs242.F10 = _rhs298
			var _1806_v52 *C3
			_ = _1806_v52
			var _nw259 *C3 = New_C3_()
			_ = _nw259
			_nw259.Ctor__(_dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()))
			_1806_v52 = _nw259
		}
		var _1807_i3 _dafny.Int
		_ = _1807_i3
		_1807_i3 = _dafny.Zero
		{
			for (_this).F24() {
				{
					if (_1807_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1807_i3 = (_1807_i3).Plus(_dafny.One)
					var _1808_v53 _dafny.Array
					_ = _1808_v53
					var _len0_44 _dafny.Int = _dafny.IntOfInt64(4)
					_ = _len0_44
					var _nw260 _dafny.Array
					_ = _nw260
					if _len0_44.Cmp(_dafny.Zero) == 0 {
						_nw260 = _dafny.NewArray(_len0_44)
					} else {
						var _init44 func(_dafny.Int) bool = func(_1809_i4 _dafny.Int) bool {
							return !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()))
						}
						_ = _init44
						var _element0_44 = _init44(_dafny.Zero)
						_ = _element0_44
						_nw260 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
						(_nw260).ArraySet1(_element0_44, 0)
						var _nativeLen0_44 = (_len0_44).Int()
						_ = _nativeLen0_44
						for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
							(_nw260).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
						}
					}
					_1808_v53 = _nw260
					var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))
					_ = _index254
					(_1808_v53).ArraySet1(!((_this).F24()), (_index254).Int())
					var _1810_v54 _dafny.Sequence
					_ = _1810_v54
					_1810_v54 = _dafny.SeqOf(_1747_v0)
					var _1811_v55 _dafny.Sequence
					_ = _1811_v55
					_1811_v55 = _dafny.SeqOf(_1810_v54, _dafny.Companion_Sequence_.Update(_1810_v54, (Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32((_1810_v54).Cardinality()))).Uint32(), _1747_v0), _1810_v54)
					var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))
					_ = _index255
					(_1808_v53).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1811_v55, _1810_v54), (_index255).Int())
					var _1812_v56 _dafny.Sequence
					_ = _1812_v56
					_1812_v56 = _dafny.SeqOf((_1808_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))).Int()).(bool))
					var _1813_v57 _dafny.MultiSet
					_ = _1813_v57
					_1813_v57 = _dafny.MultiSetOf(_1747_v0, _1747_v0, _1747_v0, _dafny.IntOfUint32((_1812_v56).Cardinality()))
					(globalState).F0 = !((_1813_v57).Intersection(_1813_v57)).Contains(_1747_v0)
					var _1814_v58 _dafny.Set
					_ = _1814_v58
					_1814_v58 = _dafny.SetOf((_1808_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))).Int()).(bool), (_1808_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))).Int()).(bool), (_this).F24())
					var _1815_v59 _dafny.Set
					_ = _1815_v59
					_1815_v59 = _dafny.SetOf(_dafny.IntOfInt64(526), (_1814_v58).Cardinality(), _dafny.IntOfInt64(272), _1747_v0)
					(globalState).F13 = _1815_v59
					var _1816_v60 _dafny.Map
					_ = _1816_v60
					_1816_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1808_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))).Int()).(bool), _1747_v0)
					var _1817_v61 _dafny.Map
					_ = _1817_v61
					_1817_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, (_1816_v60).Cardinality())
					var _1818_v62 _dafny.Map
					_ = _1818_v62
					_1818_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F24()), _1817_v61)
					_1818_v62 = (_1818_v62).Update((_1808_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1808_v53), 0))).Int()).(bool), (_1817_v61).Update(_1747_v0, _1747_v0))
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		if !((_this).F24()) {
			var _1819_v63 _dafny.Array
			_ = _1819_v63
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw261
			_1819_v63 = _nw261
			var _1820_v64 _dafny.Sequence
			_ = _1820_v64
			_1820_v64 = _dafny.UnicodeSeqOfUtf8Bytes("xlluhdhw")
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1819_v63), 0))
			_ = _index256
			(_1819_v63).ArraySet1(_1820_v64, (_index256).Int())
			var _1821_v65 _dafny.Set
			_ = _1821_v65
			_1821_v65 = _dafny.SetOf(_1747_v0)
			var _1822_v67 _dafny.MultiSet
			_ = _1822_v67
			_1822_v67 = _dafny.MultiSetOf(true)
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1819_v63), 0))
			_ = _index257
			(_1819_v63).ArraySet1(Companion_Default___.Fm36((_this).F24(), (_1821_v65).Difference(func() _dafny.Set {
				var _coll71 = _dafny.NewBuilder()
				_ = _coll71
				for _iter75 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(9), _dafny.IntOfInt64(495))); ; {
					_compr_71, _ok75 := _iter75()
					if !_ok75 {
						break
					}
					var _1823_v66 _dafny.Int
					_1823_v66 = interface{}(_compr_71).(_dafny.Int)
					if ((_dafny.IntOfInt64(9)).Cmp(_1823_v66) <= 0) && ((_1823_v66).Cmp(_dafny.IntOfInt64(495)) < 0) {
						_coll71.Add(Companion_Default___.SafeModuloInt(_1823_v66, _1747_v0))
					}
				}
				return _coll71.ToSet()
			}()), _1747_v0, (_1822_v67).Cardinality(), globalState), (_index257).Int())
			var _1824_v68 _dafny.Array
			_ = _1824_v68
			var _len0_45 _dafny.Int = _dafny.One
			_ = _len0_45
			var _nw262 _dafny.Array
			_ = _nw262
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw262 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) bool = func(_1825_i5 _dafny.Int) bool {
					return (_this).F24()
				}
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw262 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw262).ArraySet1(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw262).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1824_v68 = _nw262
			var _1826_v69 _dafny.Map
			_ = _1826_v69
			_1826_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1824_v68, (_this).F24())
			_1826_v69 = ((_1826_v69).Update(_1824_v68, (_this).F24())).Merge(_1826_v69)
			var _1827_v70 _dafny.CodePoint
			_ = _1827_v70
			_1827_v70 = _dafny.CodePoint('b')
			var _1828_v71 _dafny.Map
			_ = _1828_v71
			_1828_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_this).F24()), _1827_v70)
			var _1829_v72 _dafny.Sequence
			_ = _1829_v72
			_1829_v72 = _dafny.SeqOf(false)
			var _1830_v73 _dafny.Map
			_ = _1830_v73
			_1830_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, (_this).F24())
			var _1831_v74 _dafny.MultiSet
			_ = _1831_v74
			_1831_v74 = _dafny.MultiSetOf(_dafny.IntOfInt64(618), (_1830_v73).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_1819_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1819_v63), 0))).Int()).(_dafny.Sequence)).Cardinality())))
			var _1832_v75 _dafny.Map
			_ = _1832_v75
			_1832_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
			var _1833_v76 _dafny.Array
			_ = _1833_v76
			var _nwElement0_56 _dafny.Int = (_1828_v71).Cardinality()
			_ = _nwElement0_56
			var _nw263 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(13))
			_ = _nw263
			(_nw263).ArraySet1(_nwElement0_56, 0)
			(_nw263).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1829_v72, _dafny.SeqOf((_this).F24(), (_this).F24(), (_this).F24(), false))).Cardinality()), 1)
			(_nw263).ArraySet1(((_1821_v65).Cardinality()).Minus(_dafny.IntOfInt64(983)), 2)
			(_nw263).ArraySet1((_dafny.Zero).Minus(_1747_v0), 3)
			(_nw263).ArraySet1(_1747_v0, 4)
			(_nw263).ArraySet1(_dafny.IntOfUint32((_1820_v64).Cardinality()), 5)
			(_nw263).ArraySet1(Companion_Default___.SafeDivisionInt(_1747_v0, _1747_v0), 6)
			(_nw263).ArraySet1((_1747_v0).Times((func() _dafny.Int {
				if (_1831_v74).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qrhauoaod")).Cardinality())) {
					return (_1831_v74).Multiplicity(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qrhauoaod")).Cardinality()))
				}
				return _1747_v0
			})()), 7)
			(_nw263).ArraySet1(_1747_v0, 8)
			(_nw263).ArraySet1(_dafny.IntOfInt64(-508), 9)
			(_nw263).ArraySet1(((_1832_v75).Merge(_1832_v75)).Cardinality(), 10)
			(_nw263).ArraySet1(_1747_v0, 11)
			(_nw263).ArraySet1(_1747_v0, 12)
			_1833_v76 = _nw263
			var _1834_v77 D2
			_ = _1834_v77
			_1834_v77 = Companion_D2_.Create_DC6_(_1747_v0)
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1833_v76), 0))
			_ = _index258
			(_1833_v76).ArraySet1((_1834_v77).Dtor_cf9(), (_index258).Int())
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1833_v76), 0))
			_ = _index259
			var _rhs299 _dafny.Int = (_1747_v0).Plus(_dafny.IntOfUint32(((_1819_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1819_v63), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			_ = _rhs299
			var _rhs300 bool = ((_this).F24()) || (!((_this).F24()))
			_ = _rhs300
			var _lhs243 _dafny.Array = _1833_v76
			_ = _lhs243
			var _lhs244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1833_v76), 0))
			_ = _lhs244
			(_lhs243).ArraySet1(_rhs299, (_lhs244).Int())
			r0 = _rhs300
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_1824_v68), 0))
			_ = _index260
			(_1824_v68).ArraySet1((!((_this).F24())) || ((_this).F24()), (_index260).Int())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_1824_v68), 0))
			_ = _index261
			var _rhs301 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm52(globalState), (Companion_Default___.SafeIndex((_1830_v73).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm52(globalState)).Cardinality()))).Uint32(), (_this).F24()), _1829_v72), (_1829_v72).Select((Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32((_1829_v72).Cardinality()))).Uint32()).(bool))
			_ = _rhs301
			var _rhs302 _dafny.Int = ((_1833_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1833_v76), 0))).Int()).(_dafny.Int)).Minus(_1747_v0)
			_ = _rhs302
			var _rhs303 _dafny.Int = _1747_v0
			_ = _rhs303
			var _lhs245 _dafny.Array = _1824_v68
			_ = _lhs245
			var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_1824_v68), 0))
			_ = _lhs246
			var _lhs247 *GlobalState = globalState
			_ = _lhs247
			(_lhs245).ArraySet1(_rhs301, (_lhs246).Int())
			r1 = _rhs302
			_lhs247.F8 = _rhs303
			_1824_v68 = _1824_v68
		} else {
			r0 = (_this).F24()
			if (_this).F24() {
				var _1835_v78 _dafny.Array
				_ = _1835_v78
				var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw264
				_1835_v78 = _nw264
				var _1836_v79 _dafny.Map
				_ = _1836_v79
				_1836_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1835_v78, (_1747_v0).Cmp(_1747_v0) == 0)
				_1836_v79 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1835_v78, false)).Merge(_1836_v79)
				(globalState).F10 = (_this).F24()
				var _1837_v80 _dafny.Sequence
				_ = _1837_v80
				_1837_v80 = _dafny.UnicodeSeqOfUtf8Bytes("dxmowc")
				_1837_v80 = _dafny.Companion_Sequence_.Concatenate(_1837_v80, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("w"), Companion_Default___.Fm30(globalState)))
				var _1838_v81 *C5
				_ = _1838_v81
				var _nw265 *C5 = New_C5_()
				_ = _nw265
				_nw265.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, (_this).F24()), _dafny.IntOfInt64(895))
				_1838_v81 = _nw265
				var _1839_v82 *C2
				_ = _1839_v82
				var _nw266 *C2 = New_C2_()
				_ = _nw266
				_nw266.Ctor__(Companion_Default___.SafeDivisionInt(_1747_v0, _1747_v0), (_this).F22())
				_1839_v82 = _nw266
			} else {
				var _1840_v83 _dafny.Array
				_ = _1840_v83
				var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw267
				_1840_v83 = _nw267
				var _1841_v85 _dafny.Array
				_ = _1841_v85
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_46
				var _nw268 _dafny.Array
				_ = _nw268
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw268 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Map = (func(_1842_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_1843_i6 _dafny.Int) _dafny.Map {
							return func() _dafny.Map {
								var _coll72 = _dafny.NewMapBuilder()
								_ = _coll72
								for _iter76 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1842_v0, (_this).F24())).Cardinality(), _1842_v0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-811)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_1842_v0))).Cardinality()), _1842_v0)).Elements()); ; {
									_compr_72, _ok76 := _iter76()
									if !_ok76 {
										break
									}
									var _1844_v84 _dafny.Int
									_1844_v84 = interface{}(_compr_72).(_dafny.Int)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1842_v0, (_this).F24())).Cardinality(), _1842_v0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-811)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_1842_v0))).Cardinality()), _1842_v0), _1844_v84) {
										_coll72.Add(Companion_Default___.SafeModuloInt(_1844_v84, _1842_v0), _1842_v0)
									}
								}
								return _coll72.ToMap()
							}()
						}
					})(_1747_v0)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw268 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw268).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw268).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1841_v85 = _nw268
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_1840_v83), 0))
				_ = _index262
				(_1840_v83).ArraySet1(_1841_v85, (_index262).Int())
				var _1845_v86 _dafny.CodePoint
				_ = _1845_v86
				_1845_v86 = _dafny.CodePoint('i')
				var _1846_v87 _dafny.Array
				_ = _1846_v87
				var _nw269 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw269
				_1846_v87 = _nw269
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1846_v87), 0))
				_ = _index263
				(_1846_v87).ArraySet1((_this).F24(), (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_1840_v83), 0))
				_ = _index264
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1846_v87), 0))
				_ = _index265
				var _rhs304 _dafny.Array = _1841_v85
				_ = _rhs304
				var _rhs305 _dafny.CodePoint = Companion_Default___.Fm22((_this).F24(), !((_this).F24()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm52(globalState), Companion_Default___.Fm52(globalState))).Cardinality()), (_dafny.IntOfInt64(684)).Times(_1747_v0), globalState)
				_ = _rhs305
				var _rhs306 bool = (_this).F24()
				_ = _rhs306
				var _lhs248 _dafny.Array = _1840_v83
				_ = _lhs248
				var _lhs249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_1840_v83), 0))
				_ = _lhs249
				var _lhs250 _dafny.Array = _1846_v87
				_ = _lhs250
				var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1846_v87), 0))
				_ = _lhs251
				(_lhs248).ArraySet1(_rhs304, (_lhs249).Int())
				_1845_v86 = _rhs305
				(_lhs250).ArraySet1(_rhs306, (_lhs251).Int())
				var _1847_v88 _dafny.Set
				_ = _1847_v88
				_1847_v88 = _dafny.SetOf(_1747_v0, _1747_v0)
				var _1848_v89 _dafny.MultiSet
				_ = _1848_v89
				_1848_v89 = _dafny.MultiSetOf(_1747_v0, _1747_v0)
				var _1849_v90 D6
				_ = _1849_v90
				_1849_v90 = Companion_D6_.Create_DC23_((_this).F24(), _1747_v0, (_1846_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1846_v87), 0))).Int()).(bool))
				var _1850_v91 _dafny.Sequence
				_ = _1850_v91
				_1850_v91 = _dafny.SeqOf(_1847_v88)
				var _1851_v92 _dafny.Map
				_ = _1851_v92
				_1851_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm36((_this).F24(), _1847_v88, _1747_v0, (_dafny.Zero).Minus((_this).Fm0(_1848_v89, (_1849_v90).Dtor_cf38(), _1850_v91, (_1846_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1846_v87), 0))).Int()).(bool), globalState)), globalState), (_1747_v0).Plus(_1747_v0))
				var _1852_v93 _dafny.Sequence
				_ = _1852_v93
				_1852_v93 = _dafny.UnicodeSeqOfUtf8Bytes("uji")
				_1851_v92 = (_1851_v92).Update(_1852_v93, (_dafny.Zero).Minus(_1747_v0))
				var _1853_v94 _dafny.Array
				_ = _1853_v94
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_47
				var _nw270 _dafny.Array
				_ = _nw270
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw270 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Int = (func(_1854_v87 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_1855_i7 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1855_i7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1854_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1854_v87), 0))).Int()).(bool), true)).Cardinality())
						}
					})(_1846_v87)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw270 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw270).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw270).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1853_v94 = _nw270
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1853_v94), 0))
				_ = _index266
				(_1853_v94).ArraySet1((_dafny.IntOfInt64(713)).Plus((Companion_Default___.Fm27(_1747_v0, globalState)).Cardinality()), (_index266).Int())
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1853_v94), 0))
				_ = _index267
				(_1853_v94).ArraySet1((_1747_v0).Times(_1747_v0), (_index267).Int())
				_1853_v94 = _1853_v94
				var _1856_v95 _dafny.Map
				_ = _1856_v95
				_1856_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1846_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1846_v87), 0))).Int()).(bool), _dafny.IntOfUint32((_1852_v93).Cardinality()))
				var _1857_v96 _dafny.Set
				_ = _1857_v96
				_1857_v96 = _dafny.SetOf(Companion_Default___.Fm15(false, globalState))
				var _1858_v97 _dafny.Map
				_ = _1858_v97
				_1858_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1856_v95).Cardinality(), _1857_v96)
				var _rhs307 bool = !((_this).F24())
				_ = _rhs307
				var _rhs308 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1852_v93, _1852_v93)
				_ = _rhs308
				var _rhs309 _dafny.Map = ((_1858_v97).Merge(_1858_v97)).Update((_1853_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1853_v94), 0))).Int()).(_dafny.Int), _1857_v96)
				_ = _rhs309
				var _rhs310 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1852_v93, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-748))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_1859_v86 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1860_i8 _dafny.Int) _dafny.CodePoint {
						return _1859_v86
					}
				})(_1845_v86))))
				_ = _rhs310
				var _lhs252 *GlobalState = globalState
				_ = _lhs252
				r0 = _rhs307
				_lhs252.F10 = _rhs308
				_1858_v97 = _rhs309
				_1852_v93 = _rhs310
			}
			(globalState).F0 = (_this).F24()
			var _1861_v98 _dafny.Array
			_ = _1861_v98
			var _len0_48 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_48
			var _nw271 _dafny.Array
			_ = _nw271
			if _len0_48.Cmp(_dafny.Zero) == 0 {
				_nw271 = _dafny.NewArray(_len0_48)
			} else {
				var _init48 func(_dafny.Int) _dafny.Set = func(_1862_i9 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(!((_this).F24()), !(!((_this).F24())))
				}
				_ = _init48
				var _element0_48 = _init48(_dafny.Zero)
				_ = _element0_48
				_nw271 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
				(_nw271).ArraySet1(_element0_48, 0)
				var _nativeLen0_48 = (_len0_48).Int()
				_ = _nativeLen0_48
				for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
					(_nw271).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
				}
			}
			_1861_v98 = _nw271
			var _1863_v99 _dafny.Set
			_ = _1863_v99
			_1863_v99 = _dafny.SetOf((_this).F24())
			var _1864_v100 _dafny.Map
			_ = _1864_v100
			_1864_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1863_v99)
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_1861_v98), 0))
			_ = _index268
			(_1861_v98).ArraySet1((func() _dafny.Set {
				if (_1864_v100).Contains((_this).F24()) {
					return (_1864_v100).Get((_this).F24()).(_dafny.Set)
				}
				return _1863_v99
			})(), (_index268).Int())
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_1861_v98), 0))
			_ = _index269
			(_1861_v98).ArraySet1(_1863_v99, (_index269).Int())
			var _1865_v101 _dafny.MultiSet
			_ = _1865_v101
			_1865_v101 = _dafny.MultiSetOf(_1747_v0)
			var _1866_v102 _dafny.Set
			_ = _1866_v102
			_1866_v102 = _dafny.SetOf(_1865_v101)
			var _rhs311 _dafny.Int = _1747_v0
			_ = _rhs311
			var _rhs312 bool = ((_this).F24()) || ((_1866_v102).IsDisjointFrom(_1866_v102))
			_ = _rhs312
			var _lhs253 *GlobalState = globalState
			_ = _lhs253
			var _lhs254 *GlobalState = globalState
			_ = _lhs254
			_lhs253.F6 = _rhs311
			_lhs254.F10 = _rhs312
		}
		var _1867_v103 _dafny.Set
		_ = _1867_v103
		_1867_v103 = _dafny.SetOf(((_this).F24()) == ((_this).F24()), (_this).F24(), (_this).F24())
		_1867_v103 = _1867_v103
		var _1868_v104 D15
		_ = _1868_v104
		_1868_v104 = Companion_D15_.Create_DC51_((_this).F24(), (_this).F24())
		var _source25 D15 = (func() D15 {
			if (_this).F24() {
				return _1868_v104
			}
			return _1868_v104
		})()
		_ = _source25
		if _source25.Is_DC50() {
			(globalState).F6 = _1747_v0
			var _1869_v105 _dafny.Sequence
			_ = _1869_v105
			_1869_v105 = _dafny.SeqOf((_this).F24())
			var _1870_v106 _dafny.MultiSet
			_ = _1870_v106
			_1870_v106 = _dafny.MultiSetOf(false, (_this).F24(), (_this).F24(), (_this).F24(), (_this).F24())
			var _1871_v107 _dafny.CodePoint
			_ = _1871_v107
			_1871_v107 = _dafny.CodePoint('x')
			var _1872_v108 D10
			_ = _1872_v108
			_1872_v108 = Companion_D10_.Create_DC34_(_1871_v107, (_this).F24(), (_this).F24(), (_this).F24(), (_this).F24())
			var _1873_v109 D11
			_ = _1873_v109
			_1873_v109 = Companion_D11_.Create_DC39_((_this).F24())
			var _1874_v110 _dafny.Set
			_ = _1874_v110
			_1874_v110 = _dafny.SetOf((_dafny.MultiSetFromSeq(_1869_v105)).Cardinality())
			var _1875_v111 _dafny.Map
			_ = _1875_v111
			_1875_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1874_v110).Cardinality(), (_this).F24())
			var _1876_v112 _dafny.Map
			_ = _1876_v112
			_1876_v112 = _1875_v111
			var _1877_v113 _dafny.Sequence
			_ = _1877_v113
			_1877_v113 = _dafny.UnicodeSeqOfUtf8Bytes("lsxnhwfp")
			var _1878_v114 _dafny.Array
			_ = _1878_v114
			var _nwElement0_57 bool = (_1869_v105).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1870_v106).Contains((_this).F24()) {
					return (_1870_v106).Multiplicity((_this).F24())
				}
				return _1747_v0
			})(), _dafny.IntOfUint32((_1869_v105).Cardinality()))).Uint32()).(bool)
			_ = _nwElement0_57
			var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(22))
			_ = _nw272
			(_nw272).ArraySet1(_nwElement0_57, 0)
			(_nw272).ArraySet1((_this).F24(), 1)
			(_nw272).ArraySet1((_this).F24(), 2)
			(_nw272).ArraySet1((_this).F24(), 3)
			(_nw272).ArraySet1((func() bool {
				if (_this).F24() {
					return (_this).F24()
				}
				return (_this).F24()
			})(), 4)
			(_nw272).ArraySet1((func() bool {
				if (_this).F24() {
					return (_this).F24()
				}
				return (_this).F24()
			})(), 5)
			(_nw272).ArraySet1(true, 6)
			(_nw272).ArraySet1((_this).F24(), 7)
			(_nw272).ArraySet1((_this).F24(), 8)
			(_nw272).ArraySet1((_this).F24(), 9)
			(_nw272).ArraySet1((_this).F24(), 10)
			(_nw272).ArraySet1((_this).F24(), 11)
			(_nw272).ArraySet1(!((_1872_v108).Dtor_cf57()), 12)
			(_nw272).ArraySet1((_this).F24(), 13)
			(_nw272).ArraySet1((_1747_v0).Cmp(_1747_v0) < 0, 14)
			(_nw272).ArraySet1((_1873_v109).Dtor_cf66(), 15)
			(_nw272).ArraySet1(!((_this).F24()), 16)
			(_nw272).ArraySet1((_1747_v0).Cmp(_dafny.IntOfInt64(-636)) <= 0, 17)
			(_nw272).ArraySet1((func() bool {
				if (_this).F24() {
					return (_this).F24()
				}
				return (_this).F24()
			})(), 18)
			(_nw272).ArraySet1(true, 19)
			(_nw272).ArraySet1(Companion_Default___.Fm37(_1869_v105, (_1876_v112), _1877_v113, globalState), 20)
			(_nw272).ArraySet1(true, 21)
			_1878_v114 = _nw272
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_1878_v114), 0))
			_ = _index270
			(_1878_v114).ArraySet1(!((_this).F24()), (_index270).Int())
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_1878_v114), 0))
			_ = _index271
			(_1878_v114).ArraySet1((_1874_v110).IsSubsetOf(_1874_v110), (_index271).Int())
			var _1879_v115 *C3
			_ = _1879_v115
			var _nw273 *C3 = New_C3_()
			_ = _nw273
			_nw273.Ctor__((_this).F22())
			_1879_v115 = _nw273
			_1879_v115 = _1879_v115
			_1747_v0 = _dafny.IntOfInt64(861)
		} else if _source25.Is_DC51() {
			var _1880___mcc_h0 bool = _source25.Get_().(D15_DC51).Cf82
			_ = _1880___mcc_h0
			var _1881___mcc_h1 bool = _source25.Get_().(D15_DC51).Cf83
			_ = _1881___mcc_h1
			var _1882_cf83 bool = _1881___mcc_h1
			_ = _1882_cf83
			var _1883_cf82 bool = _1880___mcc_h0
			_ = _1883_cf82
			var _1884_v116 _dafny.Map
			_ = _1884_v116
			_1884_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_cf83, false)
			var _1885_v117 _dafny.Array
			_ = _1885_v117
			var _len0_49 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_49
			var _nw274 _dafny.Array
			_ = _nw274
			if _len0_49.Cmp(_dafny.Zero) == 0 {
				_nw274 = _dafny.NewArray(_len0_49)
			} else {
				var _init49 func(_dafny.Int) _dafny.Int = func(_1886_i10 _dafny.Int) _dafny.Int {
					return (_1886_i10).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pvvl")).Cardinality()))
				}
				_ = _init49
				var _element0_49 = _init49(_dafny.Zero)
				_ = _element0_49
				_nw274 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
				(_nw274).ArraySet1(_element0_49, 0)
				var _nativeLen0_49 = (_len0_49).Int()
				_ = _nativeLen0_49
				for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
					(_nw274).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
				}
			}
			_1885_v117 = _nw274
			var _1887_v118 _dafny.Map
			_ = _1887_v118
			_1887_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1885_v117, _1882_cf83)
			_1884_v116 = (_1884_v116).Update((func() bool {
				if (_1887_v118).Contains(_1885_v117) {
					return (_1887_v118).Get(_1885_v117).(bool)
				}
				return (_this).F24()
			})(), _1882_cf83)
			var _1888_v119 T0
			_ = _1888_v119
			var _nw275 *C1 = New_C1_()
			_ = _nw275
			_nw275.Ctor__(_dafny.IntOfInt64(779), (_this).F22())
			_1888_v119 = _nw275
			var _1889_v120 _dafny.Map
			_ = _1889_v120
			_1889_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, _1888_v119)
			var _1890_v121 _dafny.Sequence
			_ = _1890_v121
			_1890_v121 = _dafny.SeqOf(true, !((_this).F24()), false)
			var _1891_v122 _dafny.Array
			_ = _1891_v122
			var _nwElement0_58 T0 = _1888_v119
			_ = _nwElement0_58
			var _nw276 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(23))
			_ = _nw276
			(_nw276).ArraySet1(_nwElement0_58, 0)
			(_nw276).ArraySet1(_1888_v119, 1)
			(_nw276).ArraySet1(_1888_v119, 2)
			(_nw276).ArraySet1(_1888_v119, 3)
			(_nw276).ArraySet1(_1888_v119, 4)
			(_nw276).ArraySet1(_1888_v119, 5)
			(_nw276).ArraySet1(_1888_v119, 6)
			(_nw276).ArraySet1((func() T0 {
				if (_1889_v120).Contains(_dafny.IntOfUint32((_1890_v121).Cardinality())) {
					return (_1889_v120).Get(_dafny.IntOfUint32((_1890_v121).Cardinality())).(T0)
				}
				return _1888_v119
			})(), 7)
			(_nw276).ArraySet1(_1888_v119, 8)
			(_nw276).ArraySet1(_1888_v119, 9)
			(_nw276).ArraySet1(_1888_v119, 10)
			(_nw276).ArraySet1(_1888_v119, 11)
			(_nw276).ArraySet1(_1888_v119, 12)
			(_nw276).ArraySet1(_1888_v119, 13)
			(_nw276).ArraySet1((func() T0 {
				if (_this).F24() {
					return _1888_v119
				}
				return _1888_v119
			})(), 14)
			(_nw276).ArraySet1(_1888_v119, 15)
			(_nw276).ArraySet1(_1888_v119, 16)
			(_nw276).ArraySet1(_1888_v119, 17)
			(_nw276).ArraySet1(_1888_v119, 18)
			(_nw276).ArraySet1(_1888_v119, 19)
			(_nw276).ArraySet1(_1888_v119, 20)
			(_nw276).ArraySet1(_1888_v119, 21)
			(_nw276).ArraySet1(_1888_v119, 22)
			_1891_v122 = _nw276
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1891_v122), 0))
			_ = _index272
			(_1891_v122).ArraySet1(_1888_v119, (_index272).Int())
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_1891_v122), 0))
			_ = _index273
			(_1891_v122).ArraySet1(_1888_v119, (_index273).Int())
			_1883_cf82 = !(_1882_cf83) || (false)
			var _1892_v123 *C2
			_ = _1892_v123
			var _nw277 *C2 = New_C2_()
			_ = _nw277
			_nw277.Ctor__(_1747_v0, (_this).F22())
			_1892_v123 = _nw277
			_1892_v123 = _1892_v123
		} else if _source25.Is_DC52() {
			var _1893_v124 _dafny.Array
			_ = _1893_v124
			var _nw278 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw278
			_1893_v124 = _nw278
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))
			_ = _index274
			(_1893_v124).ArraySet1(true, (_index274).Int())
			var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))
			_ = _index275
			(_1893_v124).ArraySet1((_dafny.IntOfInt64(312)).Cmp(_1747_v0) < 0, (_index275).Int())
			var _1894_v125 _dafny.Sequence
			_ = _1894_v125
			_1894_v125 = _dafny.UnicodeSeqOfUtf8Bytes("elp")
			var _1895_v126 _dafny.CodePoint
			_ = _1895_v126
			_1895_v126 = _dafny.CodePoint('d')
			var _1896_v127 _dafny.Map
			_ = _1896_v127
			_1896_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool), _1895_v126)
			var _1897_v128 _dafny.Map
			_ = _1897_v128
			_1897_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(486), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1894_v125).Cardinality()))), (func() _dafny.Map {
				if false {
					return _1896_v127
				}
				return _1896_v127
			})())
			var _1898_v129 _dafny.MultiSet
			_ = _1898_v129
			_1898_v129 = _dafny.MultiSetOf(_1867_v103)
			var _1899_v130 _dafny.Sequence
			_ = _1899_v130
			_1899_v130 = _dafny.SeqOf((_1898_v129).Cardinality(), (_dafny.Zero).Minus(_1747_v0))
			var _1900_v131 _dafny.Sequence
			_ = _1900_v131
			_1900_v131 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool)), _1895_v126))
			var _1901_v132 _dafny.MultiSet
			_ = _1901_v132
			_1901_v132 = _dafny.MultiSetOf(_1747_v0, _1747_v0)
			_1897_v128 = (_1897_v128).Update(_1899_v130, (_1900_v131).Select((Companion_Default___.SafeIndex((_1901_v132).Cardinality(), _dafny.IntOfUint32((_1900_v131).Cardinality()))).Uint32()).(_dafny.Map))
			if (_this).F24() {
				var _1902_v133 *C3
				_ = _1902_v133
				var _nw279 *C3 = New_C3_()
				_ = _nw279
				_nw279.Ctor__(_dafny.Companion_Sequence_.Update((_this).F22(), (Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32(((_this).F22()).Cardinality()))).Uint32(), _1893_v124))
				_1902_v133 = _nw279
				var _1903_v134 _dafny.Array
				_ = _1903_v134
				var _nw280 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw280
				_1903_v134 = _nw280
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1903_v134), 0))
				_ = _index276
				(_1903_v134).ArraySet1((_1747_v0).Minus(_1747_v0), (_index276).Int())
				var _1904_v135 _dafny.Sequence
				_ = _1904_v135
				_1904_v135 = _dafny.SeqOf((_this).F24())
				var _1905_v136 _dafny.Map
				_ = _1905_v136
				_1905_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, (_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool))
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1903_v134), 0))
				_ = _index277
				var _rhs313 _dafny.Sequence = _dafny.SeqOf((func() bool {
					if false {
						return (_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool)
					}
					return Companion_Default___.Fm37(_1904_v135, _1905_v136, _1894_v125, globalState)
				})(), ((_this).F24()) || ((_this).F24()))
				_ = _rhs313
				var _rhs314 _dafny.Int = (_1747_v0).Minus(_1747_v0)
				_ = _rhs314
				var _rhs315 _dafny.Int = _1747_v0
				_ = _rhs315
				var _rhs316 _dafny.Array = _1893_v124
				_ = _rhs316
				var _rhs317 _dafny.Int = _1747_v0
				_ = _rhs317
				var _lhs255 *GlobalState = globalState
				_ = _lhs255
				var _lhs256 _dafny.Array = _1903_v134
				_ = _lhs256
				var _lhs257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1903_v134), 0))
				_ = _lhs257
				var _lhs258 *GlobalState = globalState
				_ = _lhs258
				_lhs255.F5 = _rhs313
				(_lhs256).ArraySet1(_rhs314, (_lhs257).Int())
				_lhs258.F8 = _rhs315
				_1893_v124 = _rhs316
				_1747_v0 = _rhs317
				_1895_v126 = _1895_v126
				r1 = (_1899_v130).Select((Companion_Default___.SafeIndex((_1903_v134).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1903_v134), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1899_v130).Cardinality()))).Uint32()).(_dafny.Int)
				_1895_v126 = _1895_v126
			} else {
				var _1906_v137 _dafny.Sequence
				_ = _1906_v137
				_1906_v137 = _dafny.SeqOf((_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool))
				var _1907_v138 _dafny.Map
				_ = _1907_v138
				_1907_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.IntOfInt64(-419)).Minus((_dafny.Zero).Minus(_1747_v0))), _dafny.Companion_Sequence_.Concatenate(_1906_v137, _1906_v137))
				var _1908_v139 _dafny.Array
				_ = _1908_v139
				var _nw281 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw281
				_1908_v139 = _nw281
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1908_v139), 0))
				_ = _index278
				(_1908_v139).ArraySet1(_1747_v0, (_index278).Int())
				var _1909_v140 _dafny.Map
				_ = _1909_v140
				_1909_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1895_v126, _1747_v0)
				var _1910_v141 _dafny.Sequence
				_ = _1910_v141
				_1910_v141 = _dafny.SeqOf(_1909_v140)
				var _1911_v142 _dafny.Map
				_ = _1911_v142
				_1911_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(367), (_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool))
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1908_v139), 0))
				_ = _index279
				var _rhs318 _dafny.Sequence = _1906_v137
				_ = _rhs318
				var _rhs319 _dafny.Map = (_1907_v138).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, _1906_v137))
				_ = _rhs319
				var _rhs320 _dafny.Int = Companion_Default___.SafeModuloInt(((_1910_v141).Select((Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32((_1910_v141).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), (_1911_v142).Cardinality())
				_ = _rhs320
				var _rhs321 bool = ((_1901_v132).Cardinality()).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1894_v125).Cardinality()), _1747_v0)) != 0
				_ = _rhs321
				var _lhs259 *GlobalState = globalState
				_ = _lhs259
				var _lhs260 _dafny.Array = _1908_v139
				_ = _lhs260
				var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1908_v139), 0))
				_ = _lhs261
				_lhs259.F5 = _rhs318
				_1907_v138 = _rhs319
				(_lhs260).ArraySet1(_rhs320, (_lhs261).Int())
				r0 = _rhs321
				var _1912_v143 _dafny.Sequence
				_ = _1912_v143
				_1912_v143 = _dafny.SeqOf(_1894_v125, _1894_v125)
				var _1913_v144 _dafny.Map
				_ = _1913_v144
				_1913_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool), (_this).F24())
				var _1914_v145 D9
				_ = _1914_v145
				_1914_v145 = Companion_D9_.Create_DC31_((_1913_v144).Cardinality(), _1893_v124)
				var _1915_v146 D9
				_ = _1915_v146
				_1915_v146 = Companion_D9_.Create_DC32_(_1914_v145)
				var _1916_v147 _dafny.Map
				_ = _1916_v147
				_1916_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1912_v143, Companion_D9_.Create_DC32_(_1915_v146))
				var _1917_v148 D9
				_ = _1917_v148
				_1917_v148 = Companion_D9_.Create_DC32_((func() D9 {
					if (_1916_v147).Contains(_1912_v143) {
						return (_1916_v147).Get(_1912_v143).(D9)
					}
					return _1915_v146
				})())
				var _1918_v149 D9
				_ = _1918_v149
				_1918_v149 = Companion_D9_.Create_DC32_(_1914_v145)
				var _1919_v150 _dafny.Map
				_ = _1919_v150
				_1919_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v149, (_this).F24())
				var _1920_v151 _dafny.Sequence
				_ = _1920_v151
				_1920_v151 = _dafny.SeqOf(_1919_v150, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v149, (_this).F24()), (_1919_v150).Merge(_1919_v150), _1919_v150)
				var _pat_let_tv67 = _1915_v146
				_ = _pat_let_tv67
				var _rhs322 _dafny.Sequence = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let38_0 D9) D9 {
					return func(_1921_dt__update__tmp_h1 D9) D9 {
						return func(_pat_let39_0 D9) D9 {
							return func(_1922_dt__update_hcf53_h0 D9) D9 {
								return Companion_D9_.Create_DC32_(_1922_dt__update_hcf53_h0)
							}(_pat_let39_0)
						}(_pat_let_tv67)
					}(_pat_let38_0)
				}(Companion_D9_.Create_DC32_(_1915_v146)), true)).Merge(_1919_v150))
				_ = _rhs322
				var _rhs323 bool = (_this).F24()
				_ = _rhs323
				var _rhs324 _dafny.CodePoint = _dafny.CodePoint('m')
				_ = _rhs324
				var _lhs262 *GlobalState = globalState
				_ = _lhs262
				_1920_v151 = _rhs322
				_lhs262.F10 = _rhs323
				_1895_v126 = _rhs324
				var _1923_v152 _dafny.MultiSet
				_ = _1923_v152
				_1923_v152 = _dafny.MultiSetOf(true)
				_1895_v126 = (_1894_v125).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1923_v152).Contains((_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool)) {
						return (_1923_v152).Multiplicity((_1893_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_1893_v124), 0))).Int()).(bool))
					}
					return (_1908_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1908_v139), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((_1894_v125).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_1894_v125 = _1894_v125
				var _1924_v153 _dafny.Set
				_ = _1924_v153
				_1924_v153 = _dafny.SetOf(_1747_v0)
				(globalState).F13 = _1924_v153
			}
			var _1925_v154 _dafny.Array
			_ = _1925_v154
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_50
			var _nw282 _dafny.Array
			_ = _nw282
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw282 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Sequence = func(_1926_i11 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F24(), (_this).F24())
				}
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw282 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw282).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw282).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_1925_v154 = _nw282
			var _1927_v155 _dafny.Map
			_ = _1927_v155
			_1927_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1899_v130, _1925_v154)
			var _1928_v156 _dafny.Array
			_ = _1928_v156
			var _nwElement0_59 _dafny.Array = _1925_v154
			_ = _nwElement0_59
			var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(26))
			_ = _nw283
			(_nw283).ArraySet1(_nwElement0_59, 0)
			(_nw283).ArraySet1(_1925_v154, 1)
			(_nw283).ArraySet1(_1925_v154, 2)
			(_nw283).ArraySet1(_1925_v154, 3)
			(_nw283).ArraySet1(_1925_v154, 4)
			(_nw283).ArraySet1(_1925_v154, 5)
			(_nw283).ArraySet1(_1925_v154, 6)
			(_nw283).ArraySet1(_1925_v154, 7)
			(_nw283).ArraySet1(_1925_v154, 8)
			(_nw283).ArraySet1(_1925_v154, 9)
			(_nw283).ArraySet1(_1925_v154, 10)
			(_nw283).ArraySet1(_1925_v154, 11)
			(_nw283).ArraySet1(_1925_v154, 12)
			(_nw283).ArraySet1(_1925_v154, 13)
			(_nw283).ArraySet1((func() _dafny.Array {
				if (_this).F24() {
					return _1925_v154
				}
				return (func() _dafny.Array {
					if (_1927_v155).Contains(_1899_v130) {
						return (_1927_v155).Get(_1899_v130).(_dafny.Array)
					}
					return _1925_v154
				})()
			})(), 14)
			(_nw283).ArraySet1(_1925_v154, 15)
			(_nw283).ArraySet1(_1925_v154, 16)
			(_nw283).ArraySet1(_1925_v154, 17)
			(_nw283).ArraySet1(_1925_v154, 18)
			(_nw283).ArraySet1(_1925_v154, 19)
			(_nw283).ArraySet1(_1925_v154, 20)
			(_nw283).ArraySet1(_1925_v154, 21)
			(_nw283).ArraySet1(_1925_v154, 22)
			(_nw283).ArraySet1(_1925_v154, 23)
			(_nw283).ArraySet1(_1925_v154, 24)
			(_nw283).ArraySet1(_1925_v154, 25)
			_1928_v156 = _nw283
			var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1928_v156), 0))
			_ = _index280
			(_1928_v156).ArraySet1(_1925_v154, (_index280).Int())
			var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_1928_v156), 0))
			_ = _index281
			(_1928_v156).ArraySet1(_1925_v154, (_index281).Int())
		} else if _source25.Is_DC49() {
			var _1929___mcc_h2 _dafny.Set = _source25.Get_().(D15_DC49).Cf81
			_ = _1929___mcc_h2
			var _1930_cf81 _dafny.Set = _1929___mcc_h2
			_ = _1930_cf81
			var _1931_v157 *C3
			_ = _1931_v157
			var _nw284 *C3 = New_C3_()
			_ = _nw284
			_nw284.Ctor__(_dafny.Companion_Sequence_.Concatenate((_this).F22(), (_this).F22()))
			_1931_v157 = _nw284
			var _1932_v158 _dafny.Array
			_ = _1932_v158
			var _nw285 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
			_ = _nw285
			_1932_v158 = _nw285
			_1932_v158 = _1932_v158
			var _rhs325 bool = (_this).F24()
			_ = _rhs325
			var _rhs326 _dafny.Int = ((_dafny.IntOfInt64(910)).Minus(_dafny.IntOfInt64(218))).Plus(_1747_v0)
			_ = _rhs326
			var _lhs263 *GlobalState = globalState
			_ = _lhs263
			var _lhs264 *GlobalState = globalState
			_ = _lhs264
			_lhs263.F0 = _rhs325
			_lhs264.F16 = _rhs326
			var _1933_v159 _dafny.MultiSet
			_ = _1933_v159
			_1933_v159 = _dafny.MultiSetOf(_1747_v0, _1747_v0, _1747_v0, _1747_v0)
			var _1934_v160 _dafny.Sequence
			_ = _1934_v160
			_1934_v160 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(974))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}(func(_1935_i13 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})))
			(globalState).F6 = (Companion_Default___.SafeDivisionInt(_1747_v0, (_1931_v157).Fm0(_1933_v159, _1747_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(861))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg100 _dafny.Int) interface{} {
					return coer100(arg100)
				}
			}((func(_1936_v0 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_1937_i12 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_1936_v0, _1936_v0)
				}
			})(_1747_v0))), (_this).F24(), globalState))).Minus(_dafny.IntOfUint32((_1934_v160).Cardinality()))
		} else {
			var _1938___mcc_h3 D15 = _source25.Get_().(D15_DC53).Cf84
			_ = _1938___mcc_h3
			var _1939_cf84 D15 = _1938___mcc_h3
			_ = _1939_cf84
			var _1940_v161 *C1
			_ = _1940_v161
			var _nw286 *C1 = New_C1_()
			_ = _nw286
			_nw286.Ctor__(_1747_v0, (_this).F22())
			_1940_v161 = _nw286
			var _1941_v162 _dafny.Sequence
			_ = _1941_v162
			_1941_v162 = _dafny.UnicodeSeqOfUtf8Bytes("wrcpc")
			var _1942_v163 D1
			_ = _1942_v163
			_1942_v163 = Companion_D1_.Create_DC3_(_1941_v162)
			var _source26 D1 = _1942_v163
			_ = _source26
			if _source26.Is_DC3() {
				var _1943___mcc_h4 _dafny.Sequence = _source26.Get_().(D1_DC3).Cf6
				_ = _1943___mcc_h4
				var _1944_cf6 _dafny.Sequence = _1943___mcc_h4
				_ = _1944_cf6
				var _1945_v164 D4
				_ = _1945_v164
				_1945_v164 = Companion_D4_.Create_DC15_(_dafny.IntOfUint32((_1941_v162).Cardinality()), (_dafny.Zero).Minus((_1940_v161).Fm34((_this).F24(), globalState)))
				_1945_v164 = _1945_v164
				_1941_v162 = _1944_cf6
				(globalState).F0 = (_this).F24()
				var _1946_v165 _dafny.Map
				_ = _1946_v165
				_1946_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1944_cf6, true)
				_1946_v165 = (_1946_v165).Update(_1944_cf6, (_this).F24())
			} else if _source26.Is_DC2() {
				var _1947___mcc_h5 _dafny.Sequence = _source26.Get_().(D1_DC2).Cf5
				_ = _1947___mcc_h5
				var _1948_cf5 _dafny.Sequence = _1947___mcc_h5
				_ = _1948_cf5
				_1948_cf5 = _dafny.UnicodeSeqOfUtf8Bytes("nqyii")
				var _1949_v166 T0
				_ = _1949_v166
				var _nw287 *C3 = New_C3_()
				_ = _nw287
				_nw287.Ctor__((_this).F22())
				_1949_v166 = _nw287
				var _1950_v167 _dafny.Array
				_ = _1950_v167
				var _nw288 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw288
				_1950_v167 = _nw288
				var _1951_v168 _dafny.Set
				_ = _1951_v168
				_1951_v168 = _dafny.SetOf(_1950_v167)
				var _1952_v169 D13
				_ = _1952_v169
				_1952_v169 = Companion_D13_.Create_DC46_((_this).F24(), _1949_v166, (_1940_v161).F30(), _1951_v168)
				var _1953_v170 _dafny.Map
				_ = _1953_v170
				_1953_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1948_cf5, (_1940_v161).F30())
				var _1954_v171 _dafny.Map
				_ = _1954_v171
				_1954_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v0, (_1940_v161).F30())
				var _1955_v172 _dafny.Sequence
				_ = _1955_v172
				_1955_v172 = _dafny.SeqOf((_dafny.Zero).Minus((_1954_v171).Cardinality()), (_1940_v161).F30(), (_1940_v161).F30())
				var _1956_v173 _dafny.Array
				_ = _1956_v173
				var _nwElement0_60 bool = (_this).F24()
				_ = _nwElement0_60
				var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(10))
				_ = _nw289
				(_nw289).ArraySet1(_nwElement0_60, 0)
				(_nw289).ArraySet1((_1952_v169).Dtor_cf73(), 1)
				(_nw289).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((func() _dafny.Int {
					if (_1953_v170).Contains(_dafny.UnicodeSeqOfUtf8Bytes("os")) {
						return (_1953_v170).Get(_dafny.UnicodeSeqOfUtf8Bytes("os")).(_dafny.Int)
					}
					return (_1940_v161).F30()
				})()), _1955_v172), 2)
				(_nw289).ArraySet1(true, 3)
				(_nw289).ArraySet1(false, 4)
				(_nw289).ArraySet1((_this).F24(), 5)
				(_nw289).ArraySet1((_this).F24(), 6)
				(_nw289).ArraySet1((_this).F24(), 7)
				(_nw289).ArraySet1(!((_this).F24()) || ((_this).F24()), 8)
				(_nw289).ArraySet1(Companion_Default___.Fm15((_this).F24(), globalState), 9)
				_1956_v173 = _nw289
				var _1957_v174 _dafny.Map
				_ = _1957_v174
				_1957_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), true)
				var _1958_v175 _dafny.Map
				_ = _1958_v175
				_1958_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1957_v174).Cardinality(), (_this).F24())
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))
				_ = _index282
				(_1956_v173).ArraySet1(!(Companion_Default___.Fm37(_dafny.SeqOf((_this).F24(), false, (_this).F24()), _1958_v175, _dafny.UnicodeSeqOfUtf8Bytes("a"), globalState)) || ((_this).F24()), (_index282).Int())
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))
				_ = _index283
				(_1956_v173).ArraySet1((_this).F24(), (_index283).Int())
				var _1959_v176 _dafny.Map
				_ = _1959_v176
				_1959_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1958_v175, (func() bool {
					if (_this).F24() {
						return (_1956_v173).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))).Int()).(bool)
					}
					return Companion_Default___.Fm15(false, globalState)
				})())
				_1959_v176 = (_1959_v176).Update(_1958_v175, (_1956_v173).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))).Int()).(bool))
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1950_v167), 0))
				_ = _index284
				(_1950_v167).ArraySet1((_1940_v161).F30(), (_index284).Int())
				var _1960_v177 _dafny.MultiSet
				_ = _1960_v177
				_1960_v177 = _dafny.MultiSetOf(_1747_v0)
				var _1961_v178 _dafny.Sequence
				_ = _1961_v178
				_1961_v178 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-361))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1962_v161 *C1) func(_dafny.Int) _dafny.Int {
					return func(_1963_i14 _dafny.Int) _dafny.Int {
						return (_1962_v161).F30()
					}
				})(_1940_v161))))
				var _1964_v179 _dafny.Set
				_ = _1964_v179
				_1964_v179 = _dafny.SetOf((_1940_v161).F30())
				var _1965_v180 _dafny.Sequence
				_ = _1965_v180
				_1965_v180 = _dafny.SeqOf(_1964_v179)
				var _1966_v181 _dafny.CodePoint
				_ = _1966_v181
				_1966_v181 = _dafny.CodePoint('k')
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1950_v167), 0))
				_ = _index285
				var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))
				_ = _index286
				var _rhs327 bool = ((_this).F24()) || ((_1956_v173).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))).Int()).(bool))
				_ = _rhs327
				var _rhs328 _dafny.Int = (_dafny.Zero).Minus((_1949_v166).Fm0(_1960_v177, _dafny.IntOfUint32((_1961_v178).Cardinality()), _1965_v180, (_this).F24(), globalState))
				_ = _rhs328
				var _rhs329 _dafny.Int = _1747_v0
				_ = _rhs329
				var _rhs330 _dafny.Int = _1747_v0
				_ = _rhs330
				var _rhs331 bool = _dafny.Companion_Sequence_.Contains(_1948_cf5, _1966_v181)
				_ = _rhs331
				var _lhs265 *GlobalState = globalState
				_ = _lhs265
				var _lhs266 _dafny.Array = _1950_v167
				_ = _lhs266
				var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_1950_v167), 0))
				_ = _lhs267
				var _lhs268 *GlobalState = globalState
				_ = _lhs268
				var _lhs269 *GlobalState = globalState
				_ = _lhs269
				var _lhs270 _dafny.Array = _1956_v173
				_ = _lhs270
				var _lhs271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1956_v173), 0))
				_ = _lhs271
				_lhs265.F0 = _rhs327
				(_lhs266).ArraySet1(_rhs328, (_lhs267).Int())
				_lhs268.F11 = _rhs329
				_lhs269.F6 = _rhs330
				(_lhs270).ArraySet1(_rhs331, (_lhs271).Int())
			} else {
				var _1967___mcc_h6 D1 = _source26.Get_().(D1_DC4).Cf7
				_ = _1967___mcc_h6
				var _1968_cf7 D1 = _1967___mcc_h6
				_ = _1968_cf7
				(globalState).F8 = _1747_v0
				var _1969_v182 D1
				_ = _1969_v182
				_1969_v182 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("cokvtw"))
				var _1970_v183 D1
				_ = _1970_v183
				_1970_v183 = Companion_D1_.Create_DC4_(_1969_v182)
				var _1971_v184 D1
				_ = _1971_v184
				_1971_v184 = Companion_D1_.Create_DC4_(_1969_v182)
				var _1972_v185 D1
				_ = _1972_v185
				_1972_v185 = Companion_D1_.Create_DC4_(_1970_v183)
				var _1973_v186 _dafny.Array
				_ = _1973_v186
				var _nwElement0_61 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1941_v162, _1941_v162)
				_ = _nwElement0_61
				var _nw290 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.One)
				_ = _nw290
				(_nw290).ArraySet1(_nwElement0_61, 0)
				_1973_v186 = _nw290
				var _1974_v187 _dafny.CodePoint
				_ = _1974_v187
				_1974_v187 = _dafny.CodePoint('y')
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_1973_v186), 0))
				_ = _index287
				(_1973_v186).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1941_v162, (Companion_Default___.SafeIndex((_1940_v161).F30(), _dafny.IntOfUint32((_1941_v162).Cardinality()))).Uint32(), _1974_v187), (Companion_Default___.SafeIndex(_1747_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1941_v162, (Companion_Default___.SafeIndex((_1940_v161).F30(), _dafny.IntOfUint32((_1941_v162).Cardinality()))).Uint32(), _1974_v187)).Cardinality()))).Uint32(), _1974_v187), (_index287).Int())
				var _1975_v188 _dafny.Array
				_ = _1975_v188
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_51
				var _nw291 _dafny.Array
				_ = _nw291
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw291 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.Int = (func(_1976_v161 *C1) func(_dafny.Int) _dafny.Int {
						return func(_1977_i15 _dafny.Int) _dafny.Int {
							return (_1977_i15).Minus((_1976_v161).F30())
						}
					})(_1940_v161)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw291 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw291).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw291).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1975_v188 = _nw291
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1975_v188), 0))
				_ = _index288
				(_1975_v188).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("krpsp")).Cardinality()), (_index288).Int())
				var _1978_v189 _dafny.MultiSet
				_ = _1978_v189
				_1978_v189 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1941_v162).Cardinality()), _dafny.IntOfInt64(999), _1747_v0)
				var _1979_v190 _dafny.Map
				_ = _1979_v190
				_1979_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(391))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg102 _dafny.Int) interface{} {
						return coer102(arg102)
					}
				}((func(_1980_v187 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1981_i16 _dafny.Int) _dafny.CodePoint {
						return _1980_v187
					}
				})(_1974_v187)))).Cardinality()))
				var _1982_v191 _dafny.Map
				_ = _1982_v191
				_1982_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1978_v189).Contains(_dafny.IntOfInt64(232)) {
						return (_1978_v189).Multiplicity(_dafny.IntOfInt64(232))
					}
					return (_1940_v161).F30()
				})(), _1979_v190)
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_1973_v186), 0))
				_ = _index289
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1975_v188), 0))
				_ = _index290
				var _rhs332 _dafny.Int = Companion_Default___.SafeDivisionInt((_1940_v161).F30(), (_1940_v161).F30())
				_ = _rhs332
				var _rhs333 D1 = Companion_D1_.Create_DC4_(_1970_v183)
				_ = _rhs333
				var _rhs334 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("cbjw")
				_ = _rhs334
				var _rhs335 _dafny.Int = (_1940_v161).F30()
				_ = _rhs335
				var _rhs336 _dafny.Map = func() _dafny.Map {
					var _coll73 = _dafny.NewMapBuilder()
					_ = _coll73
					for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(249), _dafny.IntOfInt64(61))); ; {
						_compr_73, _ok77 := _iter77()
						if !_ok77 {
							break
						}
						var _1983_v192 _dafny.Int
						_1983_v192 = interface{}(_compr_73).(_dafny.Int)
						if ((_dafny.IntOfInt64(249)).Cmp(_1983_v192) <= 0) && ((_1983_v192).Cmp(_dafny.IntOfInt64(61)) < 0) {
							_coll73.Add(Companion_Default___.SafeDivisionInt(_1983_v192, (_1940_v161).F30()), ((_1979_v190).Update((_this).F24(), _1747_v0)).Merge(_1979_v190))
						}
					}
					return _coll73.ToMap()
				}()
				_ = _rhs336
				var _lhs272 *GlobalState = globalState
				_ = _lhs272
				var _lhs273 _dafny.Array = _1973_v186
				_ = _lhs273
				var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_1973_v186), 0))
				_ = _lhs274
				var _lhs275 _dafny.Array = _1975_v188
				_ = _lhs275
				var _lhs276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1975_v188), 0))
				_ = _lhs276
				_lhs272.F16 = _rhs332
				_1972_v185 = _rhs333
				(_lhs273).ArraySet1(_rhs334, (_lhs274).Int())
				(_lhs275).ArraySet1(_rhs335, (_lhs276).Int())
				_1982_v191 = _rhs336
				var _1984_v193 D11
				_ = _1984_v193
				_1984_v193 = Companion_D11_.Create_DC36_((_this).F25())
				var _1985_v194 D11
				_ = _1985_v194
				_1985_v194 = Companion_D11_.Create_DC40_(_1984_v193)
				var _1986_v195 _dafny.Sequence
				_ = _1986_v195
				_1986_v195 = _dafny.SeqOf(_1985_v194)
				var _1987_v196 _dafny.Map
				_ = _1987_v196
				_1987_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1986_v195, _1986_v195), _1747_v0)
				_1987_v196 = (_1987_v196).Update(_dafny.Companion_Sequence_.Concatenate(_1986_v195, _dafny.SeqOf(_1985_v194)), (_1940_v161).F30())
				var _1988_v197 _dafny.Int
				_ = _1988_v197
				var _1989_v198 _dafny.Set
				_ = _1989_v198
				var _1990_v199 _dafny.Int
				_ = _1990_v199
				var _1991_v200 _dafny.Int
				_ = _1991_v200
				var _out64 _dafny.Int
				_ = _out64
				var _out65 _dafny.Set
				_ = _out65
				var _out66 _dafny.Int
				_ = _out66
				var _out67 _dafny.Int
				_ = _out67
				_out64, _out65, _out66, _out67 = (_1940_v161).M6(_dafny.IntOfInt64(104), (_this).F24(), ((_1975_v188).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1975_v188), 0))).Int()).(_dafny.Int)).Times((_1940_v161).F30()), (_1940_v161).F30(), globalState)
				_1988_v197 = _out64
				_1989_v198 = _out65
				_1990_v199 = _out66
				_1991_v200 = _out67
			}
			var _1992_v201 _dafny.Array
			_ = _1992_v201
			var _nw292 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
			_ = _nw292
			_1992_v201 = _nw292
			var _1993_v202 D11
			_ = _1993_v202
			_1993_v202 = Companion_D11_.Create_DC37_(_1992_v201)
			var _1994_v203 D11
			_ = _1994_v203
			_1994_v203 = Companion_D11_.Create_DC40_(_1993_v202)
			var _pat_let_tv68 = _1993_v202
			_ = _pat_let_tv68
			var _1995_v204 _dafny.Sequence
			_ = _1995_v204
			_1995_v204 = _dafny.SeqOf(_1994_v203, func(_pat_let40_0 D11) D11 {
				return func(_1996_dt__update__tmp_h2 D11) D11 {
					return func(_pat_let41_0 D11) D11 {
						return func(_1997_dt__update_hcf67_h0 D11) D11 {
							return Companion_D11_.Create_DC40_(_1997_dt__update_hcf67_h0)
						}(_pat_let41_0)
					}(_pat_let_tv68)
				}(_pat_let40_0)
			}(_1994_v203), _1994_v203, _1994_v203, func(_pat_let42_0 D11) D11 {
				return func(_1998_dt__update__tmp_h3 D11) D11 {
					return func(_pat_let43_0 D11) D11 {
						return func(_1999_dt__update_hcf67_h1 D11) D11 {
							return Companion_D11_.Create_DC40_(_1999_dt__update_hcf67_h1)
						}(_pat_let43_0)
					}(Companion_D11_.Create_DC39_(!((_this).F24())))
				}(_pat_let42_0)
			}(Companion_D11_.Create_DC40_(_1993_v202)))
			_1995_v204 = _dafny.Companion_Sequence_.Concatenate(_1995_v204, _1995_v204)
			var _2000_v205 _dafny.Map
			_ = _2000_v205
			_2000_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D15_.Create_DC50_(), (_this).F24())
			var _2001_v206 D15
			_ = _2001_v206
			_2001_v206 = Companion_D15_.Create_DC50_()
			_2000_v205 = (_2000_v205).Update(_2001_v206, (_this).F24())
		}
		r0 = (_this).F24()
		r1 = _1747_v0
		return r0, r1
	}
}
func (_this *C6) F24() bool {
	{
		return _this._f24
	}
}
func (_this *C6) F25() _dafny.Sequence {
	{
		return _this._f25
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f22 _dafny.Sequence
	_f23 _dafny.Set
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f22 = _dafny.EmptySeq
	_this._f23 = _dafny.EmptySet
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F22() _dafny.Sequence {
	return _this._f22
}
func (_this *C7) Ctor__(f23 _dafny.Set, f22 _dafny.Sequence) {
	{
		(_this)._f23 = f23
		(_this)._f22 = f22
	}
}
func (_this *C7) Fm0(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((((_this).F23()).Cardinality()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_dafny.IntOfInt64(44)))).Cardinality())).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(((_this).F23()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(134), false)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-324))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg103 _dafny.Int) interface{} {
				return coer103(arg103)
			}
		}(func(_2002_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Cardinality())).Cardinality()))), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(313))).Uint32(), func(coer104 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg104 _dafny.Int) interface{} {
				return coer104(arg104)
			}
		}(func(_2003_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(176)
		}))).Cardinality())))).Cardinality())))
	}
}
func (_this *C7) Fm1(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('q')
	}
}
func (_this *C7) M0(p0 bool, globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _2004_v0 _dafny.Sequence
		_ = _2004_v0
		_2004_v0 = _dafny.SeqOf(true, p0)
		var _2005_v1 _dafny.Int
		_ = _2005_v1
		_2005_v1 = _dafny.IntOfInt64(840)
		var _2006_v2 _dafny.Map
		_ = _2006_v2
		_2006_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2004_v0, _dafny.SetOf(_2005_v1))
		if !((func() _dafny.Set {
			if (_2006_v2).Contains(Companion_Default___.Fm2(globalState)) {
				return (_2006_v2).Get(Companion_Default___.Fm2(globalState)).(_dafny.Set)
			}
			return _dafny.SetOf(_2005_v1, _2005_v1)
		})()).Equals((_this).F23()) {
			var _2007_v3 _dafny.MultiSet
			_ = _2007_v3
			_2007_v3 = _dafny.MultiSetOf(_2005_v1)
			var _2008_v4 _dafny.Sequence
			_ = _2008_v4
			_2008_v4 = _dafny.SeqOf(_2007_v3, _2007_v3, _dafny.MultiSetOf(_2005_v1, _dafny.IntOfInt64(285)), (_2007_v3).Update(_2005_v1, Companion_Default___.Abs(_2005_v1)))
			var _2009_v5 *C6
			_ = _2009_v5
			var _nw293 *C6 = New_C6_()
			_ = _nw293
			_nw293.Ctor__(p0, _2008_v4, (_this).F22())
			_2009_v5 = _nw293
			var _2010_v6 _dafny.MultiSet
			_ = _2010_v6
			_2010_v6 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-446))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}(func(_2011_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})))
			var _2012_v7 _dafny.Sequence
			_ = _2012_v7
			_2012_v7 = _dafny.UnicodeSeqOfUtf8Bytes("wp")
			var _2013_v8 _dafny.Array
			_ = _2013_v8
			var _len0_52 _dafny.Int = _dafny.One
			_ = _len0_52
			var _nw294 _dafny.Array
			_ = _nw294
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw294 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) bool = (func(_2014_p0 bool) func(_dafny.Int) bool {
					return func(_2015_i1 _dafny.Int) bool {
						return _2014_p0
					}
				})(p0)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw294 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw294).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw294).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_2013_v8 = _nw294
			var _2016_v9 _dafny.Map
			_ = _2016_v9
			_2016_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2005_v1, _2013_v8)
			var _2017_v10 _dafny.Array
			_ = _2017_v10
			var _nwElement0_62 _dafny.Int = _2005_v1
			_ = _nwElement0_62
			var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(13))
			_ = _nw295
			(_nw295).ArraySet1(_nwElement0_62, 0)
			(_nw295).ArraySet1(_2005_v1, 1)
			(_nw295).ArraySet1(_2005_v1, 2)
			(_nw295).ArraySet1((func() _dafny.Int {
				if (_2010_v6).Contains(_2012_v7) {
					return (_2010_v6).Multiplicity(_2012_v7)
				}
				return _2005_v1
			})(), 3)
			(_nw295).ArraySet1(_2005_v1, 4)
			(_nw295).ArraySet1(_2005_v1, 5)
			(_nw295).ArraySet1(Companion_Default___.Fm7(_2005_v1, globalState), 6)
			(_nw295).ArraySet1(((_2016_v9).Update(_2005_v1, _2013_v8)).Cardinality(), 7)
			(_nw295).ArraySet1(_dafny.IntOfUint32((_2012_v7).Cardinality()), 8)
			(_nw295).ArraySet1(_dafny.IntOfInt64(484), 9)
			(_nw295).ArraySet1(_2005_v1, 10)
			(_nw295).ArraySet1(_2005_v1, 11)
			(_nw295).ArraySet1(_dafny.IntOfInt64(571), 12)
			_2017_v10 = _nw295
			var _2018_v11 _dafny.Set
			_ = _2018_v11
			_2018_v11 = _dafny.SetOf(_2017_v10, _2017_v10, _2017_v10)
			_2018_v11 = ((_2018_v11).Intersection(_dafny.SetOf(_2017_v10, _2017_v10, _2017_v10))).Intersection(_2018_v11)
			var _2019_v12 _dafny.Sequence
			_ = _2019_v12
			_2019_v12 = _dafny.SeqOf(_2005_v1)
			var _2020_v13 _dafny.Map
			_ = _2020_v13
			_2020_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_2009_v5).F24(), Companion_Default___.Fm15(p0, globalState)), _2019_v12)
			var _2021_v14 D10
			_ = _2021_v14
			_2021_v14 = Companion_D10_.Create_DC34_((_this).Fm1(p0, _2020_v13, globalState), (_2009_v5).F24(), false, !(p0), (_2009_v5).F24())
			var _2022_v15 _dafny.Map
			_ = _2022_v15
			_2022_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2005_v1, (_dafny.Zero).Minus(_2005_v1))
			var _2023_v16 _dafny.Map
			_ = _2023_v16
			_2023_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2021_v14, (func() _dafny.Int {
				if (_2022_v15).Contains(_2005_v1) {
					return (_2022_v15).Get(_2005_v1).(_dafny.Int)
				}
				return _dafny.IntOfInt64(490)
			})())
			_2023_v16 = (_2023_v16).Update(_2021_v14, _2005_v1)
			var _2024_v17 _dafny.Sequence
			_ = _2024_v17
			_2024_v17 = _dafny.SeqOf(_2012_v7)
			var _2025_v18 _dafny.Map
			_ = _2025_v18
			_2025_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2005_v1, _2024_v17)
			var _2026_v19 D12
			_ = _2026_v19
			_2026_v19 = Companion_D12_.Create_DC42_(_dafny.IntOfUint32((_2019_v12).Cardinality()))
			var _2027_v20 _dafny.Sequence
			_ = _2027_v20
			_2027_v20 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-56))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}((func(_2028_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_2029_i2 _dafny.Int) _dafny.Sequence {
					return _2028_v7
				}
			})(_2012_v7))), _2024_v17, (func() _dafny.Sequence {
				if (_2025_v18).Contains((_2026_v19).Dtor_cf69()) {
					return (_2025_v18).Get((_2026_v19).Dtor_cf69()).(_dafny.Sequence)
				}
				return _2024_v17
			})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}((func(_2030_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_2031_i3 _dafny.Int) _dafny.Sequence {
					return _2030_v7
				}
			})(_2012_v7))), _2024_v17)
			var _2032_v21 _dafny.Array
			_ = _2032_v21
			var _nwElement0_63 _dafny.Sequence = _2024_v17
			_ = _nwElement0_63
			var _nw296 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(8))
			_ = _nw296
			(_nw296).ArraySet1(_nwElement0_63, 0)
			(_nw296).ArraySet1(_2024_v17, 1)
			(_nw296).ArraySet1((func() _dafny.Sequence {
				if (_2025_v18).Contains(_2005_v1) {
					return (_2025_v18).Get(_2005_v1).(_dafny.Sequence)
				}
				return _2024_v17
			})(), 2)
			(_nw296).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_2027_v20).Select((Companion_Default___.SafeIndex(((_this).F23()).Cardinality(), _dafny.IntOfUint32((_2027_v20).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}(func(_2033_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), _2012_v7)), 3)
			(_nw296).ArraySet1(_2024_v17, 4)
			(_nw296).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2024_v17, _2024_v17), 5)
			(_nw296).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2024_v17, _2024_v17), 6)
			(_nw296).ArraySet1(_2024_v17, 7)
			_2032_v21 = _nw296
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_2032_v21), 0))
			_ = _index291
			(_2032_v21).ArraySet1(_2024_v17, (_index291).Int())
			var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_2032_v21), 0))
			_ = _index292
			var _rhs337 _dafny.Sequence = (func() _dafny.Sequence {
				if p0 {
					return _2019_v12
				}
				return _2019_v12
			})()
			_ = _rhs337
			var _rhs338 _dafny.Sequence = Companion_Default___.Fm53(globalState)
			_ = _rhs338
			var _lhs277 _dafny.Array = _2032_v21
			_ = _lhs277
			var _lhs278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_2032_v21), 0))
			_ = _lhs278
			_2019_v12 = _rhs337
			(_lhs277).ArraySet1(_rhs338, (_lhs278).Int())
			(globalState).F16 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_2022_v15).Contains(_2005_v1) {
					return (_2022_v15).Get(_2005_v1).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(Companion_Default___.Fm7((_dafny.MultiSetOf((_2009_v5).F24())).Cardinality(), globalState))
			})())
		} else {
			(globalState).F10 = !(p0)
			var _2034_v23 _dafny.Sequence
			_ = _2034_v23
			_2034_v23 = _dafny.SeqOf((func() _dafny.Set {
				var _coll74 = _dafny.NewBuilder()
				_ = _coll74
				for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-127), _dafny.IntOfInt64(193))); ; {
					_compr_74, _ok78 := _iter78()
					if !_ok78 {
						break
					}
					var _2035_v22 _dafny.Int
					_2035_v22 = interface{}(_compr_74).(_dafny.Int)
					if ((_dafny.IntOfInt64(-127)).Cmp(_2035_v22) <= 0) && ((_2035_v22).Cmp(_dafny.IntOfInt64(193)) < 0) {
						_coll74.Add(Companion_Default___.SafeModuloInt(_2035_v22, (_dafny.Zero).Minus(_2005_v1)))
					}
				}
				return _coll74.ToSet()
			}()).Cardinality())
			var _2036_v24 _dafny.Sequence
			_ = _2036_v24
			_2036_v24 = _dafny.SeqOf(_2034_v23)
			(globalState).F16 = (_2005_v1).Minus(_dafny.IntOfUint32((_2036_v24).Cardinality()))
			(globalState).F16 = Companion_Default___.SafeModuloInt(_2005_v1, _dafny.IntOfInt64(471))
			(globalState).F11 = _2005_v1
			var _2037_v25 _dafny.Array
			_ = _2037_v25
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_53
			var _nw297 _dafny.Array
			_ = _nw297
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw297 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) bool = func(_2038_i5 _dafny.Int) bool {
					return true
				}
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw297 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw297).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw297).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_2037_v25 = _nw297
			var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2037_v25), 0))
			_ = _index293
			(_2037_v25).ArraySet1(true, (_index293).Int())
			var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2037_v25), 0))
			_ = _index294
			(_2037_v25).ArraySet1((_2004_v0).Select((Companion_Default___.SafeIndex(_2005_v1, _dafny.IntOfUint32((_2004_v0).Cardinality()))).Uint32()).(bool), (_index294).Int())
		}
		var _2039_v26 D17
		_ = _2039_v26
		_2039_v26 = Companion_D17_.Create_DC55_((_this).F23())
		var _rhs339 _dafny.Int = (((_2039_v26).Dtor_cf86()).Union(((_this).F23()).Intersection((_this).F23()))).Cardinality()
		_ = _rhs339
		var _rhs340 _dafny.Set = (_this).F23()
		_ = _rhs340
		var _rhs341 _dafny.Int = _2005_v1
		_ = _rhs341
		var _lhs279 *GlobalState = globalState
		_ = _lhs279
		var _lhs280 *GlobalState = globalState
		_ = _lhs280
		_2005_v1 = _rhs339
		_lhs279.F13 = _rhs340
		_lhs280.F8 = _rhs341
		var _2040_v27 _dafny.Sequence
		_ = _2040_v27
		_2040_v27 = _dafny.UnicodeSeqOfUtf8Bytes("cfcba")
		var _2041_v28 _dafny.Sequence
		_ = _2041_v28
		_2041_v28 = _dafny.SeqOf(_dafny.IntOfUint32((_2040_v27).Cardinality()))
		var _2042_v29 D2
		_ = _2042_v29
		_2042_v29 = Companion_D2_.Create_DC7_(_2041_v28, p0, p0, (_dafny.Zero).Minus(_2005_v1))
		var _hi4 _dafny.Int = (_2005_v1).Plus((_2042_v29).Dtor_cf13())
		_ = _hi4
		for _2043_i6 := _2005_v1; _2043_i6.Cmp(_hi4) < 0; _2043_i6 = _2043_i6.Plus(_dafny.One) {
			if p0 {
				var _2044_v30 _dafny.Array
				_ = _2044_v30
				var _nw298 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw298
				_2044_v30 = _nw298
				var _2045_v31 _dafny.MultiSet
				_ = _2045_v31
				_2045_v31 = _dafny.MultiSetOf(p0, p0)
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))
				_ = _index295
				(_2044_v30).ArraySet1((_2045_v31).IsProperSubsetOf(_2045_v31), (_index295).Int())
				var _2046_v33 _dafny.CodePoint
				_ = _2046_v33
				_2046_v33 = _dafny.CodePoint('c')
				var _2047_v34 _dafny.Map
				_ = _2047_v34
				_2047_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2046_v33, _2043_i6)
				var _2048_v35 _dafny.MultiSet
				_ = _2048_v35
				_2048_v35 = _dafny.MultiSetOf((_2047_v34).Cardinality(), _2043_i6, (_dafny.Zero).Minus(_2043_i6), _2005_v1)
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))
				_ = _index296
				var _rhs342 bool = (false) == (p0)
				_ = _rhs342
				var _rhs343 bool = ((_this).F23()).IsSubsetOf(func() _dafny.Set {
					var _coll75 = _dafny.NewBuilder()
					_ = _coll75
					for _iter79 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_2041_v28, (Companion_Default___.SafeIndex(_2005_v1, _dafny.IntOfUint32((_2041_v28).Cardinality()))).Uint32(), _2005_v1)).Elements()); ; {
						_compr_75, _ok79 := _iter79()
						if !_ok79 {
							break
						}
						var _2049_v32 _dafny.Int
						_2049_v32 = interface{}(_compr_75).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_2041_v28, (Companion_Default___.SafeIndex(_2005_v1, _dafny.IntOfUint32((_2041_v28).Cardinality()))).Uint32(), _2005_v1), _2049_v32) {
							_coll75.Add(Companion_Default___.SafeModuloInt(_2049_v32, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kkgbxwxpm")).Cardinality()))).Cardinality())))
						}
					}
					return _coll75.ToSet()
				}())
				_ = _rhs343
				var _rhs344 _dafny.Int = (func() _dafny.Int {
					if (_2048_v35).IsSubsetOf(_dafny.MultiSetOf(_2043_i6)) {
						return _dafny.IntOfInt64(601)
					}
					return _2005_v1
				})()
				_ = _rhs344
				var _rhs345 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2041_v28, _dafny.Companion_Sequence_.Concatenate(_2041_v28, _2041_v28))
				_ = _rhs345
				var _lhs281 *GlobalState = globalState
				_ = _lhs281
				var _lhs282 _dafny.Array = _2044_v30
				_ = _lhs282
				var _lhs283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))
				_ = _lhs283
				var _lhs284 *GlobalState = globalState
				_ = _lhs284
				_lhs281.F0 = _rhs342
				(_lhs282).ArraySet1(_rhs343, (_lhs283).Int())
				_lhs284.F6 = _rhs344
				_2041_v28 = _rhs345
				var _2050_v36 _dafny.Map
				_ = _2050_v36
				_2050_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_2004_v0).Cardinality()))
				var _2051_v37 *C1
				_ = _2051_v37
				var _nw299 *C1 = New_C1_()
				_ = _nw299
				_nw299.Ctor__(((_2050_v36).Cardinality()).Times(_2005_v1), (_this).F22())
				_2051_v37 = _nw299
				_2051_v37 = _2051_v37
				var _2052_v38 _dafny.Map
				_ = _2052_v38
				_2052_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2044_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))).Int()).(bool), (_2004_v0).Select((Companion_Default___.SafeIndex((_2051_v37).F30(), _dafny.IntOfUint32((_2004_v0).Cardinality()))).Uint32()).(bool))
				_2052_v38 = _2052_v38
				var _2053_v39 _dafny.MultiSet
				_ = _2053_v39
				_2053_v39 = _dafny.MultiSetOf(_2044_v30)
				var _2054_v40 _dafny.Array
				_ = _2054_v40
				var _len0_54 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_54
				var _nw300 _dafny.Array
				_ = _nw300
				if _len0_54.Cmp(_dafny.Zero) == 0 {
					_nw300 = _dafny.NewArray(_len0_54)
				} else {
					var _init54 func(_dafny.Int) _dafny.MultiSet = (func(_2055_v33 _dafny.CodePoint) func(_dafny.Int) _dafny.MultiSet {
						return func(_2056_i7 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_2055_v33)).Intersection(_dafny.MultiSetOf(_2055_v33))
						}
					})(_2046_v33)
					_ = _init54
					var _element0_54 = _init54(_dafny.Zero)
					_ = _element0_54
					_nw300 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
					(_nw300).ArraySet1(_element0_54, 0)
					var _nativeLen0_54 = (_len0_54).Int()
					_ = _nativeLen0_54
					for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
						(_nw300).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
					}
				}
				_2054_v40 = _nw300
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_2054_v40), 0))
				_ = _index297
				(_2054_v40).ArraySet1(_dafny.MultiSetOf(_2046_v33, _2046_v33), (_index297).Int())
				var _2057_v41 _dafny.Array
				_ = _2057_v41
				var _nw301 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw301
				_2057_v41 = _nw301
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_2057_v41), 0))
				_ = _index298
				(_2057_v41).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()), (_index298).Int())
				var _2058_v42 T0
				_ = _2058_v42
				var _nw302 *C2 = New_C2_()
				_ = _nw302
				_nw302.Ctor__((_2048_v35).Cardinality(), (_this).F22())
				_2058_v42 = _nw302
				var _2059_v43 _dafny.Sequence
				_ = _2059_v43
				_2059_v43 = _dafny.SeqOf(_2041_v28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-755))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg109 _dafny.Int) interface{} {
						return coer109(arg109)
					}
				}((func(_2060_v37 *C1) func(_dafny.Int) _dafny.Int {
					return func(_2061_i8 _dafny.Int) _dafny.Int {
						return (_2060_v37).F30()
					}
				})(_2051_v37))))
				var _2062_v44 _dafny.Set
				_ = _2062_v44
				_2062_v44 = _dafny.SetOf((_2044_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))).Int()).(bool))
				var _2063_v45 _dafny.Map
				_ = _2063_v45
				_2063_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2058_v42, Companion_Default___.SafeModuloInt(_2005_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_2059_v43).Select((Companion_Default___.SafeIndex((_2062_v44).Cardinality(), _dafny.IntOfUint32((_2059_v43).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_2043_i6, _dafny.IntOfUint32(((_2059_v43).Select((Companion_Default___.SafeIndex((_2062_v44).Cardinality(), _dafny.IntOfUint32((_2059_v43).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), (_2051_v37).F30())).Cardinality())))
				var _2064_v46 _dafny.Map
				_ = _2064_v46
				_2064_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm23(_dafny.IntOfUint32((_2041_v28).Cardinality()), _2005_v1, p0, (_2044_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))).Int()).(bool), globalState)).Cardinality(), (_2044_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))).Int()).(bool))
				var _2065_v47 _dafny.MultiSet
				_ = _2065_v47
				_2065_v47 = _dafny.MultiSetOf(_2046_v33, _dafny.CodePoint('d'), _2046_v33)
				var _2066_v48 D14
				_ = _2066_v48
				_2066_v48 = Companion_D14_.Create_DC48_(_2043_i6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2044_v30, _dafny.IntOfInt64(-950)), p0)
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_2054_v40), 0))
				_ = _index299
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_2057_v41), 0))
				_ = _index300
				var _rhs346 _dafny.MultiSet = (_2053_v39).Update(_2044_v30, Companion_Default___.Abs((_dafny.Zero).Minus((_2051_v37).Fm34((func() bool {
					if (_2064_v46).Contains(_2005_v1) {
						return (_2064_v46).Get(_2005_v1).(bool)
					}
					return p0
				})(), globalState))))
				_ = _rhs346
				var _rhs347 _dafny.MultiSet = _2065_v47
				_ = _rhs347
				var _rhs348 _dafny.Int = (_dafny.Zero).Minus(_2005_v1)
				_ = _rhs348
				var _rhs349 _dafny.Map = ((func() _dafny.Map {
					if (_2044_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2044_v30), 0))).Int()).(bool) {
						return _2063_v45
					}
					return _2063_v45
				})()).Merge(_2063_v45)
				_ = _rhs349
				var _rhs350 _dafny.Int = (func() _dafny.Int {
					if (_2045_v31).Contains((_2045_v31).IsProperSubsetOf(_2045_v31)) {
						return (_2045_v31).Multiplicity((_2045_v31).IsProperSubsetOf(_2045_v31))
					}
					return Companion_Default___.SafeDivisionInt((_2066_v48).Dtor_cf78(), (_dafny.Zero).Minus(_2005_v1))
				})()
				_ = _rhs350
				var _lhs285 _dafny.Array = _2054_v40
				_ = _lhs285
				var _lhs286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_2054_v40), 0))
				_ = _lhs286
				var _lhs287 _dafny.Array = _2057_v41
				_ = _lhs287
				var _lhs288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_2057_v41), 0))
				_ = _lhs288
				var _lhs289 *GlobalState = globalState
				_ = _lhs289
				_2053_v39 = _rhs346
				(_lhs285).ArraySet1(_rhs347, (_lhs286).Int())
				(_lhs287).ArraySet1(_rhs348, (_lhs288).Int())
				_2063_v45 = _rhs349
				_lhs289.F16 = _rhs350
				var _2067_v49 _dafny.Array
				_ = _2067_v49
				var _nw303 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
				_ = _nw303
				_2067_v49 = _nw303
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_2067_v49), 0))
				_ = _index301
				(_2067_v49).ArraySet1(_2040_v27, (_index301).Int())
				var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_2067_v49), 0))
				_ = _index302
				(_2067_v49).ArraySet1(_2040_v27, (_index302).Int())
			} else {
				(globalState).F16 = Companion_Default___.SafeModuloInt(_2043_i6, _dafny.IntOfInt64(737))
				(globalState).F8 = (_2043_i6).Minus(_2005_v1)
				var _2068_v50 *C0
				_ = _2068_v50
				var _nw304 *C0 = New_C0_()
				_ = _nw304
				_nw304.Ctor__()
				_2068_v50 = _nw304
				(globalState).F5 = (func() _dafny.Sequence {
					if p0 {
						return _2004_v0
					}
					return _2004_v0
				})()
				var _2069_v51 *C4
				_ = _2069_v51
				var _nw305 *C4 = New_C4_()
				_ = _nw305
				_nw305.Ctor__(p0, (_this).F22())
				_2069_v51 = _nw305
			}
			var _2070_v52 _dafny.Map
			_ = _2070_v52
			_2070_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2043_i6, p0))
			var _2071_v53 _dafny.Map
			_ = _2071_v53
			_2071_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2005_v1, false)
			var _2072_v54 _dafny.Sequence
			_ = _2072_v54
			_2072_v54 = _dafny.SeqOf((func() _dafny.Map {
				if (_2070_v52).Contains(_2004_v0) {
					return (_2070_v52).Get(_2004_v0).(_dafny.Map)
				}
				return _2071_v53
			})())
			_2072_v54 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg110 _dafny.Int) interface{} {
					return coer110(arg110)
				}
			}((func(_2073_v53 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_2074_i9 _dafny.Int) _dafny.Map {
					return _2073_v53
				}
			})(_2071_v53))), _2072_v54)
			if p0 {
				(globalState).F10 = p0
				var _2075_v55 _dafny.Map
				_ = _2075_v55
				_2075_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_2040_v27).Cardinality())), _2043_i6))
				_2075_v55 = (_2075_v55).Update(p0, _dafny.IntOfUint32((_2040_v27).Cardinality()))
				(globalState).F11 = (_2041_v28).Select((Companion_Default___.SafeIndex(_2005_v1, _dafny.IntOfUint32((_2041_v28).Cardinality()))).Uint32()).(_dafny.Int)
				var _2076_v56 _dafny.Array
				_ = _2076_v56
				var _nwElement0_64 bool = p0
				_ = _nwElement0_64
				var _nw306 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(3))
				_ = _nw306
				(_nw306).ArraySet1(_nwElement0_64, 0)
				(_nw306).ArraySet1((p0) && (p0), 1)
				(_nw306).ArraySet1(p0, 2)
				_2076_v56 = _nw306
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2076_v56), 0))
				_ = _index303
				(_2076_v56).ArraySet1(true, (_index303).Int())
				var _2077_v57 T0
				_ = _2077_v57
				var _nw307 *C3 = New_C3_()
				_ = _nw307
				_nw307.Ctor__((_this).F22())
				_2077_v57 = _nw307
				var _2078_v58 _dafny.Map
				_ = _2078_v58
				_2078_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2005_v1).Cmp(_dafny.IntOfInt64(419)) == 0, p0)
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2076_v56), 0))
				_ = _index304
				var _rhs351 bool = (Companion_Default___.Fm7(_2005_v1, globalState)).Cmp((_dafny.MultiSetOf(_2076_v56)).Cardinality()) < 0
				_ = _rhs351
				var _rhs352 bool = (func() bool {
					if (_2078_v58).Contains(p0) {
						return (_2078_v58).Get(p0).(bool)
					}
					return p0
				})()
				_ = _rhs352
				var _rhs353 T0 = _2077_v57
				_ = _rhs353
				var _rhs354 bool = !(p0) || (p0)
				_ = _rhs354
				var _rhs355 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2004_v0, _2004_v0)
				_ = _rhs355
				var _lhs290 *GlobalState = globalState
				_ = _lhs290
				var _lhs291 _dafny.Array = _2076_v56
				_ = _lhs291
				var _lhs292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2076_v56), 0))
				_ = _lhs292
				var _lhs293 *GlobalState = globalState
				_ = _lhs293
				_lhs290.F0 = _rhs351
				(_lhs291).ArraySet1(_rhs352, (_lhs292).Int())
				_2077_v57 = _rhs353
				r0 = _rhs354
				_lhs293.F5 = _rhs355
				var _2079_v59 *C4
				_ = _2079_v59
				var _nw308 *C4 = New_C4_()
				_ = _nw308
				_nw308.Ctor__(p0, (_this).F22())
				_2079_v59 = _nw308
			} else {
				var _2080_v60 _dafny.Map
				_ = _2080_v60
				_2080_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_2043_i6)))
				var _2081_v61 _dafny.Array
				_ = _2081_v61
				var _nw309 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw309
				_2081_v61 = _nw309
				var _2082_v62 *C4
				_ = _2082_v62
				var _nw310 *C4 = New_C4_()
				_ = _nw310
				_nw310.Ctor__(!_dafny.Companion_Sequence_.Equal(_2041_v28, _dafny.SeqOf((_2080_v60).Cardinality(), _2043_i6)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2081_v61), _dafny.Companion_Sequence_.Update((_this).F22(), (Companion_Default___.SafeIndex(_2005_v1, _dafny.IntOfUint32(((_this).F22()).Cardinality()))).Uint32(), _2081_v61)))
				_2082_v62 = _nw310
				_2082_v62 = _2082_v62
				var _2083_v63 D10
				_ = _2083_v63
				_2083_v63 = Companion_D10_.Create_DC35_(_2043_i6, _2043_i6)
				(globalState).F8 = (func(_pat_let44_0 D10) D10 {
					return func(_2084_dt__update__tmp_h0 D10) D10 {
						return func(_pat_let45_0 _dafny.Int) D10 {
							return func(_2085_dt__update_hcf61_h0 _dafny.Int) D10 {
								return Companion_D10_.Create_DC35_((_2084_dt__update__tmp_h0).Dtor_cf60(), _2085_dt__update_hcf61_h0)
							}(_pat_let45_0)
						}(_2043_i6)
					}(_pat_let44_0)
				}(_2083_v63)).Dtor_cf60()
				(globalState).F10 = p0
				var _2086_v64 _dafny.Array
				_ = _2086_v64
				var _len0_55 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_55
				var _nw311 _dafny.Array
				_ = _nw311
				if _len0_55.Cmp(_dafny.Zero) == 0 {
					_nw311 = _dafny.NewArray(_len0_55)
				} else {
					var _init55 func(_dafny.Int) _dafny.Map = (func(_2087_p0 bool) func(_dafny.Int) _dafny.Map {
						return func(_2088_i10 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2087_p0, false)
						}
					})(p0)
					_ = _init55
					var _element0_55 = _init55(_dafny.Zero)
					_ = _element0_55
					_nw311 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
					(_nw311).ArraySet1(_element0_55, 0)
					var _nativeLen0_55 = (_len0_55).Int()
					_ = _nativeLen0_55
					for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
						(_nw311).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
					}
				}
				_2086_v64 = _nw311
				var _2089_v65 _dafny.Array
				_ = _2089_v65
				var _nwElement0_65 _dafny.Array = _2086_v64
				_ = _nwElement0_65
				var _nw312 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(3))
				_ = _nw312
				(_nw312).ArraySet1(_nwElement0_65, 0)
				(_nw312).ArraySet1(_2086_v64, 1)
				(_nw312).ArraySet1(_2086_v64, 2)
				_2089_v65 = _nw312
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_2089_v65), 0))
				_ = _index305
				(_2089_v65).ArraySet1((func() _dafny.Array {
					if _2082_v62.F28 {
						return _2086_v64
					}
					return _2086_v64
				})(), (_index305).Int())
				var _2090_v66 _dafny.MultiSet
				_ = _2090_v66
				_2090_v66 = _dafny.MultiSetOf(_2082_v62.F28, !(_2082_v62.F28), p0)
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_2089_v65), 0))
				_ = _index306
				var _rhs356 _dafny.Int = (_dafny.IntOfInt64(802)).Times(Companion_Default___.SafeModuloInt(_2005_v1, (_2090_v66).Cardinality()))
				_ = _rhs356
				var _rhs357 _dafny.Int = (_2090_v66).Cardinality()
				_ = _rhs357
				var _rhs358 _dafny.Array = (func() _dafny.Array {
					if _2082_v62.F28 {
						return _2086_v64
					}
					return _2086_v64
				})()
				_ = _rhs358
				var _rhs359 bool = p0
				_ = _rhs359
				var _rhs360 bool = p0
				_ = _rhs360
				var _lhs294 *GlobalState = globalState
				_ = _lhs294
				var _lhs295 *GlobalState = globalState
				_ = _lhs295
				var _lhs296 _dafny.Array = _2089_v65
				_ = _lhs296
				var _lhs297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_2089_v65), 0))
				_ = _lhs297
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				var _lhs299 *GlobalState = globalState
				_ = _lhs299
				_lhs294.F16 = _rhs356
				_lhs295.F16 = _rhs357
				(_lhs296).ArraySet1(_rhs358, (_lhs297).Int())
				_lhs298.F10 = _rhs359
				_lhs299.F0 = _rhs360
				_2041_v28 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14(_2043_i6, _2005_v1, globalState), _2041_v28)
			}
			(globalState).F0 = p0
		}
		var _2091_v67 _dafny.MultiSet
		_ = _2091_v67
		_2091_v67 = _dafny.MultiSetOf(p0, p0, p0)
		var _hi5 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfInt64(-607))).Plus((func() _dafny.Int {
			if (_2091_v67).Contains(p0) {
				return (_2091_v67).Multiplicity(p0)
			}
			return _2005_v1
		})())
		_ = _hi5
		for _2092_i11 := _dafny.IntOfInt64(621); _2092_i11.Cmp(_hi5) < 0; _2092_i11 = _2092_i11.Plus(_dafny.One) {
			var _2093_v68 _dafny.Map
			_ = _2093_v68
			_2093_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_2040_v27).Cardinality()))
			var _2094_v69 _dafny.Map
			_ = _2094_v69
			_2094_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2092_i11, _2005_v1)
			(globalState).F8 = (func() _dafny.Int {
				if (_2093_v68).Contains(p0) {
					return (_2093_v68).Get(p0).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_2094_v69).Contains(_2092_i11) {
						return (_2094_v69).Get(_2092_i11).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_2041_v28).Cardinality())
				})()
			})()
			var _2095_v70 _dafny.Array
			_ = _2095_v70
			var _len0_56 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_56
			var _nw313 _dafny.Array
			_ = _nw313
			if _len0_56.Cmp(_dafny.Zero) == 0 {
				_nw313 = _dafny.NewArray(_len0_56)
			} else {
				var _init56 func(_dafny.Int) _dafny.Sequence = (func(_2096_v1 _dafny.Int, _2097_p0 bool, _2098_i11 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_2099_i12 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_2096_v1, _2096_v1, _dafny.IntOfUint32((_dafny.SeqOf(_2097_p0, _2097_p0)).Cardinality()), (_dafny.SetOf(_2098_i11, _2098_i11, _2096_v1)).Cardinality())
					}
				})(_2005_v1, p0, _2092_i11)
				_ = _init56
				var _element0_56 = _init56(_dafny.Zero)
				_ = _element0_56
				_nw313 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
				(_nw313).ArraySet1(_element0_56, 0)
				var _nativeLen0_56 = (_len0_56).Int()
				_ = _nativeLen0_56
				for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
					(_nw313).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
				}
			}
			_2095_v70 = _nw313
			var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_2095_v70), 0))
			_ = _index307
			(_2095_v70).ArraySet1(_2041_v28, (_index307).Int())
			var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_2095_v70), 0))
			_ = _index308
			(_2095_v70).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2041_v28, (Companion_Default___.SafeIndex(Companion_Default___.Fm7(_2005_v1, globalState), _dafny.IntOfUint32((_2041_v28).Cardinality()))).Uint32(), _2092_i11), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(781))).Uint32(), func(coer111 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg111 _dafny.Int) interface{} {
					return coer111(arg111)
				}
			}((func(_2100_v69 _dafny.Map, _2101_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2102_i13 _dafny.Int) _dafny.Int {
					return (func() _dafny.Int {
						if (_2100_v69).Contains((_2100_v69).Cardinality()) {
							return (_2100_v69).Get((_2100_v69).Cardinality()).(_dafny.Int)
						}
						return _2101_v1
					})()
				}
			})(_2094_v69, _2005_v1)))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2041_v28, (Companion_Default___.SafeIndex(Companion_Default___.Fm7(_2005_v1, globalState), _dafny.IntOfUint32((_2041_v28).Cardinality()))).Uint32(), _2092_i11)).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_2005_v1)), (_index308).Int())
			if p0 {
				r3 = _2005_v1
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_2095_v70), 0))
				_ = _index309
				(_2095_v70).ArraySet1(_2041_v28, (_index309).Int())
				(globalState).F11 = _2092_i11
				r0 = true
				(globalState).F10 = p0
			} else {
				var _2103_v71 _dafny.Map
				_ = _2103_v71
				_2103_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2092_i11, p0)
				var _2104_v72 _dafny.Array
				_ = _2104_v72
				var _nwElement0_66 bool = p0
				_ = _nwElement0_66
				var _nw314 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(26))
				_ = _nw314
				(_nw314).ArraySet1(_nwElement0_66, 0)
				(_nw314).ArraySet1(p0, 1)
				(_nw314).ArraySet1(p0, 2)
				(_nw314).ArraySet1(p0, 3)
				(_nw314).ArraySet1(p0, 4)
				(_nw314).ArraySet1(Companion_Default___.Fm37(_2004_v0, _2103_v71, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(928))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}(func(_2105_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})), (Companion_Default___.SafeIndex(_2092_i11, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(928))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}(func(_2106_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}))).Cardinality()))).Uint32(), _dafny.CodePoint('n')), globalState), 5)
				(_nw314).ArraySet1((!(p0)) == (false), 6)
				(_nw314).ArraySet1(p0, 7)
				(_nw314).ArraySet1((_2005_v1).Cmp(_2005_v1) <= 0, 8)
				(_nw314).ArraySet1((_2005_v1).Cmp(_2005_v1) == 0, 9)
				(_nw314).ArraySet1(p0, 10)
				(_nw314).ArraySet1(p0, 11)
				(_nw314).ArraySet1(p0, 12)
				(_nw314).ArraySet1((func() bool {
					if (_2004_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("awkmh")).Cardinality()), _dafny.IntOfUint32((_2004_v0).Cardinality()))).Uint32()).(bool) {
						return p0
					}
					return true
				})(), 13)
				(_nw314).ArraySet1((p0) || (p0), 14)
				(_nw314).ArraySet1((_2005_v1).Cmp(_dafny.IntOfUint32((_2004_v0).Cardinality())) == 0, 15)
				(_nw314).ArraySet1((func() bool {
					if p0 {
						return p0
					}
					return p0
				})(), 16)
				(_nw314).ArraySet1((p0) && (p0), 17)
				(_nw314).ArraySet1(p0, 18)
				(_nw314).ArraySet1(!(!(p0)), 19)
				(_nw314).ArraySet1(p0, 20)
				(_nw314).ArraySet1((_2005_v1).Cmp(_2092_i11) >= 0, 21)
				(_nw314).ArraySet1((_2005_v1).Cmp(_2092_i11) > 0, 22)
				(_nw314).ArraySet1(p0, 23)
				(_nw314).ArraySet1(p0, 24)
				(_nw314).ArraySet1(p0, 25)
				_2104_v72 = _nw314
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_2104_v72), 0))
				_ = _index310
				(_2104_v72).ArraySet1(!(p0) || (p0), (_index310).Int())
				var _2107_v73 _dafny.Map
				_ = _2107_v73
				_2107_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2091_v67, _2004_v0)
				var _2108_v75 _dafny.Set
				_ = _2108_v75
				_2108_v75 = _dafny.SetOf(_2091_v67)
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_2104_v72), 0))
				_ = _index311
				(_2104_v72).ArraySet1((func() _dafny.Set {
					var _coll76 = _dafny.NewBuilder()
					_ = _coll76
					for _iter80 := _dafny.Iterate((_2107_v73).Keys().Elements()); ; {
						_compr_76, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _2109_v74 _dafny.MultiSet
						_2109_v74 = interface{}(_compr_76).(_dafny.MultiSet)
						if (_2107_v73).Contains(_2109_v74) {
							_coll76.Add(_2109_v74)
						}
					}
					return _coll76.ToSet()
				}()).IsDisjointFrom((_2108_v75).Union(_2108_v75)), (_index311).Int())
				(globalState).F8 = _2092_i11
				var _2110_v76 *C4
				_ = _2110_v76
				var _nw315 *C4 = New_C4_()
				_ = _nw315
				_nw315.Ctor__(p0, (_this).F22())
				_2110_v76 = _nw315
				var _2111_v77 _dafny.CodePoint
				_ = _2111_v77
				_2111_v77 = _dafny.CodePoint('u')
				_2111_v77 = _2111_v77
				(globalState).F16 = _2092_i11
			}
			var _2112_v78 _dafny.Array
			_ = _2112_v78
			var _nw316 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw316
			_2112_v78 = _nw316
			var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2112_v78), 0))
			_ = _index312
			(_2112_v78).ArraySet1(p0, (_index312).Int())
			var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2112_v78), 0))
			_ = _index313
			(_2112_v78).ArraySet1(p0, (_index313).Int())
		}
		var _2113_v79 _dafny.Map
		_ = _2113_v79
		_2113_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2040_v27)
		var _hi6 _dafny.Int = (_dafny.IntOfUint32((_2004_v0).Cardinality())).Times((_2113_v79).Cardinality())
		_ = _hi6
		for _2114_i15 := _2005_v1; _2114_i15.Cmp(_hi6) < 0; _2114_i15 = _2114_i15.Plus(_dafny.One) {
			(globalState).F10 = p0
			var _2115_v80 _dafny.Array
			_ = _2115_v80
			var _nw317 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw317
			_2115_v80 = _nw317
			var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))
			_ = _index314
			(_2115_v80).ArraySet1((!(p0)) || (false), (_index314).Int())
			var _2116_v81 _dafny.MultiSet
			_ = _2116_v81
			_2116_v81 = _dafny.MultiSetOf((_2114_i15).Minus(_2114_i15))
			var _2117_v82 _dafny.Map
			_ = _2117_v82
			_2117_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2005_v1)
			var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))
			_ = _index315
			var _rhs361 bool = !(true)
			_ = _rhs361
			var _rhs362 bool = true
			_ = _rhs362
			var _rhs363 _dafny.Int = (_dafny.Zero).Minus((_2116_v81).Cardinality())
			_ = _rhs363
			var _rhs364 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_2004_v0, _2004_v0, _2004_v0, _2004_v0, _dafny.Companion_Sequence_.Update(_2004_v0, (Companion_Default___.SafeIndex((_2117_v82).Cardinality(), _dafny.IntOfUint32((_2004_v0).Cardinality()))).Uint32(), p0)), (Companion_Default___.SafeIndex(_2005_v1, _dafny.IntOfUint32((_dafny.SeqOf(_2004_v0, _2004_v0, _2004_v0, _2004_v0, _dafny.Companion_Sequence_.Update(_2004_v0, (Companion_Default___.SafeIndex((_2117_v82).Cardinality(), _dafny.IntOfUint32((_2004_v0).Cardinality()))).Uint32(), p0))).Cardinality()))).Uint32(), _2004_v0)).Cardinality()))
			_ = _rhs364
			var _lhs300 *GlobalState = globalState
			_ = _lhs300
			var _lhs301 _dafny.Array = _2115_v80
			_ = _lhs301
			var _lhs302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))
			_ = _lhs302
			var _lhs303 *GlobalState = globalState
			_ = _lhs303
			var _lhs304 *GlobalState = globalState
			_ = _lhs304
			_lhs300.F0 = _rhs361
			(_lhs301).ArraySet1(_rhs362, (_lhs302).Int())
			_lhs303.F11 = _rhs363
			_lhs304.F6 = _rhs364
			if false {
				var _2118_v83 _dafny.Map
				_ = _2118_v83
				_2118_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_2118_v83 = (_2118_v83).Update((_2115_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))).Int()).(bool), (_2115_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))).Int()).(bool))
				r1 = _2115_v80
				var _2119_v84 D9
				_ = _2119_v84
				_2119_v84 = Companion_D9_.Create_DC31_(_2114_i15, _2115_v80)
				(globalState).F6 = (_2119_v84).Dtor_cf51()
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))
				_ = _index316
				(_2115_v80).ArraySet1(p0, (_index316).Int())
				var _2120_v85 _dafny.Map
				_ = _2120_v85
				_2120_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2114_i15, false)
				var _2121_v86 *C4
				_ = _2121_v86
				var _nw318 *C4 = New_C4_()
				_ = _nw318
				_nw318.Ctor__(Companion_Default___.Fm37(_2004_v0, _2120_v85, _2040_v27, globalState), (_this).F22())
				_2121_v86 = _nw318
			} else {
				var _2122_v87 _dafny.Array
				_ = _2122_v87
				var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw319
				_2122_v87 = _nw319
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2122_v87), 0))
				_ = _index317
				(_2122_v87).ArraySet1(Companion_Default___.SafeModuloInt(_2005_v1, (_2117_v82).Cardinality()), (_index317).Int())
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2122_v87), 0))
				_ = _index318
				(_2122_v87).ArraySet1(_2005_v1, (_index318).Int())
				_2122_v87 = _2122_v87
				var _2123_v88 *C1
				_ = _2123_v88
				var _nw320 *C1 = New_C1_()
				_ = _nw320
				_nw320.Ctor__((_2122_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2122_v87), 0))).Int()).(_dafny.Int), (_this).F22())
				_2123_v88 = _nw320
				var _2124_v89 _dafny.Array
				_ = _2124_v89
				var _nwElement0_67 *C1 = _2123_v88
				_ = _nwElement0_67
				var _nw321 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(27))
				_ = _nw321
				(_nw321).ArraySet1(_nwElement0_67, 0)
				(_nw321).ArraySet1(_2123_v88, 1)
				(_nw321).ArraySet1(_2123_v88, 2)
				(_nw321).ArraySet1(_2123_v88, 3)
				(_nw321).ArraySet1(_2123_v88, 4)
				(_nw321).ArraySet1(_2123_v88, 5)
				(_nw321).ArraySet1(_2123_v88, 6)
				(_nw321).ArraySet1(_2123_v88, 7)
				(_nw321).ArraySet1(_2123_v88, 8)
				(_nw321).ArraySet1(_2123_v88, 9)
				(_nw321).ArraySet1(_2123_v88, 10)
				(_nw321).ArraySet1(_2123_v88, 11)
				(_nw321).ArraySet1(_2123_v88, 12)
				(_nw321).ArraySet1(_2123_v88, 13)
				(_nw321).ArraySet1(_2123_v88, 14)
				(_nw321).ArraySet1(_2123_v88, 15)
				(_nw321).ArraySet1(_2123_v88, 16)
				(_nw321).ArraySet1(_2123_v88, 17)
				(_nw321).ArraySet1(_2123_v88, 18)
				(_nw321).ArraySet1(_2123_v88, 19)
				(_nw321).ArraySet1(_2123_v88, 20)
				(_nw321).ArraySet1(_2123_v88, 21)
				(_nw321).ArraySet1(_2123_v88, 22)
				(_nw321).ArraySet1(_2123_v88, 23)
				(_nw321).ArraySet1(_2123_v88, 24)
				(_nw321).ArraySet1(_2123_v88, 25)
				(_nw321).ArraySet1(_2123_v88, 26)
				_2124_v89 = _nw321
				var _2125_v90 D11
				_ = _2125_v90
				_2125_v90 = Companion_D11_.Create_DC37_(_2124_v89)
				var _2126_v91 D11
				_ = _2126_v91
				_2126_v91 = Companion_D11_.Create_DC40_(_2125_v90)
				var _2127_v92 _dafny.Sequence
				_ = _2127_v92
				_2127_v92 = _dafny.SeqOf(_2125_v90)
				var _2128_v93 _dafny.Map
				_ = _2128_v93
				_2128_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2122_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2122_v87), 0))).Int()).(_dafny.Int), (_2115_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))).Int()).(bool))
				var _2129_v94 *C5
				_ = _2129_v94
				var _nw322 *C5 = New_C5_()
				_ = _nw322
				_nw322.Ctor__(_2128_v93, _2005_v1)
				_2129_v94 = _nw322
				var _pat_let_tv69 = _2125_v90
				_ = _pat_let_tv69
				var _pat_let_tv70 = _2125_v90
				_ = _pat_let_tv70
				var _pat_let_tv71 = _2125_v90
				_ = _pat_let_tv71
				var _2130_v95 _dafny.Array
				_ = _2130_v95
				var _nwElement0_68 D11 = func(_pat_let46_0 D11) D11 {
					return func(_2131_dt__update__tmp_h1 D11) D11 {
						return func(_pat_let47_0 D11) D11 {
							return func(_2132_dt__update_hcf67_h0 D11) D11 {
								return Companion_D11_.Create_DC40_(_2132_dt__update_hcf67_h0)
							}(_pat_let47_0)
						}(_pat_let_tv69)
					}(_pat_let46_0)
				}(_2126_v91)
				_ = _nwElement0_68
				var _nw323 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(19))
				_ = _nw323
				(_nw323).ArraySet1(_nwElement0_68, 0)
				(_nw323).ArraySet1(_2126_v91, 1)
				(_nw323).ArraySet1(Companion_D11_.Create_DC40_(_2125_v90), 2)
				(_nw323).ArraySet1(_2126_v91, 3)
				(_nw323).ArraySet1(_2126_v91, 4)
				(_nw323).ArraySet1(_2126_v91, 5)
				(_nw323).ArraySet1(_2126_v91, 6)
				(_nw323).ArraySet1(_2126_v91, 7)
				(_nw323).ArraySet1(_2126_v91, 8)
				(_nw323).ArraySet1(func(_pat_let48_0 D11) D11 {
					return func(_2133_dt__update__tmp_h2 D11) D11 {
						return func(_pat_let49_0 D11) D11 {
							return func(_2134_dt__update_hcf67_h1 D11) D11 {
								return Companion_D11_.Create_DC40_(_2134_dt__update_hcf67_h1)
							}(_pat_let49_0)
						}(_pat_let_tv70)
					}(_pat_let48_0)
				}(_2126_v91), 9)
				(_nw323).ArraySet1(_2126_v91, 10)
				(_nw323).ArraySet1(_2126_v91, 11)
				(_nw323).ArraySet1(_2126_v91, 12)
				(_nw323).ArraySet1(func(_pat_let50_0 D11) D11 {
					return func(_2135_dt__update__tmp_h3 D11) D11 {
						return func(_pat_let51_0 D11) D11 {
							return func(_2136_dt__update_hcf67_h2 D11) D11 {
								return Companion_D11_.Create_DC40_(_2136_dt__update_hcf67_h2)
							}(_pat_let51_0)
						}(_pat_let_tv71)
					}(_pat_let50_0)
				}(_2126_v91), 13)
				(_nw323).ArraySet1((func() D11 {
					if (_2115_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))).Int()).(bool) {
						return _2126_v91
					}
					return _2126_v91
				})(), 14)
				(_nw323).ArraySet1(Companion_D11_.Create_DC40_((_2127_v92).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_2129_v94)).Cardinality(), _dafny.IntOfUint32((_2127_v92).Cardinality()))).Uint32()).(D11)), 15)
				(_nw323).ArraySet1(_2126_v91, 16)
				(_nw323).ArraySet1(Companion_D11_.Create_DC40_(_2125_v90), 17)
				(_nw323).ArraySet1(_2126_v91, 18)
				_2130_v95 = _nw323
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_2130_v95), 0))
				_ = _index319
				(_2130_v95).ArraySet1(_2126_v91, (_index319).Int())
				var _2137_v96 _dafny.Set
				_ = _2137_v96
				_2137_v96 = _dafny.SetOf(p0)
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_2130_v95), 0))
				_ = _index320
				var _rhs365 bool = !((_2115_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_2115_v80), 0))).Int()).(bool)) || ((_2137_v96).IsProperSubsetOf(_2137_v96))
				_ = _rhs365
				var _rhs366 D11 = _2126_v91
				_ = _rhs366
				var _rhs367 _dafny.Array = _2115_v80
				_ = _rhs367
				var _lhs305 *GlobalState = globalState
				_ = _lhs305
				var _lhs306 _dafny.Array = _2130_v95
				_ = _lhs306
				var _lhs307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_2130_v95), 0))
				_ = _lhs307
				_lhs305.F10 = _rhs365
				(_lhs306).ArraySet1(_rhs366, (_lhs307).Int())
				r1 = _rhs367
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_2122_v87), 0))
				_ = _index321
				(_2122_v87).ArraySet1(_dafny.IntOfInt64(625), (_index321).Int())
				(globalState).F16 = (_2129_v94).F27()
			}
			var _2138_v97 _dafny.CodePoint
			_ = _2138_v97
			_2138_v97 = _dafny.CodePoint('j')
			_2138_v97 = _2138_v97
		}
		var _2139_v98 _dafny.Array
		_ = _2139_v98
		var _len0_57 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_57
		var _nw324 _dafny.Array
		_ = _nw324
		if _len0_57.Cmp(_dafny.Zero) == 0 {
			_nw324 = _dafny.NewArray(_len0_57)
		} else {
			var _init57 func(_dafny.Int) _dafny.Sequence = (func(_2140_v28 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_2141_i16 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(93))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg114 _dafny.Int) interface{} {
							return coer114(arg114)
						}
					}((func(_2142_v28 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2143_i17 _dafny.Int) _dafny.Sequence {
							return _2142_v28
						}
					})(_2140_v28)))
				}
			})(_2041_v28)
			_ = _init57
			var _element0_57 = _init57(_dafny.Zero)
			_ = _element0_57
			_nw324 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
			(_nw324).ArraySet1(_element0_57, 0)
			var _nativeLen0_57 = (_len0_57).Int()
			_ = _nativeLen0_57
			for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
				(_nw324).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
			}
		}
		_2139_v98 = _nw324
		var _2144_v99 _dafny.Map
		_ = _2144_v99
		_2144_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2005_v1)
		var _2145_v100 _dafny.Sequence
		_ = _2145_v100
		_2145_v100 = _dafny.SeqOf(_dafny.SeqOf(_2005_v1, (_2144_v99).Cardinality()))
		var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_2139_v98), 0))
		_ = _index322
		(_2139_v98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2145_v100, _2145_v100), (_index322).Int())
		var _2146_v101 D15
		_ = _2146_v101
		_2146_v101 = Companion_D15_.Create_DC51_(p0, false)
		var _2147_v102 _dafny.CodePoint
		_ = _2147_v102
		_2147_v102 = _dafny.CodePoint('q')
		var _2148_v104 _dafny.Map
		_ = _2148_v104
		_2148_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, func() _dafny.Map {
			var _coll77 = _dafny.NewMapBuilder()
			_ = _coll77
			for _iter81 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-952))).Elements()); ; {
				_compr_77, _ok81 := _iter81()
				if !_ok81 {
					break
				}
				var _2149_v103 _dafny.Int
				_2149_v103 = interface{}(_compr_77).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-952)), _2149_v103) {
					_coll77.Add((_2149_v103).Times(_2005_v1), _2005_v1)
				}
			}
			return _coll77.ToMap()
		}())
		var _2150_v105 D0
		_ = _2150_v105
		_2150_v105 = Companion_D0_.Create_DC1_(_2005_v1, p0, _2148_v104, (_2091_v67).Cardinality(), _2147_v102)
		var _2151_v106 _dafny.Sequence
		_ = _2151_v106
		_2151_v106 = _dafny.SeqOf(_2150_v105, _2150_v105)
		var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_2139_v98), 0))
		_ = _index323
		(_2139_v98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm54(_2146_v101, _2005_v1, _2147_v102, _dafny.IntOfUint32((_2151_v106).Cardinality()), globalState), _dafny.SeqOf(_2041_v28)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm54(_2146_v101, _2005_v1, _2147_v102, _dafny.IntOfUint32((_2151_v106).Cardinality()), globalState), _dafny.SeqOf(_2041_v28))).Cardinality()))).Uint32(), _2041_v28), _2145_v100), (_index323).Int())
		r0 = p0
		var _len0_58 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_58
		var _nw325 _dafny.Array
		_ = _nw325
		if _len0_58.Cmp(_dafny.Zero) == 0 {
			_nw325 = _dafny.NewArray(_len0_58)
		} else {
			var _init58 func(_dafny.Int) bool = (func(_2152_p0 bool) func(_dafny.Int) bool {
				return func(_2153_i18 _dafny.Int) bool {
					return _2152_p0
				}
			})(p0)
			_ = _init58
			var _element0_58 = _init58(_dafny.Zero)
			_ = _element0_58
			_nw325 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
			(_nw325).ArraySet1(_element0_58, 0)
			var _nativeLen0_58 = (_len0_58).Int()
			_ = _nativeLen0_58
			for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
				(_nw325).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
			}
		}
		r1 = _nw325
		var _2154_v107 _dafny.Map
		_ = _2154_v107
		_2154_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2041_v28, _2147_v102)
		r2 = (_2154_v107).Merge(_2154_v107)
		r3 = _2005_v1
		return r0, r1, r2, r3
	}
}
func (_this *C7) F23() _dafny.Set {
	{
		return _this._f23
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
